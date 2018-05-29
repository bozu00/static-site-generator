package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "log"

  "html/template"
  yaml "gopkg.in/yaml.v2"
  "path/filepath"
  "path"
)

// structたち


type Article struct {
  Meta MetaInfo
  Title string 
  Date string
  Contents []string
  IsPusblish bool
}



type MetaInfo struct {
  Id int
  Path  string 
}

func main() {
  temp := getYamlFilePaths("./src")
  fmt.Println(temp)
  for _, ele := range temp {
	  processArticleYaml(ele)
  }

  fmt.Printf("finish")
}



func parseArticleYaml(ymlName string) Article {
  buf, err := ioutil.ReadFile("./article.yaml")
  if err != nil {
    panic(err)
  }

  // structにUnmasrshal
  var a Article 
  err = yaml.Unmarshal(buf, &a)
  if err != nil {
    panic(err)
  }

  for i, e := range a.Contents {
	  a.Contents[i] = e + "-test"
  }

  return a
}

func processArticleYaml(articleYamlPath string) {
	dirName := path.Dir(articleYamlPath)
	contentFile := path.Join(dirName, "content.txt")
	fmt.Println(contentFile)

	prefix := "src/"
	destPath := path.Join("public", dirName[len(prefix):])
	fmt.Println(destPath)
	

	articleYaml := parseArticleYaml(articleYamlPath)
	createDestHtml(destPath, articleYaml)
	// dest := "public"

}


func createDestHtml(filePath string, article Article) {
    // テンプレートをパース
	// todo yamlからテンプレート情報を抜き出して使う
    t := template.Must(template.ParseFiles("./article_template.html"))


	// file, err := os.Create(`./article_out.html`)
	dirName := path.Dir(filePath)
	if err := os.MkdirAll(dirName, 0777); err != nil {
		panic(err)
	}
	file, err := os.Create(filePath + ".html")
    if err != nil {
        // Openエラー処理
		log.Fatal(err)
    }
    defer file.Close()
    // テンプレートを描画
    // if err := t.ExecuteTemplate(w, "template000.html", article); err != nil {
    //     log.Fatal(err)
    // }
	log.Println(article)
	if err := t.Execute(file ,article); err != nil {
		log.Fatal(err)
	}
}

func getYamlFilePaths(rootDir string) []string{
	var paths []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
	if info.Name() == "article.yaml" {
		paths = append(paths, path)
	}

		return nil
	})

	if err != nil {
		fmt.Println(1, err)
	}
	return paths
}
