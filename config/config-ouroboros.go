package config

type OuroborosConfig struct {
	FeatureDetectorName string `json:"FeatureDetectorName"` //The name of the feature detection algorithm to use. Only a placeholder for now.

	SIFTConfig SIFTConfig `json:"SIFTConfig"` //SIFT feature detector specific settings.
}

type SIFTConfig struct {
	//keypoint related parameters
	SiftWorkingSize   int     `json:"Sift.Working.Resolution"` //Working resolution for sift
	NumOctave         int     `json:"Octave.Num"`              //Number of octaves to construct:
	NumScale          int     `json:"Scale.Num"`               //Number of successively blurred images to construct in each octave
	ScaleFactor       float64 `json:"Scale.Factor"`            //K, the multiplicative scaling factor given in Lowe 2004
	GaussSigma        float64 `json:Gauss.Sigma"`              //The value of sigma used in the gaussian blur kernel
	GaussWindowFactor int     `json:Gauss.Window.Factor"`      //From Lowe 2004. larger value gives less feature
}
