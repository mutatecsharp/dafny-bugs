datatype D0 = DC1(cf1: string, cf2: int, cf3: map<bool, int>, cf4: seq<int>) | DC2(cf5: map<bool, seq<int>>, cf6: bool) | DC0(cf0: seq<int>)
datatype D1 = DC4(cf8: bool, cf9: bool, cf10: string, cf11: int, cf12: int) | DC5(cf13: bool, cf14: bool, cf15: int, cf16: int) | DC6(cf17: bool, cf18: int, cf19: int, cf20: bool) | DC3(cf7: map<int, bool>) | DC7(cf21: D1)
datatype D2 = DC9(cf23: bool, cf24: int, cf25: int) | DC8(cf22: multiset<int>) | DC10(cf26: D2)
datatype D3 = DC12(cf28: bool, cf29: int) | DC11(cf27: array<int>)
datatype D4 = DC14 | DC13(cf30: map<char, int>)
datatype D5 = DC16(cf32: T0, cf33: bool, cf34: int) | DC15(cf31: multiset<bool>)
datatype D6 = DC18(cf36: bool) | DC19(cf37: int, cf38: C0, cf39: int) | DC17(cf35: seq<bool>)
datatype D7 = DC20(cf40: char)
datatype D8 = DC21(cf41: array<multiset<D3>>)
datatype D9 = DC23(cf43: int, cf44: int, cf45: bool, cf46: map<int, set<bool>>, cf47: int) | DC22(cf42: seq<C4>) | DC24(cf48: D9)
datatype D10 = DC26(cf50: bool, cf51: int, cf52: bool) | DC25(cf49: set<bool>)
class GlobalState {
	const f0 : int
	const f1 : map<bool, array<int>>
	const f2 : seq<set<int>>
	const f3 : seq<string>
	const f4 : bool
	constructor (f0 : int, f1 : map<bool, array<int>>, f2 : seq<set<int>>, f3 : seq<string>, f4 : bool) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
	}
	
}

function fm4(p0: int, globalState: GlobalState): map<int, bool> {
	map[-771 := !false] + map[0x20b := false]
}
function fm5(p0: map<bool, int>, p1: map<int, map<bool, int>>, globalState: GlobalState): string {
	"lvxthx"
}
function fm6(p0: int, p1: int, p2: bool, p3: int, globalState: GlobalState): int {
	|((multiset{0x2b5} * multiset{533, |[false]|, |(map v0 : char | v0 in {'j'} :: (v0) := (--|[|"vnemwohk"|, |[762]|]|))|}) + multiset{|DC4(!DC4(false, true, "rookixm", 0x114, |[false]|).cf9, true, "kowjt", 0x158, 0x1c4).cf10|, -0xf9, 0x125, DC4(false, true, seq(299, i0  => ('n')), 361, |(seq(-0x2e3, i1  => ('v')))|).cf12})|
}
function fm9(p0: int, globalState: GlobalState): D0 {
	DC1("khaefkx", |multiset{false, false, false, false, DC23(|map[-|DC25({false, false}).cf49| := false]|, -|multiset{'o'}|, true, map[934 := {true}], -0x34c).cf45}| % 0x70, map[!false := 0x75], [724] + [-729])
}
function fm10(p0: multiset<int>, p1: int, globalState: GlobalState): multiset<bool> {
	DC15(multiset{false, !!true, false, !!false}).cf31 * multiset([false] + [false, false])
}
function fm11(p0: char, p1: bool, p2: bool, globalState: GlobalState): char {
	'w'
}
function fm12(p0: int, p1: bool, p2: string, globalState: GlobalState): seq<int> {
	([|map[|[true]| := false]|] + [678]) + ([|"ucej"|] + [|{0x42}|])
}
function fm13(p0: string, p1: int, p2: multiset<int>, globalState: GlobalState): map<bool, int> {
	(map[!false := |{true}|] + map[!false := |"nmj"|]) + (map[true := -|map[true := 0x2ce]|] + map[true := -50])
}
function fm14(p0: bool, globalState: GlobalState): D1 {
	if (if (true) then false else true) then DC7(DC3(map[|map[true := |{true, true}|]| := false])) else DC7(DC5(true, false, |"gc"|, |"r"|))
}
function fm15(p0: int, globalState: GlobalState): map<bool, seq<int>> {
	map[-246 < 36 := [-678]]
}
function fm16(globalState: GlobalState): seq<bool> {
	[multiset{'m', 't'} !! multiset{'q'}]
}
function fm17(p0: bool, globalState: GlobalState): seq<seq<char>> {
	(seq(0x277, i0  => (['a']))) + ([['x']] + (seq(814, i1  => (['b']))))
}
function fm18(globalState: GlobalState): D6 {
	if (multiset{true} !! multiset{!!!!true, false, !false, false, false}) then DC17([true]) else DC17([!true])
}
function fm19(p0: int, globalState: GlobalState): set<multiset<int>> {
	{multiset(seq(0x312, i0  => (0xd0))), multiset{-0x246, 0x329}, multiset{|(seq(0x3d6, i1  => ('g')))|}} + ({multiset([|"seyfsdjs"|, 136]), multiset{|{|"pwmmx"|}|}, multiset{|{true, true}|}} * {multiset([|"dnlhtu"|])})
}
function fm20(p0: int, p1: multiset<bool>, p2: bool, p3: bool, globalState: GlobalState): D1 {
	match DC13(DC13(map v0 : char | v0 in ['c'] :: (v0) := (|"penvbw"|)).cf30) {
		case DC14() => DC6(true, 0x13, -0x25b, false)
		case DC13(cf30) => DC6(false, DC5(false, !true, |map[true := -945]|, 0x11c).cf15, 0x393, !!true)
	}
}
function fm21(globalState: GlobalState): set<int> {
	set v0 : int | (-0x165 <= v0) && (v0 < 743) :: (v0 % |[map[false := !false], map[true := true]]|)
}
function fm22(p0: int, p1: string, p2: int, p3: string, globalState: GlobalState): D5 {
	DC15(multiset{!!true} * multiset{false})
}
function fm23(p0: int, p1: bool, p2: int, p3: int, globalState: GlobalState): D2 {
	if (true) then DC8(multiset([-344])) else DC8(multiset{852})
}
function fm24(globalState: GlobalState): bool {
	!({458, -|map[|{true}| := false]|} > {-500}) <==> (|"bjjecxhfo"| > 53)
}
function fm25(globalState: GlobalState): map<char, int> {
	map['n' := 0x31c] + map['l' := |map[!false := |{true}|]|]
}
method m0(p0: int, p1: int, globalState: GlobalState) returns (r0: string, r1: array<int>, r2: int, r3: int) {
	var v0 := true;
	var v1: seq<bool> := [v0, v0, v0, v0, v0];
	for i0 := p1 to |v1| {
		var v3: set<int> := {-0x9c, p0};
		var v4: map<int, int> := map[i0 := 0x3d7];
		v0 := (map v2 : int | v2 in v3 :: (v2 / i0) := (p0)) != v4;
		var v5: array<int> := new int[25](i1 => i1 / p0);
		v5[500] := p0;
		var v6 := "riotpp";
		var v7 := DC4(v0, v0, v6, p1, p0);
		var v8: seq<seq<bool>> := [v1];
		var v9: C3 := new C3(|(v8[p0])[0x16c := v0]|, i0, v0);
		var v10: seq<C3> := [v9, v9, v9];
		var v11: multiset<int> := multiset{i0, p0, --|v10|, i0, |v6|};
		v5[500], r2, r3 := p1 - v7.cf12, |DC8(v11).cf22|, i0;
		r2 := (p1 * |v6|) + p0;
		var v12 := new C2(p1 / v5[500]);
	}
	r2 := p0;
	var v13 := 'q';
	v13 := match DC20(v13) {
		case DC20(cf40) => 'c'
	};
	var v14: array<int> := new int[4](i2 => i2 + 834);
	var v15: T0 := new C3(p0, p0, v0);
	var v16: multiset<T0> := multiset{v15};
	v14[835] := |v16|;
	v14[835] := p0;
	var v17 := DC13(fm25(globalState));
	v0 := match v17 {
		case DC14() => v15.f6
		case DC13(cf30) => v15.f6 <== v0
	};
	var v18: array<map<int, int>> := new map<int, int>[4];
	var v19: multiset<bool> := multiset{v15.f6, v15.f6};
	var v20: C0 := new C0(867, if (v0 in v19) then v19[v0] else v14[835]);
	var v21: map<C0, bool> := map[v20 := v15.f6];
	var v22: map<array<map<int, int>>, bool> := map[if (v0) then v18 else v18 := |v21[v20 := v15.f6]| <= 831];
	v22 := v22[v18 := v0];
	var v23 := "tjkieshnd";
	r0 := v23;
	var v24: map<bool, bool> := map[v15.f6 := v0];
	var v25: seq<int> := [p0, |v19|, |[v23, v23, v23, "eao", "iolhb"]|, |(v24[v15.f6 := v0])[v0 := v0]|];
	r1 := new int[10] [v14[835], -p1, 0x70, p0, p1 * fm6(v20.f12, v15.f5, v15.f6, v20.f13, globalState), p1, v25[v20.f12], v20.f12, v14[835] + v20.f13, v20.f12];
	r2 := v14[835];
	r3 := |[v14]|;
}
trait T0 {
	const f5 : int
	const f6 : bool
	function fm0(p0: bool, globalState: GlobalState): bool 
	method m1(p0: int, globalState: GlobalState) 
}

