package jsontest

type SmallStruct struct {
	Event string `json:"event"`
	Room  struct {
		Sid  string `json:"sid"`
		Name string `json:"name"`
	} `json:"room"`
	Participant struct {
		Sid      string `json:"sid"`
		Identity string `json:"identity"`
	} `json:"participant"`
	Track struct {
		Sid      string `json:"sid"`
		Source   int    `json:"source"`
		MimeType string `json:"mime_type"`
		Mid      string `json:"mid"`
		Stream   string `json:"stream"`
		Version  struct {
			UnixMicro int64 `json:"unix_micro"`
		} `json:"version"`
	} `json:"track"`
	Id        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
}

type MediumStruct struct {
	Person  Person      `json:"persion"`
	Company interface{} `json:"company"`
}

type BigStruct struct {
	Users  []User `json:"users"`
	Topics Topic  `json:"topics"`
	Person Person `json:"persion"`
	Record Record `json:"record"`
}

type Person struct {
	ID   string `json:"id"`
	Name struct {
		FullName   string `json:"fullName"`
		GivenName  string `json:"givenName"`
		FamilyName string `json:"familyName"`
	} `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Location string `json:"location"`
	Geo      struct {
		City    string  `json:"city"`
		State   string  `json:"state"`
		Country string  `json:"country"`
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
	} `json:"geo"`
	Bio        string `json:"bio"`
	Site       string `json:"site"`
	Avatar     string `json:"avatar"`
	Employment struct {
		Name string `json:"name"`
	} `json:"employment"`
	Facebook struct {
		Handle string `json:"handle"`
	} `json:"facebook"`
	Github struct {
		Handle    string `json:"handle"`
		ID        int    `json:"id"`
		Avatar    string `json:"avatar"`
		Company   string `json:"company"`
		Blog      string `json:"blog"`
		Followers int    `json:"followers"`
		Following int    `json:"following"`
	} `json:"github"`
	Twitter struct {
		Handle    string `json:"handle"`
		ID        int    `json:"id"`
		Bio       string `json:"bio"`
		Followers int    `json:"followers"`
		Site      string `json:"site"`
		Avatar    string `json:"avatar"`
	} `json:"twitter"`
	Linkedin struct {
		Handle string `json:"handle"`
	} `json:"linkedin"`
	Googleplus struct {
		Handle string `json:"handle"`
	} `json:"googleplus"`
	Angellist struct {
		Handle    string `json:"handle"`
		ID        int    `json:"id"`
		Bio       string `json:"bio"`
		Followers int    `json:"followers"`
		Avatar    string `json:"avatar"`
	} `json:"angellist"`
	Klout struct {
		Handle string  `json:"handle"`
		Score  float64 `json:"score"`
	} `json:"klout"`
	Foursquare struct {
		Handle string `json:"handle"`
	} `json:"foursquare"`
	Aboutme struct {
		Handle string `json:"handle"`
		Bio    string `json:"bio"`
		Avatar string `json:"avatar"`
	} `json:"aboutme"`
	Gravatar struct {
		Handle  string   `json:"handle"`
		Urls    []string `json:"urls"`
		Avatar  string   `json:"avatar"`
		Avatars []struct {
			URL  string `json:"url"`
			Type string `json:"type"`
		} `json:"avatars"`
	} `json:"gravatar"`
	Fuzzy bool `json:"fuzzy"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_template"`
}

type Topic struct {
	CanCreateTopic bool        `json:"can_create_topic"`
	MoreTopicsURL  string      `json:"more_topics_url"`
	DraftKey       string      `json:"draft_key"`
	DraftSequence  interface{} `json:"draft_sequence"`
	PerPage        int         `json:"per_page"`
	Topics         []struct {
		ID                 int         `json:"id"`
		Title              string      `json:"title"`
		FancyTitle         string      `json:"fancy_title"`
		Slug               string      `json:"slug"`
		PostsCount         int         `json:"posts_count"`
		ReplyCount         int         `json:"reply_count"`
		HighestPostNumber  int         `json:"highest_post_number"`
		ImageURL           string      `json:"image_url"`
		CreatedAt          string      `json:"created_at"`
		LastPostedAt       string      `json:"last_posted_at"`
		Bumped             bool        `json:"bumped"`
		BumpedAt           string      `json:"bumped_at"`
		Unseen             bool        `json:"unseen"`
		Pinned             bool        `json:"pinned"`
		Unpinned           interface{} `json:"unpinned"`
		Excerpt            string      `json:"excerpt"`
		Visible            bool        `json:"visible"`
		Closed             bool        `json:"closed"`
		Archived           bool        `json:"archived"`
		Bookmarked         interface{} `json:"bookmarked"`
		Liked              interface{} `json:"liked"`
		Views              int         `json:"views"`
		LikeCount          int         `json:"like_count"`
		HasSummary         bool        `json:"has_summary"`
		Archetype          string      `json:"archetype"`
		LastPosterUsername string      `json:"last_poster_username"`
		CategoryID         int         `json:"category_id"`
		PinnedGlobally     bool        `json:"pinned_globally"`
		Posters            []struct {
			Extras      string `json:"extras"`
			Description string `json:"description"`
			UserID      int    `json:"user_id"`
		} `json:"posters"`
	} `json:"topics"`
}
type Request struct {
	Track struct {
		RoomName string `json:"room_name"`
		TrackID  string `json:"track_id"`
		Output   struct {
			File struct {
				FilePath string      `json:"filepath"`
				Output   interface{} `json:"Output"`
			} `json:"File"`
		} `json:"Output"`
	} `json:"Track"`
}
type FileResult struct {
	Filename  string `json:"filename"`
	StartedAt int64  `json:"started_at"`
	Duration  int64  `json:"duration"`
	Size      int64  `json:"size"`
	Location  string `json:"location"`
}
type Record struct {
	Event      string `json:"event"`
	EgressInfo struct {
		EgressID  string `json:"egress_id"`
		RoomID    string `json:"room_id"`
		RoomName  string `json:"room_name"`
		Status    int    `json:"status"`
		StartedAt int64  `json:"started_at"`
		EndedAt   int64  `json:"ended_at"`
		UpdatedAt int64  `json:"updated_at"`
		Request   struct {
			Track struct {
				RoomName string `json:"room_name"`
				TrackID  string `json:"track_id"`
				Output   struct {
					File struct {
						FilePath string      `json:"filepath"`
						Output   interface{} `json:"Output"`
					} `json:"File"`
				} `json:"Output"`
			} `json:"Track"`
		} `json:"Request"`
		Result struct {
			File FileResult `json:"File"`
		} `json:"Result"`
		FileResults []FileResult `json:"file_results"`
	} `json:"egress_info"`
	ID        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
}
