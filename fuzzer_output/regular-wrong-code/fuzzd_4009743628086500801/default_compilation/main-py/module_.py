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
    def fm3(globalState):
        return ((_dafny.Set({_dafny.Set({False, True}), _dafny.Set({True, True})})) | (_dafny.Set({_dafny.Set({False, True}), _dafny.Set({False}), _dafny.Set({True, False}), _dafny.Set({False})}))) - ((_dafny.Set({_dafny.Set({False}), _dafny.Set({False})})))

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ed"))

    @staticmethod
    def fm5(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (_dafny.MultiSet([817])).Elements:
                d_1_v0_: int = compr_0_
                if (d_1_v0_) in (_dafny.MultiSet([817])):
                    coll0_[(d_1_v0_) * (len(_dafny.Set({True})))] = (_dafny.MultiSet([False, False])).cardinality
            return _dafny.Map(coll0_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knr")): _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_0_i0_ in range(default__.abs(220))])), 771])})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvrf")): _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(iife0_()
        )]))])}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dujbkapsg")): _dafny.SeqWithoutIsStrInference([733 for d_2_i1_ in range(default__.abs(-864))])}))

    @staticmethod
    def fm8(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uh"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_3_i0_ in range(default__.abs(-759))]))) != ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snaam"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uupaiodb"))))

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_4_i0_ in range(default__.abs(133))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bndem")))

    @staticmethod
    def fm12(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([False]) if True else _dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def fm14(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(-535, 990):
                d_5_v0_: int = compr_1_
                if ((-535) <= (d_5_v0_)) and ((d_5_v0_) < (990)):
                    def iife2_():
                        coll2_ = _dafny.Map()
                        compr_2_: int
                        for compr_2_ in _dafny.IntegerRange(286, -321):
                            d_8_v1_: int = compr_2_
                            if ((286) <= (d_8_v1_)) and ((d_8_v1_) < (-321)):
                                coll2_[default__.safeDivisionInt(d_8_v1_, len(_dafny.Map({_dafny.CodePoint('q'): len(_dafny.Map({False: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srsgjqmes"))))}))})))] = 32
                        return _dafny.Map(coll2_)
                    coll1_[(d_5_v0_) + (-90)] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-907: -667})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_6_i0_ in range(default__.abs(521))])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_7_i1_ in range(default__.abs(225))])), (_dafny.MultiSet([iife2_()
                    ])).cardinality]))
            return _dafny.Map(coll1_)
        return len((_dafny.SeqWithoutIsStrInference([iife1_()
        , _dafny.Map({len(_dafny.Map({not(True): False})): len(_dafny.Map({-841: (_dafny.MultiSet([True])).cardinality}))}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): 115})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({len((D15_DC33(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rcx")))).cf54): 692})])))

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        return (_dafny.Map({True: not(True)})) | (_dafny.Map({True: True}))

    @staticmethod
    def fm16(p0, globalState):
        return ((_dafny.Set({False, not(not(True))})) - (_dafny.Set({True, False}))) - (_dafny.Set({True, False}))

    @staticmethod
    def fm19(p0, p1, globalState):
        source0_ = D9_DC18(_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvkyd"))): len(_dafny.SeqWithoutIsStrInference([609, -451]))}): _dafny.Set({not(True)})}))
        if source0_.is_DC19:
            d_9___mcc_h0_ = source0_.cf29
            d_10___mcc_h1_ = source0_.cf30
            d_11_cf30_ = d_10___mcc_h1_
            d_12_cf29_ = d_9___mcc_h0_
            return (0) - (d_12_cf29_)
        elif source0_.is_DC20:
            d_13___mcc_h2_ = source0_.cf31
            d_14___mcc_h3_ = source0_.cf32
            d_15___mcc_h4_ = source0_.cf33
            d_16_cf33_ = d_15___mcc_h4_
            d_17_cf32_ = d_14___mcc_h3_
            d_18_cf31_ = d_13___mcc_h2_
            return len((_dafny.Set({_dafny.SeqWithoutIsStrInference([0]), _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gynrfk")))), len(_dafny.Set({not(d_16_cf33_)})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvp")))])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([432 for d_19_i0_ in range(default__.abs(38))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_18_cf31_])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gynpbladr")))])})))
        elif True:
            d_20___mcc_h5_ = source0_.cf28
            d_21_cf28_ = d_20___mcc_h5_
            return ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyoldvcp"))), 978, 255, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iflkrw"))), len(_dafny.SeqWithoutIsStrInference([True, True]))])).cardinality) + ((_dafny.MultiSet([False])).cardinality)

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(-820, 382):
                d_23_v0_: int = compr_3_
                if ((-820) <= (d_23_v0_)) and ((d_23_v0_) < (382)):
                    coll3_ = coll3_.union(_dafny.Set([default__.safeDivisionInt(d_23_v0_, 599)]))
            return _dafny.Set(coll3_)
        return ((_dafny.Set({len(_dafny.Map({False: len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqunqvo")))): len(_dafny.SeqWithoutIsStrInference([-829]))}))})), len(_dafny.SeqWithoutIsStrInference([D1_DC4(82, _dafny.CodePoint('y'), False, 805) for d_22_i0_ in range(default__.abs(495))]))})).intersection(_dafny.Set({314}))) - (iife3_()
        )

    @staticmethod
    def fm22(p0, p1, globalState):
        return (D15_DC33(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xebgvpgac")))).cf54

    @staticmethod
    def fm23(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([932, (len(_dafny.Set({len(_dafny.Set({False})), len((_dafny.Map({False: 563})))})) if True else len(_dafny.Map({len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([True, False])))})): 265}))), len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality, len(_dafny.SeqWithoutIsStrInference([False, False]))])), -658, len(_dafny.Map({D10_DC21(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False]))): False}))])

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        return default__.safeModuloInt(918, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkmbe"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))))

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vogq")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))])

    @staticmethod
    def fm26(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([False, True, False, (D1_DC4(len(_dafny.Set({(D15_DC34(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghf")), 164, False, D10_DC22(), 693)).cf57, True})), _dafny.CodePoint('c'), True, 228)).cf8])) + ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm27(globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jac"))]) if not((D15_DC32(False)).cf53) else _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwf"))]))) | ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "en"))])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sicipu"))])))

    @staticmethod
    def fm28(p0, p1, globalState):
        if (216) != (-12):
            return (_dafny.Map({not(True): (D15_DC33(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbkybdo")))).cf54})) | (_dafny.Map({not(not(not(not(not(False))))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vudhbc"))}))
        elif True:
            return (_dafny.Map({not(False): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpc"))})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgwlc"))}))

    @staticmethod
    def fm29(globalState):
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fykfcgc"))): len(_dafny.Map({True: True}))})) | (_dafny.Map({-242: 517}))) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iglpstssc"))): -663}))

    @staticmethod
    def fm30(p0, p1, p2, globalState):
        source1_ = _dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmoihfr"))])).cardinality])): 715}): True})
        d_24___mcc_h0_ = source1_
        d_25_cf18_ = d_24___mcc_h0_
        return _dafny.CodePoint('e')

    @staticmethod
    def fm31(p0, globalState):
        return ((_dafny.MultiSet([-55, len(_dafny.Map({585: 832})), (0) - (-24), (0) - ((_dafny.MultiSet([True])).cardinality), 249])) | (_dafny.MultiSet([697]))) | ((_dafny.MultiSet([747, len(_dafny.SeqWithoutIsStrInference([False, True]))])).intersection(_dafny.MultiSet([34, 873, -285, len(_dafny.Set({True}))])))

    @staticmethod
    def fm32(p0, globalState):
        def iife4_():
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(-447, 500):
                    d_27_v1_: int = compr_6_
                    if ((-447) <= (d_27_v1_)) and ((d_27_v1_) < (500)):
                        coll6_[(d_27_v1_) * (226)] = len(_dafny.Map({False: True}))
                return _dafny.Map(coll6_)
            coll4_ = _dafny.Map()
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(-447, 500):
                    d_27_v1_: int = compr_5_
                    if ((-447) <= (d_27_v1_)) and ((d_27_v1_) < (500)):
                        coll5_[(d_27_v1_) * (226)] = len(_dafny.Map({False: True}))
                return _dafny.Map(coll5_)
            compr_4_: _dafny.Map
            for compr_4_ in ((_dafny.SeqWithoutIsStrInference([_dafny.Map({164: -392}) for d_26_i0_ in range(default__.abs(138))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykxkeccls"))): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrtrgqvhp"))))}), iife5_()
            ]))).Elements:
                d_28_v0_: _dafny.Map = compr_4_
                if (d_28_v0_) in ((_dafny.SeqWithoutIsStrInference([_dafny.Map({164: -392}) for d_26_i0_ in range(default__.abs(138))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykxkeccls"))): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrtrgqvhp"))))}), iife6_()
                ]))):
                    coll4_[d_28_v0_] = _dafny.Set({not(False), False, not(not(False))})
            return _dafny.Map(coll4_)
        return iife4_()
        

    @staticmethod
    def fm33(p0, globalState):
        return False

    @staticmethod
    def fm34(p0, p1, globalState):
        source2_ = D1_DC2(False)
        if source2_.is_DC3:
            d_29___mcc_h0_ = source2_.cf5
            d_30_cf5_ = d_29___mcc_h0_
            return D1_DC2(False)
        elif source2_.is_DC4:
            d_31___mcc_h1_ = source2_.cf6
            d_32___mcc_h2_ = source2_.cf7
            d_33___mcc_h3_ = source2_.cf8
            d_34___mcc_h4_ = source2_.cf9
            d_35_cf9_ = d_34___mcc_h4_
            d_36_cf8_ = d_33___mcc_h3_
            d_37_cf7_ = d_32___mcc_h2_
            d_38_cf6_ = d_31___mcc_h1_
            return D1_DC2(d_36_cf8_)
        elif True:
            d_39___mcc_h5_ = source2_.cf4
            d_40_cf4_ = d_39___mcc_h5_
            return D1_DC2(d_40_cf4_)

    @staticmethod
    def fm35(p0, p1, p2, p3, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: str
            for compr_7_ in ((_dafny.Map({_dafny.CodePoint('c'): 902})) | (_dafny.Map({_dafny.CodePoint('y'): 712}))).keys.Elements:
                d_41_v0_: str = compr_7_
                if (d_41_v0_) in ((_dafny.Map({_dafny.CodePoint('c'): 902})) | (_dafny.Map({_dafny.CodePoint('y'): 712}))):
                    coll7_[d_41_v0_] = not(False)
            return _dafny.Map(coll7_)
        return iife7_()
        

    @staticmethod
    def fm36(globalState):
        return D1_DC4((522) - (len(_dafny.SeqWithoutIsStrInference([-290]))), (D1_DC4(721, _dafny.CodePoint('d'), False, 427)).cf7, not(not(False)), (579 if True else 602))

    @staticmethod
    def m6(globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: T0 = None
        d_42_v0_: C8
        nw0_ = C8()
        nw0_.ctor__()
        d_42_v0_ = nw0_
        nw1_ = C8()
        nw1_.ctor__()
        d_42_v0_ = nw1_
        d_43_v1_: _dafny.Array
        nw2_ = _dafny.Array(_dafny.MultiSet({}), 29)
        d_43_v1_ = nw2_
        d_44_v2_: bool
        d_44_v2_ = False
        d_45_v3_: _dafny.Map
        d_45_v3_ = _dafny.Map({d_44_v2_: _dafny.CodePoint('i')})
        d_46_v4_: str
        d_46_v4_ = _dafny.CodePoint('v')
        d_47_v5_: _dafny.MultiSet
        d_47_v5_ = _dafny.MultiSet([d_44_v2_])
        d_48_v6_: _dafny.Map
        d_48_v6_ = _dafny.Map({((d_45_v3_)[d_44_v2_] if (d_44_v2_) in (d_45_v3_) else d_46_v4_): d_47_v5_})
        index0_ = default__.safeIndex(285, (d_43_v1_).length(0))
        (d_43_v1_)[index0_] = ((d_48_v6_)[d_46_v4_] if (d_46_v4_) in (d_48_v6_) else d_47_v5_)
        d_49_v7_: _dafny.Map
        d_49_v7_ = _dafny.Map({d_44_v2_: d_44_v2_})
        d_50_v8_: int
        d_50_v8_ = 94
        d_51_v9_: _dafny.Map
        d_51_v9_ = _dafny.Map({d_50_v8_: d_44_v2_})
        index1_ = default__.safeIndex(285, (d_43_v1_).length(0))
        (d_43_v1_)[index1_] = (d_42_v0_).fm0((d_49_v7_) == (d_49_v7_), (not(d_44_v2_) if ((d_51_v9_)[d_50_v8_] if (d_50_v8_) in (d_51_v9_) else False) else d_44_v2_), d_50_v8_, d_50_v8_, globalState)
        d_52_v10_: _dafny.Array
        def lambda0_(d_53_i0_):
            return (d_53_i0_) * (800)

        init0_ = lambda0_
        nw3_ = _dafny.Array(None, 19)
        for i0_0_ in range(nw3_.length(0)):
            nw3_[i0_0_] = init0_(i0_0_)
        d_52_v10_ = nw3_
        nw4_ = _dafny.Array(int(0), 9)
        d_52_v10_ = nw4_
        d_42_v0_ = d_42_v0_
        d_54_v11_: _dafny.Seq
        d_54_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ny"))
        d_50_v8_ = (len((d_54_v11_).set(default__.safeIndex(d_50_v8_, len(d_54_v11_)), _dafny.CodePoint('y'))) if d_44_v2_ else d_50_v8_)
        d_55_v12_: D4
        d_55_v12_ = D4_DC11(D4_DC9(d_52_v10_))
        d_56_v13_: D4
        d_56_v13_ = D4_DC11(d_55_v12_)
        d_57_v14_: _dafny.Map
        d_57_v14_ = _dafny.Map({d_56_v13_: default__.fm19(False, 961, globalState)})
        d_58_v15_: _dafny.Seq
        d_58_v15_ = _dafny.SeqWithoutIsStrInference([d_50_v8_])
        d_44_v2_ = not((((d_57_v14_)[d_56_v13_] if (d_56_v13_) in (d_57_v14_) else (d_58_v15_)[default__.safeIndex(d_50_v8_, len(d_58_v15_))])) <= (d_50_v8_))
        r0 = (d_54_v11_) + (d_54_v11_)
        d_59_v16_: T0
        nw5_ = C8()
        nw5_.ctor__()
        d_59_v16_ = nw5_
        r1 = d_59_v16_
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_60_v0_: _dafny.Array
        nw6_ = _dafny.Array(False, 17)
        d_60_v0_ = nw6_
        d_61_globalState_: GlobalState
        nw7_ = GlobalState()
        nw7_.ctor__(d_60_v0_, -481)
        d_61_globalState_ = nw7_
        d_62_v1_: bool
        d_62_v1_ = False
        if d_62_v1_:
            d_63_v2_: int
            d_63_v2_ = 7
            d_63_v2_ = 547
            d_64_v3_: _dafny.MultiSet
            d_64_v3_ = _dafny.MultiSet([d_62_v1_, False])
            d_65_v4_: _dafny.Array
            nw8_ = _dafny.Array(int(0), 13)
            d_65_v4_ = nw8_
            d_66_v5_: C6
            nw9_ = C6()
            nw9_.ctor__((d_64_v3_).set(d_62_v1_, default__.abs(d_63_v2_)), d_65_v4_)
            d_66_v5_ = nw9_
            d_67_v6_: _dafny.Seq
            d_67_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yymfboxka"))
            d_67_v6_ = (d_67_v6_) + ((default__.fm4(d_62_v1_, d_62_v1_, d_62_v1_, d_61_globalState_)) + (d_67_v6_))
            d_68_v7_: _dafny.Seq
            d_68_v7_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kahofehgp"))])
            d_69_v8_: C5
            nw10_ = C5()
            nw10_.ctor__((d_68_v7_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xiko"))), len(d_68_v7_))])
            d_69_v8_ = nw10_
            d_70_v9_: _dafny.Array
            nw11_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_70_v9_ = nw11_
            index2_ = default__.safeIndex(878, (d_70_v9_).length(0))
            (d_70_v9_)[index2_] = d_60_v0_
            index3_ = default__.safeIndex(878, (d_70_v9_).length(0))
            (d_70_v9_)[index3_] = d_60_v0_
        elif True:
            d_71_v10_: int
            d_71_v10_ = 849
            d_72_v11_: _dafny.Map
            d_72_v11_ = _dafny.Map({d_71_v10_: d_71_v10_})
            d_73_v12_: _dafny.Seq
            d_73_v12_ = _dafny.SeqWithoutIsStrInference([d_62_v1_, d_62_v1_])
            d_74_v13_: _dafny.Seq
            d_74_v13_ = _dafny.SeqWithoutIsStrInference([d_71_v10_, d_71_v10_])
            d_75_v14_: _dafny.Map
            d_75_v14_ = _dafny.Map({d_71_v10_: d_62_v1_})
            d_72_v11_ = (d_72_v11_).set(len((d_73_v12_).set(default__.safeIndex(len(d_74_v13_), len(d_73_v12_)), ((d_75_v14_)[d_71_v10_] if (d_71_v10_) in (d_75_v14_) else d_62_v1_))), d_71_v10_)
            d_76_v15_: str
            d_76_v15_ = _dafny.CodePoint('f')
            d_77_v16_: _dafny.Array
            nw12_ = _dafny.Array(None, 8)
            nw12_[int(0)] = d_76_v15_
            nw12_[int(1)] = d_76_v15_
            nw12_[int(2)] = _dafny.CodePoint('f')
            nw12_[int(3)] = d_76_v15_
            nw12_[int(4)] = d_76_v15_
            nw12_[int(5)] = d_76_v15_
            nw12_[int(6)] = d_76_v15_
            nw12_[int(7)] = default__.fm30(default__.fm22(d_62_v1_, (0) - (d_71_v10_), d_61_globalState_), d_71_v10_, 581, d_61_globalState_)
            d_77_v16_ = nw12_
            d_78_v17_: _dafny.Seq
            d_78_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
            d_79_v18_: _dafny.Set
            d_79_v18_ = _dafny.Set({d_62_v1_})
            index4_ = default__.safeIndex(213, (d_77_v16_).length(0))
            (d_77_v16_)[index4_] = (d_78_v17_)[default__.safeIndex(len(d_79_v18_), len(d_78_v17_))]
            index5_ = default__.safeIndex(213, (d_77_v16_).length(0))
            (d_77_v16_)[index5_] = d_76_v15_
            index6_ = default__.safeIndex(623, (d_60_v0_).length(0))
            (d_60_v0_)[index6_] = d_62_v1_
            d_80_v19_: T1
            nw13_ = C4()
            nw13_.ctor__(232)
            d_80_v19_ = nw13_
            index7_ = default__.safeIndex(623, (d_60_v0_).length(0))
            rhs0_ = d_62_v1_
            rhs1_ = d_71_v10_
            rhs2_ = d_78_v17_
            rhs3_ = d_62_v1_
            rhs4_ = d_80_v19_
            lhs0_ = d_60_v0_
            lhs1_ = default__.safeIndex(623, (d_60_v0_).length(0))
            d_62_v1_ = rhs0_
            d_71_v10_ = rhs1_
            d_78_v17_ = rhs2_
            lhs0_[lhs1_] = rhs3_
            d_80_v19_ = rhs4_
            d_81_v20_: _dafny.Array
            nw14_ = _dafny.Array(None, 14)
            nw14_[int(0)] = (d_60_v0_)[default__.safeIndex(623, (d_60_v0_).length(0))]
            nw14_[int(1)] = (d_60_v0_)[default__.safeIndex(623, (d_60_v0_).length(0))]
            nw14_[int(2)] = (d_60_v0_)[default__.safeIndex(623, (d_60_v0_).length(0))]
            nw14_[int(3)] = d_62_v1_
            nw14_[int(4)] = (d_60_v0_)[default__.safeIndex(623, (d_60_v0_).length(0))]
            nw14_[int(5)] = (d_60_v0_)[default__.safeIndex(623, (d_60_v0_).length(0))]
            nw14_[int(6)] = d_62_v1_
            nw14_[int(7)] = (d_60_v0_)[default__.safeIndex(623, (d_60_v0_).length(0))]
            nw14_[int(8)] = (d_60_v0_)[default__.safeIndex(623, (d_60_v0_).length(0))]
            nw14_[int(9)] = d_62_v1_
            nw14_[int(10)] = d_62_v1_
            nw14_[int(11)] = True
            nw14_[int(12)] = d_62_v1_
            nw14_[int(13)] = d_62_v1_
            d_81_v20_ = nw14_
            d_82_v21_: _dafny.Seq
            d_82_v21_ = _dafny.SeqWithoutIsStrInference([d_81_v20_, d_81_v20_])
            d_83_v22_: _dafny.Array
            nw15_ = _dafny.Array(None, 26)
            nw15_[int(0)] = d_81_v20_
            nw15_[int(1)] = d_81_v20_
            nw15_[int(2)] = d_81_v20_
            nw15_[int(3)] = d_81_v20_
            nw15_[int(4)] = d_81_v20_
            nw15_[int(5)] = d_81_v20_
            nw15_[int(6)] = d_81_v20_
            nw15_[int(7)] = d_81_v20_
            nw15_[int(8)] = d_81_v20_
            nw15_[int(9)] = d_81_v20_
            nw15_[int(10)] = d_81_v20_
            nw15_[int(11)] = (d_82_v21_)[default__.safeIndex(d_71_v10_, len(d_82_v21_))]
            nw15_[int(12)] = d_81_v20_
            nw15_[int(13)] = d_81_v20_
            nw15_[int(14)] = d_81_v20_
            nw15_[int(15)] = d_81_v20_
            nw15_[int(16)] = d_81_v20_
            nw15_[int(17)] = d_81_v20_
            nw15_[int(18)] = d_81_v20_
            nw15_[int(19)] = d_81_v20_
            nw15_[int(20)] = d_81_v20_
            nw15_[int(21)] = d_81_v20_
            nw15_[int(22)] = d_81_v20_
            nw15_[int(23)] = d_81_v20_
            nw15_[int(24)] = d_81_v20_
            nw15_[int(25)] = d_81_v20_
            d_83_v22_ = nw15_
            d_84_v23_: bool
            d_85_v24_: _dafny.Map
            out0_: bool
            out1_: _dafny.Map
            out0_, out1_ = (d_80_v19_).m0(d_83_v22_, d_61_globalState_)
            d_84_v23_ = out0_
            d_85_v24_ = out1_
            d_86_v25_: C4
            nw16_ = C4()
            nw16_.ctor__(d_71_v10_)
            d_86_v25_ = nw16_
            d_86_v25_ = d_86_v25_
        d_87_v26_: int
        d_87_v26_ = 542
        d_88_v27_: _dafny.Map
        d_88_v27_ = _dafny.Map({d_87_v26_: d_87_v26_})
        d_89_v28_: _dafny.Seq
        d_89_v28_ = _dafny.SeqWithoutIsStrInference([d_87_v26_])
        d_90_v29_: _dafny.Seq
        d_90_v29_ = _dafny.SeqWithoutIsStrInference([(d_89_v28_)[default__.safeIndex(default__.fm24(904, d_62_v1_, d_62_v1_, d_87_v26_, d_61_globalState_), len(d_89_v28_))]])
        d_91_v30_: _dafny.Map
        d_91_v30_ = _dafny.Map({d_87_v26_: d_89_v28_})
        d_87_v26_ = ((d_88_v27_)[len(d_91_v30_)] if (len(d_91_v30_)) in (d_88_v27_) else (d_87_v26_ if True else d_87_v26_))
        d_92_v31_: str
        d_92_v31_ = _dafny.CodePoint('a')
        source3_ = D1_DC4(d_87_v26_, d_92_v31_, d_62_v1_, d_87_v26_)
        if source3_.is_DC3:
            d_93___mcc_h0_ = source3_.cf5
            d_94_cf5_ = d_93___mcc_h0_
            d_95_v32_: _dafny.Seq
            d_96_v33_: T0
            out2_: _dafny.Seq
            out3_: T0
            out2_, out3_ = default__.m6(d_61_globalState_)
            d_95_v32_ = out2_
            d_96_v33_ = out3_
            d_97_v34_: D9
            d_97_v34_ = D9_DC19(((d_89_v28_)[default__.safeIndex(760, len(d_89_v28_))]) + ((_dafny.MultiSet([d_62_v1_])).cardinality), d_94_cf5_)
            d_98_v35_: D2
            d_98_v35_ = D2_DC6(d_94_cf5_, d_62_v1_)
            rhs5_ = d_97_v34_
            rhs6_ = default__.safeModuloInt(d_87_v26_, (0) - (default__.fm19((d_98_v35_).cf12, d_94_cf5_, d_61_globalState_)))
            d_97_v34_ = rhs5_
            d_94_cf5_ = rhs6_
            d_99_v36_: _dafny.Map
            d_99_v36_ = _dafny.Map({d_62_v1_: d_62_v1_})
            d_100_v37_: D15
            d_100_v37_ = D15_DC31(d_99_v36_)
            d_99_v36_ = (d_100_v37_).cf52
            d_101_v38_: _dafny.Seq
            d_101_v38_ = _dafny.SeqWithoutIsStrInference([d_62_v1_, d_62_v1_])
            d_101_v38_ = (d_101_v38_).set(default__.safeIndex(d_87_v26_, len(d_101_v38_)), d_62_v1_)
        elif source3_.is_DC4:
            d_102___mcc_h1_ = source3_.cf6
            d_103___mcc_h2_ = source3_.cf7
            d_104___mcc_h3_ = source3_.cf8
            d_105___mcc_h4_ = source3_.cf9
            d_106_cf9_ = d_105___mcc_h4_
            d_107_cf8_ = d_104___mcc_h3_
            d_108_cf7_ = d_103___mcc_h2_
            d_109_cf6_ = d_102___mcc_h1_
            d_110_v39_: C3
            nw17_ = C3()
            nw17_.ctor__()
            d_110_v39_ = nw17_
            d_111_v40_: _dafny.Map
            d_111_v40_ = _dafny.Map({default__.fm14(d_87_v26_, -263, d_61_globalState_): (d_110_v39_ if d_62_v1_ else d_110_v39_)})
            d_112_v41_: _dafny.Seq
            d_112_v41_ = _dafny.SeqWithoutIsStrInference([d_107_cf8_, d_107_cf8_])
            d_111_v40_ = (d_111_v40_).set(len((d_112_v41_) + (d_112_v41_)), d_110_v39_)
            d_87_v26_ = d_109_cf6_
            d_107_cf8_ = ((d_106_cf9_) + (d_109_cf6_)) != (d_87_v26_)
        elif True:
            d_113___mcc_h5_ = source3_.cf4
            d_114_cf4_ = d_113___mcc_h5_
            d_62_v1_ = (d_87_v26_) != (d_87_v26_)
            d_115_v42_: _dafny.Seq
            d_115_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aptekw"))
            d_116_v43_: _dafny.MultiSet
            d_116_v43_ = _dafny.MultiSet([278, d_87_v26_, len(d_115_v42_)])
            d_114_cf4_ = default__.fm8(d_116_v43_, d_61_globalState_)
            d_117_v44_: _dafny.Map
            d_117_v44_ = _dafny.Map({d_114_cf4_: default__.fm4(not(d_114_cf4_), d_62_v1_, d_114_cf4_, d_61_globalState_)})
            d_118_v45_: _dafny.MultiSet
            d_118_v45_ = _dafny.MultiSet([d_114_cf4_])
            d_119_v46_: C1
            nw18_ = C1()
            nw18_.ctor__(d_118_v45_, len(default__.fm35(False, d_62_v1_, d_62_v1_, d_87_v26_, d_61_globalState_)))
            d_119_v46_ = nw18_
            d_120_v47_: _dafny.Array
            nw19_ = _dafny.Array(None, 2)
            nw19_[int(0)] = d_119_v46_
            nw19_[int(1)] = d_119_v46_
            d_120_v47_ = nw19_
            d_121_v48_: _dafny.Seq
            d_121_v48_ = _dafny.SeqWithoutIsStrInference([d_120_v47_])
            rhs7_ = (d_120_v47_) not in (d_121_v48_)
            rhs8_ = (d_119_v46_.f11) - ((0) - (d_119_v46_.f11))
            rhs9_ = (((d_117_v44_).set(d_62_v1_, d_115_v42_)).set(d_62_v1_, d_115_v42_)) | ((d_117_v44_).set(not(d_62_v1_), d_115_v42_))
            d_114_cf4_ = rhs7_
            d_87_v26_ = rhs8_
            d_117_v44_ = rhs9_
            index8_ = default__.safeIndex(286, (d_60_v0_).length(0))
            (d_60_v0_)[index8_] = True
            index9_ = default__.safeIndex(286, (d_60_v0_).length(0))
            (d_60_v0_)[index9_] = not (d_62_v1_) or (d_114_cf4_)
        d_62_v1_ = d_62_v1_
        d_87_v26_ = default__.safeDivisionInt(d_87_v26_, d_87_v26_)
        d_122_v49_: _dafny.MultiSet
        d_122_v49_ = _dafny.MultiSet([d_62_v1_, d_62_v1_, d_62_v1_, d_62_v1_, d_62_v1_])
        d_123_v50_: T1
        nw20_ = C1()
        nw20_.ctor__(d_122_v49_, d_87_v26_)
        d_123_v50_ = nw20_
        d_124_v51_: _dafny.Seq
        d_124_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmpkjd"))
        d_125_v52_: _dafny.Set
        d_125_v52_ = _dafny.Set({d_87_v26_})
        rhs10_ = d_62_v1_
        rhs11_ = d_62_v1_
        rhs12_ = d_123_v50_
        rhs13_ = default__.fm30(d_124_v51_, d_87_v26_, (d_87_v26_) + (len(d_125_v52_)), d_61_globalState_)
        rhs14_ = (d_62_v1_ if d_62_v1_ else d_62_v1_)
        d_62_v1_ = rhs10_
        d_62_v1_ = rhs11_
        d_123_v50_ = rhs12_
        d_92_v31_ = rhs13_
        d_62_v1_ = rhs14_
        d_126_v53_: D1
        d_126_v53_ = D1_DC4((0) - (d_87_v26_), d_92_v31_, d_62_v1_, default__.fm24(d_87_v26_, d_62_v1_, d_62_v1_, d_87_v26_, d_61_globalState_))
        d_127_v54_: _dafny.Array
        nw21_ = _dafny.Array(None, 11)
        nw21_[int(0)] = d_126_v53_
        nw21_[int(1)] = d_126_v53_
        nw21_[int(2)] = d_126_v53_
        nw21_[int(3)] = d_126_v53_
        nw21_[int(4)] = d_126_v53_
        nw21_[int(5)] = d_126_v53_
        nw21_[int(6)] = D1_DC4(d_87_v26_, d_92_v31_, d_62_v1_, d_87_v26_)
        nw21_[int(7)] = default__.fm36(d_61_globalState_)
        nw21_[int(8)] = d_126_v53_
        nw21_[int(9)] = d_126_v53_
        nw21_[int(10)] = default__.fm36(d_61_globalState_)
        d_127_v54_ = nw21_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_127_v54_).length(0)):
            d_128_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_128_i0_)) and ((d_128_i0_) < ((d_127_v54_).length(0)))):
                (d_127_v54_)[(d_128_i0_)] = d_126_v53_
        d_129_v55_: _dafny.Array
        nw22_ = _dafny.Array(_dafny.Array(None, 0), 13)
        d_129_v55_ = nw22_
        d_130_v56_: bool
        d_131_v57_: _dafny.Map
        out4_: bool
        out5_: _dafny.Map
        out4_, out5_ = (d_123_v50_).m0(d_129_v55_, d_61_globalState_)
        d_130_v56_ = out4_
        d_131_v57_ = out5_
        d_130_v56_ = d_130_v56_
        hi0_ = d_87_v26_
        for d_132_i1_ in range((d_87_v26_) - (d_87_v26_), hi0_):
            d_133_v58_: _dafny.Map
            d_133_v58_ = _dafny.Map({d_92_v31_: d_130_v56_})
            d_62_v1_ = (((d_133_v58_).set(d_92_v31_, False)) != (d_133_v58_) if d_62_v1_ else (d_89_v28_) != (_dafny.SeqWithoutIsStrInference([d_132_i1_])))
            d_134_v59_: _dafny.Array
            nw23_ = _dafny.Array(int(0), 24)
            d_134_v59_ = nw23_
            d_135_v60_: _dafny.Array
            nw24_ = _dafny.Array(None, 12)
            nw24_[int(0)] = d_134_v59_
            nw24_[int(1)] = d_134_v59_
            nw24_[int(2)] = d_134_v59_
            nw24_[int(3)] = d_134_v59_
            nw24_[int(4)] = d_134_v59_
            nw24_[int(5)] = d_134_v59_
            nw24_[int(6)] = d_134_v59_
            nw24_[int(7)] = d_134_v59_
            nw24_[int(8)] = d_134_v59_
            nw24_[int(9)] = d_134_v59_
            nw24_[int(10)] = d_134_v59_
            nw24_[int(11)] = d_134_v59_
            d_135_v60_ = nw24_
            d_136_v61_: _dafny.Map
            d_136_v61_ = _dafny.Map({not(d_62_v1_): d_132_i1_})
            d_137_v62_: _dafny.Map
            d_137_v62_ = d_136_v61_
            d_134_v59_ = (D14_DC30(762, d_135_v60_, d_137_v62_, -318, d_134_v59_)).cf51
            d_138_v63_: C3
            nw25_ = C3()
            nw25_.ctor__()
            d_138_v63_ = nw25_
            d_139_v64_: D11
            d_139_v64_ = D11_DC24(d_138_v63_)
            d_139_v64_ = d_139_v64_
            d_140_v65_: bool
            d_141_v66_: _dafny.Map
            out6_: bool
            out7_: _dafny.Map
            out6_, out7_ = (d_138_v63_).m0(d_129_v55_, d_61_globalState_)
            d_140_v65_ = out6_
            d_141_v66_ = out7_
        d_142_v67_: _dafny.MultiSet
        d_142_v67_ = _dafny.MultiSet([D10_DC22()])
        d_143_i2_: int
        d_143_i2_ = 0
        with _dafny.label("0"):
            while (not(False) if (d_142_v67_).issubset(d_142_v67_) else d_130_v56_):
                with _dafny.c_label("0"):
                    if (d_143_i2_) >= (100):
                        raise _dafny.Break("0")
                    d_143_i2_ = (d_143_i2_) + (1)
                    d_144_v68_: C2
                    nw26_ = C2()
                    nw26_.ctor__((d_124_v51_) + (d_124_v51_), d_124_v51_)
                    d_144_v68_ = nw26_
                    d_130_v56_ = not(not (d_62_v1_) or (d_62_v1_))
                    d_145_v69_: _dafny.Array
                    nw27_ = _dafny.Array(_dafny.Set({}), 2)
                    d_145_v69_ = nw27_
                    d_145_v69_ = ((d_145_v69_ if d_130_v56_ else d_145_v69_) if (default__.fm20(d_130_v56_, (d_144_v68_).f9, d_62_v1_, d_61_globalState_)).issubset(d_125_v52_) else d_145_v69_)
                    d_87_v26_ = default__.safeDivisionInt(d_87_v26_, 15)
                    pass
            pass
        d_146_v70_: _dafny.Array
        nw28_ = _dafny.Array(int(0), 21)
        d_146_v70_ = nw28_
        index10_ = default__.safeIndex(377, (d_146_v70_).length(0))
        (d_146_v70_)[index10_] = 72
        index11_ = default__.safeIndex(377, (d_146_v70_).length(0))
        (d_146_v70_)[index11_] = default__.safeModuloInt(-592, d_87_v26_)
        d_147_v71_: bool
        d_148_v72_: _dafny.Map
        out8_: bool
        out9_: _dafny.Map
        out8_, out9_ = (d_123_v50_).m0(d_129_v55_, d_61_globalState_)
        d_147_v71_ = out8_
        d_148_v72_ = out9_
        d_62_v1_ = not(d_62_v1_)
        d_87_v26_ = len(((d_124_v51_).set(default__.safeIndex((d_146_v70_)[default__.safeIndex(377, (d_146_v70_).length(0))], len(d_124_v51_)), d_92_v31_) if d_130_v56_ else d_124_v51_))
        d_124_v51_ = d_124_v51_
        _dafny.print(_dafny.string_of((d_60_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_61_globalState_).f0)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_61_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_62_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_87_v26_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_88_v27_) == (_dafny.Map({542: 542}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v28_) == (_dafny.SeqWithoutIsStrInference([542]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_v29_) == (_dafny.SeqWithoutIsStrInference([542]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_v30_) == (_dafny.Map({542: _dafny.SeqWithoutIsStrInference([542])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_92_v31_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v49_) == (_dafny.MultiSet([False, False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_124_v51_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_125_v52_) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v53_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v53_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v53_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v53_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[0]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[0]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[0]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[0]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[1]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[1]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[1]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[1]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[2]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[2]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[2]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[2]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[3]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[3]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[3]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[3]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[4]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[4]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[4]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[4]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[5]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[5]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[5]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[5]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[6]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[6]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[6]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[6]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[7]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[7]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[7]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[7]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[8]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[8]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[8]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[8]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[9]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[9]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[9]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[9]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[10]).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[10]).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[10]).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_127_v54_)[10]).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_130_v56_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v57_) == (_dafny.Map({1: _dafny.MultiSet([False, False, False, False, False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_142_v67_) == (_dafny.MultiSet([D10_DC22()]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_v70_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_147_v71_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_v72_) == (_dafny.Map({1: _dafny.MultiSet([False, False, False, False, False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {self.cf1.VerbatimString(True)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf3 == __o.cf3
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

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
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
        return lambda: D2_DC6(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC6(D2, NamedTuple('DC6', [('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf10 == __o.cf10
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

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(False)
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

class D4_DC10(D4, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf15 == __o.cf15
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
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

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
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC13(D6, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC15(False, int(0), int(0), int(0), None)
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

class D7_DC15(D7, NamedTuple('DC15', [('cf21', Any), ('cf22', Any), ('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC14(D7, NamedTuple('DC14', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf20 == __o.cf20
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
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)

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
        return lambda: D9_DC19(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D9_DC18)

class D9_DC19(D9, NamedTuple('DC19', [('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC20(D9, NamedTuple('DC20', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC18(D9, NamedTuple('DC18', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC18({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC18) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC22()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)

class D10_DC22(D10, NamedTuple('DC22', [])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22)
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC23(D10, NamedTuple('DC23', [('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC21(D10, NamedTuple('DC21', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC25(False, D9.default()())
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

class D11_DC25(D11, NamedTuple('DC25', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC26(D11, NamedTuple('DC26', [('cf42', Any), ('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC24(D11, NamedTuple('DC24', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D12_DC27)

class D12_DC27(D12, NamedTuple('DC27', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC27({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC27) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D13_DC28)

class D13_DC28(D13, NamedTuple('DC28', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC28({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC28) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC30(int(0), _dafny.Array(None, 0), _dafny.Map({}), int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D14_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D14_DC29)

class D14_DC30(D14, NamedTuple('DC30', [('cf47', Any), ('cf48', Any), ('cf49', Any), ('cf50', Any), ('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC30({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC30) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC29(D14, NamedTuple('DC29', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC29({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC29) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC32(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D15_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D15_DC31)

class D15_DC32(D15, NamedTuple('DC32', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC32({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC32) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC33(D15, NamedTuple('DC33', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({self.cf54.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC34(D15, NamedTuple('DC34', [('cf55', Any), ('cf56', Any), ('cf57', Any), ('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34({self.cf55.VerbatimString(True)}, {_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34) and self.cf55 == __o.cf55 and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC31(D15, NamedTuple('DC31', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC31({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC31) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D16_DC35)

class D16_DC35(D16, NamedTuple('DC35', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m0(self, p0, globalState):
        pass


class T1:
    pass

class GlobalState:
    def  __init__(self):
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1):
        (self)._f0 = f0
        (self)._f1 = f1

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass


class C1(T1, T0):
    def  __init__(self):
        self.f11: int = int(0)
        self._f10: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f10, f11):
        (self)._f10 = f10
        (self).f11 = f11

    def fm0(self, p0, p1, p2, p3, globalState):
        return (self).f10

    def fm21(self, p0, p1, p2, globalState):
        return not(not(((_dafny.Set({not(True)})).issubset(_dafny.Set({False, True})) if (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "talxrgws")))])) != (_dafny.SeqWithoutIsStrInference([self.f11])) else (self.f11) == (self.f11))))

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_149_v0_: bool
        d_149_v0_ = True
        d_150_v1_: D1
        d_150_v1_ = D1_DC2(not(d_149_v0_))
        d_151_v2_: _dafny.Seq
        d_151_v2_ = _dafny.SeqWithoutIsStrInference([False])
        d_152_v3_: D4
        d_152_v3_ = D4_DC10(not(False))
        d_153_v4_: str
        d_153_v4_ = _dafny.CodePoint('i')
        d_154_v5_: _dafny.Array
        nw29_ = _dafny.Array(None, 8)
        nw29_[int(0)] = d_149_v0_
        nw29_[int(1)] = (436) == (self.f11)
        nw29_[int(2)] = ((d_150_v1_).cf4) or (d_149_v0_)
        nw29_[int(3)] = (len((d_151_v2_).set(default__.safeIndex(self.f11, len(d_151_v2_)), d_149_v0_))) > (self.f11)
        nw29_[int(4)] = not((d_152_v3_).cf16)
        nw29_[int(5)] = d_149_v0_
        nw29_[int(6)] = not (d_149_v0_) or (d_149_v0_)
        nw29_[int(7)] = (self).fm21(_dafny.Map({d_153_v4_: 179}), self.f11, self.f11, globalState)
        d_154_v5_ = nw29_
        index12_ = default__.safeIndex(111, (d_154_v5_).length(0))
        (d_154_v5_)[index12_] = d_149_v0_
        d_155_v6_: _dafny.Map
        d_155_v6_ = _dafny.Map({d_153_v4_: self.f11})
        index13_ = default__.safeIndex(111, (d_154_v5_).length(0))
        (d_154_v5_)[index13_] = (self).fm21((d_155_v6_).set(d_153_v4_, self.f11), (0) - (default__.safeDivisionInt(567, self.f11)), self.f11, globalState)
        d_156_v7_: _dafny.Array
        nw30_ = _dafny.Array(None, 2)
        nw30_[int(0)] = self.f11
        nw30_[int(1)] = (0) - (self.f11)
        d_156_v7_ = nw30_
        index14_ = default__.safeIndex(449, (d_156_v7_).length(0))
        (d_156_v7_)[index14_] = len(default__.fm22((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))], -599, globalState))
        index15_ = default__.safeIndex(449, (d_156_v7_).length(0))
        (d_156_v7_)[index15_] = (0) - (self.f11)
        d_157_v8_: _dafny.Map
        d_157_v8_ = _dafny.Map({(d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))]: (d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]})
        hi1_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([(d_151_v2_)[default__.safeIndex(self.f11, len(d_151_v2_))]])), 441)
        for d_158_i0_ in range(len(d_157_v8_), hi1_):
            index16_ = default__.safeIndex(449, (d_156_v7_).length(0))
            (d_156_v7_)[index16_] = len(default__.fm23(d_158_i0_, globalState))
            d_159_v9_: _dafny.Seq
            d_159_v9_ = _dafny.SeqWithoutIsStrInference([(d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]])
            d_160_v10_: _dafny.Seq
            d_160_v10_ = _dafny.SeqWithoutIsStrInference([self.f11, len(d_159_v9_), (_dafny.MultiSet([not(not(d_149_v0_)), d_149_v0_, (d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))]])).cardinality])
            d_161_v11_: D2
            d_161_v11_ = D2_DC5(d_159_v9_)
            d_161_v11_ = D2_DC5(d_160_v10_)
            d_162_v12_: _dafny.Seq
            d_162_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cltsjvhb"))
            d_163_v13_: _dafny.Array
            nw31_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_163_v13_ = nw31_
            d_164_v14_: D0
            d_164_v14_ = D0_DC0(411, d_162_v12_, d_163_v13_)
            d_165_v15_: _dafny.Map
            d_165_v15_ = _dafny.Map({d_154_v5_: d_164_v14_})
            d_166_v16_: _dafny.Seq
            d_166_v16_ = _dafny.SeqWithoutIsStrInference([d_164_v14_])
            d_167_v17_: _dafny.Set
            d_167_v17_ = _dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([d_151_v2_])))})
            d_168_v18_: D0
            d_168_v18_ = D0_DC1(((d_165_v15_)[d_154_v5_] if (d_154_v5_) in (d_165_v15_) else (d_166_v16_)[default__.safeIndex(len(d_167_v17_), len(d_166_v16_))]))
            source4_ = d_168_v18_
            if source4_.is_DC0:
                d_169___mcc_h0_ = source4_.cf0
                d_170___mcc_h1_ = source4_.cf1
                d_171___mcc_h2_ = source4_.cf2
                d_172_cf2_ = d_171___mcc_h2_
                d_173_cf1_ = d_170___mcc_h1_
                d_174_cf0_ = d_169___mcc_h0_
                d_175_v19_: C0
                nw32_ = C0()
                nw32_.ctor__()
                d_175_v19_ = nw32_
                index17_ = default__.safeIndex(111, (d_154_v5_).length(0))
                (d_154_v5_)[index17_] = not ((d_160_v10_) == (d_159_v9_)) or (False)
                d_176_v20_: C0
                nw33_ = C0()
                nw33_.ctor__()
                d_176_v20_ = nw33_
                d_177_v21_: _dafny.MultiSet
                d_177_v21_ = _dafny.MultiSet([self.f11])
                d_178_v22_: _dafny.Map
                d_178_v22_ = _dafny.Map({(d_149_v0_) == (not((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))])): (self.f11) in (d_177_v21_)})
                d_178_v22_ = (d_178_v22_).set((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))], (d_150_v1_).cf4)
            elif True:
                d_179___mcc_h3_ = source4_.cf3
                d_180_cf3_ = d_179___mcc_h3_
                index18_ = default__.safeIndex(449, (d_156_v7_).length(0))
                index19_ = default__.safeIndex(449, (d_156_v7_).length(0))
                rhs15_ = self.f11
                rhs16_ = self.f11
                lhs2_ = d_156_v7_
                lhs3_ = default__.safeIndex(449, (d_156_v7_).length(0))
                lhs4_ = d_156_v7_
                lhs5_ = default__.safeIndex(449, (d_156_v7_).length(0))
                lhs2_[lhs3_] = rhs15_
                lhs4_[lhs5_] = rhs16_
                index20_ = default__.safeIndex(111, (d_154_v5_).length(0))
                (d_154_v5_)[index20_] = (self).fm21(d_155_v6_, ((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]) + (self.f11), default__.fm24(self.f11, (d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))], d_149_v0_, len(d_159_v9_), globalState), globalState)
                d_153_v4_ = d_153_v4_
                d_153_v4_ = d_153_v4_
            index21_ = default__.safeIndex(111, (d_154_v5_).length(0))
            (d_154_v5_)[index21_] = d_149_v0_
        d_181_v23_: _dafny.Seq
        d_181_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtlum"))
        hi2_ = self.f11
        for d_182_i1_ in range(default__.safeDivisionInt(len(d_181_v23_), self.f11), hi2_):
            d_183_v24_: _dafny.Seq
            d_183_v24_ = _dafny.SeqWithoutIsStrInference([d_182_i1_, 712])
            d_149_v0_ = (((self).f10).set(d_149_v0_, default__.abs(len(d_183_v24_)))).issubset((self).f10)
            index22_ = default__.safeIndex(449, (d_156_v7_).length(0))
            rhs17_ = d_149_v0_
            rhs18_ = (d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]
            rhs19_ = (d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]
            lhs6_ = self
            lhs7_ = d_156_v7_
            lhs8_ = default__.safeIndex(449, (d_156_v7_).length(0))
            r0 = rhs17_
            lhs6_.f11 = rhs18_
            lhs7_[lhs8_] = rhs19_
            d_184_v25_: _dafny.Array
            nw34_ = _dafny.Array(D2.default()(), 23)
            d_184_v25_ = nw34_
            d_185_v26_: D2
            d_185_v26_ = D2_DC6((_dafny.MultiSet(d_183_v24_)).cardinality, (d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))])
            index23_ = default__.safeIndex(57, (d_184_v25_).length(0))
            (d_184_v25_)[index23_] = d_185_v26_
            index24_ = default__.safeIndex(57, (d_184_v25_).length(0))
            (d_184_v25_)[index24_] = d_185_v26_
            d_186_v27_: _dafny.Array
            nw35_ = _dafny.Array(None, 1)
            nw35_[int(0)] = _dafny.SeqWithoutIsStrInference([d_181_v23_ for d_187_i2_ in range(default__.abs(-530))])
            d_186_v27_ = nw35_
            d_188_v28_: _dafny.Seq
            d_188_v28_ = _dafny.SeqWithoutIsStrInference([d_181_v23_])
            index25_ = default__.safeIndex(840, (d_186_v27_).length(0))
            (d_186_v27_)[index25_] = d_188_v28_
            index26_ = default__.safeIndex(840, (d_186_v27_).length(0))
            (d_186_v27_)[index26_] = (d_188_v28_) + (default__.fm25(d_181_v23_, default__.fm24((0) - (d_182_i1_), (d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))], d_149_v0_, (_dafny.MultiSet(d_183_v24_)).cardinality, globalState), d_182_i1_, globalState))
        if True:
            d_181_v23_ = d_181_v23_
            if (d_149_v0_) and (((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))]) and ((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))])):
                d_149_v0_ = d_149_v0_
                d_189_v29_: _dafny.Array
                nw36_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_189_v29_ = nw36_
                d_189_v29_ = p0
                d_190_v30_: D1
                d_190_v30_ = D1_DC3(self.f11)
                d_191_v31_: _dafny.Seq
                d_191_v31_ = _dafny.SeqWithoutIsStrInference([(d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]])
                d_192_v32_: _dafny.Map
                d_192_v32_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(d_190_v30_).cf5, self.f11]) if (d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))] else d_191_v31_): (False) or ((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))])})
                d_192_v32_ = d_192_v32_
                index27_ = default__.safeIndex(449, (d_156_v7_).length(0))
                (d_156_v7_)[index27_] = self.f11
                index28_ = default__.safeIndex(111, (d_154_v5_).length(0))
                (d_154_v5_)[index28_] = (False) or ((d_191_v31_) == (_dafny.SeqWithoutIsStrInference([self.f11])))
            elif True:
                index29_ = default__.safeIndex(449, (d_156_v7_).length(0))
                (d_156_v7_)[index29_] = (d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]
                d_149_v0_ = (d_151_v2_)[default__.safeIndex((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))], len(d_151_v2_))]
                d_150_v1_ = d_150_v1_
                d_193_v33_: _dafny.Seq
                d_193_v33_ = _dafny.SeqWithoutIsStrInference([self.f11, len(d_181_v23_), (d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]])
                d_194_v34_: _dafny.Map
                d_194_v34_ = _dafny.Map({d_149_v0_: ((_dafny.MultiSet([(d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_153_v4_ for d_195_i3_ in range(default__.abs(274))]))])).set(self.f11, default__.abs((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]))) - (_dafny.MultiSet(d_193_v33_))})
                d_196_v36_: _dafny.Map
                d_196_v36_ = _dafny.Map({self.f11: d_153_v4_})
                d_197_v37_: _dafny.MultiSet
                d_197_v37_ = _dafny.MultiSet([len((_dafny.SeqWithoutIsStrInference([d_153_v4_ for d_198_i4_ in range(default__.abs(304))])).set(default__.safeIndex(self.f11, len(_dafny.SeqWithoutIsStrInference([d_153_v4_ for d_199_i4_ in range(default__.abs(304))]))), d_153_v4_))])
                def iife8_():
                    coll8_ = _dafny.Map()
                    compr_8_: str
                    for compr_8_ in (_dafny.Map({((d_196_v36_)[self.f11] if (self.f11) in (d_196_v36_) else d_153_v4_): False})).keys.Elements:
                        d_200_v35_: str = compr_8_
                        if (d_200_v35_) in (_dafny.Map({((d_196_v36_)[self.f11] if (self.f11) in (d_196_v36_) else d_153_v4_): False})):
                            coll8_[d_200_v35_] = self.f11
                    return _dafny.Map(coll8_)
                d_194_v34_ = (d_194_v34_).set(not ((self).fm21(iife8_()
                , (d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))], len(default__.fm26((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))], globalState)), globalState)) or (d_149_v0_), d_197_v37_)
                index30_ = default__.safeIndex(111, (d_154_v5_).length(0))
                (d_154_v5_)[index30_] = (True) or ((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))])
            d_201_v38_: _dafny.Seq
            d_201_v38_ = _dafny.SeqWithoutIsStrInference([self.f11])
            d_202_v39_: D2
            d_202_v39_ = D2_DC5(d_201_v38_)
            d_202_v39_ = d_202_v39_
            d_203_v40_: C0
            nw37_ = C0()
            nw37_.ctor__()
            d_203_v40_ = nw37_
            d_204_v41_: _dafny.Array
            nw38_ = _dafny.Array(_dafny.Set({}), 8)
            d_204_v41_ = nw38_
            d_205_v42_: _dafny.Array
            nw39_ = _dafny.Array(None, 28)
            d_205_v42_ = nw39_
            d_206_v43_: _dafny.Set
            d_206_v43_ = _dafny.Set({d_205_v42_})
            index31_ = default__.safeIndex(610, (d_204_v41_).length(0))
            (d_204_v41_)[index31_] = d_206_v43_
            d_207_v44_: _dafny.Seq
            d_207_v44_ = _dafny.SeqWithoutIsStrInference([d_205_v42_, d_205_v42_, d_205_v42_, d_205_v42_])
            index32_ = default__.safeIndex(610, (d_204_v41_).length(0))
            (d_204_v41_)[index32_] = (d_206_v43_) - (_dafny.Set({(d_207_v44_)[default__.safeIndex((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))], len(d_207_v44_))], d_205_v42_}))
        elif True:
            index33_ = default__.safeIndex(449, (d_156_v7_).length(0))
            (d_156_v7_)[index33_] = (0) - (self.f11)
            index34_ = default__.safeIndex(111, (d_154_v5_).length(0))
            (d_154_v5_)[index34_] = d_149_v0_
            d_149_v0_ = not((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))])
            d_208_v45_: C0
            nw40_ = C0()
            nw40_.ctor__()
            d_208_v45_ = nw40_
            (self).f11 = ((self).f10).cardinality
        if (d_151_v2_)[default__.safeIndex(((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]) - (self.f11), len(d_151_v2_))]:
            d_209_v46_: _dafny.MultiSet
            d_209_v46_ = _dafny.MultiSet([d_153_v4_, d_153_v4_, d_153_v4_])
            d_209_v46_ = d_209_v46_
            d_210_v47_: _dafny.Seq
            d_210_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqbpjl"))])
            d_211_v48_: _dafny.Set
            d_211_v48_ = _dafny.Set({d_153_v4_, d_153_v4_, d_153_v4_, d_153_v4_, d_153_v4_})
            d_212_v49_: _dafny.MultiSet
            d_212_v49_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")), d_181_v23_, d_181_v23_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmkxsg")), d_181_v23_])
            d_213_v50_: _dafny.Seq
            d_213_v50_ = _dafny.SeqWithoutIsStrInference([d_210_v47_, _dafny.SeqWithoutIsStrInference([d_181_v23_ for d_214_i5_ in range(default__.abs(-244))])])
            d_215_v51_: _dafny.Array
            nw41_ = _dafny.Array(None, 28)
            nw41_[int(0)] = d_156_v7_
            nw41_[int(1)] = d_156_v7_
            nw41_[int(2)] = d_156_v7_
            nw41_[int(3)] = d_156_v7_
            nw41_[int(4)] = d_156_v7_
            nw41_[int(5)] = d_156_v7_
            nw41_[int(6)] = d_156_v7_
            nw41_[int(7)] = d_156_v7_
            nw41_[int(8)] = d_156_v7_
            nw41_[int(9)] = d_156_v7_
            nw41_[int(10)] = d_156_v7_
            nw41_[int(11)] = d_156_v7_
            nw41_[int(12)] = d_156_v7_
            nw41_[int(13)] = d_156_v7_
            nw41_[int(14)] = d_156_v7_
            nw41_[int(15)] = d_156_v7_
            nw41_[int(16)] = d_156_v7_
            nw41_[int(17)] = d_156_v7_
            nw41_[int(18)] = d_156_v7_
            nw41_[int(19)] = d_156_v7_
            nw41_[int(20)] = d_156_v7_
            nw41_[int(21)] = d_156_v7_
            nw41_[int(22)] = d_156_v7_
            nw41_[int(23)] = d_156_v7_
            nw41_[int(24)] = d_156_v7_
            nw41_[int(25)] = d_156_v7_
            nw41_[int(26)] = d_156_v7_
            nw41_[int(27)] = d_156_v7_
            d_215_v51_ = nw41_
            d_216_v52_: D0
            d_216_v52_ = D0_DC0(592, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajxyfq")), d_215_v51_)
            d_217_v53_: _dafny.Array
            nw42_ = _dafny.Array(None, 28)
            nw42_[int(0)] = _dafny.MultiSet(d_210_v47_)
            nw42_[int(1)] = (default__.fm27(globalState)).set(d_181_v23_, default__.abs(len(d_211_v48_)))
            nw42_[int(2)] = _dafny.MultiSet((d_210_v47_ if (d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))] else d_210_v47_))
            nw42_[int(3)] = d_212_v49_
            nw42_[int(4)] = d_212_v49_
            nw42_[int(5)] = d_212_v49_
            nw42_[int(6)] = d_212_v49_
            nw42_[int(7)] = d_212_v49_
            nw42_[int(8)] = _dafny.MultiSet([d_181_v23_, (default__.fm22((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))], self.f11, globalState)).set(default__.safeIndex(((self).f10).cardinality, len(default__.fm22((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))], self.f11, globalState))), d_153_v4_), d_181_v23_, d_181_v23_, d_181_v23_])
            nw42_[int(9)] = _dafny.MultiSet((d_213_v50_)[default__.safeIndex((d_216_v52_).cf0, len(d_213_v50_))])
            nw42_[int(10)] = (_dafny.MultiSet([d_181_v23_, d_181_v23_])) - (d_212_v49_)
            nw42_[int(11)] = d_212_v49_
            nw42_[int(12)] = _dafny.MultiSet((d_210_v47_) + (d_210_v47_))
            nw42_[int(13)] = d_212_v49_
            nw42_[int(14)] = (d_212_v49_ if (d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))] else d_212_v49_)
            nw42_[int(15)] = d_212_v49_
            nw42_[int(16)] = d_212_v49_
            nw42_[int(17)] = d_212_v49_
            nw42_[int(18)] = default__.fm27(globalState)
            nw42_[int(19)] = (d_212_v49_) | (d_212_v49_)
            nw42_[int(20)] = d_212_v49_
            nw42_[int(21)] = d_212_v49_
            nw42_[int(22)] = d_212_v49_
            nw42_[int(23)] = d_212_v49_
            nw42_[int(24)] = d_212_v49_
            nw42_[int(25)] = (_dafny.MultiSet(d_210_v47_)) - (_dafny.MultiSet([d_181_v23_]))
            nw42_[int(26)] = _dafny.MultiSet([d_181_v23_])
            nw42_[int(27)] = d_212_v49_
            d_217_v53_ = nw42_
            index35_ = default__.safeIndex(352, (d_217_v53_).length(0))
            (d_217_v53_)[index35_] = d_212_v49_
            index36_ = default__.safeIndex(352, (d_217_v53_).length(0))
            (d_217_v53_)[index36_] = _dafny.MultiSet([d_181_v23_, (d_181_v23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xce"))), d_181_v23_])
            if (d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))]:
                (self).f11 = (0) - ((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))])
                d_149_v0_ = (d_149_v0_ if d_149_v0_ else not ((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))]) or (d_149_v0_))
                index37_ = default__.safeIndex(449, (d_156_v7_).length(0))
                (d_156_v7_)[index37_] = self.f11
                d_149_v0_ = (d_153_v4_) not in (d_181_v23_)
                index38_ = default__.safeIndex(111, (d_154_v5_).length(0))
                (d_154_v5_)[index38_] = d_149_v0_
            elif True:
                d_218_v54_: _dafny.Array
                def lambda1_(d_219_v23_):
                    def lambda2_(d_220_i6_):
                        return d_219_v23_

                    return lambda2_

                init1_ = lambda1_(d_181_v23_)
                nw43_ = _dafny.Array(None, 15)
                for i0_1_ in range(nw43_.length(0)):
                    nw43_[i0_1_] = init1_(i0_1_)
                d_218_v54_ = nw43_
                d_218_v54_ = d_218_v54_
                d_221_v55_: _dafny.Map
                d_221_v55_ = _dafny.Map({self.f11: ((d_209_v46_)[d_153_v4_] if (d_153_v4_) in (d_209_v46_) else (d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))])})
                (self).f11 = (0) - (default__.safeModuloInt(-943, ((d_221_v55_)[self.f11] if (self.f11) in (d_221_v55_) else self.f11)))
                r0 = (self).fm21(_dafny.Map({d_153_v4_: (d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]}), self.f11, self.f11, globalState)
                d_222_v56_: _dafny.Set
                d_222_v56_ = _dafny.Set({False})
                d_223_v57_: D1
                d_223_v57_ = D1_DC4((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))], d_153_v4_, d_149_v0_, self.f11)
                index39_ = default__.safeIndex(449, (d_156_v7_).length(0))
                (d_156_v7_)[index39_] = ((len(d_222_v56_)) - ((d_223_v57_).cf9)) - (self.f11)
                d_210_v47_ = (d_210_v47_) + (d_210_v47_)
            index40_ = default__.safeIndex(111, (d_154_v5_).length(0))
            (d_154_v5_)[index40_] = (1) != (((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]) - (-273))
            d_224_v58_: _dafny.Array
            nw44_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
            d_224_v58_ = nw44_
            index41_ = default__.safeIndex(592, (d_224_v58_).length(0))
            (d_224_v58_)[index41_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))) + (d_181_v23_)
            d_225_v59_: _dafny.Set
            d_225_v59_ = _dafny.Set({(-451) == (self.f11)})
            index42_ = default__.safeIndex(592, (d_224_v58_).length(0))
            rhs20_ = d_181_v23_
            rhs21_ = d_225_v59_
            lhs9_ = d_224_v58_
            lhs10_ = default__.safeIndex(592, (d_224_v58_).length(0))
            lhs9_[lhs10_] = rhs20_
            d_225_v59_ = rhs21_
        elif True:
            d_226_v60_: _dafny.MultiSet
            d_226_v60_ = _dafny.MultiSet([self.f11])
            d_227_v61_: _dafny.Map
            d_227_v61_ = _dafny.Map({(d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]: _dafny.MultiSet([((d_226_v60_)[(d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]] if ((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))]) in (d_226_v60_) else len(d_181_v23_))])})
            d_228_v62_: D3
            d_228_v62_ = D3_DC8((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))])
            d_229_v63_: _dafny.Map
            d_229_v63_ = _dafny.Map({d_149_v0_: (((d_227_v61_)[(0) - ((d_228_v62_).cf14)] if ((0) - ((d_228_v62_).cf14)) in (d_227_v61_) else d_226_v60_)) - (d_226_v60_)})
            d_229_v63_ = d_229_v63_
            d_149_v0_ = (d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))]
            d_230_v64_: _dafny.Array
            def lambda3_(d_231_v23_, d_232_v7_, d_233_v4_):
                def lambda4_(d_234_i7_):
                    return ((d_231_v23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtwjl")))).set(default__.safeIndex((d_232_v7_)[default__.safeIndex(449, (d_232_v7_).length(0))], len((d_231_v23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtwjl"))))), d_233_v4_)

                return lambda4_

            init2_ = lambda3_(d_181_v23_, d_156_v7_, d_153_v4_)
            nw45_ = _dafny.Array(None, 9)
            for i0_2_ in range(nw45_.length(0)):
                nw45_[i0_2_] = init2_(i0_2_)
            d_230_v64_ = nw45_
            index43_ = default__.safeIndex(997, (d_230_v64_).length(0))
            (d_230_v64_)[index43_] = d_181_v23_
            d_235_v65_: _dafny.Map
            d_235_v65_ = _dafny.Map({d_149_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rotcipwa"))})
            index44_ = default__.safeIndex(997, (d_230_v64_).length(0))
            (d_230_v64_)[index44_] = (default__.fm22(d_149_v0_, (0) - (self.f11), globalState)) + (((d_235_v65_)[not((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))])] if (not((d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))])) in (d_235_v65_) else d_181_v23_))
            d_236_v66_: _dafny.Set
            d_236_v66_ = _dafny.Set({(d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))], (d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))]})
            d_237_v67_: _dafny.Map
            d_237_v67_ = _dafny.Map({d_154_v5_: d_149_v0_})
            d_238_v68_: _dafny.Map
            d_238_v68_ = _dafny.Map({len(d_236_v66_): (d_237_v67_) | (d_237_v67_)})
            d_238_v68_ = (d_238_v68_).set((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))], d_237_v67_)
            d_239_v69_: C0
            nw46_ = C0()
            nw46_.ctor__()
            d_239_v69_ = nw46_
        r0 = (d_153_v4_) in (d_181_v23_)
        d_240_v70_: _dafny.Array
        nw47_ = _dafny.Array(None, 17)
        nw47_[int(0)] = d_156_v7_
        nw47_[int(1)] = d_156_v7_
        nw47_[int(2)] = d_156_v7_
        nw47_[int(3)] = d_156_v7_
        nw47_[int(4)] = d_156_v7_
        nw47_[int(5)] = d_156_v7_
        nw47_[int(6)] = d_156_v7_
        nw47_[int(7)] = d_156_v7_
        nw47_[int(8)] = d_156_v7_
        nw47_[int(9)] = d_156_v7_
        nw47_[int(10)] = d_156_v7_
        nw47_[int(11)] = d_156_v7_
        nw47_[int(12)] = d_156_v7_
        nw47_[int(13)] = d_156_v7_
        nw47_[int(14)] = d_156_v7_
        nw47_[int(15)] = d_156_v7_
        nw47_[int(16)] = d_156_v7_
        d_240_v70_ = nw47_
        d_241_v71_: _dafny.Map
        d_241_v71_ = _dafny.Map({(d_154_v5_)[default__.safeIndex(111, (d_154_v5_).length(0))]: D0_DC0((d_156_v7_)[default__.safeIndex(449, (d_156_v7_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrqnahqbq")), d_240_v70_)})
        d_242_v72_: _dafny.Map
        d_242_v72_ = _dafny.Map({default__.safeDivisionInt(len(d_241_v71_), self.f11): (self).f10})
        r1 = d_242_v72_
        return r0, r1

    def m5(self, p0, p1, globalState):
        d_243_v1_: bool
        d_243_v1_ = False
        d_244_v2_: _dafny.Seq
        d_244_v2_ = _dafny.SeqWithoutIsStrInference([d_243_v1_])
        d_245_i0_: int
        d_245_i0_ = 0
        with _dafny.label("1"):
            def iife9_():
                coll9_ = _dafny.Set()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(109, 909):
                    d_251_v0_: int = compr_9_
                    if ((109) <= (d_251_v0_)) and ((d_251_v0_) < (909)):
                        coll9_ = coll9_.union(_dafny.Set([default__.safeModuloInt(d_251_v0_, p0)]))
                return _dafny.Set(coll9_)
            while not((not((p0) not in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: len(iife9_()
)})) for d_252_i1_ in range(default__.abs(767))])))) in (d_244_v2_)):
                with _dafny.c_label("1"):
                    if (d_245_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_245_i0_ = (d_245_i0_) + (1)
                    d_246_v3_: _dafny.Array
                    nw48_ = _dafny.Array(False, 5)
                    d_246_v3_ = nw48_
                    d_247_v4_: _dafny.Seq
                    d_247_v4_ = _dafny.SeqWithoutIsStrInference([d_246_v3_])
                    d_247_v4_ = d_247_v4_
                    (self).f11 = p0
                    d_248_v5_: _dafny.Map
                    d_248_v5_ = _dafny.Map({d_243_v1_: p0})
                    d_249_v6_: _dafny.Map
                    d_249_v6_ = _dafny.Map({d_248_v5_: d_243_v1_})
                    d_249_v6_ = d_249_v6_
                    d_250_v7_: _dafny.Map
                    d_250_v7_ = _dafny.Map({(p1)[default__.safeIndex(p0, len(p1))]: p0})
                    d_243_v1_ = not(((self).fm21(d_250_v7_, p0, p0, globalState) if d_243_v1_ else (self).fm21(d_250_v7_, p0, default__.fm24(p0, d_243_v1_, d_243_v1_, len(d_244_v2_), globalState), globalState)))
                    pass
            pass
        d_253_v8_: _dafny.Set
        d_253_v8_ = _dafny.Set({self.f11})
        d_243_v1_ = (d_253_v8_).issubset((d_253_v8_).intersection(d_253_v8_))
        d_254_v9_: _dafny.Seq
        d_254_v9_ = _dafny.SeqWithoutIsStrInference([self.f11, len(d_244_v2_)])
        d_255_v10_: D2
        d_255_v10_ = D2_DC5(d_254_v9_)
        pat_let_tv0_ = d_243_v1_
        def lambda5_(source5_):
            if source5_.is_DC6:
                d_256___mcc_h0_ = source5_.cf11
                d_257___mcc_h1_ = source5_.cf12
                d_258_cf12_ = d_257___mcc_h1_
                d_259_cf11_ = d_256___mcc_h0_
                return pat_let_tv0_
            elif True:
                d_260___mcc_h2_ = source5_.cf10
                d_261_cf10_ = d_260___mcc_h2_
                return (_dafny.CodePoint('o')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xiwja")))

        d_243_v1_ = lambda5_(d_255_v10_)
        d_262_v11_: D1
        d_262_v11_ = D1_DC2(d_243_v1_)
        pat_let_tv1_ = d_243_v1_
        def iife10_(_pat_let0_0):
            def iife11_(d_263_dt__update__tmp_h0_):
                def iife12_(_pat_let1_0):
                    def iife13_(d_264_dt__update_hcf4_h0_):
                        return D1_DC2(d_264_dt__update_hcf4_h0_)
                    return iife13_(_pat_let1_0)
                return iife12_(pat_let_tv1_)
            return iife11_(_pat_let0_0)
        d_262_v11_ = iife10_(D1_DC2(d_243_v1_))
        d_265_v12_: _dafny.Map
        d_265_v12_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnc")): 791})
        d_266_v13_: _dafny.Map
        d_266_v13_ = _dafny.Map({_dafny.CodePoint('n'): len(d_265_v12_)})
        d_267_v14_: _dafny.Set
        d_267_v14_ = _dafny.Set({d_243_v1_})
        d_268_v15_: _dafny.Map
        d_268_v15_ = _dafny.Map({False: d_243_v1_})
        d_269_v16_: D1
        d_269_v16_ = D1_DC4(-768, _dafny.CodePoint('o'), True, len(p1))
        d_270_v17_: _dafny.Array
        nw49_ = _dafny.Array(None, 22)
        nw49_[int(0)] = (self).fm21(d_266_v13_, (0) - (p0), p0, globalState)
        nw49_[int(1)] = (d_254_v9_) != (_dafny.SeqWithoutIsStrInference([len(d_244_v2_), self.f11]))
        nw49_[int(2)] = (d_267_v14_) != (d_267_v14_)
        nw49_[int(3)] = not(False)
        nw49_[int(4)] = d_243_v1_
        nw49_[int(5)] = d_243_v1_
        nw49_[int(6)] = d_243_v1_
        nw49_[int(7)] = d_243_v1_
        nw49_[int(8)] = True
        nw49_[int(9)] = d_243_v1_
        nw49_[int(10)] = d_243_v1_
        nw49_[int(11)] = d_243_v1_
        nw49_[int(12)] = (default__.fm24(434, (d_244_v2_)[default__.safeIndex(p0, len(d_244_v2_))], d_243_v1_, ((self).f10).cardinality, globalState)) >= (len(d_268_v15_))
        nw49_[int(13)] = d_243_v1_
        nw49_[int(14)] = d_243_v1_
        nw49_[int(15)] = (self).fm21(d_266_v13_, p0, self.f11, globalState)
        nw49_[int(16)] = d_243_v1_
        nw49_[int(17)] = (self).fm21(d_266_v13_, default__.fm24(899, d_243_v1_, False, 345, globalState), len(_dafny.SeqWithoutIsStrInference([d_243_v1_, False, d_243_v1_, d_243_v1_, d_243_v1_])), globalState)
        nw49_[int(18)] = d_243_v1_
        nw49_[int(19)] = d_243_v1_
        nw49_[int(20)] = (d_269_v16_).cf8
        nw49_[int(21)] = d_243_v1_
        d_270_v17_ = nw49_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_270_v17_).length(0)):
            d_271_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_271_i2_)) and ((d_271_i2_) < ((d_270_v17_).length(0)))):
                (d_270_v17_)[(d_271_i2_)] = ((_dafny.MultiSet(d_254_v9_)).intersection((D3_DC7(_dafny.MultiSet([709, p0]))).cf13)).isdisjoint((_dafny.MultiSet([p0, self.f11, len(p1)])) | (_dafny.MultiSet(d_254_v9_)))
        if d_243_v1_:
            pat_let_tv2_ = p0
            def iife14_(_pat_let2_0):
                def iife15_(d_272_dt__update__tmp_h1_):
                    def iife16_(_pat_let3_0):
                        def iife17_(d_273_dt__update_hcf11_h0_):
                            return D2_DC6(d_273_dt__update_hcf11_h0_, (d_272_dt__update__tmp_h1_).cf12)
                        return iife17_(_pat_let3_0)
                    return iife16_(pat_let_tv2_)
                return iife15_(_pat_let2_0)
            if ((d_243_v1_) and ((self).fm21(d_266_v13_, self.f11, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "weqkh"))), globalState))) or ((iife14_(D2_DC6(p0, d_243_v1_))).cf12):
                (self).f11 = self.f11
                d_243_v1_ = not(d_243_v1_)
                (self).f11 = self.f11
                d_274_v18_: C0
                nw50_ = C0()
                nw50_.ctor__()
                d_274_v18_ = nw50_
                d_275_v19_: _dafny.Array
                nw51_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_275_v19_ = nw51_
                d_276_v20_: _dafny.Array
                nw52_ = _dafny.Array(_dafny.Seq({}), 20)
                d_276_v20_ = nw52_
                d_277_v21_: _dafny.Seq
                d_277_v21_ = _dafny.SeqWithoutIsStrInference([d_276_v20_])
                index45_ = default__.safeIndex(984, (d_275_v19_).length(0))
                (d_275_v19_)[index45_] = (d_277_v21_)[default__.safeIndex(self.f11, len(d_277_v21_))]
                index46_ = default__.safeIndex(984, (d_275_v19_).length(0))
                (d_275_v19_)[index46_] = d_276_v20_
            elif True:
                d_255_v10_ = d_255_v10_
                d_262_v11_ = d_262_v11_
                d_278_v22_: _dafny.Array
                nw53_ = _dafny.Array(int(0), 27)
                d_278_v22_ = nw53_
                d_278_v22_ = d_278_v22_
                d_243_v1_ = (default__.safeDivisionInt(self.f11, p0)) <= ((self.f11) - (618))
                d_279_v23_: _dafny.Map
                d_279_v23_ = _dafny.Map({len(d_253_v8_): p0})
                d_280_v24_: str
                d_280_v24_ = _dafny.CodePoint('n')
                (self).f11 = default__.safeDivisionInt(default__.safeModuloInt(len(d_279_v23_), self.f11), len((p1) + ((p1).set(default__.safeIndex(779, len(p1)), d_280_v24_))))
            d_281_v25_: C0
            nw54_ = C0()
            nw54_.ctor__()
            d_281_v25_ = nw54_
            d_282_v26_: _dafny.Array
            def lambda6_(d_283_p0_):
                def lambda7_(d_284_i3_):
                    return (d_284_i3_) + (d_283_p0_)

                return lambda7_

            init3_ = lambda6_(p0)
            nw55_ = _dafny.Array(None, 12)
            for i0_3_ in range(nw55_.length(0)):
                nw55_[i0_3_] = init3_(i0_3_)
            d_282_v26_ = nw55_
            d_285_v27_: _dafny.Seq
            d_285_v27_ = _dafny.SeqWithoutIsStrInference([d_282_v26_])
            (self).f11 = (len(d_285_v27_)) + (self.f11)
            if False:
                (self).f11 = p0
                d_286_v28_: _dafny.Map
                d_286_v28_ = _dafny.Map({p0: p0})
                d_287_v29_: _dafny.Map
                d_287_v29_ = _dafny.Map({d_286_v28_: d_243_v1_})
                d_288_v30_: _dafny.Map
                d_288_v30_ = d_287_v29_
                d_289_v31_: _dafny.Array
                nw56_ = _dafny.Array(None, 6)
                nw56_[int(0)] = d_282_v26_
                nw56_[int(1)] = d_282_v26_
                nw56_[int(2)] = d_282_v26_
                nw56_[int(3)] = d_282_v26_
                nw56_[int(4)] = d_282_v26_
                nw56_[int(5)] = d_282_v26_
                d_289_v31_ = nw56_
                d_290_v32_: D0
                d_290_v32_ = D0_DC0(self.f11, p1, d_289_v31_)
                d_291_v33_: _dafny.Map
                d_291_v33_ = _dafny.Map({len((d_287_v29_)): d_290_v32_})
                d_292_v34_: _dafny.Array
                d_292_v34_ = d_270_v17_
                d_293_v35_: _dafny.Array
                nw57_ = _dafny.Array(None, 28)
                nw57_[int(0)] = d_270_v17_
                nw57_[int(1)] = d_270_v17_
                nw57_[int(2)] = d_270_v17_
                nw57_[int(3)] = (d_292_v34_)
                nw57_[int(4)] = d_270_v17_
                nw57_[int(5)] = d_270_v17_
                nw57_[int(6)] = d_270_v17_
                nw57_[int(7)] = (d_270_v17_ if (d_262_v11_).cf4 else d_270_v17_)
                nw57_[int(8)] = d_270_v17_
                nw57_[int(9)] = d_270_v17_
                nw57_[int(10)] = d_270_v17_
                nw57_[int(11)] = d_270_v17_
                nw57_[int(12)] = d_270_v17_
                nw57_[int(13)] = d_270_v17_
                nw57_[int(14)] = d_270_v17_
                nw57_[int(15)] = d_270_v17_
                nw57_[int(16)] = d_270_v17_
                nw57_[int(17)] = d_270_v17_
                nw57_[int(18)] = d_270_v17_
                nw57_[int(19)] = d_270_v17_
                nw57_[int(20)] = d_270_v17_
                nw57_[int(21)] = d_270_v17_
                nw57_[int(22)] = d_270_v17_
                nw57_[int(23)] = (d_270_v17_ if d_243_v1_ else d_270_v17_)
                nw57_[int(24)] = d_270_v17_
                nw57_[int(25)] = d_270_v17_
                nw57_[int(26)] = d_270_v17_
                nw57_[int(27)] = (d_292_v34_)
                d_293_v35_ = nw57_
                rhs22_ = d_291_v33_
                rhs23_ = d_293_v35_
                rhs24_ = ((0) - (self.f11)) - (len(d_254_v9_))
                lhs11_ = self
                d_291_v33_ = rhs22_
                d_293_v35_ = rhs23_
                lhs11_.f11 = rhs24_
                d_294_v36_: _dafny.Map
                d_294_v36_ = _dafny.Map({(d_244_v2_) + (d_244_v2_): d_244_v2_})
                d_294_v36_ = (d_294_v36_).set(d_244_v2_, d_244_v2_)
                d_295_v37_: _dafny.Map
                d_295_v37_ = _dafny.Map({p0: d_243_v1_})
                d_296_v38_: _dafny.Map
                d_296_v38_ = _dafny.Map({d_270_v17_: _dafny.MultiSet([len(d_295_v37_)])})
                d_297_v39_: _dafny.MultiSet
                d_297_v39_ = _dafny.MultiSet([len(p1)])
                d_298_v40_: str
                d_298_v40_ = _dafny.CodePoint('k')
                d_299_v41_: _dafny.Map
                d_299_v41_ = _dafny.Map({p0: d_298_v40_})
                d_300_v42_: _dafny.Map
                d_300_v42_ = _dafny.Map({(((d_296_v38_)[d_270_v17_] if (d_270_v17_) in (d_296_v38_) else d_297_v39_)) | (_dafny.MultiSet([default__.fm24(self.f11, d_243_v1_, ((d_295_v37_)[len(d_254_v9_)] if (len(d_254_v9_)) in (d_295_v37_) else d_243_v1_), len(d_299_v41_), globalState), self.f11, -835, self.f11])): default__.safeDivisionInt(((d_297_v39_)[p0] if (p0) in (d_297_v39_) else p0), self.f11)})
                d_301_v43_: D2
                d_301_v43_ = D2_DC6(p0, d_243_v1_)
                d_302_v44_: _dafny.Map
                d_302_v44_ = _dafny.Map({d_243_v1_: 727})
                rhs25_ = d_300_v42_
                rhs26_ = (self.f11 if (d_243_v1_) in (d_302_v44_) else len((((d_244_v2_).set(default__.safeIndex(p0, len(d_244_v2_)), d_243_v1_)) + (d_244_v2_)).set(default__.safeIndex(len(d_302_v44_), len(((d_244_v2_).set(default__.safeIndex(p0, len(d_244_v2_)), d_243_v1_)) + (d_244_v2_))), True)))
                rhs27_ = (d_301_v43_ if (d_253_v8_).ispropersubset(d_253_v8_) else D2_DC6(self.f11, d_243_v1_))
                rhs28_ = (d_267_v14_).issubset(_dafny.Set({False}))
                lhs12_ = self
                d_300_v42_ = rhs25_
                lhs12_.f11 = rhs26_
                d_301_v43_ = rhs27_
                d_243_v1_ = rhs28_
                d_303_v45_: _dafny.Seq
                d_303_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqinu"))
                d_303_v45_ = p1
            elif True:
                (self).f11 = default__.safeDivisionInt((0) - ((p0) - (p0)), (0) - (default__.safeModuloInt(p0, (_dafny.MultiSet([d_243_v1_])).cardinality)))
                d_304_v46_: D3
                d_304_v46_ = D3_DC8(175)
                d_305_v47_: _dafny.Map
                d_305_v47_ = _dafny.Map({d_304_v46_: d_243_v1_})
                d_306_v48_: str
                d_306_v48_ = _dafny.CodePoint('x')
                d_305_v47_ = ((_dafny.Map({d_304_v46_: (self).fm21(_dafny.Map({d_306_v48_: self.f11}), self.f11, p0, globalState)}) if d_243_v1_ else d_305_v47_)) | (_dafny.Map({d_304_v46_: True}))
                d_270_v17_ = d_270_v17_
                d_243_v1_ = d_243_v1_
                index47_ = default__.safeIndex(122, (d_270_v17_).length(0))
                (d_270_v17_)[index47_] = d_243_v1_
                index48_ = default__.safeIndex(122, (d_270_v17_).length(0))
                (d_270_v17_)[index48_] = d_243_v1_
            d_243_v1_ = ((0) - (p0)) >= (p0)
        elif True:
            d_307_v49_: _dafny.Map
            d_307_v49_ = _dafny.Map({not(not(d_243_v1_)): p0})
            (self).f11 = ((d_307_v49_)[d_243_v1_] if (d_243_v1_) in (d_307_v49_) else self.f11)
            d_308_v50_: str
            d_308_v50_ = _dafny.CodePoint('y')
            d_309_v51_: _dafny.Map
            d_309_v51_ = _dafny.Map({self.f11: d_308_v50_})
            d_309_v51_ = (d_309_v51_).set(self.f11, (d_308_v50_ if d_243_v1_ else d_308_v50_))
            d_244_v2_ = (_dafny.SeqWithoutIsStrInference([d_243_v1_])) + ((D7_DC14(d_244_v2_)).cf20)
            d_310_v52_: _dafny.Seq
            d_310_v52_ = _dafny.SeqWithoutIsStrInference([d_268_v15_])
            d_311_v53_: D4
            d_311_v53_ = D4_DC10(d_243_v1_)
            d_312_v54_: _dafny.Map
            d_312_v54_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_311_v53_).cf16: d_243_v1_})])})
            d_243_v1_ = ((d_310_v52_) + (d_310_v52_)) != (((d_312_v54_)[p0] if (p0) in (d_312_v54_) else d_310_v52_))
            d_313_v55_: _dafny.MultiSet
            d_313_v55_ = _dafny.MultiSet([default__.fm24(p0, d_243_v1_, d_243_v1_, p0, globalState), p0])
            d_314_v56_: _dafny.Seq
            d_314_v56_ = _dafny.SeqWithoutIsStrInference([d_313_v55_, d_313_v55_])
            (self).f11 = (((d_313_v55_) | ((d_314_v56_)[default__.safeIndex(self.f11, len(d_314_v56_))])).cardinality) - (p0)

    @property
    def f10(self):
        return self._f10

