function safeIndex(x: int, length: int): int
{
                                              x % length
}
trait T0 {
 const f22 : array<array<int>>
}
class GlobalState {
 var f12 : bool
 constructor (f0 : bool, f1 : int, f2 : bool, f3 : map<array<int>, seq<int>>, f4 : char, f5 : set<int>, f6 : int, f7 : int, f8 : bool, f9 : seq<bool>, f10 : int, f11 : map<    bool ,     int >, f12 : bool, f13 : array<int>, f14 : string, f15 : bool, f16 : map<      bool , int>, f17 : bool, f18 : bool, f19 : char, f20 : int, f21 : bool) {
 }
}
class C0 extends T0 {
 var f23 : int
 constructor (f23 : int, f22 : array<array<int>>) {
  this.f23 := f23;
 }
}
method Main() {
 var v0             := new int[5];
 var v1 := 0x20a;
 var v2 := "p";
 var v4 := true;
 var v6           := [          ];
 var v7                            := map[v0 := v6];
 var v8           := {v1        , 0xa4, -0x20a};
 var v9            := [                  ];
 var v10                           := map[        ];
 var v11              := new bool[12];
 var v13                        := map[            ];
 var globalState := new GlobalState(false, 265, true, v7               , 'c', v8, 0x191, -0x2bb, false, v9, 0x2e9, v10, false, v0, "qftlrogoy", true,                 v13         , false, true, 'u', -381, true);
 var v56                    := new array<int>[23]                                                                                             ;
 var v57     := new C0(v1, v56);
 var v67     := new C0(|v8|, v57.f22);
 for i6 := v67.f23 to v1 {
  v2, v11[safeIndex(148, v11.Length)], globalState.f12, v0[safeIndex(814, v0.Length)] := v2                                                      , v4, v4,     v67.f23     ;
 }
 print v0[4]      ;
}
