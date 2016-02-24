# jisho

Uses the [kanjium](https://github.com/mifunetoshiro/kanjium) database ([license](https://github.com/mifunetoshiro/kanjium/blob/master/LICENSE.txt)).

Not affiliated with the amazing [jisho.org](http://jisho.org/).

Source code available on [Gitlab](https://gitlab.com/SiegfriedEhret/jisho.go) or [Github](https://github.com/SiegfriedEhret/jisho.go).

## install

```
go install github.com/mattn/go-sqlite3
go get -u gitlab.com/SiegfriedEhret/jisho.go
```

## usage

```
./jisho.go 冥
```

You can also search multiple kanjis !

```
./jisho.go 冥 彩 子
```

## license

Under the MIT license:

```
Copyright (c) 2015-2016 Siegfried Ehret <siegfried@ehret.me>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

```
