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
datatype D0 = DC1(cf1: int, cf2: bool, cf3: int) | DC0(cf0: string) | DC2(cf4: D0)
datatype D1 = DC4(cf6: map<int, int>, cf7: bool, cf8: bool) | DC5(cf9: int, cf10: bool, cf11: int, cf12: int) | DC6(cf13: array<bool>) | DC3(cf5: seq<array<int>>)
datatype D2 = DC8(cf15: int, cf16: int, cf17: char, cf18: int, cf19: bool) | DC7(cf14: set<char>)
datatype D3 = DC10(cf21: bool) | DC11(cf22: int) | DC12(cf23: seq<bool>) | DC9(cf20: multiset<array<D1>>) | DC13(cf24: D3)
datatype D4 = DC15(cf26: int, cf27: char, cf28: int) | DC14(cf25: array<int>)
datatype D5 = DC16(cf29: T1)
datatype D6 = DC17(cf30: seq<string>)
datatype D7 = DC18(cf31: map<bool, int>)
datatype D8 = DC20(cf33: bool, cf34: int) | DC21(cf35: bool, cf36: bool, cf37: int) | DC22(cf38: bool) | DC19(cf32: seq<int>) | DC23(cf39: D8)
datatype D9 = DC25(cf41: bool, cf42: set<char>) | DC26(cf43: int) | DC27(cf44: bool, cf45: seq<int>, cf46: array<bool>, cf47: bool, cf48: bool) | DC24(cf40: array<char>) | DC28(cf49: D9)
datatype D10 = DC29(cf50: map<bool, char>)
datatype D11 = DC31(cf52: int) | DC32(cf53: string) | DC30(cf51: array<map<int, int>>) | DC33(cf54: D11)
datatype D12 = DC35 | DC36(cf56: bool, cf57: bool, cf58: int, cf59: int, cf60: map<bool, seq<int>>) | DC34(cf55: map<char, bool>) | DC37(cf61: D12)
datatype D13 = DC39(cf63: int) | DC38(cf62: set<bool>) | DC40(cf64: D13)
datatype D14 = DC42(cf66: map<array<bool>, string>) | DC41(cf65: seq<D9>)
datatype D15 = DC44(cf68: bool, cf69: seq<bool>, cf70: int, cf71: bool, cf72: bool) | DC43(cf67: map<seq<D11>, seq<int>>) | DC45(cf73: D15)
datatype D16 = DC47(cf75: int, cf76: int, cf77: C1, cf78: int) | DC46(cf74: set<string>)
datatype D17 = DC49(cf80: int, cf81: bool, cf82: bool, cf83: bool) | DC48(cf79: multiset<bool>)
datatype D18 = DC51 | DC50(cf84: map<multiset<int>, int>)
datatype D19 = DC53(cf86: int, cf87: int) | DC52(cf85: map<bool, bool>) | DC54(cf88: D19)
datatype D20 = DC56(cf90: bool, cf91: int, cf92: array<bool>, cf93: char, cf94: int) | DC57 | DC55(cf89: multiset<array<C2>>) | DC58(cf95: D20)
trait T0 {
	const f23 : int
	method m1(p0: array<bool>, p1: bool, globalState: GlobalState) returns (r0: int, r1: set<int>) 
	method m2(p0: string, globalState: GlobalState) returns (r0: int, r1: seq<int>) 
}

trait T1 extends T0 {
	const f28 : bool
	var f29 : seq<int>
}

class GlobalState {
	const f0 : bool
	const f1 : int
	const f2 : array<string>
	const f3 : int
	const f4 : bool
	const f5 : multiset<multiset<bool>>
	var f6 : string
	var f7 : bool
	var f8 : int
	var f9 : bool
	const f10 : int
	var f11 : char
	const f12 : int
	const f13 : array<set<int>>
	var f14 : int
	const f15 : int
	const f16 : bool
	const f17 : int
	const f18 : map<array<int>, int>
	const f19 : bool
	const f20 : array<int>
	var f21 : bool
	const f22 : int
	constructor (f0 : bool, f1 : int, f2 : array<string>, f3 : int, f4 : bool, f5 : multiset<multiset<bool>>, f6 : string, f7 : bool, f8 : int, f9 : bool, f10 : int, f11 : char, f12 : int, f13 : array<set<int>>, f14 : int, f15 : int, f16 : bool, f17 : int, f18 : map<array<int>, int>, f19 : bool, f20 : array<int>, f21 : bool, f22 : int) {
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
		this.f19 := f19;
		this.f20 := f20;
		this.f21 := f21;
		this.f22 := f22;
	}
	
}

class C0 {
	constructor () {
	}
	
	function fm20(p0: map<D2, string>, p1: int, p2: int, p3: int, globalState: GlobalState): bool {
		false
	}
}

class C1 extends T1, T0 {
	constructor (f28 : bool, f29 : seq<int>, f23 : int) {
		this.f28 := f28;
		this.f29 := f29;
		this.f23 := f23;
	}
	
	method m1(p0: array<bool>, p1: bool, globalState: GlobalState) returns (r0: int, r1: set<int>) {
		p0[safeIndex(216, p0.Length)] := p1;
		p0[safeIndex(216, p0.Length)] := f28;
		var v0 := "pilipchus";
		var v1: set<string> := {v0};
		f29 := fm19(|v0|, v0 !in v1, globalState);
		var v2 := 'n';
		var v3: multiset<char> := multiset{v2};
		globalState.f14, globalState.f7, r0, v3, p0[safeIndex(216, p0.Length)] := fm0(f23, p0[safeIndex(216, p0.Length)], globalState), f28, f23, v3, p1;
		var v5: map<array<bool>, int> := map[p0 := f23];
		for i0 := |[|(set v4 : int | v4 in (seq(abs(0x3c5), i1  => (f23))) :: (v4 - 0x30f))|, f23, if (p0 in v5) then v5[p0] else f23]| to f23 {
			if (p1) {
				globalState.f14 := if (v0 <= v0) then i0 * f23 else i0;
				var v6: seq<seq<bool>> := [([p1, f28, p0[safeIndex(216, p0.Length)], fm1(f23, f28, globalState), p1])[safeIndex(i0, |[p1, f28, p0[safeIndex(216, p0.Length)], fm1(f23, f28, globalState), p1]|) := f28]];
				var v7: map<int, int> := map[i0 := f23];
				var v8 := DC4(v7, p0[safeIndex(216, p0.Length)], f28);
				var v9: map<seq<seq<bool>>, bool> := map[v6 + v6[safeIndex(f23, |v6|) := [p1, p0[safeIndex(216, p0.Length)]]] := v8.cf8];
				v9 := v9[v6 := f28];
				var v10: map<bool, int> := map[p0[safeIndex(216, p0.Length)] := -0x177];
				var v11: seq<bool> := [f28];
				globalState.f14, globalState.f9 := f23 + (if (fm1(|v11|, f28, globalState) in v10) then v10[fm1(|v11|, f28, globalState)] else |(seq(abs(-960), i2  => ('v')))|), f23 > f23;
				var v12: array<int> := new int[13](i3 => i3 * 0x27b);
				var v13: seq<array<int>> := [v12];
				var v14 := DC3([v12, v12, v12]);
				var v15: array<D1> := new D1[10] [DC3(v13), v14, v14, v14, DC3(v13), v14, v14, v14, v14, v14];
				var v16: multiset<array<D1>> := multiset{v15, v15};
				var v17 := DC9(v16);
				var v18: seq<D3> := [v17, v17, v17, v17];
				v18 := (if (f28) then v18 else v18) + v18;
				globalState.f21 := f28;
			} else {
				var v19, v20, v21 := m0(p0, fm2(globalState) + v0, f29[safeIndex(f23, |f29|)], globalState);
				var v22: array<char> := new char[25];
				v22 := new char[24];
				var v23: map<int, int> := map[|v0| := v21];
				var v24: set<int> := {|v23|, f23};
				p0[safeIndex(216, p0.Length)], globalState.f7, globalState.f8 := p1, !(if (v24 > v24) then true else v0 != (seq(abs(0xc4), i4  => (v2)))), 595;
				var v25: map<int, char> := map[|f29| := v2];
				globalState.f11 := if (f23 in v25) then v25[f23] else fm6(i0, globalState);
				var v26: array<bool> := new bool[17];
				v26 := v26;
			}
			
			var v27: seq<bool> := [f28];
			var v28: seq<seq<bool>> := [v27, v27];
			p0[safeIndex(216, p0.Length)], v28, globalState.f21 := p1, v28 + v28[safeIndex(i0, |v28|) := v27], v27[safeIndex(safeDivisionInt(f23, i0), |v27|)];
			globalState.f8 := if (p0[safeIndex(216, p0.Length)]) then 950 + f23 else i0;
			globalState.f21 := f28;
		}
		var v29: seq<bool> := [false];
		var v30: map<bool, int> := map[f28 := f23];
		var v31: seq<bool> := [true, !(|v29| >= f23), p1, f28, false in v30];
		if (v31[safeIndex(f23, |v31|)]) {
			var v32: array<int> := new int[7];
			v32[safeIndex(532, v32.Length)] := |f29|;
			var v33: map<int, array<int>> := map[f23 := v32];
			var v34: map<map<int, array<int>>, bool> := map[v33 := true];
			globalState.f14, f29, v32[safeIndex(532, v32.Length)], globalState.f21, r0 := -f23, f29 + (seq(abs(288), i5  => (0x2cb))), -f23, if (v33 in v34) then v34[v33] else fm1(f23, p1, globalState), if (f28 <== !!p0[safeIndex(216, p0.Length)]) then |v31| else f23;
			globalState.f14 := v32[safeIndex(532, v32.Length)];
			var v35: seq<string> := [v0, fm2(globalState), v0, v0, v0];
			var v37: map<int, bool> := map[0x33f := f28];
			v1 := set v39 : string | v39 in ((set v36 : string | v36 in v35 :: (v36)) - (set v38 : string | v38 in map[v0 := if (|v30| in v37) then v37[|v30|] else p1] :: (v38))) :: (v39);
			var v41: set<int> := {|(set v40 : int | (0x310 <= v40) && (v40 < 0x1e0) :: (safeDivisionInt(v40, |map[f23 := multiset(f29)]|)))|, v32[safeIndex(532, v32.Length)]};
			r0 := safeDivisionInt(fm0(|fm5(v41, globalState)|, p0[safeIndex(216, p0.Length)], globalState), if (fm1(v32[safeIndex(532, v32.Length)], p1, globalState)) then v32[safeIndex(532, v32.Length)] else f23);
			if (p0[safeIndex(216, p0.Length)]) {
				p0[safeIndex(216, p0.Length)] := p0[safeIndex(216, p0.Length)];
				var v42 := new C0();
				var v43: array<string> := new string[16];
				var v44: seq<seq<bool>> := [v29, v31, [p1]];
				globalState.f9, v43 := true !in (v44[safeIndex(v32[safeIndex(532, v32.Length)], |v44|)] + v29), v43;
				globalState.f9 := p1;
				var v45: set<char> := {v2, v2};
				var v46 := DC7(v45);
				var v47: map<D2, string> := map[v46 := v0];
				var v48: multiset<bool> := multiset{v42.fm20(v47, v32[safeIndex(532, v32.Length)], |v0|, v32[safeIndex(532, v32.Length)], globalState)};
				var v49: map<multiset<bool>, string> := map[v48 := "dux"];
				var v50 := DC0(v0);
				var v51: seq<multiset<bool>> := [fm4(v50, f23, [f28], v2, globalState), v48];
				v49 := v49[v51[safeIndex(f23, |v51|)] := v0];
			} else {
				var v52: map<int, int> := map[f23 := 0x113];
				v52 := v52[v32[safeIndex(532, v32.Length)] := v32[safeIndex(532, v32.Length)]];
				v41 := fm21(v35, globalState);
				var v53, v54, v55 := m0(p0, v0[safeIndex(f23, |v0|) := v2] + v0, f23, globalState);
				v32[safeIndex(532, v32.Length)] := f23;
				globalState.f9 := !((v32[safeIndex(532, v32.Length)] != v55) <==> (v53.cf3 == v32[safeIndex(532, v32.Length)]));
			}
			
		} else {
			var v56: array<array<seq<int>>> := new array<seq<int>>[2];
			var v57: array<seq<int>> := new seq<int>[19](i6 => f29);
			var v58: seq<array<seq<int>>> := [v57, v57];
			v56[safeIndex(242, v56.Length)] := v58[safeIndex(f23, |v58|)];
			v56[safeIndex(242, v56.Length)] := v57;
			var v59: array<int> := new int[18];
			v59[safeIndex(461, v59.Length)] := f23;
			v59[safeIndex(461, v59.Length)] := f23;
			var v60: set<bool> := {!f28, p0[safeIndex(216, p0.Length)], false, p1};
			globalState.f9, p0[safeIndex(216, p0.Length)], r0, globalState.f9, v59[safeIndex(461, v59.Length)] := fm1(fm0(-|v60|, p1, globalState), f28, globalState), !f28, f23, p0[safeIndex(216, p0.Length)], safeModuloInt(safeModuloInt(v59[safeIndex(461, v59.Length)], |v0|), safeDivisionInt(f23, fm0(v59[safeIndex(461, v59.Length)], p1, globalState)));
			var v61 := DC12(v31);
			var v62 := DC10(true);
			v61 := fm22(v62, globalState);
			v59[safeIndex(461, v59.Length)] := f29[safeIndex(v59[safeIndex(461, v59.Length)], |f29|)];
		}
		
		for i7 := f23 to f23 {
			var v63: array<bool> := new bool[1] [p0[safeIndex(216, p0.Length)]];
			v63 := v63;
			globalState.f11 := v2;
			v30 := v30 + v30;
			p0[safeIndex(216, p0.Length)], globalState.f7, globalState.f7 := fm1(f23, p0[safeIndex(216, p0.Length)], globalState), f28, v29[safeIndex(i7, |v29|)];
		}
		r0 := f23;
		var v64: set<int> := {-(f23 * f23), -|v0|};
		r1 := v64;
	}
	method m2(p0: string, globalState: GlobalState) returns (r0: int, r1: seq<int>) {
		globalState.f6, globalState.f6, r0, r0 := "hoiflwc" + p0, p0 + p0, 0x281, f23 + f23;
		var v0: multiset<bool> := multiset{f28};
		var v1: map<bool, int> := map[multiset{f28} >= v0 := -f23];
		v1 := v1[f28 || f28 := 0x39a];
		globalState.f14 := -f23;
		var v2: multiset<int> := multiset{f23};
		var v3: seq<map<int, multiset<int>>> := [map[308 := v2]];
		var v4: map<int, multiset<int>> := map[f23 := v2];
		var v5: seq<map<int, multiset<int>>> := [v3[safeIndex(f23, |v3|)], v4[324 := v2], v4];
		var v6: map<int, int> := map[f23 := |v5[safeIndex(-f23, |v5|)]|];
		var v7 := 'n';
		var v8 := DC8(f23, |v6|, v7, f23, f28);
		v1 := v1[v8.cf19 := -f23];
		var v9: array<int> := new int[23];
		for i0 := -|[v9]| to f23 {
			globalState.f14 := i0 * f23;
			if (!fm1(i0, f28, globalState)) {
				globalState.f21 := !f28;
				var v10: map<bool, bool> := map[f28 := f28];
				var v11: seq<bool> := [f28, f28];
				v10 := v10[f28 := v11[safeIndex(i0, |v11|)]];
				var v12: array<array<char>> := new array<char>[19];
				var v13: array<char> := new char[3];
				v12[safeIndex(536, v12.Length)] := v13;
				v12[safeIndex(536, v12.Length)] := new char[23];
				r0 := f23;
				var v14: array<string> := new string[20](i1 => p0[safeIndex(i0, |p0|) := v7]);
				v14[safeIndex(690, v14.Length)] := p0[safeIndex(i0, |p0|) := v7] + p0;
				var v15: seq<string> := [(p0[safeIndex(i0, |p0|) := v7])[safeIndex(i0, |p0[safeIndex(i0, |p0|) := v7]|) := v7]];
				v14[safeIndex(363, v14.Length)] := v15[safeIndex(i0, |v15|)] + p0;
				var v16: seq<multiset<bool>> := [v0];
				var v17: array<multiset<bool>> := new multiset<bool>[21] [v0, v0, multiset{f28, f28, f28, f28}, v0, v0, v0, (v0[f28 := abs(f23)])[f28 := abs(f23)] + multiset{f28, true}, v0, v0, v0, v0, v0 * v0, v0[f28 := abs(-f23)], v0 + multiset{f28}, v0, v0, v0, v0 - v0, v16[safeIndex(i0, |v16|)], v0, multiset([f28]) * v0];
				v17[safeIndex(314, v17.Length)] := v0;
				var v18: set<bool> := {!f28 <== f28};
				v14[safeIndex(690, v14.Length)], v14[safeIndex(363, v14.Length)], v17[safeIndex(314, v17.Length)], v18, globalState.f6 := "bshdxkt" + p0, (seq(abs(0x80), i2  => (v7))) + "sgs", v0, v18, p0;
			} else {
				globalState.f21 := f28;
				v9[safeIndex(258, v9.Length)] := i0;
				globalState.f6, v9[safeIndex(258, v9.Length)] := "celkv", |[f28]|;
				var v19 := DC1(f23, f28, if (f23 in v2) then v2[f23] else |v2|);
				globalState.f9 := v19.cf2;
				globalState.f21 := f28;
				var v20: map<int, bool> := map[f23 := if (fm1(-v9[safeIndex(258, v9.Length)], f28, globalState)) then f28 else f28];
				v20 := v20[f23 := f28];
			}
			
			var v21 := DC15(-165, v7, safeModuloInt(f29[safeIndex(|v2|, |f29|)], f23));
			v21 := v21;
			var v22: array<bool> := new bool[22];
			v22[safeIndex(424, v22.Length)] := f28;
			v22[safeIndex(424, v22.Length)] := v8.cf19;
		}
		var i3 := 0;
		while (!(if (f28) then 0x274 != DC5(f23, f28, f23, f23).cf9 else f28))
			decreases 100 - i3
		{
			if (i3 >= 100) {
				break;
			}
			
			i3 := i3 + 1;
			var v23: seq<bool> := [f28];
			var v24 := DC10(v23[safeIndex(f23, |v23|)]);
			var v25: set<bool> := {f28};
			var v26: map<int, bool> := map[0x149 := f28];
			var v27: map<bool, bool> := map[f28 := if (f23 in v26) then v26[f23] else f28];
			var v28: array<bool> := new bool[29] [f28, fm1(f23, f28, globalState), f28, if (f28) then f28 else f28, f28, f28, v24.cf21, f28, f28, f28, f28, f28, fm1(f23, f28, globalState) <== v23[safeIndex(f23, |v23|)], if (f28) then f28 else f28, v25 == {f28}, f28, f28, f28, f28, if (f28 in v27) then v27[f28] else f28, f23 !in f29[safeIndex(f23, |f29|) := 0x80], !(-f23 < f23), f28, f28, (if (f28 in v1) then v1[f28] else f23) > -581, f28, v23[safeIndex(f23, |v23|)], !f28, f23 < f23];
			v28[safeIndex(547, v28.Length)] := if (fm1(f23, f28, globalState)) then f28 else f28;
			var v29: multiset<array<int>> := multiset{v9, v9, v9};
			v28[safeIndex(547, v28.Length)], r0, globalState.f21, globalState.f7, globalState.f21 := multiset{v9, v9} <= v29, f23, fm1(f23, f28, globalState), f28 ==> f28, fm1(f23, f28, globalState);
			var v30: map<seq<bool>, bool> := map[if (!true) then v23 else [f28] := f28];
			var v31: map<char, string> := map[v7 := p0];
			v30 := v30[v23 := v28[safeIndex(547, v28.Length)] ==> v23[safeIndex(|v31|, |v23|)]];
			var v32 := new C0();
			var v33 := new C0();
		}
		r0 := (f23 - f23) + 0x2c8;
		r1 := f29 + (if (fm1(f23, !f28, globalState)) then f29 else f29);
	}
}

class C2 extends T1 {
	const f31 : map<int, int>
	const f32 : bool
	constructor (f31 : map<int, int>, f32 : bool, f28 : bool, f29 : seq<int>, f23 : int) {
		this.f31 := f31;
		this.f32 := f32;
		this.f28 := f28;
		this.f29 := f29;
		this.f23 := f23;
	}
	
