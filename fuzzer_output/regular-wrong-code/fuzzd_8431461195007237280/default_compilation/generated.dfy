datatype D0 = DC1 | DC0(cf0: int) | DC2(cf1: D0)
datatype D1 = DC4(cf3: int, cf4: int, cf5: bool) | DC5(cf6: int) | DC6(cf7: bool, cf8: bool) | DC3(cf2: bool)
datatype D2 = DC7(cf9: seq<bool>)
datatype D3 = DC9(cf11: bool, cf12: int) | DC10(cf13: bool) | DC8(cf10: char)
datatype D4 = DC11(cf14: T0)
datatype D5 = DC13 | DC14(cf16: bool, cf17: int) | DC12(cf15: multiset<bool>)
datatype D6 = DC16(cf19: bool, cf20: array<array<bool>>) | DC15(cf18: map<char, int>) | DC17(cf21: D6)
datatype D7 = DC19(cf23: bool) | DC20(cf24: bool, cf25: bool) | DC18(cf22: set<bool>)
datatype D8 = DC21(cf26: seq<int>)
datatype D9 = DC23(cf28: seq<set<int>>, cf29: int, cf30: bool, cf31: int) | DC24(cf32: bool) | DC22(cf27: string)
datatype D10 = DC26(cf34: int, cf35: array<bool>, cf36: bool, cf37: bool, cf38: int) | DC25(cf33: array<int>)
datatype D11 = DC28(cf40: C2, cf41: int, cf42: bool) | DC27(cf39: map<string, int>)
datatype D12 = DC30(cf44: bool, cf45: bool) | DC31(cf46: bool, cf47: string, cf48: int, cf49: bool) | DC29(cf43: set<int>) | DC32(cf50: D12)
datatype D13 = DC33(cf51: multiset<string>)
class GlobalState {
	const f0 : int
	var f1 : seq<int>
	const f2 : seq<int>
	const f3 : map<bool, bool>
	const f4 : int
	const f5 : map<char, int>
	var f6 : int
	const f7 : array<map<int, int>>
	const f8 : map<array<int>, array<bool>>
	const f9 : string
	const f10 : map<array<bool>, bool>
	const f11 : seq<array<int>>
	const f12 : bool
	var f13 : seq<bool>
	const f14 : bool
	const f15 : int
	const f16 : seq<map<int, bool>>
	const f17 : bool
	constructor (f0 : int, f1 : seq<int>, f2 : seq<int>, f3 : map<bool, bool>, f4 : int, f5 : map<char, int>, f6 : int, f7 : array<map<int, int>>, f8 : map<array<int>, array<bool>>, f9 : string, f10 : map<array<bool>, bool>, f11 : seq<array<int>>, f12 : bool, f13 : seq<bool>, f14 : bool, f15 : int, f16 : seq<map<int, bool>>, f17 : bool) {
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
	}
	
}