class C2(T0):
    def  __init__(self):
        self._f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f8, f9):
        (self)._f8 = f8
        (self)._f9 = f9

    def fm0(self, p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([True])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), False, False])))) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)])))

    def fm17(self, p0, p1, p2, p3, globalState):
        return ((0) - (((_dafny.MultiSet([False])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D2_DC6(681, True)).cf12])))).cardinality)) >= (len((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))])) + (_dafny.SeqWithoutIsStrInference([-645]))))

    def fm18(self, p0, globalState):
        def iife18_():
            coll10_ = _dafny.Set()
            compr_10_: _dafny.Map
            for compr_10_ in (_dafny.MultiSet([_dafny.Map({216: not(True)})])).Elements:
                d_315_v1_: _dafny.Map = compr_10_
                if (d_315_v1_) in (_dafny.MultiSet([_dafny.Map({216: not(True)})])):
                    coll10_ = coll10_.union(_dafny.Set([d_315_v1_]))
            return _dafny.Set(coll10_)
        def iife19_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(114, 409):
                d_316_v0_: int = compr_11_
                if ((114) <= (d_316_v0_)) and ((d_316_v0_) < (409)):
                    coll11_[default__.safeDivisionInt(d_316_v0_, 736)] = True
            return _dafny.Map(coll11_)
        return not((iife18_()
        ).issubset((_dafny.Set({_dafny.Map({531: True}), _dafny.Map({-595: True})})) - (_dafny.Set({_dafny.Map({len((self).f8): True}), iife19_()
        }))))

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_317_v0_: int
        d_317_v0_ = 758
        d_318_v1_: bool
        d_318_v1_ = True
        d_317_v0_ = (d_317_v0_) - (default__.fm19(d_318_v1_, d_317_v0_, globalState))
        d_319_v2_: _dafny.Array
        nw58_ = _dafny.Array(int(0), 12)
        d_319_v2_ = nw58_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_319_v2_).length(0)):
            d_320_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_320_i0_)) and ((d_320_i0_) < ((d_319_v2_).length(0)))):
                (d_319_v2_)[(d_320_i0_)] = (d_320_i0_) - ((d_317_v0_ if d_318_v1_ else len(_dafny.SeqWithoutIsStrInference([(self).f8 for d_321_i1_ in range(default__.abs(570))]))))
        d_322_v3_: _dafny.Array
        def lambda8_(d_323_v0_):
            def lambda9_(d_324_i2_):
                return _dafny.Set({d_323_v0_})

            return lambda9_

        init4_ = lambda8_(d_317_v0_)
        nw59_ = _dafny.Array(None, 23)
        for i0_4_ in range(nw59_.length(0)):
            nw59_[i0_4_] = init4_(i0_4_)
        d_322_v3_ = nw59_
        d_325_v4_: _dafny.Set
        d_325_v4_ = _dafny.Set({(0) - (d_317_v0_), d_317_v0_, d_317_v0_})
        index49_ = default__.safeIndex(104, (d_322_v3_).length(0))
        (d_322_v3_)[index49_] = (d_325_v4_) | (d_325_v4_)
        index50_ = default__.safeIndex(264, (d_319_v2_).length(0))
        (d_319_v2_)[index50_] = (0) - (d_317_v0_)
        d_326_v5_: str
        d_326_v5_ = _dafny.CodePoint('g')
        d_327_v6_: _dafny.Set
        d_327_v6_ = _dafny.Set({d_326_v5_, d_326_v5_})
        d_328_v7_: _dafny.Map
        d_328_v7_ = _dafny.Map({d_327_v6_: d_319_v2_})
        d_329_v8_: _dafny.Array
        nw60_ = _dafny.Array(None, 24)
        nw60_[int(0)] = (d_319_v2_ if d_318_v1_ else d_319_v2_)
        nw60_[int(1)] = d_319_v2_
        nw60_[int(2)] = d_319_v2_
        nw60_[int(3)] = d_319_v2_
        nw60_[int(4)] = ((d_328_v7_)[d_327_v6_] if (d_327_v6_) in (d_328_v7_) else d_319_v2_)
        nw60_[int(5)] = d_319_v2_
        nw60_[int(6)] = d_319_v2_
        nw60_[int(7)] = d_319_v2_
        nw60_[int(8)] = d_319_v2_
        nw60_[int(9)] = d_319_v2_
        nw60_[int(10)] = d_319_v2_
        nw60_[int(11)] = d_319_v2_
        nw60_[int(12)] = d_319_v2_
        nw60_[int(13)] = d_319_v2_
        nw60_[int(14)] = d_319_v2_
        nw60_[int(15)] = d_319_v2_
        nw60_[int(16)] = d_319_v2_
        nw60_[int(17)] = (d_319_v2_ if d_318_v1_ else d_319_v2_)
        nw60_[int(18)] = d_319_v2_
        nw60_[int(19)] = d_319_v2_
        nw60_[int(20)] = d_319_v2_
        nw60_[int(21)] = d_319_v2_
        nw60_[int(22)] = d_319_v2_
        nw60_[int(23)] = (d_319_v2_ if not(d_318_v1_) else d_319_v2_)
        d_329_v8_ = nw60_
        index51_ = default__.safeIndex(104, (d_322_v3_).length(0))
        index52_ = default__.safeIndex(264, (d_319_v2_).length(0))
        rhs29_ = default__.fm20(True, (self).f9, (d_317_v0_) != (d_317_v0_), globalState)
        rhs30_ = d_317_v0_
        rhs31_ = d_329_v8_
        lhs13_ = d_322_v3_
        lhs14_ = default__.safeIndex(104, (d_322_v3_).length(0))
        lhs15_ = d_319_v2_
        lhs16_ = default__.safeIndex(264, (d_319_v2_).length(0))
        lhs13_[lhs14_] = rhs29_
        lhs15_[lhs16_] = rhs30_
        d_329_v8_ = rhs31_
        d_330_v9_: D1
        d_330_v9_ = D1_DC2(not(d_318_v1_))
        source6_ = d_330_v9_
        if source6_.is_DC3:
            d_331___mcc_h0_ = source6_.cf5
            d_332_cf5_ = d_331___mcc_h0_
            d_333_v10_: _dafny.Array
            nw61_ = _dafny.Array(False, 15)
            d_333_v10_ = nw61_
            d_334_v11_: _dafny.Map
            d_334_v11_ = _dafny.Map({d_333_v10_: d_318_v1_})
            d_335_v12_: D4
            d_335_v12_ = D4_DC10(True)
            index53_ = default__.safeIndex(264, (d_319_v2_).length(0))
            (d_319_v2_)[index53_] = len(((_dafny.Map({d_333_v10_: True})) | (d_334_v11_)) | (_dafny.Map({d_333_v10_: (d_335_v12_).cf16})))
            index54_ = default__.safeIndex(264, (d_319_v2_).length(0))
            (d_319_v2_)[index54_] = (d_332_cf5_) + (908)
            d_336_v13_: _dafny.Seq
            d_336_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndj"))
            d_337_v14_: D0
            d_337_v14_ = D0_DC0((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flum")), d_329_v8_)
            d_338_v15_: _dafny.Seq
            d_338_v15_ = _dafny.SeqWithoutIsStrInference([d_318_v1_, d_318_v1_])
            d_339_v16_: _dafny.MultiSet
            d_339_v16_ = _dafny.MultiSet([d_318_v1_])
            d_340_v17_: _dafny.Seq
            d_340_v17_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_338_v15_)).intersection(d_339_v16_), d_339_v16_, _dafny.MultiSet([d_318_v1_])])
            index55_ = default__.safeIndex(264, (d_319_v2_).length(0))
            rhs32_ = (((d_337_v14_).cf1) + (d_336_v13_)).set(default__.safeIndex(d_317_v0_, len(((d_337_v14_).cf1) + (d_336_v13_))), d_326_v5_)
            rhs33_ = len(d_340_v17_)
            lhs17_ = d_319_v2_
            lhs18_ = default__.safeIndex(264, (d_319_v2_).length(0))
            d_336_v13_ = rhs32_
            lhs17_[lhs18_] = rhs33_
            r0 = d_318_v1_
        elif source6_.is_DC4:
            d_341___mcc_h1_ = source6_.cf6
            d_342___mcc_h2_ = source6_.cf7
            d_343___mcc_h3_ = source6_.cf8
            d_344___mcc_h4_ = source6_.cf9
            d_345_cf9_ = d_344___mcc_h4_
            d_346_cf8_ = d_343___mcc_h3_
            d_347_cf7_ = d_342___mcc_h2_
            d_348_cf6_ = d_341___mcc_h1_
            d_348_cf6_ = d_317_v0_
            d_349_v18_: _dafny.Seq
            d_349_v18_ = _dafny.SeqWithoutIsStrInference([d_346_cf8_])
            d_350_v19_: _dafny.Map
            d_350_v19_ = _dafny.Map({d_347_cf7_: d_318_v1_})
            d_351_v20_: _dafny.MultiSet
            d_351_v20_ = _dafny.MultiSet([121, len(d_350_v19_), d_348_cf6_])
            d_352_v21_: _dafny.Map
            d_352_v21_ = _dafny.Map({d_348_cf6_: not(d_346_cf8_)})
            d_353_v22_: C1
            nw62_ = C1()
            nw62_.ctor__((self).fm0(d_346_cf8_, not((self).fm18(d_349_v18_, globalState)), (d_351_v20_).cardinality, len(d_352_v21_), globalState), d_317_v0_)
            d_353_v22_ = nw62_
            index56_ = default__.safeIndex(264, (d_319_v2_).length(0))
            rhs34_ = (d_349_v18_) < (_dafny.SeqWithoutIsStrInference([d_318_v1_, d_346_cf8_, d_318_v1_]))
            rhs35_ = d_317_v0_
            lhs19_ = d_319_v2_
            lhs20_ = default__.safeIndex(264, (d_319_v2_).length(0))
            d_346_cf8_ = rhs34_
            lhs19_[lhs20_] = rhs35_
            d_354_v23_: _dafny.Array
            nw63_ = _dafny.Array(D2.default()(), 28)
            d_354_v23_ = nw63_
            d_355_v24_: D2
            d_355_v24_ = D2_DC6(d_348_cf6_, d_318_v1_)
            index57_ = default__.safeIndex(68, (d_354_v23_).length(0))
            (d_354_v23_)[index57_] = d_355_v24_
            index58_ = default__.safeIndex(68, (d_354_v23_).length(0))
            (d_354_v23_)[index58_] = d_355_v24_
        elif True:
            d_356___mcc_h5_ = source6_.cf4
            d_357_cf4_ = d_356___mcc_h5_
            if d_318_v1_:
                d_358_v25_: _dafny.Map
                d_358_v25_ = _dafny.Map({d_319_v2_: len(default__.fm28(len(default__.fm29(globalState)), (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], globalState))})
                d_359_v26_: _dafny.Seq
                d_359_v26_ = _dafny.SeqWithoutIsStrInference([d_358_v25_])
                d_358_v25_ = ((d_359_v26_)[default__.safeIndex(d_317_v0_, len(d_359_v26_))]) | (d_358_v25_)
                d_317_v0_ = d_317_v0_
                d_360_v27_: D0
                d_360_v27_ = D0_DC0(338, ((self).f8) + ((self).f9), d_329_v8_)
                d_361_v28_: _dafny.Map
                d_361_v28_ = _dafny.Map({d_326_v5_: d_317_v0_})
                d_360_v27_ = D0_DC0(len((d_361_v28_) | (d_361_v28_)), _dafny.SeqWithoutIsStrInference([d_326_v5_ for d_362_i3_ in range(default__.abs(515))]), d_329_v8_)
                d_363_v29_: _dafny.Seq
                d_363_v29_ = _dafny.SeqWithoutIsStrInference([not(d_318_v1_), d_318_v1_])
                d_364_v30_: _dafny.Seq
                d_364_v30_ = _dafny.SeqWithoutIsStrInference([d_363_v29_])
                d_365_v31_: _dafny.Map
                d_365_v31_ = _dafny.Map({51: d_317_v0_})
                d_366_v32_: _dafny.Seq
                d_366_v32_ = _dafny.SeqWithoutIsStrInference([(d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], d_317_v0_, len(d_365_v31_), (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))]])
                d_367_v33_: _dafny.Map
                d_367_v33_ = _dafny.Map({d_357_cf4_: 345})
                d_368_v34_: _dafny.MultiSet
                d_368_v34_ = _dafny.MultiSet([len(d_367_v33_)])
                d_369_v35_: C0
                nw64_ = C0()
                nw64_.ctor__()
                d_369_v35_ = nw64_
                d_370_v36_: D7
                d_370_v36_ = D7_DC15((_dafny.SeqWithoutIsStrInference([d_363_v29_, _dafny.SeqWithoutIsStrInference([d_357_cf4_, d_318_v1_])])) <= (d_364_v30_), len((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_326_v5_ for d_371_i5_ in range(default__.abs(-711))])) for d_372_i4_ in range(default__.abs(-459))])) + (d_366_v32_)), (d_368_v34_).cardinality, (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], d_369_v35_)
                d_373_v37_: _dafny.MultiSet
                d_373_v37_ = _dafny.MultiSet([True])
                d_374_v38_: T1
                nw65_ = C1()
                nw65_.ctor__(d_373_v37_, (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))])
                d_374_v38_ = nw65_
                d_375_v39_: _dafny.MultiSet
                d_375_v39_ = _dafny.MultiSet([d_374_v38_, d_374_v38_])
                pat_let_tv3_ = d_369_v35_
                index59_ = default__.safeIndex(264, (d_319_v2_).length(0))
                def iife20_(_pat_let4_0):
                    def iife21_(d_376_dt__update__tmp_h0_):
                        def iife22_(_pat_let5_0):
                            def iife23_(d_377_dt__update_hcf25_h0_):
                                return D7_DC15((d_376_dt__update__tmp_h0_).cf21, (d_376_dt__update__tmp_h0_).cf22, (d_376_dt__update__tmp_h0_).cf23, (d_376_dt__update__tmp_h0_).cf24, d_377_dt__update_hcf25_h0_)
                            return iife23_(_pat_let5_0)
                        return iife22_(pat_let_tv3_)
                    return iife21_(_pat_let4_0)
                rhs36_ = len((d_364_v30_)[default__.safeIndex(d_317_v0_, len(d_364_v30_))])
                rhs37_ = (iife20_(d_370_v36_) if not((_dafny.MultiSet([d_374_v38_])).isdisjoint(d_375_v39_)) else d_370_v36_)
                lhs21_ = d_319_v2_
                lhs22_ = default__.safeIndex(264, (d_319_v2_).length(0))
                lhs21_[lhs22_] = rhs36_
                d_370_v36_ = rhs37_
                d_317_v0_ = default__.safeDivisionInt((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], default__.fm19(d_357_cf4_, (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], globalState))
            elif True:
                d_318_v1_ = not(d_357_cf4_)
                d_378_v40_: _dafny.Array
                nw66_ = _dafny.Array(_dafny.Seq({}), 1)
                d_378_v40_ = nw66_
                index60_ = default__.safeIndex(264, (d_319_v2_).length(0))
                index61_ = default__.safeIndex(264, (d_319_v2_).length(0))
                rhs38_ = (default__.safeDivisionInt((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))])) * (((0) - ((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))])) - ((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))]))
                rhs39_ = len((self).f9)
                rhs40_ = (d_378_v40_ if d_357_cf4_ else d_378_v40_)
                rhs41_ = default__.safeDivisionInt(d_317_v0_, default__.fm24((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], (self).fm17(d_326_v5_, d_357_cf4_, (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], default__.fm24(d_317_v0_, d_318_v1_, not(d_357_cf4_), (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], globalState), globalState), not(not(d_357_cf4_)), len((self).f9), globalState))
                lhs23_ = d_319_v2_
                lhs24_ = default__.safeIndex(264, (d_319_v2_).length(0))
                lhs25_ = d_319_v2_
                lhs26_ = default__.safeIndex(264, (d_319_v2_).length(0))
                d_317_v0_ = rhs38_
                lhs23_[lhs24_] = rhs39_
                d_378_v40_ = rhs40_
                lhs25_[lhs26_] = rhs41_
                d_379_v41_: _dafny.Seq
                d_379_v41_ = _dafny.SeqWithoutIsStrInference([819])
                index62_ = default__.safeIndex(222, (d_378_v40_).length(0))
                (d_378_v40_)[index62_] = d_379_v41_
                d_380_v42_: _dafny.Set
                d_380_v42_ = _dafny.Set({d_318_v1_, d_318_v1_, d_318_v1_})
                index63_ = default__.safeIndex(222, (d_378_v40_).length(0))
                (d_378_v40_)[index63_] = ((default__.fm23((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], globalState)).set(default__.safeIndex(len((d_380_v42_) - (d_380_v42_)), len(default__.fm23((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], globalState))), 778)).set(default__.safeIndex((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], len((default__.fm23((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], globalState)).set(default__.safeIndex(len((d_380_v42_) - (d_380_v42_)), len(default__.fm23((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], globalState))), 778))), len(((_dafny.SeqWithoutIsStrInference([len((self).f8)])) + (_dafny.SeqWithoutIsStrInference([d_317_v0_, d_317_v0_, 766]))).set(default__.safeIndex(d_317_v0_, len((_dafny.SeqWithoutIsStrInference([len((self).f8)])) + (_dafny.SeqWithoutIsStrInference([d_317_v0_, d_317_v0_, 766])))), (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))])))
                d_381_v43_: C0
                nw67_ = C0()
                nw67_.ctor__()
                d_381_v43_ = nw67_
                d_382_v44_: _dafny.Map
                d_382_v44_ = _dafny.Map({d_318_v1_: d_318_v1_})
                d_383_v45_: D3
                d_383_v45_ = D3_DC8(len(d_382_v44_))
                d_384_v46_: _dafny.Map
                d_384_v46_ = _dafny.Map({(D7_DC15(d_318_v1_, len((d_322_v3_)[default__.safeIndex(104, (d_322_v3_).length(0))]), (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], d_317_v0_, d_381_v43_)).cf24: (d_383_v45_).cf14})
                index64_ = default__.safeIndex(264, (d_319_v2_).length(0))
                index65_ = default__.safeIndex(222, (d_378_v40_).length(0))
                rhs42_ = (0) - (d_317_v0_)
                rhs43_ = (_dafny.SeqWithoutIsStrInference([(d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))]]) if d_357_cf4_ else (d_378_v40_)[default__.safeIndex(222, (d_378_v40_).length(0))])
                rhs44_ = True
                rhs45_ = (d_318_v1_ if ((0) - (d_317_v0_)) <= (len(d_384_v46_)) else True)
                lhs27_ = d_319_v2_
                lhs28_ = default__.safeIndex(264, (d_319_v2_).length(0))
                lhs29_ = d_378_v40_
                lhs30_ = default__.safeIndex(222, (d_378_v40_).length(0))
                lhs27_[lhs28_] = rhs42_
                lhs29_[lhs30_] = rhs43_
                d_318_v1_ = rhs44_
                d_318_v1_ = rhs45_
                d_357_cf4_ = not(d_318_v1_)
            d_385_v47_: _dafny.Array
            def lambda10_(d_386_cf4_):
                def lambda11_(d_387_i6_):
                    return _dafny.SeqWithoutIsStrInference([d_386_cf4_, d_386_cf4_, d_386_cf4_])

                return lambda11_

            init5_ = lambda10_(d_357_cf4_)
            nw68_ = _dafny.Array(None, 10)
            for i0_5_ in range(nw68_.length(0)):
                nw68_[i0_5_] = init5_(i0_5_)
            d_385_v47_ = nw68_
            d_388_v48_: _dafny.Seq
            d_388_v48_ = _dafny.SeqWithoutIsStrInference([d_318_v1_, d_357_cf4_, d_357_cf4_, False, d_318_v1_])
            index66_ = default__.safeIndex(505, (d_385_v47_).length(0))
            (d_385_v47_)[index66_] = d_388_v48_
            index67_ = default__.safeIndex(505, (d_385_v47_).length(0))
            (d_385_v47_)[index67_] = (default__.fm26((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], globalState)).set(default__.safeIndex(577, len(default__.fm26((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], globalState))), d_318_v1_)
            d_389_v49_: _dafny.Seq
            d_389_v49_ = _dafny.SeqWithoutIsStrInference([(default__.fm24(d_317_v0_, d_357_cf4_, False, len((self).f9), globalState)) - (d_317_v0_)])
            d_317_v0_ = len(d_389_v49_)
            d_390_v50_: _dafny.Seq
            d_390_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsqb"))
            d_390_v50_ = (self).f8
        d_391_v51_: _dafny.Seq
        d_391_v51_ = _dafny.SeqWithoutIsStrInference([d_318_v1_])
        d_392_v53_: _dafny.Map
        def iife24_():
            coll12_ = _dafny.Set()
            compr_12_: int
            for compr_12_ in _dafny.IntegerRange(-873, 158):
                d_393_v52_: int = compr_12_
                if ((-873) <= (d_393_v52_)) and ((d_393_v52_) < (158)):
                    coll12_ = coll12_.union(_dafny.Set([(d_393_v52_) + (len(d_391_v51_))]))
            return _dafny.Set(coll12_)
        d_392_v53_ = _dafny.Map({(d_391_v51_)[default__.safeIndex((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], len(d_391_v51_))]: iife24_()
        })
        d_394_v55_: _dafny.Seq
        def iife25_():
            coll13_ = _dafny.Set()
            compr_13_: int
            for compr_13_ in _dafny.IntegerRange(53, 673):
                d_395_v54_: int = compr_13_
                if ((53) <= (d_395_v54_)) and ((d_395_v54_) < (673)):
                    coll13_ = coll13_.union(_dafny.Set([(d_395_v54_) - (d_317_v0_)]))
            return _dafny.Set(coll13_)
        d_394_v55_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_317_v0_, (d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))]}), ((d_392_v53_)[d_318_v1_] if (d_318_v1_) in (d_392_v53_) else d_325_v4_), (d_322_v3_)[default__.safeIndex(104, (d_322_v3_).length(0))], iife25_()
        ])
        hi3_ = d_317_v0_
        for d_396_i7_ in range(len((d_394_v55_)[default__.safeIndex((d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))], len(d_394_v55_))]), hi3_):
            d_318_v1_ = d_318_v1_
            index68_ = default__.safeIndex(264, (d_319_v2_).length(0))
            (d_319_v2_)[index68_] = (d_317_v0_) - (d_396_i7_)
            d_397_v57_: _dafny.Array
            def lambda12_(d_398_i7_, d_399_v1_, d_400_v2_, d_401_v0_):
                def lambda13_(d_402_i8_):
                    def iife26_():
                        coll14_ = _dafny.Set()
                        compr_14_: _dafny.Map
                        for compr_14_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_399_v1_: (d_400_v2_)[default__.safeIndex(264, (d_400_v2_).length(0))]}), _dafny.Map({d_399_v1_: d_401_v0_}), _dafny.Map({d_399_v1_: len(_dafny.SeqWithoutIsStrInference([(self).f8, (self).f9]))})])).Elements:
                            d_403_v56_: _dafny.Map = compr_14_
                            if (d_403_v56_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_399_v1_: (d_400_v2_)[default__.safeIndex(264, (d_400_v2_).length(0))]}), _dafny.Map({d_399_v1_: d_401_v0_}), _dafny.Map({d_399_v1_: len(_dafny.SeqWithoutIsStrInference([(self).f8, (self).f9]))})])):
                                coll14_ = coll14_.union(_dafny.Set([d_403_v56_]))
                        return _dafny.Set(coll14_)
                    return (iife26_()
                    ).isdisjoint(_dafny.Set({(_dafny.Map({d_399_v1_: d_398_i7_}))}))

                return lambda13_

            init6_ = lambda12_(d_396_i7_, d_318_v1_, d_319_v2_, d_317_v0_)
            nw69_ = _dafny.Array(None, 26)
            for i0_6_ in range(nw69_.length(0)):
                nw69_[i0_6_] = init6_(i0_6_)
            d_397_v57_ = nw69_
            d_397_v57_ = d_397_v57_
            index69_ = default__.safeIndex(264, (d_319_v2_).length(0))
            (d_319_v2_)[index69_] = d_317_v0_
        d_317_v0_ = len((_dafny.SeqWithoutIsStrInference([d_326_v5_, _dafny.CodePoint('j')])) + (((_dafny.SeqWithoutIsStrInference([d_326_v5_ for d_404_i9_ in range(default__.abs(158))])).set(default__.safeIndex(d_317_v0_, len(_dafny.SeqWithoutIsStrInference([d_326_v5_ for d_405_i9_ in range(default__.abs(158))]))), d_326_v5_)) + ((self).f8)))
        r0 = d_318_v1_
        r1 = _dafny.Map({(d_319_v2_)[default__.safeIndex(264, (d_319_v2_).length(0))]: _dafny.MultiSet([d_318_v1_, d_318_v1_, d_318_v1_, d_318_v1_, d_318_v1_])})
        return r0, r1

    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9

