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
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(602, 705):
                d_0_v0_: int = compr_0_
                if ((602) <= (d_0_v0_)) and ((d_0_v0_) < (705)):
                    coll0_[default__.safeDivisionInt(d_0_v0_, 803)] = _dafny.SeqWithoutIsStrInference([-786])
            return _dafny.Map(coll0_)
        return (default__.safeDivisionInt(len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([984]))})), len(iife0_()
        ))) - ((376) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsmqogqg")))))

    @staticmethod
    def fm1(p0, p1, globalState):
        return ((_dafny.MultiSet([(_dafny.MultiSet([32])).cardinality])).intersection(_dafny.MultiSet([756]))).intersection(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([43, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({874}))]))])) + (_dafny.SeqWithoutIsStrInference([268, 281]))))

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        return (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1_i0_ in range(default__.abs(714))]))) > (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hidhmds"))), len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b')]))).cardinality for d_2_i1_ in range(default__.abs(427))]))))

    @staticmethod
    def fm3(globalState):
        return _dafny.CodePoint('v')

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return ((D8_DC26(D1_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('o'), _dafny.CodePoint('r'), _dafny.CodePoint('x')])), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cu")))).cf42) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccsyvjc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdcttheo"))))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(202, -405):
                d_3_v0_: int = compr_1_
                if ((202) <= (d_3_v0_)) and ((d_3_v0_) < (-405)):
                    coll1_[default__.safeModuloInt(d_3_v0_, len(_dafny.Map({684: True})))] = not(True)
            return _dafny.Map(coll1_)
        return (_dafny.Map({129: 577})) | ((_dafny.Map({711: len(_dafny.Map({False: len((D8_DC23(_dafny.Set({True}))).cf36)}))})) | (_dafny.Map({(0) - (len(iife1_()
        )): len((D19_DC44(_dafny.Map({not(True): False}))).cf67)})))

    @staticmethod
    def fm7(globalState):
        return _dafny.SeqWithoutIsStrInference([-417])

    @staticmethod
    def fm13(globalState):
        return (_dafny.Map({False: True})) | ((_dafny.Map({True: False})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm14(p0, p1, globalState):
        return (_dafny.MultiSet([False])) | ((_dafny.MultiSet([True, False])) | (_dafny.MultiSet([True, True, True, not(not(not(False))), True])))

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        return _dafny.Map({((0) - ((_dafny.MultiSet([155])).cardinality)) * (len(_dafny.SeqWithoutIsStrInference([True]))): (len(_dafny.SeqWithoutIsStrInference([True]))) - ((_dafny.MultiSet([_dafny.CodePoint('n'), _dafny.CodePoint('c')])).cardinality)})

    @staticmethod
    def fm16(globalState):
        if (len(_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnbdljh")))])}))) < (286):
            return (_dafny.Map({-275: False})) | (_dafny.Map({495: not(False)}))
        elif True:
            return (_dafny.Map({-59: not(False)})) | (_dafny.Map({(0) - ((0) - ((_dafny.MultiSet([_dafny.CodePoint('p')])).cardinality)): True}))

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([True])) == (_dafny.SeqWithoutIsStrInference([not(True), True, True])): (-570) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sx"))))})

    @staticmethod
    def fm18(globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(418, -290):
                d_4_v0_: int = compr_2_
                if ((418) <= (d_4_v0_)) and ((d_4_v0_) < (-290)):
                    coll2_ = coll2_.union(_dafny.Set([(d_4_v0_) + ((_dafny.MultiSet([891, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_5_i0_ in range(default__.abs(661))]))])).cardinality)]))
            return _dafny.Set(coll2_)
        return (D6_DC16(iife2_()
)).cf27

    @staticmethod
    def fm19(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([not(True), True])) + (_dafny.SeqWithoutIsStrInference([False, True])), _dafny.SeqWithoutIsStrInference([True])])

    @staticmethod
    def fm20(globalState):
        return ((D20_DC48(_dafny.Map({_dafny.SeqWithoutIsStrInference([not(not(False)), False, not(False)]): 812})) if False else D20_DC48(_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): 930})))).cf75

    @staticmethod
    def fm21(globalState):
        return D0_DC1((0) - (default__.safeModuloInt(len(_dafny.Map({_dafny.CodePoint('n'): False})), len(_dafny.SeqWithoutIsStrInference([True, not(False)])))), (len(_dafny.Set({790}))) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_6_i0_ in range(default__.abs(494))]))), (D10_DC29(len(_dafny.Map({True: len(_dafny.Map({False: True}))})), False)).cf46)

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        return _dafny.Set({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_7_i0_ in range(default__.abs(309))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_8_i1_ in range(default__.abs(132))]))})

    @staticmethod
    def fm23(globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_9_i0_ in range(default__.abs(755))])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytfqd")))])

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([814 for d_10_i0_ in range(default__.abs(-428))])

    @staticmethod
    def fm25(globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(38, -926):
                d_11_v0_: int = compr_3_
                if ((38) <= (d_11_v0_)) and ((d_11_v0_) < (-926)):
                    coll3_ = coll3_.union(_dafny.Set([(d_11_v0_) + (960)]))
            return _dafny.Set(coll3_)
        return D4_DC13((994) - ((0) - (len(_dafny.Map({True: True})))), _dafny.Map({True: len(iife3_()
)}), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "folgy"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))))

    @staticmethod
    def fm26(globalState):
        return D4_DC14((D4_DC14(D4_DC14(D4_DC12(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality]))))).cf25)

    @staticmethod
    def fm27(globalState):
        return D1_DC6()

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({not(False), True, True, not(True)})) - (_dafny.Set({not((D1_DC4(True, False, _dafny.CodePoint('e'), 554, 784)).cf8), True, False, False}))) - (_dafny.Set({True, False, True, True, not(True)}))

    @staticmethod
    def fm29(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([not(False), False])), len(_dafny.Map({-424: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([398 for d_12_i0_ in range(default__.abs(643))]))).cardinality})), len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([58, 520])): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xb"))))}): 311}))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([289]))])]))) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([19, (_dafny.MultiSet([True])).cardinality, 32, 765]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_13_i2_ in range(default__.abs(-402))])) for d_14_i1_ in range(default__.abs(684))])), len(_dafny.Set({True, False})), -421])]))

    @staticmethod
    def fm30(p0, p1, globalState):
        def iife4_():
            def iife6_():
                coll6_ = _dafny.Set()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(232, 455):
                    d_17_v1_: int = compr_6_
                    if ((232) <= (d_17_v1_)) and ((d_17_v1_) < (455)):
                        coll6_ = coll6_.union(_dafny.Set([default__.safeModuloInt(d_17_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))))]))
                return _dafny.Set(coll6_)
            coll4_ = _dafny.Map()
            def iife5_():
                coll5_ = _dafny.Set()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(232, 455):
                    d_15_v1_: int = compr_5_
                    if ((232) <= (d_15_v1_)) and ((d_15_v1_) < (455)):
                        coll5_ = coll5_.union(_dafny.Set([default__.safeModuloInt(d_15_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))))]))
                return _dafny.Set(coll5_)
            compr_4_: int
            for compr_4_ in (iife5_()
            ).Elements:
                d_16_v0_: int = compr_4_
                if (d_16_v0_) in (iife6_()
                ):
                    coll4_[default__.safeDivisionInt(d_16_v0_, 845)] = _dafny.CodePoint('s')
            return _dafny.Map(coll4_)
        return (iife4_()
        ) | ((_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_18_i0_ in range(default__.abs(393))]))): _dafny.CodePoint('e')})) | (_dafny.Map({695: _dafny.CodePoint('k')})))

    @staticmethod
    def fm31(p0, p1, globalState):
        return D6_DC19(D6_DC19(D6_DC18(False)))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return (_dafny.Map({D1_DC4(False, True, _dafny.CodePoint('r'), -724, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_19_i0_ in range(default__.abs(-318))]))): not(False)})) | (_dafny.Map({D1_DC4(True, not(False), _dafny.CodePoint('l'), -368, -261): False}))

    @staticmethod
    def fm33(p0, p1, globalState):
        return D3_DC10()

    @staticmethod
    def fm34(p0, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: _dafny.Seq
            for compr_7_ in ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqy"))})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjuydaqli"))}))).Elements:
                d_20_v0_: _dafny.Seq = compr_7_
                if (d_20_v0_) in ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqy"))})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjuydaqli"))}))):
                    coll7_[d_20_v0_] = D1_DC5(not(False), not(True), not(True))
            return _dafny.Map(coll7_)
        return iife7_()
        

    @staticmethod
    def fm35(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Set({(0) - (len(_dafny.Set({51}))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvhgqeka")))}) for d_21_i0_ in range(default__.abs(975))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([True]))), -183, (0) - (-324), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, False]), _dafny.MultiSet([not(True), True])]))])), -687}) for d_22_i1_ in range(default__.abs(348))]))

    @staticmethod
    def fm36(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: True})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(D12_DC33(_dafny.CodePoint('r'), not(False))).cf55: False}) for d_23_i0_ in range(default__.abs(-904))]))

    @staticmethod
    def fm37(globalState):
        if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycqwr"))) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_24_i0_ in range(default__.abs(737))])):
            return D8_DC26(D1_DC3((D8_DC26(D1_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('b'), _dafny.CodePoint('r')])), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndi")))).cf42), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))
        elif True:
            return D8_DC26(D1_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_25_i1_ in range(default__.abs(693))])), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_26_i2_ in range(default__.abs(263))]))

    @staticmethod
    def fm38(globalState):
        return _dafny.MultiSet([D3_DC11(442, False, True)])

    @staticmethod
    def fm39(p0, p1, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(80, 406):
                d_27_v0_: int = compr_8_
                if ((80) <= (d_27_v0_)) and ((d_27_v0_) < (406)):
                    coll8_[(d_27_v0_) - (-269)] = _dafny.Set({522, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gc")))})
            return _dafny.Map(coll8_)
        return (iife8_()
        ) | ((_dafny.Map({977: _dafny.Set({826})}) if False else _dafny.Map({len(_dafny.Set({True})): _dafny.Set({380})})))

    @staticmethod
    def fm40(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")) for d_28_i0_ in range(default__.abs(828))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fil"))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "canhopsk"))]))

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in (_dafny.Map({709: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjwyqqi")))]))})).keys.Elements:
                d_29_v0_: int = compr_9_
                if (d_29_v0_) in (_dafny.Map({709: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjwyqqi")))]))})):
                    coll9_[default__.safeModuloInt(d_29_v0_, 620)] = False
            return _dafny.Map(coll9_)
        return _dafny.Set({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nw")))])) - ((D13_DC35(_dafny.MultiSet([98]))).cf57), _dafny.MultiSet([(0) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s')])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcleahc")))])).cardinality)]), (_dafny.MultiSet([len(_dafny.Map({(0) - (-165): len(_dafny.Set({True, True, True}))}))])).intersection(_dafny.MultiSet([len(iife9_()
        ), -251, len(_dafny.SeqWithoutIsStrInference([not(True)])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqgceiacv")))]))})

    @staticmethod
    def fm42(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('f'): -853}), _dafny.Map({_dafny.CodePoint('x'): len(_dafny.Set({True}))})])

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        return D1_DC4(not (False) or (False), False, _dafny.CodePoint('i'), len((_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, not(not(True))])) for d_30_i0_ in range(default__.abs(637))])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([439])}))), default__.safeModuloInt(124, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahdj")))])).cardinality))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: int = int(0)
        d_31_v0_: _dafny.Map
        d_31_v0_ = _dafny.Map({p1: 4})
        d_32_v1_: _dafny.Seq
        d_32_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfp"))
        d_31_v0_ = (d_31_v0_).set(len(d_32_v1_), p1)
        d_33_v2_: _dafny.MultiSet
        d_33_v2_ = _dafny.MultiSet([p1])
        d_34_v3_: _dafny.Seq
        d_34_v3_ = _dafny.SeqWithoutIsStrInference([(d_33_v2_).set(len(d_31_v0_), default__.abs(default__.fm0(p0, p0, globalState)))])
        d_35_v4_: bool
        d_35_v4_ = False
        d_36_v5_: _dafny.Set
        d_36_v5_ = _dafny.Set({d_35_v4_, d_35_v4_, d_35_v4_, True})
        d_37_v6_: str
        d_37_v6_ = _dafny.CodePoint('w')
        d_38_v7_: _dafny.Map
        d_38_v7_ = _dafny.Map({d_37_v6_: p1})
        d_39_v8_: _dafny.Seq
        d_39_v8_ = _dafny.SeqWithoutIsStrInference([d_38_v7_])
        d_40_v9_: T0
        nw0_ = C1()
        nw0_.ctor__(d_34_v3_, p1, len(d_36_v5_), d_39_v8_)
        d_40_v9_ = nw0_
        d_41_v10_: _dafny.Map
        d_41_v10_ = _dafny.Map({p1: d_40_v9_})
        d_42_v11_: _dafny.Array
        nw1_ = _dafny.Array(int(0), 14)
        d_42_v11_ = nw1_
        d_43_v12_: C4
        nw2_ = C4()
        nw2_.ctor__(d_42_v11_, (d_40_v9_).f24, (d_40_v9_).f25)
        d_43_v12_ = nw2_
        d_44_v13_: _dafny.MultiSet
        d_44_v13_ = _dafny.MultiSet([d_43_v12_, d_43_v12_, d_43_v12_, d_43_v12_, d_43_v12_])
        d_45_v14_: T0
        d_45_v14_ = d_40_v9_
        d_41_v10_ = (d_41_v10_).set(default__.fm0(((d_44_v13_)[d_43_v12_] if (d_43_v12_) in (d_44_v13_) else 235), 986, globalState), (d_40_v9_))
        d_46_v15_: _dafny.Map
        d_46_v15_ = _dafny.Map({(d_43_v12_).f28: (d_40_v9_).f24})
        index0_ = default__.safeIndex(862, ((d_43_v12_).f28).length(0))
        ((d_43_v12_).f28)[index0_] = len(d_46_v15_)
        index1_ = default__.safeIndex(862, ((d_43_v12_).f28).length(0))
        ((d_43_v12_).f28)[index1_] = p1
        d_47_v16_: _dafny.Map
        d_47_v16_ = _dafny.Map({d_35_v4_: (d_43_v12_).f28})
        d_47_v16_ = (d_47_v16_) | ((d_47_v16_) | (d_47_v16_))
        hi0_ = len(_dafny.SeqWithoutIsStrInference([d_37_v6_ for d_48_i1_ in range(default__.abs(-669))]))
        for d_49_i0_ in range((0) - (((d_43_v12_).f28)[default__.safeIndex(862, ((d_43_v12_).f28).length(0))]), hi0_):
            d_50_v17_: _dafny.Array
            def lambda0_(d_51_v1_):
                def lambda1_(d_52_i2_):
                    return (_dafny.CodePoint('w')) not in (d_51_v1_)

                return lambda1_

            init0_ = lambda0_(d_32_v1_)
            nw3_ = _dafny.Array(None, 15)
            for i0_0_ in range(nw3_.length(0)):
                nw3_[i0_0_] = init0_(i0_0_)
            d_50_v17_ = nw3_
            d_53_v18_: _dafny.Map
            d_53_v18_ = _dafny.Map({(d_40_v9_).f24: (d_50_v17_ if True else d_50_v17_)})
            d_50_v17_ = ((d_53_v18_)[(d_40_v9_).f24] if ((d_40_v9_).f24) in (d_53_v18_) else d_50_v17_)
            d_54_v19_: _dafny.MultiSet
            d_54_v19_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_37_v6_ for d_55_i3_ in range(default__.abs(694))])])
            d_56_v20_: _dafny.Seq
            d_56_v20_ = _dafny.SeqWithoutIsStrInference([d_49_i0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipsloox")))])
            d_57_v21_: _dafny.Set
            d_57_v21_ = _dafny.Set({(d_54_v19_).cardinality, len(d_56_v20_)})
            d_58_v22_: D10
            d_58_v22_ = D10_DC29((d_40_v9_).f24, d_35_v4_)
            (globalState).f18 = ((((d_43_v12_).f28)[default__.safeIndex(862, ((d_43_v12_).f28).length(0))]) - (len(d_57_v21_)) if d_35_v4_ else default__.safeModuloInt((d_58_v22_).cf45, ((d_43_v12_).f28)[default__.safeIndex(862, ((d_43_v12_).f28).length(0))]))
            d_59_v23_: _dafny.Array
            nw4_ = _dafny.Array(_dafny.Seq({}), 27)
            d_59_v23_ = nw4_
            d_60_v24_: _dafny.Seq
            d_60_v24_ = _dafny.SeqWithoutIsStrInference([d_35_v4_])
            index2_ = default__.safeIndex(349, (d_59_v23_).length(0))
            (d_59_v23_)[index2_] = d_60_v24_
            index3_ = default__.safeIndex(349, (d_59_v23_).length(0))
            rhs0_ = (d_60_v24_) + ((_dafny.SeqWithoutIsStrInference([d_35_v4_, d_35_v4_, d_35_v4_, d_35_v4_])).set(default__.safeIndex(d_49_i0_, len(_dafny.SeqWithoutIsStrInference([d_35_v4_, d_35_v4_, d_35_v4_, d_35_v4_]))), d_35_v4_))
            rhs1_ = (0) - (((d_40_v9_).f24) + (p0))
            rhs2_ = d_36_v5_
            lhs0_ = d_59_v23_
            lhs1_ = default__.safeIndex(349, (d_59_v23_).length(0))
            lhs2_ = globalState
            lhs0_[lhs1_] = rhs0_
            lhs2_.f1 = rhs1_
            d_36_v5_ = rhs2_
            d_61_v25_: _dafny.MultiSet
            d_61_v25_ = _dafny.MultiSet([d_35_v4_])
            d_62_v26_: _dafny.Map
            d_62_v26_ = _dafny.Map({(d_61_v25_) - (d_61_v25_): False})
            d_62_v26_ = (d_62_v26_).set(d_61_v25_, (d_35_v4_) == (True))
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_42_v11_).length(0)):
            d_63_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_63_i4_)) and ((d_63_i4_) < ((d_42_v11_).length(0)))):
                (d_42_v11_)[(d_63_i4_)] = (d_63_i4_) * (p1)
        d_64_v27_: _dafny.Map
        d_64_v27_ = _dafny.Map({d_35_v4_: p0})
        d_65_v28_: D4
        d_65_v28_ = D4_DC13(-145, d_64_v27_, d_32_v1_)
        r0 = (0) - (len((d_65_v28_).cf24))
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_66_v0_: int
        d_66_v0_ = 590
        d_67_v1_: _dafny.Map
        d_67_v1_ = _dafny.Map({False: d_66_v0_})
        d_68_v2_: bool
        d_68_v2_ = True
        d_69_v3_: _dafny.Array
        def lambda2_(d_70_i0_):
            return (d_70_i0_) - (-719)

        init1_ = lambda2_
        nw5_ = _dafny.Array(None, 24)
        for i0_1_ in range(nw5_.length(0)):
            nw5_[i0_1_] = init1_(i0_1_)
        d_69_v3_ = nw5_
        d_71_v4_: D0
        d_71_v4_ = D0_DC0(d_69_v3_)
        d_72_v5_: _dafny.Map
        d_72_v5_ = _dafny.Map({d_69_v3_: (d_71_v4_).cf0})
        d_73_v6_: _dafny.Seq
        d_73_v6_ = _dafny.SeqWithoutIsStrInference([d_68_v2_])
        d_74_v7_: str
        d_74_v7_ = _dafny.CodePoint('c')
        d_75_v8_: D1
        d_75_v8_ = D1_DC3(_dafny.SeqWithoutIsStrInference([d_74_v7_, _dafny.CodePoint('h')]))
        d_76_v9_: _dafny.MultiSet
        d_76_v9_ = _dafny.MultiSet([d_74_v7_])
        d_77_v10_: _dafny.Set
        d_77_v10_ = _dafny.Set({_dafny.MultiSet((d_75_v8_).cf7), d_76_v9_})
        d_78_v11_: _dafny.Array
        nw6_ = _dafny.Array(False, 2)
        d_78_v11_ = nw6_
        d_79_v12_: _dafny.Map
        d_79_v12_ = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([d_66_v0_ for d_80_i1_ in range(default__.abs(-991))]))): d_78_v11_})
        d_81_v13_: _dafny.Set
        d_81_v13_ = _dafny.Set({d_68_v2_, d_68_v2_, d_68_v2_})
        d_82_v14_: _dafny.Seq
        d_82_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh"))
        d_83_v15_: D1
        d_83_v15_ = D1_DC4(d_68_v2_, d_68_v2_, (d_82_v14_)[default__.safeIndex(d_66_v0_, len(d_82_v14_))], d_66_v0_, d_66_v0_)
        d_84_v16_: _dafny.Map
        d_84_v16_ = _dafny.Map({d_66_v0_: _dafny.Map({382: d_66_v0_})})
        d_85_globalState_: GlobalState
        nw7_ = GlobalState()
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(199, 972):
                d_86_v17_: int = compr_10_
                if ((199) <= (d_86_v17_)) and ((d_86_v17_) < (972)):
                    coll10_[(d_86_v17_) - (d_66_v0_)] = d_66_v0_
            return _dafny.Map(coll10_)
        nw7_.ctor__(True, 554, True, (d_67_v1_).set(d_68_v2_, d_66_v0_), False, -61, _dafny.CodePoint('c'), d_72_v5_, ((_dafny.SeqWithoutIsStrInference([d_68_v2_, d_68_v2_, not(False)])) + (d_73_v6_)).set(default__.safeIndex((0) - (d_66_v0_), len((_dafny.SeqWithoutIsStrInference([d_68_v2_, d_68_v2_, not(False)])) + (d_73_v6_))), not(d_68_v2_)), 889, 204, d_77_v10_, 127, d_79_v12_, 996, 982, d_81_v13_, d_82_v14_, -607, _dafny.MultiSet([d_78_v11_, d_78_v11_]), d_73_v6_, _dafny.SeqWithoutIsStrInference([d_68_v2_, (d_83_v15_).cf8]), ((d_84_v16_)[d_66_v0_] if (d_66_v0_) in (d_84_v16_) else iife10_()
        ), 790)
        d_85_globalState_ = nw7_
        index4_ = default__.safeIndex(614, (d_78_v11_).length(0))
        (d_78_v11_)[index4_] = (d_66_v0_) <= (d_66_v0_)
        index5_ = default__.safeIndex(614, (d_78_v11_).length(0))
        (d_78_v11_)[index5_] = d_68_v2_
        d_74_v7_ = _dafny.CodePoint('u')
        d_87_v18_: _dafny.MultiSet
        d_87_v18_ = _dafny.MultiSet([not((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]), (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]])
        d_88_v19_: _dafny.Map
        d_88_v19_ = _dafny.Map({d_68_v2_: False})
        d_89_v20_: _dafny.Seq
        d_89_v20_ = _dafny.SeqWithoutIsStrInference([d_66_v0_, len(d_88_v19_)])
        if (default__.safeDivisionInt(d_66_v0_, ((d_87_v18_)[not((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))])] if (not((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))])) in (d_87_v18_) else d_66_v0_))) >= (default__.fm0(len(d_89_v20_), d_66_v0_, d_85_globalState_)):
            d_90_v21_: _dafny.MultiSet
            d_90_v21_ = _dafny.MultiSet([d_66_v0_])
            d_91_v22_: _dafny.Map
            d_91_v22_ = _dafny.Map({d_74_v7_: d_90_v21_})
            d_92_v23_: _dafny.Seq
            d_92_v23_ = _dafny.SeqWithoutIsStrInference([d_90_v21_])
            (d_85_globalState_).f14 = len((_dafny.SeqWithoutIsStrInference([d_90_v21_, default__.fm1(d_66_v0_, d_66_v0_, d_85_globalState_), ((d_91_v22_)[d_74_v7_] if (d_74_v7_) in (d_91_v22_) else d_90_v21_)])) + (d_92_v23_))
            d_71_v4_ = D0_DC0(d_69_v3_)
            d_93_v24_: _dafny.Set
            d_93_v24_ = _dafny.Set({default__.fm0(len(d_82_v14_), d_66_v0_, d_85_globalState_)})
            d_93_v24_ = d_93_v24_
            d_94_v25_: _dafny.Map
            d_94_v25_ = _dafny.Map({not(False): d_81_v13_})
            d_95_v26_: _dafny.Map
            d_95_v26_ = _dafny.Map({(d_81_v13_).ispropersubset(((d_94_v25_)[True] if (True) in (d_94_v25_) else _dafny.Set({d_68_v2_, default__.fm2(659, len(d_73_v6_), d_74_v7_, d_85_globalState_)}))): d_69_v3_})
            d_96_v27_: _dafny.Array
            nw8_ = _dafny.Array(_dafny.CodePoint('D'), 10)
            d_96_v27_ = nw8_
            index6_ = default__.safeIndex(707, (d_96_v27_).length(0))
            (d_96_v27_)[index6_] = (d_82_v14_)[default__.safeIndex(d_66_v0_, len(d_82_v14_))]
            d_97_v28_: _dafny.MultiSet
            d_97_v28_ = _dafny.MultiSet([d_96_v27_])
            index7_ = default__.safeIndex(707, (d_96_v27_).length(0))
            index8_ = default__.safeIndex(614, (d_78_v11_).length(0))
            rhs3_ = d_95_v26_
            rhs4_ = default__.fm3(d_85_globalState_)
            rhs5_ = d_66_v0_
            rhs6_ = ((d_97_v28_) | (d_97_v28_)).issubset(d_97_v28_)
            rhs7_ = d_66_v0_
            lhs3_ = d_96_v27_
            lhs4_ = default__.safeIndex(707, (d_96_v27_).length(0))
            lhs5_ = d_85_globalState_
            lhs6_ = d_78_v11_
            lhs7_ = default__.safeIndex(614, (d_78_v11_).length(0))
            lhs8_ = d_85_globalState_
            d_95_v26_ = rhs3_
            lhs3_[lhs4_] = rhs4_
            lhs5_.f9 = rhs5_
            lhs6_[lhs7_] = rhs6_
            lhs8_.f1 = rhs7_
            (d_85_globalState_).f9 = default__.safeModuloInt((d_66_v0_) + (765), len((d_81_v13_) - (d_81_v13_)))
        elif True:
            d_98_v29_: _dafny.Array
            nw9_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_98_v29_ = nw9_
            d_99_v30_: _dafny.Map
            d_99_v30_ = _dafny.Map({False: d_69_v3_})
            d_100_v31_: D0
            d_100_v31_ = D0_DC2((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))], (0) - (d_66_v0_), d_74_v7_)
            index9_ = default__.safeIndex(621, (d_98_v29_).length(0))
            (d_98_v29_)[index9_] = ((d_99_v30_)[not((d_100_v31_).cf4)] if (not((d_100_v31_).cf4)) in (d_99_v30_) else d_69_v3_)
            index10_ = default__.safeIndex(621, (d_98_v29_).length(0))
            (d_98_v29_)[index10_] = d_69_v3_
            d_101_v32_: D0
            d_101_v32_ = D0_DC1(215, d_66_v0_, (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))])
            d_102_v33_: _dafny.Map
            d_102_v33_ = _dafny.Map({d_66_v0_: d_82_v14_})
            d_103_v34_: _dafny.Array
            nw10_ = _dafny.Array(None, 21)
            nw10_[int(0)] = (d_82_v14_) + (d_82_v14_)
            nw10_[int(1)] = (d_82_v14_).set(default__.safeIndex(default__.fm0(d_66_v0_, d_66_v0_, d_85_globalState_), len(d_82_v14_)), d_74_v7_)
            nw10_[int(2)] = d_82_v14_
            nw10_[int(3)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orkpvq"))).set(default__.safeIndex(d_66_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orkpvq")))), default__.fm3(d_85_globalState_))
            nw10_[int(4)] = (d_82_v14_) + (default__.fm4((d_101_v32_).cf3, (0) - (len(d_89_v20_)), d_74_v7_, default__.fm2(d_66_v0_, d_66_v0_, d_74_v7_, d_85_globalState_), d_85_globalState_))
            nw10_[int(5)] = (d_82_v14_) + (d_82_v14_)
            nw10_[int(6)] = d_82_v14_
            nw10_[int(7)] = d_82_v14_
            nw10_[int(8)] = d_82_v14_
            nw10_[int(9)] = _dafny.SeqWithoutIsStrInference([d_74_v7_ for d_104_i2_ in range(default__.abs(-260))])
            nw10_[int(10)] = (((d_102_v33_)[len(d_81_v13_)] if (len(d_81_v13_)) in (d_102_v33_) else d_82_v14_)) + ((d_82_v14_).set(default__.safeIndex(len(d_89_v20_), len(d_82_v14_)), d_74_v7_))
            nw10_[int(11)] = d_82_v14_
            nw10_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmsoaylng"))
            nw10_[int(13)] = (d_82_v14_) + (d_82_v14_)
            nw10_[int(14)] = d_82_v14_
            nw10_[int(15)] = d_82_v14_
            nw10_[int(16)] = d_82_v14_
            nw10_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arypuvscv"))
            nw10_[int(18)] = d_82_v14_
            nw10_[int(19)] = (d_82_v14_).set(default__.safeIndex(d_66_v0_, len(d_82_v14_)), d_74_v7_)
            nw10_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmic"))
            d_103_v34_ = nw10_
            d_105_v35_: _dafny.Seq
            d_105_v35_ = _dafny.SeqWithoutIsStrInference([d_103_v34_, d_103_v34_])
            d_103_v34_ = (d_105_v35_)[default__.safeIndex((d_83_v15_).cf11, len(d_105_v35_))]
            (d_85_globalState_).f2 = (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]
            d_106_v36_: _dafny.Set
            d_106_v36_ = _dafny.Set({d_82_v14_, d_82_v14_, default__.fm4((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))], len(_dafny.Map({d_66_v0_: d_68_v2_})), d_74_v7_, d_68_v2_, d_85_globalState_)})
            d_107_v37_: _dafny.Map
            d_107_v37_ = _dafny.Map({(d_106_v36_) | (d_106_v36_): ((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]) or ((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))])})
            d_107_v37_ = (d_107_v37_).set(d_106_v36_, not((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]))
            d_108_v38_: _dafny.Map
            d_108_v38_ = _dafny.Map({d_66_v0_: d_74_v7_})
            d_68_v2_ = default__.fm2((d_66_v0_) + (d_66_v0_), default__.safeDivisionInt(d_66_v0_, d_66_v0_), ((d_108_v38_)[(d_101_v32_).cf1] if ((d_101_v32_).cf1) in (d_108_v38_) else d_74_v7_), d_85_globalState_)
        d_109_v39_: D1
        d_109_v39_ = D1_DC5((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))], (d_66_v0_) >= (d_66_v0_), (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))])
        d_109_v39_ = d_109_v39_
        hi1_ = (d_89_v20_)[default__.safeIndex(d_66_v0_, len(d_89_v20_))]
        for d_110_i3_ in range(d_66_v0_, hi1_):
            d_82_v14_ = (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ug"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_111_i4_ in range(default__.abs(922))]))).set(default__.safeIndex((0) - (d_66_v0_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ug"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_112_i4_ in range(default__.abs(922))])))), d_74_v7_)) + (d_82_v14_)
            d_113_v40_: int
            out0_: int
            out0_ = default__.m0(d_110_i3_, d_110_i3_, d_85_globalState_)
            d_113_v40_ = out0_
            d_114_v41_: _dafny.Map
            d_114_v41_ = _dafny.Map({d_66_v0_: d_113_v40_})
            d_115_v43_: _dafny.Seq
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: int
                for compr_11_ in (d_89_v20_).Elements:
                    d_116_v42_: int = compr_11_
                    if (d_116_v42_) in (d_89_v20_):
                        coll11_[default__.safeDivisionInt(d_116_v42_, 135)] = 597
                return _dafny.Map(coll11_)
            d_115_v43_ = _dafny.SeqWithoutIsStrInference([d_114_v41_, _dafny.Map({d_66_v0_: -223}), iife11_()
            , _dafny.Map({d_66_v0_: d_110_i3_}), d_114_v41_])
            d_117_v44_: _dafny.Map
            d_117_v44_ = _dafny.Map({d_68_v2_: d_115_v43_})
            d_117_v44_ = (d_117_v44_).set(d_68_v2_, d_115_v43_)
            index11_ = default__.safeIndex(424, (d_69_v3_).length(0))
            (d_69_v3_)[index11_] = d_66_v0_
            index12_ = default__.safeIndex(424, (d_69_v3_).length(0))
            (d_69_v3_)[index12_] = d_113_v40_
        (d_85_globalState_).f18 = len(d_82_v14_)
        d_118_v45_: _dafny.Array
        nw11_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
        d_118_v45_ = nw11_
        index13_ = default__.safeIndex(662, (d_118_v45_).length(0))
        (d_118_v45_)[index13_] = d_82_v14_
        d_119_v46_: _dafny.Map
        d_119_v46_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_82_v14_, d_82_v14_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncd")), d_82_v14_, d_82_v14_]): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbmfhj"))).set(default__.safeIndex(-133, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbmfhj")))), d_74_v7_)})
        d_120_v47_: _dafny.Seq
        d_120_v47_ = _dafny.SeqWithoutIsStrInference([d_82_v14_])
        d_121_v48_: _dafny.Seq
        d_121_v48_ = _dafny.SeqWithoutIsStrInference([d_82_v14_, (d_120_v47_)[default__.safeIndex(d_66_v0_, len(d_120_v47_))], d_82_v14_, d_82_v14_, d_82_v14_])
        index14_ = default__.safeIndex(662, (d_118_v45_).length(0))
        (d_118_v45_)[index14_] = ((d_119_v46_)[d_120_v47_] if (d_120_v47_) in (d_119_v46_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anrxvxycj")))
        rhs8_ = ((0) - (d_66_v0_) if default__.fm2(len(d_88_v19_), d_66_v0_, d_74_v7_, d_85_globalState_) else (d_66_v0_ if not(d_68_v2_) else (0) - (d_66_v0_)))
        rhs9_ = not(((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]) == (d_68_v2_))
        lhs9_ = d_85_globalState_
        lhs9_.f14 = rhs8_
        d_68_v2_ = rhs9_
        d_122_v49_: _dafny.Set
        d_122_v49_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyghqhrgx")))})
        d_123_v50_: D0
        d_123_v50_ = D0_DC1(d_66_v0_, d_66_v0_, not((d_73_v6_)[default__.safeIndex(len(_dafny.Map({760: d_122_v49_})), len(d_73_v6_))]))
        source0_ = d_123_v50_
        if source0_.is_DC1:
            d_124___mcc_h0_ = source0_.cf1
            d_125___mcc_h1_ = source0_.cf2
            d_126___mcc_h2_ = source0_.cf3
            d_127_cf3_ = d_126___mcc_h2_
            d_128_cf2_ = d_125___mcc_h1_
            d_129_cf1_ = d_124___mcc_h0_
            d_130_v51_: C0
            nw12_ = C0()
            nw12_.ctor__(not (d_127_cf3_) or (d_68_v2_))
            d_130_v51_ = nw12_
            d_131_v52_: int
            out1_: int
            out1_ = default__.m0(d_129_cf1_, default__.safeDivisionInt(d_66_v0_, d_66_v0_), d_85_globalState_)
            d_131_v52_ = out1_
            d_127_cf3_ = (d_130_v51_).f34
            d_127_cf3_ = (d_130_v51_).f34
        elif source0_.is_DC2:
            d_132___mcc_h3_ = source0_.cf4
            d_133___mcc_h4_ = source0_.cf5
            d_134___mcc_h5_ = source0_.cf6
            d_135_cf6_ = d_134___mcc_h5_
            d_136_cf5_ = d_133___mcc_h4_
            d_137_cf4_ = d_132___mcc_h3_
            d_68_v2_ = not ((d_137_cf4_ if default__.fm2(d_66_v0_, len(d_73_v6_), _dafny.CodePoint('t'), d_85_globalState_) else d_68_v2_)) or (not((d_81_v13_).issubset(d_81_v13_)))
            d_138_v53_: _dafny.Map
            d_138_v53_ = _dafny.Map({d_66_v0_: d_89_v20_})
            d_138_v53_ = d_138_v53_
            (d_85_globalState_).f9 = (d_136_cf5_) - (969)
            index15_ = default__.safeIndex(490, (d_69_v3_).length(0))
            (d_69_v3_)[index15_] = default__.safeDivisionInt(d_66_v0_, d_136_cf5_)
            index16_ = default__.safeIndex(490, (d_69_v3_).length(0))
            (d_69_v3_)[index16_] = (0) - (default__.safeDivisionInt(776, d_136_cf5_))
        elif True:
            d_139___mcc_h6_ = source0_.cf0
            d_140_cf0_ = d_139___mcc_h6_
            if (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]:
                d_66_v0_ = d_66_v0_
                d_141_v54_: _dafny.Array
                nw13_ = _dafny.Array(_dafny.Seq({}), 19)
                d_141_v54_ = nw13_
                index17_ = default__.safeIndex(738, (d_141_v54_).length(0))
                (d_141_v54_)[index17_] = d_89_v20_
                index18_ = default__.safeIndex(738, (d_141_v54_).length(0))
                (d_141_v54_)[index18_] = d_89_v20_
                def lambda3_(d_142_v0_):
                    def lambda4_(d_143_i5_):
                        return (d_143_i5_) - (d_142_v0_)

                    return lambda4_

                init2_ = lambda3_(d_66_v0_)
                nw14_ = _dafny.Array(None, 29)
                for i0_2_ in range(nw14_.length(0)):
                    nw14_[i0_2_] = init2_(i0_2_)
                d_69_v3_ = nw14_
                d_144_v55_: _dafny.Map
                d_144_v55_ = _dafny.Map({d_74_v7_: d_66_v0_})
                d_145_v56_: _dafny.Seq
                d_145_v56_ = _dafny.SeqWithoutIsStrInference([d_144_v55_])
                d_146_v57_: C4
                nw15_ = C4()
                nw15_.ctor__(d_69_v3_, d_66_v0_, d_145_v56_)
                d_146_v57_ = nw15_
                d_147_v58_: _dafny.Seq
                d_147_v58_ = _dafny.SeqWithoutIsStrInference([d_146_v57_, d_146_v57_])
                d_148_v59_: C4
                nw16_ = C4()
                nw16_.ctor__(d_140_cf0_, len(d_147_v58_), d_145_v56_)
                d_148_v59_ = nw16_
                d_149_v60_: _dafny.MultiSet
                d_149_v60_ = _dafny.MultiSet([-453, d_66_v0_])
                d_150_v61_: _dafny.Seq
                d_150_v61_ = _dafny.SeqWithoutIsStrInference([d_149_v60_, _dafny.MultiSet([d_66_v0_])])
                d_151_v62_: C1
                nw17_ = C1()
                nw17_.ctor__(d_150_v61_, d_66_v0_, ((d_67_v1_)[True] if (True) in (d_67_v1_) else (d_149_v60_).cardinality), d_145_v56_)
                d_151_v62_ = nw17_
                d_152_v63_: C1
                d_152_v63_ = d_151_v62_
                d_153_v64_: _dafny.Seq
                d_153_v64_ = _dafny.SeqWithoutIsStrInference([d_151_v62_])
                d_154_v65_: _dafny.Array
                nw18_ = _dafny.Array(None, 26)
                nw18_[int(0)] = d_151_v62_
                nw18_[int(1)] = d_151_v62_
                nw18_[int(2)] = d_151_v62_
                nw18_[int(3)] = d_151_v62_
                nw18_[int(4)] = d_151_v62_
                nw18_[int(5)] = d_151_v62_
                nw18_[int(6)] = d_151_v62_
                nw18_[int(7)] = d_151_v62_
                nw18_[int(8)] = d_151_v62_
                nw18_[int(9)] = (d_152_v63_)
                nw18_[int(10)] = d_151_v62_
                nw18_[int(11)] = d_151_v62_
                nw18_[int(12)] = d_151_v62_
                nw18_[int(13)] = d_151_v62_
                nw18_[int(14)] = (d_153_v64_)[default__.safeIndex(len((d_118_v45_)[default__.safeIndex(662, (d_118_v45_).length(0))]), len(d_153_v64_))]
                nw18_[int(15)] = d_151_v62_
                nw18_[int(16)] = d_151_v62_
                nw18_[int(17)] = d_151_v62_
                nw18_[int(18)] = (d_153_v64_)[default__.safeIndex(d_66_v0_, len(d_153_v64_))]
                nw18_[int(19)] = (d_153_v64_)[default__.safeIndex(d_151_v62_.f33, len(d_153_v64_))]
                nw18_[int(20)] = d_151_v62_
                nw18_[int(21)] = d_151_v62_
                nw18_[int(22)] = d_151_v62_
                nw18_[int(23)] = d_151_v62_
                nw18_[int(24)] = d_151_v62_
                nw18_[int(25)] = d_151_v62_
                d_154_v65_ = nw18_
                index19_ = default__.safeIndex(228, (d_154_v65_).length(0))
                (d_154_v65_)[index19_] = d_151_v62_
                index20_ = default__.safeIndex(228, (d_154_v65_).length(0))
                rhs10_ = d_74_v7_
                rhs11_ = (0) - (d_66_v0_)
                rhs12_ = d_151_v62_
                rhs13_ = ((((d_118_v45_)[default__.safeIndex(662, (d_118_v45_).length(0))]).set(default__.safeIndex(d_151_v62_.f33, len((d_118_v45_)[default__.safeIndex(662, (d_118_v45_).length(0))])), d_74_v7_)) + (d_82_v14_)) <= (d_82_v14_)
                lhs10_ = d_85_globalState_
                lhs11_ = d_154_v65_
                lhs12_ = default__.safeIndex(228, (d_154_v65_).length(0))
                d_74_v7_ = rhs10_
                lhs10_.f18 = rhs11_
                lhs11_[lhs12_] = rhs12_
                d_68_v2_ = rhs13_
            elif True:
                d_155_v67_: _dafny.Set
                d_155_v67_ = _dafny.Set({d_74_v7_, d_74_v7_})
                d_156_v68_: D4
                d_156_v68_ = D4_DC12(d_89_v20_)
                d_157_v69_: _dafny.Map
                d_157_v69_ = _dafny.Map({(d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]: d_156_v68_})
                d_158_v70_: _dafny.Map
                d_158_v70_ = _dafny.Map({d_74_v7_: len(d_157_v69_)})
                d_159_v71_: _dafny.Seq
                def iife12_():
                    coll12_ = _dafny.Map()
                    compr_12_: str
                    for compr_12_ in (d_155_v67_).Elements:
                        d_160_v66_: str = compr_12_
                        if (d_160_v66_) in (d_155_v67_):
                            coll12_[d_160_v66_] = d_66_v0_
                    return _dafny.Map(coll12_)
                d_159_v71_ = _dafny.SeqWithoutIsStrInference([iife12_()
                , d_158_v70_])
                d_161_v72_: C2
                nw19_ = C2()
                nw19_.ctor__((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))], len((d_118_v45_)[default__.safeIndex(662, (d_118_v45_).length(0))]), d_159_v71_)
                d_161_v72_ = nw19_
                d_162_v73_: _dafny.Seq
                d_162_v73_ = _dafny.SeqWithoutIsStrInference([d_161_v72_])
                d_163_v74_: _dafny.MultiSet
                d_163_v74_ = _dafny.MultiSet([(d_162_v73_)[default__.safeIndex(d_66_v0_, len(d_162_v73_))], d_161_v72_])
                d_164_v75_: _dafny.Map
                d_164_v75_ = _dafny.Map({True: d_161_v72_})
                d_165_v76_: T0
                nw20_ = C4()
                nw20_.ctor__(d_69_v3_, ((d_163_v74_)[((d_164_v75_)[(d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]] if ((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]) in (d_164_v75_) else d_161_v72_)] if (((d_164_v75_)[(d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]] if ((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]) in (d_164_v75_) else d_161_v72_)) in (d_163_v74_) else d_66_v0_), d_159_v71_)
                d_165_v76_ = nw20_
                d_165_v76_ = d_165_v76_
                d_67_v1_ = d_67_v1_
                d_166_v77_: C2
                nw21_ = C2()
                nw21_.ctor__((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))], (0) - (default__.safeDivisionInt((d_89_v20_)[default__.safeIndex((d_165_v76_).f24, len(d_89_v20_))], 687)), default__.fm42((d_165_v76_).f24, d_85_globalState_))
                d_166_v77_ = nw21_
                (d_85_globalState_).f17 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqdcd"))) + (_dafny.SeqWithoutIsStrInference([d_74_v7_ for d_167_i6_ in range(default__.abs(-23))]))
                d_168_v78_: C0
                nw22_ = C0()
                nw22_.ctor__((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))])
                d_168_v78_ = nw22_
                d_168_v78_ = d_168_v78_
            (d_85_globalState_).f2 = (default__.safeDivisionInt((0) - (default__.fm0(d_66_v0_, d_66_v0_, d_85_globalState_)), d_66_v0_)) in ((d_89_v20_) + (d_89_v20_))
            d_169_v79_: _dafny.Array
            nw23_ = _dafny.Array(_dafny.Set({}), 25)
            d_169_v79_ = nw23_
            d_170_v80_: _dafny.Seq
            d_170_v80_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_74_v7_: d_66_v0_})).set(d_74_v7_, d_66_v0_)])
            d_171_v81_: C2
            nw24_ = C2()
            nw24_.ctor__(False, d_66_v0_, (d_170_v80_))
            d_171_v81_ = nw24_
            d_172_v82_: _dafny.Map
            d_172_v82_ = _dafny.Map({d_66_v0_: 83})
            d_173_v83_: _dafny.Array
            nw25_ = _dafny.Array(None, 4)
            nw25_[int(0)] = d_172_v82_
            nw25_[int(1)] = d_172_v82_
            nw25_[int(2)] = d_172_v82_
            nw25_[int(3)] = d_172_v82_
            d_173_v83_ = nw25_
            d_174_v84_: _dafny.Map
            d_174_v84_ = _dafny.Map({d_171_v81_: d_173_v83_})
            d_175_v85_: C3
            nw26_ = C3()
            nw26_.ctor__(d_169_v79_, ((d_174_v84_)[d_171_v81_] if (d_171_v81_) in (d_174_v84_) else d_173_v83_), 869, _dafny.SeqWithoutIsStrInference([_dafny.Map({d_74_v7_: d_66_v0_}) for d_176_i7_ in range(default__.abs(510))]))
            d_175_v85_ = nw26_
            (d_85_globalState_).f18 = default__.safeDivisionInt(len(d_82_v14_), len(_dafny.SeqWithoutIsStrInference([d_175_v85_])))
            d_177_v87_: _dafny.Seq
            d_177_v87_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_74_v7_: d_66_v0_})])
            d_178_v88_: _dafny.Map
            d_178_v88_ = _dafny.Map({d_74_v7_: d_66_v0_})
            d_179_v89_: C4
            nw27_ = C4()
            nw27_.ctor__(d_140_cf0_, d_66_v0_, (d_177_v87_).set(default__.safeIndex(d_66_v0_, len(d_177_v87_)), d_178_v88_))
            d_179_v89_ = nw27_
            d_180_v90_: _dafny.Map
            d_180_v90_ = _dafny.Map({(_dafny.Map({not((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]): d_74_v7_})).set(d_171_v81_.f31, d_74_v7_): d_179_v89_})
            def iife13_():
                coll13_ = _dafny.Set()
                compr_13_: int
                for compr_13_ in _dafny.IntegerRange(62, 277):
                    d_181_v86_: int = compr_13_
                    if ((62) <= (d_181_v86_)) and ((d_181_v86_) < (277)):
                        coll13_ = coll13_.union(_dafny.Set([default__.safeModuloInt(d_181_v86_, d_66_v0_)]))
                return _dafny.Set(coll13_)
            d_79_v12_ = (d_79_v12_).set((len(iife13_()
            )) * (len(d_180_v90_)), d_78_v11_)
        hi2_ = (0) - (d_66_v0_)
        for d_182_i8_ in range(default__.fm0(d_66_v0_, d_66_v0_, d_85_globalState_), hi2_):
            index21_ = default__.safeIndex(614, (d_78_v11_).length(0))
            (d_78_v11_)[index21_] = (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]
            (d_85_globalState_).f18 = d_182_i8_
            d_183_v91_: _dafny.MultiSet
            d_183_v91_ = _dafny.MultiSet([d_182_i8_])
            index22_ = default__.safeIndex(614, (d_78_v11_).length(0))
            (d_78_v11_)[index22_] = (((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))] if (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))] else d_68_v2_)) or ((_dafny.MultiSet([d_182_i8_])) != (d_183_v91_))
            index23_ = default__.safeIndex(926, (d_69_v3_).length(0))
            (d_69_v3_)[index23_] = d_182_i8_
            d_184_v92_: _dafny.Array
            nw28_ = _dafny.Array(None, 27)
            nw28_[int(0)] = _dafny.CodePoint('y')
            nw28_[int(1)] = d_74_v7_
            nw28_[int(2)] = d_74_v7_
            nw28_[int(3)] = d_74_v7_
            nw28_[int(4)] = d_74_v7_
            nw28_[int(5)] = (d_74_v7_ if (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))] else d_74_v7_)
            nw28_[int(6)] = d_74_v7_
            nw28_[int(7)] = d_74_v7_
            nw28_[int(8)] = _dafny.CodePoint('u')
            nw28_[int(9)] = d_74_v7_
            nw28_[int(10)] = ((d_118_v45_)[default__.safeIndex(662, (d_118_v45_).length(0))])[default__.safeIndex(d_182_i8_, len((d_118_v45_)[default__.safeIndex(662, (d_118_v45_).length(0))]))]
            nw28_[int(11)] = _dafny.CodePoint('l')
            nw28_[int(12)] = d_74_v7_
            nw28_[int(13)] = d_74_v7_
            nw28_[int(14)] = d_74_v7_
            nw28_[int(15)] = d_74_v7_
            nw28_[int(16)] = d_74_v7_
            nw28_[int(17)] = d_74_v7_
            nw28_[int(18)] = d_74_v7_
            nw28_[int(19)] = default__.fm3(d_85_globalState_)
            nw28_[int(20)] = d_74_v7_
            nw28_[int(21)] = d_74_v7_
            nw28_[int(22)] = d_74_v7_
            nw28_[int(23)] = d_74_v7_
            nw28_[int(24)] = d_74_v7_
            nw28_[int(25)] = d_74_v7_
            nw28_[int(26)] = d_74_v7_
            d_184_v92_ = nw28_
            index24_ = default__.safeIndex(997, (d_184_v92_).length(0))
            (d_184_v92_)[index24_] = d_74_v7_
            d_185_v93_: _dafny.Map
            d_185_v93_ = _dafny.Map({d_68_v2_: d_74_v7_})
            d_186_v94_: _dafny.Seq
            d_186_v94_ = _dafny.SeqWithoutIsStrInference([d_78_v11_])
            d_187_v95_: _dafny.Map
            d_187_v95_ = _dafny.Map({d_78_v11_: d_78_v11_})
            d_188_v96_: C0
            nw29_ = C0()
            nw29_.ctor__(d_68_v2_)
            d_188_v96_ = nw29_
            d_189_v97_: _dafny.Seq
            d_189_v97_ = _dafny.SeqWithoutIsStrInference([d_188_v96_])
            d_190_v98_: _dafny.MultiSet
            d_190_v98_ = _dafny.MultiSet([(d_186_v94_)[default__.safeIndex(d_66_v0_, len(d_186_v94_))], d_78_v11_, ((d_187_v95_)[d_78_v11_] if (d_78_v11_) in (d_187_v95_) else (d_186_v94_)[default__.safeIndex(len(d_189_v97_), len(d_186_v94_))])])
            index25_ = default__.safeIndex(926, (d_69_v3_).length(0))
            index26_ = default__.safeIndex(997, (d_184_v92_).length(0))
            rhs14_ = ((d_73_v6_) + (d_73_v6_)).set(default__.safeIndex(default__.safeDivisionInt((0) - (len(d_185_v93_)), len(default__.fm23(d_85_globalState_))), len((d_73_v6_) + (d_73_v6_))), (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))])
            rhs15_ = default__.fm0(d_66_v0_, default__.safeDivisionInt(len(d_185_v93_), d_66_v0_), d_85_globalState_)
            rhs16_ = 525
            rhs17_ = d_74_v7_
            rhs18_ = ((d_190_v98_)[d_78_v11_] if (d_78_v11_) in (d_190_v98_) else d_182_i8_)
            lhs13_ = d_85_globalState_
            lhs14_ = d_69_v3_
            lhs15_ = default__.safeIndex(926, (d_69_v3_).length(0))
            lhs16_ = d_184_v92_
            lhs17_ = default__.safeIndex(997, (d_184_v92_).length(0))
            lhs18_ = d_85_globalState_
            lhs13_.f20 = rhs14_
            lhs14_[lhs15_] = rhs15_
            d_66_v0_ = rhs16_
            lhs16_[lhs17_] = rhs17_
            lhs18_.f9 = rhs18_
        d_191_v99_: _dafny.Array
        def lambda5_(d_192_v0_):
            def lambda6_(d_193_i9_):
                return _dafny.Set({d_192_v0_, d_192_v0_})

            return lambda6_

        init3_ = lambda5_(d_66_v0_)
        nw30_ = _dafny.Array(None, 6)
        for i0_3_ in range(nw30_.length(0)):
            nw30_[i0_3_] = init3_(i0_3_)
        d_191_v99_ = nw30_
        d_194_v100_: _dafny.Array
        def lambda7_(d_195_v6_, d_196_v0_):
            def lambda8_(d_197_i10_):
                return _dafny.Map({len(d_195_v6_): d_196_v0_})

            return lambda8_

        init4_ = lambda7_(d_73_v6_, d_66_v0_)
        nw31_ = _dafny.Array(None, 25)
        for i0_4_ in range(nw31_.length(0)):
            nw31_[i0_4_] = init4_(i0_4_)
        d_194_v100_ = nw31_
        d_198_v102_: _dafny.Map
        d_198_v102_ = _dafny.Map({d_74_v7_: d_66_v0_})
        d_199_v103_: _dafny.Seq
        d_199_v103_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_74_v7_: -1}), d_198_v102_])
        d_200_v104_: C3
        nw32_ = C3()
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: str
            for compr_14_ in (_dafny.Map({_dafny.CodePoint('q'): d_68_v2_})).keys.Elements:
                d_201_v101_: str = compr_14_
                if (d_201_v101_) in (_dafny.Map({_dafny.CodePoint('q'): d_68_v2_})):
                    coll14_[d_201_v101_] = len(d_73_v6_)
            return _dafny.Map(coll14_)
        nw32_.ctor__(d_191_v99_, d_194_v100_, default__.safeModuloInt(d_66_v0_, d_66_v0_), (_dafny.SeqWithoutIsStrInference([iife14_()
 for d_202_i11_ in range(default__.abs(775))])) + (d_199_v103_))
        d_200_v104_ = nw32_
        d_203_v105_: _dafny.Map
        d_203_v105_ = _dafny.Map({d_66_v0_: d_68_v2_})
        d_204_v106_: _dafny.Map
        d_204_v106_ = _dafny.Map({(((d_203_v105_)[len(d_73_v6_)] if (len(d_73_v6_)) in (d_203_v105_) else True) if False else (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]): default__.fm43(d_66_v0_, not(default__.fm2(d_66_v0_, d_66_v0_, d_74_v7_, d_85_globalState_)), d_68_v2_, d_85_globalState_)})
        d_204_v106_ = (d_204_v106_) | (_dafny.Map({(d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]: d_83_v15_}))
        d_205_v107_: D3
        d_205_v107_ = D3_DC9(d_73_v6_)
        source1_ = d_205_v107_
        if source1_.is_DC10:
            (d_85_globalState_).f2 = not(d_68_v2_)
            d_206_v108_: bool
            d_207_v109_: bool
            d_208_v110_: int
            out2_: bool
            out3_: bool
            out4_: int
            out2_, out3_, out4_ = (d_200_v104_).m5(d_68_v2_, d_85_globalState_)
            d_206_v108_ = out2_
            d_207_v109_ = out3_
            d_208_v110_ = out4_
            d_209_v111_: _dafny.Array
            nw33_ = _dafny.Array(_dafny.Map({}), 3)
            d_209_v111_ = nw33_
            d_209_v111_ = d_209_v111_
            (d_85_globalState_).f1 = default__.fm0(d_66_v0_, (d_66_v0_) * (d_208_v110_), d_85_globalState_)
        elif source1_.is_DC11:
            d_210___mcc_h7_ = source1_.cf18
            d_211___mcc_h8_ = source1_.cf19
            d_212___mcc_h9_ = source1_.cf20
            d_213_cf20_ = d_212___mcc_h9_
            d_214_cf19_ = d_211___mcc_h8_
            d_215_cf18_ = d_210___mcc_h7_
            d_216_v112_: _dafny.Set
            d_216_v112_ = _dafny.Set({d_73_v6_, d_73_v6_, d_73_v6_, d_73_v6_})
            d_217_v113_: _dafny.Map
            d_217_v113_ = _dafny.Map({d_214_cf19_: d_216_v112_})
            d_218_v114_: _dafny.Set
            d_218_v114_ = ((d_217_v113_)[d_68_v2_] if (d_68_v2_) in (d_217_v113_) else _dafny.Set({d_73_v6_, d_73_v6_}))
            source2_ = d_218_v114_
            d_219___mcc_h11_ = source2_
            d_220_cf65_ = d_219___mcc_h11_
            (d_85_globalState_).f17 = (d_118_v45_)[default__.safeIndex(662, (d_118_v45_).length(0))]
            (d_200_v104_).m6(d_85_globalState_)
            d_221_v115_: C4
            nw34_ = C4()
            nw34_.ctor__(d_69_v3_, d_66_v0_, d_199_v103_)
            d_221_v115_ = nw34_
            d_222_v116_: _dafny.Map
            d_222_v116_ = _dafny.Map({d_89_v20_: d_214_cf19_})
            d_223_v117_: _dafny.Seq
            d_223_v117_ = _dafny.SeqWithoutIsStrInference([d_67_v1_])
            d_224_v118_: int
            d_225_v119_: int
            d_226_v120_: _dafny.Seq
            d_227_v121_: bool
            out5_: int
            out6_: int
            out7_: _dafny.Seq
            out8_: bool
            out5_, out6_, out7_, out8_ = (d_221_v115_).m4((((d_222_v116_)[d_89_v20_] if (d_89_v20_) in (d_222_v116_) else not(True)) if not(not(d_214_cf19_)) else (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]), not(((d_223_v117_).set(default__.safeIndex(d_66_v0_, len(d_223_v117_)), d_67_v1_)) <= (d_223_v117_)), d_85_globalState_)
            d_224_v118_ = out5_
            d_225_v119_ = out6_
            d_226_v120_ = out7_
            d_227_v121_ = out8_
            (d_200_v104_).m6(d_85_globalState_)
            if not(d_214_cf19_):
                rhs19_ = (d_82_v14_) + (d_82_v14_)
                rhs20_ = d_66_v0_
                lhs19_ = d_85_globalState_
                lhs20_ = d_85_globalState_
                lhs19_.f17 = rhs19_
                lhs20_.f14 = rhs20_
                (d_85_globalState_).f14 = (d_66_v0_) * (((d_89_v20_)[default__.safeIndex(d_215_cf18_, len(d_89_v20_))]) + (len(_dafny.Set({d_78_v11_}))))
                d_228_v122_: int
                out9_: int
                out9_ = default__.m0(928, (d_215_cf18_) * (len(d_89_v20_)), d_85_globalState_)
                d_228_v122_ = out9_
                (d_85_globalState_).f21 = default__.fm23(d_85_globalState_)
                index27_ = default__.safeIndex(614, (d_78_v11_).length(0))
                (d_78_v11_)[index27_] = d_214_cf19_
            elif True:
                d_229_v123_: _dafny.Array
                nw35_ = _dafny.Array(D7.default()(), 12)
                d_229_v123_ = nw35_
                d_230_v124_: _dafny.Map
                d_230_v124_ = _dafny.Map({d_68_v2_: d_229_v123_})
                d_230_v124_ = (d_230_v124_).set(d_68_v2_, d_229_v123_)
                d_231_v125_: _dafny.Array
                nw36_ = _dafny.Array(None, 6)
                d_231_v125_ = nw36_
                d_231_v125_ = d_231_v125_
                (d_85_globalState_).f14 = (369) - (d_215_cf18_)
                (d_85_globalState_).f2 = d_213_cf20_
                (d_85_globalState_).f18 = d_66_v0_
            (d_85_globalState_).f9 = d_66_v0_
        elif True:
            d_232___mcc_h10_ = source1_.cf17
            d_233_cf17_ = d_232___mcc_h10_
            d_234_v126_: C4
            nw37_ = C4()
            nw37_.ctor__(d_69_v3_, len((d_118_v45_)[default__.safeIndex(662, (d_118_v45_).length(0))]), d_199_v103_)
            d_234_v126_ = nw37_
            d_235_v127_: D3
            d_235_v127_ = D3_DC10()
            d_236_v128_: _dafny.Seq
            d_236_v128_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([default__.fm0(d_66_v0_, d_66_v0_, d_85_globalState_), len((d_118_v45_)[default__.safeIndex(662, (d_118_v45_).length(0))])])).set(len(d_203_v105_), default__.abs(d_66_v0_))])
            d_237_v129_: C1
            nw38_ = C1()
            nw38_.ctor__((d_236_v128_).set(default__.safeIndex((d_200_v104_).fm8(d_85_globalState_), len(d_236_v128_)), _dafny.MultiSet([-838, d_66_v0_, d_66_v0_])), d_66_v0_, len(d_67_v1_), d_199_v103_)
            d_237_v129_ = nw38_
            d_238_v130_: _dafny.Map
            d_238_v130_ = _dafny.Map({(d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]: d_234_v126_})
            d_239_v131_: _dafny.MultiSet
            d_239_v131_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_237_v129_.f33, len(d_233_cf17_), (0) - (d_66_v0_)]), d_89_v20_, _dafny.SeqWithoutIsStrInference([368])])
            rhs21_ = (D3_DC11(d_237_v129_.f33, d_68_v2_, d_68_v2_)).cf20
            rhs22_ = ((d_238_v130_)[d_68_v2_] if (d_68_v2_) in (d_238_v130_) else d_234_v126_)
            rhs23_ = default__.fm33(not(((d_239_v131_).cardinality) != (d_66_v0_)), d_237_v129_.f33, d_85_globalState_)
            rhs24_ = d_66_v0_
            rhs25_ = d_237_v129_
            lhs21_ = d_85_globalState_
            lhs22_ = d_85_globalState_
            lhs21_.f2 = rhs21_
            d_234_v126_ = rhs22_
            d_235_v127_ = rhs23_
            lhs22_.f18 = rhs24_
            d_237_v129_ = rhs25_
            d_240_v132_: _dafny.Array
            nw39_ = _dafny.Array(_dafny.Map({}), 7)
            d_240_v132_ = nw39_
            d_240_v132_ = d_240_v132_
            if d_68_v2_:
                (d_85_globalState_).f9 = -974
                (d_85_globalState_).f18 = (d_66_v0_) * (-481)
                index28_ = default__.safeIndex(614, (d_78_v11_).length(0))
                (d_78_v11_)[index28_] = (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))]
                d_241_v133_: C0
                nw40_ = C0()
                nw40_.ctor__(not(True))
                d_241_v133_ = nw40_
                d_242_v134_: _dafny.Array
                nw41_ = _dafny.Array(None, 1)
                nw41_[int(0)] = d_241_v133_
                d_242_v134_ = nw41_
                index29_ = default__.safeIndex(79, (d_242_v134_).length(0))
                (d_242_v134_)[index29_] = d_241_v133_
                index30_ = default__.safeIndex(79, (d_242_v134_).length(0))
                (d_242_v134_)[index30_] = d_241_v133_
                d_243_v135_: T0
                nw42_ = C4()
                nw42_.ctor__(d_69_v3_, default__.fm0(d_237_v129_.f33, d_237_v129_.f33, d_85_globalState_), (d_199_v103_) + (_dafny.SeqWithoutIsStrInference([d_198_v102_])))
                d_243_v135_ = nw42_
                rhs26_ = d_243_v135_
                rhs27_ = (d_118_v45_)[default__.safeIndex(662, (d_118_v45_).length(0))]
                rhs28_ = (0) - ((495) - (((d_237_v129_).fm11(d_68_v2_, (d_241_v133_).f34, (d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))], d_66_v0_, d_85_globalState_)) + ((d_243_v135_).f24)))
                lhs23_ = d_85_globalState_
                lhs24_ = d_85_globalState_
                d_243_v135_ = rhs26_
                lhs23_.f17 = rhs27_
                lhs24_.f9 = rhs28_
            elif True:
                d_244_v136_: bool
                d_245_v137_: bool
                d_246_v138_: int
                out10_: bool
                out11_: bool
                out12_: int
                out10_, out11_, out12_ = (d_200_v104_).m5((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))], d_85_globalState_)
                d_244_v136_ = out10_
                d_245_v137_ = out11_
                d_246_v138_ = out12_
                (d_85_globalState_).f18 = d_66_v0_
                (d_85_globalState_).f18 = d_237_v129_.f33
                (d_85_globalState_).f14 = (0) - (d_237_v129_.f33)
                (d_85_globalState_).f18 = default__.fm0(d_66_v0_, (0) - ((d_237_v129_.f33) + (d_237_v129_.f33)), d_85_globalState_)
            d_247_v139_: _dafny.MultiSet
            d_247_v139_ = _dafny.MultiSet([d_237_v129_.f33, len(d_82_v14_)])
            index31_ = default__.safeIndex(614, (d_78_v11_).length(0))
            (d_78_v11_)[index31_] = not ((d_247_v139_).ispropersubset(d_247_v139_)) or ((d_78_v11_)[default__.safeIndex(614, (d_78_v11_).length(0))])
        (d_85_globalState_).f18 = default__.safeModuloInt((0) - (((_dafny.MultiSet([d_66_v0_, len(d_89_v20_), d_66_v0_])).cardinality) + ((0) - (d_66_v0_))), d_66_v0_)
        (d_85_globalState_).f18 = d_66_v0_
        d_203_v105_ = (d_203_v105_).set(d_66_v0_, not((d_122_v49_).issubset(d_122_v49_)))
        _dafny.print(_dafny.string_of(d_66_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_67_v1_) == (_dafny.Map({False: 590}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_68_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v3_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_71_v4_).cf0)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_72_v5_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_73_v6_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_74_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_75_v8_).cf7) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c'), _dafny.CodePoint('h')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v9_) == (_dafny.MultiSet([_dafny.CodePoint('c')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_77_v10_) == (_dafny.Set({_dafny.MultiSet([_dafny.CodePoint('c'), _dafny.CodePoint('h')]), _dafny.MultiSet([_dafny.CodePoint('c')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v11_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_79_v12_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_81_v13_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_82_v14_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_83_v15_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_83_v15_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_83_v15_).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_83_v15_).cf11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_83_v15_).cf12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_84_v16_) == (_dafny.Map({590: _dafny.Map({382: 590})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_85_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_85_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_85_globalState_).f3) == (_dafny.Map({False: 590, True: 590}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_85_globalState_).f7)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_85_globalState_).f8) == (_dafny.SeqWithoutIsStrInference([False, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_85_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_85_globalState_).f11) == (_dafny.Set({_dafny.MultiSet([_dafny.CodePoint('c'), _dafny.CodePoint('h')]), _dafny.MultiSet([_dafny.CodePoint('c')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_85_globalState_).f13)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_85_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_85_globalState_).f16) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_85_globalState_.f17).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_85_globalState_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_.f19).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_.f20) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_.f21) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_.f22) == (_dafny.Map({382: 590}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v18_) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_88_v19_) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v20_) == (_dafny.SeqWithoutIsStrInference([590, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_109_v39_).cf13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_109_v39_).cf14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_109_v39_).cf15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_118_v45_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v46_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncd")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh"))]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubmfhj"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_120_v47_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_121_v48_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbntyekh"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_122_v49_) == (_dafny.Set({9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v50_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v50_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_123_v50_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_191_v99_)[0]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_191_v99_)[1]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_191_v99_)[2]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_191_v99_)[3]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_191_v99_)[4]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_191_v99_)[5]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[0]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[1]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[2]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[3]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[4]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[5]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[6]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[7]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[8]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[9]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[10]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[11]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[12]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[13]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[14]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[15]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[16]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[17]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[18]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[19]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[20]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[21]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[22]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[23]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_194_v100_)[24]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_v102_) == (_dafny.Map({_dafny.CodePoint('u'): 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v103_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('u'): -1}), _dafny.Map({_dafny.CodePoint('u'): 525})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_200_v104_).f29)[0]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_200_v104_).f29)[1]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_200_v104_).f29)[2]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_200_v104_).f29)[3]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_200_v104_).f29)[4]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_200_v104_).f29)[5]) == (_dafny.Set({525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[0]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[1]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[2]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[3]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[4]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[5]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[6]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[7]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[8]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[9]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[10]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[11]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[12]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[13]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[14]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[15]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[16]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[17]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[18]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[19]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[20]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[21]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[22]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[23]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_200_v104_.f30)[24]) == (_dafny.Map({1: 525}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_v105_) == (_dafny.Map({525: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v106_) == (_dafny.Map({True: D1_DC4(True, True, _dafny.CodePoint('k'), 590, 590)}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_205_v107_).cf17) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0), False)
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
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
        return lambda: D1_DC4(False, False, _dafny.CodePoint('D'), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({self.cf7.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10()
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

class D3_DC10(D3, NamedTuple('DC10', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(int(0), _dafny.Map({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)

class D4_DC13(D4, NamedTuple('DC13', [('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {self.cf24.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC15(D5, NamedTuple('DC15', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(False)
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
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)

class D6_DC17(D6, NamedTuple('DC17', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC21(D1.default()(), D6.default()(), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)

class D7_DC21(D7, NamedTuple('DC21', [('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC22(D7, NamedTuple('DC22', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC24(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)

class D8_DC24(D8, NamedTuple('DC24', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC26(D8, NamedTuple('DC26', [('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26({_dafny.string_of(self.cf41)}, {self.cf42.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)

class D9_DC27(D9, NamedTuple('DC27', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC29(D10, NamedTuple('DC29', [('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC31(False, int(0), _dafny.Array(None, 0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC31(D11, NamedTuple('DC31', [('cf48', Any), ('cf49', Any), ('cf50', Any), ('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC33(_dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)

class D12_DC33(D12, NamedTuple('DC33', [('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC36(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)

class D13_DC36(D13, NamedTuple('DC36', [('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC35(D13, NamedTuple('DC35', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC37(D13, NamedTuple('DC37', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D14_DC38)

class D14_DC38(D14, NamedTuple('DC38', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC38({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC38) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D15_DC39)

class D15_DC39(D15, NamedTuple('DC39', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC39({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC39) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC41(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D16_DC40)

class D16_DC41(D16, NamedTuple('DC41', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC40(D16, NamedTuple('DC40', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC40({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC40) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D17_DC42)

class D17_DC42(D17, NamedTuple('DC42', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC42({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC42) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D18_DC43)

class D18_DC43(D18, NamedTuple('DC43', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC45(int(0), False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D19_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D19_DC46)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D19_DC47)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D19_DC44)

class D19_DC45(D19, NamedTuple('DC45', [('cf68', Any), ('cf69', Any), ('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC45({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC45) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC46(D19, NamedTuple('DC46', [('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC46({_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC46) and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC47(D19, NamedTuple('DC47', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC47({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC47) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC44(D19, NamedTuple('DC44', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC44({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC44) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC49(_dafny.MultiSet({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D20_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D20_DC48)

class D20_DC49(D20, NamedTuple('DC49', [('cf76', Any), ('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC49({_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC49) and self.cf76 == __o.cf76 and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC48(D20, NamedTuple('DC48', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC48({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC48) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: int = int(0)
        self.f2: bool = False
        self.f9: int = int(0)
        self.f14: int = int(0)
        self.f17: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f18: int = int(0)
        self.f19: _dafny.MultiSet = _dafny.MultiSet({})
        self.f20: _dafny.Seq = _dafny.Seq({})
        self.f21: _dafny.Seq = _dafny.Seq({})
        self.f22: _dafny.Map = _dafny.Map({})
        self._f0: bool = False
        self._f3: _dafny.Map = _dafny.Map({})
        self._f4: bool = False
        self._f5: int = int(0)
        self._f6: str = _dafny.CodePoint('D')
        self._f7: _dafny.Map = _dafny.Map({})
        self._f8: _dafny.Seq = _dafny.Seq({})
        self._f10: int = int(0)
        self._f11: _dafny.Set = _dafny.Set({})
        self._f12: int = int(0)
        self._f13: _dafny.Map = _dafny.Map({})
        self._f15: int = int(0)
        self._f16: _dafny.Set = _dafny.Set({})
        self._f23: int = int(0)
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
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self).f17 = f17
        (self).f18 = f18
        (self).f19 = f19
        (self).f20 = f20
        (self).f21 = f21
        (self).f22 = f22
        (self)._f23 = f23

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
    def f23(self):
        return self._f23

class C0:
    def  __init__(self):
        self._f34: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f34):
        (self)._f34 = f34

    def fm12(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference([((0) - (-780)) + (-448), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))])

    @property
    def f34(self):
        return self._f34

class C1(T0):
    def  __init__(self):
        self._f24: int = int(0)
        self._f25: _dafny.Seq = _dafny.Seq({})
        self.f33: int = int(0)
        self._f32: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f32, f33, f24, f25):
        (self)._f32 = f32
        (self).f33 = f33
        (self)._f24 = f24
        (self)._f25 = f25

    def fm10(self, globalState):
        return False

    def fm11(self, p0, p1, p2, p3, globalState):
        return default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgrpl"))), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l'), _dafny.CodePoint('b')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('y'), _dafny.CodePoint('p')]))))

    def m1(self, globalState):
        r0: bool = False
        if ((self.f33) - (self.f33)) <= (self.f33):
            d_248_v0_: bool
            d_248_v0_ = True
            (globalState).f2 = d_248_v0_
            if d_248_v0_:
                (globalState).f1 = (self).f24
                d_249_v1_: str
                d_249_v1_ = _dafny.CodePoint('i')
                d_249_v1_ = d_249_v1_
                d_250_v2_: _dafny.Array
                def lambda9_(d_251_v0_):
                    def lambda10_(d_252_i0_):
                        return d_251_v0_

                    return lambda10_

                init5_ = lambda9_(d_248_v0_)
                nw43_ = _dafny.Array(None, 9)
                for i0_5_ in range(nw43_.length(0)):
                    nw43_[i0_5_] = init5_(i0_5_)
                d_250_v2_ = nw43_
                d_253_v3_: _dafny.Seq
                d_253_v3_ = _dafny.SeqWithoutIsStrInference([d_248_v0_])
                d_254_v4_: _dafny.Set
                d_254_v4_ = _dafny.Set({d_248_v0_, (d_253_v3_)[default__.safeIndex(730, len(d_253_v3_))], d_248_v0_})
                d_250_v2_ = (d_250_v2_ if (_dafny.Set({d_248_v0_})).issubset(d_254_v4_) else d_250_v2_)
                d_255_v5_: _dafny.Map
                d_255_v5_ = _dafny.Map({self.f33: d_248_v0_})
                d_256_v6_: _dafny.Map
                d_256_v6_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_253_v3_]): (self).fm10(globalState)})
                d_257_v7_: _dafny.Map
                d_257_v7_ = _dafny.Map({d_248_v0_: d_248_v0_})
                d_258_v8_: D3
                d_258_v8_ = D3_DC9(d_253_v3_)
                d_259_v9_: _dafny.Map
                d_259_v9_ = _dafny.Map({len(d_257_v7_): (d_258_v8_).cf17})
                (globalState).f2 = ((d_255_v5_)[self.f33] if (self.f33) in (d_255_v5_) else ((d_256_v6_)[_dafny.SeqWithoutIsStrInference([((d_259_v9_)[(0) - ((self).f24)] if ((0) - ((self).f24)) in (d_259_v9_) else d_253_v3_)])] if (_dafny.SeqWithoutIsStrInference([((d_259_v9_)[(0) - ((self).f24)] if ((0) - ((self).f24)) in (d_259_v9_) else d_253_v3_)])) in (d_256_v6_) else d_248_v0_))
                d_260_v10_: _dafny.Seq
                d_260_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clyji"))
                d_261_v11_: _dafny.Set
                d_261_v11_ = _dafny.Set({d_260_v10_})
                d_262_v12_: _dafny.Seq
                d_262_v12_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((self).f24, len(d_261_v11_)), (self).f24, (_dafny.MultiSet([d_248_v0_])).cardinality])
                rhs29_ = d_250_v2_
                rhs30_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))) + (_dafny.SeqWithoutIsStrInference([d_249_v1_ for d_263_i1_ in range(default__.abs(888))])))
                rhs31_ = (self).f24
                rhs32_ = d_262_v12_
                lhs25_ = globalState
                lhs26_ = globalState
                d_250_v2_ = rhs29_
                lhs25_.f9 = rhs30_
                lhs26_.f9 = rhs31_
                d_262_v12_ = rhs32_
            elif True:
                d_264_v13_: _dafny.MultiSet
                d_264_v13_ = _dafny.MultiSet([self.f33])
                d_265_v14_: _dafny.MultiSet
                d_265_v14_ = _dafny.MultiSet([d_248_v0_, d_248_v0_, d_248_v0_])
                d_266_v15_: C0
                nw44_ = C0()
                nw44_.ctor__((((d_264_v13_).set(self.f33, default__.abs((self).f24))).cardinality) != (((d_265_v14_)[d_248_v0_] if (d_248_v0_) in (d_265_v14_) else len(_dafny.Set({True, not(d_248_v0_), d_248_v0_})))))
                d_266_v15_ = nw44_
                d_267_v16_: _dafny.Array
                def lambda11_(d_268_i2_):
                    return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_269_i3_ in range(default__.abs(648))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_270_i4_ in range(default__.abs(216))]))

                init6_ = lambda11_
                nw45_ = _dafny.Array(None, 17)
                for i0_6_ in range(nw45_.length(0)):
                    nw45_[i0_6_] = init6_(i0_6_)
                d_267_v16_ = nw45_
                d_271_v17_: _dafny.Seq
                d_271_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtih"))
                d_272_v18_: str
                d_272_v18_ = _dafny.CodePoint('q')
                index32_ = default__.safeIndex(433, (d_267_v16_).length(0))
                (d_267_v16_)[index32_] = (d_271_v17_).set(default__.safeIndex((self).fm11((d_266_v15_).f34, (d_266_v15_).f34, True, self.f33, globalState), len(d_271_v17_)), d_272_v18_)
                index33_ = default__.safeIndex(433, (d_267_v16_).length(0))
                (d_267_v16_)[index33_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcjpd"))) + (_dafny.SeqWithoutIsStrInference([d_272_v18_ for d_273_i5_ in range(default__.abs(558))]))
                d_266_v15_ = d_266_v15_
                d_274_v19_: _dafny.Map
                d_274_v19_ = _dafny.Map({(d_266_v15_).f34: d_272_v18_})
                d_275_v20_: _dafny.Map
                d_275_v20_ = _dafny.Map({not((d_266_v15_).f34): d_248_v0_})
                d_274_v19_ = (d_274_v19_).set(((d_275_v20_)[(d_266_v15_).f34] if ((d_266_v15_).f34) in (d_275_v20_) else d_248_v0_), d_272_v18_)
                d_276_v21_: D0
                d_276_v21_ = D0_DC1((self).f24, 949, ((d_266_v15_).f34 if default__.fm2(self.f33, (0) - ((self).f24), d_272_v18_, globalState) else (d_266_v15_).f34))
                d_277_v22_: _dafny.Map
                d_277_v22_ = _dafny.Map({not(d_248_v0_): self.f33})
                d_278_v23_: _dafny.Seq
                d_278_v23_ = _dafny.SeqWithoutIsStrInference([(d_266_v15_).f34, (d_266_v15_).f34, d_248_v0_])
                d_279_v24_: _dafny.Seq
                d_279_v24_ = _dafny.SeqWithoutIsStrInference([(d_277_v22_ if d_248_v0_ else d_277_v22_), (d_277_v22_).set(True, len(d_278_v23_)), d_277_v22_, d_277_v22_, d_277_v22_])
                rhs33_ = (d_266_v15_).f34
                rhs34_ = d_276_v21_
                rhs35_ = (self).f24
                rhs36_ = _dafny.SeqWithoutIsStrInference([d_277_v22_ for d_280_i6_ in range(default__.abs(655))])
                lhs27_ = globalState
                lhs28_ = globalState
                lhs27_.f2 = rhs33_
                d_276_v21_ = rhs34_
                lhs28_.f1 = rhs35_
                d_279_v24_ = rhs36_
            d_281_v25_: _dafny.Array
            nw46_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_281_v25_ = nw46_
            d_282_v26_: _dafny.Map
            d_282_v26_ = _dafny.Map({d_248_v0_: d_248_v0_})
            d_283_v27_: _dafny.Seq
            d_283_v27_ = _dafny.SeqWithoutIsStrInference([d_282_v26_])
            d_284_v28_: _dafny.Array
            nw47_ = _dafny.Array(None, 25)
            nw47_[int(0)] = default__.fm13(globalState)
            nw47_[int(1)] = _dafny.Map({d_248_v0_: d_248_v0_})
            nw47_[int(2)] = _dafny.Map({d_248_v0_: d_248_v0_})
            nw47_[int(3)] = d_282_v26_
            nw47_[int(4)] = d_282_v26_
            nw47_[int(5)] = _dafny.Map({d_248_v0_: d_248_v0_})
            nw47_[int(6)] = d_282_v26_
            nw47_[int(7)] = d_282_v26_
            nw47_[int(8)] = d_282_v26_
            nw47_[int(9)] = _dafny.Map({d_248_v0_: not(d_248_v0_)})
            nw47_[int(10)] = d_282_v26_
            nw47_[int(11)] = _dafny.Map({True: d_248_v0_})
            nw47_[int(12)] = d_282_v26_
            nw47_[int(13)] = d_282_v26_
            nw47_[int(14)] = d_282_v26_
            nw47_[int(15)] = (d_283_v27_)[default__.safeIndex((self).f24, len(d_283_v27_))]
            nw47_[int(16)] = d_282_v26_
            nw47_[int(17)] = d_282_v26_
            nw47_[int(18)] = d_282_v26_
            nw47_[int(19)] = d_282_v26_
            nw47_[int(20)] = d_282_v26_
            nw47_[int(21)] = (d_282_v26_).set(d_248_v0_, d_248_v0_)
            nw47_[int(22)] = d_282_v26_
            nw47_[int(23)] = _dafny.Map({d_248_v0_: not(d_248_v0_)})
            nw47_[int(24)] = d_282_v26_
            d_284_v28_ = nw47_
            index34_ = default__.safeIndex(428, (d_281_v25_).length(0))
            (d_281_v25_)[index34_] = d_284_v28_
            index35_ = default__.safeIndex(428, (d_281_v25_).length(0))
            (d_281_v25_)[index35_] = d_284_v28_
            d_285_v29_: _dafny.Set
            d_285_v29_ = _dafny.Set({342})
            d_285_v29_ = d_285_v29_
            d_286_v30_: _dafny.Array
            nw48_ = _dafny.Array(False, 14)
            d_286_v30_ = nw48_
            d_287_v31_: _dafny.Seq
            d_287_v31_ = _dafny.SeqWithoutIsStrInference([d_286_v30_, d_286_v30_, d_286_v30_])
            d_288_v32_: int
            d_289_v33_: bool
            d_290_v34_: int
            d_291_v35_: str
            out13_: int
            out14_: bool
            out15_: int
            out16_: str
            out13_, out14_, out15_, out16_ = (self).m8((d_287_v31_)[default__.safeIndex(self.f33, len(d_287_v31_))], globalState)
            d_288_v32_ = out13_
            d_289_v33_ = out14_
            d_290_v34_ = out15_
            d_291_v35_ = out16_
        elif True:
            d_292_v36_: bool
            d_292_v36_ = True
            if d_292_v36_:
                (globalState).f2 = d_292_v36_
                d_293_v37_: _dafny.Map
                d_293_v37_ = _dafny.Map({d_292_v36_: self.f33})
                d_294_v38_: D0
                d_294_v38_ = D0_DC1(self.f33, len(d_293_v37_), d_292_v36_)
                d_295_v39_: _dafny.Map
                d_295_v39_ = _dafny.Map({(d_294_v38_).cf3: self.f33})
                d_296_v40_: _dafny.Map
                d_296_v40_ = _dafny.Map({_dafny.CodePoint('i'): d_292_v36_})
                d_297_v41_: str
                d_297_v41_ = _dafny.CodePoint('c')
                d_298_v42_: _dafny.MultiSet
                d_298_v42_ = _dafny.MultiSet([d_292_v36_])
                d_299_v43_: D1
                d_299_v43_ = D1_DC4((self.f33) <= (len(d_295_v39_)), ((d_296_v40_)[_dafny.CodePoint('d')] if (_dafny.CodePoint('d')) in (d_296_v40_) else d_292_v36_), d_297_v41_, ((d_298_v42_)[True] if (True) in (d_298_v42_) else self.f33), (self).fm11(d_292_v36_, not(d_292_v36_), d_292_v36_, self.f33, globalState))
                d_299_v43_ = d_299_v43_
                d_300_v44_: C0
                nw49_ = C0()
                nw49_.ctor__(d_292_v36_)
                d_300_v44_ = nw49_
                d_298_v42_ = (d_298_v42_).intersection(default__.fm14(default__.fm15(224, d_292_v36_, (self).f24, (d_300_v44_).f34, globalState), self.f33, globalState))
                d_301_v45_: _dafny.Seq
                d_301_v45_ = _dafny.SeqWithoutIsStrInference([self.f33, self.f33])
                d_302_v46_: D4
                d_302_v46_ = D4_DC12(d_301_v45_)
                d_303_v47_: _dafny.Map
                d_303_v47_ = _dafny.Map({(((d_298_v42_).set((d_300_v44_).f34, default__.abs(self.f33))).cardinality) < (self.f33): (d_302_v46_).cf21})
                d_303_v47_ = (d_303_v47_).set((d_300_v44_).f34, (_dafny.SeqWithoutIsStrInference([self.f33, (self).f24])).set(default__.safeIndex(len(_dafny.Map({(0) - ((self).f24): self.f33})), len(_dafny.SeqWithoutIsStrInference([self.f33, (self).f24]))), (self).f24))
            elif True:
                d_304_v48_: _dafny.Map
                d_304_v48_ = _dafny.Map({(self).f24: ((0) - ((self).f24) if d_292_v36_ else (self).f24)})
                d_304_v48_ = (d_304_v48_).set((self).f24, (self).f24)
                d_305_v49_: _dafny.Array
                nw50_ = _dafny.Array(D0.default()(), 24)
                d_305_v49_ = nw50_
                d_305_v49_ = d_305_v49_
                d_306_v50_: D0
                d_306_v50_ = D0_DC1((0) - ((self).f24), 300, d_292_v36_)
                (globalState).f2 = not(((self).f24) == ((self.f33) + ((d_306_v50_).cf2)))
                d_307_v51_: _dafny.Seq
                d_307_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
                (globalState).f17 = d_307_v51_
                d_308_v52_: _dafny.Map
                d_308_v52_ = _dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_309_i7_ in range(default__.abs(329))])})
                d_308_v52_ = (d_308_v52_).set((self).f24, d_307_v51_)
            d_310_v53_: _dafny.Array
            def lambda12_(d_311_i8_):
                return (d_311_i8_) * (self.f33)

            init7_ = lambda12_
            nw51_ = _dafny.Array(None, 20)
            for i0_7_ in range(nw51_.length(0)):
                nw51_[i0_7_] = init7_(i0_7_)
            d_310_v53_ = nw51_
            index36_ = default__.safeIndex(353, (d_310_v53_).length(0))
            (d_310_v53_)[index36_] = (self).f24
            index37_ = default__.safeIndex(353, (d_310_v53_).length(0))
            (d_310_v53_)[index37_] = ((self).f24) * ((self).f24)
            (globalState).f14 = self.f33
            d_312_v54_: _dafny.Array
            def lambda13_(d_313_i9_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))

            init8_ = lambda13_
            nw52_ = _dafny.Array(None, 6)
            for i0_8_ in range(nw52_.length(0)):
                nw52_[i0_8_] = init8_(i0_8_)
            d_312_v54_ = nw52_
            d_314_v55_: _dafny.Seq
            d_314_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwklb"))
            index38_ = default__.safeIndex(2, (d_312_v54_).length(0))
            (d_312_v54_)[index38_] = d_314_v55_
            index39_ = default__.safeIndex(2, (d_312_v54_).length(0))
            (d_312_v54_)[index39_] = d_314_v55_
            d_315_v56_: _dafny.Map
            d_315_v56_ = _dafny.Map({d_292_v36_: d_292_v36_})
            d_316_v57_: _dafny.Map
            d_316_v57_ = _dafny.Map({d_292_v36_: len(d_315_v56_)})
            rhs37_ = (D4_DC13(self.f33, d_316_v57_, d_314_v55_)).cf22
            rhs38_ = (self).f24
            rhs39_ = d_292_v36_
            rhs40_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_317_i10_ in range(default__.abs(133))])
            lhs29_ = globalState
            lhs30_ = globalState
            lhs31_ = globalState
            lhs29_.f18 = rhs37_
            lhs30_.f9 = rhs38_
            d_292_v36_ = rhs39_
            lhs31_.f17 = rhs40_
        (globalState).f9 = (self).f24
        d_318_v58_: bool
        d_318_v58_ = True
        r0 = d_318_v58_
        d_318_v58_ = (self).fm10(globalState)
        d_319_v59_: _dafny.Seq
        d_319_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvcjan"))
        d_320_v60_: str
        d_320_v60_ = _dafny.CodePoint('t')
        d_321_v61_: _dafny.Seq
        d_321_v61_ = _dafny.SeqWithoutIsStrInference([d_319_v59_, d_319_v59_, default__.fm4(d_318_v58_, self.f33, d_320_v60_, not(d_318_v58_), globalState)])
        d_322_i11_: int
        d_322_i11_ = 0
        with _dafny.label("0"):
            while (d_321_v61_) != (_dafny.SeqWithoutIsStrInference([d_319_v59_ for d_363_i12_ in range(default__.abs(-168))])):
                with _dafny.c_label("0"):
                    if (d_322_i11_) >= (100):
                        raise _dafny.Break("0")
                    d_322_i11_ = (d_322_i11_) + (1)
                    d_323_v62_: _dafny.Set
                    d_323_v62_ = _dafny.Set({self.f33})
                    d_324_v63_: _dafny.Map
                    d_324_v63_ = _dafny.Map({d_318_v58_: len(d_323_v62_)})
                    d_325_v64_: _dafny.Map
                    d_325_v64_ = _dafny.Map({self.f33: d_320_v60_})
                    d_326_v65_: _dafny.Map
                    d_326_v65_ = _dafny.Map({len(d_325_v64_): not(d_318_v58_)})
                    d_327_v66_: _dafny.Array
                    nw53_ = _dafny.Array(None, 19)
                    nw53_[int(0)] = self.f33
                    nw53_[int(1)] = self.f33
                    nw53_[int(2)] = len(d_324_v63_)
                    nw53_[int(3)] = (self).f24
                    nw53_[int(4)] = 676
                    nw53_[int(5)] = len(d_319_v59_)
                    nw53_[int(6)] = len(default__.fm16(globalState))
                    nw53_[int(7)] = (0) - ((self).f24)
                    nw53_[int(8)] = (self).f24
                    nw53_[int(9)] = len(d_326_v65_)
                    nw53_[int(10)] = (self).f24
                    nw53_[int(11)] = ((self).f24) - (len(_dafny.SeqWithoutIsStrInference([d_319_v59_, d_319_v59_])))
                    nw53_[int(12)] = default__.safeModuloInt(self.f33, -6)
                    nw53_[int(13)] = (self).f24
                    nw53_[int(14)] = (self).f24
                    nw53_[int(15)] = len(d_326_v65_)
                    nw53_[int(16)] = self.f33
                    nw53_[int(17)] = default__.fm0(-523, self.f33, globalState)
                    nw53_[int(18)] = (self).f24
                    d_327_v66_ = nw53_
                    index40_ = default__.safeIndex(92, (d_327_v66_).length(0))
                    (d_327_v66_)[index40_] = self.f33
                    index41_ = default__.safeIndex(92, (d_327_v66_).length(0))
                    rhs41_ = (self).f24
                    rhs42_ = (d_327_v66_ if d_318_v58_ else d_327_v66_)
                    lhs32_ = d_327_v66_
                    lhs33_ = default__.safeIndex(92, (d_327_v66_).length(0))
                    lhs32_[lhs33_] = rhs41_
                    d_327_v66_ = rhs42_
                    d_328_v67_: D0
                    d_328_v67_ = D0_DC2(d_318_v58_, -118, d_320_v60_)
                    d_329_v68_: _dafny.Map
                    d_329_v68_ = _dafny.Map({d_328_v67_: d_318_v58_})
                    d_330_v69_: _dafny.Map
                    d_330_v69_ = _dafny.Map({d_318_v58_: (d_318_v58_ if d_318_v58_ else ((d_329_v68_)[d_328_v67_] if (d_328_v67_) in (d_329_v68_) else d_318_v58_))})
                    d_330_v69_ = (d_330_v69_).set(((d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))]) != ((self).f24), d_318_v58_)
                    d_331_v70_: D2
                    d_331_v70_ = D2_DC8()
                    source3_ = d_331_v70_
                    if source3_.is_DC8:
                        d_332_v71_: _dafny.Seq
                        d_332_v71_ = _dafny.SeqWithoutIsStrInference([len(d_319_v59_)])
                        d_333_v72_: _dafny.Seq
                        d_333_v72_ = _dafny.SeqWithoutIsStrInference([d_318_v58_, d_318_v58_])
                        d_334_v73_: _dafny.Array
                        nw54_ = _dafny.Array(None, 27)
                        nw54_[int(0)] = default__.fm17(790, (self).f24, 580, False, globalState)
                        nw54_[int(1)] = d_324_v63_
                        nw54_[int(2)] = d_324_v63_
                        nw54_[int(3)] = (d_324_v63_).set(d_318_v58_, (d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))])
                        nw54_[int(4)] = _dafny.Map({default__.fm2((d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))], (d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))], d_320_v60_, globalState): (self).f24})
                        nw54_[int(5)] = (_dafny.Map({d_318_v58_: 641})).set(d_318_v58_, self.f33)
                        nw54_[int(6)] = default__.fm17(len(default__.fm4(d_318_v58_, (d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))], d_320_v60_, d_318_v58_, globalState)), self.f33, len(d_332_v71_), d_318_v58_, globalState)
                        nw54_[int(7)] = _dafny.Map({d_318_v58_: -683})
                        nw54_[int(8)] = d_324_v63_
                        nw54_[int(9)] = _dafny.Map({d_318_v58_: len(_dafny.Map({_dafny.Map({d_318_v58_: (d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))]}): len(d_319_v59_)}))})
                        nw54_[int(10)] = d_324_v63_
                        nw54_[int(11)] = d_324_v63_
                        nw54_[int(12)] = d_324_v63_
                        nw54_[int(13)] = d_324_v63_
                        nw54_[int(14)] = d_324_v63_
                        nw54_[int(15)] = _dafny.Map({d_318_v58_: self.f33})
                        nw54_[int(16)] = default__.fm17(len(default__.fm18(globalState)), (self).f24, len(d_333_v72_), d_318_v58_, globalState)
                        nw54_[int(17)] = d_324_v63_
                        nw54_[int(18)] = d_324_v63_
                        nw54_[int(19)] = d_324_v63_
                        nw54_[int(20)] = d_324_v63_
                        nw54_[int(21)] = d_324_v63_
                        nw54_[int(22)] = d_324_v63_
                        nw54_[int(23)] = d_324_v63_
                        nw54_[int(24)] = (d_324_v63_).set(d_318_v58_, (self).f24)
                        nw54_[int(25)] = d_324_v63_
                        nw54_[int(26)] = d_324_v63_
                        d_334_v73_ = nw54_
                        d_335_v74_: _dafny.Map
                        d_335_v74_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_320_v60_ for d_336_i13_ in range(default__.abs(508))])): d_334_v73_})
                        d_337_v75_: _dafny.Seq
                        d_337_v75_ = _dafny.SeqWithoutIsStrInference([d_334_v73_])
                        d_335_v74_ = (d_335_v74_).set((d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))], (d_337_v75_)[default__.safeIndex(self.f33, len(d_337_v75_))])
                        d_338_v76_: D3
                        d_338_v76_ = D3_DC11(self.f33, d_318_v58_, d_318_v58_)
                        d_339_v77_: _dafny.MultiSet
                        d_339_v77_ = _dafny.MultiSet([self.f33, (self).f24, self.f33, (self).fm11(d_318_v58_, not(d_318_v58_), d_318_v58_, self.f33, globalState), (d_338_v76_).cf18])
                        d_340_v78_: _dafny.Set
                        d_340_v78_ = _dafny.Set({d_327_v66_})
                        rhs43_ = d_318_v58_
                        rhs44_ = d_339_v77_
                        rhs45_ = (d_340_v78_).intersection(d_340_v78_)
                        lhs34_ = globalState
                        lhs34_.f2 = rhs43_
                        d_339_v77_ = rhs44_
                        d_340_v78_ = rhs45_
                        d_341_v79_: _dafny.MultiSet
                        d_341_v79_ = _dafny.MultiSet([d_318_v58_, d_318_v58_])
                        d_342_v80_: _dafny.Map
                        d_342_v80_ = _dafny.Map({d_341_v79_: ((self).f24) * ((self).f24)})
                        d_342_v80_ = d_342_v80_
                        d_343_v81_: _dafny.Array
                        def lambda14_(d_344_v58_):
                            def lambda15_(d_345_i14_):
                                return d_344_v58_

                            return lambda15_

                        init9_ = lambda14_(d_318_v58_)
                        nw55_ = _dafny.Array(None, 24)
                        for i0_9_ in range(nw55_.length(0)):
                            nw55_[i0_9_] = init9_(i0_9_)
                        d_343_v81_ = nw55_
                        d_346_v82_: _dafny.Map
                        d_346_v82_ = _dafny.Map({default__.fm19(d_318_v58_, (d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))], globalState): d_343_v81_})
                        d_347_v83_: _dafny.Seq
                        d_347_v83_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_318_v58_])])
                        d_343_v81_ = ((d_346_v82_)[(d_347_v83_) + ((d_347_v83_).set(default__.safeIndex(len(d_319_v59_), len(d_347_v83_)), _dafny.SeqWithoutIsStrInference([True])))] if ((d_347_v83_) + ((d_347_v83_).set(default__.safeIndex(len(d_319_v59_), len(d_347_v83_)), _dafny.SeqWithoutIsStrInference([True])))) in (d_346_v82_) else d_343_v81_)
                    elif True:
                        d_348___mcc_h0_ = source3_.cf16
                        d_349_cf16_ = d_348___mcc_h0_
                        (globalState).f2 = (default__.safeModuloInt(len(d_326_v65_), (0) - ((d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))]))) != ((self).f24)
                        d_350_v84_: _dafny.MultiSet
                        d_350_v84_ = _dafny.MultiSet([d_318_v58_])
                        d_351_v85_: _dafny.Seq
                        d_351_v85_ = _dafny.SeqWithoutIsStrInference([(d_350_v84_).cardinality])
                        d_352_v86_: _dafny.Seq
                        d_352_v86_ = _dafny.SeqWithoutIsStrInference([d_318_v58_, d_318_v58_, d_318_v58_])
                        d_353_v87_: _dafny.Map
                        d_353_v87_ = _dafny.Map({d_352_v86_: d_318_v58_})
                        d_354_v88_: _dafny.Map
                        d_354_v88_ = _dafny.Map({d_320_v60_: d_318_v58_})
                        d_355_v89_: _dafny.MultiSet
                        d_355_v89_ = _dafny.MultiSet([d_318_v58_])
                        d_356_v90_: _dafny.Seq
                        d_356_v90_ = _dafny.SeqWithoutIsStrInference([d_350_v84_, (d_355_v89_)])
                        d_357_v91_: _dafny.Map
                        d_357_v91_ = _dafny.Map({self.f33: (self).f24})
                        d_358_v92_: D1
                        d_358_v92_ = D1_DC4(d_318_v58_, d_318_v58_, d_320_v60_, len(d_357_v91_), (self).f24)
                        d_359_v94_: _dafny.Array
                        nw56_ = _dafny.Array(None, 24)
                        nw56_[int(0)] = ((d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))]) != (self.f33)
                        nw56_[int(1)] = True
                        nw56_[int(2)] = default__.fm2((d_351_v85_)[default__.safeIndex(len(d_319_v59_), len(d_351_v85_))], self.f33, d_320_v60_, globalState)
                        nw56_[int(3)] = ((self).f24) in (d_351_v85_)
                        nw56_[int(4)] = d_318_v58_
                        nw56_[int(5)] = default__.fm2(self.f33, self.f33, _dafny.CodePoint('c'), globalState)
                        nw56_[int(6)] = d_318_v58_
                        nw56_[int(7)] = d_318_v58_
                        nw56_[int(8)] = d_318_v58_
                        nw56_[int(9)] = d_318_v58_
                        nw56_[int(10)] = not(True)
                        nw56_[int(11)] = d_318_v58_
                        nw56_[int(12)] = ((d_353_v87_)[d_352_v86_] if (d_352_v86_) in (d_353_v87_) else d_318_v58_)
                        nw56_[int(13)] = d_318_v58_
                        nw56_[int(14)] = not(default__.fm2((d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))], (d_350_v84_).cardinality, d_320_v60_, globalState))
                        nw56_[int(15)] = d_318_v58_
                        nw56_[int(16)] = d_318_v58_
                        nw56_[int(17)] = d_318_v58_
                        nw56_[int(18)] = (len(d_354_v88_)) <= (((d_356_v90_)[default__.safeIndex((d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))], len(d_356_v90_))]).cardinality)
                        nw56_[int(19)] = default__.fm2(len(d_357_v91_), (d_358_v92_).cf12, d_320_v60_, globalState)
                        nw56_[int(20)] = False
                        def iife15_():
                            coll15_ = _dafny.Map()
                            compr_15_: int
                            for compr_15_ in _dafny.IntegerRange(-804, 751):
                                d_360_v93_: int = compr_15_
                                if ((-804) <= (d_360_v93_)) and ((d_360_v93_) < (751)):
                                    coll15_[default__.safeDivisionInt(d_360_v93_, (d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))])] = len(_dafny.Map({d_318_v58_: d_320_v60_}))
                            return _dafny.Map(coll15_)
                        nw56_[int(21)] = ((d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))]) <= (len((iife15_()
                        ).set((self).f24, (d_327_v66_)[default__.safeIndex(92, (d_327_v66_).length(0))])))
                        nw56_[int(22)] = d_318_v58_
                        nw56_[int(23)] = d_318_v58_
                        d_359_v94_ = nw56_
                        d_359_v94_ = d_359_v94_
                        d_361_v95_: _dafny.Set
                        d_361_v95_ = _dafny.Set({d_352_v86_, d_352_v86_})
                        def iife16_():
                            coll16_ = _dafny.Set()
                            compr_16_: _dafny.Seq
                            for compr_16_ in (default__.fm20(globalState)).keys.Elements:
                                d_362_v96_: _dafny.Seq = compr_16_
                                if (d_362_v96_) in (default__.fm20(globalState)):
                                    coll16_ = coll16_.union(_dafny.Set([d_362_v96_]))
                            return _dafny.Set(coll16_)
                        d_361_v95_ = ((iife16_()
                        ) - (d_361_v95_)) - (d_361_v95_)
                        (globalState).f17 = d_319_v59_
                    (globalState).f1 = (len(d_323_v62_)) * (self.f33)
                    pass
            pass
        (globalState).f1 = len(((d_319_v59_) + (d_319_v59_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndbanqk"))))
        r0 = not(not(d_318_v58_))
        return r0

    def m7(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        (globalState).f18 = self.f33
        d_364_v0_: bool
        d_364_v0_ = True
        d_365_v1_: _dafny.Map
        d_365_v1_ = _dafny.Map({len(p0): d_364_v0_})
        d_366_v2_: _dafny.Seq
        d_366_v2_ = _dafny.SeqWithoutIsStrInference([len(d_365_v1_)])
        d_367_v3_: _dafny.Map
        d_367_v3_ = _dafny.Map({(d_366_v2_) + (d_366_v2_): (self).f24})
        (globalState).f1 = ((d_367_v3_)[d_366_v2_] if (d_366_v2_) in (d_367_v3_) else (self).f24)
        if d_364_v0_:
            r0 = d_364_v0_
            d_368_v4_: str
            d_368_v4_ = _dafny.CodePoint('k')
            d_369_v5_: _dafny.Map
            d_369_v5_ = _dafny.Map({546: p0})
            d_370_v6_: D1
            d_370_v6_ = D1_DC4(d_364_v0_, d_364_v0_, d_368_v4_, (self).f24, len(d_369_v5_))
            pat_let_tv0_ = d_368_v4_
            def iife17_(_pat_let0_0):
                def iife18_(d_371_dt__update__tmp_h0_):
                    def iife19_(_pat_let1_0):
                        def iife20_(d_372_dt__update_hcf10_h0_):
                            return D1_DC4((d_371_dt__update__tmp_h0_).cf8, (d_371_dt__update__tmp_h0_).cf9, d_372_dt__update_hcf10_h0_, (d_371_dt__update__tmp_h0_).cf11, (d_371_dt__update__tmp_h0_).cf12)
                        return iife20_(_pat_let1_0)
                    return iife19_(pat_let_tv0_)
                return iife18_(_pat_let0_0)
            source4_ = iife17_(d_370_v6_)
            if source4_.is_DC4:
                d_373___mcc_h0_ = source4_.cf8
                d_374___mcc_h1_ = source4_.cf9
                d_375___mcc_h2_ = source4_.cf10
                d_376___mcc_h3_ = source4_.cf11
                d_377___mcc_h4_ = source4_.cf12
                d_378_cf12_ = d_377___mcc_h4_
                d_379_cf11_ = d_376___mcc_h3_
                d_380_cf10_ = d_375___mcc_h2_
                d_381_cf9_ = d_374___mcc_h1_
                d_382_cf8_ = d_373___mcc_h0_
                d_383_v7_: _dafny.Array
                nw57_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_383_v7_ = nw57_
                d_383_v7_ = d_383_v7_
                (globalState).f2 = d_364_v0_
                d_384_v8_: D2
                d_384_v8_ = D2_DC8()
                d_385_v9_: _dafny.Set
                d_385_v9_ = _dafny.Set({d_364_v0_})
                d_386_v10_: _dafny.Map
                d_386_v10_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_368_v4_ for d_387_i0_ in range(default__.abs(-689))]): True})
                d_388_v11_: _dafny.Seq
                d_388_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qojo"))
                d_389_v12_: _dafny.Map
                d_389_v12_ = _dafny.Map({not(d_364_v0_): d_382_cf8_})
                d_390_v13_: _dafny.Seq
                d_390_v13_ = _dafny.SeqWithoutIsStrInference([d_389_v12_])
                rhs46_ = d_384_v8_
                rhs47_ = (d_382_cf8_ if (d_385_v9_).issubset(d_385_v9_) else (d_364_v0_ if d_382_cf8_ else d_381_cf9_))
                rhs48_ = (0) - ((len((d_386_v10_).set((d_388_v11_).set(default__.safeIndex(len(_dafny.Map({_dafny.MultiSet(d_366_v2_): self.f33})), len(d_388_v11_)), d_368_v4_), d_381_cf9_))) * (((self).f24) + (len(d_390_v13_))))
                lhs35_ = globalState
                d_384_v8_ = rhs46_
                d_382_cf8_ = rhs47_
                lhs35_.f1 = rhs48_
                d_391_v14_: _dafny.Map
                d_391_v14_ = _dafny.Map({d_380_cf10_: not (d_381_cf9_) or (d_381_cf9_)})
                d_391_v14_ = (d_391_v14_).set(d_368_v4_, d_364_v0_)
            elif source4_.is_DC5:
                d_392___mcc_h5_ = source4_.cf13
                d_393___mcc_h6_ = source4_.cf14
                d_394___mcc_h7_ = source4_.cf15
                d_395_cf15_ = d_394___mcc_h7_
                d_396_cf14_ = d_393___mcc_h6_
                d_397_cf13_ = d_392___mcc_h5_
                d_398_v15_: _dafny.Array
                nw58_ = _dafny.Array(_dafny.MultiSet({}), 22)
                d_398_v15_ = nw58_
                d_399_v16_: _dafny.MultiSet
                d_399_v16_ = _dafny.MultiSet([(p0)[default__.safeIndex(self.f33, len(p0))]])
                index42_ = default__.safeIndex(76, (d_398_v15_).length(0))
                (d_398_v15_)[index42_] = (d_399_v16_ if d_364_v0_ else d_399_v16_)
                d_400_v17_: _dafny.Map
                d_400_v17_ = _dafny.Map({d_364_v0_: d_399_v16_})
                index43_ = default__.safeIndex(76, (d_398_v15_).length(0))
                (d_398_v15_)[index43_] = ((d_400_v17_)[False] if (False) in (d_400_v17_) else _dafny.MultiSet(p0))
                d_401_v18_: _dafny.Map
                d_401_v18_ = _dafny.Map({len(p0): (self).f24})
                d_401_v18_ = (d_401_v18_).set(self.f33, self.f33)
                d_402_v19_: _dafny.Array
                def lambda16_(d_403_v0_):
                    def lambda17_(d_404_i1_):
                        return d_403_v0_

                    return lambda17_

                init10_ = lambda16_(d_364_v0_)
                nw59_ = _dafny.Array(None, 11)
                for i0_10_ in range(nw59_.length(0)):
                    nw59_[i0_10_] = init10_(i0_10_)
                d_402_v19_ = nw59_
                d_405_v20_: _dafny.MultiSet
                d_405_v20_ = _dafny.MultiSet([d_402_v19_])
                d_406_v21_: _dafny.Set
                d_406_v21_ = _dafny.Set({d_364_v0_})
                d_407_v22_: _dafny.Map
                d_407_v22_ = _dafny.Map({d_364_v0_: len(d_406_v21_)})
                d_408_v23_: _dafny.Seq
                d_408_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxpkr"))
                d_409_v24_: int
                out17_: int
                out17_ = default__.m0(((_dafny.MultiSet([d_402_v19_, d_402_v19_]) if d_396_cf14_ else d_405_v20_)).cardinality, (D4_DC13((0) - ((self).f24), d_407_v22_, d_408_v23_)).cf22, globalState)
                d_409_v24_ = out17_
                d_410_v25_: D3
                d_410_v25_ = D3_DC11((self).f24, d_395_cf15_, d_397_cf13_)
                (self).f33 = ((d_410_v25_).cf18) + ((0) - ((0) - (self.f33)))
            elif source4_.is_DC6:
                d_411_v26_: _dafny.Array
                def lambda18_(d_412_v0_):
                    def lambda19_(d_413_i2_):
                        return d_412_v0_

                    return lambda19_

                init11_ = lambda18_(d_364_v0_)
                nw60_ = _dafny.Array(None, 21)
                for i0_11_ in range(nw60_.length(0)):
                    nw60_[i0_11_] = init11_(i0_11_)
                d_411_v26_ = nw60_
                d_414_v27_: _dafny.Map
                d_414_v27_ = _dafny.Map({self.f33: d_411_v26_})
                d_415_v28_: _dafny.Seq
                d_415_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lp"))
                d_416_v29_: _dafny.Map
                d_416_v29_ = _dafny.Map({(d_414_v27_) | (d_414_v27_): d_415_v28_})
                d_417_v30_: _dafny.Seq
                d_417_v30_ = _dafny.SeqWithoutIsStrInference([d_415_v28_])
                d_416_v29_ = (d_416_v29_).set(d_414_v27_, ((d_417_v30_)[default__.safeIndex((0) - (len(d_366_v2_)), len(d_417_v30_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucibb"))))
                d_418_v31_: _dafny.Map
                d_418_v31_ = _dafny.Map({d_364_v0_: d_364_v0_})
                d_418_v31_ = (d_418_v31_).set((d_364_v0_) and (False), d_364_v0_)
                d_419_v32_: C0
                nw61_ = C0()
                nw61_.ctor__(True)
                d_419_v32_ = nw61_
                d_420_v33_: _dafny.Array
                nw62_ = _dafny.Array(_dafny.Set({}), 12)
                d_420_v33_ = nw62_
                d_420_v33_ = d_420_v33_
            elif True:
                d_421___mcc_h8_ = source4_.cf7
                d_422_cf7_ = d_421___mcc_h8_
                d_423_v34_: _dafny.MultiSet
                d_423_v34_ = _dafny.MultiSet([d_364_v0_, d_364_v0_])
                r0 = (d_423_v34_) == ((d_423_v34_) - (d_423_v34_))
                d_424_v35_: _dafny.Map
                d_424_v35_ = _dafny.Map({(d_364_v0_) == (d_364_v0_): ((self).f24) + (len(d_422_cf7_))})
                d_425_v36_: _dafny.Array
                nw63_ = _dafny.Array(False, 29)
                d_425_v36_ = nw63_
                rhs49_ = _dafny.Map({d_364_v0_: len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfkfe"))).set(default__.safeIndex(self.f33, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfkfe")))), d_368_v4_)) + (d_422_cf7_))})
                rhs50_ = _dafny.MultiSet([d_425_v36_])
                rhs51_ = (self).f24
                lhs36_ = globalState
                lhs37_ = globalState
                d_424_v35_ = rhs49_
                lhs36_.f19 = rhs50_
                lhs37_.f1 = rhs51_
                (globalState).f9 = self.f33
                (globalState).f2 = False
            d_426_v37_: _dafny.Array
            nw64_ = _dafny.Array(False, 3)
            d_426_v37_ = nw64_
            index44_ = default__.safeIndex(817, (d_426_v37_).length(0))
            (d_426_v37_)[index44_] = d_364_v0_
            index45_ = default__.safeIndex(817, (d_426_v37_).length(0))
            (d_426_v37_)[index45_] = default__.fm2((self).f24, self.f33, _dafny.CodePoint('w'), globalState)
            (globalState).f18 = (self).f24
            d_427_v38_: D1
            d_427_v38_ = D1_DC5((d_426_v37_)[default__.safeIndex(817, (d_426_v37_).length(0))], d_364_v0_, not(d_364_v0_))
            if (d_427_v38_).cf13:
                d_428_v39_: _dafny.Map
                d_428_v39_ = _dafny.Map({d_364_v0_: self.f33})
                d_429_v40_: D4
                d_429_v40_ = D4_DC13(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pao"))), d_428_v39_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvdcwh")))
                index46_ = default__.safeIndex(817, (d_426_v37_).length(0))
                (d_426_v37_)[index46_] = (default__.fm0(self.f33, len(((d_429_v40_).cf24).set(default__.safeIndex((self).f24, len((d_429_v40_).cf24)), d_368_v4_)), globalState)) <= ((len(d_366_v2_) if (d_426_v37_)[default__.safeIndex(817, (d_426_v37_).length(0))] else self.f33))
                d_365_v1_ = (d_365_v1_).set(self.f33, (not(d_364_v0_)) and (d_364_v0_))
                r0 = not(d_364_v0_)
                d_430_v41_: _dafny.Array
                nw65_ = _dafny.Array(_dafny.Map({}), 6)
                d_430_v41_ = nw65_
                d_431_v42_: _dafny.Map
                d_431_v42_ = _dafny.Map({False: d_364_v0_})
                index47_ = default__.safeIndex(356, (d_430_v41_).length(0))
                (d_430_v41_)[index47_] = d_431_v42_
                d_432_v43_: _dafny.Seq
                d_432_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxjwsbdy"))
                index48_ = default__.safeIndex(356, (d_430_v41_).length(0))
                def iife21_(_pat_let2_0):
                    def iife22_(d_434_dt__update__tmp_h1_):
                        def iife23_(_pat_let3_0):
                            def iife24_(d_435_dt__update_hcf1_h0_):
                                return D0_DC1(d_435_dt__update_hcf1_h0_, (d_434_dt__update__tmp_h1_).cf2, (d_434_dt__update__tmp_h1_).cf3)
                            return iife24_(_pat_let3_0)
                        return iife23_(self.f33)
                    return iife22_(_pat_let2_0)
                rhs52_ = default__.safeDivisionInt((self).f24, self.f33)
                rhs53_ = ((d_432_v43_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))) != ((d_432_v43_) + (_dafny.SeqWithoutIsStrInference([d_368_v4_ for d_433_i3_ in range(default__.abs(-744))])))
                rhs54_ = ((d_431_v42_).set((d_426_v37_)[default__.safeIndex(817, (d_426_v37_).length(0))], (d_426_v37_)[default__.safeIndex(817, (d_426_v37_).length(0))])).set((d_426_v37_)[default__.safeIndex(817, (d_426_v37_).length(0))], d_364_v0_)
                rhs55_ = (iife21_(default__.fm21(globalState))).cf2
                lhs38_ = globalState
                lhs39_ = d_430_v41_
                lhs40_ = default__.safeIndex(356, (d_430_v41_).length(0))
                lhs41_ = globalState
                lhs38_.f18 = rhs52_
                d_364_v0_ = rhs53_
                lhs39_[lhs40_] = rhs54_
                lhs41_.f1 = rhs55_
                d_436_v44_: int
                d_437_v45_: bool
                d_438_v46_: int
                d_439_v47_: str
                out18_: int
                out19_: bool
                out20_: int
                out21_: str
                out18_, out19_, out20_, out21_ = (self).m8(d_426_v37_, globalState)
                d_436_v44_ = out18_
                d_437_v45_ = out19_
                d_438_v46_ = out20_
                d_439_v47_ = out21_
            elif True:
                d_440_v48_: _dafny.Array
                nw66_ = _dafny.Array(int(0), 27)
                d_440_v48_ = nw66_
                index49_ = default__.safeIndex(378, (d_440_v48_).length(0))
                (d_440_v48_)[index49_] = (self).f24
                index50_ = default__.safeIndex(378, (d_440_v48_).length(0))
                rhs56_ = self.f33
                rhs57_ = d_364_v0_
                lhs42_ = d_440_v48_
                lhs43_ = default__.safeIndex(378, (d_440_v48_).length(0))
                lhs42_[lhs43_] = rhs56_
                r2 = rhs57_
                r0 = d_364_v0_
                d_441_v49_: _dafny.Map
                d_441_v49_ = _dafny.Map({d_370_v6_: (d_426_v37_)[default__.safeIndex(817, (d_426_v37_).length(0))]})
                d_442_v50_: _dafny.Array
                nw67_ = _dafny.Array(None, 7)
                nw67_[int(0)] = d_441_v49_
                nw67_[int(1)] = (d_441_v49_) | (d_441_v49_)
                nw67_[int(2)] = d_441_v49_
                nw67_[int(3)] = (d_441_v49_).set(d_370_v6_, (d_426_v37_)[default__.safeIndex(817, (d_426_v37_).length(0))])
                nw67_[int(4)] = (d_441_v49_) | (d_441_v49_)
                nw67_[int(5)] = d_441_v49_
                nw67_[int(6)] = d_441_v49_
                d_442_v50_ = nw67_
                d_442_v50_ = d_442_v50_
                d_443_v51_: C0
                nw68_ = C0()
                nw68_.ctor__(d_364_v0_)
                d_443_v51_ = nw68_
                d_444_v52_: _dafny.Seq
                d_444_v52_ = _dafny.SeqWithoutIsStrInference([d_368_v4_])
                d_445_v53_: D1
                d_445_v53_ = D1_DC3(d_444_v52_)
                d_446_v54_: _dafny.Set
                d_446_v54_ = _dafny.Set({d_445_v53_, d_445_v53_, D1_DC3(d_444_v52_), d_445_v53_, D1_DC3(d_444_v52_)})
                d_443_v51_ = (d_443_v51_ if not((d_446_v54_).issubset(d_446_v54_)) else d_443_v51_)
                (globalState).f18 = ((self.f33 if (d_426_v37_)[default__.safeIndex(817, (d_426_v37_).length(0))] else (self).f24)) + (214)
        elif True:
            d_447_v55_: str
            d_447_v55_ = _dafny.CodePoint('r')
            d_448_v56_: _dafny.Seq
            d_448_v56_ = _dafny.SeqWithoutIsStrInference([d_447_v55_])
            d_449_v57_: _dafny.Set
            d_449_v57_ = _dafny.Set({d_448_v56_, d_448_v56_})
            d_449_v57_ = d_449_v57_
            rhs58_ = d_364_v0_
            rhs59_ = p0
            rhs60_ = (0) - (self.f33)
            rhs61_ = (default__.fm22(self.f33, d_364_v0_, d_364_v0_, default__.fm0(self.f33, (self).f24, globalState), globalState)).intersection((d_449_v57_) | (_dafny.Set({d_448_v56_, _dafny.SeqWithoutIsStrInference([d_447_v55_ for d_450_i4_ in range(default__.abs(-411))])})))
            rhs62_ = d_447_v55_
            lhs44_ = globalState
            lhs45_ = globalState
            r2 = rhs58_
            lhs44_.f21 = rhs59_
            lhs45_.f18 = rhs60_
            d_449_v57_ = rhs61_
            d_447_v55_ = rhs62_
            if d_364_v0_:
                d_451_v58_: _dafny.MultiSet
                d_451_v58_ = _dafny.MultiSet([self.f33, self.f33])
                d_452_v59_: _dafny.Array
                nw69_ = _dafny.Array(None, 13)
                nw69_[int(0)] = p0
                nw69_[int(1)] = p0
                nw69_[int(2)] = p0
                nw69_[int(3)] = p0
                nw69_[int(4)] = (p0) + (p0)
                nw69_[int(5)] = _dafny.SeqWithoutIsStrInference([d_364_v0_, False, d_364_v0_])
                nw69_[int(6)] = (p0) + (_dafny.SeqWithoutIsStrInference([d_364_v0_]))
                nw69_[int(7)] = p0
                nw69_[int(8)] = p0
                nw69_[int(9)] = default__.fm23(globalState)
                nw69_[int(10)] = default__.fm23(globalState)
                nw69_[int(11)] = p0
                nw69_[int(12)] = (p0).set(default__.safeIndex(((d_451_v58_)[len(p0)] if (len(p0)) in (d_451_v58_) else self.f33), len(p0)), d_364_v0_)
                d_452_v59_ = nw69_
                index51_ = default__.safeIndex(49, (d_452_v59_).length(0))
                (d_452_v59_)[index51_] = p0
                index52_ = default__.safeIndex(49, (d_452_v59_).length(0))
                (d_452_v59_)[index52_] = p0
                d_453_v60_: _dafny.Seq
                d_453_v60_ = _dafny.SeqWithoutIsStrInference([d_448_v56_])
                (globalState).f2 = default__.fm2(default__.fm0(len((d_453_v60_)[default__.safeIndex(self.f33, len(d_453_v60_))]), (D3_DC11(self.f33, ((d_452_v59_)[default__.safeIndex(49, (d_452_v59_).length(0))])[default__.safeIndex((self).f24, len((d_452_v59_)[default__.safeIndex(49, (d_452_v59_).length(0))]))], d_364_v0_)).cf18, globalState), self.f33, default__.fm3(globalState), globalState)
                d_454_v61_: D4
                d_454_v61_ = D4_DC13(len(_dafny.SeqWithoutIsStrInference([d_447_v55_ for d_455_i5_ in range(default__.abs(-751))])), _dafny.Map({True: self.f33}), d_448_v56_)
                d_456_v62_: _dafny.Map
                d_456_v62_ = _dafny.Map({(d_454_v61_).cf22: -176})
                d_456_v62_ = (d_456_v62_).set(self.f33, (len(d_448_v56_) if d_364_v0_ else (self).f24))
                d_457_v63_: _dafny.Array
                nw70_ = _dafny.Array(int(0), 26)
                d_457_v63_ = nw70_
                d_458_v64_: _dafny.Map
                d_458_v64_ = _dafny.Map({self.f33: _dafny.Map({self.f33: d_364_v0_})})
                index53_ = default__.safeIndex(389, (d_457_v63_).length(0))
                (d_457_v63_)[index53_] = default__.safeDivisionInt(self.f33, len(d_458_v64_))
                index54_ = default__.safeIndex(389, (d_457_v63_).length(0))
                (d_457_v63_)[index54_] = self.f33
                d_459_v65_: C0
                nw71_ = C0()
                nw71_.ctor__(d_364_v0_)
                d_459_v65_ = nw71_
            elif True:
                d_460_v67_: _dafny.Map
                d_460_v67_ = _dafny.Map({d_364_v0_: self.f33})
                d_461_v68_: _dafny.MultiSet
                def iife25_():
                    coll17_ = _dafny.Map()
                    compr_17_: int
                    for compr_17_ in (d_366_v2_).Elements:
                        d_462_v66_: int = compr_17_
                        if (d_462_v66_) in (d_366_v2_):
                            coll17_[default__.safeDivisionInt(d_462_v66_, self.f33)] = d_364_v0_
                    return _dafny.Map(coll17_)
                d_461_v68_ = _dafny.MultiSet([iife25_()
                , _dafny.Map({(self).f24: d_364_v0_}), d_365_v1_, (d_365_v1_).set(((d_460_v67_)[True] if (True) in (d_460_v67_) else self.f33), d_364_v0_), d_365_v1_])
                d_463_v69_: D4
                d_463_v69_ = D4_DC12(d_366_v2_)
                d_464_v70_: D3
                d_464_v70_ = D3_DC11(self.f33, d_364_v0_, d_364_v0_)
                d_465_v71_: _dafny.MultiSet
                d_465_v71_ = _dafny.MultiSet([d_364_v0_, (d_464_v70_).cf19, d_364_v0_])
                d_466_v72_: _dafny.Array
                nw72_ = _dafny.Array(None, 17)
                nw72_[int(0)] = self.f33
                nw72_[int(1)] = (self).f24
                nw72_[int(2)] = ((0) - (self.f33)) * (self.f33)
                nw72_[int(3)] = self.f33
                nw72_[int(4)] = (len((d_463_v69_).cf21) if default__.fm2(self.f33, ((d_461_v68_)[d_365_v1_] if (d_365_v1_) in (d_461_v68_) else self.f33), d_447_v55_, globalState) else self.f33)
                nw72_[int(5)] = ((self).f24) * (self.f33)
                nw72_[int(6)] = (d_465_v71_).cardinality
                nw72_[int(7)] = self.f33
                nw72_[int(8)] = 100
                nw72_[int(9)] = self.f33
                nw72_[int(10)] = self.f33
                nw72_[int(11)] = (self).f24
                nw72_[int(12)] = len(d_448_v56_)
                nw72_[int(13)] = self.f33
                nw72_[int(14)] = 856
                nw72_[int(15)] = (len(d_448_v56_)) + ((self).f24)
                nw72_[int(16)] = 946
                d_466_v72_ = nw72_
                d_466_v72_ = d_466_v72_
                index55_ = default__.safeIndex(611, (d_466_v72_).length(0))
                (d_466_v72_)[index55_] = self.f33
                index56_ = default__.safeIndex(611, (d_466_v72_).length(0))
                (d_466_v72_)[index56_] = (self.f33) + (802)
                d_467_v73_: _dafny.Array
                nw73_ = _dafny.Array(False, 22)
                d_467_v73_ = nw73_
                d_468_v74_: _dafny.Set
                d_468_v74_ = _dafny.Set({d_467_v73_})
                d_469_v75_: D1
                d_469_v75_ = D1_DC5(d_364_v0_, (d_468_v74_).issubset(d_468_v74_), d_364_v0_)
                d_469_v75_ = d_469_v75_
                d_470_v76_: _dafny.Array
                nw74_ = _dafny.Array(D4.default()(), 9)
                d_470_v76_ = nw74_
                index57_ = default__.safeIndex(935, (d_470_v76_).length(0))
                (d_470_v76_)[index57_] = d_463_v69_
                d_471_v77_: D1
                d_471_v77_ = D1_DC4(d_364_v0_, d_364_v0_, d_447_v55_, self.f33, (self).f24)
                pat_let_tv1_ = d_366_v2_
                index58_ = default__.safeIndex(935, (d_470_v76_).length(0))
                def iife26_(_pat_let4_0):
                    def iife27_(d_472_dt__update__tmp_h2_):
                        def iife28_(_pat_let5_0):
                            def iife29_(d_473_dt__update_hcf21_h0_):
                                return D4_DC12(d_473_dt__update_hcf21_h0_)
                            return iife29_(_pat_let5_0)
                        return iife28_(pat_let_tv1_)
                    return iife27_(_pat_let4_0)
                rhs63_ = iife26_(d_463_v69_)
                rhs64_ = ((d_460_v67_) | (d_460_v67_) if (d_471_v77_).cf9 else d_460_v67_)
                lhs46_ = d_470_v76_
                lhs47_ = default__.safeIndex(935, (d_470_v76_).length(0))
                lhs46_[lhs47_] = rhs63_
                d_460_v67_ = rhs64_
                d_474_v78_: _dafny.Map
                d_474_v78_ = _dafny.Map({d_364_v0_: default__.fm24((d_366_v2_)[default__.safeIndex((self).f24, len(d_366_v2_))], not(d_364_v0_), (d_466_v72_)[default__.safeIndex(611, (d_466_v72_).length(0))], (0) - ((self).f24), globalState)})
                d_474_v78_ = (d_474_v78_).set(d_364_v0_, d_366_v2_)
            d_475_v79_: _dafny.Array
            nw75_ = _dafny.Array(None, 7)
            nw75_[int(0)] = d_447_v55_
            nw75_[int(1)] = d_447_v55_
            nw75_[int(2)] = d_447_v55_
            nw75_[int(3)] = _dafny.CodePoint('p')
            nw75_[int(4)] = (d_448_v56_)[default__.safeIndex((self).f24, len(d_448_v56_))]
            nw75_[int(5)] = d_447_v55_
            nw75_[int(6)] = d_447_v55_
            d_475_v79_ = nw75_
            d_476_v80_: _dafny.Seq
            d_476_v80_ = _dafny.SeqWithoutIsStrInference([d_475_v79_])
            d_477_v81_: D2
            d_477_v81_ = D2_DC7((d_476_v80_)[default__.safeIndex(234, len(d_476_v80_))])
            source5_ = d_477_v81_
            if source5_.is_DC8:
                (globalState).f18 = (self.f33) + (((self).f24 if d_364_v0_ else (self).f24))
                d_478_v82_: _dafny.Array
                nw76_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_478_v82_ = nw76_
                d_478_v82_ = d_478_v82_
                d_479_v83_: C0
                nw77_ = C0()
                nw77_.ctor__((d_364_v0_ if d_364_v0_ else d_364_v0_))
                d_479_v83_ = nw77_
                d_364_v0_ = ((self).f24) < ((self).f24)
            elif True:
                d_480___mcc_h9_ = source5_.cf16
                d_481_cf16_ = d_480___mcc_h9_
                d_482_v84_: D1
                d_482_v84_ = D1_DC5(d_364_v0_, d_364_v0_, True)
                pat_let_tv2_ = d_364_v0_
                d_483_v85_: _dafny.Map
                def iife30_(_pat_let6_0):
                    def iife31_(d_484_dt__update__tmp_h3_):
                        def iife32_(_pat_let7_0):
                            def iife33_(d_485_dt__update_hcf13_h0_):
                                return D1_DC5(d_485_dt__update_hcf13_h0_, (d_484_dt__update__tmp_h3_).cf14, (d_484_dt__update__tmp_h3_).cf15)
                            return iife33_(_pat_let7_0)
                        return iife32_(pat_let_tv2_)
                    return iife31_(_pat_let6_0)
                d_483_v85_ = _dafny.Map({(d_366_v2_)[default__.safeIndex((0) - (self.f33), len(d_366_v2_))]: iife30_(d_482_v84_)})
                d_486_v86_: _dafny.Map
                d_486_v86_ = _dafny.Map({self.f33: d_483_v85_})
                def iife34_():
                    coll18_ = _dafny.Map()
                    compr_18_: int
                    for compr_18_ in (default__.fm18(globalState)).Elements:
                        d_487_v87_: int = compr_18_
                        if (d_487_v87_) in (default__.fm18(globalState)):
                            coll18_[(d_487_v87_) * (self.f33)] = (_dafny.Map({(0) - ((0) - (self.f33)): d_482_v84_})) | (d_483_v85_)
                    return _dafny.Map(coll18_)
                d_486_v86_ = iife34_()
                
                d_488_v88_: C0
                nw78_ = C0()
                nw78_.ctor__(d_364_v0_)
                d_488_v88_ = nw78_
                (globalState).f17 = d_448_v56_
                d_489_v89_: _dafny.Map
                d_489_v89_ = _dafny.Map({d_448_v56_: not(d_364_v0_)})
                d_489_v89_ = (d_489_v89_).set(_dafny.SeqWithoutIsStrInference([d_447_v55_ for d_490_i6_ in range(default__.abs(901))]), ((d_488_v88_).f34) == ((d_488_v88_).f34))
            (globalState).f14 = (0) - ((self).f24)
        d_491_v90_: _dafny.Seq
        d_491_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frfeqaonm"))
        d_492_v91_: D4
        d_492_v91_ = D4_DC13(self.f33, _dafny.Map({d_364_v0_: self.f33}), d_491_v90_)
        source6_ = d_492_v91_
        if source6_.is_DC13:
            d_493___mcc_h10_ = source6_.cf22
            d_494___mcc_h11_ = source6_.cf23
            d_495___mcc_h12_ = source6_.cf24
            d_496_cf24_ = d_495___mcc_h12_
            d_497_cf23_ = d_494___mcc_h11_
            d_498_cf22_ = d_493___mcc_h10_
            (globalState).f2 = d_364_v0_
            r1 = (0) - (self.f33)
            if d_364_v0_:
                d_499_v92_: str
                d_499_v92_ = _dafny.CodePoint('x')
                rhs65_ = default__.fm2(self.f33, self.f33, d_499_v92_, globalState)
                rhs66_ = (self).fm11(d_364_v0_, d_364_v0_, ((self).f24) != (self.f33), (default__.fm25(globalState)).cf22, globalState)
                rhs67_ = d_364_v0_
                lhs48_ = self
                lhs49_ = globalState
                d_364_v0_ = rhs65_
                lhs48_.f33 = rhs66_
                lhs49_.f2 = rhs67_
                d_365_v1_ = (d_365_v1_).set(d_498_cf22_, True)
                (globalState).f2 = False
                r2 = (d_364_v0_) == (not(d_364_v0_))
                d_500_v93_: _dafny.Array
                nw79_ = _dafny.Array(False, 1)
                d_500_v93_ = nw79_
                d_501_v94_: _dafny.Set
                d_501_v94_ = _dafny.Set({d_500_v93_, d_500_v93_, d_500_v93_, d_500_v93_, d_500_v93_})
                d_501_v94_ = (_dafny.Set({d_500_v93_, d_500_v93_, d_500_v93_})).intersection(d_501_v94_)
            elif True:
                d_502_v95_: str
                d_502_v95_ = _dafny.CodePoint('l')
                d_503_v96_: _dafny.MultiSet
                d_503_v96_ = _dafny.MultiSet([(d_365_v1_) | (d_365_v1_)])
                rhs68_ = d_502_v95_
                rhs69_ = d_503_v96_
                d_502_v95_ = rhs68_
                d_503_v96_ = rhs69_
                (globalState).f2 = ((0) - ((self).f24)) < (self.f33)
                d_504_v97_: _dafny.Map
                d_504_v97_ = _dafny.Map({D1_DC3(d_496_cf24_): False})
                d_505_v98_: D1
                d_505_v98_ = D1_DC3(d_491_v90_)
                d_504_v97_ = (_dafny.Map({d_505_v98_: d_364_v0_})) | (d_504_v97_)
                d_364_v0_ = (self).fm10(globalState)
                (self).f33 = self.f33
            d_506_v99_: str
            d_506_v99_ = _dafny.CodePoint('n')
            d_506_v99_ = d_506_v99_
        elif source6_.is_DC12:
            d_507___mcc_h13_ = source6_.cf21
            d_508_cf21_ = d_507___mcc_h13_
            (globalState).f2 = d_364_v0_
            (globalState).f2 = ((default__.fm23(globalState)) + (p0)) == (p0)
            (globalState).f18 = ((self).f24) - (self.f33)
            d_509_v100_: _dafny.Array
            nw80_ = _dafny.Array(None, 7)
            nw80_[int(0)] = False
            nw80_[int(1)] = d_364_v0_
            nw80_[int(2)] = (d_364_v0_) == (d_364_v0_)
            nw80_[int(3)] = d_364_v0_
            nw80_[int(4)] = d_364_v0_
            nw80_[int(5)] = not((d_491_v90_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lruyipny"))))
            nw80_[int(6)] = True
            d_509_v100_ = nw80_
            index59_ = default__.safeIndex(572, (d_509_v100_).length(0))
            (d_509_v100_)[index59_] = d_364_v0_
            index60_ = default__.safeIndex(572, (d_509_v100_).length(0))
            (d_509_v100_)[index60_] = d_364_v0_
        elif True:
            d_510___mcc_h14_ = source6_.cf25
            d_511_cf25_ = d_510___mcc_h14_
            d_512_v101_: _dafny.Array
            nw81_ = _dafny.Array(_dafny.Map({}), 3)
            d_512_v101_ = nw81_
            d_513_v102_: _dafny.Map
            d_513_v102_ = _dafny.Map({d_364_v0_: d_364_v0_})
            index61_ = default__.safeIndex(638, (d_512_v101_).length(0))
            (d_512_v101_)[index61_] = d_513_v102_
            d_514_v103_: _dafny.MultiSet
            d_514_v103_ = _dafny.MultiSet([(self).f24, self.f33])
            index62_ = default__.safeIndex(638, (d_512_v101_).length(0))
            rhs70_ = (0) - ((self).f24)
            rhs71_ = d_513_v102_
            rhs72_ = ((d_514_v103_).intersection(d_514_v103_)) - (d_514_v103_)
            lhs50_ = globalState
            lhs51_ = d_512_v101_
            lhs52_ = default__.safeIndex(638, (d_512_v101_).length(0))
            lhs50_.f9 = rhs70_
            lhs51_[lhs52_] = rhs71_
            d_514_v103_ = rhs72_
            d_515_v104_: _dafny.MultiSet
            d_515_v104_ = _dafny.MultiSet([default__.fm26(globalState)])
            d_516_v105_: _dafny.Array
            def lambda20_(d_517_v0_):
                def lambda21_(d_518_i7_):
                    return d_517_v0_

                return lambda21_

            init12_ = lambda20_(d_364_v0_)
            nw82_ = _dafny.Array(None, 22)
            for i0_12_ in range(nw82_.length(0)):
                nw82_[i0_12_] = init12_(i0_12_)
            d_516_v105_ = nw82_
            d_519_v106_: _dafny.Map
            d_519_v106_ = _dafny.Map({d_516_v105_: (self).f24})
            d_520_v107_: D4
            d_520_v107_ = D4_DC13(((d_519_v106_)[d_516_v105_] if (d_516_v105_) in (d_519_v106_) else (_dafny.MultiSet([self.f33])).cardinality), default__.fm17((self).f24, len(d_366_v2_), len(_dafny.Map({d_364_v0_: d_364_v0_})), True, globalState), d_491_v90_)
            d_521_v108_: D4
            d_521_v108_ = D4_DC14(d_520_v107_)
            d_522_v109_: D4
            d_522_v109_ = D4_DC14(d_520_v107_)
            d_523_v110_: D4
            d_523_v110_ = D4_DC14(d_521_v108_)
            d_524_v111_: _dafny.Map
            d_524_v111_ = _dafny.Map({97: -620})
            d_525_v112_: str
            d_525_v112_ = _dafny.CodePoint('y')
            d_526_v113_: _dafny.Array
            nw83_ = _dafny.Array(None, 12)
            nw83_[int(0)] = ((self).f24) + ((d_366_v2_)[default__.safeIndex((0) - (((d_515_v104_).set(d_523_v110_, default__.abs(self.f33))).cardinality), len(d_366_v2_))])
            nw83_[int(1)] = self.f33
            nw83_[int(2)] = (self).f24
            nw83_[int(3)] = (0) - ((self.f33) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgboi")))))
            nw83_[int(4)] = ((d_514_v103_).cardinality) - (len(d_524_v111_))
            nw83_[int(5)] = (self).f24
            nw83_[int(6)] = len(d_365_v1_)
            nw83_[int(7)] = (self).f24
            nw83_[int(8)] = self.f33
            nw83_[int(9)] = (self).f24
            nw83_[int(10)] = (self).f24
            nw83_[int(11)] = len((d_491_v90_).set(default__.safeIndex(self.f33, len(d_491_v90_)), d_525_v112_))
            d_526_v113_ = nw83_
            d_526_v113_ = d_526_v113_
            d_527_v114_: D1
            d_527_v114_ = D1_DC6()
            d_527_v114_ = default__.fm27(globalState)
            d_528_v115_: _dafny.Map
            d_528_v115_ = _dafny.Map({d_525_v112_: (self).f24})
            d_529_v116_: _dafny.Map
            d_529_v116_ = _dafny.Map({d_528_v115_: (self).f24})
            d_530_v117_: _dafny.Set
            d_530_v117_ = _dafny.Set({d_364_v0_, d_364_v0_})
            d_531_v118_: _dafny.Seq
            d_531_v118_ = _dafny.SeqWithoutIsStrInference([(d_529_v116_).set(d_528_v115_, (self).f24), d_529_v116_, _dafny.Map({d_528_v115_: len(d_530_v117_)})])
            d_529_v116_ = ((d_531_v118_)[default__.safeIndex(22, len(d_531_v118_))]) | (d_529_v116_)
        (globalState).f14 = (self).f24
        (globalState).f14 = self.f33
        r0 = d_364_v0_
        r1 = (self).f24
        d_532_v119_: _dafny.Set
        d_532_v119_ = _dafny.Set({self.f33})
        def iife35_():
            coll19_ = _dafny.Set()
            compr_19_: int
            for compr_19_ in (d_366_v2_).Elements:
                d_533_v120_: int = compr_19_
                if (d_533_v120_) in (d_366_v2_):
                    coll19_ = coll19_.union(_dafny.Set([default__.safeDivisionInt(d_533_v120_, len(_dafny.Set({_dafny.Set({True, True})})))]))
            return _dafny.Set(coll19_)
        r2 = (iife35_()
        ).issubset(d_532_v119_)
        return r0, r1, r2

    def m8(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        r3: str = _dafny.CodePoint('D')
        d_534_v0_: bool
        d_534_v0_ = False
        r1 = (d_534_v0_ if d_534_v0_ else (d_534_v0_ if d_534_v0_ else d_534_v0_))
        d_535_v1_: _dafny.Map
        d_535_v1_ = _dafny.Map({319: (0) - ((self).f24)})
        d_536_v2_: _dafny.Seq
        d_536_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nedubbppa"))
        d_537_v3_: _dafny.Array
        nw84_ = _dafny.Array(int(0), 14)
        d_537_v3_ = nw84_
        d_538_v4_: _dafny.Set
        d_538_v4_ = _dafny.Set({d_537_v3_, d_537_v3_, d_537_v3_})
        d_539_v5_: _dafny.Set
        d_539_v5_ = _dafny.Set({len(d_536_v2_)})
        d_540_v6_: D6
        d_540_v6_ = D6_DC16(d_539_v5_)
        d_541_v7_: _dafny.Seq
        d_541_v7_ = _dafny.SeqWithoutIsStrInference([len((d_540_v6_).cf27)])
        d_542_v8_: _dafny.Seq
        d_542_v8_ = _dafny.SeqWithoutIsStrInference([d_534_v0_, d_534_v0_])
        d_543_v9_: _dafny.Array
        nw85_ = _dafny.Array(None, 26)
        nw85_[int(0)] = (self).f24
        nw85_[int(1)] = (self).f24
        nw85_[int(2)] = (self).f24
        nw85_[int(3)] = ((d_535_v1_)[self.f33] if (self.f33) in (d_535_v1_) else self.f33)
        nw85_[int(4)] = (self).f24
        nw85_[int(5)] = (self).f24
        nw85_[int(6)] = len(d_536_v2_)
        nw85_[int(7)] = len(d_536_v2_)
        nw85_[int(8)] = (self).f24
        nw85_[int(9)] = len(d_538_v4_)
        nw85_[int(10)] = default__.fm0((self).f24, 580, globalState)
        nw85_[int(11)] = self.f33
        nw85_[int(12)] = (self).f24
        nw85_[int(13)] = (self).f24
        nw85_[int(14)] = self.f33
        nw85_[int(15)] = len(d_541_v7_)
        nw85_[int(16)] = self.f33
        nw85_[int(17)] = (0) - (len(d_542_v8_))
        nw85_[int(18)] = -528
        nw85_[int(19)] = self.f33
        nw85_[int(20)] = (0) - (len(d_541_v7_))
        nw85_[int(21)] = (self).f24
        nw85_[int(22)] = (self).f24
        nw85_[int(23)] = self.f33
        nw85_[int(24)] = 108
        nw85_[int(25)] = self.f33
        d_543_v9_ = nw85_
        d_544_v10_: _dafny.MultiSet
        d_544_v10_ = _dafny.MultiSet([d_543_v9_])
        if not((d_544_v10_) == (d_544_v10_)):
            d_545_v11_: _dafny.Map
            d_545_v11_ = _dafny.Map({not(d_534_v0_): len((d_542_v8_) + (d_542_v8_))})
            d_545_v11_ = (d_545_v11_).set(d_534_v0_, (self).f24)
            d_546_v12_: _dafny.Array
            nw86_ = _dafny.Array(_dafny.Map({}), 14)
            d_546_v12_ = nw86_
            d_547_v13_: D4
            d_547_v13_ = D4_DC14(D4_DC13(self.f33, d_545_v11_, d_536_v2_))
            d_548_v14_: _dafny.Map
            d_548_v14_ = _dafny.Map({d_547_v13_: d_536_v2_})
            index63_ = default__.safeIndex(497, (d_546_v12_).length(0))
            (d_546_v12_)[index63_] = d_548_v14_
            d_549_v15_: D7
            d_549_v15_ = D7_DC20(d_548_v14_)
            index64_ = default__.safeIndex(497, (d_546_v12_).length(0))
            (d_546_v12_)[index64_] = (d_549_v15_).cf31
            (globalState).f18 = self.f33
            (globalState).f2 = d_534_v0_
            d_550_v16_: str
            d_550_v16_ = _dafny.CodePoint('u')
            d_551_v17_: _dafny.Map
            d_551_v17_ = _dafny.Map({d_534_v0_: d_534_v0_})
            d_552_v18_: _dafny.MultiSet
            d_552_v18_ = _dafny.MultiSet([len(d_551_v17_), (self).f24])
            d_553_v19_: _dafny.MultiSet
            d_553_v19_ = _dafny.MultiSet([(len(_dafny.Map({d_550_v16_: not(d_534_v0_)}))) == ((d_552_v18_).cardinality)])
            d_553_v19_ = default__.fm14(d_535_v1_, self.f33, globalState)
        elif True:
            (globalState).f2 = (d_542_v8_) != (d_542_v8_)
            if ((self).f24) != (default__.fm0((self).f24, (self).f24, globalState)):
                d_554_v20_: _dafny.Set
                d_554_v20_ = _dafny.Set({d_534_v0_, not(d_534_v0_), d_534_v0_, d_534_v0_, d_534_v0_})
                d_555_v21_: _dafny.Set
                d_555_v21_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len(d_536_v2_), len(d_554_v20_), self.f33])})
                d_556_v22_: _dafny.Array
                nw87_ = _dafny.Array(_dafny.Seq({}), 10)
                d_556_v22_ = nw87_
                index65_ = default__.safeIndex(92, (d_556_v22_).length(0))
                (d_556_v22_)[index65_] = default__.fm24(self.f33, d_534_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dn"))), self.f33, globalState)
                index66_ = default__.safeIndex(865, (p0).length(0))
                (p0)[index66_] = (len(d_539_v5_)) == ((self).f24)
                index67_ = default__.safeIndex(92, (d_556_v22_).length(0))
                index68_ = default__.safeIndex(865, (p0).length(0))
                rhs73_ = default__.fm24((0) - (self.f33), d_534_v0_, len((d_536_v2_) + (d_536_v2_)), 151, globalState)
                rhs74_ = d_555_v21_
                rhs75_ = self.f33
                rhs76_ = _dafny.SeqWithoutIsStrInference([(self).f24, 829, default__.safeModuloInt(self.f33, len(d_536_v2_))])
                rhs77_ = d_534_v0_
                lhs53_ = self
                lhs54_ = d_556_v22_
                lhs55_ = default__.safeIndex(92, (d_556_v22_).length(0))
                lhs56_ = p0
                lhs57_ = default__.safeIndex(865, (p0).length(0))
                d_541_v7_ = rhs73_
                d_555_v21_ = rhs74_
                lhs53_.f33 = rhs75_
                lhs54_[lhs55_] = rhs76_
                lhs56_[lhs57_] = rhs77_
                d_557_v23_: C0
                nw88_ = C0()
                nw88_.ctor__((self.f33) <= (self.f33))
                d_557_v23_ = nw88_
                r1 = (p0)[default__.safeIndex(865, (p0).length(0))]
                d_558_v24_: _dafny.Map
                d_558_v24_ = _dafny.Map({((d_556_v22_)[default__.safeIndex(92, (d_556_v22_).length(0))]).set(default__.safeIndex(self.f33, len((d_556_v22_)[default__.safeIndex(92, (d_556_v22_).length(0))])), self.f33): D6_DC17(d_534_v0_)})
                d_559_v25_: D6
                d_559_v25_ = D6_DC17(not((d_557_v23_).f34))
                d_558_v24_ = (d_558_v24_).set(d_541_v7_, (d_559_v25_ if (p0)[default__.safeIndex(865, (p0).length(0))] else d_559_v25_))
                d_560_v26_: C0
                nw89_ = C0()
                nw89_.ctor__((d_557_v23_).f34)
                d_560_v26_ = nw89_
            elif True:
                d_561_v27_: _dafny.Set
                d_561_v27_ = _dafny.Set({d_534_v0_})
                rhs78_ = d_534_v0_
                rhs79_ = (_dafny.Set({d_534_v0_})) - (d_561_v27_)
                rhs80_ = False
                lhs58_ = globalState
                d_534_v0_ = rhs78_
                d_561_v27_ = rhs79_
                lhs58_.f2 = rhs80_
                r1 = d_534_v0_
                d_562_v28_: D1
                d_562_v28_ = D1_DC5(d_534_v0_, d_534_v0_, d_534_v0_)
                d_563_v29_: D6
                d_563_v29_ = D6_DC17(d_534_v0_)
                pat_let_tv3_ = d_534_v0_
                d_564_v30_: D7
                def iife36_(_pat_let8_0):
                    def iife37_(d_565_dt__update__tmp_h0_):
                        def iife38_(_pat_let9_0):
                            def iife39_(d_566_dt__update_hcf15_h0_):
                                return D1_DC5((d_565_dt__update__tmp_h0_).cf13, (d_565_dt__update__tmp_h0_).cf14, d_566_dt__update_hcf15_h0_)
                            return iife39_(_pat_let9_0)
                        return iife38_(pat_let_tv3_)
                    return iife37_(_pat_let8_0)
                d_564_v30_ = D7_DC21(iife36_(d_562_v28_), D6_DC19(d_563_v29_), p0)
                d_567_v31_: _dafny.Map
                d_567_v31_ = _dafny.Map({d_534_v0_: d_564_v30_})
                d_568_v32_: _dafny.MultiSet
                d_568_v32_ = _dafny.MultiSet([d_534_v0_])
                d_569_v33_: _dafny.Map
                d_569_v33_ = _dafny.Map({d_534_v0_: True})
                rhs81_ = ((len(d_542_v8_)) - (self.f33)) - (((d_568_v32_)[d_534_v0_] if (d_534_v0_) in (d_568_v32_) else self.f33))
                rhs82_ = len(d_542_v8_)
                rhs83_ = (_dafny.Map({d_534_v0_: d_564_v30_})) | (_dafny.Map({not(d_534_v0_): d_564_v30_}))
                rhs84_ = ((self).f24) >= ((0) - ((len(d_569_v33_)) - ((self).f24)))
                lhs59_ = globalState
                lhs60_ = globalState
                lhs61_ = globalState
                lhs59_.f9 = rhs81_
                lhs60_.f18 = rhs82_
                d_567_v31_ = rhs83_
                lhs61_.f2 = rhs84_
                d_570_v34_: _dafny.Map
                d_570_v34_ = _dafny.Map({d_542_v8_: (self).f24})
                d_570_v34_ = (d_570_v34_).set((d_542_v8_) + (d_542_v8_), 997)
                (globalState).f1 = (0) - (-374)
            d_543_v9_ = d_543_v9_
            index69_ = default__.safeIndex(710, (p0).length(0))
            (p0)[index69_] = (d_534_v0_) or (d_534_v0_)
            index70_ = default__.safeIndex(710, (p0).length(0))
            (p0)[index70_] = not(not(d_534_v0_))
            d_571_v35_: _dafny.Set
            d_571_v35_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(d_536_v2_)[default__.safeIndex(self.f33, len(d_536_v2_))] for d_572_i0_ in range(default__.abs(841))])})
            d_573_v36_: C0
            nw90_ = C0()
            nw90_.ctor__(not((d_571_v35_).isdisjoint(d_571_v35_)))
            d_573_v36_ = nw90_
        (globalState).f2 = (d_538_v4_).isdisjoint(d_538_v4_)
        d_574_v37_: D1
        d_574_v37_ = D1_DC5(d_534_v0_, True, d_534_v0_)
        d_575_v38_: D6
        d_575_v38_ = D6_DC16(_dafny.Set({self.f33}))
        d_576_v39_: D7
        d_576_v39_ = D7_DC21(d_574_v37_, D6_DC19(d_575_v38_), p0)
        d_577_v40_: D7
        d_577_v40_ = D7_DC22(d_576_v39_)
        d_578_v41_: _dafny.Map
        d_578_v41_ = _dafny.Map({d_577_v40_: p0})
        d_578_v41_ = (d_578_v41_).set(D7_DC22(d_576_v39_), p0)
        d_579_v42_: _dafny.MultiSet
        d_579_v42_ = _dafny.MultiSet([(self).f24, (0) - (len(d_539_v5_))])
        d_580_v43_: _dafny.Set
        d_580_v43_ = _dafny.Set({((d_579_v42_).cardinality) < (-939), d_534_v0_, (d_534_v0_ if not(d_534_v0_) else d_534_v0_), d_534_v0_, d_534_v0_})
        d_581_v44_: _dafny.MultiSet
        d_581_v44_ = _dafny.MultiSet([(default__.fm28(d_534_v0_, len(d_536_v2_), d_534_v0_, len(_dafny.Map({d_534_v0_: self.f33})), globalState)).ispropersubset(d_580_v43_), d_534_v0_, d_534_v0_])
        d_582_v45_: D8
        d_582_v45_ = D8_DC23(_dafny.Set({d_534_v0_, d_534_v0_, d_534_v0_, not(d_534_v0_)}))
        rhs85_ = (d_582_v45_).cf36
        rhs86_ = ((830) + (self.f33)) + (689)
        rhs87_ = (d_581_v44_).set(d_534_v0_, default__.abs((self).f24))
        lhs62_ = globalState
        d_580_v43_ = rhs85_
        lhs62_.f18 = rhs86_
        d_581_v44_ = rhs87_
        hi3_ = ((self).f24) * (len(d_536_v2_))
        for d_583_i1_ in range(((self).f24 if d_534_v0_ else self.f33), hi3_):
            d_536_v2_ = d_536_v2_
            r1 = ((self).f24) < ((0) - ((default__.fm0((self).f24, 179, globalState)) - (self.f33)))
            if d_534_v0_:
                d_584_v46_: _dafny.Map
                d_584_v46_ = _dafny.Map({p0: (self).f24})
                def iife40_():
                    coll20_ = _dafny.Set()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(280, -841):
                        d_585_v47_: int = compr_20_
                        if ((280) <= (d_585_v47_)) and ((d_585_v47_) < (-841)):
                            coll20_ = coll20_.union(_dafny.Set([default__.safeModuloInt(d_585_v47_, (d_581_v44_).cardinality)]))
                    return _dafny.Set(coll20_)
                rhs88_ = default__.safeModuloInt(len(d_584_v46_), (d_583_i1_) + (self.f33))
                rhs89_ = ((d_540_v6_).cf27) != (iife40_()
                )
                lhs63_ = globalState
                lhs63_.f18 = rhs88_
                r1 = rhs89_
                d_586_v48_: _dafny.Array
                def lambda22_(d_587_v2_, d_588_i1_):
                    def lambda23_(d_589_i2_):
                        return (d_587_v2_).set(default__.safeIndex(d_588_i1_, len(d_587_v2_)), _dafny.CodePoint('v'))

                    return lambda23_

                init13_ = lambda22_(d_536_v2_, d_583_i1_)
                nw91_ = _dafny.Array(None, 16)
                for i0_13_ in range(nw91_.length(0)):
                    nw91_[i0_13_] = init13_(i0_13_)
                d_586_v48_ = nw91_
                d_586_v48_ = d_586_v48_
                d_590_v49_: int
                out22_: int
                out22_ = default__.m0(self.f33, ((self).f24) - (-447), globalState)
                d_590_v49_ = out22_
                d_591_v50_: D6
                d_591_v50_ = D6_DC17(d_534_v0_)
                d_592_v51_: str
                d_592_v51_ = _dafny.CodePoint('d')
                d_593_v52_: _dafny.Map
                d_593_v52_ = _dafny.Map({len(d_536_v2_): default__.fm4(d_534_v0_, d_590_v49_, d_592_v51_, d_534_v0_, globalState)})
                d_594_v53_: _dafny.Map
                d_594_v53_ = _dafny.Map({d_591_v50_: d_593_v52_})
                d_595_v54_: _dafny.Map
                d_595_v54_ = _dafny.Map({d_590_v49_: d_593_v52_})
                d_594_v53_ = (d_594_v53_).set(d_591_v50_, (((d_595_v54_)[d_590_v49_] if (d_590_v49_) in (d_595_v54_) else _dafny.Map({d_590_v49_: d_536_v2_}))) | (d_593_v52_))
                rhs90_ = d_534_v0_
                rhs91_ = (_dafny.SeqWithoutIsStrInference([d_592_v51_ for d_596_i3_ in range(default__.abs(212))])) + (d_536_v2_)
                lhs64_ = globalState
                lhs65_ = globalState
                lhs64_.f2 = rhs90_
                lhs65_.f17 = rhs91_
            elif True:
                (globalState).f21 = d_542_v8_
                d_597_v55_: _dafny.Map
                d_597_v55_ = _dafny.Map({d_534_v0_: len((d_541_v7_).set(default__.safeIndex(d_583_i1_, len(d_541_v7_)), d_583_i1_))})
                d_598_v56_: _dafny.Map
                d_598_v56_ = _dafny.Map({d_534_v0_: len(d_597_v55_)})
                r0 = (((d_598_v56_)[d_534_v0_] if (d_534_v0_) in (d_598_v56_) else 679)) - (default__.safeDivisionInt(d_583_i1_, self.f33))
                d_599_v57_: C0
                nw92_ = C0()
                nw92_.ctor__(d_534_v0_)
                d_599_v57_ = nw92_
                d_600_v58_: _dafny.Map
                d_600_v58_ = _dafny.Map({_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_601_i4_ in range(default__.abs(210))]))}): (d_599_v57_).f34})
                d_600_v58_ = d_600_v58_
                index71_ = default__.safeIndex(151, (p0).length(0))
                (p0)[index71_] = False
                d_602_v59_: str
                d_602_v59_ = _dafny.CodePoint('s')
                d_603_v60_: _dafny.Map
                d_603_v60_ = _dafny.Map({d_543_v9_: d_602_v59_})
                index72_ = default__.safeIndex(151, (p0).length(0))
                rhs92_ = not(((0) - (d_583_i1_)) > ((701 if True else d_583_i1_)))
                rhs93_ = (d_541_v7_)[default__.safeIndex((102) + (97), len(d_541_v7_))]
                rhs94_ = (_dafny.Map({d_537_v3_: d_602_v59_})) | ((d_603_v60_) | (_dafny.Map({d_537_v3_: d_602_v59_})))
                lhs66_ = p0
                lhs67_ = default__.safeIndex(151, (p0).length(0))
                lhs68_ = globalState
                lhs66_[lhs67_] = rhs92_
                lhs68_.f9 = rhs93_
                d_603_v60_ = rhs94_
            r1 = (134) == (d_583_i1_)
        r0 = (self.f33) * ((default__.fm0((self).f24, self.f33, globalState)) * (610))
        r1 = d_534_v0_
        d_604_v61_: str
        d_604_v61_ = _dafny.CodePoint('f')
        d_605_v62_: _dafny.Map
        d_605_v62_ = _dafny.Map({(d_581_v44_).cardinality: False})
        r2 = (d_541_v7_)[default__.safeIndex(default__.fm0((0) - ((D1_DC4(d_534_v0_, d_534_v0_, d_604_v61_, len(d_541_v7_), default__.fm0(self.f33, default__.fm0(len(d_605_v62_), (self).f24, globalState), globalState))).cf12), self.f33, globalState), len(d_541_v7_))]
        r3 = (d_536_v2_)[default__.safeIndex((-930) - (447), len(d_536_v2_))]
        return r0, r1, r2, r3

    @property
    def f32(self):
        return self._f32

class C2(T0):
    def  __init__(self):
        self._f24: int = int(0)
        self._f25: _dafny.Seq = _dafny.Seq({})
        self.f31: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f31, f24, f25):
        (self).f31 = f31
        (self)._f24 = f24
        (self)._f25 = f25

    def fm9(self, p0, p1, p2, globalState):
        return (_dafny.MultiSet([D1_DC6()])).issubset(_dafny.MultiSet([D1_DC6(), D1_DC6()]))

    def m1(self, globalState):
        r0: bool = False
        d_606_v0_: D0
        d_606_v0_ = D0_DC1((self).f24, (self).f24, self.f31)
        def lambda24_(source7_):
            if source7_.is_DC1:
                d_607___mcc_h4_ = source7_.cf1
                d_608___mcc_h5_ = source7_.cf2
                d_609___mcc_h6_ = source7_.cf3
                d_610_cf3_ = d_609___mcc_h6_
                d_611_cf2_ = d_608___mcc_h5_
                d_612_cf1_ = d_607___mcc_h4_
                return d_610_cf3_
            elif source7_.is_DC2:
                d_613___mcc_h7_ = source7_.cf4
                d_614___mcc_h8_ = source7_.cf5
                d_615___mcc_h9_ = source7_.cf6
                d_616_cf6_ = d_615___mcc_h9_
                d_617_cf5_ = d_614___mcc_h8_
                d_618_cf4_ = d_613___mcc_h7_
                return not(self.f31)
            elif True:
                d_619___mcc_h10_ = source7_.cf0
                d_620_cf0_ = d_619___mcc_h10_
                return self.f31

        if lambda24_(d_606_v0_):
            d_621_v1_: _dafny.Array
            def lambda25_(d_622_i0_):
                return (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_623_i1_ in range(default__.abs(-591))]))) != ((self).f24)

            init14_ = lambda25_
            nw93_ = _dafny.Array(None, 14)
            for i0_14_ in range(nw93_.length(0)):
                nw93_[i0_14_] = init14_(i0_14_)
            d_621_v1_ = nw93_
            rhs95_ = d_621_v1_
            rhs96_ = (self.f31) == ((self.f31) and (self.f31))
            lhs69_ = globalState
            d_621_v1_ = rhs95_
            lhs69_.f2 = rhs96_
            if self.f31:
                d_624_v2_: _dafny.Seq
                d_624_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvbaaxfn"))
                index73_ = default__.safeIndex(884, (d_621_v1_).length(0))
                (d_621_v1_)[index73_] = ((self).f24) <= (len(d_624_v2_))
                index74_ = default__.safeIndex(884, (d_621_v1_).length(0))
                (d_621_v1_)[index74_] = self.f31
                d_625_v3_: _dafny.Map
                d_625_v3_ = _dafny.Map({(self).f24: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_626_i2_ in range(default__.abs(69))]))})
                (globalState).f9 = ((self).f24) - (default__.safeDivisionInt(len(d_625_v3_), (self).f24))
                (globalState).f18 = (self).f24
                d_627_v4_: _dafny.MultiSet
                d_627_v4_ = _dafny.MultiSet([(d_621_v1_)[default__.safeIndex(884, (d_621_v1_).length(0))], (d_621_v1_)[default__.safeIndex(884, (d_621_v1_).length(0))], self.f31])
                d_628_v5_: str
                d_628_v5_ = _dafny.CodePoint('o')
                d_629_v6_: _dafny.Array
                nw94_ = _dafny.Array(None, 1)
                nw94_[int(0)] = d_628_v5_
                d_629_v6_ = nw94_
                index75_ = default__.safeIndex(531, (d_629_v6_).length(0))
                (d_629_v6_)[index75_] = (default__.fm3(globalState) if not(self.f31) else _dafny.CodePoint('t'))
                d_630_v7_: D0
                d_630_v7_ = D0_DC2(self.f31, (self).f24, d_628_v5_)
                index76_ = default__.safeIndex(884, (d_621_v1_).length(0))
                index77_ = default__.safeIndex(531, (d_629_v6_).length(0))
                rhs97_ = not((self).fm9((d_621_v1_)[default__.safeIndex(884, (d_621_v1_).length(0))], default__.fm0((self).f24, (self).f24, globalState), (self).f24, globalState))
                rhs98_ = d_627_v4_
                rhs99_ = (d_630_v7_).cf6
                lhs70_ = d_621_v1_
                lhs71_ = default__.safeIndex(884, (d_621_v1_).length(0))
                lhs72_ = d_629_v6_
                lhs73_ = default__.safeIndex(531, (d_629_v6_).length(0))
                lhs70_[lhs71_] = rhs97_
                d_627_v4_ = rhs98_
                lhs72_[lhs73_] = rhs99_
                index78_ = default__.safeIndex(884, (d_621_v1_).length(0))
                (d_621_v1_)[index78_] = ((d_621_v1_)[default__.safeIndex(884, (d_621_v1_).length(0))]) or (self.f31)
            elif True:
                r0 = self.f31
                index79_ = default__.safeIndex(816, (d_621_v1_).length(0))
                (d_621_v1_)[index79_] = (155) <= (len(_dafny.SeqWithoutIsStrInference([(self).f24 for d_631_i3_ in range(default__.abs(-796))])))
                index80_ = default__.safeIndex(816, (d_621_v1_).length(0))
                (d_621_v1_)[index80_] = self.f31
                (globalState).f2 = (self).fm9(self.f31, (self).f24, (self).f24, globalState)
                d_632_v8_: _dafny.Seq
                d_632_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuqmdjtcc"))
                d_633_v9_: str
                d_633_v9_ = _dafny.CodePoint('e')
                d_634_v10_: D0
                d_634_v10_ = D0_DC2((d_621_v1_)[default__.safeIndex(816, (d_621_v1_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qoctswai"))), d_633_v9_)
                d_635_v11_: _dafny.Map
                d_635_v11_ = _dafny.Map({(d_634_v10_).cf6: (self).f24})
                d_636_v12_: C1
                nw95_ = C1()
                nw95_.ctor__(default__.fm29(743, (self).f24, globalState), (self).f24, (self).f24, ((self).f25).set(default__.safeIndex(len(d_632_v8_), len((self).f25)), d_635_v11_))
                d_636_v12_ = nw95_
                nw96_ = C1()
                nw96_.ctor__((d_636_v12_).f32, (self).f24, (self).f24, (self).f25)
                d_636_v12_ = nw96_
            d_637_v13_: _dafny.MultiSet
            d_637_v13_ = _dafny.MultiSet([d_621_v1_, d_621_v1_, d_621_v1_])
            d_638_v14_: _dafny.Array
            def lambda26_(d_639_i4_):
                return default__.safeModuloInt(d_639_i4_, (self).f24)

            init15_ = lambda26_
            nw97_ = _dafny.Array(None, 19)
            for i0_15_ in range(nw97_.length(0)):
                nw97_[i0_15_] = init15_(i0_15_)
            d_638_v14_ = nw97_
            d_640_v15_: _dafny.Map
            d_640_v15_ = _dafny.Map({(d_637_v13_) - (d_637_v13_): d_638_v14_})
            d_640_v15_ = (d_640_v15_).set(d_637_v13_, d_638_v14_)
            d_641_v16_: D1
            d_641_v16_ = D1_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_642_i5_ in range(default__.abs(592))]))
            d_643_v17_: D8
            d_643_v17_ = D8_DC26(d_641_v16_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_644_i6_ in range(default__.abs(623))]))
            d_645_v18_: _dafny.Map
            d_645_v18_ = _dafny.Map({self.f31: len(_dafny.Map({(0) - ((self).f24): d_643_v17_}))})
            d_646_v19_: _dafny.Seq
            d_646_v19_ = _dafny.SeqWithoutIsStrInference([d_645_v18_, default__.fm17((self).f24, (self).f24, (self).f24, self.f31, globalState), d_645_v18_, d_645_v18_, d_645_v18_])
            d_646_v19_ = d_646_v19_
            d_647_v20_: D1
            d_647_v20_ = D1_DC5(True, self.f31, self.f31)
            d_648_v21_: D6
            d_648_v21_ = D6_DC19(D6_DC18(False))
            d_649_v22_: D6
            d_649_v22_ = D6_DC19(d_648_v21_)
            d_650_v23_: D6
            d_650_v23_ = D6_DC19((D7_DC21(d_647_v20_, d_649_v22_, d_621_v1_)).cf33)
            source8_ = d_649_v22_
            if source8_.is_DC17:
                d_651___mcc_h0_ = source8_.cf28
                d_652_cf28_ = d_651___mcc_h0_
                d_653_v24_: _dafny.Seq
                d_653_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibwh"))
                (globalState).f17 = d_653_v24_
                d_654_v25_: _dafny.Map
                d_654_v25_ = _dafny.Map({self.f31: d_652_cf28_})
                d_655_v26_: D6
                d_655_v26_ = D6_DC18(((d_654_v25_)[d_652_cf28_] if (d_652_cf28_) in (d_654_v25_) else d_652_cf28_))
                d_652_cf28_ = (d_655_v26_).cf29
                d_656_v27_: _dafny.Seq
                d_656_v27_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                index81_ = default__.safeIndex(655, (d_638_v14_).length(0))
                (d_638_v14_)[index81_] = default__.safeDivisionInt(len(d_656_v27_), (self).f24)
                index82_ = default__.safeIndex(655, (d_638_v14_).length(0))
                (d_638_v14_)[index82_] = default__.safeDivisionInt((self).f24, (0) - ((0) - (len(d_653_v24_))))
                d_657_v28_: _dafny.Seq
                d_657_v28_ = _dafny.SeqWithoutIsStrInference([d_652_cf28_, self.f31, d_652_cf28_])
                d_658_v29_: _dafny.MultiSet
                d_658_v29_ = _dafny.MultiSet([(d_638_v14_)[default__.safeIndex(655, (d_638_v14_).length(0))]])
                d_659_v30_: _dafny.Seq
                d_659_v30_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f24]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_657_v28_)])), _dafny.MultiSet([len(d_654_v25_)]), (d_658_v29_).set(len(d_656_v27_), default__.abs((d_638_v14_)[default__.safeIndex(655, (d_638_v14_).length(0))]))])
                d_660_v31_: C1
                nw98_ = C1()
                nw98_.ctor__(d_659_v30_, -808, (self).f24, (self).f25)
                d_660_v31_ = nw98_
            elif source8_.is_DC18:
                d_661___mcc_h1_ = source8_.cf29
                d_662_cf29_ = d_661___mcc_h1_
                d_663_v32_: str
                d_663_v32_ = _dafny.CodePoint('d')
                d_664_v33_: _dafny.Seq
                d_664_v33_ = _dafny.SeqWithoutIsStrInference([self.f31, d_662_cf29_])
                d_665_v34_: _dafny.Map
                d_665_v34_ = _dafny.Map({d_663_v32_: (d_664_v33_) + (d_664_v33_)})
                d_666_v35_: _dafny.MultiSet
                d_666_v35_ = _dafny.MultiSet([d_662_cf29_])
                d_667_v36_: _dafny.MultiSet
                d_667_v36_ = d_666_v35_
                rhs100_ = _dafny.Map({_dafny.CodePoint('q'): d_664_v33_})
                rhs101_ = d_666_v35_
                rhs102_ = False
                lhs74_ = globalState
                d_665_v34_ = rhs100_
                d_667_v36_ = rhs101_
                lhs74_.f2 = rhs102_
                d_638_v14_ = d_638_v14_
                d_638_v14_ = d_638_v14_
                d_663_v32_ = d_663_v32_
            elif source8_.is_DC16:
                d_668___mcc_h2_ = source8_.cf27
                d_669_cf27_ = d_668___mcc_h2_
                d_670_v37_: _dafny.Set
                d_670_v37_ = _dafny.Set({self.f31})
                d_671_v38_: _dafny.Map
                d_671_v38_ = _dafny.Map({default__.safeDivisionInt((self).f24, (self).f24): d_670_v37_})
                d_671_v38_ = (d_671_v38_).set(37, d_670_v37_)
                d_672_v39_: _dafny.Seq
                d_672_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
                d_673_v40_: C0
                nw99_ = C0()
                nw99_.ctor__((len(d_672_v39_)) > ((0) - ((self).f24)))
                d_673_v40_ = nw99_
                d_674_v41_: _dafny.Seq
                d_674_v41_ = _dafny.SeqWithoutIsStrInference([(d_673_v40_).f34])
                d_675_v42_: str
                d_675_v42_ = _dafny.CodePoint('g')
                d_676_v43_: D2
                d_676_v43_ = D2_DC8()
                d_677_v44_: _dafny.MultiSet
                d_677_v44_ = _dafny.MultiSet([d_676_v43_])
                d_678_v45_: _dafny.MultiSet
                d_678_v45_ = _dafny.MultiSet([(self).f24, len(_dafny.Map({(d_673_v40_).f34: d_675_v42_})), (0) - ((d_677_v44_).cardinality)])
                (globalState).f21 = (((default__.fm23(globalState)) + (d_674_v41_) if ((d_673_v40_).f34) and (not((d_673_v40_).f34)) else (d_674_v41_ if (d_673_v40_).f34 else d_674_v41_))).set(default__.safeIndex(((self).f24) * (((d_678_v45_)[(self).f24] if ((self).f24) in (d_678_v45_) else (self).f24)), len(((default__.fm23(globalState)) + (d_674_v41_) if ((d_673_v40_).f34) and (not((d_673_v40_).f34)) else (d_674_v41_ if (d_673_v40_).f34 else d_674_v41_)))), (d_672_v39_) <= (d_672_v39_))
                d_672_v39_ = _dafny.SeqWithoutIsStrInference([d_675_v42_ for d_679_i7_ in range(default__.abs(663))])
            elif True:
                d_680___mcc_h3_ = source8_.cf30
                d_681_cf30_ = d_680___mcc_h3_
                (globalState).f2 = self.f31
                (self).f31 = self.f31
                d_682_v46_: _dafny.Seq
                d_682_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ea"))
                rhs103_ = d_682_v46_
                rhs104_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_683_i8_ in range(default__.abs(180))]) if self.f31 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hdlwah")))
                rhs105_ = d_682_v46_
                rhs106_ = (-216) == ((self).f24)
                lhs75_ = globalState
                lhs76_ = globalState
                lhs77_ = globalState
                lhs75_.f17 = rhs103_
                lhs76_.f17 = rhs104_
                lhs77_.f17 = rhs105_
                r0 = rhs106_
                d_684_v47_: D8
                d_684_v47_ = D8_DC24((self).f24)
                d_685_v48_: str
                d_685_v48_ = _dafny.CodePoint('f')
                d_686_v49_: _dafny.Map
                d_686_v49_ = _dafny.Map({d_685_v48_: d_685_v48_})
                d_687_v51_: _dafny.MultiSet
                def iife41_():
                    coll21_ = _dafny.Set()
                    compr_21_: str
                    for compr_21_ in (d_686_v49_).keys.Elements:
                        d_688_v50_: str = compr_21_
                        if (d_688_v50_) in (d_686_v49_):
                            coll21_ = coll21_.union(_dafny.Set([d_688_v50_]))
                    return _dafny.Set(coll21_)
                d_687_v51_ = _dafny.MultiSet([len(iife41_()
                ), (self).f24])
                d_689_v52_: _dafny.Set
                d_689_v52_ = _dafny.Set({self.f31, True, self.f31})
                d_684_v47_ = D8_DC24(((d_687_v51_)[(self).f24] if ((self).f24) in (d_687_v51_) else len(d_689_v52_)))
        elif True:
            d_690_v53_: _dafny.Map
            d_690_v53_ = _dafny.Map({((self).f24) - ((self).f24): (self).f24})
            (globalState).f22 = d_690_v53_
            d_691_v54_: _dafny.Seq
            d_691_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
            d_692_v55_: _dafny.Seq
            d_692_v55_ = _dafny.SeqWithoutIsStrInference([(self).f24, len(d_691_v54_), 920])
            if (d_692_v55_) < (_dafny.SeqWithoutIsStrInference([(self).f24, (self).f24, (0) - ((self).f24), (self).f24, (self).f24])):
                d_693_v56_: _dafny.Map
                d_693_v56_ = _dafny.Map({self.f31: 170})
                d_694_v57_: D4
                d_694_v57_ = D4_DC13(((self).f24) - (511), (d_693_v56_) | (d_693_v56_), d_691_v54_)
                d_694_v57_ = d_694_v57_
                (globalState).f18 = (0) - ((0) - ((self).f24))
                d_690_v53_ = (_dafny.Map({(self).f24: (self).f24}) if True else d_690_v53_)
                d_695_v58_: _dafny.MultiSet
                d_695_v58_ = _dafny.MultiSet([True, self.f31])
                d_696_v59_: str
                d_696_v59_ = _dafny.CodePoint('t')
                d_697_v60_: _dafny.Map
                d_697_v60_ = _dafny.Map({(self).f24: d_691_v54_})
                d_698_v61_: C0
                nw100_ = C0()
                nw100_.ctor__(self.f31)
                d_698_v61_ = nw100_
                d_699_v62_: _dafny.Seq
                d_699_v62_ = _dafny.SeqWithoutIsStrInference([d_698_v61_, d_698_v61_])
                d_700_v63_: _dafny.Set
                d_700_v63_ = _dafny.Set({(self).f24, len(d_699_v62_)})
                d_701_v64_: _dafny.Array
                nw101_ = _dafny.Array(None, 15)
                nw101_[int(0)] = self.f31
                nw101_[int(1)] = not(False)
                nw101_[int(2)] = (D6_DC17(self.f31)).cf28
                nw101_[int(3)] = self.f31
                nw101_[int(4)] = (self.f31) == (self.f31)
                nw101_[int(5)] = (d_695_v58_).isdisjoint(d_695_v58_)
                nw101_[int(6)] = self.f31
                nw101_[int(7)] = default__.fm2(211, len(d_692_v55_), d_696_v59_, globalState)
                nw101_[int(8)] = ((self).f24) <= ((self).f24)
                nw101_[int(9)] = ((self).f24) not in ((d_697_v60_).set((self).f24, d_691_v54_))
                nw101_[int(10)] = (len(d_700_v63_)) not in (default__.fm30(d_696_v59_, not(self.f31), globalState))
                nw101_[int(11)] = (755) != ((self).f24)
                nw101_[int(12)] = True
                nw101_[int(13)] = self.f31
                nw101_[int(14)] = (d_698_v61_).f34
                d_701_v64_ = nw101_
                d_701_v64_ = d_701_v64_
                d_702_v65_: _dafny.MultiSet
                d_702_v65_ = _dafny.MultiSet([(self).f24, (self).f24, len(_dafny.SeqWithoutIsStrInference([(self).f24 for d_703_i9_ in range(default__.abs(684))])), (self).f24, default__.fm0((self).f24, (self).f24, globalState)])
                d_704_v66_: _dafny.Set
                d_704_v66_ = _dafny.Set({d_702_v65_})
                d_704_v66_ = d_704_v66_
            elif True:
                rhs107_ = False
                rhs108_ = self.f31
                rhs109_ = 613
                lhs78_ = globalState
                lhs79_ = globalState
                lhs80_ = globalState
                lhs78_.f2 = rhs107_
                lhs79_.f2 = rhs108_
                lhs80_.f1 = rhs109_
                d_705_v67_: _dafny.Array
                def lambda27_(d_706_v55_):
                    def lambda28_(d_707_i10_):
                        return (d_706_v55_) == (_dafny.SeqWithoutIsStrInference([(self).f24]))

                    return lambda28_

                init16_ = lambda27_(d_692_v55_)
                nw102_ = _dafny.Array(None, 15)
                for i0_16_ in range(nw102_.length(0)):
                    nw102_[i0_16_] = init16_(i0_16_)
                d_705_v67_ = nw102_
                d_705_v67_ = d_705_v67_
                rhs110_ = d_691_v54_
                rhs111_ = (((self).f24) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljamq"))))) - (default__.safeDivisionInt((self).f24, (self).f24))
                lhs81_ = globalState
                lhs82_ = globalState
                lhs81_.f17 = rhs110_
                lhs82_.f18 = rhs111_
                d_708_v68_: C0
                nw103_ = C0()
                nw103_.ctor__(self.f31)
                d_708_v68_ = nw103_
                (self).f31 = not (True) or ((d_708_v68_).f34)
            if self.f31:
                d_709_v69_: _dafny.Seq
                d_709_v69_ = _dafny.SeqWithoutIsStrInference([d_691_v54_])
                (self).f31 = (d_691_v54_) != (((d_709_v69_)[default__.safeIndex((self).f24, len(d_709_v69_))]) + (d_691_v54_))
                d_710_v70_: D4
                d_710_v70_ = D4_DC12(d_692_v55_)
                d_692_v55_ = (d_710_v70_).cf21
                d_711_v71_: _dafny.Map
                d_711_v71_ = _dafny.Map({self.f31: (340) - ((self).f24)})
                d_711_v71_ = (d_711_v71_).set(False, (self).f24)
                (globalState).f14 = (d_692_v55_)[default__.safeIndex((self).f24, len(d_692_v55_))]
                (globalState).f2 = not((d_692_v55_) == (_dafny.SeqWithoutIsStrInference([(self).f24, (self).f24, (self).f24])))
            elif True:
                d_712_v72_: _dafny.Set
                d_712_v72_ = _dafny.Set({(self).f24})
                d_713_v73_: _dafny.Set
                d_713_v73_ = _dafny.Set({d_712_v72_})
                (globalState).f18 = len(d_713_v73_)
                d_714_v74_: _dafny.MultiSet
                d_714_v74_ = _dafny.MultiSet([(self).f24])
                d_715_v75_: T0
                nw104_ = C1()
                nw104_.ctor__((_dafny.SeqWithoutIsStrInference([d_714_v74_])).set(default__.safeIndex(len(d_692_v55_), len(_dafny.SeqWithoutIsStrInference([d_714_v74_]))), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_716_i11_ in range(default__.abs(973))]))])), (self).f24, (self).f24, (self).f25)
                d_715_v75_ = nw104_
                d_717_v76_: _dafny.Map
                d_717_v76_ = _dafny.Map({False: d_715_v75_})
                d_718_v77_: T0
                d_718_v77_ = d_715_v75_
                d_719_v78_: _dafny.Array
                nw105_ = _dafny.Array(None, 19)
                nw105_[int(0)] = d_715_v75_
                nw105_[int(1)] = d_715_v75_
                nw105_[int(2)] = d_715_v75_
                nw105_[int(3)] = d_715_v75_
                nw105_[int(4)] = ((d_717_v76_)[self.f31] if (self.f31) in (d_717_v76_) else d_715_v75_)
                nw105_[int(5)] = d_715_v75_
                nw105_[int(6)] = d_715_v75_
                nw105_[int(7)] = d_715_v75_
                nw105_[int(8)] = d_715_v75_
                nw105_[int(9)] = d_715_v75_
                nw105_[int(10)] = d_715_v75_
                nw105_[int(11)] = d_715_v75_
                nw105_[int(12)] = d_715_v75_
                nw105_[int(13)] = d_715_v75_
                nw105_[int(14)] = d_715_v75_
                nw105_[int(15)] = d_715_v75_
                nw105_[int(16)] = d_715_v75_
                nw105_[int(17)] = d_715_v75_
                nw105_[int(18)] = (d_718_v77_)
                d_719_v78_ = nw105_
                d_720_v79_: _dafny.Seq
                d_720_v79_ = _dafny.SeqWithoutIsStrInference([d_715_v75_])
                d_721_v80_: _dafny.Map
                d_721_v80_ = _dafny.Map({self.f31: self.f31})
                index83_ = default__.safeIndex(674, (d_719_v78_).length(0))
                (d_719_v78_)[index83_] = ((d_720_v79_)[default__.safeIndex(len(d_721_v80_), len(d_720_v79_))] if self.f31 else d_715_v75_)
                d_722_v81_: _dafny.Map
                d_722_v81_ = _dafny.Map({(0) - ((0) - (-79)): d_715_v75_})
                index84_ = default__.safeIndex(674, (d_719_v78_).length(0))
                (d_719_v78_)[index84_] = (((d_722_v81_)[(self).f24] if ((self).f24) in (d_722_v81_) else d_715_v75_) if self.f31 else (d_715_v75_ if not(self.f31) else d_715_v75_))
                (globalState).f2 = not(not(not(self.f31)))
                (globalState).f9 = (d_715_v75_).f24
                d_723_v82_: _dafny.Set
                d_723_v82_ = _dafny.Set({self.f31})
                d_724_v83_: _dafny.Map
                d_724_v83_ = _dafny.Map({(d_715_v75_).f24: self.f31})
                d_725_v84_: str
                d_725_v84_ = _dafny.CodePoint('x')
                d_726_v85_: _dafny.Set
                d_726_v85_ = _dafny.Set({d_725_v84_, d_725_v84_})
                d_727_v86_: _dafny.Array
                nw106_ = _dafny.Array(None, 12)
                nw106_[int(0)] = d_723_v82_
                nw106_[int(1)] = d_723_v82_
                nw106_[int(2)] = d_723_v82_
                nw106_[int(3)] = (default__.fm28(self.f31, (self).f24, self.f31, (0) - ((d_715_v75_).f24), globalState)) | (default__.fm28(((d_724_v83_)[(0) - ((0) - ((self).f24))] if ((0) - ((0) - ((self).f24))) in (d_724_v83_) else self.f31), len(d_726_v85_), self.f31, (self).f24, globalState))
                nw106_[int(4)] = (d_723_v82_) - (d_723_v82_)
                nw106_[int(5)] = d_723_v82_
                nw106_[int(6)] = d_723_v82_
                nw106_[int(7)] = d_723_v82_
                nw106_[int(8)] = default__.fm28(self.f31, (self).f24, self.f31, 503, globalState)
                nw106_[int(9)] = _dafny.Set({self.f31, True, self.f31})
                nw106_[int(10)] = (d_723_v82_) - (d_723_v82_)
                nw106_[int(11)] = d_723_v82_
                d_727_v86_ = nw106_
                index85_ = default__.safeIndex(683, (d_727_v86_).length(0))
                (d_727_v86_)[index85_] = d_723_v82_
                index86_ = default__.safeIndex(683, (d_727_v86_).length(0))
                (d_727_v86_)[index86_] = d_723_v82_
            (globalState).f2 = self.f31
            (globalState).f17 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlgqytw"))
        d_728_v87_: _dafny.Array
        def lambda29_(d_729_i12_):
            return _dafny.CodePoint('r')

        init17_ = lambda29_
        nw107_ = _dafny.Array(None, 26)
        for i0_17_ in range(nw107_.length(0)):
            nw107_[i0_17_] = init17_(i0_17_)
        d_728_v87_ = nw107_
        d_730_v88_: str
        d_730_v88_ = _dafny.CodePoint('y')
        index87_ = default__.safeIndex(812, (d_728_v87_).length(0))
        (d_728_v87_)[index87_] = d_730_v88_
        d_731_v89_: _dafny.Seq
        d_731_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pe"))
        index88_ = default__.safeIndex(812, (d_728_v87_).length(0))
        rhs112_ = (self).f24
        rhs113_ = (self).f24
        rhs114_ = d_730_v88_
        rhs115_ = d_731_v89_
        rhs116_ = default__.safeDivisionInt(default__.safeDivisionInt((self).f24, (self).f24), default__.safeModuloInt((self).f24, len(d_731_v89_)))
        lhs83_ = globalState
        lhs84_ = globalState
        lhs85_ = d_728_v87_
        lhs86_ = default__.safeIndex(812, (d_728_v87_).length(0))
        lhs87_ = globalState
        lhs88_ = globalState
        lhs83_.f9 = rhs112_
        lhs84_.f14 = rhs113_
        lhs85_[lhs86_] = rhs114_
        lhs87_.f17 = rhs115_
        lhs88_.f9 = rhs116_
        hi4_ = (self).f24
        for d_732_i13_ in range((self).f24, hi4_):
            d_733_v90_: _dafny.Seq
            d_733_v90_ = _dafny.SeqWithoutIsStrInference([self.f31, self.f31, self.f31, self.f31])
            d_734_v91_: _dafny.Map
            d_734_v91_ = _dafny.Map({self.f31: (self).fm9(default__.fm2(d_732_i13_, 968, (d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))], globalState), len(d_733_v90_), d_732_i13_, globalState)})
            d_734_v91_ = (d_734_v91_).set(self.f31, self.f31)
            if (default__.safeDivisionInt((self).f24, d_732_i13_)) == (len(default__.fm18(globalState))):
                d_735_v92_: _dafny.Array
                nw108_ = _dafny.Array(int(0), 1)
                d_735_v92_ = nw108_
                index89_ = default__.safeIndex(285, (d_735_v92_).length(0))
                (d_735_v92_)[index89_] = (0) - (d_732_i13_)
                index90_ = default__.safeIndex(333, (d_735_v92_).length(0))
                (d_735_v92_)[index90_] = (self).f24
                index91_ = default__.safeIndex(285, (d_735_v92_).length(0))
                index92_ = default__.safeIndex(333, (d_735_v92_).length(0))
                rhs117_ = (self).f24
                rhs118_ = default__.safeModuloInt((self).f24, 225)
                lhs89_ = d_735_v92_
                lhs90_ = default__.safeIndex(285, (d_735_v92_).length(0))
                lhs91_ = d_735_v92_
                lhs92_ = default__.safeIndex(333, (d_735_v92_).length(0))
                lhs89_[lhs90_] = rhs117_
                lhs91_[lhs92_] = rhs118_
                (globalState).f2 = (self).fm9(self.f31, d_732_i13_, (len(default__.fm4(self.f31, d_732_i13_, _dafny.CodePoint('g'), self.f31, globalState))) * ((self).f24), globalState)
                d_736_v93_: _dafny.Array
                nw109_ = _dafny.Array(False, 26)
                d_736_v93_ = nw109_
                d_737_v94_: _dafny.Array
                nw110_ = _dafny.Array(D6.default()(), 13)
                d_737_v94_ = nw110_
                d_738_v95_: _dafny.Set
                d_738_v95_ = _dafny.Set({(d_735_v92_)[default__.safeIndex(285, (d_735_v92_).length(0))], d_732_i13_, len(d_731_v89_)})
                index93_ = default__.safeIndex(807, (d_737_v94_).length(0))
                (d_737_v94_)[index93_] = D6_DC16(d_738_v95_)
                d_739_v96_: D1
                d_739_v96_ = D1_DC5(self.f31, self.f31, self.f31)
                d_740_v97_: D7
                d_740_v97_ = D7_DC21(d_739_v96_, default__.fm31(self.f31, self.f31, globalState), d_736_v93_)
                index94_ = default__.safeIndex(807, (d_737_v94_).length(0))
                rhs119_ = (d_736_v93_ if self.f31 else (d_740_v97_).cf34)
                rhs120_ = D6_DC16(d_738_v95_)
                lhs93_ = d_737_v94_
                lhs94_ = default__.safeIndex(807, (d_737_v94_).length(0))
                d_736_v93_ = rhs119_
                lhs93_[lhs94_] = rhs120_
                d_741_v98_: _dafny.Array
                nw111_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_741_v98_ = nw111_
                index95_ = default__.safeIndex(286, (d_741_v98_).length(0))
                (d_741_v98_)[index95_] = (d_735_v92_ if self.f31 else d_735_v92_)
                d_742_v99_: D0
                d_742_v99_ = D0_DC0(d_735_v92_)
                index96_ = default__.safeIndex(286, (d_741_v98_).length(0))
                (d_741_v98_)[index96_] = (d_742_v99_).cf0
                (globalState).f9 = (d_732_i13_) * (((self).f24) - ((d_735_v92_)[default__.safeIndex(285, (d_735_v92_).length(0))]))
            elif True:
                d_743_v100_: _dafny.MultiSet
                d_743_v100_ = _dafny.MultiSet([d_732_i13_, d_732_i13_])
                d_744_v101_: _dafny.Seq
                d_744_v101_ = _dafny.SeqWithoutIsStrInference([d_743_v100_])
                d_745_v102_: _dafny.MultiSet
                d_745_v102_ = _dafny.MultiSet([self.f31, self.f31])
                d_746_v103_: _dafny.Set
                d_746_v103_ = _dafny.Set({not(self.f31)})
                d_747_v104_: _dafny.Map
                d_747_v104_ = _dafny.Map({d_730_v88_: (self).f24})
                d_748_v105_: C1
                nw112_ = C1()
                nw112_.ctor__((d_744_v101_) + (d_744_v101_), d_732_i13_, default__.safeModuloInt(d_732_i13_, ((d_745_v102_)[self.f31] if (self.f31) in (d_745_v102_) else len(d_746_v103_))), (_dafny.SeqWithoutIsStrInference([d_747_v104_, d_747_v104_, d_747_v104_]) if True else _dafny.SeqWithoutIsStrInference([d_747_v104_])))
                d_748_v105_ = nw112_
                (globalState).f2 = self.f31
                d_745_v102_ = d_745_v102_
                d_749_v106_: _dafny.Set
                d_749_v106_ = _dafny.Set({d_732_i13_})
                (globalState).f1 = default__.safeDivisionInt(d_748_v105_.f33, default__.safeModuloInt((0) - (d_732_i13_), len(d_749_v106_)))
                (globalState).f18 = (self).f24
            if ((self).f24) != (d_732_i13_):
                d_750_v107_: C0
                nw113_ = C0()
                nw113_.ctor__(self.f31)
                d_750_v107_ = nw113_
                d_751_v108_: _dafny.Map
                d_751_v108_ = _dafny.Map({self.f31: d_732_i13_})
                (globalState).f17 = (d_731_v89_ if default__.fm2(len(d_751_v108_), 392, default__.fm3(globalState), globalState) else d_731_v89_)
                (globalState).f18 = d_732_i13_
                d_752_v109_: _dafny.Array
                def lambda30_(d_753_v90_, d_754_v107_):
                    def lambda31_(d_755_i14_):
                        return (_dafny.SeqWithoutIsStrInference([d_753_v90_])) + (_dafny.SeqWithoutIsStrInference([d_753_v90_, _dafny.SeqWithoutIsStrInference([(d_754_v107_).f34, (d_754_v107_).f34, False])]))

                    return lambda31_

                init18_ = lambda30_(d_733_v90_, d_750_v107_)
                nw114_ = _dafny.Array(None, 5)
                for i0_18_ in range(nw114_.length(0)):
                    nw114_[i0_18_] = init18_(i0_18_)
                d_752_v109_ = nw114_
                d_756_v110_: _dafny.Seq
                d_756_v110_ = _dafny.SeqWithoutIsStrInference([d_733_v90_])
                index97_ = default__.safeIndex(737, (d_752_v109_).length(0))
                (d_752_v109_)[index97_] = (d_756_v110_) + (d_756_v110_)
                d_757_v111_: _dafny.Map
                d_757_v111_ = _dafny.Map({(0) - ((self).f24): d_733_v90_})
                index98_ = default__.safeIndex(737, (d_752_v109_).length(0))
                (d_752_v109_)[index98_] = (((d_756_v110_) + (d_756_v110_) if (d_750_v107_).f34 else (d_756_v110_ if not((d_750_v107_).f34) else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_750_v107_).f34]), ((d_757_v111_)[370] if (370) in (d_757_v111_) else d_733_v90_), d_733_v90_, d_733_v90_])))).set(default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))), len(((d_756_v110_) + (d_756_v110_) if (d_750_v107_).f34 else (d_756_v110_ if not((d_750_v107_).f34) else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_750_v107_).f34]), ((d_757_v111_)[370] if (370) in (d_757_v111_) else d_733_v90_), d_733_v90_, d_733_v90_]))))), d_733_v90_)
                d_733_v90_ = d_733_v90_
            elif True:
                d_758_v112_: _dafny.MultiSet
                d_758_v112_ = _dafny.MultiSet([(d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))], d_730_v88_])
                d_759_v113_: _dafny.Seq
                d_759_v113_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24])
                (globalState).f14 = default__.fm0(default__.fm0(233, default__.fm0((self).f24, ((d_758_v112_)[d_730_v88_] if (d_730_v88_) in (d_758_v112_) else len(d_759_v113_)), globalState), globalState), (self).f24, globalState)
                d_760_v114_: _dafny.Map
                d_760_v114_ = _dafny.Map({d_731_v89_: _dafny.SeqWithoutIsStrInference([self.f31, self.f31])})
                d_761_v115_: _dafny.Set
                d_761_v115_ = _dafny.Set({((d_760_v114_)[d_731_v89_] if (d_731_v89_) in (d_760_v114_) else d_733_v90_), d_733_v90_, default__.fm23(globalState)})
                d_762_v116_: _dafny.Seq
                d_762_v116_ = _dafny.SeqWithoutIsStrInference([d_733_v90_, d_733_v90_])
                def iife42_():
                    coll22_ = _dafny.Set()
                    compr_22_: _dafny.Seq
                    for compr_22_ in (d_762_v116_).Elements:
                        d_763_v117_: _dafny.Seq = compr_22_
                        if (d_763_v117_) in (d_762_v116_):
                            coll22_ = coll22_.union(_dafny.Set([d_763_v117_]))
                    return _dafny.Set(coll22_)
                d_761_v115_ = iife42_()
                
                d_764_v118_: _dafny.Map
                d_764_v118_ = _dafny.Map({(self).f24: len(d_761_v115_)})
                d_765_v119_: int
                out23_: int
                out23_ = default__.m0(((d_764_v118_)[(self).f24] if ((self).f24) in (d_764_v118_) else -202), (0) - (len(d_734_v91_)), globalState)
                d_765_v119_ = out23_
                (globalState).f1 = (d_732_i13_ if self.f31 else default__.safeDivisionInt(d_765_v119_, 958))
                (globalState).f14 = d_765_v119_
            d_766_v120_: _dafny.MultiSet
            d_766_v120_ = _dafny.MultiSet([self.f31])
            d_767_v121_: D6
            d_767_v121_ = D6_DC18(False)
            d_768_v122_: _dafny.Map
            d_768_v122_ = _dafny.Map({self.f31: 513})
            d_769_v123_: _dafny.Array
            nw115_ = _dafny.Array(None, 11)
            nw115_[int(0)] = (d_767_v121_).cf29
            nw115_[int(1)] = self.f31
            nw115_[int(2)] = (self).fm9(self.f31, (self).f24, (self).f24, globalState)
            nw115_[int(3)] = (self).fm9(self.f31, (self).f24, len(d_768_v122_), globalState)
            nw115_[int(4)] = self.f31
            nw115_[int(5)] = self.f31
            nw115_[int(6)] = self.f31
            nw115_[int(7)] = self.f31
            nw115_[int(8)] = self.f31
            nw115_[int(9)] = self.f31
            nw115_[int(10)] = self.f31
            d_769_v123_ = nw115_
            d_770_v124_: _dafny.Map
            d_770_v124_ = _dafny.Map({self.f31: d_769_v123_})
            d_771_v125_: _dafny.Map
            d_771_v125_ = _dafny.Map({(d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))]: ((d_766_v120_).set(not(self.f31), default__.abs(len(d_770_v124_)))) | (d_766_v120_)})
            d_771_v125_ = (d_771_v125_).set(d_730_v88_, d_766_v120_)
        d_772_v126_: D6
        d_772_v126_ = D6_DC16(default__.fm18(globalState))
        source9_ = D6_DC19(d_772_v126_)
        if source9_.is_DC17:
            d_773___mcc_h11_ = source9_.cf28
            d_774_cf28_ = d_773___mcc_h11_
            d_775_v127_: _dafny.Array
            def lambda32_(d_776_i15_):
                return (d_776_i15_) + ((self).f24)

            init19_ = lambda32_
            nw116_ = _dafny.Array(None, 22)
            for i0_19_ in range(nw116_.length(0)):
                nw116_[i0_19_] = init19_(i0_19_)
            d_775_v127_ = nw116_
            index99_ = default__.safeIndex(409, (d_775_v127_).length(0))
            (d_775_v127_)[index99_] = (self).f24
            d_777_v128_: _dafny.Set
            d_777_v128_ = _dafny.Set({(self).f24})
            d_778_v129_: _dafny.MultiSet
            d_778_v129_ = _dafny.MultiSet([d_774_cf28_, self.f31])
            d_779_v131_: _dafny.MultiSet
            def iife43_():
                coll23_ = _dafny.Set()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(103, 967):
                    d_780_v130_: int = compr_23_
                    if ((103) <= (d_780_v130_)) and ((d_780_v130_) < (967)):
                        coll23_ = coll23_.union(_dafny.Set([(d_780_v130_) + (len(d_731_v89_))]))
                return _dafny.Set(coll23_)
            d_779_v131_ = _dafny.MultiSet([d_777_v128_, _dafny.Set({len(_dafny.SeqWithoutIsStrInference([51])), (self).f24, (d_778_v129_).cardinality}), iife43_()
            ])
            index100_ = default__.safeIndex(409, (d_775_v127_).length(0))
            (d_775_v127_)[index100_] = (((d_779_v131_)[d_777_v128_] if (d_777_v128_) in (d_779_v131_) else (self).f24) if self.f31 else len(_dafny.Map({(self).f24: self.f31})))
            if (self).fm9(True, 726, (self).f24, globalState):
                index101_ = default__.safeIndex(812, (d_728_v87_).length(0))
                (d_728_v87_)[index101_] = (d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))]
                d_781_v132_: _dafny.MultiSet
                d_781_v132_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))] for d_782_i16_ in range(default__.abs(866))]), (_dafny.SeqWithoutIsStrInference([d_730_v88_ for d_783_i17_ in range(default__.abs(955))]) if self.f31 else (d_731_v89_).set(default__.safeIndex((d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))], len(d_731_v89_)), _dafny.CodePoint('o'))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqys"))])
                index102_ = default__.safeIndex(409, (d_775_v127_).length(0))
                rhs121_ = d_731_v89_
                rhs122_ = (0) - (default__.safeModuloInt((d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))], (self).f24))
                rhs123_ = d_781_v132_
                lhs95_ = d_775_v127_
                lhs96_ = default__.safeIndex(409, (d_775_v127_).length(0))
                d_731_v89_ = rhs121_
                lhs95_[lhs96_] = rhs122_
                d_781_v132_ = rhs123_
                d_784_v133_: _dafny.Seq
                d_784_v133_ = _dafny.SeqWithoutIsStrInference([d_774_cf28_])
                rhs124_ = (d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))]
                rhs125_ = (d_784_v133_) + (d_784_v133_)
                lhs97_ = globalState
                lhs98_ = globalState
                lhs97_.f1 = rhs124_
                lhs98_.f20 = rhs125_
                (self).f31 = ((d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))]) <= ((self).f24)
                d_785_v134_: _dafny.Seq
                d_785_v134_ = _dafny.SeqWithoutIsStrInference([d_775_v127_, d_775_v127_])
                d_786_v135_: D1
                d_786_v135_ = D1_DC4(d_774_cf28_, self.f31, d_730_v88_, (0) - ((d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))]), (self).f24)
                d_787_v136_: _dafny.Array
                nw117_ = _dafny.Array(None, 17)
                nw117_[int(0)] = d_775_v127_
                nw117_[int(1)] = d_775_v127_
                nw117_[int(2)] = d_775_v127_
                nw117_[int(3)] = d_775_v127_
                nw117_[int(4)] = d_775_v127_
                nw117_[int(5)] = d_775_v127_
                nw117_[int(6)] = d_775_v127_
                nw117_[int(7)] = d_775_v127_
                nw117_[int(8)] = d_775_v127_
                nw117_[int(9)] = d_775_v127_
                nw117_[int(10)] = d_775_v127_
                nw117_[int(11)] = d_775_v127_
                nw117_[int(12)] = d_775_v127_
                nw117_[int(13)] = d_775_v127_
                nw117_[int(14)] = (d_785_v134_)[default__.safeIndex((d_786_v135_).cf12, len(d_785_v134_))]
                nw117_[int(15)] = (d_785_v134_)[default__.safeIndex((self).f24, len(d_785_v134_))]
                nw117_[int(16)] = d_775_v127_
                d_787_v136_ = nw117_
                index103_ = default__.safeIndex(163, (d_787_v136_).length(0))
                (d_787_v136_)[index103_] = d_775_v127_
                index104_ = default__.safeIndex(163, (d_787_v136_).length(0))
                (d_787_v136_)[index104_] = d_775_v127_
            elif True:
                d_788_v137_: _dafny.MultiSet
                d_788_v137_ = _dafny.MultiSet([(self).f24, (self).f24])
                d_789_v138_: _dafny.Map
                d_789_v138_ = _dafny.Map({(d_774_cf28_ if d_774_cf28_ else self.f31): (d_788_v137_) - (d_788_v137_)})
                (globalState).f9 = len(d_789_v138_)
                d_790_v139_: D6
                d_790_v139_ = D6_DC16(_dafny.Set({len(d_731_v89_)}))
                d_791_v140_: _dafny.Map
                d_791_v140_ = _dafny.Map({self.f31: d_790_v139_})
                d_792_v141_: _dafny.Map
                d_792_v141_ = _dafny.Map({d_791_v140_: d_774_cf28_})
                d_792_v141_ = ((d_792_v141_) | (d_792_v141_)) | (d_792_v141_)
                d_793_v142_: _dafny.Map
                d_793_v142_ = _dafny.Map({(d_788_v137_).set((self).f24, default__.abs((self).f24)): d_730_v88_})
                d_794_v144_: _dafny.MultiSet
                d_794_v144_ = _dafny.MultiSet([(d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))]])
                def iife44_():
                    coll24_ = _dafny.Map()
                    compr_24_: str
                    for compr_24_ in (d_794_v144_).Elements:
                        d_795_v143_: str = compr_24_
                        if (d_795_v143_) in (d_794_v144_):
                            coll24_[d_795_v143_] = _dafny.Map({self.f31: (d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))]})
                    return _dafny.Map(coll24_)
                d_793_v142_ = (d_793_v142_).set(_dafny.MultiSet([len(iife44_()
                )]), (d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))])
                d_796_v145_: C0
                nw118_ = C0()
                nw118_.ctor__(self.f31)
                d_796_v145_ = nw118_
                d_797_v146_: C1
                nw119_ = C1()
                nw119_.ctor__(_dafny.SeqWithoutIsStrInference([d_788_v137_ for d_798_i18_ in range(default__.abs(286))]), (_dafny.MultiSet([(d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))]])).cardinality, (self).f24, _dafny.SeqWithoutIsStrInference([_dafny.Map({d_730_v88_: -786}) for d_799_i19_ in range(default__.abs(367))]))
                d_797_v146_ = nw119_
            (globalState).f14 = default__.safeModuloInt(-45, ((d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))]) + (519))
            d_800_v147_: _dafny.Seq
            d_800_v147_ = _dafny.SeqWithoutIsStrInference([False, d_774_cf28_, d_774_cf28_, not(d_774_cf28_), d_774_cf28_])
            if ((d_774_cf28_) in (d_800_v147_)) in (d_800_v147_):
                index105_ = default__.safeIndex(812, (d_728_v87_).length(0))
                (d_728_v87_)[index105_] = d_730_v88_
                d_801_v148_: _dafny.Map
                d_801_v148_ = _dafny.Map({d_730_v88_: self.f31})
                d_801_v148_ = (d_801_v148_).set((d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))], (d_800_v147_) != (d_800_v147_))
                r0 = not((self.f31) == ((True if self.f31 else d_774_cf28_)))
                (globalState).f2 = (d_731_v89_) != ((d_731_v89_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxrlmgp"))))
                (globalState).f1 = len(d_731_v89_)
            elif True:
                d_802_v149_: _dafny.Map
                d_802_v149_ = _dafny.Map({self.f31: (d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))]})
                d_803_v150_: _dafny.Set
                d_803_v150_ = _dafny.Set({((d_802_v149_)[d_774_cf28_] if (d_774_cf28_) in (d_802_v149_) else d_730_v88_)})
                d_804_v151_: D3
                d_804_v151_ = D3_DC11((d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))], d_774_cf28_, default__.fm2(len(d_803_v150_), (self).f24, default__.fm3(globalState), globalState))
                (globalState).f2 = (d_804_v151_).cf19
                (globalState).f18 = ((d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))]) - ((self).f24)
                d_774_cf28_ = not(((default__.fm0((d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))], (d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))], globalState)) - ((self).f24)) >= (default__.safeModuloInt((self).f24, (d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))])))
                d_805_v152_: _dafny.MultiSet
                d_805_v152_ = _dafny.MultiSet([(self).f24, (d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))], (self).f24])
                d_806_v153_: T0
                nw120_ = C1()
                nw120_.ctor__(_dafny.SeqWithoutIsStrInference([d_805_v152_, (d_805_v152_).set((d_805_v152_).cardinality, default__.abs(len(d_800_v147_))), (d_805_v152_).set((d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))], default__.abs((self).f24)), d_805_v152_, d_805_v152_]), len(d_731_v89_), (self).f24, (self).f25)
                d_806_v153_ = nw120_
                index106_ = default__.safeIndex(409, (d_775_v127_).length(0))
                rhs126_ = d_806_v153_
                rhs127_ = (len(_dafny.Map({d_775_v127_: (d_806_v153_).f24}))) >= ((d_805_v152_).cardinality)
                rhs128_ = self.f31
                rhs129_ = d_731_v89_
                rhs130_ = default__.safeModuloInt((d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))], (0) - ((d_806_v153_).f24))
                lhs99_ = globalState
                lhs100_ = globalState
                lhs101_ = d_775_v127_
                lhs102_ = default__.safeIndex(409, (d_775_v127_).length(0))
                d_806_v153_ = rhs126_
                lhs99_.f2 = rhs127_
                r0 = rhs128_
                lhs100_.f17 = rhs129_
                lhs101_[lhs102_] = rhs130_
                index107_ = default__.safeIndex(409, (d_775_v127_).length(0))
                (d_775_v127_)[index107_] = (d_775_v127_)[default__.safeIndex(409, (d_775_v127_).length(0))]
        elif source9_.is_DC18:
            d_807___mcc_h12_ = source9_.cf29
            d_808_cf29_ = d_807___mcc_h12_
            (globalState).f1 = (self).f24
            d_809_v154_: _dafny.Map
            d_809_v154_ = _dafny.Map({True: (self).f24})
            d_809_v154_ = (d_809_v154_).set(d_808_cf29_, (188) - ((self).f24))
            d_810_v155_: _dafny.Array
            nw121_ = _dafny.Array(False, 7)
            d_810_v155_ = nw121_
            index108_ = default__.safeIndex(152, (d_810_v155_).length(0))
            (d_810_v155_)[index108_] = self.f31
            d_811_v156_: _dafny.Set
            d_811_v156_ = _dafny.Set({d_810_v155_})
            d_812_v157_: _dafny.Seq
            d_812_v157_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24])
            index109_ = default__.safeIndex(152, (d_810_v155_).length(0))
            rhs131_ = (self).f24
            rhs132_ = (d_812_v157_)[default__.safeIndex((self).f24, len(d_812_v157_))]
            rhs133_ = d_808_cf29_
            rhs134_ = (d_811_v156_) | (d_811_v156_)
            lhs103_ = globalState
            lhs104_ = globalState
            lhs105_ = d_810_v155_
            lhs106_ = default__.safeIndex(152, (d_810_v155_).length(0))
            lhs103_.f18 = rhs131_
            lhs104_.f14 = rhs132_
            lhs105_[lhs106_] = rhs133_
            d_811_v156_ = rhs134_
            (globalState).f14 = (self).f24
        elif source9_.is_DC16:
            d_813___mcc_h13_ = source9_.cf27
            d_814_cf27_ = d_813___mcc_h13_
            d_815_v158_: _dafny.Map
            d_815_v158_ = _dafny.Map({self.f31: (self).f24})
            d_816_v159_: int
            out24_: int
            out24_ = default__.m0(default__.fm0((self).f24, (self).f24, globalState), ((d_815_v158_)[self.f31] if (self.f31) in (d_815_v158_) else 464), globalState)
            d_816_v159_ = out24_
            d_817_v160_: D1
            d_817_v160_ = D1_DC4(self.f31, self.f31, _dafny.CodePoint('i'), 983, (self).f24)
            d_818_v161_: _dafny.Map
            d_818_v161_ = _dafny.Map({d_817_v160_: self.f31})
            d_819_v163_: _dafny.Array
            nw122_ = _dafny.Array(None, 12)
            nw122_[int(0)] = _dafny.Map({d_817_v160_: self.f31})
            nw122_[int(1)] = d_818_v161_
            nw122_[int(2)] = (d_818_v161_) | (d_818_v161_)
            nw122_[int(3)] = (_dafny.Map({d_817_v160_: self.f31})) | (d_818_v161_)
            nw122_[int(4)] = (d_818_v161_) | (d_818_v161_)
            nw122_[int(5)] = d_818_v161_
            nw122_[int(6)] = d_818_v161_
            nw122_[int(7)] = d_818_v161_
            def iife45_():
                coll25_ = _dafny.Map()
                compr_25_: D1
                for compr_25_ in (d_818_v161_).keys.Elements:
                    d_820_v162_: D1 = compr_25_
                    if (d_820_v162_) in (d_818_v161_):
                        coll25_[d_820_v162_] = self.f31
                return _dafny.Map(coll25_)
            nw122_[int(8)] = iife45_()
            
            nw122_[int(9)] = (d_818_v161_) | (d_818_v161_)
            nw122_[int(10)] = (d_818_v161_) | (_dafny.Map({d_817_v160_: self.f31}))
            nw122_[int(11)] = (d_818_v161_) | (_dafny.Map({d_817_v160_: default__.fm2(d_816_v159_, d_816_v159_, (d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))], globalState)}))
            d_819_v163_ = nw122_
            index110_ = default__.safeIndex(469, (d_819_v163_).length(0))
            (d_819_v163_)[index110_] = (_dafny.Map({d_817_v160_: self.f31})) | (d_818_v161_)
            d_821_v164_: _dafny.MultiSet
            d_821_v164_ = _dafny.MultiSet([self.f31, self.f31])
            d_822_v165_: _dafny.Seq
            d_822_v165_ = _dafny.SeqWithoutIsStrInference([self.f31])
            d_823_v166_: D4
            d_823_v166_ = D4_DC13(d_816_v159_, d_815_v158_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "libuqyvxh")))
            index111_ = default__.safeIndex(469, (d_819_v163_).length(0))
            (d_819_v163_)[index111_] = default__.fm32((self).f24, _dafny.SeqWithoutIsStrInference([d_821_v164_, d_821_v164_, d_821_v164_, d_821_v164_]), d_822_v165_, (d_731_v89_).set(default__.safeIndex((d_823_v166_).cf22, len(d_731_v89_)), d_730_v88_), globalState)
            (globalState).f18 = default__.fm0((self).f24, d_816_v159_, globalState)
            d_824_v167_: _dafny.Array
            nw123_ = _dafny.Array(_dafny.Map({}), 9)
            d_824_v167_ = nw123_
            d_825_v168_: _dafny.Array
            nw124_ = _dafny.Array(None, 5)
            nw124_[int(0)] = (d_731_v89_) + (d_731_v89_)
            nw124_[int(1)] = (d_731_v89_ if self.f31 else default__.fm4(self.f31, d_816_v159_, (d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))], self.f31, globalState))
            nw124_[int(2)] = d_731_v89_
            nw124_[int(3)] = d_731_v89_
            nw124_[int(4)] = d_731_v89_
            d_825_v168_ = nw124_
            d_826_v169_: _dafny.Map
            d_826_v169_ = _dafny.Map({(self).f24: self.f31})
            d_827_v170_: _dafny.Seq
            d_827_v170_ = _dafny.SeqWithoutIsStrInference([d_816_v159_])
            rhs135_ = (0) - ((((self).f24) * (((d_821_v164_)[default__.fm2(default__.fm0((self).f24, d_816_v159_, globalState), (self).f24, d_730_v88_, globalState)] if (default__.fm2(default__.fm0((self).f24, d_816_v159_, globalState), (self).f24, d_730_v88_, globalState)) in (d_821_v164_) else d_816_v159_))) + ((d_816_v159_) * (d_816_v159_)))
            rhs136_ = (d_824_v167_ if False else d_824_v167_)
            rhs137_ = (default__.fm16(globalState)) == ((d_826_v169_) | (d_826_v169_))
            rhs138_ = ((d_827_v170_)[default__.safeIndex(d_816_v159_, len(d_827_v170_))]) * (len(_dafny.Map({d_816_v159_: self.f31})))
            rhs139_ = d_825_v168_
            lhs107_ = globalState
            lhs108_ = globalState
            lhs107_.f18 = rhs135_
            d_824_v167_ = rhs136_
            r0 = rhs137_
            lhs108_.f14 = rhs138_
            d_825_v168_ = rhs139_
        elif True:
            d_828___mcc_h14_ = source9_.cf30
            d_829_cf30_ = d_828___mcc_h14_
            d_830_v171_: D1
            d_830_v171_ = D1_DC6()
            d_830_v171_ = d_830_v171_
            d_831_v172_: D6
            d_831_v172_ = D6_DC17(self.f31)
            (globalState).f2 = ((d_831_v172_).cf28) and (self.f31)
            (globalState).f2 = self.f31
            d_832_v173_: _dafny.Seq
            d_832_v173_ = _dafny.SeqWithoutIsStrInference([len((d_731_v89_) + (_dafny.SeqWithoutIsStrInference([d_730_v88_ for d_833_i20_ in range(default__.abs(849))]))), ((self).f24) - ((self).f24)])
            d_832_v173_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24])
        d_834_v174_: _dafny.Array
        nw125_ = _dafny.Array(_dafny.Map({}), 18)
        d_834_v174_ = nw125_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_834_v174_).length(0)):
            d_835_i21_: int = guard_loop_1_
            if (True) and (((0) <= (d_835_i21_)) and ((d_835_i21_) < ((d_834_v174_).length(0)))):
                (d_834_v174_)[(d_835_i21_)] = _dafny.Map({self.f31: default__.safeDivisionInt((self).f24, (self).f24)})
        d_836_i22_: int
        d_836_i22_ = 0
        with _dafny.label("1"):
            while not(self.f31):
                with _dafny.c_label("1"):
                    if (d_836_i22_) >= (100):
                        raise _dafny.Break("1")
                    d_836_i22_ = (d_836_i22_) + (1)
                    if self.f31:
                        d_837_v175_: _dafny.Array
                        nw126_ = _dafny.Array(int(0), 27)
                        d_837_v175_ = nw126_
                        d_838_v176_: _dafny.Map
                        d_838_v176_ = _dafny.Map({d_837_v175_: self.f31})
                        index112_ = default__.safeIndex(849, (d_837_v175_).length(0))
                        (d_837_v175_)[index112_] = len((d_838_v176_) | ((_dafny.Map({d_837_v175_: self.f31})).set(d_837_v175_, self.f31)))
                        d_839_v177_: _dafny.MultiSet
                        d_839_v177_ = _dafny.MultiSet([(self).f24])
                        d_840_v178_: _dafny.Map
                        d_840_v178_ = _dafny.Map({default__.fm0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfeykmiww"))), (self).f24, globalState): self.f31})
                        index113_ = default__.safeIndex(849, (d_837_v175_).length(0))
                        (d_837_v175_)[index113_] = (((d_839_v177_)[(self).f24] if ((self).f24) in (d_839_v177_) else (self).f24) if self.f31 else len(d_840_v178_))
                        (globalState).f9 = (d_837_v175_)[default__.safeIndex(849, (d_837_v175_).length(0))]
                        (globalState).f2 = ((d_837_v175_)[default__.safeIndex(849, (d_837_v175_).length(0))]) > ((519) + ((0) - ((d_837_v175_)[default__.safeIndex(849, (d_837_v175_).length(0))])))
                        d_841_v179_: _dafny.Map
                        d_841_v179_ = _dafny.Map({d_837_v175_: (self).f24})
                        d_842_v180_: _dafny.Seq
                        d_842_v180_ = _dafny.SeqWithoutIsStrInference([d_837_v175_])
                        d_843_v181_: _dafny.Set
                        d_843_v181_ = _dafny.Set({self.f31})
                        d_841_v179_ = (d_841_v179_).set((d_842_v180_)[default__.safeIndex(default__.fm0(default__.fm0(len(d_843_v181_), (d_837_v175_)[default__.safeIndex(849, (d_837_v175_).length(0))], globalState), (self).f24, globalState), len(d_842_v180_))], (self).f24)
                        d_844_v182_: _dafny.Set
                        d_844_v182_ = _dafny.Set({(self).f24})
                        d_845_v183_: _dafny.Map
                        d_845_v183_ = _dafny.Map({self.f31: len(d_844_v182_)})
                        (globalState).f2 = (self).fm9(self.f31, default__.fm0(len(d_845_v183_), (d_837_v175_)[default__.safeIndex(849, (d_837_v175_).length(0))], globalState), (self).f24, globalState)
                    elif True:
                        d_846_v184_: _dafny.Map
                        d_846_v184_ = _dafny.Map({(self).fm9(self.f31, (self).f24, 593, globalState): self.f31})
                        d_846_v184_ = (d_846_v184_).set(False, False)
                        (globalState).f2 = self.f31
                        d_847_v185_: _dafny.Map
                        d_847_v185_ = _dafny.Map({(self).f24: (self).f24})
                        d_848_v186_: _dafny.Seq
                        d_848_v186_ = _dafny.SeqWithoutIsStrInference([((_dafny.Map({(self).f24: len(d_846_v184_)})).set((self).f24, -306)) | (d_847_v185_), d_847_v185_])
                        d_848_v186_ = (d_848_v186_) + (d_848_v186_)
                        (globalState).f2 = not (self.f31) or ((self.f31) == (self.f31))
                        d_849_v187_: int
                        out25_: int
                        out25_ = default__.m0((self).f24, (0) - ((self).f24), globalState)
                        d_849_v187_ = out25_
                    d_850_v188_: _dafny.MultiSet
                    d_850_v188_ = _dafny.MultiSet([(self).f24])
                    d_851_v189_: C0
                    nw127_ = C0()
                    nw127_.ctor__((_dafny.MultiSet([(self).f24, len(d_731_v89_)])).issubset(d_850_v188_))
                    d_851_v189_ = nw127_
                    if ((self).f24) > ((self).f24):
                        d_852_v190_: _dafny.Set
                        d_852_v190_ = _dafny.Set({len(d_731_v89_), (0) - ((self).f24)})
                        d_853_v191_: D6
                        d_853_v191_ = D6_DC16(d_852_v190_)
                        d_853_v191_ = d_853_v191_
                        d_854_v192_: _dafny.Map
                        d_854_v192_ = _dafny.Map({(self).f24: default__.fm18(globalState)})
                        d_855_v194_: _dafny.Seq
                        def iife46_():
                            coll26_ = _dafny.Set()
                            compr_26_: int
                            for compr_26_ in (d_850_v188_).Elements:
                                d_856_v193_: int = compr_26_
                                if (d_856_v193_) in (d_850_v188_):
                                    coll26_ = coll26_.union(_dafny.Set([(d_856_v193_) + (918)]))
                            return _dafny.Set(coll26_)
                        d_855_v194_ = _dafny.SeqWithoutIsStrInference([d_852_v190_, iife46_()
                        , d_852_v190_])
                        d_857_v195_: _dafny.Seq
                        d_857_v195_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                        d_852_v190_ = ((d_854_v192_)[(self).f24] if ((self).f24) in (d_854_v192_) else ((d_855_v194_)[default__.safeIndex(len(d_857_v195_), len(d_855_v194_))] if default__.fm2((self).f24, (self).f24, (d_728_v87_)[default__.safeIndex(812, (d_728_v87_).length(0))], globalState) else d_852_v190_))
                        d_858_v196_: _dafny.Map
                        d_858_v196_ = _dafny.Map({default__.fm23(globalState): default__.fm2((self).f24, (0) - ((self).f24), _dafny.CodePoint('n'), globalState)})
                        d_859_v197_: _dafny.Map
                        d_859_v197_ = _dafny.Map({self.f31: (d_851_v189_).f34})
                        d_860_v198_: _dafny.Seq
                        d_860_v198_ = _dafny.SeqWithoutIsStrInference([(d_851_v189_).f34])
                        d_858_v196_ = (d_858_v196_).set((d_860_v198_ if ((d_859_v197_)[(self).fm9((d_851_v189_).f34, (self).f24, (self).f24, globalState)] if ((self).fm9((d_851_v189_).f34, (self).f24, (self).f24, globalState)) in (d_859_v197_) else (d_851_v189_).f34) else d_860_v198_), (d_851_v189_).f34)
                        d_861_v199_: D1
                        d_861_v199_ = D1_DC3(d_731_v89_)
                        d_861_v199_ = D1_DC3(d_731_v89_)
                        (self).f31 = self.f31
                    elif True:
                        d_862_v200_: _dafny.Set
                        d_862_v200_ = _dafny.Set({104})
                        d_863_v201_: _dafny.Map
                        d_863_v201_ = _dafny.Map({(d_851_v189_).f34: d_862_v200_})
                        d_864_v202_: _dafny.Seq
                        d_864_v202_ = _dafny.SeqWithoutIsStrInference([not((d_851_v189_).f34)])
                        d_865_v203_: _dafny.Map
                        d_865_v203_ = _dafny.Map({d_863_v201_: (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([True])) + (d_864_v202_))).cardinality})
                        d_865_v203_ = (d_865_v203_).set(d_863_v201_, (self).f24)
                        d_866_v204_: int
                        out26_: int
                        out26_ = default__.m0((0) - ((self).f24), (self).f24, globalState)
                        d_866_v204_ = out26_
                        d_867_v205_: _dafny.Array
                        def lambda33_(d_868_i23_):
                            return not(self.f31)

                        init20_ = lambda33_
                        nw128_ = _dafny.Array(None, 26)
                        for i0_20_ in range(nw128_.length(0)):
                            nw128_[i0_20_] = init20_(i0_20_)
                        d_867_v205_ = nw128_
                        rhs140_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlpy")))) * ((0) - ((self).f24))
                        rhs141_ = d_867_v205_
                        d_866_v204_ = rhs140_
                        d_867_v205_ = rhs141_
                        d_869_v206_: _dafny.Seq
                        d_869_v206_ = _dafny.SeqWithoutIsStrInference([d_866_v204_])
                        d_869_v206_ = (d_869_v206_) + (_dafny.SeqWithoutIsStrInference([d_866_v204_ for d_870_i24_ in range(default__.abs(134))]))
                        (globalState).f2 = (d_851_v189_).f34
                    d_871_v207_: _dafny.Array
                    def lambda34_(d_872_v189_):
                        def lambda35_(d_873_i25_):
                            return (_dafny.MultiSet([(d_872_v189_).f34])) | (_dafny.MultiSet([self.f31]))

                        return lambda35_

                    init21_ = lambda34_(d_851_v189_)
                    nw129_ = _dafny.Array(None, 23)
                    for i0_21_ in range(nw129_.length(0)):
                        nw129_[i0_21_] = init21_(i0_21_)
                    d_871_v207_ = nw129_
                    d_871_v207_ = d_871_v207_
                    pass
            pass
        r0 = self.f31
        return r0


class C3(T0):
    def  __init__(self):
        self._f24: int = int(0)
        self._f25: _dafny.Seq = _dafny.Seq({})
        self.f30: _dafny.Array = _dafny.Array(None, 0)
        self._f29: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f29, f30, f24, f25):
        (self)._f29 = f29
        (self).f30 = f30
        (self)._f24 = f24
        (self)._f25 = f25

    def fm8(self, globalState):
        return (D0_DC2(False, (self).f24, _dafny.CodePoint('h'))).cf5

    def m1(self, globalState):
        r0: bool = False
        d_874_v0_: _dafny.Seq
        d_874_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olddh"))
        d_875_v1_: str
        d_875_v1_ = _dafny.CodePoint('w')
        d_876_v2_: _dafny.Seq
        d_876_v2_ = _dafny.SeqWithoutIsStrInference([(d_874_v0_).set(default__.safeIndex((self).f24, len(d_874_v0_)), d_875_v1_)])
        hi5_ = len(d_876_v2_)
        for d_877_i0_ in range((self).f24, hi5_):
            d_878_v3_: bool
            d_878_v3_ = True
            (globalState).f17 = (default__.fm4(False, (self).f24, d_875_v1_, d_878_v3_, globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pglhejcp")))
            d_879_v4_: _dafny.Array
            def lambda36_(d_880_v3_):
                def lambda37_(d_881_i1_):
                    return (d_880_v3_ if d_880_v3_ else d_880_v3_)

                return lambda37_

            init22_ = lambda36_(d_878_v3_)
            nw130_ = _dafny.Array(None, 21)
            for i0_22_ in range(nw130_.length(0)):
                nw130_[i0_22_] = init22_(i0_22_)
            d_879_v4_ = nw130_
            nw131_ = _dafny.Array(False, 2)
            d_879_v4_ = nw131_
            d_882_v5_: _dafny.Seq
            d_882_v5_ = _dafny.SeqWithoutIsStrInference([d_878_v3_])
            d_883_v6_: _dafny.MultiSet
            d_883_v6_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_878_v3_])])
            d_884_v7_: _dafny.Seq
            d_884_v7_ = _dafny.SeqWithoutIsStrInference([d_883_v6_, _dafny.MultiSet([d_882_v5_, d_882_v5_]), d_883_v6_])
            if ((d_882_v5_) + (d_882_v5_)) in ((d_884_v7_)[default__.safeIndex((self).f24, len(d_884_v7_))]):
                d_885_v8_: _dafny.Array
                nw132_ = _dafny.Array(_dafny.CodePoint('D'), 5)
                d_885_v8_ = nw132_
                d_886_v9_: D2
                d_886_v9_ = D2_DC7(d_885_v8_)
                d_887_v10_: _dafny.Array
                nw133_ = _dafny.Array(None, 6)
                nw133_[int(0)] = d_885_v8_
                nw133_[int(1)] = d_885_v8_
                nw133_[int(2)] = d_885_v8_
                nw133_[int(3)] = d_885_v8_
                nw133_[int(4)] = d_885_v8_
                nw133_[int(5)] = (d_886_v9_).cf16
                d_887_v10_ = nw133_
                d_887_v10_ = d_887_v10_
                d_888_v11_: C0
                nw134_ = C0()
                nw134_.ctor__((d_882_v5_) != (_dafny.SeqWithoutIsStrInference([True, False, d_878_v3_, d_878_v3_, d_878_v3_])))
                d_888_v11_ = nw134_
                (globalState).f18 = d_877_i0_
                index114_ = default__.safeIndex(917, (d_879_v4_).length(0))
                (d_879_v4_)[index114_] = True
                index115_ = default__.safeIndex(917, (d_879_v4_).length(0))
                (d_879_v4_)[index115_] = (d_874_v0_) < (((d_874_v0_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apcb"))), len(d_874_v0_)), d_875_v1_)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_889_i2_ in range(default__.abs(426))])))
                d_890_v12_: _dafny.MultiSet
                d_890_v12_ = _dafny.MultiSet([not((d_888_v11_).f34), (d_879_v4_)[default__.safeIndex(917, (d_879_v4_).length(0))], (d_888_v11_).f34])
                d_891_v13_: _dafny.Set
                d_891_v13_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_875_v1_ for d_892_i3_ in range(default__.abs(711))])})
                rhs142_ = ((self).f24 if (d_890_v12_).isdisjoint(_dafny.MultiSet([(d_888_v11_).f34])) else len(d_874_v0_))
                rhs143_ = ((d_891_v13_).issubset(default__.fm22(d_877_i0_, (d_888_v11_).f34, (d_888_v11_).f34, len(d_874_v0_), globalState))) and ((d_888_v11_).f34)
                rhs144_ = (self).f24
                lhs109_ = globalState
                lhs110_ = globalState
                lhs111_ = globalState
                lhs109_.f14 = rhs142_
                lhs110_.f2 = rhs143_
                lhs111_.f9 = rhs144_
            elif True:
                d_893_v14_: _dafny.Map
                d_893_v14_ = _dafny.Map({d_874_v0_: d_878_v3_})
                d_893_v14_ = (d_893_v14_).set(d_874_v0_, d_878_v3_)
                d_894_v15_: _dafny.Set
                d_894_v15_ = _dafny.Set({d_878_v3_})
                d_895_v16_: D0
                d_895_v16_ = D0_DC2((d_877_i0_) > (len(d_894_v15_)), len(_dafny.SeqWithoutIsStrInference([d_877_i0_ for d_896_i4_ in range(default__.abs(-589))])), _dafny.CodePoint('g'))
                d_895_v16_ = d_895_v16_
                d_897_v18_: _dafny.Array
                def lambda38_(d_898_i5_):
                    def iife47_():
                        coll27_ = _dafny.Set()
                        compr_27_: int
                        for compr_27_ in _dafny.IntegerRange(291, -29):
                            d_899_v17_: int = compr_27_
                            if ((291) <= (d_899_v17_)) and ((d_899_v17_) < (-29)):
                                coll27_ = coll27_.union(_dafny.Set([default__.safeModuloInt(d_899_v17_, (self).f24)]))
                        return _dafny.Set(coll27_)
                    return iife47_()
                    

                init23_ = lambda38_
                nw135_ = _dafny.Array(None, 9)
                for i0_23_ in range(nw135_.length(0)):
                    nw135_[i0_23_] = init23_(i0_23_)
                d_897_v18_ = nw135_
                rhs145_ = (self).f29
                rhs146_ = len(((d_874_v0_) + (d_874_v0_)) + (d_874_v0_))
                lhs112_ = globalState
                d_897_v18_ = rhs145_
                lhs112_.f1 = rhs146_
                d_900_v19_: D8
                d_900_v19_ = D8_DC23(d_894_v15_)
                d_900_v19_ = D8_DC23(d_894_v15_)
                d_901_v20_: _dafny.Map
                d_901_v20_ = _dafny.Map({default__.fm0((self).f24, (self).f24, globalState): (0) - ((0) - ((len(d_882_v5_)) + ((self).f24)))})
                d_901_v20_ = (d_901_v20_).set(d_877_i0_, default__.safeModuloInt(d_877_i0_, (0) - (len(d_874_v0_))))
            (globalState).f18 = default__.fm0(d_877_i0_, ((self).f24) + (d_877_i0_), globalState)
        hi6_ = (self).f24
        for d_902_i6_ in range(len(d_874_v0_), hi6_):
            (globalState).f2 = (len(d_874_v0_)) == ((d_902_i6_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhg")))))
            d_903_v21_: bool
            d_903_v21_ = True
            d_904_v22_: _dafny.Map
            d_904_v22_ = _dafny.Map({d_903_v21_: d_903_v21_})
            d_904_v22_ = (d_904_v22_).set(((self).f24) != (d_902_i6_), d_903_v21_)
            d_874_v0_ = (d_874_v0_) + ((_dafny.SeqWithoutIsStrInference([d_875_v1_ for d_905_i7_ in range(default__.abs(481))]) if d_903_v21_ else d_874_v0_))
            (globalState).f2 = d_903_v21_
        d_906_v23_: int
        out27_: int
        out27_ = default__.m0((self).f24, (self).f24, globalState)
        d_906_v23_ = out27_
        d_907_v24_: bool
        d_907_v24_ = True
        d_908_v25_: _dafny.MultiSet
        d_908_v25_ = _dafny.MultiSet([d_907_v24_])
        d_909_v26_: _dafny.Seq
        d_909_v26_ = _dafny.SeqWithoutIsStrInference([(d_908_v25_).cardinality, (self).f24, d_906_v23_])
        d_910_v27_: _dafny.Map
        d_910_v27_ = _dafny.Map({d_907_v24_: d_909_v26_})
        d_911_v28_: _dafny.Map
        d_911_v28_ = _dafny.Map({len(((d_910_v27_)[d_907_v24_] if (d_907_v24_) in (d_910_v27_) else d_909_v26_)): d_907_v24_})
        d_912_v29_: _dafny.MultiSet
        d_912_v29_ = _dafny.MultiSet([len((_dafny.Map({d_911_v28_: d_906_v23_})) | (_dafny.Map({d_911_v28_: d_906_v23_})))])
        d_912_v29_ = (_dafny.MultiSet([d_906_v23_])).intersection(d_912_v29_)
        source10_ = default__.fm33(d_907_v24_, (0) - ((self).f24), globalState)
        if source10_.is_DC10:
            d_907_v24_ = d_907_v24_
            d_913_v30_: _dafny.Array
            nw136_ = _dafny.Array(int(0), 11)
            d_913_v30_ = nw136_
            d_914_v31_: _dafny.Set
            d_914_v31_ = _dafny.Set({d_913_v30_, d_913_v30_})
            d_915_v32_: _dafny.Map
            d_915_v32_ = _dafny.Map({not(d_907_v24_): d_907_v24_})
            d_916_v33_: _dafny.Set
            d_916_v33_ = _dafny.Set({not(default__.fm2(872, (self).f24, _dafny.CodePoint('x'), globalState))})
            d_917_v34_: _dafny.Seq
            d_917_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_909_v26_)])
            d_918_v35_: _dafny.Map
            d_918_v35_ = _dafny.Map({d_875_v1_: (0) - ((self).f24)})
            d_919_v36_: C1
            nw137_ = C1()
            nw137_.ctor__(d_917_v34_, (self).f24, (self).f24, _dafny.SeqWithoutIsStrInference([d_918_v35_]))
            d_919_v36_ = nw137_
            d_920_v37_: _dafny.Map
            d_920_v37_ = _dafny.Map({default__.fm2(len(d_915_v32_), d_906_v23_, d_875_v1_, globalState): (0) - (d_906_v23_)})
            d_921_v38_: _dafny.Map
            d_921_v38_ = _dafny.Map({d_919_v36_: d_920_v37_})
            d_922_v39_: _dafny.Array
            nw138_ = _dafny.Array(None, 20)
            nw138_[int(0)] = len(d_914_v31_)
            nw138_[int(1)] = ((self).f24) * (d_906_v23_)
            nw138_[int(2)] = default__.safeDivisionInt((0) - ((self).f24), (self).f24)
            nw138_[int(3)] = (0) - (((self).f24) * (d_906_v23_))
            nw138_[int(4)] = (self).f24
            nw138_[int(5)] = d_906_v23_
            nw138_[int(6)] = len((d_915_v32_) | (d_915_v32_))
            nw138_[int(7)] = len(d_916_v33_)
            nw138_[int(8)] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwuajdbrj"))), (self).f24)
            nw138_[int(9)] = len(d_921_v38_)
            nw138_[int(10)] = (D1_DC4(d_907_v24_, d_907_v24_, d_875_v1_, (self).f24, (self).fm8(globalState))).cf12
            nw138_[int(11)] = d_919_v36_.f33
            nw138_[int(12)] = (self).f24
            nw138_[int(13)] = (self).f24
            nw138_[int(14)] = ((self).f24) - (d_919_v36_.f33)
            nw138_[int(15)] = len(d_916_v33_)
            nw138_[int(16)] = (self).f24
            nw138_[int(17)] = (d_909_v26_)[default__.safeIndex(d_906_v23_, len(d_909_v26_))]
            nw138_[int(18)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhs")))) + ((0) - (d_906_v23_))
            nw138_[int(19)] = (self).f24
            d_922_v39_ = nw138_
            d_923_v40_: _dafny.Map
            d_923_v40_ = _dafny.Map({d_920_v37_: (0) - (d_919_v36_.f33)})
            d_924_v41_: _dafny.Map
            d_924_v41_ = _dafny.Map({d_907_v24_: (d_920_v37_).set(d_907_v24_, d_919_v36_.f33)})
            nw139_ = _dafny.Array(None, 5)
            nw139_[int(0)] = (self).f24
            nw139_[int(1)] = ((d_923_v40_)[((d_924_v41_)[d_907_v24_] if (d_907_v24_) in (d_924_v41_) else d_920_v37_)] if (((d_924_v41_)[d_907_v24_] if (d_907_v24_) in (d_924_v41_) else d_920_v37_)) in (d_923_v40_) else (self).f24)
            nw139_[int(2)] = d_919_v36_.f33
            nw139_[int(3)] = (self).f24
            nw139_[int(4)] = d_906_v23_
            d_922_v39_ = nw139_
            d_925_v42_: _dafny.Array
            nw140_ = _dafny.Array(False, 25)
            d_925_v42_ = nw140_
            d_926_v43_: _dafny.Map
            d_926_v43_ = _dafny.Map({d_874_v0_: d_925_v42_})
            d_925_v42_ = ((d_926_v43_)[_dafny.SeqWithoutIsStrInference([d_875_v1_ for d_927_i8_ in range(default__.abs(-401))])] if (_dafny.SeqWithoutIsStrInference([d_875_v1_ for d_928_i8_ in range(default__.abs(-401))])) in (d_926_v43_) else d_925_v42_)
            (globalState).f17 = (d_874_v0_) + (d_874_v0_)
        elif source10_.is_DC11:
            d_929___mcc_h0_ = source10_.cf18
            d_930___mcc_h1_ = source10_.cf19
            d_931___mcc_h2_ = source10_.cf20
            d_932_cf20_ = d_931___mcc_h2_
            d_933_cf19_ = d_930___mcc_h1_
            d_934_cf18_ = d_929___mcc_h0_
            d_935_v44_: _dafny.Map
            d_935_v44_ = _dafny.Map({d_907_v24_: not (not(True)) or (True)})
            (globalState).f2 = ((d_935_v44_)[default__.fm2(d_906_v23_, d_906_v23_, d_875_v1_, globalState)] if (default__.fm2(d_906_v23_, d_906_v23_, d_875_v1_, globalState)) in (d_935_v44_) else d_933_cf19_)
            d_936_v45_: _dafny.Array
            nw141_ = _dafny.Array(None, 20)
            nw141_[int(0)] = d_933_cf19_
            nw141_[int(1)] = d_907_v24_
            nw141_[int(2)] = not(d_932_cf20_)
            nw141_[int(3)] = d_933_cf19_
            nw141_[int(4)] = d_933_cf19_
            nw141_[int(5)] = d_907_v24_
            nw141_[int(6)] = d_907_v24_
            nw141_[int(7)] = d_933_cf19_
            nw141_[int(8)] = False
            nw141_[int(9)] = True
            nw141_[int(10)] = d_907_v24_
            nw141_[int(11)] = d_932_cf20_
            nw141_[int(12)] = True
            nw141_[int(13)] = d_907_v24_
            nw141_[int(14)] = d_907_v24_
            nw141_[int(15)] = not(False)
            nw141_[int(16)] = d_907_v24_
            nw141_[int(17)] = d_932_cf20_
            nw141_[int(18)] = d_932_cf20_
            nw141_[int(19)] = d_933_cf19_
            d_936_v45_ = nw141_
            d_937_v46_: _dafny.Seq
            d_937_v46_ = _dafny.SeqWithoutIsStrInference([d_936_v45_])
            d_937_v46_ = d_937_v46_
            index116_ = default__.safeIndex(8, (d_936_v45_).length(0))
            (d_936_v45_)[index116_] = d_933_cf19_
            d_938_v47_: _dafny.Seq
            d_938_v47_ = _dafny.SeqWithoutIsStrInference([d_933_cf19_, d_907_v24_])
            d_939_v48_: _dafny.MultiSet
            d_939_v48_ = _dafny.MultiSet([d_938_v47_])
            index117_ = default__.safeIndex(8, (d_936_v45_).length(0))
            (d_936_v45_)[index117_] = ((d_939_v48_ if d_933_cf19_ else d_939_v48_)).issubset(d_939_v48_)
            index118_ = default__.safeIndex(8, (d_936_v45_).length(0))
            (d_936_v45_)[index118_] = (d_936_v45_)[default__.safeIndex(8, (d_936_v45_).length(0))]
        elif True:
            d_940___mcc_h3_ = source10_.cf17
            d_941_cf17_ = d_940___mcc_h3_
            d_942_v49_: _dafny.Seq
            d_942_v49_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_906_v23_]), _dafny.MultiSet([len(d_874_v0_)]), d_912_v29_, d_912_v29_, d_912_v29_])
            d_943_v51_: C1
            nw142_ = C1()
            def iife48_():
                coll28_ = _dafny.Map()
                compr_28_: str
                for compr_28_ in (d_874_v0_).Elements:
                    d_944_v50_: str = compr_28_
                    if (d_944_v50_) in (d_874_v0_):
                        coll28_[d_944_v50_] = d_906_v23_
                return _dafny.Map(coll28_)
            nw142_.ctor__(d_942_v49_, len(d_874_v0_), (self).f24, _dafny.SeqWithoutIsStrInference([iife48_()
            ]))
            d_943_v51_ = nw142_
            d_945_v52_: _dafny.Set
            d_945_v52_ = _dafny.Set({(d_909_v26_)[default__.safeIndex(494, len(d_909_v26_))]})
            d_946_v54_: _dafny.Map
            def iife49_():
                coll29_ = _dafny.Set()
                compr_29_: int
                for compr_29_ in _dafny.IntegerRange(350, 677):
                    d_947_v53_: int = compr_29_
                    if ((350) <= (d_947_v53_)) and ((d_947_v53_) < (677)):
                        coll29_ = coll29_.union(_dafny.Set([(d_947_v53_) * (d_906_v23_)]))
                return _dafny.Set(coll29_)
            d_946_v54_ = _dafny.Map({len((d_945_v52_) | (iife49_()
            )): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsixn"))})
            rhs147_ = ((d_946_v54_)[d_943_v51_.f33] if (d_943_v51_.f33) in (d_946_v54_) else d_874_v0_)
            rhs148_ = (d_906_v23_ if default__.fm2(len(d_874_v0_), 189, d_875_v1_, globalState) else (self).f24)
            rhs149_ = False
            lhs113_ = globalState
            lhs114_ = globalState
            lhs113_.f17 = rhs147_
            lhs114_.f14 = rhs148_
            d_907_v24_ = rhs149_
            d_948_v55_: _dafny.Array
            nw143_ = _dafny.Array(None, 28)
            nw143_[int(0)] = d_907_v24_
            nw143_[int(1)] = d_907_v24_
            nw143_[int(2)] = d_907_v24_
            nw143_[int(3)] = d_907_v24_
            nw143_[int(4)] = d_907_v24_
            nw143_[int(5)] = d_907_v24_
            nw143_[int(6)] = d_907_v24_
            nw143_[int(7)] = d_907_v24_
            nw143_[int(8)] = ((self).f24) == ((self).f24)
            nw143_[int(9)] = d_907_v24_
            nw143_[int(10)] = (_dafny.MultiSet([d_907_v24_, d_907_v24_])).isdisjoint((d_908_v25_).set(False, default__.abs(d_943_v51_.f33)))
            nw143_[int(11)] = d_907_v24_
            nw143_[int(12)] = d_907_v24_
            nw143_[int(13)] = d_907_v24_
            nw143_[int(14)] = d_907_v24_
            nw143_[int(15)] = (d_943_v51_.f33) != (d_943_v51_.f33)
            nw143_[int(16)] = d_907_v24_
            nw143_[int(17)] = not(d_907_v24_)
            nw143_[int(18)] = (d_941_cf17_)[default__.safeIndex(d_906_v23_, len(d_941_cf17_))]
            nw143_[int(19)] = d_907_v24_
            nw143_[int(20)] = d_907_v24_
            nw143_[int(21)] = (d_907_v24_ if d_907_v24_ else d_907_v24_)
            nw143_[int(22)] = not(d_907_v24_)
            nw143_[int(23)] = d_907_v24_
            nw143_[int(24)] = d_907_v24_
            nw143_[int(25)] = d_907_v24_
            nw143_[int(26)] = d_907_v24_
            nw143_[int(27)] = d_907_v24_
            d_948_v55_ = nw143_
            d_948_v55_ = d_948_v55_
            d_949_v56_: C0
            nw144_ = C0()
            nw144_.ctor__((d_874_v0_) < (d_874_v0_))
            d_949_v56_ = nw144_
        d_950_v57_: _dafny.Map
        d_950_v57_ = _dafny.Map({(0) - (len(d_874_v0_)): d_906_v23_})
        d_951_v58_: _dafny.Map
        d_951_v58_ = _dafny.Map({d_906_v23_: d_950_v57_})
        d_952_v59_: _dafny.Map
        d_952_v59_ = _dafny.Map({d_907_v24_: d_906_v23_})
        d_953_v60_: D4
        d_953_v60_ = D4_DC13(len(d_951_v58_), d_952_v59_, d_874_v0_)
        d_954_v61_: _dafny.Map
        d_954_v61_ = _dafny.Map({d_953_v60_: _dafny.SeqWithoutIsStrInference([d_875_v1_ for d_955_i9_ in range(default__.abs(496))])})
        (globalState).f1 = len((((d_954_v61_)[d_953_v60_] if (d_953_v60_) in (d_954_v61_) else d_874_v0_) if (d_907_v24_ if d_907_v24_ else d_907_v24_) else d_874_v0_))
        r0 = True
        return r0

    def m5(self, p0, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_956_v0_: _dafny.Seq
        d_956_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        d_957_v1_: D3
        d_957_v1_ = D3_DC9((d_956_v0_).set(default__.safeIndex(852, len(d_956_v0_)), p0))
        d_958_i0_: int
        d_958_i0_ = 0
        with _dafny.label("2"):
            pat_let_tv6_ = p0
            def lambda42_(source12_):
                if source12_.is_DC10:
                    return pat_let_tv6_
                elif source12_.is_DC11:
                    d_1035___mcc_h7_ = source12_.cf18
                    d_1036___mcc_h8_ = source12_.cf19
                    d_1037___mcc_h9_ = source12_.cf20
                    d_1038_cf20_ = d_1037___mcc_h9_
                    d_1039_cf19_ = d_1036___mcc_h8_
                    d_1040_cf18_ = d_1035___mcc_h7_
                    return d_1039_cf19_
                elif True:
                    d_1041___mcc_h10_ = source12_.cf17
                    d_1042_cf17_ = d_1041___mcc_h10_
                    return True

            while lambda42_(d_957_v1_):
                with _dafny.c_label("2"):
                    if (d_958_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_958_i0_ = (d_958_i0_) + (1)
                    d_959_v2_: str
                    d_959_v2_ = _dafny.CodePoint('o')
                    d_960_v3_: _dafny.Seq
                    d_960_v3_ = _dafny.SeqWithoutIsStrInference([d_959_v2_, d_959_v2_, d_959_v2_])
                    d_961_v4_: D1
                    d_961_v4_ = D1_DC3((d_960_v3_).set(default__.safeIndex((self).f24, len(d_960_v3_)), d_959_v2_))
                    d_962_v5_: D8
                    d_962_v5_ = D8_DC26(d_961_v4_, d_960_v3_)
                    source11_ = d_962_v5_
                    if source11_.is_DC24:
                        d_963___mcc_h0_ = source11_.cf37
                        d_964_cf37_ = d_963___mcc_h0_
                        d_965_v6_: D0
                        d_965_v6_ = D0_DC1(980, (self).f24, p0)
                        d_966_v7_: _dafny.Seq
                        d_966_v7_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                        d_967_v8_: D1
                        d_967_v8_ = D1_DC5(p0, p0, p0)
                        d_968_v9_: D6
                        d_968_v9_ = D6_DC18(p0)
                        d_969_v10_: _dafny.Array
                        def lambda39_(d_970_p0_):
                            def lambda40_(d_971_i1_):
                                return d_970_p0_

                            return lambda40_

                        init24_ = lambda39_(p0)
                        nw145_ = _dafny.Array(None, 17)
                        for i0_24_ in range(nw145_.length(0)):
                            nw145_[i0_24_] = init24_(i0_24_)
                        d_969_v10_ = nw145_
                        d_972_v11_: D7
                        d_972_v11_ = D7_DC21(d_967_v8_, D6_DC19(d_968_v9_), d_969_v10_)
                        d_973_v12_: _dafny.Map
                        d_973_v12_ = _dafny.Map({d_966_v7_: (d_972_v11_).cf34})
                        pat_let_tv4_ = d_956_v0_
                        pat_let_tv5_ = d_956_v0_
                        d_974_v13_: _dafny.Array
                        nw146_ = _dafny.Array(None, 27)
                        nw146_[int(0)] = d_964_cf37_
                        def iife50_(_pat_let10_0):
                            def iife51_(d_975_dt__update__tmp_h0_):
                                def iife52_(_pat_let11_0):
                                    def iife53_(d_976_dt__update_hcf3_h0_):
                                        return D0_DC1((d_975_dt__update__tmp_h0_).cf1, (d_975_dt__update__tmp_h0_).cf2, d_976_dt__update_hcf3_h0_)
                                    return iife53_(_pat_let11_0)
                                return iife52_((pat_let_tv4_)[default__.safeIndex((self).f24, len(pat_let_tv5_))])
                            return iife51_(_pat_let10_0)
                        nw146_[int(1)] = (0) - ((iife50_(d_965_v6_)).cf1)
                        nw146_[int(2)] = d_964_cf37_
                        nw146_[int(3)] = d_964_cf37_
                        nw146_[int(4)] = d_964_cf37_
                        nw146_[int(5)] = (0) - ((self).f24)
                        nw146_[int(6)] = (self).f24
                        nw146_[int(7)] = (self).f24
                        nw146_[int(8)] = d_964_cf37_
                        nw146_[int(9)] = len(d_973_v12_)
                        nw146_[int(10)] = d_964_cf37_
                        nw146_[int(11)] = d_964_cf37_
                        nw146_[int(12)] = d_964_cf37_
                        nw146_[int(13)] = d_964_cf37_
                        nw146_[int(14)] = len(d_960_v3_)
                        nw146_[int(15)] = (self).f24
                        nw146_[int(16)] = (self).f24
                        nw146_[int(17)] = 597
                        nw146_[int(18)] = (self).f24
                        nw146_[int(19)] = (self).f24
                        nw146_[int(20)] = (self).f24
                        nw146_[int(21)] = d_964_cf37_
                        nw146_[int(22)] = -506
                        nw146_[int(23)] = (self).f24
                        nw146_[int(24)] = (self).f24
                        nw146_[int(25)] = d_964_cf37_
                        nw146_[int(26)] = (self).f24
                        d_974_v13_ = nw146_
                        d_977_v14_: _dafny.Seq
                        d_977_v14_ = _dafny.SeqWithoutIsStrInference([d_974_v13_, (d_974_v13_ if p0 else d_974_v13_), d_974_v13_])
                        index119_ = default__.safeIndex(329, (d_974_v13_).length(0))
                        (d_974_v13_)[index119_] = len(_dafny.SeqWithoutIsStrInference([(0) - ((0) - (d_964_cf37_))]))
                        index120_ = default__.safeIndex(583, (d_974_v13_).length(0))
                        (d_974_v13_)[index120_] = d_964_cf37_
                        d_978_v15_: _dafny.Map
                        d_978_v15_ = _dafny.Map({d_959_v2_: (self).f24})
                        d_979_v16_: T0
                        nw147_ = C2()
                        nw147_.ctor__(p0, (self).f24, _dafny.SeqWithoutIsStrInference([d_978_v15_]))
                        d_979_v16_ = nw147_
                        d_980_v17_: _dafny.Map
                        d_980_v17_ = _dafny.Map({d_979_v16_: d_974_v13_})
                        d_981_v18_: C0
                        nw148_ = C0()
                        nw148_.ctor__(p0)
                        d_981_v18_ = nw148_
                        d_982_v19_: _dafny.Map
                        d_982_v19_ = _dafny.Map({d_981_v18_: (d_981_v18_).f34})
                        d_983_v20_: _dafny.Map
                        d_983_v20_ = _dafny.Map({(d_979_v16_).f24: True})
                        d_984_v21_: _dafny.Set
                        d_984_v21_ = _dafny.Set({d_966_v7_, d_966_v7_, d_966_v7_})
                        d_985_v22_: _dafny.Set
                        d_985_v22_ = _dafny.Set({not((d_981_v18_).f34)})
                        index121_ = default__.safeIndex(329, (d_974_v13_).length(0))
                        index122_ = default__.safeIndex(583, (d_974_v13_).length(0))
                        rhs150_ = ((d_977_v14_) + (_dafny.SeqWithoutIsStrInference([d_974_v13_]))) + (d_977_v14_)
                        rhs151_ = ((self).f24) - (len(d_980_v17_))
                        rhs152_ = ((default__.fm4(p0, len((d_982_v19_) | (d_982_v19_)), d_959_v2_, p0, globalState)).set(default__.safeIndex((d_966_v7_)[default__.safeIndex((self).f24, len(d_966_v7_))], len(default__.fm4(p0, len((d_982_v19_) | (d_982_v19_)), d_959_v2_, p0, globalState))), d_959_v2_)).set(default__.safeIndex(len(d_960_v3_), len((default__.fm4(p0, len((d_982_v19_) | (d_982_v19_)), d_959_v2_, p0, globalState)).set(default__.safeIndex((d_966_v7_)[default__.safeIndex((self).f24, len(d_966_v7_))], len(default__.fm4(p0, len((d_982_v19_) | (d_982_v19_)), d_959_v2_, p0, globalState))), d_959_v2_))), d_959_v2_)
                        rhs153_ = not ((d_981_v18_).f34) or (p0)
                        rhs154_ = default__.safeModuloInt((d_966_v7_)[default__.safeIndex(len(_dafny.Set({len(d_983_v20_), (self).f24, (self).f24, len(d_984_v21_)})), len(d_966_v7_))], len(d_985_v22_))
                        lhs115_ = d_974_v13_
                        lhs116_ = default__.safeIndex(329, (d_974_v13_).length(0))
                        lhs117_ = globalState
                        lhs118_ = d_974_v13_
                        lhs119_ = default__.safeIndex(583, (d_974_v13_).length(0))
                        d_977_v14_ = rhs150_
                        lhs115_[lhs116_] = rhs151_
                        lhs117_.f17 = rhs152_
                        r1 = rhs153_
                        lhs118_[lhs119_] = rhs154_
                        index123_ = default__.safeIndex(329, (d_974_v13_).length(0))
                        (d_974_v13_)[index123_] = (d_979_v16_).f24
                        (globalState).f1 = ((D0_DC2(not((d_981_v18_).f34), (self).f24, d_959_v2_)).cf5) * (d_964_cf37_)
                        (globalState).f2 = not((d_981_v18_).f34)
                    elif source11_.is_DC25:
                        d_986___mcc_h1_ = source11_.cf38
                        d_987___mcc_h2_ = source11_.cf39
                        d_988___mcc_h3_ = source11_.cf40
                        d_989_cf40_ = d_988___mcc_h3_
                        d_990_cf39_ = d_987___mcc_h2_
                        d_991_cf38_ = d_986___mcc_h1_
                        d_992_v23_: _dafny.Map
                        d_992_v23_ = _dafny.Map({d_991_cf38_: d_989_cf40_})
                        d_993_v24_: _dafny.Map
                        d_993_v24_ = _dafny.Map({d_990_cf39_: len((_dafny.SeqWithoutIsStrInference([d_989_cf40_, d_989_cf40_, not(not(p0))])).set(default__.safeIndex(d_991_cf38_, len(_dafny.SeqWithoutIsStrInference([d_989_cf40_, d_989_cf40_, not(not(p0))]))), ((d_992_v23_)[(self).f24] if ((self).f24) in (d_992_v23_) else d_989_cf40_)))})
                        d_994_v25_: _dafny.Map
                        d_994_v25_ = _dafny.Map({d_991_cf38_: 679})
                        d_993_v24_ = (d_993_v24_).set((d_990_cf39_) and (p0), default__.safeDivisionInt((self).f24, len(d_994_v25_)))
                        (globalState).f1 = default__.safeDivisionInt(d_991_cf38_, default__.safeDivisionInt(d_991_cf38_, (self).f24))
                        rhs155_ = p0
                        rhs156_ = d_989_cf40_
                        rhs157_ = (d_959_v2_) not in (d_960_v3_)
                        rhs158_ = (_dafny.CodePoint('b') if True else d_959_v2_)
                        rhs159_ = (self).f24
                        lhs120_ = globalState
                        lhs121_ = globalState
                        lhs122_ = globalState
                        lhs120_.f2 = rhs155_
                        lhs121_.f2 = rhs156_
                        r0 = rhs157_
                        d_959_v2_ = rhs158_
                        lhs122_.f1 = rhs159_
                        d_995_v26_: _dafny.Array
                        nw149_ = _dafny.Array(None, 19)
                        nw149_[int(0)] = True
                        nw149_[int(1)] = True
                        nw149_[int(2)] = p0
                        nw149_[int(3)] = d_990_cf39_
                        nw149_[int(4)] = not(p0)
                        nw149_[int(5)] = p0
                        nw149_[int(6)] = p0
                        nw149_[int(7)] = d_990_cf39_
                        nw149_[int(8)] = False
                        nw149_[int(9)] = d_989_cf40_
                        nw149_[int(10)] = not(p0)
                        nw149_[int(11)] = d_989_cf40_
                        nw149_[int(12)] = p0
                        nw149_[int(13)] = d_989_cf40_
                        nw149_[int(14)] = d_989_cf40_
                        nw149_[int(15)] = d_990_cf39_
                        nw149_[int(16)] = not(not(True))
                        nw149_[int(17)] = p0
                        nw149_[int(18)] = True
                        d_995_v26_ = nw149_
                        d_996_v27_: _dafny.Seq
                        d_996_v27_ = _dafny.SeqWithoutIsStrInference([d_995_v26_, d_995_v26_])
                        d_997_v28_: _dafny.Map
                        d_997_v28_ = _dafny.Map({(d_996_v27_) + (d_996_v27_): len(d_993_v24_)})
                        d_997_v28_ = (d_997_v28_).set(d_996_v27_, (0) - ((0) - (len(d_993_v24_))))
                    elif source11_.is_DC26:
                        d_998___mcc_h4_ = source11_.cf41
                        d_999___mcc_h5_ = source11_.cf42
                        d_1000_cf42_ = d_999___mcc_h5_
                        d_1001_cf41_ = d_998___mcc_h4_
                        d_1002_v29_: _dafny.Map
                        d_1002_v29_ = _dafny.Map({(self).f24: d_959_v2_})
                        d_1002_v29_ = (d_1002_v29_).set((self).f24, d_959_v2_)
                        d_1000_cf42_ = d_960_v3_
                        d_1003_v30_: _dafny.Map
                        d_1003_v30_ = _dafny.Map({(self).f24: (self).f24})
                        d_1004_v31_: _dafny.Map
                        d_1004_v31_ = _dafny.Map({p0: (self).f24})
                        d_1005_v32_: _dafny.Map
                        d_1005_v32_ = _dafny.Map({p0: p0})
                        d_1006_v33_: _dafny.Seq
                        d_1006_v33_ = _dafny.SeqWithoutIsStrInference([len(d_960_v3_), (self).f24, (self).f24, len(_dafny.Set({p0, p0})), len(d_1005_v32_)])
                        d_1007_v34_: D8
                        d_1007_v34_ = D8_DC25((self).f24, p0, p0)
                        d_1008_v35_: _dafny.Map
                        d_1008_v35_ = _dafny.Map({(default__.fm14(d_1003_v30_, ((d_1004_v31_)[p0] if (p0) in (d_1004_v31_) else len(d_1006_v33_)), globalState)).cardinality: (d_1007_v34_).cf39})
                        d_1009_v36_: _dafny.Array
                        nw150_ = _dafny.Array(False, 29)
                        d_1009_v36_ = nw150_
                        d_1010_v37_: D1
                        d_1010_v37_ = D1_DC5(p0, p0, not(p0))
                        index124_ = default__.safeIndex(800, (d_1009_v36_).length(0))
                        (d_1009_v36_)[index124_] = (d_1010_v37_).cf14
                        d_1011_v38_: _dafny.Seq
                        d_1011_v38_ = _dafny.SeqWithoutIsStrInference([d_1009_v36_, d_1009_v36_, d_1009_v36_])
                        index125_ = default__.safeIndex(800, (d_1009_v36_).length(0))
                        rhs160_ = d_1008_v35_
                        rhs161_ = not (True) or (p0)
                        rhs162_ = d_1011_v38_
                        lhs123_ = d_1009_v36_
                        lhs124_ = default__.safeIndex(800, (d_1009_v36_).length(0))
                        d_1008_v35_ = rhs160_
                        lhs123_[lhs124_] = rhs161_
                        d_1011_v38_ = rhs162_
                        d_1012_v39_: _dafny.Map
                        d_1012_v39_ = _dafny.Map({D4_DC13(len(d_1006_v33_), d_1004_v31_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "de"))): len(d_1006_v33_)})
                        d_1012_v39_ = (d_1012_v39_).set(default__.fm25(globalState), (self).fm8(globalState))
                    elif True:
                        d_1013___mcc_h6_ = source11_.cf36
                        d_1014_cf36_ = d_1013___mcc_h6_
                        d_1015_v40_: _dafny.Array
                        def lambda41_(d_1016_i2_):
                            return _dafny.MultiSet([(self).f24, (self).f24])

                        init25_ = lambda41_
                        nw151_ = _dafny.Array(None, 25)
                        for i0_25_ in range(nw151_.length(0)):
                            nw151_[i0_25_] = init25_(i0_25_)
                        d_1015_v40_ = nw151_
                        d_1017_v41_: _dafny.Array
                        nw152_ = _dafny.Array(None, 9)
                        nw152_[int(0)] = d_1015_v40_
                        nw152_[int(1)] = d_1015_v40_
                        nw152_[int(2)] = d_1015_v40_
                        nw152_[int(3)] = d_1015_v40_
                        nw152_[int(4)] = d_1015_v40_
                        nw152_[int(5)] = d_1015_v40_
                        nw152_[int(6)] = d_1015_v40_
                        nw152_[int(7)] = d_1015_v40_
                        nw152_[int(8)] = d_1015_v40_
                        d_1017_v41_ = nw152_
                        index126_ = default__.safeIndex(436, (d_1017_v41_).length(0))
                        (d_1017_v41_)[index126_] = (d_1015_v40_ if not(p0) else d_1015_v40_)
                        d_1018_v42_: _dafny.Array
                        nw153_ = _dafny.Array(False, 27)
                        d_1018_v42_ = nw153_
                        index127_ = default__.safeIndex(981, (d_1018_v42_).length(0))
                        (d_1018_v42_)[index127_] = p0
                        d_1019_v43_: D1
                        d_1019_v43_ = D1_DC5(p0, not(p0), not(p0))
                        d_1020_v44_: _dafny.Map
                        d_1020_v44_ = _dafny.Map({d_960_v3_: d_1019_v43_})
                        d_1021_v45_: _dafny.Array
                        nw154_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                        d_1021_v45_ = nw154_
                        index128_ = default__.safeIndex(89, (d_1021_v45_).length(0))
                        (d_1021_v45_)[index128_] = d_960_v3_
                        d_1022_v46_: _dafny.Seq
                        d_1022_v46_ = _dafny.SeqWithoutIsStrInference([d_1015_v40_])
                        index129_ = default__.safeIndex(436, (d_1017_v41_).length(0))
                        index130_ = default__.safeIndex(981, (d_1018_v42_).length(0))
                        index131_ = default__.safeIndex(89, (d_1021_v45_).length(0))
                        rhs163_ = (d_1022_v46_)[default__.safeIndex((self).f24, len(d_1022_v46_))]
                        rhs164_ = (d_960_v3_) + (d_960_v3_)
                        rhs165_ = (((_dafny.MultiSet([p0])).set(p0, default__.abs((self).f24))).cardinality) < ((self).f24)
                        rhs166_ = (default__.fm34(p0, globalState)) | (d_1020_v44_)
                        rhs167_ = (d_960_v3_) + (d_960_v3_)
                        lhs125_ = d_1017_v41_
                        lhs126_ = default__.safeIndex(436, (d_1017_v41_).length(0))
                        lhs127_ = globalState
                        lhs128_ = d_1018_v42_
                        lhs129_ = default__.safeIndex(981, (d_1018_v42_).length(0))
                        lhs130_ = d_1021_v45_
                        lhs131_ = default__.safeIndex(89, (d_1021_v45_).length(0))
                        lhs125_[lhs126_] = rhs163_
                        lhs127_.f17 = rhs164_
                        lhs128_[lhs129_] = rhs165_
                        d_1020_v44_ = rhs166_
                        lhs130_[lhs131_] = rhs167_
                        d_1023_v47_: _dafny.Set
                        d_1023_v47_ = _dafny.Set({(self).f24})
                        d_1024_v48_: _dafny.Set
                        d_1024_v48_ = _dafny.Set({len(_dafny.Map({(self).f24: len(d_1023_v47_)}))})
                        d_1025_v49_: _dafny.Map
                        d_1025_v49_ = _dafny.Map({(d_1023_v47_ if p0 else d_1023_v47_): 821})
                        rhs168_ = d_1025_v49_
                        rhs169_ = (self).f24
                        lhs132_ = globalState
                        d_1025_v49_ = rhs168_
                        lhs132_.f9 = rhs169_
                        index132_ = default__.safeIndex(981, (d_1018_v42_).length(0))
                        (d_1018_v42_)[index132_] = (d_1018_v42_)[default__.safeIndex(981, (d_1018_v42_).length(0))]
                        d_1026_v50_: _dafny.Map
                        d_1026_v50_ = _dafny.Map({(d_1018_v42_)[default__.safeIndex(981, (d_1018_v42_).length(0))]: (d_1018_v42_)[default__.safeIndex(981, (d_1018_v42_).length(0))]})
                        d_1027_v51_: _dafny.MultiSet
                        d_1027_v51_ = _dafny.MultiSet([d_1026_v50_])
                        d_1028_v52_: _dafny.MultiSet
                        d_1028_v52_ = _dafny.MultiSet([default__.fm0(default__.fm0((d_1027_v51_).cardinality, (self).f24, globalState), len(default__.fm17(157, -613, (self).f24, (d_1018_v42_)[default__.safeIndex(981, (d_1018_v42_).length(0))], globalState)), globalState)])
                        d_1029_v53_: _dafny.Map
                        d_1029_v53_ = _dafny.Map({d_959_v2_: (self).f24})
                        d_1030_v54_: _dafny.Map
                        d_1030_v54_ = _dafny.Map({(d_1018_v42_)[default__.safeIndex(981, (d_1018_v42_).length(0))]: (self).f25})
                        d_1031_v55_: C1
                        nw155_ = C1()
                        nw155_.ctor__(_dafny.SeqWithoutIsStrInference([d_1028_v52_]), (self).f24, default__.safeDivisionInt((self).f24, (self).f24), (_dafny.SeqWithoutIsStrInference([d_1029_v53_]) if p0 else ((d_1030_v54_)[p0] if (p0) in (d_1030_v54_) else (self).f25)))
                        d_1031_v55_ = nw155_
                    d_1032_v56_: _dafny.Map
                    d_1032_v56_ = _dafny.Map({(0) - ((0) - ((self).f24)): p0})
                    d_1033_v57_: C1
                    nw156_ = C1()
                    nw156_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(d_1032_v56_)])]), (self).f24, 939, ((self).f25) + ((self).f25))
                    d_1033_v57_ = nw156_
                    (globalState).f21 = d_956_v0_
                    d_1034_v58_: _dafny.Map
                    d_1034_v58_ = _dafny.Map({default__.fm2((self).f24, (0) - ((self).f24), d_959_v2_, globalState): (d_1033_v57_.f33) != (len(_dafny.Map({d_956_v0_: p0})))})
                    d_1034_v58_ = (d_1034_v58_).set(p0, p0)
                    pass
            pass
        d_1043_v59_: _dafny.Array
        def lambda43_(d_1044_p0_):
            def lambda44_(d_1045_i3_):
                return (_dafny.Map({d_1044_p0_: d_1044_p0_})) | (_dafny.Map({d_1044_p0_: d_1044_p0_}))

            return lambda44_

        init26_ = lambda43_(p0)
        nw157_ = _dafny.Array(None, 8)
        for i0_26_ in range(nw157_.length(0)):
            nw157_[i0_26_] = init26_(i0_26_)
        d_1043_v59_ = nw157_
        index133_ = default__.safeIndex(473, (d_1043_v59_).length(0))
        (d_1043_v59_)[index133_] = _dafny.Map({p0: p0})
        d_1046_v60_: _dafny.Set
        d_1046_v60_ = _dafny.Set({_dafny.CodePoint('i')})
        d_1047_v61_: _dafny.Array
        nw158_ = _dafny.Array(False, 6)
        d_1047_v61_ = nw158_
        index134_ = default__.safeIndex(247, (d_1047_v61_).length(0))
        (d_1047_v61_)[index134_] = not (p0) or (p0)
        d_1048_v62_: str
        d_1048_v62_ = _dafny.CodePoint('j')
        d_1049_v63_: _dafny.Map
        d_1049_v63_ = _dafny.Map({not(not(p0)): not(not(p0))})
        d_1050_v64_: _dafny.MultiSet
        d_1050_v64_ = _dafny.MultiSet([(self).f24, (self).f24])
        index135_ = default__.safeIndex(473, (d_1043_v59_).length(0))
        index136_ = default__.safeIndex(247, (d_1047_v61_).length(0))
        rhs170_ = d_1049_v63_
        rhs171_ = d_1046_v60_
        rhs172_ = ((0) - ((self).f24)) >= ((self).f24)
        rhs173_ = ((_dafny.MultiSet([(self).f24])).intersection((d_1050_v64_).intersection(d_1050_v64_))).cardinality
        rhs174_ = d_1048_v62_
        lhs133_ = d_1043_v59_
        lhs134_ = default__.safeIndex(473, (d_1043_v59_).length(0))
        lhs135_ = d_1047_v61_
        lhs136_ = default__.safeIndex(247, (d_1047_v61_).length(0))
        lhs137_ = globalState
        lhs133_[lhs134_] = rhs170_
        d_1046_v60_ = rhs171_
        lhs135_[lhs136_] = rhs172_
        lhs137_.f9 = rhs173_
        d_1048_v62_ = rhs174_
        d_1051_v65_: _dafny.Map
        d_1051_v65_ = _dafny.Map({(self).f24: default__.fm15((self).f24, p0, (self).f24, (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))], globalState)})
        d_1051_v65_ = d_1051_v65_
        if not(((d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]) == ((d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))])):
            d_1052_v66_: D3
            d_1052_v66_ = D3_DC11(((self).f24) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))), p0, True)
            d_1053_v67_: _dafny.Seq
            d_1053_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fidp"))
            rhs175_ = d_1047_v61_
            rhs176_ = d_1052_v66_
            rhs177_ = (_dafny.SeqWithoutIsStrInference([d_1048_v62_ for d_1054_i4_ in range(default__.abs(997))])) + (d_1053_v67_)
            lhs138_ = globalState
            d_1047_v61_ = rhs175_
            d_1052_v66_ = rhs176_
            lhs138_.f17 = rhs177_
            index137_ = default__.safeIndex(247, (d_1047_v61_).length(0))
            (d_1047_v61_)[index137_] = not (not ((d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]) or (not(p0))) or (p0)
            (globalState).f2 = (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]
            d_1055_v68_: _dafny.Seq
            d_1055_v68_ = _dafny.SeqWithoutIsStrInference([(self).f24])
            d_1056_v69_: D4
            d_1056_v69_ = D4_DC14(D4_DC12(d_1055_v68_))
            d_1057_v70_: _dafny.Map
            d_1057_v70_ = _dafny.Map({d_1056_v69_: d_1053_v67_})
            d_1058_v71_: D7
            d_1058_v71_ = D7_DC20(d_1057_v70_)
            d_1059_v72_: D7
            d_1059_v72_ = D7_DC22(d_1058_v71_)
            d_1059_v72_ = d_1059_v72_
            d_1060_v73_: _dafny.Map
            d_1060_v73_ = _dafny.Map({(self).f24: (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]})
            d_1061_v74_: _dafny.Map
            d_1061_v74_ = _dafny.Map({not(p0): len(d_1053_v67_)})
            d_1062_v75_: D4
            d_1062_v75_ = D4_DC13((self).f24, d_1061_v74_, d_1053_v67_)
            d_1063_v76_: _dafny.Set
            d_1063_v76_ = _dafny.Set({862})
            d_1064_v77_: _dafny.MultiSet
            d_1064_v77_ = _dafny.MultiSet([p0, p0, (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]])
            index138_ = default__.safeIndex(247, (d_1047_v61_).length(0))
            rhs178_ = ((d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]) and (not((len(d_1060_v73_)) <= ((self).f24)))
            rhs179_ = (d_1062_v75_).cf22
            rhs180_ = (0) - (default__.fm0(len(d_1063_v76_), ((d_1064_v77_)[(d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]] if ((d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]) in (d_1064_v77_) else (self).f24), globalState))
            rhs181_ = (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]
            rhs182_ = ((self).f24) - (((self).f24) + ((self).f24))
            lhs139_ = d_1047_v61_
            lhs140_ = default__.safeIndex(247, (d_1047_v61_).length(0))
            lhs141_ = globalState
            lhs142_ = globalState
            lhs143_ = globalState
            lhs139_[lhs140_] = rhs178_
            lhs141_.f14 = rhs179_
            lhs142_.f9 = rhs180_
            r0 = rhs181_
            lhs143_.f1 = rhs182_
        elif True:
            (globalState).f2 = ((p0 if p0 else (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))])) and ((p0 if p0 else default__.fm2(len(d_956_v0_), (self).f24, _dafny.CodePoint('w'), globalState)))
            d_1065_v78_: _dafny.Seq
            d_1065_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omqcl"))
            if (p0 if not((d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]) else (d_1065_v78_) < (d_1065_v78_)):
                rhs183_ = (len((d_1065_v78_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqal"))))) + ((self).f24)
                rhs184_ = p0
                lhs144_ = globalState
                lhs145_ = globalState
                lhs144_.f1 = rhs183_
                lhs145_.f2 = rhs184_
                d_1066_v79_: _dafny.Map
                d_1066_v79_ = _dafny.Map({34: not(True)})
                d_1067_v80_: _dafny.Set
                d_1067_v80_ = _dafny.Set({default__.safeDivisionInt(len(d_1066_v79_), (self).f24)})
                d_1067_v80_ = d_1067_v80_
                d_1068_v81_: _dafny.Array
                nw159_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_1068_v81_ = nw159_
                d_1069_v82_: _dafny.Map
                d_1069_v82_ = _dafny.Map({d_1049_v63_: (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]})
                rhs185_ = d_1068_v81_
                rhs186_ = (0) - ((len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]: (self).f24}) for d_1070_i5_ in range(default__.abs(294))])) if not(((self).f24) != ((0) - ((self).f24))) else len((d_1069_v82_) | (d_1069_v82_))))
                lhs146_ = globalState
                d_1068_v81_ = rhs185_
                lhs146_.f1 = rhs186_
                d_1071_v83_: _dafny.Map
                d_1071_v83_ = _dafny.Map({(self).f24: d_1047_v61_})
                d_1071_v83_ = (d_1071_v83_).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnv"))), d_1047_v61_)
                index139_ = default__.safeIndex(247, (d_1047_v61_).length(0))
                (d_1047_v61_)[index139_] = (d_956_v0_)[default__.safeIndex(default__.safeDivisionInt((self).f24, (self).f24), len(d_956_v0_))]
            elif True:
                d_1072_v84_: _dafny.Map
                d_1072_v84_ = _dafny.Map({(d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]: (self).f24})
                d_1072_v84_ = d_1072_v84_
                (globalState).f18 = (self).f24
                d_1073_v85_: _dafny.Array
                nw160_ = _dafny.Array(int(0), 5)
                d_1073_v85_ = nw160_
                index140_ = default__.safeIndex(833, (d_1073_v85_).length(0))
                (d_1073_v85_)[index140_] = ((0) - ((self).f24) if (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))] else (self).f24)
                index141_ = default__.safeIndex(833, (d_1073_v85_).length(0))
                (d_1073_v85_)[index141_] = (self).f24
                (globalState).f2 = p0
                (globalState).f14 = (d_1073_v85_)[default__.safeIndex(833, (d_1073_v85_).length(0))]
            (globalState).f9 = default__.safeDivisionInt((0) - (default__.safeDivisionInt((self).f24, (self).f24)), (self).f24)
            if p0:
                d_1074_v86_: _dafny.Map
                d_1074_v86_ = _dafny.Map({d_1048_v62_: p0})
                index142_ = default__.safeIndex(247, (d_1047_v61_).length(0))
                (d_1047_v61_)[index142_] = not(((d_1074_v86_)[d_1048_v62_] if (d_1048_v62_) in (d_1074_v86_) else p0))
                (globalState).f9 = (self).f24
                d_1075_v87_: _dafny.Array
                def lambda45_(d_1076_v0_):
                    def lambda46_(d_1077_i6_):
                        return d_1076_v0_

                    return lambda46_

                init27_ = lambda45_(d_956_v0_)
                nw161_ = _dafny.Array(None, 16)
                for i0_27_ in range(nw161_.length(0)):
                    nw161_[i0_27_] = init27_(i0_27_)
                d_1075_v87_ = nw161_
                index143_ = default__.safeIndex(750, (d_1075_v87_).length(0))
                (d_1075_v87_)[index143_] = d_956_v0_
                index144_ = default__.safeIndex(750, (d_1075_v87_).length(0))
                (d_1075_v87_)[index144_] = (d_956_v0_).set(default__.safeIndex((self).f24, len(d_956_v0_)), False)
                d_1078_v88_: _dafny.Map
                d_1078_v88_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))], (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]]): d_1047_v61_})
                d_1078_v88_ = d_1078_v88_
                (globalState).f9 = (self).f24
            elif True:
                d_1079_v89_: _dafny.Map
                d_1079_v89_ = _dafny.Map({(self).f24: d_1048_v62_})
                d_1048_v62_ = ((d_1079_v89_)[len((d_956_v0_) + (d_956_v0_))] if (len((d_956_v0_) + (d_956_v0_))) in (d_1079_v89_) else d_1048_v62_)
                r1 = p0
                d_1080_v90_: _dafny.Map
                d_1080_v90_ = _dafny.Map({default__.safeDivisionInt((self).f24, (self).fm8(globalState)): ((d_1049_v63_)[p0] if (p0) in (d_1049_v63_) else p0)})
                d_1080_v90_ = (d_1080_v90_).set((self).f24, (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))])
                d_1081_v91_: C2
                nw162_ = C2()
                nw162_.ctor__(p0, default__.safeModuloInt((self).f24, (self).f24), (self).f25)
                d_1081_v91_ = nw162_
                d_1082_v92_: _dafny.Map
                d_1082_v92_ = _dafny.Map({d_1048_v62_: not(True)})
                r0 = ((d_1082_v92_)[d_1048_v62_] if (d_1048_v62_) in (d_1082_v92_) else ((self).f24) > (123))
            d_1083_v93_: _dafny.Set
            d_1083_v93_ = _dafny.Set({(d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))], (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]})
            d_1084_v94_: D0
            d_1084_v94_ = D0_DC1(len(d_1083_v93_), (self).f24, (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))])
            d_1085_v95_: C0
            nw163_ = C0()
            nw163_.ctor__((d_1084_v94_).cf3)
            d_1085_v95_ = nw163_
        d_1047_v61_ = d_1047_v61_
        d_1086_i7_: int
        d_1086_i7_ = 0
        with _dafny.label("3"):
            while ((self).f24) != ((self).f24):
                with _dafny.c_label("3"):
                    if (d_1086_i7_) >= (100):
                        raise _dafny.Break("3")
                    d_1086_i7_ = (d_1086_i7_) + (1)
                    (globalState).f1 = (((self).f24) + ((self).f24)) + ((self).f24)
                    d_1087_v96_: _dafny.Set
                    d_1087_v96_ = _dafny.Set({True})
                    d_1087_v96_ = d_1087_v96_
                    d_1088_v97_: _dafny.Map
                    d_1088_v97_ = _dafny.Map({(self).f24: (d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]})
                    d_1088_v97_ = d_1088_v97_
                    d_1089_v98_: _dafny.MultiSet
                    d_1089_v98_ = _dafny.MultiSet([(d_1047_v61_)[default__.safeIndex(247, (d_1047_v61_).length(0))]])
                    r1 = (_dafny.MultiSet(d_956_v0_)).ispropersubset(d_1089_v98_)
                    pass
            pass
        d_1090_v99_: _dafny.Seq
        d_1090_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lq"))
        d_1091_v100_: _dafny.Set
        d_1091_v100_ = _dafny.Set({len(d_1090_v99_)})
        r0 = (((D6_DC16(d_1091_v100_)).cf27) | (d_1091_v100_)).isdisjoint((_dafny.Set({17})) - (_dafny.Set({(self).f24})))
        r1 = not((p0) == (p0))
        r2 = (self).f24
        return r0, r1, r2

    def m6(self, globalState):
        d_1092_v0_: _dafny.Map
        d_1092_v0_ = _dafny.Map({((self).f24) + ((self).f24): True})
        d_1093_v1_: bool
        d_1093_v1_ = False
        d_1092_v0_ = (d_1092_v0_).set((self).f24, d_1093_v1_)
        (globalState).f2 = not((((self).f24) - ((self).f24)) == ((self).f24))
        d_1094_v2_: D1
        d_1094_v2_ = D1_DC5(False, d_1093_v1_, d_1093_v1_)
        d_1095_v3_: D6
        d_1095_v3_ = D6_DC17(d_1093_v1_)
        d_1096_v4_: D6
        d_1096_v4_ = D6_DC19(d_1095_v3_)
        d_1097_v5_: _dafny.Seq
        d_1097_v5_ = _dafny.SeqWithoutIsStrInference([d_1093_v1_, d_1093_v1_])
        d_1098_v6_: _dafny.Array
        nw164_ = _dafny.Array(None, 12)
        nw164_[int(0)] = d_1093_v1_
        nw164_[int(1)] = d_1093_v1_
        nw164_[int(2)] = d_1093_v1_
        nw164_[int(3)] = d_1093_v1_
        nw164_[int(4)] = not(d_1093_v1_)
        nw164_[int(5)] = (d_1097_v5_)[default__.safeIndex((self).f24, len(d_1097_v5_))]
        nw164_[int(6)] = d_1093_v1_
        nw164_[int(7)] = d_1093_v1_
        nw164_[int(8)] = d_1093_v1_
        nw164_[int(9)] = d_1093_v1_
        nw164_[int(10)] = d_1093_v1_
        nw164_[int(11)] = d_1093_v1_
        d_1098_v6_ = nw164_
        d_1099_v7_: D7
        d_1099_v7_ = D7_DC21(d_1094_v2_, d_1096_v4_, d_1098_v6_)
        d_1100_v8_: str
        d_1100_v8_ = _dafny.CodePoint('k')
        pat_let_tv7_ = d_1100_v8_
        pat_let_tv8_ = globalState
        def iife54_(_pat_let12_0):
            def iife55_(d_1101_dt__update__tmp_h0_):
                def iife56_(_pat_let13_0):
                    def iife57_(d_1102_dt__update_hcf33_h0_):
                        return D7_DC21((d_1101_dt__update__tmp_h0_).cf32, d_1102_dt__update_hcf33_h0_, (d_1101_dt__update__tmp_h0_).cf34)
                    return iife57_(_pat_let13_0)
                return iife56_(D6_DC19(D6_DC18(default__.fm2((self).f24, (self).f24, pat_let_tv7_, pat_let_tv8_))))
            return iife55_(_pat_let12_0)
        source13_ = iife54_(d_1099_v7_)
        if source13_.is_DC21:
            d_1103___mcc_h0_ = source13_.cf32
            d_1104___mcc_h1_ = source13_.cf33
            d_1105___mcc_h2_ = source13_.cf34
            d_1106_cf34_ = d_1105___mcc_h2_
            d_1107_cf33_ = d_1104___mcc_h1_
            d_1108_cf32_ = d_1103___mcc_h0_
            (globalState).f2 = (d_1093_v1_) and (d_1093_v1_)
            (globalState).f2 = d_1093_v1_
            d_1109_v9_: _dafny.Seq
            d_1109_v9_ = _dafny.SeqWithoutIsStrInference([d_1100_v8_, d_1100_v8_])
            d_1110_v10_: _dafny.Set
            d_1110_v10_ = _dafny.Set({(0) - (len(d_1109_v9_)), (self).f24, (self).f24})
            d_1111_v11_: D6
            d_1111_v11_ = D6_DC18(d_1093_v1_)
            d_1112_v12_: _dafny.Seq
            d_1112_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1100_v8_ for d_1113_i0_ in range(default__.abs(291))])])
            d_1114_v13_: _dafny.MultiSet
            d_1114_v13_ = _dafny.MultiSet([d_1093_v1_])
            def iife58_(_pat_let14_0):
                def iife59_(d_1115_dt__update__tmp_h1_):
                    def iife60_(_pat_let15_0):
                        def iife61_(d_1116_dt__update_hcf29_h0_):
                            return D6_DC18(d_1116_dt__update_hcf29_h0_)
                        return iife61_(_pat_let15_0)
                    return iife60_(True)
                return iife59_(_pat_let14_0)
            rhs187_ = len((d_1112_v12_).set(default__.safeIndex((self).f24, len(d_1112_v12_)), d_1109_v9_))
            rhs188_ = d_1110_v10_
            rhs189_ = d_1097_v5_
            rhs190_ = iife58_(d_1111_v11_)
            rhs191_ = default__.fm0(((self).f24) - (((d_1114_v13_).set(d_1093_v1_, default__.abs((self).f24))).cardinality), 600, globalState)
            lhs147_ = globalState
            lhs148_ = globalState
            lhs149_ = globalState
            lhs147_.f18 = rhs187_
            d_1110_v10_ = rhs188_
            lhs148_.f20 = rhs189_
            d_1111_v11_ = rhs190_
            lhs149_.f18 = rhs191_
            d_1117_v14_: _dafny.Array
            def lambda47_(d_1118_i1_):
                return (d_1118_i1_) + ((self).f24)

            init28_ = lambda47_
            nw165_ = _dafny.Array(None, 17)
            for i0_28_ in range(nw165_.length(0)):
                nw165_[i0_28_] = init28_(i0_28_)
            d_1117_v14_ = nw165_
            d_1119_v15_: _dafny.Map
            d_1119_v15_ = _dafny.Map({d_1117_v14_: d_1093_v1_})
            d_1120_v16_: _dafny.Seq
            d_1120_v16_ = _dafny.SeqWithoutIsStrInference([d_1110_v10_])
            d_1121_v17_: _dafny.Map
            d_1121_v17_ = _dafny.Map({d_1109_v9_: d_1120_v16_})
            d_1122_v18_: _dafny.Seq
            d_1122_v18_ = _dafny.SeqWithoutIsStrInference([57])
            rhs192_ = (((self).f24) - ((self).f24)) + ((self).fm8(globalState))
            rhs193_ = default__.fm3(globalState)
            rhs194_ = d_1119_v15_
            rhs195_ = ((d_1121_v17_)[d_1109_v9_] if (d_1109_v9_) in (d_1121_v17_) else default__.fm35(not(d_1093_v1_), d_1093_v1_, globalState))
            rhs196_ = (_dafny.MultiSet(d_1122_v18_)).issubset(_dafny.MultiSet([len(d_1092_v0_), (self).f24, (self).f24, 709]))
            lhs150_ = globalState
            lhs151_ = globalState
            lhs150_.f1 = rhs192_
            d_1100_v8_ = rhs193_
            d_1119_v15_ = rhs194_
            d_1120_v16_ = rhs195_
            lhs151_.f2 = rhs196_
        elif source13_.is_DC20:
            d_1123___mcc_h3_ = source13_.cf31
            d_1124_cf31_ = d_1123___mcc_h3_
            d_1125_v19_: _dafny.Array
            nw166_ = _dafny.Array(int(0), 13)
            d_1125_v19_ = nw166_
            d_1126_v20_: _dafny.MultiSet
            d_1126_v20_ = _dafny.MultiSet([d_1093_v1_, d_1093_v1_, d_1093_v1_])
            index145_ = default__.safeIndex(428, (d_1125_v19_).length(0))
            (d_1125_v19_)[index145_] = ((self).f24) - ((d_1126_v20_).cardinality)
            d_1127_v21_: _dafny.Seq
            d_1127_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsj"))
            index146_ = default__.safeIndex(428, (d_1125_v19_).length(0))
            rhs197_ = d_1127_v21_
            rhs198_ = (self).f24
            lhs152_ = globalState
            lhs153_ = d_1125_v19_
            lhs154_ = default__.safeIndex(428, (d_1125_v19_).length(0))
            lhs152_.f17 = rhs197_
            lhs153_[lhs154_] = rhs198_
            d_1093_v1_ = d_1093_v1_
            (globalState).f2 = d_1093_v1_
            (globalState).f9 = ((d_1125_v19_)[default__.safeIndex(428, (d_1125_v19_).length(0))]) + ((d_1125_v19_)[default__.safeIndex(428, (d_1125_v19_).length(0))])
        elif True:
            d_1128___mcc_h4_ = source13_.cf35
            d_1129_cf35_ = d_1128___mcc_h4_
            d_1130_v22_: _dafny.Map
            d_1130_v22_ = _dafny.Map({d_1100_v8_: (self).f24})
            d_1131_v23_: C2
            nw167_ = C2()
            nw167_.ctor__(d_1093_v1_, -718, (_dafny.SeqWithoutIsStrInference([d_1130_v22_])).set(default__.safeIndex((self).f24, len(_dafny.SeqWithoutIsStrInference([d_1130_v22_]))), d_1130_v22_))
            d_1131_v23_ = nw167_
            d_1132_v24_: _dafny.Seq
            d_1132_v24_ = _dafny.SeqWithoutIsStrInference([(self).f24])
            d_1133_v25_: _dafny.Seq
            d_1133_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qe"))
            d_1134_v27_: _dafny.Map
            d_1134_v27_ = _dafny.Map({(self).f24: len(_dafny.SeqWithoutIsStrInference([(self).f24]))})
            d_1135_v28_: _dafny.MultiSet
            d_1135_v28_ = _dafny.MultiSet([(self).f24])
            d_1136_v29_: _dafny.Array
            nw168_ = _dafny.Array(None, 13)
            nw168_[int(0)] = (self).f24
            nw168_[int(1)] = (self).f24
            nw168_[int(2)] = (d_1132_v24_)[default__.safeIndex(len(d_1133_v25_), len(d_1132_v24_))]
            def iife62_():
                coll30_ = _dafny.Map()
                compr_30_: int
                for compr_30_ in (d_1134_v27_).keys.Elements:
                    d_1137_v26_: int = compr_30_
                    if (d_1137_v26_) in (d_1134_v27_):
                        coll30_[(d_1137_v26_) * (len(d_1097_v5_))] = -366
                return _dafny.Map(coll30_)
            nw168_[int(3)] = len(iife62_()
            )
            nw168_[int(4)] = (self).f24
            nw168_[int(5)] = (self).f24
            nw168_[int(6)] = (self).f24
            nw168_[int(7)] = (self).f24
            nw168_[int(8)] = 495
            nw168_[int(9)] = (self).f24
            nw168_[int(10)] = 488
            nw168_[int(11)] = (self).f24
            nw168_[int(12)] = (d_1135_v28_).cardinality
            d_1136_v29_ = nw168_
            d_1138_v30_: _dafny.Map
            d_1138_v30_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))): d_1136_v29_})
            d_1139_v31_: D0
            d_1139_v31_ = D0_DC0(d_1136_v29_)
            d_1140_v32_: _dafny.Set
            d_1140_v32_ = _dafny.Set({d_1136_v29_, ((d_1138_v30_)[len(_dafny.Map({d_1139_v31_: d_1136_v29_}))] if (len(_dafny.Map({d_1139_v31_: d_1136_v29_}))) in (d_1138_v30_) else d_1136_v29_), d_1136_v29_, d_1136_v29_, d_1136_v29_})
            d_1140_v32_ = ((d_1140_v32_) - (d_1140_v32_)) | (d_1140_v32_)
            (globalState).f21 = d_1097_v5_
            d_1141_v33_: _dafny.Map
            d_1141_v33_ = _dafny.Map({d_1131_v23_.f31: ((0) - ((self).f24)) > ((0) - ((self).f24))})
            d_1141_v33_ = (d_1141_v33_).set(not(d_1131_v23_.f31), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfodpobt"))) <= (d_1133_v25_))
        d_1142_v34_: _dafny.Seq
        d_1142_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
        d_1143_v35_: _dafny.Map
        d_1143_v35_ = _dafny.Map({d_1093_v1_: d_1142_v34_})
        d_1144_v36_: _dafny.Map
        d_1144_v36_ = _dafny.Map({((self).f24 if d_1093_v1_ else (self).f24): (d_1143_v35_) | (d_1143_v35_)})
        d_1144_v36_ = (d_1144_v36_).set((self).f24, _dafny.Map({d_1093_v1_: d_1142_v34_}))
        (globalState).f17 = d_1142_v34_
        d_1145_v38_: _dafny.Map
        d_1145_v38_ = _dafny.Map({d_1093_v1_: len(_dafny.SeqWithoutIsStrInference([d_1093_v1_, d_1093_v1_, False, not(d_1093_v1_)]))})
        d_1146_v39_: D4
        d_1146_v39_ = D4_DC14(D4_DC13(len(d_1142_v34_), d_1145_v38_, d_1142_v34_))
        d_1147_v40_: _dafny.Set
        d_1147_v40_ = _dafny.Set({d_1146_v39_})
        d_1148_v41_: D7
        def iife63_():
            coll31_ = _dafny.Map()
            compr_31_: D4
            for compr_31_ in (d_1147_v40_).Elements:
                d_1149_v37_: D4 = compr_31_
                if (d_1149_v37_) in (d_1147_v40_):
                    coll31_[d_1149_v37_] = d_1142_v34_
            return _dafny.Map(coll31_)
        d_1148_v41_ = D7_DC20(iife63_()
)
        d_1150_i2_: int
        d_1150_i2_ = 0
        with _dafny.label("4"):
            pat_let_tv9_ = d_1100_v8_
            pat_let_tv10_ = d_1100_v8_
            pat_let_tv11_ = d_1100_v8_
            pat_let_tv12_ = d_1100_v8_
            pat_let_tv13_ = d_1100_v8_
            pat_let_tv14_ = d_1093_v1_
            def lambda48_(source14_):
                if source14_.is_DC21:
                    d_1176___mcc_h5_ = source14_.cf32
                    d_1177___mcc_h6_ = source14_.cf33
                    d_1178___mcc_h7_ = source14_.cf34
                    d_1179_cf34_ = d_1178___mcc_h7_
                    d_1180_cf33_ = d_1177___mcc_h6_
                    d_1181_cf32_ = d_1176___mcc_h5_
                    def iife64_():
                        coll32_ = _dafny.Set()
                        compr_32_: str
                        for compr_32_ in (_dafny.Map({pat_let_tv9_: pat_let_tv10_})).keys.Elements:
                            d_1182_v42_: str = compr_32_
                            if (d_1182_v42_) in (_dafny.Map({pat_let_tv11_: pat_let_tv12_})):
                                coll32_ = coll32_.union(_dafny.Set([d_1182_v42_]))
                        return _dafny.Set(coll32_)
                    return (iife64_()
                    ).ispropersubset(_dafny.Set({pat_let_tv13_}))
                elif source14_.is_DC20:
                    d_1183___mcc_h8_ = source14_.cf31
                    d_1184_cf31_ = d_1183___mcc_h8_
                    return not(not(not(True)))
                elif True:
                    d_1185___mcc_h9_ = source14_.cf35
                    d_1186_cf35_ = d_1185___mcc_h9_
                    return (pat_let_tv14_) == (True)

            while lambda48_(d_1148_v41_):
                with _dafny.c_label("4"):
                    if (d_1150_i2_) >= (100):
                        raise _dafny.Break("4")
                    d_1150_i2_ = (d_1150_i2_) + (1)
                    d_1151_v43_: _dafny.MultiSet
                    d_1151_v43_ = _dafny.MultiSet([len(d_1142_v34_)])
                    d_1152_v44_: bool
                    d_1153_v45_: bool
                    d_1154_v46_: int
                    out28_: bool
                    out29_: bool
                    out30_: int
                    out28_, out29_, out30_ = (self).m5(default__.fm2((d_1151_v43_).cardinality, (self).f24, d_1100_v8_, globalState), globalState)
                    d_1152_v44_ = out28_
                    d_1153_v45_ = out29_
                    d_1154_v46_ = out30_
                    d_1155_v47_: _dafny.Map
                    d_1155_v47_ = _dafny.Map({d_1153_v45_: d_1093_v1_})
                    d_1156_v48_: C2
                    nw169_ = C2()
                    nw169_.ctor__(d_1152_v44_, len(d_1155_v47_), (self).f25)
                    d_1156_v48_ = nw169_
                    d_1157_v49_: _dafny.Seq
                    d_1157_v49_ = _dafny.SeqWithoutIsStrInference([d_1156_v48_])
                    d_1158_v50_: _dafny.Seq
                    d_1158_v50_ = _dafny.SeqWithoutIsStrInference([(self).f24, (0) - (len(d_1157_v49_))])
                    d_1092_v0_ = (d_1092_v0_).set((len(d_1158_v50_)) - (d_1154_v46_), d_1156_v48_.f31)
                    if d_1153_v45_:
                        (globalState).f18 = (d_1154_v46_) + ((self).f24)
                        d_1159_v51_: _dafny.Seq
                        d_1159_v51_ = _dafny.SeqWithoutIsStrInference([d_1155_v47_])
                        d_1160_v52_: _dafny.Seq
                        d_1160_v52_ = _dafny.SeqWithoutIsStrInference([d_1159_v51_, d_1159_v51_])
                        d_1161_v53_: _dafny.Map
                        d_1161_v53_ = _dafny.Map({d_1093_v1_: d_1159_v51_})
                        d_1162_v54_: _dafny.Array
                        nw170_ = _dafny.Array(None, 29)
                        nw170_[int(0)] = (d_1159_v51_ if d_1153_v45_ else d_1159_v51_)
                        nw170_[int(1)] = d_1159_v51_
                        nw170_[int(2)] = d_1159_v51_
                        nw170_[int(3)] = d_1159_v51_
                        nw170_[int(4)] = d_1159_v51_
                        nw170_[int(5)] = d_1159_v51_
                        nw170_[int(6)] = d_1159_v51_
                        nw170_[int(7)] = (default__.fm36(d_1097_v5_, globalState)) + (_dafny.SeqWithoutIsStrInference([d_1155_v47_ for d_1163_i3_ in range(default__.abs(214))]))
                        nw170_[int(8)] = d_1159_v51_
                        nw170_[int(9)] = d_1159_v51_
                        nw170_[int(10)] = d_1159_v51_
                        nw170_[int(11)] = (d_1160_v52_)[default__.safeIndex((0) - ((self).f24), len(d_1160_v52_))]
                        nw170_[int(12)] = d_1159_v51_
                        nw170_[int(13)] = d_1159_v51_
                        nw170_[int(14)] = (d_1159_v51_) + (_dafny.SeqWithoutIsStrInference([d_1155_v47_]))
                        nw170_[int(15)] = ((d_1161_v53_)[d_1156_v48_.f31] if (d_1156_v48_.f31) in (d_1161_v53_) else _dafny.SeqWithoutIsStrInference([d_1155_v47_]))
                        nw170_[int(16)] = _dafny.SeqWithoutIsStrInference([(d_1155_v47_).set(d_1093_v1_, d_1156_v48_.f31) for d_1164_i4_ in range(default__.abs(74))])
                        nw170_[int(17)] = d_1159_v51_
                        nw170_[int(18)] = (d_1159_v51_) + (d_1159_v51_)
                        nw170_[int(19)] = (d_1159_v51_) + (default__.fm36(_dafny.SeqWithoutIsStrInference([d_1093_v1_, d_1093_v1_]), globalState))
                        nw170_[int(20)] = (d_1159_v51_) + (((d_1159_v51_).set(default__.safeIndex(113, len(d_1159_v51_)), d_1155_v47_)).set(default__.safeIndex(d_1154_v46_, len((d_1159_v51_).set(default__.safeIndex(113, len(d_1159_v51_)), d_1155_v47_))), d_1155_v47_))
                        nw170_[int(21)] = _dafny.SeqWithoutIsStrInference([d_1155_v47_ for d_1165_i5_ in range(default__.abs(5))])
                        nw170_[int(22)] = (d_1159_v51_) + (d_1159_v51_)
                        nw170_[int(23)] = d_1159_v51_
                        nw170_[int(24)] = d_1159_v51_
                        nw170_[int(25)] = d_1159_v51_
                        nw170_[int(26)] = _dafny.SeqWithoutIsStrInference([d_1155_v47_, d_1155_v47_, d_1155_v47_])
                        nw170_[int(27)] = _dafny.SeqWithoutIsStrInference([d_1155_v47_])
                        nw170_[int(28)] = (d_1159_v51_) + (_dafny.SeqWithoutIsStrInference([(d_1155_v47_).set(d_1156_v48_.f31, d_1156_v48_.f31) for d_1166_i6_ in range(default__.abs(184))]))
                        d_1162_v54_ = nw170_
                        index147_ = default__.safeIndex(151, (d_1162_v54_).length(0))
                        (d_1162_v54_)[index147_] = d_1159_v51_
                        index148_ = default__.safeIndex(151, (d_1162_v54_).length(0))
                        (d_1162_v54_)[index148_] = default__.fm36(default__.fm23(globalState), globalState)
                        (globalState).f14 = (self).f24
                        d_1167_v55_: D1
                        d_1167_v55_ = D1_DC3(d_1142_v34_)
                        d_1168_v56_: D8
                        d_1168_v56_ = D8_DC26(d_1167_v55_, d_1142_v34_)
                        d_1169_v57_: _dafny.Map
                        d_1169_v57_ = _dafny.Map({(self).fm8(globalState): d_1168_v56_})
                        d_1169_v57_ = (d_1169_v57_).set(d_1154_v46_, default__.fm37(globalState))
                        d_1170_v58_: C2
                        nw171_ = C2()
                        nw171_.ctor__((d_1154_v46_) < (d_1154_v46_), (self).fm8(globalState), (self).f25)
                        d_1170_v58_ = nw171_
                    elif True:
                        d_1171_v59_: _dafny.Seq
                        d_1171_v59_ = _dafny.SeqWithoutIsStrInference([d_1151_v43_])
                        d_1172_v60_: C1
                        nw172_ = C1()
                        nw172_.ctor__((_dafny.SeqWithoutIsStrInference([d_1151_v43_ for d_1173_i7_ in range(default__.abs(694))])) + (d_1171_v59_), (self).f24, ((self).f24) * ((self).f24), (self).f25)
                        d_1172_v60_ = nw172_
                        d_1174_v61_: C1
                        nw173_ = C1()
                        nw173_.ctor__(d_1171_v59_, d_1172_v60_.f33, d_1172_v60_.f33, (self).f25)
                        d_1174_v61_ = nw173_
                        d_1093_v1_ = ((d_1092_v0_)[d_1154_v46_] if (d_1154_v46_) in (d_1092_v0_) else d_1156_v48_.f31)
                        (d_1172_v60_).f33 = default__.safeModuloInt(d_1174_v61_.f33, (self).f24)
                        d_1175_v62_: _dafny.MultiSet
                        d_1175_v62_ = _dafny.MultiSet([d_1156_v48_.f31, d_1093_v1_, (d_1172_v60_.f33) < ((self).f24), d_1156_v48_.f31, d_1156_v48_.f31])
                        d_1175_v62_ = d_1175_v62_
                    (globalState).f18 = (d_1158_v50_)[default__.safeIndex((self).f24, len(d_1158_v50_))]
                    pass
            pass

    @property
    def f29(self):
        return self._f29

class C4(T0):
    def  __init__(self):
        self._f24: int = int(0)
        self._f25: _dafny.Seq = _dafny.Seq({})
        self._f28: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f28, f24, f25):
        (self)._f28 = f28
        (self)._f24 = f24
        (self)._f25 = f25

    def fm5(self, p0, globalState):
        return D0_DC1(len((_dafny.Map({_dafny.CodePoint('j'): (self).f24})) | (_dafny.Map({_dafny.CodePoint('u'): (self).f24}))), (self).f24, True)

    def m1(self, globalState):
        r0: bool = False
        d_1187_v1_: _dafny.MultiSet
        d_1187_v1_ = _dafny.MultiSet([(self).f24])
        d_1188_v2_: bool
        d_1188_v2_ = False
        d_1189_v3_: _dafny.Seq
        d_1189_v3_ = _dafny.SeqWithoutIsStrInference([d_1188_v2_, d_1188_v2_])
        d_1190_v4_: _dafny.Map
        d_1190_v4_ = _dafny.Map({True: (self).f24})
        d_1191_v5_: _dafny.Seq
        d_1191_v5_ = _dafny.SeqWithoutIsStrInference([d_1190_v4_])
        d_1192_v6_: str
        d_1192_v6_ = _dafny.CodePoint('s')
        d_1193_v7_: D0
        d_1193_v7_ = D0_DC2(d_1188_v2_, len((d_1191_v5_)[default__.safeIndex((self).f24, len(d_1191_v5_))]), d_1192_v6_)
        d_1194_v8_: _dafny.Seq
        d_1194_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnox"))
        d_1195_v9_: _dafny.Map
        d_1195_v9_ = _dafny.Map({(self).f24: len(d_1194_v8_)})
        d_1196_v10_: _dafny.Map
        d_1196_v10_ = _dafny.Map({(self).f24: d_1195_v9_})
        d_1197_v12_: _dafny.Array
        nw174_ = _dafny.Array(False, 27)
        d_1197_v12_ = nw174_
        d_1198_v13_: _dafny.Map
        d_1198_v13_ = _dafny.Map({d_1188_v2_: d_1197_v12_})
        d_1199_v14_: _dafny.Array
        nw175_ = _dafny.Array(None, 13)
        def iife65_():
            coll33_ = _dafny.Map()
            compr_33_: int
            for compr_33_ in (d_1187_v1_).Elements:
                d_1200_v0_: int = compr_33_
                if (d_1200_v0_) in (d_1187_v1_):
                    coll33_[default__.safeModuloInt(d_1200_v0_, (self).f24)] = 98
            return _dafny.Map(coll33_)
        nw175_[int(0)] = iife65_()
        
        nw175_[int(1)] = _dafny.Map({len(_dafny.Map({d_1188_v2_: d_1188_v2_})): (self).f24})
        nw175_[int(2)] = (((_dafny.Map({(self).f24: (0) - ((self).f24)})).set((self).f24, (self).f24)).set((self).f24, 698)).set(len(d_1189_v3_), (d_1193_v7_).cf5)
        nw175_[int(3)] = default__.fm6((self).f24, (self).f24, default__.fm0((self).f24, (self).f24, globalState), (self).f24, globalState)
        nw175_[int(4)] = d_1195_v9_
        nw175_[int(5)] = d_1195_v9_
        nw175_[int(6)] = d_1195_v9_
        nw175_[int(7)] = ((d_1196_v10_)[(self).f24] if ((self).f24) in (d_1196_v10_) else d_1195_v9_)
        nw175_[int(8)] = d_1195_v9_
        nw175_[int(9)] = d_1195_v9_
        def iife66_():
            coll34_ = _dafny.Map()
            compr_34_: int
            for compr_34_ in ((default__.fm7(globalState)).set(default__.safeIndex((self).f24, len(default__.fm7(globalState))), (self).f24)).Elements:
                d_1201_v11_: int = compr_34_
                if (d_1201_v11_) in ((default__.fm7(globalState)).set(default__.safeIndex((self).f24, len(default__.fm7(globalState))), (self).f24)):
                    coll34_[(d_1201_v11_) - ((self).f24)] = (_dafny.MultiSet([_dafny.Map({len(d_1194_v8_): d_1188_v2_})])).cardinality
            return _dafny.Map(coll34_)
        nw175_[int(10)] = (_dafny.Map({(self).f24: (self).f24})) | (iife66_()
        )
        nw175_[int(11)] = (d_1195_v9_) | (d_1195_v9_)
        nw175_[int(12)] = (d_1195_v9_).set(len(d_1198_v13_), (0) - ((self).f24))
        d_1199_v14_ = nw175_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1199_v14_).length(0)):
            d_1202_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1202_i0_)) and ((d_1202_i0_) < ((d_1199_v14_).length(0)))):
                (d_1199_v14_)[(d_1202_i0_)] = _dafny.Map({(self).f24: (self).f24})
        d_1203_v15_: _dafny.Set
        d_1203_v15_ = _dafny.Set({575})
        hi7_ = ((self).f24) + (len(d_1203_v15_))
        for d_1204_i1_ in range((self).f24, hi7_):
            d_1205_v16_: _dafny.Array
            def lambda49_(d_1206_v8_):
                def lambda50_(d_1207_i2_):
                    return d_1206_v8_

                return lambda50_

            init29_ = lambda49_(d_1194_v8_)
            nw176_ = _dafny.Array(None, 9)
            for i0_29_ in range(nw176_.length(0)):
                nw176_[i0_29_] = init29_(i0_29_)
            d_1205_v16_ = nw176_
            d_1205_v16_ = d_1205_v16_
            (globalState).f14 = (len((d_1194_v8_).set(default__.safeIndex((self).f24, len(d_1194_v8_)), d_1192_v6_))) - (default__.safeModuloInt((self).f24, (self).f24))
            d_1208_v17_: _dafny.Seq
            d_1208_v17_ = _dafny.SeqWithoutIsStrInference([len(d_1194_v8_), d_1204_i1_, d_1204_i1_])
            d_1209_v18_: D1
            d_1209_v18_ = D1_DC4(((self).f24) <= (d_1204_i1_), d_1188_v2_, d_1192_v6_, len((_dafny.SeqWithoutIsStrInference([default__.fm2(d_1204_i1_, len(d_1194_v8_), d_1192_v6_, globalState)])) + (d_1189_v3_)), (d_1208_v17_)[default__.safeIndex(len((d_1194_v8_).set(default__.safeIndex(len(d_1208_v17_), len(d_1194_v8_)), d_1192_v6_)), len(d_1208_v17_))])
            d_1209_v18_ = d_1209_v18_
            index149_ = default__.safeIndex(864, ((self).f28).length(0))
            ((self).f28)[index149_] = (350) - (d_1204_i1_)
            index150_ = default__.safeIndex(864, ((self).f28).length(0))
            rhs199_ = (0) - (len(((d_1194_v8_) + (d_1194_v8_)).set(default__.safeIndex((0) - ((self).f24), len((d_1194_v8_) + (d_1194_v8_))), d_1192_v6_)))
            rhs200_ = d_1188_v2_
            lhs155_ = (self).f28
            lhs156_ = default__.safeIndex(864, ((self).f28).length(0))
            lhs157_ = globalState
            lhs155_[lhs156_] = rhs199_
            lhs157_.f2 = rhs200_
        d_1210_v19_: _dafny.MultiSet
        d_1210_v19_ = _dafny.MultiSet([d_1192_v6_, d_1192_v6_])
        d_1211_v21_: _dafny.Set
        d_1211_v21_ = _dafny.Set({d_1192_v6_, d_1192_v6_, d_1192_v6_})
        def iife71_():
            coll35_ = _dafny.Set()
            compr_35_: str
            for compr_35_ in (d_1210_v19_).Elements:
                d_1234_v20_: str = compr_35_
                if (d_1234_v20_) in (d_1210_v19_):
                    coll35_ = coll35_.union(_dafny.Set([d_1234_v20_]))
            return _dafny.Set(coll35_)
        hi8_ = (self).f24
        for d_1212_i3_ in range(len((iife71_()
        ).intersection(d_1211_v21_)), hi8_):
            if d_1188_v2_:
                index151_ = default__.safeIndex(302, (d_1197_v12_).length(0))
                def iife67_(_pat_let16_0):
                    def iife68_(d_1213_dt__update__tmp_h0_):
                        def iife69_(_pat_let17_0):
                            def iife70_(d_1214_dt__update_hcf5_h0_):
                                return D0_DC2((d_1213_dt__update__tmp_h0_).cf4, d_1214_dt__update_hcf5_h0_, (d_1213_dt__update__tmp_h0_).cf6)
                            return iife70_(_pat_let17_0)
                        return iife69_(d_1212_i3_)
                    return iife68_(_pat_let16_0)
                (d_1197_v12_)[index151_] = (iife67_(d_1193_v7_)).cf4
                index152_ = default__.safeIndex(302, (d_1197_v12_).length(0))
                (d_1197_v12_)[index152_] = default__.fm2((0) - ((d_1212_i3_) + (len(d_1195_v9_))), d_1212_i3_, d_1192_v6_, globalState)
                d_1215_v22_: _dafny.Seq
                d_1215_v22_ = _dafny.SeqWithoutIsStrInference([len(d_1194_v8_)])
                d_1216_v23_: _dafny.Seq
                d_1216_v23_ = _dafny.SeqWithoutIsStrInference([d_1187_v1_, _dafny.MultiSet([d_1212_i3_])])
                d_1217_v24_: _dafny.Map
                d_1217_v24_ = _dafny.Map({d_1192_v6_: d_1212_i3_})
                d_1218_v25_: T0
                nw177_ = C1()
                nw177_.ctor__(d_1216_v23_, (self).f24, (self).f24, _dafny.SeqWithoutIsStrInference([d_1217_v24_, d_1217_v24_, d_1217_v24_]))
                d_1218_v25_ = nw177_
                d_1219_v26_: _dafny.Map
                d_1219_v26_ = _dafny.Map({d_1218_v25_: d_1188_v2_})
                (globalState).f1 = default__.safeDivisionInt((len(d_1215_v22_)) * ((0) - (len(d_1215_v22_))), len(d_1219_v26_))
                (globalState).f18 = d_1212_i3_
                d_1220_v27_: _dafny.MultiSet
                d_1220_v27_ = _dafny.MultiSet([d_1195_v9_])
                (globalState).f2 = (d_1220_v27_).ispropersubset(d_1220_v27_)
                rhs201_ = _dafny.Set({(0) - (d_1212_i3_)})
                rhs202_ = ((d_1194_v8_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))) + (d_1194_v8_)
                rhs203_ = d_1188_v2_
                lhs158_ = globalState
                d_1203_v15_ = rhs201_
                lhs158_.f17 = rhs202_
                r0 = rhs203_
            elif True:
                d_1221_v28_: _dafny.Seq
                d_1221_v28_ = _dafny.SeqWithoutIsStrInference([d_1212_i3_])
                d_1222_v29_: D4
                d_1222_v29_ = D4_DC12(d_1221_v28_)
                d_1223_v30_: D4
                d_1223_v30_ = D4_DC14(d_1222_v29_)
                d_1224_v31_: _dafny.Map
                d_1224_v31_ = _dafny.Map({d_1223_v30_: d_1194_v8_})
                d_1225_v32_: D7
                d_1225_v32_ = D7_DC20(d_1224_v31_)
                d_1226_v33_: _dafny.MultiSet
                d_1226_v33_ = _dafny.MultiSet([d_1225_v32_, D7_DC20(d_1224_v31_), d_1225_v32_])
                d_1226_v33_ = d_1226_v33_
                (globalState).f14 = ((self).f24) * (d_1212_i3_)
                d_1227_v34_: _dafny.Map
                d_1227_v34_ = _dafny.Map({d_1188_v2_: d_1188_v2_})
                d_1227_v34_ = (d_1227_v34_).set(not(d_1188_v2_), d_1188_v2_)
                (globalState).f2 = d_1188_v2_
                d_1228_v35_: _dafny.MultiSet
                d_1228_v35_ = _dafny.MultiSet([d_1188_v2_])
                d_1229_v36_: D3
                d_1229_v36_ = D3_DC10()
                d_1230_v37_: _dafny.MultiSet
                d_1230_v37_ = _dafny.MultiSet([d_1229_v36_])
                d_1228_v35_ = _dafny.MultiSet([(d_1230_v37_).ispropersubset(d_1230_v37_)])
            index153_ = default__.safeIndex(437, (d_1197_v12_).length(0))
            (d_1197_v12_)[index153_] = False
            d_1231_v38_: _dafny.Seq
            d_1231_v38_ = _dafny.SeqWithoutIsStrInference([d_1194_v8_])
            index154_ = default__.safeIndex(437, (d_1197_v12_).length(0))
            rhs204_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xho"))) + (d_1194_v8_)) < (_dafny.SeqWithoutIsStrInference([d_1192_v6_ for d_1232_i4_ in range(default__.abs(783))]))
            rhs205_ = d_1210_v19_
            rhs206_ = _dafny.SeqWithoutIsStrInference([d_1194_v8_, d_1194_v8_])
            rhs207_ = d_1188_v2_
            lhs159_ = d_1197_v12_
            lhs160_ = default__.safeIndex(437, (d_1197_v12_).length(0))
            lhs161_ = globalState
            lhs159_[lhs160_] = rhs204_
            d_1210_v19_ = rhs205_
            d_1231_v38_ = rhs206_
            lhs161_.f2 = rhs207_
            d_1190_v4_ = (d_1190_v4_).set(not((d_1197_v12_)[default__.safeIndex(437, (d_1197_v12_).length(0))]), d_1212_i3_)
            d_1233_v39_: _dafny.Seq
            d_1233_v39_ = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f28, (self).f28])
            index155_ = default__.safeIndex(437, (d_1197_v12_).length(0))
            rhs208_ = d_1194_v8_
            rhs209_ = (d_1233_v39_).set(default__.safeIndex((self).f24, len(d_1233_v39_)), (self).f28)
            rhs210_ = d_1203_v15_
            rhs211_ = d_1188_v2_
            rhs212_ = not ((d_1197_v12_)[default__.safeIndex(437, (d_1197_v12_).length(0))]) or (d_1188_v2_)
            lhs162_ = globalState
            lhs163_ = d_1197_v12_
            lhs164_ = default__.safeIndex(437, (d_1197_v12_).length(0))
            lhs165_ = globalState
            lhs162_.f17 = rhs208_
            d_1233_v39_ = rhs209_
            d_1203_v15_ = rhs210_
            lhs163_[lhs164_] = rhs211_
            lhs165_.f2 = rhs212_
        d_1235_i5_: int
        d_1235_i5_ = 0
        with _dafny.label("5"):
            while ((self).f24) == ((self).f24):
                with _dafny.c_label("5"):
                    if (d_1235_i5_) >= (100):
                        raise _dafny.Break("5")
                    d_1235_i5_ = (d_1235_i5_) + (1)
                    pat_let_tv15_ = d_1188_v2_
                    d_1236_v40_: _dafny.MultiSet
                    def iife72_(_pat_let18_0):
                        def iife73_(d_1237_dt__update__tmp_h1_):
                            def iife74_(_pat_let19_0):
                                def iife75_(d_1238_dt__update_hcf20_h0_):
                                    return D3_DC11((d_1237_dt__update__tmp_h1_).cf18, (d_1237_dt__update__tmp_h1_).cf19, d_1238_dt__update_hcf20_h0_)
                                return iife75_(_pat_let19_0)
                            return iife74_(pat_let_tv15_)
                        return iife73_(_pat_let18_0)
                    d_1236_v40_ = _dafny.MultiSet([iife72_(D3_DC11(default__.fm0((self).f24, -151, globalState), False, d_1188_v2_)), D3_DC11((self).f24, d_1188_v2_, d_1188_v2_)])
                    if not((d_1236_v40_).issubset(default__.fm38(globalState))):
                        index156_ = default__.safeIndex(211, ((self).f28).length(0))
                        ((self).f28)[index156_] = len(d_1203_v15_)
                        index157_ = default__.safeIndex(211, ((self).f28).length(0))
                        rhs213_ = (((self).f24) + ((self).f24)) + (default__.fm0((self).f24, (self).f24, globalState))
                        rhs214_ = (self).f24
                        lhs166_ = (self).f28
                        lhs167_ = default__.safeIndex(211, ((self).f28).length(0))
                        lhs168_ = globalState
                        lhs166_[lhs167_] = rhs213_
                        lhs168_.f14 = rhs214_
                        (globalState).f17 = (_dafny.SeqWithoutIsStrInference([d_1192_v6_ for d_1239_i6_ in range(default__.abs(244))])) + (d_1194_v8_)
                        r0 = d_1188_v2_
                        index158_ = default__.safeIndex(211, ((self).f28).length(0))
                        index159_ = default__.safeIndex(211, ((self).f28).length(0))
                        rhs215_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
                        rhs216_ = (0) - (default__.safeModuloInt(337, (0) - ((self).f24)))
                        rhs217_ = (d_1210_v19_).cardinality
                        lhs169_ = globalState
                        lhs170_ = (self).f28
                        lhs171_ = default__.safeIndex(211, ((self).f28).length(0))
                        lhs172_ = (self).f28
                        lhs173_ = default__.safeIndex(211, ((self).f28).length(0))
                        lhs169_.f17 = rhs215_
                        lhs170_[lhs171_] = rhs216_
                        lhs172_[lhs173_] = rhs217_
                        d_1240_v41_: _dafny.Array
                        nw178_ = _dafny.Array(None, 28)
                        d_1240_v41_ = nw178_
                        d_1241_v42_: T0
                        nw179_ = C2()
                        nw179_.ctor__(d_1188_v2_, ((self).f28)[default__.safeIndex(211, ((self).f28).length(0))], (self).f25)
                        d_1241_v42_ = nw179_
                        d_1242_v43_: _dafny.Map
                        d_1242_v43_ = _dafny.Map({(d_1189_v3_)[default__.safeIndex(((self).f28)[default__.safeIndex(211, ((self).f28).length(0))], len(d_1189_v3_))]: d_1241_v42_})
                        index160_ = default__.safeIndex(134, (d_1240_v41_).length(0))
                        (d_1240_v41_)[index160_] = ((d_1242_v43_)[d_1188_v2_] if (d_1188_v2_) in (d_1242_v43_) else d_1241_v42_)
                        d_1243_v44_: _dafny.Map
                        d_1243_v44_ = _dafny.Map({d_1195_v9_: d_1241_v42_})
                        index161_ = default__.safeIndex(134, (d_1240_v41_).length(0))
                        (d_1240_v41_)[index161_] = ((d_1243_v44_)[(d_1195_v9_) | (d_1195_v9_)] if ((d_1195_v9_) | (d_1195_v9_)) in (d_1243_v44_) else d_1241_v42_)
                    elif True:
                        d_1244_v45_: _dafny.Array
                        nw180_ = _dafny.Array(_dafny.Array(None, 0), 28)
                        d_1244_v45_ = nw180_
                        d_1245_v46_: _dafny.Map
                        d_1245_v46_ = _dafny.Map({(self).f24: d_1188_v2_})
                        d_1246_v47_: D1
                        d_1246_v47_ = D1_DC3(default__.fm4(False, (self).f24, d_1192_v6_, d_1188_v2_, globalState))
                        d_1247_v48_: D8
                        d_1247_v48_ = D8_DC26(d_1246_v47_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))
                        d_1248_v49_: _dafny.Array
                        nw181_ = _dafny.Array(None, 19)
                        nw181_[int(0)] = d_1194_v8_
                        nw181_[int(1)] = d_1194_v8_
                        nw181_[int(2)] = d_1194_v8_
                        nw181_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "an"))
                        nw181_[int(4)] = _dafny.SeqWithoutIsStrInference([(D1_DC4(d_1188_v2_, d_1188_v2_, d_1192_v6_, (self).f24, len(_dafny.SeqWithoutIsStrInference([len(d_1194_v8_)])))).cf10 for d_1249_i7_ in range(default__.abs(671))])
                        nw181_[int(5)] = _dafny.SeqWithoutIsStrInference([d_1192_v6_ for d_1250_i8_ in range(default__.abs(187))])
                        nw181_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_1192_v6_ for d_1251_i9_ in range(default__.abs(-797))])).set(default__.safeIndex(len(d_1245_v46_), len(_dafny.SeqWithoutIsStrInference([d_1192_v6_ for d_1252_i9_ in range(default__.abs(-797))]))), d_1192_v6_)
                        nw181_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kncwpm"))
                        nw181_[int(8)] = d_1194_v8_
                        nw181_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpjutg"))
                        nw181_[int(10)] = d_1194_v8_
                        nw181_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wx"))
                        nw181_[int(12)] = d_1194_v8_
                        nw181_[int(13)] = d_1194_v8_
                        nw181_[int(14)] = _dafny.SeqWithoutIsStrInference([d_1192_v6_ for d_1253_i10_ in range(default__.abs(911))])
                        nw181_[int(15)] = d_1194_v8_
                        nw181_[int(16)] = d_1194_v8_
                        nw181_[int(17)] = d_1194_v8_
                        nw181_[int(18)] = (d_1247_v48_).cf42
                        d_1248_v49_ = nw181_
                        d_1254_v50_: D10
                        d_1254_v50_ = D10_DC28(d_1248_v49_)
                        index162_ = default__.safeIndex(552, (d_1244_v45_).length(0))
                        (d_1244_v45_)[index162_] = (d_1254_v50_).cf44
                        d_1255_v51_: _dafny.Map
                        d_1255_v51_ = _dafny.Map({d_1188_v2_: d_1248_v49_})
                        index163_ = default__.safeIndex(552, (d_1244_v45_).length(0))
                        (d_1244_v45_)[index163_] = ((d_1255_v51_)[((self).f24) < ((0) - ((self).f24))] if (((self).f24) < ((0) - ((self).f24))) in (d_1255_v51_) else d_1248_v49_)
                        d_1190_v4_ = d_1190_v4_
                        d_1256_v52_: _dafny.Array
                        nw182_ = _dafny.Array(_dafny.Set({}), 26)
                        d_1256_v52_ = nw182_
                        d_1257_v53_: _dafny.Map
                        d_1257_v53_ = _dafny.Map({d_1192_v6_: (self).f24})
                        d_1258_v54_: C3
                        nw183_ = C3()
                        nw183_.ctor__(d_1256_v52_, (d_1199_v14_ if d_1188_v2_ else d_1199_v14_), default__.safeDivisionInt((D8_DC25((self).f24, d_1188_v2_, d_1188_v2_)).cf38, (self).f24), (((self).f25) + ((self).f25)).set(default__.safeIndex((self).f24, len(((self).f25) + ((self).f25))), d_1257_v53_))
                        d_1258_v54_ = nw183_
                        d_1259_v56_: D1
                        def iife76_():
                            coll36_ = _dafny.Map()
                            compr_36_: int
                            for compr_36_ in _dafny.IntegerRange(784, 805):
                                d_1260_v55_: int = compr_36_
                                if ((784) <= (d_1260_v55_)) and ((d_1260_v55_) < (805)):
                                    coll36_[(d_1260_v55_) + ((self).f24)] = d_1192_v6_
                            return _dafny.Map(coll36_)
                        d_1259_v56_ = D1_DC4(d_1188_v2_, d_1188_v2_, d_1192_v6_, (self).f24, len(iife76_()
))
                        r0 = ((d_1259_v56_).cf12) >= ((self).f24)
                        (globalState).f2 = not(default__.fm2(((self).f24) * (len((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([d_1188_v2_]))) for d_1261_i11_ in range(default__.abs(125))])).set(default__.safeIndex((self).f24, len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([d_1188_v2_]))) for d_1262_i11_ in range(default__.abs(125))]))), (self).f24))), -970, d_1192_v6_, globalState))
                    (globalState).f2 = False
                    d_1263_v57_: _dafny.Array
                    nw184_ = _dafny.Array(_dafny.Set({}), 5)
                    d_1263_v57_ = nw184_
                    d_1264_v58_: _dafny.Map
                    d_1264_v58_ = _dafny.Map({d_1192_v6_: (self).f24})
                    d_1265_v59_: C3
                    nw185_ = C3()
                    nw185_.ctor__(d_1263_v57_, d_1199_v14_, default__.safeModuloInt((self).f24, (self).f24), _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1192_v6_: (self).f24}), d_1264_v58_]))
                    d_1265_v59_ = nw185_
                    d_1266_v60_: _dafny.Map
                    d_1266_v60_ = _dafny.Map({d_1197_v12_: False})
                    d_1267_v61_: C2
                    nw186_ = C2()
                    nw186_.ctor__(((d_1266_v60_)[d_1197_v12_] if (d_1197_v12_) in (d_1266_v60_) else not(not(d_1188_v2_))), (self).f24, ((self).f25).set(default__.safeIndex((self).f24, len((self).f25)), d_1264_v58_))
                    d_1267_v61_ = nw186_
                    pass
            pass
        d_1268_i12_: int
        d_1268_i12_ = 0
        with _dafny.label("6"):
            while d_1188_v2_:
                with _dafny.c_label("6"):
                    if (d_1268_i12_) >= (100):
                        raise _dafny.Break("6")
                    d_1268_i12_ = (d_1268_i12_) + (1)
                    (globalState).f18 = (self).f24
                    (globalState).f14 = (self).f24
                    d_1269_v62_: C2
                    nw187_ = C2()
                    nw187_.ctor__(not(d_1188_v2_), (self).f24, (self).f25)
                    d_1269_v62_ = nw187_
                    if d_1188_v2_:
                        d_1270_v65_: _dafny.Array
                        nw188_ = _dafny.Array(None, 22)
                        nw188_[int(0)] = _dafny.Set({(self).f24})
                        nw188_[int(1)] = _dafny.Set({(self).f24, (self).f24})
                        nw188_[int(2)] = d_1203_v15_
                        nw188_[int(3)] = d_1203_v15_
                        nw188_[int(4)] = _dafny.Set({(self).f24, (self).f24})
                        nw188_[int(5)] = d_1203_v15_
                        nw188_[int(6)] = d_1203_v15_
                        nw188_[int(7)] = d_1203_v15_
                        nw188_[int(8)] = d_1203_v15_
                        nw188_[int(9)] = d_1203_v15_
                        def iife77_():
                            coll37_ = _dafny.Set()
                            compr_37_: int
                            for compr_37_ in _dafny.IntegerRange(678, 758):
                                d_1271_v63_: int = compr_37_
                                if ((678) <= (d_1271_v63_)) and ((d_1271_v63_) < (758)):
                                    coll37_ = coll37_.union(_dafny.Set([(d_1271_v63_) * ((self).f24)]))
                            return _dafny.Set(coll37_)
                        nw188_[int(10)] = iife77_()
                        
                        def iife78_():
                            coll38_ = _dafny.Set()
                            compr_38_: int
                            for compr_38_ in _dafny.IntegerRange(494, -646):
                                d_1272_v64_: int = compr_38_
                                if ((494) <= (d_1272_v64_)) and ((d_1272_v64_) < (-646)):
                                    coll38_ = coll38_.union(_dafny.Set([(d_1272_v64_) + ((self).f24)]))
                            return _dafny.Set(coll38_)
                        nw188_[int(11)] = iife78_()
                        
                        nw188_[int(12)] = d_1203_v15_
                        nw188_[int(13)] = d_1203_v15_
                        nw188_[int(14)] = d_1203_v15_
                        nw188_[int(15)] = d_1203_v15_
                        nw188_[int(16)] = d_1203_v15_
                        nw188_[int(17)] = d_1203_v15_
                        nw188_[int(18)] = d_1203_v15_
                        nw188_[int(19)] = d_1203_v15_
                        nw188_[int(20)] = _dafny.Set({(self).f24})
                        nw188_[int(21)] = d_1203_v15_
                        d_1270_v65_ = nw188_
                        d_1273_v66_: C3
                        nw189_ = C3()
                        nw189_.ctor__(d_1270_v65_, d_1199_v14_, len(d_1194_v8_), (self).f25)
                        d_1273_v66_ = nw189_
                        d_1274_v67_: _dafny.Seq
                        d_1274_v67_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                        d_1275_v68_: _dafny.Seq
                        d_1275_v68_ = _dafny.SeqWithoutIsStrInference([(d_1274_v67_)[default__.safeIndex((self).f24, len(d_1274_v67_))]])
                        (globalState).f2 = (not((d_1275_v68_) < (_dafny.SeqWithoutIsStrInference([717 for d_1276_i13_ in range(default__.abs(300))]))) if d_1188_v2_ else d_1269_v62_.f31)
                        (globalState).f17 = d_1194_v8_
                        d_1277_v69_: _dafny.Map
                        d_1277_v69_ = _dafny.Map({d_1197_v12_: d_1188_v2_})
                        d_1277_v69_ = (d_1277_v69_).set(d_1197_v12_, d_1269_v62_.f31)
                        d_1278_v70_: _dafny.Map
                        d_1278_v70_ = _dafny.Map({d_1274_v67_: d_1269_v62_.f31})
                        index164_ = default__.safeIndex(638, ((self).f28).length(0))
                        ((self).f28)[index164_] = default__.safeModuloInt(100, (0) - (len(d_1278_v70_)))
                        index165_ = default__.safeIndex(638, ((self).f28).length(0))
                        ((self).f28)[index165_] = (self).f24
                    elif True:
                        d_1279_v71_: _dafny.Array
                        def lambda51_(d_1280_v15_):
                            def lambda52_(d_1281_i14_):
                                return d_1280_v15_

                            return lambda52_

                        init30_ = lambda51_(d_1203_v15_)
                        nw190_ = _dafny.Array(None, 23)
                        for i0_30_ in range(nw190_.length(0)):
                            nw190_[i0_30_] = init30_(i0_30_)
                        d_1279_v71_ = nw190_
                        d_1282_v73_: _dafny.Map
                        d_1282_v73_ = _dafny.Map({d_1192_v6_: d_1269_v62_.f31})
                        d_1283_v74_: _dafny.Map
                        d_1283_v74_ = _dafny.Map({d_1192_v6_: (self).f24})
                        d_1284_v75_: C3
                        nw191_ = C3()
                        def iife79_():
                            coll39_ = _dafny.Map()
                            compr_39_: str
                            for compr_39_ in (d_1282_v73_).keys.Elements:
                                d_1285_v72_: str = compr_39_
                                if (d_1285_v72_) in (d_1282_v73_):
                                    coll39_[d_1285_v72_] = len(_dafny.Set({(D3_DC11((self).f24, d_1269_v62_.f31, True)).cf20}))
                            return _dafny.Map(coll39_)
                        nw191_.ctor__(d_1279_v71_, d_1199_v14_, (self).f24, _dafny.SeqWithoutIsStrInference([iife79_()
                        , d_1283_v74_, _dafny.Map({d_1192_v6_: (0) - ((self).f24)}), d_1283_v74_]))
                        d_1284_v75_ = nw191_
                        d_1284_v75_ = d_1284_v75_
                        d_1286_v76_: _dafny.Map
                        d_1286_v76_ = _dafny.Map({d_1269_v62_.f31: d_1189_v3_})
                        d_1287_v77_: C2
                        nw192_ = C2()
                        nw192_.ctor__((d_1188_v2_) not in (d_1286_v76_), (self).f24, _dafny.SeqWithoutIsStrInference([d_1283_v74_]))
                        d_1287_v77_ = nw192_
                        index166_ = default__.safeIndex(42, ((self).f28).length(0))
                        ((self).f28)[index166_] = (self).f24
                        index167_ = default__.safeIndex(42, ((self).f28).length(0))
                        rhs218_ = d_1269_v62_.f31
                        rhs219_ = False
                        rhs220_ = d_1269_v62_.f31
                        rhs221_ = (self).f24
                        rhs222_ = (self).f24
                        lhs174_ = globalState
                        lhs175_ = globalState
                        lhs176_ = globalState
                        lhs177_ = (self).f28
                        lhs178_ = default__.safeIndex(42, ((self).f28).length(0))
                        lhs179_ = globalState
                        lhs174_.f2 = rhs218_
                        lhs175_.f2 = rhs219_
                        lhs176_.f2 = rhs220_
                        lhs177_[lhs178_] = rhs221_
                        lhs179_.f9 = rhs222_
                        index168_ = default__.safeIndex(42, ((self).f28).length(0))
                        ((self).f28)[index168_] = ((self).f28)[default__.safeIndex(42, ((self).f28).length(0))]
                        (globalState).f9 = ((self).f28)[default__.safeIndex(42, ((self).f28).length(0))]
                    pass
            pass
        d_1288_v78_: _dafny.Map
        d_1288_v78_ = _dafny.Map({d_1188_v2_: d_1193_v7_})
        hi9_ = len(d_1288_v78_)
        for d_1289_i15_ in range(len(default__.fm39((self).f24, d_1188_v2_, globalState)), hi9_):
            (globalState).f9 = (self).f24
            d_1290_v79_: _dafny.Array
            nw193_ = _dafny.Array(_dafny.Set({}), 16)
            d_1290_v79_ = nw193_
            d_1291_v80_: _dafny.Set
            d_1291_v80_ = _dafny.Set({True})
            index169_ = default__.safeIndex(580, (d_1290_v79_).length(0))
            (d_1290_v79_)[index169_] = (default__.fm28(d_1188_v2_, (0) - (d_1289_i15_), True, (0) - (d_1289_i15_), globalState)).intersection(d_1291_v80_)
            index170_ = default__.safeIndex(580, (d_1290_v79_).length(0))
            (d_1290_v79_)[index170_] = (d_1291_v80_) | ((d_1291_v80_) - (_dafny.Set({d_1188_v2_})))
            if (d_1289_i15_) != ((self).f24):
                d_1292_v81_: _dafny.Array
                def lambda53_(d_1293_i15_, d_1294_v8_):
                    def lambda54_(d_1295_i16_):
                        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1293_i15_, 297, d_1293_i15_, d_1293_i15_, d_1293_i15_])): d_1294_v8_})

                    return lambda54_

                init31_ = lambda53_(d_1289_i15_, d_1194_v8_)
                nw194_ = _dafny.Array(None, 16)
                for i0_31_ in range(nw194_.length(0)):
                    nw194_[i0_31_] = init31_(i0_31_)
                d_1292_v81_ = nw194_
                d_1296_v82_: _dafny.Map
                d_1296_v82_ = _dafny.Map({(self).f24: d_1194_v8_})
                index171_ = default__.safeIndex(796, (d_1292_v81_).length(0))
                (d_1292_v81_)[index171_] = d_1296_v82_
                index172_ = default__.safeIndex(796, (d_1292_v81_).length(0))
                (d_1292_v81_)[index172_] = (d_1296_v82_).set(-230, d_1194_v8_)
                (globalState).f9 = d_1289_i15_
                d_1297_v84_: _dafny.MultiSet
                def iife80_():
                    coll40_ = _dafny.Set()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(719, -613):
                        d_1298_v83_: int = compr_40_
                        if ((719) <= (d_1298_v83_)) and ((d_1298_v83_) < (-613)):
                            coll40_ = coll40_.union(_dafny.Set([default__.safeModuloInt(d_1298_v83_, (self).f24)]))
                    return _dafny.Set(coll40_)
                d_1297_v84_ = _dafny.MultiSet([(iife80_()
                ) == (d_1203_v15_), d_1188_v2_, (d_1188_v2_) == (d_1188_v2_), not (d_1188_v2_) or (d_1188_v2_)])
                d_1299_v85_: C2
                nw195_ = C2()
                nw195_.ctor__(not(d_1188_v2_), default__.safeDivisionInt(d_1289_i15_, d_1289_i15_), ((self).f25).set(default__.safeIndex(d_1289_i15_, len((self).f25)), _dafny.Map({d_1192_v6_: (self).f24})))
                d_1299_v85_ = nw195_
                rhs223_ = d_1289_i15_
                rhs224_ = _dafny.MultiSet([not(d_1188_v2_)])
                rhs225_ = (d_1188_v2_ if (d_1188_v2_ if d_1188_v2_ else d_1188_v2_) else not (d_1299_v85_.f31) or (d_1299_v85_.f31))
                rhs226_ = (default__.fm0(d_1289_i15_, (self).f24, globalState)) + (d_1289_i15_)
                rhs227_ = d_1299_v85_
                lhs180_ = globalState
                lhs181_ = globalState
                lhs180_.f9 = rhs223_
                d_1297_v84_ = rhs224_
                d_1188_v2_ = rhs225_
                lhs181_.f14 = rhs226_
                d_1299_v85_ = rhs227_
                index173_ = default__.safeIndex(768, ((self).f28).length(0))
                ((self).f28)[index173_] = d_1289_i15_
                index174_ = default__.safeIndex(768, ((self).f28).length(0))
                ((self).f28)[index174_] = d_1289_i15_
                d_1300_v86_: D11
                d_1300_v86_ = D11_DC30(d_1195_v9_)
                d_1301_v87_: _dafny.Map
                d_1301_v87_ = _dafny.Map({(d_1194_v8_) + (d_1194_v8_): d_1188_v2_})
                index175_ = default__.safeIndex(768, ((self).f28).length(0))
                rhs228_ = ((self).f28)[default__.safeIndex(768, ((self).f28).length(0))]
                rhs229_ = d_1192_v6_
                rhs230_ = (_dafny.MultiSet([d_1195_v9_])).ispropersubset(_dafny.MultiSet([d_1195_v9_, (d_1300_v86_).cf47, d_1195_v9_]))
                rhs231_ = ((d_1301_v87_)[d_1194_v8_] if (d_1194_v8_) in (d_1301_v87_) else not(d_1188_v2_))
                lhs182_ = (self).f28
                lhs183_ = default__.safeIndex(768, ((self).f28).length(0))
                lhs184_ = d_1299_v85_
                lhs185_ = globalState
                lhs182_[lhs183_] = rhs228_
                d_1192_v6_ = rhs229_
                lhs184_.f31 = rhs230_
                lhs185_.f2 = rhs231_
            elif True:
                d_1302_v88_: _dafny.Map
                d_1302_v88_ = _dafny.Map({d_1289_i15_: d_1197_v12_})
                rhs232_ = d_1188_v2_
                rhs233_ = (d_1194_v8_) + (d_1194_v8_)
                rhs234_ = d_1302_v88_
                lhs186_ = globalState
                lhs187_ = globalState
                lhs186_.f2 = rhs232_
                lhs187_.f17 = rhs233_
                d_1302_v88_ = rhs234_
                index176_ = default__.safeIndex(351, (d_1197_v12_).length(0))
                (d_1197_v12_)[index176_] = (d_1189_v3_)[default__.safeIndex(d_1289_i15_, len(d_1189_v3_))]
                index177_ = default__.safeIndex(351, (d_1197_v12_).length(0))
                (d_1197_v12_)[index177_] = (d_1289_i15_) == (len((d_1190_v4_).set(d_1188_v2_, (self).f24)))
                d_1303_v89_: _dafny.Array
                def lambda55_(d_1304_v12_, d_1305_v2_):
                    def lambda56_(d_1306_i17_):
                        return not ((d_1304_v12_)[default__.safeIndex(351, (d_1304_v12_).length(0))]) or (d_1305_v2_)

                    return lambda56_

                init32_ = lambda55_(d_1197_v12_, d_1188_v2_)
                nw196_ = _dafny.Array(None, 28)
                for i0_32_ in range(nw196_.length(0)):
                    nw196_[i0_32_] = init32_(i0_32_)
                d_1303_v89_ = nw196_
                d_1303_v89_ = d_1303_v89_
                d_1307_v90_: D8
                d_1307_v90_ = D8_DC25(d_1289_i15_, d_1188_v2_, (d_1197_v12_)[default__.safeIndex(351, (d_1197_v12_).length(0))])
                d_1308_v91_: _dafny.MultiSet
                d_1308_v91_ = _dafny.MultiSet([(d_1194_v8_).set(default__.safeIndex((self).f24, len(d_1194_v8_)), _dafny.CodePoint('a')), default__.fm4((d_1197_v12_)[default__.safeIndex(351, (d_1197_v12_).length(0))], d_1289_i15_, d_1192_v6_, (d_1307_v90_).cf40, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhh")), (d_1194_v8_) + (d_1194_v8_), d_1194_v8_])
                d_1309_v92_: _dafny.Seq
                d_1309_v92_ = _dafny.SeqWithoutIsStrInference([d_1308_v91_, (_dafny.MultiSet([d_1194_v8_, d_1194_v8_]) if (d_1197_v12_)[default__.safeIndex(351, (d_1197_v12_).length(0))] else d_1308_v91_), _dafny.MultiSet([d_1194_v8_, d_1194_v8_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqkivho"))]), ((d_1308_v91_).set(d_1194_v8_, default__.abs(d_1289_i15_))).intersection(d_1308_v91_)])
                d_1308_v91_ = (d_1309_v92_)[default__.safeIndex((self).f24, len(d_1309_v92_))]
                d_1194_v8_ = (d_1194_v8_) + (d_1194_v8_)
            index178_ = default__.safeIndex(384, ((self).f28).length(0))
            ((self).f28)[index178_] = (self).f24
            index179_ = default__.safeIndex(384, ((self).f28).length(0))
            ((self).f28)[index179_] = (self).f24
        d_1310_v93_: _dafny.Seq
        d_1310_v93_ = _dafny.SeqWithoutIsStrInference([d_1187_v1_, d_1187_v1_])
        d_1311_v94_: C1
        nw197_ = C1()
        nw197_.ctor__(d_1310_v93_, (self).f24, (self).f24, _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1192_v6_: (self).f24}) for d_1312_i18_ in range(default__.abs(526))]))
        d_1311_v94_ = nw197_
        d_1313_v95_: _dafny.Map
        d_1313_v95_ = _dafny.Map({(self).f24: d_1311_v94_})
        d_1314_v96_: _dafny.Seq
        d_1314_v96_ = _dafny.SeqWithoutIsStrInference([default__.fm0((self).f24, len(d_1313_v95_), globalState)])
        r0 = (140) == ((d_1314_v96_)[default__.safeIndex(len(d_1203_v15_), len(d_1314_v96_))])
        return r0

    def m4(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.Seq({})
        r3: bool = False
        d_1315_v0_: _dafny.Map
        d_1315_v0_ = _dafny.Map({p1: (self).f24})
        d_1316_v1_: _dafny.Seq
        d_1316_v1_ = _dafny.SeqWithoutIsStrInference([d_1315_v0_])
        d_1317_v2_: _dafny.MultiSet
        d_1317_v2_ = _dafny.MultiSet([p1, p1])
        d_1318_v3_: str
        d_1318_v3_ = _dafny.CodePoint('u')
        d_1319_v4_: _dafny.Map
        d_1319_v4_ = _dafny.Map({len(_dafny.Map({p0: d_1318_v3_})): len(_dafny.SeqWithoutIsStrInference([p1]))})
        d_1320_v5_: _dafny.Map
        d_1320_v5_ = _dafny.Map({p0: D3_DC11((self).f24, p1, p1)})
        d_1321_v6_: D12
        d_1321_v6_ = D12_DC32(d_1320_v5_)
        d_1322_v7_: _dafny.Set
        d_1322_v7_ = _dafny.Set({p0})
        d_1323_v8_: _dafny.Array
        nw198_ = _dafny.Array(None, 16)
        nw198_[int(0)] = len((d_1316_v1_) + (d_1316_v1_))
        nw198_[int(1)] = (self).f24
        nw198_[int(2)] = (self).f24
        nw198_[int(3)] = (self).f24
        nw198_[int(4)] = (self).f24
        nw198_[int(5)] = (self).f24
        nw198_[int(6)] = (self).f24
        nw198_[int(7)] = ((d_1317_v2_)[p0] if (p0) in (d_1317_v2_) else len(d_1319_v4_))
        nw198_[int(8)] = ((0) - ((self).f24) if p1 else (self).f24)
        nw198_[int(9)] = (self).f24
        nw198_[int(10)] = ((self).f24) - (len(_dafny.SeqWithoutIsStrInference([p1])))
        nw198_[int(11)] = default__.safeDivisionInt(len((d_1321_v6_).cf53), (self).f24)
        nw198_[int(12)] = len((d_1322_v7_).intersection(d_1322_v7_))
        nw198_[int(13)] = (self).f24
        nw198_[int(14)] = (self).f24
        nw198_[int(15)] = (self).f24
        d_1323_v8_ = nw198_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1323_v8_).length(0)):
            d_1324_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_1324_i0_)) and ((d_1324_i0_) < ((d_1323_v8_).length(0)))):
                (d_1323_v8_)[(d_1324_i0_)] = (d_1324_i0_) * (default__.safeModuloInt((self).f24, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spivrnx"))))))
        d_1325_v9_: C0
        nw199_ = C0()
        nw199_.ctor__(True)
        d_1325_v9_ = nw199_
        hi10_ = (0) - ((self).f24)
        for d_1326_i1_ in range((0) - (len(_dafny.Set({p1}))), hi10_):
            if (d_1325_v9_).f34:
                d_1318_v3_ = d_1318_v3_
                d_1327_v10_: _dafny.Set
                d_1327_v10_ = _dafny.Set({(self).f28})
                d_1328_v11_: _dafny.Map
                d_1328_v11_ = _dafny.Map({p0: p0})
                d_1329_v12_: _dafny.Seq
                d_1329_v12_ = _dafny.SeqWithoutIsStrInference([(d_1325_v9_).f34])
                d_1330_v13_: _dafny.Array
                nw200_ = _dafny.Array(None, 27)
                nw200_[int(0)] = p1
                nw200_[int(1)] = (d_1325_v9_).f34
                nw200_[int(2)] = (d_1325_v9_).f34
                nw200_[int(3)] = p1
                nw200_[int(4)] = (d_1325_v9_).f34
                nw200_[int(5)] = (d_1325_v9_).f34
                nw200_[int(6)] = default__.fm2(111, d_1326_i1_, d_1318_v3_, globalState)
                nw200_[int(7)] = (d_1325_v9_).f34
                nw200_[int(8)] = ((d_1328_v11_)[p1] if (p1) in (d_1328_v11_) else p0)
                nw200_[int(9)] = p1
                nw200_[int(10)] = (d_1325_v9_).f34
                nw200_[int(11)] = p0
                nw200_[int(12)] = p0
                nw200_[int(13)] = False
                nw200_[int(14)] = (d_1325_v9_).f34
                nw200_[int(15)] = True
                nw200_[int(16)] = False
                nw200_[int(17)] = p1
                nw200_[int(18)] = p0
                nw200_[int(19)] = not(True)
                nw200_[int(20)] = (d_1325_v9_).f34
                nw200_[int(21)] = (d_1325_v9_).f34
                nw200_[int(22)] = (d_1325_v9_).f34
                nw200_[int(23)] = (d_1329_v12_)[default__.safeIndex(len(_dafny.Map({(d_1325_v9_).f34: (d_1325_v9_).f34})), len(d_1329_v12_))]
                nw200_[int(24)] = (d_1325_v9_).f34
                nw200_[int(25)] = (d_1325_v9_).f34
                nw200_[int(26)] = (D6_DC17(p0)).cf28
                d_1330_v13_ = nw200_
                d_1331_v14_: _dafny.Map
                d_1331_v14_ = _dafny.Map({len(d_1327_v10_): d_1330_v13_})
                d_1331_v14_ = (d_1331_v14_).set(default__.fm0(951, -204, globalState), d_1330_v13_)
                (globalState).f2 = True
                d_1332_v15_: C0
                nw201_ = C0()
                nw201_.ctor__((d_1325_v9_).f34)
                d_1332_v15_ = nw201_
                (globalState).f1 = (self).f24
            elif True:
                (globalState).f18 = (self).f24
                index180_ = default__.safeIndex(694, ((self).f28).length(0))
                ((self).f28)[index180_] = default__.fm0((0) - (default__.fm0((self).f24, 572, globalState)), (self).f24, globalState)
                d_1333_v16_: _dafny.Array
                nw202_ = _dafny.Array(D4.default()(), 27)
                d_1333_v16_ = nw202_
                d_1334_v17_: _dafny.Map
                d_1334_v17_ = _dafny.Map({d_1333_v16_: (d_1325_v9_).f34})
                index181_ = default__.safeIndex(694, ((self).f28).length(0))
                ((self).f28)[index181_] = ((d_1326_i1_ if p0 else -994)) - (len(d_1334_v17_))
                d_1335_v18_: _dafny.Array
                def lambda57_(d_1336_v9_, d_1337_p0_):
                    def lambda58_(d_1338_i2_):
                        return ((d_1336_v9_).f34) or (d_1337_p0_)

                    return lambda58_

                init33_ = lambda57_(d_1325_v9_, p0)
                nw203_ = _dafny.Array(None, 17)
                for i0_33_ in range(nw203_.length(0)):
                    nw203_[i0_33_] = init33_(i0_33_)
                d_1335_v18_ = nw203_
                index182_ = default__.safeIndex(254, (d_1335_v18_).length(0))
                (d_1335_v18_)[index182_] = (d_1325_v9_).f34
                index183_ = default__.safeIndex(254, (d_1335_v18_).length(0))
                (d_1335_v18_)[index183_] = (d_1325_v9_).f34
                d_1335_v18_ = d_1335_v18_
                d_1339_v19_: _dafny.Array
                nw204_ = _dafny.Array(_dafny.Set({}), 14)
                d_1339_v19_ = nw204_
                d_1340_v20_: _dafny.Array
                nw205_ = _dafny.Array(_dafny.Map({}), 29)
                d_1340_v20_ = nw205_
                d_1341_v21_: _dafny.Map
                d_1341_v21_ = _dafny.Map({d_1318_v3_: d_1326_i1_})
                d_1342_v22_: T0
                nw206_ = C3()
                nw206_.ctor__(d_1339_v19_, d_1340_v20_, default__.safeModuloInt((0) - (-388), ((self).f28)[default__.safeIndex(694, ((self).f28).length(0))]), (((self).f25) + (((self).f25).set(default__.safeIndex(d_1326_i1_, len((self).f25)), d_1341_v21_))).set(default__.safeIndex((self).f24, len(((self).f25) + (((self).f25).set(default__.safeIndex(d_1326_i1_, len((self).f25)), d_1341_v21_)))), (d_1341_v21_).set(d_1318_v3_, ((self).f28)[default__.safeIndex(694, ((self).f28).length(0))])))
                d_1342_v22_ = nw206_
                d_1342_v22_ = d_1342_v22_
            d_1343_v23_: _dafny.Seq
            d_1343_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epg"))
            d_1344_v24_: _dafny.Map
            d_1344_v24_ = _dafny.Map({default__.safeModuloInt((self).f24, len(d_1343_v23_)): True})
            d_1345_v25_: _dafny.Map
            d_1345_v25_ = _dafny.Map({(d_1325_v9_).f34: d_1322_v7_})
            d_1344_v24_ = (d_1344_v24_).set((d_1326_i1_ if (d_1325_v9_).f34 else len(d_1345_v25_)), not ((d_1325_v9_).f34) or (not(p1)))
            r1 = 602
            d_1346_v26_: _dafny.MultiSet
            d_1346_v26_ = _dafny.MultiSet([(0) - ((self).f24)])
            d_1344_v24_ = (d_1344_v24_).set((d_1346_v26_).cardinality, False)
        d_1347_v27_: _dafny.Seq
        d_1347_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wlyo"))
        index184_ = default__.safeIndex(153, (d_1323_v8_).length(0))
        (d_1323_v8_)[index184_] = default__.safeModuloInt((self).f24, (0) - (len(d_1347_v27_)))
        d_1348_v28_: _dafny.Seq
        d_1348_v28_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1347_v27_))])
        index185_ = default__.safeIndex(153, (d_1323_v8_).length(0))
        (d_1323_v8_)[index185_] = (default__.fm0((self).f24, (self).f24, globalState)) * ((d_1348_v28_)[default__.safeIndex((self).f24, len(d_1348_v28_))])
        d_1349_v29_: _dafny.Map
        d_1349_v29_ = _dafny.Map({d_1318_v3_: False})
        index186_ = default__.safeIndex(153, (d_1323_v8_).length(0))
        (d_1323_v8_)[index186_] = default__.fm0((len(d_1349_v29_)) * ((self).f24), len(_dafny.SeqWithoutIsStrInference([d_1318_v3_ for d_1350_i3_ in range(default__.abs(148))])), globalState)
        d_1351_v30_: _dafny.Array
        nw207_ = _dafny.Array(_dafny.Seq({}), 26)
        d_1351_v30_ = nw207_
        d_1352_v31_: _dafny.Seq
        d_1352_v31_ = _dafny.SeqWithoutIsStrInference([p0])
        index187_ = default__.safeIndex(571, (d_1351_v30_).length(0))
        (d_1351_v30_)[index187_] = (d_1352_v31_ if (d_1325_v9_).f34 else d_1352_v31_)
        index188_ = default__.safeIndex(571, (d_1351_v30_).length(0))
        (d_1351_v30_)[index188_] = (d_1352_v31_) + (d_1352_v31_)
        r0 = (d_1323_v8_)[default__.safeIndex(153, (d_1323_v8_).length(0))]
        r1 = (self).f24
        d_1353_v32_: _dafny.Seq
        d_1353_v32_ = _dafny.SeqWithoutIsStrInference([d_1348_v28_, d_1348_v28_])
        d_1354_v33_: D4
        d_1354_v33_ = D4_DC12(_dafny.SeqWithoutIsStrInference([(d_1323_v8_)[default__.safeIndex(153, (d_1323_v8_).length(0))], (d_1323_v8_)[default__.safeIndex(153, (d_1323_v8_).length(0))], (self).f24, (self).f24]))
        d_1355_v34_: _dafny.Seq
        d_1355_v34_ = _dafny.SeqWithoutIsStrInference([d_1348_v28_, (d_1353_v32_)[default__.safeIndex((self).f24, len(d_1353_v32_))], (d_1354_v33_).cf21, (d_1348_v28_) + (d_1348_v28_)])
        r2 = (d_1353_v32_)[default__.safeIndex((self).f24, len(d_1353_v32_))]
        r3 = p0
        return r0, r1, r2, r3

    @property
    def f28(self):
        return self._f28

class C5(T0):
    def  __init__(self):
        self._f24: int = int(0)
        self._f25: _dafny.Seq = _dafny.Seq({})
        self.f27: int = int(0)
        self._f26: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25
    def ctor__(self, f26, f27, f24, f25):
        (self)._f26 = f26
        (self).f27 = f27
        (self)._f24 = f24
        (self)._f25 = f25

    def m1(self, globalState):
        r0: bool = False
        hi11_ = self.f27
        for d_1356_i0_ in range(default__.safeModuloInt(self.f27, self.f27), hi11_):
            (globalState).f14 = ((self).f24) - (self.f27)
            if default__.fm2((self.f27) + (-676), (self).f24, default__.fm3(globalState), globalState):
                d_1357_v0_: _dafny.Array
                def lambda59_(d_1358_i1_):
                    return not(not(True))

                init34_ = lambda59_
                nw208_ = _dafny.Array(None, 15)
                for i0_34_ in range(nw208_.length(0)):
                    nw208_[i0_34_] = init34_(i0_34_)
                d_1357_v0_ = nw208_
                d_1359_v1_: bool
                d_1359_v1_ = False
                index189_ = default__.safeIndex(568, (d_1357_v0_).length(0))
                (d_1357_v0_)[index189_] = d_1359_v1_
                index190_ = default__.safeIndex(568, (d_1357_v0_).length(0))
                (d_1357_v0_)[index190_] = not (d_1359_v1_) or (d_1359_v1_)
                d_1360_v2_: _dafny.Array
                nw209_ = _dafny.Array(_dafny.Seq({}), 7)
                d_1360_v2_ = nw209_
                d_1361_v3_: _dafny.Array
                nw210_ = _dafny.Array(int(0), 4)
                d_1361_v3_ = nw210_
                d_1362_v4_: T0
                nw211_ = C4()
                nw211_.ctor__(d_1361_v3_, len(_dafny.SeqWithoutIsStrInference([self.f27])), (self).f25)
                d_1362_v4_ = nw211_
                index191_ = default__.safeIndex(828, (d_1360_v2_).length(0))
                (d_1360_v2_)[index191_] = _dafny.SeqWithoutIsStrInference([d_1362_v4_])
                d_1363_v5_: _dafny.MultiSet
                d_1363_v5_ = _dafny.MultiSet([d_1359_v1_])
                d_1364_v6_: _dafny.Seq
                d_1364_v6_ = _dafny.SeqWithoutIsStrInference([d_1362_v4_, d_1362_v4_])
                index192_ = default__.safeIndex(828, (d_1360_v2_).length(0))
                (d_1360_v2_)[index192_] = ((d_1364_v6_) + (d_1364_v6_) if (d_1363_v5_).isdisjoint(d_1363_v5_) else d_1364_v6_)
                d_1365_v7_: _dafny.Array
                nw212_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                d_1365_v7_ = nw212_
                d_1366_v8_: str
                d_1366_v8_ = _dafny.CodePoint('d')
                index193_ = default__.safeIndex(466, (d_1365_v7_).length(0))
                (d_1365_v7_)[index193_] = d_1366_v8_
                index194_ = default__.safeIndex(466, (d_1365_v7_).length(0))
                (d_1365_v7_)[index194_] = d_1366_v8_
                d_1367_v9_: _dafny.Map
                d_1367_v9_ = _dafny.Map({(d_1362_v4_).f24: (d_1362_v4_).f24})
                d_1368_v10_: _dafny.Seq
                d_1368_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iccnie"))
                d_1369_v11_: _dafny.Seq
                d_1369_v11_ = _dafny.SeqWithoutIsStrInference([len(d_1368_v10_)])
                d_1370_v12_: _dafny.MultiSet
                d_1370_v12_ = _dafny.MultiSet([((d_1367_v9_)[(d_1362_v4_).f24] if ((d_1362_v4_).f24) in (d_1367_v9_) else (d_1362_v4_).f24), (d_1369_v11_)[default__.safeIndex((self).f24, len(d_1369_v11_))]])
                d_1371_v13_: int
                out31_: int
                out31_ = default__.m0((0) - (default__.safeDivisionInt(d_1356_i0_, (d_1370_v12_).cardinality)), default__.safeDivisionInt(d_1356_i0_, default__.fm0((0) - (len(d_1368_v10_)), d_1356_i0_, globalState)), globalState)
                d_1371_v13_ = out31_
                index195_ = default__.safeIndex(568, (d_1357_v0_).length(0))
                (d_1357_v0_)[index195_] = d_1359_v1_
            elif True:
                (globalState).f9 = self.f27
                d_1372_v14_: bool
                d_1372_v14_ = True
                d_1373_v15_: C0
                nw213_ = C0()
                nw213_.ctor__(d_1372_v14_)
                d_1373_v15_ = nw213_
                d_1374_v16_: D6
                d_1374_v16_ = D6_DC18(not((d_1373_v15_).f34))
                d_1375_v17_: D6
                d_1375_v17_ = D6_DC19(d_1374_v16_)
                d_1376_v18_: _dafny.Seq
                d_1376_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rflync"))
                d_1377_v19_: _dafny.Map
                d_1377_v19_ = _dafny.Map({d_1376_v18_: D6_DC18((d_1373_v15_).f34)})
                pat_let_tv16_ = d_1374_v16_
                pat_let_tv17_ = d_1377_v19_
                pat_let_tv18_ = d_1376_v18_
                pat_let_tv19_ = d_1376_v18_
                pat_let_tv20_ = d_1377_v19_
                pat_let_tv21_ = d_1374_v16_
                pat_let_tv22_ = d_1374_v16_
                d_1378_v20_: _dafny.Array
                nw214_ = _dafny.Array(None, 23)
                nw214_[int(0)] = d_1375_v17_
                nw214_[int(1)] = d_1375_v17_
                nw214_[int(2)] = d_1375_v17_
                nw214_[int(3)] = D6_DC19(d_1374_v16_)
                nw214_[int(4)] = d_1375_v17_
                nw214_[int(5)] = d_1375_v17_
                nw214_[int(6)] = d_1375_v17_
                nw214_[int(7)] = D6_DC19(d_1374_v16_)
                nw214_[int(8)] = d_1375_v17_
                nw214_[int(9)] = d_1375_v17_
                nw214_[int(10)] = D6_DC19(d_1374_v16_)
                nw214_[int(11)] = d_1375_v17_
                def iife81_(_pat_let20_0):
                    def iife82_(d_1379_dt__update__tmp_h0_):
                        def iife83_(_pat_let21_0):
                            def iife84_(d_1380_dt__update_hcf30_h0_):
                                return D6_DC19(d_1380_dt__update_hcf30_h0_)
                            return iife84_(_pat_let21_0)
                        return iife83_(pat_let_tv16_)
                    return iife82_(_pat_let20_0)
                nw214_[int(12)] = iife81_(d_1375_v17_)
                nw214_[int(13)] = d_1375_v17_
                nw214_[int(14)] = d_1375_v17_
                nw214_[int(15)] = (d_1375_v17_ if False else d_1375_v17_)
                nw214_[int(16)] = d_1375_v17_
                nw214_[int(17)] = d_1375_v17_
                nw214_[int(18)] = d_1375_v17_
                def iife85_(_pat_let22_0):
                    def iife86_(d_1381_dt__update__tmp_h1_):
                        def iife87_(_pat_let23_0):
                            def iife88_(d_1382_dt__update_hcf30_h1_):
                                return D6_DC19(d_1382_dt__update_hcf30_h1_)
                            return iife88_(_pat_let23_0)
                        return iife87_(((pat_let_tv17_)[pat_let_tv18_] if (pat_let_tv19_) in (pat_let_tv20_) else pat_let_tv21_))
                    return iife86_(_pat_let22_0)
                nw214_[int(19)] = iife85_(D6_DC19(d_1374_v16_))
                def iife89_(_pat_let24_0):
                    def iife90_(d_1383_dt__update__tmp_h2_):
                        def iife91_(_pat_let25_0):
                            def iife92_(d_1384_dt__update_hcf30_h2_):
                                return D6_DC19(d_1384_dt__update_hcf30_h2_)
                            return iife92_(_pat_let25_0)
                        return iife91_(pat_let_tv22_)
                    return iife90_(_pat_let24_0)
                nw214_[int(20)] = iife89_(d_1375_v17_)
                nw214_[int(21)] = d_1375_v17_
                nw214_[int(22)] = d_1375_v17_
                d_1378_v20_ = nw214_
                d_1378_v20_ = d_1378_v20_
                d_1385_v21_: _dafny.Seq
                d_1385_v21_ = _dafny.SeqWithoutIsStrInference([(d_1373_v15_).f34])
                (globalState).f21 = (d_1385_v21_ if (d_1373_v15_).f34 else (d_1385_v21_) + (d_1385_v21_))
                d_1386_v22_: _dafny.Array
                nw215_ = _dafny.Array(int(0), 10)
                d_1386_v22_ = nw215_
                d_1387_v23_: _dafny.MultiSet
                d_1387_v23_ = _dafny.MultiSet([d_1386_v22_, d_1386_v22_, d_1386_v22_, d_1386_v22_])
                d_1387_v23_ = (d_1387_v23_).set(d_1386_v22_, default__.abs(len((_dafny.Set({d_1386_v22_}) if (d_1373_v15_).f34 else _dafny.Set({d_1386_v22_, d_1386_v22_})))))
            (globalState).f1 = default__.safeModuloInt(891, self.f27)
            d_1388_v24_: bool
            d_1388_v24_ = False
            if d_1388_v24_:
                d_1388_v24_ = d_1388_v24_
                d_1389_v25_: bool
                out32_: bool
                out32_ = (self).m2(globalState)
                d_1389_v25_ = out32_
                d_1390_v26_: _dafny.Array
                def lambda60_(d_1391_i0_):
                    def lambda61_(d_1392_i2_):
                        return (d_1392_i2_) + (d_1391_i0_)

                    return lambda61_

                init35_ = lambda60_(d_1356_i0_)
                nw216_ = _dafny.Array(None, 18)
                for i0_35_ in range(nw216_.length(0)):
                    nw216_[i0_35_] = init35_(i0_35_)
                d_1390_v26_ = nw216_
                index196_ = default__.safeIndex(878, (d_1390_v26_).length(0))
                (d_1390_v26_)[index196_] = self.f27
                index197_ = default__.safeIndex(878, (d_1390_v26_).length(0))
                (d_1390_v26_)[index197_] = d_1356_i0_
                d_1393_v27_: str
                d_1393_v27_ = _dafny.CodePoint('o')
                d_1394_v28_: _dafny.Map
                d_1394_v28_ = _dafny.Map({d_1393_v27_: d_1356_i0_})
                d_1395_v29_: C4
                nw217_ = C4()
                nw217_.ctor__(d_1390_v26_, (d_1390_v26_)[default__.safeIndex(878, (d_1390_v26_).length(0))], (((self).f25) + ((self).f25)).set(default__.safeIndex(-423, len(((self).f25) + ((self).f25))), d_1394_v28_))
                d_1395_v29_ = nw217_
                (globalState).f9 = (self).f24
            elif True:
                d_1396_v30_: _dafny.Map
                d_1396_v30_ = _dafny.Map({d_1388_v24_: d_1356_i0_})
                d_1397_v31_: _dafny.Array
                nw218_ = _dafny.Array(None, 2)
                nw218_[int(0)] = self.f27
                nw218_[int(1)] = ((d_1396_v30_)[d_1388_v24_] if (d_1388_v24_) in (d_1396_v30_) else (0) - (self.f27))
                d_1397_v31_ = nw218_
                d_1398_v32_: T0
                nw219_ = C4()
                nw219_.ctor__(d_1397_v31_, (self).f24, (self).f25)
                d_1398_v32_ = nw219_
                d_1398_v32_ = d_1398_v32_
                rhs235_ = d_1388_v24_
                rhs236_ = default__.safeModuloInt(self.f27, d_1356_i0_)
                rhs237_ = (False if d_1388_v24_ else False)
                lhs188_ = globalState
                lhs189_ = globalState
                lhs190_ = globalState
                lhs188_.f2 = rhs235_
                lhs189_.f18 = rhs236_
                lhs190_.f2 = rhs237_
                d_1399_v33_: _dafny.Seq
                d_1399_v33_ = _dafny.SeqWithoutIsStrInference([d_1388_v24_])
                d_1400_v34_: _dafny.MultiSet
                d_1400_v34_ = _dafny.MultiSet([len(d_1399_v33_), (self).f24, d_1356_i0_, ((d_1398_v32_).f24) * ((self).f24), (564) + (self.f27)])
                d_1400_v34_ = d_1400_v34_
                (globalState).f17 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbquywj"))
                d_1401_v35_: C4
                nw220_ = C4()
                nw220_.ctor__(d_1397_v31_, (d_1398_v32_).f24, _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('u'): (self).f24}) for d_1402_i3_ in range(default__.abs(462))]))
                d_1401_v35_ = nw220_
                d_1403_v36_: _dafny.Map
                d_1403_v36_ = _dafny.Map({964: d_1401_v35_})
                d_1403_v36_ = d_1403_v36_
        d_1404_v37_: bool
        d_1404_v37_ = False
        d_1405_v38_: _dafny.Seq
        d_1405_v38_ = _dafny.SeqWithoutIsStrInference([(self).f24])
        pat_let_tv23_ = d_1404_v37_
        pat_let_tv24_ = d_1405_v38_
        d_1406_v39_: _dafny.Map
        def iife93_(_pat_let26_0):
            def iife94_(d_1407_dt__update__tmp_h3_):
                def iife95_(_pat_let27_0):
                    def iife96_(d_1408_dt__update_hcf20_h0_):
                        def iife97_(_pat_let28_0):
                            def iife98_(d_1409_dt__update_hcf18_h0_):
                                return D3_DC11(d_1409_dt__update_hcf18_h0_, (d_1407_dt__update__tmp_h3_).cf19, d_1408_dt__update_hcf20_h0_)
                            return iife98_(_pat_let28_0)
                        return iife97_(len(pat_let_tv24_))
                    return iife96_(_pat_let27_0)
                return iife95_(pat_let_tv23_)
            return iife94_(_pat_let26_0)
        d_1406_v39_ = _dafny.Map({d_1404_v37_: iife93_(D3_DC11(self.f27, d_1404_v37_, d_1404_v37_))})
        d_1410_v40_: D12
        d_1410_v40_ = D12_DC32(d_1406_v39_)
        source15_ = d_1410_v40_
        if source15_.is_DC33:
            d_1411___mcc_h0_ = source15_.cf54
            d_1412___mcc_h1_ = source15_.cf55
            d_1413_cf55_ = d_1412___mcc_h1_
            d_1414_cf54_ = d_1411___mcc_h0_
            d_1415_v41_: _dafny.Set
            d_1415_v41_ = _dafny.Set({self.f27})
            d_1413_cf55_ = not(not(((_dafny.Set({self.f27, (self).f24})) - (d_1415_v41_)).issubset(d_1415_v41_)))
            d_1416_v42_: _dafny.Seq
            d_1416_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdwmjxld"))
            r0 = default__.fm2((len(d_1416_v42_)) * ((self).f24), (self.f27) * (891), d_1414_cf54_, globalState)
            d_1417_v43_: _dafny.Array
            nw221_ = _dafny.Array(int(0), 23)
            d_1417_v43_ = nw221_
            index198_ = default__.safeIndex(655, (d_1417_v43_).length(0))
            (d_1417_v43_)[index198_] = self.f27
            d_1418_v44_: _dafny.Map
            d_1418_v44_ = _dafny.Map({d_1413_cf55_: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1419_i4_ in range(default__.abs(114))]))})
            index199_ = default__.safeIndex(655, (d_1417_v43_).length(0))
            (d_1417_v43_)[index199_] = (len(d_1418_v44_)) + (default__.safeModuloInt((self).f24, self.f27))
            (self).f27 = (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f24 for d_1420_i5_ in range(default__.abs(613))])))
        elif source15_.is_DC32:
            d_1421___mcc_h2_ = source15_.cf53
            d_1422_cf53_ = d_1421___mcc_h2_
            d_1423_v45_: _dafny.Seq
            d_1423_v45_ = _dafny.SeqWithoutIsStrInference([True, d_1404_v37_])
            d_1424_v46_: _dafny.Set
            d_1424_v46_ = _dafny.Set({d_1423_v45_})
            r0 = (d_1424_v46_).ispropersubset(d_1424_v46_)
            d_1425_v47_: _dafny.Seq
            d_1425_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_1405_v38_)])
            d_1426_v48_: C1
            nw222_ = C1()
            nw222_.ctor__(d_1425_v47_, self.f27, self.f27, _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('x'): (self).f24}) for d_1427_i6_ in range(default__.abs(-502))]))
            d_1426_v48_ = nw222_
            d_1428_v49_: str
            d_1428_v49_ = _dafny.CodePoint('w')
            d_1429_v51_: _dafny.MultiSet
            d_1429_v51_ = _dafny.MultiSet([(self).f24])
            d_1430_v52_: D13
            d_1430_v52_ = D13_DC35(d_1429_v51_)
            d_1431_v53_: _dafny.Seq
            d_1431_v53_ = _dafny.SeqWithoutIsStrInference([D3_DC11(950, d_1404_v37_, default__.fm2(543, self.f27, d_1428_v49_, globalState))])
            d_1432_v54_: _dafny.MultiSet
            d_1432_v54_ = _dafny.MultiSet([d_1404_v37_, d_1404_v37_])
            d_1433_v55_: _dafny.Array
            nw223_ = _dafny.Array(None, 16)
            nw223_[int(0)] = not(True)
            nw223_[int(1)] = d_1404_v37_
            nw223_[int(2)] = d_1404_v37_
            nw223_[int(3)] = default__.fm2(d_1426_v48_.f33, d_1426_v48_.f33, d_1428_v49_, globalState)
            nw223_[int(4)] = d_1404_v37_
            def iife99_():
                coll41_ = _dafny.Set()
                compr_41_: int
                for compr_41_ in _dafny.IntegerRange(794, -319):
                    d_1434_v50_: int = compr_41_
                    if ((794) <= (d_1434_v50_)) and ((d_1434_v50_) < (-319)):
                        coll41_ = coll41_.union(_dafny.Set([default__.safeDivisionInt(d_1434_v50_, (self).f24)]))
                return _dafny.Set(coll41_)
            nw223_[int(5)] = (_dafny.Set({d_1426_v48_.f33, d_1426_v48_.f33})).isdisjoint(iife99_()
            )
            nw223_[int(6)] = True
            nw223_[int(7)] = True
            nw223_[int(8)] = ((d_1430_v52_).cf57).isdisjoint(d_1429_v51_)
            nw223_[int(9)] = ((d_1429_v51_).set(len(d_1431_v53_), default__.abs(d_1426_v48_.f33))).issubset(d_1429_v51_)
            nw223_[int(10)] = (d_1426_v48_.f33) != (self.f27)
            nw223_[int(11)] = (d_1432_v54_).isdisjoint(d_1432_v54_)
            nw223_[int(12)] = not(default__.fm2((d_1432_v54_).cardinality, self.f27, d_1428_v49_, globalState))
            nw223_[int(13)] = d_1404_v37_
            nw223_[int(14)] = (d_1423_v45_)[default__.safeIndex((0) - (self.f27), len(d_1423_v45_))]
            nw223_[int(15)] = d_1404_v37_
            d_1433_v55_ = nw223_
            index200_ = default__.safeIndex(205, (d_1433_v55_).length(0))
            (d_1433_v55_)[index200_] = (self.f27) <= ((self).f24)
            index201_ = default__.safeIndex(205, (d_1433_v55_).length(0))
            (d_1433_v55_)[index201_] = d_1404_v37_
            d_1435_v56_: _dafny.Map
            d_1435_v56_ = _dafny.Map({d_1428_v49_: d_1404_v37_})
            index202_ = default__.safeIndex(205, (d_1433_v55_).length(0))
            rhs238_ = default__.safeDivisionInt(self.f27, default__.safeDivisionInt(d_1426_v48_.f33, 541))
            rhs239_ = (True) and ((d_1404_v37_ if (d_1433_v55_)[default__.safeIndex(205, (d_1433_v55_).length(0))] else d_1404_v37_))
            rhs240_ = (d_1435_v56_) != (d_1435_v56_)
            rhs241_ = (0) - (default__.fm0(default__.safeDivisionInt((self).f24, len(d_1405_v38_)), self.f27, globalState))
            rhs242_ = d_1404_v37_
            lhs191_ = globalState
            lhs192_ = globalState
            lhs193_ = d_1426_v48_
            lhs194_ = d_1433_v55_
            lhs195_ = default__.safeIndex(205, (d_1433_v55_).length(0))
            lhs191_.f14 = rhs238_
            r0 = rhs239_
            lhs192_.f2 = rhs240_
            lhs193_.f33 = rhs241_
            lhs194_[lhs195_] = rhs242_
        elif True:
            d_1436___mcc_h3_ = source15_.cf56
            d_1437_cf56_ = d_1436___mcc_h3_
            d_1438_v57_: _dafny.Array
            nw224_ = _dafny.Array(False, 3)
            d_1438_v57_ = nw224_
            index203_ = default__.safeIndex(383, (d_1438_v57_).length(0))
            (d_1438_v57_)[index203_] = False
            d_1439_v58_: _dafny.Seq
            d_1439_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kltkdmnl"))
            d_1440_v59_: str
            d_1440_v59_ = _dafny.CodePoint('h')
            d_1441_v60_: D8
            d_1441_v60_ = D8_DC25(self.f27, d_1404_v37_, d_1404_v37_)
            d_1442_v61_: _dafny.Map
            d_1442_v61_ = _dafny.Map({d_1440_v59_: (d_1441_v60_).cf40})
            d_1443_v62_: _dafny.Seq
            d_1443_v62_ = _dafny.SeqWithoutIsStrInference([d_1404_v37_])
            index204_ = default__.safeIndex(383, (d_1438_v57_).length(0))
            rhs243_ = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1444_i7_ in range(default__.abs(4))])) + (d_1439_v58_)) == (d_1439_v58_)
            rhs244_ = ((d_1442_v61_)[d_1440_v59_] if (d_1440_v59_) in (d_1442_v61_) else d_1404_v37_)
            rhs245_ = d_1443_v62_
            lhs196_ = d_1438_v57_
            lhs197_ = default__.safeIndex(383, (d_1438_v57_).length(0))
            lhs198_ = globalState
            lhs199_ = globalState
            lhs196_[lhs197_] = rhs243_
            lhs198_.f2 = rhs244_
            lhs199_.f20 = rhs245_
            index205_ = default__.safeIndex(383, (d_1438_v57_).length(0))
            (d_1438_v57_)[index205_] = (d_1438_v57_)[default__.safeIndex(383, (d_1438_v57_).length(0))]
            d_1445_v63_: _dafny.MultiSet
            d_1445_v63_ = _dafny.MultiSet([d_1404_v37_])
            d_1446_v64_: _dafny.Map
            d_1446_v64_ = _dafny.Map({d_1438_v57_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekhhngqe")))})
            d_1447_v65_: _dafny.Seq
            d_1447_v65_ = _dafny.SeqWithoutIsStrInference([d_1438_v57_, d_1438_v57_, d_1438_v57_, d_1438_v57_, d_1438_v57_])
            d_1448_v66_: _dafny.Array
            nw225_ = _dafny.Array(None, 24)
            nw225_[int(0)] = default__.safeDivisionInt(-453, self.f27)
            nw225_[int(1)] = (0) - (self.f27)
            nw225_[int(2)] = default__.fm0((0) - ((d_1445_v63_).cardinality), len(d_1439_v58_), globalState)
            nw225_[int(3)] = self.f27
            nw225_[int(4)] = 685
            nw225_[int(5)] = len(d_1446_v64_)
            nw225_[int(6)] = (len(d_1439_v58_)) + (self.f27)
            nw225_[int(7)] = self.f27
            nw225_[int(8)] = (len(d_1405_v38_)) * ((self).f24)
            nw225_[int(9)] = self.f27
            nw225_[int(10)] = self.f27
            nw225_[int(11)] = default__.safeDivisionInt(self.f27, -1)
            nw225_[int(12)] = self.f27
            nw225_[int(13)] = len((d_1439_v58_) + (d_1439_v58_))
            nw225_[int(14)] = self.f27
            nw225_[int(15)] = self.f27
            nw225_[int(16)] = len((d_1447_v65_).set(default__.safeIndex(len(d_1443_v62_), len(d_1447_v65_)), d_1438_v57_))
            nw225_[int(17)] = self.f27
            nw225_[int(18)] = (self).f24
            nw225_[int(19)] = self.f27
            nw225_[int(20)] = len(d_1439_v58_)
            nw225_[int(21)] = 400
            nw225_[int(22)] = (0) - ((len(d_1405_v38_)) - ((self).f24))
            nw225_[int(23)] = (662 if d_1404_v37_ else default__.fm0(-584, (self).f24, globalState))
            d_1448_v66_ = nw225_
            d_1449_v67_: _dafny.Seq
            d_1449_v67_ = _dafny.SeqWithoutIsStrInference([d_1448_v66_, d_1448_v66_, (d_1448_v66_ if d_1404_v37_ else d_1448_v66_)])
            d_1448_v66_ = (d_1449_v67_)[default__.safeIndex((self).f24, len(d_1449_v67_))]
            (globalState).f2 = d_1404_v37_
        d_1450_v68_: _dafny.MultiSet
        d_1450_v68_ = _dafny.MultiSet([not(d_1404_v37_)])
        d_1451_v69_: _dafny.Seq
        d_1451_v69_ = _dafny.SeqWithoutIsStrInference([d_1404_v37_, d_1404_v37_, d_1404_v37_])
        d_1452_v70_: _dafny.MultiSet
        d_1452_v70_ = _dafny.MultiSet([(self).f24, self.f27, (self).f24, self.f27, self.f27])
        d_1453_v71_: C1
        nw226_ = C1()
        nw226_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(d_1451_v69_)]), d_1452_v70_]), (self).f24, self.f27, _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('y'): self.f27}) for d_1454_i8_ in range(default__.abs(382))]))
        d_1453_v71_ = nw226_
        d_1455_v72_: _dafny.Set
        d_1455_v72_ = _dafny.Set({d_1453_v71_, d_1453_v71_})
        d_1456_v73_: _dafny.Array
        nw227_ = _dafny.Array(int(0), 20)
        d_1456_v73_ = nw227_
        d_1457_v75_: str
        d_1457_v75_ = _dafny.CodePoint('k')
        d_1458_v76_: _dafny.Map
        d_1458_v76_ = _dafny.Map({d_1457_v75_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uoejhqk"))})
        d_1459_v77_: T0
        nw228_ = C4()
        def iife100_():
            coll42_ = _dafny.Map()
            compr_42_: str
            for compr_42_ in (d_1458_v76_).keys.Elements:
                d_1460_v74_: str = compr_42_
                if (d_1460_v74_) in (d_1458_v76_):
                    coll42_[d_1460_v74_] = d_1453_v71_.f33
            return _dafny.Map(coll42_)
        nw228_.ctor__(d_1456_v73_, self.f27, _dafny.SeqWithoutIsStrInference([iife100_()
        ]))
        d_1459_v77_ = nw228_
        d_1461_v78_: _dafny.Set
        d_1461_v78_ = _dafny.Set({d_1459_v77_})
        d_1462_v79_: _dafny.Map
        d_1462_v79_ = _dafny.Map({(self).f24: d_1404_v37_})
        d_1463_v80_: _dafny.Seq
        d_1463_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuvexp"))
        d_1464_v81_: _dafny.Array
        nw229_ = _dafny.Array(None, 23)
        nw229_[int(0)] = (0) - ((self).f24)
        nw229_[int(1)] = (self).f24
        nw229_[int(2)] = self.f27
        nw229_[int(3)] = self.f27
        nw229_[int(4)] = self.f27
        nw229_[int(5)] = (self).f24
        nw229_[int(6)] = (self).f24
        nw229_[int(7)] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bobuyv"))))
        nw229_[int(8)] = (self).f24
        nw229_[int(9)] = (0) - ((self).f24)
        nw229_[int(10)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))
        nw229_[int(11)] = -676
        nw229_[int(12)] = default__.fm0((self).f24, self.f27, globalState)
        nw229_[int(13)] = default__.fm0(len(d_1455_v72_), self.f27, globalState)
        nw229_[int(14)] = self.f27
        nw229_[int(15)] = len(d_1461_v78_)
        nw229_[int(16)] = len(d_1462_v79_)
        nw229_[int(17)] = (self).f24
        nw229_[int(18)] = (self).f24
        nw229_[int(19)] = (self).f24
        nw229_[int(20)] = d_1453_v71_.f33
        nw229_[int(21)] = (self).f24
        nw229_[int(22)] = len(d_1463_v80_)
        d_1464_v81_ = nw229_
        d_1465_v82_: _dafny.Map
        d_1465_v82_ = _dafny.Map({default__.safeModuloInt((d_1450_v68_).cardinality, 461): d_1464_v81_})
        d_1466_v83_: _dafny.Array
        nw230_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
        d_1466_v83_ = nw230_
        index206_ = default__.safeIndex(203, (d_1466_v83_).length(0))
        (d_1466_v83_)[index206_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1467_i9_ in range(default__.abs(991))])
        index207_ = default__.safeIndex(203, (d_1466_v83_).length(0))
        rhs246_ = False
        rhs247_ = (((self).f26)[d_1404_v37_] if (d_1404_v37_) in ((self).f26) else (d_1451_v69_) < (d_1451_v69_))
        rhs248_ = (d_1465_v82_).set(len(d_1451_v69_), d_1456_v73_)
        rhs249_ = d_1404_v37_
        rhs250_ = d_1463_v80_
        lhs200_ = globalState
        lhs201_ = d_1466_v83_
        lhs202_ = default__.safeIndex(203, (d_1466_v83_).length(0))
        lhs200_.f2 = rhs246_
        d_1404_v37_ = rhs247_
        d_1465_v82_ = rhs248_
        r0 = rhs249_
        lhs201_[lhs202_] = rhs250_
        d_1456_v73_ = d_1464_v81_
        d_1468_v84_: _dafny.Map
        d_1468_v84_ = _dafny.Map({False: d_1453_v71_.f33})
        d_1469_v85_: D4
        d_1469_v85_ = D4_DC13(self.f27, d_1468_v84_, d_1463_v80_)
        d_1470_v86_: D4
        d_1470_v86_ = D4_DC14(d_1469_v85_)
        d_1471_v87_: D4
        d_1471_v87_ = D4_DC14(d_1470_v86_)
        d_1472_v88_: D4
        d_1472_v88_ = D4_DC14(d_1470_v86_)
        d_1473_v89_: _dafny.Map
        d_1473_v89_ = _dafny.Map({d_1472_v88_: d_1463_v80_})
        d_1474_v90_: D7
        d_1474_v90_ = D7_DC20(d_1473_v89_)
        d_1475_v91_: _dafny.Map
        d_1475_v91_ = _dafny.Map({d_1474_v90_: (d_1466_v83_)[default__.safeIndex(203, (d_1466_v83_).length(0))]})
        d_1475_v91_ = (d_1475_v91_).set(d_1474_v90_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aohh")))
        d_1476_v92_: _dafny.Array
        nw231_ = _dafny.Array(None, 5)
        d_1476_v92_ = nw231_
        d_1477_v93_: C4
        nw232_ = C4()
        nw232_.ctor__(d_1464_v81_, (d_1459_v77_).f24, (self).f25)
        d_1477_v93_ = nw232_
        d_1478_v94_: _dafny.Seq
        d_1478_v94_ = _dafny.SeqWithoutIsStrInference([d_1477_v93_])
        d_1479_v95_: _dafny.Map
        d_1479_v95_ = _dafny.Map({d_1404_v37_: d_1451_v69_})
        index208_ = default__.safeIndex(655, (d_1476_v92_).length(0))
        (d_1476_v92_)[index208_] = (d_1478_v94_)[default__.safeIndex(len(((d_1479_v95_)[d_1404_v37_] if (d_1404_v37_) in (d_1479_v95_) else d_1451_v69_)), len(d_1478_v94_))]
        index209_ = default__.safeIndex(655, (d_1476_v92_).length(0))
        (d_1476_v92_)[index209_] = (d_1478_v94_)[default__.safeIndex(((self).f24) * (492), len(d_1478_v94_))]
        r0 = ((d_1453_v71_).fm10(globalState)) or (not(d_1404_v37_))
        return r0

    def m2(self, globalState):
        r0: bool = False
        d_1480_v0_: bool
        d_1480_v0_ = True
        if not(d_1480_v0_):
            (globalState).f9 = (self).f24
            d_1481_v1_: _dafny.Map
            d_1481_v1_ = _dafny.Map({d_1480_v0_: len((self).f26)})
            d_1482_v2_: D4
            d_1482_v2_ = D4_DC13((self).f24, d_1481_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qiphwxsf")))
            if ((d_1482_v2_).cf22) == (self.f27):
                d_1483_v3_: _dafny.Array
                nw233_ = _dafny.Array(False, 28)
                d_1483_v3_ = nw233_
                index210_ = default__.safeIndex(372, (d_1483_v3_).length(0))
                (d_1483_v3_)[index210_] = d_1480_v0_
                index211_ = default__.safeIndex(372, (d_1483_v3_).length(0))
                (d_1483_v3_)[index211_] = d_1480_v0_
                d_1484_v4_: str
                d_1484_v4_ = _dafny.CodePoint('f')
                d_1485_v5_: C0
                nw234_ = C0()
                nw234_.ctor__(d_1480_v0_)
                d_1485_v5_ = nw234_
                d_1486_v6_: _dafny.Set
                d_1486_v6_ = _dafny.Set({self.f27})
                d_1487_v7_: _dafny.MultiSet
                d_1487_v7_ = _dafny.MultiSet([len(d_1486_v6_)])
                d_1488_v8_: _dafny.Map
                d_1488_v8_ = _dafny.Map({d_1484_v4_: (d_1487_v7_).cardinality})
                d_1489_v9_: T0
                nw235_ = C2()
                nw235_.ctor__((d_1485_v5_).f34, (self).f24, _dafny.SeqWithoutIsStrInference([d_1488_v8_]))
                d_1489_v9_ = nw235_
                d_1490_v10_: _dafny.Array
                nw236_ = _dafny.Array(int(0), 6)
                d_1490_v10_ = nw236_
                d_1491_v11_: _dafny.Map
                d_1491_v11_ = _dafny.Map({d_1489_v9_: d_1490_v10_})
                rhs251_ = (d_1489_v9_) in (d_1491_v11_)
                rhs252_ = default__.fm3(globalState)
                rhs253_ = d_1485_v5_
                lhs203_ = globalState
                lhs203_.f2 = rhs251_
                d_1484_v4_ = rhs252_
                d_1485_v5_ = rhs253_
                (globalState).f14 = (d_1489_v9_).f24
                d_1492_v12_: _dafny.MultiSet
                d_1492_v12_ = _dafny.MultiSet([(d_1483_v3_)[default__.safeIndex(372, (d_1483_v3_).length(0))]])
                d_1493_v13_: _dafny.MultiSet
                d_1493_v13_ = _dafny.MultiSet([(d_1492_v12_) - (d_1492_v12_)])
                d_1493_v13_ = (d_1493_v13_).set(d_1492_v12_, default__.abs((self).f24))
                d_1494_v14_: _dafny.Array
                nw237_ = _dafny.Array(None, 9)
                nw237_[int(0)] = d_1486_v6_
                nw237_[int(1)] = d_1486_v6_
                nw237_[int(2)] = d_1486_v6_
                nw237_[int(3)] = _dafny.Set({(d_1489_v9_).f24})
                nw237_[int(4)] = d_1486_v6_
                nw237_[int(5)] = d_1486_v6_
                nw237_[int(6)] = d_1486_v6_
                nw237_[int(7)] = d_1486_v6_
                nw237_[int(8)] = _dafny.Set({(self).f24, (self).f24})
                d_1494_v14_ = nw237_
                d_1495_v15_: _dafny.Seq
                d_1495_v15_ = _dafny.SeqWithoutIsStrInference([d_1494_v14_])
                d_1496_v16_: _dafny.Seq
                d_1496_v16_ = _dafny.SeqWithoutIsStrInference([(d_1485_v5_).f34, d_1480_v0_])
                d_1497_v17_: _dafny.Array
                nw238_ = _dafny.Array(_dafny.Map({}), 8)
                d_1497_v17_ = nw238_
                d_1498_v18_: C3
                nw239_ = C3()
                nw239_.ctor__((d_1495_v15_)[default__.safeIndex(len(d_1496_v16_), len(d_1495_v15_))], d_1497_v17_, (self).f24, (d_1489_v9_).f25)
                d_1498_v18_ = nw239_
            elif True:
                d_1499_v19_: _dafny.Array
                nw240_ = _dafny.Array(_dafny.Set({}), 2)
                d_1499_v19_ = nw240_
                d_1500_v21_: _dafny.Array
                def lambda62_(d_1501_i0_):
                    def iife101_():
                        coll43_ = _dafny.Map()
                        compr_43_: int
                        for compr_43_ in _dafny.IntegerRange(429, 778):
                            d_1502_v20_: int = compr_43_
                            if ((429) <= (d_1502_v20_)) and ((d_1502_v20_) < (778)):
                                coll43_[(d_1502_v20_) + ((self).f24)] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgqm"))))
                        return _dafny.Map(coll43_)
                    return (D11_DC30(iife101_()
)).cf47

                init36_ = lambda62_
                nw241_ = _dafny.Array(None, 13)
                for i0_36_ in range(nw241_.length(0)):
                    nw241_[i0_36_] = init36_(i0_36_)
                d_1500_v21_ = nw241_
                d_1503_v22_: T0
                nw242_ = C3()
                nw242_.ctor__(d_1499_v19_, d_1500_v21_, (self).f24, (self).f25)
                d_1503_v22_ = nw242_
                d_1504_v23_: T0
                d_1504_v23_ = d_1503_v22_
                d_1504_v23_ = d_1504_v23_
                d_1505_v24_: _dafny.MultiSet
                d_1505_v24_ = _dafny.MultiSet([d_1503_v22_, d_1503_v22_])
                d_1506_v25_: _dafny.Set
                d_1506_v25_ = _dafny.Set({(d_1505_v24_).isdisjoint(d_1505_v24_)})
                d_1507_v26_: D0
                d_1507_v26_ = D0_DC1((self).f24, 279, d_1480_v0_)
                d_1508_v27_: str
                d_1508_v27_ = _dafny.CodePoint('n')
                pat_let_tv25_ = d_1480_v0_
                def iife102_(_pat_let29_0):
                    def iife103_(d_1509_dt__update__tmp_h0_):
                        def iife104_(_pat_let30_0):
                            def iife105_(d_1510_dt__update_hcf1_h0_):
                                def iife106_(_pat_let31_0):
                                    def iife107_(d_1511_dt__update_hcf3_h0_):
                                        return D0_DC1(d_1510_dt__update_hcf1_h0_, (d_1509_dt__update__tmp_h0_).cf2, d_1511_dt__update_hcf3_h0_)
                                    return iife107_(_pat_let31_0)
                                return iife106_(pat_let_tv25_)
                            return iife105_(_pat_let30_0)
                        return iife104_(self.f27)
                    return iife103_(_pat_let29_0)
                rhs254_ = (0) - (default__.safeDivisionInt((d_1503_v22_).f24, (d_1503_v22_).f24))
                rhs255_ = _dafny.Set({(iife102_(d_1507_v26_)).cf3, default__.fm2(self.f27, default__.fm0(-397, (0) - ((d_1503_v22_).f24), globalState), d_1508_v27_, globalState)})
                rhs256_ = (self).f24
                lhs204_ = globalState
                lhs205_ = globalState
                lhs204_.f14 = rhs254_
                d_1506_v25_ = rhs255_
                lhs205_.f18 = rhs256_
                (globalState).f2 = (d_1480_v0_ if d_1480_v0_ else not(d_1480_v0_))
                d_1512_v28_: _dafny.Array
                nw243_ = _dafny.Array(int(0), 20)
                d_1512_v28_ = nw243_
                d_1512_v28_ = d_1512_v28_
                d_1513_v29_: _dafny.MultiSet
                d_1513_v29_ = _dafny.MultiSet([d_1480_v0_, d_1480_v0_])
                d_1514_v30_: _dafny.Seq
                d_1514_v30_ = _dafny.SeqWithoutIsStrInference([self.f27])
                d_1515_v31_: _dafny.Set
                d_1515_v31_ = _dafny.Set({(self).f24, len(d_1514_v30_), (self).f24, self.f27, self.f27})
                d_1516_v32_: _dafny.Set
                d_1516_v32_ = _dafny.Set({d_1508_v27_, d_1508_v27_})
                d_1517_v33_: _dafny.Array
                nw244_ = _dafny.Array(None, 29)
                nw244_[int(0)] = not (d_1480_v0_) or (d_1480_v0_)
                nw244_[int(1)] = d_1480_v0_
                nw244_[int(2)] = d_1480_v0_
                nw244_[int(3)] = ((self).f24) >= ((d_1503_v22_).f24)
                nw244_[int(4)] = d_1480_v0_
                nw244_[int(5)] = d_1480_v0_
                nw244_[int(6)] = False
                nw244_[int(7)] = d_1480_v0_
                nw244_[int(8)] = False
                nw244_[int(9)] = d_1480_v0_
                nw244_[int(10)] = d_1480_v0_
                nw244_[int(11)] = ((d_1503_v22_).f24) >= (self.f27)
                nw244_[int(12)] = (d_1507_v26_).cf3
                nw244_[int(13)] = True
                nw244_[int(14)] = d_1480_v0_
                nw244_[int(15)] = (self.f27) == ((d_1513_v29_).cardinality)
                nw244_[int(16)] = ((d_1503_v22_).f24) < ((d_1503_v22_).f24)
                nw244_[int(17)] = d_1480_v0_
                nw244_[int(18)] = d_1480_v0_
                nw244_[int(19)] = d_1480_v0_
                nw244_[int(20)] = (d_1515_v31_).issubset(d_1515_v31_)
                nw244_[int(21)] = d_1480_v0_
                nw244_[int(22)] = d_1480_v0_
                nw244_[int(23)] = not(((self).f24) < (len(d_1516_v32_)))
                nw244_[int(24)] = not(d_1480_v0_)
                nw244_[int(25)] = (d_1515_v31_).ispropersubset(_dafny.Set({self.f27, self.f27}))
                nw244_[int(26)] = (d_1480_v0_ if d_1480_v0_ else d_1480_v0_)
                nw244_[int(27)] = ((self).f24) == (self.f27)
                nw244_[int(28)] = d_1480_v0_
                d_1517_v33_ = nw244_
                index212_ = default__.safeIndex(644, (d_1517_v33_).length(0))
                (d_1517_v33_)[index212_] = d_1480_v0_
                d_1518_v34_: _dafny.Seq
                d_1518_v34_ = _dafny.SeqWithoutIsStrInference([d_1508_v27_])
                d_1519_v35_: _dafny.MultiSet
                d_1519_v35_ = _dafny.MultiSet([(self).f24, len(d_1518_v34_)])
                d_1520_v36_: C0
                nw245_ = C0()
                nw245_.ctor__(d_1480_v0_)
                d_1520_v36_ = nw245_
                d_1521_v37_: _dafny.Map
                d_1521_v37_ = _dafny.Map({d_1520_v36_: _dafny.MultiSet(d_1514_v30_)})
                index213_ = default__.safeIndex(644, (d_1517_v33_).length(0))
                (d_1517_v33_)[index213_] = ((d_1519_v35_).intersection(((d_1521_v37_)[d_1520_v36_] if (d_1520_v36_) in (d_1521_v37_) else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f24]))))).issubset(d_1519_v35_)
            (globalState).f18 = default__.safeDivisionInt((self).f24, 798)
            d_1522_v38_: str
            d_1522_v38_ = _dafny.CodePoint('j')
            d_1523_v39_: _dafny.Seq
            d_1523_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhvqnlf"))
            d_1480_v0_ = ((d_1522_v38_) in (d_1523_v39_) if d_1480_v0_ else not((d_1480_v0_ if d_1480_v0_ else d_1480_v0_)))
            (globalState).f17 = d_1523_v39_
        elif True:
            (globalState).f1 = (0) - ((self).f24)
            d_1524_v40_: str
            d_1524_v40_ = _dafny.CodePoint('t')
            if default__.fm2((self).f24, (self).f24, d_1524_v40_, globalState):
                (globalState).f14 = (self).f24
                d_1525_v41_: D12
                d_1525_v41_ = D12_DC33(d_1524_v40_, d_1480_v0_)
                d_1526_v42_: D12
                d_1526_v42_ = D12_DC34((d_1525_v41_ if not(d_1480_v0_) else d_1525_v41_))
                d_1527_v43_: _dafny.Array
                nw246_ = _dafny.Array(int(0), 26)
                d_1527_v43_ = nw246_
                d_1528_v44_: _dafny.Seq
                d_1528_v44_ = _dafny.SeqWithoutIsStrInference([d_1527_v43_])
                d_1529_v45_: T0
                nw247_ = C4()
                nw247_.ctor__((d_1528_v44_)[default__.safeIndex((self).f24, len(d_1528_v44_))], 191, _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('c'): self.f27}) for d_1530_i1_ in range(default__.abs(181))]))
                d_1529_v45_ = nw247_
                d_1531_v46_: _dafny.Seq
                d_1531_v46_ = _dafny.SeqWithoutIsStrInference([d_1480_v0_, d_1480_v0_, d_1480_v0_])
                d_1532_v47_: _dafny.Map
                d_1532_v47_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(D8_DC25((d_1529_v45_).f24, d_1480_v0_, d_1480_v0_)).cf38 for d_1533_i2_ in range(default__.abs(972))]): d_1529_v45_})
                d_1534_v48_: _dafny.Map
                d_1534_v48_ = _dafny.Map({len(_dafny.Set({(d_1529_v45_).f24, (self).f24})): (d_1529_v45_).f24})
                d_1535_v49_: _dafny.Seq
                d_1535_v49_ = _dafny.SeqWithoutIsStrInference([len(d_1534_v48_), default__.fm0((d_1529_v45_).f24, self.f27, globalState)])
                rhs257_ = d_1526_v42_
                rhs258_ = (d_1531_v46_) <= (d_1531_v46_)
                rhs259_ = ((d_1532_v47_)[d_1535_v49_] if (d_1535_v49_) in (d_1532_v47_) else d_1529_v45_)
                d_1526_v42_ = rhs257_
                d_1480_v0_ = rhs258_
                d_1529_v45_ = rhs259_
                d_1536_v50_: C4
                nw248_ = C4()
                nw248_.ctor__(d_1527_v43_, (d_1529_v45_).f24, (d_1529_v45_).f25)
                d_1536_v50_ = nw248_
                d_1537_v51_: _dafny.Map
                d_1537_v51_ = _dafny.Map({(d_1529_v45_).f24: d_1536_v50_})
                d_1538_v52_: _dafny.Map
                d_1538_v52_ = _dafny.Map({True: len(d_1537_v51_)})
                d_1539_v53_: _dafny.MultiSet
                d_1539_v53_ = _dafny.MultiSet([d_1480_v0_])
                d_1540_v54_: _dafny.Seq
                d_1540_v54_ = _dafny.SeqWithoutIsStrInference([(d_1538_v52_).set(d_1480_v0_, (d_1539_v53_).cardinality), d_1538_v52_])
                d_1540_v54_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({False: len(d_1535_v49_)})).set(d_1480_v0_, self.f27), d_1538_v52_])
                r0 = d_1480_v0_
                (globalState).f2 = (d_1480_v0_) == (d_1480_v0_)
            elif True:
                d_1541_v55_: _dafny.Seq
                d_1541_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drbtc"))
                d_1542_v56_: _dafny.Set
                d_1542_v56_ = _dafny.Set({len(d_1541_v55_)})
                d_1543_v57_: _dafny.Seq
                d_1543_v57_ = _dafny.SeqWithoutIsStrInference([d_1480_v0_, d_1480_v0_])
                d_1544_v58_: _dafny.MultiSet
                d_1544_v58_ = _dafny.MultiSet([(self).f24, len(d_1541_v55_)])
                d_1545_v60_: D6
                d_1545_v60_ = D6_DC16(d_1542_v56_)
                d_1546_v61_: _dafny.Array
                nw249_ = _dafny.Array(None, 28)
                nw249_[int(0)] = d_1542_v56_
                nw249_[int(1)] = d_1542_v56_
                nw249_[int(2)] = d_1542_v56_
                nw249_[int(3)] = _dafny.Set({self.f27, self.f27, len(d_1543_v57_)})
                nw249_[int(4)] = d_1542_v56_
                nw249_[int(5)] = d_1542_v56_
                nw249_[int(6)] = d_1542_v56_
                nw249_[int(7)] = d_1542_v56_
                def iife108_():
                    coll44_ = _dafny.Set()
                    compr_44_: int
                    for compr_44_ in (d_1544_v58_).Elements:
                        d_1547_v59_: int = compr_44_
                        if (d_1547_v59_) in (d_1544_v58_):
                            coll44_ = coll44_.union(_dafny.Set([(d_1547_v59_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omdiun"))))]))
                    return _dafny.Set(coll44_)
                nw249_[int(8)] = iife108_()
                
                nw249_[int(9)] = d_1542_v56_
                nw249_[int(10)] = d_1542_v56_
                nw249_[int(11)] = (d_1545_v60_).cf27
                nw249_[int(12)] = d_1542_v56_
                nw249_[int(13)] = d_1542_v56_
                nw249_[int(14)] = d_1542_v56_
                nw249_[int(15)] = d_1542_v56_
                nw249_[int(16)] = d_1542_v56_
                nw249_[int(17)] = d_1542_v56_
                nw249_[int(18)] = d_1542_v56_
                nw249_[int(19)] = d_1542_v56_
                nw249_[int(20)] = d_1542_v56_
                nw249_[int(21)] = _dafny.Set({510, len(d_1542_v56_)})
                nw249_[int(22)] = d_1542_v56_
                nw249_[int(23)] = d_1542_v56_
                nw249_[int(24)] = d_1542_v56_
                nw249_[int(25)] = d_1542_v56_
                nw249_[int(26)] = d_1542_v56_
                nw249_[int(27)] = d_1542_v56_
                d_1546_v61_ = nw249_
                d_1548_v62_: _dafny.Array
                nw250_ = _dafny.Array(_dafny.Map({}), 11)
                d_1548_v62_ = nw250_
                d_1549_v63_: _dafny.Seq
                d_1549_v63_ = (self).f25
                d_1550_v64_: C3
                nw251_ = C3()
                nw251_.ctor__(d_1546_v61_, d_1548_v62_, (0) - ((self).f24), (d_1549_v63_))
                d_1550_v64_ = nw251_
                d_1551_v65_: _dafny.Map
                d_1551_v65_ = _dafny.Map({d_1480_v0_: d_1550_v64_})
                d_1551_v65_ = (d_1551_v65_).set(False, d_1550_v64_)
                d_1552_v66_: _dafny.Map
                d_1552_v66_ = _dafny.Map({d_1480_v0_: self.f27})
                d_1553_v67_: _dafny.Map
                d_1553_v67_ = _dafny.Map({not(d_1480_v0_): d_1552_v66_})
                d_1553_v67_ = ((d_1553_v67_) | (d_1553_v67_)) | (d_1553_v67_)
                d_1552_v66_ = default__.fm17(default__.fm0(len(d_1541_v55_), -163, globalState), self.f27, self.f27, d_1480_v0_, globalState)
                d_1554_v68_: _dafny.Array
                nw252_ = _dafny.Array(_dafny.CodePoint('D'), 22)
                d_1554_v68_ = nw252_
                index214_ = default__.safeIndex(96, (d_1554_v68_).length(0))
                (d_1554_v68_)[index214_] = d_1524_v40_
                index215_ = default__.safeIndex(96, (d_1554_v68_).length(0))
                (d_1554_v68_)[index215_] = d_1524_v40_
                d_1555_v69_: _dafny.Map
                d_1555_v69_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(d_1543_v57_)]): (0) - (default__.safeDivisionInt((0) - (len(_dafny.Map({False: False}))), len(d_1541_v55_)))})
                d_1556_v70_: _dafny.Seq
                d_1556_v70_ = _dafny.SeqWithoutIsStrInference([self.f27])
                d_1557_v71_: _dafny.MultiSet
                d_1557_v71_ = _dafny.MultiSet([d_1480_v0_])
                d_1555_v69_ = (d_1555_v69_).set((d_1556_v70_) + (d_1556_v70_), (self.f27) - ((d_1557_v71_).cardinality))
            d_1558_v72_: _dafny.Set
            d_1558_v72_ = _dafny.Set({self.f27})
            d_1559_v73_: _dafny.Seq
            d_1559_v73_ = _dafny.SeqWithoutIsStrInference([(self).f24, -217])
            d_1560_v75_: _dafny.Array
            nw253_ = _dafny.Array(None, 14)
            nw253_[int(0)] = (d_1558_v72_) - (d_1558_v72_)
            nw253_[int(1)] = d_1558_v72_
            nw253_[int(2)] = _dafny.Set({self.f27, (self).f24, (self).f24})
            nw253_[int(3)] = (_dafny.Set({self.f27, self.f27})) | (d_1558_v72_)
            nw253_[int(4)] = d_1558_v72_
            nw253_[int(5)] = d_1558_v72_
            nw253_[int(6)] = d_1558_v72_
            nw253_[int(7)] = d_1558_v72_
            nw253_[int(8)] = d_1558_v72_
            nw253_[int(9)] = d_1558_v72_
            nw253_[int(10)] = _dafny.Set({self.f27, self.f27, (self).f24, len(d_1559_v73_)})
            def iife109_():
                coll45_ = _dafny.Set()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(234, -416):
                    d_1561_v74_: int = compr_45_
                    if ((234) <= (d_1561_v74_)) and ((d_1561_v74_) < (-416)):
                        coll45_ = coll45_.union(_dafny.Set([default__.safeModuloInt(d_1561_v74_, (self).f24)]))
                return _dafny.Set(coll45_)
            nw253_[int(11)] = iife109_()
            
            nw253_[int(12)] = d_1558_v72_
            nw253_[int(13)] = _dafny.Set({(self).f24, (self).f24})
            d_1560_v75_ = nw253_
            d_1560_v75_ = d_1560_v75_
            d_1562_v76_: _dafny.Array
            nw254_ = _dafny.Array(_dafny.Map({}), 4)
            d_1562_v76_ = nw254_
            d_1563_v77_: _dafny.Array
            nw255_ = _dafny.Array(False, 20)
            d_1563_v77_ = nw255_
            d_1564_v78_: _dafny.Map
            d_1564_v78_ = _dafny.Map({False: d_1563_v77_})
            index216_ = default__.safeIndex(190, (d_1562_v76_).length(0))
            (d_1562_v76_)[index216_] = (d_1564_v78_) | (d_1564_v78_)
            d_1565_v79_: _dafny.Map
            d_1565_v79_ = d_1564_v78_
            index217_ = default__.safeIndex(190, (d_1562_v76_).length(0))
            (d_1562_v76_)[index217_] = (d_1565_v79_)
            (globalState).f18 = 772
        (globalState).f1 = len(default__.fm40(globalState))
        (self).m3((0) - (default__.safeDivisionInt((self).f24, (self).f24)), globalState)
        d_1566_v80_: _dafny.MultiSet
        d_1566_v80_ = _dafny.MultiSet([self.f27])
        if (d_1566_v80_).ispropersubset(d_1566_v80_):
            d_1567_v81_: _dafny.Seq
            d_1567_v81_ = _dafny.SeqWithoutIsStrInference([d_1480_v0_])
            if (not (d_1480_v0_) or (d_1480_v0_) if d_1480_v0_ else (d_1567_v81_)[default__.safeIndex(-585, len(d_1567_v81_))]):
                (globalState).f18 = self.f27
                r0 = not((d_1480_v0_ if (self.f27) != ((self).f24) else not(d_1480_v0_)))
                d_1568_v82_: _dafny.Seq
                d_1568_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lofb"))
                d_1480_v0_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "th"))) != (d_1568_v82_)
                d_1569_v83_: _dafny.Set
                d_1569_v83_ = _dafny.Set({d_1480_v0_})
                d_1569_v83_ = (d_1569_v83_) | ((d_1569_v83_).intersection(d_1569_v83_))
                (globalState).f17 = ((d_1568_v82_) + (d_1568_v82_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyxkjjg")))
            elif True:
                (globalState).f2 = ((((self).f26)[d_1480_v0_] if (d_1480_v0_) in ((self).f26) else d_1480_v0_) if d_1480_v0_ else d_1480_v0_)
                d_1570_v84_: D0
                d_1570_v84_ = D0_DC1(self.f27, self.f27, d_1480_v0_)
                d_1571_v85_: _dafny.Seq
                d_1571_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qljejp"))
                d_1572_v86_: _dafny.Set
                d_1572_v86_ = _dafny.Set({(0) - (len(d_1571_v85_))})
                d_1570_v84_ = D0_DC1((self).f24, ((self).f24) * (632), (d_1572_v86_).issubset(d_1572_v86_))
                d_1573_v87_: _dafny.Array
                def lambda63_(d_1574_v0_):
                    def lambda64_(d_1575_i3_):
                        return d_1574_v0_

                    return lambda64_

                init37_ = lambda63_(d_1480_v0_)
                nw256_ = _dafny.Array(None, 18)
                for i0_37_ in range(nw256_.length(0)):
                    nw256_[i0_37_] = init37_(i0_37_)
                d_1573_v87_ = nw256_
                d_1576_v88_: _dafny.Map
                d_1576_v88_ = _dafny.Map({self.f27: d_1480_v0_})
                d_1577_v89_: _dafny.Map
                d_1577_v89_ = _dafny.Map({((d_1576_v88_)[918] if (918) in (d_1576_v88_) else d_1480_v0_): (self).f24})
                index218_ = default__.safeIndex(573, (d_1573_v87_).length(0))
                (d_1573_v87_)[index218_] = not((d_1577_v89_) != (_dafny.Map({d_1480_v0_: self.f27})))
                d_1578_v90_: _dafny.Map
                d_1578_v90_ = _dafny.Map({d_1571_v85_: 327})
                index219_ = default__.safeIndex(573, (d_1573_v87_).length(0))
                def iife110_():
                    coll46_ = _dafny.Map()
                    compr_46_: _dafny.Seq
                    for compr_46_ in (_dafny.SeqWithoutIsStrInference([d_1571_v85_ for d_1579_i4_ in range(default__.abs(422))])).Elements:
                        d_1580_v91_: _dafny.Seq = compr_46_
                        if (d_1580_v91_) in (_dafny.SeqWithoutIsStrInference([d_1571_v85_ for d_1579_i4_ in range(default__.abs(422))])):
                            coll46_[d_1580_v91_] = self.f27
                    return _dafny.Map(coll46_)
                rhs260_ = d_1480_v0_
                rhs261_ = (_dafny.Map({d_1571_v85_: default__.fm0(self.f27, self.f27, globalState)})) | ((_dafny.Map({d_1571_v85_: (self).f24})) | (iife110_()
                ))
                lhs206_ = d_1573_v87_
                lhs207_ = default__.safeIndex(573, (d_1573_v87_).length(0))
                lhs206_[lhs207_] = rhs260_
                d_1578_v90_ = rhs261_
                (globalState).f9 = (self).f24
                d_1581_v92_: _dafny.Array
                nw257_ = _dafny.Array(_dafny.Map({}), 3)
                d_1581_v92_ = nw257_
                d_1582_v93_: _dafny.MultiSet
                d_1582_v93_ = _dafny.MultiSet([(d_1573_v87_)[default__.safeIndex(573, (d_1573_v87_).length(0))]])
                d_1583_v94_: _dafny.Seq
                d_1583_v94_ = _dafny.SeqWithoutIsStrInference([len(d_1577_v89_)])
                d_1584_v95_: D3
                d_1584_v95_ = D3_DC11(self.f27, (d_1567_v81_)[default__.safeIndex(len(d_1583_v94_), len(d_1567_v81_))], True)
                index220_ = default__.safeIndex(573, (d_1573_v87_).length(0))
                rhs262_ = (d_1566_v80_) | (d_1566_v80_)
                rhs263_ = self.f27
                rhs264_ = (((d_1582_v93_).cardinality) < (self.f27) if (d_1573_v87_)[default__.safeIndex(573, (d_1573_v87_).length(0))] else (d_1584_v95_).cf20)
                rhs265_ = (d_1581_v92_ if not((d_1573_v87_)[default__.safeIndex(573, (d_1573_v87_).length(0))]) else d_1581_v92_)
                lhs208_ = globalState
                lhs209_ = d_1573_v87_
                lhs210_ = default__.safeIndex(573, (d_1573_v87_).length(0))
                d_1566_v80_ = rhs262_
                lhs208_.f9 = rhs263_
                lhs209_[lhs210_] = rhs264_
                d_1581_v92_ = rhs265_
            d_1585_v96_: _dafny.Array
            nw258_ = _dafny.Array(False, 5)
            d_1585_v96_ = nw258_
            d_1586_v97_: D3
            d_1586_v97_ = D3_DC11(self.f27, d_1480_v0_, d_1480_v0_)
            d_1587_v98_: _dafny.Array
            nw259_ = _dafny.Array(None, 27)
            nw259_[int(0)] = d_1585_v96_
            nw259_[int(1)] = d_1585_v96_
            nw259_[int(2)] = (d_1585_v96_ if d_1480_v0_ else d_1585_v96_)
            nw259_[int(3)] = d_1585_v96_
            nw259_[int(4)] = d_1585_v96_
            nw259_[int(5)] = d_1585_v96_
            nw259_[int(6)] = d_1585_v96_
            nw259_[int(7)] = d_1585_v96_
            nw259_[int(8)] = d_1585_v96_
            nw259_[int(9)] = d_1585_v96_
            nw259_[int(10)] = d_1585_v96_
            nw259_[int(11)] = d_1585_v96_
            nw259_[int(12)] = d_1585_v96_
            nw259_[int(13)] = d_1585_v96_
            nw259_[int(14)] = d_1585_v96_
            nw259_[int(15)] = d_1585_v96_
            nw259_[int(16)] = d_1585_v96_
            nw259_[int(17)] = d_1585_v96_
            nw259_[int(18)] = d_1585_v96_
            nw259_[int(19)] = d_1585_v96_
            nw259_[int(20)] = d_1585_v96_
            nw259_[int(21)] = d_1585_v96_
            nw259_[int(22)] = (d_1585_v96_ if (d_1586_v97_).cf19 else d_1585_v96_)
            nw259_[int(23)] = d_1585_v96_
            nw259_[int(24)] = d_1585_v96_
            nw259_[int(25)] = d_1585_v96_
            nw259_[int(26)] = d_1585_v96_
            d_1587_v98_ = nw259_
            index221_ = default__.safeIndex(211, (d_1587_v98_).length(0))
            (d_1587_v98_)[index221_] = d_1585_v96_
            index222_ = default__.safeIndex(211, (d_1587_v98_).length(0))
            (d_1587_v98_)[index222_] = d_1585_v96_
            d_1588_v99_: _dafny.MultiSet
            d_1588_v99_ = _dafny.MultiSet([d_1480_v0_, d_1480_v0_, d_1480_v0_, d_1480_v0_, d_1480_v0_])
            d_1589_v100_: str
            d_1589_v100_ = _dafny.CodePoint('n')
            d_1590_v101_: _dafny.Map
            d_1590_v101_ = _dafny.Map({default__.fm2((self).f24, (self).f24, d_1589_v100_, globalState): len(_dafny.Map({(self).f24: self.f27}))})
            if (_dafny.MultiSet([not(d_1480_v0_), False])).issubset((d_1588_v99_).set(d_1480_v0_, default__.abs(len(d_1590_v101_)))):
                r0 = (d_1566_v80_).isdisjoint(d_1566_v80_)
                d_1591_v102_: _dafny.Set
                d_1591_v102_ = _dafny.Set({(self).f24})
                d_1592_v103_: _dafny.Map
                d_1592_v103_ = _dafny.Map({d_1591_v102_: (self).f24})
                d_1592_v103_ = (d_1592_v103_).set(default__.fm18(globalState), default__.safeModuloInt(len(d_1567_v81_), self.f27))
                arr0_ = (d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))]
                index223_ = default__.safeIndex(777, ((d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))]).length(0))
                arr0_[index223_] = not(d_1480_v0_)
                arr1_ = (d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))]
                index224_ = default__.safeIndex(777, ((d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))]).length(0))
                arr1_[index224_] = (default__.fm0((self).f24, (self).f24, globalState)) <= (self.f27)
                (globalState).f2 = d_1480_v0_
                d_1593_v104_: _dafny.Seq
                d_1593_v104_ = _dafny.SeqWithoutIsStrInference([d_1589_v100_, d_1589_v100_, d_1589_v100_])
                d_1594_v105_: D1
                d_1594_v105_ = D1_DC3(d_1593_v104_)
                d_1595_v106_: _dafny.Set
                d_1595_v106_ = _dafny.Set({d_1566_v80_})
                d_1596_v107_: _dafny.Seq
                d_1596_v107_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_1566_v80_, d_1566_v80_, d_1566_v80_})])
                d_1597_v109_: _dafny.Seq
                d_1597_v109_ = _dafny.SeqWithoutIsStrInference([d_1566_v80_, d_1566_v80_])
                d_1598_v110_: C1
                nw260_ = C1()
                nw260_.ctor__(d_1597_v109_, (self).f24, (self).f24, (self).f25)
                d_1598_v110_ = nw260_
                d_1599_v111_: _dafny.Map
                d_1599_v111_ = _dafny.Map({d_1598_v110_: ((d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))])[default__.safeIndex(777, ((d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))]).length(0))]})
                d_1600_v112_: _dafny.Seq
                d_1600_v112_ = _dafny.SeqWithoutIsStrInference([((d_1566_v80_)[(self).f24] if ((self).f24) in (d_1566_v80_) else len(d_1599_v111_))])
                d_1601_v113_: _dafny.Map
                d_1601_v113_ = _dafny.Map({d_1598_v110_.f33: _dafny.Set({_dafny.MultiSet(d_1600_v112_), d_1566_v80_})})
                d_1602_v114_: _dafny.Map
                d_1602_v114_ = _dafny.Map({d_1566_v80_: self.f27})
                d_1603_v116_: _dafny.Array
                nw261_ = _dafny.Array(None, 27)
                nw261_[int(0)] = d_1595_v106_
                nw261_[int(1)] = d_1595_v106_
                nw261_[int(2)] = (d_1595_v106_) - (d_1595_v106_)
                nw261_[int(3)] = d_1595_v106_
                nw261_[int(4)] = d_1595_v106_
                nw261_[int(5)] = (d_1595_v106_) - (d_1595_v106_)
                nw261_[int(6)] = ((d_1596_v107_)[default__.safeIndex(len(d_1593_v104_), len(d_1596_v107_))]) - (d_1595_v106_)
                nw261_[int(7)] = d_1595_v106_
                nw261_[int(8)] = d_1595_v106_
                nw261_[int(9)] = (d_1595_v106_ if d_1480_v0_ else d_1595_v106_)
                nw261_[int(10)] = d_1595_v106_
                nw261_[int(11)] = (d_1595_v106_).intersection(default__.fm41((self).f24, default__.fm4(d_1480_v0_, (self).f24, d_1589_v100_, not(d_1480_v0_), globalState), not(not(False)), globalState))
                def iife111_():
                    coll47_ = _dafny.Set()
                    compr_47_: _dafny.MultiSet
                    for compr_47_ in (d_1595_v106_).Elements:
                        d_1604_v108_: _dafny.MultiSet = compr_47_
                        if (d_1604_v108_) in (d_1595_v106_):
                            coll47_ = coll47_.union(_dafny.Set([d_1604_v108_]))
                    return _dafny.Set(coll47_)
                nw261_[int(12)] = iife111_()
                
                nw261_[int(13)] = ((d_1596_v107_)[default__.safeIndex(len(d_1600_v112_), len(d_1596_v107_))]) - (((d_1601_v113_)[(0) - (d_1598_v110_.f33)] if ((0) - (d_1598_v110_.f33)) in (d_1601_v113_) else d_1595_v106_))
                def iife112_():
                    coll48_ = _dafny.Set()
                    compr_48_: _dafny.MultiSet
                    for compr_48_ in (d_1602_v114_).keys.Elements:
                        d_1605_v115_: _dafny.MultiSet = compr_48_
                        if (d_1605_v115_) in (d_1602_v114_):
                            coll48_ = coll48_.union(_dafny.Set([d_1605_v115_]))
                    return _dafny.Set(coll48_)
                nw261_[int(14)] = (_dafny.Set({d_1566_v80_}) if not(d_1480_v0_) else iife112_()
                )
                nw261_[int(15)] = d_1595_v106_
                nw261_[int(16)] = _dafny.Set({d_1566_v80_, _dafny.MultiSet([(self).f24, d_1598_v110_.f33])})
                nw261_[int(17)] = (d_1595_v106_) - (d_1595_v106_)
                nw261_[int(18)] = d_1595_v106_
                nw261_[int(19)] = d_1595_v106_
                nw261_[int(20)] = (d_1596_v107_)[default__.safeIndex(self.f27, len(d_1596_v107_))]
                nw261_[int(21)] = d_1595_v106_
                nw261_[int(22)] = d_1595_v106_
                nw261_[int(23)] = _dafny.Set({d_1566_v80_})
                nw261_[int(24)] = default__.fm41(self.f27, d_1593_v104_, True, globalState)
                nw261_[int(25)] = d_1595_v106_
                nw261_[int(26)] = (d_1595_v106_) - (d_1595_v106_)
                d_1603_v116_ = nw261_
                d_1606_v117_: D8
                d_1606_v117_ = D8_DC25(self.f27, ((d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))])[default__.safeIndex(777, ((d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))]).length(0))], d_1480_v0_)
                index225_ = default__.safeIndex(222, (d_1603_v116_).length(0))
                (d_1603_v116_)[index225_] = default__.fm41((self).f24, _dafny.SeqWithoutIsStrInference([d_1589_v100_ for d_1607_i5_ in range(default__.abs(513))]), (d_1606_v117_).cf39, globalState)
                d_1608_v118_: _dafny.Array
                nw262_ = _dafny.Array(int(0), 12)
                d_1608_v118_ = nw262_
                d_1609_v119_: _dafny.Array
                nw263_ = _dafny.Array(None, 4)
                nw263_[int(0)] = d_1608_v118_
                nw263_[int(1)] = d_1608_v118_
                nw263_[int(2)] = d_1608_v118_
                nw263_[int(3)] = d_1608_v118_
                d_1609_v119_ = nw263_
                index226_ = default__.safeIndex(23, (d_1609_v119_).length(0))
                (d_1609_v119_)[index226_] = d_1608_v118_
                pat_let_tv26_ = d_1593_v104_
                index227_ = default__.safeIndex(222, (d_1603_v116_).length(0))
                index228_ = default__.safeIndex(23, (d_1609_v119_).length(0))
                def iife113_(_pat_let32_0):
                    def iife114_(d_1610_dt__update__tmp_h1_):
                        def iife115_(_pat_let33_0):
                            def iife116_(d_1611_dt__update_hcf7_h0_):
                                return D1_DC3(d_1611_dt__update_hcf7_h0_)
                            return iife116_(_pat_let33_0)
                        return iife115_(pat_let_tv26_)
                    return iife114_(_pat_let32_0)
                rhs266_ = (d_1594_v105_ if d_1480_v0_ else (d_1594_v105_ if ((d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))])[default__.safeIndex(777, ((d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))]).length(0))] else iife113_(d_1594_v105_)))
                rhs267_ = d_1595_v106_
                rhs268_ = d_1608_v118_
                lhs211_ = d_1603_v116_
                lhs212_ = default__.safeIndex(222, (d_1603_v116_).length(0))
                lhs213_ = d_1609_v119_
                lhs214_ = default__.safeIndex(23, (d_1609_v119_).length(0))
                d_1594_v105_ = rhs266_
                lhs211_[lhs212_] = rhs267_
                lhs213_[lhs214_] = rhs268_
            elif True:
                (globalState).f14 = ((self).f24 if d_1480_v0_ else self.f27)
                index229_ = default__.safeIndex(211, (d_1587_v98_).length(0))
                (d_1587_v98_)[index229_] = (d_1587_v98_)[default__.safeIndex(211, (d_1587_v98_).length(0))]
                d_1612_v120_: C2
                nw264_ = C2()
                nw264_.ctor__((d_1480_v0_ if d_1480_v0_ else d_1480_v0_), (self).f24, (self).f25)
                d_1612_v120_ = nw264_
                (globalState).f14 = (0) - (self.f27)
                d_1613_v121_: _dafny.Seq
                d_1613_v121_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlgkqke"))
                d_1614_v122_: C2
                nw265_ = C2()
                nw265_.ctor__(d_1480_v0_, (self.f27) - (len(d_1613_v121_)), (self).f25)
                d_1614_v122_ = nw265_
            d_1615_v123_: _dafny.Array
            def lambda65_(d_1616_i6_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sihbc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egjfu")))

            init38_ = lambda65_
            nw266_ = _dafny.Array(None, 13)
            for i0_38_ in range(nw266_.length(0)):
                nw266_[i0_38_] = init38_(i0_38_)
            d_1615_v123_ = nw266_
            d_1617_v124_: _dafny.Seq
            d_1617_v124_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1618_i7_ in range(default__.abs(813))])])
            index230_ = default__.safeIndex(574, (d_1615_v123_).length(0))
            (d_1615_v123_)[index230_] = (d_1617_v124_)[default__.safeIndex((0) - (self.f27), len(d_1617_v124_))]
            d_1619_v125_: _dafny.Seq
            d_1619_v125_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wugxlqdqp"))
            index231_ = default__.safeIndex(574, (d_1615_v123_).length(0))
            (d_1615_v123_)[index231_] = d_1619_v125_
            d_1620_v126_: _dafny.Map
            d_1620_v126_ = _dafny.Map({(self).f24: (self.f27) * (350)})
            d_1621_v127_: _dafny.Seq
            d_1621_v127_ = _dafny.SeqWithoutIsStrInference([(d_1566_v80_).cardinality, self.f27])
            d_1620_v126_ = (d_1620_v126_).set((self).f24, (d_1621_v127_)[default__.safeIndex(self.f27, len(d_1621_v127_))])
        elif True:
            d_1622_v128_: _dafny.Seq
            d_1622_v128_ = _dafny.SeqWithoutIsStrInference([self.f27, (self).f24])
            d_1623_v129_: C1
            nw267_ = C1()
            nw267_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([self.f27, (_dafny.MultiSet([d_1480_v0_, (D10_DC29((self).f24, d_1480_v0_)).cf46, d_1480_v0_])).cardinality]) for d_1624_i8_ in range(default__.abs(-395))]), len((d_1622_v128_).set(default__.safeIndex((d_1566_v80_).cardinality, len(d_1622_v128_)), self.f27)), self.f27, (self).f25)
            d_1623_v129_ = nw267_
            d_1625_v130_: _dafny.Map
            d_1625_v130_ = _dafny.Map({d_1623_v129_: d_1480_v0_})
            r0 = (True if ((d_1625_v130_)[d_1623_v129_] if (d_1623_v129_) in (d_1625_v130_) else d_1480_v0_) else d_1480_v0_)
            d_1626_v131_: _dafny.Map
            d_1626_v131_ = _dafny.Map({d_1480_v0_: 80})
            d_1627_v132_: _dafny.Seq
            d_1627_v132_ = _dafny.SeqWithoutIsStrInference([d_1480_v0_, True])
            d_1628_v133_: _dafny.Array
            nw268_ = _dafny.Array(None, 27)
            nw268_[int(0)] = (self).f24
            nw268_[int(1)] = default__.safeDivisionInt(d_1623_v129_.f33, len(d_1626_v131_))
            nw268_[int(2)] = d_1623_v129_.f33
            nw268_[int(3)] = (self).f24
            nw268_[int(4)] = d_1623_v129_.f33
            nw268_[int(5)] = d_1623_v129_.f33
            nw268_[int(6)] = self.f27
            nw268_[int(7)] = self.f27
            nw268_[int(8)] = (self).f24
            nw268_[int(9)] = ((d_1626_v131_)[d_1480_v0_] if (d_1480_v0_) in (d_1626_v131_) else (self).f24)
            nw268_[int(10)] = self.f27
            nw268_[int(11)] = (self).f24
            nw268_[int(12)] = default__.safeDivisionInt(len(d_1627_v132_), d_1623_v129_.f33)
            nw268_[int(13)] = default__.safeModuloInt((self).f24, self.f27)
            nw268_[int(14)] = self.f27
            nw268_[int(15)] = self.f27
            nw268_[int(16)] = default__.safeModuloInt(self.f27, self.f27)
            nw268_[int(17)] = self.f27
            nw268_[int(18)] = d_1623_v129_.f33
            nw268_[int(19)] = 676
            nw268_[int(20)] = d_1623_v129_.f33
            nw268_[int(21)] = d_1623_v129_.f33
            nw268_[int(22)] = (0) - ((356) + ((self).f24))
            nw268_[int(23)] = self.f27
            nw268_[int(24)] = (self).f24
            nw268_[int(25)] = self.f27
            nw268_[int(26)] = self.f27
            d_1628_v133_ = nw268_
            index232_ = default__.safeIndex(717, (d_1628_v133_).length(0))
            (d_1628_v133_)[index232_] = d_1623_v129_.f33
            index233_ = default__.safeIndex(717, (d_1628_v133_).length(0))
            (d_1628_v133_)[index233_] = default__.safeDivisionInt((self).f24, 574)
            d_1629_v134_: str
            d_1629_v134_ = _dafny.CodePoint('r')
            d_1630_v135_: _dafny.Map
            d_1630_v135_ = _dafny.Map({((d_1566_v80_)[d_1623_v129_.f33] if (d_1623_v129_.f33) in (d_1566_v80_) else d_1623_v129_.f33): d_1629_v134_})
            d_1631_v137_: _dafny.Seq
            d_1631_v137_ = _dafny.SeqWithoutIsStrInference([d_1629_v134_, _dafny.CodePoint('f'), d_1629_v134_, d_1629_v134_, _dafny.CodePoint('q')])
            d_1632_v138_: _dafny.Seq
            def iife117_():
                coll49_ = _dafny.Map()
                compr_49_: str
                for compr_49_ in (d_1631_v137_).Elements:
                    d_1633_v136_: str = compr_49_
                    if (d_1633_v136_) in (d_1631_v137_):
                        coll49_[d_1633_v136_] = True
                return _dafny.Map(coll49_)
            d_1632_v138_ = _dafny.SeqWithoutIsStrInference([iife117_()
            ])
            d_1634_v139_: T0
            nw269_ = C1()
            nw269_.ctor__((d_1623_v129_).f32, 207, len(d_1632_v138_), (self).f25)
            d_1634_v139_ = nw269_
            d_1635_v140_: _dafny.Seq
            d_1635_v140_ = _dafny.SeqWithoutIsStrInference([d_1634_v139_])
            d_1636_v141_: C2
            nw270_ = C2()
            nw270_.ctor__((len(_dafny.SeqWithoutIsStrInference([(self).f24, len(d_1630_v135_), 797, self.f27, len(d_1635_v140_)]))) <= (d_1623_v129_.f33), (self).f24, (self).f25)
            d_1636_v141_ = nw270_
            d_1637_v142_: _dafny.Map
            d_1637_v142_ = _dafny.Map({420: d_1623_v129_.f33})
            d_1638_v143_: D8
            d_1638_v143_ = D8_DC24(((d_1637_v142_)[self.f27] if (self.f27) in (d_1637_v142_) else d_1623_v129_.f33))
            index234_ = default__.safeIndex(717, (d_1628_v133_).length(0))
            (d_1628_v133_)[index234_] = (d_1638_v143_).cf37
            d_1639_v144_: _dafny.Set
            d_1639_v144_ = _dafny.Set({d_1636_v141_.f31})
            d_1480_v0_ = (_dafny.Set({d_1480_v0_})).ispropersubset(d_1639_v144_)
        d_1480_v0_ = d_1480_v0_
        d_1640_v145_: _dafny.Array
        nw271_ = _dafny.Array(_dafny.MultiSet({}), 28)
        d_1640_v145_ = nw271_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1640_v145_).length(0)):
            d_1641_i9_: int = guard_loop_4_
            if (True) and (((0) <= (d_1641_i9_)) and ((d_1641_i9_) < ((d_1640_v145_).length(0)))):
                (d_1640_v145_)[(d_1641_i9_)] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_1480_v0_])) + (_dafny.SeqWithoutIsStrInference([d_1480_v0_])))
        r0 = not(d_1480_v0_)
        return r0

    def m3(self, p0, globalState):
        d_1642_v0_: _dafny.Seq
        d_1642_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhyjv"))
        d_1643_v1_: _dafny.MultiSet
        d_1643_v1_ = _dafny.MultiSet([len(d_1642_v0_), default__.fm0(p0, self.f27, globalState)])
        d_1644_v2_: _dafny.Map
        d_1644_v2_ = _dafny.Map({d_1643_v1_: False})
        if ((d_1644_v2_)[d_1643_v1_] if (d_1643_v1_) in (d_1644_v2_) else (default__.fm1(self.f27, 889, globalState)).issubset(d_1643_v1_)):
            if (866) > (p0):
                d_1645_v3_: bool
                d_1645_v3_ = True
                d_1646_v4_: C2
                nw272_ = C2()
                nw272_.ctor__(d_1645_v3_, p0, (self).f25)
                d_1646_v4_ = nw272_
                (globalState).f17 = d_1642_v0_
                d_1647_v5_: _dafny.Array
                nw273_ = _dafny.Array(_dafny.Seq({}), 10)
                d_1647_v5_ = nw273_
                d_1648_v6_: _dafny.Seq
                d_1648_v6_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p0 for d_1649_i0_ in range(default__.abs(770))])), 274, p0, self.f27, p0])
                index235_ = default__.safeIndex(632, (d_1647_v5_).length(0))
                (d_1647_v5_)[index235_] = (d_1648_v6_).set(default__.safeIndex(p0, len(d_1648_v6_)), self.f27)
                index236_ = default__.safeIndex(632, (d_1647_v5_).length(0))
                (d_1647_v5_)[index236_] = d_1648_v6_
                d_1650_v7_: _dafny.Array
                nw274_ = _dafny.Array(False, 13)
                d_1650_v7_ = nw274_
                index237_ = default__.safeIndex(623, (d_1650_v7_).length(0))
                (d_1650_v7_)[index237_] = (self.f27) <= (self.f27)
                index238_ = default__.safeIndex(623, (d_1650_v7_).length(0))
                (d_1650_v7_)[index238_] = d_1645_v3_
                d_1645_v3_ = d_1646_v4_.f31
            elif True:
                d_1651_v8_: _dafny.Array
                nw275_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                d_1651_v8_ = nw275_
                index239_ = default__.safeIndex(336, (d_1651_v8_).length(0))
                (d_1651_v8_)[index239_] = d_1642_v0_
                index240_ = default__.safeIndex(336, (d_1651_v8_).length(0))
                (d_1651_v8_)[index240_] = d_1642_v0_
                d_1652_v9_: bool
                d_1652_v9_ = True
                (globalState).f2 = not(d_1652_v9_)
                d_1653_v10_: C0
                nw276_ = C0()
                nw276_.ctor__((True) == (d_1652_v9_))
                d_1653_v10_ = nw276_
                d_1654_v11_: _dafny.Array
                nw277_ = _dafny.Array(None, 15)
                nw277_[int(0)] = (0) - ((self).f24)
                nw277_[int(1)] = p0
                nw277_[int(2)] = p0
                nw277_[int(3)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jk")))
                nw277_[int(4)] = (0) - ((self).f24)
                nw277_[int(5)] = p0
                nw277_[int(6)] = self.f27
                nw277_[int(7)] = (self).f24
                nw277_[int(8)] = self.f27
                nw277_[int(9)] = self.f27
                nw277_[int(10)] = (self).f24
                nw277_[int(11)] = p0
                nw277_[int(12)] = self.f27
                nw277_[int(13)] = p0
                nw277_[int(14)] = self.f27
                d_1654_v11_ = nw277_
                d_1655_v12_: _dafny.Seq
                d_1655_v12_ = _dafny.SeqWithoutIsStrInference([d_1654_v11_])
                d_1656_v13_: _dafny.Map
                d_1656_v13_ = _dafny.Map({p0: d_1654_v11_})
                d_1657_v14_: _dafny.Array
                nw278_ = _dafny.Array(None, 28)
                nw278_[int(0)] = (d_1655_v12_)[default__.safeIndex(p0, len(d_1655_v12_))]
                nw278_[int(1)] = d_1654_v11_
                nw278_[int(2)] = d_1654_v11_
                nw278_[int(3)] = (d_1654_v11_ if d_1652_v9_ else d_1654_v11_)
                nw278_[int(4)] = d_1654_v11_
                nw278_[int(5)] = d_1654_v11_
                nw278_[int(6)] = d_1654_v11_
                nw278_[int(7)] = d_1654_v11_
                nw278_[int(8)] = d_1654_v11_
                nw278_[int(9)] = (d_1655_v12_)[default__.safeIndex(731, len(d_1655_v12_))]
                nw278_[int(10)] = d_1654_v11_
                nw278_[int(11)] = d_1654_v11_
                nw278_[int(12)] = d_1654_v11_
                nw278_[int(13)] = d_1654_v11_
                nw278_[int(14)] = d_1654_v11_
                nw278_[int(15)] = (d_1654_v11_ if d_1652_v9_ else d_1654_v11_)
                nw278_[int(16)] = d_1654_v11_
                nw278_[int(17)] = d_1654_v11_
                nw278_[int(18)] = d_1654_v11_
                nw278_[int(19)] = d_1654_v11_
                nw278_[int(20)] = d_1654_v11_
                nw278_[int(21)] = d_1654_v11_
                nw278_[int(22)] = d_1654_v11_
                nw278_[int(23)] = d_1654_v11_
                nw278_[int(24)] = d_1654_v11_
                nw278_[int(25)] = d_1654_v11_
                nw278_[int(26)] = ((d_1656_v13_)[(self).f24] if ((self).f24) in (d_1656_v13_) else d_1654_v11_)
                nw278_[int(27)] = d_1654_v11_
                d_1657_v14_ = nw278_
                index241_ = default__.safeIndex(793, (d_1657_v14_).length(0))
                (d_1657_v14_)[index241_] = d_1654_v11_
                index242_ = default__.safeIndex(793, (d_1657_v14_).length(0))
                (d_1657_v14_)[index242_] = d_1654_v11_
                d_1658_v15_: _dafny.Array
                nw279_ = _dafny.Array(None, 27)
                d_1658_v15_ = nw279_
                d_1659_v16_: _dafny.Array
                nw280_ = _dafny.Array(False, 16)
                d_1659_v16_ = nw280_
                d_1660_v17_: _dafny.Map
                d_1660_v17_ = _dafny.Map({True: d_1659_v16_})
                d_1661_v18_: D16
                d_1661_v18_ = D16_DC40(d_1658_v15_)
                rhs269_ = ((691) * (len(d_1642_v0_))) * (len(d_1660_v17_))
                rhs270_ = (d_1661_v18_).cf63
                lhs215_ = self
                lhs215_.f27 = rhs269_
                d_1658_v15_ = rhs270_
            d_1662_v19_: _dafny.Array
            def lambda66_(d_1663_i1_):
                return True

            init39_ = lambda66_
            nw281_ = _dafny.Array(None, 25)
            for i0_39_ in range(nw281_.length(0)):
                nw281_[i0_39_] = init39_(i0_39_)
            d_1662_v19_ = nw281_
            d_1664_v20_: bool
            d_1664_v20_ = True
            index243_ = default__.safeIndex(167, (d_1662_v19_).length(0))
            (d_1662_v19_)[index243_] = d_1664_v20_
            d_1665_v21_: _dafny.Map
            d_1665_v21_ = _dafny.Map({d_1664_v20_: d_1664_v20_})
            d_1666_v22_: _dafny.Seq
            d_1666_v22_ = _dafny.SeqWithoutIsStrInference([d_1642_v0_, d_1642_v0_, d_1642_v0_])
            index244_ = default__.safeIndex(167, (d_1662_v19_).length(0))
            rhs271_ = d_1664_v20_
            rhs272_ = (d_1665_v21_) | ((self).f26)
            rhs273_ = len((d_1666_v22_)[default__.safeIndex((self).f24, len(d_1666_v22_))])
            lhs216_ = d_1662_v19_
            lhs217_ = default__.safeIndex(167, (d_1662_v19_).length(0))
            lhs218_ = globalState
            lhs216_[lhs217_] = rhs271_
            d_1665_v21_ = rhs272_
            lhs218_.f14 = rhs273_
            d_1667_v23_: C2
            nw282_ = C2()
            nw282_.ctor__((-986) <= (self.f27), ((0) - (self.f27)) - (p0), _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('q'): len(_dafny.SeqWithoutIsStrInference([p0 for d_1668_i3_ in range(default__.abs(84))]))}) for d_1669_i2_ in range(default__.abs(562))]))
            d_1667_v23_ = nw282_
            d_1667_v23_ = (d_1667_v23_ if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lylfv"))) != (d_1642_v0_) else d_1667_v23_)
            index245_ = default__.safeIndex(167, (d_1662_v19_).length(0))
            (d_1662_v19_)[index245_] = d_1667_v23_.f31
            d_1670_v24_: str
            d_1670_v24_ = _dafny.CodePoint('d')
            d_1671_v25_: _dafny.Set
            d_1671_v25_ = _dafny.Set({d_1662_v19_, d_1662_v19_})
            d_1672_v26_: _dafny.Seq
            d_1672_v26_ = _dafny.SeqWithoutIsStrInference([d_1671_v25_, d_1671_v25_])
            d_1673_v27_: _dafny.Map
            d_1673_v27_ = _dafny.Map({len((((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1674_i4_ in range(default__.abs(-379))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fg")))).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1675_i4_ in range(default__.abs(-379))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fg"))))), d_1670_v24_)).set(default__.safeIndex(self.f27, len(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1676_i4_ in range(default__.abs(-379))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fg")))).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1677_i4_ in range(default__.abs(-379))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fg"))))), d_1670_v24_))), d_1670_v24_)): (d_1671_v25_) | ((d_1672_v26_)[default__.safeIndex((0) - (self.f27), len(d_1672_v26_))])})
            d_1678_v28_: _dafny.Map
            d_1678_v28_ = _dafny.Map({d_1670_v24_: (self).f24})
            d_1679_v29_: _dafny.Array
            nw283_ = _dafny.Array(None, 23)
            nw283_[int(0)] = 427
            nw283_[int(1)] = p0
            nw283_[int(2)] = (self).f24
            nw283_[int(3)] = default__.fm0((self).f24, self.f27, globalState)
            nw283_[int(4)] = p0
            nw283_[int(5)] = p0
            nw283_[int(6)] = p0
            nw283_[int(7)] = p0
            nw283_[int(8)] = self.f27
            nw283_[int(9)] = (self).f24
            nw283_[int(10)] = self.f27
            nw283_[int(11)] = len(d_1678_v28_)
            nw283_[int(12)] = p0
            nw283_[int(13)] = len((self).f26)
            nw283_[int(14)] = (self).f24
            nw283_[int(15)] = p0
            nw283_[int(16)] = len(_dafny.SeqWithoutIsStrInference([len(d_1642_v0_)]))
            nw283_[int(17)] = -124
            nw283_[int(18)] = p0
            nw283_[int(19)] = self.f27
            nw283_[int(20)] = self.f27
            nw283_[int(21)] = self.f27
            nw283_[int(22)] = self.f27
            d_1679_v29_ = nw283_
            d_1680_v30_: _dafny.Seq
            d_1680_v30_ = _dafny.SeqWithoutIsStrInference([d_1679_v29_, d_1679_v29_])
            rhs274_ = len(((d_1673_v27_)[self.f27] if (self.f27) in (d_1673_v27_) else d_1671_v25_))
            rhs275_ = default__.fm2((p0) + (p0), p0, d_1670_v24_, globalState)
            rhs276_ = (((0) - (self.f27)) * (len(d_1680_v30_))) < ((self.f27) - ((self).f24))
            rhs277_ = (d_1642_v0_) + (d_1642_v0_)
            rhs278_ = True
            lhs219_ = globalState
            lhs220_ = globalState
            lhs221_ = globalState
            lhs222_ = globalState
            lhs219_.f14 = rhs274_
            d_1664_v20_ = rhs275_
            lhs220_.f2 = rhs276_
            lhs221_.f17 = rhs277_
            lhs222_.f2 = rhs278_
        elif True:
            d_1681_v31_: str
            d_1681_v31_ = _dafny.CodePoint('t')
            d_1681_v31_ = d_1681_v31_
            (globalState).f14 = p0
            d_1682_v32_: bool
            d_1682_v32_ = True
            d_1683_v33_: _dafny.Array
            nw284_ = _dafny.Array(None, 2)
            nw284_[int(0)] = d_1682_v32_
            nw284_[int(1)] = d_1682_v32_
            d_1683_v33_ = nw284_
            d_1684_v34_: _dafny.Map
            d_1684_v34_ = _dafny.Map({d_1683_v33_: not(d_1682_v32_)})
            d_1685_v35_: _dafny.Map
            d_1685_v35_ = _dafny.Map({(self).f24: d_1684_v34_})
            d_1686_v36_: _dafny.Array
            nw285_ = _dafny.Array(None, 22)
            nw285_[int(0)] = (d_1684_v34_) | (d_1684_v34_)
            nw285_[int(1)] = d_1684_v34_
            nw285_[int(2)] = d_1684_v34_
            nw285_[int(3)] = _dafny.Map({d_1683_v33_: d_1682_v32_})
            nw285_[int(4)] = d_1684_v34_
            nw285_[int(5)] = _dafny.Map({d_1683_v33_: default__.fm2((self).f24, self.f27, d_1681_v31_, globalState)})
            nw285_[int(6)] = d_1684_v34_
            nw285_[int(7)] = (d_1684_v34_) | ((d_1684_v34_).set(d_1683_v33_, not(d_1682_v32_)))
            nw285_[int(8)] = _dafny.Map({d_1683_v33_: d_1682_v32_})
            nw285_[int(9)] = _dafny.Map({d_1683_v33_: False})
            nw285_[int(10)] = _dafny.Map({d_1683_v33_: d_1682_v32_})
            nw285_[int(11)] = _dafny.Map({d_1683_v33_: d_1682_v32_})
            nw285_[int(12)] = (d_1684_v34_ if d_1682_v32_ else d_1684_v34_)
            nw285_[int(13)] = (_dafny.Map({d_1683_v33_: d_1682_v32_})) | (_dafny.Map({d_1683_v33_: d_1682_v32_}))
            nw285_[int(14)] = ((d_1685_v35_)[(0) - (p0)] if ((0) - (p0)) in (d_1685_v35_) else d_1684_v34_)
            nw285_[int(15)] = d_1684_v34_
            nw285_[int(16)] = _dafny.Map({d_1683_v33_: not(d_1682_v32_)})
            nw285_[int(17)] = d_1684_v34_
            nw285_[int(18)] = d_1684_v34_
            nw285_[int(19)] = d_1684_v34_
            nw285_[int(20)] = ((_dafny.Map({d_1683_v33_: d_1682_v32_})).set(d_1683_v33_, d_1682_v32_)) | (d_1684_v34_)
            nw285_[int(21)] = d_1684_v34_
            d_1686_v36_ = nw285_
            index246_ = default__.safeIndex(5, (d_1686_v36_).length(0))
            (d_1686_v36_)[index246_] = ((d_1684_v34_).set(d_1683_v33_, d_1682_v32_)).set(d_1683_v33_, d_1682_v32_)
            d_1687_v37_: _dafny.Seq
            d_1687_v37_ = _dafny.SeqWithoutIsStrInference([not(d_1682_v32_), d_1682_v32_])
            d_1688_v38_: _dafny.Map
            d_1688_v38_ = _dafny.Map({d_1687_v37_: self.f27})
            index247_ = default__.safeIndex(5, (d_1686_v36_).length(0))
            (d_1686_v36_)[index247_] = _dafny.Map({d_1683_v33_: (((d_1688_v38_)[default__.fm23(globalState)] if (default__.fm23(globalState)) in (d_1688_v38_) else (self).f24)) < (self.f27)})
            d_1682_v32_ = (d_1687_v37_)[default__.safeIndex((self).f24, len(d_1687_v37_))]
            (globalState).f9 = (self.f27) * ((self).f24)
        d_1689_v39_: bool
        d_1689_v39_ = True
        d_1690_i5_: int
        d_1690_i5_ = 0
        with _dafny.label("7"):
            while d_1689_v39_:
                with _dafny.c_label("7"):
                    if (d_1690_i5_) >= (100):
                        raise _dafny.Break("7")
                    d_1690_i5_ = (d_1690_i5_) + (1)
                    d_1691_v40_: _dafny.Map
                    d_1691_v40_ = _dafny.Map({(self).f24: (self).f24})
                    d_1691_v40_ = (d_1691_v40_).set((self).f24, (len(_dafny.Set({d_1689_v39_, d_1689_v39_, d_1689_v39_}))) - (p0))
                    d_1692_v41_: C0
                    nw286_ = C0()
                    nw286_.ctor__(not (d_1689_v39_) or (d_1689_v39_))
                    d_1692_v41_ = nw286_
                    d_1693_v42_: str
                    d_1693_v42_ = _dafny.CodePoint('o')
                    d_1694_v43_: _dafny.Map
                    d_1694_v43_ = _dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))})
                    d_1695_v44_: _dafny.Seq
                    d_1695_v44_ = _dafny.SeqWithoutIsStrInference([(d_1692_v41_).f34, True])
                    d_1696_v45_: _dafny.Map
                    d_1696_v45_ = _dafny.Map({(d_1695_v44_)[default__.safeIndex((self).f24, len(d_1695_v44_))]: default__.fm3(globalState)})
                    d_1697_v46_: _dafny.Array
                    nw287_ = _dafny.Array(None, 29)
                    nw287_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
                    nw287_[int(1)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgelqg"))) + (d_1642_v0_)
                    nw287_[int(2)] = d_1642_v0_
                    nw287_[int(3)] = default__.fm4(False, p0, d_1693_v42_, default__.fm2((self).f24, (self).f24, d_1693_v42_, globalState), globalState)
                    nw287_[int(4)] = d_1642_v0_
                    nw287_[int(5)] = ((d_1694_v43_)[(d_1692_v41_).f34] if ((d_1692_v41_).f34) in (d_1694_v43_) else _dafny.SeqWithoutIsStrInference([d_1693_v42_ for d_1698_i6_ in range(default__.abs(537))]))
                    nw287_[int(6)] = (d_1642_v0_) + (d_1642_v0_)
                    nw287_[int(7)] = d_1642_v0_
                    nw287_[int(8)] = ((d_1642_v0_).set(default__.safeIndex(len(d_1696_v45_), len(d_1642_v0_)), d_1693_v42_)).set(default__.safeIndex(p0, len((d_1642_v0_).set(default__.safeIndex(len(d_1696_v45_), len(d_1642_v0_)), d_1693_v42_))), _dafny.CodePoint('f'))
                    nw287_[int(9)] = d_1642_v0_
                    nw287_[int(10)] = d_1642_v0_
                    nw287_[int(11)] = (d_1642_v0_) + (d_1642_v0_)
                    nw287_[int(12)] = d_1642_v0_
                    nw287_[int(13)] = (d_1642_v0_) + (d_1642_v0_)
                    nw287_[int(14)] = d_1642_v0_
                    nw287_[int(15)] = _dafny.SeqWithoutIsStrInference([d_1693_v42_ for d_1699_i7_ in range(default__.abs(946))])
                    nw287_[int(16)] = default__.fm4(d_1689_v39_, self.f27, d_1693_v42_, (d_1692_v41_).f34, globalState)
                    nw287_[int(17)] = _dafny.SeqWithoutIsStrInference([d_1693_v42_ for d_1700_i8_ in range(default__.abs(947))])
                    nw287_[int(18)] = (d_1642_v0_) + (d_1642_v0_)
                    nw287_[int(19)] = d_1642_v0_
                    nw287_[int(20)] = (d_1642_v0_) + (d_1642_v0_)
                    nw287_[int(21)] = d_1642_v0_
                    nw287_[int(22)] = d_1642_v0_
                    nw287_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdm"))
                    nw287_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlv"))
                    nw287_[int(25)] = (d_1642_v0_).set(default__.safeIndex((self).f24, len(d_1642_v0_)), d_1693_v42_)
                    nw287_[int(26)] = d_1642_v0_
                    nw287_[int(27)] = d_1642_v0_
                    nw287_[int(28)] = (d_1642_v0_) + (d_1642_v0_)
                    d_1697_v46_ = nw287_
                    index248_ = default__.safeIndex(839, (d_1697_v46_).length(0))
                    (d_1697_v46_)[index248_] = (d_1642_v0_) + (d_1642_v0_)
                    index249_ = default__.safeIndex(839, (d_1697_v46_).length(0))
                    (d_1697_v46_)[index249_] = d_1642_v0_
                    d_1701_v47_: _dafny.Array
                    nw288_ = _dafny.Array(_dafny.Map({}), 18)
                    d_1701_v47_ = nw288_
                    d_1702_v49_: _dafny.Seq
                    d_1702_v49_ = _dafny.SeqWithoutIsStrInference([d_1643_v1_])
                    d_1703_v50_: _dafny.Map
                    d_1703_v50_ = _dafny.Map({_dafny.CodePoint('b'): (self).f24})
                    d_1704_v51_: T0
                    nw289_ = C1()
                    nw289_.ctor__(d_1702_v49_, len(d_1703_v50_), len(d_1642_v0_), (self).f25)
                    d_1704_v51_ = nw289_
                    d_1705_v52_: _dafny.Map
                    d_1705_v52_ = _dafny.Map({-30: d_1704_v51_})
                    d_1706_v53_: _dafny.Seq
                    d_1706_v53_ = _dafny.SeqWithoutIsStrInference([self.f27, p0])
                    d_1707_v54_: _dafny.Seq
                    d_1707_v54_ = _dafny.SeqWithoutIsStrInference([len(d_1705_v52_), (d_1706_v53_)[default__.safeIndex(-383, len(d_1706_v53_))]])
                    d_1708_v55_: D12
                    d_1708_v55_ = D12_DC33(d_1693_v42_, (d_1692_v41_).f34)
                    d_1709_v56_: _dafny.Map
                    d_1709_v56_ = _dafny.Map({(d_1708_v55_).cf54: (d_1692_v41_).f34})
                    d_1710_v57_: _dafny.Seq
                    d_1710_v57_ = _dafny.SeqWithoutIsStrInference([len((d_1707_v54_).set(default__.safeIndex(self.f27, len(d_1707_v54_)), 132)), (self).f24, len(d_1709_v56_), (d_1643_v1_).cardinality, 320])
                    index250_ = default__.safeIndex(918, (d_1701_v47_).length(0))
                    def iife118_():
                        coll50_ = _dafny.Map()
                        compr_50_: int
                        for compr_50_ in (d_1710_v57_).Elements:
                            d_1711_v48_: int = compr_50_
                            if (d_1711_v48_) in (d_1710_v57_):
                                coll50_[(d_1711_v48_) + (9)] = self.f27
                        return _dafny.Map(coll50_)
                    (d_1701_v47_)[index250_] = iife118_()
                    
                    d_1712_v58_: _dafny.Map
                    d_1712_v58_ = _dafny.Map({True: d_1706_v53_})
                    d_1713_v59_: _dafny.Map
                    d_1713_v59_ = _dafny.Map({(d_1692_v41_).f34: ((d_1712_v58_)[d_1689_v39_] if (d_1689_v39_) in (d_1712_v58_) else d_1707_v54_)})
                    d_1714_v60_: _dafny.MultiSet
                    d_1714_v60_ = _dafny.MultiSet([d_1692_v41_])
                    d_1715_v61_: D16
                    d_1715_v61_ = D16_DC41((d_1714_v60_).cardinality)
                    index251_ = default__.safeIndex(918, (d_1701_v47_).length(0))
                    (d_1701_v47_)[index251_] = _dafny.Map({self.f27: (len(d_1713_v59_) if default__.fm2(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvsoofsw"))), p0, d_1693_v42_, globalState) else (d_1715_v61_).cf64)})
                    pass
            pass
        (globalState).f1 = p0
        d_1716_v62_: C0
        nw290_ = C0()
        nw290_.ctor__(d_1689_v39_)
        d_1716_v62_ = nw290_
        d_1717_v63_: _dafny.Array
        nw291_ = _dafny.Array(None, 25)
        nw291_[int(0)] = d_1716_v62_
        nw291_[int(1)] = d_1716_v62_
        nw291_[int(2)] = d_1716_v62_
        nw291_[int(3)] = d_1716_v62_
        nw291_[int(4)] = d_1716_v62_
        nw291_[int(5)] = d_1716_v62_
        nw291_[int(6)] = d_1716_v62_
        nw291_[int(7)] = d_1716_v62_
        nw291_[int(8)] = d_1716_v62_
        nw291_[int(9)] = d_1716_v62_
        nw291_[int(10)] = d_1716_v62_
        nw291_[int(11)] = d_1716_v62_
        nw291_[int(12)] = d_1716_v62_
        nw291_[int(13)] = d_1716_v62_
        nw291_[int(14)] = d_1716_v62_
        nw291_[int(15)] = d_1716_v62_
        nw291_[int(16)] = d_1716_v62_
        nw291_[int(17)] = d_1716_v62_
        nw291_[int(18)] = d_1716_v62_
        nw291_[int(19)] = d_1716_v62_
        nw291_[int(20)] = d_1716_v62_
        nw291_[int(21)] = d_1716_v62_
        nw291_[int(22)] = d_1716_v62_
        nw291_[int(23)] = d_1716_v62_
        nw291_[int(24)] = d_1716_v62_
        d_1717_v63_ = nw291_
        d_1717_v63_ = d_1717_v63_
        d_1718_v64_: _dafny.Array
        def lambda67_(d_1719_p0_):
            def lambda68_(d_1720_i9_):
                return (d_1720_i9_) + (d_1719_p0_)

            return lambda68_

        init40_ = lambda67_(p0)
        nw292_ = _dafny.Array(None, 4)
        for i0_40_ in range(nw292_.length(0)):
            nw292_[i0_40_] = init40_(i0_40_)
        d_1718_v64_ = nw292_
        index252_ = default__.safeIndex(20, (d_1718_v64_).length(0))
        (d_1718_v64_)[index252_] = self.f27
        index253_ = default__.safeIndex(20, (d_1718_v64_).length(0))
        rhs279_ = (len((d_1642_v0_) + (d_1642_v0_))) <= ((self).f24)
        rhs280_ = (0) - (default__.safeDivisionInt(len((d_1642_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tytwp")))), default__.fm0((self).f24, p0, globalState)))
        lhs223_ = d_1718_v64_
        lhs224_ = default__.safeIndex(20, (d_1718_v64_).length(0))
        d_1689_v39_ = rhs279_
        lhs223_[lhs224_] = rhs280_
        d_1721_v65_: str
        d_1721_v65_ = _dafny.CodePoint('q')
        d_1722_i10_: int
        d_1722_i10_ = 0
        with _dafny.label("8"):
            while default__.fm2(p0, p0, d_1721_v65_, globalState):
                with _dafny.c_label("8"):
                    if (d_1722_i10_) >= (100):
                        raise _dafny.Break("8")
                    d_1722_i10_ = (d_1722_i10_) + (1)
                    d_1643_v1_ = d_1643_v1_
                    if (False if not (not(d_1689_v39_)) or ((((self).f26)[d_1689_v39_] if (d_1689_v39_) in ((self).f26) else (d_1716_v62_).f34)) else not((d_1716_v62_).f34)):
                        d_1723_v66_: _dafny.Seq
                        d_1723_v66_ = _dafny.SeqWithoutIsStrInference([(d_1716_v62_).f34])
                        d_1724_v67_: _dafny.Seq
                        d_1724_v67_ = _dafny.SeqWithoutIsStrInference([d_1723_v66_])
                        d_1725_v68_: _dafny.Set
                        d_1725_v68_ = _dafny.Set({(d_1724_v67_)[default__.safeIndex(len(d_1642_v0_), len(d_1724_v67_))], (d_1723_v66_) + (d_1723_v66_), d_1723_v66_})
                        d_1726_v69_: D13
                        d_1726_v69_ = D13_DC36((d_1716_v62_).f34, d_1689_v39_)
                        d_1727_v70_: _dafny.Map
                        d_1727_v70_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1689_v39_]): (d_1716_v62_).f34})
                        d_1728_v72_: _dafny.Set
                        d_1728_v72_ = d_1725_v68_
                        def iife119_():
                            coll51_ = _dafny.Set()
                            compr_51_: _dafny.Seq
                            for compr_51_ in (d_1727_v70_).keys.Elements:
                                d_1729_v71_: _dafny.Seq = compr_51_
                                if (d_1729_v71_) in (d_1727_v70_):
                                    coll51_ = coll51_.union(_dafny.Set([d_1729_v71_]))
                            return _dafny.Set(coll51_)
                        d_1725_v68_ = (d_1725_v68_ if (d_1726_v69_).cf58 else (iife119_()
                        ).intersection((d_1728_v72_)))
                        d_1730_v73_: _dafny.Seq
                        d_1730_v73_ = _dafny.SeqWithoutIsStrInference([d_1643_v1_, d_1643_v1_])
                        d_1731_v74_: _dafny.Map
                        d_1731_v74_ = _dafny.Map({d_1689_v39_: d_1730_v73_})
                        d_1732_v75_: C1
                        nw293_ = C1()
                        nw293_.ctor__(((d_1731_v74_)[(d_1716_v62_).f34] if ((d_1716_v62_).f34) in (d_1731_v74_) else d_1730_v73_), self.f27, (325) - ((self).f24), ((self).f25) + ((self).f25))
                        d_1732_v75_ = nw293_
                        d_1733_v76_: _dafny.Seq
                        d_1733_v76_ = _dafny.SeqWithoutIsStrInference([d_1732_v75_])
                        (globalState).f18 = (0) - ((len((d_1733_v76_) + (_dafny.SeqWithoutIsStrInference([d_1732_v75_, d_1732_v75_])))) * (p0))
                        d_1734_v77_: _dafny.Array
                        nw294_ = _dafny.Array(False, 9)
                        d_1734_v77_ = nw294_
                        d_1735_v78_: _dafny.Array
                        nw295_ = _dafny.Array(None, 12)
                        nw295_[int(0)] = d_1734_v77_
                        nw295_[int(1)] = d_1734_v77_
                        nw295_[int(2)] = d_1734_v77_
                        nw295_[int(3)] = d_1734_v77_
                        nw295_[int(4)] = d_1734_v77_
                        nw295_[int(5)] = d_1734_v77_
                        nw295_[int(6)] = (d_1734_v77_ if False else d_1734_v77_)
                        nw295_[int(7)] = d_1734_v77_
                        nw295_[int(8)] = d_1734_v77_
                        nw295_[int(9)] = d_1734_v77_
                        nw295_[int(10)] = d_1734_v77_
                        nw295_[int(11)] = d_1734_v77_
                        d_1735_v78_ = nw295_
                        index254_ = default__.safeIndex(71, (d_1735_v78_).length(0))
                        (d_1735_v78_)[index254_] = d_1734_v77_
                        index255_ = default__.safeIndex(71, (d_1735_v78_).length(0))
                        (d_1735_v78_)[index255_] = d_1734_v77_
                        (globalState).f1 = d_1732_v75_.f33
                    elif True:
                        d_1736_v79_: _dafny.Set
                        d_1736_v79_ = _dafny.Set({(self).f24})
                        d_1737_v80_: _dafny.Map
                        d_1737_v80_ = _dafny.Map({(d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]: d_1736_v79_})
                        d_1738_v81_: _dafny.Map
                        d_1738_v81_ = _dafny.Map({_dafny.CodePoint('h'): p0})
                        d_1739_v82_: _dafny.Map
                        d_1739_v82_ = _dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1721_v65_: (d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]}), _dafny.Map({(d_1642_v0_)[default__.safeIndex(-891, len(d_1642_v0_))]: (self).f24}), d_1738_v81_, d_1738_v81_, d_1738_v81_])})
                        d_1740_v83_: C2
                        nw296_ = C2()
                        nw296_.ctor__((((d_1737_v80_)[default__.fm0(p0, self.f27, globalState)] if (default__.fm0(p0, self.f27, globalState)) in (d_1737_v80_) else d_1736_v79_)).isdisjoint(d_1736_v79_), (d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))], ((d_1739_v82_)[(d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]] if ((d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]) in (d_1739_v82_) else (self).f25))
                        d_1740_v83_ = nw296_
                        d_1740_v83_ = d_1740_v83_
                        d_1741_v84_: D13
                        d_1741_v84_ = D13_DC35(d_1643_v1_)
                        d_1742_v85_: _dafny.Seq
                        d_1742_v85_ = _dafny.SeqWithoutIsStrInference([self.f27, (d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))], (0) - ((d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))])])
                        (d_1740_v83_).f31 = ((d_1741_v84_).cf57).isdisjoint(_dafny.MultiSet(d_1742_v85_))
                        d_1743_v86_: _dafny.Set
                        d_1743_v86_ = _dafny.Set({d_1689_v39_, d_1740_v83_.f31})
                        d_1744_v87_: _dafny.Seq
                        d_1744_v87_ = _dafny.SeqWithoutIsStrInference([d_1743_v86_])
                        d_1742_v85_ = ((_dafny.SeqWithoutIsStrInference([(d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))] for d_1745_i11_ in range(default__.abs(796))])).set(default__.safeIndex((self).f24, len(_dafny.SeqWithoutIsStrInference([(d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))] for d_1746_i11_ in range(default__.abs(796))]))), self.f27)).set(default__.safeIndex(self.f27, len((_dafny.SeqWithoutIsStrInference([(d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))] for d_1747_i11_ in range(default__.abs(796))])).set(default__.safeIndex((self).f24, len(_dafny.SeqWithoutIsStrInference([(d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))] for d_1748_i11_ in range(default__.abs(796))]))), self.f27))), len(d_1744_v87_))
                        index256_ = default__.safeIndex(20, (d_1718_v64_).length(0))
                        (d_1718_v64_)[index256_] = (p0) - ((d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))])
                        d_1749_v88_: _dafny.Map
                        d_1749_v88_ = _dafny.Map({len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuayu")): (d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]})): d_1740_v83_.f31})
                        d_1750_v89_: _dafny.Array
                        nw297_ = _dafny.Array(D4.default()(), 25)
                        d_1750_v89_ = nw297_
                        d_1751_v90_: _dafny.Map
                        d_1751_v90_ = _dafny.Map({self.f27: (d_1750_v89_ if ((d_1749_v88_)[302] if (302) in (d_1749_v88_) else d_1740_v83_.f31) else d_1750_v89_)})
                        d_1752_v91_: _dafny.Map
                        d_1752_v91_ = _dafny.Map({(d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]: d_1721_v65_})
                        rhs281_ = (d_1751_v90_) | (d_1751_v90_)
                        rhs282_ = (d_1716_v62_).f34
                        rhs283_ = (d_1752_v91_) | (d_1752_v91_)
                        lhs225_ = globalState
                        d_1751_v90_ = rhs281_
                        lhs225_.f2 = rhs282_
                        d_1752_v91_ = rhs283_
                    if (((d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))] if d_1689_v39_ else (self).f24)) == ((d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]):
                        d_1689_v39_ = (d_1716_v62_).f34
                        d_1753_v92_: _dafny.Seq
                        d_1753_v92_ = _dafny.SeqWithoutIsStrInference([d_1718_v64_])
                        d_1754_v93_: _dafny.Array
                        nw298_ = _dafny.Array(None, 7)
                        nw298_[int(0)] = d_1718_v64_
                        nw298_[int(1)] = d_1718_v64_
                        nw298_[int(2)] = d_1718_v64_
                        nw298_[int(3)] = d_1718_v64_
                        nw298_[int(4)] = d_1718_v64_
                        nw298_[int(5)] = (d_1753_v92_)[default__.safeIndex(self.f27, len(d_1753_v92_))]
                        nw298_[int(6)] = d_1718_v64_
                        d_1754_v93_ = nw298_
                        index257_ = default__.safeIndex(961, (d_1754_v93_).length(0))
                        (d_1754_v93_)[index257_] = d_1718_v64_
                        d_1755_v94_: _dafny.Seq
                        d_1755_v94_ = _dafny.SeqWithoutIsStrInference([-469])
                        d_1756_v95_: _dafny.Seq
                        d_1756_v95_ = _dafny.SeqWithoutIsStrInference([True, True])
                        d_1757_v96_: _dafny.Seq
                        d_1757_v96_ = _dafny.SeqWithoutIsStrInference([d_1756_v95_])
                        index258_ = default__.safeIndex(20, (d_1718_v64_).length(0))
                        index259_ = default__.safeIndex(961, (d_1754_v93_).length(0))
                        rhs284_ = default__.safeDivisionInt(default__.safeModuloInt(-763, len((d_1755_v94_).set(default__.safeIndex(len(default__.fm6(len(d_1755_v94_), (self).f24, (self).f24, self.f27, globalState)), len(d_1755_v94_)), p0))), (423) * (self.f27))
                        rhs285_ = (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len((d_1757_v96_)[default__.safeIndex((d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))], len(d_1757_v96_))])]))})).isdisjoint(_dafny.Set({(d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]}))
                        rhs286_ = d_1718_v64_
                        rhs287_ = d_1642_v0_
                        rhs288_ = d_1642_v0_
                        lhs226_ = d_1718_v64_
                        lhs227_ = default__.safeIndex(20, (d_1718_v64_).length(0))
                        lhs228_ = globalState
                        lhs229_ = d_1754_v93_
                        lhs230_ = default__.safeIndex(961, (d_1754_v93_).length(0))
                        lhs231_ = globalState
                        lhs226_[lhs227_] = rhs284_
                        lhs228_.f2 = rhs285_
                        lhs229_[lhs230_] = rhs286_
                        lhs231_.f17 = rhs287_
                        d_1642_v0_ = rhs288_
                        d_1758_v97_: _dafny.Array
                        nw299_ = _dafny.Array(_dafny.Seq({}), 9)
                        d_1758_v97_ = nw299_
                        d_1759_v98_: _dafny.Seq
                        d_1759_v98_ = _dafny.SeqWithoutIsStrInference([d_1758_v97_, d_1758_v97_])
                        d_1758_v97_ = (d_1759_v98_)[default__.safeIndex((0) - ((self).f24), len(d_1759_v98_))]
                        d_1760_v99_: int
                        out33_: int
                        out33_ = default__.m0(p0, (self).f24, globalState)
                        d_1760_v99_ = out33_
                        d_1761_v100_: _dafny.Map
                        d_1761_v100_ = _dafny.Map({True: default__.fm23(globalState)})
                        d_1762_v101_: _dafny.Map
                        d_1762_v101_ = _dafny.Map({(default__.fm1(((d_1643_v1_).set(self.f27, default__.abs((0) - (self.f27)))).cardinality, default__.fm0(len(((d_1761_v100_)[(d_1716_v62_).f34] if ((d_1716_v62_).f34) in (d_1761_v100_) else _dafny.SeqWithoutIsStrInference([(d_1716_v62_).f34]))), self.f27, globalState), globalState)).cardinality: d_1760_v99_})
                        d_1763_v103_: _dafny.Seq
                        def iife120_():
                            coll52_ = _dafny.Map()
                            compr_52_: int
                            for compr_52_ in _dafny.IntegerRange(593, 594):
                                d_1764_v102_: int = compr_52_
                                if ((593) <= (d_1764_v102_)) and ((d_1764_v102_) < (594)):
                                    coll52_[(d_1764_v102_) + (d_1760_v99_)] = (0) - ((d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))])
                            return _dafny.Map(coll52_)
                        d_1763_v103_ = _dafny.SeqWithoutIsStrInference([iife120_()
                        , d_1762_v101_])
                        d_1765_v104_: _dafny.Map
                        d_1765_v104_ = _dafny.Map({d_1689_v39_: p0})
                        d_1766_v105_: D3
                        d_1766_v105_ = D3_DC9(d_1756_v95_)
                        d_1767_v106_: D1
                        d_1767_v106_ = D1_DC4((d_1716_v62_).f34, d_1689_v39_, d_1721_v65_, (self).f24, len((d_1766_v105_).cf17))
                        d_1768_v108_: _dafny.Array
                        nw300_ = _dafny.Array(None, 22)
                        nw300_[int(0)] = d_1762_v101_
                        nw300_[int(1)] = (d_1762_v101_) | (d_1762_v101_)
                        nw300_[int(2)] = (d_1763_v103_)[default__.safeIndex((self).f24, len(d_1763_v103_))]
                        nw300_[int(3)] = _dafny.Map({(0) - (self.f27): (self).f24})
                        nw300_[int(4)] = (_dafny.Map({871: d_1760_v99_})) | (d_1762_v101_)
                        nw300_[int(5)] = (d_1762_v101_) | (d_1762_v101_)
                        nw300_[int(6)] = d_1762_v101_
                        nw300_[int(7)] = _dafny.Map({self.f27: (0) - (p0)})
                        nw300_[int(8)] = (d_1762_v101_).set(187, (d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))])
                        nw300_[int(9)] = d_1762_v101_
                        nw300_[int(10)] = (d_1762_v101_ if True else _dafny.Map({((d_1765_v104_)[d_1689_v39_] if (d_1689_v39_) in (d_1765_v104_) else (d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]): p0}))
                        nw300_[int(11)] = d_1762_v101_
                        nw300_[int(12)] = d_1762_v101_
                        nw300_[int(13)] = _dafny.Map({len(d_1756_v95_): d_1760_v99_})
                        nw300_[int(14)] = (d_1762_v101_ if d_1689_v39_ else default__.fm15((d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))], (d_1767_v106_).cf8, -368, (d_1716_v62_).f34, globalState))
                        def iife121_():
                            coll53_ = _dafny.Map()
                            compr_53_: int
                            for compr_53_ in _dafny.IntegerRange(794, -694):
                                d_1769_v107_: int = compr_53_
                                if ((794) <= (d_1769_v107_)) and ((d_1769_v107_) < (-694)):
                                    coll53_[(d_1769_v107_) - (p0)] = self.f27
                            return _dafny.Map(coll53_)
                        nw300_[int(15)] = (_dafny.Map({self.f27: len(d_1755_v94_)})) | (iife121_()
                        )
                        nw300_[int(16)] = d_1762_v101_
                        nw300_[int(17)] = d_1762_v101_
                        nw300_[int(18)] = d_1762_v101_
                        nw300_[int(19)] = d_1762_v101_
                        nw300_[int(20)] = _dafny.Map({d_1760_v99_: (self).f24})
                        nw300_[int(21)] = d_1762_v101_
                        d_1768_v108_ = nw300_
                        index260_ = default__.safeIndex(174, (d_1768_v108_).length(0))
                        (d_1768_v108_)[index260_] = d_1762_v101_
                        index261_ = default__.safeIndex(174, (d_1768_v108_).length(0))
                        (d_1768_v108_)[index261_] = _dafny.Map({self.f27: (d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]})
                    elif True:
                        d_1721_v65_ = _dafny.CodePoint('v')
                        d_1770_v110_: _dafny.MultiSet
                        d_1770_v110_ = _dafny.MultiSet([d_1721_v65_])
                        d_1771_v112_: C2
                        nw301_ = C2()
                        def iife122_():
                            coll54_ = _dafny.Map()
                            compr_54_: str
                            for compr_54_ in (d_1770_v110_).Elements:
                                d_1774_v109_: str = compr_54_
                                if (d_1774_v109_) in (d_1770_v110_):
                                    coll54_[d_1774_v109_] = (d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]
                            return _dafny.Map(coll54_)
                        def iife123_():
                            coll55_ = _dafny.Map()
                            compr_55_: str
                            for compr_55_ in (d_1770_v110_).Elements:
                                d_1775_v111_: str = compr_55_
                                if (d_1775_v111_) in (d_1770_v110_):
                                    coll55_[d_1775_v111_] = (self).f24
                            return _dafny.Map(coll55_)
                        nw301_.ctor__(((0) - ((self).f24)) == (318), (0) - (((d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]) - ((self).f24)), ((_dafny.SeqWithoutIsStrInference([_dafny.Map({d_1721_v65_: (d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]}) for d_1772_i12_ in range(default__.abs(244))])).set(default__.safeIndex(default__.fm0((0) - (p0), p0, globalState), len(_dafny.SeqWithoutIsStrInference([_dafny.Map({d_1721_v65_: (d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]}) for d_1773_i12_ in range(default__.abs(244))]))), iife122_()
                        )) + (_dafny.SeqWithoutIsStrInference([iife123_()
 for d_1776_i13_ in range(default__.abs(586))])))
                        d_1771_v112_ = nw301_
                        (globalState).f2 = (d_1716_v62_).f34
                        (self).f27 = (p0) * (default__.safeModuloInt((self).f24, (d_1643_v1_).cardinality))
                        (globalState).f2 = d_1771_v112_.f31
                    d_1777_v114_: _dafny.Seq
                    d_1777_v114_ = _dafny.SeqWithoutIsStrInference([(d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]])
                    d_1778_v115_: _dafny.Map
                    d_1778_v115_ = _dafny.Map({len(d_1777_v114_): p0})
                    d_1779_v116_: _dafny.Map
                    d_1779_v116_ = _dafny.Map({d_1721_v65_: d_1778_v115_})
                    d_1780_v117_: _dafny.Map
                    d_1780_v117_ = _dafny.Map({(d_1718_v64_)[default__.safeIndex(20, (d_1718_v64_).length(0))]: (d_1716_v62_).f34})
                    d_1781_v118_: _dafny.Set
                    def iife124_():
                        coll56_ = _dafny.Map()
                        compr_56_: int
                        for compr_56_ in (((d_1779_v116_)[d_1721_v65_] if (d_1721_v65_) in (d_1779_v116_) else d_1778_v115_)).keys.Elements:
                            d_1782_v113_: int = compr_56_
                            if (d_1782_v113_) in (((d_1779_v116_)[d_1721_v65_] if (d_1721_v65_) in (d_1779_v116_) else d_1778_v115_)):
                                coll56_[(d_1782_v113_) + (self.f27)] = (d_1716_v62_).f34
                        return _dafny.Map(coll56_)
                    d_1781_v118_ = _dafny.Set({self.f27, (0) - (len((iife124_()
                    ) | (d_1780_v117_)))})
                    (globalState).f18 = len(d_1781_v118_)
                    pass
            pass

    @property
    def f26(self):
        return self._f26