function fm0(globalState: GlobalState): int {
	--((0x4 % |map[|[true]| := false]|) + 0x250)
}
function fm1(p0: bool, p1: char, p2: int, globalState: GlobalState): bool {
	!!({'k'} > (if (!true) then {'t', 's', 'f'} else {'r', 'k'}))
}
function fm2(globalState: GlobalState): set<bool> {
	if (!('x' !in "anivmhual")) then {false} else {true, !true, true}
}
function fm5(p0: int, globalState: GlobalState): set<int> {
	if ('w' in "ujvkkcrk") then {|(seq(0x10, i0  => ('u')))|, |map[!false := |[-|"rpe"|]|]|} else {0xe9}
}
function fm7(p0: bool, p1: int, globalState: GlobalState): map<char, int> {
	((map v0 : char | v0 in (set v1 : char | v1 in map['p' := |map[|multiset{|{0x18}|}| := true]|] :: (v1)) :: (v0) := (|map[false := 'v']|)) + map['s' := |(seq(0x1fc, i0  => ('g')))|]) + (map v2 : char | v2 in (map v3 : char | v3 in ['e', 'g', 'b', 'o'] :: (v3) := (|[|[multiset([true]), multiset([true]), multiset{false}, multiset{true}]|, |[0x275, 0x342]|]|)) :: (v2) := (|(map v4 : int | (309 <= v4) && (v4 < 772) :: (v4 / -942) := (-|"wg"|))|))
}
function fm8(globalState: GlobalState): string {
	"kh"
}
function fm9(globalState: GlobalState): multiset<D2> {
	multiset{DC7([false])} + (multiset{DC7([true])} + multiset(seq(0x2d2, i0  => (DC7([false])))))
}
function fm10(globalState: GlobalState): char {
	'a'
}
function fm11(globalState: GlobalState): map<string, bool> {
	(map[seq(0x23, i0  => ('s')) := false] + map["nsov" := false]) + (map["k" := !true] + (map v0 : string | v0 in {"b", "v"} :: (v0) := (true)))
}
function fm14(p0: int, globalState: GlobalState): map<char, int> {
	map['t' := |{-|"pess"|, |"pj"|, |multiset([true, true])|, |map[501 := true]|}| * 0x8f]
}
function fm15(p0: char, p1: int, p2: D3, globalState: GlobalState): multiset<string> {
	if (DC6(!false, false).cf8) then DC33(multiset{"u"}).cf51 + multiset(["ddchqe", seq(0x314, i0  => ('j'))]) else DC33(multiset{seq(0x255, i1  => ('h'))}).cf51
}
function fm16(globalState: GlobalState): seq<map<int, int>> {
	[map[|[true, false, false, true, true]| := 698]]
}
function fm17(globalState: GlobalState): set<int> {
	{0x1d8 - -750}
}
function fm18(p0: int, globalState: GlobalState): map<string, int> {
	map[DC31(true, "wyhk", 0x2cc, true).cf47 := DC14(!false, 0x230).cf17] + map["qgxmchpc" := 0x19b]
}
function fm19(p0: bool, p1: bool, globalState: GlobalState): map<int, int> {
	map v0 : int | (-0x17a <= v0) && (v0 < -0xf1) :: (v0 * -0x212) := (|(seq(0xf9, i0  => (0x2b2)))| - |[false, true]|)
}
function fm20(p0: bool, globalState: GlobalState): char {
	match DC10(!false) {
		case DC9(cf11, cf12) => 'u'
		case DC10(cf13) => 'e'
		case DC8(cf10) => cf10
	}
}
function fm21(globalState: GlobalState): string {
	("ihvn" + "nvbkgws") + "qjxpor"
}
function fm22(globalState: GlobalState): seq<map<bool, int>> {
	([map[true := 0x11a], map[!true := 0x10a], map[true := |map[false := -0x6b]|], map[true := |map[false := false]|]] + [map[false := 571]]) + ([map[true := --0x25c]] + [map[true := 453]])
}
function fm23(p0: int, p1: bool, globalState: GlobalState): map<bool, bool> {
	(map[false := !true] + map[false := false]) + (map[true := true] + map[!!true := true])
}
function fm24(globalState: GlobalState): D3 {
	DC10(!true)
}
function fm25(p0: map<bool, bool>, p1: D9, p2: bool, p3: int, globalState: GlobalState): D0 {
	DC1()
}
function fm26(globalState: GlobalState): map<bool, int> {
	(map[!true := -304] + map[false := 597]) + map[false := -0x80]
}
function fm27(globalState: GlobalState): seq<bool> {
	(if (true) then [false, true] else [false, true, true]) + [false]
}
function fm28(p0: bool, p1: int, p2: bool, globalState: GlobalState): multiset<int> {
	(multiset([-633, |(set v0 : int | (-0x36c <= v0) && (v0 < 398) :: (v0 % 0x17a))|, -|(seq(0x50, i0  => ('h')))|]) - multiset{0x34}) + (if (false) then multiset{|map[|(seq(466, i1  => (map[true := true])))| := {0xa8}]|, |[|map[|map[{|{'a'}|} := false]| := 0x239]|]|, |map[true := true]|} else multiset([|{0x3d0}|]))
}
function fm29(globalState: GlobalState): D1 {
	DC4(0x127, |multiset{!true}| % 0x1a4, 0x97 <= -0x209)
}
function fm30(p0: bool, globalState: GlobalState): seq<int> {
	[-734]
}
function fm31(p0: set<bool>, p1: int, globalState: GlobalState): D1 {
	match DC29({|[0x253]|, 0x1cf}) {
		case DC30(cf44, cf45) => DC5(993)
		case DC31(cf46, cf47, cf48, cf49) => DC5(-cf48)
		case DC29(cf43) => DC5(0xd7)
		case DC32(cf50) => if (false) then DC5(-0x108) else DC5(|map[false := |{784}|]|)
	}
}
method m0(p0: int, p1: bool, p2: int, globalState: GlobalState) returns (r0: array<bool>, r1: bool, r2: int) {
	if (!p1) {
		var v0 := 'd';
		var v1 := "srfvaskfl";
		var v2 := DC0(p2);
		var v3: map<string, int> := map[v1 := v2.cf0];
		var v4: C3 := new C3(v0, v3);
		var v5: set<C3> := {v4, v4};
		var v6: seq<set<C3>> := [v5];
		var v7 := DC6(p1, p1);
		var v8: seq<bool> := [p1];
		var v9: multiset<int> := multiset{p2, p0, -p2, -p0, p0};
		var v10: array<bool> := new bool[17] [p1, p1, p1, v5 > v6[p0], v7.cf8, p1, !(-p0 < |fm21(globalState)|), p1, !false, v8[|v9|], p1, p1, p1, p1, p0 <= p0, p1, p1];
		v10[517] := !p1;
		v10[517] := p1;
		var v11: map<bool, char> := map[v10[517] := v0];
		v11 := v11[p1 := v0];
		v10[517] := true;
		globalState.f6 := |v1|;
		v10[517] := p1 <== v10[517];
	} else {
		var v12 := 'j';
		var v13 := "vqatmsp";
		var v14: array<int> := new int[6] [p0, |(seq(0x3da, i0  => ('b')))[p2 := v12]|, p0 / -|v13|, p0, -689, p2];
		v14[69] := p0;
		var v15: multiset<int> := multiset{p0};
		var v16: array<multiset<int>> := new multiset<int>[4] [v15, v15 - v15, v15, v15 * v15];
		v16[93] := fm28(p1, p2, p1, globalState);
		var v17: multiset<bool> := multiset{p1, p1, p1, !!p1};
		var v18: array<bool> := new bool[17] [p1, p1, p1, true, p1, p1, p1, p1, p1, if (p1) then p1 else p1, p1, true, fm1(p1, v12, 245, globalState), !!p1, v17 !! v17, fm1(true, v12, p2, globalState), p1];
		var v19: map<string, int> := map["fjijcedg" := p2];
		var v20: C2 := new C2(v12, v19);
		var v21: set<C2> := {v20};
		v14[69], v16[93], r0, globalState.f6, r1 := p2 / p2, v15 - v15, v18, 664, !(v20 !in v21);
		var v22: C3 := new C3(v12, map["umuxo" := v14[69]]);
		var v23: array<string> := new string[7](i1 => v13);
		v23[671] := v13;
		r1, v13, v22, v23[671] := p1, ((seq(0x26, i2  => (v12))) + v13) + (v13[v14[69] := v12])[p2 := 'q'], v22, "estdt" + "mfn";
		var v24 := DC2(DC0(p2));
		var v25 := DC27(v19);
		var v26: C1 := new C1(0x2af, p1, fm20(fm1(p1, v12, |v15|, globalState), globalState), v25.cf39 + v19[v23[671] := v14[69]]);
		var v27: map<char, int> := map[v12 := p0];
		var v28: T0 := new C0(v23[671], v27, v12, map[v13 := |map[fm0(globalState) := v14[69]]|]);
		var v29: map<T0, int> := map[v28 := p2];
		v18, v24, v26, r2, v18 := v18, v24, v26, if (v28 in v29) then v29[v28] else v14[69] - p2, v18;
		v28.f18 := v12;
		v18[297] := v26.f23;
		v18[297] := 316 in [v26.f22];
	}
	
	var v30 := "lgcj";
	for i3 := p0 to |(v30 + fm21(globalState))| {
		var v31: map<bool, int> := map[false := i3];
		globalState.f6 := if (false in v31) then v31[false] else i3;
		var v32: map<bool, string> := map[!p1 := v30 + v30];
		var v33 := 'y';
		v32 := v32[fm1(p1, v33, |(seq(0x21a, i4  => (v33)))[p2 := fm20(fm1(p1, v33, p2, globalState), globalState)]|, globalState) := v30];
		var v34: map<char, int> := map[v33 := p2];
		var v35: map<string, int> := map[v30 := p0];
		var v36 := new C0(v30[i3 := v33], v34 + v34, 'r', v35 + v35);
		var v37: array<multiset<bool>> := new multiset<bool>[3];
		var v38: map<int, array<multiset<bool>>> := map[p0 := v37];
		var v39: array<array<multiset<bool>>> := new array<multiset<bool>>[24] [v37, v37, v37, v37, v37, v37, v37, v37, v37, v37, if (p0 in v38) then v38[p0] else v37, v37, v37, v37, v37, v37, v37, v37, v37, v37, v37, v37, v37, v37];
		v39[181] := v37;
		v39[181] := v37;
	}
	var v40: array<int> := new int[2];
	v40[338] := p2 / p0;
	var v41: array<bool> := new bool[6];
	v41[360] := --p2 >= -|[p2]|;
	var v42 := 'j';
	v40[338], v41[360], v30 := 0x2cb, !p1, [v42, fm10(globalState)] + fm8(globalState);
	var v43: map<char, bool> := map[v42 := v41[360]];
	var v45: map<char, int> := map[v42 := 0x3d3];
	v43 := (map v44 : char | v44 in v45 :: (v44) := (v41[360])) + (v43 + map[v42 := v41[360]]);
	var v46: map<bool, bool> := map[!p1 := p1];
	if (-(v40[338] / |v46|) < p2) {
		var v47: map<int, bool> := map[p0 := true];
		var v48: seq<map<int, bool>> := [v47];
		if (v48 != (seq(0x1d8, i5  => (map v49 : int | (-0x3e0 <= v49) && (v49 < 0x3ba) :: (v49 / p2) := (v41[360]))))) {
			var v50 := DC23(seq(865, i6  => ({p0})), p2, p1, p2);
			globalState.f1, v41[360], v40[338], v41[360], r1 := ([v40[338]])[-p0 := v50.cf31], if (p1) then v41[360] else !p1, v40[338] % p0, v41[360], v41[360];
			var v51: map<seq<char>, int> := map[v30 := p0];
			var v52: set<bool> := {true, true, !p1};
			var v53 := DC0(p0);
			v51 := v51[v30[(fm31(v52, p0, globalState)).cf6 := v42] := v53.cf0];
			v41[360] := p1;
			var v54: map<bool, int> := map[v41[360] := p2];
			v54 := v54[v41[360] := -(-p2 / 0x77)];
			v40[338] := -0x354;
		} else {
			var v55 := DC25(v40);
			var v56: seq<bool> := [v41[360]];
			var v57: map<array<int>, seq<bool>> := map[(v55.(cf33 := v40)).cf33 := v56];
			v57 := v57;
			var v58: seq<int> := [p0, p0];
			globalState.f6, v40[338], globalState.f1, v41[360] := |(seq(0x4f, i7  => (v42)))|, (v40[338] / -p2) % p2, v58, false;
			r1 := v41[360];
			v42 := v42;
			v41[360] := !p1;
		}
		
		if (true) {
			var v59: multiset<bool> := multiset{p1};
			v41[360] := if (v59 < v59) then p1 else v41[360];
			var v60: seq<bool> := [v41[360]];
			var v61: map<string, int> := map[v30 := |v60|];
			var v62: C1 := new C1(p2, v41[360], v42, v61[v30 := p0]);
			var v63: multiset<C1> := multiset{v62, v62, v62};
			var v64: map<int, multiset<C1>> := map[v40[338] := v63];
			var v65: multiset<int> := multiset{|v64|};
			v41[360] := !(-v40[338] in v65);
			var v66: map<int, array<int>> := map[p0 := v40];
			var v67: set<int> := {p0};
			var v68: multiset<array<int>> := multiset{v40, v40, if (v41[360]) then if (|v67| in v66) then v66[|v67|] else v40 else v40};
			v68 := v68;
			v40[338] := v40[338] % (p0 * v40[338]);
			var v69 := new C0(v30, fm14(v40[338], globalState), v42, v61);
		} else {
			var v70: set<int> := {p2};
			var v71: seq<bool> := [v41[360]];
			var v72: multiset<string> := multiset{v30};
			var v73: seq<bool> := [v71[|map[|v30[|v72| := v42]| := if (p2 in v47) then v47[p2] else p1]|]];
			r1, v41[360] := (v70 * v70) !! (v70 + v70), v73[v40[338]];
			v40[338] := |v30| % p0;
			v40[338] := p2;
			v42 := v42;
			var v74 := DC3(p1);
			var v75 := DC1();
			var v76: array<seq<bool>> := new seq<bool>[25](i8 => v71);
			v76[765] := v71;
			var v77: array<array<D1>> := new array<D1>[4];
			var v78: array<D1> := new D1[17];
			v77[952] := v78;
			v74, globalState.f6, v75, v76[765], v77[952] := v74, -121, if (p1) then v75 else v75, v71, v78;
		}
		
		var v79 := DC15(v45);
		match (v79) {
			case DC16(cf19, cf20) =>
				var v80: map<int, string> := map[613 := v30];
				v80 := (v80 + v80) + (v80 + v80);
				var v82 := DC22(v30);
				var v83: seq<string> := [("qribx")[0x49 := v42], v82.cf27];
				var v84 := new C2(v42, map v81 : string | v81 in v83 :: (v81) := (|map[|{p2, p0}| := p1]|));
				var v85: set<map<int, bool>> := {v47};
				var v87: multiset<map<int, bool>> := multiset{map v86 : int | (0x2fb <= v86) && (v86 < -0x1f2) :: (v86 * |v47|) := (v41[360])};
				v85 := (v85 - (set v88 : map<int, bool> | v88 in v87 :: (v88))) - v85;
				var v89: map<string, int> := map[v30 := p2];
				var v90: C1 := new C1(p2, fm30(p1, globalState) <= [p2, v40[338]], 'u', v89);
				var v91: C3 := new C3(v42, v89 + v89);
				var v92: array<C3> := new C3[16];
				v92[439] := v91;
				var v93: multiset<bool> := multiset{v41[360], true};
				v90, v91, globalState.f6, v40[338], v92[439] := v90, v91, (|v30| * v40[338]) + (p2 - |v93|), (DC26(p2, v41, v90.f23, v90.f23, p2).(cf34 := -0xcb, cf37 := !p1, cf38 := v40[338])).cf34, v91;
			case DC15(cf18) =>
				v41[360] := v41[360];
				v30 := (v30 + v30) + v30;
				var v94: array<C3> := new C3[26];
				var v95: C3 := new C3(v42, map[v30 := p2]);
				v94[44] := v95;
				var v97: map<string, int> := map[v30 := |DC29({p2, p2}).cf43|];
				var v98: C0 := new C0(v30, map v96 : char | v96 in v30 :: (v96) := (p0), v42, v97 + v97);
				var v99: array<map<int, bool>> := new map<int, bool>[24](i9 => v47);
				v94[44], v98, v99 := v95, v98, v99;
				var v100: set<bool> := {p1};
				var v101: map<int, int> := map[p2 := p2];
				var v102: map<bool, map<int, int>> := map[v41[360] := v101];
				globalState.f6 := |(if (v100 <= v100) then v101 else (if (p1 in v102) then v102[p1] else v101) + map[p0 := p0])|;
			case DC17(cf21) =>
				var v103: map<bool, int> := map[if (true) then p1 else p1 := v40[338]];
				v103 := v103;
				v41[360] := v41[360];
				v41[360] := p2 < p2;
				globalState.f6 := p0;
		}
		
		globalState.f6, v41 := if (v41[360]) then p0 % 362 else p0 + p2, v41;
		var v104: array<array<bool>> := new array<bool>[1];
		var v105 := DC16(p1, v104);
		var v106 := DC17(v105);
		var v107 := DC17(if (p1) then v105 else v106);
		v107 := v107;
	} else {
		var v108: map<string, int> := map[v30 := v40[338]];
		var v109: C2 := new C2('f', v108);
		var v110: seq<C2> := [v109];
		var v111: set<int> := {|v110|, v40[338]};
		if (v111 !! v111) {
			var v112: map<bool, int> := map[v41[360] <==> v41[360] := p0];
			var v113: T0 := new C0(seq(525, i10  => (v42)), v45, v42, fm18(v40[338], globalState));
			var v114: seq<T0> := [v113];
			v112 := v112[v114 != v114 := v40[338] / p0];
			var v115: seq<bool> := [p1, p1];
			var v116: seq<bool> := [p1, v115[-0x1e4], p1];
			var v117: map<int, int> := map[|v115| := p2];
			v117 := v117;
			var v118: C3 := new C3('p', map["jfumgxqqe" := p0]);
			var v119 := DC24(false);
			var v120: map<C3, bool> := map[v118 := p1 <==> !v119.cf32];
			v120 := v120[v118 := p1];
			v41[360], v113.f18, globalState.f6 := v109.fm3(p1, |(v30 + v30)|, globalState), v42, (p2 / |[v41, v41]|) % v40[338];
			var v122: seq<set<int>> := [{p0}, (set v121 : int | (-0x2ad <= v121) && (v121 < 204) :: (v121 % |{v41[360]}|)) * v111];
			v122 := [v111] + [v111, v111];
		} else {
			r2 := -((-p0 - p2) / v40[338]);
			var v123: array<char> := new char[7](i11 => v30[p2]);
			v123[967] := v42;
			v123[967] := DC8(v42).cf10;
			var v124: map<set<int>, seq<bool>> := map[v111 := [p1]];
			var v125: seq<set<int>> := [v111];
			var v126 := DC23(v125, p2, p1, v40[338]);
			var v127: seq<bool> := [v41[360], v126.cf30];
			var v128: seq<seq<bool>> := [v127];
			var v129: map<seq<bool>, bool> := map[v127 := v41[360]];
			var v130: map<map<seq<bool>, bool>, seq<bool>> := map[v129 := v127];
			var v132: array<seq<bool>> := new seq<bool>[6] [(if (v111 in v124) then v124[v111] else v127) + v127, v127, v127, v128[v40[338]], v127, if ((map v131 : seq<bool> | v131 in (seq(-0xe5, i12  => (v127))) :: (v131) := (v41[360])) in v130) then v130[map v131 : seq<bool> | v131 in (seq(-0xe5, i12  => (v127))) :: (v131) := (v41[360])] else [p1, v41[360]]];
			v132 := v132;
			var v133: set<bool> := {p1};
			var v134: seq<int> := [|v133|];
			var v135: seq<map<string, int>> := [v108];
			var v136 := new C1(-v134[-p2], p1, v123[967], v135[|v127|]);
			r2 := p2;
		}
		
		r1 := (set v137 : int | v137 in v111 :: (v137 % 36)) !! v111;
		var v138: array<multiset<array<bool>>> := new multiset<array<bool>>[11];
		var v139: multiset<array<bool>> := multiset{v41};
		var v140: map<bool, multiset<array<bool>>> := map[true := v139];
		v138[166] := if (p1 in v140) then v140[p1] else v139;
		v138[166] := v139;
		var v141: seq<int> := [v40[338], fm0(globalState)];
		globalState.f1 := (if (v109.fm3(p1, p0, globalState)) then v141 else v141) + v141;
		var v142: set<bool> := {p1};
		var v143 := DC31(p1, seq(758, i13  => (v42)), |fm30(v41[360], globalState)|, p1);
		globalState.f6, r2 := if (!(v142 !! v142)) then v40[338] else |(v143.cf47 + v30)|, if (p1) then v40[338] * |map[-v40[338] := v41]| else p0;
	}
	
	var v144: map<string, int> := map[(v30[p0 := v42])[p2 := v42] := p2];
	var v145 := new C2(v42, v144);
	r0 := v41;
	r1 := if (p1) then v41[360] else p1;
	var v146: set<int> := {v40[338]};
	r2 := v145.fm4(v41[360], v42, fm10(globalState), v146 < v146, globalState);
}
trait T0 {
	var f18 : char
	const f19 : map<string, int>
}

