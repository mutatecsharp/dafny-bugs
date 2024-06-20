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
        if not(False):
            return not(True)
        elif True:
            return (761) == (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))))

    @staticmethod
    def fm1(globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvsslon"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iphflypt")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qc")))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([(0) - (len(_dafny.Set({False, True}))), -55, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suxyudnjw")))])) | (_dafny.MultiSet([len(_dafny.Set({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))])).cardinality}))]))) - ((_dafny.MultiSet([367])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([265]))))

    @staticmethod
    def fm3(p0, globalState):
        return (0) - (len(((_dafny.Map({False: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_0_i0_ in range(default__.abs(676))])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "byditj"))}))) | (_dafny.Map({(D1_DC4(_dafny.SeqWithoutIsStrInference([True, True]), False)).cf7: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))}))))

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return (_dafny.Map({True: True})) | (_dafny.Map({True: True}))

    @staticmethod
    def fm7(globalState):
        return _dafny.SeqWithoutIsStrInference([(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnuinsd")))) - (len(_dafny.SeqWithoutIsStrInference([True]))) for d_1_i0_ in range(default__.abs(619))])

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(604, 391):
                d_2_v0_: int = compr_0_
                if ((604) <= (d_2_v0_)) and ((d_2_v0_) < (391)):
                    coll0_ = coll0_.union(_dafny.Set([(d_2_v0_) + (262)]))
            return _dafny.Set(coll0_)
        return D1_DC4((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True])), (_dafny.Set({418, (_dafny.MultiSet([-592])).cardinality, len(_dafny.SeqWithoutIsStrInference([True, False, True]))})).issubset(_dafny.Set({len(iife0_()
), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([not(True)]))])).cardinality, 893})))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(-274, 697):
                d_3_v0_: int = compr_1_
                if ((-274) <= (d_3_v0_)) and ((d_3_v0_) < (697)):
                    coll1_[default__.safeModuloInt(d_3_v0_, 59)] = False
            return _dafny.Map(coll1_)
        return (iife1_()
        )

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(237, 872):
                d_4_v0_: int = compr_2_
                if ((237) <= (d_4_v0_)) and ((d_4_v0_) < (872)):
                    coll2_[(d_4_v0_) - (-833)] = True
            return _dafny.Map(coll2_)
        return ((_dafny.MultiSet([_dafny.Map({464: False})])).intersection(_dafny.MultiSet([_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqb")))})): not(True)})]))) - ((_dafny.MultiSet([_dafny.Map({965: False}), _dafny.Map({len(_dafny.Map({(0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True, False]))).cardinality): True})): True}), iife2_()
        ])) | (_dafny.MultiSet([_dafny.Map({-32: False})])))

    @staticmethod
    def fm11(p0, p1, globalState):
        return (D8_DC22(_dafny.Map({True: _dafny.CodePoint('j')}))).cf31

    @staticmethod
    def fm12(p0, p1, globalState):
        return D0_DC0(_dafny.CodePoint('m'))

    @staticmethod
    def fm13(p0, globalState):
        return _dafny.CodePoint('b')

    @staticmethod
    def fm14(p0, globalState):
        return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "buhcp"))})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqxqtv"))}))

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return _dafny.Map({len((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqow"))) for d_5_i0_ in range(default__.abs(-81))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.Map({617: 128})})) for d_6_i1_ in range(default__.abs(496))]))): -39})

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_7_v0_: _dafny.Seq
        d_7_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfhmvgywh"))
        hi0_ = ((_dafny.MultiSet([d_7_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yet")), d_7_v0_])).cardinality) * (len(default__.fm1(globalState)))
        for d_8_i0_ in range((0) - (p2), hi0_):
            if False:
                d_9_v1_: C0
                nw0_ = C0()
                nw0_.ctor__(p0)
                d_9_v1_ = nw0_
                d_10_v2_: _dafny.Seq
                d_10_v2_ = _dafny.SeqWithoutIsStrInference([p1])
                d_11_v3_: _dafny.MultiSet
                d_11_v3_ = _dafny.MultiSet([len(default__.fm11((d_9_v1_).f24, (0) - (d_8_i0_), globalState)), d_8_i0_, d_8_i0_])
                d_12_v4_: _dafny.Map
                d_12_v4_ = _dafny.Map({d_8_i0_: d_11_v3_})
                r1 = (d_9_v1_).fm4((D1_DC4(d_10_v2_, p0)).cf7, p0, (d_11_v3_ if default__.fm0(p2, (d_9_v1_).f24, _dafny.SeqWithoutIsStrInference([(D3_DC10(d_8_i0_, _dafny.CodePoint('m'))).cf17 for d_13_i1_ in range(default__.abs(337))]), globalState) else ((d_12_v4_)[d_8_i0_] if (d_8_i0_) in (d_12_v4_) else d_11_v3_)), default__.fm0(p2, p1, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_14_i2_ in range(default__.abs(532))]), globalState), globalState)
                d_15_v5_: _dafny.Seq
                d_15_v5_ = _dafny.SeqWithoutIsStrInference([d_9_v1_])
                d_16_v6_: _dafny.Seq
                d_16_v6_ = _dafny.SeqWithoutIsStrInference([d_15_v5_])
                d_17_v7_: _dafny.Set
                d_17_v7_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_9_v1_]), (d_16_v6_)[default__.safeIndex(d_8_i0_, len(d_16_v6_))], d_15_v5_, d_15_v5_, d_15_v5_})
                d_18_v8_: _dafny.Array
                nw1_ = _dafny.Array(None, 17)
                nw1_[int(0)] = d_17_v7_
                nw1_[int(1)] = d_17_v7_
                nw1_[int(2)] = d_17_v7_
                nw1_[int(3)] = (_dafny.Set({d_15_v5_})) - (d_17_v7_)
                nw1_[int(4)] = d_17_v7_
                nw1_[int(5)] = d_17_v7_
                nw1_[int(6)] = (d_17_v7_) - (d_17_v7_)
                nw1_[int(7)] = _dafny.Set({d_15_v5_})
                nw1_[int(8)] = _dafny.Set({d_15_v5_})
                nw1_[int(9)] = (d_17_v7_) | (d_17_v7_)
                nw1_[int(10)] = d_17_v7_
                nw1_[int(11)] = (d_17_v7_) | (d_17_v7_)
                nw1_[int(12)] = d_17_v7_
                nw1_[int(13)] = d_17_v7_
                nw1_[int(14)] = (_dafny.Set({d_15_v5_})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_9_v1_])}))
                nw1_[int(15)] = (d_17_v7_) | (d_17_v7_)
                nw1_[int(16)] = (d_17_v7_).intersection(d_17_v7_)
                d_18_v8_ = nw1_
                index0_ = default__.safeIndex(692, (d_18_v8_).length(0))
                (d_18_v8_)[index0_] = (d_17_v7_) | (_dafny.Set({d_15_v5_}))
                index1_ = default__.safeIndex(692, (d_18_v8_).length(0))
                (d_18_v8_)[index1_] = d_17_v7_
                (globalState).f23 = -419
                d_9_v1_ = d_9_v1_
            elif True:
                d_19_v9_: _dafny.Seq
                d_19_v9_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                r1 = (default__.safeModuloInt((d_19_v9_)[default__.safeIndex(p2, len(d_19_v9_))], d_8_i0_)) * (-624)
                d_20_v10_: _dafny.Map
                d_20_v10_ = _dafny.Map({p2: not (p0) or (False)})
                d_20_v10_ = d_20_v10_
                (globalState).f2 = p0
                d_21_v11_: str
                d_21_v11_ = _dafny.CodePoint('q')
                d_22_v12_: D0
                d_22_v12_ = D0_DC0(d_21_v11_)
                d_22_v12_ = (d_22_v12_ if p1 else default__.fm12(p1, True, globalState))
                d_23_v13_: _dafny.Seq
                d_23_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([76])])
                d_24_v15_: _dafny.Map
                d_24_v15_ = _dafny.Map({p1: d_8_i0_})
                d_25_v16_: _dafny.MultiSet
                d_25_v16_ = _dafny.MultiSet([default__.fm13(len(d_24_v15_), globalState), d_21_v11_])
                d_26_v17_: _dafny.Map
                def iife3_():
                    coll3_ = _dafny.Map()
                    compr_3_: str
                    for compr_3_ in (d_25_v16_).Elements:
                        d_27_v14_: str = compr_3_
                        if (d_27_v14_) in (d_25_v16_):
                            coll3_[d_27_v14_] = d_8_i0_
                    return _dafny.Map(coll3_)
                d_26_v17_ = _dafny.Map({iife3_()
                : p1})
                rhs0_ = True
                rhs1_ = ((d_23_v13_)[default__.safeIndex(p2, len(d_23_v13_))]) + ((_dafny.SeqWithoutIsStrInference([d_8_i0_, 701, len(d_26_v17_)])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_8_i0_, 701, len(d_26_v17_)]))), p2))
                lhs0_ = globalState
                lhs0_.f18 = rhs0_
                d_19_v9_ = rhs1_
            d_28_v18_: C0
            nw2_ = C0()
            nw2_.ctor__(p0)
            d_28_v18_ = nw2_
            if p0:
                (globalState).f18 = p1
                d_7_v0_ = (default__.fm1(globalState)) + (d_7_v0_)
                nw3_ = C0()
                nw3_.ctor__((p1) or (p0))
                d_28_v18_ = nw3_
                d_29_v19_: _dafny.Array
                nw4_ = _dafny.Array(int(0), 20)
                d_29_v19_ = nw4_
                index2_ = default__.safeIndex(3, (d_29_v19_).length(0))
                (d_29_v19_)[index2_] = (0) - ((d_8_i0_) + (d_8_i0_))
                d_30_v20_: _dafny.Map
                d_30_v20_ = _dafny.Map({(0) - (d_8_i0_): p1})
                d_31_v21_: _dafny.Seq
                d_31_v21_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(d_7_v0_)])), len(d_30_v20_)])
                d_32_v22_: D4
                d_32_v22_ = D4_DC13(d_31_v21_)
                d_33_v23_: _dafny.Array
                nw5_ = _dafny.Array(None, 5)
                nw5_[int(0)] = d_31_v21_
                nw5_[int(1)] = d_31_v21_
                nw5_[int(2)] = (d_32_v22_).cf20
                nw5_[int(3)] = d_31_v21_
                nw5_[int(4)] = _dafny.SeqWithoutIsStrInference([p2 for d_34_i3_ in range(default__.abs(799))])
                d_33_v23_ = nw5_
                index3_ = default__.safeIndex(403, (d_33_v23_).length(0))
                (d_33_v23_)[index3_] = _dafny.SeqWithoutIsStrInference([d_8_i0_])
                d_35_v24_: _dafny.Seq
                d_35_v24_ = _dafny.SeqWithoutIsStrInference([(d_28_v18_).f24])
                d_36_v25_: _dafny.Set
                d_36_v25_ = _dafny.Set({d_35_v24_})
                index4_ = default__.safeIndex(3, (d_29_v19_).length(0))
                index5_ = default__.safeIndex(403, (d_33_v23_).length(0))
                rhs2_ = len((d_36_v25_).intersection(d_36_v25_))
                rhs3_ = (d_31_v21_) + (d_31_v21_)
                lhs1_ = d_29_v19_
                lhs2_ = default__.safeIndex(3, (d_29_v19_).length(0))
                lhs3_ = d_33_v23_
                lhs4_ = default__.safeIndex(403, (d_33_v23_).length(0))
                lhs1_[lhs2_] = rhs2_
                lhs3_[lhs4_] = rhs3_
                index6_ = default__.safeIndex(3, (d_29_v19_).length(0))
                (d_29_v19_)[index6_] = (d_29_v19_)[default__.safeIndex(3, (d_29_v19_).length(0))]
            elif True:
                d_37_v26_: _dafny.MultiSet
                d_37_v26_ = _dafny.MultiSet([p1, p1])
                r0 = (p2) * (((d_37_v26_).cardinality) * (d_8_i0_))
                d_38_v27_: _dafny.Array
                nw6_ = _dafny.Array(int(0), 12)
                d_38_v27_ = nw6_
                index7_ = default__.safeIndex(365, (d_38_v27_).length(0))
                (d_38_v27_)[index7_] = p2
                d_39_v28_: _dafny.Map
                d_39_v28_ = _dafny.Map({(d_8_i0_) != (d_8_i0_): d_7_v0_})
                d_40_v29_: _dafny.Map
                d_40_v29_ = _dafny.Map({p1: 648})
                d_41_v30_: _dafny.Seq
                d_41_v30_ = _dafny.SeqWithoutIsStrInference([len(d_40_v29_), p2])
                d_42_v31_: str
                d_42_v31_ = _dafny.CodePoint('t')
                index8_ = default__.safeIndex(365, (d_38_v27_).length(0))
                rhs4_ = (d_28_v18_).fm4((d_28_v18_).f24, not((d_41_v30_) <= (d_41_v30_)), _dafny.MultiSet([d_8_i0_, p2, d_8_i0_, d_8_i0_]), (d_28_v18_).f24, globalState)
                rhs5_ = p0
                rhs6_ = p0
                rhs7_ = default__.fm14(d_42_v31_, globalState)
                lhs5_ = d_38_v27_
                lhs6_ = default__.safeIndex(365, (d_38_v27_).length(0))
                lhs7_ = globalState
                lhs8_ = globalState
                lhs5_[lhs6_] = rhs4_
                lhs7_.f2 = rhs5_
                lhs8_.f2 = rhs6_
                d_39_v28_ = rhs7_
                d_43_v32_: _dafny.Map
                d_43_v32_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_8_i0_ for d_44_i4_ in range(default__.abs(-384))])): p2})
                d_45_v33_: _dafny.Map
                d_45_v33_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_28_v18_).f24, p0]): not(True)})
                d_46_v34_: _dafny.Map
                d_46_v34_ = _dafny.Map({len(d_43_v32_): not(((d_45_v33_)[_dafny.SeqWithoutIsStrInference([p1])] if (_dafny.SeqWithoutIsStrInference([p1])) in (d_45_v33_) else (d_28_v18_).f24))})
                d_47_v35_: _dafny.MultiSet
                d_47_v35_ = _dafny.MultiSet([(d_46_v34_) | (_dafny.Map({p2: p1}))])
                index9_ = default__.safeIndex(365, (d_38_v27_).length(0))
                (d_38_v27_)[index9_] = ((d_47_v35_)[default__.fm9((0) - ((d_37_v26_).cardinality), (d_28_v18_).f24, p2, p1, globalState)] if (default__.fm9((0) - ((d_37_v26_).cardinality), (d_28_v18_).f24, p2, p1, globalState)) in (d_47_v35_) else default__.safeDivisionInt((0) - (d_8_i0_), d_8_i0_))
                d_48_v36_: D3
                d_48_v36_ = D3_DC10(p2, d_42_v31_)
                d_49_v37_: _dafny.Map
                d_49_v37_ = _dafny.Map({d_28_v18_: d_48_v36_})
                d_50_v38_: _dafny.Array
                nw7_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_50_v38_ = nw7_
                d_51_v39_: _dafny.Array
                nw8_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_51_v39_ = nw8_
                index10_ = default__.safeIndex(326, (d_50_v38_).length(0))
                (d_50_v38_)[index10_] = d_51_v39_
                d_52_v40_: D5
                d_52_v40_ = D5_DC15(d_51_v39_)
                d_53_v41_: _dafny.Seq
                d_53_v41_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_42_v31_ for d_54_i5_ in range(default__.abs(422))])])
                index11_ = default__.safeIndex(326, (d_50_v38_).length(0))
                rhs8_ = ((d_49_v37_) | (d_49_v37_)).set(d_28_v18_, d_48_v36_)
                rhs9_ = (d_52_v40_).cf22
                rhs10_ = (d_53_v41_)[default__.safeIndex((d_38_v27_)[default__.safeIndex(365, (d_38_v27_).length(0))], len(d_53_v41_))]
                lhs9_ = d_50_v38_
                lhs10_ = default__.safeIndex(326, (d_50_v38_).length(0))
                d_49_v37_ = rhs8_
                lhs9_[lhs10_] = rhs9_
                d_7_v0_ = rhs10_
                d_28_v18_ = d_28_v18_
            d_55_v42_: D3
            d_55_v42_ = D3_DC11((d_28_v18_).f24)
            d_56_v43_: C0
            nw9_ = C0()
            nw9_.ctor__((422) < ((_dafny.MultiSet([_dafny.Set({(d_55_v42_).cf18, p0})])).cardinality))
            d_56_v43_ = nw9_
        (globalState).f18 = default__.fm0(p2, p1, (d_7_v0_) + (d_7_v0_), globalState)
        hi1_ = p2
        for d_57_i6_ in range(p2, hi1_):
            d_58_v44_: str
            d_58_v44_ = _dafny.CodePoint('h')
            d_59_v45_: C0
            nw10_ = C0()
            nw10_.ctor__(p1)
            d_59_v45_ = nw10_
            d_60_v46_: _dafny.Map
            d_60_v46_ = _dafny.Map({d_59_v45_: (d_59_v45_).f24})
            d_61_v47_: _dafny.Map
            d_61_v47_ = _dafny.Map({d_57_i6_: (_dafny.Map({len(d_60_v46_): (d_59_v45_).f24})).set(d_57_i6_, p0)})
            d_62_v48_: D6
            d_62_v48_ = D6_DC18(d_61_v47_)
            d_63_v49_: _dafny.Map
            d_63_v49_ = _dafny.Map({(d_60_v46_) | (d_60_v46_): _dafny.SeqWithoutIsStrInference([d_58_v44_ for d_64_i7_ in range(default__.abs(-676))])})
            rhs11_ = d_58_v44_
            rhs12_ = (d_62_v48_).cf26
            rhs13_ = ((d_63_v49_)[_dafny.Map({d_59_v45_: (d_59_v45_).f24})] if (_dafny.Map({d_59_v45_: (d_59_v45_).f24})) in (d_63_v49_) else (d_7_v0_) + (d_7_v0_))
            rhs14_ = d_59_v45_
            rhs15_ = (d_57_i6_) != (153)
            lhs11_ = globalState
            d_58_v44_ = rhs11_
            d_61_v47_ = rhs12_
            d_7_v0_ = rhs13_
            d_59_v45_ = rhs14_
            lhs11_.f18 = rhs15_
            r1 = d_57_i6_
            d_65_v50_: _dafny.Seq
            d_65_v50_ = _dafny.SeqWithoutIsStrInference([p1])
            d_66_v51_: D1
            d_66_v51_ = D1_DC4(d_65_v50_, not(p0))
            d_67_v52_: _dafny.Seq
            d_67_v52_ = _dafny.SeqWithoutIsStrInference([d_7_v0_])
            d_68_v54_: _dafny.Seq
            d_68_v54_ = _dafny.SeqWithoutIsStrInference([len(d_65_v50_), p2])
            d_69_v55_: _dafny.Map
            d_69_v55_ = _dafny.Map({(d_59_v45_).fm4(p0, (d_59_v45_).f24, _dafny.MultiSet(d_68_v54_), p1, globalState): d_57_i6_})
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in (default__.fm15(d_57_i6_, p0, p1, globalState)).keys.Elements:
                    d_70_v53_: int = compr_4_
                    if (d_70_v53_) in (default__.fm15(d_57_i6_, p0, p1, globalState)):
                        coll4_[(d_70_v53_) * (d_57_i6_)] = p2
                return _dafny.Map(coll4_)
            rhs16_ = p0
            rhs17_ = default__.safeModuloInt(default__.safeModuloInt(len(d_67_v52_), d_57_i6_), len((iife4_()
             if p0 else d_69_v55_)))
            rhs18_ = d_66_v51_
            lhs12_ = globalState
            lhs13_ = globalState
            lhs12_.f18 = rhs16_
            lhs13_.f23 = rhs17_
            d_66_v51_ = rhs18_
            r1 = (0) - ((d_57_i6_) - (p2))
        d_71_i8_: int
        d_71_i8_ = 0
        with _dafny.label("0"):
            while default__.fm0((p2) + (p2), p1, d_7_v0_, globalState):
                with _dafny.c_label("0"):
                    if (d_71_i8_) >= (100):
                        raise _dafny.Break("0")
                    d_71_i8_ = (d_71_i8_) + (1)
                    (globalState).f18 = p1
                    d_72_v56_: D6
                    d_72_v56_ = D6_DC20()
                    d_72_v56_ = d_72_v56_
                    d_73_v57_: _dafny.Array
                    nw11_ = _dafny.Array(_dafny.Array(None, 0), 25)
                    d_73_v57_ = nw11_
                    d_74_v58_: _dafny.Array
                    nw12_ = _dafny.Array(False, 9)
                    d_74_v58_ = nw12_
                    index12_ = default__.safeIndex(686, (d_73_v57_).length(0))
                    (d_73_v57_)[index12_] = d_74_v58_
                    index13_ = default__.safeIndex(686, (d_73_v57_).length(0))
                    nw13_ = _dafny.Array(False, 23)
                    (d_73_v57_)[index13_] = nw13_
                    (globalState).f2 = p0
                    pass
            pass
        d_75_v59_: _dafny.Array
        nw14_ = _dafny.Array(_dafny.Array(None, 0), 5)
        d_75_v59_ = nw14_
        d_76_v60_: str
        d_76_v60_ = _dafny.CodePoint('v')
        d_77_v61_: _dafny.Array
        nw15_ = _dafny.Array(None, 9)
        nw15_[int(0)] = d_7_v0_
        nw15_[int(1)] = d_7_v0_
        nw15_[int(2)] = d_7_v0_
        nw15_[int(3)] = d_7_v0_
        nw15_[int(4)] = d_7_v0_
        nw15_[int(5)] = d_7_v0_
        nw15_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
        nw15_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krdlpcyb"))
        nw15_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sj"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sj")))), d_76_v60_)
        d_77_v61_ = nw15_
        index14_ = default__.safeIndex(679, (d_75_v59_).length(0))
        (d_75_v59_)[index14_] = d_77_v61_
        index15_ = default__.safeIndex(679, (d_75_v59_).length(0))
        (d_75_v59_)[index15_] = d_77_v61_
        (globalState).f2 = p1
        d_78_v62_: _dafny.Array
        def lambda0_(d_79_p0_):
            def lambda1_(d_80_i9_):
                return d_79_p0_

            return lambda1_

        init0_ = lambda0_(p0)
        nw16_ = _dafny.Array(None, 9)
        for i0_0_ in range(nw16_.length(0)):
            nw16_[i0_0_] = init0_(i0_0_)
        d_78_v62_ = nw16_
        d_81_v63_: _dafny.MultiSet
        d_81_v63_ = _dafny.MultiSet([d_78_v62_, d_78_v62_])
        r0 = ((d_81_v63_) - ((d_81_v63_).set(d_78_v62_, default__.abs(177)))).cardinality
        d_82_v64_: _dafny.Set
        d_82_v64_ = _dafny.Set({d_76_v60_, d_76_v60_})
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: str
            for compr_5_ in (d_82_v64_).Elements:
                d_83_v65_: str = compr_5_
                if (d_83_v65_) in (d_82_v64_):
                    coll5_ = coll5_.union(_dafny.Set([d_83_v65_]))
            return _dafny.Set(coll5_)
        r1 = (p2) + ((0) - (len((d_82_v64_) - (iife5_()
        ))))
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_84_v0_: _dafny.Seq
        d_84_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw"))
        d_85_v1_: _dafny.MultiSet
        d_85_v1_ = _dafny.MultiSet([d_84_v0_, d_84_v0_, d_84_v0_, d_84_v0_, d_84_v0_])
        d_86_v2_: bool
        d_86_v2_ = True
        d_87_v3_: int
        d_87_v3_ = 56
        d_88_v4_: _dafny.Map
        d_88_v4_ = _dafny.Map({d_86_v2_: d_87_v3_})
        d_89_v5_: _dafny.Map
        d_89_v5_ = _dafny.Map({d_88_v4_: d_87_v3_})
        d_90_v6_: _dafny.Array
        nw17_ = _dafny.Array(_dafny.Array(None, 0), 16)
        d_90_v6_ = nw17_
        d_91_v7_: _dafny.Map
        d_91_v7_ = _dafny.Map({d_86_v2_: d_90_v6_})
        d_92_v8_: _dafny.Array
        nw18_ = _dafny.Array(int(0), 28)
        d_92_v8_ = nw18_
        d_93_v9_: str
        d_93_v9_ = _dafny.CodePoint('n')
        d_94_v10_: _dafny.Map
        d_94_v10_ = _dafny.Map({d_86_v2_: d_93_v9_})
        d_95_v11_: _dafny.Map
        d_95_v11_ = _dafny.Map({d_87_v3_: d_86_v2_})
        d_96_v12_: _dafny.Array
        nw19_ = _dafny.Array(None, 16)
        nw19_[int(0)] = d_93_v9_
        nw19_[int(1)] = d_93_v9_
        nw19_[int(2)] = d_93_v9_
        nw19_[int(3)] = d_93_v9_
        nw19_[int(4)] = _dafny.CodePoint('w')
        nw19_[int(5)] = d_93_v9_
        nw19_[int(6)] = d_93_v9_
        nw19_[int(7)] = d_93_v9_
        nw19_[int(8)] = d_93_v9_
        nw19_[int(9)] = d_93_v9_
        nw19_[int(10)] = d_93_v9_
        nw19_[int(11)] = d_93_v9_
        nw19_[int(12)] = ((d_94_v10_)[((d_95_v11_)[d_87_v3_] if (d_87_v3_) in (d_95_v11_) else d_86_v2_)] if (((d_95_v11_)[d_87_v3_] if (d_87_v3_) in (d_95_v11_) else d_86_v2_)) in (d_94_v10_) else d_93_v9_)
        nw19_[int(13)] = d_93_v9_
        nw19_[int(14)] = d_93_v9_
        nw19_[int(15)] = _dafny.CodePoint('g')
        d_96_v12_ = nw19_
        d_97_v13_: _dafny.Seq
        d_97_v13_ = _dafny.SeqWithoutIsStrInference([d_87_v3_, d_87_v3_])
        d_98_globalState_: GlobalState
        nw20_ = GlobalState()
        nw20_.ctor__(580, d_85_v1_, False, d_89_v5_, 548, True, True, 376, 87, True, ((d_91_v7_)[False] if (False) in (d_91_v7_) else d_90_v6_), _dafny.CodePoint('w'), -257, d_92_v8_, (d_84_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_99_i0_ in range(default__.abs(-26))])), d_96_v12_, -272, d_96_v12_, False, True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dblg")), d_97_v13_, d_97_v13_, -128)
        d_98_globalState_ = nw20_
        d_100_v14_: int
        d_101_v15_: int
        out0_: int
        out1_: int
        out0_, out1_ = default__.m0(d_86_v2_, d_86_v2_, d_87_v3_, d_98_globalState_)
        d_100_v14_ = out0_
        d_101_v15_ = out1_
        d_102_v16_: _dafny.MultiSet
        d_102_v16_ = _dafny.MultiSet([d_86_v2_, d_86_v2_])
        if (d_102_v16_).ispropersubset(d_102_v16_):
            d_103_v17_: int
            d_104_v18_: int
            out2_: int
            out3_: int
            out2_, out3_ = default__.m0(d_86_v2_, (d_100_v14_) <= (d_87_v3_), (d_101_v15_) - (d_87_v3_), d_98_globalState_)
            d_103_v17_ = out2_
            d_104_v18_ = out3_
            if d_86_v2_:
                d_105_v19_: _dafny.Array
                def lambda2_(d_106_i1_):
                    return True

                init1_ = lambda2_
                nw21_ = _dafny.Array(None, 21)
                for i0_1_ in range(nw21_.length(0)):
                    nw21_[i0_1_] = init1_(i0_1_)
                d_105_v19_ = nw21_
                index16_ = default__.safeIndex(533, (d_105_v19_).length(0))
                (d_105_v19_)[index16_] = d_86_v2_
                d_107_v20_: _dafny.Seq
                d_107_v20_ = _dafny.SeqWithoutIsStrInference([d_86_v2_, d_86_v2_])
                index17_ = default__.safeIndex(533, (d_105_v19_).length(0))
                (d_105_v19_)[index17_] = (d_107_v20_)[default__.safeIndex(d_104_v18_, len(d_107_v20_))]
                d_108_v21_: _dafny.Map
                d_108_v21_ = _dafny.Map({d_86_v2_: (d_105_v19_)[default__.safeIndex(533, (d_105_v19_).length(0))]})
                d_109_v22_: _dafny.Array
                nw22_ = _dafny.Array(None, 12)
                nw22_[int(0)] = d_108_v21_
                nw22_[int(1)] = (d_108_v21_).set(d_86_v2_, (d_105_v19_)[default__.safeIndex(533, (d_105_v19_).length(0))])
                nw22_[int(2)] = d_108_v21_
                nw22_[int(3)] = d_108_v21_
                nw22_[int(4)] = (d_108_v21_).set((d_105_v19_)[default__.safeIndex(533, (d_105_v19_).length(0))], d_86_v2_)
                nw22_[int(5)] = d_108_v21_
                nw22_[int(6)] = d_108_v21_
                nw22_[int(7)] = d_108_v21_
                nw22_[int(8)] = (d_108_v21_) | (d_108_v21_)
                nw22_[int(9)] = d_108_v21_
                nw22_[int(10)] = (d_108_v21_) | (d_108_v21_)
                nw22_[int(11)] = d_108_v21_
                d_109_v22_ = nw22_
                d_109_v22_ = (d_109_v22_ if (d_105_v19_)[default__.safeIndex(533, (d_105_v19_).length(0))] else (d_109_v22_ if d_86_v2_ else d_109_v22_))
                d_95_v11_ = (d_95_v11_).set(d_100_v14_, d_86_v2_)
                d_110_v23_: D0
                d_110_v23_ = D0_DC0(d_93_v9_)
                index18_ = default__.safeIndex(533, (d_105_v19_).length(0))
                (d_105_v19_)[index18_] = default__.fm0((0) - (default__.safeDivisionInt(d_87_v3_, d_87_v3_)), not(((default__.fm1(d_98_globalState_)).set(default__.safeIndex(len(d_84_v0_), len(default__.fm1(d_98_globalState_))), (d_110_v23_).cf0)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))), d_84_v0_, d_98_globalState_)
                d_111_v24_: int
                d_112_v25_: int
                out4_: int
                out5_: int
                out4_, out5_ = default__.m0((d_105_v19_)[default__.safeIndex(533, (d_105_v19_).length(0))], d_86_v2_, d_104_v18_, d_98_globalState_)
                d_111_v24_ = out4_
                d_112_v25_ = out5_
            elif True:
                d_100_v14_ = d_87_v3_
                d_113_v26_: D1
                d_113_v26_ = D1_DC2(d_86_v2_)
                (d_98_globalState_).f18 = default__.fm0(d_103_v17_, (d_113_v26_).cf2, d_84_v0_, d_98_globalState_)
                (d_98_globalState_).f23 = d_104_v18_
                d_114_v27_: _dafny.Seq
                d_114_v27_ = _dafny.SeqWithoutIsStrInference([d_86_v2_, d_86_v2_])
                d_115_v28_: _dafny.Map
                d_115_v28_ = _dafny.Map({d_86_v2_: not(False)})
                d_116_v29_: _dafny.Array
                nw23_ = _dafny.Array(None, 17)
                nw23_[int(0)] = not (d_86_v2_) or (d_86_v2_)
                nw23_[int(1)] = False
                nw23_[int(2)] = d_86_v2_
                nw23_[int(3)] = d_86_v2_
                nw23_[int(4)] = d_86_v2_
                nw23_[int(5)] = not((d_84_v0_) < (d_84_v0_))
                nw23_[int(6)] = (d_100_v14_) == (-375)
                nw23_[int(7)] = d_86_v2_
                nw23_[int(8)] = (d_100_v14_) != (d_100_v14_)
                nw23_[int(9)] = (d_86_v2_) and (d_86_v2_)
                nw23_[int(10)] = True
                nw23_[int(11)] = default__.fm0(d_100_v14_, d_86_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")), d_98_globalState_)
                nw23_[int(12)] = d_86_v2_
                nw23_[int(13)] = d_86_v2_
                nw23_[int(14)] = d_86_v2_
                nw23_[int(15)] = True
                nw23_[int(16)] = (d_114_v27_)[default__.safeIndex(len(d_115_v28_), len(d_114_v27_))]
                d_116_v29_ = nw23_
                rhs19_ = _dafny.SeqWithoutIsStrInference([d_93_v9_ for d_117_i2_ in range(default__.abs(952))])
                rhs20_ = d_84_v0_
                rhs21_ = d_116_v29_
                rhs22_ = ((0) - (d_87_v3_)) != (d_104_v18_)
                rhs23_ = (len((d_84_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nboyosu"))))) + ((0) - (d_104_v18_))
                lhs14_ = d_98_globalState_
                lhs15_ = d_98_globalState_
                d_84_v0_ = rhs19_
                lhs14_.f20 = rhs20_
                d_116_v29_ = rhs21_
                lhs15_.f2 = rhs22_
                d_100_v14_ = rhs23_
                index19_ = default__.safeIndex(733, (d_116_v29_).length(0))
                (d_116_v29_)[index19_] = not(d_86_v2_)
                index20_ = default__.safeIndex(733, (d_116_v29_).length(0))
                (d_116_v29_)[index20_] = d_86_v2_
            d_118_v30_: _dafny.Set
            d_118_v30_ = _dafny.Set({d_103_v17_})
            d_119_v31_: _dafny.MultiSet
            d_119_v31_ = _dafny.MultiSet([d_104_v18_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfrvenm"))), d_87_v3_])
            d_120_v32_: int
            d_121_v33_: int
            out6_: int
            out7_: int
            out6_, out7_ = default__.m0(d_86_v2_, (default__.fm2(d_118_v30_, d_101_v15_, d_100_v14_, 916, d_98_globalState_)).isdisjoint(d_119_v31_), default__.fm3(d_101_v15_, d_98_globalState_), d_98_globalState_)
            d_120_v32_ = out6_
            d_121_v33_ = out7_
            d_104_v18_ = d_101_v15_
            if d_86_v2_:
                d_122_v34_: int
                d_123_v35_: int
                out8_: int
                out9_: int
                out8_, out9_ = default__.m0(False, d_86_v2_, default__.safeModuloInt(d_104_v18_, d_100_v14_), d_98_globalState_)
                d_122_v34_ = out8_
                d_123_v35_ = out9_
                d_124_v36_: D1
                d_124_v36_ = D1_DC2(d_86_v2_)
                d_125_v37_: _dafny.Array
                nw24_ = _dafny.Array(False, 26)
                d_125_v37_ = nw24_
                d_126_v38_: _dafny.Map
                d_126_v38_ = _dafny.Map({d_124_v36_: d_125_v37_})
                rhs24_ = d_86_v2_
                rhs25_ = d_126_v38_
                lhs16_ = d_98_globalState_
                lhs16_.f18 = rhs24_
                d_126_v38_ = rhs25_
                d_127_v39_: C0
                nw25_ = C0()
                nw25_.ctor__((not(d_86_v2_) if d_86_v2_ else not(d_86_v2_)))
                d_127_v39_ = nw25_
                d_128_v40_: C0
                nw26_ = C0()
                nw26_.ctor__((d_127_v39_).f24)
                d_128_v40_ = nw26_
                d_128_v40_ = d_127_v39_
                index21_ = default__.safeIndex(206, (d_125_v37_).length(0))
                (d_125_v37_)[index21_] = (d_127_v39_).f24
                index22_ = default__.safeIndex(206, (d_125_v37_).length(0))
                (d_125_v37_)[index22_] = ((d_102_v16_) - (d_102_v16_)).isdisjoint((d_102_v16_) - (d_102_v16_))
            elif True:
                d_129_v41_: _dafny.Set
                d_129_v41_ = _dafny.Set({d_93_v9_, d_93_v9_})
                d_130_v42_: _dafny.Map
                d_130_v42_ = _dafny.Map({d_84_v0_: d_129_v41_})
                d_131_v43_: _dafny.Seq
                d_131_v43_ = _dafny.SeqWithoutIsStrInference([((d_130_v42_)[d_84_v0_] if (d_84_v0_) in (d_130_v42_) else d_129_v41_)])
                d_132_v44_: int
                d_133_v45_: int
                out10_: int
                out11_: int
                out10_, out11_ = default__.m0(default__.fm0(d_87_v3_, d_86_v2_, d_84_v0_, d_98_globalState_), (d_131_v43_) <= (d_131_v43_), 495, d_98_globalState_)
                d_132_v44_ = out10_
                d_133_v45_ = out11_
                index23_ = default__.safeIndex(246, (d_92_v8_).length(0))
                (d_92_v8_)[index23_] = (d_100_v14_) * (d_120_v32_)
                index24_ = default__.safeIndex(246, (d_92_v8_).length(0))
                (d_92_v8_)[index24_] = d_133_v45_
                d_134_v46_: _dafny.Array
                def lambda3_(d_135_v2_):
                    def lambda4_(d_136_i3_):
                        return d_135_v2_

                    return lambda4_

                init2_ = lambda3_(d_86_v2_)
                nw27_ = _dafny.Array(None, 11)
                for i0_2_ in range(nw27_.length(0)):
                    nw27_[i0_2_] = init2_(i0_2_)
                d_134_v46_ = nw27_
                index25_ = default__.safeIndex(996, (d_134_v46_).length(0))
                (d_134_v46_)[index25_] = d_86_v2_
                d_137_v47_: D2
                d_137_v47_ = D2_DC6(d_84_v0_)
                d_138_v48_: _dafny.Array
                nw28_ = _dafny.Array(None, 18)
                nw28_[int(0)] = d_84_v0_
                nw28_[int(1)] = d_84_v0_
                nw28_[int(2)] = d_84_v0_
                nw28_[int(3)] = d_84_v0_
                nw28_[int(4)] = d_84_v0_
                nw28_[int(5)] = d_84_v0_
                nw28_[int(6)] = _dafny.SeqWithoutIsStrInference([d_93_v9_ for d_139_i4_ in range(default__.abs(-932))])
                nw28_[int(7)] = d_84_v0_
                nw28_[int(8)] = d_84_v0_
                nw28_[int(9)] = (d_137_v47_).cf9
                nw28_[int(10)] = d_84_v0_
                nw28_[int(11)] = d_84_v0_
                nw28_[int(12)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_140_i5_ in range(default__.abs(-501))])
                nw28_[int(13)] = d_84_v0_
                nw28_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awa"))
                nw28_[int(15)] = d_84_v0_
                nw28_[int(16)] = d_84_v0_
                nw28_[int(17)] = d_84_v0_
                d_138_v48_ = nw28_
                d_141_v49_: D1
                d_141_v49_ = D1_DC3((D1_DC3(d_138_v48_, 779, d_86_v2_)).cf3, d_100_v14_, d_86_v2_)
                index26_ = default__.safeIndex(996, (d_134_v46_).length(0))
                (d_134_v46_)[index26_] = not((d_141_v49_).cf5)
                d_142_v50_: C0
                nw29_ = C0()
                nw29_.ctor__(default__.fm0(len(d_118_v30_), (d_134_v46_)[default__.safeIndex(996, (d_134_v46_).length(0))], (d_84_v0_).set(default__.safeIndex((0) - (d_121_v33_), len(d_84_v0_)), d_93_v9_), d_98_globalState_))
                d_142_v50_ = nw29_
                d_120_v32_ = (d_92_v8_)[default__.safeIndex(246, (d_92_v8_).length(0))]
        elif True:
            d_86_v2_ = d_86_v2_
            d_143_v51_: _dafny.Array
            nw30_ = _dafny.Array(False, 2)
            d_143_v51_ = nw30_
            d_144_v52_: _dafny.Seq
            d_144_v52_ = _dafny.SeqWithoutIsStrInference([d_86_v2_])
            index27_ = default__.safeIndex(115, (d_143_v51_).length(0))
            (d_143_v51_)[index27_] = (d_144_v52_) != (d_144_v52_)
            index28_ = default__.safeIndex(115, (d_143_v51_).length(0))
            (d_143_v51_)[index28_] = (d_100_v14_) <= (default__.fm3(d_101_v15_, d_98_globalState_))
            d_145_v53_: _dafny.Array
            nw31_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
            d_145_v53_ = nw31_
            d_146_v54_: _dafny.MultiSet
            d_146_v54_ = _dafny.MultiSet([d_100_v14_, d_100_v14_])
            d_147_v55_: D1
            d_147_v55_ = D1_DC3(d_145_v53_, (d_146_v54_).cardinality, default__.fm0(d_87_v3_, (d_143_v51_)[default__.safeIndex(115, (d_143_v51_).length(0))], d_84_v0_, d_98_globalState_))
            d_148_v56_: int
            d_149_v57_: int
            out12_: int
            out13_: int
            out12_, out13_ = default__.m0(((d_102_v16_).cardinality) >= ((d_147_v55_).cf4), (d_100_v14_) == (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efngv")))), (d_97_v13_)[default__.safeIndex(d_101_v15_, len(d_97_v13_))], d_98_globalState_)
            d_148_v56_ = out12_
            d_149_v57_ = out13_
            d_150_v58_: int
            d_151_v59_: int
            out14_: int
            out15_: int
            out14_, out15_ = default__.m0(d_86_v2_, (d_143_v51_)[default__.safeIndex(115, (d_143_v51_).length(0))], 210, d_98_globalState_)
            d_150_v58_ = out14_
            d_151_v59_ = out15_
            index29_ = default__.safeIndex(115, (d_143_v51_).length(0))
            index30_ = default__.safeIndex(115, (d_143_v51_).length(0))
            rhs26_ = d_143_v51_
            rhs27_ = ((d_101_v15_) >= (d_151_v59_)) and ((_dafny.MultiSet([d_149_v57_, d_87_v3_])).issubset(d_146_v54_))
            rhs28_ = (d_143_v51_)[default__.safeIndex(115, (d_143_v51_).length(0))]
            rhs29_ = (d_148_v56_) >= ((0) - (d_149_v57_))
            rhs30_ = (0) - ((d_151_v59_) - (d_148_v56_))
            lhs17_ = d_143_v51_
            lhs18_ = default__.safeIndex(115, (d_143_v51_).length(0))
            lhs19_ = d_143_v51_
            lhs20_ = default__.safeIndex(115, (d_143_v51_).length(0))
            lhs21_ = d_98_globalState_
            lhs22_ = d_98_globalState_
            d_143_v51_ = rhs26_
            lhs17_[lhs18_] = rhs27_
            lhs19_[lhs20_] = rhs28_
            lhs21_.f2 = rhs29_
            lhs22_.f23 = rhs30_
        d_152_v60_: _dafny.Array
        nw32_ = _dafny.Array(False, 28)
        d_152_v60_ = nw32_
        d_153_v61_: _dafny.Map
        d_153_v61_ = _dafny.Map({d_152_v60_: d_84_v0_})
        (d_98_globalState_).f2 = default__.fm0(d_100_v14_, d_86_v2_, ((d_153_v61_)[d_152_v60_] if (d_152_v60_) in (d_153_v61_) else d_84_v0_), d_98_globalState_)
        d_154_v62_: C0
        nw33_ = C0()
        nw33_.ctor__(not(d_86_v2_))
        d_154_v62_ = nw33_
        d_155_v63_: _dafny.Map
        d_155_v63_ = _dafny.Map({d_86_v2_: d_154_v62_})
        d_156_v64_: _dafny.Map
        d_156_v64_ = _dafny.Map({d_155_v63_: _dafny.SeqWithoutIsStrInference([len(d_97_v13_) for d_157_i6_ in range(default__.abs(950))])})
        d_156_v64_ = (d_156_v64_).set(d_155_v63_, d_97_v13_)
        d_158_v65_: _dafny.Map
        d_158_v65_ = _dafny.Map({(d_154_v62_).f24: True})
        if ((d_154_v62_).f24) in (d_158_v65_):
            d_159_v66_: _dafny.Map
            d_159_v66_ = _dafny.Map({default__.safeDivisionInt(d_101_v15_, d_87_v3_): d_92_v8_})
            d_159_v66_ = (d_159_v66_).set(271, d_92_v8_)
            d_160_v67_: _dafny.Array
            def lambda5_(d_161_v0_):
                def lambda6_(d_162_i7_):
                    return d_161_v0_

                return lambda6_

            init3_ = lambda5_(d_84_v0_)
            nw34_ = _dafny.Array(None, 5)
            for i0_3_ in range(nw34_.length(0)):
                nw34_[i0_3_] = init3_(i0_3_)
            d_160_v67_ = nw34_
            source0_ = D1_DC3(d_160_v67_, 712, not((-154) >= (((d_102_v16_)[(d_154_v62_).f24] if ((d_154_v62_).f24) in (d_102_v16_) else -409))))
            if source0_.is_DC3:
                d_163___mcc_h0_ = source0_.cf3
                d_164___mcc_h1_ = source0_.cf4
                d_165___mcc_h2_ = source0_.cf5
                d_166_cf5_ = d_165___mcc_h2_
                d_167_cf4_ = d_164___mcc_h1_
                d_168_cf3_ = d_163___mcc_h0_
                d_169_v68_: D0
                d_169_v68_ = D0_DC1(d_96_v12_)
                d_170_v69_: _dafny.Map
                d_170_v69_ = _dafny.Map({d_100_v14_: d_169_v68_})
                d_171_v70_: int
                d_172_v71_: int
                out16_: int
                out17_: int
                out16_, out17_ = default__.m0((d_154_v62_).f24, default__.fm0(len(d_170_v69_), (d_154_v62_).f24, (d_84_v0_).set(default__.safeIndex(d_101_v15_, len(d_84_v0_)), d_93_v9_), d_98_globalState_), d_101_v15_, d_98_globalState_)
                d_171_v70_ = out16_
                d_172_v71_ = out17_
                d_86_v2_ = d_166_cf5_
                index31_ = default__.safeIndex(19, (d_152_v60_).length(0))
                (d_152_v60_)[index31_] = (False) and (d_86_v2_)
                index32_ = default__.safeIndex(19, (d_152_v60_).length(0))
                rhs31_ = (len(d_95_v11_)) <= ((0) - (d_167_cf4_))
                rhs32_ = (d_154_v62_).f24
                rhs33_ = (d_154_v62_).f24
                lhs23_ = d_152_v60_
                lhs24_ = default__.safeIndex(19, (d_152_v60_).length(0))
                lhs25_ = d_98_globalState_
                lhs23_[lhs24_] = rhs31_
                lhs25_.f2 = rhs32_
                d_166_cf5_ = rhs33_
                d_173_v72_: int
                d_174_v73_: int
                out18_: int
                out19_: int
                out18_, out19_ = default__.m0((_dafny.Set({(d_152_v60_)[default__.safeIndex(19, (d_152_v60_).length(0))]})).issubset(_dafny.Set({(d_154_v62_).f24, True, d_166_cf5_})), (d_154_v62_).f24, d_101_v15_, d_98_globalState_)
                d_173_v72_ = out18_
                d_174_v73_ = out19_
            elif source0_.is_DC4:
                d_175___mcc_h3_ = source0_.cf6
                d_176___mcc_h4_ = source0_.cf7
                d_177_cf7_ = d_176___mcc_h4_
                d_178_cf6_ = d_175___mcc_h3_
                d_179_v74_: D2
                d_179_v74_ = D2_DC7(d_101_v15_, _dafny.CodePoint('a'), d_154_v62_, (d_154_v62_).f24)
                d_180_v75_: _dafny.MultiSet
                d_180_v75_ = _dafny.MultiSet([d_154_v62_, (d_179_v74_).cf12])
                d_181_v76_: _dafny.MultiSet
                d_181_v76_ = _dafny.MultiSet([(d_180_v75_).cardinality])
                d_101_v15_ = ((d_102_v16_)[(d_154_v62_).f24] if ((d_154_v62_).f24) in (d_102_v16_) else default__.safeDivisionInt(d_100_v14_, (d_181_v76_).cardinality))
                pat_let_tv0_ = d_100_v14_
                pat_let_tv1_ = d_154_v62_
                d_182_v77_: _dafny.Array
                nw35_ = _dafny.Array(None, 16)
                nw35_[int(0)] = d_179_v74_
                nw35_[int(1)] = D2_DC7(d_101_v15_, d_93_v9_, d_154_v62_, (d_154_v62_).f24)
                nw35_[int(2)] = d_179_v74_
                nw35_[int(3)] = d_179_v74_
                nw35_[int(4)] = (d_179_v74_ if not(d_86_v2_) else d_179_v74_)
                nw35_[int(5)] = D2_DC7(d_101_v15_, d_93_v9_, d_154_v62_, d_177_cf7_)
                nw35_[int(6)] = d_179_v74_
                nw35_[int(7)] = d_179_v74_
                nw35_[int(8)] = d_179_v74_
                nw35_[int(9)] = d_179_v74_
                nw35_[int(10)] = d_179_v74_
                nw35_[int(11)] = d_179_v74_
                nw35_[int(12)] = (d_179_v74_ if (d_154_v62_).f24 else d_179_v74_)
                nw35_[int(13)] = d_179_v74_
                def iife6_(_pat_let0_0):
                    def iife7_(d_183_dt__update__tmp_h0_):
                        def iife8_(_pat_let1_0):
                            def iife9_(d_184_dt__update_hcf10_h0_):
                                def iife10_(_pat_let2_0):
                                    def iife11_(d_185_dt__update_hcf12_h0_):
                                        return D2_DC7(d_184_dt__update_hcf10_h0_, (d_183_dt__update__tmp_h0_).cf11, d_185_dt__update_hcf12_h0_, (d_183_dt__update__tmp_h0_).cf13)
                                    return iife11_(_pat_let2_0)
                                return iife10_(pat_let_tv1_)
                            return iife9_(_pat_let1_0)
                        return iife8_(pat_let_tv0_)
                    return iife7_(_pat_let0_0)
                nw35_[int(14)] = iife6_(d_179_v74_)
                nw35_[int(15)] = d_179_v74_
                d_182_v77_ = nw35_
                index33_ = default__.safeIndex(490, (d_182_v77_).length(0))
                (d_182_v77_)[index33_] = d_179_v74_
                index34_ = default__.safeIndex(311, (d_92_v8_).length(0))
                (d_92_v8_)[index34_] = d_87_v3_
                index35_ = default__.safeIndex(490, (d_182_v77_).length(0))
                index36_ = default__.safeIndex(311, (d_92_v8_).length(0))
                rhs34_ = (d_179_v74_ if (d_100_v14_) > (309) else D2_DC7(len(default__.fm6(((d_181_v76_)[d_101_v15_] if (d_101_v15_) in (d_181_v76_) else d_87_v3_), (d_154_v62_).f24, (d_154_v62_).f24, d_98_globalState_)), d_93_v9_, d_154_v62_, d_86_v2_))
                rhs35_ = d_101_v15_
                rhs36_ = d_101_v15_
                rhs37_ = d_87_v3_
                lhs26_ = d_182_v77_
                lhs27_ = default__.safeIndex(490, (d_182_v77_).length(0))
                lhs28_ = d_98_globalState_
                lhs29_ = d_98_globalState_
                lhs30_ = d_92_v8_
                lhs31_ = default__.safeIndex(311, (d_92_v8_).length(0))
                lhs26_[lhs27_] = rhs34_
                lhs28_.f23 = rhs35_
                lhs29_.f23 = rhs36_
                lhs30_[lhs31_] = rhs37_
                d_154_v62_ = d_154_v62_
                d_87_v3_ = (0) - ((d_97_v13_)[default__.safeIndex(default__.safeDivisionInt(len(d_84_v0_), len(d_84_v0_)), len(d_97_v13_))])
            elif source0_.is_DC2:
                d_186___mcc_h5_ = source0_.cf2
                d_187_cf2_ = d_186___mcc_h5_
                d_188_v78_: C0
                nw36_ = C0()
                nw36_.ctor__((d_154_v62_).f24)
                d_188_v78_ = nw36_
                d_189_v79_: int
                d_190_v80_: int
                out20_: int
                out21_: int
                out20_, out21_ = default__.m0((d_154_v62_).f24, (d_188_v78_).f24, d_101_v15_, d_98_globalState_)
                d_189_v79_ = out20_
                d_190_v80_ = out21_
                (d_98_globalState_).f18 = d_86_v2_
                (d_98_globalState_).f20 = d_84_v0_
            elif True:
                d_191___mcc_h6_ = source0_.cf8
                d_192_cf8_ = d_191___mcc_h6_
                index37_ = default__.safeIndex(362, (d_152_v60_).length(0))
                (d_152_v60_)[index37_] = (d_154_v62_).f24
                index38_ = default__.safeIndex(362, (d_152_v60_).length(0))
                (d_152_v60_)[index38_] = default__.fm0(default__.safeDivisionInt(d_101_v15_, d_101_v15_), (d_154_v62_).f24, (d_84_v0_) + (d_84_v0_), d_98_globalState_)
                d_87_v3_ = len(d_84_v0_)
                d_193_v81_: _dafny.Array
                def lambda7_(d_194_v62_):
                    def lambda8_(d_195_i8_):
                        return (False if (d_194_v62_).f24 else (d_194_v62_).f24)

                    return lambda8_

                init4_ = lambda7_(d_154_v62_)
                nw37_ = _dafny.Array(None, 21)
                for i0_4_ in range(nw37_.length(0)):
                    nw37_[i0_4_] = init4_(i0_4_)
                d_193_v81_ = nw37_
                d_193_v81_ = d_193_v81_
                d_196_v82_: D2
                d_196_v82_ = D2_DC7(d_100_v14_, d_93_v9_, d_154_v62_, (d_154_v62_).f24)
                d_197_v83_: _dafny.Map
                d_197_v83_ = _dafny.Map({d_154_v62_: not((d_154_v62_).f24)})
                d_198_v84_: int
                d_199_v85_: int
                out22_: int
                out23_: int
                out22_, out23_ = default__.m0(default__.fm0(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_200_i9_ in range(default__.abs(-210))])), (d_152_v60_)[default__.safeIndex(362, (d_152_v60_).length(0))], d_84_v0_, d_98_globalState_), default__.fm0(65, (d_196_v82_).cf13, d_84_v0_, d_98_globalState_), len(d_197_v83_), d_98_globalState_)
                d_198_v84_ = out22_
                d_199_v85_ = out23_
            d_102_v16_ = _dafny.MultiSet([d_86_v2_, (741) == (d_101_v15_), (d_154_v62_).f24])
            d_201_v86_: _dafny.Map
            d_201_v86_ = _dafny.Map({11: 619})
            d_87_v3_ = (((d_201_v86_)[(0) - (d_100_v14_)] if ((0) - (d_100_v14_)) in (d_201_v86_) else d_101_v15_)) + (len(d_84_v0_))
            d_202_v87_: int
            d_203_v88_: int
            out24_: int
            out25_: int
            out24_, out25_ = default__.m0((d_154_v62_).f24, (d_154_v62_).f24, d_101_v15_, d_98_globalState_)
            d_202_v87_ = out24_
            d_203_v88_ = out25_
        elif True:
            d_204_v89_: _dafny.Array
            nw38_ = _dafny.Array(D1.default()(), 9)
            d_204_v89_ = nw38_
            d_205_v90_: _dafny.Array
            nw39_ = _dafny.Array(None, 12)
            nw39_[int(0)] = d_84_v0_
            nw39_[int(1)] = d_84_v0_
            nw39_[int(2)] = d_84_v0_
            nw39_[int(3)] = d_84_v0_
            nw39_[int(4)] = d_84_v0_
            nw39_[int(5)] = d_84_v0_
            nw39_[int(6)] = _dafny.SeqWithoutIsStrInference([d_93_v9_ for d_206_i10_ in range(default__.abs(168))])
            nw39_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqgonfj"))
            nw39_[int(8)] = d_84_v0_
            nw39_[int(9)] = d_84_v0_
            nw39_[int(10)] = d_84_v0_
            nw39_[int(11)] = d_84_v0_
            d_205_v90_ = nw39_
            d_207_v91_: D1
            d_207_v91_ = D1_DC3(d_205_v90_, d_101_v15_, True)
            index39_ = default__.safeIndex(633, (d_204_v89_).length(0))
            (d_204_v89_)[index39_] = d_207_v91_
            index40_ = default__.safeIndex(633, (d_204_v89_).length(0))
            (d_204_v89_)[index40_] = d_207_v91_
            d_93_v9_ = _dafny.CodePoint('c')
            d_208_v92_: _dafny.Map
            d_208_v92_ = _dafny.Map({(d_100_v14_) - (d_87_v3_): d_93_v9_})
            d_208_v92_ = (d_208_v92_).set(d_87_v3_, d_93_v9_)
            (d_98_globalState_).f18 = (d_154_v62_).f24
            d_209_v93_: D1
            d_209_v93_ = D1_DC2((d_100_v14_) >= (len(d_97_v13_)))
            pat_let_tv2_ = d_86_v2_
            def iife12_(_pat_let3_0):
                def iife13_(d_210_dt__update__tmp_h1_):
                    def iife14_(_pat_let4_0):
                        def iife15_(d_211_dt__update_hcf2_h0_):
                            return D1_DC2(d_211_dt__update_hcf2_h0_)
                        return iife15_(_pat_let4_0)
                    return iife14_(pat_let_tv2_)
                return iife13_(_pat_let3_0)
            rhs38_ = d_100_v14_
            rhs39_ = iife12_((d_209_v93_ if (d_154_v62_).f24 else d_209_v93_))
            rhs40_ = d_93_v9_
            d_100_v14_ = rhs38_
            d_209_v93_ = rhs39_
            d_93_v9_ = rhs40_
        d_88_v4_ = (d_88_v4_).set(d_86_v2_, (0) - (((0) - (d_101_v15_)) - (d_87_v3_)))
        hi2_ = d_101_v15_
        for d_212_i11_ in range((d_100_v14_) + (d_101_v15_), hi2_):
            d_213_v94_: _dafny.MultiSet
            d_213_v94_ = _dafny.MultiSet([d_87_v3_])
            if ((0) - ((d_101_v15_) * (d_101_v15_))) > (((d_213_v94_)[d_87_v3_] if (d_87_v3_) in (d_213_v94_) else d_101_v15_)):
                index41_ = default__.safeIndex(704, (d_152_v60_).length(0))
                (d_152_v60_)[index41_] = d_86_v2_
                index42_ = default__.safeIndex(704, (d_152_v60_).length(0))
                (d_152_v60_)[index42_] = (d_154_v62_).f24
                d_214_v95_: _dafny.Seq
                d_214_v95_ = _dafny.SeqWithoutIsStrInference([d_96_v12_, d_96_v12_, d_96_v12_, d_96_v12_])
                (d_98_globalState_).f15 = (d_214_v95_)[default__.safeIndex(d_87_v3_, len(d_214_v95_))]
                d_93_v9_ = d_93_v9_
                d_215_v96_: _dafny.Seq
                d_215_v96_ = _dafny.SeqWithoutIsStrInference([True, (d_154_v62_).f24, d_86_v2_])
                d_216_v97_: _dafny.Array
                nw40_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_216_v97_ = nw40_
                d_217_v98_: _dafny.Array
                nw41_ = _dafny.Array(_dafny.Set({}), 24)
                d_217_v98_ = nw41_
                index43_ = default__.safeIndex(671, (d_216_v97_).length(0))
                (d_216_v97_)[index43_] = d_217_v98_
                d_218_v99_: _dafny.Seq
                d_218_v99_ = _dafny.SeqWithoutIsStrInference([d_217_v98_, d_217_v98_, d_217_v98_])
                index44_ = default__.safeIndex(671, (d_216_v97_).length(0))
                index45_ = default__.safeIndex(704, (d_152_v60_).length(0))
                rhs41_ = (0) - (d_212_i11_)
                rhs42_ = (default__.fm7(d_98_globalState_)) + (d_97_v13_)
                rhs43_ = (d_215_v96_) + (_dafny.SeqWithoutIsStrInference([(d_152_v60_)[default__.safeIndex(704, (d_152_v60_).length(0))], (d_154_v62_).f24]))
                rhs44_ = (d_218_v99_)[default__.safeIndex(d_101_v15_, len(d_218_v99_))]
                rhs45_ = (d_87_v3_) < (d_212_i11_)
                lhs32_ = d_216_v97_
                lhs33_ = default__.safeIndex(671, (d_216_v97_).length(0))
                lhs34_ = d_152_v60_
                lhs35_ = default__.safeIndex(704, (d_152_v60_).length(0))
                d_87_v3_ = rhs41_
                d_97_v13_ = rhs42_
                d_215_v96_ = rhs43_
                lhs32_[lhs33_] = rhs44_
                lhs34_[lhs35_] = rhs45_
                d_219_v100_: _dafny.Seq
                d_219_v100_ = _dafny.SeqWithoutIsStrInference([d_102_v16_, d_102_v16_])
                d_220_v101_: _dafny.Map
                d_220_v101_ = _dafny.Map({not((d_154_v62_).f24): _dafny.MultiSet(d_219_v100_)})
                d_221_v102_: _dafny.MultiSet
                d_221_v102_ = _dafny.MultiSet([_dafny.MultiSet(d_215_v96_), d_102_v16_])
                d_222_v103_: int
                d_223_v104_: int
                out26_: int
                out27_: int
                out26_, out27_ = default__.m0((d_154_v62_).f24, d_86_v2_, (((d_220_v101_)[(d_152_v60_)[default__.safeIndex(704, (d_152_v60_).length(0))]] if ((d_152_v60_)[default__.safeIndex(704, (d_152_v60_).length(0))]) in (d_220_v101_) else d_221_v102_)).cardinality, d_98_globalState_)
                d_222_v103_ = out26_
                d_223_v104_ = out27_
            elif True:
                d_95_v11_ = (d_95_v11_).set(d_100_v14_, (d_154_v62_).f24)
                d_224_v105_: _dafny.Set
                d_224_v105_ = _dafny.Set({d_86_v2_})
                d_225_v106_: _dafny.Set
                d_225_v106_ = _dafny.Set({(_dafny.Set({(d_154_v62_).f24})) - (d_224_v105_), d_224_v105_, d_224_v105_, d_224_v105_, d_224_v105_})
                d_225_v106_ = _dafny.Set({(d_224_v105_) - (d_224_v105_), d_224_v105_})
                (d_98_globalState_).f20 = ((d_84_v0_).set(default__.safeIndex(d_87_v3_, len(d_84_v0_)), d_93_v9_)) + (d_84_v0_)
                d_226_v107_: D2
                d_226_v107_ = D2_DC6(d_84_v0_)
                d_226_v107_ = d_226_v107_
                d_87_v3_ = d_212_i11_
            d_227_v108_: _dafny.Seq
            d_227_v108_ = _dafny.SeqWithoutIsStrInference([(d_154_v62_ if (d_154_v62_).f24 else d_154_v62_), d_154_v62_, d_154_v62_])
            d_228_v109_: _dafny.Array
            nw42_ = _dafny.Array(None, 7)
            nw42_[int(0)] = d_84_v0_
            nw42_[int(1)] = d_84_v0_
            nw42_[int(2)] = (d_84_v0_) + (d_84_v0_)
            nw42_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hoeogo"))
            nw42_[int(4)] = d_84_v0_
            nw42_[int(5)] = d_84_v0_
            nw42_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsmngmhxs"))
            d_228_v109_ = nw42_
            index46_ = default__.safeIndex(524, (d_228_v109_).length(0))
            (d_228_v109_)[index46_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ite"))
            d_229_v110_: _dafny.Map
            d_229_v110_ = _dafny.Map({d_84_v0_: d_84_v0_})
            index47_ = default__.safeIndex(524, (d_228_v109_).length(0))
            rhs46_ = ((d_229_v110_)[(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvomcrdu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pk")))] if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvomcrdu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pk")))) in (d_229_v110_) else (d_84_v0_) + (d_84_v0_))
            rhs47_ = d_227_v108_
            rhs48_ = ((d_84_v0_) + (d_84_v0_)) + (d_84_v0_)
            lhs36_ = d_98_globalState_
            lhs37_ = d_228_v109_
            lhs38_ = default__.safeIndex(524, (d_228_v109_).length(0))
            lhs36_.f20 = rhs46_
            d_227_v108_ = rhs47_
            lhs37_[lhs38_] = rhs48_
            d_230_v111_: _dafny.Array
            def lambda9_(d_231_v16_):
                def lambda10_(d_232_i12_):
                    return d_231_v16_

                return lambda10_

            init5_ = lambda9_(d_102_v16_)
            nw43_ = _dafny.Array(None, 27)
            for i0_5_ in range(nw43_.length(0)):
                nw43_[i0_5_] = init5_(i0_5_)
            d_230_v111_ = nw43_
            index48_ = default__.safeIndex(971, (d_230_v111_).length(0))
            (d_230_v111_)[index48_] = d_102_v16_
            index49_ = default__.safeIndex(971, (d_230_v111_).length(0))
            (d_230_v111_)[index49_] = ((d_102_v16_) | (d_102_v16_)) - ((d_102_v16_).set((d_154_v62_).f24, default__.abs(d_212_i11_)))
            d_233_v112_: _dafny.Set
            d_233_v112_ = _dafny.Set({(d_154_v62_).f24, d_86_v2_})
            source1_ = D1_DC2((_dafny.Set({d_86_v2_})).issubset(d_233_v112_))
            if source1_.is_DC3:
                d_234___mcc_h7_ = source1_.cf3
                d_235___mcc_h8_ = source1_.cf4
                d_236___mcc_h9_ = source1_.cf5
                d_237_cf5_ = d_236___mcc_h9_
                d_238_cf4_ = d_235___mcc_h8_
                d_239_cf3_ = d_234___mcc_h7_
                index50_ = default__.safeIndex(233, (d_152_v60_).length(0))
                (d_152_v60_)[index50_] = False
                index51_ = default__.safeIndex(233, (d_152_v60_).length(0))
                (d_152_v60_)[index51_] = (d_212_i11_) == ((0) - ((0) - (default__.safeDivisionInt(d_87_v3_, d_101_v15_))))
                d_240_v113_: _dafny.Map
                d_240_v113_ = _dafny.Map({d_100_v14_: d_238_cf4_})
                d_240_v113_ = (d_240_v113_).set(103, (d_100_v14_) * (d_101_v15_))
                d_241_v114_: C0
                nw44_ = C0()
                nw44_.ctor__((d_152_v60_)[default__.safeIndex(233, (d_152_v60_).length(0))])
                d_241_v114_ = nw44_
                (d_98_globalState_).f18 = (d_237_cf5_) == ((d_154_v62_).f24)
            elif source1_.is_DC4:
                d_242___mcc_h10_ = source1_.cf6
                d_243___mcc_h11_ = source1_.cf7
                d_244_cf7_ = d_243___mcc_h11_
                d_245_cf6_ = d_242___mcc_h10_
                d_246_v115_: _dafny.Map
                d_246_v115_ = _dafny.Map({d_97_v13_: (d_154_v62_).f24})
                (d_98_globalState_).f23 = (len(d_246_v115_)) * ((-803) * (-924))
                (d_98_globalState_).f23 = d_100_v14_
                d_247_v116_: int
                d_248_v117_: int
                out28_: int
                out29_: int
                out28_, out29_ = default__.m0(d_86_v2_, True, (0) - (d_212_i11_), d_98_globalState_)
                d_247_v116_ = out28_
                d_248_v117_ = out29_
                index52_ = default__.safeIndex(20, (d_152_v60_).length(0))
                (d_152_v60_)[index52_] = (d_154_v62_).f24
                index53_ = default__.safeIndex(15, (d_152_v60_).length(0))
                (d_152_v60_)[index53_] = (d_84_v0_) <= (d_84_v0_)
                d_249_v119_: D1
                d_249_v119_ = D1_DC4(d_245_cf6_, d_86_v2_)
                index54_ = default__.safeIndex(20, (d_152_v60_).length(0))
                index55_ = default__.safeIndex(15, (d_152_v60_).length(0))
                def iife16_():
                    coll6_ = _dafny.Set()
                    compr_6_: int
                    for compr_6_ in _dafny.IntegerRange(248, -932):
                        d_250_v118_: int = compr_6_
                        if ((248) <= (d_250_v118_)) and ((d_250_v118_) < (-932)):
                            coll6_ = coll6_.union(_dafny.Set([(d_250_v118_) + (d_101_v15_)]))
                    return _dafny.Set(coll6_)
                rhs49_ = (728) not in (iife16_()
                )
                rhs50_ = d_154_v62_
                rhs51_ = (d_154_v62_).f24
                rhs52_ = (d_245_cf6_) + ((d_249_v119_).cf6)
                rhs53_ = (d_154_v62_).f24
                lhs39_ = d_152_v60_
                lhs40_ = default__.safeIndex(20, (d_152_v60_).length(0))
                lhs41_ = d_98_globalState_
                lhs42_ = d_152_v60_
                lhs43_ = default__.safeIndex(15, (d_152_v60_).length(0))
                lhs39_[lhs40_] = rhs49_
                d_154_v62_ = rhs50_
                lhs41_.f2 = rhs51_
                d_245_cf6_ = rhs52_
                lhs42_[lhs43_] = rhs53_
            elif source1_.is_DC2:
                d_251___mcc_h12_ = source1_.cf2
                d_252_cf2_ = d_251___mcc_h12_
                d_100_v14_ = (default__.safeModuloInt(978, d_87_v3_)) - ((d_212_i11_) + (d_212_i11_))
                d_253_v120_: _dafny.Array
                def lambda11_(d_254_i11_):
                    def lambda12_(d_255_i13_):
                        return _dafny.Map({False: (0) - (d_254_i11_)})

                    return lambda12_

                init6_ = lambda11_(d_212_i11_)
                nw45_ = _dafny.Array(None, 2)
                for i0_6_ in range(nw45_.length(0)):
                    nw45_[i0_6_] = init6_(i0_6_)
                d_253_v120_ = nw45_
                index56_ = default__.safeIndex(385, (d_253_v120_).length(0))
                (d_253_v120_)[index56_] = d_88_v4_
                index57_ = default__.safeIndex(385, (d_253_v120_).length(0))
                (d_253_v120_)[index57_] = d_88_v4_
                d_256_v121_: _dafny.Array
                nw46_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_256_v121_ = nw46_
                d_257_v122_: _dafny.Map
                d_257_v122_ = _dafny.Map({((d_230_v111_)[default__.safeIndex(971, (d_230_v111_).length(0))]) - ((d_230_v111_)[default__.safeIndex(971, (d_230_v111_).length(0))]): d_154_v62_})
                d_258_v123_: D3
                d_258_v123_ = D3_DC9(d_152_v60_)
                d_259_v124_: _dafny.Map
                d_259_v124_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwflyn")): _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_252_cf2_, (d_154_v62_).f24])): d_154_v62_})})
                rhs54_ = d_256_v121_
                rhs55_ = d_154_v62_
                rhs56_ = (d_258_v123_).cf15
                rhs57_ = d_86_v2_
                rhs58_ = ((d_259_v124_)[d_84_v0_] if (d_84_v0_) in (d_259_v124_) else (d_257_v122_) | (d_257_v122_))
                d_256_v121_ = rhs54_
                d_154_v62_ = rhs55_
                d_152_v60_ = rhs56_
                d_252_cf2_ = rhs57_
                d_257_v122_ = rhs58_
                (d_98_globalState_).f23 = (810) + (d_101_v15_)
            elif True:
                d_260___mcc_h13_ = source1_.cf8
                d_261_cf8_ = d_260___mcc_h13_
                d_262_v125_: _dafny.Map
                d_262_v125_ = _dafny.Map({d_212_i11_: d_100_v14_})
                d_263_v126_: _dafny.Map
                d_263_v126_ = _dafny.Map({((d_158_v65_)[(d_154_v62_).f24] if ((d_154_v62_).f24) in (d_158_v65_) else (d_154_v62_).f24): d_262_v125_})
                d_264_v127_: _dafny.Map
                d_264_v127_ = _dafny.Map({d_87_v3_: ((d_88_v4_)[False] if (False) in (d_88_v4_) else len(d_263_v126_))})
                (d_98_globalState_).f23 = default__.safeDivisionInt((422 if (d_154_v62_).f24 else len(d_84_v0_)), (0) - ((0) - (((0) - (d_100_v14_)) * ((0) - (len(d_264_v127_))))))
                (d_98_globalState_).f20 = ((d_154_v62_).fm5(d_98_globalState_) if d_86_v2_ else d_84_v0_)
                d_265_v128_: _dafny.Seq
                d_265_v128_ = _dafny.SeqWithoutIsStrInference([(d_154_v62_).f24])
                d_266_v129_: _dafny.Map
                d_266_v129_ = _dafny.Map({D1_DC4((d_265_v128_).set(default__.safeIndex(d_101_v15_, len(d_265_v128_)), not((d_154_v62_).f24)), (d_154_v62_).f24): d_101_v15_})
                d_266_v129_ = (d_266_v129_).set(default__.fm8(d_101_v15_, len(d_97_v13_), not((d_154_v62_).f24), d_98_globalState_), (0) - (d_212_i11_))
                (d_98_globalState_).f23 = d_212_i11_
        index58_ = default__.safeIndex(794, (d_152_v60_).length(0))
        (d_152_v60_)[index58_] = (d_154_v62_).f24
        index59_ = default__.safeIndex(794, (d_152_v60_).length(0))
        (d_152_v60_)[index59_] = ((d_87_v3_) * (d_100_v14_)) < (d_100_v14_)
        d_267_i14_: int
        d_267_i14_ = 0
        with _dafny.label("1"):
            while ((d_152_v60_)[default__.safeIndex(794, (d_152_v60_).length(0))] if (d_152_v60_)[default__.safeIndex(794, (d_152_v60_).length(0))] else (d_152_v60_)[default__.safeIndex(794, (d_152_v60_).length(0))]):
                with _dafny.c_label("1"):
                    if (d_267_i14_) >= (100):
                        raise _dafny.Break("1")
                    d_267_i14_ = (d_267_i14_) + (1)
                    (d_98_globalState_).f2 = (d_86_v2_) or ((d_154_v62_).f24)
                    (d_98_globalState_).f23 = -84
                    d_268_v130_: _dafny.Seq
                    d_268_v130_ = _dafny.SeqWithoutIsStrInference([(d_154_v62_).f24])
                    d_269_v131_: int
                    d_270_v132_: int
                    out30_: int
                    out31_: int
                    out30_, out31_ = default__.m0((d_268_v130_)[default__.safeIndex(d_87_v3_, len(d_268_v130_))], d_86_v2_, d_100_v14_, d_98_globalState_)
                    d_269_v131_ = out30_
                    d_270_v132_ = out31_
                    d_271_v133_: int
                    d_272_v134_: int
                    out32_: int
                    out33_: int
                    out32_, out33_ = default__.m0((d_152_v60_)[default__.safeIndex(794, (d_152_v60_).length(0))], (default__.fm3(984, d_98_globalState_)) == ((d_154_v62_).fm4(False, d_86_v2_, _dafny.MultiSet(default__.fm7(d_98_globalState_)), d_86_v2_, d_98_globalState_)), d_101_v15_, d_98_globalState_)
                    d_271_v133_ = out32_
                    d_272_v134_ = out33_
                    pass
            pass
        rhs59_ = d_92_v8_
        rhs60_ = -668
        d_92_v8_ = rhs59_
        d_87_v3_ = rhs60_
        d_273_v135_: _dafny.Set
        d_273_v135_ = _dafny.Set({d_86_v2_, d_86_v2_})
        d_274_v136_: _dafny.Array
        nw47_ = _dafny.Array(None, 4)
        nw47_[int(0)] = d_154_v62_
        nw47_[int(1)] = d_154_v62_
        nw47_[int(2)] = d_154_v62_
        nw47_[int(3)] = d_154_v62_
        d_274_v136_ = nw47_
        index60_ = default__.safeIndex(597, (d_274_v136_).length(0))
        (d_274_v136_)[index60_] = d_154_v62_
        d_275_v137_: D0
        d_275_v137_ = D0_DC0(d_93_v9_)
        d_276_v138_: _dafny.Map
        d_276_v138_ = _dafny.Map({d_275_v137_: d_154_v62_})
        pat_let_tv3_ = d_93_v9_
        pat_let_tv4_ = d_93_v9_
        d_277_v139_: _dafny.Seq
        def iife17_(_pat_let5_0):
            def iife18_(d_278_dt__update__tmp_h3_):
                def iife19_(_pat_let6_0):
                    def iife20_(d_279_dt__update_hcf0_h1_):
                        return D0_DC0(d_279_dt__update_hcf0_h1_)
                    return iife20_(_pat_let6_0)
                return iife19_(pat_let_tv3_)
            return iife18_(_pat_let5_0)
        def iife21_(_pat_let7_0):
            def iife22_(d_280_dt__update__tmp_h2_):
                def iife23_(_pat_let8_0):
                    def iife24_(d_281_dt__update_hcf0_h0_):
                        return D0_DC0(d_281_dt__update_hcf0_h0_)
                    return iife24_(_pat_let8_0)
                return iife23_(pat_let_tv4_)
            return iife22_(_pat_let7_0)
        d_277_v139_ = _dafny.SeqWithoutIsStrInference([((d_276_v138_)[iife17_(d_275_v137_)] if (iife21_(d_275_v137_)) in (d_276_v138_) else d_154_v62_), d_154_v62_])
        index61_ = default__.safeIndex(597, (d_274_v136_).length(0))
        rhs61_ = d_152_v60_
        rhs62_ = d_273_v135_
        rhs63_ = not(((d_152_v60_)[default__.safeIndex(794, (d_152_v60_).length(0))] if (d_152_v60_)[default__.safeIndex(794, (d_152_v60_).length(0))] else (d_154_v62_).f24))
        rhs64_ = (d_277_v139_)[default__.safeIndex(437, len(d_277_v139_))]
        lhs44_ = d_98_globalState_
        lhs45_ = d_274_v136_
        lhs46_ = default__.safeIndex(597, (d_274_v136_).length(0))
        d_152_v60_ = rhs61_
        d_273_v135_ = rhs62_
        lhs44_.f2 = rhs63_
        lhs45_[lhs46_] = rhs64_
        nw48_ = C0()
        nw48_.ctor__(not((d_152_v60_)[default__.safeIndex(794, (d_152_v60_).length(0))]))
        d_154_v62_ = nw48_
        d_101_v15_ = (d_100_v14_) - ((d_100_v14_) + (len(d_97_v13_)))
        d_282_v140_: _dafny.MultiSet
        d_282_v140_ = _dafny.MultiSet([default__.fm9(d_101_v15_, (d_152_v60_)[default__.safeIndex(794, (d_152_v60_).length(0))], d_87_v3_, d_86_v2_, d_98_globalState_), d_95_v11_])
        d_283_v141_: _dafny.Map
        d_283_v141_ = _dafny.Map({d_282_v140_: ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "balyols")) if (d_154_v62_).f24 else d_84_v0_)).set(default__.safeIndex(d_101_v15_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "balyols")) if (d_154_v62_).f24 else d_84_v0_))), d_93_v9_)})
        index62_ = default__.safeIndex(597, (d_274_v136_).length(0))
        rhs65_ = (d_97_v13_) + (_dafny.SeqWithoutIsStrInference([d_87_v3_]))
        rhs66_ = (d_274_v136_)[default__.safeIndex(597, (d_274_v136_).length(0))]
        rhs67_ = d_152_v60_
        rhs68_ = ((d_283_v141_)[(d_282_v140_) - (default__.fm10(d_86_v2_, d_86_v2_, (d_154_v62_).f24, d_98_globalState_))] if ((d_282_v140_) - (default__.fm10(d_86_v2_, d_86_v2_, (d_154_v62_).f24, d_98_globalState_))) in (d_283_v141_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhrm")))
        lhs47_ = d_274_v136_
        lhs48_ = default__.safeIndex(597, (d_274_v136_).length(0))
        lhs49_ = d_98_globalState_
        d_97_v13_ = rhs65_
        lhs47_[lhs48_] = rhs66_
        d_152_v60_ = rhs67_
        lhs49_.f20 = rhs68_
        d_284_v142_: C0
        nw49_ = C0()
        nw49_.ctor__(True)
        d_284_v142_ = nw49_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_92_v8_).length(0)):
            d_285_i15_: int = guard_loop_0_
            if (True) and (((0) <= (d_285_i15_)) and ((d_285_i15_) < ((d_92_v8_).length(0)))):
                (d_92_v8_)[(d_285_i15_)] = default__.safeModuloInt(d_285_i15_, len((_dafny.Map({d_100_v14_: d_87_v3_}) if (d_284_v142_).f24 else _dafny.Map({d_87_v3_: d_87_v3_}))))
        _dafny.print((d_84_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_v1_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_86_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_87_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_88_v4_) == (_dafny.Map({True: 56, False: 116}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v5_) == (_dafny.Map({_dafny.Map({True: 56}): 56}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_91_v7_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_92_v8_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_93_v9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_94_v10_) == (_dafny.Map({True: _dafny.CodePoint('n')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_95_v11_) == (_dafny.Map({56: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v12_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_v13_) == (_dafny.SeqWithoutIsStrInference([56, 56, -668]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f1) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spiw"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_98_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f3) == (_dafny.Map({_dafny.Map({True: 56}): 56}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f13)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_98_globalState_).f14).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_.f15)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f17)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_98_globalState_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_98_globalState_.f20).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f21) == (_dafny.SeqWithoutIsStrInference([56, 56]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_98_globalState_).f22) == (_dafny.SeqWithoutIsStrInference([56, 56]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_98_globalState_.f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_100_v14_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_101_v15_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v16_) == (_dafny.MultiSet([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v60_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v60_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_153_v61_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v62_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_155_v63_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_156_v64_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v65_) == (_dafny.Map({False: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_267_i14_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_273_v135_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v136_)[0]).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v136_)[1]).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v136_)[2]).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v136_)[3]).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v137_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_276_v138_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_277_v139_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_v140_) == (_dafny.MultiSet([_dafny.Map({21: False, 22: False, 23: False, 24: False, 25: False, 26: False, 27: False, 28: False, 29: False, 30: False, 31: False, 32: False, 33: False, 34: False, 35: False, 36: False, 37: False, 38: False, 39: False, 40: False, 41: False, 42: False, 43: False, 44: False, 45: False, 46: False, 47: False, 48: False, 49: False, 50: False, 51: False, 52: False, 53: False, 54: False, 55: False, 56: False, 57: False, 58: False, 0: False, 1: False, 2: False, 3: False, 4: False, 5: False, 6: False, 7: False, 8: False, 9: False, 10: False, 11: False, 12: False, 13: False, 14: False, 15: False, 16: False, 17: False, 18: False, 19: False, 20: False}), _dafny.Map({56: True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_283_v141_) == (_dafny.Map({_dafny.MultiSet([_dafny.Map({21: False, 22: False, 23: False, 24: False, 25: False, 26: False, 27: False, 28: False, 29: False, 30: False, 31: False, 32: False, 33: False, 34: False, 35: False, 36: False, 37: False, 38: False, 39: False, 40: False, 41: False, 42: False, 43: False, 44: False, 45: False, 46: False, 47: False, 48: False, 49: False, 50: False, 51: False, 52: False, 53: False, 54: False, 55: False, 56: False, 57: False, 58: False, 0: False, 1: False, 2: False, 3: False, 4: False, 5: False, 6: False, 7: False, 8: False, 9: False, 10: False, 11: False, 12: False, 13: False, 14: False, 15: False, 16: False, 17: False, 18: False, 19: False, 20: False}), _dafny.Map({56: True})]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nalyols"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v142_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
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
        return lambda: D1_DC3(_dafny.Array(None, 0), int(0), False)
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
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(int(0), _dafny.CodePoint('D'), None, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC7(D2, NamedTuple('DC7', [('cf10', Any), ('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({self.cf9.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC14(D4, NamedTuple('DC14', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC16(False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)

class D5_DC16(D5, NamedTuple('DC16', [('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC17(D5, NamedTuple('DC17', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC19(_dafny.Map({}), _dafny.Set({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)

class D6_DC19(D6, NamedTuple('DC19', [('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC20(D6, NamedTuple('DC20', [])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20)
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

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
        return lambda: D8_DC23(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.CodePoint('D'), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC23(D8, NamedTuple('DC23', [('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf32)}, {self.cf33.VerbatimString(True)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f1: _dafny.MultiSet = _dafny.MultiSet({})
        self.f2: bool = False
        self.f15: _dafny.Array = _dafny.Array(None, 0)
        self.f18: bool = False
        self.f20: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f23: int = int(0)
        self._f0: int = int(0)
        self._f3: _dafny.Map = _dafny.Map({})
        self._f4: int = int(0)
        self._f5: bool = False
        self._f6: bool = False
        self._f7: int = int(0)
        self._f8: int = int(0)
        self._f9: bool = False
        self._f10: _dafny.Array = _dafny.Array(None, 0)
        self._f11: str = _dafny.CodePoint('D')
        self._f12: int = int(0)
        self._f13: _dafny.Array = _dafny.Array(None, 0)
        self._f14: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f16: int = int(0)
        self._f17: _dafny.Array = _dafny.Array(None, 0)
        self._f19: bool = False
        self._f21: _dafny.Seq = _dafny.Seq({})
        self._f22: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23):
        (self)._f0 = f0
        (self).f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self).f18 = f18
        (self)._f19 = f19
        (self).f20 = f20
        (self)._f21 = f21
        (self)._f22 = f22
        (self).f23 = f23

    @property
    def f0(self):
        return self._f0
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
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
    def f9(self):
        return self._f9
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
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
    @property
    def f19(self):
        return self._f19
    @property
    def f21(self):
        return self._f21
    @property
    def f22(self):
        return self._f22

class C0:
    def  __init__(self):
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f24):
        (self)._f24 = f24

    def fm4(self, p0, p1, p2, p3, globalState):
        return ((845 if True else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwwnotdl"))))) * ((0) - (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({-399: not((self).f24)}))) for d_286_i0_ in range(default__.abs(641))]))).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([not((self).f24)])), 980]))).cardinality))

    def fm5(self, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lffyht"))

    @property
    def f24(self):
        return self._f24
