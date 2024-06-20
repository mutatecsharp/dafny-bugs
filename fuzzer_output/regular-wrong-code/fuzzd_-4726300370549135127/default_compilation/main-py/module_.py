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
        return True

    @staticmethod
    def fm2(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([D3_DC8()])) + (_dafny.SeqWithoutIsStrInference([D3_DC8(), D3_DC8()])))).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gyruwxl"))), (123) * (758)])

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return len(((_dafny.Map({False: False})) | (_dafny.Map({False: True}))) | ((_dafny.Map({not(not(False)): True})) | (_dafny.Map({False: False}))))

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        def iife0_():
            def iife4_():
                def iife6_():
                    coll6_ = _dafny.Map()
                    compr_6_: str
                    for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(709))])).Elements:
                        d_1_v1_: str = compr_6_
                        if (d_1_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(709))])):
                            coll6_[d_1_v1_] = 904
                    return _dafny.Map(coll6_)
                coll4_ = _dafny.Map()
                def iife5_():
                    coll5_ = _dafny.Map()
                    compr_5_: str
                    for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(709))])).Elements:
                        d_1_v1_: str = compr_5_
                        if (d_1_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(709))])):
                            coll5_[d_1_v1_] = 904
                    return _dafny.Map(coll5_)
                compr_4_: str
                for compr_4_ in (iife5_()
                ).keys.Elements:
                    d_2_v0_: str = compr_4_
                    if (d_2_v0_) in (iife6_()
                    ):
                        coll4_[d_2_v0_] = _dafny.SeqWithoutIsStrInference([587])
                return _dafny.Map(coll4_)
            coll0_ = _dafny.Set()
            def iife1_():
                def iife3_():
                    coll3_ = _dafny.Map()
                    compr_3_: str
                    for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(709))])).Elements:
                        d_1_v1_: str = compr_3_
                        if (d_1_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(709))])):
                            coll3_[d_1_v1_] = 904
                    return _dafny.Map(coll3_)
                coll1_ = _dafny.Map()
                def iife2_():
                    coll2_ = _dafny.Map()
                    compr_2_: str
                    for compr_2_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(709))])).Elements:
                        d_1_v1_: str = compr_2_
                        if (d_1_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(709))])):
                            coll2_[d_1_v1_] = 904
                    return _dafny.Map(coll2_)
                compr_1_: str
                for compr_1_ in (iife2_()
                ).keys.Elements:
                    d_2_v0_: str = compr_1_
                    if (d_2_v0_) in (iife3_()
                    ):
                        coll1_[d_2_v0_] = _dafny.SeqWithoutIsStrInference([587])
                return _dafny.Map(coll1_)
            compr_0_: str
            for compr_0_ in ((_dafny.Map({_dafny.CodePoint('o'): _dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True}))])})) | (iife1_()
            )).keys.Elements:
                d_3_v2_: str = compr_0_
                if (d_3_v2_) in ((_dafny.Map({_dafny.CodePoint('o'): _dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True}))])})) | (iife4_()
                )):
                    coll0_ = coll0_.union(_dafny.Set([d_3_v2_]))
            return _dafny.Set(coll0_)
        return iife0_()
        

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        return (_dafny.MultiSet([(0) - ((_dafny.MultiSet([not(False), False])).cardinality)])) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([844]))]) if False else _dafny.MultiSet([len(_dafny.Map({True: False})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_4_i0_ in range(default__.abs(118))]))])))

    @staticmethod
    def fm6(globalState):
        return _dafny.CodePoint('q')

    @staticmethod
    def fm7(globalState):
        return (_dafny.SeqWithoutIsStrInference([not(True), False, True])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm8(globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bl")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oj"))]))

    @staticmethod
    def Main(noArgsParameter__):
        d_5_v0_: _dafny.Array
        nw0_ = _dafny.Array(int(0), 6)
        d_5_v0_ = nw0_
        d_6_v1_: bool
        d_6_v1_ = True
        d_7_v2_: _dafny.Seq
        d_7_v2_ = _dafny.SeqWithoutIsStrInference([not(d_6_v1_), d_6_v1_])
        d_8_v3_: int
        d_8_v3_ = 861
        d_9_v4_: _dafny.Set
        d_9_v4_ = _dafny.Set({d_8_v3_, d_8_v3_})
        d_10_v5_: bool
        d_10_v5_ = False
        d_11_v6_: _dafny.Array
        nw1_ = _dafny.Array(None, 12)
        nw1_[int(0)] = True
        nw1_[int(1)] = (d_10_v5_)
        nw1_[int(2)] = False
        nw1_[int(3)] = (d_7_v2_)[default__.safeIndex(d_8_v3_, len(d_7_v2_))]
        nw1_[int(4)] = d_6_v1_
        nw1_[int(5)] = d_6_v1_
        nw1_[int(6)] = not(d_6_v1_)
        nw1_[int(7)] = d_6_v1_
        nw1_[int(8)] = d_6_v1_
        nw1_[int(9)] = d_6_v1_
        nw1_[int(10)] = not(d_6_v1_)
        nw1_[int(11)] = d_6_v1_
        d_11_v6_ = nw1_
        d_12_v7_: str
        d_12_v7_ = _dafny.CodePoint('l')
        d_13_v8_: _dafny.Map
        d_13_v8_ = _dafny.Map({d_8_v3_: d_12_v7_})
        d_14_v9_: _dafny.Array
        nw2_ = _dafny.Array(_dafny.Set({}), 5)
        d_14_v9_ = nw2_
        d_15_v10_: _dafny.Map
        d_15_v10_ = _dafny.Map({d_6_v1_: d_12_v7_})
        d_16_v11_: _dafny.MultiSet
        d_16_v11_ = _dafny.MultiSet([d_15_v10_, d_15_v10_])
        d_17_v12_: D1
        d_17_v12_ = D1_DC1(d_5_v0_)
        d_18_globalState_: GlobalState
        nw3_ = GlobalState()
        nw3_.ctor__(d_5_v0_, (d_7_v2_) + (d_7_v2_), False, d_9_v4_, d_11_v6_, d_5_v0_, d_13_v8_, True, False, d_14_v9_, False, _dafny.CodePoint('s'), d_11_v6_, 242, True, 442, False, -101, 92, 999, (d_16_v11_) | (d_16_v11_), False, (d_17_v12_).cf1, (d_17_v12_).cf1, 279, 491)
        d_18_globalState_ = nw3_
        if (default__.safeDivisionInt(d_8_v3_, d_8_v3_)) >= (len(_dafny.SeqWithoutIsStrInference([d_6_v1_, d_6_v1_, d_6_v1_]))):
            d_11_v6_ = d_11_v6_
            d_19_v13_: _dafny.Seq
            d_19_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shmbaf"))
            d_20_v14_: _dafny.MultiSet
            d_20_v14_ = _dafny.MultiSet([d_19_v13_, d_19_v13_, d_19_v13_])
            d_21_v15_: _dafny.Map
            d_21_v15_ = _dafny.Map({default__.fm0(_dafny.SeqWithoutIsStrInference([d_8_v3_]), d_8_v3_, (0) - (len(d_19_v13_)), d_18_globalState_): (0) - (len(d_9_v4_))})
            d_8_v3_ = ((d_20_v14_)[_dafny.SeqWithoutIsStrInference([d_12_v7_ for d_22_i0_ in range(default__.abs(-845))])] if (_dafny.SeqWithoutIsStrInference([d_12_v7_ for d_23_i0_ in range(default__.abs(-845))])) in (d_20_v14_) else (len(d_21_v15_) if d_6_v1_ else -872))
            d_6_v1_ = (d_12_v7_) not in (d_19_v13_)
            if (d_9_v4_).ispropersubset(d_9_v4_):
                (d_18_globalState_).f15 = d_8_v3_
                d_6_v1_ = d_6_v1_
                d_24_v16_: C0
                nw4_ = C0()
                nw4_.ctor__(d_6_v1_)
                d_24_v16_ = nw4_
                (d_24_v16_).f26 = not((d_8_v3_) > (default__.safeDivisionInt(d_8_v3_, -365)))
                d_25_v17_: bool
                d_26_v18_: _dafny.Seq
                out0_: bool
                out1_: _dafny.Seq
                out0_, out1_ = (d_24_v16_).m0(d_24_v16_.f26, (0) - (len(d_19_v13_)), d_18_globalState_)
                d_25_v17_ = out0_
                d_26_v18_ = out1_
            elif True:
                d_27_v19_: C0
                nw5_ = C0()
                nw5_.ctor__((d_8_v3_) != (d_8_v3_))
                d_27_v19_ = nw5_
                d_28_v20_: _dafny.Map
                d_28_v20_ = _dafny.Map({d_8_v3_: (0) - (d_8_v3_)})
                d_29_v21_: D1
                d_29_v21_ = D1_DC2(d_8_v3_, (0) - (d_8_v3_), d_8_v3_, (d_28_v20_) | (d_28_v20_))
                d_29_v21_ = d_29_v21_
                d_30_v22_: _dafny.Array
                nw6_ = _dafny.Array(None, 12)
                nw6_[int(0)] = d_5_v0_
                nw6_[int(1)] = d_5_v0_
                nw6_[int(2)] = d_5_v0_
                nw6_[int(3)] = d_5_v0_
                nw6_[int(4)] = d_5_v0_
                nw6_[int(5)] = d_5_v0_
                nw6_[int(6)] = d_5_v0_
                nw6_[int(7)] = d_5_v0_
                nw6_[int(8)] = d_5_v0_
                nw6_[int(9)] = d_5_v0_
                nw6_[int(10)] = d_5_v0_
                nw6_[int(11)] = d_5_v0_
                d_30_v22_ = nw6_
                index0_ = default__.safeIndex(333, (d_30_v22_).length(0))
                (d_30_v22_)[index0_] = d_5_v0_
                index1_ = default__.safeIndex(333, (d_30_v22_).length(0))
                rhs0_ = d_19_v13_
                rhs1_ = d_8_v3_
                rhs2_ = d_5_v0_
                rhs3_ = d_27_v19_.f26
                rhs4_ = 850
                lhs0_ = d_18_globalState_
                lhs1_ = d_30_v22_
                lhs2_ = default__.safeIndex(333, (d_30_v22_).length(0))
                lhs3_ = d_27_v19_
                lhs4_ = d_18_globalState_
                d_19_v13_ = rhs0_
                lhs0_.f19 = rhs1_
                lhs1_[lhs2_] = rhs2_
                lhs3_.f26 = rhs3_
                lhs4_.f19 = rhs4_
                d_19_v13_ = d_19_v13_
                d_31_v23_: D1
                d_31_v23_ = D1_DC3(d_8_v3_)
                d_31_v23_ = d_31_v23_
            d_32_v24_: C0
            nw7_ = C0()
            nw7_.ctor__(d_6_v1_)
            d_32_v24_ = nw7_
            d_33_v25_: _dafny.Set
            d_33_v25_ = _dafny.Set({d_32_v24_})
            d_34_v26_: _dafny.MultiSet
            d_34_v26_ = _dafny.MultiSet([len(d_33_v25_), d_8_v3_, 721, d_8_v3_])
            d_6_v1_ = ((d_34_v26_).cardinality) >= (d_8_v3_)
        elif True:
            d_5_v0_ = d_5_v0_
            d_12_v7_ = d_12_v7_
            d_35_v27_: _dafny.Seq
            d_35_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
            d_36_v28_: _dafny.Set
            d_36_v28_ = _dafny.Set({d_6_v1_})
            (d_18_globalState_).f15 = ((len(d_35_v27_)) * (d_8_v3_)) - ((d_8_v3_ if d_6_v1_ else len(d_36_v28_)))
            d_37_v29_: _dafny.Array
            def lambda0_(d_38_v3_, d_39_v2_):
                def lambda1_(d_40_i1_):
                    return _dafny.SeqWithoutIsStrInference([D1_DC2(123, d_38_v3_, d_38_v3_, _dafny.Map({d_38_v3_: (_dafny.MultiSet(d_39_v2_)).cardinality}))])

                return lambda1_

            init0_ = lambda0_(d_8_v3_, d_7_v2_)
            nw8_ = _dafny.Array(None, 3)
            for i0_0_ in range(nw8_.length(0)):
                nw8_[i0_0_] = init0_(i0_0_)
            d_37_v29_ = nw8_
            d_41_v30_: D2
            d_41_v30_ = D2_DC4(d_37_v29_)
            d_37_v29_ = (d_41_v30_).cf7
            d_42_v31_: _dafny.Array
            nw9_ = _dafny.Array(_dafny.Map({}), 2)
            d_42_v31_ = nw9_
            d_43_v32_: _dafny.Map
            d_43_v32_ = _dafny.Map({d_6_v1_: d_6_v1_})
            index2_ = default__.safeIndex(434, (d_42_v31_).length(0))
            (d_42_v31_)[index2_] = d_43_v32_
            d_44_v33_: _dafny.MultiSet
            d_44_v33_ = _dafny.MultiSet([d_35_v27_])
            index3_ = default__.safeIndex(434, (d_42_v31_).length(0))
            rhs5_ = ((d_44_v33_) - (d_44_v33_)).cardinality
            rhs6_ = d_35_v27_
            rhs7_ = d_43_v32_
            lhs5_ = d_18_globalState_
            lhs6_ = d_42_v31_
            lhs7_ = default__.safeIndex(434, (d_42_v31_).length(0))
            lhs5_.f19 = rhs5_
            d_35_v27_ = rhs6_
            lhs6_[lhs7_] = rhs7_
        d_45_v34_: C0
        nw10_ = C0()
        nw10_.ctor__(d_6_v1_)
        d_45_v34_ = nw10_
        d_46_i2_: int
        d_46_i2_ = 0
        with _dafny.label("0"):
            while d_45_v34_.f26:
                with _dafny.c_label("0"):
                    if (d_46_i2_) >= (100):
                        raise _dafny.Break("0")
                    d_46_i2_ = (d_46_i2_) + (1)
                    d_47_v35_: _dafny.Map
                    d_47_v35_ = _dafny.Map({d_6_v1_: False})
                    d_48_v36_: _dafny.Map
                    d_48_v36_ = _dafny.Map({d_6_v1_: d_47_v35_})
                    d_49_v37_: _dafny.Map
                    d_49_v37_ = _dafny.Map({d_48_v36_: d_6_v1_})
                    d_50_v38_: _dafny.Seq
                    d_50_v38_ = _dafny.SeqWithoutIsStrInference([d_8_v3_])
                    (d_45_v34_).f26 = ((d_49_v37_)[d_48_v36_] if (d_48_v36_) in (d_49_v37_) else default__.fm0(d_50_v38_, d_8_v3_, d_8_v3_, d_18_globalState_))
                    d_51_v39_: _dafny.Array
                    nw11_ = _dafny.Array(_dafny.Seq({}), 12)
                    d_51_v39_ = nw11_
                    d_52_v40_: D2
                    d_52_v40_ = D2_DC4((d_51_v39_ if d_6_v1_ else d_51_v39_))
                    d_52_v40_ = d_52_v40_
                    d_53_v41_: _dafny.Set
                    d_53_v41_ = _dafny.Set({d_6_v1_})
                    d_54_v42_: _dafny.Array
                    nw12_ = _dafny.Array(None, 4)
                    nw12_[int(0)] = (d_53_v41_).intersection(d_53_v41_)
                    nw12_[int(1)] = (_dafny.Set({d_6_v1_})).intersection(d_53_v41_)
                    nw12_[int(2)] = d_53_v41_
                    nw12_[int(3)] = d_53_v41_
                    d_54_v42_ = nw12_
                    index4_ = default__.safeIndex(335, (d_54_v42_).length(0))
                    (d_54_v42_)[index4_] = d_53_v41_
                    index5_ = default__.safeIndex(335, (d_54_v42_).length(0))
                    (d_54_v42_)[index5_] = (d_53_v41_).intersection(_dafny.Set({d_45_v34_.f26, d_45_v34_.f26}))
                    if d_6_v1_:
                        d_47_v35_ = (d_47_v35_).set(d_45_v34_.f26, d_6_v1_)
                        d_55_v43_: _dafny.Map
                        d_55_v43_ = _dafny.Map({d_11_v6_: d_8_v3_})
                        d_56_v44_: _dafny.Map
                        d_56_v44_ = _dafny.Map({d_8_v3_: len(d_50_v38_)})
                        d_55_v43_ = (d_55_v43_).set(d_11_v6_, default__.safeModuloInt(((d_56_v44_)[len(d_50_v38_)] if (len(d_50_v38_)) in (d_56_v44_) else 999), d_8_v3_))
                        d_57_v45_: bool
                        d_58_v46_: _dafny.Seq
                        out2_: bool
                        out3_: _dafny.Seq
                        out2_, out3_ = (d_45_v34_).m0(d_45_v34_.f26, d_8_v3_, d_18_globalState_)
                        d_57_v45_ = out2_
                        d_58_v46_ = out3_
                        d_59_v47_: C0
                        nw13_ = C0()
                        nw13_.ctor__(True)
                        d_59_v47_ = nw13_
                        d_60_v48_: D1
                        d_60_v48_ = D1_DC3(len((d_9_v4_).intersection(d_9_v4_)))
                        d_61_v49_: _dafny.Array
                        nw14_ = _dafny.Array(_dafny.Array(None, 0), 15)
                        d_61_v49_ = nw14_
                        index6_ = default__.safeIndex(501, (d_61_v49_).length(0))
                        (d_61_v49_)[index6_] = d_5_v0_
                        d_62_v50_: _dafny.Seq
                        d_62_v50_ = _dafny.SeqWithoutIsStrInference([d_11_v6_])
                        d_63_v51_: _dafny.Seq
                        d_63_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfhvebs"))
                        d_64_v52_: _dafny.Map
                        d_64_v52_ = _dafny.Map({d_45_v34_.f26: d_63_v51_})
                        index7_ = default__.safeIndex(501, (d_61_v49_).length(0))
                        rhs8_ = d_59_v47_
                        rhs9_ = (D2_DC5(d_5_v0_, d_62_v50_, len(d_64_v52_))).cf10
                        rhs10_ = D1_DC3(d_8_v3_)
                        rhs11_ = d_45_v34_.f26
                        rhs12_ = d_5_v0_
                        lhs8_ = d_45_v34_
                        lhs9_ = d_61_v49_
                        lhs10_ = default__.safeIndex(501, (d_61_v49_).length(0))
                        d_59_v47_ = rhs8_
                        d_8_v3_ = rhs9_
                        d_60_v48_ = rhs10_
                        lhs8_.f26 = rhs11_
                        lhs9_[lhs10_] = rhs12_
                        d_65_v53_: _dafny.Array
                        nw15_ = _dafny.Array(None, 14)
                        nw15_[int(0)] = default__.fm6(d_18_globalState_)
                        nw15_[int(1)] = d_12_v7_
                        nw15_[int(2)] = d_12_v7_
                        nw15_[int(3)] = d_12_v7_
                        nw15_[int(4)] = d_12_v7_
                        nw15_[int(5)] = d_12_v7_
                        nw15_[int(6)] = d_12_v7_
                        nw15_[int(7)] = d_12_v7_
                        nw15_[int(8)] = d_12_v7_
                        nw15_[int(9)] = d_12_v7_
                        nw15_[int(10)] = d_12_v7_
                        nw15_[int(11)] = d_12_v7_
                        nw15_[int(12)] = d_12_v7_
                        nw15_[int(13)] = d_12_v7_
                        d_65_v53_ = nw15_
                        index8_ = default__.safeIndex(923, (d_65_v53_).length(0))
                        (d_65_v53_)[index8_] = d_12_v7_
                        index9_ = default__.safeIndex(923, (d_65_v53_).length(0))
                        (d_65_v53_)[index9_] = d_12_v7_
                    elif True:
                        nw16_ = _dafny.Array(int(0), 21)
                        (d_18_globalState_).f5 = nw16_
                        (d_18_globalState_).f15 = d_8_v3_
                        d_66_v54_: _dafny.Map
                        d_66_v54_ = _dafny.Map({D2_DC6(): d_12_v7_})
                        d_67_v55_: _dafny.Map
                        d_67_v55_ = _dafny.Map({d_8_v3_: d_66_v54_})
                        (d_18_globalState_).f19 = ((len(d_67_v55_)) - (d_8_v3_)) * (d_8_v3_)
                        d_7_v2_ = (d_7_v2_).set(default__.safeIndex(d_8_v3_, len(d_7_v2_)), d_45_v34_.f26)
                        d_12_v7_ = d_12_v7_
                    pass
            pass
        d_68_v56_: _dafny.Array
        def lambda2_(d_69_i3_):
            return D2_DC6()

        init1_ = lambda2_
        nw17_ = _dafny.Array(None, 7)
        for i0_1_ in range(nw17_.length(0)):
            nw17_[i0_1_] = init1_(i0_1_)
        d_68_v56_ = nw17_
        d_70_v57_: _dafny.Array
        nw18_ = _dafny.Array(None, 25)
        nw18_[int(0)] = d_68_v56_
        nw18_[int(1)] = d_68_v56_
        nw18_[int(2)] = d_68_v56_
        nw18_[int(3)] = (d_68_v56_ if d_45_v34_.f26 else d_68_v56_)
        nw18_[int(4)] = d_68_v56_
        nw18_[int(5)] = (d_68_v56_ if not(d_45_v34_.f26) else d_68_v56_)
        nw18_[int(6)] = d_68_v56_
        nw18_[int(7)] = d_68_v56_
        nw18_[int(8)] = d_68_v56_
        nw18_[int(9)] = d_68_v56_
        nw18_[int(10)] = d_68_v56_
        nw18_[int(11)] = d_68_v56_
        nw18_[int(12)] = d_68_v56_
        nw18_[int(13)] = d_68_v56_
        nw18_[int(14)] = (d_68_v56_ if d_6_v1_ else d_68_v56_)
        nw18_[int(15)] = d_68_v56_
        nw18_[int(16)] = d_68_v56_
        nw18_[int(17)] = d_68_v56_
        nw18_[int(18)] = d_68_v56_
        nw18_[int(19)] = d_68_v56_
        nw18_[int(20)] = d_68_v56_
        nw18_[int(21)] = d_68_v56_
        nw18_[int(22)] = d_68_v56_
        nw18_[int(23)] = d_68_v56_
        nw18_[int(24)] = d_68_v56_
        d_70_v57_ = nw18_
        d_70_v57_ = (d_70_v57_ if d_45_v34_.f26 else d_70_v57_)
        d_71_v58_: _dafny.Seq
        d_71_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xseuq"))
        index10_ = default__.safeIndex(126, (d_5_v0_).length(0))
        (d_5_v0_)[index10_] = len(d_71_v58_)
        index11_ = default__.safeIndex(126, (d_5_v0_).length(0))
        (d_5_v0_)[index11_] = (0) - (d_8_v3_)
        hi0_ = d_8_v3_
        for d_72_i4_ in range(210, hi0_):
            d_73_v59_: bool
            d_74_v60_: _dafny.Seq
            out4_: bool
            out5_: _dafny.Seq
            out4_, out5_ = (d_45_v34_).m0(d_45_v34_.f26, ((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]) * (d_8_v3_), d_18_globalState_)
            d_73_v59_ = out4_
            d_74_v60_ = out5_
            d_75_v61_: _dafny.Map
            d_75_v61_ = _dafny.Map({_dafny.CodePoint('v'): d_9_v4_})
            d_75_v61_ = (d_75_v61_).set(d_12_v7_, d_9_v4_)
            d_76_v62_: _dafny.Seq
            d_76_v62_ = _dafny.SeqWithoutIsStrInference([d_71_v58_])
            if (d_71_v58_) != (((d_76_v62_)[default__.safeIndex((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))], len(d_76_v62_))]) + (d_71_v58_)):
                (d_18_globalState_).f5 = d_5_v0_
                d_77_v63_: _dafny.Array
                nw19_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_77_v63_ = nw19_
                d_78_v64_: _dafny.Map
                d_78_v64_ = _dafny.Map({d_77_v63_: d_5_v0_})
                d_78_v64_ = (d_78_v64_).set(d_77_v63_, d_5_v0_)
                d_79_v65_: _dafny.Seq
                d_79_v65_ = _dafny.SeqWithoutIsStrInference([d_11_v6_])
                d_80_v66_: D2
                d_80_v66_ = D2_DC5(d_5_v0_, d_79_v65_, (0) - ((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]))
                pat_let_tv0_ = d_5_v0_
                d_81_v67_: _dafny.Map
                def iife7_(_pat_let0_0):
                    def iife8_(d_82_dt__update__tmp_h0_):
                        def iife9_(_pat_let1_0):
                            def iife10_(d_83_dt__update_hcf10_h0_):
                                def iife11_(_pat_let2_0):
                                    def iife12_(d_84_dt__update_hcf8_h0_):
                                        return D2_DC5(d_84_dt__update_hcf8_h0_, (d_82_dt__update__tmp_h0_).cf9, d_83_dt__update_hcf10_h0_)
                                    return iife12_(_pat_let2_0)
                                return iife11_(pat_let_tv0_)
                            return iife10_(_pat_let1_0)
                        return iife9_(762)
                    return iife8_(_pat_let0_0)
                d_81_v67_ = _dafny.Map({d_45_v34_.f26: iife7_(d_80_v66_)})
                d_81_v67_ = d_81_v67_
                d_85_v68_: bool
                d_86_v69_: _dafny.Seq
                out6_: bool
                out7_: _dafny.Seq
                out6_, out7_ = (d_45_v34_).m0((826) < (d_72_i4_), 143, d_18_globalState_)
                d_85_v68_ = out6_
                d_86_v69_ = out7_
                (d_45_v34_).f26 = d_45_v34_.f26
            elif True:
                d_87_v70_: _dafny.Map
                d_87_v70_ = _dafny.Map({(d_45_v34_.f26) or (d_6_v1_): (d_71_v58_) < (d_71_v58_)})
                d_87_v70_ = (d_87_v70_).set(d_45_v34_.f26, not (d_73_v59_) or (d_45_v34_.f26))
                d_6_v1_ = not (d_6_v1_) or (d_45_v34_.f26)
                d_88_v71_: bool
                d_89_v72_: _dafny.Seq
                out8_: bool
                out9_: _dafny.Seq
                out8_, out9_ = (d_45_v34_).m0(d_73_v59_, len(default__.fm7(d_18_globalState_)), d_18_globalState_)
                d_88_v71_ = out8_
                d_89_v72_ = out9_
                d_90_v73_: C0
                nw20_ = C0()
                nw20_.ctor__(d_88_v71_)
                d_90_v73_ = nw20_
                d_90_v73_ = d_45_v34_
                d_91_v74_: bool
                d_92_v75_: _dafny.Seq
                out10_: bool
                out11_: _dafny.Seq
                out10_, out11_ = (d_90_v73_).m0(True, d_72_i4_, d_18_globalState_)
                d_91_v74_ = out10_
                d_92_v75_ = out11_
            d_93_v76_: _dafny.Seq
            d_93_v76_ = _dafny.SeqWithoutIsStrInference([d_7_v2_, d_7_v2_])
            d_94_v77_: D3
            d_94_v77_ = D3_DC7(d_7_v2_)
            d_95_v78_: _dafny.MultiSet
            d_95_v78_ = _dafny.MultiSet([(d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]])
            d_96_v79_: _dafny.Map
            d_96_v79_ = _dafny.Map({d_95_v78_: _dafny.SeqWithoutIsStrInference([d_6_v1_])})
            d_97_v80_: _dafny.Array
            nw21_ = _dafny.Array(None, 28)
            nw21_[int(0)] = d_7_v2_
            nw21_[int(1)] = d_7_v2_
            nw21_[int(2)] = d_7_v2_
            nw21_[int(3)] = (d_7_v2_) + (d_7_v2_)
            nw21_[int(4)] = _dafny.SeqWithoutIsStrInference([d_45_v34_.f26, d_6_v1_, True, d_73_v59_, not(d_45_v34_.f26)])
            nw21_[int(5)] = d_7_v2_
            nw21_[int(6)] = (d_7_v2_ if True else d_7_v2_)
            nw21_[int(7)] = (d_7_v2_) + ((d_93_v76_)[default__.safeIndex(d_8_v3_, len(d_93_v76_))])
            nw21_[int(8)] = (d_7_v2_) + (d_7_v2_)
            nw21_[int(9)] = d_7_v2_
            nw21_[int(10)] = (_dafny.SeqWithoutIsStrInference([d_45_v34_.f26])) + (_dafny.SeqWithoutIsStrInference([d_45_v34_.f26]))
            nw21_[int(11)] = d_7_v2_
            nw21_[int(12)] = (_dafny.SeqWithoutIsStrInference([d_6_v1_, d_45_v34_.f26])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))] for d_98_i5_ in range(default__.abs(574))])), len(_dafny.SeqWithoutIsStrInference([d_6_v1_, d_45_v34_.f26]))), d_45_v34_.f26)
            nw21_[int(13)] = (d_94_v77_).cf11
            nw21_[int(14)] = d_7_v2_
            nw21_[int(15)] = (d_7_v2_).set(default__.safeIndex((0) - (len((d_45_v34_).fm1(d_73_v59_, d_18_globalState_))), len(d_7_v2_)), d_45_v34_.f26)
            nw21_[int(16)] = d_7_v2_
            nw21_[int(17)] = d_7_v2_
            nw21_[int(18)] = (d_7_v2_ if d_6_v1_ else d_7_v2_)
            nw21_[int(19)] = d_7_v2_
            nw21_[int(20)] = d_7_v2_
            nw21_[int(21)] = (d_93_v76_)[default__.safeIndex(-260, len(d_93_v76_))]
            nw21_[int(22)] = (d_7_v2_) + (d_7_v2_)
            nw21_[int(23)] = d_7_v2_
            nw21_[int(24)] = (d_7_v2_) + (d_7_v2_)
            nw21_[int(25)] = ((d_96_v79_)[_dafny.MultiSet([d_8_v3_])] if (_dafny.MultiSet([d_8_v3_])) in (d_96_v79_) else _dafny.SeqWithoutIsStrInference([d_45_v34_.f26, d_6_v1_]))
            nw21_[int(26)] = d_7_v2_
            nw21_[int(27)] = d_7_v2_
            d_97_v80_ = nw21_
            index12_ = default__.safeIndex(414, (d_97_v80_).length(0))
            (d_97_v80_)[index12_] = ((d_7_v2_).set(default__.safeIndex(d_8_v3_, len(d_7_v2_)), d_45_v34_.f26)) + (d_7_v2_)
            index13_ = default__.safeIndex(414, (d_97_v80_).length(0))
            (d_97_v80_)[index13_] = _dafny.SeqWithoutIsStrInference([not(False), d_6_v1_, d_45_v34_.f26])
        if d_6_v1_:
            if (d_10_v5_):
                (d_45_v34_).f26 = d_45_v34_.f26
                d_99_v81_: bool
                d_100_v82_: _dafny.Seq
                out12_: bool
                out13_: _dafny.Seq
                out12_, out13_ = (d_45_v34_).m0(d_45_v34_.f26, (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))], d_18_globalState_)
                d_99_v81_ = out12_
                d_100_v82_ = out13_
                (d_18_globalState_).f15 = (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]
                (d_18_globalState_).f15 = 398
                (d_45_v34_).f26 = d_99_v81_
            elif True:
                d_101_v83_: bool
                d_102_v84_: _dafny.Seq
                out14_: bool
                out15_: _dafny.Seq
                out14_, out15_ = (d_45_v34_).m0(d_45_v34_.f26, d_8_v3_, d_18_globalState_)
                d_101_v83_ = out14_
                d_102_v84_ = out15_
                index14_ = default__.safeIndex(320, (d_11_v6_).length(0))
                (d_11_v6_)[index14_] = (d_12_v7_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqoumw")))
                d_103_v85_: _dafny.Array
                nw22_ = _dafny.Array(None, 5)
                nw22_[int(0)] = d_12_v7_
                nw22_[int(1)] = d_12_v7_
                nw22_[int(2)] = d_12_v7_
                nw22_[int(3)] = d_12_v7_
                nw22_[int(4)] = d_12_v7_
                d_103_v85_ = nw22_
                d_104_v86_: _dafny.Map
                d_104_v86_ = _dafny.Map({d_12_v7_: d_103_v85_})
                d_105_v87_: _dafny.Array
                nw23_ = _dafny.Array(None, 8)
                nw23_[int(0)] = (d_103_v85_ if d_45_v34_.f26 else d_103_v85_)
                nw23_[int(1)] = d_103_v85_
                nw23_[int(2)] = d_103_v85_
                nw23_[int(3)] = d_103_v85_
                nw23_[int(4)] = d_103_v85_
                nw23_[int(5)] = ((d_104_v86_)[d_12_v7_] if (d_12_v7_) in (d_104_v86_) else d_103_v85_)
                nw23_[int(6)] = d_103_v85_
                nw23_[int(7)] = d_103_v85_
                d_105_v87_ = nw23_
                d_106_v88_: D3
                d_106_v88_ = D3_DC9((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))])
                d_107_v89_: D3
                d_107_v89_ = D3_DC10(d_106_v88_)
                index15_ = default__.safeIndex(320, (d_11_v6_).length(0))
                index16_ = default__.safeIndex(126, (d_5_v0_).length(0))
                rhs13_ = (d_10_v5_)
                rhs14_ = d_105_v87_
                rhs15_ = (d_107_v89_ if not(d_45_v34_.f26) else d_107_v89_)
                rhs16_ = (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]
                lhs11_ = d_11_v6_
                lhs12_ = default__.safeIndex(320, (d_11_v6_).length(0))
                lhs13_ = d_5_v0_
                lhs14_ = default__.safeIndex(126, (d_5_v0_).length(0))
                lhs11_[lhs12_] = rhs13_
                d_105_v87_ = rhs14_
                d_107_v89_ = rhs15_
                lhs13_[lhs14_] = rhs16_
                index17_ = default__.safeIndex(320, (d_11_v6_).length(0))
                rhs17_ = (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]
                rhs18_ = default__.fm0(_dafny.SeqWithoutIsStrInference([d_8_v3_ for d_108_i6_ in range(default__.abs(-879))]), default__.safeDivisionInt((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))], (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]), (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))], d_18_globalState_)
                lhs15_ = d_18_globalState_
                lhs16_ = d_11_v6_
                lhs17_ = default__.safeIndex(320, (d_11_v6_).length(0))
                lhs15_.f19 = rhs17_
                lhs16_[lhs17_] = rhs18_
                index18_ = default__.safeIndex(126, (d_5_v0_).length(0))
                (d_5_v0_)[index18_] = (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]
                d_109_v90_: _dafny.Seq
                d_109_v90_ = _dafny.SeqWithoutIsStrInference([d_45_v34_, d_45_v34_])
                d_110_v91_: _dafny.Seq
                d_110_v91_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_109_v90_))])
                d_111_v92_: _dafny.Map
                d_111_v92_ = _dafny.Map({d_45_v34_.f26: (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]})
                d_112_v93_: _dafny.Map
                d_112_v93_ = _dafny.Map({d_8_v3_: d_111_v92_})
                d_113_v94_: _dafny.MultiSet
                d_113_v94_ = _dafny.MultiSet([default__.fm0(d_110_v91_, (0) - (d_8_v3_), d_8_v3_, d_18_globalState_), (d_45_v34_.f26) not in (((d_112_v93_)[(d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]] if ((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]) in (d_112_v93_) else d_111_v92_))])
                d_114_v95_: _dafny.Map
                d_114_v95_ = _dafny.Map({(d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]: (d_11_v6_)[default__.safeIndex(320, (d_11_v6_).length(0))]})
                d_115_v96_: _dafny.Map
                d_115_v96_ = _dafny.Map({d_12_v7_: _dafny.MultiSet([(d_11_v6_)[default__.safeIndex(320, (d_11_v6_).length(0))], (d_11_v6_)[default__.safeIndex(320, (d_11_v6_).length(0))]])})
                index19_ = default__.safeIndex(126, (d_5_v0_).length(0))
                rhs19_ = (default__.fm3(82, d_114_v95_, (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))], d_18_globalState_)) + (default__.safeDivisionInt(d_8_v3_, (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]))
                rhs20_ = d_12_v7_
                rhs21_ = (d_11_v6_)[default__.safeIndex(320, (d_11_v6_).length(0))]
                rhs22_ = ((d_115_v96_)[d_12_v7_] if (d_12_v7_) in (d_115_v96_) else d_113_v94_)
                lhs18_ = d_5_v0_
                lhs19_ = default__.safeIndex(126, (d_5_v0_).length(0))
                lhs20_ = d_45_v34_
                lhs18_[lhs19_] = rhs19_
                d_12_v7_ = rhs20_
                lhs20_.f26 = rhs21_
                d_113_v94_ = rhs22_
            if d_6_v1_:
                index20_ = default__.safeIndex(126, (d_5_v0_).length(0))
                (d_5_v0_)[index20_] = d_8_v3_
                d_116_v97_: _dafny.Seq
                d_116_v97_ = _dafny.SeqWithoutIsStrInference([(d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]])
                d_117_v98_: _dafny.Map
                d_117_v98_ = _dafny.Map({(d_8_v3_) not in (d_116_v97_): d_45_v34_.f26})
                d_117_v98_ = (d_117_v98_).set(d_45_v34_.f26, not(d_45_v34_.f26))
                d_118_v99_: _dafny.Set
                d_118_v99_ = _dafny.Set({d_45_v34_.f26})
                d_118_v99_ = d_118_v99_
                d_119_v100_: C0
                nw24_ = C0()
                nw24_.ctor__(False)
                d_119_v100_ = nw24_
                d_120_v101_: _dafny.Seq
                d_120_v101_ = _dafny.SeqWithoutIsStrInference([d_119_v100_, d_119_v100_])
                d_121_v102_: D4
                d_121_v102_ = D4_DC11((d_120_v101_)[default__.safeIndex((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))], len(d_120_v101_))])
                d_119_v100_ = (d_121_v102_).cf14
                d_122_v103_: _dafny.Array
                nw25_ = _dafny.Array(_dafny.CodePoint('D'), 19)
                d_122_v103_ = nw25_
                d_123_v104_: _dafny.Map
                d_123_v104_ = _dafny.Map({d_5_v0_: d_122_v103_})
                d_124_v105_: _dafny.Array
                nw26_ = _dafny.Array(None, 14)
                nw26_[int(0)] = d_122_v103_
                nw26_[int(1)] = d_122_v103_
                nw26_[int(2)] = d_122_v103_
                nw26_[int(3)] = d_122_v103_
                nw26_[int(4)] = d_122_v103_
                nw26_[int(5)] = d_122_v103_
                nw26_[int(6)] = d_122_v103_
                nw26_[int(7)] = d_122_v103_
                nw26_[int(8)] = d_122_v103_
                nw26_[int(9)] = ((d_123_v104_)[d_5_v0_] if (d_5_v0_) in (d_123_v104_) else d_122_v103_)
                nw26_[int(10)] = d_122_v103_
                nw26_[int(11)] = d_122_v103_
                nw26_[int(12)] = d_122_v103_
                nw26_[int(13)] = d_122_v103_
                d_124_v105_ = nw26_
                d_125_v106_: _dafny.Map
                d_125_v106_ = _dafny.Map({d_7_v2_: d_124_v105_})
                d_125_v106_ = (d_125_v106_).set((d_7_v2_) + (d_7_v2_), d_124_v105_)
            elif True:
                d_126_v107_: C0
                nw27_ = C0()
                nw27_.ctor__(True)
                d_126_v107_ = nw27_
                (d_45_v34_).f26 = True
                d_12_v7_ = d_12_v7_
                d_127_v108_: _dafny.Map
                d_127_v108_ = _dafny.Map({d_126_v107_: d_12_v7_})
                d_128_v109_: _dafny.Seq
                d_128_v109_ = _dafny.SeqWithoutIsStrInference([len(d_127_v108_), 254])
                d_129_v110_: _dafny.Map
                d_129_v110_ = _dafny.Map({(0) - (len(_dafny.Map({d_45_v34_.f26: d_45_v34_.f26}))): d_45_v34_.f26})
                d_130_v111_: _dafny.Map
                d_130_v111_ = _dafny.Map({d_129_v110_: len(d_71_v58_)})
                rhs23_ = d_128_v109_
                rhs24_ = (len(d_130_v111_)) + (((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]) - ((_dafny.MultiSet([d_126_v107_.f26])).cardinality))
                lhs21_ = d_18_globalState_
                d_128_v109_ = rhs23_
                lhs21_.f19 = rhs24_
                d_131_v112_: _dafny.Seq
                d_131_v112_ = _dafny.SeqWithoutIsStrInference([d_71_v58_])
                d_132_v113_: _dafny.MultiSet
                d_132_v113_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))])
                d_133_v114_: _dafny.Array
                nw28_ = _dafny.Array(None, 7)
                nw28_[int(0)] = (_dafny.MultiSet([d_71_v58_, d_71_v58_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbpcgr"))])).intersection(_dafny.MultiSet(d_131_v112_))
                nw28_[int(1)] = d_132_v113_
                nw28_[int(2)] = (d_132_v113_) | (default__.fm8(d_18_globalState_))
                nw28_[int(3)] = (_dafny.MultiSet([d_71_v58_])).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihs")), default__.abs(d_8_v3_))
                nw28_[int(4)] = default__.fm8(d_18_globalState_)
                nw28_[int(5)] = d_132_v113_
                nw28_[int(6)] = (_dafny.MultiSet([d_71_v58_, d_71_v58_])) | (d_132_v113_)
                d_133_v114_ = nw28_
                index21_ = default__.safeIndex(223, (d_133_v114_).length(0))
                (d_133_v114_)[index21_] = d_132_v113_
                index22_ = default__.safeIndex(223, (d_133_v114_).length(0))
                (d_133_v114_)[index22_] = d_132_v113_
            d_11_v6_ = d_11_v6_
            (d_18_globalState_).f15 = (906) + (758)
            d_134_v115_: _dafny.Map
            d_134_v115_ = _dafny.Map({D2_DC5(d_5_v0_, _dafny.SeqWithoutIsStrInference([d_11_v6_, d_11_v6_, d_11_v6_, d_11_v6_]), len(d_7_v2_)): d_71_v58_})
            d_134_v115_ = d_134_v115_
        elif True:
            d_135_v116_: _dafny.Seq
            d_135_v116_ = _dafny.SeqWithoutIsStrInference([d_11_v6_, d_11_v6_, d_11_v6_, d_11_v6_])
            d_136_v117_: D2
            d_136_v117_ = D2_DC5(d_5_v0_, d_135_v116_, (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))])
            d_137_v118_: _dafny.Seq
            d_137_v118_ = _dafny.SeqWithoutIsStrInference([d_135_v116_])
            d_136_v117_ = D2_DC5(d_5_v0_, (d_137_v118_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_5_v0_, d_5_v0_, d_5_v0_])), len(d_137_v118_))], (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))])
            d_138_v119_: _dafny.Seq
            d_138_v119_ = _dafny.SeqWithoutIsStrInference([195])
            d_139_v120_: _dafny.Map
            d_139_v120_ = _dafny.Map({(d_7_v2_)[default__.safeIndex((d_138_v119_)[default__.safeIndex(d_8_v3_, len(d_138_v119_))], len(d_7_v2_))]: d_8_v3_})
            d_140_v121_: _dafny.MultiSet
            d_140_v121_ = _dafny.MultiSet([(d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]])
            d_141_v122_: _dafny.Map
            d_141_v122_ = _dafny.Map({(d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]: (d_140_v121_).cardinality})
            d_142_v123_: _dafny.MultiSet
            d_142_v123_ = _dafny.MultiSet([d_141_v122_])
            d_139_v120_ = (d_139_v120_).set(d_45_v34_.f26, (d_142_v123_).cardinality)
            d_143_v124_: _dafny.Map
            d_143_v124_ = _dafny.Map({not(d_6_v1_): d_45_v34_.f26})
            d_143_v124_ = (d_143_v124_).set(d_45_v34_.f26, (d_7_v2_)[default__.safeIndex(d_8_v3_, len(d_7_v2_))])
            d_144_v125_: C0
            nw29_ = C0()
            nw29_.ctor__((d_45_v34_.f26) == (d_6_v1_))
            d_144_v125_ = nw29_
            d_145_v126_: _dafny.Set
            d_145_v126_ = _dafny.Set({True})
            d_145_v126_ = (d_145_v126_) | (d_145_v126_)
        d_146_v127_: _dafny.Map
        d_146_v127_ = _dafny.Map({d_5_v0_: len(d_7_v2_)})
        (d_45_v34_).f26 = (default__.safeDivisionInt(len(_dafny.Map({False: d_6_v1_})), (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))])) != (((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]) - ((0) - (len(d_146_v127_))))
        (d_18_globalState_).f19 = (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]
        d_147_v128_: _dafny.MultiSet
        d_147_v128_ = _dafny.MultiSet([d_45_v34_.f26, d_6_v1_])
        d_148_v129_: _dafny.Map
        d_148_v129_ = _dafny.Map({len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stmah"))) + (d_71_v58_)): not((d_147_v128_).ispropersubset(d_147_v128_))})
        d_148_v129_ = (d_148_v129_).set(d_8_v3_, d_6_v1_)
        d_149_v130_: _dafny.Array
        nw30_ = _dafny.Array(_dafny.MultiSet({}), 11)
        d_149_v130_ = nw30_
        d_150_v131_: D4
        d_150_v131_ = D4_DC12(d_149_v130_, d_11_v6_, d_6_v1_)
        d_151_v132_: D4
        d_151_v132_ = D4_DC13(d_150_v131_)
        d_151_v132_ = D4_DC13(d_150_v131_)
        d_152_v133_: C0
        nw31_ = C0()
        nw31_.ctor__(d_6_v1_)
        d_152_v133_ = nw31_
        index23_ = default__.safeIndex(126, (d_5_v0_).length(0))
        rhs25_ = 872
        rhs26_ = (d_7_v2_) + (d_7_v2_)
        rhs27_ = ((d_8_v3_) * (165)) * ((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))])
        rhs28_ = _dafny.SeqWithoutIsStrInference([d_6_v1_, d_152_v133_.f26, d_45_v34_.f26])
        lhs22_ = d_5_v0_
        lhs23_ = default__.safeIndex(126, (d_5_v0_).length(0))
        lhs24_ = d_18_globalState_
        lhs22_[lhs23_] = rhs25_
        d_7_v2_ = rhs26_
        lhs24_.f19 = rhs27_
        d_7_v2_ = rhs28_
        if d_45_v34_.f26:
            (d_45_v34_).f26 = d_45_v34_.f26
            d_153_v134_: C0
            nw32_ = C0()
            nw32_.ctor__(d_45_v34_.f26)
            d_153_v134_ = nw32_
            index24_ = default__.safeIndex(126, (d_5_v0_).length(0))
            (d_5_v0_)[index24_] = default__.safeModuloInt((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))], d_8_v3_)
            d_154_v135_: _dafny.Map
            d_154_v135_ = _dafny.Map({d_45_v34_.f26: d_153_v134_.f26})
            d_155_v136_: bool
            d_156_v137_: _dafny.Seq
            out16_: bool
            out17_: _dafny.Seq
            out16_, out17_ = (d_153_v134_).m0((d_6_v1_) == ((d_7_v2_)[default__.safeIndex(d_8_v3_, len(d_7_v2_))]), len(d_154_v135_), d_18_globalState_)
            d_155_v136_ = out16_
            d_156_v137_ = out17_
            d_8_v3_ = d_8_v3_
        elif True:
            d_157_v138_: bool
            d_158_v139_: _dafny.Seq
            out18_: bool
            out19_: _dafny.Seq
            out18_, out19_ = (d_152_v133_).m0(d_45_v34_.f26, d_8_v3_, d_18_globalState_)
            d_157_v138_ = out18_
            d_158_v139_ = out19_
            index25_ = default__.safeIndex(126, (d_5_v0_).length(0))
            (d_5_v0_)[index25_] = d_8_v3_
            d_152_v133_ = d_152_v133_
            d_159_v140_: D5
            d_159_v140_ = D5_DC14(d_148_v129_)
            d_160_v141_: _dafny.MultiSet
            d_160_v141_ = _dafny.MultiSet([d_71_v58_])
            d_161_v142_: _dafny.Map
            d_161_v142_ = _dafny.Map({(d_71_v58_).set(default__.safeIndex(default__.fm3(d_8_v3_, (d_159_v140_).cf19, d_8_v3_, d_18_globalState_), len(d_71_v58_)), d_12_v7_): (d_160_v141_).cardinality})
            d_161_v142_ = (d_161_v142_).set(d_71_v58_, (d_8_v3_) + (d_8_v3_))
            if (not(d_45_v34_.f26)) == (not(d_152_v133_.f26)):
                index26_ = default__.safeIndex(126, (d_5_v0_).length(0))
                (d_5_v0_)[index26_] = (0) - (d_8_v3_)
                (d_18_globalState_).f5 = d_5_v0_
                d_12_v7_ = ((d_15_v10_)[d_45_v34_.f26] if (d_45_v34_.f26) in (d_15_v10_) else d_12_v7_)
                d_162_v143_: _dafny.Map
                d_162_v143_ = _dafny.Map({d_12_v7_: d_5_v0_})
                d_163_v144_: _dafny.Array
                nw33_ = _dafny.Array(None, 13)
                nw33_[int(0)] = ((d_162_v143_)[d_12_v7_] if (d_12_v7_) in (d_162_v143_) else d_5_v0_)
                nw33_[int(1)] = d_5_v0_
                nw33_[int(2)] = d_5_v0_
                nw33_[int(3)] = d_5_v0_
                nw33_[int(4)] = d_5_v0_
                nw33_[int(5)] = d_5_v0_
                nw33_[int(6)] = d_5_v0_
                nw33_[int(7)] = d_5_v0_
                nw33_[int(8)] = d_5_v0_
                nw33_[int(9)] = d_5_v0_
                nw33_[int(10)] = d_5_v0_
                nw33_[int(11)] = d_5_v0_
                nw33_[int(12)] = d_5_v0_
                d_163_v144_ = nw33_
                d_163_v144_ = d_163_v144_
                d_164_v145_: bool
                d_165_v146_: _dafny.Seq
                out20_: bool
                out21_: _dafny.Seq
                out20_, out21_ = (d_45_v34_).m0(d_45_v34_.f26, default__.safeModuloInt(d_8_v3_, d_8_v3_), d_18_globalState_)
                d_164_v145_ = out20_
                d_165_v146_ = out21_
            elif True:
                d_157_v138_ = ((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]) > ((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))])
                d_166_v147_: D2
                d_166_v147_ = D2_DC6()
                d_166_v147_ = d_166_v147_
                (d_152_v133_).f26 = d_45_v34_.f26
                d_167_v148_: bool
                d_168_v149_: _dafny.Seq
                out22_: bool
                out23_: _dafny.Seq
                out22_, out23_ = (d_45_v34_).m0(d_45_v34_.f26, d_8_v3_, d_18_globalState_)
                d_167_v148_ = out22_
                d_168_v149_ = out23_
                d_169_v150_: _dafny.Array
                def lambda3_(d_170_v58_):
                    def lambda4_(d_171_i7_):
                        return d_170_v58_

                    return lambda4_

                init2_ = lambda3_(d_71_v58_)
                nw34_ = _dafny.Array(None, 15)
                for i0_2_ in range(nw34_.length(0)):
                    nw34_[i0_2_] = init2_(i0_2_)
                d_169_v150_ = nw34_
                index27_ = default__.safeIndex(269, (d_169_v150_).length(0))
                (d_169_v150_)[index27_] = d_71_v58_
                index28_ = default__.safeIndex(269, (d_169_v150_).length(0))
                (d_169_v150_)[index28_] = ((d_71_v58_ if d_45_v34_.f26 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "faxxuf")))) + (d_71_v58_)
        _ingredientsd_0 = []
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_14_v9_).length(0)):
            d_172_i8_: int = guard_loop_0_
            if (True) and (((0) <= (d_172_i8_)) and ((d_172_i8_) < ((d_14_v9_).length(0)))):
                _ingredientsd_0.append((d_14_v9_, int(d_172_i8_), _dafny.Set({((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]) - ((d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]), len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-379]), _dafny.SeqWithoutIsStrInference([212 for d_173_i9_ in range(default__.abs(-899))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))], d_8_v3_, (d_5_v0_)[default__.safeIndex(126, (d_5_v0_).length(0))]])]))), 934, (0) - (d_8_v3_)})))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        d_6_v1_ = not(d_45_v34_.f26)
        _dafny.print(_dafny.string_of((d_5_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_6_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_7_v2_) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_8_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_9_v4_) == (_dafny.Set({861}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_10_v5_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_11_v6_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_12_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_13_v8_) == (_dafny.Map({861: _dafny.CodePoint('l')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_14_v9_)[0]) == (_dafny.Set({0, 3, 934, -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_14_v9_)[1]) == (_dafny.Set({0, 3, 934, -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_14_v9_)[2]) == (_dafny.Set({0, 3, 934, -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_14_v9_)[3]) == (_dafny.Set({0, 3, 934, -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_14_v9_)[4]) == (_dafny.Set({0, 3, 934, -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_15_v10_) == (_dafny.Map({True: _dafny.CodePoint('l')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_16_v11_) == (_dafny.MultiSet([_dafny.Map({True: _dafny.CodePoint('l')}), _dafny.Map({True: _dafny.CodePoint('l')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_17_v12_).cf1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f0)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f1) == (_dafny.SeqWithoutIsStrInference([False, True, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f3) == (_dafny.Set({861}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f4)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_.f5)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f6) == (_dafny.Map({861: _dafny.CodePoint('l')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_18_globalState_).f9)[0]) == (_dafny.Set({0, 3, 934, -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_18_globalState_).f9)[1]) == (_dafny.Set({0, 3, 934, -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_18_globalState_).f9)[2]) == (_dafny.Set({0, 3, 934, -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_18_globalState_).f9)[3]) == (_dafny.Set({0, 3, 934, -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_18_globalState_).f9)[4]) == (_dafny.Set({0, 3, 934, -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f12)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_18_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_18_globalState_.f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_.f20) == (_dafny.MultiSet([_dafny.Map({True: _dafny.CodePoint('l')}), _dafny.Map({True: _dafny.CodePoint('l')}), _dafny.Map({True: _dafny.CodePoint('l')}), _dafny.Map({True: _dafny.CodePoint('l')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f22)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_18_globalState_).f23)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_18_globalState_).f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_45_v34_.f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_46_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_71_v58_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_146_v127_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v128_) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_v129_) == (_dafny.Map({10: True, 1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_150_v131_).cf16)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v131_).cf17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_151_v132_).cf18).cf16)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_151_v132_).cf18).cf17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_152_v133_.f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: False
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

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
        return lambda: D1_DC2(int(0), int(0), int(0), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5(_dafny.Array(None, 0), _dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC5(D2, NamedTuple('DC5', [('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC8(D3, NamedTuple('DC8', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(_dafny.Array(None, 0), _dafny.Array(None, 0), False)
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

class D4_DC12(D4, NamedTuple('DC12', [('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC15(None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC15(D5, NamedTuple('DC15', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f5: _dafny.Array = _dafny.Array(None, 0)
        self.f15: int = int(0)
        self.f19: int = int(0)
        self.f20: _dafny.MultiSet = _dafny.MultiSet({})
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: _dafny.Seq = _dafny.Seq({})
        self._f2: bool = False
        self._f3: _dafny.Set = _dafny.Set({})
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        self._f6: _dafny.Map = _dafny.Map({})
        self._f7: bool = False
        self._f8: bool = False
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f10: bool = False
        self._f11: str = _dafny.CodePoint('D')
        self._f12: _dafny.Array = _dafny.Array(None, 0)
        self._f13: int = int(0)
        self._f14: bool = False
        self._f16: bool = False
        self._f17: int = int(0)
        self._f18: int = int(0)
        self._f21: bool = False
        self._f22: _dafny.Array = _dafny.Array(None, 0)
        self._f23: _dafny.Array = _dafny.Array(None, 0)
        self._f24: int = int(0)
        self._f25: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
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
        (self)._f18 = f18
        (self).f19 = f19
        (self).f20 = f20
        (self)._f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23
        (self)._f24 = f24
        (self)._f25 = f25

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
    def f4(self):
        return self._f4
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
    def f18(self):
        return self._f18
    @property
    def f21(self):
        return self._f21
    @property
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25

class C0:
    def  __init__(self):
        self.f26: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f26):
        (self).f26 = f26

    def fm1(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yim"))

    def m0(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.Seq({})
        d_174_v0_: _dafny.Seq
        d_174_v0_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        d_175_v1_: bool
        d_175_v1_ = p0
        d_176_v2_: _dafny.Seq
        d_176_v2_ = _dafny.SeqWithoutIsStrInference([d_174_v0_, d_174_v0_, default__.fm2((d_175_v1_), globalState)])
        d_177_v3_: _dafny.Map
        d_177_v3_ = _dafny.Map({(p1) < (p1): (d_176_v2_)[default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_178_i0_ in range(default__.abs(532))]))), len(d_176_v2_))]})
        d_179_v4_: _dafny.Array
        nw35_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
        d_179_v4_ = nw35_
        d_180_v5_: _dafny.Seq
        d_180_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylwodeou"))
        index29_ = default__.safeIndex(306, (d_179_v4_).length(0))
        (d_179_v4_)[index29_] = d_180_v5_
        d_181_v6_: _dafny.Seq
        d_181_v6_ = _dafny.SeqWithoutIsStrInference([self.f26, not(self.f26), self.f26])
        index30_ = default__.safeIndex(306, (d_179_v4_).length(0))
        rhs29_ = len((d_174_v0_) + ((d_174_v0_) + (d_174_v0_)))
        rhs30_ = (d_181_v6_)[default__.safeIndex(p1, len(d_181_v6_))]
        rhs31_ = (d_177_v3_) | (_dafny.Map({self.f26: d_174_v0_}))
        rhs32_ = d_180_v5_
        lhs25_ = globalState
        lhs26_ = d_179_v4_
        lhs27_ = default__.safeIndex(306, (d_179_v4_).length(0))
        lhs25_.f15 = rhs29_
        r0 = rhs30_
        d_177_v3_ = rhs31_
        lhs26_[lhs27_] = rhs32_
        r0 = (d_175_v1_)
        hi1_ = p1
        for d_182_i1_ in range((0) - (p1), hi1_):
            d_183_v7_: _dafny.Array
            nw36_ = _dafny.Array(int(0), 1)
            d_183_v7_ = nw36_
            index31_ = default__.safeIndex(318, (d_183_v7_).length(0))
            (d_183_v7_)[index31_] = d_182_i1_
            d_184_v8_: _dafny.Map
            d_184_v8_ = _dafny.Map({p0: p1})
            d_185_v9_: _dafny.Map
            d_185_v9_ = _dafny.Map({d_182_i1_: self.f26})
            d_186_v10_: _dafny.MultiSet
            d_186_v10_ = _dafny.MultiSet([d_182_i1_])
            index32_ = default__.safeIndex(318, (d_183_v7_).length(0))
            rhs33_ = d_183_v7_
            rhs34_ = not(False)
            rhs35_ = default__.safeDivisionInt(p1, default__.safeDivisionInt(p1, len((d_184_v8_).set(True, p1))))
            rhs36_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_187_i2_ in range(default__.abs(848))])
            rhs37_ = ((d_184_v8_)[(p1) < (default__.fm3(d_182_i1_, d_185_v9_, d_182_i1_, globalState))] if ((p1) < (default__.fm3(d_182_i1_, d_185_v9_, d_182_i1_, globalState))) in (d_184_v8_) else (d_186_v10_).cardinality)
            lhs28_ = globalState
            lhs29_ = d_183_v7_
            lhs30_ = default__.safeIndex(318, (d_183_v7_).length(0))
            lhs31_ = globalState
            lhs28_.f5 = rhs33_
            r0 = rhs34_
            lhs29_[lhs30_] = rhs35_
            d_180_v5_ = rhs36_
            lhs31_.f19 = rhs37_
            (self).f26 = self.f26
            d_188_v11_: _dafny.Map
            d_188_v11_ = _dafny.Map({(p0 if p0 else p0): self.f26})
            d_188_v11_ = (d_188_v11_).set((p0 if p0 else not(self.f26)), self.f26)
            d_189_v12_: _dafny.Map
            d_189_v12_ = _dafny.Map({d_182_i1_: d_183_v7_})
            rhs38_ = self.f26
            rhs39_ = (d_181_v6_)[default__.safeIndex((p1) * (len(d_185_v9_)), len(d_181_v6_))]
            rhs40_ = (d_182_i1_) * (len(d_189_v12_))
            rhs41_ = (0) - (len(d_174_v0_))
            rhs42_ = ((len(_dafny.Map({self.f26: self.f26}))) * ((d_183_v7_)[default__.safeIndex(318, (d_183_v7_).length(0))])) >= ((d_183_v7_)[default__.safeIndex(318, (d_183_v7_).length(0))])
            lhs32_ = globalState
            lhs33_ = globalState
            lhs34_ = self
            r0 = rhs38_
            r0 = rhs39_
            lhs32_.f15 = rhs40_
            lhs33_.f15 = rhs41_
            lhs34_.f26 = rhs42_
        d_190_v13_: _dafny.Map
        d_190_v13_ = _dafny.Map({self.f26: p1})
        d_191_v14_: _dafny.Set
        d_191_v14_ = _dafny.Set({p0, self.f26, p0})
        d_192_v15_: _dafny.Map
        d_192_v15_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_193_i3_ in range(default__.abs(456))])): self.f26})
        d_190_v13_ = ((d_190_v13_) | (d_190_v13_)) | (_dafny.Map({default__.fm0(d_174_v0_, 693, len(d_191_v14_), globalState): default__.fm3(p1, d_192_v15_, p1, globalState)}))
        (globalState).f19 = default__.fm3(len(default__.fm4((0) - (p1), self.f26, self.f26, (_dafny.MultiSet([p1, p1])).cardinality, globalState)), d_192_v15_, (0) - ((p1) + (p1)), globalState)
        d_194_v16_: _dafny.MultiSet
        d_194_v16_ = _dafny.MultiSet([self.f26, False, self.f26, self.f26, self.f26])
        d_195_v17_: _dafny.Map
        d_195_v17_ = _dafny.Map({p1: p1})
        d_196_v18_: D1
        d_196_v18_ = D1_DC2(len((d_179_v4_)[default__.safeIndex(306, (d_179_v4_).length(0))]), p1, -594, d_195_v17_)
        d_197_v19_: _dafny.Array
        nw37_ = _dafny.Array(None, 19)
        nw37_[int(0)] = self.f26
        nw37_[int(1)] = False
        nw37_[int(2)] = self.f26
        nw37_[int(3)] = p0
        nw37_[int(4)] = False
        nw37_[int(5)] = True
        nw37_[int(6)] = p0
        nw37_[int(7)] = True
        nw37_[int(8)] = self.f26
        nw37_[int(9)] = p0
        nw37_[int(10)] = self.f26
        nw37_[int(11)] = True
        nw37_[int(12)] = p0
        nw37_[int(13)] = (p0)
        nw37_[int(14)] = p0
        nw37_[int(15)] = self.f26
        nw37_[int(16)] = self.f26
        nw37_[int(17)] = p0
        nw37_[int(18)] = True
        d_197_v19_ = nw37_
        d_198_v20_: _dafny.Map
        d_198_v20_ = _dafny.Map({d_196_v18_: d_197_v19_})
        d_199_v21_: _dafny.Map
        d_199_v21_ = _dafny.Map({default__.fm3(p1, d_192_v15_, (d_194_v16_).cardinality, globalState): ((d_198_v20_)[d_196_v18_] if (d_196_v18_) in (d_198_v20_) else d_197_v19_)})
        d_199_v21_ = d_199_v21_
        d_200_v22_: _dafny.MultiSet
        d_200_v22_ = _dafny.MultiSet([p1])
        d_201_v23_: str
        d_201_v23_ = _dafny.CodePoint('c')
        d_202_v24_: _dafny.Map
        d_202_v24_ = _dafny.Map({d_200_v22_: d_201_v23_})
        r0 = ((d_202_v24_) | (d_202_v24_)) != (d_202_v24_)
        d_203_v25_: _dafny.Seq
        d_203_v25_ = _dafny.SeqWithoutIsStrInference([d_200_v22_, (d_200_v22_) - (default__.fm5(default__.fm0(d_174_v0_, p1, p1, globalState), self.f26, self.f26, globalState))])
        r1 = d_203_v25_
        return r0, r1

