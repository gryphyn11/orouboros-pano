package feature

type MatrixType interface {
	float32 | float64 | uint8
}

type Matrix[T MatrixType] struct {
	m_rows     int
	m_cols     int
	m_channels int
	m_data     []T

	/*
		const T *ptr(int r = 0) const
		{ return m_data.get() + r * m_cols * m_channels; }
		T *ptr(int r = 0)
		{ return m_data.get() + r * m_cols * m_channels; }
		const T *ptr(int r, int c) const
		{ return m_data.get() + (r * m_cols + c) * m_channels; }
		T *ptr(int r, int c)
		{ return m_data.get() + (r * m_cols + c) * m_channels; }
	*/

}

func (m *Matrix[T]) clone() *Matrix[T] {
	mat := Matrix[T]{m_rows: m.m_rows, m_cols: m.m_cols, m_channels: m.m_channels}
	src := m.m_data[:]
	dst := make([]T, len(m.m_data))
	//dst := make([]T , m_rows * m_cols * m_channels);
	copy(dst, src)
	mat.m_data = dst

	return &mat
}

func newMatrix[T MatrixType](rows int, cols int, channels int) *Matrix[T] {
	mat := Matrix[T]{m_rows: rows, m_cols: cols, m_channels: channels}
	//m_rows = rows
	//m_cols = cols
	//m_channels = channels
	mat.m_data = make([]T, rows*cols*channels)
	return &mat
}

func (m *Matrix[T]) height() int   { return m.m_rows }
func (m *Matrix[T]) width() int    { return m.m_cols }
func (m *Matrix[T]) rows() int     { return m.m_rows }
func (m *Matrix[T]) cols() int     { return m.m_cols }
func (m *Matrix[T]) channels() int { return m.m_channels }
func (m *Matrix[T]) pixels() int   { return m.m_rows * m.m_cols }

func (m *Matrix[T]) at(r int, c int, ch int) T {
	//m_assert(r < m.m_rows);
	//m_assert(c < m.m_cols);
	//m_assert(ch < m.m_channels);
	return m.m_data[(r*m.m_cols+c)*m.m_channels+ch]
}

type Mat32 Matrix[float32]
type Mat64 Matrix[float64]
