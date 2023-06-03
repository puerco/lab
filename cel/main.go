package main

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/puerco/protobom/pkg/sbom"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	NodeListObject    = decls.NewObjectType("puerco.protobom.NodeList")
	NodeListTypeValue = types.NewTypeValue("puerco.protobom.NodeList")
	NodeListType      = cel.ObjectType("puerco.protobom.NodeList")
)

type NodeList struct {
	*sbom.NodeList
}

// ConvertToNative implements ref.Val.ConvertToNative.
func (nl NodeList) ConvertToNative(typeDesc reflect.Type) (interface{}, error) {
	if reflect.TypeOf(nl).AssignableTo(typeDesc) {
		return nl, nil
	} else if reflect.TypeOf(nl.NodeList).AssignableTo(typeDesc) {
		return nl.NodeList, nil
	}
	//if reflect.TypeOf("").AssignableTo(typeDesc) {
	//		return d.URL.String(), nil
	//	}
	return nil, fmt.Errorf("type conversion error from 'NodeList' to '%v'", typeDesc)
}

// ConvertToType implements ref.Val.ConvertToType.
func (nl NodeList) ConvertToType(typeVal ref.Type) ref.Val {
	switch typeVal {
	case NodeListTypeValue:
		return nl
	case types.TypeType:
		return NodeListTypeValue

	}
	return types.NewErr("type conversion error from '%s' to '%s'", NodeListTypeValue, typeVal)
}

// Equal implements ref.Val.Equal.
func (nl NodeList) Equal(other ref.Val) ref.Val {
	// otherDur, ok := other.(NodeList)
	_, ok := other.(NodeList)
	if !ok {
		return types.MaybeNoSuchOverloadErr(other)
	}

	// TODO: Moar tests like:
	// return types.Bool(d.URL.String() == otherDur.URL.String())
	return types.True
}

// Type implements ref.Val.Type.
func (nl NodeList) Type() ref.Type {
	return NodeListTypeValue
}

// Value implements ref.Val.Value.
func (nl NodeList) Value() interface{} {
	return nl
}

/*
type NodeList []*sbom.Node

func (nl *NodeList) ConvertToNative(reflect.Type) (any, error) {
	return nil, nil
}

func (nl *NodeList) ConvertToType(ref.Type) ref.Val {
	return nil
}

func (nl *NodeList) Equal(other ref.Val) ref.Val {
	// Esto no se si esta biwen
	if nl.Type() != other.Type() {
		return types.False
	}
	return types.True
}

func (nl *NodeList) Type() ref.Type {

}

func (nl *NodeList) Value() any {
	return nil
}

*/

