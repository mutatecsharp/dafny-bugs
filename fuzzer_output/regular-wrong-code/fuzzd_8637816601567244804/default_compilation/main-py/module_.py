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
        return True

    @staticmethod
    def fm2(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_0_i0_ in range(default__.abs(719))])

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return (658) - (default__.safeModuloInt(788, len(_dafny.SeqWithoutIsStrInference([161]))))

    @staticmethod
    def fm4(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([not(not(not(False))), (False) in (_dafny.Map({not(False): 971}))])

    @staticmethod
    def fm5(p0, globalState):
        source0_ = D0_DC1(True, _dafny.Map({False: 435}), False, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1_i0_ in range(default__.abs(192))]))), 848)
        if source0_.is_DC1:
            d_2___mcc_h0_ = source0_.cf1
            d_3___mcc_h1_ = source0_.cf2
            d_4___mcc_h2_ = source0_.cf3
            d_5___mcc_h3_ = source0_.cf4
            d_6___mcc_h4_ = source0_.cf5
            d_7_cf5_ = d_6___mcc_h4_
            d_8_cf4_ = d_5___mcc_h3_
            d_9_cf3_ = d_4___mcc_h2_
            d_10_cf2_ = d_3___mcc_h1_
            d_11_cf1_ = d_2___mcc_h0_
            return _dafny.CodePoint('p')
        elif source0_.is_DC2:
            d_12___mcc_h5_ = source0_.cf6
            d_13___mcc_h6_ = source0_.cf7
            d_14___mcc_h7_ = source0_.cf8
            d_15___mcc_h8_ = source0_.cf9
            d_16_cf9_ = d_15___mcc_h8_
            d_17_cf8_ = d_14___mcc_h7_
            d_18_cf7_ = d_13___mcc_h6_
            d_19_cf6_ = d_12___mcc_h5_
            return _dafny.CodePoint('e')
        elif source0_.is_DC0:
            d_20___mcc_h9_ = source0_.cf0
            d_21_cf0_ = d_20___mcc_h9_
            return _dafny.CodePoint('u')
        elif True:
            d_22___mcc_h10_ = source0_.cf10
            d_23_cf10_ = d_22___mcc_h10_
            return _dafny.CodePoint('i')

    @staticmethod
    def m0(p0, globalState):
        d_24_v0_: _dafny.Seq
        d_24_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
        d_24_v0_ = d_24_v0_
        d_25_v2_: int
        d_25_v2_ = 795
        d_26_v3_: _dafny.Map
        d_26_v3_ = _dafny.Map({d_25_v2_: d_25_v2_})
        d_27_v4_: _dafny.Map
        d_27_v4_ = _dafny.Map({d_26_v3_: True})
        d_28_v5_: _dafny.Set
        d_28_v5_ = _dafny.Set({d_25_v2_, len(d_27_v4_)})
        d_29_v6_: _dafny.MultiSet
        d_29_v6_ = _dafny.MultiSet([d_25_v2_, d_25_v2_, default__.fm3(d_25_v2_, p0, p0, globalState), len(d_24_v0_)])
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (d_28_v5_).Elements:
                d_30_v1_: int = compr_0_
                if (d_30_v1_) in (d_28_v5_):
                    coll0_[(d_30_v1_) + (d_25_v2_)] = p0
            return _dafny.Map(coll0_)
        d_24_v0_ = (default__.fm2(len(iife0_()
        ), d_29_v6_, globalState)) + (d_24_v0_)
        d_31_v7_: _dafny.Array
        nw0_ = _dafny.Array(None, 11)
        nw0_[int(0)] = p0
        nw0_[int(1)] = p0
        nw0_[int(2)] = not(p0)
        nw0_[int(3)] = p0
        nw0_[int(4)] = p0
        nw0_[int(5)] = p0
        nw0_[int(6)] = p0
        nw0_[int(7)] = (p0) == (p0)
        nw0_[int(8)] = p0
        nw0_[int(9)] = p0
        nw0_[int(10)] = p0
        d_31_v7_ = nw0_
        index0_ = default__.safeIndex(469, (d_31_v7_).length(0))
        (d_31_v7_)[index0_] = p0
        index1_ = default__.safeIndex(469, (d_31_v7_).length(0))
        (d_31_v7_)[index1_] = not(((d_25_v2_) < (d_25_v2_)) or ((_dafny.SeqWithoutIsStrInference([p0])) == (default__.fm4(d_25_v2_, globalState))))
        d_32_v8_: _dafny.Array
        nw1_ = _dafny.Array(int(0), 24)
        d_32_v8_ = nw1_
        d_33_v9_: _dafny.Map
        d_33_v9_ = _dafny.Map({d_25_v2_: (d_31_v7_)[default__.safeIndex(469, (d_31_v7_).length(0))]})
        d_34_v10_: _dafny.Seq
        d_34_v10_ = _dafny.SeqWithoutIsStrInference([len(d_33_v9_)])
        d_35_v11_: _dafny.Seq
        d_35_v11_ = _dafny.SeqWithoutIsStrInference([(d_31_v7_)[default__.safeIndex(469, (d_31_v7_).length(0))]])
        index2_ = default__.safeIndex(533, (d_32_v8_).length(0))
        (d_32_v8_)[index2_] = (d_34_v10_)[default__.safeIndex(len(d_35_v11_), len(d_34_v10_))]
        d_36_v12_: D0
        d_36_v12_ = D0_DC2(len(d_24_v0_), (d_31_v7_)[default__.safeIndex(469, (d_31_v7_).length(0))], p0, len(d_33_v9_))
        index3_ = default__.safeIndex(533, (d_32_v8_).length(0))
        (d_32_v8_)[index3_] = (d_36_v12_).cf6
        d_37_i0_: int
        d_37_i0_ = 0
        with _dafny.label("0"):
            while p0:
                with _dafny.c_label("0"):
                    if (d_37_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_37_i0_ = (d_37_i0_) + (1)
                    index4_ = default__.safeIndex(533, (d_32_v8_).length(0))
                    (d_32_v8_)[index4_] = (0) - ((((d_32_v8_)[default__.safeIndex(533, (d_32_v8_).length(0))]) + (len(d_24_v0_))) + (d_25_v2_))
                    d_38_v13_: str
                    d_38_v13_ = _dafny.CodePoint('b')
                    index5_ = default__.safeIndex(469, (d_31_v7_).length(0))
                    (d_31_v7_)[index5_] = not(default__.fm0(p0, default__.fm5((d_32_v8_)[default__.safeIndex(533, (d_32_v8_).length(0))], globalState), d_38_v13_, d_25_v2_, globalState))
                    (globalState).f5 = d_25_v2_
                    d_39_v14_: _dafny.Array
                    nw2_ = _dafny.Array(_dafny.Map({}), 21)
                    d_39_v14_ = nw2_
                    d_40_v15_: D1
                    d_40_v15_ = D1_DC4(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_41_i1_ in range(default__.abs(845))]))
                    d_42_v16_: D1
                    d_42_v16_ = D1_DC6(d_40_v15_)
                    d_43_v17_: D1
                    d_43_v17_ = D1_DC6(d_42_v16_)
                    d_44_v18_: D1
                    d_44_v18_ = D1_DC6(d_42_v16_)
                    d_45_v19_: D1
                    d_45_v19_ = D1_DC6(d_42_v16_)
                    d_46_v20_: _dafny.Map
                    d_46_v20_ = _dafny.Map({d_45_v19_: (d_31_v7_)[default__.safeIndex(469, (d_31_v7_).length(0))]})
                    d_47_v21_: _dafny.Map
                    d_47_v21_ = _dafny.Map({d_46_v20_: default__.fm3(172, False, p0, globalState)})
                    index6_ = default__.safeIndex(216, (d_39_v14_).length(0))
                    (d_39_v14_)[index6_] = (d_47_v21_) | (d_47_v21_)
                    index7_ = default__.safeIndex(216, (d_39_v14_).length(0))
                    (d_39_v14_)[index7_] = d_47_v21_
                    pass
            pass
        d_48_v22_: str
        d_48_v22_ = _dafny.CodePoint('b')
        d_49_i2_: int
        d_49_i2_ = 0
        with _dafny.label("1"):
            while (d_24_v0_) < (((d_24_v0_) + (d_24_v0_)).set(default__.safeIndex((d_32_v8_)[default__.safeIndex(533, (d_32_v8_).length(0))], len((d_24_v0_) + (d_24_v0_))), d_48_v22_)):
                with _dafny.c_label("1"):
                    if (d_49_i2_) >= (100):
                        raise _dafny.Break("1")
                    d_49_i2_ = (d_49_i2_) + (1)
                    d_50_v24_: _dafny.Array
                    nw3_ = _dafny.Array(None, 16)
                    nw3_[int(0)] = d_33_v9_
                    nw3_[int(1)] = _dafny.Map({(d_32_v8_)[default__.safeIndex(533, (d_32_v8_).length(0))]: (d_31_v7_)[default__.safeIndex(469, (d_31_v7_).length(0))]})
                    nw3_[int(2)] = d_33_v9_
                    nw3_[int(3)] = d_33_v9_
                    nw3_[int(4)] = d_33_v9_
                    nw3_[int(5)] = d_33_v9_
                    def iife1_():
                        coll1_ = _dafny.Map()
                        compr_1_: int
                        for compr_1_ in (d_26_v3_).keys.Elements:
                            d_51_v23_: int = compr_1_
                            if (d_51_v23_) in (d_26_v3_):
                                coll1_[(d_51_v23_) - ((d_32_v8_)[default__.safeIndex(533, (d_32_v8_).length(0))])] = True
                        return _dafny.Map(coll1_)
                    nw3_[int(6)] = iife1_()
                    
                    nw3_[int(7)] = d_33_v9_
                    nw3_[int(8)] = d_33_v9_
                    nw3_[int(9)] = d_33_v9_
                    nw3_[int(10)] = d_33_v9_
                    nw3_[int(11)] = d_33_v9_
                    nw3_[int(12)] = d_33_v9_
                    nw3_[int(13)] = d_33_v9_
                    nw3_[int(14)] = d_33_v9_
                    nw3_[int(15)] = d_33_v9_
                    d_50_v24_ = nw3_
                    d_52_v25_: _dafny.Map
                    d_52_v25_ = _dafny.Map({d_50_v24_: (_dafny.SeqWithoutIsStrInference([p0, p0])).set(default__.safeIndex(len((d_34_v10_).set(default__.safeIndex(len(d_34_v10_), len(d_34_v10_)), (d_32_v8_)[default__.safeIndex(533, (d_32_v8_).length(0))])), len(_dafny.SeqWithoutIsStrInference([p0, p0]))), (d_31_v7_)[default__.safeIndex(469, (d_31_v7_).length(0))])})
                    d_52_v25_ = (d_52_v25_).set(d_50_v24_, d_35_v11_)
                    d_53_v26_: _dafny.Map
                    d_53_v26_ = _dafny.Map({(d_32_v8_)[default__.safeIndex(533, (d_32_v8_).length(0))]: d_24_v0_})
                    d_54_v27_: _dafny.Map
                    d_54_v27_ = _dafny.Map({p0: len(d_53_v26_)})
                    d_54_v27_ = d_54_v27_
                    (globalState).f0 = not(p0)
                    d_55_v28_: C0
                    nw4_ = C0()
                    nw4_.ctor__(d_25_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxfg")))
                    d_55_v28_ = nw4_
                    d_56_v29_: D1
                    d_56_v29_ = D1_DC5((d_31_v7_)[default__.safeIndex(469, (d_31_v7_).length(0))], d_55_v28_, p0)
                    (globalState).f7 = (d_56_v29_).cf14
                    pass
            pass

    @staticmethod
    def Main(noArgsParameter__):
        d_57_v0_: _dafny.Seq
        d_57_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sraobfhjd"))
        d_58_v1_: int
        d_58_v1_ = 178
        d_59_v2_: str
        d_59_v2_ = _dafny.CodePoint('q')
        d_60_v3_: _dafny.Array
        def lambda0_(d_61_v1_):
            def lambda1_(d_62_i0_):
                return _dafny.Map({(D0_DC2(d_61_v1_, True, True, len(_dafny.SeqWithoutIsStrInference([False])))).cf7: d_61_v1_})

            return lambda1_

        init0_ = lambda0_(d_58_v1_)
        nw5_ = _dafny.Array(None, 8)
        for i0_0_ in range(nw5_.length(0)):
            nw5_[i0_0_] = init0_(i0_0_)
        d_60_v3_ = nw5_
        d_63_v4_: _dafny.Array
        def lambda2_(d_64_v1_):
            def lambda3_(d_65_i1_):
                return _dafny.Set({d_64_v1_})

            return lambda3_

        init1_ = lambda2_(d_58_v1_)
        nw6_ = _dafny.Array(None, 19)
        for i0_1_ in range(nw6_.length(0)):
            nw6_[i0_1_] = init1_(i0_1_)
        d_63_v4_ = nw6_
        d_66_v5_: _dafny.Array
        def lambda4_(d_67_i2_):
            return default__.safeModuloInt(d_67_i2_, 855)

        init2_ = lambda4_
        nw7_ = _dafny.Array(None, 23)
        for i0_2_ in range(nw7_.length(0)):
            nw7_[i0_2_] = init2_(i0_2_)
        d_66_v5_ = nw7_
        d_68_v6_: _dafny.Array
        nw8_ = _dafny.Array(_dafny.Array(None, 0), 17)
        d_68_v6_ = nw8_
        d_69_v7_: bool
        d_69_v7_ = False
        d_70_v8_: _dafny.Set
        d_70_v8_ = _dafny.Set({d_59_v2_, d_59_v2_, _dafny.CodePoint('u'), d_59_v2_})
        d_71_globalState_: GlobalState
        nw9_ = GlobalState()
        nw9_.ctor__(False, 237, (d_57_v0_).set(default__.safeIndex(d_58_v1_, len(d_57_v0_)), d_59_v2_), d_60_v3_, _dafny.CodePoint('w'), -46, d_63_v4_, False, d_66_v5_, d_68_v6_, d_66_v5_, _dafny.Map({d_69_v7_: (0) - (d_58_v1_)}), False, True, d_70_v8_, True)
        d_71_globalState_ = nw9_
        d_59_v2_ = d_59_v2_
        d_72_v9_: _dafny.MultiSet
        d_72_v9_ = _dafny.MultiSet([d_69_v7_])
        hi0_ = ((d_72_v9_)[d_69_v7_] if (d_69_v7_) in (d_72_v9_) else d_58_v1_)
        for d_73_i3_ in range(d_58_v1_, hi0_):
            d_74_v10_: D0
            d_74_v10_ = D0_DC2(998, d_69_v7_, d_69_v7_, len(d_57_v0_))
            pat_let_tv0_ = d_57_v0_
            def iife2_(_pat_let0_0):
                def iife3_(d_75_dt__update__tmp_h0_):
                    def iife4_(_pat_let1_0):
                        def iife5_(d_76_dt__update_hcf9_h0_):
                            return D0_DC2((d_75_dt__update__tmp_h0_).cf6, (d_75_dt__update__tmp_h0_).cf7, (d_75_dt__update__tmp_h0_).cf8, d_76_dt__update_hcf9_h0_)
                        return iife5_(_pat_let1_0)
                    return iife4_(len(_dafny.Map({d_73_i3_: len(pat_let_tv0_)})))
                return iife3_(_pat_let0_0)
            if ((d_73_i3_) >= (d_58_v1_) if d_69_v7_ else (iife2_(d_74_v10_)).cf8):
                d_77_v11_: _dafny.Array
                def lambda5_(d_78_i4_):
                    return not(not(not(True)))

                init3_ = lambda5_
                nw10_ = _dafny.Array(None, 3)
                for i0_3_ in range(nw10_.length(0)):
                    nw10_[i0_3_] = init3_(i0_3_)
                d_77_v11_ = nw10_
                index8_ = default__.safeIndex(829, (d_77_v11_).length(0))
                (d_77_v11_)[index8_] = ((d_72_v9_).cardinality) > (d_73_i3_)
                index9_ = default__.safeIndex(829, (d_77_v11_).length(0))
                (d_77_v11_)[index9_] = not (d_69_v7_) or (d_69_v7_)
                d_79_v12_: _dafny.Seq
                d_79_v12_ = _dafny.SeqWithoutIsStrInference([d_66_v5_])
                d_80_v13_: _dafny.Map
                d_80_v13_ = _dafny.Map({(d_77_v11_)[default__.safeIndex(829, (d_77_v11_).length(0))]: d_69_v7_})
                d_81_v14_: _dafny.Array
                nw11_ = _dafny.Array(None, 5)
                nw11_[int(0)] = d_66_v5_
                nw11_[int(1)] = d_66_v5_
                nw11_[int(2)] = d_66_v5_
                nw11_[int(3)] = (d_79_v12_)[default__.safeIndex(len(d_80_v13_), len(d_79_v12_))]
                nw11_[int(4)] = d_66_v5_
                d_81_v14_ = nw11_
                index10_ = default__.safeIndex(694, (d_81_v14_).length(0))
                (d_81_v14_)[index10_] = d_66_v5_
                index11_ = default__.safeIndex(694, (d_81_v14_).length(0))
                (d_81_v14_)[index11_] = d_66_v5_
                d_82_v15_: _dafny.Map
                d_82_v15_ = _dafny.Map({d_66_v5_: not (True) or (d_69_v7_)})
                d_83_v16_: _dafny.Map
                d_83_v16_ = _dafny.Map({d_69_v7_: d_66_v5_})
                d_82_v15_ = (d_82_v15_).set(((d_83_v16_)[(d_77_v11_)[default__.safeIndex(829, (d_77_v11_).length(0))]] if ((d_77_v11_)[default__.safeIndex(829, (d_77_v11_).length(0))]) in (d_83_v16_) else d_66_v5_), (d_73_i3_) > (d_73_i3_))
                (d_71_globalState_).f5 = (d_58_v1_ if (d_77_v11_)[default__.safeIndex(829, (d_77_v11_).length(0))] else 823)
                (d_71_globalState_).f5 = (d_58_v1_) - (d_73_i3_)
            elif True:
                default__.m0(d_69_v7_, d_71_globalState_)
                (d_71_globalState_).f12 = default__.fm0((d_69_v7_) or (d_69_v7_), d_59_v2_, d_59_v2_, d_58_v1_, d_71_globalState_)
                d_59_v2_ = d_59_v2_
                d_58_v1_ = len(d_57_v0_)
                (d_71_globalState_).f0 = True
            d_84_v17_: _dafny.Map
            d_84_v17_ = _dafny.Map({True: d_69_v7_})
            d_85_v18_: _dafny.Map
            d_85_v18_ = _dafny.Map({d_84_v17_: d_57_v0_})
            (d_71_globalState_).f5 = len((((d_85_v18_)[d_84_v17_] if (d_84_v17_) in (d_85_v18_) else d_57_v0_)) + ((d_57_v0_ if d_69_v7_ else d_57_v0_)))
            d_86_v19_: C0
            nw12_ = C0()
            nw12_.ctor__(d_73_i3_, d_57_v0_)
            d_86_v19_ = nw12_
            (d_71_globalState_).f0 = (d_69_v7_ if d_69_v7_ else d_69_v7_)
        default__.m0(d_69_v7_, d_71_globalState_)
        hi1_ = d_58_v1_
        for d_87_i5_ in range((0) - ((d_58_v1_) + (d_58_v1_)), hi1_):
            rhs0_ = d_66_v5_
            rhs1_ = d_59_v2_
            rhs2_ = not (d_69_v7_) or (d_69_v7_)
            rhs3_ = d_69_v7_
            lhs0_ = d_71_globalState_
            lhs1_ = d_71_globalState_
            lhs2_ = d_71_globalState_
            lhs0_.f10 = rhs0_
            d_59_v2_ = rhs1_
            lhs1_.f12 = rhs2_
            lhs2_.f12 = rhs3_
            d_88_v20_: C0
            nw13_ = C0()
            nw13_.ctor__(d_87_i5_, (d_57_v0_) + (d_57_v0_))
            d_88_v20_ = nw13_
            d_89_v21_: _dafny.Seq
            d_89_v21_ = _dafny.SeqWithoutIsStrInference([(d_88_v20_).f16])
            d_89_v21_ = (_dafny.SeqWithoutIsStrInference([(d_88_v20_).f16 for d_90_i6_ in range(default__.abs(346))]) if d_69_v7_ else _dafny.SeqWithoutIsStrInference([d_87_i5_ for d_91_i7_ in range(default__.abs(801))]))
            d_92_v22_: D0
            d_92_v22_ = D0_DC0(d_69_v7_)
            d_93_v23_: _dafny.Map
            d_93_v23_ = _dafny.Map({d_92_v22_: (d_88_v20_).f17})
            d_94_v24_: D1
            d_94_v24_ = D1_DC4((d_88_v20_).f17)
            d_93_v23_ = (d_93_v23_).set(D0_DC0(not(d_69_v7_)), ((d_94_v24_).cf11) + (d_57_v0_))
        d_95_v25_: _dafny.MultiSet
        d_95_v25_ = _dafny.MultiSet([d_58_v1_])
        (d_71_globalState_).f7 = (default__.fm2(d_58_v1_, d_95_v25_, d_71_globalState_)) <= (_dafny.SeqWithoutIsStrInference([d_59_v2_ for d_96_i8_ in range(default__.abs(-49))]))
        d_97_v26_: _dafny.Set
        d_97_v26_ = _dafny.Set({d_69_v7_, d_69_v7_, d_69_v7_})
        d_98_v27_: _dafny.Map
        d_98_v27_ = _dafny.Map({d_97_v26_: default__.fm0(d_69_v7_, d_59_v2_, d_59_v2_, (0) - (d_58_v1_), d_71_globalState_)})
        hi2_ = 103
        for d_99_i9_ in range(len(d_98_v27_), hi2_):
            d_100_v28_: _dafny.Array
            nw14_ = _dafny.Array(None, 11)
            d_100_v28_ = nw14_
            d_101_v29_: C0
            nw15_ = C0()
            nw15_.ctor__(d_99_i9_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrcnf")))
            d_101_v29_ = nw15_
            d_102_v30_: D1
            d_102_v30_ = D1_DC5(d_69_v7_, d_101_v29_, d_69_v7_)
            pat_let_tv1_ = d_69_v7_
            pat_let_tv2_ = d_69_v7_
            index12_ = default__.safeIndex(761, (d_100_v28_).length(0))
            def iife6_(_pat_let2_0):
                def iife7_(d_103_dt__update__tmp_h1_):
                    def iife8_(_pat_let3_0):
                        def iife9_(d_104_dt__update_hcf14_h0_):
                            def iife10_(_pat_let4_0):
                                def iife11_(d_105_dt__update_hcf12_h0_):
                                    return D1_DC5(d_105_dt__update_hcf12_h0_, (d_103_dt__update__tmp_h1_).cf13, d_104_dt__update_hcf14_h0_)
                                return iife11_(_pat_let4_0)
                            return iife10_(pat_let_tv2_)
                        return iife9_(_pat_let3_0)
                    return iife8_(pat_let_tv1_)
                return iife7_(_pat_let2_0)
            (d_100_v28_)[index12_] = iife6_(d_102_v30_)
            index13_ = default__.safeIndex(761, (d_100_v28_).length(0))
            (d_100_v28_)[index13_] = d_102_v30_
            d_106_v31_: _dafny.Set
            d_106_v31_ = _dafny.Set({d_101_v29_})
            rhs4_ = (0) - ((d_101_v29_).f16)
            rhs5_ = len((d_106_v31_) | (d_106_v31_))
            lhs3_ = d_71_globalState_
            lhs4_ = d_71_globalState_
            lhs3_.f5 = rhs4_
            lhs4_.f5 = rhs5_
            d_107_v32_: D1
            d_107_v32_ = D1_DC4((d_101_v29_).f17)
            d_108_v33_: D1
            d_108_v33_ = D1_DC6(d_107_v32_)
            d_108_v33_ = D1_DC6(d_107_v32_)
            default__.m0(d_69_v7_, d_71_globalState_)
        d_109_v34_: _dafny.Map
        d_109_v34_ = _dafny.Map({False: d_69_v7_})
        d_109_v34_ = (d_109_v34_).set(False, d_69_v7_)
        d_110_v35_: C0
        nw16_ = C0()
        nw16_.ctor__(d_58_v1_, d_57_v0_)
        d_110_v35_ = nw16_
        (d_71_globalState_).f5 = len((_dafny.Map({d_58_v1_: d_110_v35_})) | (_dafny.Map({d_58_v1_: d_110_v35_})))
        d_111_v36_: _dafny.Map
        d_111_v36_ = _dafny.Map({(d_110_v35_).f16: (_dafny.SeqWithoutIsStrInference([d_59_v2_ for d_112_i10_ in range(default__.abs(816))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fal")))})
        d_111_v36_ = (d_111_v36_).set(-282, (d_110_v35_).f17)
        d_113_v37_: C0
        nw17_ = C0()
        nw17_.ctor__(default__.safeModuloInt((d_110_v35_).f16, (d_110_v35_).f16), (_dafny.SeqWithoutIsStrInference([d_59_v2_ for d_114_i11_ in range(default__.abs(41))])) + ((d_110_v35_).f17))
        d_113_v37_ = nw17_
        default__.m0(((d_113_v37_).f16) != ((d_110_v35_).f16), d_71_globalState_)
        d_115_v38_: _dafny.Map
        d_115_v38_ = _dafny.Map({d_59_v2_: len(d_57_v0_)})
        def iife12_():
            coll2_ = _dafny.Map()
            compr_2_: str
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([d_59_v2_])).Elements:
                d_116_v39_: str = compr_2_
                if (d_116_v39_) in (_dafny.SeqWithoutIsStrInference([d_59_v2_])):
                    coll2_[d_116_v39_] = len(_dafny.Set({(d_113_v37_).f16, (d_110_v35_).f16, (d_113_v37_).f16, len(d_109_v34_), d_58_v1_}))
            return _dafny.Map(coll2_)
        d_115_v38_ = (iife12_()
        ) | (d_115_v38_)
        d_69_v7_ = d_69_v7_
        (d_71_globalState_).f5 = (d_113_v37_).f16
        default__.m0(d_69_v7_, d_71_globalState_)
        d_117_v40_: _dafny.Map
        d_117_v40_ = _dafny.Map({d_69_v7_: d_110_v35_})
        (d_71_globalState_).f5 = (((d_113_v37_).f16) * (len((D2_DC7(d_117_v40_)).cf16))) + (255)
        _dafny.print((d_57_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_58_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_59_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_v3_)[0]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_v3_)[1]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_v3_)[2]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_v3_)[3]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_v3_)[4]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_v3_)[5]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_v3_)[6]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_60_v3_)[7]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[0]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[1]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[2]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[3]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[4]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[5]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[6]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[7]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[8]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[9]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[10]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[11]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[12]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[13]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[14]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[15]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[16]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[17]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_63_v4_)[18]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_66_v5_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_69_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_70_v8_) == (_dafny.Set({_dafny.CodePoint('q'), _dafny.CodePoint('u')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_71_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_71_globalState_).f2).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_71_globalState_).f3)[0]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_71_globalState_).f3)[1]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_71_globalState_).f3)[2]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_71_globalState_).f3)[3]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_71_globalState_).f3)[4]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_71_globalState_).f3)[5]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_71_globalState_).f3)[6]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_71_globalState_).f3)[7]) == (_dafny.Map({True: 178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_71_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[0]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[1]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[2]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[3]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[4]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[5]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[6]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[7]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[8]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[9]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[10]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[11]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[12]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[13]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[14]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[15]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[16]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[17]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_.f6)[18]) == (_dafny.Set({178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_71_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f8)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f10)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_globalState_).f11) == (_dafny.Map({False: -178}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_71_globalState_.f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_.f14) == (_dafny.Set({_dafny.CodePoint('q'), _dafny.CodePoint('u')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_71_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_72_v9_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_95_v25_) == (_dafny.MultiSet([178]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_v26_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v27_) == (_dafny.Map({_dafny.Set({False}): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_109_v34_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_110_v35_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_110_v35_).f17).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_111_v36_) == (_dafny.Map({178: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqfal")), -282: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sraobfhjd"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v37_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_113_v37_).f17).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_115_v38_) == (_dafny.Map({_dafny.CodePoint('q'): 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_117_v40_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, _dafny.Map({}), False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC5(False, None, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)

class D1_DC5(D1, NamedTuple('DC5', [('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({self.cf11.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(int(0), int(0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f5: int = int(0)
        self.f6: _dafny.Array = _dafny.Array(None, 0)
        self.f7: bool = False
        self.f10: _dafny.Array = _dafny.Array(None, 0)
        self.f12: bool = False
        self.f14: _dafny.Set = _dafny.Set({})
        self._f1: int = int(0)
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f4: str = _dafny.CodePoint('D')
        self._f8: _dafny.Array = _dafny.Array(None, 0)
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f11: _dafny.Map = _dafny.Map({})
        self._f13: bool = False
        self._f15: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self).f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self)._f15 = f15

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
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f11(self):
        return self._f11
    @property
    def f13(self):
        return self._f13
    @property
    def f15(self):
        return self._f15

class C0:
    def  __init__(self):
        self._f16: int = int(0)
        self._f17: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f16, f17):
        (self)._f16 = f16
        (self)._f17 = f17

    def fm1(self, p0, p1, p2, globalState):
        return len(((self).f17 if not((_dafny.MultiSet([(D0_DC1(True, _dafny.Map({True: (self).f16}), True, (self).f16, (self).f16)).cf1])).issubset(_dafny.MultiSet([False, True]))) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_118_i0_ in range(default__.abs(552))])))

    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
