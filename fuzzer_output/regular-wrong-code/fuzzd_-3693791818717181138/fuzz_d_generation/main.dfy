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
datatype D0 = DC0 | DC1(cf0: bool, cf1: seq<bool>, cf2: bool)
datatype D1 = DC2(cf3: char)
datatype D2 = DC4(cf5: bool) | DC3(cf4: int)
datatype D3 = DC6(cf7: bool, cf8: int, cf9: int) | DC5(cf6: seq<int>)
datatype D4 = DC8(cf11: int, cf12: bool) | DC9(cf13: bool, cf14: int, cf15: array<bool>, cf16: int) | DC7(cf10: string)
datatype D5 = DC11(cf18: int, cf19: string) | DC10(cf17: array<int>)
datatype D6 = DC13(cf21: bool, cf22: bool, cf23: string, cf24: bool) | DC12(cf20: array<array<int>>)
datatype D7 = DC15 | DC14(cf25: set<bool>) | DC16(cf26: D7)
datatype D8 = DC18(cf28: map<int, int>, cf29: bool, cf30: int) | DC19(cf31: bool) | DC17(cf27: set<int>)
datatype D9 = DC21(cf33: int, cf34: int) | DC22(cf35: int, cf36: int, cf37: bool, cf38: int, cf39: map<T2, int>) | DC23(cf40: T1, cf41: bool) | DC20(cf32: multiset<int>)
datatype D10 = DC25(cf43: string) | DC24(cf42: map<bool, bool>)
datatype D11 = DC27(cf45: bool, cf46: map<int, multiset<int>>) | DC28(cf47: int, cf48: array<bool>) | DC26(cf44: seq<seq<bool>>)
datatype D12 = DC30(cf50: int, cf51: int) | DC31(cf52: bool, cf53: bool, cf54: int) | DC29(cf49: C3) | DC32(cf55: D12)
datatype D13 = DC34(cf57: D11, cf58: string, cf59: bool) | DC33(cf56: array<array<bool>>) | DC35(cf60: D13)
datatype D14 = DC37(cf62: char, cf63: bool, cf64: int, cf65: T1) | DC36(cf61: T2)
datatype D15 = DC39(cf67: bool, cf68: int, cf69: bool) | DC40(cf70: bool, cf71: int) | DC38(cf66: C2)
datatype D16 = DC42(cf73: seq<bool>, cf74: seq<D15>) | DC41(cf72: seq<string>)
trait T0 {
	function fm0(globalState: GlobalState): bool 
	function fm1(p0: string, p1: D0, globalState: GlobalState): bool 
}

trait T1 extends T0 {
	var f8 : bool
	function fm2(p0: D0, globalState: GlobalState): int 
}

trait T2 {
	const f9 : bool
	method m0(p0: multiset<bool>, p1: D0, p2: bool, globalState: GlobalState) 
	method m1(p0: string, p1: bool, globalState: GlobalState) returns (r0: string, r1: int, r2: bool) 
}

class GlobalState {
	const f0 : char
	const f1 : int
	const f2 : int
	const f3 : int
	const f4 : string
	const f5 : int
	const f6 : array<bool>
	const f7 : int
	constructor (f0 : char, f1 : int, f2 : int, f3 : int, f4 : string, f5 : int, f6 : array<bool>, f7 : int) {
		this.f0 := f0;
		this.f1 := f1;
		this.f2 := f2;
		this.f3 := f3;
		this.f4 := f4;
		this.f5 := f5;
		this.f6 := f6;
		this.f7 := f7;
	}
	
}

class C0 extends T0 {
	const f12 : array<array<int>>
	constructor (f12 : array<array<int>>) {
		this.f12 := f12;
	}
	
	function fm0(globalState: GlobalState): bool {
		-0x327 > 402
	}
	function fm1(p0: string, p1: D0, globalState: GlobalState): bool {
		("eacmud" + "h") == (if (!false) then "gdjemsj" else seq(abs(0x21f), i0  => ('n')))
	}
	function fm10(p0: int, p1: int, p2: bool, globalState: GlobalState): map<int, int> {
		((map v0 : int | (0x34e <= v0) && (v0 < 0x29) :: (v0 - 450) := (0x226)) + map[-0x261 := --0x272]) + (if (true) then map v1 : int | (942 <= v1) && (v1 < 0x34d) :: (safeDivisionInt(v1, |(seq(abs(0x113), i0  => (0x3ab)))|)) := (|multiset([0x210, 570])|) else map[299 := |[0x351]|])
	}
}

class C1 extends T2, T0, T1 {
	const f10 : char
	const f11 : set<seq<D4>>
	constructor (f10 : char, f11 : set<seq<D4>>, f9 : bool, f8 : bool) {
		this.f10 := f10;
		this.f11 := f11;
		this.f9 := f9;
		this.f8 := f8;
	}
	
	function fm0(globalState: GlobalState): bool {
		f8
	}
	function fm1(p0: string, p1: D0, globalState: GlobalState): bool {
		"sapwe" in (map v0 : string | v0 in ["awbrbsnwd"] :: (v0) := (f9))
	}
	function fm2(p0: D0, globalState: GlobalState): int {
		safeModuloInt(0x15c + 0x3bf, |multiset{f9}| + |multiset(seq(abs(429), i0  => (-0x33f)))|)
	}
	function fm8(globalState: GlobalState): bool {
		f8
	}
	method m0(p0: multiset<bool>, p1: D0, p2: bool, globalState: GlobalState) {
		var v0: array<seq<bool>> := new seq<bool>[16];
		var v1: seq<bool> := [p2, p2, f9, f8, true];
		v0[safeIndex(467, v0.Length)] := v1;
		var v2 := DC4(false);
		v0[safeIndex(467, v0.Length)] := v1 + [f8, v2.cf5];
		var v3 := -0x3a0;
		var v4 := DC1(f8, v1, p2);
		var v5 := "riwfnhqa";
		for i0 := --(if (p2) then v3 else fm2(v4, globalState)) to |v5| {
			var v6: map<int, string> := map[v3 := v5];
			var v7: seq<string> := [v5];
			var v8: array<string> := new string[25] [v5, fm9(!f8, f8, globalState), v5, v5, v5, v5, v5, v5, (if (fm2(v4, globalState) in v6) then v6[fm2(v4, globalState)] else v5) + (seq(abs(0x1ae), i1  => (f10))), "fltsycqgp", v5 + v5, v5, v5, v5, "rboifjmne", v5, v5 + "mk", v5, ("nvmycqm")[safeIndex(i0, |"nvmycqm"|) := f10], v5, "tjel", v5 + (seq(abs(0x319), i2  => (f10))), v7[safeIndex(fm2(v4, globalState), |v7|)], seq(abs(0x153), i3  => (f10)), v5];
			v8[safeIndex(990, v8.Length)] := v5;
			v8[safeIndex(990, v8.Length)] := v5[safeIndex(i0, |v5|) := 'n'];
			f8 := !p2;
			f8 := !f9;
			f8 := fm8(globalState);
		}
		var v9: set<int> := {v3};
		f8 := v9 < v9;
		var i4 := 0;
		while (v3 <= v3)
			decreases 100 - i4
		{
			if (i4 >= 100) {
				break;
			}
			
			i4 := i4 + 1;
			var v10: set<bool> := {f8, false};
			var v11: seq<set<bool>> := [v10];
			var v12: map<bool, seq<set<bool>>> := map[!("yial" < v5) := v11];
			v12 := v12[f8 := v11];
			var v13: array<bool> := new bool[18];
			var v14: map<bool, bool> := map[p2 := p2];
			v13[safeIndex(518, v13.Length)] := if (false in v14) then v14[false] else fm8(globalState);
			v13[safeIndex(518, v13.Length)] := fm0(globalState);
			v13[safeIndex(518, v13.Length)] := f9;
			f8 := f8;
		}
		if (f8) {
			var v15: array<bool> := new bool[22](i5 => f8);
			v15 := v15;
			var v16: array<string> := new string[8] ["atblghtma", seq(abs(0x85), i6  => ('o')), v5, v5, v5, v5, "hvtuvijkn", v5 + v5];
			var v17: seq<string> := [v5];
			v16[safeIndex(788, v16.Length)] := v17[safeIndex(|(multiset{f8, !p2, f8})[!f8 := abs(v3)]|, |v17|)];
			var v18: map<int, int> := map[v3 := v3];
			v16[safeIndex(788, v16.Length)] := (if (p2) then v17[safeIndex(if (v3 in v18) then v18[v3] else v3, |v17|)] else v5 + v5)[safeIndex(0x2f2 * |(seq(abs(0x1ad), i7  => ('u')))|, |(if (p2) then v17[safeIndex(if (v3 in v18) then v18[v3] else v3, |v17|)] else v5 + v5)|) := f10];
			v16[safeIndex(788, v16.Length)] := (v16[safeIndex(788, v16.Length)] + v16[safeIndex(788, v16.Length)]) + v16[safeIndex(788, v16.Length)];
			var v19: map<bool, bool> := map[false := f8];
			var v20: multiset<int> := multiset{|v16[safeIndex(788, v16.Length)]|, v3, |p0|, v3};
			v19 := v19[f8 := |v19| !in v20];
			if (!f9 ==> f8) {
				f8 := safeDivisionInt(v3, v3) > (0x2fe + |v16[safeIndex(788, v16.Length)]|);
				var v21: seq<int> := [v3, v3, v3];
				var v22 := DC7("arjmj");
				var v23: map<D4, int> := map[v22 := v3];
				var v24: array<int> := new int[9] [v3, -v3, |v21|, v3 - 0x382, fm2(v4, globalState), safeDivisionInt(v3, |[v3, fm2(v4, globalState)]|), v3, v3, v3 - (if (v22 in v23) then v23[v22] else fm2(DC1(f8, v0[safeIndex(467, v0.Length)], !f8), globalState))];
				v24[safeIndex(585, v24.Length)] := v3;
				var v25: set<bool> := {f8};
				v24[safeIndex(585, v24.Length)] := |v25|;
				v15[safeIndex(103, v15.Length)] := true;
				v15[safeIndex(103, v15.Length)] := f9;
				var v26: seq<seq<int>> := [v21, v21];
				v26 := v26;
				var v27: array<map<char, map<bool, int>>> := new map<char, map<bool, int>>[4];
				var v28: map<bool, int> := map[f9 := v24[safeIndex(585, v24.Length)]];
				v27[safeIndex(156, v27.Length)] := map[f10 := v28];
				var v30: map<char, map<bool, int>> := map[f10 := v28];
				v27[safeIndex(156, v27.Length)] := (map v29 : char | v29 in (seq(abs(-0x2a), i8  => (f10))) :: (v29) := (v28)) + v30;
			} else {
				var v31: map<string, int> := map[v5 := v3];
				v31 := v31[v16[safeIndex(788, v16.Length)] := v3];
				v2 := v2;
				var v32: array<array<int>> := new array<int>[8];
				var v33 := new C0(v32);
				var v34: array<char> := new char[16](i9 => f10);
				v34[safeIndex(42, v34.Length)] := f10;
				var v35: map<int, string> := map[|v18| := v5];
				v34[safeIndex(42, v34.Length)], v15, f8, f8 := f10, v15, v3 > v3, v33.fm1(v16[safeIndex(788, v16.Length)] + (if (-v3 in v35) then v35[-v3] else v17[safeIndex(v3, |v17|)])[safeIndex(fm2(DC1(!!p2, v0[safeIndex(467, v0.Length)], f8), globalState), |(if (-v3 in v35) then v35[-v3] else v17[safeIndex(v3, |v17|)])|) := f10], v4, globalState);
				f8 := [f8, f9] <= v0[safeIndex(467, v0.Length)];
			}
			
		} else {
			var v36: array<int> := new int[18] [v3 + v3, v3, v3, v3, v3, safeModuloInt(0x14a, v3), 408, v3, fm2(DC1(false, v1, f9), globalState), v3, v3, 392 - v3, v3, v3, |v5|, safeDivisionInt(v3, v3), v3, v3];
			v36 := v36;
			v3 := v3;
			var v37 := DC10(v36);
			var v38: seq<array<int>> := [v36];
			var v39: array<array<int>> := new array<int>[19] [v36, v36, v37.cf17, v36, v36, v36, v36, v36, v36, v36, v36, v36, v36, v36, v36, v38[safeIndex(v3, |v38|)], v36, v36, v36];
			var v40 := new C0(v39);
			v36[safeIndex(49, v36.Length)] := v3;
			v36[safeIndex(49, v36.Length)], v3, v3, v5 := (fm2(v4, globalState) + v3) + safeDivisionInt(v3, v3), v3, safeDivisionInt(|"gucbasaxj"|, v3), v5;
			var v41: C0 := new C0(v40.f12);
			v41 := v40;
		}
		
		var v42 := DC6(f9, v3, |v5|);
		var v43: map<bool, D3> := map[f9 := v42];
		var v44: multiset<int> := multiset{v3};
		v43 := v43[if (p2) then p2 else f8 := DC6(f9, |v44[v3 := abs(v3)]|, 307)];
	}
	method m1(p0: string, p1: bool, globalState: GlobalState) returns (r0: string, r1: int, r2: bool) {
		var v0: array<set<int>> := new set<int>[17];
		var v1 := 0x322;
		var v2: set<int> := {v1};
		v0[safeIndex(524, v0.Length)] := v2;
		var v3: set<char> := {f10, f10};
		v0[safeIndex(524, v0.Length)] := {v1 + |v3|};
		for i0 := v1 to |[337, -v1]| {
			if (fm0(globalState)) {
				var v4 := DC11(i0, p0);
				var v5: map<bool, char> := map[p1 := f10];
				v4 := fm11(v5 + v5, f9, !(if (false) then p1 else f8), f8, globalState);
				var v6: array<bool> := new bool[4];
				v6[safeIndex(547, v6.Length)] := f9;
				var v7: seq<map<bool, char>> := [v5];
				var v8: multiset<bool> := multiset{p1};
				var v9: array<int> := new int[1] [i0];
				var v10: multiset<array<int>> := multiset{v9, v9};
				var v11 := DC10(v9);
				var v12: seq<multiset<array<int>>> := [v10, multiset{v9, v9, v11.cf17, v9, v9}, v10];
				var v13: seq<bool> := [p1];
				var v14: seq<int> := [|v13[safeIndex(i0, |v13|) := p1]|, v1, i0];
				var v15: map<int, bool> := map[i0 := p1];
				var v16 := DC1(p1, v13, f8);
				var v17: multiset<char> := multiset{f10};
				var v18: array<int> := new int[29] [|v7[safeIndex(v1, |v7|)]|, i0 + v1, if (p1 in v8) then v8[p1] else v1, v1, i0, i0, v1, v1, v1, i0, i0, v1, i0, |v12[safeIndex(|v14|, |v12|)]|, v1, v1, -560, safeDivisionInt(|v15|, i0), i0, |v14|, i0, v1, i0, v1, -(v1 + v1), safeModuloInt(i0, i0), fm2(v16.(cf0 := f9), globalState), i0 - (if (f10 in v17) then v17[f10] else i0), -0x241];
				v9[safeIndex(872, v9.Length)] := |(v14 + v14)|;
				r0, v6[safeIndex(547, v6.Length)], v9[safeIndex(872, v9.Length)] := fm9(i0 == i0, true, globalState), p1, i0 - i0;
				var v19 := DC8(|(seq(abs(0x210), i1  => ('b')))|, f9);
				v6[safeIndex(547, v6.Length)] := v19.cf12;
				var v20: seq<string> := [p0, p0, p0];
				var v21: map<seq<string>, bool> := map[v20 := v6[safeIndex(547, v6.Length)]];
				var v22 := DC5(v14);
				v6[safeIndex(547, v6.Length)] := if (fm12(p1, v22, globalState) in v21) then v21[fm12(p1, v22, globalState)] else v13[safeIndex(fm2(v16, globalState), |v13|)];
				r2 := v9[safeIndex(872, v9.Length)] > v9[safeIndex(872, v9.Length)];
			} else {
				var v23: multiset<string> := multiset{p0 + p0, p0 + p0};
				v23 := v23 - (if (p1) then multiset{p0, p0, seq(abs(0x178), i2  => (f10)), seq(abs(989), i3  => (f10))} else multiset{p0, p0, p0});
				var v24: array<array<int>> := new array<int>[11];
				var v25: C0 := new C0(v24);
				var v26: seq<C0> := [v25, v25, v25, v25, v25];
				var v27: map<bool, int> := map[p1 := -0x3df];
				var v28: seq<map<bool, int>> := [map[f8 := v1], v27];
				var v29: seq<int> := [-|v26| * v1, i0, safeModuloInt(i0, v1), i0, i0 * |v28[safeIndex(i0, |v28|)]|];
				v1, r2 := |v29|, p1;
				var v30 := new C0(v24);
				var v31: multiset<int> := multiset{v1};
				v31 := v31 * v31;
				var v32: array<seq<int>> := new seq<int>[2];
				var v33: map<char, array<seq<int>>> := map['e' := v32];
				v33 := v33['h' := v32];
			}
			
			var v34: map<bool, int> := map[true := v1];
			var v35: seq<map<bool, int>> := [v34, v34, fm13(globalState)];
			v35 := v35;
			var v36: array<int> := new int[14];
			var v37: array<array<int>> := new array<int>[4] [v36, v36, v36, v36];
			var v38 := DC12(v37);
			var v39 := new C0(v38.cf20);
			f8, f8 := !!f8, p1;
		}
		var v40: array<seq<int>> := new seq<int>[20](i4 => DC5([v1, |{f8}|]).cf6 + DC5(seq(abs(0x365), i5  => (v1))).cf6);
		v40[safeIndex(488, v40.Length)] := fm14(v1, (seq(abs(104), i6  => (f10)))[safeIndex(347, |(seq(abs(104), i6  => (f10)))|) := f10], f8, globalState);
		var v41: multiset<bool> := multiset{false};
		var v42: map<int, multiset<bool>> := map[v1 := v41];
		var v43: array<int> := new int[14];
		var v44 := DC10(v43);
		var v45: multiset<D5> := multiset{v44, DC10(v43)};
		var v46: seq<int> := [v1 + 0xeb, v1, -0x8f, v1 - |(if (|v45| in v42) then v42[|v45|] else v41)|, 388];
		v40[safeIndex(488, v40.Length)] := v46;
		var v47: array<bool> := new bool[11] [p1, p1, f8, false, p1, f8, f8, f8, v40[safeIndex(488, v40.Length)] != v40[safeIndex(488, v40.Length)], true <==> f9, p1];
		v47[safeIndex(30, v47.Length)] := v1 !in v40[safeIndex(488, v40.Length)];
		v47[safeIndex(30, v47.Length)] := f10 in p0;
		for i7 := v1 to v1 {
			var v48: seq<bool> := [v47[safeIndex(30, v47.Length)], p1, !v47[safeIndex(30, v47.Length)], v47[safeIndex(30, v47.Length)], v47[safeIndex(30, v47.Length)]];
			var v49: map<bool, int> := map[|v48| <= v1 := i7];
			var v50 := DC1(f9, v48, true);
			v49 := v49[f9 := safeDivisionInt(fm2(v50, globalState), 0x3b)];
			var v51: array<array<int>> := new array<int>[20] [v43, v43, v43, v43, v43, v43, v43, v43, v44.cf17, v43, v43, v44.cf17, v43, v43, v43, v43, v43, v43, v43, v43];
			var v52 := new C0(v51);
			var v53: array<map<int, bool>> := new map<int, bool>[25];
			var v54: map<int, bool> := map[i7 := f8];
			v53[safeIndex(167, v53.Length)] := v54;
			v53[safeIndex(167, v53.Length)] := map v55 : int | (796 <= v55) && (v55 < 0xf9) :: (v55 * v1) := (f9);
			r2 := f8;
		}
		v1 := 0x1a5;
		r0 := p0 + p0;
		r1 := v1;
		r2 := v47[safeIndex(30, v47.Length)];
	}
}

