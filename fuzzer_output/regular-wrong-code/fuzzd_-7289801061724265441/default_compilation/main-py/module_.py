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
        return False

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return (0) - (-300)

    @staticmethod
    def fm2(p0, globalState):
        return D0_DC0(_dafny.SeqWithoutIsStrInference([-446 for d_0_i0_ in range(default__.abs(905))]))

    @staticmethod
    def fm3(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1_i0_ in range(default__.abs(-787))])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fivufydg")))})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Set({(0) - (-437), 967})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pevdpvb")))}) for d_2_i1_ in range(default__.abs(390))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True)])): 1}) for d_3_i2_ in range(default__.abs(523))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({-719: (_dafny.MultiSet([True, True, False, False, False])).cardinality})])))

    @staticmethod
    def fm5(globalState):
        return _dafny.SeqWithoutIsStrInference([not(not(True))])

    @staticmethod
    def fm6(globalState):
        source0_ = D2_DC8(_dafny.SeqWithoutIsStrInference([False, False, False]))
        if source0_.is_DC9:
            d_4___mcc_h0_ = source0_.cf14
            d_5___mcc_h1_ = source0_.cf15
            d_6_cf15_ = d_5___mcc_h1_
            d_7_cf14_ = d_4___mcc_h0_
            return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "byjjmnp")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkpsd"))])
        elif True:
            d_8___mcc_h2_ = source0_.cf13
            d_9_cf13_ = d_8___mcc_h2_
            return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihjrffkga")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xiypy")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxg"))])

    @staticmethod
    def fm7(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruoqh"))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_10_i0_ in range(default__.abs(395))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irmgafus"))))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([len(_dafny.Map({False: _dafny.CodePoint('n')}))])) | (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')]))]))) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_11_i1_ in range(default__.abs(508))])) for d_12_i0_ in range(default__.abs(879))]))])).intersection(_dafny.MultiSet([768, 353])))

    @staticmethod
    def m0(globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_13_v0_: int
        d_13_v0_ = 234
        d_14_v1_: _dafny.Seq
        d_14_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmmfs"))
        d_15_v2_: _dafny.MultiSet
        d_15_v2_ = _dafny.MultiSet([len(d_14_v1_)])
        d_16_v3_: bool
        d_16_v3_ = False
        d_17_v4_: _dafny.Map
        d_17_v4_ = _dafny.Map({d_16_v3_: d_13_v0_})
        d_18_v5_: _dafny.MultiSet
        d_18_v5_ = _dafny.MultiSet([d_16_v3_])
        d_19_v6_: _dafny.Seq
        d_19_v6_ = _dafny.SeqWithoutIsStrInference([d_16_v3_, d_16_v3_])
        d_20_v7_: _dafny.Set
        d_20_v7_ = _dafny.Set({d_13_v0_, (0) - (d_13_v0_), len(d_19_v6_), 417})
        d_21_v8_: _dafny.Array
        nw0_ = _dafny.Array(None, 16)
        nw0_[int(0)] = d_13_v0_
        nw0_[int(1)] = d_13_v0_
        nw0_[int(2)] = d_13_v0_
        nw0_[int(3)] = 43
        nw0_[int(4)] = d_13_v0_
        nw0_[int(5)] = d_13_v0_
        nw0_[int(6)] = d_13_v0_
        nw0_[int(7)] = ((d_15_v2_)[d_13_v0_] if (d_13_v0_) in (d_15_v2_) else ((d_17_v4_)[default__.fm0(d_13_v0_, d_16_v3_, d_16_v3_, d_14_v1_, globalState)] if (default__.fm0(d_13_v0_, d_16_v3_, d_16_v3_, d_14_v1_, globalState)) in (d_17_v4_) else d_13_v0_))
        nw0_[int(8)] = 554
        nw0_[int(9)] = ((d_18_v5_).cardinality) + (len(d_14_v1_))
        nw0_[int(10)] = ((d_18_v5_)[d_16_v3_] if (d_16_v3_) in (d_18_v5_) else d_13_v0_)
        nw0_[int(11)] = d_13_v0_
        nw0_[int(12)] = d_13_v0_
        nw0_[int(13)] = d_13_v0_
        nw0_[int(14)] = (380) * (default__.fm1((0) - (len(d_19_v6_)), d_16_v3_, d_20_v7_, globalState))
        nw0_[int(15)] = d_13_v0_
        d_21_v8_ = nw0_
        d_22_v9_: _dafny.Map
        d_22_v9_ = _dafny.Map({d_16_v3_: d_21_v8_})
        index0_ = default__.safeIndex(381, (d_21_v8_).length(0))
        (d_21_v8_)[index0_] = len(d_22_v9_)
        index1_ = default__.safeIndex(381, (d_21_v8_).length(0))
        (d_21_v8_)[index1_] = d_13_v0_
        d_23_v10_: _dafny.Array
        nw1_ = _dafny.Array(False, 19)
        d_23_v10_ = nw1_
        index2_ = default__.safeIndex(692, (d_23_v10_).length(0))
        (d_23_v10_)[index2_] = d_16_v3_
        d_24_v11_: D3
        d_24_v11_ = D3_DC10(d_20_v7_)
        d_25_v12_: _dafny.Seq
        d_25_v12_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_13_v0_, d_16_v3_, (d_24_v11_).cf16, globalState), d_13_v0_])
        d_26_v13_: _dafny.Map
        d_26_v13_ = _dafny.Map({985: d_17_v4_})
        d_27_v14_: str
        d_27_v14_ = _dafny.CodePoint('p')
        index3_ = default__.safeIndex(692, (d_23_v10_).length(0))
        rhs0_ = d_19_v6_
        rhs1_ = d_13_v0_
        rhs2_ = d_16_v3_
        rhs3_ = ((_dafny.MultiSet(d_25_v12_)) | (default__.fm8((d_25_v12_)[default__.safeIndex(len(((d_26_v13_)[(d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))]] if ((d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))]) in (d_26_v13_) else _dafny.Map({d_16_v3_: d_13_v0_}))), len(d_25_v12_))], d_16_v3_, _dafny.Map({d_16_v3_: d_27_v14_}), len(d_19_v6_), globalState))).intersection(_dafny.MultiSet([d_13_v0_]))
        rhs4_ = d_16_v3_
        lhs0_ = globalState
        lhs1_ = d_23_v10_
        lhs2_ = default__.safeIndex(692, (d_23_v10_).length(0))
        d_19_v6_ = rhs0_
        lhs0_.f20 = rhs1_
        d_16_v3_ = rhs2_
        d_15_v2_ = rhs3_
        lhs1_[lhs2_] = rhs4_
        hi0_ = (d_13_v0_) + ((d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))])
        for d_28_i0_ in range(len(_dafny.SeqWithoutIsStrInference([d_13_v0_, d_13_v0_, (d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))]])), hi0_):
            index4_ = default__.safeIndex(692, (d_23_v10_).length(0))
            (d_23_v10_)[index4_] = (((_dafny.MultiSet([d_13_v0_])).cardinality) < (d_28_i0_)) and (((d_23_v10_)[default__.safeIndex(692, (d_23_v10_).length(0))]) == (d_16_v3_))
            d_29_v15_: _dafny.Map
            d_29_v15_ = _dafny.Map({(d_23_v10_)[default__.safeIndex(692, (d_23_v10_).length(0))]: d_23_v10_})
            d_29_v15_ = (d_29_v15_).set(((d_23_v10_)[default__.safeIndex(692, (d_23_v10_).length(0))]) or (d_16_v3_), d_23_v10_)
            d_30_v16_: _dafny.Map
            d_30_v16_ = _dafny.Map({d_21_v8_: (d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))]})
            d_31_v17_: _dafny.Map
            d_31_v17_ = _dafny.Map({True: d_16_v3_})
            d_32_v18_: _dafny.Map
            d_32_v18_ = _dafny.Map({d_31_v17_: d_16_v3_})
            d_33_v19_: _dafny.Seq
            d_33_v19_ = _dafny.SeqWithoutIsStrInference([d_32_v18_, d_32_v18_])
            d_34_v20_: _dafny.Array
            nw2_ = _dafny.Array(None, 13)
            nw2_[int(0)] = (d_30_v16_) | (d_30_v16_)
            nw2_[int(1)] = d_30_v16_
            nw2_[int(2)] = d_30_v16_
            nw2_[int(3)] = d_30_v16_
            nw2_[int(4)] = d_30_v16_
            nw2_[int(5)] = d_30_v16_
            nw2_[int(6)] = _dafny.Map({d_21_v8_: 96})
            nw2_[int(7)] = ((d_30_v16_).set(d_21_v8_, (d_15_v2_).cardinality)).set(((d_22_v9_)[d_16_v3_] if (d_16_v3_) in (d_22_v9_) else d_21_v8_), (d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))])
            nw2_[int(8)] = d_30_v16_
            nw2_[int(9)] = _dafny.Map({d_21_v8_: (d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))]})
            nw2_[int(10)] = d_30_v16_
            nw2_[int(11)] = ((d_30_v16_).set(d_21_v8_, len((d_33_v19_)[default__.safeIndex((d_25_v12_)[default__.safeIndex(len(d_25_v12_), len(d_25_v12_))], len(d_33_v19_))]))) | (d_30_v16_)
            nw2_[int(12)] = d_30_v16_
            d_34_v20_ = nw2_
            d_35_v21_: C0
            nw3_ = C0()
            nw3_.ctor__((d_23_v10_)[default__.safeIndex(692, (d_23_v10_).length(0))])
            d_35_v21_ = nw3_
            d_36_v22_: _dafny.MultiSet
            d_36_v22_ = _dafny.MultiSet([d_35_v21_])
            index5_ = default__.safeIndex(195, (d_34_v20_).length(0))
            (d_34_v20_)[index5_] = (_dafny.Map({d_21_v8_: ((d_36_v22_)[d_35_v21_] if (d_35_v21_) in (d_36_v22_) else d_13_v0_)})) | (_dafny.Map({d_21_v8_: len(d_31_v17_)}))
            index6_ = default__.safeIndex(195, (d_34_v20_).length(0))
            rhs5_ = (d_23_v10_)[default__.safeIndex(692, (d_23_v10_).length(0))]
            rhs6_ = ((d_14_v1_) + (d_14_v1_)) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))).set(default__.safeIndex(len(d_25_v12_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))), _dafny.CodePoint('e')))
            rhs7_ = default__.fm1(((0) - ((d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))])) + (d_28_i0_), (d_19_v6_) != (_dafny.SeqWithoutIsStrInference([d_16_v3_])), d_20_v7_, globalState)
            rhs8_ = d_30_v16_
            lhs3_ = globalState
            lhs4_ = d_34_v20_
            lhs5_ = default__.safeIndex(195, (d_34_v20_).length(0))
            d_16_v3_ = rhs5_
            d_16_v3_ = rhs6_
            lhs3_.f20 = rhs7_
            lhs4_[lhs5_] = rhs8_
            d_37_v23_: D0
            d_37_v23_ = D0_DC0(d_25_v12_)
            d_38_v24_: _dafny.Array
            nw4_ = _dafny.Array(None, 3)
            nw4_[int(0)] = D0_DC0(d_25_v12_)
            nw4_[int(1)] = d_37_v23_
            nw4_[int(2)] = default__.fm2((d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))], globalState)
            d_38_v24_ = nw4_
            d_39_v25_: _dafny.Map
            d_39_v25_ = _dafny.Map({d_35_v21_: d_38_v24_})
            d_39_v25_ = (d_39_v25_).set(d_35_v21_, d_38_v24_)
        index7_ = default__.safeIndex(692, (d_23_v10_).length(0))
        (d_23_v10_)[index7_] = (d_23_v10_)[default__.safeIndex(692, (d_23_v10_).length(0))]
        d_40_v26_: _dafny.Map
        d_40_v26_ = _dafny.Map({(d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))]: (d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))]})
        d_41_v27_: _dafny.Map
        d_41_v27_ = _dafny.Map({(d_13_v0_ if not(d_16_v3_) else d_13_v0_): (d_16_v3_ if default__.fm0(len(d_40_v26_), d_16_v3_, d_16_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whejvaw")), globalState) else True)})
        d_16_v3_ = not(((d_41_v27_)[d_13_v0_] if (d_13_v0_) in (d_41_v27_) else (d_23_v10_)[default__.safeIndex(692, (d_23_v10_).length(0))]))
        d_42_v28_: _dafny.Map
        d_42_v28_ = _dafny.Map({d_23_v10_: default__.fm0((d_21_v8_)[default__.safeIndex(381, (d_21_v8_).length(0))], d_16_v3_, False, d_14_v1_, globalState)})
        d_16_v3_ = ((d_42_v28_)[d_23_v10_] if (d_23_v10_) in (d_42_v28_) else (d_16_v3_) == (d_16_v3_))
        d_43_v29_: C0
        nw5_ = C0()
        nw5_.ctor__(d_16_v3_)
        d_43_v29_ = nw5_
        d_44_v30_: _dafny.MultiSet
        d_44_v30_ = _dafny.MultiSet([d_43_v29_])
        d_45_v31_: _dafny.Map
        d_45_v31_ = _dafny.Map({(d_44_v30_).cardinality: d_21_v8_})
        r0 = d_45_v31_
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_46_v0_: _dafny.Seq
        d_46_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljlofb"))
        d_47_v1_: _dafny.Array
        nw6_ = _dafny.Array(_dafny.Array(None, 0), 5)
        d_47_v1_ = nw6_
        d_48_v2_: int
        d_48_v2_ = 135
        d_49_v3_: _dafny.Seq
        d_49_v3_ = _dafny.SeqWithoutIsStrInference([d_48_v2_])
        d_50_v4_: _dafny.Array
        nw7_ = _dafny.Array(int(0), 27)
        d_50_v4_ = nw7_
        d_51_v5_: bool
        d_51_v5_ = False
        d_52_v6_: _dafny.Map
        d_52_v6_ = _dafny.Map({d_51_v5_: 584})
        d_53_v7_: str
        d_53_v7_ = _dafny.CodePoint('t')
        d_54_v8_: _dafny.Seq
        d_54_v8_ = _dafny.SeqWithoutIsStrInference([d_51_v5_])
        d_55_v9_: _dafny.Map
        d_55_v9_ = _dafny.Map({(d_46_v0_)[default__.safeIndex(d_48_v2_, len(d_46_v0_))]: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apnhjorfd")))})
        d_56_v10_: _dafny.Seq
        d_56_v10_ = _dafny.SeqWithoutIsStrInference([d_55_v9_, _dafny.Map({d_53_v7_: d_48_v2_})])
        d_57_v11_: _dafny.Map
        d_57_v11_ = _dafny.Map({d_48_v2_: (d_56_v10_)[default__.safeIndex(d_48_v2_, len(d_56_v10_))]})
        d_58_v12_: _dafny.Seq
        d_58_v12_ = _dafny.SeqWithoutIsStrInference([d_46_v0_])
        d_59_globalState_: GlobalState
        nw8_ = GlobalState()
        nw8_.ctor__(d_46_v0_, d_46_v0_, False, d_47_v1_, d_49_v3_, True, 604, False, d_50_v4_, (d_52_v6_) | (d_52_v6_), 503, (d_46_v0_).set(default__.safeIndex(d_48_v2_, len(d_46_v0_)), d_53_v7_), (d_54_v8_) + ((d_54_v8_).set(default__.safeIndex(9, len(d_54_v8_)), d_51_v5_)), 996, d_49_v3_, d_50_v4_, False, 938, (d_57_v11_) | (d_57_v11_), d_58_v12_, 465, 48)
        d_59_globalState_ = nw8_
        d_60_v13_: _dafny.Array
        def lambda0_(d_61_i0_):
            return False

        init0_ = lambda0_
        nw9_ = _dafny.Array(None, 8)
        for i0_0_ in range(nw9_.length(0)):
            nw9_[i0_0_] = init0_(i0_0_)
        d_60_v13_ = nw9_
        d_62_v14_: _dafny.Seq
        d_62_v14_ = _dafny.SeqWithoutIsStrInference([d_60_v13_])
        d_60_v13_ = (d_62_v14_)[default__.safeIndex(d_48_v2_, len(d_62_v14_))]
        hi1_ = d_48_v2_
        for d_63_i1_ in range(default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference([d_48_v2_, d_48_v2_, d_48_v2_, d_48_v2_, len(_dafny.Map({len(d_46_v0_): d_48_v2_}))])).set(default__.safeIndex(d_48_v2_, len(_dafny.SeqWithoutIsStrInference([d_48_v2_, d_48_v2_, d_48_v2_, d_48_v2_, len(_dafny.Map({len(d_46_v0_): d_48_v2_}))]))), d_48_v2_)), d_48_v2_), hi1_):
            d_64_v15_: _dafny.Map
            d_64_v15_ = _dafny.Map({d_48_v2_: 174})
            rhs9_ = d_64_v15_
            rhs10_ = (d_48_v2_) * (d_63_i1_)
            lhs6_ = d_59_globalState_
            d_64_v15_ = rhs9_
            lhs6_.f20 = rhs10_
            d_65_v16_: _dafny.Seq
            d_65_v16_ = _dafny.SeqWithoutIsStrInference([d_50_v4_])
            d_66_v17_: _dafny.Map
            d_66_v17_ = _dafny.Map({d_48_v2_: (d_65_v16_)[default__.safeIndex(d_48_v2_, len(d_65_v16_))]})
            d_67_v18_: _dafny.Seq
            d_67_v18_ = _dafny.SeqWithoutIsStrInference([d_50_v4_, d_50_v4_, ((d_66_v17_)[d_48_v2_] if (d_48_v2_) in (d_66_v17_) else d_50_v4_)])
            d_67_v18_ = (d_65_v16_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpxkaki"))), len(d_65_v16_)), d_50_v4_)
            d_54_v8_ = d_54_v8_
            d_51_v5_ = d_51_v5_
        d_52_v6_ = (d_52_v6_).set(d_51_v5_, d_48_v2_)
        d_68_v19_: _dafny.Map
        out0_: _dafny.Map
        out0_ = default__.m0(d_59_globalState_)
        d_68_v19_ = out0_
        if (False) or (False):
            d_69_v20_: _dafny.Seq
            d_69_v20_ = _dafny.SeqWithoutIsStrInference([(d_49_v3_) + (d_49_v3_), d_49_v3_, d_49_v3_])
            d_69_v20_ = d_69_v20_
            index8_ = default__.safeIndex(324, (d_60_v13_).length(0))
            (d_60_v13_)[index8_] = d_51_v5_
            index9_ = default__.safeIndex(324, (d_60_v13_).length(0))
            (d_60_v13_)[index9_] = d_51_v5_
            d_70_v21_: _dafny.MultiSet
            d_70_v21_ = _dafny.MultiSet([(d_60_v13_)[default__.safeIndex(324, (d_60_v13_).length(0))]])
            index10_ = default__.safeIndex(324, (d_60_v13_).length(0))
            (d_60_v13_)[index10_] = not(not((d_51_v5_) not in (d_70_v21_)))
            def iife0_():
                coll0_ = _dafny.Map()
                compr_0_: int
                for compr_0_ in _dafny.IntegerRange(938, 402):
                    d_71_v22_: int = compr_0_
                    if ((938) <= (d_71_v22_)) and ((d_71_v22_) < (402)):
                        coll0_[(d_71_v22_) * (d_48_v2_)] = (d_60_v13_)[default__.safeIndex(324, (d_60_v13_).length(0))]
                return _dafny.Map(coll0_)
            d_46_v0_ = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyxqgdxc"))) + (d_46_v0_)).set(default__.safeIndex(len(iife0_()
            ), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyxqgdxc"))) + (d_46_v0_))), d_53_v7_)) + (d_46_v0_)
            d_72_v23_: _dafny.Array
            nw10_ = _dafny.Array(None, 17)
            nw10_[int(0)] = default__.fm0(d_48_v2_, (d_60_v13_)[default__.safeIndex(324, (d_60_v13_).length(0))], d_51_v5_, d_46_v0_, d_59_globalState_)
            nw10_[int(1)] = d_51_v5_
            nw10_[int(2)] = d_51_v5_
            nw10_[int(3)] = not(d_51_v5_)
            nw10_[int(4)] = d_51_v5_
            nw10_[int(5)] = d_51_v5_
            nw10_[int(6)] = default__.fm0(-251, False, d_51_v5_, d_46_v0_, d_59_globalState_)
            nw10_[int(7)] = d_51_v5_
            nw10_[int(8)] = (d_60_v13_)[default__.safeIndex(324, (d_60_v13_).length(0))]
            nw10_[int(9)] = d_51_v5_
            nw10_[int(10)] = (d_60_v13_)[default__.safeIndex(324, (d_60_v13_).length(0))]
            nw10_[int(11)] = default__.fm0(len(d_46_v0_), (d_60_v13_)[default__.safeIndex(324, (d_60_v13_).length(0))], d_51_v5_, d_46_v0_, d_59_globalState_)
            nw10_[int(12)] = d_51_v5_
            nw10_[int(13)] = (d_60_v13_)[default__.safeIndex(324, (d_60_v13_).length(0))]
            nw10_[int(14)] = (d_60_v13_)[default__.safeIndex(324, (d_60_v13_).length(0))]
            nw10_[int(15)] = (d_60_v13_)[default__.safeIndex(324, (d_60_v13_).length(0))]
            nw10_[int(16)] = d_51_v5_
            d_72_v23_ = nw10_
            d_73_v24_: _dafny.Map
            d_73_v24_ = _dafny.Map({d_72_v23_: d_48_v2_})
            d_73_v24_ = ((d_73_v24_) | (d_73_v24_)) | (d_73_v24_)
        elif True:
            def lambda1_(d_74_v2_):
                def lambda2_(d_75_i2_):
                    return default__.safeDivisionInt(d_75_i2_, d_74_v2_)

                return lambda2_

            init1_ = lambda1_(d_48_v2_)
            nw11_ = _dafny.Array(None, 24)
            for i0_1_ in range(nw11_.length(0)):
                nw11_[i0_1_] = init1_(i0_1_)
            d_50_v4_ = nw11_
            d_76_v25_: _dafny.Array
            nw12_ = _dafny.Array(_dafny.MultiSet({}), 26)
            d_76_v25_ = nw12_
            d_77_v26_: _dafny.Map
            d_77_v26_ = _dafny.Map({(d_49_v3_).set(default__.safeIndex((0) - (len(d_46_v0_)), len(d_49_v3_)), d_48_v2_): d_76_v25_})
            d_78_v27_: _dafny.Map
            d_78_v27_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_54_v8_]): d_49_v3_})
            d_79_v28_: _dafny.Seq
            d_79_v28_ = _dafny.SeqWithoutIsStrInference([d_54_v8_])
            d_77_v26_ = (d_77_v26_).set(((d_78_v27_)[d_79_v28_] if (d_79_v28_) in (d_78_v27_) else d_49_v3_), d_76_v25_)
            d_80_v29_: _dafny.Set
            d_80_v29_ = _dafny.Set({d_60_v13_})
            d_48_v2_ = len(d_80_v29_)
            index11_ = default__.safeIndex(602, (d_60_v13_).length(0))
            (d_60_v13_)[index11_] = d_51_v5_
            index12_ = default__.safeIndex(602, (d_60_v13_).length(0))
            (d_60_v13_)[index12_] = d_51_v5_
            d_81_v30_: _dafny.Array
            nw13_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_81_v30_ = nw13_
            d_82_v31_: _dafny.Array
            nw14_ = _dafny.Array(None, 7)
            nw14_[int(0)] = d_50_v4_
            nw14_[int(1)] = d_50_v4_
            nw14_[int(2)] = d_50_v4_
            nw14_[int(3)] = d_50_v4_
            nw14_[int(4)] = d_50_v4_
            nw14_[int(5)] = d_50_v4_
            nw14_[int(6)] = d_50_v4_
            d_82_v31_ = nw14_
            d_83_v32_: _dafny.Seq
            d_83_v32_ = _dafny.SeqWithoutIsStrInference([d_82_v31_])
            index13_ = default__.safeIndex(588, (d_81_v30_).length(0))
            (d_81_v30_)[index13_] = (d_83_v32_)[default__.safeIndex(d_48_v2_, len(d_83_v32_))]
            d_84_v33_: _dafny.Map
            d_84_v33_ = _dafny.Map({d_48_v2_: len(d_54_v8_)})
            d_85_v34_: _dafny.MultiSet
            d_85_v34_ = _dafny.MultiSet([(d_60_v13_)[default__.safeIndex(602, (d_60_v13_).length(0))]])
            d_86_v35_: _dafny.Map
            d_86_v35_ = _dafny.Map({((d_85_v34_)[default__.fm0(d_48_v2_, (d_60_v13_)[default__.safeIndex(602, (d_60_v13_).length(0))], (d_60_v13_)[default__.safeIndex(602, (d_60_v13_).length(0))], d_46_v0_, d_59_globalState_)] if (default__.fm0(d_48_v2_, (d_60_v13_)[default__.safeIndex(602, (d_60_v13_).length(0))], (d_60_v13_)[default__.safeIndex(602, (d_60_v13_).length(0))], d_46_v0_, d_59_globalState_)) in (d_85_v34_) else d_48_v2_): (d_60_v13_)[default__.safeIndex(602, (d_60_v13_).length(0))]})
            d_87_v36_: _dafny.Set
            d_87_v36_ = _dafny.Set({len(d_52_v6_)})
            index14_ = default__.safeIndex(602, (d_60_v13_).length(0))
            index15_ = default__.safeIndex(588, (d_81_v30_).length(0))
            rhs11_ = True
            rhs12_ = d_82_v31_
            rhs13_ = _dafny.Map({default__.safeModuloInt(d_48_v2_, len(d_86_v35_)): d_48_v2_})
            rhs14_ = default__.safeDivisionInt(633, (d_48_v2_) - (default__.fm1(d_48_v2_, False, d_87_v36_, d_59_globalState_)))
            lhs7_ = d_60_v13_
            lhs8_ = default__.safeIndex(602, (d_60_v13_).length(0))
            lhs9_ = d_81_v30_
            lhs10_ = default__.safeIndex(588, (d_81_v30_).length(0))
            lhs11_ = d_59_globalState_
            lhs7_[lhs8_] = rhs11_
            lhs9_[lhs10_] = rhs12_
            d_84_v33_ = rhs13_
            lhs11_.f20 = rhs14_
        hi2_ = d_48_v2_
        for d_88_i3_ in range((941) + (d_48_v2_), hi2_):
            d_89_v37_: D0
            d_89_v37_ = D0_DC0(d_49_v3_)
            pat_let_tv0_ = d_49_v3_
            def iife1_(_pat_let0_0):
                def iife2_(d_90_dt__update__tmp_h0_):
                    def iife3_(_pat_let1_0):
                        def iife4_(d_91_dt__update_hcf0_h0_):
                            return D0_DC0(d_91_dt__update_hcf0_h0_)
                        return iife4_(_pat_let1_0)
                    return iife3_(pat_let_tv0_)
                return iife2_(_pat_let0_0)
            d_49_v3_ = (iife1_(d_89_v37_)).cf0
            d_92_v38_: D0
            d_92_v38_ = D0_DC3(d_50_v4_, d_88_i3_)
            source1_ = d_92_v38_
            if source1_.is_DC1:
                d_93___mcc_h0_ = source1_.cf1
                d_94___mcc_h1_ = source1_.cf2
                d_95___mcc_h2_ = source1_.cf3
                d_96_cf3_ = d_95___mcc_h2_
                d_97_cf2_ = d_94___mcc_h1_
                d_98_cf1_ = d_93___mcc_h0_
                d_50_v4_ = d_50_v4_
                index16_ = default__.safeIndex(928, (d_60_v13_).length(0))
                (d_60_v13_)[index16_] = not (d_97_cf2_) or (d_97_cf2_)
                d_99_v39_: _dafny.MultiSet
                d_99_v39_ = _dafny.MultiSet([not(d_51_v5_), ((0) - (d_48_v2_)) in (d_49_v3_)])
                index17_ = default__.safeIndex(928, (d_60_v13_).length(0))
                rhs15_ = not (d_51_v5_) or (d_97_cf2_)
                rhs16_ = (d_99_v39_) | (d_99_v39_)
                lhs12_ = d_60_v13_
                lhs13_ = default__.safeIndex(928, (d_60_v13_).length(0))
                lhs12_[lhs13_] = rhs15_
                d_99_v39_ = rhs16_
                index18_ = default__.safeIndex(928, (d_60_v13_).length(0))
                (d_60_v13_)[index18_] = (d_51_v5_) == ((d_88_i3_) >= (d_88_i3_))
                d_89_v37_ = (d_89_v37_ if d_97_cf2_ else default__.fm2(len(d_54_v8_), d_59_globalState_))
            elif source1_.is_DC2:
                d_100___mcc_h3_ = source1_.cf4
                d_101_cf4_ = d_100___mcc_h3_
                d_102_v40_: _dafny.Seq
                d_102_v40_ = _dafny.SeqWithoutIsStrInference([d_50_v4_, d_50_v4_])
                rhs17_ = _dafny.SeqWithoutIsStrInference([d_50_v4_, d_50_v4_, d_50_v4_, d_50_v4_])
                rhs18_ = d_60_v13_
                rhs19_ = False
                rhs20_ = (d_48_v2_) + ((d_88_i3_) - (d_48_v2_))
                lhs14_ = d_59_globalState_
                d_102_v40_ = rhs17_
                d_60_v13_ = rhs18_
                d_51_v5_ = rhs19_
                lhs14_.f20 = rhs20_
                d_51_v5_ = not (d_51_v5_) or ((not(d_51_v5_) if d_51_v5_ else d_51_v5_))
                (d_59_globalState_).f20 = d_48_v2_
                d_103_v41_: _dafny.Set
                d_103_v41_ = _dafny.Set({d_89_v37_, d_89_v37_})
                d_103_v41_ = (d_103_v41_) | (d_103_v41_)
            elif source1_.is_DC3:
                d_104___mcc_h4_ = source1_.cf5
                d_105___mcc_h5_ = source1_.cf6
                d_106_cf6_ = d_105___mcc_h5_
                d_107_cf5_ = d_104___mcc_h4_
                d_51_v5_ = (not(True)) or (not (not(d_51_v5_)) or (d_51_v5_))
                (d_59_globalState_).f20 = (d_88_i3_) * ((0) - ((d_88_i3_) - (538)))
                rhs21_ = (len(d_49_v3_)) + (d_106_cf6_)
                rhs22_ = d_51_v5_
                d_106_cf6_ = rhs21_
                d_51_v5_ = rhs22_
                d_108_v42_: _dafny.Set
                d_108_v42_ = _dafny.Set({d_48_v2_})
                (d_59_globalState_).f20 = (default__.fm1(d_88_i3_, d_51_v5_, d_108_v42_, d_59_globalState_)) - (d_88_i3_)
            elif source1_.is_DC0:
                d_109___mcc_h6_ = source1_.cf0
                d_110_cf0_ = d_109___mcc_h6_
                d_111_v43_: _dafny.Set
                d_111_v43_ = _dafny.Set({d_48_v2_, d_48_v2_, d_88_i3_})
                d_111_v43_ = d_111_v43_
                d_112_v44_: _dafny.Set
                d_112_v44_ = _dafny.Set({d_51_v5_, (d_54_v8_)[default__.safeIndex(494, len(d_54_v8_))], not((d_54_v8_)[default__.safeIndex(d_48_v2_, len(d_54_v8_))]), d_51_v5_})
                d_51_v5_ = (d_112_v44_).issubset(d_112_v44_)
                d_113_v45_: _dafny.Array
                nw15_ = _dafny.Array(_dafny.Seq({}), 19)
                d_113_v45_ = nw15_
                d_114_v46_: _dafny.Map
                d_114_v46_ = _dafny.Map({d_48_v2_: 303})
                d_115_v47_: _dafny.Seq
                d_115_v47_ = _dafny.SeqWithoutIsStrInference([d_114_v46_, d_114_v46_])
                index19_ = default__.safeIndex(335, (d_113_v45_).length(0))
                (d_113_v45_)[index19_] = (d_115_v47_) + (d_115_v47_)
                index20_ = default__.safeIndex(335, (d_113_v45_).length(0))
                (d_113_v45_)[index20_] = default__.fm3(d_59_globalState_)
                index21_ = default__.safeIndex(842, (d_60_v13_).length(0))
                (d_60_v13_)[index21_] = not((d_51_v5_ if d_51_v5_ else d_51_v5_))
                d_116_v48_: D0
                d_116_v48_ = D0_DC2(d_51_v5_)
                index22_ = default__.safeIndex(842, (d_60_v13_).length(0))
                (d_60_v13_)[index22_] = (d_116_v48_).cf4
            elif True:
                d_117___mcc_h7_ = source1_.cf7
                d_118_cf7_ = d_117___mcc_h7_
                d_119_v49_: _dafny.Map
                d_119_v49_ = _dafny.Map({d_88_i3_: d_88_i3_})
                d_120_v50_: C0
                nw16_ = C0()
                nw16_.ctor__((d_119_v49_) == (_dafny.Map({d_88_i3_: d_48_v2_})))
                d_120_v50_ = nw16_
                d_121_v51_: _dafny.MultiSet
                d_121_v51_ = _dafny.MultiSet([(d_120_v50_).f22])
                d_51_v5_ = (d_121_v51_).ispropersubset((_dafny.MultiSet(default__.fm5(d_59_globalState_))).set(d_51_v5_, default__.abs(d_48_v2_)))
                d_68_v19_ = d_68_v19_
                (d_59_globalState_).f20 = d_88_i3_
            (d_59_globalState_).f20 = (default__.safeDivisionInt(len(d_52_v6_), d_48_v2_) if d_51_v5_ else d_48_v2_)
            d_122_v52_: _dafny.Set
            d_122_v52_ = _dafny.Set({d_52_v6_})
            d_123_v53_: D1
            d_123_v53_ = D1_DC5(d_46_v0_)
            d_124_v54_: _dafny.Map
            d_124_v54_ = _dafny.Map({(d_122_v52_).ispropersubset(d_122_v52_): D0_DC3(d_50_v4_, len((d_123_v53_).cf8))})
            d_125_v55_: _dafny.MultiSet
            d_125_v55_ = _dafny.MultiSet([(d_51_v5_) == (d_51_v5_)])
            d_126_v56_: _dafny.Array
            def lambda3_(d_127_v0_):
                def lambda4_(d_128_i4_):
                    return d_127_v0_

                return lambda4_

            init2_ = lambda3_(d_46_v0_)
            nw17_ = _dafny.Array(None, 26)
            for i0_2_ in range(nw17_.length(0)):
                nw17_[i0_2_] = init2_(i0_2_)
            d_126_v56_ = nw17_
            d_129_v57_: D0
            d_129_v57_ = D0_DC1(d_50_v4_, d_51_v5_, d_126_v56_)
            rhs23_ = _dafny.Map({d_51_v5_: (d_92_v38_ if False else d_92_v38_)})
            rhs24_ = (_dafny.MultiSet(d_54_v8_)) | (_dafny.MultiSet([(d_129_v57_).cf2, d_51_v5_]))
            d_124_v54_ = rhs23_
            d_125_v55_ = rhs24_
        d_46_v0_ = d_46_v0_
        d_130_v58_: _dafny.Map
        out1_: _dafny.Map
        out1_ = default__.m0(d_59_globalState_)
        d_130_v58_ = out1_
        (d_59_globalState_).f20 = d_48_v2_
        d_131_v59_: D0
        d_131_v59_ = D0_DC2(False)
        d_132_i5_: int
        d_132_i5_ = 0
        with _dafny.label("0"):
            while not ((d_131_v59_).cf4) or (True):
                with _dafny.c_label("0"):
                    if (d_132_i5_) >= (100):
                        raise _dafny.Break("0")
                    d_132_i5_ = (d_132_i5_) + (1)
                    d_133_v60_: _dafny.Array
                    nw18_ = _dafny.Array(_dafny.MultiSet({}), 24)
                    d_133_v60_ = nw18_
                    index23_ = default__.safeIndex(801, (d_133_v60_).length(0))
                    (d_133_v60_)[index23_] = default__.fm6(d_59_globalState_)
                    d_134_v61_: _dafny.Map
                    d_134_v61_ = _dafny.Map({d_51_v5_: d_46_v0_})
                    d_135_v62_: _dafny.MultiSet
                    d_135_v62_ = _dafny.MultiSet([((d_134_v61_)[d_51_v5_] if (d_51_v5_) in (d_134_v61_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrw")))])
                    index24_ = default__.safeIndex(801, (d_133_v60_).length(0))
                    (d_133_v60_)[index24_] = ((d_135_v62_) | (d_135_v62_)).set(d_46_v0_, default__.abs(d_48_v2_))
                    (d_59_globalState_).f11 = d_46_v0_
                    d_51_v5_ = d_51_v5_
                    (d_59_globalState_).f20 = (0) - ((d_49_v3_)[default__.safeIndex(d_48_v2_, len(d_49_v3_))])
                    pass
            pass
        d_136_v63_: _dafny.Map
        out2_: _dafny.Map
        out2_ = default__.m0(d_59_globalState_)
        d_136_v63_ = out2_
        if not(not((d_48_v2_) != (27))):
            d_137_v64_: _dafny.Map
            out3_: _dafny.Map
            out3_ = default__.m0(d_59_globalState_)
            d_137_v64_ = out3_
            if d_51_v5_:
                d_50_v4_ = d_50_v4_
                d_51_v5_ = d_51_v5_
                d_138_v65_: D0
                d_138_v65_ = D0_DC0(d_49_v3_)
                d_139_v66_: _dafny.Seq
                d_139_v66_ = _dafny.SeqWithoutIsStrInference([D0_DC4(d_138_v65_), d_138_v65_, D0_DC4(d_138_v65_), d_138_v65_, d_138_v65_])
                d_140_v67_: _dafny.Map
                d_140_v67_ = _dafny.Map({d_53_v7_: (d_139_v66_)[default__.safeIndex(257, len(d_139_v66_))]})
                d_141_v68_: D0
                d_141_v68_ = D0_DC4(((d_140_v67_)[d_53_v7_] if (d_53_v7_) in (d_140_v67_) else d_138_v65_))
                d_142_v69_: _dafny.Array
                nw19_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
                d_142_v69_ = nw19_
                d_143_v70_: D0
                d_143_v70_ = D0_DC4((d_138_v65_ if d_51_v5_ else D0_DC1(d_50_v4_, d_51_v5_, d_142_v69_)))
                d_143_v70_ = D0_DC4(d_138_v65_)
                d_144_v71_: _dafny.Set
                d_144_v71_ = _dafny.Set({d_51_v5_})
                d_145_v72_: _dafny.Seq
                d_145_v72_ = _dafny.SeqWithoutIsStrInference([d_144_v71_, d_144_v71_])
                d_146_v73_: _dafny.Array
                nw20_ = _dafny.Array(None, 9)
                nw20_[int(0)] = _dafny.SeqWithoutIsStrInference([d_48_v2_])
                nw20_[int(1)] = _dafny.SeqWithoutIsStrInference([d_48_v2_])
                nw20_[int(2)] = d_49_v3_
                nw20_[int(3)] = (d_49_v3_) + (_dafny.SeqWithoutIsStrInference([len(d_144_v71_)]))
                nw20_[int(4)] = (d_49_v3_) + ((_dafny.SeqWithoutIsStrInference([d_48_v2_])).set(default__.safeIndex(d_48_v2_, len(_dafny.SeqWithoutIsStrInference([d_48_v2_]))), len(_dafny.SeqWithoutIsStrInference([d_54_v8_, d_54_v8_, d_54_v8_, d_54_v8_]))))
                nw20_[int(5)] = _dafny.SeqWithoutIsStrInference([d_48_v2_, d_48_v2_])
                nw20_[int(6)] = d_49_v3_
                nw20_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_48_v2_ for d_147_i6_ in range(default__.abs(579))])).set(default__.safeIndex((0) - (d_48_v2_), len(_dafny.SeqWithoutIsStrInference([d_48_v2_ for d_148_i6_ in range(default__.abs(579))]))), d_48_v2_)
                nw20_[int(8)] = (_dafny.SeqWithoutIsStrInference([d_48_v2_])).set(default__.safeIndex((0) - (d_48_v2_), len(_dafny.SeqWithoutIsStrInference([d_48_v2_]))), len(d_145_v72_))
                d_146_v73_ = nw20_
                index25_ = default__.safeIndex(705, (d_146_v73_).length(0))
                (d_146_v73_)[index25_] = d_49_v3_
                index26_ = default__.safeIndex(705, (d_146_v73_).length(0))
                (d_146_v73_)[index26_] = ((d_49_v3_).set(default__.safeIndex(d_48_v2_, len(d_49_v3_)), len(_dafny.SeqWithoutIsStrInference([d_48_v2_])))) + (d_49_v3_)
                d_149_v74_: _dafny.Map
                d_149_v74_ = _dafny.Map({d_48_v2_: d_48_v2_})
                d_150_v75_: _dafny.Set
                d_150_v75_ = _dafny.Set({d_48_v2_, d_48_v2_})
                (d_59_globalState_).f20 = (0) - (default__.fm1(((d_149_v74_)[d_48_v2_] if (d_48_v2_) in (d_149_v74_) else default__.fm1(d_48_v2_, True, d_150_v75_, d_59_globalState_)), (_dafny.Set({d_51_v5_})) != (d_144_v71_), d_150_v75_, d_59_globalState_))
            elif True:
                d_151_v76_: C0
                nw21_ = C0()
                nw21_.ctor__(d_51_v5_)
                d_151_v76_ = nw21_
                d_152_v77_: _dafny.Map
                d_152_v77_ = _dafny.Map({len(d_49_v3_): d_151_v76_})
                (d_59_globalState_).f20 = (len((d_152_v77_ if d_51_v5_ else d_152_v77_))) - (61)
                nw22_ = C0()
                nw22_.ctor__((d_151_v76_).f22)
                d_151_v76_ = nw22_
                d_153_v78_: _dafny.Map
                out4_: _dafny.Map
                out4_ = default__.m0(d_59_globalState_)
                d_153_v78_ = out4_
                d_154_v79_: _dafny.Map
                out5_: _dafny.Map
                out5_ = default__.m0(d_59_globalState_)
                d_154_v79_ = out5_
                d_155_v80_: _dafny.Array
                nw23_ = _dafny.Array(_dafny.Set({}), 5)
                d_155_v80_ = nw23_
                d_156_v81_: _dafny.Set
                d_156_v81_ = _dafny.Set({d_48_v2_})
                index27_ = default__.safeIndex(399, (d_155_v80_).length(0))
                (d_155_v80_)[index27_] = (_dafny.Set({d_48_v2_, len(d_49_v3_)})) - (d_156_v81_)
                index28_ = default__.safeIndex(399, (d_155_v80_).length(0))
                (d_155_v80_)[index28_] = (d_156_v81_) - (d_156_v81_)
            d_51_v5_ = d_51_v5_
            if d_51_v5_:
                d_157_v82_: _dafny.Map
                out6_: _dafny.Map
                out6_ = default__.m0(d_59_globalState_)
                d_157_v82_ = out6_
                d_158_v83_: _dafny.MultiSet
                d_158_v83_ = _dafny.MultiSet([d_50_v4_])
                d_159_v84_: _dafny.Map
                d_159_v84_ = _dafny.Map({default__.safeModuloInt((0) - (d_48_v2_), ((d_158_v83_)[d_50_v4_] if (d_50_v4_) in (d_158_v83_) else (0) - (d_48_v2_))): d_51_v5_})
                d_159_v84_ = (d_159_v84_).set(d_48_v2_, d_51_v5_)
                d_51_v5_ = d_51_v5_
                d_160_v85_: C0
                nw24_ = C0()
                nw24_.ctor__((d_54_v8_)[default__.safeIndex(d_48_v2_, len(d_54_v8_))])
                d_160_v85_ = nw24_
                d_161_v86_: _dafny.Map
                d_161_v86_ = _dafny.Map({d_160_v85_: d_51_v5_})
                d_162_v87_: _dafny.Set
                d_162_v87_ = _dafny.Set({(d_161_v86_).set(d_160_v85_, d_51_v5_), d_161_v86_, d_161_v86_})
                d_163_v88_: _dafny.Set
                d_163_v88_ = _dafny.Set({d_48_v2_})
                d_164_v89_: _dafny.Map
                d_164_v89_ = _dafny.Map({default__.safeModuloInt(default__.fm1(len(d_162_v87_), d_51_v5_, d_163_v88_, d_59_globalState_), len(_dafny.SeqWithoutIsStrInference([d_53_v7_ for d_165_i7_ in range(default__.abs(328))]))): (d_48_v2_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkst"))))})
                d_164_v89_ = d_164_v89_
                d_51_v5_ = (d_46_v0_) <= (_dafny.SeqWithoutIsStrInference([d_53_v7_]))
            elif True:
                (d_59_globalState_).f20 = (0) - (d_48_v2_)
                d_51_v5_ = d_51_v5_
                d_166_v90_: _dafny.Map
                out7_: _dafny.Map
                out7_ = default__.m0(d_59_globalState_)
                d_166_v90_ = out7_
                d_167_v91_: _dafny.Map
                d_167_v91_ = _dafny.Map({default__.safeModuloInt(len(d_52_v6_), d_48_v2_): d_53_v7_})
                index29_ = default__.safeIndex(596, (d_60_v13_).length(0))
                (d_60_v13_)[index29_] = d_51_v5_
                d_168_v92_: C0
                nw25_ = C0()
                nw25_.ctor__(d_51_v5_)
                d_168_v92_ = nw25_
                index30_ = default__.safeIndex(262, (d_60_v13_).length(0))
                (d_60_v13_)[index30_] = ((0) - (d_48_v2_)) != (d_48_v2_)
                d_169_v94_: _dafny.MultiSet
                d_169_v94_ = _dafny.MultiSet([d_48_v2_])
                d_170_v95_: _dafny.Map
                d_170_v95_ = _dafny.Map({d_50_v4_: d_168_v92_})
                index31_ = default__.safeIndex(596, (d_60_v13_).length(0))
                index32_ = default__.safeIndex(262, (d_60_v13_).length(0))
                def iife5_():
                    coll1_ = _dafny.Map()
                    compr_1_: int
                    for compr_1_ in (d_169_v94_).Elements:
                        d_171_v93_: int = compr_1_
                        if (d_171_v93_) in (d_169_v94_):
                            coll1_[default__.safeModuloInt(d_171_v93_, d_48_v2_)] = d_53_v7_
                    return _dafny.Map(coll1_)
                rhs25_ = iife5_()

                rhs26_ = (default__.fm7((d_168_v92_).f22, d_48_v2_, d_59_globalState_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqkka"))).set(default__.safeIndex(len(d_52_v6_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqkka")))), (d_46_v0_)[default__.safeIndex(d_48_v2_, len(d_46_v0_))]))
                rhs27_ = d_51_v5_
                rhs28_ = d_168_v92_
                rhs29_ = (_dafny.Map({d_50_v4_: d_168_v92_})) == (d_170_v95_)
                lhs15_ = d_59_globalState_
                lhs16_ = d_60_v13_
                lhs17_ = default__.safeIndex(596, (d_60_v13_).length(0))
                lhs18_ = d_60_v13_
                lhs19_ = default__.safeIndex(262, (d_60_v13_).length(0))
                d_167_v91_ = rhs25_
                lhs15_.f0 = rhs26_
                lhs16_[lhs17_] = rhs27_
                d_168_v92_ = rhs28_
                lhs18_[lhs19_] = rhs29_
                index33_ = default__.safeIndex(698, (d_50_v4_).length(0))
                (d_50_v4_)[index33_] = d_48_v2_
                d_172_v96_: _dafny.Set
                d_172_v96_ = _dafny.Set({d_48_v2_})
                index34_ = default__.safeIndex(698, (d_50_v4_).length(0))
                (d_50_v4_)[index34_] = default__.fm1(-810, (_dafny.SeqWithoutIsStrInference([d_48_v2_ for d_173_i8_ in range(default__.abs(-288))])) <= (d_49_v3_), d_172_v96_, d_59_globalState_)
            d_48_v2_ = d_48_v2_
        elif True:
            rhs30_ = d_53_v7_
            rhs31_ = d_51_v5_
            rhs32_ = default__.safeModuloInt(d_48_v2_, d_48_v2_)
            lhs20_ = d_59_globalState_
            d_53_v7_ = rhs30_
            d_51_v5_ = rhs31_
            lhs20_.f20 = rhs32_
            d_51_v5_ = d_51_v5_
            d_48_v2_ = d_48_v2_
            d_48_v2_ = d_48_v2_
            if (((0) - ((0) - (d_48_v2_)) if d_51_v5_ else d_48_v2_)) < (len(d_46_v0_)):
                d_53_v7_ = d_53_v7_
                d_174_v97_: _dafny.Map
                d_174_v97_ = _dafny.Map({default__.safeModuloInt(d_48_v2_, d_48_v2_): default__.safeDivisionInt(-178, d_48_v2_)})
                d_174_v97_ = d_174_v97_
                d_175_v98_: D0
                d_175_v98_ = D0_DC3(d_50_v4_, (0) - ((d_48_v2_) * (d_48_v2_)))
                d_176_v99_: _dafny.Array
                nw26_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_176_v99_ = nw26_
                d_177_v100_: _dafny.Set
                d_177_v100_ = _dafny.Set({d_48_v2_, d_48_v2_})
                index35_ = default__.safeIndex(699, (d_50_v4_).length(0))
                (d_50_v4_)[index35_] = (default__.fm1(d_48_v2_, d_51_v5_, d_177_v100_, d_59_globalState_)) - (d_48_v2_)
                d_178_v101_: _dafny.Map
                d_178_v101_ = _dafny.Map({d_49_v3_: d_48_v2_})
                index36_ = default__.safeIndex(699, (d_50_v4_).length(0))
                rhs33_ = d_175_v98_
                rhs34_ = (d_176_v99_ if not (not(d_51_v5_)) or (d_51_v5_) else d_176_v99_)
                rhs35_ = default__.fm1((len(d_178_v101_)) + (d_48_v2_), (d_46_v0_) != (d_46_v0_), d_177_v100_, d_59_globalState_)
                rhs36_ = (d_48_v2_) * (d_48_v2_)
                rhs37_ = default__.safeDivisionInt(d_48_v2_, d_48_v2_)
                lhs21_ = d_50_v4_
                lhs22_ = default__.safeIndex(699, (d_50_v4_).length(0))
                lhs23_ = d_59_globalState_
                d_175_v98_ = rhs33_
                d_176_v99_ = rhs34_
                lhs21_[lhs22_] = rhs35_
                lhs23_.f20 = rhs36_
                d_48_v2_ = rhs37_
                d_179_v102_: C0
                nw27_ = C0()
                nw27_.ctor__(True)
                d_179_v102_ = nw27_
                d_180_v103_: _dafny.Map
                d_180_v103_ = _dafny.Map({True: d_54_v8_})
                d_181_v104_: D2
                d_181_v104_ = D2_DC8(_dafny.SeqWithoutIsStrInference([d_51_v5_, d_51_v5_]))
                d_182_v105_: _dafny.Array
                nw28_ = _dafny.Array(None, 27)
                nw28_[int(0)] = d_54_v8_
                nw28_[int(1)] = _dafny.SeqWithoutIsStrInference([d_51_v5_])
                nw28_[int(2)] = (d_54_v8_) + (d_54_v8_)
                nw28_[int(3)] = d_54_v8_
                nw28_[int(4)] = d_54_v8_
                nw28_[int(5)] = d_54_v8_
                nw28_[int(6)] = (default__.fm5(d_59_globalState_)) + (d_54_v8_)
                nw28_[int(7)] = d_54_v8_
                nw28_[int(8)] = d_54_v8_
                nw28_[int(9)] = _dafny.SeqWithoutIsStrInference([d_51_v5_])
                nw28_[int(10)] = d_54_v8_
                nw28_[int(11)] = _dafny.SeqWithoutIsStrInference([True])
                nw28_[int(12)] = d_54_v8_
                nw28_[int(13)] = (d_54_v8_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fllbype"))), len(d_54_v8_)), d_51_v5_)
                nw28_[int(14)] = _dafny.SeqWithoutIsStrInference([d_51_v5_, (d_179_v102_).f22, not(not(not(d_51_v5_))), d_51_v5_, d_51_v5_])
                nw28_[int(15)] = d_54_v8_
                nw28_[int(16)] = _dafny.SeqWithoutIsStrInference([d_51_v5_])
                nw28_[int(17)] = ((d_180_v103_)[(d_179_v102_).f22] if ((d_179_v102_).f22) in (d_180_v103_) else _dafny.SeqWithoutIsStrInference([d_51_v5_]))
                nw28_[int(18)] = _dafny.SeqWithoutIsStrInference([not(d_51_v5_)])
                nw28_[int(19)] = (d_181_v104_).cf13
                nw28_[int(20)] = d_54_v8_
                nw28_[int(21)] = d_54_v8_
                nw28_[int(22)] = d_54_v8_
                nw28_[int(23)] = (d_54_v8_) + (d_54_v8_)
                nw28_[int(24)] = d_54_v8_
                nw28_[int(25)] = (d_54_v8_) + (_dafny.SeqWithoutIsStrInference([d_51_v5_, False, not(d_51_v5_), d_51_v5_]))
                nw28_[int(26)] = d_54_v8_
                d_182_v105_ = nw28_
                index37_ = default__.safeIndex(828, (d_182_v105_).length(0))
                (d_182_v105_)[index37_] = (default__.fm5(d_59_globalState_)) + (d_54_v8_)
                index38_ = default__.safeIndex(828, (d_182_v105_).length(0))
                (d_182_v105_)[index38_] = (default__.fm5(d_59_globalState_)) + ((d_54_v8_) + (d_54_v8_))
            elif True:
                index39_ = default__.safeIndex(557, (d_50_v4_).length(0))
                (d_50_v4_)[index39_] = len(d_54_v8_)
                d_183_v106_: _dafny.Array
                def lambda5_(d_184_v0_):
                    def lambda6_(d_185_i9_):
                        return d_184_v0_

                    return lambda6_

                init3_ = lambda5_(d_46_v0_)
                nw29_ = _dafny.Array(None, 17)
                for i0_3_ in range(nw29_.length(0)):
                    nw29_[i0_3_] = init3_(i0_3_)
                d_183_v106_ = nw29_
                index40_ = default__.safeIndex(696, (d_183_v106_).length(0))
                (d_183_v106_)[index40_] = default__.fm7(d_51_v5_, d_48_v2_, d_59_globalState_)
                d_186_v107_: _dafny.Set
                d_186_v107_ = _dafny.Set({_dafny.Map({d_51_v5_: d_48_v2_})})
                index41_ = default__.safeIndex(557, (d_50_v4_).length(0))
                index42_ = default__.safeIndex(696, (d_183_v106_).length(0))
                rhs38_ = d_46_v0_
                rhs39_ = default__.safeDivisionInt(default__.safeModuloInt(len(d_186_v107_), d_48_v2_), len(d_46_v0_))
                rhs40_ = _dafny.SeqWithoutIsStrInference([d_53_v7_ for d_187_i10_ in range(default__.abs(723))])
                rhs41_ = d_50_v4_
                rhs42_ = d_48_v2_
                lhs24_ = d_50_v4_
                lhs25_ = default__.safeIndex(557, (d_50_v4_).length(0))
                lhs26_ = d_183_v106_
                lhs27_ = default__.safeIndex(696, (d_183_v106_).length(0))
                lhs28_ = d_59_globalState_
                d_46_v0_ = rhs38_
                lhs24_[lhs25_] = rhs39_
                lhs26_[lhs27_] = rhs40_
                d_50_v4_ = rhs41_
                lhs28_.f20 = rhs42_
                d_188_v108_: _dafny.Map
                out8_: _dafny.Map
                out8_ = default__.m0(d_59_globalState_)
                d_188_v108_ = out8_
                index43_ = default__.safeIndex(127, (d_60_v13_).length(0))
                (d_60_v13_)[index43_] = d_51_v5_
                index44_ = default__.safeIndex(127, (d_60_v13_).length(0))
                (d_60_v13_)[index44_] = d_51_v5_
                d_189_v109_: _dafny.MultiSet
                d_189_v109_ = _dafny.MultiSet([d_48_v2_])
                d_189_v109_ = d_189_v109_
                d_190_v110_: _dafny.Set
                d_190_v110_ = _dafny.Set({d_131_v59_})
                pat_let_tv1_ = d_60_v13_
                pat_let_tv2_ = d_60_v13_
                d_191_v111_: _dafny.Map
                def iife6_(_pat_let2_0):
                    def iife7_(d_192_dt__update__tmp_h1_):
                        def iife8_(_pat_let3_0):
                            def iife9_(d_193_dt__update_hcf4_h0_):
                                return D0_DC2(d_193_dt__update_hcf4_h0_)
                            return iife9_(_pat_let3_0)
                        return iife8_((pat_let_tv2_)[default__.safeIndex(127, (pat_let_tv1_).length(0))])
                    return iife7_(_pat_let2_0)
                d_191_v111_ = _dafny.Map({(iife6_(d_131_v59_)) not in (d_190_v110_): not (not(not((d_60_v13_)[default__.safeIndex(127, (d_60_v13_).length(0))]))) or (not(False))})
                d_191_v111_ = (d_191_v111_).set(False, default__.fm0((0) - (d_48_v2_), not((d_60_v13_)[default__.safeIndex(127, (d_60_v13_).length(0))]), default__.fm0(d_48_v2_, d_51_v5_, d_51_v5_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")), d_59_globalState_), (d_183_v106_)[default__.safeIndex(696, (d_183_v106_).length(0))], d_59_globalState_))
        d_194_v112_: _dafny.Map
        out9_: _dafny.Map
        out9_ = default__.m0(d_59_globalState_)
        d_194_v112_ = out9_
        d_195_v113_: _dafny.Map
        d_195_v113_ = _dafny.Map({d_48_v2_: d_49_v3_})
        d_196_v114_: _dafny.Map
        d_196_v114_ = _dafny.Map({d_48_v2_: d_48_v2_})
        rhs43_ = (d_195_v113_).set(d_48_v2_, _dafny.SeqWithoutIsStrInference([((d_196_v114_)[d_48_v2_] if (d_48_v2_) in (d_196_v114_) else d_48_v2_)]))
        rhs44_ = d_48_v2_
        rhs45_ = ((d_46_v0_) + (d_46_v0_)) not in (d_58_v12_)
        lhs29_ = d_59_globalState_
        d_195_v113_ = rhs43_
        lhs29_.f20 = rhs44_
        d_51_v5_ = rhs45_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_50_v4_).length(0)):
            d_197_i11_: int = guard_loop_0_
            if (True) and (((0) <= (d_197_i11_)) and ((d_197_i11_) < ((d_50_v4_).length(0)))):
                (d_50_v4_)[(d_197_i11_)] = (d_197_i11_) * (d_48_v2_)
        d_198_i12_: int
        d_198_i12_ = 0
        with _dafny.label("1"):
            while d_51_v5_:
                with _dafny.c_label("1"):
                    if (d_198_i12_) >= (100):
                        raise _dafny.Break("1")
                    d_198_i12_ = (d_198_i12_) + (1)
                    d_199_v115_: C0
                    nw30_ = C0()
                    nw30_.ctor__(d_51_v5_)
                    d_199_v115_ = nw30_
                    (d_59_globalState_).f20 = 612
                    (d_59_globalState_).f20 = d_48_v2_
                    if (d_199_v115_).f22:
                        d_200_v116_: _dafny.Seq
                        d_200_v116_ = _dafny.SeqWithoutIsStrInference([d_199_v115_, d_199_v115_])
                        d_201_v117_: _dafny.Set
                        d_201_v117_ = _dafny.Set({len((_dafny.SeqWithoutIsStrInference([d_199_v115_])) + (d_200_v116_)), 73, d_48_v2_})
                        d_201_v117_ = d_201_v117_
                        (d_59_globalState_).f20 = (d_48_v2_) - (default__.safeDivisionInt(d_48_v2_, d_48_v2_))
                        d_202_v118_: _dafny.MultiSet
                        d_202_v118_ = _dafny.MultiSet([True])
                        d_203_v119_: _dafny.Map
                        d_203_v119_ = _dafny.Map({d_48_v2_: d_202_v118_})
                        d_203_v119_ = (d_203_v119_ if (d_48_v2_) >= (50) else d_203_v119_)
                        d_204_v120_: C0
                        nw31_ = C0()
                        nw31_.ctor__(True)
                        d_204_v120_ = nw31_
                        d_205_v121_: _dafny.Map
                        out10_: _dafny.Map
                        out10_ = default__.m0(d_59_globalState_)
                        d_205_v121_ = out10_
                    elif True:
                        d_206_v122_: _dafny.Array
                        def lambda7_(d_207_v0_):
                            def lambda8_(d_208_i13_):
                                return d_207_v0_

                            return lambda8_

                        init4_ = lambda7_(d_46_v0_)
                        nw32_ = _dafny.Array(None, 12)
                        for i0_4_ in range(nw32_.length(0)):
                            nw32_[i0_4_] = init4_(i0_4_)
                        d_206_v122_ = nw32_
                        d_209_v123_: D0
                        d_209_v123_ = D0_DC1(d_50_v4_, not(False), d_206_v122_)
                        rhs46_ = (((d_199_v115_).f22) and ((d_199_v115_).f22)) or ((d_209_v123_).cf2)
                        rhs47_ = d_48_v2_
                        rhs48_ = d_60_v13_
                        rhs49_ = d_48_v2_
                        lhs30_ = d_59_globalState_
                        lhs31_ = d_59_globalState_
                        d_51_v5_ = rhs46_
                        lhs30_.f20 = rhs47_
                        d_60_v13_ = rhs48_
                        lhs31_.f20 = rhs49_
                        (d_59_globalState_).f20 = d_48_v2_
                        d_210_v125_: _dafny.Seq
                        def iife10_():
                            coll2_ = _dafny.Map()
                            compr_2_: int
                            for compr_2_ in _dafny.IntegerRange(266, 17):
                                d_211_v124_: int = compr_2_
                                if ((266) <= (d_211_v124_)) and ((d_211_v124_) < (17)):
                                    coll2_[(d_211_v124_) - (len(d_46_v0_))] = d_48_v2_
                            return _dafny.Map(coll2_)
                        d_210_v125_ = _dafny.SeqWithoutIsStrInference([(iife10_()
                        ) | (d_196_v114_), (d_196_v114_) | (_dafny.Map({d_48_v2_: len(_dafny.Set({d_48_v2_}))})), d_196_v114_])
                        d_212_v126_: _dafny.Map
                        d_212_v126_ = _dafny.Map({D0_DC2((d_199_v115_).f22): d_210_v125_})
                        d_210_v125_ = ((d_212_v126_)[D0_DC2((d_199_v115_).f22)] if (D0_DC2((d_199_v115_).f22)) in (d_212_v126_) else d_210_v125_)
                        d_213_v127_: _dafny.Map
                        d_213_v127_ = _dafny.Map({d_48_v2_: d_60_v13_})
                        d_213_v127_ = d_213_v127_
                        d_214_v128_: D0
                        d_214_v128_ = D0_DC0(d_49_v3_)
                        d_215_v129_: D0
                        d_215_v129_ = D0_DC4(d_214_v128_)
                        pat_let_tv3_ = d_199_v115_
                        pat_let_tv4_ = d_214_v128_
                        d_216_v130_: _dafny.Array
                        nw33_ = _dafny.Array(None, 17)
                        nw33_[int(0)] = d_215_v129_
                        nw33_[int(1)] = d_215_v129_
                        nw33_[int(2)] = d_215_v129_
                        nw33_[int(3)] = D0_DC4(d_214_v128_)
                        nw33_[int(4)] = d_215_v129_
                        nw33_[int(5)] = d_215_v129_
                        nw33_[int(6)] = d_215_v129_
                        nw33_[int(7)] = d_215_v129_
                        nw33_[int(8)] = d_215_v129_
                        nw33_[int(9)] = d_215_v129_
                        def iife11_(_pat_let4_0):
                            def iife12_(d_217_dt__update__tmp_h2_):
                                def iife13_(_pat_let5_0):
                                    def iife14_(d_218_dt__update_hcf7_h0_):
                                        return D0_DC4(d_218_dt__update_hcf7_h0_)
                                    return iife14_(_pat_let5_0)
                                return iife13_(D0_DC2((pat_let_tv3_).f22))
                            return iife12_(_pat_let4_0)
                        nw33_[int(10)] = iife11_(d_215_v129_)
                        def iife15_(_pat_let6_0):
                            def iife16_(d_219_dt__update__tmp_h3_):
                                def iife17_(_pat_let7_0):
                                    def iife18_(d_220_dt__update_hcf7_h1_):
                                        return D0_DC4(d_220_dt__update_hcf7_h1_)
                                    return iife18_(_pat_let7_0)
                                return iife17_(pat_let_tv4_)
                            return iife16_(_pat_let6_0)
                        nw33_[int(11)] = iife15_(d_215_v129_)
                        nw33_[int(12)] = d_215_v129_
                        nw33_[int(13)] = d_215_v129_
                        nw33_[int(14)] = d_215_v129_
                        nw33_[int(15)] = D0_DC4(d_214_v128_)
                        nw33_[int(16)] = d_215_v129_
                        d_216_v130_ = nw33_
                        index45_ = default__.safeIndex(393, (d_216_v130_).length(0))
                        (d_216_v130_)[index45_] = d_215_v129_
                        index46_ = default__.safeIndex(393, (d_216_v130_).length(0))
                        (d_216_v130_)[index46_] = ((d_215_v129_ if (d_199_v115_).f22 else d_215_v129_) if d_51_v5_ else d_215_v129_)
                    pass
            pass
        _dafny.print((d_46_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_48_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_49_v3_) == (_dafny.SeqWithoutIsStrInference([135]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v4_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_51_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_52_v6_) == (_dafny.Map({False: 135}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_53_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v8_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_55_v9_) == (_dafny.Map({_dafny.CodePoint('o'): 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_56_v10_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('o'): 9}), _dafny.Map({_dafny.CodePoint('t'): 135})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_57_v11_) == (_dafny.Map({135: _dafny.Map({_dafny.CodePoint('t'): 135})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_58_v12_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljlofb"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_59_globalState_.f0).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_59_globalState_).f1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_59_globalState_).f4) == (_dafny.SeqWithoutIsStrInference([135]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_59_globalState_).f9) == (_dafny.Map({False: 584}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_59_globalState_.f11).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_59_globalState_).f12) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_59_globalState_).f14) == (_dafny.SeqWithoutIsStrInference([135]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_59_globalState_).f18) == (_dafny.Map({135: _dafny.Map({_dafny.CodePoint('t'): 135})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_globalState_.f19) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljlofb"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_59_globalState_.f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_v13_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_v13_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_v13_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_v13_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_v13_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_v13_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_v13_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_60_v13_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_62_v14_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_68_v19_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_130_v58_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v59_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_132_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_136_v63_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_194_v112_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v113_) == (_dafny.Map({1: _dafny.SeqWithoutIsStrInference([1])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_196_v114_) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_198_i12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.Array(None, 0), False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D0_DC4)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC4(D0, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC6(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D1_DC7)

class D1_DC6(D1, NamedTuple('DC6', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({self.cf8.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC7(D1, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC9(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC9(D2, NamedTuple('DC9', [('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC11(None, None, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC11(D3, NamedTuple('DC11', [('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f11: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f19: _dafny.Seq = _dafny.Seq({})
        self.f20: int = int(0)
        self._f1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f2: bool = False
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f4: _dafny.Seq = _dafny.Seq({})
        self._f5: bool = False
        self._f6: int = int(0)
        self._f7: bool = False
        self._f8: _dafny.Array = _dafny.Array(None, 0)
        self._f9: _dafny.Map = _dafny.Map({})
        self._f10: int = int(0)
        self._f12: _dafny.Seq = _dafny.Seq({})
        self._f13: int = int(0)
        self._f14: _dafny.Seq = _dafny.Seq({})
        self._f15: _dafny.Array = _dafny.Array(None, 0)
        self._f16: bool = False
        self._f17: int = int(0)
        self._f18: _dafny.Map = _dafny.Map({})
        self._f21: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f19 = f19
        (self).f20 = f20
        (self)._f21 = f21

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
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
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
    @property
    def f18(self):
        return self._f18
    @property
    def f21(self):
        return self._f21

class C0:
    def  __init__(self):
        self._f22: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f22):
        (self)._f22 = f22

    def fm4(self, p0, p1, p2, p3, globalState):
        return (self).f22

    @property
    def f22(self):
        return self._f22
