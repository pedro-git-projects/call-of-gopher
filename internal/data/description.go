package data

type Description struct {
	StrDescription string `json:"str_description"`
	AppDescription string `json:"app_description"`
	ConDescription string `json:"con_description"`
	IntDescription string `json:"int_description"`
	SizDescription string `json:"siz_descrpition"`
	PowDescription string `json:"pow_description"`
	EduDescription string `json:"edu_description"`
	DexDescription string `json:"dex_description"`
}

func (i *Investigator) SetDescription() {
	i.strenghtDescription()
	i.appearenceDescription()
	i.constitutionDescription()
	i.intelligenceDescription()
	i.sizeDescription()
	i.powerDescription()
	i.dexDescription()
	i.educationDescription()
}

func (i *Investigator) strenghtDescription() {
	switch i.Str >= 0 {
	case i.Str >= 0 && i.Str < 15:
		i.Description.StrDescription = "enfeebled: unable to even stand up of lift a cup of tea."
	case i.Str >= 15 && i.Str < 50:
		i.Description.StrDescription = "puny, weak."
	case i.Str >= 50 && i.Str < 90:
		i.Description.StrDescription = "average, for a human."
	case i.Str >= 90 && i.Str < 99:
		i.Description.StrDescription = "certainly one of the strongest people most met."
	case i.Str >= 99:
		i.Description.SizDescription = "world class. You represent the peak of the human strengh."
	}
}

func (i *Investigator) appearenceDescription() {
	switch i.App >= 0 {
	case i.App >= 0 && i.App < 15:
		i.Description.AppDescription = "so unsightly others are affected by fear, revulsion or pity."
	case i.App >= 15 && i.App < 50:
		i.Description.AppDescription = "ugly, possibly difigured due to injury or birth."
	case i.App >= 50 && i.App < 90:
		i.Description.AppDescription = "average."
	case i.App >= 90 && i.App < 99:
		i.Description.AppDescription = "naturally magnetic, one of the most beautiful people most will meet."
	case i.App >= 99:
		i.Description.AppDescription = "the height of glamour and cool, probably a supermodel or film star. Human limit."
	}
}

func (i *Investigator) constitutionDescription() {
	switch i.Con >= 0 {
	case i.Con == 0:
		i.Description.ConDescription = "dead."
	case i.Con >= 0 && i.Con < 15:
		i.Description.ConDescription = "sickly, prone to prolonged illness and probably unable to operate without assistance."
	case i.Con >= 15 && i.Con < 50:
		i.Description.ConDescription = "weak health, you're prone to bouts of ill health, great propensity for feeling pain."
	case i.Con >= 50 && i.Con < 90:
		i.Description.ConDescription = "average health."
	case i.Con >= 90 && i.Con < 99:
		i.Description.ConDescription = "shrugs off colds, hardy and hale."
	case i.Con >= 99:
		i.Description.ConDescription = "iron constitution, able to withstand great amounts of pain."
	}
}

func (i *Investigator) intelligenceDescription() {
	switch i.Int >= 0 {
	case i.Int == 0:
		i.Description.IntDescription = "no intellect. Unable to comprehend the world around you."
	case i.Int >= 1 && i.Int < 50:
		i.Description.IntDescription = "slow learner, able to undertake only the most basic math, or read beginner-level books."
	case i.Int >= 50 && i.Int < 90:
		i.Description.IntDescription = "average human intellect."
	case i.Int >= 90 && i.Int < 99:
		i.Description.IntDescription = "quick-witted, probably able to comprehend multiple languages or theorems."
	case i.Int >= 99:
		i.Description.IntDescription = "genius, comparable to Tesla. This is the limit of the human intellect."
	}
}

func (i *Investigator) sizeDescription() {
	switch i.Siz >= 0 {
	case i.Siz == 1:
		i.Description.SizDescription = "baby (1 - 12lbs)."
	case i.Siz > 1 && i.Siz <= 15:
		i.Description.SizDescription = "child or someone of very short stature, like a dwarf.(33lbs/15kg)."
	case i.Siz > 15 && i.Siz <= 65:
		i.Description.SizDescription = "average human size (moderate weight and height) (170lbs/75kg)."
	case i.Siz > 65 && i.Siz <= 80:
		i.Description.SizDescription = "very tall, strongly built, or obese (240lbs/150 kg)."
	case i.Siz > 80 && i.Siz <= 99:
		i.Description.SizDescription = "oversized in some aspect (330lbs/150kg)."
	}
}

func (i *Investigator) powerDescription() {
	switch i.Pow >= 0 {
	case i.Pow == 0:
		i.Description.PowDescription = "enfeebled minded, no willpower, no magical potential."
	case i.Pow >= 1 && i.Pow <= 15:
		i.Description.PowDescription = "weak-willed, easily dominated by those with a greater intellect or willpower."
	case i.Pow > 15 && i.Pow <= 90:
		i.Description.PowDescription = "strong willed, driven, a high potential to connect with the unseen and magical."
	case i.Pow == 100:
		i.Description.PowDescription = "iron willed, you have an strong connection to the spiritual 'realm' or the unseen world."
	}
}

func (i *Investigator) dexDescription() {
	switch i.Dex >= 0 {
	case i.Dex == 0:
		i.Description.DexDescription = "unable to move without assistance."
	case i.Dex > 0 && i.Dex <= 15:
		i.Description.DexDescription = "slow, clumsy with poor motor skills for fine manipulation."
	case i.Dex > 15 && i.Dex <= 50:
		i.Description.DexDescription = "average human dexterity."
	case i.Dex > 50 && i.Dex <= 90:
		i.Description.DexDescription = "fast, nimble and able to perform feats of fine manipulation."
	case i.Dex > 90 && i.Dex <= 99:
		i.Description.DexDescription = "world class athlete. Human maximum."
	}
}

func (i *Investigator) educationDescription() {
	switch i.Edu >= 0 {
	case i.Edu == 0:
		i.Description.EduDescription = "new born baby."
	case i.Edu >= 1 && i.Edu <= 15:
		i.Description.EduDescription = "completely uneducated in every way."
	case i.Edu > 15 && i.Edu <= 60:
		i.Description.EduDescription = "high school graduate."
	case i.Edu > 60 && i.Edu <= 70:
		i.Description.EduDescription = "college graduate."
	case i.Edu > 70 && i.Edu <= 80:
		i.Description.EduDescription = "degree level graduate."
	case i.Edu > 80 && i.Edu <= 90:
		i.Description.EduDescription = "doctorate, professor."
	case i.Edu > 90 && i.Edu <= 96:
		i.Description.EduDescription = "world class authority in your field of study."
	case i.Edu > 96:
		i.Description.EduDescription = "reached the peak of human education."
	}
}