class C2 extends T1, T2 {
	constructor (f8 : bool, f9 : bool) {
		this.f8 := f8;
		this.f9 := f9;
	}
	
	function fm2(p0: D0, globalState: GlobalState): int {
		|(if (f8) then DC5([|"efhmg"|, |{!f9, f8}|, |"jpesa"|]) else DC5([0x3dd, 490])).cf6|
	}
	function fm0(globalState: GlobalState): bool {
		0x307 != --|multiset{-788, |map[--|map[f8 := 0x2ce]| := f8]|, |[0x13f]|}|
	}
	function fm1(p0: string, p1: D0, globalState: GlobalState): bool {
		f9
	}
	function fm5(p0: bool, p1: int, globalState: GlobalState): multiset<bool> {
		multiset{f9, f9, f9, f9}
	}
	function fm6(p0: int, p1: int, globalState: GlobalState): int {
		safeDivisionInt(-safeDivisionInt(470, 464), -0x7)
	}
	method m0(p0: multiset<bool>, p1: D0, p2: bool, globalState: GlobalState) {
		var v0 := 195;
		var v1: map<int, char> := map[194 := fm7(v0, f8, |(seq(abs(0xd9), i0  => (v0)))|, globalState)];
		var v2 := 'w';
		v1 := v1[v0 := v2];
		v0 := v0 - (v0 + v0);
		var v3: map<int, int> := map[v0 := v0];
		var v4 := DC6(f9, |v3| * v0, v0);
		match (v4) {
			case DC6(cf7, cf8, cf9) =>
				v2 := if (f9) then v2 else v2;
				var v5: array<bool> := new bool[27](i1 => cf8 != cf8);
				cf7, v0, v5 := f9, safeDivisionInt(v0, cf9) * cf9, v5;
				cf9 := v0 + (cf9 - v0);
				v5 := v5;
			case DC5(cf6) =>
				var v6: array<bool> := new bool[6] [v0 != v0, f9, f8 <==> !f8, true, f9, true];
				v6[safeIndex(389, v6.Length)] := p2;
				v6[safeIndex(389, v6.Length)] := p2;
				v0 := safeDivisionInt(-v0, v0);
				var v7 := "tfsokpds";
				var v8 := DC7("ussmtod");
				var v9: seq<string> := ["ypcmbt", v7 + v8.cf10];
				v7 := v9[safeIndex(v0, |v9|)];
				var v10: seq<D4> := [v8, DC7(v7), v8];
				var v11: set<seq<D4>> := {v10};
				var v12 := new C1(v2, v11, f9, f9);
		}
		
		f8 := f9;
		var v13: map<int, bool> := map[v0 := p2];
		var v14: seq<map<int, bool>> := [v13, map[|multiset{f9}| := p2]];
		f8 := (v14 + v14) < v14;
		var v15: array<int> := new int[2](i2 => safeModuloInt(i2, 0x21c));
		match (DC10(v15)) {
			case DC11(cf18, cf19) =>
				var v16: array<bool> := new bool[29] [p2, !(f8 ==> f9), p2, p2, p2, p2, f9, f8, p2, f8, f8, -v0 == -cf18, f8, f9, f9, f8, cf18 >= fm6(cf18, cf18, globalState), p2, f8 <== f9, p2, f8, p2, f8, !f8, p2, p2, f8, cf18 > 0x228, false];
				v16[safeIndex(408, v16.Length)] := p2;
				var v17 := DC8(v0, f9);
				v16[safeIndex(408, v16.Length)] := v17.cf12;
				cf18, v0, v16[safeIndex(408, v16.Length)] := v0, |map[cf18 * v0 := if (p2) then v16[safeIndex(408, v16.Length)] else f8]|, v16[safeIndex(408, v16.Length)];
				var v18: multiset<int> := multiset{v0, -cf18};
				var v19: map<int, multiset<int>> := map[cf18 := v18];
				var v20: map<bool, bool> := map[f8 := !v16[safeIndex(408, v16.Length)]];
				var v21: map<map<bool, bool>, char> := map[v20 := v2];
				var v22: map<multiset<int>, array<int>> := map[v18 + (if (|v21| in v19) then v19[|v21|] else v18) := v15];
				v22 := v22[v18 := v15];
				var v23: set<int> := {v0 + v0, cf18, v0, cf18};
				var v24: seq<bool> := [f8];
				var v25: map<bool, seq<bool>> := map[(fm15(globalState)).cf5 := v24];
				var v26: seq<int> := [safeModuloInt(cf18, 0x37e)];
				var v27: map<bool, set<int>> := map[v16[safeIndex(408, v16.Length)] := v23];
				var v28: map<char, int> := map[v2 := cf18];
				var v29: map<int, map<char, int>> := map[cf18 := v28];
				var v30: seq<map<bool, seq<bool>>> := [v25];
				v23, v25, v26, cf18 := (if (v16[safeIndex(408, v16.Length)] in v27) then v27[v16[safeIndex(408, v16.Length)]] else {|v29|}) * v23, v30[safeIndex(-v0, |v30|)], v26, safeModuloInt(v0, v0);
			case DC10(cf17) =>
				var v31: multiset<int> := multiset{v0};
				var v32: multiset<bool> := multiset{p2 <==> f8, f9 <==> f9, p2, true, !(|v31| == v0)};
				var v33: set<bool> := {p2};
				var v34 := DC14(v33);
				var v35: seq<bool> := [v34.cf25 !! v33];
				v0, v32 := |v31[v0 := abs(v0)]|, multiset(v35);
				var v36 := "tpt";
				v15[safeIndex(72, v15.Length)] := |(v36 + v36)|;
				v0, v15[safeIndex(72, v15.Length)] := safeDivisionInt(if (false) then v0 else v0, safeModuloInt(v0, v0)), v0;
				var v37 := DC14(v33);
				var v38 := DC16(v37);
				var v39: seq<D7> := [v37];
				var v40 := DC16(v39[safeIndex(|{f9, f8}|, |v39|)]);
				v40 := v40;
				if (false) {
					v15[safeIndex(72, v15.Length)] := v0 * v0;
					var v41: map<bool, array<int>> := map[v15[safeIndex(72, v15.Length)] <= v0 := cf17];
					v41 := v41[p2 := cf17];
					v15[safeIndex(72, v15.Length)] := v15[safeIndex(72, v15.Length)];
					f8 := f9;
					var v42: seq<array<int>> := [cf17];
					var v43: map<string, array<int>> := map[v36 := cf17];
					var v44: array<array<int>> := new array<int>[18] [cf17, cf17, cf17, cf17, cf17, cf17, cf17, cf17, cf17, cf17, v42[safeIndex(v15[safeIndex(72, v15.Length)], |v42|)], cf17, cf17, cf17, if (v36 in v43) then v43[v36] else cf17, cf17, cf17, cf17];
					var v45 := new C0(v44);
				} else {
					v0 := v15[safeIndex(72, v15.Length)];
					var v46: seq<string> := ["u"];
					var v47 := DC1(f8, v35, p2);
					f8 := fm1(v36 + v46[safeIndex(0x238, |v46|)], v47.(cf2 := f9), globalState);
					v31 := v31;
					var v48: array<bool> := new bool[22](i3 => f8);
					v48[safeIndex(805, v48.Length)] := f8;
					v48[safeIndex(805, v48.Length)] := f9;
					var v49: seq<D4> := [DC7(v36)];
					var v50: set<seq<D4>> := {v49, v49};
					var v51: T2 := new C1(v2, v50, false, !p2);
					var v52: array<T2> := new T2[14] [v51, v51, v51, v51, v51, v51, v51, v51, v51, v51, v51, v51, v51, v51];
					var v53: map<string, array<T2>> := map[v36 := v52];
					var v54: map<bool, string> := map[p2 := v36[safeIndex(v15[safeIndex(72, v15.Length)], |v36|) := v2]];
					v53 := v53[(if (p2 in v54) then v54[p2] else v36) + v36 := v52];
				}
				
		}
		
	}
	method m1(p0: string, p1: bool, globalState: GlobalState) returns (r0: string, r1: int, r2: bool) {
		var v0: array<int> := new int[16];
		var v1 := 345;
		v0[safeIndex(790, v0.Length)] := v1;
		v0[safeIndex(790, v0.Length)] := fm6(v1, v1 + v1, globalState);
		var v2 := DC10(v0);
		match (v2) {
			case DC11(cf18, cf19) =>
				v0[safeIndex(790, v0.Length)] := v1;
				r1 := 0x24e;
				var v3: array<bool> := new bool[13];
				v3[safeIndex(487, v3.Length)] := if (f8) then f8 else p1;
				v3[safeIndex(487, v3.Length)] := f9;
				v0 := v0;
			case DC10(cf17) =>
				var v4 := 'o';
				var v5 := DC7(p0);
				var v6: seq<D4> := [DC7(p0), v5, v5, v5, v5];
				var v7: array<array<int>> := new array<int>[28] [cf17, v0, v0, cf17, cf17, v0, cf17, v0, v0, cf17, cf17, cf17, cf17, v0, cf17, v0, cf17, v0, cf17, cf17, v0, cf17, v0, cf17, cf17, v0, v0, cf17];
				var v8: C0 := new C0(v7);
				var v9: map<C0, bool> := map[v8 := f8];
				var v10: set<seq<D4>> := {[v5], v6, fm16(false, v1, |v9|, globalState), v6};
				var v11 := new C1(v4, v10, f9 && f8, (seq(abs(-0x3cc), i0  => (v4))) <= p0);
				f8 := true;
				cf17 := cf17;
				r1 := v0[safeIndex(790, v0.Length)];
		}
		
		var v12: seq<array<int>> := [v0, v0, v0, v0, v0];
		var v13: map<bool, array<int>> := map[f9 := v0];
		var v14: array<array<int>> := new array<int>[14] [v0, v0, v0, v0, v0, v0, v12[safeIndex(v1, |v12|)], v0, v0, v0, v0, if (f8 in v13) then v13[f8] else v2.cf17, v0, v0];
		var v15: C0 := new C0(v14);
		var v16: multiset<bool> := multiset{p1};
		var v17: seq<int> := [v0[safeIndex(790, v0.Length)]];
		var v18: set<bool> := {f8, f9, f9, f8, p1};
		var v19 := DC14(v18);
		var v20: map<int, bool> := map[v0[safeIndex(790, v0.Length)] := p1];
		var v21: seq<bool> := [f9];
		var v22 := DC1(p1, v21, v21[safeIndex(v0[safeIndex(790, v0.Length)], |v21|)]);
		var v23: array<bool> := new bool[27] [v15.fm0(globalState), p1, f9, p1, true, f8, f8, p1, p1, p1, f9, f9, f9, f9, p1, f9, f9, f8, false, f9, f9, f8, p1, p1, f9, p1, p1];
		var v24: seq<array<bool>> := [v23, v23];
		var v25: map<int, int> := map[v0[safeIndex(790, v0.Length)] := v0[safeIndex(790, v0.Length)]];
		var v26: array<bool> := new bool[26] [v1 < (if (f8 in v16) then v16[f8] else v0[safeIndex(790, v0.Length)]), f8, fm0(globalState), f9, 0x22e <= v1, f9, f9, |v17| >= v0[safeIndex(790, v0.Length)], f8, {p1, f9} !! v19.cf25, f8, f8, f8, p1, false, p1, !(|v20| > v1), v15.fm1(p0, v22, globalState), f8, !(|[|v24|, v0[safeIndex(790, v0.Length)]]| in v25), f9, p1, p1, f8 ==> f9, f8, p1];
		v23[safeIndex(815, v23.Length)] := p1;
		var v27: set<int> := {v1};
		var v28 := DC17(v27);
		r2, v15, v1, v23[safeIndex(815, v23.Length)], r2 := (v28.cf27 - v27) > {0x78}, v15, -v0[safeIndex(790, v0.Length)], p1, !(|v17| >= safeModuloInt(v1, v0[safeIndex(790, v0.Length)]));
		var v29 := DC4(f9);
		var i1 := 0;
		while (match v29.(cf5 := !v23[safeIndex(815, v23.Length)]) {
			case DC4(cf5) => if (v23[safeIndex(815, v23.Length)]) then DC19(f9).cf31 else DC13(v23[safeIndex(815, v23.Length)], f9, seq(abs(-939), i2  => ('u')), cf5).cf21
			case DC3(cf4) => if (!v23[safeIndex(815, v23.Length)]) then p1 else p1
		})
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			v0[safeIndex(790, v0.Length)] := -0x2fd;
			f8 := p1;
			var v30 := 'm';
			var v31 := DC7(seq(abs(0x377), i4  => (v30)));
			var v32: set<seq<D4>> := {(seq(abs(854), i3  => (DC7(p0))))[safeIndex(v0[safeIndex(790, v0.Length)], |(seq(abs(854), i3  => (DC7(p0))))|) := v31]};
			var v33: C1 := new C1(v30, v32 - v32, v23[safeIndex(815, v23.Length)], f9);
			v33 := v33;
			v30 := if (p1 ==> v23[safeIndex(815, v23.Length)]) then fm7(v0[safeIndex(790, v0.Length)], f9, |(seq(abs(0x3c1), i5  => (v30)))|, globalState) else if (true) then v30 else fm7(v0[safeIndex(790, v0.Length)], v23[safeIndex(815, v23.Length)], |v27|, globalState);
		}
		var i6 := 0;
		while (!p1)
			decreases 100 - i6
		{
			if (i6 >= 100) {
				break;
			}
			
			i6 := i6 + 1;
			var v34: seq<seq<bool>> := [v21];
			var v35 := DC7(p0);
			var v36: map<seq<seq<bool>>, D4> := map[v34 := v35];
			v36 := v36[v34 := fm17(f9, f8, v0[safeIndex(790, v0.Length)], globalState)];
			var v37: array<array<bool>> := new array<bool>[26];
			v37[safeIndex(925, v37.Length)] := v26;
			var v38: array<string> := new string[25];
			v38[safeIndex(813, v38.Length)] := p0 + "iudnfs";
			r2, v37[safeIndex(925, v37.Length)], v38[safeIndex(813, v38.Length)], r0 := v23[safeIndex(815, v23.Length)], v23, p0, p0;
			if (if (if (p1) then true else v23[safeIndex(815, v23.Length)]) then v23[safeIndex(815, v23.Length)] else v1 != v0[safeIndex(790, v0.Length)]) {
				v0[safeIndex(790, v0.Length)] := v0[safeIndex(790, v0.Length)] - v0[safeIndex(790, v0.Length)];
				v23[safeIndex(815, v23.Length)] := fm0(globalState);
				var v39: set<string> := {p0};
				v39 := v39;
				r1 := safeModuloInt(-safeModuloInt(|v17|, |v18|), 0xf);
				var v40 := new C0(v14);
			} else {
				f8 := v18 >= (if (f9) then v18 else v18);
				var v41: set<seq<int>> := {v17, v17 + (seq(abs(0x118), i7  => (v1))), [v0[safeIndex(790, v0.Length)], v0[safeIndex(790, v0.Length)], v1]};
				var v42 := 'h';
				var v43 := DC2(v42);
				var v44: array<char> := new char[29] ['j', if (f9) then v42 else v43.cf3, v42, v42, v42, v42, v42, v42, v42, 'n', v42, v38[safeIndex(813, v38.Length)][safeIndex(0xf0, |v38[safeIndex(813, v38.Length)]|)], v42, v42, v42, v42, v42, v42, fm7(fm6(v1, |fm18(f8, v0[safeIndex(790, v0.Length)], v0[safeIndex(790, v0.Length)], globalState)|, globalState), fm1(v38[safeIndex(813, v38.Length)], v22, globalState), v1, globalState), v42, v42, fm7(v0[safeIndex(790, v0.Length)], p1, fm2(DC1(false, [v23[safeIndex(815, v23.Length)]], f8), globalState), globalState), DC2(v42).cf3, v42, v43.cf3, v42, 'q', fm7(254, p1, -v1, globalState), if (p1) then v42 else v42];
				v44[safeIndex(878, v44.Length)] := v42;
				var v45 := DC19(p1);
				var v46: array<D6> := new D6[10];
				var v47: map<D8, array<D6>> := map[v45 := v46];
				f8, v41, v44[safeIndex(878, v44.Length)], v47 := !true, ({v17, v17[safeIndex(v0[safeIndex(790, v0.Length)], |v17|) := 0x35a], v17} - v41) - (v41 * v41), 't', v47[v45.(cf31 := f9) := v46];
				var v48 := new C0(v15.f12);
				f8 := v27 <= v27;
				v1 := fm2(v22, globalState);
			}
			
			v0 := v0;
		}
		var v49 := new C0(v15.f12);
		r0 := p0;
		r1 := |v21|;
		r2 := v23[safeIndex(815, v23.Length)];
	}
	method m3(p0: string, p1: int, globalState: GlobalState) returns (r0: int, r1: bool) {
		var v0: map<bool, bool> := map[f8 := f9];
		var v1 := DC11(|v0|, p0);
		var v2 := DC11(p1, v1.cf19[safeIndex(-p1, |v1.cf19|) := fm7(-p1, f8, p1, globalState)]);
		v1 := v2;
		var v3: multiset<multiset<bool>> := multiset{multiset{f9, f8}};
		for i0 := |v3| to if (f9) then p1 else p1 {
			var v4: multiset<seq<char>> := multiset{p0, p0};
			var v5: array<bool> := new bool[9] [multiset([p0, p0, p0, fm9(f8, !f9, globalState)]) !! v4, f9, f9, false || f8, f9, f8, f8, f8, f9];
			v5[safeIndex(706, v5.Length)] := f8;
			v5[safeIndex(706, v5.Length)] := p0 != "kggm";
			var v6: multiset<bool> := multiset{p0 <= p0, f8, f8};
			var v7: seq<bool> := [v5[safeIndex(706, v5.Length)]];
			r0, r0, v5[safeIndex(706, v5.Length)] := i0, |v6|, [v5[safeIndex(706, v5.Length)]] != (v7 + v7);
			var v8 := 'o';
			v8 := v8;
			v5[safeIndex(706, v5.Length)] := f9;
		}
		var v9: multiset<int> := multiset{p1};
		var i1 := 0;
		while (v9[p1 := abs(p1)] >= (v9 - v9))
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			r1 := !f8;
			var v10 := "xkd";
			var v11: seq<int> := [p1];
			var v12 := DC5(v11);
			var v13: array<seq<array<int>>> := new seq<array<int>>[23];
			var v14: array<int> := new int[25](i2 => i2 + |[f9, f8]|);
			var v15: seq<array<int>> := [v14, v14];
			v13[safeIndex(182, v13.Length)] := v15 + v15;
			var v16 := 'g';
			var v17: map<bool, char> := map[f9 := v16];
			v10, v12, v13[safeIndex(182, v13.Length)], r0, f8 := ("rdheeg")[safeIndex(p1, |"rdheeg"|) := if (f8 in v17) then v17[f8] else 'c'], v12, v15, p1, fm0(globalState);
			var v18: array<map<int, int>> := new map<int, int>[7](i3 => map[p1 := 0xd7] + map[p1 := p1]);
			v18 := v18;
			var v19 := DC7(v10);
			var v20: seq<D4> := [v19, v19];
			var v21: set<seq<D4>> := {v20, v20};
			var v22: T2 := new C1(v16, v21, f8, f9);
			var v23: map<T2, int> := map[v22 := p1];
			var v24: seq<map<T2, int>> := [v23 + v23, v23[v22 := p1], map[v22 := p1], v23, v23];
			v23 := v24[safeIndex(|(seq(abs(0x11f), i4  => (p1)))|, |v24|)];
		}
		var v25: array<char> := new char[29](i6 => 'y');
		forall i5 | 0 <= i5 < v25.Length {
			v25[i5] := 'c';
		}
		var v26: array<array<int>> := new array<int>[15];
		var v27: C0 := new C0(v26);
		var v28: map<int, bool> := map[|[v27, v27, v27, v27, v27]| := f9];
		v28 := v28[safeModuloInt(p1, 499) := f9];
		r0 := p1;
		r0 := p1;
		r1 := true;
	}
}

