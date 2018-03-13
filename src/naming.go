package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mxk/go-sqlite/sqlite3"
	"log"
	"strings"
	"time"
)

// checks if a string is inside an array of strings
func inside(s string, a []string) bool {
	for i := 0; i < len(a); i++ {
		if s == a[i] {
			return true
		}
	}
	return false
}

var punc = []string{".", "!", "?", ",", ";", "/"}
var vowels = []string{"a", "e", "i", "o", "u", "y"}
var consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "z"}
var stopwords = []string{"a", "able", "about", "above", "abst", "accordance", "according", "accordingly", "across", "act", "actually", "added", "adj", "affected", "affecting", "affects", "after", "afterwards", "again", "against", "ah", "all", "almost", "alone", "along", "already", "also", "although", "always", "am", "among", "amongst", "an", "and", "announce", "another", "any", "anybody", "anyhow", "anymore", "anyone", "anything", "anyway", "anyways", "anywhere", "apparently", "approximately", "are", "aren", "arent", "arise", "around", "as", "aside", "ask", "asking", "at", "auth", "available", "away", "awfully", "b", "back", "be", "became", "because", "become", "becomes", "becoming", "been", "before", "beforehand", "begin", "beginning", "beginnings", "begins", "behind", "being", "believe", "below", "beside", "besides", "between", "beyond", "biol", "both", "brief", "briefly", "but", "by", "c", "ca", "came", "can", "cannot", "can't", "cause", "causes", "certain", "certainly", "co", "com", "come", "comes", "contain", "containing", "contains", "could", "couldnt", "d", "date", "did", "didn't", "different", "do", "does", "doesn't", "doing", "done", "don't", "down", "downwards", "due", "during", "e", "each", "ed", "edu", "effect", "eg", "eight", "eighty", "either", "else", "elsewhere", "end", "ending", "enough", "especially", "et", "et-al", "etc", "even", "ever", "every", "everybody", "everyone", "everything", "everywhere", "ex", "except", "f", "far", "few", "ff", "fifth", "first", "five", "fix", "followed", "following", "follows", "for", "former", "formerly", "forth", "found", "four", "from", "further", "furthermore", "g", "gave", "get", "gets", "getting", "give", "given", "gives", "giving", "go", "goes", "gone", "got", "gotten", "h", "had", "happens", "hardly", "has", "hasn't", "have", "haven't", "having", "he", "hed", "hence", "her", "here", "hereafter", "hereby", "herein", "heres", "hereupon", "hers", "herself", "hes", "hi", "hid", "him", "himself", "his", "hither", "home", "how", "howbeit", "however", "hundred", "i", "id", "ie", "if", "i'll", "im", "immediate", "immediately", "importance", "important", "in", "inc", "indeed", "index", "information", "instead", "into", "invention", "inward", "is", "isn't", "it", "itd", "it'll", "its", "itself", "i've", "j", "just", "k", "keep", "keeps", "kept", "kg", "km", "know", "known", "knows", "l", "largely", "last", "lately", "later", "latter", "latterly", "least", "less", "lest", "let", "lets", "like", "liked", "likely", "line", "little", "'ll", "look", "looking", "looks", "ltd", "m", "made", "mainly", "make", "makes", "many", "may", "maybe", "me", "mean", "means", "meantime", "meanwhile", "merely", "mg", "might", "million", "miss", "ml", "more", "moreover", "most", "mostly", "mr", "mrs", "much", "mug", "must", "my", "myself", "n", "na", "name", "namely", "nay", "nd", "near", "nearly", "necessarily", "necessary", "need", "needs", "neither", "never", "nevertheless", "new", "next", "nine", "ninety", "no", "nobody", "non", "none", "nonetheless", "noone", "nor", "normally", "nos", "not", "noted", "nothing", "now", "nowhere", "o", "obtain", "obtained", "obviously", "of", "off", "often", "oh", "ok", "okay", "old", "omitted", "on", "once", "one", "ones", "only", "onto", "or", "ord", "other", "others", "otherwise", "ought", "our", "ours", "ourselves", "out", "outside", "over", "overall", "owing", "own", "p", "page", "pages", "part", "particular", "particularly", "past", "per", "perhaps", "placed", "please", "plus", "poorly", "possible", "possibly", "potentially", "pp", "predominantly", "present", "previously", "primarily", "probably", "promptly", "proud", "provides", "put", "q", "que", "quickly", "quite", "qv", "r", "ran", "rather", "rd", "re", "readily", "really", "recent", "recently", "ref", "refs", "regarding", "regardless", "regards", "related", "relatively", "research", "respectively", "resulted", "resulting", "results", "right", "run", "s", "said", "same", "saw", "say", "saying", "says", "sec", "section", "see", "seeing", "seem", "seemed", "seeming", "seems", "seen", "self", "selves", "sent", "seven", "several", "shall", "she", "shed", "she'll", "shes", "should", "shouldn't", "show", "showed", "shown", "showns", "shows", "significant", "significantly", "similar", "similarly", "since", "six", "slightly", "so", "some", "somebody", "somehow", "someone", "somethan", "something", "sometime", "sometimes", "somewhat", "somewhere", "soon", "sorry", "specifically", "specified", "specify", "specifying", "still", "stop", "strongly", "sub", "substantially", "successfully", "such", "sufficiently", "suggest", "sup", "sure", "t", "take", "taken", "taking", "tell", "tends", "th", "than", "thank", "thanks", "thanx", "that", "that'll", "thats", "that've", "the", "their", "theirs", "them", "themselves", "then", "thence", "there", "thereafter", "thereby", "thered", "therefore", "therein", "there'll", "thereof", "therere", "theres", "thereto", "thereupon", "there've", "these", "they", "theyd", "they'll", "theyre", "they've", "think", "this", "those", "thou", "though", "thoughh", "thousand", "throug", "through", "throughout", "thru", "thus", "til", "tip", "to", "together", "too", "took", "toward", "towards", "tried", "tries", "truly", "try", "trying", "ts", "twice", "two", "u", "un", "under", "unfortunately", "unless", "unlike", "unlikely", "until", "unto", "up", "upon", "ups", "us", "use", "used", "useful", "usefully", "usefulness", "uses", "using", "usually", "v", "value", "various", "'ve", "very", "via", "viz", "vol", "vols", "vs", "w", "want", "wants", "was", "wasnt", "way", "we", "wed", "welcome", "we'll", "went", "were", "werent", "we've", "what", "whatever", "what'll", "whats", "when", "whence", "whenever", "where", "whereafter", "whereas", "whereby", "wherein", "wheres", "whereupon", "wherever", "whether", "which", "while", "whim", "whither", "who", "whod", "whoever", "whole", "who'll", "whom", "whomever", "whos", "whose", "why", "widely", "willing", "wish", "with", "within", "without", "wont", "words", "world", "would", "wouldnt", "www", "x", "y", "yes", "yet", "you", "youd", "you'll", "your", "youre", "yours", "yourself", "yourselves", "you've", "z", "zero"}