class C3(T1, T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self):
        pass
        pass

    def fm0(self, p0, p1, p2, p3, globalState):
        if (not(not(False))) and (not(not(False))):
            return _dafny.MultiSet([not(True)])
        elif True:
            return _dafny.MultiSet([False])

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_406_v0_: int
        d_406_v0_ = 131
        d_407_v1_: _dafny.MultiSet
        d_407_v1_ = _dafny.MultiSet([d_406_v0_, -663, d_406_v0_])
        d_408_v2_: D3
        d_408_v2_ = D3_DC7(d_407_v1_)
        d_408_v2_ = d_408_v2_
        d_409_v3_: bool
        d_409_v3_ = True
        d_410_v4_: D4
        d_410_v4_ = D4_DC10(d_409_v3_)
        d_410_v4_ = d_410_v4_
        if d_409_v3_:
            d_411_v5_: _dafny.Set
            d_411_v5_ = _dafny.Set({d_409_v3_})
            d_412_v6_: _dafny.Seq
            d_412_v6_ = _dafny.SeqWithoutIsStrInference([d_409_v3_, d_409_v3_])
            d_413_v7_: _dafny.Set
            d_413_v7_ = _dafny.Set({d_406_v0_, default__.fm14(default__.fm14(d_406_v0_, len(d_411_v5_), globalState), d_406_v0_, globalState), len(d_412_v6_), (d_406_v0_) * (default__.fm14(d_406_v0_, d_406_v0_, globalState)), (0) - (d_406_v0_)})
            d_413_v7_ = d_413_v7_
            d_414_v8_: _dafny.Array
            nw70_ = _dafny.Array(False, 29)
            d_414_v8_ = nw70_
            d_415_v9_: _dafny.Map
            d_415_v9_ = _dafny.Map({d_414_v8_: d_409_v3_})
            d_415_v9_ = (d_415_v9_).set(d_414_v8_, d_409_v3_)
            d_416_v10_: _dafny.Seq
            d_416_v10_ = _dafny.SeqWithoutIsStrInference([d_412_v6_])
            d_417_v11_: str
            d_417_v11_ = _dafny.CodePoint('q')
            d_418_v12_: _dafny.Array
            def lambda14_(d_419_v0_):
                def lambda15_(d_420_i1_):
                    return default__.safeDivisionInt(d_420_i1_, d_419_v0_)

                return lambda15_

            init7_ = lambda14_(d_406_v0_)
            nw71_ = _dafny.Array(None, 11)
            for i0_7_ in range(nw71_.length(0)):
                nw71_[i0_7_] = init7_(i0_7_)
            d_418_v12_ = nw71_
            d_421_v13_: _dafny.Array
            nw72_ = _dafny.Array(None, 15)
            nw72_[int(0)] = d_418_v12_
            nw72_[int(1)] = d_418_v12_
            nw72_[int(2)] = d_418_v12_
            nw72_[int(3)] = d_418_v12_
            nw72_[int(4)] = d_418_v12_
            nw72_[int(5)] = d_418_v12_
            nw72_[int(6)] = d_418_v12_
            nw72_[int(7)] = d_418_v12_
            nw72_[int(8)] = d_418_v12_
            nw72_[int(9)] = d_418_v12_
            nw72_[int(10)] = d_418_v12_
            nw72_[int(11)] = d_418_v12_
            nw72_[int(12)] = d_418_v12_
            nw72_[int(13)] = d_418_v12_
            nw72_[int(14)] = d_418_v12_
            d_421_v13_ = nw72_
            d_422_v14_: D0
            d_422_v14_ = D0_DC0(d_406_v0_, _dafny.SeqWithoutIsStrInference([d_417_v11_ for d_423_i0_ in range(default__.abs(642))]), d_421_v13_)
            d_424_v15_: _dafny.Map
            d_424_v15_ = _dafny.Map({d_417_v11_: d_422_v14_})
            d_425_v16_: D0
            d_425_v16_ = D0_DC1(((d_424_v15_)[_dafny.CodePoint('h')] if (_dafny.CodePoint('h')) in (d_424_v15_) else d_422_v14_))
            d_426_v17_: _dafny.Map
            d_426_v17_ = _dafny.Map({d_416_v10_: d_425_v16_})
            d_426_v17_ = (d_426_v17_).set(d_416_v10_, d_425_v16_)
            d_406_v0_ = d_406_v0_
            d_427_v18_: _dafny.Array
            nw73_ = _dafny.Array(_dafny.Map({}), 21)
            d_427_v18_ = nw73_
            d_428_v19_: _dafny.MultiSet
            d_428_v19_ = _dafny.MultiSet([d_409_v3_])
            d_429_v20_: _dafny.Map
            d_429_v20_ = _dafny.Map({(0) - ((d_428_v19_).cardinality): d_417_v11_})
            index70_ = default__.safeIndex(343, (d_427_v18_).length(0))
            (d_427_v18_)[index70_] = (d_429_v20_) | (d_429_v20_)
            d_430_v22_: _dafny.Map
            def iife27_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(11, 664):
                    d_431_v21_: int = compr_15_
                    if ((11) <= (d_431_v21_)) and ((d_431_v21_) < (664)):
                        coll15_[(d_431_v21_) + (d_406_v0_)] = _dafny.CodePoint('m')
                return _dafny.Map(coll15_)
            d_430_v22_ = _dafny.Map({d_414_v8_: iife27_()
            })
            index71_ = default__.safeIndex(343, (d_427_v18_).length(0))
            (d_427_v18_)[index71_] = ((d_430_v22_)[(d_414_v8_ if False else d_414_v8_)] if ((d_414_v8_ if False else d_414_v8_)) in (d_430_v22_) else (d_429_v20_) | (_dafny.Map({(0) - (d_406_v0_): d_417_v11_})))
        elif True:
            d_432_v23_: _dafny.Array
            nw74_ = _dafny.Array(False, 17)
            d_432_v23_ = nw74_
            index72_ = default__.safeIndex(460, (d_432_v23_).length(0))
            (d_432_v23_)[index72_] = d_409_v3_
            index73_ = default__.safeIndex(460, (d_432_v23_).length(0))
            (d_432_v23_)[index73_] = not((d_407_v1_).ispropersubset(d_407_v1_))
            d_433_v24_: str
            d_433_v24_ = _dafny.CodePoint('a')
            d_434_v25_: _dafny.Seq
            d_434_v25_ = _dafny.SeqWithoutIsStrInference([d_433_v24_])
            d_435_v26_: _dafny.Map
            d_435_v26_ = _dafny.Map({d_432_v23_: d_434_v25_})
            d_434_v25_ = ((d_435_v26_)[d_432_v23_] if (d_432_v23_) in (d_435_v26_) else d_434_v25_)
            d_406_v0_ = default__.fm14(d_406_v0_, d_406_v0_, globalState)
            d_436_v27_: _dafny.Map
            d_436_v27_ = _dafny.Map({(d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]: (d_434_v25_) + ((d_434_v25_).set(default__.safeIndex(d_406_v0_, len(d_434_v25_)), d_433_v24_))})
            d_437_v28_: _dafny.Seq
            d_437_v28_ = _dafny.SeqWithoutIsStrInference([d_409_v3_])
            d_438_v29_: D1
            d_438_v29_ = D1_DC3(len(d_437_v28_))
            d_439_v30_: _dafny.MultiSet
            d_439_v30_ = _dafny.MultiSet([(d_410_v4_).cf16])
            index74_ = default__.safeIndex(460, (d_432_v23_).length(0))
            rhs46_ = d_436_v27_
            rhs47_ = ((0) - ((d_438_v29_).cf5)) == (431)
            rhs48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "atwdm"))
            rhs49_ = (((d_439_v30_)[True] if (True) in (d_439_v30_) else d_406_v0_)) * (d_406_v0_)
            lhs31_ = d_432_v23_
            lhs32_ = default__.safeIndex(460, (d_432_v23_).length(0))
            d_436_v27_ = rhs46_
            lhs31_[lhs32_] = rhs47_
            d_434_v25_ = rhs48_
            d_406_v0_ = rhs49_
            d_440_v31_: D1
            d_440_v31_ = D1_DC2(True)
            source7_ = d_440_v31_
            if source7_.is_DC3:
                d_441___mcc_h0_ = source7_.cf5
                d_442_cf5_ = d_441___mcc_h0_
                index75_ = default__.safeIndex(460, (d_432_v23_).length(0))
                (d_432_v23_)[index75_] = (d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]
                d_443_v32_: _dafny.Array
                def lambda16_(d_444_i2_):
                    return (d_444_i2_) * (-103)

                init8_ = lambda16_
                nw75_ = _dafny.Array(None, 7)
                for i0_8_ in range(nw75_.length(0)):
                    nw75_[i0_8_] = init8_(i0_8_)
                d_443_v32_ = nw75_
                d_443_v32_ = d_443_v32_
                d_445_v33_: _dafny.Map
                d_445_v33_ = _dafny.Map({650: d_406_v0_})
                d_446_v34_: _dafny.Map
                d_446_v34_ = _dafny.Map({(d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]: False})
                d_447_v35_: D1
                d_447_v35_ = D1_DC4((0) - (d_406_v0_), d_433_v24_, (d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))], d_406_v0_)
                rhs50_ = d_406_v0_
                rhs51_ = 754
                rhs52_ = (((d_445_v33_)[len(d_446_v34_)] if (len(d_446_v34_)) in (d_445_v33_) else d_406_v0_)) - (d_442_cf5_)
                rhs53_ = (d_447_v35_).cf9
                d_442_cf5_ = rhs50_
                d_442_cf5_ = rhs51_
                d_442_cf5_ = rhs52_
                d_406_v0_ = rhs53_
                d_448_v36_: _dafny.Seq
                d_448_v36_ = _dafny.SeqWithoutIsStrInference([d_432_v23_, d_432_v23_])
                d_449_v37_: _dafny.Set
                d_449_v37_ = _dafny.Set({d_406_v0_, d_406_v0_, d_406_v0_, len((d_437_v28_).set(default__.safeIndex(d_406_v0_, len(d_437_v28_)), (d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]))})
                d_450_v38_: _dafny.Set
                d_450_v38_ = _dafny.Set({(d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))], (d_410_v4_).cf16})
                index76_ = default__.safeIndex(460, (d_432_v23_).length(0))
                index77_ = default__.safeIndex(460, (d_432_v23_).length(0))
                rhs54_ = ((423) + (len(d_445_v33_))) not in ((d_449_v37_ if d_409_v3_ else d_449_v37_))
                rhs55_ = (d_406_v0_) + (d_406_v0_)
                rhs56_ = ((d_450_v38_) | (d_450_v38_)).issubset(_dafny.Set({d_409_v3_}))
                rhs57_ = d_448_v36_
                lhs33_ = d_432_v23_
                lhs34_ = default__.safeIndex(460, (d_432_v23_).length(0))
                lhs35_ = d_432_v23_
                lhs36_ = default__.safeIndex(460, (d_432_v23_).length(0))
                lhs33_[lhs34_] = rhs54_
                d_406_v0_ = rhs55_
                lhs35_[lhs36_] = rhs56_
                d_448_v36_ = rhs57_
            elif source7_.is_DC4:
                d_451___mcc_h1_ = source7_.cf6
                d_452___mcc_h2_ = source7_.cf7
                d_453___mcc_h3_ = source7_.cf8
                d_454___mcc_h4_ = source7_.cf9
                d_455_cf9_ = d_454___mcc_h4_
                d_456_cf8_ = d_453___mcc_h3_
                d_457_cf7_ = d_452___mcc_h2_
                d_458_cf6_ = d_451___mcc_h1_
                index78_ = default__.safeIndex(460, (d_432_v23_).length(0))
                (d_432_v23_)[index78_] = not(not(d_456_cf8_))
                d_459_v39_: _dafny.Array
                nw76_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_459_v39_ = nw76_
                d_460_v40_: _dafny.Map
                d_460_v40_ = _dafny.Map({(d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]: d_409_v3_})
                d_461_v41_: _dafny.Array
                nw77_ = _dafny.Array(None, 21)
                nw77_[int(0)] = d_460_v40_
                nw77_[int(1)] = d_460_v40_
                nw77_[int(2)] = d_460_v40_
                nw77_[int(3)] = d_460_v40_
                nw77_[int(4)] = d_460_v40_
                nw77_[int(5)] = d_460_v40_
                nw77_[int(6)] = _dafny.Map({d_409_v3_: (d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]})
                nw77_[int(7)] = d_460_v40_
                nw77_[int(8)] = d_460_v40_
                nw77_[int(9)] = d_460_v40_
                nw77_[int(10)] = _dafny.Map({(d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]: (d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]})
                nw77_[int(11)] = _dafny.Map({not(d_409_v3_): d_409_v3_})
                nw77_[int(12)] = (d_460_v40_).set(d_409_v3_, d_456_cf8_)
                nw77_[int(13)] = d_460_v40_
                nw77_[int(14)] = d_460_v40_
                nw77_[int(15)] = (d_460_v40_).set(d_409_v3_, False)
                nw77_[int(16)] = default__.fm15(d_458_cf6_, d_456_cf8_, d_457_cf7_, d_455_cf9_, globalState)
                nw77_[int(17)] = d_460_v40_
                nw77_[int(18)] = d_460_v40_
                nw77_[int(19)] = d_460_v40_
                nw77_[int(20)] = d_460_v40_
                d_461_v41_ = nw77_
                index79_ = default__.safeIndex(972, (d_459_v39_).length(0))
                (d_459_v39_)[index79_] = d_461_v41_
                d_462_v42_: _dafny.Array
                nw78_ = _dafny.Array(int(0), 9)
                d_462_v42_ = nw78_
                index80_ = default__.safeIndex(398, (d_462_v42_).length(0))
                (d_462_v42_)[index80_] = len(d_434_v25_)
                d_463_v43_: _dafny.Set
                d_463_v43_ = _dafny.Set({(d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))], d_456_cf8_})
                d_464_v44_: _dafny.Seq
                d_464_v44_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_456_cf8_, (d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]}), d_463_v43_])
                d_465_v45_: _dafny.Array
                nw79_ = _dafny.Array(None, 12)
                nw79_[int(0)] = (d_463_v43_).intersection(d_463_v43_)
                nw79_[int(1)] = (d_464_v44_)[default__.safeIndex(len(d_437_v28_), len(d_464_v44_))]
                nw79_[int(2)] = d_463_v43_
                nw79_[int(3)] = (d_463_v43_).intersection(d_463_v43_)
                nw79_[int(4)] = (d_463_v43_) - (d_463_v43_)
                nw79_[int(5)] = d_463_v43_
                nw79_[int(6)] = (default__.fm16(d_410_v4_, globalState)) | (d_463_v43_)
                nw79_[int(7)] = d_463_v43_
                nw79_[int(8)] = d_463_v43_
                nw79_[int(9)] = d_463_v43_
                nw79_[int(10)] = d_463_v43_
                nw79_[int(11)] = (d_463_v43_) - (_dafny.Set({d_409_v3_}))
                d_465_v45_ = nw79_
                index81_ = default__.safeIndex(326, (d_465_v45_).length(0))
                (d_465_v45_)[index81_] = d_463_v43_
                d_466_v46_: _dafny.Map
                d_466_v46_ = _dafny.Map({(0) - (d_455_cf9_): d_461_v41_})
                index82_ = default__.safeIndex(972, (d_459_v39_).length(0))
                index83_ = default__.safeIndex(398, (d_462_v42_).length(0))
                index84_ = default__.safeIndex(326, (d_465_v45_).length(0))
                rhs58_ = ((d_466_v46_)[-336] if (-336) in (d_466_v46_) else d_461_v41_)
                rhs59_ = d_406_v0_
                rhs60_ = default__.fm16(d_410_v4_, globalState)
                lhs37_ = d_459_v39_
                lhs38_ = default__.safeIndex(972, (d_459_v39_).length(0))
                lhs39_ = d_462_v42_
                lhs40_ = default__.safeIndex(398, (d_462_v42_).length(0))
                lhs41_ = d_465_v45_
                lhs42_ = default__.safeIndex(326, (d_465_v45_).length(0))
                lhs37_[lhs38_] = rhs58_
                lhs39_[lhs40_] = rhs59_
                lhs41_[lhs42_] = rhs60_
                d_467_v47_: _dafny.Map
                d_467_v47_ = _dafny.Map({d_457_cf7_: False})
                d_406_v0_ = ((len(d_434_v25_)) * (d_458_cf6_)) * (((d_462_v42_)[default__.safeIndex(398, (d_462_v42_).length(0))] if ((d_467_v47_)[d_457_cf7_] if (d_457_cf7_) in (d_467_v47_) else (d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]) else (d_462_v42_)[default__.safeIndex(398, (d_462_v42_).length(0))]))
                index85_ = default__.safeIndex(398, (d_462_v42_).length(0))
                rhs61_ = (d_462_v42_)[default__.safeIndex(398, (d_462_v42_).length(0))]
                rhs62_ = (default__.safeModuloInt(d_458_cf6_, (d_462_v42_)[default__.safeIndex(398, (d_462_v42_).length(0))]) if False else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppyxfavq"))))
                lhs43_ = d_462_v42_
                lhs44_ = default__.safeIndex(398, (d_462_v42_).length(0))
                d_458_cf6_ = rhs61_
                lhs43_[lhs44_] = rhs62_
            elif True:
                d_468___mcc_h5_ = source7_.cf4
                d_469_cf4_ = d_468___mcc_h5_
                d_470_v48_: C2
                nw80_ = C2()
                nw80_.ctor__((d_434_v25_ if d_409_v3_ else d_434_v25_), d_434_v25_)
                d_470_v48_ = nw80_
                index86_ = default__.safeIndex(460, (d_432_v23_).length(0))
                index87_ = default__.safeIndex(460, (d_432_v23_).length(0))
                rhs63_ = d_409_v3_
                rhs64_ = (d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]
                rhs65_ = ((d_406_v0_) - (d_406_v0_)) != (d_406_v0_)
                rhs66_ = not (d_409_v3_) or (((d_470_v48_).f9) < (d_434_v25_))
                lhs45_ = d_432_v23_
                lhs46_ = default__.safeIndex(460, (d_432_v23_).length(0))
                lhs47_ = d_432_v23_
                lhs48_ = default__.safeIndex(460, (d_432_v23_).length(0))
                d_469_cf4_ = rhs63_
                d_469_cf4_ = rhs64_
                lhs45_[lhs46_] = rhs65_
                lhs47_[lhs48_] = rhs66_
                index88_ = default__.safeIndex(460, (d_432_v23_).length(0))
                (d_432_v23_)[index88_] = (d_432_v23_)[default__.safeIndex(460, (d_432_v23_).length(0))]
                d_471_v49_: _dafny.Array
                nw81_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_471_v49_ = nw81_
                d_472_v50_: _dafny.Array
                nw82_ = _dafny.Array(int(0), 16)
                d_472_v50_ = nw82_
                index89_ = default__.safeIndex(814, (d_471_v49_).length(0))
                (d_471_v49_)[index89_] = d_472_v50_
                d_473_v51_: _dafny.Seq
                d_473_v51_ = _dafny.SeqWithoutIsStrInference([d_472_v50_, d_472_v50_, d_472_v50_])
                index90_ = default__.safeIndex(814, (d_471_v49_).length(0))
                (d_471_v49_)[index90_] = (d_473_v51_)[default__.safeIndex(d_406_v0_, len(d_473_v51_))]
        d_474_v52_: _dafny.Set
        d_474_v52_ = _dafny.Set({d_409_v3_})
        d_475_i3_: int
        d_475_i3_ = 0
        with _dafny.label("2"):
            while (d_474_v52_).ispropersubset(d_474_v52_):
                with _dafny.c_label("2"):
                    if (d_475_i3_) >= (100):
                        raise _dafny.Break("2")
                    d_475_i3_ = (d_475_i3_) + (1)
                    d_476_v53_: _dafny.Map
                    d_476_v53_ = _dafny.Map({d_409_v3_: -225})
                    d_476_v53_ = d_476_v53_
                    d_477_v54_: _dafny.Seq
                    d_477_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pybwqy"))
                    d_478_v55_: _dafny.Map
                    d_478_v55_ = _dafny.Map({d_406_v0_: d_477_v54_})
                    d_479_v56_: T0
                    nw83_ = C2()
                    nw83_.ctor__((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "poqcmavv"))) + (((d_478_v55_)[d_406_v0_] if (d_406_v0_) in (d_478_v55_) else d_477_v54_)), d_477_v54_)
                    d_479_v56_ = nw83_
                    d_480_v57_: _dafny.Map
                    d_480_v57_ = _dafny.Map({d_409_v3_: d_409_v3_})
                    d_479_v56_ = (d_479_v56_ if ((d_480_v57_)[not(d_409_v3_)] if (not(d_409_v3_)) in (d_480_v57_) else False) else d_479_v56_)
                    d_481_v58_: _dafny.Map
                    d_481_v58_ = _dafny.Map({d_409_v3_: d_410_v4_})
                    d_406_v0_ = default__.safeDivisionInt(d_406_v0_, len(d_481_v58_))
                    d_480_v57_ = (d_480_v57_).set(d_409_v3_, d_409_v3_)
                    pass
            pass
        d_482_v59_: _dafny.MultiSet
        d_482_v59_ = _dafny.MultiSet([d_409_v3_, d_409_v3_])
        d_483_v60_: _dafny.Seq
        d_483_v60_ = _dafny.SeqWithoutIsStrInference([d_409_v3_, d_409_v3_])
        d_484_v61_: C1
        nw84_ = C1()
        nw84_.ctor__(d_482_v59_, (len(d_483_v60_)) * (len(default__.fm16(d_410_v4_, globalState))))
        d_484_v61_ = nw84_
        if d_409_v3_:
            d_485_v62_: _dafny.Array
            nw85_ = _dafny.Array(int(0), 14)
            d_485_v62_ = nw85_
            d_486_v63_: _dafny.MultiSet
            d_486_v63_ = _dafny.MultiSet([d_485_v62_, d_485_v62_])
            d_486_v63_ = d_486_v63_
            d_487_v64_: str
            d_487_v64_ = _dafny.CodePoint('v')
            d_488_v65_: _dafny.Seq
            d_488_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ned"))
            rhs67_ = default__.fm30(d_488_v65_, d_484_v61_.f11, d_484_v61_.f11, globalState)
            rhs68_ = d_487_v64_
            rhs69_ = not(d_409_v3_)
            rhs70_ = d_484_v61_.f11
            lhs49_ = d_484_v61_
            d_487_v64_ = rhs67_
            d_487_v64_ = rhs68_
            d_409_v3_ = rhs69_
            lhs49_.f11 = rhs70_
            d_489_v66_: _dafny.Seq
            d_489_v66_ = _dafny.SeqWithoutIsStrInference([d_484_v61_.f11])
            d_489_v66_ = (d_489_v66_) + ((d_489_v66_) + (d_489_v66_))
            (d_484_v61_).f11 = d_484_v61_.f11
            (d_484_v61_).f11 = (d_484_v61_.f11) * (default__.safeModuloInt(d_484_v61_.f11, d_406_v0_))
        elif True:
            d_490_v67_: _dafny.Seq
            d_490_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwgipadpu"))
            (d_484_v61_).f11 = default__.fm14(d_406_v0_, len(d_490_v67_), globalState)
            d_491_v68_: str
            d_491_v68_ = _dafny.CodePoint('a')
            d_492_v69_: _dafny.Map
            d_492_v69_ = _dafny.Map({((d_490_v67_) + (d_490_v67_)).set(default__.safeIndex(d_406_v0_, len((d_490_v67_) + (d_490_v67_))), d_491_v68_): d_409_v3_})
            d_492_v69_ = (d_492_v69_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_493_i4_ in range(default__.abs(962))]), d_409_v3_)
            if (not((d_491_v68_) in (d_490_v67_)) if not(d_409_v3_) else not(not((default__.fm31(len(d_490_v67_), globalState)).isdisjoint(d_407_v1_)))):
                d_494_v70_: C0
                nw86_ = C0()
                nw86_.ctor__()
                d_494_v70_ = nw86_
                d_495_v71_: C1
                nw87_ = C1()
                nw87_.ctor__((d_484_v61_).fm0(d_409_v3_, d_409_v3_, len(d_474_v52_), -609, globalState), d_484_v61_.f11)
                d_495_v71_ = nw87_
                d_496_v72_: _dafny.Map
                d_496_v72_ = _dafny.Map({d_491_v68_: d_484_v61_.f11})
                d_497_v73_: _dafny.Set
                d_497_v73_ = _dafny.Set({d_490_v67_})
                d_498_v74_: D4
                d_498_v74_ = D4_DC10((d_484_v61_).fm21(d_496_v72_, len(d_497_v73_), 324, globalState))
                d_499_v75_: D4
                d_499_v75_ = D4_DC11(d_498_v74_)
                d_499_v75_ = d_499_v75_
                d_500_v76_: _dafny.Array
                nw88_ = _dafny.Array(None, 17)
                d_500_v76_ = nw88_
                d_501_v77_: T0
                nw89_ = C2()
                nw89_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcchob")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")))
                d_501_v77_ = nw89_
                index91_ = default__.safeIndex(908, (d_500_v76_).length(0))
                (d_500_v76_)[index91_] = d_501_v77_
                index92_ = default__.safeIndex(908, (d_500_v76_).length(0))
                (d_500_v76_)[index92_] = d_501_v77_
                d_502_v78_: _dafny.Map
                d_502_v78_ = _dafny.Map({d_484_v61_.f11: d_484_v61_.f11})
                d_503_v79_: _dafny.Map
                d_503_v79_ = _dafny.Map({d_502_v78_: d_474_v52_})
                d_504_v80_: _dafny.Set
                d_504_v80_ = _dafny.Set({d_484_v61_.f11, ((d_484_v61_).f10).cardinality})
                d_505_v81_: D9
                d_505_v81_ = D9_DC18(_dafny.Map({_dafny.Map({len(d_483_v60_): len(d_504_v80_)}): _dafny.Set({d_409_v3_, d_409_v3_})}))
                d_506_v82_: _dafny.Map
                d_506_v82_ = _dafny.Map({d_495_v71_.f11: d_503_v79_})
                d_507_v83_: _dafny.Map
                d_507_v83_ = _dafny.Map({d_409_v3_: d_504_v80_})
                d_508_v84_: _dafny.Map
                d_508_v84_ = _dafny.Map({d_409_v3_: ((d_507_v83_)[True] if (True) in (d_507_v83_) else d_504_v80_)})
                d_509_v86_: _dafny.Map
                d_509_v86_ = _dafny.Map({True: d_474_v52_})
                d_510_v87_: _dafny.Seq
                d_510_v87_ = _dafny.SeqWithoutIsStrInference([d_503_v79_])
                d_511_v89_: _dafny.Set
                d_511_v89_ = _dafny.Set({d_502_v78_})
                pat_let_tv4_ = d_503_v79_
                d_512_v90_: _dafny.Array
                nw90_ = _dafny.Array(None, 25)
                nw90_[int(0)] = d_503_v79_
                nw90_[int(1)] = d_503_v79_
                nw90_[int(2)] = (d_505_v81_).cf28
                nw90_[int(3)] = d_503_v79_
                nw90_[int(4)] = d_503_v79_
                nw90_[int(5)] = (d_503_v79_) | (d_503_v79_)
                nw90_[int(6)] = _dafny.Map({d_502_v78_: d_474_v52_})
                nw90_[int(7)] = (d_503_v79_) | (d_503_v79_)
                nw90_[int(8)] = (d_503_v79_ if d_409_v3_ else d_503_v79_)
                nw90_[int(9)] = (d_503_v79_) | (d_503_v79_)
                nw90_[int(10)] = ((d_506_v82_)[len(((d_507_v83_)[d_409_v3_] if (d_409_v3_) in (d_507_v83_) else _dafny.Set({d_484_v61_.f11, default__.fm19(d_409_v3_, d_484_v61_.f11, globalState)})))] if (len(((d_507_v83_)[d_409_v3_] if (d_409_v3_) in (d_507_v83_) else _dafny.Set({d_484_v61_.f11, default__.fm19(d_409_v3_, d_484_v61_.f11, globalState)})))) in (d_506_v82_) else d_503_v79_)
                nw90_[int(11)] = d_503_v79_
                nw90_[int(12)] = d_503_v79_
                nw90_[int(13)] = d_503_v79_
                nw90_[int(14)] = (default__.fm32(d_409_v3_, globalState)) | (_dafny.Map({_dafny.Map({d_406_v0_: d_484_v61_.f11}): d_474_v52_}))
                nw90_[int(15)] = d_503_v79_
                nw90_[int(16)] = _dafny.Map({d_502_v78_: d_474_v52_})
                nw90_[int(17)] = ((_dafny.Map({d_502_v78_: d_474_v52_})).set((d_502_v78_).set((0) - (d_406_v0_), d_484_v61_.f11), d_474_v52_)) | (d_503_v79_)
                def iife28_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(187, 166):
                        d_513_v85_: int = compr_16_
                        if ((187) <= (d_513_v85_)) and ((d_513_v85_) < (166)):
                            coll16_[default__.safeModuloInt(d_513_v85_, d_484_v61_.f11)] = d_491_v68_
                    return _dafny.Map(coll16_)
                nw90_[int(18)] = _dafny.Map({_dafny.Map({(((d_484_v61_).f10)[d_409_v3_] if (d_409_v3_) in ((d_484_v61_).f10) else len(iife28_()
                )): d_406_v0_}): ((d_509_v86_)[d_409_v3_] if (d_409_v3_) in (d_509_v86_) else d_474_v52_)})
                nw90_[int(19)] = d_503_v79_
                nw90_[int(20)] = (d_503_v79_) | ((d_510_v87_)[default__.safeIndex(d_484_v61_.f11, len(d_510_v87_))])
                def iife29_():
                    coll17_ = _dafny.Map()
                    compr_17_: _dafny.Map
                    for compr_17_ in (d_511_v89_).Elements:
                        d_514_v88_: _dafny.Map = compr_17_
                        if (d_514_v88_) in (d_511_v89_):
                            coll17_[d_514_v88_] = d_474_v52_
                    return _dafny.Map(coll17_)
                nw90_[int(21)] = iife29_()
                
                def iife30_(_pat_let6_0):
                    def iife31_(d_515_dt__update__tmp_h0_):
                        def iife32_(_pat_let7_0):
                            def iife33_(d_516_dt__update_hcf28_h0_):
                                return D9_DC18(d_516_dt__update_hcf28_h0_)
                            return iife33_(_pat_let7_0)
                        return iife32_(pat_let_tv4_)
                    return iife31_(_pat_let6_0)
                nw90_[int(22)] = ((iife30_(D9_DC18(_dafny.Map({d_502_v78_: d_474_v52_})))).cf28) | (default__.fm32(d_409_v3_, globalState))
                nw90_[int(23)] = _dafny.Map({d_502_v78_: default__.fm16(d_410_v4_, globalState)})
                nw90_[int(24)] = _dafny.Map({(d_502_v78_).set(d_406_v0_, 213): d_474_v52_})
                d_512_v90_ = nw90_
                index93_ = default__.safeIndex(354, (d_512_v90_).length(0))
                (d_512_v90_)[index93_] = d_503_v79_
                index94_ = default__.safeIndex(354, (d_512_v90_).length(0))
                rhs71_ = default__.safeDivisionInt(((d_502_v78_)[d_406_v0_] if (d_406_v0_) in (d_502_v78_) else d_484_v61_.f11), default__.safeDivisionInt(d_484_v61_.f11, d_495_v71_.f11))
                rhs72_ = (d_503_v79_) | (default__.fm32(False, globalState))
                lhs50_ = d_484_v61_
                lhs51_ = d_512_v90_
                lhs52_ = default__.safeIndex(354, (d_512_v90_).length(0))
                lhs50_.f11 = rhs71_
                lhs51_[lhs52_] = rhs72_
            elif True:
                d_517_v91_: _dafny.Array
                nw91_ = _dafny.Array(_dafny.CodePoint('D'), 18)
                d_517_v91_ = nw91_
                d_518_v92_: _dafny.Map
                d_518_v92_ = _dafny.Map({d_484_v61_.f11: d_517_v91_})
                d_519_v93_: _dafny.Seq
                d_519_v93_ = _dafny.SeqWithoutIsStrInference([d_517_v91_, d_517_v91_])
                d_520_v94_: _dafny.Map
                d_520_v94_ = _dafny.Map({False: d_484_v61_.f11})
                d_521_v95_: _dafny.Array
                nw92_ = _dafny.Array(None, 11)
                nw92_[int(0)] = d_517_v91_
                nw92_[int(1)] = d_517_v91_
                nw92_[int(2)] = d_517_v91_
                nw92_[int(3)] = d_517_v91_
                nw92_[int(4)] = d_517_v91_
                nw92_[int(5)] = d_517_v91_
                nw92_[int(6)] = d_517_v91_
                nw92_[int(7)] = d_517_v91_
                nw92_[int(8)] = d_517_v91_
                nw92_[int(9)] = ((d_518_v92_)[d_484_v61_.f11] if (d_484_v61_.f11) in (d_518_v92_) else (d_519_v93_)[default__.safeIndex(len(d_520_v94_), len(d_519_v93_))])
                nw92_[int(10)] = d_517_v91_
                d_521_v95_ = nw92_
                index95_ = default__.safeIndex(127, (d_521_v95_).length(0))
                (d_521_v95_)[index95_] = (d_519_v93_)[default__.safeIndex(d_484_v61_.f11, len(d_519_v93_))]
                index96_ = default__.safeIndex(127, (d_521_v95_).length(0))
                (d_521_v95_)[index96_] = (d_519_v93_)[default__.safeIndex(d_484_v61_.f11, len(d_519_v93_))]
                d_522_v96_: _dafny.Map
                d_522_v96_ = _dafny.Map({(d_484_v61_).f10: d_406_v0_})
                d_523_v97_: _dafny.Map
                d_523_v97_ = _dafny.Map({((d_522_v96_).set((d_484_v61_).f10, d_484_v61_.f11)) | (_dafny.Map({d_482_v59_: d_406_v0_})): (0) - (len(d_490_v67_))})
                (d_484_v61_).f11 = ((d_523_v97_)[d_522_v96_] if (d_522_v96_) in (d_523_v97_) else d_484_v61_.f11)
                d_490_v67_ = d_490_v67_
                (d_484_v61_).f11 = (d_484_v61_.f11 if not (d_409_v3_) or (d_409_v3_) else 393)
                d_524_v98_: _dafny.Array
                nw93_ = _dafny.Array(None, 3)
                nw93_[int(0)] = d_409_v3_
                nw93_[int(1)] = d_409_v3_
                nw93_[int(2)] = d_409_v3_
                d_524_v98_ = nw93_
                index97_ = default__.safeIndex(318, (d_524_v98_).length(0))
                (d_524_v98_)[index97_] = d_409_v3_
                index98_ = default__.safeIndex(318, (d_524_v98_).length(0))
                (d_524_v98_)[index98_] = (d_474_v52_).issubset(d_474_v52_)
            d_525_v99_: C0
            nw94_ = C0()
            nw94_.ctor__()
            d_525_v99_ = nw94_
            d_406_v0_ = (d_406_v0_) + (((_dafny.MultiSet([d_409_v3_])) | (((d_484_v61_).f10).set(False, default__.abs(506)))).cardinality)
        r0 = d_409_v3_
        d_526_v100_: D2
        d_526_v100_ = D2_DC6(d_484_v61_.f11, d_409_v3_)
        d_527_v101_: _dafny.Map
        d_527_v101_ = _dafny.Map({d_484_v61_.f11: _dafny.MultiSet([(d_526_v100_).cf12, False])})
        r1 = (d_527_v101_).set(default__.fm14(d_484_v61_.f11, (0) - ((0) - (d_406_v0_)), globalState), (d_484_v61_).f10)
        return r0, r1