class C3 extends T1, T2 {
	constructor (f8 : bool, f9 : bool) {
		this.f8 := f8;
		this.f9 := f9;
	}
	
	function fm2(p0: D0, globalState: GlobalState): int {
		(-0x2c6 * |{true}|) * 0x1b1
	}
	function fm0(globalState: GlobalState): bool {
		(set v0 : int | v0 in [--88, |[f8, f8]|] :: (safeModuloInt(v0, 0x7))) < {|multiset{f9, f8}|, |map[f8 := |(map v1 : int | (0x39 <= v1) && (v1 < 0x329) :: (v1 * -161) := (f8))|]|, 0x3b0, -0x285}
	}
	function fm1(p0: string, p1: D0, globalState: GlobalState): bool {
		f9
	}
	function fm3(globalState: GlobalState): multiset<int> {
		(multiset([0x3d0, |"ceyl"|, 0x233]) + multiset{|(seq(abs(748), i0  => (0xc9)))|}) - multiset{|[0x1bb, 0xc8, |(set v0 : int | v0 in multiset{0x2ec, 0x8e} :: (v0 + |[DC2('o').cf3]|))|]|}
	}
	method m0(p0: multiset<bool>, p1: D0, p2: bool, globalState: GlobalState) {
		var v0: array<multiset<bool>> := new multiset<bool>[24];
		v0[safeIndex(694, v0.Length)] := p0;
		var v1: seq<multiset<bool>> := [p0, multiset{f8, p2, p2, p2, f9}, p0, p0, p0];
		var v2 := 0x1db;
		v0[safeIndex(694, v0.Length)] := (if (f9) then p0 else multiset{false, f8}) * v1[safeIndex(v2, |v1|)];
		var v3: seq<int> := [v2];
		f8 := !!(|v3| == v2);
		var v4 := "ks";
		var v5 := DC3(v3[safeIndex(v2, |v3|)]);
		v4, v2, f8 := fm4(v2, v2, globalState), -(v5.cf4 * -v2), f8;
		var v6: array<bool> := new bool[7](i0 => false);
		v6 := v6;
		v6[safeIndex(170, v6.Length)] := p2;
		var v7: seq<bool> := [p2];
		v6[safeIndex(170, v6.Length)] := v7[safeIndex(-(v2 + v2), |v7|)];
		v6[safeIndex(170, v6.Length)], v2 := |"fvnhbyxkp"| == v2, v2 + 0x27d;
	}
	method m1(p0: string, p1: bool, globalState: GlobalState) returns (r0: string, r1: int, r2: bool) {
		var v0: T2 := new C2(f9, !p1);
		var v1: map<T2, bool> := map[v0 := !p1];
		var v2: map<bool, bool> := map[v0.f9 := f8];
		var v3 := 995;
		var v4: multiset<int> := multiset{|v2|, v3};
		var v5: seq<multiset<int>> := [v4];
		for i0 := |v1| to |v5| {
			var v6: seq<bool> := [f8, true, false];
			var v7: array<bool> := new bool[9] [f9, f9, false, v0.f9, v6[safeIndex(v3, |v6|)], fm0(globalState), if (v0.f9 in v2) then v2[v0.f9] else p1, p1, true];
			var v8: map<array<bool>, string> := map[v7 := p0];
			v8 := v8[v7 := p0];
			r1 := i0 * i0;
			var v9: seq<int> := [0x36b, v3];
			var v10 := DC5(v9);
			match (v10) {
				case DC6(cf7, cf8, cf9) =>
					var v11: array<int> := new int[23](i1 => i1 + -678);
					v11 := v11;
					var v12 := 'u';
					var v13: map<int, char> := map[0x214 := v12];
					v13 := v13[v3 := v12];
					var v14: map<int, bool> := map[cf8 := !f8];
					var v15 := DC1(if (|p0| in v14) then v14[|p0|] else p1, v6, v0.f9);
					cf8 := fm2(v15, globalState);
					var v16: multiset<bool> := multiset{v0.f9, |(seq(abs(0x11f), i2  => (v2)))| == |p0|};
					v16 := v16;
				case DC5(cf6) =>
					var v17: map<T2, seq<int>> := map[v0 := v9];
					r2 := (cf6 + (if (v0 in v17) then v17[v0] else seq(abs(-723), i3  => (-i0)))) <= v9;
					v3 := safeModuloInt(|(multiset{i0} + v4)|, i0);
					var v18: multiset<bool> := multiset{p1, f8, f8, f9, f8};
					r1 := -DC3(v9[safeIndex(|v18|, |v9|)]).cf4;
					var v19: set<bool> := {true, v0.f9};
					var v20 := DC14(v19);
					var v21 := DC16(v20);
					v21 := v21;
			}
			
			r2 := f9;
		}
		var v22: array<int> := new int[1](i4 => safeDivisionInt(i4, v3));
		var v23: seq<array<int>> := [v22, v22];
		var v24: seq<seq<array<int>>> := [v23, v23];
		v23 := v24[safeIndex(if (f8) then v3 else v3, |v24|)];
		r1 := 0x124;
		r1 := -v3;
		var v25: seq<bool> := [f9, v0.f9];
		var v26 := DC1(f8, v25, p1);
		var v27: array<array<int>> := new array<int>[17] [v22, v22, v22, v22, v22, v22, v22, v22, v22, v22, v22, v22, v22, v22, v22, v22, v22];
		var v28: C0 := new C0(v27);
		var v29: multiset<C0> := multiset{v28};
		var v30 := DC4(p1);
		var v31: multiset<D5> := multiset{DC10(v22)};
		var v32: array<bool> := new bool[11] [f8, if (f9) then f8 else f8, v0.f9, if (fm1("uk", v26, globalState)) then !f9 else fm1(p0, v26, globalState), v0.f9, f8, v29 >= v29, f9, !v30.cf5, f8, |v31| != 0x28d];
		v32[safeIndex(614, v32.Length)] := f8;
		v32[safeIndex(614, v32.Length)] := true;
		f8 := f9;
		r0 := p0;
		var v33: multiset<bool> := multiset{f8};
		r1 := v3 - (if (p1 in v33) then v33[p1] else v3);
		r2 := (f9 <==> v28.fm0(globalState)) <==> (v3 >= 154);
	}
	method m2(p0: bool, globalState: GlobalState) returns (r0: int, r1: bool) {
		var v0 := -0x10;
		for i0 := v0 to v0 {
			var v1: array<bool> := new bool[18];
			v1[safeIndex(297, v1.Length)] := f9;
			v1[safeIndex(297, v1.Length)] := p0;
			var v2 := "ikpfgysl";
			var v3: array<map<int, char>> := new map<int, char>[11];
			var v4: map<string, array<map<int, char>>> := map[v2 := v3];
			v4 := v4[v2 := v3];
			v1[safeIndex(297, v1.Length)] := p0;
			v1[safeIndex(297, v1.Length)] := v1[safeIndex(297, v1.Length)];
		}
		if (f9) {
			var v5: array<int> := new int[27];
			var v6: array<bool> := new bool[21] [p0, !p0, f8, false, f9, f9, f8, p0, f8, false, false, p0, f9, !f9, true, f8, f9, !true, f8, f8, f8];
			var v7: seq<array<bool>> := [v6];
			var v8: map<array<int>, bool> := map[v5 := v6 in v7];
			var v9: set<bool> := {p0, p0};
			var v10 := DC8(v0, f9);
			var v11: seq<D4> := [v10, v10, v10, v10];
			v8 := v8[v5 := |v9| != |v11|];
			var v12: map<int, int> := map[v0 := v0];
			var v13: map<bool, int> := map[f9 := v0];
			var v14: set<int> := {if (v0 in v12) then v12[v0] else |v13|, -0x1b9, v0};
			if ({498} <= v14) {
				v0 := -v0;
				var v15 := "niqye";
				var v16: seq<bool> := [p0, f9];
				var v17 := DC1(f9, v16, true);
				var v18: C2 := new C2(f9 || fm1(v15, v17, globalState), false);
				v18 := v18;
				v5[safeIndex(3, v5.Length)] := v0;
				v5[safeIndex(3, v5.Length)] := v0;
				var v19: seq<int> := [v5[safeIndex(3, v5.Length)]];
				v19 := v19;
				r0 := v5[safeIndex(3, v5.Length)];
			} else {
				v5[safeIndex(68, v5.Length)] := v0;
				v5[safeIndex(68, v5.Length)] := v0;
				var v20 := 'm';
				var v21: seq<char> := [v20, v20];
				var v23 := new C1(v21[safeIndex(-v0, |v21|)], set v22 : seq<D4> | v22 in multiset{seq(abs(0x33), i1  => (DC7("ymfwvwo")))} :: (v22), p0, !(f8 <== f9));
				var v24: array<int> := new int[22](i2 => i2 - v5[safeIndex(68, v5.Length)]);
				var v25: multiset<bool> := multiset{f9, f9, f9, !(if (v24 in v8) then v8[v24] else p0), f8};
				v23.m0(v25, DC0(), if (f9) then p0 else p0, globalState);
				r1 := v0 >= v5[safeIndex(68, v5.Length)];
				var v26: array<array<char>> := new array<char>[4];
				var v27: array<char> := new char[8](i3 => v23.f10);
				v26[safeIndex(556, v26.Length)] := v27;
				var v28: seq<bool> := [f8];
				var v29 := DC1(f9, v28, p0);
				v26[safeIndex(556, v26.Length)] := if (v23.fm1(v21, v29, globalState)) then v27 else v27;
			}
			
			var v30: array<array<array<bool>>> := new array<array<bool>>[14];
			var v31 := DC9(p0, -0x2d7, v6, v0);
			var v32: map<bool, array<bool>> := map[f8 := v6];
			var v33: array<array<bool>> := new array<bool>[29] [v6, v6, v6, v31.cf15, if (f9 in v32) then v32[f9] else v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6, v6];
			v30[safeIndex(47, v30.Length)] := v33;
			var v34: seq<bool> := [p0];
			var v35: multiset<int> := multiset{v0, v0};
			v30[safeIndex(47, v30.Length)], r0, f8 := v33, |v34|, (multiset{v0} * v35) > (v35 + v35);
			r0 := v0;
			r1 := f9;
		} else {
			var v36 := 'h';
			var v37 := "piygtl";
			var v38: map<char, string> := map[v36 := v37];
			v38 := v38[v36 := v37 + v37];
			var v39: seq<int> := [-155];
			var v40: seq<bool> := [p0];
			var v41: map<int, D4> := map[v39[safeIndex(v0, |v39|)] := DC8(|v40|, f8)];
			var v42: multiset<bool> := multiset{p0};
			var v43 := DC8(|v42|, p0);
			v41 := v41[|v40| := v43];
			v0 := v0;
			var v44: array<bool> := new bool[29];
			v44 := v44;
			var v45 := DC2(v36);
			v42 := v42 - fm19(f9, v45, v39[safeIndex(v0, |v39|)], globalState);
		}
		
		var v46: set<int> := {v0, v0};
		var v47 := "ev";
		var v48: seq<set<int>> := [v46, {|v47|, v0}, v46];
		var v49 := DC17(v48[safeIndex(v0, |v48|)]);
		match (v49) {
			case DC18(cf28, cf29, cf30) =>
				var v50: map<int, bool> := map[v0 := f9];
				var v51: C2 := new C2(cf29, cf29);
				var v52: set<string> := {v47, v47};
				var v53: multiset<int> := multiset{cf30};
				var v54 := DC20(v53);
				var v55 := DC20(v54.cf32[cf30 := abs(726)]);
				f8, v50, v51, v51, v52 := !((DC20(v53[|v53| := abs(cf30)]).cf32 + v55.cf32) !! (v53 * v53)), v50 + v50[v0 := f9], if (cf29) then v51 else v51, v51, v52 - v52;
				var v56 := 'c';
				var v57 := DC7(v47);
				var v58: seq<D4> := [v57];
				var v59: map<multiset<int>, seq<D4>> := map[v53 := v58];
				var v60: set<seq<D4>> := {v58, if (v53 in v59) then v59[v53] else [DC7(v47), v57, v57, v57, DC7(v47).(cf10 := seq(abs(754), i4  => (v56)))]};
				var v61 := DC8(v0, true);
				var v62 := new C1(v56, v60, f9, v61.cf12);
				cf30 := v0;
				var v63: array<D8> := new D8[15];
				var v64: array<seq<int>> := new seq<int>[1](i5 => [v0, |map[f8 := |v53|]|, cf30, cf30] + [v0]);
				var v65: seq<C2> := [v51];
				v63, v64, v0, v0, cf30 := if (cf29) then v63 else v63, v64, safeDivisionInt(v0, |(v53 * v53)|), cf30, |v65|;
			case DC19(cf31) =>
				var v66: seq<int> := [|v47|, v0, |v47|, v0, 0x1a];
				var v68: seq<bool> := [p0, true];
				var v70: array<int> := new int[12](i6 => i6 + v0);
				var v71: map<bool, set<int>> := map[!f8 := {v0}];
				var v72: array<set<int>> := new set<int>[23] [(set v67 : int | v67 in v66 :: (safeDivisionInt(v67, DC3(|"ieuojgwcd"|).cf4))) * v46, v46, v46, {|v68|, |v47|}, {v0}, v46, v46, v46, v46 - v46, v48[safeIndex(v0, |v48|)] * v46, v46, set v69 : int | (0x26e <= v69) && (v69 < 777) :: (v69 + DC6(p0, |multiset{p0, f9}|, v0).cf8), v46 + v46, v46, v46, {v0} + {v0, v0, |map[v70 := -v0]|}, v46, if (cf31 in v71) then v71[cf31] else v46, v46 * v46, v46 * {-0x12c}, v46, v46, v46];
				v72[safeIndex(850, v72.Length)] := {v0};
				v72[safeIndex(850, v72.Length)] := {v0 - v0};
				var v73: array<bool> := new bool[5];
				v73 := v73;
				v73 := v73;
				v70[safeIndex(921, v70.Length)] := v0;
				v70[safeIndex(921, v70.Length)] := v0;
			case DC17(cf27) =>
				var v74: array<bool> := new bool[4](i7 => DC8(v0, f9).cf11 != v0);
				v74[safeIndex(59, v74.Length)] := f8;
				var v75: map<bool, bool> := map[fm0(globalState) := f9];
				v74[safeIndex(59, v74.Length)] := if (f9 in v75) then v75[f9] else f8;
				v74[safeIndex(59, v74.Length)] := p0;
				v0, v0 := -v0, v0;
				var v76: array<D9> := new D9[16];
				var v77 := DC7(v47);
				var v78: seq<D4> := [v77, v77.(cf10 := seq(abs(0xdb), i8  => ('i')))];
				var v79: set<seq<D4>> := {v78};
				var v80: T2 := new C1(v47[safeIndex(v0, |v47|)], v79, f8, v74[safeIndex(59, v74.Length)]);
				var v81: map<T2, int> := map[v80 := v0];
				var v82 := DC22(v0, v0, p0, v0, v81);
				v76[safeIndex(159, v76.Length)] := v82;
				v76[safeIndex(159, v76.Length)] := DC22(v0, v0 * v0, p0, |fm20(globalState)|, v81 + v81[v80 := v0]);
		}
		
		var v83 := DC19(f8);
		v0 := match v83 {
			case DC18(cf28, cf29, cf30) => cf30 * |[cf30, cf30]|
			case DC19(cf31) => -|v47|
			case DC17(cf27) => 100
		};
		var v84: map<bool, int> := map[f9 := v0];
		var v85: seq<map<bool, int>> := [v84, fm13(globalState), v84];
		var i9 := 0;
		while (|v85| == v0)
			decreases 100 - i9
		{
			if (i9 >= 100) {
				break;
			}
			
			i9 := i9 + 1;
			var v86: multiset<int> := multiset{-v0, v0};
			var v87: set<bool> := {f8, v86 >= v86, f9, p0};
			var v88 := DC1(p0, [f9, f9], !p0);
			r0, v47, v0, v0, v87 := v0, v47, v0, fm2(if (p0) then v88 else v88.(cf0 := f9), globalState), v87;
			if (true) {
				var v89: seq<bool> := [f8];
				r1 := v89[safeIndex(safeDivisionInt(v0, -v0), |v89|)];
				f8 := f9;
				var v90: array<array<int>> := new array<int>[5];
				v90 := new array<int>[18];
				var v91: map<bool, char> := map[f9 := fm7(v0, f9, fm2(v88, globalState), globalState)];
				var v92: set<map<bool, char>> := {v91};
				var v93: seq<int> := [fm2(v88, globalState)];
				var v94: map<set<map<bool, char>>, int> := map[v92 := if (f9) then |v93| else v0];
				var v95: map<bool, bool> := map[f8 := true];
				v94 := v94[{v91, v91, v91} := |v95|];
				var v96 := new C0(v90);
			} else {
				var v97 := DC7("ki");
				var v98: seq<D4> := [v97, v97, v97.(cf10 := v47), v97, v97];
				var v99 := new C1('q', {v98}, v87 == v87, p0 <== f8);
				var v100: array<array<bool>> := new array<bool>[23];
				var v101: array<bool> := new bool[17] [false, f8, f9, false, false, p0, f9, f8, f9, f8, p0, f8, true, f8, f9, p0, f9];
				var v102: seq<array<bool>> := [v101, v101, v101];
				v100[safeIndex(286, v100.Length)] := v102[safeIndex(v0, |v102|)];
				var v103: seq<int> := [v0];
				var v104: array<int> := new int[9](i10 => i10 * v0);
				var v105: array<array<int>> := new array<int>[2] [v104, v104];
				var v106: C0 := new C0(v105);
				var v107: map<C0, bool> := map[v106 := f9];
				var v108 := DC9(p0, |v107|, v101, v0);
				var v109: map<int, array<bool>> := map[v0 := v108.cf15];
				v100[safeIndex(286, v100.Length)], v101, v103, v47 := if ((v0 + v0) in v109) then v109[v0 + v0] else v101, v101, ((seq(abs(-0x1a8), i11  => (-v0))) + v103) + [|{v0}|], v47;
				var v110 := new C1(v99.f10, v99.f11, p0 && true, !false);
				var v111 := 'f';
				v104[safeIndex(9, v104.Length)] := v99.fm2(v88, globalState);
				v104, v111, v104[safeIndex(9, v104.Length)] := v104, v111, v0 - v0;
				v104[safeIndex(9, v104.Length)] := safeModuloInt(v99.fm2(v88, globalState), v104[safeIndex(9, v104.Length)]);
			}
			
			var v112: array<int> := new int[26];
			v112 := new int[14];
			var v113 := 'e';
			var v114: seq<bool> := [f9, p0, f9];
			r1 := p0 && fm1(v47[safeIndex(v0, |v47|) := v113], DC1(false, v114, f8), globalState);
		}
		var v115: multiset<map<bool, int>> := multiset{v84, v84};
		if (v115 > v115) {
			f8 := !f8;
			var v116: seq<bool> := [p0, f8, f9];
			var v117 := DC1(true, v116, p0);
			r0 := -fm2(v117, globalState);
			var v118: array<int> := new int[21];
			var v119: seq<array<int>> := [v118, v118, v118, v118, v118];
			var v120: array<array<int>> := new array<int>[17] [v118, v118, v118, v118, v118, v118, v118, v118, v118, v118, v118, v118, v118, v119[safeIndex(-v0, |v119|)], v118, v118, v118];
			var v121 := new C0(v120);
			var v122: map<seq<bool>, bool> := map[[f8, f9] := f8];
			v122 := v122;
			var v123 := 'n';
			v0, v123, r1, r0, f8 := fm2(v117, globalState), if (p0) then v123 else v123, f8 ==> f8, if (v0 == v0) then v0 else v0, f9;
		} else {
			var v124: array<int> := new int[1];
			v124[safeIndex(982, v124.Length)] := v0;
			var v125: seq<bool> := [f9, f9];
			var v126 := DC1(f8, v125, f9);
			var v127 := DC1(fm1("k", v126, globalState), v125, p0);
			v124[safeIndex(982, v124.Length)] := fm2(v127.(cf2 := f9), globalState);
			var v128: array<array<int>> := new array<int>[24];
			var v129: C0 := new C0(v128);
			var v130: map<bool, C0> := map[f9 := v129];
			v130 := v130[p0 := v129];
			f8 := f8;
			r0 := v124[safeIndex(982, v124.Length)];
			var v131: set<bool> := {f9};
			v0 := |(v131 + v131)|;
		}
		
		var v132 := 'h';
		r0 := safeDivisionInt(v0, |(seq(abs(0x3c7), i12  => ('c')))[safeIndex(-265, |(seq(abs(0x3c7), i12  => ('c')))|) := v132]|);
		r1 := f8;
	}
}

