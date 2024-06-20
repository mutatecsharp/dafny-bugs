import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_
import _dafny
import System_

# Module: module_

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def abs(x):
        if (x) < (0):
            return (-1) * (x)
        elif True:
            return x

    @staticmethod
    def safeIndex(x, length):
        if (x) < (0):
            return 0
        elif (x) >= (length):
            return _dafny.euclidian_modulus(x, length)
        elif True:
            return x

    @staticmethod
    def safeDivisionInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_division(x1, x2)

    @staticmethod
    def safeModuloInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_modulus(x1, x2)

    @staticmethod
    def fm0(p0, p1, p2, globalState):
        return -184

    @staticmethod
    def fm3(p0, globalState):
        source0_ = D2_DC7(False, False, True)
        if source0_.is_DC7:
            d_0___mcc_h0_ = source0_.cf11
            d_1___mcc_h1_ = source0_.cf12
            d_2___mcc_h2_ = source0_.cf13
            d_3_cf13_ = d_2___mcc_h2_
            d_4_cf12_ = d_1___mcc_h1_
            d_5_cf11_ = d_0___mcc_h0_
            return D1_DC4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljgfacm")))
        elif True:
            d_6___mcc_h3_ = source0_.cf10
            d_7_cf10_ = d_6___mcc_h3_
            if not(False):
                return D1_DC4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))
            elif True:
                return D1_DC4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dewv")))

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        return True

    @staticmethod
    def fm6(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-548, 187):
                d_8_v0_: int = compr_0_
                if ((-548) <= (d_8_v0_)) and ((d_8_v0_) < (187)):
                    coll0_[(d_8_v0_) * (953)] = -372
            return _dafny.Map(coll0_)
        return ((len(_dafny.Map({True: len(_dafny.Map({True: iife0_()
        }))}))) - (len(_dafny.Map({False: True})))) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: -430}))]))

    @staticmethod
    def fm7(globalState):
        return D0_DC2(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfekhaf"))), 774, -567, 402}), ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False])), (0) - (-165)])) | (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d')])), (_dafny.MultiSet([False])).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_9_i0_ in range(default__.abs(-3))])), 740]))).cardinality, False)

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([853])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ks")))]))) + ((_dafny.SeqWithoutIsStrInference([121, 889, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xikkpo"))), len(_dafny.SeqWithoutIsStrInference([258])), (_dafny.MultiSet([not(not(True)), True, True, not(True)])).cardinality])) + (_dafny.SeqWithoutIsStrInference([379])))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        return _dafny.Set({(_dafny.SeqWithoutIsStrInference([832, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_10_i0_ in range(default__.abs(242))])), (0) - (-61)])) < (_dafny.SeqWithoutIsStrInference([258 for d_11_i1_ in range(default__.abs(510))]))})

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False, not(True)])) + (_dafny.SeqWithoutIsStrInference([(D0_DC1(_dafny.Set({_dafny.Set({True, not(True)})}), True)).cf3]))) + (_dafny.SeqWithoutIsStrInference([(D0_DC2(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajxt")))), -490}), 836, True)).cf6]))

    @staticmethod
    def fm11(p0, globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('o'), _dafny.CodePoint('m'), _dafny.CodePoint('f')])) - (_dafny.MultiSet([_dafny.CodePoint('f')]))).intersection(((D7_DC18(_dafny.MultiSet([_dafny.CodePoint('c')]))).cf26).intersection(_dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('y'), _dafny.CodePoint('j')])))

    @staticmethod
    def fm12(p0, p1, globalState):
        if not(False):
            return D4_DC13(D4_DC12(60, _dafny.CodePoint('e'), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_12_i0_ in range(default__.abs(115))]))])))
        elif True:
            return D4_DC13(D4_DC12(len(_dafny.Map({-739: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([380])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bntttie")))]))})), _dafny.CodePoint('g'), _dafny.MultiSet([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gix")))])).cardinality])))

    @staticmethod
    def fm13(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kem"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqlomrr")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxb"))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dugneonhq"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awqmcmgf"))])))

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return _dafny.Set({_dafny.CodePoint('d'), _dafny.CodePoint('y'), _dafny.CodePoint('y'), (_dafny.CodePoint('d') if False else _dafny.CodePoint('g')), _dafny.CodePoint('j')})

    @staticmethod
    def fm15(p0, p1, globalState):
        return _dafny.CodePoint('b')

    @staticmethod
    def fm16(p0, p1, globalState):
        return (_dafny.MultiSet([len(_dafny.Set({not(False), False}))])) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')]))])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({951})), 849]))))

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.Map({897: False})).keys.Elements:
                d_13_v0_: int = compr_1_
                if (d_13_v0_) in (_dafny.Map({897: False})):
                    coll1_[(d_13_v0_) + (len(_dafny.SeqWithoutIsStrInference([303 for d_14_i0_ in range(default__.abs(216))])))] = _dafny.CodePoint('t')
            return _dafny.Map(coll1_)
        return ((D8_DC20(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, False])), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aapfmbeu"))))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lytyvmod")))])]))).cf28) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdnvcfxx")))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(iife1_()
        )]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwpfdjks")))]))]), _dafny.SeqWithoutIsStrInference([764]), (D9_DC24(_dafny.SeqWithoutIsStrInference([369, -732]))).cf36, _dafny.SeqWithoutIsStrInference([961, len(_dafny.Set({False})), 812, 50])])))

    @staticmethod
    def fm18(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrki")))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        r3: int = int(0)
        r3 = p0
        r2 = p0
        d_15_v0_: bool
        d_15_v0_ = True
        d_16_i0_: int
        d_16_i0_ = 0
        with _dafny.label("0"):
            while d_15_v0_:
                with _dafny.c_label("0"):
                    if (d_16_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_16_i0_ = (d_16_i0_) + (1)
                    if d_15_v0_:
                        r3 = (p0) - ((p0) * (p0))
                        d_17_v1_: _dafny.Seq
                        d_17_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ym"))
                        d_18_v2_: _dafny.Set
                        d_18_v2_ = _dafny.Set({len(d_17_v1_)})
                        d_19_v3_: _dafny.Map
                        d_19_v3_ = _dafny.Map({(d_15_v0_) == (d_15_v0_): d_18_v2_})
                        def iife2_():
                            coll2_ = _dafny.Set()
                            compr_2_: int
                            for compr_2_ in _dafny.IntegerRange(-838, 345):
                                d_20_v4_: int = compr_2_
                                if ((-838) <= (d_20_v4_)) and ((d_20_v4_) < (345)):
                                    coll2_ = coll2_.union(_dafny.Set([(d_20_v4_) + (p0)]))
                            return _dafny.Set(coll2_)
                        d_19_v3_ = (d_19_v3_) | (_dafny.Map({not(d_15_v0_): iife2_()
                        }))
                        d_21_v5_: _dafny.Array
                        nw0_ = _dafny.Array(D4.default()(), 5)
                        d_21_v5_ = nw0_
                        d_22_v6_: str
                        d_22_v6_ = _dafny.CodePoint('t')
                        d_23_v7_: _dafny.Map
                        d_23_v7_ = _dafny.Map({10: -844})
                        d_24_v8_: _dafny.MultiSet
                        d_24_v8_ = _dafny.MultiSet([p0])
                        d_25_v9_: D4
                        d_25_v9_ = D4_DC12(p0, default__.fm15(d_22_v6_, d_23_v7_, globalState), d_24_v8_)
                        index0_ = default__.safeIndex(163, (d_21_v5_).length(0))
                        (d_21_v5_)[index0_] = d_25_v9_
                        index1_ = default__.safeIndex(163, (d_21_v5_).length(0))
                        (d_21_v5_)[index1_] = d_25_v9_
                        d_26_v10_: _dafny.Map
                        d_26_v10_ = _dafny.Map({768: d_17_v1_})
                        d_27_v11_: D0
                        d_27_v11_ = D0_DC3(d_22_v6_)
                        d_26_v10_ = (d_26_v10_).set(p0, default__.fm18((d_27_v11_).cf7, globalState))
                        d_28_v12_: C1
                        nw1_ = C1()
                        nw1_.ctor__((d_15_v0_ if d_15_v0_ else d_15_v0_), -333)
                        d_28_v12_ = nw1_
                        nw2_ = C1()
                        nw2_.ctor__((p0) <= (667), p0)
                        d_28_v12_ = nw2_
                    elif True:
                        d_29_v13_: str
                        d_29_v13_ = _dafny.CodePoint('u')
                        (globalState).f4 = d_29_v13_
                        d_30_v14_: _dafny.Array
                        nw3_ = _dafny.Array(_dafny.Array(None, 0), 7)
                        d_30_v14_ = nw3_
                        d_31_v15_: _dafny.Map
                        d_31_v15_ = _dafny.Map({d_15_v0_: d_30_v14_})
                        d_31_v15_ = (d_31_v15_).set((d_15_v0_) or (d_15_v0_), d_30_v14_)
                        r2 = p0
                        d_32_v16_: _dafny.Array
                        nw4_ = _dafny.Array(False, 29)
                        d_32_v16_ = nw4_
                        d_32_v16_ = d_32_v16_
                        d_33_v17_: _dafny.Seq
                        d_33_v17_ = _dafny.SeqWithoutIsStrInference([p0])
                        r0 = (p0) in (d_33_v17_)
                    d_34_v18_: _dafny.Array
                    nw5_ = _dafny.Array(int(0), 8)
                    d_34_v18_ = nw5_
                    d_35_v19_: D0
                    d_35_v19_ = D0_DC0(p0, d_15_v0_)
                    d_36_v20_: _dafny.MultiSet
                    d_36_v20_ = _dafny.MultiSet([(d_35_v19_).cf0, p0])
                    d_37_v21_: _dafny.Map
                    d_37_v21_ = _dafny.Map({(d_36_v20_).cardinality: d_34_v18_})
                    d_38_v22_: _dafny.Seq
                    d_38_v22_ = _dafny.SeqWithoutIsStrInference([d_34_v18_])
                    d_39_v23_: _dafny.Map
                    d_39_v23_ = _dafny.Map({d_15_v0_: d_34_v18_})
                    d_40_v24_: _dafny.Array
                    nw6_ = _dafny.Array(None, 24)
                    nw6_[int(0)] = d_34_v18_
                    nw6_[int(1)] = d_34_v18_
                    nw6_[int(2)] = ((d_37_v21_)[p0] if (p0) in (d_37_v21_) else d_34_v18_)
                    nw6_[int(3)] = d_34_v18_
                    nw6_[int(4)] = d_34_v18_
                    nw6_[int(5)] = d_34_v18_
                    nw6_[int(6)] = (d_38_v22_)[default__.safeIndex(p0, len(d_38_v22_))]
                    nw6_[int(7)] = d_34_v18_
                    nw6_[int(8)] = d_34_v18_
                    nw6_[int(9)] = d_34_v18_
                    nw6_[int(10)] = d_34_v18_
                    nw6_[int(11)] = d_34_v18_
                    nw6_[int(12)] = d_34_v18_
                    nw6_[int(13)] = ((d_39_v23_)[d_15_v0_] if (d_15_v0_) in (d_39_v23_) else d_34_v18_)
                    nw6_[int(14)] = d_34_v18_
                    nw6_[int(15)] = ((d_37_v21_)[p0] if (p0) in (d_37_v21_) else d_34_v18_)
                    nw6_[int(16)] = d_34_v18_
                    nw6_[int(17)] = d_34_v18_
                    nw6_[int(18)] = d_34_v18_
                    nw6_[int(19)] = d_34_v18_
                    nw6_[int(20)] = d_34_v18_
                    nw6_[int(21)] = d_34_v18_
                    nw6_[int(22)] = d_34_v18_
                    nw6_[int(23)] = d_34_v18_
                    d_40_v24_ = nw6_
                    d_41_v25_: D5
                    d_41_v25_ = D5_DC14(d_40_v24_)
                    d_41_v25_ = d_41_v25_
                    if False:
                        d_42_v26_: _dafny.Seq
                        d_42_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esfxlsqkn"))
                        r3 = default__.safeModuloInt(len(d_42_v26_), p0)
                        d_43_v27_: _dafny.Array
                        def lambda0_(d_44_v0_):
                            def lambda1_(d_45_i1_):
                                return d_44_v0_

                            return lambda1_

                        init0_ = lambda0_(d_15_v0_)
                        nw7_ = _dafny.Array(None, 12)
                        for i0_0_ in range(nw7_.length(0)):
                            nw7_[i0_0_] = init0_(i0_0_)
                        d_43_v27_ = nw7_
                        index2_ = default__.safeIndex(92, (d_43_v27_).length(0))
                        (d_43_v27_)[index2_] = (d_15_v0_) and (default__.fm6((p1)[default__.safeIndex(p0, len(p1))], p0, globalState))
                        index3_ = default__.safeIndex(92, (d_43_v27_).length(0))
                        (d_43_v27_)[index3_] = (False) or (False)
                        r3 = p0
                        d_46_v28_: str
                        d_46_v28_ = _dafny.CodePoint('b')
                        d_47_v29_: _dafny.Map
                        d_47_v29_ = _dafny.Map({True: d_46_v28_})
                        d_48_v30_: C0
                        nw8_ = C0()
                        nw8_.ctor__((d_47_v29_) == (d_47_v29_))
                        d_48_v30_ = nw8_
                        d_49_v31_: D0
                        d_49_v31_ = D0_DC3(d_46_v28_)
                        d_50_v32_: _dafny.Array
                        nw9_ = _dafny.Array(None, 16)
                        nw9_[int(0)] = D0_DC3(d_46_v28_)
                        nw9_[int(1)] = D0_DC3(d_46_v28_)
                        nw9_[int(2)] = d_49_v31_
                        nw9_[int(3)] = D0_DC3(d_46_v28_)
                        nw9_[int(4)] = d_49_v31_
                        nw9_[int(5)] = D0_DC3(d_46_v28_)
                        nw9_[int(6)] = D0_DC3(_dafny.CodePoint('c'))
                        nw9_[int(7)] = d_49_v31_
                        nw9_[int(8)] = d_49_v31_
                        nw9_[int(9)] = d_49_v31_
                        nw9_[int(10)] = d_49_v31_
                        nw9_[int(11)] = d_49_v31_
                        nw9_[int(12)] = d_49_v31_
                        nw9_[int(13)] = D0_DC3(d_46_v28_)
                        nw9_[int(14)] = d_49_v31_
                        nw9_[int(15)] = d_49_v31_
                        d_50_v32_ = nw9_
                        d_51_v33_: _dafny.Array
                        nw10_ = _dafny.Array(None, 9)
                        nw10_[int(0)] = d_50_v32_
                        nw10_[int(1)] = d_50_v32_
                        nw10_[int(2)] = d_50_v32_
                        nw10_[int(3)] = d_50_v32_
                        nw10_[int(4)] = d_50_v32_
                        nw10_[int(5)] = d_50_v32_
                        nw10_[int(6)] = d_50_v32_
                        nw10_[int(7)] = d_50_v32_
                        nw10_[int(8)] = d_50_v32_
                        d_51_v33_ = nw10_
                        index4_ = default__.safeIndex(929, (d_51_v33_).length(0))
                        (d_51_v33_)[index4_] = d_50_v32_
                        index5_ = default__.safeIndex(929, (d_51_v33_).length(0))
                        rhs0_ = d_48_v30_
                        rhs1_ = d_50_v32_
                        lhs0_ = d_51_v33_
                        lhs1_ = default__.safeIndex(929, (d_51_v33_).length(0))
                        d_48_v30_ = rhs0_
                        lhs0_[lhs1_] = rhs1_
                        index6_ = default__.safeIndex(92, (d_43_v27_).length(0))
                        (d_43_v27_)[index6_] = not((d_43_v27_)[default__.safeIndex(92, (d_43_v27_).length(0))])
                    elif True:
                        r0 = d_15_v0_
                        index7_ = default__.safeIndex(438, (d_34_v18_).length(0))
                        (d_34_v18_)[index7_] = p0
                        index8_ = default__.safeIndex(438, (d_34_v18_).length(0))
                        (d_34_v18_)[index8_] = p0
                        d_52_v34_: _dafny.Seq
                        d_52_v34_ = _dafny.SeqWithoutIsStrInference([479, default__.fm0(_dafny.Map({True: d_15_v0_}), d_15_v0_, d_15_v0_, globalState)])
                        d_53_v35_: _dafny.Set
                        d_53_v35_ = _dafny.Set({d_15_v0_, d_15_v0_})
                        d_54_v36_: _dafny.Map
                        d_54_v36_ = _dafny.Map({d_15_v0_: default__.fm4(d_15_v0_, d_53_v35_, p0, globalState)})
                        d_55_v37_: _dafny.Seq
                        d_55_v37_ = _dafny.SeqWithoutIsStrInference([(d_34_v18_)[default__.safeIndex(438, (d_34_v18_).length(0))], (len(d_52_v34_)) - (default__.fm0(d_54_v36_, d_15_v0_, d_15_v0_, globalState)), default__.safeModuloInt((d_34_v18_)[default__.safeIndex(438, (d_34_v18_).length(0))], p0)])
                        d_55_v37_ = (d_55_v37_) + (d_52_v34_)
                        d_56_v38_: str
                        d_56_v38_ = _dafny.CodePoint('g')
                        d_57_v39_: _dafny.Seq
                        d_57_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sir"))
                        d_58_v40_: _dafny.Set
                        d_58_v40_ = _dafny.Set({d_36_v20_, d_36_v20_})
                        d_59_v41_: _dafny.Array
                        nw11_ = _dafny.Array(None, 17)
                        nw11_[int(0)] = d_15_v0_
                        nw11_[int(1)] = d_15_v0_
                        nw11_[int(2)] = d_15_v0_
                        nw11_[int(3)] = not(not((d_56_v38_) not in (d_57_v39_)))
                        nw11_[int(4)] = d_15_v0_
                        nw11_[int(5)] = d_15_v0_
                        nw11_[int(6)] = (default__.fm0(d_54_v36_, d_15_v0_, d_15_v0_, globalState)) > (p0)
                        nw11_[int(7)] = d_15_v0_
                        nw11_[int(8)] = (p1)[default__.safeIndex(p0, len(p1))]
                        nw11_[int(9)] = ((d_34_v18_)[default__.safeIndex(438, (d_34_v18_).length(0))]) >= ((d_34_v18_)[default__.safeIndex(438, (d_34_v18_).length(0))])
                        nw11_[int(10)] = d_15_v0_
                        nw11_[int(11)] = d_15_v0_
                        nw11_[int(12)] = (p0) <= (77)
                        nw11_[int(13)] = (d_36_v20_) in (d_58_v40_)
                        nw11_[int(14)] = d_15_v0_
                        nw11_[int(15)] = (d_15_v0_) or (not((p1)[default__.safeIndex(p0, len(p1))]))
                        nw11_[int(16)] = not(d_15_v0_)
                        d_59_v41_ = nw11_
                        index9_ = default__.safeIndex(438, (d_34_v18_).length(0))
                        rhs2_ = p0
                        rhs3_ = default__.fm0(d_54_v36_, not(not(d_15_v0_)), d_15_v0_, globalState)
                        rhs4_ = d_59_v41_
                        rhs5_ = (default__.fm6(d_15_v0_, (d_34_v18_)[default__.safeIndex(438, (d_34_v18_).length(0))], globalState) if (d_15_v0_ if d_15_v0_ else (p1)[default__.safeIndex(696, len(p1))]) else ((d_34_v18_)[default__.safeIndex(438, (d_34_v18_).length(0))]) > ((d_34_v18_)[default__.safeIndex(438, (d_34_v18_).length(0))]))
                        lhs2_ = d_34_v18_
                        lhs3_ = default__.safeIndex(438, (d_34_v18_).length(0))
                        lhs2_[lhs3_] = rhs2_
                        r3 = rhs3_
                        d_59_v41_ = rhs4_
                        d_15_v0_ = rhs5_
                        index10_ = default__.safeIndex(438, (d_34_v18_).length(0))
                        (d_34_v18_)[index10_] = (d_34_v18_)[default__.safeIndex(438, (d_34_v18_).length(0))]
                    d_60_v42_: str
                    d_60_v42_ = _dafny.CodePoint('x')
                    d_61_v43_: _dafny.Map
                    d_61_v43_ = _dafny.Map({p0: p0})
                    d_62_v44_: _dafny.Seq
                    d_62_v44_ = _dafny.SeqWithoutIsStrInference([default__.fm15(d_60_v42_, d_61_v43_, globalState)])
                    d_63_v45_: _dafny.Map
                    d_63_v45_ = _dafny.Map({d_15_v0_: d_62_v44_})
                    d_64_v46_: _dafny.Set
                    d_64_v46_ = _dafny.Set({(False) not in (d_63_v45_), (p0) < (p0), d_15_v0_})
                    d_64_v46_ = ((d_64_v46_) | (d_64_v46_)) | (d_64_v46_)
                    pass
            pass
        d_65_v47_: _dafny.Array
        nw12_ = _dafny.Array(None, 3)
        d_65_v47_ = nw12_
        d_65_v47_ = d_65_v47_
        d_66_v48_: _dafny.Seq
        d_66_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mns"))
        d_67_v49_: D1
        d_67_v49_ = D1_DC4(d_66_v48_)
        source1_ = d_67_v49_
        if source1_.is_DC5:
            d_68___mcc_h0_ = source1_.cf9
            d_69_cf9_ = d_68___mcc_h0_
            d_70_v50_: _dafny.Set
            d_70_v50_ = _dafny.Set({d_69_cf9_})
            d_71_v51_: _dafny.Map
            d_71_v51_ = _dafny.Map({(0) - (p0): d_69_cf9_})
            d_72_v52_: _dafny.Set
            d_72_v52_ = _dafny.Set({default__.fm4(d_15_v0_, d_70_v50_, (_dafny.MultiSet([len(d_71_v51_), p0, p0, p0, p0])).cardinality, globalState)})
            if ((d_72_v52_).issubset(d_70_v50_) if d_15_v0_ else d_15_v0_):
                d_73_v53_: _dafny.MultiSet
                d_73_v53_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_15_v0_, d_69_cf9_]), p1, _dafny.SeqWithoutIsStrInference([d_15_v0_, d_69_cf9_]), _dafny.SeqWithoutIsStrInference([d_69_cf9_, d_69_cf9_, d_15_v0_, d_69_cf9_, not(d_69_cf9_)])])
                d_74_v54_: _dafny.MultiSet
                d_74_v54_ = _dafny.MultiSet([d_73_v53_, _dafny.MultiSet([p1])])
                d_75_v55_: _dafny.MultiSet
                d_75_v55_ = _dafny.MultiSet([True])
                d_76_v56_: _dafny.MultiSet
                d_76_v56_ = _dafny.MultiSet([(0) - (p0), ((d_74_v54_)[d_73_v53_] if (d_73_v53_) in (d_74_v54_) else ((d_75_v55_)[d_15_v0_] if (d_15_v0_) in (d_75_v55_) else p0)), p0])
                d_76_v56_ = d_76_v56_
                d_77_v57_: str
                d_77_v57_ = _dafny.CodePoint('c')
                d_78_v58_: _dafny.Map
                d_78_v58_ = _dafny.Map({(d_66_v48_).set(default__.safeIndex(p0, len(d_66_v48_)), d_77_v57_): p0})
                d_79_v59_: _dafny.Map
                d_79_v59_ = _dafny.Map({(default__.fm18(d_77_v57_, globalState)) + (d_66_v48_): len(d_78_v58_)})
                d_80_v60_: _dafny.Set
                d_80_v60_ = _dafny.Set({p0})
                d_81_v61_: _dafny.Set
                d_81_v61_ = _dafny.Set({p0, (_dafny.MultiSet([d_80_v60_, _dafny.Set({p0})])).cardinality})
                d_79_v59_ = (d_79_v59_).set((d_66_v48_ if d_69_cf9_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmdwy"))), len(d_80_v60_))
                d_82_v62_: _dafny.Array
                nw13_ = _dafny.Array(int(0), 24)
                d_82_v62_ = nw13_
                index11_ = default__.safeIndex(948, (d_82_v62_).length(0))
                (d_82_v62_)[index11_] = p0
                index12_ = default__.safeIndex(948, (d_82_v62_).length(0))
                (d_82_v62_)[index12_] = (p0) - (default__.safeDivisionInt(len(d_66_v48_), (0) - ((d_75_v55_).cardinality)))
                r2 = p0
                d_83_v63_: C0
                nw14_ = C0()
                nw14_.ctor__(d_15_v0_)
                d_83_v63_ = nw14_
                d_84_v64_: _dafny.Map
                d_84_v64_ = _dafny.Map({(0) - (p0): d_83_v63_})
                d_84_v64_ = (d_84_v64_).set((d_82_v62_)[default__.safeIndex(948, (d_82_v62_).length(0))], d_83_v63_)
            elif True:
                d_85_v65_: _dafny.Array
                def lambda2_(d_86_i2_):
                    return (d_86_i2_) + (310)

                init1_ = lambda2_
                nw15_ = _dafny.Array(None, 5)
                for i0_1_ in range(nw15_.length(0)):
                    nw15_[i0_1_] = init1_(i0_1_)
                d_85_v65_ = nw15_
                index13_ = default__.safeIndex(910, (d_85_v65_).length(0))
                (d_85_v65_)[index13_] = p0
                d_87_v66_: _dafny.Map
                d_87_v66_ = _dafny.Map({d_15_v0_: d_69_cf9_})
                d_88_v67_: _dafny.Map
                d_88_v67_ = _dafny.Map({d_66_v48_: d_87_v66_})
                index14_ = default__.safeIndex(910, (d_85_v65_).length(0))
                rhs6_ = default__.fm6(d_15_v0_, (0) - (default__.fm0(((d_88_v67_)[d_66_v48_] if (d_66_v48_) in (d_88_v67_) else d_87_v66_), (p1)[default__.safeIndex(len(d_66_v48_), len(p1))], d_15_v0_, globalState)), globalState)
                rhs7_ = (0) - (p0)
                rhs8_ = not(d_69_cf9_)
                rhs9_ = ((p0) + (default__.fm0(_dafny.Map({not(d_15_v0_): not(True)}), ((d_87_v66_)[d_15_v0_] if (d_15_v0_) in (d_87_v66_) else d_69_cf9_), d_15_v0_, globalState))) - (p0)
                lhs4_ = d_85_v65_
                lhs5_ = default__.safeIndex(910, (d_85_v65_).length(0))
                r0 = rhs6_
                lhs4_[lhs5_] = rhs7_
                r0 = rhs8_
                r2 = rhs9_
                d_89_v68_: T0
                nw16_ = C0()
                nw16_.ctor__(d_15_v0_)
                d_89_v68_ = nw16_
                rhs10_ = default__.safeModuloInt(p0, -220)
                rhs11_ = d_15_v0_
                rhs12_ = d_89_v68_
                r2 = rhs10_
                d_15_v0_ = rhs11_
                d_89_v68_ = rhs12_
                d_90_v69_: _dafny.Array
                def lambda3_(d_91_v0_):
                    def lambda4_(d_92_i3_):
                        return d_91_v0_

                    return lambda4_

                init2_ = lambda3_(d_15_v0_)
                nw17_ = _dafny.Array(None, 7)
                for i0_2_ in range(nw17_.length(0)):
                    nw17_[i0_2_] = init2_(i0_2_)
                d_90_v69_ = nw17_
                index15_ = default__.safeIndex(515, (d_90_v69_).length(0))
                (d_90_v69_)[index15_] = d_69_cf9_
                index16_ = default__.safeIndex(515, (d_90_v69_).length(0))
                rhs13_ = d_69_cf9_
                rhs14_ = d_15_v0_
                lhs6_ = d_90_v69_
                lhs7_ = default__.safeIndex(515, (d_90_v69_).length(0))
                d_69_cf9_ = rhs13_
                lhs6_[lhs7_] = rhs14_
                r3 = (p0) - (default__.safeDivisionInt(p0, p0))
                index17_ = default__.safeIndex(515, (d_90_v69_).length(0))
                (d_90_v69_)[index17_] = not(d_69_cf9_)
            d_93_v70_: _dafny.Array
            nw18_ = _dafny.Array(False, 26)
            d_93_v70_ = nw18_
            index18_ = default__.safeIndex(116, (d_93_v70_).length(0))
            (d_93_v70_)[index18_] = d_69_cf9_
            index19_ = default__.safeIndex(116, (d_93_v70_).length(0))
            (d_93_v70_)[index19_] = (p1)[default__.safeIndex(default__.safeModuloInt(609, len(d_66_v48_)), len(p1))]
            d_94_v71_: C1
            nw19_ = C1()
            nw19_.ctor__((d_93_v70_)[default__.safeIndex(116, (d_93_v70_).length(0))], p0)
            d_94_v71_ = nw19_
            d_95_v72_: _dafny.Map
            d_95_v72_ = _dafny.Map({d_94_v71_: p0})
            d_96_v73_: _dafny.Map
            d_96_v73_ = _dafny.Map({d_69_cf9_: len(d_95_v72_)})
            d_97_v74_: _dafny.Seq
            d_97_v74_ = _dafny.SeqWithoutIsStrInference([(d_94_v71_).f15])
            d_98_v75_: C1
            nw20_ = C1()
            nw20_.ctor__(default__.fm6(d_15_v0_, len(d_96_v73_), globalState), len(d_97_v74_))
            d_98_v75_ = nw20_
            d_99_v76_: _dafny.Set
            d_99_v76_ = _dafny.Set({d_93_v70_})
            d_100_v77_: _dafny.Map
            d_100_v77_ = _dafny.Map({d_99_v76_: d_96_v73_})
            d_100_v77_ = (d_100_v77_).set(d_99_v76_, d_96_v73_)
        elif True:
            d_101___mcc_h1_ = source1_.cf8
            d_102_cf8_ = d_101___mcc_h1_
            d_103_v78_: _dafny.Seq
            d_103_v78_ = _dafny.SeqWithoutIsStrInference([p0])
            d_104_v79_: _dafny.Map
            d_104_v79_ = _dafny.Map({p0: d_103_v78_})
            d_105_v80_: _dafny.Map
            d_105_v80_ = _dafny.Map({d_15_v0_: d_15_v0_})
            d_103_v78_ = ((((d_104_v79_)[p0] if (p0) in (d_104_v79_) else d_103_v78_)).set(default__.safeIndex(default__.fm0(d_105_v80_, d_15_v0_, d_15_v0_, globalState), len(((d_104_v79_)[p0] if (p0) in (d_104_v79_) else d_103_v78_))), p0)) + ((d_103_v78_) + (d_103_v78_))
            d_106_v81_: C1
            nw21_ = C1()
            nw21_.ctor__((p0) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm")))), p0)
            d_106_v81_ = nw21_
            d_107_v82_: _dafny.Map
            d_107_v82_ = _dafny.Map({(d_106_v81_).f15: d_15_v0_})
            d_108_v83_: _dafny.Map
            d_108_v83_ = _dafny.Map({(d_106_v81_).f15: d_107_v82_})
            d_109_v84_: _dafny.Map
            d_109_v84_ = _dafny.Map({((d_106_v81_).f15) - (len(((d_108_v83_)[p0] if (p0) in (d_108_v83_) else d_107_v82_))): (d_106_v81_).f14})
            d_107_v82_ = (d_107_v82_).set(p0, d_15_v0_)
            d_110_v85_: str
            d_110_v85_ = _dafny.CodePoint('y')
            d_111_v86_: T0
            nw22_ = C0()
            nw22_.ctor__((d_106_v81_).f14)
            d_111_v86_ = nw22_
            d_112_v87_: _dafny.Map
            d_112_v87_ = _dafny.Map({d_15_v0_: d_110_v85_})
            d_113_v88_: _dafny.Map
            d_113_v88_ = _dafny.Map({d_111_v86_: d_112_v87_})
            d_114_v89_: _dafny.MultiSet
            d_114_v89_ = _dafny.MultiSet([(0) - (len(((d_113_v88_)[d_111_v86_] if (d_111_v86_) in (d_113_v88_) else d_112_v87_)))])
            d_115_v90_: D4
            d_115_v90_ = D4_DC12((d_106_v81_).f15, d_110_v85_, d_114_v89_)
            r3 = (((d_115_v90_).cf20).cardinality) + ((0) - ((0) - ((d_103_v78_)[default__.safeIndex((d_106_v81_).f15, len(d_103_v78_))])))
        r0 = False
        d_116_v91_: str
        d_116_v91_ = _dafny.CodePoint('l')
        d_117_v92_: _dafny.MultiSet
        d_117_v92_ = _dafny.MultiSet([len(p1)])
        d_118_v93_: D4
        d_118_v93_ = D4_DC12(p0, d_116_v91_, d_117_v92_)
        pat_let_tv0_ = d_15_v0_
        pat_let_tv1_ = p1
        pat_let_tv2_ = p1
        pat_let_tv3_ = d_15_v0_
        pat_let_tv4_ = d_15_v0_
        def lambda5_(source2_):
            if source2_.is_DC12:
                d_119___mcc_h2_ = source2_.cf18
                d_120___mcc_h3_ = source2_.cf19
                d_121___mcc_h4_ = source2_.cf20
                d_122_cf20_ = d_121___mcc_h4_
                d_123_cf19_ = d_120___mcc_h3_
                d_124_cf18_ = d_119___mcc_h2_
                return not (pat_let_tv0_) or (not((pat_let_tv1_)[default__.safeIndex(d_124_cf18_, len(pat_let_tv2_))]))
            elif source2_.is_DC11:
                d_125___mcc_h5_ = source2_.cf17
                d_126_cf17_ = d_125___mcc_h5_
                return pat_let_tv3_
            elif True:
                d_127___mcc_h6_ = source2_.cf21
                d_128_cf21_ = d_127___mcc_h6_
                return pat_let_tv4_

        r0 = not(lambda5_(d_118_v93_))
        d_129_v94_: _dafny.Seq
        d_129_v94_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0, p0])
        d_130_v95_: D0
        d_130_v95_ = D0_DC0(p0, d_15_v0_)
        pat_let_tv5_ = p0
        d_131_v96_: _dafny.Array
        nw23_ = _dafny.Array(None, 14)
        nw23_[int(0)] = d_130_v95_
        nw23_[int(1)] = D0_DC0(p0, d_15_v0_)
        nw23_[int(2)] = D0_DC0(p0, default__.fm6(d_15_v0_, p0, globalState))
        nw23_[int(3)] = d_130_v95_
        nw23_[int(4)] = d_130_v95_
        nw23_[int(5)] = d_130_v95_
        nw23_[int(6)] = D0_DC0(p0, default__.fm6(d_15_v0_, p0, globalState))
        nw23_[int(7)] = d_130_v95_
        nw23_[int(8)] = d_130_v95_
        def iife3_(_pat_let0_0):
            def iife4_(d_132_dt__update__tmp_h0_):
                def iife5_(_pat_let1_0):
                    def iife6_(d_133_dt__update_hcf0_h0_):
                        return D0_DC0(d_133_dt__update_hcf0_h0_, (d_132_dt__update__tmp_h0_).cf1)
                    return iife6_(_pat_let1_0)
                return iife5_(pat_let_tv5_)
            return iife4_(_pat_let0_0)
        nw23_[int(9)] = iife3_(d_130_v95_)
        nw23_[int(10)] = d_130_v95_
        nw23_[int(11)] = d_130_v95_
        nw23_[int(12)] = D0_DC0(p0, d_15_v0_)
        nw23_[int(13)] = d_130_v95_
        d_131_v96_ = nw23_
        r1 = ((d_131_v96_ if d_15_v0_ else d_131_v96_) if (default__.fm4(d_15_v0_, _dafny.Set({d_15_v0_, d_15_v0_}), len(d_129_v94_), globalState) if not(d_15_v0_) else d_15_v0_) else d_131_v96_)
        r2 = ((p0) - ((d_117_v92_).cardinality) if d_15_v0_ else p0)
        r3 = p0
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_134_v0_: int
        d_134_v0_ = -250
        d_135_v1_: _dafny.MultiSet
        d_135_v1_ = _dafny.MultiSet([_dafny.MultiSet([d_134_v0_])])
        d_136_v2_: _dafny.Array
        nw24_ = _dafny.Array(_dafny.CodePoint('D'), 29)
        d_136_v2_ = nw24_
        d_137_v3_: _dafny.Set
        d_137_v3_ = _dafny.Set({d_136_v2_})
        d_138_v4_: _dafny.Seq
        d_138_v4_ = _dafny.SeqWithoutIsStrInference([d_134_v0_])
        d_139_v5_: _dafny.Set
        d_139_v5_ = _dafny.Set({d_138_v4_})
        d_140_v6_: bool
        d_140_v6_ = False
        d_141_v7_: _dafny.MultiSet
        d_141_v7_ = _dafny.MultiSet([d_140_v6_, d_140_v6_])
        d_142_v8_: _dafny.MultiSet
        d_142_v8_ = _dafny.MultiSet([d_141_v7_])
        d_143_v9_: _dafny.Seq
        d_143_v9_ = _dafny.SeqWithoutIsStrInference([not(d_140_v6_)])
        d_144_v10_: _dafny.Seq
        d_144_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxxycbwg"))
        d_145_v11_: _dafny.Map
        d_145_v11_ = _dafny.Map({d_140_v6_: True})
        d_146_v12_: _dafny.Map
        d_146_v12_ = _dafny.Map({d_134_v0_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ue")))})
        d_147_v13_: _dafny.Array
        nw25_ = _dafny.Array(None, 8)
        nw25_[int(0)] = d_134_v0_
        nw25_[int(1)] = len(d_144_v10_)
        nw25_[int(2)] = len(_dafny.Map({d_140_v6_: d_145_v11_}))
        nw25_[int(3)] = d_134_v0_
        nw25_[int(4)] = len(d_138_v4_)
        nw25_[int(5)] = len(d_144_v10_)
        nw25_[int(6)] = len(d_146_v12_)
        nw25_[int(7)] = d_134_v0_
        d_147_v13_ = nw25_
        d_148_globalState_: GlobalState
        nw26_ = GlobalState()
        nw26_.ctor__(870, -430, d_135_v1_, False, _dafny.CodePoint('x'), d_137_v3_, (d_139_v5_) | (d_139_v5_), _dafny.CodePoint('u'), d_142_v8_, False, 270, d_143_v9_, d_147_v13_, True)
        d_148_globalState_ = nw26_
        index20_ = default__.safeIndex(749, (d_147_v13_).length(0))
        (d_147_v13_)[index20_] = d_134_v0_
        index21_ = default__.safeIndex(749, (d_147_v13_).length(0))
        (d_147_v13_)[index21_] = default__.safeModuloInt(((0) - (len(d_144_v10_))) * (427), (0) - (d_134_v0_))
        d_149_v14_: _dafny.Map
        d_149_v14_ = _dafny.Map({(d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]: d_140_v6_})
        d_150_v15_: _dafny.Set
        d_150_v15_ = _dafny.Set({len(d_149_v14_)})
        d_151_v16_: D0
        d_151_v16_ = D0_DC2(d_150_v15_, d_134_v0_, d_140_v6_)
        source3_ = (d_151_v16_ if (d_151_v16_).cf6 else d_151_v16_)
        if source3_.is_DC0:
            d_152___mcc_h0_ = source3_.cf0
            d_153___mcc_h1_ = source3_.cf1
            d_154_cf1_ = d_153___mcc_h1_
            d_155_cf0_ = d_152___mcc_h0_
            d_156_v17_: _dafny.Array
            nw27_ = _dafny.Array(False, 5)
            d_156_v17_ = nw27_
            index22_ = default__.safeIndex(151, (d_156_v17_).length(0))
            (d_156_v17_)[index22_] = False
            d_157_v18_: _dafny.MultiSet
            d_157_v18_ = _dafny.MultiSet([(d_144_v10_) + (d_144_v10_)])
            index23_ = default__.safeIndex(50, (d_156_v17_).length(0))
            (d_156_v17_)[index23_] = d_140_v6_
            index24_ = default__.safeIndex(151, (d_156_v17_).length(0))
            index25_ = default__.safeIndex(749, (d_147_v13_).length(0))
            index26_ = default__.safeIndex(50, (d_156_v17_).length(0))
            rhs15_ = d_154_cf1_
            rhs16_ = default__.safeModuloInt(d_155_cf0_, d_134_v0_)
            rhs17_ = (len(d_144_v10_)) * ((0) - (default__.safeModuloInt(d_134_v0_, d_155_cf0_)))
            rhs18_ = d_157_v18_
            rhs19_ = (d_134_v0_) == (default__.safeModuloInt(d_134_v0_, (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]))
            lhs8_ = d_156_v17_
            lhs9_ = default__.safeIndex(151, (d_156_v17_).length(0))
            lhs10_ = d_147_v13_
            lhs11_ = default__.safeIndex(749, (d_147_v13_).length(0))
            lhs12_ = d_156_v17_
            lhs13_ = default__.safeIndex(50, (d_156_v17_).length(0))
            lhs8_[lhs9_] = rhs15_
            lhs10_[lhs11_] = rhs16_
            d_155_cf0_ = rhs17_
            d_157_v18_ = rhs18_
            lhs12_[lhs13_] = rhs19_
            if not(d_154_cf1_):
                d_158_v19_: bool
                d_159_v20_: _dafny.Array
                d_160_v21_: int
                d_161_v22_: int
                out0_: bool
                out1_: _dafny.Array
                out2_: int
                out3_: int
                out0_, out1_, out2_, out3_ = default__.m0(default__.safeModuloInt(len(d_144_v10_), -6), d_143_v9_, d_148_globalState_)
                d_158_v19_ = out0_
                d_159_v20_ = out1_
                d_160_v21_ = out2_
                d_161_v22_ = out3_
                d_162_v23_: bool
                d_163_v24_: _dafny.Array
                d_164_v25_: int
                d_165_v26_: int
                out4_: bool
                out5_: _dafny.Array
                out6_: int
                out7_: int
                out4_, out5_, out6_, out7_ = default__.m0((default__.fm0(d_145_v11_, d_154_cf1_, False, d_148_globalState_)) * (-280), d_143_v9_, d_148_globalState_)
                d_162_v23_ = out4_
                d_163_v24_ = out5_
                d_164_v25_ = out6_
                d_165_v26_ = out7_
                d_166_v27_: _dafny.MultiSet
                d_166_v27_ = _dafny.MultiSet([d_146_v12_, d_146_v12_, d_146_v12_])
                d_166_v27_ = d_166_v27_
                d_161_v22_ = d_164_v25_
                index27_ = default__.safeIndex(749, (d_147_v13_).length(0))
                (d_147_v13_)[index27_] = d_161_v22_
            elif True:
                d_167_v28_: str
                d_167_v28_ = _dafny.CodePoint('l')
                d_168_v29_: _dafny.MultiSet
                d_168_v29_ = _dafny.MultiSet([d_167_v28_])
                d_169_v30_: _dafny.Map
                d_169_v30_ = _dafny.Map({d_168_v29_: d_144_v10_})
                index28_ = default__.safeIndex(749, (d_147_v13_).length(0))
                (d_147_v13_)[index28_] = (len((d_149_v14_).set((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))], (d_156_v17_)[default__.safeIndex(151, (d_156_v17_).length(0))]))) + (len((d_169_v30_ if d_154_cf1_ else d_169_v30_)))
                d_170_v31_: bool
                d_171_v32_: _dafny.Array
                d_172_v33_: int
                d_173_v34_: int
                out8_: bool
                out9_: _dafny.Array
                out10_: int
                out11_: int
                out8_, out9_, out10_, out11_ = default__.m0((0) - (len(d_139_v5_)), d_143_v9_, d_148_globalState_)
                d_170_v31_ = out8_
                d_171_v32_ = out9_
                d_172_v33_ = out10_
                d_173_v34_ = out11_
                d_174_v35_: C0
                nw28_ = C0()
                nw28_.ctor__(d_154_cf1_)
                d_174_v35_ = nw28_
                d_175_v36_: _dafny.Map
                d_175_v36_ = _dafny.Map({not(d_140_v6_): d_155_cf0_})
                d_176_v37_: _dafny.Map
                d_176_v37_ = _dafny.Map({d_175_v36_: d_144_v10_})
                rhs20_ = d_170_v31_
                rhs21_ = (0) - (default__.safeModuloInt(d_173_v34_, len(((d_176_v37_)[d_175_v36_] if (d_175_v36_) in (d_176_v37_) else d_144_v10_))))
                d_154_cf1_ = rhs20_
                d_155_cf0_ = rhs21_
                d_177_v38_: _dafny.Set
                d_177_v38_ = _dafny.Set({d_145_v11_})
                d_178_v39_: C0
                nw29_ = C0()
                nw29_.ctor__((d_177_v38_).issubset(d_177_v38_))
                d_178_v39_ = nw29_
            index29_ = default__.safeIndex(749, (d_147_v13_).length(0))
            (d_147_v13_)[index29_] = 534
            d_179_v40_: D2
            d_179_v40_ = D2_DC6(_dafny.SeqWithoutIsStrInference([(d_156_v17_)[default__.safeIndex(151, (d_156_v17_).length(0))]]))
            d_180_v41_: _dafny.Map
            d_180_v41_ = _dafny.Map({(d_155_cf0_) * (len(d_138_v4_)): d_179_v40_})
            d_181_v42_: str
            d_181_v42_ = _dafny.CodePoint('r')
            rhs22_ = len(d_180_v41_)
            rhs23_ = d_181_v42_
            rhs24_ = default__.safeDivisionInt(d_155_cf0_, (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))])
            lhs14_ = d_148_globalState_
            d_155_cf0_ = rhs22_
            lhs14_.f4 = rhs23_
            d_155_cf0_ = rhs24_
        elif source3_.is_DC1:
            d_182___mcc_h2_ = source3_.cf2
            d_183___mcc_h3_ = source3_.cf3
            d_184_cf3_ = d_183___mcc_h3_
            d_185_cf2_ = d_182___mcc_h2_
            d_186_v43_: _dafny.Set
            d_186_v43_ = _dafny.Set({not(d_140_v6_)})
            if not(((d_186_v43_) - (d_186_v43_)).isdisjoint(_dafny.Set({not(d_140_v6_), d_184_cf3_}))):
                pat_let_tv6_ = d_150_v15_
                d_187_v44_: _dafny.Array
                nw30_ = _dafny.Array(None, 12)
                nw30_[int(0)] = (d_151_v16_ if d_184_cf3_ else d_151_v16_)
                nw30_[int(1)] = d_151_v16_
                nw30_[int(2)] = d_151_v16_
                nw30_[int(3)] = d_151_v16_
                nw30_[int(4)] = d_151_v16_
                nw30_[int(5)] = d_151_v16_
                nw30_[int(6)] = d_151_v16_
                nw30_[int(7)] = d_151_v16_
                nw30_[int(8)] = D0_DC2(d_150_v15_, (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))], d_140_v6_)
                nw30_[int(9)] = D0_DC2(d_150_v15_, default__.fm0(d_145_v11_, d_184_cf3_, d_140_v6_, d_148_globalState_), False)
                nw30_[int(10)] = d_151_v16_
                def iife7_(_pat_let2_0):
                    def iife8_(d_188_dt__update__tmp_h0_):
                        def iife9_(_pat_let3_0):
                            def iife10_(d_189_dt__update_hcf4_h0_):
                                return D0_DC2(d_189_dt__update_hcf4_h0_, (d_188_dt__update__tmp_h0_).cf5, (d_188_dt__update__tmp_h0_).cf6)
                            return iife10_(_pat_let3_0)
                        return iife9_(pat_let_tv6_)
                    return iife8_(_pat_let2_0)
                nw30_[int(11)] = iife7_(D0_DC2(d_150_v15_, d_134_v0_, d_184_cf3_))
                d_187_v44_ = nw30_
                d_187_v44_ = d_187_v44_
                index30_ = default__.safeIndex(749, (d_147_v13_).length(0))
                (d_147_v13_)[index30_] = (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]
                d_134_v0_ = default__.safeModuloInt((d_134_v0_) - (d_134_v0_), (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))])
                d_144_v10_ = d_144_v10_
                d_190_v45_: C0
                nw31_ = C0()
                nw31_.ctor__(default__.fm4(d_140_v6_, d_186_v43_, (_dafny.MultiSet([(d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]])).cardinality, d_148_globalState_))
                d_190_v45_ = nw31_
            elif True:
                d_191_v46_: D1
                d_191_v46_ = D1_DC5(((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]) <= ((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]))
                d_191_v46_ = d_191_v46_
                d_134_v0_ = (0) - (default__.safeModuloInt(len((d_149_v14_).set(d_134_v0_, d_140_v6_)), d_134_v0_))
                d_192_v47_: C1
                nw32_ = C1()
                nw32_.ctor__(d_184_cf3_, (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))])
                d_192_v47_ = nw32_
                d_193_v48_: _dafny.Map
                d_193_v48_ = _dafny.Map({d_192_v47_: (d_144_v10_)[default__.safeIndex(d_134_v0_, len(d_144_v10_))]})
                d_194_v49_: D3
                d_194_v49_ = D3_DC8(d_192_v47_)
                d_195_v50_: str
                d_195_v50_ = _dafny.CodePoint('w')
                index31_ = default__.safeIndex(749, (d_147_v13_).length(0))
                rhs25_ = ((d_193_v48_)[(d_194_v49_).cf14] if ((d_194_v49_).cf14) in (d_193_v48_) else d_195_v50_)
                rhs26_ = (d_192_v47_).f15
                rhs27_ = default__.fm6(((d_192_v47_).f14 if default__.fm6(default__.fm4(not(d_184_cf3_), d_186_v43_, len(d_138_v4_), d_148_globalState_), (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))], d_148_globalState_) else (d_192_v47_).f14), d_134_v0_, d_148_globalState_)
                rhs28_ = default__.safeDivisionInt((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))], (d_192_v47_).f15)
                lhs15_ = d_148_globalState_
                lhs16_ = d_147_v13_
                lhs17_ = default__.safeIndex(749, (d_147_v13_).length(0))
                lhs15_.f4 = rhs25_
                d_134_v0_ = rhs26_
                d_140_v6_ = rhs27_
                lhs16_[lhs17_] = rhs28_
                (d_192_v47_).m1(d_148_globalState_)
                d_134_v0_ = (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]
            d_196_v51_: D1
            d_196_v51_ = D1_DC4(d_144_v10_)
            pat_let_tv7_ = d_144_v10_
            pat_let_tv8_ = d_144_v10_
            d_197_v52_: _dafny.Array
            nw33_ = _dafny.Array(None, 7)
            nw33_[int(0)] = D1_DC4(d_144_v10_)
            nw33_[int(1)] = d_196_v51_
            nw33_[int(2)] = d_196_v51_
            nw33_[int(3)] = d_196_v51_
            def iife11_(_pat_let4_0):
                def iife12_(d_198_dt__update__tmp_h1_):
                    def iife13_(_pat_let5_0):
                        def iife14_(d_199_dt__update_hcf8_h0_):
                            return D1_DC4(d_199_dt__update_hcf8_h0_)
                        return iife14_(_pat_let5_0)
                    return iife13_(pat_let_tv7_)
                return iife12_(_pat_let4_0)
            nw33_[int(4)] = iife11_(d_196_v51_)
            nw33_[int(5)] = d_196_v51_
            def iife15_(_pat_let6_0):
                def iife16_(d_200_dt__update__tmp_h2_):
                    def iife17_(_pat_let7_0):
                        def iife18_(d_201_dt__update_hcf8_h1_):
                            return D1_DC4(d_201_dt__update_hcf8_h1_)
                        return iife18_(_pat_let7_0)
                    return iife17_(pat_let_tv8_)
                return iife16_(_pat_let6_0)
            nw33_[int(6)] = iife15_(d_196_v51_)
            d_197_v52_ = nw33_
            d_202_v53_: _dafny.Array
            nw34_ = _dafny.Array(False, 29)
            d_202_v53_ = nw34_
            rhs29_ = d_197_v52_
            rhs30_ = d_147_v13_
            rhs31_ = d_202_v53_
            d_197_v52_ = rhs29_
            d_147_v13_ = rhs30_
            d_202_v53_ = rhs31_
            d_203_v54_: bool
            d_204_v55_: _dafny.Array
            d_205_v56_: int
            d_206_v57_: int
            out12_: bool
            out13_: _dafny.Array
            out14_: int
            out15_: int
            out12_, out13_, out14_, out15_ = default__.m0((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))], d_143_v9_, d_148_globalState_)
            d_203_v54_ = out12_
            d_204_v55_ = out13_
            d_205_v56_ = out14_
            d_206_v57_ = out15_
            d_207_v58_: _dafny.Seq
            d_207_v58_ = _dafny.SeqWithoutIsStrInference([d_143_v9_, d_143_v9_])
            d_208_v59_: bool
            d_209_v60_: _dafny.Array
            d_210_v61_: int
            d_211_v62_: int
            out16_: bool
            out17_: _dafny.Array
            out18_: int
            out19_: int
            out16_, out17_, out18_, out19_ = default__.m0(-643, (d_207_v58_)[default__.safeIndex(d_206_v57_, len(d_207_v58_))], d_148_globalState_)
            d_208_v59_ = out16_
            d_209_v60_ = out17_
            d_210_v61_ = out18_
            d_211_v62_ = out19_
        elif source3_.is_DC2:
            d_212___mcc_h4_ = source3_.cf4
            d_213___mcc_h5_ = source3_.cf5
            d_214___mcc_h6_ = source3_.cf6
            d_215_cf6_ = d_214___mcc_h6_
            d_216_cf5_ = d_213___mcc_h5_
            d_217_cf4_ = d_212___mcc_h4_
            d_218_v63_: _dafny.MultiSet
            d_218_v63_ = _dafny.MultiSet([d_143_v9_, d_143_v9_, d_143_v9_])
            d_219_v64_: D0
            d_219_v64_ = D0_DC0(len(d_143_v9_), d_140_v6_)
            d_220_v65_: _dafny.Seq
            d_220_v65_ = _dafny.SeqWithoutIsStrInference([d_143_v9_])
            d_221_v66_: _dafny.Map
            d_221_v66_ = _dafny.Map({(d_219_v64_).cf1: d_220_v65_})
            d_222_v67_: _dafny.Array
            nw35_ = _dafny.Array(None, 8)
            nw35_[int(0)] = d_218_v63_
            nw35_[int(1)] = (d_218_v63_) | ((d_218_v63_).set(d_143_v9_, default__.abs(596)))
            nw35_[int(2)] = (d_218_v63_ if d_215_cf6_ else d_218_v63_)
            nw35_[int(3)] = d_218_v63_
            nw35_[int(4)] = _dafny.MultiSet([d_143_v9_, d_143_v9_])
            nw35_[int(5)] = d_218_v63_
            nw35_[int(6)] = (_dafny.MultiSet(((d_221_v66_)[d_215_cf6_] if (d_215_cf6_) in (d_221_v66_) else d_220_v65_)) if d_215_cf6_ else d_218_v63_)
            nw35_[int(7)] = d_218_v63_
            d_222_v67_ = nw35_
            index32_ = default__.safeIndex(276, (d_222_v67_).length(0))
            (d_222_v67_)[index32_] = _dafny.MultiSet([d_143_v9_, d_143_v9_, d_143_v9_])
            d_223_v68_: _dafny.Seq
            d_223_v68_ = _dafny.SeqWithoutIsStrInference([d_218_v63_, (d_218_v63_) | (d_218_v63_), (d_218_v63_) | (d_218_v63_), d_218_v63_, _dafny.MultiSet([d_143_v9_, d_143_v9_])])
            index33_ = default__.safeIndex(276, (d_222_v67_).length(0))
            (d_222_v67_)[index33_] = (d_223_v68_)[default__.safeIndex(default__.safeDivisionInt((d_138_v4_)[default__.safeIndex(504, len(d_138_v4_))], len(_dafny.SeqWithoutIsStrInference([False]))), len(d_223_v68_))]
            if d_140_v6_:
                d_224_v69_: C0
                nw36_ = C0()
                nw36_.ctor__(default__.fm6(d_140_v6_, -901, d_148_globalState_))
                d_224_v69_ = nw36_
                d_144_v10_ = (d_224_v69_).fm5(d_215_cf6_, d_216_cf5_, default__.safeDivisionInt(len(d_145_v11_), d_216_cf5_), (d_224_v69_).f16, d_148_globalState_)
                d_225_v70_: D4
                d_225_v70_ = D4_DC11(d_139_v5_)
                d_226_v71_: C1
                nw37_ = C1()
                nw37_.ctor__(default__.fm6(not(True), 999, d_148_globalState_), len((d_225_v70_).cf17))
                d_226_v71_ = nw37_
                (d_224_v69_).m1(d_148_globalState_)
                d_227_v72_: C1
                nw38_ = C1()
                nw38_.ctor__((d_144_v10_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_228_i0_ in range(default__.abs(436))])), len(d_145_v11_))
                d_227_v72_ = nw38_
            elif True:
                d_229_v73_: C0
                nw39_ = C0()
                nw39_.ctor__(d_215_cf6_)
                d_229_v73_ = nw39_
                d_230_v74_: _dafny.Seq
                d_230_v74_ = _dafny.SeqWithoutIsStrInference([d_229_v73_, d_229_v73_])
                index34_ = default__.safeIndex(749, (d_147_v13_).length(0))
                rhs32_ = ((d_151_v16_).cf5) + ((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))])
                rhs33_ = d_229_v73_
                rhs34_ = (d_138_v4_) + ((d_138_v4_) + (d_138_v4_))
                rhs35_ = d_216_cf5_
                rhs36_ = _dafny.SeqWithoutIsStrInference([d_229_v73_, d_229_v73_, d_229_v73_])
                lhs18_ = d_147_v13_
                lhs19_ = default__.safeIndex(749, (d_147_v13_).length(0))
                d_216_cf5_ = rhs32_
                d_229_v73_ = rhs33_
                d_138_v4_ = rhs34_
                lhs18_[lhs19_] = rhs35_
                d_230_v74_ = rhs36_
                index35_ = default__.safeIndex(749, (d_147_v13_).length(0))
                (d_147_v13_)[index35_] = len((_dafny.SeqWithoutIsStrInference([d_216_cf5_])) + (d_138_v4_))
                d_215_cf6_ = d_215_cf6_
                d_215_cf6_ = not (not(d_140_v6_)) or ((d_229_v73_).f16)
                d_134_v0_ = default__.safeModuloInt(d_134_v0_, d_134_v0_)
            index36_ = default__.safeIndex(749, (d_147_v13_).length(0))
            (d_147_v13_)[index36_] = d_134_v0_
            if d_215_cf6_:
                d_134_v0_ = (d_219_v64_).cf0
                d_231_v75_: _dafny.Array
                def lambda6_(d_232_v6_):
                    def lambda7_(d_233_i1_):
                        return not(d_232_v6_)

                    return lambda7_

                init3_ = lambda6_(d_140_v6_)
                nw40_ = _dafny.Array(None, 23)
                for i0_3_ in range(nw40_.length(0)):
                    nw40_[i0_3_] = init3_(i0_3_)
                d_231_v75_ = nw40_
                index37_ = default__.safeIndex(380, (d_231_v75_).length(0))
                (d_231_v75_)[index37_] = d_215_cf6_
                index38_ = default__.safeIndex(380, (d_231_v75_).length(0))
                (d_231_v75_)[index38_] = not(d_140_v6_)
                d_143_v9_ = (d_143_v9_).set(default__.safeIndex(d_216_cf5_, len(d_143_v9_)), False)
                d_140_v6_ = ((d_149_v14_)[157] if (157) in (d_149_v14_) else (d_231_v75_)[default__.safeIndex(380, (d_231_v75_).length(0))])
                d_234_v76_: str
                d_234_v76_ = _dafny.CodePoint('j')
                d_235_v77_: _dafny.Seq
                d_235_v77_ = _dafny.SeqWithoutIsStrInference([d_144_v10_])
                d_236_v78_: _dafny.Array
                nw41_ = _dafny.Array(D0.default()(), 10)
                d_236_v78_ = nw41_
                d_237_v79_: _dafny.Map
                d_237_v79_ = _dafny.Map({d_134_v0_: d_141_v7_})
                d_238_v80_: _dafny.Map
                d_238_v80_ = _dafny.Map({d_236_v78_: d_237_v79_})
                rhs37_ = d_234_v76_
                rhs38_ = d_231_v75_
                rhs39_ = (d_144_v10_) + ((d_235_v77_)[default__.safeIndex(d_216_cf5_, len(d_235_v77_))])
                rhs40_ = (len(_dafny.SeqWithoutIsStrInference([(d_231_v75_)[default__.safeIndex(380, (d_231_v75_).length(0))]]))) in (((d_238_v80_)[d_236_v78_] if (d_236_v78_) in (d_238_v80_) else d_237_v79_))
                rhs41_ = (d_231_v75_)[default__.safeIndex(380, (d_231_v75_).length(0))]
                lhs20_ = d_148_globalState_
                lhs20_.f4 = rhs37_
                d_231_v75_ = rhs38_
                d_144_v10_ = rhs39_
                d_140_v6_ = rhs40_
                d_140_v6_ = rhs41_
            elif True:
                d_239_v81_: str
                d_239_v81_ = _dafny.CodePoint('h')
                index39_ = default__.safeIndex(172, (d_136_v2_).length(0))
                (d_136_v2_)[index39_] = d_239_v81_
                index40_ = default__.safeIndex(172, (d_136_v2_).length(0))
                (d_136_v2_)[index40_] = d_239_v81_
                d_216_cf5_ = (0) - ((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))])
                d_240_v82_: _dafny.MultiSet
                d_240_v82_ = _dafny.MultiSet([len(d_149_v14_), (0) - (d_216_cf5_)])
                index41_ = default__.safeIndex(749, (d_147_v13_).length(0))
                rhs42_ = d_147_v13_
                rhs43_ = not(False)
                rhs44_ = d_216_cf5_
                rhs45_ = d_240_v82_
                lhs21_ = d_148_globalState_
                lhs22_ = d_147_v13_
                lhs23_ = default__.safeIndex(749, (d_147_v13_).length(0))
                lhs21_.f12 = rhs42_
                d_140_v6_ = rhs43_
                lhs22_[lhs23_] = rhs44_
                d_240_v82_ = rhs45_
                d_241_v83_: bool
                d_242_v84_: _dafny.Array
                d_243_v85_: int
                d_244_v86_: int
                out20_: bool
                out21_: _dafny.Array
                out22_: int
                out23_: int
                out20_, out21_, out22_, out23_ = default__.m0((0) - ((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]), default__.fm10(not(d_215_cf6_), (d_136_v2_)[default__.safeIndex(172, (d_136_v2_).length(0))], d_134_v0_, d_148_globalState_), d_148_globalState_)
                d_241_v83_ = out20_
                d_242_v84_ = out21_
                d_243_v85_ = out22_
                d_244_v86_ = out23_
                d_140_v6_ = (d_140_v6_) == (True)
        elif True:
            d_245___mcc_h7_ = source3_.cf7
            d_246_cf7_ = d_245___mcc_h7_
            d_247_v87_: _dafny.Array
            nw42_ = _dafny.Array(None, 8)
            d_247_v87_ = nw42_
            d_247_v87_ = d_247_v87_
            d_140_v6_ = (d_134_v0_) > (d_134_v0_)
            d_248_v88_: _dafny.Map
            d_248_v88_ = _dafny.Map({(d_143_v9_)[default__.safeIndex((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))], len(d_143_v9_))]: (d_149_v14_ if d_140_v6_ else d_149_v14_)})
            d_248_v88_ = (d_248_v88_) | (d_248_v88_)
            d_140_v6_ = not(((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]) > (default__.safeDivisionInt(d_134_v0_, d_134_v0_)))
        d_140_v6_ = (d_140_v6_) or (((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]) != ((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]))
        d_249_i2_: int
        d_249_i2_ = 0
        with _dafny.label("1"):
            while d_140_v6_:
                with _dafny.c_label("1"):
                    if (d_249_i2_) >= (100):
                        raise _dafny.Break("1")
                    d_249_i2_ = (d_249_i2_) + (1)
                    d_146_v12_ = (d_146_v12_).set((d_134_v0_) * (d_134_v0_), ((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]) - ((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]))
                    d_140_v6_ = (d_134_v0_) in ((d_149_v14_ if d_140_v6_ else d_149_v14_))
                    d_250_v89_: _dafny.Array
                    def lambda8_(d_251_v0_, d_252_v13_):
                        def lambda9_(d_253_i3_):
                            return (d_251_v0_) < ((d_252_v13_)[default__.safeIndex(749, (d_252_v13_).length(0))])

                        return lambda9_

                    init4_ = lambda8_(d_134_v0_, d_147_v13_)
                    nw43_ = _dafny.Array(None, 11)
                    for i0_4_ in range(nw43_.length(0)):
                        nw43_[i0_4_] = init4_(i0_4_)
                    d_250_v89_ = nw43_
                    def lambda10_(d_254_v6_):
                        def lambda11_(d_255_i4_):
                            return d_254_v6_

                        return lambda11_

                    init5_ = lambda10_(d_140_v6_)
                    nw44_ = _dafny.Array(None, 16)
                    for i0_5_ in range(nw44_.length(0)):
                        nw44_[i0_5_] = init5_(i0_5_)
                    d_250_v89_ = nw44_
                    d_256_v90_: C1
                    nw45_ = C1()
                    nw45_.ctor__(d_140_v6_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_257_i5_ in range(default__.abs(561))])))
                    d_256_v90_ = nw45_
                    index42_ = default__.safeIndex(749, (d_147_v13_).length(0))
                    rhs46_ = d_256_v90_
                    rhs47_ = 509
                    lhs24_ = d_147_v13_
                    lhs25_ = default__.safeIndex(749, (d_147_v13_).length(0))
                    d_256_v90_ = rhs46_
                    lhs24_[lhs25_] = rhs47_
                    pass
            pass
        d_258_v91_: _dafny.Array
        def lambda12_(d_259_v6_):
            def lambda13_(d_260_i6_):
                return d_259_v6_

            return lambda13_

        init6_ = lambda12_(d_140_v6_)
        nw46_ = _dafny.Array(None, 8)
        for i0_6_ in range(nw46_.length(0)):
            nw46_[i0_6_] = init6_(i0_6_)
        d_258_v91_ = nw46_
        d_261_v92_: C0
        nw47_ = C0()
        nw47_.ctor__(False)
        d_261_v92_ = nw47_
        d_262_v93_: _dafny.Array
        nw48_ = _dafny.Array(None, 2)
        nw48_[int(0)] = (d_261_v92_ if d_140_v6_ else d_261_v92_)
        nw48_[int(1)] = d_261_v92_
        d_262_v93_ = nw48_
        index43_ = default__.safeIndex(538, (d_262_v93_).length(0))
        (d_262_v93_)[index43_] = d_261_v92_
        d_263_v94_: C1
        nw49_ = C1()
        nw49_.ctor__(d_140_v6_, d_134_v0_)
        d_263_v94_ = nw49_
        d_264_v95_: _dafny.Seq
        d_264_v95_ = _dafny.SeqWithoutIsStrInference([d_263_v94_, d_263_v94_])
        d_265_v96_: str
        d_265_v96_ = _dafny.CodePoint('c')
        d_266_v97_: _dafny.Map
        d_266_v97_ = _dafny.Map({len(d_264_v95_): d_265_v96_})
        index44_ = default__.safeIndex(538, (d_262_v93_).length(0))
        rhs48_ = d_258_v91_
        rhs49_ = d_261_v92_
        rhs50_ = ((d_266_v97_)[d_134_v0_] if (d_134_v0_) in (d_266_v97_) else (d_144_v10_)[default__.safeIndex((d_138_v4_)[default__.safeIndex(-258, len(d_138_v4_))], len(d_144_v10_))])
        lhs26_ = d_262_v93_
        lhs27_ = default__.safeIndex(538, (d_262_v93_).length(0))
        lhs28_ = d_148_globalState_
        d_258_v91_ = rhs48_
        lhs26_[lhs27_] = rhs49_
        lhs28_.f4 = rhs50_
        (d_261_v92_).m1(d_148_globalState_)
        source4_ = d_151_v16_
        if source4_.is_DC0:
            d_267___mcc_h8_ = source4_.cf0
            d_268___mcc_h9_ = source4_.cf1
            d_269_cf1_ = d_268___mcc_h9_
            d_270_cf0_ = d_267___mcc_h8_
            d_134_v0_ = d_270_cf0_
            d_271_v98_: bool
            d_272_v99_: _dafny.Array
            d_273_v100_: int
            d_274_v101_: int
            out24_: bool
            out25_: _dafny.Array
            out26_: int
            out27_: int
            out24_, out25_, out26_, out27_ = default__.m0((d_263_v94_).f15, (d_143_v9_) + (_dafny.SeqWithoutIsStrInference([d_269_cf1_, not(d_269_cf1_), d_140_v6_, (d_263_v94_).f14, (d_263_v94_).f14])), d_148_globalState_)
            d_271_v98_ = out24_
            d_272_v99_ = out25_
            d_273_v100_ = out26_
            d_274_v101_ = out27_
            index45_ = default__.safeIndex(749, (d_147_v13_).length(0))
            (d_147_v13_)[index45_] = default__.safeDivisionInt(391, d_270_cf0_)
            d_144_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxrlkl"))
        elif source4_.is_DC1:
            d_275___mcc_h10_ = source4_.cf2
            d_276___mcc_h11_ = source4_.cf3
            d_277_cf3_ = d_276___mcc_h11_
            d_278_cf2_ = d_275___mcc_h10_
            d_279_v102_: _dafny.Array
            nw50_ = _dafny.Array(_dafny.Set({}), 2)
            d_279_v102_ = nw50_
            d_280_v103_: _dafny.Set
            d_280_v103_ = _dafny.Set({False, (d_263_v94_).f14})
            index46_ = default__.safeIndex(723, (d_279_v102_).length(0))
            (d_279_v102_)[index46_] = (_dafny.Set({d_277_cf3_, (d_261_v92_).f16})) - (d_280_v103_)
            index47_ = default__.safeIndex(723, (d_279_v102_).length(0))
            (d_279_v102_)[index47_] = (d_280_v103_ if (d_263_v94_).f14 else d_280_v103_)
            index48_ = default__.safeIndex(749, (d_147_v13_).length(0))
            (d_147_v13_)[index48_] = (d_263_v94_).f15
            d_281_v104_: D1
            d_281_v104_ = D1_DC5(True)
            d_282_v105_: _dafny.Map
            d_282_v105_ = _dafny.Map({D0_DC3(d_265_v96_): (d_263_v94_).f15})
            d_283_v106_: D0
            d_283_v106_ = D0_DC3(d_265_v96_)
            d_284_v107_: C1
            nw51_ = C1()
            nw51_.ctor__((d_281_v104_).cf9, ((d_282_v105_)[d_283_v106_] if (d_283_v106_) in (d_282_v105_) else (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]))
            d_284_v107_ = nw51_
            d_285_v108_: _dafny.Array
            nw52_ = _dafny.Array(_dafny.Map({}), 11)
            d_285_v108_ = nw52_
            d_286_v109_: T0
            nw53_ = C0()
            nw53_.ctor__((d_261_v92_).f16)
            d_286_v109_ = nw53_
            d_287_v110_: _dafny.Map
            d_287_v110_ = _dafny.Map({_dafny.Set({(d_263_v94_).f15}): d_286_v109_})
            index49_ = default__.safeIndex(337, (d_285_v108_).length(0))
            (d_285_v108_)[index49_] = d_287_v110_
            index50_ = default__.safeIndex(337, (d_285_v108_).length(0))
            index51_ = default__.safeIndex(749, (d_147_v13_).length(0))
            def iife19_():
                coll3_ = _dafny.Map()
                compr_3_: _dafny.Seq
                for compr_3_ in (_dafny.SeqWithoutIsStrInference([d_144_v10_, d_144_v10_])).Elements:
                    d_288_v111_: _dafny.Seq = compr_3_
                    if (d_288_v111_) in (_dafny.SeqWithoutIsStrInference([d_144_v10_, d_144_v10_])):
                        coll3_[d_288_v111_] = D0_DC1(d_278_cf2_, (d_263_v94_).f14)
                return _dafny.Map(coll3_)
            rhs51_ = (d_263_v94_).f15
            rhs52_ = len(iife19_()
            )
            rhs53_ = ((d_287_v110_) | (d_287_v110_)) | (d_287_v110_)
            rhs54_ = (d_141_v7_).cardinality
            lhs29_ = d_285_v108_
            lhs30_ = default__.safeIndex(337, (d_285_v108_).length(0))
            lhs31_ = d_147_v13_
            lhs32_ = default__.safeIndex(749, (d_147_v13_).length(0))
            d_134_v0_ = rhs51_
            d_134_v0_ = rhs52_
            lhs29_[lhs30_] = rhs53_
            lhs31_[lhs32_] = rhs54_
        elif source4_.is_DC2:
            d_289___mcc_h12_ = source4_.cf4
            d_290___mcc_h13_ = source4_.cf5
            d_291___mcc_h14_ = source4_.cf6
            d_292_cf6_ = d_291___mcc_h14_
            d_293_cf5_ = d_290___mcc_h13_
            d_294_cf4_ = d_289___mcc_h12_
            d_295_v112_: C0
            nw54_ = C0()
            nw54_.ctor__(d_140_v6_)
            d_295_v112_ = nw54_
            d_140_v6_ = not(((d_295_v112_).f16) and ((d_261_v92_).f16))
            d_140_v6_ = (len(default__.fm9(d_265_v96_, d_134_v0_, 904, _dafny.Map({not((d_263_v94_).f14): d_293_cf5_}), d_148_globalState_))) >= ((d_263_v94_).f15)
            d_296_v113_: _dafny.Array
            nw55_ = _dafny.Array(_dafny.Map({}), 13)
            d_296_v113_ = nw55_
            d_297_v114_: _dafny.Map
            d_297_v114_ = _dafny.Map({d_144_v10_: d_147_v13_})
            index52_ = default__.safeIndex(703, (d_296_v113_).length(0))
            (d_296_v113_)[index52_] = d_297_v114_
            index53_ = default__.safeIndex(703, (d_296_v113_).length(0))
            (d_296_v113_)[index53_] = (_dafny.Map({d_144_v10_: d_147_v13_})) | ((d_297_v114_) | (d_297_v114_))
        elif True:
            d_298___mcc_h15_ = source4_.cf7
            d_299_cf7_ = d_298___mcc_h15_
            index54_ = default__.safeIndex(749, (d_147_v13_).length(0))
            (d_147_v13_)[index54_] = (d_263_v94_).f15
            if (d_263_v94_).f14:
                d_300_v115_: D3
                d_300_v115_ = D3_DC10(D3_DC9(d_299_cf7_))
                pat_let_tv9_ = d_263_v94_
                d_301_v116_: _dafny.Map
                def iife20_(_pat_let8_0):
                    def iife21_(d_302_dt__update__tmp_h3_):
                        def iife22_(_pat_let9_0):
                            def iife23_(d_303_dt__update_hcf16_h0_):
                                return D3_DC10(d_303_dt__update_hcf16_h0_)
                            return iife23_(_pat_let9_0)
                        return iife22_(D3_DC8(pat_let_tv9_))
                    return iife21_(_pat_let8_0)
                d_301_v116_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_261_v92_).f16, (d_143_v9_)[default__.safeIndex(len(d_264_v95_), len(d_143_v9_))]])): iife20_(d_300_v115_)})
                d_301_v116_ = (d_301_v116_).set((d_263_v94_).f15, D3_DC10(D3_DC9(d_265_v96_)))
                d_304_v117_: T0
                nw56_ = C1()
                nw56_.ctor__((d_261_v92_).f16, (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))])
                d_304_v117_ = nw56_
                d_304_v117_ = d_304_v117_
                d_305_v118_: D3
                d_305_v118_ = D3_DC8(d_263_v94_)
                d_305_v118_ = d_305_v118_
                d_144_v10_ = d_144_v10_
                d_306_v119_: _dafny.Map
                d_306_v119_ = _dafny.Map({default__.fm6((d_261_v92_).f16, (d_263_v94_).f15, d_148_globalState_): d_143_v9_})
                d_306_v119_ = d_306_v119_
            elif True:
                d_307_v120_: _dafny.MultiSet
                d_307_v120_ = _dafny.MultiSet([d_299_cf7_])
                d_308_v121_: _dafny.Map
                d_308_v121_ = _dafny.Map({d_140_v6_: d_307_v120_})
                d_309_v122_: _dafny.Set
                d_309_v122_ = _dafny.Set({False, not((d_263_v94_).f14)})
                d_310_v123_: _dafny.Seq
                d_310_v123_ = _dafny.SeqWithoutIsStrInference([((d_308_v121_)[(d_261_v92_).f16] if ((d_261_v92_).f16) in (d_308_v121_) else (d_307_v120_).set(_dafny.CodePoint('j'), default__.abs(len(d_309_v122_))))])
                index55_ = default__.safeIndex(749, (d_147_v13_).length(0))
                (d_147_v13_)[index55_] = default__.safeModuloInt(len((d_310_v123_) + (d_310_v123_)), (d_263_v94_).f15)
                d_134_v0_ = (d_263_v94_).f15
                d_311_v124_: D1
                d_311_v124_ = D1_DC5(d_140_v6_)
                d_312_v125_: _dafny.Map
                d_312_v125_ = _dafny.Map({(d_263_v94_).f15: d_262_v93_})
                index56_ = default__.safeIndex(538, (d_262_v93_).length(0))
                rhs55_ = d_147_v13_
                rhs56_ = d_311_v124_
                rhs57_ = ((d_145_v11_)[((d_145_v11_)[(d_261_v92_).f16] if ((d_261_v92_).f16) in (d_145_v11_) else (d_261_v92_).f16)] if (((d_145_v11_)[(d_261_v92_).f16] if ((d_261_v92_).f16) in (d_145_v11_) else (d_261_v92_).f16)) in (d_145_v11_) else (d_261_v92_).f16)
                rhs58_ = (d_312_v125_) | ((d_312_v125_).set(d_134_v0_, d_262_v93_))
                rhs59_ = d_261_v92_
                lhs33_ = d_148_globalState_
                lhs34_ = d_262_v93_
                lhs35_ = default__.safeIndex(538, (d_262_v93_).length(0))
                lhs33_.f12 = rhs55_
                d_311_v124_ = rhs56_
                d_140_v6_ = rhs57_
                d_312_v125_ = rhs58_
                lhs34_[lhs35_] = rhs59_
                d_134_v0_ = d_134_v0_
                index57_ = default__.safeIndex(749, (d_147_v13_).length(0))
                rhs60_ = (default__.safeDivisionInt((d_263_v94_).f15, d_134_v0_)) + (default__.safeDivisionInt((d_263_v94_).f15, (d_263_v94_).f15))
                rhs61_ = d_145_v11_
                rhs62_ = d_299_cf7_
                rhs63_ = (d_144_v10_) + (d_144_v10_)
                lhs36_ = d_147_v13_
                lhs37_ = default__.safeIndex(749, (d_147_v13_).length(0))
                lhs36_[lhs37_] = rhs60_
                d_145_v11_ = rhs61_
                d_299_cf7_ = rhs62_
                d_144_v10_ = rhs63_
            d_147_v13_ = d_147_v13_
            d_140_v6_ = (d_138_v4_) < (d_138_v4_)
        d_313_v126_: _dafny.MultiSet
        d_313_v126_ = _dafny.MultiSet([d_265_v96_])
        d_313_v126_ = ((d_313_v126_).intersection(default__.fm11(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avyxie")), d_148_globalState_))).set(d_265_v96_, default__.abs(((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]) * ((d_263_v94_).f15)))
        d_140_v6_ = ((d_263_v94_).f15) > (len(d_144_v10_))
        pat_let_tv10_ = d_151_v16_
        pat_let_tv11_ = d_150_v15_
        pat_let_tv12_ = d_263_v94_
        pat_let_tv13_ = d_151_v16_
        def lambda14_(source6_):
            if source6_.is_DC12:
                d_314___mcc_h24_ = source6_.cf18
                d_315___mcc_h25_ = source6_.cf19
                d_316___mcc_h26_ = source6_.cf20
                d_317_cf20_ = d_316___mcc_h26_
                d_318_cf19_ = d_315___mcc_h25_
                d_319_cf18_ = d_314___mcc_h24_
                return pat_let_tv10_
            elif source6_.is_DC11:
                d_320___mcc_h27_ = source6_.cf17
                d_321_cf17_ = d_320___mcc_h27_
                return D0_DC2(pat_let_tv11_, (pat_let_tv12_).f15, False)
            elif True:
                d_322___mcc_h28_ = source6_.cf21
                d_323_cf21_ = d_322___mcc_h28_
                return pat_let_tv13_

        source5_ = lambda14_(default__.fm12(d_140_v6_, (d_263_v94_).f15, d_148_globalState_))
        if source5_.is_DC0:
            d_324___mcc_h16_ = source5_.cf0
            d_325___mcc_h17_ = source5_.cf1
            d_326_cf1_ = d_325___mcc_h17_
            d_327_cf0_ = d_324___mcc_h16_
            d_328_v127_: C1
            nw57_ = C1()
            nw57_.ctor__((d_261_v92_).f16, (d_263_v94_).f15)
            d_328_v127_ = nw57_
            d_140_v6_ = d_140_v6_
            if (d_328_v127_).f14:
                d_329_v128_: T0
                nw58_ = C1()
                nw58_.ctor__(True, (d_263_v94_).f15)
                d_329_v128_ = nw58_
                d_330_v129_: _dafny.MultiSet
                d_330_v129_ = _dafny.MultiSet([(d_263_v94_).f15, (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]])
                d_331_v130_: D4
                d_331_v130_ = D4_DC12((d_263_v94_).f15, d_265_v96_, d_330_v129_)
                d_332_v131_: _dafny.Map
                d_332_v131_ = _dafny.Map({d_138_v4_: d_331_v130_})
                d_333_v132_: _dafny.Map
                d_333_v132_ = _dafny.Map({d_329_v128_: len(d_332_v131_)})
                d_326_cf1_ = (d_333_v132_) == (d_333_v132_)
                d_334_v133_: _dafny.Map
                d_334_v133_ = _dafny.Map({(d_328_v127_).f14: d_143_v9_})
                d_335_v134_: _dafny.Map
                d_335_v134_ = _dafny.Map({d_334_v133_: d_147_v13_})
                d_147_v13_ = ((d_335_v134_)[(d_334_v133_).set((d_263_v94_).f14, d_143_v9_)] if ((d_334_v133_).set((d_263_v94_).f14, d_143_v9_)) in (d_335_v134_) else (d_147_v13_ if not(default__.fm6((d_328_v127_).f14, (d_263_v94_).f15, d_148_globalState_)) else d_147_v13_))
                index58_ = default__.safeIndex(749, (d_147_v13_).length(0))
                (d_147_v13_)[index58_] = ((0) - ((d_263_v94_).f15)) * ((d_263_v94_).f15)
                d_336_v135_: _dafny.Array
                nw59_ = _dafny.Array(D3.default()(), 28)
                d_336_v135_ = nw59_
                d_337_v136_: D3
                d_337_v136_ = D3_DC8(d_263_v94_)
                index59_ = default__.safeIndex(844, (d_336_v135_).length(0))
                (d_336_v135_)[index59_] = d_337_v136_
                index60_ = default__.safeIndex(844, (d_336_v135_).length(0))
                (d_336_v135_)[index60_] = D3_DC8(d_328_v127_)
                (d_148_globalState_).f4 = _dafny.CodePoint('w')
            elif True:
                index61_ = default__.safeIndex(749, (d_147_v13_).length(0))
                (d_147_v13_)[index61_] = (d_263_v94_).fm1(d_144_v10_, (d_328_v127_).f15, (d_263_v94_).f15, d_148_globalState_)
                (d_263_v94_).m1(d_148_globalState_)
                d_140_v6_ = d_326_cf1_
                index62_ = default__.safeIndex(321, (d_258_v91_).length(0))
                (d_258_v91_)[index62_] = not(True)
                d_338_v137_: _dafny.MultiSet
                d_338_v137_ = _dafny.MultiSet([(d_262_v93_)[default__.safeIndex(538, (d_262_v93_).length(0))], (d_262_v93_)[default__.safeIndex(538, (d_262_v93_).length(0))], d_261_v92_, d_261_v92_, (d_262_v93_)[default__.safeIndex(538, (d_262_v93_).length(0))]])
                index63_ = default__.safeIndex(321, (d_258_v91_).length(0))
                (d_258_v91_)[index63_] = (d_338_v137_).isdisjoint(d_338_v137_)
                d_339_v138_: _dafny.Set
                d_339_v138_ = _dafny.Set({(d_258_v91_)[default__.safeIndex(321, (d_258_v91_).length(0))], (d_261_v92_).f16})
                d_140_v6_ = default__.fm4((d_263_v94_).f14, d_339_v138_, (d_138_v4_)[default__.safeIndex((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))], len(d_138_v4_))], d_148_globalState_)
            d_340_v139_: _dafny.Map
            d_340_v139_ = _dafny.Map({d_141_v7_: ((d_262_v93_)[default__.safeIndex(538, (d_262_v93_).length(0))] if (d_263_v94_).f14 else d_261_v92_)})
            d_341_v140_: _dafny.Seq
            d_341_v140_ = _dafny.SeqWithoutIsStrInference([(d_144_v10_) + (d_144_v10_), _dafny.SeqWithoutIsStrInference([d_265_v96_ for d_342_i7_ in range(default__.abs(854))]), _dafny.SeqWithoutIsStrInference([d_265_v96_ for d_343_i8_ in range(default__.abs(240))])])
            d_344_v141_: _dafny.Set
            d_344_v141_ = _dafny.Set({d_140_v6_, d_326_cf1_, d_326_cf1_, (d_261_v92_).f16})
            d_345_v142_: _dafny.Map
            d_345_v142_ = _dafny.Map({(d_261_v92_).f16: (d_263_v94_).f15})
            d_346_v143_: _dafny.Set
            d_346_v143_ = _dafny.Set({d_344_v141_, d_344_v141_, d_344_v141_, default__.fm9(d_265_v96_, d_327_cf0_, (d_263_v94_).f15, (d_345_v142_).set(d_140_v6_, d_327_cf0_), d_148_globalState_), d_344_v141_})
            rhs64_ = (d_150_v15_).issubset(_dafny.Set({(d_263_v94_).f15}))
            rhs65_ = default__.fm6((False) and ((d_328_v127_).f14), d_327_cf0_, d_148_globalState_)
            rhs66_ = _dafny.Map({d_141_v7_: (d_262_v93_)[default__.safeIndex(538, (d_262_v93_).length(0))]})
            rhs67_ = (d_262_v93_)[default__.safeIndex(538, (d_262_v93_).length(0))]
            rhs68_ = default__.fm13(D0_DC1(d_346_v143_, True), (d_143_v9_)[default__.safeIndex(-196, len(d_143_v9_))], d_148_globalState_)
            d_326_cf1_ = rhs64_
            d_140_v6_ = rhs65_
            d_340_v139_ = rhs66_
            d_261_v92_ = rhs67_
            d_341_v140_ = rhs68_
        elif source5_.is_DC1:
            d_347___mcc_h18_ = source5_.cf2
            d_348___mcc_h19_ = source5_.cf3
            d_349_cf3_ = d_348___mcc_h19_
            d_350_cf2_ = d_347___mcc_h18_
            d_351_v144_: _dafny.Array
            nw60_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
            d_351_v144_ = nw60_
            index64_ = default__.safeIndex(140, (d_351_v144_).length(0))
            (d_351_v144_)[index64_] = d_144_v10_
            index65_ = default__.safeIndex(140, (d_351_v144_).length(0))
            (d_351_v144_)[index65_] = d_144_v10_
            index66_ = default__.safeIndex(749, (d_147_v13_).length(0))
            (d_147_v13_)[index66_] = (d_263_v94_).f15
            (d_263_v94_).m1(d_148_globalState_)
            d_352_v145_: _dafny.Map
            d_352_v145_ = _dafny.Map({d_258_v91_: len(d_144_v10_)})
            d_352_v145_ = (d_352_v145_).set(d_258_v91_, (d_263_v94_).f15)
        elif source5_.is_DC2:
            d_353___mcc_h20_ = source5_.cf4
            d_354___mcc_h21_ = source5_.cf5
            d_355___mcc_h22_ = source5_.cf6
            d_356_cf6_ = d_355___mcc_h22_
            d_357_cf5_ = d_354___mcc_h21_
            d_358_cf4_ = d_353___mcc_h20_
            d_140_v6_ = (d_263_v94_).f14
            d_140_v6_ = False
            d_359_v146_: C1
            nw61_ = C1()
            nw61_.ctor__(not(((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]) < (d_357_cf5_)), default__.safeDivisionInt(d_134_v0_, (d_263_v94_).f15))
            d_359_v146_ = nw61_
            d_360_v147_: bool
            d_361_v148_: _dafny.Array
            d_362_v149_: int
            d_363_v150_: int
            out28_: bool
            out29_: _dafny.Array
            out30_: int
            out31_: int
            out28_, out29_, out30_, out31_ = default__.m0((d_134_v0_ if (d_263_v94_).f14 else 524), d_143_v9_, d_148_globalState_)
            d_360_v147_ = out28_
            d_361_v148_ = out29_
            d_362_v149_ = out30_
            d_363_v150_ = out31_
        elif True:
            d_364___mcc_h23_ = source5_.cf7
            d_365_cf7_ = d_364___mcc_h23_
            d_263_v94_ = d_263_v94_
            index67_ = default__.safeIndex(615, (d_258_v91_).length(0))
            (d_258_v91_)[index67_] = True
            d_366_v151_: _dafny.Set
            d_366_v151_ = _dafny.Set({d_365_cf7_})
            d_367_v153_: _dafny.Seq
            d_367_v153_ = _dafny.SeqWithoutIsStrInference([d_366_v151_, d_366_v151_])
            d_368_v155_: _dafny.Array
            nw62_ = _dafny.Array(None, 20)
            nw62_[int(0)] = d_366_v151_
            nw62_[int(1)] = d_366_v151_
            nw62_[int(2)] = d_366_v151_
            nw62_[int(3)] = d_366_v151_
            nw62_[int(4)] = _dafny.Set({d_365_cf7_})
            nw62_[int(5)] = default__.fm14(456, d_365_cf7_, (d_263_v94_).f15, (d_263_v94_).f15, d_148_globalState_)
            nw62_[int(6)] = d_366_v151_
            def iife24_():
                def iife26_():
                    coll6_ = _dafny.Map()
                    compr_6_: str
                    for compr_6_ in ((d_367_v153_)[default__.safeIndex((d_263_v94_).f15, len(d_367_v153_))]).Elements:
                        d_369_v152_: str = compr_6_
                        if (d_369_v152_) in ((d_367_v153_)[default__.safeIndex((d_263_v94_).f15, len(d_367_v153_))]):
                            coll6_[d_369_v152_] = d_365_cf7_
                    return _dafny.Map(coll6_)
                coll4_ = _dafny.Set()
                def iife25_():
                    coll5_ = _dafny.Map()
                    compr_5_: str
                    for compr_5_ in ((d_367_v153_)[default__.safeIndex((d_263_v94_).f15, len(d_367_v153_))]).Elements:
                        d_369_v152_: str = compr_5_
                        if (d_369_v152_) in ((d_367_v153_)[default__.safeIndex((d_263_v94_).f15, len(d_367_v153_))]):
                            coll5_[d_369_v152_] = d_365_cf7_
                    return _dafny.Map(coll5_)
                compr_4_: str
                for compr_4_ in (iife25_()
                ).keys.Elements:
                    d_370_v154_: str = compr_4_
                    if (d_370_v154_) in (iife26_()
                    ):
                        coll4_ = coll4_.union(_dafny.Set([d_370_v154_]))
                return _dafny.Set(coll4_)
            nw62_[int(7)] = iife24_()
            
            nw62_[int(8)] = d_366_v151_
            nw62_[int(9)] = _dafny.Set({default__.fm15(_dafny.CodePoint('k'), _dafny.Map({(d_263_v94_).f15: d_134_v0_}), d_148_globalState_)})
            nw62_[int(10)] = (d_366_v151_ if (d_261_v92_).f16 else d_366_v151_)
            nw62_[int(11)] = (d_366_v151_ if not(False) else d_366_v151_)
            nw62_[int(12)] = d_366_v151_
            nw62_[int(13)] = d_366_v151_
            nw62_[int(14)] = d_366_v151_
            nw62_[int(15)] = (d_366_v151_) | (d_366_v151_)
            nw62_[int(16)] = d_366_v151_
            nw62_[int(17)] = _dafny.Set({d_265_v96_, d_365_cf7_})
            nw62_[int(18)] = default__.fm14(d_134_v0_, d_365_cf7_, (d_141_v7_).cardinality, 42, d_148_globalState_)
            nw62_[int(19)] = (d_366_v151_) | (d_366_v151_)
            d_368_v155_ = nw62_
            d_371_v156_: _dafny.Map
            d_371_v156_ = _dafny.Map({d_143_v9_: True})
            d_372_v157_: _dafny.Seq
            d_372_v157_ = _dafny.SeqWithoutIsStrInference([(d_262_v93_)[default__.safeIndex(538, (d_262_v93_).length(0))], d_261_v92_])
            d_373_v158_: T0
            nw63_ = C1()
            nw63_.ctor__(default__.fm6((d_263_v94_).f14, (d_263_v94_).f15, d_148_globalState_), len(_dafny.SeqWithoutIsStrInference([819, len(d_371_v156_), len(d_372_v157_), default__.fm0(d_145_v11_, not(False), (d_263_v94_).f14, d_148_globalState_)])))
            d_373_v158_ = nw63_
            d_374_v159_: _dafny.Map
            d_374_v159_ = _dafny.Map({(d_263_v94_ if (d_263_v94_).f14 else d_263_v94_): d_373_v158_})
            d_375_v160_: _dafny.Seq
            d_375_v160_ = _dafny.SeqWithoutIsStrInference([d_368_v155_])
            index68_ = default__.safeIndex(615, (d_258_v91_).length(0))
            rhs69_ = (_dafny.SeqWithoutIsStrInference([(d_263_v94_).f14])) + (d_143_v9_)
            rhs70_ = (d_143_v9_)[default__.safeIndex(d_134_v0_, len(d_143_v9_))]
            rhs71_ = (d_261_v92_).f16
            rhs72_ = len(d_374_v159_)
            rhs73_ = (d_375_v160_)[default__.safeIndex(d_134_v0_, len(d_375_v160_))]
            lhs38_ = d_258_v91_
            lhs39_ = default__.safeIndex(615, (d_258_v91_).length(0))
            d_143_v9_ = rhs69_
            lhs38_[lhs39_] = rhs70_
            d_140_v6_ = rhs71_
            d_134_v0_ = rhs72_
            d_368_v155_ = rhs73_
            d_376_v161_: _dafny.Set
            d_376_v161_ = _dafny.Set({(d_261_v92_).f16, d_140_v6_, (d_258_v91_)[default__.safeIndex(615, (d_258_v91_).length(0))], (d_261_v92_).f16, d_140_v6_})
            d_140_v6_ = default__.fm4((d_140_v6_) or ((d_261_v92_).f16), d_376_v161_, default__.safeModuloInt(d_134_v0_, -752), d_148_globalState_)
            d_377_v162_: _dafny.MultiSet
            d_377_v162_ = _dafny.MultiSet([len(_dafny.Set({(d_263_v94_).f14, not((d_258_v91_)[default__.safeIndex(615, (d_258_v91_).length(0))]), (d_263_v94_).f14, (d_261_v92_).f16, (d_258_v91_)[default__.safeIndex(615, (d_258_v91_).length(0))]})), d_134_v0_])
            d_378_v163_: _dafny.Map
            d_378_v163_ = _dafny.Map({d_365_cf7_: (d_263_v94_).f14})
            d_379_v164_: _dafny.Seq
            d_379_v164_ = _dafny.SeqWithoutIsStrInference([d_377_v162_])
            d_380_v165_: _dafny.Array
            nw64_ = _dafny.Array(None, 13)
            nw64_[int(0)] = d_377_v162_
            nw64_[int(1)] = (default__.fm16(d_140_v6_, ((d_149_v14_)[len(d_378_v163_)] if (len(d_378_v163_)) in (d_149_v14_) else ((d_145_v11_)[(d_261_v92_).f16] if ((d_261_v92_).f16) in (d_145_v11_) else (d_263_v94_).f14)), d_148_globalState_)) | (d_377_v162_)
            nw64_[int(2)] = (d_377_v162_) - (_dafny.MultiSet([-260]))
            nw64_[int(3)] = default__.fm16((d_263_v94_).f14, (d_263_v94_).f14, d_148_globalState_)
            nw64_[int(4)] = (d_377_v162_) - (_dafny.MultiSet([-871]))
            nw64_[int(5)] = (d_379_v164_)[default__.safeIndex((d_263_v94_).f15, len(d_379_v164_))]
            nw64_[int(6)] = _dafny.MultiSet([default__.fm0(d_145_v11_, (d_261_v92_).f16, (d_261_v92_).f16, d_148_globalState_)])
            nw64_[int(7)] = (d_377_v162_) - (d_377_v162_)
            nw64_[int(8)] = (_dafny.MultiSet([(d_263_v94_).f15])).intersection(d_377_v162_)
            nw64_[int(9)] = (d_379_v164_)[default__.safeIndex(d_134_v0_, len(d_379_v164_))]
            nw64_[int(10)] = _dafny.MultiSet([d_134_v0_])
            nw64_[int(11)] = (d_377_v162_) | (d_377_v162_)
            nw64_[int(12)] = _dafny.MultiSet([len(d_143_v9_), d_134_v0_])
            d_380_v165_ = nw64_
            index69_ = default__.safeIndex(787, (d_380_v165_).length(0))
            (d_380_v165_)[index69_] = d_377_v162_
            index70_ = default__.safeIndex(787, (d_380_v165_).length(0))
            (d_380_v165_)[index70_] = (d_377_v162_).intersection(d_377_v162_)
        d_381_i9_: int
        d_381_i9_ = 0
        with _dafny.label("2"):
            while not(d_140_v6_):
                with _dafny.c_label("2"):
                    if (d_381_i9_) >= (100):
                        raise _dafny.Break("2")
                    d_381_i9_ = (d_381_i9_) + (1)
                    d_382_v166_: D3
                    d_382_v166_ = D3_DC9(d_265_v96_)
                    d_383_v167_: D3
                    d_383_v167_ = D3_DC10(d_382_v166_)
                    d_384_v168_: _dafny.Set
                    d_384_v168_ = _dafny.Set({D3_DC10(d_382_v166_), d_383_v167_})
                    d_385_v169_: _dafny.Seq
                    d_385_v169_ = _dafny.SeqWithoutIsStrInference([d_384_v168_, d_384_v168_, d_384_v168_])
                    index71_ = default__.safeIndex(749, (d_147_v13_).length(0))
                    (d_147_v13_)[index71_] = len((d_385_v169_)[default__.safeIndex(((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))] if True else (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]), len(d_385_v169_))])
                    d_386_v170_: _dafny.Array
                    nw65_ = _dafny.Array(_dafny.Array(None, 0), 19)
                    d_386_v170_ = nw65_
                    d_387_v171_: D5
                    d_387_v171_ = D5_DC14(d_386_v170_)
                    d_386_v170_ = ((d_387_v171_ if d_140_v6_ else d_387_v171_)).cf22
                    d_388_v172_: C1
                    nw66_ = C1()
                    nw66_.ctor__(d_140_v6_, (d_263_v94_).f15)
                    d_388_v172_ = nw66_
                    d_389_v173_: _dafny.Map
                    d_389_v173_ = _dafny.Map({default__.fm17(d_140_v6_, d_141_v7_, len((d_144_v10_).set(default__.safeIndex((d_388_v172_).f15, len(d_144_v10_)), d_265_v96_)), d_144_v10_, d_148_globalState_): d_258_v91_})
                    d_390_v174_: _dafny.Seq
                    d_390_v174_ = _dafny.SeqWithoutIsStrInference([d_138_v4_, _dafny.SeqWithoutIsStrInference([(d_263_v94_).f15, (0) - (len(d_144_v10_))])])
                    d_389_v173_ = (d_389_v173_).set(d_390_v174_, d_258_v91_)
                    pass
            pass
        d_391_i10_: int
        d_391_i10_ = 0
        with _dafny.label("3"):
            while (d_140_v6_) and (False):
                with _dafny.c_label("3"):
                    if (d_391_i10_) >= (100):
                        raise _dafny.Break("3")
                    d_391_i10_ = (d_391_i10_) + (1)
                    d_392_v175_: _dafny.Map
                    d_392_v175_ = _dafny.Map({(d_263_v94_).f14: d_134_v0_})
                    d_392_v175_ = ((d_392_v175_) | ((d_392_v175_).set(True, (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]))) | (d_392_v175_)
                    (d_261_v92_).m1(d_148_globalState_)
                    d_393_v176_: C0
                    nw67_ = C0()
                    nw67_.ctor__((d_261_v92_).f16)
                    d_393_v176_ = nw67_
                    d_394_v178_: D0
                    def iife27_():
                        coll7_ = _dafny.Set()
                        compr_7_: _dafny.Set
                        for compr_7_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({(d_261_v92_).f16, (d_261_v92_).f16, not((d_261_v92_).f16)}) for d_395_i11_ in range(default__.abs(-916))])).Elements:
                            d_396_v177_: _dafny.Set = compr_7_
                            if (d_396_v177_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({(d_261_v92_).f16, (d_261_v92_).f16, not((d_261_v92_).f16)}) for d_395_i11_ in range(default__.abs(-916))])):
                                coll7_ = coll7_.union(_dafny.Set([d_396_v177_]))
                        return _dafny.Set(coll7_)
                    d_394_v178_ = D0_DC1(iife27_()
, (d_263_v94_).f14)
                    index72_ = default__.safeIndex(360, (d_258_v91_).length(0))
                    (d_258_v91_)[index72_] = (d_394_v178_).cf3
                    index73_ = default__.safeIndex(360, (d_258_v91_).length(0))
                    (d_258_v91_)[index73_] = ((d_263_v94_).f14) not in (d_392_v175_)
                    pass
            pass
        d_397_v179_: _dafny.Array
        nw68_ = _dafny.Array(_dafny.Map({}), 28)
        d_397_v179_ = nw68_
        index74_ = default__.safeIndex(534, (d_397_v179_).length(0))
        def iife28_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in (d_138_v4_).Elements:
                d_398_v180_: int = compr_8_
                if (d_398_v180_) in (d_138_v4_):
                    coll8_[(d_398_v180_) + ((d_263_v94_).f15)] = (d_263_v94_).f14
            return _dafny.Map(coll8_)
        (d_397_v179_)[index74_] = iife28_()
        
        index75_ = default__.safeIndex(534, (d_397_v179_).length(0))
        (d_397_v179_)[index75_] = d_149_v14_
        (d_261_v92_).m1(d_148_globalState_)
        d_399_i12_: int
        d_399_i12_ = 0
        with _dafny.label("4"):
            while ((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]) != (len((d_144_v10_) + ((d_144_v10_).set(default__.safeIndex((0) - ((0) - ((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))])), len(d_144_v10_)), d_265_v96_)))):
                with _dafny.c_label("4"):
                    if (d_399_i12_) >= (100):
                        raise _dafny.Break("4")
                    d_399_i12_ = (d_399_i12_) + (1)
                    d_400_v181_: T0
                    nw69_ = C1()
                    nw69_.ctor__(d_140_v6_, (d_263_v94_).f15)
                    d_400_v181_ = nw69_
                    d_401_v182_: _dafny.Map
                    d_401_v182_ = _dafny.Map({784: d_400_v181_})
                    d_402_v183_: D3
                    d_402_v183_ = D3_DC8(d_263_v94_)
                    d_403_v184_: _dafny.Array
                    nw70_ = _dafny.Array(None, 3)
                    nw70_[int(0)] = d_402_v183_
                    nw70_[int(1)] = d_402_v183_
                    nw70_[int(2)] = d_402_v183_
                    d_403_v184_ = nw70_
                    d_404_v185_: _dafny.Map
                    d_404_v185_ = _dafny.Map({((d_401_v182_)[d_134_v0_] if (d_134_v0_) in (d_401_v182_) else d_400_v181_): d_403_v184_})
                    d_405_v186_: _dafny.Map
                    d_405_v186_ = _dafny.Map({d_403_v184_: d_147_v13_})
                    d_140_v6_ = (((d_404_v185_)[d_400_v181_] if (d_400_v181_) in (d_404_v185_) else d_403_v184_)) in ((d_405_v186_) | (_dafny.Map({d_403_v184_: d_147_v13_})))
                    if (d_263_v94_).f14:
                        d_140_v6_ = (d_263_v94_).f14
                        d_140_v6_ = (d_261_v92_).f16
                        d_406_v187_: C1
                        nw71_ = C1()
                        nw71_.ctor__((d_261_v92_).f16, (d_263_v94_).f15)
                        d_406_v187_ = nw71_
                        d_140_v6_ = default__.fm6((d_263_v94_).f14, (d_263_v94_).f15, d_148_globalState_)
                        d_407_v188_: _dafny.Set
                        d_407_v188_ = _dafny.Set({(d_263_v94_).f14})
                        d_408_v189_: _dafny.MultiSet
                        d_408_v189_ = _dafny.MultiSet([d_407_v188_])
                        index76_ = default__.safeIndex(171, (d_258_v91_).length(0))
                        (d_258_v91_)[index76_] = (d_408_v189_).ispropersubset(d_408_v189_)
                        index77_ = default__.safeIndex(171, (d_258_v91_).length(0))
                        (d_258_v91_)[index77_] = d_140_v6_
                    elif True:
                        (d_148_globalState_).f4 = d_265_v96_
                        d_147_v13_ = d_147_v13_
                        d_409_v190_: _dafny.Seq
                        d_409_v190_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([(d_263_v94_).f14])).set(default__.safeIndex((d_138_v4_)[default__.safeIndex((d_263_v94_).f15, len(d_138_v4_))], len(_dafny.SeqWithoutIsStrInference([(d_263_v94_).f14]))), d_140_v6_), d_143_v9_, (d_143_v9_) + (d_143_v9_), d_143_v9_])
                        d_409_v190_ = (d_409_v190_) + (d_409_v190_)
                        index78_ = default__.safeIndex(864, (d_258_v91_).length(0))
                        (d_258_v91_)[index78_] = (d_263_v94_).f14
                        d_410_v191_: _dafny.Set
                        d_410_v191_ = _dafny.Set({d_140_v6_, (d_263_v94_).f14})
                        index79_ = default__.safeIndex(864, (d_258_v91_).length(0))
                        (d_258_v91_)[index79_] = (((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]) > ((d_263_v94_).f15) if d_140_v6_ else default__.fm4((d_263_v94_).f14, d_410_v191_, d_134_v0_, d_148_globalState_))
                        d_401_v182_ = (d_401_v182_).set((d_134_v0_) + ((d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]), d_400_v181_)
                    if (d_261_v92_).f16:
                        d_411_v192_: _dafny.Seq
                        d_411_v192_ = _dafny.SeqWithoutIsStrInference([d_400_v181_, d_400_v181_, d_400_v181_, d_400_v181_, d_400_v181_])
                        d_412_v193_: _dafny.Array
                        nw72_ = _dafny.Array(None, 1)
                        nw72_[int(0)] = d_411_v192_
                        d_412_v193_ = nw72_
                        index80_ = default__.safeIndex(941, (d_412_v193_).length(0))
                        (d_412_v193_)[index80_] = d_411_v192_
                        index81_ = default__.safeIndex(941, (d_412_v193_).length(0))
                        (d_412_v193_)[index81_] = d_411_v192_
                        d_144_v10_ = (d_144_v10_) + (d_144_v10_)
                        index82_ = default__.safeIndex(749, (d_147_v13_).length(0))
                        (d_147_v13_)[index82_] = ((d_263_v94_).f15) * (len((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nadw"))): (d_263_v94_).f14})) | ((d_397_v179_)[default__.safeIndex(534, (d_397_v179_).length(0))])))
                        d_413_v194_: _dafny.Array
                        nw73_ = _dafny.Array(None, 9)
                        nw73_[int(0)] = d_400_v181_
                        nw73_[int(1)] = d_400_v181_
                        nw73_[int(2)] = d_400_v181_
                        nw73_[int(3)] = d_400_v181_
                        nw73_[int(4)] = d_400_v181_
                        nw73_[int(5)] = d_400_v181_
                        nw73_[int(6)] = d_400_v181_
                        nw73_[int(7)] = d_400_v181_
                        nw73_[int(8)] = d_400_v181_
                        d_413_v194_ = nw73_
                        d_413_v194_ = d_413_v194_
                        d_414_v195_: C1
                        nw74_ = C1()
                        nw74_.ctor__((d_261_v92_).f16, ((d_263_v94_).f15 if False else (0) - ((d_400_v181_).fm1(d_144_v10_, -329, d_134_v0_, d_148_globalState_))))
                        d_414_v195_ = nw74_
                    elif True:
                        d_415_v196_: _dafny.Set
                        d_415_v196_ = _dafny.Set({d_400_v181_, d_400_v181_})
                        d_416_v197_: _dafny.Set
                        d_416_v197_ = _dafny.Set({(d_415_v196_)})
                        d_417_v198_: _dafny.Set
                        d_417_v198_ = _dafny.Set({d_400_v181_})
                        d_418_v199_: _dafny.Seq
                        d_418_v199_ = _dafny.SeqWithoutIsStrInference([d_416_v197_])
                        d_419_v200_: _dafny.Array
                        nw75_ = _dafny.Array(None, 25)
                        nw75_[int(0)] = d_416_v197_
                        nw75_[int(1)] = d_416_v197_
                        nw75_[int(2)] = d_416_v197_
                        nw75_[int(3)] = d_416_v197_
                        nw75_[int(4)] = d_416_v197_
                        nw75_[int(5)] = d_416_v197_
                        nw75_[int(6)] = (d_416_v197_).intersection(d_416_v197_)
                        nw75_[int(7)] = d_416_v197_
                        nw75_[int(8)] = d_416_v197_
                        nw75_[int(9)] = (d_416_v197_).intersection(d_416_v197_)
                        nw75_[int(10)] = (d_416_v197_) - (_dafny.Set({d_417_v198_}))
                        nw75_[int(11)] = d_416_v197_
                        nw75_[int(12)] = (d_416_v197_) - (d_416_v197_)
                        nw75_[int(13)] = d_416_v197_
                        nw75_[int(14)] = d_416_v197_
                        nw75_[int(15)] = d_416_v197_
                        nw75_[int(16)] = (d_416_v197_) - (d_416_v197_)
                        nw75_[int(17)] = d_416_v197_
                        nw75_[int(18)] = d_416_v197_
                        nw75_[int(19)] = d_416_v197_
                        nw75_[int(20)] = d_416_v197_
                        nw75_[int(21)] = d_416_v197_
                        nw75_[int(22)] = d_416_v197_
                        nw75_[int(23)] = (d_418_v199_)[default__.safeIndex((d_263_v94_).f15, len(d_418_v199_))]
                        nw75_[int(24)] = d_416_v197_
                        d_419_v200_ = nw75_
                        index83_ = default__.safeIndex(982, (d_419_v200_).length(0))
                        (d_419_v200_)[index83_] = d_416_v197_
                        index84_ = default__.safeIndex(982, (d_419_v200_).length(0))
                        (d_419_v200_)[index84_] = d_416_v197_
                        d_140_v6_ = (d_261_v92_).f16
                        (d_261_v92_).m1(d_148_globalState_)
                        d_420_v201_: D0
                        d_420_v201_ = D0_DC0((d_263_v94_).f15, (d_263_v94_).f14)
                        d_421_v202_: _dafny.Map
                        d_421_v202_ = _dafny.Map({(d_263_v94_).f14: _dafny.Map({(d_263_v94_).f15: (d_261_v92_).f16})})
                        d_420_v201_ = D0_DC0(len(((d_421_v202_).set(False, d_149_v14_)) | (d_421_v202_)), not ((d_261_v92_).f16) or (d_140_v6_))
                        index85_ = default__.safeIndex(749, (d_147_v13_).length(0))
                        (d_147_v13_)[index85_] = (d_147_v13_)[default__.safeIndex(749, (d_147_v13_).length(0))]
                    d_422_v203_: _dafny.Set
                    d_422_v203_ = _dafny.Set({d_258_v91_, d_258_v91_})
                    d_422_v203_ = (d_422_v203_).intersection(d_422_v203_)
                    pass
            pass
        d_423_v204_: D1
        d_423_v204_ = D1_DC5(d_140_v6_)
        pat_let_tv14_ = d_263_v94_
        pat_let_tv15_ = d_145_v11_
        pat_let_tv16_ = d_145_v11_
        pat_let_tv17_ = d_263_v94_
        pat_let_tv18_ = d_263_v94_
        def lambda15_(source7_):
            if source7_.is_DC5:
                d_424___mcc_h29_ = source7_.cf9
                d_425_cf9_ = d_424___mcc_h29_
                return d_425_cf9_
            elif True:
                d_426___mcc_h30_ = source7_.cf8
                d_427_cf8_ = d_426___mcc_h30_
                if ((pat_let_tv14_).f14) in (pat_let_tv15_):
                    return (pat_let_tv16_)[(pat_let_tv17_).f14]
                elif True:
                    return (pat_let_tv18_).f14

        d_140_v6_ = lambda15_(d_423_v204_)
        _dafny.print(_dafny.string_of(d_134_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v1_) == (_dafny.MultiSet([_dafny.MultiSet([-250])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v2_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_137_v3_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v4_) == (_dafny.SeqWithoutIsStrInference([-250, -250, -250]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_139_v5_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([-250])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_140_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_141_v7_) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_v8_) == (_dafny.MultiSet([_dafny.MultiSet([False, False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v9_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_144_v10_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_v11_) == (_dafny.Map({False: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_v12_) == (_dafny.Map({-250: 2, 0: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v13_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v13_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v13_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v13_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v13_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v13_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v13_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v13_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_148_globalState_).f2) == (_dafny.MultiSet([_dafny.MultiSet([-250])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_148_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_148_globalState_.f5)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_.f6) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([-250])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_148_globalState_).f8) == (_dafny.MultiSet([_dafny.MultiSet([False, False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_148_globalState_).f11) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_.f12)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_.f12)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_.f12)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_.f12)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_.f12)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_.f12)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_.f12)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_.f12)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v14_) == (_dafny.Map({84: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v15_) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_v16_).cf4) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v16_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v16_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_249_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v91_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v91_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v91_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v91_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v91_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v91_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v91_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v91_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v92_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v93_)[0]).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v93_)[1]).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v94_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v94_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_264_v95_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_265_v96_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v97_) == (_dafny.Map({2: _dafny.CodePoint('c')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v126_) == (_dafny.MultiSet([]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_381_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_391_i10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_397_v179_)[2]) == (_dafny.Map({84: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_399_i12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_423_v204_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC5(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({self.cf8.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC9(D3, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(int(0), _dafny.CodePoint('D'), _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC12(D4, NamedTuple('DC12', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC15(_dafny.Set({}), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC15(D5, NamedTuple('DC15', [('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC17(D6, NamedTuple('DC17', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC19(D7, NamedTuple('DC19', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(_dafny.Seq({}), _dafny.CodePoint('D'), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)

class D8_DC21(D8, NamedTuple('DC21', [('cf29', Any), ('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC25(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)

class D9_DC25(D9, NamedTuple('DC25', [('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f4: str = _dafny.CodePoint('D')
        self.f5: _dafny.Set = _dafny.Set({})
        self.f6: _dafny.Set = _dafny.Set({})
        self.f12: _dafny.Array = _dafny.Array(None, 0)
        self._f0: int = int(0)
        self._f1: int = int(0)
        self._f2: _dafny.MultiSet = _dafny.MultiSet({})
        self._f3: bool = False
        self._f7: str = _dafny.CodePoint('D')
        self._f8: _dafny.MultiSet = _dafny.MultiSet({})
        self._f9: bool = False
        self._f10: int = int(0)
        self._f11: _dafny.Seq = _dafny.Seq({})
        self._f13: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self).f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f13 = f13

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
    @property
    def f13(self):
        return self._f13

class C0(T0):
    def  __init__(self):
        self._f16: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f16):
        (self)._f16 = f16

    def fm1(self, p0, p1, p2, globalState):
        return (_dafny.MultiSet([(self).f16, not((self).f16)])).cardinality

    def fm2(self, p0, p1, p2, globalState):
        def iife29_():
            coll9_ = _dafny.Map()
            compr_9_: str
            for compr_9_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('m'), _dafny.CodePoint('q'), _dafny.CodePoint('g')])).Elements:
                d_428_v0_: str = compr_9_
                if (d_428_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('m'), _dafny.CodePoint('q'), _dafny.CodePoint('g')])):
                    coll9_[d_428_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmqlcma"))
            return _dafny.Map(coll9_)
        return (iife29_()
        ) | ((_dafny.Map({_dafny.CodePoint('r'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qky"))})) | (_dafny.Map({_dafny.CodePoint('l'): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_429_i0_ in range(default__.abs(467))])})))

    def fm5(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tv"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usgh")) if (self).f16 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjrsoedx"))))

    def m1(self, globalState):
        d_430_v0_: _dafny.Array
        def lambda16_(d_431_i0_):
            return (self).f16

        init7_ = lambda16_
        nw76_ = _dafny.Array(None, 4)
        for i0_7_ in range(nw76_.length(0)):
            nw76_[i0_7_] = init7_(i0_7_)
        d_430_v0_ = nw76_
        index86_ = default__.safeIndex(848, (d_430_v0_).length(0))
        (d_430_v0_)[index86_] = (self).f16
        d_432_v1_: D1
        d_432_v1_ = D1_DC5((self).f16)
        index87_ = default__.safeIndex(848, (d_430_v0_).length(0))
        def lambda17_(source8_):
            if source8_.is_DC5:
                d_433___mcc_h0_ = source8_.cf9
                d_434_cf9_ = d_433___mcc_h0_
                return (len(_dafny.Set({(self).f16}))) == (444)
            elif True:
                d_435___mcc_h1_ = source8_.cf8
                d_436_cf8_ = d_435___mcc_h1_
                return (self).f16

        (d_430_v0_)[index87_] = lambda17_(d_432_v1_)
        d_437_v2_: int
        d_437_v2_ = 96
        d_438_v3_: _dafny.Set
        d_438_v3_ = _dafny.Set({d_437_v2_, d_437_v2_})
        d_439_v4_: _dafny.Seq
        d_439_v4_ = _dafny.SeqWithoutIsStrInference([d_438_v3_])
        d_440_v5_: _dafny.Seq
        d_440_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oobutohe"))
        d_441_v6_: _dafny.Seq
        d_441_v6_ = _dafny.SeqWithoutIsStrInference([(0) - ((0) - (d_437_v2_))])
        d_442_i1_: int
        d_442_i1_ = 0
        with _dafny.label("5"):
            while (len(((d_439_v4_)[default__.safeIndex(len(d_440_v5_), len(d_439_v4_))]) | (d_438_v3_))) <= (len(d_441_v6_)):
                with _dafny.c_label("5"):
                    if (d_442_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_442_i1_ = (d_442_i1_) + (1)
                    d_443_v7_: _dafny.Array
                    def lambda18_(d_444_v0_):
                        def lambda19_(d_445_i2_):
                            return (_dafny.SeqWithoutIsStrInference([(self).f16, (self).f16])) + (_dafny.SeqWithoutIsStrInference([(d_444_v0_)[default__.safeIndex(848, (d_444_v0_).length(0))]]))

                        return lambda19_

                    init8_ = lambda18_(d_430_v0_)
                    nw77_ = _dafny.Array(None, 14)
                    for i0_8_ in range(nw77_.length(0)):
                        nw77_[i0_8_] = init8_(i0_8_)
                    d_443_v7_ = nw77_
                    d_446_v8_: _dafny.Seq
                    d_446_v8_ = _dafny.SeqWithoutIsStrInference([(self).f16, False, (d_430_v0_)[default__.safeIndex(848, (d_430_v0_).length(0))], not((d_430_v0_)[default__.safeIndex(848, (d_430_v0_).length(0))]), default__.fm6((d_430_v0_)[default__.safeIndex(848, (d_430_v0_).length(0))], d_437_v2_, globalState)])
                    index88_ = default__.safeIndex(22, (d_443_v7_).length(0))
                    (d_443_v7_)[index88_] = (d_446_v8_) + (d_446_v8_)
                    index89_ = default__.safeIndex(22, (d_443_v7_).length(0))
                    (d_443_v7_)[index89_] = (d_446_v8_).set(default__.safeIndex(d_437_v2_, len(d_446_v8_)), (d_430_v0_)[default__.safeIndex(848, (d_430_v0_).length(0))])
                    d_447_v9_: bool
                    d_448_v10_: _dafny.Array
                    d_449_v11_: int
                    d_450_v12_: int
                    out32_: bool
                    out33_: _dafny.Array
                    out34_: int
                    out35_: int
                    out32_, out33_, out34_, out35_ = default__.m0((0) - ((d_437_v2_) * (d_437_v2_)), (d_443_v7_)[default__.safeIndex(22, (d_443_v7_).length(0))], globalState)
                    d_447_v9_ = out32_
                    d_448_v10_ = out33_
                    d_449_v11_ = out34_
                    d_450_v12_ = out35_
                    d_451_v13_: _dafny.Array
                    nw78_ = _dafny.Array(int(0), 4)
                    d_451_v13_ = nw78_
                    d_452_v14_: str
                    d_452_v14_ = _dafny.CodePoint('f')
                    d_453_v15_: _dafny.Map
                    d_453_v15_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wluaxmjb")): d_452_v14_})
                    index90_ = default__.safeIndex(462, (d_451_v13_).length(0))
                    (d_451_v13_)[index90_] = (d_437_v2_) * (len(d_453_v15_))
                    index91_ = default__.safeIndex(462, (d_451_v13_).length(0))
                    (d_451_v13_)[index91_] = default__.safeModuloInt(d_437_v2_, len((d_441_v6_) + (d_441_v6_)))
                    d_454_v16_: _dafny.MultiSet
                    d_454_v16_ = _dafny.MultiSet([len(d_440_v5_), (d_451_v13_)[default__.safeIndex(462, (d_451_v13_).length(0))], 876])
                    pat_let_tv19_ = d_454_v16_
                    def iife30_(_pat_let10_0):
                        def iife31_(d_455_dt__update__tmp_h0_):
                            def iife32_(_pat_let11_0):
                                def iife33_(d_456_dt__update_hcf5_h0_):
                                    return D0_DC2((d_455_dt__update__tmp_h0_).cf4, d_456_dt__update_hcf5_h0_, (d_455_dt__update__tmp_h0_).cf6)
                                return iife33_(_pat_let11_0)
                            return iife32_((pat_let_tv19_).cardinality)
                        return iife31_(_pat_let10_0)
                    source9_ = iife30_(default__.fm7(globalState))
                    if source9_.is_DC0:
                        d_457___mcc_h2_ = source9_.cf0
                        d_458___mcc_h3_ = source9_.cf1
                        d_459_cf1_ = d_458___mcc_h3_
                        d_460_cf0_ = d_457___mcc_h2_
                        index92_ = default__.safeIndex(848, (d_430_v0_).length(0))
                        (d_430_v0_)[index92_] = d_447_v9_
                        d_461_v17_: _dafny.Map
                        d_461_v17_ = _dafny.Map({d_451_v13_: not(d_447_v9_)})
                        d_461_v17_ = (d_461_v17_).set(d_451_v13_, default__.fm6(d_447_v9_, len(d_438_v3_), globalState))
                        d_462_v18_: D0
                        d_462_v18_ = D0_DC0(len(_dafny.Set({d_437_v2_, 682})), d_447_v9_)
                        d_459_cf1_ = not(not(not(not((d_462_v18_).cf1))))
                        d_463_v19_: bool
                        d_464_v20_: _dafny.Array
                        d_465_v21_: int
                        d_466_v22_: int
                        out36_: bool
                        out37_: _dafny.Array
                        out38_: int
                        out39_: int
                        out36_, out37_, out38_, out39_ = default__.m0(((d_451_v13_)[default__.safeIndex(462, (d_451_v13_).length(0))]) - (d_460_cf0_), _dafny.SeqWithoutIsStrInference([(d_430_v0_)[default__.safeIndex(848, (d_430_v0_).length(0))]]), globalState)
                        d_463_v19_ = out36_
                        d_464_v20_ = out37_
                        d_465_v21_ = out38_
                        d_466_v22_ = out39_
                    elif source9_.is_DC1:
                        d_467___mcc_h4_ = source9_.cf2
                        d_468___mcc_h5_ = source9_.cf3
                        d_469_cf3_ = d_468___mcc_h5_
                        d_470_cf2_ = d_467___mcc_h4_
                        d_437_v2_ = 868
                        d_471_v23_: _dafny.Array
                        def lambda20_(d_472_v14_, d_473_v0_):
                            def lambda21_(d_474_i3_):
                                return _dafny.Map({d_472_v14_: (d_473_v0_)[default__.safeIndex(848, (d_473_v0_).length(0))]})

                            return lambda21_

                        init9_ = lambda20_(d_452_v14_, d_430_v0_)
                        nw79_ = _dafny.Array(None, 7)
                        for i0_9_ in range(nw79_.length(0)):
                            nw79_[i0_9_] = init9_(i0_9_)
                        d_471_v23_ = nw79_
                        d_471_v23_ = d_471_v23_
                        d_475_v24_: D0
                        d_475_v24_ = D0_DC3(d_452_v14_)
                        d_475_v24_ = d_475_v24_
                        (globalState).f12 = d_451_v13_
                    elif source9_.is_DC2:
                        d_476___mcc_h6_ = source9_.cf4
                        d_477___mcc_h7_ = source9_.cf5
                        d_478___mcc_h8_ = source9_.cf6
                        d_479_cf6_ = d_478___mcc_h8_
                        d_480_cf5_ = d_477___mcc_h7_
                        d_481_cf4_ = d_476___mcc_h6_
                        d_440_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qf"))
                        index93_ = default__.safeIndex(462, (d_451_v13_).length(0))
                        (d_451_v13_)[index93_] = d_449_v11_
                        d_482_v25_: _dafny.Array
                        nw80_ = _dafny.Array(None, 2)
                        nw80_[int(0)] = d_452_v14_
                        nw80_[int(1)] = d_452_v14_
                        d_482_v25_ = nw80_
                        index94_ = default__.safeIndex(885, (d_482_v25_).length(0))
                        (d_482_v25_)[index94_] = d_452_v14_
                        index95_ = default__.safeIndex(885, (d_482_v25_).length(0))
                        (d_482_v25_)[index95_] = d_452_v14_
                        d_483_v26_: _dafny.Map
                        d_483_v26_ = _dafny.Map({(self).f16: _dafny.Set({len(d_440_v5_), d_437_v2_, d_480_cf5_})})
                        d_484_v27_: D0
                        d_484_v27_ = D0_DC2(d_481_cf4_, len(d_446_v8_), d_447_v9_)
                        d_485_v28_: D0
                        d_485_v28_ = D0_DC2(((d_483_v26_)[d_447_v9_] if (d_447_v9_) in (d_483_v26_) else d_438_v3_), (d_484_v27_).cf5, True)
                        d_485_v28_ = default__.fm7(globalState)
                    elif True:
                        d_486___mcc_h9_ = source9_.cf7
                        d_487_cf7_ = d_486___mcc_h9_
                        d_488_v29_: bool
                        d_489_v30_: _dafny.Array
                        d_490_v31_: int
                        d_491_v32_: int
                        out40_: bool
                        out41_: _dafny.Array
                        out42_: int
                        out43_: int
                        out40_, out41_, out42_, out43_ = default__.m0(d_449_v11_, _dafny.SeqWithoutIsStrInference([(d_430_v0_)[default__.safeIndex(848, (d_430_v0_).length(0))]]), globalState)
                        d_488_v29_ = out40_
                        d_489_v30_ = out41_
                        d_490_v31_ = out42_
                        d_491_v32_ = out43_
                        d_487_cf7_ = (d_440_v5_)[default__.safeIndex(104, len(d_440_v5_))]
                        d_492_v33_: _dafny.Map
                        d_492_v33_ = _dafny.Map({-902: ((self).f16 if not(d_447_v9_) else (self).f16)})
                        d_492_v33_ = (d_492_v33_) | (d_492_v33_)
                        d_430_v0_ = d_430_v0_
                    pass
            pass
        d_493_v34_: _dafny.Array
        nw81_ = _dafny.Array(int(0), 26)
        d_493_v34_ = nw81_
        index96_ = default__.safeIndex(332, (d_493_v34_).length(0))
        (d_493_v34_)[index96_] = d_437_v2_
        index97_ = default__.safeIndex(332, (d_493_v34_).length(0))
        (d_493_v34_)[index97_] = (self).fm1(d_440_v5_, d_437_v2_, (len(d_440_v5_) if (self).f16 else d_437_v2_), globalState)
        d_494_v35_: str
        d_494_v35_ = _dafny.CodePoint('f')
        d_495_v36_: _dafny.Map
        d_495_v36_ = _dafny.Map({(self).f16: d_494_v35_})
        (globalState).f4 = ((d_495_v36_)[(self).f16] if ((self).f16) in (d_495_v36_) else _dafny.CodePoint('v'))
        d_496_v37_: _dafny.Map
        d_496_v37_ = _dafny.Map({((d_493_v34_)[default__.safeIndex(332, (d_493_v34_).length(0))]) > ((d_493_v34_)[default__.safeIndex(332, (d_493_v34_).length(0))]): d_437_v2_})
        d_497_v38_: _dafny.Map
        d_497_v38_ = _dafny.Map({d_438_v3_: -717})
        d_498_v39_: D0
        d_498_v39_ = D0_DC2(d_438_v3_, (d_493_v34_)[default__.safeIndex(332, (d_493_v34_).length(0))], (d_430_v0_)[default__.safeIndex(848, (d_430_v0_).length(0))])
        d_499_v40_: _dafny.Map
        d_499_v40_ = _dafny.Map({d_437_v2_: (d_430_v0_)[default__.safeIndex(848, (d_430_v0_).length(0))]})
        d_496_v37_ = (d_496_v37_).set((self).f16, ((d_493_v34_)[default__.safeIndex(332, (d_493_v34_).length(0))] if (self).f16 else ((d_497_v38_)[(d_498_v39_).cf4] if ((d_498_v39_).cf4) in (d_497_v38_) else len(d_499_v40_))))
        index98_ = default__.safeIndex(332, (d_493_v34_).length(0))
        (d_493_v34_)[index98_] = (d_493_v34_)[default__.safeIndex(332, (d_493_v34_).length(0))]

    @property
    def f16(self):
        return self._f16

class C1(T0):
    def  __init__(self):
        self._f14: bool = False
        self._f15: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f14, f15):
        (self)._f14 = f14
        (self)._f15 = f15

    def fm1(self, p0, p1, p2, globalState):
        return (0) - (len(((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f15, (self).f15, len(_dafny.Map({(self).f15: (self).f14})), (self).f15])]) if (self).f14 else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f15, (self).f15])]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([512]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({(self).f14})), len(_dafny.Map({(self).f14: (self).f15}))])]))))

    def fm2(self, p0, p1, p2, globalState):
        def iife34_():
            coll10_ = _dafny.Map()
            compr_10_: str
            for compr_10_ in (_dafny.Map({_dafny.CodePoint('x'): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_500_i0_ in range(default__.abs(861))])})).keys.Elements:
                d_501_v0_: str = compr_10_
                if (d_501_v0_) in (_dafny.Map({_dafny.CodePoint('x'): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_500_i0_ in range(default__.abs(861))])})):
                    coll10_[d_501_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evyrhbeq"))
            return _dafny.Map(coll10_)
        return (iife34_()
        ) | ((_dafny.Map({_dafny.CodePoint('v'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycuwefvef"))})) | (_dafny.Map({_dafny.CodePoint('h'): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_502_i1_ in range(default__.abs(581))])})))

    def m1(self, globalState):
        d_503_v0_: _dafny.Seq
        d_503_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqptj"))
        d_504_v1_: _dafny.MultiSet
        d_504_v1_ = _dafny.MultiSet([d_503_v0_, (d_503_v0_) + (d_503_v0_), d_503_v0_])
        d_505_v2_: _dafny.Seq
        d_505_v2_ = _dafny.SeqWithoutIsStrInference([d_503_v0_, d_503_v0_, (default__.fm3((self).f15, globalState)).cf8])
        d_504_v1_ = ((_dafny.MultiSet(d_505_v2_)).intersection(d_504_v1_)) | (d_504_v1_)
        d_505_v2_ = d_505_v2_
        hi0_ = (self).f15
        for d_506_i0_ in range((self).f15, hi0_):
            d_507_v3_: _dafny.Set
            d_507_v3_ = _dafny.Set({(self).f14, (self).f14, (self).f14})
            d_508_v4_: _dafny.MultiSet
            d_508_v4_ = _dafny.MultiSet([len(d_503_v0_)])
            d_509_v5_: str
            d_509_v5_ = _dafny.CodePoint('r')
            d_510_v6_: _dafny.Map
            d_510_v6_ = _dafny.Map({(self).f14: d_509_v5_})
            d_511_v7_: _dafny.Map
            d_511_v7_ = _dafny.Map({(self).f14: (self).f14})
            d_512_v8_: _dafny.Seq
            d_512_v8_ = _dafny.SeqWithoutIsStrInference([True, (self).f14])
            d_513_v9_: _dafny.Array
            nw82_ = _dafny.Array(None, 28)
            nw82_[int(0)] = (d_506_i0_) < (d_506_i0_)
            nw82_[int(1)] = (self).f14
            nw82_[int(2)] = (self).f14
            nw82_[int(3)] = True
            nw82_[int(4)] = (D0_DC2(_dafny.Set({(self).f15}), 180, True)).cf6
            nw82_[int(5)] = (self).f14
            nw82_[int(6)] = (self).f14
            nw82_[int(7)] = (self).f14
            nw82_[int(8)] = (self).f14
            nw82_[int(9)] = (d_507_v3_).issubset(d_507_v3_)
            nw82_[int(10)] = (self).f14
            nw82_[int(11)] = (self).f14
            nw82_[int(12)] = (self).f14
            nw82_[int(13)] = (self).f14
            nw82_[int(14)] = (d_506_i0_) > (((d_508_v4_)[488] if (488) in (d_508_v4_) else len(d_510_v6_)))
            nw82_[int(15)] = not (not((self).f14)) or ((self).f14)
            nw82_[int(16)] = (self).f14
            nw82_[int(17)] = (self).f14
            nw82_[int(18)] = (d_506_i0_) < (len(d_503_v0_))
            nw82_[int(19)] = not((len(d_511_v7_)) != (759))
            nw82_[int(20)] = (self).f14
            nw82_[int(21)] = (d_508_v4_).ispropersubset((d_508_v4_).set(d_506_i0_, default__.abs((self).f15)))
            nw82_[int(22)] = not(((self).f14 if (self).f14 else (self).f14))
            nw82_[int(23)] = (d_512_v8_) != (d_512_v8_)
            nw82_[int(24)] = not((self).f14)
            nw82_[int(25)] = (d_509_v5_) in (d_503_v0_)
            nw82_[int(26)] = (d_512_v8_) <= (d_512_v8_)
            nw82_[int(27)] = (self).f14
            d_513_v9_ = nw82_
            index99_ = default__.safeIndex(466, (d_513_v9_).length(0))
            (d_513_v9_)[index99_] = ((self).f15) >= (d_506_i0_)
            index100_ = default__.safeIndex(466, (d_513_v9_).length(0))
            (d_513_v9_)[index100_] = ((d_511_v7_)[(d_506_i0_) == ((self).f15)] if ((d_506_i0_) == ((self).f15)) in (d_511_v7_) else ((self).f14) or ((self).f14))
            d_514_v11_: _dafny.Array
            def lambda22_(d_515_i0_, d_516_v3_):
                def lambda23_(d_517_i1_):
                    def iife35_():
                        coll11_ = _dafny.Map()
                        compr_11_: int
                        for compr_11_ in (_dafny.Set({(self).f15, d_515_i0_, (0) - (len(d_516_v3_)), d_515_i0_, (self).f15})).Elements:
                            d_518_v10_: int = compr_11_
                            if (d_518_v10_) in (_dafny.Set({(self).f15, d_515_i0_, (0) - (len(d_516_v3_)), d_515_i0_, (self).f15})):
                                coll11_[default__.safeModuloInt(d_518_v10_, (_dafny.MultiSet([(self).f14, not((self).f14), (self).f14, (self).f14, (self).f14])).cardinality)] = -420
                        return _dafny.Map(coll11_)
                    return iife35_()
                    

                return lambda23_

            init10_ = lambda22_(d_506_i0_, d_507_v3_)
            nw83_ = _dafny.Array(None, 22)
            for i0_10_ in range(nw83_.length(0)):
                nw83_[i0_10_] = init10_(i0_10_)
            d_514_v11_ = nw83_
            def lambda24_(d_519_i0_):
                def lambda25_(d_520_i2_):
                    return _dafny.Map({d_519_i0_: (D0_DC0((self).f15, (self).f14)).cf0})

                return lambda25_

            init11_ = lambda24_(d_506_i0_)
            nw84_ = _dafny.Array(None, 19)
            for i0_11_ in range(nw84_.length(0)):
                nw84_[i0_11_] = init11_(i0_11_)
            d_514_v11_ = nw84_
            d_521_v12_: _dafny.Seq
            d_521_v12_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            d_522_v13_: D0
            d_522_v13_ = D0_DC0(len(d_521_v12_), not((self).f14))
            def iife36_(_pat_let12_0):
                def iife37_(d_523_dt__update__tmp_h0_):
                    def iife38_(_pat_let13_0):
                        def iife39_(d_524_dt__update_hcf1_h0_):
                            return D0_DC0((d_523_dt__update__tmp_h0_).cf0, d_524_dt__update_hcf1_h0_)
                        return iife39_(_pat_let13_0)
                    return iife38_((self).f14)
                return iife37_(_pat_let12_0)
            source10_ = iife36_(d_522_v13_)
            if source10_.is_DC0:
                d_525___mcc_h0_ = source10_.cf0
                d_526___mcc_h1_ = source10_.cf1
                d_527_cf1_ = d_526___mcc_h1_
                d_528_cf0_ = d_525___mcc_h0_
                d_527_cf1_ = (not(default__.fm4((self).f14, d_507_v3_, (self).f15, globalState)) if (d_513_v9_)[default__.safeIndex(466, (d_513_v9_).length(0))] else (self).f14)
                (globalState).f4 = d_509_v5_
                d_529_v14_: C0
                nw85_ = C0()
                nw85_.ctor__((d_507_v3_).ispropersubset(d_507_v3_))
                d_529_v14_ = nw85_
                index101_ = default__.safeIndex(466, (d_513_v9_).length(0))
                (d_513_v9_)[index101_] = (d_506_i0_) <= (d_506_i0_)
            elif source10_.is_DC1:
                d_530___mcc_h2_ = source10_.cf2
                d_531___mcc_h3_ = source10_.cf3
                d_532_cf3_ = d_531___mcc_h3_
                d_533_cf2_ = d_530___mcc_h2_
                d_534_v15_: int
                d_534_v15_ = 931
                d_535_v16_: _dafny.Array
                nw86_ = _dafny.Array(_dafny.Map({}), 19)
                d_535_v16_ = nw86_
                index102_ = default__.safeIndex(865, (d_535_v16_).length(0))
                (d_535_v16_)[index102_] = d_511_v7_
                index103_ = default__.safeIndex(865, (d_535_v16_).length(0))
                index104_ = default__.safeIndex(466, (d_513_v9_).length(0))
                rhs74_ = (d_521_v12_)[default__.safeIndex((self).f15, len(d_521_v12_))]
                rhs75_ = (d_511_v7_) | ((d_511_v7_) | (d_511_v7_))
                rhs76_ = (default__.safeModuloInt(d_534_v15_, d_506_i0_)) * (d_506_i0_)
                rhs77_ = ((d_513_v9_)[default__.safeIndex(466, (d_513_v9_).length(0))] if (self).f14 else not((self).f14))
                lhs40_ = d_535_v16_
                lhs41_ = default__.safeIndex(865, (d_535_v16_).length(0))
                lhs42_ = d_513_v9_
                lhs43_ = default__.safeIndex(466, (d_513_v9_).length(0))
                d_534_v15_ = rhs74_
                lhs40_[lhs41_] = rhs75_
                d_534_v15_ = rhs76_
                lhs42_[lhs43_] = rhs77_
                d_536_v17_: _dafny.Map
                d_536_v17_ = _dafny.Map({d_506_i0_: _dafny.Set({False})})
                d_537_v18_: _dafny.Map
                d_537_v18_ = _dafny.Map({d_536_v17_: (self).f15})
                d_537_v18_ = (d_537_v18_).set(d_536_v17_, -424)
                d_538_v19_: _dafny.Map
                d_538_v19_ = _dafny.Map({_dafny.MultiSet([d_506_i0_]): (self).f15})
                d_534_v15_ = default__.safeDivisionInt(((self).f15) * ((self).f15), len(d_538_v19_))
                d_532_cf3_ = not(d_532_cf3_)
            elif source10_.is_DC2:
                d_539___mcc_h4_ = source10_.cf4
                d_540___mcc_h5_ = source10_.cf5
                d_541___mcc_h6_ = source10_.cf6
                d_542_cf6_ = d_541___mcc_h6_
                d_543_cf5_ = d_540___mcc_h5_
                d_544_cf4_ = d_539___mcc_h4_
                d_542_cf6_ = not((d_544_cf4_) == (d_544_cf4_))
                (globalState).f4 = d_509_v5_
                index105_ = default__.safeIndex(466, (d_513_v9_).length(0))
                (d_513_v9_)[index105_] = ((self).f15) != ((self).fm1(d_503_v0_, (self).f15, 602, globalState))
                d_545_v20_: C0
                nw87_ = C0()
                nw87_.ctor__((d_513_v9_)[default__.safeIndex(466, (d_513_v9_).length(0))])
                d_545_v20_ = nw87_
                d_546_v21_: _dafny.Seq
                d_546_v21_ = _dafny.SeqWithoutIsStrInference([d_545_v20_, d_545_v20_, d_545_v20_])
                d_542_cf6_ = ((d_545_v20_).f16 if ((self).f15) != (len(d_546_v21_)) else False)
            elif True:
                d_547___mcc_h7_ = source10_.cf7
                d_548_cf7_ = d_547___mcc_h7_
                d_503_v0_ = d_503_v0_
                index106_ = default__.safeIndex(466, (d_513_v9_).length(0))
                (d_513_v9_)[index106_] = (self).f14
                d_549_v22_: _dafny.MultiSet
                d_549_v22_ = _dafny.MultiSet([((d_513_v9_)[default__.safeIndex(466, (d_513_v9_).length(0))] if (d_513_v9_)[default__.safeIndex(466, (d_513_v9_).length(0))] else True)])
                d_549_v22_ = d_549_v22_
                d_513_v9_ = d_513_v9_
            d_550_v23_: _dafny.Set
            d_550_v23_ = _dafny.Set({d_507_v3_})
            d_551_v24_: D0
            d_551_v24_ = D0_DC1((d_550_v23_) - (d_550_v23_), (self).f14)
            d_551_v24_ = d_551_v24_
        d_552_v25_: int
        d_552_v25_ = 253
        d_552_v25_ = d_552_v25_
        d_553_v26_: _dafny.Set
        d_553_v26_ = _dafny.Set({d_552_v25_, d_552_v25_})
        d_554_v27_: _dafny.Map
        d_554_v27_ = _dafny.Map({d_552_v25_: _dafny.Set({(self).f14, (self).f14})})
        d_555_v28_: _dafny.Set
        d_555_v28_ = _dafny.Set({(self).f14, (self).f14, (self).f14})
        d_556_v29_: D0
        d_556_v29_ = D0_DC2(d_553_v26_, len((_dafny.Set({(self).f14})) | (((d_554_v27_)[(self).f15] if ((self).f15) in (d_554_v27_) else d_555_v28_))), (self).f14)
        source11_ = d_556_v29_
        if source11_.is_DC0:
            d_557___mcc_h8_ = source11_.cf0
            d_558___mcc_h9_ = source11_.cf1
            d_559_cf1_ = d_558___mcc_h9_
            d_560_cf0_ = d_557___mcc_h8_
            d_561_v30_: _dafny.MultiSet
            d_561_v30_ = _dafny.MultiSet([(self).f15, d_552_v25_])
            d_562_v31_: _dafny.Map
            d_562_v31_ = _dafny.Map({(d_561_v30_) - (d_561_v30_): d_552_v25_})
            d_562_v31_ = (d_562_v31_) | ((d_562_v31_) | (d_562_v31_))
            if (d_559_cf1_) or ((d_559_cf1_ if d_559_cf1_ else d_559_cf1_)):
                d_552_v25_ = (self).f15
                d_563_v32_: _dafny.MultiSet
                d_563_v32_ = _dafny.MultiSet([d_559_cf1_, False])
                d_560_cf0_ = ((d_563_v32_).intersection(d_563_v32_)).cardinality
                d_564_v33_: _dafny.Seq
                d_564_v33_ = _dafny.SeqWithoutIsStrInference([d_559_cf1_])
                d_565_v34_: _dafny.Map
                d_565_v34_ = _dafny.Map({len(d_503_v0_): (self).f14})
                d_566_v35_: _dafny.Array
                nw88_ = _dafny.Array(None, 25)
                nw88_[int(0)] = not(d_559_cf1_)
                nw88_[int(1)] = d_559_cf1_
                nw88_[int(2)] = d_559_cf1_
                nw88_[int(3)] = (self).f14
                nw88_[int(4)] = ((d_564_v33_)[default__.safeIndex(d_560_cf0_, len(d_564_v33_))] if d_559_cf1_ else False)
                nw88_[int(5)] = (self).f14
                nw88_[int(6)] = d_559_cf1_
                nw88_[int(7)] = False
                nw88_[int(8)] = not(((self).f15) != ((0) - ((self).f15)))
                nw88_[int(9)] = (d_503_v0_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmjh")))
                nw88_[int(10)] = False
                nw88_[int(11)] = not(d_559_cf1_)
                nw88_[int(12)] = (self).f14
                nw88_[int(13)] = not (d_559_cf1_) or ((self).f14)
                nw88_[int(14)] = (d_561_v30_).ispropersubset(d_561_v30_)
                nw88_[int(15)] = False
                nw88_[int(16)] = ((d_565_v34_)[d_560_cf0_] if (d_560_cf0_) in (d_565_v34_) else True)
                nw88_[int(17)] = (d_563_v32_).issubset(d_563_v32_)
                nw88_[int(18)] = d_559_cf1_
                nw88_[int(19)] = (self).f14
                nw88_[int(20)] = (self).f14
                nw88_[int(21)] = True
                nw88_[int(22)] = (d_559_cf1_ if (self).f14 else True)
                nw88_[int(23)] = not(d_559_cf1_)
                nw88_[int(24)] = d_559_cf1_
                d_566_v35_ = nw88_
                d_567_v36_: str
                d_567_v36_ = _dafny.CodePoint('s')
                d_568_v37_: _dafny.Set
                d_568_v37_ = _dafny.Set({d_567_v36_})
                nw89_ = _dafny.Array(None, 15)
                nw89_[int(0)] = d_559_cf1_
                nw89_[int(1)] = not(not((self).f14))
                nw89_[int(2)] = (self).f14
                nw89_[int(3)] = d_559_cf1_
                nw89_[int(4)] = False
                nw89_[int(5)] = (d_559_cf1_) and ((self).f14)
                nw89_[int(6)] = not((self).f14)
                nw89_[int(7)] = (True) == ((self).f14)
                nw89_[int(8)] = d_559_cf1_
                nw89_[int(9)] = not(d_559_cf1_)
                nw89_[int(10)] = True
                nw89_[int(11)] = (self).f14
                nw89_[int(12)] = d_559_cf1_
                nw89_[int(13)] = (d_552_v25_) < (len(d_568_v37_))
                nw89_[int(14)] = False
                d_566_v35_ = nw89_
                d_569_v38_: C0
                nw90_ = C0()
                nw90_.ctor__(d_559_cf1_)
                d_569_v38_ = nw90_
                index107_ = default__.safeIndex(394, (d_566_v35_).length(0))
                (d_566_v35_)[index107_] = d_559_cf1_
                index108_ = default__.safeIndex(394, (d_566_v35_).length(0))
                (d_566_v35_)[index108_] = not(d_559_cf1_)
            elif True:
                d_570_v39_: _dafny.MultiSet
                d_570_v39_ = _dafny.MultiSet([d_556_v29_])
                d_571_v40_: _dafny.Seq
                d_571_v40_ = _dafny.SeqWithoutIsStrInference([D0_DC2(d_553_v26_, d_560_cf0_, d_559_cf1_), d_556_v29_])
                d_572_v41_: _dafny.Set
                d_572_v41_ = _dafny.Set({d_570_v39_, d_570_v39_, (d_570_v39_) | (_dafny.MultiSet(d_571_v40_)), _dafny.MultiSet([d_556_v29_]), _dafny.MultiSet([d_556_v29_])})
                d_573_v42_: _dafny.Map
                d_573_v42_ = _dafny.Map({d_559_cf1_: d_559_cf1_})
                rhs78_ = _dafny.Set({(d_570_v39_) - (d_570_v39_)})
                rhs79_ = d_503_v0_
                rhs80_ = default__.fm0((d_573_v42_) | (d_573_v42_), (d_559_cf1_) or ((self).f14), (_dafny.SeqWithoutIsStrInference([D0_DC3(_dafny.CodePoint('k')) for d_574_i3_ in range(default__.abs(158))])) < (_dafny.SeqWithoutIsStrInference([D0_DC3(_dafny.CodePoint('m')) for d_575_i4_ in range(default__.abs(-891))])), globalState)
                rhs81_ = d_555_v28_
                rhs82_ = (self).f14
                d_572_v41_ = rhs78_
                d_503_v0_ = rhs79_
                d_560_cf0_ = rhs80_
                d_555_v28_ = rhs81_
                d_559_cf1_ = rhs82_
                d_559_cf1_ = not((default__.fm0(d_573_v42_, not(not((self).f14)), d_559_cf1_, globalState)) == (default__.fm0(d_573_v42_, (self).f14, (self).f14, globalState)))
                d_560_cf0_ = default__.safeModuloInt((self).f15, len(default__.fm8((self).f14, (self).f14, (self).f14, d_559_cf1_, globalState)))
                d_559_cf1_ = (self).f14
                d_560_cf0_ = (self).f15
            d_576_v43_: _dafny.Array
            nw91_ = _dafny.Array(False, 26)
            d_576_v43_ = nw91_
            def lambda26_(d_577_i5_):
                return (self).f14

            init12_ = lambda26_
            nw92_ = _dafny.Array(None, 7)
            for i0_12_ in range(nw92_.length(0)):
                nw92_[i0_12_] = init12_(i0_12_)
            d_576_v43_ = nw92_
            d_578_v44_: _dafny.Array
            nw93_ = _dafny.Array(int(0), 4)
            d_578_v44_ = nw93_
            index109_ = default__.safeIndex(24, (d_578_v44_).length(0))
            (d_578_v44_)[index109_] = len(_dafny.SeqWithoutIsStrInference([(d_503_v0_)[default__.safeIndex(d_552_v25_, len(d_503_v0_))] for d_579_i6_ in range(default__.abs(-152))]))
            index110_ = default__.safeIndex(24, (d_578_v44_).length(0))
            (d_578_v44_)[index110_] = (0) - ((d_560_cf0_ if (self).f14 else (len(_dafny.SeqWithoutIsStrInference([(self).f14])) if d_559_cf1_ else d_560_cf0_)))
        elif source11_.is_DC1:
            d_580___mcc_h10_ = source11_.cf2
            d_581___mcc_h11_ = source11_.cf3
            d_582_cf3_ = d_581___mcc_h11_
            d_583_cf2_ = d_580___mcc_h10_
            d_584_v45_: str
            d_584_v45_ = _dafny.CodePoint('u')
            d_585_v46_: C0
            nw94_ = C0()
            nw94_.ctor__((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_586_i7_ in range(default__.abs(153))])) < ((d_503_v0_).set(default__.safeIndex((self).f15, len(d_503_v0_)), d_584_v45_)))
            d_585_v46_ = nw94_
            d_587_v47_: _dafny.Array
            nw95_ = _dafny.Array(False, 10)
            d_587_v47_ = nw95_
            d_587_v47_ = d_587_v47_
            index111_ = default__.safeIndex(591, (d_587_v47_).length(0))
            (d_587_v47_)[index111_] = (d_585_v46_).f16
            index112_ = default__.safeIndex(591, (d_587_v47_).length(0))
            (d_587_v47_)[index112_] = (d_585_v46_).f16
            d_588_v48_: T0
            nw96_ = C0()
            nw96_.ctor__((d_585_v46_).f16)
            d_588_v48_ = nw96_
            d_589_v49_: _dafny.Seq
            d_589_v49_ = _dafny.SeqWithoutIsStrInference([d_588_v48_])
            d_589_v49_ = d_589_v49_
        elif source11_.is_DC2:
            d_590___mcc_h12_ = source11_.cf4
            d_591___mcc_h13_ = source11_.cf5
            d_592___mcc_h14_ = source11_.cf6
            d_593_cf6_ = d_592___mcc_h14_
            d_594_cf5_ = d_591___mcc_h13_
            d_595_cf4_ = d_590___mcc_h12_
            d_596_v50_: _dafny.Seq
            d_596_v50_ = _dafny.SeqWithoutIsStrInference([d_552_v25_, len(d_553_v26_), d_594_cf5_, d_594_cf5_])
            d_597_v51_: _dafny.Seq
            d_597_v51_ = _dafny.SeqWithoutIsStrInference([False, default__.fm4(not((self).f14), d_555_v28_, (self).f15, globalState)])
            d_598_v52_: _dafny.Map
            d_598_v52_ = _dafny.Map({(self).f14: d_597_v51_})
            d_599_v53_: _dafny.Array
            nw97_ = _dafny.Array(None, 15)
            nw97_[int(0)] = (self).f15
            nw97_[int(1)] = d_594_cf5_
            nw97_[int(2)] = d_594_cf5_
            nw97_[int(3)] = (d_552_v25_ if (self).f14 else (self).f15)
            nw97_[int(4)] = d_552_v25_
            nw97_[int(5)] = 918
            nw97_[int(6)] = (self).f15
            nw97_[int(7)] = 718
            nw97_[int(8)] = (d_552_v25_) - (d_594_cf5_)
            nw97_[int(9)] = len(d_596_v50_)
            nw97_[int(10)] = (self).f15
            nw97_[int(11)] = (0) - (d_552_v25_)
            nw97_[int(12)] = 67
            nw97_[int(13)] = (len(d_598_v52_)) - (d_552_v25_)
            nw97_[int(14)] = d_552_v25_
            d_599_v53_ = nw97_
            index113_ = default__.safeIndex(241, (d_599_v53_).length(0))
            (d_599_v53_)[index113_] = (self).f15
            d_600_v54_: _dafny.Array
            nw98_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
            d_600_v54_ = nw98_
            index114_ = default__.safeIndex(678, (d_600_v54_).length(0))
            (d_600_v54_)[index114_] = d_503_v0_
            d_601_v55_: str
            d_601_v55_ = _dafny.CodePoint('d')
            d_602_v56_: _dafny.MultiSet
            d_602_v56_ = _dafny.MultiSet([d_601_v55_])
            d_603_v57_: _dafny.Map
            d_603_v57_ = _dafny.Map({(self).f14: d_552_v25_})
            index115_ = default__.safeIndex(241, (d_599_v53_).length(0))
            index116_ = default__.safeIndex(678, (d_600_v54_).length(0))
            rhs83_ = (d_602_v56_).isdisjoint((d_602_v56_).set(d_601_v55_, default__.abs(d_594_cf5_)))
            rhs84_ = (default__.fm9(d_601_v55_, (self).f15, len(_dafny.SeqWithoutIsStrInference([d_593_cf6_])), d_603_v57_, globalState)) - (d_555_v28_)
            rhs85_ = (0) - (len((_dafny.SeqWithoutIsStrInference([d_503_v0_, d_503_v0_])) + ((d_505_v2_) + (d_505_v2_))))
            rhs86_ = d_503_v0_
            lhs44_ = d_599_v53_
            lhs45_ = default__.safeIndex(241, (d_599_v53_).length(0))
            lhs46_ = d_600_v54_
            lhs47_ = default__.safeIndex(678, (d_600_v54_).length(0))
            d_593_cf6_ = rhs83_
            d_555_v28_ = rhs84_
            lhs44_[lhs45_] = rhs85_
            lhs46_[lhs47_] = rhs86_
            d_604_v58_: _dafny.Map
            d_604_v58_ = _dafny.Map({d_599_v53_: (self).f14})
            d_605_v59_: _dafny.Array
            nw99_ = _dafny.Array(None, 23)
            nw99_[int(0)] = d_593_cf6_
            nw99_[int(1)] = (self).f14
            nw99_[int(2)] = (self).f14
            nw99_[int(3)] = ((d_604_v58_)[d_599_v53_] if (d_599_v53_) in (d_604_v58_) else not(d_593_cf6_))
            nw99_[int(4)] = not(True)
            nw99_[int(5)] = d_593_cf6_
            nw99_[int(6)] = default__.fm4(True, d_555_v28_, len((d_600_v54_)[default__.safeIndex(678, (d_600_v54_).length(0))]), globalState)
            nw99_[int(7)] = not(not((d_594_cf5_) not in (d_596_v50_)))
            nw99_[int(8)] = d_593_cf6_
            nw99_[int(9)] = (self).f14
            nw99_[int(10)] = False
            nw99_[int(11)] = d_593_cf6_
            nw99_[int(12)] = (self).f14
            nw99_[int(13)] = (d_594_cf5_) == ((d_599_v53_)[default__.safeIndex(241, (d_599_v53_).length(0))])
            nw99_[int(14)] = ((d_600_v54_)[default__.safeIndex(678, (d_600_v54_).length(0))]) != ((d_600_v54_)[default__.safeIndex(678, (d_600_v54_).length(0))])
            nw99_[int(15)] = (self).f14
            nw99_[int(16)] = (self).f14
            nw99_[int(17)] = (d_593_cf6_ if d_593_cf6_ else (self).f14)
            nw99_[int(18)] = not (d_593_cf6_) or ((self).f14)
            nw99_[int(19)] = (self).f14
            nw99_[int(20)] = False
            nw99_[int(21)] = d_593_cf6_
            nw99_[int(22)] = d_593_cf6_
            d_605_v59_ = nw99_
            index117_ = default__.safeIndex(274, (d_605_v59_).length(0))
            (d_605_v59_)[index117_] = (self).f14
            index118_ = default__.safeIndex(445, (d_605_v59_).length(0))
            (d_605_v59_)[index118_] = default__.fm4((self).f14, d_555_v28_, (self).f15, globalState)
            d_606_v60_: D2
            d_606_v60_ = D2_DC6(_dafny.SeqWithoutIsStrInference([d_593_cf6_]))
            d_607_v61_: _dafny.MultiSet
            d_607_v61_ = _dafny.MultiSet([not((self).f14)])
            index119_ = default__.safeIndex(241, (d_599_v53_).length(0))
            index120_ = default__.safeIndex(274, (d_605_v59_).length(0))
            index121_ = default__.safeIndex(445, (d_605_v59_).length(0))
            def iife40_(_pat_let14_0):
                def iife41_(d_608_dt__update__tmp_h1_):
                    def iife42_(_pat_let15_0):
                        def iife43_(d_609_dt__update_hcf6_h0_):
                            def iife44_(_pat_let16_0):
                                def iife45_(d_610_dt__update_hcf5_h0_):
                                    return D0_DC2((d_608_dt__update__tmp_h1_).cf4, d_610_dt__update_hcf5_h0_, d_609_dt__update_hcf6_h0_)
                                return iife45_(_pat_let16_0)
                            return iife44_(360)
                        return iife43_(_pat_let15_0)
                    return iife42_(not ((self).f14) or ((self).f14))
                return iife41_(_pat_let14_0)
            rhs87_ = iife40_(d_556_v29_)
            rhs88_ = (self).f15
            rhs89_ = False
            rhs90_ = True
            rhs91_ = ((d_607_v61_).ispropersubset(d_607_v61_) if (_dafny.MultiSet((d_606_v60_).cf10)).ispropersubset(d_607_v61_) else not(False))
            lhs48_ = d_599_v53_
            lhs49_ = default__.safeIndex(241, (d_599_v53_).length(0))
            lhs50_ = d_605_v59_
            lhs51_ = default__.safeIndex(274, (d_605_v59_).length(0))
            lhs52_ = d_605_v59_
            lhs53_ = default__.safeIndex(445, (d_605_v59_).length(0))
            d_556_v29_ = rhs87_
            lhs48_[lhs49_] = rhs88_
            d_593_cf6_ = rhs89_
            lhs50_[lhs51_] = rhs90_
            lhs52_[lhs53_] = rhs91_
            d_611_v62_: _dafny.Array
            nw100_ = _dafny.Array(None, 16)
            d_611_v62_ = nw100_
            d_612_v63_: C0
            nw101_ = C0()
            nw101_.ctor__(not(True))
            d_612_v63_ = nw101_
            index122_ = default__.safeIndex(851, (d_611_v62_).length(0))
            (d_611_v62_)[index122_] = d_612_v63_
            index123_ = default__.safeIndex(851, (d_611_v62_).length(0))
            (d_611_v62_)[index123_] = d_612_v63_
            d_613_v64_: _dafny.Map
            d_613_v64_ = _dafny.Map({(d_605_v59_)[default__.safeIndex(274, (d_605_v59_).length(0))]: (d_600_v54_)[default__.safeIndex(678, (d_600_v54_).length(0))]})
            d_594_cf5_ = len(((d_613_v64_).set(d_593_cf6_, (d_600_v54_)[default__.safeIndex(678, (d_600_v54_).length(0))])) | ((d_613_v64_) | (d_613_v64_)))
        elif True:
            d_614___mcc_h15_ = source11_.cf7
            d_615_cf7_ = d_614___mcc_h15_
            d_616_v65_: D2
            d_616_v65_ = D2_DC7(default__.fm6((self).f14, d_552_v25_, globalState), (self).f14, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lawxj")))) != (819))
            d_616_v65_ = d_616_v65_
            if (default__.safeDivisionInt(d_552_v25_, d_552_v25_)) != ((self).f15):
                d_617_v66_: C0
                nw102_ = C0()
                nw102_.ctor__((self).f14)
                d_617_v66_ = nw102_
                d_618_v67_: _dafny.Seq
                d_618_v67_ = _dafny.SeqWithoutIsStrInference([(self).f14, not(not ((self).f14) or ((d_617_v66_).f16)), (d_617_v66_).f16, ((0) - ((self).f15)) != ((self).f15)])
                d_619_v68_: _dafny.Array
                nw103_ = _dafny.Array(False, 20)
                d_619_v68_ = nw103_
                index124_ = default__.safeIndex(777, (d_619_v68_).length(0))
                (d_619_v68_)[index124_] = (self).f14
                index125_ = default__.safeIndex(777, (d_619_v68_).length(0))
                rhs92_ = d_618_v67_
                rhs93_ = ((self).f15) <= (d_552_v25_)
                rhs94_ = d_617_v66_
                rhs95_ = d_619_v68_
                lhs54_ = d_619_v68_
                lhs55_ = default__.safeIndex(777, (d_619_v68_).length(0))
                d_618_v67_ = rhs92_
                lhs54_[lhs55_] = rhs93_
                d_617_v66_ = rhs94_
                d_619_v68_ = rhs95_
                d_620_v69_: C0
                nw104_ = C0()
                nw104_.ctor__((d_619_v68_)[default__.safeIndex(777, (d_619_v68_).length(0))])
                d_620_v69_ = nw104_
                d_621_v70_: C0
                nw105_ = C0()
                nw105_.ctor__(((d_620_v69_).fm5(False, d_552_v25_, 43, (self).f14, globalState)) <= (d_503_v0_))
                d_621_v70_ = nw105_
                d_622_v71_: D0
                d_622_v71_ = D0_DC0((self).f15, (d_620_v69_).f16)
                d_623_v72_: C0
                nw106_ = C0()
                nw106_.ctor__(((d_622_v71_).cf1) or (default__.fm6((d_621_v70_).f16, d_552_v25_, globalState)))
                d_623_v72_ = nw106_
            elif True:
                d_552_v25_ = ((self).f15 if (self).f14 else default__.safeModuloInt(d_552_v25_, d_552_v25_))
                d_624_v73_: _dafny.Map
                d_624_v73_ = _dafny.Map({d_552_v25_: d_552_v25_})
                d_624_v73_ = _dafny.Map({(self).f15: (self).f15})
                d_625_v74_: C0
                nw107_ = C0()
                nw107_.ctor__(False)
                d_625_v74_ = nw107_
                d_626_v75_: _dafny.Seq
                d_626_v75_ = _dafny.SeqWithoutIsStrInference([(self).f14, default__.fm4((d_625_v74_).f16, d_555_v28_, 231, globalState)])
                d_627_v76_: bool
                d_628_v77_: _dafny.Array
                d_629_v78_: int
                d_630_v79_: int
                out44_: bool
                out45_: _dafny.Array
                out46_: int
                out47_: int
                out44_, out45_, out46_, out47_ = default__.m0(d_552_v25_, d_626_v75_, globalState)
                d_627_v76_ = out44_
                d_628_v77_ = out45_
                d_629_v78_ = out46_
                d_630_v79_ = out47_
                d_631_v80_: D0
                d_631_v80_ = D0_DC3(d_615_cf7_)
                d_631_v80_ = d_631_v80_
            d_632_v81_: _dafny.Map
            d_632_v81_ = _dafny.Map({(self).f14: (self).f15})
            d_633_v82_: _dafny.Map
            d_633_v82_ = _dafny.Map({(self).f14: (self).f14})
            d_634_v83_: C0
            nw108_ = C0()
            nw108_.ctor__((self).f14)
            d_634_v83_ = nw108_
            d_635_v84_: _dafny.Map
            d_635_v84_ = _dafny.Map({d_552_v25_: d_634_v83_})
            d_636_v85_: _dafny.Array
            nw109_ = _dafny.Array(None, 16)
            nw109_[int(0)] = not ((self).f14) or ((self).f14)
            nw109_[int(1)] = (self).f14
            nw109_[int(2)] = (self).f14
            nw109_[int(3)] = not ((self).f14) or (default__.fm4((self).f14, d_555_v28_, len(d_632_v81_), globalState))
            nw109_[int(4)] = (self).f14
            nw109_[int(5)] = not((self).f14)
            nw109_[int(6)] = ((d_633_v82_)[(self).f14] if ((self).f14) in (d_633_v82_) else (self).f14)
            nw109_[int(7)] = (self).f14
            nw109_[int(8)] = (self).f14
            nw109_[int(9)] = (self).f14
            nw109_[int(10)] = (self).f14
            nw109_[int(11)] = (self).f14
            nw109_[int(12)] = (len(d_635_v84_)) <= (-202)
            nw109_[int(13)] = (self).f14
            nw109_[int(14)] = ((0) - ((self).f15)) > (len(_dafny.SeqWithoutIsStrInference([d_615_cf7_ for d_637_i8_ in range(default__.abs(438))])))
            nw109_[int(15)] = not((d_634_v83_).f16)
            d_636_v85_ = nw109_
            index126_ = default__.safeIndex(51, (d_636_v85_).length(0))
            (d_636_v85_)[index126_] = (self).f14
            index127_ = default__.safeIndex(51, (d_636_v85_).length(0))
            (d_636_v85_)[index127_] = (self).f14
            d_638_v86_: _dafny.Map
            d_638_v86_ = _dafny.Map({default__.fm4(True, d_555_v28_, -291, globalState): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyo")) if (d_636_v85_)[default__.safeIndex(51, (d_636_v85_).length(0))] else d_503_v0_)})
            d_638_v86_ = (d_638_v86_).set((self).f14, (d_503_v0_) + (d_503_v0_))
        d_639_v87_: bool
        d_639_v87_ = False
        d_640_v88_: _dafny.Array
        def lambda27_(d_641_v87_):
            def lambda28_(d_642_i9_):
                return d_641_v87_

            return lambda28_

        init13_ = lambda27_(d_639_v87_)
        nw110_ = _dafny.Array(None, 23)
        for i0_13_ in range(nw110_.length(0)):
            nw110_[i0_13_] = init13_(i0_13_)
        d_640_v88_ = nw110_
        d_643_v89_: _dafny.MultiSet
        d_643_v89_ = _dafny.MultiSet([d_640_v88_, d_640_v88_])
        d_644_v90_: _dafny.MultiSet
        d_644_v90_ = _dafny.MultiSet([(self).fm1(d_503_v0_, (self).f15, d_552_v25_, globalState)])
        d_639_v87_ = ((d_643_v89_).set(d_640_v88_, default__.abs((0) - ((d_644_v90_).cardinality)))).issubset(d_643_v89_)

    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