trait T1 extends T0 {
	function fm3(p0: bool, p1: int, globalState: GlobalState): bool 
	function fm4(p0: bool, p1: char, p2: char, p3: bool, globalState: GlobalState): int 
	method m1(p0: bool, globalState: GlobalState) returns (r0: D0, r1: int, r2: bool, r3: bool) 
	method m2(p0: bool, p1: int, p2: int, p3: int, globalState: GlobalState) returns (r0: int) 
}

class C0 extends T0 {
	const f20 : string
	const f21 : map<char, int>
	constructor (f20 : string, f21 : map<char, int>, f18 : char, f19 : map<string, int>) {
		this.f20 := f20;
		this.f21 := f21;
		this.f18 := f18;
		this.f19 := f19;
	}
	
	function fm6(p0: int, p1: map<bool, string>, globalState: GlobalState): bool {
		(|map[false := false]| - |f20|) != (-0x23e + 0x39c)
	}
}

class C1 extends T1 {
	var f22 : int
	const f23 : bool
	constructor (f22 : int, f23 : bool, f18 : char, f19 : map<string, int>) {
		this.f22 := f22;
		this.f23 := f23;
		this.f18 := f18;
		this.f19 := f19;
	}
	
	function fm3(p0: bool, p1: int, globalState: GlobalState): bool {
		f23
	}
	function fm4(p0: bool, p1: char, p2: char, p3: bool, globalState: GlobalState): int {
		390 * --(if (!f23) then 0x353 else 0x16e)
	}
	function fm12(globalState: GlobalState): bool {
		f23
	}
	function fm13(p0: int, p1: int, p2: int, p3: bool, globalState: GlobalState): char {
		if (!!f23) then f18 else 'b'
	}
	method m1(p0: bool, globalState: GlobalState) returns (r0: D0, r1: int, r2: bool, r3: bool) {
		var v0: map<bool, int> := map[f23 := f22];
		var v1: set<map<bool, int>> := {v0, v0};
		r3 := v1 >= ({v0} + v1);
		globalState.f6 := f22 + f22;
		var v2: multiset<int> := multiset{f22, f22, f22};
		var i0 := 0;
		while ((if (p0) then v2 else v2) >= v2[f22 := f22])
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v3 := DC10(p0);
			match (v3) {
				case DC9(cf11, cf12) =>
					var v4: seq<bool> := [!cf11];
					var v5: multiset<bool> := multiset{f23, f23, p0};
					var v6: map<int, int> := map[cf12 := |v5|];
					var v7 := DC8(f18);
					var v8: seq<int> := [f22];
					var v9: set<bool> := {cf11};
					var v10: set<int> := {cf12};
					var v11: map<bool, seq<int>> := map[cf11 := v8];
					var v12: array<bool> := new bool[21] [p0, fm1(p0, f18, f22, globalState), cf11, false, cf11, !f23, f23, fm1(f23, f18, f22, globalState), v4 < [f23, p0, p0, cf11], cf12 <= |v6|, f22 != f22, cf11, fm1(f23, v7.cf10, v8[cf12], globalState), v9 >= fm2(globalState), !false, p0, |v10| in (if (f23 in v11) then v11[f23] else v8), cf12 == cf12, f23, cf11, fm1(false, f18, -0x117, globalState)];
					v12 := v12;
					var v13 := DC2(DC1());
					var v14 := DC2(v13);
					v14 := v14;
					var v15 := "hkdt";
					var v16: set<string> := {"ylwbg", v15};
					v16 := v16 + {v15, v15, v15};
					var v17: map<char, int> := map[f18 := f22];
					var v19: seq<string> := [v15, v15, v15];
					var v20: seq<map<string, int>> := [map[seq(0xaa, i1  => (f18)) := cf12], f19, map v18 : string | v18 in v19 :: (v18) := (f22), f19];
					var v21 := new C0(v15, v17, 'x', v20[f22]);
				case DC10(cf13) =>
					cf13 := DC9(cf13, fm0(globalState)).cf11 <==> cf13;
					var v22: array<bool> := new bool[13];
					var v23: map<array<bool>, int> := map[v22 := f22];
					var v24: multiset<bool> := multiset{f23};
					var v25: set<bool> := {false};
					var v26: array<int> := new int[8] [f22 + f22, (if (v22 in v23) then v23[v22] else |v24|) * f22, f22, f22, f22, |v25|, -f22, 0xfb - f22];
					v26[392] := f22 / |fm14(fm0(globalState), globalState)|;
					var v27: seq<int> := [fm0(globalState)];
					v26[392] := |multiset(v27)| / f22;
					var v28: array<string> := new string[6];
					v28[156] := seq(0x84, i2  => (f18));
					var v29 := "vphuxrsd";
					v28[156] := v29 + "f";
					v26[392] := 0xf9 * -f22;
				case DC8(cf10) =>
					var v30: multiset<bool> := multiset{f23, p0, p0, f23, f23};
					var v31 := DC12(v30);
					var v32: map<int, multiset<bool>> := map[f22 := v31.cf15];
					var v33: seq<bool> := [p0];
					var v34: array<multiset<bool>> := new multiset<bool>[17] [v30, v30, if (|multiset{p0}| in v32) then v32[|multiset{p0}|] else multiset{f23}, v30, v30, v30, v30, v30[f23 := f22], multiset(v33), v30 * v30, v30, v30, v30, v30, v30, v30, v30 - multiset(v33)];
					v34[184] := v30 + v30;
					var v35 := "hoka";
					var v36: multiset<string> := multiset{if (p0) then v35 else v35, v35 + v35, seq(874, i3  => (f18)), v35, v35[f22 := cf10]};
					var v37: map<int, int> := map[f22 := |v30|];
					r3, v34[184], v36, globalState.f6 := p0, v30 + v30, fm15(f18, -f22, v3, globalState), -((f22 - f22) * (if (|v0| in v37) then v37[|v0|] else f22));
					var v38: map<bool, bool> := map[p0 := f23];
					var v39: set<int> := {f22, f22, 0x127};
					var v40: seq<set<int>> := [v39];
					v38 := v38[v39 !! v40[-737] := p0];
					var v41: array<int> := new int[21](i4 => i4 * f22);
					v41[162] := f22;
					v41[162] := -f22;
					var v42 := DC14(if (p0) then p0 else f23, v41[162]);
					var v43: seq<int> := [|v35|];
					v42 := v42.(cf17 := v41[162]).(cf17 := v41[162]).(cf17 := |(v43 + v43)|);
			}
			
			var v44: seq<bool> := [f23];
			var v45: map<int, int> := map[f22 := f22];
			var v46: seq<map<int, int>> := [v45, v45];
			var v47 := "gr";
			var v48 := DC7(v44);
			var v49: array<bool> := new bool[15] [false, f23, v44[446], f23, v46 < fm16(globalState), f23, p0, f18 in v47, v48.cf9 == v44, p0, p0, p0 <== f23, !p0, fm1(p0, f18, f22, globalState), p0];
			v49[732] := p0;
			v49[732] := fm1(p0, if (!p0) then f18 else f18, fm0(globalState), globalState);
			var v50 := new C0("s" + v47, map['d' := f22], f18, f19 + f19);
			var v51: array<int> := new int[12];
			v51[194] := |v50.f20|;
			v51[194] := f22;
		}
		for i5 := f22 to f22 {
			var v52: seq<int> := [i5, -0x2ac];
			var v53: seq<seq<int>> := [[i5], v52];
			globalState.f6 := i5 * (|v53| * i5);
			var v54 := "isnmbr";
			r2, r1 := f23, -|v54|;
			r3 := p0;
			r2 := !p0;
		}
		var v55 := "sk";
		var i6 := 0;
		while (if (f23) then v55 != v55 else f23)
			decreases 100 - i6
		{
			if (i6 >= 100) {
				break;
			}
			
			i6 := i6 + 1;
			var v56: map<char, int> := map[f18 := f22];
			var v57 := new C0(v55, v56, if (f23) then f18 else f18, f19 + f19);
			var v58: seq<bool> := [p0];
			var v59: map<int, seq<bool>> := map[f22 := v58];
			var v60 := DC7((if (f22 in v59) then v59[f22] else [f23]) + v58);
			var v61: set<bool> := {p0, fm1(!!p0, f18, 169, globalState), f23};
			var v62: map<int, bool> := map[f22 := f23];
			var v63: array<bool> := new bool[4] [true, v61 !! {f23}, f23, if (f22 in v62) then v62[f22] else f23];
			v63[342] := p0;
			var v64: seq<map<int, bool>> := [v62, v62];
			globalState.f6, v60, r3, globalState.f6, v63[342] := f22, v60, true, f22, v62 == v64[|v57.f20|];
			v63[342] := v58[0x182];
			var v65 := DC8(if (false) then f18 else f18);
			match (v65) {
				case DC9(cf11, cf12) =>
					var v66 := new C0("isrngtcpi", if (cf11) then v57.f21 else map[f18 := |[f22, cf12]|], f18, f19);
					var v67: array<int> := new int[26];
					v67[341] := f22 + cf12;
					v67[341] := cf12;
					var v68: array<char> := new char[11] [f18, 'w', fm13(cf12, f22, |multiset{!p0}|, f23, globalState), f18, f18, if (cf11) then f18 else f18, fm13(v67[341], cf12, f22, !false, globalState), f18, f18, f18, f18];
					v68 := v68;
					r3 := cf11;
				case DC10(cf13) =>
					var v69 := DC14(v63[342] <==> cf13, f22 - f22);
					v69 := v69;
					r1 := -f22;
					var v70: array<string> := new string[6];
					v70[367] := v57.f20;
					v70[367] := v57.f20;
					var v71: array<D2> := new D2[14];
					v71 := new D2[9](i7 => v60);
				case DC8(cf10) =>
					cf10 := cf10;
					globalState.f6 := f22;
					var v72: set<int> := {f22, 0x201};
					var v73: seq<int> := [|v62|, f22];
					v72 := {v73[f22]} * v72;
					var v74: seq<multiset<int>> := [v2];
					var v75: set<seq<int>> := {[836, f22] + (seq(264, i8  => (f22))), [|v74|, |v73|], v73};
					v75 := (set v76 : seq<int> | v76 in [v73] :: (v76)) + v75;
			}
			
		}
		var v77: array<int> := new int[2];
		var v78 := DC8('r');
		var v79: multiset<string> := multiset{v55, seq(-325, i9  => (f18)), "xdamu"};
		var v80: array<bool> := new bool[13];
		var v81: seq<int> := [f22, f22, f22, f22, f22];
		var v82: map<array<bool>, int> := map[v80 := v81[f22]];
		var v83: seq<bool> := [f23];
		v77, v78, v79, v82, globalState.f13 := v77, v78, v79, map[v80 := f22], v83;
		r0 := DC0(f22);
		r1 := f22;
		r2 := !("yjptbhmhv" != (seq(423, i10  => (f18))));
		r3 := p0;
	}
	method m2(p0: bool, p1: int, p2: int, p3: int, globalState: GlobalState) returns (r0: int) {
		var v0: array<char> := new char[15] [f18, 'q', f18, f18, fm13(p3, p2, 0x3c8, p0, globalState), f18, f18, f18, f18, f18, f18, f18, f18, f18, f18];
		var v1: set<array<char>> := {v0, v0};
		globalState.f6 := fm4(!f23, if (f23) then f18 else f18, f18, v0 in v1, globalState);
		var v2 := "jdgvd";
		var v3: map<string, bool> := map[(v2 + v2)[p2 := f18] := !fm1(f23, f18, |v2|, globalState)];
		var v4: map<int, string> := map[f22 := v2];
		v3 := v3[(if (p1 in v4) then v4[p1] else v2) + v2 := false];
		var i0 := 0;
		while (p1 >= p2)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			globalState.f6 := p2 - p1;
			var v5: map<int, char> := map[p3 := f18];
			var v6 := DC4(p1, p2, f23);
			r0 := fm4(fm12(globalState), f18, if (v6.cf4 in v5) then v5[v6.cf4] else 'y', f23, globalState);
			f22 := p1;
			var v7: array<int> := new int[7];
			v7[352] := p2 % |v5|;
			v7[352] := 365;
		}
		var v8: seq<bool> := [p0, false];
		if (p1 >= (p1 % |v8|)) {
			var v9: multiset<bool> := multiset{false, f23};
			v9 := v9 + v9;
			v8 := v8 + v8;
			var v10: set<int> := {0x12b};
			v10 := fm17(globalState);
			var v11: map<int, int> := map[f22 := p3];
			var v12: map<char, int> := map['s' := |v11|];
			var v14: seq<string> := [v2, v2];
			var v15: C0 := new C0(v2, v12, f18, map v13 : string | v13 in v14[p3 := v2] :: (v13) := (|multiset{f18, f18, 'f', 't', f18}|));
			var v16: map<C0, C0> := map[v15 := v15];
			var v17: map<bool, C0> := map[f23 := if (v15 in v16) then v16[v15] else v15];
			v17 := v17[v8 != v8 := v15];
			var v18 := DC15(v15.f21);
			var v19: multiset<map<char, int>> := multiset{v12 + v12, v18.cf18};
			v19 := v19;
		} else {
			r0 := p1;
			if (f23) {
				var v20: map<bool, char> := map[f23 := f18];
				v20 := v20[!f23 := f18];
				var v21: map<D1, bool> := map[DC6(!!false, true) := p0];
				var v22 := DC6(f23, p0);
				v21 := v21[v22 := f23];
				f22 := p2;
				globalState.f6 := -f22;
				var v23: map<bool, bool> := map[f23 := p0];
				var v24: array<int> := new int[11] [|v23| - |v8|, 0x2e8 + p1, p1, f22, fm4(f23, f18, f18, p0, globalState), fm0(globalState), p1, p3, -0x31e, p2, p2];
				var v25: seq<int> := [|v8|];
				var v26: multiset<bool> := multiset{f23, p0};
				v24[250] := v25[p2] % |v26|;
				v24[250] := p3;
			} else {
				var v27: T0 := new C0(v2, fm14(0x164, globalState), f18, f19);
				var v28: seq<T0> := [v27];
				var v29: map<bool, seq<T0>> := map[f23 := v28 + v28];
				var v30: array<bool> := new bool[25];
				var v31: map<char, int> := map['i' := f22];
				var v32: C0 := new C0("v", v31, f18, v27.f19);
				var v33: map<C0, bool> := map[v32 := !f23];
				v30[163] := if (v32 in v33) then v33[v32] else p0;
				v29, v30[163] := v29[f23 := v28], f23;
				var v34: map<bool, bool> := map[f23 := f23 <==> f23];
				v34 := v34[(seq(0x2a, i1  => (v27.f18))) <= v2 := f23];
				var v35: seq<int> := [f22, p1, p2];
				var v36: set<int> := {0x130, |v35|};
				v30[163] := fm1(p0, v27.f18, |v36| - v35[|(seq(-0x32d, i2  => (v27.f18)))|], globalState);
				var v37: seq<string> := [v32.f20, v2, v2 + v2];
				globalState.f6, v37, globalState.f6 := v35[|v3|], v37, DC14(!false, p1).cf17 % (-0xdd / 0x112);
				v2 := v32.f20[p1 := v27.f18];
			}
			
			if (fm3(!(p0 && f23), |v2| + |v2|, globalState)) {
				f18 := f18;
				var v38 := false;
				v38 := !p0;
				f22 := p2;
				var v39 := DC4(p3, p1, v38);
				var v40: map<bool, bool> := map[v39.cf5 := f23];
				v40 := v40[f23 := f23];
				f22 := -p2;
			} else {
				var v41: multiset<int> := multiset{p1, p1, p1, p1};
				var v42 := DC10(v41 >= v41);
				v42 := if (f23) then DC10(f23) else v42;
				var v43: array<int> := new int[10](i3 => i3 % p3);
				v43[753] := p3;
				var v44: set<int> := {p2, p3};
				v43[753] := p2 + (693 * |v44|);
				var v45: seq<array<char>> := [v0, v0, v0, v0, v0];
				var v46: seq<seq<array<char>>> := [v45];
				var v47: set<string> := {"tqhnwiyw", v2, v2};
				var v48: map<string, seq<array<char>>> := map[v2 := v46[|v47|]];
				v48 := v48[v2 + "jrv" := v45];
				var v49 := true;
				var v50: multiset<char> := multiset{v2[p2], f18};
				v49 := if (v49) then v8[if (f18 in v50) then v50[f18] else |v41|] else |v50| <= v43[753];
				var v51: map<bool, int> := map[v49 := f22];
				v51 := v51[p0 := 0x34b - p2];
			}
			
			if (v8[f22]) {
				var v52: array<array<bool>> := new array<bool>[24];
				var v53 := DC16(!f23, v52);
				var v54 := DC17(v53);
				var v55: map<D6, char> := map[v54 := 'y'];
				v55 := v55[v54 := f18];
				f22 := |(seq(0x29f, i4  => (p3 / |map[false := p0]|)))|;
				var v56: multiset<int> := multiset{527};
				var v57: set<map<int, char>> := {map[p2 := f18]};
				var v58: set<bool> := {f23};
				var v59: array<bool> := new bool[27] [p0, f23, f23, !fm1(p0, f18, |v56|, globalState) <== true, !f23, p0, true, fm3(false, |[f22]|, globalState), p0, !f23, f23, f23, true, v8[p2], p0, f22 >= p2, v57 > v57, p0, true, p0, !f23, f23, f23, v58 !! v58, p0, p0, f23];
				v59[600] := p0;
				v59[600] := if (p1 <= |v2|) then p0 else p2 != p2;
				var v60 := DC4(p2, 0xc0, f23);
				var v61: seq<int> := [fm4(fm1(f23, f18, p3, globalState), f18, 'd', !p0, globalState)];
				var v62: set<int> := {p3, f22, 241, p2, |v56|};
				var v63: array<D1> := new D1[24] [v60, v60, DC4(p1, p1, true), v60, DC4(p1, p3, f23), v60, v60, DC4(|v61|, |v62|, false), v60, v60, v60, v60, v60, v60, DC4(p3, p3, f23), v60, v60, v60, v60, v60, DC4(f22, p2, p0), v60, DC4(|v8|, f22, v8[p3]), v60];
				var v64: multiset<array<D1>> := multiset{v63, v63, v63};
				v64 := v64;
				var v65: map<char, int> := map[f18 := p3];
				var v66 := new C0("scd", v65, f18, f19);
			} else {
				globalState.f6 := p3;
				var v67: multiset<int> := multiset{|v8|, p3, 542};
				f22, v67 := p2, v67;
				var v68 := DC8(fm13(-p2, p1, p2, f23, globalState));
				var v69: map<D3, bool> := map[v68 := f23];
				var v70: set<bool> := {p0, !p0, p0};
				v69 := v69[v68.(cf10 := 'a') := {p0} < v70];
				var v71 := true;
				v71 := v71;
				var v72: array<int> := new int[23];
				v72[849] := p3 % p3;
				var v73: map<int, int> := map[f22 := 0x3a3];
				v72[849], globalState.f6, v71 := f22, (if (p1 in v67) then v67[p1] else if (f22 in v73) then v73[f22] else 0x2ba) / -(if (v71) then p1 else f22), p0 && false;
			}
			
			var v74 := true;
			v74 := fm1(v1 <= v1, 'y', p3, globalState);
		}
		
		globalState.f6 := -0x65;
		var i5 := 0;
		while (v8[0x129])
			decreases 100 - i5
		{
			if (i5 >= 100) {
				break;
			}
			
			i5 := i5 + 1;
			var v75: array<int> := new int[4](i6 => i6 + p3);
			var v76: map<bool, bool> := map[f23 := false];
			var v77: map<int, int> := map[p1 := -|v76|];
			v75[240] := p1 / |v77|;
			v75[240] := p1;
			var v78: multiset<int> := multiset{|[p0, f23]|, p1};
			var v79 := DC9(p0, |v78|);
			var v80 := DC4(v79.cf12, f22, f23);
			v80 := v80;
			globalState.f6 := f22;
			v75 := v75;
		}
		r0 := p1 + (0x257 % f22);
	}
}

