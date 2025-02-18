package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// MBTIProfile holds detailed information for each personality type.
type MBTIProfile struct {
	Name        string
	ShortDesc   string
	Strengths   string
	Blindspots  string
	Likes       string
	HowTheyWork string
}

// Detailed profiles for each MBTI type.
// These summaries are synthesized from your document. Adjust them as needed.
var mbtiProfiles = map[string]MBTIProfile{
	// Analysts
	"INTJ": {
		Name:        "The Architect",
		ShortDesc:   "Strategic, self-confident, and decisive.",
		Strengths:   "being hard-working, dedicated, a quick learner, thorough, and holding high standards",
		Blindspots:  "appearing aloof, overly critical, and inflexible",
		Likes:       "intellectual challenges, innovation, and deep analytical conversations",
		HowTheyWork: "by planning meticulously, working independently, and focusing on long-term goals",
	},
	"INTP": {
		Name:        "The Logician",
		ShortDesc:   "Inventive, analytical, and theoretical.",
		Strengths:   "creative thinking, logical analysis, and connecting abstract ideas",
		Blindspots:  "appearing overly critical, emotionally detached, and sometimes indecisive",
		Likes:       "exploring theories, engaging in debates, and independent research",
		HowTheyWork: "in autonomous settings that allow for deep analysis and creative problem solving",
	},
	"ENTJ": {
		Name:        "The Commander",
		ShortDesc:   "Bold, decisive, and strategic.",
		Strengths:   "efficiency, charisma, and natural leadership",
		Blindspots:  "impatience, dominating behavior, and insensitivity to emotional cues",
		Likes:       "leadership challenges, structured environments, and clear, goal-driven projects",
		HowTheyWork: "by setting clear goals, taking charge, and driving teams toward objectives",
	},
	"ENTP": {
		Name:        "The Debater",
		ShortDesc:   "Innovative, energetic, and provocative.",
		Strengths:   "quick wit, adaptability, and generating creative ideas",
		Blindspots:  "inconsistencies, being overly argumentative, and occasional impracticality",
		Likes:       "brainstorming, debating concepts, and exploring unconventional solutions",
		HowTheyWork: "in dynamic environments where spontaneity and agile thinking are rewarded",
	},
	// Diplomats
	"INFJ": {
		Name:        "The Advocate",
		ShortDesc:   "Empathetic, insightful, and principled.",
		Strengths:   "deep intuition, a caring nature, and visionary insight",
		Blindspots:  "over-idealism, reservedness, and sometimes missing practical details",
		Likes:       "meaningful connections, creative expression, and thoughtful dialogue",
		HowTheyWork: "when their actions align with core values and a long-term purpose",
	},
	"INFP": {
		Name:        "The Mediator",
		ShortDesc:   "Idealistic, empathetic, and reflective.",
		Strengths:   "compassion, creativity, and loyalty",
		Blindspots:  "over-idealism, conflict avoidance, and occasional indecision",
		Likes:       "artistic pursuits, authenticity, and exploring personal values",
		HowTheyWork: "in a caring and creative manner that values authenticity over strict efficiency",
	},
	"ENFJ": {
		Name:        "The Protagonist",
		ShortDesc:   "Charismatic, empathetic, and inspiring.",
		Strengths:   "excellent communication, supportive leadership, and vision",
		Blindspots:  "tendency to overextend, neglect personal needs, and be overly idealistic",
		Likes:       "teamwork, personal growth, and fostering community",
		HowTheyWork: "by inspiring others, building consensus, and balancing multiple perspectives",
	},
	"ENFP": {
		Name:        "The Campaigner",
		ShortDesc:   "Enthusiastic, creative, and sociable.",
		Strengths:   "optimism, innovative ideas, and an engaging personality",
		Blindspots:  "difficulty with follow-through, distractibility, and heightened emotionality",
		Likes:       "exploring new possibilities, creative expression, and spontaneous interactions",
		HowTheyWork: "in flexible environments that encourage creative risk-taking",
	},
	// Sentinels
	"ISTJ": {
		Name:        "The Logistician",
		ShortDesc:   "Reliable, practical, and methodical.",
		Strengths:   "organization, responsibility, and meticulous attention to detail",
		Blindspots:  "rigidity, resistance to change, and being overly rule-bound",
		Likes:       "clear structure, consistency, and tradition",
		HowTheyWork: "by adhering to systematic processes and valuing efficiency and order",
	},
	"ISFJ": {
		Name:        "The Defender",
		ShortDesc:   "Caring, loyal, and meticulous.",
		Strengths:   "supportiveness, conscientiousness, and attentiveness to others’ needs",
		Blindspots:  "modesty that may lead to self-neglect and resistance to change",
		Likes:       "stable environments, harmony, and nurturing relationships",
		HowTheyWork: "by working diligently behind the scenes to maintain stability and support",
	},
	"ESTJ": {
		Name:        "The Executive",
		ShortDesc:   "Decisive, organized, and pragmatic.",
		Strengths:   "strong leadership, clear communication, and robust organizational skills",
		Blindspots:  "inflexibility, a tendency to be overly controlling, and dismissiveness of alternative views",
		Likes:       "structured settings, clear responsibilities, and order",
		HowTheyWork: "by driving projects forward with clear strategies and decisive action",
	},
	"ESFJ": {
		Name:        "The Consul",
		ShortDesc:   "Social, nurturing, and responsible.",
		Strengths:   "warmth, empathy, and excellent interpersonal skills",
		Blindspots:  "over-sensitivity, people-pleasing, and conflict avoidance",
		Likes:       "community, social gatherings, and maintaining harmony",
		HowTheyWork: "by excelling in collaborative environments that value interpersonal connection",
	},
	// Explorers
	"ISTP": {
		Name:        "The Virtuoso",
		ShortDesc:   "Practical, adaptable, and resourceful.",
		Strengths:   "practical problem-solving, logical thinking, and efficiency in hands-on tasks",
		Blindspots:  "detachment, risk-taking, and occasional neglect of long-term planning",
		Likes:       "freedom, technical challenges, and novel experiences",
		HowTheyWork: "by adopting a practical, do-it-yourself attitude in dynamic settings",
	},
	"ISFP": {
		Name:        "The Adventurer",
		ShortDesc:   "Artistic, gentle, and spontaneous.",
		Strengths:   "creativity, empathy, and adaptability",
		Blindspots:  "over-sensitivity, conflict avoidance, and procrastination",
		Likes:       "art, nature, and personal freedom",
		HowTheyWork: "in a relaxed, intuitive manner that favors aesthetic and personal expression",
	},
	"ESTP": {
		Name:        "The Entrepreneur",
		ShortDesc:   "Bold, energetic, and pragmatic.",
		Strengths:   "action orientation, adaptability, and decisiveness",
		Blindspots:  "impulsiveness, overlooking details, and taking unnecessary risks",
		Likes:       "exciting challenges, rapid results, and hands-on problem solving",
		HowTheyWork: "by thriving in fast-paced environments and taking immediate action",
	},
	"ESFP": {
		Name:        "The Entertainer",
		ShortDesc:   "Lively, outgoing, and spontaneous.",
		Strengths:   "excellent people skills, creativity, and charm",
		Blindspots:  "avoidance of long-term planning, distractibility, and high sensitivity to criticism",
		Likes:       "fun social events, creative activities, and dynamic interactions",
		HowTheyWork: "by performing best in interactive, energetic settings that reward spontaneity",
	},
}

