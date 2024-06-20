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
        return (834) != ((-650) * (len(_dafny.SeqWithoutIsStrInference([not(True)]))))

    @staticmethod
    def fm1(globalState):
        def lambda0_(source0_):
            d_0___mcc_h0_ = source0_
            d_1_cf0_ = d_0___mcc_h0_
            return (_dafny.Set({d_1_cf0_})) | (_dafny.Set({d_1_cf0_, d_1_cf0_, d_1_cf0_}))

        return (0) - (len(lambda0_(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2_i0_ in range(default__.abs(597))]))))

    @staticmethod
    def fm8(globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpftgh"))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_3_i0_ in range(default__.abs(639))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jimxsttfl"))))

    @staticmethod
    def fm9(p0, globalState):
        source1_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])})
        d_4___mcc_h0_ = source1_
        d_5_cf62_ = d_4___mcc_h0_
        return _dafny.SeqWithoutIsStrInference([875 for d_6_i0_ in range(default__.abs(796))])

    @staticmethod
    def fm10(globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in (_dafny.Map({786: len(_dafny.Map({-631: -19}))})).keys.Elements:
                d_7_v0_: int = compr_0_
                if (d_7_v0_) in (_dafny.Map({786: len(_dafny.Map({-631: -19}))})):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeDivisionInt(d_7_v0_, 389)]))
            return _dafny.Set(coll0_)
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suq"))): not(False)})) for d_8_i0_ in range(default__.abs(-219))])).Elements:
                d_9_v1_: int = compr_1_
                if (d_9_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suq"))): not(False)})) for d_8_i0_ in range(default__.abs(-219))])):
                    coll1_ = coll1_.union(_dafny.Set([(d_9_v1_) + (199)]))
            return _dafny.Set(coll1_)
        return ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouixhpwjj"))), len((D9_DC20(len(_dafny.SeqWithoutIsStrInference([False, True, False, False, False])), True, False, False, _dafny.Map({len(_dafny.Set({315})): 106}))).cf38), len(_dafny.Map({-680: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f')])}))}) if False else (D9_DC17(iife0_()
)).cf31)) | ((_dafny.Set({531, len(_dafny.SeqWithoutIsStrInference([False, True, False])), len(_dafny.Set({not(False)}))})).intersection(iife1_()
        ))

    @staticmethod
    def fm12(p0, p1, globalState):
        if True:
            if True:
                return _dafny.CodePoint('f')
            elif True:
                return _dafny.CodePoint('r')
        elif True:
            return _dafny.CodePoint('t')

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False, False, True, False, False])) + ((D7_DC13(_dafny.SeqWithoutIsStrInference([True, True]))).cf23)) + ((_dafny.SeqWithoutIsStrInference([False, True]) if False else (D7_DC13(_dafny.SeqWithoutIsStrInference([False]))).cf23))

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        source2_ = D18_DC46(_dafny.Set({_dafny.Map({False: False})}))
        if source2_.is_DC47:
            d_10___mcc_h0_ = source2_.cf86
            d_11___mcc_h1_ = source2_.cf87
            d_12_cf87_ = d_11___mcc_h1_
            d_13_cf86_ = d_10___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obdympjc")))
        elif True:
            d_14___mcc_h2_ = source2_.cf85
            d_15_cf85_ = d_14___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_16_i0_ in range(default__.abs(886))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rw")))

    @staticmethod
    def fm15(p0, p1, globalState):
        return _dafny.Map({(-721) + (len(_dafny.Map({False: (_dafny.MultiSet([771, (_dafny.MultiSet([True])).cardinality])).cardinality}))): len((_dafny.SeqWithoutIsStrInference([True, not((D16_DC40(True, 499, (0) - (len(_dafny.Set({False}))))).cf74)])) + (_dafny.SeqWithoutIsStrInference([True, True, False])))})

    @staticmethod
    def fm16(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([767])) + (_dafny.SeqWithoutIsStrInference([676]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, True}))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): False})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otnqi")))}))])))

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        source3_ = D15_DC36(_dafny.CodePoint('c'), 604)
        if source3_.is_DC36:
            d_17___mcc_h0_ = source3_.cf64
            d_18___mcc_h1_ = source3_.cf65
            d_19_cf65_ = d_18___mcc_h1_
            d_20_cf64_ = d_17___mcc_h0_
            return _dafny.Set({False, not(True), not(True), not(True)})
        elif source3_.is_DC37:
            d_21___mcc_h2_ = source3_.cf66
            d_22___mcc_h3_ = source3_.cf67
            d_23___mcc_h4_ = source3_.cf68
            d_24___mcc_h5_ = source3_.cf69
            d_25_cf69_ = d_24___mcc_h5_
            d_26_cf68_ = d_23___mcc_h4_
            d_27_cf67_ = d_22___mcc_h3_
            d_28_cf66_ = d_21___mcc_h2_
            return (_dafny.Set({d_28_cf66_})).intersection(_dafny.Set({d_25_cf69_}))
        elif source3_.is_DC38:
            d_29___mcc_h6_ = source3_.cf70
            d_30___mcc_h7_ = source3_.cf71
            d_31___mcc_h8_ = source3_.cf72
            d_32_cf72_ = d_31___mcc_h8_
            d_33_cf71_ = d_30___mcc_h7_
            d_34_cf70_ = d_29___mcc_h6_
            return (D24_DC58(_dafny.Set({d_33_cf71_, d_33_cf71_}))).cf109
        elif True:
            d_35___mcc_h9_ = source3_.cf63
            d_36_cf63_ = d_35___mcc_h9_
            return _dafny.Set({True})

    @staticmethod
    def fm18(globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_37_i0_ in range(default__.abs(936))])))]), _dafny.SeqWithoutIsStrInference([-501])])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([878])]))

    @staticmethod
    def fm19(globalState):
        if (896) >= (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(True): False}))]))):
            return _dafny.Set({_dafny.CodePoint('t')})
        elif True:
            return _dafny.Set({_dafny.CodePoint('s'), _dafny.CodePoint('l'), _dafny.CodePoint('h'), _dafny.CodePoint('k'), _dafny.CodePoint('g')})

    @staticmethod
    def fm20(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False]), (_dafny.SeqWithoutIsStrInference([False]) if True else _dafny.SeqWithoutIsStrInference([not(True)]))])

    @staticmethod
    def fm23(p0, p1, globalState):
        return (_dafny.Map({True: False})) | ((_dafny.Map({False: not(True)}) if False else _dafny.Map({True: False})))

    @staticmethod
    def fm24(p0, globalState):
        return ((_dafny.Map({not(not(True)): 295})) | (_dafny.Map({False: 856}))) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skaroir")))}))

    @staticmethod
    def fm25(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False, False])) + (_dafny.SeqWithoutIsStrInference([False, not(True)]))) + (_dafny.SeqWithoutIsStrInference([not(True), False]))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        source4_ = _dafny.MultiSet([520, -680])
        d_38___mcc_h0_ = source4_
        d_39_cf18_ = d_38___mcc_h0_
        return _dafny.CodePoint('h')

    @staticmethod
    def fm27(p0, globalState):
        return (_dafny.MultiSet([264])) | (_dafny.MultiSet([len((D9_DC17(_dafny.Set({743}))).cf31), len(_dafny.SeqWithoutIsStrInference([_dafny.Set({True}) for d_40_i0_ in range(default__.abs(307))]))]))

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return D1_DC1(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, True])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))})): True}))

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        return (True if True else False)

    @staticmethod
    def fm30(globalState):
        return ((_dafny.Map({229: len(_dafny.Map({_dafny.Map({978: True}): False}))})) | (_dafny.Map({138: 704}))) | (_dafny.Map({-387: 884}))

    @staticmethod
    def fm31(p0, p1, globalState):
        source5_ = D9_DC18(not(True))
        if source5_.is_DC18:
            d_41___mcc_h0_ = source5_.cf32
            d_42_cf32_ = d_41___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpb"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypsbjt")))
        elif source5_.is_DC19:
            d_43___mcc_h1_ = source5_.cf33
            d_44_cf33_ = d_43___mcc_h1_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xa"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxt")))
        elif source5_.is_DC20:
            d_45___mcc_h2_ = source5_.cf34
            d_46___mcc_h3_ = source5_.cf35
            d_47___mcc_h4_ = source5_.cf36
            d_48___mcc_h5_ = source5_.cf37
            d_49___mcc_h6_ = source5_.cf38
            d_50_cf38_ = d_49___mcc_h6_
            d_51_cf37_ = d_48___mcc_h5_
            d_52_cf36_ = d_47___mcc_h4_
            d_53_cf35_ = d_46___mcc_h3_
            d_54_cf34_ = d_45___mcc_h2_
            if d_51_cf37_:
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eg"))
            elif True:
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uj"))
        elif True:
            d_55___mcc_h7_ = source5_.cf31
            d_56_cf31_ = d_55___mcc_h7_
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_57_i0_ in range(default__.abs(885))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_58_i1_ in range(default__.abs(865))]))

    @staticmethod
    def fm32(globalState):
        return _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaed"))), (_dafny.MultiSet([672])).isdisjoint(_dafny.MultiSet([573, -43, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, True]))])).cardinality])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))), 950])), 615, -460]))])

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, True])) + (_dafny.SeqWithoutIsStrInference([False, True]))) + (_dafny.SeqWithoutIsStrInference([True, not(False)]))

    @staticmethod
    def fm36(p0, globalState):
        source6_ = D20_DC51()
        if source6_.is_DC51:
            if not(True):
                return D2_DC4((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality, 608, -857, _dafny.CodePoint('f'), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_59_i0_ in range(default__.abs(199))]))
            elif True:
                return D2_DC4(478, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crgitqdfk"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlk"))), _dafny.CodePoint('k'), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovgnevf")))
        elif True:
            d_60___mcc_h0_ = source6_.cf92
            d_61_cf92_ = d_60___mcc_h0_
            return D2_DC4((0) - (-718), 896, -841, _dafny.CodePoint('u'), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_62_i1_ in range(default__.abs(-641))]))

    @staticmethod
    def fm37(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_63_i0_ in range(default__.abs(659))])

    @staticmethod
    def fm38(p0, globalState):
        return _dafny.CodePoint('g')

    @staticmethod
    def fm39(globalState):
        def iife2_():
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: _dafny.Seq
                for compr_4_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False})).keys.Elements:
                    d_64_v0_: _dafny.Seq = compr_4_
                    if (d_64_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False})):
                        coll4_[d_64_v0_] = not(False)
                return _dafny.Map(coll4_)
            coll2_ = _dafny.Set()
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: _dafny.Seq
                for compr_3_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False})).keys.Elements:
                    d_64_v0_: _dafny.Seq = compr_3_
                    if (d_64_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False})):
                        coll3_[d_64_v0_] = not(False)
                return _dafny.Map(coll3_)
            compr_2_: _dafny.Seq
            for compr_2_ in (iife3_()
            ).keys.Elements:
                d_65_v1_: _dafny.Seq = compr_2_
                if (d_65_v1_) in (iife4_()
                ):
                    coll2_ = coll2_.union(_dafny.Set([d_65_v1_]))
            return _dafny.Set(coll2_)
        source7_ = iife2_()

        d_66___mcc_h0_ = source7_
        d_67_cf62_ = d_66___mcc_h0_
        return D11_DC24((_dafny.MultiSet([False, False])).cardinality, D4_DC8(_dafny.CodePoint('n'), 190, False))

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofqxguj")): _dafny.SeqWithoutIsStrInference([True, True])})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vh")): _dafny.SeqWithoutIsStrInference([not(not(False))])}))

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(True)])) + ((_dafny.SeqWithoutIsStrInference([False, False])) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        return _dafny.Map({(len(_dafny.Map({_dafny.CodePoint('j'): False})) if False else 213): default__.safeDivisionInt(661, 501)})

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: str
            for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('p'), _dafny.CodePoint('a'), _dafny.CodePoint('x'), _dafny.CodePoint('d')])).Elements:
                d_68_v0_: str = compr_5_
                if (d_68_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('p'), _dafny.CodePoint('a'), _dafny.CodePoint('x'), _dafny.CodePoint('d')])):
                    coll5_ = coll5_.union(_dafny.Set([d_68_v0_]))
            return _dafny.Set(coll5_)
        return ((_dafny.SeqWithoutIsStrInference([len(iife5_()
) for d_69_i0_ in range(default__.abs(545))])) + (_dafny.SeqWithoutIsStrInference([676, -823]))) + ((_dafny.SeqWithoutIsStrInference([-438 for d_70_i1_ in range(default__.abs(313))])) + (_dafny.SeqWithoutIsStrInference([631])))

    @staticmethod
    def fm44(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, True, False]) for d_71_i0_ in range(default__.abs(-592))])

    @staticmethod
    def fm47(globalState):
        return (_dafny.Map({_dafny.Set({len(_dafny.Set({False}))}): True})) | (_dafny.Map({_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([D6_DC12(False, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True]), _dafny.SeqWithoutIsStrInference([True])]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvltui"))), D6_DC12(False, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])]), (D6_DC12(False, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(True), True])]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyruvhv")))).cf22)]))]))}): not(True)}))

    @staticmethod
    def fm48(p0, p1, p2, globalState):
        return (D1_DC1(_dafny.Map({len(_dafny.Set({876, 396})): True}))).cf1

    @staticmethod
    def fm49(p0, globalState):
        if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([587]))).ispropersubset((_dafny.MultiSet([-927]))):
            return D4_DC9(len(_dafny.SeqWithoutIsStrInference([False, True, True])), not(False))
        elif True:
            return D4_DC9(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okgwkrflp"))), True)

    @staticmethod
    def fm50(globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(996, 410):
                d_73_v0_: int = compr_6_
                if ((996) <= (d_73_v0_)) and ((d_73_v0_) < (410)):
                    coll6_[default__.safeDivisionInt(d_73_v0_, 72)] = False
            return _dafny.Map(coll6_)
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([-469, (_dafny.MultiSet([_dafny.CodePoint('u')])).cardinality]), _dafny.SeqWithoutIsStrInference([(0) - ((0) - (-827))])})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([len((D24_DC58(_dafny.Set({False, True}))).cf109) for d_72_i0_ in range(default__.abs(256))]), _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({True: (0) - (len(_dafny.Map({False: 880})))}))), 942, (_dafny.MultiSet([-878, 888])).cardinality]), _dafny.SeqWithoutIsStrInference([len(iife6_()
        )])}))

    @staticmethod
    def fm51(p0, p1, p2, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: str
            for compr_7_ in (_dafny.MultiSet([_dafny.CodePoint('n'), _dafny.CodePoint('u')])).Elements:
                d_74_v0_: str = compr_7_
                if (d_74_v0_) in (_dafny.MultiSet([_dafny.CodePoint('n'), _dafny.CodePoint('u')])):
                    coll7_[d_74_v0_] = True
            return _dafny.Map(coll7_)
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: str
            for compr_8_ in (_dafny.Map({_dafny.CodePoint('y'): False})).keys.Elements:
                d_75_v1_: str = compr_8_
                if (d_75_v1_) in (_dafny.Map({_dafny.CodePoint('y'): False})):
                    coll8_[d_75_v1_] = True
            return _dafny.Map(coll8_)
        return _dafny.SeqWithoutIsStrInference([(_dafny.Map({_dafny.CodePoint('f'): True})) | (_dafny.Map({_dafny.CodePoint('n'): not(False)})), (iife7_()
        ) | (iife8_()
        )])

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return D9_DC18(False)

    @staticmethod
    def fm53(p0, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([False, True, True])) + (_dafny.SeqWithoutIsStrInference([False])): (_dafny.Set({False})).ispropersubset(_dafny.Set({True, False, True, True, False}))})

    @staticmethod
    def fm54(p0, p1, globalState):
        return _dafny.Map({(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_76_i0_ in range(default__.abs(-113))])): 80})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n')])): (0) - (len(_dafny.SeqWithoutIsStrInference([524, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_77_i1_ in range(default__.abs(-541))])))])))})): len(_dafny.Set({868, -924}))})

    @staticmethod
    def fm55(p0, p1, p2, p3, globalState):
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: str
            for compr_9_ in (_dafny.Set({_dafny.CodePoint('q'), _dafny.CodePoint('s')})).Elements:
                d_78_v0_: str = compr_9_
                if (d_78_v0_) in (_dafny.Set({_dafny.CodePoint('q'), _dafny.CodePoint('s')})):
                    coll9_[d_78_v0_] = False
            return _dafny.Map(coll9_)
        return D16_DC40(not(not(True)), len((iife9_()
) | (_dafny.Map({_dafny.CodePoint('a'): True}))), len(_dafny.Map({415: 651})))

    @staticmethod
    def fm56(p0, p1, p2, globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_79_i0_ in range(default__.abs(82))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgydj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxhrte"))})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqymdm"))}))).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hydj"))}))

    @staticmethod
    def fm57(p0, p1, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(364, 789):
                d_80_v0_: int = compr_10_
                if ((364) <= (d_80_v0_)) and ((d_80_v0_) < (789)):
                    coll10_[(d_80_v0_) * (171)] = _dafny.MultiSet([False])
            return _dafny.Map(coll10_)
        return D15_DC35(_dafny.SeqWithoutIsStrInference([_dafny.Map({(0) - (len(iife10_()
)): _dafny.CodePoint('q')})]))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        pat_let_tv0_ = p1
        def iife11_(_pat_let0_0):
            def iife12_(d_81_dt__update__tmp_h0_):
                def iife13_(_pat_let1_0):
                    def iife14_(d_82_dt__update_hcf56_h0_):
                        return D13_DC31(d_82_dt__update_hcf56_h0_)
                    return iife14_(_pat_let1_0)
                return iife13_((0) - ((0) - (pat_let_tv0_)))
            return iife12_(_pat_let0_0)
        source8_ = iife11_(D13_DC31(p2))
        if source8_.is_DC31:
            d_83___mcc_h0_ = source8_.cf56
            d_84_cf56_ = d_83___mcc_h0_
            d_85_v0_: _dafny.Array
            def lambda1_(d_86_i0_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqcmjr"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_87_i1_ in range(default__.abs(180))]))

            init0_ = lambda1_
            nw0_ = _dafny.Array(None, 14)
            for i0_0_ in range(nw0_.length(0)):
                nw0_[i0_0_] = init0_(i0_0_)
            d_85_v0_ = nw0_
            d_85_v0_ = d_85_v0_
            d_88_v1_: D4
            d_88_v1_ = D4_DC9(-389, p0)
            d_88_v1_ = (d_88_v1_ if p0 else d_88_v1_)
            if default__.fm0((p3) - (d_84_cf56_), globalState):
                d_89_v2_: _dafny.Set
                d_89_v2_ = _dafny.Set({p2, p2})
                d_90_v3_: _dafny.Seq
                d_90_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqrd"))
                d_91_v4_: _dafny.Map
                d_91_v4_ = _dafny.Map({True: d_90_v3_})
                d_92_v5_: _dafny.MultiSet
                d_92_v5_ = _dafny.MultiSet([p0, p0, p0, p0, p0])
                d_93_v6_: D10
                d_93_v6_ = D10_DC22(len(d_91_v4_), d_92_v5_)
                d_94_v7_: _dafny.Seq
                d_94_v7_ = _dafny.SeqWithoutIsStrInference([p0])
                d_95_v8_: _dafny.Map
                d_95_v8_ = _dafny.Map({p1: d_85_v0_})
                d_96_v9_: C1
                nw1_ = C1()
                nw1_.ctor__((d_89_v2_).isdisjoint(default__.fm10(globalState)), (d_93_v6_).cf41, default__.fm16(p3, globalState), len(d_94_v7_), ((d_95_v8_)[471] if (471) in (d_95_v8_) else d_85_v0_))
                d_96_v9_ = nw1_
                d_97_v10_: C0
                nw2_ = C0()
                nw2_.ctor__(d_84_cf56_, (d_90_v3_) + (d_90_v3_))
                d_97_v10_ = nw2_
                d_97_v10_ = d_97_v10_
                d_98_v12_: T0
                nw3_ = C3()
                nw3_.ctor__(p1, d_85_v0_)
                d_98_v12_ = nw3_
                d_99_v13_: _dafny.MultiSet
                d_99_v13_ = _dafny.MultiSet([d_98_v12_, d_98_v12_])
                d_100_v14_: _dafny.MultiSet
                d_100_v14_ = _dafny.MultiSet([len(d_90_v3_), (d_99_v13_).cardinality])
                d_101_v15_: _dafny.Seq
                d_101_v15_ = _dafny.SeqWithoutIsStrInference([(d_97_v10_).f36, p2])
                d_102_v17_: _dafny.Map
                d_102_v17_ = _dafny.Map({d_96_v9_.f35: p0})
                def iife15_():
                    coll11_ = _dafny.Map()
                    compr_11_: int
                    for compr_11_ in (d_100_v14_).Elements:
                        d_103_v11_: int = compr_11_
                        if (d_103_v11_) in (d_100_v14_):
                            coll11_[default__.safeDivisionInt(d_103_v11_, (d_97_v10_).f36)] = len((d_97_v10_).f37)
                    return _dafny.Map(coll11_)
                def iife16_():
                    coll12_ = _dafny.Set()
                    compr_12_: int
                    for compr_12_ in _dafny.IntegerRange(719, 727):
                        d_104_v16_: int = compr_12_
                        if ((719) <= (d_104_v16_)) and ((d_104_v16_) < (727)):
                            coll12_ = coll12_.union(_dafny.Set([(d_104_v16_) * (105)]))
                    return _dafny.Set(coll12_)
                (globalState).f3 = (default__.safeModuloInt(len((iife15_()
                ).set((d_101_v15_)[default__.safeIndex((0) - (p3), len(d_101_v15_))], p2)), len(iife16_()
                ))) * ((len(d_102_v17_)) + (d_84_cf56_))
                d_105_v18_: C6
                nw4_ = C6()
                nw4_.ctor__((d_96_v9_).fm3(d_92_v5_, globalState), -792, (d_92_v5_) - (default__.fm32(globalState)), d_101_v15_, p2, d_85_v0_, (d_93_v6_).cf41, len(default__.fm8(globalState)))
                d_105_v18_ = nw4_
                d_106_v19_: _dafny.Array
                nw5_ = _dafny.Array(int(0), 14)
                d_106_v19_ = nw5_
                index0_ = default__.safeIndex(565, (d_106_v19_).length(0))
                (d_106_v19_)[index0_] = (d_100_v14_).cardinality
                index1_ = default__.safeIndex(565, (d_106_v19_).length(0))
                (d_106_v19_)[index1_] = ((d_97_v10_).f36 if True else ((d_92_v5_)[(d_105_v18_).f33] if ((d_105_v18_).f33) in (d_92_v5_) else d_105_v18_.f34))
            elif True:
                d_84_cf56_ = default__.safeDivisionInt(p2, p1)
                d_107_v20_: str
                d_107_v20_ = _dafny.CodePoint('s')
                d_108_v21_: _dafny.Set
                d_108_v21_ = _dafny.Set({d_107_v20_})
                d_109_v24_: _dafny.Array
                nw6_ = _dafny.Array(None, 19)
                nw6_[int(0)] = d_108_v21_
                nw6_[int(1)] = d_108_v21_
                nw6_[int(2)] = d_108_v21_
                nw6_[int(3)] = d_108_v21_
                nw6_[int(4)] = default__.fm19(globalState)
                nw6_[int(5)] = _dafny.Set({d_107_v20_})
                nw6_[int(6)] = d_108_v21_
                nw6_[int(7)] = d_108_v21_
                nw6_[int(8)] = d_108_v21_
                nw6_[int(9)] = d_108_v21_
                nw6_[int(10)] = d_108_v21_
                def iife17_():
                    def iife19_():
                        coll15_ = _dafny.Map()
                        compr_15_: str
                        for compr_15_ in (_dafny.Map({d_107_v20_: d_107_v20_})).keys.Elements:
                            d_110_v22_: str = compr_15_
                            if (d_110_v22_) in (_dafny.Map({d_107_v20_: d_107_v20_})):
                                coll15_[d_110_v22_] = False
                        return _dafny.Map(coll15_)
                    coll13_ = _dafny.Set()
                    def iife18_():
                        coll14_ = _dafny.Map()
                        compr_14_: str
                        for compr_14_ in (_dafny.Map({d_107_v20_: d_107_v20_})).keys.Elements:
                            d_110_v22_: str = compr_14_
                            if (d_110_v22_) in (_dafny.Map({d_107_v20_: d_107_v20_})):
                                coll14_[d_110_v22_] = False
                        return _dafny.Map(coll14_)
                    compr_13_: str
                    for compr_13_ in (iife18_()
                    ).keys.Elements:
                        d_111_v23_: str = compr_13_
                        if (d_111_v23_) in (iife19_()
                        ):
                            coll13_ = coll13_.union(_dafny.Set([d_111_v23_]))
                    return _dafny.Set(coll13_)
                nw6_[int(11)] = (iife17_()
                ) | (d_108_v21_)
                nw6_[int(12)] = (d_108_v21_) - (d_108_v21_)
                nw6_[int(13)] = d_108_v21_
                nw6_[int(14)] = _dafny.Set({d_107_v20_})
                nw6_[int(15)] = d_108_v21_
                nw6_[int(16)] = d_108_v21_
                nw6_[int(17)] = (d_108_v21_) | (d_108_v21_)
                nw6_[int(18)] = d_108_v21_
                d_109_v24_ = nw6_
                d_112_v25_: _dafny.Seq
                d_112_v25_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khlhpemq")))])
                d_113_v26_: _dafny.Array
                nw7_ = _dafny.Array(None, 5)
                d_113_v26_ = nw7_
                d_114_v27_: _dafny.Map
                d_114_v27_ = _dafny.Map({d_113_v26_: p0})
                d_115_v28_: _dafny.Array
                nw8_ = _dafny.Array(None, 7)
                nw8_[int(0)] = (d_112_v25_) + (_dafny.SeqWithoutIsStrInference([p3, len(d_114_v27_)]))
                nw8_[int(1)] = d_112_v25_
                nw8_[int(2)] = _dafny.SeqWithoutIsStrInference([p1, p3, 7])
                nw8_[int(3)] = (d_112_v25_) + ((d_112_v25_).set(default__.safeIndex(p1, len(d_112_v25_)), -307))
                nw8_[int(4)] = d_112_v25_
                nw8_[int(5)] = d_112_v25_
                nw8_[int(6)] = d_112_v25_
                d_115_v28_ = nw8_
                index2_ = default__.safeIndex(414, (d_115_v28_).length(0))
                (d_115_v28_)[index2_] = _dafny.SeqWithoutIsStrInference([p2 for d_116_i2_ in range(default__.abs(754))])
                index3_ = default__.safeIndex(414, (d_115_v28_).length(0))
                rhs0_ = d_109_v24_
                rhs1_ = ((d_112_v25_) + (d_112_v25_) if p0 else d_112_v25_)
                lhs0_ = d_115_v28_
                lhs1_ = default__.safeIndex(414, (d_115_v28_).length(0))
                d_109_v24_ = rhs0_
                lhs0_[lhs1_] = rhs1_
                d_117_v29_: _dafny.Map
                d_117_v29_ = _dafny.Map({d_84_cf56_: 690})
                d_118_v31_: _dafny.Seq
                d_118_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adhmxyo"))
                d_119_v32_: _dafny.Array
                nw9_ = _dafny.Array(None, 27)
                nw9_[int(0)] = d_117_v29_
                nw9_[int(1)] = d_117_v29_
                nw9_[int(2)] = d_117_v29_
                nw9_[int(3)] = d_117_v29_
                nw9_[int(4)] = d_117_v29_
                nw9_[int(5)] = d_117_v29_
                nw9_[int(6)] = d_117_v29_
                nw9_[int(7)] = d_117_v29_
                def iife20_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(-415, 552):
                        d_120_v30_: int = compr_16_
                        if ((-415) <= (d_120_v30_)) and ((d_120_v30_) < (552)):
                            coll16_[default__.safeDivisionInt(d_120_v30_, d_84_cf56_)] = p3
                    return _dafny.Map(coll16_)
                nw9_[int(8)] = iife20_()
                
                nw9_[int(9)] = d_117_v29_
                nw9_[int(10)] = d_117_v29_
                nw9_[int(11)] = d_117_v29_
                nw9_[int(12)] = d_117_v29_
                nw9_[int(13)] = d_117_v29_
                nw9_[int(14)] = _dafny.Map({p3: d_84_cf56_})
                nw9_[int(15)] = d_117_v29_
                nw9_[int(16)] = d_117_v29_
                nw9_[int(17)] = _dafny.Map({p3: d_84_cf56_})
                nw9_[int(18)] = d_117_v29_
                nw9_[int(19)] = d_117_v29_
                nw9_[int(20)] = d_117_v29_
                nw9_[int(21)] = d_117_v29_
                nw9_[int(22)] = d_117_v29_
                nw9_[int(23)] = d_117_v29_
                nw9_[int(24)] = d_117_v29_
                nw9_[int(25)] = _dafny.Map({p1: len(d_118_v31_)})
                nw9_[int(26)] = d_117_v29_
                d_119_v32_ = nw9_
                d_121_v33_: _dafny.Map
                d_121_v33_ = _dafny.Map({p3: D16_DC39(d_119_v32_)})
                d_122_v34_: _dafny.MultiSet
                d_122_v34_ = _dafny.MultiSet([d_121_v33_])
                d_122_v34_ = d_122_v34_
                d_123_v35_: _dafny.MultiSet
                d_123_v35_ = _dafny.MultiSet([d_107_v20_])
                d_124_v36_: _dafny.Array
                nw10_ = _dafny.Array(None, 10)
                nw10_[int(0)] = (0) - (d_84_cf56_)
                nw10_[int(1)] = d_84_cf56_
                nw10_[int(2)] = d_84_cf56_
                nw10_[int(3)] = d_84_cf56_
                nw10_[int(4)] = default__.safeDivisionInt((d_123_v35_).cardinality, d_84_cf56_)
                nw10_[int(5)] = p3
                nw10_[int(6)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eiavcl")))
                nw10_[int(7)] = d_84_cf56_
                nw10_[int(8)] = p2
                nw10_[int(9)] = p3
                d_124_v36_ = nw10_
                d_124_v36_ = d_124_v36_
                d_125_v37_: _dafny.MultiSet
                d_125_v37_ = _dafny.MultiSet([p0])
                d_126_v38_: C1
                nw11_ = C1()
                nw11_.ctor__(p0, d_125_v37_, d_112_v25_, d_84_cf56_, d_85_v0_)
                d_126_v38_ = nw11_
            d_127_v39_: _dafny.MultiSet
            d_127_v39_ = _dafny.MultiSet([p0])
            d_128_v40_: C2
            nw12_ = C2()
            nw12_.ctor__(d_127_v39_, (len((_dafny.Map({default__.fm0(p1, globalState): p0})).set(default__.fm0(p3, globalState), p0)) if p0 else 207))
            d_128_v40_ = nw12_
            d_128_v40_ = d_128_v40_
        elif source8_.is_DC32:
            d_129___mcc_h1_ = source8_.cf57
            d_130___mcc_h2_ = source8_.cf58
            d_131___mcc_h3_ = source8_.cf59
            d_132___mcc_h4_ = source8_.cf60
            d_133_cf60_ = d_132___mcc_h4_
            d_134_cf59_ = d_131___mcc_h3_
            d_135_cf58_ = d_130___mcc_h2_
            d_136_cf57_ = d_129___mcc_h1_
            d_137_v41_: _dafny.Seq
            d_137_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jc"))
            d_138_v42_: _dafny.Array
            def lambda2_(d_139_cf57_, d_140_p0_):
                def lambda3_(d_141_i3_):
                    return (d_139_cf57_) == ((D4_DC9(d_139_cf57_, d_140_p0_)).cf16)

                return lambda3_

            init1_ = lambda2_(d_136_cf57_, p0)
            nw13_ = _dafny.Array(None, 6)
            for i0_1_ in range(nw13_.length(0)):
                nw13_[i0_1_] = init1_(i0_1_)
            d_138_v42_ = nw13_
            rhs2_ = d_135_cf58_
            rhs3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkc"))
            rhs4_ = False
            rhs5_ = d_138_v42_
            rhs6_ = (p3 if False else d_135_cf58_)
            d_135_cf58_ = rhs2_
            d_137_v41_ = rhs3_
            d_133_cf60_ = rhs4_
            d_138_v42_ = rhs5_
            d_135_cf58_ = rhs6_
            (globalState).f15 = d_135_cf58_
            if not (d_133_cf60_) or (p0):
                d_142_v43_: _dafny.MultiSet
                d_142_v43_ = _dafny.MultiSet([p0])
                d_143_v44_: _dafny.Array
                nw14_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
                d_143_v44_ = nw14_
                d_144_v45_: _dafny.Map
                d_144_v45_ = _dafny.Map({d_133_cf60_: d_143_v44_})
                d_145_v46_: T2
                nw15_ = C6()
                nw15_.ctor__(True, p2, d_142_v43_, d_134_cf59_, p3, ((d_144_v45_)[not(not(d_133_cf60_))] if (not(not(d_133_cf60_))) in (d_144_v45_) else d_143_v44_), d_142_v43_, p2)
                d_145_v46_ = nw15_
                d_146_v47_: _dafny.Array
                nw16_ = _dafny.Array(None, 13)
                nw16_[int(0)] = d_145_v46_
                nw16_[int(1)] = d_145_v46_
                nw16_[int(2)] = d_145_v46_
                nw16_[int(3)] = d_145_v46_
                nw16_[int(4)] = d_145_v46_
                nw16_[int(5)] = d_145_v46_
                nw16_[int(6)] = d_145_v46_
                nw16_[int(7)] = d_145_v46_
                nw16_[int(8)] = d_145_v46_
                nw16_[int(9)] = d_145_v46_
                nw16_[int(10)] = d_145_v46_
                nw16_[int(11)] = d_145_v46_
                nw16_[int(12)] = d_145_v46_
                d_146_v47_ = nw16_
                d_146_v47_ = d_146_v47_
                d_147_v48_: C0
                nw17_ = C0()
                nw17_.ctor__(d_136_cf57_, d_137_v41_)
                d_147_v48_ = nw17_
                d_148_v49_: _dafny.Seq
                d_148_v49_ = _dafny.SeqWithoutIsStrInference([d_147_v48_, d_147_v48_])
                d_149_v50_: _dafny.Seq
                d_149_v50_ = _dafny.SeqWithoutIsStrInference([d_148_v49_])
                d_150_v51_: str
                d_150_v51_ = _dafny.CodePoint('w')
                d_151_v52_: _dafny.Array
                nw18_ = _dafny.Array(None, 10)
                nw18_[int(0)] = (d_145_v46_).f32
                nw18_[int(1)] = len((d_149_v50_) + (d_149_v50_))
                nw18_[int(2)] = (d_147_v48_).f36
                nw18_[int(3)] = (d_147_v48_).f36
                nw18_[int(4)] = d_136_cf57_
                nw18_[int(5)] = 913
                nw18_[int(6)] = len(default__.fm56(D4_DC8(d_150_v51_, 470, p0), p0, default__.fm57(p0, p0, globalState), globalState))
                nw18_[int(7)] = (d_145_v46_).f32
                nw18_[int(8)] = d_135_cf58_
                nw18_[int(9)] = 543
                d_151_v52_ = nw18_
                index4_ = default__.safeIndex(547, (d_151_v52_).length(0))
                (d_151_v52_)[index4_] = d_136_cf57_
                d_152_v53_: _dafny.Map
                d_152_v53_ = _dafny.Map({p3: d_136_cf57_})
                index5_ = default__.safeIndex(547, (d_151_v52_).length(0))
                rhs7_ = default__.fm8(globalState)
                rhs8_ = (D9_DC20(d_135_cf58_, not(d_133_cf60_), p0, p0, d_152_v53_)).cf34
                rhs9_ = False
                rhs10_ = not(p0)
                lhs2_ = d_151_v52_
                lhs3_ = default__.safeIndex(547, (d_151_v52_).length(0))
                lhs4_ = globalState
                d_137_v41_ = rhs7_
                lhs2_[lhs3_] = rhs8_
                lhs4_.f26 = rhs9_
                d_133_cf60_ = rhs10_
                index6_ = default__.safeIndex(239, (d_138_v42_).length(0))
                (d_138_v42_)[index6_] = p0
                index7_ = default__.safeIndex(239, (d_138_v42_).length(0))
                (d_138_v42_)[index7_] = p0
                d_153_v54_: _dafny.MultiSet
                d_153_v54_ = _dafny.MultiSet([len(d_134_cf59_), (p2) - (p1), default__.safeDivisionInt((d_151_v52_)[default__.safeIndex(547, (d_151_v52_).length(0))], (d_147_v48_).f36), -492])
                index8_ = default__.safeIndex(547, (d_151_v52_).length(0))
                rhs11_ = (d_153_v54_) | (d_153_v54_)
                rhs12_ = len((d_147_v48_).f37)
                lhs5_ = d_151_v52_
                lhs6_ = default__.safeIndex(547, (d_151_v52_).length(0))
                d_153_v54_ = rhs11_
                lhs5_[lhs6_] = rhs12_
                (globalState).f26 = (d_138_v42_)[default__.safeIndex(239, (d_138_v42_).length(0))]
            elif True:
                d_154_v55_: _dafny.Map
                d_154_v55_ = _dafny.Map({p3: (d_137_v41_)[default__.safeIndex(843, len(d_137_v41_))]})
                d_155_v56_: _dafny.Map
                d_155_v56_ = _dafny.Map({p2: d_154_v55_})
                d_156_v57_: _dafny.Seq
                d_156_v57_ = _dafny.SeqWithoutIsStrInference([((d_155_v56_)[d_135_cf58_] if (d_135_cf58_) in (d_155_v56_) else _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_157_i4_ in range(default__.abs(678))])): _dafny.CodePoint('t')}))])
                d_158_v58_: D15
                d_158_v58_ = D15_DC35(d_156_v57_)
                d_158_v58_ = (d_158_v58_ if p0 else D15_DC35(d_156_v57_))
                (globalState).f18 = default__.safeModuloInt(len(_dafny.Set({d_133_cf60_})), -822)
                d_159_v59_: D13
                d_159_v59_ = D13_DC32(14, p3, d_134_cf59_, d_133_cf60_)
                d_160_v60_: _dafny.Seq
                d_160_v60_ = _dafny.SeqWithoutIsStrInference([d_159_v59_])
                d_160_v60_ = (d_160_v60_) + (_dafny.SeqWithoutIsStrInference([d_159_v59_, d_159_v59_]))
                d_161_v61_: str
                d_161_v61_ = _dafny.CodePoint('o')
                d_162_v62_: _dafny.Array
                nw19_ = _dafny.Array(None, 12)
                nw19_[int(0)] = d_161_v61_
                nw19_[int(1)] = d_161_v61_
                nw19_[int(2)] = _dafny.CodePoint('d')
                nw19_[int(3)] = _dafny.CodePoint('u')
                nw19_[int(4)] = d_161_v61_
                nw19_[int(5)] = _dafny.CodePoint('p')
                nw19_[int(6)] = d_161_v61_
                nw19_[int(7)] = d_161_v61_
                nw19_[int(8)] = d_161_v61_
                nw19_[int(9)] = d_161_v61_
                nw19_[int(10)] = d_161_v61_
                nw19_[int(11)] = d_161_v61_
                d_162_v62_ = nw19_
                index9_ = default__.safeIndex(448, (d_162_v62_).length(0))
                (d_162_v62_)[index9_] = d_161_v61_
                index10_ = default__.safeIndex(448, (d_162_v62_).length(0))
                (d_162_v62_)[index10_] = d_161_v61_
                (globalState).f21 = (0) - (p3)
            r0 = p0
        elif source8_.is_DC30:
            d_163___mcc_h5_ = source8_.cf55
            d_164_cf55_ = d_163___mcc_h5_
            d_165_v63_: _dafny.Seq
            d_165_v63_ = _dafny.SeqWithoutIsStrInference([p3])
            if not(not((d_165_v63_) < ((_dafny.SeqWithoutIsStrInference([p3, p3, (0) - (p3)])) + (d_165_v63_)))):
                d_166_v64_: _dafny.Seq
                d_166_v64_ = _dafny.SeqWithoutIsStrInference([p0, d_164_cf55_.f35])
                d_167_v65_: _dafny.Seq
                d_167_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmttqthg"))
                d_168_v66_: _dafny.Array
                nw20_ = _dafny.Array(None, 27)
                nw20_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_169_i5_ in range(default__.abs(-212))])
                nw20_[int(1)] = d_167_v65_
                nw20_[int(2)] = d_167_v65_
                nw20_[int(3)] = d_167_v65_
                nw20_[int(4)] = d_167_v65_
                nw20_[int(5)] = d_167_v65_
                nw20_[int(6)] = d_167_v65_
                nw20_[int(7)] = d_167_v65_
                nw20_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfdjji"))
                nw20_[int(9)] = d_167_v65_
                nw20_[int(10)] = d_167_v65_
                nw20_[int(11)] = d_167_v65_
                nw20_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eet"))
                nw20_[int(13)] = d_167_v65_
                nw20_[int(14)] = d_167_v65_
                nw20_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqv"))
                nw20_[int(16)] = d_167_v65_
                nw20_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwkfera"))
                nw20_[int(18)] = d_167_v65_
                nw20_[int(19)] = d_167_v65_
                nw20_[int(20)] = d_167_v65_
                nw20_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                nw20_[int(22)] = d_167_v65_
                nw20_[int(23)] = d_167_v65_
                nw20_[int(24)] = d_167_v65_
                nw20_[int(25)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldwxsm"))
                nw20_[int(26)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fytgb"))
                d_168_v66_ = nw20_
                d_170_v67_: _dafny.MultiSet
                d_170_v67_ = _dafny.MultiSet([d_164_cf55_.f35, p0, p0, p0])
                d_171_v68_: C6
                nw21_ = C6()
                nw21_.ctor__((d_164_cf55_).fm3(_dafny.MultiSet([p0]), globalState), p2, _dafny.MultiSet(d_166_v64_), (default__.fm16(511, globalState)).set(default__.safeIndex(p1, len(default__.fm16(511, globalState))), p2), (p1) - (p2), d_168_v66_, (d_170_v67_).set(True, default__.abs((d_170_v67_).cardinality)), len(d_167_v65_))
                d_171_v68_ = nw21_
                d_172_v69_: C4
                nw22_ = C4()
                nw22_.ctor__(len((d_165_v63_) + (_dafny.SeqWithoutIsStrInference([(0) - (len(d_167_v65_)) for d_173_i6_ in range(default__.abs(-21))]))), d_168_v66_)
                d_172_v69_ = nw22_
                d_172_v69_ = d_172_v69_
                d_174_v70_: _dafny.Seq
                d_174_v70_ = d_167_v65_
                d_175_v71_: _dafny.Set
                d_175_v71_ = _dafny.Set({(d_164_cf55_).fm11(False, d_171_v68_.f34, d_174_v70_, p3, globalState)})
                d_176_v72_: _dafny.Map
                d_176_v72_ = _dafny.Map({d_164_cf55_.f35: (d_175_v71_).issubset(d_175_v71_)})
                d_176_v72_ = (d_176_v72_).set((d_171_v68_).f33, not(d_164_cf55_.f35))
                d_177_v73_: str
                d_177_v73_ = _dafny.CodePoint('n')
                d_177_v73_ = d_177_v73_
                d_178_v74_: _dafny.Map
                d_178_v74_ = _dafny.Map({len(d_167_v65_): (926) * (p1)})
                d_178_v74_ = (d_178_v74_).set(len((d_178_v74_).set((0) - (p3), 93)), d_171_v68_.f34)
            elif True:
                d_179_v75_: _dafny.MultiSet
                d_179_v75_ = _dafny.MultiSet([d_164_cf55_.f35])
                d_180_v76_: _dafny.Array
                def lambda4_(d_181_i7_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mte"))

                init2_ = lambda4_
                nw23_ = _dafny.Array(None, 26)
                for i0_2_ in range(nw23_.length(0)):
                    nw23_[i0_2_] = init2_(i0_2_)
                d_180_v76_ = nw23_
                d_182_v77_: T2
                nw24_ = C6()
                nw24_.ctor__(d_164_cf55_.f35, p3, d_179_v75_, _dafny.SeqWithoutIsStrInference([p1]), len(d_165_v63_), d_180_v76_, d_179_v75_, p3)
                d_182_v77_ = nw24_
                d_183_v78_: _dafny.Seq
                d_183_v78_ = _dafny.SeqWithoutIsStrInference([d_164_cf55_.f35])
                d_184_v79_: _dafny.Map
                d_184_v79_ = _dafny.Map({d_182_v77_: len(d_183_v78_)})
                d_185_v80_: _dafny.Map
                d_185_v80_ = _dafny.Map({d_164_cf55_.f35: d_184_v79_})
                d_186_v81_: _dafny.MultiSet
                d_186_v81_ = _dafny.MultiSet([d_184_v79_, ((d_185_v80_)[p0] if (p0) in (d_185_v80_) else d_184_v79_), _dafny.Map({d_182_v77_: p3})])
                d_187_v82_: T1
                nw25_ = C1()
                nw25_.ctor__(d_164_cf55_.f35, _dafny.MultiSet([p0]), d_165_v63_, (d_186_v81_).cardinality, d_180_v76_)
                d_187_v82_ = nw25_
                d_188_v83_: str
                d_188_v83_ = _dafny.CodePoint('u')
                rhs13_ = True
                rhs14_ = default__.fm0((d_187_v82_).fm4(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_188_v83_])}), p0, d_183_v78_, globalState), globalState)
                rhs15_ = d_187_v82_
                lhs7_ = d_164_cf55_
                lhs8_ = globalState
                lhs7_.f35 = rhs13_
                lhs8_.f26 = rhs14_
                d_187_v82_ = rhs15_
                d_189_v84_: _dafny.Array
                nw26_ = _dafny.Array(False, 8)
                d_189_v84_ = nw26_
                index11_ = default__.safeIndex(562, (d_189_v84_).length(0))
                def iife21_():
                    coll17_ = _dafny.Map()
                    compr_17_: int
                    for compr_17_ in _dafny.IntegerRange(446, 851):
                        d_190_v85_: int = compr_17_
                        if ((446) <= (d_190_v85_)) and ((d_190_v85_) < (851)):
                            coll17_[(d_190_v85_) - (531)] = p0
                    return _dafny.Map(coll17_)
                (d_189_v84_)[index11_] = (p2) >= (len(iife21_()
                ))
                d_191_v86_: _dafny.Map
                d_191_v86_ = _dafny.Map({d_164_cf55_.f35: -355})
                d_192_v87_: _dafny.Seq
                d_192_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ostrgckj"))
                d_193_v88_: _dafny.Seq
                d_193_v88_ = d_192_v87_
                index12_ = default__.safeIndex(562, (d_189_v84_).length(0))
                (d_189_v84_)[index12_] = not((d_164_cf55_).fm11(True, (len(d_191_v86_)) - ((d_187_v82_).f27), d_193_v88_, (d_182_v77_).f32, globalState))
                r0 = p0
                d_194_v89_: _dafny.Array
                def lambda5_(d_195_v77_, d_196_v82_, d_197_v75_):
                    def lambda6_(d_198_i8_):
                        return _dafny.MultiSet([(d_195_v77_).f31, (d_196_v82_).f29, d_197_v75_, (d_195_v77_).f31])

                    return lambda6_

                init3_ = lambda5_(d_182_v77_, d_187_v82_, d_179_v75_)
                nw27_ = _dafny.Array(None, 20)
                for i0_3_ in range(nw27_.length(0)):
                    nw27_[i0_3_] = init3_(i0_3_)
                d_194_v89_ = nw27_
                d_199_v90_: _dafny.MultiSet
                d_199_v90_ = _dafny.MultiSet([d_179_v75_])
                index13_ = default__.safeIndex(765, (d_194_v89_).length(0))
                (d_194_v89_)[index13_] = d_199_v90_
                d_200_v91_: _dafny.Map
                d_200_v91_ = _dafny.Map({len(d_192_v87_): d_199_v90_})
                index14_ = default__.safeIndex(765, (d_194_v89_).length(0))
                (d_194_v89_)[index14_] = (((d_200_v91_)[p3] if (p3) in (d_200_v91_) else d_199_v90_)).intersection(d_199_v90_)
                (globalState).f3 = (len(d_191_v86_)) - (904)
            d_201_v92_: _dafny.Array
            nw28_ = _dafny.Array(None, 4)
            d_201_v92_ = nw28_
            d_202_v93_: _dafny.MultiSet
            d_202_v93_ = _dafny.MultiSet([True, not(d_164_cf55_.f35)])
            d_203_v94_: _dafny.Seq
            d_203_v94_ = _dafny.SeqWithoutIsStrInference([d_202_v93_, d_202_v93_])
            d_204_v95_: C2
            nw29_ = C2()
            nw29_.ctor__((d_203_v94_)[default__.safeIndex(p3, len(d_203_v94_))], p3)
            d_204_v95_ = nw29_
            index15_ = default__.safeIndex(466, (d_201_v92_).length(0))
            (d_201_v92_)[index15_] = d_204_v95_
            d_205_v96_: _dafny.Seq
            d_205_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tytaticn"))
            d_206_v97_: _dafny.Map
            d_206_v97_ = _dafny.Map({d_164_cf55_.f35: p1})
            d_207_v98_: str
            d_207_v98_ = _dafny.CodePoint('l')
            d_208_v99_: _dafny.Set
            d_208_v99_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_207_v98_])})
            d_209_v100_: _dafny.Seq
            d_209_v100_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p1, globalState)])
            d_210_v101_: _dafny.Array
            nw30_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
            d_210_v101_ = nw30_
            d_211_v102_: C3
            nw31_ = C3()
            nw31_.ctor__(p1, d_210_v101_)
            d_211_v102_ = nw31_
            d_212_v103_: _dafny.Map
            d_212_v103_ = _dafny.Map({((d_206_v97_)[False] if (False) in (d_206_v97_) else (d_202_v93_).cardinality): d_211_v102_})
            d_213_v104_: _dafny.MultiSet
            d_213_v104_ = _dafny.MultiSet([(len(d_205_v96_)) * (len(d_206_v97_)), p2, p1, (d_164_cf55_).fm4(d_208_v99_, False, d_209_v100_, globalState), len(d_212_v103_)])
            d_214_v105_: _dafny.Set
            d_214_v105_ = _dafny.Set({p0, d_164_cf55_.f35})
            d_215_v106_: _dafny.Set
            d_215_v106_ = _dafny.Set({p1})
            d_216_v107_: _dafny.Map
            d_216_v107_ = _dafny.Map({d_164_cf55_.f35: p0})
            d_217_v108_: D10
            d_217_v108_ = D10_DC22(len(d_205_v96_), d_202_v93_)
            d_218_v109_: _dafny.MultiSet
            d_218_v109_ = _dafny.MultiSet([d_214_v105_, d_214_v105_])
            d_219_v110_: _dafny.Map
            d_219_v110_ = _dafny.Map({p3: p0})
            d_220_v111_: _dafny.Array
            nw32_ = _dafny.Array(None, 21)
            nw32_[int(0)] = default__.safeDivisionInt(p1, p1)
            nw32_[int(1)] = p1
            nw32_[int(2)] = p3
            nw32_[int(3)] = p1
            nw32_[int(4)] = p2
            nw32_[int(5)] = (default__.fm1(globalState)) * (len(d_214_v105_))
            nw32_[int(6)] = p3
            nw32_[int(7)] = default__.safeDivisionInt(p1, len(d_215_v106_))
            nw32_[int(8)] = p2
            nw32_[int(9)] = len(d_216_v107_)
            nw32_[int(10)] = p1
            nw32_[int(11)] = 307
            nw32_[int(12)] = p1
            nw32_[int(13)] = p2
            nw32_[int(14)] = (d_217_v108_).cf40
            nw32_[int(15)] = p1
            nw32_[int(16)] = ((d_218_v109_).set(d_214_v105_, default__.abs(((d_202_v93_)[p0] if (p0) in (d_202_v93_) else p1)))).cardinality
            nw32_[int(17)] = (0) - (default__.fm1(globalState))
            nw32_[int(18)] = (0) - ((p2) + (p1))
            nw32_[int(19)] = p2
            nw32_[int(20)] = (p2) - (len(d_219_v110_))
            d_220_v111_ = nw32_
            index16_ = default__.safeIndex(506, (d_220_v111_).length(0))
            (d_220_v111_)[index16_] = default__.safeModuloInt(len(d_206_v97_), p2)
            d_221_v112_: _dafny.Array
            nw33_ = _dafny.Array(None, 18)
            d_221_v112_ = nw33_
            d_222_v113_: _dafny.Seq
            d_222_v113_ = _dafny.SeqWithoutIsStrInference([d_221_v112_])
            d_223_v114_: _dafny.Map
            d_223_v114_ = _dafny.Map({p0: d_205_v96_})
            d_224_v115_: D22
            d_224_v115_ = D22_DC53(d_222_v113_)
            index17_ = default__.safeIndex(466, (d_201_v92_).length(0))
            index18_ = default__.safeIndex(506, (d_220_v111_).length(0))
            rhs16_ = d_204_v95_
            rhs17_ = ((d_213_v104_).intersection(d_213_v104_)) - (d_213_v104_)
            rhs18_ = len((d_165_v63_ if not (d_164_cf55_.f35) or (p0) else (d_165_v63_) + (d_165_v63_)))
            rhs19_ = len((_dafny.Map({p0: d_205_v96_})) | ((d_223_v114_) | (d_223_v114_)))
            rhs20_ = (d_224_v115_).cf94
            lhs9_ = d_201_v92_
            lhs10_ = default__.safeIndex(466, (d_201_v92_).length(0))
            lhs11_ = globalState
            lhs12_ = d_220_v111_
            lhs13_ = default__.safeIndex(506, (d_220_v111_).length(0))
            lhs9_[lhs10_] = rhs16_
            d_213_v104_ = rhs17_
            lhs11_.f3 = rhs18_
            lhs12_[lhs13_] = rhs19_
            d_222_v113_ = rhs20_
            d_225_v116_: C6
            nw34_ = C6()
            nw34_.ctor__(not(d_164_cf55_.f35), len(d_205_v96_), d_202_v93_, default__.fm16(p1, globalState), (d_220_v111_)[default__.safeIndex(506, (d_220_v111_).length(0))], d_210_v101_, d_202_v93_, p3)
            d_225_v116_ = nw34_
            d_226_v117_: _dafny.Map
            d_226_v117_ = _dafny.Map({d_225_v116_: (d_220_v111_)[default__.safeIndex(506, (d_220_v111_).length(0))]})
            d_227_v118_: _dafny.Map
            d_227_v118_ = _dafny.Map({d_220_v111_: d_226_v117_})
            d_227_v118_ = (d_227_v118_).set(d_220_v111_, (d_226_v117_).set(d_225_v116_, (d_220_v111_)[default__.safeIndex(506, (d_220_v111_).length(0))]))
            d_228_v119_: D4
            d_228_v119_ = D4_DC9(-28, p0)
            d_229_v120_: _dafny.Array
            nw35_ = _dafny.Array(_dafny.Map({}), 21)
            d_229_v120_ = nw35_
            d_230_v121_: _dafny.Array
            d_230_v121_ = d_229_v120_
            source9_ = (d_230_v121_ if not(((d_228_v119_).cf16) >= ((d_220_v111_)[default__.safeIndex(506, (d_220_v111_).length(0))])) else d_230_v121_)
            d_231___mcc_h7_ = source9_
            d_232_cf93_ = d_231___mcc_h7_
            d_233_v122_: _dafny.Array
            def lambda7_(d_234_p0_, d_235_v98_, d_236_v116_):
                def lambda8_(d_237_i9_):
                    return _dafny.Map({d_234_p0_: D4_DC8(d_235_v98_, (0) - (d_236_v116_.f34), d_234_p0_)})

                return lambda8_

            init4_ = lambda7_(p0, d_207_v98_, d_225_v116_)
            nw36_ = _dafny.Array(None, 1)
            for i0_4_ in range(nw36_.length(0)):
                nw36_[i0_4_] = init4_(i0_4_)
            d_233_v122_ = nw36_
            d_238_v123_: D4
            d_238_v123_ = D4_DC8(d_207_v98_, len(d_216_v107_), (d_225_v116_).f33)
            index19_ = default__.safeIndex(124, (d_233_v122_).length(0))
            (d_233_v122_)[index19_] = _dafny.Map({True: d_238_v123_})
            d_239_v124_: _dafny.Map
            d_239_v124_ = _dafny.Map({d_164_cf55_.f35: d_238_v123_})
            index20_ = default__.safeIndex(124, (d_233_v122_).length(0))
            (d_233_v122_)[index20_] = ((d_239_v124_).set(d_164_cf55_.f35, d_238_v123_)).set(not (p0) or (False), d_238_v123_)
            d_240_v125_: C0
            nw37_ = C0()
            nw37_.ctor__(p1, ((d_223_v114_)[not(d_164_cf55_.f35)] if (not(d_164_cf55_.f35)) in (d_223_v114_) else _dafny.SeqWithoutIsStrInference([d_207_v98_ for d_241_i10_ in range(default__.abs(751))])))
            d_240_v125_ = nw37_
            d_242_v126_: _dafny.Seq
            d_242_v126_ = _dafny.SeqWithoutIsStrInference([(d_201_v92_)[default__.safeIndex(466, (d_201_v92_).length(0))], d_204_v95_])
            d_243_v127_: _dafny.Map
            d_243_v127_ = _dafny.Map({(d_220_v111_)[default__.safeIndex(506, (d_220_v111_).length(0))]: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqkukel")))})
            d_244_v128_: C5
            nw38_ = C5()
            nw38_.ctor__(_dafny.Map({d_243_v127_: p1}), d_202_v93_, d_165_v63_, len(d_205_v96_), d_210_v101_)
            d_244_v128_ = nw38_
            d_245_v129_: _dafny.MultiSet
            d_245_v129_ = _dafny.MultiSet([d_244_v128_])
            nw39_ = C1()
            nw39_.ctor__((d_242_v126_) == (d_242_v126_), d_202_v93_, (_dafny.SeqWithoutIsStrInference([-554, len(d_223_v114_), (d_245_v129_).cardinality])) + (_dafny.SeqWithoutIsStrInference([p2])), len(d_165_v63_), d_210_v101_)
            d_164_cf55_ = nw39_
            d_246_v130_: _dafny.Array
            def lambda9_(d_247_v98_, d_248_v63_):
                def lambda10_(d_249_i11_):
                    return _dafny.Map({d_247_v98_: d_248_v63_})

                return lambda10_

            init5_ = lambda9_(d_207_v98_, d_165_v63_)
            nw40_ = _dafny.Array(None, 11)
            for i0_5_ in range(nw40_.length(0)):
                nw40_[i0_5_] = init5_(i0_5_)
            d_246_v130_ = nw40_
            d_250_v131_: D2
            d_250_v131_ = D2_DC4(p2, (d_220_v111_)[default__.safeIndex(506, (d_220_v111_).length(0))], (d_240_v125_).f36, d_207_v98_, d_205_v96_)
            d_251_v132_: _dafny.Map
            d_251_v132_ = _dafny.Map({(d_250_v131_).cf8: default__.fm43((d_225_v116_).f33, p0, not(p0), globalState)})
            index21_ = default__.safeIndex(383, (d_246_v130_).length(0))
            (d_246_v130_)[index21_] = d_251_v132_
            d_252_v133_: _dafny.Map
            d_252_v133_ = _dafny.Map({d_164_cf55_: d_251_v132_})
            index22_ = default__.safeIndex(383, (d_246_v130_).length(0))
            (d_246_v130_)[index22_] = ((d_252_v133_)[d_164_cf55_] if (d_164_cf55_) in (d_252_v133_) else d_251_v132_)
        elif True:
            d_253___mcc_h6_ = source8_.cf61
            d_254_cf61_ = d_253___mcc_h6_
            (globalState).f26 = p0
            d_255_v134_: _dafny.Array
            nw41_ = _dafny.Array(None, 11)
            d_255_v134_ = nw41_
            d_256_v135_: _dafny.Set
            d_256_v135_ = _dafny.Set({d_255_v134_, d_255_v134_, d_255_v134_})
            if not (p0) or ((d_256_v135_).ispropersubset(d_256_v135_)):
                (globalState).f26 = not(False)
                d_257_v136_: C0
                nw42_ = C0()
                nw42_.ctor__(p1, default__.fm8(globalState))
                d_257_v136_ = nw42_
                d_258_v137_: D6
                d_258_v137_ = D6_DC11(d_257_v136_)
                d_259_v138_: D19
                d_259_v138_ = D19_DC49(p0, p0, 807)
                d_260_v139_: _dafny.Map
                d_260_v139_ = _dafny.Map({d_258_v137_: d_259_v138_})
                d_260_v139_ = (d_260_v139_).set(d_258_v137_, d_259_v138_)
                d_261_v140_: _dafny.Array
                nw43_ = _dafny.Array(int(0), 7)
                d_261_v140_ = nw43_
                index23_ = default__.safeIndex(702, (d_261_v140_).length(0))
                (d_261_v140_)[index23_] = len(((d_257_v136_).f37) + ((d_257_v136_).f37))
                index24_ = default__.safeIndex(702, (d_261_v140_).length(0))
                (d_261_v140_)[index24_] = p1
                d_262_v141_: _dafny.Map
                d_262_v141_ = _dafny.Map({((d_257_v136_).f36) * (p3): p0})
                d_262_v141_ = (d_262_v141_).set(p2, p0)
                d_263_v142_: _dafny.Seq
                d_263_v142_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsshj"))
                d_264_v143_: _dafny.Set
                d_264_v143_ = _dafny.Set({(d_257_v136_).f36, p2})
                d_265_v144_: _dafny.Map
                d_265_v144_ = _dafny.Map({p2: (d_264_v143_).intersection(d_264_v143_)})
                d_266_v145_: str
                d_266_v145_ = _dafny.CodePoint('e')
                rhs21_ = (((d_257_v136_).f37).set(default__.safeIndex(p2, len((d_257_v136_).f37)), d_266_v145_)) + ((d_257_v136_).f37)
                rhs22_ = d_265_v144_
                d_263_v142_ = rhs21_
                d_265_v144_ = rhs22_
            elif True:
                d_267_v146_: _dafny.Seq
                d_267_v146_ = _dafny.SeqWithoutIsStrInference([(p1) > (p2), p0, p0, p0, p0])
                d_268_v147_: D4
                d_268_v147_ = D4_DC9(534, p0)
                rhs23_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
                rhs24_ = ((d_268_v147_ if p0 else default__.fm49(p1, globalState))).cf17
                rhs25_ = p3
                lhs14_ = globalState
                lhs15_ = globalState
                d_267_v146_ = rhs23_
                lhs14_.f26 = rhs24_
                lhs15_.f3 = rhs25_
                d_269_v148_: D10
                d_269_v148_ = D10_DC22(p3, _dafny.MultiSet(d_267_v146_))
                d_270_v149_: _dafny.Seq
                d_270_v149_ = _dafny.SeqWithoutIsStrInference([p3])
                d_271_v150_: _dafny.Array
                nw44_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_271_v150_ = nw44_
                d_272_v151_: C1
                nw45_ = C1()
                nw45_.ctor__(True, (d_269_v148_).cf41, d_270_v149_, p2, d_271_v150_)
                d_272_v151_ = nw45_
                d_273_v152_: _dafny.Map
                d_273_v152_ = _dafny.Map({p1: d_272_v151_})
                d_274_v153_: _dafny.Seq
                d_274_v153_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
                d_273_v152_ = (d_273_v152_).set(len(d_274_v153_), d_272_v151_)
                d_275_v154_: _dafny.MultiSet
                d_275_v154_ = _dafny.MultiSet([d_272_v151_.f35, p0, d_272_v151_.f35, True])
                d_276_v155_: _dafny.Map
                d_276_v155_ = _dafny.Map({d_275_v154_: p2})
                d_277_v156_: _dafny.Map
                d_277_v156_ = _dafny.Map({d_276_v155_: not(p0)})
                d_277_v156_ = (d_277_v156_).set((_dafny.Map({d_275_v154_: 231})) | (d_276_v155_), p0)
                d_278_v157_: C3
                nw46_ = C3()
                nw46_.ctor__(-596, d_271_v150_)
                d_278_v157_ = nw46_
                d_279_v158_: _dafny.Array
                def lambda11_(d_280_p0_, d_281_p3_):
                    def lambda12_(d_282_i12_):
                        return _dafny.Map({d_280_p0_: d_281_p3_})

                    return lambda12_

                init6_ = lambda11_(p0, p3)
                nw47_ = _dafny.Array(None, 25)
                for i0_6_ in range(nw47_.length(0)):
                    nw47_[i0_6_] = init6_(i0_6_)
                d_279_v158_ = nw47_
                rhs26_ = d_278_v157_
                rhs27_ = d_279_v158_
                d_278_v157_ = rhs26_
                d_279_v158_ = rhs27_
                d_283_v159_: _dafny.Set
                d_283_v159_ = _dafny.Set({len(d_267_v146_), p3, len(_dafny.Set({(d_275_v154_).set(p0, default__.abs(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_284_i13_ in range(default__.abs(163))]))))}))})
                d_285_v160_: str
                d_285_v160_ = _dafny.CodePoint('r')
                d_286_v161_: _dafny.Map
                d_286_v161_ = _dafny.Map({d_285_v160_: d_283_v159_})
                d_283_v159_ = ((d_286_v161_)[d_285_v160_] if (d_285_v160_) in (d_286_v161_) else d_283_v159_)
            d_287_v162_: _dafny.MultiSet
            d_287_v162_ = _dafny.MultiSet([p0, p0, p0, False])
            d_288_v163_: _dafny.Map
            d_288_v163_ = _dafny.Map({d_287_v162_: p0})
            d_289_v164_: _dafny.Seq
            d_289_v164_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_290_v165_: _dafny.Set
            d_290_v165_ = _dafny.Set({p3})
            d_291_v166_: _dafny.Map
            d_291_v166_ = _dafny.Map({p0: p3})
            d_292_v167_: _dafny.Array
            nw48_ = _dafny.Array(None, 29)
            nw48_[int(0)] = p1
            nw48_[int(1)] = p2
            nw48_[int(2)] = p3
            nw48_[int(3)] = p2
            nw48_[int(4)] = p2
            nw48_[int(5)] = p3
            nw48_[int(6)] = p3
            nw48_[int(7)] = len(d_288_v163_)
            nw48_[int(8)] = p1
            nw48_[int(9)] = len(d_289_v164_)
            nw48_[int(10)] = p1
            nw48_[int(11)] = 37
            nw48_[int(12)] = p1
            nw48_[int(13)] = p3
            nw48_[int(14)] = p3
            nw48_[int(15)] = p1
            nw48_[int(16)] = p1
            nw48_[int(17)] = (0) - (-917)
            nw48_[int(18)] = len(d_290_v165_)
            nw48_[int(19)] = p2
            nw48_[int(20)] = p2
            nw48_[int(21)] = len(d_290_v165_)
            nw48_[int(22)] = -608
            nw48_[int(23)] = 877
            nw48_[int(24)] = p3
            nw48_[int(25)] = len(d_291_v166_)
            nw48_[int(26)] = p3
            nw48_[int(27)] = p3
            nw48_[int(28)] = p2
            d_292_v167_ = nw48_
            d_293_v168_: _dafny.Set
            d_293_v168_ = _dafny.Set({p0, p0})
            d_294_v169_: _dafny.Map
            d_294_v169_ = _dafny.Map({d_293_v168_: len(_dafny.Set({True}))})
            d_295_v170_: _dafny.Map
            d_295_v170_ = _dafny.Map({d_292_v167_: len(d_294_v169_)})
            d_296_v171_: D23
            d_296_v171_ = D23_DC56(d_295_v170_)
            (globalState).f18 = len((((d_296_v171_).cf104) | (d_295_v170_)) | (_dafny.Map({d_292_v167_: p3})))
            d_297_v172_: _dafny.Map
            d_297_v172_ = _dafny.Map({len(d_289_v164_): p1})
            d_298_v173_: _dafny.Map
            d_298_v173_ = _dafny.Map({d_297_v172_: p3})
            d_299_v174_: _dafny.Seq
            d_299_v174_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            d_300_v175_: _dafny.Seq
            d_300_v175_ = _dafny.SeqWithoutIsStrInference([590, len(d_299_v174_), p1])
            d_301_v176_: _dafny.Array
            nw49_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
            d_301_v176_ = nw49_
            d_302_v177_: _dafny.Seq
            d_302_v177_ = _dafny.SeqWithoutIsStrInference([d_301_v176_])
            d_303_v178_: C3
            nw50_ = C3()
            nw50_.ctor__(p1, d_301_v176_)
            d_303_v178_ = nw50_
            d_304_v179_: _dafny.Seq
            d_304_v179_ = _dafny.SeqWithoutIsStrInference([d_303_v178_])
            d_305_v180_: C5
            nw51_ = C5()
            nw51_.ctor__((d_298_v173_) | (d_298_v173_), (_dafny.MultiSet([p0, p0])) | (d_287_v162_), d_300_v175_, p2, (d_302_v177_)[default__.safeIndex(len(d_304_v179_), len(d_302_v177_))])
            d_305_v180_ = nw51_
        d_306_v181_: _dafny.Seq
        d_306_v181_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ym"))
        d_307_v182_: str
        d_307_v182_ = _dafny.CodePoint('l')
        d_308_v183_: _dafny.MultiSet
        d_308_v183_ = _dafny.MultiSet([d_306_v181_, ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))), d_307_v182_)).set(default__.safeIndex(p2, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))), d_307_v182_))), d_307_v182_)])
        d_309_v184_: _dafny.Seq
        d_309_v184_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuf"))])
        d_310_i14_: int
        d_310_i14_ = 0
        with _dafny.label("0"):
            while (_dafny.MultiSet(d_309_v184_)).ispropersubset(d_308_v183_):
                with _dafny.c_label("0"):
                    if (d_310_i14_) >= (100):
                        raise _dafny.Break("0")
                    d_310_i14_ = (d_310_i14_) + (1)
                    d_311_v186_: _dafny.Set
                    d_311_v186_ = _dafny.Set({p3, p1})
                    d_312_v187_: _dafny.Seq
                    def iife22_():
                        coll18_ = _dafny.Set()
                        compr_18_: int
                        for compr_18_ in _dafny.IntegerRange(-577, 565):
                            d_313_v185_: int = compr_18_
                            if ((-577) <= (d_313_v185_)) and ((d_313_v185_) < (565)):
                                coll18_ = coll18_.union(_dafny.Set([(d_313_v185_) + (p3)]))
                        return _dafny.Set(coll18_)
                    d_312_v187_ = _dafny.SeqWithoutIsStrInference([iife22_()
                    , d_311_v186_])
                    d_314_v188_: D11
                    d_314_v188_ = D11_DC25(D11_DC23(_dafny.MultiSet(((d_312_v187_).set(default__.safeIndex(p1, len(d_312_v187_)), d_311_v186_)).set(default__.safeIndex(p3, len((d_312_v187_).set(default__.safeIndex(p1, len(d_312_v187_)), d_311_v186_))), d_311_v186_))))
                    d_315_v189_: D4
                    d_315_v189_ = D4_DC8(d_307_v182_, -919, p0)
                    d_316_v190_: D11
                    d_316_v190_ = D11_DC24(731, d_315_v189_)
                    pat_let_tv1_ = d_316_v190_
                    pat_let_tv2_ = d_316_v190_
                    d_317_v191_: _dafny.Array
                    nw52_ = _dafny.Array(None, 18)
                    nw52_[int(0)] = d_314_v188_
                    nw52_[int(1)] = d_314_v188_
                    nw52_[int(2)] = D11_DC25(d_316_v190_)
                    nw52_[int(3)] = default__.fm39(globalState)
                    nw52_[int(4)] = d_314_v188_
                    nw52_[int(5)] = d_314_v188_
                    nw52_[int(6)] = d_314_v188_
                    nw52_[int(7)] = D11_DC25(d_316_v190_)
                    nw52_[int(8)] = d_314_v188_
                    nw52_[int(9)] = d_314_v188_
                    def iife23_(_pat_let2_0):
                        def iife24_(d_318_dt__update__tmp_h1_):
                            def iife25_(_pat_let3_0):
                                def iife26_(d_319_dt__update_hcf45_h0_):
                                    return D11_DC25(d_319_dt__update_hcf45_h0_)
                                return iife26_(_pat_let3_0)
                            return iife25_(pat_let_tv1_)
                        return iife24_(_pat_let2_0)
                    nw52_[int(10)] = iife23_(d_314_v188_)
                    nw52_[int(11)] = d_314_v188_
                    nw52_[int(12)] = d_314_v188_
                    def iife27_(_pat_let4_0):
                        def iife28_(d_320_dt__update__tmp_h2_):
                            def iife29_(_pat_let5_0):
                                def iife30_(d_321_dt__update_hcf45_h1_):
                                    return D11_DC25(d_321_dt__update_hcf45_h1_)
                                return iife30_(_pat_let5_0)
                            return iife29_(pat_let_tv2_)
                        return iife28_(_pat_let4_0)
                    nw52_[int(13)] = iife27_(d_314_v188_)
                    nw52_[int(14)] = d_314_v188_
                    nw52_[int(15)] = d_314_v188_
                    nw52_[int(16)] = d_314_v188_
                    nw52_[int(17)] = default__.fm39(globalState)
                    d_317_v191_ = nw52_
                    index25_ = default__.safeIndex(730, (d_317_v191_).length(0))
                    (d_317_v191_)[index25_] = default__.fm39(globalState)
                    pat_let_tv3_ = d_316_v190_
                    index26_ = default__.safeIndex(730, (d_317_v191_).length(0))
                    def iife31_(_pat_let6_0):
                        def iife32_(d_322_dt__update__tmp_h3_):
                            def iife33_(_pat_let7_0):
                                def iife34_(d_323_dt__update_hcf45_h2_):
                                    return D11_DC25(d_323_dt__update_hcf45_h2_)
                                return iife34_(_pat_let7_0)
                            return iife33_(pat_let_tv3_)
                        return iife32_(_pat_let6_0)
                    (d_317_v191_)[index26_] = iife31_(D11_DC25(d_316_v190_))
                    r0 = not (not(p0)) or (((0) - (p2)) < (p1))
                    d_324_v192_: _dafny.MultiSet
                    d_324_v192_ = _dafny.MultiSet([not(False), p0])
                    d_325_v193_: bool
                    d_325_v193_ = p0
                    d_326_v194_: _dafny.Array
                    def lambda13_(d_327_v181_):
                        def lambda14_(d_328_i15_):
                            return d_327_v181_

                        return lambda14_

                    init7_ = lambda13_(d_306_v181_)
                    nw53_ = _dafny.Array(None, 12)
                    for i0_7_ in range(nw53_.length(0)):
                        nw53_[i0_7_] = init7_(i0_7_)
                    d_326_v194_ = nw53_
                    d_329_v195_: C1
                    nw54_ = C1()
                    nw54_.ctor__((p3) <= (p1), d_324_v192_, default__.fm43(p0, d_325_v193_, p0, globalState), default__.safeModuloInt(p2, p3), d_326_v194_)
                    d_329_v195_ = nw54_
                    (globalState).f18 = (0) - (((d_324_v192_).cardinality) - ((450) - (p1)))
                    pass
            pass
        d_330_v196_: _dafny.Array
        nw55_ = _dafny.Array(_dafny.Map({}), 12)
        d_330_v196_ = nw55_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_330_v196_).length(0)):
            d_331_i16_: int = guard_loop_0_
            if (True) and (((0) <= (d_331_i16_)) and ((d_331_i16_) < ((d_330_v196_).length(0)))):
                (d_330_v196_)[(d_331_i16_)] = ((_dafny.Map({p0: p3})) | (_dafny.Map({p0: p1}))) | (_dafny.Map({p0: p3}))
        d_332_i17_: int
        d_332_i17_ = 0
        with _dafny.label("1"):
            while p0:
                with _dafny.c_label("1"):
                    if (d_332_i17_) >= (100):
                        raise _dafny.Break("1")
                    d_332_i17_ = (d_332_i17_) + (1)
                    (globalState).f24 = p2
                    r0 = ((len(d_306_v181_)) + (p1)) <= (((0) - (p3)) * (len(d_306_v181_)))
                    d_333_v197_: _dafny.Seq
                    d_333_v197_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0, p0])
                    d_334_v198_: _dafny.Map
                    d_334_v198_ = _dafny.Map({len(d_333_v197_): p0})
                    r0 = (p0) or (not(((d_334_v198_)[p1] if (p1) in (d_334_v198_) else p0)))
                    d_335_v200_: D9
                    def iife35_():
                        coll19_ = _dafny.Map()
                        compr_19_: int
                        for compr_19_ in (d_334_v198_).keys.Elements:
                            d_336_v199_: int = compr_19_
                            if (d_336_v199_) in (d_334_v198_):
                                coll19_[(d_336_v199_) * (-886)] = p3
                        return _dafny.Map(coll19_)
                    d_335_v200_ = D9_DC20(89, p0, p0, p0, iife35_()
)
                    d_337_v201_: _dafny.Map
                    d_337_v201_ = _dafny.Map({(d_335_v200_).cf34: p3})
                    d_338_v202_: _dafny.Map
                    d_338_v202_ = _dafny.Map({d_337_v201_: len(d_306_v181_)})
                    d_339_v203_: _dafny.MultiSet
                    d_339_v203_ = _dafny.MultiSet([False])
                    d_340_v204_: _dafny.Seq
                    d_340_v204_ = _dafny.SeqWithoutIsStrInference([p1])
                    d_341_v205_: _dafny.Map
                    d_341_v205_ = _dafny.Map({p0: False})
                    d_342_v206_: _dafny.Array
                    def lambda15_(d_343_v181_):
                        def lambda16_(d_344_i18_):
                            return d_343_v181_

                        return lambda16_

                    init8_ = lambda15_(d_306_v181_)
                    nw56_ = _dafny.Array(None, 29)
                    for i0_8_ in range(nw56_.length(0)):
                        nw56_[i0_8_] = init8_(i0_8_)
                    d_342_v206_ = nw56_
                    d_345_v207_: _dafny.Map
                    d_345_v207_ = _dafny.Map({((d_341_v205_)[p0] if (p0) in (d_341_v205_) else p0): d_342_v206_})
                    d_346_v208_: C5
                    nw57_ = C5()
                    nw57_.ctor__(d_338_v202_, d_339_v203_, d_340_v204_, p3, ((d_345_v207_)[p0] if (p0) in (d_345_v207_) else d_342_v206_))
                    d_346_v208_ = nw57_
                    pass
            pass
        d_347_v209_: _dafny.Set
        d_347_v209_ = _dafny.Set({d_307_v182_})
        d_348_v210_: _dafny.MultiSet
        d_348_v210_ = _dafny.MultiSet([365])
        d_349_v211_: _dafny.MultiSet
        d_349_v211_ = _dafny.MultiSet([len(d_347_v209_), p1, p3, (d_348_v210_).cardinality])
        d_348_v210_ = d_348_v210_
        d_350_v212_: _dafny.Array
        nw58_ = _dafny.Array(False, 5)
        d_350_v212_ = nw58_
        d_351_v213_: _dafny.MultiSet
        d_351_v213_ = _dafny.MultiSet([d_350_v212_, d_350_v212_])
        hi0_ = (d_351_v213_).cardinality
        for d_352_i19_ in range(len(default__.fm14(p0, p0, p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owyumxl"))), globalState)), hi0_):
            d_353_v214_: _dafny.Array
            def lambda17_(d_354_v181_, d_355_p0_, d_356_i19_):
                def lambda18_(d_357_i20_):
                    return D13_DC32(len(d_354_v181_), (0) - (len(_dafny.SeqWithoutIsStrInference([d_355_p0_]))), _dafny.SeqWithoutIsStrInference([d_356_i19_]), d_355_p0_)

                return lambda18_

            init9_ = lambda17_(d_306_v181_, p0, d_352_i19_)
            nw59_ = _dafny.Array(None, 2)
            for i0_9_ in range(nw59_.length(0)):
                nw59_[i0_9_] = init9_(i0_9_)
            d_353_v214_ = nw59_
            d_358_v215_: _dafny.Seq
            d_358_v215_ = _dafny.SeqWithoutIsStrInference([p1, p3, d_352_i19_, len(_dafny.SeqWithoutIsStrInference([d_307_v182_ for d_359_i21_ in range(default__.abs(-99))])), p3])
            d_360_v216_: bool
            d_360_v216_ = False
            d_361_v217_: D13
            d_361_v217_ = D13_DC32(p3, p3, d_358_v215_, (d_360_v216_))
            index27_ = default__.safeIndex(966, (d_353_v214_).length(0))
            (d_353_v214_)[index27_] = d_361_v217_
            index28_ = default__.safeIndex(966, (d_353_v214_).length(0))
            (d_353_v214_)[index28_] = d_361_v217_
            (globalState).f26 = p0
            d_306_v181_ = (d_306_v181_) + (_dafny.SeqWithoutIsStrInference([d_307_v182_ for d_362_i22_ in range(default__.abs(809))]))
            d_363_v218_: _dafny.Array
            nw60_ = _dafny.Array(_dafny.Array(None, 0), 29)
            d_363_v218_ = nw60_
            d_364_v219_: _dafny.Array
            nw61_ = _dafny.Array(int(0), 14)
            d_364_v219_ = nw61_
            index29_ = default__.safeIndex(758, (d_363_v218_).length(0))
            (d_363_v218_)[index29_] = d_364_v219_
            index30_ = default__.safeIndex(758, (d_363_v218_).length(0))
            (d_363_v218_)[index30_] = d_364_v219_
        r0 = p0
        d_365_v220_: _dafny.Map
        d_365_v220_ = _dafny.Map({p1: p2})
        d_366_v221_: _dafny.Map
        d_366_v221_ = _dafny.Map({p0: p0})
        d_367_v222_: _dafny.Map
        d_367_v222_ = _dafny.Map({d_365_v220_: len(d_366_v221_)})
        d_368_v223_: _dafny.Seq
        d_368_v223_ = _dafny.SeqWithoutIsStrInference([p1, p3])
        d_369_v224_: _dafny.Array
        nw62_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
        d_369_v224_ = nw62_
        d_370_v225_: C5
        nw63_ = C5()
        nw63_.ctor__(d_367_v222_, _dafny.MultiSet([p0, p0]), d_368_v223_, p2, d_369_v224_)
        d_370_v225_ = nw63_
        d_371_v226_: D22
        d_371_v226_ = D22_DC54(d_307_v182_, d_370_v225_, p0, len(d_306_v181_))
        r1 = (d_371_v226_).cf98
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_372_v0_: _dafny.Array
        nw64_ = _dafny.Array(int(0), 17)
        d_372_v0_ = nw64_
        d_373_v1_: _dafny.Seq
        d_373_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yn"))
        d_374_v2_: _dafny.MultiSet
        d_374_v2_ = _dafny.MultiSet([d_373_v1_])
        d_375_v3_: str
        d_375_v3_ = _dafny.CodePoint('g')
        d_376_v4_: bool
        d_376_v4_ = False
        d_377_v5_: _dafny.Set
        d_377_v5_ = _dafny.Set({d_376_v4_})
        d_378_v6_: int
        d_378_v6_ = 491
        d_379_v7_: _dafny.Set
        d_379_v7_ = _dafny.Set({len(d_377_v5_), d_378_v6_, d_378_v6_})
        d_380_v8_: _dafny.Map
        d_380_v8_ = _dafny.Map({d_375_v3_: len(d_379_v7_)})
        d_381_v9_: _dafny.Seq
        d_381_v9_ = _dafny.SeqWithoutIsStrInference([d_378_v6_])
        d_382_v10_: _dafny.Array
        nw65_ = _dafny.Array(None, 3)
        nw65_[int(0)] = _dafny.SeqWithoutIsStrInference([len(d_380_v8_)])
        nw65_[int(1)] = d_381_v9_
        nw65_[int(2)] = d_381_v9_
        d_382_v10_ = nw65_
        d_383_v11_: _dafny.Array
        nw66_ = _dafny.Array(False, 26)
        d_383_v11_ = nw66_
        d_384_v12_: _dafny.Map
        d_384_v12_ = _dafny.Map({d_373_v1_: d_383_v11_})
        d_385_v13_: _dafny.Set
        d_385_v13_ = _dafny.Set({d_373_v1_})
        d_386_globalState_: GlobalState
        nw67_ = GlobalState()
        def iife36_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in _dafny.IntegerRange(979, 540):
                d_387_v14_: int = compr_20_
                if ((979) <= (d_387_v14_)) and ((d_387_v14_) < (540)):
                    coll20_[(d_387_v14_) - (len(d_379_v7_))] = d_378_v6_
            return _dafny.Map(coll20_)
        nw67_.ctor__(367, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thcvtep")), d_372_v0_, 880, 597, 980, False, True, (d_374_v2_) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kseko"))])), False, d_382_v10_, 683, 646, d_384_v12_, False, 368, False, (d_385_v13_).intersection(d_385_v13_), -783, 756, True, 853, iife36_()
        , d_383_v11_, 539, -892, True)
        d_386_globalState_ = nw67_
        rhs28_ = d_376_v4_
        rhs29_ = d_378_v6_
        lhs16_ = d_386_globalState_
        d_376_v4_ = rhs28_
        lhs16_.f0 = rhs29_
        if default__.fm0(default__.safeModuloInt(d_378_v6_, d_378_v6_), d_386_globalState_):
            if not (d_376_v4_) or (d_376_v4_):
                d_376_v4_ = d_376_v4_
                d_373_v1_ = (((d_373_v1_) + (d_373_v1_)).set(default__.safeIndex(d_378_v6_, len((d_373_v1_) + (d_373_v1_))), d_375_v3_)) + ((d_373_v1_) + (d_373_v1_))
                d_388_v15_: _dafny.Map
                d_388_v15_ = _dafny.Map({d_376_v4_: d_378_v6_})
                d_388_v15_ = (d_388_v15_).set(d_376_v4_, len(d_373_v1_))
                (d_386_globalState_).f26 = d_376_v4_
                d_389_v16_: _dafny.MultiSet
                d_389_v16_ = _dafny.MultiSet([d_376_v4_])
                d_390_v17_: _dafny.Seq
                d_390_v17_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_376_v4_, d_376_v4_])).intersection(d_389_v16_), d_389_v16_, (_dafny.MultiSet([d_376_v4_, d_376_v4_, d_376_v4_])) | (d_389_v16_)])
                d_390_v17_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_376_v4_]) for d_391_i0_ in range(default__.abs(-117))])
            elif True:
                (d_386_globalState_).f26 = d_376_v4_
                (d_386_globalState_).f18 = default__.fm1(d_386_globalState_)
                d_392_v18_: _dafny.Array
                nw68_ = _dafny.Array(_dafny.Map({}), 13)
                d_392_v18_ = nw68_
                rhs30_ = d_378_v6_
                rhs31_ = d_392_v18_
                lhs17_ = d_386_globalState_
                lhs17_.f24 = rhs30_
                d_392_v18_ = rhs31_
                d_393_v19_: bool
                d_394_v20_: int
                out0_: bool
                out1_: int
                out0_, out1_ = default__.m0((d_376_v4_) == (d_376_v4_), d_378_v6_, d_378_v6_, d_378_v6_, d_386_globalState_)
                d_393_v19_ = out0_
                d_394_v20_ = out1_
                d_393_v19_ = False
            d_395_v21_: _dafny.Seq
            d_395_v21_ = _dafny.SeqWithoutIsStrInference([d_372_v0_, d_372_v0_, d_372_v0_])
            d_395_v21_ = (d_395_v21_) + ((d_395_v21_) + (d_395_v21_))
            if not(d_376_v4_):
                (d_386_globalState_).f0 = d_378_v6_
                d_396_v22_: _dafny.Map
                d_396_v22_ = _dafny.Map({d_378_v6_: d_376_v4_})
                d_396_v22_ = (d_396_v22_).set(d_378_v6_, True)
                d_397_v23_: _dafny.Array
                nw69_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
                d_397_v23_ = nw69_
                d_398_v24_: C4
                nw70_ = C4()
                nw70_.ctor__(d_378_v6_, d_397_v23_)
                d_398_v24_ = nw70_
                d_399_v25_: _dafny.MultiSet
                d_399_v25_ = _dafny.MultiSet([d_378_v6_])
                d_400_v26_: C0
                nw71_ = C0()
                nw71_.ctor__((0) - (((d_399_v25_)[len(d_381_v9_)] if (len(d_381_v9_)) in (d_399_v25_) else d_378_v6_)), d_373_v1_)
                d_400_v26_ = nw71_
                d_378_v6_ = (d_400_v26_).f36
            elif True:
                d_383_v11_ = d_383_v11_
                d_401_v27_: _dafny.Map
                d_401_v27_ = _dafny.Map({205: d_376_v4_})
                (d_386_globalState_).f26 = (((d_381_v9_) + (d_381_v9_)).set(default__.safeIndex(d_378_v6_, len((d_381_v9_) + (d_381_v9_))), len(d_401_v27_))) == (_dafny.SeqWithoutIsStrInference([d_378_v6_ for d_402_i1_ in range(default__.abs(-955))]))
                (d_386_globalState_).f15 = (d_378_v6_) - (-306)
                d_403_v28_: D16
                d_403_v28_ = D16_DC40(True, (0) - (d_378_v6_), 855)
                d_404_v29_: _dafny.Map
                d_404_v29_ = _dafny.Map({d_378_v6_: d_403_v28_})
                (d_386_globalState_).f21 = (0) - ((d_378_v6_) - (len(d_404_v29_)))
                d_405_v30_: _dafny.Map
                d_405_v30_ = _dafny.Map({d_378_v6_: 189})
                d_405_v30_ = d_405_v30_
            index31_ = default__.safeIndex(443, (d_383_v11_).length(0))
            (d_383_v11_)[index31_] = d_376_v4_
            index32_ = default__.safeIndex(283, (d_372_v0_).length(0))
            (d_372_v0_)[index32_] = d_378_v6_
            d_406_v31_: _dafny.Seq
            d_406_v31_ = _dafny.SeqWithoutIsStrInference([d_381_v9_, d_381_v9_])
            d_407_v32_: D9
            d_407_v32_ = D9_DC17(d_379_v7_)
            d_408_v33_: _dafny.Seq
            d_408_v33_ = _dafny.SeqWithoutIsStrInference([d_407_v32_])
            d_409_v34_: _dafny.Map
            d_409_v34_ = _dafny.Map({d_408_v33_: d_372_v0_})
            index33_ = default__.safeIndex(443, (d_383_v11_).length(0))
            index34_ = default__.safeIndex(283, (d_372_v0_).length(0))
            rhs32_ = (d_406_v31_) != (_dafny.SeqWithoutIsStrInference([d_381_v9_]))
            rhs33_ = not(d_376_v4_)
            rhs34_ = (0) - (default__.safeModuloInt(d_378_v6_, d_378_v6_))
            rhs35_ = not (d_376_v4_) or ((d_376_v4_) or (False))
            rhs36_ = (_dafny.SeqWithoutIsStrInference([D9_DC17(_dafny.Set({d_378_v6_})) for d_410_i2_ in range(default__.abs(-536))])) not in ((d_409_v34_) | (d_409_v34_))
            lhs18_ = d_386_globalState_
            lhs19_ = d_383_v11_
            lhs20_ = default__.safeIndex(443, (d_383_v11_).length(0))
            lhs21_ = d_372_v0_
            lhs22_ = default__.safeIndex(283, (d_372_v0_).length(0))
            lhs18_.f26 = rhs32_
            lhs19_[lhs20_] = rhs33_
            lhs21_[lhs22_] = rhs34_
            d_376_v4_ = rhs35_
            d_376_v4_ = rhs36_
            d_411_v35_: _dafny.MultiSet
            d_411_v35_ = _dafny.MultiSet([(d_383_v11_)[default__.safeIndex(443, (d_383_v11_).length(0))], d_376_v4_, ((d_383_v11_)[default__.safeIndex(443, (d_383_v11_).length(0))]) == (d_376_v4_)])
            d_411_v35_ = d_411_v35_
        elif True:
            d_412_v36_: _dafny.Seq
            d_412_v36_ = _dafny.SeqWithoutIsStrInference([(d_378_v6_) > (d_378_v6_)])
            d_412_v36_ = (d_412_v36_) + (d_412_v36_)
            d_413_v37_: bool
            d_414_v38_: int
            out2_: bool
            out3_: int
            out2_, out3_ = default__.m0(d_376_v4_, len(d_381_v9_), d_378_v6_, d_378_v6_, d_386_globalState_)
            d_413_v37_ = out2_
            d_414_v38_ = out3_
            if d_376_v4_:
                d_415_v39_: _dafny.Map
                d_415_v39_ = _dafny.Map({d_378_v6_: d_376_v4_})
                d_416_v40_: bool
                d_417_v41_: int
                out4_: bool
                out5_: int
                out4_, out5_ = default__.m0(d_413_v37_, (d_414_v38_) - (d_378_v6_), d_414_v38_, (_dafny.MultiSet([d_413_v37_, d_413_v37_, ((d_415_v39_)[d_414_v38_] if (d_414_v38_) in (d_415_v39_) else d_376_v4_)])).cardinality, d_386_globalState_)
                d_416_v40_ = out4_
                d_417_v41_ = out5_
                d_375_v3_ = default__.fm12(default__.safeModuloInt(d_378_v6_, d_378_v6_), d_416_v40_, d_386_globalState_)
                d_418_v42_: _dafny.Map
                d_418_v42_ = _dafny.Map({d_375_v3_: d_416_v40_})
                d_419_v43_: D2
                d_419_v43_ = D2_DC3(d_418_v42_)
                d_420_v44_: _dafny.Map
                d_420_v44_ = _dafny.Map({d_419_v43_: d_378_v6_})
                d_420_v44_ = (d_420_v44_).set(d_419_v43_, (d_417_v41_) - (d_378_v6_))
                (d_386_globalState_).f26 = d_376_v4_
                d_421_v45_: _dafny.MultiSet
                d_421_v45_ = _dafny.MultiSet([False, d_376_v4_, not(d_376_v4_)])
                d_422_v46_: _dafny.Array
                nw72_ = _dafny.Array(None, 10)
                nw72_[int(0)] = d_373_v1_
                nw72_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
                nw72_[int(2)] = d_373_v1_
                nw72_[int(3)] = d_373_v1_
                nw72_[int(4)] = d_373_v1_
                nw72_[int(5)] = default__.fm8(d_386_globalState_)
                nw72_[int(6)] = d_373_v1_
                nw72_[int(7)] = (d_373_v1_).set(default__.safeIndex(d_414_v38_, len(d_373_v1_)), d_375_v3_)
                nw72_[int(8)] = d_373_v1_
                nw72_[int(9)] = d_373_v1_
                d_422_v46_ = nw72_
                d_423_v47_: C1
                nw73_ = C1()
                nw73_.ctor__(d_416_v40_, d_421_v45_, d_381_v9_, d_378_v6_, d_422_v46_)
                d_423_v47_ = nw73_
                d_423_v47_ = d_423_v47_
            elif True:
                d_376_v4_ = (d_376_v4_ if d_413_v37_ else d_376_v4_)
                d_424_v48_: D13
                d_424_v48_ = D13_DC32((0) - (d_378_v6_), d_414_v38_, d_381_v9_, d_413_v37_)
                d_425_v49_: bool
                d_426_v50_: int
                out6_: bool
                out7_: int
                out6_, out7_ = default__.m0(True, (len((d_424_v48_).cf59)) * (d_414_v38_), ((0) - (-481)) * (d_414_v38_), default__.safeModuloInt(default__.fm1(d_386_globalState_), d_378_v6_), d_386_globalState_)
                d_425_v49_ = out6_
                d_426_v50_ = out7_
                (d_386_globalState_).f26 = True
                d_427_v51_: _dafny.Map
                d_427_v51_ = _dafny.Map({-243: d_414_v38_})
                d_428_v52_: _dafny.Map
                d_428_v52_ = _dafny.Map({d_427_v51_: d_426_v50_})
                d_429_v53_: _dafny.MultiSet
                d_429_v53_ = _dafny.MultiSet([d_425_v49_])
                d_430_v54_: _dafny.Array
                def lambda19_(d_431_v1_):
                    def lambda20_(d_432_i3_):
                        return d_431_v1_

                    return lambda20_

                init10_ = lambda19_(d_373_v1_)
                nw74_ = _dafny.Array(None, 16)
                for i0_10_ in range(nw74_.length(0)):
                    nw74_[i0_10_] = init10_(i0_10_)
                d_430_v54_ = nw74_
                d_433_v55_: C5
                nw75_ = C5()
                nw75_.ctor__((d_428_v52_) | (default__.fm54(d_376_v4_, d_425_v49_, d_386_globalState_)), d_429_v53_, d_381_v9_, d_426_v50_, d_430_v54_)
                d_433_v55_ = nw75_
                d_434_v56_: _dafny.Map
                d_434_v56_ = _dafny.Map({d_426_v50_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijuo"))})
                d_435_v57_: C0
                nw76_ = C0()
                nw76_.ctor__(448, ((d_434_v56_)[-214] if (-214) in (d_434_v56_) else d_373_v1_))
                d_435_v57_ = nw76_
                d_435_v57_ = d_435_v57_
            d_436_v58_: _dafny.MultiSet
            d_436_v58_ = _dafny.MultiSet([d_412_v36_, d_412_v36_, d_412_v36_, d_412_v36_, d_412_v36_])
            d_437_v60_: _dafny.Set
            def iife37_():
                coll21_ = _dafny.Set()
                compr_21_: _dafny.Seq
                for compr_21_ in ((d_436_v58_).set(_dafny.SeqWithoutIsStrInference([True]), default__.abs(d_414_v38_))).Elements:
                    d_438_v59_: _dafny.Seq = compr_21_
                    if (d_438_v59_) in ((d_436_v58_).set(_dafny.SeqWithoutIsStrInference([True]), default__.abs(d_414_v38_))):
                        coll21_ = coll21_.union(_dafny.Set([d_438_v59_]))
                return _dafny.Set(coll21_)
            d_437_v60_ = (_dafny.Set({d_412_v36_, _dafny.SeqWithoutIsStrInference([d_413_v37_, d_376_v4_])})) - (iife37_()
            )
            source10_ = d_437_v60_
            d_439___mcc_h0_ = source10_
            d_440_cf62_ = d_439___mcc_h0_
            (d_386_globalState_).f26 = not (not(not(d_376_v4_))) or (d_376_v4_)
            def lambda21_(d_441_v38_):
                def lambda22_(d_442_i4_):
                    return default__.safeDivisionInt(d_442_i4_, d_441_v38_)

                return lambda22_

            init11_ = lambda21_(d_414_v38_)
            nw77_ = _dafny.Array(None, 13)
            for i0_11_ in range(nw77_.length(0)):
                nw77_[i0_11_] = init11_(i0_11_)
            d_372_v0_ = nw77_
            d_373_v1_ = _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_443_i5_ in range(default__.abs(675))])
            d_375_v3_ = d_375_v3_
            (d_386_globalState_).f18 = d_378_v6_
        d_444_v61_: D11
        d_444_v61_ = D11_DC25(D11_DC23(_dafny.MultiSet([d_379_v7_])))
        d_445_v62_: _dafny.MultiSet
        d_445_v62_ = _dafny.MultiSet([d_379_v7_])
        d_446_v63_: D11
        d_446_v63_ = D11_DC23(d_445_v62_)
        pat_let_tv4_ = d_446_v63_
        def iife38_(_pat_let8_0):
            def iife39_(d_447_dt__update__tmp_h0_):
                def iife40_(_pat_let9_0):
                    def iife41_(d_448_dt__update_hcf45_h0_):
                        return D11_DC25(d_448_dt__update_hcf45_h0_)
                    return iife41_(_pat_let9_0)
                return iife40_(pat_let_tv4_)
            return iife39_(_pat_let8_0)
        source11_ = iife38_(d_444_v61_)
        if source11_.is_DC24:
            d_449___mcc_h1_ = source11_.cf43
            d_450___mcc_h2_ = source11_.cf44
            d_451_cf44_ = d_450___mcc_h2_
            d_452_cf43_ = d_449___mcc_h1_
            d_453_v64_: _dafny.Map
            d_453_v64_ = _dafny.Map({len(d_373_v1_): d_452_cf43_})
            d_454_v65_: _dafny.MultiSet
            d_454_v65_ = _dafny.MultiSet([d_376_v4_])
            d_455_v66_: T2
            nw78_ = C2()
            nw78_.ctor__(d_454_v65_, d_378_v6_)
            d_455_v66_ = nw78_
            d_456_v67_: _dafny.Set
            d_456_v67_ = _dafny.Set({d_455_v66_, d_455_v66_, d_455_v66_})
            d_453_v64_ = (d_453_v64_).set((d_381_v9_)[default__.safeIndex(d_378_v6_, len(d_381_v9_))], len(d_456_v67_))
            d_457_v68_: _dafny.Array
            d_458_v69_: int
            out8_: _dafny.Array
            out9_: int
            out8_, out9_ = (d_455_v66_).m2(d_376_v4_, d_376_v4_, (d_454_v65_).cardinality, (0) - (d_452_cf43_), d_386_globalState_)
            d_457_v68_ = out8_
            d_458_v69_ = out9_
            d_459_v70_: _dafny.Array
            nw79_ = _dafny.Array(None, 22)
            nw79_[int(0)] = d_373_v1_
            nw79_[int(1)] = d_373_v1_
            nw79_[int(2)] = _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_460_i6_ in range(default__.abs(-447))])
            nw79_[int(3)] = d_373_v1_
            nw79_[int(4)] = d_373_v1_
            nw79_[int(5)] = d_373_v1_
            nw79_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukgaklleg"))
            nw79_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hl"))
            nw79_[int(8)] = d_373_v1_
            nw79_[int(9)] = d_373_v1_
            nw79_[int(10)] = d_373_v1_
            nw79_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wkbli"))
            nw79_[int(12)] = d_373_v1_
            nw79_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usfy"))
            nw79_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yakla"))
            nw79_[int(15)] = d_373_v1_
            nw79_[int(16)] = d_373_v1_
            nw79_[int(17)] = d_373_v1_
            nw79_[int(18)] = d_373_v1_
            nw79_[int(19)] = d_373_v1_
            nw79_[int(20)] = d_373_v1_
            nw79_[int(21)] = d_373_v1_
            d_459_v70_ = nw79_
            d_461_v71_: _dafny.Map
            d_461_v71_ = _dafny.Map({default__.fm1(d_386_globalState_): d_459_v70_})
            d_462_v72_: C4
            nw80_ = C4()
            nw80_.ctor__(default__.fm1(d_386_globalState_), ((d_461_v71_)[d_458_v69_] if (d_458_v69_) in (d_461_v71_) else d_459_v70_))
            d_462_v72_ = nw80_
            rhs37_ = d_376_v4_
            rhs38_ = d_376_v4_
            rhs39_ = d_458_v69_
            rhs40_ = d_376_v4_
            lhs23_ = d_386_globalState_
            lhs24_ = d_386_globalState_
            lhs25_ = d_386_globalState_
            d_376_v4_ = rhs37_
            lhs23_.f26 = rhs38_
            lhs24_.f15 = rhs39_
            lhs25_.f26 = rhs40_
        elif source11_.is_DC23:
            d_463___mcc_h3_ = source11_.cf42
            d_464_cf42_ = d_463___mcc_h3_
            d_465_v73_: bool
            d_466_v74_: int
            out10_: bool
            out11_: int
            out10_, out11_ = default__.m0((d_378_v6_) >= (d_378_v6_), (84) - (d_378_v6_), d_378_v6_, (d_378_v6_ if d_376_v4_ else len(d_373_v1_)), d_386_globalState_)
            d_465_v73_ = out10_
            d_466_v74_ = out11_
            if d_376_v4_:
                d_467_v75_: bool
                d_468_v76_: int
                out12_: bool
                out13_: int
                out12_, out13_ = default__.m0(d_376_v4_, d_378_v6_, d_466_v74_, d_378_v6_, d_386_globalState_)
                d_467_v75_ = out12_
                d_468_v76_ = out13_
                d_469_v77_: _dafny.MultiSet
                d_469_v77_ = _dafny.MultiSet([d_376_v4_, d_467_v75_])
                d_470_v78_: D13
                d_470_v78_ = D13_DC32(d_378_v6_, (0) - ((d_381_v9_)[default__.safeIndex((0) - (d_466_v74_), len(d_381_v9_))]), d_381_v9_, d_376_v4_)
                d_471_v79_: _dafny.Map
                d_471_v79_ = _dafny.Map({len(d_373_v1_): d_467_v75_})
                d_472_v80_: _dafny.Map
                d_472_v80_ = _dafny.Map({d_467_v75_: True})
                d_473_v81_: _dafny.Map
                d_473_v81_ = _dafny.Map({((d_471_v79_)[len(d_472_v80_)] if (len(d_472_v80_)) in (d_471_v79_) else d_376_v4_): d_468_v76_})
                d_474_v82_: _dafny.Array
                nw81_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
                d_474_v82_ = nw81_
                d_475_v83_: D17
                d_475_v83_ = D17_DC43(d_474_v82_)
                d_476_v84_: T2
                nw82_ = C6()
                nw82_.ctor__(d_465_v73_, default__.fm1(d_386_globalState_), d_469_v77_, (d_470_v78_).cf59, (d_466_v74_) - (len(d_473_v81_)), (d_475_v83_).cf82, d_469_v77_, d_468_v76_)
                d_476_v84_ = nw82_
                d_476_v84_ = (d_476_v84_ if d_467_v75_ else d_476_v84_)
                d_477_v85_: _dafny.Array
                nw83_ = _dafny.Array(None, 1)
                nw83_[int(0)] = d_471_v79_
                d_477_v85_ = nw83_
                d_478_v86_: _dafny.Array
                d_478_v86_ = d_477_v85_
                d_479_v87_: _dafny.Map
                d_479_v87_ = _dafny.Map({(0) - (d_378_v6_): (d_477_v85_)})
                d_479_v87_ = (d_479_v87_).set(d_468_v76_, d_477_v85_)
                d_480_v88_: C1
                nw84_ = C1()
                nw84_.ctor__(True, (d_469_v77_).set(d_465_v73_, default__.abs(len(d_373_v1_))), d_381_v9_, 805, d_474_v82_)
                d_480_v88_ = nw84_
                d_481_v89_: _dafny.MultiSet
                d_481_v89_ = _dafny.MultiSet([d_480_v88_])
                d_482_v90_: C6
                nw85_ = C6()
                nw85_.ctor__(((d_476_v84_).f32) == ((d_481_v89_).cardinality), default__.safeModuloInt((d_476_v84_).f32, len(d_373_v1_)), (d_476_v84_).f31, d_381_v9_, default__.safeDivisionInt((0) - (d_378_v6_), -584), d_474_v82_, (d_476_v84_).f31, d_378_v6_)
                d_482_v90_ = nw85_
                d_483_v91_: _dafny.Map
                d_483_v91_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_375_v3_ for d_484_i7_ in range(default__.abs(4))]): (d_482_v90_ if (d_482_v90_).f33 else d_482_v90_)})
                d_482_v90_ = ((d_483_v91_)[(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcvfglbs"))) + (d_373_v1_)] if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcvfglbs"))) + (d_373_v1_)) in (d_483_v91_) else d_482_v90_)
                d_485_v92_: _dafny.Seq
                d_485_v92_ = d_373_v1_
                d_486_v93_: bool
                d_487_v94_: _dafny.Array
                d_488_v95_: _dafny.Map
                out14_: bool
                out15_: _dafny.Array
                out16_: _dafny.Map
                out14_, out15_, out16_ = (d_480_v88_).m5((d_480_v88_).fm11(not((d_482_v90_).f33), len(d_373_v1_), d_485_v92_, len(d_373_v1_), d_386_globalState_), default__.fm0(d_378_v6_, d_386_globalState_), d_386_globalState_)
                d_486_v93_ = out14_
                d_487_v94_ = out15_
                d_488_v95_ = out16_
            elif True:
                (d_386_globalState_).f26 = (d_379_v7_).isdisjoint(d_379_v7_)
                d_489_v96_: D10
                d_489_v96_ = D10_DC21(d_381_v9_)
                d_490_v97_: _dafny.Map
                d_490_v97_ = _dafny.Map({(d_376_v4_ if d_376_v4_ else d_376_v4_): d_489_v96_})
                d_490_v97_ = (d_490_v97_).set(True, d_489_v96_)
                d_491_v98_: _dafny.Seq
                d_491_v98_ = d_373_v1_
                d_492_v99_: _dafny.Seq
                d_492_v99_ = _dafny.SeqWithoutIsStrInference([default__.fm14(d_376_v4_, False, d_465_v73_, d_378_v6_, d_386_globalState_)])
                d_493_v100_: _dafny.Map
                d_493_v100_ = _dafny.Map({d_466_v74_: True})
                d_494_v101_: _dafny.Array
                nw86_ = _dafny.Array(None, 27)
                nw86_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgvfq"))
                nw86_[int(1)] = d_373_v1_
                nw86_[int(2)] = d_373_v1_
                nw86_[int(3)] = d_373_v1_
                nw86_[int(4)] = d_373_v1_
                nw86_[int(5)] = (_dafny.SeqWithoutIsStrInference([d_375_v3_ for d_495_i8_ in range(default__.abs(433))])).set(default__.safeIndex(d_466_v74_, len(_dafny.SeqWithoutIsStrInference([d_375_v3_ for d_496_i8_ in range(default__.abs(433))]))), d_375_v3_)
                nw86_[int(6)] = ((d_491_v98_)).set(default__.safeIndex(default__.fm1(d_386_globalState_), len((d_491_v98_))), d_375_v3_)
                nw86_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
                nw86_[int(8)] = d_373_v1_
                nw86_[int(9)] = d_373_v1_
                nw86_[int(10)] = default__.fm31(d_465_v73_, d_376_v4_, d_386_globalState_)
                nw86_[int(11)] = _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_497_i9_ in range(default__.abs(104))])
                nw86_[int(12)] = (d_492_v99_)[default__.safeIndex(len(d_493_v100_), len(d_492_v99_))]
                nw86_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgwx"))
                nw86_[int(14)] = d_373_v1_
                nw86_[int(15)] = d_373_v1_
                nw86_[int(16)] = d_373_v1_
                nw86_[int(17)] = d_373_v1_
                nw86_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kx"))
                nw86_[int(19)] = _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_498_i10_ in range(default__.abs(486))])
                nw86_[int(20)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olpvqovqh"))).set(default__.safeIndex(d_466_v74_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olpvqovqh")))), d_375_v3_)
                nw86_[int(21)] = d_373_v1_
                nw86_[int(22)] = d_373_v1_
                nw86_[int(23)] = d_373_v1_
                nw86_[int(24)] = d_373_v1_
                nw86_[int(25)] = _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_499_i11_ in range(default__.abs(590))])
                nw86_[int(26)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfnpbmx"))
                d_494_v101_ = nw86_
                d_500_v102_: C3
                nw87_ = C3()
                nw87_.ctor__(d_466_v74_, d_494_v101_)
                d_500_v102_ = nw87_
                d_501_v103_: D13
                d_501_v103_ = D13_DC32(d_378_v6_, d_466_v74_, d_381_v9_, d_465_v73_)
                d_502_v104_: _dafny.Map
                d_502_v104_ = _dafny.Map({(d_501_v103_).cf59: d_500_v102_})
                d_500_v102_ = ((d_502_v104_)[d_381_v9_] if (d_381_v9_) in (d_502_v104_) else d_500_v102_)
                d_503_v105_: _dafny.MultiSet
                d_503_v105_ = _dafny.MultiSet([201, d_466_v74_])
                d_504_v106_: T0
                nw88_ = C3()
                nw88_.ctor__(d_378_v6_, d_494_v101_)
                d_504_v106_ = nw88_
                d_505_v107_: _dafny.Seq
                d_505_v107_ = _dafny.SeqWithoutIsStrInference([d_504_v106_, d_504_v106_])
                d_506_v108_: D9
                d_506_v108_ = D9_DC18(True)
                d_507_v109_: _dafny.Seq
                d_507_v109_ = _dafny.SeqWithoutIsStrInference([(d_506_v108_).cf32])
                d_508_v110_: D16
                d_508_v110_ = D16_DC40(d_465_v73_, (0) - (((d_503_v105_)[len((d_373_v1_).set(default__.safeIndex(len(d_505_v107_), len(d_373_v1_)), d_375_v3_))] if (len((d_373_v1_).set(default__.safeIndex(len(d_505_v107_), len(d_373_v1_)), d_375_v3_))) in (d_503_v105_) else d_378_v6_)), len(d_507_v109_))
                d_509_v111_: _dafny.MultiSet
                d_509_v111_ = _dafny.MultiSet([d_376_v4_, (d_508_v110_).cf74])
                d_510_v112_: C6
                nw89_ = C6()
                nw89_.ctor__(d_465_v73_, 549, d_509_v111_, d_381_v9_, d_466_v74_, d_494_v101_, d_509_v111_, (d_504_v106_).f27)
                d_510_v112_ = nw89_
                d_511_v113_: _dafny.Map
                d_511_v113_ = _dafny.Map({d_375_v3_: d_510_v112_})
                d_511_v113_ = (d_511_v113_).set(d_375_v3_, d_510_v112_)
                d_380_v8_ = (d_380_v8_).set(d_375_v3_, d_510_v112_.f34)
            d_512_v114_: bool
            d_513_v115_: int
            out17_: bool
            out18_: int
            out17_, out18_ = default__.m0(d_465_v73_, d_466_v74_, d_466_v74_, len(_dafny.SeqWithoutIsStrInference([len(d_373_v1_) for d_514_i12_ in range(default__.abs(391))])), d_386_globalState_)
            d_512_v114_ = out17_
            d_513_v115_ = out18_
            index35_ = default__.safeIndex(395, (d_383_v11_).length(0))
            (d_383_v11_)[index35_] = d_465_v73_
            d_515_v116_: _dafny.MultiSet
            d_515_v116_ = _dafny.MultiSet([d_513_v115_, d_513_v115_])
            index36_ = default__.safeIndex(395, (d_383_v11_).length(0))
            rhs41_ = d_512_v114_
            rhs42_ = d_465_v73_
            rhs43_ = not(not (not (d_512_v114_) or (d_512_v114_)) or ((d_512_v114_ if default__.fm0((d_515_v116_).cardinality, d_386_globalState_) else not(d_512_v114_))))
            lhs26_ = d_383_v11_
            lhs27_ = default__.safeIndex(395, (d_383_v11_).length(0))
            lhs28_ = d_386_globalState_
            lhs26_[lhs27_] = rhs41_
            lhs28_.f26 = rhs42_
            d_512_v114_ = rhs43_
        elif True:
            d_516___mcc_h4_ = source11_.cf45
            d_517_cf45_ = d_516___mcc_h4_
            if d_376_v4_:
                d_518_v117_: _dafny.Map
                d_518_v117_ = _dafny.Map({d_378_v6_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrjwq"))})
                d_519_v118_: _dafny.Map
                d_519_v118_ = _dafny.Map({d_376_v4_: d_376_v4_})
                d_520_v119_: _dafny.Array
                nw90_ = _dafny.Array(None, 27)
                nw90_[int(0)] = d_373_v1_
                nw90_[int(1)] = d_373_v1_
                nw90_[int(2)] = d_373_v1_
                nw90_[int(3)] = d_373_v1_
                nw90_[int(4)] = d_373_v1_
                nw90_[int(5)] = d_373_v1_
                nw90_[int(6)] = d_373_v1_
                nw90_[int(7)] = d_373_v1_
                nw90_[int(8)] = d_373_v1_
                nw90_[int(9)] = d_373_v1_
                nw90_[int(10)] = d_373_v1_
                nw90_[int(11)] = d_373_v1_
                nw90_[int(12)] = d_373_v1_
                nw90_[int(13)] = d_373_v1_
                nw90_[int(14)] = d_373_v1_
                nw90_[int(15)] = d_373_v1_
                nw90_[int(16)] = d_373_v1_
                nw90_[int(17)] = _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_521_i13_ in range(default__.abs(43))])
                nw90_[int(18)] = default__.fm37(d_378_v6_, d_378_v6_, d_386_globalState_)
                nw90_[int(19)] = d_373_v1_
                nw90_[int(20)] = d_373_v1_
                nw90_[int(21)] = d_373_v1_
                nw90_[int(22)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "by"))
                nw90_[int(23)] = _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_522_i14_ in range(default__.abs(31))])
                nw90_[int(24)] = d_373_v1_
                nw90_[int(25)] = ((d_518_v117_)[len(d_519_v118_)] if (len(d_519_v118_)) in (d_518_v117_) else d_373_v1_)
                nw90_[int(26)] = d_373_v1_
                d_520_v119_ = nw90_
                d_523_v120_: C4
                nw91_ = C4()
                nw91_.ctor__(len(d_381_v9_), d_520_v119_)
                d_523_v120_ = nw91_
                d_524_v121_: _dafny.MultiSet
                d_524_v121_ = _dafny.MultiSet([d_523_v120_])
                d_525_v122_: _dafny.Map
                d_525_v122_ = _dafny.Map({d_378_v6_: d_381_v9_})
                d_526_v123_: _dafny.Seq
                d_526_v123_ = _dafny.SeqWithoutIsStrInference([d_376_v4_])
                (d_386_globalState_).f21 = (((d_524_v121_).cardinality) * (len(((d_525_v122_)[len(d_526_v123_)] if (len(d_526_v123_)) in (d_525_v122_) else d_381_v9_)))) * ((d_378_v6_) - (d_378_v6_))
                d_527_v124_: _dafny.MultiSet
                d_527_v124_ = _dafny.MultiSet([d_383_v11_])
                rhs44_ = (d_527_v124_) | (d_527_v124_)
                rhs45_ = ((len(d_373_v1_)) + ((0) - (d_378_v6_))) >= (d_378_v6_)
                lhs29_ = d_386_globalState_
                d_527_v124_ = rhs44_
                lhs29_.f26 = rhs45_
                d_373_v1_ = ((d_373_v1_) + ((_dafny.SeqWithoutIsStrInference([d_375_v3_ for d_528_i15_ in range(default__.abs(8))])).set(default__.safeIndex(d_378_v6_, len(_dafny.SeqWithoutIsStrInference([d_375_v3_ for d_529_i15_ in range(default__.abs(8))]))), d_375_v3_))) + (d_373_v1_)
                (d_386_globalState_).f26 = (d_376_v4_) == (d_376_v4_)
                d_530_v125_: C0
                nw92_ = C0()
                nw92_.ctor__(745, d_373_v1_)
                d_530_v125_ = nw92_
                d_531_v126_: int
                out19_: int
                out19_ = (d_523_v120_).m11(d_530_v125_, (d_376_v4_) and (d_376_v4_), d_386_globalState_)
                d_531_v126_ = out19_
            elif True:
                (d_386_globalState_).f26 = d_376_v4_
                (d_386_globalState_).f26 = d_376_v4_
                d_376_v4_ = (d_374_v2_).isdisjoint((d_374_v2_) | (d_374_v2_))
                d_532_v127_: _dafny.Array
                nw93_ = _dafny.Array(None, 24)
                nw93_[int(0)] = d_373_v1_
                nw93_[int(1)] = _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_533_i16_ in range(default__.abs(232))])
                nw93_[int(2)] = d_373_v1_
                nw93_[int(3)] = d_373_v1_
                nw93_[int(4)] = d_373_v1_
                nw93_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_534_i17_ in range(default__.abs(195))])
                nw93_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfnbai"))
                nw93_[int(7)] = d_373_v1_
                nw93_[int(8)] = d_373_v1_
                nw93_[int(9)] = d_373_v1_
                nw93_[int(10)] = d_373_v1_
                nw93_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apfximgjv"))
                nw93_[int(12)] = d_373_v1_
                nw93_[int(13)] = d_373_v1_
                nw93_[int(14)] = d_373_v1_
                nw93_[int(15)] = _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_535_i18_ in range(default__.abs(983))])
                nw93_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gadjbunu"))
                nw93_[int(17)] = d_373_v1_
                nw93_[int(18)] = _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_536_i19_ in range(default__.abs(684))])
                nw93_[int(19)] = d_373_v1_
                nw93_[int(20)] = d_373_v1_
                nw93_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwlwggmnv"))
                nw93_[int(22)] = d_373_v1_
                nw93_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukpupoan"))
                d_532_v127_ = nw93_
                d_537_v128_: _dafny.MultiSet
                d_537_v128_ = _dafny.MultiSet([d_376_v4_])
                d_538_v129_: T2
                nw94_ = C6()
                nw94_.ctor__(d_376_v4_, d_378_v6_, _dafny.MultiSet([d_376_v4_, d_376_v4_]), d_381_v9_, len(d_379_v7_), d_532_v127_, d_537_v128_, d_378_v6_)
                d_538_v129_ = nw94_
                d_539_v130_: _dafny.Set
                d_539_v130_ = _dafny.Set({d_538_v129_})
                (d_386_globalState_).f21 = ((d_378_v6_) + (d_378_v6_) if (d_539_v130_).isdisjoint(d_539_v130_) else len(d_373_v1_))
                d_540_v131_: _dafny.Map
                d_540_v131_ = _dafny.Map({d_372_v0_: (d_538_v129_).f32})
                (d_386_globalState_).f15 = ((d_540_v131_)[d_372_v0_] if (d_372_v0_) in (d_540_v131_) else (0) - (default__.fm1(d_386_globalState_)))
            index37_ = default__.safeIndex(145, (d_383_v11_).length(0))
            (d_383_v11_)[index37_] = (d_376_v4_) or (default__.fm0(d_378_v6_, d_386_globalState_))
            index38_ = default__.safeIndex(145, (d_383_v11_).length(0))
            (d_383_v11_)[index38_] = not (d_376_v4_) or (d_376_v4_)
            d_541_v132_: _dafny.MultiSet
            d_541_v132_ = _dafny.MultiSet([d_376_v4_, (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]])
            d_542_v133_: C2
            nw95_ = C2()
            nw95_.ctor__((d_541_v132_).set(default__.fm0((0) - (d_378_v6_), d_386_globalState_), default__.abs(869)), d_378_v6_)
            d_542_v133_ = nw95_
            d_543_v134_: _dafny.Map
            d_543_v134_ = _dafny.Map({-78: d_542_v133_})
            d_544_v136_: _dafny.Seq
            def iife42_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(987, 775):
                    d_545_v135_: int = compr_22_
                    if ((987) <= (d_545_v135_)) and ((d_545_v135_) < (775)):
                        coll22_[(d_545_v135_) * (len(_dafny.Map({(d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]: d_378_v6_})))] = (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]
                return _dafny.Map(coll22_)
            d_544_v136_ = _dafny.SeqWithoutIsStrInference([iife42_()
            ])
            d_546_v137_: _dafny.Seq
            d_546_v137_ = _dafny.SeqWithoutIsStrInference([d_542_v133_, d_542_v133_])
            d_542_v133_ = ((d_543_v134_)[len((d_544_v136_) + (d_544_v136_))] if (len((d_544_v136_) + (d_544_v136_))) in (d_543_v134_) else (d_546_v137_)[default__.safeIndex(len(default__.fm48((d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))], d_376_v4_, d_378_v6_, d_386_globalState_)), len(d_546_v137_))])
            d_547_v138_: _dafny.Array
            nw96_ = _dafny.Array(None, 3)
            nw96_[int(0)] = d_373_v1_
            nw96_[int(1)] = default__.fm37(d_378_v6_, d_378_v6_, d_386_globalState_)
            nw96_[int(2)] = d_373_v1_
            d_547_v138_ = nw96_
            d_548_v139_: D17
            d_548_v139_ = D17_DC43(d_547_v138_)
            source12_ = d_548_v139_
            if source12_.is_DC44:
                d_549___mcc_h5_ = source12_.cf83
                d_550_cf83_ = d_549___mcc_h5_
                d_551_v140_: _dafny.Seq
                d_551_v140_ = _dafny.SeqWithoutIsStrInference([d_376_v4_])
                d_552_v141_: _dafny.Array
                nw97_ = _dafny.Array(None, 23)
                nw97_[int(0)] = (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]
                nw97_[int(1)] = (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]
                nw97_[int(2)] = False
                nw97_[int(3)] = d_376_v4_
                nw97_[int(4)] = d_376_v4_
                nw97_[int(5)] = (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]
                nw97_[int(6)] = (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]
                nw97_[int(7)] = d_376_v4_
                nw97_[int(8)] = ((d_542_v133_).fm5(d_386_globalState_) if False else d_376_v4_)
                nw97_[int(9)] = True
                nw97_[int(10)] = (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]
                nw97_[int(11)] = (d_551_v140_) <= (d_551_v140_)
                nw97_[int(12)] = not (d_376_v4_) or ((d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))])
                nw97_[int(13)] = d_376_v4_
                nw97_[int(14)] = (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]
                nw97_[int(15)] = (default__.fm1(d_386_globalState_)) == (d_378_v6_)
                nw97_[int(16)] = (d_542_v133_).fm21(d_386_globalState_)
                nw97_[int(17)] = True
                nw97_[int(18)] = False
                nw97_[int(19)] = (d_376_v4_) and ((d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))])
                nw97_[int(20)] = d_376_v4_
                nw97_[int(21)] = d_376_v4_
                nw97_[int(22)] = (d_541_v132_).ispropersubset(_dafny.MultiSet(d_551_v140_))
                d_552_v141_ = nw97_
                d_553_v142_: _dafny.Set
                d_553_v142_ = _dafny.Set({d_552_v141_, d_552_v141_})
                index39_ = default__.safeIndex(145, (d_383_v11_).length(0))
                rhs46_ = d_552_v141_
                rhs47_ = d_373_v1_
                rhs48_ = ((d_553_v142_) - (d_553_v142_)).ispropersubset(_dafny.Set({d_552_v141_}))
                lhs30_ = d_383_v11_
                lhs31_ = default__.safeIndex(145, (d_383_v11_).length(0))
                d_552_v141_ = rhs46_
                d_373_v1_ = rhs47_
                lhs30_[lhs31_] = rhs48_
                index40_ = default__.safeIndex(145, (d_383_v11_).length(0))
                (d_383_v11_)[index40_] = not((d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))])
                d_554_v143_: _dafny.Map
                d_554_v143_ = _dafny.Map({d_376_v4_: d_372_v0_})
                d_555_v144_: D15
                d_555_v144_ = D15_DC37((d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))], d_550_cf83_, d_372_v0_, (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))])
                d_556_v145_: _dafny.Array
                nw98_ = _dafny.Array(None, 18)
                nw98_[int(0)] = d_372_v0_
                nw98_[int(1)] = d_372_v0_
                nw98_[int(2)] = d_372_v0_
                nw98_[int(3)] = (d_372_v0_ if (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))] else ((d_554_v143_)[(d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]] if ((d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]) in (d_554_v143_) else d_372_v0_))
                nw98_[int(4)] = d_372_v0_
                nw98_[int(5)] = (d_555_v144_).cf68
                nw98_[int(6)] = d_372_v0_
                nw98_[int(7)] = d_372_v0_
                nw98_[int(8)] = d_372_v0_
                nw98_[int(9)] = d_372_v0_
                nw98_[int(10)] = d_372_v0_
                nw98_[int(11)] = d_372_v0_
                nw98_[int(12)] = d_372_v0_
                nw98_[int(13)] = d_372_v0_
                nw98_[int(14)] = d_372_v0_
                nw98_[int(15)] = d_372_v0_
                nw98_[int(16)] = d_372_v0_
                nw98_[int(17)] = d_372_v0_
                d_556_v145_ = nw98_
                index41_ = default__.safeIndex(154, (d_556_v145_).length(0))
                (d_556_v145_)[index41_] = d_372_v0_
                d_557_v146_: D12
                d_557_v146_ = D12_DC28(d_550_cf83_, d_378_v6_, d_376_v4_)
                index42_ = default__.safeIndex(154, (d_556_v145_).length(0))
                nw99_ = _dafny.Array(None, 19)
                nw99_[int(0)] = -444
                nw99_[int(1)] = d_550_cf83_
                nw99_[int(2)] = (d_550_cf83_) * (d_550_cf83_)
                nw99_[int(3)] = len(d_551_v140_)
                nw99_[int(4)] = 963
                nw99_[int(5)] = d_378_v6_
                nw99_[int(6)] = d_378_v6_
                nw99_[int(7)] = (d_378_v6_) * (d_378_v6_)
                nw99_[int(8)] = (d_378_v6_) - (d_378_v6_)
                nw99_[int(9)] = 909
                nw99_[int(10)] = d_378_v6_
                nw99_[int(11)] = d_378_v6_
                nw99_[int(12)] = ((d_557_v146_).cf52) * (d_378_v6_)
                nw99_[int(13)] = (d_550_cf83_) * (len(default__.fm9((d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))], d_386_globalState_)))
                nw99_[int(14)] = len(d_373_v1_)
                nw99_[int(15)] = d_378_v6_
                nw99_[int(16)] = d_550_cf83_
                nw99_[int(17)] = (0) - (default__.fm1(d_386_globalState_))
                nw99_[int(18)] = d_378_v6_
                (d_556_v145_)[index42_] = nw99_
                (d_386_globalState_).f15 = (0) - ((len((_dafny.SeqWithoutIsStrInference([d_376_v4_])) + (_dafny.SeqWithoutIsStrInference([True])))) * ((d_378_v6_) + (d_550_cf83_)))
            elif source12_.is_DC43:
                d_558___mcc_h6_ = source12_.cf82
                d_559_cf82_ = d_558___mcc_h6_
                d_376_v4_ = d_376_v4_
                d_560_v147_: _dafny.Seq
                d_560_v147_ = _dafny.SeqWithoutIsStrInference([d_381_v9_])
                d_561_v148_: _dafny.Set
                d_561_v148_ = _dafny.Set({(d_560_v147_)[default__.safeIndex(d_378_v6_, len(d_560_v147_))], default__.fm43(d_376_v4_, d_376_v4_, d_376_v4_, d_386_globalState_)})
                (d_386_globalState_).f15 = len(d_561_v148_)
                d_562_v149_: _dafny.Map
                d_562_v149_ = _dafny.Map({False: (d_381_v9_)[default__.safeIndex(814, len(d_381_v9_))]})
                d_563_v150_: _dafny.Map
                d_563_v150_ = _dafny.Map({d_562_v149_: (423) <= (d_378_v6_)})
                d_564_v151_: C1
                nw100_ = C1()
                nw100_.ctor__(False, d_541_v132_, d_381_v9_, d_378_v6_, d_559_cf82_)
                d_564_v151_ = nw100_
                d_565_v152_: _dafny.Map
                d_565_v152_ = _dafny.Map({d_564_v151_: d_561_v148_})
                d_563_v150_ = (d_563_v150_).set(_dafny.Map({d_376_v4_: len(((d_565_v152_)[d_564_v151_] if (d_564_v151_) in (d_565_v152_) else d_561_v148_))}), (_dafny.SeqWithoutIsStrInference([(d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]])) != (_dafny.SeqWithoutIsStrInference([d_564_v151_.f35, False, d_564_v151_.f35, d_376_v4_, (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]])))
                d_566_v153_: _dafny.Map
                d_566_v153_ = _dafny.Map({664: d_377_v5_})
                d_567_v154_: D11
                d_567_v154_ = D11_DC23(_dafny.MultiSet([d_379_v7_]))
                rhs49_ = _dafny.Map({230: _dafny.Set({d_564_v151_.f35, (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]})})
                rhs50_ = d_567_v154_
                d_566_v153_ = rhs49_
                d_567_v154_ = rhs50_
            elif True:
                d_568___mcc_h7_ = source12_.cf84
                d_569_cf84_ = d_568___mcc_h7_
                d_570_v155_: _dafny.Array
                def lambda23_(d_571_i20_):
                    return _dafny.SeqWithoutIsStrInference([False])

                init12_ = lambda23_
                nw101_ = _dafny.Array(None, 17)
                for i0_12_ in range(nw101_.length(0)):
                    nw101_[i0_12_] = init12_(i0_12_)
                d_570_v155_ = nw101_
                d_572_v156_: _dafny.Map
                d_572_v156_ = _dafny.Map({d_570_v155_: d_381_v9_})
                d_381_v9_ = ((d_572_v156_)[d_570_v155_] if (d_570_v155_) in (d_572_v156_) else d_381_v9_)
                (d_386_globalState_).f26 = (d_383_v11_)[default__.safeIndex(145, (d_383_v11_).length(0))]
                d_573_v157_: C3
                nw102_ = C3()
                nw102_.ctor__(-958, d_547_v138_)
                d_573_v157_ = nw102_
                d_574_v158_: D15
                d_574_v158_ = D15_DC36(d_375_v3_, d_378_v6_)
                d_574_v158_ = d_574_v158_
        d_575_v159_: _dafny.Map
        d_575_v159_ = _dafny.Map({(0) - (d_378_v6_): d_378_v6_})
        d_576_v160_: _dafny.Map
        d_576_v160_ = _dafny.Map({d_575_v159_: default__.fm1(d_386_globalState_)})
        d_577_v161_: _dafny.MultiSet
        d_577_v161_ = _dafny.MultiSet([d_376_v4_])
        d_578_v162_: _dafny.Array
        nw103_ = _dafny.Array(None, 5)
        nw103_[int(0)] = d_373_v1_
        nw103_[int(1)] = d_373_v1_
        nw103_[int(2)] = d_373_v1_
        nw103_[int(3)] = d_373_v1_
        nw103_[int(4)] = d_373_v1_
        d_578_v162_ = nw103_
        d_579_v163_: C5
        nw104_ = C5()
        nw104_.ctor__(d_576_v160_, (d_577_v161_).set(d_376_v4_, default__.abs(len(d_377_v5_))), (D10_DC21(d_381_v9_)).cf39, d_378_v6_, d_578_v162_)
        d_579_v163_ = nw104_
        hi1_ = d_378_v6_
        for d_580_i21_ in range(d_378_v6_, hi1_):
            index43_ = default__.safeIndex(416, (d_372_v0_).length(0))
            (d_372_v0_)[index43_] = -775
            index44_ = default__.safeIndex(416, (d_372_v0_).length(0))
            (d_372_v0_)[index44_] = d_580_i21_
            d_581_v164_: int
            d_582_v165_: int
            d_583_v166_: int
            out20_: int
            out21_: int
            out22_: int
            out20_, out21_, out22_ = (d_579_v163_).m1(d_386_globalState_)
            d_581_v164_ = out20_
            d_582_v165_ = out21_
            d_583_v166_ = out22_
            d_584_v167_: _dafny.Seq
            d_584_v167_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_376_v4_, d_376_v4_])])
            d_585_v168_: D19
            d_585_v168_ = D19_DC49(d_376_v4_, d_376_v4_, d_580_i21_)
            d_586_v169_: _dafny.Seq
            d_586_v169_ = _dafny.SeqWithoutIsStrInference([d_578_v162_])
            pat_let_tv5_ = d_376_v4_
            d_587_v170_: C6
            nw105_ = C6()
            def iife43_(_pat_let10_0):
                def iife44_(d_588_dt__update__tmp_h2_):
                    def iife45_(_pat_let11_0):
                        def iife46_(d_589_dt__update_hcf89_h0_):
                            return D19_DC49(d_589_dt__update_hcf89_h0_, (d_588_dt__update__tmp_h2_).cf90, (d_588_dt__update__tmp_h2_).cf91)
                        return iife46_(_pat_let11_0)
                    return iife45_(pat_let_tv5_)
                return iife44_(_pat_let10_0)
            nw105_.ctor__(d_376_v4_, (615) + (d_378_v6_), (d_584_v167_)[default__.safeIndex(d_580_i21_, len(d_584_v167_))], _dafny.SeqWithoutIsStrInference([(d_372_v0_)[default__.safeIndex(416, (d_372_v0_).length(0))], (iife43_(d_585_v168_)).cf91]), d_378_v6_, (d_586_v169_)[default__.safeIndex(338, len(d_586_v169_))], d_577_v161_, d_582_v165_)
            d_587_v170_ = nw105_
            d_587_v170_ = d_587_v170_
            d_590_v171_: _dafny.MultiSet
            d_590_v171_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugxwp")))])
            (d_386_globalState_).f26 = (((d_372_v0_)[default__.safeIndex(416, (d_372_v0_).length(0))]) * (len(d_381_v9_))) >= ((len(d_381_v9_)) + ((d_590_v171_).cardinality))
        d_591_v172_: D4
        d_591_v172_ = D4_DC9(d_378_v6_, False)
        d_592_v173_: _dafny.Map
        d_592_v173_ = _dafny.Map({d_376_v4_: (d_591_v172_).cf17})
        d_592_v173_ = (d_592_v173_).set(d_376_v4_, d_376_v4_)
        d_593_v174_: C3
        nw106_ = C3()
        nw106_.ctor__(923, d_578_v162_)
        d_593_v174_ = nw106_
        d_594_v175_: _dafny.Map
        d_594_v175_ = _dafny.Map({len(d_373_v1_): d_593_v174_})
        d_595_v176_: _dafny.Map
        d_595_v176_ = _dafny.Map({d_594_v175_: d_378_v6_})
        d_595_v176_ = (d_595_v176_).set(_dafny.Map({(0) - (d_378_v6_): d_593_v174_}), d_378_v6_)
        d_596_v177_: _dafny.MultiSet
        d_596_v177_ = _dafny.MultiSet([d_378_v6_, d_378_v6_])
        if ((d_579_v163_).fm4(_dafny.Set({d_373_v1_, d_373_v1_, _dafny.SeqWithoutIsStrInference([default__.fm26(d_375_v3_, (d_596_v177_).cardinality, d_378_v6_, d_386_globalState_)])}), d_376_v4_, (_dafny.SeqWithoutIsStrInference([not(d_376_v4_)])).set(default__.safeIndex(len(_dafny.Set({d_376_v4_})), len(_dafny.SeqWithoutIsStrInference([not(d_376_v4_)]))), d_376_v4_), d_386_globalState_)) < (d_378_v6_):
            (d_386_globalState_).f26 = True
            d_373_v1_ = (d_373_v1_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_597_i22_ in range(default__.abs(684))]))
            (d_386_globalState_).f0 = 522
            index45_ = default__.safeIndex(206, (d_383_v11_).length(0))
            (d_383_v11_)[index45_] = not(d_376_v4_)
            index46_ = default__.safeIndex(206, (d_383_v11_).length(0))
            (d_383_v11_)[index46_] = d_376_v4_
            d_596_v177_ = d_596_v177_
        elif True:
            d_598_v178_: _dafny.Seq
            d_598_v178_ = _dafny.SeqWithoutIsStrInference([d_373_v1_])
            d_373_v1_ = (d_598_v178_)[default__.safeIndex(default__.safeModuloInt(d_378_v6_, d_378_v6_), len(d_598_v178_))]
            d_599_v179_: _dafny.Array
            d_600_v180_: bool
            d_601_v181_: bool
            out23_: _dafny.Array
            out24_: bool
            out25_: bool
            out23_, out24_, out25_ = (d_579_v163_).m9(d_386_globalState_)
            d_599_v179_ = out23_
            d_600_v180_ = out24_
            d_601_v181_ = out25_
            index47_ = default__.safeIndex(726, (d_372_v0_).length(0))
            (d_372_v0_)[index47_] = len(d_373_v1_)
            index48_ = default__.safeIndex(726, (d_372_v0_).length(0))
            (d_372_v0_)[index48_] = d_378_v6_
            d_375_v3_ = (d_375_v3_ if d_601_v181_ else d_375_v3_)
            d_602_v182_: int
            d_603_v183_: int
            d_604_v184_: int
            out26_: int
            out27_: int
            out28_: int
            out26_, out27_, out28_ = (d_579_v163_).m1(d_386_globalState_)
            d_602_v182_ = out26_
            d_603_v183_ = out27_
            d_604_v184_ = out28_
        d_575_v159_ = _dafny.Map({default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_378_v6_])), (0) - (d_378_v6_)): d_378_v6_})
        index49_ = default__.safeIndex(515, (d_372_v0_).length(0))
        (d_372_v0_)[index49_] = d_378_v6_
        d_605_v185_: _dafny.Seq
        d_605_v185_ = _dafny.SeqWithoutIsStrInference([d_376_v4_])
        index50_ = default__.safeIndex(515, (d_372_v0_).length(0))
        (d_372_v0_)[index50_] = (0) - (((d_579_v163_).fm4(d_385_v13_, d_376_v4_, d_605_v185_, d_386_globalState_)) + (default__.safeDivisionInt((d_381_v9_)[default__.safeIndex(d_378_v6_, len(d_381_v9_))], d_378_v6_)))
        rhs51_ = (d_376_v4_ if d_376_v4_ else d_376_v4_)
        rhs52_ = d_378_v6_
        rhs53_ = (d_376_v4_ if (d_376_v4_ if d_376_v4_ else d_376_v4_) else d_376_v4_)
        rhs54_ = 451
        lhs32_ = d_386_globalState_
        lhs33_ = d_386_globalState_
        lhs32_.f26 = rhs51_
        lhs33_.f18 = rhs52_
        d_376_v4_ = rhs53_
        d_378_v6_ = rhs54_
        d_383_v11_ = d_383_v11_
        (d_386_globalState_).f26 = (len(d_592_v173_)) != ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))])
        hi2_ = (d_378_v6_) - ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))])
        for d_606_i23_ in range((d_593_v174_).fm45(d_386_globalState_), hi2_):
            d_607_v186_: D16
            d_607_v186_ = D16_DC40(d_376_v4_, (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], len((d_380_v8_).set(d_375_v3_, (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))])))
            d_608_v187_: _dafny.Map
            d_608_v187_ = _dafny.Map({419: d_375_v3_})
            d_609_v188_: D4
            d_609_v188_ = D4_DC8(((d_608_v187_)[(d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]] if ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]) in (d_608_v187_) else d_375_v3_), len(d_373_v1_), d_376_v4_)
            pat_let_tv6_ = d_375_v3_
            d_610_v189_: _dafny.Map
            def iife47_(_pat_let12_0):
                def iife48_(d_611_dt__update__tmp_h3_):
                    def iife49_(_pat_let13_0):
                        def iife50_(d_612_dt__update_hcf13_h0_):
                            return D4_DC8(d_612_dt__update_hcf13_h0_, (d_611_dt__update__tmp_h3_).cf14, (d_611_dt__update__tmp_h3_).cf15)
                        return iife50_(_pat_let13_0)
                    return iife49_(pat_let_tv6_)
                return iife48_(_pat_let12_0)
            d_610_v189_ = _dafny.Map({d_376_v4_: iife47_(d_609_v188_)})
            d_613_v190_: _dafny.MultiSet
            d_613_v190_ = _dafny.MultiSet([(_dafny.Map({d_376_v4_: d_609_v188_})).set(d_376_v4_, d_609_v188_), d_610_v189_])
            d_614_v191_: D2
            d_614_v191_ = D2_DC3(_dafny.Map({_dafny.CodePoint('c'): d_376_v4_}))
            d_615_v192_: D17
            d_615_v192_ = D17_DC43(d_578_v162_)
            d_616_v193_: D17
            d_616_v193_ = D17_DC45(d_615_v192_)
            d_617_v194_: _dafny.Seq
            d_617_v194_ = _dafny.SeqWithoutIsStrInference([d_616_v193_, D17_DC45(d_615_v192_), d_616_v193_])
            d_618_v195_: _dafny.Map
            d_618_v195_ = _dafny.Map({d_617_v194_: d_376_v4_})
            index51_ = default__.safeIndex(515, (d_372_v0_).length(0))
            rhs55_ = not(not(((d_579_v163_).fm4(d_385_v13_, False, d_605_v185_, d_386_globalState_)) > (d_606_i23_)))
            rhs56_ = (default__.fm55(d_378_v6_, d_614_v191_, d_376_v4_, len(_dafny.Set({d_376_v4_, d_376_v4_})), d_386_globalState_) if (not(True)) or (d_376_v4_) else D16_DC40(d_376_v4_, (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], d_606_i23_))
            rhs57_ = d_613_v190_
            rhs58_ = len(d_618_v195_)
            rhs59_ = d_376_v4_
            lhs34_ = d_386_globalState_
            lhs35_ = d_372_v0_
            lhs36_ = default__.safeIndex(515, (d_372_v0_).length(0))
            lhs34_.f26 = rhs55_
            d_607_v186_ = rhs56_
            d_613_v190_ = rhs57_
            lhs35_[lhs36_] = rhs58_
            d_376_v4_ = rhs59_
            (d_386_globalState_).f21 = d_606_i23_
            d_619_v196_: D2
            d_619_v196_ = D2_DC4(d_378_v6_, len(d_373_v1_), (d_593_v174_).fm45(d_386_globalState_), d_375_v3_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_620_i24_ in range(default__.abs(440))]))
            d_375_v3_ = (d_375_v3_ if ((d_596_v177_).set(d_606_i23_, default__.abs(d_606_i23_))) == (d_596_v177_) else (d_619_v196_).cf8)
            if d_376_v4_:
                d_621_v197_: _dafny.Array
                nw107_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_621_v197_ = nw107_
                d_621_v197_ = d_621_v197_
                d_622_v198_: D15
                d_622_v198_ = D15_DC37(d_376_v4_, (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], d_372_v0_, True)
                d_623_v199_: _dafny.Map
                d_623_v199_ = _dafny.Map({d_606_i23_: False})
                d_624_v200_: _dafny.Map
                d_624_v200_ = _dafny.Map({not(d_376_v4_): d_378_v6_})
                d_625_v201_: T2
                nw108_ = C2()
                nw108_.ctor__(d_577_v161_, len(d_624_v200_))
                d_625_v201_ = nw108_
                d_626_v202_: _dafny.Array
                nw109_ = _dafny.Array(None, 13)
                nw109_[int(0)] = d_622_v198_
                nw109_[int(1)] = d_622_v198_
                nw109_[int(2)] = d_622_v198_
                nw109_[int(3)] = D15_DC37(d_376_v4_, (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], d_372_v0_, d_376_v4_)
                nw109_[int(4)] = d_622_v198_
                nw109_[int(5)] = D15_DC37(d_376_v4_, len(d_623_v199_), d_372_v0_, d_376_v4_)
                nw109_[int(6)] = d_622_v198_
                nw109_[int(7)] = d_622_v198_
                nw109_[int(8)] = d_622_v198_
                nw109_[int(9)] = (d_622_v198_ if False else d_622_v198_)
                nw109_[int(10)] = d_622_v198_
                nw109_[int(11)] = D15_DC37(not((d_579_v163_).fm3(d_577_v161_, d_386_globalState_)), len(_dafny.SeqWithoutIsStrInference([d_625_v201_])), d_372_v0_, d_376_v4_)
                nw109_[int(12)] = d_622_v198_
                d_626_v202_ = nw109_
                index52_ = default__.safeIndex(1, (d_578_v162_).length(0))
                (d_578_v162_)[index52_] = d_373_v1_
                d_627_v203_: _dafny.Seq
                d_627_v203_ = _dafny.SeqWithoutIsStrInference([d_377_v5_])
                index53_ = default__.safeIndex(1, (d_578_v162_).length(0))
                rhs60_ = default__.safeModuloInt((d_378_v6_) + (d_606_i23_), d_606_i23_)
                rhs61_ = (default__.safeDivisionInt((d_625_v201_).f32, (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))])) >= ((d_381_v9_)[default__.safeIndex(len(d_627_v203_), len(d_381_v9_))])
                rhs62_ = d_626_v202_
                rhs63_ = d_373_v1_
                rhs64_ = d_378_v6_
                lhs37_ = d_386_globalState_
                lhs38_ = d_386_globalState_
                lhs39_ = d_578_v162_
                lhs40_ = default__.safeIndex(1, (d_578_v162_).length(0))
                lhs41_ = d_386_globalState_
                lhs37_.f18 = rhs60_
                lhs38_.f26 = rhs61_
                d_626_v202_ = rhs62_
                lhs39_[lhs40_] = rhs63_
                lhs41_.f3 = rhs64_
                d_628_v204_: _dafny.Set
                d_628_v204_ = _dafny.Set({d_375_v3_})
                d_629_v205_: _dafny.Map
                d_629_v205_ = _dafny.Map({d_376_v4_: d_628_v204_})
                d_630_v206_: _dafny.Map
                d_630_v206_ = _dafny.Map({d_592_v173_: ((d_629_v205_)[d_376_v4_] if (d_376_v4_) in (d_629_v205_) else d_628_v204_)})
                d_376_v4_ = (((d_623_v199_)[len(d_630_v206_)] if (len(d_630_v206_)) in (d_623_v199_) else d_376_v4_) if (d_605_v185_) == (d_605_v185_) else d_376_v4_)
                d_378_v6_ = 897
                (d_386_globalState_).f0 = d_378_v6_
            elif True:
                d_631_v207_: _dafny.Array
                nw110_ = _dafny.Array(None, 16)
                nw110_[int(0)] = d_383_v11_
                nw110_[int(1)] = d_383_v11_
                nw110_[int(2)] = d_383_v11_
                nw110_[int(3)] = d_383_v11_
                nw110_[int(4)] = d_383_v11_
                nw110_[int(5)] = d_383_v11_
                nw110_[int(6)] = d_383_v11_
                nw110_[int(7)] = d_383_v11_
                nw110_[int(8)] = d_383_v11_
                nw110_[int(9)] = d_383_v11_
                nw110_[int(10)] = d_383_v11_
                nw110_[int(11)] = d_383_v11_
                nw110_[int(12)] = d_383_v11_
                nw110_[int(13)] = d_383_v11_
                nw110_[int(14)] = d_383_v11_
                nw110_[int(15)] = d_383_v11_
                d_631_v207_ = nw110_
                d_632_v209_: _dafny.Map
                d_632_v209_ = _dafny.Map({d_375_v3_: d_376_v4_})
                d_633_v210_: D2
                def iife51_():
                    coll23_ = _dafny.Map()
                    compr_23_: str
                    for compr_23_ in (d_632_v209_).keys.Elements:
                        d_634_v208_: str = compr_23_
                        if (d_634_v208_) in (d_632_v209_):
                            coll23_[d_634_v208_] = d_376_v4_
                    return _dafny.Map(coll23_)
                d_633_v210_ = D2_DC3(iife51_()
)
                d_635_v211_: D2
                d_635_v211_ = D2_DC5(d_633_v210_)
                d_636_v212_: D2
                d_636_v212_ = D2_DC5(d_635_v211_)
                (d_386_globalState_).f26 = ((d_592_v173_)[default__.fm0(d_606_i23_, d_386_globalState_)] if (default__.fm0(d_606_i23_, d_386_globalState_)) in (d_592_v173_) else ((D16_DC42((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], d_631_v207_, (d_593_v174_).fm46(d_378_v6_, d_375_v3_, d_386_globalState_), d_636_v212_)).cf80 if d_376_v4_ else d_376_v4_))
                d_637_v213_: bool
                d_638_v214_: int
                out29_: bool
                out30_: int
                out29_, out30_ = default__.m0(((0) - ((d_593_v174_).fm45(d_386_globalState_))) > ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]), len(d_373_v1_), default__.safeDivisionInt(d_606_i23_, d_378_v6_), (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], d_386_globalState_)
                d_637_v213_ = out29_
                d_638_v214_ = out30_
                d_639_v215_: _dafny.Array
                def lambda24_(d_640_v214_, d_641_v4_):
                    def lambda25_(d_642_i25_):
                        return _dafny.Map({d_640_v214_: not(d_641_v4_)})

                    return lambda25_

                init13_ = lambda24_(d_638_v214_, d_376_v4_)
                nw111_ = _dafny.Array(None, 23)
                for i0_13_ in range(nw111_.length(0)):
                    nw111_[i0_13_] = init13_(i0_13_)
                d_639_v215_ = nw111_
                d_643_v216_: _dafny.Seq
                d_643_v216_ = _dafny.SeqWithoutIsStrInference([d_639_v215_, d_639_v215_])
                d_644_v217_: _dafny.Map
                d_644_v217_ = _dafny.Map({len((((d_373_v1_).set(default__.safeIndex(d_638_v214_, len(d_373_v1_)), d_375_v3_) if False else d_373_v1_)).set(default__.safeIndex((0) - (d_606_i23_), len(((d_373_v1_).set(default__.safeIndex(d_638_v214_, len(d_373_v1_)), d_375_v3_) if False else d_373_v1_))), d_375_v3_)): (d_643_v216_)[default__.safeIndex(d_638_v214_, len(d_643_v216_))]})
                d_644_v217_ = (d_644_v217_).set((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], d_639_v215_)
                d_645_v218_: _dafny.Seq
                d_645_v218_ = _dafny.SeqWithoutIsStrInference([d_578_v162_, d_578_v162_, d_578_v162_])
                d_646_v219_: _dafny.Map
                d_646_v219_ = _dafny.Map({(d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]: (d_645_v218_)[default__.safeIndex(283, len(d_645_v218_))]})
                d_647_v220_: C5
                nw112_ = C5()
                nw112_.ctor__((d_579_v163_).f38, d_577_v161_, d_381_v9_, (d_374_v2_).cardinality, (((d_646_v219_)[len(d_605_v185_)] if (len(d_605_v185_)) in (d_646_v219_) else d_578_v162_) if d_376_v4_ else d_578_v162_))
                d_647_v220_ = nw112_
                d_648_v221_: _dafny.Seq
                d_648_v221_ = _dafny.SeqWithoutIsStrInference([(d_373_v1_ if d_376_v4_ else d_373_v1_), d_373_v1_, d_373_v1_])
                d_648_v221_ = ((d_648_v221_).set(default__.safeIndex(938, len(d_648_v221_)), _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_649_i26_ in range(default__.abs(582))]))).set(default__.safeIndex(d_638_v214_, len((d_648_v221_).set(default__.safeIndex(938, len(d_648_v221_)), _dafny.SeqWithoutIsStrInference([d_375_v3_ for d_650_i26_ in range(default__.abs(582))])))), default__.fm14(not(d_637_v213_), d_376_v4_, d_376_v4_, d_378_v6_, d_386_globalState_))
        if d_376_v4_:
            rhs65_ = ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))] if False else default__.safeDivisionInt((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], len(d_592_v173_)))
            rhs66_ = (d_381_v9_) + (d_381_v9_)
            lhs42_ = d_386_globalState_
            lhs42_.f21 = rhs65_
            d_381_v9_ = rhs66_
            d_651_v222_: D9
            d_651_v222_ = D9_DC20(default__.fm1(d_386_globalState_), d_376_v4_, ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]) < ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]), d_376_v4_, _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_375_v3_, default__.fm12((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], d_376_v4_, d_386_globalState_)])): 870}))
            source13_ = d_651_v222_
            if source13_.is_DC18:
                d_652___mcc_h8_ = source13_.cf32
                d_653_cf32_ = d_652___mcc_h8_
                d_378_v6_ = (d_593_v174_).fm45(d_386_globalState_)
                d_654_v223_: _dafny.Seq
                d_654_v223_ = _dafny.SeqWithoutIsStrInference([((d_577_v161_).set(d_653_cf32_, default__.abs((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]))) | (d_577_v161_)])
                d_654_v223_ = (d_654_v223_).set(default__.safeIndex(default__.safeModuloInt(len(d_377_v5_), (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]), len(d_654_v223_)), default__.fm32(d_386_globalState_))
                d_381_v9_ = default__.fm9(d_653_cf32_, d_386_globalState_)
                d_655_v224_: _dafny.Map
                d_655_v224_ = _dafny.Map({d_653_cf32_: d_577_v161_})
                d_656_v225_: D11
                d_656_v225_ = D11_DC23(d_445_v62_)
                d_657_v226_: _dafny.Map
                d_657_v226_ = _dafny.Map({(d_653_cf32_) in (d_655_v224_): d_656_v225_})
                d_657_v226_ = (d_657_v226_).set((default__.fm1(d_386_globalState_)) == ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]), d_656_v225_)
            elif source13_.is_DC19:
                d_658___mcc_h9_ = source13_.cf33
                d_659_cf33_ = d_658___mcc_h9_
                d_660_v227_: _dafny.Map
                d_660_v227_ = _dafny.Map({d_605_v185_: d_578_v162_})
                d_661_v228_: D7
                d_661_v228_ = D7_DC13(d_605_v185_)
                d_662_v229_: _dafny.MultiSet
                d_662_v229_ = _dafny.MultiSet([d_575_v159_])
                d_663_v230_: C6
                nw113_ = C6()
                nw113_.ctor__(d_376_v4_, (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], d_577_v161_, (default__.fm9(d_659_cf33_, d_386_globalState_)) + (_dafny.SeqWithoutIsStrInference([(d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]])), (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], ((d_660_v227_)[(d_661_v228_).cf23] if ((d_661_v228_).cf23) in (d_660_v227_) else d_578_v162_), d_577_v161_, ((d_662_v229_)[d_575_v159_] if (d_575_v159_) in (d_662_v229_) else (0) - (len(d_377_v5_))))
                d_663_v230_ = nw113_
                d_663_v230_ = d_663_v230_
                index54_ = default__.safeIndex(216, (d_578_v162_).length(0))
                (d_578_v162_)[index54_] = (d_373_v1_) + (d_373_v1_)
                index55_ = default__.safeIndex(216, (d_578_v162_).length(0))
                (d_578_v162_)[index55_] = d_373_v1_
                d_664_v231_: D17
                d_664_v231_ = D17_DC45(D17_DC44(len(d_373_v1_)))
                d_665_v232_: D17
                d_665_v232_ = D17_DC45(d_664_v231_)
                d_666_v233_: D12
                d_666_v233_ = D12_DC27(True, d_596_v177_, d_378_v6_, d_372_v0_)
                d_667_v234_: _dafny.Map
                d_667_v234_ = _dafny.Map({d_665_v232_: ((d_666_v233_).cf50 if (d_663_v230_).f33 else d_372_v0_)})
                d_667_v234_ = (d_667_v234_).set(d_665_v232_, d_372_v0_)
                d_668_v235_: D17
                d_668_v235_ = D17_DC44(d_378_v6_)
                d_669_v236_: _dafny.Array
                nw114_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                d_669_v236_ = nw114_
                d_670_v237_: C4
                nw115_ = C4()
                nw115_.ctor__((d_668_v235_).cf83, d_669_v236_)
                d_670_v237_ = nw115_
                (d_386_globalState_).f26 = ((len(_dafny.Map({d_670_v237_: (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]}))) - (d_378_v6_)) > (744)
            elif source13_.is_DC20:
                d_671___mcc_h10_ = source13_.cf34
                d_672___mcc_h11_ = source13_.cf35
                d_673___mcc_h12_ = source13_.cf36
                d_674___mcc_h13_ = source13_.cf37
                d_675___mcc_h14_ = source13_.cf38
                d_676_cf38_ = d_675___mcc_h14_
                d_677_cf37_ = d_674___mcc_h13_
                d_678_cf36_ = d_673___mcc_h12_
                d_679_cf35_ = d_672___mcc_h11_
                d_680_cf34_ = d_671___mcc_h10_
                (d_386_globalState_).f26 = False
                d_376_v4_ = (d_379_v7_).ispropersubset(d_379_v7_)
                d_681_v238_: _dafny.Map
                d_681_v238_ = _dafny.Map({d_383_v11_: d_378_v6_})
                d_682_v239_: _dafny.Map
                d_682_v239_ = _dafny.Map({not(d_376_v4_): d_373_v1_})
                d_683_v240_: _dafny.Seq
                d_683_v240_ = _dafny.SeqWithoutIsStrInference([d_681_v238_, _dafny.Map({d_383_v11_: len(((d_682_v239_)[d_376_v4_] if (d_376_v4_) in (d_682_v239_) else (d_373_v1_).set(default__.safeIndex(d_378_v6_, len(d_373_v1_)), d_375_v3_)))})])
                d_684_v241_: _dafny.Map
                d_684_v241_ = _dafny.Map({d_378_v6_: d_681_v238_})
                d_685_v242_: _dafny.Set
                d_685_v242_ = _dafny.Set({d_375_v3_, d_375_v3_})
                (d_386_globalState_).f26 = ((d_683_v240_)[default__.safeIndex(d_378_v6_, len(d_683_v240_))]) == (((d_684_v241_)[len(d_685_v242_)] if (len(d_685_v242_)) in (d_684_v241_) else _dafny.Map({d_383_v11_: ((d_577_v161_)[d_677_cf37_] if (d_677_cf37_) in (d_577_v161_) else (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))])})))
                (d_386_globalState_).f18 = ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]) + ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))])
            elif True:
                d_686___mcc_h15_ = source13_.cf31
                d_687_cf31_ = d_686___mcc_h15_
                (d_579_v163_).m10(False, (d_381_v9_) + (d_381_v9_), d_386_globalState_)
                d_688_v243_: _dafny.Set
                d_688_v243_ = _dafny.Set({d_372_v0_, d_372_v0_})
                index56_ = default__.safeIndex(808, (d_383_v11_).length(0))
                (d_383_v11_)[index56_] = (d_688_v243_).isdisjoint(d_688_v243_)
                index57_ = default__.safeIndex(808, (d_383_v11_).length(0))
                rhs67_ = d_376_v4_
                rhs68_ = d_376_v4_
                lhs43_ = d_383_v11_
                lhs44_ = default__.safeIndex(808, (d_383_v11_).length(0))
                lhs45_ = d_386_globalState_
                lhs43_[lhs44_] = rhs67_
                lhs45_.f26 = rhs68_
                d_372_v0_ = d_372_v0_
                d_689_v244_: _dafny.Array
                nw116_ = _dafny.Array(None, 6)
                nw116_[int(0)] = True
                nw116_[int(1)] = not((d_383_v11_)[default__.safeIndex(808, (d_383_v11_).length(0))])
                nw116_[int(2)] = False
                nw116_[int(3)] = (d_383_v11_)[default__.safeIndex(808, (d_383_v11_).length(0))]
                nw116_[int(4)] = d_376_v4_
                nw116_[int(5)] = (d_579_v163_).fm3(d_577_v161_, d_386_globalState_)
                d_689_v244_ = nw116_
                d_690_v245_: C4
                nw117_ = C4()
                nw117_.ctor__(231, d_578_v162_)
                d_690_v245_ = nw117_
                d_691_v246_: _dafny.Map
                d_691_v246_ = _dafny.Map({d_689_v244_: d_690_v245_})
                d_691_v246_ = (d_691_v246_).set(d_689_v244_, d_690_v245_)
            d_692_v247_: _dafny.Set
            d_692_v247_ = _dafny.Set({d_377_v5_, d_377_v5_})
            d_693_v249_: _dafny.MultiSet
            d_693_v249_ = _dafny.MultiSet([_dafny.Set({d_376_v4_}), d_377_v5_, d_377_v5_])
            index58_ = default__.safeIndex(144, (d_383_v11_).length(0))
            def iife52_():
                coll24_ = _dafny.Set()
                compr_24_: _dafny.Set
                for compr_24_ in (d_692_v247_).Elements:
                    d_694_v248_: _dafny.Set = compr_24_
                    if (d_694_v248_) in (d_692_v247_):
                        coll24_ = coll24_.union(_dafny.Set([d_694_v248_]))
                return _dafny.Set(coll24_)
            def iife53_():
                coll25_ = _dafny.Set()
                compr_25_: _dafny.Set
                for compr_25_ in (d_693_v249_).Elements:
                    d_695_v250_: _dafny.Set = compr_25_
                    if (d_695_v250_) in (d_693_v249_):
                        coll25_ = coll25_.union(_dafny.Set([d_695_v250_]))
                return _dafny.Set(coll25_)
            (d_383_v11_)[index58_] = (iife52_()
            ).isdisjoint(iife53_()
            )
            d_696_v251_: _dafny.Set
            d_696_v251_ = _dafny.Set({d_593_v174_, d_593_v174_})
            d_697_v252_: _dafny.Seq
            d_697_v252_ = _dafny.SeqWithoutIsStrInference([d_696_v251_])
            index59_ = default__.safeIndex(144, (d_383_v11_).length(0))
            (d_383_v11_)[index59_] = not((d_697_v252_) != (d_697_v252_))
            (d_386_globalState_).f26 = (d_383_v11_)[default__.safeIndex(144, (d_383_v11_).length(0))]
            d_698_v253_: _dafny.Array
            nw118_ = _dafny.Array(_dafny.Array(None, 0), 29)
            d_698_v253_ = nw118_
            d_699_v254_: _dafny.Array
            def lambda26_(d_700_v0_, d_701_v4_):
                def lambda27_(d_702_i27_):
                    return _dafny.Map({(0) - ((d_700_v0_)[default__.safeIndex(515, (d_700_v0_).length(0))]): d_701_v4_})

                return lambda27_

            init14_ = lambda26_(d_372_v0_, d_376_v4_)
            nw119_ = _dafny.Array(None, 23)
            for i0_14_ in range(nw119_.length(0)):
                nw119_[i0_14_] = init14_(i0_14_)
            d_699_v254_ = nw119_
            index60_ = default__.safeIndex(111, (d_698_v253_).length(0))
            (d_698_v253_)[index60_] = d_699_v254_
            index61_ = default__.safeIndex(111, (d_698_v253_).length(0))
            (d_698_v253_)[index61_] = d_699_v254_
        elif True:
            d_372_v0_ = d_372_v0_
            d_703_v255_: _dafny.Set
            d_703_v255_ = _dafny.Set({default__.fm26(d_375_v3_, len(d_381_v9_), (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], d_386_globalState_), d_375_v3_})
            d_704_v256_: T2
            nw120_ = C2()
            nw120_.ctor__(_dafny.MultiSet([False]), (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))])
            d_704_v256_ = nw120_
            d_705_v257_: _dafny.MultiSet
            d_705_v257_ = _dafny.MultiSet([d_704_v256_, d_704_v256_])
            d_706_v258_: T0
            nw121_ = C4()
            nw121_.ctor__(d_378_v6_, d_578_v162_)
            d_706_v258_ = nw121_
            nw122_ = _dafny.Array(None, 19)
            nw122_[int(0)] = default__.safeModuloInt(len(d_703_v255_), default__.fm1(d_386_globalState_))
            nw122_[int(1)] = d_378_v6_
            nw122_[int(2)] = d_378_v6_
            nw122_[int(3)] = ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]) - (d_378_v6_)
            nw122_[int(4)] = (len((d_381_v9_).set(default__.safeIndex((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], len(d_381_v9_)), (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]))) + (d_378_v6_)
            nw122_[int(5)] = 802
            nw122_[int(6)] = ((d_705_v257_).cardinality if d_376_v4_ else d_378_v6_)
            nw122_[int(7)] = ((d_579_v163_).fm4(_dafny.Set({d_373_v1_, d_373_v1_}), d_376_v4_, d_605_v185_, d_386_globalState_)) - (-311)
            nw122_[int(8)] = d_378_v6_
            nw122_[int(9)] = d_378_v6_
            nw122_[int(10)] = default__.safeModuloInt((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], len(d_605_v185_))
            nw122_[int(11)] = 956
            nw122_[int(12)] = ((d_704_v256_).f32 if d_376_v4_ else (d_381_v9_)[default__.safeIndex(len(d_605_v185_), len(d_381_v9_))])
            nw122_[int(13)] = ((d_577_v161_).cardinality) * ((d_704_v256_).f32)
            nw122_[int(14)] = (0) - (default__.safeDivisionInt(d_378_v6_, (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]))
            nw122_[int(15)] = (d_704_v256_).f32
            def iife54_():
                coll26_ = _dafny.Set()
                compr_26_: str
                for compr_26_ in (d_373_v1_).Elements:
                    d_707_v259_: str = compr_26_
                    if (d_707_v259_) in (d_373_v1_):
                        coll26_ = coll26_.union(_dafny.Set([d_707_v259_]))
                return _dafny.Set(coll26_)
            nw122_[int(16)] = len(_dafny.Map({d_706_v258_: iife54_()
            }))
            nw122_[int(17)] = (d_579_v163_).fm4(d_385_v13_, d_376_v4_, d_605_v185_, d_386_globalState_)
            nw122_[int(18)] = (d_706_v258_).f27
            d_372_v0_ = nw122_
            index62_ = default__.safeIndex(940, (d_383_v11_).length(0))
            (d_383_v11_)[index62_] = (d_376_v4_) or (not(d_376_v4_))
            d_708_v260_: _dafny.Map
            d_708_v260_ = _dafny.Map({d_372_v0_: d_379_v7_})
            d_709_v263_: _dafny.Array
            nw123_ = _dafny.Array(None, 29)
            nw123_[int(0)] = d_708_v260_
            nw123_[int(1)] = d_708_v260_
            nw123_[int(2)] = d_708_v260_
            nw123_[int(3)] = d_708_v260_
            nw123_[int(4)] = _dafny.Map({d_372_v0_: d_379_v7_})
            nw123_[int(5)] = (d_708_v260_ if d_376_v4_ else d_708_v260_)
            nw123_[int(6)] = d_708_v260_
            nw123_[int(7)] = (d_708_v260_) | (d_708_v260_)
            nw123_[int(8)] = d_708_v260_
            nw123_[int(9)] = d_708_v260_
            nw123_[int(10)] = (d_708_v260_) | (d_708_v260_)
            nw123_[int(11)] = (d_708_v260_).set(d_372_v0_, d_379_v7_)
            nw123_[int(12)] = (d_708_v260_) | (d_708_v260_)
            nw123_[int(13)] = d_708_v260_
            def iife55_():
                def iife57_():
                    coll29_ = _dafny.Map()
                    compr_29_: int
                    for compr_29_ in (_dafny.SeqWithoutIsStrInference([(d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]])).Elements:
                        d_710_v261_: int = compr_29_
                        if (d_710_v261_) in (_dafny.SeqWithoutIsStrInference([(d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]])):
                            coll29_[(d_710_v261_) * ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))])] = d_376_v4_
                    return _dafny.Map(coll29_)
                coll27_ = _dafny.Set()
                def iife56_():
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in (_dafny.SeqWithoutIsStrInference([(d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]])).Elements:
                        d_710_v261_: int = compr_28_
                        if (d_710_v261_) in (_dafny.SeqWithoutIsStrInference([(d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]])):
                            coll28_[(d_710_v261_) * ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))])] = d_376_v4_
                    return _dafny.Map(coll28_)
                compr_27_: int
                for compr_27_ in (iife56_()
                ).keys.Elements:
                    d_711_v262_: int = compr_27_
                    if (d_711_v262_) in (iife57_()
                    ):
                        coll27_ = coll27_.union(_dafny.Set([(d_711_v262_) - (458)]))
                return _dafny.Set(coll27_)
            nw123_[int(14)] = _dafny.Map({d_372_v0_: iife55_()
            })
            nw123_[int(15)] = d_708_v260_
            nw123_[int(16)] = (d_708_v260_) | (d_708_v260_)
            nw123_[int(17)] = (d_708_v260_) | (_dafny.Map({d_372_v0_: d_379_v7_}))
            nw123_[int(18)] = d_708_v260_
            nw123_[int(19)] = d_708_v260_
            nw123_[int(20)] = d_708_v260_
            nw123_[int(21)] = d_708_v260_
            nw123_[int(22)] = d_708_v260_
            nw123_[int(23)] = _dafny.Map({d_372_v0_: d_379_v7_})
            nw123_[int(24)] = (d_708_v260_) | (d_708_v260_)
            nw123_[int(25)] = _dafny.Map({d_372_v0_: d_379_v7_})
            nw123_[int(26)] = (d_708_v260_) | (d_708_v260_)
            nw123_[int(27)] = d_708_v260_
            nw123_[int(28)] = _dafny.Map({d_372_v0_: d_379_v7_})
            d_709_v263_ = nw123_
            index63_ = default__.safeIndex(784, (d_709_v263_).length(0))
            (d_709_v263_)[index63_] = d_708_v260_
            d_712_v264_: D2
            d_712_v264_ = D2_DC4((0) - (d_378_v6_), (d_706_v258_).f27, ((d_704_v256_).f32) * ((d_706_v258_).f27), d_375_v3_, d_373_v1_)
            d_713_v265_: D1
            d_713_v265_ = D1_DC2(d_375_v3_, (d_704_v256_).f32)
            index64_ = default__.safeIndex(940, (d_383_v11_).length(0))
            index65_ = default__.safeIndex(784, (d_709_v263_).length(0))
            rhs69_ = default__.safeModuloInt((0) - ((d_706_v258_).f27), default__.safeModuloInt((d_704_v256_).f32, (d_704_v256_).f32))
            rhs70_ = d_376_v4_
            rhs71_ = ((d_713_v265_).cf3) * ((d_706_v258_).f27)
            rhs72_ = (d_708_v260_ if (d_376_v4_) or (d_376_v4_) else d_708_v260_)
            rhs73_ = d_712_v264_
            lhs46_ = d_386_globalState_
            lhs47_ = d_383_v11_
            lhs48_ = default__.safeIndex(940, (d_383_v11_).length(0))
            lhs49_ = d_386_globalState_
            lhs50_ = d_709_v263_
            lhs51_ = default__.safeIndex(784, (d_709_v263_).length(0))
            lhs46_.f15 = rhs69_
            lhs47_[lhs48_] = rhs70_
            lhs49_.f15 = rhs71_
            lhs50_[lhs51_] = rhs72_
            d_712_v264_ = rhs73_
            d_714_v266_: C2
            nw124_ = C2()
            nw124_.ctor__((d_704_v256_).f31, -102)
            d_714_v266_ = nw124_
            d_714_v266_ = d_714_v266_
            (d_386_globalState_).f26 = (True if d_376_v4_ else ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]) not in (_dafny.SeqWithoutIsStrInference([(d_704_v256_).f32 for d_715_i28_ in range(default__.abs(274))])))
        d_716_i29_: int
        d_716_i29_ = 0
        with _dafny.label("2"):
            while not(d_376_v4_):
                with _dafny.c_label("2"):
                    if (d_716_i29_) >= (100):
                        raise _dafny.Break("2")
                    d_716_i29_ = (d_716_i29_) + (1)
                    d_717_v267_: D8
                    d_717_v267_ = D8_DC16((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], (d_378_v6_) * (d_378_v6_))
                    d_717_v267_ = d_717_v267_
                    (d_386_globalState_).f0 = (0) - ((len(d_373_v1_)) - (((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]) + (len(d_377_v5_))))
                    if d_376_v4_:
                        (d_386_globalState_).f21 = default__.safeModuloInt(((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]) + (len(d_381_v9_)), (d_596_v177_).cardinality)
                        d_718_v268_: _dafny.Seq
                        d_718_v268_ = _dafny.SeqWithoutIsStrInference([d_381_v9_, d_381_v9_])
                        d_575_v159_ = (d_575_v159_).set((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], len(d_718_v268_))
                        index66_ = default__.safeIndex(515, (d_372_v0_).length(0))
                        (d_372_v0_)[index66_] = (d_381_v9_)[default__.safeIndex((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], len(d_381_v9_))]
                        d_719_v269_: _dafny.Array
                        def lambda28_(d_720_v4_):
                            def lambda29_(d_721_i30_):
                                return _dafny.Map({_dafny.SeqWithoutIsStrInference([d_720_v4_, d_720_v4_]): d_720_v4_})

                            return lambda29_

                        init15_ = lambda28_(d_376_v4_)
                        nw125_ = _dafny.Array(None, 10)
                        for i0_15_ in range(nw125_.length(0)):
                            nw125_[i0_15_] = init15_(i0_15_)
                        d_719_v269_ = nw125_
                        d_722_v270_: _dafny.Map
                        d_722_v270_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_376_v4_, d_376_v4_]): d_376_v4_})
                        index67_ = default__.safeIndex(704, (d_719_v269_).length(0))
                        (d_719_v269_)[index67_] = d_722_v270_
                        index68_ = default__.safeIndex(704, (d_719_v269_).length(0))
                        (d_719_v269_)[index68_] = d_722_v270_
                        (d_579_v163_).m10(not (d_376_v4_) or (d_376_v4_), d_381_v9_, d_386_globalState_)
                    elif True:
                        d_376_v4_ = d_376_v4_
                        (d_386_globalState_).f26 = d_376_v4_
                        d_723_v271_: D12
                        d_723_v271_ = D12_DC27(not(d_376_v4_), (d_596_v177_).set(489, default__.abs(d_378_v6_)), d_378_v6_, d_372_v0_)
                        d_724_v272_: _dafny.Array
                        nw126_ = _dafny.Array(None, 3)
                        nw126_[int(0)] = d_723_v271_
                        nw126_[int(1)] = d_723_v271_
                        nw126_[int(2)] = d_723_v271_
                        d_724_v272_ = nw126_
                        index69_ = default__.safeIndex(740, (d_724_v272_).length(0))
                        (d_724_v272_)[index69_] = d_723_v271_
                        d_725_v274_: _dafny.Map
                        d_725_v274_ = _dafny.Map({d_375_v3_: d_376_v4_})
                        d_726_v275_: _dafny.Seq
                        def iife58_():
                            coll30_ = _dafny.Map()
                            compr_30_: str
                            for compr_30_ in (d_725_v274_).keys.Elements:
                                d_727_v273_: str = compr_30_
                                if (d_727_v273_) in (d_725_v274_):
                                    coll30_[d_727_v273_] = (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]
                            return _dafny.Map(coll30_)
                        d_726_v275_ = _dafny.SeqWithoutIsStrInference([iife58_()
                        ])
                        index70_ = default__.safeIndex(740, (d_724_v272_).length(0))
                        rhs74_ = d_723_v271_
                        rhs75_ = (166) * (((d_381_v9_)[default__.safeIndex(d_378_v6_, len(d_381_v9_))]) + ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]))
                        rhs76_ = default__.safeDivisionInt((0) - ((0) - ((-560) + (len(d_726_v275_)))), d_378_v6_)
                        rhs77_ = (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]
                        rhs78_ = not (True) or (not(d_376_v4_))
                        lhs52_ = d_724_v272_
                        lhs53_ = default__.safeIndex(740, (d_724_v272_).length(0))
                        lhs54_ = d_386_globalState_
                        lhs55_ = d_386_globalState_
                        lhs56_ = d_386_globalState_
                        lhs57_ = d_386_globalState_
                        lhs52_[lhs53_] = rhs74_
                        lhs54_.f0 = rhs75_
                        lhs55_.f0 = rhs76_
                        lhs56_.f21 = rhs77_
                        lhs57_.f26 = rhs78_
                        d_728_v276_: _dafny.Map
                        d_728_v276_ = _dafny.Map({d_378_v6_: d_376_v4_})
                        d_728_v276_ = d_728_v276_
                        rhs79_ = d_376_v4_
                        rhs80_ = d_376_v4_
                        rhs81_ = (139) + (d_378_v6_)
                        rhs82_ = default__.fm41((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))], d_378_v6_, d_376_v4_, d_386_globalState_)
                        rhs83_ = d_381_v9_
                        lhs58_ = d_386_globalState_
                        lhs59_ = d_386_globalState_
                        lhs60_ = d_386_globalState_
                        lhs58_.f26 = rhs79_
                        lhs59_.f26 = rhs80_
                        lhs60_.f0 = rhs81_
                        d_605_v185_ = rhs82_
                        d_381_v9_ = rhs83_
                    rhs84_ = ((d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))] if (D6_DC12(d_376_v4_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_376_v4_, False]) for d_729_i31_ in range(default__.abs(721))]), d_373_v1_)).cf20 else (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))])
                    rhs85_ = (d_372_v0_)[default__.safeIndex(515, (d_372_v0_).length(0))]
                    rhs86_ = d_592_v173_
                    lhs61_ = d_386_globalState_
                    lhs62_ = d_386_globalState_
                    lhs61_.f24 = rhs84_
                    lhs62_.f18 = rhs85_
                    d_592_v173_ = rhs86_
                    pass
            pass
        _dafny.print(_dafny.string_of((d_372_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_372_v0_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_373_v1_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_374_v2_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yn"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_375_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_376_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_377_v5_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_378_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_379_v7_) == (_dafny.Set({1, 491}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_380_v8_) == (_dafny.Map({_dafny.CodePoint('g'): 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_381_v9_) == (_dafny.SeqWithoutIsStrInference([491]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_382_v10_)[0]) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_382_v10_)[1]) == (_dafny.SeqWithoutIsStrInference([491]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_382_v10_)[2]) == (_dafny.SeqWithoutIsStrInference([491]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_383_v11_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_383_v11_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_383_v11_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_384_v12_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_385_v13_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yn"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_386_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_386_globalState_).f1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_386_globalState_).f2)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_386_globalState_).f2)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_386_globalState_).f2)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_386_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_.f8) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yn"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_386_globalState_).f10)[0]) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_386_globalState_).f10)[1]) == (_dafny.SeqWithoutIsStrInference([491]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_386_globalState_).f10)[2]) == (_dafny.SeqWithoutIsStrInference([491]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_386_globalState_.f13)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_386_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_386_globalState_).f17) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yn"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_386_globalState_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_386_globalState_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_386_globalState_).f22) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_386_globalState_).f23)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_386_globalState_).f23)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_386_globalState_).f23)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_386_globalState_.f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_386_globalState_).f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_386_globalState_.f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_444_v61_).cf45).cf42) == (_dafny.MultiSet([_dafny.Set({1, 491})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_445_v62_) == (_dafny.MultiSet([_dafny.Set({1, 491})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_446_v63_).cf42) == (_dafny.MultiSet([_dafny.Set({1, 491})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_575_v159_) == (_dafny.Map({0: -491}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_576_v160_) == (_dafny.Map({_dafny.Map({491: -491}): -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_577_v161_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_578_v162_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_578_v162_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_578_v162_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_578_v162_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_578_v162_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_579_v163_).f38) == (_dafny.Map({_dafny.Map({491: -491}): -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_591_v172_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_591_v172_).cf17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_592_v173_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_594_v175_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_595_v176_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_596_v177_) == (_dafny.MultiSet([-491, -491]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_605_v185_) == (_dafny.SeqWithoutIsStrInference([False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_716_i29_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({self.cf0.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC2(_dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
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
        return lambda: D2_DC4(int(0), int(0), int(0), _dafny.CodePoint('D'), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC4(D2, NamedTuple('DC4', [('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {self.cf9.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf4 == __o.cf4
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
        return lambda: False
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC6(D3, NamedTuple('DC6', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC8(_dafny.CodePoint('D'), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D4_DC7)

class D4_DC8(D4, NamedTuple('DC8', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC7(D4, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)

class D5_DC10(D5, NamedTuple('DC10', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC12(False, _dafny.Seq({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D6_DC11)

class D6_DC12(D6, NamedTuple('DC12', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {self.cf22.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC11(D6, NamedTuple('DC11', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC11({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC11) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC14(int(0), None, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D7_DC13)

class D7_DC14(D7, NamedTuple('DC14', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC13(D7, NamedTuple('DC13', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC13({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC13) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC16(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D8_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D8_DC15)

class D8_DC16(D8, NamedTuple('DC16', [('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC16({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC16) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC15(D8, NamedTuple('DC15', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC15({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC15) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC18(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D9_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D9_DC17)

class D9_DC18(D9, NamedTuple('DC18', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC18({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC18) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC19(D9, NamedTuple('DC19', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC20(D9, NamedTuple('DC20', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC17(D9, NamedTuple('DC17', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC17({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC17) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC22(int(0), _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)

class D10_DC22(D10, NamedTuple('DC22', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC21(D10, NamedTuple('DC21', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC24(int(0), D4.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D11_DC23)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)

class D11_DC24(D11, NamedTuple('DC24', [('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC23(D11, NamedTuple('DC23', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC23({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC23) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC25(D11, NamedTuple('DC25', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC27(False, _dafny.MultiSet({}), int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D12_DC27)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D12_DC26)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)

class D12_DC27(D12, NamedTuple('DC27', [('cf47', Any), ('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC27({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC27) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC28(D12, NamedTuple('DC28', [('cf51', Any), ('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC26(D12, NamedTuple('DC26', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC26({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC26) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC29(D12, NamedTuple('DC29', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC31(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D13_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D13_DC32)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D13_DC30)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)

class D13_DC31(D13, NamedTuple('DC31', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC31({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC31) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC32(D13, NamedTuple('DC32', [('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC32({_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC32) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC30(D13, NamedTuple('DC30', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC30({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC30) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC33(D13, NamedTuple('DC33', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D14_DC34)

class D14_DC34(D14, NamedTuple('DC34', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC34({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC34) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC36(_dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D15_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D15_DC35)

class D15_DC36(D15, NamedTuple('DC36', [('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC36({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC36) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC37(D15, NamedTuple('DC37', [('cf66', Any), ('cf67', Any), ('cf68', Any), ('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf66 == __o.cf66 and self.cf67 == __o.cf67 and self.cf68 == __o.cf68 and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC38(D15, NamedTuple('DC38', [('cf70', Any), ('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC35(D15, NamedTuple('DC35', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC35({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC35) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC40(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D16_DC40)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D16_DC42)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D16_DC39)

class D16_DC40(D16, NamedTuple('DC40', [('cf74', Any), ('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC40({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC40) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC41(D16, NamedTuple('DC41', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC42(D16, NamedTuple('DC42', [('cf78', Any), ('cf79', Any), ('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC39(D16, NamedTuple('DC39', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC44(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D17_DC44)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D17_DC43)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D17_DC45)

class D17_DC44(D17, NamedTuple('DC44', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC44({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC44) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC43(D17, NamedTuple('DC43', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC43({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC43) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC45(D17, NamedTuple('DC45', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC45({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC45) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC47(int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D18_DC47)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D18_DC46)

class D18_DC47(D18, NamedTuple('DC47', [('cf86', Any), ('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC47({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC47) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC46(D18, NamedTuple('DC46', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC46({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC46) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC49(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D19_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D19_DC48)

class D19_DC49(D19, NamedTuple('DC49', [('cf89', Any), ('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC49({_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC49) and self.cf89 == __o.cf89 and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC48(D19, NamedTuple('DC48', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC48({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC48) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC51()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D20_DC51)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D20_DC50)

class D20_DC51(D20, NamedTuple('DC51', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC51'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC51)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC50(D20, NamedTuple('DC50', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC50({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC50) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D21_DC52)

class D21_DC52(D21, NamedTuple('DC52', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC52({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC52) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC54(_dafny.CodePoint('D'), None, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D22_DC54)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D22_DC55)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D22_DC53)

class D22_DC54(D22, NamedTuple('DC54', [('cf95', Any), ('cf96', Any), ('cf97', Any), ('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC54({_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC54) and self.cf95 == __o.cf95 and self.cf96 == __o.cf96 and self.cf97 == __o.cf97 and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC55(D22, NamedTuple('DC55', [('cf99', Any), ('cf100', Any), ('cf101', Any), ('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC55({_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)}, {self.cf101.VerbatimString(True)}, {self.cf102.VerbatimString(True)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC55) and self.cf99 == __o.cf99 and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC53(D22, NamedTuple('DC53', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC53({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC53) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC57(False, _dafny.Array(None, 0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D23_DC57)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D23_DC56)

class D23_DC57(D23, NamedTuple('DC57', [('cf105', Any), ('cf106', Any), ('cf107', Any), ('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC57({_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC57) and self.cf105 == __o.cf105 and self.cf106 == __o.cf106 and self.cf107 == __o.cf107 and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC56(D23, NamedTuple('DC56', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC56({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC56) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC59()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D24_DC59)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D24_DC58)

class D24_DC59(D24, NamedTuple('DC59', [])):
    def __dafnystr__(self) -> str:
        return f'D24.DC59'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC59)
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC58(D24, NamedTuple('DC58', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC58({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC58) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m1(self, globalState):
        pass


class T2:
    pass
    def m2(self, p0, p1, p2, p3, globalState):
        pass

    def m3(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f3: int = int(0)
        self.f8: _dafny.MultiSet = _dafny.MultiSet({})
        self.f13: _dafny.Map = _dafny.Map({})
        self.f15: int = int(0)
        self.f18: int = int(0)
        self.f21: int = int(0)
        self.f24: int = int(0)
        self.f26: bool = False
        self._f1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f2: _dafny.Array = _dafny.Array(None, 0)
        self._f4: int = int(0)
        self._f5: int = int(0)
        self._f6: bool = False
        self._f7: bool = False
        self._f9: bool = False
        self._f10: _dafny.Array = _dafny.Array(None, 0)
        self._f11: int = int(0)
        self._f12: int = int(0)
        self._f14: bool = False
        self._f16: bool = False
        self._f17: _dafny.Set = _dafny.Set({})
        self._f19: int = int(0)
        self._f20: bool = False
        self._f22: _dafny.Map = _dafny.Map({})
        self._f23: _dafny.Array = _dafny.Array(None, 0)
        self._f25: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self).f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self).f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23
        (self).f24 = f24
        (self)._f25 = f25
        (self).f26 = f26

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
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
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
    def f14(self):
        return self._f14
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    @property
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23
    @property
    def f25(self):
        return self._f25

class C0:
    def  __init__(self):
        self._f36: int = int(0)
        self._f37: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f36, f37):
        (self)._f36 = f36
        (self)._f37 = f37

    @property
    def f36(self):
        return self._f36
    @property
    def f37(self):
        return self._f37

class C1(T1, T0):
    def  __init__(self):
        self._f29: _dafny.MultiSet = _dafny.MultiSet({})
        self._f30: _dafny.Seq = _dafny.Seq({})
        self._f27: int = int(0)
        self._f28: _dafny.Array = _dafny.Array(None, 0)
        self.f35: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f29(self):
        return self._f29
    @property
    def f30(self):
        return self._f30
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f35, f29, f30, f27, f28):
        (self).f35 = f35
        (self)._f29 = f29
        (self)._f30 = f30
        (self)._f27 = f27
        (self)._f28 = f28

    def fm3(self, p0, globalState):
        return ((0) - (default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l')]))), len((self).f30)))) not in (_dafny.SeqWithoutIsStrInference([(self).f27 for d_730_i0_ in range(default__.abs(784))]))

    def fm4(self, p0, p1, p2, globalState):
        return default__.safeDivisionInt((self).f27, (self).f27)

    def fm2(self, p0, p1, p2, globalState):
        return ((_dafny.Map({False: (self).f27})) | (_dafny.Map({self.f35: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tcfvt"))): self.f35}))}))) | (_dafny.Map({self.f35: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "istggpsd")))}))

    def fm11(self, p0, p1, p2, p3, globalState):
        return self.f35

    def m1(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        d_731_v0_: _dafny.Array
        nw127_ = _dafny.Array(False, 22)
        d_731_v0_ = nw127_
        d_731_v0_ = d_731_v0_
        d_732_v1_: _dafny.Seq
        d_732_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltinafpk"))
        d_733_v2_: str
        d_733_v2_ = _dafny.CodePoint('j')
        d_734_v3_: _dafny.Set
        d_734_v3_ = _dafny.Set({d_732_v1_, _dafny.SeqWithoutIsStrInference([d_733_v2_]), d_732_v1_, (_dafny.SeqWithoutIsStrInference([d_733_v2_ for d_735_i0_ in range(default__.abs(595))])).set(default__.safeIndex((0) - ((self).f27), len(_dafny.SeqWithoutIsStrInference([d_733_v2_ for d_736_i0_ in range(default__.abs(595))]))), d_733_v2_), d_732_v1_})
        d_737_v4_: _dafny.Map
        d_737_v4_ = _dafny.Map({self.f35: self.f35})
        d_738_v5_: _dafny.Seq
        d_738_v5_ = _dafny.SeqWithoutIsStrInference([self.f35])
        d_739_v6_: _dafny.Seq
        d_739_v6_ = _dafny.SeqWithoutIsStrInference([self.f35, self.f35, ((d_737_v4_)[self.f35] if (self.f35) in (d_737_v4_) else self.f35), self.f35, (d_738_v5_)[default__.safeIndex((self).f27, len(d_738_v5_))]])
        rhs87_ = d_732_v1_
        rhs88_ = not(((self).f27) <= ((self).f27))
        rhs89_ = self.f35
        rhs90_ = default__.safeDivisionInt(len(d_732_v1_), (self).f27)
        rhs91_ = ((self).fm4(d_734_v3_, self.f35, d_738_v5_, globalState)) > ((self).f27)
        lhs63_ = globalState
        lhs64_ = self
        lhs65_ = globalState
        d_732_v1_ = rhs87_
        lhs63_.f26 = rhs88_
        lhs64_.f35 = rhs89_
        r0 = rhs90_
        lhs65_.f26 = rhs91_
        d_740_v7_: _dafny.Array
        nw128_ = _dafny.Array(_dafny.Seq({}), 11)
        d_740_v7_ = nw128_
        d_741_v8_: _dafny.Set
        d_741_v8_ = _dafny.Set({d_731_v0_, d_731_v0_})
        index71_ = default__.safeIndex(498, (d_731_v0_).length(0))
        (d_731_v0_)[index71_] = (_dafny.Set({d_731_v0_})).isdisjoint(d_741_v8_)
        d_742_v9_: _dafny.Array
        nw129_ = _dafny.Array(int(0), 5)
        d_742_v9_ = nw129_
        index72_ = default__.safeIndex(86, (d_742_v9_).length(0))
        (d_742_v9_)[index72_] = (self).f27
        index73_ = default__.safeIndex(498, (d_731_v0_).length(0))
        index74_ = default__.safeIndex(86, (d_742_v9_).length(0))
        rhs92_ = d_740_v7_
        rhs93_ = (self.f35 if self.f35 else (self.f35 if self.f35 else self.f35))
        rhs94_ = d_733_v2_
        rhs95_ = -114
        rhs96_ = (self).f27
        lhs66_ = d_731_v0_
        lhs67_ = default__.safeIndex(498, (d_731_v0_).length(0))
        lhs68_ = d_742_v9_
        lhs69_ = default__.safeIndex(86, (d_742_v9_).length(0))
        d_740_v7_ = rhs92_
        lhs66_[lhs67_] = rhs93_
        d_733_v2_ = rhs94_
        lhs68_[lhs69_] = rhs95_
        r2 = rhs96_
        d_743_v10_: _dafny.Set
        d_743_v10_ = _dafny.Set({(self).f27})
        d_744_v11_: _dafny.Map
        d_744_v11_ = _dafny.Map({(d_742_v9_)[default__.safeIndex(86, (d_742_v9_).length(0))]: d_743_v10_})
        d_745_v12_: _dafny.Map
        d_745_v12_ = _dafny.Map({(d_731_v0_)[default__.safeIndex(498, (d_731_v0_).length(0))]: 145})
        d_744_v11_ = (d_744_v11_).set(((self).f27) + (len(d_745_v12_)), _dafny.Set({(d_742_v9_)[default__.safeIndex(86, (d_742_v9_).length(0))], (d_742_v9_)[default__.safeIndex(86, (d_742_v9_).length(0))], (self).f27}))
        d_746_v13_: _dafny.Map
        d_746_v13_ = _dafny.Map({(self).f29: d_732_v1_})
        d_747_v14_: bool
        d_747_v14_ = (d_731_v0_)[default__.safeIndex(498, (d_731_v0_).length(0))]
        index75_ = default__.safeIndex(86, (d_742_v9_).length(0))
        index76_ = default__.safeIndex(498, (d_731_v0_).length(0))
        rhs97_ = (self).f27
        rhs98_ = default__.fm12(default__.safeDivisionInt((self).f27, (d_742_v9_)[default__.safeIndex(86, (d_742_v9_).length(0))]), (d_731_v0_)[default__.safeIndex(498, (d_731_v0_).length(0))], globalState)
        rhs99_ = d_746_v13_
        rhs100_ = (d_732_v1_) == ((_dafny.SeqWithoutIsStrInference([(d_732_v1_)[default__.safeIndex((self).f27, len(d_732_v1_))] for d_748_i1_ in range(default__.abs(-42))])) + (d_732_v1_))
        rhs101_ = (d_747_v14_)
        lhs70_ = d_742_v9_
        lhs71_ = default__.safeIndex(86, (d_742_v9_).length(0))
        lhs72_ = self
        lhs73_ = d_731_v0_
        lhs74_ = default__.safeIndex(498, (d_731_v0_).length(0))
        lhs70_[lhs71_] = rhs97_
        d_733_v2_ = rhs98_
        d_746_v13_ = rhs99_
        lhs72_.f35 = rhs100_
        lhs73_[lhs74_] = rhs101_
        d_749_v15_: C0
        nw130_ = C0()
        nw130_.ctor__((d_742_v9_)[default__.safeIndex(86, (d_742_v9_).length(0))], d_732_v1_)
        d_749_v15_ = nw130_
        r0 = ((0) - (((self).f30)[default__.safeIndex((d_742_v9_)[default__.safeIndex(86, (d_742_v9_).length(0))], len((self).f30))])) * (len(d_732_v1_))
        r1 = (d_742_v9_)[default__.safeIndex(86, (d_742_v9_).length(0))]
        r2 = (0) - (default__.safeDivisionInt((d_742_v9_)[default__.safeIndex(86, (d_742_v9_).length(0))], (d_749_v15_).f36))
        return r0, r1, r2

    def m5(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        d_750_v0_: _dafny.Seq
        d_750_v0_ = _dafny.SeqWithoutIsStrInference([p1, p0])
        d_751_v1_: _dafny.Set
        d_751_v1_ = _dafny.Set({(self).f27, (self).f27})
        (globalState).f21 = default__.safeDivisionInt(len((d_750_v0_) + (d_750_v0_)), len(d_751_v1_))
        d_752_i0_: int
        d_752_i0_ = 0
        with _dafny.label("3"):
            while p0:
                with _dafny.c_label("3"):
                    if (d_752_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_752_i0_ = (d_752_i0_) + (1)
                    if True:
                        r0 = not(p0)
                        d_753_v2_: _dafny.Seq
                        d_753_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbfgu"))
                        d_753_v2_ = d_753_v2_
                        d_754_v3_: _dafny.Map
                        d_754_v3_ = _dafny.Map({self.f35: p1})
                        d_755_v4_: _dafny.Seq
                        d_755_v4_ = d_753_v2_
                        d_754_v3_ = (d_754_v3_).set(default__.fm0((self).f27, globalState), not((self).fm11(p1, (self).f27, d_755_v4_, (self).f27, globalState)))
                        (globalState).f26 = (self.f35 if not(((self).f30) <= ((self).f30)) else self.f35)
                        (globalState).f24 = ((self).f27) * ((self).f27)
                    elif True:
                        d_756_v5_: _dafny.Seq
                        d_756_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddutqmb"))
                        d_757_v6_: str
                        d_757_v6_ = _dafny.CodePoint('m')
                        d_758_v7_: C0
                        nw131_ = C0()
                        nw131_.ctor__(-849, (d_756_v5_).set(default__.safeIndex((self).f27, len(d_756_v5_)), d_757_v6_))
                        d_758_v7_ = nw131_
                        (globalState).f3 = ((self).f27) - (default__.fm1(globalState))
                        d_756_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
                        d_759_v8_: C0
                        nw132_ = C0()
                        nw132_.ctor__(default__.safeDivisionInt(735, (0) - ((self).f27)), (d_758_v7_).f37)
                        d_759_v8_ = nw132_
                        d_750_v0_ = default__.fm13((self).f27, ((0) - ((d_759_v8_).f36) if p0 else (d_758_v7_).f36), p0, -249, globalState)
                    d_760_v9_: C0
                    nw133_ = C0()
                    nw133_.ctor__((self).f27, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyxj")))
                    d_760_v9_ = nw133_
                    d_761_v10_: _dafny.Map
                    d_761_v10_ = _dafny.Map({p0: (self).f27})
                    (globalState).f21 = default__.safeModuloInt((self).f27, (301) + (((d_761_v10_)[p0] if (p0) in (d_761_v10_) else (self).f27)))
                    d_762_v11_: _dafny.Seq
                    d_762_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxypbryxn"))
                    d_763_v12_: _dafny.Array
                    nw134_ = _dafny.Array(None, 6)
                    nw134_[int(0)] = ((d_760_v9_).f36) - ((self).f27)
                    nw134_[int(1)] = len((_dafny.SeqWithoutIsStrInference([(d_760_v9_).f36 for d_764_i1_ in range(default__.abs(373))])) + ((self).f30))
                    nw134_[int(2)] = len(d_750_v0_)
                    nw134_[int(3)] = -645
                    nw134_[int(4)] = (self).f27
                    nw134_[int(5)] = (0) - (default__.safeDivisionInt((self).f27, (d_760_v9_).f36))
                    d_763_v12_ = nw134_
                    index77_ = default__.safeIndex(90, (d_763_v12_).length(0))
                    (d_763_v12_)[index77_] = (self).f27
                    d_765_v13_: _dafny.Seq
                    d_765_v13_ = _dafny.SeqWithoutIsStrInference([d_762_v11_])
                    index78_ = default__.safeIndex(90, (d_763_v12_).length(0))
                    def iife59_():
                        coll31_ = _dafny.Set()
                        compr_31_: _dafny.Seq
                        for compr_31_ in (d_765_v13_).Elements:
                            d_766_v14_: _dafny.Seq = compr_31_
                            if (d_766_v14_) in (d_765_v13_):
                                coll31_ = coll31_.union(_dafny.Set([d_766_v14_]))
                        return _dafny.Set(coll31_)
                    def iife60_():
                        coll32_ = _dafny.Set()
                        compr_32_: _dafny.Seq
                        for compr_32_ in (d_765_v13_).Elements:
                            d_767_v15_: _dafny.Seq = compr_32_
                            if (d_767_v15_) in (d_765_v13_):
                                coll32_ = coll32_.union(_dafny.Set([d_767_v15_]))
                        return _dafny.Set(coll32_)
                    def iife61_():
                        coll33_ = _dafny.Set()
                        compr_33_: int
                        for compr_33_ in _dafny.IntegerRange(993, 216):
                            d_768_v16_: int = compr_33_
                            if ((993) <= (d_768_v16_)) and ((d_768_v16_) < (216)):
                                coll33_ = coll33_.union(_dafny.Set([(d_768_v16_) - ((d_760_v9_).f36)]))
                        return _dafny.Set(coll33_)
                    rhs102_ = ((d_760_v9_).f37) + (((d_760_v9_).f37) + (default__.fm14(self.f35, p0, (d_750_v0_)[default__.safeIndex((d_760_v9_).f36, len(d_750_v0_))], (self).f27, globalState)))
                    rhs103_ = (self).fm4((iife59_()
                    ).intersection(iife60_()
                    ), False, d_750_v0_, globalState)
                    rhs104_ = (self).f27
                    rhs105_ = not (True) or (((self).f27) not in (default__.fm15(len(iife61_()
                    ), (d_760_v9_).f36, globalState)))
                    rhs106_ = len(default__.fm16((self).f27, globalState))
                    lhs75_ = d_763_v12_
                    lhs76_ = default__.safeIndex(90, (d_763_v12_).length(0))
                    lhs77_ = globalState
                    lhs78_ = self
                    lhs79_ = globalState
                    d_762_v11_ = rhs102_
                    lhs75_[lhs76_] = rhs103_
                    lhs77_.f18 = rhs104_
                    lhs78_.f35 = rhs105_
                    lhs79_.f0 = rhs106_
                    pass
            pass
        d_769_v17_: _dafny.Seq
        d_769_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elhuc"))
        d_770_v18_: _dafny.Seq
        d_770_v18_ = d_769_v17_
        d_771_v19_: _dafny.Map
        d_771_v19_ = _dafny.Map({default__.fm0(len((self).f30), globalState): (self).fm11(p1, (self).f27, d_770_v18_, (self).f27, globalState)})
        (globalState).f18 = (0) - (default__.safeModuloInt(len(d_751_v1_), default__.safeModuloInt(len(d_771_v19_), len(_dafny.Map({(self).f27: p0})))))
        hi3_ = (self).f27
        for d_772_i2_ in range(36, hi3_):
            if p1:
                d_773_v20_: _dafny.Array
                nw135_ = _dafny.Array(None, 4)
                nw135_[int(0)] = (d_769_v17_) + (d_769_v17_)
                nw135_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_774_i3_ in range(default__.abs(402))])
                nw135_[int(2)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_775_i4_ in range(default__.abs(743))])
                nw135_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "teeic"))
                d_773_v20_ = nw135_
                d_773_v20_ = (self).f28
                d_776_v21_: _dafny.Array
                nw136_ = _dafny.Array(False, 25)
                d_776_v21_ = nw136_
                index79_ = default__.safeIndex(29, (d_776_v21_).length(0))
                (d_776_v21_)[index79_] = p0
                d_777_v22_: D4
                d_777_v22_ = D4_DC7(d_776_v21_)
                d_778_v23_: _dafny.Seq
                d_778_v23_ = _dafny.SeqWithoutIsStrInference([d_776_v21_, (d_777_v22_).cf12, d_776_v21_])
                d_779_v24_: str
                d_779_v24_ = _dafny.CodePoint('t')
                d_780_v25_: _dafny.Array
                nw137_ = _dafny.Array(None, 25)
                nw137_[int(0)] = (d_772_i2_ if self.f35 else len(default__.fm14(p1, p1, p0, (0) - ((self).f27), globalState)))
                nw137_[int(1)] = (0) - ((self).f27)
                nw137_[int(2)] = (self).f27
                nw137_[int(3)] = ((self).f30)[default__.safeIndex((self).f27, len((self).f30))]
                nw137_[int(4)] = d_772_i2_
                nw137_[int(5)] = len(d_750_v0_)
                nw137_[int(6)] = len(_dafny.Set({d_772_i2_, d_772_i2_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_781_i5_ in range(default__.abs(351))]))}))
                nw137_[int(7)] = d_772_i2_
                nw137_[int(8)] = d_772_i2_
                nw137_[int(9)] = (self).f27
                nw137_[int(10)] = -750
                nw137_[int(11)] = (self).f27
                nw137_[int(12)] = (((self).f29).set(False, default__.abs(d_772_i2_))).cardinality
                nw137_[int(13)] = len(default__.fm17(p1, self.f35, p0, globalState))
                nw137_[int(14)] = len(_dafny.Map({(0) - ((self).f27): p0}))
                nw137_[int(15)] = len((d_778_v23_) + (d_778_v23_))
                nw137_[int(16)] = (180 if self.f35 else d_772_i2_)
                nw137_[int(17)] = d_772_i2_
                nw137_[int(18)] = (0) - (len(((d_769_v17_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lylquqbbn")))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_782_i7_ in range(default__.abs(-837))])) for d_783_i6_ in range(default__.abs(651))])), len((d_769_v17_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lylquqbbn"))))), d_779_v24_)))
                nw137_[int(19)] = ((self).f27) - (len((self).f30))
                nw137_[int(20)] = (0) - ((self).f27)
                nw137_[int(21)] = ((self).f29).cardinality
                nw137_[int(22)] = d_772_i2_
                nw137_[int(23)] = 669
                nw137_[int(24)] = (self).f27
                d_780_v25_ = nw137_
                d_784_v26_: _dafny.Map
                d_784_v26_ = _dafny.Map({(d_750_v0_) + (d_750_v0_): self.f35})
                index80_ = default__.safeIndex(29, (d_776_v21_).length(0))
                rhs107_ = len(d_784_v26_)
                rhs108_ = p1
                rhs109_ = d_780_v25_
                lhs80_ = globalState
                lhs81_ = d_776_v21_
                lhs82_ = default__.safeIndex(29, (d_776_v21_).length(0))
                lhs80_.f0 = rhs107_
                lhs81_[lhs82_] = rhs108_
                d_780_v25_ = rhs109_
                d_779_v24_ = d_779_v24_
                d_785_v27_: bool
                d_786_v28_: int
                out31_: bool
                out32_: int
                out31_, out32_ = default__.m0(not(not(not(True))), 470, (self).f27, 253, globalState)
                d_785_v27_ = out31_
                d_786_v28_ = out32_
                d_787_v29_: _dafny.Map
                d_787_v29_ = _dafny.Map({(d_785_v27_ if p0 else p1): d_770_v18_})
                d_787_v29_ = (d_787_v29_).set(self.f35, d_770_v18_)
            elif True:
                r0 = p1
                d_788_v30_: str
                d_788_v30_ = _dafny.CodePoint('w')
                d_789_v31_: _dafny.MultiSet
                d_789_v31_ = _dafny.MultiSet([(self).f27, (0) - (d_772_i2_), (self).f27])
                d_790_v32_: _dafny.Map
                d_790_v32_ = _dafny.Map({d_789_v31_: d_772_i2_})
                rhs110_ = _dafny.CodePoint('e')
                rhs111_ = (d_769_v17_) <= (d_769_v17_)
                rhs112_ = (not(((0) - ((D1_DC2(d_788_v30_, 168)).cf3)) == ((self).f27)) if (d_750_v0_)[default__.safeIndex(d_772_i2_, len(d_750_v0_))] else p0)
                rhs113_ = (d_772_i2_ if (d_769_v17_) < (d_769_v17_) else (self).f27)
                rhs114_ = d_790_v32_
                lhs83_ = globalState
                lhs84_ = globalState
                lhs85_ = globalState
                d_788_v30_ = rhs110_
                lhs83_.f26 = rhs111_
                lhs84_.f26 = rhs112_
                lhs85_.f18 = rhs113_
                d_790_v32_ = rhs114_
                (globalState).f26 = (p1) or (p0)
                def iife62_():
                    coll34_ = _dafny.Map()
                    compr_34_: int
                    for compr_34_ in _dafny.IntegerRange(750, 86):
                        d_791_v33_: int = compr_34_
                        if ((750) <= (d_791_v33_)) and ((d_791_v33_) < (86)):
                            coll34_[(d_791_v33_) - (d_772_i2_)] = _dafny.Map({self.f35: d_772_i2_})
                    return _dafny.Map(coll34_)
                (globalState).f21 = default__.safeModuloInt((self).f27, len(iife62_()
                ))
                (globalState).f15 = (len(d_769_v17_)) + ((d_772_i2_) * (-170))
            d_792_v34_: str
            d_792_v34_ = _dafny.CodePoint('c')
            d_793_v35_: _dafny.Map
            d_793_v35_ = _dafny.Map({d_792_v34_: p1})
            d_794_v36_: _dafny.Array
            nw138_ = _dafny.Array(None, 4)
            nw138_[int(0)] = not(p0)
            nw138_[int(1)] = ((d_793_v35_)[d_792_v34_] if (d_792_v34_) in (d_793_v35_) else not(self.f35))
            nw138_[int(2)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihcoyjmp")))) <= ((self).f27)
            nw138_[int(3)] = False
            d_794_v36_ = nw138_
            r1 = d_794_v36_
            index81_ = default__.safeIndex(806, (d_794_v36_).length(0))
            (d_794_v36_)[index81_] = self.f35
            d_795_v37_: _dafny.MultiSet
            d_795_v37_ = _dafny.MultiSet([default__.fm16((self).f27, globalState), (self).f30])
            index82_ = default__.safeIndex(806, (d_794_v36_).length(0))
            (d_794_v36_)[index82_] = (default__.fm18(globalState)).issubset(d_795_v37_)
            d_796_v38_: _dafny.Map
            d_796_v38_ = _dafny.Map({(d_794_v36_)[default__.safeIndex(806, (d_794_v36_).length(0))]: (self).f27})
            d_797_v39_: _dafny.Map
            d_797_v39_ = _dafny.Map({(self).f28: (0) - ((d_772_i2_) - (((d_796_v38_)[self.f35] if (self.f35) in (d_796_v38_) else (self).f27)))})
            d_797_v39_ = (d_797_v39_).set((self).f28, (self).f27)
        if (((self).f30)[default__.safeIndex(2, len((self).f30))]) <= (default__.safeModuloInt((self).f27, 604)):
            d_798_v40_: D4
            d_798_v40_ = D4_DC9((self).f27, p0)
            d_799_v41_: _dafny.Set
            d_799_v41_ = _dafny.Set({(d_798_v40_).cf17})
            d_800_v42_: _dafny.Seq
            d_800_v42_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p1})])
            d_801_v43_: _dafny.MultiSet
            d_801_v43_ = _dafny.MultiSet([(self).f27])
            d_802_v44_: _dafny.MultiSet
            d_802_v44_ = _dafny.MultiSet([(self).f27, len(d_751_v1_), (d_801_v43_).cardinality, (self).f27, (self).f27])
            d_803_v45_: _dafny.MultiSet
            d_803_v45_ = d_801_v43_
            d_804_v46_: _dafny.Set
            d_804_v46_ = _dafny.Set({(d_803_v45_)})
            rhs115_ = _dafny.Set({p0, (314) < ((0) - (len(d_800_v42_))), not((d_804_v46_).isdisjoint(d_804_v46_)), default__.fm0(len(d_799_v41_), globalState), p0})
            rhs116_ = (self).f27
            lhs86_ = globalState
            d_799_v41_ = rhs115_
            lhs86_.f3 = rhs116_
            d_805_v47_: _dafny.Array
            nw139_ = _dafny.Array(int(0), 18)
            d_805_v47_ = nw139_
            d_805_v47_ = d_805_v47_
            d_806_v48_: _dafny.Array
            nw140_ = _dafny.Array(False, 24)
            d_806_v48_ = nw140_
            r1 = d_806_v48_
            d_807_v49_: _dafny.Array
            nw141_ = _dafny.Array(D2.default()(), 27)
            d_807_v49_ = nw141_
            d_808_v50_: str
            d_808_v50_ = _dafny.CodePoint('v')
            d_809_v51_: _dafny.Map
            d_809_v51_ = _dafny.Map({d_808_v50_: p1})
            d_810_v52_: D2
            d_810_v52_ = D2_DC3(d_809_v51_)
            index83_ = default__.safeIndex(532, (d_807_v49_).length(0))
            (d_807_v49_)[index83_] = d_810_v52_
            index84_ = default__.safeIndex(532, (d_807_v49_).length(0))
            (d_807_v49_)[index84_] = d_810_v52_
            r0 = default__.fm0((((self).f29)[p1] if (p1) in ((self).f29) else (self).f27), globalState)
        elif True:
            d_811_v53_: _dafny.Array
            nw142_ = _dafny.Array(False, 17)
            d_811_v53_ = nw142_
            index85_ = default__.safeIndex(937, (d_811_v53_).length(0))
            (d_811_v53_)[index85_] = (d_750_v0_)[default__.safeIndex((self).f27, len(d_750_v0_))]
            index86_ = default__.safeIndex(937, (d_811_v53_).length(0))
            (d_811_v53_)[index86_] = False
            if ((self).f27) == ((self).f27):
                d_812_v54_: _dafny.Map
                d_812_v54_ = _dafny.Map({(self).f30: 962})
                d_813_v55_: str
                d_813_v55_ = _dafny.CodePoint('o')
                d_814_v56_: _dafny.Set
                d_814_v56_ = _dafny.Set({d_769_v17_, (d_769_v17_).set(default__.safeIndex(len(d_812_v54_), len(d_769_v17_)), d_813_v55_)})
                d_815_v57_: _dafny.Array
                nw143_ = _dafny.Array(None, 8)
                nw143_[int(0)] = (self).f27
                nw143_[int(1)] = len((self).f30)
                nw143_[int(2)] = 818
                nw143_[int(3)] = default__.safeModuloInt((0) - ((0) - ((self).f27)), (self).f27)
                nw143_[int(4)] = (0) - ((self).f27)
                nw143_[int(5)] = (0) - ((self).f27)
                nw143_[int(6)] = (self).fm4(d_814_v56_, p0, d_750_v0_, globalState)
                nw143_[int(7)] = (0) - (((self).f27) * ((self).f27))
                d_815_v57_ = nw143_
                index87_ = default__.safeIndex(767, (d_815_v57_).length(0))
                (d_815_v57_)[index87_] = (0) - ((self).f27)
                index88_ = default__.safeIndex(767, (d_815_v57_).length(0))
                (d_815_v57_)[index88_] = default__.safeModuloInt((self).f27, (self).f27)
                (globalState).f26 = not(self.f35)
                index89_ = default__.safeIndex(767, (d_815_v57_).length(0))
                rhs117_ = len(_dafny.SeqWithoutIsStrInference([((self).f27) > ((self).f27), p1]))
                rhs118_ = (((self).f30)[default__.safeIndex((d_815_v57_)[default__.safeIndex(767, (d_815_v57_).length(0))], len((self).f30))]) * ((d_815_v57_)[default__.safeIndex(767, (d_815_v57_).length(0))])
                rhs119_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvpnx"))
                lhs87_ = globalState
                lhs88_ = d_815_v57_
                lhs89_ = default__.safeIndex(767, (d_815_v57_).length(0))
                lhs87_.f24 = rhs117_
                lhs88_[lhs89_] = rhs118_
                d_769_v17_ = rhs119_
                (globalState).f3 = (d_815_v57_)[default__.safeIndex(767, (d_815_v57_).length(0))]
                d_816_v58_: _dafny.Map
                d_816_v58_ = _dafny.Map({(self).f27: d_771_v19_})
                rhs120_ = not(((len(d_769_v17_)) + ((d_815_v57_)[default__.safeIndex(767, (d_815_v57_).length(0))])) >= (default__.fm1(globalState)))
                rhs121_ = (((d_816_v58_)[(self).f27] if ((self).f27) in (d_816_v58_) else d_771_v19_)) | (d_771_v19_)
                lhs90_ = globalState
                lhs90_.f26 = rhs120_
                d_771_v19_ = rhs121_
            elif True:
                d_817_v59_: _dafny.Set
                d_817_v59_ = _dafny.Set({p1})
                d_818_v60_: _dafny.Map
                d_818_v60_ = _dafny.Map({p0: (_dafny.Set({p0})) - (d_817_v59_)})
                d_818_v60_ = d_818_v60_
                (globalState).f24 = default__.safeModuloInt((self).f27, (self).f27)
                r1 = d_811_v53_
                d_819_v61_: _dafny.Array
                nw144_ = _dafny.Array(_dafny.Set({}), 3)
                d_819_v61_ = nw144_
                index90_ = default__.safeIndex(196, (d_819_v61_).length(0))
                (d_819_v61_)[index90_] = d_751_v1_
                d_820_v62_: _dafny.Map
                d_820_v62_ = _dafny.Map({(self).f27: (self).f27})
                index91_ = default__.safeIndex(196, (d_819_v61_).length(0))
                def iife63_():
                    coll35_ = _dafny.Set()
                    compr_35_: int
                    for compr_35_ in (d_820_v62_).keys.Elements:
                        d_821_v63_: int = compr_35_
                        if (d_821_v63_) in (d_820_v62_):
                            coll35_ = coll35_.union(_dafny.Set([(d_821_v63_) - ((782) * (566))]))
                    return _dafny.Set(coll35_)
                (d_819_v61_)[index91_] = iife63_()
                
                d_822_v64_: _dafny.Map
                d_822_v64_ = _dafny.Map({self.f35: (self).f27})
                d_823_v65_: _dafny.Seq
                d_823_v65_ = _dafny.SeqWithoutIsStrInference([d_822_v64_, (d_822_v64_) | (d_822_v64_), _dafny.Map({self.f35: len(d_769_v17_)}), d_822_v64_])
                d_823_v65_ = _dafny.SeqWithoutIsStrInference([d_822_v64_])
            d_824_v66_: str
            d_824_v66_ = _dafny.CodePoint('y')
            d_825_v67_: C0
            nw145_ = C0()
            nw145_.ctor__((self).f27, (d_769_v17_).set(default__.safeIndex((self).f27, len(d_769_v17_)), d_824_v66_))
            d_825_v67_ = nw145_
            d_826_v68_: D6
            d_826_v68_ = D6_DC11(d_825_v67_)
            d_827_v69_: _dafny.Array
            nw146_ = _dafny.Array(None, 23)
            nw146_[int(0)] = d_825_v67_
            nw146_[int(1)] = d_825_v67_
            nw146_[int(2)] = d_825_v67_
            nw146_[int(3)] = d_825_v67_
            nw146_[int(4)] = d_825_v67_
            nw146_[int(5)] = d_825_v67_
            nw146_[int(6)] = d_825_v67_
            nw146_[int(7)] = d_825_v67_
            nw146_[int(8)] = d_825_v67_
            nw146_[int(9)] = d_825_v67_
            nw146_[int(10)] = d_825_v67_
            nw146_[int(11)] = d_825_v67_
            nw146_[int(12)] = d_825_v67_
            nw146_[int(13)] = d_825_v67_
            nw146_[int(14)] = d_825_v67_
            nw146_[int(15)] = d_825_v67_
            nw146_[int(16)] = d_825_v67_
            nw146_[int(17)] = d_825_v67_
            nw146_[int(18)] = d_825_v67_
            nw146_[int(19)] = d_825_v67_
            nw146_[int(20)] = d_825_v67_
            nw146_[int(21)] = d_825_v67_
            nw146_[int(22)] = (d_826_v68_).cf19
            d_827_v69_ = nw146_
            rhs122_ = d_827_v69_
            rhs123_ = (((d_825_v67_).f37).set(default__.safeIndex((0) - ((d_825_v67_).f36), len((d_825_v67_).f37)), _dafny.CodePoint('b'))) + ((d_825_v67_).f37)
            rhs124_ = (d_825_v67_).f36
            lhs91_ = globalState
            d_827_v69_ = rhs122_
            d_769_v17_ = rhs123_
            lhs91_.f24 = rhs124_
            d_828_v70_: _dafny.Map
            d_828_v70_ = _dafny.Map({(d_825_v67_).f37: (0) - ((self).f27)})
            d_828_v70_ = (d_828_v70_).set(_dafny.SeqWithoutIsStrInference([d_824_v66_ for d_829_i8_ in range(default__.abs(935))]), (d_825_v67_).f36)
            d_830_v71_: _dafny.Map
            d_830_v71_ = _dafny.Map({(d_811_v53_)[default__.safeIndex(937, (d_811_v53_).length(0))]: (d_825_v67_).f36})
            d_830_v71_ = (d_830_v71_).set((d_811_v53_)[default__.safeIndex(937, (d_811_v53_).length(0))], (-461) * (len(d_771_v19_)))
        if (d_750_v0_)[default__.safeIndex((self).f27, len(d_750_v0_))]:
            (globalState).f26 = p1
            (globalState).f3 = len(((D7_DC13(_dafny.SeqWithoutIsStrInference([self.f35, True]))).cf23 if p0 else d_750_v0_))
            (globalState).f0 = (self).f27
            if (default__.fm14(self.f35, self.f35, p1, (self).f27, globalState)) == ((d_769_v17_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qk")))):
                (globalState).f15 = (self).f27
                d_831_v72_: C0
                nw147_ = C0()
                nw147_.ctor__((self).f27, d_769_v17_)
                d_831_v72_ = nw147_
                (globalState).f26 = self.f35
                d_832_v73_: _dafny.Array
                nw148_ = _dafny.Array(False, 22)
                d_832_v73_ = nw148_
                index92_ = default__.safeIndex(266, (d_832_v73_).length(0))
                (d_832_v73_)[index92_] = p0
                d_833_v74_: D6
                d_833_v74_ = D6_DC11(d_831_v72_)
                d_834_v75_: _dafny.Map
                d_834_v75_ = _dafny.Map({p1: d_833_v74_})
                d_835_v76_: _dafny.Map
                d_835_v76_ = _dafny.Map({(len(d_834_v75_)) * ((self).f27): True})
                index93_ = default__.safeIndex(266, (d_832_v73_).length(0))
                (d_832_v73_)[index93_] = ((d_835_v76_)[341] if (341) in (d_835_v76_) else self.f35)
                d_836_v77_: _dafny.Set
                d_836_v77_ = _dafny.Set({_dafny.MultiSet(d_750_v0_)})
                index94_ = default__.safeIndex(266, (d_832_v73_).length(0))
                rhs125_ = p0
                rhs126_ = not((d_836_v77_).issubset(d_836_v77_))
                rhs127_ = ((self).f27) * ((d_831_v72_).f36)
                lhs92_ = d_832_v73_
                lhs93_ = default__.safeIndex(266, (d_832_v73_).length(0))
                lhs94_ = globalState
                lhs95_ = globalState
                lhs92_[lhs93_] = rhs125_
                lhs94_.f26 = rhs126_
                lhs95_.f24 = rhs127_
            elif True:
                d_837_v78_: D8
                d_837_v78_ = D8_DC15(d_771_v19_)
                d_838_v79_: _dafny.Map
                d_838_v79_ = _dafny.Map({(d_837_v78_).cf28: p0})
                d_750_v0_ = (_dafny.SeqWithoutIsStrInference([((d_838_v79_)[_dafny.Map({self.f35: p0})] if (_dafny.Map({self.f35: p0})) in (d_838_v79_) else p0)])) + ((d_750_v0_).set(default__.safeIndex((self).f27, len(d_750_v0_)), p1))
                d_839_v80_: str
                d_839_v80_ = _dafny.CodePoint('u')
                d_840_v81_: _dafny.MultiSet
                d_840_v81_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_841_i9_ in range(default__.abs(-284))]) if p0 else (d_769_v17_).set(default__.safeIndex((self).f27, len(d_769_v17_)), d_839_v80_))])
                (globalState).f3 = ((d_840_v81_)[(d_769_v17_) + (d_769_v17_)] if ((d_769_v17_) + (d_769_v17_)) in (d_840_v81_) else (0) - (((self).f27 if not(self.f35) else len(_dafny.SeqWithoutIsStrInference([len(d_750_v0_) for d_842_i10_ in range(default__.abs(702))])))))
                (globalState).f18 = (self).f27
                (self).f35 = not(self.f35)
                d_843_v82_: _dafny.MultiSet
                d_843_v82_ = _dafny.MultiSet([(self).f27, len(d_769_v17_), (self).f27])
                d_843_v82_ = d_843_v82_
            d_844_v83_: _dafny.Array
            nw149_ = _dafny.Array(_dafny.Map({}), 24)
            d_844_v83_ = nw149_
            index95_ = default__.safeIndex(298, (d_844_v83_).length(0))
            def iife64_():
                coll36_ = _dafny.Map()
                compr_36_: str
                for compr_36_ in (default__.fm19(globalState)).Elements:
                    d_845_v84_: str = compr_36_
                    if (d_845_v84_) in (default__.fm19(globalState)):
                        coll36_[d_845_v84_] = (self).f27
                return _dafny.Map(coll36_)
            (d_844_v83_)[index95_] = iife64_()
            
            d_846_v85_: str
            d_846_v85_ = _dafny.CodePoint('t')
            d_847_v86_: _dafny.Map
            d_847_v86_ = _dafny.Map({d_846_v85_: (self).f27})
            index96_ = default__.safeIndex(298, (d_844_v83_).length(0))
            (d_844_v83_)[index96_] = d_847_v86_
        elif True:
            (globalState).f0 = (self).f27
            (globalState).f18 = (self).f27
            d_848_v87_: _dafny.Map
            d_848_v87_ = _dafny.Map({(self).f27: (self).f27})
            d_849_v88_: str
            d_849_v88_ = _dafny.CodePoint('k')
            d_850_v89_: D4
            d_850_v89_ = D4_DC8(d_849_v88_, 57, p1)
            d_848_v87_ = (d_848_v87_).set((d_850_v89_).cf14, (self).f27)
            d_851_v90_: _dafny.Array
            nw150_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_851_v90_ = nw150_
            d_852_v91_: _dafny.Seq
            d_852_v91_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_853_i11_ in range(default__.abs(208))]))])
            rhs128_ = d_851_v90_
            rhs129_ = (self).f27
            rhs130_ = (self).f30
            rhs131_ = default__.safeDivisionInt((self).f27, len(d_769_v17_))
            rhs132_ = (self).f27
            lhs96_ = globalState
            lhs97_ = globalState
            lhs98_ = globalState
            d_851_v90_ = rhs128_
            lhs96_.f21 = rhs129_
            d_852_v91_ = rhs130_
            lhs97_.f0 = rhs131_
            lhs98_.f0 = rhs132_
            d_769_v17_ = (d_769_v17_) + ((d_769_v17_) + (default__.fm14(self.f35, self.f35, p1, (self).f27, globalState)))
        r0 = p0
        d_854_v92_: _dafny.Map
        d_854_v92_ = _dafny.Map({(0) - ((self).f27): (self).f27})
        d_855_v93_: _dafny.Map
        d_855_v93_ = _dafny.Map({default__.fm1(globalState): d_854_v92_})
        d_856_v94_: str
        d_856_v94_ = _dafny.CodePoint('r')
        d_857_v95_: _dafny.Map
        d_857_v95_ = _dafny.Map({self.f35: d_856_v94_})
        nw151_ = _dafny.Array(None, 21)
        nw151_[int(0)] = self.f35
        nw151_[int(1)] = p1
        nw151_[int(2)] = (p1) and (p0)
        nw151_[int(3)] = (d_750_v0_) <= (d_750_v0_)
        nw151_[int(4)] = p0
        nw151_[int(5)] = p0
        nw151_[int(6)] = p0
        nw151_[int(7)] = p1
        nw151_[int(8)] = ((self).f27) >= ((self).f27)
        nw151_[int(9)] = p0
        nw151_[int(10)] = p0
        nw151_[int(11)] = p1
        nw151_[int(12)] = (len((self).f30)) >= ((self).f27)
        nw151_[int(13)] = (d_750_v0_)[default__.safeIndex((self).f27, len(d_750_v0_))]
        nw151_[int(14)] = p1
        nw151_[int(15)] = (self).fm11(p1, len(d_855_v93_), d_770_v18_, len(d_857_v95_), globalState)
        nw151_[int(16)] = (self).fm11(p1, (self).f27, d_770_v18_, (self).f27, globalState)
        nw151_[int(17)] = (self).fm11(p0, len(d_769_v17_), d_770_v18_, (self).f27, globalState)
        nw151_[int(18)] = p1
        nw151_[int(19)] = (self).fm11(p1, (self).f27, d_770_v18_, len(d_769_v17_), globalState)
        nw151_[int(20)] = self.f35
        r1 = nw151_
        d_858_v96_: _dafny.MultiSet
        d_858_v96_ = _dafny.MultiSet([(self).f27])
        d_859_v97_: _dafny.Map
        d_859_v97_ = _dafny.Map({len((d_769_v17_ if p1 else d_769_v17_)): not((d_858_v96_).ispropersubset(d_858_v96_))})
        r2 = d_859_v97_
        return r0, r1, r2

    def m6(self, p0, globalState):
        r0: int = int(0)
        d_860_v0_: _dafny.Array
        nw152_ = _dafny.Array(int(0), 13)
        d_860_v0_ = nw152_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_860_v0_).length(0)):
            d_861_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_861_i0_)) and ((d_861_i0_) < ((d_860_v0_).length(0)))):
                (d_860_v0_)[(d_861_i0_)] = (d_861_i0_) + (len(_dafny.Set({p0})))
        (globalState).f24 = ((self).f30)[default__.safeIndex(p0, len((self).f30))]
        d_862_v1_: _dafny.Array
        nw153_ = _dafny.Array(None, 4)
        nw153_[int(0)] = self.f35
        nw153_[int(1)] = default__.fm0((self).f27, globalState)
        nw153_[int(2)] = self.f35
        nw153_[int(3)] = self.f35
        d_862_v1_ = nw153_
        d_863_v2_: D4
        d_863_v2_ = D4_DC7(d_862_v1_)
        pat_let_tv7_ = d_862_v1_
        d_864_v3_: _dafny.Map
        def iife65_(_pat_let14_0):
            def iife66_(d_865_dt__update__tmp_h0_):
                def iife67_(_pat_let15_0):
                    def iife68_(d_866_dt__update_hcf12_h0_):
                        return D4_DC7(d_866_dt__update_hcf12_h0_)
                    return iife68_(_pat_let15_0)
                return iife67_(pat_let_tv7_)
            return iife66_(_pat_let14_0)
        d_864_v3_ = _dafny.Map({iife65_(d_863_v2_): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f35, False, self.f35, self.f35, self.f35]))})
        d_867_v4_: _dafny.Map
        d_867_v4_ = _dafny.Map({self.f35: len(d_864_v3_)})
        d_868_v5_: _dafny.Seq
        d_868_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "it"))
        d_869_v6_: _dafny.Map
        d_869_v6_ = _dafny.Map({d_868_v5_: (self).f27})
        d_870_v7_: _dafny.Map
        d_870_v7_ = _dafny.Map({(default__.fm16(((d_867_v4_)[self.f35] if (self.f35) in (d_867_v4_) else p0), globalState)) + ((self).f30): default__.safeDivisionInt((self).f27, ((d_869_v6_)[d_868_v5_] if (d_868_v5_) in (d_869_v6_) else (self).f27))})
        d_870_v7_ = (d_870_v7_).set(((self).f30) + ((self).f30), (self).f27)
        d_868_v5_ = (default__.fm14(True, self.f35, self.f35, len(d_868_v5_), globalState)) + (default__.fm14(self.f35, self.f35, self.f35, len(default__.fm14(self.f35, self.f35, self.f35, (self).f27, globalState)), globalState))
        if not((self.f35 if self.f35 else self.f35)):
            d_871_v8_: str
            d_871_v8_ = _dafny.CodePoint('w')
            d_872_v9_: D2
            d_872_v9_ = D2_DC4(len(_dafny.Map({len(_dafny.Map({self.f35: False})): self.f35})), ((self).f27) - (p0), (0) - (p0), d_871_v8_, (d_868_v5_) + (d_868_v5_))
            d_873_v10_: _dafny.Map
            d_873_v10_ = _dafny.Map({d_871_v8_: self.f35})
            d_874_v11_: D2
            d_874_v11_ = D2_DC3(d_873_v10_)
            pat_let_tv8_ = d_873_v10_
            d_875_v12_: D2
            def iife69_(_pat_let16_0):
                def iife70_(d_876_dt__update__tmp_h1_):
                    def iife71_(_pat_let17_0):
                        def iife72_(d_877_dt__update_hcf4_h0_):
                            return D2_DC3(d_877_dt__update_hcf4_h0_)
                        return iife72_(_pat_let17_0)
                    return iife71_(pat_let_tv8_)
                return iife70_(_pat_let16_0)
            d_875_v12_ = D2_DC3((iife69_(d_874_v11_)).cf4)
            d_878_v13_: D2
            d_878_v13_ = D2_DC5(d_875_v12_)
            rhs133_ = (d_868_v5_) + (d_868_v5_)
            rhs134_ = d_872_v9_
            rhs135_ = (self).f27
            rhs136_ = d_878_v13_
            rhs137_ = False
            lhs99_ = globalState
            lhs100_ = globalState
            d_868_v5_ = rhs133_
            d_872_v9_ = rhs134_
            lhs99_.f21 = rhs135_
            d_878_v13_ = rhs136_
            lhs100_.f26 = rhs137_
            d_879_v14_: _dafny.Set
            d_879_v14_ = _dafny.Set({self.f35, False})
            d_880_v15_: _dafny.Seq
            d_880_v15_ = _dafny.SeqWithoutIsStrInference([d_879_v14_])
            rhs138_ = self.f35
            rhs139_ = (((d_880_v15_) + (_dafny.SeqWithoutIsStrInference([d_879_v14_]))) + (d_880_v15_)).set(default__.safeIndex((self).f27, len(((d_880_v15_) + (_dafny.SeqWithoutIsStrInference([d_879_v14_]))) + (d_880_v15_))), d_879_v14_)
            lhs101_ = self
            lhs101_.f35 = rhs138_
            d_880_v15_ = rhs139_
            (globalState).f26 = self.f35
            d_881_v16_: _dafny.Seq
            d_881_v16_ = _dafny.SeqWithoutIsStrInference([default__.fm14(self.f35, self.f35, self.f35, p0, globalState), d_868_v5_])
            d_882_v17_: D6
            d_882_v17_ = D6_DC12(self.f35, default__.fm20(self.f35, globalState), (d_881_v16_)[default__.safeIndex(p0, len(d_881_v16_))])
            source14_ = d_882_v17_
            if source14_.is_DC12:
                d_883___mcc_h0_ = source14_.cf20
                d_884___mcc_h1_ = source14_.cf21
                d_885___mcc_h2_ = source14_.cf22
                d_886_cf22_ = d_885___mcc_h2_
                d_887_cf21_ = d_884___mcc_h1_
                d_888_cf20_ = d_883___mcc_h0_
                d_889_v18_: _dafny.Array
                def lambda30_(d_890_p0_):
                    def lambda31_(d_891_i1_):
                        return (_dafny.Set({d_890_p0_})) - ((D9_DC17(_dafny.Set({((self).f29).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vlmvh")))}))).cf31)

                    return lambda31_

                init16_ = lambda30_(p0)
                nw154_ = _dafny.Array(None, 11)
                for i0_16_ in range(nw154_.length(0)):
                    nw154_[i0_16_] = init16_(i0_16_)
                d_889_v18_ = nw154_
                rhs140_ = (d_889_v18_ if self.f35 else d_889_v18_)
                rhs141_ = d_871_v8_
                d_889_v18_ = rhs140_
                d_871_v8_ = rhs141_
                d_892_v19_: _dafny.Set
                d_892_v19_ = _dafny.Set({d_871_v8_})
                index97_ = default__.safeIndex(350, (d_860_v0_).length(0))
                (d_860_v0_)[index97_] = len(d_892_v19_)
                index98_ = default__.safeIndex(350, (d_860_v0_).length(0))
                (d_860_v0_)[index98_] = (p0) - ((p0) - ((self).f27))
                d_893_v20_: _dafny.MultiSet
                d_893_v20_ = _dafny.MultiSet([(self).f27, p0, (d_860_v0_)[default__.safeIndex(350, (d_860_v0_).length(0))], (self).f27, (self).f27])
                d_894_v21_: D8
                d_894_v21_ = D8_DC16((self).f27, (self).f27)
                d_895_v22_: _dafny.Map
                d_895_v22_ = _dafny.Map({((self).f30)[default__.safeIndex(len(default__.fm13(p0, 612, not(d_888_cf20_), len(d_886_cf22_), globalState)), len((self).f30))]: d_894_v21_})
                d_896_v23_: _dafny.Map
                d_896_v23_ = _dafny.Map({len(_dafny.Map({d_871_v8_: (self).f27})): (self).f27})
                d_897_v24_: bool
                d_898_v25_: int
                out33_: bool
                out34_: int
                out33_, out34_ = default__.m0((d_893_v20_).isdisjoint(_dafny.MultiSet([(d_860_v0_)[default__.safeIndex(350, (d_860_v0_).length(0))]])), len(d_895_v22_), (self).f27, ((d_896_v23_)[(self).f27] if ((self).f27) in (d_896_v23_) else len(d_868_v5_)), globalState)
                d_897_v24_ = out33_
                d_898_v25_ = out34_
                index99_ = default__.safeIndex(462, (d_862_v1_).length(0))
                (d_862_v1_)[index99_] = d_888_cf20_
                index100_ = default__.safeIndex(462, (d_862_v1_).length(0))
                (d_862_v1_)[index100_] = self.f35
            elif True:
                d_899___mcc_h3_ = source14_.cf19
                d_900_cf19_ = d_899___mcc_h3_
                (globalState).f26 = self.f35
                d_901_v26_: _dafny.MultiSet
                d_901_v26_ = _dafny.MultiSet([(d_900_cf19_).f36, -805])
                d_902_v27_: _dafny.Seq
                d_902_v27_ = _dafny.SeqWithoutIsStrInference([self.f35, self.f35])
                d_903_v28_: D1
                d_903_v28_ = D1_DC2(d_871_v8_, (0) - (len(d_902_v27_)))
                d_904_v29_: _dafny.Set
                d_904_v29_ = _dafny.Set({(d_901_v26_).cardinality, (d_903_v28_).cf3})
                d_905_v30_: bool
                d_906_v31_: _dafny.Array
                d_907_v32_: _dafny.Map
                out35_: bool
                out36_: _dafny.Array
                out37_: _dafny.Map
                out35_, out36_, out37_ = (self).m5(self.f35, (d_904_v29_).isdisjoint(d_904_v29_), globalState)
                d_905_v30_ = out35_
                d_906_v31_ = out36_
                d_907_v32_ = out37_
                (globalState).f24 = (self).f27
                index101_ = default__.safeIndex(822, (d_860_v0_).length(0))
                (d_860_v0_)[index101_] = ((self).f27) + (p0)
                index102_ = default__.safeIndex(822, (d_860_v0_).length(0))
                (d_860_v0_)[index102_] = default__.safeDivisionInt(len(d_868_v5_), p0)
            d_908_v33_: _dafny.Array
            nw155_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_908_v33_ = nw155_
            index103_ = default__.safeIndex(53, (d_908_v33_).length(0))
            (d_908_v33_)[index103_] = d_862_v1_
            index104_ = default__.safeIndex(53, (d_908_v33_).length(0))
            (d_908_v33_)[index104_] = d_862_v1_
        elif True:
            d_909_v34_: _dafny.Map
            d_909_v34_ = _dafny.Map({767: self.f35})
            (globalState).f21 = ((p0) - ((0) - (len(d_909_v34_)))) + ((self).f27)
            (globalState).f26 = (p0) == (default__.fm1(globalState))
            d_910_v35_: C0
            nw156_ = C0()
            nw156_.ctor__((self).f27, d_868_v5_)
            d_910_v35_ = nw156_
            d_911_v36_: _dafny.Map
            d_911_v36_ = _dafny.Map({364: d_910_v35_})
            d_912_v37_: D6
            d_912_v37_ = D6_DC12(self.f35, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f35]) for d_913_i2_ in range(default__.abs(216))]), (d_910_v35_).f37)
            d_914_v38_: D6
            d_914_v38_ = D6_DC11(((d_911_v36_)[len((d_912_v37_).cf22)] if (len((d_912_v37_).cf22)) in (d_911_v36_) else d_910_v35_))
            d_914_v38_ = d_914_v38_
            d_909_v34_ = (d_909_v34_) | (_dafny.Map({(self).f27: True}))
            (globalState).f3 = (d_910_v35_).f36
        d_915_v39_: _dafny.Map
        d_915_v39_ = _dafny.Map({self.f35: d_868_v5_})
        d_916_v40_: C0
        nw157_ = C0()
        nw157_.ctor__((self).f27, ((d_915_v39_)[True] if (True) in (d_915_v39_) else d_868_v5_))
        d_916_v40_ = nw157_
        d_917_v41_: _dafny.Map
        d_917_v41_ = _dafny.Map({self.f35: self.f35})
        d_918_v42_: D8
        d_918_v42_ = D8_DC15(d_917_v41_)
        pat_let_tv9_ = p0
        def lambda32_(source15_):
            if source15_.is_DC16:
                d_919___mcc_h4_ = source15_.cf29
                d_920___mcc_h5_ = source15_.cf30
                d_921_cf30_ = d_920___mcc_h5_
                d_922_cf29_ = d_919___mcc_h4_
                return d_922_cf29_
            elif True:
                d_923___mcc_h6_ = source15_.cf28
                d_924_cf28_ = d_923___mcc_h6_
                return pat_let_tv9_

        r0 = lambda32_(d_918_v42_)
        return r0


class C2(T2):
    def  __init__(self):
        self._f31: _dafny.MultiSet = _dafny.MultiSet({})
        self._f32: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f31(self):
        return self._f31
    @property
    def f32(self):
        return self._f32
    def ctor__(self, f31, f32):
        (self)._f31 = f31
        (self)._f32 = f32

    def fm5(self, globalState):
        def lambda33_(source16_):
            if source16_.is_DC18:
                d_925___mcc_h0_ = source16_.cf32
                d_926_cf32_ = d_925___mcc_h0_
                return (d_926_cf32_) or (d_926_cf32_)
            elif source16_.is_DC19:
                d_927___mcc_h1_ = source16_.cf33
                d_928_cf33_ = d_927___mcc_h1_
                return (D9_DC20(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okdff"))), d_928_cf33_, d_928_cf33_, d_928_cf33_, _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_928_cf33_, d_928_cf33_])): len(_dafny.Set({d_928_cf33_}))}))).cf36
            elif source16_.is_DC20:
                d_929___mcc_h2_ = source16_.cf34
                d_930___mcc_h3_ = source16_.cf35
                d_931___mcc_h4_ = source16_.cf36
                d_932___mcc_h5_ = source16_.cf37
                d_933___mcc_h6_ = source16_.cf38
                d_934_cf38_ = d_933___mcc_h6_
                d_935_cf37_ = d_932___mcc_h5_
                d_936_cf36_ = d_931___mcc_h4_
                d_937_cf35_ = d_930___mcc_h3_
                d_938_cf34_ = d_929___mcc_h2_
                return not(d_936_cf36_)
            elif True:
                d_939___mcc_h7_ = source16_.cf31
                d_940_cf31_ = d_939___mcc_h7_
                return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_941_i0_ in range(default__.abs(-888))])) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "im")))

        return not(lambda33_(D9_DC20((self).f32, not(not(True)), True, False, _dafny.Map({-15: ((_dafny.MultiSet([(self).f32, (self).f32]))).cardinality}))))

    def fm6(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]) for d_942_i0_ in range(default__.abs(-875))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, not(True)])]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(False), True])]))

    def fm21(self, globalState):
        return not((len((_dafny.Map({(self).f32: True})) | (_dafny.Map({(self).f32: True})))) != (((self).f31).cardinality))

    def fm22(self, p0, p1, p2, globalState):
        source17_ = D1_DC1(_dafny.Map({(0) - ((self).f32): not(True)}))
        if source17_.is_DC2:
            d_943___mcc_h0_ = source17_.cf2
            d_944___mcc_h1_ = source17_.cf3
            d_945_cf3_ = d_944___mcc_h1_
            d_946_cf2_ = d_943___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdr"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxqk")))
        elif True:
            d_947___mcc_h2_ = source17_.cf1
            d_948_cf1_ = d_947___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgvy"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tp")))

    def m2(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        d_949_v0_: str
        d_949_v0_ = _dafny.CodePoint('i')
        d_949_v0_ = d_949_v0_
        d_950_v1_: _dafny.Array
        nw158_ = _dafny.Array(False, 14)
        d_950_v1_ = nw158_
        d_951_v2_: _dafny.Map
        d_951_v2_ = _dafny.Map({p3: p3})
        d_952_v3_: _dafny.Map
        d_952_v3_ = _dafny.Map({d_950_v1_: d_951_v2_})
        d_952_v3_ = (d_952_v3_).set(d_950_v1_, d_951_v2_)
        if p1:
            (globalState).f26 = default__.fm0(p2, globalState)
            d_953_v4_: _dafny.MultiSet
            d_953_v4_ = _dafny.MultiSet([d_950_v1_, d_950_v1_])
            if (d_953_v4_).issubset(d_953_v4_):
                (globalState).f24 = 308
                d_954_v5_: _dafny.Map
                d_954_v5_ = _dafny.Map({p2: (self).f31})
                d_955_v6_: _dafny.Map
                d_955_v6_ = _dafny.Map({p1: p0})
                index105_ = default__.safeIndex(766, (d_950_v1_).length(0))
                (d_950_v1_)[index105_] = ((self).f31).isdisjoint(((d_954_v5_)[len(d_955_v6_)] if (len(d_955_v6_)) in (d_954_v5_) else (self).f31))
                d_956_v7_: _dafny.Seq
                d_956_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojlonqco"))
                d_957_v8_: D4
                d_957_v8_ = D4_DC8(d_949_v0_, len(d_956_v7_), p1)
                d_958_v9_: _dafny.Set
                d_958_v9_ = _dafny.Set({d_957_v8_})
                d_959_v10_: _dafny.Map
                d_959_v10_ = _dafny.Map({True: (0) - (len(default__.fm23(len(d_955_v6_), not(True), globalState)))})
                d_960_v11_: _dafny.Seq
                d_960_v11_ = _dafny.SeqWithoutIsStrInference([(len(d_956_v7_)) - (len(d_958_v9_)), (0) - (p3), p3, default__.safeModuloInt(len(d_956_v7_), len(d_959_v10_))])
                d_961_v12_: _dafny.MultiSet
                d_961_v12_ = _dafny.MultiSet([d_959_v10_])
                index106_ = default__.safeIndex(766, (d_950_v1_).length(0))
                rhs142_ = (d_956_v7_) != (d_956_v7_)
                rhs143_ = _dafny.SeqWithoutIsStrInference([p3, len((_dafny.SeqWithoutIsStrInference([p2]) if p1 else d_960_v11_))])
                rhs144_ = (0) - (((d_961_v12_)[_dafny.Map({not(p0): (0) - (p3)})] if (_dafny.Map({not(p0): (0) - (p3)})) in (d_961_v12_) else default__.safeDivisionInt((self).f32, p3)))
                rhs145_ = d_950_v1_
                rhs146_ = (self).fm21(globalState)
                lhs102_ = d_950_v1_
                lhs103_ = default__.safeIndex(766, (d_950_v1_).length(0))
                lhs104_ = globalState
                lhs105_ = globalState
                lhs102_[lhs103_] = rhs142_
                d_960_v11_ = rhs143_
                lhs104_.f15 = rhs144_
                d_950_v1_ = rhs145_
                lhs105_.f26 = rhs146_
                d_962_v13_: _dafny.Seq
                d_962_v13_ = d_956_v7_
                d_963_v14_: _dafny.Array
                nw159_ = _dafny.Array(False, 14)
                d_963_v14_ = nw159_
                d_964_v15_: bool
                d_964_v15_ = (d_950_v1_)[default__.safeIndex(766, (d_950_v1_).length(0))]
                index107_ = default__.safeIndex(8, (d_963_v14_).length(0))
                (d_963_v14_)[index107_] = d_964_v15_
                index108_ = default__.safeIndex(8, (d_963_v14_).length(0))
                rhs147_ = p1
                rhs148_ = d_962_v13_
                rhs149_ = p1
                rhs150_ = d_964_v15_
                rhs151_ = p0
                lhs106_ = globalState
                lhs107_ = globalState
                lhs108_ = d_963_v14_
                lhs109_ = default__.safeIndex(8, (d_963_v14_).length(0))
                lhs110_ = globalState
                lhs106_.f26 = rhs147_
                d_962_v13_ = rhs148_
                lhs107_.f26 = rhs149_
                lhs108_[lhs109_] = rhs150_
                lhs110_.f26 = rhs151_
                d_965_v16_: _dafny.Map
                d_965_v16_ = _dafny.Map({(d_956_v7_ if False else d_956_v7_): ((self).f31) - ((self).f31)})
                d_965_v16_ = d_965_v16_
                rhs152_ = (d_950_v1_)[default__.safeIndex(766, (d_950_v1_).length(0))]
                rhs153_ = (self).fm22((d_959_v10_) | (default__.fm24(p3, globalState)), False, ((d_959_v10_)[(d_950_v1_)[default__.safeIndex(766, (d_950_v1_).length(0))]] if ((d_950_v1_)[default__.safeIndex(766, (d_950_v1_).length(0))]) in (d_959_v10_) else ((self).f31).cardinality), globalState)
                rhs154_ = not(p1)
                lhs111_ = globalState
                lhs112_ = globalState
                lhs111_.f26 = rhs152_
                d_956_v7_ = rhs153_
                lhs112_.f26 = rhs154_
            elif True:
                d_966_v17_: _dafny.Seq
                d_966_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkvarnrk"))
                d_967_v18_: C0
                nw160_ = C0()
                nw160_.ctor__((0) - (p3), d_966_v17_)
                d_967_v18_ = nw160_
                d_968_v19_: D6
                d_968_v19_ = D6_DC11(d_967_v18_)
                d_969_v20_: _dafny.Map
                d_969_v20_ = _dafny.Map({d_968_v19_: ((d_966_v17_) + (d_966_v17_)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_949_v0_])), len((d_966_v17_) + (d_966_v17_))), d_949_v0_)})
                d_969_v20_ = (d_969_v20_).set(d_968_v19_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvq"))) + (d_966_v17_))
                d_970_v21_: _dafny.Map
                d_970_v21_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_949_v0_ for d_971_i0_ in range(default__.abs(886))]): p0})
                d_972_v22_: C0
                nw161_ = C0()
                nw161_.ctor__((len(d_970_v21_)) + ((self).f32), (d_967_v18_).f37)
                d_972_v22_ = nw161_
                d_973_v23_: _dafny.Seq
                d_973_v23_ = _dafny.SeqWithoutIsStrInference([p0])
                d_973_v23_ = _dafny.SeqWithoutIsStrInference([False, p1, ((d_967_v18_).f37) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxvupb"))), p0, p1])
                index109_ = default__.safeIndex(515, (d_950_v1_).length(0))
                (d_950_v1_)[index109_] = p1
                d_974_v24_: _dafny.Set
                d_974_v24_ = _dafny.Set({d_950_v1_})
                index110_ = default__.safeIndex(515, (d_950_v1_).length(0))
                (d_950_v1_)[index110_] = (_dafny.Set({d_950_v1_})).isdisjoint(d_974_v24_)
                (globalState).f24 = 315
            d_975_v25_: _dafny.Seq
            d_975_v25_ = _dafny.SeqWithoutIsStrInference([d_950_v1_])
            d_976_v26_: _dafny.Seq
            d_976_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbbsuc"))
            d_977_v27_: _dafny.MultiSet
            d_977_v27_ = _dafny.MultiSet([d_976_v26_, d_976_v26_])
            rhs155_ = ((self).f32) > ((len(d_975_v25_)) - (p3))
            rhs156_ = d_977_v27_
            lhs113_ = globalState
            lhs114_ = globalState
            lhs113_.f26 = rhs155_
            lhs114_.f8 = rhs156_
            (globalState).f26 = not(p0)
            (globalState).f26 = ((0) - (p3)) != ((self).f32)
        elif True:
            (globalState).f26 = (p2) >= (p2)
            (globalState).f26 = p1
            d_978_v28_: _dafny.Array
            nw162_ = _dafny.Array(None, 2)
            nw162_[int(0)] = p2
            nw162_[int(1)] = default__.fm1(globalState)
            d_978_v28_ = nw162_
            d_978_v28_ = d_978_v28_
            d_979_v29_: _dafny.Array
            nw163_ = _dafny.Array(None, 6)
            nw163_[int(0)] = d_950_v1_
            nw163_[int(1)] = d_950_v1_
            nw163_[int(2)] = d_950_v1_
            nw163_[int(3)] = d_950_v1_
            nw163_[int(4)] = d_950_v1_
            nw163_[int(5)] = d_950_v1_
            d_979_v29_ = nw163_
            d_979_v29_ = d_979_v29_
            d_980_v30_: _dafny.Array
            nw164_ = _dafny.Array(D8.default()(), 13)
            d_980_v30_ = nw164_
            d_981_v31_: D8
            d_981_v31_ = D8_DC16((self).f32, p3)
            index111_ = default__.safeIndex(261, (d_980_v30_).length(0))
            (d_980_v30_)[index111_] = d_981_v31_
            index112_ = default__.safeIndex(261, (d_980_v30_).length(0))
            rhs157_ = d_979_v29_
            rhs158_ = d_981_v31_
            rhs159_ = (0) - (((0) - (default__.fm1(globalState))) + ((0) - (p3)))
            lhs115_ = d_980_v30_
            lhs116_ = default__.safeIndex(261, (d_980_v30_).length(0))
            d_979_v29_ = rhs157_
            lhs115_[lhs116_] = rhs158_
            r1 = rhs159_
        d_982_v32_: _dafny.Map
        d_982_v32_ = _dafny.Map({(self).fm5(globalState): (self).f32})
        d_983_v33_: _dafny.Seq
        d_983_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efdnwcf"))
        d_984_v34_: D2
        d_984_v34_ = D2_DC4(len(d_982_v32_), 343, p3, d_949_v0_, d_983_v33_)
        d_985_v35_: _dafny.MultiSet
        d_985_v35_ = _dafny.MultiSet([D2_DC5(d_984_v34_)])
        d_986_v36_: D2
        d_986_v36_ = D2_DC5(d_984_v34_)
        d_987_v37_: _dafny.Map
        d_987_v37_ = _dafny.Map({p3: False})
        d_988_v38_: _dafny.Seq
        d_988_v38_ = _dafny.SeqWithoutIsStrInference([((d_987_v37_)[len(_dafny.Map({p1: p1}))] if (len(_dafny.Map({p1: p1}))) in (d_987_v37_) else p1), True])
        d_989_v39_: _dafny.Seq
        d_989_v39_ = _dafny.SeqWithoutIsStrInference([p2])
        d_990_v40_: D10
        d_990_v40_ = D10_DC21(d_989_v39_)
        d_991_v41_: _dafny.Array
        nw165_ = _dafny.Array(None, 9)
        nw165_[int(0)] = d_983_v33_
        nw165_[int(1)] = d_983_v33_
        nw165_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
        nw165_[int(3)] = d_983_v33_
        nw165_[int(4)] = d_983_v33_
        nw165_[int(5)] = d_983_v33_
        nw165_[int(6)] = d_983_v33_
        nw165_[int(7)] = d_983_v33_
        nw165_[int(8)] = d_983_v33_
        d_991_v41_ = nw165_
        d_992_v42_: C1
        nw166_ = C1()
        nw166_.ctor__((_dafny.MultiSet([d_986_v36_])).issubset(d_985_v35_), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0]))) - (_dafny.MultiSet(d_988_v38_)), (d_990_v40_).cf39, default__.safeDivisionInt(-160, (self).f32), (d_991_v41_ if p0 else d_991_v41_))
        d_992_v42_ = nw166_
        index113_ = default__.safeIndex(348, (d_950_v1_).length(0))
        (d_950_v1_)[index113_] = (p0) not in (_dafny.Set({p0}))
        pat_let_tv10_ = d_992_v42_
        pat_let_tv11_ = d_992_v42_
        index114_ = default__.safeIndex(348, (d_950_v1_).length(0))
        def lambda34_(source18_):
            if source18_.is_DC22:
                d_993___mcc_h0_ = source18_.cf40
                d_994___mcc_h1_ = source18_.cf41
                d_995_cf41_ = d_994___mcc_h1_
                d_996_cf40_ = d_993___mcc_h0_
                return pat_let_tv10_.f35
            elif True:
                d_997___mcc_h2_ = source18_.cf39
                d_998_cf39_ = d_997___mcc_h2_
                return pat_let_tv11_.f35

        (d_950_v1_)[index114_] = lambda34_(d_990_v40_)
        d_999_v43_: _dafny.Array
        nw167_ = _dafny.Array(_dafny.Seq({}), 29)
        d_999_v43_ = nw167_
        index115_ = default__.safeIndex(227, (d_999_v43_).length(0))
        (d_999_v43_)[index115_] = d_988_v38_
        d_1000_v44_: _dafny.Set
        d_1000_v44_ = _dafny.Set({True})
        d_1001_v45_: _dafny.Seq
        d_1001_v45_ = _dafny.SeqWithoutIsStrInference([d_1000_v44_])
        index116_ = default__.safeIndex(227, (d_999_v43_).length(0))
        index117_ = default__.safeIndex(348, (d_950_v1_).length(0))
        rhs160_ = (d_988_v38_) + (d_988_v38_)
        rhs161_ = (p1) == (not(False))
        rhs162_ = not((d_1000_v44_).ispropersubset((d_1001_v45_)[default__.safeIndex((0) - ((self).f32), len(d_1001_v45_))]))
        lhs117_ = d_999_v43_
        lhs118_ = default__.safeIndex(227, (d_999_v43_).length(0))
        lhs119_ = d_950_v1_
        lhs120_ = default__.safeIndex(348, (d_950_v1_).length(0))
        lhs121_ = globalState
        lhs117_[lhs118_] = rhs160_
        lhs119_[lhs120_] = rhs161_
        lhs121_.f26 = rhs162_
        r0 = d_950_v1_
        r1 = p2
        return r0, r1

    def m3(self, p0, p1, p2, p3, globalState):
        d_1002_v0_: _dafny.Array
        nw168_ = _dafny.Array(int(0), 24)
        d_1002_v0_ = nw168_
        index118_ = default__.safeIndex(100, (d_1002_v0_).length(0))
        (d_1002_v0_)[index118_] = (len(p2) if p0 else len(p2))
        d_1003_v1_: _dafny.Map
        d_1003_v1_ = _dafny.Map({p0: p0})
        d_1004_v2_: D8
        d_1004_v2_ = D8_DC15(d_1003_v1_)
        pat_let_tv12_ = p3
        pat_let_tv13_ = p3
        index119_ = default__.safeIndex(100, (d_1002_v0_).length(0))
        def lambda35_(source19_):
            if source19_.is_DC16:
                d_1005___mcc_h0_ = source19_.cf29
                d_1006___mcc_h1_ = source19_.cf30
                d_1007_cf30_ = d_1006___mcc_h1_
                d_1008_cf29_ = d_1005___mcc_h0_
                return (_dafny.MultiSet([_dafny.Set({True}), _dafny.Set({pat_let_tv12_, pat_let_tv13_})])).cardinality
            elif True:
                d_1009___mcc_h2_ = source19_.cf28
                d_1010_cf28_ = d_1009___mcc_h2_
                return -118

        (d_1002_v0_)[index119_] = lambda35_(d_1004_v2_)
        if False:
            d_1011_v3_: _dafny.Seq
            d_1011_v3_ = _dafny.SeqWithoutIsStrInference([(self).f32])
            d_1012_v4_: _dafny.Map
            d_1012_v4_ = _dafny.Map({p1: p2})
            d_1013_v5_: _dafny.Map
            d_1013_v5_ = _dafny.Map({p3: 680})
            d_1014_v6_: _dafny.Seq
            d_1014_v6_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1 for d_1015_i1_ in range(default__.abs(250))])])
            d_1016_v7_: _dafny.Array
            nw169_ = _dafny.Array(None, 9)
            nw169_[int(0)] = p2
            nw169_[int(1)] = p2
            nw169_[int(2)] = ((d_1012_v4_)[p1] if (p1) in (d_1012_v4_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnpxuk")))
            nw169_[int(3)] = p2
            nw169_[int(4)] = (self).fm22(d_1013_v5_, p3, (self).f32, globalState)
            nw169_[int(5)] = (d_1014_v6_)[default__.safeIndex((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], len(d_1014_v6_))]
            nw169_[int(6)] = p2
            nw169_[int(7)] = p2
            nw169_[int(8)] = p2
            d_1016_v7_ = nw169_
            d_1017_v8_: C1
            nw170_ = C1()
            nw170_.ctor__(p0, (self).f31, (d_1011_v3_) + (_dafny.SeqWithoutIsStrInference([(self).f32 for d_1018_i0_ in range(default__.abs(192))])), 798, d_1016_v7_)
            d_1017_v8_ = nw170_
            d_1019_v9_: _dafny.Array
            def lambda36_(d_1020_v8_):
                def lambda37_(d_1021_i2_):
                    return d_1020_v8_.f35

                return lambda37_

            init17_ = lambda36_(d_1017_v8_)
            nw171_ = _dafny.Array(None, 15)
            for i0_17_ in range(nw171_.length(0)):
                nw171_[i0_17_] = init17_(i0_17_)
            d_1019_v9_ = nw171_
            d_1022_v10_: _dafny.Map
            d_1022_v10_ = _dafny.Map({d_1019_v9_: p3})
            d_1023_v11_: C0
            nw172_ = C0()
            nw172_.ctor__((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], p2)
            d_1023_v11_ = nw172_
            d_1024_v12_: _dafny.Map
            d_1024_v12_ = _dafny.Map({d_1017_v8_.f35: d_1023_v11_})
            d_1025_v13_: _dafny.Seq
            d_1025_v13_ = _dafny.SeqWithoutIsStrInference([p1 for d_1026_i4_ in range(default__.abs(186))])
            d_1027_v14_: _dafny.Array
            nw173_ = _dafny.Array(None, 22)
            nw173_[int(0)] = d_1023_v11_
            nw173_[int(1)] = d_1023_v11_
            nw173_[int(2)] = d_1023_v11_
            nw173_[int(3)] = d_1023_v11_
            nw173_[int(4)] = d_1023_v11_
            nw173_[int(5)] = d_1023_v11_
            nw173_[int(6)] = d_1023_v11_
            nw173_[int(7)] = d_1023_v11_
            nw173_[int(8)] = d_1023_v11_
            nw173_[int(9)] = d_1023_v11_
            nw173_[int(10)] = d_1023_v11_
            nw173_[int(11)] = d_1023_v11_
            nw173_[int(12)] = d_1023_v11_
            nw173_[int(13)] = d_1023_v11_
            nw173_[int(14)] = d_1023_v11_
            nw173_[int(15)] = d_1023_v11_
            nw173_[int(16)] = d_1023_v11_
            nw173_[int(17)] = d_1023_v11_
            nw173_[int(18)] = d_1023_v11_
            nw173_[int(19)] = d_1023_v11_
            nw173_[int(20)] = d_1023_v11_
            nw173_[int(21)] = ((d_1024_v12_)[(d_1017_v8_).fm11(p3, len(_dafny.SeqWithoutIsStrInference([(d_1023_v11_).f36 for d_1028_i3_ in range(default__.abs(-262))])), d_1025_v13_, (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], globalState)] if ((d_1017_v8_).fm11(p3, len(_dafny.SeqWithoutIsStrInference([(d_1023_v11_).f36 for d_1029_i3_ in range(default__.abs(-262))])), d_1025_v13_, (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], globalState)) in (d_1024_v12_) else d_1023_v11_)
            d_1027_v14_ = nw173_
            index120_ = default__.safeIndex(351, (d_1027_v14_).length(0))
            (d_1027_v14_)[index120_] = d_1023_v11_
            d_1030_v15_: _dafny.Seq
            d_1030_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhmyrad"))
            index121_ = default__.safeIndex(351, (d_1027_v14_).length(0))
            rhs163_ = d_1022_v10_
            rhs164_ = d_1023_v11_
            rhs165_ = p2
            lhs122_ = d_1027_v14_
            lhs123_ = default__.safeIndex(351, (d_1027_v14_).length(0))
            d_1022_v10_ = rhs163_
            lhs122_[lhs123_] = rhs164_
            d_1030_v15_ = rhs165_
            (globalState).f26 = ((d_1023_v11_).f36) != ((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))])
            (globalState).f26 = (default__.safeDivisionInt(-20, len(d_1011_v3_))) not in (_dafny.SeqWithoutIsStrInference([(d_1023_v11_).f36 for d_1031_i5_ in range(default__.abs(140))]))
            d_1030_v15_ = d_1030_v15_
        elif True:
            d_1002_v0_ = d_1002_v0_
            d_1032_v16_: _dafny.Seq
            d_1032_v16_ = _dafny.SeqWithoutIsStrInference([(self).f32, (0) - (len(default__.fm25(p0, globalState))), (self).f32])
            d_1033_v17_: _dafny.Array
            nw174_ = _dafny.Array(None, 2)
            nw174_[int(0)] = p2
            nw174_[int(1)] = p2
            d_1033_v17_ = nw174_
            d_1034_v18_: C1
            nw175_ = C1()
            nw175_.ctor__(p3, (self).f31, d_1032_v16_, (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], d_1033_v17_)
            d_1034_v18_ = nw175_
            d_1035_v19_: D4
            d_1035_v19_ = D4_DC8(_dafny.CodePoint('x'), (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], p0)
            d_1036_v20_: _dafny.Seq
            d_1036_v20_ = p2
            d_1037_v21_: _dafny.Map
            d_1037_v21_ = _dafny.Map({d_1035_v19_: d_1036_v20_})
            if (d_1035_v19_) not in (d_1037_v21_):
                d_1038_v22_: C1
                nw176_ = C1()
                nw176_.ctor__(((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]) == (-977), _dafny.MultiSet([p3, not(p0)]), ((d_1032_v16_) + (d_1032_v16_)).set(default__.safeIndex((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], len((d_1032_v16_) + (d_1032_v16_))), len(p2)), (self).f32, d_1033_v17_)
                d_1038_v22_ = nw176_
                d_1039_v23_: _dafny.Set
                d_1039_v23_ = _dafny.Set({p3})
                d_1040_v24_: _dafny.MultiSet
                d_1040_v24_ = _dafny.MultiSet([len(d_1039_v23_), (self).f32, 419])
                index122_ = default__.safeIndex(100, (d_1002_v0_).length(0))
                index123_ = default__.safeIndex(100, (d_1002_v0_).length(0))
                def iife73_():
                    def iife76_():
                        coll40_ = _dafny.Set()
                        compr_40_: int
                        for compr_40_ in (d_1040_v24_).Elements:
                            d_1044_v25_: int = compr_40_
                            if (d_1044_v25_) in (d_1040_v24_):
                                def iife77_():
                                    coll41_ = _dafny.Set()
                                    compr_41_: int
                                    for compr_41_ in _dafny.IntegerRange(-421, 239):
                                        d_1045_v26_: int = compr_41_
                                        if ((-421) <= (d_1045_v26_)) and ((d_1045_v26_) < (239)):
                                            coll41_ = coll41_.union(_dafny.Set([(d_1045_v26_) + (len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdfc"))})))]))
                                    return _dafny.Set(coll41_)
                                coll40_ = coll40_.union(_dafny.Set([(d_1044_v25_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([len(iife77_()
), 543]))))]))
                        return _dafny.Set(coll40_)
                    coll37_ = _dafny.Set()
                    def iife74_():
                        coll38_ = _dafny.Set()
                        compr_38_: int
                        for compr_38_ in (d_1040_v24_).Elements:
                            d_1041_v25_: int = compr_38_
                            if (d_1041_v25_) in (d_1040_v24_):
                                def iife75_():
                                    coll39_ = _dafny.Set()
                                    compr_39_: int
                                    for compr_39_ in _dafny.IntegerRange(-421, 239):
                                        d_1042_v26_: int = compr_39_
                                        if ((-421) <= (d_1042_v26_)) and ((d_1042_v26_) < (239)):
                                            coll39_ = coll39_.union(_dafny.Set([(d_1042_v26_) + (len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdfc"))})))]))
                                    return _dafny.Set(coll39_)
                                coll38_ = coll38_.union(_dafny.Set([(d_1041_v25_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([len(iife75_()
), 543]))))]))
                        return _dafny.Set(coll38_)
                    compr_37_: int
                    for compr_37_ in (iife74_()
                    ).Elements:
                        d_1043_v27_: int = compr_37_
                        if (d_1043_v27_) in (iife76_()
                        ):
                            coll37_ = coll37_.union(_dafny.Set([default__.safeModuloInt(d_1043_v27_, 891)]))
                    return _dafny.Set(coll37_)
                rhs166_ = (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]
                rhs167_ = (((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))] if p3 else len(iife73_()
                ))) + ((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))])
                rhs168_ = 135
                rhs169_ = (d_1038_v22_).fm4(_dafny.Set({p2}), p0, _dafny.SeqWithoutIsStrInference([p0, d_1038_v22_.f35]), globalState)
                rhs170_ = (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]
                lhs124_ = globalState
                lhs125_ = d_1002_v0_
                lhs126_ = default__.safeIndex(100, (d_1002_v0_).length(0))
                lhs127_ = globalState
                lhs128_ = globalState
                lhs129_ = d_1002_v0_
                lhs130_ = default__.safeIndex(100, (d_1002_v0_).length(0))
                lhs124_.f3 = rhs166_
                lhs125_[lhs126_] = rhs167_
                lhs127_.f18 = rhs168_
                lhs128_.f24 = rhs169_
                lhs129_[lhs130_] = rhs170_
                d_1046_v28_: _dafny.Array
                nw177_ = _dafny.Array(_dafny.Set({}), 21)
                d_1046_v28_ = nw177_
                d_1047_v29_: _dafny.Seq
                d_1047_v29_ = _dafny.SeqWithoutIsStrInference([d_1034_v18_])
                d_1048_v30_: _dafny.Array
                nw178_ = _dafny.Array(None, 24)
                nw178_[int(0)] = d_1038_v22_
                nw178_[int(1)] = d_1034_v18_
                nw178_[int(2)] = d_1034_v18_
                nw178_[int(3)] = d_1034_v18_
                nw178_[int(4)] = d_1038_v22_
                nw178_[int(5)] = d_1038_v22_
                nw178_[int(6)] = d_1038_v22_
                nw178_[int(7)] = d_1038_v22_
                nw178_[int(8)] = d_1034_v18_
                nw178_[int(9)] = d_1038_v22_
                nw178_[int(10)] = d_1038_v22_
                nw178_[int(11)] = d_1038_v22_
                nw178_[int(12)] = d_1038_v22_
                nw178_[int(13)] = d_1034_v18_
                nw178_[int(14)] = d_1038_v22_
                nw178_[int(15)] = d_1034_v18_
                nw178_[int(16)] = d_1034_v18_
                nw178_[int(17)] = (d_1047_v29_)[default__.safeIndex((self).f32, len(d_1047_v29_))]
                nw178_[int(18)] = d_1034_v18_
                nw178_[int(19)] = d_1038_v22_
                nw178_[int(20)] = d_1038_v22_
                nw178_[int(21)] = d_1034_v18_
                nw178_[int(22)] = d_1038_v22_
                nw178_[int(23)] = d_1034_v18_
                d_1048_v30_ = nw178_
                d_1049_v31_: _dafny.Set
                d_1049_v31_ = _dafny.Set({d_1048_v30_})
                index124_ = default__.safeIndex(510, (d_1046_v28_).length(0))
                (d_1046_v28_)[index124_] = (d_1049_v31_) - (d_1049_v31_)
                index125_ = default__.safeIndex(510, (d_1046_v28_).length(0))
                (d_1046_v28_)[index125_] = (_dafny.Set({d_1048_v30_}) if d_1034_v18_.f35 else _dafny.Set({d_1048_v30_}))
                d_1050_v32_: _dafny.Seq
                d_1050_v32_ = _dafny.SeqWithoutIsStrInference([not(p3), ((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]) <= (-645), p0, d_1038_v22_.f35, (self).fm5(globalState)])
                d_1050_v32_ = ((_dafny.SeqWithoutIsStrInference([d_1034_v18_.f35, d_1034_v18_.f35])) + ((_dafny.SeqWithoutIsStrInference([p3])).set(default__.safeIndex((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], len(_dafny.SeqWithoutIsStrInference([p3]))), p0))) + (d_1050_v32_)
                d_1051_v33_: _dafny.Array
                def lambda38_(d_1052_p3_):
                    def lambda39_(d_1053_i6_):
                        return d_1052_p3_

                    return lambda39_

                init18_ = lambda38_(p3)
                nw179_ = _dafny.Array(None, 8)
                for i0_18_ in range(nw179_.length(0)):
                    nw179_[i0_18_] = init18_(i0_18_)
                d_1051_v33_ = nw179_
                d_1054_v34_: D4
                d_1054_v34_ = D4_DC7(d_1051_v33_)
                d_1055_v35_: _dafny.Map
                d_1055_v35_ = _dafny.Map({(self).f32: d_1051_v33_})
                d_1056_v36_: _dafny.Array
                nw180_ = _dafny.Array(None, 6)
                nw180_[int(0)] = (d_1051_v33_ if False else d_1051_v33_)
                nw180_[int(1)] = d_1051_v33_
                nw180_[int(2)] = d_1051_v33_
                nw180_[int(3)] = (d_1054_v34_).cf12
                nw180_[int(4)] = ((d_1055_v35_)[(self).f32] if ((self).f32) in (d_1055_v35_) else d_1051_v33_)
                nw180_[int(5)] = d_1051_v33_
                d_1056_v36_ = nw180_
                index126_ = default__.safeIndex(725, (d_1056_v36_).length(0))
                (d_1056_v36_)[index126_] = d_1051_v33_
                index127_ = default__.safeIndex(725, (d_1056_v36_).length(0))
                nw181_ = _dafny.Array(False, 12)
                (d_1056_v36_)[index127_] = nw181_
            elif True:
                d_1057_v37_: T1
                nw182_ = C1()
                nw182_.ctor__(False, (self).f31, d_1032_v16_, (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], d_1033_v17_)
                d_1057_v37_ = nw182_
                d_1058_v38_: _dafny.Map
                d_1058_v38_ = _dafny.Map({d_1057_v37_: p0})
                d_1058_v38_ = (d_1058_v38_).set(d_1057_v37_, (((d_1057_v37_).f30)[default__.safeIndex((self).f32, len((d_1057_v37_).f30))]) > (-231))
                d_1059_v39_: _dafny.Map
                d_1059_v39_ = _dafny.Map({(self).f32: p3})
                d_1060_v40_: C1
                nw183_ = C1()
                nw183_.ctor__(p0, (self).f31, _dafny.SeqWithoutIsStrInference([(d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], (D2_DC4((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], (0) - (len(d_1059_v39_)), (self).f32, p1, p2)).cf6]), ((_dafny.MultiSet([p0])).cardinality) * ((d_1057_v37_).f27), (d_1057_v37_).f28)
                d_1060_v40_ = nw183_
                d_1061_v41_: _dafny.Seq
                d_1061_v41_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1062_v42_: _dafny.Map
                d_1062_v42_ = _dafny.Map({d_1034_v18_.f35: (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]})
                d_1063_v43_: C0
                nw184_ = C0()
                nw184_.ctor__(len(_dafny.SeqWithoutIsStrInference([len(d_1061_v41_), len((self).fm22(d_1062_v42_, d_1034_v18_.f35, default__.fm1(globalState), globalState))])), p2)
                d_1063_v43_ = nw184_
                (globalState).f26 = not ((self).fm5(globalState)) or (not((d_1034_v18_.f35) or (d_1034_v18_.f35)))
                (d_1034_v18_).f35 = p3
            d_1064_v44_: _dafny.Map
            d_1064_v44_ = _dafny.Map({default__.fm26(p1, (self).f32, (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], globalState): (self).f32})
            (globalState).f18 = default__.safeModuloInt(((d_1064_v44_)[p1] if (p1) in (d_1064_v44_) else (self).f32), (self).f32)
            d_1065_v45_: C0
            nw185_ = C0()
            nw185_.ctor__(len(d_1032_v16_), p2)
            d_1065_v45_ = nw185_
        hi4_ = 220
        for d_1066_i7_ in range((0) - ((self).f32), hi4_):
            if (d_1066_i7_) < (d_1066_i7_):
                d_1067_v46_: C0
                nw186_ = C0()
                nw186_.ctor__((self).f32, p2)
                d_1067_v46_ = nw186_
                d_1068_v47_: _dafny.Set
                d_1068_v47_ = _dafny.Set({d_1067_v46_, d_1067_v46_, d_1067_v46_})
                (globalState).f26 = (d_1068_v47_).ispropersubset(d_1068_v47_)
                d_1069_v48_: str
                d_1069_v48_ = _dafny.CodePoint('k')
                d_1070_v49_: _dafny.Seq
                d_1070_v49_ = _dafny.SeqWithoutIsStrInference([p3])
                rhs171_ = not(default__.fm0((len(d_1070_v49_) if p3 else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1071_i8_ in range(default__.abs(90))]))), globalState))
                rhs172_ = ((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]) - (default__.safeModuloInt((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], (self).f32))
                rhs173_ = ((d_1067_v46_).f36) - (d_1066_i7_)
                rhs174_ = p1
                lhs131_ = globalState
                lhs132_ = globalState
                lhs133_ = globalState
                lhs131_.f26 = rhs171_
                lhs132_.f15 = rhs172_
                lhs133_.f18 = rhs173_
                d_1069_v48_ = rhs174_
                d_1072_v50_: _dafny.Map
                d_1072_v50_ = _dafny.Map({(d_1067_v46_).f36: p0})
                d_1073_v51_: _dafny.Array
                nw187_ = _dafny.Array(None, 23)
                nw187_[int(0)] = (d_1067_v46_).f37
                nw187_[int(1)] = (d_1067_v46_).f37
                nw187_[int(2)] = ((p2).set(default__.safeIndex((d_1067_v46_).f36, len(p2)), p1)).set(default__.safeIndex((0) - (len(d_1072_v50_)), len((p2).set(default__.safeIndex((d_1067_v46_).f36, len(p2)), p1))), p1)
                nw187_[int(3)] = (d_1067_v46_).f37
                nw187_[int(4)] = (self).fm22(default__.fm24((self).f32, globalState), p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngpnnpqao"))), globalState)
                nw187_[int(5)] = p2
                nw187_[int(6)] = p2
                nw187_[int(7)] = (p2) + (p2)
                nw187_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkphmjmvi"))) + (p2)
                nw187_[int(9)] = p2
                nw187_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmodjtc"))
                nw187_[int(11)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baincv"))) + ((d_1067_v46_).f37)
                nw187_[int(12)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hehm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgcyfxm")))
                nw187_[int(13)] = (p2) + (p2)
                nw187_[int(14)] = (d_1067_v46_).f37
                nw187_[int(15)] = p2
                nw187_[int(16)] = (d_1067_v46_).f37
                nw187_[int(17)] = (d_1067_v46_).f37
                nw187_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwkywvb"))
                nw187_[int(19)] = (d_1067_v46_).f37
                nw187_[int(20)] = p2
                nw187_[int(21)] = p2
                nw187_[int(22)] = p2
                d_1073_v51_ = nw187_
                index128_ = default__.safeIndex(124, (d_1073_v51_).length(0))
                (d_1073_v51_)[index128_] = (d_1067_v46_).f37
                index129_ = default__.safeIndex(124, (d_1073_v51_).length(0))
                (d_1073_v51_)[index129_] = (d_1067_v46_).f37
                d_1074_v53_: _dafny.MultiSet
                def iife78_():
                    coll42_ = _dafny.Set()
                    compr_42_: int
                    for compr_42_ in _dafny.IntegerRange(274, 326):
                        d_1075_v52_: int = compr_42_
                        if ((274) <= (d_1075_v52_)) and ((d_1075_v52_) < (326)):
                            coll42_ = coll42_.union(_dafny.Set([(d_1075_v52_) * (len(d_1003_v1_))]))
                    return _dafny.Set(coll42_)
                d_1074_v53_ = _dafny.MultiSet([iife78_()
                , _dafny.Set({(((self).f31)[p3] if (p3) in ((self).f31) else (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]), (d_1067_v46_).f36, (d_1067_v46_).f36, (d_1067_v46_).f36, (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]})])
                d_1076_v54_: D11
                d_1076_v54_ = D11_DC23(d_1074_v53_)
                d_1072_v50_ = (d_1072_v50_).set(((d_1067_v46_).f36) + (((d_1076_v54_).cf42).cardinality), not(p0))
                d_1077_v55_: _dafny.Array
                nw188_ = _dafny.Array(False, 14)
                d_1077_v55_ = nw188_
                index130_ = default__.safeIndex(63, (d_1077_v55_).length(0))
                (d_1077_v55_)[index130_] = not(True)
                d_1078_v56_: D2
                d_1078_v56_ = D2_DC4((0) - ((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]), (d_1067_v46_).f36, 986, _dafny.CodePoint('y'), (d_1067_v46_).f37)
                d_1079_v57_: _dafny.Map
                d_1079_v57_ = _dafny.Map({d_1078_v56_: len(_dafny.Map({p3: p3}))})
                d_1080_v58_: _dafny.Map
                d_1080_v58_ = _dafny.Map({d_1079_v57_: _dafny.Map({p0: d_1066_i7_})})
                d_1081_v59_: bool
                d_1081_v59_ = True
                d_1082_v60_: _dafny.Map
                d_1082_v60_ = _dafny.Map({(d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]: d_1081_v59_})
                d_1083_v61_: _dafny.Seq
                d_1083_v61_ = _dafny.SeqWithoutIsStrInference([d_1067_v46_])
                d_1084_v62_: _dafny.Map
                d_1084_v62_ = _dafny.Map({d_1082_v60_: _dafny.SeqWithoutIsStrInference([d_1067_v46_, (d_1083_v61_)[default__.safeIndex((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], len(d_1083_v61_))]])})
                d_1085_v63_: _dafny.Seq
                d_1085_v63_ = _dafny.SeqWithoutIsStrInference([d_1084_v62_])
                d_1086_v64_: _dafny.Seq
                d_1086_v64_ = _dafny.SeqWithoutIsStrInference([(d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]])
                d_1087_v65_: _dafny.Map
                d_1087_v65_ = _dafny.Map({len((d_1085_v63_)[default__.safeIndex((self).f32, len(d_1085_v63_))]): _dafny.Map({p0: len(d_1086_v64_)})})
                d_1088_v66_: _dafny.Map
                d_1088_v66_ = _dafny.Map({p0: (self).f32})
                d_1089_v67_: _dafny.Array
                nw189_ = _dafny.Array(None, 12)
                nw189_[int(0)] = d_1080_v58_
                nw189_[int(1)] = d_1080_v58_
                nw189_[int(2)] = (d_1080_v58_) | (d_1080_v58_)
                nw189_[int(3)] = d_1080_v58_
                nw189_[int(4)] = d_1080_v58_
                nw189_[int(5)] = d_1080_v58_
                nw189_[int(6)] = (d_1080_v58_) | (d_1080_v58_)
                nw189_[int(7)] = (_dafny.Map({_dafny.Map({d_1078_v56_: (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]}): (_dafny.Map({p3: (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]})).set(p0, d_1066_i7_)})).set(d_1079_v57_, ((d_1087_v65_)[(d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]] if ((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]) in (d_1087_v65_) else d_1088_v66_))
                nw189_[int(8)] = d_1080_v58_
                nw189_[int(9)] = d_1080_v58_
                nw189_[int(10)] = d_1080_v58_
                nw189_[int(11)] = (d_1080_v58_) | (d_1080_v58_)
                d_1089_v67_ = nw189_
                index131_ = default__.safeIndex(553, (d_1089_v67_).length(0))
                (d_1089_v67_)[index131_] = d_1080_v58_
                d_1090_v68_: D4
                d_1090_v68_ = D4_DC8(p1, 802, False)
                index132_ = default__.safeIndex(124, (d_1073_v51_).length(0))
                index133_ = default__.safeIndex(63, (d_1077_v55_).length(0))
                index134_ = default__.safeIndex(553, (d_1089_v67_).length(0))
                rhs175_ = (D11_DC24(((d_1088_v66_)[p0] if (p0) in (d_1088_v66_) else (d_1067_v46_).f36), d_1090_v68_)).cf43
                rhs176_ = (d_1067_v46_).f37
                rhs177_ = True
                rhs178_ = (d_1067_v46_).f36
                rhs179_ = d_1080_v58_
                lhs134_ = globalState
                lhs135_ = d_1073_v51_
                lhs136_ = default__.safeIndex(124, (d_1073_v51_).length(0))
                lhs137_ = d_1077_v55_
                lhs138_ = default__.safeIndex(63, (d_1077_v55_).length(0))
                lhs139_ = globalState
                lhs140_ = d_1089_v67_
                lhs141_ = default__.safeIndex(553, (d_1089_v67_).length(0))
                lhs134_.f24 = rhs175_
                lhs135_[lhs136_] = rhs176_
                lhs137_[lhs138_] = rhs177_
                lhs139_.f21 = rhs178_
                lhs140_[lhs141_] = rhs179_
            elif True:
                d_1091_v70_: _dafny.Seq
                d_1091_v70_ = _dafny.SeqWithoutIsStrInference([(self).f32])
                def iife79_():
                    coll43_ = _dafny.Map()
                    compr_43_: int
                    for compr_43_ in (d_1091_v70_).Elements:
                        d_1092_v69_: int = compr_43_
                        if (d_1092_v69_) in (d_1091_v70_):
                            coll43_[default__.safeDivisionInt(d_1092_v69_, d_1066_i7_)] = p1
                    return _dafny.Map(coll43_)
                (globalState).f24 = len(iife79_()
                )
                d_1093_v71_: _dafny.Map
                d_1093_v71_ = _dafny.Map({p0: p2})
                d_1094_v72_: D10
                d_1094_v72_ = D10_DC22(len(((d_1093_v71_)[p3] if (p3) in (d_1093_v71_) else p2)), _dafny.MultiSet([p0]))
                (self).m7(d_1003_v1_, (d_1094_v72_).cf40, globalState)
                (globalState).f26 = not(((self).f32) < (default__.fm1(globalState)))
                d_1095_v73_: _dafny.Array
                def lambda40_(d_1096_p0_):
                    def lambda41_(d_1097_i9_):
                        return not (not(True)) or (d_1096_p0_)

                    return lambda41_

                init19_ = lambda40_(p0)
                nw190_ = _dafny.Array(None, 14)
                for i0_19_ in range(nw190_.length(0)):
                    nw190_[i0_19_] = init19_(i0_19_)
                d_1095_v73_ = nw190_
                index135_ = default__.safeIndex(512, (d_1095_v73_).length(0))
                (d_1095_v73_)[index135_] = p3
                d_1098_v74_: _dafny.Map
                d_1098_v74_ = _dafny.Map({p3: (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]})
                d_1099_v75_: _dafny.Seq
                d_1099_v75_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1100_v76_: _dafny.Seq
                d_1100_v76_ = _dafny.SeqWithoutIsStrInference([d_1099_v75_, d_1099_v75_])
                index136_ = default__.safeIndex(512, (d_1095_v73_).length(0))
                (d_1095_v73_)[index136_] = (default__.fm27(len(d_1100_v76_), globalState)).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(d_1098_v74_)])))
                d_1101_v77_: _dafny.Map
                d_1101_v77_ = _dafny.Map({(d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]: p3})
                d_1102_v78_: D1
                d_1102_v78_ = D1_DC1(d_1101_v77_)
                d_1103_v79_: _dafny.Seq
                d_1103_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqpsjmtb"))
                d_1104_v80_: _dafny.Seq
                d_1104_v80_ = _dafny.SeqWithoutIsStrInference([d_1095_v73_, d_1095_v73_])
                d_1105_v81_: _dafny.MultiSet
                d_1105_v81_ = _dafny.MultiSet([d_1066_i7_, (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], d_1066_i7_, (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], (self).f32])
                pat_let_tv14_ = d_1002_v0_
                pat_let_tv15_ = d_1002_v0_
                pat_let_tv16_ = d_1095_v73_
                pat_let_tv17_ = d_1095_v73_
                pat_let_tv18_ = d_1101_v77_
                def iife80_(_pat_let18_0):
                    def iife81_(d_1106_dt__update__tmp_h0_):
                        def iife82_(_pat_let19_0):
                            def iife83_(d_1107_dt__update_hcf1_h0_):
                                return D1_DC1(d_1107_dt__update_hcf1_h0_)
                            return iife83_(_pat_let19_0)
                        return iife82_(_dafny.Map({(pat_let_tv15_)[default__.safeIndex(100, (pat_let_tv14_).length(0))]: (pat_let_tv17_)[default__.safeIndex(512, (pat_let_tv16_).length(0))]}))
                    return iife81_(_pat_let18_0)
                def iife84_(_pat_let20_0):
                    def iife85_(d_1108_dt__update__tmp_h1_):
                        def iife86_(_pat_let21_0):
                            def iife87_(d_1109_dt__update_hcf1_h1_):
                                return D1_DC1(d_1109_dt__update_hcf1_h1_)
                            return iife87_(_pat_let21_0)
                        return iife86_(pat_let_tv18_)
                    return iife85_(_pat_let20_0)
                rhs180_ = (d_1104_v80_)[default__.safeIndex((d_1105_v81_).cardinality, len(d_1104_v80_))]
                rhs181_ = (iife80_(d_1102_v78_) if p0 else iife84_(default__.fm28(_dafny.CodePoint('k'), (0) - ((0) - (d_1066_i7_)), (d_1095_v73_)[default__.safeIndex(512, (d_1095_v73_).length(0))], (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], globalState)))
                rhs182_ = d_1103_v79_
                d_1095_v73_ = rhs180_
                d_1102_v78_ = rhs181_
                d_1103_v79_ = rhs182_
            index137_ = default__.safeIndex(100, (d_1002_v0_).length(0))
            (d_1002_v0_)[index137_] = default__.safeModuloInt(d_1066_i7_, 408)
            (globalState).f26 = p3
            if p0:
                d_1110_v82_: _dafny.Seq
                d_1110_v82_ = _dafny.SeqWithoutIsStrInference([d_1066_i7_])
                (globalState).f18 = ((self).f32) + ((d_1110_v82_)[default__.safeIndex((self).f32, len(d_1110_v82_))])
                (globalState).f3 = d_1066_i7_
                d_1111_v83_: _dafny.Map
                d_1111_v83_ = _dafny.Map({p0: d_1066_i7_})
                d_1112_v84_: _dafny.Set
                d_1112_v84_ = _dafny.Set({p3})
                (globalState).f0 = default__.safeModuloInt((self).f32, default__.safeDivisionInt((0) - (len((self).fm22(d_1111_v83_, p0, len(d_1112_v84_), globalState))), d_1066_i7_))
                (globalState).f26 = not (p3) or (True)
                d_1113_v85_: C0
                nw191_ = C0()
                nw191_.ctor__(default__.safeDivisionInt((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], d_1066_i7_), ((self).fm22(d_1111_v83_, True, (self).f32, globalState)) + ((self).fm22(d_1111_v83_, p0, 306, globalState)))
                d_1113_v85_ = nw191_
            elif True:
                (globalState).f26 = p0
                d_1114_v86_: _dafny.Array
                nw192_ = _dafny.Array(False, 25)
                d_1114_v86_ = nw192_
                index138_ = default__.safeIndex(108, (d_1114_v86_).length(0))
                (d_1114_v86_)[index138_] = p0
                d_1115_v87_: _dafny.Seq
                d_1115_v87_ = _dafny.SeqWithoutIsStrInference([not(p0)])
                index139_ = default__.safeIndex(108, (d_1114_v86_).length(0))
                (d_1114_v86_)[index139_] = (d_1115_v87_)[default__.safeIndex(d_1066_i7_, len(d_1115_v87_))]
                d_1116_v88_: _dafny.MultiSet
                d_1116_v88_ = _dafny.MultiSet([p2])
                (globalState).f8 = d_1116_v88_
                (globalState).f26 = p0
                index140_ = default__.safeIndex(108, (d_1114_v86_).length(0))
                (d_1114_v86_)[index140_] = p0
        d_1117_v89_: str
        d_1117_v89_ = _dafny.CodePoint('g')
        d_1117_v89_ = p1
        _ingredientsd_0 = []
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1002_v0_).length(0)):
            d_1118_i10_: int = guard_loop_2_
            if (True) and (((0) <= (d_1118_i10_)) and ((d_1118_i10_) < ((d_1002_v0_).length(0)))):
                _ingredientsd_0.append((d_1002_v0_, int(d_1118_i10_), (d_1118_i10_) * ((d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))])))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        if not (p0) or (not(p3)):
            (globalState).f15 = (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]
            d_1119_v90_: _dafny.Seq
            d_1119_v90_ = _dafny.SeqWithoutIsStrInference([len(p2), (self).f32])
            d_1120_v91_: _dafny.Array
            nw193_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
            d_1120_v91_ = nw193_
            d_1121_v92_: C1
            nw194_ = C1()
            nw194_.ctor__(p3, (self).f31, d_1119_v90_, (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))], d_1120_v91_)
            d_1121_v92_ = nw194_
            d_1122_v93_: _dafny.Array
            nw195_ = _dafny.Array(None, 5)
            nw195_[int(0)] = d_1121_v92_
            nw195_[int(1)] = d_1121_v92_
            nw195_[int(2)] = d_1121_v92_
            nw195_[int(3)] = d_1121_v92_
            nw195_[int(4)] = d_1121_v92_
            d_1122_v93_ = nw195_
            d_1122_v93_ = d_1122_v93_
            d_1123_v94_: _dafny.Map
            d_1123_v94_ = _dafny.Map({(default__.fm1(globalState)) + ((self).f32): not(p0)})
            d_1123_v94_ = (d_1123_v94_).set(len(d_1123_v94_), not(p0))
            d_1124_v95_: _dafny.MultiSet
            d_1124_v95_ = _dafny.MultiSet([(self).f32])
            d_1125_v96_: D4
            d_1125_v96_ = D4_DC9((d_1124_v95_).cardinality, p3)
            d_1125_v96_ = d_1125_v96_
        elif True:
            index141_ = default__.safeIndex(100, (d_1002_v0_).length(0))
            (d_1002_v0_)[index141_] = (d_1002_v0_)[default__.safeIndex(100, (d_1002_v0_).length(0))]
            (globalState).f26 = p0
            d_1126_v97_: _dafny.Array
            nw196_ = _dafny.Array(None, 3)
            d_1126_v97_ = nw196_
            d_1126_v97_ = d_1126_v97_
            d_1127_v98_: _dafny.Array
            nw197_ = _dafny.Array(False, 11)
            d_1127_v98_ = nw197_
            index142_ = default__.safeIndex(698, (d_1127_v98_).length(0))
            (d_1127_v98_)[index142_] = False
            index143_ = default__.safeIndex(698, (d_1127_v98_).length(0))
            (d_1127_v98_)[index143_] = p0
            d_1128_v99_: _dafny.Set
            d_1128_v99_ = _dafny.Set({p0, (self).fm5(globalState)})
            (globalState).f26 = (d_1128_v99_).issubset(d_1128_v99_)

    def m7(self, p0, p1, globalState):
        (globalState).f3 = (p1) * (p1)
        if (p1) != (p1):
            d_1129_v0_: _dafny.Array
            nw198_ = _dafny.Array(_dafny.CodePoint('D'), 20)
            d_1129_v0_ = nw198_
            d_1130_v1_: str
            d_1130_v1_ = _dafny.CodePoint('d')
            index144_ = default__.safeIndex(857, (d_1129_v0_).length(0))
            (d_1129_v0_)[index144_] = d_1130_v1_
            d_1131_v2_: bool
            d_1131_v2_ = True
            d_1132_v3_: _dafny.Set
            d_1132_v3_ = _dafny.Set({d_1131_v2_})
            index145_ = default__.safeIndex(857, (d_1129_v0_).length(0))
            rhs183_ = default__.safeModuloInt(len((d_1132_v3_ if d_1131_v2_ else d_1132_v3_)), (self).f32)
            rhs184_ = not(d_1131_v2_)
            rhs185_ = d_1130_v1_
            lhs142_ = globalState
            lhs143_ = globalState
            lhs144_ = d_1129_v0_
            lhs145_ = default__.safeIndex(857, (d_1129_v0_).length(0))
            lhs142_.f21 = rhs183_
            lhs143_.f26 = rhs184_
            lhs144_[lhs145_] = rhs185_
            (globalState).f24 = (self).f32
            d_1133_v4_: _dafny.Array
            def lambda42_(d_1134_i0_):
                return (d_1134_i0_) - ((self).f32)

            init20_ = lambda42_
            nw199_ = _dafny.Array(None, 7)
            for i0_20_ in range(nw199_.length(0)):
                nw199_[i0_20_] = init20_(i0_20_)
            d_1133_v4_ = nw199_
            rhs186_ = (default__.fm29(d_1131_v2_, True, d_1131_v2_, globalState))
            rhs187_ = d_1131_v2_
            rhs188_ = d_1133_v4_
            lhs146_ = globalState
            lhs147_ = globalState
            lhs146_.f26 = rhs186_
            lhs147_.f26 = rhs187_
            d_1133_v4_ = rhs188_
            d_1135_v5_: _dafny.Array
            nw200_ = _dafny.Array(False, 13)
            d_1135_v5_ = nw200_
            index146_ = default__.safeIndex(976, (d_1135_v5_).length(0))
            (d_1135_v5_)[index146_] = (self).fm21(globalState)
            d_1136_v6_: _dafny.Seq
            d_1136_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxwtlr"))
            index147_ = default__.safeIndex(976, (d_1135_v5_).length(0))
            (d_1135_v5_)[index147_] = (d_1136_v6_) <= (d_1136_v6_)
            index148_ = default__.safeIndex(539, (d_1133_v4_).length(0))
            (d_1133_v4_)[index148_] = -949
            index149_ = default__.safeIndex(539, (d_1133_v4_).length(0))
            (d_1133_v4_)[index149_] = default__.fm1(globalState)
        elif True:
            d_1137_v7_: bool
            d_1137_v7_ = False
            d_1138_v8_: bool
            d_1139_v9_: int
            out38_: bool
            out39_: int
            out38_, out39_ = default__.m0(not(d_1137_v7_), (0) - ((p1) - (p1)), (self).f32, (p1) + (p1), globalState)
            d_1138_v8_ = out38_
            d_1139_v9_ = out39_
            d_1140_v10_: _dafny.Seq
            d_1140_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmvkrjqv"))
            d_1141_v11_: C0
            nw201_ = C0()
            nw201_.ctor__((d_1139_v9_) * (default__.fm1(globalState)), d_1140_v10_)
            d_1141_v11_ = nw201_
            d_1142_v12_: _dafny.Set
            d_1142_v12_ = _dafny.Set({(0) - (p1)})
            d_1143_v13_: _dafny.MultiSet
            d_1143_v13_ = _dafny.MultiSet([(self).f32])
            d_1144_v14_: _dafny.Map
            d_1144_v14_ = _dafny.Map({d_1137_v7_: d_1143_v13_})
            d_1145_v15_: _dafny.Set
            d_1145_v15_ = _dafny.Set({d_1142_v12_, _dafny.Set({d_1139_v9_, (d_1141_v11_).f36, (0) - (len(d_1144_v14_)), (d_1141_v11_).f36}), _dafny.Set({(self).f32, 322}), d_1142_v12_, (d_1142_v12_) - (d_1142_v12_)})
            d_1146_v16_: _dafny.Seq
            d_1146_v16_ = _dafny.SeqWithoutIsStrInference([d_1138_v8_, d_1137_v7_])
            rhs189_ = d_1145_v15_
            rhs190_ = d_1137_v7_
            rhs191_ = (d_1139_v9_) - (len(d_1146_v16_))
            lhs148_ = globalState
            d_1145_v15_ = rhs189_
            d_1138_v8_ = rhs190_
            lhs148_.f21 = rhs191_
            rhs192_ = d_1146_v16_
            rhs193_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "coh")) if not(d_1137_v7_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbd")))
            d_1146_v16_ = rhs192_
            d_1140_v10_ = rhs193_
            if (d_1137_v7_ if not((d_1146_v16_)[default__.safeIndex((self).f32, len(d_1146_v16_))]) else d_1137_v7_):
                (globalState).f26 = d_1138_v8_
                d_1147_v17_: _dafny.Map
                d_1147_v17_ = _dafny.Map({d_1139_v9_: (self).f32})
                d_1147_v17_ = (d_1147_v17_).set(len((d_1140_v10_) + ((d_1141_v11_).f37)), default__.fm1(globalState))
                d_1148_v18_: _dafny.Seq
                d_1148_v18_ = _dafny.SeqWithoutIsStrInference([(d_1141_v11_).f36])
                d_1149_v19_: _dafny.Map
                d_1149_v19_ = _dafny.Map({d_1138_v8_: d_1148_v18_})
                d_1150_v20_: _dafny.Map
                d_1150_v20_ = _dafny.Map({not(d_1138_v8_): len(d_1149_v19_)})
                rhs194_ = len((d_1150_v20_).set(d_1138_v8_, p1))
                rhs195_ = not(d_1138_v8_)
                rhs196_ = (0) - ((0) - (d_1139_v9_))
                rhs197_ = ((d_1139_v9_) + (p1)) * ((d_1141_v11_).f36)
                lhs149_ = globalState
                lhs150_ = globalState
                lhs151_ = globalState
                lhs152_ = globalState
                lhs149_.f21 = rhs194_
                lhs150_.f26 = rhs195_
                lhs151_.f21 = rhs196_
                lhs152_.f18 = rhs197_
                d_1151_v21_: str
                d_1151_v21_ = _dafny.CodePoint('g')
                d_1152_v22_: C0
                nw202_ = C0()
                nw202_.ctor__(default__.fm1(globalState), ((d_1141_v11_).f37).set(default__.safeIndex(605, len((d_1141_v11_).f37)), d_1151_v21_))
                d_1152_v22_ = nw202_
                (globalState).f26 = (d_1137_v7_ if d_1138_v8_ else d_1137_v7_)
            elif True:
                d_1138_v8_ = d_1138_v8_
                d_1140_v10_ = (d_1141_v11_).f37
                (globalState).f18 = d_1139_v9_
                (globalState).f26 = (((self).f31) | ((self).f31)) != ((self).f31)
                d_1153_v24_: _dafny.Seq
                def iife88_():
                    coll44_ = _dafny.Map()
                    compr_44_: int
                    for compr_44_ in _dafny.IntegerRange(454, 314):
                        d_1154_v23_: int = compr_44_
                        if ((454) <= (d_1154_v23_)) and ((d_1154_v23_) < (314)):
                            coll44_[(d_1154_v23_) + (len(_dafny.Map({d_1137_v7_: d_1139_v9_})))] = (d_1141_v11_).f36
                    return _dafny.Map(coll44_)
                d_1153_v24_ = _dafny.SeqWithoutIsStrInference([iife88_()
                ])
                d_1155_v25_: _dafny.Map
                d_1155_v25_ = _dafny.Map({d_1141_v11_: d_1139_v9_})
                d_1156_v26_: _dafny.Map
                d_1156_v26_ = _dafny.Map({(d_1141_v11_).f36: ((d_1155_v25_)[d_1141_v11_] if (d_1141_v11_) in (d_1155_v25_) else len((d_1141_v11_).f37))})
                d_1157_v27_: _dafny.Seq
                d_1157_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1141_v11_).f36: 63}), ((d_1153_v24_)[default__.safeIndex((self).f32, len(d_1153_v24_))]).set(371, d_1139_v9_), default__.fm30(globalState), d_1156_v26_, d_1156_v26_])
                d_1158_v28_: _dafny.Map
                d_1158_v28_ = _dafny.Map({len((d_1157_v27_)[default__.safeIndex((d_1141_v11_).f36, len(d_1157_v27_))]): not (not(d_1138_v8_)) or ((self).fm5(globalState))})
                d_1158_v28_ = d_1158_v28_
        d_1159_v29_: _dafny.Array
        nw203_ = _dafny.Array(_dafny.CodePoint('D'), 14)
        d_1159_v29_ = nw203_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1159_v29_).length(0)):
            d_1160_i1_: int = guard_loop_3_
            if (True) and (((0) <= (d_1160_i1_)) and ((d_1160_i1_) < ((d_1159_v29_).length(0)))):
                (d_1159_v29_)[(d_1160_i1_)] = (_dafny.CodePoint('u') if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ic"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hksuqeycu"))) else _dafny.CodePoint('r'))
        d_1161_v30_: bool
        d_1161_v30_ = True
        (globalState).f26 = d_1161_v30_
        d_1162_v31_: _dafny.MultiSet
        d_1162_v31_ = _dafny.MultiSet([d_1161_v30_])
        d_1162_v31_ = (self).f31
        d_1163_v32_: _dafny.Seq
        d_1163_v32_ = _dafny.SeqWithoutIsStrInference([False])
        d_1164_v33_: _dafny.MultiSet
        d_1164_v33_ = _dafny.MultiSet([d_1163_v32_])
        d_1165_v34_: _dafny.Array
        nw204_ = _dafny.Array(None, 1)
        nw204_[int(0)] = (d_1164_v33_).cardinality
        d_1165_v34_ = nw204_
        d_1165_v34_ = d_1165_v34_

    def m8(self, p0, p1, globalState):
        d_1166_v0_: _dafny.Array
        nw205_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
        d_1166_v0_ = nw205_
        index150_ = default__.safeIndex(509, (d_1166_v0_).length(0))
        (d_1166_v0_)[index150_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1167_i0_ in range(default__.abs(155))])
        d_1168_v1_: _dafny.Seq
        d_1168_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwtdknsgw"))
        d_1169_v2_: bool
        d_1169_v2_ = False
        d_1170_v3_: _dafny.Map
        d_1170_v3_ = _dafny.Map({len(d_1168_v1_): d_1169_v2_})
        d_1171_v4_: _dafny.Map
        d_1171_v4_ = _dafny.Map({((d_1170_v3_)[p1] if (p1) in (d_1170_v3_) else d_1169_v2_): d_1168_v1_})
        d_1172_v5_: _dafny.Set
        d_1172_v5_ = _dafny.Set({not(d_1169_v2_), d_1169_v2_})
        d_1173_v6_: _dafny.Map
        d_1173_v6_ = _dafny.Map({d_1172_v5_: not(d_1169_v2_)})
        index151_ = default__.safeIndex(700, (d_1166_v0_).length(0))
        (d_1166_v0_)[index151_] = ((d_1171_v4_)[((d_1173_v6_)[d_1172_v5_] if (d_1172_v5_) in (d_1173_v6_) else d_1169_v2_)] if (((d_1173_v6_)[d_1172_v5_] if (d_1172_v5_) in (d_1173_v6_) else d_1169_v2_)) in (d_1171_v4_) else d_1168_v1_)
        d_1174_v7_: _dafny.MultiSet
        d_1174_v7_ = _dafny.MultiSet([d_1169_v2_])
        index152_ = default__.safeIndex(509, (d_1166_v0_).length(0))
        index153_ = default__.safeIndex(700, (d_1166_v0_).length(0))
        rhs198_ = ((self).f32) - (978)
        rhs199_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1175_i1_ in range(default__.abs(228))])
        rhs200_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1176_i2_ in range(default__.abs(547))])
        rhs201_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hucvcbfcp"))
        rhs202_ = d_1174_v7_
        lhs153_ = globalState
        lhs154_ = d_1166_v0_
        lhs155_ = default__.safeIndex(509, (d_1166_v0_).length(0))
        lhs156_ = d_1166_v0_
        lhs157_ = default__.safeIndex(700, (d_1166_v0_).length(0))
        lhs153_.f18 = rhs198_
        lhs154_[lhs155_] = rhs199_
        lhs156_[lhs157_] = rhs200_
        d_1168_v1_ = rhs201_
        d_1174_v7_ = rhs202_
        d_1177_v8_: _dafny.Set
        d_1177_v8_ = _dafny.Set({(self).f32, (self).f32})
        d_1178_v9_: _dafny.Seq
        d_1178_v9_ = _dafny.SeqWithoutIsStrInference([d_1177_v8_, d_1177_v8_])
        d_1179_v10_: _dafny.MultiSet
        d_1179_v10_ = _dafny.MultiSet([p1, p1, p1])
        d_1180_v11_: _dafny.Map
        d_1180_v11_ = _dafny.Map({d_1169_v2_: (self).f32})
        d_1181_v12_: D8
        d_1181_v12_ = D8_DC16((0) - ((0) - (len(d_1178_v9_))), ((d_1179_v10_)[len(d_1180_v11_)] if (len(d_1180_v11_)) in (d_1179_v10_) else (self).f32))
        pat_let_tv19_ = d_1169_v2_
        def lambda43_(source20_):
            if source20_.is_DC16:
                d_1182___mcc_h0_ = source20_.cf29
                d_1183___mcc_h1_ = source20_.cf30
                d_1184_cf30_ = d_1183___mcc_h1_
                d_1185_cf29_ = d_1182___mcc_h0_
                return pat_let_tv19_
            elif True:
                d_1186___mcc_h2_ = source20_.cf28
                d_1187_cf28_ = d_1186___mcc_h2_
                return True

        (globalState).f26 = lambda43_(d_1181_v12_)
        d_1188_v13_: _dafny.Map
        d_1188_v13_ = _dafny.Map({not(d_1169_v2_): d_1169_v2_})
        d_1189_v14_: _dafny.Seq
        d_1189_v14_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1190_v15_: _dafny.Map
        d_1190_v15_ = _dafny.Map({_dafny.MultiSet(d_1189_v14_): False})
        d_1191_v16_: D8
        d_1191_v16_ = D8_DC15(d_1188_v13_)
        d_1192_v17_: _dafny.Array
        nw206_ = _dafny.Array(None, 19)
        nw206_[int(0)] = d_1188_v13_
        nw206_[int(1)] = (d_1188_v13_).set(d_1169_v2_, ((d_1190_v15_)[_dafny.MultiSet(d_1189_v14_)] if (_dafny.MultiSet(d_1189_v14_)) in (d_1190_v15_) else not(d_1169_v2_)))
        nw206_[int(2)] = d_1188_v13_
        nw206_[int(3)] = d_1188_v13_
        nw206_[int(4)] = (d_1188_v13_) | (_dafny.Map({d_1169_v2_: not(d_1169_v2_)}))
        nw206_[int(5)] = d_1188_v13_
        nw206_[int(6)] = _dafny.Map({False: d_1169_v2_})
        nw206_[int(7)] = d_1188_v13_
        nw206_[int(8)] = (d_1188_v13_).set(d_1169_v2_, d_1169_v2_)
        nw206_[int(9)] = d_1188_v13_
        nw206_[int(10)] = d_1188_v13_
        nw206_[int(11)] = (d_1188_v13_) | (default__.fm23(p1, d_1169_v2_, globalState))
        nw206_[int(12)] = (d_1191_v16_).cf28
        nw206_[int(13)] = (_dafny.Map({d_1169_v2_: d_1169_v2_})) | (d_1188_v13_)
        nw206_[int(14)] = d_1188_v13_
        nw206_[int(15)] = d_1188_v13_
        nw206_[int(16)] = (d_1188_v13_) | (d_1188_v13_)
        nw206_[int(17)] = d_1188_v13_
        nw206_[int(18)] = d_1188_v13_
        d_1192_v17_ = nw206_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1192_v17_).length(0)):
            d_1193_i3_: int = guard_loop_4_
            if (True) and (((0) <= (d_1193_i3_)) and ((d_1193_i3_) < ((d_1192_v17_).length(0)))):
                (d_1192_v17_)[(d_1193_i3_)] = (d_1188_v13_) | (d_1188_v13_)
        d_1194_i4_: int
        d_1194_i4_ = 0
        with _dafny.label("4"):
            while ((((d_1174_v7_).set(d_1169_v2_, default__.abs(p1))).intersection((self).f31)).cardinality) == (p1):
                with _dafny.c_label("4"):
                    if (d_1194_i4_) >= (100):
                        raise _dafny.Break("4")
                    d_1194_i4_ = (d_1194_i4_) + (1)
                    d_1195_v18_: C0
                    nw207_ = C0()
                    nw207_.ctor__((self).f32, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1196_i5_ in range(default__.abs(521))]))
                    d_1195_v18_ = nw207_
                    nw208_ = C0()
                    nw208_.ctor__(default__.safeDivisionInt(len(d_1177_v8_), p1), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwjbo")))
                    d_1195_v18_ = nw208_
                    d_1197_v19_: _dafny.Array
                    nw209_ = _dafny.Array(False, 20)
                    d_1197_v19_ = nw209_
                    d_1198_v20_: _dafny.Map
                    d_1198_v20_ = _dafny.Map({d_1197_v19_: (_dafny.MultiSet([d_1174_v7_, (self).f31])).cardinality})
                    d_1199_v21_: str
                    d_1199_v21_ = _dafny.CodePoint('g')
                    d_1200_v22_: D2
                    d_1200_v22_ = D2_DC4((self).f32, len((d_1195_v18_).f37), ((d_1179_v10_)[((d_1198_v20_)[d_1197_v19_] if (d_1197_v19_) in (d_1198_v20_) else p1)] if (((d_1198_v20_)[d_1197_v19_] if (d_1197_v19_) in (d_1198_v20_) else p1)) in (d_1179_v10_) else (self).f32), d_1199_v21_, d_1168_v1_)
                    source21_ = d_1200_v22_
                    if source21_.is_DC4:
                        d_1201___mcc_h3_ = source21_.cf5
                        d_1202___mcc_h4_ = source21_.cf6
                        d_1203___mcc_h5_ = source21_.cf7
                        d_1204___mcc_h6_ = source21_.cf8
                        d_1205___mcc_h7_ = source21_.cf9
                        d_1206_cf9_ = d_1205___mcc_h7_
                        d_1207_cf8_ = d_1204___mcc_h6_
                        d_1208_cf7_ = d_1203___mcc_h5_
                        d_1209_cf6_ = d_1202___mcc_h4_
                        d_1210_cf5_ = d_1201___mcc_h3_
                        d_1211_v23_: _dafny.Seq
                        d_1211_v23_ = _dafny.SeqWithoutIsStrInference([(d_1195_v18_).f37])
                        d_1209_cf6_ = (((_dafny.MultiSet(d_1211_v23_)).cardinality) + ((d_1195_v18_).f36) if d_1169_v2_ else (d_1195_v18_).f36)
                        d_1209_cf6_ = len((_dafny.Set({False, d_1169_v2_})) - ((_dafny.Set({d_1169_v2_})) - (d_1172_v5_)))
                        d_1212_v24_: bool
                        d_1212_v24_ = d_1169_v2_
                        d_1213_v25_: _dafny.Map
                        d_1213_v25_ = _dafny.Map({d_1197_v19_: d_1169_v2_})
                        d_1214_v26_: _dafny.Seq
                        d_1214_v26_ = _dafny.SeqWithoutIsStrInference([(p1) >= (d_1209_cf6_), (d_1169_v2_), ((d_1213_v25_)[d_1197_v19_] if (d_1197_v19_) in (d_1213_v25_) else d_1169_v2_), d_1169_v2_, d_1169_v2_])
                        d_1215_v27_: _dafny.MultiSet
                        d_1215_v27_ = _dafny.MultiSet([d_1197_v19_])
                        d_1216_v28_: _dafny.Array
                        def lambda44_(d_1217_v10_):
                            def lambda45_(d_1218_i6_):
                                return d_1217_v10_

                            return lambda45_

                        init21_ = lambda44_(d_1179_v10_)
                        nw210_ = _dafny.Array(None, 27)
                        for i0_21_ in range(nw210_.length(0)):
                            nw210_[i0_21_] = init21_(i0_21_)
                        d_1216_v28_ = nw210_
                        index154_ = default__.safeIndex(689, (d_1216_v28_).length(0))
                        (d_1216_v28_)[index154_] = (d_1179_v10_ if d_1169_v2_ else d_1179_v10_)
                        index155_ = default__.safeIndex(689, (d_1216_v28_).length(0))
                        rhs203_ = not(d_1169_v2_)
                        rhs204_ = d_1214_v26_
                        rhs205_ = (d_1215_v27_) - (((d_1215_v27_).set(d_1197_v19_, default__.abs((d_1195_v18_).f36))) | (d_1215_v27_))
                        rhs206_ = (_dafny.MultiSet((d_1189_v14_) + (d_1189_v14_))) - ((d_1179_v10_).intersection(d_1179_v10_))
                        rhs207_ = (0) - (((163) * (-469)) - (883))
                        lhs158_ = globalState
                        lhs159_ = d_1216_v28_
                        lhs160_ = default__.safeIndex(689, (d_1216_v28_).length(0))
                        lhs161_ = globalState
                        lhs158_.f26 = rhs203_
                        d_1214_v26_ = rhs204_
                        d_1215_v27_ = rhs205_
                        lhs159_[lhs160_] = rhs206_
                        lhs161_.f24 = rhs207_
                        d_1219_v29_: C0
                        nw211_ = C0()
                        nw211_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlhbrjm"))), ((d_1168_v1_ if d_1169_v2_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjfolha")))).set(default__.safeIndex((d_1195_v18_).f36, len((d_1168_v1_ if d_1169_v2_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjfolha"))))), ((d_1166_v0_)[default__.safeIndex(509, (d_1166_v0_).length(0))])[default__.safeIndex((d_1195_v18_).f36, len((d_1166_v0_)[default__.safeIndex(509, (d_1166_v0_).length(0))]))]))
                        d_1219_v29_ = nw211_
                    elif source21_.is_DC3:
                        d_1220___mcc_h8_ = source21_.cf4
                        d_1221_cf4_ = d_1220___mcc_h8_
                        (globalState).f18 = p1
                        d_1199_v21_ = ((d_1199_v21_ if d_1169_v2_ else d_1199_v21_) if d_1169_v2_ else d_1199_v21_)
                        d_1222_v30_: _dafny.Map
                        d_1222_v30_ = _dafny.Map({d_1169_v2_: d_1188_v13_})
                        (globalState).f24 = len(d_1222_v30_)
                        d_1223_v31_: bool
                        d_1224_v32_: int
                        out40_: bool
                        out41_: int
                        out40_, out41_ = default__.m0(d_1169_v2_, 935, p1, 299, globalState)
                        d_1223_v31_ = out40_
                        d_1224_v32_ = out41_
                    elif True:
                        d_1225___mcc_h9_ = source21_.cf10
                        d_1226_cf10_ = d_1225___mcc_h9_
                        d_1189_v14_ = d_1189_v14_
                        (globalState).f26 = not(d_1169_v2_)
                        rhs208_ = (d_1195_v18_).f36
                        rhs209_ = not (d_1169_v2_) or (d_1169_v2_)
                        rhs210_ = not(((d_1195_v18_).f36) < (((d_1195_v18_).f36) + (len(_dafny.Map({d_1169_v2_: True})))))
                        lhs162_ = globalState
                        lhs163_ = globalState
                        lhs162_.f24 = rhs208_
                        d_1169_v2_ = rhs209_
                        lhs163_.f26 = rhs210_
                        (globalState).f26 = d_1169_v2_
                    d_1227_v33_: _dafny.MultiSet
                    d_1227_v33_ = _dafny.MultiSet([(d_1166_v0_)[default__.safeIndex(509, (d_1166_v0_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlmcu")), (d_1195_v18_).f37])
                    (globalState).f8 = d_1227_v33_
                    d_1228_v34_: _dafny.MultiSet
                    d_1228_v34_ = _dafny.MultiSet([d_1177_v8_, d_1177_v8_])
                    d_1229_v35_: D11
                    d_1229_v35_ = D11_DC23(d_1228_v34_)
                    source22_ = d_1229_v35_
                    if source22_.is_DC24:
                        d_1230___mcc_h10_ = source22_.cf43
                        d_1231___mcc_h11_ = source22_.cf44
                        d_1232_cf44_ = d_1231___mcc_h11_
                        d_1233_cf43_ = d_1230___mcc_h10_
                        (globalState).f26 = d_1169_v2_
                        d_1234_v36_: _dafny.Seq
                        d_1234_v36_ = _dafny.SeqWithoutIsStrInference([d_1169_v2_])
                        (globalState).f3 = (len(d_1234_v36_)) * ((d_1195_v18_).f36)
                        index156_ = default__.safeIndex(844, (d_1197_v19_).length(0))
                        (d_1197_v19_)[index156_] = d_1169_v2_
                        d_1235_v37_: D4
                        d_1235_v37_ = D4_DC9(p1, d_1169_v2_)
                        index157_ = default__.safeIndex(844, (d_1197_v19_).length(0))
                        (d_1197_v19_)[index157_] = ((d_1188_v13_)[not ((d_1235_v37_).cf17) or (True)] if (not ((d_1235_v37_).cf17) or (True)) in (d_1188_v13_) else not(d_1169_v2_))
                        d_1236_v38_: _dafny.Array
                        def lambda46_(d_1237_p1_):
                            def lambda47_(d_1238_i7_):
                                return (d_1238_i7_) + (d_1237_p1_)

                            return lambda47_

                        init22_ = lambda46_(p1)
                        nw212_ = _dafny.Array(None, 29)
                        for i0_22_ in range(nw212_.length(0)):
                            nw212_[i0_22_] = init22_(i0_22_)
                        d_1236_v38_ = nw212_
                        index158_ = default__.safeIndex(733, (d_1236_v38_).length(0))
                        (d_1236_v38_)[index158_] = d_1233_cf43_
                        d_1239_v39_: _dafny.Array
                        nw213_ = _dafny.Array(_dafny.Array(None, 0), 23)
                        d_1239_v39_ = nw213_
                        d_1240_v40_: _dafny.Array
                        nw214_ = _dafny.Array(None, 10)
                        nw214_[int(0)] = d_1195_v18_
                        nw214_[int(1)] = d_1195_v18_
                        nw214_[int(2)] = d_1195_v18_
                        nw214_[int(3)] = d_1195_v18_
                        nw214_[int(4)] = d_1195_v18_
                        nw214_[int(5)] = d_1195_v18_
                        nw214_[int(6)] = d_1195_v18_
                        nw214_[int(7)] = d_1195_v18_
                        nw214_[int(8)] = d_1195_v18_
                        nw214_[int(9)] = d_1195_v18_
                        d_1240_v40_ = nw214_
                        d_1241_v41_: _dafny.Seq
                        d_1241_v41_ = _dafny.SeqWithoutIsStrInference([d_1240_v40_, d_1240_v40_])
                        index159_ = default__.safeIndex(317, (d_1239_v39_).length(0))
                        (d_1239_v39_)[index159_] = (d_1241_v41_)[default__.safeIndex(683, len(d_1241_v41_))]
                        index160_ = default__.safeIndex(733, (d_1236_v38_).length(0))
                        index161_ = default__.safeIndex(317, (d_1239_v39_).length(0))
                        rhs211_ = d_1177_v8_
                        rhs212_ = default__.safeModuloInt(d_1233_cf43_, d_1233_cf43_)
                        rhs213_ = (((d_1197_v19_)[default__.safeIndex(844, (d_1197_v19_).length(0))] if (d_1197_v19_)[default__.safeIndex(844, (d_1197_v19_).length(0))] else (d_1197_v19_)[default__.safeIndex(844, (d_1197_v19_).length(0))])) == (d_1169_v2_)
                        rhs214_ = (self).f32
                        rhs215_ = d_1240_v40_
                        lhs164_ = d_1236_v38_
                        lhs165_ = default__.safeIndex(733, (d_1236_v38_).length(0))
                        lhs166_ = globalState
                        lhs167_ = globalState
                        lhs168_ = d_1239_v39_
                        lhs169_ = default__.safeIndex(317, (d_1239_v39_).length(0))
                        d_1177_v8_ = rhs211_
                        lhs164_[lhs165_] = rhs212_
                        lhs166_.f26 = rhs213_
                        lhs167_.f15 = rhs214_
                        lhs168_[lhs169_] = rhs215_
                    elif source22_.is_DC23:
                        d_1242___mcc_h12_ = source22_.cf42
                        d_1243_cf42_ = d_1242___mcc_h12_
                        d_1168_v1_ = (_dafny.SeqWithoutIsStrInference([d_1199_v21_ for d_1244_i8_ in range(default__.abs(713))])) + ((d_1195_v18_).f37)
                        d_1245_v42_: C1
                        nw215_ = C1()
                        nw215_.ctor__(d_1169_v2_, (self).f31, (_dafny.SeqWithoutIsStrInference([(d_1195_v18_).f36 for d_1246_i9_ in range(default__.abs(9))])) + (d_1189_v14_), (self).f32, d_1166_v0_)
                        d_1245_v42_ = nw215_
                        d_1247_v43_: _dafny.Seq
                        d_1247_v43_ = _dafny.SeqWithoutIsStrInference([d_1245_v42_.f35])
                        (globalState).f26 = (((d_1195_v18_).f36 if True else (d_1195_v18_).f36)) < (len(d_1247_v43_))
                        d_1248_v44_: _dafny.Seq
                        d_1248_v44_ = _dafny.SeqWithoutIsStrInference([d_1247_v43_])
                        d_1249_v45_: D6
                        d_1249_v45_ = D6_DC12(not(d_1245_v42_.f35), d_1248_v44_, ((d_1171_v4_)[d_1245_v42_.f35] if (d_1245_v42_.f35) in (d_1171_v4_) else d_1168_v1_))
                        (d_1245_v42_).f35 = (d_1199_v21_) in ((d_1249_v45_).cf22)
                    elif True:
                        d_1250___mcc_h13_ = source22_.cf45
                        d_1251_cf45_ = d_1250___mcc_h13_
                        (globalState).f26 = ((d_1174_v7_) | ((d_1174_v7_).set(d_1169_v2_, default__.abs((d_1195_v18_).f36)))).isdisjoint(d_1174_v7_)
                        (globalState).f0 = (0) - ((d_1195_v18_).f36)
                        d_1252_v46_: C1
                        nw216_ = C1()
                        nw216_.ctor__(d_1169_v2_, (self).f31, _dafny.SeqWithoutIsStrInference([(d_1195_v18_).f36 for d_1253_i10_ in range(default__.abs(-32))]), len(d_1172_v5_), d_1166_v0_)
                        d_1252_v46_ = nw216_
                        d_1254_v47_: _dafny.Set
                        d_1254_v47_ = _dafny.Set({d_1252_v46_})
                        d_1255_v48_: _dafny.Array
                        nw217_ = _dafny.Array(None, 18)
                        nw217_[int(0)] = (self).f32
                        nw217_[int(1)] = (d_1195_v18_).f36
                        nw217_[int(2)] = (self).f32
                        nw217_[int(3)] = (d_1195_v18_).f36
                        nw217_[int(4)] = len((d_1198_v20_).set(d_1197_v19_, p1))
                        nw217_[int(5)] = (d_1195_v18_).f36
                        nw217_[int(6)] = p1
                        nw217_[int(7)] = p1
                        nw217_[int(8)] = (d_1195_v18_).f36
                        nw217_[int(9)] = ((d_1195_v18_).f36) - ((0) - (p1))
                        nw217_[int(10)] = 204
                        nw217_[int(11)] = p1
                        nw217_[int(12)] = (d_1195_v18_).f36
                        nw217_[int(13)] = len(d_1168_v1_)
                        nw217_[int(14)] = (0) - (len(d_1254_v47_))
                        nw217_[int(15)] = (0) - (len((d_1168_v1_).set(default__.safeIndex(p1, len(d_1168_v1_)), d_1199_v21_)))
                        nw217_[int(16)] = (d_1195_v18_).f36
                        nw217_[int(17)] = (d_1195_v18_).f36
                        d_1255_v48_ = nw217_
                        index162_ = default__.safeIndex(703, (d_1255_v48_).length(0))
                        (d_1255_v48_)[index162_] = (0) - ((d_1195_v18_).f36)
                        pat_let_tv20_ = d_1199_v21_
                        pat_let_tv21_ = d_1195_v18_
                        pat_let_tv22_ = globalState
                        index163_ = default__.safeIndex(703, (d_1255_v48_).length(0))
                        def iife89_(_pat_let22_0):
                            def iife90_(d_1256_dt__update__tmp_h1_):
                                def iife91_(_pat_let23_0):
                                    def iife92_(d_1257_dt__update_hcf8_h0_):
                                        return D2_DC4((d_1256_dt__update__tmp_h1_).cf5, (d_1256_dt__update__tmp_h1_).cf6, (d_1256_dt__update__tmp_h1_).cf7, d_1257_dt__update_hcf8_h0_, (d_1256_dt__update__tmp_h1_).cf9)
                                    return iife92_(_pat_let23_0)
                                return iife91_(default__.fm26(pat_let_tv20_, 677, (pat_let_tv21_).f36, pat_let_tv22_))
                            return iife90_(_pat_let22_0)
                        (d_1255_v48_)[index163_] = (iife89_(d_1200_v22_)).cf7
                        d_1258_v49_: _dafny.Seq
                        d_1258_v49_ = _dafny.SeqWithoutIsStrInference([(d_1255_v48_ if ((d_1188_v13_)[d_1169_v2_] if (d_1169_v2_) in (d_1188_v13_) else d_1169_v2_) else d_1255_v48_)])
                        d_1259_v50_: D10
                        d_1259_v50_ = D10_DC21(d_1189_v14_)
                        d_1255_v48_ = (d_1258_v49_)[default__.safeIndex(len((d_1189_v14_) + ((d_1259_v50_).cf39)), len(d_1258_v49_))]
                    pass
            pass
        d_1260_v51_: _dafny.MultiSet
        d_1260_v51_ = _dafny.MultiSet([_dafny.Set({(self).f32})])
        d_1260_v51_ = d_1260_v51_
        d_1169_v2_ = d_1169_v2_


class C3(T0):
    def  __init__(self):
        self._f27: int = int(0)
        self._f28: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f27, f28):
        (self)._f27 = f27
        (self)._f28 = f28

    def fm2(self, p0, p1, p2, globalState):
        return _dafny.Map({True: (self).f27})

    def fm45(self, globalState):
        return ((0) - ((self).f27)) - ((self).f27)

    def fm46(self, p0, p1, globalState):
        return ((not(False)) or (False)) == (not((True if (D16_DC40(False, 418, (0) - ((self).f27))).cf74 else True)))


class C4(T0):
    def  __init__(self):
        self._f27: int = int(0)
        self._f28: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f27, f28):
        (self)._f27 = f27
        (self)._f28 = f28

    def fm2(self, p0, p1, p2, globalState):
        if (False) not in (_dafny.SeqWithoutIsStrInference([False, True])):
            return _dafny.Map({True: (self).f27})
        elif True:
            return _dafny.Map({False: (self).f27})

    def fm34(self, p0, p1, globalState):
        return ((self).f27) - ((self).f27)

    def fm35(self, p0, p1, globalState):
        return not(True)

    def m11(self, p0, p1, globalState):
        r0: int = int(0)
        d_1261_v0_: _dafny.Map
        d_1261_v0_ = _dafny.Map({p1: p1})
        d_1262_v1_: _dafny.Seq
        d_1262_v1_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        d_1263_v2_: _dafny.Seq
        d_1263_v2_ = _dafny.SeqWithoutIsStrInference([d_1262_v1_])
        d_1264_v3_: D6
        d_1264_v3_ = D6_DC12(p1, d_1263_v2_, (p0).f37)
        d_1265_v4_: str
        d_1265_v4_ = _dafny.CodePoint('k')
        (globalState).f18 = len(((d_1264_v3_).cf22 if ((d_1261_v0_)[p1] if (p1) in (d_1261_v0_) else p1) else (((p0).f37) + ((p0).f37)).set(default__.safeIndex((self).f27, len(((p0).f37) + ((p0).f37))), d_1265_v4_)))
        d_1266_i0_: int
        d_1266_i0_ = 0
        with _dafny.label("5"):
            while not(not ((p1) and (p1)) or (p1)):
                with _dafny.c_label("5"):
                    if (d_1266_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_1266_i0_ = (d_1266_i0_) + (1)
                    d_1267_v5_: _dafny.Seq
                    d_1267_v5_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({p1})), 277])
                    d_1268_v6_: _dafny.Set
                    d_1268_v6_ = _dafny.Set({-76})
                    d_1269_v7_: D2
                    d_1269_v7_ = D2_DC4((self).f27, (p0).f36, (p0).f36, d_1265_v4_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njlrvlaga")))
                    d_1270_v8_: C1
                    nw218_ = C1()
                    nw218_.ctor__(p1, _dafny.MultiSet(d_1262_v1_), (d_1267_v5_) + (d_1267_v5_), default__.safeDivisionInt((self).fm34(d_1268_v6_, d_1269_v7_, globalState), len(d_1268_v6_)), (self).f28)
                    d_1270_v8_ = nw218_
                    d_1271_v9_: _dafny.Array
                    def lambda48_(d_1272_p1_, d_1273_v8_):
                        def lambda49_(d_1274_i1_):
                            return _dafny.SeqWithoutIsStrInference([d_1272_p1_, d_1273_v8_.f35])

                        return lambda49_

                    init23_ = lambda48_(p1, d_1270_v8_)
                    nw219_ = _dafny.Array(None, 8)
                    for i0_23_ in range(nw219_.length(0)):
                        nw219_[i0_23_] = init23_(i0_23_)
                    d_1271_v9_ = nw219_
                    index164_ = default__.safeIndex(593, (d_1271_v9_).length(0))
                    (d_1271_v9_)[index164_] = (d_1263_v2_)[default__.safeIndex(len(d_1268_v6_), len(d_1263_v2_))]
                    d_1275_v10_: _dafny.MultiSet
                    d_1275_v10_ = _dafny.MultiSet([p1])
                    pat_let_tv23_ = d_1270_v8_
                    index165_ = default__.safeIndex(593, (d_1271_v9_).length(0))
                    def iife93_(_pat_let24_0):
                        def iife94_(d_1276_dt__update__tmp_h0_):
                            def iife95_(_pat_let25_0):
                                def iife96_(d_1277_dt__update_hcf55_h0_):
                                    return D13_DC30(d_1277_dt__update_hcf55_h0_)
                                return iife96_(_pat_let25_0)
                            return iife95_(pat_let_tv23_)
                        return iife94_(_pat_let24_0)
                    rhs216_ = (iife93_(D13_DC30(d_1270_v8_))).cf55
                    rhs217_ = True
                    rhs218_ = (((d_1262_v1_).set(default__.safeIndex((p0).f36, len(d_1262_v1_)), d_1270_v8_.f35)).set(default__.safeIndex((p0).f36, len((d_1262_v1_).set(default__.safeIndex((p0).f36, len(d_1262_v1_)), d_1270_v8_.f35))), p1)).set(default__.safeIndex(default__.safeModuloInt((p0).f36, 688), len(((d_1262_v1_).set(default__.safeIndex((p0).f36, len(d_1262_v1_)), d_1270_v8_.f35)).set(default__.safeIndex((p0).f36, len((d_1262_v1_).set(default__.safeIndex((p0).f36, len(d_1262_v1_)), d_1270_v8_.f35))), p1))), (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1, (d_1270_v8_).fm3(d_1275_v10_, globalState)]), d_1262_v1_, d_1262_v1_])) == (d_1263_v2_))
                    lhs170_ = globalState
                    lhs171_ = d_1271_v9_
                    lhs172_ = default__.safeIndex(593, (d_1271_v9_).length(0))
                    d_1270_v8_ = rhs216_
                    lhs170_.f26 = rhs217_
                    lhs171_[lhs172_] = rhs218_
                    (globalState).f3 = (p0).f36
                    (globalState).f26 = (p1) and (not (d_1270_v8_.f35) or ((d_1262_v1_)[default__.safeIndex((self).f27, len(d_1262_v1_))]))
                    if p1:
                        d_1278_v11_: _dafny.Seq
                        d_1278_v11_ = _dafny.SeqWithoutIsStrInference([d_1269_v7_, default__.fm36(not(p1), globalState)])
                        pat_let_tv24_ = p0
                        pat_let_tv25_ = globalState
                        def iife97_(_pat_let26_0):
                            def iife98_(d_1279_dt__update__tmp_h1_):
                                def iife99_(_pat_let27_0):
                                    def iife100_(d_1280_dt__update_hcf6_h0_):
                                        def iife101_(_pat_let28_0):
                                            def iife102_(d_1281_dt__update_hcf5_h0_):
                                                return D2_DC4(d_1281_dt__update_hcf5_h0_, d_1280_dt__update_hcf6_h0_, (d_1279_dt__update__tmp_h1_).cf7, (d_1279_dt__update__tmp_h1_).cf8, (d_1279_dt__update__tmp_h1_).cf9)
                                            return iife102_(_pat_let28_0)
                                        return iife101_(default__.fm1(pat_let_tv25_))
                                    return iife100_(_pat_let27_0)
                                return iife99_((pat_let_tv24_).f36)
                            return iife98_(_pat_let26_0)
                        d_1278_v11_ = _dafny.SeqWithoutIsStrInference([iife97_(d_1269_v7_), d_1269_v7_, d_1269_v7_])
                        d_1282_v12_: _dafny.Set
                        d_1282_v12_ = _dafny.Set({d_1270_v8_.f35})
                        d_1283_v13_: T2
                        nw220_ = C2()
                        nw220_.ctor__(d_1275_v10_, (p0).f36)
                        d_1283_v13_ = nw220_
                        d_1284_v14_: _dafny.Map
                        d_1284_v14_ = _dafny.Map({d_1270_v8_: d_1283_v13_})
                        d_1285_v15_: _dafny.Seq
                        d_1285_v15_ = _dafny.SeqWithoutIsStrInference([(p0).f37])
                        d_1286_v16_: _dafny.Array
                        nw221_ = _dafny.Array(None, 11)
                        nw221_[int(0)] = len((p0).f37)
                        nw221_[int(1)] = (p0).f36
                        nw221_[int(2)] = len((d_1271_v9_)[default__.safeIndex(593, (d_1271_v9_).length(0))])
                        nw221_[int(3)] = len((p0).f37)
                        nw221_[int(4)] = (self).f27
                        nw221_[int(5)] = (p0).f36
                        nw221_[int(6)] = (d_1283_v13_).f32
                        nw221_[int(7)] = (p0).f36
                        nw221_[int(8)] = len(_dafny.SeqWithoutIsStrInference([d_1265_v4_ for d_1287_i2_ in range(default__.abs(557))]))
                        nw221_[int(9)] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([((p0).f37)[default__.safeIndex((self).f27, len((p0).f37))] for d_1288_i4_ in range(default__.abs(106))])) for d_1289_i3_ in range(default__.abs(693))]))
                        nw221_[int(10)] = (p0).f36
                        d_1286_v16_ = nw221_
                        d_1290_v17_: _dafny.Map
                        d_1290_v17_ = _dafny.Map({d_1262_v1_: d_1286_v16_})
                        d_1291_v18_: _dafny.Seq
                        d_1291_v18_ = _dafny.SeqWithoutIsStrInference([((d_1290_v17_)[(d_1271_v9_)[default__.safeIndex(593, (d_1271_v9_).length(0))]] if ((d_1271_v9_)[default__.safeIndex(593, (d_1271_v9_).length(0))]) in (d_1290_v17_) else d_1286_v16_), d_1286_v16_, d_1286_v16_])
                        d_1292_v19_: _dafny.Map
                        d_1292_v19_ = _dafny.Map({(p0).f36: (d_1283_v13_).f32})
                        d_1293_v20_: _dafny.Map
                        d_1293_v20_ = _dafny.Map({d_1270_v8_.f35: len(d_1292_v19_)})
                        d_1294_v21_: _dafny.MultiSet
                        d_1294_v21_ = _dafny.MultiSet([(d_1291_v18_)[default__.safeIndex(len(d_1293_v20_), len(d_1291_v18_))], d_1286_v16_])
                        d_1295_v22_: _dafny.MultiSet
                        d_1295_v22_ = _dafny.MultiSet([_dafny.MultiSet([d_1286_v16_, d_1286_v16_, d_1286_v16_]), d_1294_v21_, d_1294_v21_])
                        d_1296_v23_: _dafny.Map
                        d_1296_v23_ = _dafny.Map({default__.fm37((p0).f36, (self).fm34(d_1268_v6_, D2_DC4(len(d_1282_v12_), len(d_1284_v14_), 662, d_1265_v4_, (d_1285_v15_)[default__.safeIndex((p0).f36, len(d_1285_v15_))]), globalState), globalState): ((d_1295_v22_)[_dafny.MultiSet([d_1286_v16_])] if (_dafny.MultiSet([d_1286_v16_])) in (d_1295_v22_) else (0) - ((self).f27))})
                        (globalState).f24 = (0) - (((d_1296_v23_)[(p0).f37] if ((p0).f37) in (d_1296_v23_) else (p0).f36))
                        d_1297_v24_: _dafny.Seq
                        d_1297_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqatqn"))
                        d_1297_v24_ = (((p0).f37) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqpon"))).set(default__.safeIndex((d_1283_v13_).f32, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqpon")))), d_1265_v4_))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qq")))
                        r0 = 695
                        (globalState).f26 = d_1270_v8_.f35
                    elif True:
                        d_1298_v25_: _dafny.Seq
                        d_1298_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjg"))
                        d_1299_v26_: _dafny.Seq
                        d_1299_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1265_v4_, _dafny.CodePoint('a')])])
                        d_1300_v27_: _dafny.Map
                        d_1300_v27_ = _dafny.Map({d_1265_v4_: 985})
                        d_1301_v28_: _dafny.MultiSet
                        d_1301_v28_ = _dafny.MultiSet([len(d_1300_v27_)])
                        d_1302_v29_: D9
                        d_1302_v29_ = D9_DC19(p1)
                        d_1303_v30_: _dafny.Array
                        nw222_ = _dafny.Array(None, 17)
                        nw222_[int(0)] = d_1270_v8_.f35
                        nw222_[int(1)] = (False if p1 else d_1270_v8_.f35)
                        nw222_[int(2)] = True
                        nw222_[int(3)] = False
                        nw222_[int(4)] = d_1270_v8_.f35
                        nw222_[int(5)] = (_dafny.SeqWithoutIsStrInference([d_1265_v4_])) in ((d_1299_v26_).set(default__.safeIndex((self).f27, len(d_1299_v26_)), (p0).f37))
                        nw222_[int(6)] = ((self).fm35(d_1265_v4_, default__.fm1(globalState), globalState) if p1 else d_1270_v8_.f35)
                        nw222_[int(7)] = True
                        nw222_[int(8)] = ((p0).f36) != ((p0).f36)
                        nw222_[int(9)] = (d_1301_v28_).ispropersubset(_dafny.MultiSet(d_1267_v5_))
                        nw222_[int(10)] = ((p0).f36) >= ((p0).f36)
                        nw222_[int(11)] = (d_1275_v10_) == (_dafny.MultiSet([d_1270_v8_.f35, True]))
                        nw222_[int(12)] = not(d_1270_v8_.f35)
                        nw222_[int(13)] = p1
                        nw222_[int(14)] = d_1270_v8_.f35
                        nw222_[int(15)] = (d_1302_v29_).cf33
                        nw222_[int(16)] = False
                        d_1303_v30_ = nw222_
                        index166_ = default__.safeIndex(885, (d_1303_v30_).length(0))
                        (d_1303_v30_)[index166_] = not(p1)
                        index167_ = default__.safeIndex(885, (d_1303_v30_).length(0))
                        rhs219_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpkm"))
                        rhs220_ = not(not (p1) or (d_1270_v8_.f35))
                        lhs173_ = d_1303_v30_
                        lhs174_ = default__.safeIndex(885, (d_1303_v30_).length(0))
                        d_1298_v25_ = rhs219_
                        lhs173_[lhs174_] = rhs220_
                        d_1304_v31_: C0
                        nw223_ = C0()
                        nw223_.ctor__(len(d_1298_v25_), (p0).f37)
                        d_1304_v31_ = nw223_
                        (globalState).f0 = default__.fm1(globalState)
                        d_1305_v32_: D4
                        d_1305_v32_ = D4_DC8(d_1265_v4_, 444, p1)
                        d_1306_v33_: D11
                        d_1306_v33_ = D11_DC24((d_1304_v31_).f36, d_1305_v32_)
                        d_1306_v33_ = d_1306_v33_
                        d_1307_v34_: D13
                        d_1307_v34_ = D13_DC31((p0).f36)
                        d_1308_v35_: D13
                        d_1308_v35_ = D13_DC33(d_1307_v34_)
                        d_1308_v35_ = d_1308_v35_
                    pass
            pass
        (globalState).f3 = (self).f27
        (globalState).f21 = ((p0).f36) - (-794)
        rhs221_ = not(p1)
        rhs222_ = p1
        rhs223_ = not(p1)
        lhs175_ = globalState
        lhs176_ = globalState
        lhs177_ = globalState
        lhs175_.f26 = rhs221_
        lhs176_.f26 = rhs222_
        lhs177_.f26 = rhs223_
        d_1309_v36_: _dafny.Map
        d_1309_v36_ = _dafny.Map({p1: (self).f27})
        d_1310_v37_: _dafny.Set
        d_1310_v37_ = _dafny.Set({d_1309_v36_, (d_1309_v36_).set(True, 677), (d_1309_v36_).set(default__.fm0(len(d_1262_v1_), globalState), (self).f27), d_1309_v36_})
        hi5_ = len((d_1310_v37_) - (d_1310_v37_))
        for d_1311_i5_ in range(default__.safeModuloInt((self).f27, (self).f27), hi5_):
            if not(((p0).f36) == (431)):
                d_1312_v38_: _dafny.Map
                d_1312_v38_ = _dafny.Map({(p0).f36: d_1262_v1_})
                d_1312_v38_ = (d_1312_v38_).set(((self).f27) - (d_1311_i5_), _dafny.SeqWithoutIsStrInference([p1]))
                d_1313_v39_: _dafny.Map
                d_1313_v39_ = _dafny.Map({(p0).f36: default__.fm0((p0).f36, globalState)})
                d_1314_v40_: _dafny.Map
                d_1314_v40_ = _dafny.Map({p1: d_1313_v39_})
                d_1315_v41_: _dafny.Array
                nw224_ = _dafny.Array(None, 4)
                nw224_[int(0)] = d_1314_v40_
                nw224_[int(1)] = (d_1314_v40_) | (d_1314_v40_)
                nw224_[int(2)] = d_1314_v40_
                nw224_[int(3)] = (d_1314_v40_).set(p1, _dafny.Map({(0) - (d_1311_i5_): False}))
                d_1315_v41_ = nw224_
                index168_ = default__.safeIndex(535, (d_1315_v41_).length(0))
                (d_1315_v41_)[index168_] = d_1314_v40_
                index169_ = default__.safeIndex(535, (d_1315_v41_).length(0))
                (d_1315_v41_)[index169_] = ((d_1314_v40_) | (d_1314_v40_)) | (d_1314_v40_)
                d_1316_v42_: _dafny.MultiSet
                d_1316_v42_ = _dafny.MultiSet([422])
                d_1317_v43_: C0
                nw225_ = C0()
                nw225_.ctor__(((0) - ((p0).f36) if p1 else (d_1316_v42_).cardinality), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1318_i6_ in range(default__.abs(46))]))
                d_1317_v43_ = nw225_
                d_1319_v44_: _dafny.Array
                nw226_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_1319_v44_ = nw226_
                nw227_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_1319_v44_ = nw227_
                d_1320_v45_: _dafny.Array
                nw228_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                d_1320_v45_ = nw228_
                index170_ = default__.safeIndex(416, (d_1320_v45_).length(0))
                (d_1320_v45_)[index170_] = _dafny.CodePoint('c')
                index171_ = default__.safeIndex(416, (d_1320_v45_).length(0))
                (d_1320_v45_)[index171_] = default__.fm38(_dafny.Map({True: False}), globalState)
            elif True:
                d_1321_v46_: _dafny.Array
                def lambda50_(d_1322_v4_, d_1323_i5_, d_1324_p1_):
                    def lambda51_(d_1325_i7_):
                        return D4_DC8(d_1322_v4_, d_1323_i5_, d_1324_p1_)

                    return lambda51_

                init24_ = lambda50_(d_1265_v4_, d_1311_i5_, p1)
                nw229_ = _dafny.Array(None, 4)
                for i0_24_ in range(nw229_.length(0)):
                    nw229_[i0_24_] = init24_(i0_24_)
                d_1321_v46_ = nw229_
                d_1326_v47_: _dafny.Array
                nw230_ = _dafny.Array(D10.default()(), 26)
                d_1326_v47_ = nw230_
                d_1327_v48_: _dafny.Seq
                d_1327_v48_ = _dafny.SeqWithoutIsStrInference([320])
                d_1328_v49_: D10
                d_1328_v49_ = D10_DC21(d_1327_v48_)
                index172_ = default__.safeIndex(315, (d_1326_v47_).length(0))
                (d_1326_v47_)[index172_] = d_1328_v49_
                index173_ = default__.safeIndex(315, (d_1326_v47_).length(0))
                rhs224_ = (default__.safeModuloInt(691, (p0).f36)) - ((p0).f36)
                rhs225_ = p1
                rhs226_ = d_1261_v0_
                rhs227_ = d_1321_v46_
                rhs228_ = d_1328_v49_
                lhs178_ = globalState
                lhs179_ = globalState
                lhs180_ = d_1326_v47_
                lhs181_ = default__.safeIndex(315, (d_1326_v47_).length(0))
                lhs178_.f21 = rhs224_
                lhs179_.f26 = rhs225_
                d_1261_v0_ = rhs226_
                d_1321_v46_ = rhs227_
                lhs180_[lhs181_] = rhs228_
                d_1329_v51_: _dafny.Set
                d_1329_v51_ = _dafny.Set({(p0).f36})
                d_1330_v52_: _dafny.Map
                d_1330_v52_ = _dafny.Map({p1: p0})
                d_1331_v53_: D2
                d_1331_v53_ = D2_DC4((0) - ((self).f27), len(d_1330_v52_), -286, d_1265_v4_, (p0).f37)
                d_1332_v54_: _dafny.Map
                d_1332_v54_ = _dafny.Map({(p0).f36: p1})
                d_1333_v55_: _dafny.MultiSet
                d_1333_v55_ = _dafny.MultiSet([default__.fm0((self).f27, globalState)])
                d_1334_v56_: _dafny.Array
                nw231_ = _dafny.Array(None, 28)
                nw231_[int(0)] = default__.safeDivisionInt(d_1311_i5_, d_1311_i5_)
                nw231_[int(1)] = (0) - (d_1311_i5_)
                nw231_[int(2)] = (d_1311_i5_) * ((p0).f36)
                nw231_[int(3)] = d_1311_i5_
                nw231_[int(4)] = (d_1327_v48_)[default__.safeIndex(d_1311_i5_, len(d_1327_v48_))]
                nw231_[int(5)] = len(d_1262_v1_)
                nw231_[int(6)] = (p0).f36
                def iife103_():
                    coll45_ = _dafny.Set()
                    compr_45_: int
                    for compr_45_ in _dafny.IntegerRange(865, 146):
                        d_1335_v50_: int = compr_45_
                        if ((865) <= (d_1335_v50_)) and ((d_1335_v50_) < (146)):
                            coll45_ = coll45_.union(_dafny.Set([(d_1335_v50_) * ((p0).f36)]))
                    return _dafny.Set(coll45_)
                nw231_[int(7)] = default__.safeDivisionInt((p0).f36, len(iife103_()
                ))
                nw231_[int(8)] = (0) - ((self).f27)
                nw231_[int(9)] = (p0).f36
                nw231_[int(10)] = (self).f27
                nw231_[int(11)] = (p0).f36
                nw231_[int(12)] = ((self).fm34(d_1329_v51_, d_1331_v53_, globalState)) - ((0) - (len(d_1332_v54_)))
                nw231_[int(13)] = (p0).f36
                nw231_[int(14)] = ((d_1309_v36_)[p1] if (p1) in (d_1309_v36_) else (self).f27)
                nw231_[int(15)] = ((d_1333_v55_)[p1] if (p1) in (d_1333_v55_) else (p0).f36)
                nw231_[int(16)] = (0) - ((p0).f36)
                nw231_[int(17)] = d_1311_i5_
                nw231_[int(18)] = (self).f27
                nw231_[int(19)] = (p0).f36
                nw231_[int(20)] = (p0).f36
                nw231_[int(21)] = len(_dafny.SeqWithoutIsStrInference([d_1265_v4_ for d_1336_i8_ in range(default__.abs(300))]))
                nw231_[int(22)] = ((p0).f36) + ((p0).f36)
                nw231_[int(23)] = d_1311_i5_
                nw231_[int(24)] = default__.safeModuloInt(594, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1337_i9_ in range(default__.abs(884))])))
                nw231_[int(25)] = (self).f27
                nw231_[int(26)] = d_1311_i5_
                nw231_[int(27)] = (p0).f36
                d_1334_v56_ = nw231_
                index174_ = default__.safeIndex(55, (d_1334_v56_).length(0))
                (d_1334_v56_)[index174_] = default__.safeDivisionInt((p0).f36, len(d_1327_v48_))
                index175_ = default__.safeIndex(55, (d_1334_v56_).length(0))
                (d_1334_v56_)[index175_] = d_1311_i5_
                (globalState).f3 = d_1311_i5_
                d_1338_v57_: _dafny.Array
                nw232_ = _dafny.Array(False, 29)
                d_1338_v57_ = nw232_
                d_1338_v57_ = (d_1338_v57_ if (_dafny.MultiSet([p1, p1, not(default__.fm0(736, globalState))])).issubset(d_1333_v55_) else d_1338_v57_)
                d_1339_v58_: _dafny.Seq
                d_1339_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efdl"))
                d_1339_v58_ = default__.fm37((self).f27, d_1311_i5_, globalState)
            (globalState).f26 = not(p1)
            d_1340_v59_: _dafny.Array
            nw233_ = _dafny.Array(D13.default()(), 20)
            d_1340_v59_ = nw233_
            d_1341_v60_: _dafny.MultiSet
            d_1341_v60_ = _dafny.MultiSet([(0) - (len(d_1262_v1_)), (self).f27])
            d_1342_v61_: _dafny.MultiSet
            d_1342_v61_ = _dafny.MultiSet([p1, p1])
            d_1343_v62_: D13
            d_1343_v62_ = D13_DC32(((d_1341_v60_)[(p0).f36] if ((p0).f36) in (d_1341_v60_) else d_1311_i5_), (d_1342_v61_).cardinality, _dafny.SeqWithoutIsStrInference([d_1311_i5_, (0) - (len(d_1261_v0_))]), p1)
            index176_ = default__.safeIndex(162, (d_1340_v59_).length(0))
            (d_1340_v59_)[index176_] = (d_1343_v62_ if p1 else d_1343_v62_)
            d_1344_v63_: _dafny.Set
            d_1344_v63_ = _dafny.Set({d_1262_v1_})
            d_1345_v64_: _dafny.Map
            d_1345_v64_ = _dafny.Map({(self).f27: d_1344_v63_})
            d_1346_v65_: _dafny.Set
            d_1346_v65_ = _dafny.Set({d_1262_v1_})
            index177_ = default__.safeIndex(162, (d_1340_v59_).length(0))
            rhs229_ = ((p0).f36) - ((p0).f36)
            rhs230_ = (p0).f36
            rhs231_ = ((d_1346_v65_)).issubset(((d_1345_v64_)[(p0).f36] if ((p0).f36) in (d_1345_v64_) else _dafny.Set({d_1262_v1_})))
            rhs232_ = d_1343_v62_
            lhs182_ = globalState
            lhs183_ = globalState
            lhs184_ = globalState
            lhs185_ = d_1340_v59_
            lhs186_ = default__.safeIndex(162, (d_1340_v59_).length(0))
            lhs182_.f0 = rhs229_
            lhs183_.f21 = rhs230_
            lhs184_.f26 = rhs231_
            lhs185_[lhs186_] = rhs232_
            d_1265_v4_ = d_1265_v4_
        r0 = (self).f27
        return r0

    def m12(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        if not(p2):
            d_1347_v0_: _dafny.Seq
            d_1347_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_1348_v1_: _dafny.Seq
            d_1348_v1_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qppgnht"))) + (d_1347_v0_), default__.fm37((self).f27, len(d_1347_v0_), globalState)])
            d_1349_v2_: _dafny.Map
            d_1349_v2_ = _dafny.Map({True: p0})
            d_1350_v3_: _dafny.MultiSet
            d_1350_v3_ = _dafny.MultiSet([p3])
            d_1351_v4_: _dafny.Seq
            d_1351_v4_ = _dafny.SeqWithoutIsStrInference([p0, p2, False, p0])
            d_1352_v5_: _dafny.Seq
            d_1352_v5_ = _dafny.SeqWithoutIsStrInference([d_1351_v4_])
            d_1353_v6_: D6
            d_1353_v6_ = D6_DC12(p0, d_1352_v5_, d_1347_v0_)
            rhs233_ = d_1348_v1_
            rhs234_ = _dafny.Map({p0: (d_1353_v6_).cf20})
            rhs235_ = d_1350_v3_
            rhs236_ = default__.safeDivisionInt((self).f27, (D4_DC8(p3, len(_dafny.Set({p2})), p2)).cf14)
            lhs187_ = globalState
            d_1348_v1_ = rhs233_
            d_1349_v2_ = rhs234_
            d_1350_v3_ = rhs235_
            lhs187_.f0 = rhs236_
            d_1347_v0_ = (d_1347_v0_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))) + (d_1347_v0_))
            d_1348_v1_ = d_1348_v1_
            (globalState).f24 = ((self).f27 if not(p0) else (self).f27)
            (globalState).f21 = (self).f27
        elif True:
            d_1354_v7_: _dafny.Array
            def lambda52_(d_1355_p0_):
                def lambda53_(d_1356_i0_):
                    return d_1355_p0_

                return lambda53_

            init25_ = lambda52_(p0)
            nw234_ = _dafny.Array(None, 2)
            for i0_25_ in range(nw234_.length(0)):
                nw234_[i0_25_] = init25_(i0_25_)
            d_1354_v7_ = nw234_
            index178_ = default__.safeIndex(625, (d_1354_v7_).length(0))
            (d_1354_v7_)[index178_] = (True) and (p0)
            d_1357_v8_: _dafny.Map
            d_1357_v8_ = _dafny.Map({p2: p0})
            index179_ = default__.safeIndex(625, (d_1354_v7_).length(0))
            (d_1354_v7_)[index179_] = not(not ((p0 if p2 else (self).fm35(default__.fm38(d_1357_v8_, globalState), (self).f27, globalState))) or (False))
            d_1358_v11_: _dafny.Seq
            d_1358_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
            d_1359_v12_: D2
            d_1359_v12_ = D2_DC4(len(d_1358_v11_), (self).f27, (self).f27, p3, d_1358_v11_)
            def iife104_():
                coll46_ = _dafny.Set()
                compr_46_: int
                for compr_46_ in _dafny.IntegerRange(234, -710):
                    d_1360_v9_: int = compr_46_
                    if ((234) <= (d_1360_v9_)) and ((d_1360_v9_) < (-710)):
                        def iife105_():
                            coll47_ = _dafny.Map()
                            compr_47_: int
                            for compr_47_ in _dafny.IntegerRange(-807, 827):
                                d_1361_v10_: int = compr_47_
                                if ((-807) <= (d_1361_v10_)) and ((d_1361_v10_) < (827)):
                                    coll47_[default__.safeDivisionInt(d_1361_v10_, (self).f27)] = p3
                            return _dafny.Map(coll47_)
                        coll46_ = coll46_.union(_dafny.Set([default__.safeModuloInt(d_1360_v9_, len(iife105_()
))]))
                return _dafny.Set(coll46_)
            r3 = (((self).f27) + ((self).f27)) == ((self).fm34(iife104_()
            , d_1359_v12_, globalState))
            if p2:
                d_1362_v13_: D11
                d_1362_v13_ = D11_DC25(default__.fm39(globalState))
                d_1363_v14_: _dafny.Array
                nw235_ = _dafny.Array(_dafny.Map({}), 24)
                d_1363_v14_ = nw235_
                d_1364_v15_: _dafny.Seq
                d_1364_v15_ = _dafny.SeqWithoutIsStrInference([p0, p2])
                d_1365_v16_: _dafny.Map
                d_1365_v16_ = _dafny.Map({d_1358_v11_: d_1364_v15_})
                index180_ = default__.safeIndex(165, (d_1363_v14_).length(0))
                (d_1363_v14_)[index180_] = (d_1365_v16_).set(d_1358_v11_, d_1364_v15_)
                d_1366_v17_: _dafny.Map
                d_1366_v17_ = _dafny.Map({False: (self).f27})
                d_1367_v18_: D11
                d_1367_v18_ = D11_DC24(166, D4_DC8(p3, len(d_1366_v17_), p2))
                d_1368_v19_: _dafny.MultiSet
                d_1368_v19_ = _dafny.MultiSet([p0])
                d_1369_v20_: C2
                nw236_ = C2()
                nw236_.ctor__(d_1368_v19_, (len(d_1358_v11_)) - ((self).f27))
                d_1369_v20_ = nw236_
                d_1370_v21_: _dafny.MultiSet
                d_1370_v21_ = _dafny.MultiSet([len(_dafny.Map({(self).f27: 437}))])
                d_1371_v22_: _dafny.Set
                d_1371_v22_ = _dafny.Set({((d_1370_v21_)[len(d_1358_v11_)] if (len(d_1358_v11_)) in (d_1370_v21_) else (self).f27), (self).f27})
                d_1372_v23_: _dafny.MultiSet
                d_1372_v23_ = _dafny.MultiSet([d_1371_v22_, d_1371_v22_])
                d_1373_v24_: D11
                d_1373_v24_ = D11_DC23(d_1372_v23_)
                d_1374_v25_: D11
                d_1374_v25_ = D11_DC25(d_1373_v24_)
                d_1375_v26_: D12
                d_1375_v26_ = D12_DC26(d_1369_v20_)
                d_1376_v27_: _dafny.Seq
                d_1376_v27_ = _dafny.SeqWithoutIsStrInference([(d_1375_v26_).cf46, d_1369_v20_, d_1369_v20_])
                index181_ = default__.safeIndex(165, (d_1363_v14_).length(0))
                rhs237_ = D11_DC25(d_1374_v25_)
                rhs238_ = default__.fm40(154, p0, not ((d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))]) or (p2), default__.fm41((self).f27, (self).f27, p2, globalState), globalState)
                rhs239_ = default__.fm39(globalState)
                rhs240_ = -619
                rhs241_ = (d_1376_v27_)[default__.safeIndex((self).f27, len(d_1376_v27_))]
                lhs188_ = d_1363_v14_
                lhs189_ = default__.safeIndex(165, (d_1363_v14_).length(0))
                lhs190_ = globalState
                d_1362_v13_ = rhs237_
                lhs188_[lhs189_] = rhs238_
                d_1367_v18_ = rhs239_
                lhs190_.f3 = rhs240_
                d_1369_v20_ = rhs241_
                index182_ = default__.safeIndex(625, (d_1354_v7_).length(0))
                (d_1354_v7_)[index182_] = p2
                d_1377_v28_: _dafny.Seq
                d_1377_v28_ = _dafny.SeqWithoutIsStrInference([d_1358_v11_, d_1358_v11_])
                index183_ = default__.safeIndex(625, (d_1354_v7_).length(0))
                rhs242_ = (d_1377_v28_)[default__.safeIndex((self).f27, len(d_1377_v28_))]
                rhs243_ = (self).f27
                rhs244_ = -720
                rhs245_ = not((848) < ((self).f27))
                lhs191_ = globalState
                lhs192_ = globalState
                lhs193_ = d_1354_v7_
                lhs194_ = default__.safeIndex(625, (d_1354_v7_).length(0))
                d_1358_v11_ = rhs242_
                lhs191_.f15 = rhs243_
                lhs192_.f0 = rhs244_
                lhs193_[lhs194_] = rhs245_
                d_1378_v29_: C0
                nw237_ = C0()
                nw237_.ctor__(len((d_1364_v15_) + (d_1364_v15_)), d_1358_v11_)
                d_1378_v29_ = nw237_
                d_1379_v30_: _dafny.Map
                d_1379_v30_ = _dafny.Map({754: len(_dafny.Map({478: True}))})
                d_1380_v31_: _dafny.Map
                d_1380_v31_ = _dafny.Map({(d_1378_v29_).f36: p2})
                d_1381_v32_: D9
                d_1381_v32_ = D9_DC20(len(d_1364_v15_), ((d_1380_v31_)[(d_1378_v29_).f36] if ((d_1378_v29_).f36) in (d_1380_v31_) else (d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))]), p0, (d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))], d_1379_v30_)
                d_1382_v34_: _dafny.Seq
                d_1382_v34_ = _dafny.SeqWithoutIsStrInference([(d_1378_v29_).f36, default__.fm1(globalState)])
                d_1383_v37_: _dafny.Seq
                def iife106_():
                    coll48_ = _dafny.Set()
                    compr_48_: int
                    for compr_48_ in (d_1382_v34_).Elements:
                        d_1384_v36_: int = compr_48_
                        if (d_1384_v36_) in (d_1382_v34_):
                            coll48_ = coll48_.union(_dafny.Set([default__.safeDivisionInt(d_1384_v36_, 29)]))
                    return _dafny.Set(coll48_)
                d_1383_v37_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len(iife106_()
                ): (d_1378_v29_).f36}), d_1379_v30_])
                d_1385_v38_: _dafny.Map
                d_1385_v38_ = _dafny.Map({len(d_1364_v15_): p3})
                d_1386_v39_: _dafny.Seq
                d_1386_v39_ = _dafny.SeqWithoutIsStrInference([d_1385_v38_])
                d_1387_v40_: D15
                d_1387_v40_ = D15_DC35(d_1386_v39_)
                d_1388_v41_: _dafny.Array
                nw238_ = _dafny.Array(None, 29)
                nw238_[int(0)] = (d_1379_v30_) | (_dafny.Map({len((d_1378_v29_).f37): (d_1378_v29_).f36}))
                nw238_[int(1)] = (d_1381_v32_).cf38
                nw238_[int(2)] = _dafny.Map({(self).f27: (self).f27})
                def iife107_():
                    coll49_ = _dafny.Map()
                    compr_49_: int
                    for compr_49_ in (_dafny.Map({len(d_1382_v34_): (d_1378_v29_).f36})).keys.Elements:
                        d_1389_v33_: int = compr_49_
                        if (d_1389_v33_) in (_dafny.Map({len(d_1382_v34_): (d_1378_v29_).f36})):
                            coll49_[(d_1389_v33_) + ((d_1378_v29_).f36)] = 319
                    return _dafny.Map(coll49_)
                nw238_[int(3)] = (d_1379_v30_) | (iife107_()
                )
                nw238_[int(4)] = d_1379_v30_
                nw238_[int(5)] = (d_1379_v30_ if False else d_1379_v30_)
                nw238_[int(6)] = d_1379_v30_
                def iife108_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in _dafny.IntegerRange(32, 657):
                        d_1390_v35_: int = compr_50_
                        if ((32) <= (d_1390_v35_)) and ((d_1390_v35_) < (657)):
                            coll50_[default__.safeModuloInt(d_1390_v35_, (d_1378_v29_).f36)] = (d_1378_v29_).f36
                    return _dafny.Map(coll50_)
                nw238_[int(7)] = (d_1379_v30_) | (iife108_()
                )
                nw238_[int(8)] = _dafny.Map({(self).f27: (self).f27})
                nw238_[int(9)] = d_1379_v30_
                nw238_[int(10)] = d_1379_v30_
                nw238_[int(11)] = d_1379_v30_
                nw238_[int(12)] = default__.fm42(((d_1368_v19_)[p2] if (p2) in (d_1368_v19_) else (self).f27), True, (d_1378_v29_).f36, globalState)
                nw238_[int(13)] = d_1379_v30_
                nw238_[int(14)] = d_1379_v30_
                nw238_[int(15)] = d_1379_v30_
                nw238_[int(16)] = (d_1379_v30_) | (d_1379_v30_)
                nw238_[int(17)] = d_1379_v30_
                nw238_[int(18)] = d_1379_v30_
                nw238_[int(19)] = (d_1383_v37_)[default__.safeIndex(((_dafny.MultiSet((d_1387_v40_).cf63)).set(d_1385_v38_, default__.abs((d_1378_v29_).f36))).cardinality, len(d_1383_v37_))]
                nw238_[int(20)] = d_1379_v30_
                nw238_[int(21)] = d_1379_v30_
                nw238_[int(22)] = (d_1379_v30_) | ((d_1379_v30_).set((self).fm34(d_1371_v22_, d_1359_v12_, globalState), (d_1378_v29_).f36))
                nw238_[int(23)] = d_1379_v30_
                nw238_[int(24)] = (d_1379_v30_) | (((d_1381_v32_).cf38).set((d_1378_v29_).f36, (self).f27))
                nw238_[int(25)] = d_1379_v30_
                nw238_[int(26)] = d_1379_v30_
                nw238_[int(27)] = _dafny.Map({len(d_1364_v15_): (d_1378_v29_).f36})
                nw238_[int(28)] = _dafny.Map({(self).f27: (d_1378_v29_).f36})
                d_1388_v41_ = nw238_
                d_1391_v42_: D16
                d_1391_v42_ = D16_DC39(d_1388_v41_)
                d_1392_v43_: bool
                d_1392_v43_ = (d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))]
                d_1393_v44_: D9
                d_1393_v44_ = D9_DC18(True)
                rhs246_ = (d_1378_v29_).f36
                rhs247_ = (d_1391_v42_).cf73
                rhs248_ = True
                rhs249_ = (default__.fm43(p2, p2, p0, globalState)) + (d_1382_v34_)
                rhs250_ = not (False) or ((d_1393_v44_).cf32)
                lhs195_ = globalState
                lhs196_ = globalState
                lhs195_.f21 = rhs246_
                d_1388_v41_ = rhs247_
                lhs196_.f26 = rhs248_
                d_1382_v34_ = rhs249_
                r1 = rhs250_
            elif True:
                d_1394_v45_: D9
                d_1394_v45_ = D9_DC19(p0)
                d_1394_v45_ = (d_1394_v45_ if (d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))] else d_1394_v45_)
                d_1395_v46_: C0
                nw239_ = C0()
                nw239_.ctor__((self).f27, d_1358_v11_)
                d_1395_v46_ = nw239_
                d_1396_v47_: _dafny.Array
                def lambda54_(d_1397_v46_):
                    def lambda55_(d_1398_i1_):
                        return (d_1398_i1_) - ((d_1397_v46_).f36)

                    return lambda55_

                init26_ = lambda54_(d_1395_v46_)
                nw240_ = _dafny.Array(None, 20)
                for i0_26_ in range(nw240_.length(0)):
                    nw240_[i0_26_] = init26_(i0_26_)
                d_1396_v47_ = nw240_
                d_1399_v48_: _dafny.Map
                d_1399_v48_ = _dafny.Map({(self).f27: p2})
                d_1400_v49_: _dafny.Map
                d_1400_v49_ = _dafny.Map({d_1396_v47_: len(d_1399_v48_)})
                d_1401_v50_: _dafny.Map
                d_1401_v50_ = _dafny.Map({d_1400_v49_: p0})
                (globalState).f21 = ((self).f27 if ((d_1401_v50_)[d_1400_v49_] if (d_1400_v49_) in (d_1401_v50_) else p2) else 729)
                (globalState).f26 = (len(d_1358_v11_)) < ((d_1395_v46_).f36)
                index184_ = default__.safeIndex(625, (d_1354_v7_).length(0))
                (d_1354_v7_)[index184_] = p0
            r3 = (d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))]
            index185_ = default__.safeIndex(748, ((self).f28).length(0))
            ((self).f28)[index185_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "naxcrpkh"))
            d_1402_v51_: _dafny.MultiSet
            d_1402_v51_ = _dafny.MultiSet([(d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))]])
            d_1403_v52_: D15
            d_1403_v52_ = D15_DC38((self).f27, True, 245)
            d_1404_v53_: bool
            d_1404_v53_ = (d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))]
            d_1405_v55_: D16
            def iife109_():
                coll51_ = _dafny.Set()
                compr_51_: int
                for compr_51_ in (default__.fm43((d_1403_v52_).cf71, d_1404_v53_, (d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))], globalState)).Elements:
                    d_1406_v54_: int = compr_51_
                    if (d_1406_v54_) in (default__.fm43((d_1403_v52_).cf71, d_1404_v53_, (d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))], globalState)):
                        coll51_ = coll51_.union(_dafny.Set([(d_1406_v54_) + (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dp"))): False})))]))
                return _dafny.Set(coll51_)
            d_1405_v55_ = D16_DC40(not((d_1402_v51_).ispropersubset(d_1402_v51_)), (0) - (default__.safeModuloInt((self).fm34(iife109_()
, d_1359_v12_, globalState), (0) - ((self).f27))), (self).f27)
            d_1407_v56_: _dafny.Array
            def lambda56_(d_1408_p0_):
                def lambda57_(d_1409_i2_):
                    return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, d_1408_p0_])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1408_p0_])]))

                return lambda57_

            init27_ = lambda56_(p0)
            nw241_ = _dafny.Array(None, 23)
            for i0_27_ in range(nw241_.length(0)):
                nw241_[i0_27_] = init27_(i0_27_)
            d_1407_v56_ = nw241_
            index186_ = default__.safeIndex(532, (d_1407_v56_).length(0))
            (d_1407_v56_)[index186_] = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0]) for d_1410_i3_ in range(default__.abs(18))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_1354_v7_)[default__.safeIndex(625, (d_1354_v7_).length(0))]]) for d_1411_i4_ in range(default__.abs(558))]))
            d_1412_v57_: D4
            d_1412_v57_ = D4_DC7(d_1354_v7_)
            pat_let_tv26_ = p2
            index187_ = default__.safeIndex(748, ((self).f28).length(0))
            index188_ = default__.safeIndex(532, (d_1407_v56_).length(0))
            def iife110_(_pat_let29_0):
                def iife111_(d_1413_dt__update__tmp_h1_):
                    def iife112_(_pat_let30_0):
                        def iife113_(d_1414_dt__update_hcf74_h0_):
                            def iife114_(_pat_let31_0):
                                def iife115_(d_1415_dt__update_hcf76_h0_):
                                    return D16_DC40(d_1414_dt__update_hcf74_h0_, (d_1413_dt__update__tmp_h1_).cf75, d_1415_dt__update_hcf76_h0_)
                                return iife115_(_pat_let31_0)
                            return iife114_(default__.safeDivisionInt((self).f27, 527))
                        return iife113_(_pat_let30_0)
                    return iife112_(pat_let_tv26_)
                return iife111_(_pat_let29_0)
            rhs251_ = d_1358_v11_
            rhs252_ = iife110_(d_1405_v55_)
            rhs253_ = default__.fm44(globalState)
            rhs254_ = (default__.safeDivisionInt((self).f27, (self).f27)) + ((155) + ((self).f27))
            rhs255_ = (d_1412_v57_).cf12
            lhs197_ = (self).f28
            lhs198_ = default__.safeIndex(748, ((self).f28).length(0))
            lhs199_ = d_1407_v56_
            lhs200_ = default__.safeIndex(532, (d_1407_v56_).length(0))
            lhs201_ = globalState
            lhs197_[lhs198_] = rhs251_
            d_1405_v55_ = rhs252_
            lhs199_[lhs200_] = rhs253_
            lhs201_.f15 = rhs254_
            d_1354_v7_ = rhs255_
        d_1416_v58_: _dafny.Map
        d_1416_v58_ = _dafny.Map({p0: p0})
        d_1417_v59_: _dafny.Array
        nw242_ = _dafny.Array(None, 25)
        nw242_[int(0)] = default__.fm38(d_1416_v58_, globalState)
        nw242_[int(1)] = p3
        nw242_[int(2)] = default__.fm38(_dafny.Map({p0: p0}), globalState)
        nw242_[int(3)] = p3
        nw242_[int(4)] = p3
        nw242_[int(5)] = p3
        nw242_[int(6)] = p3
        nw242_[int(7)] = p3
        nw242_[int(8)] = p3
        nw242_[int(9)] = _dafny.CodePoint('c')
        nw242_[int(10)] = p3
        nw242_[int(11)] = p3
        nw242_[int(12)] = _dafny.CodePoint('g')
        nw242_[int(13)] = p3
        nw242_[int(14)] = p3
        nw242_[int(15)] = p3
        nw242_[int(16)] = p3
        nw242_[int(17)] = _dafny.CodePoint('b')
        nw242_[int(18)] = p3
        nw242_[int(19)] = p3
        nw242_[int(20)] = p3
        nw242_[int(21)] = p3
        nw242_[int(22)] = p3
        nw242_[int(23)] = default__.fm38(d_1416_v58_, globalState)
        nw242_[int(24)] = p3
        d_1417_v59_ = nw242_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1417_v59_).length(0)):
            d_1418_i5_: int = guard_loop_5_
            if (True) and (((0) <= (d_1418_i5_)) and ((d_1418_i5_) < ((d_1417_v59_).length(0)))):
                (d_1417_v59_)[(d_1418_i5_)] = _dafny.CodePoint('y')
        d_1419_v60_: _dafny.Seq
        d_1419_v60_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_1420_v61_: T0
        nw243_ = C3()
        nw243_.ctor__((len(d_1419_v60_)) + ((self).f27), (self).f28)
        d_1420_v61_ = nw243_
        d_1420_v61_ = d_1420_v61_
        d_1421_v62_: D12
        d_1421_v62_ = D12_DC28((self).f27, (0) - ((d_1420_v61_).f27), p0)
        d_1422_i6_: int
        d_1422_i6_ = 0
        with _dafny.label("6"):
            while default__.fm0(len(((d_1419_v60_) + (default__.fm41((d_1420_v61_).f27, (d_1420_v61_).f27, (d_1421_v62_).cf53, globalState))).set(default__.safeIndex((d_1420_v61_).f27, len((d_1419_v60_) + (default__.fm41((d_1420_v61_).f27, (d_1420_v61_).f27, (d_1421_v62_).cf53, globalState)))), p0)), globalState):
                with _dafny.c_label("6"):
                    if (d_1422_i6_) >= (100):
                        raise _dafny.Break("6")
                    d_1422_i6_ = (d_1422_i6_) + (1)
                    if p2:
                        d_1423_v63_: C3
                        nw244_ = C3()
                        nw244_.ctor__(342, (d_1420_v61_).f28)
                        d_1423_v63_ = nw244_
                        d_1424_v64_: _dafny.Seq
                        d_1424_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - ((d_1420_v61_).f27)])])
                        d_1424_v64_ = d_1424_v64_
                        d_1425_v65_: _dafny.Seq
                        d_1425_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                        d_1426_v66_: _dafny.Set
                        d_1426_v66_ = _dafny.Set({(d_1420_v61_).f27, default__.fm1(globalState), (self).f27, (d_1420_v61_).f27})
                        d_1427_v67_: _dafny.Set
                        d_1427_v67_ = _dafny.Set({(d_1419_v60_).set(default__.safeIndex(len(d_1426_v66_), len(d_1419_v60_)), p2)})
                        d_1428_v68_: _dafny.Array
                        nw245_ = _dafny.Array(None, 10)
                        nw245_[int(0)] = (d_1420_v61_).f27
                        nw245_[int(1)] = (d_1420_v61_).f27
                        nw245_[int(2)] = (self).f27
                        nw245_[int(3)] = (d_1420_v61_).f27
                        nw245_[int(4)] = (d_1420_v61_).f27
                        nw245_[int(5)] = (d_1420_v61_).f27
                        nw245_[int(6)] = default__.safeModuloInt(len(d_1425_v65_), (self).f27)
                        nw245_[int(7)] = ((d_1420_v61_).f27) + (len(d_1427_v67_))
                        nw245_[int(8)] = len(default__.fm47(globalState))
                        nw245_[int(9)] = (d_1420_v61_).f27
                        d_1428_v68_ = nw245_
                        index189_ = default__.safeIndex(786, (d_1428_v68_).length(0))
                        (d_1428_v68_)[index189_] = default__.safeDivisionInt((d_1420_v61_).f27, 595)
                        d_1429_v69_: _dafny.Array
                        nw246_ = _dafny.Array(False, 4)
                        d_1429_v69_ = nw246_
                        index190_ = default__.safeIndex(439, (d_1429_v69_).length(0))
                        (d_1429_v69_)[index190_] = not (p0) or (p0)
                        d_1430_v70_: _dafny.Map
                        d_1430_v70_ = _dafny.Map({p0: (d_1420_v61_).f27})
                        d_1431_v71_: _dafny.Map
                        d_1431_v71_ = _dafny.Map({p0: len(d_1430_v70_)})
                        d_1432_v72_: _dafny.Map
                        d_1432_v72_ = _dafny.Map({(d_1420_v61_).f27: p2})
                        index191_ = default__.safeIndex(786, (d_1428_v68_).length(0))
                        index192_ = default__.safeIndex(439, (d_1429_v69_).length(0))
                        rhs256_ = (d_1420_v61_).f27
                        rhs257_ = ((0) - ((self).f27)) > (len(((d_1430_v70_).set(False, (d_1420_v61_).f27)).set(p0, len(d_1432_v72_))))
                        rhs258_ = p0
                        rhs259_ = default__.fm0((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nngcbxfxb"))) if p0 else (d_1420_v61_).f27), globalState)
                        rhs260_ = (len(_dafny.Map({p0: (0) - ((d_1420_v61_).f27)}))) > ((d_1420_v61_).f27)
                        lhs202_ = d_1428_v68_
                        lhs203_ = default__.safeIndex(786, (d_1428_v68_).length(0))
                        lhs204_ = globalState
                        lhs205_ = globalState
                        lhs206_ = d_1429_v69_
                        lhs207_ = default__.safeIndex(439, (d_1429_v69_).length(0))
                        lhs202_[lhs203_] = rhs256_
                        lhs204_.f26 = rhs257_
                        r1 = rhs258_
                        lhs205_.f26 = rhs259_
                        lhs206_[lhs207_] = rhs260_
                        d_1428_v68_ = d_1428_v68_
                        (globalState).f0 = default__.safeModuloInt((d_1428_v68_)[default__.safeIndex(786, (d_1428_v68_).length(0))], 477)
                    elif True:
                        d_1433_v73_: _dafny.Map
                        d_1433_v73_ = _dafny.Map({(0) - (default__.safeModuloInt(default__.fm1(globalState), (self).f27)): p0})
                        d_1433_v73_ = (default__.fm48(p2, p2, ((_dafny.MultiSet([(d_1420_v61_).f27, (0) - ((d_1420_v61_).f27)])).set((d_1420_v61_).f27, default__.abs(-216))).cardinality, globalState)) | (d_1433_v73_)
                        d_1434_v74_: _dafny.Seq
                        d_1434_v74_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_1419_v60_) for d_1435_i7_ in range(default__.abs(33))]), _dafny.SeqWithoutIsStrInference([(d_1420_v61_).f27 for d_1436_i8_ in range(default__.abs(942))])])
                        rhs261_ = (default__.fm1(globalState) if p0 else (_dafny.MultiSet((d_1434_v74_)[default__.safeIndex((self).f27, len(d_1434_v74_))])).cardinality)
                        rhs262_ = p2
                        lhs208_ = globalState
                        lhs209_ = globalState
                        lhs208_.f15 = rhs261_
                        lhs209_.f26 = rhs262_
                        d_1437_v75_: _dafny.Seq
                        d_1437_v75_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cscw"))])
                        d_1438_v76_: _dafny.Seq
                        d_1438_v76_ = _dafny.SeqWithoutIsStrInference([(d_1420_v61_).f27, 892, (d_1420_v61_).f27])
                        d_1439_v77_: _dafny.MultiSet
                        d_1439_v77_ = _dafny.MultiSet([p2])
                        d_1440_v78_: _dafny.Set
                        d_1440_v78_ = _dafny.Set({(self).f27})
                        d_1441_v79_: _dafny.Map
                        d_1441_v79_ = _dafny.Map({p2: (d_1420_v61_).f27})
                        rhs263_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwi"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqpplff"))) for d_1442_i9_ in range(default__.abs(84))])
                        rhs264_ = ((d_1439_v77_) | (d_1439_v77_)).isdisjoint((d_1439_v77_).set(p0, default__.abs((d_1420_v61_).f27)))
                        rhs265_ = d_1438_v76_
                        rhs266_ = not((d_1440_v78_).isdisjoint(d_1440_v78_))
                        rhs267_ = (len(d_1441_v79_)) - ((self).f27)
                        lhs210_ = globalState
                        d_1437_v75_ = rhs263_
                        r1 = rhs264_
                        d_1438_v76_ = rhs265_
                        r1 = rhs266_
                        lhs210_.f0 = rhs267_
                        (globalState).f18 = ((0) - ((866) - ((0) - (default__.fm1(globalState))))) * ((d_1420_v61_).f27)
                        d_1443_v80_: bool
                        d_1444_v81_: int
                        out42_: bool
                        out43_: int
                        out42_, out43_ = default__.m0(p0, (d_1420_v61_).f27, (d_1420_v61_).f27, (d_1420_v61_).f27, globalState)
                        d_1443_v80_ = out42_
                        d_1444_v81_ = out43_
                    d_1445_v82_: _dafny.Seq
                    d_1445_v82_ = _dafny.SeqWithoutIsStrInference([(d_1420_v61_).f27, (d_1420_v61_).f27])
                    d_1446_v83_: _dafny.Seq
                    d_1446_v83_ = _dafny.SeqWithoutIsStrInference([d_1445_v82_])
                    (globalState).f3 = (125) * (((d_1420_v61_).f27) * (len(d_1446_v83_)))
                    d_1447_v84_: _dafny.Map
                    d_1447_v84_ = _dafny.Map({(self).f27: True})
                    d_1448_v85_: _dafny.Seq
                    d_1448_v85_ = _dafny.SeqWithoutIsStrInference([d_1447_v84_, d_1447_v84_])
                    d_1449_v86_: _dafny.Seq
                    d_1449_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cyisvag"))
                    d_1450_v87_: _dafny.MultiSet
                    d_1450_v87_ = _dafny.MultiSet([p2])
                    d_1451_v88_: _dafny.MultiSet
                    d_1451_v88_ = _dafny.MultiSet([(self).f27, len(d_1419_v60_)])
                    d_1452_v89_: _dafny.Array
                    nw247_ = _dafny.Array(None, 27)
                    nw247_[int(0)] = (d_1420_v61_).f27
                    nw247_[int(1)] = (len((d_1448_v85_)[default__.safeIndex(360, len(d_1448_v85_))])) * ((d_1420_v61_).f27)
                    nw247_[int(2)] = (d_1420_v61_).f27
                    nw247_[int(3)] = default__.safeDivisionInt(len(default__.fm41(len(_dafny.Map({p0: p2})), (self).f27, p0, globalState)), (0) - ((self).f27))
                    nw247_[int(4)] = (d_1420_v61_).f27
                    nw247_[int(5)] = (d_1420_v61_).f27
                    nw247_[int(6)] = -541
                    nw247_[int(7)] = (d_1420_v61_).f27
                    nw247_[int(8)] = (d_1420_v61_).f27
                    nw247_[int(9)] = (self).f27
                    nw247_[int(10)] = (d_1420_v61_).f27
                    nw247_[int(11)] = (self).f27
                    nw247_[int(12)] = (self).f27
                    nw247_[int(13)] = (d_1420_v61_).f27
                    nw247_[int(14)] = (d_1420_v61_).f27
                    nw247_[int(15)] = (self).f27
                    nw247_[int(16)] = (self).f27
                    nw247_[int(17)] = (default__.fm1(globalState)) * (len(d_1449_v86_))
                    nw247_[int(18)] = len(_dafny.Set({(d_1419_v60_)[default__.safeIndex(len(_dafny.Set({(d_1419_v60_)[default__.safeIndex(9, len(d_1419_v60_))], p0})), len(d_1419_v60_))], p2}))
                    nw247_[int(19)] = ((d_1420_v61_).f27) + (len(d_1445_v82_))
                    nw247_[int(20)] = (self).f27
                    nw247_[int(21)] = (d_1420_v61_).f27
                    nw247_[int(22)] = (0) - ((d_1420_v61_).f27)
                    nw247_[int(23)] = ((self).f27) * ((d_1420_v61_).f27)
                    nw247_[int(24)] = -743
                    nw247_[int(25)] = ((d_1450_v87_)[p2] if (p2) in (d_1450_v87_) else ((d_1451_v88_)[(d_1420_v61_).f27] if ((d_1420_v61_).f27) in (d_1451_v88_) else (d_1450_v87_).cardinality))
                    nw247_[int(26)] = (d_1420_v61_).f27
                    d_1452_v89_ = nw247_
                    index193_ = default__.safeIndex(634, (d_1452_v89_).length(0))
                    (d_1452_v89_)[index193_] = (self).f27
                    index194_ = default__.safeIndex(634, (d_1452_v89_).length(0))
                    (d_1452_v89_)[index194_] = len(d_1445_v82_)
                    d_1453_v90_: _dafny.Set
                    d_1453_v90_ = _dafny.Set({p2, p0, p0})
                    (globalState).f0 = default__.safeDivisionInt((d_1420_v61_).f27, ((d_1451_v88_)[(d_1452_v89_)[default__.safeIndex(634, (d_1452_v89_).length(0))]] if ((d_1452_v89_)[default__.safeIndex(634, (d_1452_v89_).length(0))]) in (d_1451_v88_) else len(d_1453_v90_)))
                    pass
            pass
        d_1454_v91_: _dafny.Seq
        d_1454_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cd"))
        d_1455_v92_: D4
        d_1455_v92_ = D4_DC8(p3, len(d_1454_v91_), p0)
        (globalState).f24 = (d_1455_v92_).cf14
        d_1456_v93_: _dafny.Array
        nw248_ = _dafny.Array(_dafny.Map({}), 10)
        d_1456_v93_ = nw248_
        d_1457_v94_: D16
        d_1457_v94_ = D16_DC39(d_1456_v93_)
        d_1458_v95_: _dafny.Array
        nw249_ = _dafny.Array(False, 3)
        d_1458_v95_ = nw249_
        d_1459_v96_: _dafny.Map
        d_1459_v96_ = _dafny.Map({d_1457_v94_: d_1458_v95_})
        d_1459_v96_ = d_1459_v96_
        r0 = d_1419_v60_
        r1 = (default__.safeDivisionInt((self).f27, (0) - ((d_1420_v61_).f27))) < (987)
        d_1460_v97_: _dafny.Array
        nw250_ = _dafny.Array(_dafny.Array(None, 0), 23)
        d_1460_v97_ = nw250_
        d_1461_v98_: _dafny.Seq
        d_1461_v98_ = _dafny.SeqWithoutIsStrInference([d_1460_v97_, d_1460_v97_, d_1460_v97_, d_1460_v97_])
        d_1462_v99_: _dafny.Seq
        d_1462_v99_ = _dafny.SeqWithoutIsStrInference([len(d_1454_v91_)])
        r2 = (d_1461_v98_)[default__.safeIndex(len(d_1462_v99_), len(d_1461_v98_))]
        d_1463_v100_: _dafny.MultiSet
        d_1463_v100_ = _dafny.MultiSet([len(d_1454_v91_)])
        d_1464_v101_: _dafny.MultiSet
        d_1464_v101_ = d_1463_v100_
        pat_let_tv27_ = p2
        def lambda58_(source23_):
            d_1465___mcc_h0_ = source23_
            d_1466_cf18_ = d_1465___mcc_h0_
            return pat_let_tv27_

        r3 = lambda58_(d_1464_v101_)
        return r0, r1, r2, r3


class C5(T1, T0):
    def  __init__(self):
        self._f29: _dafny.MultiSet = _dafny.MultiSet({})
        self._f30: _dafny.Seq = _dafny.Seq({})
        self._f27: int = int(0)
        self._f28: _dafny.Array = _dafny.Array(None, 0)
        self._f38: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f29(self):
        return self._f29
    @property
    def f30(self):
        return self._f30
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f38, f29, f30, f27, f28):
        (self)._f38 = f38
        (self)._f29 = f29
        (self)._f30 = f30
        (self)._f27 = f27
        (self)._f28 = f28

    def fm3(self, p0, globalState):
        return not (False) or (True)

    def fm4(self, p0, p1, p2, globalState):
        if (True) in ((self).f29):
            return ((self).f29)[True]
        elif True:
            return (0) - (((self).f27) * ((self).f27))

    def fm2(self, p0, p1, p2, globalState):
        source24_ = D8_DC16((self).f27, (0) - ((_dafny.MultiSet([205])).cardinality))
        if source24_.is_DC16:
            d_1467___mcc_h0_ = source24_.cf29
            d_1468___mcc_h1_ = source24_.cf30
            d_1469_cf30_ = d_1468___mcc_h1_
            d_1470_cf29_ = d_1467___mcc_h0_
            return (_dafny.Map({True: d_1470_cf29_})) | (_dafny.Map({not(not(True)): (self).f27}))
        elif True:
            d_1471___mcc_h2_ = source24_.cf28
            d_1472_cf28_ = d_1471___mcc_h2_
            return _dafny.Map({True: (self).f27})

    def m1(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        d_1473_v0_: bool
        d_1473_v0_ = True
        d_1474_v1_: _dafny.Seq
        d_1474_v1_ = _dafny.SeqWithoutIsStrInference([not((((self).f29).set(not(d_1473_v0_), default__.abs((self).f27))).isdisjoint((self).f29))])
        if (d_1474_v1_)[default__.safeIndex(default__.fm1(globalState), len(d_1474_v1_))]:
            d_1475_v2_: _dafny.MultiSet
            d_1475_v2_ = _dafny.MultiSet([(self).f27])
            d_1476_v3_: _dafny.Map
            d_1476_v3_ = _dafny.Map({(d_1475_v2_).cardinality: (self).f27})
            d_1477_v4_: D9
            d_1477_v4_ = D9_DC20((self).f27, d_1473_v0_, d_1473_v0_, d_1473_v0_, d_1476_v3_)
            pat_let_tv28_ = d_1473_v0_
            pat_let_tv29_ = d_1473_v0_
            pat_let_tv30_ = d_1473_v0_
            pat_let_tv31_ = d_1473_v0_
            def iife117_(_pat_let33_0):
                def iife118_(d_1478_dt__update__tmp_h0_):
                    def iife119_(_pat_let34_0):
                        def iife120_(d_1479_dt__update_hcf36_h0_):
                            return D9_DC20((d_1478_dt__update__tmp_h0_).cf34, (d_1478_dt__update__tmp_h0_).cf35, d_1479_dt__update_hcf36_h0_, (d_1478_dt__update__tmp_h0_).cf37, (d_1478_dt__update__tmp_h0_).cf38)
                        return iife120_(_pat_let34_0)
                    return iife119_(pat_let_tv28_)
                return iife118_(_pat_let33_0)
            def iife116_(_pat_let32_0):
                def iife121_(d_1480_dt__update__tmp_h1_):
                    def iife122_(_pat_let35_0):
                        def iife123_(d_1481_dt__update_hcf37_h0_):
                            def iife124_(_pat_let36_0):
                                def iife125_(d_1482_dt__update_hcf34_h0_):
                                    def iife126_(_pat_let37_0):
                                        def iife127_(d_1483_dt__update_hcf35_h0_):
                                            return D9_DC20(d_1482_dt__update_hcf34_h0_, d_1483_dt__update_hcf35_h0_, (d_1480_dt__update__tmp_h1_).cf36, d_1481_dt__update_hcf37_h0_, (d_1480_dt__update__tmp_h1_).cf38)
                                        return iife127_(_pat_let37_0)
                                    return iife126_((pat_let_tv29_ if pat_let_tv30_ else pat_let_tv31_))
                                return iife125_(_pat_let36_0)
                            return iife124_((self).f27)
                        return iife123_(_pat_let35_0)
                    return iife122_(not(False))
                return iife121_(_pat_let32_0)
            source25_ = iife116_(iife117_(d_1477_v4_))
            if source25_.is_DC18:
                d_1484___mcc_h0_ = source25_.cf32
                d_1485_cf32_ = d_1484___mcc_h0_
                d_1486_v5_: _dafny.Seq
                d_1486_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fq"))
                d_1487_v6_: C0
                nw251_ = C0()
                nw251_.ctor__((self).f27, d_1486_v5_)
                d_1487_v6_ = nw251_
                d_1488_v7_: D6
                d_1488_v7_ = D6_DC11(d_1487_v6_)
                d_1489_v8_: _dafny.Map
                d_1489_v8_ = _dafny.Map({d_1485_cf32_: d_1488_v7_})
                d_1490_v9_: _dafny.Map
                d_1490_v9_ = _dafny.Map({(self).f27: d_1489_v8_})
                rhs268_ = (self).f27
                rhs269_ = default__.safeDivisionInt((self).f27, ((self).f27) + ((self).f27))
                rhs270_ = ((self).f27) > ((self).f27)
                rhs271_ = (((self).f30)[default__.safeIndex((self).f27, len((self).f30))]) in (d_1490_v9_)
                rhs272_ = (d_1487_v6_).f36
                lhs211_ = globalState
                lhs212_ = globalState
                lhs213_ = globalState
                lhs214_ = globalState
                lhs211_.f18 = rhs268_
                lhs212_.f0 = rhs269_
                lhs213_.f26 = rhs270_
                d_1473_v0_ = rhs271_
                lhs214_.f18 = rhs272_
                d_1491_v10_: _dafny.Array
                def lambda59_(d_1492_v0_):
                    def lambda60_(d_1493_i0_):
                        return d_1492_v0_

                    return lambda60_

                init28_ = lambda59_(d_1473_v0_)
                nw252_ = _dafny.Array(None, 21)
                for i0_28_ in range(nw252_.length(0)):
                    nw252_[i0_28_] = init28_(i0_28_)
                d_1491_v10_ = nw252_
                d_1494_v11_: _dafny.Set
                d_1494_v11_ = _dafny.Set({(d_1487_v6_).f36})
                d_1495_v12_: _dafny.MultiSet
                d_1495_v12_ = _dafny.MultiSet([d_1494_v11_])
                d_1496_v13_: D11
                d_1496_v13_ = D11_DC23(d_1495_v12_)
                d_1497_v14_: _dafny.Map
                d_1497_v14_ = _dafny.Map({_dafny.Set({d_1496_v13_, d_1496_v13_, d_1496_v13_}): d_1491_v10_})
                d_1498_v15_: _dafny.Set
                d_1498_v15_ = _dafny.Set({d_1496_v13_})
                d_1491_v10_ = ((d_1497_v14_)[(d_1498_v15_) | (d_1498_v15_)] if ((d_1498_v15_) | (d_1498_v15_)) in (d_1497_v14_) else d_1491_v10_)
                d_1476_v3_ = d_1476_v3_
                d_1491_v10_ = d_1491_v10_
            elif source25_.is_DC19:
                d_1499___mcc_h1_ = source25_.cf33
                d_1500_cf33_ = d_1499___mcc_h1_
                d_1501_v16_: str
                d_1501_v16_ = _dafny.CodePoint('h')
                d_1502_v17_: _dafny.Map
                d_1502_v17_ = _dafny.Map({d_1501_v16_: (self).f27})
                d_1503_v19_: _dafny.MultiSet
                def iife128_():
                    coll52_ = _dafny.Map()
                    compr_52_: str
                    for compr_52_ in (_dafny.MultiSet([d_1501_v16_, _dafny.CodePoint('o'), d_1501_v16_])).Elements:
                        d_1504_v18_: str = compr_52_
                        if (d_1504_v18_) in (_dafny.MultiSet([d_1501_v16_, _dafny.CodePoint('o'), d_1501_v16_])):
                            coll52_[d_1504_v18_] = ((d_1476_v3_)[(self).f27] if ((self).f27) in (d_1476_v3_) else (self).f27)
                    return _dafny.Map(coll52_)
                d_1503_v19_ = _dafny.MultiSet([d_1502_v17_, _dafny.Map({d_1501_v16_: (self).f27}), d_1502_v17_, iife128_()
                ])
                d_1505_v20_: _dafny.Map
                d_1505_v20_ = _dafny.Map({d_1473_v0_: not(d_1473_v0_)})
                d_1506_v21_: _dafny.Seq
                d_1506_v21_ = _dafny.SeqWithoutIsStrInference([d_1505_v20_])
                d_1507_v22_: _dafny.Set
                d_1507_v22_ = _dafny.Set({(self).f27, (self).f27, (self).f27, (self).f27, len(default__.fm31(d_1473_v0_, d_1500_cf33_, globalState))})
                d_1508_v23_: _dafny.Map
                d_1508_v23_ = _dafny.Map({(d_1474_v1_)[default__.safeIndex((self).f27, len(d_1474_v1_))]: len(_dafny.Map({d_1501_v16_: (self).f27}))})
                d_1509_v24_: _dafny.Seq
                d_1509_v24_ = _dafny.SeqWithoutIsStrInference([d_1501_v16_])
                d_1510_v25_: _dafny.Array
                nw253_ = _dafny.Array(None, 19)
                nw253_[int(0)] = ((d_1503_v19_)[_dafny.Map({d_1501_v16_: (0) - ((self).f27)})] if (_dafny.Map({d_1501_v16_: (0) - ((self).f27)})) in (d_1503_v19_) else (self).f27)
                nw253_[int(1)] = (self).f27
                nw253_[int(2)] = (self).f27
                nw253_[int(3)] = len(_dafny.SeqWithoutIsStrInference([d_1500_cf33_, d_1500_cf33_]))
                nw253_[int(4)] = ((self).f27 if d_1473_v0_ else (0) - ((self).f27))
                nw253_[int(5)] = 19
                nw253_[int(6)] = len((d_1506_v21_)[default__.safeIndex((self).f27, len(d_1506_v21_))])
                nw253_[int(7)] = (D10_DC22(len(d_1505_v20_), (self).f29)).cf40
                nw253_[int(8)] = (self).f27
                nw253_[int(9)] = (self).f27
                nw253_[int(10)] = default__.safeDivisionInt(967, (self).f27)
                nw253_[int(11)] = (self).f27
                nw253_[int(12)] = (self).f27
                nw253_[int(13)] = (((self).f30)[default__.safeIndex((self).f27, len((self).f30))]) + (-925)
                nw253_[int(14)] = (self).f27
                nw253_[int(15)] = ((self).f27 if d_1500_cf33_ else len(d_1507_v22_))
                nw253_[int(16)] = len(d_1508_v23_)
                nw253_[int(17)] = (0) - (len((_dafny.SeqWithoutIsStrInference([d_1501_v16_ for d_1511_i1_ in range(default__.abs(891))])) + (d_1509_v24_)))
                nw253_[int(18)] = (self).f27
                d_1510_v25_ = nw253_
                index195_ = default__.safeIndex(757, (d_1510_v25_).length(0))
                (d_1510_v25_)[index195_] = ((self).f27) - ((self).f27)
                d_1512_v27_: _dafny.Map
                def iife129_():
                    coll53_ = _dafny.Map()
                    compr_53_: int
                    for compr_53_ in _dafny.IntegerRange(593, -287):
                        d_1513_v26_: int = compr_53_
                        if ((593) <= (d_1513_v26_)) and ((d_1513_v26_) < (-287)):
                            coll53_[default__.safeDivisionInt(d_1513_v26_, (self).f27)] = d_1473_v0_
                    return _dafny.Map(coll53_)
                d_1512_v27_ = _dafny.Map({((self).f29).ispropersubset((self).f29): iife129_()
                })
                index196_ = default__.safeIndex(757, (d_1510_v25_).length(0))
                rhs273_ = (self).f27
                rhs274_ = len(((d_1509_v24_).set(default__.safeIndex((self).f27, len(d_1509_v24_)), d_1501_v16_)) + (d_1509_v24_))
                rhs275_ = len(d_1512_v27_)
                rhs276_ = ((d_1505_v20_)[d_1473_v0_] if (d_1473_v0_) in (d_1505_v20_) else True)
                lhs215_ = globalState
                lhs216_ = d_1510_v25_
                lhs217_ = default__.safeIndex(757, (d_1510_v25_).length(0))
                lhs218_ = globalState
                r0 = rhs273_
                lhs215_.f24 = rhs274_
                lhs216_[lhs217_] = rhs275_
                lhs218_.f26 = rhs276_
                d_1514_v28_: C1
                nw254_ = C1()
                nw254_.ctor__(d_1500_cf33_, (self).f29, (self).f30, (0) - ((0) - ((self).f27)), (self).f28)
                d_1514_v28_ = nw254_
                d_1515_v29_: _dafny.Seq
                d_1515_v29_ = _dafny.SeqWithoutIsStrInference([d_1509_v24_, d_1509_v24_])
                d_1516_v30_: _dafny.Map
                d_1516_v30_ = _dafny.Map({d_1500_cf33_: (d_1515_v29_).set(default__.safeIndex((d_1510_v25_)[default__.safeIndex(757, (d_1510_v25_).length(0))], len(d_1515_v29_)), _dafny.SeqWithoutIsStrInference([d_1501_v16_, d_1501_v16_, _dafny.CodePoint('u')]))})
                (d_1514_v28_).f35 = (((d_1516_v30_)[d_1473_v0_] if (d_1473_v0_) in (d_1516_v30_) else d_1515_v29_)) == ((d_1515_v29_) + (d_1515_v29_))
                d_1517_v31_: _dafny.Array
                nw255_ = _dafny.Array(False, 22)
                d_1517_v31_ = nw255_
                d_1518_v32_: _dafny.Map
                d_1518_v32_ = _dafny.Map({d_1517_v31_: False})
                d_1473_v0_ = (d_1517_v31_) not in (d_1518_v32_)
            elif source25_.is_DC20:
                d_1519___mcc_h2_ = source25_.cf34
                d_1520___mcc_h3_ = source25_.cf35
                d_1521___mcc_h4_ = source25_.cf36
                d_1522___mcc_h5_ = source25_.cf37
                d_1523___mcc_h6_ = source25_.cf38
                d_1524_cf38_ = d_1523___mcc_h6_
                d_1525_cf37_ = d_1522___mcc_h5_
                d_1526_cf36_ = d_1521___mcc_h4_
                d_1527_cf35_ = d_1520___mcc_h3_
                d_1528_cf34_ = d_1519___mcc_h2_
                d_1529_v33_: _dafny.Map
                d_1529_v33_ = _dafny.Map({d_1526_cf36_: d_1473_v0_})
                d_1530_v34_: _dafny.Map
                d_1530_v34_ = _dafny.Map({((d_1529_v33_)[d_1473_v0_] if (d_1473_v0_) in (d_1529_v33_) else d_1525_cf37_): (self).f27})
                d_1531_v35_: _dafny.Array
                nw256_ = _dafny.Array(int(0), 16)
                d_1531_v35_ = nw256_
                d_1532_v36_: _dafny.Map
                d_1532_v36_ = _dafny.Map({d_1530_v34_: d_1531_v35_})
                d_1533_v37_: _dafny.Array
                nw257_ = _dafny.Array(_dafny.MultiSet({}), 17)
                d_1533_v37_ = nw257_
                index197_ = default__.safeIndex(333, (d_1533_v37_).length(0))
                (d_1533_v37_)[index197_] = (self).f29
                d_1534_v38_: _dafny.Set
                d_1534_v38_ = _dafny.Set({d_1528_cf34_, d_1528_cf34_, (self).f27, (self).f27, d_1528_cf34_})
                index198_ = default__.safeIndex(333, (d_1533_v37_).length(0))
                rhs277_ = ((d_1474_v1_) + (d_1474_v1_)).set(default__.safeIndex((self).f27, len((d_1474_v1_) + (d_1474_v1_))), d_1525_cf37_)
                rhs278_ = (len((d_1474_v1_) + (d_1474_v1_))) >= (d_1528_cf34_)
                rhs279_ = ((d_1532_v36_).set(d_1530_v34_, d_1531_v35_)) | (d_1532_v36_)
                rhs280_ = (d_1534_v38_).issubset(d_1534_v38_)
                rhs281_ = (((self).f29) - (default__.fm32(globalState))) | (_dafny.MultiSet([False]))
                lhs219_ = globalState
                lhs220_ = d_1533_v37_
                lhs221_ = default__.safeIndex(333, (d_1533_v37_).length(0))
                d_1474_v1_ = rhs277_
                lhs219_.f26 = rhs278_
                d_1532_v36_ = rhs279_
                d_1525_cf37_ = rhs280_
                lhs220_[lhs221_] = rhs281_
                d_1535_v39_: _dafny.Set
                d_1535_v39_ = _dafny.Set({d_1526_cf36_, d_1527_cf35_, d_1527_cf35_})
                d_1527_cf35_ = not((d_1535_v39_).ispropersubset(d_1535_v39_))
                d_1536_v40_: str
                d_1536_v40_ = _dafny.CodePoint('y')
                d_1537_v41_: D4
                d_1537_v41_ = D4_DC9(d_1528_cf34_, d_1527_cf35_)
                d_1538_v42_: _dafny.Set
                d_1538_v42_ = _dafny.Set({d_1537_v41_, d_1537_v41_})
                d_1539_v43_: D1
                d_1539_v43_ = D1_DC2(d_1536_v40_, len(d_1538_v42_))
                (globalState).f18 = (((d_1533_v37_)[default__.safeIndex(333, (d_1533_v37_).length(0))])[d_1526_cf36_] if (d_1526_cf36_) in ((d_1533_v37_)[default__.safeIndex(333, (d_1533_v37_).length(0))]) else (d_1539_v43_).cf3)
                (globalState).f18 = (d_1528_cf34_) - ((self).f27)
            elif True:
                d_1540___mcc_h7_ = source25_.cf31
                d_1541_cf31_ = d_1540___mcc_h7_
                d_1542_v44_: C0
                nw258_ = C0()
                nw258_.ctor__((0) - (len((self).f30)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tidvowswu")))
                d_1542_v44_ = nw258_
                d_1543_v45_: D6
                d_1543_v45_ = D6_DC11(d_1542_v44_)
                pat_let_tv32_ = d_1542_v44_
                d_1544_v46_: _dafny.Array
                nw259_ = _dafny.Array(None, 4)
                nw259_[int(0)] = d_1543_v45_
                nw259_[int(1)] = d_1543_v45_
                nw259_[int(2)] = D6_DC11(d_1542_v44_)
                def iife130_(_pat_let38_0):
                    def iife131_(d_1545_dt__update__tmp_h2_):
                        def iife132_(_pat_let39_0):
                            def iife133_(d_1546_dt__update_hcf19_h0_):
                                return D6_DC11(d_1546_dt__update_hcf19_h0_)
                            return iife133_(_pat_let39_0)
                        return iife132_(pat_let_tv32_)
                    return iife131_(_pat_let38_0)
                nw259_[int(3)] = iife130_(D6_DC11(d_1542_v44_))
                d_1544_v46_ = nw259_
                pat_let_tv33_ = d_1542_v44_
                index199_ = default__.safeIndex(259, (d_1544_v46_).length(0))
                def iife134_(_pat_let40_0):
                    def iife135_(d_1547_dt__update__tmp_h3_):
                        def iife136_(_pat_let41_0):
                            def iife137_(d_1548_dt__update_hcf19_h1_):
                                return D6_DC11(d_1548_dt__update_hcf19_h1_)
                            return iife137_(_pat_let41_0)
                        return iife136_(pat_let_tv33_)
                    return iife135_(_pat_let40_0)
                (d_1544_v46_)[index199_] = iife134_(d_1543_v45_)
                pat_let_tv34_ = d_1542_v44_
                index200_ = default__.safeIndex(259, (d_1544_v46_).length(0))
                def iife138_(_pat_let42_0):
                    def iife139_(d_1549_dt__update__tmp_h4_):
                        def iife140_(_pat_let43_0):
                            def iife141_(d_1550_dt__update_hcf19_h2_):
                                return D6_DC11(d_1550_dt__update_hcf19_h2_)
                            return iife141_(_pat_let43_0)
                        return iife140_(pat_let_tv34_)
                    return iife139_(_pat_let42_0)
                (d_1544_v46_)[index200_] = iife138_(d_1543_v45_)
                (globalState).f26 = not(d_1473_v0_)
                (globalState).f0 = ((self).f27) * ((self).f27)
                d_1551_v47_: C0
                nw260_ = C0()
                nw260_.ctor__((d_1542_v44_).f36, ((d_1542_v44_).f37 if not(d_1473_v0_) else (d_1542_v44_).f37))
                d_1551_v47_ = nw260_
            if d_1473_v0_:
                d_1552_v48_: _dafny.Seq
                d_1552_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aljehft"))
                d_1553_v49_: C0
                nw261_ = C0()
                nw261_.ctor__(len(d_1552_v48_), d_1552_v48_)
                d_1553_v49_ = nw261_
                d_1554_v50_: D6
                d_1554_v50_ = D6_DC11(d_1553_v49_)
                pat_let_tv35_ = d_1553_v49_
                def iife142_(_pat_let44_0):
                    def iife143_(d_1555_dt__update__tmp_h5_):
                        def iife144_(_pat_let45_0):
                            def iife145_(d_1556_dt__update_hcf19_h3_):
                                return D6_DC11(d_1556_dt__update_hcf19_h3_)
                            return iife145_(_pat_let45_0)
                        return iife144_(pat_let_tv35_)
                    return iife143_(_pat_let44_0)
                d_1554_v50_ = iife142_(d_1554_v50_)
                (globalState).f26 = not(d_1473_v0_)
                d_1557_v51_: _dafny.Array
                def lambda61_(d_1558_v49_):
                    def lambda62_(d_1559_i2_):
                        return default__.safeDivisionInt(d_1559_i2_, (d_1558_v49_).f36)

                    return lambda62_

                init29_ = lambda61_(d_1553_v49_)
                nw262_ = _dafny.Array(None, 5)
                for i0_29_ in range(nw262_.length(0)):
                    nw262_[i0_29_] = init29_(i0_29_)
                d_1557_v51_ = nw262_
                d_1560_v52_: _dafny.Map
                d_1560_v52_ = _dafny.Map({(d_1477_v4_).cf34: d_1557_v51_})
                d_1561_v53_: _dafny.Seq
                d_1561_v53_ = _dafny.SeqWithoutIsStrInference([d_1557_v51_, d_1557_v51_, d_1557_v51_, d_1557_v51_, d_1557_v51_])
                d_1560_v52_ = (d_1560_v52_).set((d_1553_v49_).f36, (d_1561_v53_)[default__.safeIndex((d_1553_v49_).f36, len(d_1561_v53_))])
                d_1562_v54_: str
                d_1562_v54_ = _dafny.CodePoint('r')
                d_1563_v55_: D4
                d_1563_v55_ = D4_DC8(d_1562_v54_, (self).f27, False)
                d_1564_v56_: _dafny.Map
                d_1564_v56_ = _dafny.Map({d_1473_v0_: (self).f27})
                d_1565_v57_: _dafny.Map
                d_1565_v57_ = _dafny.Map({(_dafny.MultiSet([(d_1563_v55_).cf15, d_1473_v0_])).isdisjoint((self).f29): len(d_1564_v56_)})
                d_1564_v56_ = (d_1564_v56_).set(d_1473_v0_, (0) - ((self).f27))
                (globalState).f26 = (self).fm3((self).f29, globalState)
            elif True:
                d_1566_v58_: _dafny.Array
                nw263_ = _dafny.Array(False, 28)
                d_1566_v58_ = nw263_
                d_1567_v59_: _dafny.Seq
                d_1567_v59_ = _dafny.SeqWithoutIsStrInference([d_1566_v58_])
                d_1568_v60_: C0
                nw264_ = C0()
                nw264_.ctor__(len(d_1567_v59_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwnfm")))
                d_1568_v60_ = nw264_
                d_1569_v61_: str
                d_1569_v61_ = _dafny.CodePoint('k')
                d_1569_v61_ = d_1569_v61_
                d_1473_v0_ = d_1473_v0_
                d_1570_v62_: _dafny.Set
                d_1570_v62_ = _dafny.Set({d_1473_v0_})
                d_1571_v63_: _dafny.Set
                d_1571_v63_ = _dafny.Set({d_1570_v62_})
                (globalState).f26 = not(not((len(d_1571_v63_)) > ((d_1568_v60_).f36)))
                (globalState).f21 = (d_1568_v60_).f36
            d_1572_v64_: D10
            d_1572_v64_ = D10_DC22(((self).f27) - (((self).f29).cardinality), ((self).f29) | ((self).f29))
            d_1572_v64_ = d_1572_v64_
            d_1573_v65_: _dafny.Map
            d_1573_v65_ = _dafny.Map({(self).f27: d_1473_v0_})
            d_1573_v65_ = ((d_1573_v65_) | (d_1573_v65_)) | ((d_1573_v65_) | (d_1573_v65_))
            d_1574_v66_: _dafny.Seq
            d_1574_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxncklgcj"))
            d_1575_v67_: D6
            d_1575_v67_ = D6_DC12(True, _dafny.SeqWithoutIsStrInference([d_1474_v1_, d_1474_v1_, d_1474_v1_]), d_1574_v66_)
            d_1576_v68_: _dafny.Map
            d_1576_v68_ = _dafny.Map({d_1575_v67_: True})
            d_1577_v69_: _dafny.Array
            def lambda63_(d_1578_i3_):
                return False

            init30_ = lambda63_
            nw265_ = _dafny.Array(None, 4)
            for i0_30_ in range(nw265_.length(0)):
                nw265_[i0_30_] = init30_(i0_30_)
            d_1577_v69_ = nw265_
            d_1579_v70_: _dafny.Map
            d_1579_v70_ = _dafny.Map({True: d_1473_v0_})
            rhs282_ = (self).f27
            rhs283_ = d_1576_v68_
            rhs284_ = d_1577_v69_
            rhs285_ = ((d_1473_v0_) not in (d_1579_v70_)) or (d_1473_v0_)
            rhs286_ = d_1473_v0_
            lhs222_ = globalState
            lhs223_ = globalState
            r2 = rhs282_
            d_1576_v68_ = rhs283_
            d_1577_v69_ = rhs284_
            lhs222_.f26 = rhs285_
            lhs223_.f26 = rhs286_
        elif True:
            (globalState).f26 = d_1473_v0_
            d_1474_v1_ = d_1474_v1_
            d_1580_v71_: _dafny.Map
            d_1580_v71_ = _dafny.Map({default__.fm0((self).f27, globalState): d_1473_v0_})
            d_1581_v72_: _dafny.Seq
            d_1581_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukkh"))
            d_1582_v73_: C0
            nw266_ = C0()
            nw266_.ctor__(len(d_1580_v71_), d_1581_v72_)
            d_1582_v73_ = nw266_
            d_1583_v74_: _dafny.Map
            d_1583_v74_ = _dafny.Map({(d_1582_v73_).f36: d_1582_v73_})
            d_1584_v75_: _dafny.Array
            nw267_ = _dafny.Array(None, 22)
            nw267_[int(0)] = d_1582_v73_
            nw267_[int(1)] = d_1582_v73_
            nw267_[int(2)] = d_1582_v73_
            nw267_[int(3)] = d_1582_v73_
            nw267_[int(4)] = ((d_1583_v74_)[(0) - ((self).f27)] if ((0) - ((self).f27)) in (d_1583_v74_) else d_1582_v73_)
            nw267_[int(5)] = d_1582_v73_
            nw267_[int(6)] = d_1582_v73_
            nw267_[int(7)] = d_1582_v73_
            nw267_[int(8)] = d_1582_v73_
            nw267_[int(9)] = d_1582_v73_
            nw267_[int(10)] = d_1582_v73_
            nw267_[int(11)] = d_1582_v73_
            nw267_[int(12)] = d_1582_v73_
            nw267_[int(13)] = d_1582_v73_
            nw267_[int(14)] = d_1582_v73_
            nw267_[int(15)] = d_1582_v73_
            nw267_[int(16)] = d_1582_v73_
            nw267_[int(17)] = d_1582_v73_
            nw267_[int(18)] = d_1582_v73_
            nw267_[int(19)] = d_1582_v73_
            nw267_[int(20)] = d_1582_v73_
            nw267_[int(21)] = d_1582_v73_
            d_1584_v75_ = nw267_
            index201_ = default__.safeIndex(986, (d_1584_v75_).length(0))
            (d_1584_v75_)[index201_] = d_1582_v73_
            d_1585_v76_: D2
            d_1585_v76_ = D2_DC4((self).f27, len(_dafny.Set({d_1473_v0_})), (self).f27, _dafny.CodePoint('s'), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iga")))
            index202_ = default__.safeIndex(986, (d_1584_v75_).length(0))
            rhs287_ = ((self).f27) + ((0) - ((d_1582_v73_).f36))
            rhs288_ = d_1582_v73_
            rhs289_ = len((d_1585_v76_).cf9)
            lhs224_ = globalState
            lhs225_ = d_1584_v75_
            lhs226_ = default__.safeIndex(986, (d_1584_v75_).length(0))
            lhs227_ = globalState
            lhs224_.f15 = rhs287_
            lhs225_[lhs226_] = rhs288_
            lhs227_.f3 = rhs289_
            d_1586_v77_: str
            d_1586_v77_ = _dafny.CodePoint('e')
            d_1586_v77_ = d_1586_v77_
            (globalState).f21 = (self).f27
        d_1587_v78_: _dafny.Seq
        d_1587_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odaravso"))
        d_1588_v79_: _dafny.Map
        d_1588_v79_ = _dafny.Map({(0) - ((self).f27): d_1473_v0_})
        d_1589_v80_: C2
        nw268_ = C2()
        nw268_.ctor__((self).f29, (self).f27)
        d_1589_v80_ = nw268_
        d_1590_v81_: D12
        d_1590_v81_ = D12_DC26(d_1589_v80_)
        d_1591_v82_: _dafny.Set
        d_1591_v82_ = _dafny.Set({d_1589_v80_, d_1589_v80_, d_1589_v80_, (d_1590_v81_).cf46, d_1589_v80_})
        d_1592_v83_: _dafny.MultiSet
        d_1592_v83_ = _dafny.MultiSet([(self).f27, default__.fm1(globalState)])
        d_1593_v84_: _dafny.Array
        nw269_ = _dafny.Array(None, 13)
        nw269_[int(0)] = (d_1474_v1_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yspuqcihq"))), len(d_1474_v1_))]
        nw269_[int(1)] = d_1473_v0_
        nw269_[int(2)] = d_1473_v0_
        nw269_[int(3)] = (d_1587_v78_) < (d_1587_v78_)
        nw269_[int(4)] = d_1473_v0_
        nw269_[int(5)] = d_1473_v0_
        nw269_[int(6)] = d_1473_v0_
        nw269_[int(7)] = False
        nw269_[int(8)] = ((self).f27) < ((self).f27)
        nw269_[int(9)] = ((self).f27) < ((self).f27)
        nw269_[int(10)] = (default__.fm33(d_1588_v79_, len(d_1591_v82_), len(d_1474_v1_), (d_1592_v83_).cardinality, globalState)) == (d_1474_v1_)
        nw269_[int(11)] = d_1473_v0_
        nw269_[int(12)] = True
        d_1593_v84_ = nw269_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1593_v84_).length(0)):
            d_1594_i4_: int = guard_loop_6_
            if (True) and (((0) <= (d_1594_i4_)) and ((d_1594_i4_) < ((d_1593_v84_).length(0)))):
                (d_1593_v84_)[(d_1594_i4_)] = not((175) > (871))
        d_1595_v85_: _dafny.Map
        d_1595_v85_ = _dafny.Map({(self).f27: d_1593_v84_})
        d_1596_v86_: _dafny.Map
        d_1596_v86_ = _dafny.Map({d_1473_v0_: (self).f27})
        d_1595_v85_ = (d_1595_v85_).set(((d_1596_v86_)[default__.fm0((self).f27, globalState)] if (default__.fm0((self).f27, globalState)) in (d_1596_v86_) else len(d_1588_v79_)), d_1593_v84_)
        d_1597_v87_: str
        d_1597_v87_ = _dafny.CodePoint('k')
        d_1598_v88_: _dafny.Map
        d_1598_v88_ = _dafny.Map({d_1473_v0_: d_1597_v87_})
        d_1598_v88_ = (d_1598_v88_).set(d_1473_v0_, d_1597_v87_)
        hi6_ = 18
        for d_1599_i5_ in range((self).f27, hi6_):
            (globalState).f18 = (0) - (-374)
            (globalState).f26 = (420) != (871)
            if (not((_dafny.MultiSet([d_1473_v0_])).ispropersubset((_dafny.MultiSet([d_1473_v0_, d_1473_v0_])).set(False, default__.abs(d_1599_i5_)))) if (self).fm3((self).f29, globalState) else d_1473_v0_):
                d_1473_v0_ = not(True)
                rhs290_ = d_1473_v0_
                rhs291_ = d_1587_v78_
                d_1473_v0_ = rhs290_
                d_1587_v78_ = rhs291_
                d_1600_v89_: _dafny.Array
                nw270_ = _dafny.Array(int(0), 4)
                d_1600_v89_ = nw270_
                index203_ = default__.safeIndex(178, (d_1600_v89_).length(0))
                (d_1600_v89_)[index203_] = d_1599_i5_
                index204_ = default__.safeIndex(164, (d_1593_v84_).length(0))
                (d_1593_v84_)[index204_] = d_1473_v0_
                index205_ = default__.safeIndex(178, (d_1600_v89_).length(0))
                index206_ = default__.safeIndex(164, (d_1593_v84_).length(0))
                rhs292_ = default__.safeModuloInt((self).f27, 316)
                rhs293_ = (d_1599_i5_) == ((self).f27)
                rhs294_ = d_1593_v84_
                rhs295_ = default__.safeModuloInt(d_1599_i5_, len(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])))
                rhs296_ = (_dafny.MultiSet([(self).f27, (self).f27, d_1599_i5_])).ispropersubset(_dafny.MultiSet([(self).f27, d_1599_i5_]))
                lhs228_ = d_1600_v89_
                lhs229_ = default__.safeIndex(178, (d_1600_v89_).length(0))
                lhs230_ = d_1593_v84_
                lhs231_ = default__.safeIndex(164, (d_1593_v84_).length(0))
                lhs228_[lhs229_] = rhs292_
                d_1473_v0_ = rhs293_
                d_1593_v84_ = rhs294_
                r2 = rhs295_
                lhs230_[lhs231_] = rhs296_
                d_1601_v90_: _dafny.MultiSet
                d_1601_v90_ = d_1592_v83_
                d_1602_v91_: _dafny.Map
                d_1602_v91_ = _dafny.Map({d_1601_v90_: (d_1593_v84_)[default__.safeIndex(164, (d_1593_v84_).length(0))]})
                d_1603_v92_: _dafny.Seq
                d_1603_v92_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1601_v90_: (d_1593_v84_)[default__.safeIndex(164, (d_1593_v84_).length(0))]}), _dafny.Map({d_1601_v90_: (d_1593_v84_)[default__.safeIndex(164, (d_1593_v84_).length(0))]})])
                d_1604_v93_: _dafny.Set
                d_1604_v93_ = _dafny.Set({(d_1593_v84_)[default__.safeIndex(164, (d_1593_v84_).length(0))]})
                d_1602_v91_ = (d_1603_v92_)[default__.safeIndex((len(d_1604_v93_)) + (d_1599_i5_), len(d_1603_v92_))]
                (globalState).f3 = (self).f27
            elif True:
                (globalState).f0 = d_1599_i5_
                d_1605_v94_: _dafny.Array
                def lambda64_(d_1606_v1_):
                    def lambda65_(d_1607_i6_):
                        return (d_1607_i6_) - (len(d_1606_v1_))

                    return lambda65_

                init31_ = lambda64_(d_1474_v1_)
                nw271_ = _dafny.Array(None, 21)
                for i0_31_ in range(nw271_.length(0)):
                    nw271_[i0_31_] = init31_(i0_31_)
                d_1605_v94_ = nw271_
                index207_ = default__.safeIndex(562, (d_1605_v94_).length(0))
                (d_1605_v94_)[index207_] = 311
                index208_ = default__.safeIndex(562, (d_1605_v94_).length(0))
                (d_1605_v94_)[index208_] = d_1599_i5_
                d_1608_v95_: _dafny.Seq
                d_1608_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hfpe"))
                d_1609_v96_: C0
                nw272_ = C0()
                nw272_.ctor__(((self).f29).cardinality, (d_1608_v95_))
                d_1609_v96_ = nw272_
                d_1610_v97_: _dafny.Map
                d_1610_v97_ = _dafny.Map({d_1473_v0_: d_1609_v96_})
                d_1611_v98_: C2
                nw273_ = C2()
                nw273_.ctor__((self).f29, len(d_1610_v97_))
                d_1611_v98_ = nw273_
                d_1587_v78_ = (d_1609_v96_).f37
                d_1612_v99_: _dafny.Map
                d_1612_v99_ = _dafny.Map({(d_1605_v94_)[default__.safeIndex(562, (d_1605_v94_).length(0))]: (d_1605_v94_)[default__.safeIndex(562, (d_1605_v94_).length(0))]})
                d_1613_v100_: _dafny.Set
                d_1613_v100_ = _dafny.Set({(d_1589_v80_).fm21(globalState), d_1473_v0_})
                d_1614_v101_: T0
                nw274_ = C3()
                nw274_.ctor__(len(d_1613_v100_), (self).f28)
                d_1614_v101_ = nw274_
                d_1615_v102_: _dafny.Set
                d_1615_v102_ = _dafny.Set({d_1614_v101_, d_1614_v101_})
                d_1616_v103_: _dafny.Map
                d_1616_v103_ = _dafny.Map({d_1599_i5_: d_1596_v86_})
                d_1612_v99_ = (d_1612_v99_).set(len((_dafny.Map({d_1599_i5_: _dafny.Map({d_1473_v0_: len(d_1615_v102_)})})) | (d_1616_v103_)), d_1599_i5_)
            d_1617_v104_: D4
            d_1617_v104_ = D4_DC7(d_1593_v84_)
            d_1617_v104_ = d_1617_v104_
        hi7_ = (self).f27
        for d_1618_i7_ in range((0) - ((self).f27), hi7_):
            d_1619_v105_: _dafny.Array
            nw275_ = _dafny.Array(_dafny.Map({}), 9)
            d_1619_v105_ = nw275_
            index209_ = default__.safeIndex(289, (d_1619_v105_).length(0))
            (d_1619_v105_)[index209_] = (d_1588_v79_) | (d_1588_v79_)
            index210_ = default__.safeIndex(289, (d_1619_v105_).length(0))
            (d_1619_v105_)[index210_] = d_1588_v79_
            d_1620_v106_: _dafny.Array
            def lambda66_(d_1621_v0_):
                def lambda67_(d_1622_i8_):
                    return (_dafny.SeqWithoutIsStrInference([(self).f30]) if d_1621_v0_ else _dafny.SeqWithoutIsStrInference([(self).f30 for d_1623_i9_ in range(default__.abs(-529))]))

                return lambda67_

            init32_ = lambda66_(d_1473_v0_)
            nw276_ = _dafny.Array(None, 27)
            for i0_32_ in range(nw276_.length(0)):
                nw276_[i0_32_] = init32_(i0_32_)
            d_1620_v106_ = nw276_
            d_1624_v107_: _dafny.Seq
            d_1624_v107_ = _dafny.SeqWithoutIsStrInference([(self).f30, (self).f30])
            index211_ = default__.safeIndex(991, (d_1620_v106_).length(0))
            (d_1620_v106_)[index211_] = d_1624_v107_
            index212_ = default__.safeIndex(991, (d_1620_v106_).length(0))
            (d_1620_v106_)[index212_] = (d_1624_v107_) + (d_1624_v107_)
            d_1625_v108_: _dafny.Seq
            d_1625_v108_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1618_i7_: d_1473_v0_})])
            index213_ = default__.safeIndex(289, (d_1619_v105_).length(0))
            (d_1619_v105_)[index213_] = (d_1625_v108_)[default__.safeIndex(d_1618_i7_, len(d_1625_v108_))]
            r0 = (self).f27
        d_1626_v109_: _dafny.Map
        d_1626_v109_ = _dafny.Map({((self).f27) + ((self).f27): 53})
        r0 = ((d_1626_v109_)[(self).f27] if ((self).f27) in (d_1626_v109_) else (0) - ((self).f27))
        d_1627_v110_: D9
        d_1627_v110_ = D9_DC20((self).f27, True, d_1473_v0_, False, d_1626_v109_)
        d_1628_v111_: _dafny.Array
        def lambda68_(d_1629_i10_):
            return (d_1629_i10_) * ((0) - ((self).f27))

        init33_ = lambda68_
        nw277_ = _dafny.Array(None, 13)
        for i0_33_ in range(nw277_.length(0)):
            nw277_[i0_33_] = init33_(i0_33_)
        d_1628_v111_ = nw277_
        d_1630_v112_: D12
        d_1630_v112_ = D12_DC27(d_1473_v0_, d_1592_v83_, len(d_1587_v78_), d_1628_v111_)
        pat_let_tv36_ = d_1473_v0_
        d_1631_v113_: _dafny.Seq
        def iife146_(_pat_let46_0):
            def iife147_(d_1632_dt__update__tmp_h6_):
                def iife148_(_pat_let47_0):
                    def iife149_(d_1633_dt__update_hcf47_h0_):
                        return D12_DC27(d_1633_dt__update_hcf47_h0_, (d_1632_dt__update__tmp_h6_).cf48, (d_1632_dt__update__tmp_h6_).cf49, (d_1632_dt__update__tmp_h6_).cf50)
                    return iife149_(_pat_let47_0)
                return iife148_(pat_let_tv36_)
            return iife147_(_pat_let46_0)
        d_1631_v113_ = _dafny.SeqWithoutIsStrInference([iife146_(D12_DC27(d_1473_v0_, d_1592_v83_, 468, d_1628_v111_)), d_1630_v112_])
        r1 = default__.safeDivisionInt(((self).f30)[default__.safeIndex((d_1627_v110_).cf34, len((self).f30))], default__.safeDivisionInt((self).f27, ((_dafny.MultiSet(d_1631_v113_)).set(d_1630_v112_, default__.abs((self).f27))).cardinality))
        r2 = ((self).f27 if d_1473_v0_ else len((self).f30))
        return r0, r1, r2

    def m9(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        r2: bool = False
        d_1634_v0_: bool
        d_1634_v0_ = False
        if (((self).f27) + (len((self).fm2(d_1634_v0_, d_1634_v0_, d_1634_v0_, globalState)))) >= ((self).f27):
            d_1635_v1_: _dafny.Seq
            d_1635_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            d_1636_v2_: bool
            d_1637_v3_: int
            out44_: bool
            out45_: int
            out44_, out45_ = default__.m0(True, len(d_1635_v1_), ((self).f27) * ((self).f27), (self).f27, globalState)
            d_1636_v2_ = out44_
            d_1637_v3_ = out45_
            d_1638_v4_: C2
            nw278_ = C2()
            nw278_.ctor__((self).f29, (self).f27)
            d_1638_v4_ = nw278_
            d_1639_v5_: D12
            d_1639_v5_ = D12_DC29(D12_DC26(d_1638_v4_))
            d_1640_v6_: _dafny.Map
            d_1640_v6_ = _dafny.Map({d_1639_v5_: d_1636_v2_})
            d_1641_v7_: _dafny.Map
            d_1641_v7_ = _dafny.Map({d_1636_v2_: d_1634_v0_})
            d_1640_v6_ = (d_1640_v6_).set(d_1639_v5_, ((d_1641_v7_)[d_1636_v2_] if (d_1636_v2_) in (d_1641_v7_) else d_1634_v0_))
            d_1642_v8_: _dafny.Array
            nw279_ = _dafny.Array(int(0), 2)
            d_1642_v8_ = nw279_
            d_1643_v9_: _dafny.Seq
            d_1643_v9_ = _dafny.SeqWithoutIsStrInference([d_1642_v8_])
            if (d_1643_v9_) < (d_1643_v9_):
                d_1644_v10_: _dafny.Array
                def lambda69_(d_1645_v0_, d_1646_v2_):
                    def lambda70_(d_1647_i0_):
                        return _dafny.SeqWithoutIsStrInference([d_1645_v0_, d_1645_v0_, d_1646_v2_])

                    return lambda70_

                init34_ = lambda69_(d_1634_v0_, d_1636_v2_)
                nw280_ = _dafny.Array(None, 26)
                for i0_34_ in range(nw280_.length(0)):
                    nw280_[i0_34_] = init34_(i0_34_)
                d_1644_v10_ = nw280_
                d_1648_v11_: _dafny.Seq
                d_1648_v11_ = _dafny.SeqWithoutIsStrInference([not((d_1638_v4_).fm21(globalState))])
                index214_ = default__.safeIndex(910, (d_1644_v10_).length(0))
                (d_1644_v10_)[index214_] = d_1648_v11_
                index215_ = default__.safeIndex(910, (d_1644_v10_).length(0))
                (d_1644_v10_)[index215_] = d_1648_v11_
                d_1649_v12_: C3
                nw281_ = C3()
                nw281_.ctor__((self).f27, (self).f28)
                d_1649_v12_ = nw281_
                d_1650_v13_: _dafny.Map
                d_1650_v13_ = _dafny.Map({(not(d_1634_v0_) if False else d_1634_v0_): default__.safeDivisionInt((self).f27, (self).f27)})
                d_1650_v13_ = d_1650_v13_
                index216_ = default__.safeIndex(788, (d_1642_v8_).length(0))
                (d_1642_v8_)[index216_] = len(d_1648_v11_)
                d_1651_v14_: _dafny.Array
                nw282_ = _dafny.Array(None, 1)
                nw282_[int(0)] = d_1634_v0_
                d_1651_v14_ = nw282_
                index217_ = default__.safeIndex(788, (d_1642_v8_).length(0))
                rhs297_ = d_1634_v0_
                rhs298_ = d_1651_v14_
                rhs299_ = d_1637_v3_
                rhs300_ = ((520) >= (len(d_1635_v1_))) or ((d_1648_v11_) <= ((d_1644_v10_)[default__.safeIndex(910, (d_1644_v10_).length(0))]))
                lhs232_ = globalState
                lhs233_ = d_1642_v8_
                lhs234_ = default__.safeIndex(788, (d_1642_v8_).length(0))
                lhs235_ = globalState
                lhs232_.f26 = rhs297_
                r0 = rhs298_
                lhs233_[lhs234_] = rhs299_
                lhs235_.f26 = rhs300_
                d_1652_v15_: _dafny.Set
                d_1652_v15_ = _dafny.Set({d_1651_v14_})
                d_1653_v16_: D12
                d_1653_v16_ = D12_DC28(507, len(d_1652_v15_), d_1636_v2_)
                d_1654_v17_: _dafny.Map
                d_1654_v17_ = _dafny.Map({d_1653_v16_: d_1635_v1_})
                d_1654_v17_ = (d_1654_v17_).set(d_1653_v16_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fcnwjnotv"))) + (d_1635_v1_))
            elif True:
                d_1655_v18_: _dafny.Array
                nw283_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_1655_v18_ = nw283_
                d_1655_v18_ = d_1655_v18_
                (globalState).f24 = (self).f27
                nw284_ = _dafny.Array(False, 25)
                r0 = nw284_
                r2 = d_1634_v0_
                r1 = d_1636_v2_
            d_1641_v7_ = (d_1641_v7_).set(d_1634_v0_, d_1636_v2_)
            d_1656_v19_: _dafny.Seq
            d_1656_v19_ = _dafny.SeqWithoutIsStrInference([d_1636_v2_, True, False, d_1636_v2_])
            if (d_1656_v19_)[default__.safeIndex((self).f27, len(d_1656_v19_))]:
                d_1657_v20_: str
                d_1657_v20_ = _dafny.CodePoint('x')
                d_1658_v21_: _dafny.Array
                nw285_ = _dafny.Array(None, 13)
                nw285_[int(0)] = d_1657_v20_
                nw285_[int(1)] = d_1657_v20_
                nw285_[int(2)] = d_1657_v20_
                nw285_[int(3)] = d_1657_v20_
                nw285_[int(4)] = (d_1635_v1_)[default__.safeIndex((self).f27, len(d_1635_v1_))]
                nw285_[int(5)] = d_1657_v20_
                nw285_[int(6)] = d_1657_v20_
                nw285_[int(7)] = d_1657_v20_
                nw285_[int(8)] = d_1657_v20_
                nw285_[int(9)] = default__.fm38(d_1641_v7_, globalState)
                nw285_[int(10)] = d_1657_v20_
                nw285_[int(11)] = d_1657_v20_
                nw285_[int(12)] = (D1_DC2(d_1657_v20_, d_1637_v3_)).cf2
                d_1658_v21_ = nw285_
                d_1659_v22_: _dafny.Map
                d_1659_v22_ = _dafny.Map({D13_DC31(d_1637_v3_): d_1658_v21_})
                d_1660_v23_: D13
                d_1660_v23_ = D13_DC31(d_1637_v3_)
                d_1661_v24_: _dafny.Array
                nw286_ = _dafny.Array(None, 12)
                nw286_[int(0)] = d_1658_v21_
                nw286_[int(1)] = d_1658_v21_
                nw286_[int(2)] = d_1658_v21_
                nw286_[int(3)] = d_1658_v21_
                nw286_[int(4)] = d_1658_v21_
                nw286_[int(5)] = ((d_1659_v22_)[d_1660_v23_] if (d_1660_v23_) in (d_1659_v22_) else d_1658_v21_)
                nw286_[int(6)] = d_1658_v21_
                nw286_[int(7)] = d_1658_v21_
                nw286_[int(8)] = d_1658_v21_
                nw286_[int(9)] = d_1658_v21_
                nw286_[int(10)] = d_1658_v21_
                nw286_[int(11)] = d_1658_v21_
                d_1661_v24_ = nw286_
                d_1662_v25_: _dafny.Seq
                d_1662_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, d_1636_v2_, d_1634_v0_]), d_1656_v19_, d_1656_v19_])
                d_1663_v26_: D6
                d_1663_v26_ = D6_DC12(d_1636_v2_, d_1662_v25_, d_1635_v1_)
                d_1664_v27_: _dafny.Seq
                d_1664_v27_ = _dafny.SeqWithoutIsStrInference([d_1663_v26_])
                d_1665_v28_: _dafny.Set
                d_1665_v28_ = _dafny.Set({d_1635_v1_, d_1635_v1_, d_1635_v1_})
                rhs301_ = ((self).fm4(d_1665_v28_, d_1636_v2_, _dafny.SeqWithoutIsStrInference([d_1634_v0_]), globalState)) * ((d_1637_v3_) + (d_1637_v3_))
                rhs302_ = d_1661_v24_
                rhs303_ = d_1636_v2_
                rhs304_ = _dafny.SeqWithoutIsStrInference([d_1663_v26_ for d_1666_i1_ in range(default__.abs(906))])
                rhs305_ = len(d_1635_v1_)
                lhs236_ = globalState
                lhs237_ = globalState
                lhs236_.f0 = rhs301_
                d_1661_v24_ = rhs302_
                d_1634_v0_ = rhs303_
                d_1664_v27_ = rhs304_
                lhs237_.f21 = rhs305_
                d_1667_v29_: _dafny.MultiSet
                d_1667_v29_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtdcmrxwf"))])
                (globalState).f3 = ((d_1667_v29_).set(d_1635_v1_, default__.abs((self).f27))).cardinality
                d_1668_v30_: C3
                nw287_ = C3()
                nw287_.ctor__(918, (self).f28)
                d_1668_v30_ = nw287_
                d_1669_v31_: T0
                nw288_ = C3()
                nw288_.ctor__(d_1637_v3_, (self).f28)
                d_1669_v31_ = nw288_
                d_1670_v32_: _dafny.Map
                d_1670_v32_ = _dafny.Map({d_1669_v31_: _dafny.CodePoint('k')})
                d_1670_v32_ = d_1670_v32_
                d_1671_v33_: C3
                nw289_ = C3()
                nw289_.ctor__((self).f27, (d_1669_v31_).f28)
                d_1671_v33_ = nw289_
            elif True:
                (globalState).f21 = d_1637_v3_
                d_1672_v34_: _dafny.Seq
                d_1672_v34_ = _dafny.SeqWithoutIsStrInference([d_1656_v19_])
                d_1673_v35_: D6
                d_1673_v35_ = D6_DC12(d_1634_v0_, d_1672_v34_, d_1635_v1_)
                (globalState).f26 = (((d_1673_v35_).cf21) + (d_1672_v34_)) <= (d_1672_v34_)
                d_1674_v36_: C3
                nw290_ = C3()
                nw290_.ctor__((0) - (d_1637_v3_), (self).f28)
                d_1674_v36_ = nw290_
                d_1675_v37_: str
                d_1675_v37_ = _dafny.CodePoint('k')
                rhs306_ = d_1634_v0_
                rhs307_ = d_1675_v37_
                rhs308_ = (self).f27
                lhs238_ = globalState
                d_1636_v2_ = rhs306_
                d_1675_v37_ = rhs307_
                lhs238_.f0 = rhs308_
                d_1676_v38_: _dafny.Map
                d_1676_v38_ = _dafny.Map({d_1675_v37_: d_1634_v0_})
                d_1677_v39_: _dafny.Map
                d_1677_v39_ = _dafny.Map({d_1637_v3_: d_1675_v37_})
                d_1678_v40_: _dafny.Seq
                d_1678_v40_ = _dafny.SeqWithoutIsStrInference([((self).f29).set(d_1636_v2_, default__.abs((self).f27))])
                d_1676_v38_ = (d_1676_v38_).set(((d_1677_v39_)[d_1637_v3_] if (d_1637_v3_) in (d_1677_v39_) else d_1675_v37_), ((d_1678_v40_)[default__.safeIndex((self).f27, len(d_1678_v40_))]).isdisjoint((self).f29))
        elif True:
            d_1679_v41_: _dafny.Array
            nw291_ = _dafny.Array(None, 23)
            nw291_[int(0)] = d_1634_v0_
            nw291_[int(1)] = d_1634_v0_
            nw291_[int(2)] = d_1634_v0_
            nw291_[int(3)] = d_1634_v0_
            nw291_[int(4)] = d_1634_v0_
            nw291_[int(5)] = d_1634_v0_
            nw291_[int(6)] = not(d_1634_v0_)
            nw291_[int(7)] = d_1634_v0_
            nw291_[int(8)] = not(not(d_1634_v0_))
            nw291_[int(9)] = d_1634_v0_
            nw291_[int(10)] = d_1634_v0_
            nw291_[int(11)] = d_1634_v0_
            nw291_[int(12)] = d_1634_v0_
            nw291_[int(13)] = d_1634_v0_
            nw291_[int(14)] = not(d_1634_v0_)
            nw291_[int(15)] = d_1634_v0_
            nw291_[int(16)] = d_1634_v0_
            nw291_[int(17)] = True
            nw291_[int(18)] = d_1634_v0_
            nw291_[int(19)] = True
            nw291_[int(20)] = d_1634_v0_
            nw291_[int(21)] = d_1634_v0_
            nw291_[int(22)] = d_1634_v0_
            d_1679_v41_ = nw291_
            d_1680_v42_: _dafny.MultiSet
            d_1680_v42_ = _dafny.MultiSet([d_1679_v41_, d_1679_v41_, d_1679_v41_])
            d_1681_v43_: C4
            nw292_ = C4()
            nw292_.ctor__(((d_1680_v42_) | (d_1680_v42_)).cardinality, (self).f28)
            d_1681_v43_ = nw292_
            (globalState).f0 = default__.safeModuloInt((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wibmqbfun")))) * ((self).f27), (self).f27)
            (globalState).f21 = (self).f27
            if d_1634_v0_:
                d_1682_v44_: C3
                nw293_ = C3()
                def iife150_(_pat_let48_0):
                    def iife151_(d_1683_dt__update__tmp_h0_):
                        def iife152_(_pat_let49_0):
                            def iife153_(d_1684_dt__update_hcf82_h0_):
                                return D17_DC43(d_1684_dt__update_hcf82_h0_)
                            return iife153_(_pat_let49_0)
                        return iife152_((self).f28)
                    return iife151_(_pat_let48_0)
                nw293_.ctor__((self).f27, (iife150_(D17_DC43((self).f28))).cf82)
                d_1682_v44_ = nw293_
                (globalState).f26 = d_1634_v0_
                d_1634_v0_ = d_1634_v0_
                d_1685_v45_: _dafny.Map
                d_1685_v45_ = _dafny.Map({d_1634_v0_: d_1634_v0_})
                d_1686_v46_: _dafny.Seq
                d_1686_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikcash"))
                d_1685_v45_ = (d_1685_v45_).set(d_1634_v0_, (len(d_1686_v46_)) > ((self).f27))
                d_1687_v47_: _dafny.Array
                def lambda71_(d_1688_v0_):
                    def lambda72_(d_1689_i2_):
                        return d_1688_v0_

                    return lambda72_

                init35_ = lambda71_(d_1634_v0_)
                nw294_ = _dafny.Array(None, 18)
                for i0_35_ in range(nw294_.length(0)):
                    nw294_[i0_35_] = init35_(i0_35_)
                d_1687_v47_ = nw294_
                d_1690_v48_: _dafny.Seq
                d_1690_v48_ = _dafny.SeqWithoutIsStrInference([d_1687_v47_])
                r2 = (d_1690_v48_) != (d_1690_v48_)
            elif True:
                index218_ = default__.safeIndex(603, (d_1679_v41_).length(0))
                (d_1679_v41_)[index218_] = d_1634_v0_
                d_1691_v49_: _dafny.Array
                def lambda73_(d_1692_i3_):
                    return default__.safeModuloInt(d_1692_i3_, (self).f27)

                init36_ = lambda73_
                nw295_ = _dafny.Array(None, 3)
                for i0_36_ in range(nw295_.length(0)):
                    nw295_[i0_36_] = init36_(i0_36_)
                d_1691_v49_ = nw295_
                index219_ = default__.safeIndex(629, (d_1691_v49_).length(0))
                (d_1691_v49_)[index219_] = ((self).f27 if default__.fm0((0) - ((self).f27), globalState) else -75)
                d_1693_v50_: D17
                d_1693_v50_ = D17_DC44((self).f27)
                d_1694_v51_: D17
                d_1694_v51_ = D17_DC45(d_1693_v50_)
                d_1695_v52_: _dafny.Map
                d_1695_v52_ = _dafny.Map({d_1634_v0_: d_1634_v0_})
                d_1696_v53_: _dafny.Set
                d_1696_v53_ = _dafny.Set({d_1695_v52_, d_1695_v52_, d_1695_v52_})
                d_1697_v54_: _dafny.Seq
                d_1697_v54_ = _dafny.SeqWithoutIsStrInference([d_1696_v53_])
                d_1698_v55_: D18
                d_1698_v55_ = D18_DC46((d_1697_v54_)[default__.safeIndex((self).f27, len(d_1697_v54_))])
                index220_ = default__.safeIndex(603, (d_1679_v41_).length(0))
                index221_ = default__.safeIndex(629, (d_1691_v49_).length(0))
                rhs309_ = ((0) - ((self).f27)) not in ((self).f30)
                rhs310_ = (self).f27
                rhs311_ = d_1694_v51_
                rhs312_ = default__.safeDivisionInt((self).f27, len(((d_1698_v55_).cf85) | (_dafny.Set({d_1695_v52_, _dafny.Map({d_1634_v0_: False}), d_1695_v52_}))))
                rhs313_ = (self).f27
                lhs239_ = d_1679_v41_
                lhs240_ = default__.safeIndex(603, (d_1679_v41_).length(0))
                lhs241_ = d_1691_v49_
                lhs242_ = default__.safeIndex(629, (d_1691_v49_).length(0))
                lhs243_ = globalState
                lhs244_ = globalState
                lhs239_[lhs240_] = rhs309_
                lhs241_[lhs242_] = rhs310_
                d_1694_v51_ = rhs311_
                lhs243_.f3 = rhs312_
                lhs244_.f15 = rhs313_
                d_1699_v56_: C2
                nw296_ = C2()
                nw296_.ctor__(_dafny.MultiSet([(d_1679_v41_)[default__.safeIndex(603, (d_1679_v41_).length(0))], d_1634_v0_, not((d_1679_v41_)[default__.safeIndex(603, (d_1679_v41_).length(0))])]), 268)
                d_1699_v56_ = nw296_
                d_1700_v57_: _dafny.Array
                nw297_ = _dafny.Array(False, 9)
                d_1700_v57_ = nw297_
                d_1701_v58_: _dafny.Map
                d_1701_v58_ = _dafny.Map({(d_1691_v49_)[default__.safeIndex(629, (d_1691_v49_).length(0))]: d_1700_v57_})
                d_1701_v58_ = d_1701_v58_
                d_1702_v59_: _dafny.MultiSet
                d_1702_v59_ = _dafny.MultiSet([(self).f27])
                d_1703_v60_: _dafny.Set
                d_1703_v60_ = _dafny.Set({d_1702_v59_, d_1702_v59_})
                d_1704_v61_: _dafny.Map
                d_1704_v61_ = _dafny.Map({D17_DC44((self).f27): d_1703_v60_})
                d_1705_v62_: _dafny.Seq
                d_1705_v62_ = _dafny.SeqWithoutIsStrInference([d_1634_v0_])
                d_1706_v63_: _dafny.Map
                d_1706_v63_ = _dafny.Map({d_1705_v62_: len(d_1701_v58_)})
                d_1707_v64_: D17
                d_1707_v64_ = D17_DC44(len(d_1706_v63_))
                def iife154_(_pat_let50_0):
                    def iife155_(d_1708_dt__update__tmp_h1_):
                        def iife156_(_pat_let51_0):
                            def iife157_(d_1709_dt__update_hcf83_h0_):
                                return D17_DC44(d_1709_dt__update_hcf83_h0_)
                            return iife157_(_pat_let51_0)
                        return iife156_(len((self).f30))
                    return iife155_(_pat_let50_0)
                d_1704_v61_ = (d_1704_v61_).set((iife154_(d_1707_v64_) if d_1634_v0_ else d_1707_v64_), d_1703_v60_)
                d_1710_v65_: _dafny.Set
                d_1710_v65_ = _dafny.Set({d_1691_v49_})
                index222_ = default__.safeIndex(603, (d_1679_v41_).length(0))
                rhs314_ = ((d_1710_v65_) - (_dafny.Set({d_1691_v49_}))).issubset(d_1710_v65_)
                rhs315_ = not(d_1634_v0_)
                rhs316_ = d_1691_v49_
                rhs317_ = (d_1679_v41_)[default__.safeIndex(603, (d_1679_v41_).length(0))]
                lhs245_ = d_1679_v41_
                lhs246_ = default__.safeIndex(603, (d_1679_v41_).length(0))
                lhs247_ = globalState
                lhs248_ = globalState
                lhs245_[lhs246_] = rhs314_
                lhs247_.f26 = rhs315_
                d_1691_v49_ = rhs316_
                lhs248_.f26 = rhs317_
            d_1711_v66_: _dafny.Set
            d_1711_v66_ = _dafny.Set({(self).f27})
            d_1712_v67_: D9
            d_1712_v67_ = D9_DC17(d_1711_v66_)
            d_1713_v68_: _dafny.Map
            d_1713_v68_ = _dafny.Map({d_1712_v67_: (self).f27})
            d_1714_v69_: _dafny.Seq
            d_1714_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "heeuww"))
            d_1715_v70_: C0
            nw298_ = C0()
            nw298_.ctor__(len(d_1713_v68_), d_1714_v69_)
            d_1715_v70_ = nw298_
            d_1716_v71_: _dafny.Seq
            d_1716_v71_ = _dafny.SeqWithoutIsStrInference([d_1715_v70_])
            d_1717_v72_: int
            out46_: int
            out46_ = (d_1681_v43_).m11((d_1716_v71_)[default__.safeIndex((d_1715_v70_).f36, len(d_1716_v71_))], d_1634_v0_, globalState)
            d_1717_v72_ = out46_
        d_1718_v73_: str
        d_1718_v73_ = _dafny.CodePoint('g')
        d_1719_v74_: _dafny.Seq
        d_1719_v74_ = _dafny.SeqWithoutIsStrInference([d_1718_v73_])
        d_1720_v75_: _dafny.MultiSet
        d_1720_v75_ = _dafny.MultiSet([(self).f27, (self).f27])
        d_1721_v76_: _dafny.Set
        d_1721_v76_ = _dafny.Set({d_1719_v74_, (d_1719_v74_).set(default__.safeIndex((d_1720_v75_).cardinality, len(d_1719_v74_)), d_1718_v73_), d_1719_v74_, d_1719_v74_})
        d_1722_v77_: _dafny.Map
        d_1722_v77_ = _dafny.Map({d_1718_v73_: d_1634_v0_})
        d_1723_v78_: _dafny.Map
        d_1723_v78_ = _dafny.Map({(self).f27: False})
        d_1724_v79_: _dafny.Seq
        d_1724_v79_ = _dafny.SeqWithoutIsStrInference([((d_1723_v78_)[len((d_1719_v74_).set(default__.safeIndex((self).f27, len(d_1719_v74_)), d_1718_v73_))] if (len((d_1719_v74_).set(default__.safeIndex((self).f27, len(d_1719_v74_)), d_1718_v73_))) in (d_1723_v78_) else d_1634_v0_)])
        d_1725_v80_: D15
        d_1725_v80_ = D15_DC36((d_1718_v73_ if d_1634_v0_ else d_1718_v73_), (self).fm4(d_1721_v76_, ((d_1722_v77_)[d_1718_v73_] if (d_1718_v73_) in (d_1722_v77_) else d_1634_v0_), d_1724_v79_, globalState))
        d_1725_v80_ = d_1725_v80_
        d_1719_v74_ = (default__.fm31(d_1634_v0_, d_1634_v0_, globalState)) + (d_1719_v74_)
        d_1726_i4_: int
        d_1726_i4_ = 0
        with _dafny.label("7"):
            while d_1634_v0_:
                with _dafny.c_label("7"):
                    if (d_1726_i4_) >= (100):
                        raise _dafny.Break("7")
                    d_1726_i4_ = (d_1726_i4_) + (1)
                    d_1727_v81_: C1
                    nw299_ = C1()
                    nw299_.ctor__(((self).f27) >= (len((self).f30)), (self).f29, (self).f30, (self).f27, (self).f28)
                    d_1727_v81_ = nw299_
                    d_1728_v82_: _dafny.Seq
                    d_1728_v82_ = _dafny.SeqWithoutIsStrInference([d_1727_v81_])
                    d_1727_v81_ = ((d_1728_v82_)[default__.safeIndex((0) - (len((self).f30)), len(d_1728_v82_))] if d_1634_v0_ else d_1727_v81_)
                    d_1719_v74_ = d_1719_v74_
                    (globalState).f21 = (self).f27
                    d_1729_v83_: _dafny.MultiSet
                    d_1729_v83_ = _dafny.MultiSet([d_1718_v73_])
                    d_1730_v84_: _dafny.Map
                    d_1730_v84_ = _dafny.Map({(d_1634_v0_) and (d_1727_v81_.f35): d_1729_v83_})
                    d_1730_v84_ = (d_1730_v84_).set((d_1724_v79_) < (d_1724_v79_), d_1729_v83_)
                    pass
            pass
        d_1731_v85_: C3
        nw300_ = C3()
        nw300_.ctor__((self).f27, (self).f28)
        d_1731_v85_ = nw300_
        if d_1634_v0_:
            d_1634_v0_ = default__.fm0(860, globalState)
            index223_ = default__.safeIndex(243, ((self).f28).length(0))
            ((self).f28)[index223_] = d_1719_v74_
            d_1732_v86_: _dafny.Array
            def lambda74_(d_1733_v0_):
                def lambda75_(d_1734_i5_):
                    return d_1733_v0_

                return lambda75_

            init37_ = lambda74_(d_1634_v0_)
            nw301_ = _dafny.Array(None, 28)
            for i0_37_ in range(nw301_.length(0)):
                nw301_[i0_37_] = init37_(i0_37_)
            d_1732_v86_ = nw301_
            d_1735_v87_: _dafny.Set
            d_1735_v87_ = _dafny.Set({(self).f27})
            d_1736_v88_: D4
            d_1736_v88_ = D4_DC8(d_1718_v73_, len(d_1735_v87_), d_1634_v0_)
            index224_ = default__.safeIndex(304, (d_1732_v86_).length(0))
            (d_1732_v86_)[index224_] = (d_1724_v79_)[default__.safeIndex((d_1736_v88_).cf14, len(d_1724_v79_))]
            d_1737_v89_: D6
            d_1737_v89_ = D6_DC12(default__.fm0((self).f27, globalState), default__.fm44(globalState), (default__.fm31(d_1634_v0_, d_1634_v0_, globalState)).set(default__.safeIndex((self).f27, len(default__.fm31(d_1634_v0_, d_1634_v0_, globalState))), default__.fm38(_dafny.Map({d_1634_v0_: d_1634_v0_}), globalState)))
            index225_ = default__.safeIndex(243, ((self).f28).length(0))
            index226_ = default__.safeIndex(304, (d_1732_v86_).length(0))
            rhs318_ = (d_1737_v89_).cf22
            rhs319_ = (d_1634_v0_) == (False)
            lhs249_ = (self).f28
            lhs250_ = default__.safeIndex(243, ((self).f28).length(0))
            lhs251_ = d_1732_v86_
            lhs252_ = default__.safeIndex(304, (d_1732_v86_).length(0))
            lhs249_[lhs250_] = rhs318_
            lhs251_[lhs252_] = rhs319_
            d_1738_v90_: _dafny.Array
            def lambda76_(d_1739_v73_, d_1740_v0_):
                def lambda77_(d_1741_i6_):
                    return _dafny.Map({_dafny.SeqWithoutIsStrInference([D1_DC2(d_1739_v73_, (self).f27) for d_1742_i7_ in range(default__.abs(882))]): d_1740_v0_})

                return lambda77_

            init38_ = lambda76_(d_1718_v73_, d_1634_v0_)
            nw302_ = _dafny.Array(None, 1)
            for i0_38_ in range(nw302_.length(0)):
                nw302_[i0_38_] = init38_(i0_38_)
            d_1738_v90_ = nw302_
            d_1743_v91_: _dafny.Map
            d_1743_v91_ = _dafny.Map({(self).f27: -418})
            d_1744_v92_: _dafny.Map
            d_1744_v92_ = _dafny.Map({(d_1732_v86_)[default__.safeIndex(304, (d_1732_v86_).length(0))]: _dafny.Set({d_1634_v0_})})
            d_1745_v93_: D1
            d_1745_v93_ = D1_DC2(d_1718_v73_, ((d_1743_v91_)[len(d_1724_v79_)] if (len(d_1724_v79_)) in (d_1743_v91_) else len(d_1744_v92_)))
            d_1746_v94_: _dafny.Seq
            d_1746_v94_ = _dafny.SeqWithoutIsStrInference([d_1745_v93_])
            d_1747_v95_: _dafny.Map
            d_1747_v95_ = _dafny.Map({d_1746_v94_: (d_1732_v86_)[default__.safeIndex(304, (d_1732_v86_).length(0))]})
            index227_ = default__.safeIndex(773, (d_1738_v90_).length(0))
            (d_1738_v90_)[index227_] = (_dafny.Map({d_1746_v94_: (d_1731_v85_).fm46((self).f27, d_1718_v73_, globalState)})) | (d_1747_v95_)
            index228_ = default__.safeIndex(773, (d_1738_v90_).length(0))
            (d_1738_v90_)[index228_] = d_1747_v95_
            index229_ = default__.safeIndex(304, (d_1732_v86_).length(0))
            (d_1732_v86_)[index229_] = (d_1732_v86_)[default__.safeIndex(304, (d_1732_v86_).length(0))]
            d_1748_v96_: _dafny.Seq
            d_1748_v96_ = _dafny.SeqWithoutIsStrInference([(self).f30])
            d_1749_v97_: _dafny.Map
            d_1749_v97_ = _dafny.Map({d_1634_v0_: d_1718_v73_})
            d_1750_v98_: _dafny.Map
            d_1750_v98_ = _dafny.Map({(len(d_1748_v96_)) != (len(d_1749_v97_)): (self).f27})
            (globalState).f21 = ((d_1750_v98_)[(d_1732_v86_)[default__.safeIndex(304, (d_1732_v86_).length(0))]] if ((d_1732_v86_)[default__.safeIndex(304, (d_1732_v86_).length(0))]) in (d_1750_v98_) else (self).f27)
        elif True:
            d_1634_v0_ = d_1634_v0_
            d_1751_v99_: _dafny.Set
            d_1751_v99_ = _dafny.Set({d_1718_v73_, d_1718_v73_})
            d_1752_v100_: D19
            d_1752_v100_ = D19_DC48(d_1751_v99_)
            def iife158_():
                coll54_ = _dafny.Set()
                compr_54_: _dafny.Seq
                for compr_54_ in (_dafny.Map({default__.fm37((self).f27, len((d_1752_v100_).cf88), globalState): d_1634_v0_})).keys.Elements:
                    d_1753_v101_: _dafny.Seq = compr_54_
                    if (d_1753_v101_) in (_dafny.Map({default__.fm37((self).f27, len((d_1752_v100_).cf88), globalState): d_1634_v0_})):
                        coll54_ = coll54_.union(_dafny.Set([d_1753_v101_]))
                return _dafny.Set(coll54_)
            (globalState).f3 = (363) - (len((iife158_()
            ) | (d_1721_v76_)))
            d_1754_v102_: _dafny.Map
            d_1754_v102_ = _dafny.Map({d_1634_v0_: d_1634_v0_})
            d_1754_v102_ = (d_1754_v102_).set(((self).f27) == (len(d_1719_v74_)), d_1634_v0_)
            d_1755_v103_: _dafny.Map
            d_1755_v103_ = _dafny.Map({(self).f27: (self).f27})
            d_1756_v104_: T0
            nw303_ = C3()
            nw303_.ctor__((self).f27, (self).f28)
            d_1756_v104_ = nw303_
            d_1757_v105_: _dafny.Array
            nw304_ = _dafny.Array(None, 29)
            nw304_[int(0)] = (self).f27
            nw304_[int(1)] = len(_dafny.Map({((d_1755_v103_)[(self).f27] if ((self).f27) in (d_1755_v103_) else len(default__.fm48(d_1634_v0_, d_1634_v0_, len(_dafny.Set({d_1756_v104_})), globalState))): (self).f27}))
            nw304_[int(2)] = 477
            nw304_[int(3)] = (d_1756_v104_).f27
            nw304_[int(4)] = (self).f27
            nw304_[int(5)] = (d_1756_v104_).f27
            nw304_[int(6)] = len(d_1719_v74_)
            nw304_[int(7)] = (self).f27
            nw304_[int(8)] = (self).f27
            nw304_[int(9)] = (self).f27
            nw304_[int(10)] = (d_1756_v104_).f27
            nw304_[int(11)] = (0) - ((self).f27)
            nw304_[int(12)] = len(d_1723_v78_)
            nw304_[int(13)] = (self).f27
            nw304_[int(14)] = (self).f27
            nw304_[int(15)] = (self).f27
            nw304_[int(16)] = len(d_1719_v74_)
            nw304_[int(17)] = (d_1756_v104_).f27
            nw304_[int(18)] = (self).f27
            nw304_[int(19)] = (d_1756_v104_).f27
            nw304_[int(20)] = 386
            nw304_[int(21)] = (d_1756_v104_).f27
            nw304_[int(22)] = (d_1756_v104_).f27
            nw304_[int(23)] = (self).f27
            nw304_[int(24)] = (self).f27
            nw304_[int(25)] = (self).fm4(d_1721_v76_, False, d_1724_v79_, globalState)
            nw304_[int(26)] = (d_1756_v104_).f27
            nw304_[int(27)] = (d_1756_v104_).f27
            nw304_[int(28)] = 235
            d_1757_v105_ = nw304_
            d_1758_v106_: _dafny.Set
            d_1758_v106_ = _dafny.Set({d_1757_v105_, d_1757_v105_, d_1757_v105_, d_1757_v105_})
            d_1759_v107_: _dafny.Map
            d_1759_v107_ = _dafny.Map({not((self).fm3((self).f29, globalState)): d_1755_v103_})
            d_1760_v108_: _dafny.Array
            nw305_ = _dafny.Array(None, 17)
            nw305_[int(0)] = d_1755_v103_
            nw305_[int(1)] = d_1755_v103_
            nw305_[int(2)] = _dafny.Map({(self).f27: (self).f27})
            nw305_[int(3)] = (d_1755_v103_) | (d_1755_v103_)
            nw305_[int(4)] = d_1755_v103_
            nw305_[int(5)] = (d_1755_v103_).set((self).f27, (self).f27)
            nw305_[int(6)] = _dafny.Map({(self).f27: (self).f27})
            nw305_[int(7)] = d_1755_v103_
            nw305_[int(8)] = d_1755_v103_
            nw305_[int(9)] = _dafny.Map({-591: len(d_1758_v106_)})
            nw305_[int(10)] = (_dafny.Map({(self).f27: (d_1756_v104_).f27})).set(len(d_1724_v79_), (self).f27)
            nw305_[int(11)] = d_1755_v103_
            nw305_[int(12)] = (_dafny.Map({(d_1756_v104_).f27: (self).f27}) if d_1634_v0_ else d_1755_v103_)
            nw305_[int(13)] = d_1755_v103_
            nw305_[int(14)] = ((d_1759_v107_)[((d_1754_v102_)[(d_1724_v79_)[default__.safeIndex((d_1756_v104_).f27, len(d_1724_v79_))]] if ((d_1724_v79_)[default__.safeIndex((d_1756_v104_).f27, len(d_1724_v79_))]) in (d_1754_v102_) else d_1634_v0_)] if (((d_1754_v102_)[(d_1724_v79_)[default__.safeIndex((d_1756_v104_).f27, len(d_1724_v79_))]] if ((d_1724_v79_)[default__.safeIndex((d_1756_v104_).f27, len(d_1724_v79_))]) in (d_1754_v102_) else d_1634_v0_)) in (d_1759_v107_) else d_1755_v103_)
            nw305_[int(15)] = d_1755_v103_
            nw305_[int(16)] = _dafny.Map({230: len(d_1719_v74_)})
            d_1760_v108_ = nw305_
            index230_ = default__.safeIndex(720, (d_1760_v108_).length(0))
            (d_1760_v108_)[index230_] = (d_1755_v103_) | (d_1755_v103_)
            d_1761_v110_: _dafny.Map
            d_1761_v110_ = _dafny.Map({(d_1756_v104_).f27: d_1718_v73_})
            d_1762_v111_: _dafny.Map
            d_1762_v111_ = _dafny.Map({(d_1756_v104_).f27: d_1761_v110_})
            d_1763_v112_: _dafny.Seq
            def iife159_():
                coll55_ = _dafny.Map()
                compr_55_: int
                for compr_55_ in _dafny.IntegerRange(124, 382):
                    d_1764_v109_: int = compr_55_
                    if ((124) <= (d_1764_v109_)) and ((d_1764_v109_) < (382)):
                        coll55_[default__.safeModuloInt(d_1764_v109_, (d_1756_v104_).f27)] = _dafny.CodePoint('t')
                return _dafny.Map(coll55_)
            d_1763_v112_ = _dafny.SeqWithoutIsStrInference([iife159_()
            , (d_1761_v110_).set(len(_dafny.SeqWithoutIsStrInference([not(d_1634_v0_)])), d_1718_v73_), d_1761_v110_, ((d_1762_v111_)[(self).f27] if ((self).f27) in (d_1762_v111_) else d_1761_v110_)])
            d_1765_v113_: D15
            d_1765_v113_ = D15_DC35(d_1763_v112_)
            d_1766_v114_: _dafny.Seq
            d_1766_v114_ = _dafny.SeqWithoutIsStrInference([D15_DC35(d_1763_v112_), d_1765_v113_])
            d_1767_v115_: _dafny.Seq
            d_1767_v115_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([D15_DC35(_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1756_v104_).f27: d_1718_v73_})])) for d_1768_i8_ in range(default__.abs(362))])).set(default__.safeIndex(-918, len(_dafny.SeqWithoutIsStrInference([D15_DC35(_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1756_v104_).f27: d_1718_v73_})])) for d_1769_i8_ in range(default__.abs(362))]))), d_1765_v113_), (d_1766_v114_) + (d_1766_v114_), (d_1766_v114_) + (d_1766_v114_)])
            d_1770_v116_: C0
            nw306_ = C0()
            nw306_.ctor__((self).f27, d_1719_v74_)
            d_1770_v116_ = nw306_
            index231_ = default__.safeIndex(720, (d_1760_v108_).length(0))
            rhs320_ = ((d_1755_v103_) | (d_1755_v103_)) | (d_1755_v103_)
            rhs321_ = d_1634_v0_
            rhs322_ = d_1767_v115_
            rhs323_ = d_1770_v116_
            lhs253_ = d_1760_v108_
            lhs254_ = default__.safeIndex(720, (d_1760_v108_).length(0))
            lhs253_[lhs254_] = rhs320_
            r2 = rhs321_
            d_1767_v115_ = rhs322_
            d_1770_v116_ = rhs323_
            d_1771_v117_: _dafny.Seq
            d_1771_v117_ = _dafny.SeqWithoutIsStrInference([d_1757_v105_])
            d_1771_v117_ = (d_1771_v117_) + (d_1771_v117_)
        d_1772_v118_: D9
        d_1772_v118_ = D9_DC19(not(d_1634_v0_))
        d_1773_v119_: _dafny.Set
        d_1773_v119_ = _dafny.Set({True, True, d_1634_v0_, d_1634_v0_, d_1634_v0_})
        nw307_ = _dafny.Array(None, 3)
        nw307_[int(0)] = not ((d_1772_v118_).cf33) or (default__.fm0((self).f27, globalState))
        nw307_[int(1)] = d_1634_v0_
        nw307_[int(2)] = not((d_1773_v119_).ispropersubset(d_1773_v119_))
        r0 = nw307_
        d_1774_v120_: D7
        d_1774_v120_ = D7_DC13(d_1724_v79_)
        pat_let_tv37_ = d_1634_v0_
        pat_let_tv38_ = d_1634_v0_
        def lambda78_(source26_):
            if source26_.is_DC14:
                d_1775___mcc_h0_ = source26_.cf24
                d_1776___mcc_h1_ = source26_.cf25
                d_1777___mcc_h2_ = source26_.cf26
                d_1778___mcc_h3_ = source26_.cf27
                d_1779_cf27_ = d_1778___mcc_h3_
                d_1780_cf26_ = d_1777___mcc_h2_
                d_1781_cf25_ = d_1776___mcc_h1_
                d_1782_cf24_ = d_1775___mcc_h0_
                return pat_let_tv37_
            elif True:
                d_1783___mcc_h4_ = source26_.cf23
                d_1784_cf23_ = d_1783___mcc_h4_
                return pat_let_tv38_

        r1 = lambda78_(d_1774_v120_)
        r2 = (_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])) <= ((((self).f30) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({((self).f30)[default__.safeIndex(58, len((self).f30))]: d_1634_v0_}))]))).set(default__.safeIndex(((_dafny.MultiSet((self).f30)).set((self).f27, default__.abs((self).f27))).cardinality, len(((self).f30) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({((self).f30)[default__.safeIndex(58, len((self).f30))]: d_1634_v0_}))])))), 669))
        return r0, r1, r2

    def m10(self, p0, p1, globalState):
        d_1785_v0_: str
        d_1785_v0_ = _dafny.CodePoint('i')
        d_1786_v1_: _dafny.Seq
        d_1786_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xggvd"))
        d_1787_v2_: D2
        d_1787_v2_ = D2_DC4((self).f27, 521, (0) - ((self).f27), d_1785_v0_, d_1786_v1_)
        d_1788_v3_: _dafny.Set
        d_1788_v3_ = _dafny.Set({d_1787_v2_})
        hi8_ = (len(d_1788_v3_) if p0 else (self).f27)
        for d_1789_i0_ in range(((self).f27 if p0 else (self).f27), hi8_):
            d_1790_v4_: _dafny.Map
            d_1790_v4_ = _dafny.Map({(self).f27: p0})
            d_1791_v5_: C2
            nw308_ = C2()
            nw308_.ctor__(_dafny.MultiSet([p0]), len((d_1790_v4_) | (d_1790_v4_)))
            d_1791_v5_ = nw308_
            (globalState).f26 = p0
            d_1792_v6_: _dafny.Seq
            d_1792_v6_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1792_v6_ = d_1792_v6_
            d_1793_v7_: _dafny.Array
            nw309_ = _dafny.Array(_dafny.CodePoint('D'), 21)
            d_1793_v7_ = nw309_
            d_1794_v8_: _dafny.Map
            d_1794_v8_ = _dafny.Map({d_1792_v6_: d_1793_v7_})
            d_1795_v9_: _dafny.Map
            d_1795_v9_ = _dafny.Map({d_1792_v6_: ((d_1794_v8_)[d_1792_v6_] if (d_1792_v6_) in (d_1794_v8_) else d_1793_v7_)})
            d_1795_v9_ = (d_1795_v9_).set(_dafny.SeqWithoutIsStrInference([not(p0), p0]), d_1793_v7_)
        d_1796_v10_: _dafny.Array
        nw310_ = _dafny.Array(_dafny.Map({}), 4)
        d_1796_v10_ = nw310_
        d_1797_v11_: D16
        d_1797_v11_ = D16_DC39(d_1796_v10_)
        pat_let_tv39_ = d_1796_v10_
        pat_let_tv40_ = d_1796_v10_
        def iife161_(_pat_let53_0):
            def iife162_(d_1798_dt__update__tmp_h0_):
                def iife163_(_pat_let54_0):
                    def iife164_(d_1799_dt__update_hcf73_h0_):
                        return D16_DC39(d_1799_dt__update_hcf73_h0_)
                    return iife164_(_pat_let54_0)
                return iife163_(pat_let_tv39_)
            return iife162_(_pat_let53_0)
        def iife160_(_pat_let52_0):
            def iife165_(d_1800_dt__update__tmp_h1_):
                def iife166_(_pat_let55_0):
                    def iife167_(d_1801_dt__update_hcf73_h1_):
                        return D16_DC39(d_1801_dt__update_hcf73_h1_)
                    return iife167_(_pat_let55_0)
                return iife166_(pat_let_tv40_)
            return iife165_(_pat_let52_0)
        d_1797_v11_ = iife160_(iife161_(d_1797_v11_))
        d_1802_v12_: _dafny.Seq
        d_1802_v12_ = _dafny.SeqWithoutIsStrInference([((self).f27) <= (386)])
        rhs324_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        rhs325_ = default__.fm1(globalState)
        rhs326_ = (self).f27
        lhs255_ = globalState
        lhs256_ = globalState
        d_1802_v12_ = rhs324_
        lhs255_.f24 = rhs325_
        lhs256_.f0 = rhs326_
        d_1803_v13_: _dafny.Map
        d_1803_v13_ = _dafny.Map({425: (0) - ((self).f27)})
        d_1804_v14_: _dafny.Map
        d_1804_v14_ = _dafny.Map({((d_1803_v13_)[((self).f30)[default__.safeIndex((self).f27, len((self).f30))]] if (((self).f30)[default__.safeIndex((self).f27, len((self).f30))]) in (d_1803_v13_) else (0) - ((self).f27)): False})
        d_1804_v14_ = (d_1804_v14_).set(len(default__.fm48(p0, p0, -276, globalState)), p0)
        d_1785_v0_ = d_1785_v0_
        d_1805_v15_: _dafny.Map
        d_1805_v15_ = _dafny.Map({not(default__.fm0((self).f27, globalState)): True})
        d_1806_v16_: D8
        d_1806_v16_ = D8_DC15((d_1805_v15_) | (d_1805_v15_))
        source27_ = d_1806_v16_
        if source27_.is_DC16:
            d_1807___mcc_h0_ = source27_.cf29
            d_1808___mcc_h1_ = source27_.cf30
            d_1809_cf30_ = d_1808___mcc_h1_
            d_1810_cf29_ = d_1807___mcc_h0_
            d_1811_v17_: _dafny.Map
            d_1811_v17_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jiyhtfc"))})
            d_1812_v18_: _dafny.Array
            nw311_ = _dafny.Array(None, 20)
            nw311_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnsowybbf"))
            nw311_[int(1)] = d_1786_v1_
            nw311_[int(2)] = d_1786_v1_
            nw311_[int(3)] = default__.fm31(p0, False, globalState)
            nw311_[int(4)] = (d_1786_v1_) + (d_1786_v1_)
            nw311_[int(5)] = (default__.fm37((self).f27, 422, globalState)) + (d_1786_v1_)
            nw311_[int(6)] = d_1786_v1_
            nw311_[int(7)] = d_1786_v1_
            nw311_[int(8)] = d_1786_v1_
            nw311_[int(9)] = d_1786_v1_
            nw311_[int(10)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yo"))) + (d_1786_v1_)).set(default__.safeIndex(d_1809_cf30_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yo"))) + (d_1786_v1_))), d_1785_v0_)
            nw311_[int(11)] = d_1786_v1_
            nw311_[int(12)] = (d_1786_v1_) + (_dafny.SeqWithoutIsStrInference([d_1785_v0_ for d_1813_i1_ in range(default__.abs(-646))]))
            nw311_[int(13)] = (_dafny.SeqWithoutIsStrInference([d_1785_v0_ for d_1814_i2_ in range(default__.abs(-14))])) + (d_1786_v1_)
            nw311_[int(14)] = d_1786_v1_
            nw311_[int(15)] = d_1786_v1_
            nw311_[int(16)] = ((d_1811_v17_)[True] if (True) in (d_1811_v17_) else d_1786_v1_)
            nw311_[int(17)] = d_1786_v1_
            nw311_[int(18)] = d_1786_v1_
            nw311_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggdic"))
            d_1812_v18_ = nw311_
            d_1815_v19_: _dafny.Map
            d_1815_v19_ = _dafny.Map({(len(d_1786_v1_)) - (len(d_1786_v1_)): (self).f28})
            d_1812_v18_ = ((d_1815_v19_)[d_1810_cf29_] if (d_1810_cf29_) in (d_1815_v19_) else (self).f28)
            d_1816_v20_: _dafny.Array
            nw312_ = _dafny.Array(D12.default()(), 16)
            d_1816_v20_ = nw312_
            d_1817_v21_: _dafny.Map
            d_1817_v21_ = _dafny.Map({d_1816_v20_: d_1785_v0_})
            d_1818_v22_: _dafny.Map
            d_1818_v22_ = _dafny.Map({d_1810_cf29_: d_1785_v0_})
            rhs327_ = ((self).f27) * ((p1)[default__.safeIndex(d_1810_cf29_, len(p1))])
            rhs328_ = p0
            rhs329_ = _dafny.Map({d_1816_v20_: ((d_1818_v22_)[(self).f27] if ((self).f27) in (d_1818_v22_) else d_1785_v0_)})
            rhs330_ = -672
            lhs257_ = globalState
            lhs258_ = globalState
            lhs259_ = globalState
            lhs257_.f0 = rhs327_
            lhs258_.f26 = rhs328_
            d_1817_v21_ = rhs329_
            lhs259_.f18 = rhs330_
            d_1819_v23_: _dafny.Seq
            d_1819_v23_ = _dafny.SeqWithoutIsStrInference([d_1786_v1_])
            d_1820_v24_: T0
            nw313_ = C4()
            nw313_.ctor__((len(d_1786_v1_)) - (d_1810_cf29_), (self).f28)
            d_1820_v24_ = nw313_
            d_1821_v25_: _dafny.Array
            def lambda79_(d_1822_cf29_):
                def lambda80_(d_1823_i3_):
                    return default__.safeDivisionInt(d_1823_i3_, d_1822_cf29_)

                return lambda80_

            init39_ = lambda79_(d_1810_cf29_)
            nw314_ = _dafny.Array(None, 28)
            for i0_39_ in range(nw314_.length(0)):
                nw314_[i0_39_] = init39_(i0_39_)
            d_1821_v25_ = nw314_
            d_1824_v26_: _dafny.MultiSet
            d_1824_v26_ = _dafny.MultiSet([(737) < (len((d_1786_v1_).set(default__.safeIndex((self).f27, len(d_1786_v1_)), d_1785_v0_))), False, p0, (p0 if p0 else p0)])
            d_1825_v27_: _dafny.Map
            d_1825_v27_ = _dafny.Map({False: (self).f27})
            rhs331_ = d_1819_v23_
            rhs332_ = d_1820_v24_
            rhs333_ = (d_1821_v25_ if p0 else d_1821_v25_)
            rhs334_ = (d_1810_cf29_) - ((len(d_1825_v27_)) + ((self).f27))
            rhs335_ = _dafny.MultiSet(((default__.fm41(d_1810_cf29_, len(d_1804_v14_), (self).fm3(_dafny.MultiSet([p0]), globalState), globalState)) + (d_1802_v12_)).set(default__.safeIndex(d_1810_cf29_, len((default__.fm41(d_1810_cf29_, len(d_1804_v14_), (self).fm3(_dafny.MultiSet([p0]), globalState), globalState)) + (d_1802_v12_))), p0))
            lhs260_ = globalState
            d_1819_v23_ = rhs331_
            d_1820_v24_ = rhs332_
            d_1821_v25_ = rhs333_
            lhs260_.f0 = rhs334_
            d_1824_v26_ = rhs335_
            d_1826_v28_: C0
            nw315_ = C0()
            nw315_.ctor__(len(d_1786_v1_), d_1786_v1_)
            d_1826_v28_ = nw315_
            d_1827_v29_: D6
            d_1827_v29_ = D6_DC11(d_1826_v28_)
            source28_ = d_1827_v29_
            if source28_.is_DC12:
                d_1828___mcc_h3_ = source28_.cf20
                d_1829___mcc_h4_ = source28_.cf21
                d_1830___mcc_h5_ = source28_.cf22
                d_1831_cf22_ = d_1830___mcc_h5_
                d_1832_cf21_ = d_1829___mcc_h4_
                d_1833_cf20_ = d_1828___mcc_h3_
                d_1834_v30_: C0
                nw316_ = C0()
                nw316_.ctor__((d_1820_v24_).f27, (d_1826_v28_).f37)
                d_1834_v30_ = nw316_
                d_1810_cf29_ = (d_1826_v28_).f36
                d_1835_v31_: D4
                d_1835_v31_ = D4_DC8(d_1785_v0_, d_1809_cf30_, False)
                d_1836_v32_: _dafny.Set
                d_1836_v32_ = _dafny.Set({(d_1834_v30_).f36, (d_1834_v30_).f36, d_1810_cf29_, d_1810_cf29_, (d_1820_v24_).f27})
                d_1837_v33_: _dafny.Array
                nw317_ = _dafny.Array(None, 15)
                nw317_[int(0)] = p0
                nw317_[int(1)] = p0
                nw317_[int(2)] = (self).fm3((self).f29, globalState)
                nw317_[int(3)] = (d_1802_v12_) <= (d_1802_v12_)
                nw317_[int(4)] = d_1833_cf20_
                nw317_[int(5)] = p0
                nw317_[int(6)] = p0
                nw317_[int(7)] = (d_1809_cf30_) in (d_1836_v32_)
                nw317_[int(8)] = d_1833_cf20_
                nw317_[int(9)] = (default__.fm49((self).f27, globalState)).cf17
                nw317_[int(10)] = d_1833_cf20_
                nw317_[int(11)] = (d_1810_cf29_) >= ((d_1820_v24_).f27)
                nw317_[int(12)] = p0
                nw317_[int(13)] = False
                nw317_[int(14)] = d_1833_cf20_
                d_1837_v33_ = nw317_
                index232_ = default__.safeIndex(594, (d_1837_v33_).length(0))
                (d_1837_v33_)[index232_] = False
                pat_let_tv41_ = d_1834_v30_
                pat_let_tv42_ = d_1810_cf29_
                index233_ = default__.safeIndex(594, (d_1837_v33_).length(0))
                def iife168_(_pat_let56_0):
                    def iife169_(d_1838_dt__update__tmp_h2_):
                        def iife170_(_pat_let57_0):
                            def iife171_(d_1839_dt__update_hcf15_h0_):
                                return D4_DC8((d_1838_dt__update__tmp_h2_).cf13, (d_1838_dt__update__tmp_h2_).cf14, d_1839_dt__update_hcf15_h0_)
                            return iife171_(_pat_let57_0)
                        return iife170_(((pat_let_tv41_).f36) == (pat_let_tv42_))
                    return iife169_(_pat_let56_0)
                rhs336_ = d_1809_cf30_
                rhs337_ = iife168_(D4_DC8(d_1785_v0_, (d_1826_v28_).f36, d_1833_cf20_))
                rhs338_ = not(((d_1826_v28_).f36) > (len(d_1802_v12_)))
                lhs261_ = globalState
                lhs262_ = d_1837_v33_
                lhs263_ = default__.safeIndex(594, (d_1837_v33_).length(0))
                lhs261_.f0 = rhs336_
                d_1835_v31_ = rhs337_
                lhs262_[lhs263_] = rhs338_
                (globalState).f24 = (d_1834_v30_).f36
            elif True:
                d_1840___mcc_h6_ = source28_.cf19
                d_1841_cf19_ = d_1840___mcc_h6_
                rhs339_ = d_1821_v25_
                rhs340_ = ((d_1804_v14_)[(d_1820_v24_).f27] if ((d_1820_v24_).f27) in (d_1804_v14_) else p0)
                rhs341_ = ((d_1820_v24_).f27) * (d_1810_cf29_)
                rhs342_ = (not(p0) if p0 else False)
                rhs343_ = d_1803_v13_
                lhs264_ = globalState
                lhs265_ = globalState
                lhs266_ = globalState
                d_1821_v25_ = rhs339_
                lhs264_.f26 = rhs340_
                lhs265_.f24 = rhs341_
                lhs266_.f26 = rhs342_
                d_1803_v13_ = rhs343_
                d_1805_v15_ = (d_1805_v15_).set(p0, p0)
                d_1842_v34_: C1
                nw318_ = C1()
                nw318_.ctor__(p0, (d_1824_v26_ if p0 else (self).f29), (self).f30, ((d_1820_v24_).f27) * (d_1810_cf29_), (d_1820_v24_).f28)
                d_1842_v34_ = nw318_
                (globalState).f0 = default__.safeDivisionInt(len(d_1786_v1_), (d_1826_v28_).f36)
        elif True:
            d_1843___mcc_h2_ = source27_.cf28
            d_1844_cf28_ = d_1843___mcc_h2_
            d_1845_v35_: _dafny.Array
            nw319_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
            d_1845_v35_ = nw319_
            d_1846_v36_: C0
            nw320_ = C0()
            nw320_.ctor__((self).f27, d_1786_v1_)
            d_1846_v36_ = nw320_
            d_1847_v37_: _dafny.Array
            nw321_ = _dafny.Array(False, 20)
            d_1847_v37_ = nw321_
            index234_ = default__.safeIndex(391, (d_1847_v37_).length(0))
            (d_1847_v37_)[index234_] = p0
            index235_ = default__.safeIndex(391, (d_1847_v37_).length(0))
            rhs344_ = d_1845_v35_
            rhs345_ = d_1846_v36_
            rhs346_ = not(p0)
            lhs267_ = d_1847_v37_
            lhs268_ = default__.safeIndex(391, (d_1847_v37_).length(0))
            d_1845_v35_ = rhs344_
            d_1846_v36_ = rhs345_
            lhs267_[lhs268_] = rhs346_
            d_1848_v38_: _dafny.MultiSet
            d_1848_v38_ = _dafny.MultiSet([(d_1847_v37_)[default__.safeIndex(391, (d_1847_v37_).length(0))]])
            d_1849_v39_: _dafny.Array
            nw322_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_1849_v39_ = nw322_
            d_1850_v40_: C3
            nw323_ = C3()
            nw323_.ctor__((d_1846_v36_).f36, (self).f28)
            d_1850_v40_ = nw323_
            d_1851_v41_: _dafny.Seq
            d_1851_v41_ = _dafny.SeqWithoutIsStrInference([d_1850_v40_, d_1850_v40_])
            d_1852_v42_: _dafny.Array
            nw324_ = _dafny.Array(None, 26)
            nw324_[int(0)] = (d_1846_v36_).f36
            nw324_[int(1)] = (d_1846_v36_).f36
            nw324_[int(2)] = (self).f27
            nw324_[int(3)] = (self).f27
            nw324_[int(4)] = len(d_1851_v41_)
            nw324_[int(5)] = len(d_1805_v15_)
            nw324_[int(6)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtwh"))).set(default__.safeIndex((d_1846_v36_).f36, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtwh")))), d_1785_v0_))
            nw324_[int(7)] = (self).f27
            nw324_[int(8)] = (d_1846_v36_).f36
            nw324_[int(9)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrfkfxp")))
            nw324_[int(10)] = default__.fm1(globalState)
            nw324_[int(11)] = (self).f27
            nw324_[int(12)] = (d_1846_v36_).f36
            nw324_[int(13)] = ((self).f30)[default__.safeIndex(896, len((self).f30))]
            nw324_[int(14)] = (d_1846_v36_).f36
            nw324_[int(15)] = 943
            nw324_[int(16)] = (d_1846_v36_).f36
            nw324_[int(17)] = (self).f27
            nw324_[int(18)] = 719
            nw324_[int(19)] = (self).f27
            nw324_[int(20)] = 306
            nw324_[int(21)] = (self).f27
            nw324_[int(22)] = 947
            nw324_[int(23)] = -688
            nw324_[int(24)] = (self).f27
            nw324_[int(25)] = (self).f27
            d_1852_v42_ = nw324_
            index236_ = default__.safeIndex(29, (d_1849_v39_).length(0))
            (d_1849_v39_)[index236_] = d_1852_v42_
            d_1853_v43_: _dafny.MultiSet
            d_1853_v43_ = _dafny.MultiSet([d_1785_v0_])
            d_1854_v44_: C2
            nw325_ = C2()
            nw325_.ctor__(d_1848_v38_, (self).f27)
            d_1854_v44_ = nw325_
            d_1855_v45_: _dafny.Map
            d_1855_v45_ = _dafny.Map({d_1854_v44_: (d_1847_v37_)[default__.safeIndex(391, (d_1847_v37_).length(0))]})
            index237_ = default__.safeIndex(29, (d_1849_v39_).length(0))
            rhs347_ = (((d_1853_v43_)[d_1785_v0_] if (d_1785_v0_) in (d_1853_v43_) else len(d_1855_v45_))) - ((d_1846_v36_).f36)
            rhs348_ = d_1848_v38_
            rhs349_ = d_1852_v42_
            rhs350_ = p0
            lhs269_ = globalState
            lhs270_ = d_1849_v39_
            lhs271_ = default__.safeIndex(29, (d_1849_v39_).length(0))
            lhs272_ = globalState
            lhs269_.f15 = rhs347_
            d_1848_v38_ = rhs348_
            lhs270_[lhs271_] = rhs349_
            lhs272_.f26 = rhs350_
            rhs351_ = default__.safeModuloInt((self).f27, (self).f27)
            rhs352_ = d_1847_v37_
            lhs273_ = globalState
            lhs273_.f24 = rhs351_
            d_1847_v37_ = rhs352_
            d_1804_v14_ = default__.fm48(not(True), False, (0) - ((d_1846_v36_).f36), globalState)

    @property
    def f38(self):
        return self._f38

class C6(T1, T2, T0):
    def  __init__(self):
        self._f29: _dafny.MultiSet = _dafny.MultiSet({})
        self._f30: _dafny.Seq = _dafny.Seq({})
        self._f27: int = int(0)
        self._f28: _dafny.Array = _dafny.Array(None, 0)
        self._f31: _dafny.MultiSet = _dafny.MultiSet({})
        self._f32: int = int(0)
        self.f34: int = int(0)
        self._f33: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f29(self):
        return self._f29
    @property
    def f30(self):
        return self._f30
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    @property
    def f31(self):
        return self._f31
    @property
    def f32(self):
        return self._f32
    def ctor__(self, f33, f34, f29, f30, f27, f28, f31, f32):
        (self)._f33 = f33
        (self).f34 = f34
        (self)._f29 = f29
        (self)._f30 = f30
        (self)._f27 = f27
        (self)._f28 = f28
        (self)._f31 = f31
        (self)._f32 = f32

    def fm3(self, p0, globalState):
        return (self).f33

    def fm4(self, p0, p1, p2, globalState):
        return len(_dafny.Set({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f33, False, (self).f33, False])): (self).f33})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rio")))}))

    def fm2(self, p0, p1, p2, globalState):
        def iife172_():
            coll56_ = _dafny.Set()
            compr_56_: int
            for compr_56_ in _dafny.IntegerRange(-947, -397):
                d_1856_v0_: int = compr_56_
                if ((-947) <= (d_1856_v0_)) and ((d_1856_v0_) < (-397)):
                    coll56_ = coll56_.union(_dafny.Set([(d_1856_v0_) - (len(_dafny.SeqWithoutIsStrInference([(self).f33, (self).f33])))]))
            return _dafny.Set(coll56_)
        return _dafny.Map({(_dafny.Set({(self).f27, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spkmp")))), 4})).isdisjoint(iife172_()
        ): (self).f32})

    def fm5(self, globalState):
        return (self).f33

    def fm6(self, p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f33, (self).f33]) for d_1857_i0_ in range(default__.abs(-678))])

    def fm7(self, globalState):
        return len((D1_DC1(_dafny.Map({self.f34: not((self).f33)}))).cf1)

    def m1(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        d_1858_v0_: str
        d_1858_v0_ = _dafny.CodePoint('f')
        d_1859_v1_: _dafny.Map
        d_1859_v1_ = _dafny.Map({d_1858_v0_: (self).f33})
        d_1860_v2_: _dafny.Seq
        d_1860_v2_ = _dafny.SeqWithoutIsStrInference([d_1859_v1_, d_1859_v1_])
        d_1861_v3_: D2
        d_1861_v3_ = D2_DC3(d_1859_v1_)
        d_1862_v6_: _dafny.Set
        d_1862_v6_ = _dafny.Set({d_1858_v0_})
        d_1863_v7_: _dafny.Array
        nw326_ = _dafny.Array(None, 22)
        nw326_[int(0)] = d_1859_v1_
        nw326_[int(1)] = (d_1859_v1_).set(d_1858_v0_, not((self).fm3((self).f29, globalState)))
        nw326_[int(2)] = d_1859_v1_
        nw326_[int(3)] = (d_1860_v2_)[default__.safeIndex(self.f34, len(d_1860_v2_))]
        nw326_[int(4)] = (d_1861_v3_).cf4
        nw326_[int(5)] = d_1859_v1_
        nw326_[int(6)] = d_1859_v1_
        nw326_[int(7)] = _dafny.Map({d_1858_v0_: (self).f33})
        def iife173_():
            coll57_ = _dafny.Map()
            compr_57_: str
            for compr_57_ in (_dafny.Map({d_1858_v0_: (self).f33})).keys.Elements:
                d_1864_v4_: str = compr_57_
                if (d_1864_v4_) in (_dafny.Map({d_1858_v0_: (self).f33})):
                    coll57_[d_1864_v4_] = (self).f33
            return _dafny.Map(coll57_)
        nw326_[int(8)] = (d_1859_v1_) | (iife173_()
        )
        nw326_[int(9)] = d_1859_v1_
        nw326_[int(10)] = d_1859_v1_
        nw326_[int(11)] = d_1859_v1_
        nw326_[int(12)] = (d_1861_v3_).cf4
        nw326_[int(13)] = d_1859_v1_
        nw326_[int(14)] = d_1859_v1_
        nw326_[int(15)] = (d_1859_v1_) | (d_1859_v1_)
        nw326_[int(16)] = (d_1860_v2_)[default__.safeIndex(self.f34, len(d_1860_v2_))]
        nw326_[int(17)] = (d_1859_v1_ if (self).f33 else d_1859_v1_)
        nw326_[int(18)] = (d_1859_v1_) | (d_1859_v1_)
        nw326_[int(19)] = d_1859_v1_
        def iife174_():
            coll58_ = _dafny.Map()
            compr_58_: str
            for compr_58_ in (d_1862_v6_).Elements:
                d_1865_v5_: str = compr_58_
                if (d_1865_v5_) in (d_1862_v6_):
                    coll58_[d_1865_v5_] = (self).f33
            return _dafny.Map(coll58_)
        nw326_[int(20)] = (iife174_()
        ) | (d_1859_v1_)
        nw326_[int(21)] = d_1859_v1_
        d_1863_v7_ = nw326_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_1863_v7_).length(0)):
            d_1866_i0_: int = guard_loop_7_
            if (True) and (((0) <= (d_1866_i0_)) and ((d_1866_i0_) < ((d_1863_v7_).length(0)))):
                (d_1863_v7_)[(d_1866_i0_)] = _dafny.Map({d_1858_v0_: (self).f33})
        d_1867_v8_: _dafny.Map
        d_1867_v8_ = _dafny.Map({not (True) or (not(not((self).f33))): (len(_dafny.SeqWithoutIsStrInference([110, len(default__.fm8(globalState))]))) > ((self).f27)})
        d_1867_v8_ = (d_1867_v8_).set((self).f33, (self).f33)
        (globalState).f26 = (self).f33
        (globalState).f18 = ((self).f30)[default__.safeIndex(len((self).f30), len((self).f30))]
        d_1858_v0_ = d_1858_v0_
        d_1868_v9_: _dafny.Set
        d_1868_v9_ = _dafny.Set({(self).f33})
        d_1869_v10_: _dafny.Map
        d_1869_v10_ = _dafny.Map({(self).f33: d_1868_v9_})
        hi9_ = len((default__.fm9((self).f33, globalState)).set(default__.safeIndex(self.f34, len(default__.fm9((self).f33, globalState))), default__.fm1(globalState)))
        for d_1870_i1_ in range(len((_dafny.SeqWithoutIsStrInference([d_1858_v0_ for d_1895_i2_ in range(default__.abs(784))])).set(default__.safeIndex(len(d_1869_v10_), len(_dafny.SeqWithoutIsStrInference([d_1858_v0_ for d_1896_i2_ in range(default__.abs(784))]))), d_1858_v0_)), hi9_):
            d_1871_v11_: _dafny.Map
            d_1871_v11_ = _dafny.Map({(self).f32: (self).f33})
            d_1872_v12_: _dafny.MultiSet
            d_1872_v12_ = _dafny.MultiSet([d_1871_v11_])
            d_1873_v13_: _dafny.Seq
            d_1873_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhq"))
            d_1874_v14_: _dafny.Map
            d_1874_v14_ = _dafny.Map({(d_1872_v12_).cardinality: d_1873_v13_})
            if (default__.safeModuloInt(len(((d_1874_v14_)[(self).f27] if ((self).f27) in (d_1874_v14_) else d_1873_v13_)), (self).f32)) == ((self).f27):
                d_1875_v15_: _dafny.Seq
                d_1875_v15_ = _dafny.SeqWithoutIsStrInference([((d_1871_v11_)[len((self).f30)] if (len((self).f30)) in (d_1871_v11_) else (self).f33)])
                d_1876_v16_: _dafny.Array
                nw327_ = _dafny.Array(None, 28)
                nw327_[int(0)] = (self).f33
                nw327_[int(1)] = (self).f33
                nw327_[int(2)] = (self).f33
                nw327_[int(3)] = (self).f33
                nw327_[int(4)] = (self).f33
                nw327_[int(5)] = (self).f33
                nw327_[int(6)] = (self).f33
                nw327_[int(7)] = (self).f33
                nw327_[int(8)] = not((self).f33)
                nw327_[int(9)] = (self).f33
                nw327_[int(10)] = (self).f33
                nw327_[int(11)] = (self).f33
                nw327_[int(12)] = (self).f33
                nw327_[int(13)] = not((self).f33)
                nw327_[int(14)] = True
                nw327_[int(15)] = (self).f33
                nw327_[int(16)] = (self).f33
                nw327_[int(17)] = (d_1875_v15_)[default__.safeIndex((self).f27, len(d_1875_v15_))]
                nw327_[int(18)] = (self).f33
                nw327_[int(19)] = (self).f33
                nw327_[int(20)] = (self).f33
                nw327_[int(21)] = (self).f33
                nw327_[int(22)] = (self).f33
                nw327_[int(23)] = not((self).f33)
                nw327_[int(24)] = (self).f33
                nw327_[int(25)] = (self).f33
                nw327_[int(26)] = (self).f33
                nw327_[int(27)] = (self).f33
                d_1876_v16_ = nw327_
                d_1877_v17_: _dafny.Map
                d_1877_v17_ = _dafny.Map({((self).f32) * (len(d_1871_v11_)): d_1876_v16_})
                d_1877_v17_ = (d_1877_v17_).set((self).f32, d_1876_v16_)
                (globalState).f26 = ((d_1867_v8_)[(len(d_1868_v9_)) >= ((self).f27)] if ((len(d_1868_v9_)) >= ((self).f27)) in (d_1867_v8_) else (self).f33)
                index238_ = default__.safeIndex(588, (d_1876_v16_).length(0))
                (d_1876_v16_)[index238_] = not ((self).f33) or ((self).f33)
                d_1878_v18_: _dafny.Array
                nw328_ = _dafny.Array(int(0), 2)
                d_1878_v18_ = nw328_
                index239_ = default__.safeIndex(312, (d_1878_v18_).length(0))
                (d_1878_v18_)[index239_] = (self).f27
                index240_ = default__.safeIndex(588, (d_1876_v16_).length(0))
                index241_ = default__.safeIndex(312, (d_1878_v18_).length(0))
                rhs353_ = (self).fm3(_dafny.MultiSet([(self).f33]), globalState)
                rhs354_ = -998
                rhs355_ = d_1876_v16_
                lhs274_ = d_1876_v16_
                lhs275_ = default__.safeIndex(588, (d_1876_v16_).length(0))
                lhs276_ = d_1878_v18_
                lhs277_ = default__.safeIndex(312, (d_1878_v18_).length(0))
                lhs274_[lhs275_] = rhs353_
                lhs276_[lhs277_] = rhs354_
                d_1876_v16_ = rhs355_
                d_1879_v19_: _dafny.MultiSet
                d_1879_v19_ = _dafny.MultiSet([(len(d_1875_v15_)) - (d_1870_i1_), (0) - (d_1870_i1_), (self).f27, len(d_1873_v13_), ((d_1878_v18_)[default__.safeIndex(312, (d_1878_v18_).length(0))]) * (d_1870_i1_)])
                d_1879_v19_ = d_1879_v19_
                d_1880_v20_: _dafny.Map
                d_1880_v20_ = _dafny.Map({(d_1876_v16_)[default__.safeIndex(588, (d_1876_v16_).length(0))]: (self).fm7(globalState)})
                d_1881_v22_: _dafny.Set
                def iife175_():
                    coll59_ = _dafny.Map()
                    compr_59_: int
                    for compr_59_ in _dafny.IntegerRange(215, 925):
                        d_1882_v21_: int = compr_59_
                        if ((215) <= (d_1882_v21_)) and ((d_1882_v21_) < (925)):
                            coll59_[default__.safeDivisionInt(d_1882_v21_, len(_dafny.Set({self.f34})))] = (self).f33
                    return _dafny.Map(coll59_)
                d_1881_v22_ = _dafny.Set({len(iife175_()
                )})
                d_1883_v23_: _dafny.Map
                d_1883_v23_ = _dafny.Map({d_1880_v20_: len((d_1881_v22_ if (d_1876_v16_)[default__.safeIndex(588, (d_1876_v16_).length(0))] else default__.fm10(globalState)))})
                d_1883_v23_ = (d_1883_v23_ if (d_1876_v16_)[default__.safeIndex(588, (d_1876_v16_).length(0))] else d_1883_v23_)
            elif True:
                r2 = (self.f34) + ((d_1870_i1_) - ((self).f32))
                r0 = (self).f32
                d_1884_v24_: _dafny.Map
                d_1884_v24_ = _dafny.Map({len(d_1868_v9_): d_1867_v8_})
                d_1885_v25_: _dafny.Seq
                d_1885_v25_ = _dafny.SeqWithoutIsStrInference([d_1867_v8_, d_1867_v8_, ((d_1884_v24_)[d_1870_i1_] if (d_1870_i1_) in (d_1884_v24_) else d_1867_v8_)])
                d_1886_v26_: _dafny.Seq
                d_1886_v26_ = _dafny.SeqWithoutIsStrInference([(d_1885_v25_)[default__.safeIndex(d_1870_i1_, len(d_1885_v25_))], _dafny.Map({(self).f33: (self).f33}), (d_1867_v8_) | (d_1867_v8_), (_dafny.Map({(self).f33: (self).f33})) | (d_1867_v8_), d_1867_v8_])
                d_1886_v26_ = d_1886_v26_
                d_1887_v27_: _dafny.Map
                d_1887_v27_ = _dafny.Map({((d_1859_v1_)[d_1858_v0_] if (d_1858_v0_) in (d_1859_v1_) else (self).f33): (self).f27})
                d_1888_v28_: T1
                nw329_ = C1()
                nw329_.ctor__(False, (self).f29, (((self).f30).set(default__.safeIndex(d_1870_i1_, len((self).f30)), len(d_1887_v27_))).set(default__.safeIndex(self.f34, len(((self).f30).set(default__.safeIndex(d_1870_i1_, len((self).f30)), len(d_1887_v27_)))), (self).f27), (self).f27, (self).f28)
                d_1888_v28_ = nw329_
                d_1889_v29_: _dafny.Map
                d_1889_v29_ = _dafny.Map({d_1873_v13_: d_1888_v28_})
                d_1889_v29_ = (d_1889_v29_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nl")), d_1888_v28_)
                d_1890_v30_: _dafny.Seq
                d_1890_v30_ = _dafny.SeqWithoutIsStrInference([not((self).f33), (self).f33, (self).f33, (self).f33, not(not((self).f33))])
                d_1891_v31_: _dafny.Array
                nw330_ = _dafny.Array(None, 1)
                nw330_[int(0)] = d_1890_v30_
                d_1891_v31_ = nw330_
                index242_ = default__.safeIndex(440, (d_1891_v31_).length(0))
                (d_1891_v31_)[index242_] = d_1890_v30_
                index243_ = default__.safeIndex(440, (d_1891_v31_).length(0))
                (d_1891_v31_)[index243_] = (d_1890_v30_).set(default__.safeIndex((self).f27, len(d_1890_v30_)), (self).f33)
            d_1892_v32_: _dafny.Array
            nw331_ = _dafny.Array(None, 27)
            d_1892_v32_ = nw331_
            d_1893_v33_: _dafny.Seq
            d_1893_v33_ = _dafny.SeqWithoutIsStrInference([(self).f33])
            d_1894_v34_: T2
            nw332_ = C2()
            nw332_.ctor__(_dafny.MultiSet(d_1893_v33_), (self).f32)
            d_1894_v34_ = nw332_
            index244_ = default__.safeIndex(743, (d_1892_v32_).length(0))
            (d_1892_v32_)[index244_] = (d_1894_v34_ if (self).f33 else d_1894_v34_)
            index245_ = default__.safeIndex(743, (d_1892_v32_).length(0))
            (d_1892_v32_)[index245_] = d_1894_v34_
            rhs356_ = d_1870_i1_
            rhs357_ = (self).f33
            lhs278_ = globalState
            lhs279_ = globalState
            lhs278_.f15 = rhs356_
            lhs279_.f26 = rhs357_
            d_1873_v13_ = d_1873_v13_
        d_1897_v35_: _dafny.Seq
        d_1897_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asxhdikym"))
        r0 = default__.safeModuloInt(default__.fm1(globalState), (len(d_1897_v35_)) * (self.f34))
        r1 = (self).f27
        r2 = (self).f27
        return r0, r1, r2

    def m2(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        d_1898_v0_: _dafny.Seq
        d_1898_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1899_v1_: _dafny.Seq
        d_1899_v1_ = _dafny.SeqWithoutIsStrInference([d_1898_v0_])
        d_1900_v2_: _dafny.Seq
        d_1900_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmsbclr"))
        pat_let_tv43_ = d_1899_v1_
        d_1901_v3_: bool
        d_1902_v4_: int
        out47_: bool
        out48_: int
        def iife176_(_pat_let58_0):
            def iife177_(d_1903_dt__update__tmp_h0_):
                def iife178_(_pat_let59_0):
                    def iife179_(d_1904_dt__update_hcf21_h0_):
                        return D6_DC12((d_1903_dt__update__tmp_h0_).cf20, d_1904_dt__update_hcf21_h0_, (d_1903_dt__update__tmp_h0_).cf22)
                    return iife179_(_pat_let59_0)
                return iife178_(pat_let_tv43_)
            return iife177_(_pat_let58_0)
        out47_, out48_ = default__.m0((self).fm5(globalState), ((self).f27) + (len((iife176_(D6_DC12((self).f33, d_1899_v1_, d_1900_v2_))).cf21)), (0) - (self.f34), self.f34, globalState)
        d_1901_v3_ = out47_
        d_1902_v4_ = out48_
        d_1905_i0_: int
        d_1905_i0_ = 0
        with _dafny.label("8"):
            while ((self).f29).ispropersubset((self).f29):
                with _dafny.c_label("8"):
                    if (d_1905_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_1905_i0_ = (d_1905_i0_) + (1)
                    d_1906_v5_: _dafny.Map
                    d_1906_v5_ = _dafny.Map({d_1901_v3_: ((self).f27) < ((self).f32)})
                    d_1906_v5_ = (d_1906_v5_) | (d_1906_v5_)
                    d_1907_v6_: _dafny.Map
                    d_1907_v6_ = _dafny.Map({self.f34: not(d_1901_v3_)})
                    (globalState).f15 = ((len((self).fm2((self).f33, d_1901_v3_, p1, globalState))) * (len(d_1907_v6_))) + (len(d_1907_v6_))
                    d_1908_v7_: _dafny.Set
                    d_1908_v7_ = _dafny.Set({(self).fm5(globalState)})
                    (globalState).f26 = ((len(d_1898_v0_)) != ((self).f32)) in ((d_1908_v7_).intersection(d_1908_v7_))
                    d_1909_v8_: _dafny.Array
                    nw333_ = _dafny.Array(None, 24)
                    nw333_[int(0)] = d_1901_v3_
                    nw333_[int(1)] = (self).f33
                    nw333_[int(2)] = not(p0)
                    nw333_[int(3)] = False
                    nw333_[int(4)] = (self).f33
                    nw333_[int(5)] = p1
                    nw333_[int(6)] = (self).f33
                    nw333_[int(7)] = p1
                    nw333_[int(8)] = d_1901_v3_
                    nw333_[int(9)] = d_1901_v3_
                    nw333_[int(10)] = p1
                    nw333_[int(11)] = d_1901_v3_
                    nw333_[int(12)] = not(p0)
                    nw333_[int(13)] = default__.fm0((self).f27, globalState)
                    nw333_[int(14)] = (self).f33
                    nw333_[int(15)] = True
                    nw333_[int(16)] = p1
                    nw333_[int(17)] = p1
                    nw333_[int(18)] = (d_1898_v0_)[default__.safeIndex(p3, len(d_1898_v0_))]
                    nw333_[int(19)] = d_1901_v3_
                    nw333_[int(20)] = p0
                    nw333_[int(21)] = p1
                    nw333_[int(22)] = p1
                    nw333_[int(23)] = (self).f33
                    d_1909_v8_ = nw333_
                    d_1910_v9_: _dafny.Map
                    d_1910_v9_ = _dafny.Map({not(p0): d_1909_v8_})
                    d_1910_v9_ = (d_1910_v9_).set(p1, d_1909_v8_)
                    pass
            pass
        d_1901_v3_ = (self).f33
        (globalState).f15 = ((self).f32) + (d_1902_v4_)
        hi10_ = (self).f32
        for d_1911_i1_ in range(395, hi10_):
            (globalState).f18 = (0) - (self.f34)
            d_1912_v10_: str
            d_1912_v10_ = _dafny.CodePoint('t')
            d_1913_v11_: _dafny.Set
            d_1913_v11_ = _dafny.Set({d_1912_v10_})
            if (d_1912_v10_) in (d_1913_v11_):
                d_1914_v12_: _dafny.Map
                d_1914_v12_ = _dafny.Map({self.f34: d_1901_v3_})
                d_1915_v13_: D1
                d_1915_v13_ = D1_DC1(d_1914_v12_)
                d_1916_v14_: _dafny.Map
                d_1916_v14_ = _dafny.Map({d_1915_v13_: p0})
                d_1916_v14_ = (d_1916_v14_).set(d_1915_v13_, not(d_1901_v3_))
                (globalState).f0 = d_1911_i1_
                d_1917_v15_: _dafny.Set
                d_1918_v16_: int
                d_1919_v17_: _dafny.Array
                out49_: _dafny.Set
                out50_: int
                out51_: _dafny.Array
                out49_, out50_, out51_ = (self).m4(default__.fm0((self).f27, globalState), self.f34, p1, p1, globalState)
                d_1917_v15_ = out49_
                d_1918_v16_ = out50_
                d_1919_v17_ = out51_
                d_1900_v2_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1920_i2_ in range(default__.abs(370))])) + ((d_1900_v2_ if p1 else d_1900_v2_))
                d_1921_v18_: C2
                nw334_ = C2()
                nw334_.ctor__((self).f31, (0) - (len(d_1898_v0_)))
                d_1921_v18_ = nw334_
            elif True:
                d_1922_v19_: _dafny.Array
                nw335_ = _dafny.Array(_dafny.Seq({}), 6)
                d_1922_v19_ = nw335_
                index246_ = default__.safeIndex(398, (d_1922_v19_).length(0))
                (d_1922_v19_)[index246_] = (self).f30
                d_1923_v20_: _dafny.Map
                d_1923_v20_ = _dafny.Map({p0: p1})
                index247_ = default__.safeIndex(398, (d_1922_v19_).length(0))
                rhs358_ = default__.fm14(((self).f32) <= ((self).f27), ((d_1923_v20_)[(self).f33] if ((self).f33) in (d_1923_v20_) else p1), p1, (self).f27, globalState)
                rhs359_ = p3
                rhs360_ = (self).f32
                rhs361_ = p0
                rhs362_ = (self).f30
                lhs280_ = globalState
                lhs281_ = globalState
                lhs282_ = d_1922_v19_
                lhs283_ = default__.safeIndex(398, (d_1922_v19_).length(0))
                d_1900_v2_ = rhs358_
                lhs280_.f21 = rhs359_
                lhs281_.f0 = rhs360_
                d_1901_v3_ = rhs361_
                lhs282_[lhs283_] = rhs362_
                d_1924_v21_: _dafny.Array
                def lambda81_(d_1925_p2_):
                    def lambda82_(d_1926_i3_):
                        return _dafny.Set({d_1925_p2_})

                    return lambda82_

                init40_ = lambda81_(p2)
                nw336_ = _dafny.Array(None, 9)
                for i0_40_ in range(nw336_.length(0)):
                    nw336_[i0_40_] = init40_(i0_40_)
                d_1924_v21_ = nw336_
                d_1927_v22_: _dafny.Array
                nw337_ = _dafny.Array(None, 6)
                d_1927_v22_ = nw337_
                d_1928_v23_: _dafny.Map
                d_1928_v23_ = _dafny.Map({d_1924_v21_: d_1927_v22_})
                d_1928_v23_ = (d_1928_v23_).set(d_1924_v21_, d_1927_v22_)
                d_1929_v24_: _dafny.Map
                d_1929_v24_ = _dafny.Map({default__.fm0((self).f27, globalState): (self).f28})
                d_1930_v25_: C1
                nw338_ = C1()
                nw338_.ctor__(p1, (_dafny.MultiSet(default__.fm13(p2, (self).f32, p0, 421, globalState))) | (_dafny.MultiSet([p1])), (((d_1922_v19_)[default__.safeIndex(398, (d_1922_v19_).length(0))]).set(default__.safeIndex((self).f27, len((d_1922_v19_)[default__.safeIndex(398, (d_1922_v19_).length(0))])), p3)).set(default__.safeIndex(default__.fm1(globalState), len(((d_1922_v19_)[default__.safeIndex(398, (d_1922_v19_).length(0))]).set(default__.safeIndex((self).f27, len((d_1922_v19_)[default__.safeIndex(398, (d_1922_v19_).length(0))])), p3))), default__.fm1(globalState)), p3, ((d_1929_v24_)[(self).f33] if ((self).f33) in (d_1929_v24_) else (self).f28))
                d_1930_v25_ = nw338_
                d_1930_v25_ = d_1930_v25_
                d_1931_v26_: _dafny.Array
                nw339_ = _dafny.Array(None, 2)
                nw339_[int(0)] = p1
                nw339_[int(1)] = False
                d_1931_v26_ = nw339_
                d_1932_v27_: _dafny.MultiSet
                d_1932_v27_ = _dafny.MultiSet([d_1931_v26_])
                d_1933_v28_: bool
                d_1934_v29_: _dafny.Array
                d_1935_v30_: _dafny.Map
                out52_: bool
                out53_: _dafny.Array
                out54_: _dafny.Map
                out52_, out53_, out54_ = (d_1930_v25_).m5((d_1932_v27_).isdisjoint(_dafny.MultiSet([d_1931_v26_])), p1, globalState)
                d_1933_v28_ = out52_
                d_1934_v29_ = out53_
                d_1935_v30_ = out54_
                d_1936_v32_: _dafny.Set
                d_1936_v32_ = _dafny.Set({d_1911_i1_})
                d_1937_v33_: _dafny.Map
                d_1937_v33_ = _dafny.Map({d_1911_i1_: len(d_1936_v32_)})
                d_1938_v34_: _dafny.MultiSet
                d_1938_v34_ = _dafny.MultiSet([_dafny.Map({d_1911_i1_: ((self).f31).cardinality}), d_1937_v33_])
                d_1939_v35_: D20
                d_1939_v35_ = D20_DC50(d_1938_v34_)
                d_1940_v36_: T0
                nw340_ = C5()
                def iife180_():
                    coll60_ = _dafny.Map()
                    compr_60_: _dafny.Map
                    for compr_60_ in ((d_1939_v35_).cf92).Elements:
                        d_1941_v31_: _dafny.Map = compr_60_
                        if (d_1941_v31_) in ((d_1939_v35_).cf92):
                            coll60_[d_1941_v31_] = 4
                    return _dafny.Map(coll60_)
                nw340_.ctor__(iife180_()
                , _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1, False, d_1930_v25_.f35])), (self).f30, self.f34, (self).f28)
                d_1940_v36_ = nw340_
                d_1942_v37_: _dafny.Map
                d_1942_v37_ = _dafny.Map({d_1940_v36_: p3})
                d_1943_v38_: D10
                d_1943_v38_ = D10_DC22(((d_1942_v37_)[d_1940_v36_] if (d_1940_v36_) in (d_1942_v37_) else 596), (self).f31)
                pat_let_tv44_ = d_1898_v0_
                pat_let_tv45_ = d_1898_v0_
                def iife181_(_pat_let60_0):
                    def iife182_(d_1944_dt__update__tmp_h1_):
                        def iife183_(_pat_let61_0):
                            def iife184_(d_1945_dt__update_hcf41_h0_):
                                return D10_DC22((d_1944_dt__update__tmp_h1_).cf40, d_1945_dt__update_hcf41_h0_)
                            return iife184_(_pat_let61_0)
                        return iife183_(((self).f31).intersection(_dafny.MultiSet([(pat_let_tv44_)[default__.safeIndex(-773, len(pat_let_tv45_))]])))
                    return iife182_(_pat_let60_0)
                d_1943_v38_ = iife181_(d_1943_v38_)
            (globalState).f26 = ((d_1898_v0_) + (d_1898_v0_)) < ((d_1898_v0_) + (_dafny.SeqWithoutIsStrInference([not(p0)])))
            d_1946_v39_: _dafny.Set
            d_1946_v39_ = _dafny.Set({(self).f30, (self).f30})
            (globalState).f26 = ((self).f30) not in ((d_1946_v39_ if p1 else d_1946_v39_))
        d_1947_v40_: _dafny.Map
        d_1947_v40_ = _dafny.Map({550: (_dafny.MultiSet((self).f30)).cardinality})
        d_1901_v3_ = not ((len(d_1947_v40_)) <= ((self).f32)) or ((self).f33)
        d_1948_v41_: _dafny.Array
        nw341_ = _dafny.Array(False, 1)
        d_1948_v41_ = nw341_
        r0 = (d_1948_v41_ if ((self).f32) > (self.f34) else d_1948_v41_)
        r1 = ((_dafny.MultiSet([(self).fm3((self).f31, globalState), d_1901_v3_, (self).f33])).cardinality) + (((self).f31).cardinality)
        return r0, r1

    def m3(self, p0, p1, p2, p3, globalState):
        d_1949_v0_: _dafny.Array
        def lambda83_(d_1950_i0_):
            return (self).f33

        init41_ = lambda83_
        nw342_ = _dafny.Array(None, 27)
        for i0_41_ in range(nw342_.length(0)):
            nw342_[i0_41_] = init41_(i0_41_)
        d_1949_v0_ = nw342_
        d_1949_v0_ = d_1949_v0_
        d_1951_v1_: _dafny.Seq
        d_1951_v1_ = _dafny.SeqWithoutIsStrInference([d_1949_v0_, d_1949_v0_, d_1949_v0_])
        d_1952_v2_: _dafny.Map
        d_1952_v2_ = _dafny.Map({p2: (p3 if p0 else p0)})
        d_1953_v3_: _dafny.Map
        d_1953_v3_ = _dafny.Map({(self).f33: (self).f33})
        d_1954_v4_: _dafny.MultiSet
        d_1954_v4_ = _dafny.MultiSet([(self).f27])
        rhs363_ = d_1951_v1_
        rhs364_ = (self).fm3((self).f31, globalState)
        rhs365_ = ((d_1952_v2_) | (_dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yskuy"))).set(default__.safeIndex((self).f27, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yskuy")))), p1): p3})) if (self).f33 else (d_1952_v2_) | (d_1952_v2_))
        rhs366_ = not (not((self).f33)) or (((d_1953_v3_)[p0] if (p0) in (d_1953_v3_) else p3))
        rhs367_ = (d_1954_v4_).issubset(d_1954_v4_)
        lhs284_ = globalState
        lhs285_ = globalState
        lhs286_ = globalState
        d_1951_v1_ = rhs363_
        lhs284_.f26 = rhs364_
        d_1952_v2_ = rhs365_
        lhs285_.f26 = rhs366_
        lhs286_.f26 = rhs367_
        d_1955_i1_: int
        d_1955_i1_ = 0
        with _dafny.label("9"):
            while (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msled"))) == (_dafny.SeqWithoutIsStrInference([p1 for d_1965_i2_ in range(default__.abs(492))])):
                with _dafny.c_label("9"):
                    if (d_1955_i1_) >= (100):
                        raise _dafny.Break("9")
                    d_1955_i1_ = (d_1955_i1_) + (1)
                    def iife185_():
                        coll61_ = _dafny.Set()
                        compr_61_: int
                        for compr_61_ in _dafny.IntegerRange(-80, 763):
                            d_1956_v5_: int = compr_61_
                            if ((-80) <= (d_1956_v5_)) and ((d_1956_v5_) < (763)):
                                coll61_ = coll61_.union(_dafny.Set([default__.safeModuloInt(d_1956_v5_, (self).f32)]))
                        return _dafny.Set(coll61_)
                    (globalState).f0 = len(iife185_()
                    )
                    d_1957_v6_: _dafny.Map
                    d_1957_v6_ = _dafny.Map({p1: _dafny.Map({self.f34: (p2)[default__.safeIndex((self).f32, len(p2))]})})
                    d_1958_v7_: _dafny.Map
                    d_1958_v7_ = _dafny.Map({self.f34: _dafny.CodePoint('j')})
                    d_1959_v8_: C2
                    nw343_ = C2()
                    nw343_.ctor__((self).f31, len(((d_1957_v6_)[p1] if (p1) in (d_1957_v6_) else d_1958_v7_)))
                    d_1959_v8_ = nw343_
                    index248_ = default__.safeIndex(899, (d_1949_v0_).length(0))
                    (d_1949_v0_)[index248_] = p3
                    d_1960_v9_: T1
                    nw344_ = C1()
                    nw344_.ctor__(p3, ((self).f31).set(p3, default__.abs(default__.fm1(globalState))), _dafny.SeqWithoutIsStrInference([self.f34]), self.f34, (self).f28)
                    d_1960_v9_ = nw344_
                    d_1961_v10_: _dafny.Map
                    d_1961_v10_ = _dafny.Map({d_1960_v9_: self.f34})
                    index249_ = default__.safeIndex(899, (d_1949_v0_).length(0))
                    (d_1949_v0_)[index249_] = ((self.f34) < (len(d_1961_v10_))) or (p0)
                    d_1962_v11_: _dafny.Set
                    d_1962_v11_ = _dafny.Set({p2, p2})
                    d_1963_v12_: _dafny.Seq
                    d_1963_v12_ = _dafny.SeqWithoutIsStrInference([False])
                    d_1964_v13_: _dafny.Map
                    d_1964_v13_ = _dafny.Map({(d_1960_v9_).fm4(d_1962_v11_, (self).f33, d_1963_v12_, globalState): p0})
                    d_1964_v13_ = (d_1964_v13_).set((self).f32, p3)
                    pass
            pass
        d_1966_v14_: _dafny.Set
        d_1966_v14_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([self.f34, self.f34]), (self).f30, (self).f30, (self).f30, (self).f30})
        (globalState).f26 = (d_1966_v14_).ispropersubset(default__.fm50(globalState))
        hi11_ = self.f34
        for d_1967_i3_ in range((0) - ((self).f32), hi11_):
            d_1968_v15_: _dafny.Map
            d_1968_v15_ = _dafny.Map({p3: p2})
            d_1969_v16_: _dafny.Map
            d_1969_v16_ = _dafny.Map({len(((d_1968_v15_)[p0] if (p0) in (d_1968_v15_) else p2)): d_1967_i3_})
            d_1970_v18_: _dafny.Array
            nw345_ = _dafny.Array(None, 14)
            nw345_[int(0)] = d_1969_v16_
            nw345_[int(1)] = default__.fm15((self).f27, d_1967_i3_, globalState)
            def iife186_():
                coll62_ = _dafny.Map()
                compr_62_: int
                for compr_62_ in _dafny.IntegerRange(-310, 360):
                    d_1971_v17_: int = compr_62_
                    if ((-310) <= (d_1971_v17_)) and ((d_1971_v17_) < (360)):
                        coll62_[default__.safeModuloInt(d_1971_v17_, (self).f27)] = -274
                return _dafny.Map(coll62_)
            nw345_[int(2)] = iife186_()
            
            nw345_[int(3)] = d_1969_v16_
            nw345_[int(4)] = default__.fm30(globalState)
            nw345_[int(5)] = d_1969_v16_
            nw345_[int(6)] = d_1969_v16_
            nw345_[int(7)] = _dafny.Map({d_1967_i3_: d_1967_i3_})
            nw345_[int(8)] = d_1969_v16_
            nw345_[int(9)] = d_1969_v16_
            nw345_[int(10)] = d_1969_v16_
            nw345_[int(11)] = d_1969_v16_
            nw345_[int(12)] = d_1969_v16_
            nw345_[int(13)] = default__.fm42((self).f32, p3, d_1967_i3_, globalState)
            d_1970_v18_ = nw345_
            d_1972_v19_: _dafny.MultiSet
            d_1972_v19_ = _dafny.MultiSet([d_1970_v18_, d_1970_v18_, d_1970_v18_])
            (globalState).f0 = ((d_1972_v19_)[d_1970_v18_] if (d_1970_v18_) in (d_1972_v19_) else (self.f34) - ((self).f32))
            d_1973_v20_: _dafny.Seq
            d_1973_v20_ = _dafny.SeqWithoutIsStrInference([D18_DC47(len(((self).f30).set(default__.safeIndex(d_1967_i3_, len((self).f30)), d_1967_i3_)), p1)])
            (globalState).f3 = len(default__.fm51((d_1973_v20_) + (_dafny.SeqWithoutIsStrInference([D18_DC47((self).f27, p1)])), (d_1953_v3_) != (d_1953_v3_), (d_1968_v15_) | (d_1968_v15_), globalState))
            (globalState).f0 = (default__.safeDivisionInt(d_1967_i3_, self.f34) if p0 else 861)
            d_1974_v21_: _dafny.Array
            nw346_ = _dafny.Array(int(0), 24)
            d_1974_v21_ = nw346_
            d_1975_v22_: _dafny.Seq
            d_1975_v22_ = _dafny.SeqWithoutIsStrInference([(self).f33, p3, p0])
            index250_ = default__.safeIndex(837, (d_1974_v21_).length(0))
            (d_1974_v21_)[index250_] = (d_1967_i3_) + ((self).fm4(_dafny.Set({p2, _dafny.SeqWithoutIsStrInference([p1 for d_1976_i4_ in range(default__.abs(566))])}), (self).f33, d_1975_v22_, globalState))
            d_1977_v23_: _dafny.Seq
            d_1977_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbd"))
            index251_ = default__.safeIndex(837, (d_1974_v21_).length(0))
            rhs368_ = (self).f32
            rhs369_ = d_1970_v18_
            rhs370_ = (d_1977_v23_) + (d_1977_v23_)
            lhs287_ = d_1974_v21_
            lhs288_ = default__.safeIndex(837, (d_1974_v21_).length(0))
            lhs287_[lhs288_] = rhs368_
            d_1970_v18_ = rhs369_
            d_1977_v23_ = rhs370_
        d_1978_v24_: _dafny.Map
        d_1978_v24_ = _dafny.Map({76: p3})
        d_1979_v26_: _dafny.Map
        d_1979_v26_ = _dafny.Map({(self).f32: (self).f27})
        d_1980_v28_: D1
        d_1980_v28_ = D1_DC1(d_1978_v24_)
        d_1981_v29_: _dafny.Seq
        d_1981_v29_ = _dafny.SeqWithoutIsStrInference([(self).f33, p0])
        d_1982_v30_: _dafny.Seq
        d_1982_v30_ = _dafny.SeqWithoutIsStrInference([d_1978_v24_])
        d_1983_v32_: _dafny.Set
        d_1983_v32_ = _dafny.Set({((self).f30)[default__.safeIndex((self).f32, len((self).f30))], (0) - (len(d_1953_v3_)), self.f34, (self).f32, self.f34})
        d_1984_v34_: _dafny.Array
        nw347_ = _dafny.Array(None, 19)
        nw347_[int(0)] = d_1978_v24_
        nw347_[int(1)] = _dafny.Map({(self).f32: (self).f33})
        nw347_[int(2)] = d_1978_v24_
        def iife187_():
            coll63_ = _dafny.Map()
            compr_63_: int
            for compr_63_ in (d_1979_v26_).keys.Elements:
                d_1985_v25_: int = compr_63_
                if (d_1985_v25_) in (d_1979_v26_):
                    coll63_[(d_1985_v25_) - (len((self).f30))] = (self).f33
            return _dafny.Map(coll63_)
        def iife188_():
            coll64_ = _dafny.Map()
            compr_64_: int
            for compr_64_ in _dafny.IntegerRange(432, -747):
                d_1986_v27_: int = compr_64_
                if ((432) <= (d_1986_v27_)) and ((d_1986_v27_) < (-747)):
                    coll64_[default__.safeDivisionInt(d_1986_v27_, (self).f32)] = p3
            return _dafny.Map(coll64_)
        nw347_[int(3)] = (iife187_()
        ) | (iife188_()
        )
        nw347_[int(4)] = d_1978_v24_
        nw347_[int(5)] = ((d_1980_v28_).cf1).set(len(d_1979_v26_), p3)
        nw347_[int(6)] = d_1978_v24_
        nw347_[int(7)] = (d_1978_v24_) | (_dafny.Map({len(_dafny.Map({len(d_1981_v29_): p3})): p3}))
        nw347_[int(8)] = (d_1982_v30_)[default__.safeIndex(len(d_1953_v3_), len(d_1982_v30_))]
        nw347_[int(9)] = d_1978_v24_
        nw347_[int(10)] = _dafny.Map({self.f34: False})
        nw347_[int(11)] = default__.fm48(p3, p3, self.f34, globalState)
        nw347_[int(12)] = _dafny.Map({self.f34: not(p3)})
        def iife189_():
            coll65_ = _dafny.Map()
            compr_65_: int
            for compr_65_ in (d_1983_v32_).Elements:
                d_1987_v31_: int = compr_65_
                if (d_1987_v31_) in (d_1983_v32_):
                    coll65_[default__.safeDivisionInt(d_1987_v31_, (D13_DC32((self).f32, (self).f32, (self).f30, p3)).cf57)] = (self).f33
            return _dafny.Map(coll65_)
        nw347_[int(13)] = iife189_()
        
        nw347_[int(14)] = (d_1978_v24_) | (default__.fm48(p0, p3, (self).f27, globalState))
        nw347_[int(15)] = d_1978_v24_
        def iife190_():
            coll66_ = _dafny.Map()
            compr_66_: int
            for compr_66_ in ((self).f30).Elements:
                d_1988_v33_: int = compr_66_
                if (d_1988_v33_) in ((self).f30):
                    coll66_[(d_1988_v33_) + (len(d_1953_v3_))] = (self).f33
            return _dafny.Map(coll66_)
        nw347_[int(16)] = (d_1978_v24_) | (iife190_()
        )
        nw347_[int(17)] = (_dafny.Map({(self).f32: p3})) | (d_1978_v24_)
        nw347_[int(18)] = d_1978_v24_
        d_1984_v34_ = nw347_
        d_1984_v34_ = d_1984_v34_

    def m4(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_1989_v0_: _dafny.Map
        d_1989_v0_ = _dafny.Map({(self).f32: (self).f32})
        d_1990_v1_: _dafny.Seq
        d_1990_v1_ = _dafny.SeqWithoutIsStrInference([(self).f33, p3])
        d_1991_v2_: _dafny.Map
        d_1991_v2_ = _dafny.Map({d_1989_v0_: len(d_1990_v1_)})
        d_1992_v3_: _dafny.Seq
        d_1992_v3_ = _dafny.SeqWithoutIsStrInference([(self).f30, (self).f30, (self).f30])
        d_1993_v4_: _dafny.Seq
        d_1993_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfggacs"))
        d_1994_v5_: str
        d_1994_v5_ = _dafny.CodePoint('p')
        d_1995_v6_: C5
        nw348_ = C5()
        nw348_.ctor__(d_1991_v2_, (self).f29, (d_1992_v3_)[default__.safeIndex((self).f32, len(d_1992_v3_))], (len((d_1993_v4_).set(default__.safeIndex(p1, len(d_1993_v4_)), d_1994_v5_))) - ((self).f27), (self).f28)
        d_1995_v6_ = nw348_
        d_1996_v7_: C3
        nw349_ = C3()
        nw349_.ctor__(90, (self).f28)
        d_1996_v7_ = nw349_
        d_1997_v9_: _dafny.Set
        d_1997_v9_ = _dafny.Set({self.f34, self.f34})
        d_1998_v10_: D11
        def iife191_():
            coll67_ = _dafny.Set()
            compr_67_: int
            for compr_67_ in _dafny.IntegerRange(-381, 691):
                d_1999_v8_: int = compr_67_
                if ((-381) <= (d_1999_v8_)) and ((d_1999_v8_) < (691)):
                    coll67_ = coll67_.union(_dafny.Set([default__.safeModuloInt(d_1999_v8_, p1)]))
            return _dafny.Set(coll67_)
        d_1998_v10_ = D11_DC23(_dafny.MultiSet([iife191_()
, d_1997_v9_, d_1997_v9_]))
        source29_ = (d_1998_v10_ if p0 else d_1998_v10_)
        if source29_.is_DC24:
            d_2000___mcc_h0_ = source29_.cf43
            d_2001___mcc_h1_ = source29_.cf44
            d_2002_cf44_ = d_2001___mcc_h1_
            d_2003_cf43_ = d_2000___mcc_h0_
            d_2004_v11_: _dafny.Array
            def lambda84_(d_2005_i0_):
                return (d_2005_i0_) * ((self).f32)

            init42_ = lambda84_
            nw350_ = _dafny.Array(None, 18)
            for i0_42_ in range(nw350_.length(0)):
                nw350_[i0_42_] = init42_(i0_42_)
            d_2004_v11_ = nw350_
            d_2004_v11_ = d_2004_v11_
            d_2006_v12_: _dafny.Seq
            d_2006_v12_ = _dafny.SeqWithoutIsStrInference([d_1990_v1_, (_dafny.SeqWithoutIsStrInference([(self).f33])).set(default__.safeIndex((self).f27, len(_dafny.SeqWithoutIsStrInference([(self).f33]))), p0)])
            d_2007_v13_: D6
            d_2007_v13_ = D6_DC12(False, d_2006_v12_, default__.fm31(p3, p2, globalState))
            d_2007_v13_ = d_2007_v13_
            d_2008_v14_: _dafny.Array
            def lambda85_(d_2009_v0_):
                def lambda86_(d_2010_i1_):
                    return d_2009_v0_

                return lambda86_

            init43_ = lambda85_(d_1989_v0_)
            nw351_ = _dafny.Array(None, 11)
            for i0_43_ in range(nw351_.length(0)):
                nw351_[i0_43_] = init43_(i0_43_)
            d_2008_v14_ = nw351_
            source30_ = D16_DC39(d_2008_v14_)
            if source30_.is_DC40:
                d_2011___mcc_h4_ = source30_.cf74
                d_2012___mcc_h5_ = source30_.cf75
                d_2013___mcc_h6_ = source30_.cf76
                d_2014_cf76_ = d_2013___mcc_h6_
                d_2015_cf75_ = d_2012___mcc_h5_
                d_2016_cf74_ = d_2011___mcc_h4_
                d_2017_v15_: _dafny.Seq
                d_2017_v15_ = _dafny.SeqWithoutIsStrInference([d_1993_v4_])
                d_1993_v4_ = (d_2017_v15_)[default__.safeIndex(d_2015_cf75_, len(d_2017_v15_))]
                d_2018_v16_: _dafny.Seq
                d_2018_v16_ = _dafny.SeqWithoutIsStrInference([d_2003_cf43_])
                d_2019_v17_: _dafny.Seq
                d_2019_v17_ = _dafny.SeqWithoutIsStrInference([d_1989_v0_])
                rhs371_ = (self).f30
                rhs372_ = (p0 if (self).f33 else default__.fm0(d_2014_cf76_, globalState))
                rhs373_ = d_2019_v17_
                lhs289_ = globalState
                d_2018_v16_ = rhs371_
                lhs289_.f26 = rhs372_
                d_2019_v17_ = rhs373_
                d_2016_cf74_ = not(((self).f27) <= ((d_2015_cf75_) * (self.f34)))
                d_2004_v11_ = d_2004_v11_
            elif source30_.is_DC41:
                d_2020___mcc_h7_ = source30_.cf77
                d_2021_cf77_ = d_2020___mcc_h7_
                d_2022_v19_: _dafny.MultiSet
                def iife192_():
                    coll68_ = _dafny.Set()
                    compr_68_: int
                    for compr_68_ in _dafny.IntegerRange(799, 419):
                        d_2023_v18_: int = compr_68_
                        if ((799) <= (d_2023_v18_)) and ((d_2023_v18_) < (419)):
                            coll68_ = coll68_.union(_dafny.Set([default__.safeModuloInt(d_2023_v18_, d_2003_cf43_)]))
                    return _dafny.Set(coll68_)
                d_2022_v19_ = _dafny.MultiSet([d_1997_v9_, iife192_()
                ])
                d_2024_v20_: D12
                d_2024_v20_ = D12_DC28(((self).f32) * ((0) - ((d_2022_v19_).cardinality)), ((self).f30)[default__.safeIndex(-352, len((self).f30))], (self).f33)
                d_2024_v20_ = d_2024_v20_
                d_2025_v21_: _dafny.Array
                nw352_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_2025_v21_ = nw352_
                d_2026_v22_: T1
                nw353_ = C1()
                nw353_.ctor__(p0, _dafny.MultiSet([p2]), ((self).f30).set(default__.safeIndex((self).f32, len((self).f30)), len(d_1990_v1_)), (0) - ((0) - (d_2003_cf43_)), (self).f28)
                d_2026_v22_ = nw353_
                d_2027_v23_: _dafny.Array
                nw354_ = _dafny.Array(None, 18)
                nw354_[int(0)] = d_2026_v22_
                nw354_[int(1)] = d_2026_v22_
                nw354_[int(2)] = d_2026_v22_
                nw354_[int(3)] = d_2026_v22_
                nw354_[int(4)] = d_2026_v22_
                nw354_[int(5)] = d_2026_v22_
                nw354_[int(6)] = d_2026_v22_
                nw354_[int(7)] = d_2026_v22_
                nw354_[int(8)] = d_2026_v22_
                nw354_[int(9)] = d_2026_v22_
                nw354_[int(10)] = d_2026_v22_
                nw354_[int(11)] = d_2026_v22_
                nw354_[int(12)] = d_2026_v22_
                nw354_[int(13)] = d_2026_v22_
                nw354_[int(14)] = d_2026_v22_
                nw354_[int(15)] = d_2026_v22_
                nw354_[int(16)] = d_2026_v22_
                nw354_[int(17)] = d_2026_v22_
                d_2027_v23_ = nw354_
                index252_ = default__.safeIndex(659, (d_2025_v21_).length(0))
                (d_2025_v21_)[index252_] = d_2027_v23_
                index253_ = default__.safeIndex(659, (d_2025_v21_).length(0))
                (d_2025_v21_)[index253_] = d_2027_v23_
                (globalState).f24 = (len((_dafny.SeqWithoutIsStrInference([d_1994_v5_ for d_2028_i2_ in range(default__.abs(32))])) + (d_1993_v4_))) + (len(d_1990_v1_))
                (globalState).f21 = default__.safeModuloInt((self).f32, (self).f27)
            elif source30_.is_DC42:
                d_2029___mcc_h8_ = source30_.cf78
                d_2030___mcc_h9_ = source30_.cf79
                d_2031___mcc_h10_ = source30_.cf80
                d_2032___mcc_h11_ = source30_.cf81
                d_2033_cf81_ = d_2032___mcc_h11_
                d_2034_cf80_ = d_2031___mcc_h10_
                d_2035_cf79_ = d_2030___mcc_h9_
                d_2036_cf78_ = d_2029___mcc_h8_
                (globalState).f3 = (self).f27
                d_2037_v24_: _dafny.Array
                nw355_ = _dafny.Array(None, 4)
                nw355_[int(0)] = d_2034_cf80_
                nw355_[int(1)] = True
                nw355_[int(2)] = (self).f33
                nw355_[int(3)] = ((self).f32) > (p1)
                d_2037_v24_ = nw355_
                index254_ = default__.safeIndex(738, (d_2037_v24_).length(0))
                (d_2037_v24_)[index254_] = p0
                index255_ = default__.safeIndex(738, (d_2037_v24_).length(0))
                (d_2037_v24_)[index255_] = (d_1990_v1_)[default__.safeIndex(len((d_1993_v4_) + (d_1993_v4_)), len(d_1990_v1_))]
                d_2038_v25_: C2
                nw356_ = C2()
                nw356_.ctor__((self).f29, (0) - ((753) * (p1)))
                d_2038_v25_ = nw356_
                index256_ = default__.safeIndex(544, (d_2004_v11_).length(0))
                (d_2004_v11_)[index256_] = d_2003_cf43_
                d_2039_v26_: T0
                nw357_ = C3()
                nw357_.ctor__((self).f27, (self).f28)
                d_2039_v26_ = nw357_
                d_2040_v27_: D16
                d_2040_v27_ = D16_DC40((d_1990_v1_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpyjyi"))), len(d_1990_v1_))], p1, (self).f27)
                d_2041_v28_: _dafny.Map
                d_2041_v28_ = _dafny.Map({d_2039_v26_: (d_2040_v27_).cf76})
                index257_ = default__.safeIndex(544, (d_2004_v11_).length(0))
                rhs374_ = d_2038_v25_
                rhs375_ = 446
                rhs376_ = (_dafny.Map({d_2039_v26_: ((self).f29).cardinality}) if (self).f33 else (d_2041_v28_ if (self).f33 else d_2041_v28_))
                lhs290_ = d_2004_v11_
                lhs291_ = default__.safeIndex(544, (d_2004_v11_).length(0))
                d_2038_v25_ = rhs374_
                lhs290_[lhs291_] = rhs375_
                d_2041_v28_ = rhs376_
                (globalState).f21 = (d_2004_v11_)[default__.safeIndex(544, (d_2004_v11_).length(0))]
            elif True:
                d_2042___mcc_h12_ = source30_.cf73
                d_2043_cf73_ = d_2042___mcc_h12_
                rhs377_ = default__.safeDivisionInt(556, p1)
                rhs378_ = (d_1990_v1_)[default__.safeIndex(len((self).f30), len(d_1990_v1_))]
                lhs292_ = globalState
                lhs293_ = globalState
                lhs292_.f15 = rhs377_
                lhs293_.f26 = rhs378_
                d_2044_v29_: _dafny.Map
                d_2044_v29_ = _dafny.Map({p0: len(_dafny.Set({p3}))})
                d_2045_v30_: _dafny.Map
                d_2045_v30_ = _dafny.Map({(self).f27: d_2044_v29_})
                rhs379_ = p2
                rhs380_ = (d_1993_v4_ if not(p3) else d_1993_v4_)
                rhs381_ = (len((d_1993_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nq"))))) + ((self).f32)
                rhs382_ = (self.f34 if p0 else (0) - (default__.safeDivisionInt(len(d_1993_v4_), 760)))
                rhs383_ = len((d_2045_v30_) | (d_2045_v30_))
                lhs294_ = globalState
                lhs295_ = globalState
                lhs296_ = globalState
                lhs297_ = globalState
                lhs294_.f26 = rhs379_
                d_1993_v4_ = rhs380_
                lhs295_.f18 = rhs381_
                lhs296_.f21 = rhs382_
                lhs297_.f21 = rhs383_
                d_2046_v31_: _dafny.Array
                nw358_ = _dafny.Array(None, 2)
                nw358_[int(0)] = d_1993_v4_
                nw358_[int(1)] = (d_1993_v4_ if p3 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yggtvlhp")))
                d_2046_v31_ = nw358_
                rhs384_ = (d_1993_v4_)[default__.safeIndex(((self).f31).cardinality, len(d_1993_v4_))]
                rhs385_ = (self).f28
                rhs386_ = default__.safeDivisionInt(len(d_1990_v1_), d_2003_cf43_)
                lhs298_ = globalState
                d_1994_v5_ = rhs384_
                d_2046_v31_ = rhs385_
                lhs298_.f24 = rhs386_
                (globalState).f0 = (self).f27
            r0 = d_1997_v9_
        elif source29_.is_DC23:
            d_2047___mcc_h2_ = source29_.cf42
            d_2048_cf42_ = d_2047___mcc_h2_
            (globalState).f26 = (self).fm3(((self).f31) | (_dafny.MultiSet([p2])), globalState)
            (globalState).f26 = (self).f33
            d_2049_v32_: D9
            d_2049_v32_ = D9_DC19(p0)
            d_2050_v33_: _dafny.MultiSet
            d_2050_v33_ = _dafny.MultiSet([d_2049_v32_])
            d_2051_v34_: C1
            nw359_ = C1()
            nw359_.ctor__(p2, (self).f29, ((self).f30).set(default__.safeIndex(((d_2050_v33_)[d_2049_v32_] if (d_2049_v32_) in (d_2050_v33_) else self.f34), len((self).f30)), (self).f27), len(default__.fm25(p2, globalState)), (self).f28)
            d_2051_v34_ = nw359_
            d_2052_v35_: _dafny.Array
            nw360_ = _dafny.Array(False, 6)
            d_2052_v35_ = nw360_
            index258_ = default__.safeIndex(105, (d_2052_v35_).length(0))
            (d_2052_v35_)[index258_] = ((self).f33) or ((self).f33)
            d_2053_v36_: bool
            d_2053_v36_ = d_2051_v34_.f35
            d_2054_v37_: _dafny.Map
            d_2054_v37_ = _dafny.Map({p2: d_1994_v5_})
            d_2055_v38_: _dafny.Map
            d_2055_v38_ = _dafny.Map({d_2054_v37_: (self).f33})
            d_2056_v39_: _dafny.Map
            d_2056_v39_ = _dafny.Map({d_2053_v36_: ((d_2055_v38_)[d_2054_v37_] if (d_2054_v37_) in (d_2055_v38_) else p3)})
            d_2057_v40_: _dafny.Map
            d_2057_v40_ = _dafny.Map({p2: d_2056_v39_})
            index259_ = default__.safeIndex(105, (d_2052_v35_).length(0))
            (d_2052_v35_)[index259_] = (((d_2057_v40_)[p2] if (p2) in (d_2057_v40_) else d_2056_v39_)) in (_dafny.SeqWithoutIsStrInference([d_2056_v39_ for d_2058_i3_ in range(default__.abs(140))]))
        elif True:
            d_2059___mcc_h3_ = source29_.cf45
            d_2060_cf45_ = d_2059___mcc_h3_
            d_2061_v41_: _dafny.MultiSet
            d_2061_v41_ = _dafny.MultiSet([284])
            d_2062_v42_: T0
            nw361_ = C5()
            nw361_.ctor__((d_1995_v6_).f38, (self).f31, (self).f30, (self).f32, (self).f28)
            d_2062_v42_ = nw361_
            d_2063_v43_: _dafny.Map
            d_2063_v43_ = _dafny.Map({d_2062_v42_: True})
            d_2064_v44_: _dafny.Array
            nw362_ = _dafny.Array(None, 20)
            nw362_[int(0)] = p3
            nw362_[int(1)] = (self).f33
            nw362_[int(2)] = p3
            nw362_[int(3)] = True
            nw362_[int(4)] = p3
            nw362_[int(5)] = p2
            nw362_[int(6)] = p0
            nw362_[int(7)] = not(p0)
            nw362_[int(8)] = p2
            nw362_[int(9)] = ((self).f33 if False else not(p0))
            nw362_[int(10)] = not (not(p0)) or (p0)
            nw362_[int(11)] = not(not(p0))
            nw362_[int(12)] = p0
            nw362_[int(13)] = not(p0)
            nw362_[int(14)] = (d_2061_v41_) == ((d_2061_v41_).set(p1, default__.abs(len(d_2063_v43_))))
            nw362_[int(15)] = p3
            nw362_[int(16)] = p2
            nw362_[int(17)] = p0
            nw362_[int(18)] = True
            nw362_[int(19)] = p0
            d_2064_v44_ = nw362_
            index260_ = default__.safeIndex(226, (d_2064_v44_).length(0))
            (d_2064_v44_)[index260_] = (d_1990_v1_) <= (d_1990_v1_)
            index261_ = default__.safeIndex(226, (d_2064_v44_).length(0))
            (d_2064_v44_)[index261_] = default__.fm0((self).f32, globalState)
            index262_ = default__.safeIndex(226, (d_2064_v44_).length(0))
            (d_2064_v44_)[index262_] = p0
            d_2065_v45_: C1
            nw363_ = C1()
            nw363_.ctor__(p0, ((self).f31).intersection((self).f29), (self).f30, (d_2062_v42_).f27, (self).f28)
            d_2065_v45_ = nw363_
            d_2066_v46_: _dafny.Set
            d_2066_v46_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogu")), (d_1993_v4_).set(default__.safeIndex(p1, len(d_1993_v4_)), d_1994_v5_)})
            rhs387_ = default__.safeDivisionInt(p1, (self).f27)
            rhs388_ = d_2065_v45_
            rhs389_ = d_2066_v46_
            lhs299_ = self
            lhs299_.f34 = rhs387_
            d_2065_v45_ = rhs388_
            d_2066_v46_ = rhs389_
            d_2067_v47_: _dafny.Array
            def lambda87_(d_2068_i4_):
                return (d_2068_i4_) + (self.f34)

            init44_ = lambda87_
            nw364_ = _dafny.Array(None, 19)
            for i0_44_ in range(nw364_.length(0)):
                nw364_[i0_44_] = init44_(i0_44_)
            d_2067_v47_ = nw364_
            d_2069_v48_: _dafny.MultiSet
            d_2069_v48_ = _dafny.MultiSet([d_2067_v47_, d_2067_v47_, d_2067_v47_, d_2067_v47_])
            d_2069_v48_ = d_2069_v48_
        d_2070_v49_: D4
        d_2070_v49_ = D4_DC8(d_1994_v5_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kybibih"))), p2)
        d_2071_v50_: D11
        d_2071_v50_ = D11_DC24((0) - ((self).f32), (d_2070_v49_ if (self).f33 else d_2070_v49_))
        source31_ = d_2071_v50_
        if source31_.is_DC24:
            d_2072___mcc_h13_ = source31_.cf43
            d_2073___mcc_h14_ = source31_.cf44
            d_2074_cf44_ = d_2073___mcc_h14_
            d_2075_cf43_ = d_2072___mcc_h13_
            d_2076_v51_: _dafny.Map
            d_2076_v51_ = _dafny.Map({p2: ((self).f30) != (_dafny.SeqWithoutIsStrInference([self.f34, d_2075_cf43_]))})
            d_2076_v51_ = (d_2076_v51_).set(p2, (self).f33)
            if p3:
                d_2077_v52_: _dafny.Array
                def lambda88_(d_2078_i5_):
                    return (d_2078_i5_) - ((self).f32)

                init45_ = lambda88_
                nw365_ = _dafny.Array(None, 4)
                for i0_45_ in range(nw365_.length(0)):
                    nw365_[i0_45_] = init45_(i0_45_)
                d_2077_v52_ = nw365_
                d_2079_v53_: T2
                nw366_ = C2()
                nw366_.ctor__((self).f29, p1)
                d_2079_v53_ = nw366_
                d_2080_v54_: _dafny.Seq
                d_2080_v54_ = _dafny.SeqWithoutIsStrInference([d_2079_v53_])
                index263_ = default__.safeIndex(820, (d_2077_v52_).length(0))
                (d_2077_v52_)[index263_] = (len(d_2080_v54_)) * ((self).f27)
                index264_ = default__.safeIndex(820, (d_2077_v52_).length(0))
                (d_2077_v52_)[index264_] = ((self).f32) * (len((d_1993_v4_) + (d_1993_v4_)))
                d_2081_v55_: _dafny.Set
                d_2081_v55_ = _dafny.Set({(self).f33, False})
                d_2082_v56_: _dafny.Array
                nw367_ = _dafny.Array(None, 10)
                nw367_[int(0)] = True
                nw367_[int(1)] = (self).f33
                nw367_[int(2)] = p2
                nw367_[int(3)] = False
                nw367_[int(4)] = (self).f33
                nw367_[int(5)] = p2
                nw367_[int(6)] = p0
                nw367_[int(7)] = p2
                nw367_[int(8)] = True
                nw367_[int(9)] = p2
                d_2082_v56_ = nw367_
                d_2083_v57_: _dafny.Array
                nw368_ = _dafny.Array(None, 8)
                nw368_[int(0)] = d_2082_v56_
                nw368_[int(1)] = d_2082_v56_
                nw368_[int(2)] = d_2082_v56_
                nw368_[int(3)] = d_2082_v56_
                nw368_[int(4)] = d_2082_v56_
                nw368_[int(5)] = d_2082_v56_
                nw368_[int(6)] = d_2082_v56_
                nw368_[int(7)] = d_2082_v56_
                d_2083_v57_ = nw368_
                d_2084_v58_: D2
                d_2084_v58_ = D2_DC4((d_2079_v53_).f32, 982, self.f34, d_1994_v5_, d_1993_v4_)
                d_2085_v59_: D16
                d_2085_v59_ = D16_DC42(p1, d_2083_v57_, p3, D2_DC5(d_2084_v58_))
                d_2086_v60_: _dafny.MultiSet
                d_2086_v60_ = _dafny.MultiSet([(399) - (len(_dafny.Map({d_2081_v55_: p3}))), self.f34, self.f34, (d_2085_v59_).cf78, (d_2079_v53_).f32])
                d_2087_v61_: _dafny.Seq
                d_2087_v61_ = _dafny.SeqWithoutIsStrInference([d_1997_v9_])
                index265_ = default__.safeIndex(820, (d_2077_v52_).length(0))
                (d_2077_v52_)[index265_] = ((d_2086_v60_)[(p1) + ((d_2077_v52_)[default__.safeIndex(820, (d_2077_v52_).length(0))])] if ((p1) + ((d_2077_v52_)[default__.safeIndex(820, (d_2077_v52_).length(0))])) in (d_2086_v60_) else len((d_2087_v61_).set(default__.safeIndex(p1, len(d_2087_v61_)), d_1997_v9_)))
                index266_ = default__.safeIndex(820, (d_2077_v52_).length(0))
                rhs390_ = p0
                rhs391_ = (d_2079_v53_).f32
                lhs300_ = globalState
                lhs301_ = d_2077_v52_
                lhs302_ = default__.safeIndex(820, (d_2077_v52_).length(0))
                lhs300_.f26 = rhs390_
                lhs301_[lhs302_] = rhs391_
                (globalState).f26 = True
                d_2088_v62_: D12
                d_2088_v62_ = D12_DC27(p0, d_2086_v60_, ((d_2086_v60_)[len(d_1997_v9_)] if (len(d_1997_v9_)) in (d_2086_v60_) else p1), d_2077_v52_)
                d_2089_v63_: _dafny.Seq
                d_2089_v63_ = _dafny.SeqWithoutIsStrInference([d_2088_v62_])
                (globalState).f24 = (len((d_2089_v63_) + (_dafny.SeqWithoutIsStrInference([d_2088_v62_])))) - (len((_dafny.Set({len((d_1993_v4_).set(default__.safeIndex(830, len(d_1993_v4_)), _dafny.CodePoint('k')))})).intersection(d_1997_v9_)))
            elif True:
                (globalState).f0 = default__.safeModuloInt(d_2075_cf43_, (self).f32)
                d_2090_v64_: _dafny.Seq
                d_2090_v64_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_2091_v65_: D9
                d_2091_v65_ = D9_DC18(p0)
                rhs392_ = default__.fm9((self).f33, globalState)
                rhs393_ = not(p3)
                rhs394_ = default__.fm52((self).f33, default__.fm0((self).f32, globalState), False, 914, globalState)
                rhs395_ = (d_1990_v1_) + (d_1990_v1_)
                rhs396_ = 844
                lhs303_ = globalState
                lhs304_ = globalState
                d_2090_v64_ = rhs392_
                lhs303_.f26 = rhs393_
                d_2091_v65_ = rhs394_
                d_1990_v1_ = rhs395_
                lhs304_.f21 = rhs396_
                (globalState).f18 = self.f34
                (globalState).f3 = self.f34
                d_2092_v66_: _dafny.Array
                def lambda89_(d_2093_cf43_, d_2094_p1_, d_2095_v0_):
                    def lambda90_(d_2096_i6_):
                        return (_dafny.Map({d_2093_cf43_: _dafny.Map({self.f34: d_2094_p1_})})) != (_dafny.Map({d_2093_cf43_: d_2095_v0_}))

                    return lambda90_

                init46_ = lambda89_(d_2075_cf43_, p1, d_1989_v0_)
                nw369_ = _dafny.Array(None, 4)
                for i0_46_ in range(nw369_.length(0)):
                    nw369_[i0_46_] = init46_(i0_46_)
                d_2092_v66_ = nw369_
                index267_ = default__.safeIndex(146, (d_2092_v66_).length(0))
                (d_2092_v66_)[index267_] = p3
                d_2097_v67_: _dafny.Map
                d_2097_v67_ = _dafny.Map({len(d_1990_v1_): d_1994_v5_})
                d_2098_v68_: _dafny.Array
                nw370_ = _dafny.Array(None, 5)
                nw370_[int(0)] = (len(d_2097_v67_)) * ((d_1996_v7_).fm45(globalState))
                nw370_[int(1)] = 638
                nw370_[int(2)] = (p1) + (d_2075_cf43_)
                nw370_[int(3)] = (d_2075_cf43_) * (p1)
                nw370_[int(4)] = d_2075_cf43_
                d_2098_v68_ = nw370_
                index268_ = default__.safeIndex(565, (d_2098_v68_).length(0))
                (d_2098_v68_)[index268_] = (self).f32
                d_2099_v69_: D2
                d_2099_v69_ = D2_DC4((self).fm7(globalState), (self).f32, d_2075_cf43_, d_1994_v5_, default__.fm37(len(d_1997_v9_), self.f34, globalState))
                d_2100_v70_: _dafny.Map
                d_2100_v70_ = _dafny.Map({d_2099_v69_: p3})
                d_2101_v71_: _dafny.Seq
                d_2101_v71_ = _dafny.SeqWithoutIsStrInference([d_2100_v70_])
                index269_ = default__.safeIndex(146, (d_2092_v66_).length(0))
                index270_ = default__.safeIndex(565, (d_2098_v68_).length(0))
                rhs397_ = ((len((D8_DC15(d_2076_v51_)).cf28)) + (p1)) <= (len(d_2076_v51_))
                rhs398_ = (_dafny.Map({default__.fm36(p3, globalState): p2})) not in (_dafny.MultiSet((d_2101_v71_) + (d_2101_v71_)))
                rhs399_ = (self).fm7(globalState)
                rhs400_ = default__.safeDivisionInt(870, d_2075_cf43_)
                lhs305_ = globalState
                lhs306_ = d_2092_v66_
                lhs307_ = default__.safeIndex(146, (d_2092_v66_).length(0))
                lhs308_ = d_2098_v68_
                lhs309_ = default__.safeIndex(565, (d_2098_v68_).length(0))
                lhs310_ = globalState
                lhs305_.f26 = rhs397_
                lhs306_[lhs307_] = rhs398_
                lhs308_[lhs309_] = rhs399_
                lhs310_.f3 = rhs400_
            d_2102_v72_: _dafny.Array
            def lambda91_(d_2103_i7_):
                return (self).f30

            init47_ = lambda91_
            nw371_ = _dafny.Array(None, 22)
            for i0_47_ in range(nw371_.length(0)):
                nw371_[i0_47_] = init47_(i0_47_)
            d_2102_v72_ = nw371_
            index271_ = default__.safeIndex(456, (d_2102_v72_).length(0))
            (d_2102_v72_)[index271_] = _dafny.SeqWithoutIsStrInference([(0) - ((0) - ((self).f32)) for d_2104_i8_ in range(default__.abs(-680))])
            d_2105_v73_: _dafny.Map
            d_2105_v73_ = _dafny.Map({-289: True})
            d_2106_v74_: _dafny.Array
            nw372_ = _dafny.Array(None, 22)
            nw372_[int(0)] = (self).f33
            nw372_[int(1)] = p0
            nw372_[int(2)] = not (not(p0)) or (p0)
            nw372_[int(3)] = (p3) == (p0)
            nw372_[int(4)] = (self).f33
            nw372_[int(5)] = p3
            nw372_[int(6)] = (p0) and (p0)
            nw372_[int(7)] = p3
            nw372_[int(8)] = p3
            nw372_[int(9)] = (_dafny.SeqWithoutIsStrInference([p1, (self).f27])) == (default__.fm16((self).fm7(globalState), globalState))
            nw372_[int(10)] = (self).f33
            nw372_[int(11)] = (len(d_2105_v73_)) < (self.f34)
            nw372_[int(12)] = ((self).f31).issubset(default__.fm32(globalState))
            nw372_[int(13)] = not(p2)
            nw372_[int(14)] = (d_1997_v9_).ispropersubset(_dafny.Set({525}))
            nw372_[int(15)] = p3
            nw372_[int(16)] = default__.fm0((self).fm7(globalState), globalState)
            nw372_[int(17)] = p3
            nw372_[int(18)] = (self).fm5(globalState)
            nw372_[int(19)] = (p1) < ((self).f32)
            nw372_[int(20)] = not(p3)
            nw372_[int(21)] = p0
            d_2106_v74_ = nw372_
            index272_ = default__.safeIndex(426, (d_2106_v74_).length(0))
            (d_2106_v74_)[index272_] = False
            index273_ = default__.safeIndex(456, (d_2102_v72_).length(0))
            index274_ = default__.safeIndex(426, (d_2106_v74_).length(0))
            rhs401_ = not (p0) or (p3)
            rhs402_ = p1
            rhs403_ = _dafny.SeqWithoutIsStrInference([self.f34])
            rhs404_ = (self).f33
            rhs405_ = ((d_1990_v1_) + (_dafny.SeqWithoutIsStrInference([p2]))) + (d_1990_v1_)
            lhs311_ = globalState
            lhs312_ = globalState
            lhs313_ = d_2102_v72_
            lhs314_ = default__.safeIndex(456, (d_2102_v72_).length(0))
            lhs315_ = d_2106_v74_
            lhs316_ = default__.safeIndex(426, (d_2106_v74_).length(0))
            lhs311_.f26 = rhs401_
            lhs312_.f24 = rhs402_
            lhs313_[lhs314_] = rhs403_
            lhs315_[lhs316_] = rhs404_
            d_1990_v1_ = rhs405_
            (globalState).f3 = d_2075_cf43_
        elif source31_.is_DC23:
            d_2107___mcc_h15_ = source31_.cf42
            d_2108_cf42_ = d_2107___mcc_h15_
            d_2109_v75_: D15
            d_2109_v75_ = D15_DC36(d_1994_v5_, (self).f32)
            source32_ = d_2109_v75_
            if source32_.is_DC36:
                d_2110___mcc_h17_ = source32_.cf64
                d_2111___mcc_h18_ = source32_.cf65
                d_2112_cf65_ = d_2111___mcc_h18_
                d_2113_cf64_ = d_2110___mcc_h17_
                d_2114_v76_: D16
                d_2114_v76_ = D16_DC40((d_1993_v4_) < ((d_1993_v4_).set(default__.safeIndex(p1, len(d_1993_v4_)), _dafny.CodePoint('g'))), self.f34, (0) - (((0) - ((self).f27) if not(p0) else (self).f27)))
                d_2115_v77_: _dafny.Array
                nw373_ = _dafny.Array(False, 21)
                d_2115_v77_ = nw373_
                d_2116_v78_: _dafny.Seq
                d_2116_v78_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({True: d_1990_v1_})])
                d_2117_v79_: _dafny.Map
                d_2117_v79_ = _dafny.Map({p1: d_1997_v9_})
                d_2118_v80_: _dafny.Map
                d_2118_v80_ = _dafny.Map({(self).f33: p0})
                pat_let_tv46_ = d_2116_v78_
                pat_let_tv47_ = d_2116_v78_
                def iife193_(_pat_let62_0):
                    def iife194_(d_2119_dt__update__tmp_h0_):
                        def iife195_(_pat_let63_0):
                            def iife196_(d_2120_dt__update_hcf75_h0_):
                                return D16_DC40((d_2119_dt__update__tmp_h0_).cf74, d_2120_dt__update_hcf75_h0_, (d_2119_dt__update__tmp_h0_).cf76)
                            return iife196_(_pat_let63_0)
                        return iife195_(len((pat_let_tv46_)[default__.safeIndex((self).f27, len(pat_let_tv47_))]))
                    return iife194_(_pat_let62_0)
                rhs406_ = iife193_(D16_DC40(p3, d_2112_cf65_, p1))
                rhs407_ = len(((d_2117_v79_)[(len(d_2118_v80_)) * (self.f34)] if ((len(d_2118_v80_)) * (self.f34)) in (d_2117_v79_) else d_1997_v9_))
                rhs408_ = (d_2115_v77_ if (self).f33 else d_2115_v77_)
                lhs317_ = globalState
                d_2114_v76_ = rhs406_
                lhs317_.f15 = rhs407_
                d_2115_v77_ = rhs408_
                (globalState).f0 = len((d_1993_v4_).set(default__.safeIndex(len(d_1993_v4_), len(d_1993_v4_)), d_1994_v5_))
                d_2121_v81_: T0
                nw374_ = C4()
                nw374_.ctor__(p1, (self).f28)
                d_2121_v81_ = nw374_
                d_2122_v82_: _dafny.Seq
                d_2122_v82_ = d_1993_v4_
                d_2123_v83_: _dafny.MultiSet
                d_2123_v83_ = _dafny.MultiSet([d_1993_v4_])
                d_2124_v84_: D7
                d_2124_v84_ = D7_DC14((self).f27, d_2121_v81_, d_2122_v82_, d_2123_v83_)
                d_2125_v85_: _dafny.Map
                d_2125_v85_ = _dafny.Map({(self).f27: d_2124_v84_})
                d_2125_v85_ = (d_2125_v85_).set(165, d_2124_v84_)
                d_2126_v86_: _dafny.Map
                d_2126_v86_ = _dafny.Map({d_1994_v5_: p1})
                d_2126_v86_ = (d_2126_v86_).set(_dafny.CodePoint('n'), (self).f27)
            elif source32_.is_DC37:
                d_2127___mcc_h19_ = source32_.cf66
                d_2128___mcc_h20_ = source32_.cf67
                d_2129___mcc_h21_ = source32_.cf68
                d_2130___mcc_h22_ = source32_.cf69
                d_2131_cf69_ = d_2130___mcc_h22_
                d_2132_cf68_ = d_2129___mcc_h21_
                d_2133_cf67_ = d_2128___mcc_h20_
                d_2134_cf66_ = d_2127___mcc_h19_
                d_2135_v87_: _dafny.Array
                nw375_ = _dafny.Array(False, 7)
                d_2135_v87_ = nw375_
                index275_ = default__.safeIndex(143, (d_2135_v87_).length(0))
                (d_2135_v87_)[index275_] = d_2131_cf69_
                index276_ = default__.safeIndex(143, (d_2135_v87_).length(0))
                (d_2135_v87_)[index276_] = (-729) != (((self).f30)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1994_v5_ for d_2136_i9_ in range(default__.abs(654))])), len((self).f30))])
                d_2137_v88_: _dafny.Set
                d_2137_v88_ = _dafny.Set({d_2134_cf66_, (d_2135_v87_)[default__.safeIndex(143, (d_2135_v87_).length(0))]})
                d_2138_v89_: _dafny.MultiSet
                d_2138_v89_ = _dafny.MultiSet([(d_1990_v1_ if d_2131_cf69_ else d_1990_v1_), d_1990_v1_, (default__.fm41(d_2133_cf67_, len(d_2137_v88_), p0, globalState)) + (d_1990_v1_)])
                d_2138_v89_ = _dafny.MultiSet((default__.fm44(globalState)) + (_dafny.SeqWithoutIsStrInference([d_1990_v1_])))
                d_2139_v90_: _dafny.Array
                nw376_ = _dafny.Array(_dafny.CodePoint('D'), 9)
                d_2139_v90_ = nw376_
                index277_ = default__.safeIndex(440, (d_2139_v90_).length(0))
                (d_2139_v90_)[index277_] = d_1994_v5_
                index278_ = default__.safeIndex(440, (d_2139_v90_).length(0))
                (d_2139_v90_)[index278_] = d_1994_v5_
                d_2140_v91_: _dafny.Seq
                d_2140_v91_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_2140_v91_ = (self).f30
            elif source32_.is_DC38:
                d_2141___mcc_h23_ = source32_.cf70
                d_2142___mcc_h24_ = source32_.cf71
                d_2143___mcc_h25_ = source32_.cf72
                d_2144_cf72_ = d_2143___mcc_h25_
                d_2145_cf71_ = d_2142___mcc_h24_
                d_2146_cf70_ = d_2141___mcc_h23_
                d_2147_v92_: _dafny.Map
                d_2147_v92_ = _dafny.Map({True: (d_1997_v9_).intersection(d_1997_v9_)})
                (globalState).f3 = len(((d_2147_v92_)[(_dafny.SeqWithoutIsStrInference([d_1994_v5_ for d_2148_i10_ in range(default__.abs(466))])) <= (d_1993_v4_)] if ((_dafny.SeqWithoutIsStrInference([d_1994_v5_ for d_2149_i10_ in range(default__.abs(466))])) <= (d_1993_v4_)) in (d_2147_v92_) else d_1997_v9_))
                d_2150_v93_: _dafny.Array
                nw377_ = _dafny.Array(_dafny.Map({}), 12)
                d_2150_v93_ = nw377_
                d_2150_v93_ = d_2150_v93_
                (globalState).f15 = p1
                d_2151_v94_: _dafny.Array
                nw378_ = _dafny.Array(None, 3)
                nw378_[int(0)] = d_2144_cf72_
                nw378_[int(1)] = d_2144_cf72_
                nw378_[int(2)] = (self).f32
                d_2151_v94_ = nw378_
                d_2152_v95_: _dafny.Map
                d_2152_v95_ = _dafny.Map({d_2151_v94_: (self).f29})
                d_2153_v96_: C5
                nw379_ = C5()
                nw379_.ctor__(_dafny.Map({d_1989_v0_: len((self).f30)}), (self).f29, ((self).f30).set(default__.safeIndex(((((d_2152_v95_)[d_2151_v94_] if (d_2151_v94_) in (d_2152_v95_) else (self).f29)).set(False, default__.abs(len(((self).f30).set(default__.safeIndex(self.f34, len((self).f30)), d_2144_cf72_))))).cardinality, len((self).f30)), self.f34), (p1 if p2 else (((self).f29)[p2] if (p2) in ((self).f29) else p1)), (self).f28)
                d_2153_v96_ = nw379_
            elif True:
                d_2154___mcc_h26_ = source32_.cf63
                d_2155_cf63_ = d_2154___mcc_h26_
                d_2156_v97_: D8
                d_2156_v97_ = D8_DC16(self.f34, (self).f27)
                d_2157_v98_: _dafny.Map
                d_2157_v98_ = _dafny.Map({d_2156_v97_: len(d_1989_v0_)})
                (globalState).f0 = len(d_2157_v98_)
                r0 = ((d_1997_v9_) | (d_1997_v9_)) | (d_1997_v9_)
                d_2158_v99_: C2
                nw380_ = C2()
                nw380_.ctor__(_dafny.MultiSet([p0]), self.f34)
                d_2158_v99_ = nw380_
                d_2159_v100_: _dafny.Map
                d_2159_v100_ = _dafny.Map({d_1990_v1_: False})
                d_2159_v100_ = ((_dafny.Map({d_1990_v1_: p0}) if p3 else d_2159_v100_)) | (default__.fm53(d_1994_v5_, globalState))
            d_2160_v101_: _dafny.Array
            nw381_ = _dafny.Array(_dafny.MultiSet({}), 4)
            d_2160_v101_ = nw381_
            index279_ = default__.safeIndex(367, (d_2160_v101_).length(0))
            (d_2160_v101_)[index279_] = (self).f31
            index280_ = default__.safeIndex(367, (d_2160_v101_).length(0))
            (d_2160_v101_)[index280_] = ((self).f29) | ((self).f29)
            (globalState).f26 = (self).f33
            d_2161_v102_: _dafny.MultiSet
            d_2161_v102_ = _dafny.MultiSet([(self).f32])
            d_2162_v103_: C2
            nw382_ = C2()
            nw382_.ctor__(_dafny.MultiSet([p0, (self).f33, (self).f33, True, p3]), (d_2161_v102_).cardinality)
            d_2162_v103_ = nw382_
        elif True:
            d_2163___mcc_h16_ = source31_.cf45
            d_2164_cf45_ = d_2163___mcc_h16_
            d_2165_v104_: C4
            nw383_ = C4()
            nw383_.ctor__(self.f34, (self).f28)
            d_2165_v104_ = nw383_
            d_2166_v106_: _dafny.MultiSet
            d_2166_v106_ = _dafny.MultiSet([(self).f27, p1, p1])
            d_2167_v107_: _dafny.Array
            nw384_ = _dafny.Array(None, 27)
            nw384_[int(0)] = not((self).f33)
            nw384_[int(1)] = p0
            nw384_[int(2)] = p0
            nw384_[int(3)] = p0
            nw384_[int(4)] = p0
            nw384_[int(5)] = (d_1990_v1_) != (d_1990_v1_)
            nw384_[int(6)] = p2
            nw384_[int(7)] = (len(d_1993_v4_)) < ((0) - ((self).f27))
            nw384_[int(8)] = not (False) or (p3)
            nw384_[int(9)] = (self.f34) <= (p1)
            nw384_[int(10)] = not ((self).f33) or (p2)
            nw384_[int(11)] = not (p0) or (p3)
            def iife197_():
                coll69_ = _dafny.Map()
                compr_69_: int
                for compr_69_ in _dafny.IntegerRange(371, 979):
                    d_2168_v105_: int = compr_69_
                    if ((371) <= (d_2168_v105_)) and ((d_2168_v105_) < (979)):
                        coll69_[(d_2168_v105_) - (566)] = (self).f32
                return _dafny.Map(coll69_)
            nw384_[int(12)] = ((self).f32) in (iife197_()
            )
            nw384_[int(13)] = (self).f33
            nw384_[int(14)] = (d_2166_v106_).issubset(d_2166_v106_)
            nw384_[int(15)] = p2
            nw384_[int(16)] = p3
            nw384_[int(17)] = not((d_1995_v6_).fm3((self).f31, globalState))
            nw384_[int(18)] = (d_1990_v1_)[default__.safeIndex((self).f32, len(d_1990_v1_))]
            nw384_[int(19)] = p0
            nw384_[int(20)] = (self).f33
            nw384_[int(21)] = p2
            nw384_[int(22)] = p0
            nw384_[int(23)] = p3
            nw384_[int(24)] = p0
            nw384_[int(25)] = p2
            nw384_[int(26)] = p3
            d_2167_v107_ = nw384_
            index281_ = default__.safeIndex(249, (d_2167_v107_).length(0))
            (d_2167_v107_)[index281_] = (self).f33
            d_2169_v108_: _dafny.Map
            d_2169_v108_ = _dafny.Map({((self).f27) <= (p1): not((d_2165_v104_).fm35(d_1994_v5_, (self).f27, globalState))})
            d_2170_v109_: _dafny.Map
            d_2170_v109_ = _dafny.Map({(self).f32: d_1993_v4_})
            index282_ = default__.safeIndex(249, (d_2167_v107_).length(0))
            rhs409_ = (d_1994_v5_) not in ((((d_2170_v109_)[self.f34] if (self.f34) in (d_2170_v109_) else _dafny.SeqWithoutIsStrInference([d_1994_v5_ for d_2171_i11_ in range(default__.abs(228))]))) + (_dafny.SeqWithoutIsStrInference([d_1994_v5_ for d_2172_i12_ in range(default__.abs(7))])))
            rhs410_ = d_2169_v108_
            lhs318_ = d_2167_v107_
            lhs319_ = default__.safeIndex(249, (d_2167_v107_).length(0))
            lhs318_[lhs319_] = rhs409_
            d_2169_v108_ = rhs410_
            d_2173_v110_: _dafny.Array
            nw385_ = _dafny.Array(int(0), 6)
            d_2173_v110_ = nw385_
            index283_ = default__.safeIndex(656, (d_2173_v110_).length(0))
            (d_2173_v110_)[index283_] = (self).f32
            index284_ = default__.safeIndex(656, (d_2173_v110_).length(0))
            (d_2173_v110_)[index284_] = default__.fm1(globalState)
            d_1993_v4_ = default__.fm37(((d_1989_v0_)[333] if (333) in (d_1989_v0_) else (self).f32), 474, globalState)
        hi12_ = (self).f27
        for d_2174_i13_ in range((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))) + (self.f34), hi12_):
            r1 = (d_2174_i13_) * (d_2174_i13_)
            d_2175_v111_: C3
            nw386_ = C3()
            nw386_.ctor__((self).f32, (self).f28)
            d_2175_v111_ = nw386_
            (globalState).f26 = not (p3) or ((self).f33)
            d_2176_v112_: _dafny.Array
            def lambda92_(d_2177_p0_):
                def lambda93_(d_2178_i14_):
                    return _dafny.Map({d_2177_p0_: True})

                return lambda93_

            init48_ = lambda92_(p0)
            nw387_ = _dafny.Array(None, 13)
            for i0_48_ in range(nw387_.length(0)):
                nw387_[i0_48_] = init48_(i0_48_)
            d_2176_v112_ = nw387_
            d_2179_v113_: _dafny.Set
            d_2179_v113_ = _dafny.Set({d_1993_v4_, default__.fm8(globalState)})
            rhs411_ = default__.safeDivisionInt(((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([831, (self).f27, (self).f27, d_2174_i13_]))).cardinality) - (len(d_1993_v4_)), p1)
            rhs412_ = default__.fm15((self).fm4(d_2179_v113_, not(p3), d_1990_v1_, globalState), d_2174_i13_, globalState)
            rhs413_ = d_2176_v112_
            rhs414_ = p1
            lhs320_ = globalState
            lhs321_ = globalState
            lhs320_.f18 = rhs411_
            d_1989_v0_ = rhs412_
            d_2176_v112_ = rhs413_
            lhs321_.f15 = rhs414_
        d_2180_v114_: _dafny.Set
        d_2180_v114_ = _dafny.Set({(self).f31})
        (globalState).f21 = ((self.f34) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgatgwca"))))) * (len(d_2180_v114_))
        d_2181_v115_: _dafny.MultiSet
        d_2181_v115_ = _dafny.MultiSet([self.f34, p1])
        d_2182_v116_: D4
        d_2182_v116_ = D4_DC9(self.f34, p3)
        d_2183_v117_: _dafny.Map
        d_2183_v117_ = _dafny.Map({(d_2181_v115_).cardinality: d_2182_v116_})
        def iife198_():
            coll70_ = _dafny.Set()
            compr_70_: int
            for compr_70_ in (d_2183_v117_).keys.Elements:
                d_2184_v118_: int = compr_70_
                if (d_2184_v118_) in (d_2183_v117_):
                    coll70_ = coll70_.union(_dafny.Set([(d_2184_v118_) * (default__.safeDivisionInt(-421, len(_dafny.Map({_dafny.CodePoint('r'): (_dafny.MultiSet([280])).cardinality}))))]))
            return _dafny.Set(coll70_)
        r0 = iife198_()
        
        r1 = 245
        nw388_ = _dafny.Array(_dafny.Array(None, 0), 15)
        r2 = nw388_
        return r0, r1, r2

    @property
    def f33(self):
        return self._f33