class C4(T0, T1):
    def  __init__(self):
        self.f7: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f7):
        (self).f7 = f7

    def fm0(self, p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([False])) | (_dafny.MultiSet([False, False, False, True, True]))) - (_dafny.MultiSet([not(not(not(True)))]))

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_528_v0_: C3
        nw95_ = C3()
        nw95_.ctor__()
        d_528_v0_ = nw95_
        r0 = (471) not in (_dafny.SeqWithoutIsStrInference([self.f7]))
        d_529_v1_: _dafny.Seq
        d_529_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmcwwhb"))
        d_530_v2_: C2
        nw96_ = C2()
        nw96_.ctor__(d_529_v1_, d_529_v1_)
        d_530_v2_ = nw96_
        d_531_v3_: bool
        d_531_v3_ = True
        d_532_v4_: _dafny.Seq
        d_532_v4_ = _dafny.SeqWithoutIsStrInference([True, True, d_531_v3_])
        r0 = (d_531_v3_ if d_531_v3_ else (d_532_v4_)[default__.safeIndex(self.f7, len(d_532_v4_))])
        d_533_v5_: _dafny.Array
        nw97_ = _dafny.Array(int(0), 19)
        d_533_v5_ = nw97_
        d_534_v6_: D4
        d_534_v6_ = D4_DC9(d_533_v5_)
        d_534_v6_ = (d_534_v6_ if d_531_v3_ else d_534_v6_)
        d_535_i0_: int
        d_535_i0_ = 0
        with _dafny.label("3"):
            while True:
                with _dafny.c_label("3"):
                    if (d_535_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_535_i0_ = (d_535_i0_) + (1)
                    d_529_v1_ = (d_530_v2_).f8
                    rhs73_ = (0) - ((self.f7) * (default__.fm14(self.f7, self.f7, globalState)))
                    rhs74_ = 907
                    rhs75_ = (False) or (d_531_v3_)
                    lhs53_ = self
                    lhs54_ = self
                    lhs53_.f7 = rhs73_
                    lhs54_.f7 = rhs74_
                    r0 = rhs75_
                    d_536_v7_: _dafny.Array
                    nw98_ = _dafny.Array(_dafny.Array(None, 0), 27)
                    d_536_v7_ = nw98_
                    d_537_v8_: _dafny.Array
                    nw99_ = _dafny.Array(None, 26)
                    nw99_[int(0)] = d_533_v5_
                    nw99_[int(1)] = d_533_v5_
                    nw99_[int(2)] = d_533_v5_
                    nw99_[int(3)] = d_533_v5_
                    nw99_[int(4)] = d_533_v5_
                    nw99_[int(5)] = d_533_v5_
                    nw99_[int(6)] = d_533_v5_
                    nw99_[int(7)] = d_533_v5_
                    nw99_[int(8)] = d_533_v5_
                    nw99_[int(9)] = d_533_v5_
                    nw99_[int(10)] = d_533_v5_
                    nw99_[int(11)] = d_533_v5_
                    nw99_[int(12)] = d_533_v5_
                    nw99_[int(13)] = d_533_v5_
                    nw99_[int(14)] = d_533_v5_
                    nw99_[int(15)] = d_533_v5_
                    nw99_[int(16)] = d_533_v5_
                    nw99_[int(17)] = d_533_v5_
                    nw99_[int(18)] = d_533_v5_
                    nw99_[int(19)] = d_533_v5_
                    nw99_[int(20)] = d_533_v5_
                    nw99_[int(21)] = d_533_v5_
                    nw99_[int(22)] = d_533_v5_
                    nw99_[int(23)] = d_533_v5_
                    nw99_[int(24)] = d_533_v5_
                    nw99_[int(25)] = d_533_v5_
                    d_537_v8_ = nw99_
                    index99_ = default__.safeIndex(504, (d_536_v7_).length(0))
                    (d_536_v7_)[index99_] = d_537_v8_
                    index100_ = default__.safeIndex(504, (d_536_v7_).length(0))
                    (d_536_v7_)[index100_] = d_537_v8_
                    index101_ = default__.safeIndex(375, (d_533_v5_).length(0))
                    (d_533_v5_)[index101_] = self.f7
                    d_538_v9_: _dafny.MultiSet
                    d_538_v9_ = _dafny.MultiSet([d_531_v3_, d_531_v3_])
                    d_539_v10_: _dafny.MultiSet
                    d_539_v10_ = _dafny.MultiSet([self.f7, ((d_538_v9_)[d_531_v3_] if (d_531_v3_) in (d_538_v9_) else self.f7)])
                    d_540_v11_: _dafny.Seq
                    d_540_v11_ = _dafny.SeqWithoutIsStrInference([d_529_v1_])
                    index102_ = default__.safeIndex(375, (d_533_v5_).length(0))
                    rhs76_ = 613
                    rhs77_ = self.f7
                    rhs78_ = d_531_v3_
                    rhs79_ = default__.fm19((d_539_v10_).isdisjoint(d_539_v10_), len(d_540_v11_), globalState)
                    lhs55_ = self
                    lhs56_ = d_533_v5_
                    lhs57_ = default__.safeIndex(375, (d_533_v5_).length(0))
                    lhs58_ = self
                    lhs55_.f7 = rhs76_
                    lhs56_[lhs57_] = rhs77_
                    d_531_v3_ = rhs78_
                    lhs58_.f7 = rhs79_
                    pass
            pass
        d_541_v12_: str
        d_541_v12_ = _dafny.CodePoint('i')
        r0 = (d_530_v2_).fm17(d_541_v12_, d_531_v3_, self.f7, self.f7, globalState)
        d_542_v13_: _dafny.Map
        d_542_v13_ = _dafny.Map({self.f7: _dafny.MultiSet(d_532_v4_)})
        d_543_v14_: _dafny.MultiSet
        d_543_v14_ = _dafny.MultiSet([d_531_v3_])
        r1 = ((d_542_v13_) | (d_542_v13_)) | ((d_542_v13_) | (_dafny.Map({self.f7: d_543_v14_})))
        return r0, r1

    def m4(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r0 = p0
        hi4_ = self.f7
        for d_544_i0_ in range(p2, hi4_):
            d_545_v0_: C0
            nw100_ = C0()
            nw100_.ctor__()
            d_545_v0_ = nw100_
            d_546_v1_: D7
            d_546_v1_ = D7_DC15(p0, self.f7, d_544_i0_, self.f7, d_545_v0_)
            d_547_v2_: _dafny.MultiSet
            d_547_v2_ = _dafny.MultiSet([p3])
            d_548_v3_: C1
            nw101_ = C1()
            nw101_.ctor__(d_547_v2_, d_544_i0_)
            d_548_v3_ = nw101_
            d_549_v4_: _dafny.Map
            d_549_v4_ = _dafny.Map({d_548_v3_: p3})
            d_550_v5_: _dafny.Map
            d_550_v5_ = _dafny.Map({p0: len(d_549_v4_)})
            d_551_v6_: _dafny.Map
            d_551_v6_ = _dafny.Map({(p0) == (not(p3)): len(_dafny.Map({d_546_v1_: d_550_v5_}))})
            d_550_v5_ = d_551_v6_
            d_552_v7_: _dafny.Array
            nw102_ = _dafny.Array(False, 1)
            d_552_v7_ = nw102_
            index103_ = default__.safeIndex(283, (d_552_v7_).length(0))
            (d_552_v7_)[index103_] = p3
            index104_ = default__.safeIndex(283, (d_552_v7_).length(0))
            (d_552_v7_)[index104_] = ((p0) in ((d_548_v3_).f10) if (p3 if p3 else p3) else (self.f7) <= (-700))
            d_553_v8_: _dafny.Array
            def lambda17_(d_554_i1_):
                return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_555_i2_ in range(default__.abs(766))])

            init9_ = lambda17_
            nw103_ = _dafny.Array(None, 19)
            for i0_9_ in range(nw103_.length(0)):
                nw103_[i0_9_] = init9_(i0_9_)
            d_553_v8_ = nw103_
            d_556_v9_: _dafny.Seq
            d_556_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "muqj"))
            index105_ = default__.safeIndex(912, (d_553_v8_).length(0))
            (d_553_v8_)[index105_] = (d_556_v9_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yifcb"))).set(default__.safeIndex(d_544_i0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yifcb")))), _dafny.CodePoint('f')))
            index106_ = default__.safeIndex(912, (d_553_v8_).length(0))
            (d_553_v8_)[index106_] = (d_556_v9_ if p0 else (d_556_v9_) + (d_556_v9_))
            (self).f7 = (p2) * (p2)
        d_557_v10_: _dafny.Array
        def lambda18_(d_558_i4_):
            return (d_558_i4_) - ((0) - (self.f7))

        init10_ = lambda18_
        nw104_ = _dafny.Array(None, 20)
        for i0_10_ in range(nw104_.length(0)):
            nw104_[i0_10_] = init10_(i0_10_)
        d_557_v10_ = nw104_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_557_v10_).length(0)):
            d_559_i3_: int = guard_loop_3_
            if (True) and (((0) <= (d_559_i3_)) and ((d_559_i3_) < ((d_557_v10_).length(0)))):
                (d_557_v10_)[(d_559_i3_)] = default__.safeDivisionInt(d_559_i3_, -97)
        d_560_v11_: _dafny.Array
        nw105_ = _dafny.Array(None, 4)
        d_560_v11_ = nw105_
        d_561_v12_: _dafny.Seq
        d_561_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjbqnymx"))
        d_562_v13_: T0
        nw106_ = C2()
        nw106_.ctor__(d_561_v12_, d_561_v12_)
        d_562_v13_ = nw106_
        index107_ = default__.safeIndex(669, (d_560_v11_).length(0))
        (d_560_v11_)[index107_] = d_562_v13_
        index108_ = default__.safeIndex(669, (d_560_v11_).length(0))
        (d_560_v11_)[index108_] = d_562_v13_
        d_563_i5_: int
        d_563_i5_ = 0
        with _dafny.label("4"):
            while p0:
                with _dafny.c_label("4"):
                    if (d_563_i5_) >= (100):
                        raise _dafny.Break("4")
                    d_563_i5_ = (d_563_i5_) + (1)
                    d_564_v14_: _dafny.Seq
                    d_564_v14_ = _dafny.SeqWithoutIsStrInference([p3])
                    d_565_v15_: _dafny.Seq
                    d_565_v15_ = _dafny.SeqWithoutIsStrInference([d_564_v14_, d_564_v14_])
                    d_566_v16_: _dafny.Map
                    d_566_v16_ = _dafny.Map({(d_565_v15_) + (d_565_v15_): not(p3)})
                    r0 = ((d_566_v16_)[_dafny.SeqWithoutIsStrInference([d_564_v14_ for d_567_i6_ in range(default__.abs(699))])] if (_dafny.SeqWithoutIsStrInference([d_564_v14_ for d_568_i6_ in range(default__.abs(699))])) in (d_566_v16_) else default__.fm33(d_561_v12_, globalState))
                    d_569_v17_: _dafny.Array
                    nw107_ = _dafny.Array(_dafny.Array(None, 0), 29)
                    d_569_v17_ = nw107_
                    d_570_v18_: bool
                    d_571_v19_: _dafny.Map
                    out10_: bool
                    out11_: _dafny.Map
                    out10_, out11_ = (d_562_v13_).m0(d_569_v17_, globalState)
                    d_570_v18_ = out10_
                    d_571_v19_ = out11_
                    d_572_v20_: str
                    d_572_v20_ = _dafny.CodePoint('v')
                    d_572_v20_ = d_572_v20_
                    index109_ = default__.safeIndex(981, (p1).length(0))
                    (p1)[index109_] = 25
                    d_573_v21_: _dafny.Array
                    nw108_ = _dafny.Array(D4.default()(), 1)
                    d_573_v21_ = nw108_
                    d_574_v22_: D4
                    d_574_v22_ = D4_DC9(d_557_v10_)
                    d_575_v23_: D4
                    d_575_v23_ = D4_DC11(d_574_v22_)
                    index110_ = default__.safeIndex(194, (d_573_v21_).length(0))
                    (d_573_v21_)[index110_] = D4_DC11(d_574_v22_)
                    d_576_v24_: _dafny.Seq
                    d_576_v24_ = _dafny.SeqWithoutIsStrInference([self.f7])
                    d_577_v25_: _dafny.Seq
                    d_577_v25_ = _dafny.SeqWithoutIsStrInference([len(d_576_v24_)])
                    d_578_v26_: _dafny.MultiSet
                    d_578_v26_ = _dafny.MultiSet([self.f7, len(d_576_v24_)])
                    d_579_v27_: D1
                    d_579_v27_ = D1_DC4(len(_dafny.SeqWithoutIsStrInference([p0, p0, d_570_v18_])), d_572_v20_, p0, p2)
                    d_580_v28_: _dafny.Array
                    nw109_ = _dafny.Array(None, 12)
                    nw109_[int(0)] = p3
                    nw109_[int(1)] = p3
                    nw109_[int(2)] = True
                    nw109_[int(3)] = d_570_v18_
                    nw109_[int(4)] = p3
                    nw109_[int(5)] = (_dafny.MultiSet([len((d_564_v14_).set(default__.safeIndex((0) - (p2), len(d_564_v14_)), False))])).ispropersubset(d_578_v26_)
                    nw109_[int(6)] = (d_570_v18_ if (d_579_v27_).cf8 else not(p0))
                    nw109_[int(7)] = ((d_564_v14_)[default__.safeIndex(self.f7, len(d_564_v14_))]) and (p0)
                    nw109_[int(8)] = (d_561_v12_) <= (d_561_v12_)
                    nw109_[int(9)] = True
                    nw109_[int(10)] = d_570_v18_
                    nw109_[int(11)] = not (not(p3)) or (d_570_v18_)
                    d_580_v28_ = nw109_
                    index111_ = default__.safeIndex(435, (d_580_v28_).length(0))
                    (d_580_v28_)[index111_] = (p0 if not(p3) else not(d_570_v18_))
                    d_581_v29_: _dafny.MultiSet
                    d_581_v29_ = _dafny.MultiSet([d_570_v18_])
                    index112_ = default__.safeIndex(981, (p1).length(0))
                    index113_ = default__.safeIndex(194, (d_573_v21_).length(0))
                    index114_ = default__.safeIndex(435, (d_580_v28_).length(0))
                    rhs80_ = ((d_581_v29_)[p3] if (p3) in (d_581_v29_) else p2)
                    rhs81_ = D4_DC11(d_575_v23_)
                    rhs82_ = (d_564_v14_) != (d_564_v14_)
                    lhs59_ = p1
                    lhs60_ = default__.safeIndex(981, (p1).length(0))
                    lhs61_ = d_573_v21_
                    lhs62_ = default__.safeIndex(194, (d_573_v21_).length(0))
                    lhs63_ = d_580_v28_
                    lhs64_ = default__.safeIndex(435, (d_580_v28_).length(0))
                    lhs59_[lhs60_] = rhs80_
                    lhs61_[lhs62_] = rhs81_
                    lhs63_[lhs64_] = rhs82_
                    pass
            pass
        d_582_v30_: _dafny.MultiSet
        d_582_v30_ = _dafny.MultiSet([p3])
        d_583_v31_: _dafny.Seq
        d_583_v31_ = _dafny.SeqWithoutIsStrInference([False])
        d_584_v32_: _dafny.Seq
        d_584_v32_ = _dafny.SeqWithoutIsStrInference([d_582_v30_, d_582_v30_, _dafny.MultiSet(d_583_v31_)])
        d_585_v33_: D2
        d_585_v33_ = D2_DC6(self.f7, p0)
        d_586_v34_: _dafny.Array
        nw110_ = _dafny.Array(None, 7)
        nw110_[int(0)] = p1
        nw110_[int(1)] = d_557_v10_
        nw110_[int(2)] = p1
        nw110_[int(3)] = d_557_v10_
        nw110_[int(4)] = d_557_v10_
        nw110_[int(5)] = d_557_v10_
        nw110_[int(6)] = p1
        d_586_v34_ = nw110_
        d_587_v35_: D0
        d_587_v35_ = D0_DC0((0) - ((d_585_v33_).cf11), d_561_v12_, d_586_v34_)
        d_588_v36_: C1
        nw111_ = C1()
        nw111_.ctor__((D10_DC23((d_584_v32_)[default__.safeIndex(self.f7, len(d_584_v32_))], p2, True, p0)).cf35, (d_587_v35_).cf0)
        d_588_v36_ = nw111_
        d_589_v37_: _dafny.Set
        d_589_v37_ = _dafny.Set({d_588_v36_, d_588_v36_, (d_588_v36_ if p3 else d_588_v36_)})
        d_589_v37_ = d_589_v37_
        r0 = p0
        return r0