// Question represents a single test question.
type Question struct {
	ID       int
	Question string
	OptionA  string
	OptionB  string
	MappingA string // letter for option A
	MappingB string // letter for option B
}

// testQuestions contains all 70 questions.
var testQuestions = []Question{
	{1, "At a party do you:", "Interact with many, including strangers", "Interact with a few, known to you", "E", "I"},
	{2, "Are you more:", "Realistic than speculative", "Speculative than realistic", "S", "N"},
	{3, "Is it worse to:", "Have your head in the clouds", "Be in a rut", "N", "S"},
	{4, "Are you more impressed by:", "Principles", "Emotions", "T", "F"},
	{5, "Are you more drawn toward the:", "Convincing", "Touching", "T", "F"},
	{6, "Do you prefer to work:", "To deadlines", "Just 'whenever'", "J", "P"},
	{7, "Do you tend to choose:", "Rather carefully", "Somewhat impulsively", "J", "P"},
	{8, "At parties do you:", "Stay late, with increasing energy", "Leave early with decreased energy", "E", "I"},
	{9, "Are you more attracted to:", "Sensible people", "Imaginative people", "S", "N"},
	{10, "Are you more interested in:", "What is actual", "What is possible", "S", "N"},
	{11, "In judging others are you more swayed by:", "Laws than circumstances", "Circumstances than laws", "T", "F"},
	{12, "In approaching others is your inclination to be somewhat:", "Objective", "Personal", "T", "F"},
	{13, "Are you more:", "Punctual", "Leisurely", "J", "P"},
	{14, "Does it bother you more having things:", "Incomplete", "Completed", "J", "P"},
	{15, "In your social groups do you:", "Keep abreast of others' happenings", "Get behind on the news", "E", "I"},
	{16, "In doing ordinary things are you more likely to:", "Do it the usual way", "Do it your own way", "J", "P"},
	{17, "Writers should:", "Say what they mean and mean what they say", "Express things more by use of analogy", "T", "F"},
	{18, "Which appeals to you more:", "Consistency of thought", "Harmonious human relationships", "T", "F"},
	{19, "Are you more comfortable in making:", "Logical judgments", "Value judgments", "T", "F"},
	{20, "Do you want things:", "Settled and decided", "Unsettled and undecided", "J", "P"},
	{21, "Would you say you are more:", "Serious and determined", "Easy-going", "T", "F"},
	{22, "In phoning do you:", "Rarely question that it will all be said", "Rehearse what you'll say", "E", "I"},
	{23, "Facts:", "Speak for themselves", "Illustrate principles", "T", "F"},
	{24, "Are visionaries:", "Somewhat annoying", "Rather fascinating", "S", "N"},
	{25, "Are you more often:", "A cool-headed person", "A warm-hearted person", "T", "F"},
	{26, "Is it worse to be:", "Unjust", "Merciless", "T", "F"},
	{27, "Should one usually let events occur:", "By careful selection and choice", "Randomly and by chance", "J", "P"},
	{28, "Do you feel better about:", "Having purchased", "Having the option to buy", "T", "F"},
	{29, "In company do you:", "Initiate conversation", "Wait to be approached", "E", "I"},
	{30, "Common sense is:", "Rarely questionable", "Frequently questionable", "T", "F"},
	{31, "Children often do not:", "Make themselves useful enough", "Exercise their fantasy enough", "S", "N"},
	{32, "In making decisions do you feel more comfortable with:", "Standards", "Feelings", "T", "F"},
	{33, "Are you more:", "Firm than gentle", "Gentle than firm", "T", "F"},
	{34, "Which is more admirable:", "The ability to organize and be methodical", "The ability to adapt and make do", "J", "P"},
	{35, "Do you put more value on:", "Infinite", "Open-minded", "T", "F"},
	{36, "Does new and non-routine interaction with others:", "Stimulate and energize you", "Tax your reserves", "E", "I"},
	{37, "Are you more frequently:", "A practical sort of person", "A fanciful sort of person", "S", "N"},
	{38, "Are you more likely to:", "See how others are useful", "See how others see", "T", "F"},
	{39, "Which is more satisfying:", "To discuss an issue thoroughly", "To arrive at agreement on an issue", "T", "F"},
	{40, "Which rules you more:", "Your head", "Your heart", "T", "F"},
	{41, "Are you more comfortable with work that is:", "Contracted", "Done on a casual basis", "J", "P"},
	{42, "Do you tend to look for:", "The orderly", "Whatever turns up", "J", "P"},
	{43, "Do you prefer:", "Many friends with brief contact", "A few friends with more lengthy contact", "E", "I"},
	{44, "Do you go more by:", "Facts", "Principles", "T", "F"},
	{45, "Are you more interested in:", "Production and distribution", "Design and research", "S", "N"},
	{46, "Which seems the greater error:", "To be too logical", "To be too sentimental", "T", "F"},
	{47, "Do you value in yourself more that you are:", "Unwavering", "Devoted", "T", "F"},
	{48, "Do you more often prefer the:", "Final and unalterable statement", "Tentative and preliminary statement", "J", "P"},
	{49, "Are you more comfortable:", "After a decision", "Before a decision", "J", "P"},
	{50, "Do you:", "Speak easily and at length with strangers", "Find little to say to strangers", "E", "I"},
	{51, "Are you more likely to trust your:", "Experience", "Hunch", "S", "N"},
	{52, "Do you feel:", "More practical than ingenious", "More ingenious than practical", "S", "N"},
	{53, "Which person is more to be complimented – one of:", "Clear reason", "Strong feeling", "T", "F"},
	{54, "Are you inclined more to be:", "Fair-minded", "Sympathetic", "T", "F"},
	{55, "Is it preferable mostly to:", "Make sure things are arranged", "Just let things happen", "J", "P"},
	{56, "In relationships should most things be:", "Re-negotiable", "Random and circumstantial", "J", "P"},
	{57, "When the phone rings do you:", "Hasten to get to it first", "Hope someone else will answer", "E", "I"},
	{58, "Do you prize more in yourself:", "A strong sense of reality", "A vivid imagination", "S", "N"},
	{59, "Are you drawn more to:", "Fundamentals", "Overtones", "S", "N"},
	{60, "Which seems the greater error:", "To be too passionate", "To be too objective", "T", "F"},
	{61, "Do you see yourself as basically:", "Hard-headed", "Soft-hearted", "T", "F"},
	{62, "Which situation appeals to you more:", "The structured and scheduled", "The unstructured and unscheduled", "J", "P"},
	{63, "Are you a person that is more:", "Routinized", "Whimsical", "J", "P"},
	{64, "Are you more inclined to be:", "Easy to approach", "Somewhat reserved", "E", "I"},
	{65, "In writings do you prefer:", "The more literal", "The more figurative", "S", "N"},
	{66, "Is it harder for you to:", "Identify with others", "Utilize others", "F", "T"},
	{67, "Which do you wish more for yourself:", "Clarity of reason", "Strength of compassion", "T", "F"},
	{68, "Which is the greater fault:", "Being indiscriminate", "Being critical", "T", "F"},
	{69, "Do you prefer the:", "Planned event", "Unplanned event", "J", "P"},
	{70, "Do you tend to be more:", "Deliberate than spontaneous", "Spontaneous than deliberate", "J", "P"},
}

