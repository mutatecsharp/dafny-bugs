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
    def fm1(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('w')]))

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        return _dafny.CodePoint('s')

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({not(False)})).intersection(_dafny.Set({not(False)}))) | (_dafny.Set({True}))

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return ((548) + (len(_dafny.SeqWithoutIsStrInference([412])))) >= (-314)

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([421, -17])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, not(False), False])), len(_dafny.SeqWithoutIsStrInference([False]))])])))

    @staticmethod
    def fm6(p0, p1, globalState):
        return (100) + ((D1_DC3(339, not(not(not(not(True)))))).cf6)

    @staticmethod
    def fm7(p0, p1, globalState):
        source0_ = D3_DC8((_dafny.MultiSet([-463, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))), 851, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdurb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amh"))])), 331])).cardinality)
        if source0_.is_DC8:
            d_0___mcc_h0_ = source0_.cf13
            d_1_cf13_ = d_0___mcc_h0_
            return D1_DC3(d_1_cf13_, False)
        elif source0_.is_DC7:
            d_2___mcc_h1_ = source0_.cf12
            d_3_cf12_ = d_2___mcc_h1_
            return D1_DC3((_dafny.MultiSet([True, True])).cardinality, False)
        elif True:
            d_4___mcc_h2_ = source0_.cf14
            d_5_cf14_ = d_4___mcc_h2_
            return D1_DC3((_dafny.MultiSet([len(_dafny.Set({not(False)}))])).cardinality, False)

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(767, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nux")))), (len(_dafny.Set({-871, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqqxre")))}))) * (len(_dafny.Map({not(True): True})))])

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        source1_ = D2_DC6()
        if source1_.is_DC5:
            d_6___mcc_h0_ = source1_.cf9
            d_7___mcc_h1_ = source1_.cf10
            d_8___mcc_h2_ = source1_.cf11
            d_9_cf11_ = d_8___mcc_h2_
            d_10_cf10_ = d_7___mcc_h1_
            d_11_cf9_ = d_6___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([D0_DC0(_dafny.CodePoint('y')), D0_DC0(_dafny.CodePoint('j')), D0_DC0(_dafny.CodePoint('w')), D0_DC0(_dafny.CodePoint('h'))])
        elif source1_.is_DC6:
            return _dafny.SeqWithoutIsStrInference([D0_DC0(_dafny.CodePoint('b')), D0_DC0(_dafny.CodePoint('p'))])
        elif True:
            d_12___mcc_h3_ = source1_.cf8
            d_13_cf8_ = d_12___mcc_h3_
            return _dafny.SeqWithoutIsStrInference([D0_DC0(_dafny.CodePoint('y')), D0_DC0(_dafny.CodePoint('r'))])

    @staticmethod
    def m0(globalState):
        r0: bool = False
        r1: _dafny.Set = _dafny.Set({})
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_14_v0_: int
        d_14_v0_ = 506
        d_15_v1_: _dafny.Seq
        d_15_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skwjs"))
        d_16_v2_: _dafny.Map
        d_16_v2_ = _dafny.Map({default__.safeModuloInt(d_14_v0_, len(d_15_v1_)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))})
        d_17_v3_: _dafny.Seq
        d_17_v3_ = _dafny.SeqWithoutIsStrInference([d_15_v1_])
        d_16_v2_ = (d_16_v2_).set((d_14_v0_) * (61), ((d_17_v3_)[default__.safeIndex(d_14_v0_, len(d_17_v3_))]) + (d_15_v1_))
        d_18_v4_: _dafny.Array
        nw0_ = _dafny.Array(_dafny.Set({}), 17)
        d_18_v4_ = nw0_
        d_19_v5_: bool
        d_19_v5_ = True
        d_20_v6_: D1
        d_20_v6_ = D1_DC3((0) - (d_14_v0_), d_19_v5_)
        d_21_v7_: _dafny.Map
        d_21_v7_ = _dafny.Map({d_19_v5_: _dafny.Set({d_20_v6_})})
        pat_let_tv0_ = d_14_v0_
        index0_ = default__.safeIndex(874, (d_18_v4_).length(0))
        def iife0_(_pat_let0_0):
            def iife1_(d_22_dt__update__tmp_h0_):
                def iife2_(_pat_let1_0):
                    def iife3_(d_23_dt__update_hcf6_h0_):
                        return D1_DC3(d_23_dt__update_hcf6_h0_, (d_22_dt__update__tmp_h0_).cf7)
                    return iife3_(_pat_let1_0)
                return iife2_(pat_let_tv0_)
            return iife1_(_pat_let0_0)
        (d_18_v4_)[index0_] = ((d_21_v7_)[d_19_v5_] if (d_19_v5_) in (d_21_v7_) else _dafny.Set({d_20_v6_, iife0_(d_20_v6_), d_20_v6_}))
        d_24_v8_: _dafny.Set
        d_24_v8_ = _dafny.Set({d_20_v6_})
        index1_ = default__.safeIndex(874, (d_18_v4_).length(0))
        (d_18_v4_)[index1_] = d_24_v8_
        d_25_i0_: int
        d_25_i0_ = 0
        with _dafny.label("0"):
            while (d_19_v5_) == (d_19_v5_):
                with _dafny.c_label("0"):
                    if (d_25_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_25_i0_ = (d_25_i0_) + (1)
                    d_14_v0_ = (0) - (d_14_v0_)
                    d_26_v9_: _dafny.Array
                    def lambda0_(d_27_v0_):
                        def lambda1_(d_28_i1_):
                            return default__.safeDivisionInt(d_28_i1_, d_27_v0_)

                        return lambda1_

                    init0_ = lambda0_(d_14_v0_)
                    nw1_ = _dafny.Array(None, 29)
                    for i0_0_ in range(nw1_.length(0)):
                        nw1_[i0_0_] = init0_(i0_0_)
                    d_26_v9_ = nw1_
                    index2_ = default__.safeIndex(54, (d_26_v9_).length(0))
                    (d_26_v9_)[index2_] = d_14_v0_
                    index3_ = default__.safeIndex(54, (d_26_v9_).length(0))
                    (d_26_v9_)[index3_] = d_14_v0_
                    d_29_v10_: str
                    d_29_v10_ = _dafny.CodePoint('k')
                    d_30_v11_: D0
                    d_30_v11_ = D0_DC0(d_29_v10_)
                    d_31_v12_: _dafny.Seq
                    d_31_v12_ = _dafny.SeqWithoutIsStrInference([d_30_v11_])
                    d_32_v13_: _dafny.Seq
                    d_32_v13_ = _dafny.SeqWithoutIsStrInference([(d_26_v9_)[default__.safeIndex(54, (d_26_v9_).length(0))]])
                    d_33_v14_: _dafny.Seq
                    d_33_v14_ = _dafny.SeqWithoutIsStrInference([True, d_19_v5_, d_19_v5_, not(d_19_v5_)])
                    d_34_v15_: _dafny.Seq
                    d_34_v15_ = _dafny.SeqWithoutIsStrInference([(d_26_v9_)[default__.safeIndex(54, (d_26_v9_).length(0))], (d_32_v13_)[default__.safeIndex(len(_dafny.Map({650: d_19_v5_})), len(d_32_v13_))], len(d_33_v14_), d_14_v0_, len(d_15_v1_)])
                    rhs0_ = ((d_31_v12_) + (default__.fm9(d_14_v0_, (d_26_v9_)[default__.safeIndex(54, (d_26_v9_).length(0))], d_19_v5_, d_19_v5_, globalState))) + (d_31_v12_)
                    rhs1_ = default__.fm4(d_15_v1_, (_dafny.SeqWithoutIsStrInference([len(d_15_v1_) for d_35_i2_ in range(default__.abs(811))])) + (d_32_v13_), d_19_v5_, d_19_v5_, globalState)
                    d_31_v12_ = rhs0_
                    d_19_v5_ = rhs1_
                    r0 = d_19_v5_
                    pass
            pass
        d_36_v16_: _dafny.Seq
        d_36_v16_ = _dafny.SeqWithoutIsStrInference([d_14_v0_])
        (globalState).f6 = default__.safeDivisionInt(default__.safeDivisionInt(706, d_14_v0_), (d_14_v0_) - ((0) - (len(_dafny.Map({not(default__.fm4(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_37_i3_ in range(default__.abs(476))]), d_36_v16_, d_19_v5_, d_19_v5_, globalState)): d_14_v0_})))))
        d_38_v17_: C0
        nw2_ = C0()
        nw2_.ctor__()
        d_38_v17_ = nw2_
        d_39_v18_: _dafny.Map
        d_39_v18_ = _dafny.Map({d_14_v0_: d_38_v17_})
        d_40_v19_: _dafny.MultiSet
        d_40_v19_ = _dafny.MultiSet([d_39_v18_, d_39_v18_])
        d_41_v20_: _dafny.Set
        d_41_v20_ = _dafny.Set({d_19_v5_})
        d_42_v21_: _dafny.Map
        d_42_v21_ = _dafny.Map({((d_40_v19_)[d_39_v18_] if (d_39_v18_) in (d_40_v19_) else default__.fm6(d_14_v0_, d_19_v5_, globalState)): (d_41_v20_).ispropersubset(_dafny.Set({not(d_19_v5_)}))})
        if ((d_42_v21_)[(d_14_v0_) * (d_14_v0_)] if ((d_14_v0_) * (d_14_v0_)) in (d_42_v21_) else True):
            d_43_v22_: C0
            nw3_ = C0()
            nw3_.ctor__()
            d_43_v22_ = nw3_
            d_44_v23_: D0
            d_44_v23_ = D0_DC1((d_14_v0_) == (d_14_v0_), d_14_v0_, d_19_v5_, (d_14_v0_) * (d_14_v0_))
            d_44_v23_ = d_44_v23_
            d_45_v24_: _dafny.Seq
            d_45_v24_ = _dafny.SeqWithoutIsStrInference([False, d_19_v5_])
            d_46_v25_: _dafny.Seq
            d_46_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_19_v5_, d_19_v5_]), _dafny.SeqWithoutIsStrInference([False])])
            d_47_v26_: _dafny.Seq
            d_47_v26_ = _dafny.SeqWithoutIsStrInference([d_38_v17_])
            rhs2_ = (d_45_v24_) not in (d_46_v25_)
            rhs3_ = d_19_v5_
            rhs4_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_48_i4_ in range(default__.abs(883))])), len(d_47_v26_))
            lhs0_ = globalState
            d_19_v5_ = rhs2_
            r0 = rhs3_
            lhs0_.f6 = rhs4_
            d_49_v27_: _dafny.MultiSet
            d_49_v27_ = _dafny.MultiSet([d_14_v0_, d_14_v0_, d_14_v0_])
            d_50_v28_: _dafny.MultiSet
            d_50_v28_ = _dafny.MultiSet([d_19_v5_])
            d_51_v29_: str
            d_51_v29_ = _dafny.CodePoint('g')
            d_52_v30_: _dafny.Map
            d_52_v30_ = _dafny.Map({d_51_v29_: not(True)})
            d_53_v31_: _dafny.Array
            nw4_ = _dafny.Array(None, 23)
            nw4_[int(0)] = d_19_v5_
            nw4_[int(1)] = (d_20_v6_).cf7
            nw4_[int(2)] = (d_15_v1_) == (d_15_v1_)
            nw4_[int(3)] = (True if not(d_19_v5_) else d_19_v5_)
            nw4_[int(4)] = d_19_v5_
            nw4_[int(5)] = d_19_v5_
            nw4_[int(6)] = d_19_v5_
            nw4_[int(7)] = d_19_v5_
            nw4_[int(8)] = (d_14_v0_) != (d_14_v0_)
            nw4_[int(9)] = (len(d_36_v16_)) < (d_14_v0_)
            nw4_[int(10)] = d_19_v5_
            nw4_[int(11)] = d_19_v5_
            nw4_[int(12)] = d_19_v5_
            nw4_[int(13)] = not((d_19_v5_) or (False))
            nw4_[int(14)] = ((0) - (d_14_v0_)) > (len(d_41_v20_))
            nw4_[int(15)] = (d_49_v27_).isdisjoint(d_49_v27_)
            nw4_[int(16)] = d_19_v5_
            nw4_[int(17)] = (d_50_v28_).ispropersubset(d_50_v28_)
            nw4_[int(18)] = d_19_v5_
            nw4_[int(19)] = True
            nw4_[int(20)] = ((d_52_v30_)[d_51_v29_] if (d_51_v29_) in (d_52_v30_) else d_19_v5_)
            nw4_[int(21)] = not(d_19_v5_)
            nw4_[int(22)] = d_19_v5_
            d_53_v31_ = nw4_
            index4_ = default__.safeIndex(602, (d_53_v31_).length(0))
            (d_53_v31_)[index4_] = d_19_v5_
            index5_ = default__.safeIndex(602, (d_53_v31_).length(0))
            (d_53_v31_)[index5_] = not (d_19_v5_) or (d_19_v5_)
            d_54_v32_: _dafny.Map
            d_54_v32_ = _dafny.Map({d_19_v5_: d_51_v29_})
            index6_ = default__.safeIndex(602, (d_53_v31_).length(0))
            index7_ = default__.safeIndex(602, (d_53_v31_).length(0))
            rhs5_ = (d_53_v31_)[default__.safeIndex(602, (d_53_v31_).length(0))]
            rhs6_ = (((0) - (d_14_v0_)) - (d_14_v0_)) < (d_14_v0_)
            rhs7_ = (len(d_15_v1_)) >= (d_14_v0_)
            rhs8_ = (d_15_v1_).set(default__.safeIndex(d_14_v0_, len(d_15_v1_)), d_51_v29_)
            rhs9_ = d_54_v32_
            lhs1_ = d_53_v31_
            lhs2_ = default__.safeIndex(602, (d_53_v31_).length(0))
            lhs3_ = d_53_v31_
            lhs4_ = default__.safeIndex(602, (d_53_v31_).length(0))
            r0 = rhs5_
            lhs1_[lhs2_] = rhs6_
            lhs3_[lhs4_] = rhs7_
            r2 = rhs8_
            d_54_v32_ = rhs9_
        elif True:
            (globalState).f7 = (d_20_v6_).cf6
            (globalState).f6 = (0) - ((d_36_v16_)[default__.safeIndex(d_14_v0_, len(d_36_v16_))])
            d_55_v33_: _dafny.Array
            def lambda2_(d_56_v0_):
                def lambda3_(d_57_i5_):
                    return (d_57_i5_) - (d_56_v0_)

                return lambda3_

            init1_ = lambda2_(d_14_v0_)
            nw5_ = _dafny.Array(None, 21)
            for i0_1_ in range(nw5_.length(0)):
                nw5_[i0_1_] = init1_(i0_1_)
            d_55_v33_ = nw5_
            d_58_v34_: _dafny.MultiSet
            d_58_v34_ = _dafny.MultiSet([d_55_v33_, d_55_v33_])
            if (_dafny.MultiSet([d_55_v33_])) == (d_58_v34_):
                d_59_v35_: _dafny.Seq
                d_59_v35_ = _dafny.SeqWithoutIsStrInference([d_19_v5_, d_19_v5_])
                d_59_v35_ = (d_59_v35_) + (d_59_v35_)
                d_19_v5_ = ((d_14_v0_) - (d_14_v0_)) < (d_14_v0_)
                (globalState).f8 = ((len(d_59_v35_)) * (d_14_v0_)) >= (245)
                index8_ = default__.safeIndex(699, (d_55_v33_).length(0))
                (d_55_v33_)[index8_] = (d_14_v0_) - (d_14_v0_)
                d_60_v36_: _dafny.Map
                d_60_v36_ = _dafny.Map({d_55_v33_: default__.safeDivisionInt(d_14_v0_, d_14_v0_)})
                d_61_v37_: _dafny.Seq
                d_61_v37_ = _dafny.SeqWithoutIsStrInference([d_55_v33_])
                d_62_v38_: _dafny.Map
                d_62_v38_ = _dafny.Map({len(d_61_v37_): d_14_v0_})
                d_63_v39_: _dafny.Set
                d_63_v39_ = _dafny.Set({d_62_v38_, d_62_v38_})
                d_64_v40_: _dafny.Map
                d_64_v40_ = _dafny.Map({d_19_v5_: d_19_v5_})
                d_65_v41_: _dafny.Map
                d_65_v41_ = _dafny.Map({d_19_v5_: d_38_v17_})
                index9_ = default__.safeIndex(699, (d_55_v33_).length(0))
                rhs10_ = d_19_v5_
                rhs11_ = ((d_60_v36_)[d_55_v33_] if (d_55_v33_) in (d_60_v36_) else d_14_v0_)
                rhs12_ = d_14_v0_
                rhs13_ = d_19_v5_
                rhs14_ = (D0_DC1(d_19_v5_, len(d_63_v39_), ((d_64_v40_)[default__.fm4(d_15_v1_, _dafny.SeqWithoutIsStrInference([d_14_v0_]), d_19_v5_, d_19_v5_, globalState)] if (default__.fm4(d_15_v1_, _dafny.SeqWithoutIsStrInference([d_14_v0_]), d_19_v5_, d_19_v5_, globalState)) in (d_64_v40_) else d_19_v5_), len(d_65_v41_))).cf1
                lhs5_ = d_55_v33_
                lhs6_ = default__.safeIndex(699, (d_55_v33_).length(0))
                lhs7_ = globalState
                lhs8_ = globalState
                r0 = rhs10_
                lhs5_[lhs6_] = rhs11_
                lhs7_.f6 = rhs12_
                r0 = rhs13_
                lhs8_.f8 = rhs14_
                d_59_v35_ = (d_59_v35_) + (d_59_v35_)
            elif True:
                (globalState).f6 = d_14_v0_
                d_66_v42_: _dafny.Seq
                d_66_v42_ = _dafny.SeqWithoutIsStrInference([d_38_v17_, d_38_v17_])
                d_38_v17_ = (d_66_v42_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_67_i6_ in range(default__.abs(473))])), len(d_66_v42_))]
                d_68_v43_: _dafny.Seq
                d_68_v43_ = _dafny.SeqWithoutIsStrInference([d_19_v5_, d_19_v5_])
                d_69_v44_: _dafny.MultiSet
                d_69_v44_ = _dafny.MultiSet([d_14_v0_])
                d_70_v45_: _dafny.Map
                d_70_v45_ = _dafny.Map({d_68_v43_: (d_69_v44_).isdisjoint(d_69_v44_)})
                rhs15_ = default__.safeDivisionInt(d_14_v0_, default__.safeDivisionInt(-380, d_14_v0_))
                rhs16_ = (len((d_15_v1_) + (d_15_v1_))) * (d_14_v0_)
                rhs17_ = d_70_v45_
                lhs9_ = globalState
                lhs10_ = globalState
                lhs9_.f3 = rhs15_
                lhs10_.f3 = rhs16_
                d_70_v45_ = rhs17_
                d_71_v46_: C0
                nw6_ = C0()
                nw6_.ctor__()
                d_71_v46_ = nw6_
                (globalState).f3 = d_14_v0_
            d_72_v47_: _dafny.Array
            nw7_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_72_v47_ = nw7_
            d_73_v48_: _dafny.Array
            def lambda4_(d_74_v5_):
                def lambda5_(d_75_i7_):
                    return d_74_v5_

                return lambda5_

            init2_ = lambda4_(d_19_v5_)
            nw8_ = _dafny.Array(None, 17)
            for i0_2_ in range(nw8_.length(0)):
                nw8_[i0_2_] = init2_(i0_2_)
            d_73_v48_ = nw8_
            index10_ = default__.safeIndex(702, (d_72_v47_).length(0))
            (d_72_v47_)[index10_] = d_73_v48_
            index11_ = default__.safeIndex(702, (d_72_v47_).length(0))
            (d_72_v47_)[index11_] = d_73_v48_
            d_76_v49_: D0
            d_76_v49_ = D0_DC1(d_19_v5_, 34, not(d_19_v5_), (0) - (len((d_41_v20_) | (d_41_v20_))))
            source2_ = d_76_v49_
            if source2_.is_DC1:
                d_77___mcc_h0_ = source2_.cf1
                d_78___mcc_h1_ = source2_.cf2
                d_79___mcc_h2_ = source2_.cf3
                d_80___mcc_h3_ = source2_.cf4
                d_81_cf4_ = d_80___mcc_h3_
                d_82_cf3_ = d_79___mcc_h2_
                d_83_cf2_ = d_78___mcc_h1_
                d_84_cf1_ = d_77___mcc_h0_
                r2 = d_15_v1_
                d_73_v48_ = (d_72_v47_)[default__.safeIndex(702, (d_72_v47_).length(0))]
                (globalState).f3 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vspevqskp")))
                d_85_v50_: _dafny.Map
                d_85_v50_ = _dafny.Map({d_82_cf3_: d_81_cf4_})
                d_16_v2_ = (d_16_v2_).set((len(d_85_v50_)) * (d_14_v0_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_86_i8_ in range(default__.abs(655))]))
            elif True:
                d_87___mcc_h4_ = source2_.cf0
                d_88_cf0_ = d_87___mcc_h4_
                (globalState).f3 = (d_36_v16_)[default__.safeIndex(len(d_41_v20_), len(d_36_v16_))]
                d_89_v51_: _dafny.MultiSet
                d_89_v51_ = _dafny.MultiSet([d_19_v5_])
                d_90_v52_: _dafny.Set
                d_90_v52_ = _dafny.Set({d_89_v51_, d_89_v51_, d_89_v51_})
                d_91_v53_: _dafny.Map
                d_91_v53_ = _dafny.Map({True: d_90_v52_})
                d_90_v52_ = ((d_91_v53_)[default__.fm4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")), d_36_v16_, not(not(d_19_v5_)), not(d_19_v5_), globalState)] if (default__.fm4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")), d_36_v16_, not(not(d_19_v5_)), not(d_19_v5_), globalState)) in (d_91_v53_) else d_90_v52_)
                (globalState).f8 = d_19_v5_
                d_92_v54_: _dafny.Array
                nw9_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_92_v54_ = nw9_
                d_93_v55_: _dafny.Array
                def lambda6_(d_94_i9_):
                    return _dafny.SeqWithoutIsStrInference([False])

                init3_ = lambda6_
                nw10_ = _dafny.Array(None, 19)
                for i0_3_ in range(nw10_.length(0)):
                    nw10_[i0_3_] = init3_(i0_3_)
                d_93_v55_ = nw10_
                index12_ = default__.safeIndex(478, (d_92_v54_).length(0))
                (d_92_v54_)[index12_] = d_93_v55_
                d_95_v56_: D1
                d_95_v56_ = D1_DC2(_dafny.Map({default__.fm4(d_15_v1_, d_36_v16_, d_19_v5_, d_19_v5_, globalState): d_14_v0_}))
                d_96_v57_: _dafny.Seq
                d_96_v57_ = _dafny.SeqWithoutIsStrInference([d_93_v55_, d_93_v55_])
                d_97_v58_: _dafny.Seq
                d_97_v58_ = _dafny.SeqWithoutIsStrInference([d_93_v55_, (d_96_v57_)[default__.safeIndex((0) - (d_14_v0_), len(d_96_v57_))], d_93_v55_])
                index13_ = default__.safeIndex(478, (d_92_v54_).length(0))
                rhs18_ = (d_96_v57_)[default__.safeIndex(d_14_v0_, len(d_96_v57_))]
                rhs19_ = (default__.fm6(d_14_v0_, d_19_v5_, globalState)) * (default__.fm6(d_14_v0_, d_19_v5_, globalState))
                rhs20_ = d_95_v56_
                lhs11_ = d_92_v54_
                lhs12_ = default__.safeIndex(478, (d_92_v54_).length(0))
                lhs13_ = globalState
                lhs11_[lhs12_] = rhs18_
                lhs13_.f6 = rhs19_
                d_95_v56_ = rhs20_
        d_98_v59_: _dafny.Seq
        d_98_v59_ = _dafny.SeqWithoutIsStrInference([d_41_v20_, d_41_v20_])
        d_99_i10_: int
        d_99_i10_ = 0
        with _dafny.label("1"):
            while (len(d_36_v16_)) <= (len((d_98_v59_)[default__.safeIndex(d_14_v0_, len(d_98_v59_))])):
                with _dafny.c_label("1"):
                    if (d_99_i10_) >= (100):
                        raise _dafny.Break("1")
                    d_99_i10_ = (d_99_i10_) + (1)
                    d_100_v60_: _dafny.Array
                    def lambda7_(d_101_v0_):
                        def lambda8_(d_102_i11_):
                            return (d_102_i11_) * (d_101_v0_)

                        return lambda8_

                    init4_ = lambda7_(d_14_v0_)
                    nw11_ = _dafny.Array(None, 7)
                    for i0_4_ in range(nw11_.length(0)):
                        nw11_[i0_4_] = init4_(i0_4_)
                    d_100_v60_ = nw11_
                    d_103_v61_: _dafny.MultiSet
                    d_103_v61_ = _dafny.MultiSet([d_100_v60_, d_100_v60_, d_100_v60_, d_100_v60_])
                    rhs21_ = (D2_DC4(d_100_v60_)).cf8
                    rhs22_ = d_14_v0_
                    rhs23_ = d_19_v5_
                    rhs24_ = d_19_v5_
                    rhs25_ = ((d_103_v61_).set(d_100_v60_, default__.abs((d_14_v0_) - (d_14_v0_)))).cardinality
                    lhs14_ = globalState
                    lhs15_ = globalState
                    lhs16_ = globalState
                    d_100_v60_ = rhs21_
                    d_14_v0_ = rhs22_
                    lhs14_.f8 = rhs23_
                    lhs15_.f8 = rhs24_
                    lhs16_.f6 = rhs25_
                    if d_19_v5_:
                        d_104_v62_: _dafny.Map
                        d_104_v62_ = _dafny.Map({292: d_14_v0_})
                        d_105_v63_: str
                        d_105_v63_ = _dafny.CodePoint('p')
                        d_106_v64_: _dafny.MultiSet
                        d_106_v64_ = _dafny.MultiSet([d_105_v63_])
                        d_107_v65_: _dafny.Array
                        nw12_ = _dafny.Array(None, 21)
                        nw12_[int(0)] = d_19_v5_
                        nw12_[int(1)] = d_19_v5_
                        nw12_[int(2)] = default__.fm4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owms")), d_36_v16_, d_19_v5_, d_19_v5_, globalState)
                        nw12_[int(3)] = d_19_v5_
                        nw12_[int(4)] = d_19_v5_
                        nw12_[int(5)] = d_19_v5_
                        nw12_[int(6)] = d_19_v5_
                        nw12_[int(7)] = (len(d_104_v62_)) > (d_14_v0_)
                        nw12_[int(8)] = False
                        nw12_[int(9)] = d_19_v5_
                        nw12_[int(10)] = (d_106_v64_).ispropersubset(d_106_v64_)
                        nw12_[int(11)] = d_19_v5_
                        nw12_[int(12)] = d_19_v5_
                        nw12_[int(13)] = d_19_v5_
                        nw12_[int(14)] = d_19_v5_
                        nw12_[int(15)] = d_19_v5_
                        nw12_[int(16)] = d_19_v5_
                        nw12_[int(17)] = d_19_v5_
                        nw12_[int(18)] = d_19_v5_
                        nw12_[int(19)] = d_19_v5_
                        nw12_[int(20)] = d_19_v5_
                        d_107_v65_ = nw12_
                        index14_ = default__.safeIndex(836, (d_107_v65_).length(0))
                        (d_107_v65_)[index14_] = (d_19_v5_ if d_19_v5_ else True)
                        index15_ = default__.safeIndex(836, (d_107_v65_).length(0))
                        (d_107_v65_)[index15_] = d_19_v5_
                        d_38_v17_ = d_38_v17_
                        d_108_v66_: _dafny.Array
                        nw13_ = _dafny.Array(None, 23)
                        d_108_v66_ = nw13_
                        index16_ = default__.safeIndex(624, (d_108_v66_).length(0))
                        (d_108_v66_)[index16_] = d_38_v17_
                        d_109_v67_: _dafny.Seq
                        d_109_v67_ = _dafny.SeqWithoutIsStrInference([d_19_v5_, (d_107_v65_)[default__.safeIndex(836, (d_107_v65_).length(0))]])
                        d_110_v68_: _dafny.MultiSet
                        d_110_v68_ = _dafny.MultiSet([d_19_v5_, False])
                        d_111_v69_: _dafny.Set
                        d_111_v69_ = _dafny.Set({len(d_109_v67_), (d_110_v68_).cardinality, d_14_v0_})
                        index17_ = default__.safeIndex(624, (d_108_v66_).length(0))
                        rhs26_ = d_38_v17_
                        rhs27_ = d_38_v17_
                        rhs28_ = ((len(d_111_v69_)) * (d_14_v0_)) + (d_14_v0_)
                        rhs29_ = (708) * (len(d_42_v21_))
                        lhs17_ = d_108_v66_
                        lhs18_ = default__.safeIndex(624, (d_108_v66_).length(0))
                        lhs19_ = globalState
                        lhs20_ = globalState
                        d_38_v17_ = rhs26_
                        lhs17_[lhs18_] = rhs27_
                        lhs19_.f3 = rhs28_
                        lhs20_.f3 = rhs29_
                        (globalState).f6 = d_14_v0_
                        d_112_v70_: C0
                        nw14_ = C0()
                        nw14_.ctor__()
                        d_112_v70_ = nw14_
                    elif True:
                        (globalState).f8 = d_19_v5_
                        (globalState).f3 = d_14_v0_
                        d_113_v71_: _dafny.Set
                        d_113_v71_ = _dafny.Set({d_14_v0_, (0) - (d_14_v0_)})
                        d_19_v5_ = (len(d_113_v71_)) in ((d_36_v16_) + (d_36_v16_))
                        d_114_v72_: D0
                        d_114_v72_ = D0_DC1(False, (0) - (d_14_v0_), not(d_19_v5_), d_14_v0_)
                        d_115_v73_: D2
                        d_115_v73_ = D2_DC5(d_14_v0_, d_14_v0_, d_114_v72_)
                        index18_ = default__.safeIndex(531, (d_100_v60_).length(0))
                        (d_100_v60_)[index18_] = default__.safeModuloInt(d_14_v0_, (d_115_v73_).cf9)
                        index19_ = default__.safeIndex(531, (d_100_v60_).length(0))
                        (d_100_v60_)[index19_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "poflqgduh")))
                        d_116_v74_: _dafny.Array
                        nw15_ = _dafny.Array(int(0), 19)
                        d_116_v74_ = nw15_
                        d_117_v75_: _dafny.Map
                        d_117_v75_ = _dafny.Map({False: d_116_v74_})
                        d_118_v76_: _dafny.Seq
                        d_118_v76_ = _dafny.SeqWithoutIsStrInference([d_117_v75_, d_117_v75_])
                        d_119_v77_: D3
                        d_119_v77_ = D3_DC7(d_117_v75_)
                        d_120_v78_: _dafny.Seq
                        d_120_v78_ = _dafny.SeqWithoutIsStrInference([d_19_v5_])
                        d_121_v79_: _dafny.Array
                        nw16_ = _dafny.Array(None, 20)
                        nw16_[int(0)] = d_117_v75_
                        nw16_[int(1)] = d_117_v75_
                        nw16_[int(2)] = d_117_v75_
                        nw16_[int(3)] = d_117_v75_
                        nw16_[int(4)] = d_117_v75_
                        nw16_[int(5)] = d_117_v75_
                        nw16_[int(6)] = _dafny.Map({default__.fm4(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_122_i12_ in range(default__.abs(-688))]), _dafny.SeqWithoutIsStrInference([len(d_42_v21_), d_14_v0_]), d_19_v5_, d_19_v5_, globalState): d_116_v74_})
                        nw16_[int(7)] = d_117_v75_
                        nw16_[int(8)] = _dafny.Map({d_19_v5_: d_116_v74_})
                        nw16_[int(9)] = (d_117_v75_) | ((d_118_v76_)[default__.safeIndex(d_14_v0_, len(d_118_v76_))])
                        nw16_[int(10)] = d_117_v75_
                        nw16_[int(11)] = d_117_v75_
                        nw16_[int(12)] = (d_119_v77_).cf12
                        nw16_[int(13)] = d_117_v75_
                        nw16_[int(14)] = d_117_v75_
                        nw16_[int(15)] = d_117_v75_
                        nw16_[int(16)] = (d_117_v75_) | (_dafny.Map({True: d_116_v74_}))
                        nw16_[int(17)] = (d_117_v75_) | (d_117_v75_)
                        nw16_[int(18)] = _dafny.Map({(d_120_v78_)[default__.safeIndex((d_100_v60_)[default__.safeIndex(531, (d_100_v60_).length(0))], len(d_120_v78_))]: d_116_v74_})
                        nw16_[int(19)] = d_117_v75_
                        d_121_v79_ = nw16_
                        index20_ = default__.safeIndex(438, (d_121_v79_).length(0))
                        (d_121_v79_)[index20_] = (d_117_v75_) | (_dafny.Map({d_19_v5_: d_116_v74_}))
                        index21_ = default__.safeIndex(438, (d_121_v79_).length(0))
                        (d_121_v79_)[index21_] = d_117_v75_
                    d_123_v80_: D0
                    d_123_v80_ = D0_DC1(True, d_14_v0_, d_19_v5_, 848)
                    (globalState).f7 = default__.safeModuloInt((d_123_v80_).cf2, 34)
                    d_124_v81_: C0
                    nw17_ = C0()
                    nw17_.ctor__()
                    d_124_v81_ = nw17_
                    pass
            pass
        r0 = d_19_v5_
        def iife4_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_14_v0_: d_14_v0_}))])) + (d_36_v16_)).Elements:
                d_125_v82_: int = compr_0_
                if (d_125_v82_) in ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_14_v0_: d_14_v0_}))])) + (d_36_v16_)):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeDivisionInt(d_125_v82_, 669)]))
            return _dafny.Set(coll0_)
        r1 = iife4_()
        
        d_126_v83_: _dafny.Map
        d_126_v83_ = _dafny.Map({d_19_v5_: d_14_v0_})
        r2 = ((default__.fm1(len(d_126_v83_), d_19_v5_, globalState)) + (d_15_v1_)) + (d_15_v1_)
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_127_v0_: int
        d_127_v0_ = -183
        d_128_v1_: _dafny.MultiSet
        d_128_v1_ = _dafny.MultiSet([d_127_v0_])
        d_129_v2_: _dafny.Array
        nw18_ = _dafny.Array(int(0), 19)
        d_129_v2_ = nw18_
        d_130_v3_: _dafny.Seq
        d_130_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqq"))
        d_131_v4_: _dafny.Set
        d_131_v4_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwijw")), d_130_v3_})
        d_132_globalState_: GlobalState
        nw19_ = GlobalState()
        nw19_.ctor__(d_128_v1_, d_129_v2_, False, 434, 588, 870, 77, -390, False, True, d_131_v4_, True, 980, 718)
        d_132_globalState_ = nw19_
        d_133_v5_: bool
        d_133_v5_ = False
        d_134_v6_: _dafny.Map
        d_134_v6_ = _dafny.Map({_dafny.Map({d_127_v0_: not(d_133_v5_)}): d_127_v0_})
        d_135_v7_: _dafny.Map
        d_135_v7_ = _dafny.Map({-421: False})
        d_134_v6_ = (d_134_v6_).set(d_135_v7_, (0) - (d_127_v0_))
        d_136_v8_: C0
        nw20_ = C0()
        nw20_.ctor__()
        d_136_v8_ = nw20_
        d_137_v9_: _dafny.Array
        nw21_ = _dafny.Array(False, 28)
        d_137_v9_ = nw21_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_137_v9_).length(0)):
            d_138_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_138_i0_)) and ((d_138_i0_) < ((d_137_v9_).length(0)))):
                (d_137_v9_)[(d_138_i0_)] = d_133_v5_
        d_139_v10_: bool
        d_140_v11_: _dafny.Set
        d_141_v12_: _dafny.Seq
        out0_: bool
        out1_: _dafny.Set
        out2_: _dafny.Seq
        out0_, out1_, out2_ = default__.m0(d_132_globalState_)
        d_139_v10_ = out0_
        d_140_v11_ = out1_
        d_141_v12_ = out2_
        d_142_i1_: int
        d_142_i1_ = 0
        with _dafny.label("2"):
            while d_139_v10_:
                with _dafny.c_label("2"):
                    if (d_142_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_142_i1_ = (d_142_i1_) + (1)
                    d_139_v10_ = not (not(d_133_v5_)) or (not((d_127_v0_) < (d_127_v0_)))
                    d_143_v13_: bool
                    d_144_v14_: _dafny.Set
                    d_145_v15_: _dafny.Seq
                    out3_: bool
                    out4_: _dafny.Set
                    out5_: _dafny.Seq
                    out3_, out4_, out5_ = default__.m0(d_132_globalState_)
                    d_143_v13_ = out3_
                    d_144_v14_ = out4_
                    d_145_v15_ = out5_
                    (d_132_globalState_).f3 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tiuwpcrys")))
                    d_146_v16_: _dafny.Map
                    d_146_v16_ = _dafny.Map({d_127_v0_: d_141_v12_})
                    d_141_v12_ = (d_145_v15_) + (((d_146_v16_)[d_127_v0_] if (d_127_v0_) in (d_146_v16_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_147_i2_ in range(default__.abs(747))])))
                    pass
            pass
        d_148_i3_: int
        d_148_i3_ = 0
        with _dafny.label("3"):
            while (d_139_v10_ if (d_135_v7_) == (_dafny.Map({d_127_v0_: True})) else d_139_v10_):
                with _dafny.c_label("3"):
                    if (d_148_i3_) >= (100):
                        raise _dafny.Break("3")
                    d_148_i3_ = (d_148_i3_) + (1)
                    (d_132_globalState_).f3 = d_127_v0_
                    d_149_v17_: C0
                    nw22_ = C0()
                    nw22_.ctor__()
                    d_149_v17_ = nw22_
                    d_130_v3_ = (d_130_v3_) + ((_dafny.SeqWithoutIsStrInference([(D0_DC0(_dafny.CodePoint('x'))).cf0 for d_150_i4_ in range(default__.abs(-155))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_151_i5_ in range(default__.abs(-596))])))
                    d_152_v18_: _dafny.Map
                    d_152_v18_ = _dafny.Map({d_133_v5_: (d_133_v5_) == (d_139_v10_)})
                    (d_132_globalState_).f6 = len(d_152_v18_)
                    pass
            pass
        d_153_i6_: int
        d_153_i6_ = 0
        with _dafny.label("4"):
            while d_133_v5_:
                with _dafny.c_label("4"):
                    if (d_153_i6_) >= (100):
                        raise _dafny.Break("4")
                    d_153_i6_ = (d_153_i6_) + (1)
                    d_154_v19_: str
                    d_154_v19_ = _dafny.CodePoint('v')
                    d_155_v20_: D0
                    d_155_v20_ = D0_DC0(d_154_v19_)
                    rhs30_ = d_155_v20_
                    rhs31_ = False
                    lhs21_ = d_132_globalState_
                    d_155_v20_ = rhs30_
                    lhs21_.f8 = rhs31_
                    d_130_v3_ = d_130_v3_
                    d_156_v21_: _dafny.Array
                    nw23_ = _dafny.Array(_dafny.Array(None, 0), 5)
                    d_156_v21_ = nw23_
                    index22_ = default__.safeIndex(128, (d_156_v21_).length(0))
                    (d_156_v21_)[index22_] = d_129_v2_
                    index23_ = default__.safeIndex(128, (d_156_v21_).length(0))
                    (d_156_v21_)[index23_] = d_129_v2_
                    d_133_v5_ = d_133_v5_
                    pass
            pass
        d_157_v22_: str
        d_157_v22_ = _dafny.CodePoint('q')
        d_158_v23_: D0
        d_158_v23_ = D0_DC0(d_157_v22_)
        d_158_v23_ = d_158_v23_
        d_139_v10_ = d_133_v5_
        d_159_v24_: D0
        d_159_v24_ = D0_DC1(d_139_v10_, d_127_v0_, True, d_127_v0_)
        d_160_i7_: int
        d_160_i7_ = 0
        with _dafny.label("5"):
            pat_let_tv1_ = d_139_v10_
            pat_let_tv2_ = d_133_v5_
            def lambda9_(source3_):
                if source3_.is_DC1:
                    d_168___mcc_h0_ = source3_.cf1
                    d_169___mcc_h1_ = source3_.cf2
                    d_170___mcc_h2_ = source3_.cf3
                    d_171___mcc_h3_ = source3_.cf4
                    d_172_cf4_ = d_171___mcc_h3_
                    d_173_cf3_ = d_170___mcc_h2_
                    d_174_cf2_ = d_169___mcc_h1_
                    d_175_cf1_ = d_168___mcc_h0_
                    return pat_let_tv1_
                elif True:
                    d_176___mcc_h4_ = source3_.cf0
                    d_177_cf0_ = d_176___mcc_h4_
                    return pat_let_tv2_

            while lambda9_(d_159_v24_):
                with _dafny.c_label("5"):
                    if (d_160_i7_) >= (100):
                        raise _dafny.Break("5")
                    d_160_i7_ = (d_160_i7_) + (1)
                    (d_132_globalState_).f6 = d_127_v0_
                    d_161_v25_: C0
                    nw24_ = C0()
                    nw24_.ctor__()
                    d_161_v25_ = nw24_
                    d_162_v26_: _dafny.Map
                    d_162_v26_ = _dafny.Map({d_127_v0_: d_161_v25_})
                    (d_132_globalState_).f8 = (d_127_v0_) in ((d_162_v26_) | (d_162_v26_))
                    if d_133_v5_:
                        (d_132_globalState_).f8 = d_139_v10_
                        d_163_v27_: _dafny.Array
                        nw25_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
                        d_163_v27_ = nw25_
                        index24_ = default__.safeIndex(578, (d_163_v27_).length(0))
                        (d_163_v27_)[index24_] = d_130_v3_
                        index25_ = default__.safeIndex(578, (d_163_v27_).length(0))
                        rhs32_ = (default__.fm1(d_127_v0_, d_139_v10_, d_132_globalState_)).set(default__.safeIndex(d_127_v0_, len(default__.fm1(d_127_v0_, d_139_v10_, d_132_globalState_))), default__.fm2(d_141_v12_, d_133_v5_, d_139_v10_, d_132_globalState_))
                        rhs33_ = d_139_v10_
                        rhs34_ = ((d_127_v0_) - (d_127_v0_)) - (len(d_130_v3_))
                        lhs22_ = d_163_v27_
                        lhs23_ = default__.safeIndex(578, (d_163_v27_).length(0))
                        lhs24_ = d_132_globalState_
                        lhs25_ = d_132_globalState_
                        lhs22_[lhs23_] = rhs32_
                        lhs24_.f8 = rhs33_
                        lhs25_.f3 = rhs34_
                        d_133_v5_ = d_139_v10_
                        d_164_v28_: _dafny.Array
                        nw26_ = _dafny.Array(None, 15)
                        nw26_[int(0)] = d_158_v23_
                        nw26_[int(1)] = D0_DC0(d_157_v22_)
                        nw26_[int(2)] = d_158_v23_
                        nw26_[int(3)] = d_158_v23_
                        nw26_[int(4)] = d_158_v23_
                        nw26_[int(5)] = d_158_v23_
                        nw26_[int(6)] = d_158_v23_
                        nw26_[int(7)] = d_158_v23_
                        nw26_[int(8)] = D0_DC0(d_157_v22_)
                        nw26_[int(9)] = d_158_v23_
                        nw26_[int(10)] = D0_DC0(d_157_v22_)
                        nw26_[int(11)] = d_158_v23_
                        nw26_[int(12)] = D0_DC0(_dafny.CodePoint('c'))
                        nw26_[int(13)] = d_158_v23_
                        nw26_[int(14)] = d_158_v23_
                        d_164_v28_ = nw26_
                        index26_ = default__.safeIndex(170, (d_164_v28_).length(0))
                        (d_164_v28_)[index26_] = D0_DC0(d_157_v22_)
                        index27_ = default__.safeIndex(578, (d_163_v27_).length(0))
                        index28_ = default__.safeIndex(170, (d_164_v28_).length(0))
                        rhs35_ = (d_130_v3_) + (d_130_v3_)
                        rhs36_ = d_158_v23_
                        lhs26_ = d_163_v27_
                        lhs27_ = default__.safeIndex(578, (d_163_v27_).length(0))
                        lhs28_ = d_164_v28_
                        lhs29_ = default__.safeIndex(170, (d_164_v28_).length(0))
                        lhs26_[lhs27_] = rhs35_
                        lhs28_[lhs29_] = rhs36_
                        (d_132_globalState_).f7 = d_127_v0_
                    elif True:
                        d_159_v24_ = d_159_v24_
                        d_165_v29_: C0
                        nw27_ = C0()
                        nw27_.ctor__()
                        d_165_v29_ = nw27_
                        d_157_v22_ = d_157_v22_
                        d_166_v30_: _dafny.Seq
                        d_166_v30_ = _dafny.SeqWithoutIsStrInference([d_133_v5_])
                        (d_132_globalState_).f8 = (d_166_v30_) != ((d_136_v8_).fm0(d_139_v10_, d_130_v3_, d_132_globalState_))
                        d_167_v31_: C0
                        nw28_ = C0()
                        nw28_.ctor__()
                        d_167_v31_ = nw28_
                    pass
            pass
        d_178_v32_: _dafny.Set
        d_178_v32_ = _dafny.Set({d_133_v5_})
        d_179_v33_: _dafny.Seq
        d_179_v33_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_133_v5_, d_139_v10_, d_139_v10_, d_139_v10_, d_139_v10_}), d_178_v32_, d_178_v32_])
        if (d_179_v33_) != ((d_179_v33_) + (_dafny.SeqWithoutIsStrInference([d_178_v32_]))):
            d_180_v34_: _dafny.Map
            d_180_v34_ = _dafny.Map({d_130_v3_: d_139_v10_})
            d_180_v34_ = d_180_v34_
            if not((d_178_v32_).issubset(default__.fm3(d_139_v10_, d_139_v10_, d_159_v24_, d_133_v5_, d_132_globalState_))):
                d_178_v32_ = d_178_v32_
                d_181_v35_: _dafny.Seq
                d_181_v35_ = _dafny.SeqWithoutIsStrInference([d_127_v0_, (0) - (d_127_v0_), d_127_v0_])
                (d_132_globalState_).f8 = default__.fm4(d_141_v12_, d_181_v35_, False, d_133_v5_, d_132_globalState_)
                d_182_v37_: _dafny.Array
                def lambda10_(d_183_v5_, d_184_v0_, d_185_v10_, d_186_v12_, d_187_v22_):
                    def lambda11_(d_188_i8_):
                        def iife5_():
                            coll1_ = _dafny.Map()
                            compr_1_: _dafny.Seq
                            for compr_1_ in (_dafny.SeqWithoutIsStrInference([(d_186_v12_).set(default__.safeIndex((0) - (d_184_v0_), len(d_186_v12_)), d_187_v22_), _dafny.SeqWithoutIsStrInference([d_187_v22_])])).Elements:
                                d_189_v36_: _dafny.Seq = compr_1_
                                if (d_189_v36_) in (_dafny.SeqWithoutIsStrInference([(d_186_v12_).set(default__.safeIndex((0) - (d_184_v0_), len(d_186_v12_)), d_187_v22_), _dafny.SeqWithoutIsStrInference([d_187_v22_])])):
                                    coll1_[d_189_v36_] = D0_DC1(not(d_183_v5_), d_184_v0_, d_183_v5_, len(_dafny.Map({d_187_v22_: len(_dafny.Map({False: False}))})))
                            return _dafny.Map(coll1_)
                        return ((D1_DC2(_dafny.Map({d_183_v5_: d_184_v0_}))).cf5) | (_dafny.Map({d_185_v10_: len(iife5_()
                        )}))

                    return lambda11_

                init5_ = lambda10_(d_133_v5_, d_127_v0_, d_139_v10_, d_141_v12_, d_157_v22_)
                nw29_ = _dafny.Array(None, 25)
                for i0_5_ in range(nw29_.length(0)):
                    nw29_[i0_5_] = init5_(i0_5_)
                d_182_v37_ = nw29_
                d_182_v37_ = d_182_v37_
                index29_ = default__.safeIndex(98, (d_137_v9_).length(0))
                (d_137_v9_)[index29_] = d_139_v10_
                index30_ = default__.safeIndex(98, (d_137_v9_).length(0))
                (d_137_v9_)[index30_] = d_133_v5_
                index31_ = default__.safeIndex(833, (d_129_v2_).length(0))
                (d_129_v2_)[index31_] = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngxstq"))) + (d_141_v12_)).set(default__.safeIndex(d_127_v0_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngxstq"))) + (d_141_v12_))), d_157_v22_))
                index32_ = default__.safeIndex(833, (d_129_v2_).length(0))
                (d_129_v2_)[index32_] = (0) - (d_127_v0_)
            elif True:
                d_190_v38_: _dafny.Array
                def lambda12_(d_191_v22_):
                    def lambda13_(d_192_i9_):
                        return d_191_v22_

                    return lambda13_

                init6_ = lambda12_(d_157_v22_)
                nw30_ = _dafny.Array(None, 16)
                for i0_6_ in range(nw30_.length(0)):
                    nw30_[i0_6_] = init6_(i0_6_)
                d_190_v38_ = nw30_
                index33_ = default__.safeIndex(867, (d_190_v38_).length(0))
                (d_190_v38_)[index33_] = d_157_v22_
                index34_ = default__.safeIndex(867, (d_190_v38_).length(0))
                (d_190_v38_)[index34_] = d_157_v22_
                (d_132_globalState_).f6 = ((0) - (len(_dafny.SeqWithoutIsStrInference([default__.fm2(_dafny.SeqWithoutIsStrInference([(d_190_v38_)[default__.safeIndex(867, (d_190_v38_).length(0))] for d_193_i10_ in range(default__.abs(895))]), d_139_v10_, ((d_180_v34_)[d_141_v12_] if (d_141_v12_) in (d_180_v34_) else default__.fm4(d_141_v12_, _dafny.SeqWithoutIsStrInference([d_127_v0_]), d_139_v10_, True, d_132_globalState_)), d_132_globalState_)]))) if True else d_127_v0_)
                d_194_v39_: _dafny.Seq
                d_194_v39_ = _dafny.SeqWithoutIsStrInference([d_130_v3_, (_dafny.SeqWithoutIsStrInference([(d_190_v38_)[default__.safeIndex(867, (d_190_v38_).length(0))] for d_195_i11_ in range(default__.abs(966))])) + (_dafny.SeqWithoutIsStrInference([(d_190_v38_)[default__.safeIndex(867, (d_190_v38_).length(0))] for d_196_i12_ in range(default__.abs(-219))]))])
                d_194_v39_ = _dafny.SeqWithoutIsStrInference([d_130_v3_, d_130_v3_, (_dafny.SeqWithoutIsStrInference([d_157_v22_ for d_197_i13_ in range(default__.abs(279))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngtw"))), _dafny.SeqWithoutIsStrInference([(d_190_v38_)[default__.safeIndex(867, (d_190_v38_).length(0))] for d_198_i14_ in range(default__.abs(782))]), d_141_v12_])
                d_199_v40_: _dafny.Seq
                d_199_v40_ = _dafny.SeqWithoutIsStrInference([not(d_133_v5_)])
                (d_132_globalState_).f8 = ((d_178_v32_) != (_dafny.Set({not(d_139_v10_)}))) in ((d_199_v40_) + (d_199_v40_))
                d_200_v41_: _dafny.Array
                nw31_ = _dafny.Array(_dafny.MultiSet({}), 2)
                d_200_v41_ = nw31_
                d_201_v42_: _dafny.MultiSet
                d_201_v42_ = _dafny.MultiSet([d_128_v1_])
                index35_ = default__.safeIndex(706, (d_200_v41_).length(0))
                (d_200_v41_)[index35_] = (d_201_v42_) - (_dafny.MultiSet([d_128_v1_, d_128_v1_]))
                index36_ = default__.safeIndex(706, (d_200_v41_).length(0))
                (d_200_v41_)[index36_] = default__.fm5(330, (d_139_v10_) not in (d_178_v32_), (d_133_v5_) and (True), d_132_globalState_)
            d_202_v43_: bool
            d_203_v44_: _dafny.Set
            d_204_v45_: _dafny.Seq
            out6_: bool
            out7_: _dafny.Set
            out8_: _dafny.Seq
            out6_, out7_, out8_ = default__.m0(d_132_globalState_)
            d_202_v43_ = out6_
            d_203_v44_ = out7_
            d_204_v45_ = out8_
            d_202_v43_ = d_133_v5_
            d_205_v46_: _dafny.Array
            nw32_ = _dafny.Array(D0.default()(), 12)
            d_205_v46_ = nw32_
            index37_ = default__.safeIndex(27, (d_205_v46_).length(0))
            (d_205_v46_)[index37_] = d_159_v24_
            d_206_v47_: _dafny.Seq
            d_206_v47_ = _dafny.SeqWithoutIsStrInference([d_127_v0_])
            d_207_v48_: _dafny.Seq
            d_207_v48_ = _dafny.SeqWithoutIsStrInference([default__.fm4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okbhypfhd")), d_206_v47_, d_133_v5_, d_133_v5_, d_132_globalState_)])
            d_208_v49_: _dafny.Seq
            d_208_v49_ = _dafny.SeqWithoutIsStrInference([d_137_v9_, d_137_v9_])
            index38_ = default__.safeIndex(27, (d_205_v46_).length(0))
            rhs37_ = ((len(d_204_v45_)) * ((_dafny.MultiSet(d_207_v48_)).cardinality)) - (((d_128_v1_)[d_127_v0_] if (d_127_v0_) in (d_128_v1_) else default__.fm6(899, d_133_v5_, d_132_globalState_)))
            rhs38_ = _dafny.CodePoint('t')
            rhs39_ = not (d_202_v43_) or (d_202_v43_)
            rhs40_ = D0_DC1(d_202_v43_, d_127_v0_, d_133_v5_, d_127_v0_)
            rhs41_ = (d_208_v49_)[default__.safeIndex(d_127_v0_, len(d_208_v49_))]
            lhs30_ = d_132_globalState_
            lhs31_ = d_205_v46_
            lhs32_ = default__.safeIndex(27, (d_205_v46_).length(0))
            lhs30_.f7 = rhs37_
            d_157_v22_ = rhs38_
            d_133_v5_ = rhs39_
            lhs31_[lhs32_] = rhs40_
            d_137_v9_ = rhs41_
        elif True:
            (d_132_globalState_).f6 = (default__.fm6(default__.fm6(d_127_v0_, d_139_v10_, d_132_globalState_), d_139_v10_, d_132_globalState_)) + ((d_127_v0_ if d_133_v5_ else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nowxxf")))))
            index39_ = default__.safeIndex(856, (d_137_v9_).length(0))
            (d_137_v9_)[index39_] = not (d_139_v10_) or (d_133_v5_)
            index40_ = default__.safeIndex(856, (d_137_v9_).length(0))
            (d_137_v9_)[index40_] = False
            d_209_v50_: _dafny.Array
            def lambda14_(d_210_v9_, d_211_v5_):
                def lambda15_(d_212_i15_):
                    return _dafny.SeqWithoutIsStrInference([(d_210_v9_)[default__.safeIndex(856, (d_210_v9_).length(0))], (d_210_v9_)[default__.safeIndex(856, (d_210_v9_).length(0))], d_211_v5_])

                return lambda15_

            init7_ = lambda14_(d_137_v9_, d_133_v5_)
            nw33_ = _dafny.Array(None, 22)
            for i0_7_ in range(nw33_.length(0)):
                nw33_[i0_7_] = init7_(i0_7_)
            d_209_v50_ = nw33_
            index41_ = default__.safeIndex(762, (d_209_v50_).length(0))
            (d_209_v50_)[index41_] = _dafny.SeqWithoutIsStrInference([d_139_v10_])
            d_213_v51_: _dafny.Seq
            d_213_v51_ = _dafny.SeqWithoutIsStrInference([d_139_v10_])
            index42_ = default__.safeIndex(762, (d_209_v50_).length(0))
            (d_209_v50_)[index42_] = (d_213_v51_) + (d_213_v51_)
            d_214_v52_: _dafny.Seq
            d_214_v52_ = _dafny.SeqWithoutIsStrInference([d_136_v8_, d_136_v8_])
            d_214_v52_ = d_214_v52_
            (d_132_globalState_).f3 = (0) - (d_127_v0_)
        d_215_v53_: _dafny.Array
        nw34_ = _dafny.Array(_dafny.Array(None, 0), 8)
        d_215_v53_ = nw34_
        d_216_v54_: _dafny.Array
        nw35_ = _dafny.Array(_dafny.Seq({}), 18)
        d_216_v54_ = nw35_
        d_217_v55_: _dafny.Seq
        d_217_v55_ = _dafny.SeqWithoutIsStrInference([d_136_v8_])
        index43_ = default__.safeIndex(689, (d_216_v54_).length(0))
        (d_216_v54_)[index43_] = (d_217_v55_ if d_139_v10_ else d_217_v55_)
        d_218_v56_: _dafny.Map
        d_218_v56_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_157_v22_, d_157_v22_])): d_127_v0_})
        d_219_v57_: _dafny.Seq
        d_219_v57_ = _dafny.SeqWithoutIsStrInference([d_127_v0_])
        pat_let_tv3_ = d_139_v10_
        pat_let_tv4_ = d_140_v11_
        pat_let_tv5_ = d_140_v11_
        index44_ = default__.safeIndex(689, (d_216_v54_).length(0))
        def lambda16_(source4_):
            if source4_.is_DC3:
                d_220___mcc_h5_ = source4_.cf6
                d_221___mcc_h6_ = source4_.cf7
                d_222_cf7_ = d_221___mcc_h6_
                d_223_cf6_ = d_220___mcc_h5_
                return pat_let_tv3_
            elif True:
                d_224___mcc_h7_ = source4_.cf5
                d_225_cf5_ = d_224___mcc_h7_
                return (pat_let_tv4_).isdisjoint(pat_let_tv5_)

        rhs42_ = lambda16_(default__.fm7(d_133_v5_, d_133_v5_, d_132_globalState_))
        rhs43_ = d_215_v53_
        rhs44_ = d_133_v5_
        rhs45_ = (((d_217_v55_) + (d_217_v55_)) + (d_217_v55_)).set(default__.safeIndex(d_127_v0_, len(((d_217_v55_) + (d_217_v55_)) + (d_217_v55_))), d_136_v8_)
        rhs46_ = ((_dafny.SeqWithoutIsStrInference([d_127_v0_])) <= (d_219_v57_) if default__.fm4(default__.fm1(d_127_v0_, d_139_v10_, d_132_globalState_), default__.fm8(d_139_v10_, (d_218_v56_).set(316, d_127_v0_), d_127_v0_, d_132_globalState_), d_133_v5_, default__.fm4(d_141_v12_, d_219_v57_, d_139_v10_, False, d_132_globalState_), d_132_globalState_) else not(d_133_v5_))
        lhs33_ = d_216_v54_
        lhs34_ = default__.safeIndex(689, (d_216_v54_).length(0))
        d_133_v5_ = rhs42_
        d_215_v53_ = rhs43_
        d_139_v10_ = rhs44_
        lhs33_[lhs34_] = rhs45_
        d_133_v5_ = rhs46_
        d_226_v58_: bool
        d_227_v59_: _dafny.Set
        d_228_v60_: _dafny.Seq
        out9_: bool
        out10_: _dafny.Set
        out11_: _dafny.Seq
        out9_, out10_, out11_ = default__.m0(d_132_globalState_)
        d_226_v58_ = out9_
        d_227_v59_ = out10_
        d_228_v60_ = out11_
        d_127_v0_ = (0) - (d_127_v0_)
        d_229_v61_: C0
        nw36_ = C0()
        nw36_.ctor__()
        d_229_v61_ = nw36_
        d_230_v62_: _dafny.MultiSet
        d_230_v62_ = _dafny.MultiSet([d_226_v58_])
        d_231_v63_: D1
        d_231_v63_ = D1_DC3((d_230_v62_).cardinality, d_226_v58_)
        d_232_v64_: _dafny.Array
        nw37_ = _dafny.Array(None, 6)
        nw37_[int(0)] = d_178_v32_
        nw37_[int(1)] = d_178_v32_
        nw37_[int(2)] = (d_178_v32_).intersection((d_179_v33_)[default__.safeIndex(d_127_v0_, len(d_179_v33_))])
        nw37_[int(3)] = d_178_v32_
        nw37_[int(4)] = d_178_v32_
        nw37_[int(5)] = (d_178_v32_) - (default__.fm3(d_139_v10_, not(d_133_v5_), d_159_v24_, (d_231_v63_).cf7, d_132_globalState_))
        d_232_v64_ = nw37_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_232_v64_).length(0)):
            d_233_i16_: int = guard_loop_1_
            if (True) and (((0) <= (d_233_i16_)) and ((d_233_i16_) < ((d_232_v64_).length(0)))):
                (d_232_v64_)[(d_233_i16_)] = d_178_v32_
        _dafny.print(_dafny.string_of(d_127_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_v1_) == (_dafny.MultiSet([-183]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v2_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_130_v3_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v4_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwijw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqq"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_132_globalState_).f0) == (_dafny.MultiSet([-183]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_132_globalState_).f1)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_132_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_132_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_132_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_132_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_132_globalState_).f10) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwijw")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqq"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_133_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v6_) == (_dafny.Map({_dafny.Map({-183: True}): -183, _dafny.Map({-421: False}): 183}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v7_) == (_dafny.Map({-421: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v9_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_139_v10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_140_v11_) == (_dafny.Set({0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_141_v12_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_142_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_148_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_153_i6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_157_v22_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v23_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v24_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v24_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v24_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v24_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_178_v32_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v33_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({False}), _dafny.Set({False}), _dafny.Set({False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_216_v54_)[5])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_217_v55_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v56_) == (_dafny.Map({2: -183}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v57_) == (_dafny.SeqWithoutIsStrInference([-183]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_226_v58_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_227_v59_) == (_dafny.Set({0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_228_v60_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_230_v62_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v63_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v63_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v64_)[0]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v64_)[1]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v64_)[2]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v64_)[3]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v64_)[4]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_232_v64_)[5]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
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
        return lambda: D1_DC3(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5(int(0), int(0), D0.default()())
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

class D2_DC5(D2, NamedTuple('DC5', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC8(D3, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f3: int = int(0)
        self.f6: int = int(0)
        self.f7: int = int(0)
        self.f8: bool = False
        self._f0: _dafny.MultiSet = _dafny.MultiSet({})
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f2: bool = False
        self._f4: int = int(0)
        self._f5: int = int(0)
        self._f9: bool = False
        self._f10: _dafny.Set = _dafny.Set({})
        self._f11: bool = False
        self._f12: int = int(0)
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self).f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
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
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
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

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm0(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True]) if False else _dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True])))