	function fm18(p0: map<bool, int>, p1: string, p2: bool, globalState: GlobalState): int {
		-|((if (f28) then multiset(seq(abs(779), i0  => ('x'))) else multiset{'k', 'g'}) * (multiset{'s'} * multiset{'n'}))|
	}
	method m1(p0: array<bool>, p1: bool, globalState: GlobalState) returns (r0: int, r1: set<int>) {
		var v0 := "rovoqj";
		var v1: map<string, bool> := map[v0 := true];
		var v2: map<string, map<string, bool>> := map[v0 + v0 := v1 + v1];
		v2 := v2[seq(abs(-0x188), i0  => ('u')) := v1 + v1];
		var v3: array<seq<bool>> := new seq<bool>[26](i1 => [f32, p1] + [p1]);
		v3 := v3;
		var i2 := 0;
		while (f23 <= f23)
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			var v4: multiset<string> := multiset{v0, v0};
			var v5: T0 := new C1(false, f29, |v0|);
			var v6: multiset<T0> := multiset{v5, v5, v5};
			var v7 := 'f';
			var v8: array<int> := new int[11] [f23, -(if (v0 in v4) then v4[v0] else f23), |{-334, |v6|}|, 118, v5.f23, f29[safeIndex(248, |f29|)], f23, f23, f23, f23, |[v7]|];
			var v9: seq<array<int>> := [v8, v8, v8, v8, v8];
			var v10: map<seq<array<int>>, int> := map[v9 := v5.f23];
			v10 := v10[([v8] + v9)[safeIndex(v5.f23, |([v8] + v9)|) := v8] := f23];
			var v11: array<seq<C0>> := new seq<C0>[16];
			var v12: C0 := new C0();
			var v13: seq<C0> := [v12];
			v11[safeIndex(624, v11.Length)] := v13;
			v11[safeIndex(624, v11.Length)] := if (f28) then v13 else v13;
			var v14 := DC4(f31, p1, true);
			if (v14.cf8) {
				var v15 := DC3(v9);
				v15 := v15;
				globalState.f9 := p1;
				var v16: multiset<bool> := multiset{f28};
				globalState.f9 := (v16 >= v16) || p1;
				globalState.f7, globalState.f8, globalState.f8 := !(safeModuloInt(v5.f23, v5.f23) <= v5.f23), -f23, v5.f23;
				globalState.f7 := f32 <==> (v0 < v0[safeIndex(v5.f23, |v0|) := v7]);
			} else {
				var v18: seq<string> := [v0, "tocc"];
				var v19: map<int, map<string, int>> := map[safeDivisionInt(v5.f23, v5.f23) := map v17 : string | v17 in v18[safeIndex(0x1da, |v18|) := v0] :: (v17) := (|multiset{f28}|)];
				var v20: map<string, int> := map[v0 := -v5.f23];
				v19 := v19[f23 := v20];
				var v21: map<int, bool> := map[fm0(v5.f23, p1, globalState) := p1];
				v21 := v21[f23 * v5.f23 := fm1(v5.f23, f28, globalState)];
				globalState.f14 := fm0(v5.f23, p1, globalState) * (-0x2a5 * v5.f23);
				var v22: map<bool, int> := map[f32 := 0x170];
				v22 := v22[p1 := |fm2(globalState)| - v5.f23];
				v0, globalState.f14 := (v0 + (seq(abs(-846), i3  => (v7)))) + v0, f23;
			}
			
			var v23: array<D0> := new D0[2];
			var v24 := DC1(f23, f32, f23);
			v23[safeIndex(326, v23.Length)] := v24;
			v23[safeIndex(326, v23.Length)] := v24;
		}
		var v25: set<bool> := {p1};
		v25 := v25;
		var i4 := 0;
		while (f28)
			decreases 100 - i4
		{
			if (i4 >= 100) {
				break;
			}
			
			i4 := i4 + 1;
			var v26: seq<bool> := [f32, f32];
			var v27: multiset<int> := multiset{fm0(f23, p1, globalState), f23};
			var v28: seq<array<bool>> := [p0];
			var v29: multiset<string> := multiset{v0};
			var v30: array<int> := new int[26] [fm0(f23, f32, globalState), f23, |(if (v26[safeIndex(f23, |v26|)]) then v0 else seq(abs(0xb), i5  => ('k')))|, f23, |"hpcxes"|, f23, |((seq(abs(0x1e1), i6  => ('g'))) + v0)|, f23, f23, f23, safeDivisionInt(f23, f23), f23, if (p1) then f23 else 222, -f23, 583, f23, |v0| * |(multiset{p1})[p1 := abs(f23)]|, f23, f23, f23 * f23, |v25|, f23, if (|v28| in v27) then v27[|v28|] else f23, |v29|, f23, f23];
			v30 := v30;
			var v31: multiset<bool> := multiset{f32, false};
			var v32: seq<string> := ["thg", ("suocs")[safeIndex(|v31|, |"suocs"|) := 'p'], seq(abs(247), i7  => (DC15(f23, 'f', |v0|).cf27)), v0, v0];
			var v33: seq<seq<bool>> := [v26];
			var v34: map<string, seq<seq<bool>>> := map[v32[safeIndex(fm0(348, !f28, globalState), |v32|)] := v33];
			v34 := v34[v0 := v33];
			var v35: map<int, array<int>> := map[f23 := v30];
			var v36: array<array<int>> := new array<int>[13] [if (f32) then v30 else v30, v30, v30, v30, v30, if (f23 in v35) then v35[f23] else v30, v30, v30, v30, v30, v30, v30, v30];
			v36[safeIndex(502, v36.Length)] := v30;
			v36[safeIndex(502, v36.Length)] := v30;
			var v37 := new C0();
		}
		var v38: seq<bool> := [f32];
		var v39: multiset<int> := multiset{f23};
		v38 := v38[safeIndex(fm0(|map[|v39| := multiset(v38)]|, !p1, globalState), |v38|) := fm1(f23, f32, globalState)] + v38;
		r0 := f23;
		var v40: set<int> := {f23, |f29|};
		r1 := v40 + v40;
	}
	method m2(p0: string, globalState: GlobalState) returns (r0: int, r1: seq<int>) {
		var v0: array<bool> := new bool[29];
		var v1: seq<string> := [p0, p0, p0 + p0];
		var v2 := DC0(p0);
		var v3 := DC5(f23, f32, f23, f23);
		globalState.f7, v0, globalState.f9, v1 := "wcapto" == p0, v0, match v2 {
			case DC1(cf1, cf2, cf3) => if (cf2) then cf2 else f32
			case DC0(cf0) => f28
			case DC2(cf4) => p0[safeIndex(f23, |p0|) := DC8(f23, f23, 'o', |p0|, f28).cf17] != (p0[safeIndex(f23, |p0|) := 'i'])[safeIndex(f23, |p0[safeIndex(f23, |p0|) := 'i']|) := 'b']
		}, match v3 {
			case DC4(cf6, cf7, cf8) => [p0]
			case DC5(cf9, cf10, cf11, cf12) => DC17(v1).cf30 + [p0]
			case DC6(cf13) => [p0, seq(abs(-0x2c), i0  => ('f')), p0, p0]
			case DC3(cf5) => v1
		};
		var v4 := DC6(v0);
		var v5: multiset<array<bool>> := multiset{v4.cf13};
		v5 := v5 + v5;
		globalState.f21 := f32 ==> f28;
		for i1 := f23 - f23 to f23 {
			var v6: multiset<int> := multiset{f23};
			if (i1 in (v6 - v6)) {
				globalState.f7 := true;
				var v7: map<int, bool> := map[-f23 := f32];
				v7 := if (!false) then v7 else v7;
				var v8: map<int, int> := map[i1 := -(f23 * f23)];
				v8 := v8[i1 := |f31|];
				globalState.f8 := |(p0 + (p0 + p0))[safeIndex(i1 - |v5[v0 := abs(f23)]|, |(p0 + (p0 + p0))|) := 'b']|;
				var v9: multiset<bool> := multiset{f32, f28};
				var v10: seq<bool> := [if (f23 in v7) then v7[f23] else f28, f32];
				var v11: map<bool, bool> := map[!f28 := f28];
				var v12: seq<map<bool, bool>> := [map[f28 := f28], v11];
				var v13: array<int> := new int[24] [fm0(i1, f32, globalState), f23, i1, i1, i1, f23, |v10|, 362, i1, |v12[safeIndex(i1, |v12|)]|, |p0|, f23, |v9[true := abs(i1)]|, fm0(f23, f28, globalState), f23, 0x180, |p0|, i1, 501, |v9|, f23, |v10|, i1, i1];
				var v14 := new C1(f32, f29, if (!f28 in v9) then v9[!f28] else -|map[f32 := v13]|);
			} else {
				var v15: array<int> := new int[27](i2 => safeModuloInt(i2, -f23));
				v15[safeIndex(445, v15.Length)] := i1;
				v15[safeIndex(445, v15.Length)] := f23;
				globalState.f21 := f32 || !f32;
				globalState.f7 := (p0 + p0) != fm2(globalState);
				var v16: map<bool, array<int>> := map[f32 <==> f28 := v15];
				v16 := v16 + v16;
				var v17: map<int, bool> := map[i1 := multiset{true, f32, f32} !! multiset{true}];
				globalState.f9 := !(if (i1 in v17) then v17[i1] else f32);
			}
			
			var v18: map<bool, string> := map[f32 := seq(abs(-469), i3  => ('d'))];
			var v19 := new C1(f32, f29, |(if (f32 in v18) then v18[f32] else p0)|);
			var v20: seq<bool> := [f28, true, f32, f32, fm1(|(seq(abs(131), i4  => ('t')))|, f32, globalState)];
			var v21: map<int, string> := map[f23 := p0];
			var v22: map<seq<bool>, string> := map[v20 := p0 + (if (i1 in v21) then v21[i1] else "hkctqqp")];
			v22 := v22[v20 := p0];
			var v23: array<D1> := new D1[15];
			v23[safeIndex(72, v23.Length)] := v3;
			v23[safeIndex(72, v23.Length)] := if (f32 <== f32) then fm23(globalState) else v3;
		}
		globalState.f14 := f23;
		var v24: array<D1> := new D1[11];
		forall i5 | 0 <= i5 < v24.Length {
			v24[i5] := DC4(map[f23 := f23], f28, !f32);
		}
		r0 := f23;
		r1 := ((seq(abs(-0x39a), i6  => (|p0|))) + (seq(abs(454), i7  => (f23))))[safeIndex(f23, |((seq(abs(-0x39a), i6  => (|p0|))) + (seq(abs(454), i7  => (f23))))|) := f23];
	}
	method m6(p0: bool, globalState: GlobalState) returns (r0: bool, r1: map<int, array<bool>>, r2: D1, r3: array<bool>) {
		var v0: multiset<bool> := multiset{f28, f28, p0, true};
		var v1: map<bool, int> := map[false := |v0|];
		var v2 := "fnnsolgl";
		var v3 := DC8(fm18(v1, v2, f28, globalState), f23, 'a', 0x32c, false);
		var v4 := DC8(f23, f23, 't', v3.cf16, v3.cf19);
		var v5: set<D2> := {v4, v3};
		v5 := v5;
		var v6: array<int> := new int[8];
		v6[safeIndex(669, v6.Length)] := 0x210;
		v6[safeIndex(669, v6.Length)] := f23;
		for i0 := |f29| - f23 to v6[safeIndex(669, v6.Length)] - v6[safeIndex(669, v6.Length)] {
			globalState.f8 := v6[safeIndex(669, v6.Length)];
			var v7 := 'w';
			var v8: map<bool, char> := map[p0 := v7];
			v6[safeIndex(669, v6.Length)] := if (v8[f32 := v7] != v8) then if (p0 in v1) then v1[p0] else f23 else f23 * -0x33a;
			var v9 := new C1(f28, [i0] + f29, i0);
			f29 := f29;
		}
		var v10: seq<bool> := [p0, f28];
		var v11 := DC12(v10);
		v11 := if (p0 <==> p0) then v11 else v11;
		var v12: array<bool> := new bool[25];
		v12[safeIndex(779, v12.Length)] := p0;
		v12[safeIndex(779, v12.Length)] := f32;
		var i1 := 0;
		while (v12[safeIndex(779, v12.Length)])
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			r0 := !v12[safeIndex(779, v12.Length)];
			var v13: seq<string> := [v2, v2];
			globalState.f21 := v2 !in v13;
			var v14: seq<map<bool, int>> := [map[false := |DC18(v1).cf31|]];
			globalState.f14 := fm18(v14[safeIndex(f23, |v14|)], (seq(abs(339), i2  => ('h'))) + v2, v12[safeIndex(779, v12.Length)], globalState);
			v1 := v1[f32 := f23];
		}
		r0 := f28;
		var v15: map<int, array<bool>> := map[f23 * f23 := v12];
		r1 := v15;
		var v16 := DC4(f31, v12[safeIndex(779, v12.Length)], v12[safeIndex(779, v12.Length)]);
		r2 := v16;
		r3 := v12;
	}
}

class C3 extends T0, T1 {
	const f30 : seq<multiset<bool>>
	constructor (f30 : seq<multiset<bool>>, f23 : int, f28 : bool, f29 : seq<int>) {
		this.f30 := f30;
		this.f23 := f23;
		this.f28 := f28;
		this.f29 := f29;
	}
	