function fm4(p0: int, p1: int, globalState: GlobalState): string {
	if (|[|[true, false, !false, false]|, -825, --0x157, 0xb9]| > |"piggfikm"|) then "rehg" else "hm" + "clvp"
}
function fm7(p0: int, p1: bool, p2: int, globalState: GlobalState): char {
	if ("bxb" == "p") then 't' else 'i'
}
function fm9(p0: bool, p1: bool, globalState: GlobalState): string {
	"kivrkb" + "fqfafs"
}
function fm11(p0: map<bool, char>, p1: bool, p2: bool, p3: bool, globalState: GlobalState): D5 {
	DC11(|(map[231 := true] + (map v0 : int | v0 in [|map[-0x103 := !true]|, 0x36, 0xcf, 0x349, |{0x306}|] :: (v0 - 0x3e6) := (!!!!false)))|, DC7("ekmhsfs").cf10)
}
function fm12(p0: bool, p1: D3, globalState: GlobalState): seq<string> {
	DC41(["tfnuqeu"]).cf72
}
function fm13(globalState: GlobalState): map<bool, int> {
	map[!true := 0xf0] + (map[true := |"bjdopi"|] + map[false := |{false}|])
}
function fm14(p0: int, p1: string, p2: bool, globalState: GlobalState): seq<int> {
	[safeDivisionInt(|map[false := false]|, |map['d' := |{0x64}|]|), -881, 0x1d9, |{true}|]
}
function fm15(globalState: GlobalState): D2 {
	DC4(true)
}
function fm16(p0: bool, p1: int, p2: int, globalState: GlobalState): seq<D4> {
	[if (true) then DC7("gwsfpeakx") else DC7(seq(abs(0x239), i0  => ('n')))]
}
function fm17(p0: bool, p1: bool, p2: int, globalState: GlobalState): D4 {
	DC7("ysmgae")
}
function fm18(p0: bool, p1: int, p2: int, globalState: GlobalState): multiset<int> {
	multiset{|"yriwev"|, |map[[0x38c] := !true]|, 640} * multiset{0xb6, |[true, !!true, false, true]|}
}
function fm19(p0: bool, p1: D1, p2: int, globalState: GlobalState): multiset<bool> {
	(multiset{false, true} - multiset{true}) - multiset{true, !false, !true, false, false}
}
function fm20(globalState: GlobalState): map<int, bool> {
	map v0 : int | (0x24d <= v0) && (v0 < 0x35b) :: (v0 * |"nqxdyp"|) := (true)
}
function fm21(p0: int, globalState: GlobalState): map<map<int, int>, int> {
	map v0 : map<int, int> | v0 in {map[|[|[false]|, -0x372]| := |[true]|]} :: (v0) := (DC6(!true, 0x233, 0x372).cf8)
}
function fm22(p0: int, p1: int, p2: int, p3: bool, globalState: GlobalState): map<bool, bool> {
	(map[false := false] + DC24(map[true := true]).cf42) + (map[false := false] + map[true := false])
}
function fm23(p0: bool, p1: set<string>, p2: int, globalState: GlobalState): set<int> {
	if (!!!!true) then (set v0 : int | v0 in multiset{|map[false := |{152}|]|, 0x72, |(seq(abs(0x142), i0  => ('i')))|} :: (v0 - |{!true}|)) * (set v1 : int | v1 in [|"chjhmkgm"|, |map[true := false]|, 0x36b] :: (v1 + -0x31b)) else {|(set v2 : int | (0x377 <= v2) && (v2 < 374) :: (v2 * |multiset([|multiset{false}|])|))|}
}
function fm24(globalState: GlobalState): D7 {
	DC16(DC15())
}
function fm25(p0: int, p1: int, p2: int, globalState: GlobalState): D12 {
	DC30(0x17a, |multiset([true] + [true])|)
}
function fm26(p0: map<bool, int>, p1: int, globalState: GlobalState): set<bool> {
	{if (!true) then true else false}
}
function fm27(globalState: GlobalState): seq<bool> {
	match DC18(map[|map[-|map[!!true := 798]| := 95]| := -|multiset{|[0x2be]|}|], true, 138) {
		case DC18(cf28, cf29, cf30) => [cf29] + [cf29, cf29, cf29]
		case DC19(cf31) => [cf31] + [cf31, cf31]
		case DC17(cf27) => [false, true, !true] + [true]
	}
}
method m4(p0: int, p1: bool, p2: string, globalState: GlobalState) returns (r0: int, r1: bool) {
	var v0: multiset<int> := multiset{p0};
	var v1: map<bool, bool> := map[p1 := p1];
	var v2: T1 := new C3(if (p1 in v1) then v1[p1] else p1, p1);
	var v3: set<T1> := {v2};
	var v4: seq<bool> := [v2.f8];
	var v5: seq<int> := [|v4|];
	var v6 := DC1(p1, [p1, v4[safeIndex(-p0, |v4|)]], false);
	var v7: T2 := new C3(p1, p1);
	var v8: map<T2, int> := map[v7 := p0];
	var v9 := DC22(p0, |v5|, p1, -v2.fm2(v6, globalState), v8);
	var v10: multiset<int> := multiset{if (|v3| in v0) then v0[|v3|] else -700, v5[safeIndex(v9.cf35, |v5|)]};
	var v11: array<set<int>> := new set<int>[1];
	var v12: set<int> := {p0};
	var v13: seq<set<int>> := [v12, v12, v12, v12];
	v11[safeIndex(922, v11.Length)] := v13[safeIndex(0x221, |v13|)];
	var v14: map<bool, int> := map[v7.f9 := p0];
	var v15 := DC11(|v14|, p2);
	var v16 := 'm';
	var v17: array<int> := new int[20] [p0, p0, 0x304 + p0, p0, |"en"|, -(p0 - p0), p0, p0, p0 + p0, p0, 0x8d, v2.fm2(v6, globalState), v5[safeIndex((v15.(cf19 := (seq(abs(-0x254), i0  => ('d')))[safeIndex(p0, |(seq(abs(-0x254), i0  => ('d')))|) := v16])).cf18, |v5|)], p0, p0, p0, p0 + p0, p0, p0, p0 * p0];
	v0, v11[safeIndex(922, v11.Length)], v17, r0 := multiset{p0}, v12 * v12, v17, safeModuloInt(v5[safeIndex(p0, |v5|)], |([v7.f9] + v4)|);
	forall i1 | 0 <= i1 < v17.Length {
		v17[i1] := i1 * p0;
	}
	v2.f8 := match v6 {
		case DC0() => v7.f9
		case DC1(cf0, cf1, cf2) => !false
	};
	for i2 := p0 to p0 * p0 {
		var v18, v19, v20 := v7.m1(p2 + p2, multiset([v2.fm0(globalState), v2.f8]) >= multiset{v2.f8, p1}, globalState);
		var v21: map<string, bool> := map[v18 := v2.f8];
		r1 := |v21| > p0;
		var v22 := DC7(fm9(v7.f9, v2.f8, globalState));
		var v23: seq<D4> := [v22, v22, v22, v22, v22];
		var v24 := new C1(v16, {v23, v23}, i2 in v5, v2.f8);
		var v25: C3 := new C3(v20, v2.f8);
		v25 := new C3(v2.f8 ==> v20, if (v7.f9) then v20 else v2.f8);
	}
	var v26: array<set<bool>> := new set<bool>[2](i4 => {v2.f8, !v2.f8, v7.f9});
	forall i3 | 0 <= i3 < v26.Length {
		v26[i3] := {v7.f9, v2.f8};
	}
	for i5 := p0 * p0 to -(p0 + p0) {
		r0 := 672 * |p2|;
		var v27: multiset<bool> := multiset{v2.f8, true};
		var v28: map<int, multiset<bool>> := map[p0 := v27];
		r0 := safeModuloInt(|v28|, v2.fm2(v6, globalState));
		var v29 := DC10(v17);
		var v30: multiset<array<int>> := multiset{v29.cf17};
		v30 := v30 + (multiset{v17, v17} * v30);
		var v31: array<bool> := new bool[17](i6 => !v2.f8);
		v31[safeIndex(989, v31.Length)] := v7.f9;
		v31[safeIndex(989, v31.Length)] := v2.f8;
	}
	r0 := p0;
	var v32 := DC30(p0, p0);
	var v33: map<D12, string> := map[DC32(v32) := p2];
	var v34: set<string> := {if (DC32(v32) in v33) then v33[DC32(v32)] else p2};
	r1 := if (v2.f8) then v2.f8 else {"kl"} !! v34;
}
method Main() {
	var v0 := "xl";
	var v1: array<bool> := new bool[18](i0 => true);
	var globalState := new GlobalState('i', 0x72, -0x3e1, 0x1ef, v0, -0x268, v1, -0x339);
	var v2 := false;
	var v3: seq<bool> := [v2, v2, v2, v2];
	match (match DC1(v2, v3, v2).(cf1 := v3) {
		case DC0() => DC0()
		case DC1(cf0, cf1, cf2) => DC0()
	}) {
		case DC0() =>
			if (v2) {
				var v4 := new C3(v2, v2);
				var v5 := 'p';
				v5, v2, v3, v5 := 'c', v2, v3, v5;
				var v6: array<string> := new string[6] [v0 + v0, "fblhxtr", v0, v0 + v0, v0, "whybwpehq"];
				var v7: map<bool, string> := map[v2 := v0];
				var v8 := 873;
				var v9: seq<int> := [v8];
				var v10: map<int, string> := map[|v9| := v0];
				v6[safeIndex(598, v6.Length)] := if (v2 in v7) then v7[v2] else if (v8 in v10) then v10[v8] else seq(abs(621), i1  => (v5));
				v6[safeIndex(598, v6.Length)] := v0;
				var v11 := DC21(v8, v8);
				var v12: map<D9, int> := map[v11 := v8];
				var v13 := DC1(v2, v3, v2);
				var v14: array<array<int>> := new array<int>[11];
				var v15: T0 := new C0(v14);
				var v16: map<int, T0> := map[|v0| := v15];
				var v17: seq<map<int, T0>> := [v16];
				var v18: set<array<bool>> := {v1, v1};
				var v19: map<bool, int> := map[v2 := v8];
				var v20: map<int, int> := map[v8 := v4.fm2(v13, globalState)];
				var v21: map<int, bool> := map[v8 := true];
				var v22: map<bool, bool> := map[v2 := !v2];
				var v23: multiset<map<bool, bool>> := multiset{v22, v22};
				var v24 := DC11(|v23|, v6[safeIndex(598, v6.Length)]);
				var v25: array<char> := new char[4](i2 => v5);
				var v26: map<array<char>, bool> := map[v25 := v2];
				var v27: array<int> := new int[29] [v8, v8, v8, v8, v8, v8, if (v11.(cf33 := v8) in v12) then v12[v11.(cf33 := v8)] else v8, safeDivisionInt(v4.fm2(v13, globalState), v8), safeDivisionInt(v8, v4.fm2(DC1(v2, v3, v2), globalState)), |(v17 + v17)|, -|v18|, v8, if (v2) then v8 else v8, v8, v8, |fm21(|v19|, globalState)|, v8 + v8, v4.fm2(v13, globalState), |v9[safeIndex(v8, |v9|) := v8]|, v8, if (|v0| in v20) then v20[|v0|] else v8, -(|v21| + |v3|), -0x278 - |map[v2 := v8]|, -0x12e, v8, v8, if (v2) then v24.cf18 else |v26|, -(v8 * v8), v8];
				v27[safeIndex(815, v27.Length)] := -v8;
				var v28: multiset<int> := multiset{v4.fm2(v13, globalState), 0x14e, v8, safeModuloInt(v8, v8), v8};
				v27[safeIndex(815, v27.Length)] := if (v4.fm2(DC1(v2, [v2], v2), globalState) in v28) then v28[v4.fm2(DC1(v2, [v2], v2), globalState)] else -0x153;
				v1[safeIndex(897, v1.Length)] := v2;
				var v29: multiset<bool> := multiset{v15.fm0(globalState)};
				v1[safeIndex(897, v1.Length)] := v27[safeIndex(815, v27.Length)] >= (if (v2 in v29) then v29[v2] else v8);
			} else {
				var v30 := -507;
				var v31: seq<int> := [-v30];
				var v32 := 'e';
				v0 := v0[safeIndex(v31[safeIndex(0x234, |v31|)], |v0|) := v32];
				var v33: map<bool, array<bool>> := map[v2 := v1];
				v1 := if (v2 in v33) then v33[v2] else v1;
				var v34: set<string> := {seq(abs(0x214), i3  => (v32)), v0};
				v34 := v34;
				var v35: map<bool, int> := map[v2 := |v0|];
				v35 := v35[v2 := v30];
				v30 := v30 * v30;
			}
			
			var v36 := 115;
			v36 := v36;
			var v37: map<int, bool> := map[0x2c := v2];
			v37 := v37[v36 := if (v2) then false else v2];
			var v38 := 'c';
			var v40 := DC7(v0);
			var v41: seq<D4> := [v40, v40, v40, v40, v40];
			var v42: set<seq<D4>> := {v41};
			var v43: C1 := new C1(v38, (set v39 : seq<D4> | v39 in (seq(abs(-0x17f), i4  => ([DC7(seq(abs(413), i5  => (v38)))]))) :: (v39)) + v42, v2, v2);
			v43 := v43;
		case DC1(cf0, cf1, cf2) =>
			var v44 := 0x3b4;
			cf0 := (|v0| <= v44) && cf2;
			var v45 := DC8(safeModuloInt(-v44, v44), cf2);
			match (v45) {
				case DC8(cf11, cf12) =>
					var v46: map<bool, bool> := map[cf2 := cf2];
					var v47: seq<map<bool, bool>> := [v46];
					var v48: multiset<bool> := multiset{cf0};
					var v49 := DC6(cf0, cf11, |v48|);
					var v50: array<map<bool, bool>> := new map<bool, bool>[13] [v46 + v46, v46 + v46, v46, v47[safeIndex(v44, |v47|)], map[v2 := v49.cf7], v46 + v46, v46, v46[v2 := cf2] + map[v2 := !cf0], v46, v46 + v46, v46 + v46, fm22(cf11, |v0|, 0xf0, false, globalState), v46[cf12 := cf12]];
					v50[safeIndex(793, v50.Length)] := v46;
					var v51: array<int> := new int[27];
					var v52 := DC24(v46);
					v50[safeIndex(793, v50.Length)], v51 := v52.cf42, v51;
					var v53: seq<seq<bool>> := [v3, v3, [cf0], v3, [true]];
					var v54 := DC26(v53);
					v44 := safeModuloInt(safeDivisionInt(v44, -cf11), |v54.cf44|);
					var v55: set<int> := {|v0|, |v3|};
					v44 := cf11 - -|map[!false := v55]|;
					v51[safeIndex(66, v51.Length)] := v44;
					v51[safeIndex(66, v51.Length)] := (if (cf2) then v44 else v44) + v44;
				case DC9(cf13, cf14, cf15, cf16) =>
					cf13 := cf13;
					var v56: set<bool> := {cf0, v2};
					var v57: multiset<int> := multiset{|v56|};
					var v58: multiset<multiset<int>> := multiset{v57};
					var v59: set<int> := {|v58|};
					v1[safeIndex(685, v1.Length)] := if (cf13) then v2 else false;
					var v60 := DC3(cf14);
					var v61: C2 := new C2(cf16 == v60.cf4, v2);
					v59, v2, v1[safeIndex(685, v1.Length)], cf14, v61 := fm23(cf13, {v0}, cf16, globalState), |v0| >= v44, cf13, cf14, v61;
					var v62 := new C2(!(v57 !! multiset{|fm14(|(seq(abs(158), i6  => ('i')))|, v0, v2, globalState)|}), v1[safeIndex(685, v1.Length)]);
					var v63 := new C2(if (v2) then cf0 else v62.fm0(globalState), v2);
				case DC7(cf10) =>
					v44 := (v44 - -0x35e) - v44;
					cf2 := (seq(abs(0x152), i7  => ('d'))) <= "hduwoi";
					var v64: map<int, int> := map[v44 := |cf1| + 0x35b];
					v44 := -(if (|cf10| in v64) then v64[|cf10|] else v44);
					var v65: array<array<C3>> := new array<C3>[4];
					var v66: C3 := new C3(v2, cf2);
					var v67 := DC29(v66);
					var v68: array<C3> := new C3[28] [v66, v66, v66, v66, v67.cf49, v66, v66, v66, v66, v66, v66, v66, v66, v66, v66, v66, v66, v66, v66, v66, v66, v66, v66, DC29(v66).cf49, v66, v66, v66, v66];
					v65[safeIndex(675, v65.Length)] := v68;
					v65[safeIndex(675, v65.Length)] := v68;
			}
			
			v44 := (v44 - v44) * safeModuloInt(v44, v44);
			if (if (v2) then cf2 else cf2) {
				var v69: array<array<int>> := new array<int>[16];
				var v70 := new C0(v69);
				var v71 := new C0(v69);
				var v72: array<int> := new int[22];
				v72[safeIndex(223, v72.Length)] := -|v3|;
				var v73: set<int> := {0x1aa};
				var v74: seq<int> := [v44, v44];
				cf2, v72[safeIndex(223, v72.Length)] := v73 >= v73, safeDivisionInt(-v44 + v44, -|v74| - v44);
				var v75, v76 := m4(v44, if (cf2) then cf0 else v2, v0, globalState);
				var v77: set<bool> := {v2};
				cf0 := v77 >= (v77 * v77);
			} else {
				var v78: C3 := new C3(cf0, cf2);
				var v79 := DC29(v78);
				var v80: array<D12> := new D12[19] [v79, v79, v79, if (v2) then DC29(v78) else v79, if (cf2) then v79 else DC29(v78), v79, v79, v79, v79, v79, v79.(cf49 := v78), v79, v79, v79, v79, v79, v79, v79, v79];
				v80 := v80;
				var v81: map<int, D7> := map[v44 := fm24(globalState)];
				var v82: set<bool> := {cf2};
				var v83 := DC14(v82);
				v81 := v81[-v44 := DC16(v83)];
				var v84, v85 := m4(safeDivisionInt(v44, v44), false, v0, globalState);
				var v86: array<multiset<bool>> := new multiset<bool>[2];
				v86 := v86;
				v1[safeIndex(280, v1.Length)] := v85;
				var v87: set<int> := {v84, v44};
				v1[safeIndex(280, v1.Length)] := v87 < (set v88 : int | (-626 <= v88) && (v88 < 0x157) :: (safeModuloInt(v88, v44)));
			}
			
	}
	
	var v89 := -166;
	var v90, v91 := m4(safeModuloInt(0x33b, v89), v2, fm4(v89, 0x262, globalState), globalState);
	var v92: array<array<int>> := new array<int>[11];
	var v93 := new C0(v92);
	var v94: map<bool, int> := map[v2 := v89];
	var v95: multiset<int> := multiset{|v94|};
	var v96: map<bool, int> := map[v2 := |v95|];
	var v97: map<int, int> := map[|(v94 + v94)| := v89 * v90];
	var v98: C3 := new C3(v2, v2);
	var v99 := DC29(v98);
	var v100 := DC32(v99);
	var v101: array<D12> := new D12[8] [v100, v100, v100, v100, v100, DC32(DC29(v98)), v100, DC32(v99)];
	v101[safeIndex(682, v101.Length)] := v100;
	var v102: array<int> := new int[9](i8 => safeDivisionInt(i8, v90));
	v97, v101[safeIndex(682, v101.Length)], v102, v91, v90 := v97, v100, v102, !((if (v91) then v90 else v89) <= v90), -v90;
	var v103: seq<int> := [v90];
	v103 := v103;
	v1[safeIndex(875, v1.Length)] := v0 < "rggpcxlop";
	var v104 := 'q';
	var v105: array<string> := new string[8] [fm4(v90, |(seq(abs(-0x2e1), i9  => ([v91, v2, v91, v91])))|, globalState), v0, if (v91) then v0 else v0, v0, "hw", (if (v91) then v0 else v0)[safeIndex(v89, |(if (v91) then v0 else v0)|) := v104], v0, seq(abs(0x24d), i10  => ('j'))];
	v105[safeIndex(762, v105.Length)] := v0;
	var v106: map<array<bool>, bool> := map[v1 := v2];
	var v107: map<bool, bool> := map[v91 := v2];
	v91, v90, v1[safeIndex(875, v1.Length)], v105[safeIndex(762, v105.Length)] := -(if (v91) then -v89 else v89) >= (if (v89 in v97) then v97[v89] else -|v106|), if (v2) then v89 else (fm25(v90, v89, 0x30d, globalState)).cf51 + v90, match DC24(v107).(cf42 := v107) {
		case DC25(cf43) => false
		case DC24(cf42) => v89 == 0x341
	}, v0;
	var v108 := DC7(v105[safeIndex(762, v105.Length)]);
	var v109: set<seq<D4>> := {[v108, v108, v108, v108]};
	var v110: seq<C3> := [v98];
	var v111 := new C1(v104, v109 + v109, v110 == v110, v91);
	var i11 := 0;
	while (v91)
		decreases 100 - i11
	{
		if (i11 >= 100) {
			break;
		}
		
		i11 := i11 + 1;
		v105[safeIndex(762, v105.Length)] := v105[safeIndex(762, v105.Length)] + v0;
		if (v1[safeIndex(875, v1.Length)]) {
			var v112: multiset<bool> := multiset{!!v2, v1[safeIndex(875, v1.Length)]};
			var v113 := DC0();
			v98.m0(v112, v113, v2, globalState);
			v2 := v2;
			var v114: map<int, char> := map[637 := v111.f10];
			var v115: map<int, map<int, char>> := map[v89 := v114];
			v91 := v89 in v115;
			var v116 := new C0(v93.f12);
			var v117: set<bool> := {v1[safeIndex(875, v1.Length)], v91};
			var v118: seq<set<bool>> := [v117];
			var v119 := DC14(v118[safeIndex(|(seq(abs(0xd6), i12  => (v89)))|, |v118|)]);
			v119 := v119.(cf25 := v117);
		} else {
			var v120 := DC14({v1[safeIndex(875, v1.Length)]});
			var v121 := DC16(v120);
			v121 := DC16(v120);
			var v122: array<D3> := new D3[5];
			var v123 := DC6(v2, 0x347, v90);
			v122[safeIndex(224, v122.Length)] := v123;
			var v124: set<int> := {|(seq(abs(0x343), i13  => ('x')))|, v89};
			var v125 := DC1(!(v124 != v124), v3, v91 || true);
			v122[safeIndex(224, v122.Length)], v125, v104 := v123, v125, v104;
			var v126: set<bool> := {v2, v1[safeIndex(875, v1.Length)]};
			var v127 := DC3(if (|v126| in v95) then v95[|v126|] else v90);
			v89 := v127.cf4;
			v89 := v90;
			v89 := |(fm9(v1[safeIndex(875, v1.Length)], v1[safeIndex(875, v1.Length)], globalState) + fm4(v90, 0x399, globalState))|;
		}
		
		var v128: map<int, multiset<int>> := map[|(seq(abs(0xf7), i14  => (v90)))| := v95];
		var v129 := DC27(true, v128);
		v91 := !!v129.cf45;
		var v130: multiset<C3> := multiset{v98, v98, v98, v98};
		var v131 := DC30(885, -|v130|);
		v94 := v94[v91 <==> v1[safeIndex(875, v1.Length)] := v131.cf51 + v89];
	}
	for i15 := 0x3a1 to -|v0| + 0xcb {
		var v133: set<int> := {|v0|, |(map v132 : int | (0x364 <= v132) && (v132 < 0x267) :: (safeDivisionInt(v132, v89)) := (|multiset([{|v3|}])|))|};
		if (v133 > {v90}) {
			var v134: T1 := new C2(v1[safeIndex(875, v1.Length)], v2);
			var v135: array<T1> := new T1[4] [v134, v134, v134, v134];
			v135[safeIndex(539, v135.Length)] := v134;
			var v136: map<bool, T1> := map[v2 := v134];
			v135[safeIndex(539, v135.Length)] := if (false in v136) then v136[false] else v134;
			v104 := v111.f10;
			v90 := v89 - safeDivisionInt(i15, i15);
			v1[safeIndex(875, v1.Length)] := v93.fm0(globalState);
			var v137, v138 := v98.m2(v134.f8, globalState);
		} else {
			v94 := v94[v1[safeIndex(875, v1.Length)] := i15];
			v104 := 'g';
			v105[safeIndex(762, v105.Length)] := "cubgc";
			var v139: array<multiset<bool>> := new multiset<bool>[24](i16 => multiset{!v2} * multiset(v3));
			var v140: set<bool> := {v91};
			var v141: map<set<int>, set<bool>> := map[{|v97|} := v140];
			v0, v90, v139 := "ovh", |v141[v133 := v140 + v140]|, v139;
			var v142 := new C0(v93.f12);
		}
		
		var v143 := DC1(v2, v3, false);
		var v144, v145 := v98.m2(v98.fm1(v105[safeIndex(762, v105.Length)], v143, globalState), globalState);
		v145 := v2;
		v145 := v144 >= v90;
	}
	v2 := v1[safeIndex(875, v1.Length)];
	match (DC27(v2, map v146 : int | (421 <= v146) && (v146 < 465) :: (v146 - v90) := (v95))) {
		case DC27(cf45, cf46) =>
			if (v2) {
				v90 := v89;
				v108 := v108;
				v98 := v98;
				v89 := v90;
				var v147: C2 := new C2(v91, v2);
				var v148: multiset<string> := multiset{v105[safeIndex(762, v105.Length)]};
				var v149 := DC1(v2, v3, v1[safeIndex(875, v1.Length)]);
				var v150: map<char, int> := map[v111.f10 := v89];
				v1[safeIndex(875, v1.Length)], v147, v89, v91, v89 := v148 > v148, v147, v111.fm2(v149, globalState), v3[safeIndex(if (v111.f10 in v150) then v150[v111.f10] else v90, |v3|)], v147.fm2(v149, globalState);
			} else {
				var v151: seq<array<int>> := [v102, v102, v102, v102, v102];
				cf45 := v151 < v151;
				v2 := !v91;
				v102[safeIndex(83, v102.Length)] := |(multiset(v103) + v95)|;
				var v152: seq<D4> := [v108];
				var v153: T2 := new C1(v104, {v152, v152, v152}, true !in v3, v2);
				var v154: seq<T2> := [v153];
				v102[safeIndex(83, v102.Length)], v153, v1[safeIndex(875, v1.Length)] := 296, if (v91) then v154[safeIndex(|"xj"|, |v154|)] else v153, v1[safeIndex(875, v1.Length)];
				var v155: multiset<char> := multiset{v104, v111.f10};
				v1[safeIndex(875, v1.Length)] := v155 > v155;
				var v156, v157 := m4(v89, cf45, v105[safeIndex(762, v105.Length)], globalState);
			}
			
			v1[safeIndex(875, v1.Length)] := v111.fm0(globalState);
			var v158 := DC2(v104);
			var v159: map<int, char> := map[v89 := v158.cf3];
			v89, v89 := v89, safeModuloInt(safeDivisionInt(v89, |v159|), v90 + v89);
			if (v91) {
				var v160 := new C0(v93.f12);
				var v161 := new C1(v111.f10, v111.f11, v1[safeIndex(875, v1.Length)], cf45);
				v1[safeIndex(875, v1.Length)] := if (v91) then v2 else v1[safeIndex(875, v1.Length)];
				v1[safeIndex(875, v1.Length)] := !v91;
				v2 := false;
			} else {
				var v162, v163 := v98.m2(!(v1[safeIndex(875, v1.Length)] <==> v1[safeIndex(875, v1.Length)]), globalState);
				v2 := true;
				var v164: multiset<bool> := multiset{v163, v91};
				var v165 := DC0();
				v111.m0(v164, v165, if (v2) then v163 else cf45, globalState);
				var v166: array<D8> := new D8[20](i17 => DC19(v1[safeIndex(875, v1.Length)]));
				var v167 := DC19(cf45);
				v166[safeIndex(842, v166.Length)] := v167;
				v166[safeIndex(842, v166.Length)] := v167.(cf31 := v91);
				v89 := |v105[safeIndex(762, v105.Length)]|;
			}
			
		case DC28(cf47, cf48) =>
			var v168: array<array<bool>> := new array<bool>[19];
			var v169: array<array<array<bool>>> := new array<array<bool>>[11] [if (v91) then v168 else v168, v168, v168, v168, v168, v168, v168, v168, if (v91) then v168 else v168, v168, v168];
			v169[safeIndex(367, v169.Length)] := v168;
			var v170 := DC33(v168);
			v169[safeIndex(367, v169.Length)] := v170.cf56;
			var v171: seq<seq<bool>> := [[v91]];
			var v172 := DC26(v171);
			var v173 := new C2(DC34(v172, v105[safeIndex(762, v105.Length)], v1[safeIndex(875, v1.Length)]).cf59, v89 > v90);
			v97 := v97[cf47 := -v90];
			var v174 := DC1(v2, v3, v91);
			var v175 := new C2(v2, v3[safeIndex(v98.fm2(v174, globalState), |v3|)]);
		case DC26(cf44) =>
			v89, v89, v89 := -v90, v89, 373;
			var v176: set<int> := {v89 - -v90};
			var v177: T2 := new C3(v2, v91);
			var v178 := DC22(|v94|, v90, v2, v89, map[DC36(v177).cf61 := v89]);
			var v179: seq<D9> := [v178, v178, v178.(cf36 := v90), v178];
			var v180: set<string> := {v105[safeIndex(762, v105.Length)]};
			var v182: map<T2, bool> := map[v177 := v2];
			v176, v179, v89, v105[safeIndex(762, v105.Length)] := fm23(v2, set v181 : string | v181 in v180 :: (v181), v89 - |v182|, globalState), v179, safeDivisionInt(|v3|, -v89), "qwk" + v105[safeIndex(762, v105.Length)];
			var v183: C2 := new C2(v1[safeIndex(875, v1.Length)], v2);
			var v184: multiset<C2> := multiset{v183};
			var v185 := DC1(v184 <= v184, v3, v1[safeIndex(875, v1.Length)]);
			var v186: multiset<char> := multiset{v104};
			var v187: seq<C2> := [v183, v183, v183, v183, v183];
			v1[safeIndex(875, v1.Length)], v91, v185, v90, v183 := true, (v1[safeIndex(875, v1.Length)] <== v177.f9) ==> (v186 < multiset{fm7(v89, v177.f9, v90, globalState)}), v185.(cf0 := v2), v89, v187[safeIndex(v89, |v187|)];
			var v188: array<array<seq<char>>> := new array<seq<char>>[16];
			v188[safeIndex(384, v188.Length)] := v105;
			var v189: seq<array<int>> := [v102];
			var v190: map<array<bool>, array<seq<char>>> := map[v1 := v105];
			v188[safeIndex(384, v188.Length)], v1[safeIndex(875, v1.Length)], v91, v189 := if (v1 in v190) then v190[v1] else v105, v2, v1[safeIndex(875, v1.Length)], v189;
	}
	
	v1[safeIndex(875, v1.Length)] := v91;
	var v191 := DC13(v91, true, v105[safeIndex(762, v105.Length)], v91);
	match (v191.(cf24 := v91, cf22 := v1[safeIndex(875, v1.Length)])) {
		case DC13(cf21, cf22, cf23, cf24) =>
			v1[safeIndex(875, v1.Length)], cf22, v89 := !cf21, !v2, v98.fm2(DC1(!v3[safeIndex(v90, |v3|)], v3, v2), globalState);
			var v192: T1 := new C2(v2, cf21);
			v192 := if ((seq(abs(-0x1d1), i18  => (v3))) == [v3, v3]) then v192 else v192;
			v91 := cf22;
			cf22 := v1[safeIndex(875, v1.Length)] <== cf21;
		case DC12(cf20) =>
			v90 := -v90;
			v2 := !!v98.fm0(globalState);
			var v193: array<char> := new char[11];
			var v194: seq<array<char>> := [v193];
			v193 := v194[safeIndex(v90, |v194|)];
			v90 := v90;
	}
	
	var v195 := DC12(v93.f12);
	v195 := v195;
	var v196: multiset<bool> := multiset{v91, v2, v91};
	var v197: seq<array<bool>> := [v1, v1, v1];
	var v198: set<int> := {155, v90, |v196|, 0x1b4, |v197|};
	var v199 := DC17(v198);
	if (match if (v1[safeIndex(875, v1.Length)]) then v199 else v199 {
		case DC18(cf28, cf29, cf30) => v2
		case DC19(cf31) => v95 !! v95
		case DC17(cf27) => !!(|v97| <= |DC24(v107).cf42|)
	}) {
		v105[safeIndex(762, v105.Length)], v90 := v0, |(v103 + v103)| * v89;
		var v200: map<int, bool> := map[-452 := v91];
		var v201: multiset<map<int, bool>> := multiset{v200};
		v89, v90, v103, v90 := v90, v89 + -447, fm14(v90, v105[safeIndex(762, v105.Length)], v2, globalState), -|((if (v91) then v201 else v201) - v201)|;
		var v202: map<C1, bool> := map[v111 := v91];
		v202 := v202[v111 := v1[safeIndex(875, v1.Length)]];
		v1 := v1;
		v89 := -|(fm4(v89, v90, globalState) + ((seq(abs(-0x2f1), i19  => (v111.f10))) + v0))|;
	} else {
		match (DC6(DC13(false, v91, v0, v91).cf22, -|v94|, v90)) {
			case DC6(cf7, cf8, cf9) =>
				var v203 := new C3(v2, v3[safeIndex(|v3|, |v3|)]);
				var v204, v205, v206 := v98.m1(v0, cf7, globalState);
				v2 := v90 >= (cf9 + cf9);
				var v207 := new C2(v206 || true, v206);
			case DC5(cf6) =>
				var v208: T2 := new C3(v1[safeIndex(875, v1.Length)], v91);
				var v209 := DC36(v208);
				var v210: map<int, D14> := map[safeDivisionInt(v89, v89) := v209];
				v210 := v210[if (v1[safeIndex(875, v1.Length)]) then v90 else v90 := v209];
				v102[safeIndex(989, v102.Length)] := |(seq(abs(0x8b), i20  => (v111.f10)))[safeIndex(v90, |(seq(abs(0x8b), i20  => (v111.f10)))|) := v104]|;
				v102[safeIndex(989, v102.Length)] := v89;
				v91 := (v90 * 0x3a1) >= -v89;
				var v211: map<bool, string> := map[v208.f9 := seq(abs(-0x17f), i21  => (v111.f10))];
				v102[safeIndex(989, v102.Length)] := |v211[v111.f10 !in (seq(abs(0x19d), i22  => (v111.f10))) := v105[safeIndex(762, v105.Length)] + v0]|;
		}
		
		var v212, v213, v214 := v111.m1(v105[safeIndex(762, v105.Length)], v2, globalState);
		var v215 := DC31(v2, v1[safeIndex(875, v1.Length)], v90);
		v89, v92, v1[safeIndex(875, v1.Length)], v1[safeIndex(875, v1.Length)] := safeDivisionInt(v213, v90), v92, v215.cf53, v214;
		var v216 := DC0();
		v98.m0(v196, v216, v2, globalState);
		v102[safeIndex(131, v102.Length)] := v90;
		v102[safeIndex(131, v102.Length)] := safeModuloInt(safeDivisionInt(390, v213), v89);
	}
	
	var v217: map<int, multiset<int>> := map[v89 := v95];
	var v218 := DC27(v1[safeIndex(875, v1.Length)], v217);
	match (v218) {
		case DC27(cf45, cf46) =>
			v90 := v90;
			v1[safeIndex(875, v1.Length)] := cf45;
			var v219 := DC31(cf45, cf45, |([v1[safeIndex(875, v1.Length)]] + [v1[safeIndex(875, v1.Length)], v2])|);
			cf45, v90, v219 := v90 == v90, -(v89 - safeModuloInt(v90, |v97|)), v219;
			v107 := v107[v91 && v1[safeIndex(875, v1.Length)] := multiset{-0x3cd} !! multiset{v89}];
		case DC28(cf47, cf48) =>
			v102[safeIndex(395, v102.Length)] := -(v89 * v89);
			v1[safeIndex(875, v1.Length)], v102[safeIndex(395, v102.Length)], v2 := ([v89, cf47] + v103) != v103, if (v2) then v89 else v89, v1[safeIndex(875, v1.Length)];
			var v220 := DC31(v2, v2, 0x270);
			match (v220) {
				case DC30(cf50, cf51) =>
					var v221: C2 := new C2(v2, if (v91 in v107) then v107[v91] else v2);
					var v222: map<C2, array<bool>> := map[v221 := cf48];
					v222 := v222[v221 := cf48];
					v1[safeIndex(875, v1.Length)] := true;
					cf51 := cf51;
					var v223 := new C3(true, v2);
				case DC31(cf52, cf53, cf54) =>
					v1[safeIndex(875, v1.Length)] := (v95 + multiset(v103)) >= multiset{|v105[safeIndex(762, v105.Length)]|};
					var v224, v225 := m4(|[v89]|, v91, "xexs", globalState);
					var v226, v227 := m4(v90, -v89 < v224, v105[safeIndex(762, v105.Length)], globalState);
					v94 := v94[v91 := v224 - -cf47];
				case DC29(cf49) =>
					var v228: T1 := new C3(v2, v2);
					v228 := v228;
					var v229 := DC21(808, -v102[safeIndex(395, v102.Length)]);
					var v230: array<int> := new int[26] [v89, 585, v90 + v102[safeIndex(395, v102.Length)], v90, if (v89 in v97) then v97[v89] else v89, cf47, v89 + v89, |v105[safeIndex(762, v105.Length)]|, v102[safeIndex(395, v102.Length)] + v89, cf47, v90, |v0|, -v90 - v90, v102[safeIndex(395, v102.Length)], -v102[safeIndex(395, v102.Length)], |v3|, v90, cf47, cf47, v229.cf33 * v89, v89 + |(seq(abs(0x1a5), i23  => (v111.f10)))|, v102[safeIndex(395, v102.Length)] - v90, -0x1f7, -v102[safeIndex(395, v102.Length)], cf47, v102[safeIndex(395, v102.Length)]];
					v230, v89 := v230, v90;
					var v231: T2 := new C3(v228.f8, v1[safeIndex(875, v1.Length)]);
					var v232: seq<D4> := [v108];
					v231 := new C1(v111.f10, v111.f11 - {v232, v232}, v1[safeIndex(875, v1.Length)], v231.f9);
					var v233 := DC1(v2, v3, v228.f8);
					v2 := (!v228.f8 || v111.fm1(v105[safeIndex(762, v105.Length)], v233, globalState)) <==> v1[safeIndex(875, v1.Length)];
				case DC32(cf55) =>
					var v234: array<map<array<D9>, int>> := new map<array<D9>, int>[12];
					var v235: T1 := new C3(v1[safeIndex(875, v1.Length)], v91);
					var v236 := DC23(v235, v2);
					var v237: array<D9> := new D9[15] [v236, v236, v236, v236, DC23(v235, !v2), v236.(cf41 := v2), v236, v236, v236, v236, v236, v236, v236, DC23(v235, !v1[safeIndex(875, v1.Length)]), v236];
					var v238: map<int, C3> := map[v90 := v98];
					var v239: map<array<D9>, int> := map[v237 := |v238[|v3| := v98]|];
					v234[safeIndex(962, v234.Length)] := v239;
					v234[safeIndex(962, v234.Length)] := v239 + v239;
					v102[safeIndex(395, v102.Length)] := v103[safeIndex(safeDivisionInt(0x14a, cf47), |v103|)];
					v102[safeIndex(395, v102.Length)] := v102[safeIndex(395, v102.Length)] + -v102[safeIndex(395, v102.Length)];
					var v240 := DC1(v235.f8, v3, v91);
					var v241, v242, v243 := v98.m1(v105[safeIndex(762, v105.Length)], v98.fm1(seq(abs(0x113), i24  => (v111.f10)), v240, globalState), globalState);
			}
			
			var v244 := new C2(!v1[safeIndex(875, v1.Length)] && v91, v2);
			v1[safeIndex(875, v1.Length)] := v2;
		case DC26(cf44) =>
			var v245: T0 := new C0(v92);
			var v246: map<int, bool> := map[|(v0 + v0)| := v1[safeIndex(875, v1.Length)]];
			v91, v245, v1[safeIndex(875, v1.Length)], v90 := v2, v245, if (v89 in v246) then v246[v89] else v91 <==> v91, v90 + |fm26(v94, v90, globalState)|;
			var v247 := DC1(v111.fm0(globalState), fm27(globalState), v1[safeIndex(875, v1.Length)]);
			if (v93.fm1(v0 + v0, v247, globalState)) {
				v1[safeIndex(875, v1.Length)] := v1[safeIndex(875, v1.Length)];
				v90 := 135;
				v90 := v89 * |v105[safeIndex(762, v105.Length)]|;
				var v248: T2 := new C1(v111.f10, v111.f11, v91, true);
				var v249: seq<T2> := [v248];
				var v250: map<T2, bool> := map[v249[safeIndex(v89, |v249|)] := !false];
				v250 := v250[v248 := false];
				v102[safeIndex(968, v102.Length)] := |"mnwoqpjhn"|;
				v102[safeIndex(968, v102.Length)], v1[safeIndex(875, v1.Length)], v89 := safeDivisionInt(v90, v89), v248.f9, v89 * v89;
			} else {
				v90 := v89;
				v91 := v91;
				var v251: array<multiset<bool>> := new multiset<bool>[26];
				v251[safeIndex(609, v251.Length)] := v196;
				var v252 := DC2(v111.f10);
				v251[safeIndex(609, v251.Length)] := fm19(v2, v252.(cf3 := v111.f10), 589 + |v105[safeIndex(762, v105.Length)]|, globalState);
				var v253: map<string, bool> := map[(fm4(v89, v89, globalState))[safeIndex(v90, |fm4(v89, v89, globalState)|) := v111.f10] := v1[safeIndex(875, v1.Length)]];
				var v254: map<int, map<string, bool>> := map[v89 * 0x1a4 := v253];
				v254 := v254[safeModuloInt(v90, |"tbnxyymh"|) := if (|v97| in v254) then v254[|v97|] else v253];
				var v255, v256, v257 := v98.m1(v105[safeIndex(762, v105.Length)] + v0, v1[safeIndex(875, v1.Length)] ==> false, globalState);
			}
			
			var v258: multiset<array<int>> := multiset{v102};
			var v259: T1 := new C1(v111.f10, v111.f11, !(multiset{v102} !! v258), true);
			var v260: map<int, set<seq<D4>>> := map[v89 := v109];
			v259 := new C1(v111.f10, (if (v90 in v260) then v260[v90] else if (v90 in v260) then v260[v90] else v111.f11) * v109, v2, v91);
			if (v2) {
				var v261: set<array<int>> := {v102};
				var v262: C2 := new C2(v91, v261 >= v261);
				var v263: multiset<string> := multiset{v0};
				var v264 := DC38(v262);
				var v265: map<bool, C2> := map[|v198| != 0x307 := v264.cf66];
				var v266: map<array<int>, string> := map[v102 := v0];
				v1[safeIndex(875, v1.Length)], v262, v0 := v263 !! v263, if ((if (v259.f8) then v91 else v1[safeIndex(875, v1.Length)]) in v265) then v265[if (v259.f8) then v91 else v1[safeIndex(875, v1.Length)]] else v262, v0 + (if (v102 in v266) then v266[v102] else v105[safeIndex(762, v105.Length)]);
				var v267: seq<multiset<int>> := [v95];
				var v268: set<bool> := {v95 >= v267[safeIndex(v89, |v267|)], v2};
				v268 := v268;
				v89, v105[safeIndex(762, v105.Length)] := v89, v0;
				v105[safeIndex(762, v105.Length)] := v0;
				var v269: map<int, C1> := map[v89 := v111];
				v269 := v269;
			} else {
				var v270: set<bool> := {v259.f8};
				var v271 := new C1(fm7(|map[v111 := |v198|]|, v2, v89, globalState), v111.f11 + v111.f11, v270 > v270, v91);
				v93.f12[safeIndex(370, v93.f12.Length)] := DC10(v102).cf17;
				v93.f12[safeIndex(370, v93.f12.Length)] := v102;
				var v272: map<int, set<int>> := map[v89 := v198];
				v246 := v246[929 := (if (v90 in v272) then v272[v90] else {v90, |v198|, v90}) > v198];
				v259.f8 := v2;
				v1[safeIndex(875, v1.Length)] := !v259.f8 ==> v259.f8;
			}
			
	}
	
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
	print v1[13], "\n";
	print v1[14], "\n";
	print v1[15], "\n";
	print v1[16], "\n";
	print v1[17], "\n";
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
	print globalState.f7, "\n";
	print v2, "\n";
	print v3 == [false, false, false, false], "\n";
	print v89, "\n";
	print v90, "\n";
	print v91, "\n";
	print v94 == map[false := 14], "\n";
	print v95 == multiset{1}, "\n";
	print v96 == map[false := 1], "\n";
	print v97 == map[1 := -27058], "\n";
	print v99.cf49.f8, "\n";
	print v99.cf49.f9, "\n";
	print v100.cf55.cf49.f8, "\n";
	print v100.cf55.cf49.f9, "\n";
	print v101[0].cf55.cf49.f8, "\n";
	print v101[0].cf55.cf49.f9, "\n";
	print v101[1].cf55.cf49.f8, "\n";
	print v101[1].cf55.cf49.f9, "\n";
	print v101[2].cf55.cf49.f8, "\n";
	print v101[2].cf55.cf49.f9, "\n";
	print v101[3].cf55.cf49.f8, "\n";
	print v101[3].cf55.cf49.f9, "\n";
	print v101[4].cf55.cf49.f8, "\n";
	print v101[4].cf55.cf49.f9, "\n";
	print v101[5].cf55.cf49.f8, "\n";
	print v101[5].cf55.cf49.f9, "\n";
	print v101[6].cf55.cf49.f8, "\n";
	print v101[6].cf55.cf49.f9, "\n";
	print v101[7].cf55.cf49.f8, "\n";
	print v101[7].cf55.cf49.f9, "\n";
	print v102[0], "\n";
	print v102[1], "\n";
	print v102[2], "\n";
	print v102[3], "\n";
	print v102[4], "\n";
	print v102[5], "\n";
	print v102[6], "\n";
	print v102[7], "\n";
	print v102[8], "\n";
	print v103 == [1, -881, 473, 1], "\n";
	print v104, "\n";
	print v105[0], "\n";
	print v105[1], "\n";
	print v105[2], "\n";
	print v105[3], "\n";
	print v105[4], "\n";
	print v105[5], "\n";
	print v105[6], "\n";
	print v105[7] == ['j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j', 'j'], "\n";
	print |v106|, "\n";
	print v107 == map[false := true], "\n";
	print v108.cf10, "\n";
	print v109 == {[DC7("xe"), DC7("xe"), DC7("xe"), DC7("xe")]}, "\n";
	print |v110|, "\n";
	print v111.f10, "\n";
	print v111.f11 == {[DC7("xe"), DC7("xe"), DC7("xe"), DC7("xe")]}, "\n";
	print i11, "\n";
	print v191.cf21, "\n";
	print v191.cf22, "\n";
	print v191.cf23, "\n";
	print v191.cf24, "\n";
	print v196 == multiset{true, true, false}, "\n";
	print |v197|, "\n";
	print v198 == {155, -161, 3, 436}, "\n";
	print v199.cf27 == {155, -161, 3, 436}, "\n";
	print v217 == map[-761 := multiset{1}], "\n";
	print v218.cf45, "\n";
	print v218.cf46 == map[-761 := multiset{1}], "\n";
}