class C0 {
	const f12 : int
	const f13 : int
	constructor (f12 : int, f13 : int) {
		this.f12 := f12;
		this.f13 := f13;
	}
	
	function fm7(p0: seq<int>, p1: map<bool, bool>, p2: int, p3: multiset<int>, globalState: GlobalState): map<int, bool> {
		(map v0 : int | (0x39d <= v0) && (v0 < 0x18a) :: (v0 * f12) := (true)) + (map[f12 := true] + DC3(map v1 : int | (0x18d <= v1) && (v1 < 0x74) :: (v1 + f13) := (false)).cf7)
	}
	function fm8(p0: bool, p1: string, p2: char, globalState: GlobalState): bool {
		!true || false
	}
}

class C1 extends T0 {
	var f10 : int
	const f11 : int
	constructor (f10 : int, f11 : int, f5 : int, f6 : bool) {
		this.f10 := f10;
		this.f11 := f11;
		this.f5 := f5;
		this.f6 := f6;
	}
	
	function fm0(p0: bool, globalState: GlobalState): bool {
		f6
	}
	function fm2(p0: set<bool>, p1: int, p2: int, globalState: GlobalState): bool {
		f11 == -(f5 + f11)
	}
	function fm3(p0: seq<int>, p1: bool, p2: int, globalState: GlobalState): map<bool, int> {
		map[f6 := f10]
	}
	method m1(p0: int, globalState: GlobalState) {
		var v0 := "ju";
		var v1: multiset<string> := multiset{if (false) then v0 else "kcobgiw", seq(0x3b7, i0  => ('r')), v0 + (seq(0, i1  => ('f')))};
		f10 := if (v0 in v1) then v1[v0] else p0;
		var v2: array<bool> := new bool[15];
		var v3: map<bool, int> := map[f6 := p0];
		var v4: map<int, int> := map[|v3| := f10];
		var v5: set<bool> := {true, f6};
		var v6: seq<int> := [|v5|];
		var v7: map<bool, seq<int>> := map[f6 := v6];
		var v8 := DC2(v7, !f6);
		var v9: set<D0> := {v8};
		for i2 := f5 / |fm4(|[v2]|, globalState)| to if (|v0| in v4) then v4[|v0|] else |v9| {
			var v10: seq<map<bool, seq<int>>> := [map[f6 := v6]];
			var v11: map<int, bool> := map[f5 := false];
			v3 := v3[DC2(v10[p0], f6).cf6 := |v11|];
			var v12: map<int, string> := map[i2 := v0];
			var v13 := 'b';
			var v14 := DC1(v0[f10 := v13], f5, v3, v6);
			var v15: seq<string> := [v0, v0];
			var v16: map<int, map<bool, int>> := map[0x28c := v3];
			var v17: array<string> := new string[24] [v0 + "md", if (p0 in v12) then v12[p0] else seq(0x46, i3  => ('l')), v0[p0 := v13] + v0, v14.cf1[f11 := v13], v0, "o", "th", v15[f10], v0, v0, v0, v0, fm5(v3, v16, globalState), if (f11 in v12) then v12[f11] else v0, v0, "chm" + v0, v0, v0, v0[f11 := v13], "u" + v0, DC1("ylxgswe", fm6(f5, f5, f6, f5, globalState), v3, v6).cf1, v0 + v0, v0 + "soi", v0];
			v17[567] := v0;
			var v18 := false;
			var v19: seq<bool> := [fm0(fm0(true, globalState), globalState), v18];
			var v20: seq<seq<bool>> := [v19];
			f10, f10, v17[567], v18 := i2 * (f10 * p0), p0, v0, f10 < |(v20 + v20)|;
			f10 := i2;
			f10 := -(f11 - p0);
		}
		var v22: map<string, int> := map[fm5(v3, map v21 : int | (0x6f <= v21) && (v21 < 211) :: (v21 * f11) := (v3), globalState) := f11];
		f10 := |v22|;
		for i4 := f5 to f10 {
			var v23 := new C0(fm6(-|v0|, i4, f6, f5, globalState), f11);
			if (if (f6) then f6 else f6 <==> f6) {
				var v24, v25, v26, v27 := m0(f10, v6[v23.f13], globalState);
				var v28: multiset<bool> := multiset{f6};
				var v29: map<multiset<bool>, int> := map[v28 := -v23.f12];
				var v30: multiset<int> := multiset{v23.f12, p0, |v24|, f10, 0x264};
				var v31: map<int, multiset<bool>> := map[f11 := v28];
				var v32 := DC8(v30);
				var v33: map<array<bool>, int> := map[v2 := -|"l"|];
				var v34: array<multiset<int>> := new multiset<int>[15] [v30, v30 - v30, multiset(v6), v30, v30, v30, v30, v30 + v30, v30, v30 - v30, multiset((seq(0xb4, i5  => (v23.f13))) + (fm9(v23.f13, globalState)).cf4), multiset(v6), v30, multiset(v6 + [|(if (i4 in v31) then v31[i4] else v28)|, v23.f13, |v0|]), v32.cf22[if (v2 in v33) then v33[v2] else |v3| := |(seq(0x1c9, i6  => ('u')))|]];
				v34[573] := v30;
				var v35 := true;
				v2[837] := v35;
				var v36: seq<bool> := [true, f6, false];
				var v37 := DC1(v24, v27, v3, [-v23.f13]);
				var v38 := 'd';
				v29, v34[573], v35, v2[837] := v29, v30, (v36 + [f6, !f6]) <= [v35], v23.fm8(v35, v37.cf1, v38, globalState);
				v24 := v24;
				v34[573] := v30;
				v27, v2[837], v25, f10, v35 := f5 % f11, v24 < ((seq(0x6e, i7  => (v38))) + (seq(125, i8  => ('k')))[v23.f12 := v38]), v25, v27, v35;
			} else {
				var v39 := 'w';
				v39 := v39;
				f10 := if (!f6) then |v0| else v23.f12;
				var v40: array<multiset<bool>> := new multiset<bool>[15];
				var v41: multiset<bool> := multiset{f6};
				var v42: seq<bool> := [f6, f6];
				var v43: multiset<int> := multiset{|v0|, p0, f10, |v42|};
				v40[794] := v41 - fm10(v43, v23.f13, globalState);
				v40[794] := v41;
				v39 := v39;
				var v44 := true;
				v44 := |(v0 + "qchhltun")| < -(i4 + f10);
			}
			
			if (f6) {
				var v45 := new C0(0x191, if (f6) then i4 else -0x2);
				f10 := -f11;
				var v46: array<int> := new int[1];
				var v47: set<array<int>> := {v46, v46, v46};
				v46[131] := f5;
				var v48: map<int, string> := map[v45.f12 := v0];
				f10, v47, v0, v46[131] := (0x1de % v45.f12) + |(v48 + v48)|, v47, "ihvp", v45.f13 % v23.f12;
				v2 := v2;
				var v49: array<string> := new string[24];
				var v50: seq<bool> := [f6];
				v49, v50 := if (v46[131] >= 0x73) then v49 else v49, (v50 + v50) + v50;
			} else {
				var v51: array<int> := new int[3](i9 => i9 / v23.f12);
				v51[153] := p0;
				var v52: seq<bool> := [!f6];
				var v53 := 'e';
				var v54: map<int, bool> := map[|v4| := f6];
				var v55: map<int, map<bool, int>> := map[v23.f13 := v3];
				v51[153] := |((v0 + v0)[|v52| := fm11(v53, true, f6, globalState)] + (fm5(v3[f6 := |v54|], v55, globalState) + v0))|;
				f10 := v23.f12;
				v2[22] := f6 ==> f6;
				var v56 := true;
				var v57: array<array<int>> := new array<int>[25];
				v57[211] := v51;
				var v58: multiset<bool> := multiset{v56, f6};
				var v59: multiset<int> := multiset{v23.f12};
				v2[22], v56, v57[211] := f6 <==> (v58 > fm10(v59, 100, globalState)), v56, v51;
				var v60: map<bool, bool> := map[v56 := false];
				v53 := if (!(|v60| <= 294)) then v53 else v53;
				v56 := !(v52 < [v52[v23.f13]]);
			}
			
			var v61 := new C0(0x3b6 % f5, p0);
		}
		var v62: map<bool, bool> := map[f6 := f6];
		f10 := |v62| * |(v6[p0 := f10] + (seq(0x65, i10  => (0x2c7))))|;
		var v63 := false;
		v63 := !f6;
	}
	method m4(p0: map<int, map<int, bool>>, p1: seq<map<bool, int>>, p2: int, p3: int, globalState: GlobalState) {
		var v0: seq<bool> := [true, f6, fm0(f6, globalState)];
		var v1: map<bool, seq<bool>> := map[!f6 := v0];
		v1 := v1[!false := v0];
		for i0 := p3 to f5 {
			var v2: multiset<int> := multiset{p2};
			var v3: seq<int> := [|v2|, fm6(p3, f5, f6, -0x303, globalState)];
			var v4: multiset<int> := multiset{f10, v3[p2], 0x3a1};
			var v5: seq<seq<int>> := [v3];
			var v6 := new C0(f5, if (|v5| in v4) then v4[|v5|] else f10);
			var v7 := DC2(map[f6 := v3], f6);
			var v8: map<bool, seq<int>> := map[f6 := [p3]];
			var v9: multiset<D0> := multiset{v7, if (f6) then v7 else DC2(v8, false).(cf6 := f6)};
			v9 := v9;
			var v10 := false;
			v10 := v10;
			var v11 := DC0(seq(0x1c7, i1  => (p3)));
			match (v11) {
				case DC1(cf1, cf2, cf3, cf4) =>
					var v12: array<int> := new int[22];
					v12 := v12;
					v10, v3, v10 := f6, cf4, v10;
					var v13: multiset<bool> := multiset{v10};
					var v14 := new C0(fm6(f10, if (false in v13) then v13[false] else p2, v10, v6.f13, globalState), v6.f12);
					v12[446] := f5;
					v12[446] := f11 * v6.f13;
				case DC2(cf5, cf6) =>
					var v15: set<bool> := {cf6};
					var v16: set<int> := {f10, p3, v6.f12, |v15|, i0};
					var v17: array<int> := new int[11] [|v16|, p3, |v3|, f11, v6.f12, v6.f13, f10 * v6.f12, |"wecoqytg"|, f11, v6.f12, f10];
					v17 := v17;
					var v18: array<bool> := new bool[21](i2 => v10);
					var v19: array<array<bool>> := new array<bool>[29] [v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18, v18];
					v19[898] := v18;
					v19[898] := v18;
					var v20: seq<array<bool>> := [v19[898]];
					var v21: map<int, seq<array<bool>>> := map[|[v6.f12]| := v20[f11 := v18]];
					var v22, v23, v24, v25 := m0(f10 - |(if (v6.f13 in v21) then v21[v6.f13] else v20)|, 0x1a9, globalState);
					v10 := fm2({f6, v10, cf6, cf6, v10}, -v24, |v22|, globalState);
				case DC0(cf0) =>
					var v26: map<C0, bool> := map[v6 := v10];
					var v27: set<map<C0, bool>> := {v26, v26};
					var v28 := 'n';
					var v29 := "xcfy";
					var v30 := DC6(f6, |v27|, v6.f12, v28 !in v29[v6.f13 := v28]);
					v30 := v30;
					v10 := v10;
					var v31: multiset<bool> := multiset{false};
					var v32: map<bool, bool> := map[v10 <== f6 := v31 > v31];
					v32 := v32[v10 := v10];
					f10 := p2;
			}
			
		}
		var v33: seq<int> := [p2, f10, f10];
		var v34: set<int> := {f11};
		var v35: map<bool, bool> := map[f6 := f6];
		var v36: array<int> := new int[5](i3 => i3 % p3);
		var v37: map<int, array<int>> := map[f10 := v36];
		var v38 := DC11(v36);
		v33 := fm12(|v34| * f10, !(|map[|v35| := if (f10 in v37) then v37[f10] else v38.cf27]| == 431), seq(-214, i4  => ('v')), globalState);
		var v39: map<bool, int> := map[f6 := |"vbenawu"|];
		var v40: multiset<int> := multiset{p2, f5};
		var v41: map<int, multiset<int>> := map[|v39| := v40 * v40];
		var v42: map<bool, seq<int>> := map[f6 := v33];
		var v43 := DC2(v42, f6);
		var v44: set<bool> := {!true};
		v41 := v41[if (v43.cf6) then f11 else fm6(|v44|, p3, false, 0x189, globalState) := v40[f5 := 515]];
		var v45 := 'i';
		var v46: map<char, int> := map[v45 := p3];
		var v47 := DC13(v46);
		var v48 := new C0(f11, if (!false) then |(v47.(cf30 := v46)).cf30| else p2);
		forall i5 | 0 <= i5 < v36.Length {
			v36[i5] := i5 + (f11 % v48.f12);
		}
	}
}