	function fm12(p0: int, p1: bool, p2: bool, globalState: GlobalState): int {
		-(f23 + safeModuloInt(|(seq(abs(-0x249), i0  => ('r')))|, 0x3ce))
	}
	method m1(p0: array<bool>, p1: bool, globalState: GlobalState) returns (r0: int, r1: set<int>) {
		var v0: array<bool> := new bool[22];
		forall i0 | 0 <= i0 < v0.Length {
			v0[i0] := DC8(|multiset([map[f23 := f23]])|, f23, 'p', |"nasts"|, true).cf19;
		}
		if (safeDivisionInt(f23, f23) < (f23 * f23)) {
			var v1: set<bool> := {p1, p1};
			globalState.f21, globalState.f21 := (v1 - {f28}) > v1, f28;
			var v2: map<bool, bool> := map[f28 := false];
			p0[safeIndex(868, p0.Length)] := if (p1 in v2) then v2[p1] else p1;
			p0[safeIndex(868, p0.Length)] := f28;
			var v3: array<map<int, bool>> := new map<int, bool>[20];
			var v4: map<int, bool> := map[f23 := true];
			v3[safeIndex(243, v3.Length)] := v4;
			v3[safeIndex(243, v3.Length)] := v4;
			p0[safeIndex(868, p0.Length)] := p0[safeIndex(868, p0.Length)] !in v2;
			var v5 := DC5(f23, p0[safeIndex(868, p0.Length)], f23, f23);
			var v6: seq<D1> := [v5];
			var v7: map<seq<D1>, bool> := map[v6 := p1];
			var v8: map<int, map<int, int>> := map[f23 := map[-0xe1 := f23]];
			var v9: multiset<bool> := multiset{f28};
			var v10 := DC11(fm0(|v9|, f28, globalState));
			var v11: map<int, int> := map[|v9| := v10.cf22];
			var v12: set<int> := {f23};
			var v13 := DC4(v11, f28 || p1, |v12| != f23);
			v7, globalState.f14, v8, v13, globalState.f14 := v7, f23, fm13('q', fm1(f23, f28, globalState), if (fm1(f23, p1, globalState)) then f23 else f23, if (p1) then p0[safeIndex(868, p0.Length)] else fm1(f23, p0[safeIndex(868, p0.Length)], globalState), globalState), v13, 0x1d9;
		} else {
			globalState.f14 := f23;
			var v14: set<bool> := {p1};
			var v15 := 'h';
			var v16: map<bool, int> := map[p1 := f23];
			var v17: multiset<int> := multiset{f23, fm12(f23, f28, p1, globalState), f23};
			var v18 := "htcv";
			var v19: array<int> := new int[28] [f23, safeDivisionInt(f23, f23), f23, f23, f23, |"n"|, |v14| * f23, f23, |fm14(f23, v15, f28, globalState)|, f23, f23, f23, fm12(f23, fm1(f23, p1, globalState), !p1, globalState) + 0x1bf, f23, 0x275, fm12(fm0(f23, p1, globalState), f28, f28, globalState), f23, |v16| * (if (0x36c in v17) then v17[0x36c] else f23), f23 + fm12(|v18|, f28, p1, globalState), 547 + f23, |v18|, if (p1) then |f29| else f23, f23, f23, safeModuloInt(f23, f23), |v17[f23 := abs(fm12(f23, f28, f28, globalState))]|, 0x3c6, f23];
			v19[safeIndex(726, v19.Length)] := f23;
			var v20: map<bool, bool> := map[f28 := f23 <= 141];
			v19[safeIndex(726, v19.Length)] := -|v20|;
			globalState.f8 := fm0(f23, p1, globalState);
			if (if (f28) then f28 else p1) {
				var v21: seq<bool> := [false, p1, f28];
				var v22: multiset<seq<bool>> := multiset{v21 + v21};
				var v23: multiset<bool> := multiset{p1};
				var v24: map<int, multiset<bool>> := map[v19[safeIndex(726, v19.Length)] := v23];
				v22 := fm15((DC10(p1).(cf21 := p1)).cf21, v24, globalState);
				var v25: seq<array<bool>> := [v0, v0];
				var v26: array<array<bool>> := new array<bool>[14] [v0, p0, p0, v0, if (f28) then p0 else v0, v0, v25[safeIndex(|multiset(f29)|, |v25|)], v0, if (p1) then p0 else v0, p0, p0, v0, p0, p0];
				v26[safeIndex(281, v26.Length)] := p0;
				v26[safeIndex(281, v26.Length)] := p0;
				v14 := v14;
				var v27: array<array<int>> := new array<int>[15];
				v27[safeIndex(602, v27.Length)] := v19;
				var v28: map<int, bool> := map[v19[safeIndex(726, v19.Length)] := p1];
				var v29: map<string, map<bool, int>> := map[v18 := if (p1) then map[v21[safeIndex(v19[safeIndex(726, v19.Length)], |v21|)] := |v23|] else map[if (|v18| in v28) then v28[|v18|] else f28 := v19[safeIndex(726, v19.Length)]]];
				v16, v27[safeIndex(602, v27.Length)], v19[safeIndex(726, v19.Length)], globalState.f9, globalState.f14 := if (v18 in v29) then v29[v18] else v16, v19, fm12(v19[safeIndex(726, v19.Length)], false, p1, globalState), f28, v19[safeIndex(726, v19.Length)] - f23;
				globalState.f8 := v19[safeIndex(726, v19.Length)];
			} else {
				var v30: map<array<bool>, int> := map[v0 := f23 * 0x1b2];
				v30 := v30[v0 := f23];
				v19[safeIndex(726, v19.Length)] := v19[safeIndex(726, v19.Length)];
				var v31: seq<seq<int>> := [f29 + f29, f29];
				var v32: map<int, seq<int>> := map[0x187 := f29];
				v31 := (v31[safeIndex(if (f23 in v17) then v17[f23] else f23, |v31|) := f29])[safeIndex(|[p1]|, |v31[safeIndex(if (f23 in v17) then v17[f23] else f23, |v31|) := f29]|) := if (v19[safeIndex(726, v19.Length)] in v32) then v32[v19[safeIndex(726, v19.Length)]] else [v19[safeIndex(726, v19.Length)]]] + (v31 + v31);
				m5(globalState);
				var v33 := DC14(v19);
				v19 := v33.cf25;
			}
			
			var v34: map<int, int> := map[DC1(v19[safeIndex(726, v19.Length)], f28, v19[safeIndex(726, v19.Length)]).cf1 := f23];
			var v35: seq<map<int, int>> := [fm16(f28, f23, f23, globalState)];
			v34 := v34[f23 := |(v35[safeIndex(v19[safeIndex(726, v19.Length)], |v35|)] + map[|v18| := f23])|];
		}
		
		globalState.f14 := f23;
		var v36 := "tjvi";
		var v37 := 'y';
		if (!("ii" == v36[safeIndex(f23, |v36|) := v37])) {
			v0[safeIndex(956, v0.Length)] := f28;
			v0[safeIndex(956, v0.Length)] := !DC10(f28).cf21;
			var v38: array<multiset<bool>> := new multiset<bool>[14](i1 => multiset{v0[safeIndex(956, v0.Length)]});
			v38 := v38;
			var v40 := DC4(map v39 : int | (0x39 <= v39) && (v39 < 0x3d8) :: (v39 + f23) := (f23), v0[safeIndex(956, v0.Length)], p1);
			v0[safeIndex(956, v0.Length)] := fm1(f23, false, globalState) <==> v40.cf7;
			globalState.f6 := v36 + "l";
			globalState.f8 := f29[safeIndex(f23, |f29|)];
		} else {
			v0[safeIndex(296, v0.Length)] := p1;
			v0[safeIndex(296, v0.Length)] := p1;
			var v41: array<int> := new int[11](i2 => i2 + f23);
			var v42: seq<bool> := [v0[safeIndex(296, v0.Length)], p1];
			v41[safeIndex(509, v41.Length)] := |v42|;
			v41[safeIndex(509, v41.Length)] := fm12(f23, v0[safeIndex(296, v0.Length)], v0[safeIndex(296, v0.Length)], globalState);
			var v43: array<string> := new string[18](i3 => v36 + v36);
			v43[safeIndex(35, v43.Length)] := v36;
			var v44: multiset<char> := multiset{v37};
			var v45: map<bool, int> := map[true := v41[safeIndex(509, v41.Length)]];
			var v46: seq<string> := ["qkbs", "uxecrdjl"];
			var v47: set<bool> := {true, f28};
			var v48: multiset<set<bool>> := multiset{v47};
			var v49 := DC15(|v48|, v37, f23);
			v41[safeIndex(509, v41.Length)], v41[safeIndex(509, v41.Length)], v43[safeIndex(35, v43.Length)], v41[safeIndex(509, v41.Length)] := v41[safeIndex(509, v41.Length)], fm0(f23, v44[v37 := abs(|v45|)] > fm17(f28, |v36|, globalState), globalState), v36 + v36, if ((seq(abs(39), i4  => (v36))) < v46) then if (v0[safeIndex(296, v0.Length)]) then v49.cf28 else fm12(f23, f28, fm1(f23, p1, globalState), globalState) else v41[safeIndex(509, v41.Length)] - f23;
			m5(globalState);
			v0[safeIndex(296, v0.Length)] := v41[safeIndex(509, v41.Length)] == |v45|;
		}
		
		for i5 := -f23 to 581 {
			var v50: map<bool, bool> := map[p1 := !p1];
			var v51: map<int, array<bool>> := map[|v50| := p0];
			var v52: array<int> := new int[17](i6 => i6 + |{f29}|);
			var v53: map<array<bool>, array<int>> := map[if (f23 in v51) then v51[f23] else p0 := v52];
			var v54 := DC14(if (p0 in v53) then v53[p0] else v52);
			v54 := DC14(v52);
			var v55: map<int, int> := map[f23 := f23];
			var v56 := DC4(v55, p1, p1);
			match (v56) {
				case DC4(cf6, cf7, cf8) =>
					globalState.f14 := |v36| * (f23 * f23);
					var v57: map<bool, int> := map[f28 := f23];
					var v58: multiset<int> := multiset{|v57|};
					cf7, v54, v37, globalState.f14, v54 := v58 >= v58, v54, v37, i5 + f23, v54;
					var v59: map<bool, map<bool, bool>> := map[fm1(0x277, cf8, globalState) ==> !cf8 := v50];
					var v60: seq<map<bool, bool>> := [v50];
					v59 := v59[f28 := v50[p1 := cf8] + v60[safeIndex(f23, |v60|)]];
					var v61: map<map<array<int>, set<char>>, int> := map[map[v52 := {v37, v37}] := -f23];
					var v62: set<char> := {v37, 'e', 'a', v37};
					var v63: map<array<int>, set<char>> := map[v52 := v62];
					v61 := v61[v63 + v63 := 0x238];
				case DC5(cf9, cf10, cf11, cf12) =>
					var v64: array<T1> := new T1[21];
					var v65: T1 := new C2(v55, false, cf10, [0x3ca], f29[safeIndex(-f23, |f29|)]);
					var v66 := DC16(v65);
					v64[safeIndex(821, v64.Length)] := v66.cf29;
					v64[safeIndex(821, v64.Length)] := v65;
					var v67: seq<bool> := [f28, p1, cf10, f28];
					cf11 := cf11 - safeDivisionInt(|v67|, f23);
					globalState.f14 := cf12;
					v0[safeIndex(836, v0.Length)] := !true;
					v0[safeIndex(836, v0.Length)] := cf10;
				case DC6(cf13) =>
					var v68: C1 := new C1(f28, f29, fm0(f23, f28, globalState));
					v68 := v68;
					v50 := v50[p1 := !p1];
					var v69: map<bool, char> := map[f28 := v37];
					v69 := v69[f28 := v37];
					var v70: map<bool, map<bool, bool>> := map[false := v50];
					globalState.f9 := f28 in v70;
				case DC3(cf5) =>
					globalState.f7, r0 := p1, 0x198;
					v52[safeIndex(936, v52.Length)] := f23;
					v52[safeIndex(936, v52.Length)] := fm12(f23, p1, f28, globalState);
					var v71, v72, v73 := m0(if (f28) then p0 else v0, "nplonay", i5, globalState);
					globalState.f14 := i5 + v73;
			}
			
			globalState.f14 := i5 - |v36|;
			v0[safeIndex(107, v0.Length)] := p1;
			v0[safeIndex(107, v0.Length)] := fm1(f23, !p1, globalState);
		}
		var v74: multiset<bool> := multiset{!true};
		var v75: set<int> := {f23, fm0(f23, f28, globalState)};
		if (--(|v74| - |v75|) == f23) {
			if (v75 > v75) {
				globalState.f7 := false;
				var v76: array<int> := new int[21](i7 => safeDivisionInt(i7, f23));
				var v77: array<array<int>> := new array<int>[20] [v76, v76, v76, v76, v76, v76, v76, v76, v76, v76, v76, v76, v76, v76, v76, if (f28) then v76 else v76, v76, v76, v76, v76];
				v77[safeIndex(71, v77.Length)] := v76;
				v77[safeIndex(71, v77.Length)] := v76;
				var v78: multiset<int> := multiset{safeDivisionInt(f23, f23)};
				globalState.f8 := if (f23 in v78) then v78[f23] else f23;
				globalState.f21 := f28;
				var v79: multiset<array<int>> := multiset{v76, v76, v76};
				v79 := v79;
			} else {
				var v80 := DC19(f29[safeIndex(f23, |f29|) := f23]);
				var v81: multiset<seq<int>> := multiset{f29, v80.cf32};
				globalState.f21 := v81 == v81;
				globalState.f6 := fm2(globalState);
				v0 := p0;
				var v82: array<int> := new int[21];
				v82[safeIndex(127, v82.Length)] := f23;
				v82[safeIndex(127, v82.Length)] := f23;
				var v83 := DC11(f23);
				var v84: map<set<int>, D3> := map[v75 := v83];
				var v86: seq<set<int>> := [v75];
				v84, v82[safeIndex(127, v82.Length)] := (map v85 : set<int> | v85 in v86 :: (v85) := (v83)) + v84, safeModuloInt(-|(v36 + "udxyoam")|, v82[safeIndex(127, v82.Length)]);
			}
			
			globalState.f21 := f23 == 805;
			var v87: seq<bool> := [p1];
			p0[safeIndex(347, p0.Length)] := v87[safeIndex(f23, |v87|) := f28] < [true, f28];
			p0[safeIndex(347, p0.Length)] := p1;
			var v89: map<bool, map<char, int>> := map[true := map v88 : char | v88 in v36 :: (v88) := (f23)];
			var v90: set<bool> := {p0[safeIndex(347, p0.Length)]};
			var v91: map<int, seq<int>> := map[fm0(|map[v37 := |v89|]|, true, globalState) := (f29[safeIndex(|v90|, |f29|) := f23])[safeIndex(f23, |f29[safeIndex(|v90|, |f29|) := f23]|) := 0x3af]];
			var v92: array<int> := new int[17];
			var v93: map<array<int>, seq<int>> := map[v92 := f29];
			var v94: C1 := new C1(false, if (v92 in v93) then v93[v92] else f29, f23);
			var v95: map<seq<int>, C1> := map[if (f23 in v91) then v91[f23] else [|v91|] := v94];
			v95 := v95[f29 + [f23] := v94];
			globalState.f21 := true;
		} else {
			var v96: seq<bool> := [f28, false, p1];
			r0, v96, globalState.f9 := f23, v96, p1;
			var v97: map<int, bool> := map[f23 := f28];
			v97 := v97[|v36| := !true];
			var v98: array<array<char>> := new array<char>[27];
			var v99: array<char> := new char[28](i8 => v37);
			v98[safeIndex(139, v98.Length)] := v99;
			globalState.f14, v98[safeIndex(139, v98.Length)], globalState.f9, globalState.f14 := -f29[safeIndex(f23, |f29|)], v99, !p1, f23;
			var v100: map<int, int> := map[f23 := f23];
			var v101 := new C2(v100, p1, !p1 || p1, [f23, f23] + f29, if (f28) then f23 else f23);
			globalState.f8 := f23;
		}
		
		r0 := f23;
		var v102: multiset<int> := multiset{f23};
		var v103: map<int, int> := map[f23 := f23];
		var v104: map<char, int> := map[v37 := if (|v103| in v102) then v102[|v103|] else f23];
		r1 := v75 - {|(set v105 : char | v105 in v104 :: (v105))|, f23, |v36|};
	}
	method m2(p0: string, globalState: GlobalState) returns (r0: int, r1: seq<int>) {
		var v0: array<int> := new int[26](i0 => safeDivisionInt(i0, f23));
		var v1: seq<array<int>> := [v0];
		var v2 := DC3(v1);
		var v3: multiset<D1> := multiset{v2, v2, v2};
		var v4: map<seq<bool>, multiset<D1>> := map[[false] := v3];
		var v5: seq<bool> := [f28];
		v4 := v4[v5 := v3];
		v0[safeIndex(260, v0.Length)] := f23;
		var v6: seq<string> := [seq(abs(0x13d), i1  => ('m'))];
		var v7 := DC17(v6);
		v0[safeIndex(260, v0.Length)] := match if (f28) then v7 else DC17(v6) {
			case DC17(cf30) => f23
		};
		var i2 := 0;
		while (f28 ==> (f23 in f29))
			decreases 100 - i2
		{
			if (i2 >= 100) {
				break;
			}
			
			i2 := i2 + 1;
			var v8: multiset<bool> := multiset{f28};
			var v9 := 'g';
			var v10 := DC8(f23, |v8|, v9, |p0|, f28);
			globalState.f7, globalState.f8, globalState.f21 := !!v10.cf19, f23 + (0x2b4 + v0[safeIndex(260, v0.Length)]), v8[f28 := abs(v0[safeIndex(260, v0.Length)])] != v8;
			globalState.f21 := false;
			globalState.f11 := v9;
			globalState.f8 := v0[safeIndex(260, v0.Length)];
		}
		var v11: array<char> := new char[15](i3 => 'p');
		var v12: map<array<char>, bool> := map[v11 := f28];
		v12 := v12[v11 := f28];
		if (false) {
			var v14: map<int, bool> := map[|v5| := f28];
			var v15: seq<seq<int>> := [[|v14|]];
			var v16 := new C2(map v13 : int | v13 in v15[safeIndex(f23, |v15|)] :: (v13 - f23) := (f23), f28, f28, f29, v0[safeIndex(260, v0.Length)]);
			var v17: multiset<bool> := multiset{v16.f32, v16.f32, f28};
			globalState.f6 := v6[safeIndex(|v17|, |v6|)];
			var v18: array<string> := new string[17](i4 => p0);
			v18 := v18;
			var v19: array<bool> := new bool[21];
			var v20: map<array<bool>, bool> := map[v19 := 0x3b1 >= --0x22b];
			globalState.f21 := if (v19 in v20) then v20[v19] else f28;
			v19[safeIndex(409, v19.Length)] := v16.f32;
			var v21: map<bool, bool> := map[v16.f32 := v16.f32];
			var v22: map<bool, int> := map[f28 := |v21|];
			var v23: multiset<int> := multiset{39, v0[safeIndex(260, v0.Length)]};
			v19[safeIndex(409, v19.Length)] := (|v22| - v0[safeIndex(260, v0.Length)]) != (if (f23 in v23) then v23[f23] else |v5|);
		} else {
			var v24: map<int, int> := map[-|v6| := 0x3f];
			var v25: T1 := new C2(v24, f28, f28, f29, v0[safeIndex(260, v0.Length)]);
			var v26 := DC16(v25);
			match (v26) {
				case DC16(cf29) =>
					var v27: multiset<int> := multiset{v25.f23, cf29.f23, -|v5|};
					var v28: multiset<multiset<int>> := multiset{v27, multiset(f29), v27 + v27, v27[v0[safeIndex(260, v0.Length)] := abs(v25.f23)], v27};
					var v29: multiset<array<char>> := multiset{v11, v11, v11};
					r0, v28 := safeModuloInt(-f23, f23 + |p0|), multiset{(multiset{v25.f23, |v29|, v25.f23, v0[safeIndex(260, v0.Length)], |p0|})[v25.f23 := abs(cf29.f23)], v27, v27 * v27, v27};
					var v30: array<multiset<bool>> := new multiset<bool>[27](i5 => multiset{false, !v25.f28, cf29.f28, cf29.f28} + multiset{v25.f28, false, !false, v25.f28});
					globalState.f21, v30 := v25.f28, v30;
					var v31 := DC11(v0[safeIndex(260, v0.Length)]);
					var v32: map<seq<D3>, int> := map[[v31.(cf22 := v25.f23)] + (seq(abs(0x1ea), i6  => (v31))) := -0xad];
					var v33: seq<D3> := [DC11(cf29.f23), v31];
					v32 := v32[if (cf29.f28) then v33 else v33 := v0[safeIndex(260, v0.Length)]];
					var v34: multiset<array<int>> := multiset{v0, v0, v0};
					globalState.f7 := v34 > (multiset{v0})[v1[safeIndex(cf29.f23, |v1|)] := abs(v25.f23)];
			}
			
			var v35: map<array<int>, seq<int>> := map[v0 := v25.f29];
			var v36: map<bool, map<array<int>, seq<int>>> := map[false := v35[v0 := f29]];
			v36 := v36[v25.f28 := v35];
			if (!fm1(v25.f23, v25.f28, globalState)) {
				globalState.f7 := v25.f28;
				v11 := v11;
				globalState.f9 := false;
				v25 := v25;
				var v37: array<bool> := new bool[5] [v25.f28, v25.f28, fm1(-v0[safeIndex(260, v0.Length)], v25.f28, globalState), v25.f28, fm1(v0[safeIndex(260, v0.Length)], v25.f28, globalState)];
				v37[safeIndex(829, v37.Length)] := f28;
				var v38 := 'a';
				var v39 := DC24(v11);
				globalState.f21, globalState.f14, v11, v37[safeIndex(829, v37.Length)], v37 := v38 !in p0, -0x2e5, v39.cf40, (seq(abs(729), i7  => (v38))) == p0, v37;
			} else {
				var v40: map<seq<bool>, bool> := map[v5 := f28];
				var v42 := 'w';
				var v43: map<char, bool> := map[v42 := v25.f28];
				v40 := if ((map v41 : int | v41 in (seq(abs(0x302), i8  => (|(seq(abs(0xb1), i9  => ('y')))|))) :: (safeModuloInt(v41, v25.f23)) := (v0[safeIndex(260, v0.Length)])) == fm16(f28, |v43|, f23, globalState)) then map[v5 := f28] else v40;
				var v44 := DC26(0x224);
				var v45: array<D9> := new D9[22] [DC26(v0[safeIndex(260, v0.Length)]), v44, v44, v44, v44, v44, v44, v44, v44, v44, v44, v44, v44, DC26(0x25c), v44.(cf43 := f23), v44, v44, v44, v44.(cf43 := v0[safeIndex(260, v0.Length)]), v44, v44, v44];
				v45[safeIndex(94, v45.Length)] := v44;
				v45[safeIndex(94, v45.Length)] := v44;
				v0[safeIndex(260, v0.Length)] := 0x2d1;
				var v46: array<array<multiset<bool>>> := new array<multiset<bool>>[9];
				var v47: multiset<bool> := multiset{true};
				var v48: seq<set<bool>> := [{f28}];
				var v49 := DC0(p0);
				var v50: multiset<int> := multiset{|p0|};
				var v51: array<multiset<bool>> := new multiset<bool>[22] [v47, v47, v47, v47, v47, v47, v47, multiset{v25.f28, f28}, multiset{v25.f28, f28}, v47, v47, v47, (f30[safeIndex(f23, |f30|)])[f28 := abs(|v48|)], v47, v47, v47[v25.f28 := abs(v25.f23)], multiset{v25.f28, v25.f28}, fm4(v49, f23, v5, v42, globalState), multiset(v5), multiset{v25.f28, f28}, v47, fm4(v49, if (|v47| in v50) then v50[|v47|] else v25.f23, v5, v42, globalState)];
				v46[safeIndex(491, v46.Length)] := v51;
				var v52: map<int, bool> := map[v25.f23 := f28];
				var v53: array<bool> := new bool[27](i10 => false);
				var v54: map<array<bool>, bool> := map[v53 := fm1(v25.f23, f28, globalState)];
				globalState.f21, globalState.f8, v46[safeIndex(491, v46.Length)] := |v52| >= |v54|, v25.f23, v51;
				var v55, v56, v57 := m0(v53, p0 + p0[safeIndex(f23, |p0|) := v42], v25.f23, globalState);
			}
			
			var v58 := DC22(f28);
			globalState.f8 := if (v58.cf38) then f23 else v0[safeIndex(260, v0.Length)];
			globalState.f14 := f23;
		}
		
		var v59 := new C0();
		var v60: map<string, int> := map["rtk" := f23];
		var v61: map<int, string> := map[0x1f3 := "gun"];
		var v62 := 'a';
		var v63: set<int> := {f23, -0x26c, f23, v0[safeIndex(260, v0.Length)], f23};
		var v64: map<int, int> := map[if ((if (-0xbf in v61) then v61[-0xbf] else p0)[safeIndex(f23, |(if (-0xbf in v61) then v61[-0xbf] else p0)|) := v62] in v60) then v60[(if (-0xbf in v61) then v61[-0xbf] else p0)[safeIndex(f23, |(if (-0xbf in v61) then v61[-0xbf] else p0)|) := v62]] else |v63| := -v0[safeIndex(260, v0.Length)]];
		r0 := if (-f23 in v64) then v64[-f23] else v0[safeIndex(260, v0.Length)];
		r1 := (f29 + f29)[safeIndex(v0[safeIndex(260, v0.Length)], |(f29 + f29)|) := v0[safeIndex(260, v0.Length)]];
	}
	method m5(globalState: GlobalState) {
		var v0: array<int> := new int[1];
		forall i0 | 0 <= i0 < v0.Length {
			v0[i0] := safeModuloInt(i0, |f29| - |"fgr"|);
		}
		var i1 := 0;
		while (!f28)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			var v1 := new C0();
			var v2: array<bool> := new bool[23];
			var v3: set<int> := {-f23};
			var v4: map<int, int> := map[f23 := f23];
			v2[safeIndex(174, v2.Length)] := v3 < {if (f29[safeIndex(f23, |f29|)] in v4) then v4[f29[safeIndex(f23, |f29|)]] else f23};
			v2[safeIndex(174, v2.Length)] := !(if (!f28) then (set v5 : int | (0x3d3 <= v5) && (v5 < 0x2af) :: (v5 + f23)) !! v3 else if (f28) then f28 else f28);
			var v6: set<bool> := {false, f28};
			var v7: seq<set<bool>> := [v6];
			globalState.f14, v6, globalState.f8 := f23, v7[safeIndex(f23, |v7|)], f29[safeIndex(f23, |f29|)];
			v2[safeIndex(174, v2.Length)] := f28;
		}
		var v8 := "kgw";
		var v9: multiset<int> := multiset{|v8|, f23};
		var v10: map<int, seq<int>> := map[f23 := f29];
		for i2 := |([|fm17(f28, fm0(f23, f28, globalState), globalState)|, f23] + f29)| to if (0x2ad in v9) then v9[0x2ad] else |(if (0x104 in v10) then v10[0x104] else [f23, f23])| {
			var v11: seq<bool> := [f28];
			var v12: multiset<bool> := multiset{!f28, f28};
			var v13: map<bool, bool> := map[false := true];
			var v14: map<bool, map<bool, bool>> := map[f28 := v13];
			var v15: set<int> := {i2, safeModuloInt(f23, |v9|), i2 - |v11|, |v12[f28 := abs(i2)]| * f23, |v14|};
			v15 := if (false) then v15 else v15;
			var v16 := 'e';
			var v17: map<int, int> := map[|f29| := |[(v8[safeIndex(f23, |v8|) := v16])[safeIndex(f23, |v8[safeIndex(f23, |v8|) := v16]|) := v16], v8]|];
			var v18 := DC14(v0);
			var v19 := new C2(v17, fm1(i2, f28, globalState), f28, f29[safeIndex(i2, |f29|) := i2] + f29, |[v18, v18]|);
			v17 := v17[|v8| := (fm24(v19.f32, i2, globalState)).cf43];
			if (f28) {
				v0[safeIndex(888, v0.Length)] := i2;
				var v20: array<char> := new char[6] ['n', v16, v16, v16, v16, v16];
				var v21 := DC24(v20);
				var v22: array<array<char>> := new array<char>[19] [v20, v20, v20, v20, v20, v20, v20, v20, v20, v20, v20, v20, v20, v20, v20, v20, v21.cf40, v20, v20];
				var v23: array<seq<int>> := new seq<int>[11];
				v23[safeIndex(742, v23.Length)] := f29;
				v0[safeIndex(888, v0.Length)], v22, v23[safeIndex(742, v23.Length)], globalState.f21 := 0x3aa, v22, f29, f28;
				var v24 := DC5(i2, v19.f32, 0x11c, f23);
				globalState.f9 := v24.cf10;
				globalState.f11 := v16;
				var v25: map<int, bool> := map[f23 := v19.f32];
				var v26: map<bool, int> := map[f28 := |v25|];
				globalState.f14 := v19.fm18(if (f28) then v26 else fm25(|multiset{v16}|, v19.f32, globalState), (DC0("lvnjskw").(cf0 := v8[safeIndex(-v0[safeIndex(888, v0.Length)], |v8|) := v16])).cf0 + (seq(abs(387), i3  => (v16))), v19.f32, globalState);
				globalState.f7 := v19.f32;
			} else {
				var v27: map<bool, int> := map[f28 := i2];
				var v28 := DC18(v27);
				v28 := DC18(v27);
				f29 := [v19.fm18(v27, "epl", v11[safeIndex(i2, |v11|)], globalState), fm0(i2, v19.f32, globalState), f23];
				v0[safeIndex(661, v0.Length)] := -|"pnttsu"|;
				globalState.f8, v0[safeIndex(661, v0.Length)], globalState.f14, globalState.f14 := safeModuloInt(|v8|, f23), safeModuloInt(f23, safeDivisionInt(f23, |"bqnxnel"|)), f23, if (fm1(|v19.f31|, v19.f32, globalState)) then safeDivisionInt(f23, i2) else i2;
				var v29: array<map<string, bool>> := new map<string, bool>[8];
				var v30: map<string, bool> := map[v8 := v19.f32];
				v29[safeIndex(532, v29.Length)] := if (v19.f32) then v30 else map["pinimjpwv" := f28];
				var v31: array<int> := new int[25](i4 => i4 * v0[safeIndex(661, v0.Length)]);
				v29[safeIndex(532, v29.Length)], v31, globalState.f14 := v30[v8 := !v19.f32], v31, safeDivisionInt(i2, -v0[safeIndex(661, v0.Length)]);
				var v32: map<bool, char> := map[fm1(564, f28, globalState) := v16];
				var v33 := DC29(map[v19.f32 := v16]);
				var v34: multiset<map<bool, char>> := multiset{v32 + v32, v33.cf50, if (fm1(i2, v19.f32, globalState)) then v32 else v32};
				v34 := v34;
			}
			
		}
		globalState.f14 := f23;
		globalState.f14 := f23;
		globalState.f9 := if (f28) then f28 else f28;
	}
}

class C4 extends T0 {
	const f27 : array<int>
	constructor (f27 : array<int>, f23 : int) {
		this.f27 := f27;
		this.f23 := f23;
	}
	
