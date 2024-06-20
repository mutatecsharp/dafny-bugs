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
        return (True) == (not(not((True if True else False))))

    @staticmethod
    def fm1(p0, p1, globalState):
        source0_ = D10_DC23(len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False])), 11, (_dafny.MultiSet([473, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwihgjxua")))])).cardinality, (0) - (len(_dafny.SeqWithoutIsStrInference([-810, len(_dafny.SeqWithoutIsStrInference([True, False]))])))})), not(True), True)
        if source0_.is_DC23:
            d_0___mcc_h0_ = source0_.cf33
            d_1___mcc_h1_ = source0_.cf34
            d_2___mcc_h2_ = source0_.cf35
            d_3_cf35_ = d_2___mcc_h2_
            d_4_cf34_ = d_1___mcc_h1_
            d_5_cf33_ = d_0___mcc_h0_
            return len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_5_cf33_, 841])): d_4_cf34_}))
        elif True:
            d_6___mcc_h3_ = source0_.cf32
            d_7_cf32_ = d_6___mcc_h3_
            if (d_7_cf32_).f16:
                return len(_dafny.SeqWithoutIsStrInference([-846]))
            elif True:
                return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhf")))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({933, len(_dafny.Map({200: True}))})).intersection(_dafny.Set({len(_dafny.Map({-918: not(False)}))}))) | (_dafny.Set({-219, len(_dafny.SeqWithoutIsStrInference([False, not(not(True))])), 623, 881}))

    @staticmethod
    def fm10(p0, globalState):
        def lambda0_(source1_):
            if source1_.is_DC2:
                d_8___mcc_h0_ = source1_.cf2
                d_9_cf2_ = d_8___mcc_h0_
                def iife0_():
                    coll0_ = _dafny.Map()
                    compr_0_: int
                    for compr_0_ in _dafny.IntegerRange(892, 185):
                        d_11_v0_: int = compr_0_
                        if ((892) <= (d_11_v0_)) and ((d_11_v0_) < (185)):
                            coll0_[(d_11_v0_) + (712)] = len(_dafny.SeqWithoutIsStrInference([d_9_cf2_, d_9_cf2_]))
                    return _dafny.Map(coll0_)
                return (_dafny.MultiSet([-14, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_10_i0_ in range(default__.abs(541))])), 958, len(iife0_()
                )])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-701 for d_12_i1_ in range(default__.abs(-142))])))
            elif source1_.is_DC3:
                d_13___mcc_h1_ = source1_.cf3
                d_14___mcc_h2_ = source1_.cf4
                d_15_cf4_ = d_14___mcc_h2_
                d_16_cf3_ = d_13___mcc_h1_
                return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([944 for d_17_i2_ in range(default__.abs(911))])) + (_dafny.SeqWithoutIsStrInference([290, 674])))
            elif True:
                d_18___mcc_h3_ = source1_.cf1
                d_19_cf1_ = d_18___mcc_h3_
                return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_19_cf1_, d_19_cf1_, True]))]))) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_19_cf1_]))]))

        return (lambda0_(D1_DC3(not(not(False)), 18))).cardinality

    @staticmethod
    def fm11(globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(not(False)), not(not(True))])) | (_dafny.MultiSet([True])) for d_20_i0_ in range(default__.abs(756))])

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([False])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm15(globalState):
        return ((_dafny.Map({True: D3_DC8(True, 257, 9)})) | (_dafny.Map({True: D3_DC8(True, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True}))])), len(_dafny.SeqWithoutIsStrInference([True])))}))) | (_dafny.Map({True: D3_DC8(True, len(_dafny.SeqWithoutIsStrInference([True])), (0) - (-201))}))

    @staticmethod
    def fm16(globalState):
        return D3_DC8(False, (633) * ((D15_DC39(True, len(_dafny.Map({(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([D8_DC18(), D8_DC18()]), _dafny.SeqWithoutIsStrInference([D8_DC18()])])).cardinality: True})), not(not(False)), 981, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([263])) for d_21_i0_ in range(default__.abs(470))])))).cf59), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lriyttbrc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jftuf")))))

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        return default__.safeModuloInt(((_dafny.MultiSet([False, False, False, True, True])).cardinality) + (len(_dafny.SeqWithoutIsStrInference([(0) - ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True}))]))))) for d_22_i0_ in range(default__.abs(-45))]))), (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('j')]))) + (271))

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([False, not((-922) > (583)), (False) or (not(False))])

    @staticmethod
    def fm19(p0, globalState):
        return ((D15_DC39(False, len(_dafny.Map({False: not(True)})), True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cynnxklv"))), -334)).cf59) + (-349)

    @staticmethod
    def fm20(globalState):
        if not (not(not(False))) or (False):
            def iife1_():
                def iife3_():
                    coll3_ = _dafny.Map()
                    compr_3_: int
                    for compr_3_ in _dafny.IntegerRange(650, 968):
                        d_23_v1_: int = compr_3_
                        if ((650) <= (d_23_v1_)) and ((d_23_v1_) < (968)):
                            coll3_[default__.safeDivisionInt(d_23_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubretlxf"))))] = not(True)
                    return _dafny.Map(coll3_)
                coll1_ = _dafny.Map()
                def iife2_():
                    coll2_ = _dafny.Map()
                    compr_2_: int
                    for compr_2_ in _dafny.IntegerRange(650, 968):
                        d_23_v1_: int = compr_2_
                        if ((650) <= (d_23_v1_)) and ((d_23_v1_) < (968)):
                            coll2_[default__.safeDivisionInt(d_23_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubretlxf"))))] = not(True)
                    return _dafny.Map(coll2_)
                compr_1_: int
                for compr_1_ in (iife2_()
                ).keys.Elements:
                    d_24_v0_: int = compr_1_
                    if (d_24_v0_) in (iife3_()
                    ):
                        coll1_[default__.safeModuloInt(d_24_v0_, len(_dafny.SeqWithoutIsStrInference([False])))] = _dafny.SeqWithoutIsStrInference([809])
                return _dafny.Map(coll1_)
            return iife1_()
            
        elif True:
            return _dafny.Map({len(_dafny.Set({not(False), not(False)})): _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, True, True, False])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjccwlgyi"))), 12])})

    @staticmethod
    def fm21(p0, p1, p2, globalState):
        return _dafny.MultiSet([not((not(False) if True else True)), False, ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_25_i0_ in range(default__.abs(309))]))) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_26_i1_ in range(default__.abs(613))]))])

    @staticmethod
    def fm22(globalState):
        return (_dafny.Map({not(False): not(False)})) | (_dafny.Map({False: not(not(False))}))

    @staticmethod
    def fm23(p0, p1, globalState):
        return _dafny.Set({(-93) != ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, not(False)]))).cardinality), (_dafny.Set({235, 10})).issubset(_dafny.Set({-751}))})

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(188, 665):
                d_27_v0_: int = compr_4_
                if ((188) <= (d_27_v0_)) and ((d_27_v0_) < (665)):
                    coll4_[(d_27_v0_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_29_i0_ in range(default__.abs(-960))])))] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_28_i1_ in range(default__.abs(954))]))
            return _dafny.Map(coll4_)
        source2_ = D3_DC6(len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ap")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mthdwg")))])).cardinality, -666])), True, False, len(iife4_()
), False)
        if source2_.is_DC6:
            d_30___mcc_h0_ = source2_.cf7
            d_31___mcc_h1_ = source2_.cf8
            d_32___mcc_h2_ = source2_.cf9
            d_33___mcc_h3_ = source2_.cf10
            d_34___mcc_h4_ = source2_.cf11
            d_35_cf11_ = d_34___mcc_h4_
            d_36_cf10_ = d_33___mcc_h3_
            d_37_cf9_ = d_32___mcc_h2_
            d_38_cf8_ = d_31___mcc_h1_
            d_39_cf7_ = d_30___mcc_h0_
            return _dafny.MultiSet([608])
        elif source2_.is_DC7:
            d_40___mcc_h5_ = source2_.cf12
            d_41_cf12_ = d_40___mcc_h5_
            return (_dafny.MultiSet([d_41_cf12_, d_41_cf12_])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([321, d_41_cf12_])))
        elif source2_.is_DC8:
            d_42___mcc_h6_ = source2_.cf13
            d_43___mcc_h7_ = source2_.cf14
            d_44___mcc_h8_ = source2_.cf15
            d_45_cf15_ = d_44___mcc_h8_
            d_46_cf14_ = d_43___mcc_h7_
            d_47_cf13_ = d_42___mcc_h6_
            return _dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference([(D14_DC36(d_45_cf15_, d_47_cf13_, d_47_cf13_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqwuhixo"))), d_47_cf13_)).cf51])))])
        elif True:
            d_48___mcc_h9_ = source2_.cf6
            d_49_cf6_ = d_48___mcc_h9_
            return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wiaxnjh")))])

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmhrvhsnu")))) + (207) for d_50_i0_ in range(default__.abs(841))])

    @staticmethod
    def fm26(p0, globalState):
        source3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
        d_51___mcc_h0_ = source3_
        d_52_cf20_ = d_51___mcc_h0_
        return _dafny.CodePoint('w')

    @staticmethod
    def fm27(globalState):
        return (_dafny.Map({-319: not(True)})) | (_dafny.Map({(0) - ((_dafny.MultiSet([len(_dafny.Map({(_dafny.MultiSet([True])).cardinality: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_53_i0_ in range(default__.abs(36))])}))])).cardinality): False}))

    @staticmethod
    def fm28(p0, p1, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in ((D17_DC44(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('v')])).cardinality])), 728]))).cf70).Elements:
                d_54_v0_: int = compr_5_
                if (d_54_v0_) in ((D17_DC44(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('v')])).cardinality])), 728]))).cf70):
                    coll5_[(d_54_v0_) + (866)] = True
            return _dafny.Map(coll5_)
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "butfuvwqw"))): _dafny.Map({170: True})})) | (_dafny.Map({(_dafny.MultiSet([False, not(False), False, True])).cardinality: iife5_()
        }))) | (_dafny.Map({(0) - ((0) - (-349)): _dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(D9_DC20(True, _dafny.CodePoint('e'))).cf29, True])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwvavjx")))})): False})}))

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        source4_ = D8_DC18()
        if source4_.is_DC18:
            return D1_DC1(True)
        elif True:
            d_55___mcc_h0_ = source4_.cf27
            d_56_cf27_ = d_55___mcc_h0_
            return D1_DC1(True)

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        return D9_DC20((True if not(False) else not(False)), _dafny.CodePoint('j'))

    @staticmethod
    def m0(globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        r3: _dafny.Map = _dafny.Map({})
        d_57_v0_: bool
        d_57_v0_ = False
        d_58_v1_: _dafny.Array
        nw0_ = _dafny.Array(False, 26)
        d_58_v1_ = nw0_
        d_59_v2_: _dafny.Array
        nw1_ = _dafny.Array(None, 6)
        nw1_[int(0)] = d_58_v1_
        nw1_[int(1)] = d_58_v1_
        nw1_[int(2)] = d_58_v1_
        nw1_[int(3)] = d_58_v1_
        nw1_[int(4)] = d_58_v1_
        nw1_[int(5)] = d_58_v1_
        d_59_v2_ = nw1_
        d_60_v3_: D3
        d_60_v3_ = D3_DC5(d_59_v2_)
        d_61_v4_: C1
        nw2_ = C1()
        nw2_.ctor__((d_57_v0_ if d_57_v0_ else d_57_v0_), d_60_v3_)
        d_61_v4_ = nw2_
        d_62_v5_: int
        d_62_v5_ = 188
        if (d_57_v0_ if d_57_v0_ else default__.fm0(d_62_v5_, d_62_v5_, _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_61_v4_).f16]) for d_63_i0_ in range(default__.abs(635))]), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkrihtkwb")): not((d_61_v4_).f16)}), globalState)):
            d_64_v6_: _dafny.MultiSet
            d_64_v6_ = _dafny.MultiSet([d_57_v0_, d_57_v0_])
            d_65_v7_: _dafny.Seq
            d_65_v7_ = _dafny.SeqWithoutIsStrInference([d_64_v6_])
            d_66_v8_: _dafny.Seq
            d_66_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmqmsbnsb"))
            d_67_v9_: str
            d_67_v9_ = _dafny.CodePoint('d')
            d_68_v10_: _dafny.Map
            d_68_v10_ = _dafny.Map({d_66_v8_: (d_61_v4_).fm4(d_62_v5_, True, d_67_v9_, globalState)})
            if ((False if (d_61_v4_).f16 else default__.fm0(d_62_v5_, d_62_v5_, d_65_v7_, d_68_v10_, globalState))) == ((d_61_v4_).f16):
                r1 = not(d_57_v0_)
                d_58_v1_ = d_58_v1_
                index0_ = default__.safeIndex(729, (d_58_v1_).length(0))
                (d_58_v1_)[index0_] = d_57_v0_
                d_69_v11_: _dafny.Seq
                d_69_v11_ = _dafny.SeqWithoutIsStrInference([(d_61_v4_).f16])
                index1_ = default__.safeIndex(729, (d_58_v1_).length(0))
                rhs0_ = d_67_v9_
                rhs1_ = (d_69_v11_)[default__.safeIndex(d_62_v5_, len(d_69_v11_))]
                rhs2_ = (d_61_v4_).f16
                lhs0_ = d_58_v1_
                lhs1_ = default__.safeIndex(729, (d_58_v1_).length(0))
                d_67_v9_ = rhs0_
                lhs0_[lhs1_] = rhs1_
                d_57_v0_ = rhs2_
                index2_ = default__.safeIndex(729, (d_58_v1_).length(0))
                (d_58_v1_)[index2_] = (d_61_v4_).f16
                d_70_v12_: _dafny.Array
                nw3_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
                d_70_v12_ = nw3_
                d_70_v12_ = d_70_v12_
            elif True:
                r2 = not (not((d_61_v4_).f16)) or ((d_61_v4_).f16)
                d_71_v13_: _dafny.Map
                d_71_v13_ = _dafny.Map({(0) - (d_62_v5_): (d_61_v4_).f16})
                d_72_v14_: _dafny.Seq
                d_72_v14_ = _dafny.SeqWithoutIsStrInference([len(d_71_v13_)])
                d_73_v15_: _dafny.Set
                d_73_v15_ = _dafny.Set({d_62_v5_, len(d_72_v14_), d_62_v5_})
                d_74_v16_: _dafny.Map
                d_74_v16_ = _dafny.Map({d_58_v1_: d_66_v8_})
                d_75_v17_: _dafny.Map
                d_75_v17_ = _dafny.Map({len(d_73_v15_): d_74_v16_})
                d_75_v17_ = (d_75_v17_).set(184, d_74_v16_)
                d_68_v10_ = (d_68_v10_).set(_dafny.SeqWithoutIsStrInference([d_67_v9_ for d_76_i1_ in range(default__.abs(592))]), (d_61_v4_).f16)
                d_77_v18_: _dafny.MultiSet
                d_77_v18_ = _dafny.MultiSet([d_67_v9_])
                d_78_v19_: D15
                d_78_v19_ = D15_DC39((d_61_v4_).f16, (d_77_v18_).cardinality, False, len(_dafny.SeqWithoutIsStrInference([d_67_v9_ for d_79_i2_ in range(default__.abs(768))])), d_62_v5_)
                d_80_v20_: _dafny.Map
                d_80_v20_ = _dafny.Map({d_62_v5_: d_78_v19_})
                d_80_v20_ = (d_80_v20_).set((d_62_v5_) * (d_62_v5_), d_78_v19_)
                (globalState).f2 = d_62_v5_
            d_81_v21_: _dafny.Array
            def lambda1_(d_82_v5_):
                def lambda2_(d_83_i3_):
                    return (d_83_i3_) * (d_82_v5_)

                return lambda2_

            init0_ = lambda1_(d_62_v5_)
            nw4_ = _dafny.Array(None, 3)
            for i0_0_ in range(nw4_.length(0)):
                nw4_[i0_0_] = init0_(i0_0_)
            d_81_v21_ = nw4_
            index3_ = default__.safeIndex(731, (d_81_v21_).length(0))
            (d_81_v21_)[index3_] = default__.safeModuloInt(127, d_62_v5_)
            index4_ = default__.safeIndex(731, (d_81_v21_).length(0))
            (d_81_v21_)[index4_] = default__.safeModuloInt(d_62_v5_, d_62_v5_)
            r1 = (d_61_v4_).f16
            if (d_61_v4_).f16:
                (globalState).f4 = d_81_v21_
                index5_ = default__.safeIndex(678, (d_59_v2_).length(0))
                (d_59_v2_)[index5_] = d_58_v1_
                index6_ = default__.safeIndex(678, (d_59_v2_).length(0))
                (d_59_v2_)[index6_] = d_58_v1_
                index7_ = default__.safeIndex(731, (d_81_v21_).length(0))
                (d_81_v21_)[index7_] = (d_81_v21_)[default__.safeIndex(731, (d_81_v21_).length(0))]
                d_84_v22_: _dafny.Map
                d_84_v22_ = _dafny.Map({d_57_v0_: (d_64_v6_).cardinality})
                d_85_v23_: _dafny.Set
                d_85_v23_ = _dafny.Set({len(d_66_v8_)})
                d_86_v24_: _dafny.Seq
                d_86_v24_ = _dafny.SeqWithoutIsStrInference([d_85_v23_, d_85_v23_])
                d_87_v25_: C0
                nw5_ = C0()
                nw5_.ctor__(default__.fm0(default__.fm19(len(d_84_v22_), globalState), (d_81_v21_)[default__.safeIndex(731, (d_81_v21_).length(0))], d_65_v7_, d_68_v10_, globalState), (len(d_86_v24_)) - (d_62_v5_))
                d_87_v25_ = nw5_
                d_88_v26_: T0
                nw6_ = C2()
                nw6_.ctor__(d_57_v0_, ((d_87_v25_).f12) and (True))
                d_88_v26_ = nw6_
                arr0_ = (d_59_v2_)[default__.safeIndex(678, (d_59_v2_).length(0))]
                index8_ = default__.safeIndex(136, ((d_59_v2_)[default__.safeIndex(678, (d_59_v2_).length(0))]).length(0))
                arr0_[index8_] = (d_61_v4_).f16
                index9_ = default__.safeIndex(731, (d_81_v21_).length(0))
                arr1_ = (d_59_v2_)[default__.safeIndex(678, (d_59_v2_).length(0))]
                index10_ = default__.safeIndex(136, ((d_59_v2_)[default__.safeIndex(678, (d_59_v2_).length(0))]).length(0))
                rhs3_ = (d_81_v21_)[default__.safeIndex(731, (d_81_v21_).length(0))]
                rhs4_ = d_87_v25_
                rhs5_ = d_88_v26_
                rhs6_ = not((d_61_v4_).f16)
                rhs7_ = (d_88_v26_).fm4(994, not(False), _dafny.CodePoint('h'), globalState)
                lhs2_ = d_81_v21_
                lhs3_ = default__.safeIndex(731, (d_81_v21_).length(0))
                lhs4_ = (d_59_v2_)[default__.safeIndex(678, (d_59_v2_).length(0))]
                lhs5_ = default__.safeIndex(136, ((d_59_v2_)[default__.safeIndex(678, (d_59_v2_).length(0))]).length(0))
                lhs2_[lhs3_] = rhs3_
                d_87_v25_ = rhs4_
                d_88_v26_ = rhs5_
                r1 = rhs6_
                lhs4_[lhs5_] = rhs7_
                d_89_v27_: _dafny.Array
                nw7_ = _dafny.Array(_dafny.Map({}), 23)
                d_89_v27_ = nw7_
                index11_ = default__.safeIndex(852, (d_89_v27_).length(0))
                (d_89_v27_)[index11_] = d_68_v10_
                d_90_v28_: _dafny.Seq
                d_90_v28_ = _dafny.SeqWithoutIsStrInference([(d_87_v25_).f13, d_62_v5_])
                index12_ = default__.safeIndex(731, (d_81_v21_).length(0))
                index13_ = default__.safeIndex(852, (d_89_v27_).length(0))
                index14_ = default__.safeIndex(731, (d_81_v21_).length(0))
                rhs8_ = ((d_81_v21_)[default__.safeIndex(731, (d_81_v21_).length(0))] if (d_87_v25_).f12 else (d_81_v21_)[default__.safeIndex(731, (d_81_v21_).length(0))])
                rhs9_ = d_68_v10_
                rhs10_ = ((d_61_v4_).f16 if (default__.fm0((d_81_v21_)[default__.safeIndex(731, (d_81_v21_).length(0))], d_62_v5_, d_65_v7_, d_68_v10_, globalState) if False else (d_61_v4_).f16) else not((d_62_v5_) < ((d_81_v21_)[default__.safeIndex(731, (d_81_v21_).length(0))])))
                rhs11_ = ((d_87_v25_).f13) + ((len(_dafny.SeqWithoutIsStrInference([d_67_v9_ for d_91_i4_ in range(default__.abs(204))]))) - ((d_81_v21_)[default__.safeIndex(731, (d_81_v21_).length(0))]))
                rhs12_ = ((d_81_v21_)[default__.safeIndex(731, (d_81_v21_).length(0))]) + ((d_90_v28_)[default__.safeIndex((0) - ((d_87_v25_).f13), len(d_90_v28_))])
                lhs6_ = d_81_v21_
                lhs7_ = default__.safeIndex(731, (d_81_v21_).length(0))
                lhs8_ = d_89_v27_
                lhs9_ = default__.safeIndex(852, (d_89_v27_).length(0))
                lhs10_ = d_81_v21_
                lhs11_ = default__.safeIndex(731, (d_81_v21_).length(0))
                lhs6_[lhs7_] = rhs8_
                lhs8_[lhs9_] = rhs9_
                d_57_v0_ = rhs10_
                lhs10_[lhs11_] = rhs11_
                d_62_v5_ = rhs12_
            elif True:
                d_66_v8_ = d_66_v8_
                d_92_v29_: _dafny.Set
                d_92_v29_ = _dafny.Set({d_58_v1_})
                d_93_v30_: _dafny.Seq
                d_93_v30_ = _dafny.SeqWithoutIsStrInference([D7_DC14(d_92_v29_)])
                d_93_v30_ = d_93_v30_
                d_94_v31_: C2
                nw8_ = C2()
                nw8_.ctor__((not(False) if (d_61_v4_).fm4(567, (d_61_v4_).f16, d_67_v9_, globalState) else d_57_v0_), True)
                d_94_v31_ = nw8_
                d_94_v31_ = d_94_v31_
                d_95_v32_: _dafny.Seq
                d_95_v32_ = _dafny.SeqWithoutIsStrInference([d_66_v8_])
                d_96_v33_: C1
                nw9_ = C1()
                nw9_.ctor__((d_66_v8_) < ((d_95_v32_)[default__.safeIndex(len(d_66_v8_), len(d_95_v32_))]), d_61_v4_.f17)
                d_96_v33_ = nw9_
                d_67_v9_ = d_67_v9_
            d_97_v34_: _dafny.Map
            d_97_v34_ = _dafny.Map({d_57_v0_: ((d_81_v21_)[default__.safeIndex(731, (d_81_v21_).length(0))]) >= (d_62_v5_)})
            d_97_v34_ = (d_97_v34_).set(not ((d_61_v4_).f16) or (d_57_v0_), not (d_57_v0_) or ((d_61_v4_).f16))
        elif True:
            d_98_v35_: C2
            nw10_ = C2()
            nw10_.ctor__((d_61_v4_).f16, (d_61_v4_).f16)
            d_98_v35_ = nw10_
            d_98_v35_ = d_98_v35_
            d_99_v36_: _dafny.Map
            d_99_v36_ = _dafny.Map({(d_61_v4_).f16: d_98_v35_.f14})
            d_100_v37_: D9
            d_100_v37_ = D9_DC20(not(((d_99_v36_)[(d_98_v35_).f15] if ((d_98_v35_).f15) in (d_99_v36_) else (d_98_v35_).f15)), _dafny.CodePoint('t'))
            d_101_v38_: _dafny.Seq
            d_101_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kplyjx"))
            d_102_v39_: _dafny.Seq
            d_102_v39_ = _dafny.SeqWithoutIsStrInference([647, d_62_v5_])
            d_103_v40_: D16
            d_103_v40_ = D16_DC42((d_102_v39_)[default__.safeIndex(36, len(d_102_v39_))], (d_61_v4_).f16)
            d_104_v41_: str
            d_104_v41_ = _dafny.CodePoint('a')
            d_105_v42_: _dafny.Map
            d_105_v42_ = _dafny.Map({d_62_v5_: _dafny.SeqWithoutIsStrInference([d_104_v41_ for d_106_i5_ in range(default__.abs(571))])})
            rhs13_ = not(False)
            rhs14_ = default__.fm30((d_103_v40_).cf65, 210, d_104_v41_, d_62_v5_, globalState)
            rhs15_ = ((d_105_v42_)[(0) - (len(d_101_v38_))] if ((0) - (len(d_101_v38_))) in (d_105_v42_) else default__.fm2(d_98_v35_.f14, d_62_v5_, 875, True, globalState))
            rhs16_ = ((970) - (d_62_v5_)) < ((0) - (d_62_v5_))
            lhs12_ = d_98_v35_
            r1 = rhs13_
            d_100_v37_ = rhs14_
            d_101_v38_ = rhs15_
            lhs12_.f14 = rhs16_
            r1 = (d_61_v4_).f16
            if d_57_v0_:
                d_107_v43_: T0
                nw11_ = C1()
                nw11_.ctor__((d_61_v4_).f16, d_61_v4_.f17)
                d_107_v43_ = nw11_
                d_108_v44_: _dafny.Map
                d_108_v44_ = _dafny.Map({d_101_v38_: default__.safeDivisionInt(d_62_v5_, d_62_v5_)})
                d_109_v45_: _dafny.Seq
                d_109_v45_ = _dafny.SeqWithoutIsStrInference([d_107_v43_, d_107_v43_, d_107_v43_])
                d_110_v47_: _dafny.Set
                d_110_v47_ = _dafny.Set({d_101_v38_})
                def iife6_():
                    coll6_ = _dafny.Map()
                    compr_6_: _dafny.Seq
                    for compr_6_ in (d_110_v47_).Elements:
                        d_111_v46_: _dafny.Seq = compr_6_
                        if (d_111_v46_) in (d_110_v47_):
                            coll6_[d_111_v46_] = (0) - (len(_dafny.Set({(0) - (d_62_v5_)})))
                    return _dafny.Map(coll6_)
                rhs17_ = (d_109_v45_)[default__.safeIndex(d_62_v5_, len(d_109_v45_))]
                rhs18_ = d_62_v5_
                rhs19_ = d_60_v3_
                rhs20_ = d_62_v5_
                rhs21_ = (iife6_()
                ).set(_dafny.SeqWithoutIsStrInference([d_104_v41_ for d_112_i6_ in range(default__.abs(-734))]), d_62_v5_)
                lhs13_ = globalState
                lhs14_ = globalState
                d_107_v43_ = rhs17_
                lhs13_.f2 = rhs18_
                d_60_v3_ = rhs19_
                lhs14_.f2 = rhs20_
                d_108_v44_ = rhs21_
                d_113_v48_: _dafny.MultiSet
                d_113_v48_ = _dafny.MultiSet([d_107_v43_])
                d_114_v49_: _dafny.Seq
                d_114_v49_ = _dafny.SeqWithoutIsStrInference([(d_61_v4_).f16, (d_98_v35_).fm12(True, ((d_113_v48_)[d_107_v43_] if (d_107_v43_) in (d_113_v48_) else d_62_v5_), _dafny.MultiSet([d_62_v5_]), len(_dafny.SeqWithoutIsStrInference([d_104_v41_ for d_115_i7_ in range(default__.abs(786))])), globalState)])
                d_116_v50_: _dafny.Map
                d_116_v50_ = _dafny.Map({d_62_v5_: d_57_v0_})
                d_117_v51_: T1
                nw12_ = C3()
                nw12_.ctor__(d_62_v5_, d_116_v50_)
                d_117_v51_ = nw12_
                d_118_v52_: _dafny.Map
                d_118_v52_ = _dafny.Map({d_57_v0_: len(d_114_v49_)})
                d_119_v53_: _dafny.Set
                d_119_v53_ = _dafny.Set({198})
                d_120_v54_: _dafny.MultiSet
                d_120_v54_ = _dafny.MultiSet([d_98_v35_.f14, (d_98_v35_).f15])
                d_121_v55_: D15
                d_121_v55_ = D15_DC39(d_98_v35_.f14, len(d_101_v38_), (d_61_v4_).f16, (d_120_v54_).cardinality, d_117_v51_.f10)
                d_122_v56_: int
                d_122_v56_ = d_117_v51_.f10
                d_123_v57_: _dafny.Map
                d_123_v57_ = _dafny.Map({d_122_v56_: d_98_v35_.f14})
                d_124_v58_: _dafny.Array
                nw13_ = _dafny.Array(None, 9)
                nw13_[int(0)] = d_62_v5_
                nw13_[int(1)] = len(d_99_v36_)
                nw13_[int(2)] = (d_62_v5_) - (((d_118_v52_)[d_98_v35_.f14] if (d_98_v35_.f14) in (d_118_v52_) else len(d_101_v38_)))
                nw13_[int(3)] = len(default__.fm18(len(d_119_v53_), (d_121_v55_).cf58, d_117_v51_.f10, globalState))
                nw13_[int(4)] = d_62_v5_
                nw13_[int(5)] = (default__.fm10((d_98_v35_).f15, globalState)) - (len(default__.fm25(d_62_v5_, default__.fm17(d_98_v35_.f14, d_117_v51_.f10, (d_98_v35_).f15, len(d_123_v57_), globalState), d_117_v51_.f10, d_62_v5_, globalState)))
                nw13_[int(6)] = d_117_v51_.f10
                nw13_[int(7)] = d_62_v5_
                nw13_[int(8)] = d_62_v5_
                d_124_v58_ = nw13_
                index15_ = default__.safeIndex(82, (d_124_v58_).length(0))
                (d_124_v58_)[index15_] = len((d_117_v51_).f11)
                index16_ = default__.safeIndex(82, (d_124_v58_).length(0))
                rhs22_ = ((d_114_v49_).set(default__.safeIndex(d_117_v51_.f10, len(d_114_v49_)), d_98_v35_.f14)).set(default__.safeIndex((d_120_v54_).cardinality, len((d_114_v49_).set(default__.safeIndex(d_117_v51_.f10, len(d_114_v49_)), d_98_v35_.f14))), (d_101_v38_) < (d_101_v38_))
                rhs23_ = d_117_v51_.f10
                rhs24_ = d_117_v51_
                rhs25_ = d_62_v5_
                rhs26_ = d_117_v51_.f10
                lhs15_ = globalState
                lhs16_ = d_124_v58_
                lhs17_ = default__.safeIndex(82, (d_124_v58_).length(0))
                d_114_v49_ = rhs22_
                lhs15_.f2 = rhs23_
                d_117_v51_ = rhs24_
                lhs16_[lhs17_] = rhs25_
                d_62_v5_ = rhs26_
                d_102_v39_ = d_102_v39_
                r0 = (0) - ((d_102_v39_)[default__.safeIndex((0) - (d_62_v5_), len(d_102_v39_))])
                d_125_v59_: _dafny.MultiSet
                d_125_v59_ = _dafny.MultiSet([d_124_v58_, d_124_v58_])
                index17_ = default__.safeIndex(830, (d_58_v1_).length(0))
                (d_58_v1_)[index17_] = ((d_124_v58_)[default__.safeIndex(82, (d_124_v58_).length(0))]) == ((d_125_v59_).cardinality)
                index18_ = default__.safeIndex(830, (d_58_v1_).length(0))
                rhs27_ = d_58_v1_
                rhs28_ = len((_dafny.Map({(D3_DC8((d_61_v4_).f16, d_117_v51_.f10, (0) - (d_117_v51_.f10))).cf13: (d_61_v4_).f16})).set(False, not((D3_DC6((d_124_v58_)[default__.safeIndex(82, (d_124_v58_).length(0))], d_57_v0_, (d_61_v4_).f16, d_117_v51_.f10, (d_98_v35_).f15)).cf8)))
                rhs29_ = d_98_v35_.f14
                rhs30_ = default__.safeDivisionInt(d_117_v51_.f10, d_117_v51_.f10)
                lhs18_ = d_58_v1_
                lhs19_ = default__.safeIndex(830, (d_58_v1_).length(0))
                d_58_v1_ = rhs27_
                d_62_v5_ = rhs28_
                lhs18_[lhs19_] = rhs29_
                r0 = rhs30_
            elif True:
                (d_98_v35_).f14 = (d_62_v5_) >= ((376 if (d_61_v4_).f16 else len(d_101_v38_)))
                d_126_v60_: _dafny.Set
                d_126_v60_ = _dafny.Set({d_102_v39_, d_102_v39_})
                r0 = (0) - (len(d_126_v60_))
                d_127_v61_: _dafny.Map
                d_127_v61_ = _dafny.Map({(d_62_v5_) - (d_62_v5_): d_62_v5_})
                rhs31_ = d_62_v5_
                rhs32_ = d_127_v61_
                lhs20_ = globalState
                lhs20_.f2 = rhs31_
                d_127_v61_ = rhs32_
                d_128_v62_: _dafny.Array
                nw14_ = _dafny.Array(int(0), 19)
                d_128_v62_ = nw14_
                index19_ = default__.safeIndex(333, (d_128_v62_).length(0))
                (d_128_v62_)[index19_] = len(d_102_v39_)
                d_129_v63_: _dafny.Map
                d_129_v63_ = _dafny.Map({(d_61_v4_).f16: (0) - (d_62_v5_)})
                d_130_v64_: D9
                d_130_v64_ = D9_DC19(d_129_v63_)
                d_131_v65_: D9
                d_131_v65_ = D9_DC21(d_130_v64_)
                d_132_v66_: _dafny.Array
                nw15_ = _dafny.Array(None, 2)
                nw15_[int(0)] = D9_DC21(d_130_v64_)
                nw15_[int(1)] = d_131_v65_
                d_132_v66_ = nw15_
                index20_ = default__.safeIndex(928, (d_132_v66_).length(0))
                (d_132_v66_)[index20_] = d_131_v65_
                d_133_v67_: _dafny.Map
                d_133_v67_ = _dafny.Map({d_62_v5_: (d_61_v4_).f16})
                index21_ = default__.safeIndex(333, (d_128_v62_).length(0))
                index22_ = default__.safeIndex(928, (d_132_v66_).length(0))
                rhs33_ = d_61_v4_
                rhs34_ = (d_98_v35_.f14) == (not (d_57_v0_) or (((d_133_v67_)[d_62_v5_] if (d_62_v5_) in (d_133_v67_) else d_98_v35_.f14)))
                rhs35_ = default__.fm19(d_62_v5_, globalState)
                rhs36_ = d_131_v65_
                lhs21_ = d_128_v62_
                lhs22_ = default__.safeIndex(333, (d_128_v62_).length(0))
                lhs23_ = d_132_v66_
                lhs24_ = default__.safeIndex(928, (d_132_v66_).length(0))
                d_61_v4_ = rhs33_
                d_57_v0_ = rhs34_
                lhs21_[lhs22_] = rhs35_
                lhs23_[lhs24_] = rhs36_
                d_134_v68_: _dafny.Seq
                d_134_v68_ = _dafny.SeqWithoutIsStrInference([(d_98_v35_).f15, True])
                index23_ = default__.safeIndex(333, (d_128_v62_).length(0))
                (d_128_v62_)[index23_] = len(d_134_v68_)
            d_58_v1_ = d_58_v1_
        d_135_v69_: C0
        nw16_ = C0()
        nw16_.ctor__(d_57_v0_, d_62_v5_)
        d_135_v69_ = nw16_
        d_136_v70_: _dafny.Map
        d_136_v70_ = _dafny.Map({(d_61_v4_).f16: d_135_v69_})
        d_136_v70_ = (d_136_v70_).set((d_61_v4_).f16, d_135_v69_)
        r1 = False
        r2 = d_57_v0_
        index24_ = default__.safeIndex(991, (d_58_v1_).length(0))
        (d_58_v1_)[index24_] = d_57_v0_
        d_137_v71_: _dafny.Array
        def lambda3_(d_138_v0_):
            def lambda4_(d_139_i8_):
                return (d_139_i8_) + (len(_dafny.SeqWithoutIsStrInference([d_138_v0_])))

            return lambda4_

        init1_ = lambda3_(d_57_v0_)
        nw17_ = _dafny.Array(None, 4)
        for i0_1_ in range(nw17_.length(0)):
            nw17_[i0_1_] = init1_(i0_1_)
        d_137_v71_ = nw17_
        d_140_v72_: _dafny.Seq
        d_140_v72_ = _dafny.SeqWithoutIsStrInference([(d_135_v69_).f13])
        index25_ = default__.safeIndex(991, (d_58_v1_).length(0))
        rhs37_ = d_137_v71_
        rhs38_ = (default__.fm10(d_57_v0_, globalState)) + (((d_135_v69_).f13) - ((d_140_v72_)[default__.safeIndex((d_135_v69_).f13, len(d_140_v72_))]))
        rhs39_ = d_57_v0_
        lhs25_ = globalState
        lhs26_ = d_58_v1_
        lhs27_ = default__.safeIndex(991, (d_58_v1_).length(0))
        lhs25_.f4 = rhs37_
        d_62_v5_ = rhs38_
        lhs26_[lhs27_] = rhs39_
        r0 = 809
        d_141_v73_: str
        d_141_v73_ = _dafny.CodePoint('c')
        d_142_v74_: _dafny.Seq
        d_142_v74_ = _dafny.SeqWithoutIsStrInference([d_141_v73_])
        r1 = (len(d_142_v74_)) < (d_62_v5_)
        r2 = not((((d_58_v1_)[default__.safeIndex(991, (d_58_v1_).length(0))]) == (d_57_v0_) if (d_135_v69_).f12 else d_57_v0_))
        d_143_v75_: _dafny.Map
        d_143_v75_ = _dafny.Map({(d_58_v1_)[default__.safeIndex(991, (d_58_v1_).length(0))]: (d_135_v69_).f13})
        r3 = ((d_143_v75_ if (d_61_v4_).f16 else d_143_v75_) if (d_135_v69_).f12 else d_143_v75_)
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_144_v0_: _dafny.Array
        nw18_ = _dafny.Array(_dafny.Array(None, 0), 4)
        d_144_v0_ = nw18_
        d_145_v1_: int
        d_145_v1_ = 581
        d_146_v2_: _dafny.MultiSet
        d_146_v2_ = _dafny.MultiSet([d_145_v1_, d_145_v1_])
        d_147_v3_: _dafny.Seq
        d_147_v3_ = _dafny.SeqWithoutIsStrInference([(d_146_v2_).cardinality])
        d_148_v4_: bool
        d_148_v4_ = False
        d_149_v5_: _dafny.Map
        d_149_v5_ = _dafny.Map({d_145_v1_: d_148_v4_})
        d_150_v6_: _dafny.Seq
        d_150_v6_ = _dafny.SeqWithoutIsStrInference([d_148_v4_])
        d_151_v7_: _dafny.Set
        d_151_v7_ = _dafny.Set({len(d_150_v6_)})
        d_152_v8_: int
        d_152_v8_ = d_145_v1_
        d_153_v9_: _dafny.Seq
        d_153_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "joueqk"))
        d_154_v10_: _dafny.Array
        nw19_ = _dafny.Array(None, 29)
        nw19_[int(0)] = -822
        nw19_[int(1)] = d_145_v1_
        nw19_[int(2)] = d_145_v1_
        nw19_[int(3)] = d_145_v1_
        nw19_[int(4)] = d_145_v1_
        nw19_[int(5)] = d_145_v1_
        nw19_[int(6)] = (0) - (d_145_v1_)
        nw19_[int(7)] = 288
        nw19_[int(8)] = d_145_v1_
        nw19_[int(9)] = d_145_v1_
        nw19_[int(10)] = len(d_147_v3_)
        nw19_[int(11)] = len((_dafny.SeqWithoutIsStrInference([d_149_v5_, _dafny.Map({55: d_148_v4_}), _dafny.Map({d_145_v1_: d_148_v4_})])).set(default__.safeIndex(len(d_151_v7_), len(_dafny.SeqWithoutIsStrInference([d_149_v5_, _dafny.Map({55: d_148_v4_}), _dafny.Map({d_145_v1_: d_148_v4_})]))), d_149_v5_))
        nw19_[int(12)] = d_145_v1_
        nw19_[int(13)] = d_145_v1_
        nw19_[int(14)] = len(d_149_v5_)
        nw19_[int(15)] = d_145_v1_
        nw19_[int(16)] = d_145_v1_
        nw19_[int(17)] = d_145_v1_
        nw19_[int(18)] = (d_152_v8_)
        nw19_[int(19)] = d_145_v1_
        nw19_[int(20)] = d_145_v1_
        nw19_[int(21)] = 576
        nw19_[int(22)] = d_145_v1_
        nw19_[int(23)] = len(d_153_v9_)
        nw19_[int(24)] = d_145_v1_
        nw19_[int(25)] = -713
        nw19_[int(26)] = d_145_v1_
        nw19_[int(27)] = d_145_v1_
        nw19_[int(28)] = d_145_v1_
        d_154_v10_ = nw19_
        d_155_globalState_: GlobalState
        nw20_ = GlobalState()
        nw20_.ctor__(True, False, 572, d_144_v0_, d_154_v10_, 807, 286, d_154_v10_, -297, 926)
        d_155_globalState_ = nw20_
        d_156_v11_: _dafny.Array
        nw21_ = _dafny.Array(_dafny.MultiSet({}), 2)
        d_156_v11_ = nw21_
        d_157_v12_: _dafny.Map
        d_157_v12_ = _dafny.Map({d_145_v1_: d_156_v11_})
        d_157_v12_ = (d_157_v12_).set(d_145_v1_, d_156_v11_)
        d_158_v13_: _dafny.MultiSet
        d_158_v13_ = _dafny.MultiSet([d_148_v4_, d_148_v4_, d_148_v4_, d_148_v4_])
        d_159_v14_: _dafny.Seq
        d_159_v14_ = _dafny.SeqWithoutIsStrInference([d_158_v13_])
        d_160_v15_: _dafny.Map
        d_160_v15_ = _dafny.Map({d_153_v9_: d_148_v4_})
        if default__.fm0(d_145_v1_, d_145_v1_, d_159_v14_, d_160_v15_, d_155_globalState_):
            d_161_v16_: int
            d_162_v17_: bool
            d_163_v18_: bool
            d_164_v19_: _dafny.Map
            out0_: int
            out1_: bool
            out2_: bool
            out3_: _dafny.Map
            out0_, out1_, out2_, out3_ = default__.m0(d_155_globalState_)
            d_161_v16_ = out0_
            d_162_v17_ = out1_
            d_163_v18_ = out2_
            d_164_v19_ = out3_
            d_165_v20_: _dafny.Array
            def lambda5_(d_166_v17_):
                def lambda6_(d_167_i0_):
                    return (D1_DC2(d_166_v17_)).cf2

                return lambda6_

            init2_ = lambda5_(d_162_v17_)
            nw22_ = _dafny.Array(None, 11)
            for i0_2_ in range(nw22_.length(0)):
                nw22_[i0_2_] = init2_(i0_2_)
            d_165_v20_ = nw22_
            index26_ = default__.safeIndex(428, (d_165_v20_).length(0))
            (d_165_v20_)[index26_] = (d_145_v1_) <= (len(d_150_v6_))
            index27_ = default__.safeIndex(428, (d_165_v20_).length(0))
            rhs40_ = False
            rhs41_ = d_163_v18_
            rhs42_ = 261
            rhs43_ = (d_161_v16_) >= (d_145_v1_)
            lhs28_ = d_155_globalState_
            lhs29_ = d_165_v20_
            lhs30_ = default__.safeIndex(428, (d_165_v20_).length(0))
            d_162_v17_ = rhs40_
            d_162_v17_ = rhs41_
            lhs28_.f2 = rhs42_
            lhs29_[lhs30_] = rhs43_
            d_168_v21_: D1
            d_168_v21_ = D1_DC2(d_162_v17_)
            d_168_v21_ = D1_DC2(True)
            d_163_v18_ = (d_147_v3_) == (_dafny.SeqWithoutIsStrInference([d_161_v16_]))
            d_169_v22_: int
            d_170_v23_: bool
            d_171_v24_: bool
            d_172_v25_: _dafny.Map
            out4_: int
            out5_: bool
            out6_: bool
            out7_: _dafny.Map
            out4_, out5_, out6_, out7_ = default__.m0(d_155_globalState_)
            d_169_v22_ = out4_
            d_170_v23_ = out5_
            d_171_v24_ = out6_
            d_172_v25_ = out7_
        elif True:
            d_173_v26_: int
            d_174_v27_: bool
            d_175_v28_: bool
            d_176_v29_: _dafny.Map
            out8_: int
            out9_: bool
            out10_: bool
            out11_: _dafny.Map
            out8_, out9_, out10_, out11_ = default__.m0(d_155_globalState_)
            d_173_v26_ = out8_
            d_174_v27_ = out9_
            d_175_v28_ = out10_
            d_176_v29_ = out11_
            d_148_v4_ = default__.fm0((d_145_v1_) - (d_145_v1_), d_173_v26_, (d_159_v14_) + (d_159_v14_), d_160_v15_, d_155_globalState_)
            d_154_v10_ = (d_154_v10_ if d_174_v27_ else d_154_v10_)
            if default__.fm0(d_145_v1_, d_145_v1_, d_159_v14_, (d_160_v15_) | (d_160_v15_), d_155_globalState_):
                d_177_v30_: _dafny.Array
                def lambda7_(d_178_v1_, d_179_v28_):
                    def lambda8_(d_180_i1_):
                        return _dafny.Map({d_178_v1_: not(d_179_v28_)})

                    return lambda8_

                init3_ = lambda7_(d_145_v1_, d_175_v28_)
                nw23_ = _dafny.Array(None, 23)
                for i0_3_ in range(nw23_.length(0)):
                    nw23_[i0_3_] = init3_(i0_3_)
                d_177_v30_ = nw23_
                d_177_v30_ = d_177_v30_
                index28_ = default__.safeIndex(317, (d_154_v10_).length(0))
                (d_154_v10_)[index28_] = d_173_v26_
                d_181_v31_: _dafny.Array
                nw24_ = _dafny.Array(None, 21)
                nw24_[int(0)] = d_152_v8_
                nw24_[int(1)] = d_152_v8_
                nw24_[int(2)] = d_173_v26_
                nw24_[int(3)] = d_152_v8_
                nw24_[int(4)] = 386
                nw24_[int(5)] = d_152_v8_
                nw24_[int(6)] = default__.fm1(d_174_v27_, d_148_v4_, d_155_globalState_)
                nw24_[int(7)] = 911
                nw24_[int(8)] = d_152_v8_
                nw24_[int(9)] = d_152_v8_
                nw24_[int(10)] = d_152_v8_
                nw24_[int(11)] = d_152_v8_
                nw24_[int(12)] = d_152_v8_
                nw24_[int(13)] = d_152_v8_
                nw24_[int(14)] = d_152_v8_
                nw24_[int(15)] = d_152_v8_
                nw24_[int(16)] = d_152_v8_
                nw24_[int(17)] = d_152_v8_
                nw24_[int(18)] = d_152_v8_
                nw24_[int(19)] = d_145_v1_
                nw24_[int(20)] = len(d_176_v29_)
                d_181_v31_ = nw24_
                index29_ = default__.safeIndex(94, (d_181_v31_).length(0))
                (d_181_v31_)[index29_] = d_152_v8_
                d_182_v32_: _dafny.Array
                nw25_ = _dafny.Array(False, 4)
                d_182_v32_ = nw25_
                index30_ = default__.safeIndex(58, (d_182_v32_).length(0))
                (d_182_v32_)[index30_] = d_174_v27_
                index31_ = default__.safeIndex(317, (d_154_v10_).length(0))
                index32_ = default__.safeIndex(94, (d_181_v31_).length(0))
                index33_ = default__.safeIndex(58, (d_182_v32_).length(0))
                rhs44_ = (d_173_v26_) * (d_173_v26_)
                rhs45_ = d_152_v8_
                rhs46_ = default__.fm0(len(d_149_v5_), d_145_v1_, d_159_v14_, _dafny.Map({d_153_v9_: d_148_v4_}), d_155_globalState_)
                rhs47_ = d_173_v26_
                lhs31_ = d_154_v10_
                lhs32_ = default__.safeIndex(317, (d_154_v10_).length(0))
                lhs33_ = d_181_v31_
                lhs34_ = default__.safeIndex(94, (d_181_v31_).length(0))
                lhs35_ = d_182_v32_
                lhs36_ = default__.safeIndex(58, (d_182_v32_).length(0))
                lhs31_[lhs32_] = rhs44_
                lhs33_[lhs34_] = rhs45_
                lhs35_[lhs36_] = rhs46_
                d_145_v1_ = rhs47_
                d_183_v33_: _dafny.Seq
                d_183_v33_ = _dafny.SeqWithoutIsStrInference([d_153_v9_, d_153_v9_])
                d_184_v34_: _dafny.MultiSet
                d_184_v34_ = _dafny.MultiSet([d_153_v9_])
                d_175_v28_ = (d_184_v34_).issubset(_dafny.MultiSet(d_183_v33_))
                index34_ = default__.safeIndex(58, (d_182_v32_).length(0))
                (d_182_v32_)[index34_] = d_175_v28_
                d_185_v35_: int
                d_186_v36_: bool
                d_187_v37_: bool
                d_188_v38_: _dafny.Map
                out12_: int
                out13_: bool
                out14_: bool
                out15_: _dafny.Map
                out12_, out13_, out14_, out15_ = default__.m0(d_155_globalState_)
                d_185_v35_ = out12_
                d_186_v36_ = out13_
                d_187_v37_ = out14_
                d_188_v38_ = out15_
            elif True:
                d_175_v28_ = d_148_v4_
                (d_155_globalState_).f2 = d_173_v26_
                index35_ = default__.safeIndex(553, (d_154_v10_).length(0))
                (d_154_v10_)[index35_] = d_145_v1_
                index36_ = default__.safeIndex(553, (d_154_v10_).length(0))
                (d_154_v10_)[index36_] = (d_145_v1_) * (len((d_153_v9_) + (d_153_v9_)))
                d_175_v28_ = (d_175_v28_ if False else d_175_v28_)
                d_189_v39_: str
                d_189_v39_ = _dafny.CodePoint('r')
                d_190_v40_: _dafny.Map
                d_190_v40_ = _dafny.Map({d_189_v39_: d_174_v27_})
                index37_ = default__.safeIndex(553, (d_154_v10_).length(0))
                index38_ = default__.safeIndex(553, (d_154_v10_).length(0))
                rhs48_ = d_148_v4_
                rhs49_ = 137
                rhs50_ = ((-485) * (d_173_v26_)) <= (d_173_v26_)
                rhs51_ = True
                rhs52_ = (len(d_190_v40_)) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))))
                lhs37_ = d_154_v10_
                lhs38_ = default__.safeIndex(553, (d_154_v10_).length(0))
                lhs39_ = d_154_v10_
                lhs40_ = default__.safeIndex(553, (d_154_v10_).length(0))
                d_175_v28_ = rhs48_
                lhs37_[lhs38_] = rhs49_
                d_148_v4_ = rhs50_
                d_175_v28_ = rhs51_
                lhs39_[lhs40_] = rhs52_
            d_174_v27_ = not (d_148_v4_) or ((d_173_v26_) <= (d_145_v1_))
        hi0_ = d_145_v1_
        for d_191_i2_ in range(d_145_v1_, hi0_):
            source5_ = d_152_v8_
            d_192___mcc_h0_ = source5_
            d_193_cf0_ = d_192___mcc_h0_
            d_148_v4_ = (False) or (d_148_v4_)
            d_194_v41_: _dafny.Array
            nw26_ = _dafny.Array(_dafny.Seq({}), 23)
            d_194_v41_ = nw26_
            d_195_v42_: _dafny.Seq
            d_195_v42_ = _dafny.SeqWithoutIsStrInference([d_154_v10_])
            d_196_v43_: _dafny.Map
            d_196_v43_ = _dafny.Map({d_148_v4_: d_148_v4_})
            rhs53_ = (d_195_v42_)[default__.safeIndex(len(_dafny.Map({(d_150_v6_)[default__.safeIndex(len(default__.fm2(not(d_148_v4_), d_193_cf0_, len(d_196_v43_), d_148_v4_, d_155_globalState_)), len(d_150_v6_))]: False})), len(d_195_v42_))]
            rhs54_ = d_194_v41_
            rhs55_ = False
            lhs41_ = d_155_globalState_
            lhs41_.f4 = rhs53_
            d_194_v41_ = rhs54_
            d_148_v4_ = rhs55_
            d_196_v43_ = (d_196_v43_).set(d_148_v4_, d_148_v4_)
            d_153_v9_ = d_153_v9_
            d_197_v44_: _dafny.Set
            d_197_v44_ = _dafny.Set({d_153_v9_})
            d_197_v44_ = (d_197_v44_) - (d_197_v44_)
            (d_155_globalState_).f2 = (d_145_v1_) + (default__.safeDivisionInt(d_191_i2_, 566))
            d_198_v45_: C0
            nw27_ = C0()
            nw27_.ctor__(False, d_145_v1_)
            d_198_v45_ = nw27_
            d_199_v46_: _dafny.MultiSet
            d_199_v46_ = _dafny.MultiSet([d_198_v45_, d_198_v45_])
            d_200_v47_: _dafny.Array
            nw28_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_200_v47_ = nw28_
            d_201_v48_: C1
            nw29_ = C1()
            nw29_.ctor__((d_199_v46_).issubset(d_199_v46_), D3_DC5(d_200_v47_))
            d_201_v48_ = nw29_
        d_202_v49_: _dafny.Array
        nw30_ = _dafny.Array(_dafny.Array(None, 0), 7)
        d_202_v49_ = nw30_
        d_203_v50_: D3
        d_203_v50_ = D3_DC5(d_202_v49_)
        d_204_v51_: C1
        nw31_ = C1()
        nw31_.ctor__(d_148_v4_, (d_203_v50_ if d_148_v4_ else d_203_v50_))
        d_204_v51_ = nw31_
        if d_148_v4_:
            d_153_v9_ = d_153_v9_
            d_205_v52_: int
            d_206_v53_: bool
            d_207_v54_: bool
            d_208_v55_: _dafny.Map
            out16_: int
            out17_: bool
            out18_: bool
            out19_: _dafny.Map
            out16_, out17_, out18_, out19_ = default__.m0(d_155_globalState_)
            d_205_v52_ = out16_
            d_206_v53_ = out17_
            d_207_v54_ = out18_
            d_208_v55_ = out19_
            d_209_v56_: _dafny.Array
            nw32_ = _dafny.Array(_dafny.Seq({}), 6)
            d_209_v56_ = nw32_
            d_209_v56_ = d_209_v56_
            d_204_v51_ = d_204_v51_
            d_207_v54_ = False
        elif True:
            d_210_v57_: _dafny.Array
            nw33_ = _dafny.Array(False, 8)
            d_210_v57_ = nw33_
            d_210_v57_ = d_210_v57_
            d_204_v51_ = d_204_v51_
            d_211_v58_: C2
            nw34_ = C2()
            nw34_.ctor__((d_151_v7_).issubset(_dafny.Set({414, d_145_v1_, d_145_v1_, d_145_v1_, d_145_v1_})), True)
            d_211_v58_ = nw34_
            d_211_v58_ = d_211_v58_
            d_212_v59_: str
            d_212_v59_ = _dafny.CodePoint('q')
            d_213_v60_: _dafny.Map
            d_213_v60_ = _dafny.Map({d_212_v59_: d_145_v1_})
            d_214_v61_: _dafny.Map
            d_214_v61_ = _dafny.Map({(d_213_v60_) | (d_213_v60_): default__.safeModuloInt(-347, d_145_v1_)})
            d_214_v61_ = (d_214_v61_) | (d_214_v61_)
            d_215_v62_: _dafny.Array
            nw35_ = _dafny.Array(D3.default()(), 10)
            d_215_v62_ = nw35_
            d_216_v63_: D3
            d_216_v63_ = D3_DC8((d_204_v51_).f16, d_145_v1_, d_145_v1_)
            index39_ = default__.safeIndex(502, (d_215_v62_).length(0))
            (d_215_v62_)[index39_] = d_216_v63_
            index40_ = default__.safeIndex(502, (d_215_v62_).length(0))
            (d_215_v62_)[index40_] = d_216_v63_
        hi1_ = default__.safeModuloInt(d_145_v1_, d_145_v1_)
        for d_217_i3_ in range(default__.fm17(not(False), d_145_v1_, d_148_v4_, d_145_v1_, d_155_globalState_), hi1_):
            d_218_v64_: _dafny.Map
            d_218_v64_ = _dafny.Map({(0) - (d_145_v1_): d_145_v1_})
            d_219_v65_: _dafny.Set
            d_219_v65_ = _dafny.Set({d_153_v9_, d_153_v9_})
            d_218_v64_ = (d_218_v64_).set(len((d_219_v65_) - (d_219_v65_)), (d_217_i3_) - (len(d_218_v64_)))
            d_220_v66_: _dafny.Array
            nw36_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
            d_220_v66_ = nw36_
            d_221_v67_: _dafny.Map
            d_221_v67_ = _dafny.Map({d_220_v66_: d_217_i3_})
            d_221_v67_ = (d_221_v67_).set(d_220_v66_, d_145_v1_)
            index41_ = default__.safeIndex(971, (d_154_v10_).length(0))
            (d_154_v10_)[index41_] = d_217_i3_
            index42_ = default__.safeIndex(971, (d_154_v10_).length(0))
            (d_154_v10_)[index42_] = d_217_i3_
            rhs56_ = default__.fm2(d_148_v4_, (0) - ((d_154_v10_)[default__.safeIndex(971, (d_154_v10_).length(0))]), d_145_v1_, not(d_148_v4_), d_155_globalState_)
            rhs57_ = default__.safeDivisionInt(897, d_217_i3_)
            rhs58_ = (d_154_v10_)[default__.safeIndex(971, (d_154_v10_).length(0))]
            lhs42_ = d_155_globalState_
            lhs43_ = d_155_globalState_
            d_153_v9_ = rhs56_
            lhs42_.f2 = rhs57_
            lhs43_.f2 = rhs58_
        if d_148_v4_:
            d_222_v68_: T0
            nw37_ = C2()
            nw37_.ctor__((d_204_v51_).f16, (d_204_v51_).f16)
            d_222_v68_ = nw37_
            d_223_v69_: D4
            d_223_v69_ = D4_DC10(d_222_v68_, 84, d_145_v1_)
            source6_ = d_223_v69_
            if source6_.is_DC10:
                d_224___mcc_h1_ = source6_.cf17
                d_225___mcc_h2_ = source6_.cf18
                d_226___mcc_h3_ = source6_.cf19
                d_227_cf19_ = d_226___mcc_h3_
                d_228_cf18_ = d_225___mcc_h2_
                d_229_cf17_ = d_224___mcc_h1_
                d_153_v9_ = d_153_v9_
                d_230_v70_: D8
                d_230_v70_ = D8_DC18()
                d_231_v71_: _dafny.Map
                d_231_v71_ = _dafny.Map({(0) - (d_228_cf18_): d_230_v70_})
                d_232_v72_: str
                d_232_v72_ = _dafny.CodePoint('a')
                rhs59_ = ((d_204_v51_).f16 if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bonxg")))) < ((d_153_v9_).set(default__.safeIndex(d_145_v1_, len(d_153_v9_)), d_232_v72_)) else (d_204_v51_).f16)
                rhs60_ = d_231_v71_
                rhs61_ = (d_145_v1_) >= ((d_227_cf19_) - (d_145_v1_))
                d_148_v4_ = rhs59_
                d_231_v71_ = rhs60_
                d_148_v4_ = rhs61_
                d_233_v73_: _dafny.Array
                nw38_ = _dafny.Array(None, 26)
                nw38_[int(0)] = d_204_v51_
                nw38_[int(1)] = d_204_v51_
                nw38_[int(2)] = d_204_v51_
                nw38_[int(3)] = d_204_v51_
                nw38_[int(4)] = d_204_v51_
                nw38_[int(5)] = d_204_v51_
                nw38_[int(6)] = d_204_v51_
                nw38_[int(7)] = d_204_v51_
                nw38_[int(8)] = d_204_v51_
                nw38_[int(9)] = d_204_v51_
                nw38_[int(10)] = d_204_v51_
                nw38_[int(11)] = d_204_v51_
                nw38_[int(12)] = d_204_v51_
                nw38_[int(13)] = d_204_v51_
                nw38_[int(14)] = d_204_v51_
                nw38_[int(15)] = d_204_v51_
                nw38_[int(16)] = d_204_v51_
                nw38_[int(17)] = d_204_v51_
                nw38_[int(18)] = d_204_v51_
                nw38_[int(19)] = d_204_v51_
                nw38_[int(20)] = d_204_v51_
                nw38_[int(21)] = d_204_v51_
                nw38_[int(22)] = d_204_v51_
                nw38_[int(23)] = d_204_v51_
                nw38_[int(24)] = d_204_v51_
                nw38_[int(25)] = (d_204_v51_ if d_148_v4_ else d_204_v51_)
                d_233_v73_ = nw38_
                index43_ = default__.safeIndex(327, (d_233_v73_).length(0))
                (d_233_v73_)[index43_] = d_204_v51_
                index44_ = default__.safeIndex(327, (d_233_v73_).length(0))
                (d_233_v73_)[index44_] = d_204_v51_
                d_234_v74_: _dafny.Map
                d_234_v74_ = _dafny.Map({(0) - (d_145_v1_): d_228_cf18_})
                d_234_v74_ = (d_234_v74_).set((d_145_v1_) + (d_145_v1_), d_145_v1_)
            elif True:
                d_235___mcc_h4_ = source6_.cf16
                d_236_cf16_ = d_235___mcc_h4_
                d_148_v4_ = (d_204_v51_).f16
                (d_155_globalState_).f2 = 412
                d_237_v75_: int
                d_238_v76_: int
                out20_: int
                out21_: int
                out20_, out21_ = (d_204_v51_).m5(d_145_v1_, d_155_globalState_)
                d_237_v75_ = out20_
                d_238_v76_ = out21_
                d_239_v77_: int
                d_240_v78_: int
                out22_: int
                out23_: int
                out22_, out23_ = (d_204_v51_).m5((d_147_v3_)[default__.safeIndex(d_145_v1_, len(d_147_v3_))], d_155_globalState_)
                d_239_v77_ = out22_
                d_240_v78_ = out23_
            d_241_v79_: _dafny.Array
            nw39_ = _dafny.Array(False, 6)
            d_241_v79_ = nw39_
            index45_ = default__.safeIndex(399, (d_241_v79_).length(0))
            (d_241_v79_)[index45_] = ((d_204_v51_).f16 if (d_204_v51_).f16 else (d_204_v51_).f16)
            index46_ = default__.safeIndex(399, (d_241_v79_).length(0))
            (d_241_v79_)[index46_] = not((d_204_v51_).f16)
            (d_155_globalState_).f2 = d_145_v1_
            source7_ = D11_DC26()
            if source7_.is_DC25:
                d_242___mcc_h5_ = source7_.cf37
                d_243_cf37_ = d_242___mcc_h5_
                d_244_v80_: D6
                d_244_v80_ = D6_DC13((d_241_v79_)[default__.safeIndex(399, (d_241_v79_).length(0))])
                d_245_v81_: C1
                nw40_ = C1()
                nw40_.ctor__((d_244_v80_).cf22, d_204_v51_.f17)
                d_245_v81_ = nw40_
                index47_ = default__.safeIndex(399, (d_241_v79_).length(0))
                (d_241_v79_)[index47_] = (d_241_v79_)[default__.safeIndex(399, (d_241_v79_).length(0))]
                index48_ = default__.safeIndex(399, (d_241_v79_).length(0))
                (d_241_v79_)[index48_] = (d_204_v51_).f16
                d_246_v82_: _dafny.Array
                nw41_ = _dafny.Array(_dafny.Set({}), 7)
                d_246_v82_ = nw41_
                d_247_v83_: _dafny.Map
                d_247_v83_ = _dafny.Map({d_147_v3_: d_145_v1_})
                d_248_v84_: _dafny.MultiSet
                d_248_v84_ = _dafny.MultiSet([d_243_cf37_])
                index49_ = default__.safeIndex(852, (d_246_v82_).length(0))
                def iife7_():
                    coll7_ = _dafny.Set()
                    compr_7_: _dafny.Seq
                    for compr_7_ in ((d_247_v83_).set(d_147_v3_, (d_248_v84_).cardinality)).keys.Elements:
                        d_249_v85_: _dafny.Seq = compr_7_
                        if (d_249_v85_) in ((d_247_v83_).set(d_147_v3_, (d_248_v84_).cardinality)):
                            coll7_ = coll7_.union(_dafny.Set([d_249_v85_]))
                    return _dafny.Set(coll7_)
                (d_246_v82_)[index49_] = iife7_()
                
                index50_ = default__.safeIndex(852, (d_246_v82_).length(0))
                (d_246_v82_)[index50_] = _dafny.Set({d_147_v3_, ((d_147_v3_).set(default__.safeIndex(d_145_v1_, len(d_147_v3_)), len(_dafny.SeqWithoutIsStrInference([d_243_cf37_ for d_250_i4_ in range(default__.abs(577))])))) + (d_147_v3_), default__.fm25(d_145_v1_, d_145_v1_, d_145_v1_, len(d_151_v7_), d_155_globalState_)})
            elif source7_.is_DC26:
                index51_ = default__.safeIndex(896, (d_154_v10_).length(0))
                (d_154_v10_)[index51_] = len(d_153_v9_)
                index52_ = default__.safeIndex(896, (d_154_v10_).length(0))
                (d_154_v10_)[index52_] = d_145_v1_
                index53_ = default__.safeIndex(399, (d_241_v79_).length(0))
                (d_241_v79_)[index53_] = (d_204_v51_).f16
                d_251_v86_: str
                d_251_v86_ = _dafny.CodePoint('b')
                d_148_v4_ = (not(not((d_204_v51_).fm4(d_145_v1_, (d_204_v51_).f16, d_251_v86_, d_155_globalState_))) if (d_241_v79_)[default__.safeIndex(399, (d_241_v79_).length(0))] else (d_241_v79_)[default__.safeIndex(399, (d_241_v79_).length(0))])
                d_252_v87_: _dafny.Array
                nw42_ = _dafny.Array(None, 1)
                nw42_[int(0)] = d_153_v9_
                d_252_v87_ = nw42_
                index54_ = default__.safeIndex(818, (d_252_v87_).length(0))
                (d_252_v87_)[index54_] = d_153_v9_
                d_253_v88_: C2
                nw43_ = C2()
                nw43_.ctor__((d_204_v51_).f16, (d_241_v79_)[default__.safeIndex(399, (d_241_v79_).length(0))])
                d_253_v88_ = nw43_
                d_254_v89_: _dafny.Map
                d_254_v89_ = _dafny.Map({d_253_v88_: d_253_v88_.f14})
                d_255_v90_: _dafny.Seq
                d_255_v90_ = d_153_v9_
                d_256_v91_: D12
                d_256_v91_ = D12_DC28(d_254_v89_)
                d_257_v92_: _dafny.Map
                d_257_v92_ = _dafny.Map({d_145_v1_: d_204_v51_})
                index55_ = default__.safeIndex(818, (d_252_v87_).length(0))
                rhs62_ = ((d_153_v9_) + ((d_255_v90_))) + (d_153_v9_)
                rhs63_ = (d_256_v91_).cf39
                rhs64_ = (d_154_v10_)[default__.safeIndex(896, (d_154_v10_).length(0))]
                rhs65_ = (d_204_v51_ if d_148_v4_ else ((d_257_v92_)[(0) - ((0) - (len(d_150_v6_)))] if ((0) - ((0) - (len(d_150_v6_)))) in (d_257_v92_) else d_204_v51_))
                rhs66_ = d_251_v86_
                lhs44_ = d_252_v87_
                lhs45_ = default__.safeIndex(818, (d_252_v87_).length(0))
                lhs46_ = d_155_globalState_
                lhs44_[lhs45_] = rhs62_
                d_254_v89_ = rhs63_
                lhs46_.f2 = rhs64_
                d_204_v51_ = rhs65_
                d_251_v86_ = rhs66_
            elif source7_.is_DC24:
                d_258___mcc_h6_ = source7_.cf36
                d_259_cf36_ = d_258___mcc_h6_
                d_260_v93_: _dafny.Set
                d_260_v93_ = _dafny.Set({(d_204_v51_).f16})
                d_261_v94_: _dafny.Map
                d_261_v94_ = _dafny.Map({d_145_v1_: d_145_v1_})
                index56_ = default__.safeIndex(399, (d_241_v79_).length(0))
                (d_241_v79_)[index56_] = ((default__.fm23(d_261_v94_, (d_204_v51_).f16, d_155_globalState_)) | (_dafny.Set({True, (d_204_v51_).f16}))).ispropersubset(d_260_v93_)
                d_262_v95_: _dafny.Map
                d_262_v95_ = _dafny.Map({False: d_145_v1_})
                d_262_v95_ = (d_262_v95_).set((d_148_v4_ if (d_204_v51_).f16 else (d_204_v51_).f16), default__.safeModuloInt((0) - ((0) - (d_145_v1_)), d_145_v1_))
                (d_155_globalState_).f2 = len(d_149_v5_)
                (d_155_globalState_).f2 = d_145_v1_
            elif True:
                d_263___mcc_h7_ = source7_.cf38
                d_264_cf38_ = d_263___mcc_h7_
                d_265_v96_: str
                d_265_v96_ = _dafny.CodePoint('y')
                d_265_v96_ = d_265_v96_
                d_266_v97_: T1
                nw44_ = C3()
                nw44_.ctor__((d_145_v1_ if (d_204_v51_).f16 else (d_145_v1_)), default__.fm27(d_155_globalState_))
                d_266_v97_ = nw44_
                index57_ = default__.safeIndex(399, (d_241_v79_).length(0))
                index58_ = default__.safeIndex(399, (d_241_v79_).length(0))
                rhs67_ = (default__.fm19(d_266_v97_.f10, d_155_globalState_)) == (d_266_v97_.f10)
                rhs68_ = (d_266_v97_.f10) >= (d_145_v1_)
                rhs69_ = d_266_v97_
                lhs47_ = d_241_v79_
                lhs48_ = default__.safeIndex(399, (d_241_v79_).length(0))
                lhs49_ = d_241_v79_
                lhs50_ = default__.safeIndex(399, (d_241_v79_).length(0))
                lhs47_[lhs48_] = rhs67_
                lhs49_[lhs50_] = rhs68_
                d_266_v97_ = rhs69_
                index59_ = default__.safeIndex(399, (d_241_v79_).length(0))
                (d_241_v79_)[index59_] = (d_204_v51_).f16
                index60_ = default__.safeIndex(765, (d_154_v10_).length(0))
                (d_154_v10_)[index60_] = d_145_v1_
                index61_ = default__.safeIndex(765, (d_154_v10_).length(0))
                (d_154_v10_)[index61_] = d_266_v97_.f10
            d_267_v98_: _dafny.Map
            d_267_v98_ = _dafny.Map({(d_241_v79_)[default__.safeIndex(399, (d_241_v79_).length(0))]: len((d_147_v3_).set(default__.safeIndex(d_145_v1_, len(d_147_v3_)), d_145_v1_))})
            d_268_v99_: str
            d_268_v99_ = _dafny.CodePoint('e')
            index62_ = default__.safeIndex(399, (d_241_v79_).length(0))
            rhs70_ = (d_145_v1_) <= (d_145_v1_)
            rhs71_ = (d_241_v79_)[default__.safeIndex(399, (d_241_v79_).length(0))]
            rhs72_ = d_153_v9_
            rhs73_ = (d_222_v68_).fm4((0) - (d_145_v1_), (d_153_v9_) == (d_153_v9_), d_268_v99_, d_155_globalState_)
            rhs74_ = d_267_v98_
            lhs51_ = d_241_v79_
            lhs52_ = default__.safeIndex(399, (d_241_v79_).length(0))
            d_148_v4_ = rhs70_
            d_148_v4_ = rhs71_
            d_153_v9_ = rhs72_
            lhs51_[lhs52_] = rhs73_
            d_267_v98_ = rhs74_
        elif True:
            d_269_v100_: _dafny.Array
            nw45_ = _dafny.Array(False, 29)
            d_269_v100_ = nw45_
            d_270_v101_: D12
            d_270_v101_ = D12_DC29(d_269_v100_, d_148_v4_)
            d_148_v4_ = (d_270_v101_).cf41
            d_271_v102_: T0
            nw46_ = C2()
            nw46_.ctor__((len((d_147_v3_).set(default__.safeIndex(len(d_153_v9_), len(d_147_v3_)), d_145_v1_))) >= (d_145_v1_), (d_145_v1_) <= (d_145_v1_))
            d_271_v102_ = nw46_
            d_272_v103_: _dafny.Map
            d_272_v103_ = _dafny.Map({d_145_v1_: d_149_v5_})
            rhs75_ = d_271_v102_
            rhs76_ = default__.fm28(len((d_150_v6_) + (_dafny.SeqWithoutIsStrInference([(d_204_v51_).f16, default__.fm0(891, d_145_v1_, d_159_v14_, _dafny.Map({d_153_v9_: (d_204_v51_).f16}), d_155_globalState_)]))), True, d_155_globalState_)
            d_271_v102_ = rhs75_
            d_272_v103_ = rhs76_
            d_148_v4_ = (D1_DC1(d_148_v4_)).cf1
            d_273_v104_: T1
            nw47_ = C3()
            nw47_.ctor__(len(d_147_v3_), d_149_v5_)
            d_273_v104_ = nw47_
            d_274_v105_: C2
            nw48_ = C2()
            nw48_.ctor__(d_148_v4_, (d_150_v6_)[default__.safeIndex(d_273_v104_.f10, len(d_150_v6_))])
            d_274_v105_ = nw48_
            d_275_v106_: _dafny.Map
            d_275_v106_ = _dafny.Map({d_273_v104_: d_274_v105_})
            d_276_v107_: _dafny.Seq
            d_276_v107_ = _dafny.SeqWithoutIsStrInference([d_275_v106_])
            d_277_v108_: C0
            nw49_ = C0()
            nw49_.ctor__((d_204_v51_).f16, len(d_276_v107_))
            d_277_v108_ = nw49_
            if not((d_204_v51_).f16):
                rhs77_ = ((d_204_v51_).f16 if (d_204_v51_).f16 else d_274_v105_.f14)
                rhs78_ = ((771) * (265)) - ((d_277_v108_).f13)
                rhs79_ = d_269_v100_
                rhs80_ = d_277_v108_
                lhs53_ = d_274_v105_
                lhs54_ = d_155_globalState_
                lhs53_.f14 = rhs77_
                lhs54_.f2 = rhs78_
                d_269_v100_ = rhs79_
                d_277_v108_ = rhs80_
                d_278_v109_: str
                d_278_v109_ = _dafny.CodePoint('s')
                rhs81_ = d_273_v104_.f10
                rhs82_ = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_279_i5_ in range(default__.abs(111))])) + (default__.fm2((d_204_v51_).f16, (d_277_v108_).f13, (0) - ((d_277_v108_).f13), d_274_v105_.f14, d_155_globalState_))).set(default__.safeIndex(len(d_150_v6_), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_280_i5_ in range(default__.abs(111))])) + (default__.fm2((d_204_v51_).f16, (d_277_v108_).f13, (0) - ((d_277_v108_).f13), d_274_v105_.f14, d_155_globalState_)))), d_278_v109_)
                rhs83_ = d_153_v9_
                lhs55_ = d_155_globalState_
                lhs55_.f2 = rhs81_
                d_153_v9_ = rhs82_
                d_153_v9_ = rhs83_
                index63_ = default__.safeIndex(892, (d_269_v100_).length(0))
                (d_269_v100_)[index63_] = not(True)
                index64_ = default__.safeIndex(892, (d_269_v100_).length(0))
                (d_269_v100_)[index64_] = (d_274_v105_).f15
                d_281_v110_: C3
                nw50_ = C3()
                nw50_.ctor__((d_147_v3_)[default__.safeIndex((d_277_v108_).f13, len(d_147_v3_))], ((d_273_v104_).f11).set((d_277_v108_).f13, d_148_v4_))
                d_281_v110_ = nw50_
                d_281_v110_ = d_281_v110_
                d_282_v111_: C0
                nw51_ = C0()
                nw51_.ctor__(not((d_277_v108_).f12), d_273_v104_.f10)
                d_282_v111_ = nw51_
            elif True:
                (d_274_v105_).f14 = (d_274_v105_).f15
                d_283_v112_: D10
                d_283_v112_ = D10_DC23((d_277_v108_).f13, d_274_v105_.f14, (d_204_v51_).f16)
                d_284_v113_: _dafny.Map
                d_284_v113_ = _dafny.Map({(len(d_150_v6_)) >= ((d_283_v112_).cf33): (d_274_v105_.f14) and ((d_277_v108_).f12)})
                d_284_v113_ = (d_284_v113_) | (default__.fm22(d_155_globalState_))
                (d_155_globalState_).f2 = (0) - (default__.safeModuloInt(d_145_v1_, d_145_v1_))
                (d_273_v104_).f10 = (d_277_v108_).f13
                d_285_v114_: int
                d_286_v115_: bool
                d_287_v116_: bool
                d_288_v117_: _dafny.Map
                out24_: int
                out25_: bool
                out26_: bool
                out27_: _dafny.Map
                out24_, out25_, out26_, out27_ = default__.m0(d_155_globalState_)
                d_285_v114_ = out24_
                d_286_v115_ = out25_
                d_287_v116_ = out26_
                d_288_v117_ = out27_
        index65_ = default__.safeIndex(884, (d_154_v10_).length(0))
        (d_154_v10_)[index65_] = 192
        index66_ = default__.safeIndex(884, (d_154_v10_).length(0))
        (d_154_v10_)[index66_] = d_145_v1_
        d_289_v118_: C3
        nw52_ = C3()
        nw52_.ctor__(d_145_v1_, d_149_v5_)
        d_289_v118_ = nw52_
        d_290_v119_: _dafny.Map
        d_290_v119_ = _dafny.Map({_dafny.CodePoint('m'): d_289_v118_})
        d_291_v120_: str
        d_291_v120_ = _dafny.CodePoint('j')
        d_290_v119_ = (d_290_v119_).set(d_291_v120_, d_289_v118_)
        d_292_v121_: _dafny.Array
        nw53_ = _dafny.Array(None, 5)
        nw53_[int(0)] = d_148_v4_
        nw53_[int(1)] = (d_148_v4_ if True else d_148_v4_)
        nw53_[int(2)] = (d_204_v51_).f16
        nw53_[int(3)] = (d_204_v51_).f16
        nw53_[int(4)] = (d_148_v4_) or (d_148_v4_)
        d_292_v121_ = nw53_
        index67_ = default__.safeIndex(366, (d_292_v121_).length(0))
        (d_292_v121_)[index67_] = True
        d_293_v122_: D9
        d_293_v122_ = D9_DC20(d_148_v4_, d_291_v120_)
        d_294_v123_: _dafny.MultiSet
        d_294_v123_ = _dafny.MultiSet([d_293_v122_, d_293_v122_, d_293_v122_, d_293_v122_])
        index68_ = default__.safeIndex(366, (d_292_v121_).length(0))
        (d_292_v121_)[index68_] = (d_294_v123_).isdisjoint(d_294_v123_)
        d_295_i6_: int
        d_295_i6_ = 0
        with _dafny.label("0"):
            while d_148_v4_:
                with _dafny.c_label("0"):
                    if (d_295_i6_) >= (100):
                        raise _dafny.Break("0")
                    d_295_i6_ = (d_295_i6_) + (1)
                    index69_ = default__.safeIndex(884, (d_154_v10_).length(0))
                    (d_154_v10_)[index69_] = len(d_153_v9_)
                    d_296_v125_: _dafny.MultiSet
                    d_296_v125_ = _dafny.MultiSet([d_291_v120_, d_291_v120_])
                    d_297_v126_: _dafny.Map
                    def iife8_():
                        coll8_ = _dafny.Map()
                        compr_8_: str
                        for compr_8_ in (d_296_v125_).Elements:
                            d_298_v124_: str = compr_8_
                            if (d_298_v124_) in (d_296_v125_):
                                coll8_[d_298_v124_] = d_145_v1_
                        return _dafny.Map(coll8_)
                    d_297_v126_ = _dafny.Map({(d_204_v51_).f16: default__.fm29(d_145_v1_, (iife8_()
                    ).set(d_291_v120_, len(d_147_v3_)), len(d_153_v9_), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_204_v51_).f16]))).cardinality, d_155_globalState_)})
                    d_299_v127_: D3
                    d_299_v127_ = D3_DC7(len(d_149_v5_))
                    d_300_v128_: _dafny.Map
                    d_300_v128_ = _dafny.Map({(d_289_v118_).fm7(d_297_v126_, True, d_145_v1_, d_155_globalState_): d_299_v127_})
                    d_301_v129_: _dafny.Map
                    d_301_v129_ = _dafny.Map({(d_204_v51_).f16: len(d_153_v9_)})
                    d_302_v130_: D4
                    d_302_v130_ = D4_DC9(default__.fm14(d_300_v128_, d_148_v4_, d_301_v129_, (d_292_v121_)[default__.safeIndex(366, (d_292_v121_).length(0))], d_155_globalState_))
                    d_302_v130_ = d_302_v130_
                    d_303_v131_: _dafny.Set
                    d_303_v131_ = _dafny.Set({(d_150_v6_) + (d_150_v6_), d_150_v6_})
                    d_304_v132_: _dafny.Seq
                    d_304_v132_ = _dafny.SeqWithoutIsStrInference([d_301_v129_])
                    d_305_v133_: _dafny.Map
                    d_305_v133_ = _dafny.Map({d_148_v4_: (d_150_v6_).set(default__.safeIndex(d_145_v1_, len(d_150_v6_)), (d_292_v121_)[default__.safeIndex(366, (d_292_v121_).length(0))])})
                    d_306_v134_: _dafny.Seq
                    d_306_v134_ = _dafny.SeqWithoutIsStrInference([((d_305_v133_)[False] if (False) in (d_305_v133_) else d_150_v6_), d_150_v6_])
                    d_307_v136_: _dafny.Seq
                    def iife9_():
                        coll9_ = _dafny.Set()
                        compr_9_: _dafny.Seq
                        for compr_9_ in (d_306_v134_).Elements:
                            d_308_v135_: _dafny.Seq = compr_9_
                            if (d_308_v135_) in (d_306_v134_):
                                coll9_ = coll9_.union(_dafny.Set([d_308_v135_]))
                        return _dafny.Set(coll9_)
                    d_307_v136_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({_dafny.SeqWithoutIsStrInference([(d_204_v51_).f16]), _dafny.SeqWithoutIsStrInference([d_148_v4_, (d_204_v51_).f16]), d_150_v6_, _dafny.SeqWithoutIsStrInference([(d_289_v118_).fm5(d_291_v120_, d_304_v132_, d_148_v4_, d_155_globalState_)]), d_150_v6_})) | (d_303_v131_), ((D13_DC31(d_303_v131_)).cf43).intersection(iife9_()
                    ), d_303_v131_, d_303_v131_])
                    d_309_v137_: _dafny.Map
                    d_309_v137_ = _dafny.Map({(d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))]: (d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))]})
                    d_310_v138_: _dafny.Map
                    d_310_v138_ = _dafny.Map({(d_204_v51_).f16: d_309_v137_})
                    d_303_v131_ = (d_307_v136_)[default__.safeIndex(len((d_310_v138_) | (d_310_v138_)), len(d_307_v136_))]
                    d_311_v139_: _dafny.MultiSet
                    d_311_v139_ = _dafny.MultiSet([_dafny.MultiSet([len(_dafny.Map({(d_147_v3_)[default__.safeIndex(d_145_v1_, len(d_147_v3_))]: d_291_v120_}))]), d_146_v2_, d_146_v2_])
                    d_311_v139_ = d_311_v139_
                    pass
            pass
        index70_ = default__.safeIndex(366, (d_292_v121_).length(0))
        (d_292_v121_)[index70_] = True
        d_312_v140_: _dafny.Array
        nw54_ = _dafny.Array(None, 18)
        d_312_v140_ = nw54_
        d_313_v141_: _dafny.Seq
        d_313_v141_ = _dafny.SeqWithoutIsStrInference([d_312_v140_])
        if (len(d_313_v141_)) <= (d_145_v1_):
            d_314_v142_: D14
            d_314_v142_ = D14_DC35(_dafny.SeqWithoutIsStrInference([d_204_v51_]))
            d_315_v143_: _dafny.Seq
            d_315_v143_ = _dafny.SeqWithoutIsStrInference([d_204_v51_])
            pat_let_tv0_ = d_315_v143_
            pat_let_tv1_ = d_315_v143_
            def iife10_(_pat_let0_0):
                def iife11_(d_316_dt__update__tmp_h5_):
                    def iife12_(_pat_let1_0):
                        def iife13_(d_317_dt__update_hcf50_h0_):
                            return D14_DC35(d_317_dt__update_hcf50_h0_)
                        return iife13_(_pat_let1_0)
                    return iife12_(pat_let_tv0_)
                return iife11_(_pat_let0_0)
            def iife14_(_pat_let2_0):
                def iife15_(d_318_dt__update__tmp_h6_):
                    def iife16_(_pat_let3_0):
                        def iife17_(d_319_dt__update_hcf50_h1_):
                            return D14_DC35(d_319_dt__update_hcf50_h1_)
                        return iife17_(_pat_let3_0)
                    return iife16_(pat_let_tv1_)
                return iife15_(_pat_let2_0)
            d_148_v4_ = not(((d_151_v7_) - (d_151_v7_)).issubset((d_151_v7_).intersection(_dafny.Set({len(((iife10_(d_314_v142_)).cf50).set(default__.safeIndex(d_145_v1_, len((iife14_(d_314_v142_)).cf50)), d_204_v51_))}))))
            index71_ = default__.safeIndex(884, (d_154_v10_).length(0))
            def iife18_():
                coll10_ = _dafny.Set()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(63, 992):
                    d_320_v144_: int = compr_10_
                    if ((63) <= (d_320_v144_)) and ((d_320_v144_) < (992)):
                        coll10_ = coll10_.union(_dafny.Set([default__.safeModuloInt(d_320_v144_, len(d_153_v9_))]))
                return _dafny.Set(coll10_)
            (d_154_v10_)[index71_] = len(iife18_()
            )
            d_321_v145_: D11
            d_321_v145_ = D11_DC26()
            d_321_v145_ = d_321_v145_
            d_322_v146_: int
            d_323_v147_: int
            out28_: int
            out29_: int
            out28_, out29_ = (d_204_v51_).m5((0) - ((d_145_v1_) - (len(d_153_v9_))), d_155_globalState_)
            d_322_v146_ = out28_
            d_323_v147_ = out29_
            if (d_204_v51_).f16:
                d_324_v148_: _dafny.Array
                nw55_ = _dafny.Array(None, 5)
                d_324_v148_ = nw55_
                d_324_v148_ = (d_324_v148_ if True else d_324_v148_)
                rhs84_ = d_323_v147_
                rhs85_ = not((d_204_v51_).f16)
                rhs86_ = (0) - ((d_145_v1_) - (d_145_v1_))
                d_323_v147_ = rhs84_
                d_148_v4_ = rhs85_
                d_145_v1_ = rhs86_
                d_325_v149_: _dafny.Seq
                d_325_v149_ = _dafny.SeqWithoutIsStrInference([d_289_v118_])
                d_326_v150_: _dafny.Map
                d_326_v150_ = _dafny.Map({(d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))]: (d_325_v149_)[default__.safeIndex((d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))], len(d_325_v149_))]})
                d_326_v150_ = (d_326_v150_).set(default__.fm10((d_204_v51_).f16, d_155_globalState_), d_289_v118_)
                d_148_v4_ = True
                d_327_v151_: _dafny.Seq
                d_327_v151_ = _dafny.SeqWithoutIsStrInference([(d_153_v9_).set(default__.safeIndex(d_322_v146_, len(d_153_v9_)), d_291_v120_)])
                d_145_v1_ = ((d_146_v2_)[d_145_v1_] if (d_145_v1_) in (d_146_v2_) else len((d_327_v151_)[default__.safeIndex(d_322_v146_, len(d_327_v151_))]))
            elif True:
                d_328_v152_: int
                d_329_v153_: int
                out30_: int
                out31_: int
                out30_, out31_ = (d_204_v51_).m5((d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))], d_155_globalState_)
                d_328_v152_ = out30_
                d_329_v153_ = out31_
                d_330_v154_: T1
                nw56_ = C3()
                nw56_.ctor__(d_329_v153_, d_149_v5_)
                d_330_v154_ = nw56_
                d_331_v155_: _dafny.MultiSet
                d_331_v155_ = _dafny.MultiSet([d_330_v154_])
                d_332_v156_: D15
                d_332_v156_ = D15_DC38(d_330_v154_)
                d_333_v157_: _dafny.MultiSet
                d_333_v157_ = _dafny.MultiSet([d_331_v155_, d_331_v155_, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_332_v156_).cf57, d_330_v154_, d_330_v154_])), d_331_v155_])
                d_333_v157_ = d_333_v157_
                d_334_v159_: D7
                def iife19_():
                    coll11_ = _dafny.Set()
                    compr_11_: int
                    for compr_11_ in _dafny.IntegerRange(941, 43):
                        d_335_v158_: int = compr_11_
                        if ((941) <= (d_335_v158_)) and ((d_335_v158_) < (43)):
                            coll11_ = coll11_.union(_dafny.Set([(d_335_v158_) + (d_323_v147_)]))
                    return _dafny.Set(coll11_)
                d_334_v159_ = D7_DC15(len(iife19_()
), (d_204_v51_).f16)
                d_336_v160_: T0
                nw57_ = C2()
                nw57_.ctor__(((d_334_v159_).cf25 if (d_204_v51_).f16 else (d_204_v51_).f16), not(d_148_v4_))
                d_336_v160_ = nw57_
                d_337_v161_: _dafny.Map
                d_337_v161_ = _dafny.Map({d_148_v4_: (d_204_v51_).f16})
                d_338_v162_: _dafny.Map
                d_338_v162_ = _dafny.Map({d_292_v121_: not(((d_337_v161_)[d_148_v4_] if (d_148_v4_) in (d_337_v161_) else (d_204_v51_).f16))})
                rhs87_ = (d_204_v51_).f16
                rhs88_ = d_336_v160_
                rhs89_ = (d_338_v162_) | (d_338_v162_)
                d_148_v4_ = rhs87_
                d_336_v160_ = rhs88_
                d_338_v162_ = rhs89_
                index72_ = default__.safeIndex(884, (d_154_v10_).length(0))
                (d_154_v10_)[index72_] = d_323_v147_
                d_339_v163_: _dafny.Set
                d_339_v163_ = _dafny.Set({d_291_v120_})
                index73_ = default__.safeIndex(366, (d_292_v121_).length(0))
                (d_292_v121_)[index73_] = not((_dafny.Set({d_291_v120_})).issubset((d_339_v163_) | (d_339_v163_)))
        elif True:
            d_340_v164_: D12
            d_340_v164_ = D12_DC29(d_292_v121_, d_148_v4_)
            d_341_v165_: D1
            d_341_v165_ = D1_DC1((d_204_v51_).f16)
            d_342_v166_: _dafny.Map
            d_342_v166_ = _dafny.Map({(d_204_v51_).f16: d_341_v165_})
            d_340_v164_ = (D12_DC29(d_292_v121_, (d_204_v51_).f16) if ((d_289_v118_).fm7(d_342_v166_, (d_204_v51_).f16, 575, d_155_globalState_)) and (True) else d_340_v164_)
            index74_ = default__.safeIndex(366, (d_292_v121_).length(0))
            (d_292_v121_)[index74_] = d_148_v4_
            d_343_v167_: _dafny.Set
            d_343_v167_ = _dafny.Set({d_149_v5_})
            d_148_v4_ = (d_343_v167_) != ((d_343_v167_).intersection(d_343_v167_))
            if (d_150_v6_)[default__.safeIndex((d_145_v1_) * ((_dafny.MultiSet([not(default__.fm0((d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))], d_145_v1_, d_159_v14_, _dafny.Map({d_153_v9_: default__.fm0((d_147_v3_)[default__.safeIndex(-797, len(d_147_v3_))], len(d_153_v9_), d_159_v14_, d_160_v15_, d_155_globalState_)}), d_155_globalState_)), (d_204_v51_).f16, d_148_v4_])).cardinality), len(d_150_v6_))]:
                d_344_v168_: C1
                nw58_ = C1()
                nw58_.ctor__(not(not((d_146_v2_) != (_dafny.MultiSet([d_145_v1_])))), d_203_v50_)
                d_344_v168_ = nw58_
                d_345_v169_: _dafny.Array
                nw59_ = _dafny.Array(_dafny.Set({}), 19)
                d_345_v169_ = nw59_
                d_346_v170_: _dafny.Set
                d_346_v170_ = _dafny.Set({d_153_v9_})
                index75_ = default__.safeIndex(741, (d_345_v169_).length(0))
                (d_345_v169_)[index75_] = (d_346_v170_).intersection(_dafny.Set({d_153_v9_}))
                index76_ = default__.safeIndex(741, (d_345_v169_).length(0))
                (d_345_v169_)[index76_] = d_346_v170_
                d_145_v1_ = default__.fm17(d_148_v4_, d_145_v1_, (d_204_v51_).f16, len(_dafny.Map({(d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))]: False})), d_155_globalState_)
                d_292_v121_ = d_292_v121_
                d_347_v171_: _dafny.Seq
                d_347_v171_ = _dafny.SeqWithoutIsStrInference([d_289_v118_, d_289_v118_, d_289_v118_, d_289_v118_, d_289_v118_])
                d_348_v172_: _dafny.Set
                d_348_v172_ = _dafny.Set({d_347_v171_})
                d_349_v173_: _dafny.Map
                d_349_v173_ = _dafny.Map({((d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))]) * ((0) - ((0) - (d_145_v1_))): (0) - (len(d_348_v172_))})
                def iife20_():
                    coll12_ = _dafny.Map()
                    compr_12_: int
                    for compr_12_ in (d_147_v3_).Elements:
                        d_350_v174_: int = compr_12_
                        if (d_350_v174_) in (d_147_v3_):
                            coll12_[(d_350_v174_) - ((d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))])] = 39
                    return _dafny.Map(coll12_)
                d_349_v173_ = iife20_()
                
            elif True:
                d_148_v4_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yx"))) < (d_153_v9_)
                d_148_v4_ = (d_292_v121_)[default__.safeIndex(366, (d_292_v121_).length(0))]
                d_351_v175_: D16
                d_351_v175_ = D16_DC41(d_159_v14_)
                (d_155_globalState_).f2 = len((d_351_v175_).cf64)
                d_352_v176_: _dafny.Map
                d_352_v176_ = _dafny.Map({(d_204_v51_).f16: (d_204_v51_).f16})
                d_353_v177_: _dafny.Map
                d_353_v177_ = _dafny.Map({(d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))]: (len(d_352_v176_)) - (len(d_147_v3_))})
                d_353_v177_ = (d_353_v177_).set(len((d_353_v177_) | (d_353_v177_)), 881)
                d_354_v178_: bool
                d_355_v179_: int
                d_356_v180_: bool
                d_357_v181_: bool
                out32_: bool
                out33_: int
                out34_: bool
                out35_: bool
                out32_, out33_, out34_, out35_ = (d_289_v118_).m2(d_155_globalState_)
                d_354_v178_ = out32_
                d_355_v179_ = out33_
                d_356_v180_ = out34_
                d_357_v181_ = out35_
            d_358_v182_: C2
            nw60_ = C2()
            nw60_.ctor__((d_204_v51_).f16, (d_204_v51_).f16)
            d_358_v182_ = nw60_
            d_359_v183_: _dafny.Map
            d_359_v183_ = _dafny.Map({d_358_v182_: d_148_v4_})
            source8_ = D12_DC28((d_359_v183_) | (d_359_v183_))
            if source8_.is_DC29:
                d_360___mcc_h8_ = source8_.cf40
                d_361___mcc_h9_ = source8_.cf41
                d_362_cf41_ = d_361___mcc_h9_
                d_363_cf40_ = d_360___mcc_h8_
                rhs90_ = d_145_v1_
                rhs91_ = (d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))]
                lhs56_ = d_155_globalState_
                lhs57_ = d_155_globalState_
                lhs56_.f2 = rhs90_
                lhs57_.f2 = rhs91_
                d_289_v118_ = d_289_v118_
                d_364_v184_: int
                d_365_v185_: int
                out36_: int
                out37_: int
                out36_, out37_ = (d_204_v51_).m5(-863, d_155_globalState_)
                d_364_v184_ = out36_
                d_365_v185_ = out37_
                (d_358_v182_).m4((d_358_v182_).f15, d_364_v184_, d_155_globalState_)
            elif source8_.is_DC28:
                d_366___mcc_h10_ = source8_.cf39
                d_367_cf39_ = d_366___mcc_h10_
                d_368_v186_: C1
                nw61_ = C1()
                nw61_.ctor__((-981) == (default__.fm17(not((d_358_v182_).f15), -403, (d_204_v51_).f16, (d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))], d_155_globalState_)), d_204_v51_.f17)
                d_368_v186_ = nw61_
                d_358_v182_ = d_358_v182_
                (d_155_globalState_).f2 = d_145_v1_
                d_369_v187_: _dafny.Array
                d_369_v187_ = d_154_v10_
                d_370_v188_: _dafny.Array
                nw62_ = _dafny.Array(None, 2)
                nw62_[int(0)] = d_369_v187_
                nw62_[int(1)] = d_154_v10_
                d_370_v188_ = nw62_
                index77_ = default__.safeIndex(628, (d_370_v188_).length(0))
                (d_370_v188_)[index77_] = d_369_v187_
                d_371_v190_: T0
                nw63_ = C2()
                def iife21_():
                    coll13_ = _dafny.Set()
                    compr_13_: int
                    for compr_13_ in (_dafny.Map({default__.fm10((d_368_v186_).f16, d_155_globalState_): d_146_v2_})).keys.Elements:
                        d_372_v189_: int = compr_13_
                        if (d_372_v189_) in (_dafny.Map({default__.fm10((d_368_v186_).f16, d_155_globalState_): d_146_v2_})):
                            coll13_ = coll13_.union(_dafny.Set([(d_372_v189_) - (len(_dafny.Set({False})))]))
                    return _dafny.Set(coll13_)
                nw63_.ctor__((d_151_v7_) != (iife21_()
                ), True)
                d_371_v190_ = nw63_
                index78_ = default__.safeIndex(628, (d_370_v188_).length(0))
                index79_ = default__.safeIndex(884, (d_154_v10_).length(0))
                rhs92_ = (d_369_v187_ if not((d_358_v182_).f15) else d_369_v187_)
                rhs93_ = 286
                rhs94_ = d_371_v190_
                rhs95_ = d_158_v13_
                rhs96_ = d_145_v1_
                lhs58_ = d_370_v188_
                lhs59_ = default__.safeIndex(628, (d_370_v188_).length(0))
                lhs60_ = d_154_v10_
                lhs61_ = default__.safeIndex(884, (d_154_v10_).length(0))
                lhs62_ = d_155_globalState_
                lhs58_[lhs59_] = rhs92_
                lhs60_[lhs61_] = rhs93_
                d_371_v190_ = rhs94_
                d_158_v13_ = rhs95_
                lhs62_.f2 = rhs96_
            elif True:
                d_373___mcc_h11_ = source8_.cf42
                d_374_cf42_ = d_373___mcc_h11_
                d_375_v191_: T0
                nw64_ = C1()
                nw64_.ctor__(((d_149_v5_)[128] if (128) in (d_149_v5_) else True), d_204_v51_.f17)
                d_375_v191_ = nw64_
                rhs97_ = d_375_v191_
                rhs98_ = default__.safeDivisionInt(d_145_v1_, (0) - ((d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))]))
                rhs99_ = ((d_145_v1_) - ((d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))])) * (d_145_v1_)
                rhs100_ = d_151_v7_
                lhs63_ = d_155_globalState_
                lhs64_ = d_155_globalState_
                d_375_v191_ = rhs97_
                lhs63_.f2 = rhs98_
                lhs64_.f2 = rhs99_
                d_151_v7_ = rhs100_
                d_376_v192_: C0
                nw65_ = C0()
                nw65_.ctor__(d_358_v182_.f14, d_145_v1_)
                d_376_v192_ = nw65_
                d_377_v193_: _dafny.Map
                d_377_v193_ = _dafny.Map({d_358_v182_.f14: d_376_v192_})
                d_378_v194_: D13
                d_378_v194_ = D13_DC33(d_292_v121_, d_358_v182_, d_377_v193_, (d_204_v51_).f16)
                d_379_v195_: _dafny.Map
                d_379_v195_ = _dafny.Map({(0) - ((0) - (((d_146_v2_)[(d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))]] if ((d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))]) in (d_146_v2_) else default__.fm19(d_145_v1_, d_155_globalState_)))): (d_378_v194_).cf45})
                d_379_v195_ = (d_379_v195_).set(d_145_v1_, d_292_v121_)
                d_380_v196_: bool
                d_381_v197_: int
                d_382_v198_: bool
                d_383_v199_: bool
                out38_: bool
                out39_: int
                out40_: bool
                out41_: bool
                out38_, out39_, out40_, out41_ = (d_289_v118_).m2(d_155_globalState_)
                d_380_v196_ = out38_
                d_381_v197_ = out39_
                d_382_v198_ = out40_
                d_383_v199_ = out41_
                (d_155_globalState_).f2 = (0) - ((d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))])
        d_384_i7_: int
        d_384_i7_ = 0
        with _dafny.label("1"):
            while (d_292_v121_)[default__.safeIndex(366, (d_292_v121_).length(0))]:
                with _dafny.c_label("1"):
                    if (d_384_i7_) >= (100):
                        raise _dafny.Break("1")
                    d_384_i7_ = (d_384_i7_) + (1)
                    d_385_v200_: _dafny.Array
                    nw66_ = _dafny.Array(D3.default()(), 13)
                    d_385_v200_ = nw66_
                    d_386_v201_: D3
                    d_386_v201_ = D3_DC6((d_154_v10_)[default__.safeIndex(884, (d_154_v10_).length(0))], (d_204_v51_).f16, (d_204_v51_).f16, d_145_v1_, (d_204_v51_).f16)
                    index80_ = default__.safeIndex(522, (d_385_v200_).length(0))
                    (d_385_v200_)[index80_] = d_386_v201_
                    index81_ = default__.safeIndex(522, (d_385_v200_).length(0))
                    (d_385_v200_)[index81_] = d_386_v201_
                    d_148_v4_ = (d_292_v121_)[default__.safeIndex(366, (d_292_v121_).length(0))]
                    d_387_v202_: _dafny.Set
                    d_387_v202_ = _dafny.Set({(d_204_v51_).f16})
                    d_388_v203_: _dafny.Seq
                    d_388_v203_ = _dafny.SeqWithoutIsStrInference([(d_387_v202_) | (d_387_v202_), d_387_v202_])
                    d_388_v203_ = (d_388_v203_) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({(d_204_v51_).f16}), d_387_v202_, d_387_v202_, d_387_v202_]))
                    d_148_v4_ = (d_153_v9_) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yecri"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_389_i8_ in range(default__.abs(156))])))
                    pass
            pass
        d_148_v4_ = (d_292_v121_)[default__.safeIndex(366, (d_292_v121_).length(0))]
        d_148_v4_ = not((d_150_v6_)[default__.safeIndex((0) - (d_145_v1_), len(d_150_v6_))])
        _dafny.print(_dafny.string_of(d_145_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_v2_) == (_dafny.MultiSet([581, 581]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v3_) == (_dafny.SeqWithoutIsStrInference([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_148_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v5_) == (_dafny.Map({581: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v6_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v7_) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v8_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_153_v9_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v10_)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_155_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_.f4)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_.f4)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_.f4)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_.f4)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_globalState_).f7)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_157_v12_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v13_) == (_dafny.MultiSet([False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v14_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, False, False, False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v15_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "joueqk")): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v51_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_290_v119_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_291_v120_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v121_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v121_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v121_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v121_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v121_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v122_).cf29))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v122_).cf30))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v123_) == (_dafny.MultiSet([D9_DC20(False, _dafny.CodePoint('j')), D9_DC20(False, _dafny.CodePoint('j')), D9_DC20(False, _dafny.CodePoint('j')), D9_DC20(False, _dafny.CodePoint('j'))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_295_i6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_313_v141_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_384_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: int(0)
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
        return lambda: D1_DC2(False)
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

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
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
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC4(D2, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC6(int(0), False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D3_DC5)

class D3_DC6(D3, NamedTuple('DC6', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC5(D3, NamedTuple('DC5', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC5({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC5) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(None, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC10(D4, NamedTuple('DC10', [('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)

class D5_DC11(D5, NamedTuple('DC11', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({self.cf20.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC13(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)

class D6_DC13(D6, NamedTuple('DC13', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC12(D6, NamedTuple('DC12', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC15(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)

class D7_DC15(D7, NamedTuple('DC15', [('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC14(D7, NamedTuple('DC14', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC18()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)

class D8_DC18(D8, NamedTuple('DC18', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC17(D8, NamedTuple('DC17', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC20(False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)

class D9_DC20(D9, NamedTuple('DC20', [('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC19(D9, NamedTuple('DC19', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC21(D9, NamedTuple('DC21', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC23(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)

class D10_DC23(D10, NamedTuple('DC23', [('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC22(D10, NamedTuple('DC22', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC25(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)

class D11_DC25(D11, NamedTuple('DC25', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC26(D11, NamedTuple('DC26', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC24(D11, NamedTuple('DC24', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC29(_dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)

class D12_DC29(D12, NamedTuple('DC29', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC28(D12, NamedTuple('DC28', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC30(D12, NamedTuple('DC30', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC32(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D13_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D13_DC31)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)

class D13_DC32(D13, NamedTuple('DC32', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC32({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC32) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC33(D13, NamedTuple('DC33', [('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC31(D13, NamedTuple('DC31', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC31({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC31) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC34(D13, NamedTuple('DC34', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC36(int(0), False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)

class D14_DC36(D14, NamedTuple('DC36', [('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC37(D14, NamedTuple('DC37', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC35(D14, NamedTuple('DC35', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC39(False, int(0), False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D15_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D15_DC40)

class D15_DC39(D15, NamedTuple('DC39', [('cf58', Any), ('cf59', Any), ('cf60', Any), ('cf61', Any), ('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC39({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC39) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC38(D15, NamedTuple('DC38', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC40(D15, NamedTuple('DC40', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC40({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC40) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC42(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D16_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D16_DC43)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)

class D16_DC42(D16, NamedTuple('DC42', [('cf65', Any), ('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC43(D16, NamedTuple('DC43', [('cf67', Any), ('cf68', Any), ('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC43({_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC43) and self.cf67 == __o.cf67 and self.cf68 == __o.cf68 and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC41(D16, NamedTuple('DC41', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC45(_dafny.Map({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D17_DC45)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D17_DC44)

class D17_DC45(D17, NamedTuple('DC45', [('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC45({_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC45) and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC44(D17, NamedTuple('DC44', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC44({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC44) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    @property
    def f10(self):
        return self._f10
    @f10.setter
    def f10(self, value):
        self._f10 = value
    def m1(self, globalState):
        pass

    def m2(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: int = int(0)
        self.f4: _dafny.Array = _dafny.Array(None, 0)
        self._f0: bool = False
        self._f1: bool = False
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f5: int = int(0)
        self._f6: int = int(0)
        self._f7: _dafny.Array = _dafny.Array(None, 0)
        self._f8: int = int(0)
        self._f9: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9

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

class C0:
    def  __init__(self):
        self._f12: bool = False
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f12, f13):
        (self)._f12 = f12
        (self)._f13 = f13

    def fm8(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqbqv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))

    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13

class C1(T0):
    def  __init__(self):
        self.f17: D3 = D3.default()()
        self._f16: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f16, f17):
        (self)._f16 = f16
        (self).f17 = f17

    def fm3(self, globalState):
        source9_ = 310
        d_390___mcc_h0_ = source9_
        d_391_cf0_ = d_390___mcc_h0_
        return _dafny.Map({(0) - (-795): _dafny.Map({(self).f16: d_391_cf0_})})

    def fm4(self, p0, p1, p2, globalState):
        return (self).f16

    def m5(self, p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_392_v0_: D1
        d_392_v0_ = D1_DC3((self).f16, default__.fm1((self).f16, True, globalState))
        def lambda9_(source11_):
            if source11_.is_DC2:
                d_393___mcc_h4_ = source11_.cf2
                d_394_cf2_ = d_393___mcc_h4_
                return D4_DC9(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_394_cf2_])))
            elif source11_.is_DC3:
                d_395___mcc_h5_ = source11_.cf3
                d_396___mcc_h6_ = source11_.cf4
                d_397_cf4_ = d_396___mcc_h6_
                d_398_cf3_ = d_395___mcc_h5_
                return D4_DC9(_dafny.MultiSet([d_398_cf3_, (self).f16]))
            elif True:
                d_399___mcc_h7_ = source11_.cf1
                d_400_cf1_ = d_399___mcc_h7_
                return D4_DC9(_dafny.MultiSet([True]))

        source10_ = lambda9_(d_392_v0_)
        if source10_.is_DC10:
            d_401___mcc_h0_ = source10_.cf17
            d_402___mcc_h1_ = source10_.cf18
            d_403___mcc_h2_ = source10_.cf19
            d_404_cf19_ = d_403___mcc_h2_
            d_405_cf18_ = d_402___mcc_h1_
            d_406_cf17_ = d_401___mcc_h0_
            d_407_v1_: _dafny.Map
            d_407_v1_ = _dafny.Map({(self).f16: (d_404_cf19_ if (self).f16 else d_405_cf18_)})
            d_407_v1_ = (d_407_v1_).set((self).f16, d_404_cf19_)
            d_408_v2_: D3
            d_408_v2_ = D3_DC7((p0) - (d_405_cf18_))
            d_409_v3_: _dafny.Map
            d_409_v3_ = _dafny.Map({(self).f16: (self).f16})
            d_408_v2_ = (d_408_v2_ if (len(d_409_v3_)) >= (p0) else (d_408_v2_ if (self).f16 else d_408_v2_))
            d_410_v4_: _dafny.Seq
            d_410_v4_ = _dafny.SeqWithoutIsStrInference([980])
            d_411_v5_: _dafny.Seq
            d_411_v5_ = _dafny.SeqWithoutIsStrInference([-191, (d_410_v4_)[default__.safeIndex((0) - (p0), len(d_410_v4_))]])
            d_411_v5_ = (d_410_v4_).set(default__.safeIndex(p0, len(d_410_v4_)), p0)
            d_412_v6_: _dafny.Seq
            d_412_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "raf"))
            (globalState).f2 = len(default__.fm18((len(d_412_v6_) if (self).f16 else d_405_cf18_), (self).f16, d_404_cf19_, globalState))
        elif True:
            d_413___mcc_h3_ = source10_.cf16
            d_414_cf16_ = d_413___mcc_h3_
            d_415_v7_: bool
            d_415_v7_ = True
            d_415_v7_ = (self).f16
            d_416_v8_: _dafny.Seq
            d_416_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fg"))
            d_417_v9_: str
            d_417_v9_ = _dafny.CodePoint('t')
            d_418_v10_: _dafny.MultiSet
            d_418_v10_ = _dafny.MultiSet([d_417_v9_, _dafny.CodePoint('h'), d_417_v9_])
            d_419_v11_: _dafny.MultiSet
            d_419_v11_ = _dafny.MultiSet([(0) - (p0)])
            d_420_v12_: _dafny.Set
            d_420_v12_ = _dafny.Set({(d_419_v11_).cardinality})
            d_421_v14_: _dafny.Seq
            d_421_v14_ = _dafny.SeqWithoutIsStrInference([d_416_v8_])
            d_422_v15_: _dafny.Seq
            def iife22_():
                coll14_ = _dafny.Map()
                compr_14_: _dafny.Seq
                for compr_14_ in ((d_421_v14_).set(default__.safeIndex(p0, len(d_421_v14_)), d_416_v8_)).Elements:
                    d_424_v13_: _dafny.Seq = compr_14_
                    if (d_424_v13_) in ((d_421_v14_).set(default__.safeIndex(p0, len(d_421_v14_)), d_416_v8_)):
                        coll14_[d_424_v13_] = (self).f16
                return _dafny.Map(coll14_)
            d_422_v15_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p0, p0, _dafny.SeqWithoutIsStrInference([d_414_cf16_ for d_423_i0_ in range(default__.abs(42))]), iife22_()
            , globalState), d_415_v7_])
            d_425_v16_: _dafny.Array
            nw67_ = _dafny.Array(None, 22)
            nw67_[int(0)] = (self).f16
            nw67_[int(1)] = ((self).f16 if not((self).f16) else (self).f16)
            nw67_[int(2)] = (self).f16
            nw67_[int(3)] = (self).f16
            nw67_[int(4)] = (d_416_v8_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmuiq")))
            nw67_[int(5)] = (d_418_v10_).isdisjoint(d_418_v10_)
            nw67_[int(6)] = (self).f16
            nw67_[int(7)] = not(d_415_v7_)
            nw67_[int(8)] = (d_420_v12_).issubset(d_420_v12_)
            nw67_[int(9)] = (d_415_v7_) == (d_415_v7_)
            nw67_[int(10)] = d_415_v7_
            nw67_[int(11)] = d_415_v7_
            nw67_[int(12)] = d_415_v7_
            nw67_[int(13)] = d_415_v7_
            nw67_[int(14)] = not(not(((self).f16) and (d_415_v7_)))
            nw67_[int(15)] = d_415_v7_
            nw67_[int(16)] = d_415_v7_
            nw67_[int(17)] = (self).f16
            nw67_[int(18)] = d_415_v7_
            nw67_[int(19)] = (self).f16
            nw67_[int(20)] = (d_422_v15_) == ((d_422_v15_).set(default__.safeIndex(p0, len(d_422_v15_)), (self).f16))
            nw67_[int(21)] = d_415_v7_
            d_425_v16_ = nw67_
            index82_ = default__.safeIndex(295, (d_425_v16_).length(0))
            (d_425_v16_)[index82_] = (self).f16
            index83_ = default__.safeIndex(295, (d_425_v16_).length(0))
            (d_425_v16_)[index83_] = True
            d_426_v17_: _dafny.Seq
            d_426_v17_ = _dafny.SeqWithoutIsStrInference([d_414_cf16_])
            d_427_v18_: _dafny.Map
            d_427_v18_ = _dafny.Map({d_416_v8_: d_415_v7_})
            d_428_v19_: _dafny.Map
            d_428_v19_ = _dafny.Map({d_417_v9_: d_427_v18_})
            d_429_v20_: _dafny.Map
            d_429_v20_ = _dafny.Map({(d_425_v16_)[default__.safeIndex(295, (d_425_v16_).length(0))]: default__.fm0(p0, len(default__.fm18(p0, (self).f16, p0, globalState)), d_426_v17_, ((d_428_v19_)[_dafny.CodePoint('b')] if (_dafny.CodePoint('b')) in (d_428_v19_) else d_427_v18_), globalState)})
            d_430_v21_: _dafny.Set
            d_430_v21_ = _dafny.Set({(d_425_v16_)[default__.safeIndex(295, (d_425_v16_).length(0))]})
            d_431_v22_: _dafny.Seq
            d_431_v22_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_425_v16_)[default__.safeIndex(295, (d_425_v16_).length(0))], False}), d_430_v21_, _dafny.Set({(d_425_v16_)[default__.safeIndex(295, (d_425_v16_).length(0))], not((self).f16)})])
            d_429_v20_ = _dafny.Map({default__.fm0(len((d_431_v22_)[default__.safeIndex(p0, len(d_431_v22_))]), p0, d_426_v17_, d_427_v18_, globalState): (self).f16})
            d_432_v23_: _dafny.Array
            nw68_ = _dafny.Array(_dafny.Seq({}), 29)
            d_432_v23_ = nw68_
            index84_ = default__.safeIndex(723, (d_432_v23_).length(0))
            (d_432_v23_)[index84_] = d_422_v15_
            index85_ = default__.safeIndex(723, (d_432_v23_).length(0))
            (d_432_v23_)[index85_] = d_422_v15_
        d_433_i1_: int
        d_433_i1_ = 0
        with _dafny.label("2"):
            while (-271) >= (p0):
                with _dafny.c_label("2"):
                    if (d_433_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_433_i1_ = (d_433_i1_) + (1)
                    d_434_v24_: bool
                    d_434_v24_ = False
                    d_434_v24_ = (self).f16
                    r1 = p0
                    d_435_v25_: _dafny.MultiSet
                    d_435_v25_ = _dafny.MultiSet([(0) - (default__.fm19(p0, globalState))])
                    d_436_v26_: _dafny.Map
                    d_436_v26_ = _dafny.Map({p0: p0})
                    d_437_v27_: _dafny.Array
                    nw69_ = _dafny.Array(int(0), 26)
                    d_437_v27_ = nw69_
                    d_438_v28_: _dafny.Map
                    d_438_v28_ = _dafny.Map({not(d_434_v24_): d_437_v27_})
                    d_439_v29_: _dafny.Array
                    nw70_ = _dafny.Array(None, 11)
                    nw70_[int(0)] = (self).f16
                    nw70_[int(1)] = d_434_v24_
                    nw70_[int(2)] = True
                    nw70_[int(3)] = (self).f16
                    nw70_[int(4)] = d_434_v24_
                    nw70_[int(5)] = d_434_v24_
                    nw70_[int(6)] = (self).f16
                    nw70_[int(7)] = (self).f16
                    nw70_[int(8)] = (self).f16
                    nw70_[int(9)] = True
                    nw70_[int(10)] = (self).f16
                    d_439_v29_ = nw70_
                    d_440_v30_: _dafny.Set
                    d_440_v30_ = _dafny.Set({d_439_v29_, d_439_v29_})
                    d_441_v31_: D7
                    d_441_v31_ = D7_DC14(d_440_v30_)
                    d_442_v32_: C0
                    nw71_ = C0()
                    nw71_.ctor__(d_434_v24_, p0)
                    d_442_v32_ = nw71_
                    d_443_v33_: _dafny.Map
                    d_443_v33_ = _dafny.Map({d_442_v32_: default__.fm19(p0, globalState)})
                    d_444_v34_: _dafny.Seq
                    d_444_v34_ = _dafny.SeqWithoutIsStrInference([d_434_v24_, (d_442_v32_).f12, (d_442_v32_).f12, (self).f16, d_434_v24_])
                    d_445_v35_: _dafny.Seq
                    d_445_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ie"))
                    d_446_v36_: _dafny.Map
                    d_446_v36_ = _dafny.Map({(d_442_v32_).f12: d_445_v35_})
                    d_447_v37_: _dafny.Map
                    d_447_v37_ = _dafny.Map({d_434_v24_: len(d_446_v36_)})
                    d_448_v38_: _dafny.Array
                    nw72_ = _dafny.Array(None, 24)
                    nw72_[int(0)] = p0
                    nw72_[int(1)] = ((d_436_v26_)[452] if (452) in (d_436_v26_) else len(d_438_v28_))
                    nw72_[int(2)] = p0
                    nw72_[int(3)] = p0
                    nw72_[int(4)] = p0
                    nw72_[int(5)] = (0) - (p0)
                    nw72_[int(6)] = p0
                    nw72_[int(7)] = p0
                    nw72_[int(8)] = 689
                    nw72_[int(9)] = len((d_441_v31_).cf23)
                    nw72_[int(10)] = p0
                    nw72_[int(11)] = p0
                    nw72_[int(12)] = p0
                    nw72_[int(13)] = len(d_443_v33_)
                    nw72_[int(14)] = p0
                    nw72_[int(15)] = len(default__.fm20(globalState))
                    nw72_[int(16)] = 499
                    nw72_[int(17)] = p0
                    nw72_[int(18)] = (d_442_v32_).f13
                    nw72_[int(19)] = p0
                    nw72_[int(20)] = len(d_444_v34_)
                    nw72_[int(21)] = ((d_447_v37_)[True] if (True) in (d_447_v37_) else default__.fm19(p0, globalState))
                    nw72_[int(22)] = (d_442_v32_).f13
                    nw72_[int(23)] = p0
                    d_448_v38_ = nw72_
                    d_449_v39_: _dafny.Array
                    d_449_v39_ = d_448_v38_
                    d_450_v40_: _dafny.Seq
                    d_450_v40_ = _dafny.SeqWithoutIsStrInference([len(d_445_v35_)])
                    rhs101_ = d_434_v24_
                    rhs102_ = False
                    rhs103_ = (_dafny.MultiSet(d_450_v40_)) - (((d_435_v25_).set(len(d_445_v35_), default__.abs(p0))).intersection(d_435_v25_))
                    rhs104_ = d_437_v27_
                    d_434_v24_ = rhs101_
                    d_434_v24_ = rhs102_
                    d_435_v25_ = rhs103_
                    d_449_v39_ = rhs104_
                    (globalState).f2 = p0
                    pass
            pass
        d_451_i2_: int
        d_451_i2_ = 0
        with _dafny.label("3"):
            while not(((self).f16) and (not((self).f16))):
                with _dafny.c_label("3"):
                    if (d_451_i2_) >= (100):
                        raise _dafny.Break("3")
                    d_451_i2_ = (d_451_i2_) + (1)
                    (globalState).f2 = p0
                    d_452_v41_: C0
                    nw73_ = C0()
                    nw73_.ctor__((self).f16, p0)
                    d_452_v41_ = nw73_
                    if (self).f16:
                        d_453_v42_: bool
                        d_453_v42_ = False
                        d_454_v43_: _dafny.MultiSet
                        d_454_v43_ = _dafny.MultiSet([(d_452_v41_).f13])
                        d_455_v44_: str
                        d_455_v44_ = _dafny.CodePoint('o')
                        d_453_v42_ = (self).fm4(((d_452_v41_).f13) * ((d_454_v43_).cardinality), (p0) <= ((0) - ((d_452_v41_).f13)), d_455_v44_, globalState)
                        d_453_v42_ = not((d_452_v41_).f12)
                        d_456_v45_: _dafny.Seq
                        d_456_v45_ = _dafny.SeqWithoutIsStrInference([(d_452_v41_).f13, (d_452_v41_).f13])
                        d_457_v46_: _dafny.Seq
                        d_457_v46_ = _dafny.SeqWithoutIsStrInference([(d_456_v45_)[default__.safeIndex(p0, len(d_456_v45_))], default__.fm19((d_452_v41_).f13, globalState), p0, default__.fm19((d_452_v41_).f13, globalState)])
                        d_458_v47_: _dafny.Array
                        def lambda10_(d_459_v41_):
                            def lambda11_(d_460_i3_):
                                return (d_460_i3_) + ((d_459_v41_).f13)

                            return lambda11_

                        init4_ = lambda10_(d_452_v41_)
                        nw74_ = _dafny.Array(None, 12)
                        for i0_4_ in range(nw74_.length(0)):
                            nw74_[i0_4_] = init4_(i0_4_)
                        d_458_v47_ = nw74_
                        index86_ = default__.safeIndex(556, (d_458_v47_).length(0))
                        (d_458_v47_)[index86_] = (d_452_v41_).f13
                        d_461_v48_: _dafny.Array
                        def lambda12_(d_462_v42_, d_463_v41_):
                            def lambda13_(d_464_i4_):
                                return (_dafny.MultiSet([d_462_v42_])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_463_v41_).f12])))

                            return lambda13_

                        init5_ = lambda12_(d_453_v42_, d_452_v41_)
                        nw75_ = _dafny.Array(None, 12)
                        for i0_5_ in range(nw75_.length(0)):
                            nw75_[i0_5_] = init5_(i0_5_)
                        d_461_v48_ = nw75_
                        d_465_v49_: _dafny.MultiSet
                        d_465_v49_ = _dafny.MultiSet([(d_452_v41_).f12])
                        index87_ = default__.safeIndex(898, (d_461_v48_).length(0))
                        (d_461_v48_)[index87_] = (d_465_v49_) | (default__.fm21((d_452_v41_).f13, False, (d_452_v41_).f13, globalState))
                        d_466_v50_: _dafny.Set
                        d_466_v50_ = _dafny.Set({53})
                        index88_ = default__.safeIndex(556, (d_458_v47_).length(0))
                        index89_ = default__.safeIndex(898, (d_461_v48_).length(0))
                        rhs105_ = not((d_466_v50_).isdisjoint(d_466_v50_))
                        rhs106_ = ((d_456_v45_) + (d_456_v45_)) + (d_457_v46_)
                        rhs107_ = default__.fm19(((d_454_v43_).cardinality) - ((d_465_v49_).cardinality), globalState)
                        rhs108_ = d_465_v49_
                        lhs65_ = d_458_v47_
                        lhs66_ = default__.safeIndex(556, (d_458_v47_).length(0))
                        lhs67_ = d_461_v48_
                        lhs68_ = default__.safeIndex(898, (d_461_v48_).length(0))
                        d_453_v42_ = rhs105_
                        d_456_v45_ = rhs106_
                        lhs65_[lhs66_] = rhs107_
                        lhs67_[lhs68_] = rhs108_
                        d_467_v51_: C0
                        nw76_ = C0()
                        nw76_.ctor__((d_452_v41_).f12, (0) - ((d_458_v47_)[default__.safeIndex(556, (d_458_v47_).length(0))]))
                        d_467_v51_ = nw76_
                        d_468_v52_: _dafny.Map
                        d_468_v52_ = _dafny.Map({d_453_v42_: (self).f16})
                        d_469_v53_: _dafny.Array
                        nw77_ = _dafny.Array(None, 29)
                        nw77_[int(0)] = default__.fm22(globalState)
                        nw77_[int(1)] = _dafny.Map({(self).f16: (d_452_v41_).f12})
                        nw77_[int(2)] = (default__.fm22(globalState)) | (d_468_v52_)
                        nw77_[int(3)] = (default__.fm22(globalState)) | (d_468_v52_)
                        nw77_[int(4)] = _dafny.Map({(self).f16: (d_452_v41_).f12})
                        nw77_[int(5)] = _dafny.Map({(d_467_v51_).f12: (self).f16})
                        nw77_[int(6)] = d_468_v52_
                        nw77_[int(7)] = d_468_v52_
                        nw77_[int(8)] = d_468_v52_
                        nw77_[int(9)] = d_468_v52_
                        nw77_[int(10)] = (d_468_v52_) | (d_468_v52_)
                        nw77_[int(11)] = d_468_v52_
                        nw77_[int(12)] = d_468_v52_
                        nw77_[int(13)] = d_468_v52_
                        nw77_[int(14)] = d_468_v52_
                        nw77_[int(15)] = (d_468_v52_) | (d_468_v52_)
                        nw77_[int(16)] = d_468_v52_
                        nw77_[int(17)] = d_468_v52_
                        nw77_[int(18)] = d_468_v52_
                        nw77_[int(19)] = d_468_v52_
                        nw77_[int(20)] = (d_468_v52_ if (d_467_v51_).f12 else d_468_v52_)
                        nw77_[int(21)] = d_468_v52_
                        nw77_[int(22)] = (d_468_v52_) | (d_468_v52_)
                        nw77_[int(23)] = (_dafny.Map({(d_452_v41_).f12: (d_452_v41_).f12})) | (d_468_v52_)
                        nw77_[int(24)] = (d_468_v52_) | (d_468_v52_)
                        nw77_[int(25)] = (d_468_v52_).set(d_453_v42_, (self).f16)
                        nw77_[int(26)] = d_468_v52_
                        nw77_[int(27)] = d_468_v52_
                        nw77_[int(28)] = _dafny.Map({(d_452_v41_).f12: (d_467_v51_).f12})
                        d_469_v53_ = nw77_
                        index90_ = default__.safeIndex(932, (d_469_v53_).length(0))
                        (d_469_v53_)[index90_] = (d_468_v52_) | (d_468_v52_)
                        index91_ = default__.safeIndex(932, (d_469_v53_).length(0))
                        (d_469_v53_)[index91_] = d_468_v52_
                    elif True:
                        r1 = p0
                        d_470_v54_: _dafny.Array
                        nw78_ = _dafny.Array(False, 25)
                        d_470_v54_ = nw78_
                        d_470_v54_ = d_470_v54_
                        d_470_v54_ = d_470_v54_
                        d_471_v55_: bool
                        d_471_v55_ = True
                        d_472_v56_: str
                        d_472_v56_ = _dafny.CodePoint('m')
                        d_473_v57_: _dafny.Seq
                        d_473_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
                        d_471_v55_ = (d_472_v56_) in (d_473_v57_)
                        d_474_v58_: C0
                        nw79_ = C0()
                        nw79_.ctor__(((d_452_v41_).f13) != (734), (p0) + (p0))
                        d_474_v58_ = nw79_
                    if (d_452_v41_).f12:
                        d_475_v59_: bool
                        d_475_v59_ = True
                        d_476_v60_: str
                        d_476_v60_ = _dafny.CodePoint('v')
                        d_477_v61_: _dafny.Seq
                        d_477_v61_ = _dafny.SeqWithoutIsStrInference([d_476_v60_, d_476_v60_, d_476_v60_, d_476_v60_, d_476_v60_])
                        d_475_v59_ = (d_477_v61_) <= (d_477_v61_)
                        d_476_v60_ = d_476_v60_
                        d_478_v62_: _dafny.Array
                        def lambda14_(d_479_v61_, d_480_v41_):
                            def lambda15_(d_481_i5_):
                                return D7_DC15(len(d_479_v61_), (d_480_v41_).f12)

                            return lambda15_

                        init6_ = lambda14_(d_477_v61_, d_452_v41_)
                        nw80_ = _dafny.Array(None, 7)
                        for i0_6_ in range(nw80_.length(0)):
                            nw80_[i0_6_] = init6_(i0_6_)
                        d_478_v62_ = nw80_
                        d_482_v63_: D7
                        d_482_v63_ = D7_DC15((d_452_v41_).f13, (self).f16)
                        index92_ = default__.safeIndex(867, (d_478_v62_).length(0))
                        (d_478_v62_)[index92_] = d_482_v63_
                        index93_ = default__.safeIndex(867, (d_478_v62_).length(0))
                        (d_478_v62_)[index93_] = d_482_v63_
                        d_475_v59_ = (d_452_v41_).f12
                        rhs109_ = (p0) * ((d_452_v41_).f13)
                        rhs110_ = (default__.fm19((d_452_v41_).f13, globalState)) >= (p0)
                        rhs111_ = (d_452_v41_).f13
                        rhs112_ = (default__.safeModuloInt((d_452_v41_).f13, (d_452_v41_).f13)) + ((d_452_v41_).f13)
                        lhs69_ = globalState
                        lhs70_ = globalState
                        lhs69_.f2 = rhs109_
                        d_475_v59_ = rhs110_
                        lhs70_.f2 = rhs111_
                        r1 = rhs112_
                    elif True:
                        d_483_v64_: bool
                        d_483_v64_ = True
                        d_483_v64_ = d_483_v64_
                        d_484_v65_: _dafny.Map
                        d_484_v65_ = _dafny.Map({(d_452_v41_).f12: (d_452_v41_).f12})
                        d_485_v66_: _dafny.Set
                        d_485_v66_ = _dafny.Set({p0, (d_452_v41_).f13})
                        d_486_v67_: D3
                        d_486_v67_ = D3_DC8((self).f16, len(d_485_v66_), p0)
                        d_487_v68_: _dafny.MultiSet
                        d_487_v68_ = _dafny.MultiSet([(d_486_v67_).cf15])
                        d_488_v69_: _dafny.Map
                        d_488_v69_ = _dafny.Map({(d_452_v41_).f13: (self).f16})
                        d_489_v70_: D3
                        d_489_v70_ = D3_DC8((d_452_v41_).f12, (d_452_v41_).f13, ((d_487_v68_)[p0] if (p0) in (d_487_v68_) else len(d_488_v69_)))
                        d_490_v71_: _dafny.MultiSet
                        d_490_v71_ = _dafny.MultiSet([(len(d_484_v65_)) * ((d_489_v70_).cf14), (d_452_v41_).f13])
                        rhs113_ = (d_490_v71_) | (_dafny.MultiSet([(d_452_v41_).f13]))
                        rhs114_ = p0
                        rhs115_ = default__.fm19((d_452_v41_).f13, globalState)
                        rhs116_ = (self).f16
                        lhs71_ = globalState
                        d_490_v71_ = rhs113_
                        lhs71_.f2 = rhs114_
                        r0 = rhs115_
                        d_483_v64_ = rhs116_
                        d_491_v72_: int
                        d_491_v72_ = (d_452_v41_).f13
                        d_492_v73_: _dafny.MultiSet
                        d_492_v73_ = _dafny.MultiSet([D1_DC3((d_452_v41_).f12, d_491_v72_), d_392_v0_])
                        r0 = (((d_492_v73_)[d_392_v0_] if (d_392_v0_) in (d_492_v73_) else default__.fm19(p0, globalState))) - ((0) - ((d_452_v41_).f13))
                        d_493_v74_: _dafny.Array
                        nw81_ = _dafny.Array(False, 10)
                        d_493_v74_ = nw81_
                        index94_ = default__.safeIndex(832, (d_493_v74_).length(0))
                        (d_493_v74_)[index94_] = (d_452_v41_).f12
                        d_494_v75_: _dafny.Seq
                        d_494_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eotb"))
                        d_495_v76_: _dafny.Map
                        d_495_v76_ = _dafny.Map({len(d_494_v75_): d_486_v67_})
                        d_496_v78_: str
                        d_496_v78_ = _dafny.CodePoint('d')
                        d_497_v79_: _dafny.Map
                        d_497_v79_ = _dafny.Map({(d_452_v41_).f13: d_496_v78_})
                        index95_ = default__.safeIndex(832, (d_493_v74_).length(0))
                        def iife23_():
                            coll15_ = _dafny.Map()
                            compr_15_: int
                            for compr_15_ in (d_497_v79_).keys.Elements:
                                d_498_v77_: int = compr_15_
                                if (d_498_v77_) in (d_497_v79_):
                                    coll15_[(d_498_v77_) + (748)] = d_486_v67_
                            return _dafny.Map(coll15_)
                        rhs117_ = (p0) < (p0)
                        rhs118_ = ((d_495_v76_ if d_483_v64_ else d_495_v76_)) | ((d_495_v76_) | (iife23_()
                        ))
                        lhs72_ = d_493_v74_
                        lhs73_ = default__.safeIndex(832, (d_493_v74_).length(0))
                        lhs72_[lhs73_] = rhs117_
                        d_495_v76_ = rhs118_
                        d_499_v80_: C0
                        nw82_ = C0()
                        nw82_.ctor__((d_452_v41_).f12, (d_452_v41_).f13)
                        d_499_v80_ = nw82_
                    pass
            pass
        d_500_v81_: bool
        d_500_v81_ = True
        rhs119_ = (self).f16
        rhs120_ = p0
        lhs74_ = globalState
        d_500_v81_ = rhs119_
        lhs74_.f2 = rhs120_
        hi2_ = p0
        for d_501_i6_ in range(p0, hi2_):
            d_502_v82_: int
            d_503_v83_: bool
            d_504_v84_: bool
            d_505_v85_: _dafny.Map
            out42_: int
            out43_: bool
            out44_: bool
            out45_: _dafny.Map
            out42_, out43_, out44_, out45_ = default__.m0(globalState)
            d_502_v82_ = out42_
            d_503_v83_ = out43_
            d_504_v84_ = out44_
            d_505_v85_ = out45_
            d_506_v86_: _dafny.Array
            nw83_ = _dafny.Array(int(0), 23)
            d_506_v86_ = nw83_
            index96_ = default__.safeIndex(246, (d_506_v86_).length(0))
            (d_506_v86_)[index96_] = default__.fm19((0) - (p0), globalState)
            d_507_v87_: _dafny.Array
            nw84_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
            d_507_v87_ = nw84_
            d_508_v88_: _dafny.Seq
            d_508_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jasbi"))
            d_509_v89_: _dafny.Seq
            d_509_v89_ = d_508_v88_
            index97_ = default__.safeIndex(487, (d_507_v87_).length(0))
            (d_507_v87_)[index97_] = d_509_v89_
            d_510_v90_: _dafny.Map
            d_510_v90_ = _dafny.Map({d_501_i6_: (self).f16})
            d_511_v91_: _dafny.Seq
            d_511_v91_ = _dafny.SeqWithoutIsStrInference([d_503_v83_])
            d_512_v92_: _dafny.Map
            d_512_v92_ = _dafny.Map({d_501_i6_: p0})
            d_513_v93_: _dafny.MultiSet
            d_513_v93_ = _dafny.MultiSet([((d_512_v92_)[d_501_i6_] if (d_501_i6_) in (d_512_v92_) else 267)])
            d_514_v94_: _dafny.MultiSet
            d_514_v94_ = _dafny.MultiSet([d_500_v81_])
            d_515_v95_: _dafny.Seq
            d_515_v95_ = _dafny.SeqWithoutIsStrInference([d_514_v94_])
            d_516_v96_: _dafny.Map
            d_516_v96_ = _dafny.Map({d_508_v88_: d_500_v81_})
            d_517_v97_: D3
            d_517_v97_ = D3_DC6(d_502_v82_, ((d_510_v90_)[len(d_511_v91_)] if (len(d_511_v91_)) in (d_510_v90_) else False), not(default__.fm0((d_513_v93_).cardinality, len(d_511_v91_), d_515_v95_, d_516_v96_, globalState)), len(_dafny.Map({d_502_v82_: (self).f16})), True)
            d_518_v98_: _dafny.Seq
            d_518_v98_ = _dafny.SeqWithoutIsStrInference([d_502_v82_, d_501_i6_])
            index98_ = default__.safeIndex(246, (d_506_v86_).length(0))
            index99_ = default__.safeIndex(487, (d_507_v87_).length(0))
            rhs121_ = (d_517_v97_).cf7
            rhs122_ = d_501_i6_
            rhs123_ = len((d_518_v98_) + ((d_518_v98_ if d_504_v84_ else d_518_v98_)))
            rhs124_ = p0
            rhs125_ = d_509_v89_
            lhs75_ = d_506_v86_
            lhs76_ = default__.safeIndex(246, (d_506_v86_).length(0))
            lhs77_ = globalState
            lhs78_ = globalState
            lhs79_ = d_507_v87_
            lhs80_ = default__.safeIndex(487, (d_507_v87_).length(0))
            d_502_v82_ = rhs121_
            lhs75_[lhs76_] = rhs122_
            lhs77_.f2 = rhs123_
            lhs78_.f2 = rhs124_
            lhs79_[lhs80_] = rhs125_
            d_504_v84_ = (p0) >= (default__.safeDivisionInt(d_502_v82_, 965))
            d_519_v99_: str
            d_519_v99_ = _dafny.CodePoint('e')
            d_520_v100_: _dafny.Map
            d_520_v100_ = _dafny.Map({not(d_500_v81_): False})
            d_519_v99_ = (d_519_v99_ if ((d_520_v100_)[d_504_v84_] if (d_504_v84_) in (d_520_v100_) else True) else _dafny.CodePoint('m'))
        d_521_v101_: _dafny.Seq
        d_521_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tri"))
        r0 = len(d_521_v101_)
        r0 = p0
        r1 = p0
        return r0, r1

    @property
    def f16(self):
        return self._f16

class C2(T0):
    def  __init__(self):
        self.f14: bool = False
        self._f15: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f14, f15):
        (self).f14 = f14
        (self)._f15 = f15

    def fm3(self, globalState):
        return _dafny.Map({(531) + (len(_dafny.SeqWithoutIsStrInference([(self).f15, self.f14, False]))): (_dafny.Map({self.f14: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_522_i0_ in range(default__.abs(886))]))})) | (_dafny.Map({self.f14: 730}))})

    def fm4(self, p0, p1, p2, globalState):
        return self.f14

    def fm12(self, p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([(_dafny.MultiSet([self.f14])).cardinality, 73])).issubset(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.Set({self.f14}))])})) for d_523_i0_ in range(default__.abs(766))]))]))

    def fm13(self, p0, p1, p2, globalState):
        return (406) > ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgislaabn")))) - (169))

    def m3(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_524_v0_: int
        d_524_v0_ = 55
        d_525_v1_: C0
        nw85_ = C0()
        nw85_.ctor__((self).f15, d_524_v0_)
        d_525_v1_ = nw85_
        d_526_v2_: _dafny.Seq
        d_526_v2_ = _dafny.SeqWithoutIsStrInference([d_525_v1_, d_525_v1_])
        d_527_v3_: _dafny.MultiSet
        d_527_v3_ = _dafny.MultiSet([not((self).f15)])
        d_528_v4_: D3
        d_528_v4_ = D3_DC7(d_524_v0_)
        d_529_v5_: _dafny.Map
        d_529_v5_ = _dafny.Map({True: d_528_v4_})
        d_530_v6_: _dafny.Map
        d_530_v6_ = _dafny.Map({self.f14: d_524_v0_})
        d_531_v7_: _dafny.Seq
        d_531_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vedhrd"))
        d_532_v8_: _dafny.Map
        d_532_v8_ = _dafny.Map({(d_525_v1_).f13: len(d_530_v6_)})
        d_533_v10_: _dafny.Set
        d_533_v10_ = _dafny.Set({(d_525_v1_).f13, (d_525_v1_).f13})
        d_534_v11_: str
        d_534_v11_ = _dafny.CodePoint('i')
        d_535_v12_: _dafny.Seq
        d_535_v12_ = _dafny.SeqWithoutIsStrInference([len(d_533_v10_), (d_525_v1_).f13, (_dafny.MultiSet([d_534_v11_, d_534_v11_])).cardinality, d_524_v0_])
        d_536_v13_: _dafny.Array
        nw86_ = _dafny.Array(None, 26)
        nw86_[int(0)] = (d_524_v0_) - (d_524_v0_)
        nw86_[int(1)] = d_524_v0_
        nw86_[int(2)] = d_524_v0_
        nw86_[int(3)] = -457
        nw86_[int(4)] = d_524_v0_
        nw86_[int(5)] = (0) - (d_524_v0_)
        nw86_[int(6)] = d_524_v0_
        nw86_[int(7)] = (d_524_v0_) - (452)
        nw86_[int(8)] = d_524_v0_
        nw86_[int(9)] = len((_dafny.Map({(d_526_v2_)[default__.safeIndex((d_525_v1_).f13, len(d_526_v2_))]: (d_525_v1_).f13})) | (_dafny.Map({d_525_v1_: d_524_v0_})))
        nw86_[int(10)] = ((d_527_v3_).set((d_525_v1_).f12, default__.abs(d_524_v0_))).cardinality
        nw86_[int(11)] = (d_525_v1_).f13
        nw86_[int(12)] = (default__.fm14(d_529_v5_, self.f14, d_530_v6_, (self).f15, globalState)).cardinality
        nw86_[int(13)] = d_524_v0_
        nw86_[int(14)] = ((_dafny.MultiSet([(self).f15, (self).f15])).cardinality) + (len(d_531_v7_))
        nw86_[int(15)] = (0) - (d_524_v0_)
        nw86_[int(16)] = (d_525_v1_).f13
        def iife24_():
            coll16_ = _dafny.Map()
            compr_16_: _dafny.Seq
            for compr_16_ in (_dafny.Map({p0: (d_525_v1_).f13})).keys.Elements:
                d_537_v9_: _dafny.Seq = compr_16_
                if (d_537_v9_) in (_dafny.Map({p0: (d_525_v1_).f13})):
                    coll16_[d_537_v9_] = False
            return _dafny.Map(coll16_)
        nw86_[int(17)] = len(((d_532_v8_).set((d_525_v1_).f13, (d_525_v1_).f13)).set(len(iife24_()
        ), -304))
        nw86_[int(18)] = (d_525_v1_).f13
        nw86_[int(19)] = (d_525_v1_).f13
        nw86_[int(20)] = ((d_525_v1_).f13) - (len(d_535_v12_))
        nw86_[int(21)] = -752
        nw86_[int(22)] = default__.safeDivisionInt(117, (d_525_v1_).f13)
        nw86_[int(23)] = ((d_527_v3_ if (self).f15 else d_527_v3_)).cardinality
        nw86_[int(24)] = d_524_v0_
        nw86_[int(25)] = (293 if self.f14 else d_524_v0_)
        d_536_v13_ = nw86_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_536_v13_).length(0)):
            d_538_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_538_i0_)) and ((d_538_i0_) < ((d_536_v13_).length(0)))):
                (d_536_v13_)[(d_538_i0_)] = (d_538_i0_) - ((0) - (d_524_v0_))
        d_539_v14_: _dafny.Array
        nw87_ = _dafny.Array(None, 13)
        d_539_v14_ = nw87_
        d_539_v14_ = d_539_v14_
        d_540_v15_: C0
        nw88_ = C0()
        nw88_.ctor__(True, default__.safeDivisionInt(d_524_v0_, d_524_v0_))
        d_540_v15_ = nw88_
        d_541_v16_: _dafny.Map
        d_541_v16_ = _dafny.Map({(d_540_v15_).f13: False})
        d_542_v17_: _dafny.Seq
        d_542_v17_ = _dafny.SeqWithoutIsStrInference([d_541_v16_, (d_541_v16_) | (d_541_v16_), d_541_v16_])
        d_542_v17_ = d_542_v17_
        d_543_v18_: _dafny.Map
        d_543_v18_ = _dafny.Map({(self).f15: D3_DC8(self.f14, len((d_535_v12_).set(default__.safeIndex((d_540_v15_).f13, len(d_535_v12_)), (d_525_v1_).f13)), d_524_v0_)})
        d_544_v19_: D3
        d_544_v19_ = D3_DC8((self).fm4((d_540_v15_).f13, (d_540_v15_).f12, d_534_v11_, globalState), d_524_v0_, (d_540_v15_).f13)
        d_545_v20_: _dafny.Map
        d_545_v20_ = _dafny.Map({self.f14: (d_540_v15_).f12})
        d_546_v21_: _dafny.Set
        d_546_v21_ = _dafny.Set({True, (d_540_v15_).f12})
        d_547_v22_: _dafny.Array
        nw89_ = _dafny.Array(None, 25)
        nw89_[int(0)] = (d_543_v18_) | (d_543_v18_)
        nw89_[int(1)] = (d_543_v18_) | (d_543_v18_)
        nw89_[int(2)] = d_543_v18_
        nw89_[int(3)] = d_543_v18_
        nw89_[int(4)] = d_543_v18_
        nw89_[int(5)] = (_dafny.Map({self.f14: d_544_v19_})) | (d_543_v18_)
        nw89_[int(6)] = ((d_543_v18_).set(not((d_525_v1_).f12), d_544_v19_)) | (d_543_v18_)
        nw89_[int(7)] = (default__.fm15(globalState)) | (_dafny.Map({(d_525_v1_).f12: D3_DC8(self.f14, (d_540_v15_).f13, -421)}))
        nw89_[int(8)] = _dafny.Map({(d_540_v15_).f12: default__.fm16(globalState)})
        nw89_[int(9)] = d_543_v18_
        nw89_[int(10)] = d_543_v18_
        nw89_[int(11)] = d_543_v18_
        nw89_[int(12)] = (d_543_v18_) | ((d_543_v18_).set(True, D3_DC8((d_540_v15_).f12, d_524_v0_, d_524_v0_)))
        nw89_[int(13)] = _dafny.Map({(d_540_v15_).f12: D3_DC8(False, default__.fm17(((d_545_v20_)[self.f14] if (self.f14) in (d_545_v20_) else self.f14), d_524_v0_, (d_540_v15_).f12, (d_525_v1_).f13, globalState), len(d_546_v21_))})
        nw89_[int(14)] = (d_543_v18_) | (d_543_v18_)
        nw89_[int(15)] = d_543_v18_
        nw89_[int(16)] = (d_543_v18_) | (d_543_v18_)
        nw89_[int(17)] = (d_543_v18_) | (d_543_v18_)
        nw89_[int(18)] = ((D6_DC12(d_543_v18_)).cf21) | (_dafny.Map({(d_525_v1_).f12: d_544_v19_}))
        nw89_[int(19)] = d_543_v18_
        nw89_[int(20)] = (d_543_v18_).set((self).f15, d_544_v19_)
        nw89_[int(21)] = d_543_v18_
        nw89_[int(22)] = (_dafny.Map({(d_525_v1_).f12: d_544_v19_})).set((self).f15, d_544_v19_)
        nw89_[int(23)] = d_543_v18_
        nw89_[int(24)] = d_543_v18_
        d_547_v22_ = nw89_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_547_v22_).length(0)):
            d_548_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_548_i1_)) and ((d_548_i1_) < ((d_547_v22_).length(0)))):
                (d_547_v22_)[(d_548_i1_)] = d_543_v18_
        d_549_v23_: _dafny.Seq
        d_549_v23_ = d_531_v7_
        d_549_v23_ = d_531_v7_
        r0 = (d_525_v1_).f12
        nw90_ = _dafny.Array(False, 27)
        r1 = nw90_
        return r0, r1

    def m4(self, p0, p1, globalState):
        (globalState).f2 = (p1) + (p1)
        d_550_v0_: str
        d_550_v0_ = _dafny.CodePoint('v')
        d_551_v1_: _dafny.Map
        d_551_v1_ = _dafny.Map({False: p1})
        d_552_v2_: _dafny.Map
        d_552_v2_ = _dafny.Map({((d_551_v1_)[p0] if (p0) in (d_551_v1_) else p1): d_550_v0_})
        d_553_v3_: _dafny.MultiSet
        d_553_v3_ = _dafny.MultiSet([(self).f15])
        d_550_v0_ = ((d_552_v2_)[(692) * (((d_553_v3_)[self.f14] if (self.f14) in (d_553_v3_) else 648))] if ((692) * (((d_553_v3_)[self.f14] if (self.f14) in (d_553_v3_) else 648))) in (d_552_v2_) else d_550_v0_)
        d_554_v4_: int
        d_554_v4_ = p1
        d_555_i0_: int
        d_555_i0_ = 0
        with _dafny.label("4"):
            def lambda16_(source12_):
                d_565___mcc_h0_ = source12_
                d_566_cf0_ = d_565___mcc_h0_
                return self.f14

            while lambda16_(d_554_v4_):
                with _dafny.c_label("4"):
                    if (d_555_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_555_i0_ = (d_555_i0_) + (1)
                    d_556_v5_: _dafny.MultiSet
                    d_556_v5_ = _dafny.MultiSet([685, (d_553_v3_).cardinality, 525, p1])
                    d_557_v6_: _dafny.Array
                    nw91_ = _dafny.Array(_dafny.Array(None, 0), 3)
                    d_557_v6_ = nw91_
                    d_558_v7_: D3
                    d_558_v7_ = D3_DC5(d_557_v6_)
                    d_559_v8_: T0
                    nw92_ = C1()
                    nw92_.ctor__(self.f14, d_558_v7_)
                    d_559_v8_ = nw92_
                    d_560_v9_: _dafny.Map
                    d_560_v9_ = _dafny.Map({(self).f15: d_559_v8_})
                    (globalState).f2 = ((((d_556_v5_)[len(d_560_v9_)] if (len(d_560_v9_)) in (d_556_v5_) else p1)) + (p1)) * ((0) - (p1))
                    (self).f14 = p0
                    d_561_v10_: _dafny.Seq
                    d_561_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                    d_562_v11_: _dafny.Map
                    d_562_v11_ = _dafny.Map({(p0) or (p0): ((0) - (len(d_561_v10_))) != (p1)})
                    d_562_v11_ = default__.fm22(globalState)
                    d_563_v12_: _dafny.Array
                    nw93_ = _dafny.Array(None, 3)
                    nw93_[int(0)] = (self).f15
                    nw93_[int(1)] = not(p0)
                    nw93_[int(2)] = p0
                    d_563_v12_ = nw93_
                    d_564_v13_: _dafny.MultiSet
                    d_564_v13_ = _dafny.MultiSet([d_563_v12_])
                    d_564_v13_ = (d_564_v13_) - ((d_564_v13_).set(d_563_v12_, default__.abs(len(d_551_v1_))))
                    pass
            pass
        d_567_i1_: int
        d_567_i1_ = 0
        with _dafny.label("5"):
            while p0:
                with _dafny.c_label("5"):
                    if (d_567_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_567_i1_ = (d_567_i1_) + (1)
                    (self).f14 = (self).f15
                    d_568_v14_: _dafny.Set
                    d_568_v14_ = _dafny.Set({d_550_v0_, _dafny.CodePoint('y')})
                    d_569_v15_: _dafny.Map
                    d_569_v15_ = _dafny.Map({(self).f15: p0})
                    d_570_v16_: C0
                    nw94_ = C0()
                    nw94_.ctor__(p0, len(d_569_v15_))
                    d_570_v16_ = nw94_
                    d_571_v17_: _dafny.Seq
                    d_571_v17_ = _dafny.SeqWithoutIsStrInference([d_570_v16_])
                    d_572_v18_: D3
                    d_572_v18_ = D3_DC8((self).f15, (d_570_v16_).f13, (d_570_v16_).f13)
                    d_573_v19_: _dafny.Array
                    nw95_ = _dafny.Array(None, 20)
                    nw95_[int(0)] = (self).f15
                    nw95_[int(1)] = self.f14
                    nw95_[int(2)] = True
                    nw95_[int(3)] = p0
                    nw95_[int(4)] = p0
                    nw95_[int(5)] = (p1) > (p1)
                    nw95_[int(6)] = (self).f15
                    nw95_[int(7)] = (self).f15
                    nw95_[int(8)] = (self).fm4(p1, p0, d_550_v0_, globalState)
                    nw95_[int(9)] = False
                    nw95_[int(10)] = (d_568_v14_).isdisjoint(d_568_v14_)
                    nw95_[int(11)] = (_dafny.SeqWithoutIsStrInference([d_570_v16_, d_570_v16_])) < (d_571_v17_)
                    nw95_[int(12)] = p0
                    nw95_[int(13)] = (True) and ((d_570_v16_).f12)
                    nw95_[int(14)] = True
                    nw95_[int(15)] = (self).f15
                    nw95_[int(16)] = (d_570_v16_).f12
                    nw95_[int(17)] = (d_570_v16_).f12
                    nw95_[int(18)] = ((self).f15 if (d_572_v18_).cf13 else (d_570_v16_).f12)
                    nw95_[int(19)] = self.f14
                    d_573_v19_ = nw95_
                    index100_ = default__.safeIndex(66, (d_573_v19_).length(0))
                    (d_573_v19_)[index100_] = (self).f15
                    d_574_v20_: D7
                    d_574_v20_ = D7_DC15((d_570_v16_).f13, self.f14)
                    index101_ = default__.safeIndex(66, (d_573_v19_).length(0))
                    (d_573_v19_)[index101_] = (self.f14 if not(self.f14) else ((d_574_v20_).cf24) <= (p1))
                    d_575_v21_: _dafny.Map
                    d_575_v21_ = _dafny.Map({(d_570_v16_).f13: len(_dafny.SeqWithoutIsStrInference([d_550_v0_ for d_576_i2_ in range(default__.abs(-457))]))})
                    d_577_v22_: _dafny.Set
                    d_577_v22_ = _dafny.Set({(d_570_v16_).f13, (0) - (((d_575_v21_)[(d_570_v16_).f13] if ((d_570_v16_).f13) in (d_575_v21_) else -421))})
                    d_578_v23_: _dafny.Array
                    nw96_ = _dafny.Array(None, 19)
                    nw96_[int(0)] = d_573_v19_
                    nw96_[int(1)] = d_573_v19_
                    nw96_[int(2)] = d_573_v19_
                    nw96_[int(3)] = d_573_v19_
                    nw96_[int(4)] = d_573_v19_
                    nw96_[int(5)] = d_573_v19_
                    nw96_[int(6)] = d_573_v19_
                    nw96_[int(7)] = d_573_v19_
                    nw96_[int(8)] = d_573_v19_
                    nw96_[int(9)] = d_573_v19_
                    nw96_[int(10)] = d_573_v19_
                    nw96_[int(11)] = d_573_v19_
                    nw96_[int(12)] = d_573_v19_
                    nw96_[int(13)] = d_573_v19_
                    nw96_[int(14)] = d_573_v19_
                    nw96_[int(15)] = d_573_v19_
                    nw96_[int(16)] = d_573_v19_
                    nw96_[int(17)] = d_573_v19_
                    nw96_[int(18)] = d_573_v19_
                    d_578_v23_ = nw96_
                    d_579_v24_: D3
                    d_579_v24_ = D3_DC5(d_578_v23_)
                    d_580_v25_: T0
                    nw97_ = C1()
                    nw97_.ctor__(False, d_579_v24_)
                    d_580_v25_ = nw97_
                    d_581_v26_: _dafny.Map
                    d_581_v26_ = _dafny.Map({d_580_v25_: d_577_v22_})
                    d_577_v22_ = ((d_581_v26_)[d_580_v25_] if (d_580_v25_) in (d_581_v26_) else (d_577_v22_).intersection(d_577_v22_))
                    d_582_v28_: _dafny.Seq
                    d_582_v28_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_583_v29_: _dafny.Seq
                    d_583_v29_ = _dafny.SeqWithoutIsStrInference([d_553_v3_])
                    d_584_v30_: _dafny.Seq
                    d_584_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "leigq"))
                    d_585_v31_: _dafny.Map
                    d_585_v31_ = _dafny.Map({d_584_v30_: (d_573_v19_)[default__.safeIndex(66, (d_573_v19_).length(0))]})
                    d_586_v32_: _dafny.Set
                    def iife25_():
                        coll17_ = _dafny.Map()
                        compr_17_: int
                        for compr_17_ in (d_582_v28_).Elements:
                            d_587_v27_: int = compr_17_
                            if (d_587_v27_) in (d_582_v28_):
                                coll17_[(d_587_v27_) * ((d_570_v16_).f13)] = (self).f15
                        return _dafny.Map(coll17_)
                    d_586_v32_ = _dafny.Set({self.f14, default__.fm0(p1, len(iife25_()
                    ), d_583_v29_, d_585_v31_, globalState)})
                    d_588_v33_: _dafny.Seq
                    d_588_v33_ = _dafny.SeqWithoutIsStrInference([d_586_v32_])
                    d_589_v34_: _dafny.MultiSet
                    d_589_v34_ = _dafny.MultiSet([((d_553_v3_).cardinality) * (len(d_586_v32_)), p1, p1, default__.safeModuloInt((0) - (len((d_588_v33_)[default__.safeIndex((d_570_v16_).f13, len(d_588_v33_))])), p1), (p1) - (p1)])
                    (globalState).f2 = (0) - (((d_589_v34_)[p1] if (p1) in (d_589_v34_) else p1))
                    pass
            pass
        d_590_v35_: _dafny.Seq
        d_590_v35_ = _dafny.SeqWithoutIsStrInference([self.f14, self.f14])
        d_591_v36_: bool
        d_592_v37_: _dafny.Array
        out46_: bool
        out47_: _dafny.Array
        out46_, out47_ = (self).m3(d_590_v35_, globalState)
        d_591_v36_ = out46_
        d_592_v37_ = out47_
        d_593_v38_: D4
        d_593_v38_ = D4_DC9((d_553_v3_).set(d_591_v36_, default__.abs(p1)))
        d_593_v38_ = (d_593_v38_ if p0 else d_593_v38_)

    @property
    def f15(self):
        return self._f15

class C3(T1, T0):
    def  __init__(self):
        self._f10: int = int(0)
        self._f11: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f10(self):
        return self._f10
    @f10.setter
    def f10(self, value):
        self._f10 = value
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f10, f11):
        (self).f10 = f10
        (self)._f11 = f11

    def fm5(self, p0, p1, p2, globalState):
        def iife26_():
            coll18_ = _dafny.Set()
            compr_18_: int
            for compr_18_ in (_dafny.Set({self.f10})).Elements:
                d_594_v0_: int = compr_18_
                if (d_594_v0_) in (_dafny.Set({self.f10})):
                    coll18_ = coll18_.union(_dafny.Set([default__.safeModuloInt(d_594_v0_, 777)]))
            return _dafny.Set(coll18_)
        return not (True) or ((_dafny.Set({246})).ispropersubset(iife26_()
        ))

    def fm6(self, p0, globalState):
        source13_ = D1_DC2(not(True))
        if source13_.is_DC2:
            d_595___mcc_h0_ = source13_.cf2
            d_596_cf2_ = d_595___mcc_h0_
            return D1_DC2(d_596_cf2_)
        elif source13_.is_DC3:
            d_597___mcc_h1_ = source13_.cf3
            d_598___mcc_h2_ = source13_.cf4
            d_599_cf4_ = d_598___mcc_h2_
            d_600_cf3_ = d_597___mcc_h1_
            return D1_DC2(d_600_cf3_)
        elif True:
            d_601___mcc_h3_ = source13_.cf1
            d_602_cf1_ = d_601___mcc_h3_
            return D1_DC2(False)

    def fm3(self, globalState):
        return (_dafny.Map({self.f10: _dafny.Map({True: self.f10})})) | (_dafny.Map({len((self).f11): _dafny.Map({True: self.f10})}))

    def fm4(self, p0, p1, p2, globalState):
        return (self.f10) >= ((0) - (self.f10))

    def fm7(self, p0, p1, p2, globalState):
        return not(((not(True) if not(True) else False) if True else False))

    def m1(self, globalState):
        r0: int = int(0)
        d_603_v0_: bool
        d_603_v0_ = True
        d_603_v0_ = (d_603_v0_ if d_603_v0_ else d_603_v0_)
        d_604_v1_: _dafny.Seq
        d_604_v1_ = _dafny.SeqWithoutIsStrInference([self.f10])
        d_605_v2_: _dafny.Set
        d_605_v2_ = _dafny.Set({self.f10, self.f10, self.f10})
        if (((d_604_v1_)[default__.safeIndex(len(d_605_v2_), len(d_604_v1_))]) <= (self.f10) if d_603_v0_ else d_603_v0_):
            d_606_v3_: C0
            nw98_ = C0()
            nw98_.ctor__((self.f10) >= ((0) - (self.f10)), (0) - (self.f10))
            d_606_v3_ = nw98_
            if (d_603_v0_) or ((d_606_v3_).f12):
                (self).f10 = (d_606_v3_).f13
                d_607_v4_: int
                d_608_v5_: bool
                d_609_v6_: bool
                d_610_v7_: _dafny.Map
                out48_: int
                out49_: bool
                out50_: bool
                out51_: _dafny.Map
                out48_, out49_, out50_, out51_ = default__.m0(globalState)
                d_607_v4_ = out48_
                d_608_v5_ = out49_
                d_609_v6_ = out50_
                d_610_v7_ = out51_
                d_611_v8_: _dafny.Array
                def lambda17_(d_612_v3_):
                    def lambda18_(d_613_i0_):
                        return (len(_dafny.SeqWithoutIsStrInference([(d_612_v3_).f12]))) == ((d_612_v3_).f13)

                    return lambda18_

                init7_ = lambda17_(d_606_v3_)
                nw99_ = _dafny.Array(None, 21)
                for i0_7_ in range(nw99_.length(0)):
                    nw99_[i0_7_] = init7_(i0_7_)
                d_611_v8_ = nw99_
                d_611_v8_ = d_611_v8_
                d_614_v9_: _dafny.Array
                def lambda19_(d_615_v1_):
                    def lambda20_(d_616_i1_):
                        return default__.safeModuloInt(d_616_i1_, len(d_615_v1_))

                    return lambda20_

                init8_ = lambda19_(d_604_v1_)
                nw100_ = _dafny.Array(None, 29)
                for i0_8_ in range(nw100_.length(0)):
                    nw100_[i0_8_] = init8_(i0_8_)
                d_614_v9_ = nw100_
                d_617_v10_: _dafny.Array
                nw101_ = _dafny.Array(None, 23)
                d_617_v10_ = nw101_
                d_618_v11_: _dafny.Map
                d_618_v11_ = _dafny.Map({d_617_v10_: d_614_v9_})
                d_619_v12_: _dafny.Array
                d_619_v12_ = d_614_v9_
                d_620_v13_: _dafny.Seq
                d_620_v13_ = _dafny.SeqWithoutIsStrInference([d_614_v9_, d_614_v9_, d_614_v9_, d_614_v9_, d_614_v9_])
                d_621_v14_: _dafny.Array
                nw102_ = _dafny.Array(None, 29)
                nw102_[int(0)] = (d_614_v9_ if d_609_v6_ else d_614_v9_)
                nw102_[int(1)] = d_614_v9_
                nw102_[int(2)] = d_614_v9_
                nw102_[int(3)] = d_614_v9_
                nw102_[int(4)] = d_614_v9_
                nw102_[int(5)] = d_614_v9_
                nw102_[int(6)] = d_614_v9_
                nw102_[int(7)] = d_614_v9_
                nw102_[int(8)] = d_614_v9_
                nw102_[int(9)] = ((d_618_v11_)[d_617_v10_] if (d_617_v10_) in (d_618_v11_) else d_614_v9_)
                nw102_[int(10)] = d_614_v9_
                nw102_[int(11)] = (d_614_v9_ if (d_606_v3_).f12 else d_614_v9_)
                nw102_[int(12)] = d_614_v9_
                nw102_[int(13)] = d_614_v9_
                nw102_[int(14)] = d_614_v9_
                nw102_[int(15)] = d_614_v9_
                nw102_[int(16)] = d_614_v9_
                nw102_[int(17)] = d_614_v9_
                nw102_[int(18)] = (d_619_v12_)
                nw102_[int(19)] = d_614_v9_
                nw102_[int(20)] = d_614_v9_
                nw102_[int(21)] = d_614_v9_
                nw102_[int(22)] = d_614_v9_
                nw102_[int(23)] = d_614_v9_
                nw102_[int(24)] = d_614_v9_
                nw102_[int(25)] = d_614_v9_
                nw102_[int(26)] = d_614_v9_
                nw102_[int(27)] = d_614_v9_
                nw102_[int(28)] = (d_620_v13_)[default__.safeIndex(d_607_v4_, len(d_620_v13_))]
                d_621_v14_ = nw102_
                index102_ = default__.safeIndex(339, (d_621_v14_).length(0))
                (d_621_v14_)[index102_] = d_614_v9_
                index103_ = default__.safeIndex(339, (d_621_v14_).length(0))
                (d_621_v14_)[index103_] = d_614_v9_
                index104_ = default__.safeIndex(781, (d_611_v8_).length(0))
                (d_611_v8_)[index104_] = d_608_v5_
                d_622_v15_: _dafny.Array
                nw103_ = _dafny.Array(_dafny.Set({}), 13)
                d_622_v15_ = nw103_
                d_623_v16_: _dafny.Seq
                d_623_v16_ = _dafny.SeqWithoutIsStrInference([(d_606_v3_).f12])
                d_624_v17_: _dafny.Map
                d_624_v17_ = _dafny.Map({d_623_v16_: self.f10})
                index105_ = default__.safeIndex(255, (d_622_v15_).length(0))
                (d_622_v15_)[index105_] = default__.fm9(d_603_v0_, d_624_v17_, d_607_v4_, _dafny.CodePoint('u'), globalState)
                d_625_v18_: _dafny.Array
                nw104_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_625_v18_ = nw104_
                d_626_v19_: _dafny.Seq
                d_626_v19_ = _dafny.SeqWithoutIsStrInference([d_625_v18_])
                d_627_v20_: _dafny.Array
                nw105_ = _dafny.Array(None, 12)
                nw105_[int(0)] = (D3_DC5(d_625_v18_)).cf6
                nw105_[int(1)] = d_625_v18_
                nw105_[int(2)] = d_625_v18_
                nw105_[int(3)] = (d_626_v19_)[default__.safeIndex((d_606_v3_).f13, len(d_626_v19_))]
                nw105_[int(4)] = d_625_v18_
                nw105_[int(5)] = d_625_v18_
                nw105_[int(6)] = d_625_v18_
                nw105_[int(7)] = d_625_v18_
                nw105_[int(8)] = d_625_v18_
                nw105_[int(9)] = d_625_v18_
                nw105_[int(10)] = (d_625_v18_ if True else d_625_v18_)
                nw105_[int(11)] = d_625_v18_
                d_627_v20_ = nw105_
                index106_ = default__.safeIndex(868, (d_627_v20_).length(0))
                (d_627_v20_)[index106_] = d_625_v18_
                d_628_v21_: str
                d_628_v21_ = _dafny.CodePoint('m')
                d_629_v22_: _dafny.Seq
                d_629_v22_ = _dafny.SeqWithoutIsStrInference([d_628_v21_, d_628_v21_])
                d_630_v23_: _dafny.Map
                d_630_v23_ = _dafny.Map({default__.fm10(d_608_v5_, globalState): -133})
                d_631_v24_: _dafny.MultiSet
                d_631_v24_ = _dafny.MultiSet([d_603_v0_])
                d_632_v25_: D4
                d_632_v25_ = D4_DC9(d_631_v24_)
                d_633_v26_: _dafny.Array
                nw106_ = _dafny.Array(None, 20)
                nw106_[int(0)] = d_628_v21_
                nw106_[int(1)] = d_628_v21_
                nw106_[int(2)] = d_628_v21_
                nw106_[int(3)] = (d_629_v22_)[default__.safeIndex((d_606_v3_).f13, len(d_629_v22_))]
                nw106_[int(4)] = d_628_v21_
                nw106_[int(5)] = (d_629_v22_)[default__.safeIndex(((d_630_v23_)[((d_610_v7_)[(d_606_v3_).f12] if ((d_606_v3_).f12) in (d_610_v7_) else ((d_632_v25_).cf16).cardinality)] if (((d_610_v7_)[(d_606_v3_).f12] if ((d_606_v3_).f12) in (d_610_v7_) else ((d_632_v25_).cf16).cardinality)) in (d_630_v23_) else (d_606_v3_).f13), len(d_629_v22_))]
                nw106_[int(6)] = d_628_v21_
                nw106_[int(7)] = d_628_v21_
                nw106_[int(8)] = d_628_v21_
                nw106_[int(9)] = d_628_v21_
                nw106_[int(10)] = (d_628_v21_ if d_608_v5_ else d_628_v21_)
                nw106_[int(11)] = _dafny.CodePoint('p')
                nw106_[int(12)] = _dafny.CodePoint('y')
                nw106_[int(13)] = d_628_v21_
                nw106_[int(14)] = d_628_v21_
                nw106_[int(15)] = d_628_v21_
                nw106_[int(16)] = d_628_v21_
                nw106_[int(17)] = d_628_v21_
                nw106_[int(18)] = _dafny.CodePoint('c')
                nw106_[int(19)] = d_628_v21_
                d_633_v26_ = nw106_
                index107_ = default__.safeIndex(563, (d_633_v26_).length(0))
                (d_633_v26_)[index107_] = d_628_v21_
                d_634_v27_: _dafny.Seq
                d_634_v27_ = _dafny.SeqWithoutIsStrInference([d_631_v24_, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_603_v0_]))])
                d_635_v28_: _dafny.Map
                d_635_v28_ = _dafny.Map({d_629_v22_: d_609_v6_})
                d_636_v29_: _dafny.Set
                d_636_v29_ = _dafny.Set({d_603_v0_, default__.fm0(len(d_604_v1_), self.f10, d_634_v27_, d_635_v28_, globalState), (d_606_v3_).f12, d_603_v0_})
                d_637_v31_: _dafny.Map
                d_637_v31_ = _dafny.Map({d_629_v22_: (d_606_v3_).f13})
                d_638_v32_: _dafny.Seq
                def iife27_():
                    coll19_ = _dafny.Map()
                    compr_19_: _dafny.Seq
                    for compr_19_ in (d_637_v31_).keys.Elements:
                        d_639_v30_: _dafny.Seq = compr_19_
                        if (d_639_v30_) in (d_637_v31_):
                            coll19_[d_639_v30_] = d_608_v5_
                    return _dafny.Map(coll19_)
                d_638_v32_ = _dafny.SeqWithoutIsStrInference([iife27_()
                ])
                index108_ = default__.safeIndex(781, (d_611_v8_).length(0))
                index109_ = default__.safeIndex(255, (d_622_v15_).length(0))
                index110_ = default__.safeIndex(868, (d_627_v20_).length(0))
                index111_ = default__.safeIndex(563, (d_633_v26_).length(0))
                rhs126_ = (default__.fm0((0) - (len(d_636_v29_)), self.f10, default__.fm11(globalState), (d_638_v32_)[default__.safeIndex((d_606_v3_).f13, len(d_638_v32_))], globalState) if d_608_v5_ else d_603_v0_)
                rhs127_ = ((d_606_v3_).f13) * (self.f10)
                rhs128_ = d_605_v2_
                rhs129_ = d_625_v18_
                rhs130_ = (d_629_v22_)[default__.safeIndex(d_607_v4_, len(d_629_v22_))]
                lhs81_ = d_611_v8_
                lhs82_ = default__.safeIndex(781, (d_611_v8_).length(0))
                lhs83_ = d_622_v15_
                lhs84_ = default__.safeIndex(255, (d_622_v15_).length(0))
                lhs85_ = d_627_v20_
                lhs86_ = default__.safeIndex(868, (d_627_v20_).length(0))
                lhs87_ = d_633_v26_
                lhs88_ = default__.safeIndex(563, (d_633_v26_).length(0))
                lhs81_[lhs82_] = rhs126_
                d_607_v4_ = rhs127_
                lhs83_[lhs84_] = rhs128_
                lhs85_[lhs86_] = rhs129_
                lhs87_[lhs88_] = rhs130_
            elif True:
                d_604_v1_ = d_604_v1_
                r0 = ((d_606_v3_).f13) - (-909)
                d_640_v33_: _dafny.Array
                nw107_ = _dafny.Array(False, 3)
                d_640_v33_ = nw107_
                rhs131_ = default__.safeModuloInt(default__.safeDivisionInt((d_606_v3_).f13, (d_606_v3_).f13), default__.safeModuloInt((d_606_v3_).f13, self.f10))
                rhs132_ = d_640_v33_
                r0 = rhs131_
                d_640_v33_ = rhs132_
                d_641_v34_: _dafny.Seq
                d_641_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajenxyshb"))
                r0 = default__.safeModuloInt((d_604_v1_)[default__.safeIndex((d_606_v3_).f13, len(d_604_v1_))], len(d_641_v34_))
                d_642_v35_: _dafny.Seq
                d_642_v35_ = d_641_v34_
                d_643_v36_: _dafny.Seq
                d_643_v36_ = _dafny.SeqWithoutIsStrInference([default__.fm2(False, 996, (d_606_v3_).f13, not(d_603_v0_), globalState), ((d_641_v34_)) + (d_641_v34_)])
                d_643_v36_ = d_643_v36_
            d_644_v37_: _dafny.Seq
            d_644_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adcgwephx"))
            d_644_v37_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lndvj"))) + (d_644_v37_)
            if (D3_DC8(not((d_606_v3_).f12), len(d_605_v2_), (d_606_v3_).f13)).cf13:
                d_603_v0_ = (d_606_v3_).f12
                d_645_v38_: _dafny.Array
                def lambda21_(d_646_v0_):
                    def lambda22_(d_647_i2_):
                        return d_646_v0_

                    return lambda22_

                init9_ = lambda21_(d_603_v0_)
                nw108_ = _dafny.Array(None, 8)
                for i0_9_ in range(nw108_.length(0)):
                    nw108_[i0_9_] = init9_(i0_9_)
                d_645_v38_ = nw108_
                index112_ = default__.safeIndex(887, (d_645_v38_).length(0))
                (d_645_v38_)[index112_] = d_603_v0_
                index113_ = default__.safeIndex(887, (d_645_v38_).length(0))
                (d_645_v38_)[index113_] = d_603_v0_
                d_648_v39_: D3
                d_648_v39_ = D3_DC8((d_645_v38_)[default__.safeIndex(887, (d_645_v38_).length(0))], self.f10, 704)
                index114_ = default__.safeIndex(887, (d_645_v38_).length(0))
                (d_645_v38_)[index114_] = (len(d_604_v1_)) > (default__.safeModuloInt((d_648_v39_).cf15, len(d_644_v37_)))
                d_603_v0_ = not(d_603_v0_)
                d_649_v40_: _dafny.Seq
                d_649_v40_ = d_644_v37_
                d_650_v41_: _dafny.Map
                d_650_v41_ = _dafny.Map({(d_649_v40_): (0) - ((d_606_v3_).f13)})
                d_650_v41_ = (d_650_v41_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wayojdhbc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwne"))), (d_606_v3_).f13)
            elif True:
                d_651_v42_: _dafny.Array
                def lambda23_(d_652_v0_):
                    def lambda24_(d_653_i3_):
                        return d_652_v0_

                    return lambda24_

                init10_ = lambda23_(d_603_v0_)
                nw109_ = _dafny.Array(None, 8)
                for i0_10_ in range(nw109_.length(0)):
                    nw109_[i0_10_] = init10_(i0_10_)
                d_651_v42_ = nw109_
                index115_ = default__.safeIndex(117, (d_651_v42_).length(0))
                (d_651_v42_)[index115_] = not(not(not((d_606_v3_).f12)))
                index116_ = default__.safeIndex(117, (d_651_v42_).length(0))
                (d_651_v42_)[index116_] = False
                d_606_v3_ = d_606_v3_
                d_603_v0_ = (d_651_v42_)[default__.safeIndex(117, (d_651_v42_).length(0))]
                index117_ = default__.safeIndex(117, (d_651_v42_).length(0))
                index118_ = default__.safeIndex(117, (d_651_v42_).length(0))
                rhs133_ = (d_606_v3_).f12
                rhs134_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_654_i4_ in range(default__.abs(377))])) + (d_644_v37_)
                rhs135_ = (d_606_v3_).f12
                rhs136_ = (((d_606_v3_).f12) and ((d_651_v42_)[default__.safeIndex(117, (d_651_v42_).length(0))])) or ((d_606_v3_).f12)
                lhs89_ = d_651_v42_
                lhs90_ = default__.safeIndex(117, (d_651_v42_).length(0))
                lhs91_ = d_651_v42_
                lhs92_ = default__.safeIndex(117, (d_651_v42_).length(0))
                d_603_v0_ = rhs133_
                d_644_v37_ = rhs134_
                lhs89_[lhs90_] = rhs135_
                lhs91_[lhs92_] = rhs136_
                d_655_v43_: _dafny.Array
                def lambda25_(d_656_i5_):
                    return (d_656_i5_) - (self.f10)

                init11_ = lambda25_
                nw110_ = _dafny.Array(None, 6)
                for i0_11_ in range(nw110_.length(0)):
                    nw110_[i0_11_] = init11_(i0_11_)
                d_655_v43_ = nw110_
                d_657_v44_: _dafny.Map
                d_657_v44_ = _dafny.Map({(d_651_v42_)[default__.safeIndex(117, (d_651_v42_).length(0))]: (d_651_v42_)[default__.safeIndex(117, (d_651_v42_).length(0))]})
                d_658_v45_: D1
                d_658_v45_ = D1_DC3(d_603_v0_, default__.fm1(((d_657_v44_)[(d_606_v3_).f12] if ((d_606_v3_).f12) in (d_657_v44_) else True), d_603_v0_, globalState))
                d_659_v46_: _dafny.Map
                def iife28_(_pat_let4_0):
                    def iife29_(d_660_dt__update__tmp_h1_):
                        def iife30_(_pat_let5_0):
                            def iife31_(d_661_dt__update_hcf3_h0_):
                                return D1_DC3(d_661_dt__update_hcf3_h0_, (d_660_dt__update__tmp_h1_).cf4)
                            return iife31_(_pat_let5_0)
                        return iife30_(not(False))
                    return iife29_(_pat_let4_0)
                d_659_v46_ = _dafny.Map({d_655_v43_: iife28_(d_658_v45_)})
                (globalState).f2 = len(d_659_v46_)
            d_662_v47_: D1
            d_662_v47_ = D1_DC2(d_603_v0_)
            source14_ = d_662_v47_
            if source14_.is_DC2:
                d_663___mcc_h0_ = source14_.cf2
                d_664_cf2_ = d_663___mcc_h0_
                (self).f10 = (0) - ((len((d_604_v1_) + (_dafny.SeqWithoutIsStrInference([(d_604_v1_)[default__.safeIndex(-325, len(d_604_v1_))]]))) if (d_605_v2_).ispropersubset(d_605_v2_) else default__.fm10((d_606_v3_).f12, globalState)))
                d_665_v48_: _dafny.Map
                d_665_v48_ = _dafny.Map({d_664_cf2_: default__.fm10((d_606_v3_).f12, globalState)})
                d_666_v49_: _dafny.Map
                d_666_v49_ = _dafny.Map({d_644_v37_: d_664_cf2_})
                d_667_v50_: _dafny.Map
                d_667_v50_ = _dafny.Map({(d_606_v3_).f12: default__.fm0(len(d_665_v48_), (d_606_v3_).f13, default__.fm11(globalState), d_666_v49_, globalState)})
                (self).f10 = (0) - (len(d_667_v50_))
                d_665_v48_ = (d_665_v48_).set(True, (d_606_v3_).f13)
                d_668_v51_: _dafny.Map
                d_668_v51_ = _dafny.Map({_dafny.Map({not(d_603_v0_): (d_606_v3_).f13}): (d_606_v3_).f12})
                d_668_v51_ = (d_668_v51_).set(d_665_v48_, not (d_603_v0_) or ((d_606_v3_).f12))
            elif source14_.is_DC3:
                d_669___mcc_h1_ = source14_.cf3
                d_670___mcc_h2_ = source14_.cf4
                d_671_cf4_ = d_670___mcc_h2_
                d_672_cf3_ = d_669___mcc_h1_
                d_673_v52_: _dafny.Seq
                d_673_v52_ = _dafny.SeqWithoutIsStrInference([d_672_cf3_, d_603_v0_])
                d_674_v53_: _dafny.MultiSet
                d_674_v53_ = _dafny.MultiSet([(d_606_v3_).f13, self.f10])
                d_675_v54_: D3
                d_675_v54_ = D3_DC8((d_673_v52_) <= (d_673_v52_), (d_606_v3_).f13, ((d_674_v53_)[self.f10] if (self.f10) in (d_674_v53_) else self.f10))
                d_675_v54_ = d_675_v54_
                (globalState).f2 = self.f10
                d_676_v55_: _dafny.MultiSet
                d_676_v55_ = _dafny.MultiSet([d_603_v0_, d_672_cf3_])
                d_677_v56_: _dafny.Map
                d_677_v56_ = _dafny.Map({d_676_v55_: (d_606_v3_).f12})
                d_678_v57_: D3
                d_678_v57_ = D3_DC6((d_606_v3_).f13, d_603_v0_, (d_606_v3_).f12, self.f10, (d_606_v3_).f12)
                d_677_v56_ = (d_677_v56_).set((d_676_v55_).set((d_606_v3_).f12, default__.abs(len(d_605_v2_))), ((d_678_v57_).cf8 if d_603_v0_ else (d_606_v3_).f12))
                d_603_v0_ = (d_606_v3_).f12
            elif True:
                d_679___mcc_h3_ = source14_.cf1
                d_680_cf1_ = d_679___mcc_h3_
                d_681_v58_: _dafny.Map
                d_681_v58_ = _dafny.Map({(d_606_v3_).f13: not((d_606_v3_).f12)})
                d_681_v58_ = (d_681_v58_).set((d_606_v3_).f13, (d_606_v3_).f12)
                d_682_v59_: D1
                d_682_v59_ = D1_DC1((d_606_v3_).f12)
                d_683_v60_: _dafny.Seq
                d_683_v60_ = _dafny.SeqWithoutIsStrInference([d_603_v0_, (d_606_v3_).f12])
                d_684_v61_: _dafny.Map
                d_684_v61_ = _dafny.Map({(d_683_v60_)[default__.safeIndex((0) - (self.f10), len(d_683_v60_))]: not((d_606_v3_).f12)})
                d_680_cf1_ = ((self).fm7(_dafny.Map({(d_606_v3_).f12: d_682_v59_}), d_680_cf1_, (d_606_v3_).f13, globalState)) or ((len(d_684_v61_)) != (self.f10))
                (globalState).f2 = len(d_644_v37_)
                d_685_v62_: _dafny.Array
                def lambda26_(d_686_i6_):
                    return default__.safeDivisionInt(d_686_i6_, 850)

                init12_ = lambda26_
                nw111_ = _dafny.Array(None, 15)
                for i0_12_ in range(nw111_.length(0)):
                    nw111_[i0_12_] = init12_(i0_12_)
                d_685_v62_ = nw111_
                d_687_v63_: _dafny.Map
                d_687_v63_ = _dafny.Map({self.f10: d_685_v62_})
                d_688_v64_: _dafny.Array
                nw112_ = _dafny.Array(None, 11)
                nw112_[int(0)] = d_685_v62_
                nw112_[int(1)] = d_685_v62_
                nw112_[int(2)] = d_685_v62_
                nw112_[int(3)] = d_685_v62_
                nw112_[int(4)] = d_685_v62_
                nw112_[int(5)] = d_685_v62_
                nw112_[int(6)] = ((d_687_v63_)[self.f10] if (self.f10) in (d_687_v63_) else d_685_v62_)
                nw112_[int(7)] = d_685_v62_
                nw112_[int(8)] = d_685_v62_
                nw112_[int(9)] = d_685_v62_
                nw112_[int(10)] = d_685_v62_
                d_688_v64_ = nw112_
                index119_ = default__.safeIndex(513, (d_688_v64_).length(0))
                (d_688_v64_)[index119_] = d_685_v62_
                index120_ = default__.safeIndex(513, (d_688_v64_).length(0))
                (d_688_v64_)[index120_] = d_685_v62_
        elif True:
            d_689_v65_: _dafny.Array
            def lambda27_(d_690_v0_):
                def lambda28_(d_691_i7_):
                    return default__.safeDivisionInt(d_691_i7_, len(_dafny.SeqWithoutIsStrInference([False, d_690_v0_])))

                return lambda28_

            init13_ = lambda27_(d_603_v0_)
            nw113_ = _dafny.Array(None, 5)
            for i0_13_ in range(nw113_.length(0)):
                nw113_[i0_13_] = init13_(i0_13_)
            d_689_v65_ = nw113_
            index121_ = default__.safeIndex(275, (d_689_v65_).length(0))
            (d_689_v65_)[index121_] = self.f10
            d_692_v66_: _dafny.Set
            d_692_v66_ = _dafny.Set({d_603_v0_, d_603_v0_, d_603_v0_})
            d_693_v67_: _dafny.Map
            d_693_v67_ = _dafny.Map({d_603_v0_: self.f10})
            index122_ = default__.safeIndex(275, (d_689_v65_).length(0))
            rhs137_ = d_603_v0_
            rhs138_ = (d_692_v66_).ispropersubset(d_692_v66_)
            rhs139_ = (self.f10) + (-305)
            rhs140_ = (0) - (len(d_693_v67_))
            lhs93_ = d_689_v65_
            lhs94_ = default__.safeIndex(275, (d_689_v65_).length(0))
            lhs95_ = globalState
            d_603_v0_ = rhs137_
            d_603_v0_ = rhs138_
            lhs93_[lhs94_] = rhs139_
            lhs95_.f2 = rhs140_
            d_694_v68_: _dafny.Array
            nw114_ = _dafny.Array(False, 21)
            d_694_v68_ = nw114_
            d_695_v69_: _dafny.Seq
            d_695_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cirg"))
            index123_ = default__.safeIndex(628, (d_694_v68_).length(0))
            (d_694_v68_)[index123_] = (d_695_v69_) == (d_695_v69_)
            d_696_v70_: _dafny.Array
            nw115_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
            d_696_v70_ = nw115_
            d_697_v71_: _dafny.Array
            nw116_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_697_v71_ = nw116_
            d_698_v72_: T0
            nw117_ = C1()
            nw117_.ctor__(d_603_v0_, D3_DC5(d_697_v71_))
            d_698_v72_ = nw117_
            d_699_v73_: _dafny.MultiSet
            d_699_v73_ = _dafny.MultiSet([d_698_v72_])
            d_700_v74_: _dafny.MultiSet
            d_700_v74_ = _dafny.MultiSet([d_603_v0_])
            d_701_v75_: _dafny.Map
            d_701_v75_ = _dafny.Map({d_700_v74_: d_603_v0_})
            d_702_v76_: _dafny.Seq
            d_702_v76_ = _dafny.SeqWithoutIsStrInference([d_603_v0_, d_603_v0_, d_603_v0_, d_603_v0_, d_603_v0_])
            d_703_v77_: _dafny.Map
            d_703_v77_ = _dafny.Map({len(d_701_v75_): _dafny.MultiSet([(d_689_v65_)[default__.safeIndex(275, (d_689_v65_).length(0))], len(d_702_v76_)])})
            d_704_v78_: _dafny.Map
            d_704_v78_ = _dafny.Map({d_703_v77_: d_696_v70_})
            d_705_v81_: _dafny.Map
            d_705_v81_ = _dafny.Map({self.f10: len(d_695_v69_)})
            d_706_v82_: str
            d_706_v82_ = _dafny.CodePoint('a')
            d_707_v83_: _dafny.Set
            d_707_v83_ = _dafny.Set({d_706_v82_})
            index124_ = default__.safeIndex(628, (d_694_v68_).length(0))
            def iife32_():
                coll20_ = _dafny.Map()
                compr_20_: int
                for compr_20_ in _dafny.IntegerRange(-398, 496):
                    d_708_v79_: int = compr_20_
                    if ((-398) <= (d_708_v79_)) and ((d_708_v79_) < (496)):
                        coll20_[(d_708_v79_) - (self.f10)] = _dafny.MultiSet([len(d_692_v66_), (d_689_v65_)[default__.safeIndex(275, (d_689_v65_).length(0))]])
                return _dafny.Map(coll20_)
            def iife33_():
                coll21_ = _dafny.Map()
                compr_21_: int
                for compr_21_ in (d_705_v81_).keys.Elements:
                    d_709_v80_: int = compr_21_
                    if (d_709_v80_) in (d_705_v81_):
                        coll21_[default__.safeModuloInt(d_709_v80_, -284)] = _dafny.MultiSet([self.f10])
                return _dafny.Map(coll21_)
            def iife34_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(-398, 496):
                    d_710_v79_: int = compr_22_
                    if ((-398) <= (d_710_v79_)) and ((d_710_v79_) < (496)):
                        coll22_[(d_710_v79_) - (self.f10)] = _dafny.MultiSet([len(d_692_v66_), (d_689_v65_)[default__.safeIndex(275, (d_689_v65_).length(0))]])
                return _dafny.Map(coll22_)
            def iife35_():
                coll23_ = _dafny.Map()
                compr_23_: int
                for compr_23_ in (d_705_v81_).keys.Elements:
                    d_711_v80_: int = compr_23_
                    if (d_711_v80_) in (d_705_v81_):
                        coll23_[default__.safeModuloInt(d_711_v80_, -284)] = _dafny.MultiSet([self.f10])
                return _dafny.Map(coll23_)
            rhs141_ = (d_605_v2_).isdisjoint(d_605_v2_)
            rhs142_ = ((d_699_v73_)[d_698_v72_] if (d_698_v72_) in (d_699_v73_) else 265)
            rhs143_ = ((d_704_v78_)[(iife32_()
            ) | (iife33_()
            )] if ((iife34_()
            ) | (iife35_()
            )) in (d_704_v78_) else d_696_v70_)
            rhs144_ = ((d_707_v83_ if d_603_v0_ else _dafny.Set({d_706_v82_, _dafny.CodePoint('i')}))).ispropersubset(d_707_v83_)
            rhs145_ = (default__.fm23(d_705_v81_, d_603_v0_, globalState)).isdisjoint(d_692_v66_)
            lhs96_ = d_694_v68_
            lhs97_ = default__.safeIndex(628, (d_694_v68_).length(0))
            lhs96_[lhs97_] = rhs141_
            r0 = rhs142_
            d_696_v70_ = rhs143_
            d_603_v0_ = rhs144_
            d_603_v0_ = rhs145_
            index125_ = default__.safeIndex(628, (d_694_v68_).length(0))
            (d_694_v68_)[index125_] = ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ethwsioyp")))) - ((d_689_v65_)[default__.safeIndex(275, (d_689_v65_).length(0))])) > (((default__.fm24(d_603_v0_, d_702_v76_, (d_694_v68_)[default__.safeIndex(628, (d_694_v68_).length(0))], globalState)) | (_dafny.MultiSet(d_604_v1_))).cardinality)
            d_712_v84_: _dafny.Map
            d_712_v84_ = _dafny.Map({True: False})
            index126_ = default__.safeIndex(628, (d_694_v68_).length(0))
            (d_694_v68_)[index126_] = (self.f10) > (len(d_712_v84_))
            index127_ = default__.safeIndex(628, (d_694_v68_).length(0))
            (d_694_v68_)[index127_] = (d_694_v68_)[default__.safeIndex(628, (d_694_v68_).length(0))]
        (globalState).f2 = default__.fm10(d_603_v0_, globalState)
        d_713_v85_: _dafny.Seq
        d_713_v85_ = _dafny.SeqWithoutIsStrInference([d_603_v0_, True])
        d_714_v86_: _dafny.Map
        d_714_v86_ = _dafny.Map({d_713_v85_: (0) - (self.f10)})
        d_715_v88_: _dafny.Seq
        d_715_v88_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_713_v85_)])
        d_716_v89_: _dafny.Seq
        d_716_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjyrhbv"))
        d_717_v90_: _dafny.Map
        d_717_v90_ = _dafny.Map({d_716_v89_: True})
        d_718_v91_: C2
        nw118_ = C2()
        def iife36_():
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in (d_605_v2_).Elements:
                d_719_v87_: int = compr_24_
                if (d_719_v87_) in (d_605_v2_):
                    coll24_[default__.safeDivisionInt(d_719_v87_, self.f10)] = (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')]))])).cardinality
            return _dafny.Map(coll24_)
        nw118_.ctor__(d_603_v0_, default__.fm0(len(iife36_()
        ), self.f10, d_715_v88_, d_717_v90_, globalState))
        d_718_v91_ = nw118_
        d_720_v92_: _dafny.Seq
        d_720_v92_ = _dafny.SeqWithoutIsStrInference([d_718_v91_, d_718_v91_])
        d_714_v86_ = (d_714_v86_).set(d_713_v85_, len((d_720_v92_) + (d_720_v92_)))
        d_721_v93_: _dafny.MultiSet
        d_721_v93_ = _dafny.MultiSet([len(d_713_v85_)])
        d_722_v94_: _dafny.MultiSet
        d_722_v94_ = _dafny.MultiSet([self.f10, self.f10, len(d_716_v89_), (d_721_v93_).cardinality])
        d_723_v95_: _dafny.MultiSet
        d_723_v95_ = _dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.CodePoint('i')})) for d_724_i9_ in range(default__.abs(567))]))])
        d_725_v96_: _dafny.Array
        nw119_ = _dafny.Array(None, 15)
        nw119_[int(0)] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([22 for d_726_i8_ in range(default__.abs(339))])) + (d_604_v1_))
        nw119_[int(1)] = d_722_v94_
        nw119_[int(2)] = d_722_v94_
        nw119_[int(3)] = _dafny.MultiSet([self.f10, len((self).f11), (_dafny.MultiSet([(d_718_v91_).f15])).cardinality, len(d_605_v2_)])
        nw119_[int(4)] = (d_722_v94_ if False else d_722_v94_)
        nw119_[int(5)] = (d_721_v93_).set(self.f10, default__.abs(self.f10))
        nw119_[int(6)] = (d_722_v94_).intersection(d_721_v93_)
        nw119_[int(7)] = default__.fm24((d_718_v91_).f15, d_713_v85_, d_718_v91_.f14, globalState)
        nw119_[int(8)] = (d_722_v94_).intersection((_dafny.MultiSet([self.f10])).set((d_723_v95_).cardinality, default__.abs(self.f10)))
        nw119_[int(9)] = (d_722_v94_).intersection(d_721_v93_)
        nw119_[int(10)] = (default__.fm24((d_718_v91_).f15, default__.fm18(self.f10, (d_718_v91_).f15, len((_dafny.SeqWithoutIsStrInference([(d_718_v91_).f15])).set(default__.safeIndex(258, len(_dafny.SeqWithoutIsStrInference([(d_718_v91_).f15]))), True)), globalState), (d_718_v91_).f15, globalState)).intersection(d_722_v94_)
        nw119_[int(11)] = _dafny.MultiSet([self.f10, (0) - (self.f10)])
        nw119_[int(12)] = d_721_v93_
        nw119_[int(13)] = d_721_v93_
        nw119_[int(14)] = d_721_v93_
        d_725_v96_ = nw119_
        index128_ = default__.safeIndex(259, (d_725_v96_).length(0))
        (d_725_v96_)[index128_] = _dafny.MultiSet(d_604_v1_)
        index129_ = default__.safeIndex(259, (d_725_v96_).length(0))
        (d_725_v96_)[index129_] = ((_dafny.MultiSet([self.f10, 552])).intersection(d_722_v94_)).intersection(d_721_v93_)
        d_727_v97_: _dafny.Seq
        d_727_v97_ = d_716_v89_
        d_728_v98_: _dafny.Array
        def lambda29_(d_729_i10_):
            return default__.safeModuloInt(d_729_i10_, self.f10)

        init14_ = lambda29_
        nw120_ = _dafny.Array(None, 15)
        for i0_14_ in range(nw120_.length(0)):
            nw120_[i0_14_] = init14_(i0_14_)
        d_728_v98_ = nw120_
        d_730_v99_: _dafny.Map
        d_730_v99_ = _dafny.Map({d_727_v97_: d_728_v98_})
        d_731_v100_: D8
        d_731_v100_ = D8_DC17(d_730_v99_)
        (globalState).f2 = len(((d_730_v99_) | ((d_731_v100_).cf27)) | (d_730_v99_))
        r0 = self.f10
        return r0

    def m2(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_732_v0_: D6
        d_732_v0_ = D6_DC13((self.f10) > ((0) - (self.f10)))
        source15_ = d_732_v0_
        if source15_.is_DC13:
            d_733___mcc_h0_ = source15_.cf22
            d_734_cf22_ = d_733___mcc_h0_
            r0 = d_734_cf22_
            d_735_v1_: _dafny.Array
            nw121_ = _dafny.Array(False, 4)
            d_735_v1_ = nw121_
            d_736_v2_: D7
            d_736_v2_ = D7_DC14(_dafny.Set({d_735_v1_, d_735_v1_, d_735_v1_, d_735_v1_}))
            d_737_v3_: _dafny.Set
            d_737_v3_ = _dafny.Set({d_735_v1_})
            pat_let_tv2_ = d_737_v3_
            def iife37_(_pat_let6_0):
                def iife38_(d_738_dt__update__tmp_h0_):
                    def iife39_(_pat_let7_0):
                        def iife40_(d_739_dt__update_hcf23_h0_):
                            return D7_DC14(d_739_dt__update_hcf23_h0_)
                        return iife40_(_pat_let7_0)
                    return iife39_(pat_let_tv2_)
                return iife38_(_pat_let6_0)
            source16_ = iife37_(d_736_v2_)
            if source16_.is_DC15:
                d_740___mcc_h2_ = source16_.cf24
                d_741___mcc_h3_ = source16_.cf25
                d_742_cf25_ = d_741___mcc_h3_
                d_743_cf24_ = d_740___mcc_h2_
                (globalState).f2 = d_743_cf24_
                (self).f10 = (default__.safeModuloInt(self.f10, 317)) - (((0) - ((0) - (d_743_cf24_))) * (self.f10))
                d_744_v4_: _dafny.Map
                d_744_v4_ = _dafny.Map({d_742_cf25_: d_743_cf24_})
                d_745_v5_: _dafny.MultiSet
                d_745_v5_ = _dafny.MultiSet([d_743_cf24_, self.f10])
                d_746_v6_: _dafny.MultiSet
                d_746_v6_ = _dafny.MultiSet([d_742_cf25_])
                d_747_v7_: D9
                d_747_v7_ = D9_DC19(d_744_v4_)
                d_748_v8_: _dafny.MultiSet
                d_748_v8_ = _dafny.MultiSet([d_744_v4_, (d_744_v4_) | (d_744_v4_), (d_744_v4_).set(d_742_cf25_, self.f10), d_744_v4_, (_dafny.Map({d_734_cf22_: ((d_745_v5_)[((d_746_v6_)[d_734_cf22_] if (d_734_cf22_) in (d_746_v6_) else d_743_cf24_)] if (((d_746_v6_)[d_734_cf22_] if (d_734_cf22_) in (d_746_v6_) else d_743_cf24_)) in (d_745_v5_) else d_743_cf24_)})) | ((d_747_v7_).cf28)])
                d_749_v9_: _dafny.Seq
                d_749_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wiwrmiodu"))
                rhs146_ = d_748_v8_
                rhs147_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjdplvq"))
                d_748_v8_ = rhs146_
                d_749_v9_ = rhs147_
                d_750_v10_: int
                d_751_v11_: bool
                d_752_v12_: bool
                d_753_v13_: _dafny.Map
                out52_: int
                out53_: bool
                out54_: bool
                out55_: _dafny.Map
                out52_, out53_, out54_, out55_ = default__.m0(globalState)
                d_750_v10_ = out52_
                d_751_v11_ = out53_
                d_752_v12_ = out54_
                d_753_v13_ = out55_
            elif source16_.is_DC14:
                d_754___mcc_h4_ = source16_.cf23
                d_755_cf23_ = d_754___mcc_h4_
                d_756_v14_: str
                d_756_v14_ = _dafny.CodePoint('j')
                d_757_v15_: D9
                d_757_v15_ = D9_DC20(d_734_cf22_, d_756_v14_)
                d_758_v16_: _dafny.Array
                nw122_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_758_v16_ = nw122_
                d_759_v17_: D3
                d_759_v17_ = D3_DC5(d_758_v16_)
                d_760_v18_: C1
                nw123_ = C1()
                nw123_.ctor__(((d_757_v15_).cf29) or (d_734_cf22_), d_759_v17_)
                d_760_v18_ = nw123_
                d_761_v19_: D10
                d_761_v19_ = D10_DC22(d_760_v18_)
                pat_let_tv3_ = d_760_v18_
                def iife41_(_pat_let8_0):
                    def iife42_(d_762_dt__update__tmp_h1_):
                        def iife43_(_pat_let9_0):
                            def iife44_(d_763_dt__update_hcf32_h0_):
                                return D10_DC22(d_763_dt__update_hcf32_h0_)
                            return iife44_(_pat_let9_0)
                        return iife43_(pat_let_tv3_)
                    return iife42_(_pat_let8_0)
                d_760_v18_ = (iife41_(d_761_v19_)).cf32
                d_764_v20_: _dafny.Map
                d_764_v20_ = _dafny.Map({self.f10: (self).f11})
                d_765_v21_: _dafny.Seq
                d_765_v21_ = _dafny.SeqWithoutIsStrInference([self.f10, self.f10])
                d_766_v22_: _dafny.Map
                d_766_v22_ = _dafny.Map({d_764_v20_: (d_765_v21_)[default__.safeIndex(self.f10, len(d_765_v21_))]})
                def iife45_():
                    coll25_ = _dafny.Map()
                    compr_25_: int
                    for compr_25_ in ((self).f11).keys.Elements:
                        d_767_v23_: int = compr_25_
                        if (d_767_v23_) in ((self).f11):
                            coll25_[(d_767_v23_) - (self.f10)] = (self).f11
                    return _dafny.Map(coll25_)
                d_766_v22_ = (d_766_v22_).set(iife45_()
                , self.f10)
                d_768_v24_: _dafny.Seq
                d_768_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                d_769_v25_: _dafny.Array
                nw124_ = _dafny.Array(None, 18)
                nw124_[int(0)] = d_765_v21_
                nw124_[int(1)] = default__.fm25(self.f10, self.f10, len(d_768_v24_), self.f10, globalState)
                nw124_[int(2)] = d_765_v21_
                nw124_[int(3)] = d_765_v21_
                nw124_[int(4)] = d_765_v21_
                nw124_[int(5)] = d_765_v21_
                nw124_[int(6)] = _dafny.SeqWithoutIsStrInference([self.f10])
                nw124_[int(7)] = _dafny.SeqWithoutIsStrInference([self.f10])
                nw124_[int(8)] = d_765_v21_
                nw124_[int(9)] = d_765_v21_
                nw124_[int(10)] = d_765_v21_
                nw124_[int(11)] = d_765_v21_
                nw124_[int(12)] = _dafny.SeqWithoutIsStrInference([self.f10, 758])
                nw124_[int(13)] = _dafny.SeqWithoutIsStrInference([self.f10 for d_770_i0_ in range(default__.abs(903))])
                nw124_[int(14)] = d_765_v21_
                nw124_[int(15)] = _dafny.SeqWithoutIsStrInference([self.f10])
                nw124_[int(16)] = _dafny.SeqWithoutIsStrInference([self.f10])
                nw124_[int(17)] = d_765_v21_
                d_769_v25_ = nw124_
                d_771_v26_: C2
                nw125_ = C2()
                nw125_.ctor__((d_760_v18_).f16, (((self).f11)[(0) - (self.f10)] if ((0) - (self.f10)) in ((self).f11) else (d_760_v18_).f16))
                d_771_v26_ = nw125_
                d_772_v27_: _dafny.Map
                d_772_v27_ = _dafny.Map({d_769_v25_: d_771_v26_})
                d_772_v27_ = (d_772_v27_).set(d_769_v25_, d_771_v26_)
                d_773_v28_: _dafny.Array
                def lambda30_(d_774_v21_):
                    def lambda31_(d_775_i1_):
                        return (d_775_i1_) - ((d_774_v21_)[default__.safeIndex(self.f10, len(d_774_v21_))])

                    return lambda31_

                init15_ = lambda30_(d_765_v21_)
                nw126_ = _dafny.Array(None, 26)
                for i0_15_ in range(nw126_.length(0)):
                    nw126_[i0_15_] = init15_(i0_15_)
                d_773_v28_ = nw126_
                index130_ = default__.safeIndex(324, (d_773_v28_).length(0))
                (d_773_v28_)[index130_] = (self.f10) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dq"))))
                d_776_v29_: _dafny.Seq
                d_776_v29_ = _dafny.SeqWithoutIsStrInference([d_734_cf22_])
                index131_ = default__.safeIndex(324, (d_773_v28_).length(0))
                (d_773_v28_)[index131_] = ((default__.fm10((d_771_v26_).f15, globalState)) + (len(d_776_v29_)) if ((d_765_v21_)[default__.safeIndex(self.f10, len(d_765_v21_))]) < ((0) - (self.f10)) else len(d_768_v24_))
            elif True:
                d_777___mcc_h5_ = source16_.cf26
                d_778_cf26_ = d_777___mcc_h5_
                r2 = not((d_734_cf22_) and (d_734_cf22_))
                (globalState).f2 = (187) * (self.f10)
                d_779_v30_: C2
                nw127_ = C2()
                nw127_.ctor__(d_734_cf22_, False)
                d_779_v30_ = nw127_
                d_780_v31_: _dafny.Seq
                d_780_v31_ = _dafny.SeqWithoutIsStrInference([(d_779_v30_).f15])
                d_781_v32_: bool
                d_782_v33_: _dafny.Array
                out56_: bool
                out57_: _dafny.Array
                out56_, out57_ = (d_779_v30_).m3(((d_780_v31_).set(default__.safeIndex(self.f10, len(d_780_v31_)), d_734_cf22_)) + (_dafny.SeqWithoutIsStrInference([not((d_779_v30_).f15)])), globalState)
                d_781_v32_ = out56_
                d_782_v33_ = out57_
            d_783_v34_: _dafny.Seq
            d_783_v34_ = _dafny.SeqWithoutIsStrInference([d_734_cf22_])
            d_784_v35_: D9
            d_784_v35_ = D9_DC20((d_783_v34_) == (d_783_v34_), default__.fm26((((self).f11)[self.f10] if (self.f10) in ((self).f11) else d_734_cf22_), globalState))
            d_785_v36_: _dafny.MultiSet
            d_785_v36_ = _dafny.MultiSet([len(default__.fm27(globalState))])
            d_786_v37_: str
            d_786_v37_ = _dafny.CodePoint('f')
            d_784_v35_ = D9_DC20((_dafny.MultiSet([self.f10])).issubset(d_785_v36_), d_786_v37_)
            d_787_v38_: _dafny.Seq
            d_787_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bicxihw"))
            d_788_v39_: _dafny.Map
            d_788_v39_ = _dafny.Map({d_734_cf22_: self.f10})
            d_789_v40_: _dafny.Map
            d_789_v40_ = _dafny.Map({self.f10: len(_dafny.SeqWithoutIsStrInference([d_786_v37_ for d_790_i2_ in range(default__.abs(-765))]))})
            r1 = (len((d_787_v38_) + ((d_787_v38_).set(default__.safeIndex(617, len(d_787_v38_)), d_786_v37_)))) * ((((d_788_v39_)[d_734_cf22_] if (d_734_cf22_) in (d_788_v39_) else self.f10)) * (((d_789_v40_)[(d_785_v36_).cardinality] if ((d_785_v36_).cardinality) in (d_789_v40_) else self.f10)))
        elif True:
            d_791___mcc_h1_ = source15_.cf21
            d_792_cf21_ = d_791___mcc_h1_
            (globalState).f2 = 447
            (globalState).f2 = self.f10
            d_793_v41_: bool
            d_793_v41_ = False
            d_794_v42_: _dafny.Array
            nw128_ = _dafny.Array(None, 3)
            nw128_[int(0)] = d_793_v41_
            nw128_[int(1)] = d_793_v41_
            nw128_[int(2)] = d_793_v41_
            d_794_v42_ = nw128_
            d_795_v43_: _dafny.Array
            nw129_ = _dafny.Array(None, 15)
            nw129_[int(0)] = d_794_v42_
            nw129_[int(1)] = d_794_v42_
            nw129_[int(2)] = d_794_v42_
            nw129_[int(3)] = d_794_v42_
            nw129_[int(4)] = d_794_v42_
            nw129_[int(5)] = d_794_v42_
            nw129_[int(6)] = d_794_v42_
            nw129_[int(7)] = d_794_v42_
            nw129_[int(8)] = d_794_v42_
            nw129_[int(9)] = d_794_v42_
            nw129_[int(10)] = d_794_v42_
            nw129_[int(11)] = (D11_DC24(d_794_v42_)).cf36
            nw129_[int(12)] = d_794_v42_
            nw129_[int(13)] = d_794_v42_
            nw129_[int(14)] = d_794_v42_
            d_795_v43_ = nw129_
            rhs148_ = d_795_v43_
            rhs149_ = self.f10
            lhs98_ = globalState
            d_795_v43_ = rhs148_
            lhs98_.f2 = rhs149_
            d_796_v44_: int
            d_797_v45_: bool
            d_798_v46_: bool
            d_799_v47_: _dafny.Map
            out58_: int
            out59_: bool
            out60_: bool
            out61_: _dafny.Map
            out58_, out59_, out60_, out61_ = default__.m0(globalState)
            d_796_v44_ = out58_
            d_797_v45_ = out59_
            d_798_v46_ = out60_
            d_799_v47_ = out61_
        d_800_v48_: bool
        d_800_v48_ = False
        d_801_v49_: C0
        nw130_ = C0()
        nw130_.ctor__(d_800_v48_, self.f10)
        d_801_v49_ = nw130_
        d_802_v50_: D1
        d_802_v50_ = D1_DC1(d_800_v48_)
        d_803_v51_: C2
        nw131_ = C2()
        nw131_.ctor__((d_801_v49_).f12, (d_801_v49_).f12)
        d_803_v51_ = nw131_
        d_804_v52_: _dafny.Map
        d_804_v52_ = _dafny.Map({(d_802_v50_).cf1: d_803_v51_})
        d_805_v53_: _dafny.Seq
        d_805_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntqkq"))
        d_806_v54_: _dafny.Set
        d_806_v54_ = _dafny.Set({len(d_804_v52_), (d_801_v49_).f13, (d_801_v49_).f13, len(d_805_v53_), self.f10})
        d_807_v55_: _dafny.Map
        d_807_v55_ = _dafny.Map({not((d_801_v49_).f12): len(d_806_v54_)})
        d_808_v56_: _dafny.Seq
        d_808_v56_ = _dafny.SeqWithoutIsStrInference([d_807_v55_])
        d_809_v57_: _dafny.Array
        nw132_ = _dafny.Array(None, 13)
        nw132_[int(0)] = d_807_v55_
        nw132_[int(1)] = d_807_v55_
        nw132_[int(2)] = d_807_v55_
        nw132_[int(3)] = d_807_v55_
        nw132_[int(4)] = d_807_v55_
        nw132_[int(5)] = ((d_808_v56_)[default__.safeIndex((d_801_v49_).f13, len(d_808_v56_))]) | (d_807_v55_)
        nw132_[int(6)] = d_807_v55_
        nw132_[int(7)] = _dafny.Map({d_800_v48_: (d_801_v49_).f13})
        nw132_[int(8)] = (d_807_v55_) | (_dafny.Map({not(d_800_v48_): self.f10}))
        nw132_[int(9)] = (d_807_v55_) | (d_807_v55_)
        nw132_[int(10)] = _dafny.Map({(d_803_v51_).f15: len(d_805_v53_)})
        nw132_[int(11)] = (d_807_v55_) | (_dafny.Map({(d_803_v51_).f15: (d_801_v49_).f13}))
        nw132_[int(12)] = d_807_v55_
        d_809_v57_ = nw132_
        index132_ = default__.safeIndex(456, (d_809_v57_).length(0))
        (d_809_v57_)[index132_] = d_807_v55_
        d_810_v58_: _dafny.Array
        nw133_ = _dafny.Array(_dafny.CodePoint('D'), 5)
        d_810_v58_ = nw133_
        d_811_v59_: str
        d_811_v59_ = _dafny.CodePoint('g')
        index133_ = default__.safeIndex(966, (d_810_v58_).length(0))
        (d_810_v58_)[index133_] = d_811_v59_
        d_812_v60_: D9
        d_812_v60_ = D9_DC19(d_807_v55_)
        pat_let_tv4_ = d_803_v51_
        pat_let_tv5_ = d_801_v49_
        pat_let_tv6_ = d_801_v49_
        index134_ = default__.safeIndex(456, (d_809_v57_).length(0))
        index135_ = default__.safeIndex(966, (d_810_v58_).length(0))
        def lambda32_(source17_):
            if source17_.is_DC20:
                d_813___mcc_h6_ = source17_.cf29
                d_814___mcc_h7_ = source17_.cf30
                d_815_cf30_ = d_814___mcc_h7_
                d_816_cf29_ = d_813___mcc_h6_
                return (pat_let_tv4_).f15
            elif source17_.is_DC19:
                d_817___mcc_h8_ = source17_.cf28
                d_818_cf28_ = d_817___mcc_h8_
                return (pat_let_tv5_).f12
            elif True:
                d_819___mcc_h9_ = source17_.cf31
                d_820_cf31_ = d_819___mcc_h9_
                return (self.f10) != ((pat_let_tv6_).f13)

        rhs150_ = lambda32_((d_812_v60_ if d_800_v48_ else d_812_v60_))
        rhs151_ = (d_807_v55_) | (d_807_v55_)
        rhs152_ = d_811_v59_
        rhs153_ = ((self.f10) + (self.f10)) - (default__.fm19(len(d_805_v53_), globalState))
        lhs99_ = d_809_v57_
        lhs100_ = default__.safeIndex(456, (d_809_v57_).length(0))
        lhs101_ = d_810_v58_
        lhs102_ = default__.safeIndex(966, (d_810_v58_).length(0))
        lhs103_ = self
        r0 = rhs150_
        lhs99_[lhs100_] = rhs151_
        lhs101_[lhs102_] = rhs152_
        lhs103_.f10 = rhs153_
        (globalState).f2 = (d_801_v49_).f13
        d_821_v61_: _dafny.Seq
        d_821_v61_ = _dafny.SeqWithoutIsStrInference([d_805_v53_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjimyu"))])
        d_822_v62_: int
        d_822_v62_ = (0) - ((d_801_v49_).f13)
        source18_ = D1_DC3((_dafny.SeqWithoutIsStrInference([(D11_DC25((d_810_v58_)[default__.safeIndex(966, (d_810_v58_).length(0))])).cf37 for d_823_i3_ in range(default__.abs(924))])) != ((d_821_v61_)[default__.safeIndex(self.f10, len(d_821_v61_))]), d_822_v62_)
        if source18_.is_DC2:
            d_824___mcc_h10_ = source18_.cf2
            d_825_cf2_ = d_824___mcc_h10_
            d_800_v48_ = (d_801_v49_).f12
            (globalState).f2 = ((d_801_v49_).f13 if (d_801_v49_).f12 else (len(_dafny.Map({(d_801_v49_).f13: d_803_v51_.f14}))) - ((d_801_v49_).f13))
            (globalState).f2 = (d_801_v49_).f13
            d_826_v63_: _dafny.Seq
            d_826_v63_ = _dafny.SeqWithoutIsStrInference([(d_803_v51_).f15, d_803_v51_.f14, (d_803_v51_).f15, d_800_v48_])
            d_827_v64_: C2
            nw134_ = C2()
            nw134_.ctor__((d_826_v63_) != (d_826_v63_), (d_801_v49_).f12)
            d_827_v64_ = nw134_
        elif source18_.is_DC3:
            d_828___mcc_h11_ = source18_.cf3
            d_829___mcc_h12_ = source18_.cf4
            d_830_cf4_ = d_829___mcc_h12_
            d_831_cf3_ = d_828___mcc_h11_
            d_832_v65_: C0
            nw135_ = C0()
            nw135_.ctor__(d_831_cf3_, self.f10)
            d_832_v65_ = nw135_
            d_833_v66_: D10
            d_833_v66_ = D10_DC23((d_801_v49_).f13, (d_801_v49_).f12, d_831_cf3_)
            d_833_v66_ = d_833_v66_
            d_811_v59_ = d_811_v59_
            d_800_v48_ = d_831_cf3_
        elif True:
            d_834___mcc_h13_ = source18_.cf1
            d_835_cf1_ = d_834___mcc_h13_
            d_836_v67_: _dafny.Set
            d_836_v67_ = _dafny.Set({(d_801_v49_).f12})
            (d_803_v51_).f14 = (not(d_835_cf1_)) or ((d_836_v67_).isdisjoint(d_836_v67_))
            r3 = not((d_803_v51_).f15)
            d_837_v68_: _dafny.Array
            def lambda33_(d_838_v49_):
                def lambda34_(d_839_i4_):
                    return not((d_838_v49_).f12)

                return lambda34_

            init16_ = lambda33_(d_801_v49_)
            nw136_ = _dafny.Array(None, 21)
            for i0_16_ in range(nw136_.length(0)):
                nw136_[i0_16_] = init16_(i0_16_)
            d_837_v68_ = nw136_
            d_837_v68_ = d_837_v68_
            d_840_v69_: int
            d_841_v70_: bool
            d_842_v71_: bool
            d_843_v72_: _dafny.Map
            out62_: int
            out63_: bool
            out64_: bool
            out65_: _dafny.Map
            out62_, out63_, out64_, out65_ = default__.m0(globalState)
            d_840_v69_ = out62_
            d_841_v70_ = out63_
            d_842_v71_ = out64_
            d_843_v72_ = out65_
        (d_803_v51_).f14 = (d_803_v51_).f15
        r0 = ((d_801_v49_).f13) <= ((d_801_v49_).f13)
        r1 = (d_801_v49_).f13
        r2 = d_800_v48_
        d_844_v73_: _dafny.MultiSet
        d_844_v73_ = _dafny.MultiSet([d_800_v48_])
        r3 = (d_803_v51_).fm13(len((_dafny.SeqWithoutIsStrInference([d_803_v51_.f14])) + ((_dafny.SeqWithoutIsStrInference([d_800_v48_, not(d_803_v51_.f14), not((d_801_v49_).f12), (d_803_v51_).f15])).set(default__.safeIndex((0) - ((d_801_v49_).f13), len(_dafny.SeqWithoutIsStrInference([d_800_v48_, not(d_803_v51_.f14), not((d_801_v49_).f12), (d_803_v51_).f15]))), (d_803_v51_).f15))), (d_844_v73_).issubset(d_844_v73_), 22, globalState)
        return r0, r1, r2, r3