class C2 {
	var f9 : int
	constructor (f9 : int) {
		this.f9 := f9;
	}
	
	function fm1(p0: int, p1: char, globalState: GlobalState): bool {
		-|(map['r' := true] + map['t' := true])| < f9
	}
	method m3(p0: string, globalState: GlobalState) returns (r0: map<map<char, D0>, int>) {
		f9 := f9;
		var v0 := false;
		var i0 := 0;
		while (v0 || v0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v1: T0 := new C1(f9, |[v0, v0, v0, v0]|, f9, v0);
			var v2: set<T0> := {if (v0) then v1 else v1};
			var v3: map<string, set<T0>> := map[p0 := {v1}];
			var v4: seq<set<T0>> := [v2];
			v2 := (if (p0 in v3) then v3[p0] else v4[249]) - (v2 - v2);
			var v5: map<int, string> := map[v1.f5 / 0x265 := seq(193, i1  => ('d'))];
			v5 := v5[v1.f5 := if (v1.f6) then "byj" else p0];
			var v6: array<bool> := new bool[2];
			v6[353] := v0;
			v6[353] := !!(v1.f5 == fm6(v1.f5, f9, v0, v1.f5, globalState));
			var v7: map<bool, int> := map[v0 := f9];
			var v8: map<bool, map<bool, int>> := map[v6[353] := v7];
			var v9: seq<int> := [v1.f5];
			var v10 := DC1(p0, f9, v7, v9);
			var v11: map<int, map<bool, int>> := map[f9 := v7];
			var v12: multiset<int> := multiset{v1.f5};
			var v13 := DC6(v0, |v12|, v9[v1.f5], v1.f6);
			var v14: seq<bool> := [v6[353], v6[353], v6[353], v1.f6, v13.cf20];
			var v15: multiset<int> := multiset{|v14|, v1.f5, f9, v1.f5, f9};
			var v16 := DC8(v15);
			var v17: array<map<bool, int>> := new map<bool, int>[19] [if (v1.f6 in v8) then v8[v1.f6] else map[v6[353] := v1.f5], v7, v7, v10.cf3 + v7, v7, v7, v7, v7, v7, v7, v7, v7, v7 + (if (v1.f5 in v11) then v11[v1.f5] else v7), v7 + v7, v7, v7[v1.f6 := v1.f5], v7, fm13(p0, if (-v1.f5 in v15) then v15[-v1.f5] else 0x90, v16.cf22, globalState), v7];
			v17 := v17;
		}
		var v18: map<int, int> := map[f9 := f9];
		var v19: seq<int> := [-0x219, fm6(|v18|, 0xd2, v0, f9, globalState)];
		var i2 := 0;
		while (v19 == (if (v0) then v19 else [f9]))
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			var v20: seq<bool> := [v0];
			match (fm14(v0 !in v20, globalState)) {
				case DC4(cf8, cf9, cf10, cf11, cf12) =>
					cf9, cf9 := true, cf9;
					var v21: multiset<bool> := multiset{v0, cf9, false, !(cf8 <== cf9)};
					v21 := v21;
					cf9 := cf8;
					var v22: map<bool, seq<int>> := map[cf8 := fm12(cf11, v0, cf10, globalState)];
					cf9 := v22 != (fm15(cf11, globalState) + v22);
				case DC5(cf13, cf14, cf15, cf16) =>
					cf13 := false;
					v20 := fm16(globalState);
					var v23: multiset<int> := multiset{575};
					var v24: multiset<bool> := multiset{v0, false, cf13, v0};
					var v25 := DC8(v23 + multiset{f9, if (v0 in v24) then v24[v0] else cf16, f9});
					v25 := v25.(cf22 := v23).(cf22 := multiset{cf16, cf15, cf16});
					var v26 := 'c';
					v26 := v26;
				case DC6(cf17, cf18, cf19, cf20) =>
					v0 := cf20;
					var v27, v28, v29, v30 := m0(cf18, fm6(f9, f9, cf20, f9, globalState), globalState);
					var v31: map<bool, bool> := map[cf17 := v0];
					var v32: map<int, map<bool, bool>> := map[0x287 := v31];
					var v33: C1 := new C1(v30 * |(if (0x3d8 in v32) then v32[0x3d8] else v31)|, 0x16d, v29, cf20);
					v33 := v33;
					var v34 := DC9(v0, v33.f11, cf19);
					var v35: set<bool> := {v0};
					var v36 := DC6(cf17, |v35|, |v27|, v0);
					var v37: map<bool, seq<int>> := map[v36.cf17 := v19];
					var v38 := DC2(v37, cf20);
					var v39: multiset<bool> := multiset{cf20};
					var v40 := 'u';
					var v41: array<bool> := new bool[28] [!v0 <==> cf20, v0, DC9(v34.cf23, 0x329, |p0|).cf23, !cf20, v0, v0, cf17 <== v0, cf17, !!(v0 ==> cf20), cf20, cf17, v0 !in v31, cf17, v34.cf23, cf20, v0, cf17, v0, v0 <== v38.cf6, cf20, multiset{!cf20} !! v39, cf20, v19 == v19, v33.fm2(v35, v33.f11, cf18, globalState), cf20, fm1(v33.f10, v40, globalState), v20[v30], v0];
					v41[564] := cf17;
					var v42: map<bool, int> := map[v0 := |v27|];
					var v43: seq<map<bool, int>> := [v42];
					var v44 := DC1(p0, v29, v43[v33.f10], seq(0x22d, i3  => (v30)));
					v41[564] := v40 in v44.cf1;
				case DC3(cf7) =>
					var v45: array<char> := new char[19](i4 => if (true) then 'w' else 'l');
					var v46 := 'f';
					v45[407] := v46;
					v45[407] := v46;
					var v47: array<bool> := new bool[22];
					v47[588] := !v0;
					v47[588] := v0;
					var v48: multiset<int> := multiset{-f9, f9, f9, f9};
					var v49 := DC8(v48);
					v49 := v49.(cf22 := v48);
					var v50: array<array<bool>> := new array<bool>[5];
					v50[834] := v47;
					v50[834] := new bool[18];
				case DC7(cf21) =>
					var v51: map<int, bool> := map[0x1f4 := v0];
					var v52 := DC3(v51);
					var v53 := DC7(v52);
					v53 := v53;
					var v54, v55, v56, v57 := m0(f9, f9, globalState);
					v55[163] := |(map v58 : int | (809 <= v58) && (v58 < 0xba) :: (v58 % f9) := (v54))|;
					v55[163] := v57;
					var v59: set<bool> := {v0, v0};
					var v60 := DC5(v0, v0, |v59|, 0x5a);
					var v61: map<D1, seq<int>> := map[v60 := v19];
					v61 := v61[v60 := [|v54|] + v19];
			}
			
			var v62: array<bool> := new bool[1];
			v62[252] := v20[-|p0|];
			v62[252] := v0;
			v62[252] := v62[252] || v62[252];
			v62[252] := f9 == f9;
		}
		v0 := v0 <==> v0;
		var v63 := DC14();
		v63 := v63;
		var v64: map<int, bool> := map[f9 := v0];
		v64 := v64[f9 := v0];
		var v65 := 'h';
		var v66 := DC0([-0x398]);
		var v67: map<char, D0> := map[v65 := v66];
		var v68: set<bool> := {v0, v0};
		var v69: map<map<char, D0>, int> := map[v67 := |v68|];
		var v71: map<bool, map<map<char, D0>, int>> := map[!v0 := v69];
		var v74: multiset<char> := multiset{v65};
		var v75: set<map<char, D0>> := {map v73 : char | v73 in v74 :: (v73) := (v66), map[v65 := v66.(cf0 := v19)]};
		var v76: map<int, set<map<char, D0>>> := map[f9 := v75];
		var v77: seq<map<map<char, D0>, int>> := [v69, map v70 : map<char, D0> | v70 in (if (v0 in v71) then v71[v0] else map v72 : map<char, D0> | v72 in (if (f9 in v76) then v76[f9] else v75) :: (v72) := (f9)) :: (v70) := (f9), v69];
		r0 := v77[f9];
	}
}