	function fm10(p0: string, p1: bool, p2: map<bool, int>, p3: bool, globalState: GlobalState): bool {
		({false, false} - {true, false}) >= ({false} - {true, !true, false})
	}
	function fm11(globalState: GlobalState): seq<bool> {
		[true] + (if (true) then [false] else [false])
	}
	method m1(p0: array<bool>, p1: bool, globalState: GlobalState) returns (r0: int, r1: set<int>) {
		var v0: array<D1> := new D1[15];
		var v1: seq<multiset<array<D1>>> := [multiset{v0}];
		var v2 := "ru";
		var v3: map<bool, multiset<array<D1>>> := map[p1 := v1[safeIndex(|v2|, |v1|)]];
		var v4: multiset<array<D1>> := multiset{v0};
		var v5 := DC9(v4);
		r0, globalState.f8, r0, globalState.f14, globalState.f6 := |(if (p1 in v3) then v3[p1] else v5.cf20)|, f23, f23, safeModuloInt(f23, if (p1) then fm0(f23, p1, globalState) else f23), ("taruujebt" + v2) + v2;
		var v6: map<bool, int> := map[p1 := f23];
		var v7 := DC18(v6);
		var v8: multiset<D7> := multiset{v7, v7, v7};
		var v9: seq<int> := [|v8|];
		var v10 := new C1(fm1(f23, true, globalState), v9 + v9, f23);
		var v11: array<seq<int>> := new seq<int>[15](i0 => v9);
		v11[safeIndex(122, v11.Length)] := [f23] + v9;
		p0[safeIndex(428, p0.Length)] := f23 <= f23;
		v11[safeIndex(122, v11.Length)], globalState.f7, globalState.f21, p0[safeIndex(428, p0.Length)], globalState.f21 := DC27(p1, v9, p0, p1, p1).cf45, p1, p1, p1, f23 > (f23 + f23);
		var i1 := 0;
		while (p0[safeIndex(428, p0.Length)])
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			globalState.f8 := f23;
			var v12: array<D5> := new D5[11];
			var v13 := DC10(p0[safeIndex(428, p0.Length)]);
			v12, globalState.f8, v13 := v12, f23, DC10(p0[safeIndex(428, p0.Length)]).(cf21 := -0x271 > 856);
			var v14: set<array<int>> := {f27};
			var v15: map<set<array<int>>, bool> := map[v14 := p1];
			globalState.f21 := if (v14 in v15) then v15[v14] else true;
			var v16: array<int> := new int[8];
			v16 := v16;
		}
		var v17: array<map<set<bool>, int>> := new map<set<bool>, int>[16](i2 => map[{false, p1, p0[safeIndex(428, p0.Length)]} := f23] + map[{!p0[safeIndex(428, p0.Length)], p0[safeIndex(428, p0.Length)], true, p0[safeIndex(428, p0.Length)], p1} := f23]);
		var v18: set<bool> := {!p0[safeIndex(428, p0.Length)]};
		v17[safeIndex(476, v17.Length)] := map[v18 := f23];
		var v20: map<set<bool>, seq<int>> := map[v18 := v9];
		v17[safeIndex(476, v17.Length)] := map v19 : set<bool> | v19 in v20 :: (v19) := (f23);
		if (p0[safeIndex(428, p0.Length)]) {
			var v21: map<int, int> := map[f23 := f23];
			r0 := fm0(if (|v2| in v21) then v21[|v2|] else f23, false, globalState);
			globalState.f9 := p0[safeIndex(428, p0.Length)];
			var v22: multiset<bool> := multiset{p0[safeIndex(428, p0.Length)], !p1, !p1};
			var v23: seq<multiset<bool>> := [v22];
			var v24: T0 := new C3(v23, fm0(f23, p1, globalState), p0[safeIndex(428, p0.Length)], v11[safeIndex(122, v11.Length)]);
			var v25: map<T0, bool> := map[v24 := p0[safeIndex(428, p0.Length)]];
			var v26 := new C3(if (p0[safeIndex(428, p0.Length)]) then v23 else v23, -f23 - |v25|, p0[safeIndex(428, p0.Length)], v11[safeIndex(122, v11.Length)]);
			var v29: seq<map<int, int>> := [v21, v21];
			var v30: set<int> := {f23, |v21|};
			var v31: array<map<int, int>> := new map<int, int>[16] [v21, v21, v21, v21, if (fm1(if (p0[safeIndex(428, p0.Length)] in v22) then v22[p0[safeIndex(428, p0.Length)]] else 0x28b, p0[safeIndex(428, p0.Length)], globalState)) then v21 else v21, v21, v21, v21, v21, v21, v21 + map[|v21| := v24.f23], map v27 : int | v27 in (map v28 : int | (118 <= v28) && (v28 < 492) :: (safeDivisionInt(v28, |{DC10(!p1)}|)) := (multiset{|v18|})) :: (v27 - v24.f23) := (v24.f23), v29[safeIndex(|v2|, |v29|)], (map[-v24.f23 := v24.f23])[|v30| := f23], v21, v21];
			var v32: seq<bool> := [fm10(seq(abs(868), i3  => ('e')), false, v6, p0[safeIndex(428, p0.Length)], globalState), p1, p1, p1, !p0[safeIndex(428, p0.Length)]];
			f27[safeIndex(32, f27.Length)] := -fm0(v11[safeIndex(122, v11.Length)][safeIndex(|v32|, |v11[safeIndex(122, v11.Length)]|)], p0[safeIndex(428, p0.Length)], globalState);
			var v33 := DC30(v31);
			globalState.f14, v31, f27[safeIndex(32, f27.Length)], globalState.f14 := v24.f23, v33.cf51, f23, f23;
			globalState.f14 := |v9|;
		} else {
			var v34: map<int, char> := map[fm0(f23, p0[safeIndex(428, p0.Length)], globalState) := fm6(-f23, globalState)];
			globalState.f8 := -|v34|;
			var v35: map<int, int> := map[f23 := f23];
			var v36 := new C2(v35, p0[safeIndex(428, p0.Length)], |v2| != -f23, v9 + (seq(abs(397), i4  => (f23))), f23);
			var v37: array<bool> := new bool[21];
			v37 := p0;
			var v38: set<C2> := {v36};
			var v39: seq<bool> := [p1, p0[safeIndex(428, p0.Length)], fm1(f23, v36.f32, globalState), p1, v38 !! v38];
			if (v39[safeIndex(-f23, |v39|)]) {
				p0[safeIndex(428, p0.Length)] := f23 < f23;
				var v40: array<map<bool, int>> := new map<bool, int>[24] [v6, map[v36.f32 := f23], map[p0[safeIndex(428, p0.Length)] := f23], v6 + v6, v6, v6[v36.f32 := f23], v6, v6, if (v36.f32) then v6 else map[v36.f32 := f23], v6, v6 + v6, v6 + v6, v6[false := f23] + v6, v6, v6 + v6, (map[p1 := f23])[v36.f32 := |v2|], v6, v6, v6, v6[p0[safeIndex(428, p0.Length)] := |v6|], v6, v6, v6, map[p0[safeIndex(428, p0.Length)] := f23]];
				v40[safeIndex(449, v40.Length)] := v6;
				var v41 := 'v';
				var v42: map<char, bool> := map[v41 := p1];
				var v43 := DC34(v42);
				v40[safeIndex(449, v40.Length)] := map[p1 := |v43.cf55|] + v6;
				f27[safeIndex(259, f27.Length)] := f23;
				var v44: map<int, bool> := map[f23 := p0[safeIndex(428, p0.Length)]];
				var v45 := DC21(v36.f32, p0[safeIndex(428, p0.Length)], |v44|);
				f27[safeIndex(259, f27.Length)] := safeDivisionInt(0x31e, v45.cf37 * f23);
				var v46, v47, v48, v49 := v36.m6(v18 >= {true, true, v36.f32}, globalState);
				p0[safeIndex(428, p0.Length)] := p1;
			} else {
				p0[safeIndex(428, p0.Length)] := f23 > (33 * 0x7f);
				f27[safeIndex(914, f27.Length)] := f23;
				f27[safeIndex(914, f27.Length)] := safeDivisionInt(f23, safeModuloInt(f23, f23));
				var v50 := 'n';
				globalState.f21 := v2[safeIndex(f27[safeIndex(914, f27.Length)], |v2|) := v50] <= (v2 + v2);
				var v51 := DC11(453);
				p0[safeIndex(428, p0.Length)] := v39[safeIndex(-(if (v36.f32) then f27[safeIndex(914, f27.Length)] else v51.cf22), |v39|)];
				globalState.f21 := false;
			}
			
			var v52: set<int> := {f23};
			var v53 := new C2(v36.f31[f23 := 797], false, fm10(v2, p1, v6[p0[safeIndex(428, p0.Length)] := f23], fm1(-f23, v36.f32, globalState), globalState), v9, safeModuloInt(-|v52|, -f23));
		}
		
		r0 := f23;
		var v54: set<int> := {|v6|, f23};
		r1 := v54 + (v54 * v54);
	}
	method m2(p0: string, globalState: GlobalState) returns (r0: int, r1: seq<int>) {
		globalState.f8 := f23;
		var v0 := false;
		var i0 := 0;
		while ((if (false) then fm10(p0, true, map[v0 := f23], v0, globalState) else v0) || v0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			f27[safeIndex(562, f27.Length)] := safeDivisionInt(f23, f23);
			var v1: seq<bool> := [v0];
			var v2: array<int> := new int[3] [|("uovdktobe" + p0)|, safeModuloInt(|v1|, f23), f23];
			var v3: seq<int> := [f23];
			var v4: map<int, int> := map[f23 := |v3|];
			var v5: set<bool> := {v0};
			var v6: map<bool, int> := map[!v0 := |DC38(v5).cf62|];
			var v7: C2 := new C2(v4, fm10(p0, v0, v6, true, globalState), v0, v3, 0x1d1);
			f27[safeIndex(562, f27.Length)], globalState.f21, v2, v7 := f23 - f23, v0, f27, if (v0) then v7 else v7;
			var v8 := 'n';
			var v9: set<int> := {DC20(v0, f27[safeIndex(562, f27.Length)]).cf34, |(("dxmrtsdjn")[safeIndex(f27[safeIndex(562, f27.Length)], |"dxmrtsdjn"|) := v8])[safeIndex(0x200, |("dxmrtsdjn")[safeIndex(f27[safeIndex(562, f27.Length)], |"dxmrtsdjn"|) := v8]|) := v8]|};
			var v10: map<bool, bool> := map[v0 := v0];
			var v11: map<int, bool> := map[|v9| := if (v7.f32 in v10) then v10[v7.f32] else v0];
			var v12: array<bool> := new bool[12] [true, v7.f32, true, v7.f32, v0, v7.f32 || v0, v1 != v1, v0, v0 <==> v0, if (v7.f32) then v7.f32 else v7.f32, false, f27[safeIndex(562, f27.Length)] == |[v0, v0, v7.f32, if (-f23 in v11) then v11[-f23] else v7.f32]|];
			v12 := v12;
			var v13 := DC27(|v9| != f23, v3, v12, v0, v7.f32);
			var v14 := DC18(map[v7.f32 := 0x3bd]);
			f27[safeIndex(562, f27.Length)], v13 := fm0(|p0|, if (f27[safeIndex(562, f27.Length)] in v11) then v11[f27[safeIndex(562, f27.Length)]] else if (v7.f32 in v10) then v10[v7.f32] else fm10(p0, v0, v14.cf31, v0, globalState), globalState), v13;
			var v15: map<int, seq<bool>> := map[f27[safeIndex(562, f27.Length)] := (v1 + v1)[safeIndex(|p0|, |(v1 + v1)|) := v7.f32]];
			globalState.f9, v15, r0 := f23 == (|multiset{f27[safeIndex(562, f27.Length)]}| + f27[safeIndex(562, f27.Length)]), (map[f27[safeIndex(562, f27.Length)] := v1] + v15) + v15[|v1| := fm11(globalState)], -f23;
		}
		var v16: set<bool> := {!v0, !!v0};
		globalState.f8 := -match fm26(v16, globalState) {
			case DC1(cf1, cf2, cf3) => cf3 - cf1
			case DC0(cf0) => 0xb - f23
			case DC2(cf4) => DC26(f23).cf43
		};
		var i1 := 0;
		while (!v0)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			if (v0) {
				var v17 := 's';
				var v18: multiset<bool> := multiset{v0, v0, false, v0 || fm1(|[v17]|, v0, globalState), f23 >= f23};
				v18 := multiset{v0, v0} - v18;
				var v19: map<bool, char> := map[f23 == f23 := v17];
				var v20: seq<int> := [f23, 419, f23];
				var v21: map<bool, seq<int>> := map[v0 := seq(abs(-174), i2  => (f23))];
				var v22 := DC36(v0, true, f23, f23, v21);
				var v23: array<bool> := new bool[26] [0x255 >= f23, v0, v0, v0, v0 <== v0, true, v0 <== v0, v0, true, v0, f23 > -f23, v0, !v0, v0, v0, v17 !in p0, v0, v0, !(v20 == v20), v0 || v0, v22.cf57, v0, !true, v0, v0, v0];
				v23[safeIndex(36, v23.Length)] := !v0;
				var v24: array<char> := new char[6](i3 => 'o');
				v24[safeIndex(605, v24.Length)] := v17;
				v19, v23[safeIndex(36, v23.Length)], v23, v24[safeIndex(605, v24.Length)] := v19 + v19, v0 ==> !v0, v23, v17;
				v23[safeIndex(36, v23.Length)] := fm1(f23, !v23[safeIndex(36, v23.Length)], globalState) ==> v0;
				var v25: map<bool, bool> := map[v0 := false];
				var v26: seq<multiset<bool>> := [v18, v18, (v18[v23[safeIndex(36, v23.Length)] := abs(|v25|)])[v23[safeIndex(36, v23.Length)] := abs(0x2a6)], v18, v18];
				var v27 := new C3(v26, f23, !v23[safeIndex(36, v23.Length)], v20);
				globalState.f21 := v0;
			} else {
				var v28 := 'x';
				globalState.f11 := v28;
				var v29: array<array<bool>> := new array<bool>[15];
				v29 := v29;
				var v31: map<bool, int> := map[v0 := f23];
				var v32: set<int> := {fm0(f23, v0, globalState), f23, f23, fm0(-f23, fm10(seq(abs(0x214), i4  => (v28)), v0, v31, v0, globalState), globalState), f23};
				var v33: array<set<int>> := new set<int>[2] [(set v30 : int | (0x18d <= v30) && (v30 < -0xe7) :: (v30 - f23)) + {f23}, v32];
				v33[safeIndex(754, v33.Length)] := v32;
				v33[safeIndex(754, v33.Length)] := v32;
				var v34: map<int, bool> := map[|[!v0]| := v0];
				v34 := v34 + v34;
				var v35: map<int, int> := map[f23 := f23];
				v35 := v35[fm0(-0xf1, !v0, globalState) := f23];
			}
			
			var v36: seq<int> := [f23];
			var v37 := DC20(v0, f23);
			var v38: seq<bool> := [v0, v0, v37.cf33, v0, v0];
			var v39: array<bool> := new bool[18] [fm1(f23, true, globalState), v0, v16 >= v16, v0, !v0, |v36| >= f23, v0, v38[safeIndex(|"hm"|, |v38|)], v0, v0, v0, v0, false, v0, fm1(|p0|, v0, globalState), !!!v0, v0 ==> v0, v0];
			v39[safeIndex(808, v39.Length)] := v0;
			v39[safeIndex(808, v39.Length)] := !v0;
			if (v0) {
				var v40: map<seq<int>, bool> := map[v36 := v0];
				v40 := (v40 + v40) + v40;
				var v41: map<bool, bool> := map[fm1(f23, v39[safeIndex(808, v39.Length)], globalState) := DC1(0x3c7, v39[safeIndex(808, v39.Length)], f23).cf3 <= 872];
				globalState.f7 := if ((-f23 <= f23) in v41) then v41[-f23 <= f23] else f23 <= 0x155;
				globalState.f14 := fm0(f23, 0x50 <= f23, globalState);
				var v42: map<bool, int> := map[v39[safeIndex(808, v39.Length)] := f23];
				var v43 := DC18(v42 + v42);
				v43 := v43;
				globalState.f21 := true ==> v0;
			} else {
				var v44, v45, v46 := m0(v39, p0, f23, globalState);
				var v47: array<seq<int>> := new seq<int>[17];
				var v48: C1 := new C1(v39[safeIndex(808, v39.Length)], v36, f23);
				var v49: map<C1, array<seq<int>>> := map[v48 := v47];
				var v50: array<array<seq<int>>> := new array<seq<int>>[23] [v47, v47, v47, v47, v47, if (v39[safeIndex(808, v39.Length)]) then v47 else if (v48 in v49) then v49[v48] else v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47, v47];
				v50[safeIndex(507, v50.Length)] := v47;
				var v51 := 'r';
				var v52: multiset<char> := multiset{v51, 'o', v51};
				f27[safeIndex(136, f27.Length)] := if (v51 in v52) then v52[v51] else |[v39[safeIndex(808, v39.Length)]]|;
				v50[safeIndex(507, v50.Length)], f27[safeIndex(136, f27.Length)] := v47, -v46 + (if (v39[safeIndex(808, v39.Length)]) then f23 else f23);
				var v53: multiset<int> := multiset{f23};
				var v54: seq<string> := [p0, "sxaeugrwk"];
				var v55: map<int, int> := map[(if (|multiset{v46, f27[safeIndex(136, f27.Length)]}| in v53) then v53[|multiset{v46, f27[safeIndex(136, f27.Length)]}|] else f23) - f27[safeIndex(136, f27.Length)] := |v54[safeIndex(f27[safeIndex(136, f27.Length)], |v54|)]|];
				v55 := v55;
				var v56: array<map<map<int, bool>, bool>> := new map<map<int, bool>, bool>[25](i5 => map[v45 := v0]);
				var v57: map<map<int, bool>, bool> := map[v45 := false];
				v56[safeIndex(452, v56.Length)] := v57;
				var v58: seq<seq<string>> := [seq(abs(0x3cf), i6  => ("o"))];
				var v59: seq<set<int>> := [fm21((v58[safeIndex(v46, |v58|)])[safeIndex(v46, |v58[safeIndex(v46, |v58|)]|) := p0], globalState)];
				v56[safeIndex(452, v56.Length)], globalState.f21, v59, globalState.f8 := map[v45 := v0], false, v59, f23;
				globalState.f14 := fm0(DC1(v46, !v39[safeIndex(808, v39.Length)], v46).cf3, !(f27[safeIndex(136, f27.Length)] <= v46), globalState);
			}
			
			if (true) {
				var v60: array<map<int, int>> := new map<int, int>[13];
				var v61 := DC33(DC30(v60));
				var v62 := DC30(v60);
				var v63: array<D11> := new D11[18] [v61, v61, v61, v61, v61, v61, v61, v61, DC33(v62), v61, v61, v61, v61, v61, v61, v61, v61, v61];
				var v64: map<int, array<D11>> := map[-safeModuloInt(|v36|, -v36[safeIndex(0x12f, |v36|)]) := v63];
				v64 := v64[0x375 := v63];
				var v65: C1 := new C1(v39[safeIndex(808, v39.Length)], v36, f23);
				var v66: seq<C1> := [v65];
				var v67: map<bool, bool> := map[v39[safeIndex(808, v39.Length)] := fm1(|v66|, fm1(|v16|, v39[safeIndex(808, v39.Length)], globalState), globalState)];
				var v68: map<int, set<bool>> := map[|v67| := v16];
				var v69: map<map<int, set<bool>>, int> := map[v68[f23 := v16] + v68 := f23];
				globalState.f14 := if ((if (v39[safeIndex(808, v39.Length)]) then v68 else map[f23 := v16]) in v69) then v69[if (v39[safeIndex(808, v39.Length)]) then v68 else map[f23 := v16]] else f23;
				var v70: array<C2> := new C2[10];
				var v71: map<int, int> := map[f23 := f23];
				var v72: map<bool, int> := map[v39[safeIndex(808, v39.Length)] := |"eys"|];
				var v73: multiset<int> := multiset{|(seq(abs(-0xfb), i7  => (-0x1bd)))|};
				var v74: C2 := new C2(v71, v0, v39[safeIndex(808, v39.Length)], v36, -(if (fm10("bms", v0, v72, true, globalState) in v72) then v72[fm10("bms", v0, v72, true, globalState)] else |v73|));
				v70[safeIndex(156, v70.Length)] := v74;
				var v75 := 'o';
				var v76: map<C2, int> := map[v74 := |map[!v0 := v75]|];
				globalState.f9, globalState.f21, globalState.f14, v70[safeIndex(156, v70.Length)] := !v0, v74.f32, safeDivisionInt(f23, if (v74 in v76) then v76[v74] else f23), v74;
				v39[safeIndex(808, v39.Length)] := v0;
				var v77: map<bool, C1> := map[true := v65];
				v77 := v77;
			} else {
				var v78: map<bool, int> := map[if (v39[safeIndex(808, v39.Length)]) then v39[safeIndex(808, v39.Length)] else v39[safeIndex(808, v39.Length)] := f23];
				var v79: map<int, int> := map[safeDivisionInt(f23, f23) := f23];
				globalState.f7, r0, globalState.f8, v78 := v0, if (f23 in v79) then v79[f23] else f23, |(v78 + v78[v0 := 0x1b5])|, v78 + v78;
				var v80: map<bool, string> := map[v39[safeIndex(808, v39.Length)] := p0];
				var v81 := DC32(p0);
				var v82: array<map<bool, string>> := new map<bool, string>[11] [v80, v80, v80, map[!fm1(|(map[v39[safeIndex(808, v39.Length)] := v0])[v39[safeIndex(808, v39.Length)] := v0]|, v0, globalState) := "lqs"], v80, v80, fm27(globalState), v80, map[v0 := v81.cf53], v80 + v80, v80];
				v82[safeIndex(933, v82.Length)] := v80;
				var v83: map<seq<bool>, map<bool, string>> := map[v38 := v80 + v80[v0 := p0]];
				v82[safeIndex(933, v82.Length)] := if ((v38 + v38) in v83) then v83[v38 + v38] else fm27(globalState);
				var v84 := DC12(v38);
				var v85: array<D6> := new D6[6](i8 => DC17([p0, p0, p0]));
				var v86: seq<string> := [fm2(globalState)];
				var v87 := DC17(v86);
				v85[safeIndex(292, v85.Length)] := v87;
				v84, v85[safeIndex(292, v85.Length)], globalState.f8 := v84, v87, f23;
				globalState.f14 := f23;
				var v88: map<int, bool> := map[f23 := v39[safeIndex(808, v39.Length)]];
				v79 := v79[|v88| := 0x3a5];
			}
			
		}
		var v89 := DC38(v16);
		var v90 := DC40(v89);
		match (v90.(cf64 := DC40(v89)).(cf64 := v89)) {
			case DC39(cf63) =>
				var v91: array<bool> := new bool[22];
				var v92: set<array<bool>> := {v91};
				var v93: seq<array<bool>> := [v91];
				var v94: seq<set<array<bool>>> := [v92, {v91}];
				var v95: array<set<array<bool>>> := new set<array<bool>>[19] [if (v0) then v92 else v92, v92, v92, {v91, v93[safeIndex(cf63, |v93|)]}, v92, v92, v92 + v92, v92, v92, v92, v92, v94[safeIndex(cf63, |v94|)], v92, v92, v92, v92, v92, v92, v92];
				v95[safeIndex(446, v95.Length)] := v92;
				var v96: array<char> := new char[23];
				var v97: seq<D9> := [DC24(v96)];
				var v98: seq<int> := [|v97|];
				var v99 := DC27(v0, v98, v91, false, true);
				v95[safeIndex(446, v95.Length)] := {v99.cf46, v91};
				f27[safeIndex(197, f27.Length)] := f23;
				var v100: C1 := new C1(!v0, v98, f23);
				var v101: map<bool, C1> := map[v0 := v100];
				var v102: seq<bool> := [v0];
				f27[safeIndex(197, f27.Length)] := fm0(|v101|, v102[safeIndex(f23, |v102|)], globalState);
				var v103: map<int, bool> := map[fm0(f27[safeIndex(197, f27.Length)], v0, globalState) := v0];
				var v104: multiset<int> := multiset{f23, |v103|, f23};
				if ((|v104| + f27[safeIndex(197, f27.Length)]) < f27[safeIndex(197, f27.Length)]) {
					v102 := v102[safeIndex(|(seq(abs(0x36e), i9  => (|v102|)))|, |v102|) := v0] + ([v0] + v102);
					f27[safeIndex(197, f27.Length)] := |(p0 + ("chtuw" + p0))[safeIndex(safeModuloInt(f27[safeIndex(197, f27.Length)], f27[safeIndex(197, f27.Length)]), |(p0 + ("chtuw" + p0))|) := p0[safeIndex(f27[safeIndex(197, f27.Length)], |p0|)]]|;
					var v105 := DC11(if (cf63 in v104) then v104[cf63] else f27[safeIndex(197, f27.Length)]);
					v105 := v105;
					var v106: map<int, int> := map[f27[safeIndex(197, f27.Length)] := f23];
					v106 := (map v107 : int | (311 <= v107) && (v107 < 520) :: (v107 + f23) := (cf63)) + v106;
					var v108 := new C0();
				} else {
					var v109 := 'n';
					var v110: multiset<char> := multiset{v109};
					var v112 := DC25(v0, set v111 : char | v111 in v110 :: (v111));
					var v113 := DC28(v112);
					v113 := v113;
					var v114: map<bool, int> := map[v0 := f27[safeIndex(197, f27.Length)]];
					var v115 := DC26(f23);
					globalState.f6, globalState.f9, globalState.f9, globalState.f9, globalState.f8 := p0, (f27[safeIndex(197, f27.Length)] + |v114|) >= safeDivisionInt(f23, cf63), v0, v0, safeModuloInt(-0x199, v115.cf43);
					v0 := fm10("vrkt", v0, map[v0 := f27[safeIndex(197, f27.Length)]], v0, globalState) ==> v0;
					f27[safeIndex(197, f27.Length)] := f27[safeIndex(197, f27.Length)];
					var v116: multiset<bool> := multiset{v0};
					var v117: seq<multiset<bool>> := [v116, v116];
					var v118 := new C3(v117 + v117, cf63, fm1(cf63, v0, globalState), v98);
				}
				
				var v119 := DC26(f23);
				var v120 := DC28(v119);
				var v121: seq<D9> := [if (v0) then v120 else v120];
				var v122 := DC41(v121);
				v121 := v121 + v122.cf65;
			case DC38(cf62) =>
				var v123: array<bool> := new bool[21](i10 => v0);
				var v124, v125, v126 := m0(v123, seq(abs(779), i11  => ('k')), |"lwigmupbn"|, globalState);
				if (v0) {
					var v127 := 'v';
					globalState.f11 := v127;
					var v128: seq<string> := [p0, "fq"];
					var v129 := DC17(v128);
					v128 := v128 + (v129.cf30 + v128);
					var v130: map<array<bool>, bool> := map[if (v0) then v123 else v123 := if (v0) then v0 else v0];
					v130 := v130[v123 := p0 != p0];
					var v131: T1 := new C1(v0, seq(abs(0x303), i12  => (-0xcd)), |"vfnpyft"|);
					v131 := v131;
					f27[safeIndex(913, f27.Length)] := f23;
					f27[safeIndex(913, f27.Length)] := f23;
				} else {
					var v132: map<int, char> := map[v126 := fm6(0x3c2, globalState)];
					var v133 := 'j';
					var v134: map<char, bool> := map[v133 := !v0];
					var v135: seq<int> := [f23];
					var v136: seq<int> := [|v134|, f23, v135[safeIndex(v126, |v135|)]];
					globalState.f11, r1 := if (f23 in v132) then v132[f23] else v133, v135 + (seq(abs(336), i13  => (f23)));
					var v137: seq<map<int, int>> := [map[f23 := 396]];
					var v138: map<seq<map<int, int>>, string> := map[v137 := seq(abs(-831), i14  => (v133))];
					v138 := v138[v137 + v137 := p0 + (seq(abs(0x10f), i15  => (v133)))];
					v123[safeIndex(403, v123.Length)] := !(-v126 <= 495);
					v123[safeIndex(403, v123.Length)] := !v0;
					var v139: multiset<char> := multiset{'w', v133, v133};
					globalState.f6, globalState.f7, globalState.f8 := p0, fm1(v135[safeIndex(if (v133 in v139) then v139[v133] else |"ew"|, |v135|)], fm1(|(set v140 : char | v140 in v134 :: (v140))|, true, globalState), globalState), v126;
					var v141: seq<bool> := [v126 == -100];
					v123[safeIndex(403, v123.Length)] := v141[safeIndex(|p0|, |v141|)];
				}
				
				if (v0) {
					var v142: seq<int> := [|"ufjwpiq"|];
					globalState.f11 := fm6(-v142[safeIndex(f23, |v142|)], globalState);
					var v143 := new C1(v0 <== fm1(f23, !v0, globalState), v142, f23);
					globalState.f8 := v126;
					var v144: multiset<bool> := multiset{!v0, v0 <== v0, v0, v0};
					v144 := v144;
					globalState.f7 := !v0;
				} else {
					var v145: seq<int> := [v126];
					globalState.f8 := v145[safeIndex(v126, |v145|)] - f23;
					v123[safeIndex(193, v123.Length)] := v0;
					var v146: seq<bool> := [v0, v0];
					var v147: map<seq<bool>, int> := map[v146 := 0x10b];
					v123[safeIndex(193, v123.Length)] := (if (true) then fm28(v126, globalState) else v147) != v147;
					var v148 := 'e';
					var v149 := DC15(f23, p0[safeIndex(f23, |p0|)], f23);
					var v150: set<char> := {v148};
					var v151: map<D2, char> := map[DC7(v150) := v148];
					var v152 := DC7(v150);
					var v153 := DC8(v126, |v125|, v148, -f23, if (v126 in v125) then v125[v126] else v123[safeIndex(193, v123.Length)]);
					var v154: array<char> := new char[19] [v148, v148, v148, fm6(0x213, globalState), v148, v149.cf27, v148, 'y', v148, v148, v148, v148, fm6(|v146|, globalState), if (v152 in v151) then v151[v152] else v148, v148, 'u', v148, v153.cf17, 'y'];
					v154 := v154;
					var v155: map<bool, string> := map[v123[safeIndex(193, v123.Length)] <==> v0 := p0];
					v155 := v155[v123[safeIndex(193, v123.Length)] := p0];
					var v157: C3 := new C3(fm29(v126, set v156 : int | (899 <= v156) && (v156 < 229) :: (v156 - f23), globalState), f23, v123[safeIndex(193, v123.Length)], v145 + v145);
					v157 := v157;
				}
				
				v0 := !false ==> v0;
			case DC40(cf64) =>
				var v158: array<D1> := new D1[25];
				var v159: multiset<array<D1>> := multiset{v158};
				var v160: seq<D3> := [DC9(v159)];
				var v161: multiset<int> := multiset{0x189};
				var v162: map<multiset<int>, bool> := map[multiset{f23} := fm1(f23, v0, globalState)];
				var v163: seq<bool> := [!v0];
				var v164: set<int> := {-f23};
				var v165 := 'p';
				var v166: map<char, int> := map[v165 := f23];
				var v167: multiset<bool> := multiset{v0, false};
				var v168: map<bool, bool> := map[v0 := v0];
				var v169: map<bool, int> := map[v0 := |v168|];
				var v170: array<bool> := new bool[27] [v0, v160 < v160, v161[f23 := abs(f23)] >= v161, v0, if (multiset([|v163|]) in v162) then v162[multiset([|v163|])] else v0, v0, v164 > v164, -0x378 != -f23, true, -(if (v165 in v166) then v166[v165] else |v161|) < f23, f23 > f23, v0, !(v0 <==> v0), f23 <= f23, v0, !v0, v0, v0, v0, !v0, |v167| != f23, fm10(p0, v0, v169, v0, globalState), fm10(p0, v0, v169, v0, globalState), v0, v0, v0, v0];
				v170 := v170;
				var v171: array<int> := new int[6](i16 => i16 * f23);
				var v172 := DC18(map[v0 := f23]);
				r0, v171, v172, r0 := fm0(-f23, true, globalState), v171, DC18(map[v0 := f23]).(cf31 := v169), f23;
				var v173: array<array<int>> := new array<int>[16];
				v173 := v173;
				v170[safeIndex(618, v170.Length)] := if (!v0) then true else v0;
				v170[safeIndex(618, v170.Length)] := v0;
		}
		
		var v174 := DC22(v0);
		globalState.f7 := (v174.(cf38 := v0)).cf38;
		r0 := f23;
		r1 := [|[-f23, f23]|, f23 * f23, f23];
	}
}

