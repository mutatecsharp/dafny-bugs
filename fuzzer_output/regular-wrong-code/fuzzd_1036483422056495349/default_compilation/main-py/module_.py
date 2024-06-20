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
    def fm0(p0, p1, p2, p3, globalState):
        return default__.safeDivisionInt(len((_dafny.SeqWithoutIsStrInference([997])) + (_dafny.SeqWithoutIsStrInference([-776, (_dafny.MultiSet([-746, 499])).cardinality, 549]))), 38)

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return (D1_DC4(not(False), False, True)).cf6

    @staticmethod
    def fm2(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([39, (_dafny.MultiSet([633, (0) - (-268)])).cardinality])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_0_i0_ in range(default__.abs(-431))]))]))

    @staticmethod
    def fm3(globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjsfret"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwkap"))))

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-185 for d_1_i0_ in range(default__.abs(945))])): 886})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tysocxty"))): (_dafny.MultiSet([False])).cardinality}))

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False, not(False), not(False), False]))) + ((_dafny.SeqWithoutIsStrInference([True, True])) + (_dafny.SeqWithoutIsStrInference([False, True, not(True), True]))))

    @staticmethod
    def fm8(globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([542, 145]): False})) | ((D13_DC34(_dafny.Map({_dafny.SeqWithoutIsStrInference([531, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pq")))]): True}))).cf45)

    @staticmethod
    def fm9(p0, p1, globalState):
        return _dafny.Map({(413) <= ((_dafny.MultiSet([603])).cardinality): ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({_dafny.Set({686}), _dafny.Set({815})}))), len(_dafny.SeqWithoutIsStrInference([D11_DC30(False) for d_2_i0_ in range(default__.abs(-851))]))]))})

    @staticmethod
    def fm10(p0, globalState):
        return D2_DC5(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvfc"))): len(_dafny.SeqWithoutIsStrInference([True]))}))

    @staticmethod
    def fm11(globalState):
        return D1_DC3(((_dafny.MultiSet([False, not(False)])).cardinality) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epbbayn")))))

    @staticmethod
    def fm12(globalState):
        return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ap")))])

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: str
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_3_i0_ in range(default__.abs(139))])).Elements:
                d_4_v0_: str = compr_0_
                if (d_4_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_3_i0_ in range(default__.abs(139))])):
                    coll0_[d_4_v0_] = True
            return _dafny.Map(coll0_)
        return _dafny.Set({(len(iife0_()
        )) > (587), (True if False else True), (False) or (False)})

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        return _dafny.Set({_dafny.CodePoint('h')})

    @staticmethod
    def fm15(p0, p1, globalState):
        if (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, True})), len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([True]))) for d_5_i0_ in range(default__.abs(973))]))]))) > (231):
            return _dafny.CodePoint('k')
        elif True:
            return _dafny.CodePoint('c')

    @staticmethod
    def fm16(p0, globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pyah"))])

    @staticmethod
    def fm17(globalState):
        if False:
            return D6_DC17(380, True)
        elif True:
            return D6_DC17(-419, True)

    @staticmethod
    def fm18(p0, globalState):
        return D5_DC13(((_dafny.MultiSet([False])).cardinality) - (len(_dafny.Map({854: _dafny.CodePoint('q')}))))

    @staticmethod
    def fm19(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([D1_DC3(len(_dafny.SeqWithoutIsStrInference([338]))), D1_DC3(553), D1_DC3(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))), D1_DC3((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-257]))).cardinality), D1_DC3(len(_dafny.SeqWithoutIsStrInference([295, 780])))])).intersection(_dafny.MultiSet([D1_DC3(567)]))).intersection((_dafny.MultiSet([D1_DC3(602), D1_DC3(len(_dafny.Map({True: D1_DC3(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htapck"))))}))), D1_DC3(-138)])).intersection(_dafny.MultiSet([D1_DC3(967)])))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        hi0_ = p0
        for d_6_i0_ in range(default__.safeDivisionInt(p0, p0), hi0_):
            d_7_v0_: bool
            d_7_v0_ = False
            d_8_v1_: _dafny.Map
            d_8_v1_ = _dafny.Map({-239: d_7_v0_})
            d_9_v2_: D0
            d_9_v2_ = D0_DC0(d_7_v0_)
            d_10_v3_: _dafny.Set
            d_10_v3_ = _dafny.Set({d_6_i0_})
            d_11_v4_: C1
            nw0_ = C1()
            nw0_.ctor__(d_7_v0_, (p1) - (len((d_8_v1_).set(p0, (d_9_v2_).cf0))), p0, (d_10_v3_).ispropersubset(_dafny.Set({p0, p0})))
            d_11_v4_ = nw0_
            r0 = default__.safeModuloInt(p1, p0)
            d_12_v6_: _dafny.Map
            def iife1_():
                coll1_ = _dafny.Set()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(947, 518):
                    d_13_v5_: int = compr_1_
                    if ((947) <= (d_13_v5_)) and ((d_13_v5_) < (518)):
                        coll1_ = coll1_.union(_dafny.Set([default__.safeDivisionInt(d_13_v5_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_14_i1_ in range(default__.abs(491))])))]))
                return _dafny.Set(coll1_)
            d_12_v6_ = _dafny.Map({p1: iife1_()
            })
            d_15_v7_: _dafny.Map
            d_15_v7_ = _dafny.Map({d_11_v4_.f21: len(d_12_v6_)})
            d_16_v8_: _dafny.Seq
            d_16_v8_ = _dafny.SeqWithoutIsStrInference([-998, p1])
            d_17_v9_: D5
            d_17_v9_ = D5_DC13((d_16_v8_)[default__.safeIndex(d_6_i0_, len(d_16_v8_))])
            d_18_v10_: _dafny.Set
            d_18_v10_ = _dafny.Set({d_17_v9_})
            d_19_v11_: _dafny.Set
            d_19_v11_ = _dafny.Set({d_18_v10_})
            d_20_v12_: _dafny.Map
            d_20_v12_ = _dafny.Map({d_15_v7_: d_19_v11_})
            d_21_v13_: _dafny.MultiSet
            d_21_v13_ = _dafny.MultiSet([d_11_v4_.f21])
            d_20_v12_ = (d_20_v12_).set(_dafny.Map({d_7_v0_: (d_21_v13_).cardinality}), d_19_v11_)
            d_22_v14_: C1
            nw1_ = C1()
            nw1_.ctor__(d_7_v0_, 350, p1, d_7_v0_)
            d_22_v14_ = nw1_
        d_23_v15_: bool
        d_23_v15_ = False
        d_24_v16_: _dafny.Seq
        d_24_v16_ = _dafny.SeqWithoutIsStrInference([d_23_v15_, d_23_v15_])
        if not((d_24_v16_)[default__.safeIndex(p1, len(d_24_v16_))]):
            d_25_v17_: str
            d_25_v17_ = _dafny.CodePoint('m')
            d_26_v18_: _dafny.Map
            d_26_v18_ = _dafny.Map({d_23_v15_: 837})
            d_27_v19_: _dafny.MultiSet
            d_27_v19_ = _dafny.MultiSet([p0, ((d_26_v18_)[d_23_v15_] if (d_23_v15_) in (d_26_v18_) else p1)])
            d_28_v20_: D2
            d_28_v20_ = D2_DC6(d_27_v19_)
            d_29_v21_: D2
            d_29_v21_ = D2_DC7(d_28_v20_)
            d_25_v17_ = default__.fm15(d_29_v21_, not(d_23_v15_), globalState)
            d_30_v22_: _dafny.Array
            nw2_ = _dafny.Array(None, 2)
            nw2_[int(0)] = d_25_v17_
            nw2_[int(1)] = d_25_v17_
            d_30_v22_ = nw2_
            index0_ = default__.safeIndex(540, (d_30_v22_).length(0))
            (d_30_v22_)[index0_] = _dafny.CodePoint('a')
            d_31_v23_: _dafny.MultiSet
            d_31_v23_ = _dafny.MultiSet([d_23_v15_, d_23_v15_])
            d_32_v24_: _dafny.Seq
            d_32_v24_ = _dafny.SeqWithoutIsStrInference([d_31_v23_])
            d_33_v25_: _dafny.Seq
            d_33_v25_ = _dafny.SeqWithoutIsStrInference([p0])
            index1_ = default__.safeIndex(540, (d_30_v22_).length(0))
            rhs0_ = (_dafny.SeqWithoutIsStrInference([len(d_32_v24_)])) != (d_33_v25_)
            rhs1_ = not((d_23_v15_) not in (d_31_v23_))
            rhs2_ = d_25_v17_
            lhs0_ = globalState
            lhs1_ = globalState
            lhs2_ = d_30_v22_
            lhs3_ = default__.safeIndex(540, (d_30_v22_).length(0))
            lhs0_.f13 = rhs0_
            lhs1_.f13 = rhs1_
            lhs2_[lhs3_] = rhs2_
            d_34_v26_: T1
            nw3_ = C1()
            nw3_.ctor__(not(d_23_v15_), p0, p0, d_23_v15_)
            d_34_v26_ = nw3_
            d_35_v27_: _dafny.Seq
            d_35_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
            d_36_v28_: _dafny.MultiSet
            d_36_v28_ = _dafny.MultiSet([d_35_v27_])
            d_37_v29_: D0
            d_37_v29_ = D0_DC1((d_34_v26_).f19, len(d_24_v16_), d_23_v15_)
            d_38_v30_: _dafny.Map
            d_38_v30_ = _dafny.Map({p0: d_35_v27_})
            d_39_v31_: _dafny.Array
            nw4_ = _dafny.Array(None, 21)
            nw4_[int(0)] = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfnc")) for d_40_i2_ in range(default__.abs(128))]))
            nw4_[int(1)] = d_34_v26_.f20
            nw4_[int(2)] = d_34_v26_.f20
            nw4_[int(3)] = (d_34_v26_).f19
            nw4_[int(4)] = len((d_24_v16_) + (d_24_v16_))
            nw4_[int(5)] = (-199) - (len(d_35_v27_))
            nw4_[int(6)] = d_34_v26_.f20
            nw4_[int(7)] = d_34_v26_.f20
            nw4_[int(8)] = len((_dafny.SeqWithoutIsStrInference([579, (d_34_v26_).f19])) + (d_33_v25_))
            nw4_[int(9)] = ((d_36_v28_).intersection(default__.fm16((d_37_v29_).cf3, globalState))).cardinality
            nw4_[int(10)] = d_34_v26_.f20
            nw4_[int(11)] = len(d_35_v27_)
            nw4_[int(12)] = len(d_38_v30_)
            nw4_[int(13)] = (d_34_v26_).f19
            nw4_[int(14)] = d_34_v26_.f20
            nw4_[int(15)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))
            nw4_[int(16)] = p1
            nw4_[int(17)] = d_34_v26_.f20
            nw4_[int(18)] = ((default__.fm12(globalState)).set(p1, default__.abs(p1))).cardinality
            nw4_[int(19)] = -579
            nw4_[int(20)] = d_34_v26_.f20
            d_39_v31_ = nw4_
            index2_ = default__.safeIndex(980, (d_39_v31_).length(0))
            (d_39_v31_)[index2_] = -751
            d_41_v32_: T0
            nw5_ = C0()
            nw5_.ctor__(D2_DC6(d_27_v19_), -841, (d_34_v26_).f18)
            d_41_v32_ = nw5_
            d_42_v33_: _dafny.Map
            d_42_v33_ = _dafny.Map({d_41_v32_: p1})
            index3_ = default__.safeIndex(980, (d_39_v31_).length(0))
            rhs3_ = d_23_v15_
            rhs4_ = _dafny.SeqWithoutIsStrInference([((d_42_v33_)[d_41_v32_] if (d_41_v32_) in (d_42_v33_) else p1), (d_34_v26_).f19, (p1 if d_23_v15_ else d_34_v26_.f20), (d_34_v26_).f19])
            rhs5_ = d_34_v26_
            rhs6_ = (p0) + (p0)
            lhs4_ = globalState
            lhs5_ = d_39_v31_
            lhs6_ = default__.safeIndex(980, (d_39_v31_).length(0))
            lhs4_.f13 = rhs3_
            d_33_v25_ = rhs4_
            d_34_v26_ = rhs5_
            lhs5_[lhs6_] = rhs6_
            d_43_v34_: _dafny.Array
            nw6_ = _dafny.Array(None, 4)
            d_43_v34_ = nw6_
            d_44_v35_: _dafny.Map
            d_44_v35_ = _dafny.Map({d_43_v34_: (d_34_v26_).f18})
            d_44_v35_ = (d_44_v35_).set(d_43_v34_, d_23_v15_)
            d_45_v36_: _dafny.Set
            d_45_v36_ = _dafny.Set({len((d_33_v25_).set(default__.safeIndex((d_34_v26_).f19, len(d_33_v25_)), p1))})
            (d_34_v26_).f20 = (0) - ((0) - (default__.safeModuloInt(len(d_33_v25_), default__.safeDivisionInt(p1, len(d_45_v36_)))))
        elif True:
            (globalState).f13 = d_23_v15_
            d_46_v37_: _dafny.Array
            nw7_ = _dafny.Array(None, 23)
            d_46_v37_ = nw7_
            d_47_v38_: D6
            d_47_v38_ = D6_DC16(p1)
            index4_ = default__.safeIndex(761, (d_46_v37_).length(0))
            (d_46_v37_)[index4_] = d_47_v38_
            index5_ = default__.safeIndex(761, (d_46_v37_).length(0))
            (d_46_v37_)[index5_] = D6_DC16(p1)
            d_48_v39_: _dafny.Array
            nw8_ = _dafny.Array(None, 9)
            d_48_v39_ = nw8_
            d_49_v40_: _dafny.MultiSet
            d_49_v40_ = _dafny.MultiSet([p1, p1])
            d_50_v41_: D2
            d_50_v41_ = D2_DC6(d_49_v40_)
            d_51_v42_: T0
            nw9_ = C0()
            nw9_.ctor__(d_50_v41_, p1, d_23_v15_)
            d_51_v42_ = nw9_
            d_52_v43_: D6
            d_52_v43_ = D6_DC15(d_51_v42_)
            index6_ = default__.safeIndex(633, (d_48_v39_).length(0))
            (d_48_v39_)[index6_] = d_52_v43_
            index7_ = default__.safeIndex(633, (d_48_v39_).length(0))
            (d_48_v39_)[index7_] = d_52_v43_
            (globalState).f13 = ((d_24_v16_) + (d_24_v16_)) < (d_24_v16_)
            d_53_v44_: C0
            nw10_ = C0()
            nw10_.ctor__(D2_DC6(d_49_v40_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhrsgbmu"))), d_23_v15_)
            d_53_v44_ = nw10_
            d_54_v45_: _dafny.Seq
            d_54_v45_ = _dafny.SeqWithoutIsStrInference([d_53_v44_])
            d_55_v46_: _dafny.Array
            nw11_ = _dafny.Array(None, 4)
            nw11_[int(0)] = d_54_v45_
            nw11_[int(1)] = (d_54_v45_) + (d_54_v45_)
            nw11_[int(2)] = d_54_v45_
            nw11_[int(3)] = d_54_v45_
            d_55_v46_ = nw11_
            index8_ = default__.safeIndex(530, (d_55_v46_).length(0))
            (d_55_v46_)[index8_] = d_54_v45_
            d_56_v47_: _dafny.Seq
            d_56_v47_ = _dafny.SeqWithoutIsStrInference([d_53_v44_])
            index9_ = default__.safeIndex(530, (d_55_v46_).length(0))
            (d_55_v46_)[index9_] = ((d_56_v47_ if (d_51_v42_).f18 else d_54_v45_))
        d_57_v48_: C1
        nw12_ = C1()
        nw12_.ctor__(d_23_v15_, p1, (0) - (p1), d_23_v15_)
        d_57_v48_ = nw12_
        d_58_v49_: D4
        d_58_v49_ = D4_DC9(d_57_v48_)
        d_59_v50_: _dafny.Seq
        d_59_v50_ = _dafny.SeqWithoutIsStrInference([(d_58_v49_).cf13, d_57_v48_])
        d_60_v51_: _dafny.Array
        nw13_ = _dafny.Array(None, 20)
        nw13_[int(0)] = d_57_v48_
        nw13_[int(1)] = d_57_v48_
        nw13_[int(2)] = (d_59_v50_)[default__.safeIndex(p1, len(d_59_v50_))]
        nw13_[int(3)] = d_57_v48_
        nw13_[int(4)] = d_57_v48_
        nw13_[int(5)] = d_57_v48_
        nw13_[int(6)] = d_57_v48_
        nw13_[int(7)] = d_57_v48_
        nw13_[int(8)] = d_57_v48_
        nw13_[int(9)] = d_57_v48_
        nw13_[int(10)] = d_57_v48_
        nw13_[int(11)] = d_57_v48_
        nw13_[int(12)] = d_57_v48_
        nw13_[int(13)] = d_57_v48_
        nw13_[int(14)] = d_57_v48_
        nw13_[int(15)] = d_57_v48_
        nw13_[int(16)] = d_57_v48_
        nw13_[int(17)] = d_57_v48_
        nw13_[int(18)] = d_57_v48_
        nw13_[int(19)] = d_57_v48_
        d_60_v51_ = nw13_
        d_61_v52_: _dafny.Seq
        d_61_v52_ = _dafny.SeqWithoutIsStrInference([d_60_v51_, d_60_v51_, d_60_v51_])
        d_62_v53_: D6
        d_62_v53_ = D6_DC16(p1)
        d_63_v54_: _dafny.MultiSet
        d_63_v54_ = _dafny.MultiSet([p1, len(_dafny.SeqWithoutIsStrInference([p0 for d_64_i3_ in range(default__.abs(507))]))])
        d_65_v55_: D2
        d_65_v55_ = D2_DC6(d_63_v54_)
        d_66_v56_: C0
        nw14_ = C0()
        nw14_.ctor__(d_65_v55_, p1, d_23_v15_)
        d_66_v56_ = nw14_
        d_67_v57_: _dafny.Seq
        d_67_v57_ = _dafny.SeqWithoutIsStrInference([d_66_v56_])
        d_68_v58_: _dafny.Seq
        d_68_v58_ = _dafny.SeqWithoutIsStrInference([479, len(_dafny.Set({(d_62_v53_).cf22, len(d_67_v57_)}))])
        rhs7_ = (d_61_v52_)[default__.safeIndex(((d_68_v58_)[default__.safeIndex(212, len(d_68_v58_))]) + (p0), len(d_61_v52_))]
        rhs8_ = p0
        rhs9_ = (d_66_v56_).f23
        rhs10_ = default__.fm1(p0, (d_57_v48_.f21) or (d_23_v15_), (_dafny.MultiSet([d_57_v48_.f21])).intersection(_dafny.MultiSet([not(d_23_v15_)])), globalState)
        lhs7_ = globalState
        lhs8_ = globalState
        lhs9_ = globalState
        d_60_v51_ = rhs7_
        lhs7_.f9 = rhs8_
        lhs8_.f5 = rhs9_
        lhs9_.f13 = rhs10_
        if (default__.safeDivisionInt((d_66_v56_).f23, (0) - (len(default__.fm9(d_57_v48_.f21, True, globalState))))) >= (36):
            d_69_v59_: _dafny.Set
            d_69_v59_ = _dafny.Set({d_23_v15_, d_57_v48_.f21, d_23_v15_})
            d_70_v60_: D9
            d_70_v60_ = D9_DC23(d_69_v59_)
            pat_let_tv0_ = d_57_v48_
            pat_let_tv1_ = d_57_v48_
            def iife2_(_pat_let0_0):
                def iife3_(d_71_dt__update__tmp_h0_):
                    def iife4_(_pat_let1_0):
                        def iife5_(d_72_dt__update_hcf32_h0_):
                            return D9_DC23(d_72_dt__update_hcf32_h0_)
                        return iife5_(_pat_let1_0)
                    return iife4_(_dafny.Set({pat_let_tv0_.f21, pat_let_tv1_.f21}))
                return iife3_(_pat_let0_0)
            d_69_v59_ = (d_69_v59_) | ((iife2_(d_70_v60_)).cf32)
            (globalState).f3 = default__.safeModuloInt(default__.fm0(929, p1, 503, p0, globalState), 263)
            d_73_v61_: C1
            nw15_ = C1()
            nw15_.ctor__(d_57_v48_.f21, p0, 537, not((d_57_v48_).fm5(False, -758, globalState)))
            d_73_v61_ = nw15_
            d_74_v62_: _dafny.Map
            d_74_v62_ = _dafny.Map({d_66_v56_: d_69_v59_})
            d_75_v63_: D1
            d_75_v63_ = D1_DC3((len(d_74_v62_)) + (len(_dafny.SeqWithoutIsStrInference([d_73_v61_.f21, d_57_v48_.f21, d_73_v61_.f21]))))
            source0_ = d_75_v63_
            if source0_.is_DC3:
                d_76___mcc_h0_ = source0_.cf5
                d_77_cf5_ = d_76___mcc_h0_
                d_78_v64_: _dafny.Map
                d_78_v64_ = _dafny.Map({(d_66_v56_).f23: d_73_v61_.f21})
                d_78_v64_ = (d_78_v64_).set(p0, d_73_v61_.f21)
                d_79_v65_: _dafny.Array
                nw16_ = _dafny.Array(False, 6)
                d_79_v65_ = nw16_
                d_80_v66_: _dafny.Array
                nw17_ = _dafny.Array(None, 1)
                nw17_[int(0)] = (d_66_v56_).f23
                d_80_v66_ = nw17_
                d_81_v67_: _dafny.Map
                d_81_v67_ = _dafny.Map({d_79_v65_: d_80_v66_})
                d_82_v68_: D7
                d_82_v68_ = D7_DC18(d_81_v67_)
                d_82_v68_ = D7_DC18((d_81_v67_) | (d_81_v67_))
                d_23_v15_ = not(((d_66_v56_).f23) >= ((d_66_v56_).f23))
                (d_73_v61_).f21 = not(d_57_v48_.f21)
            elif source0_.is_DC4:
                d_83___mcc_h1_ = source0_.cf6
                d_84___mcc_h2_ = source0_.cf7
                d_85___mcc_h3_ = source0_.cf8
                d_86_cf8_ = d_85___mcc_h3_
                d_87_cf7_ = d_84___mcc_h2_
                d_88_cf6_ = d_83___mcc_h1_
                d_89_v70_: _dafny.Array
                def lambda0_(d_90_cf7_, d_91_v48_, d_92_v56_):
                    def lambda1_(d_93_i4_):
                        def iife6_():
                            coll2_ = _dafny.Map()
                            compr_2_: int
                            for compr_2_ in _dafny.IntegerRange(418, 981):
                                d_94_v69_: int = compr_2_
                                if ((418) <= (d_94_v69_)) and ((d_94_v69_) < (981)):
                                    coll2_[(d_94_v69_) - ((d_92_v56_).f23)] = d_91_v48_.f21
                            return _dafny.Map(coll2_)
                        return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqlnui"))): d_91_v48_.f21}) if d_90_cf7_ else iife6_()
                        )

                    return lambda1_

                init0_ = lambda0_(d_87_cf7_, d_57_v48_, d_66_v56_)
                nw18_ = _dafny.Array(None, 29)
                for i0_0_ in range(nw18_.length(0)):
                    nw18_[i0_0_] = init0_(i0_0_)
                d_89_v70_ = nw18_
                d_95_v71_: _dafny.Map
                d_95_v71_ = _dafny.Map({(d_66_v56_).f23: d_57_v48_.f21})
                index10_ = default__.safeIndex(918, (d_89_v70_).length(0))
                (d_89_v70_)[index10_] = d_95_v71_
                index11_ = default__.safeIndex(918, (d_89_v70_).length(0))
                (d_89_v70_)[index11_] = d_95_v71_
                d_96_v72_: _dafny.Map
                d_96_v72_ = _dafny.Map({d_57_v48_.f21: d_57_v48_.f21})
                d_97_v73_: C1
                nw19_ = C1()
                nw19_.ctor__(d_57_v48_.f21, (615) * (p1), len(d_96_v72_), d_57_v48_.f21)
                d_97_v73_ = nw19_
                (globalState).f9 = (d_66_v56_).f23
                d_98_v74_: _dafny.Seq
                d_98_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaky"))
                (globalState).f4 = d_98_v74_
            elif True:
                d_99___mcc_h4_ = source0_.cf4
                d_100_cf4_ = d_99___mcc_h4_
                d_101_v75_: int
                d_102_v76_: bool
                d_103_v77_: int
                d_104_v78_: int
                out0_: int
                out1_: bool
                out2_: int
                out3_: int
                out0_, out1_, out2_, out3_ = (d_57_v48_).m1(d_57_v48_.f21, d_73_v61_.f21, globalState)
                d_101_v75_ = out0_
                d_102_v76_ = out1_
                d_103_v77_ = out2_
                d_104_v78_ = out3_
                d_105_v79_: _dafny.Array
                nw20_ = _dafny.Array(False, 15)
                d_105_v79_ = nw20_
                d_106_v80_: _dafny.Seq
                d_106_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
                d_107_v81_: _dafny.Map
                d_107_v81_ = _dafny.Map({d_103_v77_: len(d_24_v16_)})
                d_108_v82_: _dafny.Map
                d_108_v82_ = _dafny.Map({d_57_v48_.f21: d_107_v81_})
                d_109_v83_: _dafny.Array
                nw21_ = _dafny.Array(None, 17)
                nw21_[int(0)] = (d_66_v56_).f23
                nw21_[int(1)] = p0
                nw21_[int(2)] = (d_66_v56_).f23
                nw21_[int(3)] = len(d_106_v80_)
                nw21_[int(4)] = (d_66_v56_).f23
                nw21_[int(5)] = len(d_108_v82_)
                nw21_[int(6)] = (0) - ((d_66_v56_).f23)
                nw21_[int(7)] = len(d_68_v58_)
                nw21_[int(8)] = (d_66_v56_).f23
                nw21_[int(9)] = d_103_v77_
                nw21_[int(10)] = (d_66_v56_).f23
                nw21_[int(11)] = p0
                nw21_[int(12)] = p0
                nw21_[int(13)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_110_i5_ in range(default__.abs(330))]))
                nw21_[int(14)] = d_103_v77_
                nw21_[int(15)] = (0) - ((d_66_v56_).f23)
                nw21_[int(16)] = d_103_v77_
                d_109_v83_ = nw21_
                d_111_v84_: _dafny.Map
                d_111_v84_ = _dafny.Map({d_105_v79_: d_109_v83_})
                d_112_v85_: D7
                d_112_v85_ = D7_DC18(d_111_v84_)
                d_112_v85_ = d_112_v85_
                d_113_v86_: D6
                d_113_v86_ = D6_DC17((d_66_v56_).f23, True)
                d_114_v87_: _dafny.Set
                d_114_v87_ = _dafny.Set({d_113_v86_})
                d_115_v88_: _dafny.Map
                d_115_v88_ = _dafny.Map({d_103_v77_: d_114_v87_})
                d_115_v88_ = (d_115_v88_).set(d_103_v77_, _dafny.Set({d_113_v86_, default__.fm17(globalState), D6_DC17(p1, d_23_v15_)}))
                d_116_v89_: _dafny.Map
                d_116_v89_ = _dafny.Map({d_109_v83_: d_106_v80_})
                (globalState).f5 = len((d_116_v89_) | ((d_116_v89_).set(d_109_v83_, d_106_v80_)))
            r0 = (d_66_v56_).f23
        elif True:
            d_117_v90_: _dafny.Array
            nw22_ = _dafny.Array(int(0), 1)
            d_117_v90_ = nw22_
            index12_ = default__.safeIndex(983, (d_117_v90_).length(0))
            (d_117_v90_)[index12_] = 250
            index13_ = default__.safeIndex(983, (d_117_v90_).length(0))
            (d_117_v90_)[index13_] = len(d_24_v16_)
            d_57_v48_ = d_57_v48_
            d_118_v91_: _dafny.Map
            d_118_v91_ = _dafny.Map({(d_66_v56_).f23: p1})
            d_119_v93_: D9
            def iife7_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(665, 12):
                    d_120_v92_: int = compr_3_
                    if ((665) <= (d_120_v92_)) and ((d_120_v92_) < (12)):
                        coll3_[(d_120_v92_) - (330)] = (d_66_v56_).f23
                return _dafny.Map(coll3_)
            d_119_v93_ = D9_DC24((d_118_v91_) != (iife7_()
), (0) - ((d_66_v56_).f23), not(d_57_v48_.f21))
            source1_ = d_119_v93_
            if source1_.is_DC24:
                d_121___mcc_h5_ = source1_.cf33
                d_122___mcc_h6_ = source1_.cf34
                d_123___mcc_h7_ = source1_.cf35
                d_124_cf35_ = d_123___mcc_h7_
                d_125_cf34_ = d_122___mcc_h6_
                d_126_cf33_ = d_121___mcc_h5_
                d_127_v94_: _dafny.Array
                def lambda2_(d_128_v48_):
                    def lambda3_(d_129_i6_):
                        return d_128_v48_.f21

                    return lambda3_

                init1_ = lambda2_(d_57_v48_)
                nw23_ = _dafny.Array(None, 27)
                for i0_1_ in range(nw23_.length(0)):
                    nw23_[i0_1_] = init1_(i0_1_)
                d_127_v94_ = nw23_
                d_130_v95_: _dafny.Map
                d_130_v95_ = _dafny.Map({d_57_v48_.f21: d_127_v94_})
                d_131_v96_: _dafny.Array
                nw24_ = _dafny.Array(None, 4)
                nw24_[int(0)] = d_127_v94_
                nw24_[int(1)] = d_127_v94_
                nw24_[int(2)] = ((d_130_v95_)[d_57_v48_.f21] if (d_57_v48_.f21) in (d_130_v95_) else d_127_v94_)
                nw24_[int(3)] = d_127_v94_
                d_131_v96_ = nw24_
                index14_ = default__.safeIndex(337, (d_131_v96_).length(0))
                (d_131_v96_)[index14_] = d_127_v94_
                d_132_v97_: _dafny.Seq
                d_132_v97_ = _dafny.SeqWithoutIsStrInference([(D10_DC25(d_127_v94_)).cf36])
                d_133_v98_: _dafny.Set
                d_133_v98_ = _dafny.Set({d_57_v48_.f21})
                index15_ = default__.safeIndex(337, (d_131_v96_).length(0))
                (d_131_v96_)[index15_] = (d_132_v97_)[default__.safeIndex(len(d_133_v98_), len(d_132_v97_))]
                d_134_v99_: _dafny.Seq
                d_134_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qggfdxe"))
                d_135_v100_: _dafny.MultiSet
                d_135_v100_ = _dafny.MultiSet([True])
                rhs11_ = _dafny.SeqWithoutIsStrInference([p0 for d_136_i7_ in range(default__.abs(663))])
                rhs12_ = (0) - (d_125_cf34_)
                rhs13_ = (d_24_v16_)[default__.safeIndex(((d_117_v90_)[default__.safeIndex(983, (d_117_v90_).length(0))]) - (len(d_134_v99_)), len(d_24_v16_))]
                rhs14_ = (d_117_v90_)[default__.safeIndex(983, (d_117_v90_).length(0))]
                rhs15_ = ((d_135_v100_)[(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ursmty")))) < ((d_66_v56_).f23)] if ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ursmty")))) < ((d_66_v56_).f23)) in (d_135_v100_) else ((d_66_v56_).f23) * ((d_66_v56_).f23))
                lhs10_ = globalState
                lhs11_ = d_57_v48_
                lhs12_ = globalState
                d_68_v58_ = rhs11_
                lhs10_.f3 = rhs12_
                lhs11_.f21 = rhs13_
                lhs12_.f3 = rhs14_
                d_125_cf34_ = rhs15_
                d_137_v101_: _dafny.Map
                d_137_v101_ = _dafny.Map({d_126_cf33_: len(_dafny.Map({(_dafny.MultiSet([(d_66_v56_).f23])).cardinality: d_57_v48_}))})
                d_124_cf35_ = ((len(d_137_v101_)) - ((d_135_v100_).cardinality)) <= ((d_117_v90_)[default__.safeIndex(983, (d_117_v90_).length(0))])
                d_138_v102_: _dafny.Map
                d_138_v102_ = _dafny.Map({(d_117_v90_)[default__.safeIndex(983, (d_117_v90_).length(0))]: d_134_v99_})
                index16_ = default__.safeIndex(983, (d_117_v90_).length(0))
                def iife8_():
                    coll4_ = _dafny.Map()
                    compr_4_: int
                    for compr_4_ in _dafny.IntegerRange(410, 451):
                        d_139_v103_: int = compr_4_
                        if ((410) <= (d_139_v103_)) and ((d_139_v103_) < (451)):
                            coll4_[(d_139_v103_) + (169)] = d_134_v99_
                    return _dafny.Map(coll4_)
                rhs16_ = d_131_v96_
                rhs17_ = default__.safeDivisionInt(len((d_132_v97_) + (d_132_v97_)), (0) - (len((d_138_v102_) | (iife8_()
                ))))
                lhs13_ = d_117_v90_
                lhs14_ = default__.safeIndex(983, (d_117_v90_).length(0))
                d_131_v96_ = rhs16_
                lhs13_[lhs14_] = rhs17_
            elif True:
                d_140___mcc_h8_ = source1_.cf32
                d_141_cf32_ = d_140___mcc_h8_
                d_142_v104_: _dafny.Array
                def lambda4_(d_143_v56_, d_144_v48_):
                    def lambda5_(d_145_i8_):
                        return _dafny.Map({D5_DC13((d_143_v56_).f23): d_144_v48_.f21})

                    return lambda5_

                init2_ = lambda4_(d_66_v56_, d_57_v48_)
                nw25_ = _dafny.Array(None, 12)
                for i0_2_ in range(nw25_.length(0)):
                    nw25_[i0_2_] = init2_(i0_2_)
                d_142_v104_ = nw25_
                d_146_v105_: D5
                d_146_v105_ = D5_DC13(p0)
                d_147_v106_: _dafny.Map
                d_147_v106_ = _dafny.Map({d_146_v105_: d_57_v48_.f21})
                index17_ = default__.safeIndex(534, (d_142_v104_).length(0))
                (d_142_v104_)[index17_] = d_147_v106_
                d_148_v108_: _dafny.Seq
                d_148_v108_ = _dafny.SeqWithoutIsStrInference([d_146_v105_, default__.fm18(d_23_v15_, globalState), d_146_v105_])
                pat_let_tv2_ = d_66_v56_
                index18_ = default__.safeIndex(534, (d_142_v104_).length(0))
                def iife9_():
                    coll5_ = _dafny.Map()
                    compr_5_: D5
                    for compr_5_ in (d_148_v108_).Elements:
                        d_149_v107_: D5 = compr_5_
                        if (d_149_v107_) in (d_148_v108_):
                            coll5_[d_149_v107_] = d_23_v15_
                    return _dafny.Map(coll5_)
                def iife10_(_pat_let2_0):
                    def iife11_(d_150_dt__update__tmp_h1_):
                        def iife12_(_pat_let3_0):
                            def iife13_(d_151_dt__update_hcf19_h0_):
                                return D5_DC13(d_151_dt__update_hcf19_h0_)
                            return iife13_(_pat_let3_0)
                        return iife12_((pat_let_tv2_).f23)
                    return iife11_(_pat_let2_0)
                rhs18_ = d_57_v48_.f21
                rhs19_ = d_57_v48_.f21
                rhs20_ = (iife9_()
                ).set(iife10_(d_146_v105_), d_57_v48_.f21)
                lhs15_ = globalState
                lhs16_ = d_142_v104_
                lhs17_ = default__.safeIndex(534, (d_142_v104_).length(0))
                lhs15_.f13 = rhs18_
                d_23_v15_ = rhs19_
                lhs16_[lhs17_] = rhs20_
                nw26_ = C1()
                nw26_.ctor__(d_57_v48_.f21, (d_66_v56_).f23, (d_66_v56_).f23, (d_57_v48_.f21 if d_57_v48_.f21 else d_57_v48_.f21))
                d_57_v48_ = nw26_
                (globalState).f13 = d_57_v48_.f21
                (globalState).f9 = (d_66_v56_).f23
            d_152_v109_: _dafny.Set
            d_152_v109_ = _dafny.Set({d_66_v56_, d_66_v56_, d_66_v56_})
            index19_ = default__.safeIndex(983, (d_117_v90_).length(0))
            (d_117_v90_)[index19_] = (0) - (len(d_152_v109_))
            if d_23_v15_:
                d_153_v110_: _dafny.Map
                d_153_v110_ = _dafny.Map({(d_66_v56_).f23: d_57_v48_})
                d_153_v110_ = (d_153_v110_) | (d_153_v110_)
                d_154_v111_: _dafny.Map
                d_154_v111_ = _dafny.Map({d_23_v15_: p0})
                d_154_v111_ = d_154_v111_
                (globalState).f13 = d_57_v48_.f21
                d_155_v112_: T1
                nw27_ = C1()
                nw27_.ctor__(d_57_v48_.f21, (d_66_v56_).f23, (d_66_v56_).f23, True)
                d_155_v112_ = nw27_
                d_156_v113_: str
                d_156_v113_ = _dafny.CodePoint('t')
                d_157_v114_: _dafny.Map
                d_157_v114_ = _dafny.Map({d_57_v48_.f21: d_57_v48_})
                d_158_v115_: _dafny.MultiSet
                d_158_v115_ = _dafny.MultiSet([d_117_v90_, d_117_v90_])
                rhs21_ = d_155_v112_
                rhs22_ = not (d_57_v48_.f21) or ((d_155_v112_).f18)
                rhs23_ = d_156_v113_
                rhs24_ = ((d_157_v114_)[(_dafny.MultiSet([d_117_v90_])).ispropersubset(d_158_v115_)] if ((_dafny.MultiSet([d_117_v90_])).ispropersubset(d_158_v115_)) in (d_157_v114_) else d_57_v48_)
                lhs18_ = globalState
                d_155_v112_ = rhs21_
                lhs18_.f13 = rhs22_
                d_156_v113_ = rhs23_
                d_57_v48_ = rhs24_
                d_159_v116_: _dafny.Array
                nw28_ = _dafny.Array(None, 6)
                nw28_[int(0)] = d_156_v113_
                nw28_[int(1)] = d_156_v113_
                nw28_[int(2)] = d_156_v113_
                nw28_[int(3)] = d_156_v113_
                nw28_[int(4)] = d_156_v113_
                nw28_[int(5)] = d_156_v113_
                d_159_v116_ = nw28_
                index20_ = default__.safeIndex(388, (d_159_v116_).length(0))
                (d_159_v116_)[index20_] = d_156_v113_
                d_160_v117_: D2
                d_160_v117_ = D2_DC6(_dafny.MultiSet([789]))
                d_161_v118_: D2
                d_161_v118_ = D2_DC7(d_160_v117_)
                index21_ = default__.safeIndex(388, (d_159_v116_).length(0))
                (d_159_v116_)[index21_] = default__.fm15(d_161_v118_, (d_57_v48_).fm5((d_155_v112_).f18, default__.fm0(760, p0, (d_66_v56_).f23, 189, globalState), globalState), globalState)
            elif True:
                d_162_v119_: str
                d_162_v119_ = _dafny.CodePoint('k')
                d_163_v120_: D2
                d_163_v120_ = D2_DC6(d_63_v54_)
                d_164_v121_: D2
                d_164_v121_ = D2_DC7(D2_DC7(d_163_v120_))
                d_165_v122_: _dafny.Array
                nw29_ = _dafny.Array(None, 6)
                nw29_[int(0)] = d_162_v119_
                nw29_[int(1)] = d_162_v119_
                nw29_[int(2)] = default__.fm15(d_164_v121_, d_57_v48_.f21, globalState)
                nw29_[int(3)] = d_162_v119_
                nw29_[int(4)] = (d_162_v119_ if d_57_v48_.f21 else d_162_v119_)
                nw29_[int(5)] = d_162_v119_
                d_165_v122_ = nw29_
                index22_ = default__.safeIndex(751, (d_165_v122_).length(0))
                (d_165_v122_)[index22_] = d_162_v119_
                d_166_v123_: _dafny.Map
                d_166_v123_ = _dafny.Map({True: d_57_v48_})
                d_167_v124_: T1
                nw30_ = C1()
                nw30_.ctor__((d_57_v48_).fm5(d_57_v48_.f21, p1, globalState), (_dafny.MultiSet(d_24_v16_)).cardinality, len(d_166_v123_), d_57_v48_.f21)
                d_167_v124_ = nw30_
                d_168_v125_: D10
                d_168_v125_ = D10_DC26(d_167_v124_)
                d_169_v126_: _dafny.Array
                nw31_ = _dafny.Array(None, 4)
                nw31_[int(0)] = d_168_v125_
                nw31_[int(1)] = D10_DC26(d_167_v124_)
                nw31_[int(2)] = d_168_v125_
                nw31_[int(3)] = D10_DC26(d_167_v124_)
                d_169_v126_ = nw31_
                d_170_v127_: _dafny.Array
                nw32_ = _dafny.Array(None, 3)
                nw32_[int(0)] = d_117_v90_
                nw32_[int(1)] = d_117_v90_
                nw32_[int(2)] = d_117_v90_
                d_170_v127_ = nw32_
                index23_ = default__.safeIndex(599, (d_170_v127_).length(0))
                (d_170_v127_)[index23_] = d_117_v90_
                d_171_v128_: D11
                d_171_v128_ = D11_DC29(d_117_v90_)
                index24_ = default__.safeIndex(751, (d_165_v122_).length(0))
                index25_ = default__.safeIndex(599, (d_170_v127_).length(0))
                rhs25_ = d_68_v58_
                rhs26_ = d_162_v119_
                rhs27_ = d_169_v126_
                rhs28_ = (d_171_v128_).cf40
                lhs19_ = d_165_v122_
                lhs20_ = default__.safeIndex(751, (d_165_v122_).length(0))
                lhs21_ = d_170_v127_
                lhs22_ = default__.safeIndex(599, (d_170_v127_).length(0))
                d_68_v58_ = rhs25_
                lhs19_[lhs20_] = rhs26_
                d_169_v126_ = rhs27_
                lhs21_[lhs22_] = rhs28_
                d_172_v129_: int
                d_173_v130_: bool
                d_174_v131_: int
                d_175_v132_: int
                out4_: int
                out5_: bool
                out6_: int
                out7_: int
                out4_, out5_, out6_, out7_ = (d_167_v124_).m1((d_24_v16_)[default__.safeIndex(d_167_v124_.f20, len(d_24_v16_))], d_23_v15_, globalState)
                d_172_v129_ = out4_
                d_173_v130_ = out5_
                d_174_v131_ = out6_
                d_175_v132_ = out7_
                d_176_v133_: _dafny.MultiSet
                d_176_v133_ = _dafny.MultiSet([d_57_v48_.f21, (d_167_v124_).f18])
                d_177_v134_: _dafny.Seq
                d_177_v134_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rihyj"))
                d_178_v135_: _dafny.Map
                d_178_v135_ = _dafny.Map({default__.fm0((d_176_v133_).cardinality, 600, -288, len(_dafny.SeqWithoutIsStrInference([d_162_v119_ for d_179_i9_ in range(default__.abs(684))])), globalState): ((d_177_v134_).set(default__.safeIndex((d_117_v90_)[default__.safeIndex(983, (d_117_v90_).length(0))], len(d_177_v134_)), (d_165_v122_)[default__.safeIndex(751, (d_165_v122_).length(0))])) + (d_177_v134_)})
                d_178_v135_ = (d_178_v135_).set(((d_176_v133_)[d_57_v48_.f21] if (d_57_v48_.f21) in (d_176_v133_) else d_174_v131_), d_177_v134_)
                d_180_v136_: T0
                nw33_ = C0()
                nw33_.ctor__(d_66_v56_.f22, p1, d_173_v130_)
                d_180_v136_ = nw33_
                d_181_v137_: D6
                d_181_v137_ = D6_DC15(d_180_v136_)
                d_182_v138_: _dafny.Set
                d_182_v138_ = _dafny.Set({(d_181_v137_).cf21})
                d_183_v139_: _dafny.Map
                d_183_v139_ = _dafny.Map({len(d_182_v138_): d_57_v48_.f21})
                (globalState).f3 = ((0) - ((d_68_v58_)[default__.safeIndex(len(d_183_v139_), len(d_68_v58_))])) + (d_174_v131_)
                d_184_v140_: _dafny.Array
                def lambda6_(d_185_v124_):
                    def lambda7_(d_186_i10_):
                        return not(not((d_185_v124_).f18))

                    return lambda7_

                init3_ = lambda6_(d_167_v124_)
                nw34_ = _dafny.Array(None, 15)
                for i0_3_ in range(nw34_.length(0)):
                    nw34_[i0_3_] = init3_(i0_3_)
                d_184_v140_ = nw34_
                d_187_v141_: _dafny.Map
                d_187_v141_ = _dafny.Map({d_184_v140_: len(d_118_v91_)})
                d_188_v142_: _dafny.Map
                d_188_v142_ = _dafny.Map({False: d_187_v141_})
                rhs29_ = (d_24_v16_)[default__.safeIndex(d_167_v124_.f20, len(d_24_v16_))]
                rhs30_ = d_24_v16_
                rhs31_ = (d_187_v141_) == ((((d_188_v142_)[d_57_v48_.f21] if (d_57_v48_.f21) in (d_188_v142_) else _dafny.Map({d_184_v140_: d_175_v132_}))) | (_dafny.Map({d_184_v140_: default__.fm0(d_174_v131_, (d_66_v56_).f23, len(d_24_v16_), d_167_v124_.f20, globalState)})))
                rhs32_ = (d_180_v136_).f18
                rhs33_ = (d_172_v129_) - ((d_66_v56_).f23)
                lhs23_ = d_57_v48_
                lhs24_ = d_57_v48_
                lhs25_ = globalState
                lhs26_ = globalState
                lhs23_.f21 = rhs29_
                d_24_v16_ = rhs30_
                lhs24_.f21 = rhs31_
                lhs25_.f13 = rhs32_
                lhs26_.f9 = rhs33_
        d_189_v143_: _dafny.Map
        d_189_v143_ = _dafny.Map({p1: 366})
        d_190_v144_: _dafny.MultiSet
        d_190_v144_ = _dafny.MultiSet([d_189_v143_, d_189_v143_])
        d_191_v145_: D7
        d_191_v145_ = D7_DC20(d_190_v144_)
        pat_let_tv3_ = d_57_v48_
        pat_let_tv4_ = d_190_v144_
        pat_let_tv5_ = d_190_v144_
        def lambda8_(source2_):
            if source2_.is_DC19:
                d_192___mcc_h9_ = source2_.cf26
                d_193___mcc_h10_ = source2_.cf27
                d_194___mcc_h11_ = source2_.cf28
                d_195_cf28_ = d_194___mcc_h11_
                d_196_cf27_ = d_193___mcc_h10_
                d_197_cf26_ = d_192___mcc_h9_
                return d_195_cf28_
            elif source2_.is_DC20:
                d_198___mcc_h12_ = source2_.cf29
                d_199_cf29_ = d_198___mcc_h12_
                return True
            elif source2_.is_DC18:
                d_200___mcc_h13_ = source2_.cf25
                d_201_cf25_ = d_200___mcc_h13_
                return pat_let_tv3_.f21
            elif True:
                d_202___mcc_h14_ = source2_.cf30
                d_203_cf30_ = d_202___mcc_h14_
                return (pat_let_tv4_).issubset(pat_let_tv5_)

        (d_57_v48_).f21 = lambda8_(d_191_v145_)
        d_204_v146_: _dafny.Seq
        d_204_v146_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
        if not(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ev"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlhryx")))) < (d_204_v146_)):
            if d_57_v48_.f21:
                d_205_v147_: _dafny.Map
                d_205_v147_ = _dafny.Map({default__.fm1((_dafny.MultiSet([True])).cardinality, not(not(d_57_v48_.f21)), _dafny.MultiSet([d_23_v15_]), globalState): len(d_68_v58_)})
                d_206_v148_: T1
                nw35_ = C1()
                nw35_.ctor__(d_23_v15_, p1, (d_68_v58_)[default__.safeIndex((d_66_v56_).f23, len(d_68_v58_))], False)
                d_206_v148_ = nw35_
                d_207_v149_: D10
                d_207_v149_ = D10_DC26(d_206_v148_)
                d_208_v150_: _dafny.Map
                d_208_v150_ = _dafny.Map({d_205_v147_: d_207_v149_})
                d_208_v150_ = (d_208_v150_).set(d_205_v147_, d_207_v149_)
                d_209_v151_: D4
                d_209_v151_ = D4_DC10(d_57_v48_.f21, ((d_189_v143_)[(0) - (p1)] if ((0) - (p1)) in (d_189_v143_) else -796), p1)
                d_210_v152_: D1
                d_210_v152_ = D1_DC3((d_209_v151_).cf15)
                d_211_v153_: _dafny.MultiSet
                d_211_v153_ = _dafny.MultiSet([d_210_v152_, d_210_v152_])
                d_212_v154_: D9
                d_212_v154_ = D9_DC24(False, (d_206_v148_).f19, d_57_v48_.f21)
                rhs34_ = (((d_206_v148_).f19) - (p1) if d_57_v48_.f21 else (d_206_v148_).f19)
                rhs35_ = (d_212_v154_).cf35
                rhs36_ = default__.fm19(d_206_v148_.f20, len((d_68_v58_) + (d_68_v58_)), len(_dafny.SeqWithoutIsStrInference([(d_66_v56_).f23 for d_213_i11_ in range(default__.abs(-197))])), globalState)
                rhs37_ = d_57_v48_.f21
                rhs38_ = d_57_v48_.f21
                lhs27_ = globalState
                lhs28_ = d_57_v48_
                lhs29_ = globalState
                lhs30_ = d_57_v48_
                lhs27_.f3 = rhs34_
                lhs28_.f21 = rhs35_
                d_211_v153_ = rhs36_
                lhs29_.f13 = rhs37_
                lhs30_.f21 = rhs38_
                d_57_v48_ = d_57_v48_
                d_214_v155_: _dafny.Array
                def lambda9_(d_215_i12_):
                    return default__.safeDivisionInt(d_215_i12_, 110)

                init4_ = lambda9_
                nw36_ = _dafny.Array(None, 15)
                for i0_4_ in range(nw36_.length(0)):
                    nw36_[i0_4_] = init4_(i0_4_)
                d_214_v155_ = nw36_
                index26_ = default__.safeIndex(754, (d_214_v155_).length(0))
                (d_214_v155_)[index26_] = p0
                index27_ = default__.safeIndex(754, (d_214_v155_).length(0))
                (d_214_v155_)[index27_] = (-685) + (d_206_v148_.f20)
                (globalState).f13 = (d_206_v148_).fm5(d_57_v48_.f21, (d_206_v148_).f19, globalState)
            elif True:
                d_57_v48_ = d_57_v48_
                (globalState).f13 = d_23_v15_
                d_216_v156_: _dafny.MultiSet
                d_216_v156_ = _dafny.MultiSet([d_23_v15_, d_57_v48_.f21])
                d_217_v157_: C1
                nw37_ = C1()
                nw37_.ctor__((d_216_v156_) == (d_216_v156_), (0) - (((d_66_v56_).f23) * (len(d_204_v146_))), p0, (d_57_v48_.f21 if d_57_v48_.f21 else d_57_v48_.f21))
                d_217_v157_ = nw37_
                d_218_v158_: _dafny.Array
                nw38_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_218_v158_ = nw38_
                d_218_v158_ = d_218_v158_
                d_219_v159_: str
                d_219_v159_ = _dafny.CodePoint('k')
                d_220_v160_: _dafny.Set
                d_220_v160_ = _dafny.Set({d_57_v48_, d_217_v157_, d_57_v48_, d_57_v48_, d_57_v48_})
                d_219_v159_ = (d_219_v159_ if d_57_v48_.f21 else ((d_204_v146_)[default__.safeIndex(len(d_220_v160_), len(d_204_v146_))] if d_217_v157_.f21 else d_219_v159_))
            (globalState).f9 = (p0 if (263) < (len(d_204_v146_)) else (d_66_v56_).f23)
            if (p1) <= (914):
                d_221_v161_: int
                d_222_v162_: bool
                d_223_v163_: int
                d_224_v164_: int
                out8_: int
                out9_: bool
                out10_: int
                out11_: int
                out8_, out9_, out10_, out11_ = (d_57_v48_).m1((d_24_v16_)[default__.safeIndex((d_66_v56_).f23, len(d_24_v16_))], True, globalState)
                d_221_v161_ = out8_
                d_222_v162_ = out9_
                d_223_v163_ = out10_
                d_224_v164_ = out11_
                d_225_v165_: str
                d_225_v165_ = _dafny.CodePoint('x')
                d_225_v165_ = d_225_v165_
                d_226_v166_: _dafny.Map
                d_226_v166_ = _dafny.Map({d_57_v48_.f21: len(d_204_v146_)})
                d_227_v167_: _dafny.MultiSet
                d_227_v167_ = _dafny.MultiSet([d_57_v48_.f21, d_23_v15_])
                d_228_v168_: C1
                nw39_ = C1()
                nw39_.ctor__(default__.fm1(len(d_226_v166_), d_57_v48_.f21, d_227_v167_, globalState), (0) - ((d_66_v56_).f23), p0, default__.fm1((d_66_v56_).f23, d_222_v162_, d_227_v167_, globalState))
                d_228_v168_ = nw39_
                d_229_v169_: _dafny.Array
                def lambda10_(d_230_v164_, d_231_v15_):
                    def lambda11_(d_232_i13_):
                        return (d_232_i13_) + (len(_dafny.Map({d_230_v164_: d_231_v15_})))

                    return lambda11_

                init5_ = lambda10_(d_224_v164_, d_23_v15_)
                nw40_ = _dafny.Array(None, 6)
                for i0_5_ in range(nw40_.length(0)):
                    nw40_[i0_5_] = init5_(i0_5_)
                d_229_v169_ = nw40_
                index28_ = default__.safeIndex(785, (d_229_v169_).length(0))
                (d_229_v169_)[index28_] = (d_66_v56_).f23
                d_233_v170_: _dafny.MultiSet
                d_233_v170_ = _dafny.MultiSet([d_68_v58_])
                d_234_v171_: _dafny.Array
                def lambda12_(d_235_i14_):
                    return False

                init6_ = lambda12_
                nw41_ = _dafny.Array(None, 7)
                for i0_6_ in range(nw41_.length(0)):
                    nw41_[i0_6_] = init6_(i0_6_)
                d_234_v171_ = nw41_
                d_236_v172_: _dafny.Map
                d_236_v172_ = _dafny.Map({d_234_v171_: d_57_v48_.f21})
                d_237_v173_: _dafny.Set
                d_237_v173_ = _dafny.Set({_dafny.MultiSet([not(d_23_v15_)]), d_227_v167_, d_227_v167_, d_227_v167_, d_227_v167_})
                index29_ = default__.safeIndex(785, (d_229_v169_).length(0))
                (d_229_v169_)[index29_] = ((d_233_v170_).cardinality) - (default__.fm0(len(d_204_v146_), 785, (0) - (len(d_236_v172_)), len(d_237_v173_), globalState))
                d_238_v174_: _dafny.Array
                nw42_ = _dafny.Array(_dafny.Set({}), 17)
                d_238_v174_ = nw42_
                d_239_v175_: _dafny.Set
                d_239_v175_ = _dafny.Set({(d_66_v56_).f23})
                index30_ = default__.safeIndex(326, (d_238_v174_).length(0))
                (d_238_v174_)[index30_] = d_239_v175_
                index31_ = default__.safeIndex(326, (d_238_v174_).length(0))
                (d_238_v174_)[index31_] = (d_239_v175_) - (d_239_v175_)
            elif True:
                d_240_v176_: _dafny.Array
                nw43_ = _dafny.Array(None, 7)
                nw43_[int(0)] = not(((d_68_v58_)[default__.safeIndex(len(d_204_v146_), len(d_68_v58_))]) != ((d_66_v56_).f23))
                nw43_[int(1)] = not(d_57_v48_.f21)
                nw43_[int(2)] = d_57_v48_.f21
                nw43_[int(3)] = d_57_v48_.f21
                nw43_[int(4)] = d_57_v48_.f21
                nw43_[int(5)] = d_23_v15_
                nw43_[int(6)] = d_57_v48_.f21
                d_240_v176_ = nw43_
                index32_ = default__.safeIndex(329, (d_240_v176_).length(0))
                (d_240_v176_)[index32_] = not (d_57_v48_.f21) or (d_57_v48_.f21)
                index33_ = default__.safeIndex(329, (d_240_v176_).length(0))
                (d_240_v176_)[index33_] = d_23_v15_
                d_241_v177_: C0
                nw44_ = C0()
                nw44_.ctor__(D2_DC6(d_63_v54_), (len(d_204_v146_)) + ((d_66_v56_).f23), d_57_v48_.f21)
                d_241_v177_ = nw44_
                d_242_v178_: D5
                d_242_v178_ = D5_DC13((d_66_v56_).f23)
                pat_let_tv6_ = d_241_v177_
                pat_let_tv7_ = p0
                pat_let_tv8_ = d_66_v56_
                pat_let_tv9_ = d_66_v56_
                pat_let_tv10_ = globalState
                def iife14_(_pat_let4_0):
                    def iife15_(d_243_dt__update__tmp_h2_):
                        def iife16_(_pat_let5_0):
                            def iife17_(d_244_dt__update_hcf19_h1_):
                                return D5_DC13(d_244_dt__update_hcf19_h1_)
                            return iife17_(_pat_let5_0)
                        return iife16_((0) - (default__.fm0((pat_let_tv6_).f23, pat_let_tv7_, (pat_let_tv8_).f23, (pat_let_tv9_).f23, pat_let_tv10_)))
                    return iife15_(_pat_let4_0)
                d_242_v178_ = iife14_(d_242_v178_)
                d_245_v179_: _dafny.Seq
                d_245_v179_ = _dafny.SeqWithoutIsStrInference([d_241_v177_])
                d_245_v179_ = d_245_v179_
                d_246_v180_: int
                d_247_v181_: bool
                d_248_v182_: int
                d_249_v183_: int
                out12_: int
                out13_: bool
                out14_: int
                out15_: int
                out12_, out13_, out14_, out15_ = (d_57_v48_).m1(d_23_v15_, d_57_v48_.f21, globalState)
                d_246_v180_ = out12_
                d_247_v181_ = out13_
                d_248_v182_ = out14_
                d_249_v183_ = out15_
            d_250_v184_: _dafny.MultiSet
            d_250_v184_ = _dafny.MultiSet([d_57_v48_.f21])
            d_251_v185_: _dafny.Set
            d_251_v185_ = _dafny.Set({d_57_v48_.f21, d_23_v15_})
            d_252_v186_: C1
            nw45_ = C1()
            nw45_.ctor__(not((default__.fm1((d_66_v56_).f23, d_57_v48_.f21, d_250_v184_, globalState)) == (d_57_v48_.f21)), (366) + (p1), (d_66_v56_).f23, (d_251_v185_).ispropersubset(d_251_v185_))
            d_252_v186_ = nw45_
            d_253_v187_: D6
            d_253_v187_ = D6_DC17(((d_63_v54_)[123] if (123) in (d_63_v54_) else (d_66_v56_).f23), d_252_v186_.f21)
            d_254_v188_: C1
            nw46_ = C1()
            nw46_.ctor__(d_57_v48_.f21, (d_253_v187_).cf23, p1, (len(d_204_v146_)) == (p0))
            d_254_v188_ = nw46_
        elif True:
            d_23_v15_ = (d_57_v48_.f21) or (True)
            if d_57_v48_.f21:
                d_255_v189_: _dafny.MultiSet
                d_255_v189_ = _dafny.MultiSet([d_57_v48_.f21])
                d_256_v190_: C1
                nw47_ = C1()
                nw47_.ctor__(d_57_v48_.f21, default__.safeModuloInt((d_255_v189_).cardinality, p0), (d_66_v56_).f23, not(d_57_v48_.f21))
                d_256_v190_ = nw47_
                (d_57_v48_).f21 = False
                d_257_v191_: _dafny.Map
                d_257_v191_ = _dafny.Map({d_57_v48_.f21: p1})
                d_258_v192_: _dafny.Map
                d_258_v192_ = _dafny.Map({(d_57_v48_).fm5(d_57_v48_.f21, 308, globalState): d_257_v191_})
                d_258_v192_ = (d_258_v192_).set(d_57_v48_.f21, _dafny.Map({d_57_v48_.f21: p0}))
                (d_256_v190_).f21 = d_256_v190_.f21
                d_259_v193_: _dafny.Map
                d_259_v193_ = _dafny.Map({p1: (D6_DC14(d_66_v56_)).cf20})
                d_260_v194_: T1
                nw48_ = C1()
                nw48_.ctor__(False, (d_66_v56_).f23, ((d_255_v189_).set(d_57_v48_.f21, default__.abs((d_66_v56_).f23))).cardinality, d_57_v48_.f21)
                d_260_v194_ = nw48_
                d_261_v195_: _dafny.Map
                d_261_v195_ = _dafny.Map({d_260_v194_: d_57_v48_.f21})
                d_262_v196_: _dafny.Map
                d_262_v196_ = _dafny.Map({d_66_v56_: (d_67_v57_)[default__.safeIndex((d_66_v56_).f23, len(d_67_v57_))]})
                d_263_v197_: _dafny.Array
                nw49_ = _dafny.Array(None, 19)
                nw49_[int(0)] = ((d_259_v193_)[len(d_261_v195_)] if (len(d_261_v195_)) in (d_259_v193_) else d_66_v56_)
                nw49_[int(1)] = d_66_v56_
                nw49_[int(2)] = d_66_v56_
                nw49_[int(3)] = d_66_v56_
                nw49_[int(4)] = d_66_v56_
                nw49_[int(5)] = ((d_262_v196_)[d_66_v56_] if (d_66_v56_) in (d_262_v196_) else d_66_v56_)
                nw49_[int(6)] = (d_66_v56_ if d_57_v48_.f21 else d_66_v56_)
                nw49_[int(7)] = d_66_v56_
                nw49_[int(8)] = d_66_v56_
                nw49_[int(9)] = d_66_v56_
                nw49_[int(10)] = d_66_v56_
                nw49_[int(11)] = d_66_v56_
                nw49_[int(12)] = (d_67_v57_)[default__.safeIndex((d_260_v194_).f19, len(d_67_v57_))]
                nw49_[int(13)] = d_66_v56_
                nw49_[int(14)] = d_66_v56_
                nw49_[int(15)] = d_66_v56_
                nw49_[int(16)] = d_66_v56_
                nw49_[int(17)] = d_66_v56_
                nw49_[int(18)] = d_66_v56_
                d_263_v197_ = nw49_
                index34_ = default__.safeIndex(12, (d_263_v197_).length(0))
                (d_263_v197_)[index34_] = d_66_v56_
                index35_ = default__.safeIndex(12, (d_263_v197_).length(0))
                nw50_ = C0()
                nw50_.ctor__(d_66_v56_.f22, (d_260_v194_).f19, not(d_23_v15_))
                (d_263_v197_)[index35_] = nw50_
            elif True:
                d_264_v198_: T0
                nw51_ = C0()
                nw51_.ctor__(d_65_v55_, p1, d_57_v48_.f21)
                d_264_v198_ = nw51_
                d_264_v198_ = d_264_v198_
                (globalState).f3 = (0) - ((d_66_v56_).f23)
                d_265_v199_: _dafny.Map
                d_265_v199_ = _dafny.Map({p0: d_68_v58_})
                d_265_v199_ = (d_265_v199_).set(default__.safeModuloInt(p0, (d_66_v56_).f23), (_dafny.SeqWithoutIsStrInference([(d_66_v56_).f23])).set(default__.safeIndex((d_66_v56_).f23, len(_dafny.SeqWithoutIsStrInference([(d_66_v56_).f23]))), p1))
                d_68_v58_ = _dafny.SeqWithoutIsStrInference([(d_66_v56_).f23, (d_66_v56_).f23])
                d_266_v200_: _dafny.Array
                def lambda13_(d_267_v48_, d_268_v198_):
                    def lambda14_(d_269_i15_):
                        return (_dafny.MultiSet([d_267_v48_.f21, d_267_v48_.f21])).issubset(_dafny.MultiSet([(d_268_v198_).f18]))

                    return lambda14_

                init7_ = lambda13_(d_57_v48_, d_264_v198_)
                nw52_ = _dafny.Array(None, 11)
                for i0_7_ in range(nw52_.length(0)):
                    nw52_[i0_7_] = init7_(i0_7_)
                d_266_v200_ = nw52_
                index36_ = default__.safeIndex(802, (d_266_v200_).length(0))
                (d_266_v200_)[index36_] = d_57_v48_.f21
                d_270_v201_: _dafny.Set
                d_270_v201_ = _dafny.Set({d_266_v200_})
                d_271_v202_: _dafny.Set
                d_271_v202_ = d_270_v201_
                d_272_v203_: T1
                nw53_ = C1()
                nw53_.ctor__(d_57_v48_.f21, p1, (d_66_v56_).f23, True)
                d_272_v203_ = nw53_
                d_273_v204_: _dafny.Map
                d_273_v204_ = _dafny.Map({d_272_v203_: (d_272_v203_).f19})
                d_274_v205_: _dafny.Seq
                d_274_v205_ = _dafny.SeqWithoutIsStrInference([d_272_v203_, d_272_v203_, d_272_v203_])
                d_275_v206_: _dafny.Map
                d_275_v206_ = _dafny.Map({d_57_v48_.f21: (d_273_v204_).set((d_274_v205_)[default__.safeIndex((d_66_v56_).f23, len(d_274_v205_))], p1)})
                d_276_v207_: _dafny.Map
                d_276_v207_ = _dafny.Map({len(((d_275_v206_)[d_57_v48_.f21] if (d_57_v48_.f21) in (d_275_v206_) else d_273_v204_)): d_189_v143_})
                index37_ = default__.safeIndex(802, (d_266_v200_).length(0))
                rhs39_ = not(((_dafny.Set({d_266_v200_, d_266_v200_, d_266_v200_, d_266_v200_, d_266_v200_}) if d_57_v48_.f21 else (d_271_v202_))).isdisjoint((_dafny.Set({d_266_v200_})) - (d_270_v201_)))
                rhs40_ = p1
                rhs41_ = not((default__.safeModuloInt(p1, p0)) >= (len(d_276_v207_)))
                lhs31_ = d_266_v200_
                lhs32_ = default__.safeIndex(802, (d_266_v200_).length(0))
                lhs33_ = globalState
                lhs34_ = globalState
                lhs31_[lhs32_] = rhs39_
                lhs33_.f5 = rhs40_
                lhs34_.f13 = rhs41_
            d_277_v208_: D2
            d_277_v208_ = D2_DC5((d_189_v143_ if d_57_v48_.f21 else d_189_v143_))
            rhs42_ = d_189_v143_
            rhs43_ = D2_DC5(d_189_v143_)
            rhs44_ = d_23_v15_
            lhs35_ = globalState
            d_189_v143_ = rhs42_
            d_277_v208_ = rhs43_
            lhs35_.f13 = rhs44_
            d_278_v209_: _dafny.Array
            def lambda15_(d_279_v16_):
                def lambda16_(d_280_i16_):
                    return (d_280_i16_) + (len(d_279_v16_))

                return lambda16_

            init8_ = lambda15_(d_24_v16_)
            nw54_ = _dafny.Array(None, 10)
            for i0_8_ in range(nw54_.length(0)):
                nw54_[i0_8_] = init8_(i0_8_)
            d_278_v209_ = nw54_
            index38_ = default__.safeIndex(760, (d_278_v209_).length(0))
            (d_278_v209_)[index38_] = p0
            d_281_v210_: _dafny.Array
            def lambda17_(d_282_v146_):
                def lambda18_(d_283_i17_):
                    return d_282_v146_

                return lambda18_

            init9_ = lambda17_(d_204_v146_)
            nw55_ = _dafny.Array(None, 18)
            for i0_9_ in range(nw55_.length(0)):
                nw55_[i0_9_] = init9_(i0_9_)
            d_281_v210_ = nw55_
            index39_ = default__.safeIndex(722, (d_281_v210_).length(0))
            (d_281_v210_)[index39_] = d_204_v146_
            d_284_v211_: _dafny.Array
            def lambda19_(d_285_v48_):
                def lambda20_(d_286_i18_):
                    return d_285_v48_.f21

                return lambda20_

            init10_ = lambda19_(d_57_v48_)
            nw56_ = _dafny.Array(None, 26)
            for i0_10_ in range(nw56_.length(0)):
                nw56_[i0_10_] = init10_(i0_10_)
            d_284_v211_ = nw56_
            index40_ = default__.safeIndex(784, (d_284_v211_).length(0))
            (d_284_v211_)[index40_] = (d_68_v58_) < (d_68_v58_)
            d_287_v212_: _dafny.Set
            d_287_v212_ = _dafny.Set({d_57_v48_.f21})
            d_288_v213_: T0
            nw57_ = C0()
            nw57_.ctor__(d_66_v56_.f22, len(d_287_v212_), not(d_23_v15_))
            d_288_v213_ = nw57_
            d_289_v214_: _dafny.Set
            d_289_v214_ = _dafny.Set({d_288_v213_, d_288_v213_})
            d_290_v215_: str
            d_290_v215_ = _dafny.CodePoint('j')
            index41_ = default__.safeIndex(760, (d_278_v209_).length(0))
            index42_ = default__.safeIndex(722, (d_281_v210_).length(0))
            index43_ = default__.safeIndex(784, (d_284_v211_).length(0))
            rhs45_ = p1
            rhs46_ = len(_dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('v') if False else _dafny.CodePoint('e')) for d_291_i19_ in range(default__.abs(468))]))
            rhs47_ = (d_204_v146_).set(default__.safeIndex(p1, len(d_204_v146_)), d_290_v215_)
            rhs48_ = d_57_v48_.f21
            rhs49_ = (_dafny.Set({d_288_v213_, d_288_v213_, d_288_v213_}) if (d_290_v215_) in (d_204_v146_) else _dafny.Set({d_288_v213_}))
            lhs36_ = d_278_v209_
            lhs37_ = default__.safeIndex(760, (d_278_v209_).length(0))
            lhs38_ = globalState
            lhs39_ = d_281_v210_
            lhs40_ = default__.safeIndex(722, (d_281_v210_).length(0))
            lhs41_ = d_284_v211_
            lhs42_ = default__.safeIndex(784, (d_284_v211_).length(0))
            lhs36_[lhs37_] = rhs45_
            lhs38_.f5 = rhs46_
            lhs39_[lhs40_] = rhs47_
            lhs41_[lhs42_] = rhs48_
            d_289_v214_ = rhs49_
            d_292_v216_: _dafny.Map
            d_292_v216_ = _dafny.Map({d_57_v48_.f21: d_57_v48_.f21})
            d_293_v217_: D1
            d_293_v217_ = D1_DC4(not((d_288_v213_).f18), ((d_57_v48_).fm5(d_23_v15_, (d_66_v56_).f23, globalState) if (d_24_v16_)[default__.safeIndex((d_66_v56_).f23, len(d_24_v16_))] else d_57_v48_.f21), (p0) > (len(d_292_v216_)))
            index44_ = default__.safeIndex(760, (d_278_v209_).length(0))
            rhs50_ = (d_278_v209_)[default__.safeIndex(760, (d_278_v209_).length(0))]
            rhs51_ = len(_dafny.SeqWithoutIsStrInference([len((d_281_v210_)[default__.safeIndex(722, (d_281_v210_).length(0))])]))
            rhs52_ = d_293_v217_
            lhs43_ = globalState
            lhs44_ = d_278_v209_
            lhs45_ = default__.safeIndex(760, (d_278_v209_).length(0))
            lhs43_.f5 = rhs50_
            lhs44_[lhs45_] = rhs51_
            d_293_v217_ = rhs52_
        d_294_v218_: T0
        nw58_ = C0()
        nw58_.ctor__(d_65_v55_, p0, d_57_v48_.f21)
        d_294_v218_ = nw58_
        d_295_v219_: _dafny.Map
        d_295_v219_ = _dafny.Map({d_66_v56_: d_294_v218_})
        r0 = len(((d_295_v219_).set(d_66_v56_, d_294_v218_)) | (d_295_v219_))
        d_296_v220_: _dafny.Array
        def lambda21_(d_297_v56_):
            def lambda22_(d_298_i20_):
                return (d_298_i20_) - ((d_297_v56_).f23)

            return lambda22_

        init11_ = lambda21_(d_66_v56_)
        nw59_ = _dafny.Array(None, 23)
        for i0_11_ in range(nw59_.length(0)):
            nw59_[i0_11_] = init11_(i0_11_)
        d_296_v220_ = nw59_
        d_299_v221_: _dafny.Map
        d_299_v221_ = _dafny.Map({690: d_296_v220_})
        d_300_v222_: _dafny.MultiSet
        d_300_v222_ = _dafny.MultiSet([d_57_v48_.f21])
        r1 = (d_299_v221_).set((d_300_v222_).cardinality, d_296_v220_)
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_301_v0_: _dafny.Array
        nw60_ = _dafny.Array(_dafny.Seq({}), 19)
        d_301_v0_ = nw60_
        d_302_v1_: _dafny.Array
        nw61_ = _dafny.Array(False, 2)
        d_302_v1_ = nw61_
        d_303_v2_: _dafny.Seq
        d_303_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqrmmy"))
        d_304_v3_: bool
        d_304_v3_ = False
        d_305_globalState_: GlobalState
        nw62_ = GlobalState()
        nw62_.ctor__(d_301_v0_, True, d_302_v1_, 0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okaofhw")), -352, 893, d_303_v2_, 868, 685, True, 229, False, False, 157, False, True, (d_302_v1_ if d_304_v3_ else d_302_v1_))
        d_305_globalState_ = nw62_
        d_306_v4_: _dafny.Array
        def lambda23_(d_307_v2_):
            def lambda24_(d_308_i0_):
                return default__.safeDivisionInt(d_308_i0_, len(d_307_v2_))

            return lambda24_

        init12_ = lambda23_(d_303_v2_)
        nw63_ = _dafny.Array(None, 14)
        for i0_12_ in range(nw63_.length(0)):
            nw63_[i0_12_] = init12_(i0_12_)
        d_306_v4_ = nw63_
        d_309_v5_: _dafny.Map
        d_309_v5_ = _dafny.Map({d_304_v3_: d_306_v4_})
        d_310_v6_: _dafny.Seq
        d_310_v6_ = _dafny.SeqWithoutIsStrInference([((d_309_v5_)[d_304_v3_] if (d_304_v3_) in (d_309_v5_) else d_306_v4_), d_306_v4_, d_306_v4_])
        d_310_v6_ = (d_310_v6_) + (d_310_v6_)
        d_311_v7_: int
        d_311_v7_ = 604
        (d_305_globalState_).f9 = ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([d_304_v3_]))))) * (default__.fm0(len(d_303_v2_), len(d_303_v2_), d_311_v7_, d_311_v7_, d_305_globalState_))
        d_312_v8_: _dafny.MultiSet
        d_312_v8_ = _dafny.MultiSet([d_306_v4_, d_306_v4_])
        if not(((_dafny.MultiSet([d_306_v4_])).set(d_306_v4_, default__.abs(d_311_v7_))) == ((d_312_v8_) | (d_312_v8_))):
            d_313_v9_: _dafny.Array
            def lambda25_(d_314_v7_):
                def lambda26_(d_315_i1_):
                    return _dafny.Set({d_314_v7_})

                return lambda26_

            init13_ = lambda25_(d_311_v7_)
            nw64_ = _dafny.Array(None, 24)
            for i0_13_ in range(nw64_.length(0)):
                nw64_[i0_13_] = init13_(i0_13_)
            d_313_v9_ = nw64_
            d_316_v10_: str
            d_316_v10_ = _dafny.CodePoint('f')
            d_317_v11_: _dafny.Seq
            d_317_v11_ = _dafny.SeqWithoutIsStrInference([d_304_v3_, d_304_v3_, d_304_v3_, (d_311_v7_) != (d_311_v7_), (len(d_303_v2_)) > (d_311_v7_)])
            rhs53_ = (((0) - (d_311_v7_)) * ((0) - (d_311_v7_))) - (d_311_v7_)
            rhs54_ = (d_313_v9_ if (d_311_v7_) <= (d_311_v7_) else d_313_v9_)
            rhs55_ = d_316_v10_
            rhs56_ = d_301_v0_
            rhs57_ = (_dafny.SeqWithoutIsStrInference([d_304_v3_, d_304_v3_, default__.fm1(d_311_v7_, d_304_v3_, _dafny.MultiSet([d_304_v3_]), d_305_globalState_)]) if d_304_v3_ else d_317_v11_)
            lhs46_ = d_305_globalState_
            lhs46_.f9 = rhs53_
            d_313_v9_ = rhs54_
            d_316_v10_ = rhs55_
            d_301_v0_ = rhs56_
            d_317_v11_ = rhs57_
            d_318_v12_: _dafny.MultiSet
            d_318_v12_ = _dafny.MultiSet([d_304_v3_])
            d_319_v13_: _dafny.MultiSet
            d_319_v13_ = _dafny.MultiSet([d_302_v1_])
            d_320_v14_: int
            d_321_v15_: _dafny.Map
            out16_: int
            out17_: _dafny.Map
            out16_, out17_ = default__.m0(((d_318_v12_)[d_304_v3_] if (d_304_v3_) in (d_318_v12_) else (d_319_v13_).cardinality), d_311_v7_, d_305_globalState_)
            d_320_v14_ = out16_
            d_321_v15_ = out17_
            d_322_v16_: _dafny.Set
            d_322_v16_ = _dafny.Set({d_304_v3_, d_304_v3_})
            d_323_v17_: _dafny.Seq
            d_323_v17_ = _dafny.SeqWithoutIsStrInference([615])
            d_324_v18_: _dafny.Map
            d_324_v18_ = _dafny.Map({(0) - (d_311_v7_): (d_317_v11_)[default__.safeIndex(d_320_v14_, len(d_317_v11_))]})
            d_325_v19_: _dafny.Map
            d_325_v19_ = _dafny.Map({default__.fm0(len(d_322_v16_), (d_323_v17_)[default__.safeIndex(d_311_v7_, len(d_323_v17_))], d_311_v7_, len(d_324_v18_), d_305_globalState_): d_304_v3_})
            d_326_v20_: _dafny.Seq
            d_326_v20_ = _dafny.SeqWithoutIsStrInference([len(d_325_v19_)])
            d_327_v21_: _dafny.Seq
            d_327_v21_ = _dafny.SeqWithoutIsStrInference([d_311_v7_, default__.safeModuloInt(len(d_323_v17_), d_320_v14_), d_311_v7_])
            d_328_v22_: _dafny.Map
            d_328_v22_ = _dafny.Map({d_304_v3_: d_304_v3_})
            d_329_v23_: _dafny.MultiSet
            d_329_v23_ = _dafny.MultiSet([d_311_v7_])
            d_330_v24_: _dafny.Seq
            d_330_v24_ = _dafny.SeqWithoutIsStrInference([d_329_v23_])
            d_331_v25_: _dafny.Seq
            d_331_v25_ = _dafny.SeqWithoutIsStrInference([d_330_v24_, d_330_v24_])
            d_332_v26_: _dafny.Map
            d_332_v26_ = _dafny.Map({d_311_v7_: d_311_v7_})
            rhs58_ = default__.fm2(d_316_v10_, d_304_v3_, d_305_globalState_)
            rhs59_ = (d_311_v7_ if (568) != (len(d_328_v22_)) else len(((d_331_v25_)[default__.safeIndex(((d_332_v26_)[515] if (515) in (d_332_v26_) else d_320_v14_), len(d_331_v25_))]) + (d_330_v24_)))
            d_326_v20_ = rhs58_
            d_320_v14_ = rhs59_
            (d_305_globalState_).f13 = d_304_v3_
            (d_305_globalState_).f3 = d_320_v14_
        elif True:
            if (False) == (d_304_v3_):
                (d_305_globalState_).f2 = d_302_v1_
                d_333_v27_: str
                d_333_v27_ = _dafny.CodePoint('r')
                d_334_v28_: _dafny.Map
                d_334_v28_ = _dafny.Map({d_304_v3_: d_333_v27_})
                d_335_v29_: _dafny.Seq
                d_335_v29_ = _dafny.SeqWithoutIsStrInference([d_334_v28_])
                (d_305_globalState_).f5 = len((d_335_v29_) + (d_335_v29_))
                d_336_v30_: _dafny.Map
                d_336_v30_ = _dafny.Map({(D0_DC0(d_304_v3_)).cf0: d_311_v7_})
                d_337_v31_: _dafny.Seq
                d_337_v31_ = _dafny.SeqWithoutIsStrInference([d_304_v3_, d_304_v3_])
                rhs60_ = d_336_v30_
                rhs61_ = ((d_337_v31_ if False else d_337_v31_)) + (d_337_v31_)
                rhs62_ = d_304_v3_
                lhs47_ = d_305_globalState_
                d_336_v30_ = rhs60_
                d_337_v31_ = rhs61_
                lhs47_.f13 = rhs62_
                (d_305_globalState_).f3 = 176
                d_338_v32_: _dafny.Array
                nw65_ = _dafny.Array(_dafny.Set({}), 12)
                d_338_v32_ = nw65_
                d_339_v33_: _dafny.Set
                d_339_v33_ = _dafny.Set({d_311_v7_, d_311_v7_, d_311_v7_, -714, (0) - ((0) - (d_311_v7_))})
                index45_ = default__.safeIndex(220, (d_338_v32_).length(0))
                (d_338_v32_)[index45_] = d_339_v33_
                index46_ = default__.safeIndex(220, (d_338_v32_).length(0))
                (d_338_v32_)[index46_] = d_339_v33_
            elif True:
                d_340_v34_: _dafny.Seq
                d_340_v34_ = _dafny.SeqWithoutIsStrInference([d_304_v3_])
                d_341_v35_: _dafny.Map
                d_341_v35_ = _dafny.Map({d_304_v3_: d_304_v3_})
                d_342_v36_: _dafny.Seq
                d_342_v36_ = _dafny.SeqWithoutIsStrInference([(d_340_v34_) + (d_340_v34_), ((D1_DC2(_dafny.SeqWithoutIsStrInference([d_304_v3_, d_304_v3_]))).cf4) + (_dafny.SeqWithoutIsStrInference([((d_341_v35_)[d_304_v3_] if (d_304_v3_) in (d_341_v35_) else d_304_v3_)]))])
                d_340_v34_ = (d_342_v36_)[default__.safeIndex(d_311_v7_, len(d_342_v36_))]
                d_343_v37_: _dafny.Map
                d_343_v37_ = _dafny.Map({(d_311_v7_) + (d_311_v7_): d_304_v3_})
                d_343_v37_ = (d_343_v37_).set(d_311_v7_, (d_303_v2_) != (d_303_v2_))
                (d_305_globalState_).f9 = (0) - (((0) - ((d_311_v7_) * (d_311_v7_)) if d_304_v3_ else d_311_v7_))
                index47_ = default__.safeIndex(886, (d_302_v1_).length(0))
                (d_302_v1_)[index47_] = False
                index48_ = default__.safeIndex(886, (d_302_v1_).length(0))
                (d_302_v1_)[index48_] = (d_304_v3_ if d_304_v3_ else False)
                d_344_v38_: D1
                d_344_v38_ = D1_DC4(d_304_v3_, (d_302_v1_)[default__.safeIndex(886, (d_302_v1_).length(0))], d_304_v3_)
                d_345_v39_: _dafny.MultiSet
                d_345_v39_ = _dafny.MultiSet([d_344_v38_, d_344_v38_])
                index49_ = default__.safeIndex(220, (d_306_v4_).length(0))
                (d_306_v4_)[index49_] = ((d_345_v39_)[d_344_v38_] if (d_344_v38_) in (d_345_v39_) else d_311_v7_)
                d_346_v40_: _dafny.Array
                def lambda27_(d_347_v1_, d_348_v7_, d_349_v2_):
                    def lambda28_(d_350_i2_):
                        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghaimy")) if (d_347_v1_)[default__.safeIndex(886, (d_347_v1_).length(0))] else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvj"))).set(default__.safeIndex((0) - (d_348_v7_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvj")))), _dafny.CodePoint('k')))).set(default__.safeIndex(len(d_349_v2_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghaimy")) if (d_347_v1_)[default__.safeIndex(886, (d_347_v1_).length(0))] else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvj"))).set(default__.safeIndex((0) - (d_348_v7_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvj")))), _dafny.CodePoint('k'))))), _dafny.CodePoint('l'))

                    return lambda28_

                init14_ = lambda27_(d_302_v1_, d_311_v7_, d_303_v2_)
                nw66_ = _dafny.Array(None, 22)
                for i0_14_ in range(nw66_.length(0)):
                    nw66_[i0_14_] = init14_(i0_14_)
                d_346_v40_ = nw66_
                index50_ = default__.safeIndex(35, (d_346_v40_).length(0))
                (d_346_v40_)[index50_] = d_303_v2_
                d_351_v41_: _dafny.MultiSet
                d_351_v41_ = _dafny.MultiSet([(d_302_v1_)[default__.safeIndex(886, (d_302_v1_).length(0))]])
                d_352_v42_: _dafny.Map
                d_352_v42_ = _dafny.Map({False: d_303_v2_})
                d_353_v43_: _dafny.Seq
                d_353_v43_ = _dafny.SeqWithoutIsStrInference([(d_303_v2_) + (d_303_v2_), ((d_352_v42_)[d_304_v3_] if (d_304_v3_) in (d_352_v42_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_354_i3_ in range(default__.abs(502))])), d_303_v2_])
                index51_ = default__.safeIndex(220, (d_306_v4_).length(0))
                index52_ = default__.safeIndex(35, (d_346_v40_).length(0))
                rhs63_ = default__.fm1(d_311_v7_, (d_303_v2_) < (d_303_v2_), (_dafny.MultiSet([d_304_v3_, (d_302_v1_)[default__.safeIndex(886, (d_302_v1_).length(0))]])) - (d_351_v41_), d_305_globalState_)
                rhs64_ = default__.safeDivisionInt(d_311_v7_, default__.safeModuloInt((0) - (len(d_303_v2_)), d_311_v7_))
                rhs65_ = d_304_v3_
                rhs66_ = (d_353_v43_)[default__.safeIndex(d_311_v7_, len(d_353_v43_))]
                lhs48_ = d_305_globalState_
                lhs49_ = d_306_v4_
                lhs50_ = default__.safeIndex(220, (d_306_v4_).length(0))
                lhs51_ = d_346_v40_
                lhs52_ = default__.safeIndex(35, (d_346_v40_).length(0))
                lhs48_.f13 = rhs63_
                lhs49_[lhs50_] = rhs64_
                d_304_v3_ = rhs65_
                lhs51_[lhs52_] = rhs66_
            d_355_v44_: _dafny.Seq
            d_355_v44_ = _dafny.SeqWithoutIsStrInference([(0) - (default__.safeDivisionInt(d_311_v7_, d_311_v7_))])
            d_355_v44_ = (d_355_v44_ if not((not(d_304_v3_) if d_304_v3_ else d_304_v3_)) else (d_355_v44_) + (d_355_v44_))
            if (not (d_304_v3_) or (d_304_v3_) if (d_311_v7_) > (810) else (d_311_v7_) > (d_311_v7_)):
                d_356_v45_: _dafny.MultiSet
                d_356_v45_ = _dafny.MultiSet([d_304_v3_])
                d_357_v46_: _dafny.Set
                d_357_v46_ = _dafny.Set({True, default__.fm1(d_311_v7_, d_304_v3_, d_356_v45_, d_305_globalState_)})
                index53_ = default__.safeIndex(261, (d_306_v4_).length(0))
                (d_306_v4_)[index53_] = default__.fm0(len(d_357_v46_), d_311_v7_, d_311_v7_, d_311_v7_, d_305_globalState_)
                d_358_v47_: _dafny.Set
                d_358_v47_ = _dafny.Set({d_303_v2_, d_303_v2_, default__.fm3(d_305_globalState_), d_303_v2_, d_303_v2_})
                index54_ = default__.safeIndex(261, (d_306_v4_).length(0))
                (d_306_v4_)[index54_] = (0) - (default__.fm0(d_311_v7_, 780, d_311_v7_, len((d_358_v47_ if d_304_v3_ else d_358_v47_)), d_305_globalState_))
                (d_305_globalState_).f13 = d_304_v3_
                d_304_v3_ = not(d_304_v3_)
                d_359_v48_: str
                d_359_v48_ = _dafny.CodePoint('b')
                rhs67_ = d_304_v3_
                rhs68_ = d_311_v7_
                rhs69_ = d_359_v48_
                lhs53_ = d_305_globalState_
                lhs54_ = d_305_globalState_
                lhs53_.f13 = rhs67_
                lhs54_.f9 = rhs68_
                d_359_v48_ = rhs69_
                d_360_v49_: int
                d_361_v50_: _dafny.Map
                out18_: int
                out19_: _dafny.Map
                out18_, out19_ = default__.m0(d_311_v7_, d_311_v7_, d_305_globalState_)
                d_360_v49_ = out18_
                d_361_v50_ = out19_
            elif True:
                d_362_v51_: str
                d_362_v51_ = _dafny.CodePoint('a')
                d_362_v51_ = ((d_303_v2_)[default__.safeIndex(d_311_v7_, len(d_303_v2_))] if d_304_v3_ else (d_362_v51_ if d_304_v3_ else d_362_v51_))
                d_363_v52_: _dafny.Map
                d_363_v52_ = _dafny.Map({d_304_v3_: 628})
                d_364_v53_: _dafny.MultiSet
                d_364_v53_ = _dafny.MultiSet([d_363_v52_, d_363_v52_, (d_363_v52_).set(d_304_v3_, d_311_v7_)])
                (d_305_globalState_).f13 = (d_364_v53_).isdisjoint((d_364_v53_) | (d_364_v53_))
                d_365_v54_: _dafny.MultiSet
                d_365_v54_ = _dafny.MultiSet([d_302_v1_, d_302_v1_])
                d_366_v55_: _dafny.MultiSet
                d_366_v55_ = _dafny.MultiSet([d_304_v3_, d_304_v3_])
                (d_305_globalState_).f5 = (((d_365_v54_) | ((d_365_v54_).set(d_302_v1_, default__.abs(d_311_v7_)))).cardinality) + ((len(d_363_v52_)) - ((d_366_v55_).cardinality))
                d_367_v56_: _dafny.Seq
                d_367_v56_ = _dafny.SeqWithoutIsStrInference([not(d_304_v3_)])
                d_368_v57_: _dafny.Map
                d_368_v57_ = _dafny.Map({default__.fm0(d_311_v7_, d_311_v7_, d_311_v7_, len(d_367_v56_), d_305_globalState_): d_311_v7_})
                d_369_v59_: _dafny.MultiSet
                d_369_v59_ = _dafny.MultiSet([d_311_v7_, d_311_v7_])
                d_370_v60_: _dafny.Seq
                d_370_v60_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({248: d_311_v7_})])
                d_371_v61_: D2
                d_371_v61_ = D2_DC5(_dafny.Map({d_311_v7_: ((d_363_v52_)[d_304_v3_] if (d_304_v3_) in (d_363_v52_) else -956)}))
                d_372_v64_: _dafny.Array
                nw67_ = _dafny.Array(None, 27)
                nw67_[int(0)] = d_368_v57_
                nw67_[int(1)] = d_368_v57_
                nw67_[int(2)] = (d_368_v57_) | (d_368_v57_)
                def iife18_():
                    coll6_ = _dafny.Map()
                    compr_6_: int
                    for compr_6_ in (d_369_v59_).Elements:
                        d_373_v58_: int = compr_6_
                        if (d_373_v58_) in (d_369_v59_):
                            coll6_[default__.safeDivisionInt(d_373_v58_, d_311_v7_)] = (0) - (d_311_v7_)
                    return _dafny.Map(coll6_)
                nw67_[int(3)] = iife18_()
                
                nw67_[int(4)] = d_368_v57_
                nw67_[int(5)] = d_368_v57_
                nw67_[int(6)] = d_368_v57_
                nw67_[int(7)] = d_368_v57_
                nw67_[int(8)] = (d_368_v57_) | (d_368_v57_)
                nw67_[int(9)] = (d_370_v60_)[default__.safeIndex(498, len(d_370_v60_))]
                nw67_[int(10)] = d_368_v57_
                nw67_[int(11)] = (d_368_v57_) | (d_368_v57_)
                nw67_[int(12)] = d_368_v57_
                nw67_[int(13)] = _dafny.Map({d_311_v7_: d_311_v7_})
                nw67_[int(14)] = (d_368_v57_ if d_304_v3_ else (d_370_v60_)[default__.safeIndex(d_311_v7_, len(d_370_v60_))])
                nw67_[int(15)] = (d_371_v61_).cf9
                nw67_[int(16)] = (d_368_v57_) | (d_368_v57_)
                nw67_[int(17)] = d_368_v57_
                nw67_[int(18)] = _dafny.Map({len(d_303_v2_): default__.fm0(d_311_v7_, len(d_368_v57_), d_311_v7_, len(d_303_v2_), d_305_globalState_)})
                nw67_[int(19)] = d_368_v57_
                nw67_[int(20)] = d_368_v57_
                nw67_[int(21)] = default__.fm4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")), True, d_304_v3_, d_304_v3_, d_305_globalState_)
                nw67_[int(22)] = d_368_v57_
                def iife19_():
                    coll7_ = _dafny.Map()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(-970, 932):
                        d_374_v62_: int = compr_7_
                        if ((-970) <= (d_374_v62_)) and ((d_374_v62_) < (932)):
                            coll7_[(d_374_v62_) - (d_311_v7_)] = len(_dafny.SeqWithoutIsStrInference([723]))
                    return _dafny.Map(coll7_)
                nw67_[int(23)] = iife19_()
                
                def iife20_():
                    coll8_ = _dafny.Map()
                    compr_8_: int
                    for compr_8_ in _dafny.IntegerRange(-721, 577):
                        d_375_v63_: int = compr_8_
                        if ((-721) <= (d_375_v63_)) and ((d_375_v63_) < (577)):
                            coll8_[(d_375_v63_) * (len(_dafny.Set({d_304_v3_})))] = d_311_v7_
                    return _dafny.Map(coll8_)
                nw67_[int(24)] = iife20_()
                
                nw67_[int(25)] = d_368_v57_
                nw67_[int(26)] = (default__.fm4(d_303_v2_, d_304_v3_, not(False), d_304_v3_, d_305_globalState_) if False else (_dafny.Map({d_311_v7_: len(d_303_v2_)})).set(d_311_v7_, len(d_363_v52_)))
                d_372_v64_ = nw67_
                d_372_v64_ = d_372_v64_
                d_376_v65_: C1
                nw68_ = C1()
                nw68_.ctor__(d_304_v3_, (d_311_v7_) * (d_311_v7_), -522, d_304_v3_)
                d_376_v65_ = nw68_
            index55_ = default__.safeIndex(34, (d_306_v4_).length(0))
            (d_306_v4_)[index55_] = d_311_v7_
            index56_ = default__.safeIndex(34, (d_306_v4_).length(0))
            (d_306_v4_)[index56_] = default__.safeDivisionInt(d_311_v7_, 425)
            d_377_v66_: _dafny.Map
            d_377_v66_ = _dafny.Map({d_304_v3_: d_311_v7_})
            d_378_v67_: _dafny.Seq
            d_378_v67_ = _dafny.SeqWithoutIsStrInference([d_304_v3_, d_304_v3_])
            d_379_v68_: int
            d_380_v69_: _dafny.Map
            out20_: int
            out21_: _dafny.Map
            out20_, out21_ = default__.m0(default__.fm0((d_306_v4_)[default__.safeIndex(34, (d_306_v4_).length(0))], (0) - ((d_306_v4_)[default__.safeIndex(34, (d_306_v4_).length(0))]), default__.fm0(len((d_377_v66_)), d_311_v7_, len(_dafny.SeqWithoutIsStrInference([(d_306_v4_)[default__.safeIndex(34, (d_306_v4_).length(0))] for d_381_i4_ in range(default__.abs(263))])), len(d_378_v67_), d_305_globalState_), len(d_303_v2_), d_305_globalState_), 659, d_305_globalState_)
            d_379_v68_ = out20_
            d_380_v69_ = out21_
        d_382_v70_: _dafny.Map
        d_382_v70_ = _dafny.Map({d_311_v7_: d_311_v7_})
        def iife21_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in (d_382_v70_).keys.Elements:
                d_383_v71_: int = compr_9_
                if (d_383_v71_) in (d_382_v70_):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeModuloInt(d_383_v71_, 240)]))
            return _dafny.Set(coll9_)
        (d_305_globalState_).f3 = ((len(iife21_()
        )) * (d_311_v7_)) + (d_311_v7_)
        (d_305_globalState_).f9 = -648
        d_384_v72_: D0
        d_384_v72_ = D0_DC0(True)
        source3_ = d_384_v72_
        if source3_.is_DC1:
            d_385___mcc_h0_ = source3_.cf1
            d_386___mcc_h1_ = source3_.cf2
            d_387___mcc_h2_ = source3_.cf3
            d_388_cf3_ = d_387___mcc_h2_
            d_389_cf2_ = d_386___mcc_h1_
            d_390_cf1_ = d_385___mcc_h0_
            d_391_v73_: _dafny.Map
            d_391_v73_ = _dafny.Map({d_311_v7_: d_304_v3_})
            index57_ = default__.safeIndex(862, (d_306_v4_).length(0))
            (d_306_v4_)[index57_] = default__.safeDivisionInt(len(_dafny.Set({((d_391_v73_)[415] if (415) in (d_391_v73_) else d_388_cf3_)})), d_389_cf2_)
            d_392_v74_: _dafny.MultiSet
            d_392_v74_ = _dafny.MultiSet([d_304_v3_, d_304_v3_, d_304_v3_, d_388_cf3_])
            d_393_v75_: _dafny.Seq
            d_393_v75_ = _dafny.SeqWithoutIsStrInference([d_388_cf3_, d_304_v3_])
            d_394_v76_: _dafny.Seq
            d_394_v76_ = _dafny.SeqWithoutIsStrInference([d_389_cf2_])
            d_395_v77_: _dafny.Seq
            d_395_v77_ = _dafny.SeqWithoutIsStrInference([d_394_v76_])
            index58_ = default__.safeIndex(862, (d_306_v4_).length(0))
            (d_306_v4_)[index58_] = (default__.safeModuloInt(d_390_cf1_, 750) if (_dafny.MultiSet(d_393_v75_)).issubset(d_392_v74_) else (987) - ((0) - (len((d_395_v77_)[default__.safeIndex(d_390_cf1_, len(d_395_v77_))]))))
            d_396_v78_: _dafny.Map
            d_396_v78_ = _dafny.Map({d_304_v3_: d_304_v3_})
            d_397_v79_: int
            d_398_v80_: _dafny.Map
            out22_: int
            out23_: _dafny.Map
            out22_, out23_ = default__.m0((0) - (((d_392_v74_).cardinality) - (len(d_396_v78_))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmuinekm"))), d_305_globalState_)
            d_397_v79_ = out22_
            d_398_v80_ = out23_
            d_399_v81_: _dafny.Map
            d_399_v81_ = _dafny.Map({d_304_v3_: d_311_v7_})
            d_400_v82_: _dafny.Map
            d_400_v82_ = _dafny.Map({d_399_v81_: d_304_v3_})
            d_393_v75_ = (d_393_v75_).set(default__.safeIndex((d_306_v4_)[default__.safeIndex(862, (d_306_v4_).length(0))], len(d_393_v75_)), (d_397_v79_) == (len(d_400_v82_)))
            d_396_v78_ = (d_396_v78_).set(False, False)
        elif True:
            d_401___mcc_h3_ = source3_.cf0
            d_402_cf0_ = d_401___mcc_h3_
            (d_305_globalState_).f4 = d_303_v2_
            (d_305_globalState_).f4 = (d_303_v2_) + (d_303_v2_)
            d_403_v83_: T1
            nw69_ = C1()
            nw69_.ctor__(not(d_402_cf0_), d_311_v7_, d_311_v7_, d_402_cf0_)
            d_403_v83_ = nw69_
            d_404_v84_: _dafny.Map
            d_404_v84_ = _dafny.Map({d_403_v83_: d_311_v7_})
            d_405_v85_: _dafny.Seq
            d_405_v85_ = _dafny.SeqWithoutIsStrInference([(d_311_v7_) - (d_311_v7_), d_311_v7_, 516, len(d_404_v84_), d_403_v83_.f20])
            d_405_v85_ = d_405_v85_
            d_406_v86_: D1
            d_406_v86_ = D1_DC3((d_403_v83_).f19)
            d_407_v87_: _dafny.Map
            d_407_v87_ = _dafny.Map({False: d_406_v86_})
            d_408_v88_: _dafny.Seq
            d_408_v88_ = _dafny.SeqWithoutIsStrInference([d_407_v87_, d_407_v87_, d_407_v87_])
            (d_305_globalState_).f13 = (d_403_v83_).fm5(False, len(d_408_v88_), d_305_globalState_)
        d_409_v89_: C1
        nw70_ = C1()
        nw70_.ctor__(False, len(default__.fm8(d_305_globalState_)), d_311_v7_, not(d_304_v3_))
        d_409_v89_ = nw70_
        d_410_v90_: _dafny.Map
        d_410_v90_ = _dafny.Map({d_409_v89_: d_303_v2_})
        d_411_v91_: D4
        d_411_v91_ = D4_DC9(d_409_v89_)
        d_410_v90_ = (d_410_v90_).set((d_411_v91_).cf13, d_303_v2_)
        d_412_v92_: str
        d_412_v92_ = _dafny.CodePoint('t')
        d_413_v93_: T1
        nw71_ = C1()
        nw71_.ctor__(d_304_v3_, default__.safeModuloInt(d_311_v7_, d_311_v7_), d_311_v7_, d_409_v89_.f21)
        d_413_v93_ = nw71_
        d_414_v94_: _dafny.Seq
        d_414_v94_ = _dafny.SeqWithoutIsStrInference([not((d_413_v93_).f18), d_409_v89_.f21])
        d_415_v95_: _dafny.Map
        d_415_v95_ = _dafny.Map({d_414_v94_: d_304_v3_})
        d_416_v96_: _dafny.Map
        d_416_v96_ = _dafny.Map({d_301_v0_: (d_413_v93_).f18})
        d_417_v97_: _dafny.MultiSet
        d_417_v97_ = _dafny.MultiSet([((d_415_v95_)[d_414_v94_] if (d_414_v94_) in (d_415_v95_) else d_409_v89_.f21), d_304_v3_, ((d_416_v96_)[d_301_v0_] if (d_301_v0_) in (d_416_v96_) else False)])
        d_418_v98_: _dafny.Map
        d_418_v98_ = _dafny.Map({d_304_v3_: d_304_v3_})
        d_419_v99_: _dafny.Seq
        d_419_v99_ = _dafny.SeqWithoutIsStrInference([d_418_v98_, _dafny.Map({(d_413_v93_).f18: d_304_v3_})])
        d_420_v100_: _dafny.Seq
        d_420_v100_ = _dafny.SeqWithoutIsStrInference([-186])
        d_421_v101_: _dafny.Map
        d_421_v101_ = _dafny.Map({(d_413_v93_).f18: (d_413_v93_).f19})
        d_422_v102_: D0
        d_422_v102_ = D0_DC1(d_311_v7_, (d_420_v100_)[default__.safeIndex(len(d_421_v101_), len(d_420_v100_))], d_304_v3_)
        d_423_v103_: _dafny.Map
        d_423_v103_ = _dafny.Map({d_311_v7_: d_413_v93_})
        rhs70_ = _dafny.CodePoint('s')
        rhs71_ = (d_303_v2_ if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nd")))
        rhs72_ = (d_413_v93_).fm5((d_419_v99_) == ((d_419_v99_).set(default__.safeIndex(len(_dafny.Map({d_412_v92_: len(d_382_v70_)})), len(d_419_v99_)), default__.fm9(d_409_v89_.f21, (d_413_v93_).f18, d_305_globalState_))), (d_422_v102_).cf2, d_305_globalState_)
        rhs73_ = ((d_423_v103_)[default__.safeDivisionInt(d_413_v93_.f20, d_311_v7_)] if (default__.safeDivisionInt(d_413_v93_.f20, d_311_v7_)) in (d_423_v103_) else d_413_v93_)
        rhs74_ = d_417_v97_
        d_412_v92_ = rhs70_
        d_303_v2_ = rhs71_
        d_304_v3_ = rhs72_
        d_413_v93_ = rhs73_
        d_417_v97_ = rhs74_
        d_424_v104_: _dafny.MultiSet
        d_424_v104_ = _dafny.MultiSet([d_409_v89_])
        (d_409_v89_).f21 = (d_304_v3_) and (default__.fm1(((d_424_v104_)[d_409_v89_] if (d_409_v89_) in (d_424_v104_) else (0) - (d_413_v93_.f20)), d_304_v3_, d_417_v97_, d_305_globalState_))
        d_425_v105_: D2
        d_425_v105_ = D2_DC5(d_382_v70_)
        d_425_v105_ = default__.fm10(d_311_v7_, d_305_globalState_)
        if (default__.fm2(d_412_v92_, d_304_v3_, d_305_globalState_)) != ((d_420_v100_) + (d_420_v100_)):
            d_426_v106_: C1
            nw72_ = C1()
            nw72_.ctor__((d_413_v93_).f18, (d_413_v93_).f19, d_311_v7_, d_409_v89_.f21)
            d_426_v106_ = nw72_
            (d_305_globalState_).f13 = (d_413_v93_).f18
            if (d_413_v93_).fm5(d_304_v3_, d_311_v7_, d_305_globalState_):
                (d_305_globalState_).f3 = (0) - ((d_413_v93_).f19)
                (d_409_v89_).f21 = False
                (d_409_v89_).f21 = d_426_v106_.f21
                d_427_v107_: _dafny.MultiSet
                d_427_v107_ = _dafny.MultiSet([d_413_v93_.f20])
                d_428_v108_: D2
                d_428_v108_ = D2_DC6(d_427_v107_)
                d_429_v109_: _dafny.Set
                d_429_v109_ = _dafny.Set({(d_413_v93_).f19, (d_413_v93_).f19, d_311_v7_})
                d_430_v110_: T0
                nw73_ = C0()
                nw73_.ctor__(D2_DC6(_dafny.MultiSet([len(d_429_v109_), (d_413_v93_).f19, d_413_v93_.f20])), (d_413_v93_).f19, (d_413_v93_).f18)
                d_430_v110_ = nw73_
                d_431_v111_: _dafny.Seq
                d_431_v111_ = _dafny.SeqWithoutIsStrInference([d_430_v110_, d_430_v110_, d_430_v110_, d_430_v110_, d_430_v110_])
                d_432_v112_: C0
                nw74_ = C0()
                nw74_.ctor__(d_428_v108_, d_413_v93_.f20, (d_409_v89_).fm5(d_409_v89_.f21, len(d_431_v111_), d_305_globalState_))
                d_432_v112_ = nw74_
                d_433_v113_: _dafny.Seq
                d_433_v113_ = _dafny.SeqWithoutIsStrInference([d_432_v112_])
                d_434_v114_: C1
                nw75_ = C1()
                nw75_.ctor__(d_409_v89_.f21, len((d_433_v113_) + (d_433_v113_)), (d_311_v7_) * ((d_413_v93_).f19), ((d_432_v112_).f23) > ((default__.fm11(d_305_globalState_)).cf5))
                d_434_v114_ = nw75_
                index59_ = default__.safeIndex(505, (d_306_v4_).length(0))
                (d_306_v4_)[index59_] = d_413_v93_.f20
                index60_ = default__.safeIndex(505, (d_306_v4_).length(0))
                rhs75_ = (d_413_v93_.f20) * (d_413_v93_.f20)
                rhs76_ = default__.safeDivisionInt(default__.safeModuloInt(-158, 654), d_413_v93_.f20)
                rhs77_ = (_dafny.Set({(d_413_v93_).f19, (d_417_v97_).cardinality})).isdisjoint((d_429_v109_).intersection(d_429_v109_))
                rhs78_ = d_301_v0_
                rhs79_ = False
                lhs55_ = d_306_v4_
                lhs56_ = default__.safeIndex(505, (d_306_v4_).length(0))
                lhs57_ = d_305_globalState_
                lhs58_ = d_305_globalState_
                lhs59_ = d_305_globalState_
                lhs55_[lhs56_] = rhs75_
                lhs57_.f9 = rhs76_
                lhs58_.f13 = rhs77_
                d_301_v0_ = rhs78_
                lhs59_.f13 = rhs79_
            elif True:
                d_435_v115_: _dafny.Array
                def lambda29_(d_436_v101_):
                    def lambda30_(d_437_i5_):
                        return d_436_v101_

                    return lambda30_

                init15_ = lambda29_(d_421_v101_)
                nw76_ = _dafny.Array(None, 21)
                for i0_15_ in range(nw76_.length(0)):
                    nw76_[i0_15_] = init15_(i0_15_)
                d_435_v115_ = nw76_
                d_438_v116_: _dafny.Map
                d_438_v116_ = _dafny.Map({d_414_v94_: d_435_v115_})
                d_438_v116_ = (d_438_v116_).set(_dafny.SeqWithoutIsStrInference([(d_413_v93_).f18, (d_413_v93_).f18, (d_413_v93_).f18, d_409_v89_.f21, not(d_426_v106_.f21)]), d_435_v115_)
                d_439_v117_: D1
                d_439_v117_ = D1_DC2(_dafny.SeqWithoutIsStrInference([(d_413_v93_).f18, d_409_v89_.f21, d_426_v106_.f21]))
                d_440_v118_: D2
                d_440_v118_ = D2_DC6(_dafny.MultiSet([115, (d_413_v93_).f19]))
                d_441_v119_: _dafny.MultiSet
                d_441_v119_ = _dafny.MultiSet([(d_413_v93_).f19])
                pat_let_tv11_ = d_441_v119_
                d_442_v120_: C0
                nw77_ = C0()
                def iife22_(_pat_let6_0):
                    def iife23_(d_443_dt__update__tmp_h0_):
                        def iife24_(_pat_let7_0):
                            def iife25_(d_444_dt__update_hcf10_h0_):
                                return D2_DC6(d_444_dt__update_hcf10_h0_)
                            return iife25_(_pat_let7_0)
                        return iife24_(pat_let_tv11_)
                    return iife23_(_pat_let6_0)
                nw77_.ctor__(iife22_(d_440_v118_), (d_413_v93_).f19, d_426_v106_.f21)
                d_442_v120_ = nw77_
                d_445_v121_: _dafny.Map
                d_445_v121_ = _dafny.Map({d_442_v120_: d_413_v93_.f20})
                d_446_v122_: _dafny.Map
                d_446_v122_ = _dafny.Map({d_439_v117_: d_445_v121_})
                d_446_v122_ = d_446_v122_
                d_442_v120_ = d_442_v120_
                d_447_v123_: _dafny.Map
                d_447_v123_ = _dafny.Map({d_426_v106_: d_413_v93_.f20})
                d_448_v124_: _dafny.Array
                d_449_v125_: _dafny.Set
                out24_: _dafny.Array
                out25_: _dafny.Set
                out24_, out25_ = (d_413_v93_).m2((d_441_v119_) - (_dafny.MultiSet([d_311_v7_, d_413_v93_.f20])), ((d_447_v123_)[d_409_v89_] if (d_409_v89_) in (d_447_v123_) else (d_413_v93_).f19), d_302_v1_, d_303_v2_, d_305_globalState_)
                d_448_v124_ = out24_
                d_449_v125_ = out25_
                d_450_v126_: _dafny.Map
                d_450_v126_ = _dafny.Map({d_409_v89_.f21: d_302_v1_})
                d_450_v126_ = (d_450_v126_).set(d_409_v89_.f21, d_302_v1_)
            d_451_v127_: _dafny.Set
            d_451_v127_ = _dafny.Set({d_302_v1_, d_302_v1_})
            if (((d_413_v93_).f19) * (len(d_451_v127_))) == (default__.safeDivisionInt((d_413_v93_).f19, (d_413_v93_).f19)):
                d_452_v128_: _dafny.Array
                nw78_ = _dafny.Array(None, 18)
                d_452_v128_ = nw78_
                d_453_v129_: _dafny.Map
                d_453_v129_ = _dafny.Map({d_452_v128_: d_306_v4_})
                d_454_v130_: _dafny.Set
                d_454_v130_ = _dafny.Set({_dafny.CodePoint('v'), _dafny.CodePoint('d')})
                d_455_v131_: _dafny.MultiSet
                d_455_v131_ = _dafny.MultiSet([default__.fm0(d_413_v93_.f20, d_413_v93_.f20, (d_413_v93_).f19, len(d_454_v130_), d_305_globalState_)])
                d_456_v132_: C1
                nw79_ = C1()
                nw79_.ctor__(d_409_v89_.f21, (d_413_v93_).f19, default__.fm0(d_413_v93_.f20, (d_413_v93_).f19, len(d_453_v129_), d_311_v7_, d_305_globalState_), (D0_DC1((d_455_v131_).cardinality, (d_413_v93_).f19, d_426_v106_.f21)).cf3)
                d_456_v132_ = nw79_
                (d_305_globalState_).f9 = d_311_v7_
                index61_ = default__.safeIndex(615, (d_302_v1_).length(0))
                (d_302_v1_)[index61_] = not(d_409_v89_.f21)
                index62_ = default__.safeIndex(615, (d_302_v1_).length(0))
                (d_302_v1_)[index62_] = ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfbbox")))) + (d_413_v93_.f20)) < ((d_413_v93_).f19)
                d_382_v70_ = (d_382_v70_).set(d_413_v93_.f20, (0) - (len(d_418_v98_)))
                d_303_v2_ = ((d_410_v90_)[d_426_v106_] if (d_426_v106_) in (d_410_v90_) else d_303_v2_)
            elif True:
                (d_305_globalState_).f5 = (0) - (default__.safeModuloInt(d_413_v93_.f20, ((d_417_v97_) - (d_417_v97_)).cardinality))
                (d_409_v89_).f21 = ((_dafny.MultiSet([(d_413_v93_).f19, d_413_v93_.f20])).cardinality) != ((0) - (d_311_v7_))
                d_304_v3_ = (d_413_v93_.f20) in (d_382_v70_)
                d_457_v133_: int
                d_458_v134_: bool
                d_459_v135_: int
                d_460_v136_: int
                out26_: int
                out27_: bool
                out28_: int
                out29_: int
                out26_, out27_, out28_, out29_ = (d_409_v89_).m1(d_426_v106_.f21, (d_426_v106_.f21) or (False), d_305_globalState_)
                d_457_v133_ = out26_
                d_458_v134_ = out27_
                d_459_v135_ = out28_
                d_460_v136_ = out29_
                rhs80_ = (d_413_v93_).f19
                rhs81_ = (0) - (d_413_v93_.f20)
                lhs60_ = d_305_globalState_
                d_459_v135_ = rhs80_
                lhs60_.f3 = rhs81_
            (d_413_v93_).f20 = -525
        elif True:
            d_461_v137_: _dafny.Seq
            d_461_v137_ = _dafny.SeqWithoutIsStrInference([(d_414_v94_ if (d_413_v93_).f18 else _dafny.SeqWithoutIsStrInference([(d_413_v93_).f18])), _dafny.SeqWithoutIsStrInference([d_409_v89_.f21])])
            d_462_v138_: _dafny.Array
            nw80_ = _dafny.Array(_dafny.Seq({}), 21)
            d_462_v138_ = nw80_
            index63_ = default__.safeIndex(670, (d_462_v138_).length(0))
            (d_462_v138_)[index63_] = (d_414_v94_) + (d_414_v94_)
            index64_ = default__.safeIndex(670, (d_462_v138_).length(0))
            rhs82_ = (((d_413_v93_).f19) + ((d_413_v93_).f19)) + (d_311_v7_)
            rhs83_ = d_409_v89_.f21
            rhs84_ = d_461_v137_
            rhs85_ = d_414_v94_
            rhs86_ = ((d_413_v93_.f20) >= ((d_413_v93_).f19) if (_dafny.SeqWithoutIsStrInference([d_412_v92_ for d_463_i6_ in range(default__.abs(79))])) != (d_303_v2_) else ((d_413_v93_).f18) or ((d_413_v93_).f18))
            lhs61_ = d_305_globalState_
            lhs62_ = d_409_v89_
            lhs63_ = d_462_v138_
            lhs64_ = default__.safeIndex(670, (d_462_v138_).length(0))
            lhs65_ = d_409_v89_
            lhs61_.f9 = rhs82_
            lhs62_.f21 = rhs83_
            d_461_v137_ = rhs84_
            lhs63_[lhs64_] = rhs85_
            lhs65_.f21 = rhs86_
            d_464_v139_: D1
            d_464_v139_ = D1_DC3((d_413_v93_).f19)
            pat_let_tv12_ = d_413_v93_
            def iife26_(_pat_let8_0):
                def iife27_(d_465_dt__update__tmp_h1_):
                    def iife28_(_pat_let9_0):
                        def iife29_(d_466_dt__update_hcf5_h0_):
                            return D1_DC3(d_466_dt__update_hcf5_h0_)
                        return iife29_(_pat_let9_0)
                    return iife28_(pat_let_tv12_.f20)
                return iife27_(_pat_let8_0)
            d_464_v139_ = iife26_(d_464_v139_)
            index65_ = default__.safeIndex(670, (d_462_v138_).length(0))
            rhs87_ = (len(d_420_v100_) if d_409_v89_.f21 else (d_413_v93_).f19)
            rhs88_ = (d_303_v2_) < (d_303_v2_)
            rhs89_ = (d_414_v94_) + (((d_462_v138_)[default__.safeIndex(670, (d_462_v138_).length(0))]) + (d_414_v94_))
            rhs90_ = d_303_v2_
            lhs66_ = d_305_globalState_
            lhs67_ = d_409_v89_
            lhs68_ = d_462_v138_
            lhs69_ = default__.safeIndex(670, (d_462_v138_).length(0))
            lhs66_.f3 = rhs87_
            lhs67_.f21 = rhs88_
            lhs68_[lhs69_] = rhs89_
            d_303_v2_ = rhs90_
            if d_409_v89_.f21:
                index66_ = default__.safeIndex(963, (d_302_v1_).length(0))
                (d_302_v1_)[index66_] = not((d_413_v93_).f18)
                index67_ = default__.safeIndex(963, (d_302_v1_).length(0))
                (d_302_v1_)[index67_] = d_304_v3_
                d_467_v140_: _dafny.Array
                nw81_ = _dafny.Array(False, 4)
                d_467_v140_ = nw81_
                (d_305_globalState_).f2 = d_467_v140_
                d_468_v141_: _dafny.Array
                def lambda31_(d_469_v102_):
                    def lambda32_(d_470_i7_):
                        return d_469_v102_

                    return lambda32_

                init16_ = lambda31_(d_422_v102_)
                nw82_ = _dafny.Array(None, 25)
                for i0_16_ in range(nw82_.length(0)):
                    nw82_[i0_16_] = init16_(i0_16_)
                d_468_v141_ = nw82_
                index68_ = default__.safeIndex(468, (d_468_v141_).length(0))
                (d_468_v141_)[index68_] = d_422_v102_
                index69_ = default__.safeIndex(468, (d_468_v141_).length(0))
                (d_468_v141_)[index69_] = D0_DC1((d_413_v93_).f19, default__.fm0((d_413_v93_).f19, (d_413_v93_).f19, d_413_v93_.f20, (d_413_v93_).f19, d_305_globalState_), (len(d_382_v70_)) <= (d_413_v93_.f20))
                d_471_v142_: _dafny.Array
                nw83_ = _dafny.Array(_dafny.Seq({}), 27)
                d_471_v142_ = nw83_
                d_471_v142_ = d_471_v142_
                index70_ = default__.safeIndex(473, (d_306_v4_).length(0))
                (d_306_v4_)[index70_] = (d_413_v93_).f19
                index71_ = default__.safeIndex(473, (d_306_v4_).length(0))
                (d_306_v4_)[index71_] = len((d_419_v99_) + (d_419_v99_))
            elif True:
                (d_305_globalState_).f3 = (d_413_v93_).f19
                d_472_v143_: int
                d_473_v144_: _dafny.Map
                out30_: int
                out31_: _dafny.Map
                out30_, out31_ = default__.m0(d_413_v93_.f20, len(default__.fm2(d_412_v92_, d_409_v89_.f21, d_305_globalState_)), d_305_globalState_)
                d_472_v143_ = out30_
                d_473_v144_ = out31_
                (d_305_globalState_).f9 = d_472_v143_
                rhs91_ = d_409_v89_.f21
                rhs92_ = d_306_v4_
                rhs93_ = d_472_v143_
                lhs70_ = d_305_globalState_
                d_304_v3_ = rhs91_
                d_306_v4_ = rhs92_
                lhs70_.f3 = rhs93_
                index72_ = default__.safeIndex(788, (d_302_v1_).length(0))
                (d_302_v1_)[index72_] = (d_413_v93_).f18
                index73_ = default__.safeIndex(788, (d_302_v1_).length(0))
                (d_302_v1_)[index73_] = d_409_v89_.f21
            d_474_v145_: _dafny.Map
            d_474_v145_ = _dafny.Map({d_414_v94_: (0) - ((d_413_v93_).f19)})
            d_475_v146_: _dafny.MultiSet
            d_475_v146_ = _dafny.MultiSet([len(d_420_v100_), len(d_474_v145_)])
            d_476_v147_: _dafny.Array
            d_477_v148_: _dafny.Set
            out32_: _dafny.Array
            out33_: _dafny.Set
            out32_, out33_ = (d_409_v89_).m2(d_475_v146_, (d_413_v93_).f19, d_302_v1_, default__.fm3(d_305_globalState_), d_305_globalState_)
            d_476_v147_ = out32_
            d_477_v148_ = out33_
        index74_ = default__.safeIndex(150, (d_306_v4_).length(0))
        (d_306_v4_)[index74_] = (d_413_v93_).f19
        d_478_v149_: _dafny.Array
        def lambda33_(d_479_v101_):
            def lambda34_(d_480_i8_):
                return d_479_v101_

            return lambda34_

        init17_ = lambda33_(d_421_v101_)
        nw84_ = _dafny.Array(None, 28)
        for i0_17_ in range(nw84_.length(0)):
            nw84_[i0_17_] = init17_(i0_17_)
        d_478_v149_ = nw84_
        d_481_v150_: _dafny.Map
        d_481_v150_ = d_421_v101_
        index75_ = default__.safeIndex(779, (d_478_v149_).length(0))
        (d_478_v149_)[index75_] = (d_481_v150_ if not(d_409_v89_.f21) else d_481_v150_)
        index76_ = default__.safeIndex(150, (d_306_v4_).length(0))
        index77_ = default__.safeIndex(779, (d_478_v149_).length(0))
        rhs94_ = d_413_v93_.f20
        rhs95_ = d_481_v150_
        rhs96_ = (default__.fm0(d_413_v93_.f20, (0) - (d_413_v93_.f20), default__.fm0(-511, d_311_v7_, d_413_v93_.f20, d_311_v7_, d_305_globalState_), d_413_v93_.f20, d_305_globalState_)) < ((len(d_414_v94_) if d_304_v3_ else d_413_v93_.f20))
        lhs71_ = d_306_v4_
        lhs72_ = default__.safeIndex(150, (d_306_v4_).length(0))
        lhs73_ = d_478_v149_
        lhs74_ = default__.safeIndex(779, (d_478_v149_).length(0))
        lhs75_ = d_409_v89_
        lhs71_[lhs72_] = rhs94_
        lhs73_[lhs74_] = rhs95_
        lhs75_.f21 = rhs96_
        pat_let_tv13_ = d_412_v92_
        pat_let_tv14_ = d_412_v92_
        pat_let_tv15_ = d_412_v92_
        def lambda35_(source4_):
            if source4_.is_DC6:
                d_482___mcc_h4_ = source4_.cf10
                d_483_cf10_ = d_482___mcc_h4_
                return pat_let_tv13_
            elif source4_.is_DC5:
                d_484___mcc_h5_ = source4_.cf9
                d_485_cf9_ = d_484___mcc_h5_
                return pat_let_tv14_
            elif True:
                d_486___mcc_h6_ = source4_.cf11
                d_487_cf11_ = d_486___mcc_h6_
                return pat_let_tv15_

        d_412_v92_ = lambda35_((d_425_v105_ if True else d_425_v105_))
        (d_305_globalState_).f3 = 440
        d_488_v151_: _dafny.Array
        d_489_v152_: _dafny.Set
        out34_: _dafny.Array
        out35_: _dafny.Set
        out34_, out35_ = (d_409_v89_).m2(default__.fm12(d_305_globalState_), (d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], d_302_v1_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmgvn")) if d_409_v89_.f21 else d_303_v2_), d_305_globalState_)
        d_488_v151_ = out34_
        d_489_v152_ = out35_
        if ((d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))]) < (d_413_v93_.f20):
            d_490_v153_: _dafny.MultiSet
            d_490_v153_ = _dafny.MultiSet([d_413_v93_.f20])
            d_491_v154_: _dafny.MultiSet
            d_491_v154_ = _dafny.MultiSet([(d_413_v93_).f19, default__.fm0((d_490_v153_).cardinality, (d_413_v93_).f19, (0) - (len(d_303_v2_)), (d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], d_305_globalState_)])
            (d_413_v93_).f20 = default__.fm0(d_311_v7_, 790, ((d_491_v154_)[len(d_420_v100_)] if (len(d_420_v100_)) in (d_491_v154_) else default__.fm0(d_311_v7_, d_311_v7_, 250, len(d_420_v100_), d_305_globalState_)), (d_413_v93_).f19, d_305_globalState_)
            d_492_v155_: _dafny.Map
            d_492_v155_ = _dafny.Map({d_409_v89_: d_409_v89_.f21})
            d_493_v156_: D1
            d_493_v156_ = D1_DC3(len(d_489_v152_))
            source5_ = (d_493_v156_ if (d_492_v155_) == (d_492_v155_) else d_493_v156_)
            if source5_.is_DC3:
                d_494___mcc_h7_ = source5_.cf5
                d_495_cf5_ = d_494___mcc_h7_
                d_496_v157_: D5
                d_496_v157_ = D5_DC12(d_413_v93_)
                d_497_v158_: _dafny.Map
                d_497_v158_ = _dafny.Map({d_409_v89_.f21: d_413_v93_})
                d_498_v159_: _dafny.Array
                nw85_ = _dafny.Array(None, 18)
                nw85_[int(0)] = d_413_v93_
                nw85_[int(1)] = d_413_v93_
                nw85_[int(2)] = d_413_v93_
                nw85_[int(3)] = d_413_v93_
                nw85_[int(4)] = d_413_v93_
                nw85_[int(5)] = d_413_v93_
                nw85_[int(6)] = (d_496_v157_).cf18
                nw85_[int(7)] = d_413_v93_
                nw85_[int(8)] = ((d_497_v158_)[d_409_v89_.f21] if (d_409_v89_.f21) in (d_497_v158_) else d_413_v93_)
                nw85_[int(9)] = d_413_v93_
                nw85_[int(10)] = d_413_v93_
                nw85_[int(11)] = d_413_v93_
                nw85_[int(12)] = d_413_v93_
                nw85_[int(13)] = d_413_v93_
                nw85_[int(14)] = d_413_v93_
                nw85_[int(15)] = d_413_v93_
                nw85_[int(16)] = d_413_v93_
                nw85_[int(17)] = (d_413_v93_ if d_409_v89_.f21 else d_413_v93_)
                d_498_v159_ = nw85_
                rhs97_ = d_498_v159_
                rhs98_ = d_303_v2_
                rhs99_ = d_409_v89_
                lhs76_ = d_305_globalState_
                d_498_v159_ = rhs97_
                lhs76_.f4 = rhs98_
                d_409_v89_ = rhs99_
                (d_305_globalState_).f3 = d_495_cf5_
                (d_409_v89_).f21 = d_304_v3_
                (d_305_globalState_).f4 = (d_303_v2_ if d_409_v89_.f21 else d_303_v2_)
            elif source5_.is_DC4:
                d_499___mcc_h8_ = source5_.cf6
                d_500___mcc_h9_ = source5_.cf7
                d_501___mcc_h10_ = source5_.cf8
                d_502_cf8_ = d_501___mcc_h10_
                d_503_cf7_ = d_500___mcc_h9_
                d_504_cf6_ = d_499___mcc_h8_
                d_421_v101_ = _dafny.Map({(d_413_v93_).f18: (d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))]})
                d_503_cf7_ = (d_413_v93_).f18
                d_504_cf6_ = d_503_cf7_
                d_505_v160_: _dafny.Seq
                d_505_v160_ = _dafny.SeqWithoutIsStrInference([d_303_v2_, _dafny.SeqWithoutIsStrInference([d_412_v92_ for d_506_i9_ in range(default__.abs(410))]), d_303_v2_])
                d_414_v94_ = (_dafny.SeqWithoutIsStrInference([(d_304_v3_) == (d_409_v89_.f21), d_304_v3_])).set(default__.safeIndex(len(d_505_v160_), len(_dafny.SeqWithoutIsStrInference([(d_304_v3_) == (d_409_v89_.f21), d_304_v3_]))), (len(default__.fm13((d_413_v93_).f18, ((d_418_v98_)[d_304_v3_] if (d_304_v3_) in (d_418_v98_) else d_502_cf8_), d_413_v93_.f20, len(_dafny.SeqWithoutIsStrInference([389])), d_305_globalState_))) < (d_413_v93_.f20))
            elif True:
                d_507___mcc_h11_ = source5_.cf4
                d_508_cf4_ = d_507___mcc_h11_
                (d_409_v89_).f21 = not((d_303_v2_) <= (d_303_v2_))
                index78_ = default__.safeIndex(150, (d_306_v4_).length(0))
                index79_ = default__.safeIndex(150, (d_306_v4_).length(0))
                rhs100_ = d_413_v93_.f20
                rhs101_ = d_384_v72_
                rhs102_ = (d_413_v93_.f20) - (d_311_v7_)
                lhs77_ = d_306_v4_
                lhs78_ = default__.safeIndex(150, (d_306_v4_).length(0))
                lhs79_ = d_306_v4_
                lhs80_ = default__.safeIndex(150, (d_306_v4_).length(0))
                lhs77_[lhs78_] = rhs100_
                d_384_v72_ = rhs101_
                lhs79_[lhs80_] = rhs102_
                index80_ = default__.safeIndex(150, (d_306_v4_).length(0))
                (d_306_v4_)[index80_] = ((0) - ((d_413_v93_).f19)) + (d_413_v93_.f20)
                d_509_v161_: _dafny.Seq
                d_509_v161_ = _dafny.SeqWithoutIsStrInference([d_303_v2_])
                index81_ = default__.safeIndex(150, (d_306_v4_).length(0))
                rhs103_ = d_306_v4_
                rhs104_ = d_409_v89_.f21
                rhs105_ = d_509_v161_
                rhs106_ = 88
                rhs107_ = d_412_v92_
                lhs81_ = d_306_v4_
                lhs82_ = default__.safeIndex(150, (d_306_v4_).length(0))
                d_306_v4_ = rhs103_
                d_304_v3_ = rhs104_
                d_509_v161_ = rhs105_
                lhs81_[lhs82_] = rhs106_
                d_412_v92_ = rhs107_
            index82_ = default__.safeIndex(312, (d_302_v1_).length(0))
            (d_302_v1_)[index82_] = (not(d_409_v89_.f21)) == ((d_413_v93_).f18)
            index83_ = default__.safeIndex(312, (d_302_v1_).length(0))
            (d_302_v1_)[index83_] = d_409_v89_.f21
            d_510_v162_: C1
            nw86_ = C1()
            nw86_.ctor__(((d_413_v93_).f18 if True else (D0_DC0(d_304_v3_)).cf0), d_413_v93_.f20, d_413_v93_.f20, d_409_v89_.f21)
            d_510_v162_ = nw86_
            if (len(d_303_v2_)) > ((d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))]):
                (d_305_globalState_).f5 = d_311_v7_
                (d_413_v93_).f20 = ((d_413_v93_).f19) * (d_413_v93_.f20)
                d_511_v163_: D2
                d_511_v163_ = D2_DC6(d_491_v154_)
                d_512_v164_: D2
                d_512_v164_ = D2_DC6((d_511_v163_).cf10)
                d_513_v165_: T0
                nw87_ = C0()
                nw87_.ctor__(d_512_v164_, d_413_v93_.f20, (d_413_v93_).f18)
                d_513_v165_ = nw87_
                d_514_v166_: _dafny.Seq
                d_514_v166_ = _dafny.SeqWithoutIsStrInference([d_513_v165_, d_513_v165_, d_513_v165_])
                d_515_v167_: _dafny.Map
                d_515_v167_ = _dafny.Map({(d_514_v166_)[default__.safeIndex((d_413_v93_).f19, len(d_514_v166_))]: d_409_v89_.f21})
                d_515_v167_ = (d_515_v167_).set(d_513_v165_, (d_413_v93_).f18)
                d_516_v168_: D2
                d_516_v168_ = D2_DC5(d_382_v70_)
                d_517_v169_: D2
                d_517_v169_ = D2_DC7(d_516_v168_)
                d_518_v170_: D2
                d_518_v170_ = D2_DC7(d_516_v168_)
                d_519_v171_: _dafny.Map
                d_519_v171_ = _dafny.Map({d_518_v170_: _dafny.Map({d_510_v162_.f21: len(d_303_v2_)})})
                d_519_v171_ = (d_519_v171_).set((d_518_v170_ if False else d_518_v170_), d_421_v101_)
                d_520_v172_: _dafny.Map
                d_520_v172_ = _dafny.Map({d_413_v93_.f20: (d_413_v93_).f18})
                d_520_v172_ = (d_520_v172_).set(default__.safeDivisionInt(d_413_v93_.f20, (d_413_v93_).f19), not(not(default__.fm1((d_413_v93_).f19, d_510_v162_.f21, (_dafny.MultiSet([(d_413_v93_).f18])).set((d_413_v93_).fm5(d_510_v162_.f21, d_413_v93_.f20, d_305_globalState_), default__.abs((d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))])), d_305_globalState_))))
            elif True:
                d_521_v173_: _dafny.Array
                nw88_ = _dafny.Array(False, 7)
                d_521_v173_ = nw88_
                d_522_v174_: _dafny.Array
                d_523_v175_: _dafny.Set
                out36_: _dafny.Array
                out37_: _dafny.Set
                out36_, out37_ = (d_413_v93_).m2((d_491_v154_).set((0) - (d_413_v93_.f20), default__.abs(d_413_v93_.f20)), len(_dafny.SeqWithoutIsStrInference([d_409_v89_.f21])), d_521_v173_, _dafny.SeqWithoutIsStrInference([d_412_v92_ for d_524_i10_ in range(default__.abs(368))]), d_305_globalState_)
                d_522_v174_ = out36_
                d_523_v175_ = out37_
                d_525_v176_: _dafny.Array
                nw89_ = _dafny.Array(None, 15)
                d_525_v176_ = nw89_
                index84_ = default__.safeIndex(767, (d_525_v176_).length(0))
                (d_525_v176_)[index84_] = d_409_v89_
                index85_ = default__.safeIndex(767, (d_525_v176_).length(0))
                (d_525_v176_)[index85_] = d_510_v162_
                d_526_v177_: _dafny.Map
                d_526_v177_ = _dafny.Map({(d_303_v2_ if (d_422_v102_).cf3 else d_303_v2_): d_409_v89_.f21})
                d_526_v177_ = (d_526_v177_).set((d_303_v2_) + (d_303_v2_), d_409_v89_.f21)
                d_527_v178_: D2
                d_527_v178_ = D2_DC6(default__.fm12(d_305_globalState_))
                d_528_v179_: _dafny.Map
                d_528_v179_ = _dafny.Map({d_311_v7_: _dafny.MultiSet(d_414_v94_)})
                d_529_v180_: C0
                nw90_ = C0()
                nw90_.ctor__(D2_DC6(d_491_v154_), (d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], default__.fm1(len(d_303_v2_), d_409_v89_.f21, ((d_528_v179_)[d_311_v7_] if (d_311_v7_) in (d_528_v179_) else d_417_v97_), d_305_globalState_))
                d_529_v180_ = nw90_
                d_530_v181_: _dafny.Map
                d_530_v181_ = _dafny.Map({d_529_v180_: not((d_413_v93_).f18)})
                d_531_v182_: T0
                nw91_ = C0()
                nw91_.ctor__(d_527_v178_, (len(d_530_v181_)) - ((d_417_v97_).cardinality), not(d_409_v89_.f21))
                d_531_v182_ = nw91_
                d_531_v182_ = d_531_v182_
                (d_305_globalState_).f13 = d_510_v162_.f21
        elif True:
            if (d_304_v3_ if d_409_v89_.f21 else False):
                d_532_v183_: _dafny.Array
                d_533_v184_: _dafny.Set
                out38_: _dafny.Array
                out39_: _dafny.Set
                out38_, out39_ = (d_409_v89_).m2(_dafny.MultiSet([608]), d_413_v93_.f20, d_302_v1_, d_303_v2_, d_305_globalState_)
                d_532_v183_ = out38_
                d_533_v184_ = out39_
                def lambda36_(d_534_v89_):
                    def lambda37_(d_535_i11_):
                        return d_534_v89_.f21

                    return lambda37_

                init18_ = lambda36_(d_409_v89_)
                nw92_ = _dafny.Array(None, 14)
                for i0_18_ in range(nw92_.length(0)):
                    nw92_[i0_18_] = init18_(i0_18_)
                d_302_v1_ = nw92_
                d_536_v185_: _dafny.Map
                d_536_v185_ = _dafny.Map({d_414_v94_: d_420_v100_})
                d_536_v185_ = (d_536_v185_).set((_dafny.SeqWithoutIsStrInference([d_304_v3_, d_409_v89_.f21, True, d_304_v3_, d_409_v89_.f21])).set(default__.safeIndex(816, len(_dafny.SeqWithoutIsStrInference([d_304_v3_, d_409_v89_.f21, True, d_304_v3_, d_409_v89_.f21]))), d_409_v89_.f21), d_420_v100_)
                d_537_v186_: _dafny.Set
                d_537_v186_ = _dafny.Set({(d_413_v93_).f18})
                d_538_v187_: _dafny.MultiSet
                d_538_v187_ = _dafny.MultiSet([d_413_v93_.f20, len(d_537_v186_)])
                pat_let_tv16_ = d_538_v187_
                d_539_v188_: C0
                nw93_ = C0()
                def iife31_(_pat_let11_0):
                    def iife32_(d_540_dt__update__tmp_h2_):
                        def iife33_(_pat_let12_0):
                            def iife34_(d_541_dt__update_hcf10_h1_):
                                return D2_DC6(d_541_dt__update_hcf10_h1_)
                            return iife34_(_pat_let12_0)
                        return iife33_(pat_let_tv16_)
                    return iife32_(_pat_let11_0)
                def iife30_(_pat_let10_0):
                    def iife35_(d_542_dt__update__tmp_h3_):
                        def iife36_(_pat_let13_0):
                            def iife37_(d_543_dt__update_hcf10_h2_):
                                return D2_DC6(d_543_dt__update_hcf10_h2_)
                            return iife37_(_pat_let13_0)
                        return iife36_(_dafny.MultiSet([286]))
                    return iife35_(_pat_let10_0)
                nw93_.ctor__(iife30_(iife31_(D2_DC6(d_538_v187_))), (d_413_v93_).f19, (d_413_v93_).f18)
                d_539_v188_ = nw93_
                d_544_v189_: _dafny.Set
                d_544_v189_ = _dafny.Set({(D6_DC14(d_539_v188_)).cf20})
                d_545_v190_: _dafny.MultiSet
                d_545_v190_ = _dafny.MultiSet([(d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], (d_413_v93_).f19, len(d_544_v189_), 360])
                d_546_v191_: _dafny.Array
                d_547_v192_: _dafny.Set
                out40_: _dafny.Array
                out41_: _dafny.Set
                out40_, out41_ = (d_409_v89_).m2(d_538_v187_, d_413_v93_.f20, d_302_v1_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xduc"))) + (_dafny.SeqWithoutIsStrInference([d_412_v92_ for d_548_i12_ in range(default__.abs(-917))])), d_305_globalState_)
                d_546_v191_ = out40_
                d_547_v192_ = out41_
                d_549_v193_: _dafny.Map
                d_549_v193_ = _dafny.Map({(d_413_v93_).f18: d_412_v92_})
                d_550_v194_: _dafny.Map
                d_550_v194_ = _dafny.Map({d_409_v89_.f21: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qafu"))).set(default__.safeIndex((d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qafu")))), d_412_v92_)})
                d_551_v195_: _dafny.Array
                d_552_v196_: _dafny.Set
                out42_: _dafny.Array
                out43_: _dafny.Set
                out42_, out43_ = (d_409_v89_).m2((_dafny.MultiSet([d_413_v93_.f20, default__.fm0((d_413_v93_).f19, (d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], default__.fm0(81, d_311_v7_, (d_413_v93_).f19, (d_413_v93_).f19, d_305_globalState_), (d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], d_305_globalState_)])).intersection(d_545_v190_), default__.safeDivisionInt(default__.fm0(len(d_547_v192_), 892, default__.fm0((d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], (d_413_v93_).f19, len(d_544_v189_), d_413_v93_.f20, d_305_globalState_), len(_dafny.Map({(d_413_v93_).f19: True})), d_305_globalState_), len(d_549_v193_)), d_302_v1_, ((d_550_v194_)[(d_413_v93_).f18] if ((d_413_v93_).f18) in (d_550_v194_) else d_303_v2_), d_305_globalState_)
                d_551_v195_ = out42_
                d_552_v196_ = out43_
            elif True:
                d_553_v197_: _dafny.Set
                d_553_v197_ = _dafny.Set({d_412_v92_})
                pat_let_tv17_ = d_553_v197_
                pat_let_tv18_ = d_409_v89_
                pat_let_tv19_ = d_413_v93_
                pat_let_tv20_ = d_413_v93_
                pat_let_tv21_ = d_305_globalState_
                pat_let_tv22_ = d_413_v93_
                def iife38_(_pat_let14_0):
                    def iife39_(d_554_dt__update__tmp_h4_):
                        def iife40_(_pat_let15_0):
                            def iife41_(d_555_dt__update_hcf3_h0_):
                                def iife42_(_pat_let16_0):
                                    def iife43_(d_556_dt__update_hcf1_h0_):
                                        return D0_DC1(d_556_dt__update_hcf1_h0_, (d_554_dt__update__tmp_h4_).cf2, d_555_dt__update_hcf3_h0_)
                                    return iife43_(_pat_let16_0)
                                return iife42_(pat_let_tv22_.f20)
                            return iife41_(_pat_let15_0)
                        return iife40_((pat_let_tv17_).issubset(default__.fm14(pat_let_tv18_.f21, (pat_let_tv19_).f18, (pat_let_tv20_).f19, pat_let_tv21_)))
                    return iife39_(_pat_let14_0)
                d_422_v102_ = iife38_(d_422_v102_)
                d_421_v101_ = (d_421_v101_).set(False, (d_413_v93_).f19)
                d_557_v198_: _dafny.MultiSet
                d_557_v198_ = _dafny.MultiSet([len(d_489_v152_), 122])
                d_558_v200_: _dafny.MultiSet
                d_558_v200_ = _dafny.MultiSet([d_303_v2_, d_303_v2_])
                d_559_v201_: _dafny.Array
                d_560_v202_: _dafny.Set
                out44_: _dafny.Array
                out45_: _dafny.Set
                def iife44_():
                    coll10_ = _dafny.Map()
                    compr_10_: _dafny.Seq
                    for compr_10_ in (d_558_v200_).Elements:
                        d_561_v199_: _dafny.Seq = compr_10_
                        if (d_561_v199_) in (d_558_v200_):
                            coll10_[d_561_v199_] = (d_413_v93_).f18
                    return _dafny.Map(coll10_)
                out44_, out45_ = (d_409_v89_).m2(d_557_v198_, len(iife44_()
                ), d_302_v1_, d_303_v2_, d_305_globalState_)
                d_559_v201_ = out44_
                d_560_v202_ = out45_
                d_562_v203_: _dafny.Map
                d_562_v203_ = _dafny.Map({d_302_v1_: d_306_v4_})
                d_563_v204_: D7
                d_563_v204_ = D7_DC18(d_562_v203_)
                d_562_v203_ = ((d_563_v204_).cf25) | ((d_562_v203_) | (_dafny.Map({d_302_v1_: d_306_v4_})))
                (d_409_v89_).f21 = ((default__.fm7((d_413_v93_).f19, (d_413_v93_).f19, (d_413_v93_).f18, d_305_globalState_)).intersection(d_417_v97_)).issubset(_dafny.MultiSet([d_304_v3_, (d_413_v93_).f18]))
            d_564_v205_: _dafny.Seq
            d_564_v205_ = _dafny.SeqWithoutIsStrInference([d_302_v1_, d_302_v1_, d_302_v1_])
            d_565_v206_: _dafny.Array
            nw94_ = _dafny.Array(None, 15)
            nw94_[int(0)] = (d_564_v205_)[default__.safeIndex((d_417_v97_).cardinality, len(d_564_v205_))]
            nw94_[int(1)] = d_302_v1_
            nw94_[int(2)] = d_302_v1_
            nw94_[int(3)] = d_302_v1_
            nw94_[int(4)] = d_302_v1_
            nw94_[int(5)] = d_302_v1_
            nw94_[int(6)] = d_302_v1_
            nw94_[int(7)] = d_302_v1_
            nw94_[int(8)] = d_302_v1_
            nw94_[int(9)] = d_302_v1_
            nw94_[int(10)] = d_302_v1_
            nw94_[int(11)] = d_302_v1_
            nw94_[int(12)] = (d_564_v205_)[default__.safeIndex(d_413_v93_.f20, len(d_564_v205_))]
            nw94_[int(13)] = d_302_v1_
            nw94_[int(14)] = d_302_v1_
            d_565_v206_ = nw94_
            d_565_v206_ = d_565_v206_
            d_566_v207_: D6
            d_566_v207_ = D6_DC17((d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], d_304_v3_)
            rhs108_ = d_413_v93_.f20
            rhs109_ = (925) > (((d_413_v93_).f19) * ((d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))]))
            rhs110_ = d_311_v7_
            rhs111_ = (d_566_v207_).cf23
            lhs83_ = d_413_v93_
            lhs84_ = d_305_globalState_
            lhs85_ = d_305_globalState_
            lhs86_ = d_305_globalState_
            lhs83_.f20 = rhs108_
            lhs84_.f13 = rhs109_
            lhs85_.f5 = rhs110_
            lhs86_.f9 = rhs111_
            if not((((d_303_v2_).set(default__.safeIndex((d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], len(d_303_v2_)), d_412_v92_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fa")))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))):
                (d_305_globalState_).f13 = ((D1_DC3(d_311_v7_)).cf5) != (d_413_v93_.f20)
                (d_305_globalState_).f9 = d_413_v93_.f20
                d_567_v208_: int
                d_568_v209_: bool
                d_569_v210_: int
                d_570_v211_: int
                out46_: int
                out47_: bool
                out48_: int
                out49_: int
                out46_, out47_, out48_, out49_ = (d_413_v93_).m1(d_409_v89_.f21, (d_413_v93_).f18, d_305_globalState_)
                d_567_v208_ = out46_
                d_568_v209_ = out47_
                d_569_v210_ = out48_
                d_570_v211_ = out49_
                d_571_v212_: D7
                d_571_v212_ = D7_DC19(d_409_v89_.f21, d_412_v92_, d_568_v209_)
                d_572_v213_: C1
                nw95_ = C1()
                nw95_.ctor__(d_409_v89_.f21, (d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))], (d_413_v93_).f19, (d_571_v212_).cf28)
                d_572_v213_ = nw95_
                d_573_v214_: _dafny.MultiSet
                d_573_v214_ = _dafny.MultiSet([(d_413_v93_).f19])
                d_574_v215_: D2
                d_574_v215_ = D2_DC6(d_573_v214_)
                d_575_v216_: C0
                nw96_ = C0()
                nw96_.ctor__(d_574_v215_, (d_413_v93_).f19, d_304_v3_)
                d_575_v216_ = nw96_
                d_576_v217_: _dafny.Map
                d_576_v217_ = _dafny.Map({d_575_v216_: -175})
                d_577_v218_: _dafny.Map
                d_577_v218_ = _dafny.Map({888: d_303_v2_})
                d_578_v219_: _dafny.MultiSet
                d_578_v219_ = _dafny.MultiSet([d_576_v217_, _dafny.Map({d_575_v216_: len(d_577_v218_)})])
                d_579_v220_: C1
                nw97_ = C1()
                nw97_.ctor__((d_413_v93_).f18, d_311_v7_, (len(default__.fm13(d_572_v213_.f21, (d_413_v93_).f18, ((d_578_v219_)[d_576_v217_] if (d_576_v217_) in (d_578_v219_) else d_413_v93_.f20), (0) - ((d_575_v216_).f23), d_305_globalState_))) - ((_dafny.MultiSet(d_420_v100_)).cardinality), d_409_v89_.f21)
                d_579_v220_ = nw97_
            elif True:
                (d_305_globalState_).f9 = ((d_413_v93_).f19 if default__.fm1((d_420_v100_)[default__.safeIndex(d_413_v93_.f20, len(d_420_v100_))], d_304_v3_, d_417_v97_, d_305_globalState_) else (d_413_v93_).f19)
                index86_ = default__.safeIndex(150, (d_306_v4_).length(0))
                (d_306_v4_)[index86_] = ((d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))]) + (d_413_v93_.f20)
                def iife45_():
                    coll11_ = _dafny.Map()
                    compr_11_: int
                    for compr_11_ in _dafny.IntegerRange(800, 926):
                        d_580_v221_: int = compr_11_
                        if ((800) <= (d_580_v221_)) and ((d_580_v221_) < (926)):
                            coll11_[(d_580_v221_) * ((d_413_v93_).f19)] = d_311_v7_
                    return _dafny.Map(coll11_)
                (d_305_globalState_).f5 = len(_dafny.Map({_dafny.MultiSet([d_413_v93_.f20]): iife45_()
                }))
                d_581_v222_: _dafny.MultiSet
                d_581_v222_ = _dafny.MultiSet([len(d_489_v152_)])
                d_582_v223_: D2
                d_582_v223_ = D2_DC6(d_581_v222_)
                d_583_v224_: T0
                nw98_ = C0()
                nw98_.ctor__(d_582_v223_, d_413_v93_.f20, True)
                d_583_v224_ = nw98_
                d_584_v225_: _dafny.Seq
                d_584_v225_ = _dafny.SeqWithoutIsStrInference([d_583_v224_, d_583_v224_])
                d_585_v226_: T0
                nw99_ = C0()
                nw99_.ctor__(d_582_v223_, len(d_584_v225_), d_304_v3_)
                d_585_v226_ = nw99_
                d_586_v227_: _dafny.Set
                d_586_v227_ = _dafny.Set({d_585_v226_})
                (d_305_globalState_).f3 = len(d_586_v227_)
                d_587_v228_: _dafny.Set
                d_587_v228_ = _dafny.Set({d_303_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njxobfr")), d_303_v2_, d_303_v2_})
                d_588_v230_: _dafny.Map
                def iife46_():
                    coll12_ = _dafny.Set()
                    compr_12_: _dafny.Seq
                    for compr_12_ in (d_587_v228_).Elements:
                        d_589_v229_: _dafny.Seq = compr_12_
                        if (d_589_v229_) in (d_587_v228_):
                            coll12_ = coll12_.union(_dafny.Set([d_589_v229_]))
                    return _dafny.Set(coll12_)
                d_588_v230_ = _dafny.Map({iife46_()
                : d_306_v4_})
                d_588_v230_ = (d_588_v230_).set(d_587_v228_, d_306_v4_)
            d_590_v231_: _dafny.MultiSet
            d_590_v231_ = _dafny.MultiSet([d_413_v93_.f20, (d_413_v93_.f20) + ((d_413_v93_).f19), len(d_420_v100_)])
            d_590_v231_ = ((d_590_v231_) - (d_590_v231_)).set(len(d_414_v94_), default__.abs((d_306_v4_)[default__.safeIndex(150, (d_306_v4_).length(0))]))
        _dafny.print(_dafny.string_of((d_302_v1_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v1_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_303_v2_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_304_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_globalState_.f2)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_305_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_305_globalState_.f4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_305_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_305_globalState_).f7).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_305_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_305_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_305_globalState_).f17)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v4_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_309_v5_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_310_v6_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_311_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v8_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_382_v70_) == (_dafny.Map({604: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_384_v72_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_409_v89_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_410_v90_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_411_v91_).cf13.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_411_v91_).cf13).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_411_v91_).cf13.f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_411_v91_).cf13).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_412_v92_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_413_v93_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_413_v93_.f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_413_v93_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_414_v94_) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_415_v95_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, False]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_416_v96_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_417_v97_) == (_dafny.MultiSet([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_418_v98_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_419_v99_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: False}), _dafny.Map({False: False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_420_v100_) == (_dafny.SeqWithoutIsStrInference([-186]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_421_v101_) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_422_v102_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_422_v102_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_422_v102_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_423_v103_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_424_v104_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_425_v105_).cf9) == (_dafny.Map({4: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[0])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[1])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[2])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[3])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[4])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[5])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[6])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[7])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[8])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[9])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[10])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[11])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[12])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[13])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[14])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[15])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[16])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[17])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[18])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[19])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[20])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[21])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[22])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[23])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[24])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[25])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[26])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_478_v149_)[27])) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_481_v150_)) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[0]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[1]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[2]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[3]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[4]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[5]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[6]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[7]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[8]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[9]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[10]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[11]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[12]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[13]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[14]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[15]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[16]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[17]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[18]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[19]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[20]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_488_v151_)[21]).cf10) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_489_v152_) == (_dafny.Set({2, -525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(_dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC8(D3, NamedTuple('DC8', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC10(D4, NamedTuple('DC10', [('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC13(D5, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC15(None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC15(D6, NamedTuple('DC15', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC14(D6, NamedTuple('DC14', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(False, _dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC19(D7, NamedTuple('DC19', [('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC22(D8, NamedTuple('DC22', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC24(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)

class D9_DC24(D9, NamedTuple('DC24', [('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC26(None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC26(D10, NamedTuple('DC26', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC27(D10, NamedTuple('DC27', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC25(D10, NamedTuple('DC25', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC30(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)

class D11_DC30(D11, NamedTuple('DC30', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)

class D12_DC33(D12, NamedTuple('DC33', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC35(_dafny.CodePoint('D'), D9.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)

class D13_DC35(D13, NamedTuple('DC35', [('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC34(D13, NamedTuple('DC34', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC36(D13, NamedTuple('DC36', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    @property
    def f20(self):
        return self._f20
    @f20.setter
    def f20(self, value):
        self._f20 = value
    def m1(self, p0, p1, globalState):
        pass

    def m2(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: _dafny.Array = _dafny.Array(None, 0)
        self.f3: int = int(0)
        self.f4: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f5: int = int(0)
        self.f9: int = int(0)
        self.f13: bool = False
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: bool = False
        self._f6: int = int(0)
        self._f7: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f8: int = int(0)
        self._f10: bool = False
        self._f11: int = int(0)
        self._f12: bool = False
        self._f14: int = int(0)
        self._f15: bool = False
        self._f16: bool = False
        self._f17: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
    @property
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17

class C0(T0):
    def  __init__(self):
        self._f18: bool = False
        self.f22: D2 = D2.default()()
        self._f23: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f22, f23, f18):
        (self).f22 = f22
        (self)._f23 = f23
        (self)._f18 = f18

    @property
    def f23(self):
        return self._f23

class C1(T1, T0):
    def  __init__(self):
        self._f20: int = int(0)
        self._f19: int = int(0)
        self._f18: bool = False
        self.f21: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f20(self):
        return self._f20
    @f20.setter
    def f20(self, value):
        self._f20 = value
    @property
    def f19(self):
        return self._f19
    @property
    def f18(self):
        return self._f18
    def ctor__(self, f21, f19, f20, f18):
        (self).f21 = f21
        (self)._f19 = f19
        (self).f20 = f20
        (self)._f18 = f18

    def fm5(self, p0, p1, globalState):
        return ((_dafny.MultiSet([354, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_591_i0_ in range(default__.abs(654))])), self.f20])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([444])))).isdisjoint(_dafny.MultiSet([(self).f19]))

    def fm6(self, p0, p1, p2, p3, globalState):
        def iife47_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in (_dafny.SeqWithoutIsStrInference([(0) - (self.f20), (self).f19])).Elements:
                d_592_v0_: int = compr_13_
                if (d_592_v0_) in (_dafny.SeqWithoutIsStrInference([(0) - (self.f20), (self).f19])):
                    coll13_[(d_592_v0_) - ((self).f19)] = (self).f19
            return _dafny.Map(coll13_)
        return ((_dafny.Map({989: self.f20})) | (_dafny.Map({(0) - (self.f20): self.f20}))) | ((iife47_()
        ) | (_dafny.Map({self.f20: len(_dafny.SeqWithoutIsStrInference([self.f21]))})))

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        r3: int = int(0)
        d_593_v0_: _dafny.Array
        def lambda38_(d_594_i0_):
            return (d_594_i0_) - ((self).f19)

        init19_ = lambda38_
        nw100_ = _dafny.Array(None, 8)
        for i0_19_ in range(nw100_.length(0)):
            nw100_[i0_19_] = init19_(i0_19_)
        d_593_v0_ = nw100_
        index87_ = default__.safeIndex(449, (d_593_v0_).length(0))
        (d_593_v0_)[index87_] = default__.safeModuloInt(self.f20, self.f20)
        d_595_v1_: _dafny.Map
        d_595_v1_ = _dafny.Map({self.f20: self.f20})
        d_596_v2_: _dafny.MultiSet
        d_596_v2_ = _dafny.MultiSet([self.f20])
        d_597_v3_: D2
        d_597_v3_ = D2_DC6(d_596_v2_)
        d_598_v4_: _dafny.MultiSet
        d_598_v4_ = _dafny.MultiSet([self.f21])
        d_599_v5_: T0
        nw101_ = C0()
        nw101_.ctor__(d_597_v3_, ((d_598_v4_)[not(p1)] if (not(p1)) in (d_598_v4_) else self.f20), (self).f18)
        d_599_v5_ = nw101_
        d_600_v6_: _dafny.Map
        d_600_v6_ = _dafny.Map({d_599_v5_: d_593_v0_})
        d_601_v7_: _dafny.Seq
        d_601_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oelkbtu"))
        index88_ = default__.safeIndex(449, (d_593_v0_).length(0))
        (d_593_v0_)[index88_] = (((d_595_v1_)[(self).f19] if ((self).f19) in (d_595_v1_) else len((d_600_v6_).set(d_599_v5_, d_593_v0_)))) - (len(d_601_v7_))
        d_602_v8_: _dafny.Array
        nw102_ = _dafny.Array(_dafny.MultiSet({}), 6)
        d_602_v8_ = nw102_
        _ingredientsd_0 = []
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_602_v8_).length(0)):
            d_603_i1_: int = guard_loop_0_
            if (True) and (((0) <= (d_603_i1_)) and ((d_603_i1_) < ((d_602_v8_).length(0)))):
                _ingredientsd_0.append((d_602_v8_, int(d_603_i1_), (d_596_v2_).set(self.f20, default__.abs((d_593_v0_)[default__.safeIndex(449, (d_593_v0_).length(0))]))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        hi1_ = len(d_601_v7_)
        for d_604_i2_ in range(719, hi1_):
            d_605_v9_: _dafny.Set
            d_605_v9_ = _dafny.Set({p0, p0, True})
            d_606_v10_: _dafny.MultiSet
            d_606_v10_ = _dafny.MultiSet([d_605_v9_, d_605_v9_, d_605_v9_])
            (globalState).f5 = ((d_595_v1_)[((d_606_v10_)[d_605_v9_] if (d_605_v9_) in (d_606_v10_) else (self).f19)] if (((d_606_v10_)[d_605_v9_] if (d_605_v9_) in (d_606_v10_) else (self).f19)) in (d_595_v1_) else ((d_595_v1_)[self.f20] if (self.f20) in (d_595_v1_) else d_604_i2_))
            d_607_v11_: D1
            d_607_v11_ = D1_DC3(980)
            d_608_v12_: _dafny.Map
            d_608_v12_ = _dafny.Map({((self).f18 if (d_599_v5_).f18 else (self).f18): (0) - ((d_607_v11_).cf5)})
            d_609_v13_: _dafny.Map
            d_609_v13_ = _dafny.Map({(d_599_v5_).f18: default__.fm1((d_593_v0_)[default__.safeIndex(449, (d_593_v0_).length(0))], (d_599_v5_).f18, d_598_v4_, globalState)})
            d_610_v14_: _dafny.Map
            d_610_v14_ = _dafny.Map({d_599_v5_: (d_593_v0_)[default__.safeIndex(449, (d_593_v0_).length(0))]})
            d_611_v15_: _dafny.Map
            d_611_v15_ = _dafny.Map({d_596_v2_: d_610_v14_})
            d_608_v12_ = (d_608_v12_).set((self).f18, default__.fm0(d_604_i2_, d_604_i2_, len(d_609_v13_), len(d_611_v15_), globalState))
            r0 = (d_593_v0_)[default__.safeIndex(449, (d_593_v0_).length(0))]
            index89_ = default__.safeIndex(449, (d_593_v0_).length(0))
            (d_593_v0_)[index89_] = (d_607_v11_).cf5
        d_612_v16_: D0
        d_612_v16_ = D0_DC0((self).f18)
        d_613_v17_: _dafny.Seq
        d_613_v17_ = _dafny.SeqWithoutIsStrInference([(d_593_v0_)[default__.safeIndex(449, (d_593_v0_).length(0))], self.f20])
        d_614_v18_: _dafny.Array
        nw103_ = _dafny.Array(None, 11)
        nw103_[int(0)] = True
        nw103_[int(1)] = False
        nw103_[int(2)] = p1
        nw103_[int(3)] = not((d_612_v16_).cf0)
        nw103_[int(4)] = not((self.f20) == (self.f20))
        nw103_[int(5)] = p1
        nw103_[int(6)] = not(False)
        nw103_[int(7)] = p0
        nw103_[int(8)] = (d_599_v5_).f18
        nw103_[int(9)] = p1
        nw103_[int(10)] = (_dafny.SeqWithoutIsStrInference([(0) - ((d_593_v0_)[default__.safeIndex(449, (d_593_v0_).length(0))]) for d_615_i3_ in range(default__.abs(632))])) < (d_613_v17_)
        d_614_v18_ = nw103_
        index90_ = default__.safeIndex(650, (d_614_v18_).length(0))
        (d_614_v18_)[index90_] = (self).fm5(False, self.f20, globalState)
        index91_ = default__.safeIndex(650, (d_614_v18_).length(0))
        (d_614_v18_)[index91_] = (p0 if (d_598_v4_).issubset(d_598_v4_) else p0)
        d_598_v4_ = (_dafny.MultiSet([True, p0])) | ((default__.fm7((self).f19, (d_593_v0_)[default__.safeIndex(449, (d_593_v0_).length(0))], p1, globalState)) | (d_598_v4_))
        d_616_v19_: _dafny.Seq
        d_616_v19_ = _dafny.SeqWithoutIsStrInference([(d_599_v5_).f18, (d_614_v18_)[default__.safeIndex(650, (d_614_v18_).length(0))]])
        r3 = ((self.f20 if (self).f18 else ((d_596_v2_)[len(d_616_v19_)] if (len(d_616_v19_)) in (d_596_v2_) else self.f20))) - ((self).f19)
        r0 = (self).f19
        d_617_v20_: _dafny.Seq
        d_617_v20_ = _dafny.SeqWithoutIsStrInference([d_599_v5_])
        r1 = (d_599_v5_) in (d_617_v20_)
        r2 = (d_593_v0_)[default__.safeIndex(449, (d_593_v0_).length(0))]
        r3 = self.f20
        return r0, r1, r2, r3

    def m2(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Set = _dafny.Set({})
        d_618_v0_: _dafny.Array
        nw104_ = _dafny.Array(int(0), 27)
        d_618_v0_ = nw104_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_618_v0_).length(0)):
            d_619_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_619_i0_)) and ((d_619_i0_) < ((d_618_v0_).length(0)))):
                (d_618_v0_)[(d_619_i0_)] = (d_619_i0_) - (self.f20)
        d_620_v1_: D2
        d_620_v1_ = D2_DC6(p0)
        d_621_v2_: _dafny.Seq
        d_621_v2_ = _dafny.SeqWithoutIsStrInference([p1])
        pat_let_tv23_ = d_620_v1_
        d_622_v3_: _dafny.MultiSet
        def iife48_(_pat_let17_0):
            def iife49_(d_623_dt__update__tmp_h0_):
                def iife50_(_pat_let18_0):
                    def iife51_(d_624_dt__update_hcf10_h0_):
                        return D2_DC6(d_624_dt__update_hcf10_h0_)
                    return iife51_(_pat_let18_0)
                return iife50_((pat_let_tv23_).cf10)
            return iife49_(_pat_let17_0)
        d_622_v3_ = _dafny.MultiSet([d_620_v1_, d_620_v1_, iife48_(D2_DC6(_dafny.MultiSet(d_621_v2_)))])
        d_625_i1_: int
        d_625_i1_ = 0
        with _dafny.label("0"):
            while (self.f21 if (self).f18 else (d_622_v3_).isdisjoint(d_622_v3_)):
                with _dafny.c_label("0"):
                    if (d_625_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_625_i1_ = (d_625_i1_) + (1)
                    d_626_v6_: _dafny.Array
                    def lambda39_(d_627_i2_):
                        def iife52_():
                            coll14_ = _dafny.Set()
                            compr_14_: int
                            for compr_14_ in _dafny.IntegerRange(-686, 230):
                                d_628_v4_: int = compr_14_
                                if ((-686) <= (d_628_v4_)) and ((d_628_v4_) < (230)):
                                    coll14_ = coll14_.union(_dafny.Set([(d_628_v4_) - (self.f20)]))
                            return _dafny.Set(coll14_)
                        def iife53_():
                            coll15_ = _dafny.Map()
                            compr_15_: int
                            for compr_15_ in _dafny.IntegerRange(955, -330):
                                d_629_v5_: int = compr_15_
                                if ((955) <= (d_629_v5_)) and ((d_629_v5_) < (-330)):
                                    coll15_[(d_629_v5_) - (32)] = False
                            return _dafny.Map(coll15_)
                        return (_dafny.Map({_dafny.Map({977: self.f20}): _dafny.Map({(self).f19: self.f21})})) | (_dafny.Map({_dafny.Map({len(iife52_()
                        ): self.f20}): iife53_()
                        }))

                    init20_ = lambda39_
                    nw105_ = _dafny.Array(None, 29)
                    for i0_20_ in range(nw105_.length(0)):
                        nw105_[i0_20_] = init20_(i0_20_)
                    d_626_v6_ = nw105_
                    d_630_v7_: str
                    d_630_v7_ = _dafny.CodePoint('e')
                    d_631_v8_: _dafny.Map
                    d_631_v8_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_630_v7_])): len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yh"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yh")))), d_630_v7_))})
                    d_632_v9_: _dafny.Map
                    d_632_v9_ = _dafny.Map({(self).f19: self.f21})
                    d_633_v10_: _dafny.Map
                    d_633_v10_ = _dafny.Map({d_631_v8_: d_632_v9_})
                    index92_ = default__.safeIndex(543, (d_626_v6_).length(0))
                    (d_626_v6_)[index92_] = d_633_v10_
                    index93_ = default__.safeIndex(543, (d_626_v6_).length(0))
                    def iife54_():
                        coll16_ = _dafny.Map()
                        compr_16_: _dafny.Map
                        for compr_16_ in (_dafny.Set({d_631_v8_})).Elements:
                            d_634_v11_: _dafny.Map = compr_16_
                            if (d_634_v11_) in (_dafny.Set({d_631_v8_})):
                                coll16_[d_634_v11_] = d_632_v9_
                        return _dafny.Map(coll16_)
                    (d_626_v6_)[index93_] = ((d_633_v10_ if True else d_633_v10_) if self.f21 else iife54_()
                    )
                    if (d_630_v7_) in (default__.fm3(globalState)):
                        (globalState).f5 = p1
                        d_635_v12_: _dafny.Set
                        d_635_v12_ = _dafny.Set({(self).f19})
                        index94_ = default__.safeIndex(677, (p2).length(0))
                        (p2)[index94_] = (d_635_v12_).issubset(d_635_v12_)
                        d_636_v13_: D0
                        d_636_v13_ = D0_DC0((self).f18)
                        index95_ = default__.safeIndex(677, (p2).length(0))
                        (p2)[index95_] = (d_636_v13_).cf0
                        (globalState).f5 = (0) - (self.f20)
                        d_637_v14_: _dafny.Set
                        d_637_v14_ = _dafny.Set({self.f21, self.f21})
                        d_637_v14_ = (d_637_v14_).intersection((d_637_v14_).intersection(d_637_v14_))
                        d_638_v15_: _dafny.Map
                        d_638_v15_ = _dafny.Map({d_621_v2_: (self).f19})
                        d_638_v15_ = (d_638_v15_).set(d_621_v2_, (706) * ((0) - (len(p3))))
                    elif True:
                        (globalState).f13 = not((self).f18)
                        d_639_v16_: _dafny.Set
                        d_639_v16_ = _dafny.Set({self.f21, self.f21, (self).f18})
                        d_640_v17_: _dafny.Map
                        d_640_v17_ = _dafny.Map({811: d_639_v16_})
                        d_640_v17_ = d_640_v17_
                        d_641_v18_: D1
                        d_641_v18_ = D1_DC4((self).f18, False, self.f21)
                        (globalState).f13 = ((self).f18 if (d_641_v18_).cf7 else self.f21)
                        d_642_v19_: _dafny.Set
                        d_642_v19_ = _dafny.Set({self.f20, (0) - ((0) - (p1)), default__.fm0(p1, -806, p1, p1, globalState)})
                        d_643_v20_: C0
                        nw106_ = C0()
                        nw106_.ctor__(D2_DC6(p0), (self).f19, self.f21)
                        d_643_v20_ = nw106_
                        d_644_v21_: _dafny.Map
                        d_644_v21_ = _dafny.Map({d_642_v19_: d_643_v20_})
                        d_644_v21_ = (d_644_v21_).set((d_642_v19_) - (d_642_v19_), d_643_v20_)
                        d_645_v22_: _dafny.Seq
                        d_645_v22_ = _dafny.SeqWithoutIsStrInference([p3])
                        d_646_v23_: _dafny.Map
                        d_646_v23_ = _dafny.Map({p3: (d_645_v22_) < (d_645_v22_)})
                        d_646_v23_ = d_646_v23_
                    index96_ = default__.safeIndex(236, (d_618_v0_).length(0))
                    (d_618_v0_)[index96_] = default__.safeDivisionInt((self).f19, self.f20)
                    index97_ = default__.safeIndex(236, (d_618_v0_).length(0))
                    (d_618_v0_)[index97_] = (p1) * ((self).f19)
                    d_647_v24_: C0
                    nw107_ = C0()
                    nw107_.ctor__(d_620_v1_, (self).f19, (self).f18)
                    d_647_v24_ = nw107_
                    pass
            pass
        d_648_v25_: C0
        nw108_ = C0()
        nw108_.ctor__(D2_DC6(p0), p1, self.f21)
        d_648_v25_ = nw108_
        d_649_v26_: _dafny.Set
        d_649_v26_ = _dafny.Set({d_648_v25_, d_648_v25_})
        index98_ = default__.safeIndex(658, (d_618_v0_).length(0))
        (d_618_v0_)[index98_] = (len(d_649_v26_)) - (825)
        index99_ = default__.safeIndex(658, (d_618_v0_).length(0))
        (d_618_v0_)[index99_] = p1
        d_650_v27_: str
        d_650_v27_ = _dafny.CodePoint('y')
        d_651_v28_: _dafny.Set
        d_651_v28_ = _dafny.Set({d_650_v27_, d_650_v27_, d_650_v27_})
        d_652_v29_: _dafny.MultiSet
        d_652_v29_ = _dafny.MultiSet([d_651_v28_])
        d_653_v30_: _dafny.Map
        d_653_v30_ = _dafny.Map({(self).f18: (d_652_v29_).set(d_651_v28_, default__.abs(len(d_621_v2_)))})
        d_654_v31_: _dafny.Seq
        d_654_v31_ = _dafny.SeqWithoutIsStrInference([((d_653_v30_)[self.f21] if (self.f21) in (d_653_v30_) else d_652_v29_), _dafny.MultiSet([d_651_v28_, _dafny.Set({d_650_v27_})])])
        d_652_v29_ = (d_654_v31_)[default__.safeIndex((len(p3)) * (self.f20), len(d_654_v31_))]
        d_655_v32_: _dafny.Array
        nw109_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
        d_655_v32_ = nw109_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_655_v32_).length(0)):
            d_656_i3_: int = guard_loop_2_
            if (True) and (((0) <= (d_656_i3_)) and ((d_656_i3_) < ((d_655_v32_).length(0)))):
                (d_655_v32_)[(d_656_i3_)] = (p3) + (_dafny.SeqWithoutIsStrInference([d_650_v27_ for d_657_i4_ in range(default__.abs(-898))]))
        d_618_v0_ = d_618_v0_
        d_658_v33_: _dafny.Array
        def lambda40_(d_659_v25_):
            def lambda41_(d_660_i5_):
                return d_659_v25_.f22

            return lambda41_

        init21_ = lambda40_(d_648_v25_)
        nw110_ = _dafny.Array(None, 22)
        for i0_21_ in range(nw110_.length(0)):
            nw110_[i0_21_] = init21_(i0_21_)
        d_658_v33_ = nw110_
        r0 = d_658_v33_
        r1 = _dafny.Set({(self).f19, (d_648_v25_).f23})
        return r0, r1