class C3 extends T0 {
	const f8 : int
	constructor (f8 : int, f5 : int, f6 : bool) {
		this.f8 := f8;
		this.f5 := f5;
		this.f6 := f6;
	}
	
	function fm0(p0: bool, globalState: GlobalState): bool {
		!(f8 >= f5)
	}
	method m1(p0: int, globalState: GlobalState) {
		var i0 := 0;
		while (f6)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v0 := false;
			v0 := !v0;
			var v1: multiset<bool> := multiset{f6, v0};
			var v2 := DC15(v1);
			var v3: seq<int> := [|multiset{f6, f6}|, f8, |v2.cf31|, p0];
			var v4 := new C2(v3[f5] + p0);
			var v5: set<int> := {p0, f5, p0, f5};
			var v6: map<int, int> := map[v4.f9 := 408];
			var v7: map<bool, bool> := map[v0 := v0];
			var v8: array<int> := new int[15] [fm6(0x38, v4.f9, v0, p0, globalState), f8, v4.f9, v4.f9, v4.f9, f8 % f5, -v4.f9, v4.f9 * |v5|, fm6(v4.f9, f8, v0, f5, globalState), f8, if (f5 in v6) then v6[f5] else v3[-p0], v4.f9, v4.f9 % p0, |((seq(0x379, i1  => (v4.f9))) + v3)|, if (v0) then -|v7| else p0];
			v8[279] := v4.f9;
			var v9: array<array<bool>> := new array<bool>[20];
			var v10: array<bool> := new bool[29];
			v9[888] := v10;
			v8[279], v9[888], v4.f9 := --f5, v10, f5;
			v9[888][338] := f6;
			var v11: T0 := new C1(v8[279], f5, v4.f9, f6);
			var v12 := 'b';
			var v13: seq<char> := [v12];
			var v14 := DC16(v11, false, |v13|);
			v0, v9[888][338] := (v14.(cf33 := v0)).cf33, f6;
		}
		if (p0 <= p0) {
			var v15 := 'o';
			var v16: map<int, char> := map[p0 := v15];
			var v17: multiset<bool> := multiset{f6, f6, f6, fm0(f6, globalState), f6};
			var v18: C0 := new C0(|v17|, f8);
			var v19: map<C0, char> := map[v18 := v15];
			var v20: multiset<map<int, char>> := multiset{v16, v16[|v19| := fm11(v15, f6, f6, globalState)]};
			var v21: seq<multiset<map<int, char>>> := [v20];
			var v22: multiset<int> := multiset{v18.f13};
			var v23: multiset<int> := multiset{|v22|, fm6(v18.f12, f5, f6, v18.f12, globalState), fm6(v18.f13, f8, f6, v18.f13, globalState), f5, f8};
			v20 := if (f6) then v21[|v23|] else v20;
			var v24: seq<char> := [v15];
			var v25: seq<seq<char>> := [v24];
			var v26: map<int, bool> := map[f8 := !f6];
			var v27: seq<seq<seq<char>>> := [fm17(if (p0 in v26) then v26[p0] else f6, globalState)];
			v25 := v27[|((map v28 : int | (0x77 <= v28) && (v28 < 0x73) :: (v28 / v18.f13) := (!f6)) + (map v29 : int | (223 <= v29) && (v29 < 0x1b9) :: (v29 + v18.f13) := (true)))|];
			var v30: array<C2> := new C2[29];
			var v31: set<char> := {v15, v15, v15, v15};
			var v32: C2 := new C2(|v31|);
			v30[190] := v32;
			v30[190] := v32;
			var v33 := true;
			v33 := f6;
			v15 := 'r';
		} else {
			var v34: map<bool, bool> := map[!f6 := f6];
			var v35: seq<map<bool, bool>> := [v34, v34];
			var v36: map<bool, int> := map[f6 := |[f6, f6]|];
			var v37 := new C1(|((seq(0x112, i2  => (map[f6 := f6]))) + v35)|, f5, |v36|, if (f6) then f6 else f6);
			var v38: array<seq<bool>> := new seq<bool>[10];
			var v39: seq<bool> := [true];
			var v40: map<int, seq<bool>> := map[f8 := v39];
			var v41 := DC17(v39);
			v38 := new seq<bool>[28] [(v39 + v39)[|(if (v37.f10 in v40) then v40[v37.f10] else v39)| := f6], v39, v39 + v39, DC17(v39).cf35, [f6, f6], fm16(globalState), fm16(globalState), v39 + v39, v39, v39, v39 + v39, (fm18(globalState)).cf35, [true], v39, v39, v39 + v41.cf35, v39 + v39, v39, [f6], [f6], [f6, f6], v39, (v39 + v39)[v37.f10 := false], v39, v39 + v39, v39, [true], v39];
			var v42: array<bool> := new bool[18](i3 => f6);
			var v43: seq<int> := [f5];
			var v44: seq<int> := [|v43|];
			v42[947] := v44 <= v43;
			var v45: array<int> := new int[11];
			var v46 := DC11(v45);
			v42[947], v46 := false, v46;
			var v47 := 'n';
			v47 := v47;
			var v48: multiset<int> := multiset{0x26e};
			var v49: seq<multiset<int>> := [v48];
			var v50 := "fmwqpjm";
			var v51: seq<string> := [v50];
			var v52: map<multiset<int>, string> := map[v49[f5] := v51[v37.f11]];
			v42[947] := if (v37.fm0(f6, globalState)) then !(v37.f10 > |v52|) else v48 > v48;
		}
		
