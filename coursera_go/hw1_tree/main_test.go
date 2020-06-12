package main

import (
	"bytes"
	"testing"
)

const testFullResult = `├───project
│	├───file.txt (19b)
│	└───gopher.png (70372b)
├───static
│	├───a_lorem
│	│	├───dolor.txt (empty)
│	│	├───gopher.png (70372b)
│	│	└───ipsum
│	│		└───gopher.png (70372b)
│	├───css
│	│	└───body.css (28b)
│	├───empty.txt (empty)
│	├───html
│	│	└───index.html (57b)
│	├───js
│	│	└───site.js (10b)
│	└───z_lorem
│		├───dolor.txt (empty)
│		├───gopher.png (70372b)
│		└───ipsum
│			└───gopher.png (70372b)
├───zline
│	├───empty.txt (empty)
│	└───lorem
│		├───dolor.txt (empty)
│		├───gopher.png (70372b)
│		└───ipsum
│			└───gopher.png (70372b)
└───zzfile.txt (empty)
`

func TestTreeFull(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "D:\\Postgres", true)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}

	result := out.String()

	if result != testFullResult {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testFullResult)
	}
}

const testDirResult = `|---(by-Regina-O.-Obe,-Leo-S.-Hsu)-PostgreSQL-Up-and--3382905-(z-lib.org).pdf
|---2020-05-27-raspios-buster-full-armhf.zip
|---Debian_9.5.0_stud.part01.rar
|---Debian_9.5.0_stud.part02.rar
|---Debian_9.5.0_stud.part03.rar
|---Debian_9.5.0_stud.part04.rar
|---Debian_9.5.0_stud.part05.rar
|---Debian_9.5.0_stud.part06.rar
|---Debian_9.5.0_stud.part07.rar
|---Debian_9.5.0_stud.part08.rar
|---Debian_9.5.0_stud.part09.rar
|---Debian_9.5.0_stud.part10.rar
|---My_Debian
|       |---Logs
|       |       |---VBox.log
|       |       |---VBox.log.1
|       |       |---VBox.log.2
|       |       |---VBox.log.3
|       |       |___VBoxHardening.log
|       |---My_Debian.vbox
|       |---My_Debian.vbox-prev
|       |___My_Debian.vdi
|---Tekhnologii_proektirovania_baz_dannykh_2019_Osipov.pdf
|---sql_primer.pdf
|---sqlprimer.pdf
|___База
        |---Commands.txt
        |___new 1.txt

`

func TestTreeDir(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "D:\\Postgres", false)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testDirResult {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testDirResult)
	}
}