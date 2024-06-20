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
    def fm0(p0, p1, globalState):
        source0_ = D8_DC23()
        if source0_.is_DC23:
            return (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_0_i0_ in range(default__.abs(7))]))) < (len(_dafny.Map({not(True): 25})))
        elif source0_.is_DC24:
            d_1___mcc_h0_ = source0_.cf45
            d_2_cf45_ = d_1___mcc_h0_
            return True
        elif source0_.is_DC22:
            d_3___mcc_h1_ = source0_.cf44
            d_4_cf44_ = d_3___mcc_h1_
            return (False) in (_dafny.SeqWithoutIsStrInference([True, False]))
        elif True:
            d_5___mcc_h2_ = source0_.cf46
            d_6_cf46_ = d_5___mcc_h2_
            return True

    @staticmethod
    def fm1(globalState):
        return D0_DC0(len(_dafny.Map({112: not(True)})), (0) - (default__.safeModuloInt(-891, 290)))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return default__.safeDivisionInt((636) - (len(_dafny.Set({-371, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qywhj"))), len(_dafny.SeqWithoutIsStrInference([616])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_7_i0_ in range(default__.abs(929))]))}))), 503)

    @staticmethod
    def fm3(p0, globalState):
        return D0_DC1()

    @staticmethod
    def fm4(p0, p1, globalState):
        if (_dafny.SeqWithoutIsStrInference([False, True, False, True, False])) == (_dafny.SeqWithoutIsStrInference([False])):
            return D1_DC2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dudmm")))
        elif True:
            return D1_DC2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([868, 482])) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])), (0) - (-324)])) + (_dafny.SeqWithoutIsStrInference([822])))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtqhqq")))])

    @staticmethod
    def fm7(p0, p1, globalState):
        source1_ = D7_DC21(D7_DC20(True, not(False), -30, True))
        if source1_.is_DC20:
            d_8___mcc_h0_ = source1_.cf39
            d_9___mcc_h1_ = source1_.cf40
            d_10___mcc_h2_ = source1_.cf41
            d_11___mcc_h3_ = source1_.cf42
            d_12_cf42_ = d_11___mcc_h3_
            d_13_cf41_ = d_10___mcc_h2_
            d_14_cf40_ = d_9___mcc_h1_
            d_15_cf39_ = d_8___mcc_h0_
            return _dafny.CodePoint('b')
        elif source1_.is_DC19:
            d_16___mcc_h4_ = source1_.cf38
            d_17_cf38_ = d_16___mcc_h4_
            return _dafny.CodePoint('w')
        elif True:
            d_18___mcc_h5_ = source1_.cf43
            d_19_cf43_ = d_18___mcc_h5_
            return _dafny.CodePoint('k')

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return _dafny.Set({_dafny.SeqWithoutIsStrInference([702])})

    @staticmethod
    def fm9(p0, globalState):
        return ((_dafny.Set({False, True})).intersection(_dafny.Set({not(not(not(True)))}))) | ((_dafny.Set({False, False})) | (_dafny.Set({True, True})))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        source2_ = D1_DC6(D1_DC3(len(_dafny.SeqWithoutIsStrInference([True, not(True)])), -624))
        if source2_.is_DC3:
            d_20___mcc_h0_ = source2_.cf3
            d_21___mcc_h1_ = source2_.cf4
            d_22_cf4_ = d_21___mcc_h1_
            d_23_cf3_ = d_20___mcc_h0_
            if not(False):
                return D1_DC6(D1_DC5(d_22_cf4_, True, d_22_cf4_, True))
            elif True:
                return D1_DC6(D1_DC3(d_22_cf4_, d_23_cf3_))
        elif source2_.is_DC4:
            d_24___mcc_h2_ = source2_.cf5
            d_25___mcc_h3_ = source2_.cf6
            d_26_cf6_ = d_25___mcc_h3_
            d_27_cf5_ = d_24___mcc_h2_
            return D1_DC6(D1_DC6(D1_DC5(d_26_cf6_, d_27_cf5_, d_26_cf6_, d_27_cf5_)))
        elif source2_.is_DC5:
            d_28___mcc_h4_ = source2_.cf7
            d_29___mcc_h5_ = source2_.cf8
            d_30___mcc_h6_ = source2_.cf9
            d_31___mcc_h7_ = source2_.cf10
            d_32_cf10_ = d_31___mcc_h7_
            d_33_cf9_ = d_30___mcc_h6_
            d_34_cf8_ = d_29___mcc_h5_
            d_35_cf7_ = d_28___mcc_h4_
            return D1_DC6(D1_DC6(D1_DC6(D1_DC5((0) - (d_35_cf7_), d_32_cf10_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xexlenl"))), d_34_cf8_))))
        elif source2_.is_DC2:
            d_36___mcc_h8_ = source2_.cf2
            d_37_cf2_ = d_36___mcc_h8_
            return D1_DC6(D1_DC6(D1_DC2(d_37_cf2_)))
        elif True:
            d_38___mcc_h9_ = source2_.cf11
            d_39_cf11_ = d_38___mcc_h9_
            if False:
                return D1_DC6(D1_DC2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bno"))))
            elif True:
                return D1_DC6(D1_DC4(False, 755))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return D1_DC4((True if True else False), len((_dafny.SeqWithoutIsStrInference([not(not(True))])) + (_dafny.SeqWithoutIsStrInference([True, False, not(False), False]))))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dm")) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yx")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "badgldb")))

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        return (_dafny.Set({22, -190, -212})) | ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i')])), -603]))})) - (_dafny.Set({(0) - (len(_dafny.Map({297: True}))), (0) - (-80), 377, -598})))

    @staticmethod
    def fm14(p0, globalState):
        source3_ = D8_DC23()
        if source3_.is_DC23:
            return _dafny.MultiSet([True])
        elif source3_.is_DC24:
            d_40___mcc_h0_ = source3_.cf45
            d_41_cf45_ = d_40___mcc_h0_
            return _dafny.MultiSet([d_41_cf45_, False, d_41_cf45_, d_41_cf45_])
        elif source3_.is_DC22:
            d_42___mcc_h1_ = source3_.cf44
            d_43_cf44_ = d_42___mcc_h1_
            return _dafny.MultiSet([(D1_DC4(True, 371)).cf5, False])
        elif True:
            d_44___mcc_h2_ = source3_.cf46
            d_45_cf46_ = d_44___mcc_h2_
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))) - (_dafny.MultiSet([False]))

    @staticmethod
    def m0(p0, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        d_46_v0_: str
        d_46_v0_ = _dafny.CodePoint('r')
        d_47_v1_: D2
        d_47_v1_ = D2_DC10(-185, d_46_v0_)
        source4_ = d_47_v1_
        if source4_.is_DC8:
            d_48___mcc_h0_ = source4_.cf13
            d_49_cf13_ = d_48___mcc_h0_
            d_50_v2_: bool
            d_50_v2_ = True
            d_51_v3_: _dafny.Seq
            d_51_v3_ = _dafny.SeqWithoutIsStrInference([d_50_v2_])
            d_52_v4_: C0
            nw0_ = C0()
            nw0_.ctor__((d_51_v3_) < (d_51_v3_))
            d_52_v4_ = nw0_
            d_52_v4_ = d_52_v4_
            d_53_v5_: _dafny.Seq
            d_53_v5_ = _dafny.SeqWithoutIsStrInference([(d_51_v3_) + (_dafny.SeqWithoutIsStrInference([d_50_v2_]))])
            d_54_v6_: int
            d_54_v6_ = 978
            d_53_v5_ = ((d_53_v5_).set(default__.safeIndex(default__.fm2(_dafny.MultiSet([d_46_v0_, _dafny.CodePoint('e')]), d_54_v6_, d_54_v6_, False, globalState), len(d_53_v5_)), _dafny.SeqWithoutIsStrInference([(d_52_v4_).f29, d_50_v2_]))) + (d_53_v5_)
            d_52_v4_ = d_52_v4_
            (globalState).f27 = d_54_v6_
        elif source4_.is_DC9:
            d_55___mcc_h1_ = source4_.cf14
            d_56___mcc_h2_ = source4_.cf15
            d_57_cf15_ = d_56___mcc_h2_
            d_58_cf14_ = d_55___mcc_h1_
            d_59_v8_: _dafny.Array
            def lambda0_(d_60_cf15_, d_61_cf14_):
                def lambda1_(d_62_i0_):
                    def iife0_():
                        coll0_ = _dafny.Map()
                        compr_0_: int
                        for compr_0_ in (_dafny.Map({660: True})).keys.Elements:
                            d_63_v7_: int = compr_0_
                            if (d_63_v7_) in (_dafny.Map({660: True})):
                                coll0_[default__.safeDivisionInt(d_63_v7_, len(_dafny.Set({len(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([284, -907]))]): (d_60_cf15_).f29}))})))] = d_61_cf14_
                        return _dafny.Map(coll0_)
                    return (d_62_i0_) - (len(iife0_()
                    ))

                return lambda1_

            init0_ = lambda0_(d_57_cf15_, d_58_cf14_)
            nw1_ = _dafny.Array(None, 19)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_59_v8_ = nw1_
            d_64_v9_: D4
            d_64_v9_ = D4_DC12(d_59_v8_)
            (globalState).f28 = (d_64_v9_).cf19
            if (d_57_cf15_).f29:
                d_65_v10_: int
                d_65_v10_ = 581
                d_66_v11_: _dafny.Seq
                d_66_v11_ = _dafny.SeqWithoutIsStrInference([d_65_v10_])
                d_67_v12_: _dafny.Array
                nw2_ = _dafny.Array(False, 5)
                d_67_v12_ = nw2_
                d_68_v13_: _dafny.Map
                d_68_v13_ = _dafny.Map({d_67_v12_: d_57_cf15_})
                d_57_cf15_ = (d_57_cf15_ if not((d_65_v10_) >= ((d_66_v11_)[default__.safeIndex(len(d_68_v13_), len(d_66_v11_))])) else d_57_cf15_)
                (globalState).f13 = (d_66_v11_)[default__.safeIndex(d_65_v10_, len(d_66_v11_))]
                d_69_v14_: _dafny.Array
                nw3_ = _dafny.Array(D2.default()(), 5)
                d_69_v14_ = nw3_
                index0_ = default__.safeIndex(12, (d_69_v14_).length(0))
                (d_69_v14_)[index0_] = d_47_v1_
                index1_ = default__.safeIndex(12, (d_69_v14_).length(0))
                (d_69_v14_)[index1_] = d_47_v1_
                d_70_v15_: _dafny.Map
                d_70_v15_ = _dafny.Map({(d_57_cf15_).f29: (d_57_cf15_).f29})
                d_70_v15_ = (d_70_v15_).set(default__.fm0(p0, default__.fm0(p0, False, globalState), globalState), d_58_cf14_)
                (globalState).f13 = 249
            elif True:
                d_71_v16_: C0
                nw4_ = C0()
                nw4_.ctor__(not((d_57_cf15_).f29))
                d_71_v16_ = nw4_
                (globalState).f14 = (True if False else True)
                d_72_v17_: _dafny.Set
                d_72_v17_ = _dafny.Set({d_58_cf14_, (d_57_cf15_).f29})
                (globalState).f13 = len(d_72_v17_)
                (globalState).f14 = (d_57_cf15_).f29
                d_73_v18_: int
                d_73_v18_ = 927
                d_74_v19_: _dafny.Seq
                d_74_v19_ = _dafny.SeqWithoutIsStrInference([d_73_v18_, d_73_v18_, d_73_v18_])
                d_75_v20_: _dafny.Set
                d_75_v20_ = _dafny.Set({d_74_v19_, d_74_v19_})
                (globalState).f14 = (len(d_75_v20_)) == (d_73_v18_)
            d_76_v21_: int
            d_76_v21_ = 82
            d_77_v22_: _dafny.Map
            d_77_v22_ = _dafny.Map({d_76_v21_: ((0) - (d_76_v21_) if (d_57_cf15_).f29 else d_76_v21_)})
            d_78_v23_: _dafny.Seq
            d_78_v23_ = _dafny.SeqWithoutIsStrInference([d_76_v21_])
            d_77_v22_ = (d_77_v22_).set(d_76_v21_, (d_76_v21_) - (len(d_78_v23_)))
            d_79_v24_: C0
            nw5_ = C0()
            nw5_.ctor__(d_58_cf14_)
            d_79_v24_ = nw5_
        elif source4_.is_DC10:
            d_80___mcc_h3_ = source4_.cf16
            d_81___mcc_h4_ = source4_.cf17
            d_82_cf17_ = d_81___mcc_h4_
            d_83_cf16_ = d_80___mcc_h3_
            d_84_v25_: bool
            d_84_v25_ = False
            d_85_v26_: _dafny.Array
            def lambda2_(d_86_v25_):
                def lambda3_(d_87_i1_):
                    return d_86_v25_

                return lambda3_

            init1_ = lambda2_(d_84_v25_)
            nw6_ = _dafny.Array(None, 13)
            for i0_1_ in range(nw6_.length(0)):
                nw6_[i0_1_] = init1_(i0_1_)
            d_85_v26_ = nw6_
            d_88_v27_: _dafny.Seq
            d_88_v27_ = _dafny.SeqWithoutIsStrInference([d_84_v25_, d_84_v25_])
            d_89_v28_: _dafny.Seq
            d_89_v28_ = _dafny.SeqWithoutIsStrInference([d_88_v27_])
            d_90_v29_: _dafny.Map
            d_90_v29_ = _dafny.Map({d_89_v28_: d_85_v26_})
            d_91_v30_: D5
            d_91_v30_ = D5_DC15(p0, d_84_v25_, d_84_v25_, d_84_v25_, d_85_v26_)
            d_92_v31_: _dafny.Array
            nw7_ = _dafny.Array(None, 22)
            nw7_[int(0)] = (d_85_v26_ if d_84_v25_ else d_85_v26_)
            nw7_[int(1)] = d_85_v26_
            nw7_[int(2)] = d_85_v26_
            nw7_[int(3)] = d_85_v26_
            nw7_[int(4)] = d_85_v26_
            nw7_[int(5)] = d_85_v26_
            nw7_[int(6)] = d_85_v26_
            nw7_[int(7)] = d_85_v26_
            nw7_[int(8)] = d_85_v26_
            nw7_[int(9)] = d_85_v26_
            nw7_[int(10)] = d_85_v26_
            nw7_[int(11)] = d_85_v26_
            nw7_[int(12)] = d_85_v26_
            nw7_[int(13)] = d_85_v26_
            nw7_[int(14)] = d_85_v26_
            nw7_[int(15)] = ((d_90_v29_)[d_89_v28_] if (d_89_v28_) in (d_90_v29_) else d_85_v26_)
            nw7_[int(16)] = d_85_v26_
            nw7_[int(17)] = d_85_v26_
            nw7_[int(18)] = (d_91_v30_).cf29
            nw7_[int(19)] = d_85_v26_
            nw7_[int(20)] = d_85_v26_
            nw7_[int(21)] = d_85_v26_
            d_92_v31_ = nw7_
            d_93_v32_: D0
            d_93_v32_ = D0_DC0(d_83_cf16_, d_83_cf16_)
            def iife1_(_pat_let0_0):
                def iife2_(d_94_dt__update__tmp_h0_):
                    def iife3_(_pat_let1_0):
                        def iife4_(d_95_dt__update_hcf0_h0_):
                            return D0_DC0(d_95_dt__update_hcf0_h0_, (d_94_dt__update__tmp_h0_).cf1)
                        return iife4_(_pat_let1_0)
                    return iife3_(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfraklcw"))))
                return iife2_(_pat_let0_0)
            d_92_v31_ = (d_92_v31_ if (-176) < ((iife1_(d_93_v32_)).cf0) else d_92_v31_)
            (globalState).f14 = d_84_v25_
            d_96_v33_: _dafny.Seq
            d_96_v33_ = _dafny.SeqWithoutIsStrInference([d_83_cf16_])
            (globalState).f27 = (d_96_v33_)[default__.safeIndex(d_83_cf16_, len(d_96_v33_))]
            d_97_v34_: C0
            nw8_ = C0()
            nw8_.ctor__(d_84_v25_)
            d_97_v34_ = nw8_
            d_98_v35_: _dafny.Map
            d_98_v35_ = _dafny.Map({(d_83_cf16_) != (d_83_cf16_): d_97_v34_})
            rhs0_ = (d_98_v35_).set((d_97_v34_).f29, d_97_v34_)
            rhs1_ = default__.fm2(_dafny.MultiSet([_dafny.CodePoint('y'), d_82_cf17_, d_82_cf17_, default__.fm7(p0, (0) - ((0) - (d_83_cf16_)), globalState)]), (d_83_cf16_ if (d_97_v34_).f29 else d_83_cf16_), (0) - (d_83_cf16_), True, globalState)
            rhs2_ = (d_97_v34_).f29
            lhs0_ = globalState
            d_98_v35_ = rhs0_
            d_83_cf16_ = rhs1_
            lhs0_.f14 = rhs2_
        elif True:
            d_99___mcc_h5_ = source4_.cf12
            d_100_cf12_ = d_99___mcc_h5_
            d_101_v36_: bool
            d_101_v36_ = True
            d_102_v37_: C0
            nw9_ = C0()
            nw9_.ctor__((d_101_v36_ if d_101_v36_ else d_101_v36_))
            d_102_v37_ = nw9_
            d_103_v38_: _dafny.MultiSet
            d_103_v38_ = _dafny.MultiSet([(d_102_v37_).f29])
            d_104_v39_: C0
            nw10_ = C0()
            nw10_.ctor__((d_101_v36_) in (d_103_v38_))
            d_104_v39_ = nw10_
            d_105_v40_: C0
            nw11_ = C0()
            nw11_.ctor__((d_102_v37_).f29)
            d_105_v40_ = nw11_
            d_106_v41_: int
            d_106_v41_ = 556
            d_107_v42_: D1
            d_107_v42_ = D1_DC4((d_105_v40_).f29, d_106_v41_)
            d_101_v36_ = (d_107_v42_).cf5
        d_108_v43_: bool
        d_108_v43_ = True
        d_109_i2_: int
        d_109_i2_ = 0
        with _dafny.label("0"):
            while not(d_108_v43_):
                with _dafny.c_label("0"):
                    if (d_109_i2_) >= (100):
                        raise _dafny.Break("0")
                    d_109_i2_ = (d_109_i2_) + (1)
                    (globalState).f14 = d_108_v43_
                    (globalState).f14 = ((d_108_v43_ if d_108_v43_ else d_108_v43_) if not(d_108_v43_) else d_108_v43_)
                    d_110_v44_: C0
                    nw12_ = C0()
                    nw12_.ctor__(False)
                    d_110_v44_ = nw12_
                    d_111_v45_: _dafny.MultiSet
                    d_111_v45_ = _dafny.MultiSet([False])
                    d_112_v46_: _dafny.Seq
                    d_112_v46_ = _dafny.SeqWithoutIsStrInference([(d_110_v44_).f29])
                    d_113_v47_: _dafny.Array
                    nw13_ = _dafny.Array(None, 25)
                    nw13_[int(0)] = d_111_v45_
                    nw13_[int(1)] = d_111_v45_
                    nw13_[int(2)] = d_111_v45_
                    nw13_[int(3)] = _dafny.MultiSet([(d_110_v44_).f29])
                    nw13_[int(4)] = d_111_v45_
                    nw13_[int(5)] = d_111_v45_
                    nw13_[int(6)] = d_111_v45_
                    nw13_[int(7)] = d_111_v45_
                    nw13_[int(8)] = (D7_DC19(_dafny.MultiSet(d_112_v46_))).cf38
                    nw13_[int(9)] = _dafny.MultiSet([default__.fm0(p0, d_108_v43_, globalState)])
                    nw13_[int(10)] = (d_111_v45_).set((d_110_v44_).f29, default__.abs(518))
                    nw13_[int(11)] = _dafny.MultiSet(d_112_v46_)
                    nw13_[int(12)] = d_111_v45_
                    nw13_[int(13)] = d_111_v45_
                    nw13_[int(14)] = d_111_v45_
                    nw13_[int(15)] = d_111_v45_
                    nw13_[int(16)] = _dafny.MultiSet([not(d_108_v43_), (d_110_v44_).f29, not(d_108_v43_), d_108_v43_])
                    nw13_[int(17)] = d_111_v45_
                    nw13_[int(18)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))
                    nw13_[int(19)] = _dafny.MultiSet([(d_110_v44_).f29, d_108_v43_])
                    nw13_[int(20)] = d_111_v45_
                    nw13_[int(21)] = d_111_v45_
                    nw13_[int(22)] = d_111_v45_
                    nw13_[int(23)] = d_111_v45_
                    nw13_[int(24)] = d_111_v45_
                    d_113_v47_ = nw13_
                    d_114_v48_: int
                    d_114_v48_ = 2
                    d_115_v49_: _dafny.Seq
                    d_115_v49_ = _dafny.SeqWithoutIsStrInference([d_114_v48_, -565, d_114_v48_])
                    d_116_v50_: _dafny.Seq
                    d_116_v50_ = _dafny.SeqWithoutIsStrInference([d_115_v49_])
                    d_117_v51_: _dafny.Map
                    d_117_v51_ = _dafny.Map({d_113_v47_: (d_116_v50_)[default__.safeIndex(d_114_v48_, len(d_116_v50_))]})
                    d_117_v51_ = (d_117_v51_).set(d_113_v47_, d_115_v49_)
                    pass
            pass
        d_118_v52_: int
        d_118_v52_ = 366
        d_119_v53_: _dafny.Set
        d_119_v53_ = _dafny.Set({d_118_v52_})
        d_120_v54_: _dafny.Set
        d_120_v54_ = _dafny.Set({len(d_119_v53_)})
        (globalState).f14 = ((d_119_v53_) == (d_119_v53_)) == (d_108_v43_)
        d_121_v55_: _dafny.Map
        d_121_v55_ = _dafny.Map({d_108_v43_: d_46_v0_})
        d_121_v55_ = (d_121_v55_).set(not(d_108_v43_), d_46_v0_)
        d_122_i3_: int
        d_122_i3_ = 0
        with _dafny.label("1"):
            while d_108_v43_:
                with _dafny.c_label("1"):
                    if (d_122_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_122_i3_ = (d_122_i3_) + (1)
                    if d_108_v43_:
                        d_123_v56_: _dafny.Map
                        d_123_v56_ = _dafny.Map({d_108_v43_: len(default__.fm9(d_118_v52_, globalState))})
                        d_124_v57_: C0
                        nw14_ = C0()
                        nw14_.ctor__(d_108_v43_)
                        d_124_v57_ = nw14_
                        d_125_v58_: _dafny.Set
                        d_125_v58_ = _dafny.Set({d_124_v57_, d_124_v57_, d_124_v57_})
                        d_123_v56_ = (d_123_v56_).set((False if d_108_v43_ else d_108_v43_), len(d_125_v58_))
                        (globalState).f14 = (d_124_v57_).f29
                        d_126_v59_: _dafny.Set
                        d_126_v59_ = _dafny.Set({(d_124_v57_).f29, d_108_v43_})
                        d_127_v60_: _dafny.Seq
                        d_127_v60_ = _dafny.SeqWithoutIsStrInference([d_118_v52_])
                        d_128_v61_: C0
                        nw15_ = C0()
                        nw15_.ctor__((_dafny.SeqWithoutIsStrInference([len(default__.fm13((d_124_v57_).f29, d_126_v59_, d_127_v60_, d_118_v52_, globalState)), d_118_v52_])) < (d_127_v60_))
                        d_128_v61_ = nw15_
                        d_129_v62_: _dafny.Seq
                        d_129_v62_ = _dafny.SeqWithoutIsStrInference([d_108_v43_, (d_128_v61_).f29])
                        d_130_v63_: D4
                        d_130_v63_ = D4_DC13(d_128_v61_, _dafny.CodePoint('w'), d_46_v0_, d_129_v62_)
                        d_131_v64_: _dafny.Map
                        d_131_v64_ = _dafny.Map({d_118_v52_: d_128_v61_})
                        d_132_v65_: _dafny.Array
                        nw16_ = _dafny.Array(None, 8)
                        nw16_[int(0)] = (d_130_v63_).cf20
                        nw16_[int(1)] = d_128_v61_
                        nw16_[int(2)] = d_124_v57_
                        nw16_[int(3)] = d_128_v61_
                        nw16_[int(4)] = d_124_v57_
                        nw16_[int(5)] = d_124_v57_
                        nw16_[int(6)] = d_124_v57_
                        nw16_[int(7)] = (d_124_v57_ if (d_128_v61_).f29 else ((d_131_v64_)[d_118_v52_] if (d_118_v52_) in (d_131_v64_) else d_128_v61_))
                        d_132_v65_ = nw16_
                        index2_ = default__.safeIndex(153, (d_132_v65_).length(0))
                        (d_132_v65_)[index2_] = d_124_v57_
                        index3_ = default__.safeIndex(153, (d_132_v65_).length(0))
                        (d_132_v65_)[index3_] = d_124_v57_
                        r2 = (len(_dafny.SeqWithoutIsStrInference([d_46_v0_ for d_133_i4_ in range(default__.abs(932))]))) * (d_118_v52_)
                    elif True:
                        d_134_v66_: _dafny.Array
                        def lambda4_(d_135_v0_, d_136_v1_):
                            def lambda5_(d_137_i5_):
                                return (_dafny.SeqWithoutIsStrInference([d_135_v0_ for d_138_i6_ in range(default__.abs(383))])).set(default__.safeIndex((d_136_v1_).cf16, len(_dafny.SeqWithoutIsStrInference([d_135_v0_ for d_139_i6_ in range(default__.abs(383))]))), d_135_v0_)

                            return lambda5_

                        init2_ = lambda4_(d_46_v0_, d_47_v1_)
                        nw17_ = _dafny.Array(None, 12)
                        for i0_2_ in range(nw17_.length(0)):
                            nw17_[i0_2_] = init2_(i0_2_)
                        d_134_v66_ = nw17_
                        index4_ = default__.safeIndex(8, (d_134_v66_).length(0))
                        (d_134_v66_)[index4_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txf"))
                        index5_ = default__.safeIndex(8, (d_134_v66_).length(0))
                        (d_134_v66_)[index5_] = p0
                        d_140_v67_: _dafny.MultiSet
                        d_140_v67_ = _dafny.MultiSet([d_46_v0_, (p0)[default__.safeIndex((0) - (d_118_v52_), len(p0))]])
                        (globalState).f25 = default__.fm2(d_140_v67_, d_118_v52_, d_118_v52_, False, globalState)
                        d_141_v68_: C0
                        nw18_ = C0()
                        nw18_.ctor__(d_108_v43_)
                        d_141_v68_ = nw18_
                        rhs3_ = d_141_v68_
                        rhs4_ = d_108_v43_
                        d_141_v68_ = rhs3_
                        d_108_v43_ = rhs4_
                        d_142_v69_: _dafny.Array
                        nw19_ = _dafny.Array(int(0), 15)
                        d_142_v69_ = nw19_
                        d_143_v70_: _dafny.Map
                        d_143_v70_ = _dafny.Map({d_142_v69_: d_141_v68_})
                        d_144_v71_: _dafny.Seq
                        d_144_v71_ = _dafny.SeqWithoutIsStrInference([d_141_v68_, ((d_143_v70_)[d_142_v69_] if (d_142_v69_) in (d_143_v70_) else d_141_v68_)])
                        d_145_v72_: _dafny.Map
                        d_145_v72_ = _dafny.Map({d_118_v52_: d_141_v68_})
                        d_146_v73_: D8
                        d_146_v73_ = D8_DC22(d_144_v71_)
                        d_147_v74_: _dafny.Map
                        d_147_v74_ = _dafny.Map({d_108_v43_: d_144_v71_})
                        d_148_v75_: _dafny.Seq
                        d_148_v75_ = _dafny.SeqWithoutIsStrInference([not(not((d_141_v68_).f29)), True])
                        d_149_v76_: _dafny.Array
                        nw20_ = _dafny.Array(None, 26)
                        nw20_[int(0)] = d_144_v71_
                        nw20_[int(1)] = _dafny.SeqWithoutIsStrInference([d_141_v68_, d_141_v68_])
                        nw20_[int(2)] = d_144_v71_
                        nw20_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_141_v68_])) + ((d_144_v71_).set(default__.safeIndex(d_118_v52_, len(d_144_v71_)), d_141_v68_))
                        nw20_[int(4)] = _dafny.SeqWithoutIsStrInference([d_141_v68_, d_141_v68_, d_141_v68_])
                        nw20_[int(5)] = _dafny.SeqWithoutIsStrInference([d_141_v68_, d_141_v68_, ((d_145_v72_)[d_118_v52_] if (d_118_v52_) in (d_145_v72_) else d_141_v68_), d_141_v68_, d_141_v68_])
                        nw20_[int(6)] = (d_144_v71_) + (d_144_v71_)
                        nw20_[int(7)] = (d_144_v71_) + ((d_146_v73_).cf44)
                        nw20_[int(8)] = _dafny.SeqWithoutIsStrInference([d_141_v68_])
                        nw20_[int(9)] = d_144_v71_
                        nw20_[int(10)] = d_144_v71_
                        nw20_[int(11)] = (d_144_v71_ if (d_141_v68_).f29 else (d_144_v71_).set(default__.safeIndex(d_118_v52_, len(d_144_v71_)), d_141_v68_))
                        nw20_[int(12)] = d_144_v71_
                        nw20_[int(13)] = d_144_v71_
                        nw20_[int(14)] = (d_144_v71_) + (d_144_v71_)
                        nw20_[int(15)] = d_144_v71_
                        nw20_[int(16)] = d_144_v71_
                        nw20_[int(17)] = d_144_v71_
                        nw20_[int(18)] = ((d_147_v74_)[d_108_v43_] if (d_108_v43_) in (d_147_v74_) else d_144_v71_)
                        nw20_[int(19)] = (d_144_v71_ if (d_141_v68_).f29 else d_144_v71_)
                        nw20_[int(20)] = d_144_v71_
                        nw20_[int(21)] = d_144_v71_
                        nw20_[int(22)] = d_144_v71_
                        nw20_[int(23)] = (d_144_v71_) + (d_144_v71_)
                        nw20_[int(24)] = ((d_144_v71_).set(default__.safeIndex(len((d_134_v66_)[default__.safeIndex(8, (d_134_v66_).length(0))]), len(d_144_v71_)), (d_144_v71_)[default__.safeIndex(len(d_148_v75_), len(d_144_v71_))])) + ((_dafny.SeqWithoutIsStrInference([d_141_v68_])).set(default__.safeIndex(d_118_v52_, len(_dafny.SeqWithoutIsStrInference([d_141_v68_]))), d_141_v68_))
                        nw20_[int(25)] = d_144_v71_
                        d_149_v76_ = nw20_
                        index6_ = default__.safeIndex(266, (d_149_v76_).length(0))
                        (d_149_v76_)[index6_] = (d_144_v71_) + (d_144_v71_)
                        d_150_v77_: _dafny.Map
                        d_150_v77_ = _dafny.Map({d_141_v68_: d_108_v43_})
                        d_151_v78_: _dafny.Set
                        d_151_v78_ = _dafny.Set({not((d_141_v68_).f29), ((d_150_v77_)[d_141_v68_] if (d_141_v68_) in (d_150_v77_) else not(False)), (d_141_v68_).f29, (d_141_v68_).f29, (d_141_v68_).f29})
                        index7_ = default__.safeIndex(266, (d_149_v76_).length(0))
                        rhs5_ = d_144_v71_
                        rhs6_ = ((d_141_v68_).f29) and (not((d_118_v52_) != (len(d_151_v78_))))
                        rhs7_ = False
                        lhs1_ = d_149_v76_
                        lhs2_ = default__.safeIndex(266, (d_149_v76_).length(0))
                        lhs3_ = globalState
                        lhs1_[lhs2_] = rhs5_
                        lhs3_.f14 = rhs6_
                        d_108_v43_ = rhs7_
                        d_141_v68_ = d_141_v68_
                    d_152_v79_: C0
                    nw21_ = C0()
                    nw21_.ctor__(d_108_v43_)
                    d_152_v79_ = nw21_
                    d_153_v80_: _dafny.Seq
                    d_153_v80_ = _dafny.SeqWithoutIsStrInference([d_152_v79_])
                    d_154_v81_: D8
                    d_154_v81_ = D8_DC22(d_153_v80_)
                    d_155_v82_: _dafny.Array
                    nw22_ = _dafny.Array(False, 5)
                    d_155_v82_ = nw22_
                    d_156_v83_: D5
                    d_156_v83_ = D5_DC15(_dafny.SeqWithoutIsStrInference([d_46_v0_ for d_157_i7_ in range(default__.abs(418))]), d_108_v43_, False, (d_152_v79_).f29, d_155_v82_)
                    d_158_v84_: _dafny.Map
                    d_158_v84_ = _dafny.Map({d_155_v82_: d_156_v83_})
                    rhs8_ = d_46_v0_
                    rhs9_ = p0
                    rhs10_ = d_118_v52_
                    rhs11_ = d_154_v81_
                    rhs12_ = len((d_158_v84_) | (d_158_v84_))
                    lhs4_ = globalState
                    lhs5_ = globalState
                    lhs4_.f2 = rhs8_
                    r0 = rhs9_
                    d_118_v52_ = rhs10_
                    d_154_v81_ = rhs11_
                    lhs5_.f17 = rhs12_
                    d_159_v85_: _dafny.MultiSet
                    d_159_v85_ = _dafny.MultiSet([d_108_v43_])
                    d_160_v86_: _dafny.Map
                    d_160_v86_ = _dafny.Map({(d_152_v79_).f29: d_118_v52_})
                    d_161_v87_: _dafny.Map
                    d_161_v87_ = _dafny.Map({(0) - (d_118_v52_): d_118_v52_})
                    d_162_v88_: _dafny.Seq
                    d_162_v88_ = _dafny.SeqWithoutIsStrInference([d_118_v52_])
                    d_163_v89_: _dafny.Seq
                    d_163_v89_ = _dafny.SeqWithoutIsStrInference([d_108_v43_, False])
                    d_164_v90_: _dafny.Array
                    nw23_ = _dafny.Array(None, 11)
                    nw23_[int(0)] = len(d_160_v86_)
                    nw23_[int(1)] = (0) - (d_118_v52_)
                    nw23_[int(2)] = d_118_v52_
                    nw23_[int(3)] = d_118_v52_
                    nw23_[int(4)] = len(d_161_v87_)
                    nw23_[int(5)] = d_118_v52_
                    nw23_[int(6)] = len(d_162_v88_)
                    nw23_[int(7)] = len(d_163_v89_)
                    nw23_[int(8)] = d_118_v52_
                    nw23_[int(9)] = d_118_v52_
                    nw23_[int(10)] = (0) - (d_118_v52_)
                    d_164_v90_ = nw23_
                    d_165_v91_: _dafny.Map
                    d_165_v91_ = _dafny.Map({(d_159_v85_).cardinality: d_164_v90_})
                    d_165_v91_ = (d_165_v91_).set((d_118_v52_) * (941), d_164_v90_)
                    d_166_v92_: C0
                    nw24_ = C0()
                    nw24_.ctor__((d_152_v79_).f29)
                    d_166_v92_ = nw24_
                    pass
            pass
        hi0_ = default__.safeModuloInt(len(d_120_v54_), d_118_v52_)
        for d_167_i8_ in range(d_118_v52_, hi0_):
            if d_108_v43_:
                d_168_v93_: C0
                nw25_ = C0()
                nw25_.ctor__(default__.fm0(p0, d_108_v43_, globalState))
                d_168_v93_ = nw25_
                d_169_v94_: _dafny.Seq
                d_169_v94_ = _dafny.SeqWithoutIsStrInference([d_168_v93_, d_168_v93_, d_168_v93_])
                d_170_v95_: _dafny.Array
                nw26_ = _dafny.Array(None, 15)
                nw26_[int(0)] = d_168_v93_
                nw26_[int(1)] = d_168_v93_
                nw26_[int(2)] = d_168_v93_
                nw26_[int(3)] = d_168_v93_
                nw26_[int(4)] = d_168_v93_
                nw26_[int(5)] = d_168_v93_
                nw26_[int(6)] = d_168_v93_
                nw26_[int(7)] = d_168_v93_
                nw26_[int(8)] = d_168_v93_
                nw26_[int(9)] = (d_169_v94_)[default__.safeIndex(d_118_v52_, len(d_169_v94_))]
                nw26_[int(10)] = d_168_v93_
                nw26_[int(11)] = d_168_v93_
                nw26_[int(12)] = d_168_v93_
                nw26_[int(13)] = d_168_v93_
                nw26_[int(14)] = d_168_v93_
                d_170_v95_ = nw26_
                index8_ = default__.safeIndex(874, (d_170_v95_).length(0))
                (d_170_v95_)[index8_] = d_168_v93_
                index9_ = default__.safeIndex(874, (d_170_v95_).length(0))
                (d_170_v95_)[index9_] = d_168_v93_
                d_171_v96_: _dafny.Seq
                d_171_v96_ = _dafny.SeqWithoutIsStrInference([d_108_v43_, default__.fm0(p0, d_108_v43_, globalState)])
                d_172_v97_: _dafny.MultiSet
                d_172_v97_ = _dafny.MultiSet([d_118_v52_])
                d_173_v98_: _dafny.Map
                d_173_v98_ = _dafny.Map({d_118_v52_: (d_168_v93_).f29})
                d_174_v99_: _dafny.Set
                d_174_v99_ = _dafny.Set({((d_173_v98_)[d_167_i8_] if (d_167_i8_) in (d_173_v98_) else not(True))})
                d_175_v100_: _dafny.Array
                nw27_ = _dafny.Array(None, 25)
                nw27_[int(0)] = d_108_v43_
                nw27_[int(1)] = d_108_v43_
                nw27_[int(2)] = d_108_v43_
                nw27_[int(3)] = False
                nw27_[int(4)] = d_108_v43_
                nw27_[int(5)] = d_108_v43_
                nw27_[int(6)] = d_108_v43_
                nw27_[int(7)] = (d_168_v93_).f29
                nw27_[int(8)] = True
                nw27_[int(9)] = (d_168_v93_).f29
                nw27_[int(10)] = (d_168_v93_).f29
                nw27_[int(11)] = d_108_v43_
                nw27_[int(12)] = (d_168_v93_).f29
                nw27_[int(13)] = (d_168_v93_).f29
                nw27_[int(14)] = d_108_v43_
                nw27_[int(15)] = d_108_v43_
                nw27_[int(16)] = (d_168_v93_).f29
                nw27_[int(17)] = (d_168_v93_).f29
                nw27_[int(18)] = not(d_108_v43_)
                nw27_[int(19)] = True
                nw27_[int(20)] = (d_168_v93_).f29
                nw27_[int(21)] = True
                nw27_[int(22)] = d_108_v43_
                nw27_[int(23)] = True
                nw27_[int(24)] = (d_168_v93_).f29
                d_175_v100_ = nw27_
                d_176_v101_: _dafny.Map
                d_176_v101_ = _dafny.Map({d_175_v100_: True})
                d_177_v102_: _dafny.Array
                nw28_ = _dafny.Array(None, 28)
                nw28_[int(0)] = d_108_v43_
                nw28_[int(1)] = (d_171_v96_)[default__.safeIndex(len(d_169_v94_), len(d_171_v96_))]
                nw28_[int(2)] = (False) and (False)
                nw28_[int(3)] = (d_168_v93_).f29
                nw28_[int(4)] = (_dafny.MultiSet([d_118_v52_])) != (d_172_v97_)
                nw28_[int(5)] = (True) == (default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shqrkmdpp")), True, globalState))
                nw28_[int(6)] = d_108_v43_
                nw28_[int(7)] = d_108_v43_
                nw28_[int(8)] = (d_168_v93_).f29
                nw28_[int(9)] = (d_168_v93_).f29
                nw28_[int(10)] = d_108_v43_
                nw28_[int(11)] = (len(d_174_v99_)) < (d_118_v52_)
                nw28_[int(12)] = (len(p0)) < (-454)
                nw28_[int(13)] = d_108_v43_
                nw28_[int(14)] = default__.fm0(p0, not(False), globalState)
                nw28_[int(15)] = (d_168_v93_).f29
                nw28_[int(16)] = d_108_v43_
                nw28_[int(17)] = (not(((d_176_v101_)[d_175_v100_] if (d_175_v100_) in (d_176_v101_) else d_108_v43_))) not in (d_171_v96_)
                nw28_[int(18)] = ((d_168_v93_).f29) or (d_108_v43_)
                nw28_[int(19)] = d_108_v43_
                nw28_[int(20)] = d_108_v43_
                nw28_[int(21)] = (d_167_i8_) > ((0) - (d_118_v52_))
                nw28_[int(22)] = True
                nw28_[int(23)] = (d_168_v93_).f29
                nw28_[int(24)] = d_108_v43_
                nw28_[int(25)] = d_108_v43_
                nw28_[int(26)] = (d_168_v93_).f29
                nw28_[int(27)] = (D6_DC18(p0, len(d_171_v96_), d_108_v43_, 834)).cf36
                d_177_v102_ = nw28_
                d_177_v102_ = d_175_v100_
                d_178_v103_: _dafny.Map
                d_178_v103_ = _dafny.Map({default__.fm7(p0, d_167_i8_, globalState): _dafny.SeqWithoutIsStrInference([_dafny.Map({True: d_118_v52_}) for d_179_i9_ in range(default__.abs(848))])})
                d_180_v104_: _dafny.Map
                d_180_v104_ = _dafny.Map({(d_168_v93_).f29: d_167_i8_})
                d_181_v105_: _dafny.Seq
                d_181_v105_ = _dafny.SeqWithoutIsStrInference([d_180_v104_])
                d_178_v103_ = (d_178_v103_).set(d_46_v0_, d_181_v105_)
                d_182_v106_: _dafny.MultiSet
                d_182_v106_ = _dafny.MultiSet([(d_168_v93_).f29])
                d_183_v107_: _dafny.Seq
                d_183_v107_ = _dafny.SeqWithoutIsStrInference([((d_182_v106_)[not((d_168_v93_).f29)] if (not((d_168_v93_).f29)) in (d_182_v106_) else d_167_i8_), d_118_v52_])
                d_184_v108_: _dafny.Array
                nw29_ = _dafny.Array(int(0), 10)
                d_184_v108_ = nw29_
                d_185_v109_: _dafny.Set
                d_185_v109_ = _dafny.Set({d_184_v108_, d_184_v108_})
                (globalState).f27 = (d_183_v107_)[default__.safeIndex(len((d_185_v109_).intersection(d_185_v109_)), len(d_183_v107_))]
                d_186_v110_: C0
                nw30_ = C0()
                nw30_.ctor__(True)
                d_186_v110_ = nw30_
            elif True:
                d_187_v111_: _dafny.MultiSet
                d_187_v111_ = _dafny.MultiSet([True, d_108_v43_])
                (globalState).f17 = default__.safeDivisionInt(d_167_i8_, default__.safeDivisionInt(((d_187_v111_)[d_108_v43_] if (d_108_v43_) in (d_187_v111_) else d_118_v52_), 892))
                (globalState).f14 = d_108_v43_
                d_188_v112_: _dafny.Seq
                d_188_v112_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvspt"))])
                d_189_v113_: _dafny.Map
                d_189_v113_ = _dafny.Map({len(d_119_v53_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjrw"))})
                d_190_v114_: _dafny.Array
                nw31_ = _dafny.Array(None, 20)
                nw31_[int(0)] = (d_188_v112_)[default__.safeIndex(d_118_v52_, len(d_188_v112_))]
                nw31_[int(1)] = p0
                nw31_[int(2)] = _dafny.SeqWithoutIsStrInference([d_46_v0_ for d_191_i10_ in range(default__.abs(317))])
                nw31_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iattnj"))
                nw31_[int(4)] = (p0) + (p0)
                nw31_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amks"))
                nw31_[int(6)] = p0
                nw31_[int(7)] = (p0) + (p0)
                nw31_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uut"))
                nw31_[int(9)] = p0
                nw31_[int(10)] = p0
                nw31_[int(11)] = p0
                nw31_[int(12)] = p0
                nw31_[int(13)] = p0
                nw31_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rivncvdvu"))
                nw31_[int(15)] = (d_188_v112_)[default__.safeIndex(d_118_v52_, len(d_188_v112_))]
                nw31_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pukcc"))
                nw31_[int(17)] = p0
                nw31_[int(18)] = ((d_189_v113_)[len(p0)] if (len(p0)) in (d_189_v113_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpehwcpgi")))
                nw31_[int(19)] = p0
                d_190_v114_ = nw31_
                index10_ = default__.safeIndex(80, (d_190_v114_).length(0))
                (d_190_v114_)[index10_] = p0
                d_192_v115_: _dafny.Map
                d_192_v115_ = _dafny.Map({d_108_v43_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cugx"))})
                index11_ = default__.safeIndex(80, (d_190_v114_).length(0))
                (d_190_v114_)[index11_] = (p0) + ((((d_192_v115_)[d_108_v43_] if (d_108_v43_) in (d_192_v115_) else p0)) + (p0))
                (globalState).f17 = ((d_118_v52_) - (d_167_i8_)) + (d_118_v52_)
                d_193_v116_: D2
                d_193_v116_ = D2_DC8(p0)
                d_194_v117_: _dafny.Seq
                d_194_v117_ = _dafny.SeqWithoutIsStrInference([default__.fm14(d_193_v116_, globalState), d_187_v111_, d_187_v111_, d_187_v111_, d_187_v111_])
                d_195_v118_: _dafny.Map
                d_195_v118_ = _dafny.Map({(d_194_v117_)[default__.safeIndex(d_118_v52_, len(d_194_v117_))]: d_108_v43_})
                d_195_v118_ = d_195_v118_
            d_118_v52_ = len(p0)
            d_196_v119_: _dafny.Seq
            d_196_v119_ = _dafny.SeqWithoutIsStrInference([d_118_v52_])
            d_196_v119_ = d_196_v119_
            d_197_v120_: _dafny.Map
            d_197_v120_ = _dafny.Map({d_108_v43_: d_167_i8_})
            d_197_v120_ = d_197_v120_
        r0 = p0
        d_198_v121_: _dafny.MultiSet
        d_198_v121_ = _dafny.MultiSet([False])
        d_199_v122_: _dafny.Seq
        d_199_v122_ = _dafny.SeqWithoutIsStrInference([((d_198_v121_)[d_108_v43_] if (d_108_v43_) in (d_198_v121_) else d_118_v52_)])
        d_200_v123_: _dafny.Seq
        d_200_v123_ = _dafny.SeqWithoutIsStrInference([d_108_v43_])
        d_201_v124_: _dafny.Map
        d_201_v124_ = _dafny.Map({d_118_v52_: len(d_200_v123_)})
        d_202_v125_: _dafny.Set
        d_202_v125_ = _dafny.Set({d_201_v124_})
        d_203_v126_: _dafny.MultiSet
        d_203_v126_ = _dafny.MultiSet([d_118_v52_, d_118_v52_])
        d_204_v127_: _dafny.Map
        d_204_v127_ = _dafny.Map({d_118_v52_: d_108_v43_})
        d_205_v128_: _dafny.Map
        d_205_v128_ = _dafny.Map({d_204_v127_: False})
        nw32_ = _dafny.Array(None, 25)
        nw32_[int(0)] = len(d_199_v122_)
        nw32_[int(1)] = len(d_200_v123_)
        nw32_[int(2)] = d_118_v52_
        nw32_[int(3)] = len(d_202_v125_)
        nw32_[int(4)] = d_118_v52_
        nw32_[int(5)] = d_118_v52_
        nw32_[int(6)] = d_118_v52_
        nw32_[int(7)] = default__.safeDivisionInt((0) - (d_118_v52_), len(p0))
        nw32_[int(8)] = (0) - (d_118_v52_)
        nw32_[int(9)] = d_118_v52_
        nw32_[int(10)] = (721 if d_108_v43_ else d_118_v52_)
        nw32_[int(11)] = len(((p0) + (p0)).set(default__.safeIndex(d_118_v52_, len((p0) + (p0))), _dafny.CodePoint('b')))
        nw32_[int(12)] = 580
        nw32_[int(13)] = len(default__.fm9(d_118_v52_, globalState))
        nw32_[int(14)] = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_118_v52_, d_118_v52_]))) | (d_203_v126_)).cardinality
        nw32_[int(15)] = d_118_v52_
        nw32_[int(16)] = (d_118_v52_) + (d_118_v52_)
        nw32_[int(17)] = ((d_198_v121_)[d_108_v43_] if (d_108_v43_) in (d_198_v121_) else d_118_v52_)
        nw32_[int(18)] = d_118_v52_
        nw32_[int(19)] = d_118_v52_
        nw32_[int(20)] = default__.safeDivisionInt(len(p0), d_118_v52_)
        nw32_[int(21)] = (len(default__.fm12(d_118_v52_, -213, d_108_v43_, globalState))) * (d_118_v52_)
        nw32_[int(22)] = len(p0)
        nw32_[int(23)] = default__.safeDivisionInt((0) - (d_118_v52_), len(d_205_v128_))
        nw32_[int(24)] = (-200) - ((0) - (d_118_v52_))
        r1 = nw32_
        d_206_v129_: D7
        d_206_v129_ = D7_DC19(d_198_v121_)
        d_207_v130_: D7
        d_207_v130_ = D7_DC21(d_206_v129_)
        pat_let_tv0_ = d_118_v52_
        pat_let_tv1_ = d_203_v126_
        def lambda6_(source5_):
            if source5_.is_DC20:
                d_208___mcc_h6_ = source5_.cf39
                d_209___mcc_h7_ = source5_.cf40
                d_210___mcc_h8_ = source5_.cf41
                d_211___mcc_h9_ = source5_.cf42
                d_212_cf42_ = d_211___mcc_h9_
                d_213_cf41_ = d_210___mcc_h8_
                d_214_cf40_ = d_209___mcc_h7_
                d_215_cf39_ = d_208___mcc_h6_
                return d_213_cf41_
            elif source5_.is_DC19:
                d_216___mcc_h10_ = source5_.cf38
                d_217_cf38_ = d_216___mcc_h10_
                return pat_let_tv0_
            elif True:
                d_218___mcc_h11_ = source5_.cf43
                d_219_cf43_ = d_218___mcc_h11_
                return (pat_let_tv1_).cardinality

        r2 = (0) - (lambda6_(d_207_v130_))
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_220_v0_: bool
        d_220_v0_ = True
        d_221_v1_: _dafny.MultiSet
        d_221_v1_ = _dafny.MultiSet([d_220_v0_, d_220_v0_])
        d_222_v2_: _dafny.Seq
        d_222_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_220_v0_])), d_221_v1_])
        d_223_v3_: _dafny.Seq
        d_223_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujuh"))
        d_224_v4_: _dafny.Seq
        d_224_v4_ = _dafny.SeqWithoutIsStrInference([d_220_v0_, d_220_v0_, d_220_v0_])
        d_225_v5_: _dafny.Map
        d_225_v5_ = _dafny.Map({(d_224_v4_)[default__.safeIndex(len(d_223_v3_), len(d_224_v4_))]: d_224_v4_})
        d_226_v6_: _dafny.Array
        nw33_ = _dafny.Array(_dafny.Map({}), 21)
        d_226_v6_ = nw33_
        d_227_v7_: int
        d_227_v7_ = 773
        d_228_v8_: _dafny.Set
        d_228_v8_ = _dafny.Set({d_220_v0_})
        d_229_v9_: _dafny.Array
        nw34_ = _dafny.Array(None, 20)
        nw34_[int(0)] = d_227_v7_
        nw34_[int(1)] = -741
        nw34_[int(2)] = d_227_v7_
        nw34_[int(3)] = d_227_v7_
        nw34_[int(4)] = -617
        nw34_[int(5)] = d_227_v7_
        nw34_[int(6)] = len(d_228_v8_)
        nw34_[int(7)] = d_227_v7_
        nw34_[int(8)] = 744
        nw34_[int(9)] = d_227_v7_
        nw34_[int(10)] = d_227_v7_
        nw34_[int(11)] = 462
        nw34_[int(12)] = d_227_v7_
        nw34_[int(13)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvc")))
        nw34_[int(14)] = d_227_v7_
        nw34_[int(15)] = d_227_v7_
        nw34_[int(16)] = d_227_v7_
        nw34_[int(17)] = 726
        nw34_[int(18)] = d_227_v7_
        nw34_[int(19)] = d_227_v7_
        d_229_v9_ = nw34_
        d_230_v10_: _dafny.MultiSet
        d_230_v10_ = _dafny.MultiSet([d_229_v9_])
        d_231_v12_: _dafny.Map
        d_231_v12_ = _dafny.Map({d_220_v0_: not(d_220_v0_)})
        d_232_v13_: _dafny.Seq
        d_232_v13_ = _dafny.SeqWithoutIsStrInference([d_231_v12_, d_231_v12_])
        d_233_v14_: _dafny.Array
        def lambda7_(d_234_v0_):
            def lambda8_(d_235_i0_):
                return d_234_v0_

            return lambda8_

        init3_ = lambda7_(d_220_v0_)
        nw35_ = _dafny.Array(None, 16)
        for i0_3_ in range(nw35_.length(0)):
            nw35_[i0_3_] = init3_(i0_3_)
        d_233_v14_ = nw35_
        d_236_v15_: _dafny.Map
        d_236_v15_ = _dafny.Map({155: d_220_v0_})
        d_237_v16_: str
        d_237_v16_ = _dafny.CodePoint('a')
        d_238_v17_: _dafny.Set
        d_238_v17_ = _dafny.Set({d_237_v16_, d_237_v16_})
        d_239_v18_: _dafny.Map
        d_239_v18_ = _dafny.Map({d_236_v15_: d_238_v17_})
        d_240_globalState_: GlobalState
        nw36_ = GlobalState()
        def iife5_():
            coll1_ = _dafny.Map()
            compr_1_: _dafny.Map
            for compr_1_ in (d_232_v13_).Elements:
                d_241_v11_: _dafny.Map = compr_1_
                if (d_241_v11_) in (d_232_v13_):
                    coll1_[d_241_v11_] = d_227_v7_
            return _dafny.Map(coll1_)
        nw36_.ctor__(d_222_v2_, -624, _dafny.CodePoint('l'), 951, True, 868, d_223_v3_, d_225_v5_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwvoes"))) + (d_223_v3_), d_226_v6_, 113, (d_221_v1_).set(d_220_v0_, default__.abs(d_227_v7_)), 664, 310, False, 536, 495, -530, (d_230_v10_) - (_dafny.MultiSet([d_229_v9_, d_229_v9_, d_229_v9_])), iife5_()
        , d_233_v14_, 294, True, 119, (d_239_v18_) | (_dafny.Map({d_236_v15_: d_238_v17_})), -643, d_221_v1_, 239, d_229_v9_)
        d_240_globalState_ = nw36_
        d_242_i1_: int
        d_242_i1_ = 0
        with _dafny.label("2"):
            while not(not (d_220_v0_) or (d_220_v0_)):
                with _dafny.c_label("2"):
                    if (d_242_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_242_i1_ = (d_242_i1_) + (1)
                    d_243_v19_: C0
                    nw37_ = C0()
                    nw37_.ctor__((len(d_236_v15_)) <= (579))
                    d_243_v19_ = nw37_
                    d_244_v20_: _dafny.Seq
                    d_244_v20_ = _dafny.SeqWithoutIsStrInference([d_223_v3_, _dafny.SeqWithoutIsStrInference([d_237_v16_ for d_245_i2_ in range(default__.abs(-469))])])
                    d_220_v0_ = default__.fm0((d_244_v20_)[default__.safeIndex(149, len(d_244_v20_))], d_220_v0_, d_240_globalState_)
                    d_229_v9_ = d_229_v9_
                    (d_240_globalState_).f27 = d_227_v7_
                    pass
            pass
        d_246_v21_: _dafny.Seq
        d_247_v22_: _dafny.Array
        d_248_v23_: int
        out0_: _dafny.Seq
        out1_: _dafny.Array
        out2_: int
        out0_, out1_, out2_ = default__.m0(d_223_v3_, d_240_globalState_)
        d_246_v21_ = out0_
        d_247_v22_ = out1_
        d_248_v23_ = out2_
        hi1_ = len(_dafny.SeqWithoutIsStrInference([d_237_v16_ for d_249_i4_ in range(default__.abs(169))]))
        for d_250_i3_ in range(len(d_246_v21_), hi1_):
            d_231_v12_ = (d_231_v12_).set((d_228_v8_).ispropersubset(d_228_v8_), d_220_v0_)
            source6_ = default__.fm1(d_240_globalState_)
            if source6_.is_DC0:
                d_251___mcc_h0_ = source6_.cf0
                d_252___mcc_h1_ = source6_.cf1
                d_253_cf1_ = d_252___mcc_h1_
                d_254_cf0_ = d_251___mcc_h0_
                index12_ = default__.safeIndex(800, (d_247_v22_).length(0))
                (d_247_v22_)[index12_] = default__.safeDivisionInt(d_248_v23_, d_248_v23_)
                d_255_v24_: C0
                nw38_ = C0()
                nw38_.ctor__(d_220_v0_)
                d_255_v24_ = nw38_
                d_256_v25_: _dafny.Array
                nw39_ = _dafny.Array(None, 1)
                nw39_[int(0)] = (d_255_v24_ if d_220_v0_ else d_255_v24_)
                d_256_v25_ = nw39_
                d_257_v26_: _dafny.Map
                d_257_v26_ = _dafny.Map({(d_255_v24_).f29: d_227_v7_})
                index13_ = default__.safeIndex(800, (d_247_v22_).length(0))
                rhs13_ = (((d_257_v26_)[(d_255_v24_).f29] if ((d_255_v24_).f29) in (d_257_v26_) else d_253_cf1_) if (d_220_v0_) and ((d_255_v24_).f29) else d_248_v23_)
                rhs14_ = d_250_i3_
                rhs15_ = d_256_v25_
                lhs6_ = d_247_v22_
                lhs7_ = default__.safeIndex(800, (d_247_v22_).length(0))
                lhs6_[lhs7_] = rhs13_
                d_254_cf0_ = rhs14_
                d_256_v25_ = rhs15_
                d_258_v27_: C0
                nw40_ = C0()
                nw40_.ctor__(d_220_v0_)
                d_258_v27_ = nw40_
                d_259_v28_: _dafny.Array
                nw41_ = _dafny.Array(None, 18)
                nw41_[int(0)] = d_230_v10_
                nw41_[int(1)] = (d_230_v10_ if False else d_230_v10_)
                nw41_[int(2)] = _dafny.MultiSet([d_229_v9_, d_229_v9_, d_229_v9_])
                nw41_[int(3)] = _dafny.MultiSet([d_229_v9_, d_229_v9_])
                nw41_[int(4)] = d_230_v10_
                nw41_[int(5)] = d_230_v10_
                nw41_[int(6)] = d_230_v10_
                nw41_[int(7)] = d_230_v10_
                nw41_[int(8)] = d_230_v10_
                nw41_[int(9)] = d_230_v10_
                nw41_[int(10)] = d_230_v10_
                nw41_[int(11)] = (d_230_v10_).set(d_229_v9_, default__.abs(437))
                nw41_[int(12)] = d_230_v10_
                nw41_[int(13)] = (d_230_v10_) - (d_230_v10_)
                nw41_[int(14)] = d_230_v10_
                nw41_[int(15)] = d_230_v10_
                nw41_[int(16)] = d_230_v10_
                nw41_[int(17)] = (d_230_v10_) | (_dafny.MultiSet([d_229_v9_]))
                d_259_v28_ = nw41_
                index14_ = default__.safeIndex(977, (d_259_v28_).length(0))
                (d_259_v28_)[index14_] = _dafny.MultiSet([d_229_v9_, d_229_v9_, d_229_v9_])
                d_260_v29_: _dafny.Array
                nw42_ = _dafny.Array(D0.default()(), 17)
                d_260_v29_ = nw42_
                d_261_v30_: D0
                d_261_v30_ = D0_DC0(781, (0) - (d_253_cf1_))
                pat_let_tv2_ = d_253_cf1_
                index15_ = default__.safeIndex(94, (d_260_v29_).length(0))
                def iife6_(_pat_let2_0):
                    def iife7_(d_262_dt__update__tmp_h0_):
                        def iife8_(_pat_let3_0):
                            def iife9_(d_263_dt__update_hcf0_h0_):
                                return D0_DC0(d_263_dt__update_hcf0_h0_, (d_262_dt__update__tmp_h0_).cf1)
                            return iife9_(_pat_let3_0)
                        return iife8_(pat_let_tv2_)
                    return iife7_(_pat_let2_0)
                (d_260_v29_)[index15_] = iife6_(d_261_v30_)
                index16_ = default__.safeIndex(977, (d_259_v28_).length(0))
                index17_ = default__.safeIndex(94, (d_260_v29_).length(0))
                rhs16_ = _dafny.MultiSet([d_229_v9_, d_229_v9_, d_229_v9_, d_229_v9_, d_229_v9_])
                rhs17_ = D0_DC0(d_254_cf0_, d_254_cf0_)
                rhs18_ = (d_258_v27_).f29
                rhs19_ = d_255_v24_
                rhs20_ = ((0) - (d_253_cf1_)) + ((d_253_cf1_ if (d_258_v27_).f29 else d_254_cf0_))
                lhs8_ = d_259_v28_
                lhs9_ = default__.safeIndex(977, (d_259_v28_).length(0))
                lhs10_ = d_260_v29_
                lhs11_ = default__.safeIndex(94, (d_260_v29_).length(0))
                lhs12_ = d_240_globalState_
                lhs13_ = d_240_globalState_
                lhs8_[lhs9_] = rhs16_
                lhs10_[lhs11_] = rhs17_
                lhs12_.f14 = rhs18_
                d_255_v24_ = rhs19_
                lhs13_.f13 = rhs20_
                d_264_v31_: _dafny.Map
                d_264_v31_ = _dafny.Map({True: d_258_v27_})
                d_264_v31_ = (d_264_v31_).set(not((d_258_v27_).f29), d_255_v24_)
            elif True:
                (d_240_globalState_).f27 = (d_227_v7_) + (316)
                index18_ = default__.safeIndex(335, (d_233_v14_).length(0))
                (d_233_v14_)[index18_] = (d_220_v0_) == (d_220_v0_)
                index19_ = default__.safeIndex(335, (d_233_v14_).length(0))
                (d_233_v14_)[index19_] = not(d_220_v0_)
                d_248_v23_ = d_248_v23_
                d_265_v32_: _dafny.Array
                def lambda9_(d_266_v0_, d_267_i3_):
                    def lambda10_(d_268_i5_):
                        return (_dafny.Map({d_266_v0_: 889})) | (_dafny.Map({d_266_v0_: d_267_i3_}))

                    return lambda10_

                init4_ = lambda9_(d_220_v0_, d_250_i3_)
                nw43_ = _dafny.Array(None, 22)
                for i0_4_ in range(nw43_.length(0)):
                    nw43_[i0_4_] = init4_(i0_4_)
                d_265_v32_ = nw43_
                d_269_v33_: _dafny.Map
                d_269_v33_ = _dafny.Map({(d_233_v14_)[default__.safeIndex(335, (d_233_v14_).length(0))]: 878})
                index20_ = default__.safeIndex(88, (d_265_v32_).length(0))
                (d_265_v32_)[index20_] = (d_269_v33_ if ((d_231_v12_)[(d_233_v14_)[default__.safeIndex(335, (d_233_v14_).length(0))]] if ((d_233_v14_)[default__.safeIndex(335, (d_233_v14_).length(0))]) in (d_231_v12_) else d_220_v0_) else d_269_v33_)
                index21_ = default__.safeIndex(88, (d_265_v32_).length(0))
                (d_265_v32_)[index21_] = d_269_v33_
            d_270_v34_: _dafny.Array
            nw44_ = _dafny.Array(_dafny.Seq({}), 7)
            d_270_v34_ = nw44_
            index22_ = default__.safeIndex(919, (d_270_v34_).length(0))
            (d_270_v34_)[index22_] = (d_224_v4_) + (d_224_v4_)
            index23_ = default__.safeIndex(919, (d_270_v34_).length(0))
            (d_270_v34_)[index23_] = (d_224_v4_) + (d_224_v4_)
            d_271_v35_: C0
            nw45_ = C0()
            nw45_.ctor__(d_220_v0_)
            d_271_v35_ = nw45_
        d_272_v36_: _dafny.Array
        nw46_ = _dafny.Array(_dafny.MultiSet({}), 21)
        d_272_v36_ = nw46_
        d_273_v37_: _dafny.MultiSet
        d_273_v37_ = _dafny.MultiSet([d_233_v14_, d_233_v14_, d_233_v14_, d_233_v14_, d_233_v14_])
        index24_ = default__.safeIndex(107, (d_272_v36_).length(0))
        (d_272_v36_)[index24_] = d_273_v37_
        index25_ = default__.safeIndex(107, (d_272_v36_).length(0))
        (d_272_v36_)[index25_] = ((d_273_v37_) - (d_273_v37_)).intersection((d_273_v37_).intersection(d_273_v37_))
        if (d_248_v23_) <= (d_248_v23_):
            d_237_v16_ = d_237_v16_
            d_274_v38_: _dafny.Seq
            d_275_v39_: _dafny.Array
            d_276_v40_: int
            out3_: _dafny.Seq
            out4_: _dafny.Array
            out5_: int
            out3_, out4_, out5_ = default__.m0(d_223_v3_, d_240_globalState_)
            d_274_v38_ = out3_
            d_275_v39_ = out4_
            d_276_v40_ = out5_
            d_277_v41_: _dafny.Seq
            d_278_v42_: _dafny.Array
            d_279_v43_: int
            out6_: _dafny.Seq
            out7_: _dafny.Array
            out8_: int
            out6_, out7_, out8_ = default__.m0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "konvcdp")), d_240_globalState_)
            d_277_v41_ = out6_
            d_278_v42_ = out7_
            d_279_v43_ = out8_
            if d_220_v0_:
                (d_240_globalState_).f13 = d_227_v7_
                d_280_v44_: C0
                nw47_ = C0()
                nw47_.ctor__(True)
                d_280_v44_ = nw47_
                d_281_v45_: _dafny.Map
                d_281_v45_ = _dafny.Map({309: d_276_v40_})
                d_282_v46_: C0
                nw48_ = C0()
                nw48_.ctor__((d_281_v45_) == (d_281_v45_))
                d_282_v46_ = nw48_
                d_223_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
                (d_240_globalState_).f14 = ((d_236_v15_)[d_279_v43_] if (d_279_v43_) in (d_236_v15_) else ((0) - ((0) - (d_276_v40_))) <= (d_276_v40_))
            elif True:
                (d_240_globalState_).f27 = 157
                d_283_v47_: _dafny.Map
                d_283_v47_ = _dafny.Map({d_237_v16_: d_233_v14_})
                rhs21_ = ((d_283_v47_) | (_dafny.Map({d_237_v16_: d_233_v14_}))).set(d_237_v16_, d_233_v14_)
                rhs22_ = default__.fm0(_dafny.SeqWithoutIsStrInference([d_237_v16_ for d_284_i6_ in range(default__.abs(-390))]), (d_236_v15_) != (d_236_v15_), d_240_globalState_)
                rhs23_ = d_221_v1_
                lhs14_ = d_240_globalState_
                d_283_v47_ = rhs21_
                lhs14_.f14 = rhs22_
                d_221_v1_ = rhs23_
                d_285_v48_: _dafny.Seq
                d_285_v48_ = _dafny.SeqWithoutIsStrInference([d_248_v23_])
                d_286_v49_: _dafny.Map
                d_286_v49_ = _dafny.Map({d_276_v40_: 663})
                d_287_v50_: _dafny.Seq
                d_287_v50_ = _dafny.SeqWithoutIsStrInference([745, (d_285_v48_)[default__.safeIndex(len(d_286_v49_), len(d_285_v48_))]])
                (d_240_globalState_).f13 = default__.safeDivisionInt(57, (d_279_v43_) + (len(d_287_v50_)))
                d_288_v51_: C0
                nw49_ = C0()
                nw49_.ctor__((d_285_v48_) <= (d_287_v50_))
                d_288_v51_ = nw49_
                d_289_v52_: _dafny.MultiSet
                d_289_v52_ = _dafny.MultiSet([d_237_v16_, d_237_v16_])
                d_290_v53_: _dafny.Set
                d_290_v53_ = _dafny.Set({default__.fm2(d_289_v52_, d_248_v23_, d_227_v7_, d_220_v0_, d_240_globalState_)})
                (d_240_globalState_).f25 = len(((d_236_v15_).set(len(d_290_v53_), (d_288_v51_).f29)).set(886, d_220_v0_))
            d_291_v54_: C0
            nw50_ = C0()
            nw50_.ctor__(d_220_v0_)
            d_291_v54_ = nw50_
        elif True:
            (d_240_globalState_).f2 = d_237_v16_
            (d_240_globalState_).f25 = d_248_v23_
            d_220_v0_ = (d_227_v7_) >= (d_248_v23_)
            (d_240_globalState_).f27 = (0) - (d_227_v7_)
            d_231_v12_ = (d_231_v12_).set(d_220_v0_, (d_220_v0_ if default__.fm0(d_223_v3_, d_220_v0_, d_240_globalState_) else ((d_236_v15_)[d_248_v23_] if (d_248_v23_) in (d_236_v15_) else d_220_v0_)))
        d_292_v55_: C0
        nw51_ = C0()
        nw51_.ctor__((d_237_v16_) in (d_246_v21_))
        d_292_v55_ = nw51_
        (d_240_globalState_).f27 = d_248_v23_
        if (d_292_v55_).f29:
            (d_240_globalState_).f13 = d_248_v23_
            (d_240_globalState_).f14 = (d_292_v55_).f29
            index26_ = default__.safeIndex(595, (d_233_v14_).length(0))
            (d_233_v14_)[index26_] = d_220_v0_
            d_293_v56_: D1
            d_293_v56_ = D1_DC2(d_246_v21_)
            index27_ = default__.safeIndex(595, (d_233_v14_).length(0))
            (d_233_v14_)[index27_] = (((d_223_v3_).set(default__.safeIndex(d_248_v23_, len(d_223_v3_)), d_237_v16_)) + ((d_293_v56_).cf2)) == (d_223_v3_)
            (d_240_globalState_).f14 = not((d_233_v14_)[default__.safeIndex(595, (d_233_v14_).length(0))])
            d_294_v57_: _dafny.Set
            d_294_v57_ = _dafny.Set({d_248_v23_})
            d_295_v58_: _dafny.Seq
            d_295_v58_ = _dafny.SeqWithoutIsStrInference([d_227_v7_, (0) - (-473), d_227_v7_])
            d_296_v59_: _dafny.Seq
            d_296_v59_ = _dafny.SeqWithoutIsStrInference([d_294_v57_, _dafny.Set({(0) - ((d_295_v58_)[default__.safeIndex(161, len(d_295_v58_))])})])
            d_297_v60_: _dafny.MultiSet
            d_297_v60_ = _dafny.MultiSet([d_294_v57_, (d_296_v59_)[default__.safeIndex(d_248_v23_, len(d_296_v59_))]])
            d_297_v60_ = d_297_v60_
        elif True:
            d_298_v61_: D1
            d_298_v61_ = D1_DC4((d_292_v55_).f29, d_227_v7_)
            index28_ = default__.safeIndex(305, (d_229_v9_).length(0))
            (d_229_v9_)[index28_] = (d_298_v61_).cf6
            d_299_v62_: _dafny.Map
            d_299_v62_ = _dafny.Map({(d_292_v55_).f29: d_231_v12_})
            d_300_v63_: _dafny.Map
            d_300_v63_ = _dafny.Map({len(d_299_v62_): d_237_v16_})
            d_301_v64_: _dafny.Map
            d_301_v64_ = _dafny.Map({len(d_300_v63_): _dafny.SeqWithoutIsStrInference([d_220_v0_, (d_292_v55_).f29, d_220_v0_, (d_292_v55_).f29, (d_292_v55_).f29])})
            d_302_v65_: _dafny.Map
            d_302_v65_ = _dafny.Map({(d_301_v64_) | (d_301_v64_): d_248_v23_})
            index29_ = default__.safeIndex(305, (d_229_v9_).length(0))
            (d_229_v9_)[index29_] = ((d_302_v65_)[d_301_v64_] if (d_301_v64_) in (d_302_v65_) else d_248_v23_)
            d_303_v66_: _dafny.Seq
            d_303_v66_ = _dafny.SeqWithoutIsStrInference([238, (d_229_v9_)[default__.safeIndex(305, (d_229_v9_).length(0))]])
            d_304_v67_: _dafny.Seq
            d_304_v67_ = _dafny.SeqWithoutIsStrInference([d_248_v23_, (d_303_v66_)[default__.safeIndex(d_227_v7_, len(d_303_v66_))]])
            rhs24_ = d_303_v66_
            rhs25_ = (d_292_v55_).f29
            lhs15_ = d_240_globalState_
            d_303_v66_ = rhs24_
            lhs15_.f14 = rhs25_
            d_305_v68_: D0
            d_305_v68_ = D0_DC1()
            d_305_v68_ = default__.fm3(d_227_v7_, d_240_globalState_)
            if d_220_v0_:
                d_220_v0_ = False
                (d_240_globalState_).f17 = d_227_v7_
                d_306_v69_: _dafny.Seq
                d_307_v70_: _dafny.Array
                d_308_v71_: int
                out9_: _dafny.Seq
                out10_: _dafny.Array
                out11_: int
                out9_, out10_, out11_ = default__.m0(d_246_v21_, d_240_globalState_)
                d_306_v69_ = out9_
                d_307_v70_ = out10_
                d_308_v71_ = out11_
                d_309_v72_: _dafny.Map
                d_309_v72_ = _dafny.Map({(d_292_v55_).f29: len(_dafny.SeqWithoutIsStrInference([d_248_v23_, d_308_v71_]))})
                rhs26_ = d_308_v71_
                rhs27_ = default__.fm0(d_246_v21_, not(d_220_v0_), d_240_globalState_)
                rhs28_ = d_220_v0_
                rhs29_ = ((d_229_v9_)[default__.safeIndex(305, (d_229_v9_).length(0))]) + (default__.safeModuloInt(((d_309_v72_)[(d_292_v55_).f29] if ((d_292_v55_).f29) in (d_309_v72_) else d_248_v23_), (d_229_v9_)[default__.safeIndex(305, (d_229_v9_).length(0))]))
                lhs16_ = d_240_globalState_
                lhs17_ = d_240_globalState_
                lhs18_ = d_240_globalState_
                lhs19_ = d_240_globalState_
                lhs16_.f25 = rhs26_
                lhs17_.f14 = rhs27_
                lhs18_.f14 = rhs28_
                lhs19_.f25 = rhs29_
                (d_240_globalState_).f14 = (((d_292_v55_).f29 if not(default__.fm0(d_246_v21_, (d_292_v55_).f29, d_240_globalState_)) else (d_292_v55_).f29)) and ((d_292_v55_).f29)
            elif True:
                d_310_v73_: D1
                d_310_v73_ = D1_DC2(d_246_v21_)
                d_310_v73_ = default__.fm4((509) >= (d_227_v7_), (d_292_v55_).f29, d_240_globalState_)
                d_311_v74_: _dafny.Set
                d_311_v74_ = _dafny.Set({d_246_v21_})
                d_312_v75_: _dafny.Seq
                d_312_v75_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_313_i7_ in range(default__.abs(-172))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_314_i8_ in range(default__.abs(57))])])
                def iife10_():
                    coll2_ = _dafny.Set()
                    compr_2_: _dafny.Seq
                    for compr_2_ in (d_312_v75_).Elements:
                        d_315_v76_: _dafny.Seq = compr_2_
                        if (d_315_v76_) in (d_312_v75_):
                            coll2_ = coll2_.union(_dafny.Set([d_315_v76_]))
                    return _dafny.Set(coll2_)
                rhs30_ = (d_311_v74_).ispropersubset((d_311_v74_) - (iife10_()
                ))
                rhs31_ = (d_224_v4_)[default__.safeIndex((d_229_v9_)[default__.safeIndex(305, (d_229_v9_).length(0))], len(d_224_v4_))]
                rhs32_ = d_305_v68_
                d_220_v0_ = rhs30_
                d_220_v0_ = rhs31_
                d_305_v68_ = rhs32_
                index30_ = default__.safeIndex(305, (d_229_v9_).length(0))
                (d_229_v9_)[index30_] = (0) - ((d_248_v23_) - (((0) - (d_248_v23_)) - ((d_229_v9_)[default__.safeIndex(305, (d_229_v9_).length(0))])))
                (d_240_globalState_).f17 = (0) - ((d_304_v67_)[default__.safeIndex(len(d_303_v66_), len(d_304_v67_))])
                (d_240_globalState_).f14 = False
            d_316_v77_: _dafny.Seq
            d_317_v78_: _dafny.Array
            d_318_v79_: int
            out12_: _dafny.Seq
            out13_: _dafny.Array
            out14_: int
            out12_, out13_, out14_ = default__.m0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arg")), d_240_globalState_)
            d_316_v77_ = out12_
            d_317_v78_ = out13_
            d_318_v79_ = out14_
        if (d_227_v7_) <= (d_248_v23_):
            d_319_v80_: _dafny.Seq
            d_320_v81_: _dafny.Array
            d_321_v82_: int
            out15_: _dafny.Seq
            out16_: _dafny.Array
            out17_: int
            out15_, out16_, out17_ = default__.m0(d_223_v3_, d_240_globalState_)
            d_319_v80_ = out15_
            d_320_v81_ = out16_
            d_321_v82_ = out17_
            d_322_v83_: _dafny.Seq
            d_323_v84_: _dafny.Array
            d_324_v85_: int
            out18_: _dafny.Seq
            out19_: _dafny.Array
            out20_: int
            out18_, out19_, out20_ = default__.m0(d_223_v3_, d_240_globalState_)
            d_322_v83_ = out18_
            d_323_v84_ = out19_
            d_324_v85_ = out20_
            (d_240_globalState_).f14 = (d_292_v55_).f29
            d_325_v86_: _dafny.Array
            nw52_ = _dafny.Array(_dafny.CodePoint('D'), 24)
            d_325_v86_ = nw52_
            d_325_v86_ = d_325_v86_
            index31_ = default__.safeIndex(579, (d_233_v14_).length(0))
            (d_233_v14_)[index31_] = (d_292_v55_).f29
            index32_ = default__.safeIndex(579, (d_233_v14_).length(0))
            rhs33_ = (0) - ((0) - (d_321_v82_))
            rhs34_ = d_220_v0_
            lhs20_ = d_240_globalState_
            lhs21_ = d_233_v14_
            lhs22_ = default__.safeIndex(579, (d_233_v14_).length(0))
            lhs20_.f17 = rhs33_
            lhs21_[lhs22_] = rhs34_
        elif True:
            d_326_v87_: _dafny.MultiSet
            d_326_v87_ = _dafny.MultiSet([d_237_v16_, d_237_v16_, d_237_v16_, d_237_v16_, d_237_v16_])
            d_327_v88_: _dafny.Set
            d_327_v88_ = _dafny.Set({d_248_v23_, default__.fm2(d_326_v87_, d_227_v7_, 331, True, d_240_globalState_), (len(d_236_v15_) if d_220_v0_ else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_328_i9_ in range(default__.abs(-400))]))), d_248_v23_})
            def iife11_():
                coll3_ = _dafny.Set()
                compr_3_: int
                for compr_3_ in (default__.fm5(d_220_v0_, d_248_v23_, (d_292_v55_).f29, d_240_globalState_)).Elements:
                    d_329_v89_: int = compr_3_
                    if (d_329_v89_) in (default__.fm5(d_220_v0_, d_248_v23_, (d_292_v55_).f29, d_240_globalState_)):
                        coll3_ = coll3_.union(_dafny.Set([(d_329_v89_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khn"))))]))
                return _dafny.Set(coll3_)
            d_327_v88_ = ((d_327_v88_) - (d_327_v88_)).intersection(iife11_()
            )
            d_330_v90_: C0
            nw53_ = C0()
            nw53_.ctor__((d_292_v55_).f29)
            d_330_v90_ = nw53_
            if (d_292_v55_).f29:
                d_331_v91_: _dafny.Seq
                d_331_v91_ = _dafny.SeqWithoutIsStrInference([(d_246_v21_) + (d_223_v3_), d_223_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhyao"))])
                d_331_v91_ = default__.fm6(d_246_v21_, D1_DC2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yuhsgwf"))), (0) - (d_248_v23_), (d_227_v7_) >= (d_248_v23_), d_240_globalState_)
                (d_240_globalState_).f17 = d_248_v23_
                d_332_v92_: _dafny.Map
                d_332_v92_ = _dafny.Map({(0) - ((0) - (d_227_v7_)): d_330_v90_})
                (d_240_globalState_).f14 = (len(_dafny.SeqWithoutIsStrInference([(d_292_v55_).f29]))) not in ((d_332_v92_) | ((d_332_v92_).set((0) - (d_248_v23_), d_330_v90_)))
                index33_ = default__.safeIndex(25, (d_229_v9_).length(0))
                (d_229_v9_)[index33_] = (0) - (d_248_v23_)
                index34_ = default__.safeIndex(25, (d_229_v9_).length(0))
                (d_229_v9_)[index34_] = default__.safeDivisionInt(default__.safeModuloInt(d_248_v23_, d_227_v7_), d_248_v23_)
                d_333_v93_: _dafny.Array
                nw54_ = _dafny.Array(None, 3)
                nw54_[int(0)] = d_326_v87_
                nw54_[int(1)] = d_326_v87_
                nw54_[int(2)] = d_326_v87_
                d_333_v93_ = nw54_
                d_334_v94_: _dafny.Set
                d_334_v94_ = _dafny.Set({d_327_v88_})
                d_335_v95_: D2
                d_335_v95_ = D2_DC7(d_334_v94_)
                d_336_v96_: _dafny.Map
                d_336_v96_ = _dafny.Map({d_333_v93_: ((d_335_v95_).cf12).isdisjoint(d_334_v94_)})
                (d_240_globalState_).f14 = ((d_336_v96_)[d_333_v93_] if (d_333_v93_) in (d_336_v96_) else (d_292_v55_).f29)
            elif True:
                d_337_v97_: _dafny.Map
                d_337_v97_ = _dafny.Map({d_330_v90_: d_330_v90_})
                d_330_v90_ = ((d_337_v97_)[d_292_v55_] if (d_292_v55_) in (d_337_v97_) else d_330_v90_)
                d_338_v98_: C0
                nw55_ = C0()
                nw55_.ctor__((d_292_v55_).f29)
                d_338_v98_ = nw55_
                d_339_v99_: _dafny.Seq
                d_340_v100_: _dafny.Array
                d_341_v101_: int
                out21_: _dafny.Seq
                out22_: _dafny.Array
                out23_: int
                out21_, out22_, out23_ = default__.m0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkmvx")), d_240_globalState_)
                d_339_v99_ = out21_
                d_340_v100_ = out22_
                d_341_v101_ = out23_
                (d_240_globalState_).f14 = (d_224_v4_)[default__.safeIndex((0) - (d_227_v7_), len(d_224_v4_))]
                d_342_v102_: _dafny.Array
                def lambda11_(d_343_v3_):
                    def lambda12_(d_344_i10_):
                        return (d_343_v3_)[default__.safeIndex(832, len(d_343_v3_))]

                    return lambda12_

                init5_ = lambda11_(d_223_v3_)
                nw56_ = _dafny.Array(None, 24)
                for i0_5_ in range(nw56_.length(0)):
                    nw56_[i0_5_] = init5_(i0_5_)
                d_342_v102_ = nw56_
                index35_ = default__.safeIndex(784, (d_342_v102_).length(0))
                (d_342_v102_)[index35_] = d_237_v16_
                index36_ = default__.safeIndex(784, (d_342_v102_).length(0))
                rhs35_ = (d_224_v4_) + (d_224_v4_)
                rhs36_ = default__.fm7((d_223_v3_ if (d_292_v55_).f29 else d_246_v21_), (0) - (d_227_v7_), d_240_globalState_)
                lhs23_ = d_342_v102_
                lhs24_ = default__.safeIndex(784, (d_342_v102_).length(0))
                d_224_v4_ = rhs35_
                lhs23_[lhs24_] = rhs36_
            (d_240_globalState_).f13 = ((0) - ((d_248_v23_) * (-758))) * (len(_dafny.SeqWithoutIsStrInference([d_227_v7_ for d_345_i11_ in range(default__.abs(782))])))
            d_346_v103_: _dafny.MultiSet
            d_346_v103_ = _dafny.MultiSet([d_227_v7_, -464, d_227_v7_, d_248_v23_])
            if (d_346_v103_).ispropersubset(d_346_v103_):
                d_227_v7_ = d_227_v7_
                (d_240_globalState_).f14 = (d_292_v55_).f29
                (d_240_globalState_).f17 = d_248_v23_
                d_347_v104_: _dafny.Seq
                d_347_v104_ = _dafny.SeqWithoutIsStrInference([d_227_v7_])
                d_348_v105_: _dafny.Map
                d_348_v105_ = _dafny.Map({d_347_v104_: (d_292_v55_).f29})
                d_349_v106_: _dafny.Seq
                d_349_v106_ = _dafny.SeqWithoutIsStrInference([default__.fm2(_dafny.MultiSet(d_223_v3_), len(d_348_v105_), d_248_v23_, d_220_v0_, d_240_globalState_)])
                index37_ = default__.safeIndex(209, (d_229_v9_).length(0))
                (d_229_v9_)[index37_] = default__.safeModuloInt(d_248_v23_, len(d_349_v106_))
                index38_ = default__.safeIndex(209, (d_229_v9_).length(0))
                (d_229_v9_)[index38_] = d_248_v23_
                index39_ = default__.safeIndex(209, (d_229_v9_).length(0))
                (d_229_v9_)[index39_] = len(d_349_v106_)
            elif True:
                d_350_v107_: _dafny.Map
                d_350_v107_ = _dafny.Map({D2_DC10(d_227_v7_, d_237_v16_): _dafny.MultiSet([(d_330_v90_).f29, d_220_v0_, (d_330_v90_).f29, (d_292_v55_).f29, (d_292_v55_).f29])})
                d_351_v108_: D2
                d_351_v108_ = D2_DC10(d_248_v23_, d_237_v16_)
                d_220_v0_ = ((d_292_v55_).f29) not in (((d_350_v107_)[d_351_v108_] if (d_351_v108_) in (d_350_v107_) else d_221_v1_))
                d_246_v21_ = d_223_v3_
                d_352_v109_: _dafny.Array
                nw57_ = _dafny.Array(_dafny.CodePoint('D'), 11)
                d_352_v109_ = nw57_
                d_352_v109_ = d_352_v109_
                d_353_v110_: _dafny.Array
                nw58_ = _dafny.Array(None, 1)
                nw58_[int(0)] = d_229_v9_
                d_353_v110_ = nw58_
                index40_ = default__.safeIndex(520, (d_353_v110_).length(0))
                (d_353_v110_)[index40_] = d_247_v22_
                index41_ = default__.safeIndex(874, (d_233_v14_).length(0))
                (d_233_v14_)[index41_] = d_220_v0_
                d_354_v111_: _dafny.Map
                d_354_v111_ = _dafny.Map({(d_292_v55_).f29: (d_221_v1_).cardinality})
                d_355_v112_: _dafny.Map
                d_355_v112_ = _dafny.Map({d_227_v7_: d_354_v111_})
                index42_ = default__.safeIndex(520, (d_353_v110_).length(0))
                index43_ = default__.safeIndex(874, (d_233_v14_).length(0))
                rhs37_ = d_229_v9_
                rhs38_ = default__.fm0(d_223_v3_, ((d_330_v90_).f29) in (((d_355_v112_)[d_248_v23_] if (d_248_v23_) in (d_355_v112_) else d_354_v111_)), d_240_globalState_)
                rhs39_ = (d_292_v55_).f29
                lhs25_ = d_353_v110_
                lhs26_ = default__.safeIndex(520, (d_353_v110_).length(0))
                lhs27_ = d_233_v14_
                lhs28_ = default__.safeIndex(874, (d_233_v14_).length(0))
                lhs29_ = d_240_globalState_
                lhs25_[lhs26_] = rhs37_
                lhs27_[lhs28_] = rhs38_
                lhs29_.f14 = rhs39_
                d_356_v113_: _dafny.Seq
                d_357_v114_: _dafny.Array
                d_358_v115_: int
                out24_: _dafny.Seq
                out25_: _dafny.Array
                out26_: int
                out24_, out25_, out26_ = default__.m0((d_223_v3_) + (d_246_v21_), d_240_globalState_)
                d_356_v113_ = out24_
                d_357_v114_ = out25_
                d_358_v115_ = out26_
        if not((d_248_v23_) <= (d_248_v23_)):
            d_359_v116_: C0
            nw59_ = C0()
            nw59_.ctor__(True)
            d_359_v116_ = nw59_
            d_220_v0_ = d_220_v0_
            (d_240_globalState_).f14 = (d_292_v55_).f29
            d_360_v117_: _dafny.Set
            d_360_v117_ = _dafny.Set({d_227_v7_, d_227_v7_})
            d_361_v118_: _dafny.Set
            d_361_v118_ = _dafny.Set({d_360_v117_})
            d_362_v119_: D2
            d_362_v119_ = D2_DC7(d_361_v118_)
            pat_let_tv3_ = d_361_v118_
            pat_let_tv4_ = d_361_v118_
            def iife12_(_pat_let4_0):
                def iife13_(d_363_dt__update__tmp_h1_):
                    def iife14_(_pat_let5_0):
                        def iife15_(d_364_dt__update_hcf12_h0_):
                            return D2_DC7(d_364_dt__update_hcf12_h0_)
                        return iife15_(_pat_let5_0)
                    return iife14_((pat_let_tv3_) - (pat_let_tv4_))
                return iife13_(_pat_let4_0)
            d_362_v119_ = iife12_(D2_DC7(d_361_v118_))
            d_365_v120_: _dafny.Map
            d_365_v120_ = _dafny.Map({d_237_v16_: (d_359_v116_).f29})
            rhs40_ = d_248_v23_
            rhs41_ = d_365_v120_
            rhs42_ = d_227_v7_
            lhs30_ = d_240_globalState_
            lhs31_ = d_240_globalState_
            lhs30_.f17 = rhs40_
            d_365_v120_ = rhs41_
            lhs31_.f27 = rhs42_
        elif True:
            d_366_v121_: _dafny.Seq
            d_367_v122_: _dafny.Array
            d_368_v123_: int
            out27_: _dafny.Seq
            out28_: _dafny.Array
            out29_: int
            out27_, out28_, out29_ = default__.m0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqs")), d_240_globalState_)
            d_366_v121_ = out27_
            d_367_v122_ = out28_
            d_368_v123_ = out29_
            index44_ = default__.safeIndex(390, (d_233_v14_).length(0))
            (d_233_v14_)[index44_] = (d_292_v55_).f29
            index45_ = default__.safeIndex(390, (d_233_v14_).length(0))
            (d_233_v14_)[index45_] = (d_292_v55_).f29
            d_369_v124_: C0
            nw60_ = C0()
            nw60_.ctor__(d_220_v0_)
            d_369_v124_ = nw60_
            index46_ = default__.safeIndex(390, (d_233_v14_).length(0))
            (d_233_v14_)[index46_] = d_220_v0_
            d_370_v125_: _dafny.Map
            d_370_v125_ = _dafny.Map({d_220_v0_: d_366_v121_})
            d_371_v126_: _dafny.MultiSet
            d_371_v126_ = _dafny.MultiSet([d_248_v23_, d_227_v7_, len(((d_370_v125_)[(d_292_v55_).f29] if ((d_292_v55_).f29) in (d_370_v125_) else d_246_v21_)), default__.safeModuloInt(d_227_v7_, d_368_v123_)])
            d_372_v127_: _dafny.Map
            d_372_v127_ = _dafny.Map({d_248_v23_: _dafny.MultiSet([d_368_v123_, d_248_v23_, d_227_v7_, d_227_v7_, 101])})
            d_371_v126_ = (d_371_v126_ if not((d_233_v14_)[default__.safeIndex(390, (d_233_v14_).length(0))]) else (d_371_v126_ if (d_369_v124_).f29 else ((d_372_v127_)[d_368_v123_] if (d_368_v123_) in (d_372_v127_) else d_371_v126_)))
        d_373_v128_: _dafny.Seq
        d_374_v129_: _dafny.Array
        d_375_v130_: int
        out30_: _dafny.Seq
        out31_: _dafny.Array
        out32_: int
        out30_, out31_, out32_ = default__.m0(d_246_v21_, d_240_globalState_)
        d_373_v128_ = out30_
        d_374_v129_ = out31_
        d_375_v130_ = out32_
        d_292_v55_ = (d_292_v55_ if False else d_292_v55_)
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_233_v14_).length(0)):
            d_376_i12_: int = guard_loop_0_
            if (True) and (((0) <= (d_376_i12_)) and ((d_376_i12_) < ((d_233_v14_).length(0)))):
                (d_233_v14_)[(d_376_i12_)] = (d_292_v55_).f29
        if d_220_v0_:
            d_377_v131_: _dafny.Array
            nw61_ = _dafny.Array(_dafny.Array(None, 0), 29)
            d_377_v131_ = nw61_
            d_377_v131_ = (d_377_v131_ if d_220_v0_ else d_377_v131_)
            (d_240_globalState_).f13 = default__.safeDivisionInt(d_375_v130_, default__.safeModuloInt(d_248_v23_, d_227_v7_))
            (d_240_globalState_).f27 = d_375_v130_
            d_378_v132_: _dafny.Seq
            d_379_v133_: _dafny.Array
            d_380_v134_: int
            out33_: _dafny.Seq
            out34_: _dafny.Array
            out35_: int
            out33_, out34_, out35_ = default__.m0(d_246_v21_, d_240_globalState_)
            d_378_v132_ = out33_
            d_379_v133_ = out34_
            d_380_v134_ = out35_
            d_381_v135_: _dafny.Map
            d_381_v135_ = _dafny.Map({len(_dafny.Map({d_375_v130_: (d_292_v55_).f29})): (0) - (d_227_v7_)})
            d_382_v136_: _dafny.Set
            d_382_v136_ = _dafny.Set({d_227_v7_, (0) - (len(d_381_v135_))})
            d_383_v137_: _dafny.Map
            d_383_v137_ = _dafny.Map({d_248_v23_: len(d_382_v136_)})
            index47_ = default__.safeIndex(135, (d_229_v9_).length(0))
            (d_229_v9_)[index47_] = ((d_383_v137_)[d_248_v23_] if (d_248_v23_) in (d_383_v137_) else -935)
            index48_ = default__.safeIndex(135, (d_229_v9_).length(0))
            rhs43_ = d_220_v0_
            rhs44_ = -171
            rhs45_ = default__.safeDivisionInt(d_248_v23_, d_380_v134_)
            lhs32_ = d_229_v9_
            lhs33_ = default__.safeIndex(135, (d_229_v9_).length(0))
            d_220_v0_ = rhs43_
            d_375_v130_ = rhs44_
            lhs32_[lhs33_] = rhs45_
        elif True:
            d_384_v138_: _dafny.Map
            d_384_v138_ = _dafny.Map({d_237_v16_: d_220_v0_})
            d_385_v139_: _dafny.Seq
            d_385_v139_ = _dafny.SeqWithoutIsStrInference([d_227_v7_, len(_dafny.Set({395, d_227_v7_})), len(d_384_v138_), len(d_231_v12_), d_248_v23_])
            d_386_v140_: _dafny.Set
            d_386_v140_ = _dafny.Set({d_385_v139_})
            d_387_v141_: _dafny.MultiSet
            d_387_v141_ = _dafny.MultiSet([d_385_v139_, d_385_v139_])
            d_388_v143_: _dafny.Array
            nw62_ = _dafny.Array(None, 13)
            nw62_[int(0)] = d_386_v140_
            nw62_[int(1)] = d_386_v140_
            nw62_[int(2)] = (d_386_v140_).intersection(d_386_v140_)
            nw62_[int(3)] = default__.fm8(d_227_v7_, True, (d_292_v55_).f29, (d_292_v55_).f29, d_240_globalState_)
            nw62_[int(4)] = (d_386_v140_) | (d_386_v140_)
            def iife16_():
                coll4_ = _dafny.Set()
                compr_4_: _dafny.Seq
                for compr_4_ in (d_387_v141_).Elements:
                    d_389_v142_: _dafny.Seq = compr_4_
                    if (d_389_v142_) in (d_387_v141_):
                        coll4_ = coll4_.union(_dafny.Set([d_389_v142_]))
                return _dafny.Set(coll4_)
            nw62_[int(5)] = iife16_()
            
            nw62_[int(6)] = d_386_v140_
            nw62_[int(7)] = d_386_v140_
            nw62_[int(8)] = d_386_v140_
            nw62_[int(9)] = d_386_v140_
            nw62_[int(10)] = _dafny.Set({d_385_v139_, d_385_v139_})
            nw62_[int(11)] = (d_386_v140_) | (_dafny.Set({d_385_v139_, d_385_v139_}))
            nw62_[int(12)] = d_386_v140_
            d_388_v143_ = nw62_
            index49_ = default__.safeIndex(425, (d_388_v143_).length(0))
            (d_388_v143_)[index49_] = d_386_v140_
            index50_ = default__.safeIndex(425, (d_388_v143_).length(0))
            rhs46_ = not(d_220_v0_)
            rhs47_ = (d_386_v140_).intersection(default__.fm8(d_248_v23_, False, (d_292_v55_).f29, d_220_v0_, d_240_globalState_))
            lhs34_ = d_240_globalState_
            lhs35_ = d_388_v143_
            lhs36_ = default__.safeIndex(425, (d_388_v143_).length(0))
            lhs34_.f14 = rhs46_
            lhs35_[lhs36_] = rhs47_
            d_248_v23_ = d_248_v23_
            if (d_224_v4_)[default__.safeIndex(d_375_v130_, len(d_224_v4_))]:
                d_390_v144_: D1
                d_390_v144_ = D1_DC4((d_292_v55_).f29, d_375_v130_)
                d_220_v0_ = (d_390_v144_).cf5
                d_391_v145_: _dafny.Set
                d_391_v145_ = d_228_v8_
                d_392_v146_: _dafny.Seq
                d_392_v146_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_292_v55_).f29, (d_292_v55_).f29}), default__.fm9(d_248_v23_, d_240_globalState_)])
                d_393_v147_: _dafny.Array
                nw63_ = _dafny.Array(None, 18)
                nw63_[int(0)] = d_228_v8_
                nw63_[int(1)] = (_dafny.Set({d_220_v0_, (d_390_v144_).cf5}) if (d_292_v55_).f29 else d_228_v8_)
                nw63_[int(2)] = (d_228_v8_ if d_220_v0_ else d_228_v8_)
                nw63_[int(3)] = default__.fm9(d_248_v23_, d_240_globalState_)
                nw63_[int(4)] = (d_391_v145_)
                nw63_[int(5)] = d_228_v8_
                nw63_[int(6)] = _dafny.Set({d_220_v0_, (d_292_v55_).f29, (d_292_v55_).f29, (d_292_v55_).f29})
                nw63_[int(7)] = (d_228_v8_).intersection(d_228_v8_)
                nw63_[int(8)] = (_dafny.Set({(d_292_v55_).f29})).intersection(_dafny.Set({(d_292_v55_).f29, (d_292_v55_).f29}))
                nw63_[int(9)] = d_228_v8_
                nw63_[int(10)] = (d_392_v146_)[default__.safeIndex(d_227_v7_, len(d_392_v146_))]
                nw63_[int(11)] = (_dafny.Set({d_220_v0_, not((d_292_v55_).f29)})) - (default__.fm9(629, d_240_globalState_))
                nw63_[int(12)] = d_228_v8_
                nw63_[int(13)] = (d_228_v8_) - (d_228_v8_)
                nw63_[int(14)] = d_228_v8_
                nw63_[int(15)] = d_228_v8_
                nw63_[int(16)] = d_228_v8_
                nw63_[int(17)] = (d_228_v8_) | (d_228_v8_)
                d_393_v147_ = nw63_
                d_394_v148_: D1
                d_394_v148_ = D1_DC5(d_227_v7_, (d_292_v55_).f29, d_375_v130_, (d_292_v55_).f29)
                d_395_v149_: D1
                d_395_v149_ = D1_DC6(d_394_v148_)
                d_396_v150_: D2
                d_396_v150_ = D2_DC9(d_220_v0_, d_292_v55_)
                rhs48_ = d_393_v147_
                rhs49_ = (True) not in ((d_231_v12_) | (d_231_v12_))
                rhs50_ = default__.fm10(d_220_v0_, 620, (d_292_v55_).f29, d_240_globalState_)
                rhs51_ = (d_396_v150_).cf14
                lhs37_ = d_240_globalState_
                lhs38_ = d_240_globalState_
                d_393_v147_ = rhs48_
                lhs37_.f14 = rhs49_
                d_395_v149_ = rhs50_
                lhs38_.f14 = rhs51_
                index51_ = default__.safeIndex(937, (d_229_v9_).length(0))
                (d_229_v9_)[index51_] = d_248_v23_
                index52_ = default__.safeIndex(937, (d_229_v9_).length(0))
                (d_229_v9_)[index52_] = d_375_v130_
                (d_240_globalState_).f14 = (d_292_v55_).f29
                d_220_v0_ = (d_227_v7_) <= (d_227_v7_)
            elif True:
                d_397_v151_: _dafny.Array
                nw64_ = _dafny.Array(_dafny.Map({}), 2)
                d_397_v151_ = nw64_
                d_398_v152_: _dafny.MultiSet
                d_398_v152_ = _dafny.MultiSet([d_248_v23_])
                d_399_v153_: D2
                d_399_v153_ = D2_DC9((d_292_v55_).f29, d_292_v55_)
                d_400_v154_: D2
                d_400_v154_ = D2_DC9((d_292_v55_).f29, (d_399_v153_).cf15)
                d_401_v155_: _dafny.Map
                d_401_v155_ = _dafny.Map({not((default__.fm11(d_375_v130_, d_398_v152_, (d_292_v55_).f29, d_240_globalState_)).cf5): d_399_v153_})
                index53_ = default__.safeIndex(796, (d_397_v151_).length(0))
                (d_397_v151_)[index53_] = d_401_v155_
                index54_ = default__.safeIndex(796, (d_397_v151_).length(0))
                (d_397_v151_)[index54_] = d_401_v155_
                d_246_v21_ = ((d_246_v21_).set(default__.safeIndex(d_248_v23_, len(d_246_v21_)), d_237_v16_)) + (d_373_v128_)
                d_402_v156_: C0
                nw65_ = C0()
                nw65_.ctor__(d_220_v0_)
                d_402_v156_ = nw65_
                index55_ = default__.safeIndex(49, (d_247_v22_).length(0))
                (d_247_v22_)[index55_] = d_227_v7_
                d_403_v157_: _dafny.Array
                nw66_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_403_v157_ = nw66_
                index56_ = default__.safeIndex(105, (d_403_v157_).length(0))
                (d_403_v157_)[index56_] = d_233_v14_
                d_404_v158_: _dafny.Set
                d_404_v158_ = _dafny.Set({d_227_v7_, len(d_231_v12_), -263, (0) - (len(_dafny.Set({d_385_v139_, d_385_v139_}))), len(_dafny.Set({(0) - (d_248_v23_)}))})
                d_405_v159_: _dafny.Map
                d_405_v159_ = _dafny.Map({(d_402_v156_).f29: d_247_v22_})
                index57_ = default__.safeIndex(49, (d_247_v22_).length(0))
                index58_ = default__.safeIndex(105, (d_403_v157_).length(0))
                rhs52_ = (0) - ((d_375_v130_) * (len(d_404_v158_)))
                rhs53_ = ((d_405_v159_)[(len(d_246_v21_)) != (d_227_v7_)] if ((len(d_246_v21_)) != (d_227_v7_)) in (d_405_v159_) else d_247_v22_)
                rhs54_ = d_399_v153_
                rhs55_ = d_233_v14_
                rhs56_ = default__.fm0(d_223_v3_, (d_292_v55_).f29, d_240_globalState_)
                lhs39_ = d_247_v22_
                lhs40_ = default__.safeIndex(49, (d_247_v22_).length(0))
                lhs41_ = d_240_globalState_
                lhs42_ = d_403_v157_
                lhs43_ = default__.safeIndex(105, (d_403_v157_).length(0))
                lhs44_ = d_240_globalState_
                lhs39_[lhs40_] = rhs52_
                lhs41_.f28 = rhs53_
                d_400_v154_ = rhs54_
                lhs42_[lhs43_] = rhs55_
                lhs44_.f14 = rhs56_
                (d_240_globalState_).f17 = default__.safeDivisionInt(d_227_v7_, d_227_v7_)
            d_406_v160_: _dafny.Set
            d_406_v160_ = _dafny.Set({d_248_v23_})
            d_407_v161_: D1
            d_407_v161_ = D1_DC5(d_227_v7_, d_220_v0_, default__.safeDivisionInt(len(d_406_v160_), d_248_v23_), d_220_v0_)
            source7_ = d_407_v161_
            if source7_.is_DC3:
                d_408___mcc_h2_ = source7_.cf3
                d_409___mcc_h3_ = source7_.cf4
                d_410_cf4_ = d_409___mcc_h3_
                d_411_cf3_ = d_408___mcc_h2_
                (d_240_globalState_).f14 = not((d_246_v21_) < (d_246_v21_))
                d_412_v162_: _dafny.Seq
                d_412_v162_ = _dafny.SeqWithoutIsStrInference([d_223_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msee"))])
                d_413_v163_: _dafny.Set
                d_413_v163_ = _dafny.Set({d_374_v129_})
                d_414_v164_: D4
                d_414_v164_ = D4_DC12(d_229_v9_)
                rhs57_ = (d_248_v23_) - (len((d_412_v162_)[default__.safeIndex(d_248_v23_, len(d_412_v162_))]))
                rhs58_ = (d_413_v163_).ispropersubset((d_413_v163_) - (_dafny.Set({(d_414_v164_).cf19, d_229_v9_, d_374_v129_})))
                lhs45_ = d_240_globalState_
                lhs46_ = d_240_globalState_
                lhs45_.f13 = rhs57_
                lhs46_.f14 = rhs58_
                d_415_v165_: _dafny.MultiSet
                d_415_v165_ = _dafny.MultiSet([d_224_v4_])
                d_415_v165_ = ((d_415_v165_).set(d_224_v4_, default__.abs(d_248_v23_))) - (d_415_v165_)
                rhs59_ = (d_292_v55_).f29
                rhs60_ = (d_224_v4_) + ((d_224_v4_) + (d_224_v4_))
                lhs47_ = d_240_globalState_
                lhs47_.f14 = rhs59_
                d_224_v4_ = rhs60_
            elif source7_.is_DC4:
                d_416___mcc_h4_ = source7_.cf5
                d_417___mcc_h5_ = source7_.cf6
                d_418_cf6_ = d_417___mcc_h5_
                d_419_cf5_ = d_416___mcc_h4_
                d_420_v166_: _dafny.Seq
                d_420_v166_ = _dafny.SeqWithoutIsStrInference([d_223_v3_])
                (d_240_globalState_).f14 = not(not(not((d_373_v128_) not in (d_420_v166_))))
                (d_240_globalState_).f27 = (d_385_v139_)[default__.safeIndex(d_248_v23_, len(d_385_v139_))]
                d_421_v167_: C0
                nw67_ = C0()
                nw67_.ctor__(d_419_cf5_)
                d_421_v167_ = nw67_
                d_422_v168_: _dafny.Array
                nw68_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_422_v168_ = nw68_
                d_423_v169_: D5
                d_423_v169_ = D5_DC14(d_422_v168_)
                d_424_v170_: _dafny.Array
                nw69_ = _dafny.Array(None, 28)
                nw69_[int(0)] = d_422_v168_
                nw69_[int(1)] = d_422_v168_
                nw69_[int(2)] = d_422_v168_
                nw69_[int(3)] = d_422_v168_
                nw69_[int(4)] = d_422_v168_
                nw69_[int(5)] = d_422_v168_
                nw69_[int(6)] = d_422_v168_
                nw69_[int(7)] = d_422_v168_
                nw69_[int(8)] = d_422_v168_
                nw69_[int(9)] = d_422_v168_
                nw69_[int(10)] = d_422_v168_
                nw69_[int(11)] = d_422_v168_
                nw69_[int(12)] = d_422_v168_
                nw69_[int(13)] = d_422_v168_
                nw69_[int(14)] = d_422_v168_
                nw69_[int(15)] = d_422_v168_
                nw69_[int(16)] = d_422_v168_
                nw69_[int(17)] = d_422_v168_
                nw69_[int(18)] = d_422_v168_
                nw69_[int(19)] = (d_422_v168_ if (d_292_v55_).f29 else d_422_v168_)
                nw69_[int(20)] = d_422_v168_
                nw69_[int(21)] = d_422_v168_
                nw69_[int(22)] = d_422_v168_
                nw69_[int(23)] = (d_423_v169_).cf24
                nw69_[int(24)] = d_422_v168_
                nw69_[int(25)] = d_422_v168_
                nw69_[int(26)] = d_422_v168_
                nw69_[int(27)] = d_422_v168_
                d_424_v170_ = nw69_
                index59_ = default__.safeIndex(743, (d_424_v170_).length(0))
                (d_424_v170_)[index59_] = d_422_v168_
                index60_ = default__.safeIndex(743, (d_424_v170_).length(0))
                (d_424_v170_)[index60_] = (d_422_v168_ if (d_292_v55_).f29 else d_422_v168_)
            elif source7_.is_DC5:
                d_425___mcc_h6_ = source7_.cf7
                d_426___mcc_h7_ = source7_.cf8
                d_427___mcc_h8_ = source7_.cf9
                d_428___mcc_h9_ = source7_.cf10
                d_429_cf10_ = d_428___mcc_h9_
                d_430_cf9_ = d_427___mcc_h8_
                d_431_cf8_ = d_426___mcc_h7_
                d_432_cf7_ = d_425___mcc_h6_
                d_433_v171_: _dafny.Array
                nw70_ = _dafny.Array(None, 1)
                nw70_[int(0)] = d_292_v55_
                d_433_v171_ = nw70_
                index61_ = default__.safeIndex(687, (d_433_v171_).length(0))
                (d_433_v171_)[index61_] = d_292_v55_
                index62_ = default__.safeIndex(687, (d_433_v171_).length(0))
                (d_433_v171_)[index62_] = d_292_v55_
                d_434_v172_: D1
                d_434_v172_ = D1_DC2(d_246_v21_)
                d_435_v173_: _dafny.Seq
                d_436_v174_: _dafny.Array
                d_437_v175_: int
                out36_: _dafny.Seq
                out37_: _dafny.Array
                out38_: int
                out36_, out37_, out38_ = default__.m0((d_434_v172_).cf2, d_240_globalState_)
                d_435_v173_ = out36_
                d_436_v174_ = out37_
                d_437_v175_ = out38_
                (d_240_globalState_).f17 = (d_385_v139_)[default__.safeIndex(d_227_v7_, len(d_385_v139_))]
                d_438_v176_: _dafny.MultiSet
                d_438_v176_ = _dafny.MultiSet([d_237_v16_, d_237_v16_])
                d_439_v177_: _dafny.MultiSet
                d_439_v177_ = _dafny.MultiSet([default__.fm2((d_438_v176_).set(d_237_v16_, default__.abs(d_227_v7_)), d_430_cf9_, d_437_v175_, (d_292_v55_).f29, d_240_globalState_), len(d_385_v139_), d_430_cf9_])
                d_439_v177_ = d_439_v177_
            elif source7_.is_DC2:
                d_440___mcc_h10_ = source7_.cf2
                d_441_cf2_ = d_440___mcc_h10_
                d_442_v178_: _dafny.Seq
                d_443_v179_: _dafny.Array
                d_444_v180_: int
                out39_: _dafny.Seq
                out40_: _dafny.Array
                out41_: int
                out39_, out40_, out41_ = default__.m0((d_441_cf2_).set(default__.safeIndex(-918, len(d_441_cf2_)), (d_373_v128_)[default__.safeIndex(-862, len(d_373_v128_))]), d_240_globalState_)
                d_442_v178_ = out39_
                d_443_v179_ = out40_
                d_444_v180_ = out41_
                (d_240_globalState_).f11 = ((_dafny.MultiSet([(d_292_v55_).f29])) | ((d_221_v1_).set(d_220_v0_, default__.abs(d_444_v180_)))).intersection(d_221_v1_)
                index63_ = default__.safeIndex(806, (d_233_v14_).length(0))
                (d_233_v14_)[index63_] = (d_292_v55_).f29
                index64_ = default__.safeIndex(806, (d_233_v14_).length(0))
                (d_233_v14_)[index64_] = not(d_220_v0_)
                d_445_v181_: _dafny.MultiSet
                d_445_v181_ = _dafny.MultiSet([-929, d_227_v7_])
                d_446_v182_: D6
                d_446_v182_ = D6_DC16(d_445_v181_)
                index65_ = default__.safeIndex(806, (d_233_v14_).length(0))
                rhs61_ = d_292_v55_
                rhs62_ = ((d_224_v4_).set(default__.safeIndex(d_375_v130_, len(d_224_v4_)), True) if (d_233_v14_)[default__.safeIndex(806, (d_233_v14_).length(0))] else d_224_v4_)
                rhs63_ = not((d_445_v181_).issubset((d_446_v182_).cf30))
                rhs64_ = d_224_v4_
                lhs48_ = d_233_v14_
                lhs49_ = default__.safeIndex(806, (d_233_v14_).length(0))
                d_292_v55_ = rhs61_
                d_224_v4_ = rhs62_
                lhs48_[lhs49_] = rhs63_
                d_224_v4_ = rhs64_
            elif True:
                d_447___mcc_h11_ = source7_.cf11
                d_448_cf11_ = d_447___mcc_h11_
                d_449_v183_: _dafny.Seq
                d_450_v184_: _dafny.Array
                d_451_v185_: int
                out42_: _dafny.Seq
                out43_: _dafny.Array
                out44_: int
                out42_, out43_, out44_ = default__.m0((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynbc"))).set(default__.safeIndex(d_227_v7_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynbc")))), d_237_v16_), d_240_globalState_)
                d_449_v183_ = out42_
                d_450_v184_ = out43_
                d_451_v185_ = out44_
                d_236_v15_ = d_236_v15_
                d_452_v186_: _dafny.Seq
                d_453_v187_: _dafny.Array
                d_454_v188_: int
                out45_: _dafny.Seq
                out46_: _dafny.Array
                out47_: int
                out45_, out46_, out47_ = default__.m0((d_246_v21_) + (d_246_v21_), d_240_globalState_)
                d_452_v186_ = out45_
                d_453_v187_ = out46_
                d_454_v188_ = out47_
                d_455_v189_: _dafny.Map
                d_455_v189_ = _dafny.Map({d_451_v185_: d_374_v129_})
                (d_240_globalState_).f14 = default__.fm0((d_449_v183_) + (default__.fm12(d_227_v7_, len(d_455_v189_), (d_292_v55_).f29, d_240_globalState_)), (d_292_v55_).f29, d_240_globalState_)
            (d_240_globalState_).f27 = (0) - ((d_375_v130_) * (len(d_385_v139_)))
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_233_v14_).length(0)):
            d_456_i13_: int = guard_loop_1_
            if (True) and (((0) <= (d_456_i13_)) and ((d_456_i13_) < ((d_233_v14_).length(0)))):
                (d_233_v14_)[(d_456_i13_)] = (d_292_v55_).f29
        d_220_v0_ = (d_292_v55_).f29
        _dafny.print(_dafny.string_of(d_220_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_v1_) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_222_v2_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True]), _dafny.MultiSet([True, True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_223_v3_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_224_v4_) == (_dafny.SeqWithoutIsStrInference([True, True, True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_225_v5_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([True, True, True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_227_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v8_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v9_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_230_v10_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v12_) == (_dafny.Map({True: False, False: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_v13_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: False}), _dafny.Map({True: False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v14_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v15_) == (_dafny.Map({155: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_237_v16_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v17_) == (_dafny.Set({_dafny.CodePoint('a')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_239_v18_) == (_dafny.Map({_dafny.Map({155: True}): _dafny.Set({_dafny.CodePoint('a')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f0) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True]), _dafny.MultiSet([True, True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_240_globalState_).f6).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f7) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([True, True, True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_240_globalState_).f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f11) == (_dafny.MultiSet([True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f18).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f19) == (_dafny.Map({_dafny.Map({True: False}): 773}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f20)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f24) == (_dafny.Map({_dafny.Map({155: True}): _dafny.Set({_dafny.CodePoint('a')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f26) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_.f28)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_242_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_246_v21_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v22_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_248_v23_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_272_v36_)[2]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_273_v37_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v55_).f29))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_373_v128_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v129_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_375_v130_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1)
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({self.cf2.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D2_DC10)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({self.cf13.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC10(D2, NamedTuple('DC10', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC10({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)

class D3_DC11(D3, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(None, _dafny.CodePoint('D'), _dafny.CodePoint('D'), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC13(D4, NamedTuple('DC13', [('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC15(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, False, False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC15(D5, NamedTuple('DC15', [('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({self.cf25.VerbatimString(True)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(_dafny.CodePoint('D'), False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC17(D6, NamedTuple('DC17', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({self.cf34.VerbatimString(True)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC20(False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC20(D7, NamedTuple('DC20', [('cf39', Any), ('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)

class D8_DC23(D8, NamedTuple('DC23', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f2: str = _dafny.CodePoint('D')
        self.f11: _dafny.MultiSet = _dafny.MultiSet({})
        self.f13: int = int(0)
        self.f14: bool = False
        self.f17: int = int(0)
        self.f25: int = int(0)
        self.f27: int = int(0)
        self.f28: _dafny.Array = _dafny.Array(None, 0)
        self._f0: _dafny.Seq = _dafny.Seq({})
        self._f1: int = int(0)
        self._f3: int = int(0)
        self._f4: bool = False
        self._f5: int = int(0)
        self._f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f7: _dafny.Map = _dafny.Map({})
        self._f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f10: int = int(0)
        self._f12: int = int(0)
        self._f15: int = int(0)
        self._f16: int = int(0)
        self._f18: _dafny.MultiSet = _dafny.MultiSet({})
        self._f19: _dafny.Map = _dafny.Map({})
        self._f20: _dafny.Array = _dafny.Array(None, 0)
        self._f21: int = int(0)
        self._f22: bool = False
        self._f23: int = int(0)
        self._f24: _dafny.Map = _dafny.Map({})
        self._f26: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27, f28):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
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
        (self).f13 = f13
        (self).f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self).f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23
        (self)._f24 = f24
        (self).f25 = f25
        (self)._f26 = f26
        (self).f27 = f27
        (self).f28 = f28

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
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
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
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
    def f26(self):
        return self._f26

class C0:
    def  __init__(self):
        self._f29: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f29):
        (self)._f29 = f29

    @property
    def f29(self):
        return self._f29
