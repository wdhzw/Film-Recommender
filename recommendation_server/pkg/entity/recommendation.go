package entity

type Recommendation struct {
    UserID string
    Movies []Movie
}

type Movie struct {
    ID       string
    Title    string
    // other movie properties
}
