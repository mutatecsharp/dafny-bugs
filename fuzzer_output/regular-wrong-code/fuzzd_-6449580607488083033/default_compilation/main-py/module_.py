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
    def fm0(p0, globalState):
        return D0_DC1((len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.CodePoint('m'), _dafny.CodePoint('d'), _dafny.CodePoint('v'), _dafny.CodePoint('x'), _dafny.CodePoint('l')})), 216]))) * (906), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_0_i0_ in range(default__.abs(875))]) if not(True) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyfwejrpi"))), True, (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('x'): 168}))]))), True)

    @staticmethod
    def fm1(globalState):
        return (((D14_DC32(True, len(_dafny.Set({-907})), _dafny.Map({True: True}), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s')]), True)).cf55) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "racyt")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irsogxd"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "maxquypl"))))

    @staticmethod
    def fm2(p0, p1, globalState):
        return (len((_dafny.Map({932: 220}) if True else _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgfhyrgb"))): len(_dafny.SeqWithoutIsStrInference([False]))})))) - (417)

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        source0_ = (_dafny.MultiSet([-389]) if not(True) else _dafny.MultiSet([(0) - ((0) - (-550))]))
        d_1___mcc_h0_ = source0_
        d_2_cf18_ = d_1___mcc_h0_
        return (_dafny.SeqWithoutIsStrInference([(d_2_cf18_).cardinality])) + (_dafny.SeqWithoutIsStrInference([67]))

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        if (len(_dafny.Set({True}))) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([879, len(_dafny.Map({_dafny.Set({False}): False}))]))):
            return _dafny.CodePoint('o')
        elif True:
            return _dafny.CodePoint('a')

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        return False

    @staticmethod
    def fm10(p0, p1, globalState):
        return D1_DC3((_dafny.CodePoint('u') if True else _dafny.CodePoint('p')))

    @staticmethod
    def fm12(p0, p1, globalState):
        return ((_dafny.Set({-113})) | (_dafny.Set({len(_dafny.Set({not(True)}))}))).intersection((D6_DC13(7, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwvbfw"))), _dafny.Set({(0) - (-416), 686}))).cf22)

    @staticmethod
    def fm13(p0, globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([not(True), True])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([True]), (_dafny.SeqWithoutIsStrInference([True]))}))) - ((_dafny.Set({_dafny.SeqWithoutIsStrInference([False, False])})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([True, False])})))

    @staticmethod
    def fm21(p0, p1, p2, globalState):
        return (_dafny.MultiSet([len(_dafny.Map({not(True): True}))])) - ((_dafny.MultiSet([-786, 57])))

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        return D1_DC4(not (False) or (False), (_dafny.Set({True, False})) not in (_dafny.SeqWithoutIsStrInference([_dafny.Set({True}), _dafny.Set({False}), _dafny.Set({False})])))

    @staticmethod
    def fm23(globalState):
        return _dafny.SeqWithoutIsStrInference([True, not(True), True, not((-314) != (len(_dafny.Set({653})))), not(((0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True)]))).cardinality)) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (len(_dafny.Map({True: False}))): 125})) for d_3_i0_ in range(default__.abs(597))])))])

    @staticmethod
    def fm24(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(595, 298):
                d_4_v0_: int = compr_0_
                if ((595) <= (d_4_v0_)) and ((d_4_v0_) < (298)):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeModuloInt(d_4_v0_, 642)]))
            return _dafny.Set(coll0_)
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pejhxpg"))), 161])).Elements:
                d_5_v1_: int = compr_1_
                if (d_5_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pejhxpg"))), 161])):
                    coll1_ = coll1_.union(_dafny.Set([(d_5_v1_) * (-835)]))
            return _dafny.Set(coll1_)
        return (iife0_()
        ).intersection((iife1_()
        ) | (_dafny.Set({24, -476, len(_dafny.Map({262: _dafny.CodePoint('v')}))})))

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return _dafny.Map({D0_DC0(True): ((_dafny.MultiSet([(0) - ((_dafny.MultiSet([True, True, True, not(True), False])).cardinality)])) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, True])), 782, -196, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ef")))), 464]))).cardinality})

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        if not (False) or (True):
            return _dafny.MultiSet([_dafny.Map({not(True): _dafny.CodePoint('h')}), _dafny.Map({False: _dafny.CodePoint('u')}), _dafny.Map({True: _dafny.CodePoint('d')}), _dafny.Map({False: _dafny.CodePoint('o')})])
        elif True:
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({True: _dafny.CodePoint('s')})]))).intersection(_dafny.MultiSet([_dafny.Map({True: _dafny.CodePoint('e')})]))

    @staticmethod
    def fm27(p0, globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(not(False))]))) - (_dafny.MultiSet([True, not(False)]))).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        return D0_DC0((_dafny.CodePoint('e')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufhvh"))))

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({True})) - (_dafny.Set({True, False}))).intersection(_dafny.Set({True, True}))

    @staticmethod
    def fm30(globalState):
        return (_dafny.MultiSet([-988, (0) - (len(_dafny.SeqWithoutIsStrInference([632])))]) if False else _dafny.MultiSet([len(_dafny.Map({not(not(True)): len(_dafny.Map({False: len(_dafny.Set({True, not(False)}))}))}))]))

    @staticmethod
    def fm31(globalState):
        return _dafny.Set({_dafny.CodePoint('v')})

    @staticmethod
    def fm32(globalState):
        source1_ = D0_DC2()
        if source1_.is_DC1:
            d_6___mcc_h0_ = source1_.cf1
            d_7___mcc_h1_ = source1_.cf2
            d_8___mcc_h2_ = source1_.cf3
            d_9___mcc_h3_ = source1_.cf4
            d_10___mcc_h4_ = source1_.cf5
            d_11_cf5_ = d_10___mcc_h4_
            d_12_cf4_ = d_9___mcc_h3_
            d_13_cf3_ = d_8___mcc_h2_
            d_14_cf2_ = d_7___mcc_h1_
            d_15_cf1_ = d_6___mcc_h0_
            return D4_DC10(_dafny.CodePoint('m'), _dafny.CodePoint('a'))
        elif source1_.is_DC2:
            if True:
                return D4_DC10(_dafny.CodePoint('s'), _dafny.CodePoint('d'))
            elif True:
                return D4_DC10(_dafny.CodePoint('w'), _dafny.CodePoint('c'))
        elif True:
            d_16___mcc_h5_ = source1_.cf0
            d_17_cf0_ = d_16___mcc_h5_
            return D4_DC10(_dafny.CodePoint('w'), _dafny.CodePoint('f'))

    @staticmethod
    def fm33(p0, p1, globalState):
        return (_dafny.Map({not(not(False)): _dafny.Map({_dafny.SeqWithoutIsStrInference([True, True]): not(False)})})) | (_dafny.Map({not(True): _dafny.Map({_dafny.SeqWithoutIsStrInference([True, True, False]): True})}))

    @staticmethod
    def fm34(p0, p1, globalState):
        if (not(not(True))) == (False):
            return D11_DC25(_dafny.Map({True: True}))
        elif True:
            return D11_DC25(_dafny.Map({not(True): not(False)}))

    @staticmethod
    def fm35(p0, p1, globalState):
        return _dafny.MultiSet([_dafny.CodePoint('l')])

    @staticmethod
    def fm36(p0, p1, p2, p3, globalState):
        return _dafny.Map({False: (_dafny.Map({True: 138})) not in (_dafny.MultiSet([_dafny.Map({True: (0) - (len(_dafny.Set({False})))}), _dafny.Map({True: 687}), _dafny.Map({True: len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([566])}))}), _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality]))}), _dafny.Map({False: 459})]))})

    @staticmethod
    def fm37(p0, globalState):
        return _dafny.Map({(_dafny.MultiSet([True])) | (_dafny.MultiSet([not(False), True])): (D4_DC10(_dafny.CodePoint('h'), _dafny.CodePoint('v'))).cf16})

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        return D2_DC6(_dafny.Set({len(_dafny.Map({297: D14_DC32(True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjwmcxnlo"))), _dafny.Map({True: True}), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t'), _dafny.CodePoint('f')]), False)}))}))

    @staticmethod
    def fm39(p0, p1, globalState):
        return D4_DC9(_dafny.Set({True}))

    @staticmethod
    def fm40(p0, globalState):
        return D14_DC32((_dafny.Set({_dafny.CodePoint('n'), _dafny.CodePoint('f')})).ispropersubset((D22_DC50(_dafny.Set({_dafny.CodePoint('w')}))).cf80), (290) + (-436), (_dafny.Map({not(False): False})) | (_dafny.Map({False: False})), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), _dafny.CodePoint('b')])), True)

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        return D6_DC14(((_dafny.MultiSet([(0) - ((0) - (len(_dafny.Map({len(_dafny.Map({715: 609})): len(_dafny.Map({330: not(True)}))}))))])).intersection(_dafny.MultiSet([260, -500]))).cardinality, 586, (_dafny.MultiSet([not(False)])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))))

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([848]))) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))]))).Elements:
                d_18_v0_: int = compr_2_
                if (d_18_v0_) in ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([848]))) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))]))):
                    coll2_[default__.safeModuloInt(d_18_v0_, -263)] = True
            return _dafny.Map(coll2_)
        return iife2_()
        

    @staticmethod
    def fm43(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(-464, 557):
                d_19_v0_: int = compr_3_
                if ((-464) <= (d_19_v0_)) and ((d_19_v0_) < (557)):
                    coll3_[default__.safeModuloInt(d_19_v0_, (0) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, True])).cardinality, len(_dafny.Set({False, True}))]))))] = 200
            return _dafny.Map(coll3_)
        return iife3_()
        

    @staticmethod
    def fm44(globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])) for d_24_i4_ in range(default__.abs(-771))])).Elements:
                d_25_v0_: int = compr_4_
                if (d_25_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])) for d_24_i4_ in range(default__.abs(-771))])):
                    coll4_ = coll4_.union(_dafny.Set([(d_25_v0_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-688])) for d_26_i5_ in range(default__.abs(167))]))))]))
            return _dafny.Set(coll4_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(False): 328})) for d_20_i1_ in range(default__.abs(94))])) for d_21_i0_ in range(default__.abs(282))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})) for d_22_i2_ in range(default__.abs(28))])), _dafny.MultiSet([647, -625]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('k'): True}))])), _dafny.MultiSet([645, 732, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_23_i3_ in range(default__.abs(697))]))), len(iife4_()
        )]), _dafny.MultiSet([687, len(_dafny.SeqWithoutIsStrInference([True])), -742])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([586])))]))])])))

    @staticmethod
    def m0(p0, globalState):
        r0: int = int(0)
        d_27_v0_: int
        d_27_v0_ = 629
        d_28_i0_: int
        d_28_i0_ = 0
        with _dafny.label("0"):
            while (d_27_v0_) < (d_27_v0_):
                with _dafny.c_label("0"):
                    if (d_28_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_28_i0_ = (d_28_i0_) + (1)
                    if p0:
                        d_29_v1_: C1
                        nw0_ = C1()
                        nw0_.ctor__(p0)
                        d_29_v1_ = nw0_
                        (globalState).f2 = d_27_v0_
                        d_27_v0_ = (0) - (d_27_v0_)
                        d_30_v2_: D0
                        d_30_v2_ = D0_DC0((d_29_v1_).f32)
                        d_31_v3_: C2
                        nw1_ = C2()
                        nw1_.ctor__(d_27_v0_, p0, (0) - (-851))
                        d_31_v3_ = nw1_
                        d_32_v4_: _dafny.Map
                        d_32_v4_ = _dafny.Map({d_30_v2_: len(_dafny.SeqWithoutIsStrInference([d_31_v3_]))})
                        d_33_v5_: D10
                        d_33_v5_ = D10_DC23()
                        d_34_v6_: _dafny.Seq
                        d_34_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehxqgq"))
                        d_35_v7_: _dafny.MultiSet
                        d_35_v7_ = _dafny.MultiSet([(d_31_v3_).f31, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_36_i1_ in range(default__.abs(614))])), len(d_34_v6_), 411])
                        d_37_v8_: _dafny.Map
                        d_37_v8_ = _dafny.Map({not((d_29_v1_).f32): (d_29_v1_).f32})
                        d_38_v9_: _dafny.Seq
                        d_38_v9_ = _dafny.SeqWithoutIsStrInference([(d_29_v1_).f32, p0, p0])
                        d_39_v10_: _dafny.Array
                        nw2_ = _dafny.Array(None, 26)
                        nw2_[int(0)] = default__.fm5((d_29_v1_).f32, d_32_v4_, (d_31_v3_).f31, globalState)
                        nw2_[int(1)] = p0
                        nw2_[int(2)] = (-184) >= ((d_31_v3_).f31)
                        nw2_[int(3)] = p0
                        nw2_[int(4)] = p0
                        nw2_[int(5)] = (d_29_v1_).f32
                        nw2_[int(6)] = (d_29_v1_).f32
                        nw2_[int(7)] = p0
                        nw2_[int(8)] = p0
                        nw2_[int(9)] = (len(_dafny.Set({d_33_v5_, d_33_v5_}))) <= (d_27_v0_)
                        nw2_[int(10)] = (d_27_v0_) < (362)
                        nw2_[int(11)] = (d_29_v1_).f32
                        nw2_[int(12)] = (d_27_v0_) > (d_27_v0_)
                        nw2_[int(13)] = p0
                        nw2_[int(14)] = (d_29_v1_).f32
                        nw2_[int(15)] = (d_35_v7_) == (d_35_v7_)
                        nw2_[int(16)] = (d_29_v1_).f32
                        nw2_[int(17)] = (len(d_34_v6_)) == (144)
                        nw2_[int(18)] = (d_37_v8_) != ((d_37_v8_).set(p0, (d_38_v9_)[default__.safeIndex(d_27_v0_, len(d_38_v9_))]))
                        nw2_[int(19)] = (len(d_34_v6_)) < ((d_31_v3_).f31)
                        nw2_[int(20)] = p0
                        nw2_[int(21)] = (d_29_v1_).f32
                        nw2_[int(22)] = p0
                        nw2_[int(23)] = (d_29_v1_).f32
                        nw2_[int(24)] = not (p0) or ((d_29_v1_).f32)
                        nw2_[int(25)] = not(((d_31_v3_).f31) <= ((d_31_v3_).f31))
                        d_39_v10_ = nw2_
                        rhs0_ = d_39_v10_
                        rhs1_ = not((d_38_v9_)[default__.safeIndex((d_31_v3_).f31, len(d_38_v9_))])
                        lhs0_ = globalState
                        d_39_v10_ = rhs0_
                        lhs0_.f1 = rhs1_
                        d_40_v11_: _dafny.MultiSet
                        d_41_v12_: _dafny.Map
                        out0_: _dafny.MultiSet
                        out1_: _dafny.Map
                        out0_, out1_ = (d_29_v1_).m7((d_31_v3_).fm7(globalState), d_30_v2_, d_34_v6_, globalState)
                        d_40_v11_ = out0_
                        d_41_v12_ = out1_
                    elif True:
                        d_42_v13_: C2
                        nw3_ = C2()
                        nw3_.ctor__(d_27_v0_, not(p0), default__.safeModuloInt(d_27_v0_, d_27_v0_))
                        d_42_v13_ = nw3_
                        (globalState).f0 = (0) - (d_27_v0_)
                        d_43_v14_: _dafny.Seq
                        d_43_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "myjeh"))
                        (globalState).f1 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iplftk"))) + (d_43_v14_)) <= ((d_43_v14_) + (d_43_v14_))
                        d_44_v15_: _dafny.Map
                        d_44_v15_ = _dafny.Map({d_27_v0_: p0})
                        d_45_v16_: _dafny.Map
                        d_45_v16_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfaji"))): d_27_v0_})
                        (globalState).f1 = (default__.fm43(len(d_44_v15_), p0, (_dafny.MultiSet([d_27_v0_])).cardinality, (0) - ((d_42_v13_).f31), globalState)) not in (_dafny.SeqWithoutIsStrInference([d_45_v16_]))
                        d_46_v17_: _dafny.Map
                        d_46_v17_ = _dafny.Map({p0: p0})
                        d_47_v18_: _dafny.Map
                        d_47_v18_ = _dafny.Map({d_27_v0_: d_46_v17_})
                        d_48_v19_: _dafny.Seq
                        d_48_v19_ = _dafny.SeqWithoutIsStrInference([75, d_27_v0_, (d_42_v13_).f31])
                        d_49_v20_: _dafny.Map
                        d_49_v20_ = _dafny.Map({((d_47_v18_)[len(d_48_v19_)] if (len(d_48_v19_)) in (d_47_v18_) else d_46_v17_): (d_42_v13_).f31})
                        d_49_v20_ = (d_49_v20_).set(d_46_v17_, d_27_v0_)
                    d_50_v21_: D0
                    d_50_v21_ = D0_DC0(p0)
                    pat_let_tv0_ = p0
                    d_51_v22_: _dafny.Map
                    def iife5_(_pat_let0_0):
                        def iife6_(d_52_dt__update__tmp_h0_):
                            def iife7_(_pat_let1_0):
                                def iife8_(d_53_dt__update_hcf0_h0_):
                                    return D0_DC0(d_53_dt__update_hcf0_h0_)
                                return iife8_(_pat_let1_0)
                            return iife7_(pat_let_tv0_)
                        return iife6_(_pat_let0_0)
                    d_51_v22_ = _dafny.Map({iife5_(d_50_v21_): d_27_v0_})
                    d_54_v23_: _dafny.Set
                    d_54_v23_ = _dafny.Set({p0, False})
                    d_55_v24_: C2
                    nw4_ = C2()
                    nw4_.ctor__((0) - ((d_27_v0_) - (d_27_v0_)), default__.fm5(p0, d_51_v22_, d_27_v0_, globalState), len(d_54_v23_))
                    d_55_v24_ = nw4_
                    d_56_v25_: _dafny.Seq
                    d_56_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ww"))
                    d_57_v26_: _dafny.Seq
                    d_57_v26_ = _dafny.SeqWithoutIsStrInference([(0) - (((d_55_v24_).f31) - (len(d_56_v25_))), d_27_v0_, default__.safeModuloInt((d_55_v24_).f31, (d_55_v24_).f31)])
                    (globalState).f0 = (0) - (len(d_57_v26_))
                    d_58_v27_: _dafny.MultiSet
                    d_58_v27_ = _dafny.MultiSet([p0])
                    d_58_v27_ = d_58_v27_
                    pass
            pass
        d_59_v28_: _dafny.MultiSet
        d_59_v28_ = _dafny.MultiSet([d_27_v0_, d_27_v0_, default__.fm2(d_27_v0_, d_27_v0_, globalState)])
        d_60_v29_: _dafny.Seq
        d_60_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfda"))
        d_61_v30_: _dafny.Array
        nw5_ = _dafny.Array(None, 11)
        nw5_[int(0)] = d_27_v0_
        nw5_[int(1)] = default__.safeDivisionInt(d_27_v0_, d_27_v0_)
        nw5_[int(2)] = d_27_v0_
        nw5_[int(3)] = (((d_59_v28_)[d_27_v0_] if (d_27_v0_) in (d_59_v28_) else d_27_v0_)) + (d_27_v0_)
        nw5_[int(4)] = d_27_v0_
        nw5_[int(5)] = default__.safeModuloInt(d_27_v0_, d_27_v0_)
        nw5_[int(6)] = (_dafny.MultiSet(default__.fm44(globalState))).cardinality
        nw5_[int(7)] = d_27_v0_
        nw5_[int(8)] = d_27_v0_
        nw5_[int(9)] = len(d_60_v29_)
        nw5_[int(10)] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ewsm"))), d_27_v0_)
        d_61_v30_ = nw5_
        index0_ = default__.safeIndex(194, (d_61_v30_).length(0))
        (d_61_v30_)[index0_] = d_27_v0_
        index1_ = default__.safeIndex(194, (d_61_v30_).length(0))
        (d_61_v30_)[index1_] = d_27_v0_
        (globalState).f0 = d_27_v0_
        d_60_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fba"))
        (globalState).f0 = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_62_i2_ in range(default__.abs(324))]))
        if not(not (p0) or (not(p0))):
            (globalState).f0 = ((d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))]) + (default__.safeDivisionInt(d_27_v0_, d_27_v0_))
            d_63_v31_: _dafny.Map
            d_63_v31_ = _dafny.Map({d_60_v29_: (0) - (d_27_v0_)})
            d_63_v31_ = (d_63_v31_).set(d_60_v29_, 970)
            (globalState).f2 = (d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))]
            d_64_v32_: _dafny.Map
            d_64_v32_ = _dafny.Map({d_27_v0_: p0})
            d_64_v32_ = (d_64_v32_).set(len(_dafny.Set({p0, p0})), p0)
            d_65_v33_: D0
            d_65_v33_ = D0_DC0(True)
            d_66_v34_: D0
            d_66_v34_ = D0_DC0(default__.fm5(False, _dafny.Map({d_65_v33_: len(d_60_v29_)}), d_27_v0_, globalState))
            d_67_v35_: _dafny.Map
            d_67_v35_ = _dafny.Map({d_65_v33_: (d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))]})
            d_68_v36_: C3
            nw6_ = C3()
            nw6_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nj")), not(p0), p0, (d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))])
            d_68_v36_ = nw6_
            d_69_v37_: _dafny.Set
            d_69_v37_ = _dafny.Set({d_68_v36_})
            d_70_v38_: _dafny.Seq
            d_70_v38_ = _dafny.SeqWithoutIsStrInference([False])
            d_71_v39_: _dafny.Array
            nw7_ = _dafny.Array(None, 23)
            nw7_[int(0)] = (d_68_v36_).f30
            nw7_[int(1)] = (d_70_v38_)[default__.safeIndex((d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))], len(d_70_v38_))]
            nw7_[int(2)] = p0
            nw7_[int(3)] = (d_68_v36_).f30
            nw7_[int(4)] = (d_68_v36_).f30
            nw7_[int(5)] = (d_68_v36_).f30
            nw7_[int(6)] = p0
            nw7_[int(7)] = p0
            nw7_[int(8)] = (d_68_v36_).f30
            nw7_[int(9)] = p0
            nw7_[int(10)] = False
            nw7_[int(11)] = (d_68_v36_).f30
            nw7_[int(12)] = p0
            nw7_[int(13)] = (d_68_v36_).f30
            nw7_[int(14)] = not((d_68_v36_).f30)
            nw7_[int(15)] = True
            nw7_[int(16)] = (d_68_v36_).f30
            nw7_[int(17)] = not(p0)
            nw7_[int(18)] = p0
            nw7_[int(19)] = (d_68_v36_).f30
            nw7_[int(20)] = p0
            nw7_[int(21)] = p0
            nw7_[int(22)] = p0
            d_71_v39_ = nw7_
            d_72_v40_: C0
            nw8_ = C0()
            nw8_.ctor__(d_60_v29_, default__.fm5(not(p0), d_67_v35_, (d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))], globalState), len(d_69_v37_), d_71_v39_, (d_68_v36_).f30)
            d_72_v40_ = nw8_
            d_73_v41_: D17
            d_73_v41_ = D17_DC39(D17_DC38())
            d_74_v42_: C2
            nw9_ = C2()
            nw9_.ctor__((d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))], False, len(_dafny.Map({default__.fm5(p0, d_67_v35_, len(d_70_v38_), globalState): (d_68_v36_).f30})))
            d_74_v42_ = nw9_
            d_75_v43_: D17
            d_75_v43_ = D17_DC37(d_74_v42_)
            d_76_v44_: _dafny.Seq
            d_76_v44_ = _dafny.SeqWithoutIsStrInference([(d_73_v41_).cf64, d_75_v43_, d_75_v43_, d_75_v43_])
            d_77_v45_: _dafny.Seq
            d_77_v45_ = _dafny.SeqWithoutIsStrInference([d_72_v40_.f35])
            pat_let_tv1_ = d_76_v44_
            pat_let_tv2_ = d_61_v30_
            pat_let_tv3_ = d_61_v30_
            pat_let_tv4_ = d_76_v44_
            def iife9_(_pat_let2_0):
                def iife10_(d_78_dt__update__tmp_h1_):
                    def iife11_(_pat_let3_0):
                        def iife12_(d_79_dt__update_hcf64_h0_):
                            return D17_DC39(d_79_dt__update_hcf64_h0_)
                        return iife12_(_pat_let3_0)
                    return iife11_((pat_let_tv1_)[default__.safeIndex((pat_let_tv3_)[default__.safeIndex(194, (pat_let_tv2_).length(0))], len(pat_let_tv4_))])
                return iife10_(_pat_let2_0)
            rhs2_ = d_72_v40_
            rhs3_ = iife9_(d_73_v41_)
            rhs4_ = ((d_70_v38_) + (d_70_v38_)) + (d_70_v38_)
            rhs5_ = (d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))]
            rhs6_ = ((0) - (default__.safeModuloInt((d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))], d_27_v0_))) - ((len(d_77_v45_)) * ((d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))]))
            lhs1_ = globalState
            lhs2_ = globalState
            d_72_v40_ = rhs2_
            d_73_v41_ = rhs3_
            d_70_v38_ = rhs4_
            lhs1_.f2 = rhs5_
            lhs2_.f0 = rhs6_
        elif True:
            d_80_v46_: str
            d_80_v46_ = _dafny.CodePoint('r')
            d_80_v46_ = d_80_v46_
            d_81_v47_: _dafny.Array
            nw10_ = _dafny.Array(None, 3)
            d_81_v47_ = nw10_
            d_81_v47_ = d_81_v47_
            d_82_v48_: _dafny.Seq
            d_82_v48_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_60_v29_)[default__.safeIndex((d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))], len(d_60_v29_))] for d_83_i3_ in range(default__.abs(794))]), d_60_v29_])
            d_84_v49_: _dafny.Array
            nw11_ = _dafny.Array(None, 12)
            d_84_v49_ = nw11_
            d_85_v50_: _dafny.Map
            d_85_v50_ = _dafny.Map({d_82_v48_: d_84_v49_})
            d_86_v51_: D21
            d_86_v51_ = D21_DC48(d_84_v49_)
            d_85_v50_ = (d_85_v50_).set((d_82_v48_) + (d_82_v48_), (d_86_v51_).cf77)
            d_87_v52_: _dafny.Array
            def lambda0_(d_88_p0_):
                def lambda1_(d_89_i4_):
                    return d_88_p0_

                return lambda1_

            init0_ = lambda0_(p0)
            nw12_ = _dafny.Array(None, 29)
            for i0_0_ in range(nw12_.length(0)):
                nw12_[i0_0_] = init0_(i0_0_)
            d_87_v52_ = nw12_
            index2_ = default__.safeIndex(70, (d_87_v52_).length(0))
            (d_87_v52_)[index2_] = (d_59_v28_).issubset(_dafny.MultiSet([(0) - ((d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))])]))
            index3_ = default__.safeIndex(70, (d_87_v52_).length(0))
            (d_87_v52_)[index3_] = p0
            d_90_v53_: C1
            nw13_ = C1()
            nw13_.ctor__((d_87_v52_)[default__.safeIndex(70, (d_87_v52_).length(0))])
            d_90_v53_ = nw13_
            d_91_v54_: D10
            d_91_v54_ = D10_DC24(d_27_v0_, d_90_v53_, d_80_v46_)
            d_92_v55_: D10
            d_92_v55_ = D10_DC24((d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))], (d_91_v54_).cf37, d_80_v46_)
            pat_let_tv5_ = d_90_v53_
            pat_let_tv6_ = d_80_v46_
            def iife13_(_pat_let4_0):
                def iife14_(d_93_dt__update__tmp_h2_):
                    def iife15_(_pat_let5_0):
                        def iife16_(d_94_dt__update_hcf37_h0_):
                            def iife17_(_pat_let6_0):
                                def iife18_(d_95_dt__update_hcf38_h0_):
                                    return D10_DC24((d_93_dt__update__tmp_h2_).cf36, d_94_dt__update_hcf37_h0_, d_95_dt__update_hcf38_h0_)
                                return iife18_(_pat_let6_0)
                            return iife17_(pat_let_tv6_)
                        return iife16_(_pat_let5_0)
                    return iife15_(pat_let_tv5_)
                return iife14_(_pat_let4_0)
            d_92_v55_ = iife13_(d_92_v55_)
        r0 = (d_61_v30_)[default__.safeIndex(194, (d_61_v30_).length(0))]
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_96_v0_: _dafny.Array
        nw14_ = _dafny.Array(_dafny.Set({}), 24)
        d_96_v0_ = nw14_
        d_97_v1_: _dafny.Seq
        d_97_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oduxsblwp"))
        d_98_v2_: bool
        d_98_v2_ = False
        d_99_v3_: int
        d_99_v3_ = 189
        d_100_v4_: _dafny.Map
        d_100_v4_ = _dafny.Map({d_98_v2_: d_99_v3_})
        d_101_v5_: _dafny.Seq
        d_101_v5_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_98_v2_: d_99_v3_}), d_100_v4_])
        d_102_v6_: _dafny.Set
        d_102_v6_ = _dafny.Set({d_99_v3_})
        d_103_globalState_: GlobalState
        nw15_ = GlobalState()
        nw15_.ctor__(281, False, -82, d_96_v0_, True, -974, 777, False, (d_97_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))), d_101_v5_, -345, True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrxdsa")), _dafny.Map({d_98_v2_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxiyve"))}), 362, 72, 327, (d_102_v6_) | (d_102_v6_), True, -736, True, True, False, False, True)
        d_103_globalState_ = nw15_
        d_98_v2_ = d_98_v2_
        d_104_v7_: _dafny.Array
        def lambda2_(d_105_v2_, d_106_v1_):
            def lambda3_(d_107_i0_):
                return not((d_105_v2_) and ((D0_DC1(-205, d_106_v1_, d_105_v2_, len(d_106_v1_), d_105_v2_)).cf5))

            return lambda3_

        init1_ = lambda2_(d_98_v2_, d_97_v1_)
        nw16_ = _dafny.Array(None, 26)
        for i0_1_ in range(nw16_.length(0)):
            nw16_[i0_1_] = init1_(i0_1_)
        d_104_v7_ = nw16_
        index4_ = default__.safeIndex(810, (d_104_v7_).length(0))
        (d_104_v7_)[index4_] = not (d_98_v2_) or (d_98_v2_)
        index5_ = default__.safeIndex(810, (d_104_v7_).length(0))
        (d_104_v7_)[index5_] = (d_99_v3_) <= (d_99_v3_)
        d_108_i1_: int
        d_108_i1_ = 0
        with _dafny.label("1"):
            while not ((len(d_100_v4_)) <= (d_99_v3_)) or (not(False)):
                with _dafny.c_label("1"):
                    if (d_108_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_108_i1_ = (d_108_i1_) + (1)
                    (d_103_globalState_).f0 = d_99_v3_
                    d_109_v8_: _dafny.MultiSet
                    d_109_v8_ = _dafny.MultiSet([d_98_v2_])
                    d_110_v9_: str
                    d_110_v9_ = _dafny.CodePoint('w')
                    d_111_v10_: _dafny.MultiSet
                    d_111_v10_ = _dafny.MultiSet([d_110_v9_])
                    d_112_v11_: D0
                    d_112_v11_ = D0_DC1(315, default__.fm1(d_103_globalState_), (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], default__.fm2(((d_109_v8_)[d_98_v2_] if (d_98_v2_) in (d_109_v8_) else (d_111_v10_).cardinality), d_99_v3_, d_103_globalState_), d_98_v2_)
                    d_113_v12_: _dafny.Seq
                    d_113_v12_ = _dafny.SeqWithoutIsStrInference([d_99_v3_])
                    d_114_v13_: _dafny.Map
                    d_114_v13_ = _dafny.Map({len(d_113_v12_): d_99_v3_})
                    d_115_v14_: _dafny.Array
                    nw17_ = _dafny.Array(None, 9)
                    nw17_[int(0)] = default__.fm0(d_99_v3_, d_103_globalState_)
                    nw17_[int(1)] = d_112_v11_
                    nw17_[int(2)] = d_112_v11_
                    nw17_[int(3)] = d_112_v11_
                    nw17_[int(4)] = default__.fm0(d_99_v3_, d_103_globalState_)
                    nw17_[int(5)] = d_112_v11_
                    nw17_[int(6)] = d_112_v11_
                    nw17_[int(7)] = d_112_v11_
                    nw17_[int(8)] = D0_DC1(((d_114_v13_)[d_99_v3_] if (d_99_v3_) in (d_114_v13_) else d_99_v3_), d_97_v1_, False, d_99_v3_, d_98_v2_)
                    d_115_v14_ = nw17_
                    index6_ = default__.safeIndex(295, (d_115_v14_).length(0))
                    (d_115_v14_)[index6_] = d_112_v11_
                    index7_ = default__.safeIndex(295, (d_115_v14_).length(0))
                    (d_115_v14_)[index7_] = d_112_v11_
                    d_116_v15_: _dafny.Map
                    d_116_v15_ = _dafny.Map({d_99_v3_: d_98_v2_})
                    (d_103_globalState_).f0 = (d_113_v12_)[default__.safeIndex(len(d_116_v15_), len(d_113_v12_))]
                    d_109_v8_ = (d_109_v8_) | (d_109_v8_)
                    pass
            pass
        hi0_ = d_99_v3_
        for d_117_i2_ in range(d_99_v3_, hi0_):
            d_118_v16_: str
            d_118_v16_ = _dafny.CodePoint('o')
            d_118_v16_ = d_118_v16_
            d_119_v17_: _dafny.Seq
            d_119_v17_ = _dafny.SeqWithoutIsStrInference([d_99_v3_])
            d_120_v18_: _dafny.MultiSet
            d_120_v18_ = _dafny.MultiSet([(d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], d_98_v2_, d_98_v2_, d_98_v2_])
            d_121_v19_: _dafny.Array
            nw18_ = _dafny.Array(None, 9)
            nw18_[int(0)] = -393
            nw18_[int(1)] = (d_99_v3_ if False else d_99_v3_)
            nw18_[int(2)] = default__.safeDivisionInt(d_117_i2_, d_117_i2_)
            nw18_[int(3)] = d_99_v3_
            nw18_[int(4)] = d_99_v3_
            nw18_[int(5)] = d_117_i2_
            nw18_[int(6)] = len(((d_119_v17_).set(default__.safeIndex(d_117_i2_, len(d_119_v17_)), d_99_v3_)) + (d_119_v17_))
            nw18_[int(7)] = ((d_120_v18_).cardinality if (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))] else d_117_i2_)
            nw18_[int(8)] = (d_119_v17_)[default__.safeIndex(d_99_v3_, len(d_119_v17_))]
            d_121_v19_ = nw18_
            index8_ = default__.safeIndex(324, (d_121_v19_).length(0))
            (d_121_v19_)[index8_] = default__.safeModuloInt(d_99_v3_, (d_120_v18_).cardinality)
            index9_ = default__.safeIndex(324, (d_121_v19_).length(0))
            (d_121_v19_)[index9_] = default__.safeDivisionInt(d_99_v3_, d_99_v3_)
            d_122_v20_: D0
            d_122_v20_ = D0_DC0((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))])
            source2_ = d_122_v20_
            if source2_.is_DC1:
                d_123___mcc_h0_ = source2_.cf1
                d_124___mcc_h1_ = source2_.cf2
                d_125___mcc_h2_ = source2_.cf3
                d_126___mcc_h3_ = source2_.cf4
                d_127___mcc_h4_ = source2_.cf5
                d_128_cf5_ = d_127___mcc_h4_
                d_129_cf4_ = d_126___mcc_h3_
                d_130_cf3_ = d_125___mcc_h2_
                d_131_cf2_ = d_124___mcc_h1_
                d_132_cf1_ = d_123___mcc_h0_
                index10_ = default__.safeIndex(810, (d_104_v7_).length(0))
                (d_104_v7_)[index10_] = d_128_cf5_
                (d_103_globalState_).f0 = d_117_i2_
                d_133_v21_: _dafny.Seq
                d_133_v21_ = _dafny.SeqWithoutIsStrInference([d_130_cf3_, d_98_v2_])
                (d_103_globalState_).f1 = (d_133_v21_)[default__.safeIndex(d_132_cf1_, len(d_133_v21_))]
                d_134_v22_: D0
                d_134_v22_ = D0_DC2()
                d_135_v23_: _dafny.Array
                nw19_ = _dafny.Array(None, 15)
                nw19_[int(0)] = d_134_v22_
                nw19_[int(1)] = d_134_v22_
                nw19_[int(2)] = d_134_v22_
                nw19_[int(3)] = D0_DC2()
                nw19_[int(4)] = d_134_v22_
                nw19_[int(5)] = d_134_v22_
                nw19_[int(6)] = d_134_v22_
                nw19_[int(7)] = d_134_v22_
                nw19_[int(8)] = d_134_v22_
                nw19_[int(9)] = (d_134_v22_ if d_98_v2_ else d_134_v22_)
                nw19_[int(10)] = D0_DC2()
                nw19_[int(11)] = D0_DC2()
                nw19_[int(12)] = d_134_v22_
                nw19_[int(13)] = D0_DC2()
                nw19_[int(14)] = d_134_v22_
                d_135_v23_ = nw19_
                index11_ = default__.safeIndex(880, (d_135_v23_).length(0))
                (d_135_v23_)[index11_] = d_134_v22_
                d_136_v24_: _dafny.Map
                d_136_v24_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(d_128_cf5_), d_130_cf3_, d_98_v2_])): (d_117_i2_) * ((d_121_v19_)[default__.safeIndex(324, (d_121_v19_).length(0))])})
                d_137_v25_: _dafny.Array
                def lambda4_(d_138_cf2_):
                    def lambda5_(d_139_i3_):
                        return d_138_cf2_

                    return lambda5_

                init2_ = lambda4_(d_131_cf2_)
                nw20_ = _dafny.Array(None, 18)
                for i0_2_ in range(nw20_.length(0)):
                    nw20_[i0_2_] = init2_(i0_2_)
                d_137_v25_ = nw20_
                index12_ = default__.safeIndex(504, (d_137_v25_).length(0))
                (d_137_v25_)[index12_] = (d_131_cf2_ if d_130_cf3_ else _dafny.SeqWithoutIsStrInference([d_118_v16_ for d_140_i4_ in range(default__.abs(363))]))
                d_141_v26_: _dafny.Array
                nw21_ = _dafny.Array(None, 15)
                nw21_[int(0)] = d_119_v17_
                nw21_[int(1)] = d_119_v17_
                nw21_[int(2)] = (d_119_v17_) + (d_119_v17_)
                nw21_[int(3)] = (d_119_v17_) + (d_119_v17_)
                nw21_[int(4)] = d_119_v17_
                nw21_[int(5)] = (d_119_v17_ if d_98_v2_ else _dafny.SeqWithoutIsStrInference([(d_121_v19_)[default__.safeIndex(324, (d_121_v19_).length(0))], d_99_v3_, d_99_v3_, d_132_cf1_, len(d_131_cf2_)]))
                nw21_[int(6)] = d_119_v17_
                nw21_[int(7)] = d_119_v17_
                nw21_[int(8)] = d_119_v17_
                nw21_[int(9)] = (d_119_v17_) + (d_119_v17_)
                nw21_[int(10)] = d_119_v17_
                nw21_[int(11)] = _dafny.SeqWithoutIsStrInference([d_117_i2_, d_132_cf1_, (d_121_v19_)[default__.safeIndex(324, (d_121_v19_).length(0))], d_99_v3_, d_99_v3_])
                nw21_[int(12)] = (_dafny.SeqWithoutIsStrInference([(d_121_v19_)[default__.safeIndex(324, (d_121_v19_).length(0))], d_132_cf1_])) + (d_119_v17_)
                nw21_[int(13)] = default__.fm3(d_132_cf1_, d_130_cf3_, d_99_v3_, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], d_103_globalState_)
                nw21_[int(14)] = d_119_v17_
                d_141_v26_ = nw21_
                index13_ = default__.safeIndex(967, (d_141_v26_).length(0))
                (d_141_v26_)[index13_] = d_119_v17_
                index14_ = default__.safeIndex(880, (d_135_v23_).length(0))
                index15_ = default__.safeIndex(504, (d_137_v25_).length(0))
                index16_ = default__.safeIndex(967, (d_141_v26_).length(0))
                def iife19_():
                    coll5_ = _dafny.Set()
                    compr_5_: int
                    for compr_5_ in _dafny.IntegerRange(388, 253):
                        d_142_v27_: int = compr_5_
                        if ((388) <= (d_142_v27_)) and ((d_142_v27_) < (253)):
                            coll5_ = coll5_.union(_dafny.Set([(d_142_v27_) + (len(d_119_v17_))]))
                    return _dafny.Set(coll5_)
                rhs7_ = d_134_v22_
                rhs8_ = d_136_v24_
                rhs9_ = d_97_v1_
                rhs10_ = ((d_119_v17_) + (_dafny.SeqWithoutIsStrInference([d_99_v3_, default__.fm2(d_117_i2_, d_132_cf1_, d_103_globalState_)]))).set(default__.safeIndex(default__.fm2(d_99_v3_, len(iife19_()
                ), d_103_globalState_), len((d_119_v17_) + (_dafny.SeqWithoutIsStrInference([d_99_v3_, default__.fm2(d_117_i2_, d_132_cf1_, d_103_globalState_)])))), (len(d_133_v21_)) * (d_129_cf4_))
                lhs3_ = d_135_v23_
                lhs4_ = default__.safeIndex(880, (d_135_v23_).length(0))
                lhs5_ = d_137_v25_
                lhs6_ = default__.safeIndex(504, (d_137_v25_).length(0))
                lhs7_ = d_141_v26_
                lhs8_ = default__.safeIndex(967, (d_141_v26_).length(0))
                lhs3_[lhs4_] = rhs7_
                d_136_v24_ = rhs8_
                lhs5_[lhs6_] = rhs9_
                lhs7_[lhs8_] = rhs10_
            elif source2_.is_DC2:
                d_143_v28_: _dafny.Seq
                d_143_v28_ = _dafny.SeqWithoutIsStrInference([d_119_v17_, d_119_v17_])
                index17_ = default__.safeIndex(810, (d_104_v7_).length(0))
                rhs11_ = not(d_98_v2_)
                rhs12_ = _dafny.SeqWithoutIsStrInference([d_119_v17_ for d_144_i5_ in range(default__.abs(-947))])
                rhs13_ = (0) - (d_117_i2_)
                lhs9_ = d_104_v7_
                lhs10_ = default__.safeIndex(810, (d_104_v7_).length(0))
                lhs11_ = d_103_globalState_
                lhs9_[lhs10_] = rhs11_
                d_143_v28_ = rhs12_
                lhs11_.f0 = rhs13_
                d_145_v29_: _dafny.MultiSet
                d_145_v29_ = _dafny.MultiSet([d_118_v16_, d_118_v16_, default__.fm4(D0_DC1(d_99_v3_, d_97_v1_, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], d_117_i2_, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]), not(d_98_v2_), (d_121_v19_)[default__.safeIndex(324, (d_121_v19_).length(0))], d_103_globalState_)])
                index18_ = default__.safeIndex(810, (d_104_v7_).length(0))
                index19_ = default__.safeIndex(324, (d_121_v19_).length(0))
                rhs14_ = d_145_v29_
                rhs15_ = ((d_121_v19_)[default__.safeIndex(324, (d_121_v19_).length(0))]) <= ((d_121_v19_)[default__.safeIndex(324, (d_121_v19_).length(0))])
                rhs16_ = d_98_v2_
                rhs17_ = d_117_i2_
                lhs12_ = d_104_v7_
                lhs13_ = default__.safeIndex(810, (d_104_v7_).length(0))
                lhs14_ = d_121_v19_
                lhs15_ = default__.safeIndex(324, (d_121_v19_).length(0))
                d_145_v29_ = rhs14_
                lhs12_[lhs13_] = rhs15_
                d_98_v2_ = rhs16_
                lhs14_[lhs15_] = rhs17_
                (d_103_globalState_).f2 = d_117_i2_
                d_98_v2_ = (d_99_v3_) < (d_117_i2_)
            elif True:
                d_146___mcc_h5_ = source2_.cf0
                d_147_cf0_ = d_146___mcc_h5_
                (d_103_globalState_).f1 = d_98_v2_
                d_148_v30_: int
                out2_: int
                out2_ = default__.m0(not(d_147_cf0_), d_103_globalState_)
                d_148_v30_ = out2_
                d_149_v31_: _dafny.Map
                d_149_v31_ = _dafny.Map({d_122_v20_: d_117_i2_})
                d_147_cf0_ = not(default__.fm5(d_147_cf0_, (d_149_v31_) | (d_149_v31_), d_148_v30_, d_103_globalState_))
                d_150_v32_: int
                out3_: int
                out3_ = default__.m0((d_98_v2_) == (d_98_v2_), d_103_globalState_)
                d_150_v32_ = out3_
            d_151_v33_: C1
            nw22_ = C1()
            nw22_.ctor__((_dafny.Set({False})) == ((default__.fm39((d_121_v19_)[default__.safeIndex(324, (d_121_v19_).length(0))], not((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]), d_103_globalState_)).cf15))
            d_151_v33_ = nw22_
        d_152_v34_: _dafny.Seq
        d_152_v34_ = _dafny.SeqWithoutIsStrInference([not(d_98_v2_), (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]])
        d_153_v35_: _dafny.Seq
        d_153_v35_ = _dafny.SeqWithoutIsStrInference([d_99_v3_])
        d_154_v36_: str
        d_154_v36_ = _dafny.CodePoint('r')
        d_155_v38_: D0
        d_155_v38_ = D0_DC0((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))])
        d_156_v39_: _dafny.Seq
        d_156_v39_ = _dafny.SeqWithoutIsStrInference([D0_DC0((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]), D0_DC0(d_98_v2_), d_155_v38_])
        d_157_v40_: _dafny.Map
        def iife20_():
            coll6_ = _dafny.Map()
            compr_6_: D0
            for compr_6_ in (d_156_v39_).Elements:
                d_158_v37_: D0 = compr_6_
                if (d_158_v37_) in (d_156_v39_):
                    coll6_[d_158_v37_] = len(d_100_v4_)
            return _dafny.Map(coll6_)
        d_157_v40_ = _dafny.Map({d_154_v36_: iife20_()
        })
        d_159_v41_: C3
        nw23_ = C3()
        nw23_.ctor__((d_97_v1_ if not((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]) else d_97_v1_), d_98_v2_, default__.fm5((d_152_v34_)[default__.safeIndex(len(d_153_v35_), len(d_152_v34_))], ((d_157_v40_)[d_154_v36_] if (d_154_v36_) in (d_157_v40_) else _dafny.Map({d_155_v38_: len(d_97_v1_)})), d_99_v3_, d_103_globalState_), default__.safeDivisionInt(d_99_v3_, d_99_v3_))
        d_159_v41_ = nw23_
        d_160_v42_: C2
        nw24_ = C2()
        nw24_.ctor__((0) - (d_99_v3_), not((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]), d_99_v3_)
        d_160_v42_ = nw24_
        d_160_v42_ = d_160_v42_
        d_161_v43_: _dafny.Array
        nw25_ = _dafny.Array(_dafny.Map({}), 20)
        d_161_v43_ = nw25_
        d_162_v44_: _dafny.Array
        nw26_ = _dafny.Array(None, 21)
        d_162_v44_ = nw26_
        index20_ = default__.safeIndex(390, (d_161_v43_).length(0))
        (d_161_v43_)[index20_] = _dafny.Map({(d_160_v42_).f31: d_162_v44_})
        d_163_v45_: _dafny.Map
        d_163_v45_ = _dafny.Map({(d_160_v42_).f31: d_162_v44_})
        index21_ = default__.safeIndex(390, (d_161_v43_).length(0))
        (d_161_v43_)[index21_] = d_163_v45_
        hi1_ = (d_160_v42_).f31
        for d_164_i6_ in range((d_159_v41_).fm6(d_98_v2_, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], d_99_v3_, d_103_globalState_), hi1_):
            if (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]:
                d_165_v46_: _dafny.Map
                d_165_v46_ = _dafny.Map({d_164_i6_: d_152_v34_})
                d_166_v47_: _dafny.Seq
                d_166_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_99_v3_: d_152_v34_}), d_165_v46_, d_165_v46_])
                d_100_v4_ = (d_100_v4_).set((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], len((d_166_v47_)[default__.safeIndex((d_160_v42_).f31, len(d_166_v47_))]))
                d_154_v36_ = _dafny.CodePoint('f')
                d_167_v48_: _dafny.Map
                d_167_v48_ = _dafny.Map({d_98_v2_: d_97_v1_})
                (d_103_globalState_).f1 = (not((d_159_v41_).f30) if (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))] else ((0) - (d_99_v3_)) != (len(((d_167_v48_)[d_98_v2_] if (d_98_v2_) in (d_167_v48_) else (d_159_v41_).f29))))
                d_168_v49_: _dafny.Map
                d_168_v49_ = _dafny.Map({(d_160_v42_).fm11(_dafny.Map({d_97_v1_: d_97_v1_}), (d_160_v42_).f31, len((d_159_v41_).f29), d_103_globalState_): d_159_v41_})
                d_168_v49_ = (d_168_v49_).set(d_164_i6_, d_159_v41_)
                d_169_v50_: _dafny.Array
                nw27_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_169_v50_ = nw27_
                d_170_v51_: _dafny.Array
                def lambda6_(d_171_v42_):
                    def lambda7_(d_172_i7_):
                        return (d_172_i7_) - ((d_171_v42_).f31)

                    return lambda7_

                init3_ = lambda6_(d_160_v42_)
                nw28_ = _dafny.Array(None, 17)
                for i0_3_ in range(nw28_.length(0)):
                    nw28_[i0_3_] = init3_(i0_3_)
                d_170_v51_ = nw28_
                index22_ = default__.safeIndex(355, (d_169_v50_).length(0))
                (d_169_v50_)[index22_] = d_170_v51_
                d_173_v52_: D12
                d_173_v52_ = D12_DC28(d_170_v51_)
                index23_ = default__.safeIndex(355, (d_169_v50_).length(0))
                (d_169_v50_)[index23_] = (d_173_v52_).cf44
            elif True:
                (d_103_globalState_).f2 = default__.safeModuloInt(805, d_164_i6_)
                rhs18_ = (d_159_v41_).f29
                rhs19_ = (d_160_v42_).fm6((d_159_v41_).f30, (d_159_v41_).f30, (0) - ((d_160_v42_).f31), d_103_globalState_)
                rhs20_ = d_98_v2_
                lhs16_ = d_103_globalState_
                d_97_v1_ = rhs18_
                lhs16_.f2 = rhs19_
                d_98_v2_ = rhs20_
                d_174_v53_: _dafny.Set
                d_174_v53_ = _dafny.Set({d_104_v7_, d_104_v7_, d_104_v7_, d_104_v7_})
                (d_103_globalState_).f1 = (len(((d_97_v1_).set(default__.safeIndex((d_160_v42_).f31, len(d_97_v1_)), d_154_v36_)) + (d_97_v1_))) == (((d_160_v42_).f31) * (len(d_174_v53_)))
                d_104_v7_ = (d_104_v7_ if (d_159_v41_).f30 else d_104_v7_)
                d_175_v54_: _dafny.Array
                def lambda8_(d_176_v41_, d_177_v3_):
                    def lambda9_(d_178_i8_):
                        return D9_DC19(D6_DC14(len((d_176_v41_).f29), d_177_v3_, _dafny.MultiSet([(d_176_v41_).f30])))

                    return lambda9_

                init4_ = lambda8_(d_159_v41_, d_99_v3_)
                nw29_ = _dafny.Array(None, 29)
                for i0_4_ in range(nw29_.length(0)):
                    nw29_[i0_4_] = init4_(i0_4_)
                d_175_v54_ = nw29_
                d_175_v54_ = d_175_v54_
            d_179_v55_: _dafny.Array
            nw30_ = _dafny.Array(_dafny.Map({}), 22)
            d_179_v55_ = nw30_
            d_180_v56_: _dafny.Map
            d_180_v56_ = _dafny.Map({((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]) and ((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]): d_179_v55_})
            d_179_v55_ = ((d_180_v56_)[(d_159_v41_).f30] if ((d_159_v41_).f30) in (d_180_v56_) else d_179_v55_)
            (d_103_globalState_).f2 = (0) - ((default__.safeModuloInt((d_160_v42_).f31, (d_160_v42_).f31)) + (377))
            d_181_v57_: C1
            nw31_ = C1()
            nw31_.ctor__((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))])
            d_181_v57_ = nw31_
        if not(not ((d_159_v41_).f30) or (not(((d_159_v41_).f29) != (d_97_v1_)))):
            if (d_159_v41_).f30:
                d_182_v58_: C1
                nw32_ = C1()
                nw32_.ctor__((d_159_v41_).f30)
                d_182_v58_ = nw32_
                d_182_v58_ = d_182_v58_
                d_183_v59_: _dafny.MultiSet
                d_183_v59_ = _dafny.MultiSet([(d_159_v41_).f30, (d_182_v58_).f32, (d_182_v58_).f32])
                d_184_v60_: _dafny.Map
                d_184_v60_ = _dafny.Map({(d_160_v42_).f31: not((d_159_v41_).f30)})
                rhs21_ = (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]
                rhs22_ = ((d_160_v42_).f31 if False else (len((d_159_v41_).f29)) + ((d_160_v42_).f31))
                rhs23_ = (d_159_v41_).f30
                rhs24_ = (d_160_v42_).f31
                rhs25_ = len((_dafny.Map({((d_183_v59_)[(d_159_v41_).f30] if ((d_159_v41_).f30) in (d_183_v59_) else (d_160_v42_).f31): (d_159_v41_).f30})) | (d_184_v60_))
                lhs17_ = d_103_globalState_
                lhs18_ = d_103_globalState_
                lhs19_ = d_103_globalState_
                lhs20_ = d_103_globalState_
                lhs21_ = d_103_globalState_
                lhs17_.f1 = rhs21_
                lhs18_.f0 = rhs22_
                lhs19_.f1 = rhs23_
                lhs20_.f2 = rhs24_
                lhs21_.f0 = rhs25_
                d_185_v61_: D17
                d_185_v61_ = D17_DC37(d_160_v42_)
                d_186_v62_: _dafny.Seq
                d_186_v62_ = _dafny.SeqWithoutIsStrInference([d_160_v42_, d_160_v42_, (d_185_v61_).cf63, (d_160_v42_ if (d_182_v58_).f32 else d_160_v42_), d_160_v42_])
                d_186_v62_ = ((d_186_v62_).set(default__.safeIndex(len(d_152_v34_), len(d_186_v62_)), d_160_v42_)) + (_dafny.SeqWithoutIsStrInference([d_160_v42_, d_160_v42_, d_160_v42_, d_160_v42_, d_160_v42_]))
                (d_103_globalState_).f2 = ((d_160_v42_).f31 if (d_152_v34_)[default__.safeIndex((d_160_v42_).f31, len(d_152_v34_))] else d_99_v3_)
                d_187_v63_: _dafny.Array
                nw33_ = _dafny.Array(int(0), 19)
                d_187_v63_ = nw33_
                d_188_v64_: _dafny.Seq
                d_188_v64_ = _dafny.SeqWithoutIsStrInference([d_187_v63_, (d_187_v63_ if (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))] else d_187_v63_), d_187_v63_, d_187_v63_])
                (d_103_globalState_).f0 = len(d_188_v64_)
            elif True:
                d_189_v65_: C4
                nw34_ = C4()
                nw34_.ctor__(d_99_v3_, (d_159_v41_).f29, (d_159_v41_).f30, (d_160_v42_).f31)
                d_189_v65_ = nw34_
                d_190_v66_: _dafny.Array
                nw35_ = _dafny.Array(int(0), 3)
                d_190_v66_ = nw35_
                d_190_v66_ = d_190_v66_
                d_191_v67_: D16
                d_191_v67_ = D16_DC36((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], d_153_v35_, d_104_v7_)
                d_192_v68_: int
                d_193_v69_: bool
                d_194_v70_: _dafny.Map
                d_195_v71_: _dafny.Map
                out4_: int
                out5_: bool
                out6_: _dafny.Map
                out7_: _dafny.Map
                out4_, out5_, out6_, out7_ = (d_189_v65_).m2((d_191_v67_).cf60, d_103_globalState_)
                d_192_v68_ = out4_
                d_193_v69_ = out5_
                d_194_v70_ = out6_
                d_195_v71_ = out7_
                d_97_v1_ = (d_159_v41_).f29
                d_196_v72_: _dafny.Array
                nw36_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_196_v72_ = nw36_
                index24_ = default__.safeIndex(668, (d_196_v72_).length(0))
                (d_196_v72_)[index24_] = d_190_v66_
                d_197_v73_: _dafny.Array
                nw37_ = _dafny.Array(None, 7)
                nw37_[int(0)] = d_153_v35_
                nw37_[int(1)] = (d_153_v35_) + (d_153_v35_)
                nw37_[int(2)] = d_153_v35_
                nw37_[int(3)] = _dafny.SeqWithoutIsStrInference([(d_160_v42_).f31 for d_198_i9_ in range(default__.abs(-357))])
                nw37_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_189_v65_).f27 for d_199_i10_ in range(default__.abs(430))])
                nw37_[int(5)] = d_153_v35_
                nw37_[int(6)] = _dafny.SeqWithoutIsStrInference([(d_160_v42_).f31 for d_200_i11_ in range(default__.abs(768))])
                d_197_v73_ = nw37_
                index25_ = default__.safeIndex(805, (d_197_v73_).length(0))
                (d_197_v73_)[index25_] = (d_153_v35_) + (d_153_v35_)
                index26_ = default__.safeIndex(668, (d_196_v72_).length(0))
                index27_ = default__.safeIndex(805, (d_197_v73_).length(0))
                rhs26_ = d_190_v66_
                rhs27_ = (d_160_v42_).f31
                rhs28_ = (default__.fm40(d_98_v2_, d_103_globalState_)).cf55
                rhs29_ = _dafny.SeqWithoutIsStrInference([d_99_v3_, len((default__.fm1(d_103_globalState_) if (d_159_v41_).f30 else (d_159_v41_).f29)), (d_160_v42_).f31])
                lhs22_ = d_196_v72_
                lhs23_ = default__.safeIndex(668, (d_196_v72_).length(0))
                lhs24_ = d_103_globalState_
                lhs25_ = d_189_v65_
                lhs26_ = d_197_v73_
                lhs27_ = default__.safeIndex(805, (d_197_v73_).length(0))
                lhs22_[lhs23_] = rhs26_
                lhs24_.f2 = rhs27_
                lhs25_.f28 = rhs28_
                lhs26_[lhs27_] = rhs29_
            (d_103_globalState_).f1 = (d_159_v41_).f30
            d_152_v34_ = (d_152_v34_) + (d_152_v34_)
            d_201_v74_: _dafny.Map
            d_201_v74_ = _dafny.Map({(d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]: d_98_v2_})
            d_202_v75_: C4
            nw38_ = C4()
            nw38_.ctor__(len(d_201_v74_), (d_97_v1_).set(default__.safeIndex(len((d_159_v41_).f29), len(d_97_v1_)), d_154_v36_), d_98_v2_, (d_160_v42_).f31)
            d_202_v75_ = nw38_
            d_203_v76_: _dafny.MultiSet
            d_203_v76_ = _dafny.MultiSet([d_202_v75_])
            d_204_v77_: _dafny.Seq
            d_204_v77_ = _dafny.SeqWithoutIsStrInference([d_202_v75_])
            d_205_v78_: _dafny.MultiSet
            d_205_v78_ = _dafny.MultiSet(d_204_v77_)
            d_206_v79_: _dafny.Set
            d_206_v79_ = _dafny.Set({d_203_v76_, (d_203_v76_) - (d_203_v76_), d_203_v76_, (d_205_v78_), d_203_v76_})
            rhs30_ = d_99_v3_
            rhs31_ = default__.fm2(((d_202_v75_).f27) - (d_99_v3_), (d_202_v75_).fm6(d_98_v2_, (d_159_v41_).f30, d_99_v3_, d_103_globalState_), d_103_globalState_)
            rhs32_ = _dafny.Set({_dafny.MultiSet([d_202_v75_, d_202_v75_, d_202_v75_])})
            rhs33_ = default__.safeModuloInt((d_159_v41_).fm9(d_103_globalState_), 915)
            lhs28_ = d_103_globalState_
            lhs29_ = d_103_globalState_
            d_99_v3_ = rhs30_
            lhs28_.f0 = rhs31_
            d_206_v79_ = rhs32_
            lhs29_.f0 = rhs33_
            d_207_v80_: _dafny.Set
            d_207_v80_ = _dafny.Set({d_152_v34_, d_152_v34_, d_152_v34_, d_152_v34_})
            d_208_v81_: _dafny.Seq
            d_208_v81_ = _dafny.SeqWithoutIsStrInference([d_152_v34_, _dafny.SeqWithoutIsStrInference([(d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]]), d_152_v34_, d_152_v34_, d_152_v34_])
            def iife21_():
                coll7_ = _dafny.Set()
                compr_7_: _dafny.Seq
                for compr_7_ in (d_208_v81_).Elements:
                    d_209_v82_: _dafny.Seq = compr_7_
                    if (d_209_v82_) in (d_208_v81_):
                        coll7_ = coll7_.union(_dafny.Set([d_209_v82_]))
                return _dafny.Set(coll7_)
            (d_103_globalState_).f0 = (0) - (len((d_207_v80_) | (iife21_()
            )))
        elif True:
            d_97_v1_ = d_97_v1_
            if (d_98_v2_ if not((d_159_v41_).f30) else not(True)):
                (d_103_globalState_).f0 = d_99_v3_
                (d_103_globalState_).f0 = 338
                d_210_v83_: C1
                nw39_ = C1()
                nw39_.ctor__(False)
                d_210_v83_ = nw39_
                d_211_v84_: C4
                nw40_ = C4()
                nw40_.ctor__(903, (d_159_v41_).f29, (d_210_v83_).f32, (d_160_v42_).f31)
                d_211_v84_ = nw40_
                d_212_v85_: _dafny.MultiSet
                d_212_v85_ = _dafny.MultiSet([d_211_v84_, d_211_v84_])
                d_213_v86_: D19
                d_213_v86_ = D19_DC41(d_211_v84_)
                d_214_v87_: _dafny.Seq
                d_214_v87_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_211_v84_).f27})])
                d_215_v88_: _dafny.Seq
                d_215_v88_ = _dafny.SeqWithoutIsStrInference([d_153_v35_])
                d_216_v89_: _dafny.Seq
                d_216_v89_ = _dafny.SeqWithoutIsStrInference([d_104_v7_])
                d_217_v90_: _dafny.Map
                d_217_v90_ = _dafny.Map({len(_dafny.Set({(d_211_v84_).f27})): (d_97_v1_).set(default__.safeIndex((d_211_v84_).f27, len(d_97_v1_)), d_154_v36_)})
                d_218_v91_: _dafny.Map
                d_218_v91_ = _dafny.Map({not((d_159_v41_).f30): d_104_v7_})
                d_219_v92_: D20
                d_219_v92_ = D20_DC44(d_218_v91_)
                d_220_v93_: _dafny.Map
                d_220_v93_ = _dafny.Map({(d_219_v92_).cf73: _dafny.SeqWithoutIsStrInference([d_154_v36_ for d_221_i13_ in range(default__.abs(140))])})
                d_222_v94_: _dafny.Set
                d_222_v94_ = _dafny.Set({not((d_160_v42_).fm7(d_103_globalState_)), (d_210_v83_).f32})
                d_223_v95_: _dafny.Map
                d_223_v95_ = _dafny.Map({d_222_v94_: False})
                d_224_v96_: _dafny.Map
                d_224_v96_ = _dafny.Map({d_223_v95_: (d_160_v42_).f31})
                d_225_v97_: _dafny.Array
                nw41_ = _dafny.Array(None, 22)
                nw41_[int(0)] = ((d_212_v85_)[(d_213_v86_).cf66] if ((d_213_v86_).cf66) in (d_212_v85_) else (d_211_v84_).f27)
                nw41_[int(1)] = (len((d_159_v41_).f29)) - ((d_211_v84_).f27)
                nw41_[int(2)] = (d_160_v42_).fm11(_dafny.Map({_dafny.SeqWithoutIsStrInference([d_154_v36_ for d_226_i12_ in range(default__.abs(629))]): d_97_v1_}), d_99_v3_, (d_160_v42_).f31, d_103_globalState_)
                nw41_[int(3)] = default__.safeDivisionInt(d_99_v3_, (d_160_v42_).f31)
                nw41_[int(4)] = default__.safeDivisionInt(len(d_214_v87_), (d_211_v84_).f27)
                nw41_[int(5)] = (0) - ((d_211_v84_).f27)
                nw41_[int(6)] = len(d_215_v88_)
                nw41_[int(7)] = (d_160_v42_).f31
                nw41_[int(8)] = (671) + ((d_211_v84_).f27)
                nw41_[int(9)] = (0) - ((d_160_v42_).f31)
                nw41_[int(10)] = d_99_v3_
                nw41_[int(11)] = len(_dafny.Map({(d_216_v89_)[default__.safeIndex(len(d_217_v90_), len(d_216_v89_))]: (d_160_v42_).f31}))
                nw41_[int(12)] = (d_160_v42_).f31
                nw41_[int(13)] = default__.safeModuloInt((d_160_v42_).f31, (d_160_v42_).f31)
                nw41_[int(14)] = (d_211_v84_).f27
                nw41_[int(15)] = len(default__.fm1(d_103_globalState_))
                nw41_[int(16)] = (0) - ((d_160_v42_).f31)
                nw41_[int(17)] = (0) - ((len(((d_220_v93_)[d_218_v91_] if (d_218_v91_) in (d_220_v93_) else (d_159_v41_).f29))) * (890))
                nw41_[int(18)] = d_99_v3_
                nw41_[int(19)] = -639
                nw41_[int(20)] = (d_160_v42_).f31
                nw41_[int(21)] = ((d_224_v96_)[d_223_v95_] if (d_223_v95_) in (d_224_v96_) else (d_211_v84_).f27)
                d_225_v97_ = nw41_
                index28_ = default__.safeIndex(553, (d_225_v97_).length(0))
                (d_225_v97_)[index28_] = (d_160_v42_).f31
                index29_ = default__.safeIndex(553, (d_225_v97_).length(0))
                (d_225_v97_)[index29_] = (((d_211_v84_).f27) * ((d_160_v42_).f31)) * (len((d_159_v41_).f29))
                index30_ = default__.safeIndex(810, (d_104_v7_).length(0))
                (d_104_v7_)[index30_] = (d_160_v42_).fm7(d_103_globalState_)
            elif True:
                (d_103_globalState_).f1 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) < ((d_159_v41_).f29)
                d_153_v35_ = _dafny.SeqWithoutIsStrInference([len((d_159_v41_).f29)])
                rhs34_ = (d_159_v41_).f29
                rhs35_ = (d_159_v41_).f30
                lhs30_ = d_103_globalState_
                d_97_v1_ = rhs34_
                lhs30_.f1 = rhs35_
                d_227_v98_: _dafny.Set
                d_227_v98_ = _dafny.Set({d_97_v1_})
                (d_103_globalState_).f1 = (d_227_v98_).issubset(d_227_v98_)
                d_228_v99_: _dafny.Map
                d_228_v99_ = _dafny.Map({(d_159_v41_).f30: d_98_v2_})
                d_229_v100_: _dafny.Set
                d_229_v100_ = _dafny.Set({(d_159_v41_).f30})
                d_230_v101_: C4
                nw42_ = C4()
                nw42_.ctor__((d_160_v42_).f31, (d_159_v41_).f29, (d_159_v41_).f30, len(d_229_v100_))
                d_230_v101_ = nw42_
                d_231_v102_: _dafny.Set
                d_231_v102_ = _dafny.Set({d_230_v101_, d_230_v101_})
                d_232_v103_: D10
                d_232_v103_ = D10_DC23()
                d_233_v104_: C0
                nw43_ = C0()
                nw43_.ctor__(d_230_v101_.f28, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], d_99_v3_, d_104_v7_, (d_159_v41_).f30)
                d_233_v104_ = nw43_
                d_234_v105_: _dafny.Map
                d_234_v105_ = _dafny.Map({d_233_v104_: d_153_v35_})
                d_235_v107_: _dafny.MultiSet
                d_235_v107_ = _dafny.MultiSet([((d_100_v4_)[(d_233_v104_).fm17(not((d_159_v41_).f30), (d_159_v41_).f30, d_103_globalState_)] if ((d_233_v104_).fm17(not((d_159_v41_).f30), (d_159_v41_).f30, d_103_globalState_)) in (d_100_v4_) else (d_160_v42_).f31), (d_160_v42_).f31, (d_160_v42_).f31])
                d_236_v108_: T0
                nw44_ = C3()
                nw44_.ctor__(d_230_v101_.f28, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], 637)
                d_236_v108_ = nw44_
                d_237_v109_: _dafny.Map
                d_237_v109_ = _dafny.Map({(d_160_v42_).f31: d_236_v108_})
                d_238_v110_: _dafny.Array
                def lambda10_(d_239_v3_):
                    def lambda11_(d_240_i14_):
                        return default__.safeDivisionInt(d_240_i14_, d_239_v3_)

                    return lambda11_

                init5_ = lambda10_(d_99_v3_)
                nw45_ = _dafny.Array(None, 10)
                for i0_5_ in range(nw45_.length(0)):
                    nw45_[i0_5_] = init5_(i0_5_)
                d_238_v110_ = nw45_
                d_241_v111_: _dafny.Seq
                d_241_v111_ = _dafny.SeqWithoutIsStrInference([d_238_v110_])
                d_242_v112_: _dafny.Seq
                d_242_v112_ = _dafny.SeqWithoutIsStrInference([default__.fm3(len(d_237_v109_), (d_159_v41_).f30, len(d_241_v111_), not((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]), d_103_globalState_)])
                d_243_v113_: _dafny.Array
                nw46_ = _dafny.Array(None, 28)
                nw46_[int(0)] = _dafny.SeqWithoutIsStrInference([(d_159_v41_).fm6((d_159_v41_).f30, (d_159_v41_).f30, (d_160_v42_).f31, d_103_globalState_)])
                nw46_[int(1)] = d_153_v35_
                nw46_[int(2)] = (d_153_v35_) + (_dafny.SeqWithoutIsStrInference([125, len(d_228_v99_), (d_160_v42_).f31, (0) - ((d_160_v42_).f31)]))
                nw46_[int(3)] = ((d_153_v35_).set(default__.safeIndex((0) - ((d_160_v42_).f31), len(d_153_v35_)), (d_160_v42_).f31) if (d_159_v41_).f30 else d_153_v35_)
                nw46_[int(4)] = d_153_v35_
                nw46_[int(5)] = (d_153_v35_) + (d_153_v35_)
                nw46_[int(6)] = d_153_v35_
                nw46_[int(7)] = _dafny.SeqWithoutIsStrInference([(d_160_v42_).f31])
                nw46_[int(8)] = d_153_v35_
                nw46_[int(9)] = d_153_v35_
                nw46_[int(10)] = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_231_v102_])).cardinality])
                nw46_[int(11)] = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Set({d_232_v103_, D10_DC23()}): d_104_v7_})), len(d_153_v35_), (d_160_v42_).f31, d_99_v3_, (0) - ((d_160_v42_).f31)])
                nw46_[int(12)] = _dafny.SeqWithoutIsStrInference([811])
                nw46_[int(13)] = d_153_v35_
                nw46_[int(14)] = d_153_v35_
                nw46_[int(15)] = d_153_v35_
                nw46_[int(16)] = d_153_v35_
                nw46_[int(17)] = _dafny.SeqWithoutIsStrInference([307, (d_160_v42_).f31])
                nw46_[int(18)] = ((d_234_v105_)[d_233_v104_] if (d_233_v104_) in (d_234_v105_) else d_153_v35_)
                nw46_[int(19)] = (d_153_v35_) + (d_153_v35_)
                nw46_[int(20)] = d_153_v35_
                nw46_[int(21)] = (d_153_v35_).set(default__.safeIndex(d_99_v3_, len(d_153_v35_)), d_99_v3_)
                nw46_[int(22)] = d_153_v35_
                nw46_[int(23)] = d_153_v35_
                nw46_[int(24)] = (d_153_v35_) + (d_153_v35_)
                def iife22_():
                    coll8_ = _dafny.Map()
                    compr_8_: int
                    for compr_8_ in (d_235_v107_).Elements:
                        d_244_v106_: int = compr_8_
                        if (d_244_v106_) in (d_235_v107_):
                            coll8_[(d_244_v106_) - ((d_160_v42_).f31)] = 166
                    return _dafny.Map(coll8_)
                nw46_[int(25)] = (d_153_v35_) + ((d_153_v35_).set(default__.safeIndex(len(iife22_()
                ), len(d_153_v35_)), (d_160_v42_).f31))
                nw46_[int(26)] = d_153_v35_
                nw46_[int(27)] = (d_242_v112_)[default__.safeIndex((0) - ((d_236_v108_).f26), len(d_242_v112_))]
                d_243_v113_ = nw46_
                d_245_v114_: _dafny.Seq
                d_245_v114_ = _dafny.SeqWithoutIsStrInference([(d_243_v113_ if (d_159_v41_).f30 else d_243_v113_), d_243_v113_])
                d_243_v113_ = (d_245_v114_)[default__.safeIndex((d_236_v108_).f26, len(d_245_v114_))]
            d_246_v115_: T1
            nw47_ = C0()
            nw47_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkpbjugb")), (d_159_v41_).f30, (d_160_v42_).f31, d_104_v7_, ((d_159_v41_).f30) or (not((d_159_v41_).fm7(d_103_globalState_))))
            d_246_v115_ = nw47_
            d_246_v115_ = d_246_v115_
            d_247_v116_: _dafny.Map
            d_247_v116_ = _dafny.Map({not (True) or (d_98_v2_): (d_159_v41_).f30})
            d_248_v117_: _dafny.MultiSet
            d_248_v117_ = _dafny.MultiSet([d_98_v2_, (d_159_v41_).f30, d_98_v2_, (d_159_v41_).f30])
            d_247_v116_ = (d_247_v116_).set((default__.fm27(d_99_v3_, d_103_globalState_)).ispropersubset(d_248_v117_), d_98_v2_)
            (d_103_globalState_).f1 = (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]
        d_249_v118_: _dafny.Map
        d_249_v118_ = _dafny.Map({(d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]: d_154_v36_})
        if ((d_249_v118_) != (d_249_v118_)) in (d_152_v34_):
            d_250_v119_: _dafny.Set
            d_250_v119_ = _dafny.Set({(d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]})
            index31_ = default__.safeIndex(122, (d_96_v0_).length(0))
            (d_96_v0_)[index31_] = d_250_v119_
            index32_ = default__.safeIndex(122, (d_96_v0_).length(0))
            (d_96_v0_)[index32_] = d_250_v119_
            d_251_v120_: D1
            d_251_v120_ = D1_DC4(d_98_v2_, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))])
            d_252_v121_: D1
            d_252_v121_ = D1_DC5(d_251_v120_)
            d_253_v122_: D0
            d_253_v122_ = D0_DC1(d_99_v3_, (d_159_v41_).f29, (d_159_v41_).f30, (d_160_v42_).f31, (d_159_v41_).f30)
            d_254_v123_: bool
            d_255_v124_: bool
            out8_: bool
            out9_: bool
            out8_, out9_ = (d_160_v42_).m5(d_98_v2_, d_252_v121_, default__.fm4(d_253_v122_, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], len((d_159_v41_).f29), d_103_globalState_), d_103_globalState_)
            d_254_v123_ = out8_
            d_255_v124_ = out9_
            d_154_v36_ = d_154_v36_
            d_99_v3_ = (((0) - (d_99_v3_)) - (554)) - ((len(d_153_v35_)) * (d_99_v3_))
            d_256_v125_: _dafny.Array
            nw48_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_256_v125_ = nw48_
            d_257_v126_: _dafny.Map
            d_257_v126_ = _dafny.Map({172: d_152_v34_})
            d_258_v127_: C4
            nw49_ = C4()
            nw49_.ctor__(-205, (d_159_v41_).f29, (d_159_v41_).f30, len(d_250_v119_))
            d_258_v127_ = nw49_
            d_259_v128_: _dafny.Set
            d_259_v128_ = _dafny.Set({d_258_v127_, d_258_v127_})
            d_260_v129_: _dafny.Array
            nw50_ = _dafny.Array(None, 25)
            nw50_[int(0)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eblm")))
            nw50_[int(1)] = (d_160_v42_).f31
            nw50_[int(2)] = (d_160_v42_).f31
            nw50_[int(3)] = (d_160_v42_).f31
            nw50_[int(4)] = d_99_v3_
            nw50_[int(5)] = -139
            nw50_[int(6)] = 420
            nw50_[int(7)] = d_99_v3_
            nw50_[int(8)] = (d_160_v42_).f31
            nw50_[int(9)] = d_99_v3_
            nw50_[int(10)] = (d_160_v42_).f31
            nw50_[int(11)] = len(d_257_v126_)
            nw50_[int(12)] = 952
            nw50_[int(13)] = len((d_96_v0_)[default__.safeIndex(122, (d_96_v0_).length(0))])
            nw50_[int(14)] = d_99_v3_
            nw50_[int(15)] = (d_160_v42_).f31
            nw50_[int(16)] = (d_160_v42_).f31
            nw50_[int(17)] = (d_160_v42_).f31
            nw50_[int(18)] = (d_160_v42_).f31
            nw50_[int(19)] = d_99_v3_
            nw50_[int(20)] = 576
            nw50_[int(21)] = (d_160_v42_).f31
            nw50_[int(22)] = len(d_259_v128_)
            nw50_[int(23)] = len(_dafny.SeqWithoutIsStrInference([d_254_v123_, d_254_v123_, False]))
            nw50_[int(24)] = (d_258_v127_).f27
            d_260_v129_ = nw50_
            index33_ = default__.safeIndex(64, (d_256_v125_).length(0))
            (d_256_v125_)[index33_] = d_260_v129_
            index34_ = default__.safeIndex(64, (d_256_v125_).length(0))
            index35_ = default__.safeIndex(810, (d_104_v7_).length(0))
            rhs36_ = d_260_v129_
            rhs37_ = default__.safeModuloInt(89, (0) - (((d_160_v42_).f31) * ((d_160_v42_).f31)))
            rhs38_ = (d_258_v127_).fm7(d_103_globalState_)
            lhs31_ = d_256_v125_
            lhs32_ = default__.safeIndex(64, (d_256_v125_).length(0))
            lhs33_ = d_103_globalState_
            lhs34_ = d_104_v7_
            lhs35_ = default__.safeIndex(810, (d_104_v7_).length(0))
            lhs31_[lhs32_] = rhs36_
            lhs33_.f0 = rhs37_
            lhs34_[lhs35_] = rhs38_
        elif True:
            d_261_v130_: _dafny.Seq
            d_261_v130_ = _dafny.SeqWithoutIsStrInference([d_159_v41_, d_159_v41_])
            d_262_v131_: bool
            out10_: bool
            out10_ = (d_160_v42_).m6(_dafny.Map({(d_160_v42_).f31: (d_160_v42_).f31}), len(d_261_v130_), 257, d_103_globalState_)
            d_262_v131_ = out10_
            if not(((d_97_v1_) < (d_97_v1_) if (d_159_v41_).f30 else (d_159_v41_).f30)):
                d_263_v132_: _dafny.Array
                nw51_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_263_v132_ = nw51_
                d_264_v133_: _dafny.Array
                def lambda12_(d_265_v42_):
                    def lambda13_(d_266_i15_):
                        return (d_266_i15_) + ((d_265_v42_).f31)

                    return lambda13_

                init6_ = lambda12_(d_160_v42_)
                nw52_ = _dafny.Array(None, 1)
                for i0_6_ in range(nw52_.length(0)):
                    nw52_[i0_6_] = init6_(i0_6_)
                d_264_v133_ = nw52_
                index36_ = default__.safeIndex(987, (d_263_v132_).length(0))
                (d_263_v132_)[index36_] = d_264_v133_
                index37_ = default__.safeIndex(987, (d_263_v132_).length(0))
                (d_263_v132_)[index37_] = d_264_v133_
                d_267_v134_: T0
                nw53_ = C4()
                nw53_.ctor__((0) - ((d_160_v42_).f31), (d_159_v41_).f29, (d_159_v41_).f30, (d_160_v42_).f31)
                d_267_v134_ = nw53_
                d_268_v135_: _dafny.Set
                d_268_v135_ = _dafny.Set({d_267_v134_})
                d_269_v136_: _dafny.Seq
                d_269_v136_ = _dafny.SeqWithoutIsStrInference([d_268_v135_])
                d_270_v137_: D20
                d_270_v137_ = D20_DC46(True, (d_160_v42_).f31)
                d_271_v138_: bool
                out11_: bool
                out11_ = (d_160_v42_).m6(_dafny.Map({(default__.fm35((0) - (d_99_v3_), (_dafny.MultiSet(d_269_v136_)).cardinality, d_103_globalState_)).cardinality: (d_270_v137_).cf76}), d_99_v3_, (d_267_v134_).fm6(d_98_v2_, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], d_99_v3_, d_103_globalState_), d_103_globalState_)
                d_271_v138_ = out11_
                d_272_v139_: D2
                d_272_v139_ = D2_DC6(d_102_v6_)
                d_273_v140_: _dafny.Seq
                d_273_v140_ = _dafny.SeqWithoutIsStrInference([D2_DC6(d_102_v6_), d_272_v139_, d_272_v139_])
                d_273_v140_ = ((d_273_v140_) + (_dafny.SeqWithoutIsStrInference([d_272_v139_ for d_274_i16_ in range(default__.abs(512))]))).set(default__.safeIndex((d_99_v3_ if (d_159_v41_).f30 else (d_160_v42_).f31), len((d_273_v140_) + (_dafny.SeqWithoutIsStrInference([d_272_v139_ for d_275_i16_ in range(default__.abs(512))])))), D2_DC6(d_102_v6_))
                d_276_v141_: _dafny.MultiSet
                d_276_v141_ = _dafny.MultiSet([(d_160_v42_).f31])
                d_277_v142_: bool
                out12_: bool
                out12_ = (d_159_v41_).m1(d_276_v141_, (d_159_v41_).f30, d_98_v2_, d_103_globalState_)
                d_277_v142_ = out12_
                pat_let_tv7_ = d_160_v42_
                pat_let_tv8_ = d_160_v42_
                def iife23_(_pat_let7_0):
                    def iife24_(d_278_dt__update__tmp_h0_):
                        def iife25_(_pat_let8_0):
                            def iife26_(d_279_dt__update_hcf24_h0_):
                                def iife27_(_pat_let9_0):
                                    def iife28_(d_280_dt__update_hcf23_h0_):
                                        return D6_DC14(d_280_dt__update_hcf23_h0_, d_279_dt__update_hcf24_h0_, (d_278_dt__update__tmp_h0_).cf25)
                                    return iife28_(_pat_let9_0)
                                return iife27_((pat_let_tv8_).f31)
                            return iife26_(_pat_let8_0)
                        return iife25_((pat_let_tv7_).f31)
                    return iife24_(_pat_let7_0)
                d_99_v3_ = (iife23_(default__.fm41(d_277_v142_, True, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], d_103_globalState_))).cf23
            elif True:
                d_102_v6_ = (D2_DC6(d_102_v6_)).cf10
                index38_ = default__.safeIndex(810, (d_104_v7_).length(0))
                (d_104_v7_)[index38_] = d_98_v2_
                d_281_v143_: C1
                nw54_ = C1()
                nw54_.ctor__((d_159_v41_).f30)
                d_281_v143_ = nw54_
                d_282_v144_: D10
                d_282_v144_ = D10_DC24(547, d_281_v143_, d_154_v36_)
                d_283_v145_: _dafny.MultiSet
                d_283_v145_ = _dafny.MultiSet([len((d_159_v41_).f29)])
                d_284_v146_: _dafny.MultiSet
                d_284_v146_ = _dafny.MultiSet([d_154_v36_])
                d_285_v147_: _dafny.Seq
                d_285_v147_ = _dafny.SeqWithoutIsStrInference([d_160_v42_])
                d_286_v148_: D19
                d_286_v148_ = D19_DC43((d_285_v147_)[default__.safeIndex((d_160_v42_).f31, len(d_285_v147_))], (d_160_v42_).f31, (d_281_v143_).f32)
                d_287_v149_: _dafny.MultiSet
                d_287_v149_ = _dafny.MultiSet([True, d_98_v2_])
                d_288_v150_: D6
                d_288_v150_ = D6_DC13((d_160_v42_).f31, (d_160_v42_).f31, _dafny.Set({(d_287_v149_).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygqeqe"))), (d_160_v42_).f31}))
                d_289_v151_: _dafny.Array
                nw55_ = _dafny.Array(None, 10)
                nw55_[int(0)] = (d_159_v41_).fm6((d_159_v41_).f30, d_98_v2_, (d_160_v42_).f31, d_103_globalState_)
                nw55_[int(1)] = 969
                nw55_[int(2)] = len(d_97_v1_)
                nw55_[int(3)] = (d_160_v42_).f31
                nw55_[int(4)] = (d_160_v42_).f31
                nw55_[int(5)] = (d_288_v150_).cf20
                nw55_[int(6)] = default__.fm2((d_160_v42_).f31, 421, d_103_globalState_)
                nw55_[int(7)] = (d_160_v42_).f31
                nw55_[int(8)] = (d_160_v42_).f31
                nw55_[int(9)] = (d_160_v42_).f31
                d_289_v151_ = nw55_
                d_290_v152_: _dafny.Set
                d_290_v152_ = _dafny.Set({d_289_v151_})
                d_291_v153_: T2
                nw56_ = C0()
                nw56_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjpy")), False, d_99_v3_, d_104_v7_, (d_281_v143_).f32)
                d_291_v153_ = nw56_
                d_292_v154_: _dafny.Array
                nw57_ = _dafny.Array(None, 24)
                nw57_[int(0)] = (d_160_v42_).f31
                nw57_[int(1)] = (d_160_v42_).f31
                nw57_[int(2)] = (d_160_v42_).f31
                nw57_[int(3)] = len(_dafny.SeqWithoutIsStrInference([d_154_v36_ for d_293_i17_ in range(default__.abs(776))]))
                nw57_[int(4)] = (d_160_v42_).f31
                nw57_[int(5)] = (d_160_v42_).f31
                nw57_[int(6)] = (d_282_v144_).cf36
                nw57_[int(7)] = 187
                nw57_[int(8)] = (0) - ((d_160_v42_).f31)
                nw57_[int(9)] = default__.safeModuloInt((d_160_v42_).f31, (d_160_v42_).f31)
                nw57_[int(10)] = (d_160_v42_).f31
                nw57_[int(11)] = d_99_v3_
                nw57_[int(12)] = default__.safeDivisionInt((d_160_v42_).f31, (d_160_v42_).f31)
                nw57_[int(13)] = (d_160_v42_).f31
                nw57_[int(14)] = d_99_v3_
                nw57_[int(15)] = (0) - ((d_283_v145_).cardinality)
                nw57_[int(16)] = (d_99_v3_ if False else len(d_152_v34_))
                nw57_[int(17)] = ((d_284_v146_).cardinality if (d_281_v143_).f32 else (d_160_v42_).f31)
                nw57_[int(18)] = 979
                nw57_[int(19)] = (d_160_v42_).f31
                nw57_[int(20)] = (0) - ((d_99_v3_ if (d_281_v143_).f32 else (d_160_v42_).f31))
                nw57_[int(21)] = (d_286_v148_).cf71
                nw57_[int(22)] = len((d_290_v152_) | (d_290_v152_))
                nw57_[int(23)] = len(_dafny.Set({d_291_v153_, d_291_v153_}))
                d_292_v154_ = nw57_
                index39_ = default__.safeIndex(706, (d_289_v151_).length(0))
                (d_289_v151_)[index39_] = (d_160_v42_).f31
                d_294_v155_: C3
                nw58_ = C3()
                nw58_.ctor__(default__.fm1(d_103_globalState_), d_262_v131_, d_262_v131_, 56)
                d_294_v155_ = nw58_
                d_295_v156_: _dafny.Map
                d_295_v156_ = _dafny.Map({d_99_v3_: (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]})
                index40_ = default__.safeIndex(706, (d_289_v151_).length(0))
                rhs39_ = (len(d_295_v156_)) - ((d_160_v42_).f31)
                rhs40_ = d_294_v155_
                rhs41_ = (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]
                lhs36_ = d_289_v151_
                lhs37_ = default__.safeIndex(706, (d_289_v151_).length(0))
                lhs38_ = d_103_globalState_
                lhs36_[lhs37_] = rhs39_
                d_294_v155_ = rhs40_
                lhs38_.f1 = rhs41_
                d_296_v157_: _dafny.Set
                d_296_v157_ = _dafny.Set({(d_281_v143_).f32, (d_281_v143_).f32})
                d_297_v158_: C2
                nw59_ = C2()
                nw59_.ctor__((d_289_v151_)[default__.safeIndex(706, (d_289_v151_).length(0))], (d_281_v143_).fm14(d_296_v157_, (d_289_v151_)[default__.safeIndex(706, (d_289_v151_).length(0))], d_103_globalState_), (d_289_v151_)[default__.safeIndex(706, (d_289_v151_).length(0))])
                d_297_v158_ = nw59_
                d_100_v4_ = d_100_v4_
            d_298_v159_: _dafny.Array
            def lambda14_(d_299_v36_):
                def lambda15_(d_300_i18_):
                    return _dafny.SeqWithoutIsStrInference([d_299_v36_ for d_301_i19_ in range(default__.abs(-591))])

                return lambda15_

            init7_ = lambda14_(d_154_v36_)
            nw60_ = _dafny.Array(None, 26)
            for i0_7_ in range(nw60_.length(0)):
                nw60_[i0_7_] = init7_(i0_7_)
            d_298_v159_ = nw60_
            d_302_v160_: _dafny.Seq
            d_302_v160_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_154_v36_ for d_303_i20_ in range(default__.abs(-774))])])
            d_304_v161_: _dafny.MultiSet
            d_304_v161_ = _dafny.MultiSet([(d_160_v42_).f31])
            d_305_v162_: T1
            nw61_ = C0()
            nw61_.ctor__(d_97_v1_, d_262_v131_, (d_304_v161_).cardinality, d_104_v7_, (d_159_v41_).f30)
            d_305_v162_ = nw61_
            d_306_v163_: _dafny.Seq
            d_306_v163_ = _dafny.SeqWithoutIsStrInference([d_305_v162_])
            index41_ = default__.safeIndex(929, (d_298_v159_).length(0))
            (d_298_v159_)[index41_] = (d_97_v1_) + ((d_302_v160_)[default__.safeIndex(len(d_306_v163_), len(d_302_v160_))])
            index42_ = default__.safeIndex(929, (d_298_v159_).length(0))
            (d_298_v159_)[index42_] = (d_159_v41_).f29
            d_99_v3_ = (d_160_v42_).f31
            d_307_v164_: _dafny.Map
            d_307_v164_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrd")): d_98_v2_})
            d_307_v164_ = (d_307_v164_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdm")), d_262_v131_)
        d_308_v165_: _dafny.Set
        d_308_v165_ = _dafny.Set({(d_159_v41_).f30})
        if (d_308_v165_).issubset(d_308_v165_):
            d_309_v166_: _dafny.Array
            nw62_ = _dafny.Array(None, 1)
            nw62_[int(0)] = default__.safeModuloInt((d_160_v42_).f31, (d_160_v42_).f31)
            d_309_v166_ = nw62_
            index43_ = default__.safeIndex(641, (d_309_v166_).length(0))
            (d_309_v166_)[index43_] = -575
            d_310_v167_: _dafny.MultiSet
            d_310_v167_ = _dafny.MultiSet([d_99_v3_])
            index44_ = default__.safeIndex(641, (d_309_v166_).length(0))
            (d_309_v166_)[index44_] = (-461) - ((d_310_v167_).cardinality)
            d_311_v168_: _dafny.Seq
            d_311_v168_ = _dafny.SeqWithoutIsStrInference([d_102_v6_, d_102_v6_, d_102_v6_, d_102_v6_, d_102_v6_])
            d_312_v169_: _dafny.Map
            d_312_v169_ = _dafny.Map({80: d_98_v2_})
            d_313_v170_: D12
            d_313_v170_ = D12_DC29(381, d_98_v2_, d_311_v168_, len(_dafny.Map({d_98_v2_: len(d_312_v169_)})), -170)
            source3_ = (d_313_v170_ if d_98_v2_ else d_313_v170_)
            if source3_.is_DC29:
                d_314___mcc_h6_ = source3_.cf45
                d_315___mcc_h7_ = source3_.cf46
                d_316___mcc_h8_ = source3_.cf47
                d_317___mcc_h9_ = source3_.cf48
                d_318___mcc_h10_ = source3_.cf49
                d_319_cf49_ = d_318___mcc_h10_
                d_320_cf48_ = d_317___mcc_h9_
                d_321_cf47_ = d_316___mcc_h8_
                d_322_cf46_ = d_315___mcc_h7_
                d_323_cf45_ = d_314___mcc_h6_
                d_324_v171_: _dafny.Seq
                d_324_v171_ = _dafny.SeqWithoutIsStrInference([d_152_v34_, d_152_v34_, d_152_v34_])
                d_152_v34_ = (d_152_v34_) + ((d_324_v171_)[default__.safeIndex(d_323_cf45_, len(d_324_v171_))])
                d_325_v172_: C3
                nw63_ = C3()
                nw63_.ctor__((d_159_v41_).f29, d_322_cf46_, (d_159_v41_).f30, d_323_cf45_)
                d_325_v172_ = nw63_
                d_97_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kh"))
                index45_ = default__.safeIndex(810, (d_104_v7_).length(0))
                (d_104_v7_)[index45_] = (d_159_v41_).f30
            elif True:
                d_326___mcc_h11_ = source3_.cf44
                d_327_cf44_ = d_326___mcc_h11_
                d_328_v173_: _dafny.MultiSet
                d_328_v173_ = _dafny.MultiSet([not((d_159_v41_).f30)])
                index46_ = default__.safeIndex(810, (d_104_v7_).length(0))
                index47_ = default__.safeIndex(810, (d_104_v7_).length(0))
                rhs42_ = (d_98_v2_) and ((d_159_v41_).f30)
                rhs43_ = default__.safeModuloInt((d_160_v42_).fm6(True, True, (0) - ((d_160_v42_).f31), d_103_globalState_), (d_153_v35_)[default__.safeIndex(len(d_97_v1_), len(d_153_v35_))])
                rhs44_ = not(d_98_v2_)
                rhs45_ = (d_328_v173_).isdisjoint((d_328_v173_).intersection(d_328_v173_))
                rhs46_ = d_162_v44_
                lhs39_ = d_104_v7_
                lhs40_ = default__.safeIndex(810, (d_104_v7_).length(0))
                lhs41_ = d_103_globalState_
                lhs42_ = d_104_v7_
                lhs43_ = default__.safeIndex(810, (d_104_v7_).length(0))
                lhs44_ = d_103_globalState_
                lhs39_[lhs40_] = rhs42_
                lhs41_.f0 = rhs43_
                lhs42_[lhs43_] = rhs44_
                lhs44_.f1 = rhs45_
                d_162_v44_ = rhs46_
                d_329_v174_: _dafny.Array
                nw64_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_329_v174_ = nw64_
                d_330_v175_: _dafny.Map
                d_330_v175_ = _dafny.Map({d_104_v7_: d_329_v174_})
                index48_ = default__.safeIndex(641, (d_309_v166_).length(0))
                index49_ = default__.safeIndex(810, (d_104_v7_).length(0))
                rhs47_ = (d_160_v42_).f31
                rhs48_ = ((d_160_v42_).f31) <= (d_99_v3_)
                rhs49_ = ((d_330_v175_)[d_104_v7_] if (d_104_v7_) in (d_330_v175_) else d_329_v174_)
                rhs50_ = default__.fm21(default__.fm1(d_103_globalState_), (d_160_v42_).f31, 82, d_103_globalState_)
                rhs51_ = (d_159_v41_).f30
                lhs45_ = d_309_v166_
                lhs46_ = default__.safeIndex(641, (d_309_v166_).length(0))
                lhs47_ = d_104_v7_
                lhs48_ = default__.safeIndex(810, (d_104_v7_).length(0))
                lhs45_[lhs46_] = rhs47_
                lhs47_[lhs48_] = rhs48_
                d_329_v174_ = rhs49_
                d_310_v167_ = rhs50_
                d_98_v2_ = rhs51_
                d_331_v176_: bool
                out13_: bool
                out13_ = (d_160_v42_).m1(d_310_v167_, ((d_328_v173_).cardinality) >= ((d_160_v42_).f31), (d_159_v41_).f30, d_103_globalState_)
                d_331_v176_ = out13_
                d_98_v2_ = (d_159_v41_).f30
            d_332_v177_: C1
            nw65_ = C1()
            nw65_.ctor__(False)
            d_332_v177_ = nw65_
            d_333_v178_: _dafny.Map
            d_333_v178_ = _dafny.Map({d_332_v177_: (d_159_v41_ if (d_332_v177_).f32 else d_159_v41_)})
            d_334_v179_: _dafny.Seq
            d_334_v179_ = _dafny.SeqWithoutIsStrInference([d_333_v178_])
            d_335_v180_: T2
            nw66_ = C0()
            nw66_.ctor__((d_159_v41_).f29, (d_159_v41_).f30, (0) - (len((d_159_v41_).f29)), d_104_v7_, (d_159_v41_).f30)
            d_335_v180_ = nw66_
            rhs52_ = (d_334_v179_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_335_v180_, d_335_v180_, d_335_v180_, d_335_v180_])), len(d_334_v179_))]
            rhs53_ = d_104_v7_
            rhs54_ = (d_160_v42_).fm6(not(((d_160_v42_).f31) <= (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhjne"))))), (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], default__.safeModuloInt(((d_310_v167_)[(d_160_v42_).f31] if ((d_160_v42_).f31) in (d_310_v167_) else (d_309_v166_)[default__.safeIndex(641, (d_309_v166_).length(0))]), (d_309_v166_)[default__.safeIndex(641, (d_309_v166_).length(0))]), d_103_globalState_)
            rhs55_ = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([283, 161, 353])), (d_160_v42_).f31)
            lhs49_ = d_103_globalState_
            lhs50_ = d_103_globalState_
            d_333_v178_ = rhs52_
            d_104_v7_ = rhs53_
            lhs49_.f2 = rhs54_
            lhs50_.f0 = rhs55_
            index50_ = default__.safeIndex(810, (d_104_v7_).length(0))
            rhs56_ = (d_98_v2_) not in (_dafny.MultiSet([d_335_v180_.f34, (d_159_v41_).f30]))
            rhs57_ = (d_332_v177_).f32
            rhs58_ = d_309_v166_
            rhs59_ = (d_160_v42_).f31
            rhs60_ = (d_159_v41_).f29
            lhs51_ = d_335_v180_
            lhs52_ = d_104_v7_
            lhs53_ = default__.safeIndex(810, (d_104_v7_).length(0))
            lhs51_.f34 = rhs56_
            lhs52_[lhs53_] = rhs57_
            d_309_v166_ = rhs58_
            d_99_v3_ = rhs59_
            d_97_v1_ = rhs60_
            d_336_v181_: _dafny.MultiSet
            d_336_v181_ = _dafny.MultiSet([(d_159_v41_).f30, d_98_v2_])
            if not((((d_336_v181_)[d_98_v2_] if (d_98_v2_) in (d_336_v181_) else d_99_v3_)) == ((d_160_v42_).f31)):
                d_337_v182_: C4
                nw67_ = C4()
                nw67_.ctor__((d_160_v42_).f31, d_97_v1_, (d_332_v177_).f32, (d_160_v42_).f31)
                d_337_v182_ = nw67_
                d_338_v183_: _dafny.Map
                d_338_v183_ = _dafny.Map({d_337_v182_: (d_160_v42_).f31})
                d_339_v184_: _dafny.Map
                d_339_v184_ = _dafny.Map({(d_159_v41_).f30: ((d_312_v169_)[(d_337_v182_).f27] if ((d_337_v182_).f27) in (d_312_v169_) else (d_159_v41_).f30)})
                d_340_v185_: D14
                d_340_v185_ = D14_DC32((d_159_v41_).f30, d_99_v3_, d_339_v184_, (d_159_v41_).f29, (d_332_v177_).f32)
                d_341_v186_: C4
                nw68_ = C4()
                nw68_.ctor__(((d_338_v183_)[d_337_v182_] if (d_337_v182_) in (d_338_v183_) else (d_160_v42_).f31), (d_340_v185_).cf55, (d_332_v177_).f32, len(d_100_v4_))
                d_341_v186_ = nw68_
                d_342_v187_: _dafny.Array
                nw69_ = _dafny.Array(None, 27)
                nw69_[int(0)] = d_309_v166_
                nw69_[int(1)] = d_309_v166_
                nw69_[int(2)] = d_309_v166_
                nw69_[int(3)] = (d_309_v166_ if (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))] else d_309_v166_)
                nw69_[int(4)] = d_309_v166_
                nw69_[int(5)] = d_309_v166_
                nw69_[int(6)] = (d_309_v166_ if True else d_309_v166_)
                nw69_[int(7)] = d_309_v166_
                nw69_[int(8)] = d_309_v166_
                nw69_[int(9)] = d_309_v166_
                nw69_[int(10)] = d_309_v166_
                nw69_[int(11)] = d_309_v166_
                nw69_[int(12)] = d_309_v166_
                nw69_[int(13)] = d_309_v166_
                nw69_[int(14)] = d_309_v166_
                nw69_[int(15)] = d_309_v166_
                nw69_[int(16)] = d_309_v166_
                nw69_[int(17)] = d_309_v166_
                nw69_[int(18)] = d_309_v166_
                nw69_[int(19)] = d_309_v166_
                nw69_[int(20)] = d_309_v166_
                nw69_[int(21)] = d_309_v166_
                nw69_[int(22)] = d_309_v166_
                nw69_[int(23)] = d_309_v166_
                nw69_[int(24)] = d_309_v166_
                nw69_[int(25)] = d_309_v166_
                nw69_[int(26)] = d_309_v166_
                d_342_v187_ = nw69_
                index51_ = default__.safeIndex(641, (d_309_v166_).length(0))
                rhs61_ = (d_152_v34_)[default__.safeIndex((d_160_v42_).f31, len(d_152_v34_))]
                rhs62_ = d_99_v3_
                rhs63_ = d_342_v187_
                rhs64_ = len((d_153_v35_).set(default__.safeIndex((d_309_v166_)[default__.safeIndex(641, (d_309_v166_).length(0))], len(d_153_v35_)), (d_160_v42_).f31))
                lhs54_ = d_335_v180_
                lhs55_ = d_309_v166_
                lhs56_ = default__.safeIndex(641, (d_309_v166_).length(0))
                lhs57_ = d_103_globalState_
                lhs54_.f34 = rhs61_
                lhs55_[lhs56_] = rhs62_
                d_342_v187_ = rhs63_
                lhs57_.f0 = rhs64_
                d_100_v4_ = (d_100_v4_).set((d_159_v41_).f30, -43)
                (d_103_globalState_).f2 = (d_160_v42_).f31
                d_102_v6_ = d_102_v6_
            elif True:
                d_343_v188_: _dafny.Map
                d_343_v188_ = _dafny.Map({d_154_v36_: d_104_v7_})
                d_344_v189_: _dafny.MultiSet
                d_344_v189_ = _dafny.MultiSet([(d_343_v188_) | (_dafny.Map({d_154_v36_: (d_335_v180_).f33}))])
                (d_103_globalState_).f0 = ((d_344_v189_)[d_343_v188_] if (d_343_v188_) in (d_344_v189_) else (d_160_v42_).f31)
                d_97_v1_ = ((d_159_v41_).f29) + (((d_159_v41_).f29) + ((d_159_v41_).f29))
                (d_103_globalState_).f0 = (d_160_v42_).f31
                index52_ = default__.safeIndex(641, (d_309_v166_).length(0))
                (d_309_v166_)[index52_] = default__.safeModuloInt(((d_160_v42_).f31) * ((d_160_v42_).f31), d_99_v3_)
                index53_ = default__.safeIndex(641, (d_309_v166_).length(0))
                (d_309_v166_)[index53_] = (d_309_v166_)[default__.safeIndex(641, (d_309_v166_).length(0))]
        elif True:
            d_153_v35_ = ((default__.fm3((d_160_v42_).f31, (d_159_v41_).f30, (d_160_v42_).f31, d_98_v2_, d_103_globalState_)) + (d_153_v35_)) + (d_153_v35_)
            d_345_v190_: _dafny.Map
            d_345_v190_ = _dafny.Map({(d_152_v34_)[default__.safeIndex(len(d_152_v34_), len(d_152_v34_))]: True})
            d_346_v191_: _dafny.MultiSet
            d_346_v191_ = _dafny.MultiSet([d_345_v190_, d_345_v190_, (d_345_v190_) | (d_345_v190_), (d_345_v190_).set(not((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]), True), (d_345_v190_).set((d_159_v41_).f30, (d_159_v41_).f30)])
            d_346_v191_ = ((_dafny.MultiSet([d_345_v190_])) - (_dafny.MultiSet([d_345_v190_, d_345_v190_, d_345_v190_]))) - (_dafny.MultiSet([d_345_v190_]))
            d_347_v192_: _dafny.Seq
            d_347_v192_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqwbfped")), (d_159_v41_).f29, d_97_v1_, d_97_v1_])
            d_347_v192_ = (d_347_v192_) + (d_347_v192_)
            d_348_v193_: _dafny.Map
            d_348_v193_ = _dafny.Map({d_153_v35_: d_104_v7_})
            d_348_v193_ = d_348_v193_
            d_349_v194_: _dafny.Array
            nw70_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_349_v194_ = nw70_
            d_350_v195_: _dafny.Array
            nw71_ = _dafny.Array(None, 2)
            nw71_[int(0)] = len(d_345_v190_)
            nw71_[int(1)] = d_99_v3_
            d_350_v195_ = nw71_
            index54_ = default__.safeIndex(194, (d_349_v194_).length(0))
            (d_349_v194_)[index54_] = d_350_v195_
            d_351_v196_: _dafny.Map
            d_351_v196_ = _dafny.Map({((d_160_v42_).f31) + (d_99_v3_): (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]})
            d_352_v197_: _dafny.Seq
            d_352_v197_ = _dafny.SeqWithoutIsStrInference([d_152_v34_])
            index55_ = default__.safeIndex(194, (d_349_v194_).length(0))
            rhs65_ = d_350_v195_
            rhs66_ = (d_350_v195_ if ((d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))] if (d_159_v41_).f30 else (d_159_v41_).f30) else d_350_v195_)
            rhs67_ = default__.fm42((default__.fm3(len(_dafny.SeqWithoutIsStrInference([(d_160_v42_).f31])), (d_159_v41_).f30, (d_160_v42_).f31, (d_159_v41_).f30, d_103_globalState_)) + ((d_153_v35_).set(default__.safeIndex((d_160_v42_).f31, len(d_153_v35_)), (d_160_v42_).f31)), d_98_v2_, d_352_v197_, d_103_globalState_)
            lhs58_ = d_349_v194_
            lhs59_ = default__.safeIndex(194, (d_349_v194_).length(0))
            lhs58_[lhs59_] = rhs65_
            d_350_v195_ = rhs66_
            d_351_v196_ = rhs67_
        d_353_v198_: _dafny.Map
        d_353_v198_ = _dafny.Map({d_97_v1_: (d_159_v41_).f29})
        d_354_v199_: _dafny.MultiSet
        d_354_v199_ = _dafny.MultiSet([(d_160_v42_).fm11(d_353_v198_, len((d_159_v41_).f29), (d_160_v42_).f31, d_103_globalState_), d_99_v3_, (d_160_v42_).f31, (d_160_v42_).f31])
        d_354_v199_ = (d_354_v199_) | (d_354_v199_)
        d_355_v200_: _dafny.Map
        d_355_v200_ = _dafny.Map({D0_DC0(False): d_99_v3_})
        index56_ = default__.safeIndex(810, (d_104_v7_).length(0))
        (d_104_v7_)[index56_] = not(default__.fm5((_dafny.MultiSet([(d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))], not((d_159_v41_).f30), (d_159_v41_).f30, (d_104_v7_)[default__.safeIndex(810, (d_104_v7_).length(0))]])).isdisjoint(_dafny.MultiSet([(d_159_v41_).f30])), d_355_v200_, d_99_v3_, d_103_globalState_))
        d_356_v201_: _dafny.Map
        d_356_v201_ = _dafny.Map({d_104_v7_: (d_152_v34_) + (d_152_v34_)})
        d_356_v201_ = (d_356_v201_).set(d_104_v7_, (_dafny.SeqWithoutIsStrInference([(d_159_v41_).f30])) + (d_152_v34_))
        d_357_v202_: _dafny.Array
        def lambda16_(d_358_v3_):
            def lambda17_(d_359_i21_):
                return default__.safeModuloInt(d_359_i21_, d_358_v3_)

            return lambda17_

        init8_ = lambda16_(d_99_v3_)
        nw72_ = _dafny.Array(None, 21)
        for i0_8_ in range(nw72_.length(0)):
            nw72_[i0_8_] = init8_(i0_8_)
        d_357_v202_ = nw72_
        index57_ = default__.safeIndex(95, (d_357_v202_).length(0))
        (d_357_v202_)[index57_] = 923
        index58_ = default__.safeIndex(95, (d_357_v202_).length(0))
        rhs68_ = d_99_v3_
        rhs69_ = (d_97_v1_)[default__.safeIndex(default__.safeModuloInt((d_160_v42_).f31, (d_160_v42_).f31), len(d_97_v1_))]
        rhs70_ = (224) > (d_99_v3_)
        rhs71_ = d_152_v34_
        rhs72_ = (d_160_v42_).f31
        lhs60_ = d_357_v202_
        lhs61_ = default__.safeIndex(95, (d_357_v202_).length(0))
        d_99_v3_ = rhs68_
        d_154_v36_ = rhs69_
        d_98_v2_ = rhs70_
        d_152_v34_ = rhs71_
        lhs60_[lhs61_] = rhs72_
        d_360_v203_: _dafny.Array
        nw73_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
        d_360_v203_ = nw73_
        index59_ = default__.safeIndex(652, (d_360_v203_).length(0))
        (d_360_v203_)[index59_] = ((d_159_v41_).f29) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cd")))
        index60_ = default__.safeIndex(652, (d_360_v203_).length(0))
        (d_360_v203_)[index60_] = d_97_v1_
        _dafny.print((d_97_v1_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_98_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_99_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v4_) == (_dafny.Map({False: -43}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v5_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: 189}), _dafny.Map({False: 189})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v6_) == (_dafny.Set({189}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_103_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_103_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_103_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_103_globalState_).f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_103_globalState_).f9) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: 189}), _dafny.Map({False: 189})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_103_globalState_).f12).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_103_globalState_).f13) == (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxiyve"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_103_globalState_).f17) == (_dafny.Set({189}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_108_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v34_) == (_dafny.SeqWithoutIsStrInference([True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v35_) == (_dafny.SeqWithoutIsStrInference([9]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_154_v36_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v38_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_v39_) == (_dafny.SeqWithoutIsStrInference([D0_DC0(True), D0_DC0(False), D0_DC0(True)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v40_) == (_dafny.Map({_dafny.CodePoint('r'): _dafny.Map({D0_DC0(True): 1, D0_DC0(False): 1})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_159_v41_).f29).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v41_).f30))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v42_).f31))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_161_v43_)[10])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_163_v45_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_249_v118_) == (_dafny.Map({True: _dafny.CodePoint('r')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_308_v165_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_353_v198_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oduxsblwp")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oduxsblwp"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_354_v199_) == (_dafny.MultiSet([991, 991, -189, -189, -189, -189, -189, -189]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_355_v200_) == (_dafny.Map({D0_DC0(False): -189}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_356_v201_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v202_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_360_v203_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0), False)
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {self.cf2.VerbatimString(True)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2)
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
        return lambda: D1_DC4(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(_dafny.CodePoint('D'), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC10(D4, NamedTuple('DC10', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)

class D5_DC11(D5, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC13(int(0), int(0), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)

class D6_DC13(D6, NamedTuple('DC13', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC14(D6, NamedTuple('DC14', [('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC12(D6, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)

class D7_DC15(D7, NamedTuple('DC15', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC17(D0.default()(), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D8_DC16)

class D8_DC17(D8, NamedTuple('DC17', [('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC16(D8, NamedTuple('DC16', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC16({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC16) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC19(D6.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D9_DC18)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)

class D9_DC19(D9, NamedTuple('DC19', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC18(D9, NamedTuple('DC18', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC18({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC18) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC20(D9, NamedTuple('DC20', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC22(_dafny.Array(None, 0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)

class D10_DC22(D10, NamedTuple('DC22', [('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({_dafny.string_of(self.cf34)}, {self.cf35.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC23(D10, NamedTuple('DC23', [])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23)
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC24(D10, NamedTuple('DC24', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC21(D10, NamedTuple('DC21', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC26(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)

class D11_DC26(D11, NamedTuple('DC26', [('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC25(D11, NamedTuple('DC25', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC29(int(0), False, _dafny.Seq({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)

class D12_DC29(D12, NamedTuple('DC29', [('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC28(D12, NamedTuple('DC28', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D13_DC30)

class D13_DC30(D13, NamedTuple('DC30', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC30({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC30) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC32(False, int(0), _dafny.Map({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)

class D14_DC32(D14, NamedTuple('DC32', [('cf52', Any), ('cf53', Any), ('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {self.cf55.VerbatimString(True)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC31(D14, NamedTuple('DC31', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC34(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)

class D15_DC34(D15, NamedTuple('DC34', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC33(D15, NamedTuple('DC33', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC36(False, _dafny.Seq({}), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D16_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D16_DC35)

class D16_DC36(D16, NamedTuple('DC36', [('cf60', Any), ('cf61', Any), ('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC36({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC36) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC35(D16, NamedTuple('DC35', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC38()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D17_DC38)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D17_DC37)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D17_DC39)

class D17_DC38(D17, NamedTuple('DC38', [])):
    def __dafnystr__(self) -> str:
        return f'D17.DC38'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC38)
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC37(D17, NamedTuple('DC37', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC37({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC37) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC39(D17, NamedTuple('DC39', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC39({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC39) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D18_DC40)

class D18_DC40(D18, NamedTuple('DC40', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC40({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC40) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC42(_dafny.MultiSet({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D19_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D19_DC43)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D19_DC41)

class D19_DC42(D19, NamedTuple('DC42', [('cf67', Any), ('cf68', Any), ('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC42({_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC42) and self.cf67 == __o.cf67 and self.cf68 == __o.cf68 and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC43(D19, NamedTuple('DC43', [('cf70', Any), ('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC43({_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC43) and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC41(D19, NamedTuple('DC41', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC41({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC41) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC45(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D20_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D20_DC46)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D20_DC47)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D20_DC44)

class D20_DC45(D20, NamedTuple('DC45', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC45({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC45) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC46(D20, NamedTuple('DC46', [('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC46({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC46) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC47(D20, NamedTuple('DC47', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC47'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC47)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC44(D20, NamedTuple('DC44', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC44({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC44) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC49(False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D21_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D21_DC48)

class D21_DC49(D21, NamedTuple('DC49', [('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC49({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC49) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC48(D21, NamedTuple('DC48', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC48({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC48) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC51(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D22_DC51)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D22_DC50)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D22_DC52)

class D22_DC51(D22, NamedTuple('DC51', [('cf81', Any), ('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC51({_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC51) and self.cf81 == __o.cf81 and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC50(D22, NamedTuple('DC50', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC50({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC50) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC52(D22, NamedTuple('DC52', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC52({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC52) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, p2, globalState):
        pass


class T1:
    pass

class T2:
    pass
    @property
    def f34(self):
        return self._f34
    @f34.setter
    def f34(self, value):
        self._f34 = value

class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f1: bool = False
        self.f2: int = int(0)
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f4: bool = False
        self._f5: int = int(0)
        self._f6: int = int(0)
        self._f7: bool = False
        self._f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f9: _dafny.Seq = _dafny.Seq({})
        self._f10: int = int(0)
        self._f11: bool = False
        self._f12: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f13: _dafny.Map = _dafny.Map({})
        self._f14: int = int(0)
        self._f15: int = int(0)
        self._f16: int = int(0)
        self._f17: _dafny.Set = _dafny.Set({})
        self._f18: bool = False
        self._f19: int = int(0)
        self._f20: bool = False
        self._f21: bool = False
        self._f22: bool = False
        self._f23: bool = False
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24):
        (self).f0 = f0
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
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23
        (self)._f24 = f24

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

class C0(T0, T1, T2):
    def  __init__(self):
        self._f34: bool = False
        self._f25: bool = False
        self._f26: int = int(0)
        self._f33: _dafny.Array = _dafny.Array(None, 0)
        self.f35: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f34(self):
        return self._f34
    @f34.setter
    def f34(self, value):
        self._f34 = value
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    @property
    def f33(self):
        return self._f33
    def ctor__(self, f35, f25, f26, f33, f34):
        (self).f35 = f35
        (self)._f25 = f25
        (self)._f26 = f26
        (self)._f33 = f33
        (self).f34 = f34

    def fm6(self, p0, p1, p2, globalState):
        if (_dafny.MultiSet([_dafny.Set({(self).f25, not(self.f34)})])).isdisjoint(_dafny.MultiSet([_dafny.Set({(self).f25, self.f34})])):
            return (self).f26
        elif True:
            return (self).f26

    def fm7(self, globalState):
        return (self).f25

    def fm16(self, globalState):
        return (162) + ((self).f26)

    def fm17(self, p0, p1, globalState):
        return ((self).f26) != (((self).f26) - ((self).f26))

    def fm18(self, p0, p1, p2, globalState):
        def iife29_():
            coll9_ = _dafny.Map()
            compr_9_: _dafny.Seq
            for compr_9_ in ((_dafny.MultiSet([self.f35, self.f35, self.f35, self.f35, self.f35])) - (_dafny.MultiSet([self.f35, self.f35, self.f35]))).Elements:
                d_361_v0_: _dafny.Seq = compr_9_
                if (d_361_v0_) in ((_dafny.MultiSet([self.f35, self.f35, self.f35, self.f35, self.f35])) - (_dafny.MultiSet([self.f35, self.f35, self.f35]))):
                    coll9_[d_361_v0_] = (self).f25
            return _dafny.Map(coll9_)
        return iife29_()
        

    def fm19(self, p0, p1, p2, globalState):
        def iife30_():
            coll10_ = _dafny.Map()
            compr_10_: str
            for compr_10_ in (self.f35).Elements:
                d_362_v0_: str = compr_10_
                if (d_362_v0_) in (self.f35):
                    coll10_[d_362_v0_] = self.f34
            return _dafny.Map(coll10_)
        return ((_dafny.MultiSet([len(iife30_()
        )])).intersection((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f26]))))).cardinality

    def fm20(self, p0, globalState):
        source4_ = D2_DC7((self).f25, True, (self).f26)
        if source4_.is_DC7:
            d_363___mcc_h0_ = source4_.cf11
            d_364___mcc_h1_ = source4_.cf12
            d_365___mcc_h2_ = source4_.cf13
            d_366_cf13_ = d_365___mcc_h2_
            d_367_cf12_ = d_364___mcc_h1_
            d_368_cf11_ = d_363___mcc_h0_
            return _dafny.CodePoint('m')
        elif True:
            d_369___mcc_h3_ = source4_.cf10
            d_370_cf10_ = d_369___mcc_h3_
            return _dafny.CodePoint('w')

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        (globalState).f2 = (self).f26
        d_371_v0_: _dafny.MultiSet
        d_371_v0_ = _dafny.MultiSet([self.f35, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_372_i0_ in range(default__.abs(449))])])
        (globalState).f2 = ((d_371_v0_)[(self.f35) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfhbn")))] if ((self.f35) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfhbn")))) in (d_371_v0_) else ((self).f26 if p1 else (p0).cardinality))
        d_373_v1_: _dafny.Array
        nw74_ = _dafny.Array(_dafny.Seq({}), 8)
        d_373_v1_ = nw74_
        d_374_v2_: _dafny.Seq
        d_374_v2_ = _dafny.SeqWithoutIsStrInference([(self).f33])
        index61_ = default__.safeIndex(429, (d_373_v1_).length(0))
        (d_373_v1_)[index61_] = d_374_v2_
        d_375_v3_: _dafny.Array
        nw75_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
        d_375_v3_ = nw75_
        index62_ = default__.safeIndex(289, (d_375_v3_).length(0))
        (d_375_v3_)[index62_] = self.f35
        d_376_v4_: _dafny.Array
        nw76_ = _dafny.Array(int(0), 23)
        d_376_v4_ = nw76_
        d_377_v5_: _dafny.Map
        d_377_v5_ = _dafny.Map({d_376_v4_: default__.safeModuloInt((self).f26, 574)})
        d_378_v6_: _dafny.Map
        d_378_v6_ = _dafny.Map({(self).f25: (d_374_v2_) + (((d_374_v2_).set(default__.safeIndex(396, len(d_374_v2_)), (self).f33)).set(default__.safeIndex(len(self.f35), len((d_374_v2_).set(default__.safeIndex(396, len(d_374_v2_)), (self).f33))), (self).f33))})
        d_379_v7_: _dafny.Map
        d_379_v7_ = _dafny.Map({(self).f26: (self).f26})
        d_380_v8_: _dafny.Map
        d_380_v8_ = _dafny.Map({len(d_379_v7_): p1})
        index63_ = default__.safeIndex(429, (d_373_v1_).length(0))
        index64_ = default__.safeIndex(289, (d_375_v3_).length(0))
        rhs73_ = ((d_378_v6_)[((d_380_v8_)[(self).f26] if ((self).f26) in (d_380_v8_) else self.f34)] if (((d_380_v8_)[(self).f26] if ((self).f26) in (d_380_v8_) else self.f34)) in (d_378_v6_) else (d_374_v2_) + (_dafny.SeqWithoutIsStrInference([(self).f33, (self).f33])))
        rhs74_ = (self).f25
        rhs75_ = self.f35
        rhs76_ = d_377_v5_
        lhs62_ = d_373_v1_
        lhs63_ = default__.safeIndex(429, (d_373_v1_).length(0))
        lhs64_ = globalState
        lhs65_ = d_375_v3_
        lhs66_ = default__.safeIndex(289, (d_375_v3_).length(0))
        lhs62_[lhs63_] = rhs73_
        lhs64_.f1 = rhs74_
        lhs65_[lhs66_] = rhs75_
        d_377_v5_ = rhs76_
        index65_ = default__.safeIndex(197, (d_376_v4_).length(0))
        (d_376_v4_)[index65_] = default__.safeModuloInt((self).f26, len((d_375_v3_)[default__.safeIndex(289, (d_375_v3_).length(0))]))
        index66_ = default__.safeIndex(197, (d_376_v4_).length(0))
        rhs77_ = p1
        rhs78_ = (0) - ((self).f26)
        rhs79_ = p2
        lhs67_ = d_376_v4_
        lhs68_ = default__.safeIndex(197, (d_376_v4_).length(0))
        r0 = rhs77_
        lhs67_[lhs68_] = rhs78_
        r0 = rhs79_
        (globalState).f2 = default__.safeDivisionInt((self).f26, ((d_376_v4_)[default__.safeIndex(197, (d_376_v4_).length(0))]) * ((d_376_v4_)[default__.safeIndex(197, (d_376_v4_).length(0))]))
        d_381_v9_: _dafny.Map
        d_381_v9_ = _dafny.Map({self.f34: (self).f26})
        (globalState).f2 = ((self).f26 if (d_381_v9_) != (_dafny.Map({self.f34: len((d_375_v3_)[default__.safeIndex(289, (d_375_v3_).length(0))])})) else (len(_dafny.SeqWithoutIsStrInference([not(self.f34)]))) - ((d_376_v4_)[default__.safeIndex(197, (d_376_v4_).length(0))]))
        r0 = p2
        return r0

    def m9(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_382_i0_: int
        d_382_i0_ = 0
        with _dafny.label("2"):
            while (True) and ((self).f25):
                with _dafny.c_label("2"):
                    if (d_382_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_382_i0_ = (d_382_i0_) + (1)
                    (globalState).f1 = ((self).f26) == ((self).f26)
                    index67_ = default__.safeIndex(435, ((self).f33).length(0))
                    ((self).f33)[index67_] = self.f34
                    index68_ = default__.safeIndex(435, ((self).f33).length(0))
                    rhs80_ = default__.safeModuloInt(-188, (0) - ((self).f26))
                    rhs81_ = self.f34
                    rhs82_ = self.f34
                    rhs83_ = ((self).f26 if ((self).f26) <= ((self).f26) else (self).f26)
                    lhs69_ = globalState
                    lhs70_ = (self).f33
                    lhs71_ = default__.safeIndex(435, ((self).f33).length(0))
                    lhs72_ = globalState
                    lhs73_ = globalState
                    lhs69_.f0 = rhs80_
                    lhs70_[lhs71_] = rhs81_
                    lhs72_.f1 = rhs82_
                    lhs73_.f2 = rhs83_
                    d_383_v0_: _dafny.Seq
                    d_383_v0_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                    (globalState).f2 = default__.fm2(((d_383_v0_)[default__.safeIndex((self).fm16(globalState), len(d_383_v0_))] if self.f34 else -33), (0) - (len(d_383_v0_)), globalState)
                    d_384_v1_: _dafny.Set
                    d_384_v1_ = _dafny.Set({((self).f33)[default__.safeIndex(435, ((self).f33).length(0))], ((self).f33)[default__.safeIndex(435, ((self).f33).length(0))]})
                    d_385_v2_: _dafny.Map
                    d_385_v2_ = _dafny.Map({(self).f26: len(d_384_v1_)})
                    d_386_v3_: D2
                    d_386_v3_ = D2_DC7((self).f25, self.f34, ((d_385_v2_)[(self).f26] if ((self).f26) in (d_385_v2_) else (d_383_v0_)[default__.safeIndex((self).f26, len(d_383_v0_))]))
                    d_387_v4_: _dafny.Array
                    nw77_ = _dafny.Array(None, 21)
                    nw77_[int(0)] = d_386_v3_
                    nw77_[int(1)] = d_386_v3_
                    nw77_[int(2)] = d_386_v3_
                    nw77_[int(3)] = d_386_v3_
                    nw77_[int(4)] = d_386_v3_
                    nw77_[int(5)] = d_386_v3_
                    nw77_[int(6)] = D2_DC7(self.f34, ((self).f33)[default__.safeIndex(435, ((self).f33).length(0))], (self).f26)
                    nw77_[int(7)] = d_386_v3_
                    nw77_[int(8)] = D2_DC7((self).f25, ((self).f33)[default__.safeIndex(435, ((self).f33).length(0))], 442)
                    nw77_[int(9)] = d_386_v3_
                    nw77_[int(10)] = d_386_v3_
                    nw77_[int(11)] = d_386_v3_
                    nw77_[int(12)] = d_386_v3_
                    nw77_[int(13)] = d_386_v3_
                    nw77_[int(14)] = (d_386_v3_ if (self).f25 else d_386_v3_)
                    nw77_[int(15)] = d_386_v3_
                    nw77_[int(16)] = d_386_v3_
                    nw77_[int(17)] = d_386_v3_
                    nw77_[int(18)] = d_386_v3_
                    nw77_[int(19)] = d_386_v3_
                    nw77_[int(20)] = d_386_v3_
                    d_387_v4_ = nw77_
                    index69_ = default__.safeIndex(703, (d_387_v4_).length(0))
                    (d_387_v4_)[index69_] = d_386_v3_
                    index70_ = default__.safeIndex(703, (d_387_v4_).length(0))
                    (d_387_v4_)[index70_] = d_386_v3_
                    pass
            pass
        d_388_v5_: _dafny.Map
        d_388_v5_ = _dafny.Map({default__.safeModuloInt((self).f26, (self).f26): p0})
        d_388_v5_ = (d_388_v5_).set(((p0).intersection(p0)).cardinality, (p0).set((self).f25, default__.abs((self).f26)))
        (self).f34 = self.f34
        (self).f34 = self.f34
        d_389_v6_: _dafny.Map
        d_389_v6_ = _dafny.Map({(self).f26: (self).f26})
        d_390_v7_: _dafny.Map
        d_390_v7_ = _dafny.Map({(self).f25: d_389_v6_})
        hi2_ = default__.safeModuloInt(((p0)[(self).f25] if ((self).f25) in (p0) else len(d_390_v7_)), (self).f26)
        for d_391_i1_ in range((0) - (((self).fm6((self).f25, (self).f25, (self).f26, globalState)) + ((self).f26)), hi2_):
            (globalState).f2 = 480
            (globalState).f1 = (d_391_i1_) <= ((self).f26)
            (globalState).f1 = (self).f25
            (globalState).f1 = True
        (globalState).f0 = (self).f26
        d_392_v8_: str
        d_392_v8_ = _dafny.CodePoint('g')
        r0 = d_392_v8_
        r1 = (self).f33
        return r0, r1


class C1:
    def  __init__(self):
        self._f32: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f32):
        (self)._f32 = f32

    def fm14(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehss"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_393_i0_ in range(default__.abs(556))]))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbxgttvgb")))

    def fm15(self, p0, p1, p2, p3, globalState):
        return (len(_dafny.SeqWithoutIsStrInference([644 for d_394_i0_ in range(default__.abs(967))]))) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_395_i1_ in range(default__.abs(727))])), -971, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_396_i2_ in range(default__.abs(-428))])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_397_i3_ in range(default__.abs(907))])), 943])))

    def m7(self, p0, p1, p2, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: _dafny.Map = _dafny.Map({})
        d_398_v0_: _dafny.Array
        def lambda18_(d_399_i0_):
            return (self).f32

        init9_ = lambda18_
        nw78_ = _dafny.Array(None, 7)
        for i0_9_ in range(nw78_.length(0)):
            nw78_[i0_9_] = init9_(i0_9_)
        d_398_v0_ = nw78_
        d_400_v1_: str
        d_400_v1_ = _dafny.CodePoint('o')
        d_401_v2_: _dafny.Map
        d_401_v2_ = _dafny.Map({d_400_v1_: (self).f32})
        index71_ = default__.safeIndex(20, (d_398_v0_).length(0))
        (d_398_v0_)[index71_] = ((d_401_v2_)[d_400_v1_] if (d_400_v1_) in (d_401_v2_) else (self).f32)
        index72_ = default__.safeIndex(20, (d_398_v0_).length(0))
        (d_398_v0_)[index72_] = (self).f32
        d_402_v3_: _dafny.Seq
        d_402_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qby"))
        d_402_v3_ = p2
        d_403_v4_: int
        d_403_v4_ = 238
        d_404_v5_: _dafny.Map
        d_404_v5_ = _dafny.Map({p1: d_403_v4_})
        if default__.fm5(not((self).f32), d_404_v5_, d_403_v4_, globalState):
            d_405_v6_: _dafny.Map
            d_405_v6_ = _dafny.Map({d_403_v4_: p0})
            index73_ = default__.safeIndex(20, (d_398_v0_).length(0))
            (d_398_v0_)[index73_] = ((d_405_v6_).set(d_403_v4_, (self).f32)) == ((d_405_v6_) | (d_405_v6_))
            index74_ = default__.safeIndex(20, (d_398_v0_).length(0))
            rhs84_ = d_403_v4_
            rhs85_ = not((D0_DC1(d_403_v4_, p2, (self).f32, (0) - (d_403_v4_), p0)).cf5)
            lhs74_ = globalState
            lhs75_ = d_398_v0_
            lhs76_ = default__.safeIndex(20, (d_398_v0_).length(0))
            lhs74_.f2 = rhs84_
            lhs75_[lhs76_] = rhs85_
            (globalState).f1 = (self).f32
            d_406_v7_: _dafny.Set
            d_406_v7_ = _dafny.Set({False})
            d_407_v8_: D4
            d_407_v8_ = D4_DC9(d_406_v7_)
            rhs86_ = (d_402_v3_) + (d_402_v3_)
            rhs87_ = default__.fm2(d_403_v4_, len(((d_407_v8_).cf15) - (d_406_v7_)), globalState)
            rhs88_ = (0) - (((0) - (d_403_v4_) if (self).f32 else len((d_406_v7_) - (_dafny.Set({(self).f32})))))
            lhs77_ = globalState
            lhs78_ = globalState
            d_402_v3_ = rhs86_
            lhs77_.f2 = rhs87_
            lhs78_.f0 = rhs88_
            (globalState).f0 = d_403_v4_
        elif True:
            d_408_v9_: _dafny.Map
            d_408_v9_ = _dafny.Map({(d_398_v0_)[default__.safeIndex(20, (d_398_v0_).length(0))]: 347})
            d_409_v10_: _dafny.Seq
            d_409_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f32: d_403_v4_}), d_408_v9_])
            d_410_v11_: _dafny.Map
            d_410_v11_ = _dafny.Map({d_403_v4_: _dafny.MultiSet(d_409_v10_)})
            d_411_v12_: _dafny.Set
            d_411_v12_ = _dafny.Set({(self).f32, True, p0})
            d_412_v13_: _dafny.MultiSet
            d_412_v13_ = _dafny.MultiSet([(d_408_v9_).set((d_398_v0_)[default__.safeIndex(20, (d_398_v0_).length(0))], d_403_v4_)])
            d_413_v14_: _dafny.MultiSet
            d_413_v14_ = ((d_410_v11_)[len(d_411_v12_)] if (len(d_411_v12_)) in (d_410_v11_) else d_412_v13_)
            r0 = d_413_v14_
            d_414_v15_: D0
            d_414_v15_ = D0_DC2()
            d_415_v16_: _dafny.Map
            d_415_v16_ = _dafny.Map({d_414_v15_: (self).f32})
            if (d_414_v15_) in (d_415_v16_):
                d_416_v17_: _dafny.Array
                def lambda19_(d_417_v4_):
                    def lambda20_(d_418_i1_):
                        return (d_418_i1_) * (d_417_v4_)

                    return lambda20_

                init10_ = lambda19_(d_403_v4_)
                nw79_ = _dafny.Array(None, 29)
                for i0_10_ in range(nw79_.length(0)):
                    nw79_[i0_10_] = init10_(i0_10_)
                d_416_v17_ = nw79_
                rhs89_ = (d_403_v4_) < ((d_403_v4_) + (d_403_v4_))
                rhs90_ = d_416_v17_
                lhs79_ = globalState
                lhs79_.f1 = rhs89_
                d_416_v17_ = rhs90_
                d_402_v3_ = (d_402_v3_) + (d_402_v3_)
                (globalState).f2 = d_403_v4_
                d_419_v18_: D0
                d_419_v18_ = D0_DC1(d_403_v4_, d_402_v3_, p0, d_403_v4_, (self).f32)
                d_420_v19_: C0
                nw80_ = C0()
                def iife31_(_pat_let10_0):
                    def iife32_(d_421_dt__update__tmp_h0_):
                        def iife33_(_pat_let11_0):
                            def iife34_(d_422_dt__update_hcf5_h0_):
                                return D0_DC1((d_421_dt__update__tmp_h0_).cf1, (d_421_dt__update__tmp_h0_).cf2, (d_421_dt__update__tmp_h0_).cf3, (d_421_dt__update__tmp_h0_).cf4, d_422_dt__update_hcf5_h0_)
                            return iife34_(_pat_let11_0)
                        return iife33_((self).f32)
                    return iife32_(_pat_let10_0)
                nw80_.ctor__((iife31_(d_419_v18_)).cf2, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yggum"))) != (d_402_v3_), d_403_v4_, d_398_v0_, (d_398_v0_)[default__.safeIndex(20, (d_398_v0_).length(0))])
                d_420_v19_ = nw80_
                (globalState).f1 = p0
            elif True:
                (globalState).f0 = d_403_v4_
                d_400_v1_ = d_400_v1_
                d_423_v20_: _dafny.Seq
                d_423_v20_ = _dafny.SeqWithoutIsStrInference([(self).f32])
                d_424_v21_: _dafny.Array
                nw81_ = _dafny.Array(None, 10)
                nw81_[int(0)] = (224) - (d_403_v4_)
                nw81_[int(1)] = d_403_v4_
                nw81_[int(2)] = 931
                nw81_[int(3)] = d_403_v4_
                nw81_[int(4)] = d_403_v4_
                nw81_[int(5)] = (d_403_v4_) - (-173)
                nw81_[int(6)] = len(p2)
                nw81_[int(7)] = d_403_v4_
                nw81_[int(8)] = default__.safeDivisionInt(len(d_423_v20_), d_403_v4_)
                nw81_[int(9)] = (0) - (d_403_v4_)
                d_424_v21_ = nw81_
                d_425_v22_: D0
                d_425_v22_ = D0_DC1(791, d_402_v3_, (self).f32, d_403_v4_, True)
                index75_ = default__.safeIndex(661, (d_424_v21_).length(0))
                (d_424_v21_)[index75_] = default__.fm2(d_403_v4_, (d_425_v22_).cf1, globalState)
                d_426_v23_: _dafny.Map
                d_426_v23_ = _dafny.Map({d_403_v4_: d_403_v4_})
                d_427_v24_: _dafny.Map
                d_427_v24_ = _dafny.Map({d_424_v21_: d_426_v23_})
                d_428_v25_: _dafny.Map
                d_428_v25_ = _dafny.Map({p0: d_427_v24_})
                d_429_v27_: _dafny.Seq
                d_429_v27_ = _dafny.SeqWithoutIsStrInference([d_403_v4_])
                d_430_v28_: _dafny.MultiSet
                d_430_v28_ = _dafny.MultiSet([len(d_429_v27_), d_403_v4_])
                d_431_v29_: _dafny.Seq
                d_431_v29_ = _dafny.SeqWithoutIsStrInference([p1])
                index76_ = default__.safeIndex(661, (d_424_v21_).length(0))
                index77_ = default__.safeIndex(20, (d_398_v0_).length(0))
                def iife35_():
                    coll11_ = _dafny.Map()
                    compr_11_: int
                    for compr_11_ in (d_430_v28_).Elements:
                        d_432_v26_: int = compr_11_
                        if (d_432_v26_) in (d_430_v28_):
                            coll11_[(d_432_v26_) - (d_403_v4_)] = d_403_v4_
                    return _dafny.Map(coll11_)
                rhs91_ = d_403_v4_
                rhs92_ = (((d_428_v25_)[(d_398_v0_)[default__.safeIndex(20, (d_398_v0_).length(0))]] if ((d_398_v0_)[default__.safeIndex(20, (d_398_v0_).length(0))]) in (d_428_v25_) else d_427_v24_)) | (_dafny.Map({d_424_v21_: (iife35_()
                ).set(617, d_403_v4_)}))
                rhs93_ = (len(d_431_v29_)) > (d_403_v4_)
                rhs94_ = (d_403_v4_) + (d_403_v4_)
                lhs80_ = d_424_v21_
                lhs81_ = default__.safeIndex(661, (d_424_v21_).length(0))
                lhs82_ = d_398_v0_
                lhs83_ = default__.safeIndex(20, (d_398_v0_).length(0))
                lhs84_ = globalState
                lhs80_[lhs81_] = rhs91_
                d_427_v24_ = rhs92_
                lhs82_[lhs83_] = rhs93_
                lhs84_.f2 = rhs94_
                d_408_v9_ = (d_408_v9_) | (d_408_v9_)
                d_402_v3_ = _dafny.SeqWithoutIsStrInference([d_400_v1_ for d_433_i2_ in range(default__.abs(-795))])
            d_434_v30_: _dafny.Map
            d_434_v30_ = _dafny.Map({default__.fm21(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adicjv")), d_403_v4_, d_403_v4_, globalState): (self).f32})
            d_434_v30_ = d_434_v30_
            d_400_v1_ = d_400_v1_
            d_435_v31_: _dafny.Map
            d_435_v31_ = _dafny.Map({(d_403_v4_) + (d_403_v4_): p0})
            d_436_v32_: _dafny.MultiSet
            d_436_v32_ = _dafny.MultiSet([_dafny.CodePoint('v'), d_400_v1_])
            d_435_v31_ = (d_435_v31_).set(d_403_v4_, (d_436_v32_).issubset(d_436_v32_))
        d_437_v33_: _dafny.Set
        d_437_v33_ = _dafny.Set({(self).f32})
        d_438_v34_: D1
        d_438_v34_ = D1_DC4((self).fm14(d_437_v33_, 614, globalState), (d_398_v0_)[default__.safeIndex(20, (d_398_v0_).length(0))])
        d_439_v35_: _dafny.MultiSet
        d_439_v35_ = _dafny.MultiSet([d_438_v34_])
        index78_ = default__.safeIndex(20, (d_398_v0_).length(0))
        (d_398_v0_)[index78_] = (d_439_v35_).isdisjoint(d_439_v35_)
        pat_let_tv9_ = d_403_v4_
        pat_let_tv10_ = d_403_v4_
        pat_let_tv11_ = d_403_v4_
        def lambda21_(source5_):
            if source5_.is_DC4:
                d_440___mcc_h0_ = source5_.cf7
                d_441___mcc_h1_ = source5_.cf8
                d_442_cf8_ = d_441___mcc_h1_
                d_443_cf7_ = d_440___mcc_h0_
                return -400
            elif source5_.is_DC3:
                d_444___mcc_h2_ = source5_.cf6
                d_445_cf6_ = d_444___mcc_h2_
                return default__.safeDivisionInt(pat_let_tv9_, pat_let_tv10_)
            elif True:
                d_446___mcc_h3_ = source5_.cf9
                d_447_cf9_ = d_446___mcc_h3_
                return pat_let_tv11_

        rhs95_ = default__.fm22((len(d_437_v33_)) + (d_403_v4_), d_403_v4_, d_400_v1_, globalState)
        rhs96_ = True
        rhs97_ = ((d_402_v3_) + (d_402_v3_)) + (p2)
        rhs98_ = lambda21_(d_438_v34_)
        rhs99_ = 230
        lhs85_ = globalState
        lhs86_ = globalState
        lhs87_ = globalState
        d_438_v34_ = rhs95_
        lhs85_.f1 = rhs96_
        d_402_v3_ = rhs97_
        lhs86_.f0 = rhs98_
        lhs87_.f0 = rhs99_
        d_448_v36_: _dafny.Seq
        d_448_v36_ = _dafny.SeqWithoutIsStrInference([p0, (d_400_v1_) in (d_402_v3_), (self).f32, (d_398_v0_)[default__.safeIndex(20, (d_398_v0_).length(0))]])
        d_448_v36_ = default__.fm23(globalState)
        d_449_v37_: _dafny.Map
        d_449_v37_ = _dafny.Map({(d_398_v0_)[default__.safeIndex(20, (d_398_v0_).length(0))]: d_403_v4_})
        d_450_v38_: _dafny.MultiSet
        d_450_v38_ = _dafny.MultiSet([d_449_v37_, d_449_v37_])
        d_451_v39_: _dafny.MultiSet
        d_451_v39_ = d_450_v38_
        r0 = d_451_v39_
        d_452_v40_: _dafny.Map
        d_452_v40_ = _dafny.Map({(d_403_v4_) + (d_403_v4_): (self).fm14(d_437_v33_, len(d_448_v36_), globalState)})
        r1 = d_452_v40_
        return r0, r1

    def m8(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: bool = False
        d_453_v0_: str
        d_453_v0_ = _dafny.CodePoint('o')
        d_454_v1_: _dafny.Map
        d_454_v1_ = _dafny.Map({d_453_v0_: p1})
        (globalState).f1 = ((self).f32) == (not(((d_454_v1_)[d_453_v0_] if (d_453_v0_) in (d_454_v1_) else (self).f32)))
        hi3_ = p2
        for d_455_i0_ in range(default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfgmwfj"))), p2), hi3_):
            d_456_v2_: _dafny.Seq
            d_456_v2_ = _dafny.SeqWithoutIsStrInference([(self).f32, (self).f32, p1, p1, p1])
            d_457_v3_: _dafny.Set
            d_457_v3_ = _dafny.Set({p2})
            (globalState).f1 = (d_456_v2_)[default__.safeIndex(len((d_457_v3_) - (d_457_v3_)), len(d_456_v2_))]
            (globalState).f1 = (p2) != (d_455_i0_)
            r1 = (self).f32
            (globalState).f1 = ((default__.fm24((self).f32, globalState)).issubset(d_457_v3_)) or (p1)
        if p1:
            if p1:
                d_458_v4_: _dafny.Seq
                d_458_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qh"))
                d_459_v5_: _dafny.Map
                d_459_v5_ = _dafny.Map({(p2) - (p2): (d_458_v4_) + (d_458_v4_)})
                d_460_v6_: _dafny.Set
                d_460_v6_ = _dafny.Set({(self).f32, p1})
                rhs100_ = len(default__.fm1(globalState))
                rhs101_ = (p1) == (p1)
                rhs102_ = (default__.safeModuloInt(p2, (0) - (p2)) if p1 else p2)
                rhs103_ = ((d_459_v5_)[p2] if (p2) in (d_459_v5_) else d_458_v4_)
                rhs104_ = (len(d_460_v6_)) == (113)
                lhs88_ = globalState
                lhs89_ = globalState
                lhs90_ = globalState
                lhs91_ = globalState
                lhs88_.f0 = rhs100_
                lhs89_.f1 = rhs101_
                lhs90_.f2 = rhs102_
                r0 = rhs103_
                lhs91_.f1 = rhs104_
                (globalState).f2 = (p2) - (p2)
                d_461_v7_: C0
                nw82_ = C0()
                nw82_.ctor__((d_458_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgruw"))), (self).f32, p2, p0, (d_458_v4_) != (d_458_v4_))
                d_461_v7_ = nw82_
                d_462_v8_: _dafny.Map
                d_462_v8_ = _dafny.Map({p1: p1})
                d_463_v9_: _dafny.MultiSet
                d_463_v9_ = _dafny.MultiSet([(self).f32, ((d_462_v8_)[True] if (True) in (d_462_v8_) else (self).f32), (self).f32, (self).f32, not(not(not(p1)))])
                rhs105_ = (((self).f32) or (p1) if (d_463_v9_).isdisjoint(d_463_v9_) else True)
                rhs106_ = d_461_v7_
                rhs107_ = d_453_v0_
                rhs108_ = (_dafny.SeqWithoutIsStrInference([d_453_v0_ for d_464_i1_ in range(default__.abs(19))])) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))
                lhs92_ = globalState
                lhs93_ = globalState
                lhs92_.f1 = rhs105_
                d_461_v7_ = rhs106_
                d_453_v0_ = rhs107_
                lhs93_.f1 = rhs108_
                (globalState).f0 = len(((_dafny.SeqWithoutIsStrInference([d_453_v0_ for d_465_i2_ in range(default__.abs(-248))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_453_v0_ for d_466_i2_ in range(default__.abs(-248))]))), (d_461_v7_.f35)[default__.safeIndex(p2, len(d_461_v7_.f35))])).set(default__.safeIndex((p2) * (p2), len((_dafny.SeqWithoutIsStrInference([d_453_v0_ for d_467_i2_ in range(default__.abs(-248))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_453_v0_ for d_468_i2_ in range(default__.abs(-248))]))), (d_461_v7_.f35)[default__.safeIndex(p2, len(d_461_v7_.f35))]))), (d_461_v7_.f35)[default__.safeIndex((0) - ((0) - (p2)), len(d_461_v7_.f35))]))
                d_469_v10_: _dafny.Array
                nw83_ = _dafny.Array(int(0), 16)
                d_469_v10_ = nw83_
                index79_ = default__.safeIndex(78, (d_469_v10_).length(0))
                (d_469_v10_)[index79_] = p2
                index80_ = default__.safeIndex(78, (d_469_v10_).length(0))
                (d_469_v10_)[index80_] = p2
            elif True:
                d_470_v11_: _dafny.Array
                nw84_ = _dafny.Array(int(0), 6)
                d_470_v11_ = nw84_
                d_471_v12_: _dafny.Set
                d_471_v12_ = _dafny.Set({d_470_v11_})
                d_472_v13_: _dafny.MultiSet
                d_472_v13_ = _dafny.MultiSet([p2])
                d_473_v14_: _dafny.Seq
                d_473_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpv"))
                d_474_v15_: _dafny.Seq
                d_474_v15_ = _dafny.SeqWithoutIsStrInference([p1])
                d_475_v16_: _dafny.Array
                nw85_ = _dafny.Array(None, 8)
                nw85_[int(0)] = d_472_v13_
                nw85_[int(1)] = _dafny.MultiSet([p2, (0) - (len(d_473_v14_)), p2])
                nw85_[int(2)] = (d_472_v13_) - (d_472_v13_)
                nw85_[int(3)] = (default__.fm21(d_473_v14_, len(d_474_v15_), p2, globalState)) | (d_472_v13_)
                nw85_[int(4)] = d_472_v13_
                nw85_[int(5)] = (d_472_v13_).intersection(d_472_v13_)
                nw85_[int(6)] = d_472_v13_
                nw85_[int(7)] = d_472_v13_
                d_475_v16_ = nw85_
                index81_ = default__.safeIndex(28, (d_475_v16_).length(0))
                (d_475_v16_)[index81_] = d_472_v13_
                d_476_v17_: _dafny.MultiSet
                d_476_v17_ = d_472_v13_
                d_477_v18_: _dafny.Set
                d_477_v18_ = _dafny.Set({d_476_v17_, d_476_v17_, d_476_v17_, d_472_v13_})
                d_478_v19_: D6
                d_478_v19_ = D6_DC12(d_471_v12_)
                index82_ = default__.safeIndex(28, (d_475_v16_).length(0))
                rhs109_ = (0) - (p2)
                rhs110_ = ((d_478_v19_).cf19) - (d_471_v12_)
                rhs111_ = d_472_v13_
                rhs112_ = ((d_477_v18_) | (d_477_v18_)).intersection(d_477_v18_)
                lhs94_ = globalState
                lhs95_ = d_475_v16_
                lhs96_ = default__.safeIndex(28, (d_475_v16_).length(0))
                lhs94_.f2 = rhs109_
                d_471_v12_ = rhs110_
                lhs95_[lhs96_] = rhs111_
                d_477_v18_ = rhs112_
                d_479_v20_: _dafny.Array
                def lambda22_(d_480_v14_):
                    def lambda23_(d_481_i3_):
                        return d_480_v14_

                    return lambda23_

                init11_ = lambda22_(d_473_v14_)
                nw86_ = _dafny.Array(None, 24)
                for i0_11_ in range(nw86_.length(0)):
                    nw86_[i0_11_] = init11_(i0_11_)
                d_479_v20_ = nw86_
                d_479_v20_ = d_479_v20_
                (globalState).f0 = (553) + (740)
                (globalState).f0 = p2
                (globalState).f1 = (self).f32
            d_482_v21_: D4
            d_482_v21_ = D4_DC10(d_453_v0_, _dafny.CodePoint('k'))
            d_483_v22_: D1
            d_483_v22_ = D1_DC3(d_453_v0_)
            pat_let_tv12_ = d_483_v22_
            def iife36_(_pat_let12_0):
                def iife37_(d_484_dt__update__tmp_h1_):
                    def iife38_(_pat_let13_0):
                        def iife39_(d_485_dt__update_hcf16_h0_):
                            return D4_DC10(d_485_dt__update_hcf16_h0_, (d_484_dt__update__tmp_h1_).cf17)
                        return iife39_(_pat_let13_0)
                    return iife38_((pat_let_tv12_).cf6)
                return iife37_(_pat_let12_0)
            d_482_v21_ = iife36_(d_482_v21_)
            d_486_v23_: _dafny.Seq
            d_486_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
            d_487_v24_: C0
            nw87_ = C0()
            nw87_.ctor__(d_486_v23_, ((self).f32) and (p1), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "norpc"))), p0, p1)
            d_487_v24_ = nw87_
            (globalState).f0 = (0) - (p2)
            index83_ = default__.safeIndex(708, (p0).length(0))
            (p0)[index83_] = (p2) <= (p2)
            d_488_v25_: _dafny.Array
            nw88_ = _dafny.Array(int(0), 24)
            d_488_v25_ = nw88_
            index84_ = default__.safeIndex(708, (p0).length(0))
            rhs113_ = (p1) == ((p2) != (len(d_486_v23_)))
            rhs114_ = d_488_v25_
            lhs97_ = p0
            lhs98_ = default__.safeIndex(708, (p0).length(0))
            lhs97_[lhs98_] = rhs113_
            d_488_v25_ = rhs114_
        elif True:
            d_489_v26_: _dafny.Array
            nw89_ = _dafny.Array(int(0), 15)
            d_489_v26_ = nw89_
            index85_ = default__.safeIndex(113, (d_489_v26_).length(0))
            (d_489_v26_)[index85_] = p2
            d_490_v27_: _dafny.Seq
            d_490_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scemdf"))
            index86_ = default__.safeIndex(113, (d_489_v26_).length(0))
            rhs115_ = (0) - (len((d_490_v27_) + ((d_490_v27_) + (_dafny.SeqWithoutIsStrInference([d_453_v0_ for d_491_i4_ in range(default__.abs(195))])))))
            rhs116_ = p2
            lhs99_ = d_489_v26_
            lhs100_ = default__.safeIndex(113, (d_489_v26_).length(0))
            lhs101_ = globalState
            lhs99_[lhs100_] = rhs115_
            lhs101_.f0 = rhs116_
            (globalState).f2 = p2
            if (default__.safeModuloInt(p2, 660)) <= ((d_489_v26_)[default__.safeIndex(113, (d_489_v26_).length(0))]):
                r1 = False
                d_492_v28_: _dafny.Array
                def lambda24_(d_493_v0_):
                    def lambda25_(d_494_i5_):
                        return d_493_v0_

                    return lambda25_

                init12_ = lambda24_(d_453_v0_)
                nw90_ = _dafny.Array(None, 5)
                for i0_12_ in range(nw90_.length(0)):
                    nw90_[i0_12_] = init12_(i0_12_)
                d_492_v28_ = nw90_
                index87_ = default__.safeIndex(493, (d_492_v28_).length(0))
                (d_492_v28_)[index87_] = _dafny.CodePoint('h')
                index88_ = default__.safeIndex(493, (d_492_v28_).length(0))
                (d_492_v28_)[index88_] = d_453_v0_
                d_495_v29_: _dafny.Set
                d_495_v29_ = _dafny.Set({(d_489_v26_)[default__.safeIndex(113, (d_489_v26_).length(0))]})
                (globalState).f1 = (len((d_495_v29_) | (_dafny.Set({(d_489_v26_)[default__.safeIndex(113, (d_489_v26_).length(0))]})))) > ((d_489_v26_)[default__.safeIndex(113, (d_489_v26_).length(0))])
                d_496_v30_: _dafny.Array
                def lambda26_(d_497_p1_, d_498_p2_):
                    def lambda27_(d_499_i6_):
                        return (_dafny.SeqWithoutIsStrInference([d_497_p1_])) + (_dafny.SeqWithoutIsStrInference([(D2_DC7(d_497_p1_, (self).f32, d_498_p2_)).cf11]))

                    return lambda27_

                init13_ = lambda26_(p1, p2)
                nw91_ = _dafny.Array(None, 5)
                for i0_13_ in range(nw91_.length(0)):
                    nw91_[i0_13_] = init13_(i0_13_)
                d_496_v30_ = nw91_
                d_500_v31_: _dafny.Seq
                d_500_v31_ = _dafny.SeqWithoutIsStrInference([p1])
                d_501_v32_: _dafny.Seq
                d_501_v32_ = d_500_v31_
                index89_ = default__.safeIndex(834, (d_496_v30_).length(0))
                (d_496_v30_)[index89_] = (d_501_v32_)
                d_502_v33_: D6
                d_502_v33_ = D6_DC12(_dafny.Set({d_489_v26_}))
                d_503_v34_: _dafny.Set
                d_503_v34_ = _dafny.Set({d_489_v26_})
                index90_ = default__.safeIndex(834, (d_496_v30_).length(0))
                (d_496_v30_)[index90_] = ((d_500_v31_) + (d_500_v31_) if (_dafny.MultiSet([d_502_v33_])) == (_dafny.MultiSet([D6_DC12(d_503_v34_), d_502_v33_])) else (d_500_v31_) + (_dafny.SeqWithoutIsStrInference([(self).f32, (self).f32])))
                (globalState).f1 = (self).f32
            elif True:
                d_504_v35_: _dafny.Set
                d_504_v35_ = _dafny.Set({p1})
                d_505_v36_: C0
                nw92_ = C0()
                nw92_.ctor__(d_490_v27_, (self).fm14(d_504_v35_, (d_489_v26_)[default__.safeIndex(113, (d_489_v26_).length(0))], globalState), (0) - (p2), p0, p1)
                d_505_v36_ = nw92_
                d_506_v37_: _dafny.Set
                d_506_v37_ = _dafny.Set({d_505_v36_})
                d_506_v37_ = d_506_v37_
                (globalState).f2 = (d_489_v26_)[default__.safeIndex(113, (d_489_v26_).length(0))]
                d_504_v35_ = d_504_v35_
                index91_ = default__.safeIndex(113, (d_489_v26_).length(0))
                (d_489_v26_)[index91_] = (p2) * ((0) - ((d_489_v26_)[default__.safeIndex(113, (d_489_v26_).length(0))]))
                d_507_v38_: C0
                nw93_ = C0()
                nw93_.ctor__(d_505_v36_.f35, (self).f32, len(d_490_v27_), p0, (d_505_v36_).fm7(globalState))
                d_507_v38_ = nw93_
            (globalState).f1 = (self).f32
            d_489_v26_ = d_489_v26_
        d_508_v39_: D0
        d_508_v39_ = D0_DC0((self).f32)
        d_508_v39_ = d_508_v39_
        d_509_v40_: _dafny.Array
        def lambda28_(d_510_p1_):
            def lambda29_(d_511_i7_):
                return _dafny.Map({d_510_p1_: (self).f32})

            return lambda29_

        init14_ = lambda28_(p1)
        nw94_ = _dafny.Array(None, 15)
        for i0_14_ in range(nw94_.length(0)):
            nw94_[i0_14_] = init14_(i0_14_)
        d_509_v40_ = nw94_
        d_512_v41_: D8
        d_512_v41_ = D8_DC16(d_509_v40_)
        d_513_v42_: _dafny.Map
        d_513_v42_ = _dafny.Map({(d_512_v41_).cf27: not(p1)})
        d_514_v43_: _dafny.Seq
        d_514_v43_ = _dafny.SeqWithoutIsStrInference([p1, (self).f32, not(False)])
        d_515_v44_: _dafny.Seq
        d_515_v44_ = _dafny.SeqWithoutIsStrInference([p2])
        d_513_v42_ = (d_513_v42_).set(d_509_v40_, (d_514_v43_)[default__.safeIndex(len(d_515_v44_), len(d_514_v43_))])
        d_516_v45_: D1
        d_516_v45_ = D1_DC3(d_453_v0_)
        d_517_v46_: D1
        d_517_v46_ = D1_DC5(d_516_v45_)
        source6_ = d_517_v46_
        if source6_.is_DC4:
            d_518___mcc_h0_ = source6_.cf7
            d_519___mcc_h1_ = source6_.cf8
            d_520_cf8_ = d_519___mcc_h1_
            d_521_cf7_ = d_518___mcc_h0_
            if (d_514_v43_)[default__.safeIndex(p2, len(d_514_v43_))]:
                pat_let_tv13_ = p1
                d_522_v47_: _dafny.Seq
                def iife40_(_pat_let14_0):
                    def iife41_(d_523_dt__update__tmp_h2_):
                        def iife42_(_pat_let15_0):
                            def iife43_(d_524_dt__update_hcf0_h0_):
                                return D0_DC0(d_524_dt__update_hcf0_h0_)
                            return iife43_(_pat_let15_0)
                        return iife42_(pat_let_tv13_)
                    return iife41_(_pat_let14_0)
                d_522_v47_ = _dafny.SeqWithoutIsStrInference([d_508_v39_, iife40_(d_508_v39_), D0_DC0(d_521_cf7_)])
                d_522_v47_ = d_522_v47_
                (globalState).f1 = (self).f32
                index92_ = default__.safeIndex(816, (p0).length(0))
                (p0)[index92_] = (self).f32
                d_525_v48_: _dafny.Set
                d_525_v48_ = _dafny.Set({p2})
                d_526_v49_: _dafny.MultiSet
                d_526_v49_ = _dafny.MultiSet([p1])
                d_527_v50_: D6
                d_527_v50_ = D6_DC14(444, len(d_525_v48_), d_526_v49_)
                d_528_v51_: _dafny.Map
                d_528_v51_ = _dafny.Map({default__.fm2(p2, default__.fm2(p2, len(_dafny.Map({(self).f32: (d_527_v50_).cf24})), globalState), globalState): p2})
                index93_ = default__.safeIndex(816, (p0).length(0))
                (p0)[index93_] = not((d_528_v51_) != ((d_528_v51_) | (d_528_v51_)))
                d_529_v52_: _dafny.Array
                nw95_ = _dafny.Array(int(0), 29)
                d_529_v52_ = nw95_
                index94_ = default__.safeIndex(533, (d_529_v52_).length(0))
                (d_529_v52_)[index94_] = p2
                d_530_v53_: _dafny.Set
                d_530_v53_ = _dafny.Set({(p0)[default__.safeIndex(816, (p0).length(0))], d_520_cf8_, (self).f32})
                d_531_v54_: _dafny.Set
                d_531_v54_ = _dafny.Set({(self).f32, (self).fm14(d_530_v53_, p2, globalState)})
                d_532_v55_: D0
                d_532_v55_ = D0_DC1(p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcg")), not((p0)[default__.safeIndex(816, (p0).length(0))]), p2, (self).fm14(d_530_v53_, p2, globalState))
                index95_ = default__.safeIndex(533, (d_529_v52_).length(0))
                rhs117_ = default__.fm4(d_532_v55_, (p0)[default__.safeIndex(816, (p0).length(0))], p2, globalState)
                rhs118_ = p2
                lhs102_ = d_529_v52_
                lhs103_ = default__.safeIndex(533, (d_529_v52_).length(0))
                d_453_v0_ = rhs117_
                lhs102_[lhs103_] = rhs118_
                d_533_v56_: _dafny.Map
                d_533_v56_ = _dafny.Map({(self).f32: (d_515_v44_) <= (d_515_v44_)})
                d_534_v57_: _dafny.Array
                nw96_ = _dafny.Array(None, 20)
                nw96_[int(0)] = d_514_v43_
                nw96_[int(1)] = (d_514_v43_) + ((d_514_v43_).set(default__.safeIndex((d_529_v52_)[default__.safeIndex(533, (d_529_v52_).length(0))], len(d_514_v43_)), p1))
                nw96_[int(2)] = d_514_v43_
                nw96_[int(3)] = _dafny.SeqWithoutIsStrInference([not((self).f32)])
                nw96_[int(4)] = _dafny.SeqWithoutIsStrInference([d_521_cf7_])
                nw96_[int(5)] = d_514_v43_
                nw96_[int(6)] = d_514_v43_
                nw96_[int(7)] = d_514_v43_
                nw96_[int(8)] = (_dafny.SeqWithoutIsStrInference([(self).f32, (d_514_v43_)[default__.safeIndex((d_529_v52_)[default__.safeIndex(533, (d_529_v52_).length(0))], len(d_514_v43_))]])).set(default__.safeIndex((d_529_v52_)[default__.safeIndex(533, (d_529_v52_).length(0))], len(_dafny.SeqWithoutIsStrInference([(self).f32, (d_514_v43_)[default__.safeIndex((d_529_v52_)[default__.safeIndex(533, (d_529_v52_).length(0))], len(d_514_v43_))]]))), (p0)[default__.safeIndex(816, (p0).length(0))])
                nw96_[int(9)] = (d_514_v43_) + (_dafny.SeqWithoutIsStrInference([d_521_cf7_]))
                nw96_[int(10)] = d_514_v43_
                nw96_[int(11)] = d_514_v43_
                nw96_[int(12)] = d_514_v43_
                nw96_[int(13)] = d_514_v43_
                nw96_[int(14)] = _dafny.SeqWithoutIsStrInference([True])
                nw96_[int(15)] = d_514_v43_
                nw96_[int(16)] = d_514_v43_
                nw96_[int(17)] = d_514_v43_
                nw96_[int(18)] = d_514_v43_
                nw96_[int(19)] = d_514_v43_
                d_534_v57_ = nw96_
                index96_ = default__.safeIndex(586, (d_534_v57_).length(0))
                (d_534_v57_)[index96_] = d_514_v43_
                index97_ = default__.safeIndex(533, (d_529_v52_).length(0))
                index98_ = default__.safeIndex(586, (d_534_v57_).length(0))
                rhs119_ = p2
                rhs120_ = d_521_cf7_
                rhs121_ = d_533_v56_
                rhs122_ = d_514_v43_
                lhs104_ = d_529_v52_
                lhs105_ = default__.safeIndex(533, (d_529_v52_).length(0))
                lhs106_ = d_534_v57_
                lhs107_ = default__.safeIndex(586, (d_534_v57_).length(0))
                lhs104_[lhs105_] = rhs119_
                d_521_cf7_ = rhs120_
                d_533_v56_ = rhs121_
                lhs106_[lhs107_] = rhs122_
            elif True:
                d_535_v58_: _dafny.Map
                d_535_v58_ = _dafny.Map({p1: 609})
                (globalState).f0 = default__.safeDivisionInt(p2, ((d_535_v58_)[not((self).f32)] if (not((self).f32)) in (d_535_v58_) else (0) - (p2)))
                d_536_v59_: _dafny.Array
                nw97_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_536_v59_ = nw97_
                index99_ = default__.safeIndex(148, (d_536_v59_).length(0))
                (d_536_v59_)[index99_] = p0
                index100_ = default__.safeIndex(148, (d_536_v59_).length(0))
                (d_536_v59_)[index100_] = p0
                (globalState).f0 = (p2) + (p2)
                (globalState).f0 = default__.safeModuloInt(-529, (default__.fm2(p2, len(_dafny.Map({625: False})), globalState)) * (p2))
                arr0_ = (d_536_v59_)[default__.safeIndex(148, (d_536_v59_).length(0))]
                index101_ = default__.safeIndex(478, ((d_536_v59_)[default__.safeIndex(148, (d_536_v59_).length(0))]).length(0))
                arr0_[index101_] = p1
                d_537_v60_: _dafny.Seq
                d_537_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                d_538_v61_: T0
                nw98_ = C0()
                nw98_.ctor__((d_537_v60_) + (d_537_v60_), (self).f32, default__.safeModuloInt(p2, p2), (d_536_v59_)[default__.safeIndex(148, (d_536_v59_).length(0))], True)
                d_538_v61_ = nw98_
                d_539_v62_: C0
                nw99_ = C0()
                nw99_.ctor__(d_537_v60_, False, p2, p0, (self).f32)
                d_539_v62_ = nw99_
                d_540_v63_: D6
                d_540_v63_ = D6_DC14((d_538_v61_).f26, len(_dafny.Map({(self).f32: d_514_v43_})), _dafny.MultiSet([not(d_521_cf7_)]))
                arr1_ = (d_536_v59_)[default__.safeIndex(148, (d_536_v59_).length(0))]
                index102_ = default__.safeIndex(478, ((d_536_v59_)[default__.safeIndex(148, (d_536_v59_).length(0))]).length(0))
                rhs123_ = ((d_540_v63_).cf23) != (p2)
                rhs124_ = ((d_538_v61_).f26) - ((d_538_v61_).f26)
                rhs125_ = d_538_v61_
                rhs126_ = d_539_v62_
                lhs108_ = (d_536_v59_)[default__.safeIndex(148, (d_536_v59_).length(0))]
                lhs109_ = default__.safeIndex(478, ((d_536_v59_)[default__.safeIndex(148, (d_536_v59_).length(0))]).length(0))
                lhs110_ = globalState
                lhs108_[lhs109_] = rhs123_
                lhs110_.f0 = rhs124_
                d_538_v61_ = rhs125_
                d_539_v62_ = rhs126_
            index103_ = default__.safeIndex(946, (p0).length(0))
            (p0)[index103_] = d_520_cf8_
            index104_ = default__.safeIndex(946, (p0).length(0))
            (p0)[index104_] = True
            d_541_v64_: _dafny.Array
            def lambda30_(d_542_v46_):
                def lambda31_(d_543_i8_):
                    return d_542_v46_

                return lambda31_

            init15_ = lambda30_(d_517_v46_)
            nw100_ = _dafny.Array(None, 15)
            for i0_15_ in range(nw100_.length(0)):
                nw100_[i0_15_] = init15_(i0_15_)
            d_541_v64_ = nw100_
            d_541_v64_ = d_541_v64_
            (globalState).f1 = (self).f32
        elif source6_.is_DC3:
            d_544___mcc_h2_ = source6_.cf6
            d_545_cf6_ = d_544___mcc_h2_
            d_546_v65_: _dafny.MultiSet
            d_547_v66_: _dafny.Map
            out14_: _dafny.MultiSet
            out15_: _dafny.Map
            out14_, out15_ = (self).m7(p1, d_508_v39_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_548_i9_ in range(default__.abs(666))]), globalState)
            d_546_v65_ = out14_
            d_547_v66_ = out15_
            d_549_v67_: _dafny.Set
            d_549_v67_ = _dafny.Set({p1})
            d_550_v68_: D4
            d_550_v68_ = D4_DC9((d_549_v67_) | (d_549_v67_))
            source7_ = d_550_v68_
            if source7_.is_DC10:
                d_551___mcc_h4_ = source7_.cf16
                d_552___mcc_h5_ = source7_.cf17
                d_553_cf17_ = d_552___mcc_h5_
                d_554_cf16_ = d_551___mcc_h4_
                d_555_v69_: _dafny.Array
                nw101_ = _dafny.Array(_dafny.MultiSet({}), 15)
                d_555_v69_ = nw101_
                d_556_v70_: _dafny.Seq
                d_556_v70_ = _dafny.SeqWithoutIsStrInference([d_555_v69_, d_555_v69_, d_555_v69_])
                d_557_v71_: _dafny.Seq
                d_557_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cc"))
                d_555_v69_ = (d_556_v70_)[default__.safeIndex(len((d_557_v71_) + (d_557_v71_)), len(d_556_v70_))]
                index105_ = default__.safeIndex(269, (p0).length(0))
                (p0)[index105_] = p1
                index106_ = default__.safeIndex(269, (p0).length(0))
                (p0)[index106_] = True
                index107_ = default__.safeIndex(269, (p0).length(0))
                (p0)[index107_] = not((d_554_cf16_) not in (d_557_v71_))
                d_558_v72_: _dafny.Map
                d_558_v72_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukvt")): p1})
                d_558_v72_ = d_558_v72_
            elif True:
                d_559___mcc_h6_ = source7_.cf15
                d_560_cf15_ = d_559___mcc_h6_
                d_561_v73_: _dafny.Seq
                d_561_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qten"))
                d_562_v74_: C0
                nw102_ = C0()
                nw102_.ctor__(d_561_v73_, not((self).f32), p2, p0, True)
                d_562_v74_ = nw102_
                d_563_v75_: _dafny.Seq
                d_563_v75_ = _dafny.SeqWithoutIsStrInference([d_562_v74_, d_562_v74_])
                d_564_v76_: _dafny.Map
                d_564_v76_ = _dafny.Map({(d_514_v43_)[default__.safeIndex(p2, len(d_514_v43_))]: ((d_563_v75_).set(default__.safeIndex(p2, len(d_563_v75_)), d_562_v74_)) + (d_563_v75_)})
                d_564_v76_ = (d_564_v76_).set(p1, d_563_v75_)
                d_565_v77_: _dafny.Array
                def lambda32_(d_566_p2_):
                    def lambda33_(d_567_i10_):
                        return default__.safeDivisionInt(d_567_i10_, d_566_p2_)

                    return lambda33_

                init16_ = lambda32_(p2)
                nw103_ = _dafny.Array(None, 29)
                for i0_16_ in range(nw103_.length(0)):
                    nw103_[i0_16_] = init16_(i0_16_)
                d_565_v77_ = nw103_
                d_568_v78_: _dafny.Array
                nw104_ = _dafny.Array(None, 4)
                nw104_[int(0)] = d_565_v77_
                nw104_[int(1)] = d_565_v77_
                nw104_[int(2)] = d_565_v77_
                nw104_[int(3)] = d_565_v77_
                d_568_v78_ = nw104_
                d_568_v78_ = d_568_v78_
                d_569_v79_: _dafny.MultiSet
                d_569_v79_ = _dafny.MultiSet([(self).f32])
                d_570_v80_: D6
                d_570_v80_ = D6_DC14((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_545_cf6_ for d_571_i11_ in range(default__.abs(63))]))).cardinality, p2, d_569_v79_)
                d_572_v81_: _dafny.Set
                d_572_v81_ = _dafny.Set({p2})
                d_573_v82_: D6
                d_573_v82_ = D6_DC13((d_570_v80_).cf23, -119, d_572_v81_)
                d_574_v83_: _dafny.Map
                d_574_v83_ = _dafny.Map({(self).f32: p2})
                d_573_v82_ = D6_DC13(((d_574_v83_)[not(not(True))] if (not(not(True))) in (d_574_v83_) else 682), default__.safeModuloInt(len(d_514_v43_), p2), d_572_v81_)
                (globalState).f2 = p2
            d_575_v84_: _dafny.MultiSet
            d_575_v84_ = _dafny.MultiSet([(self).f32, (self).f32, not(p1)])
            d_575_v84_ = (_dafny.MultiSet([p1])) - (d_575_v84_)
            if (self).f32:
                d_515_v44_ = (d_515_v44_) + (d_515_v44_)
                r0 = default__.fm1(globalState)
                d_576_v85_: _dafny.Array
                nw105_ = _dafny.Array(D0.default()(), 23)
                d_576_v85_ = nw105_
                d_577_v86_: D0
                d_577_v86_ = D0_DC2()
                index108_ = default__.safeIndex(68, (d_576_v85_).length(0))
                (d_576_v85_)[index108_] = d_577_v86_
                index109_ = default__.safeIndex(68, (d_576_v85_).length(0))
                (d_576_v85_)[index109_] = D0_DC2()
                d_578_v87_: _dafny.Map
                d_578_v87_ = _dafny.Map({p1: (0) - (p2)})
                d_579_v88_: _dafny.MultiSet
                d_579_v88_ = _dafny.MultiSet([d_578_v87_])
                rhs127_ = p1
                rhs128_ = (d_579_v88_) - (d_579_v88_)
                lhs111_ = globalState
                lhs111_.f1 = rhs127_
                d_546_v65_ = rhs128_
                d_580_v89_: _dafny.Seq
                d_580_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbsu"))
                d_581_v90_: C0
                nw106_ = C0()
                nw106_.ctor__(d_580_v89_, p1, p2, p0, not(p1))
                d_581_v90_ = nw106_
            elif True:
                d_582_v91_: _dafny.Seq
                d_582_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doq"))
                d_583_v92_: _dafny.Array
                nw107_ = _dafny.Array(None, 9)
                nw107_[int(0)] = d_582_v91_
                nw107_[int(1)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqocpbbe"))) + (d_582_v91_)
                nw107_[int(2)] = d_582_v91_
                nw107_[int(3)] = d_582_v91_
                nw107_[int(4)] = (d_582_v91_) + (d_582_v91_)
                nw107_[int(5)] = d_582_v91_
                nw107_[int(6)] = d_582_v91_
                nw107_[int(7)] = d_582_v91_
                nw107_[int(8)] = d_582_v91_
                d_583_v92_ = nw107_
                d_583_v92_ = d_583_v92_
                index110_ = default__.safeIndex(509, (d_583_v92_).length(0))
                (d_583_v92_)[index110_] = d_582_v91_
                index111_ = default__.safeIndex(509, (d_583_v92_).length(0))
                rhs129_ = ((p2) + (724)) != (p2)
                rhs130_ = d_582_v91_
                lhs112_ = globalState
                lhs113_ = d_583_v92_
                lhs114_ = default__.safeIndex(509, (d_583_v92_).length(0))
                lhs112_.f1 = rhs129_
                lhs113_[lhs114_] = rhs130_
                r1 = ((len(d_549_v67_)) > (p2)) == (p1)
                d_584_v93_: _dafny.Array
                def lambda34_(d_585_i12_):
                    return default__.safeModuloInt(d_585_i12_, -966)

                init17_ = lambda34_
                nw108_ = _dafny.Array(None, 17)
                for i0_17_ in range(nw108_.length(0)):
                    nw108_[i0_17_] = init17_(i0_17_)
                d_584_v93_ = nw108_
                index112_ = default__.safeIndex(307, (d_584_v93_).length(0))
                (d_584_v93_)[index112_] = (0) - (p2)
                index113_ = default__.safeIndex(307, (d_584_v93_).length(0))
                (d_584_v93_)[index113_] = (p2) + (len(_dafny.SeqWithoutIsStrInference([d_545_cf6_ for d_586_i13_ in range(default__.abs(478))])))
                index114_ = default__.safeIndex(307, (d_584_v93_).length(0))
                index115_ = default__.safeIndex(307, (d_584_v93_).length(0))
                rhs131_ = -988
                rhs132_ = p2
                lhs115_ = d_584_v93_
                lhs116_ = default__.safeIndex(307, (d_584_v93_).length(0))
                lhs117_ = d_584_v93_
                lhs118_ = default__.safeIndex(307, (d_584_v93_).length(0))
                lhs115_[lhs116_] = rhs131_
                lhs117_[lhs118_] = rhs132_
        elif True:
            d_587___mcc_h3_ = source6_.cf9
            d_588_cf9_ = d_587___mcc_h3_
            d_589_v94_: _dafny.Seq
            d_589_v94_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmndt"))
            d_590_v95_: C0
            nw109_ = C0()
            nw109_.ctor__(d_589_v94_, (self).f32, p2, p0, p1)
            d_590_v95_ = nw109_
            d_591_v96_: _dafny.Array
            nw110_ = _dafny.Array(_dafny.Map({}), 25)
            d_591_v96_ = nw110_
            d_592_v97_: D4
            d_592_v97_ = D4_DC10(_dafny.CodePoint('o'), _dafny.CodePoint('p'))
            d_593_v98_: _dafny.Array
            nw111_ = _dafny.Array(None, 26)
            nw111_[int(0)] = d_453_v0_
            nw111_[int(1)] = d_453_v0_
            nw111_[int(2)] = _dafny.CodePoint('i')
            nw111_[int(3)] = d_453_v0_
            nw111_[int(4)] = d_453_v0_
            nw111_[int(5)] = d_453_v0_
            nw111_[int(6)] = (d_592_v97_).cf16
            nw111_[int(7)] = d_453_v0_
            nw111_[int(8)] = d_453_v0_
            nw111_[int(9)] = d_453_v0_
            nw111_[int(10)] = _dafny.CodePoint('b')
            nw111_[int(11)] = d_453_v0_
            nw111_[int(12)] = _dafny.CodePoint('s')
            nw111_[int(13)] = d_453_v0_
            nw111_[int(14)] = _dafny.CodePoint('o')
            nw111_[int(15)] = _dafny.CodePoint('x')
            nw111_[int(16)] = d_453_v0_
            nw111_[int(17)] = (d_590_v95_.f35)[default__.safeIndex(p2, len(d_590_v95_.f35))]
            nw111_[int(18)] = d_453_v0_
            nw111_[int(19)] = d_453_v0_
            nw111_[int(20)] = d_453_v0_
            nw111_[int(21)] = d_453_v0_
            nw111_[int(22)] = d_453_v0_
            nw111_[int(23)] = d_453_v0_
            nw111_[int(24)] = d_453_v0_
            nw111_[int(25)] = d_453_v0_
            d_593_v98_ = nw111_
            d_594_v99_: _dafny.Map
            d_594_v99_ = _dafny.Map({p2: d_593_v98_})
            d_595_v100_: _dafny.Map
            d_595_v100_ = _dafny.Map({d_594_v99_: d_591_v96_})
            d_591_v96_ = ((d_595_v100_)[(d_594_v99_) | (d_594_v99_)] if ((d_594_v99_) | (d_594_v99_)) in (d_595_v100_) else d_591_v96_)
            d_596_v101_: _dafny.Array
            nw112_ = _dafny.Array(int(0), 16)
            d_596_v101_ = nw112_
            index116_ = default__.safeIndex(456, (d_596_v101_).length(0))
            (d_596_v101_)[index116_] = p2
            index117_ = default__.safeIndex(796, (p0).length(0))
            (p0)[index117_] = (d_514_v43_) < (_dafny.SeqWithoutIsStrInference([p1, p1]))
            index118_ = default__.safeIndex(456, (d_596_v101_).length(0))
            index119_ = default__.safeIndex(796, (p0).length(0))
            rhs133_ = (p2) * ((len((d_590_v95_.f35).set(default__.safeIndex(439, len(d_590_v95_.f35)), d_453_v0_))) * (p2))
            rhs134_ = False
            lhs119_ = d_596_v101_
            lhs120_ = default__.safeIndex(456, (d_596_v101_).length(0))
            lhs121_ = p0
            lhs122_ = default__.safeIndex(796, (p0).length(0))
            lhs119_[lhs120_] = rhs133_
            lhs121_[lhs122_] = rhs134_
            (globalState).f1 = ((p0)[default__.safeIndex(796, (p0).length(0))]) == (p1)
        d_597_v102_: _dafny.Seq
        d_597_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "km"))
        r0 = d_597_v102_
        r1 = p1
        return r0, r1

    @property
    def f32(self):
        return self._f32

class C2(T0):
    def  __init__(self):
        self._f25: bool = False
        self._f26: int = int(0)
        self._f31: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f31, f25, f26):
        (self)._f31 = f31
        (self)._f25 = f25
        (self)._f26 = f26

    def fm6(self, p0, p1, p2, globalState):
        return (self).f31

    def fm7(self, globalState):
        return (D0_DC0((self).f25)).cf0

    def fm11(self, p0, p1, p2, globalState):
        def lambda35_(source8_):
            if source8_.is_DC1:
                d_598___mcc_h0_ = source8_.cf1
                d_599___mcc_h1_ = source8_.cf2
                d_600___mcc_h2_ = source8_.cf3
                d_601___mcc_h3_ = source8_.cf4
                d_602___mcc_h4_ = source8_.cf5
                d_603_cf5_ = d_602___mcc_h4_
                d_604_cf4_ = d_601___mcc_h3_
                d_605_cf3_ = d_600___mcc_h2_
                d_606_cf2_ = d_599___mcc_h1_
                d_607_cf1_ = d_598___mcc_h0_
                return d_604_cf4_
            elif source8_.is_DC2:
                return default__.safeModuloInt((self).f26, (self).f31)
            elif True:
                d_608___mcc_h5_ = source8_.cf0
                d_609_cf0_ = d_608___mcc_h5_
                return default__.safeDivisionInt(523, (self).f26)

        return (0) - (lambda35_(D0_DC1((self).f31, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lyagi")), (self).f25, -991, (self).f25)))

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        hi4_ = (self).f31
        for d_610_i0_ in range((self).f31, hi4_):
            d_611_v0_: _dafny.Map
            d_611_v0_ = _dafny.Map({p1: p2})
            d_611_v0_ = (d_611_v0_).set((self).f25, p2)
            d_612_v1_: D0
            d_612_v1_ = D0_DC2()
            d_613_v2_: _dafny.Map
            d_613_v2_ = _dafny.Map({64: p1})
            d_614_v3_: _dafny.Seq
            d_614_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxrsfndo"))
            d_615_v4_: _dafny.Map
            d_615_v4_ = _dafny.Map({154: len(d_614_v3_)})
            d_616_v5_: _dafny.Set
            d_616_v5_ = _dafny.Set({(p0).cardinality})
            d_617_v6_: D2
            d_617_v6_ = D2_DC6(_dafny.Set({(self).f26}))
            rhs135_ = (self).f25
            rhs136_ = d_612_v1_
            rhs137_ = (self).fm6(p2, ((d_613_v2_)[((d_615_v4_)[(self).f26] if ((self).f26) in (d_615_v4_) else d_610_i0_)] if (((d_615_v4_)[(self).f26] if ((self).f26) in (d_615_v4_) else d_610_i0_)) in (d_613_v2_) else p2), default__.safeDivisionInt((self).f26, d_610_i0_), globalState)
            rhs138_ = (d_616_v5_) == (((d_617_v6_).cf10) - (d_616_v5_))
            lhs123_ = globalState
            lhs124_ = globalState
            lhs125_ = globalState
            lhs123_.f1 = rhs135_
            d_612_v1_ = rhs136_
            lhs124_.f0 = rhs137_
            lhs125_.f1 = rhs138_
            d_618_v7_: D0
            d_618_v7_ = D0_DC0(True)
            d_619_v8_: _dafny.Map
            d_619_v8_ = _dafny.Map({d_618_v7_: (self).f26})
            def iife44_():
                coll12_ = _dafny.Set()
                compr_12_: str
                for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_620_i1_ in range(default__.abs(-118))])).Elements:
                    d_621_v9_: str = compr_12_
                    if (d_621_v9_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_620_i1_ in range(default__.abs(-118))])):
                        coll12_ = coll12_.union(_dafny.Set([d_621_v9_]))
                return _dafny.Set(coll12_)
            (globalState).f2 = (d_610_i0_ if not(not(default__.fm5(p1, d_619_v8_, len(iife44_()
            ), globalState))) else (self).f26)
            d_622_v10_: _dafny.Map
            d_622_v10_ = _dafny.Map({p2: d_614_v3_})
            d_622_v10_ = (d_622_v10_).set(default__.fm5(p1, d_619_v8_, (self).f31, globalState), d_614_v3_)
        (globalState).f0 = (self).f31
        d_623_v11_: _dafny.Seq
        d_623_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mt"))
        d_624_v12_: D1
        d_624_v12_ = D1_DC4(p1, p1)
        pat_let_tv14_ = d_624_v12_
        pat_let_tv15_ = p1
        pat_let_tv16_ = p2
        pat_let_tv17_ = p2
        pat_let_tv18_ = p1
        pat_let_tv19_ = p1
        def lambda36_(source9_):
            if source9_.is_DC4:
                d_625___mcc_h0_ = source9_.cf7
                d_626___mcc_h1_ = source9_.cf8
                d_627_cf8_ = d_626___mcc_h1_
                d_628_cf7_ = d_625___mcc_h0_
                return (pat_let_tv14_).cf8
            elif source9_.is_DC3:
                d_629___mcc_h2_ = source9_.cf6
                d_630_cf6_ = d_629___mcc_h2_
                return (self).f25
            elif True:
                d_631___mcc_h3_ = source9_.cf9
                d_632_cf9_ = d_631___mcc_h3_
                return (_dafny.MultiSet([(self).f25])).ispropersubset(_dafny.MultiSet([pat_let_tv15_, pat_let_tv16_, pat_let_tv17_, pat_let_tv18_, pat_let_tv19_]))

        rhs139_ = p1
        rhs140_ = default__.fm1(globalState)
        rhs141_ = lambda36_(d_624_v12_)
        rhs142_ = not(not(not(p1)))
        rhs143_ = ((self).f31) < ((self).f26)
        lhs126_ = globalState
        lhs127_ = globalState
        r0 = rhs139_
        d_623_v11_ = rhs140_
        lhs126_.f1 = rhs141_
        lhs127_.f1 = rhs142_
        r0 = rhs143_
        hi5_ = ((self).f31) * ((self).f26)
        for d_633_i2_ in range(default__.safeModuloInt((self).f26, (self).f31), hi5_):
            d_634_v13_: _dafny.Map
            d_634_v13_ = _dafny.Map({(self).f31: p2})
            d_635_v14_: _dafny.Map
            d_635_v14_ = _dafny.Map({d_634_v13_: True})
            d_636_v17_: D0
            d_636_v17_ = D0_DC0(p1)
            d_637_v18_: D0
            def iife45_():
                def iife47_():
                    coll15_ = _dafny.Map()
                    compr_15_: D0
                    for compr_15_ in (_dafny.SeqWithoutIsStrInference([d_636_v17_, d_636_v17_])).Elements:
                        d_638_v16_: D0 = compr_15_
                        if (d_638_v16_) in (_dafny.SeqWithoutIsStrInference([d_636_v17_, d_636_v17_])):
                            coll15_[d_638_v16_] = (self).f25
                    return _dafny.Map(coll15_)
                coll13_ = _dafny.Map()
                def iife46_():
                    coll14_ = _dafny.Map()
                    compr_14_: D0
                    for compr_14_ in (_dafny.SeqWithoutIsStrInference([d_636_v17_, d_636_v17_])).Elements:
                        d_638_v16_: D0 = compr_14_
                        if (d_638_v16_) in (_dafny.SeqWithoutIsStrInference([d_636_v17_, d_636_v17_])):
                            coll14_[d_638_v16_] = (self).f25
                    return _dafny.Map(coll14_)
                compr_13_: D0
                for compr_13_ in (iife46_()
                ).keys.Elements:
                    d_639_v15_: D0 = compr_13_
                    if (d_639_v15_) in (iife47_()
                    ):
                        coll13_[d_639_v15_] = len(_dafny.Map({d_633_i2_: (self).f31}))
                return _dafny.Map(coll13_)
            d_637_v18_ = D0_DC1(d_633_i2_, d_623_v11_, p1, (self).f26, ((d_634_v13_)[d_633_i2_] if (d_633_i2_) in (d_634_v13_) else default__.fm5((self).f25, iife45_()
, len(d_623_v11_), globalState)))
            d_640_v19_: _dafny.Map
            d_640_v19_ = _dafny.Map({len(d_635_v14_): (0) - ((d_637_v18_).cf4)})
            d_641_v20_: _dafny.Map
            d_641_v20_ = _dafny.Map({d_623_v11_: d_623_v11_})
            d_640_v19_ = (d_640_v19_).set((self).f26, (self).fm11(d_641_v20_, (self).f31, (self).f31, globalState))
            (globalState).f0 = ((281) * ((self).f31)) * ((self).f31)
            d_642_v21_: _dafny.Map
            d_642_v21_ = _dafny.Map({p2: ((p0).cardinality) != (len(d_623_v11_))})
            d_642_v21_ = (d_642_v21_).set(p2, not(p1))
            d_640_v19_ = (d_640_v19_).set((len(d_642_v21_)) - ((self).f31), (d_633_i2_) * ((self).f31))
        d_643_v22_: _dafny.Seq
        d_643_v22_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f31), 537])
        d_644_v23_: _dafny.Map
        d_644_v23_ = _dafny.Map({p1: len(d_643_v22_)})
        d_645_v24_: _dafny.Set
        d_645_v24_ = _dafny.Set({D2_DC6(default__.fm12(len(d_644_v23_), (self).f31, globalState))})
        d_646_v25_: _dafny.Set
        d_646_v25_ = _dafny.Set({(self).f31})
        d_647_v26_: D2
        d_647_v26_ = D2_DC6(d_646_v25_)
        d_645_v24_ = (d_645_v24_) - (_dafny.Set({d_647_v26_, D2_DC6(_dafny.Set({(self).f26, 969, (self).f31, (self).f31})), d_647_v26_}))
        hi6_ = ((self).f26) + ((self).f26)
        for d_648_i3_ in range((self).f26, hi6_):
            (globalState).f0 = -20
            d_649_v27_: _dafny.Array
            nw113_ = _dafny.Array(_dafny.Set({}), 22)
            d_649_v27_ = nw113_
            d_650_v28_: _dafny.Seq
            d_650_v28_ = _dafny.SeqWithoutIsStrInference([(self).f25])
            d_651_v29_: _dafny.Set
            d_651_v29_ = _dafny.Set({d_650_v28_, d_650_v28_})
            index120_ = default__.safeIndex(498, (d_649_v27_).length(0))
            (d_649_v27_)[index120_] = d_651_v29_
            index121_ = default__.safeIndex(498, (d_649_v27_).length(0))
            (d_649_v27_)[index121_] = default__.fm13(p2, globalState)
            d_652_v30_: D0
            d_652_v30_ = D0_DC2()
            source10_ = d_652_v30_
            if source10_.is_DC1:
                d_653___mcc_h4_ = source10_.cf1
                d_654___mcc_h5_ = source10_.cf2
                d_655___mcc_h6_ = source10_.cf3
                d_656___mcc_h7_ = source10_.cf4
                d_657___mcc_h8_ = source10_.cf5
                d_658_cf5_ = d_657___mcc_h8_
                d_659_cf4_ = d_656___mcc_h7_
                d_660_cf3_ = d_655___mcc_h6_
                d_661_cf2_ = d_654___mcc_h5_
                d_662_cf1_ = d_653___mcc_h4_
                d_663_v31_: _dafny.Array
                nw114_ = _dafny.Array(_dafny.MultiSet({}), 12)
                d_663_v31_ = nw114_
                d_664_v32_: _dafny.MultiSet
                d_664_v32_ = _dafny.MultiSet([d_644_v23_])
                index122_ = default__.safeIndex(274, (d_663_v31_).length(0))
                (d_663_v31_)[index122_] = (_dafny.MultiSet([d_644_v23_])) | (d_664_v32_)
                d_665_v33_: _dafny.MultiSet
                d_665_v33_ = d_664_v32_
                index123_ = default__.safeIndex(274, (d_663_v31_).length(0))
                (d_663_v31_)[index123_] = (d_665_v33_)
                d_666_v34_: D2
                d_666_v34_ = D2_DC7((self).f25, p2, (self).f31)
                d_667_v35_: _dafny.Array
                nw115_ = _dafny.Array(None, 7)
                nw115_[int(0)] = d_658_cf5_
                nw115_[int(1)] = d_658_cf5_
                nw115_[int(2)] = True
                nw115_[int(3)] = (d_658_cf5_) and (p2)
                nw115_[int(4)] = (d_666_v34_).cf12
                nw115_[int(5)] = d_658_cf5_
                nw115_[int(6)] = p2
                d_667_v35_ = nw115_
                index124_ = default__.safeIndex(607, (d_667_v35_).length(0))
                (d_667_v35_)[index124_] = p1
                index125_ = default__.safeIndex(607, (d_667_v35_).length(0))
                (d_667_v35_)[index125_] = p1
                d_668_v36_: str
                d_668_v36_ = _dafny.CodePoint('c')
                d_668_v36_ = d_668_v36_
                d_669_v37_: _dafny.Map
                d_669_v37_ = _dafny.Map({d_659_cf4_: d_662_cf1_})
                d_670_v38_: _dafny.Map
                d_670_v38_ = _dafny.Map({d_669_v37_: d_667_v35_})
                rhs144_ = (d_670_v38_) != ((d_670_v38_) | (d_670_v38_))
                rhs145_ = _dafny.Set({d_662_cf1_, (self).f26, default__.safeDivisionInt(960, d_659_cf4_)})
                rhs146_ = ((self).f26 if (d_667_v35_)[default__.safeIndex(607, (d_667_v35_).length(0))] else ((self).f31) - (d_659_cf4_))
                rhs147_ = (self).f31
                rhs148_ = False
                lhs128_ = globalState
                lhs129_ = globalState
                lhs130_ = globalState
                lhs128_.f1 = rhs144_
                d_646_v25_ = rhs145_
                lhs129_.f2 = rhs146_
                lhs130_.f0 = rhs147_
                d_660_cf3_ = rhs148_
            elif source10_.is_DC2:
                d_671_v39_: _dafny.Seq
                d_671_v39_ = _dafny.SeqWithoutIsStrInference([d_623_v11_])
                d_672_v40_: _dafny.Map
                d_672_v40_ = _dafny.Map({(self).f26: (self).f31})
                def iife48_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(914, -407):
                        d_673_v41_: int = compr_16_
                        if ((914) <= (d_673_v41_)) and ((d_673_v41_) < (-407)):
                            coll16_[(d_673_v41_) + ((self).f31)] = (0) - ((self).f31)
                    return _dafny.Map(coll16_)
                d_643_v22_ = default__.fm3(d_648_i3_, not(((0) - ((self).f31)) == (len(d_671_v39_))), 766, (d_672_v40_) != (iife48_()
                ), globalState)
                (globalState).f1 = (self).fm7(globalState)
                d_674_v42_: C1
                nw116_ = C1()
                nw116_.ctor__(p2)
                d_674_v42_ = nw116_
                (globalState).f1 = (d_650_v28_)[default__.safeIndex((self).f26, len(d_650_v28_))]
            elif True:
                d_675___mcc_h9_ = source10_.cf0
                d_676_cf0_ = d_675___mcc_h9_
                d_677_v43_: D0
                d_677_v43_ = D0_DC0(d_676_cf0_)
                d_678_v44_: _dafny.Array
                nw117_ = _dafny.Array(False, 6)
                d_678_v44_ = nw117_
                d_679_v45_: C0
                nw118_ = C0()
                nw118_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "me")), (d_677_v43_).cf0, d_648_i3_, d_678_v44_, False)
                d_679_v45_ = nw118_
                d_680_v46_: C1
                nw119_ = C1()
                nw119_.ctor__((p1 if p1 else not(p2)))
                d_680_v46_ = nw119_
                d_650_v28_ = d_650_v28_
                d_681_v47_: _dafny.Seq
                d_681_v47_ = d_650_v28_
                d_682_v48_: _dafny.Seq
                d_682_v48_ = _dafny.SeqWithoutIsStrInference([d_650_v28_, d_681_v47_, d_681_v47_, d_681_v47_, d_681_v47_])
                d_683_v49_: _dafny.Seq
                d_683_v49_ = _dafny.SeqWithoutIsStrInference([d_678_v44_])
                (globalState).f1 = (d_682_v48_) == (((_dafny.SeqWithoutIsStrInference([d_681_v47_, d_681_v47_, d_681_v47_, d_681_v47_])).set(default__.safeIndex(len(d_683_v49_), len(_dafny.SeqWithoutIsStrInference([d_681_v47_, d_681_v47_, d_681_v47_, d_681_v47_]))), _dafny.SeqWithoutIsStrInference([p1]))) + (d_682_v48_))
            d_684_v50_: _dafny.MultiSet
            d_684_v50_ = _dafny.MultiSet([not(p1), True])
            d_685_v51_: _dafny.Map
            d_685_v51_ = _dafny.Map({D0_DC0(p2): (d_684_v50_).cardinality})
            d_686_v52_: _dafny.Seq
            d_686_v52_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_623_v11_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ferpotds"))), len(d_623_v11_))] for d_687_i4_ in range(default__.abs(928))])])
            if (default__.fm5(p1, d_685_v51_, (self).f31, globalState)) and ((-875) > (len((d_686_v52_)[default__.safeIndex((self).f26, len(d_686_v52_))]))):
                d_688_v53_: _dafny.Map
                d_688_v53_ = _dafny.Map({d_648_i3_: (d_650_v28_)[default__.safeIndex(d_648_i3_, len(d_650_v28_))]})
                d_689_v54_: str
                d_689_v54_ = _dafny.CodePoint('r')
                d_690_v55_: _dafny.Array
                nw120_ = _dafny.Array(None, 12)
                nw120_[int(0)] = (0) - ((self).f26)
                nw120_[int(1)] = (self).f31
                nw120_[int(2)] = (self).f26
                nw120_[int(3)] = len(((d_688_v53_).set((self).f31, (self).f25)) | (d_688_v53_))
                nw120_[int(4)] = default__.safeModuloInt((0) - ((self).f26), (self).f26)
                nw120_[int(5)] = ((p0)[926] if (926) in (p0) else (self).f31)
                nw120_[int(6)] = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))).set(default__.safeIndex((self).f31, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))), d_689_v54_)).set(default__.safeIndex((self).f26, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))).set(default__.safeIndex((self).f31, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))), d_689_v54_))), _dafny.CodePoint('s')))
                nw120_[int(7)] = d_648_i3_
                nw120_[int(8)] = default__.safeDivisionInt((self).f26, (self).f31)
                nw120_[int(9)] = d_648_i3_
                nw120_[int(10)] = (self).f31
                nw120_[int(11)] = default__.fm2(d_648_i3_, -150, globalState)
                d_690_v55_ = nw120_
                index126_ = default__.safeIndex(590, (d_690_v55_).length(0))
                (d_690_v55_)[index126_] = ((self).f31 if p1 else len(_dafny.SeqWithoutIsStrInference([(self).f26, d_648_i3_])))
                index127_ = default__.safeIndex(590, (d_690_v55_).length(0))
                (d_690_v55_)[index127_] = d_648_i3_
                d_691_v56_: _dafny.Array
                def lambda37_(d_692_p2_):
                    def lambda38_(d_693_i5_):
                        return d_692_p2_

                    return lambda38_

                init18_ = lambda37_(p2)
                nw121_ = _dafny.Array(None, 23)
                for i0_18_ in range(nw121_.length(0)):
                    nw121_[i0_18_] = init18_(i0_18_)
                d_691_v56_ = nw121_
                d_694_v57_: C0
                nw122_ = C0()
                nw122_.ctor__(d_623_v11_, (self).f25, 296, d_691_v56_, p2)
                d_694_v57_ = nw122_
                index128_ = default__.safeIndex(555, (d_691_v56_).length(0))
                (d_691_v56_)[index128_] = p2
                d_695_v58_: D0
                d_695_v58_ = D0_DC0(p2)
                pat_let_tv20_ = p1
                index129_ = default__.safeIndex(555, (d_691_v56_).length(0))
                def iife49_(_pat_let16_0):
                    def iife50_(d_696_dt__update__tmp_h0_):
                        def iife51_(_pat_let17_0):
                            def iife52_(d_697_dt__update_hcf0_h0_):
                                return D0_DC0(d_697_dt__update_hcf0_h0_)
                            return iife52_(_pat_let17_0)
                        return iife51_(pat_let_tv20_)
                    return iife50_(_pat_let16_0)
                (d_691_v56_)[index129_] = ((self).f25) and (default__.fm5((self).f25, (d_685_v51_).set(iife49_(d_695_v58_), (d_690_v55_)[default__.safeIndex(590, (d_690_v55_).length(0))]), (self).f26, globalState))
                d_698_v59_: _dafny.Seq
                d_698_v59_ = _dafny.SeqWithoutIsStrInference([d_691_v56_])
                d_699_v60_: _dafny.Seq
                d_699_v60_ = _dafny.SeqWithoutIsStrInference([(d_698_v59_)[default__.safeIndex((self).f31, len(d_698_v59_))]])
                d_699_v60_ = ((d_698_v59_) + (d_699_v60_) if (p0).ispropersubset(p0) else d_699_v60_)
                (globalState).f2 = d_648_i3_
            elif True:
                d_700_v61_: _dafny.Array
                nw123_ = _dafny.Array(None, 5)
                nw123_[int(0)] = (self).f25
                nw123_[int(1)] = (self).f25
                nw123_[int(2)] = p1
                nw123_[int(3)] = not(p1)
                nw123_[int(4)] = p1
                d_700_v61_ = nw123_
                d_701_v62_: T0
                nw124_ = C0()
                nw124_.ctor__(d_623_v11_, p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kiivsm"))), d_700_v61_, p2)
                d_701_v62_ = nw124_
                d_702_v63_: _dafny.Seq
                d_702_v63_ = _dafny.SeqWithoutIsStrInference([d_701_v62_])
                d_703_v64_: _dafny.Map
                d_703_v64_ = _dafny.Map({True: d_702_v63_})
                d_704_v65_: _dafny.Map
                d_704_v65_ = _dafny.Map({False: ((d_703_v64_)[(d_701_v62_).f25] if ((d_701_v62_).f25) in (d_703_v64_) else d_702_v63_)})
                d_705_v66_: D9
                d_705_v66_ = D9_DC18(d_703_v64_)
                d_704_v65_ = (((d_705_v66_).cf30).set((d_701_v62_).f25, d_702_v63_)) | (d_703_v64_)
                (globalState).f1 = (True if (d_701_v62_).f25 else (d_623_v11_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmubls"))))
                (globalState).f2 = len(d_650_v28_)
                d_706_v67_: int
                out16_: int
                out16_ = default__.m0(p1, globalState)
                d_706_v67_ = out16_
                d_707_v68_: _dafny.Map
                d_707_v68_ = _dafny.Map({d_623_v11_: d_706_v67_})
                (globalState).f1 = (_dafny.Map({d_623_v11_: (self).fm6((self).f25, (self).f25, d_648_i3_, globalState)})) == (d_707_v68_)
        r0 = (self).fm7(globalState)
        return r0

    def m5(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        d_708_v0_: _dafny.Set
        d_708_v0_ = _dafny.Set({p0, (self).f25, (self).f25, (self).f25, not ((self).f25) or (not(p0))})
        d_709_v1_: _dafny.Array
        nw125_ = _dafny.Array(False, 4)
        d_709_v1_ = nw125_
        index130_ = default__.safeIndex(303, (d_709_v1_).length(0))
        (d_709_v1_)[index130_] = p0
        d_710_v2_: D0
        d_710_v2_ = D0_DC0(p0)
        d_711_v3_: _dafny.Map
        d_711_v3_ = _dafny.Map({default__.fm5((self).f25, default__.fm25((self).f26, (d_710_v2_).cf0, p0, globalState), 345, globalState): d_708_v0_})
        d_712_v4_: _dafny.Seq
        d_712_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nf"))
        d_713_v5_: _dafny.Seq
        d_713_v5_ = _dafny.SeqWithoutIsStrInference([default__.fm26((self).f26, p2, p0, globalState)])
        d_714_v6_: _dafny.Map
        d_714_v6_ = _dafny.Map({(self).f25: p2})
        d_715_v7_: _dafny.MultiSet
        d_715_v7_ = _dafny.MultiSet([d_714_v6_, d_714_v6_])
        d_716_v8_: _dafny.Map
        d_716_v8_ = _dafny.Map({d_710_v2_: (self).f31})
        index131_ = default__.safeIndex(303, (d_709_v1_).length(0))
        rhs149_ = ((d_711_v3_)[not (p0) or ((self).f25)] if (not (p0) or ((self).f25)) in (d_711_v3_) else d_708_v0_)
        rhs150_ = (0) - ((len((d_712_v4_) + (d_712_v4_)) if (self).f25 else (self).f31))
        rhs151_ = ((d_715_v7_).set(d_714_v6_, default__.abs((self).f31))).ispropersubset((d_713_v5_)[default__.safeIndex((0) - ((self).f31), len(d_713_v5_))])
        rhs152_ = ((p2) in (_dafny.SeqWithoutIsStrInference([p2 for d_717_i0_ in range(default__.abs(263))])) if default__.fm5((self).f25, d_716_v8_, (self).f31, globalState) else (self).f25)
        lhs131_ = globalState
        lhs132_ = globalState
        lhs133_ = d_709_v1_
        lhs134_ = default__.safeIndex(303, (d_709_v1_).length(0))
        d_708_v0_ = rhs149_
        lhs131_.f0 = rhs150_
        lhs132_.f1 = rhs151_
        lhs133_[lhs134_] = rhs152_
        hi7_ = (self).f31
        for d_718_i1_ in range((self).f31, hi7_):
            d_719_v9_: _dafny.Array
            def lambda39_(d_720_v4_):
                def lambda40_(d_721_i2_):
                    return d_720_v4_

                return lambda40_

            init19_ = lambda39_(d_712_v4_)
            nw126_ = _dafny.Array(None, 19)
            for i0_19_ in range(nw126_.length(0)):
                nw126_[i0_19_] = init19_(i0_19_)
            d_719_v9_ = nw126_
            d_722_v10_: _dafny.Seq
            d_722_v10_ = _dafny.SeqWithoutIsStrInference([(self).f31])
            d_723_v11_: _dafny.Map
            d_723_v11_ = _dafny.Map({d_719_v9_: d_722_v10_})
            d_723_v11_ = (d_723_v11_).set(d_719_v9_, d_722_v10_)
            d_724_v12_: _dafny.Map
            d_724_v12_ = _dafny.Map({d_722_v10_: p2})
            d_725_v13_: C0
            nw127_ = C0()
            nw127_.ctor__(d_712_v4_, not(((self).f26) > ((self).f31)), len(((d_712_v4_).set(default__.safeIndex(438, len(d_712_v4_)), ((d_724_v12_)[d_722_v10_] if (d_722_v10_) in (d_724_v12_) else p2))).set(default__.safeIndex(d_718_i1_, len((d_712_v4_).set(default__.safeIndex(438, len(d_712_v4_)), ((d_724_v12_)[d_722_v10_] if (d_722_v10_) in (d_724_v12_) else p2)))), p2)), d_709_v1_, not((self).f25))
            d_725_v13_ = nw127_
            index132_ = default__.safeIndex(303, (d_709_v1_).length(0))
            (d_709_v1_)[index132_] = not((self).f25)
            d_726_v14_: _dafny.Map
            d_726_v14_ = _dafny.Map({(d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))]: d_718_i1_})
            d_727_v15_: _dafny.MultiSet
            d_727_v15_ = _dafny.MultiSet([_dafny.Map({(d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))]: (self).f31}), (d_726_v14_).set((d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))], -691)])
            d_728_v16_: _dafny.MultiSet
            d_728_v16_ = d_727_v15_
            source11_ = _dafny.MultiSet([d_726_v14_])
            d_729___mcc_h0_ = source11_
            d_730_cf14_ = d_729___mcc_h0_
            d_731_v17_: _dafny.MultiSet
            d_731_v17_ = _dafny.MultiSet([p0])
            d_732_v18_: C0
            nw128_ = C0()
            nw128_.ctor__(d_725_v13_.f35, (d_731_v17_).isdisjoint(d_731_v17_), len((d_708_v0_).intersection(_dafny.Set({p0}))), d_709_v1_, (default__.fm27((self).f26, globalState)).isdisjoint(d_731_v17_))
            d_732_v18_ = nw128_
            d_733_v19_: _dafny.Seq
            d_733_v19_ = _dafny.SeqWithoutIsStrInference([not(p0)])
            index133_ = default__.safeIndex(303, (d_709_v1_).length(0))
            rhs153_ = d_733_v19_
            rhs154_ = d_718_i1_
            rhs155_ = ((d_731_v17_).intersection(d_731_v17_)).ispropersubset(d_731_v17_)
            lhs135_ = globalState
            lhs136_ = d_709_v1_
            lhs137_ = default__.safeIndex(303, (d_709_v1_).length(0))
            d_733_v19_ = rhs153_
            lhs135_.f0 = rhs154_
            lhs136_[lhs137_] = rhs155_
            (globalState).f2 = default__.safeDivisionInt(((self).f26) - ((self).f26), (self).f31)
            d_734_v20_: _dafny.Array
            nw129_ = _dafny.Array(int(0), 10)
            d_734_v20_ = nw129_
            index134_ = default__.safeIndex(28, (d_734_v20_).length(0))
            (d_734_v20_)[index134_] = (self).f26
            index135_ = default__.safeIndex(28, (d_734_v20_).length(0))
            (d_734_v20_)[index135_] = (self).f26
        d_735_v21_: _dafny.Array
        nw130_ = _dafny.Array(_dafny.Array(None, 0), 12)
        d_735_v21_ = nw130_
        d_736_v22_: _dafny.MultiSet
        d_736_v22_ = _dafny.MultiSet([(self).f25, (self).f25, (d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))]])
        d_737_v23_: _dafny.Seq
        d_737_v23_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_738_v24_: _dafny.Set
        d_738_v24_ = _dafny.Set({(self).f31, (self).f31})
        d_739_v25_: _dafny.Map
        d_739_v25_ = _dafny.Map({(_dafny.MultiSet(d_737_v23_)).cardinality: d_738_v24_})
        d_740_v26_: D6
        d_740_v26_ = D6_DC13((self).f31, (self).f31, ((d_739_v25_)[(self).f26] if ((self).f26) in (d_739_v25_) else _dafny.Set({(self).f26})))
        d_741_v27_: _dafny.Array
        nw131_ = _dafny.Array(None, 2)
        d_741_v27_ = nw131_
        d_742_v28_: _dafny.Seq
        d_742_v28_ = _dafny.SeqWithoutIsStrInference([d_741_v27_, d_741_v27_])
        d_743_v29_: _dafny.Array
        nw132_ = _dafny.Array(None, 10)
        nw132_[int(0)] = (self).f26
        nw132_[int(1)] = ((d_736_v22_)[p0] if (p0) in (d_736_v22_) else (self).f26)
        nw132_[int(2)] = len(d_712_v4_)
        nw132_[int(3)] = 880
        nw132_[int(4)] = (self).f31
        nw132_[int(5)] = (self).f31
        nw132_[int(6)] = len(d_712_v4_)
        nw132_[int(7)] = (self).f31
        nw132_[int(8)] = (d_740_v26_).cf20
        nw132_[int(9)] = len(d_742_v28_)
        d_743_v29_ = nw132_
        index136_ = default__.safeIndex(236, (d_735_v21_).length(0))
        (d_735_v21_)[index136_] = d_743_v29_
        index137_ = default__.safeIndex(236, (d_735_v21_).length(0))
        (d_735_v21_)[index137_] = d_743_v29_
        hi8_ = len(_dafny.Map({(self).f31: (self).f26}))
        for d_744_i3_ in range((self).f31, hi8_):
            (globalState).f0 = d_744_i3_
            arr2_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            index138_ = default__.safeIndex(170, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            arr2_[index138_] = d_744_i3_
            arr3_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            index139_ = default__.safeIndex(77, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            arr3_[index139_] = (self).f26
            d_745_v30_: _dafny.Map
            d_745_v30_ = _dafny.Map({p2: default__.fm5((d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))], d_716_v8_, (self).f26, globalState)})
            d_746_v31_: _dafny.Set
            d_746_v31_ = _dafny.Set({d_745_v30_, d_745_v30_, d_745_v30_})
            d_747_v32_: _dafny.MultiSet
            d_747_v32_ = _dafny.MultiSet([(0) - (len(d_746_v31_)), (self).f26, -489])
            arr4_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            index140_ = default__.safeIndex(170, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            arr5_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            index141_ = default__.safeIndex(77, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            rhs156_ = (self).f26
            rhs157_ = ((d_736_v22_)[p0] if (p0) in (d_736_v22_) else (self).f31)
            rhs158_ = ((self).f31) * (d_744_i3_)
            rhs159_ = default__.safeModuloInt(d_744_i3_, (d_747_v32_).cardinality)
            lhs138_ = globalState
            lhs139_ = globalState
            lhs140_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            lhs141_ = default__.safeIndex(170, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            lhs142_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            lhs143_ = default__.safeIndex(77, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            lhs138_.f2 = rhs156_
            lhs139_.f0 = rhs157_
            lhs140_[lhs141_] = rhs158_
            lhs142_[lhs143_] = rhs159_
            d_748_v33_: C1
            nw133_ = C1()
            nw133_.ctor__((d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))])
            d_748_v33_ = nw133_
            d_749_v35_: _dafny.Map
            d_749_v35_ = _dafny.Map({d_712_v4_: d_712_v4_})
            d_750_v36_: _dafny.Map
            d_750_v36_ = _dafny.Map({(self).fm11(d_749_v35_, len(_dafny.SeqWithoutIsStrInference([D9_DC19(D6_DC14(len(_dafny.Map({d_744_i3_: (self).f31})), (self).f31, d_736_v22_)) for d_751_i4_ in range(default__.abs(169))])), d_744_i3_, globalState): (d_747_v32_).cardinality})
            d_752_v37_: bool
            out17_: bool
            def iife53_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in ((d_747_v32_).set((self).f26, default__.abs((self).f31))).Elements:
                    d_753_v34_: int = compr_17_
                    if (d_753_v34_) in ((d_747_v32_).set((self).f26, default__.abs((self).f31))):
                        coll17_[(d_753_v34_) + (((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))])[default__.safeIndex(170, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))])] = d_744_i3_
                return _dafny.Map(coll17_)
            out17_ = (self).m6((iife53_()
            ) | (d_750_v36_), ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))])[default__.safeIndex(170, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))], (self).f31, globalState)
            d_752_v37_ = out17_
        (globalState).f0 = (0) - ((self).f26)
        d_754_v39_: _dafny.Seq
        d_754_v39_ = _dafny.SeqWithoutIsStrInference([801])
        d_755_v40_: D0
        def iife54_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in (d_754_v39_).Elements:
                d_756_v38_: int = compr_18_
                if (d_756_v38_) in (d_754_v39_):
                    coll18_[(d_756_v38_) - (619)] = p0
            return _dafny.Map(coll18_)
        d_755_v40_ = D0_DC1((self).f26, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ac")), (d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))], len(d_737_v23_), default__.fm5((d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))], d_716_v8_, len(iife54_()
), globalState))
        d_757_v41_: _dafny.Map
        d_757_v41_ = _dafny.Map({default__.fm2(len(d_712_v4_), len(d_737_v23_), globalState): p0})
        pat_let_tv21_ = d_757_v41_
        pat_let_tv22_ = p0
        def iife55_(_pat_let18_0):
            def iife56_(d_758_dt__update__tmp_h1_):
                def iife57_(_pat_let19_0):
                    def iife58_(d_759_dt__update_hcf1_h0_):
                        def iife59_(_pat_let20_0):
                            def iife60_(d_760_dt__update_hcf5_h0_):
                                def iife61_(_pat_let21_0):
                                    def iife62_(d_761_dt__update_hcf4_h0_):
                                        return D0_DC1(d_759_dt__update_hcf1_h0_, (d_758_dt__update__tmp_h1_).cf2, (d_758_dt__update__tmp_h1_).cf3, d_761_dt__update_hcf4_h0_, d_760_dt__update_hcf5_h0_)
                                    return iife62_(_pat_let21_0)
                                return iife61_((self).f31)
                            return iife60_(_pat_let20_0)
                        return iife59_(pat_let_tv22_)
                    return iife58_(_pat_let19_0)
                return iife57_(len(pat_let_tv21_))
            return iife56_(_pat_let18_0)
        if (len((iife55_(d_755_v40_)).cf2)) >= ((self).f26):
            index142_ = default__.safeIndex(303, (d_709_v1_).length(0))
            (d_709_v1_)[index142_] = ((0) - ((self).f31)) != ((self).f31)
            (globalState).f1 = (self).f25
            (globalState).f0 = default__.safeModuloInt((self).f26, ((0) - ((self).f26)) + ((self).f31))
            arr6_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            index143_ = default__.safeIndex(189, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            arr6_[index143_] = (self).f31
            arr7_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            index144_ = default__.safeIndex(189, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            arr7_[index144_] = (self).f26
            d_762_v42_: _dafny.Map
            d_762_v42_ = _dafny.Map({(self).f25: (_dafny.MultiSet(default__.fm23(globalState))).cardinality})
            d_762_v42_ = (d_762_v42_).set(p0, (self).f31)
        elif True:
            (globalState).f0 = 432
            (globalState).f0 = ((len(d_712_v4_)) * ((self).f26)) + (800)
            d_763_v43_: D6
            d_763_v43_ = D6_DC14(len(d_712_v4_), (self).f26, d_736_v22_)
            d_764_v44_: D9
            d_764_v44_ = D9_DC19(d_763_v43_)
            d_765_v45_: D9
            d_765_v45_ = D9_DC20(d_764_v44_)
            d_766_v46_: D9
            d_766_v46_ = D9_DC20(d_765_v45_)
            d_767_v47_: _dafny.Map
            d_767_v47_ = _dafny.Map({d_743_v29_: d_766_v46_})
            d_768_v48_: _dafny.Map
            d_768_v48_ = _dafny.Map({(d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))]: d_767_v47_})
            d_768_v48_ = (d_768_v48_).set((d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))], d_767_v47_)
            d_769_v49_: C1
            nw134_ = C1()
            nw134_.ctor__((self).f25)
            d_769_v49_ = nw134_
            arr8_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            index145_ = default__.safeIndex(294, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            arr8_[index145_] = (d_736_v22_).cardinality
            arr9_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            index146_ = default__.safeIndex(294, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            index147_ = default__.safeIndex(236, (d_735_v21_).length(0))
            rhs160_ = d_769_v49_
            rhs161_ = not(p0)
            rhs162_ = (self).f26
            rhs163_ = d_743_v29_
            lhs144_ = globalState
            lhs145_ = (d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]
            lhs146_ = default__.safeIndex(294, ((d_735_v21_)[default__.safeIndex(236, (d_735_v21_).length(0))]).length(0))
            lhs147_ = d_735_v21_
            lhs148_ = default__.safeIndex(236, (d_735_v21_).length(0))
            d_769_v49_ = rhs160_
            lhs144_.f1 = rhs161_
            lhs145_[lhs146_] = rhs162_
            lhs147_[lhs148_] = rhs163_
            d_770_v50_: _dafny.Map
            d_770_v50_ = _dafny.Map({(self).fm7(globalState): (d_769_v49_).f32})
            d_771_v51_: _dafny.Map
            d_771_v51_ = _dafny.Map({(d_754_v39_)[default__.safeIndex((self).f31, len(d_754_v39_))]: d_770_v50_})
            d_772_v52_: D9
            d_772_v52_ = D9_DC19(d_763_v43_)
            d_773_v53_: _dafny.Map
            d_773_v53_ = _dafny.Map({(d_770_v50_ if not((d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))]) else ((d_771_v51_)[(self).f31] if ((self).f31) in (d_771_v51_) else d_770_v50_)): d_772_v52_})
            d_773_v53_ = d_773_v53_
        r0 = not((d_737_v23_) == (_dafny.SeqWithoutIsStrInference([(d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))], (d_709_v1_)[default__.safeIndex(303, (d_709_v1_).length(0))]])))
        d_774_v54_: _dafny.MultiSet
        d_774_v54_ = _dafny.MultiSet([(self).fm6(False, False, (self).f31, globalState), (self).f26, (self).f31])
        pat_let_tv23_ = p0
        def lambda41_(source12_):
            d_775___mcc_h1_ = source12_
            d_776_cf18_ = d_775___mcc_h1_
            return pat_let_tv23_

        r1 = lambda41_(d_774_v54_)
        return r0, r1

    def m6(self, p0, p1, p2, globalState):
        r0: bool = False
        d_777_v0_: str
        d_777_v0_ = _dafny.CodePoint('y')
        d_777_v0_ = d_777_v0_
        if ((self).f25 if False else default__.fm5((self).f25, _dafny.Map({D0_DC0((self).f25): p2}), 726, globalState)):
            d_778_v1_: _dafny.MultiSet
            d_778_v1_ = _dafny.MultiSet([-40])
            d_779_v2_: _dafny.Seq
            d_779_v2_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2, (0) - (((d_778_v1_)[7] if (7) in (d_778_v1_) else p2))])
            d_780_v3_: _dafny.Set
            d_780_v3_ = _dafny.Set({(self).f26, (self).f31})
            d_781_v4_: D6
            d_781_v4_ = D6_DC13((self).f26, (self).f26, d_780_v3_)
            d_782_v5_: _dafny.Seq
            d_782_v5_ = _dafny.SeqWithoutIsStrInference([d_779_v2_, d_779_v2_, _dafny.SeqWithoutIsStrInference([(self).f26 for d_783_i1_ in range(default__.abs(218))])])
            d_784_v6_: _dafny.Seq
            d_784_v6_ = _dafny.SeqWithoutIsStrInference([(self).f25])
            d_785_v7_: _dafny.Array
            nw135_ = _dafny.Array(None, 26)
            nw135_[int(0)] = (d_779_v2_).set(default__.safeIndex((self).f26, len(d_779_v2_)), (self).f26)
            nw135_[int(1)] = (d_779_v2_) + (d_779_v2_)
            nw135_[int(2)] = (d_779_v2_) + (d_779_v2_)
            nw135_[int(3)] = (d_779_v2_) + (d_779_v2_)
            nw135_[int(4)] = d_779_v2_
            nw135_[int(5)] = d_779_v2_
            nw135_[int(6)] = default__.fm3(p1, False, p2, (self).f25, globalState)
            nw135_[int(7)] = d_779_v2_
            nw135_[int(8)] = (d_779_v2_) + (d_779_v2_)
            nw135_[int(9)] = d_779_v2_
            nw135_[int(10)] = d_779_v2_
            nw135_[int(11)] = _dafny.SeqWithoutIsStrInference([p2, p2, p1])
            nw135_[int(12)] = d_779_v2_
            nw135_[int(13)] = d_779_v2_
            nw135_[int(14)] = d_779_v2_
            nw135_[int(15)] = d_779_v2_
            nw135_[int(16)] = ((d_779_v2_) + (d_779_v2_)).set(default__.safeIndex((d_781_v4_).cf20, len((d_779_v2_) + (d_779_v2_))), 383)
            nw135_[int(17)] = (d_779_v2_) + (d_779_v2_)
            nw135_[int(18)] = d_779_v2_
            nw135_[int(19)] = (_dafny.SeqWithoutIsStrInference([327, (self).f31, p2, (self).f26])) + (_dafny.SeqWithoutIsStrInference([(self).f26 for d_786_i0_ in range(default__.abs(-679))]))
            nw135_[int(20)] = _dafny.SeqWithoutIsStrInference([(self).f26, p1, p1])
            nw135_[int(21)] = ((d_779_v2_) + (d_779_v2_)).set(default__.safeIndex(296, len((d_779_v2_) + (d_779_v2_))), p1)
            nw135_[int(22)] = (d_782_v5_)[default__.safeIndex(len(d_784_v6_), len(d_782_v5_))]
            nw135_[int(23)] = (d_779_v2_) + (_dafny.SeqWithoutIsStrInference([(self).f26 for d_787_i2_ in range(default__.abs(723))]))
            nw135_[int(24)] = d_779_v2_
            nw135_[int(25)] = d_779_v2_
            d_785_v7_ = nw135_
            index148_ = default__.safeIndex(161, (d_785_v7_).length(0))
            (d_785_v7_)[index148_] = d_779_v2_
            d_788_v8_: _dafny.MultiSet
            d_788_v8_ = _dafny.MultiSet([True])
            index149_ = default__.safeIndex(161, (d_785_v7_).length(0))
            (d_785_v7_)[index149_] = ((d_782_v5_)[default__.safeIndex(((self).f31) * ((self).f26), len(d_782_v5_))]).set(default__.safeIndex(((d_788_v8_).set((self).f25, default__.abs((0) - ((self).f26)))).cardinality, len((d_782_v5_)[default__.safeIndex(((self).f31) * ((self).f26), len(d_782_v5_))])), (self).f31)
            d_789_v9_: int
            out18_: int
            out18_ = default__.m0(not((self).f25), globalState)
            d_789_v9_ = out18_
            d_790_v10_: int
            out19_: int
            out19_ = default__.m0((self).f25, globalState)
            d_790_v10_ = out19_
            if not((self).f25):
                d_791_v11_: _dafny.Seq
                d_791_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxl"))
                d_792_v12_: _dafny.Array
                nw136_ = _dafny.Array(None, 28)
                nw136_[int(0)] = (self).f25
                nw136_[int(1)] = not((self).f25)
                nw136_[int(2)] = (self).f25
                nw136_[int(3)] = not((self).f25)
                nw136_[int(4)] = (self).f25
                nw136_[int(5)] = (self).f25
                nw136_[int(6)] = (self).f25
                nw136_[int(7)] = (self).f25
                nw136_[int(8)] = (self).f25
                nw136_[int(9)] = (self).f25
                nw136_[int(10)] = (self).f25
                nw136_[int(11)] = (self).f25
                nw136_[int(12)] = (self).f25
                nw136_[int(13)] = (self).f25
                nw136_[int(14)] = (self).f25
                nw136_[int(15)] = (self).f25
                nw136_[int(16)] = (self).f25
                nw136_[int(17)] = (self).f25
                nw136_[int(18)] = (self).f25
                nw136_[int(19)] = (self).f25
                nw136_[int(20)] = (self).f25
                nw136_[int(21)] = (self).f25
                nw136_[int(22)] = False
                nw136_[int(23)] = (self).f25
                nw136_[int(24)] = (self).f25
                nw136_[int(25)] = True
                nw136_[int(26)] = (self).f25
                nw136_[int(27)] = (self).f25
                d_792_v12_ = nw136_
                d_793_v13_: C0
                nw137_ = C0()
                nw137_.ctor__(d_791_v11_, (self).f25, (self).f26, d_792_v12_, not((self).f25))
                d_793_v13_ = nw137_
                d_794_v14_: C1
                nw138_ = C1()
                nw138_.ctor__((False if (self).f25 else (self).f25))
                d_794_v14_ = nw138_
                d_795_v15_: _dafny.Array
                nw139_ = _dafny.Array(int(0), 10)
                d_795_v15_ = nw139_
                index150_ = default__.safeIndex(408, (d_795_v15_).length(0))
                (d_795_v15_)[index150_] = (self).f31
                index151_ = default__.safeIndex(408, (d_795_v15_).length(0))
                (d_795_v15_)[index151_] = p1
                d_796_v16_: _dafny.Set
                d_796_v16_ = _dafny.Set({d_795_v15_, d_795_v15_})
                d_797_v17_: D6
                d_797_v17_ = D6_DC12(d_796_v16_)
                d_798_v18_: D6
                d_798_v18_ = D6_DC12((d_796_v16_) | ((d_797_v17_).cf19))
                d_797_v17_ = d_798_v18_
                index152_ = default__.safeIndex(872, (d_792_v12_).length(0))
                (d_792_v12_)[index152_] = ((d_794_v14_).f32 if (self).f25 else (self).f25)
                d_799_v19_: _dafny.Set
                d_799_v19_ = _dafny.Set({d_793_v13_.f35})
                index153_ = default__.safeIndex(408, (d_795_v15_).length(0))
                index154_ = default__.safeIndex(872, (d_792_v12_).length(0))
                rhs164_ = default__.safeModuloInt(default__.safeDivisionInt((self).f31, d_789_v9_), (self).f26)
                rhs165_ = (d_794_v14_).f32
                rhs166_ = 727
                rhs167_ = ((d_799_v19_).intersection(d_799_v19_)).ispropersubset(d_799_v19_)
                lhs149_ = globalState
                lhs150_ = d_795_v15_
                lhs151_ = default__.safeIndex(408, (d_795_v15_).length(0))
                lhs152_ = d_792_v12_
                lhs153_ = default__.safeIndex(872, (d_792_v12_).length(0))
                d_790_v10_ = rhs164_
                lhs149_.f1 = rhs165_
                lhs150_[lhs151_] = rhs166_
                lhs152_[lhs153_] = rhs167_
            elif True:
                d_800_v20_: _dafny.Map
                d_800_v20_ = _dafny.Map({((d_778_v1_)[d_789_v9_] if (d_789_v9_) in (d_778_v1_) else -600): (d_784_v6_ if (self).f25 else d_784_v6_)})
                def iife63_():
                    coll19_ = _dafny.Map()
                    compr_19_: int
                    for compr_19_ in ((d_785_v7_)[default__.safeIndex(161, (d_785_v7_).length(0))]).Elements:
                        d_801_v21_: int = compr_19_
                        if (d_801_v21_) in ((d_785_v7_)[default__.safeIndex(161, (d_785_v7_).length(0))]):
                            coll19_[(d_801_v21_) - ((0) - (p2))] = d_784_v6_
                    return _dafny.Map(coll19_)
                d_800_v20_ = (iife63_()
                ).set(default__.safeDivisionInt(len((d_785_v7_)[default__.safeIndex(161, (d_785_v7_).length(0))]), p1), d_784_v6_)
                (globalState).f1 = (self).f25
                d_785_v7_ = d_785_v7_
                d_802_v22_: _dafny.Array
                def lambda42_(d_803_i3_):
                    return (d_803_i3_) - (948)

                init20_ = lambda42_
                nw140_ = _dafny.Array(None, 28)
                for i0_20_ in range(nw140_.length(0)):
                    nw140_[i0_20_] = init20_(i0_20_)
                d_802_v22_ = nw140_
                d_802_v22_ = d_802_v22_
                r0 = (p1) > (p1)
            r0 = not((self).f25)
        elif True:
            d_804_v23_: _dafny.Array
            def lambda43_(d_805_p2_):
                def lambda44_(d_806_i4_):
                    return default__.safeModuloInt(d_806_i4_, (0) - (d_805_p2_))

                return lambda44_

            init21_ = lambda43_(p2)
            nw141_ = _dafny.Array(None, 13)
            for i0_21_ in range(nw141_.length(0)):
                nw141_[i0_21_] = init21_(i0_21_)
            d_804_v23_ = nw141_
            index155_ = default__.safeIndex(361, (d_804_v23_).length(0))
            (d_804_v23_)[index155_] = ((self).f31) - (p2)
            index156_ = default__.safeIndex(361, (d_804_v23_).length(0))
            (d_804_v23_)[index156_] = ((0) - (p2) if (self).f25 else (self).f26)
            d_807_v24_: _dafny.Seq
            d_807_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwperytjt"))
            d_808_v25_: _dafny.Map
            d_808_v25_ = _dafny.Map({(self).f25: p2})
            source13_ = default__.fm28(d_807_v24_, (d_807_v24_)[default__.safeIndex((self).f31, len(d_807_v24_))], default__.safeModuloInt(len(_dafny.Map({(self).f25: (_dafny.Map({(self).f25: p2})).set((self).f25, ((d_808_v25_)[(self).f25] if ((self).f25) in (d_808_v25_) else (self).f31))})), default__.fm2(p1, (0) - (len((d_808_v25_).set((self).f25, (d_804_v23_)[default__.safeIndex(361, (d_804_v23_).length(0))]))), globalState)), globalState)
            if source13_.is_DC1:
                d_809___mcc_h0_ = source13_.cf1
                d_810___mcc_h1_ = source13_.cf2
                d_811___mcc_h2_ = source13_.cf3
                d_812___mcc_h3_ = source13_.cf4
                d_813___mcc_h4_ = source13_.cf5
                d_814_cf5_ = d_813___mcc_h4_
                d_815_cf4_ = d_812___mcc_h3_
                d_816_cf3_ = d_811___mcc_h2_
                d_817_cf2_ = d_810___mcc_h1_
                d_818_cf1_ = d_809___mcc_h0_
                d_819_v26_: _dafny.Array
                def lambda45_(d_820_cf5_):
                    def lambda46_(d_821_i6_):
                        return d_820_cf5_

                    return lambda46_

                init22_ = lambda45_(d_814_cf5_)
                nw142_ = _dafny.Array(None, 27)
                for i0_22_ in range(nw142_.length(0)):
                    nw142_[i0_22_] = init22_(i0_22_)
                d_819_v26_ = nw142_
                d_822_v27_: T1
                nw143_ = C0()
                nw143_.ctor__(_dafny.SeqWithoutIsStrInference([d_777_v0_ for d_823_i5_ in range(default__.abs(685))]), d_816_cf3_, p2, d_819_v26_, (self).f25)
                d_822_v27_ = nw143_
                d_824_v28_: _dafny.Seq
                d_824_v28_ = _dafny.SeqWithoutIsStrInference([d_822_v27_])
                d_825_v29_: _dafny.Set
                d_825_v29_ = _dafny.Set({d_816_cf3_})
                d_826_v30_: _dafny.Seq
                d_826_v30_ = _dafny.SeqWithoutIsStrInference([d_825_v29_])
                d_827_v31_: _dafny.Map
                d_827_v31_ = _dafny.Map({d_824_v28_: (d_826_v30_) + (_dafny.SeqWithoutIsStrInference([d_825_v29_]))})
                d_827_v31_ = (d_827_v31_).set(d_824_v28_, d_826_v30_)
                d_828_v32_: D2
                d_828_v32_ = D2_DC7(d_816_cf3_, False, d_815_cf4_)
                d_829_v33_: _dafny.Seq
                d_829_v33_ = _dafny.SeqWithoutIsStrInference([d_814_cf5_])
                d_830_v34_: _dafny.Map
                d_830_v34_ = _dafny.Map({(self).f31: d_816_cf3_})
                d_831_v35_: _dafny.MultiSet
                d_831_v35_ = _dafny.MultiSet([d_818_cf1_, (self).f31])
                d_832_v36_: _dafny.Map
                d_832_v36_ = _dafny.Map({len(d_829_v33_): not(((d_830_v34_)[((d_831_v35_)[(d_831_v35_).cardinality] if ((d_831_v35_).cardinality) in (d_831_v35_) else (self).f26)] if (((d_831_v35_)[(d_831_v35_).cardinality] if ((d_831_v35_).cardinality) in (d_831_v35_) else (self).f26)) in (d_830_v34_) else d_816_cf3_))})
                d_833_v37_: _dafny.MultiSet
                d_833_v37_ = _dafny.MultiSet([d_832_v36_])
                d_834_v38_: _dafny.Map
                d_834_v38_ = _dafny.Map({d_828_v32_: ((_dafny.MultiSet([_dafny.Map({-923: d_814_cf5_})])).set(d_830_v34_, default__.abs(d_818_cf1_))) - (d_833_v37_)})
                d_835_v39_: _dafny.Seq
                d_835_v39_ = _dafny.SeqWithoutIsStrInference([d_833_v37_])
                d_834_v38_ = (d_834_v38_).set(d_828_v32_, (d_835_v39_)[default__.safeIndex((self).f26, len(d_835_v39_))])
                d_836_v40_: _dafny.Set
                d_836_v40_ = _dafny.Set({d_777_v0_})
                d_837_v41_: int
                out20_: int
                out20_ = default__.m0(not((d_777_v0_) not in (d_836_v40_)), globalState)
                d_837_v41_ = out20_
                index157_ = default__.safeIndex(361, (d_804_v23_).length(0))
                (d_804_v23_)[index157_] = ((self).f26) * ((0) - (d_815_cf4_))
            elif source13_.is_DC2:
                d_838_v42_: _dafny.Map
                d_838_v42_ = _dafny.Map({(d_804_v23_)[default__.safeIndex(361, (d_804_v23_).length(0))]: (p1) == ((d_804_v23_)[default__.safeIndex(361, (d_804_v23_).length(0))])})
                d_838_v42_ = (d_838_v42_).set((0) - (default__.safeDivisionInt((self).f26, (self).f31)), not((self).f25))
                d_839_v43_: _dafny.Map
                d_839_v43_ = _dafny.Map({(self).f25: d_807_v24_})
                d_840_v44_: _dafny.Array
                def lambda47_(d_841_i7_):
                    return (self).f25

                init23_ = lambda47_
                nw144_ = _dafny.Array(None, 26)
                for i0_23_ in range(nw144_.length(0)):
                    nw144_[i0_23_] = init23_(i0_23_)
                d_840_v44_ = nw144_
                d_842_v45_: _dafny.Map
                d_842_v45_ = _dafny.Map({(self).f25: d_840_v44_})
                d_843_v46_: T0
                nw145_ = C0()
                nw145_.ctor__(d_807_v24_, not(False), len(d_839_v43_), ((d_842_v45_)[(self).f25] if ((self).f25) in (d_842_v45_) else d_840_v44_), (self).f25)
                d_843_v46_ = nw145_
                d_844_v47_: _dafny.Seq
                d_844_v47_ = _dafny.SeqWithoutIsStrInference([d_843_v46_, d_843_v46_])
                d_845_v48_: _dafny.Map
                d_845_v48_ = _dafny.Map({(self).f25: d_844_v47_})
                d_846_v49_: D9
                d_846_v49_ = D9_DC18(d_845_v48_)
                pat_let_tv24_ = d_845_v48_
                d_847_v50_: _dafny.Array
                nw146_ = _dafny.Array(None, 29)
                nw146_[int(0)] = d_846_v49_
                nw146_[int(1)] = d_846_v49_
                nw146_[int(2)] = d_846_v49_
                nw146_[int(3)] = d_846_v49_
                nw146_[int(4)] = D9_DC18(d_845_v48_)
                nw146_[int(5)] = D9_DC18(d_845_v48_)
                def iife64_(_pat_let22_0):
                    def iife65_(d_848_dt__update__tmp_h0_):
                        def iife66_(_pat_let23_0):
                            def iife67_(d_849_dt__update_hcf30_h0_):
                                return D9_DC18(d_849_dt__update_hcf30_h0_)
                            return iife67_(_pat_let23_0)
                        return iife66_(pat_let_tv24_)
                    return iife65_(_pat_let22_0)
                nw146_[int(6)] = iife64_(d_846_v49_)
                nw146_[int(7)] = d_846_v49_
                nw146_[int(8)] = D9_DC18(d_845_v48_)
                nw146_[int(9)] = d_846_v49_
                nw146_[int(10)] = d_846_v49_
                nw146_[int(11)] = d_846_v49_
                nw146_[int(12)] = d_846_v49_
                nw146_[int(13)] = d_846_v49_
                nw146_[int(14)] = d_846_v49_
                nw146_[int(15)] = d_846_v49_
                nw146_[int(16)] = d_846_v49_
                nw146_[int(17)] = d_846_v49_
                nw146_[int(18)] = d_846_v49_
                nw146_[int(19)] = d_846_v49_
                nw146_[int(20)] = d_846_v49_
                nw146_[int(21)] = d_846_v49_
                nw146_[int(22)] = d_846_v49_
                nw146_[int(23)] = d_846_v49_
                nw146_[int(24)] = d_846_v49_
                nw146_[int(25)] = D9_DC18(d_845_v48_)
                nw146_[int(26)] = D9_DC18(d_845_v48_)
                nw146_[int(27)] = d_846_v49_
                nw146_[int(28)] = d_846_v49_
                d_847_v50_ = nw146_
                d_850_v51_: _dafny.Array
                nw147_ = _dafny.Array(None, 25)
                nw147_[int(0)] = d_847_v50_
                nw147_[int(1)] = d_847_v50_
                nw147_[int(2)] = d_847_v50_
                nw147_[int(3)] = d_847_v50_
                nw147_[int(4)] = d_847_v50_
                nw147_[int(5)] = d_847_v50_
                nw147_[int(6)] = d_847_v50_
                nw147_[int(7)] = d_847_v50_
                nw147_[int(8)] = d_847_v50_
                nw147_[int(9)] = d_847_v50_
                nw147_[int(10)] = d_847_v50_
                nw147_[int(11)] = d_847_v50_
                nw147_[int(12)] = d_847_v50_
                nw147_[int(13)] = d_847_v50_
                nw147_[int(14)] = d_847_v50_
                nw147_[int(15)] = d_847_v50_
                nw147_[int(16)] = d_847_v50_
                nw147_[int(17)] = d_847_v50_
                nw147_[int(18)] = d_847_v50_
                nw147_[int(19)] = d_847_v50_
                nw147_[int(20)] = d_847_v50_
                nw147_[int(21)] = d_847_v50_
                nw147_[int(22)] = d_847_v50_
                nw147_[int(23)] = d_847_v50_
                nw147_[int(24)] = d_847_v50_
                d_850_v51_ = nw147_
                index158_ = default__.safeIndex(59, (d_850_v51_).length(0))
                (d_850_v51_)[index158_] = d_847_v50_
                index159_ = default__.safeIndex(59, (d_850_v51_).length(0))
                nw148_ = _dafny.Array(D9.default()(), 3)
                (d_850_v51_)[index159_] = nw148_
                d_851_v52_: C1
                nw149_ = C1()
                nw149_.ctor__((self).f25)
                d_851_v52_ = nw149_
                d_851_v52_ = d_851_v52_
                d_852_v53_: _dafny.Map
                d_852_v53_ = _dafny.Map({956: (d_807_v24_)[default__.safeIndex((self).f31, len(d_807_v24_))]})
                d_853_v54_: D4
                d_853_v54_ = D4_DC10(d_777_v0_, d_777_v0_)
                d_852_v53_ = (d_852_v53_).set(p1, (d_853_v54_).cf16)
            elif True:
                d_854___mcc_h5_ = source13_.cf0
                d_855_cf0_ = d_854___mcc_h5_
                d_856_v55_: int
                out21_: int
                out21_ = default__.m0(not(True), globalState)
                d_856_v55_ = out21_
                (globalState).f0 = default__.safeModuloInt((self).f26, (self).f26)
                d_857_v56_: _dafny.Seq
                d_857_v56_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f25, False, (self).f25])])
                d_858_v57_: _dafny.Array
                nw150_ = _dafny.Array(None, 23)
                nw150_[int(0)] = d_855_cf0_
                nw150_[int(1)] = (self).f25
                nw150_[int(2)] = (self).f25
                nw150_[int(3)] = (self).f25
                nw150_[int(4)] = (self).f25
                nw150_[int(5)] = not(d_855_cf0_)
                nw150_[int(6)] = (self).f25
                nw150_[int(7)] = d_855_cf0_
                nw150_[int(8)] = (self).f25
                nw150_[int(9)] = (p1) != ((d_804_v23_)[default__.safeIndex(361, (d_804_v23_).length(0))])
                nw150_[int(10)] = d_855_cf0_
                nw150_[int(11)] = d_855_cf0_
                nw150_[int(12)] = (self).f25
                nw150_[int(13)] = False
                nw150_[int(14)] = not(d_855_cf0_)
                nw150_[int(15)] = d_855_cf0_
                nw150_[int(16)] = (self).f25
                nw150_[int(17)] = d_855_cf0_
                nw150_[int(18)] = d_855_cf0_
                nw150_[int(19)] = True
                nw150_[int(20)] = (len(_dafny.Set({len(d_857_v56_)}))) in (p0)
                nw150_[int(21)] = d_855_cf0_
                nw150_[int(22)] = d_855_cf0_
                d_858_v57_ = nw150_
                index160_ = default__.safeIndex(671, (d_858_v57_).length(0))
                (d_858_v57_)[index160_] = (self).f25
                index161_ = default__.safeIndex(671, (d_858_v57_).length(0))
                (d_858_v57_)[index161_] = (d_807_v24_) == (d_807_v24_)
                d_859_v58_: _dafny.Map
                d_859_v58_ = _dafny.Map({True: (self).f25})
                d_860_v59_: C1
                nw151_ = C1()
                nw151_.ctor__(d_855_cf0_)
                d_860_v59_ = nw151_
                d_861_v60_: _dafny.Seq
                d_861_v60_ = _dafny.SeqWithoutIsStrInference([d_860_v59_, d_860_v59_, d_860_v59_])
                d_862_v61_: _dafny.Seq
                d_862_v61_ = _dafny.SeqWithoutIsStrInference([d_861_v60_])
                d_859_v58_ = ((d_859_v58_).set((self).f25, True)).set(((self).f26) == (len(d_862_v61_)), (d_807_v24_) != (d_807_v24_))
            (globalState).f1 = True
            (globalState).f1 = ((self).f25 if (self).f25 else False)
            index162_ = default__.safeIndex(361, (d_804_v23_).length(0))
            (d_804_v23_)[index162_] = (self).f26
        d_863_i8_: int
        d_863_i8_ = 0
        with _dafny.label("3"):
            while not((self).f25):
                with _dafny.c_label("3"):
                    if (d_863_i8_) >= (100):
                        raise _dafny.Break("3")
                    d_863_i8_ = (d_863_i8_) + (1)
                    d_864_v62_: D0
                    d_864_v62_ = D0_DC2()
                    d_865_v63_: _dafny.Map
                    d_865_v63_ = _dafny.Map({(self).f25: d_864_v62_})
                    d_866_v64_: _dafny.Seq
                    d_866_v64_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(len(d_865_v63_), (self).f31)])
                    (globalState).f0 = (d_866_v64_)[default__.safeIndex(761, len(d_866_v64_))]
                    d_867_v65_: _dafny.Set
                    d_867_v65_ = _dafny.Set({default__.fm2(-334, (self).f31, globalState), p1})
                    if not(not((d_867_v65_).isdisjoint(d_867_v65_))):
                        d_868_v66_: _dafny.Seq
                        d_868_v66_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25, (self).f25, (self).f25])
                        d_869_v67_: _dafny.Seq
                        d_869_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynxpxku"))
                        d_870_v68_: _dafny.Map
                        d_870_v68_ = _dafny.Map({p2: d_869_v67_})
                        d_871_v69_: _dafny.Map
                        d_871_v69_ = _dafny.Map({len(((d_870_v68_)[-3] if (-3) in (d_870_v68_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_872_i9_ in range(default__.abs(0))]))): (self).f25})
                        d_873_v70_: _dafny.Map
                        d_873_v70_ = _dafny.Map({(self).f25: p2})
                        d_874_v71_: _dafny.Seq
                        d_874_v71_ = _dafny.SeqWithoutIsStrInference([d_873_v70_, d_873_v70_])
                        d_875_v72_: _dafny.Seq
                        d_875_v72_ = (((d_868_v66_).set(default__.safeIndex((0) - (len(d_871_v69_)), len(d_868_v66_)), (self).f25)).set(default__.safeIndex(807, len((d_868_v66_).set(default__.safeIndex((0) - (len(d_871_v69_)), len(d_868_v66_)), (self).f25))), (self).f25)) + ((d_868_v66_).set(default__.safeIndex(len(d_874_v71_), len(d_868_v66_)), (self).f25))
                        d_875_v72_ = d_875_v72_
                        d_876_v73_: _dafny.MultiSet
                        d_876_v73_ = _dafny.MultiSet([d_777_v0_, d_777_v0_, (d_777_v0_ if (self).f25 else d_777_v0_)])
                        d_877_v74_: D0
                        d_877_v74_ = D0_DC1((self).f31, d_869_v67_, (self).f25, (self).f31, (self).f25)
                        d_878_v76_: _dafny.Set
                        d_878_v76_ = _dafny.Set({(self).f25, True})
                        d_879_v77_: _dafny.Array
                        nw152_ = _dafny.Array(None, 13)
                        nw152_[int(0)] = default__.safeDivisionInt(p2, p1)
                        nw152_[int(1)] = ((self).f31 if (self).f25 else p1)
                        nw152_[int(2)] = p1
                        nw152_[int(3)] = (self).f26
                        nw152_[int(4)] = (d_877_v74_).cf4
                        nw152_[int(5)] = 22
                        nw152_[int(6)] = ((self).f26) * ((self).f31)
                        nw152_[int(7)] = p2
                        def iife68_():
                            coll20_ = _dafny.Map()
                            compr_20_: int
                            for compr_20_ in (p0).keys.Elements:
                                d_880_v75_: int = compr_20_
                                if (d_880_v75_) in (p0):
                                    coll20_[(d_880_v75_) * ((self).f26)] = not((self).f25)
                            return _dafny.Map(coll20_)
                        nw152_[int(8)] = len(iife68_()
                        )
                        nw152_[int(9)] = (len(d_866_v64_)) * (p2)
                        nw152_[int(10)] = (self).f31
                        nw152_[int(11)] = (len(d_878_v76_)) - ((self).f31)
                        nw152_[int(12)] = (self).f31
                        d_879_v77_ = nw152_
                        index163_ = default__.safeIndex(605, (d_879_v77_).length(0))
                        (d_879_v77_)[index163_] = p1
                        d_881_v78_: D2
                        d_881_v78_ = D2_DC7((self).f25, (self).f25, p1)
                        index164_ = default__.safeIndex(605, (d_879_v77_).length(0))
                        rhs168_ = d_876_v73_
                        rhs169_ = ((d_881_v78_).cf13) - (p1)
                        rhs170_ = (self).f26
                        rhs171_ = not (not(((self).f25 if not((self).f25) else (self).f25))) or ((self).f25)
                        rhs172_ = (self).fm7(globalState)
                        lhs154_ = globalState
                        lhs155_ = d_879_v77_
                        lhs156_ = default__.safeIndex(605, (d_879_v77_).length(0))
                        lhs157_ = globalState
                        d_876_v73_ = rhs168_
                        lhs154_.f2 = rhs169_
                        lhs155_[lhs156_] = rhs170_
                        lhs157_.f1 = rhs171_
                        r0 = rhs172_
                        d_882_v79_: _dafny.Seq
                        d_882_v79_ = _dafny.SeqWithoutIsStrInference([d_869_v67_, d_869_v67_])
                        d_869_v67_ = (d_869_v67_) + ((d_869_v67_) + (((d_882_v79_)[default__.safeIndex((0) - (-414), len(d_882_v79_))]).set(default__.safeIndex(len((d_869_v67_).set(default__.safeIndex(len(d_869_v67_), len(d_869_v67_)), d_777_v0_)), len((d_882_v79_)[default__.safeIndex((0) - (-414), len(d_882_v79_))])), d_777_v0_)))
                        d_883_v80_: _dafny.MultiSet
                        d_883_v80_ = _dafny.MultiSet([(self).f25])
                        rhs173_ = ((d_883_v80_).set((self).f25, default__.abs(p2))) - (d_883_v80_)
                        rhs174_ = not(not((self).f25))
                        rhs175_ = (d_878_v76_).intersection(d_878_v76_)
                        lhs158_ = globalState
                        d_883_v80_ = rhs173_
                        lhs158_.f1 = rhs174_
                        d_878_v76_ = rhs175_
                        d_884_v81_: D6
                        d_884_v81_ = D6_DC13(990, p1, _dafny.Set({(self).f26, (d_879_v77_)[default__.safeIndex(605, (d_879_v77_).length(0))], len(d_869_v67_)}))
                        d_885_v82_: _dafny.Map
                        d_885_v82_ = _dafny.Map({d_777_v0_: (0) - ((d_884_v81_).cf21)})
                        d_885_v82_ = d_885_v82_
                    elif True:
                        d_886_v83_: _dafny.Seq
                        d_886_v83_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kg"))
                        d_887_v84_: D0
                        d_887_v84_ = D0_DC0((self).f25)
                        d_888_v85_: _dafny.Map
                        d_888_v85_ = _dafny.Map({d_887_v84_: (0) - ((self).f31)})
                        d_889_v86_: _dafny.Array
                        nw153_ = _dafny.Array(False, 20)
                        d_889_v86_ = nw153_
                        d_890_v87_: T2
                        nw154_ = C0()
                        nw154_.ctor__(d_886_v83_, default__.fm5(True, d_888_v85_, 814, globalState), (self).f26, d_889_v86_, False)
                        d_890_v87_ = nw154_
                        d_891_v88_: _dafny.Seq
                        d_891_v88_ = _dafny.SeqWithoutIsStrInference([False])
                        index165_ = default__.safeIndex(918, (d_889_v86_).length(0))
                        (d_889_v86_)[index165_] = (d_891_v88_)[default__.safeIndex(p2, len(d_891_v88_))]
                        d_892_v89_: _dafny.Array
                        def lambda48_(d_893_v64_):
                            def lambda49_(d_894_i10_):
                                return _dafny.Map({d_893_v64_: True})

                            return lambda49_

                        init24_ = lambda48_(d_866_v64_)
                        nw155_ = _dafny.Array(None, 10)
                        for i0_24_ in range(nw155_.length(0)):
                            nw155_[i0_24_] = init24_(i0_24_)
                        d_892_v89_ = nw155_
                        d_895_v91_: _dafny.Seq
                        d_895_v91_ = _dafny.SeqWithoutIsStrInference([d_866_v64_])
                        index166_ = default__.safeIndex(973, (d_892_v89_).length(0))
                        def iife69_():
                            coll21_ = _dafny.Map()
                            compr_21_: _dafny.Seq
                            for compr_21_ in (d_895_v91_).Elements:
                                d_896_v90_: _dafny.Seq = compr_21_
                                if (d_896_v90_) in (d_895_v91_):
                                    coll21_[d_896_v90_] = (self).f25
                            return _dafny.Map(coll21_)
                        (d_892_v89_)[index166_] = iife69_()
                        
                        d_897_v92_: _dafny.Set
                        d_897_v92_ = _dafny.Set({d_890_v87_.f34})
                        d_898_v93_: _dafny.Map
                        d_898_v93_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([p2, len(d_886_v83_)])})
                        index167_ = default__.safeIndex(918, (d_889_v86_).length(0))
                        index168_ = default__.safeIndex(973, (d_892_v89_).length(0))
                        rhs176_ = (self).f31
                        rhs177_ = d_890_v87_
                        rhs178_ = (self).f25
                        rhs179_ = ((p2 if (self).f25 else (D2_DC7(True, d_890_v87_.f34, len(d_897_v92_))).cf13)) * (p2)
                        rhs180_ = (_dafny.Map({((d_898_v93_)[-64] if (-64) in (d_898_v93_) else d_866_v64_): True})) | (_dafny.Map({d_866_v64_: (self).f25}))
                        lhs159_ = globalState
                        lhs160_ = d_889_v86_
                        lhs161_ = default__.safeIndex(918, (d_889_v86_).length(0))
                        lhs162_ = globalState
                        lhs163_ = d_892_v89_
                        lhs164_ = default__.safeIndex(973, (d_892_v89_).length(0))
                        lhs159_.f0 = rhs176_
                        d_890_v87_ = rhs177_
                        lhs160_[lhs161_] = rhs178_
                        lhs162_.f2 = rhs179_
                        lhs163_[lhs164_] = rhs180_
                        d_899_v94_: _dafny.Map
                        d_899_v94_ = _dafny.Map({d_886_v83_: d_886_v83_})
                        d_899_v94_ = d_899_v94_
                        d_900_v95_: _dafny.Array
                        nw156_ = _dafny.Array(int(0), 7)
                        d_900_v95_ = nw156_
                        index169_ = default__.safeIndex(508, (d_900_v95_).length(0))
                        (d_900_v95_)[index169_] = p2
                        index170_ = default__.safeIndex(508, (d_900_v95_).length(0))
                        (d_900_v95_)[index170_] = (self).f31
                        d_901_v96_: D1
                        d_901_v96_ = D1_DC4((d_889_v86_)[default__.safeIndex(918, (d_889_v86_).length(0))], True)
                        (globalState).f0 = default__.fm2(((d_900_v95_)[default__.safeIndex(508, (d_900_v95_).length(0))] if (d_901_v96_).cf7 else p1), (d_900_v95_)[default__.safeIndex(508, (d_900_v95_).length(0))], globalState)
                        d_902_v97_: _dafny.Set
                        d_902_v97_ = _dafny.Set({d_867_v65_})
                        index171_ = default__.safeIndex(918, (d_889_v86_).length(0))
                        (d_889_v86_)[index171_] = (d_902_v97_).isdisjoint(d_902_v97_)
                    d_903_v98_: _dafny.Seq
                    d_903_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbigaspv"))
                    d_904_v99_: _dafny.Set
                    d_904_v99_ = _dafny.Set({(self).f25})
                    d_905_v100_: _dafny.Map
                    d_905_v100_ = _dafny.Map({len(d_903_v98_): d_904_v99_})
                    d_906_v101_: _dafny.Seq
                    d_906_v101_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                    d_905_v100_ = (d_905_v100_).set((_dafny.MultiSet([d_903_v98_])).cardinality, (_dafny.Set({(d_906_v101_)[default__.safeIndex(p1, len(d_906_v101_))], (self).f25})) - (default__.fm29((self).fm7(globalState), p0, d_777_v0_, (self).f26, globalState)))
                    d_904_v99_ = d_904_v99_
                    pass
            pass
        d_907_v102_: int
        out22_: int
        out22_ = default__.m0((self).f25, globalState)
        d_907_v102_ = out22_
        d_908_v103_: _dafny.MultiSet
        d_908_v103_ = _dafny.MultiSet([default__.fm2((_dafny.MultiSet([p2])).cardinality, (self).f31, globalState)])
        d_909_v104_: _dafny.Seq
        d_909_v104_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        d_910_v105_: _dafny.MultiSet
        d_910_v105_ = _dafny.MultiSet([(d_909_v104_)[default__.safeIndex((self).f26, len(d_909_v104_))]])
        d_911_v106_: _dafny.Array
        nw157_ = _dafny.Array(None, 17)
        nw157_[int(0)] = d_908_v103_
        nw157_[int(1)] = d_908_v103_
        nw157_[int(2)] = d_910_v105_
        nw157_[int(3)] = d_908_v103_
        nw157_[int(4)] = d_908_v103_
        nw157_[int(5)] = d_908_v103_
        nw157_[int(6)] = default__.fm30(globalState)
        nw157_[int(7)] = d_908_v103_
        nw157_[int(8)] = d_908_v103_
        nw157_[int(9)] = default__.fm30(globalState)
        nw157_[int(10)] = d_908_v103_
        nw157_[int(11)] = d_910_v105_
        nw157_[int(12)] = d_908_v103_
        nw157_[int(13)] = (d_910_v105_).set(p1, default__.abs((self).f26))
        nw157_[int(14)] = d_908_v103_
        nw157_[int(15)] = d_908_v103_
        nw157_[int(16)] = d_908_v103_
        d_911_v106_ = nw157_
        d_912_v107_: _dafny.Seq
        d_912_v107_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnxuic"))
        d_913_v108_: D0
        d_913_v108_ = D0_DC1(d_907_v102_, d_912_v107_, True, p1, (self).f25)
        d_914_v109_: _dafny.MultiSet
        d_914_v109_ = _dafny.MultiSet([True, (self).f25, False])
        rhs181_ = d_911_v106_
        rhs182_ = (d_913_v108_).cf4
        rhs183_ = (d_914_v109_).issubset(d_914_v109_)
        rhs184_ = (d_909_v104_) + ((d_909_v104_) + (d_909_v104_))
        lhs165_ = globalState
        lhs166_ = globalState
        d_911_v106_ = rhs181_
        lhs165_.f2 = rhs182_
        lhs166_.f1 = rhs183_
        d_909_v104_ = rhs184_
        d_915_v110_: _dafny.Array
        nw158_ = _dafny.Array(False, 15)
        d_915_v110_ = nw158_
        index172_ = default__.safeIndex(869, (d_915_v110_).length(0))
        (d_915_v110_)[index172_] = (self).f25
        d_916_v111_: D0
        d_916_v111_ = D0_DC0((self).f25)
        index173_ = default__.safeIndex(869, (d_915_v110_).length(0))
        def lambda50_(source14_):
            if source14_.is_DC1:
                d_917___mcc_h6_ = source14_.cf1
                d_918___mcc_h7_ = source14_.cf2
                d_919___mcc_h8_ = source14_.cf3
                d_920___mcc_h9_ = source14_.cf4
                d_921___mcc_h10_ = source14_.cf5
                d_922_cf5_ = d_921___mcc_h10_
                d_923_cf4_ = d_920___mcc_h9_
                d_924_cf3_ = d_919___mcc_h8_
                d_925_cf2_ = d_918___mcc_h7_
                d_926_cf1_ = d_917___mcc_h6_
                return (self).f25
            elif source14_.is_DC2:
                return (self).f25
            elif True:
                d_927___mcc_h11_ = source14_.cf0
                d_928_cf0_ = d_927___mcc_h11_
                return d_928_cf0_

        (d_915_v110_)[index173_] = lambda50_(d_916_v111_)
        r0 = (d_915_v110_)[default__.safeIndex(869, (d_915_v110_).length(0))]
        return r0

    @property
    def f31(self):
        return self._f31

class C3(T0):
    def  __init__(self):
        self._f25: bool = False
        self._f26: int = int(0)
        self._f29: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f30: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f29, f30, f25, f26):
        (self)._f29 = f29
        (self)._f30 = f30
        (self)._f25 = f25
        (self)._f26 = f26

    def fm6(self, p0, p1, p2, globalState):
        return ((self).f26) * (((self).f26) - ((self).f26))

    def fm7(self, globalState):
        return (926) != ((self).f26)

    def fm9(self, globalState):
        return (0) - ((self).f26)

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        d_929_v0_: str
        d_929_v0_ = _dafny.CodePoint('u')
        d_929_v0_ = (default__.fm10((self).f25, p2, globalState)).cf6
        d_930_v1_: _dafny.Map
        d_930_v1_ = _dafny.Map({(self).f30: p1})
        if ((self).f25 if (True if ((d_930_v1_)[True] if (True) in (d_930_v1_) else (self).f25) else (self).f30) else p2):
            d_931_v2_: _dafny.Map
            d_931_v2_ = _dafny.Map({(self).f26: not ((self).f30) or ((self).f30)})
            d_931_v2_ = (d_931_v2_).set((self).f26, (self).f25)
            d_932_v3_: _dafny.Array
            nw159_ = _dafny.Array(False, 6)
            d_932_v3_ = nw159_
            d_933_v4_: _dafny.Seq
            d_933_v4_ = _dafny.SeqWithoutIsStrInference([d_932_v3_])
            d_934_v5_: _dafny.Array
            nw160_ = _dafny.Array(None, 26)
            nw160_[int(0)] = d_932_v3_
            nw160_[int(1)] = d_932_v3_
            nw160_[int(2)] = d_932_v3_
            nw160_[int(3)] = d_932_v3_
            nw160_[int(4)] = d_932_v3_
            nw160_[int(5)] = d_932_v3_
            nw160_[int(6)] = d_932_v3_
            nw160_[int(7)] = d_932_v3_
            nw160_[int(8)] = d_932_v3_
            nw160_[int(9)] = d_932_v3_
            nw160_[int(10)] = d_932_v3_
            nw160_[int(11)] = d_932_v3_
            nw160_[int(12)] = d_932_v3_
            nw160_[int(13)] = d_932_v3_
            nw160_[int(14)] = d_932_v3_
            nw160_[int(15)] = d_932_v3_
            nw160_[int(16)] = d_932_v3_
            nw160_[int(17)] = d_932_v3_
            nw160_[int(18)] = d_932_v3_
            nw160_[int(19)] = d_932_v3_
            nw160_[int(20)] = d_932_v3_
            nw160_[int(21)] = d_932_v3_
            nw160_[int(22)] = d_932_v3_
            nw160_[int(23)] = (d_932_v3_ if (self).f30 else d_932_v3_)
            nw160_[int(24)] = (d_933_v4_)[default__.safeIndex((self).f26, len(d_933_v4_))]
            nw160_[int(25)] = (d_933_v4_)[default__.safeIndex((self).f26, len(d_933_v4_))]
            d_934_v5_ = nw160_
            d_935_v6_: _dafny.Set
            d_935_v6_ = _dafny.Set({(self).f30, (self).f25})
            d_936_v7_: D1
            d_936_v7_ = D1_DC5(D1_DC3(d_929_v0_))
            d_937_v8_: D1
            d_937_v8_ = D1_DC5(d_936_v7_)
            pat_let_tv25_ = d_936_v7_
            d_938_v9_: _dafny.MultiSet
            def iife70_(_pat_let24_0):
                def iife71_(d_939_dt__update__tmp_h0_):
                    def iife72_(_pat_let25_0):
                        def iife73_(d_940_dt__update_hcf9_h0_):
                            return D1_DC5(d_940_dt__update_hcf9_h0_)
                        return iife73_(_pat_let25_0)
                    return iife72_(pat_let_tv25_)
                return iife71_(_pat_let24_0)
            d_938_v9_ = _dafny.MultiSet([D1_DC5(d_936_v7_), iife70_(d_937_v8_), d_937_v8_])
            rhs185_ = not ((d_935_v6_).isdisjoint(d_935_v6_)) or (p2)
            rhs186_ = (self).f26
            rhs187_ = d_934_v5_
            rhs188_ = ((d_931_v2_)[((d_938_v9_)[d_937_v8_] if (d_937_v8_) in (d_938_v9_) else (self).f26)] if (((d_938_v9_)[d_937_v8_] if (d_937_v8_) in (d_938_v9_) else (self).f26)) in (d_931_v2_) else ((self).f25) or (p1))
            lhs167_ = globalState
            lhs168_ = globalState
            r0 = rhs185_
            lhs167_.f2 = rhs186_
            d_934_v5_ = rhs187_
            lhs168_.f1 = rhs188_
            d_941_v10_: _dafny.Array
            def lambda51_(d_942_i0_):
                return _dafny.SeqWithoutIsStrInference([True, True])

            init25_ = lambda51_
            nw161_ = _dafny.Array(None, 8)
            for i0_25_ in range(nw161_.length(0)):
                nw161_[i0_25_] = init25_(i0_25_)
            d_941_v10_ = nw161_
            d_943_v11_: _dafny.Seq
            d_943_v11_ = _dafny.SeqWithoutIsStrInference([p1])
            d_944_v12_: D0
            d_944_v12_ = D0_DC0((self).f30)
            d_945_v13_: _dafny.Map
            d_945_v13_ = _dafny.Map({d_944_v12_: (self).f26})
            index174_ = default__.safeIndex(41, (d_941_v10_).length(0))
            (d_941_v10_)[index174_] = (d_943_v11_).set(default__.safeIndex((self).f26, len(d_943_v11_)), default__.fm5((self).f30, d_945_v13_, (self).f26, globalState))
            index175_ = default__.safeIndex(41, (d_941_v10_).length(0))
            (d_941_v10_)[index175_] = _dafny.SeqWithoutIsStrInference([p1])
            d_946_v14_: D0
            d_946_v14_ = D0_DC1(len(_dafny.SeqWithoutIsStrInference([d_929_v0_ for d_947_i1_ in range(default__.abs(689))])), (self).f29, (self).f30, (self).f26, p2)
            d_929_v0_ = default__.fm4(d_946_v14_, (self).f25, (0) - ((self).f26), globalState)
            (globalState).f0 = (self).f26
        elif True:
            d_948_v15_: _dafny.Seq
            d_948_v15_ = _dafny.SeqWithoutIsStrInference([p2])
            d_949_v16_: T0
            nw162_ = C2()
            nw162_.ctor__((0) - ((self).f26), not(not((self).f25)), ((self).f26 if p2 else len(_dafny.SeqWithoutIsStrInference([len(d_948_v15_), (p0).cardinality]))))
            d_949_v16_ = nw162_
            d_949_v16_ = (d_949_v16_ if p1 else d_949_v16_)
            d_948_v15_ = d_948_v15_
            (globalState).f1 = (D0_DC1((0) - ((d_949_v16_).f26), (self).f29, True, (d_949_v16_).f26, not(((d_930_v1_)[True] if (True) in (d_930_v1_) else p1)))).cf5
            rhs189_ = (self).f30
            rhs190_ = (0) - ((self).f26)
            lhs169_ = globalState
            lhs170_ = globalState
            lhs169_.f1 = rhs189_
            lhs170_.f2 = rhs190_
            d_950_v17_: _dafny.Set
            d_950_v17_ = _dafny.Set({default__.fm2((self).f26, (self).f26, globalState), (self).f26})
            d_951_v18_: D6
            d_951_v18_ = D6_DC13((0) - ((self).f26), (d_949_v16_).f26, d_950_v17_)
            d_952_v19_: _dafny.Map
            d_952_v19_ = _dafny.Map({(d_951_v18_).cf21: (self).f26})
            d_952_v19_ = d_952_v19_
        d_953_i2_: int
        d_953_i2_ = 0
        with _dafny.label("4"):
            while not(((self).f26) == (((self).f26) - ((self).f26))):
                with _dafny.c_label("4"):
                    if (d_953_i2_) >= (100):
                        raise _dafny.Break("4")
                    d_953_i2_ = (d_953_i2_) + (1)
                    d_954_v20_: _dafny.Array
                    nw163_ = _dafny.Array(False, 28)
                    d_954_v20_ = nw163_
                    index176_ = default__.safeIndex(723, (d_954_v20_).length(0))
                    (d_954_v20_)[index176_] = (self).f25
                    d_955_v21_: _dafny.Seq
                    d_955_v21_ = _dafny.SeqWithoutIsStrInference([False])
                    d_956_v22_: _dafny.Seq
                    d_956_v22_ = _dafny.SeqWithoutIsStrInference([d_955_v21_, d_955_v21_])
                    d_957_v23_: _dafny.MultiSet
                    d_957_v23_ = _dafny.MultiSet([(d_956_v22_)[default__.safeIndex((self).f26, len(d_956_v22_))], d_955_v21_, d_955_v21_])
                    index177_ = default__.safeIndex(723, (d_954_v20_).length(0))
                    (d_954_v20_)[index177_] = not((d_955_v21_) not in (d_957_v23_))
                    if not(p1):
                        d_958_v24_: _dafny.Map
                        d_958_v24_ = _dafny.Map({(self).f29: ((self).f26) + ((self).f26)})
                        d_959_v25_: _dafny.Array
                        nw164_ = _dafny.Array(_dafny.Seq({}), 4)
                        d_959_v25_ = nw164_
                        index178_ = default__.safeIndex(587, (d_959_v25_).length(0))
                        (d_959_v25_)[index178_] = default__.fm3(973, (d_954_v20_)[default__.safeIndex(723, (d_954_v20_).length(0))], (self).f26, (self).f30, globalState)
                        d_960_v26_: _dafny.Seq
                        d_960_v26_ = _dafny.SeqWithoutIsStrInference([(self).f26, len(_dafny.SeqWithoutIsStrInference([(self).f26]))])
                        index179_ = default__.safeIndex(587, (d_959_v25_).length(0))
                        rhs191_ = d_958_v24_
                        rhs192_ = d_960_v26_
                        lhs171_ = d_959_v25_
                        lhs172_ = default__.safeIndex(587, (d_959_v25_).length(0))
                        d_958_v24_ = rhs191_
                        lhs171_[lhs172_] = rhs192_
                        d_961_v27_: _dafny.Array
                        nw165_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                        d_961_v27_ = nw165_
                        index180_ = default__.safeIndex(840, (d_961_v27_).length(0))
                        (d_961_v27_)[index180_] = d_929_v0_
                        d_962_v28_: _dafny.Set
                        d_962_v28_ = _dafny.Set({((self).f26) >= ((self).f26)})
                        pat_let_tv26_ = d_929_v0_
                        index181_ = default__.safeIndex(840, (d_961_v27_).length(0))
                        def iife74_(_pat_let26_0):
                            def iife75_(d_963_dt__update__tmp_h1_):
                                def iife76_(_pat_let27_0):
                                    def iife77_(d_964_dt__update_hcf17_h0_):
                                        return D4_DC10((d_963_dt__update__tmp_h1_).cf16, d_964_dt__update_hcf17_h0_)
                                    return iife77_(_pat_let27_0)
                                return iife76_(pat_let_tv26_)
                            return iife75_(_pat_let26_0)
                        rhs193_ = d_954_v20_
                        rhs194_ = (d_955_v21_) + (d_955_v21_)
                        rhs195_ = (iife74_(D4_DC10(d_929_v0_, d_929_v0_))).cf17
                        rhs196_ = (self).f26
                        rhs197_ = d_962_v28_
                        lhs173_ = d_961_v27_
                        lhs174_ = default__.safeIndex(840, (d_961_v27_).length(0))
                        lhs175_ = globalState
                        d_954_v20_ = rhs193_
                        d_955_v21_ = rhs194_
                        lhs173_[lhs174_] = rhs195_
                        lhs175_.f0 = rhs196_
                        d_962_v28_ = rhs197_
                        d_965_v29_: _dafny.Map
                        d_965_v29_ = _dafny.Map({d_954_v20_: (d_954_v20_)[default__.safeIndex(723, (d_954_v20_).length(0))]})
                        d_966_v30_: _dafny.Seq
                        d_966_v30_ = _dafny.SeqWithoutIsStrInference([d_965_v29_, d_965_v29_, d_965_v29_, d_965_v29_, d_965_v29_])
                        d_967_v31_: C2
                        nw166_ = C2()
                        nw166_.ctor__((self).f26, not ((self).f25) or (False), (len((d_966_v30_)[default__.safeIndex((self).f26, len(d_966_v30_))])) * (len(d_955_v21_)))
                        d_967_v31_ = nw166_
                        d_929_v0_ = (d_961_v27_)[default__.safeIndex(840, (d_961_v27_).length(0))]
                        index182_ = default__.safeIndex(587, (d_959_v25_).length(0))
                        (d_959_v25_)[index182_] = ((d_960_v26_ if p1 else (d_959_v25_)[default__.safeIndex(587, (d_959_v25_).length(0))])) + ((d_959_v25_)[default__.safeIndex(587, (d_959_v25_).length(0))])
                    elif True:
                        d_968_v32_: D0
                        d_968_v32_ = D0_DC0(p2)
                        d_969_v33_: _dafny.Map
                        d_969_v33_ = _dafny.Map({d_968_v32_: 649})
                        d_970_v34_: _dafny.Map
                        d_970_v34_ = _dafny.Map({326: (0) - (len((self).f29))})
                        d_971_v35_: _dafny.MultiSet
                        d_971_v35_ = _dafny.MultiSet([(d_954_v20_)[default__.safeIndex(723, (d_954_v20_).length(0))], default__.fm5((self).f25, d_969_v33_, ((d_970_v34_)[(self).f26] if ((self).f26) in (d_970_v34_) else (self).f26), globalState)])
                        (globalState).f1 = not((True) in (d_971_v35_))
                        d_972_v36_: _dafny.Array
                        nw167_ = _dafny.Array(int(0), 22)
                        d_972_v36_ = nw167_
                        d_973_v37_: _dafny.Seq
                        d_973_v37_ = _dafny.SeqWithoutIsStrInference([d_972_v36_])
                        d_972_v36_ = (d_972_v36_ if default__.fm5(p1, d_969_v33_, (self).f26, globalState) else (d_973_v37_)[default__.safeIndex((self).f26, len(d_973_v37_))])
                        d_974_v38_: _dafny.Array
                        nw168_ = _dafny.Array(_dafny.Set({}), 17)
                        d_974_v38_ = nw168_
                        d_975_v39_: _dafny.Map
                        d_975_v39_ = _dafny.Map({(self).f26: d_971_v35_})
                        d_976_v40_: D6
                        d_976_v40_ = D6_DC14((self).f26, (self).f26, ((d_975_v39_)[796] if (796) in (d_975_v39_) else (d_971_v35_).set((self).f25, default__.abs((self).f26))))
                        d_977_v41_: D9
                        d_977_v41_ = D9_DC19(d_976_v40_)
                        d_978_v42_: D9
                        d_978_v42_ = D9_DC20(d_977_v41_)
                        d_979_v43_: D0
                        d_979_v43_ = D0_DC1(-977, (self).f29, p1, (self).f26, (d_955_v21_)[default__.safeIndex((self).f26, len(d_955_v21_))])
                        pat_let_tv27_ = p1
                        pat_let_tv28_ = p1
                        d_980_v44_: _dafny.Map
                        def iife78_(_pat_let28_0):
                            def iife79_(d_981_dt__update__tmp_h2_):
                                def iife80_(_pat_let29_0):
                                    def iife81_(d_982_dt__update_hcf3_h0_):
                                        def iife82_(_pat_let30_0):
                                            def iife83_(d_983_dt__update_hcf4_h0_):
                                                def iife84_(_pat_let31_0):
                                                    def iife85_(d_984_dt__update_hcf5_h0_):
                                                        return D0_DC1((d_981_dt__update__tmp_h2_).cf1, (d_981_dt__update__tmp_h2_).cf2, d_982_dt__update_hcf3_h0_, d_983_dt__update_hcf4_h0_, d_984_dt__update_hcf5_h0_)
                                                    return iife85_(_pat_let31_0)
                                                return iife84_(pat_let_tv28_)
                                            return iife83_(_pat_let30_0)
                                        return iife82_((self).f26)
                                    return iife81_(_pat_let29_0)
                                return iife80_(pat_let_tv27_)
                            return iife79_(_pat_let28_0)
                        d_980_v44_ = _dafny.Map({d_978_v42_: _dafny.Set({d_929_v0_, default__.fm4(iife78_(d_979_v43_), (d_954_v20_)[default__.safeIndex(723, (d_954_v20_).length(0))], (self).f26, globalState), d_929_v0_, d_929_v0_})})
                        d_985_v45_: _dafny.Set
                        d_985_v45_ = _dafny.Set({d_929_v0_, d_929_v0_})
                        pat_let_tv29_ = d_978_v42_
                        pat_let_tv30_ = d_978_v42_
                        index183_ = default__.safeIndex(404, (d_974_v38_).length(0))
                        def iife86_(_pat_let32_0):
                            def iife87_(d_986_dt__update__tmp_h4_):
                                def iife88_(_pat_let33_0):
                                    def iife89_(d_987_dt__update_hcf32_h1_):
                                        return D9_DC20(d_987_dt__update_hcf32_h1_)
                                    return iife89_(_pat_let33_0)
                                return iife88_((pat_let_tv29_).cf32)
                            return iife87_(_pat_let32_0)
                        def iife90_(_pat_let34_0):
                            def iife91_(d_988_dt__update__tmp_h3_):
                                def iife92_(_pat_let35_0):
                                    def iife93_(d_989_dt__update_hcf32_h0_):
                                        return D9_DC20(d_989_dt__update_hcf32_h0_)
                                    return iife93_(_pat_let35_0)
                                return iife92_((pat_let_tv30_).cf32)
                            return iife91_(_pat_let34_0)
                        (d_974_v38_)[index183_] = ((d_980_v44_)[iife86_(d_978_v42_)] if (iife90_(d_978_v42_)) in (d_980_v44_) else d_985_v45_)
                        index184_ = default__.safeIndex(404, (d_974_v38_).length(0))
                        (d_974_v38_)[index184_] = default__.fm31(globalState)
                        index185_ = default__.safeIndex(952, (d_972_v36_).length(0))
                        (d_972_v36_)[index185_] = (self).f26
                        index186_ = default__.safeIndex(952, (d_972_v36_).length(0))
                        (d_972_v36_)[index186_] = 625
                        d_990_v46_: _dafny.Seq
                        d_990_v46_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                        d_991_v47_: _dafny.Set
                        d_991_v47_ = _dafny.Set({(self).fm7(globalState), p1})
                        d_992_v48_: _dafny.Array
                        nw169_ = _dafny.Array(None, 21)
                        nw169_[int(0)] = d_990_v46_
                        nw169_[int(1)] = default__.fm3((d_972_v36_)[default__.safeIndex(952, (d_972_v36_).length(0))], False, (self).f26, True, globalState)
                        nw169_[int(2)] = (d_990_v46_) + (d_990_v46_)
                        nw169_[int(3)] = _dafny.SeqWithoutIsStrInference([352])
                        nw169_[int(4)] = d_990_v46_
                        nw169_[int(5)] = d_990_v46_
                        nw169_[int(6)] = d_990_v46_
                        nw169_[int(7)] = d_990_v46_
                        nw169_[int(8)] = d_990_v46_
                        nw169_[int(9)] = default__.fm3((self).f26, p2, len(d_991_v47_), False, globalState)
                        nw169_[int(10)] = (_dafny.SeqWithoutIsStrInference([(self).f26, (d_972_v36_)[default__.safeIndex(952, (d_972_v36_).length(0))], 784])) + (d_990_v46_)
                        nw169_[int(11)] = d_990_v46_
                        nw169_[int(12)] = d_990_v46_
                        nw169_[int(13)] = (_dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])) + (d_990_v46_)
                        nw169_[int(14)] = d_990_v46_
                        nw169_[int(15)] = (d_990_v46_).set(default__.safeIndex((self).f26, len(d_990_v46_)), (self).f26)
                        nw169_[int(16)] = d_990_v46_
                        nw169_[int(17)] = d_990_v46_
                        nw169_[int(18)] = _dafny.SeqWithoutIsStrInference([(0) - ((self).f26)])
                        nw169_[int(19)] = d_990_v46_
                        nw169_[int(20)] = _dafny.SeqWithoutIsStrInference([(d_972_v36_)[default__.safeIndex(952, (d_972_v36_).length(0))]])
                        d_992_v48_ = nw169_
                        index187_ = default__.safeIndex(689, (d_992_v48_).length(0))
                        (d_992_v48_)[index187_] = d_990_v46_
                        index188_ = default__.safeIndex(689, (d_992_v48_).length(0))
                        (d_992_v48_)[index188_] = d_990_v46_
                    d_993_v49_: C2
                    nw170_ = C2()
                    nw170_.ctor__(144, (p2) == ((self).f25), ((self).f26) * ((self).f26))
                    d_993_v49_ = nw170_
                    d_994_v50_: _dafny.Map
                    d_994_v50_ = _dafny.Map({(d_954_v20_)[default__.safeIndex(723, (d_954_v20_).length(0))]: (self).f26})
                    d_994_v50_ = (d_994_v50_).set(p1, (0) - ((d_993_v49_).f31))
                    pass
            pass
        (globalState).f0 = (self).f26
        d_995_i3_: int
        d_995_i3_ = 0
        with _dafny.label("5"):
            while (len(_dafny.SeqWithoutIsStrInference([True]))) >= (default__.safeDivisionInt((self).f26, default__.fm2((self).f26, (self).f26, globalState))):
                with _dafny.c_label("5"):
                    if (d_995_i3_) >= (100):
                        raise _dafny.Break("5")
                    d_995_i3_ = (d_995_i3_) + (1)
                    d_996_v51_: int
                    out23_: int
                    out23_ = default__.m0((self).f25, globalState)
                    d_996_v51_ = out23_
                    (globalState).f0 = -983
                    d_930_v1_ = (d_930_v1_).set((self).f25, p2)
                    d_929_v0_ = d_929_v0_
                    pass
            pass
        hi9_ = (self).f26
        for d_997_i4_ in range(((self).f26) * ((0) - ((self).f26)), hi9_):
            d_998_v52_: _dafny.Array
            def lambda52_(d_999_p1_):
                def lambda53_(d_1000_i5_):
                    return d_999_p1_

                return lambda53_

            init26_ = lambda52_(p1)
            nw171_ = _dafny.Array(None, 2)
            for i0_26_ in range(nw171_.length(0)):
                nw171_[i0_26_] = init26_(i0_26_)
            d_998_v52_ = nw171_
            d_1001_v53_: _dafny.Seq
            d_1001_v53_ = _dafny.SeqWithoutIsStrInference([d_998_v52_])
            d_1002_v54_: _dafny.Map
            d_1002_v54_ = _dafny.Map({(self).f30: d_998_v52_})
            d_1001_v53_ = _dafny.SeqWithoutIsStrInference([((d_1002_v54_)[not(not(p2))] if (not(not(p2))) in (d_1002_v54_) else d_998_v52_), d_998_v52_, d_998_v52_, d_998_v52_])
            (globalState).f0 = (d_997_i4_) + ((self).f26)
            d_1003_v55_: _dafny.Array
            nw172_ = _dafny.Array(D2.default()(), 17)
            d_1003_v55_ = nw172_
            d_1004_v56_: _dafny.Map
            d_1004_v56_ = _dafny.Map({p1: (self).f26})
            d_1005_v57_: _dafny.MultiSet
            d_1005_v57_ = _dafny.MultiSet([d_1004_v56_, d_1004_v56_])
            d_1006_v58_: _dafny.MultiSet
            d_1006_v58_ = d_1005_v57_
            rhs198_ = d_1003_v55_
            rhs199_ = d_1006_v58_
            rhs200_ = (self).f26
            lhs176_ = globalState
            d_1003_v55_ = rhs198_
            d_1006_v58_ = rhs199_
            lhs176_.f2 = rhs200_
            d_1007_v59_: _dafny.Array
            def lambda54_(d_1008_i6_):
                return _dafny.SeqWithoutIsStrInference([(self).f26])

            init27_ = lambda54_
            nw173_ = _dafny.Array(None, 5)
            for i0_27_ in range(nw173_.length(0)):
                nw173_[i0_27_] = init27_(i0_27_)
            d_1007_v59_ = nw173_
            d_1009_v60_: _dafny.Seq
            d_1009_v60_ = _dafny.SeqWithoutIsStrInference([d_997_i4_])
            index189_ = default__.safeIndex(635, (d_1007_v59_).length(0))
            (d_1007_v59_)[index189_] = d_1009_v60_
            index190_ = default__.safeIndex(635, (d_1007_v59_).length(0))
            (d_1007_v59_)[index190_] = d_1009_v60_
        d_1010_v61_: D1
        d_1010_v61_ = D1_DC3(d_929_v0_)
        d_1011_v62_: _dafny.Seq
        d_1011_v62_ = _dafny.SeqWithoutIsStrInference([d_1010_v61_])
        d_1012_v63_: _dafny.Map
        d_1012_v63_ = _dafny.Map({(self).f26: d_1011_v62_})
        r0 = ((((d_1012_v63_)[(self).f26] if ((self).f26) in (d_1012_v63_) else d_1011_v62_)) + (d_1011_v62_)) <= (_dafny.SeqWithoutIsStrInference([d_1010_v61_ for d_1013_i7_ in range(default__.abs(304))]))
        return r0

    def m4(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        source15_ = default__.fm32(globalState)
        if source15_.is_DC10:
            d_1014___mcc_h0_ = source15_.cf16
            d_1015___mcc_h1_ = source15_.cf17
            d_1016_cf17_ = d_1015___mcc_h1_
            d_1017_cf16_ = d_1014___mcc_h0_
            d_1018_v0_: _dafny.MultiSet
            d_1018_v0_ = _dafny.MultiSet([(self).f26, (self).f26])
            d_1019_v1_: _dafny.MultiSet
            d_1019_v1_ = d_1018_v0_
            d_1020_v2_: _dafny.Map
            d_1020_v2_ = _dafny.Map({p1: p1})
            d_1021_v3_: _dafny.Map
            d_1021_v3_ = _dafny.Map({(0) - (p0): not(p1)})
            d_1022_v4_: _dafny.Array
            nw174_ = _dafny.Array(None, 21)
            nw174_[int(0)] = True
            nw174_[int(1)] = ((d_1020_v2_)[p1] if (p1) in (d_1020_v2_) else False)
            nw174_[int(2)] = (self).f25
            nw174_[int(3)] = p1
            nw174_[int(4)] = p2
            nw174_[int(5)] = p1
            nw174_[int(6)] = (self).f25
            nw174_[int(7)] = p1
            nw174_[int(8)] = ((d_1021_v3_)[-745] if (-745) in (d_1021_v3_) else p2)
            nw174_[int(9)] = p2
            nw174_[int(10)] = not(p1)
            nw174_[int(11)] = p2
            nw174_[int(12)] = (self).f25
            nw174_[int(13)] = True
            nw174_[int(14)] = (self).f25
            nw174_[int(15)] = p1
            nw174_[int(16)] = (self).f30
            nw174_[int(17)] = p2
            nw174_[int(18)] = p2
            nw174_[int(19)] = p2
            nw174_[int(20)] = p1
            d_1022_v4_ = nw174_
            d_1023_v5_: T0
            nw175_ = C0()
            nw175_.ctor__((self).f29, True, p0, d_1022_v4_, (self).f30)
            d_1023_v5_ = nw175_
            d_1024_v6_: _dafny.Set
            d_1024_v6_ = _dafny.Set({(self).f25})
            d_1025_v7_: _dafny.Map
            d_1025_v7_ = _dafny.Map({d_1023_v5_: len(d_1024_v6_)})
            d_1026_v8_: _dafny.Seq
            d_1026_v8_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, ((d_1025_v7_)[d_1023_v5_] if (d_1023_v5_) in (d_1025_v7_) else (self).f26)])
            d_1027_v9_: _dafny.MultiSet
            d_1027_v9_ = _dafny.MultiSet([_dafny.MultiSet(d_1026_v8_)])
            d_1028_v10_: _dafny.Array
            def lambda55_(d_1029_v2_):
                def lambda56_(d_1030_i0_):
                    return d_1029_v2_

                return lambda56_

            init28_ = lambda55_(d_1020_v2_)
            nw176_ = _dafny.Array(None, 6)
            for i0_28_ in range(nw176_.length(0)):
                nw176_[i0_28_] = init28_(i0_28_)
            d_1028_v10_ = nw176_
            d_1031_v11_: _dafny.Map
            d_1031_v11_ = _dafny.Map({(d_1023_v5_).f26: D8_DC16(d_1028_v10_)})
            d_1032_v12_: _dafny.MultiSet
            d_1032_v12_ = _dafny.MultiSet([d_1031_v11_, d_1031_v11_])
            d_1033_v13_: _dafny.MultiSet
            d_1033_v13_ = _dafny.MultiSet([default__.fm1(globalState)])
            d_1034_v14_: _dafny.Array
            nw177_ = _dafny.Array(None, 21)
            nw177_[int(0)] = (self).f26
            nw177_[int(1)] = (d_1027_v9_).cardinality
            nw177_[int(2)] = (d_1023_v5_).f26
            nw177_[int(3)] = (self).f26
            nw177_[int(4)] = ((0) - (((d_1032_v12_)[d_1031_v11_] if (d_1031_v11_) in (d_1032_v12_) else (d_1023_v5_).f26))) + ((self).f26)
            nw177_[int(5)] = ((d_1033_v13_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gesulblvl"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gesulblvl"))) in (d_1033_v13_) else (0) - ((d_1023_v5_).f26))
            nw177_[int(6)] = (self).f26
            nw177_[int(7)] = (len(d_1021_v3_) if (self).f30 else p0)
            nw177_[int(8)] = (d_1023_v5_).f26
            nw177_[int(9)] = p0
            nw177_[int(10)] = (d_1026_v8_)[default__.safeIndex((d_1023_v5_).f26, len(d_1026_v8_))]
            nw177_[int(11)] = (d_1023_v5_).f26
            nw177_[int(12)] = -274
            nw177_[int(13)] = len(d_1026_v8_)
            nw177_[int(14)] = len(((self).f29) + ((self).f29))
            nw177_[int(15)] = (self).f26
            nw177_[int(16)] = (312) - ((d_1023_v5_).f26)
            nw177_[int(17)] = p0
            nw177_[int(18)] = (self).fm6((self).f25, (self).f25, 132, globalState)
            nw177_[int(19)] = (len(_dafny.Map({(self).f25: p1}))) - ((d_1023_v5_).f26)
            nw177_[int(20)] = (self).f26
            d_1034_v14_ = nw177_
            index191_ = default__.safeIndex(661, (d_1034_v14_).length(0))
            (d_1034_v14_)[index191_] = p0
            index192_ = default__.safeIndex(661, (d_1034_v14_).length(0))
            (d_1034_v14_)[index192_] = len(((self).f29) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ix"))))
            d_1035_v15_: _dafny.Set
            d_1035_v15_ = _dafny.Set({((d_1034_v14_)[default__.safeIndex(661, (d_1034_v14_).length(0))] if p2 else (d_1023_v5_).f26)})
            d_1036_v16_: _dafny.Map
            d_1036_v16_ = _dafny.Map({(d_1023_v5_).f26: default__.fm1(globalState)})
            d_1037_v17_: _dafny.Seq
            d_1037_v17_ = _dafny.SeqWithoutIsStrInference([p1, p2, p1, (len((d_1036_v16_).set((d_1034_v14_)[default__.safeIndex(661, (d_1034_v14_).length(0))], (self).f29))) < ((d_1023_v5_).f26)])
            d_1038_v18_: _dafny.Array
            nw178_ = _dafny.Array(D6.default()(), 9)
            d_1038_v18_ = nw178_
            d_1039_v19_: D6
            d_1039_v19_ = D6_DC13(default__.fm2(default__.fm2((self).f26, 102, globalState), (d_1034_v14_)[default__.safeIndex(661, (d_1034_v14_).length(0))], globalState), (d_1023_v5_).f26, d_1035_v15_)
            index193_ = default__.safeIndex(649, (d_1038_v18_).length(0))
            (d_1038_v18_)[index193_] = d_1039_v19_
            pat_let_tv31_ = d_1034_v14_
            pat_let_tv32_ = d_1034_v14_
            index194_ = default__.safeIndex(649, (d_1038_v18_).length(0))
            def iife94_(_pat_let36_0):
                def iife95_(d_1040_dt__update__tmp_h1_):
                    def iife96_(_pat_let37_0):
                        def iife97_(d_1041_dt__update_hcf21_h0_):
                            return D6_DC13((d_1040_dt__update__tmp_h1_).cf20, d_1041_dt__update_hcf21_h0_, (d_1040_dt__update__tmp_h1_).cf22)
                        return iife97_(_pat_let37_0)
                    return iife96_((pat_let_tv32_)[default__.safeIndex(661, (pat_let_tv31_).length(0))])
                return iife95_(_pat_let36_0)
            rhs201_ = d_1035_v15_
            rhs202_ = d_1020_v2_
            rhs203_ = _dafny.SeqWithoutIsStrInference([(d_1023_v5_).f25])
            rhs204_ = iife94_(d_1039_v19_)
            rhs205_ = d_1017_cf16_
            lhs177_ = d_1038_v18_
            lhs178_ = default__.safeIndex(649, (d_1038_v18_).length(0))
            d_1035_v15_ = rhs201_
            d_1020_v2_ = rhs202_
            d_1037_v17_ = rhs203_
            lhs177_[lhs178_] = rhs204_
            d_1017_cf16_ = rhs205_
            index195_ = default__.safeIndex(661, (d_1034_v14_).length(0))
            (d_1034_v14_)[index195_] = (d_1034_v14_)[default__.safeIndex(661, (d_1034_v14_).length(0))]
            (globalState).f1 = p2
        elif True:
            d_1042___mcc_h2_ = source15_.cf15
            d_1043_cf15_ = d_1042___mcc_h2_
            d_1044_v20_: _dafny.Map
            d_1044_v20_ = _dafny.Map({(self).f26: p2})
            r1 = ((self).f26 if not(((d_1044_v20_)[(self).f26] if ((self).f26) in (d_1044_v20_) else (self).f30)) else (self).f26)
            d_1045_v21_: _dafny.Array
            def lambda57_(d_1046_p1_):
                def lambda58_(d_1047_i1_):
                    return d_1046_p1_

                return lambda58_

            init29_ = lambda57_(p1)
            nw179_ = _dafny.Array(None, 2)
            for i0_29_ in range(nw179_.length(0)):
                nw179_[i0_29_] = init29_(i0_29_)
            d_1045_v21_ = nw179_
            index196_ = default__.safeIndex(717, (d_1045_v21_).length(0))
            (d_1045_v21_)[index196_] = ((self).f29) == ((self).f29)
            d_1048_v22_: _dafny.Map
            d_1048_v22_ = _dafny.Map({_dafny.Map({len(d_1044_v20_): p2}): p1})
            d_1049_v23_: str
            d_1049_v23_ = _dafny.CodePoint('d')
            d_1050_v24_: _dafny.Map
            d_1050_v24_ = _dafny.Map({d_1049_v23_: (self).f26})
            index197_ = default__.safeIndex(717, (d_1045_v21_).length(0))
            (d_1045_v21_)[index197_] = ((len(d_1048_v22_)) + (p0)) not in (_dafny.MultiSet([len(d_1050_v24_)]))
            d_1051_v25_: C1
            nw180_ = C1()
            nw180_.ctor__(p1)
            d_1051_v25_ = nw180_
            d_1052_v26_: C2
            nw181_ = C2()
            nw181_.ctor__(p0, (d_1045_v21_)[default__.safeIndex(717, (d_1045_v21_).length(0))], -12)
            d_1052_v26_ = nw181_
        hi10_ = p0
        for d_1053_i2_ in range((606 if p2 else (self).f26), hi10_):
            d_1054_v27_: _dafny.Array
            nw182_ = _dafny.Array(int(0), 27)
            d_1054_v27_ = nw182_
            index198_ = default__.safeIndex(712, (d_1054_v27_).length(0))
            (d_1054_v27_)[index198_] = d_1053_i2_
            d_1055_v28_: _dafny.MultiSet
            d_1055_v28_ = _dafny.MultiSet([len(default__.fm3(d_1053_i2_, p2, -651, p1, globalState))])
            d_1056_v29_: _dafny.MultiSet
            d_1056_v29_ = _dafny.MultiSet([d_1055_v28_])
            index199_ = default__.safeIndex(712, (d_1054_v27_).length(0))
            (d_1054_v27_)[index199_] = ((d_1056_v29_)[default__.fm21((self).f29, (self).f26, p0, globalState)] if (default__.fm21((self).f29, (self).f26, p0, globalState)) in (d_1056_v29_) else default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_1053_i2_ for d_1057_i3_ in range(default__.abs(833))])), len((self).f29)))
            d_1058_v30_: _dafny.MultiSet
            d_1058_v30_ = _dafny.MultiSet([(self).f30])
            d_1059_v31_: C2
            nw183_ = C2()
            nw183_.ctor__((d_1058_v30_).cardinality, False, (0) - ((self).f26))
            d_1059_v31_ = nw183_
            d_1060_v32_: _dafny.MultiSet
            d_1060_v32_ = _dafny.MultiSet([d_1059_v31_])
            if (_dafny.MultiSet([d_1059_v31_, d_1059_v31_])).ispropersubset(d_1060_v32_):
                r1 = p0
                d_1061_v33_: _dafny.Array
                nw184_ = _dafny.Array(_dafny.Seq({}), 5)
                d_1061_v33_ = nw184_
                d_1062_v34_: _dafny.Seq
                d_1062_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({False: (d_1054_v27_)[default__.safeIndex(712, (d_1054_v27_).length(0))]})])
                index200_ = default__.safeIndex(926, (d_1061_v33_).length(0))
                (d_1061_v33_)[index200_] = d_1062_v34_
                index201_ = default__.safeIndex(926, (d_1061_v33_).length(0))
                (d_1061_v33_)[index201_] = d_1062_v34_
                d_1063_v35_: D0
                d_1063_v35_ = D0_DC0((self).f30)
                (globalState).f1 = default__.fm5((self).f30, _dafny.Map({d_1063_v35_: (self).f26}), (self).f26, globalState)
                d_1064_v36_: _dafny.Set
                d_1064_v36_ = _dafny.Set({(d_1054_v27_)[default__.safeIndex(712, (d_1054_v27_).length(0))]})
                d_1065_v37_: _dafny.Seq
                d_1065_v37_ = _dafny.SeqWithoutIsStrInference([(d_1059_v31_).f31])
                r0 = (len(d_1064_v36_)) - (len(d_1065_v37_))
                d_1066_v38_: D0
                d_1066_v38_ = D0_DC1((0) - (d_1053_i2_), (self).f29, p1, 145, p1)
                (globalState).f0 = (d_1066_v38_).cf1
            elif True:
                d_1067_v39_: D1
                d_1067_v39_ = D1_DC4((self).f30, (self).f25)
                d_1068_v40_: _dafny.Seq
                d_1068_v40_ = _dafny.SeqWithoutIsStrInference([p1, (d_1067_v39_).cf7, (self).f30])
                (globalState).f2 = (len(d_1068_v40_)) * (((d_1054_v27_)[default__.safeIndex(712, (d_1054_v27_).length(0))] if p2 else 720))
                d_1069_v41_: _dafny.Seq
                d_1069_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pq"))
                d_1070_v42_: _dafny.MultiSet
                d_1070_v42_ = _dafny.MultiSet([d_1069_v41_, (self).f29, (self).f29])
                d_1071_v43_: _dafny.Set
                d_1071_v43_ = _dafny.Set({(d_1054_v27_)[default__.safeIndex(712, (d_1054_v27_).length(0))]})
                d_1072_v44_: _dafny.Array
                nw185_ = _dafny.Array(None, 23)
                nw185_[int(0)] = ((d_1059_v31_).f31) <= (-210)
                nw185_[int(1)] = (d_1070_v42_).ispropersubset(_dafny.MultiSet([(self).f29, d_1069_v41_, d_1069_v41_]))
                nw185_[int(2)] = p2
                nw185_[int(3)] = (self).f30
                nw185_[int(4)] = not((p1) and ((self).f25))
                nw185_[int(5)] = ((self).f25 if True else p2)
                nw185_[int(6)] = (self).f30
                nw185_[int(7)] = (d_1055_v28_).issubset(d_1055_v28_)
                nw185_[int(8)] = (self).f30
                nw185_[int(9)] = (self).f25
                nw185_[int(10)] = not ((self).f30) or ((self).f25)
                nw185_[int(11)] = p2
                nw185_[int(12)] = (self).f25
                nw185_[int(13)] = p1
                nw185_[int(14)] = (self).f25
                nw185_[int(15)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwtpj"))) != (d_1069_v41_)
                nw185_[int(16)] = not((self).f30)
                nw185_[int(17)] = (self).f25
                nw185_[int(18)] = p1
                nw185_[int(19)] = (_dafny.Set({(d_1059_v31_).f31})).ispropersubset(d_1071_v43_)
                nw185_[int(20)] = (self).f25
                nw185_[int(21)] = not((self).f30)
                nw185_[int(22)] = p1
                d_1072_v44_ = nw185_
                index202_ = default__.safeIndex(540, (d_1072_v44_).length(0))
                (d_1072_v44_)[index202_] = ((d_1058_v30_).cardinality) < (d_1053_i2_)
                d_1073_v45_: _dafny.Array
                def lambda59_(d_1074_v28_):
                    def lambda60_(d_1075_i4_):
                        return d_1074_v28_

                    return lambda60_

                init30_ = lambda59_(d_1055_v28_)
                nw186_ = _dafny.Array(None, 25)
                for i0_30_ in range(nw186_.length(0)):
                    nw186_[i0_30_] = init30_(i0_30_)
                d_1073_v45_ = nw186_
                d_1076_v46_: D10
                d_1076_v46_ = D10_DC21(d_1073_v45_)
                index203_ = default__.safeIndex(540, (d_1072_v44_).length(0))
                index204_ = default__.safeIndex(712, (d_1054_v27_).length(0))
                rhs206_ = ((self).f29) + ((self).f29)
                rhs207_ = True
                rhs208_ = default__.safeDivisionInt(-169, ((0) - ((d_1054_v27_)[default__.safeIndex(712, (d_1054_v27_).length(0))])) * (d_1053_i2_))
                rhs209_ = (d_1076_v46_).cf33
                rhs210_ = (self).f25
                lhs179_ = d_1072_v44_
                lhs180_ = default__.safeIndex(540, (d_1072_v44_).length(0))
                lhs181_ = d_1054_v27_
                lhs182_ = default__.safeIndex(712, (d_1054_v27_).length(0))
                lhs183_ = globalState
                d_1069_v41_ = rhs206_
                lhs179_[lhs180_] = rhs207_
                lhs181_[lhs182_] = rhs208_
                d_1073_v45_ = rhs209_
                lhs183_.f1 = rhs210_
                d_1077_v47_: _dafny.Seq
                d_1077_v47_ = _dafny.SeqWithoutIsStrInference([d_1054_v27_])
                d_1077_v47_ = ((d_1077_v47_).set(default__.safeIndex((d_1059_v31_).f31, len(d_1077_v47_)), d_1054_v27_)) + ((d_1077_v47_) + (d_1077_v47_))
                (globalState).f1 = not((self).f25)
                d_1078_v48_: D0
                d_1078_v48_ = D0_DC0((d_1072_v44_)[default__.safeIndex(540, (d_1072_v44_).length(0))])
                (globalState).f1 = (d_1078_v48_).cf0
            d_1079_v49_: _dafny.Array
            def lambda61_(d_1080_i5_):
                return (self).f25

            init31_ = lambda61_
            nw187_ = _dafny.Array(None, 10)
            for i0_31_ in range(nw187_.length(0)):
                nw187_[i0_31_] = init31_(i0_31_)
            d_1079_v49_ = nw187_
            index205_ = default__.safeIndex(398, (d_1079_v49_).length(0))
            (d_1079_v49_)[index205_] = p2
            d_1081_v50_: _dafny.Array
            def lambda62_(d_1082_p2_):
                def lambda63_(d_1083_i6_):
                    return _dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([d_1082_p2_, (self).f30]): _dafny.CodePoint('p')})), (self).f26])

                return lambda63_

            init32_ = lambda62_(p2)
            nw188_ = _dafny.Array(None, 23)
            for i0_32_ in range(nw188_.length(0)):
                nw188_[i0_32_] = init32_(i0_32_)
            d_1081_v50_ = nw188_
            d_1084_v51_: _dafny.Seq
            d_1084_v51_ = _dafny.SeqWithoutIsStrInference([p1])
            index206_ = default__.safeIndex(398, (d_1079_v49_).length(0))
            rhs211_ = (d_1084_v51_) < (_dafny.SeqWithoutIsStrInference([(self).f30, (self).f30]))
            rhs212_ = (d_1059_v31_).fm7(globalState)
            rhs213_ = d_1081_v50_
            lhs184_ = d_1079_v49_
            lhs185_ = default__.safeIndex(398, (d_1079_v49_).length(0))
            r2 = rhs211_
            lhs184_[lhs185_] = rhs212_
            d_1081_v50_ = rhs213_
        if (self).f25:
            d_1085_v52_: _dafny.Map
            d_1085_v52_ = _dafny.Map({p0: (self).fm7(globalState)})
            d_1086_v53_: _dafny.MultiSet
            d_1086_v53_ = _dafny.MultiSet([(self).f26])
            d_1087_v54_: _dafny.Seq
            d_1087_v54_ = _dafny.SeqWithoutIsStrInference([((d_1085_v52_)[(d_1086_v53_).cardinality] if ((d_1086_v53_).cardinality) in (d_1085_v52_) else (self).f30)])
            d_1088_v55_: _dafny.Map
            d_1088_v55_ = _dafny.Map({d_1087_v54_: (self).f30})
            d_1089_v56_: _dafny.Map
            d_1089_v56_ = _dafny.Map({p2: d_1088_v55_})
            d_1090_v57_: _dafny.MultiSet
            d_1090_v57_ = _dafny.MultiSet([True, (self).f25, p1])
            d_1091_v58_: D6
            d_1091_v58_ = D6_DC14((self).f26, -684, d_1090_v57_)
            r0 = len((d_1089_v56_) | ((default__.fm33(d_1091_v58_, p0, globalState)) | (d_1089_v56_)))
            (globalState).f1 = ((self).f26) < ((self).f26)
            d_1092_v59_: _dafny.Array
            def lambda64_(d_1093_p0_):
                def lambda65_(d_1094_i7_):
                    return default__.safeModuloInt(d_1094_i7_, d_1093_p0_)

                return lambda65_

            init33_ = lambda64_(p0)
            nw189_ = _dafny.Array(None, 1)
            for i0_33_ in range(nw189_.length(0)):
                nw189_[i0_33_] = init33_(i0_33_)
            d_1092_v59_ = nw189_
            index207_ = default__.safeIndex(303, (d_1092_v59_).length(0))
            (d_1092_v59_)[index207_] = (self).f26
            index208_ = default__.safeIndex(303, (d_1092_v59_).length(0))
            (d_1092_v59_)[index208_] = default__.safeModuloInt(default__.safeModuloInt(221, p0), (self).f26)
            r0 = (self).f26
            d_1095_v60_: _dafny.Array
            def lambda66_(d_1096_p2_):
                def lambda67_(d_1097_i8_):
                    return (D11_DC25(_dafny.Map({d_1096_p2_: d_1096_p2_}))).cf39

                return lambda67_

            init34_ = lambda66_(p2)
            nw190_ = _dafny.Array(None, 26)
            for i0_34_ in range(nw190_.length(0)):
                nw190_[i0_34_] = init34_(i0_34_)
            d_1095_v60_ = nw190_
            d_1098_v61_: D8
            d_1098_v61_ = D8_DC16(d_1095_v60_)
            d_1099_v62_: _dafny.Map
            d_1099_v62_ = _dafny.Map({d_1087_v54_: d_1098_v61_})
            d_1100_v63_: C2
            nw191_ = C2()
            nw191_.ctor__(len(d_1099_v62_), p1, (0) - (len(d_1087_v54_)))
            d_1100_v63_ = nw191_
        elif True:
            d_1101_v64_: _dafny.Map
            d_1101_v64_ = _dafny.Map({(self).f26: 545})
            d_1101_v64_ = d_1101_v64_
            d_1102_v65_: _dafny.Array
            def lambda68_(d_1103_i9_):
                return default__.safeModuloInt(d_1103_i9_, 843)

            init35_ = lambda68_
            nw192_ = _dafny.Array(None, 7)
            for i0_35_ in range(nw192_.length(0)):
                nw192_[i0_35_] = init35_(i0_35_)
            d_1102_v65_ = nw192_
            d_1104_v66_: _dafny.Array
            def lambda69_(d_1105_i10_):
                return _dafny.SeqWithoutIsStrInference([(self).f29, (self).f29, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1106_i11_ in range(default__.abs(128))])])

            init36_ = lambda69_
            nw193_ = _dafny.Array(None, 5)
            for i0_36_ in range(nw193_.length(0)):
                nw193_[i0_36_] = init36_(i0_36_)
            d_1104_v66_ = nw193_
            d_1107_v67_: _dafny.Seq
            d_1107_v67_ = _dafny.SeqWithoutIsStrInference([(self).f29])
            index209_ = default__.safeIndex(56, (d_1104_v66_).length(0))
            (d_1104_v66_)[index209_] = (d_1107_v67_) + (d_1107_v67_)
            d_1108_v68_: str
            d_1108_v68_ = _dafny.CodePoint('c')
            d_1109_v69_: _dafny.MultiSet
            d_1109_v69_ = _dafny.MultiSet([397, (self).f26, (self).f26])
            d_1110_v70_: _dafny.Map
            d_1110_v70_ = _dafny.Map({default__.fm28((self).f29, d_1108_v68_, (self).f26, globalState): len((self).f29)})
            d_1111_v73_: D0
            d_1111_v73_ = D0_DC0((self).f30)
            d_1112_v74_: _dafny.Seq
            d_1112_v74_ = _dafny.SeqWithoutIsStrInference([d_1111_v73_, d_1111_v73_])
            index210_ = default__.safeIndex(56, (d_1104_v66_).length(0))
            def iife98_():
                def iife100_():
                    coll24_ = _dafny.Map()
                    compr_24_: D0
                    for compr_24_ in (d_1112_v74_).Elements:
                        d_1113_v72_: D0 = compr_24_
                        if (d_1113_v72_) in (d_1112_v74_):
                            coll24_[d_1113_v72_] = len(_dafny.Map({(self).f30: (self).f30}))
                    return _dafny.Map(coll24_)
                coll22_ = _dafny.Map()
                def iife99_():
                    coll23_ = _dafny.Map()
                    compr_23_: D0
                    for compr_23_ in (d_1112_v74_).Elements:
                        d_1113_v72_: D0 = compr_23_
                        if (d_1113_v72_) in (d_1112_v74_):
                            coll23_[d_1113_v72_] = len(_dafny.Map({(self).f30: (self).f30}))
                    return _dafny.Map(coll23_)
                compr_22_: D0
                for compr_22_ in (iife99_()
                ).keys.Elements:
                    d_1114_v71_: D0 = compr_22_
                    if (d_1114_v71_) in (iife100_()
                    ):
                        coll22_[d_1114_v71_] = len((self).f29)
                return _dafny.Map(coll22_)
            rhs214_ = d_1102_v65_
            rhs215_ = (d_1107_v67_ if (self).fm7(globalState) else d_1107_v67_)
            rhs216_ = (self).f25
            rhs217_ = default__.fm5((d_1109_v69_).isdisjoint(default__.fm21(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncdfeu")), (self).f26, (self).f26, globalState)), (d_1110_v70_ if False else iife98_()
            ), ((self).f26) * (((d_1101_v64_)[(self).f26] if ((self).f26) in (d_1101_v64_) else (0) - ((self).f26))), globalState)
            rhs218_ = d_1108_v68_
            lhs186_ = d_1104_v66_
            lhs187_ = default__.safeIndex(56, (d_1104_v66_).length(0))
            d_1102_v65_ = rhs214_
            lhs186_[lhs187_] = rhs215_
            r2 = rhs216_
            r2 = rhs217_
            d_1108_v68_ = rhs218_
            d_1115_v75_: _dafny.Array
            nw194_ = _dafny.Array(_dafny.Set({}), 12)
            d_1115_v75_ = nw194_
            d_1115_v75_ = d_1115_v75_
            d_1116_v76_: _dafny.Array
            nw195_ = _dafny.Array(False, 22)
            d_1116_v76_ = nw195_
            index211_ = default__.safeIndex(734, (d_1116_v76_).length(0))
            (d_1116_v76_)[index211_] = (self).f25
            index212_ = default__.safeIndex(734, (d_1116_v76_).length(0))
            (d_1116_v76_)[index212_] = (default__.fm2(p0, (0) - ((0) - (len((d_1107_v67_)[default__.safeIndex(533, len(d_1107_v67_))]))), globalState)) >= ((self).f26)
            d_1117_v77_: _dafny.Set
            d_1117_v77_ = _dafny.Set({len(_dafny.Set({False, (self).f25, p1})), p0})
            d_1118_v78_: _dafny.Seq
            d_1118_v78_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({(self).f26})).isdisjoint(d_1117_v77_), True])
            d_1119_v79_: C0
            nw196_ = C0()
            nw196_.ctor__((self).f29, (d_1116_v76_)[default__.safeIndex(734, (d_1116_v76_).length(0))], (self).f26, d_1116_v76_, False)
            d_1119_v79_ = nw196_
            d_1120_v80_: _dafny.Map
            d_1120_v80_ = _dafny.Map({d_1119_v79_: d_1118_v78_})
            d_1118_v78_ = ((d_1120_v80_)[d_1119_v79_] if (d_1119_v79_) in (d_1120_v80_) else default__.fm23(globalState))
        d_1121_v81_: D11
        d_1121_v81_ = D11_DC26((self).f26, (self).f26, (self).f25)
        d_1122_v82_: _dafny.Map
        d_1122_v82_ = _dafny.Map({(d_1121_v81_).cf41: (self).fm6(not(p2), (d_1121_v81_).cf42, (0) - (p0), globalState)})
        d_1123_v83_: _dafny.Seq
        d_1123_v83_ = _dafny.SeqWithoutIsStrInference([p0, (self).f26])
        d_1122_v82_ = (d_1122_v82_).set(len((_dafny.SeqWithoutIsStrInference([(self).f26 for d_1124_i12_ in range(default__.abs(-976))])) + (d_1123_v83_)), p0)
        d_1125_v84_: _dafny.Array
        def lambda70_(d_1126_i13_):
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1127_i14_ in range(default__.abs(638))])) == ((self).f29)

        init37_ = lambda70_
        nw197_ = _dafny.Array(None, 3)
        for i0_37_ in range(nw197_.length(0)):
            nw197_[i0_37_] = init37_(i0_37_)
        d_1125_v84_ = nw197_
        index213_ = default__.safeIndex(882, (d_1125_v84_).length(0))
        (d_1125_v84_)[index213_] = p1
        index214_ = default__.safeIndex(882, (d_1125_v84_).length(0))
        (d_1125_v84_)[index214_] = not((self).f30)
        hi11_ = default__.safeModuloInt((self).f26, (0) - ((self).f26))
        for d_1128_i15_ in range(p0, hi11_):
            d_1129_v85_: _dafny.Array
            def lambda71_(d_1130_v84_):
                def lambda72_(d_1131_i16_):
                    return (_dafny.Set({False})) - (_dafny.Set({(d_1130_v84_)[default__.safeIndex(882, (d_1130_v84_).length(0))]}))

                return lambda72_

            init38_ = lambda71_(d_1125_v84_)
            nw198_ = _dafny.Array(None, 24)
            for i0_38_ in range(nw198_.length(0)):
                nw198_[i0_38_] = init38_(i0_38_)
            d_1129_v85_ = nw198_
            d_1132_v86_: _dafny.Map
            d_1132_v86_ = _dafny.Map({len(d_1123_v83_): (self).f25})
            d_1133_v87_: _dafny.Set
            d_1133_v87_ = _dafny.Set({((d_1132_v86_)[d_1128_i15_] if (d_1128_i15_) in (d_1132_v86_) else (self).f30)})
            index215_ = default__.safeIndex(58, (d_1129_v85_).length(0))
            (d_1129_v85_)[index215_] = (d_1133_v87_) | (_dafny.Set({(self).f30}))
            index216_ = default__.safeIndex(58, (d_1129_v85_).length(0))
            (d_1129_v85_)[index216_] = (d_1133_v87_).intersection(d_1133_v87_)
            d_1134_v88_: _dafny.Map
            d_1134_v88_ = _dafny.Map({p2: (d_1125_v84_)[default__.safeIndex(882, (d_1125_v84_).length(0))]})
            r2 = ((d_1134_v88_)[(self).f25] if ((self).f25) in (d_1134_v88_) else (d_1125_v84_)[default__.safeIndex(882, (d_1125_v84_).length(0))])
            d_1135_v89_: _dafny.Seq
            d_1135_v89_ = _dafny.SeqWithoutIsStrInference([(d_1125_v84_)[default__.safeIndex(882, (d_1125_v84_).length(0))]])
            d_1136_v90_: _dafny.Set
            d_1136_v90_ = _dafny.Set({p0, d_1128_i15_, (_dafny.MultiSet((d_1135_v89_).set(default__.safeIndex(p0, len(d_1135_v89_)), True))).cardinality})
            d_1137_v91_: C2
            nw199_ = C2()
            nw199_.ctor__(-927, (self).f25, len(d_1136_v90_))
            d_1137_v91_ = nw199_
            d_1138_v92_: _dafny.Seq
            d_1138_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxlotbcg"))
            rhs219_ = (self).f26
            rhs220_ = p1
            rhs221_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            lhs188_ = globalState
            lhs189_ = globalState
            lhs188_.f0 = rhs219_
            lhs189_.f1 = rhs220_
            d_1138_v92_ = rhs221_
        d_1139_v93_: C2
        nw200_ = C2()
        nw200_.ctor__((self).f26, True, p0)
        d_1139_v93_ = nw200_
        d_1140_v94_: _dafny.Map
        d_1140_v94_ = _dafny.Map({(self).f30: d_1139_v93_})
        r0 = len((d_1140_v94_) | ((d_1140_v94_) | (d_1140_v94_)))
        r1 = (self).f26
        r2 = (self).f30
        return r0, r1, r2

    @property
    def f29(self):
        return self._f29
    @property
    def f30(self):
        return self._f30

class C4(T0):
    def  __init__(self):
        self._f25: bool = False
        self._f26: int = int(0)
        self.f28: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f27: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26
    def ctor__(self, f27, f28, f25, f26):
        (self)._f27 = f27
        (self).f28 = f28
        (self)._f25 = f25
        (self)._f26 = f26

    def fm6(self, p0, p1, p2, globalState):
        return (self).f27

    def fm7(self, globalState):
        return not(not(not((D0_DC0(True)).cf0)))

    def fm8(self, p0, globalState):
        source16_ = D0_DC2()
        if source16_.is_DC1:
            d_1141___mcc_h0_ = source16_.cf1
            d_1142___mcc_h1_ = source16_.cf2
            d_1143___mcc_h2_ = source16_.cf3
            d_1144___mcc_h3_ = source16_.cf4
            d_1145___mcc_h4_ = source16_.cf5
            d_1146_cf5_ = d_1145___mcc_h4_
            d_1147_cf4_ = d_1144___mcc_h3_
            d_1148_cf3_ = d_1143___mcc_h2_
            d_1149_cf2_ = d_1142___mcc_h1_
            d_1150_cf1_ = d_1141___mcc_h0_
            return D0_DC2()
        elif source16_.is_DC2:
            return D0_DC2()
        elif True:
            d_1151___mcc_h5_ = source16_.cf0
            d_1152_cf0_ = d_1151___mcc_h5_
            return D0_DC2()

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        (globalState).f1 = (self).f25
        d_1153_v0_: _dafny.Seq
        d_1153_v0_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25])
        if (d_1153_v0_)[default__.safeIndex((self).f27, len(d_1153_v0_))]:
            d_1154_v1_: _dafny.Set
            d_1154_v1_ = _dafny.Set({self.f28})
            d_1155_v2_: D0
            d_1155_v2_ = D0_DC0(p2)
            d_1156_v3_: _dafny.Map
            d_1156_v3_ = _dafny.Map({d_1155_v2_: (self).f26})
            d_1157_v4_: _dafny.Seq
            d_1157_v4_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f27])
            d_1158_v5_: _dafny.Array
            nw201_ = _dafny.Array(None, 25)
            nw201_[int(0)] = (True) or (p2)
            nw201_[int(1)] = p2
            nw201_[int(2)] = (self).f25
            nw201_[int(3)] = p2
            nw201_[int(4)] = (self).f25
            nw201_[int(5)] = (self).f25
            nw201_[int(6)] = not(p1)
            nw201_[int(7)] = not(p1)
            nw201_[int(8)] = True
            nw201_[int(9)] = True
            nw201_[int(10)] = p1
            nw201_[int(11)] = (d_1154_v1_).issubset(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrebsokg")), self.f28}))
            nw201_[int(12)] = (d_1153_v0_)[default__.safeIndex((self).f27, len(d_1153_v0_))]
            nw201_[int(13)] = default__.fm5((d_1155_v2_).cf0, (d_1156_v3_).set(D0_DC0((self).f25), (d_1157_v4_)[default__.safeIndex((self).f27, len(d_1157_v4_))]), (self).f26, globalState)
            nw201_[int(14)] = (self).f25
            nw201_[int(15)] = (self).f25
            nw201_[int(16)] = False
            nw201_[int(17)] = not(p2)
            nw201_[int(18)] = (self).f25
            nw201_[int(19)] = p1
            nw201_[int(20)] = not (p2) or ((self).f25)
            nw201_[int(21)] = (self).f25
            nw201_[int(22)] = (self).f25
            nw201_[int(23)] = p2
            nw201_[int(24)] = p1
            d_1158_v5_ = nw201_
            index217_ = default__.safeIndex(191, (d_1158_v5_).length(0))
            (d_1158_v5_)[index217_] = (len(self.f28)) <= ((self).f27)
            index218_ = default__.safeIndex(191, (d_1158_v5_).length(0))
            (d_1158_v5_)[index218_] = p2
            if (self).f25:
                d_1159_v6_: _dafny.Array
                nw202_ = _dafny.Array(None, 1)
                nw202_[int(0)] = self.f28
                d_1159_v6_ = nw202_
                index219_ = default__.safeIndex(244, (d_1159_v6_).length(0))
                (d_1159_v6_)[index219_] = (self.f28) + (self.f28)
                index220_ = default__.safeIndex(244, (d_1159_v6_).length(0))
                (d_1159_v6_)[index220_] = default__.fm1(globalState)
                d_1160_v7_: _dafny.Map
                d_1160_v7_ = _dafny.Map({(d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))]: (self).f26})
                d_1161_v8_: D0
                d_1161_v8_ = D0_DC2()
                d_1162_v9_: _dafny.MultiSet
                d_1162_v9_ = _dafny.MultiSet([d_1161_v8_, d_1161_v8_])
                d_1163_v10_: _dafny.Seq
                d_1163_v10_ = _dafny.SeqWithoutIsStrInference([d_1161_v8_])
                d_1164_v11_: D0
                d_1164_v11_ = D0_DC1(535, (d_1159_v6_)[default__.safeIndex(244, (d_1159_v6_).length(0))], p2, (self).f26, (d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))])
                rhs222_ = (_dafny.MultiSet([d_1161_v8_, d_1161_v8_, d_1161_v8_, D0_DC2()])).issubset((d_1162_v9_).intersection(_dafny.MultiSet(d_1163_v10_)))
                rhs223_ = (self).f25
                rhs224_ = (default__.safeModuloInt((self).f26, 50)) in (_dafny.Set({(self).f26, (self).f27, (self).f27, len(_dafny.Set({False}))}))
                rhs225_ = (d_1164_v11_).cf4
                rhs226_ = (d_1160_v7_) | (d_1160_v7_)
                lhs190_ = globalState
                lhs191_ = globalState
                lhs192_ = globalState
                lhs193_ = globalState
                lhs190_.f1 = rhs222_
                lhs191_.f1 = rhs223_
                lhs192_.f1 = rhs224_
                lhs193_.f2 = rhs225_
                d_1160_v7_ = rhs226_
                index221_ = default__.safeIndex(191, (d_1158_v5_).length(0))
                (d_1158_v5_)[index221_] = (self).f25
                d_1165_v12_: str
                d_1165_v12_ = _dafny.CodePoint('b')
                d_1165_v12_ = d_1165_v12_
                d_1166_v13_: _dafny.Array
                nw203_ = _dafny.Array(_dafny.Set({}), 4)
                d_1166_v13_ = nw203_
                d_1167_v14_: _dafny.Array
                nw204_ = _dafny.Array(_dafny.CodePoint('D'), 6)
                d_1167_v14_ = nw204_
                d_1168_v15_: _dafny.Set
                d_1168_v15_ = _dafny.Set({len(d_1157_v4_), (self).f27, (self).f26, (self).f27})
                index222_ = default__.safeIndex(643, (d_1166_v13_).length(0))
                (d_1166_v13_)[index222_] = (_dafny.Set({696, len(_dafny.Map({(self).f27: d_1167_v14_}))})).intersection(d_1168_v15_)
                index223_ = default__.safeIndex(643, (d_1166_v13_).length(0))
                rhs227_ = (len(default__.fm1(globalState))) - (default__.safeModuloInt(261, (self).f26))
                rhs228_ = d_1168_v15_
                rhs229_ = d_1165_v12_
                lhs194_ = globalState
                lhs195_ = d_1166_v13_
                lhs196_ = default__.safeIndex(643, (d_1166_v13_).length(0))
                lhs194_.f0 = rhs227_
                lhs195_[lhs196_] = rhs228_
                d_1165_v12_ = rhs229_
            elif True:
                (globalState).f1 = (d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))]
                d_1169_v16_: _dafny.Set
                d_1169_v16_ = _dafny.Set({(d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))], (d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))]})
                d_1170_v17_: str
                d_1170_v17_ = _dafny.CodePoint('i')
                d_1171_v18_: _dafny.Map
                d_1171_v18_ = _dafny.Map({d_1169_v16_: d_1170_v17_})
                d_1172_v19_: _dafny.Map
                d_1172_v19_ = _dafny.Map({(d_1171_v18_).set(d_1169_v16_, d_1170_v17_): (0) - (((self).f26) - ((self).f27))})
                d_1172_v19_ = (d_1172_v19_).set((d_1171_v18_) | (d_1171_v18_), (self).f26)
                d_1155_v2_ = d_1155_v2_
                d_1173_v20_: T0
                nw205_ = C3()
                nw205_.ctor__((_dafny.SeqWithoutIsStrInference([(self.f28)[default__.safeIndex((self).f27, len(self.f28))] for d_1174_i0_ in range(default__.abs(-475))])) + (self.f28), p1, ((self).f26) != ((0) - ((self).f27)), (self).f27)
                d_1173_v20_ = nw205_
                d_1175_v21_: _dafny.Array
                nw206_ = _dafny.Array(int(0), 20)
                d_1175_v21_ = nw206_
                d_1176_v22_: _dafny.Map
                d_1176_v22_ = _dafny.Map({d_1175_v21_: (d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))]})
                d_1177_v23_: _dafny.Map
                d_1177_v23_ = _dafny.Map({p1: ((0) - (len(d_1176_v22_))) - ((d_1173_v20_).f26)})
                d_1178_v24_: D0
                d_1178_v24_ = D0_DC1(len(self.f28), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmdoif")), p2, (self).f26, (self).f25)
                rhs230_ = (self).f25
                rhs231_ = d_1173_v20_
                rhs232_ = not ((d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))]) or ((d_1173_v20_).fm7(globalState))
                rhs233_ = (_dafny.Map({(self).f25: (self).f26}) if (d_1173_v20_).f25 else d_1177_v23_)
                rhs234_ = d_1178_v24_
                lhs197_ = globalState
                lhs198_ = globalState
                lhs197_.f1 = rhs230_
                d_1173_v20_ = rhs231_
                lhs198_.f1 = rhs232_
                d_1177_v23_ = rhs233_
                d_1178_v24_ = rhs234_
                d_1179_v25_: _dafny.Array
                nw207_ = _dafny.Array(_dafny.CodePoint('D'), 25)
                d_1179_v25_ = nw207_
                d_1180_v26_: _dafny.Seq
                d_1180_v26_ = _dafny.SeqWithoutIsStrInference([d_1179_v25_, d_1179_v25_])
                (globalState).f1 = ((d_1173_v20_).f26) <= (len(d_1180_v26_))
            index224_ = default__.safeIndex(191, (d_1158_v5_).length(0))
            (d_1158_v5_)[index224_] = (d_1153_v0_)[default__.safeIndex((self).f27, len(d_1153_v0_))]
            d_1181_v27_: _dafny.Array
            def lambda73_(d_1182_i1_):
                return self.f28

            init39_ = lambda73_
            nw208_ = _dafny.Array(None, 19)
            for i0_39_ in range(nw208_.length(0)):
                nw208_[i0_39_] = init39_(i0_39_)
            d_1181_v27_ = nw208_
            index225_ = default__.safeIndex(720, (d_1181_v27_).length(0))
            (d_1181_v27_)[index225_] = (self.f28) + (self.f28)
            index226_ = default__.safeIndex(720, (d_1181_v27_).length(0))
            (d_1181_v27_)[index226_] = self.f28
            d_1183_v28_: _dafny.Set
            d_1183_v28_ = _dafny.Set({d_1153_v0_})
            if (d_1183_v28_).ispropersubset(default__.fm13((d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))], globalState)):
                index227_ = default__.safeIndex(191, (d_1158_v5_).length(0))
                (d_1158_v5_)[index227_] = p2
                d_1184_v29_: _dafny.MultiSet
                d_1184_v29_ = _dafny.MultiSet([d_1158_v5_])
                d_1185_v30_: _dafny.Map
                d_1185_v30_ = _dafny.Map({p2: (d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))]})
                d_1184_v29_ = ((d_1184_v29_) | ((d_1184_v29_).set(d_1158_v5_, default__.abs(len(d_1185_v30_))))) - ((d_1184_v29_) | (d_1184_v29_))
                d_1186_v31_: int
                d_1187_v32_: bool
                d_1188_v33_: _dafny.Map
                d_1189_v34_: _dafny.Map
                out24_: int
                out25_: bool
                out26_: _dafny.Map
                out27_: _dafny.Map
                out24_, out25_, out26_, out27_ = (self).m2(p1, globalState)
                d_1186_v31_ = out24_
                d_1187_v32_ = out25_
                d_1188_v33_ = out26_
                d_1189_v34_ = out27_
                d_1190_v35_: str
                d_1190_v35_ = _dafny.CodePoint('b')
                d_1190_v35_ = d_1190_v35_
                d_1191_v36_: _dafny.Seq
                d_1191_v36_ = _dafny.SeqWithoutIsStrInference([(d_1158_v5_ if p2 else d_1158_v5_), d_1158_v5_, d_1158_v5_])
                rhs235_ = ((self.f28) < ((d_1181_v27_)[default__.safeIndex(720, (d_1181_v27_).length(0))]) if p2 else (self).f25)
                rhs236_ = _dafny.SeqWithoutIsStrInference([d_1158_v5_, d_1158_v5_, d_1158_v5_, d_1158_v5_, d_1158_v5_])
                d_1187_v32_ = rhs235_
                d_1191_v36_ = rhs236_
            elif True:
                (self).f28 = self.f28
                d_1192_v37_: _dafny.MultiSet
                d_1192_v37_ = _dafny.MultiSet([(self).f25, p1])
                (globalState).f1 = ((d_1192_v37_) - (_dafny.MultiSet([p1, p1]))).ispropersubset((_dafny.MultiSet([(self).f25, (self).f25, (d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))], p1, p1])) | (d_1192_v37_))
                d_1193_v38_: _dafny.Map
                d_1193_v38_ = _dafny.Map({p2: (d_1181_v27_)[default__.safeIndex(720, (d_1181_v27_).length(0))]})
                d_1194_v39_: _dafny.Seq
                d_1194_v39_ = _dafny.SeqWithoutIsStrInference([(d_1181_v27_)[default__.safeIndex(720, (d_1181_v27_).length(0))], default__.fm1(globalState), self.f28, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wddepf")), ((d_1193_v38_)[(d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))]] if ((d_1158_v5_)[default__.safeIndex(191, (d_1158_v5_).length(0))]) in (d_1193_v38_) else (d_1181_v27_)[default__.safeIndex(720, (d_1181_v27_).length(0))])])
                d_1195_v40_: _dafny.Set
                d_1195_v40_ = _dafny.Set({default__.fm5((self).f25, d_1156_v3_, len(d_1153_v0_), globalState)})
                d_1196_v41_: C3
                nw209_ = C3()
                nw209_.ctor__((d_1194_v39_)[default__.safeIndex(len(d_1195_v40_), len(d_1194_v39_))], (d_1195_v40_) != (d_1195_v40_), p1, (0) - ((self).f26))
                d_1196_v41_ = nw209_
                (globalState).f1 = not (p2) or (not (True) or (p1))
                d_1197_v42_: _dafny.Map
                d_1197_v42_ = _dafny.Map({(d_1196_v41_).f30: True})
                d_1198_v43_: D11
                d_1198_v43_ = D11_DC25(d_1197_v42_)
                d_1199_v44_: str
                d_1199_v44_ = _dafny.CodePoint('j')
                d_1198_v43_ = default__.fm34(((d_1196_v41_).f29).set(default__.safeIndex((self).f26, len((d_1196_v41_).f29)), d_1199_v44_), (self).f27, globalState)
        elif True:
            d_1200_v45_: _dafny.Seq
            d_1200_v45_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f27, (self).f27, 844])
            d_1201_v46_: _dafny.Set
            d_1201_v46_ = _dafny.Set({p2})
            d_1202_v47_: _dafny.Map
            d_1202_v47_ = _dafny.Map({(self).f26: d_1201_v46_})
            (globalState).f1 = (((d_1200_v45_)[default__.safeIndex(161, len(d_1200_v45_))] if (self).f25 else (self).f27)) in (d_1202_v47_)
            r0 = p1
            d_1203_v48_: _dafny.Map
            d_1203_v48_ = _dafny.Map({(self).f27: (649) > ((self).f26)})
            d_1204_v49_: _dafny.MultiSet
            d_1204_v49_ = _dafny.MultiSet([p0])
            d_1205_v50_: D0
            d_1205_v50_ = D0_DC1((d_1204_v49_).cardinality, self.f28, p1, (self).f26, p2)
            d_1203_v48_ = (d_1203_v48_).set(default__.safeModuloInt(-915, (self).f27), (d_1205_v50_).cf3)
            d_1206_v51_: str
            d_1206_v51_ = _dafny.CodePoint('s')
            d_1207_v52_: _dafny.Map
            d_1207_v52_ = _dafny.Map({d_1206_v51_: p2})
            d_1208_v54_: _dafny.Set
            d_1208_v54_ = _dafny.Set({d_1206_v51_})
            def iife101_():
                coll25_ = _dafny.Map()
                compr_25_: str
                for compr_25_ in (d_1208_v54_).Elements:
                    d_1209_v53_: str = compr_25_
                    if (d_1209_v53_) in (d_1208_v54_):
                        coll25_[d_1209_v53_] = True
                return _dafny.Map(coll25_)
            d_1207_v52_ = (_dafny.Map({_dafny.CodePoint('h'): p2})) | (iife101_()
            )
            (self).f28 = default__.fm1(globalState)
        d_1210_v55_: _dafny.Seq
        d_1210_v55_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        if (default__.safeDivisionInt((self).f27, (self).f27)) != ((d_1210_v55_)[default__.safeIndex((d_1210_v55_)[default__.safeIndex((self).f27, len(d_1210_v55_))], len(d_1210_v55_))]):
            d_1211_v56_: D1
            d_1211_v56_ = D1_DC4(False, p1)
            source17_ = d_1211_v56_
            if source17_.is_DC4:
                d_1212___mcc_h0_ = source17_.cf7
                d_1213___mcc_h1_ = source17_.cf8
                d_1214_cf8_ = d_1213___mcc_h1_
                d_1215_cf7_ = d_1212___mcc_h0_
                d_1216_v57_: D6
                d_1216_v57_ = D6_DC14(-775, (self).f26, _dafny.MultiSet([False, not(d_1215_cf7_)]))
                d_1217_v58_: _dafny.Set
                d_1217_v58_ = _dafny.Set({(d_1216_v57_).cf23})
                d_1218_v59_: _dafny.Map
                d_1218_v59_ = _dafny.Map({(self).f27: (self).f26})
                def iife102_():
                    coll26_ = _dafny.Set()
                    compr_26_: int
                    for compr_26_ in (d_1218_v59_).keys.Elements:
                        d_1219_v60_: int = compr_26_
                        if (d_1219_v60_) in (d_1218_v59_):
                            coll26_ = coll26_.union(_dafny.Set([(d_1219_v60_) - (746)]))
                    return _dafny.Set(coll26_)
                rhs237_ = ((self.f28) + (self.f28)) + (self.f28)
                rhs238_ = (len(d_1217_v58_)) > ((self).f26)
                rhs239_ = ((iife102_()
                ) - (d_1217_v58_)).intersection(_dafny.Set({(self).f27}))
                lhs199_ = self
                lhs199_.f28 = rhs237_
                d_1214_cf8_ = rhs238_
                d_1217_v58_ = rhs239_
                d_1153_v0_ = d_1153_v0_
                d_1220_v61_: D2
                d_1220_v61_ = D2_DC7(d_1215_cf7_, d_1214_cf8_, (self).f27)
                d_1221_v62_: _dafny.Array
                nw210_ = _dafny.Array(None, 20)
                nw210_[int(0)] = len(self.f28)
                nw210_[int(1)] = (self).f26
                nw210_[int(2)] = ((d_1210_v55_)[default__.safeIndex(-480, len(d_1210_v55_))]) - ((self).f27)
                nw210_[int(3)] = default__.safeDivisionInt(901, (p0).cardinality)
                nw210_[int(4)] = (self).f26
                nw210_[int(5)] = default__.fm2((self).f27, (p0).cardinality, globalState)
                nw210_[int(6)] = ((0) - (len(d_1217_v58_))) - ((self).f26)
                nw210_[int(7)] = (self).f26
                nw210_[int(8)] = default__.fm2((self).f26, default__.fm2((self).f27, len(d_1210_v55_), globalState), globalState)
                nw210_[int(9)] = len(d_1153_v0_)
                nw210_[int(10)] = ((self).f27) - ((d_1220_v61_).cf13)
                nw210_[int(11)] = (self).f26
                nw210_[int(12)] = (self).f26
                nw210_[int(13)] = (d_1216_v57_).cf24
                nw210_[int(14)] = (self).f27
                nw210_[int(15)] = ((self).f26) * ((self).f26)
                nw210_[int(16)] = (0) - (default__.safeModuloInt((self).f27, (self).f27))
                nw210_[int(17)] = (self).f27
                nw210_[int(18)] = (self).f27
                nw210_[int(19)] = (self).f27
                d_1221_v62_ = nw210_
                nw211_ = _dafny.Array(int(0), 5)
                d_1221_v62_ = nw211_
                d_1222_v63_: _dafny.Map
                d_1222_v63_ = _dafny.Map({d_1211_v56_: (self).f27})
                index228_ = default__.safeIndex(602, (d_1221_v62_).length(0))
                (d_1221_v62_)[index228_] = (0) - ((self).f26)
                index229_ = default__.safeIndex(602, (d_1221_v62_).length(0))
                rhs240_ = _dafny.Map({D1_DC4(p1, True): (8) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1223_i2_ in range(default__.abs(-569))])))})
                rhs241_ = default__.safeDivisionInt((0) - ((self).f27), (self).f26)
                lhs200_ = d_1221_v62_
                lhs201_ = default__.safeIndex(602, (d_1221_v62_).length(0))
                d_1222_v63_ = rhs240_
                lhs200_[lhs201_] = rhs241_
            elif source17_.is_DC3:
                d_1224___mcc_h2_ = source17_.cf6
                d_1225_cf6_ = d_1224___mcc_h2_
                d_1226_v64_: D2
                d_1226_v64_ = D2_DC6(_dafny.Set({(self).f26}))
                d_1226_v64_ = d_1226_v64_
                d_1227_v65_: _dafny.Map
                d_1227_v65_ = _dafny.Map({(self).f27: (self).f27})
                d_1228_v66_: _dafny.Map
                d_1228_v66_ = _dafny.Map({p2: (self).f27})
                d_1229_v67_: D0
                d_1229_v67_ = D0_DC1((self).f27, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssmnmk")), False, (self).f27, (self).f25)
                d_1230_v68_: _dafny.Array
                nw212_ = _dafny.Array(None, 27)
                nw212_[int(0)] = ((d_1227_v65_)[(self).f27] if ((self).f27) in (d_1227_v65_) else (self).f27)
                nw212_[int(1)] = default__.safeModuloInt(len((d_1228_v66_).set(p1, (self).f26)), (self).f27)
                nw212_[int(2)] = (d_1229_v67_).cf4
                nw212_[int(3)] = ((self).f26) * ((self).f26)
                nw212_[int(4)] = ((d_1228_v66_)[not(p2)] if (not(p2)) in (d_1228_v66_) else -159)
                nw212_[int(5)] = (_dafny.MultiSet(self.f28)).cardinality
                nw212_[int(6)] = len(_dafny.SeqWithoutIsStrInference([p2]))
                nw212_[int(7)] = default__.safeDivisionInt((self).f27, (self).f27)
                nw212_[int(8)] = (self).f27
                nw212_[int(9)] = (self).f26
                nw212_[int(10)] = (self).f26
                nw212_[int(11)] = 619
                nw212_[int(12)] = default__.fm2((self).f26, 731, globalState)
                nw212_[int(13)] = (0) - (-263)
                nw212_[int(14)] = ((self).f26) * ((self).f26)
                nw212_[int(15)] = (self).f26
                nw212_[int(16)] = (self).f27
                nw212_[int(17)] = default__.safeModuloInt((self).f27, (self).f26)
                nw212_[int(18)] = (self).f26
                nw212_[int(19)] = (self).f27
                nw212_[int(20)] = ((self).f26) * ((self).f26)
                nw212_[int(21)] = (self).f26
                nw212_[int(22)] = ((p0)[(self).f26] if ((self).f26) in (p0) else (self).f26)
                nw212_[int(23)] = (self).f26
                nw212_[int(24)] = (self).f27
                nw212_[int(25)] = (self).f26
                nw212_[int(26)] = len(self.f28)
                d_1230_v68_ = nw212_
                nw213_ = _dafny.Array(int(0), 5)
                d_1230_v68_ = nw213_
                (globalState).f0 = (self).f26
                (globalState).f1 = (p1 if not (p2) or (p2) else not ((self).f25) or (p1))
            elif True:
                d_1231___mcc_h3_ = source17_.cf9
                d_1232_cf9_ = d_1231___mcc_h3_
                d_1233_v69_: _dafny.Array
                nw214_ = _dafny.Array(int(0), 12)
                d_1233_v69_ = nw214_
                def lambda74_(d_1234_i3_):
                    return default__.safeModuloInt(d_1234_i3_, (self).f26)

                init40_ = lambda74_
                nw215_ = _dafny.Array(None, 8)
                for i0_40_ in range(nw215_.length(0)):
                    nw215_[i0_40_] = init40_(i0_40_)
                d_1233_v69_ = nw215_
                d_1235_v70_: D0
                d_1235_v70_ = D0_DC1(len(self.f28), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dl")), True, (self).f26, not(False))
                d_1236_v71_: str
                d_1236_v71_ = _dafny.CodePoint('q')
                (self).f28 = (((d_1235_v70_).cf2) + ((self.f28 if (self).f25 else self.f28))).set(default__.safeIndex((len(_dafny.Map({p2: p2}))) + ((self).f26), len(((d_1235_v70_).cf2) + ((self.f28 if (self).f25 else self.f28)))), d_1236_v71_)
                d_1211_v56_ = D1_DC4(((self).f26) > ((self).f26), (self).f25)
                d_1237_v72_: _dafny.Array
                nw216_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_1237_v72_ = nw216_
                d_1238_v73_: _dafny.Array
                def lambda75_(d_1239_p1_):
                    def lambda76_(d_1240_i4_):
                        return _dafny.Map({not((self).f25): d_1239_p1_})

                    return lambda76_

                init41_ = lambda75_(p1)
                nw217_ = _dafny.Array(None, 4)
                for i0_41_ in range(nw217_.length(0)):
                    nw217_[i0_41_] = init41_(i0_41_)
                d_1238_v73_ = nw217_
                d_1241_v74_: D8
                d_1241_v74_ = D8_DC16(d_1238_v73_)
                d_1242_v75_: _dafny.Map
                d_1242_v75_ = _dafny.Map({p2: p1})
                d_1243_v76_: D12
                d_1243_v76_ = D12_DC28(d_1233_v69_)
                rhs242_ = 576
                rhs243_ = ((d_1243_v76_).cf44 if not(default__.fm5((self).f25, default__.fm25((self).f26, (self).f25, (self).f25, globalState), len(d_1242_v75_), globalState)) else d_1233_v69_)
                rhs244_ = d_1237_v72_
                rhs245_ = d_1241_v74_
                rhs246_ = (self).f25
                lhs202_ = globalState
                lhs203_ = globalState
                lhs202_.f2 = rhs242_
                d_1233_v69_ = rhs243_
                d_1237_v72_ = rhs244_
                d_1241_v74_ = rhs245_
                lhs203_.f1 = rhs246_
            d_1244_v77_: int
            out28_: int
            out28_ = default__.m0((True if (self).f25 else (self).f25), globalState)
            d_1244_v77_ = out28_
            d_1245_v78_: _dafny.Array
            nw218_ = _dafny.Array(None, 19)
            nw218_[int(0)] = (self).f25
            nw218_[int(1)] = (self).f25
            nw218_[int(2)] = p1
            nw218_[int(3)] = p1
            nw218_[int(4)] = p2
            nw218_[int(5)] = not((self).f25)
            nw218_[int(6)] = (self).fm7(globalState)
            nw218_[int(7)] = (self).f25
            nw218_[int(8)] = p2
            nw218_[int(9)] = (self).f25
            nw218_[int(10)] = p2
            nw218_[int(11)] = p2
            nw218_[int(12)] = (self).f25
            nw218_[int(13)] = p2
            nw218_[int(14)] = p1
            nw218_[int(15)] = p2
            nw218_[int(16)] = False
            nw218_[int(17)] = (self).f25
            nw218_[int(18)] = p2
            d_1245_v78_ = nw218_
            d_1246_v79_: str
            d_1246_v79_ = _dafny.CodePoint('p')
            d_1247_v80_: _dafny.Map
            d_1247_v80_ = _dafny.Map({d_1245_v78_: d_1246_v79_})
            d_1248_v81_: _dafny.Set
            d_1248_v81_ = _dafny.Set({d_1244_v77_, d_1244_v77_, (self).f27, len(d_1247_v80_), len(self.f28)})
            d_1249_v82_: _dafny.MultiSet
            d_1249_v82_ = _dafny.MultiSet([d_1248_v81_])
            d_1249_v82_ = (d_1249_v82_) - ((d_1249_v82_ if (self).f25 else d_1249_v82_))
            d_1250_v83_: _dafny.Array
            nw219_ = _dafny.Array(_dafny.Seq({}), 21)
            d_1250_v83_ = nw219_
            d_1250_v83_ = d_1250_v83_
            d_1251_v84_: C1
            nw220_ = C1()
            nw220_.ctor__((self).f25)
            d_1251_v84_ = nw220_
            d_1252_v85_: _dafny.Map
            d_1252_v85_ = _dafny.Map({p1: (d_1251_v84_ if False else d_1251_v84_)})
            d_1252_v85_ = (d_1252_v85_).set(not(p2), d_1251_v84_)
        elif True:
            if (self).f25:
                d_1253_v86_: str
                d_1253_v86_ = _dafny.CodePoint('x')
                d_1254_v87_: D6
                d_1254_v87_ = D6_DC14((self).f26, (self).f26, _dafny.MultiSet([p2, (self).f25]))
                d_1255_v88_: D9
                d_1255_v88_ = D9_DC19(d_1254_v87_)
                d_1256_v89_: D9
                d_1256_v89_ = D9_DC20((d_1255_v88_ if p1 else d_1255_v88_))
                d_1257_v90_: _dafny.Set
                d_1257_v90_ = _dafny.Set({not((d_1153_v0_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([self.f28 for d_1258_i5_ in range(default__.abs(760))])), len(d_1153_v0_))])})
                pat_let_tv33_ = d_1255_v88_
                def iife103_(_pat_let38_0):
                    def iife104_(d_1259_dt__update__tmp_h0_):
                        def iife105_(_pat_let39_0):
                            def iife106_(d_1260_dt__update_hcf32_h0_):
                                return D9_DC20(d_1260_dt__update_hcf32_h0_)
                            return iife106_(_pat_let39_0)
                        return iife105_(D9_DC20(pat_let_tv33_))
                    return iife104_(_pat_let38_0)
                rhs247_ = d_1253_v86_
                rhs248_ = not (p2) or (not((d_1257_v90_) == (d_1257_v90_)))
                rhs249_ = iife103_(d_1256_v89_)
                rhs250_ = (len((self.f28 if (self).f25 else _dafny.SeqWithoutIsStrInference([d_1253_v86_ for d_1261_i6_ in range(default__.abs(110))])))) < ((len(d_1210_v55_)) * ((self).f26))
                lhs204_ = globalState
                d_1253_v86_ = rhs247_
                r0 = rhs248_
                d_1256_v89_ = rhs249_
                lhs204_.f1 = rhs250_
                d_1262_v91_: C3
                nw221_ = C3()
                nw221_.ctor__(self.f28, p2, p2, (self).f26)
                d_1262_v91_ = nw221_
                d_1263_v92_: _dafny.Seq
                d_1263_v92_ = _dafny.SeqWithoutIsStrInference([default__.fm1(globalState)])
                d_1263_v92_ = _dafny.SeqWithoutIsStrInference([self.f28, _dafny.SeqWithoutIsStrInference([d_1253_v86_ for d_1264_i7_ in range(default__.abs(150))]), self.f28, self.f28])
                d_1265_v93_: _dafny.Array
                nw222_ = _dafny.Array(_dafny.CodePoint('D'), 16)
                d_1265_v93_ = nw222_
                d_1266_v94_: D10
                d_1266_v94_ = D10_DC22(d_1265_v93_, self.f28)
                d_1267_v95_: _dafny.Map
                d_1267_v95_ = _dafny.Map({self.f28: (d_1266_v94_).cf35})
                (globalState).f1 = (((d_1267_v95_)[self.f28] if (self.f28) in (d_1267_v95_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rj")))) <= (self.f28)
                d_1268_v96_: _dafny.Array
                nw223_ = _dafny.Array(int(0), 21)
                d_1268_v96_ = nw223_
                index230_ = default__.safeIndex(211, (d_1268_v96_).length(0))
                (d_1268_v96_)[index230_] = (self).f27
                index231_ = default__.safeIndex(211, (d_1268_v96_).length(0))
                (d_1268_v96_)[index231_] = (self).f27
            elif True:
                (globalState).f1 = p2
                (globalState).f1 = p1
                d_1269_v97_: _dafny.Map
                d_1269_v97_ = _dafny.Map({not(p2): False})
                d_1270_v98_: C1
                nw224_ = C1()
                nw224_.ctor__(p2)
                d_1270_v98_ = nw224_
                d_1271_v99_: str
                d_1271_v99_ = _dafny.CodePoint('i')
                d_1272_v100_: D10
                d_1272_v100_ = D10_DC24((self).f27, d_1270_v98_, d_1271_v99_)
                d_1273_v101_: _dafny.Map
                d_1273_v101_ = _dafny.Map({d_1272_v100_: _dafny.MultiSet([(self).f27])})
                d_1269_v97_ = (d_1269_v97_).set(not(False), ((_dafny.MultiSet([p2])).cardinality) == ((((d_1273_v101_)[d_1272_v100_] if (d_1272_v100_) in (d_1273_v101_) else p0)).cardinality))
                d_1274_v102_: C1
                nw225_ = C1()
                nw225_.ctor__(p1)
                d_1274_v102_ = nw225_
                (globalState).f1 = (self).f25
            d_1275_v103_: _dafny.Array
            nw226_ = _dafny.Array(False, 13)
            d_1275_v103_ = nw226_
            index232_ = default__.safeIndex(58, (d_1275_v103_).length(0))
            (d_1275_v103_)[index232_] = (self).f25
            index233_ = default__.safeIndex(58, (d_1275_v103_).length(0))
            (d_1275_v103_)[index233_] = True
            d_1276_v104_: _dafny.Array
            nw227_ = _dafny.Array(None, 15)
            nw227_[int(0)] = d_1275_v103_
            nw227_[int(1)] = d_1275_v103_
            nw227_[int(2)] = d_1275_v103_
            nw227_[int(3)] = d_1275_v103_
            nw227_[int(4)] = d_1275_v103_
            nw227_[int(5)] = d_1275_v103_
            nw227_[int(6)] = d_1275_v103_
            nw227_[int(7)] = d_1275_v103_
            nw227_[int(8)] = d_1275_v103_
            nw227_[int(9)] = d_1275_v103_
            nw227_[int(10)] = d_1275_v103_
            nw227_[int(11)] = d_1275_v103_
            nw227_[int(12)] = (d_1275_v103_ if p2 else d_1275_v103_)
            nw227_[int(13)] = d_1275_v103_
            nw227_[int(14)] = d_1275_v103_
            d_1276_v104_ = nw227_
            index234_ = default__.safeIndex(103, (d_1276_v104_).length(0))
            (d_1276_v104_)[index234_] = d_1275_v103_
            d_1277_v105_: _dafny.Seq
            d_1277_v105_ = _dafny.SeqWithoutIsStrInference([d_1275_v103_, d_1275_v103_, d_1275_v103_, d_1275_v103_])
            index235_ = default__.safeIndex(103, (d_1276_v104_).length(0))
            (d_1276_v104_)[index235_] = (d_1277_v105_)[default__.safeIndex(952, len(d_1277_v105_))]
            d_1278_v106_: C1
            nw228_ = C1()
            nw228_.ctor__(((d_1275_v103_)[default__.safeIndex(58, (d_1275_v103_).length(0))]) and (p2))
            d_1278_v106_ = nw228_
            if (default__.safeModuloInt((self).f27, (self).f26)) >= ((self).f27):
                d_1279_v107_: _dafny.Map
                d_1279_v107_ = _dafny.Map({(self).f26: (d_1278_v106_).fm14(_dafny.Set({p2}), (self).f26, globalState)})
                d_1280_v108_: _dafny.Set
                d_1280_v108_ = _dafny.Set({len(d_1279_v107_), (self).f26})
                def iife107_():
                    coll27_ = _dafny.Set()
                    compr_27_: int
                    for compr_27_ in _dafny.IntegerRange(-465, -266):
                        d_1281_v109_: int = compr_27_
                        if ((-465) <= (d_1281_v109_)) and ((d_1281_v109_) < (-266)):
                            coll27_ = coll27_.union(_dafny.Set([(d_1281_v109_) * ((self).f27)]))
                    return _dafny.Set(coll27_)
                d_1280_v108_ = (d_1280_v108_) - (iife107_()
                )
                d_1282_v110_: str
                d_1282_v110_ = _dafny.CodePoint('m')
                d_1283_v111_: _dafny.MultiSet
                d_1283_v111_ = _dafny.MultiSet([d_1282_v110_])
                d_1284_v112_: _dafny.Map
                d_1284_v112_ = _dafny.Map({(d_1278_v106_).f32: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qtwjddafo")))})
                d_1285_v113_: _dafny.Seq
                d_1285_v113_ = _dafny.SeqWithoutIsStrInference([d_1283_v111_])
                d_1286_v114_: _dafny.Array
                nw229_ = _dafny.Array(None, 22)
                nw229_[int(0)] = d_1283_v111_
                nw229_[int(1)] = d_1283_v111_
                nw229_[int(2)] = d_1283_v111_
                nw229_[int(3)] = (d_1283_v111_).intersection(d_1283_v111_)
                nw229_[int(4)] = d_1283_v111_
                nw229_[int(5)] = d_1283_v111_
                nw229_[int(6)] = (d_1283_v111_ if p1 else default__.fm35((self).f27, len(d_1284_v112_), globalState))
                nw229_[int(7)] = d_1283_v111_
                nw229_[int(8)] = (d_1283_v111_) - (_dafny.MultiSet([d_1282_v110_]))
                nw229_[int(9)] = (d_1283_v111_) - (_dafny.MultiSet([d_1282_v110_]))
                nw229_[int(10)] = d_1283_v111_
                nw229_[int(11)] = _dafny.MultiSet([d_1282_v110_])
                nw229_[int(12)] = (d_1283_v111_ if (d_1275_v103_)[default__.safeIndex(58, (d_1275_v103_).length(0))] else _dafny.MultiSet([d_1282_v110_, d_1282_v110_, d_1282_v110_]))
                nw229_[int(13)] = d_1283_v111_
                nw229_[int(14)] = (d_1283_v111_).set(d_1282_v110_, default__.abs((0) - ((self).f26)))
                nw229_[int(15)] = d_1283_v111_
                nw229_[int(16)] = _dafny.MultiSet([d_1282_v110_, (self.f28)[default__.safeIndex(len(d_1153_v0_), len(self.f28))]])
                nw229_[int(17)] = (d_1285_v113_)[default__.safeIndex(((_dafny.MultiSet([(d_1278_v106_).f32, True])).set((self).f25, default__.abs((self).f27))).cardinality, len(d_1285_v113_))]
                nw229_[int(18)] = d_1283_v111_
                nw229_[int(19)] = d_1283_v111_
                nw229_[int(20)] = (d_1283_v111_) - (_dafny.MultiSet([d_1282_v110_, d_1282_v110_]))
                nw229_[int(21)] = d_1283_v111_
                d_1286_v114_ = nw229_
                index236_ = default__.safeIndex(135, (d_1286_v114_).length(0))
                (d_1286_v114_)[index236_] = d_1283_v111_
                index237_ = default__.safeIndex(135, (d_1286_v114_).length(0))
                (d_1286_v114_)[index237_] = d_1283_v111_
                (globalState).f0 = (self).f27
                d_1287_v115_: _dafny.Map
                d_1287_v115_ = _dafny.Map({(d_1276_v104_)[default__.safeIndex(103, (d_1276_v104_).length(0))]: (self).f26})
                d_1288_v116_: int
                d_1289_v117_: bool
                d_1290_v118_: _dafny.Map
                d_1291_v119_: _dafny.Map
                out29_: int
                out30_: bool
                out31_: _dafny.Map
                out32_: _dafny.Map
                out29_, out30_, out31_, out32_ = (self).m2(((self).f26) == (len(d_1287_v115_)), globalState)
                d_1288_v116_ = out29_
                d_1289_v117_ = out30_
                d_1290_v118_ = out31_
                d_1291_v119_ = out32_
                r0 = p1
            elif True:
                d_1292_v120_: _dafny.Set
                d_1292_v120_ = _dafny.Set({(self).f27})
                d_1293_v121_: _dafny.Array
                nw230_ = _dafny.Array(None, 10)
                nw230_[int(0)] = default__.fm21(self.f28, (self).f27, len(d_1292_v120_), globalState)
                nw230_[int(1)] = p0
                nw230_[int(2)] = (p0).set((self).f26, default__.abs(len(self.f28)))
                nw230_[int(3)] = p0
                nw230_[int(4)] = _dafny.MultiSet([(self).f26])
                nw230_[int(5)] = p0
                nw230_[int(6)] = _dafny.MultiSet(d_1210_v55_)
                nw230_[int(7)] = p0
                nw230_[int(8)] = p0
                nw230_[int(9)] = p0
                d_1293_v121_ = nw230_
                d_1294_v122_: D10
                d_1294_v122_ = D10_DC21(d_1293_v121_)
                index238_ = default__.safeIndex(58, (d_1275_v103_).length(0))
                rhs251_ = D10_DC21(d_1293_v121_)
                rhs252_ = (self).f25
                lhs205_ = d_1275_v103_
                lhs206_ = default__.safeIndex(58, (d_1275_v103_).length(0))
                d_1294_v122_ = rhs251_
                lhs205_[lhs206_] = rhs252_
                d_1295_v123_: _dafny.Set
                d_1295_v123_ = _dafny.Set({(self).f25})
                d_1296_v124_: _dafny.Seq
                d_1296_v124_ = _dafny.SeqWithoutIsStrInference([d_1295_v123_])
                d_1297_v125_: _dafny.MultiSet
                d_1297_v125_ = _dafny.MultiSet([(d_1278_v106_).f32, (self).f25, (d_1278_v106_).f32, (d_1278_v106_).fm14(d_1295_v123_, (self).f27, globalState), (d_1275_v103_)[default__.safeIndex(58, (d_1275_v103_).length(0))]])
                rhs253_ = d_1296_v124_
                rhs254_ = p1
                rhs255_ = (False if (d_1297_v125_) == (_dafny.MultiSet([(d_1278_v106_).f32, not(p1)])) else p1)
                rhs256_ = (0) - ((self).f26)
                lhs207_ = globalState
                lhs208_ = globalState
                lhs209_ = globalState
                d_1296_v124_ = rhs253_
                lhs207_.f1 = rhs254_
                lhs208_.f1 = rhs255_
                lhs209_.f0 = rhs256_
                d_1298_v126_: _dafny.Array
                nw231_ = _dafny.Array(_dafny.Seq({}), 10)
                d_1298_v126_ = nw231_
                index239_ = default__.safeIndex(572, (d_1298_v126_).length(0))
                (d_1298_v126_)[index239_] = d_1153_v0_
                d_1299_v127_: _dafny.Seq
                d_1299_v127_ = d_1153_v0_
                index240_ = default__.safeIndex(572, (d_1298_v126_).length(0))
                rhs257_ = d_1299_v127_
                rhs258_ = len((self.f28) + (self.f28))
                rhs259_ = ((_dafny.MultiSet([(d_1278_v106_).f32, not((d_1278_v106_).f32)])) | (default__.fm27(len((_dafny.Map({(d_1276_v104_)[default__.safeIndex(103, (d_1276_v104_).length(0))]: (d_1275_v103_)[default__.safeIndex(58, (d_1275_v103_).length(0))]})).set((d_1276_v104_)[default__.safeIndex(103, (d_1276_v104_).length(0))], True)), globalState)) if p2 else _dafny.MultiSet(d_1153_v0_))
                lhs210_ = d_1298_v126_
                lhs211_ = default__.safeIndex(572, (d_1298_v126_).length(0))
                lhs212_ = globalState
                lhs210_[lhs211_] = rhs257_
                lhs212_.f0 = rhs258_
                d_1297_v125_ = rhs259_
                (globalState).f2 = (self).f27
                d_1300_v129_: _dafny.Seq
                def iife108_():
                    coll28_ = _dafny.Set()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(141, 841):
                        d_1301_v128_: int = compr_28_
                        if ((141) <= (d_1301_v128_)) and ((d_1301_v128_) < (841)):
                            coll28_ = coll28_.union(_dafny.Set([default__.safeDivisionInt(d_1301_v128_, (self).f27)]))
                    return _dafny.Set(coll28_)
                d_1300_v129_ = _dafny.SeqWithoutIsStrInference([(d_1292_v120_) | (d_1292_v120_), (d_1292_v120_) - (iife108_()
                ), d_1292_v120_])
                rhs260_ = len((d_1300_v129_)[default__.safeIndex((self).f26, len(d_1300_v129_))])
                rhs261_ = (d_1278_v106_).f32
                lhs213_ = globalState
                lhs214_ = globalState
                lhs213_.f0 = rhs260_
                lhs214_.f1 = rhs261_
        d_1302_v130_: _dafny.Array
        nw232_ = _dafny.Array(False, 17)
        d_1302_v130_ = nw232_
        index241_ = default__.safeIndex(894, (d_1302_v130_).length(0))
        (d_1302_v130_)[index241_] = p2
        index242_ = default__.safeIndex(894, (d_1302_v130_).length(0))
        (d_1302_v130_)[index242_] = p1
        d_1303_v131_: D1
        d_1303_v131_ = D1_DC4(p2, (d_1302_v130_)[default__.safeIndex(894, (d_1302_v130_).length(0))])
        d_1304_v132_: D1
        d_1304_v132_ = D1_DC5(d_1303_v131_)
        def lambda77_(source18_):
            if source18_.is_DC4:
                d_1305___mcc_h4_ = source18_.cf7
                d_1306___mcc_h5_ = source18_.cf8
                d_1307_cf8_ = d_1306___mcc_h5_
                d_1308_cf7_ = d_1305___mcc_h4_
                return len((_dafny.Set({len(self.f28), (self).f27})).intersection(_dafny.Set({(self).f26})))
            elif source18_.is_DC3:
                d_1309___mcc_h6_ = source18_.cf6
                d_1310_cf6_ = d_1309___mcc_h6_
                return (self).f27
            elif True:
                d_1311___mcc_h7_ = source18_.cf9
                d_1312_cf9_ = d_1311___mcc_h7_
                return (self).f27

        (globalState).f0 = lambda77_(d_1304_v132_)
        source19_ = d_1153_v0_
        d_1313___mcc_h8_ = source19_
        d_1314_cf26_ = d_1313___mcc_h8_
        (globalState).f0 = (self).f27
        (globalState).f1 = ((self).f26) >= ((self).f26)
        d_1315_v133_: str
        d_1315_v133_ = _dafny.CodePoint('q')
        d_1315_v133_ = (d_1315_v133_ if p2 else d_1315_v133_)
        (globalState).f0 = (self).f27
        r0 = (p2 if (d_1302_v130_)[default__.safeIndex(894, (d_1302_v130_).length(0))] else (d_1302_v130_)[default__.safeIndex(894, (d_1302_v130_).length(0))])
        return r0

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        r3: _dafny.Map = _dafny.Map({})
        d_1316_v0_: _dafny.Seq
        d_1316_v0_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25, p0, p0, not((self).f25)])
        d_1317_v1_: _dafny.Seq
        d_1317_v1_ = _dafny.SeqWithoutIsStrInference([d_1316_v0_])
        hi12_ = ((self).f27) - ((0) - (len(d_1317_v1_)))
        for d_1318_i0_ in range((self).f27, hi12_):
            d_1319_v2_: str
            d_1319_v2_ = _dafny.CodePoint('s')
            d_1320_v3_: _dafny.Map
            d_1320_v3_ = _dafny.Map({d_1319_v2_: (self).f25})
            d_1321_v4_: _dafny.Array
            nw233_ = _dafny.Array(None, 16)
            nw233_[int(0)] = (self).f25
            nw233_[int(1)] = p0
            nw233_[int(2)] = False
            nw233_[int(3)] = not(p0)
            nw233_[int(4)] = p0
            nw233_[int(5)] = (self).f25
            nw233_[int(6)] = p0
            nw233_[int(7)] = p0
            nw233_[int(8)] = (self).f25
            nw233_[int(9)] = p0
            nw233_[int(10)] = (self).f25
            nw233_[int(11)] = (self).f25
            nw233_[int(12)] = ((d_1320_v3_)[d_1319_v2_] if (d_1319_v2_) in (d_1320_v3_) else False)
            nw233_[int(13)] = p0
            nw233_[int(14)] = (self).f25
            nw233_[int(15)] = (self).f25
            d_1321_v4_ = nw233_
            d_1322_v5_: _dafny.Map
            d_1322_v5_ = _dafny.Map({self.f28: d_1321_v4_})
            d_1323_v6_: _dafny.Map
            d_1323_v6_ = _dafny.Map({((d_1322_v5_)[self.f28] if (self.f28) in (d_1322_v5_) else d_1321_v4_): p0})
            d_1323_v6_ = (d_1323_v6_).set(d_1321_v4_, True)
            rhs262_ = (self).f25
            rhs263_ = True
            lhs215_ = globalState
            lhs216_ = globalState
            lhs215_.f1 = rhs262_
            lhs216_.f1 = rhs263_
            d_1324_v7_: _dafny.Array
            nw234_ = _dafny.Array(_dafny.Seq({}), 5)
            d_1324_v7_ = nw234_
            d_1325_v8_: _dafny.Array
            nw235_ = _dafny.Array(int(0), 16)
            d_1325_v8_ = nw235_
            d_1326_v9_: _dafny.Seq
            d_1326_v9_ = _dafny.SeqWithoutIsStrInference([d_1325_v8_])
            index243_ = default__.safeIndex(875, (d_1324_v7_).length(0))
            (d_1324_v7_)[index243_] = d_1326_v9_
            index244_ = default__.safeIndex(578, (d_1321_v4_).length(0))
            (d_1321_v4_)[index244_] = (self).f25
            d_1327_v10_: D0
            d_1327_v10_ = D0_DC1(d_1318_i0_, self.f28, (self).f25, (self).f26, (self).f25)
            index245_ = default__.safeIndex(875, (d_1324_v7_).length(0))
            index246_ = default__.safeIndex(578, (d_1321_v4_).length(0))
            rhs264_ = True
            rhs265_ = (d_1326_v9_) + (_dafny.SeqWithoutIsStrInference([d_1325_v8_, d_1325_v8_, d_1325_v8_]))
            rhs266_ = p0
            rhs267_ = default__.fm4(d_1327_v10_, (self).f25, default__.safeDivisionInt((self).f26, (self).f27), globalState)
            rhs268_ = (d_1316_v0_)[default__.safeIndex(((self).f27) - ((self).f27), len(d_1316_v0_))]
            lhs217_ = d_1324_v7_
            lhs218_ = default__.safeIndex(875, (d_1324_v7_).length(0))
            lhs219_ = d_1321_v4_
            lhs220_ = default__.safeIndex(578, (d_1321_v4_).length(0))
            lhs221_ = globalState
            r1 = rhs264_
            lhs217_[lhs218_] = rhs265_
            lhs219_[lhs220_] = rhs266_
            d_1319_v2_ = rhs267_
            lhs221_.f1 = rhs268_
            d_1321_v4_ = d_1321_v4_
        d_1328_v11_: str
        d_1328_v11_ = _dafny.CodePoint('t')
        d_1329_v12_: _dafny.Map
        d_1329_v12_ = _dafny.Map({d_1328_v11_: p0})
        d_1330_v13_: _dafny.Map
        d_1330_v13_ = _dafny.Map({d_1329_v12_: (self).f26})
        d_1330_v13_ = d_1330_v13_
        d_1331_v14_: D1
        d_1331_v14_ = D1_DC4((d_1316_v0_)[default__.safeIndex(default__.fm2((self).f26, (self).f26, globalState), len(d_1316_v0_))], p0)
        d_1332_v15_: _dafny.Seq
        d_1332_v15_ = _dafny.SeqWithoutIsStrInference([d_1331_v14_])
        d_1333_v16_: D0
        d_1333_v16_ = D0_DC2()
        d_1334_v17_: D8
        d_1334_v17_ = D8_DC17(d_1333_v16_, p0)
        pat_let_tv34_ = p0
        def iife109_(_pat_let40_0):
            def iife110_(d_1335_dt__update__tmp_h0_):
                def iife111_(_pat_let41_0):
                    def iife112_(d_1336_dt__update_hcf7_h0_):
                        return D1_DC4(d_1336_dt__update_hcf7_h0_, (d_1335_dt__update__tmp_h0_).cf8)
                    return iife112_(_pat_let41_0)
                return iife111_(pat_let_tv34_)
            return iife110_(_pat_let40_0)
        d_1332_v15_ = _dafny.SeqWithoutIsStrInference([(d_1331_v14_ if (d_1334_v17_).cf29 else iife109_(d_1331_v14_))])
        d_1337_v18_: _dafny.MultiSet
        d_1337_v18_ = _dafny.MultiSet([(self).f26, len(self.f28)])
        d_1338_v19_: _dafny.Set
        d_1338_v19_ = _dafny.Set({p0})
        d_1339_v20_: _dafny.Map
        d_1339_v20_ = _dafny.Map({d_1337_v18_: (d_1338_v19_).isdisjoint(d_1338_v19_)})
        def iife113_():
            coll29_ = _dafny.Map()
            compr_29_: _dafny.MultiSet
            for compr_29_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f27])), d_1337_v18_])).Elements:
                d_1340_v21_: _dafny.MultiSet = compr_29_
                if (d_1340_v21_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f27])), d_1337_v18_])):
                    coll29_[d_1340_v21_] = p0
            return _dafny.Map(coll29_)
        d_1339_v20_ = iife113_()
        
        r0 = (self).f26
        r1 = p0
        r0 = (0) - (default__.fm2((self).f27, len((self.f28 if p0 else self.f28)), globalState))
        r1 = (self).f25
        d_1341_v22_: _dafny.MultiSet
        d_1341_v22_ = _dafny.MultiSet([(self).f25, (self).f25, (self).f25])
        d_1342_v23_: _dafny.Map
        d_1342_v23_ = _dafny.Map({d_1341_v22_: (self).f27})
        r2 = (d_1342_v23_) | ((d_1342_v23_ if (self).f25 else d_1342_v23_))
        d_1343_v24_: _dafny.Map
        d_1343_v24_ = _dafny.Map({((self).f27) + ((self).f26): (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1344_i1_ in range(default__.abs(396))]) if True else self.f28)})
        r3 = d_1343_v24_
        return r0, r1, r2, r3

    def m3(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        r2: int = int(0)
        if (self).f25:
            r2 = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1345_i0_ in range(default__.abs(398))]))
            d_1346_v0_: _dafny.Array
            def lambda78_(d_1347_i1_):
                return (d_1347_i1_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f27 for d_1348_i2_ in range(default__.abs(-692))])), _dafny.MultiSet([659]), _dafny.MultiSet([-667, (self).f27]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f27]))])))

            init42_ = lambda78_
            nw236_ = _dafny.Array(None, 20)
            for i0_42_ in range(nw236_.length(0)):
                nw236_[i0_42_] = init42_(i0_42_)
            d_1346_v0_ = nw236_
            index247_ = default__.safeIndex(624, (d_1346_v0_).length(0))
            (d_1346_v0_)[index247_] = (0) - ((self).f27)
            d_1349_v1_: D12
            d_1349_v1_ = D12_DC28(d_1346_v0_)
            d_1350_v2_: _dafny.Array
            nw237_ = _dafny.Array(False, 27)
            d_1350_v2_ = nw237_
            index248_ = default__.safeIndex(491, (d_1350_v2_).length(0))
            (d_1350_v2_)[index248_] = True
            d_1351_v3_: _dafny.Set
            d_1351_v3_ = _dafny.Set({(self).f25, (self).f25, (self).f25})
            index249_ = default__.safeIndex(624, (d_1346_v0_).length(0))
            index250_ = default__.safeIndex(491, (d_1350_v2_).length(0))
            rhs269_ = ((self).f26 if (self).f25 else 229)
            rhs270_ = (self).f27
            rhs271_ = self.f28
            rhs272_ = d_1349_v1_
            rhs273_ = (d_1351_v3_).issubset(_dafny.Set({not((self).f25)}))
            lhs222_ = globalState
            lhs223_ = d_1346_v0_
            lhs224_ = default__.safeIndex(624, (d_1346_v0_).length(0))
            lhs225_ = self
            lhs226_ = d_1350_v2_
            lhs227_ = default__.safeIndex(491, (d_1350_v2_).length(0))
            lhs222_.f0 = rhs269_
            lhs223_[lhs224_] = rhs270_
            lhs225_.f28 = rhs271_
            d_1349_v1_ = rhs272_
            lhs226_[lhs227_] = rhs273_
            (self).f28 = self.f28
            index251_ = default__.safeIndex(624, (d_1346_v0_).length(0))
            rhs274_ = (d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))]
            rhs275_ = (self).fm6((self).f25, (self).f25, (self).f26, globalState)
            lhs228_ = d_1346_v0_
            lhs229_ = default__.safeIndex(624, (d_1346_v0_).length(0))
            lhs230_ = globalState
            lhs228_[lhs229_] = rhs274_
            lhs230_.f2 = rhs275_
            if ((d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))]) <= ((d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))]):
                d_1352_v4_: _dafny.Map
                d_1352_v4_ = _dafny.Map({(self).f27: (d_1350_v2_)[default__.safeIndex(491, (d_1350_v2_).length(0))]})
                d_1353_v5_: C1
                nw238_ = C1()
                nw238_.ctor__((self).f25)
                d_1353_v5_ = nw238_
                d_1354_v6_: _dafny.MultiSet
                d_1354_v6_ = _dafny.MultiSet([d_1353_v5_])
                d_1355_v8_: D0
                d_1355_v8_ = D0_DC0((d_1350_v2_)[default__.safeIndex(491, (d_1350_v2_).length(0))])
                d_1356_v9_: _dafny.Set
                d_1356_v9_ = _dafny.Set({d_1355_v8_})
                d_1357_v10_: _dafny.Map
                def iife114_():
                    coll30_ = _dafny.Map()
                    compr_30_: D0
                    for compr_30_ in (d_1356_v9_).Elements:
                        d_1358_v7_: D0 = compr_30_
                        if (d_1358_v7_) in (d_1356_v9_):
                            coll30_[d_1358_v7_] = (self).f26
                    return _dafny.Map(coll30_)
                d_1357_v10_ = _dafny.Map({(self).f26: ((d_1352_v4_)[((d_1354_v6_)[d_1353_v5_] if (d_1353_v5_) in (d_1354_v6_) else (d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))])] if (((d_1354_v6_)[d_1353_v5_] if (d_1353_v5_) in (d_1354_v6_) else (d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))])) in (d_1352_v4_) else default__.fm5((D1_DC4((self).f25, (self).f25)).cf8, iife114_()
                , (d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))], globalState))})
                d_1359_v11_: _dafny.Seq
                d_1359_v11_ = _dafny.SeqWithoutIsStrInference([(d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))], (d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))]])
                d_1360_v12_: _dafny.Seq
                d_1360_v12_ = _dafny.SeqWithoutIsStrInference([(d_1353_v5_).f32])
                d_1357_v10_ = (d_1357_v10_).set((len(d_1357_v10_)) + (len((d_1359_v11_).set(default__.safeIndex((d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))], len(d_1359_v11_)), (d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))]))), ((d_1360_v12_)[default__.safeIndex((self).f27, len(d_1360_v12_))] if (d_1350_v2_)[default__.safeIndex(491, (d_1350_v2_).length(0))] else (d_1350_v2_)[default__.safeIndex(491, (d_1350_v2_).length(0))]))
                index252_ = default__.safeIndex(491, (d_1350_v2_).length(0))
                (d_1350_v2_)[index252_] = ((d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))]) != ((d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))])
                d_1361_v13_: _dafny.Map
                d_1361_v13_ = _dafny.Map({d_1346_v0_: (self).f25})
                index253_ = default__.safeIndex(624, (d_1346_v0_).length(0))
                (d_1346_v0_)[index253_] = len(d_1361_v13_)
                index254_ = default__.safeIndex(491, (d_1350_v2_).length(0))
                (d_1350_v2_)[index254_] = not((self).f25)
                d_1350_v2_ = d_1350_v2_
            elif True:
                index255_ = default__.safeIndex(624, (d_1346_v0_).length(0))
                (d_1346_v0_)[index255_] = (0) - ((self).f26)
                d_1362_v14_: _dafny.Map
                d_1362_v14_ = _dafny.Map({(self).f25: (self).f27})
                d_1363_v15_: _dafny.Map
                d_1363_v15_ = _dafny.Map({((d_1362_v14_)[(self).f25] if ((self).f25) in (d_1362_v14_) else len(self.f28)): self.f28})
                d_1363_v15_ = (d_1363_v15_).set((d_1346_v0_)[default__.safeIndex(624, (d_1346_v0_).length(0))], (self.f28) + (self.f28))
                d_1364_v16_: str
                d_1364_v16_ = _dafny.CodePoint('u')
                d_1365_v17_: _dafny.Map
                d_1365_v17_ = _dafny.Map({(self).f25: ((self.f28)[default__.safeIndex((self).f26, len(self.f28))] if (d_1350_v2_)[default__.safeIndex(491, (d_1350_v2_).length(0))] else d_1364_v16_)})
                d_1366_v18_: _dafny.Seq
                d_1366_v18_ = _dafny.SeqWithoutIsStrInference([False])
                d_1365_v17_ = (d_1365_v17_).set((d_1350_v2_)[default__.safeIndex(491, (d_1350_v2_).length(0))], (d_1364_v16_ if (d_1366_v18_)[default__.safeIndex((self).f27, len(d_1366_v18_))] else d_1364_v16_))
                d_1367_v19_: C1
                nw239_ = C1()
                nw239_.ctor__(False)
                d_1367_v19_ = nw239_
                d_1368_v20_: _dafny.Seq
                d_1368_v20_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])
                d_1369_v21_: D11
                d_1369_v21_ = D11_DC26((self).f26, 797, (self).f25)
                d_1368_v20_ = default__.fm3((d_1369_v21_).cf41, not(True), (self).f26, (self).f25, globalState)
        elif True:
            d_1370_v22_: D0
            d_1370_v22_ = D0_DC0((self).f25)
            d_1371_v23_: _dafny.Map
            d_1371_v23_ = _dafny.Map({default__.fm5(True, _dafny.Map({d_1370_v22_: (self).f26}), (self).f27, globalState): (self).f25})
            d_1371_v23_ = _dafny.Map({(self).f25: not((self).f25)})
            (globalState).f0 = (self).f27
            d_1372_v24_: str
            d_1372_v24_ = _dafny.CodePoint('l')
            d_1373_v25_: _dafny.Map
            d_1373_v25_ = _dafny.Map({d_1372_v24_: (self).f27})
            r1 = ((self).f26) >= (((self).f27) - (((d_1373_v25_)[_dafny.CodePoint('f')] if (_dafny.CodePoint('f')) in (d_1373_v25_) else (self).f27)))
            d_1374_v26_: _dafny.Array
            nw240_ = _dafny.Array(int(0), 24)
            d_1374_v26_ = nw240_
            index256_ = default__.safeIndex(352, (d_1374_v26_).length(0))
            (d_1374_v26_)[index256_] = (self).f26
            index257_ = default__.safeIndex(352, (d_1374_v26_).length(0))
            (d_1374_v26_)[index257_] = (self).f26
            if (self).f25:
                rhs276_ = (self.f28) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))
                rhs277_ = (self).f26
                lhs231_ = self
                lhs232_ = globalState
                lhs231_.f28 = rhs276_
                lhs232_.f0 = rhs277_
                d_1375_v27_: C2
                nw241_ = C2()
                nw241_.ctor__((self).f27, False, (self).f26)
                d_1375_v27_ = nw241_
                d_1376_v28_: C1
                nw242_ = C1()
                nw242_.ctor__(not(False))
                d_1376_v28_ = nw242_
                d_1377_v29_: _dafny.Array
                def lambda79_(d_1378_i3_):
                    return (self.f28) + (self.f28)

                init43_ = lambda79_
                nw243_ = _dafny.Array(None, 12)
                for i0_43_ in range(nw243_.length(0)):
                    nw243_[i0_43_] = init43_(i0_43_)
                d_1377_v29_ = nw243_
                index258_ = default__.safeIndex(838, (d_1377_v29_).length(0))
                (d_1377_v29_)[index258_] = self.f28
                index259_ = default__.safeIndex(838, (d_1377_v29_).length(0))
                (d_1377_v29_)[index259_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgnv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqotdmdy")))) + ((self.f28) + (self.f28))
                d_1379_v30_: _dafny.Seq
                d_1379_v30_ = _dafny.SeqWithoutIsStrInference([(self).f25])
                d_1380_v31_: C3
                nw244_ = C3()
                nw244_.ctor__(self.f28, (self).f25, (self).f25, len(d_1379_v30_))
                d_1380_v31_ = nw244_
                d_1381_v32_: _dafny.Map
                d_1381_v32_ = _dafny.Map({d_1380_v31_: True})
                d_1382_v33_: _dafny.Seq
                d_1382_v33_ = _dafny.SeqWithoutIsStrInference([d_1380_v31_])
                d_1383_v34_: _dafny.MultiSet
                d_1383_v34_ = _dafny.MultiSet([(self).f25, True, False])
                d_1384_v35_: _dafny.Map
                d_1384_v35_ = _dafny.Map({len((d_1377_v29_)[default__.safeIndex(838, (d_1377_v29_).length(0))]): d_1383_v34_})
                d_1381_v32_ = (d_1381_v32_).set((d_1382_v33_)[default__.safeIndex((0) - ((((d_1384_v35_)[(self).f26] if ((self).f26) in (d_1384_v35_) else d_1383_v34_)).cardinality), len(d_1382_v33_))], (d_1376_v28_).f32)
            elif True:
                d_1385_v36_: D1
                d_1385_v36_ = D1_DC4((self).f25, (self).f25)
                (globalState).f1 = (d_1385_v36_).cf7
                (globalState).f1 = (self).f25
                (globalState).f2 = default__.safeModuloInt(len(self.f28), -450)
                d_1386_v37_: _dafny.Seq
                d_1386_v37_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_1386_v37_ = d_1386_v37_
                d_1387_v38_: _dafny.Set
                d_1387_v38_ = _dafny.Set({(self).f25})
                (globalState).f1 = ((_dafny.Set({(self).f25})) - (d_1387_v38_)).issubset(d_1387_v38_)
        d_1388_v39_: str
        d_1388_v39_ = _dafny.CodePoint('o')
        d_1389_v40_: D1
        d_1389_v40_ = D1_DC3(d_1388_v39_)
        def lambda80_(source20_):
            if source20_.is_DC4:
                d_1390___mcc_h0_ = source20_.cf7
                d_1391___mcc_h1_ = source20_.cf8
                d_1392_cf8_ = d_1391___mcc_h1_
                d_1393_cf7_ = d_1390___mcc_h0_
                def iife115_():
                    coll31_ = _dafny.Set()
                    compr_31_: int
                    for compr_31_ in _dafny.IntegerRange(-349, 15):
                        d_1394_v41_: int = compr_31_
                        if ((-349) <= (d_1394_v41_)) and ((d_1394_v41_) < (15)):
                            coll31_ = coll31_.union(_dafny.Set([(d_1394_v41_) - (len(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])))]))
                    return _dafny.Set(coll31_)
                return (iife115_()
                ).ispropersubset(_dafny.Set({(self).f27}))
            elif source20_.is_DC3:
                d_1395___mcc_h2_ = source20_.cf6
                d_1396_cf6_ = d_1395___mcc_h2_
                return (self).f25
            elif True:
                d_1397___mcc_h3_ = source20_.cf9
                d_1398_cf9_ = d_1397___mcc_h3_
                return (self).f25

        if lambda80_(d_1389_v40_):
            d_1388_v39_ = d_1388_v39_
            (globalState).f0 = (self).f26
            d_1399_v42_: _dafny.Array
            def lambda81_(d_1400_i4_):
                return _dafny.Map({(self).f25: (self).f25})

            init44_ = lambda81_
            nw245_ = _dafny.Array(None, 19)
            for i0_44_ in range(nw245_.length(0)):
                nw245_[i0_44_] = init44_(i0_44_)
            d_1399_v42_ = nw245_
            d_1401_v43_: D8
            d_1401_v43_ = D8_DC16(d_1399_v42_)
            d_1402_v44_: _dafny.Array
            nw246_ = _dafny.Array(None, 7)
            nw246_[int(0)] = D8_DC16(d_1399_v42_)
            nw246_[int(1)] = d_1401_v43_
            nw246_[int(2)] = d_1401_v43_
            nw246_[int(3)] = d_1401_v43_
            nw246_[int(4)] = d_1401_v43_
            nw246_[int(5)] = d_1401_v43_
            nw246_[int(6)] = d_1401_v43_
            d_1402_v44_ = nw246_
            d_1403_v45_: _dafny.Map
            d_1403_v45_ = _dafny.Map({(self).f25: d_1402_v44_})
            d_1404_v46_: _dafny.Seq
            d_1404_v46_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f25: d_1402_v44_})])
            d_1403_v45_ = (d_1404_v46_)[default__.safeIndex((self).f27, len(d_1404_v46_))]
            d_1405_v47_: _dafny.Seq
            d_1405_v47_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])
            (globalState).f2 = (703) * (len(_dafny.Map({d_1405_v47_: (self).f25})))
            d_1406_v48_: C3
            nw247_ = C3()
            nw247_.ctor__(self.f28, (self).f25, (self).f25, -752)
            d_1406_v48_ = nw247_
        elif True:
            d_1407_v49_: _dafny.Set
            d_1407_v49_ = _dafny.Set({(self).f25})
            if ((d_1407_v49_) | (d_1407_v49_)).isdisjoint(_dafny.Set({(self).f25, (self).f25})):
                d_1408_v50_: _dafny.Array
                nw248_ = _dafny.Array(int(0), 15)
                d_1408_v50_ = nw248_
                d_1408_v50_ = d_1408_v50_
                d_1409_v51_: _dafny.Map
                d_1409_v51_ = _dafny.Map({(self).f25: (self).f25})
                d_1409_v51_ = (d_1409_v51_).set((self).f25, not((self).f25))
                d_1410_v52_: _dafny.Array
                def lambda82_(d_1411_i5_):
                    return False

                init45_ = lambda82_
                nw249_ = _dafny.Array(None, 3)
                for i0_45_ in range(nw249_.length(0)):
                    nw249_[i0_45_] = init45_(i0_45_)
                d_1410_v52_ = nw249_
                d_1412_v53_: C0
                nw250_ = C0()
                nw250_.ctor__(self.f28, (self).f25, (self).f26, d_1410_v52_, (self).f25)
                d_1412_v53_ = nw250_
                d_1413_v54_: C3
                nw251_ = C3()
                nw251_.ctor__(self.f28, (self).f25, (self).f25, (self).f26)
                d_1413_v54_ = nw251_
                d_1414_v55_: _dafny.Array
                nw252_ = _dafny.Array(None, 12)
                nw252_[int(0)] = d_1413_v54_
                nw252_[int(1)] = d_1413_v54_
                nw252_[int(2)] = d_1413_v54_
                nw252_[int(3)] = d_1413_v54_
                nw252_[int(4)] = d_1413_v54_
                nw252_[int(5)] = d_1413_v54_
                nw252_[int(6)] = d_1413_v54_
                nw252_[int(7)] = d_1413_v54_
                nw252_[int(8)] = d_1413_v54_
                nw252_[int(9)] = (d_1413_v54_ if False else d_1413_v54_)
                nw252_[int(10)] = d_1413_v54_
                nw252_[int(11)] = d_1413_v54_
                d_1414_v55_ = nw252_
                d_1415_v56_: _dafny.Seq
                d_1415_v56_ = _dafny.SeqWithoutIsStrInference([d_1410_v52_])
                d_1416_v57_: D0
                d_1416_v57_ = D0_DC0((d_1413_v54_).f30)
                d_1417_v58_: _dafny.Map
                d_1417_v58_ = _dafny.Map({d_1416_v57_: (0) - ((self).f26)})
                d_1418_v59_: _dafny.MultiSet
                d_1418_v59_ = _dafny.MultiSet([(self).f25])
                d_1419_v60_: _dafny.Map
                d_1419_v60_ = _dafny.Map({(d_1418_v59_).cardinality: (self).f26})
                rhs278_ = (self).f26
                rhs279_ = (d_1414_v55_ if (d_1410_v52_) in (d_1415_v56_) else d_1414_v55_)
                rhs280_ = default__.fm5((d_1413_v54_).f30, d_1417_v58_, (self).f27, globalState)
                rhs281_ = (0) - ((((self).f26 if (d_1413_v54_).f30 else (self).f26) if ((d_1413_v54_).fm9(globalState)) not in (d_1419_v60_) else default__.safeModuloInt((self).f27, (self).f26)))
                lhs233_ = globalState
                lhs234_ = globalState
                lhs235_ = globalState
                lhs233_.f0 = rhs278_
                d_1414_v55_ = rhs279_
                lhs234_.f1 = rhs280_
                lhs235_.f2 = rhs281_
                d_1420_v61_: _dafny.Array
                def lambda83_(d_1421_v39_):
                    def lambda84_(d_1422_i6_):
                        return _dafny.SeqWithoutIsStrInference([d_1421_v39_ for d_1423_i7_ in range(default__.abs(496))])

                    return lambda84_

                init46_ = lambda83_(d_1388_v39_)
                nw253_ = _dafny.Array(None, 23)
                for i0_46_ in range(nw253_.length(0)):
                    nw253_[i0_46_] = init46_(i0_46_)
                d_1420_v61_ = nw253_
                d_1420_v61_ = d_1420_v61_
            elif True:
                d_1424_v62_: _dafny.MultiSet
                d_1424_v62_ = _dafny.MultiSet([(self).f27])
                d_1425_v63_: C2
                nw254_ = C2()
                nw254_.ctor__((((d_1424_v62_)[824] if (824) in (d_1424_v62_) else (self).f26)) + ((self).f27), (self).f25, (self).f27)
                d_1425_v63_ = nw254_
                d_1425_v63_ = d_1425_v63_
                (globalState).f0 = (default__.safeDivisionInt((d_1425_v63_).f31, (self).f27)) * ((d_1425_v63_).f31)
                d_1426_v64_: _dafny.Array
                nw255_ = _dafny.Array(_dafny.Map({}), 13)
                d_1426_v64_ = nw255_
                d_1427_v65_: _dafny.Array
                nw256_ = _dafny.Array(None, 9)
                nw256_[int(0)] = (self).f27
                nw256_[int(1)] = (self).f26
                nw256_[int(2)] = (d_1425_v63_).f31
                nw256_[int(3)] = (d_1425_v63_).f31
                nw256_[int(4)] = (d_1425_v63_).f31
                nw256_[int(5)] = (self).f27
                nw256_[int(6)] = (self).f27
                nw256_[int(7)] = (d_1425_v63_).f31
                nw256_[int(8)] = 514
                d_1427_v65_ = nw256_
                d_1428_v66_: _dafny.Map
                d_1428_v66_ = _dafny.Map({d_1427_v65_: (self).f25})
                index260_ = default__.safeIndex(567, (d_1426_v64_).length(0))
                (d_1426_v64_)[index260_] = d_1428_v66_
                index261_ = default__.safeIndex(567, (d_1426_v64_).length(0))
                (d_1426_v64_)[index261_] = d_1428_v66_
                (globalState).f0 = ((d_1424_v62_)[271] if (271) in (d_1424_v62_) else (self).f27)
                d_1429_v67_: _dafny.MultiSet
                d_1429_v67_ = _dafny.MultiSet([self.f28, self.f28])
                d_1429_v67_ = (d_1429_v67_) - ((d_1429_v67_).intersection(d_1429_v67_))
            d_1430_v68_: _dafny.Map
            d_1430_v68_ = _dafny.Map({(self).f27: (self).f25})
            d_1430_v68_ = (d_1430_v68_).set(len(self.f28), not((self).f25))
            d_1431_v69_: _dafny.Seq
            d_1431_v69_ = _dafny.SeqWithoutIsStrInference([self.f28])
            d_1432_v70_: C3
            nw257_ = C3()
            nw257_.ctor__((self.f28) + ((d_1431_v69_)[default__.safeIndex((self).f26, len(d_1431_v69_))]), (self).f25, (self).f25, (self).f27)
            d_1432_v70_ = nw257_
            d_1433_v71_: _dafny.Array
            nw258_ = _dafny.Array(_dafny.Seq({}), 5)
            d_1433_v71_ = nw258_
            d_1434_v72_: _dafny.Seq
            d_1434_v72_ = _dafny.SeqWithoutIsStrInference([(d_1432_v70_).f30])
            index262_ = default__.safeIndex(380, (d_1433_v71_).length(0))
            (d_1433_v71_)[index262_] = (d_1434_v72_) + (d_1434_v72_)
            index263_ = default__.safeIndex(380, (d_1433_v71_).length(0))
            (d_1433_v71_)[index263_] = d_1434_v72_
            d_1435_v73_: _dafny.MultiSet
            d_1435_v73_ = _dafny.MultiSet([len(self.f28), (self).f26, (self).fm6((self).f25, (d_1432_v70_).f30, (0) - ((self).f27), globalState)])
            (globalState).f0 = ((d_1435_v73_)[len(_dafny.Map({d_1407_v49_: (self).fm7(globalState)}))] if (len(_dafny.Map({d_1407_v49_: (self).fm7(globalState)}))) in (d_1435_v73_) else len((self.f28) + (self.f28)))
        hi13_ = 801
        for d_1436_i8_ in range(((self).f26) + ((self).f27), hi13_):
            d_1437_v74_: _dafny.Array
            def lambda85_(d_1438_i9_):
                return (self).f25

            init47_ = lambda85_
            nw259_ = _dafny.Array(None, 4)
            for i0_47_ in range(nw259_.length(0)):
                nw259_[i0_47_] = init47_(i0_47_)
            d_1437_v74_ = nw259_
            index264_ = default__.safeIndex(96, (d_1437_v74_).length(0))
            (d_1437_v74_)[index264_] = (self).f25
            d_1439_v75_: _dafny.Array
            nw260_ = _dafny.Array(_dafny.MultiSet({}), 21)
            d_1439_v75_ = nw260_
            d_1440_v76_: _dafny.MultiSet
            d_1440_v76_ = _dafny.MultiSet([d_1437_v74_, d_1437_v74_, d_1437_v74_])
            index265_ = default__.safeIndex(518, (d_1439_v75_).length(0))
            (d_1439_v75_)[index265_] = d_1440_v76_
            d_1441_v77_: _dafny.Array
            nw261_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_1441_v77_ = nw261_
            d_1442_v78_: _dafny.Array
            nw262_ = _dafny.Array(_dafny.MultiSet({}), 25)
            d_1442_v78_ = nw262_
            index266_ = default__.safeIndex(497, (d_1441_v77_).length(0))
            (d_1441_v77_)[index266_] = d_1442_v78_
            d_1443_v79_: _dafny.Seq
            d_1443_v79_ = _dafny.SeqWithoutIsStrInference([(self).f25, (not((self).f25) if (self).f25 else (self).f25), ((self).f25) or ((self).f25)])
            d_1444_v80_: _dafny.Map
            d_1444_v80_ = _dafny.Map({(self).f27: d_1442_v78_})
            d_1445_v81_: _dafny.Set
            d_1445_v81_ = _dafny.Set({(self).f25})
            index267_ = default__.safeIndex(96, (d_1437_v74_).length(0))
            index268_ = default__.safeIndex(518, (d_1439_v75_).length(0))
            index269_ = default__.safeIndex(497, (d_1441_v77_).length(0))
            rhs282_ = not((self).f25)
            rhs283_ = (d_1443_v79_)[default__.safeIndex((self).f27, len(d_1443_v79_))]
            rhs284_ = d_1440_v76_
            rhs285_ = ((d_1444_v80_)[len(default__.fm36((self).f27, len(self.f28), self.f28, (self).f27, globalState))] if (len(default__.fm36((self).f27, len(self.f28), self.f28, (self).f27, globalState))) in (d_1444_v80_) else (d_1442_v78_ if (self).f25 else d_1442_v78_))
            rhs286_ = len(d_1445_v81_)
            lhs236_ = d_1437_v74_
            lhs237_ = default__.safeIndex(96, (d_1437_v74_).length(0))
            lhs238_ = globalState
            lhs239_ = d_1439_v75_
            lhs240_ = default__.safeIndex(518, (d_1439_v75_).length(0))
            lhs241_ = d_1441_v77_
            lhs242_ = default__.safeIndex(497, (d_1441_v77_).length(0))
            lhs243_ = globalState
            lhs236_[lhs237_] = rhs282_
            lhs238_.f1 = rhs283_
            lhs239_[lhs240_] = rhs284_
            lhs241_[lhs242_] = rhs285_
            lhs243_.f0 = rhs286_
            (globalState).f2 = default__.fm2((0) - (default__.fm2((self).f26, (self).f27, globalState)), (self).f27, globalState)
            d_1446_v82_: _dafny.Array
            nw263_ = _dafny.Array(_dafny.Set({}), 5)
            d_1446_v82_ = nw263_
            d_1447_v83_: D0
            d_1447_v83_ = D0_DC0((d_1437_v74_)[default__.safeIndex(96, (d_1437_v74_).length(0))])
            d_1448_v84_: _dafny.Map
            d_1448_v84_ = _dafny.Map({(d_1437_v74_)[default__.safeIndex(96, (d_1437_v74_).length(0))]: (self).f25})
            d_1449_v85_: _dafny.Map
            d_1449_v85_ = _dafny.Map({d_1447_v83_: len((d_1448_v84_).set(False, (self).f25))})
            index270_ = default__.safeIndex(171, (d_1446_v82_).length(0))
            (d_1446_v82_)[index270_] = (d_1445_v81_).intersection(_dafny.Set({(d_1437_v74_)[default__.safeIndex(96, (d_1437_v74_).length(0))], (self).f25, (d_1437_v74_)[default__.safeIndex(96, (d_1437_v74_).length(0))], (self).f25, default__.fm5((d_1437_v74_)[default__.safeIndex(96, (d_1437_v74_).length(0))], d_1449_v85_, len(_dafny.SeqWithoutIsStrInference([(self).f27 for d_1450_i10_ in range(default__.abs(-270))])), globalState)}))
            index271_ = default__.safeIndex(171, (d_1446_v82_).length(0))
            (d_1446_v82_)[index271_] = d_1445_v81_
            d_1451_v86_: _dafny.Map
            d_1451_v86_ = _dafny.Map({605: (self).f27})
            d_1452_v87_: _dafny.MultiSet
            d_1452_v87_ = _dafny.MultiSet([(self).f26])
            d_1453_v88_: _dafny.Map
            d_1453_v88_ = _dafny.Map({len((d_1446_v82_)[default__.safeIndex(171, (d_1446_v82_).length(0))]): d_1452_v87_})
            d_1454_v89_: C2
            nw264_ = C2()
            nw264_.ctor__((self).f27, (d_1437_v74_)[default__.safeIndex(96, (d_1437_v74_).length(0))], (0) - (len(d_1453_v88_)))
            d_1454_v89_ = nw264_
            d_1455_v90_: _dafny.Set
            d_1455_v90_ = _dafny.Set({d_1454_v89_})
            d_1456_v91_: _dafny.Set
            d_1456_v91_ = _dafny.Set({(self).f26, d_1436_i8_, len(d_1455_v90_)})
            d_1457_v92_: _dafny.Seq
            d_1457_v92_ = _dafny.SeqWithoutIsStrInference([len(d_1456_v91_), d_1436_i8_, (self).f26, (d_1454_v89_).f31, (d_1454_v89_).f31])
            (globalState).f2 = default__.fm2(len((default__.fm3((self).f27, (d_1437_v74_)[default__.safeIndex(96, (d_1437_v74_).length(0))], len(d_1451_v86_), not(not((self).f25)), globalState)) + (d_1457_v92_)), (self).f27, globalState)
        if not (not ((self).f25) or (not((self).f25))) or ((False if (self).f25 else (self).f25)):
            (globalState).f1 = (self).f25
            d_1458_v93_: _dafny.Seq
            d_1458_v93_ = _dafny.SeqWithoutIsStrInference([False])
            d_1459_v94_: _dafny.Map
            d_1459_v94_ = _dafny.Map({True: (d_1458_v93_)[default__.safeIndex((self).f27, len(d_1458_v93_))]})
            d_1459_v94_ = (d_1459_v94_).set((self).f25, (self).f25)
            d_1460_v95_: C2
            nw265_ = C2()
            nw265_.ctor__((self).f26, (self).f25, (self).f26)
            d_1460_v95_ = nw265_
            d_1461_v96_: _dafny.Seq
            d_1461_v96_ = _dafny.SeqWithoutIsStrInference([d_1460_v95_, d_1460_v95_])
            d_1462_v97_: _dafny.Map
            d_1462_v97_ = _dafny.Map({not((self).f25): (d_1461_v96_)[default__.safeIndex(-67, len(d_1461_v96_))]})
            d_1463_v98_: _dafny.Seq
            d_1463_v98_ = _dafny.SeqWithoutIsStrInference([((d_1462_v97_)[(self).f25] if ((self).f25) in (d_1462_v97_) else d_1460_v95_)])
            d_1464_v99_: _dafny.MultiSet
            d_1464_v99_ = _dafny.MultiSet([(d_1463_v98_) + (d_1461_v96_)])
            d_1465_v100_: _dafny.MultiSet
            d_1465_v100_ = (d_1464_v99_).set(d_1463_v98_, default__.abs((self).f26))
            d_1466_v101_: _dafny.Seq
            d_1466_v101_ = _dafny.SeqWithoutIsStrInference([d_1463_v98_])
            d_1464_v99_ = ((d_1464_v99_) | (_dafny.MultiSet([d_1461_v96_]))) | ((_dafny.MultiSet(d_1466_v101_)))
            d_1467_v102_: _dafny.Map
            d_1467_v102_ = _dafny.Map({(d_1460_v95_).f31: (self).f25})
            d_1467_v102_ = (d_1467_v102_).set((d_1460_v95_).f31, True)
            (globalState).f1 = (self).f25
        elif True:
            d_1468_v103_: _dafny.MultiSet
            d_1468_v103_ = _dafny.MultiSet([(self).f26])
            d_1469_v104_: _dafny.Map
            d_1469_v104_ = _dafny.Map({541: (d_1468_v103_).cardinality})
            d_1470_v105_: _dafny.Map
            d_1470_v105_ = _dafny.Map({d_1469_v104_: (self).f25})
            d_1471_v106_: _dafny.Seq
            d_1471_v106_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f25, ((d_1470_v105_)[d_1469_v104_] if (d_1469_v104_) in (d_1470_v105_) else not((self).f25)), (self).f25])
            if (d_1471_v106_) <= (d_1471_v106_):
                d_1472_v107_: _dafny.Array
                nw266_ = _dafny.Array(int(0), 28)
                d_1472_v107_ = nw266_
                d_1473_v108_: D12
                d_1473_v108_ = D12_DC28(d_1472_v107_)
                d_1474_v109_: _dafny.Array
                nw267_ = _dafny.Array(None, 27)
                nw267_[int(0)] = d_1472_v107_
                nw267_[int(1)] = d_1472_v107_
                nw267_[int(2)] = d_1472_v107_
                nw267_[int(3)] = d_1472_v107_
                nw267_[int(4)] = d_1472_v107_
                nw267_[int(5)] = d_1472_v107_
                nw267_[int(6)] = d_1472_v107_
                nw267_[int(7)] = d_1472_v107_
                nw267_[int(8)] = d_1472_v107_
                nw267_[int(9)] = (d_1472_v107_ if False else d_1472_v107_)
                nw267_[int(10)] = d_1472_v107_
                nw267_[int(11)] = d_1472_v107_
                nw267_[int(12)] = d_1472_v107_
                nw267_[int(13)] = (d_1473_v108_).cf44
                nw267_[int(14)] = d_1472_v107_
                nw267_[int(15)] = d_1472_v107_
                nw267_[int(16)] = (d_1472_v107_ if (self).f25 else d_1472_v107_)
                nw267_[int(17)] = d_1472_v107_
                nw267_[int(18)] = d_1472_v107_
                nw267_[int(19)] = (d_1472_v107_ if (self).f25 else d_1472_v107_)
                nw267_[int(20)] = d_1472_v107_
                nw267_[int(21)] = d_1472_v107_
                nw267_[int(22)] = d_1472_v107_
                nw267_[int(23)] = d_1472_v107_
                nw267_[int(24)] = d_1472_v107_
                nw267_[int(25)] = d_1472_v107_
                nw267_[int(26)] = d_1472_v107_
                d_1474_v109_ = nw267_
                index272_ = default__.safeIndex(7, (d_1474_v109_).length(0))
                (d_1474_v109_)[index272_] = d_1472_v107_
                index273_ = default__.safeIndex(7, (d_1474_v109_).length(0))
                (d_1474_v109_)[index273_] = d_1472_v107_
                d_1475_v110_: _dafny.Array
                def lambda86_(d_1476_i11_):
                    return (self).f25

                init48_ = lambda86_
                nw268_ = _dafny.Array(None, 25)
                for i0_48_ in range(nw268_.length(0)):
                    nw268_[i0_48_] = init48_(i0_48_)
                d_1475_v110_ = nw268_
                index274_ = default__.safeIndex(573, (d_1475_v110_).length(0))
                (d_1475_v110_)[index274_] = not ((self).f25) or (True)
                index275_ = default__.safeIndex(573, (d_1475_v110_).length(0))
                (d_1475_v110_)[index275_] = (self).f25
                (globalState).f0 = default__.safeModuloInt((self).f26, (self).f27)
                d_1477_v111_: C3
                nw269_ = C3()
                nw269_.ctor__(self.f28, (self).f25, False, (self).f27)
                d_1477_v111_ = nw269_
                d_1478_v112_: _dafny.Map
                d_1478_v112_ = _dafny.Map({((d_1468_v103_)[(self).f27] if ((self).f27) in (d_1468_v103_) else 205): d_1477_v111_})
                d_1469_v104_ = (d_1469_v104_).set((-314) - (len(d_1478_v112_)), (self).f27)
                index276_ = default__.safeIndex(573, (d_1475_v110_).length(0))
                (d_1475_v110_)[index276_] = (d_1468_v103_) == (d_1468_v103_)
            elif True:
                d_1479_v113_: D1
                d_1479_v113_ = D1_DC4((self).f25, (self).f25)
                d_1480_v114_: D1
                d_1480_v114_ = D1_DC5(D1_DC5(d_1479_v113_))
                d_1481_v115_: _dafny.Seq
                d_1481_v115_ = _dafny.SeqWithoutIsStrInference([d_1480_v114_, d_1480_v114_])
                d_1482_v116_: _dafny.Set
                d_1482_v116_ = _dafny.Set({d_1481_v115_})
                d_1483_v117_: _dafny.Seq
                d_1483_v117_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D1_DC5(d_1479_v113_) for d_1484_i12_ in range(default__.abs(839))])])
                d_1485_v119_: _dafny.Map
                def iife116_():
                    coll32_ = _dafny.Set()
                    compr_32_: _dafny.Seq
                    for compr_32_ in (d_1483_v117_).Elements:
                        d_1486_v118_: _dafny.Seq = compr_32_
                        if (d_1486_v118_) in (d_1483_v117_):
                            coll32_ = coll32_.union(_dafny.Set([d_1486_v118_]))
                    return _dafny.Set(coll32_)
                d_1485_v119_ = _dafny.Map({(self).f25: (d_1482_v116_) - (iife116_()
                )})
                d_1485_v119_ = (d_1485_v119_).set((self).f25, d_1482_v116_)
                d_1487_v120_: _dafny.Map
                d_1487_v120_ = _dafny.Map({True: (self).f27})
                d_1488_v121_: _dafny.Array
                nw270_ = _dafny.Array(False, 22)
                d_1488_v121_ = nw270_
                r0 = (d_1488_v121_ if (d_1487_v120_) != (d_1487_v120_) else d_1488_v121_)
                r1 = (d_1388_v39_) in (default__.fm1(globalState))
                (globalState).f0 = (((d_1469_v104_)[(self).f26] if ((self).f26) in (d_1469_v104_) else (self).f26)) - ((self).f27)
                r1 = (self).f25
            d_1489_v122_: _dafny.Map
            d_1489_v122_ = _dafny.Map({(0) - ((self).f26): _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f25: (self).f25})) for d_1490_i13_ in range(default__.abs(295))])): (self).f25})})
            r2 = default__.fm2(len(d_1489_v122_), (self).f26, globalState)
            d_1491_v123_: int
            out33_: int
            out33_ = default__.m0((d_1471_v106_)[default__.safeIndex((self).f26, len(d_1471_v106_))], globalState)
            d_1491_v123_ = out33_
            r2 = (self).f27
            (globalState).f1 = (self).f25
        (globalState).f2 = (0) - (((len(_dafny.SeqWithoutIsStrInference([782]))) + ((self).f26)) + ((0) - ((self).f27)))
        d_1492_v124_: D0
        d_1492_v124_ = D0_DC0((self).f25)
        d_1493_v125_: _dafny.Map
        d_1493_v125_ = _dafny.Map({default__.fm28(_dafny.SeqWithoutIsStrInference([d_1388_v39_ for d_1494_i14_ in range(default__.abs(973))]), d_1388_v39_, -187, globalState): (self).f26})
        d_1495_v126_: _dafny.MultiSet
        d_1495_v126_ = _dafny.MultiSet([True, (self).f25])
        if default__.fm5((d_1492_v124_).cf0, d_1493_v125_, ((d_1495_v126_)[(self).f25] if ((self).f25) in (d_1495_v126_) else (self).f27), globalState):
            d_1496_v127_: _dafny.Map
            d_1496_v127_ = _dafny.Map({(self).f25: (self).f27})
            d_1497_v128_: _dafny.Map
            d_1497_v128_ = _dafny.Map({(self).f26: (len(_dafny.SeqWithoutIsStrInference([(self).f25]))) != (len(d_1496_v127_))})
            d_1498_v129_: _dafny.MultiSet
            d_1498_v129_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umwt")))])
            d_1499_v130_: D11
            d_1499_v130_ = D11_DC26((0) - ((self).f26), (self).f27, (self).f25)
            d_1500_v131_: _dafny.Set
            d_1500_v131_ = _dafny.Set({(self).f27, 649})
            d_1501_v132_: _dafny.MultiSet
            d_1501_v132_ = _dafny.MultiSet([((d_1498_v129_)[(d_1499_v130_).cf41] if ((d_1499_v130_).cf41) in (d_1498_v129_) else default__.fm2((self).f27, len(d_1500_v131_), globalState)), (self).f26, (self).f27])
            d_1497_v128_ = (d_1497_v128_).set((d_1501_v132_).cardinality, (self).f25)
            d_1502_v133_: _dafny.Map
            d_1502_v133_ = _dafny.Map({d_1495_v126_: d_1388_v39_})
            d_1503_v134_: D14
            d_1503_v134_ = D14_DC31((d_1502_v133_).set(_dafny.MultiSet([(self).f25, (self).f25]), d_1388_v39_))
            d_1504_v135_: _dafny.Array
            nw271_ = _dafny.Array(None, 22)
            nw271_[int(0)] = (d_1502_v133_) | (d_1502_v133_)
            nw271_[int(1)] = d_1502_v133_
            nw271_[int(2)] = d_1502_v133_
            nw271_[int(3)] = ((d_1503_v134_).cf51).set(d_1495_v126_, d_1388_v39_)
            nw271_[int(4)] = d_1502_v133_
            nw271_[int(5)] = _dafny.Map({default__.fm27((self).f26, globalState): _dafny.CodePoint('n')})
            nw271_[int(6)] = d_1502_v133_
            nw271_[int(7)] = _dafny.Map({d_1495_v126_: d_1388_v39_})
            nw271_[int(8)] = d_1502_v133_
            nw271_[int(9)] = d_1502_v133_
            nw271_[int(10)] = d_1502_v133_
            nw271_[int(11)] = _dafny.Map({d_1495_v126_: d_1388_v39_})
            nw271_[int(12)] = d_1502_v133_
            nw271_[int(13)] = d_1502_v133_
            nw271_[int(14)] = d_1502_v133_
            nw271_[int(15)] = _dafny.Map({d_1495_v126_: d_1388_v39_})
            nw271_[int(16)] = (d_1502_v133_) | (d_1502_v133_)
            nw271_[int(17)] = d_1502_v133_
            nw271_[int(18)] = (d_1502_v133_) | (_dafny.Map({d_1495_v126_: d_1388_v39_}))
            nw271_[int(19)] = default__.fm37((self).f25, globalState)
            nw271_[int(20)] = d_1502_v133_
            nw271_[int(21)] = d_1502_v133_
            d_1504_v135_ = nw271_
            d_1505_v136_: _dafny.Seq
            d_1505_v136_ = _dafny.SeqWithoutIsStrInference([d_1500_v131_, _dafny.Set({(self).f27})])
            index277_ = default__.safeIndex(377, (d_1504_v135_).length(0))
            (d_1504_v135_)[index277_] = default__.fm37((D12_DC29((self).f27, (self).f25, d_1505_v136_, (0) - ((self).f26), len(d_1496_v127_))).cf46, globalState)
            index278_ = default__.safeIndex(377, (d_1504_v135_).length(0))
            (d_1504_v135_)[index278_] = d_1502_v133_
            source21_ = default__.fm34(self.f28, 277, globalState)
            if source21_.is_DC26:
                d_1506___mcc_h4_ = source21_.cf40
                d_1507___mcc_h5_ = source21_.cf41
                d_1508___mcc_h6_ = source21_.cf42
                d_1509_cf42_ = d_1508___mcc_h6_
                d_1510_cf41_ = d_1507___mcc_h5_
                d_1511_cf40_ = d_1506___mcc_h4_
                d_1512_v137_: _dafny.Seq
                d_1512_v137_ = _dafny.SeqWithoutIsStrInference([d_1511_cf40_, (self).f27, d_1510_cf41_, (self).f26])
                d_1512_v137_ = d_1512_v137_
                d_1513_v138_: _dafny.Array
                def lambda87_(d_1514_i15_):
                    return (self).f25

                init49_ = lambda87_
                nw272_ = _dafny.Array(None, 24)
                for i0_49_ in range(nw272_.length(0)):
                    nw272_[i0_49_] = init49_(i0_49_)
                d_1513_v138_ = nw272_
                d_1515_v139_: C0
                nw273_ = C0()
                nw273_.ctor__(self.f28, d_1509_cf42_, (self).f27, d_1513_v138_, d_1509_cf42_)
                d_1515_v139_ = nw273_
                d_1516_v140_: _dafny.Array
                nw274_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_1516_v140_ = nw274_
                d_1516_v140_ = d_1516_v140_
                d_1517_v141_: _dafny.Map
                d_1517_v141_ = _dafny.Map({False: d_1509_cf42_})
                d_1518_v142_: D14
                d_1518_v142_ = D14_DC32((self).f25, (self).f26, d_1517_v141_, _dafny.SeqWithoutIsStrInference([d_1388_v39_]), (self).f25)
                d_1519_v143_: C1
                nw275_ = C1()
                nw275_.ctor__((self).f25)
                d_1519_v143_ = nw275_
                d_1520_v144_: D10
                d_1520_v144_ = D10_DC24((d_1518_v142_).cf53, d_1519_v143_, d_1388_v39_)
                (globalState).f2 = (d_1520_v144_).cf36
            elif source21_.is_DC25:
                d_1521___mcc_h7_ = source21_.cf39
                d_1522_cf39_ = d_1521___mcc_h7_
                d_1523_v145_: _dafny.Array
                nw276_ = _dafny.Array(int(0), 11)
                d_1523_v145_ = nw276_
                rhs287_ = (self).f27
                rhs288_ = d_1523_v145_
                rhs289_ = (self).f25
                lhs244_ = globalState
                lhs245_ = globalState
                lhs244_.f2 = rhs287_
                d_1523_v145_ = rhs288_
                lhs245_.f1 = rhs289_
                d_1524_v146_: _dafny.Array
                def lambda88_(d_1525_i16_):
                    return not((self).f25)

                init50_ = lambda88_
                nw277_ = _dafny.Array(None, 10)
                for i0_50_ in range(nw277_.length(0)):
                    nw277_[i0_50_] = init50_(i0_50_)
                d_1524_v146_ = nw277_
                index279_ = default__.safeIndex(374, (d_1524_v146_).length(0))
                (d_1524_v146_)[index279_] = not((self).f25)
                index280_ = default__.safeIndex(374, (d_1524_v146_).length(0))
                (d_1524_v146_)[index280_] = (self).f25
                d_1526_v147_: _dafny.Array
                nw278_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_1526_v147_ = nw278_
                d_1527_v148_: D15
                d_1527_v148_ = D15_DC33(d_1526_v147_)
                d_1526_v147_ = ((d_1527_v148_ if (d_1524_v146_)[default__.safeIndex(374, (d_1524_v146_).length(0))] else D15_DC33(d_1526_v147_))).cf57
                (globalState).f2 = (self).f27
            elif True:
                d_1528___mcc_h8_ = source21_.cf43
                d_1529_cf43_ = d_1528___mcc_h8_
                (globalState).f1 = (self).f25
                d_1530_v149_: _dafny.Array
                nw279_ = _dafny.Array(int(0), 22)
                d_1530_v149_ = nw279_
                index281_ = default__.safeIndex(473, (d_1530_v149_).length(0))
                (d_1530_v149_)[index281_] = default__.safeModuloInt((d_1495_v126_).cardinality, default__.fm2((self).f27, (self).f27, globalState))
                d_1531_v150_: _dafny.Map
                d_1531_v150_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1388_v39_ for d_1532_i17_ in range(default__.abs(277))])): (self).f27})
                d_1533_v151_: _dafny.Map
                d_1533_v151_ = _dafny.Map({d_1531_v150_: (d_1498_v129_).cardinality})
                index282_ = default__.safeIndex(473, (d_1530_v149_).length(0))
                (d_1530_v149_)[index282_] = default__.safeDivisionInt(len((d_1533_v151_) | (d_1533_v151_)), default__.fm2((self).f27, (self).f26, globalState))
                r2 = ((self).f27) * ((self).f27)
                (globalState).f2 = default__.safeModuloInt((self).f27, (self).f26)
            d_1534_v153_: _dafny.Map
            def iife117_():
                coll33_ = _dafny.Set()
                compr_33_: int
                for compr_33_ in _dafny.IntegerRange(881, 80):
                    d_1535_v152_: int = compr_33_
                    if ((881) <= (d_1535_v152_)) and ((d_1535_v152_) < (80)):
                        coll33_ = coll33_.union(_dafny.Set([default__.safeModuloInt(d_1535_v152_, 328)]))
                return _dafny.Set(coll33_)
            d_1534_v153_ = _dafny.Map({(self).f26: iife117_()
            })
            d_1536_v154_: D2
            d_1536_v154_ = D2_DC6(((d_1534_v153_)[(self).f27] if ((self).f27) in (d_1534_v153_) else d_1500_v131_))
            d_1537_v155_: _dafny.Map
            d_1537_v155_ = _dafny.Map({len(d_1500_v131_): (self).f26})
            d_1538_v156_: _dafny.Map
            d_1538_v156_ = _dafny.Map({(self).f25: (self).f25})
            d_1539_v157_: D14
            d_1539_v157_ = D14_DC32((self).f25, (self).f26, d_1538_v156_, self.f28, False)
            pat_let_tv35_ = d_1538_v156_
            def iife118_(_pat_let42_0):
                def iife119_(d_1540_dt__update__tmp_h1_):
                    def iife120_(_pat_let43_0):
                        def iife121_(d_1541_dt__update_hcf10_h0_):
                            return D2_DC6(d_1541_dt__update_hcf10_h0_)
                        return iife121_(_pat_let43_0)
                    return iife120_((_dafny.Set({(self).f26, (self).f27, 700})) | (_dafny.Set({(self).f27})))
                return iife119_(_pat_let42_0)
            def iife122_(_pat_let44_0):
                def iife123_(d_1542_dt__update__tmp_h2_):
                    def iife124_(_pat_let45_0):
                        def iife125_(d_1543_dt__update_hcf54_h0_):
                            return D14_DC32((d_1542_dt__update__tmp_h2_).cf52, (d_1542_dt__update__tmp_h2_).cf53, d_1543_dt__update_hcf54_h0_, (d_1542_dt__update__tmp_h2_).cf55, (d_1542_dt__update__tmp_h2_).cf56)
                        return iife125_(_pat_let45_0)
                    return iife124_(pat_let_tv35_)
                return iife123_(_pat_let44_0)
            rhs290_ = (self).f25
            rhs291_ = iife118_(default__.fm38(((d_1497_v128_)[(self).f27] if ((self).f27) in (d_1497_v128_) else (self).f25), (self).f26, (0) - (len(d_1537_v155_)), default__.fm5((self).f25, _dafny.Map({D0_DC0((self).f25): (self).f27}), (self).fm6((self).f25, not((self).f25), 220, globalState), globalState), globalState))
            rhs292_ = default__.safeModuloInt(((d_1496_v127_)[True] if (True) in (d_1496_v127_) else (iife122_(d_1539_v157_)).cf53), (self).f26)
            rhs293_ = (self).f27
            lhs246_ = globalState
            lhs247_ = globalState
            lhs248_ = globalState
            lhs246_.f1 = rhs290_
            d_1536_v154_ = rhs291_
            lhs247_.f2 = rhs292_
            lhs248_.f2 = rhs293_
            d_1544_v158_: _dafny.Array
            def lambda89_(d_1545_i18_):
                return True

            init51_ = lambda89_
            nw280_ = _dafny.Array(None, 17)
            for i0_51_ in range(nw280_.length(0)):
                nw280_[i0_51_] = init51_(i0_51_)
            d_1544_v158_ = nw280_
            d_1546_v159_: C0
            nw281_ = C0()
            nw281_.ctor__(self.f28, ((self).f25) and ((self).f25), (self).f27, d_1544_v158_, (self).f25)
            d_1546_v159_ = nw281_
        elif True:
            d_1547_v160_: _dafny.Map
            d_1547_v160_ = _dafny.Map({(self).f26: len(_dafny.SeqWithoutIsStrInference([(self).f26]))})
            d_1548_v161_: _dafny.MultiSet
            d_1548_v161_ = _dafny.MultiSet([self.f28])
            d_1549_v162_: _dafny.Array
            def lambda90_(d_1550_i19_):
                return (self).f25

            init52_ = lambda90_
            nw282_ = _dafny.Array(None, 14)
            for i0_52_ in range(nw282_.length(0)):
                nw282_[i0_52_] = init52_(i0_52_)
            d_1549_v162_ = nw282_
            d_1551_v163_: C0
            nw283_ = C0()
            nw283_.ctor__(self.f28, (self).f25, (self).f27, d_1549_v162_, (self).f25)
            d_1551_v163_ = nw283_
            d_1552_v164_: _dafny.Map
            d_1552_v164_ = _dafny.Map({d_1551_v163_: (self).f25})
            d_1553_v165_: D0
            d_1553_v165_ = D0_DC1(39, d_1551_v163_.f35, False, (_dafny.MultiSet([(self).f27, 434, (self).f26])).cardinality, (self).f25)
            d_1554_v166_: C3
            nw284_ = C3()
            nw284_.ctor__(_dafny.SeqWithoutIsStrInference([d_1388_v39_ for d_1555_i20_ in range(default__.abs(277))]), (self).f25, (self).f25, (self).f26)
            d_1554_v166_ = nw284_
            d_1556_v169_: D16
            d_1556_v169_ = D16_DC35(d_1547_v160_)
            d_1557_v172_: _dafny.MultiSet
            d_1557_v172_ = _dafny.MultiSet([(self).f26])
            d_1558_v174_: _dafny.Seq
            d_1558_v174_ = _dafny.SeqWithoutIsStrInference([d_1547_v160_, d_1547_v160_])
            d_1559_v175_: _dafny.Map
            d_1559_v175_ = _dafny.Map({True: d_1388_v39_})
            d_1560_v176_: _dafny.Map
            d_1560_v176_ = _dafny.Map({len(d_1559_v175_): self.f28})
            d_1561_v177_: _dafny.Array
            nw285_ = _dafny.Array(None, 29)
            nw285_[int(0)] = d_1547_v160_
            nw285_[int(1)] = _dafny.Map({(self).f26: ((d_1548_v161_)[(self.f28).set(default__.safeIndex(len(d_1552_v164_), len(self.f28)), default__.fm4(d_1553_v165_, (self).f25, len(_dafny.Map({d_1554_v166_: (d_1554_v166_).f30})), globalState))] if ((self.f28).set(default__.safeIndex(len(d_1552_v164_), len(self.f28)), default__.fm4(d_1553_v165_, (self).f25, len(_dafny.Map({d_1554_v166_: (d_1554_v166_).f30})), globalState))) in (d_1548_v161_) else (self).f26)})
            def iife126_():
                coll34_ = _dafny.Map()
                compr_34_: int
                for compr_34_ in _dafny.IntegerRange(83, 340):
                    d_1562_v167_: int = compr_34_
                    if ((83) <= (d_1562_v167_)) and ((d_1562_v167_) < (340)):
                        coll34_[(d_1562_v167_) - ((self).f27)] = (self).f27
                return _dafny.Map(coll34_)
            nw285_[int(2)] = (iife126_()
            ) | (d_1547_v160_)
            nw285_[int(3)] = (d_1547_v160_) | (d_1547_v160_)
            nw285_[int(4)] = (_dafny.Map({(self).f27: (self).f27})).set((self).f27, (self).f26)
            nw285_[int(5)] = d_1547_v160_
            nw285_[int(6)] = (_dafny.Map({(0) - ((self).f26): (self).fm6((d_1554_v166_).f30, not((d_1554_v166_).f30), (self).f26, globalState)})) | (d_1547_v160_)
            nw285_[int(7)] = d_1547_v160_
            nw285_[int(8)] = _dafny.Map({(self).f27: (self).f27})
            nw285_[int(9)] = d_1547_v160_
            nw285_[int(10)] = (d_1547_v160_) | (d_1547_v160_)
            def iife127_():
                coll35_ = _dafny.Map()
                compr_35_: int
                for compr_35_ in _dafny.IntegerRange(491, -882):
                    d_1563_v168_: int = compr_35_
                    if ((491) <= (d_1563_v168_)) and ((d_1563_v168_) < (-882)):
                        coll35_[(d_1563_v168_) + ((_dafny.MultiSet((d_1554_v166_).f29)).cardinality)] = 343
                return _dafny.Map(coll35_)
            nw285_[int(11)] = iife127_()
            
            nw285_[int(12)] = d_1547_v160_
            nw285_[int(13)] = (d_1556_v169_).cf59
            def iife128_():
                coll36_ = _dafny.Map()
                compr_36_: int
                for compr_36_ in _dafny.IntegerRange(580, 30):
                    d_1564_v170_: int = compr_36_
                    if ((580) <= (d_1564_v170_)) and ((d_1564_v170_) < (30)):
                        coll36_[(d_1564_v170_) - (len(_dafny.Map({(self).f27: (d_1554_v166_).f30})))] = (self).f26
                return _dafny.Map(coll36_)
            nw285_[int(14)] = (iife128_()
            ) | (d_1547_v160_)
            def iife129_():
                coll37_ = _dafny.Map()
                compr_37_: int
                for compr_37_ in (d_1557_v172_).Elements:
                    d_1565_v171_: int = compr_37_
                    if (d_1565_v171_) in (d_1557_v172_):
                        coll37_[default__.safeModuloInt(d_1565_v171_, (self).f27)] = (self).f26
                return _dafny.Map(coll37_)
            nw285_[int(15)] = (iife129_()
            ) | (d_1547_v160_)
            def iife130_():
                coll38_ = _dafny.Map()
                compr_38_: int
                for compr_38_ in _dafny.IntegerRange(77, -552):
                    d_1566_v173_: int = compr_38_
                    if ((77) <= (d_1566_v173_)) and ((d_1566_v173_) < (-552)):
                        coll38_[(d_1566_v173_) + (-918)] = (self).f26
                return _dafny.Map(coll38_)
            nw285_[int(16)] = iife130_()
            
            nw285_[int(17)] = (d_1558_v174_)[default__.safeIndex((self).f27, len(d_1558_v174_))]
            nw285_[int(18)] = (d_1547_v160_ if (self).f25 else d_1547_v160_)
            nw285_[int(19)] = d_1547_v160_
            nw285_[int(20)] = _dafny.Map({(self).f26: (self).f27})
            nw285_[int(21)] = d_1547_v160_
            nw285_[int(22)] = d_1547_v160_
            nw285_[int(23)] = (d_1547_v160_) | (d_1547_v160_)
            nw285_[int(24)] = (d_1547_v160_) | (d_1547_v160_)
            nw285_[int(25)] = d_1547_v160_
            nw285_[int(26)] = (d_1547_v160_).set(159, len(d_1560_v176_))
            nw285_[int(27)] = (d_1547_v160_) | (d_1547_v160_)
            nw285_[int(28)] = d_1547_v160_
            d_1561_v177_ = nw285_
            nw286_ = _dafny.Array(_dafny.Map({}), 25)
            d_1561_v177_ = nw286_
            d_1567_v178_: _dafny.Array
            def lambda91_(d_1568_i21_):
                return default__.safeModuloInt(d_1568_i21_, (self).f26)

            init53_ = lambda91_
            nw287_ = _dafny.Array(None, 13)
            for i0_53_ in range(nw287_.length(0)):
                nw287_[i0_53_] = init53_(i0_53_)
            d_1567_v178_ = nw287_
            d_1567_v178_ = (d_1567_v178_ if (d_1554_v166_).f30 else d_1567_v178_)
            r1 = (d_1554_v166_).f30
            d_1388_v39_ = d_1388_v39_
            d_1569_v179_: _dafny.Array
            nw288_ = _dafny.Array(D0.default()(), 8)
            d_1569_v179_ = nw288_
            d_1570_v180_: _dafny.Set
            d_1570_v180_ = _dafny.Set({d_1569_v179_})
            (globalState).f1 = (d_1570_v180_).ispropersubset(_dafny.Set({d_1569_v179_, d_1569_v179_}))
        d_1571_v181_: _dafny.Array
        nw289_ = _dafny.Array(False, 29)
        d_1571_v181_ = nw289_
        r0 = d_1571_v181_
        d_1572_v182_: _dafny.Seq
        d_1572_v182_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_1573_v183_: _dafny.MultiSet
        d_1573_v183_ = _dafny.MultiSet([(self).f27])
        d_1574_v184_: _dafny.MultiSet
        d_1574_v184_ = _dafny.MultiSet([d_1572_v182_, _dafny.SeqWithoutIsStrInference([(self).f25])])
        r1 = not(((d_1572_v182_).set(default__.safeIndex(default__.fm2((self).f27, ((d_1573_v183_)[len(d_1572_v182_)] if (len(d_1572_v182_)) in (d_1573_v183_) else (self).f26), globalState), len(d_1572_v182_)), (self).f25)) in (d_1574_v184_))
        r2 = default__.safeDivisionInt(((self).f26 if (self).f25 else (self).f26), (self).f27)
        return r0, r1, r2

    @property
    def f27(self):
        return self._f27