// ------------------ Handler Functions ------------------

// Interaction Explorer Handler (Home Page)
func interactionFormHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		MBTITypes []string
	}{MBTITypes: mbtiTypes}
	if err := interactionFormTmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Interaction Result Handler
func interactionResultHandler(w http.ResponseWriter, r *http.Request) {
	type1 := strings.ToUpper(strings.TrimSpace(r.FormValue("person1")))
	type2 := strings.ToUpper(strings.TrimSpace(r.FormValue("person2")))
	profile1, ok1 := mbtiProfiles[type1]
	if !ok1 {
		profile1 = MBTIProfile{Name: "Unknown", ShortDesc: "No data available.", Strengths: "N/A", Blindspots: "N/A", Likes: "N/A", HowTheyWork: "N/A"}
	}
	profile2, ok2 := mbtiProfiles[type2]
	if !ok2 {
		profile2 = MBTIProfile{Name: "Unknown", ShortDesc: "No data available.", Strengths: "N/A", Blindspots: "N/A", Likes: "N/A", HowTheyWork: "N/A"}
	}
	interaction := getInteractionDetail(type1, type2)
	data := struct {
		Type1, Type2 string
		Profile1     MBTIProfile
		Profile2     MBTIProfile
		Interaction  string
	}{
		Type1:       type1,
		Profile1:    profile1,
		Type2:       type2,
		Profile2:    profile2,
		Interaction: interaction,
	}
	if err := interactionResultTmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Myers Briggs Test Handler
func testHandler(w http.ResponseWriter, r *http.Request) {
	if err := testTmpl.Execute(w, testQuestions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Myers Briggs Test Result Handler
func testResultHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse test responses", http.StatusBadRequest)
		return
	}
	// Count responses for each dichotomy.
	dimensions := map[string]int{"E": 0, "I": 0, "S": 0, "N": 0, "T": 0, "F": 0, "J": 0, "P": 0}
	for _, q := range testQuestions {
		ans := r.FormValue("q" + strconv.Itoa(q.ID))
		dimensions[ans]++
	}
	mbti := ""
	if dimensions["E"] >= dimensions["I"] {
		mbti += "E"
	} else {
		mbti += "I"
	}
	if dimensions["S"] >= dimensions["N"] {
		mbti += "S"
	} else {
		mbti += "N"
	}
	if dimensions["T"] >= dimensions["F"] {
		mbti += "T"
	} else {
		mbti += "F"
	}
	if dimensions["J"] >= dimensions["P"] {
		mbti += "J"
	} else {
		mbti += "P"
	}
	profile, exists := mbtiProfiles[mbti]
	data := struct {
		MBTIType string
		Profile  *MBTIProfile
	}{
		MBTIType: mbti,
	}
	if exists {
		data.Profile = &profile
	}
	if err := testResultTmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ------------------ Templates ------------------

// Define the list of MBTI types (for selection boxes).
var mbtiTypes = []string{
	"ISTJ", "ISFJ", "INFJ", "INTJ",
	"ISTP", "ISFP", "INFP", "INTP",
	"ESTP", "ESFP", "ENFP", "ENTP",
	"ESTJ", "ESFJ", "ENFJ", "ENTJ",
}

// Template for the Interaction Explorer Form (Home Page)
var interactionFormTmpl = template.Must(template.New("interactionForm").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>MBTI Interaction Explorer</title>
    <style>
      body { background: #f0f2f5; font-family: Arial, sans-serif; margin: 0; padding: 20px; color: #333; }
      .container { max-width: 1000px; margin: 40px auto; background: #fff; border-radius: 8px; padding: 30px; box-shadow: 0 2px 8px rgba(0,0,0,0.1); }
      h1 { text-align: center; margin-bottom: 10px; color: #007BFF; }
      p.description { text-align: center; margin-bottom: 40px; font-size: 16px; }
      .button-link { text-align: center; margin-bottom: 30px; }
      .button-link a { background-color: #28a745; color: #fff; padding: 10px 20px; text-decoration: none; border-radius: 8px; font-size: 16px; }
      .flex-container { display: flex; gap: 40px; flex-wrap: wrap; }
      .mbti-group { flex: 1; min-width: 300px; }
      .mbti-group h2 { margin-bottom: 20px; }
      .your-group h2 { color: #007BFF; }
      .their-group h2 { color: #FF8C00; }
      .grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 15px; }
      .box { border: 2px solid #ddd; border-radius: 8px; padding: 20px; text-align: center; cursor: pointer; transition: border-color 0.3s, background-color 0.3s; }
      .your-group .box:hover { border-color: #007BFF; background-color: #e7f1ff; }
      .your-group .box.selected { border-color: #007BFF; background-color: #e7f1ff; }
      .their-group .box:hover { border-color: #FF8C00; background-color: #fff0e6; }
      .their-group .box.selected { border-color: #FF8C00; background-color: #fff0e6; }
      .box input { display: none; }
      .submit-btn { display: block; width: 100%; padding: 15px; font-size: 18px; border: none; background-color: #007BFF; color: #fff; border-radius: 8px; cursor: pointer; transition: background-color 0.3s; margin-top: 20px; }
      .submit-btn:hover { background-color: #0056b3; }
    </style>
  </head>
  <body>
    <div class="container">
      <h1>MBTI Interaction Explorer</h1>
      <p class="description">Select your personality type and your counterpart's type to reveal detailed insights on how you interact.</p>
      <div class="button-link">
        <a href="/test">Take the Full Myers Briggs Test</a>
      </div>
      <form method="GET" action="/result">
        <div class="flex-container">
          <div class="mbti-group your-group">
            <h2>Your Personality</h2>
            <div class="grid">
              {{range .MBTITypes}}
                <label class="box">
                  <input type="radio" name="person1" value="{{.}}" required>
                  <span>{{.}}</span>
                </label>
              {{end}}
            </div>
          </div>
          <div class="mbti-group their-group">
            <h2>Their Personality</h2>
            <div class="grid">
              {{range .MBTITypes}}
                <label class="box">
                  <input type="radio" name="person2" value="{{.}}" required>
                  <span>{{.}}</span>
                </label>
              {{end}}
            </div>
          </div>
        </div>
        <input type="submit" value="Discover Interaction" class="submit-btn" />
      </form>
    </div>
    <script>
      const boxes = document.querySelectorAll('.box');
      boxes.forEach(box => {
        const input = box.querySelector('input');
        input.addEventListener('change', () => {
          const groupName = input.name;
          document.querySelectorAll('input[name="' + groupName + '"]').forEach(i => { i.parentElement.classList.remove('selected'); });
          if(input.checked){ box.classList.add('selected'); }
        });
      });
    </script>
  </body>
</html>
`))

// Template for Interaction Result
var interactionResultTmpl = template.Must(template.New("result").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>MBTI Interaction Result</title>
    <style>
      body { background: #f0f2f5; font-family: Arial, sans-serif; margin: 0; padding: 20px; color: #333; }
      .container { max-width: 1000px; margin: 40px auto; background: #fff; border-radius: 8px; padding: 30px; box-shadow: 0 2px 8px rgba(0,0,0,0.1); }
      h1 { text-align: center; margin-bottom: 40px; color: #007BFF; }
      .flex-container { display: flex; gap: 40px; flex-wrap: wrap; }
      .result-box { flex: 1; min-width: 300px; border: 2px solid #ddd; border-radius: 8px; padding: 20px; }
      .your-result h2 { color: #007BFF; }
      .their-result h2 { color: #FF8C00; }
      .interaction-box { border: 2px dashed #666; border-radius: 8px; padding: 20px; margin-top: 40px; background-color: #fffbea; }
      .interaction-box h2 { color: #333; }
      .back-link { display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #007BFF; color: #fff; text-decoration: none; border-radius: 8px; transition: background-color 0.3s; }
      .back-link:hover { background-color: #0056b3; }
      .description { white-space: pre-wrap; }
    </style>
  </head>
  <body>
    <div class="container">
      <h1>Interaction Result</h1>
      <div class="flex-container">
        <div class="result-box your-result">
          <h2>Your Personality ({{.Type1}} – {{.Profile1.Name}})</h2>
          <p><strong>Short Description:</strong> {{.Profile1.ShortDesc}}</p>
          <p><strong>Strengths:</strong> {{.Profile1.Strengths}}</p>
          <p><strong>Blindspots:</strong> {{.Profile1.Blindspots}}</p>
          <p><strong>Likes:</strong> {{.Profile1.Likes}}</p>
          <p><strong>How They Work:</strong> {{.Profile1.HowTheyWork}}</p>
        </div>
        <div class="result-box their-result">
          <h2>Their Personality ({{.Type2}} – {{.Profile2.Name}})</h2>
          <p><strong>Short Description:</strong> {{.Profile2.ShortDesc}}</p>
          <p><strong>Strengths:</strong> {{.Profile2.Strengths}}</p>
          <p><strong>Blindspots:</strong> {{.Profile2.Blindspots}}</p>
          <p><strong>Likes:</strong> {{.Profile2.Likes}}</p>
          <p><strong>How They Work:</strong> {{.Profile2.HowTheyWork}}</p>
        </div>
      </div>
      <div class="interaction-box">
        <h2>How You Interact</h2>
        <p class="description">{{.Interaction}}</p>
      </div>
      <a href="/" class="back-link">Back</a>
    </div>
  </body>
</html>
`))

// Template for Myers Briggs Test
var testTmpl = template.Must(template.New("test").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Myers Briggs Test</title>
    <style>
      body { background: #f0f2f5; font-family: Arial, sans-serif; margin: 0; padding: 20px; color: #333; }
      .container { max-width: 900px; margin: 40px auto; background: #fff; border-radius: 8px; padding: 30px; box-shadow: 0 2px 8px rgba(0,0,0,0.1); }
      h1 { text-align: center; margin-bottom: 20px; color: #28a745; }
      form { margin-top: 20px; }
      .question { margin-bottom: 20px; }
      .question p { font-weight: bold; }
      .options { margin-top: 10px; }
      .options label { display: block; margin-bottom: 10px; }
      .submit-btn { display: block; width: 100%; padding: 15px; font-size: 18px; border: none; background-color: #28a745; color: #fff; border-radius: 8px; cursor: pointer; transition: background-color 0.3s; margin-top: 20px; }
      .submit-btn:hover { background-color: #218838; }
      .back-link { display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #007BFF; color: #fff; text-decoration: none; border-radius: 8px; transition: background-color 0.3s; }
      .back-link:hover { background-color: #0056b3; }
    </style>
  </head>
  <body>
    <div class="container">
      <h1>Myers Briggs Test</h1>
      <form method="POST" action="/testresult">
        {{range .}}
          <div class="question">
            <p>Question {{.ID}}: {{.Question}}</p>
            <div class="options">
              <label>
                <input type="radio" name="q{{.ID}}" value="{{.MappingA}}" required>
                {{.OptionA}}
              </label>
              <label>
                <input type="radio" name="q{{.ID}}" value="{{.MappingB}}" required>
                {{.OptionB}}
              </label>
            </div>
          </div>
        {{end}}
        <input type="submit" value="Submit Test" class="submit-btn" />
      </form>
      <a href="/" class="back-link">Back</a>
    </div>
  </body>
</html>
`))

// Template for Myers Briggs Test Result
var testResultTmpl = template.Must(template.New("testresult").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Myers Briggs Test Result</title>
    <style>
      body { background: #f0f2f5; font-family: Arial, sans-serif; margin: 0; padding: 20px; color: #333; }
      .container { max-width: 900px; margin: 40px auto; background: #fff; border-radius: 8px; padding: 30px; box-shadow: 0 2px 8px rgba(0,0,0,0.1); text-align: center; }
      h1 { color: #28a745; margin-bottom: 20px; }
      .result { font-size: 24px; margin-bottom: 20px; }
      .profile { text-align: left; margin-top: 20px; }
      .profile p { margin: 5px 0; }
      .back-link { display: inline-block; margin-top: 20px; padding: 10px 20px; background-color: #007BFF; color: #fff; text-decoration: none; border-radius: 8px; transition: background-color 0.3s; }
      .back-link:hover { background-color: #0056b3; }
    </style>
  </head>
  <body>
    <div class="container">
      <h1>Test Result</h1>
      <div class="result">Your Myers Briggs Type is: <strong>{{.MBTIType}}</strong></div>
      {{with .Profile}}
      <div class="profile">
        <p><strong>{{.Name}}</strong> ({{$.MBTIType}})</p>
        <p><strong>Description:</strong> {{.ShortDesc}}</p>
        <p><strong>Strengths:</strong> {{.Strengths}}</p>
        <p><strong>Blindspots:</strong> {{.Blindspots}}</p>
        <p><strong>Likes:</strong> {{.Likes}}</p>
        <p><strong>How They Work:</strong> {{.HowTheyWork}}</p>
      </div>
      {{end}}
      <a href="/" class="back-link">Back to Home</a>
    </div>
  </body>
</html>
`))

// ------------------ getInteractionDetail Function ------------------

// getInteractionDetail provides a detailed, objective interaction analysis.
// It includes specific analyses for key pairings based on document insights,
// and falls back to a generic comparison for other pairings.
func getInteractionDetail(type1, type2 string) string {
	profile1, ok1 := mbtiProfiles[type1]
	profile2, ok2 := mbtiProfiles[type2]
	if !ok1 || !ok2 {
		return "Detailed interaction information is not available for one or both personality types."
	}

	// Specific detailed interactions based on document insights:
	if (type1 == "ENFJ" && type2 == "ENTJ") || (type1 == "ENTJ" && type2 == "ENFJ") {
		return "When an ENFJ meets an ENTJ, the ENFJ’s nurturing, people-centered approach meets the ENTJ’s results-driven, strategic mindset. The document notes that ENFJs invest in personal growth and empathetic connections, while ENTJs focus on efficiency and goal achievement. For success, the ENTJ should temper decisiveness with empathy, and the ENFJ should appreciate clear, structured direction."
	} else if (type1 == "ENFJ" && type2 == "ENFP") || (type1 == "ENFP" && type2 == "ENFJ") {
		return "An ENFJ paired with an ENFP creates an energetic, inspiring mix. Both value deep relationships, yet ENFJs bring structure and consistency while ENFPs offer spontaneity and creative energy. The document suggests that mutual support and open communication—where the ENFJ provides guidance and the ENFP injects fresh ideas—can lead to a harmonious and productive interaction."
	} else if (type1 == "ENTJ" && type2 == "ENTP") || (type1 == "ENTP" && type2 == "ENTJ") {
		return "When ENTJs and ENTPs interact, their shared drive for achievement and innovation fuels dynamic debate and idea generation. The ENTJ’s focus on structured, long-term planning complements the ENTP’s talent for challenging conventions and generating new solutions. Their collaboration can be highly effective if both remain open to balancing structure with creative exploration."
	} else if type1 == type2 {
		return "When two " + type1 + " individuals interact, they generally share similar values, work habits, and communication styles. This alignment leads to clear expectations and coordinated efforts. However, similar blindspots may also be reinforced, so it’s important for both to remain open to alternative viewpoints."
	}

	// Generic detailed interaction analysis for all other pairings:
	generic := "Detailed Interaction Analysis:\n\n"
	generic += "Your Profile (" + type1 + " - " + profile1.Name + "):\n"
	generic += "• Strengths: " + profile1.Strengths + ".\n"
	generic += "• Blindspots: " + profile1.Blindspots + ".\n"
	generic += "• Likes: " + profile1.Likes + ".\n"
	generic += "• Working Style: " + profile1.HowTheyWork + ".\n\n"
	generic += "Their Profile (" + type2 + " - " + profile2.Name + "):\n"
	generic += "• Strengths: " + profile2.Strengths + ".\n"
	generic += "• Blindspots: " + profile2.Blindspots + ".\n"
	generic += "• Likes: " + profile2.Likes + ".\n"
	generic += "• Working Style: " + profile2.HowTheyWork + ".\n\n"
	generic += "Comparison & Insights:\n"
	generic += "1. Strengths: Your focus on " + profile1.Strengths + " provides a stable foundation, while they contribute through " + profile2.Strengths + ".\n\n"
	generic += "2. Blindspots: You are cautious about " + profile1.Blindspots + ", whereas they might sometimes overlook these aspects due to " + profile2.Blindspots + ". Open dialogue can help bridge this gap.\n\n"
	generic += "3. Preferences: You value " + profile1.Likes + ", suggesting a need for structure, while they appreciate " + profile2.Likes + ", indicating a preference for flexibility. Finding common ground is key.\n\n"
	generic += "4. Working Styles: Your method of " + profile1.HowTheyWork + " contrasts with their approach of " + profile2.HowTheyWork + ". Aligning your goals and maintaining clear communication can harmonize these differences.\n\n"
	generic += "Overall, although there are clear differences in strengths and approaches, a mutual effort to understand and integrate these factors can foster a balanced, productive relationship."
	return generic
}

func main() {
	http.HandleFunc("/", interactionFormHandler)
	http.HandleFunc("/result", interactionResultHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/testresult", testResultHandler)
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