class C5(T0):
    def  __init__(self):
        self._f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f6):
        (self)._f6 = f6

    def fm0(self, p0, p1, p2, p3, globalState):
        if not (False) or (False):
            return _dafny.MultiSet([True, False, True, not(not(True)), not(False)])
        elif True:
            return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False, True, True]))

    def fm13(self, p0, p1, p2, globalState):
        if (_dafny.MultiSet([565, len((self).f6), len(_dafny.SeqWithoutIsStrInference([825]))])).ispropersubset(_dafny.MultiSet([156, len(_dafny.Set({True})), -408, len(_dafny.SeqWithoutIsStrInference([True, True, False, False]))])):
            return len(((self).f6) + ((self).f6))
        elif True:
            return default__.safeDivisionInt(-330, len(_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_590_i0_ in range(default__.abs(602))]), (self).f6})))

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_591_v0_: int
        d_591_v0_ = 112
        d_592_v1_: bool
        d_592_v1_ = False
        d_593_v2_: _dafny.MultiSet
        d_593_v2_ = _dafny.MultiSet([d_591_v0_])
        d_594_v3_: _dafny.Seq
        d_594_v3_ = _dafny.SeqWithoutIsStrInference([d_591_v0_])
        d_595_v4_: _dafny.Seq
        d_595_v4_ = _dafny.SeqWithoutIsStrInference([d_592_v1_])
        d_596_v5_: _dafny.Array
        nw112_ = _dafny.Array(None, 20)
        nw112_[int(0)] = d_591_v0_
        nw112_[int(1)] = (d_591_v0_) - (d_591_v0_)
        nw112_[int(2)] = (0) - (default__.safeDivisionInt(d_591_v0_, d_591_v0_))
        nw112_[int(3)] = d_591_v0_
        nw112_[int(4)] = ((d_593_v2_ if d_592_v1_ else _dafny.MultiSet(d_594_v3_))).cardinality
        nw112_[int(5)] = d_591_v0_
        nw112_[int(6)] = d_591_v0_
        nw112_[int(7)] = d_591_v0_
        nw112_[int(8)] = len(d_594_v3_)
        nw112_[int(9)] = 686
        nw112_[int(10)] = d_591_v0_
        nw112_[int(11)] = 967
        nw112_[int(12)] = (d_591_v0_) - (d_591_v0_)
        nw112_[int(13)] = d_591_v0_
        nw112_[int(14)] = d_591_v0_
        nw112_[int(15)] = len((_dafny.SeqWithoutIsStrInference([d_592_v1_, d_592_v1_])) + (d_595_v4_))
        nw112_[int(16)] = default__.safeModuloInt(d_591_v0_, 753)
        nw112_[int(17)] = d_591_v0_
        nw112_[int(18)] = d_591_v0_
        nw112_[int(19)] = (self).fm13(d_592_v1_, d_591_v0_, d_592_v1_, globalState)
        d_596_v5_ = nw112_
        d_597_v6_: _dafny.MultiSet
        d_597_v6_ = _dafny.MultiSet([(D3_DC7(d_593_v2_)).cf13, d_593_v2_, d_593_v2_])
        index115_ = default__.safeIndex(175, (d_596_v5_).length(0))
        (d_596_v5_)[index115_] = ((d_597_v6_)[d_593_v2_] if (d_593_v2_) in (d_597_v6_) else d_591_v0_)
        index116_ = default__.safeIndex(175, (d_596_v5_).length(0))
        (d_596_v5_)[index116_] = d_591_v0_
        d_598_v7_: _dafny.Seq
        d_598_v7_ = _dafny.SeqWithoutIsStrInference([d_594_v3_])
        d_599_v8_: D2
        d_599_v8_ = D2_DC5((d_598_v7_)[default__.safeIndex((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], len(d_598_v7_))])
        d_600_v9_: _dafny.Map
        d_600_v9_ = _dafny.Map({(d_599_v8_).cf10: d_592_v1_})
        d_600_v9_ = (d_600_v9_).set(d_594_v3_, d_592_v1_)
        if d_592_v1_:
            d_601_v10_: _dafny.Seq
            d_601_v10_ = _dafny.SeqWithoutIsStrInference([d_593_v2_])
            index117_ = default__.safeIndex(175, (d_596_v5_).length(0))
            rhs83_ = not(not((((d_601_v10_)[default__.safeIndex(len((self).f6), len(d_601_v10_))]) - (d_593_v2_)).isdisjoint(d_593_v2_)))
            rhs84_ = d_591_v0_
            lhs65_ = d_596_v5_
            lhs66_ = default__.safeIndex(175, (d_596_v5_).length(0))
            r0 = rhs83_
            lhs65_[lhs66_] = rhs84_
            d_602_v11_: D4
            d_602_v11_ = D4_DC9(d_596_v5_)
            d_603_v12_: _dafny.Set
            d_603_v12_ = _dafny.Set({(d_602_v11_).cf15, d_596_v5_, d_596_v5_, d_596_v5_})
            if ((d_594_v3_)[default__.safeIndex(len(d_603_v12_), len(d_594_v3_))]) <= (d_591_v0_):
                d_604_v13_: str
                d_604_v13_ = _dafny.CodePoint('a')
                d_605_v14_: _dafny.Map
                d_605_v14_ = _dafny.Map({d_591_v0_: d_604_v13_})
                d_606_v15_: C2
                nw113_ = C2()
                nw113_.ctor__(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_607_i0_ in range(default__.abs(945))])) + ((self).f6)).set(default__.safeIndex(len(d_595_v4_), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_608_i0_ in range(default__.abs(945))])) + ((self).f6))), d_604_v13_), ((self).f6).set(default__.safeIndex(len((self).f6), len((self).f6)), ((d_605_v14_)[(d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]] if ((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]) in (d_605_v14_) else d_604_v13_)))
                d_606_v15_ = nw113_
                r0 = d_592_v1_
                d_609_v16_: _dafny.Array
                def lambda19_(d_610_v1_):
                    def lambda20_(d_611_i1_):
                        return d_610_v1_

                    return lambda20_

                init11_ = lambda19_(d_592_v1_)
                nw114_ = _dafny.Array(None, 21)
                for i0_11_ in range(nw114_.length(0)):
                    nw114_[i0_11_] = init11_(i0_11_)
                d_609_v16_ = nw114_
                d_609_v16_ = d_609_v16_
                d_612_v17_: _dafny.Seq
                d_612_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrp"))
                d_613_v18_: _dafny.Array
                nw115_ = _dafny.Array(_dafny.Seq({}), 18)
                d_613_v18_ = nw115_
                d_614_v19_: _dafny.Set
                d_614_v19_ = _dafny.Set({d_591_v0_})
                d_615_v20_: _dafny.Map
                d_615_v20_ = _dafny.Map({d_591_v0_: d_592_v1_})
                d_616_v21_: _dafny.Seq
                d_616_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oexlyc")), (d_606_v15_).f8])
                index118_ = default__.safeIndex(320, (d_613_v18_).length(0))
                (d_613_v18_)[index118_] = (_dafny.SeqWithoutIsStrInference([len(d_614_v19_), (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]])) + (_dafny.SeqWithoutIsStrInference([default__.fm19(d_592_v1_, len(d_615_v20_), globalState), 129, len((d_616_v21_)[default__.safeIndex((d_593_v2_).cardinality, len(d_616_v21_))])]))
                index119_ = default__.safeIndex(320, (d_613_v18_).length(0))
                index120_ = default__.safeIndex(175, (d_596_v5_).length(0))
                rhs85_ = d_604_v13_
                rhs86_ = (d_606_v15_).fm18(d_595_v4_, globalState)
                rhs87_ = _dafny.SeqWithoutIsStrInference([(D1_DC4((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], d_604_v13_, d_592_v1_, (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])).cf7 for d_617_i2_ in range(default__.abs(981))])
                rhs88_ = (d_594_v3_) + (d_594_v3_)
                rhs89_ = (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]
                lhs67_ = d_613_v18_
                lhs68_ = default__.safeIndex(320, (d_613_v18_).length(0))
                lhs69_ = d_596_v5_
                lhs70_ = default__.safeIndex(175, (d_596_v5_).length(0))
                d_604_v13_ = rhs85_
                d_592_v1_ = rhs86_
                d_612_v17_ = rhs87_
                lhs67_[lhs68_] = rhs88_
                lhs69_[lhs70_] = rhs89_
                d_618_v22_: C0
                nw116_ = C0()
                nw116_.ctor__()
                d_618_v22_ = nw116_
            elif True:
                d_619_v23_: _dafny.Array
                nw117_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_619_v23_ = nw117_
                index121_ = default__.safeIndex(378, (d_619_v23_).length(0))
                (d_619_v23_)[index121_] = d_596_v5_
                index122_ = default__.safeIndex(378, (d_619_v23_).length(0))
                (d_619_v23_)[index122_] = d_596_v5_
                d_620_v24_: _dafny.Seq
                d_620_v24_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
                d_621_v25_: str
                d_621_v25_ = _dafny.CodePoint('c')
                d_622_v26_: _dafny.MultiSet
                d_622_v26_ = _dafny.MultiSet([d_621_v25_])
                rhs90_ = d_594_v3_
                rhs91_ = ((len(d_620_v24_)) - (((d_622_v26_).set(d_621_v25_, default__.abs((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]))).cardinality)) + ((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])
                d_594_v3_ = rhs90_
                d_591_v0_ = rhs91_
                d_623_v27_: _dafny.Map
                d_623_v27_ = _dafny.Map({True: d_591_v0_})
                d_623_v27_ = (d_623_v27_).set(d_592_v1_, d_591_v0_)
                d_624_v28_: D4
                d_624_v28_ = D4_DC9((d_619_v23_)[default__.safeIndex(378, (d_619_v23_).length(0))])
                d_625_v29_: D4
                d_625_v29_ = D4_DC11(d_624_v28_)
                d_626_v30_: D4
                d_626_v30_ = D4_DC11(d_624_v28_)
                d_627_v31_: _dafny.Seq
                d_627_v31_ = _dafny.SeqWithoutIsStrInference([d_626_v30_, d_626_v30_, d_626_v30_, d_626_v30_, d_626_v30_])
                d_628_v32_: _dafny.Map
                d_628_v32_ = _dafny.Map({True: d_627_v31_})
                d_629_v33_: _dafny.Map
                d_629_v33_ = _dafny.Map({(d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]: len(d_628_v32_)})
                d_630_v36_: _dafny.Array
                nw118_ = _dafny.Array(None, 13)
                nw118_[int(0)] = default__.fm29(globalState)
                nw118_[int(1)] = (d_629_v33_).set((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], d_591_v0_)
                nw118_[int(2)] = d_629_v33_
                def iife34_():
                    coll18_ = _dafny.Map()
                    compr_18_: int
                    for compr_18_ in _dafny.IntegerRange(744, 9):
                        d_631_v34_: int = compr_18_
                        if ((744) <= (d_631_v34_)) and ((d_631_v34_) < (9)):
                            coll18_[default__.safeModuloInt(d_631_v34_, d_591_v0_)] = (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]
                    return _dafny.Map(coll18_)
                nw118_[int(3)] = iife34_()
                
                nw118_[int(4)] = d_629_v33_
                def iife35_():
                    coll19_ = _dafny.Map()
                    compr_19_: int
                    for compr_19_ in _dafny.IntegerRange(446, 923):
                        d_632_v35_: int = compr_19_
                        if ((446) <= (d_632_v35_)) and ((d_632_v35_) < (923)):
                            coll19_[default__.safeModuloInt(d_632_v35_, (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])] = d_591_v0_
                    return _dafny.Map(coll19_)
                nw118_[int(5)] = iife35_()
                
                nw118_[int(6)] = d_629_v33_
                nw118_[int(7)] = d_629_v33_
                nw118_[int(8)] = (_dafny.Map({(d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]: (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]})).set(d_591_v0_, (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])
                nw118_[int(9)] = d_629_v33_
                nw118_[int(10)] = d_629_v33_
                nw118_[int(11)] = d_629_v33_
                nw118_[int(12)] = (d_629_v33_).set(((d_623_v27_)[d_592_v1_] if (d_592_v1_) in (d_623_v27_) else (_dafny.MultiSet([d_592_v1_, d_592_v1_, d_592_v1_])).cardinality), (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])
                d_630_v36_ = nw118_
                d_633_v37_: _dafny.Seq
                d_633_v37_ = _dafny.SeqWithoutIsStrInference([d_630_v36_])
                d_630_v36_ = (d_633_v37_)[default__.safeIndex((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], len(d_633_v37_))]
                d_591_v0_ = (default__.fm19(d_592_v1_, len((self).f6), globalState)) * (152)
            d_602_v11_ = d_602_v11_
            d_634_v38_: D9
            d_634_v38_ = D9_DC20(d_592_v1_, d_592_v1_, d_592_v1_)
            source8_ = d_634_v38_
            if source8_.is_DC19:
                d_635___mcc_h0_ = source8_.cf29
                d_636___mcc_h1_ = source8_.cf30
                d_637_cf30_ = d_636___mcc_h1_
                d_638_cf29_ = d_635___mcc_h0_
                d_639_v39_: C0
                nw119_ = C0()
                nw119_.ctor__()
                d_639_v39_ = nw119_
                d_640_v40_: _dafny.Array
                nw120_ = _dafny.Array(None, 10)
                nw120_[int(0)] = d_639_v39_
                nw120_[int(1)] = d_639_v39_
                nw120_[int(2)] = d_639_v39_
                nw120_[int(3)] = d_639_v39_
                nw120_[int(4)] = d_639_v39_
                nw120_[int(5)] = d_639_v39_
                nw120_[int(6)] = d_639_v39_
                nw120_[int(7)] = d_639_v39_
                nw120_[int(8)] = d_639_v39_
                nw120_[int(9)] = d_639_v39_
                d_640_v40_ = nw120_
                d_641_v41_: _dafny.Seq
                d_641_v41_ = _dafny.SeqWithoutIsStrInference([d_640_v40_, d_640_v40_, d_640_v40_, d_640_v40_])
                d_642_v42_: _dafny.Array
                nw121_ = _dafny.Array(None, 20)
                nw121_[int(0)] = d_640_v40_
                nw121_[int(1)] = d_640_v40_
                nw121_[int(2)] = d_640_v40_
                nw121_[int(3)] = d_640_v40_
                nw121_[int(4)] = d_640_v40_
                nw121_[int(5)] = d_640_v40_
                nw121_[int(6)] = d_640_v40_
                nw121_[int(7)] = (d_640_v40_ if d_592_v1_ else (d_641_v41_)[default__.safeIndex((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], len(d_641_v41_))])
                nw121_[int(8)] = d_640_v40_
                nw121_[int(9)] = d_640_v40_
                nw121_[int(10)] = d_640_v40_
                nw121_[int(11)] = d_640_v40_
                nw121_[int(12)] = d_640_v40_
                nw121_[int(13)] = d_640_v40_
                nw121_[int(14)] = d_640_v40_
                nw121_[int(15)] = d_640_v40_
                nw121_[int(16)] = d_640_v40_
                nw121_[int(17)] = (d_641_v41_)[default__.safeIndex(d_638_cf29_, len(d_641_v41_))]
                nw121_[int(18)] = d_640_v40_
                nw121_[int(19)] = d_640_v40_
                d_642_v42_ = nw121_
                d_643_v43_: str
                d_643_v43_ = _dafny.CodePoint('r')
                d_644_v44_: _dafny.Map
                d_644_v44_ = _dafny.Map({d_591_v0_: d_642_v42_})
                d_645_v45_: _dafny.Map
                d_645_v45_ = _dafny.Map({d_592_v1_: d_642_v42_})
                rhs92_ = (((d_644_v44_)[d_638_cf29_] if (d_638_cf29_) in (d_644_v44_) else d_642_v42_) if d_592_v1_ else (((d_645_v45_)[d_592_v1_] if (d_592_v1_) in (d_645_v45_) else d_642_v42_) if d_592_v1_ else d_642_v42_))
                rhs93_ = default__.fm30(((self).f6) + ((self).f6), len(d_594_v3_), (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], globalState)
                d_642_v42_ = rhs92_
                d_643_v43_ = rhs93_
                d_646_v46_: C2
                nw122_ = C2()
                nw122_.ctor__(((self).f6) + ((self).f6), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjvex")))
                d_646_v46_ = nw122_
                d_646_v46_ = d_646_v46_
                d_647_v47_: _dafny.MultiSet
                d_647_v47_ = _dafny.MultiSet([d_592_v1_, d_592_v1_])
                d_648_v48_: _dafny.Set
                d_648_v48_ = _dafny.Set({((d_647_v47_)[not(d_592_v1_)] if (not(d_592_v1_)) in (d_647_v47_) else d_638_cf29_)})
                d_592_v1_ = (d_648_v48_).isdisjoint(d_648_v48_)
                d_649_v49_: _dafny.Map
                d_649_v49_ = _dafny.Map({((d_646_v46_).f9).set(default__.safeIndex((self).fm13(d_592_v1_, d_638_cf29_, d_592_v1_, globalState), len((d_646_v46_).f9)), _dafny.CodePoint('u')): 618})
                d_649_v49_ = (d_649_v49_).set((self).f6, d_637_cf30_)
            elif source8_.is_DC20:
                d_650___mcc_h2_ = source8_.cf31
                d_651___mcc_h3_ = source8_.cf32
                d_652___mcc_h4_ = source8_.cf33
                d_653_cf33_ = d_652___mcc_h4_
                d_654_cf32_ = d_651___mcc_h3_
                d_655_cf31_ = d_650___mcc_h2_
                index123_ = default__.safeIndex(175, (d_596_v5_).length(0))
                (d_596_v5_)[index123_] = d_591_v0_
                index124_ = default__.safeIndex(175, (d_596_v5_).length(0))
                (d_596_v5_)[index124_] = d_591_v0_
                d_656_v50_: _dafny.MultiSet
                d_656_v50_ = _dafny.MultiSet([d_592_v1_, d_655_cf31_])
                d_657_v51_: T0
                nw123_ = C4()
                nw123_.ctor__(d_591_v0_)
                d_657_v51_ = nw123_
                d_658_v52_: int
                d_659_v53_: str
                out12_: int
                out13_: str
                out12_, out13_ = (self).m3(not(((_dafny.MultiSet([d_653_cf33_, d_655_cf31_])).set(d_654_cf32_, default__.abs(16))).issubset(d_656_v50_)), d_657_v51_, globalState)
                d_658_v52_ = out12_
                d_659_v53_ = out13_
                index125_ = default__.safeIndex(175, (d_596_v5_).length(0))
                (d_596_v5_)[index125_] = d_658_v52_
            elif True:
                d_660___mcc_h5_ = source8_.cf28
                d_661_cf28_ = d_660___mcc_h5_
                d_662_v54_: _dafny.Array
                nw124_ = _dafny.Array(_dafny.CodePoint('D'), 22)
                d_662_v54_ = nw124_
                d_663_v55_: str
                d_663_v55_ = _dafny.CodePoint('x')
                index126_ = default__.safeIndex(69, (d_662_v54_).length(0))
                (d_662_v54_)[index126_] = d_663_v55_
                index127_ = default__.safeIndex(69, (d_662_v54_).length(0))
                (d_662_v54_)[index127_] = d_663_v55_
                d_591_v0_ = default__.fm14(len(_dafny.Set({d_592_v1_, d_592_v1_})), (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], globalState)
                d_664_v56_: _dafny.Set
                d_664_v56_ = _dafny.Set({not((d_663_v55_) in ((self).f6))})
                d_664_v56_ = d_664_v56_
                d_591_v0_ = (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]
            r0 = d_592_v1_
        elif True:
            d_665_v58_: _dafny.Array
            def lambda21_(d_666_v5_, d_667_v1_):
                def lambda22_(d_668_i3_):
                    def iife36_():
                        coll20_ = _dafny.Set()
                        compr_20_: int
                        for compr_20_ in (_dafny.Map({(d_666_v5_)[default__.safeIndex(175, (d_666_v5_).length(0))]: d_667_v1_})).keys.Elements:
                            d_669_v57_: int = compr_20_
                            if (d_669_v57_) in (_dafny.Map({(d_666_v5_)[default__.safeIndex(175, (d_666_v5_).length(0))]: d_667_v1_})):
                                coll20_ = coll20_.union(_dafny.Set([(d_669_v57_) + ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False])])).cardinality)]))
                        return _dafny.Set(coll20_)
                    return iife36_()
                    

                return lambda22_

            init12_ = lambda21_(d_596_v5_, d_592_v1_)
            nw125_ = _dafny.Array(None, 5)
            for i0_12_ in range(nw125_.length(0)):
                nw125_[i0_12_] = init12_(i0_12_)
            d_665_v58_ = nw125_
            d_670_v59_: _dafny.Set
            d_670_v59_ = _dafny.Set({(d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], d_591_v0_})
            index128_ = default__.safeIndex(936, (d_665_v58_).length(0))
            (d_665_v58_)[index128_] = d_670_v59_
            index129_ = default__.safeIndex(936, (d_665_v58_).length(0))
            (d_665_v58_)[index129_] = d_670_v59_
            d_592_v1_ = not(d_592_v1_)
            d_671_v60_: D4
            d_671_v60_ = D4_DC10(d_592_v1_)
            d_672_v61_: _dafny.Array
            nw126_ = _dafny.Array(None, 13)
            nw126_[int(0)] = d_592_v1_
            nw126_[int(1)] = d_592_v1_
            nw126_[int(2)] = d_592_v1_
            nw126_[int(3)] = False
            nw126_[int(4)] = not(d_592_v1_)
            nw126_[int(5)] = d_592_v1_
            nw126_[int(6)] = (d_671_v60_).cf16
            nw126_[int(7)] = not(d_592_v1_)
            nw126_[int(8)] = d_592_v1_
            nw126_[int(9)] = d_592_v1_
            nw126_[int(10)] = d_592_v1_
            nw126_[int(11)] = d_592_v1_
            nw126_[int(12)] = d_592_v1_
            d_672_v61_ = nw126_
            d_673_v62_: _dafny.Array
            d_673_v62_ = d_672_v61_
            d_672_v61_ = (d_673_v62_)
            d_591_v0_ = default__.safeDivisionInt(d_591_v0_, (0) - ((0) - ((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])))
            index130_ = default__.safeIndex(175, (d_596_v5_).length(0))
            (d_596_v5_)[index130_] = (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]
        hi5_ = d_591_v0_
        for d_674_i4_ in range(default__.safeDivisionInt((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgtil"))))), hi5_):
            d_594_v3_ = ((d_594_v3_) + (d_594_v3_)) + (d_594_v3_)
            r0 = d_592_v1_
            index131_ = default__.safeIndex(175, (d_596_v5_).length(0))
            (d_596_v5_)[index131_] = (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]
            index132_ = default__.safeIndex(175, (d_596_v5_).length(0))
            (d_596_v5_)[index132_] = default__.safeDivisionInt(d_674_i4_, (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])
        d_675_v63_: _dafny.Set
        d_675_v63_ = _dafny.Set({(self).f6, (self).f6, (self).f6})
        d_676_i5_: int
        d_676_i5_ = 0
        with _dafny.label("5"):
            while (_dafny.Set({(self).f6})).isdisjoint(d_675_v63_):
                with _dafny.c_label("5"):
                    if (d_676_i5_) >= (100):
                        raise _dafny.Break("5")
                    d_676_i5_ = (d_676_i5_) + (1)
                    d_677_v64_: C3
                    nw127_ = C3()
                    nw127_.ctor__()
                    d_677_v64_ = nw127_
                    d_678_v65_: D11
                    d_678_v65_ = D11_DC24(d_677_v64_)
                    d_679_v66_: _dafny.Seq
                    d_679_v66_ = _dafny.SeqWithoutIsStrInference([d_677_v64_, d_677_v64_, d_677_v64_])
                    d_680_v67_: _dafny.Array
                    nw128_ = _dafny.Array(None, 11)
                    nw128_[int(0)] = d_677_v64_
                    nw128_[int(1)] = d_677_v64_
                    nw128_[int(2)] = d_677_v64_
                    nw128_[int(3)] = d_677_v64_
                    nw128_[int(4)] = d_677_v64_
                    nw128_[int(5)] = d_677_v64_
                    nw128_[int(6)] = d_677_v64_
                    nw128_[int(7)] = (d_678_v65_).cf39
                    nw128_[int(8)] = d_677_v64_
                    nw128_[int(9)] = d_677_v64_
                    nw128_[int(10)] = (d_679_v66_)[default__.safeIndex(270, len(d_679_v66_))]
                    d_680_v67_ = nw128_
                    index133_ = default__.safeIndex(419, (d_680_v67_).length(0))
                    (d_680_v67_)[index133_] = d_677_v64_
                    index134_ = default__.safeIndex(419, (d_680_v67_).length(0))
                    nw129_ = C3()
                    nw129_.ctor__()
                    (d_680_v67_)[index134_] = nw129_
                    d_681_v68_: _dafny.Map
                    d_681_v68_ = _dafny.Map({default__.safeModuloInt((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvbygep"))})
                    d_682_v69_: _dafny.Array
                    nw130_ = _dafny.Array(None, 3)
                    nw130_[int(0)] = d_592_v1_
                    nw130_[int(1)] = d_592_v1_
                    nw130_[int(2)] = d_592_v1_
                    d_682_v69_ = nw130_
                    d_683_v70_: _dafny.Array
                    d_683_v70_ = d_682_v69_
                    d_684_v71_: _dafny.Set
                    d_684_v71_ = _dafny.Set({(d_683_v70_), d_682_v69_})
                    d_681_v68_ = (d_681_v68_).set(default__.safeModuloInt(d_591_v0_, len(d_684_v71_)), ((self).f6) + ((self).f6))
                    d_685_v72_: _dafny.Array
                    nw131_ = _dafny.Array(_dafny.Array(None, 0), 27)
                    d_685_v72_ = nw131_
                    d_685_v72_ = p0
                    d_686_v73_: _dafny.Set
                    d_686_v73_ = _dafny.Set({False, True})
                    d_687_v74_: _dafny.Map
                    d_687_v74_ = _dafny.Map({(d_591_v0_) - (d_591_v0_): d_686_v73_})
                    d_688_v75_: _dafny.Map
                    d_688_v75_ = _dafny.Map({d_593_v2_: ((0) - ((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])) - ((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])})
                    index135_ = default__.safeIndex(175, (d_596_v5_).length(0))
                    rhs94_ = ((d_688_v75_)[(_dafny.MultiSet([(d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], d_591_v0_])).intersection(default__.fm31((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], globalState))] if ((_dafny.MultiSet([(d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], d_591_v0_])).intersection(default__.fm31((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], globalState))) in (d_688_v75_) else (0) - ((self).fm13(not(d_592_v1_), default__.fm24(d_591_v0_, False, d_592_v1_, len((self).f6), globalState), d_592_v1_, globalState)))
                    rhs95_ = d_687_v74_
                    lhs71_ = d_596_v5_
                    lhs72_ = default__.safeIndex(175, (d_596_v5_).length(0))
                    lhs71_[lhs72_] = rhs94_
                    d_687_v74_ = rhs95_
                    pass
            pass
        if d_592_v1_:
            d_689_v76_: _dafny.MultiSet
            d_689_v76_ = _dafny.MultiSet([d_592_v1_])
            d_690_v77_: D10
            d_690_v77_ = D10_DC21(d_689_v76_)
            d_691_v78_: _dafny.Map
            d_691_v78_ = _dafny.Map({len(d_595_v4_): (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]})
            d_692_v80_: _dafny.Map
            def iife37_():
                coll21_ = _dafny.Set()
                compr_21_: int
                for compr_21_ in (d_691_v78_).keys.Elements:
                    d_693_v79_: int = compr_21_
                    if (d_693_v79_) in (d_691_v78_):
                        coll21_ = coll21_.union(_dafny.Set([default__.safeModuloInt(d_693_v79_, len(_dafny.Map({(D9_DC19(-120, 390)).cf29: len(_dafny.Map({False: 768}))})))]))
                return _dafny.Set(coll21_)
            d_692_v80_ = _dafny.Map({d_690_v77_: len(iife37_()
            )})
            d_692_v80_ = (d_692_v80_).set(d_690_v77_, 283)
            d_694_v81_: C1
            nw132_ = C1()
            nw132_.ctor__((d_689_v76_).intersection(d_689_v76_), (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])
            d_694_v81_ = nw132_
            if d_592_v1_:
                index136_ = default__.safeIndex(175, (d_596_v5_).length(0))
                (d_596_v5_)[index136_] = d_591_v0_
                index137_ = default__.safeIndex(175, (d_596_v5_).length(0))
                (d_596_v5_)[index137_] = (0) - (-96)
                d_695_v82_: str
                d_695_v82_ = _dafny.CodePoint('v')
                d_695_v82_ = ((self).f6)[default__.safeIndex((0) - ((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]), len((self).f6))]
                d_592_v1_ = not(default__.fm33((self).f6, globalState))
                d_696_v83_: _dafny.Array
                nw133_ = _dafny.Array(False, 22)
                d_696_v83_ = nw133_
                d_696_v83_ = d_696_v83_
            elif True:
                index138_ = default__.safeIndex(175, (d_596_v5_).length(0))
                (d_596_v5_)[index138_] = d_694_v81_.f11
                d_595_v4_ = ((d_595_v4_) + (_dafny.SeqWithoutIsStrInference([d_592_v1_]))) + (d_595_v4_)
                d_697_v84_: C0
                nw134_ = C0()
                nw134_.ctor__()
                d_697_v84_ = nw134_
                nw135_ = _dafny.Array(int(0), 5)
                d_596_v5_ = nw135_
                d_591_v0_ = (0) - (d_694_v81_.f11)
            d_698_v85_: C0
            nw136_ = C0()
            nw136_.ctor__()
            d_698_v85_ = nw136_
            d_699_v86_: D7
            d_699_v86_ = D7_DC15(d_592_v1_, d_694_v81_.f11, d_591_v0_, len(_dafny.Map({(d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]: d_694_v81_.f11})), d_698_v85_)
            d_700_v87_: _dafny.Map
            d_700_v87_ = _dafny.Map({(d_699_v86_).cf24: d_592_v1_})
            d_701_v88_: _dafny.Map
            d_701_v88_ = _dafny.Map({d_700_v87_: d_591_v0_})
            d_702_v90_: str
            d_702_v90_ = _dafny.CodePoint('m')
            d_703_v91_: _dafny.Map
            d_703_v91_ = _dafny.Map({d_694_v81_.f11: _dafny.Map({d_702_v90_: d_592_v1_})})
            def iife38_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in (d_703_v91_).keys.Elements:
                    d_704_v89_: int = compr_22_
                    if (d_704_v89_) in (d_703_v91_):
                        coll22_[(d_704_v89_) - (d_694_v81_.f11)] = not(False)
                return _dafny.Map(coll22_)
            d_701_v88_ = (d_701_v88_).set(iife38_()
            , (self).fm13(d_592_v1_, ((d_691_v78_)[-45] if (-45) in (d_691_v78_) else d_591_v0_), d_592_v1_, globalState))
            d_705_v92_: _dafny.Array
            nw137_ = _dafny.Array(_dafny.CodePoint('D'), 29)
            d_705_v92_ = nw137_
            d_705_v92_ = d_705_v92_
        elif True:
            d_706_v94_: _dafny.Array
            def lambda23_(d_707_v5_, d_708_v2_):
                def lambda24_(d_709_i6_):
                    def iife39_():
                        coll23_ = _dafny.Set()
                        compr_23_: int
                        for compr_23_ in (d_708_v2_).Elements:
                            d_710_v93_: int = compr_23_
                            if (d_710_v93_) in (d_708_v2_):
                                coll23_ = coll23_.union(_dafny.Set([(d_710_v93_) * (647)]))
                        return _dafny.Set(coll23_)
                    return (_dafny.Set({(d_707_v5_)[default__.safeIndex(175, (d_707_v5_).length(0))], (d_707_v5_)[default__.safeIndex(175, (d_707_v5_).length(0))]})) | (iife39_()
                    )

                return lambda24_

            init13_ = lambda23_(d_596_v5_, d_593_v2_)
            nw138_ = _dafny.Array(None, 25)
            for i0_13_ in range(nw138_.length(0)):
                nw138_[i0_13_] = init13_(i0_13_)
            d_706_v94_ = nw138_
            d_711_v95_: _dafny.Set
            d_711_v95_ = _dafny.Set({(d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]})
            index139_ = default__.safeIndex(201, (d_706_v94_).length(0))
            def iife40_():
                coll24_ = _dafny.Set()
                compr_24_: int
                for compr_24_ in _dafny.IntegerRange(734, 782):
                    d_712_v96_: int = compr_24_
                    if ((734) <= (d_712_v96_)) and ((d_712_v96_) < (782)):
                        coll24_ = coll24_.union(_dafny.Set([(d_712_v96_) - ((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])]))
                return _dafny.Set(coll24_)
            (d_706_v94_)[index139_] = (d_711_v95_) - (iife40_()
            )
            index140_ = default__.safeIndex(201, (d_706_v94_).length(0))
            def iife41_():
                coll25_ = _dafny.Set()
                compr_25_: int
                for compr_25_ in (_dafny.SeqWithoutIsStrInference([d_591_v0_ for d_713_i7_ in range(default__.abs(9))])).Elements:
                    d_714_v97_: int = compr_25_
                    if (d_714_v97_) in (_dafny.SeqWithoutIsStrInference([d_591_v0_ for d_713_i7_ in range(default__.abs(9))])):
                        coll25_ = coll25_.union(_dafny.Set([(d_714_v97_) * (907)]))
                return _dafny.Set(coll25_)
            (d_706_v94_)[index140_] = (d_711_v95_) - (iife41_()
            )
            d_715_v98_: _dafny.Set
            d_715_v98_ = _dafny.Set({d_592_v1_, d_592_v1_})
            d_716_v99_: _dafny.MultiSet
            d_716_v99_ = _dafny.MultiSet([_dafny.Map({(0) - (len(d_715_v98_)): (0) - ((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])}), _dafny.Map({default__.fm24(d_591_v0_, d_592_v1_, d_592_v1_, d_591_v0_, globalState): d_591_v0_})])
            d_717_v100_: _dafny.Array
            nw139_ = _dafny.Array(None, 25)
            nw139_[int(0)] = True
            nw139_[int(1)] = d_592_v1_
            nw139_[int(2)] = True
            nw139_[int(3)] = d_592_v1_
            nw139_[int(4)] = d_592_v1_
            nw139_[int(5)] = (d_592_v1_) == (not(d_592_v1_))
            nw139_[int(6)] = (d_716_v99_).issubset(d_716_v99_)
            nw139_[int(7)] = d_592_v1_
            nw139_[int(8)] = d_592_v1_
            nw139_[int(9)] = d_592_v1_
            nw139_[int(10)] = d_592_v1_
            nw139_[int(11)] = (d_592_v1_) in ((_dafny.MultiSet(d_595_v4_)).set(d_592_v1_, default__.abs((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])))
            nw139_[int(12)] = not((d_591_v0_) <= (317))
            nw139_[int(13)] = d_592_v1_
            nw139_[int(14)] = not (d_592_v1_) or (d_592_v1_)
            nw139_[int(15)] = d_592_v1_
            nw139_[int(16)] = not((d_592_v1_ if d_592_v1_ else d_592_v1_))
            nw139_[int(17)] = d_592_v1_
            nw139_[int(18)] = d_592_v1_
            nw139_[int(19)] = d_592_v1_
            nw139_[int(20)] = d_592_v1_
            nw139_[int(21)] = d_592_v1_
            nw139_[int(22)] = default__.fm33(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lemmjobsq")), globalState)
            nw139_[int(23)] = not(d_592_v1_)
            nw139_[int(24)] = d_592_v1_
            d_717_v100_ = nw139_
            index141_ = default__.safeIndex(49, (d_717_v100_).length(0))
            (d_717_v100_)[index141_] = (d_592_v1_) == (d_592_v1_)
            index142_ = default__.safeIndex(49, (d_717_v100_).length(0))
            rhs96_ = d_595_v4_
            rhs97_ = d_596_v5_
            rhs98_ = d_592_v1_
            rhs99_ = ((self).fm13(d_592_v1_, (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))], True, globalState)) > (d_591_v0_)
            lhs73_ = d_717_v100_
            lhs74_ = default__.safeIndex(49, (d_717_v100_).length(0))
            d_595_v4_ = rhs96_
            d_596_v5_ = rhs97_
            lhs73_[lhs74_] = rhs98_
            r0 = rhs99_
            rhs100_ = d_596_v5_
            rhs101_ = d_594_v3_
            d_596_v5_ = rhs100_
            d_594_v3_ = rhs101_
            d_718_v101_: _dafny.Array
            nw140_ = _dafny.Array(_dafny.Map({}), 24)
            d_718_v101_ = nw140_
            d_719_v102_: _dafny.Map
            d_719_v102_ = _dafny.Map({d_718_v101_: default__.safeModuloInt((0) - ((d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]), (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])})
            d_719_v102_ = (d_719_v102_).set(d_718_v101_, (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))])
            index143_ = default__.safeIndex(175, (d_596_v5_).length(0))
            (d_596_v5_)[index143_] = (d_596_v5_)[default__.safeIndex(175, (d_596_v5_).length(0))]
        r0 = d_592_v1_
        d_720_v103_: _dafny.MultiSet
        d_720_v103_ = _dafny.MultiSet([d_592_v1_])
        d_721_v104_: _dafny.Map
        d_721_v104_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_722_i8_ in range(default__.abs(706))])): d_720_v103_})
        d_723_v105_: _dafny.Seq
        d_723_v105_ = _dafny.SeqWithoutIsStrInference([(d_721_v104_ if d_592_v1_ else d_721_v104_), d_721_v104_, d_721_v104_, d_721_v104_])
        r1 = (d_723_v105_)[default__.safeIndex((0) - (default__.safeModuloInt(len((self).f6), d_591_v0_)), len(d_723_v105_))]
        return r0, r1

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        r1: str = _dafny.CodePoint('D')
        d_724_v0_: int
        d_724_v0_ = -822
        d_725_v1_: _dafny.Seq
        d_725_v1_ = _dafny.SeqWithoutIsStrInference([d_724_v0_, d_724_v0_])
        d_726_v2_: D1
        d_726_v2_ = D1_DC4((d_725_v1_)[default__.safeIndex(d_724_v0_, len(d_725_v1_))], _dafny.CodePoint('b'), False, len((self).f6))
        d_727_i0_: int
        d_727_i0_ = 0
        with _dafny.label("6"):
            pat_let_tv5_ = p0
            pat_let_tv6_ = p0
            def lambda25_(source9_):
                if source9_.is_DC3:
                    d_738___mcc_h0_ = source9_.cf5
                    d_739_cf5_ = d_738___mcc_h0_
                    return not((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])])) < (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([pat_let_tv5_, False]) for d_740_i1_ in range(default__.abs(755))])))
                elif source9_.is_DC4:
                    d_741___mcc_h1_ = source9_.cf6
                    d_742___mcc_h2_ = source9_.cf7
                    d_743___mcc_h3_ = source9_.cf8
                    d_744___mcc_h4_ = source9_.cf9
                    d_745_cf9_ = d_744___mcc_h4_
                    d_746_cf8_ = d_743___mcc_h3_
                    d_747_cf7_ = d_742___mcc_h2_
                    d_748_cf6_ = d_741___mcc_h1_
                    return pat_let_tv6_
                elif True:
                    d_749___mcc_h5_ = source9_.cf4
                    d_750_cf4_ = d_749___mcc_h5_
                    return not(not(d_750_cf4_))

            while lambda25_(d_726_v2_):
                with _dafny.c_label("6"):
                    if (d_727_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_727_i0_ = (d_727_i0_) + (1)
                    d_728_v3_: bool
                    d_728_v3_ = False
                    d_728_v3_ = p0
                    d_729_v4_: _dafny.Array
                    nw141_ = _dafny.Array(int(0), 24)
                    d_729_v4_ = nw141_
                    d_730_v5_: D2
                    d_730_v5_ = D2_DC6(d_724_v0_, p0)
                    index144_ = default__.safeIndex(185, (d_729_v4_).length(0))
                    (d_729_v4_)[index144_] = (d_730_v5_).cf11
                    d_731_v7_: _dafny.Map
                    d_731_v7_ = _dafny.Map({d_728_v3_: d_728_v3_})
                    d_732_v8_: _dafny.Set
                    d_732_v8_ = _dafny.Set({len(d_731_v7_), (0) - (d_724_v0_)})
                    d_733_v9_: _dafny.Seq
                    def iife42_():
                        coll26_ = _dafny.Set()
                        compr_26_: int
                        for compr_26_ in _dafny.IntegerRange(553, 891):
                            d_734_v6_: int = compr_26_
                            if ((553) <= (d_734_v6_)) and ((d_734_v6_) < (891)):
                                coll26_ = coll26_.union(_dafny.Set([default__.safeModuloInt(d_734_v6_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ba"))))]))
                        return _dafny.Set(coll26_)
                    d_733_v9_ = _dafny.SeqWithoutIsStrInference([d_728_v3_, (d_732_v8_).ispropersubset(iife42_()
                    ), p0, not(d_728_v3_)])
                    d_735_v10_: _dafny.Seq
                    d_735_v10_ = _dafny.SeqWithoutIsStrInference([d_733_v9_])
                    index145_ = default__.safeIndex(185, (d_729_v4_).length(0))
                    rhs102_ = (d_724_v0_) + (d_724_v0_)
                    rhs103_ = default__.safeDivisionInt(d_724_v0_, (355 if d_728_v3_ else d_724_v0_))
                    rhs104_ = d_724_v0_
                    rhs105_ = ((default__.fm26(d_724_v0_, globalState)) + (d_733_v9_)) + ((d_735_v10_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aaqjua"))), len(d_735_v10_))])
                    lhs75_ = d_729_v4_
                    lhs76_ = default__.safeIndex(185, (d_729_v4_).length(0))
                    r0 = rhs102_
                    r0 = rhs103_
                    lhs75_[lhs76_] = rhs104_
                    d_733_v9_ = rhs105_
                    d_736_v11_: C3
                    nw142_ = C3()
                    nw142_.ctor__()
                    d_736_v11_ = nw142_
                    d_737_v12_: _dafny.Map
                    d_737_v12_ = _dafny.Map({default__.safeModuloInt((d_729_v4_)[default__.safeIndex(185, (d_729_v4_).length(0))], (d_729_v4_)[default__.safeIndex(185, (d_729_v4_).length(0))]): (self).f6})
                    index146_ = default__.safeIndex(185, (d_729_v4_).length(0))
                    index147_ = default__.safeIndex(185, (d_729_v4_).length(0))
                    rhs106_ = len(d_737_v12_)
                    rhs107_ = -470
                    rhs108_ = d_728_v3_
                    lhs77_ = d_729_v4_
                    lhs78_ = default__.safeIndex(185, (d_729_v4_).length(0))
                    lhs79_ = d_729_v4_
                    lhs80_ = default__.safeIndex(185, (d_729_v4_).length(0))
                    lhs77_[lhs78_] = rhs106_
                    lhs79_[lhs80_] = rhs107_
                    d_728_v3_ = rhs108_
                    pass
            pass
        d_751_v13_: _dafny.Array
        def lambda26_(d_752_v0_):
            def lambda27_(d_753_i2_):
                return (d_753_i2_) + (d_752_v0_)

            return lambda27_

        init14_ = lambda26_(d_724_v0_)
        nw143_ = _dafny.Array(None, 24)
        for i0_14_ in range(nw143_.length(0)):
            nw143_[i0_14_] = init14_(i0_14_)
        d_751_v13_ = nw143_
        d_751_v13_ = d_751_v13_
        d_754_v14_: _dafny.Array
        nw144_ = _dafny.Array(False, 13)
        d_754_v14_ = nw144_
        d_755_v15_: _dafny.Map
        d_755_v15_ = _dafny.Map({len((self).f6): d_754_v14_})
        if (d_755_v15_) == (d_755_v15_):
            index148_ = default__.safeIndex(172, (d_754_v14_).length(0))
            (d_754_v14_)[index148_] = p0
            d_756_v16_: _dafny.Array
            nw145_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
            d_756_v16_ = nw145_
            index149_ = default__.safeIndex(133, (d_756_v16_).length(0))
            (d_756_v16_)[index149_] = ((self).f6) + ((self).f6)
            index150_ = default__.safeIndex(172, (d_754_v14_).length(0))
            index151_ = default__.safeIndex(133, (d_756_v16_).length(0))
            rhs109_ = p0
            rhs110_ = len(((self).f6) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_757_i3_ in range(default__.abs(948))])))
            rhs111_ = (((self).f6) + ((self).f6)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npyc")))
            lhs81_ = d_754_v14_
            lhs82_ = default__.safeIndex(172, (d_754_v14_).length(0))
            lhs83_ = d_756_v16_
            lhs84_ = default__.safeIndex(133, (d_756_v16_).length(0))
            lhs81_[lhs82_] = rhs109_
            r0 = rhs110_
            lhs83_[lhs84_] = rhs111_
            d_724_v0_ = (-516) + ((len(_dafny.Map({d_724_v0_: p0}))) - ((d_725_v1_)[default__.safeIndex(d_724_v0_, len(d_725_v1_))]))
            index152_ = default__.safeIndex(133, (d_756_v16_).length(0))
            (d_756_v16_)[index152_] = (self).f6
            index153_ = default__.safeIndex(172, (d_754_v14_).length(0))
            (d_754_v14_)[index153_] = (d_754_v14_)[default__.safeIndex(172, (d_754_v14_).length(0))]
            index154_ = default__.safeIndex(172, (d_754_v14_).length(0))
            (d_754_v14_)[index154_] = p0
        elif True:
            r0 = -976
            d_758_v17_: bool
            d_758_v17_ = False
            d_759_v18_: _dafny.Seq
            d_759_v18_ = _dafny.SeqWithoutIsStrInference([p0, d_758_v17_])
            d_760_v19_: D1
            d_760_v19_ = D1_DC2((d_759_v18_)[default__.safeIndex(d_724_v0_, len(d_759_v18_))])
            d_761_v20_: str
            d_761_v20_ = _dafny.CodePoint('g')
            pat_let_tv7_ = d_724_v0_
            pat_let_tv8_ = d_761_v20_
            pat_let_tv9_ = d_758_v17_
            def iife44_(_pat_let9_0):
                def iife45_(d_762_dt__update__tmp_h0_):
                    def iife46_(_pat_let10_0):
                        def iife47_(d_763_dt__update_hcf9_h0_):
                            def iife48_(_pat_let11_0):
                                def iife49_(d_764_dt__update_hcf7_h0_):
                                    return D1_DC4((d_762_dt__update__tmp_h0_).cf6, d_764_dt__update_hcf7_h0_, (d_762_dt__update__tmp_h0_).cf8, d_763_dt__update_hcf9_h0_)
                                return iife49_(_pat_let11_0)
                            return iife48_(pat_let_tv8_)
                        return iife47_(_pat_let10_0)
                    return iife46_(pat_let_tv7_)
                return iife45_(_pat_let9_0)
            def iife43_(_pat_let8_0):
                def iife50_(d_765_dt__update__tmp_h1_):
                    def iife51_(_pat_let12_0):
                        def iife52_(d_766_dt__update_hcf8_h0_):
                            return D1_DC4((d_765_dt__update__tmp_h1_).cf6, (d_765_dt__update__tmp_h1_).cf7, d_766_dt__update_hcf8_h0_, (d_765_dt__update__tmp_h1_).cf9)
                        return iife52_(_pat_let12_0)
                    return iife51_(pat_let_tv9_)
                return iife50_(_pat_let8_0)
            rhs112_ = p0
            rhs113_ = not(not (False) or (False))
            rhs114_ = (d_760_v19_).cf4
            rhs115_ = (iife43_(iife44_(d_726_v2_))).cf7
            rhs116_ = default__.fm33((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsfeu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mosbxp"))), globalState)
            d_758_v17_ = rhs112_
            d_758_v17_ = rhs113_
            d_758_v17_ = rhs114_
            r1 = rhs115_
            d_758_v17_ = rhs116_
            d_725_v1_ = default__.fm23((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cho")))) * (d_724_v0_), globalState)
            d_767_v21_: _dafny.Map
            d_767_v21_ = _dafny.Map({(False) == (p0): d_751_v13_})
            d_767_v21_ = (d_767_v21_).set(d_758_v17_, d_751_v13_)
            d_768_v22_: _dafny.Map
            d_768_v22_ = _dafny.Map({p0: True})
            d_758_v17_ = (d_758_v17_) in ((d_768_v22_) | (d_768_v22_))
        if p0:
            d_769_v23_: _dafny.Array
            nw146_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_769_v23_ = nw146_
            d_770_v24_: _dafny.Map
            d_770_v24_ = _dafny.Map({d_724_v0_: D0_DC0(d_724_v0_, (self).f6, d_769_v23_)})
            d_771_v25_: _dafny.Seq
            d_771_v25_ = _dafny.SeqWithoutIsStrInference([d_770_v24_])
            d_772_v26_: T1
            nw147_ = C3()
            nw147_.ctor__()
            d_772_v26_ = nw147_
            d_773_v27_: D11
            d_773_v27_ = D11_DC26(d_771_v25_, d_772_v26_)
            pat_let_tv10_ = d_770_v24_
            pat_let_tv11_ = d_770_v24_
            pat_let_tv12_ = d_770_v24_
            pat_let_tv13_ = d_770_v24_
            d_774_v28_: _dafny.Map
            def iife53_(_pat_let13_0):
                def iife54_(d_776_dt__update__tmp_h2_):
                    def iife55_(_pat_let14_0):
                        def iife56_(d_777_dt__update_hcf42_h0_):
                            return D11_DC26(d_777_dt__update_hcf42_h0_, (d_776_dt__update__tmp_h2_).cf43)
                        return iife56_(_pat_let14_0)
                    return iife55_(_dafny.SeqWithoutIsStrInference([pat_let_tv10_, pat_let_tv11_, pat_let_tv12_, pat_let_tv13_]))
                return iife54_(_pat_let13_0)
            d_774_v28_ = _dafny.Map({((self).f6) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_775_i4_ in range(default__.abs(748))])): iife53_(d_773_v27_)})
            d_774_v28_ = (d_774_v28_) | ((_dafny.Map({p0: d_773_v27_})) | (_dafny.Map({p0: d_773_v27_})))
            d_778_v29_: _dafny.Seq
            d_778_v29_ = _dafny.SeqWithoutIsStrInference([p0, p0, not (p0) or (p0), p0])
            d_778_v29_ = (d_778_v29_) + ((d_778_v29_) + (d_778_v29_))
            d_724_v0_ = d_724_v0_
            d_779_v30_: bool
            d_779_v30_ = True
            d_780_v31_: D0
            d_780_v31_ = D0_DC0(d_724_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqbg")), d_769_v23_)
            d_779_v30_ = default__.fm33((d_780_v31_).cf1, globalState)
            d_781_v32_: C1
            nw148_ = C1()
            nw148_.ctor__(_dafny.MultiSet(d_778_v29_), default__.safeDivisionInt(d_724_v0_, d_724_v0_))
            d_781_v32_ = nw148_
        elif True:
            d_782_v33_: C0
            nw149_ = C0()
            nw149_.ctor__()
            d_782_v33_ = nw149_
            d_724_v0_ = d_724_v0_
            d_724_v0_ = -933
            d_783_v34_: bool
            d_783_v34_ = False
            d_784_v35_: _dafny.Seq
            d_784_v35_ = _dafny.SeqWithoutIsStrInference([d_783_v34_])
            d_783_v34_ = not (d_783_v34_) or ((d_784_v35_)[default__.safeIndex(len(_dafny.Map({d_783_v34_: d_783_v34_})), len(d_784_v35_))])
            index155_ = default__.safeIndex(621, (d_751_v13_).length(0))
            (d_751_v13_)[index155_] = d_724_v0_
            index156_ = default__.safeIndex(621, (d_751_v13_).length(0))
            (d_751_v13_)[index156_] = d_724_v0_
        d_785_v36_: bool
        d_785_v36_ = False
        d_785_v36_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybptaos"))) < (((self).f6) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgjwul"))))
        d_786_v37_: _dafny.Seq
        d_786_v37_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_724_v0_])])
        d_787_i5_: int
        d_787_i5_ = 0
        with _dafny.label("7"):
            while ((d_724_v0_) + (d_724_v0_)) == ((_dafny.MultiSet((d_786_v37_)[default__.safeIndex(d_724_v0_, len(d_786_v37_))])).cardinality):
                with _dafny.c_label("7"):
                    if (d_787_i5_) >= (100):
                        raise _dafny.Break("7")
                    d_787_i5_ = (d_787_i5_) + (1)
                    d_788_v38_: _dafny.Seq
                    d_788_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
                    d_789_v39_: _dafny.Set
                    d_789_v39_ = _dafny.Set({d_724_v0_, (0) - (d_724_v0_)})
                    d_790_v40_: _dafny.Map
                    d_790_v40_ = _dafny.Map({d_724_v0_: d_789_v39_})
                    index157_ = default__.safeIndex(257, (d_754_v14_).length(0))
                    (d_754_v14_)[index157_] = (((d_790_v40_)[d_724_v0_] if (d_724_v0_) in (d_790_v40_) else d_789_v39_)).issubset(default__.fm20(d_785_v36_, d_788_v38_, default__.fm33(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_791_i6_ in range(default__.abs(290))]), globalState), globalState))
                    index158_ = default__.safeIndex(257, (d_754_v14_).length(0))
                    rhs117_ = ((self).f6) < ((default__.fm22(d_785_v36_, d_724_v0_, globalState)) + ((self).f6))
                    rhs118_ = ((self).f6) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xq")))
                    rhs119_ = d_785_v36_
                    rhs120_ = d_785_v36_
                    rhs121_ = d_724_v0_
                    lhs85_ = d_754_v14_
                    lhs86_ = default__.safeIndex(257, (d_754_v14_).length(0))
                    d_785_v36_ = rhs117_
                    d_788_v38_ = rhs118_
                    lhs85_[lhs86_] = rhs119_
                    d_785_v36_ = rhs120_
                    r0 = rhs121_
                    d_792_v41_: _dafny.MultiSet
                    d_792_v41_ = _dafny.MultiSet([259, default__.safeModuloInt(-880, (0) - (d_724_v0_))])
                    d_792_v41_ = d_792_v41_
                    d_793_v42_: D4
                    d_793_v42_ = D4_DC9(d_751_v13_)
                    d_794_v43_: _dafny.Array
                    nw150_ = _dafny.Array(None, 2)
                    nw150_[int(0)] = d_751_v13_
                    nw150_[int(1)] = (d_793_v42_).cf15
                    d_794_v43_ = nw150_
                    index159_ = default__.safeIndex(404, (d_794_v43_).length(0))
                    (d_794_v43_)[index159_] = d_751_v13_
                    d_795_v44_: _dafny.Map
                    d_795_v44_ = _dafny.Map({p0: d_724_v0_})
                    d_796_v45_: _dafny.MultiSet
                    d_796_v45_ = _dafny.MultiSet([(d_754_v14_)[default__.safeIndex(257, (d_754_v14_).length(0))], p0, False, d_785_v36_])
                    index160_ = default__.safeIndex(404, (d_794_v43_).length(0))
                    nw151_ = _dafny.Array(None, 24)
                    nw151_[int(0)] = d_724_v0_
                    nw151_[int(1)] = d_724_v0_
                    nw151_[int(2)] = d_724_v0_
                    nw151_[int(3)] = d_724_v0_
                    nw151_[int(4)] = d_724_v0_
                    nw151_[int(5)] = (0) - ((d_724_v0_) - (d_724_v0_))
                    nw151_[int(6)] = d_724_v0_
                    nw151_[int(7)] = d_724_v0_
                    nw151_[int(8)] = default__.safeDivisionInt(d_724_v0_, (0) - (d_724_v0_))
                    nw151_[int(9)] = d_724_v0_
                    nw151_[int(10)] = d_724_v0_
                    nw151_[int(11)] = (len(d_795_v44_) if d_785_v36_ else d_724_v0_)
                    nw151_[int(12)] = d_724_v0_
                    nw151_[int(13)] = d_724_v0_
                    nw151_[int(14)] = d_724_v0_
                    nw151_[int(15)] = (len(_dafny.Set({not(default__.fm33((self).f6, globalState))}))) - (902)
                    nw151_[int(16)] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_785_v36_])), 351)
                    nw151_[int(17)] = d_724_v0_
                    nw151_[int(18)] = (0) - (d_724_v0_)
                    nw151_[int(19)] = d_724_v0_
                    nw151_[int(20)] = default__.fm14(((d_796_v45_)[p0] if (p0) in (d_796_v45_) else d_724_v0_), d_724_v0_, globalState)
                    nw151_[int(21)] = len((self).f6)
                    nw151_[int(22)] = (_dafny.MultiSet([not((d_754_v14_)[default__.safeIndex(257, (d_754_v14_).length(0))])])).cardinality
                    nw151_[int(23)] = d_724_v0_
                    (d_794_v43_)[index160_] = nw151_
                    d_797_v46_: _dafny.Seq
                    d_797_v46_ = _dafny.SeqWithoutIsStrInference([d_785_v36_, True])
                    d_798_v47_: _dafny.Seq
                    d_798_v47_ = _dafny.SeqWithoutIsStrInference([(len(d_797_v46_)) >= (d_724_v0_)])
                    index161_ = default__.safeIndex(257, (d_754_v14_).length(0))
                    (d_754_v14_)[index161_] = (d_798_v47_)[default__.safeIndex((d_725_v1_)[default__.safeIndex(d_724_v0_, len(d_725_v1_))], len(d_798_v47_))]
                    pass
            pass
        r0 = d_724_v0_
        d_799_v48_: str
        d_799_v48_ = _dafny.CodePoint('j')
        r1 = d_799_v48_
        return r0, r1

    @property
    def f6(self):
        return self._f6

