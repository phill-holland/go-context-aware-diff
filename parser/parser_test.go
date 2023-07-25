package parser_test

import (
	"parser"
	"testing"
)

func TestComparisonAddFunctionAndUpdateLine(t *testing.T) {

	program1 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
    int c = 145;
}`

	program2 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hellfo! \" world!";
    int c = 145;
    int g = 444;
}

int func(int a)
{
    int b = a;
    return b + 5;
}`

	expected := `#include <iostream>
using namespace std;
void main()
{
- cout << "hello! \" world!";
+ cout << "hellfo! \" world!";
int c = 145;
+ int g = 444;
}
+ int func(int a)
+ {
+ int b = a;
+ return b + 5;
+ }
`
	p1 := parser.Parse(program1)
	p2 := parser.Parse(program2)

	result := p1.Compare(p2)
	data := result.Print()
	if expected != data {
		t.Errorf("Outputted program does not match expected")
	}
}

func TestComparisonRemoveSingleLine(t *testing.T) {

	program1 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
    int c = 145;
}`

	program2 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
}`

	expected := `#include <iostream>
using namespace std;
void main()
{
cout << "hello! \" world!";
- int c = 145;
}
`

	p1 := parser.Parse(program1)
	p2 := parser.Parse(program2)

	result := p1.Compare(p2)
	data := result.Print()
	if expected != data {
		t.Errorf("Outputted program does not match expected")
	}
}

func TestComparisonRemoveSingleFunction(t *testing.T) {

	program1 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
}

int func(int a)
{
    int b = a;
    return b + 5;
}`

	program2 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
}`

	expected := `#include <iostream>
using namespace std;
void main()
{
cout << "hello! \" world!";
}
- int func(int a)
- {
- int b = a;
- return b + 5;
- }
`

	p1 := parser.Parse(program1)
	p2 := parser.Parse(program2)

	result := p1.Compare(p2)
	data := result.Print()
	if expected != data {
		t.Errorf("Outputted program does not match expected")
	}
}

func TestComparisonInsertSeveralLines(t *testing.T) {

	program1 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
    cout << "monkey monkey";
    int a = 10;
}`

	program2 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
    int b = 56;
    cout << "bananas";
    cout << "monkey monkey";
    int a = 10;
}`

	expected := `#include <iostream>
using namespace std;
void main()
{
cout << "hello! \" world!";
+ int b = 56;
+ cout << "bananas";
cout << "monkey monkey";
int a = 10;
}
`

	p1 := parser.Parse(program1)
	p2 := parser.Parse(program2)

	result := p1.Compare(p2)
	data := result.Print()
	if expected != data {
		t.Errorf("Outputted program does not match expected")
	}
}

func TestComparisonAddForLoop(t *testing.T) {

	program1 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
    int w = 0;
    int c = 145;
}`

	program2 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
    int w = 0;
    for(int a = 10; a < 100; ++a)
    {
        w = w * 2;
        cout << "weeeeee!";
    }
    int c = 145;
}`

	expected := `#include <iostream>
using namespace std;
void main()
{
cout << "hello! \" world!";
int w = 0;
+ for(int a = 10; a < 100; ++a)
+ {
+ w = w * 2;
+ cout << "weeeeee!";
+ }
int c = 145;
}
`

	p1 := parser.Parse(program1)
	p2 := parser.Parse(program2)

	result := p1.Compare(p2)
	data := result.Print()
	if expected != data {
		t.Errorf("Outputted program does not match expected")
	}
}

func TestComparisonRemoveForLoop(t *testing.T) {

	program1 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
    int w = 0;
    for(int a = 10; a < 100; ++a)
    {
        w = w * 2;
        cout << "weeeeee!";
    }
    int c = 145;
}`

	program2 := `#include <iostream>

using namespace std;

void main()
{
    cout << "hello! \" world!";
    int w = 0;
    int c = 145;
}`

	expected := `#include <iostream>
using namespace std;
void main()
{
cout << "hello! \" world!";
int w = 0;
- for(int a = 10; a < 100; ++a)
- {
- w = w * 2;
- cout << "weeeeee!";
- }
int c = 145;
}
`

	p1 := parser.Parse(program1)
	p2 := parser.Parse(program2)

	result := p1.Compare(p2)
	data := result.Print()
	if expected != data {
		t.Errorf("Outputted program does not match expected")
	}
}
