package javbus

// 影片类型
type MovieType string

const (
	Normal     MovieType = "normal"
	Uncensored MovieType = "uncensored"
)

// 磁力链接类型

type MagnetType string

const (
	All   MagnetType = "all"
	Exist MagnetType = "exist"
)

// 过滤器类型
type FilterType string

const (
	Star     FilterType = "star"
	Genre    FilterType = "genre"
	Director FilterType = "director"
	Studio   FilterType = "studio"
	Label    FilterType = "label"
	Series   FilterType = "series"
)

// 排序规则
type SortBy string

const (
	Date SortBy = "date"
	Size SortBy = "size"
)

// 排序顺序逆序
type SortOrder string

const (
	Asc  SortOrder = "asc"
	Desc SortOrder = "desc"
)

// 属性
type Property struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 影片
type Movie struct {
	Date  string   `json:"date"`
	Title string   `json:"title"`
	ID    string   `json:"id"`
	Img   string   `json:"img"`
	Tags  []string `json:"tags"`
}

// 磁力链接
type Magnet struct {
	ID          string `json:"id"`
	Link        string `json:"link"`
	IsHD        bool   `json:"isHD"`
	Title       string `json:"title"`
	Size        string `json:"size"`
	NumberSize  int    `json:"numberSize"`
	ShareDate   string `json:"shareDate"`
	HasSubtitle bool   `json:"hasSubtitle"`
}

// 图片大小
type ImageSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// 样片图
type Sample struct {
	Alt       string `json:"alt"`
	ID        string `json:"id"`
	Thumbnail string `json:"thumbnail"`
	Src       string `json:"src"`
}

// 影片详情
type MovieDetail struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Img         string     `json:"img"`
	Date        string     `json:"date"`
	VideoLength int        `json:"videoLength"`
	Director    *Property  `json:"director"`
	Producer    *Property  `json:"producer"`
	Publisher   *Property  `json:"publisher"`
	Series      *Property  `json:"series"`
	Genres      []Property `json:"genres"`
	Stars       []Property `json:"stars"`
	ImageSize   *ImageSize `json:"imageSize"`
	Samples     []Sample   `json:"samples"`
	GID         string     `json:"gid"`
	UC          string     `json:"uc"`
}

// 分页
type Pagination struct {
	CurrentPage int   `json:"currentPage"`
	HasNextPage bool  `json:"hasNextPage"`
	NextPage    *int  `json:"nextPage"`
	Pages       []int `json:"pages"`
}

// 影片页
type MoviesPage struct {
	Movies     []Movie    `json:"movies"`
	Pagination Pagination `json:"pagination"`
}

// 演员信息
type StarInfo struct {
	Avatar     string `json:"avatar"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Birthday   string `json:"birthday"`
	Age        int    `json:"age"`
	Height     int    `json:"height"`
	Bust       string `json:"bust"`
	Waistline  string `json:"waistline"`
	Hipline    string `json:"hipline"`
	Birthplace string `json:"birthplace"`
	Hobby      string `json:"hobby"`
}

// 搜索页
type SearchMoviesPage struct {
	Movies     []Movie    `json:"movies"`
	Pagination Pagination `json:"pagination"`
	Keyword    string     `json:"keyword"`
}

// 个人信息Map
var starInfoMap = map[string]string{
	"birthday":   "生日: ",
	"age":        "年齡: ",
	"height":     "身高: ",
	"bust":       "胸圍: ",
	"waistline":  "腰圍: ",
	"hipline":    "臀圍: ",
	"birthplace": "出生地: ",
	"hobby":      "愛好: ",
}