class C5 extends T0 {
	const f25 : int
	const f26 : int
	constructor (f25 : int, f26 : int, f23 : int) {
		this.f25 := f25;
		this.f26 := f26;
		this.f23 := f23;
	}
	
	function fm8(p0: D0, p1: set<bool>, p2: bool, p3: int, globalState: GlobalState): map<map<int, bool>, bool> {
		map[(map v0 : int | v0 in (map v1 : int | v1 in multiset{f26, f26} :: (v1 - |(map v2 : char | v2 in DC7(set v3 : char | v3 in (seq(abs(0x223), i0  => ('u'))) :: (v3)).cf14 :: (v2) := (true))|) := (false)) :: (safeModuloInt(v0, f25)) := (false)) + map[|"vdal"| := false] := true && true]
	}
	method m1(p0: array<bool>, p1: bool, globalState: GlobalState) returns (r0: int, r1: set<int>) {
		var v0: map<bool, int> := map[p1 := -0x13d];
		globalState.f21, r0 := p1, |fm9(p1, -f25, map[p1 := f26] + v0, globalState)|;
		var v1: seq<int> := [f23];
		var v2 := new C1(fm1(f26, p1, globalState) || p1, v1 + (seq(abs(0xc), i0  => (-0x332))), -f26);
		var v3 := 'm';
		var v4: map<char, array<bool>> := map[v3 := p0];
		var v5 := DC8(340, -216, v3, f26, p1);
		v4 := v4[v5.cf17 := p0];
		var v6 := DC4(fm16(p1, f23, fm0(f26, false, globalState), globalState), p1, p1);
		var i1 := 0;
		while (v6.cf7)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			globalState.f21 := p1;
			var v7: array<int> := new int[4];
			var v8: map<int, int> := map[f23 := f23];
			v7[safeIndex(507, v7.Length)] := safeDivisionInt(-f23, |v8|);
			v7[safeIndex(507, v7.Length)] := -(if (f26 > f25) then f23 else f25);
			p0[safeIndex(281, p0.Length)] := p1;
			p0[safeIndex(281, p0.Length)] := f26 <= ---fm0(f25, p1, globalState);
			var v9 := "jyx";
			var v10: map<int, bool> := map[0x2fa := p0[safeIndex(281, p0.Length)]];
			var v11: map<string, map<int, bool>> := map[v9 := v10];
			v11, v7[safeIndex(507, v7.Length)] := (v11[("gxiuegtec")[safeIndex(-0x77, |"gxiuegtec"|) := v3] := v10])[fm2(globalState) := v10], -|v9|;
		}
		v0 := v0[true := f26];
		for i2 := |v1| to f25 {
			var v12: seq<bool> := [f23 > --0x348];
			var v13 := DC19(v1);
			var v14 := "nxqooq";
			globalState.f7, globalState.f21, v12, globalState.f21 := !(0x4 != (-f23 * f25)), safeModuloInt(f26, f23) != safeDivisionInt(f26, f26), (fm30(f23, v13, v14, f25, globalState) + v12) + (v12 + v12), f23 > f26;
			var v15 := DC10(p1);
			var v16: map<D3, int> := map[DC10(p1) := fm0(f25, p1, globalState)];
			if (v15 in v16) {
				var v17: map<bool, seq<int>> := map[false := v1[safeIndex(i2, |v1|) := f26]];
				var v18 := DC36(false, p1, f25, f25, v17);
				var v19: map<int, char> := map[|map[p1 := v14]| := v3];
				var v20: seq<map<int, char>> := [map[f26 := v3], v19];
				var v21: set<int> := {f26};
				globalState.f8, globalState.f7, globalState.f6, globalState.f8, v18 := f23, p1, v14, safeModuloInt(|v20[safeIndex(f26, |v20|)]|, |(v21 + v21)|), fm31(f25, globalState);
				var v22: map<int, bool> := map[0x158 := p1];
				var v23: map<bool, bool> := map[p1 := p1];
				var v24 := DC1(f26, p1, |v23|);
				v22 := v22[f25 := v24.cf2];
				var v25 := new C0();
				p0[safeIndex(200, p0.Length)] := p1;
				p0[safeIndex(200, p0.Length)] := p1;
				var v26: array<array<bool>> := new array<bool>[9];
				var v27: array<bool> := new bool[15];
				v26[safeIndex(402, v26.Length)] := v27;
				v26[safeIndex(402, v26.Length)] := v27;
			} else {
				var v28: multiset<int> := multiset{|v14|};
				globalState.f7 := p1 <==> (multiset{f26} >= v28);
				var v29: array<int> := new int[27](i3 => i3 * i2);
				var v30: map<bool, array<int>> := map[p1 := v29];
				var v31: map<int, map<bool, array<int>>> := map[safeModuloInt(f26, i2) := v30];
				v31 := v31[f25 := v30];
				var v32: map<int, int> := map[-fm0(f25, fm1(f23, p1, globalState), globalState) := f26];
				var v33: map<seq<int>, bool> := map[v1 := p1];
				var v34 := new C2(v32, v33 == v33, p1, [f25, f26] + v1, f25);
				var v35: array<D12> := new D12[18](i4 => DC37(DC35()));
				var v36 := DC35();
				var v37 := DC37(v36);
				v35[safeIndex(391, v35.Length)] := v37;
				v35[safeIndex(391, v35.Length)] := v37;
				v29[safeIndex(528, v29.Length)] := f25;
				v29[safeIndex(528, v29.Length)] := f23 - f26;
			}
			
			r0 := -f26 * f25;
			var v38: array<int> := new int[7](i5 => i5 - f25);
			v38 := v38;
		}
		r0 := f23;
		var v39: set<int> := {f23};
		r1 := (v39 * v39) * v39;
	}
	method m2(p0: string, globalState: GlobalState) returns (r0: int, r1: seq<int>) {
		var v0 := false;
		var v1: set<bool> := {v0, v0};
		var v2: map<int, int> := map[f25 := f25];
		var v3 := DC39(if (f23 in v2) then v2[f23] else f23);
		var v4: set<int> := {|v1|};
		var v5: array<int> := new int[12] [f23, |p0|, f23, f23, 0x2bc, f25, v3.cf63, -f25, f26, f25, |v4|, 0x3b3];
		var v6: map<set<bool>, array<int>> := map[v1 := v5];
		var v7 := DC38(v1);
		var v8: map<bool, set<bool>> := map[fm1(f23, !v0, globalState) := v7.cf62];
		var v9 := DC0(p0);
		v6, v1, globalState.f6 := map[if (false in v8) then v8[false] else v1 := v5], v1 - v1, v9.cf0 + p0;
		v1 := v1;
		var v10 := 'k';
		globalState.f11 := v10;
		r0 := safeModuloInt(f26 - f25, f26);
		var v11: array<char> := new char[3] [v10, v10, v10];
		var v12 := DC24(v11);
		var v13: multiset<D9> := multiset{v12};
		var v14: seq<D9> := [DC24(v11), v12, v12];
		globalState.f8 := |((v13 - v13) - multiset(v14))|;
		var v15: map<char, bool> := map['i' := v0];
		v15 := v15[v10 := v0];
		r0 := safeModuloInt(-f25, |p0|);
		r1 := [f25];
	}
	method m3(globalState: GlobalState) returns (r0: bool, r1: seq<map<bool, int>>, r2: map<int, bool>) {
		var v0: map<int, int> := map[f26 := f26];
		var v1 := false;
		var v2 := DC5(0x3ba, v1, f23, f26);
		var v3: seq<int> := [f25];
		var v4: map<bool, bool> := map[v1 := v1];
		var v5: T1 := new C2(v0, v1, v2.cf10, v3, |v4|);
		var v6: seq<T1> := [v5];
		var v7 := "uckqsq";
		var v8: seq<string> := [v7];
		for i0 := |(v6 + v6[safeIndex(-858, |v6|) := v5])[safeIndex(v5.f23, |(v6 + v6[safeIndex(-858, |v6|) := v5])|) := v5]| to if (false) then |v8[safeIndex(f23, |v8|)]| else f26 {
			globalState.f9 := f25 != f26;
			var v9: array<int> := new int[10];
			v9 := new int[28](i1 => safeDivisionInt(i1, f23));
			r0 := if (true) then v5.f28 else v5.f28;
			v9 := v9;
		}
		if (v1) {
			var v10: map<string, bool> := map[v7 := v5.f28];
			v10 := v10[v7 := v1];
			v5.f29 := v3;
			globalState.f7 := false;
			var v11: multiset<bool> := multiset{v5.f28, v5.f28};
			var v12: seq<multiset<bool>> := [multiset{v5.f28}];
			var v13 := new C3([v11, v11] + v12, f25, v5.f28, v3);
			var v14 := 'a';
			var v15: seq<bool> := [!v1];
			var v16: array<bool> := new bool[27] [v5.f28, v5.f28, v5.f28, true, v5.f28, v5.f28, v5.f28, !v1, true, v5.f28, v5.f28, v1, v1, v15[safeIndex(f26, |v15|)], v1, v1, fm1(-0xf6, v1, globalState), v1, v5.f28, v5.f28, false, v5.f28, v1, v15[safeIndex(0x164, |v15|)], true, v5.f28, v1];
			var v17 := DC27(v5.f28, seq(abs(19), i2  => (f23)), v16, v5.f28, v15[safeIndex(f23, |v15|)]);
			var v18 := DC28(v17);
			var v19: map<char, D9> := map[v14 := v18];
			v19 := v19[v14 := v18];
		} else {
			var v20 := new C0();
			globalState.f14 := -(v5.f23 + v5.f23);
			var v21 := DC32(v7);
			match (v21) {
				case DC31(cf52) =>
					var v22: set<int> := {v5.f23, 0x14f, |v4|, f23, v3[safeIndex(cf52, |v3|)]};
					var v23 := new C3(fm29(fm0(0x17f, v5.f28, globalState), v22, globalState), v5.f23, v5.f28, [f23] + v5.f29);
					v4 := v4[fm1(|map["mnssog" := 387]|, v5.f28, globalState) := v1];
					var v24: set<bool> := {v5.f28, v1};
					v24 := v24;
					var v25: map<D9, int> := map[fm24(v5.f28, |v7|, globalState) := v5.f23];
					var v26 := DC26(v5.f23);
					v25 := v25[v26 := v5.f23];
				case DC32(cf53) =>
					globalState.f14, globalState.f6 := v5.f23, seq(abs(0x2b7), i3  => ('f'));
					var v27: seq<bool> := [v5.f28];
					v1 := v5.f28 !in v27;
					var v28: set<string> := {seq(abs(915), i4  => ('h'))};
					globalState.f7 := v28 != {cf53, v7};
					v3 := v5.f29;
				case DC30(cf51) =>
					globalState.f8 := safeModuloInt(f26, -v5.f23) * |v7|;
					var v29: set<int> := {v5.f23};
					var v30: map<bool, int> := map[v1 := |map[v20 := v5.f28]|];
					var v31: map<bool, map<bool, int>> := map[false := v30];
					globalState.f21 := (-fm0(|v29|, !v5.f28, globalState) != f23) in v31;
					globalState.f7 := v5.f28;
					var v32: map<int, bool> := map[|"r"| := false];
					var v33: C2 := new C2(v0, !v5.f28, v1, v5.f29, |v32[v5.f23 := v5.f28]|);
					var v34: multiset<bool> := multiset{!true};
					var v35 := DC26(v5.f23);
					var v36: set<bool> := {false};
					var v37: array<D9> := new D9[17] [DC26(v5.f23), DC26(if (v5.f28 in v34) then v34[v5.f28] else 0x1a2), v35, v35, fm24(v5.f28, 0x9f, globalState), DC26(f25), v35, v35, v35, v35, v35, v35, DC26(|v36|), v35, v35, v35, v35];
					var v38: map<C2, array<D9>> := map[v33 := v37];
					v38 := v38[v33 := v37];
				case DC33(cf54) =>
					var v39 := DC25(v1, {'v'});
					var v40 := DC7(v39.cf42);
					var v41: map<D2, string> := map[v40 := v7];
					var v42: set<C0> := {v20};
					var v43: set<bool> := {v1, v5.f28, v5.f28};
					var v44: seq<bool> := [v20.fm20(v41, v5.f23, f26, |v42|, globalState) ==> v5.f28, v5.f28, if (v5.f28) then v5.f28 else v5.f28, |v43| >= v5.f23];
					globalState.f21 := v44[safeIndex(if (v5.f28) then f25 else 414, |v44|)];
					var v45: array<seq<seq<int>>> := new seq<seq<int>>[20];
					var v46: seq<seq<int>> := [[v5.f23, -v5.f23]];
					v45[safeIndex(694, v45.Length)] := v46;
					v45[safeIndex(694, v45.Length)] := (seq(abs(0x1fa), i5  => (v5.f29)))[safeIndex(if (v1) then v5.f23 else |v8|, |(seq(abs(0x1fa), i5  => (v5.f29)))|) := v5.f29];
					var v47: array<char> := new char[26];
					var v48 := DC24(v47);
					var v49 := DC28(v48);
					v49 := v49;
					var v50: array<set<map<char, bool>>> := new set<map<char, bool>>[11];
					var v51 := 'l';
					var v52: map<char, bool> := map[v51 := v5.f28];
					var v53: set<map<char, bool>> := {v52};
					var v54: map<map<char, bool>, int> := map[v52 := v5.f23];
					v50[safeIndex(119, v50.Length)] := v53 + (set v55 : map<char, bool> | v55 in v54 :: (v55));
					globalState.f14, v50[safeIndex(119, v50.Length)], globalState.f6 := v5.f23, v53 - v53, v7[safeIndex(|v4|, |v7|) := v51];
			}
			
			var v56: seq<bool> := [v1, v5.f28, v5.f28];
			var v57 := new C1(!!v56[safeIndex(v5.f23, |v56|)], v5.f29, f23 + f23);
			var v58 := 'y';
			var v59: array<int> := new int[17](i6 => i6 * v5.f23);
			var v60: seq<array<int>> := [v59, v59];
			var v61 := DC3(v60);
			var v62: map<char, D1> := map[v58 := v61];
			var v63: set<bool> := {v5.f28, !v5.f28, v5.f28};
			var v64: map<set<bool>, map<char, D1>> := map[v63 := v62 + v62];
			v62 := if ((v63 - v63) in v64) then v64[v63 - v63] else v62;
		}
		
		globalState.f7 := v5.f28;
		for i7 := f23 to v5.f23 - v5.f23 {
			var v65: array<int> := new int[18](i8 => safeDivisionInt(i8, f23));
			var v66: C4 := new C4(v65, |v7|);
			var v67: multiset<C4> := multiset{v66};
			var v68 := DC32(v7);
			var v69: multiset<seq<char>> := multiset{(v68.(cf53 := v7)).cf53, v7};
			if (fm1(if (v66 in v67) then v67[v66] else f23, v69 <= v69, globalState)) {
				var v70: array<bool> := new bool[21](i9 => v5.f28 <==> !v5.f28);
				v70[safeIndex(470, v70.Length)] := v5.f28;
				v70[safeIndex(470, v70.Length)] := v5.f28;
				var v71: array<map<int, int>> := new map<int, int>[24];
				var v72: map<bool, array<map<int, int>>> := map[v5.f28 := v71];
				v72 := v72[true <== !true := v71];
				var v73 := 'b';
				globalState.f11 := v73;
				var v74: array<string> := new string[7] [v7, v7 + v7, ("vrrj")[safeIndex(fm0(f26, v5.f28, globalState), |"vrrj"|) := v73], v7, fm2(globalState), seq(abs(0x17b), i10  => (v73)), "etxm" + v7];
				var v75: map<bool, int> := map[v5.f28 := i7];
				v74[safeIndex(926, v74.Length)] := (v7 + v7)[safeIndex(if (v1 in v75) then v75[v1] else |v75|, |(v7 + v7)|) := v73];
				v74[safeIndex(926, v74.Length)] := v7;
				var v76: multiset<char> := multiset{v73};
				var v77 := DC20(v5.f28, i7);
				var v78: C1 := new C1(v5.f28, v3, v5.f23);
				var v79: array<C1> := new C1[18] [v78, v78, v78, v78, v78, v78, v78, v78, v78, v78, v78, v78, v78, v78, v78, v78, v78, v78];
				var v80: map<bool, array<C1>> := map[v5.f28 := v79];
				var v81: map<bool, seq<int>> := map[!v70[safeIndex(470, v70.Length)] := v5.f29];
				var v82: multiset<int> := multiset{|v81|};
				globalState.f14 := fm0(|(multiset(v5.f29))[|v76| := abs(v77.cf34)]| - -|fm2(globalState)|, !(multiset{|v80|, v5.f23} < v82), globalState);
			} else {
				v66.f27[safeIndex(610, v66.f27.Length)] := |("hoqqdqjyb" + (seq(abs(-545), i11  => ('t'))))|;
				v66.f27[safeIndex(610, v66.f27.Length)] := v5.f23 + |v7|;
				var v83 := new C0();
				var v84: seq<bool> := [false];
				var v85 := 'l';
				var v86: set<char> := {v85};
				var v87 := DC25(true, v86);
				var v88: map<bool, int> := map[v5.f28 := f26];
				var v89 := DC21(v66.fm10(seq(abs(224), i12  => (v85)), v5.f28, v88, v84[safeIndex(|"ixco"|, |v84|)], globalState), v5.f28, v5.f23);
				var v90: multiset<char> := multiset{'q'};
				var v92 := DC7(set v91 : char | v91 in v90 :: (v91));
				var v93: map<D2, string> := map[v92 := v7];
				var v94: array<bool> := new bool[21](i13 => v1);
				var v95: set<seq<int>> := {v3, v5.f29, v3};
				var v96: array<bool> := new bool[27] [v1, v1, v84[safeIndex(|v7|, |v84|)], v1, false, v87.cf41, if (v5.f28) then v1 else false, v5.f28, v5.f28, v5.f28, v5.f28, v5.f28, v5.f28, false, v7 != v7, true, v5.f28, v89.cf36, v1, v5.f28, !(v83.fm20(v93, |v7|, v5.f23, v5.f23, globalState) !in map[v5.f28 := v94]), v5.f28, v5.f28, v5.f28, true, v5.f28, v95 >= v95];
				var v97: seq<array<bool>> := [v96, v96, v96];
				v96 := v97[safeIndex(v5.f23, |v97|)];
				globalState.f8 := fm0(|(seq(abs(0x221), i14  => (f26)))|, v1, globalState);
				var v98: multiset<seq<bool>> := multiset{v84};
				globalState.f8 := (if (v84 in v98) then v98[v84] else f23) + v66.f27[safeIndex(610, v66.f27.Length)];
			}
			
			r0 := v1;
			var v99 := 'g';
			globalState.f21 := v99 in (seq(abs(0x3a3), i15  => (DC15(f26, v99, |v7|).cf27)));
			match (DC22(v1)) {
				case DC20(cf33, cf34) =>
					globalState.f21 := fm1(f26, fm1(v5.f23, !cf33, globalState), globalState);
					v66.f27[safeIndex(182, v66.f27.Length)] := 0x3b9;
					var v100: set<bool> := {v1};
					v99, v66.f27[safeIndex(182, v66.f27.Length)], globalState.f11 := if (cf33) then v7[safeIndex(i7, |v7|)] else 'b', fm0(safeDivisionInt(|v100|, |v5.f29|), cf33, globalState), fm6(f25, globalState);
					var v101: array<bool> := new bool[5](i16 => v5.f28);
					var v102: map<array<bool>, bool> := map[v101 := v66.f27[safeIndex(182, v66.f27.Length)] <= |multiset{v1, v5.f28}|];
					var v103: map<bool, int> := map[v5.f28 := v66.f27[safeIndex(182, v66.f27.Length)]];
					v102 := v102[v101 := v66.fm10(v7, v5.f28, v103, v1, globalState)];
					globalState.f11, globalState.f8 := 'v', -f26;
				case DC21(cf35, cf36, cf37) =>
					v66.f27[safeIndex(129, v66.f27.Length)] := cf37;
					v66.f27[safeIndex(129, v66.f27.Length)] := -i7;
					var v104: array<bool> := new bool[6](i17 => v5.f28);
					v104[safeIndex(284, v104.Length)] := v5.f28;
					v104[safeIndex(284, v104.Length)] := true;
					v104[safeIndex(284, v104.Length)] := !(v99 !in v7);
					var v105: map<array<bool>, bool> := map[v104 := cf35];
					v105 := v105;
				case DC22(cf38) =>
					var v106: array<bool> := new bool[8](i18 => v5.f28);
					v106[safeIndex(436, v106.Length)] := v1;
					v106[safeIndex(436, v106.Length)] := cf38;
					globalState.f11 := 'l';
					var v107 := new C0();
					v4 := v4[!true := i7 in v5.f29];
				case DC19(cf32) =>
					var v108: map<string, bool> := map["vflk" := v5.f28];
					var v109: array<bool> := new bool[21] [v1, v5.f28, v5.f28, v5.f28, v5.f28, v5.f28, v1, v5.f28, v5.f28, v5.f28, DC20(v5.f28, f23).cf33, false, v1, true, v1, v5.f28, if (v7 in v108) then v108[v7] else v5.f28, fm1(i7, v1, globalState), v1, v5.f28, !v5.f28];
					var v110: map<array<bool>, char> := map[v109 := v99];
					globalState.f11 := if (v109 in v110) then v110[v109] else v99;
					var v111, v112, v113 := m0(v109, v7 + v7, i7, globalState);
					globalState.f14 := f25 - i7;
					var v114: array<char> := new char[5](i19 => v99);
					var v115: seq<bool> := [v5.f28];
					var v116 := DC12(v115);
					var v117: map<D9, D3> := map[DC24(v114) := v116];
					v117 := v117;
				case DC23(cf39) =>
					var v118: map<bool, int> := map[v5.f28 := v5.f23];
					var v119: seq<map<bool, int>> := [v118];
					v4 := v4[v66.fm10(v7, v1, (map[v66.fm10(seq(abs(-438), i20  => (v99)), v5.f28, v118, true, globalState) := 922])[v5.f28 := fm0(|v119[safeIndex(f25, |v119|)]|, v5.f28, globalState)], v5.f28, globalState) := v1];
					var v120: array<array<bool>> := new array<bool>[20];
					var v121: array<bool> := new bool[25];
					v120[safeIndex(181, v120.Length)] := v121;
					var v122: multiset<char> := multiset{v99};
					var v123: map<int, bool> := map[f23 := v5.f28];
					globalState.f14, v1, v120[safeIndex(181, v120.Length)], globalState.f14, r2 := safeModuloInt(safeDivisionInt(|v8|, f25), 0x1eb), if (!v5.f28 in v4) then v4[!v5.f28] else v99 in v7[safeIndex(v5.f23, |v7|) := v99], v121, if (v99 in v122) then v122[v99] else 0x285, (v123[-0x2f1 := v5.f28] + (map[i7 := v5.f28])[f25 := v1])[safeDivisionInt(|(map v124 : int | (-391 <= v124) && (v124 < 0x20c) :: (safeModuloInt(v124, 909)) := (-v5.f23))|, v5.f23) := true];
					v123 := v123[i7 := v5.f28];
					var v125: seq<bool> := [v5.f28];
					var v126: map<seq<bool>, bool> := map[v125 := f23 < f23];
					globalState.f6, r0, v126, globalState.f7, globalState.f8 := v7, v5.f28, v126 + map[v125 := v5.f28], !v5.f28, (-332 + f23) + v5.f23;
			}
			
		}
		for i21 := -v5.f23 to safeDivisionInt(v5.f23, f26) {
			globalState.f14 := fm0(f25, if (v5.f28 in v4) then v4[v5.f28] else v5.f28, globalState);
			var v127 := DC35();
			var v128: array<bool> := new bool[5];
			var v129: map<D12, array<bool>> := map[v127 := v128];
			var v130, v131 := v5.m1(if (v127 in v129) then v129[v127] else v128, v5.f28 && v5.f28, globalState);
			var v132: array<int> := new int[15](i22 => safeDivisionInt(i22, v130));
			var v133: seq<array<int>> := [v132];
			var v134 := DC3(v133);
			match (v134) {
				case DC4(cf6, cf7, cf8) =>
					var v135 := new C3(seq(abs(0x43), i23  => (multiset{true, true})), v5.f23, v5.f28, v5.f29);
					globalState.f14 := f23;
					var v136 := new C4(v132, i21);
					var v137 := DC12([v1]);
					var v138: map<D3, bool> := map[v137 := v5.f28];
					var v139: seq<bool> := [!v5.f28, !cf8];
					globalState.f21 := if (v137.(cf23 := v139) in v138) then v138[v137.(cf23 := v139)] else true;
				case DC5(cf9, cf10, cf11, cf12) =>
					globalState.f14 := i21 - v5.f23;
					var v140: seq<bool> := [v1, v5.f28];
					globalState.f21 := if (v140 <= v140) then cf10 else f26 <= i21;
					v8 := v8;
					var v141 := new C0();
				case DC6(cf13) =>
					globalState.f8 := -f23;
					globalState.f9 := v1;
					cf13 := if (v5.f28) then cf13 else cf13;
					var v142: map<seq<int>, int> := map[v5.f29 := i21];
					v142 := map[[fm0(0x1c7, v1, globalState), f25] + [v5.f23] := 0x3e1];
				case DC3(cf5) =>
					var v143: map<bool, int> := map[v5.f28 := v130];
					var v144: seq<bool> := [v5.f28];
					v143 := v143[v1 := |v144|];
					var v145: multiset<bool> := multiset{v5.f28};
					globalState.f14 := safeDivisionInt(|v145|, -f23 + |multiset{v5.f23}|);
					var v146: array<string> := new string[13](i24 => v7);
					v146[safeIndex(404, v146.Length)] := v7;
					v146[safeIndex(404, v146.Length)] := (seq(abs(-84), i25  => ('x'))) + v7;
					v7 := (v7 + v146[safeIndex(404, v146.Length)]) + v7;
			}
			
			if (v1) {
				var v147: array<set<int>> := new set<int>[18](i26 => v131);
				v147[safeIndex(867, v147.Length)] := set v149 : int | v149 in (set v148 : int | (0x344 <= v148) && (v148 < 653) :: (v148 * 0x19)) :: (safeModuloInt(v149, |multiset{'b', 'e'}|));
				v147[safeIndex(867, v147.Length)] := v131;
				var v150, v151 := v5.m2(v7 + v7, globalState);
				var v152 := DC1(-v130, v1, f26);
				globalState.f14, v1, globalState.f21 := v5.f23, v5.f28, v152.cf2;
				var v153: map<int, bool> := map[v5.f23 := v5.f28];
				var v154: map<bool, int> := map[v1 := v130];
				v128[safeIndex(568, v128.Length)] := !(|v153| > -(if (v5.f28 in v154) then v154[v5.f28] else v5.f23));
				var v155 := 'u';
				var v156 := DC32("mxdqyef");
				var v157: array<string> := new string[26] [v7, v7 + v7, v7, v8[safeIndex(fm0(v5.f23, v5.f28, globalState), |v8|)], v7[safeIndex(|v8[safeIndex(f23, |v8|)]|, |v7|) := v155], v7, v7 + v7[safeIndex(v150, |v7|) := v155], v7, seq(abs(0x20d), i27  => ('n')), "tfynwpxj" + v7, v7 + v7, seq(abs(62), i28  => (v155)), seq(abs(-0x342), i29  => (v155)), v156.cf53, "u", fm2(globalState), v7, v7 + v7, v7, "g", v7, v7, v7, v7, "esqnlyvg", "tkq" + v7];
				v157[safeIndex(881, v157.Length)] := if (v5.f28) then "nxukkyffw" else v7;
				var v158: array<map<int, int>> := new map<int, int>[1] [v0];
				var v159 := DC30(v158);
				var v160: map<char, bool> := map[if (true) then v155 else 'e' := v5.f28];
				var v161: map<string, string> := map[v7 := v7];
				v128[safeIndex(568, v128.Length)], v157[safeIndex(881, v157.Length)], v159, v160 := !(v5.f23 < v5.f23), ((if ("culpysq" in v161) then v161["culpysq"] else v7) + v7) + (v7 + v7), v159, v160 + v160;
				v7 := v157[safeIndex(881, v157.Length)];
			} else {
				v8 := seq(abs(0xd9), i30  => ("tp"));
				globalState.f8 := f23 + v5.f23;
				var v162: map<bool, int> := map[v5.f28 := f26];
				var v163 := new C1(v5.f28, [|multiset{v162}|], v5.f23);
				v132[safeIndex(872, v132.Length)] := f25;
				var v164 := 'i';
				var v165: multiset<char> := multiset{v164, v164};
				v132[safeIndex(872, v132.Length)] := (f26 + (if (v164 in v165) then v165[v164] else f26)) + f26;
				var v166 := DC31(i21);
				var v167: seq<D11> := [v166, v166, v166];
				var v168: map<seq<D11>, seq<int>> := map[v167 := v5.f29];
				var v170: set<seq<D11>> := {fm32(globalState)};
				var v172: multiset<seq<D11>> := multiset{[v166], v167, ([v166])[safeIndex(v5.f23, |[v166]|) := v166], v167};
				var v173: array<map<seq<D11>, seq<int>>> := new map<seq<D11>, seq<int>>[26] [v168[v167 := v5.f29], v168, map[v167 := seq(abs(0x1ac), i31  => (|[multiset{f26}, multiset(v5.f29), multiset{v130, 0x17f}, multiset{i21, v132[safeIndex(872, v132.Length)]}]|))], v168 + DC43(v168).cf67, map v169 : seq<D11> | v169 in v170 :: (v169) := (v5.f29), v168, v168, map[v167 := v5.f29], v168, v168, v168, v168, map v171 : seq<D11> | v171 in v172 :: (v171) := (v3), v168, if (v5.f28) then v168 else v168, fm33(|{|v5.f29|, i21, -f23}|, v5.f28, globalState), v168, v168, v168, v168, v168, v168 + v168, map[v167 := v3], v168, v168 + v168[v167 := v5.f29], v168];
				v173[safeIndex(41, v173.Length)] := v168;
				var v174: seq<bool> := [v5.f28];
				v173[safeIndex(41, v173.Length)] := map[(seq(abs(0x344), i32  => (v166)))[safeIndex(|v174|, |(seq(abs(0x344), i32  => (v166)))|) := v166] := v5.f29 + [v132[safeIndex(872, v132.Length)]]];
			}
			
		}
		var v175: set<string> := {if (v5.f28) then v7 else "seef"};
		var v176: map<int, set<string>> := map[f25 := v175];
		var v177 := DC46(v175);
		v175 := if (f26 in v176) then v176[f26] else if (v5.f28) then v175 else v177.cf74;
		var v178: map<string, bool> := map[v7 := v5.f28];
		r0 := !(if (("ulwvqx" + v7) in v178) then v178["ulwvqx" + v7] else v5.f28);
		var v179: map<bool, int> := map[v1 := f23];
		var v180: seq<map<bool, int>> := [v179, v179];
		var v181: set<int> := {|v8|};
		var v182: set<set<int>> := {{-|v181|, 0x3bd, 0x34b, f26}};
		r1 := ((v180 + (seq(abs(681), i33  => (v179)))) + (v180 + [map[v5.f28 := f25], v179, v179, v179]))[safeIndex(|v182|, |((v180 + (seq(abs(681), i33  => (v179)))) + (v180 + [map[v5.f28 := f25], v179, v179, v179]))|) := v179];
		r2 := fm5(v181, globalState);
	}
	method m4(p0: bool, globalState: GlobalState) returns (r0: bool, r1: int, r2: int) {
		globalState.f7 := p0;
		var v0 := DC31(f25);
		var v1: seq<D11> := [v0, v0];
		var v2 := "fvpo";
		var v3: seq<int> := [|v2|, f25];
		var v4: map<seq<D11>, seq<int>> := map[v1 := v3];
		var v5 := DC43(v4 + v4);
		match (v5) {
			case DC44(cf68, cf69, cf70, cf71, cf72) =>
				var v6: array<array<D12>> := new array<D12>[5];
				var v7: map<char, bool> := map['u' := p0];
				var v8 := DC34(v7);
				var v9: array<D12> := new D12[21] [DC34(v7), DC34(v7), v8, v8, v8.(cf55 := v7), v8, v8, v8.(cf55 := v7).(cf55 := v7), v8, DC34(v7), v8, v8, v8, v8, DC34(v7), v8, v8, v8, v8, v8, v8];
				v6[safeIndex(643, v6.Length)] := v9;
				v6[safeIndex(643, v6.Length)] := v9;
				var v10: map<bool, bool> := map[cf72 := cf68];
				var v11: seq<string> := [v2];
				var v12: seq<seq<string>> := [v11];
				var v13: map<int, bool> := map[f25 := cf71];
				var v14: set<bool> := {cf72};
				var v15: array<int> := new int[27] [|v10|, f26, |v12[safeIndex(|multiset{!cf71}|, |v12|)]|, |v13|, f25, f23, f26, |map[|[-|v14|]| := f26]|, 353, f25, f25, -646, |"rocyxmwas"|, cf70, |v13|, 390, f25, f23, cf70, f23, f23, f26, f25, f23, f23, f23, f23];
				var v16: map<array<int>, bool> := map[v15 := cf71];
				var v17: array<bool> := new bool[4](i0 => cf71);
				v17[safeIndex(644, v17.Length)] := cf71;
				v16, v17[safeIndex(644, v17.Length)] := v16, if (cf72 in v10) then v10[cf72] else !(f25 <= f23);
				var v18: map<seq<int>, bool> := map[v3 := cf68];
				cf68, v18, r1 := cf72, map[v3 := cf72], cf70;
				globalState.f6 := fm2(globalState);
			case DC43(cf67) =>
				var v19: multiset<bool> := multiset{p0, p0};
				var v20: seq<multiset<bool>> := [v19];
				var v21 := new C3([v19] + v20, 0x3d7 + f23, p0, v3);
				var v22: set<multiset<bool>> := {v19[p0 := abs(190)], v19};
				if (v22 == (set v23 : multiset<bool> | v23 in map[v19 := p0] :: (v23))) {
					var v24: array<bool> := new bool[9] [p0, true, fm1(f26, p0, globalState), p0, p0, p0, fm1(f25, p0, globalState), p0, p0];
					var v25: seq<array<bool>> := [v24, v24, v24, v24];
					v25 := v25;
					var v26 := 'j';
					var v27: map<int, string> := map[-f26 := v2];
					var v28: map<bool, bool> := map[p0 := p0];
					var v29: seq<string> := [v2];
					var v30: array<string> := new string[23] [v2, "pfojm" + "mmwijrbtt", seq(abs(0xac), i1  => ('s')), "vabaqredp", v2, v2[safeIndex(f25, |v2|) := v26], if (|v28| in v27) then v27[|v28|] else v2, v29[safeIndex(195, |v29|)] + v2, v2, "untl", "vfldre" + v2, v2, v2, seq(abs(0x15d), i2  => (v26)), v2, v2, v2, v2, v2, v2 + (fm34(0x1f4, globalState)).cf0, "k", if (p0) then v2 else v2, v2 + v2[safeIndex(f23, |v2|) := v26]];
					v30[safeIndex(464, v30.Length)] := v2;
					v30[safeIndex(464, v30.Length)] := v2;
					globalState.f6 := v30[safeIndex(464, v30.Length)] + "qwammv";
					var v31: C0 := new C0();
					var v32: seq<C0> := [v31, v31, v31];
					var v33 := DC48(v19);
					var v34: set<bool> := {p0};
					v32 := if (fm35(p0, false, v33.cf79, globalState) !! v34) then v32 + v32 else v32;
					v30[safeIndex(464, v30.Length)] := v2;
				} else {
					r0 := fm1(f23, p0, globalState);
					var v35: map<int, int> := map[f26 := f26];
					var v36: array<int> := new int[15];
					var v37: seq<array<int>> := [v36, v36, v36];
					var v38 := new C1(p0, seq(abs(-0x2f4), i3  => (f26)), safeModuloInt(if (|v37| in v35) then v35[|v37|] else f26, f23));
					r2 := -0x172;
					globalState.f9 := p0;
					var v39: map<bool, bool> := map[p0 := p0];
					var v40: seq<bool> := [p0, p0, p0];
					var v41: seq<bool> := [p0, v40[safeIndex(|multiset{f25}|, |v40|)], p0, p0];
					v39 := v39[p0 := v40[safeIndex(f25, |v40|)]];
				}
				
				r2 := safeDivisionInt(f25, f25);
				var v42: array<bool> := new bool[10](i4 => false);
				var v43 := DC4(map[f25 := f25], p0, p0);
				v42[safeIndex(604, v42.Length)] := v43.cf8;
				v42[safeIndex(604, v42.Length)] := p0;
			case DC45(cf73) =>
				globalState.f14 := -f25 - f25;
				var v44: array<int> := new int[2] [f25, 79];
				var v45: seq<array<int>> := [v44, v44, v44];
				v45 := v45;
				var v46: map<int, int> := map[f23 := fm0(f23, false, globalState)];
				var v47 := new C2(v46, if (p0) then p0 else p0, if (true) then p0 else p0, v3 + v3, f23);
				var v48 := new C0();
		}
		
		globalState.f7 := p0 && p0;
		globalState.f9 := p0 ==> p0;
		for i5 := f23 to f23 {
			var v49: seq<bool> := [p0];
			var v50: seq<seq<bool>> := [v49, ([fm1(|multiset(v3)|, p0, globalState)])[safeIndex(fm0(f26, p0, globalState), |[fm1(|multiset(v3)|, p0, globalState)]|) := p0], v49[safeIndex(f25, |v49|) := p0]];
			var v51: map<int, int> := map[fm0(f23, p0, globalState) := f26];
			var v52 := DC19(v3);
			var v53: array<seq<bool>> := new seq<bool>[27] [v49, v50[safeIndex(f25, |v50|)], v49, v49, v49, v49 + v49, v49, v49 + v49, v49, if (p0) then [true] else v49, v49, fm30(if (|v3[safeIndex(-f23, |v3|) := 0x264]| in v51) then v51[|v3[safeIndex(-f23, |v3|) := 0x264]|] else f26, v52, v2, i5, globalState), v49, v49, fm30(f23, v52, v2, f25, globalState), v49, if (true) then v49 else v49, v49, v49[safeIndex(261, |v49|) := p0], v49, [p0, p0], v49, v49 + v49, v49, fm30(f23, v52, v2, i5, globalState) + [p0, true, !p0], v49[safeIndex(-|v2|, |v49|) := true], v49];
			v53 := v53;
			globalState.f8, globalState.f6, globalState.f9 := f26, (v2 + v2) + v2, p0;
			var v54: array<int> := new int[25];
			var v55 := DC20(true, i5);
			var v56 := DC23(v55);
			var v57: T1 := new C2(v51, p0, !p0, v3, f25);
			var v58: seq<T1> := [v57];
			var v59: seq<array<int>> := [v54, v54];
			var v60: set<int> := {f26, f26};
			var v61: map<bool, T1> := map[v57.f28 := v57];
			var v62 := DC16(v57);
			var v63: array<T1> := new T1[28] [v57, v58[safeIndex(|v59|, |v58|)], v57, v57, v57, v57, v57, v57, v57, v58[safeIndex(i5, |v58|)], v57, v58[safeIndex(|v60|, |v58|)], v58[safeIndex(f25, |v58|)], if (!v57.f28 in v61) then v61[!v57.f28] else v57, v57, v62.cf29, v57, v57, v57, v57, v57, v62.cf29, v57, v57, v57, v57, v57, v57];
			v63[safeIndex(638, v63.Length)] := v57;
			globalState.f14, v54, globalState.f9, v56, v63[safeIndex(638, v63.Length)] := f25 * v57.f23, v54, p0, v56, v57;
			globalState.f7 := f26 == (if (v57.f28) then |v49| else f25);
		}
		var i6 := 0;
		while (f26 > f23)
			decreases 100 - i6
		{
			if (i6 >= 100) {
				break;
			}
			
			i6 := i6 + 1;
			var v64: map<bool, int> := map[p0 := f26];
			var v65 := DC18(v64 + v64);
			match (v65) {
				case DC18(cf31) =>
					var v66: seq<bool> := [true];
					var v67: array<bool> := new bool[14];
					var v68: map<seq<bool>, array<bool>> := map[v66 := v67];
					var v69: seq<map<seq<bool>, array<bool>>> := [v68];
					var v70: map<map<seq<bool>, array<bool>>, int> := map[v69[safeIndex(f23, |v69|)] := -(-0x35d - f26)];
					v70 := v70[v68 := f23];
					globalState.f21 := true;
					var v71: array<map<int, bool>> := new map<int, bool>[18];
					var v72: map<int, bool> := map[0x12c := !p0];
					v71[safeIndex(978, v71.Length)] := v72[f23 := p0];
					v71[safeIndex(978, v71.Length)] := v72;
					var v73: map<int, map<int, bool>> := map[-(f25 - |(seq(abs(0x2c1), i7  => (f25)))|) := map[f23 := !p0]];
					v73 := v73[|[-0xc9, 0x256]| := map[f25 := p0]];
			}
			
			var v74: seq<D15> := [v5, v5, v5];
			v74 := v74;
			var v75: C0 := new C0();
			var v76: map<int, C0> := map[f23 := v75];
			v76 := v76[safeModuloInt(f23, f25) := v75];
			globalState.f8 := v3[safeIndex(f26, |v3|)] + (if (p0) then f23 else |v64[p0 := f26]|);
		}
		r0 := p0;
		r1 := f25;
		var v77 := DC0(v2);
		var v78 := DC2(v77);
		r2 := match v78 {
			case DC1(cf1, cf2, cf3) => f26 + f23
			case DC0(cf0) => safeModuloInt(|(map v79 : seq<int> | v79 in (set v80 : seq<int> | v80 in (seq(abs(-0x12a), i8  => (v3))) :: (v80)) :: (v79) := (f25))|, 0x16a)
			case DC2(cf4) => f26
		};
	}
}