class C6(T0):
    def  __init__(self):
        self.f4: _dafny.MultiSet = _dafny.MultiSet({})
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f4, f5):
        (self).f4 = f4
        (self)._f5 = f5

    def fm0(self, p0, p1, p2, p3, globalState):
        return ((self.f4) | (self.f4)) - (self.f4)

    def fm9(self, globalState):
        return len(((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_800_i0_ in range(default__.abs(788))])): False}) if False else _dafny.Map({len(_dafny.Map({True: False})): False}))) | ((_dafny.Map({(0) - (len(_dafny.Map({True: False}))): True}) if False else _dafny.Map({440: True}))))

    def fm10(self, p0, p1, p2, globalState):
        return (D2_DC6(929, False)).cf12

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_801_v0_: _dafny.Array
        nw152_ = _dafny.Array(_dafny.Array(None, 0), 22)
        d_801_v0_ = nw152_
        d_802_v1_: _dafny.Seq
        d_802_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxl"))
        d_803_v2_: str
        d_803_v2_ = _dafny.CodePoint('v')
        d_804_v3_: int
        d_804_v3_ = 403
        d_805_v4_: bool
        d_805_v4_ = False
        d_806_v5_: D1
        d_806_v5_ = D1_DC2(True)
        d_807_v6_: _dafny.Array
        nw153_ = _dafny.Array(None, 23)
        nw153_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_808_i0_ in range(default__.abs(656))])
        nw153_[int(1)] = d_802_v1_
        nw153_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pet"))
        nw153_[int(3)] = d_802_v1_
        nw153_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpfmyj"))
        nw153_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipd"))
        nw153_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "voauqmys"))
        nw153_[int(7)] = d_802_v1_
        nw153_[int(8)] = d_802_v1_
        nw153_[int(9)] = d_802_v1_
        nw153_[int(10)] = d_802_v1_
        nw153_[int(11)] = default__.fm11(d_803_v2_, d_804_v3_, d_805_v4_, d_806_v5_, globalState)
        nw153_[int(12)] = d_802_v1_
        nw153_[int(13)] = d_802_v1_
        nw153_[int(14)] = d_802_v1_
        nw153_[int(15)] = d_802_v1_
        nw153_[int(16)] = d_802_v1_
        nw153_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
        nw153_[int(18)] = d_802_v1_
        nw153_[int(19)] = d_802_v1_
        nw153_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usguexh"))
        nw153_[int(21)] = d_802_v1_
        nw153_[int(22)] = d_802_v1_
        d_807_v6_ = nw153_
        index162_ = default__.safeIndex(44, (d_801_v0_).length(0))
        (d_801_v0_)[index162_] = d_807_v6_
        d_809_v7_: D2
        d_809_v7_ = D2_DC6(d_804_v3_, d_805_v4_)
        d_810_v8_: _dafny.Seq
        d_810_v8_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_804_v3_, 294]) for d_811_i1_ in range(default__.abs(166))]))])
        d_812_v9_: _dafny.Seq
        d_812_v9_ = _dafny.SeqWithoutIsStrInference([d_810_v8_, d_810_v8_, d_810_v8_, _dafny.SeqWithoutIsStrInference([d_804_v3_ for d_813_i2_ in range(default__.abs(906))]), d_810_v8_])
        d_814_v10_: _dafny.MultiSet
        d_814_v10_ = _dafny.MultiSet([d_804_v3_])
        d_815_v11_: _dafny.Map
        d_815_v11_ = _dafny.Map({d_805_v4_: d_814_v10_})
        index163_ = default__.safeIndex(44, (d_801_v0_).length(0))
        rhs122_ = d_807_v6_
        rhs123_ = (self).fm10(d_809_v7_, (len((d_812_v9_)[default__.safeIndex(len(default__.fm12(d_804_v3_, d_804_v3_, globalState)), len(d_812_v9_))])) + (933), ((d_815_v11_)[d_805_v4_] if (d_805_v4_) in (d_815_v11_) else d_814_v10_), globalState)
        rhs124_ = False
        lhs87_ = d_801_v0_
        lhs88_ = default__.safeIndex(44, (d_801_v0_).length(0))
        lhs87_[lhs88_] = rhs122_
        r0 = rhs123_
        d_805_v4_ = rhs124_
        d_816_i3_: int
        d_816_i3_ = 0
        with _dafny.label("8"):
            while not((d_802_v1_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omtjm")))):
                with _dafny.c_label("8"):
                    if (d_816_i3_) >= (100):
                        raise _dafny.Break("8")
                    d_816_i3_ = (d_816_i3_) + (1)
                    rhs125_ = d_806_v5_
                    rhs126_ = (0) - (((d_804_v3_) + (d_804_v3_)) - (199))
                    d_806_v5_ = rhs125_
                    d_804_v3_ = rhs126_
                    d_817_v12_: _dafny.Seq
                    d_817_v12_ = _dafny.SeqWithoutIsStrInference([False, d_805_v4_, True, d_805_v4_])
                    d_804_v3_ = default__.safeModuloInt(len(d_817_v12_), d_804_v3_)
                    d_818_v13_: int
                    d_819_v14_: _dafny.Map
                    d_820_v15_: bool
                    out14_: int
                    out15_: _dafny.Map
                    out16_: bool
                    out14_, out15_, out16_ = (self).m2(globalState)
                    d_818_v13_ = out14_
                    d_819_v14_ = out15_
                    d_820_v15_ = out16_
                    if d_820_v15_:
                        index164_ = default__.safeIndex(995, ((self).f5).length(0))
                        ((self).f5)[index164_] = (self).fm9(globalState)
                        index165_ = default__.safeIndex(995, ((self).f5).length(0))
                        ((self).f5)[index165_] = (d_809_v7_).cf11
                        d_810_v8_ = ((d_810_v8_) + ((_dafny.SeqWithoutIsStrInference([174 for d_821_i4_ in range(default__.abs(20))])).set(default__.safeIndex(d_818_v13_, len(_dafny.SeqWithoutIsStrInference([174 for d_822_i4_ in range(default__.abs(20))]))), 844))).set(default__.safeIndex(((d_814_v10_)[((self).f5)[default__.safeIndex(995, ((self).f5).length(0))]] if (((self).f5)[default__.safeIndex(995, ((self).f5).length(0))]) in (d_814_v10_) else -687), len((d_810_v8_) + ((_dafny.SeqWithoutIsStrInference([174 for d_823_i4_ in range(default__.abs(20))])).set(default__.safeIndex(d_818_v13_, len(_dafny.SeqWithoutIsStrInference([174 for d_824_i4_ in range(default__.abs(20))]))), 844)))), -367)
                        index166_ = default__.safeIndex(995, ((self).f5).length(0))
                        ((self).f5)[index166_] = d_818_v13_
                        index167_ = default__.safeIndex(316, (d_807_v6_).length(0))
                        (d_807_v6_)[index167_] = d_802_v1_
                        index168_ = default__.safeIndex(316, (d_807_v6_).length(0))
                        (d_807_v6_)[index168_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgiljws")) if not(d_820_v15_) else d_802_v1_)
                        d_818_v13_ = (d_810_v8_)[default__.safeIndex(((self).f5)[default__.safeIndex(995, ((self).f5).length(0))], len(d_810_v8_))]
                    elif True:
                        d_825_v16_: _dafny.Array
                        nw154_ = _dafny.Array(False, 22)
                        d_825_v16_ = nw154_
                        index169_ = default__.safeIndex(937, (d_825_v16_).length(0))
                        (d_825_v16_)[index169_] = d_820_v15_
                        index170_ = default__.safeIndex(937, (d_825_v16_).length(0))
                        (d_825_v16_)[index170_] = not((self).fm10(d_809_v7_, d_804_v3_, d_814_v10_, globalState))
                        d_810_v8_ = (d_810_v8_).set(default__.safeIndex((0) - (len(d_802_v1_)), len(d_810_v8_)), d_804_v3_)
                        d_804_v3_ = len(d_810_v8_)
                        index171_ = default__.safeIndex(256, (d_807_v6_).length(0))
                        (d_807_v6_)[index171_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvrjpar"))
                        index172_ = default__.safeIndex(256, (d_807_v6_).length(0))
                        (d_807_v6_)[index172_] = d_802_v1_
                        d_826_v17_: C1
                        nw155_ = C1()
                        nw155_.ctor__(self.f4, d_818_v13_)
                        d_826_v17_ = nw155_
                    pass
            pass
        d_827_i5_: int
        d_827_i5_ = 0
        with _dafny.label("9"):
            while d_805_v4_:
                with _dafny.c_label("9"):
                    if (d_827_i5_) >= (100):
                        raise _dafny.Break("9")
                    d_827_i5_ = (d_827_i5_) + (1)
                    d_828_v18_: _dafny.Array
                    nw156_ = _dafny.Array(_dafny.Array(None, 0), 19)
                    d_828_v18_ = nw156_
                    d_829_v19_: _dafny.Map
                    d_829_v19_ = _dafny.Map({d_804_v3_: D0_DC0(d_804_v3_, d_802_v1_, d_828_v18_)})
                    d_830_v20_: _dafny.Seq
                    d_830_v20_ = _dafny.SeqWithoutIsStrInference([d_829_v19_, d_829_v19_, d_829_v19_])
                    d_831_v21_: T1
                    nw157_ = C1()
                    nw157_.ctor__(self.f4, d_804_v3_)
                    d_831_v21_ = nw157_
                    d_832_v22_: D11
                    d_832_v22_ = D11_DC26(d_830_v20_, d_831_v21_)
                    d_833_v23_: _dafny.MultiSet
                    d_833_v23_ = _dafny.MultiSet([d_832_v22_])
                    d_834_v24_: _dafny.Seq
                    d_834_v24_ = _dafny.SeqWithoutIsStrInference([d_805_v4_])
                    d_835_v25_: _dafny.MultiSet
                    d_835_v25_ = _dafny.MultiSet([d_834_v24_, default__.fm12(d_804_v3_, d_804_v3_, globalState)])
                    rhs127_ = d_833_v23_
                    rhs128_ = default__.safeDivisionInt((d_835_v25_).cardinality, (d_804_v3_ if d_805_v4_ else 579))
                    d_833_v23_ = rhs127_
                    d_804_v3_ = rhs128_
                    d_836_v26_: _dafny.Map
                    d_836_v26_ = _dafny.Map({default__.fm33(d_802_v1_, globalState): _dafny.Map({d_834_v24_: d_804_v3_})})
                    d_837_v27_: _dafny.Map
                    d_837_v27_ = _dafny.Map({d_834_v24_: d_804_v3_})
                    d_836_v26_ = (d_836_v26_).set(d_805_v4_, d_837_v27_)
                    d_838_v28_: C0
                    nw158_ = C0()
                    nw158_.ctor__()
                    d_838_v28_ = nw158_
                    d_839_v29_: D7
                    d_839_v29_ = D7_DC15(d_805_v4_, d_804_v3_, d_804_v3_, d_804_v3_, d_838_v28_)
                    d_840_v30_: D10
                    d_840_v30_ = D10_DC22()
                    d_841_v31_: _dafny.Seq
                    d_841_v31_ = _dafny.SeqWithoutIsStrInference([d_840_v30_])
                    d_842_v32_: _dafny.Map
                    d_842_v32_ = _dafny.Map({d_805_v4_: (0) - (d_804_v3_)})
                    d_843_v33_: D9
                    d_843_v33_ = D9_DC19(d_804_v3_, len(d_842_v32_))
                    d_844_v34_: _dafny.Array
                    nw159_ = _dafny.Array(None, 18)
                    nw159_[int(0)] = d_804_v3_
                    nw159_[int(1)] = (d_839_v29_).cf22
                    nw159_[int(2)] = (697) * (d_804_v3_)
                    nw159_[int(3)] = d_804_v3_
                    nw159_[int(4)] = -819
                    nw159_[int(5)] = d_804_v3_
                    nw159_[int(6)] = d_804_v3_
                    nw159_[int(7)] = (0) - (d_804_v3_)
                    nw159_[int(8)] = d_804_v3_
                    nw159_[int(9)] = -642
                    nw159_[int(10)] = d_804_v3_
                    nw159_[int(11)] = d_804_v3_
                    nw159_[int(12)] = d_804_v3_
                    nw159_[int(13)] = d_804_v3_
                    nw159_[int(14)] = len((_dafny.SeqWithoutIsStrInference([d_840_v30_, d_840_v30_])) + (d_841_v31_))
                    nw159_[int(15)] = default__.safeModuloInt((d_843_v33_).cf29, d_804_v3_)
                    nw159_[int(16)] = (self).fm9(globalState)
                    nw159_[int(17)] = (d_810_v8_)[default__.safeIndex((0) - (default__.fm24(d_804_v3_, True, d_805_v4_, d_804_v3_, globalState)), len(d_810_v8_))]
                    d_844_v34_ = nw159_
                    d_844_v34_ = d_844_v34_
                    d_810_v8_ = _dafny.SeqWithoutIsStrInference([(len(_dafny.Map({d_804_v3_: d_805_v4_}))) - (668) for d_845_i6_ in range(default__.abs(466))])
                    pass
            pass
        d_846_v35_: _dafny.Array
        nw160_ = _dafny.Array(False, 26)
        d_846_v35_ = nw160_
        rhs129_ = (d_809_v7_).cf12
        rhs130_ = d_846_v35_
        r0 = rhs129_
        d_846_v35_ = rhs130_
        d_804_v3_ = default__.safeModuloInt(d_804_v3_, d_804_v3_)
        d_847_v36_: C1
        nw161_ = C1()
        nw161_.ctor__(self.f4, d_804_v3_)
        d_847_v36_ = nw161_
        r0 = d_805_v4_
        d_848_v37_: _dafny.Map
        d_848_v37_ = _dafny.Map({(0) - (-879): self.f4})
        r1 = (d_848_v37_)
        return r0, r1

    def m2(self, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        d_849_v0_: int
        d_849_v0_ = 392
        r0 = d_849_v0_
        d_850_v1_: bool
        d_850_v1_ = False
        d_851_v2_: D10
        d_851_v2_ = D10_DC21(_dafny.MultiSet([d_850_v1_]))
        d_851_v2_ = d_851_v2_
        r2 = not(d_850_v1_)
        d_852_v3_: _dafny.Array
        nw162_ = _dafny.Array(False, 24)
        d_852_v3_ = nw162_
        index173_ = default__.safeIndex(718, (d_852_v3_).length(0))
        (d_852_v3_)[index173_] = d_850_v1_
        index174_ = default__.safeIndex(718, (d_852_v3_).length(0))
        (d_852_v3_)[index174_] = d_850_v1_
        hi6_ = d_849_v0_
        for d_853_i0_ in range(d_849_v0_, hi6_):
            r0 = d_853_i0_
            d_849_v0_ = (d_853_i0_) - (d_853_i0_)
            d_854_v4_: _dafny.Array
            nw163_ = _dafny.Array(None, 2)
            nw163_[int(0)] = (self).f5
            nw163_[int(1)] = (self).f5
            d_854_v4_ = nw163_
            d_855_v5_: D0
            d_855_v5_ = D0_DC0(d_849_v0_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_856_i1_ in range(default__.abs(79))]), d_854_v4_)
            r2 = default__.fm33((d_855_v5_).cf1, globalState)
            if (d_852_v3_)[default__.safeIndex(718, (d_852_v3_).length(0))]:
                d_857_v6_: _dafny.MultiSet
                d_857_v6_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_853_i0_, 901])])
                d_857_v6_ = (d_857_v6_) | ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_849_v0_ for d_858_i2_ in range(default__.abs(776))])])) | (d_857_v6_))
                d_859_v7_: str
                d_859_v7_ = _dafny.CodePoint('t')
                d_860_v8_: _dafny.Seq
                d_860_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cu"))
                d_861_v9_: _dafny.Set
                d_861_v9_ = _dafny.Set({d_860_v8_})
                d_862_v10_: _dafny.Seq
                d_862_v10_ = _dafny.SeqWithoutIsStrInference([((0) - (d_849_v0_)) != (len(d_861_v9_))])
                index175_ = default__.safeIndex(718, (d_852_v3_).length(0))
                rhs131_ = (d_862_v10_)[default__.safeIndex((d_853_i0_) - (d_853_i0_), len(d_862_v10_))]
                rhs132_ = _dafny.CodePoint('f')
                rhs133_ = d_849_v0_
                lhs89_ = d_852_v3_
                lhs90_ = default__.safeIndex(718, (d_852_v3_).length(0))
                lhs89_[lhs90_] = rhs131_
                d_859_v7_ = rhs132_
                r0 = rhs133_
                d_859_v7_ = d_859_v7_
                d_863_v11_: _dafny.Array
                nw164_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_863_v11_ = nw164_
                d_864_v12_: _dafny.Seq
                d_864_v12_ = _dafny.SeqWithoutIsStrInference([d_853_i0_])
                d_865_v13_: _dafny.Map
                d_865_v13_ = _dafny.Map({(d_852_v3_)[default__.safeIndex(718, (d_852_v3_).length(0))]: d_852_v3_})
                d_866_v14_: _dafny.Array
                d_866_v14_ = d_852_v3_
                d_867_v15_: _dafny.Array
                nw165_ = _dafny.Array(None, 3)
                nw165_[int(0)] = d_852_v3_
                nw165_[int(1)] = d_852_v3_
                nw165_[int(2)] = ((d_865_v13_)[d_850_v1_] if (d_850_v1_) in (d_865_v13_) else (d_866_v14_))
                d_867_v15_ = nw165_
                d_868_v16_: _dafny.Map
                d_868_v16_ = _dafny.Map({d_864_v12_: d_867_v15_})
                index176_ = default__.safeIndex(195, (d_863_v11_).length(0))
                (d_863_v11_)[index176_] = ((d_868_v16_)[d_864_v12_] if (d_864_v12_) in (d_868_v16_) else d_867_v15_)
                index177_ = default__.safeIndex(195, (d_863_v11_).length(0))
                (d_863_v11_)[index177_] = d_867_v15_
                d_850_v1_ = (d_852_v3_)[default__.safeIndex(718, (d_852_v3_).length(0))]
            elif True:
                d_849_v0_ = d_853_i0_
                d_869_v17_: D3
                d_869_v17_ = D3_DC8(((0) - (d_849_v0_) if d_850_v1_ else d_853_i0_))
                d_869_v17_ = d_869_v17_
                d_870_v18_: str
                d_870_v18_ = _dafny.CodePoint('w')
                d_870_v18_ = d_870_v18_
                d_871_v19_: _dafny.Seq
                d_871_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjisbuee"))
                d_872_v20_: _dafny.Seq
                d_872_v20_ = _dafny.SeqWithoutIsStrInference([d_871_v19_, d_871_v19_, d_871_v19_])
                d_873_v21_: _dafny.Seq
                d_873_v21_ = d_872_v20_
                d_874_v22_: _dafny.Map
                d_874_v22_ = _dafny.Map({d_849_v0_: len((d_873_v21_))})
                d_874_v22_ = (d_874_v22_).set(d_849_v0_, d_849_v0_)
                index178_ = default__.safeIndex(651, ((self).f5).length(0))
                ((self).f5)[index178_] = d_853_i0_
                index179_ = default__.safeIndex(651, ((self).f5).length(0))
                ((self).f5)[index179_] = d_853_i0_
        d_875_v23_: D4
        d_875_v23_ = D4_DC10(d_850_v1_)
        d_876_v24_: _dafny.Seq
        d_876_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
        d_850_v1_ = (default__.fm33(d_876_v24_, globalState) if (default__.fm16(d_875_v23_, globalState)).issubset(_dafny.Set({d_850_v1_, (d_852_v3_)[default__.safeIndex(718, (d_852_v3_).length(0))]})) else (d_852_v3_)[default__.safeIndex(718, (d_852_v3_).length(0))])
        d_877_v25_: _dafny.Seq
        d_877_v25_ = _dafny.SeqWithoutIsStrInference([d_852_v3_])
        d_878_v27_: str
        d_878_v27_ = _dafny.CodePoint('v')
        d_879_v28_: _dafny.Map
        def iife57_():
            coll27_ = _dafny.Map()
            compr_27_: str
            for compr_27_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), d_878_v27_])).Elements:
                d_880_v26_: str = compr_27_
                if (d_880_v26_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), d_878_v27_])):
                    coll27_[d_880_v26_] = d_849_v0_
            return _dafny.Map(coll27_)
        d_879_v28_ = _dafny.Map({default__.fm33(d_876_v24_, globalState): len(iife57_()
        )})
        r0 = default__.safeDivisionInt(default__.safeModuloInt(len(d_877_v25_), d_849_v0_), ((d_879_v28_)[default__.fm33(d_876_v24_, globalState)] if (default__.fm33(d_876_v24_, globalState)) in (d_879_v28_) else d_849_v0_))
        d_881_v29_: _dafny.Seq
        d_881_v29_ = _dafny.SeqWithoutIsStrInference([(0) - (d_849_v0_), len(d_876_v24_)])
        d_882_v30_: _dafny.Map
        d_882_v30_ = _dafny.Map({D2_DC5(d_881_v29_): 101})
        d_883_v31_: _dafny.Map
        d_883_v31_ = _dafny.Map({(d_882_v30_) | (d_882_v30_): d_850_v1_})
        r1 = d_883_v31_
        r2 = (d_852_v3_)[default__.safeIndex(718, (d_852_v3_).length(0))]
        return r0, r1, r2

    @property
    def f5(self):
        return self._f5