class C2 extends T1 {
	constructor (f18 : char, f19 : map<string, int>) {
		this.f18 := f18;
		this.f19 := f19;
	}
	
	function fm3(p0: bool, p1: int, globalState: GlobalState): bool {
		if (|"fywgn"| <= 896) then |"bnebdr"| <= 0x149 else true
	}
	function fm4(p0: bool, p1: char, p2: char, p3: bool, globalState: GlobalState): int {
		-(|([false] + [false, false, !true, false])| + |(seq(0x3cb, i0  => (|[true, !!false, true, !!true, !false]|)))|)
	}
	method m1(p0: bool, globalState: GlobalState) returns (r0: D0, r1: int, r2: bool, r3: bool) {
		var v0: array<int> := new int[15](i0 => i0 - |"fugltux"|);
		var v1 := -0x25d;
		v0[438] := v1;
		var v2: array<map<int, D0>> := new map<int, D0>[28](i1 => map[|map["yxgl" := p0]| := DC0(|(seq(-0x229, i2  => (v1)))|)] + map[v1 := DC0(v1)]);
		v0[969] := v1;
		var v3 := "uuruhxs";
		var v4: map<int, string> := map[v1 := v3];
		var v5: array<bool> := new bool[5];
		var v6: multiset<array<bool>> := multiset{v5, v5};
		var v7: seq<bool> := [!p0];
		var v8 := DC7(v7);
		r2, f18, v0[438], v2, v0[969] := |v4| <= (579 * |v6|), f18, v1, v2, match v8 {
			case DC7(cf9) => v1 * v1
		};
		var v9: map<char, int> := map[f18 := v0[438]];
		var v10 := new C0(v3, v9, f18, f19);
		v1 := v1 * 988;
		var v11: map<int, int> := map[v0[438] := v0[438]];
		var v12: seq<int> := [|v11| + v1, v0[438], v0[438]];
		r1 := v12[v1];
		for i3 := v0[438] to v1 {
			var v13 := m4(i3 * v1, fm0(globalState), v5, globalState);
			r3 := v13;
			var v14 := DC1();
			match (v14) {
				case DC1() =>
					var v15: multiset<bool> := multiset{v13};
					v0[438] := |v15|;
					globalState.f13 := (v7 + v7) + [v13];
					v5[264] := false;
					v5[56] := v13;
					var v16: multiset<array<int>> := multiset{v0, v0, v0};
					var v17: seq<multiset<array<int>>> := [v16, v16];
					v13, globalState.f6, v13, v5[264], v5[56] := v13, fm0(globalState), v17[v1] != v16[v0 := |v10.f20|], i3 >= (v0[438] % v1), p0;
					v13 := v13;
				case DC0(cf0) =>
					var v18: set<bool> := {v13};
					v18 := v18;
					globalState.f6 := v1 * v1;
					var v19: array<D2> := new D2[20];
					v19[883] := v8.(cf9 := [v13]);
					v19[883] := v8;
					var v20: map<string, bool> := map[v3 := !p0];
					var v21: map<int, map<string, bool>> := map[if (v13) then -0xc4 else i3 := map[v10.f20 := v13] + v20];
					v21 := v21[v0[438] := if (v13) then fm11(globalState) else v20[v10.f20 := true]];
				case DC2(cf1) =>
					var v22: T0 := new C0("qfwpi", map[f18 := v1], 'm', map[v10.f20[v0[438] := f18] := v0[438]]);
					var v23 := DC11(v22);
					var v24 := DC11(v23.cf14);
					v22 := (v24.(cf14 := v22)).cf14;
					var v26: multiset<string> := multiset{v3};
					var v27 := new C0(v10.f20, v9, 'l', map v25 : string | v25 in v26 :: (v25) := (|v10.f20|));
					var v28: map<C0, int> := map[v10 := -v1];
					v13, v1 := (if (v27 in v28) then v28[v27] else v0[438]) < v1, 0x6;
					var v29: T1 := new C1(v1, v13, f18, fm18(|v10.f20|, globalState));
					var v30: map<seq<bool>, T1> := map[v7 := v29];
					var v31: map<bool, map<seq<bool>, T1>> := map[false := map[v7 := v29]];
					v30, v1 := (if (p0 in v31) then v31[p0] else v30) + v30, v1 - |(seq(357, i4  => (v29.f18)))|;
			}
			
			r2 := !p0;
		}
		v5[496] := p0;
		v0[438], v5[496], f18, v1 := v0[438], v10.f20 == v10.f20, f18, v1;
		r0 := DC0(v0[438] - fm4(p0, f18, 'r', v5[496], globalState));
		r1 := if (v0[438] in v11) then v11[v0[438]] else |(seq(380, i5  => (v11)))|;
		r2 := if (v5[496]) then p0 else v0[438] < v1;
		var v32: set<bool> := {!false};
		r3 := v32 >= v32;
	}
	method m2(p0: bool, p1: int, p2: int, p3: int, globalState: GlobalState) returns (r0: int) {
		r0 := p2;
		if (p0) {
			r0 := p3;
			r0 := p3;
			var v0: array<bool> := new bool[20];
			v0 := v0;
			v0[256] := p0;
			v0[256] := p3 > -p3;
			v0[256] := true <==> (p2 > p3);
		} else {
			var v1 := false;
			var v2: set<bool> := {v1};
			v1 := v2 >= v2;
			var v3 := "wrx";
			var v4: map<bool, bool> := map[v1 := !v1];
			var v5: map<char, int> := map[v3[|v4|] := p2];
			var v6 := DC15(v5 + v5);
			var v7: array<string> := new string[8] [v3, v3, v3, v3[p2 := 'p'], v3, v3, v3, v3];
			var v8: array<bool> := new bool[17];
			var v9: multiset<bool> := multiset{v1, !false, v1, fm1(v1, 'r', p2, globalState)};
			v8[138] := v9 >= v9;
			var v10 := DC0(-p1);
			v6, v7, v8[138], v10, f18 := DC15(v5 + v5), v7, fm1(v1, f18, -405, globalState), DC0(p2), f18;
			var v11 := DC6(true, p0);
			v8[138] := v11.cf7;
			var v12: map<int, bool> := map[p1 := !v8[138]];
			var v13: map<int, bool> := map[|v3| := if (p3 in v12) then v12[p3] else p0];
			v13 := v13[p1 := v1];
			globalState.f6 := fm0(globalState);
		}
		
		var v14 := true;
		var v15: set<bool> := {p0};
		var v16 := DC18(v15);
		v14 := v15 == v16.cf22;
		var v17: map<bool, bool> := map[v14 := p0];
		var v18: map<int, int> := map[|(v17[p0 := v14] + v17)| := p2];
		var v19 := "oiyohkc";
		var v20: array<int> := new int[15];
		v20[605] := p1;
		var v21: multiset<bool> := multiset{v14, v14};
		var v22 := DC12(v21);
		v18, v19, v20[605] := fm19(p0, p0, globalState) + map[p2 := p2], v19, match v22 {
			case DC13() => |multiset([|"k"|] + DC21([p2]).cf26)|
			case DC14(cf16, cf17) => -|(seq(0x17a, i0  => ('s')))|
			case DC12(cf15) => p1
		};
		var v23: map<bool, int> := map[v14 := p3];
		v23 := v23[p0 := p2];
		var v24: multiset<int> := multiset{296};
		var v25: C1 := new C1(p3, v20[605] in v24, 'y', f19 + f19);
		var v26: seq<C1> := [v25];
		v25 := v26[p3];
		r0 := p1;
	}
	method m4(p0: int, p1: int, p2: array<bool>, globalState: GlobalState) returns (r0: bool) {
		for i0 := p1 + p1 to p1 {
			var v0 := false;
			var v1: map<bool, int> := map[v0 := p0];
			var v2: multiset<bool> := multiset{v0};
			var v3: set<int> := {|v1|, |v2|};
			var v4: map<set<int>, int> := map[v3 := p1 * p0];
			v4 := v4[v3 := fm0(globalState)];
			var v5 := "gtyc";
			globalState.f6 := p1 - |[v5, v5]|;
			v0 := true;
			var v6: map<string, bool> := map[v5 := v0];
			p2[135] := if (v5 in v6) then v6[v5] else v0;
			var v7 := DC13();
			var v8: map<D5, int> := map[v7 := i0];
			p2[135], v8, v0 := v0 !in fm2(globalState), v8, fm1(v0 || v0, 'v', -p1, globalState);
		}
		var v9 := false;
		var i1 := 0;
		while (v9)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			globalState.f6 := p0 - --p1;
			var v10 := "yesur";
			var v11: map<bool, char> := map[v9 := v10[p1]];
			var v12: seq<map<bool, char>> := [map[v9 := 'c']];
			var v13: map<int, string> := map[p0 := v10];
			v11 := v12[|(if (-0x300 in v13) then v13[-0x300] else seq(0xaf, i2  => (f18)))|];
			r0 := fm1(v9, f18, p1, globalState);
			var v14: map<bool, bool> := map[!v9 := v9];
			var v15: seq<bool> := [v9];
			var v16: map<int, char> := map[-0x26f := f18];
			var v17: map<seq<bool>, map<int, char>> := map[v15 := v16];
			var v19: set<seq<bool>> := {v15, v15, v15};
			var v20: array<bool> := new bool[21] [v9, v9, if (v9) then true else !(if (v9 in v14) then v14[v9] else v9), v9, v9, v9 in v15, v9, true, true, !(if (v9) then v9 else true), v9, v15[|map[p0 := true]|], v9, v9, v9 ==> v9, v9, (set v18 : seq<bool> | v18 in v17 :: (v18)) !! v19, v9, v9, v9, v9];
			var v21: map<char, int> := map[v10[0x13c] := p1];
			var v22: T0 := new C0(seq(0xd6, i3  => (f18)), v21, f18, f19);
			var v23: map<bool, array<bool>> := map[v9 := v20];
			v20, globalState.f6, v22, f18 := p2, |(if (v9) then v23 else v23 + v23)|, if (v9) then v22 else v22, fm20(v9, globalState);
		}
		var v24: array<array<bool>> := new array<bool>[28];
		v24[400] := p2;
		v24[400] := p2;
		var v25: multiset<set<int>> := multiset{{-174, p1}};
		var v26: map<bool, int> := map[v9 := p0];
		var v27 := "jcyiogjr";
		var v28: set<int> := {|v26|, |v27|};
		var v29: seq<int> := [p0];
		v25 := (multiset{v28})[v28 + (set v30 : int | v30 in v29 :: (v30 / |(set v32 : int | v32 in (map v31 : int | (0x32e <= v31) && (v31 < 0x1de) :: (v31 + -0x291) := (0x34e)) :: (v32 * |(seq(0x373, i4  => ('t')))|))|)) := fm0(globalState)];
		var i5 := 0;
		while (v28 !! v28)
			decreases 100 - i5
		{
			if (i5 >= 100) {
				break;
			}
			
			i5 := i5 + 1;
			globalState.f6 := p0;
			globalState.f6 := p0;
			var v33: map<char, int> := map[f18 := p0];
			var v34: T0 := new C0(v27, v33, f18, f19);
			var v35 := DC11(v34);
			match (v35) {
				case DC11(cf14) =>
					var v36: seq<bool> := [v9];
					v27 := ((fm21(globalState))[p1 := f18])[p0 := cf14.f18] + (v27 + v27)[|v36| := f18];
					var v37: seq<map<bool, int>> := [v26];
					var v38: map<int, seq<map<bool, int>>> := map[p0 := fm22(globalState)];
					var v39: seq<seq<map<bool, int>>> := [if (p1 in v38) then v38[p1] else v37];
					v37 := v39[(DC14(v9, p1).(cf16 := !v9).(cf17 := p0)).cf17];
					var v40: map<string, bool> := map[v27 := v9];
					v40 := v40[v27 := v9];
					globalState.f6 := -fm0(globalState);
			}
			
			var v41 := DC21(v29);
			v41 := v41;
		}
		var v42: array<int> := new int[9];
		forall i6 | 0 <= i6 < v42.Length {
			v42[i6] := i6 + (p0 % p1);
		}
		r0 := !v9;
	}
}

