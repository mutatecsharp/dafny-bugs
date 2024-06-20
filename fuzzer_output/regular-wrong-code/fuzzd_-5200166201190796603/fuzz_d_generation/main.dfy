function abs(x: int): int {
	if (x < 0) then -1 * x else x
}
function safeIndex(x: int, length: int): int 
	requires length > 0
{
	if (x < 0) then 0 else if (x >= length) then x % length else x
}
function safeDivisionInt(x1: int, x2: int): int {
	if (x2 == 0) then x1 else x1 / x2
}
function safeModuloInt(x1: int, x2: int): int {
	if (x2 == 0) then x1 else x1 % x2
}
datatype D0 = DC1(cf1: int, cf2: bool, cf3: int) | DC2 | DC3 | DC0(cf0: int)
datatype D1 = DC4(cf4: string)
datatype D2 = DC6(cf6: map<bool, bool>, cf7: bool) | DC7(cf8: int, cf9: array<T0>) | DC5(cf5: multiset<int>)
datatype D3 = DC8(cf10: map<T0, array<int>>)
datatype D4 = DC10 | DC11(cf12: int) | DC9(cf11: seq<T0>) | DC12(cf13: D4)
datatype D5 = DC13(cf14: map<D0, int>)
datatype D6 = DC15(cf16: int, cf17: int) | DC14(cf15: array<bool>) | DC16(cf18: D6)
datatype D7 = DC18(cf20: int) | DC17(cf19: array<int>) | DC19(cf21: D7)
datatype D8 = DC21(cf23: int, cf24: C0) | DC20(cf22: char)
datatype D9 = DC23(cf26: int) | DC24(cf27: int, cf28: C0, cf29: char, cf30: int) | DC25 | DC22(cf25: set<bool>)
datatype D10 = DC27(cf32: bool, cf33: bool) | DC28(cf34: int, cf35: bool, cf36: int, cf37: T0, cf38: int) | DC26(cf31: map<map<bool, bool>, C0>)
datatype D11 = DC30(cf40: bool) | DC29(cf39: set<char>)
trait T0 {
}

class GlobalState {
	const f0 : bool
	const f1 : bool
	var f2 : bool
	var f3 : int
	var f4 : int
	const f5 : int
	const f6 : array<int>
	const f7 : bool
	const f8 : array<int>
	const f9 : array<bool>
	const f10 : int
	const f11 : bool
	const f12 : string
	const f13 : int
	const f14 : bool
	constructor (f0 : bool, f1 : bool, f2 : bool, f3 : int, f4 : int, f5 : int, f6 : array<int>, f7 : bool, f8 : array<int>, f9 : array<bool>, f10 : int, f11 : bool, f12 : string, f13 : int, f14 : bool) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
		this.f7 := f7;
		this.f8 := f8;
		this.f9 := f9;
		this.f10 := f10;
		this.f11 := f11;
		this.f12 := f12;
		this.f13 := f13;
		this.f14 := f14;
	}
	
}

class C0 extends T0 {
	const f15 : map<D0, int>
	constructor (f15 : map<D0, int>) {
		this.f15 := f15;
	}
	
}

