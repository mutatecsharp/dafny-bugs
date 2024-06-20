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
datatype D0 = DC1 | DC0(cf0: map<seq<bool>, array<int>>) | DC2(cf1: D0)
datatype D1 = DC4(cf3: set<string>) | DC3(cf2: array<multiset<bool>>) | DC5(cf4: D1)
datatype D2 = DC6(cf5: map<bool, bool>)
datatype D3 = DC7(cf6: int)
datatype D4 = DC8(cf7: bool)
datatype D5 = DC10(cf9: set<int>, cf10: map<multiset<int>, string>) | DC11(cf11: bool, cf12: int, cf13: bool, cf14: array<bool>, cf15: int) | DC9(cf8: C0) | DC12(cf16: D5)
datatype D6 = DC13(cf17: string)
datatype D7 = DC15(cf19: C0, cf20: int, cf21: int, cf22: char) | DC16(cf23: char, cf24: array<C0>, cf25: int, cf26: int, cf27: bool) | DC17(cf28: bool, cf29: int, cf30: bool, cf31: bool) | DC14(cf18: multiset<int>)
class GlobalState {
	const f0 : bool
	const f1 : string
	const f2 : map<char, array<int>>
	const f3 : set<bool>
	var f4 : map<string, int>
	const f5 : int
	const f6 : bool
	var f7 : int
	var f8 : int
	var f9 : seq<bool>
	const f10 : bool
	const f11 : array<int>
	const f12 : bool
	const f13 : map<bool, bool>
	const f14 : int
	const f15 : int
	const f16 : bool
	const f17 : bool
	const f18 : bool
	constructor (f0 : bool, f1 : string, f2 : map<char, array<int>>, f3 : set<bool>, f4 : map<string, int>, f5 : int, f6 : bool, f7 : int, f8 : int, f9 : seq<bool>, f10 : bool, f11 : array<int>, f12 : bool, f13 : map<bool, bool>, f14 : int, f15 : int, f16 : bool, f17 : bool, f18 : bool) {
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
		this.f15 := f15;
		this.f16 := f16;
		this.f17 := f17;
		this.f18 := f18;
	}
	
}

class C0 {
	var f19 : bool
	const f20 : map<bool, bool>
	constructor (f19 : bool, f20 : map<bool, bool>) {
		this.f19 := f19;
		this.f20 := f20;
	}
	
	function fm2(p0: char, p1: bool, globalState: GlobalState): int {
		-|({|"d"|, |f20|, |[|{|"p"|}|]|} * ({689} + {0x1d4, |"doisvt"|}))|
	}
	function fm3(p0: bool, p1: int, globalState: GlobalState): D0 {
		match DC1() {
			case DC1() => DC1()
			case DC0(cf0) => DC1()
			case DC2(cf1) => DC1()
		}
	}
}