class C3 extends T1 {
	constructor (f18 : char, f19 : map<string, int>) {
		this.f18 := f18;
		this.f19 := f19;
	}
	
	function fm3(p0: bool, p1: int, globalState: GlobalState): bool {
		((seq(0x202, i0  => (|[0x3a9, 0x2a6]|))) + (seq(0x370, i1  => (307)))) != (seq(0x287, i2  => (-125)))
	}
	function fm4(p0: bool, p1: char, p2: char, p3: bool, globalState: GlobalState): int {
		|(multiset{265} * multiset{0x175, 0x35a})| % |("hkdhj" + "wtgkshg")|
	}
	method m1(p0: bool, globalState: GlobalState) returns (r0: D0, r1: int, r2: bool, r3: bool) {
		var i0 := 0;
		while (p0)
			decreases 100 - i0
		{
			if (i0 >= 100) {
				break;
			}
			
			i0 := i0 + 1;
			var v0 := 652;
			var v1: multiset<int> := multiset{v0};
			var v2: seq<char> := [f18];
			var v3: set<int> := {-v0, v0};
			var v4: seq<bool> := [false, p0];
			var v5 := DC7(v4);
			var v6: map<seq<bool>, int> := map[v5.cf9 := v0];
			var v7: map<int, bool> := map[|v6| := p0];
			var v8: map<int, bool> := map[|v7| := p0];
			var v9: array<int> := new int[21](i1 => i1 % v0);
			var v10: map<array<int>, bool> := map[v9 := fm3(p0, v0, globalState)];
			var v11: array<bool> := new bool[23] [p0 || p0, v1 <= v1, false, DC4(v0, v0, true).cf5, v2 == v2, false, p0, p0, p0 || true, if (p0) then p0 else p0, p0, p0, !p0, !(fm5(v0, globalState) !! v3), |v2| > |v7|, v0 < v0, p0, p0, if (v9 in v10) then v10[v9] else p0, v0 > v0, v1 !! v1, p0, p0 || p0];
			v11[994] := p0;
			v9[977] := v0;
			r3, globalState.f6, v11[994], v9[977] := v0 > v0, v0 * v0, p0, 620;
			var v12: map<int, int> := map[|v2| := v0];
			v7 := (map[v0 := p0])[-(v9[977] * |v12|) := v11[994]];
			var v13 := new C0((if (DC3(p0).cf2) then v2 else "as")[v9[977] := f18], fm7(v11[994], v0, globalState), f18, f19 + f19);
			var v14 := new C0(v13.f20, v13.f21, f18, f19);
		}
		var v15 := 0x356;
		var v16: multiset<int> := multiset{v15};
		var v17: array<int> := new int[21] [v15, v15, v15, v15, 200, v15, v15, v15, -v15, 0x1f1, v15, v15, v15, |v16|, v15, v15, v15, v15, v15, 192, -38];
		var v18: map<array<int>, char> := map[v17 := f18];
		var v19 := "kbgkabcen";
		for i2 := |v18| to |v19| - v15 {
			var v20 := DC8(f18);
			var v21: map<bool, bool> := map[p0 := p0];
			var v22: seq<bool> := [p0];
			var v23: map<map<bool, bool>, int> := map[v21 := |v22|];
			var v24: seq<bool> := [p0, p0 ==> fm1(p0, v20.cf10, if (v21 in v23) then v23[v21] else i2, globalState), p0];
			r1 := |v22|;
			globalState.f6 := v15;
			globalState.f6 := v15;
			r1 := if ("pqpigsw" != "mlp") then v15 else v15;
		}
		for i3 := v15 to v15 {
			var v25: array<string> := new string[17](i4 => v19);
			v25[962] := v19;
			v25[962] := v19 + v19;
			v17 := if (false) then v17 else v17;
			var v26: map<int, bool> := map[i3 := p0];
			var v27: set<bool> := {p0};
			var v28: seq<bool> := [if (|v16| in v26) then v26[|v16|] else p0, p0, p0, fm1(p0, f18, |v27|, globalState), p0];
			var v29: seq<bool> := [true, v28[v15]];
			var v30: map<char, int> := map['j' := |v19|];
			var v31: T0 := new C0(v25[962], v30, f18, f19);
			var v32: set<T0> := {v31, v31, v31};
			globalState.f6, r3 := if (v28[v15]) then v15 else i3, !(v32 > (v32 + v32));
			globalState.f6 := i3;
		}
		var v33: array<bool> := new bool[9](i5 => p0);
		v33 := v33;
		var v34 := DC3(false);
		r2 := !v34.cf2 && p0;
		var v35 := new C0(if (p0) then v19 else v19, fm7(p0, v15, globalState), 'g', f19);
		var v36 := DC0(-v15);
		r0 := v36;
		var v37: seq<int> := [v15, v15, v15];
		r1 := v15 * (|v37| - v15);
		r2 := p0;
		r3 := p0;
	}
	method m2(p0: bool, p1: int, p2: int, p3: int, globalState: GlobalState) returns (r0: int) {
		var v0: array<int> := new int[8](i0 => i0 + p3);
		v0[974] := p3;
		var v1 := true;
		var v2: set<bool> := {v1};
		var v3: map<int, set<bool>> := map[p3 := v2];
		var v4 := DC9(v1, fm0(globalState));
		v0[974], v1, v3 := -match v4 {
			case DC9(cf11, cf12) => cf12
			case DC10(cf13) => p2 % p2
			case DC8(cf10) => p3
		}, false, v3;
		var i1 := 0;
		while (v1)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			globalState.f6 := v0[974] * p3;
			var v5 := "qxtvrh";
			var v6: map<char, int> := map[f18 := 308];
			var v7 := new C0(v5, v6, f18, f19[v5 := v0[974]]);
			v0[974] := fm0(globalState);
			if ((p3 - p2) <= -v0[974]) {
				var v8: seq<int> := [v0[974]];
				globalState.f1 := v8 + [p2];
				var v9: map<int, map<int, bool>> := map[p2 % p1 := map[p1 := p0]];
				v9 := v9[v0[974] := map[fm0(globalState) := p0]];
				var v10 := DC3(true);
				var v11: map<int, D1> := map[p3 := v10];
				var v12: T0 := new C0(fm8(globalState), v7.f21, f18, (map[v7.f20 := |v11|])["vnfqve" := -fm0(globalState)]);
				var v13: array<bool> := new bool[23](i2 => p0);
				var v14: map<T0, array<bool>> := map[v12 := v13];
				v14 := v14[v12 := v13];
				var v15: seq<bool> := [true, v1];
				var v16: map<int, bool> := map[p1 := true];
				var v17 := DC7(v15[|v16| := p0]);
				var v18: multiset<D2> := multiset{v17};
				var v19: seq<D2> := [DC7(v15)];
				var v20: seq<multiset<D2>> := [multiset(v19)];
				var v21: array<multiset<D2>> := new multiset<D2>[11] [v18 * v18, multiset{v17, v17, v17, DC7(v15), DC7([p0, p0, false, v1, p0])}, v18, v18 - v18, multiset{v17, v17}, fm9(globalState), v18, v20[v0[974]], v20[v0[974]], v18, (multiset{v17})[v17 := fm4(v1, f18, 'u', v1, globalState)]];
				v21[790] := v18;
				var v22: multiset<bool> := multiset{v1};
				v21[790] := v18[v17 := |v22|] * v18;
				var v23: set<int> := {-0xc0, v0[974], v0[974], v0[974]};
				v1 := (v23 + v23) >= v23;
			} else {
				var v24: map<string, bool> := map[fm8(globalState) + v7.f20 := v1];
				v1 := if (("fvox" + v7.f20) in v24) then v24["fvox" + v7.f20] else v1;
				var v25: array<seq<int>> := new seq<int>[19];
				var v26: set<int> := {p3};
				var v27: seq<int> := [p1, v0[974], p1, |v26|, p1];
				v25[462] := v27;
				var v28: array<set<bool>> := new set<bool>[22](i3 => v2);
				v28[562] := {v1, v1, v1, p0, !v1};
				var v29: array<bool> := new bool[19];
				v29[615] := v1;
				var v30: seq<bool> := [!v1];
				v1, v4, v25[462], v28[562], v29[615] := v30[p1], DC9(p0, p3), (v27 + v27) + v27, (v2 + v2) * {p0}, (if (p0) then v27 else v27) < (v27 + v27);
				v0[974] := 600;
				v29[615] := v30[p1];
				v29 := v29;
			}
			
		}
		var v31: map<int, bool> := map[p3 := p0];
		var v32: seq<map<int, bool>> := [v31, v31, v31];
		v0[974] := -(if (p0 ==> !v1) then v0[974] else |multiset((seq(-0x93, i4  => (map[|multiset{f18, f18, f18}| := v1]))) + v32)|);
		var v33: map<int, int> := map[v0[974] := 0x87];
		for i5 := p3 to |v33| {
			var v34 := DC5(p1);
			var v35: array<D1> := new D1[6] [v34, if (v1) then v34 else DC5(0x2cd), v34, v34, v34, v34];
			v35[63] := v34;
			v35[63] := DC5(p2);
			v1 := p0 && v1;
			var v36: array<bool> := new bool[10];
			v36[991] := !p0;
			var v37: multiset<int> := multiset{-p1, p3, p1};
			var v38: map<bool, bool> := map[v1 := p0];
			var v39: T1 := new C2(f18, f19);
			var v40: map<T1, bool> := map[v39 := v1];
			var v41 := DC4(p1, |v40|, v1);
			v36[991], v2, v1 := fm3(v1, if (-p3 in v37) then v37[-p3] else |v38|, globalState) && (if (p0) then fm1(p0, fm10(globalState), i5, globalState) else v1), v2, v41.cf5;
			var v42: array<array<array<int>>> := new array<array<int>>[15];
			var v43: array<array<int>> := new array<int>[14];
			v42[390] := v43;
			var v44: seq<bool> := [v1];
			var v45: set<seq<bool>> := {(v44 + [v36[991], v1, v1, true, !p0])[p3 := v1]};
			v1, v42[390], v45 := v36[991], v43, v45;
		}
		var v46: C2 := new C2('q', f19);
		var v47: map<C2, bool> := map[v46 := p0];
		var v48: set<map<C2, bool>> := {v47, v47, v47, v47};
		var v49: seq<set<map<C2, bool>>> := [v48];
		var v50: multiset<bool> := multiset{p0};
		if (!(v49[|v50|] != (v48 - v48))) {
			var v51: seq<bool> := [p0];
			var v52: array<bool> := new bool[29] [p0, p0, true, p0, v1, v1, v1, v51[p1], v1, p0, v1, p0, v1, v1, v1, v1, p0, p0, v1, true, v1, p0, v1, p0, p0, v1, v51[-|map[v1 := v0]|], p0, p0];
			var v53: map<bool, array<bool>> := map[p0 := v52];
			var v54: array<string> := new string[12];
			var v55: map<array<bool>, array<string>> := map[if (v1 in v53) then v53[v1] else v52 := v54];
			var v56: map<array<string>, bool> := map[if (v52 in v55) then v55[v52] else v54 := p0];
			v56 := v56[v54 := p0];
			var v57: map<bool, bool> := map[p0 := p0];
			var v58: seq<map<bool, bool>> := [v57];
			var v59: multiset<map<bool, bool>> := multiset{v57, v57, v57, v57 + v58[p3], fm23(p1, v1, globalState)};
			v59 := v59;
			var v60 := new C1(p2 + p2, p0, f18, fm18(v0[974], globalState));
			var v61: seq<int> := [p2];
			var v62: map<seq<int>, bool> := map[v61 := v1];
			v1 := v1 ==> (true || (if ((seq(-491, i6  => (0x83))) in v62) then v62[seq(-491, i6  => (0x83))] else false));
			v0[974] := |v51[|v2| := !v1]| % |v31|;
		} else {
			globalState.f6, r0, v1 := p2, p2, if (v2 !! v2) then true else !false;
			var v63: seq<array<int>> := [v0, v0, v0, v0, v0];
			var v64 := DC4(-v0[974], v0[974], v1);
			var v65: map<bool, int> := map[p0 := v0[974]];
			var v66: multiset<int> := multiset{p1};
			var v67 := "wyt";
			var v68: array<bool> := new bool[28] [p0, v1, v1, v1, !true, p0, v1, v0 !in v63, v1, !p0, if (v0[974] in v31) then v31[v0[974]] else v64.cf5, p0 in v65, v1, v66 >= v66, v1, v4.cf11 <==> p0, if (p2 in v31) then v31[p2] else !p0, true, true, v1 ==> p0, !p0, v67 != v67, p0, if (v1) then false else v1, v1, !true, if (!v1) then p0 else v1, p0];
			v68[459] := v1;
			var v69: map<multiset<bool>, int> := map[v50 := p2];
			var v71: set<multiset<bool>> := {v50, v50};
			var v72: map<multiset<bool>, bool> := map[v50 := p0];
			v68[459], v1, v67 := ((set v70 : multiset<bool> | v70 in v69 :: (v70)) + v71) > (set v73 : multiset<bool> | v73 in v72[v50 := p0] :: (v73)), v1, v67;
			if (true) {
				v1 := (if (false) then false else p0) <== v68[459];
				v33 := v33;
				var v75: seq<bool> := [v68[459]];
				var v76 := DC7(v75);
				var v77: seq<D2> := [v76, DC7(v75)];
				var v78: seq<seq<D2>> := [v77];
				var v79: map<map<int, bool>, seq<D2>> := map[map v74 : int | (0x130 <= v74) && (v74 < -0x3b1) :: (v74 - p1) := (p0) := v78[p3]];
				var v80: map<set<bool>, map<map<int, bool>, seq<D2>>> := map[v2 := v79];
				v80 := v80[if (v68[459]) then v2 else {v68[459], p0, !v68[459]} := v79];
				v68[459] := p0;
				r0 := p3;
			} else {
				var v81 := DC8('d');
				var v82: map<char, int> := map[v81.cf10 := p3];
				var v83 := new C0(v67, v82, f18, f19 + f19);
				var v84: C0 := new C0(fm8(globalState), v83.f21, f18, f19);
				v84 := v83;
				var v85 := DC6(v1, v68[459]);
				var v86: map<int, D1> := map[p3 := v85];
				v86 := v86[v0[974] := v85];
				var v87 := new C0(if (!!true) then v84.f20 else v67, v82 + v82, f18, f19);
				globalState.f6 := (p3 - p1) % 259;
			}
			
			var v88: map<bool, seq<bool>> := map[false := [v1, p0]];
			var v89 := DC10(!p0);
			var v90: map<map<bool, seq<bool>>, D3> := map[v88 := v89];
			v90 := v90[v88 := fm24(globalState)];
			var v92: map<char, bool> := map[f18 := v1];
			var v93 := DC15(map v91 : char | v91 in v92 :: (v91) := (|"y"|));
			var v94 := DC17(v93);
			var v95: map<D6, bool> := map[v94 := v68[459]];
			var v96: seq<map<D6, bool>> := [v95];
			var v97: array<map<D6, bool>> := new map<D6, bool>[29] [v95, v95 + v95[v94 := p0], v95, v95, v95 + v95, v95, v95, v95, v95, v95, v96[fm0(globalState)], v95, v95, v95, v95, v95, v95, if (v1) then v95 else v95, v95, v95, v95[v94 := fm3(!v68[459], if (v68[459] in v50) then v50[v68[459]] else |v67|, globalState)], v95, v95, v95, v95 + v95[v94 := v68[459]], v95 + v95, v95, map[DC17(DC15(map[f18 := p3])) := v1], v95];
			v97 := v97;
		}
		
