package com.example;

import org.apache.commons.lang3.RandomUtils;
import org.apache.commons.lang3.StringUtils;
import org.apache.commons.lang3.tuple.Pair;
import org.apache.commons.text.WordUtils;

/**
 * A Shakespearean Insult Generator of the highest caliber.
 * Uses commons-lang3 (SLSA SNAPSHOT!) and commons-text to craft
 * insults worthy of the Globe Theatre.
 */
public class InsultGenerator {

    private static final String[] ADJECTIVES_1 = {
        "artless", "bawdy", "beslubbering", "bootless", "churlish",
        "cockered", "clouted", "craven", "currish", "dankish",
        "dissembling", "droning", "errant", "fawning", "fobbing",
        "froward", "gleeking", "goatish", "gorbellied", "impertinent",
        "infectious", "jarring", "loggerheaded", "lumpish", "mammering",
        "mangled", "mewling", "paunchy", "pribbling", "puking",
        "puny", "qualling", "rank", "reeky", "roguish",
        "ruttish", "saucy", "spleeny", "spongy", "surly",
        "tottering", "unmuzzled", "vain", "venomed", "villainous",
        "warped", "wayward", "weedy", "yeasty"
    };

    private static final String[] ADJECTIVES_2 = {
        "base-court", "bat-fowling", "beef-witted", "beetle-headed",
        "boil-brained", "clay-brained", "common-kissing", "crook-pated",
        "dismal-dreaming", "dizzy-eyed", "dog-hearted", "dread-bolted",
        "earth-vexing", "elf-skinned", "fat-kidneyed", "fen-sucked",
        "flap-mouthed", "fly-bitten", "fool-born", "full-gorged",
        "guts-griping", "half-faced", "hasty-witted", "hedge-born",
        "hell-hated", "idle-headed", "ill-breeding", "ill-nurtured",
        "knotty-pated", "milk-livered", "motley-minded", "onion-eyed",
        "plume-plucked", "pottle-deep", "pox-marked", "reeling-ripe",
        "rough-hewn", "rude-growing", "rump-fed", "shard-borne",
        "sheep-biting", "spur-galled", "swag-bellied", "tardy-gaited",
        "tickle-brained", "toad-spotted", "unchin-snouted", "weather-bitten"
    };

    private static final String[] NOUNS = {
        "apple-john", "baggage", "barnacle", "bladder", "boar-pig",
        "bugbear", "bum-bailey", "canker-blossom", "clack-dish", "clotpole",
        "coxcomb", "codpiece", "death-token", "dewberry", "flap-dragon",
        "flax-wench", "flirt-gill", "foot-licker", "fustilarian", "giglet",
        "gudgeon", "haggard", "harpy", "hedge-pig", "horn-beast",
        "hugger-mugger", "jolthead", "lewdster", "lout", "maggot-pie",
        "malt-worm", "mammet", "measle", "minnow", "miscreant",
        "moldwarp", "mumblecrust", "nut-hook", "pigeon-egg", "pignut",
        "puttock", "pumpion", "ratsbane", "scut", "skainsmate",
        "strumpet", "varlot", "vassal", "whey-face", "wagtail"
    };

    public static Pair<String, String> generateInsult() {
        String adj1 = ADJECTIVES_1[RandomUtils.secure().randomInt(0, ADJECTIVES_1.length)];
        String adj2 = ADJECTIVES_2[RandomUtils.secure().randomInt(0, ADJECTIVES_2.length)];
        String noun = NOUNS[RandomUtils.secure().randomInt(0, NOUNS.length)];

        String insult = String.format("Thou %s, %s %s!", adj1, adj2, noun);
        return Pair.of(insult, noun);
    }

    public static void main(String[] args) {
        String banner = StringUtils.repeat("*", 60);
        String title = StringUtils.center("SHAKESPEAREAN INSULT GENERATOR", 60);
        String subtitle = StringUtils.center("~ Powered by commons-lang3 SLSA SNAPSHOT ~", 60);

        System.out.println(banner);
        System.out.println(title);
        System.out.println(subtitle);
        System.out.println(banner);
        System.out.println();

        int count = 5;
        if (args.length > 0) {
            try {
                count = Integer.parseInt(args[0]);
            } catch (NumberFormatException e) {
                // Default to 5
            }
        }

        System.out.println("Hear ye, hear ye! The Bard doth summon " + count + " insult(s):\n");

        for (int i = 1; i <= count; i++) {
            Pair<String, String> result = generateInsult();
            String insult = result.getLeft();
            String noun = result.getRight();

            // Use commons-text to capitalize words nicely for the dramatic version
            String dramatic = WordUtils.capitalizeFully(insult, ' ', '-');

            System.out.printf("  %d. %s%n", i, insult);
            System.out.printf("     (Dramatic reading: %s)%n%n", dramatic);
        }

        // Bonus: use commons-lang3 to build a dramatic exit
        String exit = StringUtils.center("~ Exeunt, pursued by a bear ~", 60);
        System.out.println(banner);
        System.out.println(exit);
        System.out.println(banner);
    }
}