class C6 extends T0 {
	const f24 : int
	constructor (f24 : int, f23 : int) {
		this.f24 := f24;
		this.f23 := f23;
	}
	
	method m1(p0: array<bool>, p1: bool, globalState: GlobalState) returns (r0: int, r1: set<int>) {
		var v0: map<int, array<bool>> := map[|[f24, f23]| := p0];
		for i0 := f23 to |v0| {
			r0 := f24;
			var v1: map<bool, bool> := map[p1 := !p1];
			var v2: set<map<bool, bool>> := {if (p1) then v1 else map[!p1 := p1], v1, v1, v1};
			var v3 := 'j';
			var v4: multiset<char> := multiset{'n', v3};
			v2 := {v1[fm1(|v4|, p1, globalState) := p1], fm7(f24, p1, p1, globalState) + v1, v1, map[p1 := p1] + v1};
			p0[safeIndex(897, p0.Length)] := p1;
			var v5: multiset<bool> := multiset{p1, p1};
			var v6: seq<bool> := [p1];
			p0[safeIndex(897, p0.Length)] := v5 !! fm4(DC0(seq(abs(591), i1  => (v3))), i0, v6, v3, globalState);
			var v7 := "cduaj";
			var v8: seq<string> := [v7];
			globalState.f14 := -|v8|;
		}
		var v9: seq<int> := [f23];
		globalState.f14 := v9[safeIndex(0x1b7, |v9|)];
		p0[safeIndex(487, p0.Length)] := fm0(f23, p1, globalState) != f24;
		var v10: map<int, int> := map[f23 := f24];
		var v11 := "l";
		var v12: set<int> := {f24, if (f24 in v10) then v10[f24] else |v11|};
		var v13: multiset<bool> := multiset{fm1(|v12|, false, globalState), p1};
		var v14: set<bool> := {p1, p1};
		var v15: T0 := new C3([v13[true := abs(|v14|)], multiset{p1, p1}, v13, v13, multiset{p1}], f23, p1, v9);
		var v16: seq<T0> := [v15];
		p0[safeIndex(487, p0.Length)] := v15 !in (v16[safeIndex(|map[p1 := p1]|, |v16|) := v15] + v16);
		var v17: map<bool, int> := map[p1 := 0x388];
		var v18: seq<map<bool, int>> := [map[p0[safeIndex(487, p0.Length)] := v15.f23], v17];
		r0 := if (v17 in v18) then f23 * f24 else -954;
		for i2 := v15.f23 to safeDivisionInt(f24, f23) {
			var v19: map<bool, bool> := map[false := !!p1];
			v19 := v19[p1 := fm1(v15.f23, p1, globalState)];
			var v20: seq<bool> := [fm1(f23, p0[safeIndex(487, p0.Length)], globalState), p1, v13 !! v13, p1, if (p1) then !p1 else true];
			globalState.f7 := v20[safeIndex(|fm2(globalState)|, |v20|)];
			var v21 := 'n';
			var v22: map<char, seq<int>> := map[v21 := v9 + [v15.f23]];
			v22 := v22[v11[safeIndex(v15.f23, |v11|)] := v9];
			v20 := v20 + v20;
		}
		var v23: array<bool> := new bool[15];
		forall i3 | 0 <= i3 < v23.Length {
			v23[i3] := v11 <= v11;
		}
		r0 := f24;
		var v24: multiset<int> := multiset{f23};
		var v25: map<multiset<int>, int> := map[v24 := |v13|];
		var v26 := DC50(v25);
		r1 := {v15.f23, |(v26.cf84 + v25)|, f23};
	}
	method m2(p0: string, globalState: GlobalState) returns (r0: int, r1: seq<int>) {
		var v0 := true;
		var v1: map<bool, int> := map[v0 := f23];
		var v2 := 'u';
		var v3: set<char> := {v2};
		var v4: map<set<char>, string> := map[v3 := p0];
		var v5: set<int> := {|(if (v3 in v4) then v4[v3] else (p0[safeIndex(f23, |p0|) := v2])[safeIndex(f23, |p0[safeIndex(f23, |p0|) := v2]|) := v2])|, f23};
		v1 := v1[{f23} > v5 := f24];
		v2 := v2;
		if (v0) {
			var v6 := DC10(!v0);
			match (v6) {
				case DC10(cf21) =>
					var v7: array<array<bool>> := new array<bool>[22];
					var v8: array<bool> := new bool[13](i0 => cf21);
					v7[safeIndex(540, v7.Length)] := v8;
					var v9: seq<int> := [f24];
					var v10: C3 := new C3(fm29(|v9|, {f23, f23}, globalState), f24, false, v9 + v9);
					globalState.f14, v2, v7[safeIndex(540, v7.Length)], v10, v8 := -f23, v2, v8, v10, v8;
					globalState.f6 := fm2(globalState);
					globalState.f21 := v0;
					v8[safeIndex(682, v8.Length)] := cf21;
					v8[safeIndex(682, v8.Length)] := cf21;
				case DC11(cf22) =>
					var v11: map<set<int>, bool> := map[v5 := p0 <= p0[safeIndex(f23, |p0|) := v2]];
					v11 := v11[v5 := v0];
					var v12: multiset<bool> := multiset{v0};
					var v13: array<int> := new int[8] [f24, cf22, cf22, |v1|, f24, f24, |v12|, f24];
					var v14: map<array<int>, bool> := map[v13 := v0];
					var v15: multiset<int> := multiset{|map[v0 := if (v13 in v14) then v14[v13] else v0]|, |p0|, f23, cf22, cf22};
					var v16: array<bool> := new bool[11];
					var v17: map<multiset<int>, array<bool>> := map[v15 := v16];
					var v18: map<int, map<multiset<int>, array<bool>>> := map[-f24 := v17];
					v18 := v18[f24 := v17];
					var v19 := DC7(v3);
					var v20: map<D2, bool> := map[v19 := v0];
					var v21: map<map<D2, bool>, int> := map[v20 + map[v19 := v0] := f24];
					v21 := v21[v20 := f24];
					var v22: C5 := new C5(|p0|, -|"r"|, f24);
					v22 := v22;
				case DC12(cf23) =>
					var v23: map<int, int> := map[f23 := f24];
					var v24: array<int> := new int[12] [f23, fm0(f24, v0, globalState), f24, if (false) then f24 else f24, if (f24 in v23) then v23[f24] else f23, f24, f23 - |p0|, |cf23| + -fm0(f23, v0, globalState), fm0(-|v23|, v0, globalState), f24, f24, f24];
					v24 := v24;
					var v25: seq<string> := [p0, p0];
					var v26: seq<string> := [p0[safeIndex(|v25|, |p0|) := v2]];
					var v27 := DC32(v26[safeIndex(f24, |v26|)]);
					var v28: multiset<D11> := multiset{v27, v27};
					var v29: map<multiset<D11>, seq<int>> := map[v28 := ([f23, f24, |cf23|, f24, f24])[safeIndex(f24, |[f23, f24, |cf23|, f24, f24]|) := fm0(f24, v0, globalState)]];
					var v30: seq<int> := [f23, f24, f23, f23, f24];
					v29 := v29[v28 := v30 + v30];
					var v31: map<bool, string> := map[v0 := "bxnn"];
					v31 := if (!v0) then v31 + v31 else v31;
					globalState.f8 := -f23;
				case DC9(cf20) =>
					globalState.f9 := v0;
					globalState.f8 := f23;
					globalState.f6 := p0;
					var v32: multiset<bool> := multiset{v0, true, true};
					var v33: seq<bool> := [v0, v0, v0];
					var v34: array<bool> := new bool[21] [v0, -|p0| >= f23, !v0, !v0, v0, !v0, fm1(f23, v0, globalState), v2 in p0, v0, v0 <==> !v0, true, fm1(f23, v0, globalState), v0, v32 > v32, v0, v0, v0, "thw" <= p0, v0, v33[safeIndex(f23, |v33|)], v0];
					v34 := v34;
				case DC13(cf24) =>
					var v35: multiset<bool> := multiset{v0};
					var v36: set<set<bool>> := {fm35(v0, !v0, v35, globalState)};
					var v37: seq<int> := [f24];
					var v38: seq<seq<int>> := [v37];
					var v39: map<bool, seq<int>> := map[v0 := [f23, f24]];
					var v40 := DC36(true, v0, f23, f24, v39);
					var v41: multiset<int> := multiset{f24, f24, -f24, f24};
					var v42: array<int> := new int[17] [f24, f23, f23 + f23, f24, f23, f23, f23 * f24, f24, safeDivisionInt(196, f24), f24, f23, |v38|, v40.cf59, fm0(if (f24 in v41) then v41[f24] else -104, v0, globalState), f23, f23 * |multiset{f23, f24, f23, |p0|}|, f23];
					v42[safeIndex(645, v42.Length)] := f24;
					var v43: set<bool> := {v0};
					v36, v42[safeIndex(645, v42.Length)], r0 := {v43}, |p0|, |({v0, v0} * v43)|;
					var v44: array<bool> := new bool[15](i1 => true);
					var v45, v46, v47 := m0(v44, "hqminyem", f23, globalState);
					v0 := f23 != (v42[safeIndex(645, v42.Length)] * |(seq(abs(-0x208), i2  => (v2)))|);
					v1 := if (v41 > v41) then if (v0) then v1 else v1 else v1;
			}
			
			var v48: array<bool> := new bool[2] [v0, false];
			v48 := v48;
			var v49: array<char> := new char[25];
			var v50: map<bool, char> := map[v0 := v2];
			v49 := new char[6] [v2, v2, v2, if (false in v50) then v50[false] else fm6(f23, globalState), if (v0 in v50) then v50[v0] else v2, fm6(f23, globalState)];
			globalState.f14 := 0x22e;
			globalState.f9, globalState.f7 := true, false;
		} else {
			globalState.f8 := f23;
			var v51 := DC39(-fm0(f24, v0, globalState));
			v51 := if (v0) then v51 else v51;
			var v52: multiset<bool> := multiset{true, v0};
			var v53: T1 := new C3([v52, v52, v52, v52], f23, v0 <==> v0, seq(abs(0xcd), i3  => (|p0|)));
			globalState.f11, globalState.f14, globalState.f21, v53 := v2, v53.f23, fm1(f23, v53.f28, globalState), if (v53.f28) then v53 else v53;
			globalState.f7 := v53.f28;
			var v54: map<bool, string> := map[v53.f28 := p0];
			globalState.f6 := ((seq(abs(350), i4  => (v2))) + p0) + ((seq(abs(635), i5  => ('k'))) + (if (v0 in v54) then v54[v0] else p0))[safeIndex(v53.f23, |((seq(abs(635), i5  => ('k'))) + (if (v0 in v54) then v54[v0] else p0))|) := fm6(-66, globalState)];
		}
		
		r0 := f23 * f23;
		var v55: array<int> := new int[3];
		forall i6 | 0 <= i6 < v55.Length {
			v55[i6] := safeModuloInt(i6, f24 * f24);
		}
		var v56: map<int, bool> := map[f24 := v0];
		var v57: map<map<int, bool>, int> := map[v56 := f24];
		globalState.f9 := (v2 in "r") || (f23 < |v57|);
		r0 := 0xa6 - f23;
		var v58: seq<int> := [f24, f23];
		r1 := v58;
	}
}