		var v53 := "whxnlgjoh";
		for i4 := f5 - fm6(p0, f5, f6, p0, globalState) to -|v53| {
			var v54 := 0x160;
			var v55: seq<bool> := [false];
			var v56: set<bool> := {f6};
			var v57: seq<int> := [|v55|, |v53|, |v56|];
			v54 := |(v57 + v57)| / v54;
			var v58 := true;
			var v59: multiset<int> := multiset{|v53|};
			var v60: map<multiset<int>, bool> := map[v59 := f6];
			v58 := if (f6) then !(|v60| >= 0x39c) else v58 <==> f6;
			var v61 := 't';
			var v62: array<int> := new int[15];
			var v63 := DC18(f6);
			v58, v54, v61, v62, v54 := fm0(v63.cf36 <==> f6, globalState), p0, v61, v62, p0;
			v58 := f6;
		}
		var v64 := 'o';
		var v65: array<char> := new char[4] [v64, v64, v64, v64];
		var v66: seq<array<char>> := [v65];
		var v67: multiset<bool> := multiset{f6, true};
		var v68 := DC15(v67);
		var v69: array<seq<bool>> := new seq<bool>[7](i5 => [f6, f6] + [f6, f6]);
		var v70: seq<bool> := [f6, f6];
		v69[417] := v70 + v70;
		v69[680] := v70;
		var v71 := false;
		var v72 := DC7(DC4(v71, v71, seq(0x1fc, i6  => (v64)), p0, fm6(f8, -0x3e0, v71, 0x396, globalState)));
		var v73: map<D1, multiset<bool>> := map[v72 := v67];
		v66, v68, v69[417], v69[680], v71 := v66, DC15(v67).(cf31 := if (DC7(fm14(v71, globalState)) in v73) then v73[DC7(fm14(v71, globalState))] else v67), v70 + v70, fm16(globalState), v71;
		var v74: array<int> := new int[25](i8 => i8 - p0);
		forall i7 | 0 <= i7 < v74.Length {
			v74[i7] := i7 / p0;
		}
		var v75: set<bool> := {v71, f6};
		var v76: seq<int> := [-|v75|];
		v74[399] := v76[|v53|];
		v74[399] := f5 + |(v75 - v75)|;
	}
	method m2(p0: int, p1: string, p2: int, p3: bool, globalState: GlobalState) returns (r0: array<array<D0>>, r1: D0) {
		var v0 := -0x189;
		var v1: set<int> := {p2, 0x1ae};
		var v2: multiset<int> := multiset{|v1|, v0};
		var v3: T0 := new C1(|v2|, f8, v0, f6);
		var v4 := DC16(v3, f6, p0);
		var v5: T0 := new C1(-929, v4.cf34, p0, v3.f6);
		var v6: seq<int> := [|p1|, 0x3c8];
		var v7: set<bool> := {f6, f6, v5.f6, f6};
		var v8: seq<bool> := [v5.f6, true];
		var v9: map<int, int> := map[|v7| := |v8|];
		var v10: map<int, bool> := map[0xb := v3.f6];
		var v11: map<bool, bool> := map[v5.f6 := if (v5.f5 in v10) then v10[v5.f5] else v5.f6];
		var v12: array<int> := new int[9] [p0, |v6|, v0, |(v9 + v9)|, 0x299 + 0x30c, v5.f5, v3.f5 / f5, v3.f5, fm6(p0, |v11|, v5.f6, fm6(fm6(v3.f5, f8, v3.f6, v5.f5, globalState), v0, p3, p0, globalState), globalState)];
		v12[109] := f8 % v3.f5;
		var v13 := DC6(v3.f6 || f6, -v3.f5, |"a"|, v5.f6);
		var v14: multiset<bool> := multiset{!v3.f6};
		var v15: map<bool, T0> := map[!v3.f6 := v5];
		var v16: map<bool, seq<int>> := map[p3 := v6];
		v0, v5, v12[109], v13, v5 := fm6(|fm19(|v14|, globalState)|, -(|p1| * p0), p3, |"doyklvo"|, globalState), if (f6 in v15) then v15[f6] else v3, fm6(v3.f5, |(v6 + (if (v3.f6 in v16) then v16[v3.f6] else v6))|, v3.fm0(p3, globalState), v0 % p2, globalState), if (fm0(p3, globalState)) then v13 else fm20(-|p1|, v14, f6, f6, globalState), v3;
		if (if (v3.f6 in v11) then v11[v3.f6] else fm0(v5.f6, globalState)) {
			var v17 := new C0(p2, |v7|);
			if (v14 >= v14) {
				v12[109] := v17.f12;
				v0 := fm6(v0 * v5.f5, -p2, v3.f6, v17.f13 * p2, globalState);
				var v18: map<int, string> := map[v17.f13 := seq(737, i0  => ('d'))];
				var v19: array<string> := new string[17] [if (v17.f12 in v18) then v18[v17.f12] else p1, p1, "xvmci", seq(0x3c2, i1  => ('a')), seq(0x38c, i2  => ('v')), "ea", p1, p1, p1, p1, seq(0x354, i3  => ('r')), p1, p1, p1, p1, seq(0x207, i4  => ('y')), p1];
				var v20: map<int, array<string>> := map[v5.f5 := v19];
				v20 := v20[-f8 := v19];
				v19[980] := p1;
				var v21: array<C2> := new C2[28];
				var v22 := 'g';
				var v23 := DC20(v22);
				var v24: array<bool> := new bool[22] [v3.f6, v3.f6, f6, !(if (v3.f6 in v11) then v11[v3.f6] else v3.f6) <== v3.f6, v3.f6, -v5.f5 != v17.f12, if (v3.f6) then !v3.f6 else v3.f6, v3.f6, v5.f6 ==> v3.f6, v3.f6, true, v3.f6, v3.f6, v3.f6, v5.f6 && f6, v3.f6, v5.fm0(v5.f6, globalState), v3.f6, v3.fm0(v5.f6, globalState), p3, v5.f6, !(if (v17.fm8(p3, "vjkgdpeu", v23.cf40, globalState) in v11) then v11[v17.fm8(p3, "vjkgdpeu", v23.cf40, globalState)] else v3.f6)];
				var v25: map<int, array<bool>> := map[129 := v24];
				v19[980], v0, v21, v24 := p1 + p1, 0x1e1, v21, if (v0 in v25) then v25[v0] else v24;
				var v26 := DC17(v8);
				v26 := v26.(cf35 := v8 + v8);
			} else {
				var v27 := DC4(v5.f6, p3, p1 + "ahcbwma", v5.f5, |v10| / v17.f12);
				v27 := v27;
				var v29: seq<set<int>> := [v1, set v28 : int | (626 <= v28) && (v28 < 0x34b) :: (v28 + v5.f5), v1];
				var v31: array<set<int>> := new set<int>[20] [v1 - v1, v1, v1, fm21(globalState), v1, v29[|p1|], {v3.f5} + v1, v1 - {-219}, v1, v1, v1, set v30 : int | (-558 <= v30) && (v30 < -0x2fc) :: (v30 / v12[109]), v1 + v1, v1, v1, v1, v1, v1, v1, v1];
				v31[347] := v1;
				v31[347] := v1 * (v1 + v1);
				var v32: array<bool> := new bool[5](i5 => v3.f6);
				v32[721] := v3.f6;
				v32[721] := false ==> ((seq(-432, i6  => ('w'))) != p1);
				v32[721], v12[109] := false ==> v32[721], (p0 % 0x2a0) + v3.f5;
				var v33 := DC11(v12);
				v33 := v33;
			}
			
			var v34: array<bool> := new bool[9] [v3.f6, v3.f6, p3, p3, v3.f6, v3.f6, v3.f6, false, v5.f6];
			v34[913] := v5.f6;
			v34[913] := p3;
			v34[913] := v5.f6;
			v12[109] := v12[109];
		} else {
			var v35: map<seq<int>, bool> := map[v6 := v3.f6];
			var v36: array<bool> := new bool[23] [true, v5.f6, v3.f6, v1 !! v1, v5.f6, true, v3.f6, v5.f6, v3.f6, p3, v5.fm0(v3.f6, globalState), f8 < v0, if (p3) then v3.f6 else p3, v5.f6, v5.f6, f6, if ([p2, v5.f5, v12[109], v12[109]] in v35) then v35[[p2, v5.f5, v12[109], v12[109]]] else f6, true, v3.f6, f6, f6, p3 <== v5.f6, v3.f6 <== v8[-|multiset{fm0(v3.f6, globalState), f6}|]];
			v36[375] := p3;
			v36[375] := p3;
			var v37: map<int, array<bool>> := map[p2 := v36];
			v37 := v37[v12[109] := v36];
			v8 := fm16(globalState);
			v36[375] := v5.f6;
			var v38: seq<string> := [p1];
			var v39 := DC4(p0 > v0, v3.fm0(p3, globalState), p1, v3.f5 - |v38[-v12[109]]|, v3.f5);
			match (v39) {
				case DC4(cf8, cf9, cf10, cf11, cf12) =>
					var v40: map<bool, int> := map[cf9 := v3.f5];
					var v41: map<map<bool, bool>, bool> := map[v11 := -fm6(|v40|, v3.f5, p3, |"htbyp"|, globalState) != v3.f5];
					v41 := v41;
					var v42 := 'e';
					v42 := v42;
					var v43: array<T0> := new T0[29];
					v43[199] := v5;
					v43[199] := if (p3 in v15) then v15[p3] else v5;
					cf8 := v3.f6;
				case DC5(cf13, cf14, cf15, cf16) =>
					cf15 := v5.f5;
					var v44 := 'e';
					v12, cf14, cf13 := v12, DC4(v5.f6, if (false in v11) then v11[false] else p3, p1, p0, v5.f5).cf8, v44 !in "vhrqm";
					cf16 := |(seq(138, i7  => (v3.f5)))|;
					cf14 := v5.f6;
				case DC6(cf17, cf18, cf19, cf20) =>
					var v45: C2 := new C2(v3.f5);
					var v46: map<array<int>, C2> := map[v12 := v45];
					var v47: seq<array<int>> := [v12];
					var v48: map<C2, set<int>> := map[if (v47[v3.f5] in v46) then v46[v47[v3.f5]] else v45 := {v3.f5, p2}];
					v48 := v48;
					v36 := v36;
					var v49 := 'k';
					cf20 := !v45.fm1(p0, v49, globalState);
					var v50: array<array<map<int, int>>> := new array<map<int, int>>[27];
					var v52: seq<map<int, int>> := [v9];
					var v53: array<map<int, int>> := new map<int, int>[15] [v9, v9, v9, v9, map v51 : int | v51 in v9 :: (v51 * cf19) := (cf19), v9[v3.f5 := v5.f5], v9, v9, v9, v52[-0x108], v9, v9[0x3c7 := f8], map[515 := v0], v9, v9];
					v50[473] := v53;
					var v54: seq<array<map<int, int>>> := [v53, v53, v53, v53];
					v50[473] := v54[-573];
				case DC3(cf7) =>
					v36[375] := v5.f6;
					v12[109], v12[109] := v5.f5, v3.f5;
					var v55 := DC11(v12);
					var v56: map<D3, int> := map[v55 := v5.f5];
					v56 := v56[if (f6) then v55 else DC11(v12) := v3.f5];
					var v57: map<int, string> := map[v5.f5 := p1];
					v57 := v57[v5.f5 + p0 := p1];
				case DC7(cf21) =>
					v36[375] := v3.f6;
					v0 := p2;
					var v58: array<array<bool>> := new array<bool>[9];
					var v59: seq<array<bool>> := [v36, v36];
					var v60: map<bool, int> := map[true := v13.cf19];
					v36[375], v58, v12[109], v12[109] := !(true <==> (v36 !in v59)), v58, |(v14 - v14)| % |v60|, fm6(v0, v3.f5 + p2, true, 0xc5, globalState);
					v36[375] := true;
			}
			
		}
		
		var v61 := new C2(v12[109] * -v0);
		var i8 := 0;
		while (!v5.f6)
			decreases 100 - i8
		{
			if (i8 >= 100) {
				break;
			}
			
			i8 := i8 + 1;
			var v62: array<string> := new string[24];
			var v63: map<seq<char>, array<string>> := map[if (!false) then p1 else seq(0x1b9, i9  => ('b')) := v62];
			v63 := v63[p1 + p1 := v62];
			var v64: map<bool, int> := map[v3.f6 := v3.f5];
			var v65: seq<seq<int>> := [v6, seq(141, i10  => (p2))];
			v64 := v64[v65[v5.f5] <= v6 := -v5.f5];
			var v66 := new C2(0x249 % -f8);
			v62[434] := p1;
			var v67 := true;
			var v68: map<int, map<bool, int>> := map[v5.f5 := v64];
			var v69: seq<string> := [fm5(map[p3 := |[v5.f5]|], v68, globalState), p1];
			v62[434], v61.f9, v67, v67, v67 := p1 + v69[|v8|], f5, p3, (v3.f5 <= v3.f5) ==> f6, v61.f9 != v5.f5;
		}
		var v70 := DC11(v12);
		var v71: seq<D3> := [v70];
		var v72: multiset<D3> := multiset{v70};
		var v73: array<multiset<D3>> := new multiset<D3>[6] [multiset(v71 + v71), multiset{v70} - v72, v72, v72, v72, v72 - v72];
		v73 := DC21(v73).cf41;
		var v74: array<map<int, int>> := new map<int, int>[2](i11 => v9);
		v74 := v74;
		var v75: array<array<D0>> := new array<D0>[21];
		var v76: seq<array<array<D0>>> := [v75, v75];
		r0 := v76[v0];
		var v77 := DC0(v6);
		r1 := v77;
	}
}

