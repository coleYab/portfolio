package data

import "portfolio/types"

var Projects = []types.Project{
	{
		ImageURL:    "https://cdn.freelogovectors.net/wp-content/uploads/2023/02/react-logo-freelogovectors.net_.png",
		GithubURL:   "https://github.com/coleYab/Enawga",
		Title:       "Enawga",
		Description: "A simple a web based chat app",
		TechStacks:  []string{"JavaScript", "React", "JSX", "Express"},
		Stars:       1,
	},
	{
		ImageURL:    "https://i1.wp.com/www.mwendengao.com/wp-content/uploads/2018/04/Safaricom-Mpesa.png",
		GithubURL:   "https://github.com/coleYab/mpesasdk",
		Title:       "MpesaSDK",
		Description: "A simple mpesa sdk build with golang",
		TechStacks:  []string{"Go"},
		Stars:       1,
	},
	{
		ImageURL:    "https://w7.pngwing.com/pngs/566/160/png-transparent-golang-hd-logo-thumbnail.png",
		GithubURL:   "https://github.com/coleYab/podex",
		Title:       "Podex",
		Description: "A simle cli repl build while learning golang",
		TechStacks:  []string{"Go"},
		Stars:       1,
	},
	{
		ImageURL:    "https://1000logos.net/wp-content/uploads/2020/08/Python-Logo.png",
		GithubURL:   "https://github.com/coleYab/amharic_se",
		Title:       "Amharic SE",
		Description: "A simple amharic search engine build as ISR assignment",
		TechStacks:  []string{"Python", "Flask", "HTML"},
		Stars:       3,
	},
	{
		ImageURL:    "https://w7.pngwing.com/pngs/724/306/png-transparent-c-logo-c-programming-language-icon-letter-c-blue-logo-computer-program-thumbnail.png",
		GithubURL:   "https://github.com/coleYab/Maze",
		Title:       "Maze",
		Description: "A simple raycasting game build with C",
		TechStacks:  []string{"C", "SDL2", "C++"},
		Stars:       1,
	},
	{
		ImageURL:    "https://upload.wikimedia.org/wikipedia/commons/6/6a/JavaScript-logo.png",
		GithubURL:   "https://github.com/coleYab/alx-file_manager",
		Title:       "File Manager",
		Description: "A simple file manager build with express js",
		TechStacks:  []string{"javascript", "express", "redis", "mongodb"},
		Stars:       1,
	},
}