function fm0(p0: int, p1: int, globalState: GlobalState): int {
	--(if (true) then if (!!!false) then 0x245 else |map[-0x3 := |[|"xhl"|]|]| else |(multiset{0x3d5, -0x332} - multiset{|(map v0 : int | v0 in {--229, |[true]|} :: (safeModuloInt(v0, -0x2b5)) := (|map[-902 := 'd']|))|, |multiset{-|(set v1 : int | (430 <= v1) && (v1 < 54) :: (safeModuloInt(v1, 0x2e2)))|}|})|)
}
function fm1(globalState: GlobalState): string {
	((seq(abs(669), i0  => ('j'))) + (seq(abs(0x2a9), i1  => ('b')))) + ("tjfujd" + "cnhbwhio")
}
function fm2(p0: bool, globalState: GlobalState): D0 {
	DC1(977, false, -0x3a2)
}
function fm3(p0: bool, p1: multiset<int>, p2: bool, globalState: GlobalState): multiset<bool> {
	multiset{true, !false} + (multiset{false, true, false} - multiset{true})
}
function fm4(p0: bool, p1: bool, globalState: GlobalState): multiset<int> {
	multiset([0x3b9, |map[!false := |[true, true]|]|, -0x303] + [-|"b"|])
}
function fm5(p0: map<bool, int>, p1: bool, p2: int, p3: int, globalState: GlobalState): seq<string> {
	["txsljff"] + (seq(abs(0x12e), i0  => ("cuu")))
}
function fm6(p0: bool, p1: int, globalState: GlobalState): bool {
	match DC1(369, false, 448) {
		case DC1(cf1, cf2, cf3) => cf3 < cf1
		case DC2() => false <==> true
		case DC3() => !(true !in map[!false := !false])
		case DC0(cf0) => -cf0 > 917
	}
}
function fm7(p0: char, p1: int, p2: int, p3: bool, globalState: GlobalState): set<char> {
	({'h'} + {'q'}) + DC29({'m', 'u'}).cf39
}
function fm8(p0: int, globalState: GlobalState): seq<int> {
	(seq(abs(636), i0  => (0x1b5))) + [-774]
}
function fm9(p0: int, p1: int, globalState: GlobalState): set<int> {
	{|multiset{!false}|, 0x300, 0x35f} - {|(map v0 : char | v0 in ['d'] :: (v0) := (DC23(|"jtgcdb"|)))|}
}
function fm10(globalState: GlobalState): seq<bool> {
	[false] + ([true, !false, true] + [true, !true])
}
method m0(p0: bool, globalState: GlobalState) {
	var v0: array<int> := new int[28](i0 => safeDivisionInt(i0, -0x34));
	var v1 := 0x1d9;
	v0[safeIndex(723, v0.Length)] := fm0(v1, v1, globalState);
	var v2: seq<bool> := [p0];
	var v3: seq<int> := [v1, |v2|];
	var v4: seq<seq<int>> := [v3];
	v0[safeIndex(723, v0.Length)] := -safeDivisionInt(v1, v1) + |v4|;
	var i1 := 0;
	while (v2[safeIndex(v1 - v1, |v2|)])
		decreases 100 - i1
	{
		if (i1 >= 100) {
			break;
		}
		
		i1 := i1 + 1;
		var v5: array<string> := new string[26](i2 => "kvewhxik");
		var v6 := "mqwfqaq";
		v5[safeIndex(542, v5.Length)] := v6 + v6;
		var v7: array<map<D0, multiset<bool>>> := new map<D0, multiset<bool>>[2];
		var v8 := DC0(v1);
		var v9: multiset<bool> := multiset{p0, p0, p0, true, p0};
		v7[safeIndex(907, v7.Length)] := map[v8 := v9];
		var v10: multiset<int> := multiset{0x25f};
		var v11: map<D0, multiset<bool>> := map[v8 := fm3(p0, v10, p0, globalState)];
		v5[safeIndex(542, v5.Length)], globalState.f4, v7[safeIndex(907, v7.Length)], globalState.f2 := v6, 106 + safeDivisionInt(v0[safeIndex(723, v0.Length)], v0[safeIndex(723, v0.Length)]), v11 + v11, !(v9 > v9);
		v0[safeIndex(723, v0.Length)] := v1;
		var v13: array<bool> := new bool[23](i3 => (set v12 : int | (0x29f <= v12) && (v12 < -0x83) :: (v12 * DC1(|v10|, p0, v0[safeIndex(723, v0.Length)]).cf3)) !! {v1});
		v13[safeIndex(484, v13.Length)] := true;
		v13[safeIndex(484, v13.Length)] := !fm6((seq(abs(698), i4  => ('k'))) != v5[safeIndex(542, v5.Length)], v0[safeIndex(723, v0.Length)], globalState);
		var v14 := DC1(v1, p0, v0[safeIndex(723, v0.Length)]);
		var v15: map<D0, int> := map[v14 := v0[safeIndex(723, v0.Length)]];
		var v16 := new C0(v15);
	}
	for i5 := v0[safeIndex(723, v0.Length)] to safeDivisionInt(v0[safeIndex(723, v0.Length)], v0[safeIndex(723, v0.Length)]) {
		var v17: set<bool> := {p0, p0};
		var v18 := 'y';
		var v19: map<set<bool>, char> := map[v17 := v18];
		var v20 := DC22(v17);
		v19 := v19[v20.cf25 := v18];
		var v21 := DC1(i5, p0, v1);
		var v22: map<D0, int> := map[v21 := -29];
		var v23 := new C0(v22);
		globalState.f3 := i5;
		var v24 := "lfxeijna";
		var v25: multiset<string> := multiset{v24, v24, "mv", v24};
		var v26: multiset<int> := multiset{i5, v3[safeIndex(|v25|, |v3|)]};
		var v27: map<bool, int> := map[v1 == (if (v1 in v26) then v26[v1] else v1) := safeModuloInt(i5, v1)];
		v27 := v27[p0 := v1];
	}
	if (p0) {
		var v29: map<int, bool> := map[v1 := p0];
		var v30 := DC1(v0[safeIndex(723, v0.Length)], if (v0[safeIndex(723, v0.Length)] in v29) then v29[v0[safeIndex(723, v0.Length)]] else p0, v1);
		var v31: map<D0, int> := map[v30 := v0[safeIndex(723, v0.Length)]];
		var v32: C0 := new C0(map v28 : D0 | v28 in v31 :: (v28) := (v1));
		var v33 := 'l';
		var v34 := DC24(fm0(v0[safeIndex(723, v0.Length)], v0[safeIndex(723, v0.Length)], globalState), v32, v33, v1);
		var v35: map<D9, int> := map[v34 := -v1];
		var v36: multiset<bool> := multiset{p0};
		var v37 := "qaoscxjb";
		v35 := v35[v34 := safeModuloInt(if (p0 in v36) then v36[p0] else |v37|, v0[safeIndex(723, v0.Length)])];
		globalState.f2 := p0;
		globalState.f2 := p0;
		var v38: map<bool, char> := map[p0 := v33];
		var v39 := DC20(v33);
		var v40: array<char> := new char[26] ['q', v33, v33, v33, DC24(v0[safeIndex(723, v0.Length)], v32, 'g', v1).cf29, v33, v33, v33, 'x', if (p0 in v38) then v38[p0] else v33, 's', v33, v33, v33, v33, 'w', v33, v33, v39.cf22, v33, v33, v33, v33, 'k', v33, v33];
		v40[safeIndex(167, v40.Length)] := v33;
		var v41: set<int> := {v0[safeIndex(723, v0.Length)]};
		v0[safeIndex(723, v0.Length)], v40[safeIndex(167, v40.Length)], globalState.f3 := if (p0) then -0x21 else |v41|, v33, v1;
		var v42: map<string, int> := map[if (true) then v37 else seq(abs(-0x92), i6  => (v40[safeIndex(167, v40.Length)])) := v0[safeIndex(723, v0.Length)] - 0x1dc];
		v42 := v42;
	} else {
		if (p0) {
			var v43 := "cnu";
			var v44 := 'n';
			var v45 := DC1(0x157 - v1, v43 <= v43[safeIndex(0x14f, |v43|) := v44], |v2|);
			v45 := v45;
			globalState.f4 := v0[safeIndex(723, v0.Length)];
			var v46 := DC3();
			var v47: map<D0, array<int>> := map[v46 := v0];
			v47 := v47[v46 := v0];
			var v48: map<bool, bool> := map[p0 := p0];
			var v50: map<D0, seq<int>> := map[v45 := v3];
			var v51: C0 := new C0(map v49 : D0 | v49 in v50 :: (v49) := (|v43|));
			var v52 := DC26(map[v48 := v51]);
			v1 := |v52.cf31| * v1;
			v44 := v44;
		} else {
			var v53 := "kqlo";
			var v54: array<seq<bool>> := new seq<bool>[1](i7 => v2 + v2);
			v54[safeIndex(619, v54.Length)] := [!true, p0, p0, p0, p0];
			v53, globalState.f4, v54[safeIndex(619, v54.Length)], v0[safeIndex(723, v0.Length)] := if (p0) then v53 else v53, v0[safeIndex(723, v0.Length)], v2, v0[safeIndex(723, v0.Length)];
			var v55: map<int, int> := map[v1 := v0[safeIndex(723, v0.Length)]];
			var v57: seq<map<int, int>> := [map[fm0(v0[safeIndex(723, v0.Length)], 0x39b, globalState) := |multiset{v1, -0x87}|]];
			var v58: multiset<map<int, int>> := multiset{v55, v55, map v56 : int | (0x2ae <= v56) && (v56 < 0x2f6) :: (v56 + 0x183) := (v1), v57[safeIndex(v1, |v57|)], v55};
			globalState.f2 := v58 > (v58 + v58);
			var v59: array<bool> := new bool[18];
			var v60 := DC17(v0);
			var v61: map<int, array<int>> := map[0x1c3 := v60.cf19];
			v59[safeIndex(637, v59.Length)] := v61 != v61;
			v59[safeIndex(637, v59.Length)] := p0;
			globalState.f2 := v59[safeIndex(637, v59.Length)];
			globalState.f2 := -v0[safeIndex(723, v0.Length)] != 0x26d;
		}
		
		v3 := [494];
		v0 := v0;
		var v62 := DC1(v0[safeIndex(723, v0.Length)], v2[safeIndex(v1, |v2|)], v0[safeIndex(723, v0.Length)]);
		var v63: map<D0, int> := map[v62 := v1];
		var v64: C0 := new C0(v63);
		var v65 := 'k';
		var v66 := DC24(v1, v64, v65, v0[safeIndex(723, v0.Length)]);
		var v67: map<D9, bool> := map[v66 := p0];
		var v68: array<seq<bool>> := new seq<bool>[9] [v2, v2, v2, [p0, p0], if (p0) then v2 else v2, [!false, if (v66 in v67) then v67[v66] else p0], [p0, p0], v2, v2];
		var v69: array<D6> := new D6[21];
		var v70 := DC15(v1, v1);
		v69[safeIndex(512, v69.Length)] := v70;
		v68, v69[safeIndex(512, v69.Length)] := v68, v70.(cf17 := v1);
		var v71: array<T0> := new T0[26];
		var v72: map<array<T0>, bool> := map[v71 := p0];
		globalState.f4, globalState.f2, v0[safeIndex(723, v0.Length)], v0[safeIndex(723, v0.Length)] := -((v0[safeIndex(723, v0.Length)] * -v1) - v0[safeIndex(723, v0.Length)]), p0 <==> !(v0[safeIndex(723, v0.Length)] <= |v72|), v0[safeIndex(723, v0.Length)], fm0(v1, v1, globalState);
	}
	
	var v73 := 'f';
	var v74: seq<char> := [v73];
	for i8 := if (p0) then v1 else |v74| to v1 {
		globalState.f2 := p0 ==> p0;
		var v75: set<bool> := {if (p0) then p0 else p0};
		v75 := v75;
		var v76: multiset<bool> := multiset{false};
		v76 := multiset{p0, p0, p0, p0, true} + (v76 * v76);
		var v77: set<int> := {v1, i8};
		var v78 := DC1(i8, p0, |v77|);
		var v79: map<D0, int> := map[v78.(cf3 := i8) := v1];
		var v80: T0 := new C0(v79);
		var v81: array<T0> := new T0[2] [v80, v80];
		v81 := if (p0) then v81 else v81;
	}
	var v82: set<seq<int>> := {v3};
	globalState.f2 := if (v82 > v82) then p0 else v1 > v1;
}
method Main() {
	var v0 := 0x36d;
	var v1: set<int> := {-v0};
	var v2 := "yxsvjmbqp";
	var v3: set<string> := {v2, v2};
	var v4: array<int> := new int[20] [v0, v0, v0, v0, v0, 0x30e, v0, v0, v0, DC0(v0).cf0, v0, v0, |v1|, -|v3|, v0, v0, v0, -v0, v0, 0xd9];
	var v5: array<bool> := new bool[12](i0 => DC1(v0, true, v0).cf2);
	var globalState := new GlobalState(false, true, true, -0x2bd, 0x2f5, 0x3af, v4, true, v4, v5, 0x2ad, false, v2, 0x33a, true);
	var i1 := 0;
	while (v0 == v0)
		decreases 100 - i1
	{
		if (i1 >= 100) {
			break;
		}
		
		i1 := i1 + 1;
		var v6: map<int, int> := map[v0 := v0];
		v4[safeIndex(621, v4.Length)] := |(v6 + v6)|;
		var v7 := true;
		var v8 := 'i';
		v4[safeIndex(621, v4.Length)] := fm0(if (v7) then v0 else -650, |(if (true) then seq(abs(0x1b6), i2  => ('u')) else v2)[safeIndex(v0, |(if (true) then seq(abs(0x1b6), i2  => ('u')) else v2)|) := v8]|, globalState);
		v4[safeIndex(621, v4.Length)] := fm0(v0, v0, globalState);
		if (v0 < v0) {
			var v9: multiset<int> := multiset{v0};
			var v10: seq<string> := [v2];
			var v11: multiset<int> := multiset{fm0(v4[safeIndex(621, v4.Length)], |v9|, globalState), v4[safeIndex(621, v4.Length)], |v10[safeIndex(if (v0 in v9) then v9[v0] else v0, |v10|)]|, v0};
			var v12: array<string> := new string[20](i3 => v2);
			v12[safeIndex(607, v12.Length)] := v2;
			var v13 := DC4(v2);
			v4[safeIndex(621, v4.Length)], v9, v2, globalState.f2, v12[safeIndex(607, v12.Length)] := v0, v11, v2, v7, v13.cf4;
			var v14: seq<int> := [v4[safeIndex(621, v4.Length)], v4[safeIndex(621, v4.Length)], v4[safeIndex(621, v4.Length)], v0];
			globalState.f2 := v0 !in v14;
			var v15: multiset<string> := multiset{v2, v12[safeIndex(607, v12.Length)]};
			v12[safeIndex(607, v12.Length)], globalState.f2, globalState.f2, globalState.f4 := v2[safeIndex(v4[safeIndex(621, v4.Length)], |v2|) := v8] + v13.cf4, v0 != (v0 * v14[safeIndex(v0, |v14|)]), v15 > v15, v0;
			var v16 := DC1(-v4[safeIndex(621, v4.Length)], v7, |v2|);
			var v17 := new C0(map[v16 := v4[safeIndex(621, v4.Length)]]);
			globalState.f3 := v4[safeIndex(621, v4.Length)] - v4[safeIndex(621, v4.Length)];
		} else {
			m0("lkmnn" != v2, globalState);
			var v18: array<array<bool>> := new array<bool>[3] [if (v7) then v5 else v5, v5, v5];
			v18[safeIndex(506, v18.Length)] := v5;
			v18[safeIndex(506, v18.Length)] := v5;
			var v19 := DC2();
			v19 := v19;
			var v20: array<string> := new seq<char>[4](i4 => if (v7) then seq(abs(0x396), i5  => (v8)) else v2);
			var v21: map<bool, string> := map[!v7 := "hcnuuyv"];
			v20[safeIndex(931, v20.Length)] := if (v7 in v21) then v21[v7] else fm1(globalState);
			v20[safeIndex(931, v20.Length)] := v2 + v2;
			var v22: multiset<int> := multiset{v4[safeIndex(621, v4.Length)], v4[safeIndex(621, v4.Length)]};
			m0(0x241 !in v22, globalState);
		}
		
		if (v0 == v4[safeIndex(621, v4.Length)]) {
			globalState.f3 := ---(safeModuloInt(v4[safeIndex(621, v4.Length)], v0) + v0);
			globalState.f4 := v0;
			v4[safeIndex(621, v4.Length)] := -v0;
			var v23: array<map<bool, int>> := new map<bool, int>[13](i6 => map[false := -0x2] + map[v7 := |map[v4[safeIndex(621, v4.Length)] := v7]|]);
			var v24: map<bool, int> := map[v7 := v0];
			v23[safeIndex(845, v23.Length)] := v24;
			v23[safeIndex(845, v23.Length)] := v24[v7 := 0x19a];
			var v25: map<int, multiset<int>> := map[v0 := multiset{0x12c, 0x3a}];
			var v26: seq<int> := [v0, v0];
			var v27 := DC5(multiset(v26));
			v25 := v25[v0 := v27.cf5];
		} else {
			var v29: multiset<bool> := multiset{v7};
			var v30 := DC1(v4[safeIndex(621, v4.Length)], v7, v4[safeIndex(621, v4.Length)]);
			var v31: map<D0, int> := map[v30 := v0];
			var v32: T0 := new C0(v31);
			var v33: multiset<T0> := multiset{v32, v32};
			var v34 := DC1(|v33|, v7, v0);
			var v35: multiset<D0> := multiset{DC1(v4[safeIndex(621, v4.Length)], v7, if (v7 in v29) then v29[v7] else v0), v34};
			var v36: T0 := new C0(map v28 : D0 | v28 in v35 :: (v28) := (v4[safeIndex(621, v4.Length)]));
			var v37: array<int> := new int[11];
			var v38: map<T0, array<int>> := map[v32 := v37];
			var v39: C0 := new C0(v31);
			var v40: set<C0> := {v39};
			v7, globalState.f2, v7, globalState.f2, v38 := v7, v7, (v40 - v40) !! (v40 * {v39, v39, v39}), v7, v38 + DC8(v38).cf10;
			globalState.f2, v4[safeIndex(621, v4.Length)] := v7, v4[safeIndex(621, v4.Length)] + -v4[safeIndex(621, v4.Length)];
			var v41: set<bool> := {v7, v7, v7, v7, v7};
			var v42: seq<set<bool>> := [v41, v41, v41, v41];
			var v43: map<bool, multiset<set<bool>>> := map[v7 := multiset(v42)];
			var v44: multiset<set<bool>> := multiset{v41, v41, v41};
			var v45: map<int, multiset<set<bool>>> := map[|v41| := v44];
			v43 := v43[v7 := if (v0 in v45) then v45[v0] else v44];
			var v46: array<multiset<int>> := new multiset<int>[4];
			v46 := v46;
			var v47 := new C0(map[fm2(v7, globalState) := v4[safeIndex(621, v4.Length)]]);
		}
		
	}
	var v48 := false;
	var v49: seq<bool> := [v48];
	if (v48 <==> (v49 < [true])) {
		var v50: multiset<string> := multiset{v2, "eueosg" + v2, v2};
		v0 := |v50|;
		var v51 := 'y';
		v51 := v51;
		var v52 := DC1(v0, v48, fm0(v0, v0, globalState));
		var v54: map<D0, int> := map[v52 := v0];
		var v55 := new C0(map[v52 := v0] + (map v53 : D0 | v53 in v54 :: (v53) := (v52.cf3)));
		var v56: seq<int> := [v0];
		if (v0 !in v56) {
			v5[safeIndex(151, v5.Length)] := v48;
			var v57: map<int, int> := map[v0 := v0];
			var v58: T0 := new C0(v54);
			var v59: seq<T0> := [v58, v58];
			var v60: map<int, seq<T0>> := map[safeModuloInt(|v57|, v0) := DC9(v59).cf11 + v59];
			var v61: map<T0, int> := map[v58 := v0];
			var v62: map<T0, map<T0, int>> := map[v58 := if (v48) then v61 else v61];
			v5[safeIndex(151, v5.Length)], v60, v62 := v48, map[v0 := v59] + v60, v62;
			var v63 := DC13(v54);
			var v64: C0 := new C0(v63.cf14);
			v64 := v64;
			v5[safeIndex(151, v5.Length)] := if (v0 == |multiset{v0}|) then v5[safeIndex(151, v5.Length)] else v48;
			v58 := new C0(v55.f15);
			v5[safeIndex(151, v5.Length)] := if (v51 !in v2) then false else v5[safeIndex(151, v5.Length)];
		} else {
			v4[safeIndex(553, v4.Length)] := |v56|;
			var v65: multiset<bool> := multiset{v48};
			var v66: map<int, multiset<bool>> := map[v0 := multiset{v48}];
			var v67: multiset<int> := multiset{v0};
			var v68: array<multiset<bool>> := new multiset<bool>[22] [multiset{!true, v48}, v65, v65, v65, v65, multiset(v49), v65, v65, v65, multiset(v49), if (|(seq(abs(0x37c), i7  => (v0)))[safeIndex(v0, |(seq(abs(0x37c), i7  => (v0)))|) := v0]| in v66) then v66[|(seq(abs(0x37c), i7  => (v0)))[safeIndex(v0, |(seq(abs(0x37c), i7  => (v0)))|) := v0]|] else v65, v65, v65, v65, v65, v65, v65, multiset(v49), v65, fm3(true, v67, v48, globalState), multiset(v49), v65];
			var v69: map<bool, array<bool>> := map[v48 := v5];
			var v70: map<array<multiset<bool>>, array<bool>> := map[v68 := if (v48 in v69) then v69[v48] else v5];
			var v71: seq<array<bool>> := [v5];
			var v72: array<array<bool>> := new array<bool>[25] [v5, v5, if (v68 in v70) then v70[v68] else v5, v5, v5, v5, v5, v5, v5, v5, v5, v71[safeIndex(v0, |v71|)], v5, v5, v5, v5, DC14(v5).cf15, v5, if (v48) then v5 else v5, v5, v5, v5, if (v48) then v5 else v5, v5, v5];
			v72[safeIndex(378, v72.Length)] := v5;
			var v73: map<seq<char>, int> := map[v2 := v0];
			globalState.f3, v2, v4[safeIndex(553, v4.Length)], v72[safeIndex(378, v72.Length)] := safeDivisionInt(safeDivisionInt(-v0, v0), if (v2 in v73) then v73[v2] else v0), v2, 0xd1 - 602, v5;
			var v74 := new C0(map[v52 := v0]);
			var v75: map<int, string> := map[v0 := v2];
			v75 := v75[v4[safeIndex(553, v4.Length)] := seq(abs(0x284), i8  => (v51))];
			globalState.f4 := -(v4[safeIndex(553, v4.Length)] + 820);
			v56, v4[safeIndex(553, v4.Length)] := v56, |fm4(!v48, v48, globalState)|;
		}
		
		globalState.f2 := v48;
	} else {
		var v76 := 'x';
		var v77: map<char, bool> := map[v76 := v2 == v2];
		var v78: map<int, bool> := map[v0 := v48];
		var v79 := DC1(116, v48, v0);
		var v80: map<D0, int> := map[v79 := v0];
		var v81: C0 := new C0(v80[v79 := v0]);
		var v82: map<C0, C0> := map[v81 := v81];
		var v83: map<map<C0, C0>, int> := map[map[v81 := v81] := v0];
		var v84: seq<int> := [|v49|];
		v77, v78 := if (v82[v81 := v81] !in v83) then v77 else v77, v78[v0 := (set v85 : int | v85 in {v84[safeIndex(v0, |v84|)]} :: (safeModuloInt(v85, -0x5d))) > (set v86 : int | (0x11 <= v86) && (v86 < 0x257) :: (v86 * 0x388))];
		var v87: map<seq<string>, array<int>> := map[[v2, "jlvlgk"] := v4];
		v87 := v87[fm5(map[v48 := v0], v48, v0, v0, globalState) := v4];
		m0(v48, globalState);
		var v88: map<array<int>, int> := map[v4 := v0];
		v88 := v88;
		v4[safeIndex(557, v4.Length)] := v0 * v0;
		var v89: array<seq<int>> := new seq<int>[1](i9 => v84);
		var v90: map<bool, seq<int>> := map[fm6(fm6(v48, |{v4, v4, v4, DC17(v4).cf19, v4}|, globalState), v0, globalState) := [-v0, v0]];
		var v91: map<bool, bool> := map[v48 := !v48];
		var v92 := DC6(v91, v48);
		v89[safeIndex(76, v89.Length)] := if (!v92.cf7 in v90) then v90[!v92.cf7] else [v0];
		v4[safeIndex(557, v4.Length)], globalState.f3, v89[safeIndex(76, v89.Length)] := v0 * v0, (v0 * v0) + v0, v84;
	}
	
	var i10 := 0;
	while (false)
		decreases 100 - i10
	{
		if (i10 >= 100) {
			break;
		}
		
		i10 := i10 + 1;
		var v93 := DC1(v0, v48, v0);
		var v94: map<D0, int> := map[v93 := -v0];
		var v95 := DC13(v94);
		match (v95) {
			case DC13(cf14) =>
				var v96: map<bool, bool> := map[v48 := v48];
				v48 := (|v96| <= v0) <==> v48;
				m0(v48, globalState);
				v0 := if (!v48) then v0 else v0;
				var v97 := DC6(v96, v48);
				v5[safeIndex(371, v5.Length)] := v97.cf7 <==> v48;
				v5[safeIndex(371, v5.Length)] := v48;
		}
		
		globalState.f2 := !v48;
		v0, globalState.f4 := |v2|, v0;
		if (v48) {
			var v98 := DC0(v0);
			globalState.f2 := v98.cf0 > v0;
			v2 := "uuglv";
			var v99 := new C0(v94);
			globalState.f4 := |fm1(globalState)| + 0x2f7;
			var v100: seq<int> := [v0, 0x261, v0];
			v4[safeIndex(27, v4.Length)] := if (v48) then v0 else |v100|;
			v4[safeIndex(27, v4.Length)] := v0;
		} else {
			var v101: map<int, int> := map[v0 * v0 := v0 - v0];
			v101 := v101;
			globalState.f3 := v0;
			m0(v49 <= [v48], globalState);
			globalState.f3 := v0;
			var v102 := new C0(map[v93.(cf3 := v0) := v0]);
		}
		
	}
	match (DC0(v0)) {
		case DC1(cf1, cf2, cf3) =>
			var v103 := 'o';
			v103 := 'u';
			var v104: array<array<int>> := new array<int>[27];
			v104[safeIndex(490, v104.Length)] := v4;
			v104[safeIndex(490, v104.Length)] := new int[22];
			v5[safeIndex(207, v5.Length)] := v48;
			v5[safeIndex(207, v5.Length)] := v48;
			cf1 := fm0(safeModuloInt(cf3, cf1), fm0(cf3, cf3, globalState), globalState);
		case DC2() =>
			globalState.f2 := v48;
			v5[safeIndex(690, v5.Length)] := false;
			var v105: map<bool, bool> := map[v48 := v48];
			var v106: multiset<int> := multiset{|v105|};
			var v107 := DC5(v106);
			var v108: map<D2, bool> := map[if (v48) then DC5(v106) else v107 := v0 == 0x3b9];
			globalState.f2, v5[safeIndex(690, v5.Length)], globalState.f4, globalState.f2 := !!(0x3bc < -v0), true, v0, if (v107 in v108) then v108[v107] else v48;
			v0 := v0;
			var v109 := DC6(v105, v5[safeIndex(690, v5.Length)]);
			var v110 := 'x';
			var v111: map<D2, char> := map[v109 := v110];
			var v112: set<char> := {if (v109 in v111) then v111[v109] else v110};
			var v113: seq<int> := [fm0(v0, |"jfgpdshm"|, globalState), v0];
			var v114: seq<int> := [|multiset(v113)|];
			v112, globalState.f2, v110 := fm7(v110, |(fm8(v0, globalState) + v114)|, v0, v48, globalState), v5[safeIndex(690, v5.Length)], v110;
		case DC3() =>
			globalState.f2 := v48;
			var v115 := DC0(v0);
			v115 := v115;
			globalState.f2 := v48;
			if (v48) {
				var v116 := DC1(|v2|, !v48, v0);
				var v117: map<D0, int> := map[v116 := |v49|];
				var v118 := new C0(v117);
				v2 := v2 + (v2 + v2);
				var v119 := new C0(v118.f15);
				v119 := v118;
				v5[safeIndex(125, v5.Length)] := v0 < v0;
				v5[safeIndex(125, v5.Length)] := fm6(false, 0x25b, globalState);
			} else {
				v2 := (v2 + v2) + v2;
				var v120 := DC1(--v0, v48, v0);
				var v121: map<D0, int> := map[v120 := v0];
				var v122 := new C0(v121);
				var v124: seq<D0> := [v120, v120];
				var v125: T0 := new C0(map v123 : D0 | v123 in v124 :: (v123) := (v0));
				v4[safeIndex(339, v4.Length)] := v0;
				var v126: map<bool, bool> := map[v48 := v48];
				var v127 := DC6(v126, v48 && v48);
				v125, v4[safeIndex(339, v4.Length)], globalState.f2, globalState.f3, v127 := v125, safeModuloInt(v0, v0), v48, v0, v127;
				var v128 := 'c';
				var v129 := DC20(v128);
				v128 := v129.cf22;
				var v130: array<T0> := new T0[6];
				v130[safeIndex(44, v130.Length)] := v125;
				var v131: map<bool, T0> := map[v48 := v125];
				v130[safeIndex(44, v130.Length)] := if (v48 in v131) then v131[v48] else v125;
			}
			
		case DC0(cf0) =>
			var v132 := 'i';
			var v133: multiset<char> := multiset{v132, v132, v132, v132};
			globalState.f2 := fm6(v48, |v133|, globalState);
			v5[safeIndex(591, v5.Length)] := v49[safeIndex(v0, |v49|)];
			var v134: seq<int> := [v0];
			v5[safeIndex(591, v5.Length)] := safeModuloInt(0x2b6, |[v48]|) <= v134[safeIndex(cf0, |v134|)];
			var v135: set<bool> := {v48};
			var v136: map<D0, int> := map[DC1(|v135|, v48, v0) := -cf0];
			var v137: C0 := new C0(v136);
			var v138: map<bool, C0> := map[v48 := v137];
			if (fm9(|v138|, fm0(cf0, v0, globalState), globalState) >= v1) {
				globalState.f4 := safeModuloInt(cf0, cf0) - -0xa7;
				var v139: map<char, bool> := map['v' := v5[safeIndex(591, v5.Length)]];
				m0(if (v132 in v139) then v139[v132] else true, globalState);
				v137 := v137;
				var v141: multiset<int> := multiset{cf0};
				var v142 := DC1(if (cf0 in v141) then v141[cf0] else v0, v5[safeIndex(591, v5.Length)], cf0);
				var v143: map<D0, bool> := map[v142 := v5[safeIndex(591, v5.Length)]];
				var v144 := new C0((map v140 : D0 | v140 in v143 :: (v140) := (v0)) + map[v142 := v0]);
				globalState.f2 := v132 !in (seq(abs(0x2ee), i11  => ('j')));
			} else {
				v132 := DC20(v132).cf22;
				m0(v5[safeIndex(591, v5.Length)], globalState);
				globalState.f3 := v0;
				var v145: map<bool, int> := map[v5[safeIndex(591, v5.Length)] := 189];
				var v146: map<int, int> := map[v0 := if (true in v145) then v145[true] else cf0];
				var v147: map<int, int> := map[if (cf0 in v146) then v146[cf0] else |"yeoger"| := cf0];
				var v148: seq<map<int, int>> := [map[0x31c := 480]];
				v147 := (v148[safeIndex(cf0, |v148|)] + map[|v2| := 0x1f7]) + v147;
				v2 := "vfpkjomp";
			}
			
			var v149: map<seq<bool>, C0> := map[fm10(globalState) := v137];
			cf0 := |v149|;
	}
	
	v4[safeIndex(642, v4.Length)] := v0;
	var v150: set<array<bool>> := {v5, v5, v5, v5};
	v4[safeIndex(642, v4.Length)] := safeModuloInt(|v150|, 0x148);
	var v151: map<int, bool> := map[v0 := fm6(v48, v4[safeIndex(642, v4.Length)], globalState)];
	v5[safeIndex(359, v5.Length)] := v48;
	v4[safeIndex(642, v4.Length)], v48, v151, v4, v5[safeIndex(359, v5.Length)] := v4[safeIndex(642, v4.Length)] + -v0, v48, v151, v4, 0x3bc != v0;
	var v152: multiset<set<int>> := multiset{v1};
	var v153: seq<multiset<set<int>>> := [v152];
	var i12 := 0;
	while ((v152 + v152) > (v153[safeIndex(v0, |v153|)] + v152))
		decreases 100 - i12
	{
		if (i12 >= 100) {
			break;
		}
		
		i12 := i12 + 1;
		var v154 := DC1(v4[safeIndex(642, v4.Length)], v5[safeIndex(359, v5.Length)], v4[safeIndex(642, v4.Length)]);
		var v155: map<D0, int> := map[v154 := v0];
		var v156: C0 := new C0(v155);
		var v157: map<C0, string> := map[v156 := v2];
		v5[safeIndex(359, v5.Length)], v2, v5[safeIndex(359, v5.Length)] := v5[safeIndex(359, v5.Length)], if (v156 in v157) then v157[v156] else v2, true && !v48;
		var v158 := 'd';
		var v159 := DC20(v158);
		v48 := v159 !in (seq(abs(0x3d6), i13  => (v159)));
		v156 := v156;
		var v160: map<bool, C0> := map[!v5[safeIndex(359, v5.Length)] := v156];
		v160 := v160;
	}
	globalState.f2, v49 := v48, [v48];
	globalState.f2 := v48;
	var v161: array<map<bool, int>> := new map<bool, int>[4](i15 => map[v48 := v0]);
	forall i14 | 0 <= i14 < v161.Length {
		v161[i14] := map[false := v0] + map[v5[safeIndex(359, v5.Length)] := v4[safeIndex(642, v4.Length)]];
	}
	m0(v5[safeIndex(359, v5.Length)] <== v48, globalState);
	var v162 := DC4(DC4(v2).cf4);
	v4[safeIndex(642, v4.Length)] := match v162 {
		case DC4(cf4) => -(v0 * |multiset{-v4[safeIndex(642, v4.Length)]}|)
	};
	v0 := v0;
	var v163 := 'e';
	var v164: map<char, bool> := map[v163 := v48];
	var v165: map<bool, int> := map[v48 := |v164|];
	var v166 := DC1(v0, v48, -|v165|);
	var v167: map<D0, int> := map[v166 := v0];
	var v168: C0 := new C0(v167);
	var v169: multiset<map<C0, C0>> := multiset{map[v168 := v168]};
	var v170 := DC21(v0, v168);
	var v171: map<C0, C0> := map[v170.cf24 := v168];
	v169 := v169[v171 := abs(v0)];
	var i16 := 0;
	while (v5[safeIndex(359, v5.Length)] || v48)
		decreases 100 - i16
	{
		if (i16 >= 100) {
			break;
		}
		
		i16 := i16 + 1;
		v165 := v165[v48 := v0];
		v0 := -0x3c7 + |(v1 + v1)|;
		v2 := v2;
		v165 := v165[v48 := v4[safeIndex(642, v4.Length)]];
	}
	var v172: array<array<bool>> := new array<bool>[3];
	var v173: map<array<array<bool>>, bool> := map[v172 := v48];
	v173 := v173[v172 := v48];
	print v0, "\n";
	print v1 == {-877}, "\n";
	print v2, "\n";
	print v3 == {"yxsvjmbqp"}, "\n";
	print v4[0], "\n";
	print v4[1], "\n";
	print v4[2], "\n";
	print v4[3], "\n";
	print v4[4], "\n";
	print v4[5], "\n";
	print v4[6], "\n";
	print v4[7], "\n";
	print v4[8], "\n";
	print v4[9], "\n";
	print v4[10], "\n";
	print v4[11], "\n";
	print v4[12], "\n";
	print v4[13], "\n";
	print v4[14], "\n";
	print v4[15], "\n";
	print v4[16], "\n";
	print v4[17], "\n";
	print v4[18], "\n";
	print v4[19], "\n";
	print v5[0], "\n";
	print v5[1], "\n";
	print v5[2], "\n";
	print v5[3], "\n";
	print v5[4], "\n";
	print v5[5], "\n";
	print v5[6], "\n";
	print v5[7], "\n";
	print v5[8], "\n";
	print v5[9], "\n";
	print v5[10], "\n";
	print v5[11], "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2, "\n";
	print globalState.f3, "\n";
	print globalState.f4, "\n";
	print globalState.f5, "\n";
	print globalState.f6[0], "\n";
	print globalState.f6[1], "\n";
	print globalState.f6[2], "\n";
	print globalState.f6[3], "\n";
	print globalState.f6[4], "\n";
	print globalState.f6[5], "\n";
	print globalState.f6[6], "\n";
	print globalState.f6[7], "\n";
	print globalState.f6[8], "\n";
	print globalState.f6[9], "\n";
	print globalState.f6[10], "\n";
	print globalState.f6[11], "\n";
	print globalState.f6[12], "\n";
	print globalState.f6[13], "\n";
	print globalState.f6[14], "\n";
	print globalState.f6[15], "\n";
	print globalState.f6[16], "\n";
	print globalState.f6[17], "\n";
	print globalState.f6[18], "\n";
	print globalState.f6[19], "\n";
	print globalState.f7, "\n";
	print globalState.f8[0], "\n";
	print globalState.f8[1], "\n";
	print globalState.f8[2], "\n";
	print globalState.f8[3], "\n";
	print globalState.f8[4], "\n";
	print globalState.f8[5], "\n";
	print globalState.f8[6], "\n";
	print globalState.f8[7], "\n";
	print globalState.f8[8], "\n";
	print globalState.f8[9], "\n";
	print globalState.f8[10], "\n";
	print globalState.f8[11], "\n";
	print globalState.f8[12], "\n";
	print globalState.f8[13], "\n";
	print globalState.f8[14], "\n";
	print globalState.f8[15], "\n";
	print globalState.f8[16], "\n";
	print globalState.f8[17], "\n";
	print globalState.f8[18], "\n";
	print globalState.f8[19], "\n";
	print globalState.f9[0], "\n";
	print globalState.f9[1], "\n";
	print globalState.f9[2], "\n";
	print globalState.f9[3], "\n";
	print globalState.f9[4], "\n";
	print globalState.f9[5], "\n";
	print globalState.f9[6], "\n";
	print globalState.f9[7], "\n";
	print globalState.f9[8], "\n";
	print globalState.f9[9], "\n";
	print globalState.f9[10], "\n";
	print globalState.f9[11], "\n";
	print globalState.f10, "\n";
	print globalState.f11, "\n";
	print globalState.f12, "\n";
	print globalState.f13, "\n";
	print globalState.f14, "\n";
	print i1, "\n";
	print v48, "\n";
	print v49 == [false], "\n";
	print i10, "\n";
	print |v150|, "\n";
	print v151 == map[3 := false], "\n";
	print v152 == multiset{{-877}}, "\n";
	print v153 == [multiset{{-877}}], "\n";
	print i12, "\n";
	print v161[0] == map[false := 3, true := -2], "\n";
	print v161[1] == map[false := 3, true := -2], "\n";
	print v161[2] == map[false := 3, true := -2], "\n";
	print v161[3] == map[false := 3, true := -2], "\n";
	print v162.cf4, "\n";
	print v163, "\n";
	print v164 == map['e' := false], "\n";
	print v165 == map[false := -3], "\n";
	print v166.cf1, "\n";
	print v166.cf2, "\n";
	print v166.cf3, "\n";
	print v167 == map[DC1(3, false, -1) := 3], "\n";
	print v168.f15 == map[DC1(3, false, -1) := 3], "\n";
	print |v169|, "\n";
	print v170.cf23, "\n";
	print v170.cf24.f15 == map[DC1(3, false, -1) := 3], "\n";
	print |v171|, "\n";
	print i16, "\n";
	print |v173|, "\n";
}