class C7(T0):
    def  __init__(self):
        self._f3: bool = False
        self._f2: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self, f2, f3):
        (self)._f2 = f2
        (self)._f3 = f3

    def fm0(self, p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([(self).f3, ((0) - ((_dafny.MultiSet([(self).f3, (self).f3, (self).f3])).cardinality)) > (-476)])

    def fm6(self, p0, p1, p2, p3, globalState):
        def iife58_():
            coll28_ = _dafny.Map()
            compr_28_: int
            for compr_28_ in _dafny.IntegerRange(747, -512):
                d_884_v0_: int = compr_28_
                if ((747) <= (d_884_v0_)) and ((d_884_v0_) < (-512)):
                    coll28_[(d_884_v0_) + ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtxa")))))] = (self).f3
            return _dafny.Map(coll28_)
        return (0) - ((len((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(iife58_()
        ), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwbgbfuq")))}))]))])): (self).f3})) | (_dafny.Map({512: (self).f3})))) - (default__.safeModuloInt(len(_dafny.Map({(self).f3: not((self).f3)})), len(_dafny.SeqWithoutIsStrInference([-762])))))

    def fm7(self, p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxbvgut"))

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_885_v0_: int
        d_885_v0_ = -934
        d_885_v0_ = d_885_v0_
        d_886_v1_: _dafny.Array
        def lambda28_(d_887_i0_):
            return (self).f3

        init15_ = lambda28_
        nw166_ = _dafny.Array(None, 28)
        for i0_15_ in range(nw166_.length(0)):
            nw166_[i0_15_] = init15_(i0_15_)
        d_886_v1_ = nw166_
        d_888_v2_: str
        d_888_v2_ = _dafny.CodePoint('o')
        d_889_v3_: _dafny.Map
        d_889_v3_ = _dafny.Map({(self).f3: d_888_v2_})
        d_890_v4_: _dafny.MultiSet
        d_890_v4_ = _dafny.MultiSet([d_889_v3_])
        index180_ = default__.safeIndex(462, (d_886_v1_).length(0))
        (d_886_v1_)[index180_] = ((d_890_v4_).set((d_889_v3_).set((self).f3, d_888_v2_), default__.abs(161))).isdisjoint(d_890_v4_)
        d_891_v5_: _dafny.MultiSet
        d_891_v5_ = _dafny.MultiSet([d_885_v0_, d_885_v0_, d_885_v0_])
        index181_ = default__.safeIndex(462, (d_886_v1_).length(0))
        (d_886_v1_)[index181_] = default__.fm8(d_891_v5_, globalState)
        if ((d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))] if (self).f3 else (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]):
            r0 = (self).f3
            d_892_v6_: _dafny.Seq
            d_892_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prubmuoud"))
            d_893_v7_: _dafny.Seq
            d_893_v7_ = _dafny.SeqWithoutIsStrInference([d_885_v0_])
            d_894_v8_: D2
            d_894_v8_ = D2_DC5(d_893_v7_)
            d_895_v9_: D1
            d_895_v9_ = D1_DC4((0) - (d_885_v0_), d_888_v2_, False, len(d_892_v6_))
            d_896_v10_: C2
            nw167_ = C2()
            nw167_.ctor__(d_892_v6_, (self).fm7(d_894_v8_, d_885_v0_, d_894_v8_, d_895_v9_, globalState))
            d_896_v10_ = nw167_
            index182_ = default__.safeIndex(462, (d_886_v1_).length(0))
            (d_886_v1_)[index182_] = (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]
            d_897_v11_: _dafny.Map
            d_897_v11_ = _dafny.Map({d_888_v2_: d_885_v0_})
            d_898_v12_: _dafny.Map
            d_898_v12_ = _dafny.Map({d_888_v2_: len((d_897_v11_) | (d_897_v11_))})
            d_898_v12_ = (d_898_v12_).set(default__.fm30(d_892_v6_, d_885_v0_, len(_dafny.SeqWithoutIsStrInference([d_888_v2_ for d_899_i1_ in range(default__.abs(384))])), globalState), d_885_v0_)
            d_892_v6_ = d_892_v6_
        elif True:
            d_900_v13_: _dafny.Seq
            d_900_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snuoik"))
            d_901_v14_: C5
            nw168_ = C5()
            nw168_.ctor__(d_900_v13_)
            d_901_v14_ = nw168_
            d_885_v0_ = d_885_v0_
            d_902_v15_: D1
            d_902_v15_ = D1_DC2((self).f3)
            d_903_v16_: _dafny.MultiSet
            d_903_v16_ = _dafny.MultiSet([(self).f3, (d_902_v15_).cf4])
            d_904_v17_: C1
            nw169_ = C1()
            nw169_.ctor__(d_903_v16_, d_885_v0_)
            d_904_v17_ = nw169_
            if (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]:
                d_905_v18_: _dafny.Array
                nw170_ = _dafny.Array(int(0), 29)
                d_905_v18_ = nw170_
                index183_ = default__.safeIndex(738, (d_905_v18_).length(0))
                (d_905_v18_)[index183_] = d_904_v17_.f11
                d_906_v19_: _dafny.Seq
                d_906_v19_ = _dafny.SeqWithoutIsStrInference([d_904_v17_.f11])
                index184_ = default__.safeIndex(738, (d_905_v18_).length(0))
                (d_905_v18_)[index184_] = (d_904_v17_.f11 if (len(d_906_v19_)) > (len(_dafny.SeqWithoutIsStrInference([d_888_v2_ for d_907_i2_ in range(default__.abs(-709))]))) else (d_885_v0_) + ((_dafny.MultiSet([d_904_v17_.f11, d_904_v17_.f11, d_885_v0_])).cardinality))
                (d_904_v17_).f11 = (d_904_v17_.f11) + (default__.safeDivisionInt((d_905_v18_)[default__.safeIndex(738, (d_905_v18_).length(0))], d_904_v17_.f11))
                index185_ = default__.safeIndex(715, ((self).f2).length(0))
                ((self).f2)[index185_] = d_905_v18_
                index186_ = default__.safeIndex(715, ((self).f2).length(0))
                ((self).f2)[index186_] = d_905_v18_
                index187_ = default__.safeIndex(462, (d_886_v1_).length(0))
                (d_886_v1_)[index187_] = (((d_901_v14_).f6) + ((d_901_v14_).f6)) <= ((d_901_v14_).f6)
                r0 = (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]
            elif True:
                d_908_v20_: _dafny.Array
                nw171_ = _dafny.Array(_dafny.Seq({}), 7)
                d_908_v20_ = nw171_
                d_909_v21_: _dafny.Seq
                d_909_v21_ = _dafny.SeqWithoutIsStrInference([(d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]])
                index188_ = default__.safeIndex(653, (d_908_v20_).length(0))
                (d_908_v20_)[index188_] = (d_909_v21_).set(default__.safeIndex(d_904_v17_.f11, len(d_909_v21_)), (d_909_v21_)[default__.safeIndex(d_885_v0_, len(d_909_v21_))])
                d_910_v22_: _dafny.Array
                nw172_ = _dafny.Array(None, 8)
                nw172_[int(0)] = (0) - (d_904_v17_.f11)
                nw172_[int(1)] = ((d_891_v5_)[d_904_v17_.f11] if (d_904_v17_.f11) in (d_891_v5_) else len(d_909_v21_))
                nw172_[int(2)] = d_904_v17_.f11
                nw172_[int(3)] = d_904_v17_.f11
                nw172_[int(4)] = d_885_v0_
                nw172_[int(5)] = (d_901_v14_).fm13((d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))], d_904_v17_.f11, (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))], globalState)
                nw172_[int(6)] = d_904_v17_.f11
                nw172_[int(7)] = d_904_v17_.f11
                d_910_v22_ = nw172_
                index189_ = default__.safeIndex(628, ((self).f2).length(0))
                ((self).f2)[index189_] = d_910_v22_
                d_911_v23_: _dafny.Map
                d_911_v23_ = _dafny.Map({len(d_900_v13_): d_903_v16_})
                d_912_v24_: _dafny.Map
                d_912_v24_ = d_911_v23_
                d_913_v25_: _dafny.Map
                d_913_v25_ = _dafny.Map({(self).f3: d_904_v17_.f11})
                index190_ = default__.safeIndex(653, (d_908_v20_).length(0))
                index191_ = default__.safeIndex(628, ((self).f2).length(0))
                rhs134_ = (d_909_v21_) + ((d_909_v21_).set(default__.safeIndex(len(d_913_v25_), len(d_909_v21_)), not((d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))])))
                rhs135_ = (self).f3
                rhs136_ = d_888_v2_
                rhs137_ = d_910_v22_
                rhs138_ = d_912_v24_
                lhs91_ = d_908_v20_
                lhs92_ = default__.safeIndex(653, (d_908_v20_).length(0))
                lhs93_ = (self).f2
                lhs94_ = default__.safeIndex(628, ((self).f2).length(0))
                lhs91_[lhs92_] = rhs134_
                r0 = rhs135_
                d_888_v2_ = rhs136_
                lhs93_[lhs94_] = rhs137_
                d_912_v24_ = rhs138_
                d_914_v26_: _dafny.Map
                d_914_v26_ = _dafny.Map({(d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]: not(not(not((self).f3)))})
                d_914_v26_ = (d_914_v26_).set((self).f3, (self).f3)
                d_900_v13_ = ((d_901_v14_).f6) + (_dafny.SeqWithoutIsStrInference([d_888_v2_ for d_915_i3_ in range(default__.abs(123))]))
                index192_ = default__.safeIndex(462, (d_886_v1_).length(0))
                (d_886_v1_)[index192_] = ((d_914_v26_)[((d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))] if (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))] else (self).f3)] if (((d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))] if (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))] else (self).f3)) in (d_914_v26_) else True)
                index193_ = default__.safeIndex(462, (d_886_v1_).length(0))
                (d_886_v1_)[index193_] = not(False)
            d_916_v27_: _dafny.Seq
            d_916_v27_ = _dafny.SeqWithoutIsStrInference([True])
            r0 = (d_916_v27_)[default__.safeIndex(d_904_v17_.f11, len(d_916_v27_))]
        d_888_v2_ = d_888_v2_
        d_917_v28_: _dafny.Map
        d_917_v28_ = _dafny.Map({d_885_v0_: d_891_v5_})
        d_918_i4_: int
        d_918_i4_ = 0
        with _dafny.label("10"):
            def iife65_():
                coll31_ = _dafny.Map()
                compr_31_: int
                for compr_31_ in _dafny.IntegerRange(213, 546):
                    d_945_v29_: int = compr_31_
                    if ((213) <= (d_945_v29_)) and ((d_945_v29_) < (546)):
                        coll31_[default__.safeDivisionInt(d_945_v29_, d_885_v0_)] = d_885_v0_
                return _dafny.Map(coll31_)
            while (default__.fm31(len(iife65_()
            ), globalState)).ispropersubset(((d_917_v28_)[d_885_v0_] if (d_885_v0_) in (d_917_v28_) else d_891_v5_)):
                with _dafny.c_label("10"):
                    if (d_918_i4_) >= (100):
                        raise _dafny.Break("10")
                    d_918_i4_ = (d_918_i4_) + (1)
                    r0 = (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]
                    if (self).f3:
                        d_885_v0_ = d_885_v0_
                        d_919_v30_: C4
                        nw173_ = C4()
                        nw173_.ctor__((0) - (d_885_v0_))
                        d_919_v30_ = nw173_
                        d_920_v31_: _dafny.Map
                        d_920_v31_ = _dafny.Map({default__.safeDivisionInt(d_919_v30_.f7, 257): (self).f3})
                        d_921_v32_: _dafny.Seq
                        d_921_v32_ = _dafny.SeqWithoutIsStrInference([-445, d_885_v0_])
                        d_920_v31_ = ((d_920_v31_).set(len(d_921_v32_), (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))])) | (d_920_v31_)
                        d_922_v33_: _dafny.Array
                        nw174_ = _dafny.Array(int(0), 16)
                        d_922_v33_ = nw174_
                        d_923_v34_: _dafny.Set
                        d_923_v34_ = _dafny.Set({d_885_v0_})
                        d_924_v35_: D1
                        d_924_v35_ = D1_DC4(len(d_923_v34_), d_888_v2_, not((self).f3), -214)
                        d_925_v36_: _dafny.Seq
                        d_925_v36_ = _dafny.SeqWithoutIsStrInference([d_888_v2_, d_888_v2_, d_888_v2_, (d_924_v35_).cf7, d_888_v2_])
                        d_926_v37_: _dafny.Seq
                        d_926_v37_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f3: (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]})])
                        index194_ = default__.safeIndex(408, (d_922_v33_).length(0))
                        (d_922_v33_)[index194_] = len(((d_925_v36_) + (d_925_v36_)).set(default__.safeIndex(len(d_926_v37_), len((d_925_v36_) + (d_925_v36_))), d_888_v2_))
                        index195_ = default__.safeIndex(408, (d_922_v33_).length(0))
                        (d_922_v33_)[index195_] = d_885_v0_
                        index196_ = default__.safeIndex(408, (d_922_v33_).length(0))
                        index197_ = default__.safeIndex(462, (d_886_v1_).length(0))
                        rhs139_ = default__.safeModuloInt(d_919_v30_.f7, (d_885_v0_) - (d_919_v30_.f7))
                        rhs140_ = True
                        rhs141_ = d_925_v36_
                        rhs142_ = d_885_v0_
                        lhs95_ = d_922_v33_
                        lhs96_ = default__.safeIndex(408, (d_922_v33_).length(0))
                        lhs97_ = d_886_v1_
                        lhs98_ = default__.safeIndex(462, (d_886_v1_).length(0))
                        lhs99_ = d_919_v30_
                        lhs95_[lhs96_] = rhs139_
                        lhs97_[lhs98_] = rhs140_
                        d_925_v36_ = rhs141_
                        lhs99_.f7 = rhs142_
                    elif True:
                        d_888_v2_ = d_888_v2_
                        d_927_v38_: _dafny.Map
                        d_927_v38_ = _dafny.Map({_dafny.MultiSet([d_885_v0_]): (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]})
                        d_928_v39_: _dafny.Seq
                        d_928_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftwmf"))
                        pat_let_tv14_ = d_886_v1_
                        pat_let_tv15_ = d_886_v1_
                        pat_let_tv16_ = d_927_v38_
                        pat_let_tv17_ = globalState
                        d_929_v40_: _dafny.Map
                        def iife59_(_pat_let15_0):
                            def iife60_(d_930_dt__update__tmp_h0_):
                                def iife61_(_pat_let16_0):
                                    def iife62_(d_931_dt__update_hcf14_h0_):
                                        return D3_DC8(d_931_dt__update_hcf14_h0_)
                                    return iife62_(_pat_let16_0)
                                return iife61_(default__.fm19((pat_let_tv15_)[default__.safeIndex(462, (pat_let_tv14_).length(0))], len(pat_let_tv16_), pat_let_tv17_))
                            return iife60_(_pat_let15_0)
                        d_929_v40_ = _dafny.Map({(iife59_(D3_DC8(d_885_v0_)) if (self).f3 else D3_DC8((_dafny.MultiSet([(self).f3])).cardinality)): default__.fm33(d_928_v39_, globalState)})
                        d_932_v42_: _dafny.Map
                        d_932_v42_ = _dafny.Map({(self).f3: False})
                        d_933_v43_: _dafny.Map
                        d_933_v43_ = _dafny.Map({(d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]: len(d_932_v42_)})
                        d_934_v44_: _dafny.Seq
                        d_934_v44_ = _dafny.SeqWithoutIsStrInference([(d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]])
                        d_935_v45_: _dafny.Seq
                        d_935_v45_ = _dafny.SeqWithoutIsStrInference([len(d_934_v44_), d_885_v0_])
                        d_936_v46_: _dafny.Set
                        d_936_v46_ = _dafny.Set({d_933_v43_, d_933_v43_, (d_933_v43_).set((d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))], len(d_935_v45_)), d_933_v43_})
                        d_937_v47_: D1
                        d_937_v47_ = D1_DC3(len(d_936_v46_))
                        d_938_v48_: D3
                        d_938_v48_ = D3_DC8((d_937_v47_).cf5)
                        d_939_v49_: _dafny.Seq
                        d_939_v49_ = _dafny.SeqWithoutIsStrInference([d_938_v48_])
                        def iife63_():
                            coll29_ = _dafny.Map()
                            compr_29_: D3
                            for compr_29_ in (d_939_v49_).Elements:
                                d_940_v41_: D3 = compr_29_
                                if (d_940_v41_) in (d_939_v49_):
                                    coll29_[d_940_v41_] = (self).f3
                            return _dafny.Map(coll29_)
                        d_929_v40_ = iife63_()
                        
                        index198_ = default__.safeIndex(462, (d_886_v1_).length(0))
                        (d_886_v1_)[index198_] = (((d_891_v5_).set(d_885_v0_, default__.abs(469))) - (d_891_v5_)).ispropersubset(d_891_v5_)
                        index199_ = default__.safeIndex(462, (d_886_v1_).length(0))
                        (d_886_v1_)[index199_] = (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]
                        index200_ = default__.safeIndex(462, (d_886_v1_).length(0))
                        (d_886_v1_)[index200_] = ((d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]) and ((self).f3)
                    d_941_v50_: _dafny.Map
                    d_941_v50_ = _dafny.Map({not ((d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]) or ((self).f3): d_886_v1_})
                    d_942_v51_: _dafny.Seq
                    d_942_v51_ = _dafny.SeqWithoutIsStrInference([(d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))], True])
                    index201_ = default__.safeIndex(462, (d_886_v1_).length(0))
                    def iife64_():
                        coll30_ = _dafny.Map()
                        compr_30_: int
                        for compr_30_ in _dafny.IntegerRange(902, 637):
                            d_943_v52_: int = compr_30_
                            if ((902) <= (d_943_v52_)) and ((d_943_v52_) < (637)):
                                coll30_[default__.safeDivisionInt(d_943_v52_, 470)] = False
                        return _dafny.Map(coll30_)
                    rhs143_ = ((d_942_v51_) != (d_942_v51_) if (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))] else not(not((len(_dafny.SeqWithoutIsStrInference([len(iife64_()
                    ), d_885_v0_, 119]))) > (314))))
                    rhs144_ = (d_941_v50_).set((d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))], d_886_v1_)
                    rhs145_ = not ((self).f3) or ((self).f3)
                    rhs146_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))
                    lhs100_ = d_886_v1_
                    lhs101_ = default__.safeIndex(462, (d_886_v1_).length(0))
                    r0 = rhs143_
                    d_941_v50_ = rhs144_
                    lhs100_[lhs101_] = rhs145_
                    d_885_v0_ = rhs146_
                    d_944_v53_: _dafny.Seq
                    d_944_v53_ = _dafny.SeqWithoutIsStrInference([d_885_v0_])
                    index202_ = default__.safeIndex(462, (d_886_v1_).length(0))
                    rhs147_ = (self).f3
                    rhs148_ = d_944_v53_
                    lhs102_ = d_886_v1_
                    lhs103_ = default__.safeIndex(462, (d_886_v1_).length(0))
                    lhs102_[lhs103_] = rhs147_
                    d_944_v53_ = rhs148_
                    pass
            pass
        d_946_v54_: _dafny.Map
        d_946_v54_ = _dafny.Map({d_885_v0_: (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]})
        index203_ = default__.safeIndex(462, (d_886_v1_).length(0))
        (d_886_v1_)[index203_] = ((d_946_v54_)[d_885_v0_] if (d_885_v0_) in (d_946_v54_) else (self).f3)
        r0 = (self).f3
        d_947_v55_: _dafny.MultiSet
        d_947_v55_ = _dafny.MultiSet([(d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))], (d_886_v1_)[default__.safeIndex(462, (d_886_v1_).length(0))]])
        d_948_v56_: _dafny.Map
        d_948_v56_ = _dafny.Map({len(d_946_v54_): d_947_v55_})
        r1 = (d_948_v56_) | (d_948_v56_)
        return r0, r1

    @property
    def f3(self):
        return self._f3
    @property
    def f2(self):
        return self._f2