		var v98 := "dv";
		var v99 := DC22(v98);
		var v100: array<string> := new seq<char>[16] [seq(0x29e, i8  => ('a')), v98, v99.cf27, "ude", v98, v98, v98, v98 + v98, seq(168, i9  => (f18)), v98 + v98, v98, v99.cf27, v98 + v98, v98, v98, "o" + v98];
		forall i7 | 0 <= i7 < v100.Length {
			v100[i7] := v98;
		}
		r0 := p3;
	}
	method m3(p0: int, p1: bool, p2: multiset<int>, p3: seq<seq<char>>, globalState: GlobalState) returns (r0: int) {
		var v0 := "bfbxyfb";
		var v1: map<int, bool> := map[|p2| := p1];
		var v2: seq<int> := [|[p1, p1]|, p0, p0, p0, |v0|];
		for i0 := -|v0| to |v1| - v2[0x2fc] {
			v0 := v0;
			var v3: array<map<int, int>> := new map<int, int>[6];
			var v4: map<int, int> := map[p0 := |p2|];
			v3[973] := v4;
			v3[973] := if (p1 && p1) then v4 else v4;
			var v5, v6, v7 := m0(fm4(p1, 'e', f18, p1, globalState), p1, i0, globalState);
			var v8: multiset<multiset<int>> := multiset{p2, p2, multiset{v7, fm0(globalState), -v7, p0, |v0|}, multiset{i0} + p2};
			var v9: set<int> := {fm0(globalState), |v0|, |v0|};
			r0, globalState.f6, globalState.f6, v0 := i0, if (p2 in v8) then v8[p2] else v7, |v9| / fm0(globalState), fm8(globalState);
		}
		var v10: C2 := new C2(f18, f19);
		var v11: seq<C2> := [v10];
		var v12: map<bool, seq<C2>> := map[p1 := if (p1) then v11 else v11];
		v12 := (v12 + map[false := v11]) + map[p1 := v11];
		var v13: multiset<bool> := multiset{false, false, !p1, p1, v0 <= v0};
		v13 := v13;
		var v14 := true;
		v14 := !v14;
		v14 := !p1;
		var i1 := 0;
		while (v14)
			decreases 100 - i1
		{
			if (i1 >= 100) {
				break;
			}
			
			i1 := i1 + 1;
			v14 := p1;
			f18 := v0[|(v0 + v0)|];
			v14 := v14;
			var v15: set<int> := {fm4(p1, v0[p0], f18, v14, globalState)};
			v14 := v15 >= v15;
		}
		r0 := p0;
	}
}

