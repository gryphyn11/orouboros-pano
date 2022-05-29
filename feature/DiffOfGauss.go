package feature

import (
	"math"
	"main/config"
)


type GaussianPyramid[T MatrixType]struct {
	Nscale int
	Data   []*Matrix[T] // len = nscale
	Mag    []*Matrix[T] // len = nscale
	Ort    []*Matrix[T]
	// len = nscale, value \in [0, 2 * pi]
	//void cal_mag_ort(int);
	W int
	H int
}

func newGaussianPyramid(mat *Matrix[T], nscale int) *GaussianPyramid {
	pyr := GaussianPyramid[T]{Nscale: nscale, W: m.width(), H: m.height()}
	data = make(*Matrix[T], num_scale)
	mag = make(*Matrix[T], num_scale)
	ort = make(*Matrix[T], num_scale)
	
	return &pyr
}

type GaussianPyramidBuilder[T MatrixType]struct {
	Nscale int
	ImgMat *Matrix[T]
	config SIFTConfig

}

func newGaussPyramidBuilder(mat *Matrix[T], nscale int) *GaussianPyramidBuilder {
	pyr := GaussianPyramidBuilder[T]{Nscale: nscale, W: m.width(), H: m.height()}
	data = make(*Matrix[T], num_scale)
	mag = make(*Matrix[T], num_scale)
	ort = make(*Matrix[T], num_scale)
	
	return &pyr
}


func (pyr *GaussianPyramidBuilder[T]) buildPyramids(){

	if (m.channels() == 3)
		data[0] = rgb2grey();
	else
		data[0] = m.clone();

	blurer := makeMultiStageGaussianBlurrer(nscale, SIFTConfig , SCALE_FACTOR);
	for (int i = 1; i < nscale; i ++) {
		data[i] = blurer.blur(data[0], i);	// sigma needs a better one
		cal_mag_ort(i);
	}



func (pyr *GaussianPyramidBuilder) cal_mag_ort(int i) {
	const Mat32f& orig = data[i];
	int w = orig.width(), h = orig.height();
	mag[i] = Mat32f(h, w, 1);
	ort[i] = Mat32f(h, w, 1);
	REP(y, h) {
		float *mag_row = mag[i].ptr(y),
					*ort_row = ort[i].ptr(y);
		const float *orig_row = orig.ptr(y),
					*orig_plus = orig.ptr(y + 1),
					*orig_minus = orig.ptr(y - 1);
		// x == 0:
		mag_row[0] = 0;
		ort_row[0] = M_PI;

		REPL(x, 1, w-1) {
			if (between(y, 1, h - 1)) {
				float dy = orig_plus[x] - orig_minus[x],
							dx = orig_row[x + 1] - orig_row[x - 1];
				mag_row[x] = hypotf(dx, dy);
				// approx here cause break working on myself/small*. fix later
				// when dx==dy==0, no need to set ort
				ort_row[x] = fast_atan(dy, dx) + M_PI;
			} else {
				mag_row[x] = 0;
				ort_row[x] = M_PI;
			}
		}
		// x == w-1
		mag_row[w-1] = 0;
		ort_row[w-1] = M_PI;

	}
}
/*
	GaussianPyramid( Mat32f&, int num_scale);

	inline const Mat32f& get(int i) const { return data[i]; }

	inline const Mat32f& get_mag(int i) const { return mag[i]; }

	inline const Mat32f& get_ort(int i) const { return ort[i]; }

	int get_len() const { return nscale; }
*/

/*
test_extrema(const char* fname, int mode) {
	auto mat = read_img(fname);

	ScaleSpace ss(mat, NUM_OCTAVE, NUM_SCALE);
	DOGSpace dog(ss);
	ExtremaDetector ex(dog);

	PlaneDrawer pld(mat);
	if (mode == 0) {
		auto extrema = ex.get_raw_extrema();
		PP(extrema.size());
		for (auto &i : extrema)
			pld.cross(i, LABEL_LEN / 2);
	} else if (mode == 1) {
		auto extrema = ex.get_extrema();
		cout << extrema.size() << endl;
		for (auto &i : extrema) {
			Coor c{(int)(i.real_coor.x * mat.width()), (int)(i.real_coor.y * mat.height())};
			pld.cross(c, LABEL_LEN / 2);
		}
	}
	write_rgb(IMGFILE(extrema), mat);
}

*/





type scaleSpace struct {
	noctave int
	nscale  int
	origw   int
	origh   int

	pyramids []*GaussianPyramidBuilder // len = noctave
}

func newScaleSpace(mat *Matrix[float32], num_octave int, num_scale int) *scaleSpace {
	ss := scaleSpace{noctave: num_octave, nscale: num_scale, origw: mat.width(), origh: mat.height()}
	ss.pyramids = make([]*GaussianPyramidBuilder, num_octave)

	ss.pyramids[:0] = newMatrix[float32](mat, num_scale)
	for i := 0; i < num_octave; i++ {
		factor := math.Pow(SCALE_FACTOR, -i)
		neww := math.Ceil(float64(ss.origw) * factor)
		newh := math.Ceil(float64(ss.origh) * factor)

		//assert(neww > 5 && newh > 5);
		resized := newMatrix[float32](newh, neww, 3)
		resize(mat, resized)
		pyramids.emplace_back(resized, nscale)
	}

	return &ss
}








class DOGSpace {

	public:
		// Calculate difference of a list of image
		// diff[0] = orig[1] - orig[0]

		typedef std::vector<Mat32f> DOG;	// len = nscale - 1

		int noctave, nscale;
		int origw, origh;

		std::vector<DOG> dogs;		// len = noctave

		DOGSpace(const DOGSpace&) = delete;
		DOGSpace& operator = (const DOGSpace&) = delete;

		Mat32f diff(const Mat32f& img1, const Mat32f& img2) const;
		DOGSpace(ScaleSpace&);

};