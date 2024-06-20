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
        return (len((_dafny.SeqWithoutIsStrInference([False, False, False])) + ((_dafny.SeqWithoutIsStrInference([True]))))) > ((0) - ((len(_dafny.Set({True, not(False)}))) + (559)))

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        return ((_dafny.Set({False})).intersection(_dafny.Set({False, False, False, False, True}))) | ((_dafny.Set({not(False)})) | (_dafny.Set({True, not(True)})))

    @staticmethod
    def fm3(globalState):
        return (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uybruf")))), 695, -92, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txqeg"))), 559])) + (_dafny.SeqWithoutIsStrInference([657, -678]))

    @staticmethod
    def fm4(p0, p1, globalState):
        return _dafny.Map({(0) - (-387): (0) - (default__.safeDivisionInt(93, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuqnfvfjl")))))})

    @staticmethod
    def fm5(p0, globalState):
        return D0_DC0(_dafny.Map({not(not(False)): False}))

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return (_dafny.Map({(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trfmv"))))): not(True)})) | (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ataulj")))): True}))

    @staticmethod
    def fm7(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_0_i0_ in range(default__.abs(571))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mclcxtfx"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1_i1_ in range(default__.abs(3))])))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return default__.safeDivisionInt(161, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvcutvx")) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "entn")))))

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        return ((_dafny.Map({not(False): True}) if not(False) else _dafny.Map({False: not(False)}))) | ((_dafny.Map({not(True): True})) | (_dafny.Map({True: False})))

    @staticmethod
    def fm10(p0, p1, globalState):
        if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehe"))):
            return _dafny.CodePoint('h')
        elif True:
            return _dafny.CodePoint('i')

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        d_2_v0_: bool
        d_2_v0_ = False
        d_3_i0_: int
        d_3_i0_ = 0
        with _dafny.label("0"):
            while d_2_v0_:
                with _dafny.c_label("0"):
                    if (d_3_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_3_i0_ = (d_3_i0_) + (1)
                    d_4_v1_: C0
                    nw0_ = C0()
                    nw0_.ctor__(p1)
                    d_4_v1_ = nw0_
                    d_5_v2_: D3
                    d_5_v2_ = D3_DC8(D3_DC6(d_4_v1_))
                    source0_ = (d_5_v2_ if True else d_5_v2_)
                    if source0_.is_DC7:
                        d_6_v3_: _dafny.Set
                        d_6_v3_ = _dafny.Set({d_2_v0_, d_2_v0_})
                        d_7_v4_: _dafny.Set
                        d_7_v4_ = _dafny.Set({d_2_v0_})
                        (globalState).f6 = ((d_7_v4_)).issubset(d_6_v3_)
                        d_8_v5_: C0
                        nw1_ = C0()
                        nw1_.ctor__((d_4_v1_).f22)
                        d_8_v5_ = nw1_
                        d_9_v6_: _dafny.Seq
                        d_9_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ack"))
                        d_10_v7_: D0
                        d_10_v7_ = D0_DC1(len(d_9_v6_), d_2_v0_)
                        d_11_v8_: _dafny.Seq
                        d_11_v8_ = _dafny.SeqWithoutIsStrInference([((d_10_v7_).cf1) >= (693), (d_6_v3_).isdisjoint(d_6_v3_)])
                        d_11_v8_ = (d_11_v8_) + (_dafny.SeqWithoutIsStrInference([False, d_2_v0_]))
                        (globalState).f6 = not(d_2_v0_)
                    elif source0_.is_DC6:
                        d_12___mcc_h0_ = source0_.cf8
                        d_13_cf8_ = d_12___mcc_h0_
                        d_14_v9_: _dafny.Seq
                        d_14_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyy"))
                        d_14_v9_ = d_14_v9_
                        (globalState).f6 = not(d_2_v0_)
                        d_15_v10_: _dafny.Map
                        d_15_v10_ = _dafny.Map({-222: d_2_v0_})
                        d_16_v11_: _dafny.Map
                        d_16_v11_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")): (d_13_cf8_).fm1(((d_15_v10_)[p2] if (p2) in (d_15_v10_) else d_2_v0_), _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqfg"))) for d_17_i1_ in range(default__.abs(984))]), globalState)})
                        d_18_v12_: str
                        d_18_v12_ = _dafny.CodePoint('m')
                        d_16_v11_ = (d_16_v11_).set(d_14_v9_, (d_18_v12_) not in (d_14_v9_))
                        d_19_v13_: C0
                        nw2_ = C0()
                        nw2_.ctor__(((d_4_v1_).f22).set(d_2_v0_, p0))
                        d_19_v13_ = nw2_
                    elif True:
                        d_20___mcc_h1_ = source0_.cf9
                        d_21_cf9_ = d_20___mcc_h1_
                        d_22_v14_: _dafny.Array
                        nw3_ = _dafny.Array(_dafny.Array(None, 0), 9)
                        d_22_v14_ = nw3_
                        index0_ = default__.safeIndex(703, (d_22_v14_).length(0))
                        (d_22_v14_)[index0_] = p0
                        index1_ = default__.safeIndex(703, (d_22_v14_).length(0))
                        (d_22_v14_)[index1_] = p0
                        d_23_v15_: _dafny.Seq
                        d_23_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
                        d_24_v16_: _dafny.Seq
                        d_24_v16_ = _dafny.SeqWithoutIsStrInference([d_4_v1_, d_4_v1_, d_4_v1_, d_4_v1_, d_4_v1_])
                        d_25_v17_: _dafny.Map
                        d_25_v17_ = _dafny.Map({False: len(d_23_v15_)})
                        d_26_v18_: D3
                        d_26_v18_ = D3_DC7()
                        d_27_v19_: _dafny.Map
                        d_27_v19_ = _dafny.Map({d_2_v0_: d_26_v18_})
                        d_28_v20_: _dafny.Set
                        d_28_v20_ = _dafny.Set({d_2_v0_, d_2_v0_})
                        d_29_v21_: str
                        d_29_v21_ = _dafny.CodePoint('t')
                        d_30_v22_: _dafny.Array
                        nw4_ = _dafny.Array(None, 26)
                        nw4_[int(0)] = 740
                        nw4_[int(1)] = p2
                        nw4_[int(2)] = p2
                        nw4_[int(3)] = p2
                        nw4_[int(4)] = (992) - (p2)
                        nw4_[int(5)] = default__.safeDivisionInt(p2, p2)
                        nw4_[int(6)] = p2
                        nw4_[int(7)] = default__.safeDivisionInt(p2, p2)
                        nw4_[int(8)] = (0) - (-561)
                        nw4_[int(9)] = p2
                        nw4_[int(10)] = len(d_24_v16_)
                        nw4_[int(11)] = default__.safeModuloInt(p2, p2)
                        nw4_[int(12)] = (0) - (len(_dafny.SeqWithoutIsStrInference([True])))
                        nw4_[int(13)] = p2
                        nw4_[int(14)] = (0) - (len(d_25_v17_))
                        nw4_[int(15)] = (p2) - (p2)
                        nw4_[int(16)] = len((d_27_v19_) | (_dafny.Map({d_2_v0_: D3_DC7()})))
                        nw4_[int(17)] = p2
                        nw4_[int(18)] = (0) - (p2)
                        nw4_[int(19)] = p2
                        nw4_[int(20)] = p2
                        nw4_[int(21)] = (0) - ((p2) * (p2))
                        nw4_[int(22)] = (p2) * (len(d_28_v20_))
                        nw4_[int(23)] = p2
                        nw4_[int(24)] = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eybkhsaop"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddyvcpun")))).set(default__.safeIndex(p2, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eybkhsaop"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddyvcpun"))))), d_29_v21_))
                        nw4_[int(25)] = -209
                        d_30_v22_ = nw4_
                        index2_ = default__.safeIndex(783, (d_30_v22_).length(0))
                        (d_30_v22_)[index2_] = p2
                        d_31_v23_: _dafny.MultiSet
                        d_31_v23_ = _dafny.MultiSet([d_4_v1_])
                        d_32_v24_: D3
                        d_32_v24_ = D3_DC6(d_4_v1_)
                        d_33_v25_: _dafny.Map
                        d_33_v25_ = _dafny.Map({d_32_v24_: d_2_v0_})
                        index3_ = default__.safeIndex(783, (d_30_v22_).length(0))
                        rhs0_ = ((d_31_v23_).intersection(d_31_v23_)) == (_dafny.MultiSet([d_4_v1_]))
                        rhs1_ = d_23_v15_
                        rhs2_ = (len(d_33_v25_) if d_2_v0_ else (p2) * (p2))
                        rhs3_ = 650
                        lhs0_ = d_30_v22_
                        lhs1_ = default__.safeIndex(783, (d_30_v22_).length(0))
                        lhs2_ = globalState
                        r0 = rhs0_
                        d_23_v15_ = rhs1_
                        lhs0_[lhs1_] = rhs2_
                        lhs2_.f7 = rhs3_
                        d_34_v26_: _dafny.Seq
                        d_34_v26_ = _dafny.SeqWithoutIsStrInference([d_23_v15_])
                        d_35_v27_: _dafny.MultiSet
                        d_35_v27_ = _dafny.MultiSet([d_2_v0_, not(d_2_v0_), d_2_v0_, not(d_2_v0_), ((d_4_v1_).fm1(d_2_v0_, d_34_v26_, globalState) if d_2_v0_ else d_2_v0_)])
                        rhs4_ = d_35_v27_
                        rhs5_ = (d_30_v22_)[default__.safeIndex(783, (d_30_v22_).length(0))]
                        lhs3_ = globalState
                        d_35_v27_ = rhs4_
                        lhs3_.f20 = rhs5_
                        d_36_v28_: _dafny.Seq
                        d_36_v28_ = _dafny.SeqWithoutIsStrInference([False, d_2_v0_, d_2_v0_, d_2_v0_])
                        index4_ = default__.safeIndex(73, (p0).length(0))
                        (p0)[index4_] = (d_2_v0_) not in (d_36_v28_)
                        index5_ = default__.safeIndex(73, (p0).length(0))
                        (p0)[index5_] = d_2_v0_
                    d_37_v29_: _dafny.Array
                    nw5_ = _dafny.Array(int(0), 15)
                    d_37_v29_ = nw5_
                    d_38_v30_: _dafny.MultiSet
                    d_38_v30_ = _dafny.MultiSet([d_2_v0_])
                    d_39_v31_: _dafny.Seq
                    d_39_v31_ = _dafny.SeqWithoutIsStrInference([d_4_v1_])
                    d_40_v32_: _dafny.Seq
                    d_40_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nb"))
                    index6_ = default__.safeIndex(202, (d_37_v29_).length(0))
                    (d_37_v29_)[index6_] = (((d_38_v30_)[d_2_v0_] if (d_2_v0_) in (d_38_v30_) else len(d_39_v31_))) * (default__.fm8(p2, len(d_40_v32_), d_2_v0_, d_2_v0_, globalState))
                    d_41_v33_: _dafny.Map
                    d_41_v33_ = _dafny.Map({d_4_v1_: d_2_v0_})
                    d_42_v34_: str
                    d_42_v34_ = _dafny.CodePoint('p')
                    index7_ = default__.safeIndex(202, (d_37_v29_).length(0))
                    rhs6_ = d_2_v0_
                    rhs7_ = (not (not(d_2_v0_)) or (d_2_v0_)) == (d_2_v0_)
                    rhs8_ = ((len(default__.fm9(p2, p2, d_2_v0_, globalState))) * (p2)) + (p2)
                    rhs9_ = ((0) - (p2) if (d_40_v32_) != ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "necjelxnj"))).set(default__.safeIndex(len(d_41_v33_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "necjelxnj")))), d_42_v34_)) else p2)
                    rhs10_ = (p2) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))) * (p2))
                    lhs4_ = globalState
                    lhs5_ = globalState
                    lhs6_ = globalState
                    lhs7_ = d_37_v29_
                    lhs8_ = default__.safeIndex(202, (d_37_v29_).length(0))
                    lhs9_ = globalState
                    lhs4_.f6 = rhs6_
                    lhs5_.f4 = rhs7_
                    lhs6_.f14 = rhs8_
                    lhs7_[lhs8_] = rhs9_
                    lhs9_.f7 = rhs10_
                    d_2_v0_ = d_2_v0_
                    d_43_v35_: _dafny.Seq
                    d_43_v35_ = _dafny.SeqWithoutIsStrInference([d_2_v0_, d_2_v0_])
                    d_44_v36_: _dafny.Set
                    d_44_v36_ = _dafny.Set({d_2_v0_, not((d_43_v35_)[default__.safeIndex((0) - ((d_37_v29_)[default__.safeIndex(202, (d_37_v29_).length(0))]), len(d_43_v35_))])})
                    d_45_v37_: _dafny.Map
                    d_45_v37_ = _dafny.Map({d_2_v0_: (d_44_v36_) == (d_44_v36_)})
                    d_46_v38_: _dafny.Seq
                    d_46_v38_ = _dafny.SeqWithoutIsStrInference([default__.fm8(len(d_40_v32_), p2, d_2_v0_, d_2_v0_, globalState)])
                    d_47_v39_: D1
                    d_47_v39_ = D1_DC2(p1)
                    d_48_v40_: D1
                    d_48_v40_ = D1_DC4(d_47_v39_)
                    d_49_v41_: _dafny.Seq
                    d_49_v41_ = _dafny.SeqWithoutIsStrInference([d_48_v40_])
                    d_50_v42_: _dafny.Set
                    d_50_v42_ = _dafny.Set({len(d_49_v41_)})
                    d_51_v43_: _dafny.Map
                    d_51_v43_ = _dafny.Map({not(d_2_v0_): _dafny.CodePoint('n')})
                    d_52_v44_: _dafny.MultiSet
                    d_52_v44_ = _dafny.MultiSet([d_51_v43_, d_51_v43_])
                    rhs11_ = d_4_v1_
                    rhs12_ = (d_45_v37_) | ((d_45_v37_).set(d_2_v0_, d_2_v0_))
                    rhs13_ = (d_37_v29_)[default__.safeIndex(202, (d_37_v29_).length(0))]
                    rhs14_ = default__.fm10(d_50_v42_, (d_52_v44_) - (d_52_v44_), globalState)
                    rhs15_ = d_46_v38_
                    lhs10_ = globalState
                    d_4_v1_ = rhs11_
                    d_45_v37_ = rhs12_
                    lhs10_.f14 = rhs13_
                    d_42_v34_ = rhs14_
                    d_46_v38_ = rhs15_
                    pass
            pass
        r1 = d_2_v0_
        (globalState).f20 = p2
        d_53_v45_: _dafny.Seq
        d_53_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqgwm"))
        d_54_v46_: _dafny.Seq
        d_54_v46_ = _dafny.SeqWithoutIsStrInference([(d_53_v45_) == (d_53_v45_), (p2) < ((0) - (p2)), (d_2_v0_) or (d_2_v0_), d_2_v0_, d_2_v0_])
        d_55_v48_: _dafny.Seq
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(209, 925):
                d_56_v47_: int = compr_0_
                if ((209) <= (d_56_v47_)) and ((d_56_v47_) < (925)):
                    coll0_[default__.safeModuloInt(d_56_v47_, p2)] = d_2_v0_
            return _dafny.Map(coll0_)
        d_55_v48_ = _dafny.SeqWithoutIsStrInference([len(iife0_()
        )])
        r0 = (d_54_v46_)[default__.safeIndex(default__.fm8(len(d_55_v48_), p2, d_2_v0_, False, globalState), len(d_54_v46_))]
        d_57_v49_: _dafny.Array
        nw6_ = _dafny.Array(False, 1)
        d_57_v49_ = nw6_
        d_57_v49_ = p0
        d_58_v50_: _dafny.Map
        d_58_v50_ = _dafny.Map({d_57_v49_: (d_55_v48_)[default__.safeIndex(p2, len(d_55_v48_))]})
        d_58_v50_ = (d_58_v50_) | (d_58_v50_)
        d_59_v51_: D0
        d_59_v51_ = D0_DC1(p2, d_2_v0_)
        pat_let_tv0_ = d_2_v0_
        def lambda0_(source1_):
            if source1_.is_DC1:
                d_60___mcc_h2_ = source1_.cf1
                d_61___mcc_h3_ = source1_.cf2
                d_62_cf2_ = d_61___mcc_h3_
                d_63_cf1_ = d_60___mcc_h2_
                return pat_let_tv0_
            elif True:
                d_64___mcc_h4_ = source1_.cf0
                d_65_cf0_ = d_64___mcc_h4_
                return False

        r0 = lambda0_(d_59_v51_)
        d_66_v52_: _dafny.MultiSet
        d_66_v52_ = _dafny.MultiSet([p2, p2])
        r1 = (d_66_v52_).isdisjoint(d_66_v52_)
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_67_v0_: str
        d_67_v0_ = _dafny.CodePoint('w')
        d_68_v1_: int
        d_68_v1_ = 557
        d_69_v2_: _dafny.Array
        nw7_ = _dafny.Array(int(0), 4)
        d_69_v2_ = nw7_
        d_70_v3_: _dafny.Map
        d_70_v3_ = _dafny.Map({d_69_v2_: d_68_v1_})
        d_71_v4_: _dafny.Set
        d_71_v4_ = _dafny.Set({len(d_70_v3_), d_68_v1_})
        d_72_v5_: _dafny.MultiSet
        d_72_v5_ = _dafny.MultiSet([d_71_v4_])
        d_73_v6_: bool
        d_73_v6_ = True
        d_74_v7_: _dafny.MultiSet
        d_74_v7_ = _dafny.MultiSet([d_73_v6_])
        d_75_v8_: _dafny.Map
        d_75_v8_ = _dafny.Map({d_73_v6_: d_68_v1_})
        d_76_v9_: _dafny.Array
        nw8_ = _dafny.Array(None, 22)
        nw8_[int(0)] = d_68_v1_
        nw8_[int(1)] = (d_72_v5_).cardinality
        nw8_[int(2)] = len(_dafny.SeqWithoutIsStrInference([d_67_v0_ for d_77_i1_ in range(default__.abs(-906))]))
        nw8_[int(3)] = d_68_v1_
        nw8_[int(4)] = d_68_v1_
        nw8_[int(5)] = d_68_v1_
        nw8_[int(6)] = d_68_v1_
        nw8_[int(7)] = d_68_v1_
        nw8_[int(8)] = len(_dafny.Map({d_68_v1_: d_73_v6_}))
        nw8_[int(9)] = d_68_v1_
        nw8_[int(10)] = d_68_v1_
        nw8_[int(11)] = d_68_v1_
        nw8_[int(12)] = d_68_v1_
        nw8_[int(13)] = d_68_v1_
        nw8_[int(14)] = d_68_v1_
        nw8_[int(15)] = d_68_v1_
        nw8_[int(16)] = d_68_v1_
        nw8_[int(17)] = ((d_74_v7_)[d_73_v6_] if (d_73_v6_) in (d_74_v7_) else d_68_v1_)
        nw8_[int(18)] = len(d_75_v8_)
        nw8_[int(19)] = d_68_v1_
        nw8_[int(20)] = d_68_v1_
        nw8_[int(21)] = 875
        d_76_v9_ = nw8_
        d_78_v10_: _dafny.Map
        d_78_v10_ = _dafny.Map({d_73_v6_: not(d_73_v6_)})
        d_79_v11_: _dafny.Set
        d_79_v11_ = _dafny.Set({d_75_v8_, d_75_v8_, d_75_v8_})
        d_80_v12_: _dafny.Array
        nw9_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
        d_80_v12_ = nw9_
        d_81_v13_: _dafny.Set
        d_81_v13_ = _dafny.Set({d_73_v6_})
        d_82_v14_: _dafny.Seq
        d_82_v14_ = _dafny.SeqWithoutIsStrInference([d_81_v13_])
        d_83_v15_: _dafny.Seq
        d_83_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "je"))
        d_84_globalState_: GlobalState
        nw10_ = GlobalState()
        nw10_.ctor__(_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_85_i0_ in range(default__.abs(-417))]): d_67_v0_}), d_76_v9_, d_78_v10_, d_79_v11_, True, d_80_v12_, True, 929, d_74_v7_, 319, -434, True, (d_71_v4_).intersection(d_71_v4_), d_76_v9_, -685, True, ((d_82_v14_)[default__.safeIndex(d_68_v1_, len(d_82_v14_))]) | (_dafny.Set({d_73_v6_})), d_69_v2_, (_dafny.Map({False: len(d_83_v15_)})) | (_dafny.Map({True: d_68_v1_})), 581, 37, d_83_v15_)
        d_84_globalState_ = nw10_
        if True:
            index8_ = default__.safeIndex(196, (d_69_v2_).length(0))
            (d_69_v2_)[index8_] = (0) - (d_68_v1_)
            index9_ = default__.safeIndex(196, (d_69_v2_).length(0))
            (d_69_v2_)[index9_] = d_68_v1_
            d_75_v8_ = (d_75_v8_).set((d_74_v7_).issubset(d_74_v7_), ((d_69_v2_)[default__.safeIndex(196, (d_69_v2_).length(0))]) * ((d_69_v2_)[default__.safeIndex(196, (d_69_v2_).length(0))]))
            d_86_v16_: _dafny.Array
            nw11_ = _dafny.Array(False, 19)
            d_86_v16_ = nw11_
            d_86_v16_ = d_86_v16_
            d_83_v15_ = _dafny.SeqWithoutIsStrInference([d_67_v0_ for d_87_i2_ in range(default__.abs(971))])
            d_73_v6_ = not (True) or (d_73_v6_)
        elif True:
            d_88_v17_: _dafny.Array
            def lambda1_(d_89_v15_):
                def lambda2_(d_90_i3_):
                    return (d_89_v15_) == (d_89_v15_)

                return lambda2_

            init0_ = lambda1_(d_83_v15_)
            nw12_ = _dafny.Array(None, 2)
            for i0_0_ in range(nw12_.length(0)):
                nw12_[i0_0_] = init0_(i0_0_)
            d_88_v17_ = nw12_
            index10_ = default__.safeIndex(675, (d_88_v17_).length(0))
            (d_88_v17_)[index10_] = d_73_v6_
            d_91_v18_: _dafny.MultiSet
            d_91_v18_ = _dafny.MultiSet([805, d_68_v1_])
            d_92_v19_: _dafny.Seq
            d_92_v19_ = _dafny.SeqWithoutIsStrInference([d_68_v1_])
            d_93_v20_: _dafny.MultiSet
            d_93_v20_ = _dafny.MultiSet([(((d_91_v18_)[d_68_v1_] if (d_68_v1_) in (d_91_v18_) else d_68_v1_)) - (d_68_v1_), d_68_v1_, (len(d_92_v19_)) + (d_68_v1_), d_68_v1_])
            index11_ = default__.safeIndex(475, (d_69_v2_).length(0))
            (d_69_v2_)[index11_] = default__.safeDivisionInt(d_68_v1_, d_68_v1_)
            d_94_v21_: _dafny.Map
            d_94_v21_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(d_83_v15_), d_68_v1_]): _dafny.MultiSet([d_68_v1_])})
            index12_ = default__.safeIndex(675, (d_88_v17_).length(0))
            index13_ = default__.safeIndex(475, (d_69_v2_).length(0))
            rhs16_ = d_73_v6_
            rhs17_ = ((d_94_v21_)[d_92_v19_] if (d_92_v19_) in (d_94_v21_) else d_93_v20_)
            rhs18_ = (-971) - (d_68_v1_)
            rhs19_ = d_68_v1_
            lhs11_ = d_88_v17_
            lhs12_ = default__.safeIndex(675, (d_88_v17_).length(0))
            lhs13_ = d_69_v2_
            lhs14_ = default__.safeIndex(475, (d_69_v2_).length(0))
            lhs15_ = d_84_globalState_
            lhs11_[lhs12_] = rhs16_
            d_93_v20_ = rhs17_
            lhs13_[lhs14_] = rhs18_
            lhs15_.f20 = rhs19_
            (d_84_globalState_).f6 = default__.fm0(_dafny.Map({d_68_v1_: (d_69_v2_)[default__.safeIndex(475, (d_69_v2_).length(0))]}), d_68_v1_, (d_88_v17_)[default__.safeIndex(675, (d_88_v17_).length(0))], d_73_v6_, d_84_globalState_)
            d_76_v9_ = d_76_v9_
            d_95_v22_: _dafny.Map
            d_95_v22_ = _dafny.Map({(d_69_v2_)[default__.safeIndex(475, (d_69_v2_).length(0))]: d_68_v1_})
            d_96_v23_: _dafny.Map
            d_96_v23_ = _dafny.Map({len(d_95_v22_): (d_88_v17_)[default__.safeIndex(675, (d_88_v17_).length(0))]})
            d_97_v24_: _dafny.Map
            d_97_v24_ = _dafny.Map({d_96_v23_: 446})
            d_73_v6_ = (d_97_v24_) != (d_97_v24_)
            d_81_v13_ = d_81_v13_
        d_98_v25_: _dafny.Array
        nw13_ = _dafny.Array(_dafny.Array(None, 0), 16)
        d_98_v25_ = nw13_
        d_99_v26_: _dafny.Array
        nw14_ = _dafny.Array(_dafny.CodePoint('D'), 25)
        d_99_v26_ = nw14_
        index14_ = default__.safeIndex(919, (d_98_v25_).length(0))
        (d_98_v25_)[index14_] = d_99_v26_
        index15_ = default__.safeIndex(919, (d_98_v25_).length(0))
        (d_98_v25_)[index15_] = d_99_v26_
        (d_84_globalState_).f4 = d_73_v6_
        d_100_v27_: _dafny.Map
        d_100_v27_ = _dafny.Map({d_73_v6_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfucpd"))})
        d_83_v15_ = ((d_100_v27_)[d_73_v6_] if (d_73_v6_) in (d_100_v27_) else d_83_v15_)
        (d_84_globalState_).f4 = (d_68_v1_) < (d_68_v1_)
        if not (d_73_v6_) or (not(d_73_v6_)):
            (d_84_globalState_).f14 = -584
            d_101_v28_: _dafny.Seq
            d_101_v28_ = _dafny.SeqWithoutIsStrInference([d_73_v6_, d_73_v6_, d_73_v6_])
            d_102_v29_: _dafny.Map
            d_102_v29_ = _dafny.Map({d_68_v1_: len(d_71_v4_)})
            d_103_v30_: _dafny.Array
            nw15_ = _dafny.Array(None, 11)
            nw15_[int(0)] = d_73_v6_
            nw15_[int(1)] = d_73_v6_
            nw15_[int(2)] = d_73_v6_
            nw15_[int(3)] = d_73_v6_
            nw15_[int(4)] = (d_101_v28_)[default__.safeIndex(d_68_v1_, len(d_101_v28_))]
            nw15_[int(5)] = d_73_v6_
            nw15_[int(6)] = d_73_v6_
            nw15_[int(7)] = d_73_v6_
            nw15_[int(8)] = default__.fm0(d_102_v29_, d_68_v1_, d_73_v6_, d_73_v6_, d_84_globalState_)
            nw15_[int(9)] = d_73_v6_
            nw15_[int(10)] = not(not(d_73_v6_))
            d_103_v30_ = nw15_
            d_104_v31_: _dafny.Map
            d_104_v31_ = _dafny.Map({d_73_v6_: d_103_v30_})
            d_105_v32_: _dafny.Map
            d_105_v32_ = _dafny.Map({False: d_74_v7_})
            d_106_v33_: _dafny.Map
            d_106_v33_ = _dafny.Map({_dafny.Set({default__.fm0(d_102_v29_, d_68_v1_, d_73_v6_, d_73_v6_, d_84_globalState_), False}): ((d_105_v32_)[True] if (True) in (d_105_v32_) else _dafny.MultiSet([d_73_v6_]))})
            d_107_v34_: bool
            d_108_v35_: bool
            out0_: bool
            out1_: bool
            out0_, out1_ = default__.m0(d_103_v30_, (d_104_v31_).set(True, d_103_v30_), len((d_106_v33_ if d_73_v6_ else d_106_v33_)), d_84_globalState_)
            d_107_v34_ = out0_
            d_108_v35_ = out1_
            d_109_v36_: _dafny.Map
            d_109_v36_ = _dafny.Map({d_76_v9_: False})
            d_110_v37_: _dafny.Seq
            d_110_v37_ = _dafny.SeqWithoutIsStrInference([(d_68_v1_) - ((_dafny.MultiSet([236])).cardinality), d_68_v1_, default__.safeModuloInt(len((d_109_v36_).set(d_76_v9_, d_107_v34_)), d_68_v1_)])
            d_111_v38_: D0
            d_111_v38_ = D0_DC0(d_78_v10_)
            d_112_v39_: _dafny.Map
            d_112_v39_ = _dafny.Map({d_68_v1_: d_108_v35_})
            d_113_v40_: _dafny.Map
            d_113_v40_ = _dafny.Map({len((d_111_v38_).cf0): d_112_v39_})
            d_110_v37_ = (d_110_v37_).set(default__.safeIndex(d_68_v1_, len(d_110_v37_)), len((d_113_v40_) | (_dafny.Map({d_68_v1_: d_112_v39_}))))
            d_114_v41_: C0
            nw16_ = C0()
            nw16_.ctor__(d_104_v31_)
            d_114_v41_ = nw16_
            d_115_v42_: C0
            nw17_ = C0()
            nw17_.ctor__(d_104_v31_)
            d_115_v42_ = nw17_
        elif True:
            (d_84_globalState_).f20 = default__.safeDivisionInt((d_68_v1_) - ((0) - (d_68_v1_)), d_68_v1_)
            d_116_v43_: _dafny.Array
            nw18_ = _dafny.Array(None, 7)
            nw18_[int(0)] = d_73_v6_
            nw18_[int(1)] = d_73_v6_
            nw18_[int(2)] = True
            nw18_[int(3)] = d_73_v6_
            nw18_[int(4)] = d_73_v6_
            nw18_[int(5)] = True
            nw18_[int(6)] = d_73_v6_
            d_116_v43_ = nw18_
            d_117_v44_: _dafny.Map
            d_117_v44_ = _dafny.Map({d_73_v6_: d_116_v43_})
            d_118_v45_: C0
            nw19_ = C0()
            nw19_.ctor__(d_117_v44_)
            d_118_v45_ = nw19_
            d_118_v45_ = d_118_v45_
            d_119_v46_: _dafny.Map
            d_119_v46_ = _dafny.Map({d_68_v1_: d_68_v1_})
            d_120_v47_: bool
            d_121_v48_: bool
            out2_: bool
            out3_: bool
            out2_, out3_ = default__.m0(d_116_v43_, (d_118_v45_).f22, ((d_119_v46_)[d_68_v1_] if (d_68_v1_) in (d_119_v46_) else d_68_v1_), d_84_globalState_)
            d_120_v47_ = out2_
            d_121_v48_ = out3_
            d_122_v49_: _dafny.Map
            d_122_v49_ = _dafny.Map({d_78_v10_: len(d_83_v15_)})
            (d_84_globalState_).f14 = ((d_122_v49_)[d_78_v10_] if (d_78_v10_) in (d_122_v49_) else d_68_v1_)
            d_123_v50_: _dafny.Map
            d_123_v50_ = _dafny.Map({(0) - (d_68_v1_): d_121_v48_})
            d_123_v50_ = (d_123_v50_) | (_dafny.Map({-604: d_120_v47_}))
        (d_84_globalState_).f14 = d_68_v1_
        d_124_v51_: _dafny.Array
        nw20_ = _dafny.Array(None, 23)
        nw20_[int(0)] = d_73_v6_
        nw20_[int(1)] = d_73_v6_
        nw20_[int(2)] = False
        nw20_[int(3)] = d_73_v6_
        nw20_[int(4)] = d_73_v6_
        nw20_[int(5)] = d_73_v6_
        nw20_[int(6)] = d_73_v6_
        nw20_[int(7)] = d_73_v6_
        nw20_[int(8)] = d_73_v6_
        nw20_[int(9)] = d_73_v6_
        nw20_[int(10)] = d_73_v6_
        nw20_[int(11)] = d_73_v6_
        nw20_[int(12)] = d_73_v6_
        nw20_[int(13)] = d_73_v6_
        nw20_[int(14)] = d_73_v6_
        nw20_[int(15)] = d_73_v6_
        nw20_[int(16)] = d_73_v6_
        nw20_[int(17)] = d_73_v6_
        nw20_[int(18)] = d_73_v6_
        nw20_[int(19)] = not(((d_78_v10_)[True] if (True) in (d_78_v10_) else d_73_v6_))
        nw20_[int(20)] = d_73_v6_
        nw20_[int(21)] = d_73_v6_
        nw20_[int(22)] = d_73_v6_
        d_124_v51_ = nw20_
        d_125_v52_: _dafny.Map
        d_125_v52_ = _dafny.Map({False: d_124_v51_})
        d_126_v53_: bool
        d_127_v54_: bool
        out4_: bool
        out5_: bool
        out4_, out5_ = default__.m0(d_124_v51_, d_125_v52_, (d_68_v1_) * (d_68_v1_), d_84_globalState_)
        d_126_v53_ = out4_
        d_127_v54_ = out5_
        d_128_v56_: _dafny.Seq
        d_128_v56_ = _dafny.SeqWithoutIsStrInference([d_78_v10_, d_78_v10_])
        d_129_v57_: _dafny.Map
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.Map
            for compr_1_ in (d_128_v56_).Elements:
                d_130_v55_: _dafny.Map = compr_1_
                if (d_130_v55_) in (d_128_v56_):
                    coll1_[d_130_v55_] = not(d_126_v53_)
            return _dafny.Map(coll1_)
        d_129_v57_ = _dafny.Map({iife1_()
        : d_126_v53_})
        d_131_v58_: _dafny.Map
        d_131_v58_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyot"))): d_68_v1_})
        d_132_v59_: _dafny.Map
        d_132_v59_ = _dafny.Map({_dafny.Map({d_127_v54_: d_127_v54_}): default__.fm0(d_131_v58_, d_68_v1_, d_126_v53_, False, d_84_globalState_)})
        (d_84_globalState_).f4 = ((d_129_v57_)[d_132_v59_] if (d_132_v59_) in (d_129_v57_) else default__.fm0(d_131_v58_, d_68_v1_, d_127_v54_, d_127_v54_, d_84_globalState_))
        (d_84_globalState_).f4 = d_127_v54_
        d_133_v60_: D1
        d_133_v60_ = D1_DC2(d_125_v52_)
        d_134_v61_: C0
        nw21_ = C0()
        nw21_.ctor__((d_133_v60_).cf3)
        d_134_v61_ = nw21_
        d_134_v61_ = d_134_v61_
        hi0_ = (0) - (d_68_v1_)
        for d_135_i4_ in range(d_68_v1_, hi0_):
            d_136_v62_: _dafny.Array
            nw22_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_136_v62_ = nw22_
            d_137_v63_: _dafny.Map
            d_137_v63_ = _dafny.Map({d_126_v53_: d_81_v13_})
            d_138_v64_: _dafny.Seq
            d_138_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_67_v0_ for d_139_i5_ in range(default__.abs(282))])])
            d_140_v65_: _dafny.Array
            nw23_ = _dafny.Array(None, 14)
            nw23_[int(0)] = _dafny.Set({d_127_v54_})
            nw23_[int(1)] = d_81_v13_
            nw23_[int(2)] = d_81_v13_
            nw23_[int(3)] = d_81_v13_
            nw23_[int(4)] = d_81_v13_
            nw23_[int(5)] = d_81_v13_
            nw23_[int(6)] = d_81_v13_
            nw23_[int(7)] = d_81_v13_
            nw23_[int(8)] = d_81_v13_
            nw23_[int(9)] = ((d_137_v63_)[d_127_v54_] if (d_127_v54_) in (d_137_v63_) else default__.fm2(d_127_v54_, (d_134_v61_).fm1(d_127_v54_, (d_138_v64_).set(default__.safeIndex(d_68_v1_, len(d_138_v64_)), d_83_v15_), d_84_globalState_), len(d_83_v15_), d_84_globalState_))
            nw23_[int(10)] = _dafny.Set({d_126_v53_})
            nw23_[int(11)] = d_81_v13_
            nw23_[int(12)] = d_81_v13_
            nw23_[int(13)] = d_81_v13_
            d_140_v65_ = nw23_
            index16_ = default__.safeIndex(118, (d_136_v62_).length(0))
            (d_136_v62_)[index16_] = d_140_v65_
            index17_ = default__.safeIndex(118, (d_136_v62_).length(0))
            (d_136_v62_)[index17_] = (d_140_v65_ if d_127_v54_ else d_140_v65_)
            d_141_v66_: _dafny.Array
            d_141_v66_ = d_69_v2_
            d_142_v67_: _dafny.MultiSet
            d_142_v67_ = _dafny.MultiSet([(d_141_v66_), d_76_v9_, d_76_v9_, d_69_v2_, d_76_v9_])
            d_143_v68_: _dafny.Seq
            d_143_v68_ = _dafny.SeqWithoutIsStrInference([not(True), d_127_v54_])
            d_144_v69_: _dafny.Map
            d_144_v69_ = _dafny.Map({d_142_v67_: d_143_v68_})
            d_144_v69_ = (d_144_v69_).set(_dafny.MultiSet([d_76_v9_, d_69_v2_]), (d_143_v68_) + (d_143_v68_))
            (d_84_globalState_).f4 = ((d_126_v53_) or ((d_134_v61_).fm1(True, d_138_v64_, d_84_globalState_)) if d_73_v6_ else d_126_v53_)
            d_145_v70_: D0
            d_145_v70_ = D0_DC0(d_78_v10_)
            pat_let_tv1_ = d_78_v10_
            d_146_v71_: _dafny.Seq
            def iife2_(_pat_let0_0):
                def iife3_(d_147_dt__update__tmp_h0_):
                    def iife4_(_pat_let1_0):
                        def iife5_(d_148_dt__update_hcf0_h0_):
                            return D0_DC0(d_148_dt__update_hcf0_h0_)
                        return iife5_(_pat_let1_0)
                    return iife4_(pat_let_tv1_)
                return iife3_(_pat_let0_0)
            d_146_v71_ = _dafny.SeqWithoutIsStrInference([D0_DC0(d_78_v10_), iife2_(d_145_v70_)])
            rhs20_ = 495
            rhs21_ = d_135_i4_
            rhs22_ = (_dafny.SeqWithoutIsStrInference([d_145_v70_])) + (d_146_v71_)
            rhs23_ = ((d_68_v1_) + (d_135_i4_)) * ((d_68_v1_) * (d_135_i4_))
            rhs24_ = d_68_v1_
            lhs16_ = d_84_globalState_
            lhs17_ = d_84_globalState_
            lhs18_ = d_84_globalState_
            d_68_v1_ = rhs20_
            lhs16_.f14 = rhs21_
            d_146_v71_ = rhs22_
            lhs17_.f7 = rhs23_
            lhs18_.f20 = rhs24_
        if default__.fm0(d_131_v58_, d_68_v1_, d_127_v54_, d_73_v6_, d_84_globalState_):
            index18_ = default__.safeIndex(508, (d_80_v12_).length(0))
            (d_80_v12_)[index18_] = d_83_v15_
            d_149_v72_: _dafny.MultiSet
            d_149_v72_ = _dafny.MultiSet([746, d_68_v1_])
            d_150_v73_: _dafny.Seq
            d_150_v73_ = _dafny.SeqWithoutIsStrInference([(d_149_v72_).cardinality])
            d_151_v74_: _dafny.Seq
            d_151_v74_ = _dafny.SeqWithoutIsStrInference([d_150_v73_, d_150_v73_, d_150_v73_])
            index19_ = default__.safeIndex(508, (d_80_v12_).length(0))
            rhs25_ = ((d_151_v74_)[default__.safeIndex(d_68_v1_, len(d_151_v74_))]) < (default__.fm3(d_84_globalState_))
            rhs26_ = (d_68_v1_) - (d_68_v1_)
            rhs27_ = not((((d_149_v72_)[d_68_v1_] if (d_68_v1_) in (d_149_v72_) else d_68_v1_)) == (d_68_v1_))
            rhs28_ = (((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])) + (d_83_v15_)).set(default__.safeIndex(d_68_v1_, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])) + (d_83_v15_))), d_67_v0_)).set(default__.safeIndex(default__.safeDivisionInt(d_68_v1_, d_68_v1_), len(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])) + (d_83_v15_)).set(default__.safeIndex(d_68_v1_, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])) + (d_83_v15_))), d_67_v0_))), (d_67_v0_ if d_127_v54_ else d_67_v0_))
            lhs19_ = d_84_globalState_
            lhs20_ = d_80_v12_
            lhs21_ = default__.safeIndex(508, (d_80_v12_).length(0))
            d_127_v54_ = rhs25_
            lhs19_.f7 = rhs26_
            d_127_v54_ = rhs27_
            lhs20_[lhs21_] = rhs28_
            d_152_v75_: _dafny.Seq
            d_152_v75_ = _dafny.SeqWithoutIsStrInference([d_125_v52_, d_125_v52_])
            d_153_v76_: C0
            nw24_ = C0()
            nw24_.ctor__((d_152_v75_)[default__.safeIndex(d_68_v1_, len(d_152_v75_))])
            d_153_v76_ = nw24_
            d_154_v77_: C0
            nw25_ = C0()
            nw25_.ctor__((d_134_v61_).f22)
            d_154_v77_ = nw25_
            d_131_v58_ = default__.fm4((d_68_v1_ if d_126_v53_ else d_68_v1_), d_68_v1_, d_84_globalState_)
            d_155_v78_: _dafny.Set
            d_155_v78_ = _dafny.Set({d_67_v0_, d_67_v0_})
            d_156_v79_: _dafny.Map
            d_156_v79_ = _dafny.Map({D0_DC0(d_78_v10_): len(d_155_v78_)})
            d_157_v80_: _dafny.Map
            d_157_v80_ = _dafny.Map({d_127_v54_: d_149_v72_})
            d_158_v81_: bool
            d_159_v82_: bool
            out6_: bool
            out7_: bool
            out6_, out7_ = default__.m0(d_124_v51_, (d_153_v76_).f22, (((d_156_v79_)[default__.fm5(d_149_v72_, d_84_globalState_)] if (default__.fm5(d_149_v72_, d_84_globalState_)) in (d_156_v79_) else 315)) - ((((d_157_v80_)[d_73_v6_] if (d_73_v6_) in (d_157_v80_) else d_149_v72_)).cardinality), d_84_globalState_)
            d_158_v81_ = out6_
            d_159_v82_ = out7_
        elif True:
            if ((d_134_v61_).fm1(d_126_v53_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")) for d_160_i6_ in range(default__.abs(-683))]), d_84_globalState_)) and ((d_68_v1_) >= (d_68_v1_)):
                d_161_v83_: _dafny.Map
                d_161_v83_ = _dafny.Map({d_68_v1_: d_83_v15_})
                d_83_v15_ = ((d_161_v83_)[d_68_v1_] if (d_68_v1_) in (d_161_v83_) else d_83_v15_)
                d_162_v84_: C0
                nw26_ = C0()
                nw26_.ctor__((d_134_v61_).f22)
                d_162_v84_ = nw26_
                index20_ = default__.safeIndex(435, (d_76_v9_).length(0))
                (d_76_v9_)[index20_] = d_68_v1_
                d_163_v85_: D0
                d_163_v85_ = D0_DC1(d_68_v1_, d_127_v54_)
                d_164_v86_: _dafny.MultiSet
                d_164_v86_ = _dafny.MultiSet([d_163_v85_])
                d_165_v87_: _dafny.Map
                d_165_v87_ = _dafny.Map({d_164_v86_: d_127_v54_})
                d_166_v88_: _dafny.Array
                nw27_ = _dafny.Array(None, 2)
                nw27_[int(0)] = d_165_v87_
                nw27_[int(1)] = d_165_v87_
                d_166_v88_ = nw27_
                index21_ = default__.safeIndex(867, (d_166_v88_).length(0))
                (d_166_v88_)[index21_] = (d_165_v87_) | (d_165_v87_)
                d_167_v89_: _dafny.Seq
                d_167_v89_ = _dafny.SeqWithoutIsStrInference([d_73_v6_, d_127_v54_, not(not(True))])
                d_168_v90_: _dafny.Array
                d_168_v90_ = d_69_v2_
                d_169_v91_: _dafny.Array
                d_169_v91_ = (d_168_v90_)
                d_170_v92_: _dafny.Seq
                d_170_v92_ = _dafny.SeqWithoutIsStrInference([-376, 212])
                d_171_v93_: _dafny.Map
                d_171_v93_ = _dafny.Map({d_169_v91_: (0) - (len((d_170_v92_).set(default__.safeIndex(d_68_v1_, len(d_170_v92_)), d_68_v1_)))})
                d_172_v94_: D0
                d_172_v94_ = D0_DC0(d_78_v10_)
                d_173_v95_: _dafny.Map
                d_173_v95_ = _dafny.Map({d_172_v94_: d_134_v61_})
                d_174_v96_: _dafny.Map
                d_174_v96_ = _dafny.Map({len((d_171_v93_) | (d_171_v93_)): ((d_173_v95_)[d_172_v94_] if (d_172_v94_) in (d_173_v95_) else d_162_v84_)})
                d_175_v97_: D3
                d_175_v97_ = D3_DC6(d_162_v84_)
                d_176_v98_: _dafny.MultiSet
                d_176_v98_ = _dafny.MultiSet([d_81_v13_])
                index22_ = default__.safeIndex(435, (d_76_v9_).length(0))
                index23_ = default__.safeIndex(867, (d_166_v88_).length(0))
                rhs29_ = d_68_v1_
                rhs30_ = ((d_174_v96_)[d_68_v1_] if (d_68_v1_) in (d_174_v96_) else (d_175_v97_).cf8)
                rhs31_ = (d_176_v98_).cardinality
                rhs32_ = ((d_165_v87_) | (d_165_v87_)) | (d_165_v87_)
                rhs33_ = d_167_v89_
                lhs22_ = d_76_v9_
                lhs23_ = default__.safeIndex(435, (d_76_v9_).length(0))
                lhs24_ = d_84_globalState_
                lhs25_ = d_166_v88_
                lhs26_ = default__.safeIndex(867, (d_166_v88_).length(0))
                lhs22_[lhs23_] = rhs29_
                d_134_v61_ = rhs30_
                lhs24_.f7 = rhs31_
                lhs25_[lhs26_] = rhs32_
                d_167_v89_ = rhs33_
                (d_84_globalState_).f4 = (not(d_127_v54_)) or (d_127_v54_)
                d_83_v15_ = d_83_v15_
            elif True:
                d_177_v99_: C0
                nw28_ = C0()
                nw28_.ctor__((d_134_v61_).f22)
                d_177_v99_ = nw28_
                d_83_v15_ = d_83_v15_
                d_127_v54_ = d_127_v54_
                d_178_v100_: C0
                nw29_ = C0()
                nw29_.ctor__((d_134_v61_).f22)
                d_178_v100_ = nw29_
                d_179_v101_: _dafny.Seq
                d_179_v101_ = _dafny.SeqWithoutIsStrInference([d_178_v100_])
                d_180_v102_: _dafny.Array
                nw30_ = _dafny.Array(None, 21)
                nw30_[int(0)] = d_134_v61_
                nw30_[int(1)] = (d_179_v101_)[default__.safeIndex(d_68_v1_, len(d_179_v101_))]
                nw30_[int(2)] = d_134_v61_
                nw30_[int(3)] = d_178_v100_
                nw30_[int(4)] = d_178_v100_
                nw30_[int(5)] = d_177_v99_
                nw30_[int(6)] = d_134_v61_
                nw30_[int(7)] = d_134_v61_
                nw30_[int(8)] = d_178_v100_
                nw30_[int(9)] = d_178_v100_
                nw30_[int(10)] = d_134_v61_
                nw30_[int(11)] = d_178_v100_
                nw30_[int(12)] = d_177_v99_
                nw30_[int(13)] = d_178_v100_
                nw30_[int(14)] = d_177_v99_
                nw30_[int(15)] = d_134_v61_
                nw30_[int(16)] = d_134_v61_
                nw30_[int(17)] = d_177_v99_
                nw30_[int(18)] = d_178_v100_
                nw30_[int(19)] = d_178_v100_
                nw30_[int(20)] = d_178_v100_
                d_180_v102_ = nw30_
                d_181_v103_: _dafny.Set
                d_181_v103_ = _dafny.Set({d_180_v102_})
                rhs34_ = not(d_126_v53_)
                rhs35_ = d_73_v6_
                rhs36_ = len(((d_181_v103_) | (d_181_v103_)) | ((d_181_v103_) - (_dafny.Set({d_180_v102_}))))
                lhs27_ = d_84_globalState_
                lhs28_ = d_84_globalState_
                d_73_v6_ = rhs34_
                lhs27_.f6 = rhs35_
                lhs28_.f14 = rhs36_
            index24_ = default__.safeIndex(129, (d_76_v9_).length(0))
            (d_76_v9_)[index24_] = d_68_v1_
            d_182_v104_: _dafny.Map
            d_182_v104_ = _dafny.Map({d_127_v54_: d_76_v9_})
            index25_ = default__.safeIndex(129, (d_76_v9_).length(0))
            (d_76_v9_)[index25_] = len(d_182_v104_)
            d_183_v105_: _dafny.Map
            d_183_v105_ = _dafny.Map({len(d_83_v15_): d_127_v54_})
            d_184_v106_: _dafny.Map
            d_184_v106_ = _dafny.Map({-955: not(((d_183_v105_)[d_68_v1_] if (d_68_v1_) in (d_183_v105_) else d_73_v6_))})
            d_184_v106_ = d_184_v106_
            (d_84_globalState_).f6 = True
            d_185_v107_: _dafny.Seq
            d_185_v107_ = _dafny.SeqWithoutIsStrInference([d_127_v54_, d_126_v53_, d_127_v54_])
            d_186_v108_: _dafny.Seq
            d_186_v108_ = _dafny.SeqWithoutIsStrInference([(d_76_v9_)[default__.safeIndex(129, (d_76_v9_).length(0))]])
            d_187_v109_: _dafny.Map
            d_187_v109_ = _dafny.Map({d_185_v107_: len(default__.fm6(d_186_v108_, d_126_v53_, d_73_v6_, d_84_globalState_))})
            index26_ = default__.safeIndex(129, (d_76_v9_).length(0))
            rhs37_ = (((d_76_v9_)[default__.safeIndex(129, (d_76_v9_).length(0))]) - ((d_76_v9_)[default__.safeIndex(129, (d_76_v9_).length(0))]) if d_127_v54_ else (d_76_v9_)[default__.safeIndex(129, (d_76_v9_).length(0))])
            rhs38_ = not(d_73_v6_)
            rhs39_ = (0) - (((d_187_v109_)[(d_185_v107_) + (d_185_v107_)] if ((d_185_v107_) + (d_185_v107_)) in (d_187_v109_) else (d_76_v9_)[default__.safeIndex(129, (d_76_v9_).length(0))]))
            rhs40_ = (d_76_v9_)[default__.safeIndex(129, (d_76_v9_).length(0))]
            lhs29_ = d_76_v9_
            lhs30_ = default__.safeIndex(129, (d_76_v9_).length(0))
            lhs31_ = d_84_globalState_
            lhs32_ = d_84_globalState_
            lhs33_ = d_84_globalState_
            lhs29_[lhs30_] = rhs37_
            lhs31_.f4 = rhs38_
            lhs32_.f14 = rhs39_
            lhs33_.f14 = rhs40_
        (d_84_globalState_).f14 = (d_68_v1_) * (d_68_v1_)
        d_188_i7_: int
        d_188_i7_ = 0
        with _dafny.label("1"):
            while (d_83_v15_) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pt"))) + (_dafny.SeqWithoutIsStrInference([d_67_v0_ for d_204_i8_ in range(default__.abs(834))]))):
                with _dafny.c_label("1"):
                    if (d_188_i7_) >= (100):
                        raise _dafny.Break("1")
                    d_188_i7_ = (d_188_i7_) + (1)
                    d_189_v110_: D3
                    d_189_v110_ = D3_DC6(d_134_v61_)
                    d_189_v110_ = D3_DC6((d_189_v110_).cf8)
                    d_190_v111_: C0
                    nw31_ = C0()
                    nw31_.ctor__(d_125_v52_)
                    d_190_v111_ = nw31_
                    d_191_v112_: _dafny.Seq
                    d_191_v112_ = ((d_100_v27_)[d_73_v6_] if (d_73_v6_) in (d_100_v27_) else d_83_v15_)
                    d_83_v15_ = (((d_191_v112_)) + (d_83_v15_)) + (d_83_v15_)
                    d_192_v113_: _dafny.Seq
                    d_192_v113_ = _dafny.SeqWithoutIsStrInference([d_83_v15_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_193_i10_ in range(default__.abs(285))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikudn"))])
                    if default__.fm0(d_131_v58_, ((0) - ((0) - (d_68_v1_))) + (len(_dafny.SeqWithoutIsStrInference([90 for d_194_i9_ in range(default__.abs(6))]))), d_126_v53_, (d_190_v111_).fm1(d_73_v6_, d_192_v113_, d_84_globalState_), d_84_globalState_):
                        (d_84_globalState_).f14 = (d_68_v1_) - (default__.safeDivisionInt(915, d_68_v1_))
                        d_83_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sv"))
                        d_83_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbjlpyup"))
                        d_195_v114_: C0
                        nw32_ = C0()
                        nw32_.ctor__(((d_134_v61_).f22) | (((d_134_v61_).f22).set(d_73_v6_, d_124_v51_)))
                        d_195_v114_ = nw32_
                        d_196_v115_: _dafny.Seq
                        d_196_v115_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_190_v111_: d_190_v111_})), d_68_v1_])
                        d_197_v117_: _dafny.Map
                        def iife6_():
                            coll2_ = _dafny.Set()
                            compr_2_: int
                            for compr_2_ in (d_196_v115_).Elements:
                                d_198_v116_: int = compr_2_
                                if (d_198_v116_) in (d_196_v115_):
                                    coll2_ = coll2_.union(_dafny.Set([(d_198_v116_) * (len(_dafny.SeqWithoutIsStrInference([(0) - (-117) for d_199_i11_ in range(default__.abs(783))])))]))
                            return _dafny.Set(coll2_)
                        d_197_v117_ = _dafny.Map({(d_83_v15_) + (default__.fm7(iife6_()
                        , d_84_globalState_)): _dafny.Map({d_68_v1_: d_68_v1_})})
                        d_200_v118_: D3
                        d_200_v118_ = D3_DC8(D3_DC7())
                        rhs41_ = d_73_v6_
                        rhs42_ = d_197_v117_
                        rhs43_ = d_76_v9_
                        rhs44_ = d_200_v118_
                        lhs34_ = d_84_globalState_
                        lhs34_.f4 = rhs41_
                        d_197_v117_ = rhs42_
                        d_76_v9_ = rhs43_
                        d_200_v118_ = rhs44_
                    elif True:
                        d_83_v15_ = d_83_v15_
                        d_201_v119_: _dafny.Array
                        nw33_ = _dafny.Array(D3.default()(), 3)
                        d_201_v119_ = nw33_
                        rhs45_ = (D3_DC6(d_190_v111_)).cf8
                        rhs46_ = d_201_v119_
                        rhs47_ = d_83_v15_
                        d_134_v61_ = rhs45_
                        d_201_v119_ = rhs46_
                        d_83_v15_ = rhs47_
                        d_202_v120_: C0
                        nw34_ = C0()
                        nw34_.ctor__(d_125_v52_)
                        d_202_v120_ = nw34_
                        d_203_v121_: C0
                        nw35_ = C0()
                        nw35_.ctor__((d_134_v61_).f22)
                        d_203_v121_ = nw35_
                        d_75_v8_ = (d_75_v8_).set(d_127_v54_, d_68_v1_)
                    pass
            pass
        def iife15_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(372, 739):
                d_220_v122_: int = compr_3_
                if ((372) <= (d_220_v122_)) and ((d_220_v122_) < (739)):
                    coll3_[default__.safeDivisionInt(d_220_v122_, d_68_v1_)] = d_126_v53_
            return _dafny.Map(coll3_)
        hi1_ = 492
        for d_205_i12_ in range((len(iife15_()
        ) if d_73_v6_ else d_68_v1_), hi1_):
            d_206_v123_: D3
            d_206_v123_ = D3_DC6(d_134_v61_)
            d_207_v124_: D3
            d_207_v124_ = D3_DC8(d_206_v123_)
            pat_let_tv2_ = d_206_v123_
            pat_let_tv3_ = d_206_v123_
            d_208_v125_: _dafny.Array
            nw36_ = _dafny.Array(None, 6)
            nw36_[int(0)] = d_207_v124_
            nw36_[int(1)] = d_207_v124_
            nw36_[int(2)] = d_207_v124_
            def iife7_(_pat_let2_0):
                def iife8_(d_209_dt__update__tmp_h1_):
                    def iife9_(_pat_let3_0):
                        def iife10_(d_210_dt__update_hcf9_h0_):
                            return D3_DC8(d_210_dt__update_hcf9_h0_)
                        return iife10_(_pat_let3_0)
                    return iife9_(pat_let_tv2_)
                return iife8_(_pat_let2_0)
            nw36_[int(3)] = iife7_(d_207_v124_)
            def iife11_(_pat_let4_0):
                def iife12_(d_211_dt__update__tmp_h2_):
                    def iife13_(_pat_let5_0):
                        def iife14_(d_212_dt__update_hcf9_h1_):
                            return D3_DC8(d_212_dt__update_hcf9_h1_)
                        return iife14_(_pat_let5_0)
                    return iife13_(pat_let_tv3_)
                return iife12_(_pat_let4_0)
            nw36_[int(4)] = iife11_(d_207_v124_)
            nw36_[int(5)] = d_207_v124_
            d_208_v125_ = nw36_
            d_213_v126_: _dafny.Map
            d_213_v126_ = _dafny.Map({d_208_v125_: d_76_v9_})
            index27_ = default__.safeIndex(778, (d_124_v51_).length(0))
            (d_124_v51_)[index27_] = d_127_v54_
            d_214_v127_: _dafny.Map
            d_214_v127_ = _dafny.Map({d_67_v0_: d_134_v61_})
            index28_ = default__.safeIndex(778, (d_124_v51_).length(0))
            rhs48_ = ((d_213_v126_).set(d_208_v125_, d_69_v2_)) | ((d_213_v126_) | (d_213_v126_))
            rhs49_ = d_124_v51_
            rhs50_ = ((d_214_v127_)[_dafny.CodePoint('t')] if (_dafny.CodePoint('t')) in (d_214_v127_) else d_134_v61_)
            rhs51_ = (d_127_v54_) or (d_73_v6_)
            lhs35_ = d_124_v51_
            lhs36_ = default__.safeIndex(778, (d_124_v51_).length(0))
            d_213_v126_ = rhs48_
            d_124_v51_ = rhs49_
            d_134_v61_ = rhs50_
            lhs35_[lhs36_] = rhs51_
            d_215_v128_: C0
            nw37_ = C0()
            nw37_.ctor__((d_134_v61_).f22)
            d_215_v128_ = nw37_
            index29_ = default__.safeIndex(739, (d_69_v2_).length(0))
            (d_69_v2_)[index29_] = d_68_v1_
            d_216_v129_: _dafny.MultiSet
            d_216_v129_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_67_v0_ for d_217_i13_ in range(default__.abs(198))])), d_205_i12_])
            d_218_v130_: _dafny.Map
            d_218_v130_ = _dafny.Map({(d_216_v129_) - (d_216_v129_): _dafny.SeqWithoutIsStrInference([653])})
            index30_ = default__.safeIndex(739, (d_69_v2_).length(0))
            index31_ = default__.safeIndex(778, (d_124_v51_).length(0))
            rhs52_ = (len(_dafny.SeqWithoutIsStrInference([d_68_v1_, d_205_i12_, d_68_v1_, d_205_i12_, 658]))) - (default__.safeModuloInt(d_68_v1_, d_68_v1_))
            rhs53_ = d_218_v130_
            rhs54_ = (d_134_v61_ if not(d_73_v6_) else d_215_v128_)
            rhs55_ = not((d_124_v51_)[default__.safeIndex(778, (d_124_v51_).length(0))])
            lhs37_ = d_69_v2_
            lhs38_ = default__.safeIndex(739, (d_69_v2_).length(0))
            lhs39_ = d_124_v51_
            lhs40_ = default__.safeIndex(778, (d_124_v51_).length(0))
            lhs37_[lhs38_] = rhs52_
            d_218_v130_ = rhs53_
            d_215_v128_ = rhs54_
            lhs39_[lhs40_] = rhs55_
            d_219_v131_: _dafny.Seq
            d_219_v131_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uiiomhn"))])
            index32_ = default__.safeIndex(778, (d_124_v51_).length(0))
            rhs56_ = (default__.fm0(d_131_v58_, d_68_v1_, (d_215_v128_).fm1(not(d_126_v53_), _dafny.SeqWithoutIsStrInference([default__.fm7(d_71_v4_, d_84_globalState_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odeavklj"))]), d_84_globalState_), (d_215_v128_).fm1(True, d_219_v131_, d_84_globalState_), d_84_globalState_) if (d_124_v51_)[default__.safeIndex(778, (d_124_v51_).length(0))] else d_126_v53_)
            rhs57_ = (d_69_v2_)[default__.safeIndex(739, (d_69_v2_).length(0))]
            lhs41_ = d_124_v51_
            lhs42_ = default__.safeIndex(778, (d_124_v51_).length(0))
            lhs43_ = d_84_globalState_
            lhs41_[lhs42_] = rhs56_
            lhs43_.f20 = rhs57_
        _dafny.print(_dafny.string_of(d_67_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_68_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v2_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v2_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_70_v3_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_v4_) == (_dafny.Set({1, 557}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_72_v5_) == (_dafny.MultiSet([_dafny.Set({1, 557})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_73_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_74_v7_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_v8_) == (_dafny.Map({True: 310249}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v10_) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_79_v11_) == (_dafny.Set({_dafny.Map({True: 557})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_80_v12_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_81_v13_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v14_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_83_v15_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f0) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x')]): _dafny.CodePoint('w')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f1)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f2) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f3) == (_dafny.Set({_dafny.Map({True: 557})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_84_globalState_).f5)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f8) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f12) == (_dafny.Set({1, 557}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f13)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f16) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f17)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f17)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_84_globalState_).f18) == (_dafny.Map({False: 2, True: 557}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_globalState_.f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_84_globalState_).f21).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v27_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfucpd"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v51_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_125_v52_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_126_v53_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_127_v54_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_v56_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False}), _dafny.Map({True: False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v57_) == (_dafny.Map({_dafny.Map({_dafny.Map({True: False}): True}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v58_) == (_dafny.Map({387: -10}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v59_) == (_dafny.Map({_dafny.Map({False: False}): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_133_v60_).cf3)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_134_v61_).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_188_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
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
        return lambda: D1_DC3(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC5(D2, NamedTuple('DC5', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC7()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC7(D3, NamedTuple('DC7', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC9(D4, NamedTuple('DC9', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({self.cf10.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)

class D5_DC10(D5, NamedTuple('DC10', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D6_DC11)

class D6_DC11(D6, NamedTuple('DC11', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC11({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC11) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f4: bool = False
        self.f6: bool = False
        self.f7: int = int(0)
        self.f14: int = int(0)
        self.f20: int = int(0)
        self._f0: _dafny.Map = _dafny.Map({})
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f2: _dafny.Map = _dafny.Map({})
        self._f3: _dafny.Set = _dafny.Set({})
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f8: _dafny.MultiSet = _dafny.MultiSet({})
        self._f9: int = int(0)
        self._f10: int = int(0)
        self._f11: bool = False
        self._f12: _dafny.Set = _dafny.Set({})
        self._f13: _dafny.Array = _dafny.Array(None, 0)
        self._f15: bool = False
        self._f16: _dafny.Set = _dafny.Set({})
        self._f17: _dafny.Array = _dafny.Array(None, 0)
        self._f18: _dafny.Map = _dafny.Map({})
        self._f19: int = int(0)
        self._f21: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self).f20 = f20
        (self)._f21 = f21

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
    def f5(self):
        return self._f5
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
    def f15(self):
        return self._f15
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
    def f19(self):
        return self._f19
    @property
    def f21(self):
        return self._f21

class C0:
    def  __init__(self):
        self._f22: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f22):
        (self)._f22 = f22

    def fm1(self, p0, p1, globalState):
        source2_ = D0_DC1(len(_dafny.Set({not(False), True, True})), not(False))
        if source2_.is_DC1:
            d_221___mcc_h0_ = source2_.cf1
            d_222___mcc_h1_ = source2_.cf2
            d_223_cf2_ = d_222___mcc_h1_
            d_224_cf1_ = d_221___mcc_h0_
            return (D0_DC1((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpsgaxsnn")))), d_223_cf2_)).cf2
        elif True:
            d_225___mcc_h2_ = source2_.cf0
            d_226_cf0_ = d_225___mcc_h2_
            return not (not(False)) or (True)

    @property
    def f22(self):
        return self._f22
