// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package expressions

type T struct {
	x, y, z int
}

var (
	a, b, c, d, e int;
	under_bar int;
	longIdentifier1, longIdentifier2, longIdentifier3 int;
	t0, t1, t2 T;
	s string;
	p *int;
)


func _() {
	// no spaces around simple or parenthesized expressions
	_ = a+b;
	_ = a+b+c;
	_ = a+b-c;
	_ = a-b-c;
	_ = a+(b*c);
	_ = a+(b/c);
	_ = a-(b%c);
	_ = 1+a;
	_ = a+1;
	_ = a+b+1;
	_ = s[1:2];
	_ = s[a:b];
	_ = s[0:len(s)];
	_ = s[0]<<1;
	_ = (s[0]<<1)&0xf;
	_ = s[0] << 2 | s[1] >> 4;
	_ = "foo"+s;
	_ = s+"foo";
	_ = 'a'+'b';

	// spaces around expressions of different precedence or expressions containing spaces
	_ = a + -b;
	_ = a - ^b;
	_ = a / *p;
	_ = a + b*c;
	_ = 1 + b*c;
	_ = a + 2*c;
	_ = a + c*2;
	_ = 1 + 2*3;
	_ = s[1 : 2*3];
	_ = s[a : b-c];
	_ = s[a+b : len(s)];
	_ = s[len(s) : -a];
	_ = s[a : len(s)+1];
	_ = s[a : len(s)+1]+s;

	// spaces around operators with equal or lower precedence than comparisons
	_ = a == b;
	_ = a != b;
	_ = a > b;
	_ = a >= b;
	_ = a < b;
	_ = a <= b;
	_ = a < b && c > d;
	_ = a < b || c > d;

	// spaces around "long" operands
	_ = a + longIdentifier1;
	_ = longIdentifier1 + a;
	_ = longIdentifier1 + longIdentifier2 * longIdentifier3;
	_ = s + "a longer string";

	// some selected cases
	_ = a + t0.x;
	_ = a + t0.x + t1.x * t2.x;
	_ = a + b + c + d + e + 2*3;
	_ = a + b + c + 2*3 + d + e;
	_ = (a+b+c)*2;
	_ = a - b + c - d + (a+b+c) + d&e;
	_ = under_bar-1;
}


func _() {
	_ = T{};
	_ = struct{}{};
	_ = [10]T{};
	_ = [...]T{};
	_ = []T{};
	_ = map[int]T{};

	_ = (T){};
	_ = (struct{}){};
	_ = ([10]T){};
	_ = ([...]T){};
	_ = ([]T){};
	_ = (map[int]T){};
}


func _() {
	// TODO respect source line breaks in multi-line expressions
	_ = a < b ||
		b < a;
	// TODO(gri): add more test cases
	// TODO(gri): these comments should be indented
}
