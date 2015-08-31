mstranslator
======================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/mstranslator/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/goFbAlbum?status.svg)](https://godoc.org/github.com/kkdai/mstranslator)  [![Build Status](https://travis-ci.org/kkdai/mstranslator.svg)](https://travis-ci.org/kkdai/mstranslator)
 
The "mstranslator" is a [Micrsoft Translator Service](https://www.microsoft.com/translator/) client which written by Golang. 


What is Microsoft Translator 
---------------
![image](https://pbs.twimg.com/profile_images/599262624996462592/N61dCOMr_normal.png)

Microsoft Translator is a cloud based automatic translation service. (Refer [here](https://www.microsoft.com/translator/) for more detail). 

Here is another site for [Microsoft Translator API](https://www.microsoft.com/translator/api.aspx).

Installation
---------------

        go get github.com/kkdai/mstranslator

How to use it
---------------

Sign-up for Microsoft Translator API (see [here](http://blogs.msdn.com/b/translation/p/gettingstarted1.aspx) for more detail) and get your developer credentials. Use the client ID and secret to instantiate a translator as shown below.

```go
        package main
        
        import (
        	"fmt"
        	"log"
        
                ms "github.com/kkdai/mstranslator"
                )
        
        func main() {
        	msClient := ms.NewClient("YourClientID", "YourClientSecret")
        
        	//Translate "Hello World" from English to France.
        	translation, err := msClient.Translate("Hello World!", "en", "de")
        	if err != nil {
        		log.Panicf("Error : %s", err.Error())
        	}
        	fmt.Println(translation) //Hallo Welt!        
        }
```

Check [example](example/main.go) for more detail.


Implemented APIs
---------------

- [Detect](https://msdn.microsoft.com/en-us/library/ff512411.aspx)
- [DetectArray](https://msdn.microsoft.com/en-us/library/ff512412.aspx)        
- [GetLanguageNames](https://msdn.microsoft.com/en-us/library/ff512414.aspx)
- [GetLanguagesForTranslate](https://msdn.microsoft.com/en-us/library/ff512416.aspx)
- [GetLanguagesForSpeak](https://msdn.microsoft.com/en-us/library/ff512415.aspx)
- [GetTranslations](https://msdn.microsoft.com/en-us/library/ff512417.aspx)
- [Translate](https://msdn.microsoft.com/en-us/library/ff512421.aspx)
- [TransformText](https://msdn.microsoft.com/en-us/library/dn876735.aspx)
- [Speak](https://msdn.microsoft.com/en-us/library/ff512420.aspx)


Unimplement APIs (Yet)
---------------

- [AddTranslation](https://msdn.microsoft.com/en-us/library/ff512408.aspx)
- [AddTranslationArray](https://msdn.microsoft.com/en-us/library/ff512409.aspx)
- [BreakSentences](https://msdn.microsoft.com/en-us/library/ff512410.aspx)
- [GetTranslationsArray](https://msdn.microsoft.com/en-us/library/ff512418.aspx)
- [TranslateArray](https://msdn.microsoft.com/en-us/library/ff512422.aspx)


Contribute
---------------

Please open up an issue on GitHub before you put a lot efforts on pull request.
The code submitting to PR must be filtered with `gofmt`

Inspired
---------------

This project is inspired by [https://github.com/st3v/translator](https://github.com/st3v/translator). 


License
---------------

This package is licensed under MIT license. See LICENSE for details.


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/kkdai/mstranslator/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