func main() {
	// Crea el environment de CEL
	env, err := cel.NewEnv(
		cel.Types(&sbom.Document{}),
		cel.Types(&sbom.NodeList{}),
		//cel.Types(cel.ObjectType("puerco.protobom.Document")),
		//cel.Types(cel.ObjectType("puerco.protobom.NodeList")),
		//ext.NativeTypes(reflect.TypeOf(&sbom.NodeList{})),
		//ext.NativeTypes(reflect.TypeOf(&sbom.NodeList{})),
		cel.Variable("sbom",
			cel.ObjectType("puerco.protobom.Document"),
		),

		cel.Function(
			"files",
			cel.MemberOverload(
				"sbom_files_binding", []*cel.Type{cel.ObjectType("puerco.protobom.Document")}, NodeListType,
				cel.UnaryBinding(
					func(lhs ref.Val) ref.Val {
						nl := NodeList{
							NodeList: &sbom.NodeList{},
						}
						bom, ok := lhs.Value().(*sbom.Document)
						if !ok {
							return types.NewErr("unable to convert sbom to native (wrong type?)")
						}
						nl.Edges = bom.Edges
						for _, n := range bom.Nodes {
							if n.Type == sbom.Node_FILE {
								nl.NodeList.Nodes = append(nl.NodeList.Nodes, n)
							}
						}
						return nl
						// val := types.NewObjectTypeValue("puerco.protobom.NodeList")
						// return &NodeList{}
						// return val
						// return cel.ObjectType("puerco.protobom.NodeList")
					},
				),
			),
		),
	)

	if err != nil {
		logrus.Fatal(fmt.Errorf("Creating SBOM type: %w", err))
	}

	// Compie the CEL program
	// ast, iss := env.Compile(`sbom.metadata.id == 'SPDX-DOCID'`)
	ast, iss := env.Compile(`sbom.files()`)
	if iss.Err() != nil {
		logrus.Fatal(iss.Err())
	}

	program, err := env.Program(ast)
	if err != nil {
		logrus.Fatalf("creating program: %v", err)
	}

	/*
		if !reflect.DeepEqual(ast.OutputType(), cel.BoolType) {
			glog.Exitf(
				"Got %v, wanted %v result type", ast.OutputType(), celType)
		}
	*/
	// fmt.Printf("%s\n\n", strings.ReplaceAll(expr, "\t", " "))
	/*
		varMap, isMap := vars.(map[string]any)
		if !isMap {
			fmt.Printf("(%T)\n", vars)
		} else {
			for k, v := range varMap {
				switch val := v.(type) {
				case proto.Message:
					bytes, err := prototext.Marshal(val)
					if err != nil {
						glog.Exitf("failed to marshal proto to text: %v", val)
					}
					fmt.Printf("%s = %s", k, string(bytes))
				case map[string]any:
					b, _ := json.MarshalIndent(v, "", "  ")
					fmt.Printf("%s = %v\n", k, string(b))
				case uint64:
					fmt.Printf("%s = %vu\n", k, v)
				default:
					fmt.Printf("%s = %v\n", k, v)
				}
			}
		}
	*/
	fmt.Println()

	val := map[string]interface{}{
		"sbom": &sbom.Document{
			Metadata: &sbom.Metadata{
				Id:      "SPDX-DOCID",
				Version: "",
				Name:    "",
				Date:    &timestamppb.Timestamp{},
				Tools:   []*sbom.Tool{},
				Authors: []*sbom.Person{},
				Comment: "",
			},
			RootElements: []string{},
			Nodes: []*sbom.Node{
				{
					Id:                 "package-nginx",
					Type:               sbom.Node_PACKAGE,
					Name:               "nginx",
					Version:            "1.24.0-1.fc38",
					ExternalReferences: []*sbom.ExternalReference{},
				},
				{
					Id:                 "file-nginx-conf",
					Type:               sbom.Node_FILE,
					Name:               "/etc/nginx.conf",
					ExternalReferences: []*sbom.ExternalReference{},
				},
			},
			Edges: []*sbom.Edge{
				{
					Type: sbom.Edge_contains,
					From: "package-nginx",
					To:   []string{"file-nginx-conf"},
				},
			},
		},
	}

	//bytes, err := prototext.Marshal(val)
	//if err != nil {
	//	logrus.Fatalf("failed to marshal proto to text: %v", val)
	//}
	/// fmt.Printf("%s = %s", string(bytes))

	out, det, err := program.Eval(val)
	if err != nil {
		logrus.Fatal(err)
	}
	report(out, det, err)
	fmt.Println()

	//eval(program, request(auth("user:me@acme.co", claims), time.Now()))
	fmt.Println()

	/*
		env, err := cel.NewEnv(
			cel.Variable("name", cel.StringType),
			cel.Variable("group", cel.StringType),
		)

		ast, issues := env.Compile(`name.startsWith("/groups/" + group)`)

		if issues != nil && issues.Err() != nil {
			log.Fatalf("type-check error: %s", issues.Err())
		}
		prg, err := env.Program(ast)
		if err != nil {
			log.Fatalf("program construction error: %s", err)
		}

		// The `out` var contains the output of a successful evaluation.
		// The `details' var would contain intermediate evaluation state if enabled as
		// a cel.ProgramOption. This can be useful for visualizing how the `out` value
		// was arrive at.
		out, details, err := prg.Eval(map[string]interface{}{
			"name":  "/groups/acme.co/documents/secret-stuff",
			"group": "acme.co"})
		fmt.Println(out)     // 'true'
		fmt.Println(details) // 'true'
	*/
}

func report(result ref.Val, details *cel.EvalDetails, err error) {
	fmt.Println("------ result ------")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("value: %v (%T)\n", result.Value(), result)
	}
	if details != nil {
		fmt.Printf("\n------ eval states ------\n")
		state := details.State()
		stateIDs := state.IDs()
		ids := make([]int, len(stateIDs), len(stateIDs))
		for i, id := range stateIDs {
			ids[i] = int(id)
		}
		sort.Ints(ids)
		for _, id := range ids {
			v, found := state.Value(int64(id))
			if !found {
				continue
			}
			fmt.Printf("%d: %v (%T)\n", id, v, v)
		}
	}
}