function fm0(p0: bool, p1: bool, p2: string, globalState: GlobalState): bool {
	(|[true, !false]| * 0x348) < (-0x201 + 0x372)
}
function fm1(p0: bool, p1: int, p2: int, globalState: GlobalState): int {
	0x2e0
}
function fm4(p0: bool, p1: bool, p2: bool, p3: int, globalState: GlobalState): string {
	"euwd" + "d"
}
function fm5(p0: multiset<bool>, globalState: GlobalState): set<string> {
	set v0 : string | v0 in multiset{seq(abs(0x241), i0  => ('e')), seq(abs(-0xd3), i1  => ('k')), "bmpbyeb", "jeeg"} :: (v0)
}
function fm6(p0: seq<seq<int>>, p1: bool, p2: bool, globalState: GlobalState): D1 {
	DC4((set v0 : string | v0 in (seq(abs(240), i0  => (seq(abs(0x2cb), i1  => ('m'))))) :: (v0)) * {"khuk", "gb", "gbthcxlck"})
}
function fm7(globalState: GlobalState): seq<int> {
	([582, DC17(false, |map[0x79 := 0x385]|, false, true).cf29, -450] + [630, 0x2ca]) + [0x225]
}
function fm8(p0: int, p1: bool, globalState: GlobalState): seq<bool> {
	[0x210 != |map[!!true := DC17(true, |[0x2c2, 932]|, true, !true).cf30]|, |multiset{|(set v0 : int | (0x1d5 <= v0) && (v0 < -0x2f) :: (v0 - 437))|, -|{false}|, |map[false := map[!!true := |map[false := true]|]]|, 581, --0x30}| >= 0x231, false, !(|[[!true, !true], [true, true, true], [false, false, false, false, true], [!true, true, true], [true, false, true]]| == |multiset{false, false}|), false]
}
function fm9(p0: int, p1: int, p2: int, globalState: GlobalState): map<bool, bool> {
	(map[true := true] + map[true := true]) + (map[true := false] + map[false := true])
}
function fm10(p0: seq<int>, p1: D4, globalState: GlobalState): char {
	if (true || !true) then 'j' else 'u'
}
function fm11(p0: char, p1: bool, p2: int, p3: int, globalState: GlobalState): D4 {
	DC8((seq(abs(0x41), i0  => ('x'))) == "xm")
}
function fm12(p0: seq<seq<int>>, p1: bool, p2: bool, globalState: GlobalState): set<bool> {
	({false} - {true, false}) + ({true, false} - {true})
}
function fm13(p0: bool, p1: int, globalState: GlobalState): map<multiset<int>, int> {
	map[multiset([|[|map[0x3da := |multiset{true}|]|]|, |{'k'}|]) + multiset{|[-994, |map[false := -991]|]|} := safeModuloInt(|map[-577 := map[0x294 := DC7(0x197)]]|, 0x248)]
}
method m0(globalState: GlobalState) returns (r0: int, r1: int) {
	var v0 := 0x12d;
	var v1 := false;
	var v2: map<bool, char> := map[v1 := 'l'];
	for i0 := v0 to -|(v2 + map[v1 := 'k'])| {
		r1 := v0;
		v0 := v0;
		if (v1) {
			var v3 := "r";
			var v4 := 'l';
			globalState.f9, globalState.f7, v3 := ([v1, v3 < (seq(abs(0xc4), i1  => ('f'))), v1, v1, v1 && fm0(v1, v1, v3, globalState)])[safeIndex(|(seq(abs(0x10b), i2  => (seq(abs(-0x343), i3  => (i0)))))|, |[v1, v3 < (seq(abs(0xc4), i1  => ('f'))), v1, v1, v1 && fm0(v1, v1, v3, globalState)]|) := v1 && v1], i0, ("quhelj" + v3[safeIndex(v0, |v3|) := v4]) + v3;
			var v5: array<seq<int>> := new seq<int>[27](i4 => [i0] + [i0]);
			var v6: array<bool> := new bool[26];
			var v7: array<array<bool>> := new array<bool>[9];
			v7[safeIndex(651, v7.Length)] := v6;
			var v8: map<bool, array<bool>> := map[v1 := v6];
			var v9 := DC11(v1, i0, true, v6, 0x13f);
			var v10: map<char, string> := map[v3[safeIndex(v0, |v3|)] := v3];
			var v11: set<int> := {fm1(false, v0, DC7(v0).cf6, globalState)};
			v5, v6, v7[safeIndex(651, v7.Length)], v1 := v5, if (!v1 in v8) then v8[!v1] else v6, v9.cf14, |v10| > safeModuloInt(v0, |v11|);
			v0 := i0;
			var v12: array<int> := new int[3];
			var v13: seq<int> := [v0];
			var v14: map<bool, bool> := map[v1 := v1];
			var v15: C0 := new C0(v1, v14);
			var v16: seq<int> := [|v13|, |map[v15 := v0]|];
			var v17: map<int, bool> := map[|v13| := v1];
			var v18: multiset<int> := multiset{i0, v0};
			var v19: seq<bool> := [v1, if ((if (-v0 in v18) then v18[-v0] else v0) in v17) then v17[if (-v0 in v18) then v18[-v0] else v0] else v15.f19, v15.f19];
			var v20: map<array<int>, seq<bool>> := map[v12 := v19];
			v20 := v20;
			v4 := v4;
		} else {
			globalState.f8 := v0;
			r1, v1 := v0 + v0, v1;
			var v21: array<int> := new int[22];
			v21 := v21;
			var v22 := "xbhg";
			var v23: C0 := new C0(v1, fm9(v0, |v22|, fm1(v1, i0, |map["q" := v0]|, globalState), globalState));
			var v24 := 'y';
			var v25: map<bool, int> := map[v1 := 770];
			v1, v23, r0, v22, v1 := fm0(v23.f19, v0 >= v0, v22[safeIndex(i0, |v22|) := v24], globalState), v23, -(i0 + |v25|), v22, if (v1) then v23.f19 else v1;
			v0 := i0;
		}
		
		var v26: array<bool> := new bool[2];
		var v27: seq<array<bool>> := [v26, v26, if (v1) then v26 else v26];
		v26 := v27[safeIndex(0x3d8, |v27|)];
	}
	r0 := v0;
	var v28 := 'f';
	var v29: seq<char> := [v28, v28, v28];
	var v30: seq<bool> := [v29 < (seq(abs(0x175), i5  => (v28))), v1];
	var v31: map<int, int> := map[v0 := |(seq(abs(0x2c5), i6  => (v0)))|];
	r1, v1 := (if (!v1) then 0xc else v0) * v0, v30[safeIndex(v0 * (if (|v29| in v31) then v31[|v29|] else fm1(v1, v0, v0, globalState)), |v30|)];
	var v32: array<string> := new string[3];
	var v33 := DC13(v29);
	v32[safeIndex(704, v32.Length)] := v33.cf17;
	v32[safeIndex(704, v32.Length)] := v29;
	var v34: array<bool> := new bool[27](i7 => v0 < v0);
	v34[safeIndex(798, v34.Length)] := 0x1e3 > v0;
	var v35: array<int> := new int[9];
	var v36: map<seq<bool>, array<int>> := map[v30 := v35];
	var v37 := DC0(v36 + v36);
	v1, r1, v34[safeIndex(798, v34.Length)], v37 := v1, v0, v1, v37.(cf0 := v36);
	if (v0 < v0) {
		v34[safeIndex(798, v34.Length)] := !v1;
		r0 := v0;
		globalState.f7 := v0;
		if (v1) {
			var v38: map<bool, bool> := map[v1 := !v1];
			var v39: seq<map<bool, bool>> := [v38, v38, v38, v38, v38[v34[safeIndex(798, v34.Length)] := v1]];
			var v40: seq<int> := [fm1(true, |v31|, v0, globalState), v0];
			var v41: array<map<bool, bool>> := new map<bool, bool>[21] [v38, v38, map[v1 := v34[safeIndex(798, v34.Length)]], v38, v38, v38, v38, v38, v38, map[v34[safeIndex(798, v34.Length)] := v34[safeIndex(798, v34.Length)]], v38, map[v1 := v34[safeIndex(798, v34.Length)]], v38, v39[safeIndex(|v40|, |v39|)], v38, map[v1 := v34[safeIndex(798, v34.Length)]], v38, v38, map[v34[safeIndex(798, v34.Length)] := v1], v38, v38];
			var v42: map<bool, array<map<bool, bool>>> := map[false := v41];
			var v43: seq<array<map<bool, bool>>> := [v41, v41, v41];
			var v44: set<int> := {v0, |v40|};
			var v45 := DC8(v1);
			var v46: array<array<map<bool, bool>>> := new array<map<bool, bool>>[27] [v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, v41, if (v34[safeIndex(798, v34.Length)] in v42) then v42[v34[safeIndex(798, v34.Length)]] else v41, v41, v41, v43[safeIndex(|v44|, |v43|)], if (v1) then v41 else v41, v41, v41, v41, v41, v41, v41, if (v45.cf7) then v41 else v41];
			v46[safeIndex(74, v46.Length)] := v41;
			v46[safeIndex(74, v46.Length)] := new map<bool, bool>[23](i8 => v38);
			var v47: C0 := new C0(v1, v38);
			v47 := v47;
			r1 := v0;
			v1 := (v0 - |v38|) != (-|v29[safeIndex(v0, |v29|) := v28]| - -v0);
			globalState.f8 := v0;
		} else {
			v0 := -v0;
			var v48: map<bool, bool> := map[v1 := v1];
			var v49: C0 := new C0(v1, v48);
			var v50: seq<array<int>> := [v35];
			var v51: multiset<array<int>> := multiset{v35, v50[safeIndex(|v29|, |v50|)], v35, v35};
			var v52: map<int, C0> := map[v0 - v0 := v49];
			var v53: array<D1> := new D1[5];
			var v54 := DC4({"qvldbtq", "aiws"});
			var v55 := DC5(v54);
			v53[safeIndex(147, v53.Length)] := v55;
			var v56: multiset<seq<bool>> := multiset{v30};
			v49, v51, globalState.f7, v52, v53[safeIndex(147, v53.Length)] := if (!v1) then v49 else v49, v51, -|(multiset([v30]) * v56)| + 0x4f, v52 + v52, v55;
			v2 := v2[v1 := v28];
			v34[safeIndex(798, v34.Length)] := v1 in v30;
			var v57: map<char, bool> := map[v28 := v49.f19];
			v57 := v57 + v57;
		}
		
		r0 := -v0;
	} else {
		v35[safeIndex(494, v35.Length)] := v0;
		v35[safeIndex(494, v35.Length)] := v0;
		var v58: set<seq<bool>> := {[v34[safeIndex(798, v34.Length)]], fm8(-0x7, v1, globalState)};
		v34[safeIndex(798, v34.Length)] := v58 > {v30};
		v0 := v0;
		var v59: set<int> := {v0, v35[safeIndex(494, v35.Length)]};
		var v60: multiset<int> := multiset{v35[safeIndex(494, v35.Length)], v35[safeIndex(494, v35.Length)]};
		var v61: map<multiset<int>, string> := map[DC14(v60).cf18 := v32[safeIndex(704, v32.Length)]];
		var v62 := DC10(v59, v61);
		match (v62) {
			case DC10(cf9, cf10) =>
				v35[safeIndex(494, v35.Length)] := v0;
				var v63: set<string> := {seq(abs(0x13e), i9  => (v28))};
				v1 := v63 < (v63 + v63);
				var v64: array<multiset<bool>> := new multiset<bool>[6];
				var v65 := DC3(v64);
				var v66: array<D1> := new D1[28] [v65, v65, v65, v65, v65, v65, v65, DC3(v64), v65, v65, v65, if (v1) then v65 else v65, v65, v65, v65, v65, DC3(DC3(v64).cf2), if (v1) then v65 else v65, v65, DC3(v64), v65, DC3(v64), v65, v65.(cf2 := v64), v65, v65, v65, v65];
				v66[safeIndex(691, v66.Length)] := v65;
				v66[safeIndex(691, v66.Length)] := v65;
				var v67 := DC14(v60);
				v67, v35[safeIndex(494, v35.Length)] := if (fm0(!!v1, v1, "i", globalState)) then v67 else v67, v35[safeIndex(494, v35.Length)];
			case DC11(cf11, cf12, cf13, cf14, cf15) =>
				var v68 := DC11(v30[safeIndex(cf12, |v30|)], v0, v1, cf14, -cf12);
				v68 := v68;
				v0 := if (cf11) then v0 else v0;
				cf11 := cf11;
				var v69: C0 := new C0(v35[safeIndex(494, v35.Length)] < v0, map[v34[safeIndex(798, v34.Length)] := v1]);
				v69 := v69;
			case DC9(cf8) =>
				var v70 := DC7(0x60);
				var v71: map<bool, int> := map[!v1 := v70.cf6];
				var v72: array<D5> := new D5[20];
				var v73: multiset<array<D5>> := multiset{v72};
				v71 := v71[v34[safeIndex(798, v34.Length)] := safeDivisionInt(|v73|, |v32[safeIndex(704, v32.Length)]|)];
				var v74 := DC11(cf8.f19, v35[safeIndex(494, v35.Length)], v34[safeIndex(798, v34.Length)], v34, v35[safeIndex(494, v35.Length)]);
				var v75: multiset<bool> := multiset{v34[safeIndex(798, v34.Length)]};
				var v76: seq<multiset<bool>> := [multiset([cf8.f19, v1]), v75, v75];
				var v77: array<int> := new int[11] [v0, v74.cf15 + |cf8.f20|, v0, |"ivkit"|, v35[safeIndex(494, v35.Length)], -v35[safeIndex(494, v35.Length)], -(v0 + v35[safeIndex(494, v35.Length)]), v0, |v76|, v0, v0];
				v77 := v77;
				var v78: set<string> := {"pbduko"};
				v78 := v78 + v78;
				var v79: array<seq<bool>> := new seq<bool>[26](i10 => v30);
				var v80: map<array<seq<bool>>, int> := map[v79 := v35[safeIndex(494, v35.Length)]];
				v80 := v80[v79 := |multiset{v35[safeIndex(494, v35.Length)]}|];
			case DC12(cf16) =>
				var v81: map<bool, bool> := map[v1 := !v34[safeIndex(798, v34.Length)]];
				v1, globalState.f8, v34[safeIndex(798, v34.Length)] := v34[safeIndex(798, v34.Length)], |v81|, true;
				var v82: array<C0> := new C0[4];
				v82 := v82;
				globalState.f8 := safeModuloInt(-0x6d, -v35[safeIndex(494, v35.Length)]);
				var v83 := new C0(v34[safeIndex(798, v34.Length)], v81[false := v34[safeIndex(798, v34.Length)]]);
		}
		
		var v84: set<bool> := {v1};
		var v85: map<bool, set<bool>> := map[v34[safeIndex(798, v34.Length)] := v84];
		var v86: seq<set<bool>> := [if (v34[safeIndex(798, v34.Length)] in v85) then v85[v34[safeIndex(798, v34.Length)]] else v84, v84];
		v86 := v86;
	}
	
	r0 := v0;
	r1 := safeModuloInt(v0, 0x22a);
}
method Main() {
	var v0 := 'u';
	var v1: array<int> := new int[13](i0 => safeDivisionInt(i0, |"y"|));
	var v2: map<char, array<int>> := map[v0 := v1];
	var v3 := false;
	var v4: set<bool> := {v3, v3};
	var v5 := 0xac;
	var v6: seq<bool> := [v3, v3, v3];
	var v7: map<bool, bool> := map[v3 := false];
	var globalState := new GlobalState(true, "pvfmfbioj", v2, v4, map[seq(abs(0x1ec), i1  => (v0)) := v5], -0x3d8, false, 0x94, -0x178, v6 + v6, false, v1, true, v7, 363, 0x2a8, true, false, false);
	var v8: multiset<bool> := multiset{v3};
	v3 := v3 !in (multiset{v3, v3, v3} * v8);
	var v9 := "qu";
	var i2 := 0;
	while (fm0(v3, v3, v9, globalState))
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		var v10: array<bool> := new bool[28] [v3, v3, v3, v3, v3, v3, !v3, fm0(v3, v3, seq(abs(0x22e), i3  => (v0)), globalState), v3, false, v3, false, false, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, fm0(v3, true, "w", globalState)];
		var v11: multiset<array<bool>> := multiset{v10, v10};
		v11 := (v11[v10 := abs(v5)] - multiset{v10}) - v11;
		v10[safeIndex(42, v10.Length)] := true;
		var v12: multiset<char> := multiset{v9[safeIndex(|(seq(abs(0x2e), i4  => (|v7|)))|, |v9|)]};
		var v13: map<int, bool> := map[|(v12 + v12)| := v3];
		v10[safeIndex(42, v10.Length)], v13 := v3, v13;
		var v14: multiset<seq<bool>> := multiset{v6};
		var v15: array<string> := new string[8];
		var v16: map<multiset<seq<bool>>, array<string>> := map[v14 := v15];
		v16 := v16[v14 := v15];
		v3 := v10[safeIndex(42, v10.Length)] && v10[safeIndex(42, v10.Length)];
	}
	var v17, v18 := m0(globalState);
	for i5 := safeDivisionInt(v5, v17) to v17 {
		v3 := true;
		var v19: map<array<int>, int> := map[v1 := |v6|];
		var v20: map<char, map<array<int>, int>> := map[v0 := v19];
		v20 := v20[v0 := v19];
		v7 := v7;
		var v21, v22 := m0(globalState);
	}
	globalState.f8 := v18;
	v3 := v3;
	var v23: array<array<bool>> := new array<bool>[27];
	var v24: array<bool> := new bool[8];
	v23[safeIndex(114, v23.Length)] := if (v3) then v24 else v24;
	v23[safeIndex(114, v23.Length)] := v24;
	var v25: array<string> := new seq<char>[15](i6 => (seq(abs(-0x232), i7  => (v0))) + "uueqdwt");
	v25[safeIndex(195, v25.Length)] := v9;
	v25[safeIndex(195, v25.Length)] := v9;
	globalState.f9 := v6;
	var i8 := 0;
	while (!((v17 == v18) <== v3))
		decreases 100 - i8
	{
		if (i8 >= 100) {
			break;
		}
		
		i8 := i8 + 1;
		v0 := 'h';
		var v26: multiset<set<bool>> := multiset{v4};
		v1[safeIndex(640, v1.Length)] := fm1(v3, |v26|, v18, globalState);
		var v27: multiset<char> := multiset{v0, v0};
		v1[safeIndex(640, v1.Length)] := fm1(v27 >= v27, -v17 * v18, v17, globalState);
		var v28: array<int> := new int[24](i9 => i9 * v18);
		var v29: map<bool, array<int>> := map[v3 := v28];
		var v30: map<seq<bool>, array<int>> := map[[v3] := v28];
		var v31 := DC0(map[v6 := v28]);
		var v32: array<map<seq<bool>, array<int>>> := new map<seq<bool>, array<int>>[22] [map[v6 := if (v3 in v29) then v29[v3] else v28] + v30, v30 + v30, if (v3) then map[v6 := v28] else v30, v30, v30, v30, v30, v30, v30, v30, v30, v31.cf0, v30, v30, v30 + v30, v30, v30, map[v6 := v28], v31.cf0 + v30, v31.cf0, v30[v6 := v28], v30];
		v32[safeIndex(845, v32.Length)] := v30;
		var v33: map<bool, map<seq<bool>, array<int>>> := map[v3 := v30];
		v32[safeIndex(845, v32.Length)] := if (v3) then v30 else if (v3 in v33) then v33[v3] else v30;
		globalState.f8 := v18;
	}
	var i10 := 0;
	while (v25[safeIndex(195, v25.Length)] <= v25[safeIndex(195, v25.Length)])
		decreases 100 - i10
	{
		if (i10 >= 100) {
			break;
		}
		
		i10 := i10 + 1;
		var v34 := new C0(false, v7);
		v1[safeIndex(51, v1.Length)] := v5;
		v1[safeIndex(51, v1.Length)] := -0x2ba;
		var v35: array<multiset<bool>> := new multiset<bool>[29] [v8, multiset{v34.f19, v3, v34.f19, if (v3 in v7) then v7[v3] else v3} - multiset{v34.f19, v3}, multiset{fm0(v3, v3, fm4(v34.f19, v3, !v3, |v4|, globalState), globalState), v34.f19}, multiset(v6), v8, v8[v3 := abs(v1[safeIndex(51, v1.Length)])], v8 - multiset{v3}, v8, v8[v34.f19 := abs(v18)], v8, v8, multiset([v34.f19]), v8, v8[false := abs(v1[safeIndex(51, v1.Length)])], v8, v8 - multiset{v34.f19, v34.f19}, v8, multiset{v34.f19, v34.f19, v34.f19}, v8, v8 + multiset{v3}, multiset{v3}, v8, v8[v3 := abs(v17)], v8, v8, multiset{v34.f19, v3, v3}, v8, v8[fm0(v3, v3, v25[safeIndex(195, v25.Length)], globalState) := abs(v1[safeIndex(51, v1.Length)])], v8];
		v35 := DC3(v35).cf2;
		v34.f19 := v34.f19;
	}
	match (DC1()) {
		case DC1() =>
			var v36: map<bool, set<string>> := map[fm0(!v3, v3, v25[safeIndex(195, v25.Length)], globalState) := fm5(v8, globalState)];
			var v37: set<string> := {"odbnb"};
			var v38 := DC4(if (!v3 in v36) then v36[!v3] else v37);
			var v39: seq<int> := [40];
			var v40: seq<seq<int>> := [v39, fm7(globalState)];
			v38 := fm6(v40, v3, v3, globalState).(cf3 := v37);
			v3 := v3 && v3;
			var v41: map<seq<int>, bool> := map[fm7(globalState) := v3];
			v41 := v41[v39 := v3 <==> v3];
			v24[safeIndex(386, v24.Length)] := v3;
			v24[safeIndex(386, v24.Length)] := |v25[safeIndex(195, v25.Length)]| >= |v4|;
		case DC0(cf0) =>
			globalState.f8 := 0x227;
			var v42: C0 := new C0(v3 <== v3, DC6(v7).cf5 + v7);
			v42 := v42;
			globalState.f8 := v5;
			v23[safeIndex(114, v23.Length)][safeIndex(386, v23[safeIndex(114, v23.Length)].Length)] := v3;
			v23[safeIndex(114, v23.Length)][safeIndex(386, v23[safeIndex(114, v23.Length)].Length)] := v18 < -0x27e;
		case DC2(cf1) =>
			var v43: array<C0> := new C0[21];
			v43 := v43;
			if (v3) {
				globalState.f7 := v5;
				v24[safeIndex(122, v24.Length)] := v3;
				var v44: map<bool, char> := map[v3 := 'g'];
				var v45: set<int> := {v5, v5, fm1(v3, v5, |v44|, globalState), safeDivisionInt(-0xb2, v17), v18};
				globalState.f8, globalState.f9, v24[safeIndex(122, v24.Length)], v45 := if (fm0(v3, v3, v9, globalState)) then v5 else v18, (if (v17 <= v18) then v6 else fm8(v18, v3, globalState))[safeIndex(v5, |(if (v17 <= v18) then v6 else fm8(v18, v3, globalState))|) := v3 ==> !false], if (v3) then !(v25[safeIndex(195, v25.Length)] <= v25[safeIndex(195, v25.Length)]) else -0x151 > |multiset{v23[safeIndex(114, v23.Length)], v24, v24}|, v45;
				var v46, v47 := m0(globalState);
				var v48: C0 := new C0(!true, v7);
				v48 := v48;
				v3 := v48.fm2(v0, fm0(v48.f19, v24[safeIndex(122, v24.Length)], v9, globalState), globalState) >= v5;
			} else {
				var v49, v50 := m0(globalState);
				v24[safeIndex(830, v24.Length)] := fm8(v17, v3, globalState) < fm8(v49, v3, globalState);
				v24[safeIndex(830, v24.Length)] := !v3 && v3;
				var v51: map<string, int> := map["mnvbm" := v50];
				globalState.f7 := if (v9 in v51) then v51[v9] else v50;
				var v52: map<int, int> := map[v49 := v49];
				var v53 := new C0(v3, fm9(if (DC7(|v25[safeIndex(195, v25.Length)]|).cf6 in v52) then v52[DC7(|v25[safeIndex(195, v25.Length)]|).cf6] else v49, v5, |v9|, globalState));
				var v54: map<array<int>, int> := map[v1 := v50 - v5];
				v54 := v54[v1 := -v49];
			}
			
			v0 := v0;
			v3 := !true;
	}
	
	var i11 := 0;
	while (v3)
		decreases 100 - i11
	{
		if (i11 >= 100) {
			break;
		}
		
		i11 := i11 + 1;
		v3, globalState.f8, globalState.f8 := v3 <== v3, v17, 0x2e8;
		if (fm0(v3, v3, "xahmwdx", globalState)) {
			var v55: map<int, char> := map[|v9| * v18 := 'p'];
			v55 := v55[v5 := v0];
			var v56 := new C0(v3, v7);
			var v57, v58 := m0(globalState);
			var v59 := DC8(!v3);
			v56.f19 := v59.cf7;
			var v60: seq<int> := [if (v56.f19 in v8) then v8[v56.f19] else v18, v5];
			var v61: map<C0, seq<int>> := map[v56 := v60];
			v61 := v61[v56 := v60];
		} else {
			var v62 := new C0(v3, v7 + v7);
			v7 := v7[fm0(v3, v3, v25[safeIndex(195, v25.Length)], globalState) := !v3];
			var v63: map<int, int> := map[safeDivisionInt(fm1(v3, v18, -v5, globalState), v17) := safeDivisionInt(v17, v18)];
			v63 := v63[v18 * v18 := |multiset{-v17, v62.fm2(v0, v62.f19, globalState), v5}|];
			v7 := v7[false := v3];
			var v64 := new C0(v62.f19 <==> v6[safeIndex(v18, |v6|)], v7);
		}
		
		v1[safeIndex(844, v1.Length)] := v18;
		globalState.f8, globalState.f7, v1[safeIndex(844, v1.Length)] := |v25[safeIndex(195, v25.Length)]|, safeDivisionInt(v17, v17), v5;
		var v65: multiset<string> := multiset{v25[safeIndex(195, v25.Length)][safeIndex(v17, |v25[safeIndex(195, v25.Length)]|) := v0], v9};
		var v66: seq<string> := [v25[safeIndex(195, v25.Length)][safeIndex(v5, |v25[safeIndex(195, v25.Length)]|) := v0]];
		v3 := (v65 + v65) > multiset(v66);
	}
	v3 := v3;
	if (v3) {
		if (false && v3) {
			var v67: C0 := new C0(v3, v7);
			v67 := v67;
			v1 := new int[15](i12 => i12 * v5);
			globalState.f8 := v18;
			v23[safeIndex(114, v23.Length)][safeIndex(41, v23[safeIndex(114, v23.Length)].Length)] := v5 != v17;
			var v68: set<int> := {v17, 0x97};
			var v70: map<set<int>, seq<bool>> := map[set v69 : int | (0xa8 <= v69) && (v69 < 0x6a) :: (v69 * v5) := v6];
			v23[safeIndex(114, v23.Length)][safeIndex(41, v23[safeIndex(114, v23.Length)].Length)] := v68 !in v70;
			v67 := v67;
		} else {
			var v71: C0 := new C0(if (v3 in v7) then v7[v3] else v3, v7);
			var v72 := DC9(v71);
			v71 := (v72.(cf8 := v71)).cf8;
			var v73: array<multiset<bool>> := new multiset<bool>[8] [v8, v8, v8, v8, v8, v8, v8, multiset{false}];
			v73 := new multiset<bool>[19];
			var v74: seq<string> := [v25[safeIndex(195, v25.Length)], v9];
			v9 := v25[safeIndex(195, v25.Length)] + (v74[safeIndex(0x21e, |v74|)] + v9);
			v71 := v71;
			v1[safeIndex(217, v1.Length)] := -0x342;
			var v75: set<int> := {v71.fm2(v0, !!v71.f19, globalState), v18};
			var v76: multiset<int> := multiset{|v8|};
			var v77: map<multiset<int>, int> := map[v76[|v9| := abs(v5)] := v18];
			var v78: seq<int> := [v17];
			v1[safeIndex(217, v1.Length)], v75, v0, v77, v3 := v78[safeIndex(-v18, |v78|)], if (v3) then v75 else v75, fm10(v78, fm11(v0, v71.f19, v17, v18, globalState), globalState), v77 + v77, v3;
		}
		
		var v79: seq<int> := [v17];
		var v80: seq<seq<int>> := [v79, seq(abs(-204), i13  => (|multiset{v17, 0x16a}|))];
		v4 := fm12(v80, v5 >= v5, if (!false in v7) then v7[!false] else v3, globalState);
		if ((if (v6[safeIndex(v18, |v6|)]) then v3 else v3) <== v3) {
			var v81: array<set<int>> := new set<int>[24];
			v81 := v81;
			v23 := v23;
			var v82: seq<array<bool>> := [v24, v24, v23[safeIndex(114, v23.Length)]];
			v82 := v82 + v82;
			v24[safeIndex(277, v24.Length)] := !true;
			v24[safeIndex(277, v24.Length)] := (if (v3 in v7) then v7[v3] else v3) && v3;
			globalState.f8 := v17;
		} else {
			var v83: C0 := new C0(v3, v7);
			v83 := v83;
			globalState.f7 := safeModuloInt(v5, v18);
			var v84: seq<C0> := [v83];
			v83 := v84[safeIndex(v5 - v5, |v84|)];
			globalState.f8 := |[v83.f19, !v3, v3, if (v3) then v83.f19 else v3]|;
			v24[safeIndex(721, v24.Length)] := v83.f19;
			var v85: map<int, bool> := map[0x1ba := v3];
			v24[safeIndex(721, v24.Length)], v85 := v83.f19, v85;
		}
		
		var v87: multiset<int> := multiset{-|v7|, 0x2b5, v17};
		var v88: set<multiset<int>> := {multiset{v17}, v87, v87, v87, v87};
		var v89 := DC10({v5, v5, v18}, map v86 : multiset<int> | v86 in v88 :: (v86) := (v9));
		var v90: seq<D5> := [if (v3) then v89 else v89, v89, v89];
		var v91: set<int> := {v5};
		var v92: map<multiset<int>, string> := map[v87 := v25[safeIndex(195, v25.Length)]];
		v1, v90, v0 := v1, ([v89, DC10(v91, v92)])[safeIndex(v17, |[v89, DC10(v91, v92)]|) := v89], v0;
		var v93 := DC13(v9);
		var v94 := new C0(fm0(v3, v3, v93.cf17, globalState), v7);
	} else {
		if ((v17 * 0x335) > v5) {
			v1[safeIndex(458, v1.Length)] := v18 - fm1(v3, v18, -v17, globalState);
			var v95: seq<array<int>> := [v1];
			var v96: map<int, int> := map[v18 := v18];
			globalState.f7, v1, v1[safeIndex(458, v1.Length)] := if (v3) then |v9| else -v17 - v17, v95[safeIndex(|{v3, v3}|, |v95|)], safeDivisionInt(-v18, safeModuloInt(if (v5 in v96) then v96[v5] else v17, v5));
			var v97: multiset<int> := multiset{v18, v1[safeIndex(458, v1.Length)]};
			var v98: map<multiset<int>, int> := map[v97 := |v25[safeIndex(195, v25.Length)]|];
			var v99: map<int, map<multiset<int>, int>> := map[v1[safeIndex(458, v1.Length)] := v98];
			var v100: seq<map<bool, bool>> := [v7, v7, v7];
			var v102: C0 := new C0(v3, v7);
			var v103: map<C0, bool> := map[v102 := v3];
			var v104: array<map<multiset<int>, int>> := new map<multiset<int>, int>[8] [if (v3) then v98[v97[v1[safeIndex(458, v1.Length)] := abs(v18)] := v1[safeIndex(458, v1.Length)]] else if (|v100[safeIndex(v18, |v100|)]| in v99) then v99[|v100[safeIndex(v18, |v100|)]|] else v98, (map v101 : multiset<int> | v101 in [v97] :: (v101) := (v5)) + map[v97 := v5], v98, (map[v97 := v5])[v97 := |v103|], map[multiset{v17, 274, v1[safeIndex(458, v1.Length)], v17} := v17], fm13(v3, -v5, globalState), v98, v98 + v98];
			v104[safeIndex(16, v104.Length)] := v98;
			globalState.f8, v18, v9, v104[safeIndex(16, v104.Length)] := -v1[safeIndex(458, v1.Length)], fm1(true, v17, v1[safeIndex(458, v1.Length)], globalState), (v25[safeIndex(195, v25.Length)] + (seq(abs(905), i14  => (v0)))) + (v9 + (seq(abs(934), i15  => (v0)))), v98;
			var v105: seq<array<bool>> := [v24, v24, v24, v24, v23[safeIndex(114, v23.Length)]];
			v105 := v105;
			var v106: seq<int> := [v1[safeIndex(458, v1.Length)], -v17];
			var v107: seq<seq<int>> := [v106, v106, v106, [v17, |[v1[safeIndex(458, v1.Length)]]|, v5]];
			var v108: set<string> := {"rjjbn"};
			v3, v106, v102.f19, v25[safeIndex(195, v25.Length)] := v3, v107[safeIndex(v5, |v107|)] + [|v108|], v3, "mwwcq" + v9;
			v102.f19 := false <==> v102.f19;
		} else {
			v3 := v3;
			var v109 := new C0(v3, v7);
			v1[safeIndex(32, v1.Length)] := v18 * v17;
			var v110: map<multiset<bool>, int> := map[v8 := v17];
			v1[safeIndex(32, v1.Length)] := -safeDivisionInt(|v110|, safeModuloInt(v5, v18));
			v109.f19 := v109.f19;
			globalState.f7 := -v1[safeIndex(32, v1.Length)] - v17;
		}
		
		var v111 := new C0(v6[safeIndex(v17, |v6|)], v7);
		v7 := v7[v3 := v111.f19];
		var v112: map<char, C0> := map[v0 := v111];
		v112 := if (v3) then (v112[v0 := v111])['w' := v111] else v112;
		v3 := v111.f19;
	}
	
	v7 := v7[v3 := v3];
	print v0, "\n";
	print v1[0], "\n";
	print v1[1], "\n";
	print v1[2], "\n";
	print v1[3], "\n";
	print v1[4], "\n";
	print v1[5], "\n";
	print v1[6], "\n";
	print v1[7], "\n";
	print v1[8], "\n";
	print v1[9], "\n";
	print v1[10], "\n";
	print v1[11], "\n";
	print v1[12], "\n";
	print |v2|, "\n";
	print v3, "\n";
	print v4 == {false}, "\n";
	print v5, "\n";
	print v6 == [false, false, false], "\n";
	print v7 == map[false := false], "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print |globalState.f2|, "\n";
	print globalState.f3 == {false}, "\n";
	print globalState.f4 == map[['u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u', 'u'] := 172], "\n";
	print globalState.f5, "\n";
	print globalState.f6, "\n";
	print globalState.f7, "\n";
	print globalState.f8, "\n";
	print globalState.f9 == [false, false, false], "\n";
	print globalState.f10, "\n";
	print globalState.f11[0], "\n";
	print globalState.f11[1], "\n";
	print globalState.f11[2], "\n";
	print globalState.f11[3], "\n";
	print globalState.f11[4], "\n";
	print globalState.f11[5], "\n";
	print globalState.f11[6], "\n";
	print globalState.f11[7], "\n";
	print globalState.f11[8], "\n";
	print globalState.f11[9], "\n";
	print globalState.f11[10], "\n";
	print globalState.f11[11], "\n";
	print globalState.f11[12], "\n";
	print globalState.f12, "\n";
	print globalState.f13 == map[false := false], "\n";
	print globalState.f14, "\n";
	print globalState.f15, "\n";
	print globalState.f16, "\n";
	print globalState.f17, "\n";
	print globalState.f18, "\n";
	print v8 == multiset{false}, "\n";
	print v9, "\n";
	print i2, "\n";
	print v17, "\n";
	print v18, "\n";
	print v23[6][2], "\n";
	print v24[2], "\n";
	print v25[0], "\n";
	print v25[1], "\n";
	print v25[2], "\n";
	print v25[3], "\n";
	print v25[4], "\n";
	print v25[5], "\n";
	print v25[6], "\n";
	print v25[7], "\n";
	print v25[8], "\n";
	print v25[9], "\n";
	print v25[10], "\n";
	print v25[11], "\n";
	print v25[12], "\n";
	print v25[13], "\n";
	print v25[14], "\n";
	print i8, "\n";
	print i10, "\n";
	print i11, "\n";
}