function fm0(p0: int, p1: bool, globalState: GlobalState): int {
	if (true) then safeModuloInt(-0x8a, |map[false := false]|) else 0x21d
}
function fm1(p0: int, p1: bool, globalState: GlobalState): bool {
	true
}
function fm2(globalState: GlobalState): string {
	("ko" + (seq(abs(0x128), i0  => ('d')))) + (seq(abs(0x4e), i1  => ('i')))
}
function fm3(p0: int, globalState: GlobalState): map<char, bool> {
	map v0 : char | v0 in ['s'] :: (v0) := (["ac"] == ["n", "dubvfe", seq(abs(0x98), i0  => ('w')), "pt", "k"])
}
function fm4(p0: D0, p1: int, p2: seq<bool>, p3: char, globalState: GlobalState): multiset<bool> {
	multiset(([false, false] + [false]) + ([true] + [true, false, false, true, false]))
}
function fm5(p0: set<int>, globalState: GlobalState): map<int, bool> {
	map[|{{!!false}, {true, false, true, true, DC25(true, {'d'}).cf41}, {false}, {true}}| := multiset{!true, false} != multiset([true])]
}
function fm6(p0: int, globalState: GlobalState): char {
	'b'
}
function fm7(p0: int, p1: bool, p2: bool, globalState: GlobalState): map<bool, bool> {
	(map[false := false] + map[false := !true]) + (map[false := true] + map[!true := !false])
}
function fm9(p0: bool, p1: int, p2: map<bool, int>, globalState: GlobalState): map<multiset<bool>, char> {
	map[multiset{false, false, false, false, true} * multiset{false, false} := 'l']
}
function fm13(p0: char, p1: bool, p2: int, p3: bool, globalState: GlobalState): map<int, map<int, int>> {
	if (false) then map[|(seq(abs(671), i0  => ('p')))| := map[0x111 := |multiset{!false, false}|]] + map[-33 := map v0 : int | (228 <= v0) && (v0 < 940) :: (safeModuloInt(v0, 160)) := (0x317)] else map[--0x309 := map[|map[false := false]| := -0x292]]
}
function fm14(p0: int, p1: char, p2: bool, globalState: GlobalState): multiset<map<bool, int>> {
	multiset{map[false := 0x1cb], map[false := -0x1e0]} + multiset([map[false := |{true}|]] + [map[true := 0xcf]])
}
function fm15(p0: bool, p1: map<int, multiset<bool>>, globalState: GlobalState): multiset<seq<bool>> {
	(if (true) then multiset{[true]} else multiset([[true]])) - multiset{[false]}
}
function fm16(p0: bool, p1: int, p2: int, globalState: GlobalState): map<int, int> {
	if (0x3dc != |map[true := false]|) then map[0x1de := 0x18a] else map[0x69 := 0x1c3] + map[0xfb := -748]
}
function fm17(p0: bool, p1: int, globalState: GlobalState): multiset<char> {
	if (true || false) then multiset{'s'} else multiset{'c', 'o', 'i'}
}
function fm19(p0: int, p1: bool, globalState: GlobalState): seq<int> {
	[|multiset{'f', 'q', 'y'}|] + ([-|{"kwrvsluv", "mtfd"}|, |multiset([|(set v0 : D0 | v0 in [DC2(DC0(seq(abs(0x279), i0  => ('x'))))] :: (v0))|])|, |{false}|] + [|"yrcmuol"|])
}
function fm21(p0: seq<string>, globalState: GlobalState): set<int> {
	{0xaf, 0x28c, -0x3d0, 896} + {|"b"|}
}
function fm22(p0: D3, globalState: GlobalState): D3 {
	DC12([!true, false] + [!true])
}
function fm23(globalState: GlobalState): D1 {
	DC5(-(0xb6 - |(seq(abs(357), i0  => ([true, !!false])))|), -0x2cb != |"lucmkk"|, -424, |((set v0 : int | (0x2a1 <= v0) && (v0 < 0xf5) :: (v0 - 937)) + {|multiset{|[|{map[false := false]}|]|, |(seq(abs(0x149), i1  => ('l')))|}|})|)
}
function fm24(p0: bool, p1: int, globalState: GlobalState): D9 {
	DC26(-|("yglvl" + "emwcyv")|)
}
function fm25(p0: int, p1: bool, globalState: GlobalState): map<bool, int> {
	map[true := --0x10d] + map[true := 773]
}
function fm26(p0: set<bool>, globalState: GlobalState): D0 {
	DC2(DC1(873, !true, |[0x17c]|))
}
function fm27(globalState: GlobalState): map<bool, string> {
	if (!(multiset{false, false} > multiset([!false, false, !true]))) then map[!false := "ivfd"] else map[true := "jbqnyc"] + map[false := "jjxglfc"]
}
function fm28(p0: int, globalState: GlobalState): map<seq<bool>, int> {
	map[[false] := 0x364] + (map[[true, false] := -|[false, false]|] + (map v0 : seq<bool> | v0 in {[true]} :: (v0) := (200)))
}
function fm29(p0: int, p1: set<int>, globalState: GlobalState): seq<multiset<bool>> {
	((seq(abs(171), i0  => (multiset{false, false, true}))) + [multiset{true}]) + [multiset{!false}, multiset([true, false]), multiset{true}, multiset{true}, multiset{true, false}]
}
function fm30(p0: int, p1: D8, p2: string, p3: int, globalState: GlobalState): seq<bool> {
	([true] + [false]) + ([!false] + [false])
}
function fm31(p0: int, globalState: GlobalState): D12 {
	if (!!true) then DC36(true, false, 125, 710, map[true := seq(abs(674), i0  => (-0x1b6))]) else if (!true) then DC36(false, true, 0x302, |[0x3dc]|, map[false := [735, |"p"|, -0xe0]]) else DC36(false, !false, 0x3d1, |[false, false]|, map[true := [|[|[false]|, -67, |map[!!true := -|[true]|]|]|]])
}
function fm32(globalState: GlobalState): seq<D11> {
	([DC31(-0x2dc)] + (seq(abs(0x2a9), i0  => (DC31(0x246))))) + [DC31(|multiset([DC1(-|{true, false}|, true, 0x14d).cf2, false])|), DC31(|{!true}|), DC31(|[true]|), DC31(|[multiset([|['l']|])]|)]
}
function fm33(p0: int, p1: bool, globalState: GlobalState): map<seq<D11>, seq<int>> {
	(map[seq(abs(-70), i0  => (DC31(807))) := [-|"aguorpcjs"|]] + map[[DC31(|"habj"|)] := [-|"nv"|, 0xc8]]) + (map[[DC31(|"kdusdcx"|)] := [582]] + (map v0 : seq<D11> | v0 in map[[DC31(|"to"|)] := |"bffhds"|] :: (v0) := ([|map[true := -0x2b3]|])))
}
function fm34(p0: int, globalState: GlobalState): D0 {
	DC0("mlwjdjdmf")
}
function fm35(p0: bool, p1: bool, p2: multiset<bool>, globalState: GlobalState): set<bool> {
	({!!true, false} + {true}) + {false, true}
}
method m0(p0: array<bool>, p1: string, p2: int, globalState: GlobalState) returns (r0: D0, r1: map<int, bool>, r2: int) {
	r2 := p2;
	var v0: array<int> := new int[20];
	var v1 := false;
	var v2: map<array<int>, bool> := map[v0 := v1];
	var v3: seq<int> := [p2];
	var v4 := DC20(true, v3[safeIndex(p2, |v3|)]);
	var v5 := new C5(|v2|, safeModuloInt(p2, p2), v4.cf34);
	var v6: seq<string> := [p1];
	var v7 := 'x';
	var v8: C4 := new C4(v0, v5.f25);
	var v9: seq<C4> := [v8, v8];
	var v10: multiset<C4> := multiset{v9[safeIndex(v5.f25, |v9|)]};
	var v11: seq<bool> := [fm1(0x37f, v1, globalState)];
	var v12: set<int> := {p2};
	var v13: map<int, int> := map[v5.f25 := v5.f26];
	globalState.f7, r2, v1, globalState.f8 := |p1[safeIndex(DC21(v1, v1, -|v6[safeIndex(p2, |v6|)]|).cf37, |p1|) := v7]| < safeModuloInt(--|v10|, |v11|), v5.f25, v12 >= v12, safeModuloInt(-|v13|, p2);
	var v14: map<bool, bool> := map[v1 := v1];
	var v15: C0 := new C0();
	var v16: map<C0, int> := map[v15 := 0x31d];
	var v17: map<map<C0, int>, map<bool, bool>> := map[v16 := v14];
	var v18: array<map<bool, bool>> := new map<bool, bool>[7] [v14, v14, map[!v1 := v1], map[!v1 := v1], if (map[v15 := v5.f26] in v17) then v17[map[v15 := v5.f26]] else v14, v14, v14];
	forall i0 | 0 <= i0 < v18.Length {
		v18[i0] := map[false := v1] + DC52(v14).cf85;
	}
	for i1 := v5.f25 to p2 {
		var v19: array<bool> := new bool[16];
		v19 := v19;
		r2 := fm0(v5.f25, v1, globalState);
		globalState.f7 := v1;
		globalState.f6 := p1;
	}
	v1 := v1;
	var v20: C2 := new C2(v13, v1, v1, v3, p2);
	var v21: array<C2> := new C2[1] [v20];
	var v22: multiset<array<C2>> := multiset{v21, v21, v21, v21};
	var v23 := DC55(v22);
	var v24 := DC1(fm0(p2, false, globalState), v23.cf89 >= v22, v5.f26);
	r0 := v24;
	var v25: map<bool, seq<int>> := map[v1 := v3];
	var v26 := DC36(!v20.f32, v1, |p1|, p2, v25);
	r1 := match v26 {
		case DC35() => (map[-|v20.f31| := true])[v5.f26 := v20.f32]
		case DC36(cf56, cf57, cf58, cf59, cf60) => (map v27 : int | v27 in [|(seq(abs(0x34a), i2  => (v7)))|, |[-950]|] :: (v27 * v5.f26) := (v20.f32)) + (map v28 : int | (-214 <= v28) && (v28 < 460) :: (safeDivisionInt(v28, cf58)) := (cf56))[|map[true := |p1|]| := cf57]
		case DC34(cf55) => map v29 : int | (0x3dd <= v29) && (v29 < 320) :: (safeModuloInt(v29, v5.f25)) := (v1)
		case DC37(cf61) => map[v5.f25 := v1]
	};
	r2 := v5.f25;
}
method Main() {
	var v0 := "hg";
	var v1 := DC0(v0);
	var v2 := 't';
	var v3: array<string> := new string[12] [v0, "mexrhnq", v0, v0, v0, v1.cf0, v0, "qu", "wggsbwf", "rnjar", v0, v0[safeIndex(-282, |v0|) := v2]];
	var v4 := true;
	var v5: multiset<bool> := multiset{v4, true, v4};
	var v6: multiset<multiset<bool>> := multiset{v5};
	var v7: array<set<int>> := new set<int>[5];
	var v8: array<int> := new int[16];
	var v9 := -0x14;
	var v10: map<array<int>, int> := map[v8 := v9];
	var globalState := new GlobalState(true, 0x38b, v3, -486, false, v6, v0, false, -0x3d1, false, 0x278, 'f', 0xb3, v7, 435, 0x32e, true, 558, v10 + v10, false, v8, false, 0x27a);
	var i0 := 0;
	while (v4)
		decreases 100 - i0
	{
		if (i0 >= 100) {
			break;
		}
		
		i0 := i0 + 1;
		var v11: array<bool> := new bool[16];
		var v12, v13, v14 := m0(v11, v0, v9, globalState);
		var v15: map<int, int> := map[v9 := 921];
		v8[safeIndex(91, v8.Length)] := if (|v0| in v15) then v15[|v0|] else 0x3cd;
		v8[safeIndex(91, v8.Length)] := fm0(v9, v4, globalState);
		var v16: array<array<string>> := new array<string>[22] [v3, v3, v3, v3, v3, if (v4) then v3 else v3, if (v4) then v3 else v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3, v3];
		v16[safeIndex(99, v16.Length)] := v3;
		v16[safeIndex(99, v16.Length)] := v3;
		var v17: seq<bool> := [v4, |v0| > v14, v4];
		v17 := v17;
	}
	var v18: array<map<int, int>> := new map<int, int>[20];
	var v19: map<int, int> := map[378 := |multiset{v9}|];
	v18[safeIndex(957, v18.Length)] := v19;
	var v21: seq<int> := [0x1d3];
	v18[safeIndex(957, v18.Length)] := v19[0x2b0 := v9] + ((map v20 : int | v20 in v21 :: (v20 - 0x2dc) := (v9)) + v19);
	var v22: array<bool> := new bool[20];
	var v23: set<bool> := {v4, v4};
	v22[safeIndex(547, v22.Length)] := {v4, v4, v4, v4} < v23;
	var v24: map<char, bool> := map[v2 := fm1(v9, v4, globalState)];
	var v25: multiset<string> := multiset{"jxuh", fm2(globalState), v0};
	globalState.f7, globalState.f14, globalState.f14, v22[safeIndex(547, v22.Length)] := v4 ==> (map[v2 := !!v4] != v24), v9 * (if (v0 in v25) then v25[v0] else v9), safeModuloInt(v9, v9), !true;
	var v26: seq<bool> := [!v4];
	for i1 := v9 to |(if (true) then v26[safeIndex(v9, |v26|) := v22[safeIndex(547, v22.Length)]] else v26)| {
		globalState.f14 := v9;
		var v27: array<seq<bool>> := new seq<bool>[9];
		var v28: map<int, array<seq<bool>>> := map[i1 := v27];
		v28 := v28[safeModuloInt(i1, v9) := v27];
		v4 := v22[safeIndex(547, v22.Length)];
		v8[safeIndex(240, v8.Length)] := v9;
		v8[safeIndex(240, v8.Length)] := safeDivisionInt(|{v4}|, |v21| + v9);
	}
	var i2 := 0;
	while ((v9 * v9) != safeDivisionInt(v9, -v9))
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		globalState.f8 := |[v23 + v23, v23, v23, v23]|;
		var v29, v30, v31 := m0(v22, v0, v9, globalState);
		var v32: seq<map<char, bool>> := [fm3(v9, globalState), v24, map[v2 := v4], v24, map[v2 := v4]];
		var v33, v34, v35 := m0(v22, seq(abs(0x2ef), i3  => (v2)), |map[|v32| := v31]| * |fm4(DC0("k"), 998, v26, v2, globalState)|, globalState);
		v8[safeIndex(279, v8.Length)] := -0x18;
		var v36: map<bool, array<int>> := map[v22[safeIndex(547, v22.Length)] := v8];
		v8[safeIndex(279, v8.Length)] := fm0(v31, v9 < |v36|, globalState);
	}
	globalState.f6 := "ahdth" + "syrkxpe";
	globalState.f14 := -safeModuloInt(if (v4) then v9 else v9, -|fm2(globalState)|);
	if (v22[safeIndex(547, v22.Length)]) {
		globalState.f7 := v4;
		globalState.f8 := safeDivisionInt(|v0|, v9 + 30);
		var v37: set<int> := {v9};
		var v38: seq<map<int, bool>> := [fm5(v37, globalState)];
		globalState.f8 := safeModuloInt(v9, safeDivisionInt(-|v0|, |v38|));
		var v39: map<bool, int> := map[v4 := v9];
		var v40 := DC1(v9, v22[safeIndex(547, v22.Length)], v9);
		v39 := (map[true := |v0[safeIndex(v9, |v0|) := v2]|])[v40.cf2 := -355] + v39;
		v21 := v21;
	} else {
		var v41: multiset<int> := multiset{v9, |v21|};
		globalState.f14, globalState.f6, v22[safeIndex(547, v22.Length)], v4, globalState.f21 := v9, (v0[safeIndex(v9, |v0|) := v2])[safeIndex(v9, |v0[safeIndex(v9, |v0|) := v2]|) := v2], v4, v41 >= v41, v26[safeIndex(|v26| * |(fm2(globalState))[safeIndex(|v21|, |fm2(globalState)|) := v2]|, |v26|)];
		globalState.f11 := fm6(v9, globalState);
		globalState.f9 := DC1(v9, v22[safeIndex(547, v22.Length)], v9).cf2;
		globalState.f7, v22[safeIndex(547, v22.Length)], v22[safeIndex(547, v22.Length)] := if (v22[safeIndex(547, v22.Length)]) then v22[safeIndex(547, v22.Length)] else v4, v4, v22[safeIndex(547, v22.Length)];
		var v42, v43, v44 := m0(v22, seq(abs(-195), i4  => (v2)), -v9, globalState);
	}
	
	for i5 := v9 to v9 {
		v22[safeIndex(547, v22.Length)] := v0 != v0;
		globalState.f21 := v4;
		globalState.f8 := i5 + fm0(0x37c, v4, globalState);
		var v45, v46, v47 := m0(if (v22[safeIndex(547, v22.Length)]) then v22 else v22, v0, i5, globalState);
	}
	var i6 := 0;
	while (v9 >= -886)
		decreases 100 - i6
	{
		if (i6 >= 100) {
			break;
		}
		
		i6 := i6 + 1;
		var v48: map<int, bool> := map[v9 := v4];
		v8, globalState.f14, globalState.f21, globalState.f7, globalState.f14 := v8, |v48| - -v9, fm0(v9, v4, globalState) == 0x2e7, if (v4) then v4 else v4, v9;
		globalState.f9 := v22[safeIndex(547, v22.Length)];
		var v49: seq<seq<bool>> := [[v22[safeIndex(547, v22.Length)]], v26];
		globalState.f14 := safeModuloInt(v9, v9 - |v49[safeIndex(v9, |v49|)]|);
		var v50: array<char> := new char[13];
		v50[safeIndex(963, v50.Length)] := v2;
		var v51: multiset<int> := multiset{v9};
		globalState.f6, globalState.f11, v50[safeIndex(963, v50.Length)] := "ghjmshpn" + v0, v0[safeIndex(|v51|, |v0|)], v2;
	}
	var v52, v53, v54 := m0(v22, fm2(globalState), v9, globalState);
	var v55: set<int> := {|v0|, v9};
	for i7 := |(v5 - v5)| to |v55| {
		globalState.f8, v54, v8 := --|fm2(globalState)|, safeDivisionInt(-737, i7), v8;
		var v56: array<map<bool, bool>> := new map<bool, bool>[28];
		v56 := v56;
		var v57: array<seq<array<int>>> := new seq<array<int>>[21];
		var v58: seq<array<int>> := [v8];
		var v59 := DC3(v58);
		v57[safeIndex(996, v57.Length)] := v59.cf5;
		v57[safeIndex(996, v57.Length)] := v58;
		var v60 := DC5(v54, true, |v55|, fm0(v9, v4, globalState));
		v18[safeIndex(957, v18.Length)] := v18[safeIndex(957, v18.Length)][fm0(fm0(v54, v4, globalState), v22[safeIndex(547, v22.Length)], globalState) + --|[v9, i7, v60.cf12, i7]| := v9];
	}
	var v61: multiset<int> := multiset{v9, |([0x4b])[safeIndex(v9, |[0x4b]|) := v54]|};
	if (v4 <==> (v61[|{v1, DC0(v0)}| := abs(|v0|)] !! v61)) {
		globalState.f9 := v4;
		v8[safeIndex(268, v8.Length)] := -v9;
		v8[safeIndex(268, v8.Length)] := -0x26;
		var v62: C0 := new C0();
		var v63: seq<multiset<bool>> := [v5, multiset{v22[safeIndex(547, v22.Length)]}];
		var v64: C3 := new C3(v63, v8[safeIndex(268, v8.Length)], !v4, [v9]);
		var v65: multiset<map<C0, C3>> := multiset{map[v62 := v64]};
		var v66: map<C0, C3> := map[v62 := v64];
		var v67 := new C3(fm29(v9, {v54, v9}, globalState), v8[safeIndex(268, v8.Length)], v8[safeIndex(268, v8.Length)] !in v21, v21[safeIndex(if (v66 in v65) then v65[v66] else v54, |v21|) := v9]);
		v22, v8[safeIndex(268, v8.Length)], globalState.f7 := v22, v54, fm1(v9, if (false) then !true else v22[safeIndex(547, v22.Length)], globalState);
		v22[safeIndex(547, v22.Length)] := v0 < v0;
	} else {
		var v68: map<bool, bool> := map[v4 := !v22[safeIndex(547, v22.Length)]];
		var v69: map<bool, int> := map[v22[safeIndex(547, v22.Length)] := |v68|];
		v9 := if (v22[safeIndex(547, v22.Length)] in v5) then v5[v22[safeIndex(547, v22.Length)]] else if (v22[safeIndex(547, v22.Length)] in v69) then v69[v22[safeIndex(547, v22.Length)]] else v9;
		var v70: map<string, bool> := map[seq(abs(0x15b), i8  => (v2)) := v22[safeIndex(547, v22.Length)]];
		v70 := v70[v0 := v4];
		var v71: seq<seq<bool>> := [v26, v26];
		v26 := v26 + (v71[safeIndex(-|v0|, |v71|)] + v26);
		var v72: seq<map<bool, bool>> := [v68];
		v72 := v72;
		v8[safeIndex(201, v8.Length)] := v54;
		globalState.f7, v8[safeIndex(201, v8.Length)], v9, v9 := v4, v9, v9, v9;
	}
	
	var v73: map<bool, bool> := map[!v4 := v22[safeIndex(547, v22.Length)]];
	v73 := v73[v4 := v4];
	var v74: array<C6> := new C6[4];
	var v75: map<bool, array<C6>> := map[v22[safeIndex(547, v22.Length)] := v74];
	var v76: T1 := new C2(v18[safeIndex(957, v18.Length)], v22[safeIndex(547, v22.Length)], v4, fm19(v9, v22[safeIndex(547, v22.Length)], globalState), |v75|);
	var v77: array<multiset<char>> := new multiset<char>[11];
	var v78: map<T1, array<multiset<char>>> := map[v76 := v77];
	v78 := v78[v76 := v77];
	globalState.f7 := safeModuloInt(v54, |v0[safeIndex(v9, |v0|) := v2]|) <= v9;
	print v0, "\n";
	print v1.cf0, "\n";
	print v2, "\n";
	print v3[0], "\n";
	print v3[1], "\n";
	print v3[2], "\n";
	print v3[3], "\n";
	print v3[4], "\n";
	print v3[5], "\n";
	print v3[6], "\n";
	print v3[7], "\n";
	print v3[8], "\n";
	print v3[9], "\n";
	print v3[10], "\n";
	print v3[11], "\n";
	print v4, "\n";
	print v5 == multiset{true, true, true}, "\n";
	print v6 == multiset{multiset{true, true, true}}, "\n";
	print v8[0], "\n";
	print v8[7], "\n";
	print v8[9], "\n";
	print v8[11], "\n";
	print v9, "\n";
	print |v10|, "\n";
	print globalState.f0, "\n";
	print globalState.f1, "\n";
	print globalState.f2[0], "\n";
	print globalState.f2[1], "\n";
	print globalState.f2[2], "\n";
	print globalState.f2[3], "\n";
	print globalState.f2[4], "\n";
	print globalState.f2[5], "\n";
	print globalState.f2[6], "\n";
	print globalState.f2[7], "\n";
	print globalState.f2[8], "\n";
	print globalState.f2[9], "\n";
	print globalState.f2[10], "\n";
	print globalState.f2[11], "\n";
	print globalState.f3, "\n";
	print globalState.f4, "\n";
	print globalState.f5 == multiset{multiset{true, true, true}}, "\n";
	print globalState.f6, "\n";
	print globalState.f7, "\n";
	print globalState.f8, "\n";
	print globalState.f9, "\n";
	print globalState.f10, "\n";
	print globalState.f11, "\n";
	print globalState.f12, "\n";
	print globalState.f14, "\n";
	print globalState.f15, "\n";
	print globalState.f16, "\n";
	print globalState.f17, "\n";
	print |globalState.f18|, "\n";
	print globalState.f19, "\n";
	print globalState.f20[0], "\n";
	print globalState.f20[7], "\n";
	print globalState.f20[9], "\n";
	print globalState.f20[11], "\n";
	print globalState.f21, "\n";
	print globalState.f22, "\n";
	print i0, "\n";
	print v18[17] == map[378 := 1, 688 := -20, -265 := -20, 4 := -20], "\n";
	print v19 == map[378 := 1], "\n";
	print v21 == [467], "\n";
	print v22[7], "\n";
	print v23 == {true}, "\n";
	print v24 == map['t' := true], "\n";
	print v25 == multiset{"jxuh", "koddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii", "hg"}, "\n";
	print v26 == [false, false, false], "\n";
	print i2, "\n";
	print i6, "\n";
	print v52.cf1, "\n";
	print v52.cf2, "\n";
	print v52.cf3, "\n";
	print v53 == map[0 := false, -1 := false, 1 := true], "\n";
	print v54, "\n";
	print v55 == {2, -20}, "\n";
	print v61 == multiset{-20, 1}, "\n";
	print v73 == map[false := false, true := true], "\n";
	print |v75|, "\n";
	print v76.f28, "\n";
	print v76.f29 == [3, -2, 1, 1, 7], "\n";
	print v76.f23, "\n";
	print |v78|, "\n";
}