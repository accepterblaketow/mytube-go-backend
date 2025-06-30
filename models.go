
package main

type Video struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Tags        string `json:"tags"`
    URL         string `json:"url"`
}