method Main() {
	var v0 := 140;
	var v1: seq<int> := [-v0];
	var v2 := true;
	var v3: map<bool, bool> := map[!true := v2];
	var v4 := 'q';
	var v5: map<char, int> := map[v4 := -v0];
	var v7: array<map<int, int>> := new map<int, int>[10](i0 => map v6 : int | (0x18f <= v6) && (v6 < -703) :: (v6 % 668) := (v0));
	var v8: set<bool> := {v2};
	var v9: array<int> := new int[12] [v0, v0, v0, v0, v0, 44, v0, v0, v0, v0, v0, |v8|];
	var v10: array<bool> := new bool[5];
	var v11: map<array<int>, array<bool>> := map[v9 := v10];
	var v12 := "oalh";
	var v13: seq<array<int>> := [v9];
	var v14: seq<bool> := [false, v2];
	var v15: map<int, bool> := map[v0 := v2];
	var v16: map<int, bool> := map[v0 := if (|"qgoah"| in v15) then v15[|"qgoah"|] else v2];
	var v17: seq<map<int, bool>> := [v16];
	var globalState := new GlobalState(0x203, v1, v1, v3, 0x20b, v5, 914, v7, v11, v12, map[v10 := v2], [v9, v9, v9, v9, v9] + v13, false, v14, false, 682, [v16] + v17, true);
	for i1 := v0 to -(0x386 % v0) {
		v2 := (if (v2) then i1 else -i1) < (-v0 * v0);
		var v18: seq<string> := [v12, v12];
		v18 := [v12] + v18;
		var v19: set<array<int>> := {v9};
		v2, v2, globalState.f1, v0 := v2, v2, v1, |(v19 * v19)| / |v1|;
		v10[545] := v2;
		var v20: map<int, char> := map[fm0(globalState) := if (v2) then v4 else v4];
		var v21 := DC0(fm0(globalState));
		v4, globalState.f6, v10[545], v2, v0 := if (0x120 in v20) then v20[0x120] else v4, v21.cf0 - v0, v2, if (false) then false else v2, v21.cf0;
	}
	var i2 := 0;
	while (v2)
		decreases 100 - i2
	{
		if (i2 >= 100) {
			break;
		}
		
		i2 := i2 + 1;
		v3 := v3[!(0x388 <= v0) := v2];
		v4, globalState.f6 := v4, v0;
		var v22: map<array<bool>, bool> := map[v10 := true];
		globalState.f6 := if (v2) then v0 else --533 * |v22|;
		v9[307] := v0;
		v9[307] := -0x1d - v0;
	}
	v9[465] := v0;
	v9[465] := v0 % -|"nxertem"|;
	var v23, v24, v25 := m0(-v9[465], fm1(if (|"mjgtu"| in v15) then v15[|"mjgtu"|] else v2, v4, -0x105, globalState), |v12[v0 := v4]| * -v0, globalState);
	v12 := v12;
	var v26: map<bool, seq<bool>> := map[v24 := v14];
	v9[465], v14 := fm0(globalState), (if (v24 in v26) then v26[v24] else v14) + v14;
	v24 := v2;
	var v27: array<char> := new char[1](i3 => v4);
	var v28: map<int, array<char>> := map[v9[465] := v27];
	if (fm2(globalState) >= {fm1(v2, v4, |v28|, globalState), v24, if (v24 in v3) then v3[v24] else v24}) {
		var v29 := DC1();
		v29 := v29;
		var v30: map<bool, int> := map[v2 := 0x213];
		var v31 := new C3(fm20(false, globalState), map[v12 := |v30|]);
		globalState.f6 := |(map v32 : int | (0x248 <= v32) && (v32 < 0x14e) :: (v32 - v0) := (DC12(multiset{v2})))| * (v25 + v9[465]);
		v23[790] := !(if (v24) then v24 else v2);
		var v33: set<int> := {v25};
		var v34: seq<set<int>> := [v33];
		var v35: map<int, seq<set<int>>> := map[v0 := v34];
		v23[790], v2, v29, v2, v24 := !v24, v2, fm25(v3, DC23(if (v0 in v35) then v35[v0] else v34, v9[465], v2, v0), fm1(false, v4, v0, globalState), v9[465], globalState), v2 ==> (v25 < |fm26(globalState)|), !v24;
		v2 := v25 != v25;
	} else {
		var v36: array<string> := new string[1];
		var v37: map<array<string>, array<bool>> := map[v36 := v10];
		v37 := v37[v36 := v23];
		var v38: array<seq<bool>> := new seq<bool>[16](i4 => v14 + v14);
		v38 := v38;
		var v39: array<array<bool>> := new array<bool>[19];
		var v40 := DC16(v2, v39);
		match (v40) {
			case DC16(cf19, cf20) =>
				var v41 := DC18(v8);
				v9[465] := v25 - (|v41.cf22| * |map[cf19 := v25]|);
				cf19 := v24;
				var v42: multiset<int> := multiset{v9[465]};
				v10[284] := !(|v12| !in v42);
				v10[284] := v2;
				v0 := fm0(globalState);
			case DC15(cf18) =>
				var v43, v44, v45 := m0(|v12|, fm0(globalState) in (seq(0x9e, i5  => (629))), |v12|, globalState);
				var v46: C3 := new C3(v4, map[v12 := v0]);
				v46 := v46;
				var v47: multiset<bool> := multiset{v44, true};
				var v48 := DC10(multiset([v2]) < v47);
				v48 := v48;
				var v49: array<array<string>> := new array<string>[10];
				var v50: map<bool, array<string>> := map[v2 := v36];
				v49[151] := if (v44 in v50) then v50[v44] else v36;
				v49[151] := v36;
			case DC17(cf21) =>
				var v51: map<string, int> := map[v12 := |DC22(seq(0x3a4, i6  => (v4))).cf27|];
				var v52: C3 := new C3(v4, v51);
				v52 := v52;
				var v53: map<bool, array<int>> := map[v24 := if (v2) then v9 else v13[v0]];
				v9 := if (v52.fm3(v2, -v0, globalState) in v53) then v53[v52.fm3(v2, -v0, globalState)] else v9;
				var v54 := new C3(v4, v51 + v51);
				var v55, v56, v57, v58 := v52.m1(fm1(v2, 'm', |v53|, globalState), globalState);
		}
		
		var v59: map<string, int> := map[v12 := v0];
		var v60: C1 := new C1(v9[465], v24, v4, v59);
		var v61: map<C1, int> := map[v60 := -(v0 % v25)];
		v61 := v61[v60 := v25];
		var v62: multiset<bool> := multiset{v24, true};
		var v63: map<char, multiset<bool>> := map[v4 := multiset([v60.f23]) * v62];
		v63 := v63[v4 := v62];
	}
	
	var v64: array<D7> := new D7[18];
	forall i7 | 0 <= i7 < v64.Length {
		v64[i7] := DC20(v2, v2);
	}
	var v65: map<int, int> := map[v25 := 0x1ca];
	var i8 := 0;
	while ((-|v65| * |v12|) >= v25)
		decreases 100 - i8
	{
		if (i8 >= 100) {
			break;
		}
		
		i8 := i8 + 1;
		v25 := |(v1[0x68 := -v25])[v25 / v0 := 123]|;
		var v66: set<int> := {v9[465], v9[465]};
		var v67: seq<set<int>> := [v66, v66];
		v67 := if (v2) then [v66, v66, v66, v66, {v0}] else [set v68 : int | (-240 <= v68) && (v68 < 336) :: (v68 - |v15|), {0x5a}, v66, v66, v66];
		var v70: seq<string> := [seq(0xd8, i9  => (v4))];
		var v71 := new C2(v4, map v69 : string | v69 in v70 :: (v69) := (v9[465]));
		v65 := v65[-v9[465] := 378 / v25];
	}
	var v72: array<C0> := new C0[27];
	var v73: map<array<C0>, seq<int>> := map[v72 := v1];
	v73 := v73[v72 := v1];
	var v74 := DC6(v12 <= v12, v24);
	match (v74) {
		case DC4(cf3, cf4, cf5) =>
			var v75: map<string, int> := map[v12 := cf4];
			var v76: T0 := new C0("ikrmgew", map['o' := v9[465]], v4, v75);
			var v77: multiset<T0> := multiset{v76};
			var v78: seq<multiset<T0>> := [multiset{v76}, multiset{v76}, v77[v76 := fm0(globalState)], multiset{v76, v76}];
			v78 := v78;
			var v79: set<int> := {0x314, v25, 0x3d3, cf3, v9[465]};
			var v80: seq<set<int>> := [v79, v79];
			var v81, v82, v83 := m0(-DC23(v80, v25, v2, 0x115).cf31, v2, v0, globalState);
			if (v2) {
				v82 := !(cf5 <==> fm1(v24, v76.f18, v0, globalState));
				globalState.f6 := v0;
				var v86: seq<array<bool>> := [v10];
				var v87: array<array<bool>> := new array<bool>[29] [v81, v81, v23, v10, v81, v81, v10, v10, v81, v81, v81, v10, v23, v23, v23, v81, v81, v10, v10, v23, v23, v10, v23, v23, v10, v23, v23, v81, v86[|v8|]];
				var v88 := DC16(cf5, v87);
				var v89 := DC17(v88);
				var v90: map<D6, bool> := map[v89 := cf5];
				var v93: array<map<int, bool>> := new map<int, bool>[18] [map[-|v12| := v2], v15, v16, v16, v16, map[|v8| := fm1(true, v4, v83, globalState)], map[cf4 := true], map v84 : int | (919 <= v84) && (v84 < 0x23a) :: (v84 % v0) := (cf5), v16, v15, v16, (map v85 : int | (0x2bb <= v85) && (v85 < 764) :: (v85 + v25) := (cf5)) + v15, map[v0 := !(if (v89 in v90) then v90[v89] else v2)] + v16, map v91 : int | (0x3a <= v91) && (v91 < 0x1a8) :: (v91 - 335) := (true), v15, v16, map v92 : int | (0xe6 <= v92) && (v92 < 64) :: (v92 - v0) := (v24), v16];
				v93[567] := map[|v12| := !v24];
				v93[567] := if (fm1(v82, v76.f18, v9[465], globalState)) then v15 else v16;
				var v94 := DC25(v9);
				var v95: array<array<int>> := new array<int>[9] [v9, v9, v9, v94.cf33, v9, v9, v9, v13[v83], if (v24) then v9 else v9];
				v95[323] := v9;
				globalState.f1, v95[323], v9[465] := v1, v9, v25;
				v24 := v82;
			} else {
				var v96, v97, v98 := m0(cf3, v25 == |fm2(globalState)|, |v79|, globalState);
				v24 := false;
				var v99: T1 := new C3(v76.f18, map[v12 := v25]);
				var v100: C2 := new C2(v76.f18, v75);
				var v101: map<T1, C2> := map[v99 := v100];
				var v102: multiset<map<T1, C2>> := multiset{v101 + v101};
				var v103: seq<map<T1, C2>> := [v101];
				v102 := multiset(v103) - v102;
				v83 := v0;
				v0 := v98;
			}
			
			v4 := v4;
		case DC5(cf6) =>
			v8 := if (!v24 && v2) then if (v2) then fm2(globalState) else v8 else v8;
			var v104: T0 := new C0(v12, v5, v4, map["vf" := v9[465]]);
			v104 := v104;
			var v105: array<seq<int>> := new seq<int>[13];
			v105 := v105;
			var v106: multiset<bool> := multiset{if (false in v3) then v3[false] else v2, !v2};
			var v107 := DC4(-|v106|, -v25, v2);
			var v108: seq<set<bool>> := [v8];
			v9[465], v107, v108, globalState.f13 := v9[465], DC4(v0, v25, true), v108, fm27(globalState);
		case DC6(cf7, cf8) =>
			var v109, v110, v111 := m0(v0, cf7 && v2, v0, globalState);
			var v112: map<string, int> := map[v12 := -|map[v25 := cf7]|];
			var v113: map<int, map<string, int>> := map[-v25 := v112];
			var v114 := new C1(-(v25 / v111), true, v4, if (v0 in v113) then v113[v0] else v112);
			var v115 := v114.m2(!cf8, v114.f22, v9[465], v111, globalState);
			var v116 := new C0(seq(0x66, i10  => (v4)), v5, v4, v112);
		case DC3(cf2) =>
			var v117: map<string, int> := map[v12 := v9[465]];
			var v118: multiset<char> := multiset{v4};
			var v119 := new C0(v12 + v12, v5, v4, v117 + (map[v12 := |v118[v4 := v25]|])["rtoijma" := v9[465]]);
			var v120 := DC14(v24, v9[465]);
			var v121: map<map<D5, bool>, int> := map[map[v120 := v2] := v25];
			var v122: map<D5, bool> := map[v120 := v24];
			v121 := map[v122 := if (v9[465] in v65) then v65[v9[465]] else v25];
			v4, cf2 := v4, if (v14[v25] in v3) then v3[v14[v25]] else v24;
			var v123: array<string> := new string[12] [v12, seq(0x1a0, i11  => (v4)), "p", v12, fm8(globalState), v119.f20, if (!v24) then seq(0x31, i12  => (v4)) else "f", seq(0x1d2, i13  => (v4)), "l" + v119.f20, v12, "art", "fhfctskoe"];
			v123[463] := v12 + v12[v0 := v4];
			v123[463] := v12;
	}
	
	var i14 := 0;
	while (false)
		decreases 100 - i14
	{
		if (i14 >= 100) {
			break;
		}
		
		i14 := i14 + 1;
		var v124: C0 := new C0(v12, v5, v4, map[v12 := |{|v14|}|]);
		v124 := v124;
		var v125: map<string, int> := map[v12 := |v12|];
		var v126 := new C0(v12, v124.f21, if (v24) then v4 else v4, v125 + v125);
		var v127: array<seq<bool>> := new seq<bool>[22];
		v127[279] := [v24] + v14;
		v127[279] := v14;
		var v128: map<array<int>, map<string, int>> := map[v9 := v125];
		var v129 := new C3(v4, (if (v9 in v128) then v128[v9] else v125) + v125);
	}
	var v130: map<bool, int> := map[v24 := v0];
	for i15 := |v130| to fm0(globalState) {
		v2 := if (v0 >= i15) then !v2 else v24;
		var v131: map<int, map<int, int>> := map[i15 := map[if (v2 in v130) then v130[v2] else i15 := |v8|]];
		v131 := v131;
		v3 := v3;
		v9 := v9;
	}
	v25 := v0;
	if (v24) {
		var v132 := DC8(fm20(v24, globalState));
		var v134: seq<string> := [v12];
		var v135: C3 := new C3(v132.cf10, map v133 : string | v133 in v134 :: (v133) := (v25));
		var v136: map<C3, bool> := map[v135 := if (v2) then v2 else !v135.fm3(v24, |v12|, globalState)];
		var v137: seq<multiset<int>> := [fm28(v2, v9[465], v24, globalState)];
		v136 := v136[v135 := |v137[v25]| != v0];
		globalState.f6 := v9[465];
		var v138: map<string, int> := map[v12 := fm0(globalState)];
		var v139 := new C1(v25, v2, v4, v138);
		var v140: array<array<int>> := new array<int>[23];
		v140[399] := v9;
		v140[399] := v9;
		v12 := fm8(globalState);
	} else {
		globalState.f6 := v25;
		var v141: array<array<int>> := new array<int>[10] [v9, v9, v9, v9, v9, v9, v9, v9, v9, v9];
		v141[146] := v9;
		v141[146] := v9;
		var v142 := DC14(false, 0x33);
		v24, v0, v2 := (fm29(globalState)).cf5, (-v9[465] - |v12|) - v142.cf17, v9[465] < -v25;
		var v143: multiset<bool> := multiset{v24, v24, v2};
		var v144, v145, v146 := m0(|v143|, v24, -(v25 * |fm30(v2, globalState)|), globalState);
		v145 := v24;
	}
	
}