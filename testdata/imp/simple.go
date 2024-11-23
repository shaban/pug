:go:func Simple( u *model.User, st []model.Story )

:go:import "github.com/Shaban/pug/testdata/imp/model"

html
    body
        h1= u.FirstName
        p Here's a list of your favorite colors:
        ul
            each colorName in u.FavoriteColors
                li= colorName

        ul
            each story in st
                li= story.StoryId
                li= story.UserId
                li= story.UserName