type classifier struct {
	wordDb *sql.DB
}

// checks what type of word it is (noun, verb, etc.). The second return value is whether or not the word is english.
// BUG: google
//
func (c classifier) getWordType(word string) (string, bool) {
	rows, err := c.wordDb.Query("SELECT type FROM words WHERE word=? LIMIT(1)", word)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	rows.Next()
	var wordType string
	if err := rows.Scan(&wordType); err != nil {
		return "", false
	}
	return wordType, true
}

// replaces all the punctuation in a string with a space
func (c classifier) cleanPunc(s *string) {
	for i := 0; i < len(punc); i++ {
		*s = strings.Replace(*s, punc[i], " ", -1)
	}
}

// checks if the given string is the name of something from various conditions/algorithms
func (c classifier) isName(word string) bool {
	/*
	 * Check list so far:
	 * 1. Capital first letter
	 * 2. Vowels touching consonants (implying that there are letters which are both vowels and consonants in the word)
	 *     - no more than 4 consonants in a row
	 * 3. It is not in the constant list of stopwords
	 * 4. The word is not in the wordnet dictionary <<< BUG
	 * 4.5 OR: The word is a noun in wordnet. BUG: Plural nouns since that is a common way to start a sentence with a noun (i.e. "Cats are a very nice animal.")
	 *
	 * Issues:
	 * - wordnet has some companies in it (i.e. Google, Medium, Apple)
	 *   sometimes this is simply due to the fact that they are nouns in the english language.
	 */
	if len(word) > 2 && string(word[0]) == strings.ToUpper(string(word[0])) {
		// from now on the word will be all lower, for processing purposes
		word = strings.ToLower(word)
		// make sure that each consonant is either touching a vowel, or in a row of MAX 3 consonants
		for c := 0; c < len(word); c++ {
			// NOTE: we can't assume if it isn't inside consonants it is a vowel since these two checks double as our check that this is a letter
			// check
			if inside(string(word[c]), consonants) {
				i := 1
				for c-i >= 0 || c+i < len(word) {
					if c-i >= 0 {
						if inside(string(word[c-i]), vowels) {
							break
						}
					}
					if c+i < len(word) {
						if inside(string(word[c+i]), vowels) {
							break
						}
					}
					i++
				}
				//fmt.Println(i)
				// check that we did not have to move farther than 3 characters to find a vowel
				if i > 3 {
					return false
				}
			}
		}
		// affirm the word isn't a "stopword"
		if inside(word, stopwords) {
			return false
		}

		// given that the code is still executing by this point, we know that check #2 has passed
		// now check the `wordnet`
		wordType, isEnglish := c.getWordType(word)

		if wordType == "noun" || !isEnglish {
			// also check if the word is plural and re-query
			if word[len(word)-1] == 's' {
				// finally, if the word checks out both plural and singular, it is a name (hopefully)
				wordType, isEnglish = c.getWordType(word[:len(word)-1])
				if wordType == "noun" || !isEnglish {
					return true
				}
			} else {
				// if the word isn't plural, it checks out as a name (hopefully)
				return true
			}
		} else {
			return false
		}
	} else {
		return false
	}
	return false
}

// returns a hashmap of all names mentioned in a piece of text
func (c classifier) findNamesInText(text string) map[string]int {
	c.cleanPunc(&text)
	words := strings.Split(text, " ")
	mentions := map[string]int{}

	for i := 0; i < len(words); i++ {
		//fmt.Println(words[i])
		if c.isName(words[i]) {
			mentions[strings.ToLower(words[i])] += 1
		}
	}

	return mentions
}

func main() {
	// setting up the classifier and database to go with it
	db, err := sql.Open("sqlite3", "wordnet.dict")
	if err != nil {
		log.Fatalf("error opening DB (%s)", err)
	}
	c := classifier{wordDb: db}
	// TESTING CODE:
	text := "Congratulations. I hope Samsung nurtures the technical excellence at Joyent.Please, please continue to develop your public cloud offerings. Having options other that the myopic, me-too, feature-matching, monoculture that is AWS/GCE/Azure is incredibly important.That said, for my use profile, you guys need to work on your price competitiveness. Hopefully Samsung will inject the necessary cash for economies-of-scale."
	t := time.Now()
	c.findNamesInText(text)
	fmt.Println(time.Now().Sub(t))
}