class C4 extends T0 {
	var f7 : int
	constructor (f7 : int, f5 : int, f6 : bool) {
		this.f7 := f7;
		this.f5 := f5;
		this.f6 := f6;
	}
	
	function fm0(p0: bool, globalState: GlobalState): bool {
		match DC0([f7]) {
			case DC1(cf1, cf2, cf3, cf4) => true
			case DC2(cf5, cf6) => f6 in {f6}
			case DC0(cf0) => !f6
		}
	}
	method m1(p0: int, globalState: GlobalState) {
		var v0: seq<bool> := [f6, f6];
		var v1: seq<int> := [f7, p0, p0, |v0|, p0];
		var v2 := DC2(map[false := v1], f6);
		var v3: set<int> := {f7, f7};
		var v4 := DC6(f6, f7, |v3|, false);
		var v5: C0 := new C0(|"fcehnr"|, v4.cf19);
		var v6: map<C0, int> := map[v5 := v5.f12];
		var v7: multiset<map<C0, int>> := multiset{v6, v6};
		var v8: map<int, int> := map[-0x66 := f7];
		var v9: C3 := new C3(v5.f12, -|v8|, f6);
		var v10: T0 := new C3(|v7|, |{v9, v9}|, f6);
		var v11: set<T0> := {v10};
		var v12 := "nhstj";
		var v13: array<bool> := new bool[17] [v2.cf6, 0x38 > p0, f6, v1 == v1, f6, f6, f6, f7 !in v1, f6, !f6, f6, f6, v11 !! v11, v12 == v12, f6 && f6, v10.f6, f6];
		v13[428] := v10.f6;
		v13[428] := v10.f6;
		v13[428] := v10.f6;
		if (f6) {
			var v14: array<multiset<D3>> := new multiset<D3>[17];
			var v15 := DC21(v14);
			v13[428], v15, v13[428], v12, f7 := false, v15, !v10.f6, v12, v9.f8;
			v13[428] := if (v13[428]) then v13[428] else v10.f6;
			var v16: multiset<string> := multiset{"jijr"};
			v16 := v16 - v16;
			var v17, v18, v19, v20 := m0(v5.f12 * f7, 253, globalState);
			v19 := f7;
		} else {
			var v21 := DC16(v10, v10.f6, --v10.f5);
			var v22: map<bool, int> := map[v21.cf33 := v1[f7]];
			var v23: map<int, bool> := map[|v22| := true];
			var v24 := new C3(v9.f8, |v23|, false);
			var v25 := new C3(v9.f8, f7 * (if (v13[428] in v22) then v22[v13[428]] else |multiset{v13[428]}|), v24.f8 > v24.f8);
			var v26 := 'e';
			v12 := (v12 + v12) + v12[fm6(-0x222, v5.f12, v10.f6, v5.f13, globalState) := v26];
			var v27 := new C0(v25.f8, v10.f5);
			f7 := fm6(v27.f12, f7, v10.f6, f5, globalState);
		}
		
		var i0 := 0;
		while (v10.f6)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v28: array<D5> := new D5[29](i1 => DC15(multiset(v0)));
			var v29: multiset<bool> := multiset{v10.f6};
			var v30 := DC15(v29[f6 := |v12|]);
			v28[784] := v30.(cf31 := fm10(multiset{v5.f12, v5.f12, -0x35f}, f5, globalState));
			var v31: C1 := new C1(|['i']|, f5, v9.f8, false);
			var v32: map<bool, C1> := map[v10.f6 := v31];
			var v33: map<bool, int> := map[f6 := f7];
			v28[784] := fm22(|(v32 + v32)|, v12, v31.f11 - DC1(v12, v9.f8, v33, v1).cf2, seq(0x7e, i2  => ('o')), globalState);
			var v34: multiset<int> := multiset{-v5.f12, |{f5}|};
			var v35 := DC8(v34);
			v35 := fm23(f7, |v12| < |v1|, 237, v5.f12 + |v12|, globalState);
			var v36: map<int, bool> := map[v9.f8 := v12 != v12];
			var v37 := DC16(v10, f6, v5.f13);
			v36 := v36[v37.cf34 := f6];
			f7, v36 := -v5.f12, (map[v10.f5 := v10.f6])[v10.f5 * v5.f13 := v13[428]];
		}
		v12 := v12;
		var v38: array<int> := new int[5] [f5, f7, v5.f12, v10.f5, f5];
		var v39: multiset<array<int>> := multiset{v38};
		var v40: array<multiset<array<int>>> := new multiset<array<int>>[7] [v39, v39 * v39, multiset{v38, v38, v38, v38, v38}, v39 + multiset{v38}, v39, multiset{v38, v38}, v39];
		v40[860] := v39;
		v40[860] := v39;
	}
}