class C8(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self):
        pass
        pass

    def fm0(self, p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([False, not(not(not(False))), (_dafny.Map({78: True})) != (_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, True]))])])).cardinality: False}))])

    def fm1(self, p0, p1, p2, globalState):
        return (_dafny.MultiSet([(0) - (len(_dafny.Set({False})))])).intersection((_dafny.MultiSet([len(_dafny.Map({True: (_dafny.MultiSet([276])).cardinality})), -163])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-353: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nyftlaqe")))})) for d_949_i0_ in range(default__.abs(136))]))))

    def fm2(self, p0, p1, p2, globalState):
        return not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyqqsd"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtkslws"))))

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_950_v0_: int
        d_950_v0_ = -859
        hi7_ = d_950_v0_
        for d_951_i0_ in range(d_950_v0_, hi7_):
            d_952_v1_: bool
            d_952_v1_ = True
            d_953_v2_: _dafny.Set
            d_953_v2_ = _dafny.Set({d_952_v1_, d_952_v1_})
            d_954_v3_: _dafny.Set
            d_954_v3_ = _dafny.Set({d_953_v2_})
            d_954_v3_ = default__.fm3(globalState)
            d_955_v4_: _dafny.Seq
            d_955_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwkemvrcr"))
            d_956_v5_: _dafny.Seq
            d_956_v5_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_957_i1_ in range(default__.abs(269))]), d_955_v4_])
            d_950_v0_ = default__.safeModuloInt(d_950_v0_, len((d_956_v5_)[default__.safeIndex(d_950_v0_, len(d_956_v5_))]))
            d_958_v6_: _dafny.Array
            nw175_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_958_v6_ = nw175_
            d_959_v7_: D0
            d_959_v7_ = D0_DC0(d_950_v0_, d_955_v4_, d_958_v6_)
            d_960_v8_: D0
            d_960_v8_ = D0_DC1(D0_DC1(d_959_v7_))
            d_961_v9_: D0
            d_961_v9_ = D0_DC1(D0_DC1(d_960_v8_))
            source10_ = d_961_v9_
            if source10_.is_DC0:
                d_962___mcc_h0_ = source10_.cf0
                d_963___mcc_h1_ = source10_.cf1
                d_964___mcc_h2_ = source10_.cf2
                d_965_cf2_ = d_964___mcc_h2_
                d_966_cf1_ = d_963___mcc_h1_
                d_967_cf0_ = d_962___mcc_h0_
                d_968_v10_: _dafny.Map
                d_968_v10_ = _dafny.Map({d_952_v1_: d_955_v4_})
                d_969_v11_: str
                d_969_v11_ = _dafny.CodePoint('s')
                d_970_v12_: D1
                d_970_v12_ = D1_DC4(d_950_v0_, d_969_v11_, d_952_v1_, d_950_v0_)
                d_971_v13_: _dafny.Array
                nw176_ = _dafny.Array(None, 23)
                nw176_[int(0)] = ((d_968_v10_)[not(d_952_v1_)] if (not(d_952_v1_)) in (d_968_v10_) else d_955_v4_)
                nw176_[int(1)] = (d_955_v4_).set(default__.safeIndex(d_950_v0_, len(d_955_v4_)), d_969_v11_)
                nw176_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
                nw176_[int(3)] = d_955_v4_
                nw176_[int(4)] = (_dafny.SeqWithoutIsStrInference([d_969_v11_ for d_972_i2_ in range(default__.abs(561))])) + (d_966_cf1_)
                nw176_[int(5)] = d_955_v4_
                nw176_[int(6)] = ((d_966_cf1_) + (_dafny.SeqWithoutIsStrInference([d_969_v11_ for d_973_i3_ in range(default__.abs(13))]))).set(default__.safeIndex(d_950_v0_, len((d_966_cf1_) + (_dafny.SeqWithoutIsStrInference([d_969_v11_ for d_974_i3_ in range(default__.abs(13))])))), _dafny.CodePoint('c'))
                nw176_[int(7)] = d_955_v4_
                nw176_[int(8)] = d_966_cf1_
                nw176_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "moycxky")))
                nw176_[int(10)] = d_966_cf1_
                nw176_[int(11)] = d_955_v4_
                nw176_[int(12)] = (_dafny.SeqWithoutIsStrInference([d_969_v11_ for d_975_i4_ in range(default__.abs(458))])) + (d_955_v4_)
                nw176_[int(13)] = d_955_v4_
                nw176_[int(14)] = d_966_cf1_
                nw176_[int(15)] = d_966_cf1_
                nw176_[int(16)] = d_966_cf1_
                nw176_[int(17)] = (d_966_cf1_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_950_v0_ for d_976_i5_ in range(default__.abs(53))])), len(d_966_cf1_)), d_969_v11_)
                nw176_[int(18)] = d_966_cf1_
                nw176_[int(19)] = d_966_cf1_
                nw176_[int(20)] = d_955_v4_
                nw176_[int(21)] = (d_966_cf1_ if (d_970_v12_).cf8 else d_966_cf1_)
                nw176_[int(22)] = default__.fm4(d_952_v1_, d_952_v1_, d_952_v1_, globalState)
                d_971_v13_ = nw176_
                d_971_v13_ = d_971_v13_
                d_977_v14_: _dafny.MultiSet
                d_977_v14_ = _dafny.MultiSet([d_950_v0_, d_967_cf0_, len(_dafny.SeqWithoutIsStrInference([d_969_v11_ for d_978_i6_ in range(default__.abs(578))]))])
                d_979_v15_: _dafny.Map
                d_979_v15_ = _dafny.Map({d_967_cf0_: d_977_v14_})
                d_979_v15_ = (d_979_v15_).set(d_967_cf0_, _dafny.MultiSet([d_950_v0_]))
                r0 = (d_952_v1_) == (d_952_v1_)
                d_980_v16_: _dafny.Seq
                d_980_v16_ = _dafny.SeqWithoutIsStrInference([d_952_v1_])
                d_981_v17_: _dafny.Map
                d_981_v17_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_952_v1_, d_952_v1_, d_952_v1_])) + (d_980_v16_): d_980_v16_})
                d_981_v17_ = (d_981_v17_).set(d_980_v16_, d_980_v16_)
            elif True:
                d_982___mcc_h3_ = source10_.cf3
                d_983_cf3_ = d_982___mcc_h3_
                d_984_v18_: _dafny.Array
                nw177_ = _dafny.Array(D1.default()(), 15)
                d_984_v18_ = nw177_
                d_985_v19_: D1
                d_985_v19_ = D1_DC3(d_950_v0_)
                index204_ = default__.safeIndex(450, (d_984_v18_).length(0))
                (d_984_v18_)[index204_] = d_985_v19_
                index205_ = default__.safeIndex(450, (d_984_v18_).length(0))
                (d_984_v18_)[index205_] = d_985_v19_
                d_986_v20_: D1
                d_986_v20_ = D1_DC2(d_952_v1_)
                d_986_v20_ = d_986_v20_
                d_987_v22_: _dafny.Map
                def iife66_():
                    coll32_ = _dafny.Map()
                    compr_32_: _dafny.Seq
                    for compr_32_ in ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abvray"))])).set(d_955_v4_, default__.abs(len(d_955_v4_)))).Elements:
                        d_988_v21_: _dafny.Seq = compr_32_
                        if (d_988_v21_) in ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abvray"))])).set(d_955_v4_, default__.abs(len(d_955_v4_)))):
                            coll32_[d_988_v21_] = (D2_DC5(_dafny.SeqWithoutIsStrInference([d_950_v0_]))).cf10
                    return _dafny.Map(coll32_)
                d_987_v22_ = _dafny.Map({iife66_()
                : d_952_v1_})
                d_987_v22_ = (d_987_v22_).set(default__.fm5(globalState), not((d_952_v1_) or (d_952_v1_)))
                d_989_v23_: _dafny.MultiSet
                d_989_v23_ = _dafny.MultiSet([d_952_v1_])
                d_990_v24_: C1
                nw178_ = C1()
                nw178_.ctor__((d_989_v23_).set(d_952_v1_, default__.abs(d_950_v0_)), d_950_v0_)
                d_990_v24_ = nw178_
            d_991_v25_: T0
            nw179_ = C2()
            nw179_.ctor__(d_955_v4_, d_955_v4_)
            d_991_v25_ = nw179_
            d_992_v26_: _dafny.Map
            d_992_v26_ = _dafny.Map({d_991_v25_: 219})
            d_993_v27_: _dafny.Map
            d_993_v27_ = _dafny.Map({d_952_v1_: d_992_v26_})
            d_992_v26_ = (d_992_v26_ if d_952_v1_ else ((d_993_v27_)[d_952_v1_] if (d_952_v1_) in (d_993_v27_) else d_992_v26_))
        hi8_ = (d_950_v0_) * (d_950_v0_)
        for d_994_i7_ in range((0) - (d_950_v0_), hi8_):
            r0 = (d_994_i7_) < (d_950_v0_)
            d_995_v28_: _dafny.Seq
            d_995_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erb"))
            d_996_v29_: _dafny.Array
            def lambda29_(d_997_v0_):
                def lambda30_(d_998_i8_):
                    return _dafny.SeqWithoutIsStrInference([d_997_v0_, d_997_v0_, len(_dafny.Map({not(True): True}))])

                return lambda30_

            init16_ = lambda29_(d_950_v0_)
            nw180_ = _dafny.Array(None, 10)
            for i0_16_ in range(nw180_.length(0)):
                nw180_[i0_16_] = init16_(i0_16_)
            d_996_v29_ = nw180_
            d_999_v30_: C0
            nw181_ = C0()
            nw181_.ctor__()
            d_999_v30_ = nw181_
            d_1000_v31_: bool
            d_1000_v31_ = False
            d_1001_v32_: _dafny.Map
            d_1001_v32_ = _dafny.Map({d_1000_v31_: d_994_i7_})
            d_1002_v33_: _dafny.Map
            d_1002_v33_ = _dafny.Map({d_1000_v31_: d_1000_v31_})
            rhs149_ = d_995_v28_
            rhs150_ = d_996_v29_
            rhs151_ = ((d_1001_v32_)[not (d_1000_v31_) or (d_1000_v31_)] if (not (d_1000_v31_) or (d_1000_v31_)) in (d_1001_v32_) else (d_994_i7_ if d_1000_v31_ else d_950_v0_))
            rhs152_ = d_999_v30_
            rhs153_ = ((0) - (len((d_1002_v33_) | (d_1002_v33_)))) * (380)
            d_995_v28_ = rhs149_
            d_996_v29_ = rhs150_
            d_950_v0_ = rhs151_
            d_999_v30_ = rhs152_
            d_950_v0_ = rhs153_
            d_1003_v34_: D9
            d_1003_v34_ = D9_DC20(d_1000_v31_, d_1000_v31_, d_1000_v31_)
            pat_let_tv18_ = d_1000_v31_
            d_1004_v35_: _dafny.Seq
            def iife67_(_pat_let17_0):
                def iife68_(d_1005_dt__update__tmp_h0_):
                    def iife69_(_pat_let18_0):
                        def iife70_(d_1006_dt__update_hcf31_h0_):
                            return D9_DC20(d_1006_dt__update_hcf31_h0_, (d_1005_dt__update__tmp_h0_).cf32, (d_1005_dt__update__tmp_h0_).cf33)
                        return iife70_(_pat_let18_0)
                    return iife69_(pat_let_tv18_)
                return iife68_(_pat_let17_0)
            d_1004_v35_ = _dafny.SeqWithoutIsStrInference([iife67_(D9_DC20(d_1000_v31_, d_1000_v31_, d_1000_v31_)), D9_DC20(d_1000_v31_, True, d_1000_v31_), d_1003_v34_, d_1003_v34_, d_1003_v34_])
            d_1007_v36_: D14
            d_1007_v36_ = D14_DC29(d_1004_v35_)
            d_1004_v35_ = (d_1007_v36_).cf46
            d_1000_v31_ = d_1000_v31_
        d_1008_v37_: bool
        d_1008_v37_ = False
        d_1009_v38_: _dafny.Seq
        d_1009_v38_ = _dafny.SeqWithoutIsStrInference([d_1008_v37_])
        d_1010_i9_: int
        d_1010_i9_ = 0
        with _dafny.label("11"):
            while (d_1009_v38_) in (_dafny.SeqWithoutIsStrInference([d_1009_v38_, d_1009_v38_])):
                with _dafny.c_label("11"):
                    if (d_1010_i9_) >= (100):
                        raise _dafny.Break("11")
                    d_1010_i9_ = (d_1010_i9_) + (1)
                    d_1008_v37_ = d_1008_v37_
                    d_1011_v39_: _dafny.Map
                    d_1011_v39_ = _dafny.Map({_dafny.Map({d_950_v0_: d_950_v0_}): (d_950_v0_) <= (d_950_v0_)})
                    d_1012_v40_: _dafny.Map
                    d_1012_v40_ = _dafny.Map({d_950_v0_: d_950_v0_})
                    d_1011_v39_ = (d_1011_v39_).set(d_1012_v40_, d_1008_v37_)
                    d_1013_v41_: _dafny.Array
                    nw182_ = _dafny.Array(None, 16)
                    nw182_[int(0)] = d_1008_v37_
                    nw182_[int(1)] = d_1008_v37_
                    nw182_[int(2)] = d_1008_v37_
                    nw182_[int(3)] = d_1008_v37_
                    nw182_[int(4)] = d_1008_v37_
                    nw182_[int(5)] = d_1008_v37_
                    nw182_[int(6)] = d_1008_v37_
                    nw182_[int(7)] = False
                    nw182_[int(8)] = d_1008_v37_
                    nw182_[int(9)] = d_1008_v37_
                    nw182_[int(10)] = False
                    nw182_[int(11)] = d_1008_v37_
                    nw182_[int(12)] = d_1008_v37_
                    nw182_[int(13)] = d_1008_v37_
                    nw182_[int(14)] = d_1008_v37_
                    nw182_[int(15)] = d_1008_v37_
                    d_1013_v41_ = nw182_
                    index206_ = default__.safeIndex(942, (p0).length(0))
                    (p0)[index206_] = d_1013_v41_
                    index207_ = default__.safeIndex(942, (p0).length(0))
                    (p0)[index207_] = d_1013_v41_
                    d_1014_v42_: _dafny.Map
                    d_1014_v42_ = _dafny.Map({d_1013_v41_: not(not(not (d_1008_v37_) or (d_1008_v37_)))})
                    d_1008_v37_ = ((d_1014_v42_)[d_1013_v41_] if (d_1013_v41_) in (d_1014_v42_) else d_1008_v37_)
                    pass
            pass
        d_1015_v43_: _dafny.Array
        nw183_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
        d_1015_v43_ = nw183_
        d_1016_v44_: _dafny.Seq
        d_1016_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhh"))
        d_1017_v45_: _dafny.Map
        d_1017_v45_ = _dafny.Map({d_1008_v37_: d_1016_v44_})
        index208_ = default__.safeIndex(651, (d_1015_v43_).length(0))
        (d_1015_v43_)[index208_] = ((d_1017_v45_)[d_1008_v37_] if (d_1008_v37_) in (d_1017_v45_) else d_1016_v44_)
        d_1018_v47_: _dafny.Map
        def iife71_():
            coll33_ = _dafny.Map()
            compr_33_: int
            for compr_33_ in _dafny.IntegerRange(445, 767):
                d_1019_v46_: int = compr_33_
                if ((445) <= (d_1019_v46_)) and ((d_1019_v46_) < (767)):
                    coll33_[(d_1019_v46_) + (d_950_v0_)] = d_1008_v37_
            return _dafny.Map(coll33_)
        d_1018_v47_ = _dafny.Map({len(iife71_()
        ): d_1008_v37_})
        index209_ = default__.safeIndex(651, (d_1015_v43_).length(0))
        (d_1015_v43_)[index209_] = (d_1016_v44_) + (default__.fm22(d_1008_v37_, len(d_1018_v47_), globalState))
        d_1008_v37_ = d_1008_v37_
        d_950_v0_ = d_950_v0_
        r0 = d_1008_v37_
        d_1020_v48_: _dafny.Array
        nw184_ = _dafny.Array(int(0), 9)
        d_1020_v48_ = nw184_
        d_1021_v49_: _dafny.Seq
        d_1021_v49_ = _dafny.SeqWithoutIsStrInference([d_1020_v48_])
        d_1022_v50_: _dafny.Seq
        d_1022_v50_ = _dafny.SeqWithoutIsStrInference([(d_1015_v43_)[default__.safeIndex(651, (d_1015_v43_).length(0))]])
        d_1023_v51_: _dafny.Seq
        d_1023_v51_ = _dafny.SeqWithoutIsStrInference([(d_1021_v49_)[default__.safeIndex(len((d_1022_v50_)[default__.safeIndex(d_950_v0_, len(d_1022_v50_))]), len(d_1021_v49_))]])
        d_1024_v52_: _dafny.MultiSet
        d_1024_v52_ = _dafny.MultiSet([True])
        r1 = _dafny.Map({len(d_1023_v51_): d_1024_v52_})
        return r0, r1

    def m1(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        d_1025_v0_: _dafny.Seq
        d_1025_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "goxqvwwrh"))
        d_1025_v0_ = default__.fm4(not (p0) or (p0), p0, p0, globalState)
        d_1026_v1_: int
        d_1026_v1_ = 669
        d_1027_v2_: _dafny.Seq
        d_1027_v2_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1026_v1_ = len(d_1027_v2_)
        d_1028_v3_: _dafny.MultiSet
        d_1028_v3_ = _dafny.MultiSet([d_1026_v1_, -45, d_1026_v1_])
        d_1028_v3_ = d_1028_v3_
        d_1029_v4_: D1
        d_1029_v4_ = D1_DC3(d_1026_v1_)
        d_1030_v5_: D9
        d_1030_v5_ = D9_DC19(d_1026_v1_, d_1026_v1_)
        pat_let_tv19_ = d_1030_v5_
        def iife72_(_pat_let19_0):
            def iife73_(d_1031_dt__update__tmp_h0_):
                def iife74_(_pat_let20_0):
                    def iife75_(d_1032_dt__update_hcf5_h0_):
                        return D1_DC3(d_1032_dt__update_hcf5_h0_)
                    return iife75_(_pat_let20_0)
                return iife74_((pat_let_tv19_).cf29)
            return iife73_(_pat_let19_0)
        source11_ = iife72_(d_1029_v4_)
        if source11_.is_DC3:
            d_1033___mcc_h0_ = source11_.cf5
            d_1034_cf5_ = d_1033___mcc_h0_
            d_1035_v6_: _dafny.Seq
            d_1035_v6_ = _dafny.SeqWithoutIsStrInference([d_1026_v1_, d_1026_v1_, d_1034_cf5_, d_1034_cf5_])
            d_1036_v7_: _dafny.Map
            d_1036_v7_ = _dafny.Map({d_1026_v1_: d_1035_v6_})
            d_1037_v8_: _dafny.Map
            d_1037_v8_ = _dafny.Map({d_1026_v1_: p0})
            d_1038_v9_: D10
            d_1038_v9_ = D10_DC21(_dafny.MultiSet([((d_1037_v8_)[d_1026_v1_] if (d_1026_v1_) in (d_1037_v8_) else False), p0, p0]))
            d_1039_v10_: _dafny.Map
            d_1039_v10_ = _dafny.Map({p0: d_1034_cf5_})
            d_1040_v11_: str
            d_1040_v11_ = _dafny.CodePoint('w')
            d_1041_v12_: _dafny.Array
            nw185_ = _dafny.Array(None, 12)
            nw185_[int(0)] = (0) - (len(((d_1036_v7_)[696] if (696) in (d_1036_v7_) else d_1035_v6_)))
            nw185_[int(1)] = d_1026_v1_
            nw185_[int(2)] = ((d_1038_v9_).cf34).cardinality
            nw185_[int(3)] = d_1034_cf5_
            nw185_[int(4)] = (d_1034_cf5_) - (d_1026_v1_)
            nw185_[int(5)] = len(d_1035_v6_)
            nw185_[int(6)] = ((d_1039_v10_)[p0] if (p0) in (d_1039_v10_) else d_1034_cf5_)
            nw185_[int(7)] = d_1026_v1_
            nw185_[int(8)] = (0) - (len(default__.fm11(d_1040_v11_, d_1034_cf5_, True, default__.fm34(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbdinbg"))), d_1034_cf5_, globalState), globalState)))
            nw185_[int(9)] = default__.fm24(857, True, p0, d_1026_v1_, globalState)
            nw185_[int(10)] = len(d_1035_v6_)
            nw185_[int(11)] = 321
            d_1041_v12_ = nw185_
            d_1041_v12_ = d_1041_v12_
            d_1042_v13_: C1
            nw186_ = C1()
            nw186_.ctor__(_dafny.MultiSet(d_1027_v2_), len(d_1037_v8_))
            d_1042_v13_ = nw186_
            d_1043_v14_: _dafny.Map
            d_1043_v14_ = _dafny.Map({p0: d_1042_v13_})
            d_1044_v15_: _dafny.Map
            d_1044_v15_ = _dafny.Map({p0: p0})
            d_1045_v16_: _dafny.Seq
            d_1045_v16_ = _dafny.SeqWithoutIsStrInference([d_1041_v12_, d_1041_v12_])
            d_1046_v17_: _dafny.Array
            nw187_ = _dafny.Array(None, 19)
            nw187_[int(0)] = d_1041_v12_
            nw187_[int(1)] = d_1041_v12_
            nw187_[int(2)] = d_1041_v12_
            nw187_[int(3)] = d_1041_v12_
            nw187_[int(4)] = d_1041_v12_
            nw187_[int(5)] = (d_1045_v16_)[default__.safeIndex(d_1034_cf5_, len(d_1045_v16_))]
            nw187_[int(6)] = (d_1045_v16_)[default__.safeIndex(d_1026_v1_, len(d_1045_v16_))]
            nw187_[int(7)] = d_1041_v12_
            nw187_[int(8)] = (d_1045_v16_)[default__.safeIndex(704, len(d_1045_v16_))]
            nw187_[int(9)] = d_1041_v12_
            nw187_[int(10)] = d_1041_v12_
            nw187_[int(11)] = d_1041_v12_
            nw187_[int(12)] = d_1041_v12_
            nw187_[int(13)] = d_1041_v12_
            nw187_[int(14)] = d_1041_v12_
            nw187_[int(15)] = d_1041_v12_
            nw187_[int(16)] = d_1041_v12_
            nw187_[int(17)] = d_1041_v12_
            nw187_[int(18)] = d_1041_v12_
            d_1046_v17_ = nw187_
            d_1047_v18_: _dafny.Map
            d_1047_v18_ = d_1039_v10_
            d_1048_v19_: _dafny.Set
            d_1048_v19_ = _dafny.Set({p0})
            d_1049_v20_: D14
            d_1049_v20_ = D14_DC30(d_1026_v1_, d_1046_v17_, d_1047_v18_, len(d_1048_v19_), d_1041_v12_)
            d_1050_v21_: _dafny.Array
            nw188_ = _dafny.Array(None, 11)
            nw188_[int(0)] = (d_1034_cf5_) > (len((d_1043_v14_).set(not(p0), d_1042_v13_)))
            nw188_[int(1)] = (not(((d_1044_v15_)[((d_1044_v15_)[p0] if (p0) in (d_1044_v15_) else p0)] if (((d_1044_v15_)[p0] if (p0) in (d_1044_v15_) else p0)) in (d_1044_v15_) else p0)) if p0 else p0)
            nw188_[int(2)] = (d_1026_v1_) < (d_1026_v1_)
            nw188_[int(3)] = p0
            nw188_[int(4)] = p0
            nw188_[int(5)] = False
            nw188_[int(6)] = ((d_1049_v20_).cf47) not in (d_1035_v6_)
            nw188_[int(7)] = False
            nw188_[int(8)] = p0
            nw188_[int(9)] = p0
            nw188_[int(10)] = False
            d_1050_v21_ = nw188_
            r1 = d_1050_v21_
            d_1035_v6_ = (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1051_i1_ in range(default__.abs(-426))]))) for d_1052_i0_ in range(default__.abs(481))])) + ((d_1035_v6_) + (d_1035_v6_))
            d_1053_v22_: C1
            nw189_ = C1()
            nw189_.ctor__((d_1042_v13_).f10, d_1042_v13_.f11)
            d_1053_v22_ = nw189_
        elif source11_.is_DC4:
            d_1054___mcc_h1_ = source11_.cf6
            d_1055___mcc_h2_ = source11_.cf7
            d_1056___mcc_h3_ = source11_.cf8
            d_1057___mcc_h4_ = source11_.cf9
            d_1058_cf9_ = d_1057___mcc_h4_
            d_1059_cf8_ = d_1056___mcc_h3_
            d_1060_cf7_ = d_1055___mcc_h2_
            d_1061_cf6_ = d_1054___mcc_h1_
            r0 = (p0 if d_1059_cf8_ else (p0) == (d_1059_cf8_))
            if not((d_1026_v1_) < (d_1058_cf9_)):
                d_1062_v23_: _dafny.Array
                nw190_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_1062_v23_ = nw190_
                d_1063_v24_: C7
                nw191_ = C7()
                nw191_.ctor__(d_1062_v23_, p0)
                d_1063_v24_ = nw191_
                d_1064_v25_: _dafny.Map
                d_1064_v25_ = _dafny.Map({True: True})
                d_1065_v26_: _dafny.Set
                d_1065_v26_ = _dafny.Set({(0) - (d_1026_v1_), d_1026_v1_, len(_dafny.SeqWithoutIsStrInference([len(d_1064_v25_), d_1026_v1_])), d_1026_v1_})
                d_1066_v27_: _dafny.Array
                nw192_ = _dafny.Array(None, 14)
                nw192_[int(0)] = len(_dafny.Set({d_1063_v24_}))
                nw192_[int(1)] = d_1026_v1_
                nw192_[int(2)] = d_1058_cf9_
                nw192_[int(3)] = d_1061_cf6_
                nw192_[int(4)] = len(d_1065_v26_)
                nw192_[int(5)] = (0) - ((d_1028_v3_).cardinality)
                nw192_[int(6)] = d_1026_v1_
                nw192_[int(7)] = d_1061_cf6_
                nw192_[int(8)] = len(d_1027_v2_)
                nw192_[int(9)] = (d_1063_v24_).fm6((d_1063_v24_).f3, d_1061_cf6_, d_1026_v1_, d_1060_cf7_, globalState)
                nw192_[int(10)] = len(d_1025_v0_)
                nw192_[int(11)] = 210
                nw192_[int(12)] = d_1058_cf9_
                nw192_[int(13)] = len((d_1064_v25_).set((d_1063_v24_).f3, d_1059_cf8_))
                d_1066_v27_ = nw192_
                d_1067_v28_: _dafny.Array
                def lambda31_(d_1068_v25_):
                    def lambda32_(d_1069_i2_):
                        return d_1068_v25_

                    return lambda32_

                init17_ = lambda31_(d_1064_v25_)
                nw193_ = _dafny.Array(None, 18)
                for i0_17_ in range(nw193_.length(0)):
                    nw193_[i0_17_] = init17_(i0_17_)
                d_1067_v28_ = nw193_
                d_1070_v29_: _dafny.Map
                d_1070_v29_ = _dafny.Map({d_1066_v27_: d_1067_v28_})
                d_1070_v29_ = (d_1070_v29_).set(d_1066_v27_, d_1067_v28_)
                d_1071_v30_: _dafny.Array
                def lambda33_(d_1072_v0_):
                    def lambda34_(d_1073_i3_):
                        return (d_1072_v0_) + (d_1072_v0_)

                    return lambda34_

                init18_ = lambda33_(d_1025_v0_)
                nw194_ = _dafny.Array(None, 11)
                for i0_18_ in range(nw194_.length(0)):
                    nw194_[i0_18_] = init18_(i0_18_)
                d_1071_v30_ = nw194_
                index210_ = default__.safeIndex(429, (d_1071_v30_).length(0))
                (d_1071_v30_)[index210_] = (d_1025_v0_) + (d_1025_v0_)
                index211_ = default__.safeIndex(429, (d_1071_v30_).length(0))
                (d_1071_v30_)[index211_] = ((_dafny.SeqWithoutIsStrInference([d_1060_cf7_ for d_1074_i4_ in range(default__.abs(111))])) + (default__.fm22((d_1063_v24_).f3, d_1061_cf6_, globalState))).set(default__.safeIndex(d_1058_cf9_, len((_dafny.SeqWithoutIsStrInference([d_1060_cf7_ for d_1075_i4_ in range(default__.abs(111))])) + (default__.fm22((d_1063_v24_).f3, d_1061_cf6_, globalState)))), _dafny.CodePoint('w'))
                d_1076_v31_: C5
                nw195_ = C5()
                nw195_.ctor__((d_1071_v30_)[default__.safeIndex(429, (d_1071_v30_).length(0))])
                d_1076_v31_ = nw195_
                d_1025_v0_ = ((d_1076_v31_).f6) + ((d_1071_v30_)[default__.safeIndex(429, (d_1071_v30_).length(0))])
                d_1077_v32_: _dafny.Map
                d_1077_v32_ = _dafny.Map({(d_1058_cf9_) - (d_1026_v1_): False})
                d_1077_v32_ = (d_1077_v32_).set(d_1058_cf9_, d_1059_cf8_)
            elif True:
                d_1078_v33_: C4
                nw196_ = C4()
                nw196_.ctor__(default__.safeModuloInt(len(d_1025_v0_), d_1026_v1_))
                d_1078_v33_ = nw196_
                d_1078_v33_ = d_1078_v33_
                d_1079_v34_: C0
                nw197_ = C0()
                nw197_.ctor__()
                d_1079_v34_ = nw197_
                d_1080_v35_: _dafny.Array
                nw198_ = _dafny.Array(False, 14)
                d_1080_v35_ = nw198_
                d_1081_v36_: _dafny.Set
                d_1081_v36_ = _dafny.Set({d_1080_v35_, d_1080_v35_})
                rhs154_ = d_1059_cf8_
                rhs155_ = d_1081_v36_
                rhs156_ = d_1078_v33_.f7
                rhs157_ = p0
                rhs158_ = p0
                r0 = rhs154_
                d_1081_v36_ = rhs155_
                d_1061_cf6_ = rhs156_
                d_1059_cf8_ = rhs157_
                r0 = rhs158_
                d_1082_v37_: C2
                nw199_ = C2()
                nw199_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm")), ((d_1025_v0_).set(default__.safeIndex(d_1078_v33_.f7, len(d_1025_v0_)), d_1060_cf7_)).set(default__.safeIndex(d_1026_v1_, len((d_1025_v0_).set(default__.safeIndex(d_1078_v33_.f7, len(d_1025_v0_)), d_1060_cf7_))), d_1060_cf7_))
                d_1082_v37_ = nw199_
                d_1083_v38_: _dafny.Map
                d_1083_v38_ = _dafny.Map({d_1026_v1_: d_1082_v37_})
                d_1083_v38_ = (d_1083_v38_).set(d_1026_v1_, d_1082_v37_)
                d_1084_v39_: _dafny.Seq
                d_1084_v39_ = _dafny.SeqWithoutIsStrInference([d_1080_v35_])
                d_1085_v40_: _dafny.MultiSet
                d_1085_v40_ = _dafny.MultiSet([d_1080_v35_, (d_1080_v35_ if not(d_1059_cf8_) else (d_1084_v39_)[default__.safeIndex(d_1026_v1_, len(d_1084_v39_))]), d_1080_v35_, d_1080_v35_, d_1080_v35_])
                d_1058_cf9_ = ((d_1085_v40_)[d_1080_v35_] if (d_1080_v35_) in (d_1085_v40_) else default__.safeModuloInt(d_1026_v1_, 382))
            d_1086_v41_: _dafny.Array
            nw200_ = _dafny.Array(int(0), 7)
            d_1086_v41_ = nw200_
            d_1086_v41_ = d_1086_v41_
            d_1087_v42_: _dafny.Array
            def lambda35_(d_1088_cf9_, d_1089_cf8_, d_1090_cf6_):
                def lambda36_(d_1091_i5_):
                    return (_dafny.Set({_dafny.SeqWithoutIsStrInference([d_1088_cf9_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1089_cf8_]))).cardinality, d_1090_cf6_])})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_1092_i6_ in range(default__.abs(174))])}))

                return lambda36_

            init19_ = lambda35_(d_1058_cf9_, d_1059_cf8_, d_1061_cf6_)
            nw201_ = _dafny.Array(None, 25)
            for i0_19_ in range(nw201_.length(0)):
                nw201_[i0_19_] = init19_(i0_19_)
            d_1087_v42_ = nw201_
            d_1093_v43_: _dafny.Seq
            d_1093_v43_ = _dafny.SeqWithoutIsStrInference([d_1058_cf9_])
            d_1094_v44_: _dafny.Set
            d_1094_v44_ = _dafny.Set({d_1093_v43_})
            index212_ = default__.safeIndex(177, (d_1087_v42_).length(0))
            (d_1087_v42_)[index212_] = d_1094_v44_
            index213_ = default__.safeIndex(177, (d_1087_v42_).length(0))
            def iife76_():
                coll34_ = _dafny.Set()
                compr_34_: _dafny.Seq
                for compr_34_ in ((_dafny.Map({d_1093_v43_: d_1026_v1_})).set(_dafny.SeqWithoutIsStrInference([d_1061_cf6_, d_1058_cf9_]), 437)).keys.Elements:
                    d_1095_v45_: _dafny.Seq = compr_34_
                    if (d_1095_v45_) in ((_dafny.Map({d_1093_v43_: d_1026_v1_})).set(_dafny.SeqWithoutIsStrInference([d_1061_cf6_, d_1058_cf9_]), 437)):
                        coll34_ = coll34_.union(_dafny.Set([d_1095_v45_]))
                return _dafny.Set(coll34_)
            (d_1087_v42_)[index213_] = iife76_()
            
        elif True:
            d_1096___mcc_h5_ = source11_.cf4
            d_1097_cf4_ = d_1096___mcc_h5_
            d_1098_v46_: _dafny.Array
            nw202_ = _dafny.Array(int(0), 11)
            d_1098_v46_ = nw202_
            d_1099_v47_: _dafny.Map
            d_1099_v47_ = _dafny.Map({p0: d_1098_v46_})
            d_1100_v48_: D4
            d_1100_v48_ = D4_DC9(d_1098_v46_)
            d_1101_v49_: D4
            d_1101_v49_ = D4_DC11(d_1100_v48_)
            d_1102_v50_: _dafny.Map
            d_1102_v50_ = _dafny.Map({len(d_1099_v47_): d_1101_v49_})
            d_1102_v50_ = (d_1102_v50_).set((0) - (d_1026_v1_), d_1101_v49_)
            d_1103_v51_: _dafny.Map
            d_1103_v51_ = _dafny.Map({d_1098_v46_: False})
            d_1104_v52_: D9
            d_1104_v52_ = D9_DC20(p0, p0, d_1097_cf4_)
            d_1103_v51_ = (d_1103_v51_).set(d_1098_v46_, (d_1104_v52_).cf33)
            index214_ = default__.safeIndex(212, (d_1098_v46_).length(0))
            (d_1098_v46_)[index214_] = d_1026_v1_
            d_1105_v53_: str
            d_1105_v53_ = _dafny.CodePoint('x')
            d_1106_v54_: _dafny.Set
            d_1106_v54_ = _dafny.Set({d_1105_v53_})
            d_1107_v55_: _dafny.Seq
            d_1107_v55_ = _dafny.SeqWithoutIsStrInference([d_1106_v54_, _dafny.Set({d_1105_v53_}), d_1106_v54_])
            d_1108_v56_: _dafny.Seq
            d_1108_v56_ = _dafny.SeqWithoutIsStrInference([(0) - (d_1026_v1_), d_1026_v1_, (d_1029_v4_).cf5, d_1026_v1_, len(d_1107_v55_)])
            d_1109_v57_: _dafny.Array
            nw203_ = _dafny.Array(False, 18)
            d_1109_v57_ = nw203_
            d_1110_v58_: _dafny.Seq
            d_1110_v58_ = _dafny.SeqWithoutIsStrInference([d_1109_v57_])
            d_1111_v59_: _dafny.Map
            d_1111_v59_ = _dafny.Map({d_1026_v1_: d_1026_v1_})
            d_1112_v60_: _dafny.MultiSet
            d_1112_v60_ = _dafny.MultiSet([default__.fm8(d_1028_v3_, globalState), False])
            d_1113_v61_: C6
            nw204_ = C6()
            nw204_.ctor__(d_1112_v60_, d_1098_v46_)
            d_1113_v61_ = nw204_
            d_1114_v62_: _dafny.Map
            d_1114_v62_ = _dafny.Map({d_1113_v61_: d_1109_v57_})
            d_1115_v63_: _dafny.Set
            d_1115_v63_ = _dafny.Set({d_1109_v57_, d_1109_v57_})
            index215_ = default__.safeIndex(212, (d_1098_v46_).length(0))
            rhs159_ = d_1026_v1_
            rhs160_ = (0) - (default__.safeDivisionInt(d_1026_v1_, (d_1026_v1_) + (d_1026_v1_)))
            rhs161_ = (d_1108_v56_ if d_1097_cf4_ else d_1108_v56_)
            rhs162_ = d_1026_v1_
            rhs163_ = (d_1115_v63_).ispropersubset(_dafny.Set({d_1109_v57_, d_1109_v57_, (d_1110_v58_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference([d_1026_v1_])).set(default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference([d_1111_v59_]))), len(_dafny.SeqWithoutIsStrInference([d_1026_v1_]))), d_1026_v1_)), len(d_1110_v58_))], ((d_1114_v62_)[d_1113_v61_] if (d_1113_v61_) in (d_1114_v62_) else d_1109_v57_)}))
            lhs104_ = d_1098_v46_
            lhs105_ = default__.safeIndex(212, (d_1098_v46_).length(0))
            d_1026_v1_ = rhs159_
            lhs104_[lhs105_] = rhs160_
            d_1108_v56_ = rhs161_
            d_1026_v1_ = rhs162_
            r0 = rhs163_
            d_1116_v64_: _dafny.Map
            d_1116_v64_ = _dafny.Map({(d_1098_v46_)[default__.safeIndex(212, (d_1098_v46_).length(0))]: d_1027_v2_})
            d_1117_v65_: C0
            nw205_ = C0()
            nw205_.ctor__()
            d_1117_v65_ = nw205_
            d_1118_v66_: D7
            d_1118_v66_ = D7_DC15(True, -529, d_1026_v1_, (d_1098_v46_)[default__.safeIndex(212, (d_1098_v46_).length(0))], d_1117_v65_)
            d_1119_v67_: _dafny.MultiSet
            d_1119_v67_ = _dafny.MultiSet([d_1118_v66_, d_1118_v66_])
            d_1120_v68_: _dafny.Map
            d_1120_v68_ = _dafny.Map({d_1119_v67_: d_1025_v0_})
            d_1121_v69_: _dafny.Map
            d_1121_v69_ = _dafny.Map({len(((d_1116_v64_)[(d_1098_v46_)[default__.safeIndex(212, (d_1098_v46_).length(0))]] if ((d_1098_v46_)[default__.safeIndex(212, (d_1098_v46_).length(0))]) in (d_1116_v64_) else (_dafny.SeqWithoutIsStrInference([d_1097_cf4_])).set(default__.safeIndex(len(d_1108_v56_), len(_dafny.SeqWithoutIsStrInference([d_1097_cf4_]))), p0))): (d_1026_v1_) >= (len(d_1120_v68_))})
            d_1122_v70_: T1
            nw206_ = C1()
            nw206_.ctor__(d_1112_v60_, (d_1098_v46_)[default__.safeIndex(212, (d_1098_v46_).length(0))])
            d_1122_v70_ = nw206_
            d_1123_v71_: _dafny.MultiSet
            d_1123_v71_ = _dafny.MultiSet([d_1122_v70_])
            d_1124_v72_: D2
            d_1124_v72_ = D2_DC6(d_1026_v1_, d_1097_cf4_)
            d_1125_v73_: _dafny.Map
            d_1125_v73_ = _dafny.Map({d_1113_v61_: False})
            d_1121_v69_ = (d_1121_v69_).set(((d_1123_v71_).cardinality) + ((d_1028_v3_).cardinality), (d_1113_v61_).fm10(d_1124_v72_, len(d_1025_v0_), (self).fm1(d_1026_v1_, len(_dafny.SeqWithoutIsStrInference([len(d_1125_v73_)])), d_1108_v56_, globalState), globalState))
        d_1126_v74_: _dafny.Seq
        d_1126_v74_ = _dafny.SeqWithoutIsStrInference([len((d_1025_v0_).set(default__.safeIndex((d_1028_v3_).cardinality, len(d_1025_v0_)), default__.fm30(d_1025_v0_, d_1026_v1_, d_1026_v1_, globalState))), d_1026_v1_])
        hi9_ = d_1026_v1_
        for d_1127_i7_ in range(len(d_1126_v74_), hi9_):
            d_1128_v75_: _dafny.Map
            d_1128_v75_ = _dafny.Map({d_1127_i7_: 223})
            d_1128_v75_ = (d_1128_v75_).set(-742, d_1026_v1_)
            d_1129_v76_: _dafny.Set
            d_1129_v76_ = _dafny.Set({d_1026_v1_, d_1127_i7_, len(d_1025_v0_), d_1127_i7_, 924})
            d_1130_v77_: _dafny.Map
            d_1130_v77_ = _dafny.Map({(d_1129_v76_ if p0 else d_1129_v76_): d_1126_v74_})
            d_1131_v79_: _dafny.Map
            d_1131_v79_ = _dafny.Map({p0: d_1126_v74_})
            def iife77_():
                coll35_ = _dafny.Set()
                compr_35_: int
                for compr_35_ in _dafny.IntegerRange(452, -33):
                    d_1132_v78_: int = compr_35_
                    if ((452) <= (d_1132_v78_)) and ((d_1132_v78_) < (-33)):
                        coll35_ = coll35_.union(_dafny.Set([(d_1132_v78_) - (265)]))
                return _dafny.Set(coll35_)
            d_1130_v77_ = (d_1130_v77_).set((d_1129_v76_ if p0 else iife77_()
            ), ((d_1131_v79_)[p0] if (p0) in (d_1131_v79_) else d_1126_v74_))
            d_1026_v1_ = d_1026_v1_
            d_1133_v80_: _dafny.Array
            nw207_ = _dafny.Array(False, 25)
            d_1133_v80_ = nw207_
            d_1134_v81_: _dafny.Seq
            d_1134_v81_ = _dafny.SeqWithoutIsStrInference([d_1133_v80_])
            r2 = (d_1134_v81_) != (d_1134_v81_)
        r0 = default__.fm8(d_1028_v3_, globalState)
        d_1135_v82_: D4
        d_1135_v82_ = D4_DC10(p0)
        r0 = (d_1135_v82_).cf16
        d_1136_v83_: _dafny.Array
        nw208_ = _dafny.Array(False, 9)
        d_1136_v83_ = nw208_
        r1 = d_1136_v83_
        d_1137_v84_: _dafny.Map
        d_1137_v84_ = _dafny.Map({d_1026_v1_: d_1126_v74_})
        r2 = default__.fm8(_dafny.MultiSet(((d_1137_v84_)[d_1026_v1_] if (d_1026_v1_) in (d_1137_v84_) else d_1126_v74_)), globalState)
        return r0, r1, r2