method Main() {
	var v0 := false;
	var v1 := 0x52;
	var v2 := 'd';
	var v3: multiset<char> := multiset{v2};
	var v4: map<bool, int> := map[v0 := v1];
	var v5: map<bool, bool> := map[v0 := v0];
	var v6: map<int, bool> := map[|v5| := v0];
	var v7 := "oxasa";
	var v8: seq<bool> := [v0];
	var v9: array<int> := new int[29] [v1, if (v2 in v3) then v3[v2] else 0x57, 0x368, |v4|, v1, v1, v1, -v1, |v6|, |[v1, v1]|, -v1, |v7|, v1, |v7|, |(seq(325, i0  => (v2)))|, v1, v1, v1, v1, v1, |v8|, v1, |(seq(0x2f, i1  => (v2)))|, v1, v1, -v1, v1, v1, v1];
	var v10: map<bool, array<int>> := map[v0 := v9];
	var v11: set<int> := {v1};
	var v12: seq<string> := [seq(0x1db, i2  => (v2))];
	var globalState := new GlobalState(0x2fe, v10, [v11], v12, true);
	var v13, v14, v15, v16 := m0(0x35d + |(seq(0x1eb, i3  => (v2)))|, v1, globalState);
	var v17: multiset<int> := multiset{v16, v15, v15};
	v16 := v16 % |(v17 * v17)|;
	var v18: array<multiset<int>> := new multiset<int>[6];
	v18[816] := v17;
	var v19: seq<multiset<int>> := [v17];
	var v20: seq<int> := [0x2d0];
	v18[816] := (v17 * v19[|v13|]) * (multiset(v20) + multiset((DC1(v13, v15, v4, ([|v11|, v1])[v15 := v15]).(cf3 := v4, cf1 := v7)).cf4));
	var v21: map<int, int> := map[v16 := 0x7d % v16];
	var v22 := DC1("vewema", v1, v4, v20);
	v21 := v21[v15 - v22.cf2 := -(0x2e5 * v16)];
	v15 := v16;
	var v23: array<map<bool, bool>> := new map<bool, bool>[27](i4 => v5);
	v23[758] := v5 + v5;
	v23[758] := v5 + (v5 + v5);
	v2 := 'y';
	var v24: array<bool> := new bool[2];
	var v25: map<array<bool>, bool> := map[v24 := !v0];
	v25 := v25[v24 := !(v0 ==> v0)];
	if (|v11| >= v1) {
		var v26: map<bool, seq<int>> := map[v0 := v20];
		var v27 := DC2(v26, !v0);
		match (v27) {
			case DC1(cf1, cf2, cf3, cf4) =>
				v0 := ((seq(-964, i5  => (v2)))[v16 := v2] + "l")[v15 := v2] == v7;
				var v28 := new C0(cf2, -v1);
				cf2 := v28.f13 / cf2;
				v15 := v16 - v28.f13;
			case DC2(cf5, cf6) =>
				v15 := v16;
				v24[847] := cf6;
				v9[116] := v15;
				var v29: C0 := new C0(v1, v15);
				var v30: map<C0, bool> := map[v29 := true];
				var v31: map<int, map<bool, int>> := map[|v30| := v4];
				v24[187] := v15 > |fm5(v4, v31, globalState)|;
				var v32: set<bool> := {true};
				v24[847], v18[816], v9[116], v24, v24[187] := v0 <==> false, v17 + v18[816], v29.f12, v24, !(if (false) then v32 >= v32 else v0);
				v2 := fm11(v2, !v24[847] || v24[847], v29.fm8(v24[847], v7, v2, globalState), globalState);
				var v33: T0 := new C4(v29.f12, v9[116], cf6);
				v33 := v33;
			case DC0(cf0) =>
				var v34, v35, v36, v37 := m0(v15, v1, globalState);
				var v38: array<string> := new string[11];
				v38[373] := v13;
				v38[373] := v13;
				var v39: array<char> := new char[20](i6 => 'u');
				v39[60] := v2;
				v39[60] := v2;
				v15 := v1;
		}
		
		var v40 := DC20(v2);
		v40 := DC20(v2);
		v14 := new int[20];
		v2 := fm11(v2, v0, v0 ==> fm24(globalState), globalState);
		var v41: T0 := new C1(v15, v15, v15, v0);
		var v42: array<T0> := new T0[3] [if (v0) then v41 else v41, v41, v41];
		v42[472] := v41;
		v42[472] := if (v41.f6) then v41 else v41;
	} else {
		var v43 := new C3(v16, --0x1d5 - v1, v0);
		v16 := v43.f8;
		var v44: set<bool> := {v0};
		var v45: map<set<bool>, string> := map[v44 := v7];
		var v46: T0 := new C1(v43.f8, 431, v15, v0);
		var v47: set<T0> := {v46, v46};
		v0 := v2 !in (if (v44 in v45) then v45[v44] else v13)[|v47| := v2];
		v0, v16 := if (v0) then true else v46.f6, (v15 + v1) * -(|v8| - |v8|);
		v2, v16 := v2, -v16;
	}
	
	v1 := v20[v15] % v15;
	if (fm24(globalState)) {
		v0 := v0;
		var v49 := DC3(map v48 : int | v48 in {v15} :: (v48 % 236) := (v0));
		var v50: C3 := new C3(-0x12e, |v20|, v0);
		var v51: T0 := new C1(v16, v15, |map[v49 := v50]|, v0);
		var v52 := DC16(v51, v51.f6, v16);
		var v53 := new C0(fm6(|multiset{v0}|, v1, v0, v1, globalState) + fm6(|v20|, v16, v0, |"puoy"|, globalState), -v16 % v52.cf34);
		v2 := fm11(v2, v0 && v51.f6, v51.f6, globalState);
		v2 := v2;
		v9[658] := v53.f13 - v53.f13;
		v13, v1, v9[658] := v13, v53.f12 + v15, v1;
	} else {
		var v54: array<seq<bool>> := new seq<bool>[4];
		v54 := v54;
		v24 := v24;
		v0 := v0;
		var v55: map<bool, seq<int>> := map[true := v20];
		match (DC12(DC2(v55, fm24(globalState)).cf6, |v17|).(cf28 := v0)) {
			case DC12(cf28, cf29) =>
				var v56: map<string, int> := map["skjvslqs" := v15];
				var v57: map<map<string, int>, array<bool>> := map[v56 := v24];
				var v58: seq<array<bool>> := [v24];
				v57 := v57[v56 := v58[cf29]];
				var v59, v60, v61, v62 := m0(v15, v1, globalState);
				var v63: C0 := new C0(v1, v16);
				var v64: map<C0, int> := map[v63 := |v18[816]|];
				var v65: multiset<map<C0, int>> := multiset{v64};
				v6 := if (fm24(globalState)) then v6[cf29 := v0] else map[if (v64 in v65) then v65[v64] else v61 := v0];
				v9[831] := cf29;
				var v66: set<bool> := {v0};
				v9[831] := |v66|;
			case DC11(cf27) =>
				var v67: map<set<int>, bool> := map[v11 * v11 := !fm24(globalState)];
				v67 := v67[v11 := v0];
				v1 := |(v22.cf3 + fm13("ptcfinu", v16, v19[v15], globalState))|;
				var v68: map<int, map<bool, int>> := map[v1 := map[v0 := v16]];
				var v69: C0 := new C0(v15 + 0x2d2, |fm5(v4, v68, globalState)|);
				var v70: T0 := new C3(v69.f12, v16, v0);
				v69, v70 := v69, v70;
				var v71: array<string> := new string[24];
				v71[314] := v7 + v7;
				v71[314] := v7;
		}
		
		var v72, v73, v74, v75 := m0(v1 % v15, v16, globalState);
	}
	
	v9 := v14;
	var v76: C4 := new C4(v16, v15, v0);
	var v77: seq<C4> := [v76];
	var v78 := DC22(v77);
	var v79: multiset<C4> := multiset{v76};
	var v81 := DC10(DC8(v17[v1 := -|(map v80 : char | v80 in v7 :: (v80) := (!v0))|]));
	var v82 := DC8(v18[816]);
	match (if (multiset((v78.(cf42 := v77)).cf42) > v79) then v81 else DC10(v82)) {
		case DC9(cf23, cf24, cf25) =>
			v1 := v15;
			v13 := seq(-0x2d, i7  => (v7[v76.f7]));
			cf23 := cf23;
			var v83: C0 := new C0(v1, v15);
			var v84: map<bool, C0> := map[v0 := v83];
			var v85 := new C2(|v84|);
		case DC8(cf22) =>
			v14[5] := v76.f7 + -v76.f7;
			v14[5] := v76.f7;
			v76.m1(v76.f7, globalState);
			var v86: array<seq<bool>> := new seq<bool>[21];
			v86[224] := v8[|v8| := v0];
			v86[224] := (if (v0) then [v0, true] else v8) + v8;
			var v87: C0 := new C0(v1, |v8|);
			var v88: map<set<C0>, int> := map[{v87} := v87.f13];
			var v89: set<C0> := {v87};
			var v90: map<map<set<C0>, int>, bool> := map[(v88[v89 := |v21|])[v89 := |v11|] := !!v0];
			v90 := v90[v88 := !v87.fm8(v0, v22.cf1, v2, globalState)];
		case DC10(cf26) =>
			v15 := fm6(v15, v76.f7 % v76.f7, !v76.fm0(false, globalState), |v18[816]|, globalState);
			var v91: array<C4> := new C4[4] [v76, v76, v76, v76];
			var v92 := DC11(v14);
			var v93: map<array<C4>, D3> := map[v91 := v92];
			v93 := v93[v91 := v92];
			v21 := v21[v15 := if (true) then v76.f7 else v15];
			var v94: map<bool, array<bool>> := map[v0 := v24];
			var v95: T0 := new C3(v16, v1, v0);
			var v96: set<T0> := {v95};
			var v97: map<map<bool, array<bool>>, set<T0>> := map[v94 := v96];
			v97 := v97[v94 := v96];
	}
	
	var i8 := 0;
	while (v76.fm0(v1 <= 0x17f, globalState))
		decreases 100 - i8
	{
		if (i8 >= 100) {
			break;
		}
		
		i8 := i8 + 1;
		var v98: set<array<bool>> := {v24, v24};
		var v99, v100, v101, v102 := m0(|v98| % v76.f7, if (v0) then v15 else v76.f7, globalState);
		v0 := if (v99 < v13) then fm24(globalState) else v0;
		var v103 := new C0(v101, v15);
		var v104: C1 := new C1(v15, fm6(v16, v102, v0, v101, globalState), v103.f12, v0);
		var v105 := DC14();
		var v106: map<C1, D4> := map[v104 := v105];
		v76.f7 := fm6(0x94 / v76.f7, |v106|, v16 != v76.f7, v16 * v1, globalState);
	}
	var v107, v108, v109, v110 := m0(if (v16 in v18[816]) then v18[816][v16] else v16, v15, globalState);
	var v111: multiset<bool> := multiset{v0, !v0};
	v0 := (v0 !in v111) <==> v0;
}