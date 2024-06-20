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
        def lambda0_(source0_):
            if source0_.is_DC72:
                d_0___mcc_h0_ = source0_.cf131
                d_1___mcc_h1_ = source0_.cf132
                d_2___mcc_h2_ = source0_.cf133
                d_3_cf133_ = d_2___mcc_h2_
                d_4_cf132_ = d_1___mcc_h1_
                d_5_cf131_ = d_0___mcc_h0_
                return d_5_cf131_
            elif True:
                d_6___mcc_h3_ = source0_.cf130
                d_7_cf130_ = d_6___mcc_h3_
                if not((D1_DC5(207, False, False)).cf8):
                    return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glnsynp")))
                elif True:
                    return -309

        def iife0_():
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in (_dafny.SeqWithoutIsStrInference([(D11_DC34(True, False, 950, False, False)).cf60 for d_8_i0_ in range(default__.abs(494))])).Elements:
                    d_9_v1_: int = compr_2_
                    if (d_9_v1_) in (_dafny.SeqWithoutIsStrInference([(D11_DC34(True, False, 950, False, False)).cf60 for d_8_i0_ in range(default__.abs(494))])):
                        coll2_[default__.safeModuloInt(d_9_v1_, -948)] = False
                return _dafny.Map(coll2_)
            coll0_ = _dafny.Map()
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: int
                for compr_1_ in (_dafny.SeqWithoutIsStrInference([(D11_DC34(True, False, 950, False, False)).cf60 for d_8_i0_ in range(default__.abs(494))])).Elements:
                    d_9_v1_: int = compr_1_
                    if (d_9_v1_) in (_dafny.SeqWithoutIsStrInference([(D11_DC34(True, False, 950, False, False)).cf60 for d_8_i0_ in range(default__.abs(494))])):
                        coll1_[default__.safeModuloInt(d_9_v1_, -948)] = False
                return _dafny.Map(coll1_)
            compr_0_: int
            for compr_0_ in (iife1_()
            ).keys.Elements:
                d_10_v0_: int = compr_0_
                if (d_10_v0_) in (iife2_()
                ):
                    coll0_[default__.safeModuloInt(d_10_v0_, len(_dafny.Map({len(_dafny.Set({(_dafny.MultiSet([500])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))})): not(False)})))] = len(_dafny.SeqWithoutIsStrInference([357]))
            return _dafny.Map(coll0_)
        return (0) - (lambda0_((D27_DC72(141, not(False), len(_dafny.Set({True}))) if (D29_DC77(len(iife0_()
), True)).cf141 else D27_DC72(353, True, len(_dafny.SeqWithoutIsStrInference([not(True)]))))))

    @staticmethod
    def fm5(globalState):
        return _dafny.CodePoint('f')

    @staticmethod
    def fm7(p0, p1, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyrbgem"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjbkt"))), 417, len(_dafny.Set({False, not(False), True}))])).Elements:
                d_11_v0_: int = compr_3_
                if (d_11_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyrbgem"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjbkt"))), 417, len(_dafny.Set({False, not(False), True}))])):
                    coll3_[default__.safeDivisionInt(d_11_v0_, 156)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybcsqf"))
            return _dafny.Map(coll3_)
        return (len((_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([278, 715]))})) | (_dafny.Map({True: (_dafny.MultiSet([310, len(_dafny.Map({-538: not((D1_DC3(True)).cf4)}))])).cardinality})))) in ((iife3_()
        ) | (_dafny.Map({len(_dafny.Map({False: not(True)})): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_12_i0_ in range(default__.abs(638))])})))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        if not(True):
            return (_dafny.SeqWithoutIsStrInference([False, not(True)])) + (_dafny.SeqWithoutIsStrInference([False]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([False])

    @staticmethod
    def fm9(p0, globalState):
        return _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvuywn")): (235 if True else len(_dafny.Map({len(_dafny.Set({False, False})): len(_dafny.Map({(0) - (len(_dafny.Map({False: False}))): False}))})))})

    @staticmethod
    def fm11(p0, p1, globalState):
        if (True if True else True):
            return _dafny.CodePoint('j')
        elif True:
            return _dafny.CodePoint('i')

    @staticmethod
    def fm14(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('i')

    @staticmethod
    def fm15(p0, p1, globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ku")): _dafny.Set({True})})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxknbf")): _dafny.Set({False, True})}))

    @staticmethod
    def fm18(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oihqs"))

    @staticmethod
    def fm19(p0, p1, p2, globalState):
        if (not(False)) == (True):
            return D4_DC15(166)
        elif True:
            return D4_DC15(561)

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        if False:
            return (_dafny.SeqWithoutIsStrInference([469])) + (_dafny.SeqWithoutIsStrInference([392]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))), len(_dafny.SeqWithoutIsStrInference([not(True)]))])

    @staticmethod
    def fm21(p0, p1, globalState):
        return _dafny.CodePoint('n')

    @staticmethod
    def fm22(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(D1_DC4(True)).cf5])) + (_dafny.SeqWithoutIsStrInference([True, False, True, True]))) + ((_dafny.SeqWithoutIsStrInference([True, True])) + (_dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def fm23(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wv")) for d_13_i0_ in range(default__.abs(-729))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvir")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mskma"))])))

    @staticmethod
    def fm25(p0, p1, globalState):
        if False:
            return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_14_i0_ in range(default__.abs(-540))]))])
        elif True:
            return (_dafny.MultiSet((D4_DC13(_dafny.SeqWithoutIsStrInference([141, len(_dafny.Map({False: 110}))]))).cf29)).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([262])))

    @staticmethod
    def fm27(p0, p1, globalState):
        if (_dafny.MultiSet([len(_dafny.Set({not(True)})), (_dafny.MultiSet([False])).cardinality])).isdisjoint((D13_DC37(_dafny.MultiSet([985]))).cf67):
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(-151, 941):
                    d_15_v0_: int = compr_4_
                    if ((-151) <= (d_15_v0_)) and ((d_15_v0_) < (941)):
                        coll4_[(d_15_v0_) * (810)] = not(False)
                return _dafny.Map(coll4_)
            return D3_DC10(not(False), (0) - ((D17_DC49(len(_dafny.Map({False: False})))).cf93), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uoj")): len(iife4_()
)}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nq")), -299)
        elif True:
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(201, 625):
                    d_16_v1_: int = compr_5_
                    if ((201) <= (d_16_v1_)) and ((d_16_v1_) < (625)):
                        coll5_[default__.safeDivisionInt(d_16_v1_, 31)] = 513
                return _dafny.Map(coll5_)
            return D3_DC10(not(False), len(iife5_()
), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hw")): len(_dafny.Map({_dafny.CodePoint('x'): False}))}), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_17_i0_ in range(default__.abs(706))]), len(_dafny.Set({_dafny.SeqWithoutIsStrInference([True])})))

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return D1_DC2(_dafny.CodePoint('v'))

    @staticmethod
    def fm29(globalState):
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: int
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))})), len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, False])).cardinality, 60]))])).Elements:
                d_19_v0_: int = compr_6_
                if (d_19_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))})), len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, False])).cardinality, 60]))])):
                    coll6_ = coll6_.union(_dafny.Set([(d_19_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnkd"))))]))
            return _dafny.Set(coll6_)
        return (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_18_i0_ in range(default__.abs(203))])), 715, len(_dafny.Map({-60: len(_dafny.SeqWithoutIsStrInference([(0) - (-858)]))})), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anvvoyeki"))))})) | (iife6_()
        )

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(True if not(False) else False)])

    @staticmethod
    def fm31(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etpmvu"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nouanwc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gexyprjao"))))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return _dafny.Set({True, (True) in (_dafny.Map({True: 334})), (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnufsim"))): True})) == (_dafny.Map({(0) - (len(_dafny.Map({669: 596}))): True})), (74) == (len(_dafny.Map({len((D4_DC13(_dafny.SeqWithoutIsStrInference([617]))).cf29): 508}))), True})

    @staticmethod
    def fm33(globalState):
        return _dafny.MultiSet([(750) - ((0) - (len((D17_DC47(_dafny.Map({len(_dafny.Set({457, 967, len(_dafny.SeqWithoutIsStrInference([True, True])), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True}))]))).cardinality, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: (_dafny.MultiSet([False])).cardinality}))]))})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ac")))}))).cf89)))])

    @staticmethod
    def fm34(p0, globalState):
        return _dafny.CodePoint('s')

    @staticmethod
    def fm35(p0, p1, globalState):
        source1_ = D29_DC77((0) - (len(_dafny.Map({False: False}))), False)
        if source1_.is_DC77:
            d_20___mcc_h0_ = source1_.cf140
            d_21___mcc_h1_ = source1_.cf141
            d_22_cf141_ = d_21___mcc_h1_
            d_23_cf140_ = d_20___mcc_h0_
            return (_dafny.Map({d_22_cf141_: d_22_cf141_})) | (_dafny.Map({d_22_cf141_: d_22_cf141_}))
        elif source1_.is_DC76:
            d_24___mcc_h2_ = source1_.cf139
            d_25_cf139_ = d_24___mcc_h2_
            return _dafny.Map({False: False})
        elif True:
            d_26___mcc_h3_ = source1_.cf142
            d_27_cf142_ = d_26___mcc_h3_
            return (_dafny.Map({False: not((D1_DC4(False)).cf5)})) | (_dafny.Map({not(False): True}))

    @staticmethod
    def fm36(p0, p1, p2, p3, globalState):
        return D6_DC22(D6_DC19(_dafny.Map({False: 198})))

    @staticmethod
    def fm37(p0, globalState):
        return D10_DC30(True, (False) or (True), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkuqqt")), (_dafny.Set({(D2_DC8(True, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxsilkell")))).cf17})) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubigmjyk"))})))

    @staticmethod
    def fm38(p0, p1, p2, globalState):
        source2_ = D12_DC36(932, -817, len(_dafny.Map({D4_DC15(525): False})))
        if source2_.is_DC36:
            d_28___mcc_h0_ = source2_.cf64
            d_29___mcc_h1_ = source2_.cf65
            d_30___mcc_h2_ = source2_.cf66
            d_31_cf66_ = d_30___mcc_h2_
            d_32_cf65_ = d_29___mcc_h1_
            d_33_cf64_ = d_28___mcc_h0_
            if False:
                return _dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('o')])): False})])
            elif True:
                def iife7_():
                    coll7_ = _dafny.Map()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(642, -446):
                        d_34_v0_: int = compr_7_
                        if ((642) <= (d_34_v0_)) and ((d_34_v0_) < (-446)):
                            coll7_[default__.safeModuloInt(d_34_v0_, d_33_cf64_)] = False
                    return _dafny.Map(coll7_)
                return _dafny.SeqWithoutIsStrInference([iife7_()
                , _dafny.Map({(_dafny.MultiSet([True])).cardinality: True}), _dafny.Map({d_33_cf64_: True}), _dafny.Map({d_31_cf66_: True}), _dafny.Map({170: True})])
        elif True:
            d_35___mcc_h3_ = source2_.cf63
            d_36_cf63_ = d_35___mcc_h3_
            return _dafny.SeqWithoutIsStrInference([_dafny.Map({len((D4_DC13(_dafny.SeqWithoutIsStrInference([87]))).cf29): not(True)}), _dafny.Map({(_dafny.MultiSet([680])).cardinality: True})])

    @staticmethod
    def fm39(globalState):
        return D2_DC8((not(True)) and (True), (-665) >= (len(_dafny.Map({_dafny.Map({False: False}): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))}))), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_37_i0_ in range(default__.abs(765))]))

    @staticmethod
    def fm42(p0, p1, p2, globalState):
        return _dafny.Map({default__.safeModuloInt(138, 292): (True) in (_dafny.SeqWithoutIsStrInference([False]))})

    @staticmethod
    def fm43(p0, p1, globalState):
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(821, 478):
                d_38_v0_: int = compr_8_
                if ((821) <= (d_38_v0_)) and ((d_38_v0_) < (478)):
                    def iife9_():
                        def iife11_():
                            coll11_ = _dafny.Map()
                            compr_11_: int
                            for compr_11_ in _dafny.IntegerRange(858, -901):
                                d_39_v2_: int = compr_11_
                                if ((858) <= (d_39_v2_)) and ((d_39_v2_) < (-901)):
                                    coll11_[(d_39_v2_) + (-710)] = False
                            return _dafny.Map(coll11_)
                        coll9_ = _dafny.Map()
                        def iife10_():
                            coll10_ = _dafny.Map()
                            compr_10_: int
                            for compr_10_ in _dafny.IntegerRange(858, -901):
                                d_39_v2_: int = compr_10_
                                if ((858) <= (d_39_v2_)) and ((d_39_v2_) < (-901)):
                                    coll10_[(d_39_v2_) + (-710)] = False
                            return _dafny.Map(coll10_)
                        compr_9_: int
                        for compr_9_ in (_dafny.Map({len(iife10_()
                        ): -968})).keys.Elements:
                            d_40_v1_: int = compr_9_
                            if (d_40_v1_) in (_dafny.Map({len(iife11_()
                            ): -968})):
                                coll9_[(d_40_v1_) - (len(_dafny.SeqWithoutIsStrInference([False, False])))] = -240
                        return _dafny.Map(coll9_)
                    coll8_ = coll8_.union(_dafny.Set([default__.safeModuloInt(d_38_v0_, len(iife9_()
))]))
            return _dafny.Set(coll8_)
        return _dafny.SeqWithoutIsStrInference([(_dafny.Set({(0) - (len(_dafny.Map({False: _dafny.CodePoint('d')})))})).issubset(iife8_()
        )])

    @staticmethod
    def fm44(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([((D1_DC2(_dafny.CodePoint('m'))).cf3 if not(not(True)) else _dafny.CodePoint('o')) for d_41_i0_ in range(default__.abs(917))])

    @staticmethod
    def fm45(p0, p1, p2, globalState):
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in (_dafny.Map({818: _dafny.SeqWithoutIsStrInference([999])})).keys.Elements:
                d_42_v0_: int = compr_12_
                if (d_42_v0_) in (_dafny.Map({818: _dafny.SeqWithoutIsStrInference([999])})):
                    def iife13_():
                        coll13_ = _dafny.Set()
                        compr_13_: int
                        for compr_13_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-496])), 812])).Elements:
                            d_43_v1_: int = compr_13_
                            if (d_43_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-496])), 812])):
                                coll13_ = coll13_.union(_dafny.Set([default__.safeDivisionInt(d_43_v1_, 781)]))
                        return _dafny.Set(coll13_)
                    coll12_[default__.safeModuloInt(d_42_v0_, 81)] = (_dafny.Set({839, 801, 3, 257})) - (iife13_()
                    )
            return _dafny.Map(coll12_)
        return iife12_()
        

    @staticmethod
    def fm46(p0, p1, p2, p3, globalState):
        source3_ = D20_DC57(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbuaxrp")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gljasqy")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_44_i0_ in range(default__.abs(-683))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_45_i1_ in range(default__.abs(445))])]))
        if source3_.is_DC58:
            d_46___mcc_h0_ = source3_.cf110
            d_47_cf110_ = d_46___mcc_h0_
            return D3_DC11(712, D2_DC8(d_47_cf110_, d_47_cf110_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_48_i2_ in range(default__.abs(663))])), _dafny.Set({D0_DC0(_dafny.MultiSet([_dafny.CodePoint('q')])), D0_DC0(_dafny.MultiSet([_dafny.CodePoint('d')])), D0_DC0(_dafny.MultiSet([_dafny.CodePoint('s'), _dafny.CodePoint('e')]))}), d_47_cf110_)
        elif source3_.is_DC57:
            d_49___mcc_h1_ = source3_.cf109
            d_50_cf109_ = d_49___mcc_h1_
            return D3_DC11(168, D2_DC8(not(not(True)), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vkqorhiqs"))), _dafny.Set({D0_DC0(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('u')]))), D0_DC0(_dafny.MultiSet([_dafny.CodePoint('f'), _dafny.CodePoint('n')]))}), not(not(False)))
        elif True:
            d_51___mcc_h2_ = source3_.cf111
            d_52_cf111_ = d_51___mcc_h2_
            return D3_DC11(-374, D2_DC8(False, not(not(False)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfaootvex"))), _dafny.Set({D0_DC0(_dafny.MultiSet([_dafny.CodePoint('c'), _dafny.CodePoint('u'), _dafny.CodePoint('d'), _dafny.CodePoint('s'), _dafny.CodePoint('n')]))}), True)

    @staticmethod
    def fm47(p0, globalState):
        source4_ = D8_DC25(102)
        if source4_.is_DC25:
            d_53___mcc_h0_ = source4_.cf44
            d_54_cf44_ = d_53___mcc_h0_
            return _dafny.CodePoint('e')
        elif True:
            d_55___mcc_h1_ = source4_.cf43
            d_56_cf43_ = d_55___mcc_h1_
            return _dafny.CodePoint('c')

    @staticmethod
    def fm48(globalState):
        return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_57_i0_ in range(default__.abs(-474))])), (0) - (len((_dafny.Set({_dafny.MultiSet([True]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, not(not(False))])), _dafny.MultiSet([True])})) | (_dafny.Set({_dafny.MultiSet([False, True, True]), _dafny.MultiSet([False, True])}))))])

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        if not((D1_DC4(not(False))).cf5):
            return D1_DC2(_dafny.CodePoint('e'))
        elif True:
            return D1_DC2(_dafny.CodePoint('r'))

    @staticmethod
    def fm50(globalState):
        return ((_dafny.Map({False: -974})) | (_dafny.Map({True: -291}))) | ((_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([738]))).cardinality, (_dafny.MultiSet([True])).cardinality]))})) | (_dafny.Map({False: (0) - ((0) - (-32))})))

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        return (_dafny.Map({not(False): not(True)})) | ((_dafny.Map({True: False})) | (_dafny.Map({True: False})))

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return D6_DC22(D6_DC19(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([False, True]))})))

    @staticmethod
    def fm53(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([(_dafny.Set({True, not(True), True})) | (_dafny.Set({True})), (_dafny.Set({True, not(True), True}) if False else _dafny.Set({False, True})), _dafny.Set({not(not(True)), not(True), False})])

    @staticmethod
    def fm54(p0, globalState):
        return (((D8_DC24(_dafny.MultiSet([True, True]))).cf43).intersection(_dafny.MultiSet([True, False]))) | ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))) - (_dafny.MultiSet([True, not(not(True)), not(False), True])))

    @staticmethod
    def fm55(globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Set({False, True}) for d_58_i0_ in range(default__.abs(212))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.Set({True})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({True, True}) for d_59_i1_ in range(default__.abs(-464))])))

    @staticmethod
    def fm56(p0, globalState):
        return D0_DC0(_dafny.MultiSet([_dafny.CodePoint('h'), _dafny.CodePoint('m'), _dafny.CodePoint('t'), _dafny.CodePoint('v'), _dafny.CodePoint('t')]))

    @staticmethod
    def fm57(p0, p1, globalState):
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "judqia"))), 22})).Elements:
                d_60_v0_: int = compr_14_
                if (d_60_v0_) in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "judqia"))), 22})):
                    coll14_[(d_60_v0_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_61_i0_ in range(default__.abs(997))])))] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            return _dafny.Map(coll14_)
        return iife14_()
        

    @staticmethod
    def fm58(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(not(True))]) for d_62_i0_ in range(default__.abs(136))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]) for d_63_i1_ in range(default__.abs(834))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False]) for d_64_i2_ in range(default__.abs(784))])))

    @staticmethod
    def fm59(p0, p1, p2, globalState):
        return D4_DC14(_dafny.CodePoint('o'), (-6) + (len(_dafny.SeqWithoutIsStrInference([False]))))

    @staticmethod
    def fm60(globalState):
        def iife15_():
            coll15_ = _dafny.Map()
            compr_15_: int
            for compr_15_ in _dafny.IntegerRange(706, 284):
                d_66_v0_: int = compr_15_
                if ((706) <= (d_66_v0_)) and ((d_66_v0_) < (284)):
                    coll15_[default__.safeModuloInt(d_66_v0_, len(_dafny.Set({(D14_DC41(True, True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yurnal")))).cf73})))] = len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(not(False)): (0) - (-169)}))]))
            return _dafny.Map(coll15_)
        return ((_dafny.Map({(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdyjih"))))): -814})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_65_i0_ in range(default__.abs(14))])): len(_dafny.Map({False: (_dafny.MultiSet([not(not(False)), False, False])).cardinality}))}))) | ((_dafny.Map({888: 920})) | (iife15_()
        ))

    @staticmethod
    def fm61(p0, p1, p2, p3, globalState):
        return _dafny.Set({_dafny.Set({len(_dafny.SeqWithoutIsStrInference([813, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aebiwlg"))]))])), 864, len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kw")))])) for d_67_i0_ in range(default__.abs(42))]))}))}), _dafny.Set({500, len(_dafny.Map({252: D18_DC52(False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vol")))}))})})

    @staticmethod
    def fm62(p0, globalState):
        return (_dafny.MultiSet([_dafny.CodePoint('m')])) - ((_dafny.MultiSet([_dafny.CodePoint('x')])).intersection(_dafny.MultiSet([_dafny.CodePoint('s'), _dafny.CodePoint('v'), _dafny.CodePoint('o'), _dafny.CodePoint('f'), _dafny.CodePoint('d')])))

    @staticmethod
    def fm63(p0, p1, p2, globalState):
        source5_ = D1_DC4(not(not(not(True))))
        if source5_.is_DC3:
            d_68___mcc_h0_ = source5_.cf4
            d_69_cf4_ = d_68___mcc_h0_
            return D17_DC49(len(_dafny.Map({d_69_cf4_: 770})))
        elif source5_.is_DC4:
            d_70___mcc_h1_ = source5_.cf5
            d_71_cf5_ = d_70___mcc_h1_
            if d_71_cf5_:
                return D17_DC49(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_71_cf5_, d_71_cf5_})), 606])))
            elif True:
                return D17_DC49(1)
        elif source5_.is_DC5:
            d_72___mcc_h2_ = source5_.cf6
            d_73___mcc_h3_ = source5_.cf7
            d_74___mcc_h4_ = source5_.cf8
            d_75_cf8_ = d_74___mcc_h4_
            d_76_cf7_ = d_73___mcc_h3_
            d_77_cf6_ = d_72___mcc_h2_
            if d_76_cf7_:
                return D17_DC49(d_77_cf6_)
            elif True:
                return D17_DC49(d_77_cf6_)
        elif True:
            d_78___mcc_h5_ = source5_.cf3
            d_79_cf3_ = d_78___mcc_h5_
            return D17_DC49(348)

    @staticmethod
    def fm64(p0, globalState):
        return (_dafny.MultiSet([_dafny.Map({421: False}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prdjb"))): True})])).intersection(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.Map({794: True}) for d_80_i0_ in range(default__.abs(792))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({581: False})]))))

    @staticmethod
    def fm65(p0, p1, globalState):
        return D4_DC13(_dafny.SeqWithoutIsStrInference([124, len(_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_81_i0_ in range(default__.abs(532))])})), 359]))

    @staticmethod
    def fm66(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-881: len(_dafny.SeqWithoutIsStrInference([-753]))})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_82_i1_ in range(default__.abs(-693))]))]) for d_83_i0_ in range(default__.abs(-78))])

    @staticmethod
    def fm67(p0, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([691]))).Elements:
                d_84_v0_: int = compr_16_
                if (d_84_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([691]))):
                    coll16_[(d_84_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))))] = not(True)
            return _dafny.Map(coll16_)
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in _dafny.IntegerRange(-889, 452):
                d_85_v1_: int = compr_17_
                if ((-889) <= (d_85_v1_)) and ((d_85_v1_) < (452)):
                    coll17_[default__.safeDivisionInt(d_85_v1_, len(_dafny.SeqWithoutIsStrInference([957, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True}))]))])))] = False
            return _dafny.Map(coll17_)
        return (_dafny.Map({iife16_()
        : -355})) | ((_dafny.Map({_dafny.Map({-105: False}): -778})) | (_dafny.Map({iife17_()
        : 357})))

    @staticmethod
    def fm68(p0, p1, globalState):
        return _dafny.Set({_dafny.CodePoint('i'), _dafny.CodePoint('v')})

    @staticmethod
    def fm69(p0, p1, p2, p3, globalState):
        return (True) or (True)

    @staticmethod
    def fm70(globalState):
        return D2_DC6((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oj"))))

    @staticmethod
    def fm71(p0, p1, p2, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False, not(False)])): False})

    @staticmethod
    def fm72(p0, p1, p2, p3, globalState):
        return D20_DC59(D20_DC58(True))

    @staticmethod
    def fm73(p0, p1, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False, True, not(False), False]), _dafny.SeqWithoutIsStrInference([False, True]), _dafny.SeqWithoutIsStrInference([True, False]), _dafny.SeqWithoutIsStrInference([False, False])])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([not(True)])]))) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])])).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False])])))

    @staticmethod
    def fm74(p0, p1, p2, globalState):
        return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqff"))})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_86_i0_ in range(default__.abs(716))])}))

    @staticmethod
    def fm75(p0, globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(825, 646):
                d_87_v0_: int = compr_18_
                if ((825) <= (d_87_v0_)) and ((d_87_v0_) < (646)):
                    coll18_[default__.safeModuloInt(d_87_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))))] = False
            return _dafny.Map(coll18_)
        return D1_DC5((0) - (len(iife18_()
)), (False if True else True), False)

    @staticmethod
    def fm76(p0, globalState):
        return _dafny.Map({(D1_DC5((D29_DC77(len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejqxpdtsx"))), 426, (_dafny.MultiSet([True])).cardinality, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([281])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igemumhdf")))]))})), True)).cf140, not(False), False) if True else D1_DC5(len(_dafny.Set({True, True})), False, True)): not(True)})

    @staticmethod
    def fm77(globalState):
        return _dafny.Map({(193) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrhhgdpfh")))): _dafny.CodePoint('q')})

    @staticmethod
    def fm78(p0, p1, globalState):
        return _dafny.Map({True: _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c'), _dafny.CodePoint('q')]))).cardinality: _dafny.CodePoint('c')})})

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.Seq({})
        d_88_v0_: str
        d_88_v0_ = _dafny.CodePoint('r')
        d_89_v1_: _dafny.MultiSet
        d_89_v1_ = _dafny.MultiSet([d_88_v0_])
        d_90_v2_: D0
        d_90_v2_ = D0_DC0((d_89_v1_) - (default__.fm62(p0, globalState)))
        source6_ = d_90_v2_
        if source6_.is_DC1:
            d_91___mcc_h0_ = source6_.cf1
            d_92___mcc_h1_ = source6_.cf2
            d_93_cf2_ = d_92___mcc_h1_
            d_94_cf1_ = d_91___mcc_h0_
            d_95_v3_: _dafny.Set
            d_95_v3_ = _dafny.Set({d_94_cf1_, d_94_cf1_, p0})
            r0 = (d_95_v3_).issubset(d_95_v3_)
            d_96_v4_: _dafny.Seq
            d_96_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tg"))
            d_97_v5_: D2
            d_97_v5_ = D2_DC8(p0, p1, d_96_v4_)
            d_98_v6_: _dafny.Seq
            d_98_v6_ = _dafny.SeqWithoutIsStrInference([d_94_cf1_])
            d_99_v7_: _dafny.Array
            nw0_ = _dafny.Array(None, 21)
            nw0_[int(0)] = p0
            nw0_[int(1)] = (D3_DC11(p2, d_97_v5_, _dafny.Set({d_90_v2_}), not(d_94_cf1_))).cf27
            nw0_[int(2)] = p1
            nw0_[int(3)] = not(((_dafny.SeqWithoutIsStrInference([p0])).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference([p0]))), d_94_cf1_)) <= (d_98_v6_))
            nw0_[int(4)] = p1
            nw0_[int(5)] = p0
            nw0_[int(6)] = d_94_cf1_
            nw0_[int(7)] = p1
            nw0_[int(8)] = p0
            nw0_[int(9)] = p0
            nw0_[int(10)] = d_94_cf1_
            nw0_[int(11)] = p0
            nw0_[int(12)] = default__.fm69(True, len(d_98_v6_), d_89_v1_, p2, globalState)
            nw0_[int(13)] = (p2) != (p2)
            nw0_[int(14)] = True
            nw0_[int(15)] = not (d_94_cf1_) or (True)
            nw0_[int(16)] = (d_94_cf1_ if d_94_cf1_ else p0)
            nw0_[int(17)] = not(False)
            nw0_[int(18)] = p1
            nw0_[int(19)] = p1
            nw0_[int(20)] = p1
            d_99_v7_ = nw0_
            d_100_v8_: C1
            nw1_ = C1()
            nw1_.ctor__(d_99_v7_, d_93_cf2_, p2, p2, d_94_cf1_, d_96_v4_)
            d_100_v8_ = nw1_
            d_101_v9_: _dafny.Seq
            d_101_v9_ = _dafny.SeqWithoutIsStrInference([d_100_v8_])
            d_102_v10_: _dafny.Array
            nw2_ = _dafny.Array(None, 20)
            nw2_[int(0)] = (0) - (len(d_96_v4_))
            nw2_[int(1)] = p2
            nw2_[int(2)] = p3
            nw2_[int(3)] = p2
            nw2_[int(4)] = p3
            nw2_[int(5)] = p3
            nw2_[int(6)] = len(d_101_v9_)
            nw2_[int(7)] = p2
            nw2_[int(8)] = d_93_cf2_
            nw2_[int(9)] = d_93_cf2_
            nw2_[int(10)] = len(_dafny.Map({602: d_94_cf1_}))
            nw2_[int(11)] = p3
            nw2_[int(12)] = d_93_cf2_
            nw2_[int(13)] = 802
            nw2_[int(14)] = p3
            nw2_[int(15)] = p2
            nw2_[int(16)] = d_93_cf2_
            nw2_[int(17)] = p2
            nw2_[int(18)] = len((d_96_v4_).set(default__.safeIndex(d_93_cf2_, len(d_96_v4_)), d_88_v0_))
            nw2_[int(19)] = len(d_96_v4_)
            d_102_v10_ = nw2_
            d_103_v11_: _dafny.Map
            d_103_v11_ = _dafny.Map({p0: d_94_cf1_})
            d_104_v12_: _dafny.Seq
            d_104_v12_ = _dafny.SeqWithoutIsStrInference([d_103_v11_])
            d_105_v13_: D19
            d_105_v13_ = D19_DC56(d_102_v10_, d_94_cf1_, p2, d_104_v12_)
            index0_ = default__.safeIndex(435, (d_99_v7_).length(0))
            (d_99_v7_)[index0_] = (d_105_v13_).cf106
            index1_ = default__.safeIndex(435, (d_99_v7_).length(0))
            (d_99_v7_)[index1_] = d_94_cf1_
            if d_94_cf1_:
                d_106_v14_: _dafny.Array
                def lambda1_(d_107_v4_):
                    def lambda2_(d_108_i0_):
                        return (d_107_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ib")))

                    return lambda2_

                init0_ = lambda1_(d_96_v4_)
                nw3_ = _dafny.Array(None, 29)
                for i0_0_ in range(nw3_.length(0)):
                    nw3_[i0_0_] = init0_(i0_0_)
                d_106_v14_ = nw3_
                index2_ = default__.safeIndex(495, (d_106_v14_).length(0))
                (d_106_v14_)[index2_] = d_96_v4_
                index3_ = default__.safeIndex(495, (d_106_v14_).length(0))
                (d_106_v14_)[index3_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyhfc"))) + (_dafny.SeqWithoutIsStrInference([d_88_v0_ for d_109_i1_ in range(default__.abs(879))]))
                index4_ = default__.safeIndex(435, (d_99_v7_).length(0))
                index5_ = default__.safeIndex(495, (d_106_v14_).length(0))
                rhs0_ = d_94_cf1_
                rhs1_ = d_96_v4_
                rhs2_ = d_96_v4_
                lhs0_ = d_99_v7_
                lhs1_ = default__.safeIndex(435, (d_99_v7_).length(0))
                lhs2_ = d_106_v14_
                lhs3_ = default__.safeIndex(495, (d_106_v14_).length(0))
                lhs0_[lhs1_] = rhs0_
                d_96_v4_ = rhs1_
                lhs2_[lhs3_] = rhs2_
                d_100_v8_ = d_100_v8_
                d_110_v15_: _dafny.Map
                d_110_v15_ = _dafny.Map({False: p3})
                d_111_v16_: D10
                d_111_v16_ = D10_DC30(d_94_cf1_, default__.fm7(False, d_93_cf2_, globalState), ((d_106_v14_)[default__.safeIndex(495, (d_106_v14_).length(0))]) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_112_i2_ in range(default__.abs(701))])), (d_110_v15_) == (d_110_v15_))
                d_111_v16_ = d_111_v16_
                (globalState).f6 = d_93_cf2_
            elif True:
                d_113_v17_: C8
                nw4_ = C8()
                nw4_.ctor__(not(p0), p1)
                d_113_v17_ = nw4_
                d_93_cf2_ = default__.safeDivisionInt(default__.safeModuloInt(p2, p2), (0) - (((0) - (d_93_cf2_)) * (p2)))
                d_114_v18_: _dafny.Map
                d_114_v18_ = _dafny.Map({d_113_v17_.f21: (d_96_v4_).set(default__.safeIndex(p2, len(d_96_v4_)), d_88_v0_)})
                d_115_v19_: D28
                d_115_v19_ = D28_DC73(d_114_v18_)
                (globalState).f6 = (0) - (len((d_115_v19_).cf134))
                d_116_v20_: _dafny.Array
                def lambda3_(d_117_v6_):
                    def lambda4_(d_118_i3_):
                        return _dafny.MultiSet(d_117_v6_)

                    return lambda4_

                init1_ = lambda3_(d_98_v6_)
                nw5_ = _dafny.Array(None, 6)
                for i0_1_ in range(nw5_.length(0)):
                    nw5_[i0_1_] = init1_(i0_1_)
                d_116_v20_ = nw5_
                d_119_v21_: _dafny.Array
                nw6_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_119_v21_ = nw6_
                d_120_v22_: _dafny.Array
                nw7_ = _dafny.Array(_dafny.Map({}), 8)
                d_120_v22_ = nw7_
                index6_ = default__.safeIndex(513, (d_119_v21_).length(0))
                (d_119_v21_)[index6_] = d_120_v22_
                d_121_v23_: C0
                nw8_ = C0()
                nw8_.ctor__()
                d_121_v23_ = nw8_
                d_122_v24_: _dafny.Map
                d_122_v24_ = _dafny.Map({d_113_v17_.f21: default__.fm0((d_99_v7_)[default__.safeIndex(435, (d_99_v7_).length(0))], globalState)})
                index7_ = default__.safeIndex(513, (d_119_v21_).length(0))
                index8_ = default__.safeIndex(435, (d_99_v7_).length(0))
                rhs3_ = (d_116_v20_ if (len(d_96_v4_)) == (len(d_122_v24_)) else d_116_v20_)
                rhs4_ = d_120_v22_
                rhs5_ = (d_113_v17_).f22
                rhs6_ = (_dafny.CodePoint('n')) not in (_dafny.SeqWithoutIsStrInference([d_88_v0_ for d_123_i4_ in range(default__.abs(-692))]))
                rhs7_ = d_121_v23_
                lhs4_ = d_119_v21_
                lhs5_ = default__.safeIndex(513, (d_119_v21_).length(0))
                lhs6_ = d_99_v7_
                lhs7_ = default__.safeIndex(435, (d_99_v7_).length(0))
                d_116_v20_ = rhs3_
                lhs4_[lhs5_] = rhs4_
                lhs6_[lhs7_] = rhs5_
                d_94_cf1_ = rhs6_
                d_121_v23_ = rhs7_
                rhs8_ = True
                rhs9_ = 229
                lhs8_ = globalState
                lhs9_ = globalState
                lhs8_.f2 = rhs8_
                lhs9_.f6 = rhs9_
            d_88_v0_ = d_88_v0_
        elif True:
            d_124___mcc_h2_ = source6_.cf0
            d_125_cf0_ = d_124___mcc_h2_
            d_126_v25_: _dafny.Seq
            d_126_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yeodlan"))
            d_127_v26_: _dafny.MultiSet
            d_127_v26_ = _dafny.MultiSet([p2, p2, p2])
            d_128_v27_: _dafny.Map
            d_128_v27_ = _dafny.Map({d_126_v25_: p2})
            d_129_v28_: D3
            d_129_v28_ = D3_DC10(p0, p2, d_128_v27_, d_126_v25_, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obmxpq")))))
            d_130_v29_: C9
            nw9_ = C9()
            nw9_.ctor__((d_126_v25_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xiuooa"))), (p0) or (p1), 343, ((d_127_v26_)[(d_129_v28_).cf23] if ((d_129_v28_).cf23) in (d_127_v26_) else len(d_126_v25_)))
            d_130_v29_ = nw9_
            (globalState).f6 = p2
            d_131_v30_: _dafny.Array
            def lambda5_(d_132_v29_, d_133_p0_, d_134_p1_):
                def lambda6_(d_135_i5_):
                    return (_dafny.Set({(d_132_v29_).f20, not(d_133_p0_)})) - (_dafny.Set({d_134_p1_, (d_132_v29_).f20}))

                return lambda6_

            init2_ = lambda5_(d_130_v29_, p0, p1)
            nw10_ = _dafny.Array(None, 5)
            for i0_2_ in range(nw10_.length(0)):
                nw10_[i0_2_] = init2_(i0_2_)
            d_131_v30_ = nw10_
            d_136_v31_: _dafny.Set
            d_136_v31_ = _dafny.Set({p1})
            index9_ = default__.safeIndex(467, (d_131_v30_).length(0))
            (d_131_v30_)[index9_] = d_136_v31_
            index10_ = default__.safeIndex(467, (d_131_v30_).length(0))
            (d_131_v30_)[index10_] = ((_dafny.Set({p0})).intersection(d_136_v31_)) | (d_136_v31_)
            (globalState).f6 = (p2 if p0 else p2)
        d_137_v32_: _dafny.Map
        d_137_v32_ = _dafny.Map({p2: True})
        d_138_v33_: _dafny.Seq
        d_138_v33_ = _dafny.SeqWithoutIsStrInference([d_137_v32_, _dafny.Map({p3: p0})])
        d_139_v34_: _dafny.Map
        d_139_v34_ = _dafny.Map({len((d_138_v33_)[default__.safeIndex(p2, len(d_138_v33_))]): p1})
        (globalState).f7 = default__.fm69(p1, (0) - (p3), d_89_v1_, default__.safeDivisionInt(len(d_139_v34_), p2), globalState)
        d_140_v35_: _dafny.Array
        nw11_ = _dafny.Array(False, 25)
        d_140_v35_ = nw11_
        index11_ = default__.safeIndex(651, (d_140_v35_).length(0))
        (d_140_v35_)[index11_] = (p2) <= (p2)
        index12_ = default__.safeIndex(651, (d_140_v35_).length(0))
        rhs10_ = (p3) == (default__.fm0(p0, globalState))
        rhs11_ = p0
        lhs10_ = globalState
        lhs11_ = d_140_v35_
        lhs12_ = default__.safeIndex(651, (d_140_v35_).length(0))
        lhs10_.f7 = rhs10_
        lhs11_[lhs12_] = rhs11_
        d_141_v36_: _dafny.Set
        d_141_v36_ = _dafny.Set({p0, (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))]})
        d_142_v37_: _dafny.Set
        d_142_v37_ = _dafny.Set({(0) - (p3)})
        d_143_v38_: _dafny.Seq
        d_143_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prushd"))
        d_144_v39_: T1
        nw12_ = C4()
        nw12_.ctor__(d_143_v38_, 430, p0, d_143_v38_, p2, p2)
        d_144_v39_ = nw12_
        d_145_v40_: _dafny.Map
        d_145_v40_ = _dafny.Map({p2: d_144_v39_})
        d_146_v41_: _dafny.Map
        d_146_v41_ = _dafny.Map({p1: p3})
        d_147_v42_: _dafny.Seq
        d_147_v42_ = _dafny.SeqWithoutIsStrInference([len(d_146_v41_), (d_144_v39_).f12])
        d_148_v43_: C2
        nw13_ = C2()
        nw13_.ctor__(len(d_145_v40_), d_143_v38_, p3, (0) - (len(d_147_v42_)))
        d_148_v43_ = nw13_
        d_149_v44_: D18
        d_149_v44_ = D18_DC53(520, d_142_v37_, p2, d_148_v43_)
        source7_ = (D18_DC53(len(d_141_v36_), d_142_v37_, p2, d_148_v43_) if (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))] else d_149_v44_)
        if source7_.is_DC52:
            d_150___mcc_h3_ = source7_.cf96
            d_151___mcc_h4_ = source7_.cf97
            d_152___mcc_h5_ = source7_.cf98
            d_153_cf98_ = d_152___mcc_h5_
            d_154_cf97_ = d_151___mcc_h4_
            d_155_cf96_ = d_150___mcc_h3_
            d_88_v0_ = d_88_v0_
            d_156_v45_: D4
            d_156_v45_ = D4_DC15(len(_dafny.Map({d_148_v43_.f27: (0) - (len(_dafny.SeqWithoutIsStrInference([d_88_v0_ for d_157_i6_ in range(default__.abs(488))])))})))
            d_158_v46_: D4
            d_158_v46_ = D4_DC16(d_156_v45_)
            d_159_v47_: D4
            d_159_v47_ = D4_DC16(d_158_v46_)
            d_160_v48_: _dafny.Map
            d_160_v48_ = _dafny.Map({d_155_cf96_: p0})
            d_161_v49_: _dafny.Seq
            d_161_v49_ = _dafny.SeqWithoutIsStrInference([d_155_cf96_])
            rhs12_ = ((d_160_v48_).set((d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))], d_155_cf96_)) | (d_160_v48_)
            rhs13_ = (d_161_v49_) + ((d_161_v49_ if d_155_cf96_ else default__.fm8(p0, (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))], p1, len(d_153_cf98_), globalState)))
            rhs14_ = d_159_v47_
            lhs13_ = globalState
            lhs13_.f1 = rhs12_
            r1 = rhs13_
            d_159_v47_ = rhs14_
            d_162_v50_: _dafny.Array
            def lambda7_(d_163_v39_):
                def lambda8_(d_164_i7_):
                    return (d_163_v39_).f13

                return lambda8_

            init3_ = lambda7_(d_144_v39_)
            nw14_ = _dafny.Array(None, 10)
            for i0_3_ in range(nw14_.length(0)):
                nw14_[i0_3_] = init3_(i0_3_)
            d_162_v50_ = nw14_
            index13_ = default__.safeIndex(470, (d_162_v50_).length(0))
            (d_162_v50_)[index13_] = d_143_v38_
            index14_ = default__.safeIndex(470, (d_162_v50_).length(0))
            (d_162_v50_)[index14_] = d_143_v38_
            index15_ = default__.safeIndex(470, (d_162_v50_).length(0))
            (d_162_v50_)[index15_] = (d_153_cf98_).set(default__.safeIndex(p2, len(d_153_cf98_)), default__.fm11(d_155_cf96_, p3, globalState))
        elif source7_.is_DC53:
            d_165___mcc_h6_ = source7_.cf99
            d_166___mcc_h7_ = source7_.cf100
            d_167___mcc_h8_ = source7_.cf101
            d_168___mcc_h9_ = source7_.cf102
            d_169_cf102_ = d_168___mcc_h9_
            d_170_cf101_ = d_167___mcc_h8_
            d_171_cf100_ = d_166___mcc_h7_
            d_172_cf99_ = d_165___mcc_h6_
            if not((d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))]):
                (globalState).f2 = True
                d_173_v51_: T0
                nw15_ = C1()
                nw15_.ctor__(d_140_v35_, (0) - ((len(_dafny.Set({d_148_v43_.f27, (d_144_v39_).f11}))) * (d_169_cf102_.f27)), (0) - ((0) - (default__.fm0(p1, globalState))), len(_dafny.SeqWithoutIsStrInference([p2 for d_174_i8_ in range(default__.abs(-158))])), (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))], (d_144_v39_).f13)
                d_173_v51_ = nw15_
                d_173_v51_ = d_173_v51_
                d_175_v52_: _dafny.Array
                nw16_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_175_v52_ = nw16_
                d_176_v53_: _dafny.Array
                nw17_ = _dafny.Array(None, 7)
                nw17_[int(0)] = (d_144_v39_).f11
                nw17_[int(1)] = (d_173_v51_).f12
                nw17_[int(2)] = p2
                nw17_[int(3)] = default__.fm0(not(p1), globalState)
                nw17_[int(4)] = (d_144_v39_).f12
                nw17_[int(5)] = d_170_cf101_
                nw17_[int(6)] = d_170_cf101_
                d_176_v53_ = nw17_
                index16_ = default__.safeIndex(929, (d_175_v52_).length(0))
                (d_175_v52_)[index16_] = d_176_v53_
                d_177_v54_: _dafny.Map
                d_177_v54_ = _dafny.Map({d_148_v43_: (d_144_v39_).f12})
                d_178_v55_: _dafny.Seq
                d_178_v55_ = _dafny.SeqWithoutIsStrInference([(d_144_v39_).f13, _dafny.SeqWithoutIsStrInference([d_88_v0_ for d_179_i9_ in range(default__.abs(635))])])
                index17_ = default__.safeIndex(929, (d_175_v52_).length(0))
                nw18_ = _dafny.Array(None, 17)
                nw18_[int(0)] = d_169_cf102_.f27
                nw18_[int(1)] = (d_144_v39_).f11
                nw18_[int(2)] = d_172_cf99_
                nw18_[int(3)] = 554
                nw18_[int(4)] = (d_144_v39_).f12
                nw18_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oruuaf")))
                nw18_[int(6)] = ((d_173_v51_).f12) - ((_dafny.MultiSet([(d_173_v51_).f12])).cardinality)
                nw18_[int(7)] = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))).set(default__.safeIndex((d_144_v39_).f11, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))), d_88_v0_)) + (d_143_v38_))
                nw18_[int(8)] = ((0) - (len((d_144_v39_).f13))) - (d_170_cf101_)
                nw18_[int(9)] = ((d_146_v41_)[not((d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))])] if (not((d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))])) in (d_146_v41_) else d_148_v43_.f27)
                nw18_[int(10)] = -472
                nw18_[int(11)] = len(d_147_v42_)
                nw18_[int(12)] = ((d_177_v54_)[d_148_v43_] if (d_148_v43_) in (d_177_v54_) else (d_144_v39_).f12)
                nw18_[int(13)] = (d_144_v39_).f12
                nw18_[int(14)] = (d_144_v39_).f12
                nw18_[int(15)] = p3
                nw18_[int(16)] = len(d_178_v55_)
                (d_175_v52_)[index17_] = nw18_
                d_172_cf99_ = (0) - (default__.safeDivisionInt((d_148_v43_).fm40((d_144_v39_).f12, not(p1), d_88_v0_, (d_169_cf102_).fm41(globalState), globalState), ((d_144_v39_).f11 if p1 else 19)))
                d_180_v56_: _dafny.MultiSet
                d_180_v56_ = _dafny.MultiSet([p0, p1, not(p1)])
                d_181_v57_: _dafny.Seq
                d_181_v57_ = _dafny.SeqWithoutIsStrInference([p1, (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))]])
                d_182_v58_: _dafny.Array
                nw19_ = _dafny.Array(None, 3)
                nw19_[int(0)] = (d_180_v56_).intersection(d_180_v56_)
                nw19_[int(1)] = (d_180_v56_).intersection(d_180_v56_)
                nw19_[int(2)] = _dafny.MultiSet(d_181_v57_)
                d_182_v58_ = nw19_
                index18_ = default__.safeIndex(535, (d_182_v58_).length(0))
                (d_182_v58_)[index18_] = (d_180_v56_ if (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))] else d_180_v56_)
                index19_ = default__.safeIndex(535, (d_182_v58_).length(0))
                (d_182_v58_)[index19_] = d_180_v56_
            elif True:
                d_183_v59_: _dafny.MultiSet
                d_183_v59_ = _dafny.MultiSet([(d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))], (D22_DC63(d_148_v43_.f27, p0, (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))])).cf119, (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))], p0])
                (globalState).f2 = (d_183_v59_) != (d_183_v59_)
                d_184_v60_: T2
                nw20_ = C4()
                nw20_.ctor__((d_144_v39_).f13, ((d_144_v39_).f11) + (len(d_146_v41_)), (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))], (d_143_v38_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uejn"))), (0) - (p3), 720)
                d_184_v60_ = nw20_
                d_185_v61_: _dafny.MultiSet
                d_185_v61_ = _dafny.MultiSet([d_140_v35_])
                rhs15_ = _dafny.CodePoint('s')
                rhs16_ = d_137_v32_
                rhs17_ = (d_148_v43_.f27) + ((p2) - (((d_183_v59_)[not((d_184_v60_).f15)] if (not((d_184_v60_).f15)) in (d_183_v59_) else (d_185_v61_).cardinality)))
                rhs18_ = d_184_v60_
                lhs14_ = d_148_v43_
                d_88_v0_ = rhs15_
                d_137_v32_ = rhs16_
                lhs14_.f27 = rhs17_
                d_184_v60_ = rhs18_
                d_148_v43_ = d_169_cf102_
                d_186_v62_: _dafny.Map
                d_186_v62_ = _dafny.Map({False: not(p0)})
                d_187_v63_: _dafny.Map
                d_187_v63_ = _dafny.Map({(d_184_v60_).f14: d_186_v62_})
                d_188_v65_: _dafny.Map
                d_188_v65_ = _dafny.Map({d_88_v0_: (d_144_v39_).f13})
                d_189_v66_: _dafny.Map
                def iife19_():
                    coll19_ = _dafny.Map()
                    compr_19_: str
                    for compr_19_ in (d_188_v65_).keys.Elements:
                        d_190_v64_: str = compr_19_
                        if (d_190_v64_) in (d_188_v65_):
                            coll19_[d_190_v64_] = (d_184_v60_).f12
                    return _dafny.Map(coll19_)
                d_189_v66_ = _dafny.Map({D25_DC67(((d_187_v63_)[len((d_144_v39_).f13)] if (len((d_144_v39_).f13)) in (d_187_v63_) else default__.fm35((d_184_v60_).f15, d_147_v42_, globalState))): iife19_()
                })
                d_191_v67_: D25
                d_191_v67_ = D25_DC67(default__.fm35(p0, d_147_v42_, globalState))
                d_192_v68_: _dafny.Map
                d_192_v68_ = _dafny.Map({d_88_v0_: d_148_v43_.f27})
                d_189_v66_ = (d_189_v66_).set(d_191_v67_, (d_192_v68_) | (_dafny.Map({default__.fm47(p3, globalState): p3})))
                (d_148_v43_).f27 = 846
            if False:
                d_193_v69_: D3
                d_193_v69_ = D3_DC10((d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))], d_170_cf101_, _dafny.Map({(d_144_v39_).f13: (d_144_v39_).f11}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqm")), (d_144_v39_).f11)
                d_194_v70_: C2
                nw21_ = C2()
                nw21_.ctor__(d_172_cf99_, ((d_193_v69_).cf22) + ((d_144_v39_).f13), (d_144_v39_).f12, (0) - ((p2) * ((d_144_v39_).f12)))
                d_194_v70_ = nw21_
                d_195_v71_: D12
                d_195_v71_ = D12_DC35(d_141_v36_)
                d_196_v72_: _dafny.MultiSet
                d_196_v72_ = _dafny.MultiSet([d_195_v71_, D12_DC35(d_141_v36_), d_195_v71_])
                d_197_v73_: _dafny.Map
                d_197_v73_ = _dafny.Map({p1: d_196_v72_})
                d_198_v74_: D29
                d_198_v74_ = D29_DC76(d_196_v72_)
                d_199_v75_: C3
                nw22_ = C3()
                nw22_.ctor__(p1, _dafny.SeqWithoutIsStrInference([d_88_v0_ for d_200_i10_ in range(default__.abs(260))]), (d_144_v39_).f13, len((d_197_v73_).set(p0, (d_198_v74_).cf139)), p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsnwettne")), d_148_v43_.f27, -834)
                d_199_v75_ = nw22_
                d_201_v76_: D24
                d_201_v76_ = D24_DC65(d_199_v75_)
                d_201_v76_ = d_201_v76_
                d_148_v43_ = d_148_v43_
                d_202_v77_: _dafny.Array
                def lambda9_(d_203_v39_):
                    def lambda10_(d_204_i11_):
                        return (d_204_i11_) + ((d_203_v39_).f11)

                    return lambda10_

                init4_ = lambda9_(d_144_v39_)
                nw23_ = _dafny.Array(None, 28)
                for i0_4_ in range(nw23_.length(0)):
                    nw23_[i0_4_] = init4_(i0_4_)
                d_202_v77_ = nw23_
                d_205_v78_: _dafny.Map
                d_205_v78_ = _dafny.Map({False: d_202_v77_})
                d_205_v78_ = (d_205_v78_).set(p1, d_202_v77_)
                d_206_v79_: C10
                nw24_ = C10()
                nw24_.ctor__(((d_144_v39_).f13) + (d_143_v38_), ((d_199_v75_).f25 if (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))] else default__.fm7(False, p3, globalState)), d_143_v38_, ((d_144_v39_).f12) - (len(_dafny.SeqWithoutIsStrInference([d_172_cf99_, d_148_v43_.f27]))), p0, (_dafny.SeqWithoutIsStrInference([d_88_v0_ for d_207_i12_ in range(default__.abs(643))])) + ((d_144_v39_).f13), d_148_v43_.f27, ((d_144_v39_).f12) - ((0) - (p2)))
                d_206_v79_ = nw24_
            elif True:
                d_208_v80_: _dafny.Map
                d_208_v80_ = _dafny.Map({not(p1): (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))]})
                index20_ = default__.safeIndex(651, (d_140_v35_).length(0))
                (d_140_v35_)[index20_] = ((d_208_v80_)[(p0 if p1 else p0)] if ((p0 if p1 else p0)) in (d_208_v80_) else not(True))
                d_138_v33_ = (_dafny.SeqWithoutIsStrInference([d_139_v34_])) + (_dafny.SeqWithoutIsStrInference([(d_139_v34_).set(p2, p1) for d_209_i13_ in range(default__.abs(775))]))
                d_137_v32_ = (d_137_v32_).set((d_144_v39_).f12, (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))])
                r0 = p1
                d_210_v81_: _dafny.Seq
                d_210_v81_ = _dafny.SeqWithoutIsStrInference([not((p0) and (p0)), p0, p0, (p0) or (p0), (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))]])
                index21_ = default__.safeIndex(651, (d_140_v35_).length(0))
                (d_140_v35_)[index21_] = (d_210_v81_)[default__.safeIndex((d_144_v39_).f11, len(d_210_v81_))]
            d_211_v82_: _dafny.Array
            nw25_ = _dafny.Array(int(0), 24)
            d_211_v82_ = nw25_
            d_212_v83_: _dafny.Map
            d_212_v83_ = _dafny.Map({p3: d_211_v82_})
            d_213_v84_: _dafny.Seq
            d_213_v84_ = _dafny.SeqWithoutIsStrInference([d_211_v82_])
            d_214_v85_: _dafny.Array
            nw26_ = _dafny.Array(None, 11)
            nw26_[int(0)] = d_211_v82_
            nw26_[int(1)] = d_211_v82_
            nw26_[int(2)] = d_211_v82_
            nw26_[int(3)] = d_211_v82_
            nw26_[int(4)] = d_211_v82_
            nw26_[int(5)] = d_211_v82_
            nw26_[int(6)] = ((d_212_v83_)[d_172_cf99_] if (d_172_cf99_) in (d_212_v83_) else d_211_v82_)
            nw26_[int(7)] = d_211_v82_
            nw26_[int(8)] = d_211_v82_
            nw26_[int(9)] = d_211_v82_
            nw26_[int(10)] = (d_213_v84_)[default__.safeIndex(p2, len(d_213_v84_))]
            d_214_v85_ = nw26_
            d_215_v86_: _dafny.Map
            d_215_v86_ = _dafny.Map({(d_144_v39_).f12: (d_144_v39_).f11})
            index22_ = default__.safeIndex(342, (d_214_v85_).length(0))
            (d_214_v85_)[index22_] = (d_213_v84_)[default__.safeIndex(len((d_215_v86_).set(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_216_i14_ in range(default__.abs(446))])), d_172_cf99_)), len(d_213_v84_))]
            index23_ = default__.safeIndex(342, (d_214_v85_).length(0))
            (d_214_v85_)[index23_] = d_211_v82_
            d_217_v87_: _dafny.Seq
            d_217_v87_ = _dafny.SeqWithoutIsStrInference([default__.fm69(p0, (d_144_v39_).f11, d_89_v1_, d_148_v43_.f27, globalState)])
            source8_ = D15_DC43(len((d_217_v87_) + (d_217_v87_)), p0, ((d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))] if p0 else p1))
            if source8_.is_DC43:
                d_218___mcc_h12_ = source8_.cf77
                d_219___mcc_h13_ = source8_.cf78
                d_220___mcc_h14_ = source8_.cf79
                d_221_cf79_ = d_220___mcc_h14_
                d_222_cf78_ = d_219___mcc_h13_
                d_223_cf77_ = d_218___mcc_h12_
                (d_148_v43_).f27 = ((d_144_v39_).f11) * (d_172_cf99_)
                d_224_v88_: _dafny.Map
                d_224_v88_ = _dafny.Map({d_140_v35_: p1})
                (d_169_cf102_).f27 = (len(_dafny.Map({(d_144_v39_).f11: ((d_224_v88_)[d_140_v35_] if (d_140_v35_) in (d_224_v88_) else d_221_cf79_)})) if p0 else (0) - ((d_144_v39_).f12))
                d_137_v32_ = (d_137_v32_).set(default__.safeModuloInt(d_169_cf102_.f27, d_148_v43_.f27), (p3) not in (_dafny.Map({d_170_cf101_: d_140_v35_})))
                d_225_v89_: _dafny.Array
                nw27_ = _dafny.Array(None, 6)
                nw27_[int(0)] = d_140_v35_
                nw27_[int(1)] = d_140_v35_
                nw27_[int(2)] = d_140_v35_
                nw27_[int(3)] = d_140_v35_
                nw27_[int(4)] = d_140_v35_
                nw27_[int(5)] = d_140_v35_
                d_225_v89_ = nw27_
                d_225_v89_ = d_225_v89_
            elif True:
                d_226___mcc_h15_ = source8_.cf76
                d_227_cf76_ = d_226___mcc_h15_
                (globalState).f7 = p0
                d_228_v90_: _dafny.Array
                d_228_v90_ = d_211_v82_
                (d_148_v43_).m19(((d_144_v39_).f11) * (p3), p3, d_228_v90_, d_169_cf102_.f27, globalState)
                d_229_v91_: _dafny.Array
                def lambda11_(d_230_v42_):
                    def lambda12_(d_231_i15_):
                        return d_230_v42_

                    return lambda12_

                init5_ = lambda11_(d_147_v42_)
                nw28_ = _dafny.Array(None, 27)
                for i0_5_ in range(nw28_.length(0)):
                    nw28_[i0_5_] = init5_(i0_5_)
                d_229_v91_ = nw28_
                d_232_v92_: D11
                d_232_v92_ = D11_DC34(p0, (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))], 668, False, (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))])
                index24_ = default__.safeIndex(78, (d_229_v91_).length(0))
                (d_229_v91_)[index24_] = _dafny.SeqWithoutIsStrInference([d_148_v43_.f27, p2, 391, p2, default__.fm0((d_232_v92_).cf58, globalState)])
                d_233_v93_: _dafny.Map
                d_233_v93_ = _dafny.Map({(d_144_v39_).f11: d_147_v42_})
                index25_ = default__.safeIndex(78, (d_229_v91_).length(0))
                index26_ = default__.safeIndex(651, (d_140_v35_).length(0))
                rhs19_ = ((d_233_v93_)[d_148_v43_.f27] if (d_148_v43_.f27) in (d_233_v93_) else ((_dafny.SeqWithoutIsStrInference([(d_144_v39_).f12 for d_234_i16_ in range(default__.abs(-990))])).set(default__.safeIndex(d_169_cf102_.f27, len(_dafny.SeqWithoutIsStrInference([(d_144_v39_).f12 for d_235_i16_ in range(default__.abs(-990))]))), p2)) + (d_147_v42_))
                rhs20_ = p1
                rhs21_ = not (not((d_144_v39_).fm1((d_144_v39_).f12, p1, True, globalState))) or ((d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))])
                rhs22_ = (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))]
                lhs15_ = d_229_v91_
                lhs16_ = default__.safeIndex(78, (d_229_v91_).length(0))
                lhs17_ = globalState
                lhs18_ = d_140_v35_
                lhs19_ = default__.safeIndex(651, (d_140_v35_).length(0))
                lhs15_[lhs16_] = rhs19_
                lhs17_.f2 = rhs20_
                r0 = rhs21_
                lhs18_[lhs19_] = rhs22_
                d_143_v38_ = (d_144_v39_).f13
        elif source7_.is_DC51:
            d_236___mcc_h10_ = source7_.cf95
            d_237_cf95_ = d_236___mcc_h10_
            d_143_v38_ = ((d_144_v39_).f13) + ((d_144_v39_).f13)
            d_238_v94_: _dafny.Seq
            d_238_v94_ = _dafny.SeqWithoutIsStrInference([d_143_v38_, (d_144_v39_).f13, (d_144_v39_).f13])
            d_239_v95_: _dafny.MultiSet
            d_239_v95_ = _dafny.MultiSet([p1, (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))]])
            d_238_v94_ = default__.fm23(((d_239_v95_)[p0] if (p0) in (d_239_v95_) else (d_144_v39_).f11), (d_144_v39_).f11, globalState)
            d_88_v0_ = default__.fm47(d_148_v43_.f27, globalState)
            d_240_v96_: C3
            nw29_ = C3()
            nw29_.ctor__(p0, d_143_v38_, ((d_144_v39_).f13) + (d_143_v38_), (d_148_v43_.f27) * (p2), (p0) or ((d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))]), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ceiapmhef"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "estmoun"))), (d_144_v39_).f11, (d_144_v39_).f11)
            d_240_v96_ = nw29_
        elif True:
            d_241___mcc_h11_ = source7_.cf103
            d_242_cf103_ = d_241___mcc_h11_
            (globalState).f2 = p0
            d_243_v97_: _dafny.Array
            nw30_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
            d_243_v97_ = nw30_
            index27_ = default__.safeIndex(328, (d_243_v97_).length(0))
            (d_243_v97_)[index27_] = (d_144_v39_).f13
            d_244_v98_: C4
            nw31_ = C4()
            nw31_.ctor__((d_144_v39_).f13, d_148_v43_.f27, p1, (d_144_v39_).f13, (d_144_v39_).f12, len(d_137_v32_))
            d_244_v98_ = nw31_
            d_245_v99_: _dafny.Map
            d_245_v99_ = _dafny.Map({(d_244_v98_ if (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))] else d_244_v98_): ((d_144_v39_).f13) + ((d_144_v39_).f13)})
            index28_ = default__.safeIndex(328, (d_243_v97_).length(0))
            (d_243_v97_)[index28_] = ((d_245_v99_)[d_244_v98_] if (d_244_v98_) in (d_245_v99_) else _dafny.SeqWithoutIsStrInference([d_88_v0_ for d_246_i17_ in range(default__.abs(-64))]))
            d_88_v0_ = d_88_v0_
            (globalState).f7 = not (not(not(p0))) or ((d_148_v43_).fm41(globalState))
        d_247_i18_: int
        d_247_i18_ = 0
        with _dafny.label("0"):
            while not(((d_144_v39_).f12) >= (d_148_v43_.f27)):
                with _dafny.c_label("0"):
                    if (d_247_i18_) >= (100):
                        raise _dafny.Break("0")
                    d_247_i18_ = (d_247_i18_) + (1)
                    (d_148_v43_).f27 = (d_144_v39_).f12
                    index29_ = default__.safeIndex(651, (d_140_v35_).length(0))
                    (d_140_v35_)[index29_] = (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))]
                    (globalState).f7 = (d_140_v35_)[default__.safeIndex(651, (d_140_v35_).length(0))]
                    index30_ = default__.safeIndex(651, (d_140_v35_).length(0))
                    (d_140_v35_)[index30_] = p0
                    pass
            pass
        index31_ = default__.safeIndex(651, (d_140_v35_).length(0))
        (d_140_v35_)[index31_] = (d_148_v43_).fm1(((d_144_v39_).f11) * (p2), p0, p1, globalState)
        r0 = not(True)
        d_248_v100_: _dafny.Seq
        d_248_v100_ = _dafny.SeqWithoutIsStrInference([not(p1)])
        r1 = d_248_v100_
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_249_v0_: bool
        d_249_v0_ = False
        d_250_v1_: _dafny.Seq
        d_250_v1_ = _dafny.SeqWithoutIsStrInference([d_249_v0_])
        d_251_v2_: int
        d_251_v2_ = 67
        d_252_v3_: _dafny.Map
        d_252_v3_ = _dafny.Map({d_249_v0_: d_249_v0_})
        d_253_v4_: _dafny.Array
        nw32_ = _dafny.Array(None, 6)
        nw32_[int(0)] = d_249_v0_
        nw32_[int(1)] = d_249_v0_
        nw32_[int(2)] = d_249_v0_
        nw32_[int(3)] = d_249_v0_
        nw32_[int(4)] = (d_250_v1_)[default__.safeIndex(d_251_v2_, len(d_250_v1_))]
        nw32_[int(5)] = d_249_v0_
        d_253_v4_ = nw32_
        d_254_v5_: _dafny.Seq
        d_254_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlvjuryaj"))
        d_255_globalState_: GlobalState
        nw33_ = GlobalState()
        nw33_.ctor__(((d_250_v1_).set(default__.safeIndex(d_251_v2_, len(d_250_v1_)), d_249_v0_)) + (d_250_v1_), d_252_v3_, False, d_253_v4_, -370, _dafny.CodePoint('m'), 250, True, 750, 972, d_254_v5_)
        d_255_globalState_ = nw33_
        d_256_v6_: _dafny.Seq
        d_256_v6_ = _dafny.SeqWithoutIsStrInference([d_251_v2_, len(_dafny.SeqWithoutIsStrInference([d_251_v2_]))])
        (d_255_globalState_).f6 = (d_256_v6_)[default__.safeIndex(d_251_v2_, len(d_256_v6_))]
        d_257_v7_: _dafny.MultiSet
        d_257_v7_ = _dafny.MultiSet([d_251_v2_, d_251_v2_])
        d_258_v8_: bool
        d_259_v9_: _dafny.Seq
        out0_: bool
        out1_: _dafny.Seq
        out0_, out1_ = default__.m0((d_251_v2_) in (d_257_v7_), d_249_v0_, d_251_v2_, default__.fm0(d_249_v0_, d_255_globalState_), d_255_globalState_)
        d_258_v8_ = out0_
        d_259_v9_ = out1_
        d_260_v10_: str
        d_260_v10_ = _dafny.CodePoint('j')
        d_261_v11_: _dafny.MultiSet
        d_261_v11_ = _dafny.MultiSet([d_260_v10_, d_260_v10_, d_260_v10_, d_260_v10_, d_260_v10_])
        d_262_v12_: _dafny.Array
        nw34_ = _dafny.Array(None, 4)
        nw34_[int(0)] = d_261_v11_
        nw34_[int(1)] = d_261_v11_
        nw34_[int(2)] = d_261_v11_
        nw34_[int(3)] = (D0_DC0(d_261_v11_)).cf0
        d_262_v12_ = nw34_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_262_v12_).length(0)):
            d_263_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_263_i0_)) and ((d_263_i0_) < ((d_262_v12_).length(0)))):
                (d_262_v12_)[(d_263_i0_)] = ((_dafny.MultiSet([d_260_v10_, (D1_DC2(d_260_v10_)).cf3, (D1_DC2(d_260_v10_)).cf3, d_260_v10_])).intersection(d_261_v11_) if d_258_v8_ else d_261_v11_)
        d_264_v13_: bool
        d_265_v14_: _dafny.Seq
        out2_: bool
        out3_: _dafny.Seq
        out2_, out3_ = default__.m0((not(False)) or (d_258_v8_), d_249_v0_, d_251_v2_, (0) - (d_251_v2_), d_255_globalState_)
        d_264_v13_ = out2_
        d_265_v14_ = out3_
        d_266_v15_: _dafny.Seq
        d_266_v15_ = _dafny.SeqWithoutIsStrInference([d_253_v4_])
        d_267_v16_: _dafny.Seq
        d_267_v16_ = _dafny.SeqWithoutIsStrInference([d_266_v15_])
        d_268_v17_: _dafny.Map
        d_268_v17_ = _dafny.Map({(0) - ((d_261_v11_).cardinality): d_251_v2_})
        d_269_v18_: _dafny.Map
        d_269_v18_ = _dafny.Map({len(d_268_v17_): d_251_v2_})
        d_270_v19_: _dafny.Map
        d_270_v19_ = _dafny.Map({d_251_v2_: len(d_268_v17_)})
        d_271_v20_: _dafny.Map
        d_271_v20_ = _dafny.Map({d_251_v2_: (d_270_v19_).set(d_251_v2_, (d_256_v6_)[default__.safeIndex(d_251_v2_, len(d_256_v6_))])})
        d_272_v21_: _dafny.Array
        nw35_ = _dafny.Array(None, 9)
        nw35_[int(0)] = d_251_v2_
        nw35_[int(1)] = (0) - (((d_261_v11_)[d_260_v10_] if (d_260_v10_) in (d_261_v11_) else d_251_v2_))
        nw35_[int(2)] = d_251_v2_
        nw35_[int(3)] = 809
        nw35_[int(4)] = (default__.fm0(d_264_v13_, d_255_globalState_)) * ((0) - (d_251_v2_))
        nw35_[int(5)] = (_dafny.MultiSet(((d_256_v6_).set(default__.safeIndex(d_251_v2_, len(d_256_v6_)), d_251_v2_)) + (d_256_v6_))).cardinality
        nw35_[int(6)] = d_251_v2_
        nw35_[int(7)] = len((d_267_v16_)[default__.safeIndex(len(((d_271_v20_)[d_251_v2_] if (d_251_v2_) in (d_271_v20_) else d_270_v19_)), len(d_267_v16_))])
        nw35_[int(8)] = d_251_v2_
        d_272_v21_ = nw35_
        index32_ = default__.safeIndex(953, (d_272_v21_).length(0))
        (d_272_v21_)[index32_] = d_251_v2_
        d_273_v22_: _dafny.MultiSet
        d_273_v22_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_251_v2_])])
        index33_ = default__.safeIndex(953, (d_272_v21_).length(0))
        (d_272_v21_)[index33_] = default__.fm0((d_251_v2_) >= (((d_273_v22_)[d_256_v6_] if (d_256_v6_) in (d_273_v22_) else d_251_v2_)), d_255_globalState_)
        if d_264_v13_:
            d_274_v23_: _dafny.Map
            d_274_v23_ = _dafny.Map({d_264_v13_: (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]})
            d_275_v24_: _dafny.Map
            d_275_v24_ = _dafny.Map({d_249_v0_: default__.safeModuloInt(len(d_274_v23_), d_251_v2_)})
            d_274_v23_ = (d_274_v23_).set(not(((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]) > ((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))])), (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))])
            (d_255_globalState_).f2 = (d_249_v0_ if d_249_v0_ else True)
            d_276_v25_: C0
            nw36_ = C0()
            nw36_.ctor__()
            d_276_v25_ = nw36_
            d_277_v26_: _dafny.Set
            d_277_v26_ = _dafny.Set({d_249_v0_})
            d_278_v27_: D16
            d_278_v27_ = D16_DC46(d_277_v26_, d_259_v9_, _dafny.MultiSet([d_258_v8_, ((d_252_v3_)[d_258_v8_] if (d_258_v8_) in (d_252_v3_) else d_258_v8_), d_264_v13_]))
            d_279_v28_: bool
            d_280_v29_: _dafny.Seq
            out4_: bool
            out5_: _dafny.Seq
            out4_, out5_ = (d_276_v25_).m12((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efayxjylm"))) + (d_254_v5_), d_275_v24_, (d_250_v1_) == ((d_278_v27_).cf87), default__.safeModuloInt(((d_275_v24_)[d_264_v13_] if (d_264_v13_) in (d_275_v24_) else (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]), (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]), d_255_globalState_)
            d_279_v28_ = out4_
            d_280_v29_ = out5_
            index34_ = default__.safeIndex(953, (d_272_v21_).length(0))
            rhs23_ = (0) - ((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))])
            rhs24_ = d_258_v8_
            rhs25_ = d_251_v2_
            lhs20_ = d_272_v21_
            lhs21_ = default__.safeIndex(953, (d_272_v21_).length(0))
            lhs22_ = d_255_globalState_
            lhs23_ = d_255_globalState_
            lhs20_[lhs21_] = rhs23_
            lhs22_.f2 = rhs24_
            lhs23_.f6 = rhs25_
        elif True:
            d_264_v13_ = d_258_v8_
            d_281_v30_: C8
            nw37_ = C8()
            nw37_.ctor__(False, d_264_v13_)
            d_281_v30_ = nw37_
            d_282_v31_: _dafny.Array
            def lambda13_(d_283_v5_):
                def lambda14_(d_284_i1_):
                    return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpi"))) + (d_283_v5_)

                return lambda14_

            init6_ = lambda13_(d_254_v5_)
            nw38_ = _dafny.Array(None, 25)
            for i0_6_ in range(nw38_.length(0)):
                nw38_[i0_6_] = init6_(i0_6_)
            d_282_v31_ = nw38_
            index35_ = default__.safeIndex(834, (d_282_v31_).length(0))
            (d_282_v31_)[index35_] = (d_254_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcoui")))
            d_285_v32_: _dafny.Map
            d_285_v32_ = _dafny.Map({d_281_v30_.f21: d_260_v10_})
            d_286_v33_: _dafny.Array
            nw39_ = _dafny.Array(_dafny.Map({}), 22)
            d_286_v33_ = nw39_
            index36_ = default__.safeIndex(323, (d_286_v33_).length(0))
            (d_286_v33_)[index36_] = default__.fm51(d_258_v8_, default__.fm0((d_281_v30_).f22, d_255_globalState_), (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], len(d_259_v9_), d_255_globalState_)
            d_287_v34_: _dafny.Map
            d_287_v34_ = _dafny.Map({d_282_v31_: not(d_258_v8_)})
            index37_ = default__.safeIndex(834, (d_282_v31_).length(0))
            index38_ = default__.safeIndex(323, (d_286_v33_).length(0))
            rhs26_ = (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]
            rhs27_ = default__.fm69(((d_287_v34_)[d_282_v31_] if (d_282_v31_) in (d_287_v34_) else d_264_v13_), len((d_254_v5_ if (d_281_v30_).f22 else d_254_v5_)), _dafny.MultiSet(d_254_v5_), d_251_v2_, d_255_globalState_)
            rhs28_ = (d_254_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))
            rhs29_ = d_285_v32_
            rhs30_ = d_252_v3_
            lhs24_ = d_255_globalState_
            lhs25_ = d_281_v30_
            lhs26_ = d_282_v31_
            lhs27_ = default__.safeIndex(834, (d_282_v31_).length(0))
            lhs28_ = d_286_v33_
            lhs29_ = default__.safeIndex(323, (d_286_v33_).length(0))
            lhs24_.f6 = rhs26_
            lhs25_.f21 = rhs27_
            lhs26_[lhs27_] = rhs28_
            d_285_v32_ = rhs29_
            lhs28_[lhs29_] = rhs30_
            d_288_v35_: C5
            nw40_ = C5()
            nw40_.ctor__()
            d_288_v35_ = nw40_
            d_289_v36_: _dafny.Set
            d_289_v36_ = _dafny.Set({d_258_v8_, True, d_281_v30_.f21, (d_281_v30_).f22})
            d_290_v37_: _dafny.Map
            d_290_v37_ = _dafny.Map({d_288_v35_: d_289_v36_})
            d_291_v38_: _dafny.Seq
            d_291_v38_ = _dafny.SeqWithoutIsStrInference([d_290_v37_])
            d_292_v39_: _dafny.MultiSet
            d_292_v39_ = _dafny.MultiSet([d_281_v30_.f21])
            index39_ = default__.safeIndex(953, (d_272_v21_).length(0))
            rhs31_ = 636
            rhs32_ = ((d_291_v38_)[default__.safeIndex((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], len(d_291_v38_))] if (d_281_v30_).f22 else d_290_v37_)
            rhs33_ = (d_292_v39_).ispropersubset(default__.fm54((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], d_255_globalState_))
            lhs30_ = d_272_v21_
            lhs31_ = default__.safeIndex(953, (d_272_v21_).length(0))
            lhs32_ = d_255_globalState_
            lhs30_[lhs31_] = rhs31_
            d_290_v37_ = rhs32_
            lhs32_.f2 = rhs33_
            d_293_v40_: T1
            nw41_ = C4()
            nw41_.ctor__(d_254_v5_, default__.safeModuloInt((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], d_251_v2_), d_258_v8_, d_254_v5_, 254, (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))])
            d_293_v40_ = nw41_
            d_293_v40_ = d_293_v40_
        d_294_v41_: C1
        nw42_ = C1()
        nw42_.ctor__(d_253_v4_, d_251_v2_, 102, (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], default__.fm7(d_249_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bu"))), d_255_globalState_), d_254_v5_)
        d_294_v41_ = nw42_
        d_295_i2_: int
        d_295_i2_ = 0
        with _dafny.label("1"):
            while d_249_v0_:
                with _dafny.c_label("1"):
                    if (d_295_i2_) >= (100):
                        raise _dafny.Break("1")
                    d_295_i2_ = (d_295_i2_) + (1)
                    (d_255_globalState_).f2 = d_258_v8_
                    if d_249_v0_:
                        (d_294_v41_).f24 = d_294_v41_.f24
                        (d_255_globalState_).f6 = (0) - ((d_256_v6_)[default__.safeIndex(len(d_256_v6_), len(d_256_v6_))])
                        index40_ = default__.safeIndex(953, (d_272_v21_).length(0))
                        (d_272_v21_)[index40_] = len(d_254_v5_)
                        index41_ = default__.safeIndex(953, (d_272_v21_).length(0))
                        (d_272_v21_)[index41_] = d_251_v2_
                        d_296_v42_: _dafny.Set
                        d_296_v42_ = _dafny.Set({(d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]})
                        d_297_v43_: _dafny.Set
                        d_297_v43_ = _dafny.Set({d_296_v42_})
                        d_298_v44_: _dafny.Map
                        d_298_v44_ = _dafny.Map({d_297_v43_: ((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]) <= ((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))])})
                        d_298_v44_ = (d_298_v44_).set(d_297_v43_, d_258_v8_)
                    elif True:
                        index42_ = default__.safeIndex(953, (d_272_v21_).length(0))
                        (d_272_v21_)[index42_] = d_251_v2_
                        d_264_v13_ = ((d_258_v8_ if d_249_v0_ else d_258_v8_)) == (d_249_v0_)
                        (d_294_v41_).m14(d_255_globalState_)
                        d_299_v45_: bool
                        d_300_v46_: int
                        d_301_v47_: int
                        out6_: bool
                        out7_: int
                        out8_: int
                        out6_, out7_, out8_ = (d_294_v41_).m1((d_254_v5_).set(default__.safeIndex(d_251_v2_, len(d_254_v5_)), d_260_v10_), d_251_v2_, d_258_v8_, ((0) - (len(d_256_v6_)) if d_264_v13_ else 498), d_255_globalState_)
                        d_299_v45_ = out6_
                        d_300_v46_ = out7_
                        d_301_v47_ = out8_
                        arr0_ = d_294_v41_.f24
                        index43_ = default__.safeIndex(888, (d_294_v41_.f24).length(0))
                        arr0_[index43_] = d_299_v45_
                        arr1_ = d_294_v41_.f24
                        index44_ = default__.safeIndex(888, (d_294_v41_.f24).length(0))
                        rhs34_ = True
                        rhs35_ = False
                        rhs36_ = (d_257_v7_) - (d_257_v7_)
                        lhs33_ = d_294_v41_.f24
                        lhs34_ = default__.safeIndex(888, (d_294_v41_.f24).length(0))
                        d_258_v8_ = rhs34_
                        lhs33_[lhs34_] = rhs35_
                        d_257_v7_ = rhs36_
                    d_302_v48_: _dafny.Set
                    d_302_v48_ = _dafny.Set({d_251_v2_, 78})
                    d_303_v49_: D16
                    d_303_v49_ = D16_DC45(d_253_v4_, d_260_v10_, len(d_302_v48_), d_251_v2_, d_258_v8_)
                    d_304_v50_: T3
                    nw43_ = C10()
                    nw43_.ctor__((d_254_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmwbirfxq"))), ((d_254_v5_).set(default__.safeIndex(d_251_v2_, len(d_254_v5_)), d_260_v10_)) != (d_254_v5_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dq")), (len(d_265_v14_)) + (d_251_v2_), not (d_249_v0_) or ((d_303_v49_).cf85), d_254_v5_, default__.safeDivisionInt((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], d_251_v2_), d_251_v2_)
                    d_304_v50_ = nw43_
                    d_304_v50_ = d_304_v50_
                    d_305_v51_: _dafny.Seq
                    d_305_v51_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_304_v50_).f14, (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]})])
                    d_306_v52_: _dafny.Map
                    d_306_v52_ = _dafny.Map({(d_304_v50_).f15: (d_304_v50_).f14})
                    d_307_v53_: D1
                    d_307_v53_ = D1_DC4((d_302_v48_).ispropersubset((d_305_v51_)[default__.safeIndex((d_294_v41_).fm2(len(d_306_v52_), d_264_v13_, (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], d_258_v8_, d_255_globalState_), len(d_305_v51_))]))
                    source9_ = d_307_v53_
                    if source9_.is_DC3:
                        d_308___mcc_h0_ = source9_.cf4
                        d_309_cf4_ = d_308___mcc_h0_
                        d_260_v10_ = d_260_v10_
                        arr2_ = d_294_v41_.f24
                        index45_ = default__.safeIndex(934, (d_294_v41_.f24).length(0))
                        arr2_[index45_] = not(d_258_v8_)
                        arr3_ = d_294_v41_.f24
                        index46_ = default__.safeIndex(934, (d_294_v41_.f24).length(0))
                        arr3_[index46_] = (default__.fm0(d_264_v13_, d_255_globalState_)) != (len(d_306_v52_))
                        d_306_v52_ = (d_306_v52_).set(((d_252_v3_)[d_264_v13_] if (d_264_v13_) in (d_252_v3_) else d_258_v8_), (d_304_v50_).f14)
                        (d_255_globalState_).f6 = (default__.fm0(False, d_255_globalState_)) * ((d_304_v50_).f11)
                    elif source9_.is_DC4:
                        d_310___mcc_h1_ = source9_.cf5
                        d_311_cf5_ = d_310___mcc_h1_
                        index47_ = default__.safeIndex(953, (d_272_v21_).length(0))
                        (d_272_v21_)[index47_] = (0) - ((d_304_v50_).f11)
                        arr4_ = d_294_v41_.f24
                        index48_ = default__.safeIndex(538, (d_294_v41_.f24).length(0))
                        arr4_[index48_] = (d_257_v7_).issubset(d_257_v7_)
                        arr5_ = d_294_v41_.f24
                        index49_ = default__.safeIndex(538, (d_294_v41_.f24).length(0))
                        rhs37_ = d_302_v48_
                        rhs38_ = d_264_v13_
                        lhs35_ = d_294_v41_.f24
                        lhs36_ = default__.safeIndex(538, (d_294_v41_.f24).length(0))
                        d_302_v48_ = rhs37_
                        lhs35_[lhs36_] = rhs38_
                        d_312_v54_: _dafny.Map
                        d_312_v54_ = _dafny.Map({(d_304_v50_).f15: d_304_v50_.f16})
                        d_313_v55_: bool
                        d_314_v56_: int
                        d_315_v57_: int
                        out9_: bool
                        out10_: int
                        out11_: int
                        out9_, out10_, out11_ = (d_304_v50_).m1(((d_312_v54_)[(d_294_v41_.f24)[default__.safeIndex(538, (d_294_v41_.f24).length(0))]] if ((d_294_v41_.f24)[default__.safeIndex(538, (d_294_v41_.f24).length(0))]) in (d_312_v54_) else d_304_v50_.f16), (len(_dafny.Map({d_251_v2_: d_258_v8_}))) * ((d_304_v50_).f12), (d_304_v50_).f15, 852, d_255_globalState_)
                        d_313_v55_ = out9_
                        d_314_v56_ = out10_
                        d_315_v57_ = out11_
                        d_316_v58_: int
                        d_317_v59_: bool
                        d_318_v60_: T0
                        out12_: int
                        out13_: bool
                        out14_: T0
                        out12_, out13_, out14_ = (d_304_v50_).m3(_dafny.Map({d_253_v4_: (d_256_v6_)[default__.safeIndex((d_304_v50_).f14, len(d_256_v6_))]}), default__.fm69((d_294_v41_.f24)[default__.safeIndex(538, (d_294_v41_.f24).length(0))], 906, d_261_v11_, (d_304_v50_).f11, d_255_globalState_), d_255_globalState_)
                        d_316_v58_ = out12_
                        d_317_v59_ = out13_
                        d_318_v60_ = out14_
                    elif source9_.is_DC5:
                        d_319___mcc_h2_ = source9_.cf6
                        d_320___mcc_h3_ = source9_.cf7
                        d_321___mcc_h4_ = source9_.cf8
                        d_322_cf8_ = d_321___mcc_h4_
                        d_323_cf7_ = d_320___mcc_h3_
                        d_324_cf6_ = d_319___mcc_h2_
                        d_260_v10_ = _dafny.CodePoint('b')
                        (d_255_globalState_).f6 = (d_304_v50_).f12
                        rhs39_ = not((d_304_v50_).f15)
                        rhs40_ = d_258_v8_
                        rhs41_ = (d_294_v41_).fm1(len(_dafny.SeqWithoutIsStrInference([d_323_cf7_])), (d_304_v50_).f15, d_249_v0_, d_255_globalState_)
                        lhs37_ = d_255_globalState_
                        lhs38_ = d_255_globalState_
                        lhs39_ = d_255_globalState_
                        lhs37_.f2 = rhs39_
                        lhs38_.f7 = rhs40_
                        lhs39_.f2 = rhs41_
                        d_325_v61_: _dafny.Array
                        nw44_ = _dafny.Array(_dafny.Seq({}), 14)
                        d_325_v61_ = nw44_
                        index50_ = default__.safeIndex(679, (d_325_v61_).length(0))
                        (d_325_v61_)[index50_] = default__.fm20(d_258_v8_, d_249_v0_, d_322_cf8_, d_255_globalState_)
                        index51_ = default__.safeIndex(679, (d_325_v61_).length(0))
                        (d_325_v61_)[index51_] = d_256_v6_
                    elif True:
                        d_326___mcc_h5_ = source9_.cf3
                        d_327_cf3_ = d_326___mcc_h5_
                        index52_ = default__.safeIndex(953, (d_272_v21_).length(0))
                        (d_272_v21_)[index52_] = (((d_306_v52_)[d_249_v0_] if (d_249_v0_) in (d_306_v52_) else d_251_v2_)) - (293)
                        d_260_v10_ = d_327_cf3_
                        d_328_v62_: _dafny.Array
                        d_329_v63_: _dafny.Set
                        out15_: _dafny.Array
                        out16_: _dafny.Set
                        out15_, out16_ = (d_294_v41_).m13(_dafny.Set({d_264_v13_}), d_255_globalState_)
                        d_328_v62_ = out15_
                        d_329_v63_ = out16_
                        d_327_cf3_ = d_327_cf3_
                    pass
            pass
        d_330_v64_: C5
        nw45_ = C5()
        nw45_.ctor__()
        d_330_v64_ = nw45_
        nw46_ = C5()
        nw46_.ctor__()
        d_330_v64_ = nw46_
        hi0_ = default__.fm0((D2_DC8(d_264_v13_, d_249_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))).cf16, d_255_globalState_)
        for d_331_i3_ in range(d_251_v2_, hi0_):
            d_332_v65_: _dafny.Seq
            d_332_v65_ = _dafny.SeqWithoutIsStrInference([d_254_v5_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_333_i4_ in range(default__.abs(453))]), d_254_v5_, d_254_v5_, d_254_v5_])
            d_264_v13_ = ((d_294_v41_).fm2(d_251_v2_, d_258_v8_, d_331_i3_, d_264_v13_, d_255_globalState_)) == (len(d_332_v65_))
            index53_ = default__.safeIndex(953, (d_272_v21_).length(0))
            (d_272_v21_)[index53_] = (d_331_i3_) * (d_331_i3_)
            d_260_v10_ = d_260_v10_
            d_334_v66_: _dafny.Array
            nw47_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_334_v66_ = nw47_
            d_335_v67_: _dafny.Array
            d_335_v67_ = d_334_v66_
            d_336_v68_: _dafny.Array
            nw48_ = _dafny.Array(None, 21)
            nw48_[int(0)] = d_334_v66_
            nw48_[int(1)] = d_334_v66_
            nw48_[int(2)] = d_334_v66_
            nw48_[int(3)] = d_334_v66_
            nw48_[int(4)] = d_334_v66_
            nw48_[int(5)] = d_334_v66_
            nw48_[int(6)] = d_334_v66_
            nw48_[int(7)] = d_334_v66_
            nw48_[int(8)] = (d_334_v66_ if d_258_v8_ else d_334_v66_)
            nw48_[int(9)] = d_334_v66_
            nw48_[int(10)] = d_334_v66_
            nw48_[int(11)] = d_334_v66_
            nw48_[int(12)] = d_334_v66_
            nw48_[int(13)] = d_334_v66_
            nw48_[int(14)] = d_334_v66_
            nw48_[int(15)] = d_334_v66_
            nw48_[int(16)] = d_334_v66_
            nw48_[int(17)] = d_334_v66_
            nw48_[int(18)] = d_334_v66_
            nw48_[int(19)] = (d_335_v67_)
            nw48_[int(20)] = d_334_v66_
            d_336_v68_ = nw48_
            index54_ = default__.safeIndex(509, (d_336_v68_).length(0))
            (d_336_v68_)[index54_] = d_334_v66_
            index55_ = default__.safeIndex(509, (d_336_v68_).length(0))
            (d_336_v68_)[index55_] = d_334_v66_
        d_337_v69_: _dafny.Map
        d_337_v69_ = _dafny.Map({d_272_v21_: d_258_v8_})
        d_252_v3_ = (d_252_v3_).set(((d_337_v69_)[d_272_v21_] if (d_272_v21_) in (d_337_v69_) else d_264_v13_), d_264_v13_)
        (d_255_globalState_).f6 = (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in _dafny.IntegerRange(935, -797):
                d_338_v70_: int = compr_20_
                if ((935) <= (d_338_v70_)) and ((d_338_v70_) < (-797)):
                    coll20_[default__.safeModuloInt(d_338_v70_, (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))])] = (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]
            return _dafny.Map(coll20_)
        def iife21_():
            coll21_ = _dafny.Map()
            compr_21_: int
            for compr_21_ in _dafny.IntegerRange(181, 804):
                d_339_v71_: int = compr_21_
                if ((181) <= (d_339_v71_)) and ((d_339_v71_) < (804)):
                    coll21_[default__.safeModuloInt(d_339_v71_, d_251_v2_)] = (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]
            return _dafny.Map(coll21_)
        (d_255_globalState_).f2 = (iife20_()
        ) != (iife21_()
        )
        d_340_v72_: C0
        nw49_ = C0()
        nw49_.ctor__()
        d_340_v72_ = nw49_
        d_341_v73_: _dafny.MultiSet
        d_341_v73_ = _dafny.MultiSet([d_340_v72_])
        (d_255_globalState_).f2 = ((d_256_v6_)[default__.safeIndex(((d_341_v73_)[d_340_v72_] if (d_340_v72_) in (d_341_v73_) else (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]), len(d_256_v6_))]) >= ((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))])
        d_342_v74_: _dafny.Array
        def lambda15_(d_343_v14_):
            def lambda16_(d_344_i5_):
                return d_343_v14_

            return lambda16_

        init7_ = lambda15_(d_265_v14_)
        nw50_ = _dafny.Array(None, 16)
        for i0_7_ in range(nw50_.length(0)):
            nw50_[i0_7_] = init7_(i0_7_)
        d_342_v74_ = nw50_
        index56_ = default__.safeIndex(540, (d_342_v74_).length(0))
        (d_342_v74_)[index56_] = d_259_v9_
        index57_ = default__.safeIndex(540, (d_342_v74_).length(0))
        (d_342_v74_)[index57_] = d_265_v14_
        d_345_i6_: int
        d_345_i6_ = 0
        with _dafny.label("2"):
            while d_249_v0_:
                with _dafny.c_label("2"):
                    if (d_345_i6_) >= (100):
                        raise _dafny.Break("2")
                    d_345_i6_ = (d_345_i6_) + (1)
                    index58_ = default__.safeIndex(867, (d_253_v4_).length(0))
                    (d_253_v4_)[index58_] = (d_251_v2_) < ((0) - (d_251_v2_))
                    d_346_v75_: _dafny.Seq
                    d_346_v75_ = _dafny.SeqWithoutIsStrInference([d_256_v6_])
                    index59_ = default__.safeIndex(867, (d_253_v4_).length(0))
                    (d_253_v4_)[index59_] = (d_340_v72_).fm17((d_256_v6_ if d_258_v8_ else (d_346_v75_)[default__.safeIndex((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], len(d_346_v75_))]), d_255_globalState_)
                    d_249_v0_ = ((0) - ((688) - ((0) - ((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))])))) != (d_251_v2_)
                    d_347_v76_: bool
                    d_348_v77_: int
                    d_349_v78_: int
                    out17_: bool
                    out18_: int
                    out19_: int
                    out17_, out18_, out19_ = (d_294_v41_).m1(d_254_v5_, (58) * (len(d_254_v5_)), d_249_v0_, (0) - (d_251_v2_), d_255_globalState_)
                    d_347_v76_ = out17_
                    d_348_v77_ = out18_
                    d_349_v78_ = out19_
                    d_350_v79_: _dafny.Map
                    d_350_v79_ = _dafny.Map({d_251_v2_: d_259_v9_})
                    d_351_v80_: D10
                    d_351_v80_ = D10_DC28(d_350_v79_)
                    pat_let_tv0_ = d_350_v79_
                    pat_let_tv1_ = d_350_v79_
                    def iife23_(_pat_let1_0):
                        def iife24_(d_352_dt__update__tmp_h0_):
                            def iife25_(_pat_let2_0):
                                def iife26_(d_353_dt__update_hcf47_h0_):
                                    return D10_DC28(d_353_dt__update_hcf47_h0_)
                                return iife26_(_pat_let2_0)
                            return iife25_(pat_let_tv0_)
                        return iife24_(_pat_let1_0)
                    def iife22_(_pat_let0_0):
                        def iife27_(d_354_dt__update__tmp_h1_):
                            def iife28_(_pat_let3_0):
                                def iife29_(d_355_dt__update_hcf47_h1_):
                                    return D10_DC28(d_355_dt__update_hcf47_h1_)
                                return iife29_(_pat_let3_0)
                            return iife28_(pat_let_tv1_)
                        return iife27_(_pat_let0_0)
                    source10_ = iife22_(iife23_(d_351_v80_))
                    if source10_.is_DC29:
                        index60_ = default__.safeIndex(953, (d_272_v21_).length(0))
                        (d_272_v21_)[index60_] = default__.safeDivisionInt((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], d_251_v2_)
                        d_356_v81_: D0
                        d_356_v81_ = D0_DC0(d_261_v11_)
                        d_357_v82_: bool
                        d_358_v83_: int
                        d_359_v84_: int
                        out20_: bool
                        out21_: int
                        out22_: int
                        out20_, out21_, out22_ = (d_294_v41_).m1((d_254_v5_) + (default__.fm31(d_260_v10_, d_356_v81_, d_255_globalState_)), d_349_v78_, d_249_v0_, len((d_254_v5_) + (d_254_v5_)), d_255_globalState_)
                        d_357_v82_ = out20_
                        d_358_v83_ = out21_
                        d_359_v84_ = out22_
                        index61_ = default__.safeIndex(867, (d_253_v4_).length(0))
                        (d_253_v4_)[index61_] = d_357_v82_
                        d_360_v85_: C7
                        nw51_ = C7()
                        nw51_.ctor__()
                        d_360_v85_ = nw51_
                    elif source10_.is_DC30:
                        d_361___mcc_h6_ = source10_.cf48
                        d_362___mcc_h7_ = source10_.cf49
                        d_363___mcc_h8_ = source10_.cf50
                        d_364___mcc_h9_ = source10_.cf51
                        d_365_cf51_ = d_364___mcc_h9_
                        d_366_cf50_ = d_363___mcc_h8_
                        d_367_cf49_ = d_362___mcc_h7_
                        d_368_cf48_ = d_361___mcc_h6_
                        d_254_v5_ = d_366_cf50_
                        d_369_v86_: _dafny.Set
                        d_369_v86_ = _dafny.Set({d_349_v78_, d_349_v78_, (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]})
                        d_370_v87_: D0
                        d_370_v87_ = D0_DC0(_dafny.MultiSet(d_366_cf50_))
                        d_371_v88_: C2
                        nw52_ = C2()
                        nw52_.ctor__(d_349_v78_, default__.fm31(d_260_v10_, d_370_v87_, d_255_globalState_), (0) - (d_251_v2_), d_251_v2_)
                        d_371_v88_ = nw52_
                        d_372_v89_: D18
                        d_372_v89_ = D18_DC53(495, d_369_v86_, (0) - (d_251_v2_), d_371_v88_)
                        d_373_v90_: C8
                        nw53_ = C8()
                        nw53_.ctor__(d_258_v8_, d_365_cf51_)
                        d_373_v90_ = nw53_
                        d_374_v91_: _dafny.Map
                        d_374_v91_ = _dafny.Map({len((d_372_v89_).cf100): d_373_v90_})
                        d_375_v92_: C1
                        nw54_ = C1()
                        nw54_.ctor__(d_294_v41_.f24, (d_256_v6_)[default__.safeIndex(d_349_v78_, len(d_256_v6_))], (d_294_v41_).fm2(len(d_374_v91_), (d_253_v4_)[default__.safeIndex(867, (d_253_v4_).length(0))], (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], d_368_cf48_, d_255_globalState_), len((d_366_cf50_ if d_365_cf51_ else d_254_v5_)), (d_340_v72_).fm17(_dafny.SeqWithoutIsStrInference([d_251_v2_]), d_255_globalState_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cv")))
                        d_375_v92_ = nw54_
                        (d_375_v92_).m14(d_255_globalState_)
                        (d_255_globalState_).f6 = d_251_v2_
                    elif source10_.is_DC31:
                        d_376___mcc_h10_ = source10_.cf52
                        d_377___mcc_h11_ = source10_.cf53
                        d_378___mcc_h12_ = source10_.cf54
                        d_379___mcc_h13_ = source10_.cf55
                        d_380_cf55_ = d_379___mcc_h13_
                        d_381_cf54_ = d_378___mcc_h12_
                        d_382_cf53_ = d_377___mcc_h11_
                        d_383_cf52_ = d_376___mcc_h10_
                        (d_255_globalState_).f2 = (d_330_v64_).fm13(d_380_cf55_, _dafny.CodePoint('v'), d_382_cf53_, d_255_globalState_)
                        d_384_v93_: _dafny.Map
                        d_384_v93_ = _dafny.Map({(0) - (d_383_cf52_): d_347_v76_})
                        d_385_v94_: _dafny.Seq
                        d_385_v94_ = _dafny.SeqWithoutIsStrInference([d_384_v93_])
                        d_386_v95_: D15
                        d_386_v95_ = D15_DC43(d_348_v77_, d_264_v13_, d_258_v8_)
                        d_387_v96_: _dafny.MultiSet
                        d_387_v96_ = _dafny.MultiSet([d_249_v0_])
                        d_388_v97_: C6
                        nw55_ = C6()
                        nw55_.ctor__(d_385_v94_, d_381_cf54_, default__.fm7((d_386_v95_).cf78, (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))], d_255_globalState_), d_254_v5_, (0) - ((d_387_v96_).cardinality), d_349_v78_)
                        d_388_v97_ = nw55_
                        d_389_v98_: _dafny.Map
                        d_389_v98_ = _dafny.Map({d_388_v97_: d_249_v0_})
                        d_390_v99_: _dafny.Set
                        d_390_v99_ = _dafny.Set({(d_383_cf52_) != ((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))]), d_347_v76_, d_249_v0_, d_264_v13_, (default__.fm7(d_249_v0_, d_348_v77_, d_255_globalState_) if d_347_v76_ else ((d_389_v98_)[d_388_v97_] if (d_388_v97_) in (d_389_v98_) else d_347_v76_))})
                        d_391_v100_: _dafny.Map
                        d_391_v100_ = _dafny.Map({d_260_v10_: d_381_cf54_})
                        d_392_v101_: _dafny.Set
                        d_392_v101_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))), (d_340_v72_).fm16(d_347_v76_, d_349_v78_, (d_294_v41_).fm2(232, d_264_v13_, -916, d_347_v76_, d_255_globalState_), d_380_cf55_, d_255_globalState_), (0) - (d_251_v2_)})
                        d_390_v99_ = default__.fm32((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aeqdtt"))) != (d_254_v5_), d_391_v100_, d_392_v101_, d_347_v76_, d_255_globalState_)
                        d_393_v102_: D1
                        d_393_v102_ = D1_DC2(_dafny.CodePoint('s'))
                        index62_ = default__.safeIndex(867, (d_253_v4_).length(0))
                        rhs42_ = (d_393_v102_).cf3
                        rhs43_ = d_390_v99_
                        rhs44_ = not((not(d_258_v8_) if default__.fm7(d_249_v0_, d_349_v78_, d_255_globalState_) else (False) and (d_258_v8_)))
                        lhs40_ = d_253_v4_
                        lhs41_ = default__.safeIndex(867, (d_253_v4_).length(0))
                        d_260_v10_ = rhs42_
                        d_390_v99_ = rhs43_
                        lhs40_[lhs41_] = rhs44_
                        d_394_v103_: C4
                        nw56_ = C4()
                        nw56_.ctor__(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vg")), d_348_v77_, d_347_v76_, d_254_v5_, d_383_cf52_, d_383_cf52_)
                        d_394_v103_ = nw56_
                        d_395_v104_: _dafny.Set
                        d_395_v104_ = _dafny.Set({d_394_v103_})
                        d_347_v76_ = ((0) - (default__.safeModuloInt(d_383_cf52_, len(d_395_v104_)))) == (len((d_390_v99_) | (d_390_v99_)))
                    elif source10_.is_DC28:
                        d_396___mcc_h14_ = source10_.cf47
                        d_397_cf47_ = d_396___mcc_h14_
                        d_347_v76_ = d_258_v8_
                        d_398_v105_: C0
                        nw57_ = C0()
                        nw57_.ctor__()
                        d_398_v105_ = nw57_
                        d_399_v106_: _dafny.Map
                        d_399_v106_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_260_v10_ for d_400_i7_ in range(default__.abs(918))]): d_264_v13_})
                        d_401_v107_: _dafny.Map
                        d_401_v107_ = _dafny.Map({d_249_v0_: d_399_v106_})
                        index63_ = default__.safeIndex(867, (d_253_v4_).length(0))
                        (d_253_v4_)[index63_] = not((d_249_v0_ if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmnbmsk"))) not in (((d_401_v107_)[d_249_v0_] if (d_249_v0_) in (d_401_v107_) else d_399_v106_)) else (d_253_v4_)[default__.safeIndex(867, (d_253_v4_).length(0))]))
                        d_402_v108_: T3
                        nw58_ = C10()
                        nw58_.ctor__(d_254_v5_, (d_330_v64_).fm13(d_251_v2_, d_260_v10_, d_260_v10_, d_255_globalState_), d_254_v5_, d_348_v77_, d_258_v8_, d_254_v5_, d_349_v78_, (d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))])
                        d_402_v108_ = nw58_
                        d_403_v109_: _dafny.Map
                        d_403_v109_ = _dafny.Map({(d_272_v21_ if (d_253_v4_)[default__.safeIndex(867, (d_253_v4_).length(0))] else d_272_v21_): d_402_v108_})
                        d_403_v109_ = (d_403_v109_).set(d_272_v21_, d_402_v108_)
                    elif True:
                        d_404___mcc_h15_ = source10_.cf56
                        d_405_cf56_ = d_404___mcc_h15_
                        d_406_v110_: _dafny.Map
                        d_406_v110_ = _dafny.Map({(d_253_v4_)[default__.safeIndex(867, (d_253_v4_).length(0))]: d_348_v77_})
                        d_407_v111_: bool
                        d_408_v112_: _dafny.Seq
                        out23_: bool
                        out24_: _dafny.Seq
                        out23_, out24_ = (d_340_v72_).m12(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dp")), d_406_v110_, d_264_v13_, d_349_v78_, d_255_globalState_)
                        d_407_v111_ = out23_
                        d_408_v112_ = out24_
                        (d_255_globalState_).f2 = True
                        d_409_v113_: _dafny.Array
                        nw59_ = _dafny.Array(_dafny.Array(None, 0), 20)
                        d_409_v113_ = nw59_
                        index64_ = default__.safeIndex(940, (d_409_v113_).length(0))
                        (d_409_v113_)[index64_] = d_272_v21_
                        index65_ = default__.safeIndex(940, (d_409_v113_).length(0))
                        (d_409_v113_)[index65_] = d_272_v21_
                        d_410_v114_: _dafny.Map
                        d_410_v114_ = _dafny.Map({d_294_v41_.f24: d_348_v77_})
                        d_411_v115_: C4
                        nw60_ = C4()
                        nw60_.ctor__(_dafny.SeqWithoutIsStrInference([d_260_v10_ for d_412_i8_ in range(default__.abs(670))]), (d_251_v2_) - (default__.fm0(d_347_v76_, d_255_globalState_)), not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdtdq"))) == (d_254_v5_)), (d_254_v5_ if d_258_v8_ else d_254_v5_), d_349_v78_, ((D5_DC18(d_410_v114_, d_349_v78_)).cf36) * (len(d_254_v5_)))
                        d_411_v115_ = nw60_
                        index66_ = default__.safeIndex(953, (d_272_v21_).length(0))
                        rhs45_ = ((0) - (d_348_v77_) if not(True) else (d_348_v77_) - (d_349_v78_))
                        rhs46_ = default__.safeDivisionInt(((_dafny.MultiSet((d_342_v74_)[default__.safeIndex(540, (d_342_v74_).length(0))])).cardinality) * ((0) - ((d_272_v21_)[default__.safeIndex(953, (d_272_v21_).length(0))])), (0) - (d_349_v78_))
                        rhs47_ = not((d_253_v4_)[default__.safeIndex(867, (d_253_v4_).length(0))])
                        rhs48_ = d_411_v115_
                        rhs49_ = (d_406_v110_).set((d_249_v0_ if (d_253_v4_)[default__.safeIndex(867, (d_253_v4_).length(0))] else d_264_v13_), default__.safeModuloInt(-699, 90))
                        lhs42_ = d_272_v21_
                        lhs43_ = default__.safeIndex(953, (d_272_v21_).length(0))
                        lhs44_ = d_255_globalState_
                        d_349_v78_ = rhs45_
                        lhs42_[lhs43_] = rhs46_
                        lhs44_.f2 = rhs47_
                        d_411_v115_ = rhs48_
                        d_406_v110_ = rhs49_
                    pass
            pass
        _dafny.print(_dafny.string_of(d_249_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_250_v1_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_251_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v3_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_253_v4_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_253_v4_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_253_v4_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_253_v4_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_253_v4_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_253_v4_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_254_v5_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_globalState_).f0) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_255_globalState_.f1) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_255_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_globalState_).f3)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_globalState_).f3)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_globalState_).f3)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_globalState_).f3)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_globalState_).f3)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_255_globalState_).f3)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_255_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_255_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_255_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_255_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_255_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_255_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_255_globalState_).f10).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v6_) == (_dafny.SeqWithoutIsStrInference([67, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_257_v7_) == (_dafny.MultiSet([67, 67]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_258_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v9_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_260_v10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v11_) == (_dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v12_)[0]) == (_dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v12_)[1]) == (_dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v12_)[2]) == (_dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_262_v12_)[3]) == (_dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_264_v13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_265_v14_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_266_v15_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_267_v16_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_268_v17_) == (_dafny.Map({-5: 67}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v18_) == (_dafny.Map({1: 67}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v19_) == (_dafny.Map({67: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_v20_) == (_dafny.Map({67: _dafny.Map({67: 1})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v21_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v21_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v21_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v21_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v21_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v21_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v21_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v21_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v21_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_273_v22_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([67])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v41_.f24)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v41_.f24)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v41_.f24)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v41_.f24)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v41_.f24)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v41_.f24)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_295_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_337_v69_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_341_v73_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[0]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[1]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[2]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[3]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[4]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[5]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[6]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[7]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[8]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[9]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[10]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[11]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[12]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[13]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[14]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v74_)[15]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_345_i6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, int(0))
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
        return lambda: D1_DC3(False)
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

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(_dafny.Array(None, 0), _dafny.Set({}), D1.default()(), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC7(D2, NamedTuple('DC7', [('cf10', Any), ('cf11', Any), ('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {self.cf17.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({self.cf9.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(False, int(0), _dafny.Map({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
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
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)

class D3_DC10(D3, NamedTuple('DC10', [('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {self.cf22.VerbatimString(True)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14(_dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D4_DC16)

class D4_DC14(D4, NamedTuple('DC14', [('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC16(D4, NamedTuple('DC16', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC16({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC16) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC18(_dafny.Map({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)

class D5_DC18(D5, NamedTuple('DC18', [('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC17(D5, NamedTuple('DC17', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC20(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D6_DC21)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D6_DC22)

class D6_DC20(D6, NamedTuple('DC20', [('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC21(D6, NamedTuple('DC21', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC22(D6, NamedTuple('DC22', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC22({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC22) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D7_DC23)

class D7_DC23(D7, NamedTuple('DC23', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC23({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC23) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC25(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)

class D8_DC25(D8, NamedTuple('DC25', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC27(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)

class D9_DC27(D9, NamedTuple('DC27', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D10_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D10_DC31)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D10_DC32)

class D10_DC29(D10, NamedTuple('DC29', [])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29)
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC30(D10, NamedTuple('DC30', [('cf48', Any), ('cf49', Any), ('cf50', Any), ('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {self.cf50.VerbatimString(True)}, {_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC31(D10, NamedTuple('DC31', [('cf52', Any), ('cf53', Any), ('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC31({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC31) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC32(D10, NamedTuple('DC32', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC32({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC32) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC34(False, False, int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D11_DC34)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)

class D11_DC34(D11, NamedTuple('DC34', [('cf58', Any), ('cf59', Any), ('cf60', Any), ('cf61', Any), ('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC34({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC34) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC33(D11, NamedTuple('DC33', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC36(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)

class D12_DC36(D12, NamedTuple('DC36', [('cf64', Any), ('cf65', Any), ('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65 and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC35(D12, NamedTuple('DC35', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC38(D3.default()(), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D13_DC39)

class D13_DC38(D13, NamedTuple('DC38', [('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC37(D13, NamedTuple('DC37', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC39(D13, NamedTuple('DC39', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC39({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC39) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC41(False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D14_DC40)

class D14_DC41(D14, NamedTuple('DC41', [('cf73', Any), ('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {self.cf75.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC40(D14, NamedTuple('DC40', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC40({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC40) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC43(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)

class D15_DC43(D15, NamedTuple('DC43', [('cf77', Any), ('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC42(D15, NamedTuple('DC42', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC45(_dafny.Array(None, 0), _dafny.CodePoint('D'), int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D16_DC46)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)

class D16_DC45(D16, NamedTuple('DC45', [('cf81', Any), ('cf82', Any), ('cf83', Any), ('cf84', Any), ('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf81 == __o.cf81 and self.cf82 == __o.cf82 and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC46(D16, NamedTuple('DC46', [('cf86', Any), ('cf87', Any), ('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC46({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC46) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC44(D16, NamedTuple('DC44', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC48(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D17_DC48)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D17_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D17_DC50)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D17_DC47)

class D17_DC48(D17, NamedTuple('DC48', [('cf90', Any), ('cf91', Any), ('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC48({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC48) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC49(D17, NamedTuple('DC49', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC49({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC49) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC50(D17, NamedTuple('DC50', [('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC50({_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC50) and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC47(D17, NamedTuple('DC47', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC47({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC47) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC52(False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D18_DC52)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D18_DC53)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D18_DC51)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D18_DC54)

class D18_DC52(D18, NamedTuple('DC52', [('cf96', Any), ('cf97', Any), ('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC52({_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)}, {self.cf98.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC52) and self.cf96 == __o.cf96 and self.cf97 == __o.cf97 and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC53(D18, NamedTuple('DC53', [('cf99', Any), ('cf100', Any), ('cf101', Any), ('cf102', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC53({_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC53) and self.cf99 == __o.cf99 and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC51(D18, NamedTuple('DC51', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC51({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC51) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC54(D18, NamedTuple('DC54', [('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC54({_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC54) and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC56(_dafny.Array(None, 0), False, int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D19_DC56)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D19_DC55)

class D19_DC56(D19, NamedTuple('DC56', [('cf105', Any), ('cf106', Any), ('cf107', Any), ('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC56({_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC56) and self.cf105 == __o.cf105 and self.cf106 == __o.cf106 and self.cf107 == __o.cf107 and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC55(D19, NamedTuple('DC55', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC55({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC55) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC58(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D20_DC58)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D20_DC57)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D20_DC59)

class D20_DC58(D20, NamedTuple('DC58', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC58({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC58) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC57(D20, NamedTuple('DC57', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC57({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC57) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC59(D20, NamedTuple('DC59', [('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC59({_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC59) and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC61(False, False, _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D21_DC61)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D21_DC60)

class D21_DC61(D21, NamedTuple('DC61', [('cf113', Any), ('cf114', Any), ('cf115', Any), ('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC61({_dafny.string_of(self.cf113)}, {_dafny.string_of(self.cf114)}, {_dafny.string_of(self.cf115)}, {_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC61) and self.cf113 == __o.cf113 and self.cf114 == __o.cf114 and self.cf115 == __o.cf115 and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC60(D21, NamedTuple('DC60', [('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC60({_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC60) and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC63(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D22_DC63)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D22_DC62)

class D22_DC63(D22, NamedTuple('DC63', [('cf118', Any), ('cf119', Any), ('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC63({_dafny.string_of(self.cf118)}, {_dafny.string_of(self.cf119)}, {_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC63) and self.cf118 == __o.cf118 and self.cf119 == __o.cf119 and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC62(D22, NamedTuple('DC62', [('cf117', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC62({_dafny.string_of(self.cf117)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC62) and self.cf117 == __o.cf117
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D23_DC64)

class D23_DC64(D23, NamedTuple('DC64', [('cf121', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC64({_dafny.string_of(self.cf121)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC64) and self.cf121 == __o.cf121
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC66(None, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D24_DC66)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D24_DC65)

class D24_DC66(D24, NamedTuple('DC66', [('cf123', Any), ('cf124', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC66({_dafny.string_of(self.cf123)}, {_dafny.string_of(self.cf124)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC66) and self.cf123 == __o.cf123 and self.cf124 == __o.cf124
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC65(D24, NamedTuple('DC65', [('cf122', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC65({_dafny.string_of(self.cf122)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC65) and self.cf122 == __o.cf122
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC68()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D25_DC68)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D25_DC67)

class D25_DC68(D25, NamedTuple('DC68', [])):
    def __dafnystr__(self) -> str:
        return f'D25.DC68'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC68)
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC67(D25, NamedTuple('DC67', [('cf125', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC67({_dafny.string_of(self.cf125)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC67) and self.cf125 == __o.cf125
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC70(None, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D26_DC70)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D26_DC69)

class D26_DC70(D26, NamedTuple('DC70', [('cf127', Any), ('cf128', Any), ('cf129', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC70({_dafny.string_of(self.cf127)}, {_dafny.string_of(self.cf128)}, {_dafny.string_of(self.cf129)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC70) and self.cf127 == __o.cf127 and self.cf128 == __o.cf128 and self.cf129 == __o.cf129
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC69(D26, NamedTuple('DC69', [('cf126', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC69({_dafny.string_of(self.cf126)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC69) and self.cf126 == __o.cf126
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC72(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D27_DC72)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D27_DC71)

class D27_DC72(D27, NamedTuple('DC72', [('cf131', Any), ('cf132', Any), ('cf133', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC72({_dafny.string_of(self.cf131)}, {_dafny.string_of(self.cf132)}, {_dafny.string_of(self.cf133)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC72) and self.cf131 == __o.cf131 and self.cf132 == __o.cf132 and self.cf133 == __o.cf133
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC71(D27, NamedTuple('DC71', [('cf130', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC71({_dafny.string_of(self.cf130)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC71) and self.cf130 == __o.cf130
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC74(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D28_DC74)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D28_DC75)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D28_DC73)

class D28_DC74(D28, NamedTuple('DC74', [('cf135', Any), ('cf136', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC74({_dafny.string_of(self.cf135)}, {_dafny.string_of(self.cf136)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC74) and self.cf135 == __o.cf135 and self.cf136 == __o.cf136
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC75(D28, NamedTuple('DC75', [('cf137', Any), ('cf138', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC75({_dafny.string_of(self.cf137)}, {_dafny.string_of(self.cf138)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC75) and self.cf137 == __o.cf137 and self.cf138 == __o.cf138
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC73(D28, NamedTuple('DC73', [('cf134', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC73({_dafny.string_of(self.cf134)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC73) and self.cf134 == __o.cf134
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC77(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D29_DC77)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D29_DC76)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D29_DC78)

class D29_DC77(D29, NamedTuple('DC77', [('cf140', Any), ('cf141', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC77({_dafny.string_of(self.cf140)}, {_dafny.string_of(self.cf141)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC77) and self.cf140 == __o.cf140 and self.cf141 == __o.cf141
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC76(D29, NamedTuple('DC76', [('cf139', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC76({_dafny.string_of(self.cf139)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC76) and self.cf139 == __o.cf139
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC78(D29, NamedTuple('DC78', [('cf142', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC78({_dafny.string_of(self.cf142)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC78) and self.cf142 == __o.cf142
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, p2, p3, globalState):
        pass


class T1:
    pass

class T2:
    pass
    def m2(self, p0, p1, globalState):
        pass


class T3:
    pass
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    def m3(self, p0, p1, globalState):
        pass

    def m4(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: _dafny.Map = _dafny.Map({})
        self.f2: bool = False
        self.f6: int = int(0)
        self.f7: bool = False
        self._f0: _dafny.Seq = _dafny.Seq({})
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f4: int = int(0)
        self._f5: str = _dafny.CodePoint('D')
        self._f8: int = int(0)
        self._f9: int = int(0)
        self._f10: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10):
        (self)._f0 = f0
        (self).f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10

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
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm16(self, p0, p1, p2, p3, globalState):
        return (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_413_i0_ in range(default__.abs(404))]))) * (243)

    def fm17(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False, True, False])])) != (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False, True, True, False])]))

    def m12(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.Seq({})
        (globalState).f6 = (0) - ((0) - (p3))
        d_414_v0_: _dafny.Seq
        d_414_v0_ = _dafny.SeqWithoutIsStrInference([201])
        r0 = (self).fm17(d_414_v0_, globalState)
        (globalState).f7 = p2
        d_415_v1_: D2
        d_415_v1_ = D2_DC8((self).fm17(_dafny.SeqWithoutIsStrInference([p3 for d_416_i0_ in range(default__.abs(994))]), globalState), p2, p0)
        pat_let_tv2_ = p3
        pat_let_tv3_ = p3
        pat_let_tv4_ = p3
        pat_let_tv5_ = p2
        def lambda17_(source11_):
            if source11_.is_DC7:
                d_417___mcc_h0_ = source11_.cf10
                d_418___mcc_h1_ = source11_.cf11
                d_419___mcc_h2_ = source11_.cf12
                d_420___mcc_h3_ = source11_.cf13
                d_421___mcc_h4_ = source11_.cf14
                d_422_cf14_ = d_421___mcc_h4_
                d_423_cf13_ = d_420___mcc_h3_
                d_424_cf12_ = d_419___mcc_h2_
                d_425_cf11_ = d_418___mcc_h1_
                d_426_cf10_ = d_417___mcc_h0_
                return (pat_let_tv2_) == (pat_let_tv3_)
            elif source11_.is_DC8:
                d_427___mcc_h5_ = source11_.cf15
                d_428___mcc_h6_ = source11_.cf16
                d_429___mcc_h7_ = source11_.cf17
                d_430_cf17_ = d_429___mcc_h7_
                d_431_cf16_ = d_428___mcc_h6_
                d_432_cf15_ = d_427___mcc_h5_
                return (D0_DC1(d_431_cf16_, pat_let_tv4_)).cf1
            elif True:
                d_433___mcc_h8_ = source11_.cf9
                d_434_cf9_ = d_433___mcc_h8_
                return pat_let_tv5_

        (globalState).f7 = lambda17_(d_415_v1_)
        d_435_v2_: _dafny.MultiSet
        d_435_v2_ = _dafny.MultiSet([p2, p2, True, p2])
        d_436_v3_: _dafny.Seq
        d_436_v3_ = _dafny.SeqWithoutIsStrInference([p2])
        d_437_v4_: _dafny.MultiSet
        d_437_v4_ = _dafny.MultiSet([(d_435_v2_).set(p2, default__.abs((0) - (len(d_436_v3_))))])
        d_437_v4_ = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_435_v2_, d_435_v2_, (_dafny.MultiSet([p2])) - (d_435_v2_), d_435_v2_, (d_435_v2_ if (self).fm17(d_414_v0_, globalState) else d_435_v2_)]))
        (globalState).f7 = p2
        r0 = p2
        d_438_v5_: D1
        d_438_v5_ = D1_DC4(p2)
        d_439_v6_: _dafny.Seq
        d_439_v6_ = _dafny.SeqWithoutIsStrInference([d_438_v5_, d_438_v5_, d_438_v5_])
        r1 = (d_439_v6_) + ((d_439_v6_) + (d_439_v6_))
        return r0, r1


class C1(T0, T2, T1):
    def  __init__(self):
        self._f11: int = int(0)
        self._f12: int = int(0)
        self._f14: int = int(0)
        self._f15: bool = False
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f24: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
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
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f24, f11, f12, f14, f15, f13):
        (self).f24 = f24
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f13 = f13

    def fm2(self, p0, p1, p2, p3, globalState):
        return (self).f11

    def fm3(self, p0, p1, p2, globalState):
        return ((_dafny.Map({(self).f15: (self).f11})) | (_dafny.Map({(self).f15: (self).f12}))) | ((D6_DC19(_dafny.Map({(self).f15: 0}))).cf37)

    def fm1(self, p0, p1, p2, globalState):
        return True

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        hi1_ = (self).f12
        for d_440_i0_ in range(p1, hi1_):
            d_441_v0_: _dafny.Array
            nw61_ = _dafny.Array(int(0), 21)
            d_441_v0_ = nw61_
            d_442_v1_: _dafny.Map
            d_442_v1_ = _dafny.Map({len(p0): p2})
            d_443_v2_: _dafny.Map
            d_443_v2_ = _dafny.Map({d_441_v0_: ((d_442_v1_)[912] if (912) in (d_442_v1_) else p2)})
            d_443_v2_ = (d_443_v2_).set(d_441_v0_, not(((self).f15 if True else p2)))
            d_444_v3_: D4
            d_444_v3_ = D4_DC15((self).f11)
            d_445_v4_: _dafny.Array
            nw62_ = _dafny.Array(None, 25)
            nw62_[int(0)] = d_444_v3_
            nw62_[int(1)] = d_444_v3_
            nw62_[int(2)] = d_444_v3_
            nw62_[int(3)] = D4_DC15(d_440_i0_)
            nw62_[int(4)] = d_444_v3_
            nw62_[int(5)] = D4_DC15((self).f14)
            nw62_[int(6)] = d_444_v3_
            nw62_[int(7)] = d_444_v3_
            nw62_[int(8)] = d_444_v3_
            def iife30_(_pat_let4_0):
                def iife31_(d_446_dt__update__tmp_h0_):
                    def iife32_(_pat_let5_0):
                        def iife33_(d_447_dt__update_hcf32_h0_):
                            return D4_DC15(d_447_dt__update_hcf32_h0_)
                        return iife33_(_pat_let5_0)
                    return iife32_((self).f12)
                return iife31_(_pat_let4_0)
            nw62_[int(9)] = iife30_(d_444_v3_)
            nw62_[int(10)] = D4_DC15(p1)
            nw62_[int(11)] = d_444_v3_
            nw62_[int(12)] = default__.fm19(p2, (self).f15, _dafny.SeqWithoutIsStrInference([p1 for d_448_i1_ in range(default__.abs(915))]), globalState)
            nw62_[int(13)] = D4_DC15(p1)
            nw62_[int(14)] = d_444_v3_
            nw62_[int(15)] = d_444_v3_
            nw62_[int(16)] = d_444_v3_
            nw62_[int(17)] = d_444_v3_
            nw62_[int(18)] = D4_DC15((self).f12)
            nw62_[int(19)] = d_444_v3_
            nw62_[int(20)] = d_444_v3_
            nw62_[int(21)] = d_444_v3_
            nw62_[int(22)] = d_444_v3_
            nw62_[int(23)] = d_444_v3_
            nw62_[int(24)] = d_444_v3_
            d_445_v4_ = nw62_
            index67_ = default__.safeIndex(309, (d_445_v4_).length(0))
            (d_445_v4_)[index67_] = D4_DC15(634)
            index68_ = default__.safeIndex(309, (d_445_v4_).length(0))
            rhs50_ = d_444_v3_
            rhs51_ = p3
            rhs52_ = d_441_v0_
            rhs53_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcpelm")))) * (p1)
            lhs45_ = d_445_v4_
            lhs46_ = default__.safeIndex(309, (d_445_v4_).length(0))
            lhs47_ = globalState
            lhs48_ = globalState
            lhs45_[lhs46_] = rhs50_
            lhs47_.f6 = rhs51_
            d_441_v0_ = rhs52_
            lhs48_.f6 = rhs53_
            if not(True):
                d_449_v5_: _dafny.Seq
                d_449_v5_ = _dafny.SeqWithoutIsStrInference([not((self).f15), True])
                d_450_v6_: _dafny.Map
                d_450_v6_ = _dafny.Map({497: p3})
                d_451_v7_: _dafny.Set
                d_451_v7_ = _dafny.Set({p2})
                d_452_v8_: _dafny.MultiSet
                d_452_v8_ = _dafny.MultiSet([len(((self).f13).set(default__.safeIndex(len(_dafny.Set({p1})), len((self).f13)), ((self).f13)[default__.safeIndex(d_440_i0_, len((self).f13))])), len((d_449_v5_).set(default__.safeIndex(((d_450_v6_)[len(d_451_v7_)] if (len(d_451_v7_)) in (d_450_v6_) else (self).f14), len(d_449_v5_)), True))])
                d_453_v9_: _dafny.Seq
                d_453_v9_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f12, p1])).ispropersubset(d_452_v8_), p2])
                (globalState).f2 = (d_449_v5_)[default__.safeIndex((self).f11, len(d_449_v5_))]
                (globalState).f2 = p2
                d_454_v10_: C0
                nw63_ = C0()
                nw63_.ctor__()
                d_454_v10_ = nw63_
                d_455_v11_: _dafny.Array
                def lambda18_(d_456_i0_):
                    def lambda19_(d_457_i2_):
                        return _dafny.SeqWithoutIsStrInference([d_456_i0_, (self).f12, (self).f12])

                    return lambda19_

                init8_ = lambda18_(d_440_i0_)
                nw64_ = _dafny.Array(None, 17)
                for i0_8_ in range(nw64_.length(0)):
                    nw64_[i0_8_] = init8_(i0_8_)
                d_455_v11_ = nw64_
                d_458_v12_: _dafny.Seq
                d_458_v12_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p2, globalState), d_440_i0_])
                index69_ = default__.safeIndex(370, (d_455_v11_).length(0))
                (d_455_v11_)[index69_] = d_458_v12_
                index70_ = default__.safeIndex(370, (d_455_v11_).length(0))
                (d_455_v11_)[index70_] = d_458_v12_
                (self).f24 = self.f24
            elif True:
                d_459_v13_: _dafny.MultiSet
                d_459_v13_ = _dafny.MultiSet([(self).f11])
                d_460_v14_: _dafny.MultiSet
                d_460_v14_ = _dafny.MultiSet([True])
                rhs54_ = (self).f12
                rhs55_ = (self).f15
                rhs56_ = (self).fm2((self).f11, p2, (d_459_v13_).cardinality, (d_460_v14_) != (d_460_v14_), globalState)
                rhs57_ = (self).f14
                lhs49_ = globalState
                lhs50_ = globalState
                r2 = rhs54_
                lhs49_.f7 = rhs55_
                lhs50_.f6 = rhs56_
                r2 = rhs57_
                d_461_v15_: _dafny.Map
                d_461_v15_ = _dafny.Map({(self).f15: d_440_i0_})
                d_462_v16_: _dafny.Map
                d_462_v16_ = _dafny.Map({p1: ((d_461_v15_).set((self).f15, d_440_i0_)) | (d_461_v15_)})
                d_461_v15_ = ((d_462_v16_)[((self).f12 if (self).f15 else (self).f14)] if (((self).f12 if (self).f15 else (self).f14)) in (d_462_v16_) else d_461_v15_)
                d_463_v17_: _dafny.Seq
                d_463_v17_ = _dafny.SeqWithoutIsStrInference([self.f24, self.f24, self.f24, self.f24, self.f24])
                d_464_v18_: _dafny.Array
                nw65_ = _dafny.Array(None, 9)
                nw65_[int(0)] = (d_463_v17_)[default__.safeIndex(p1, len(d_463_v17_))]
                nw65_[int(1)] = self.f24
                nw65_[int(2)] = self.f24
                nw65_[int(3)] = self.f24
                nw65_[int(4)] = self.f24
                nw65_[int(5)] = self.f24
                nw65_[int(6)] = self.f24
                nw65_[int(7)] = self.f24
                nw65_[int(8)] = self.f24
                d_464_v18_ = nw65_
                index71_ = default__.safeIndex(942, (d_464_v18_).length(0))
                (d_464_v18_)[index71_] = self.f24
                index72_ = default__.safeIndex(942, (d_464_v18_).length(0))
                (d_464_v18_)[index72_] = self.f24
                d_465_v19_: _dafny.Set
                d_465_v19_ = _dafny.Set({True})
                r2 = (0) - (default__.safeModuloInt((self).f14, (len(d_465_v19_)) + ((0) - (d_440_i0_))))
                d_466_v20_: _dafny.Map
                d_466_v20_ = _dafny.Map({d_441_v0_: p0})
                d_467_v21_: _dafny.MultiSet
                d_467_v21_ = _dafny.MultiSet([((d_466_v20_)[d_441_v0_] if (d_441_v0_) in (d_466_v20_) else p0), (self).f13, (self).f13, (self).f13])
                d_468_v22_: _dafny.Seq
                d_468_v22_ = _dafny.SeqWithoutIsStrInference([(self).f15, p2])
                d_469_v23_: _dafny.Map
                d_469_v23_ = _dafny.Map({(d_464_v18_)[default__.safeIndex(942, (d_464_v18_).length(0))]: (self).f15})
                d_470_v24_: _dafny.Map
                d_470_v24_ = _dafny.Map({d_468_v22_: d_469_v23_})
                d_471_v25_: _dafny.Seq
                d_471_v25_ = _dafny.SeqWithoutIsStrInference([len(((d_470_v24_)[_dafny.SeqWithoutIsStrInference([p2])] if (_dafny.SeqWithoutIsStrInference([p2])) in (d_470_v24_) else d_469_v23_)), default__.fm0(p2, globalState), ((self).f12) + ((self).f14)])
                rhs58_ = len(d_471_v25_)
                rhs59_ = d_467_v21_
                lhs51_ = globalState
                lhs51_.f6 = rhs58_
                d_467_v21_ = rhs59_
            if not(not ((self).f15) or ((p2 if p2 else p2))):
                d_472_v26_: D2
                d_472_v26_ = D2_DC6((self).f13)
                d_473_v27_: _dafny.Seq
                d_473_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sladrwsg"))
                d_474_v28_: C0
                nw66_ = C0()
                nw66_.ctor__()
                d_474_v28_ = nw66_
                d_475_v29_: _dafny.Map
                d_475_v29_ = _dafny.Map({d_474_v28_: _dafny.Map({p2: (self).f15})})
                d_476_v30_: _dafny.Map
                d_476_v30_ = _dafny.Map({(self).f15: (self).f15})
                d_477_v31_: D2
                d_477_v31_ = D2_DC8(p2, p2, p0)
                d_478_v32_: str
                d_478_v32_ = _dafny.CodePoint('u')
                rhs60_ = (self).f15
                rhs61_ = (self).f15
                rhs62_ = (self).f15
                rhs63_ = d_472_v26_
                rhs64_ = (p0).set(default__.safeIndex(len((((d_475_v29_)[d_474_v28_] if (d_474_v28_) in (d_475_v29_) else d_476_v30_)).set((d_477_v31_).cf15, (self).f15)), len(p0)), d_478_v32_)
                lhs52_ = globalState
                lhs53_ = globalState
                lhs52_.f7 = rhs60_
                r0 = rhs61_
                lhs53_.f2 = rhs62_
                d_472_v26_ = rhs63_
                d_473_v27_ = rhs64_
                index73_ = default__.safeIndex(991, (d_441_v0_).length(0))
                (d_441_v0_)[index73_] = default__.fm0((self).f15, globalState)
                index74_ = default__.safeIndex(991, (d_441_v0_).length(0))
                rhs65_ = (self).f14
                rhs66_ = not ((p2 if (self).f15 else p2)) or ((self).f15)
                rhs67_ = ((d_472_v26_).cf9) + (d_473_v27_)
                lhs54_ = d_441_v0_
                lhs55_ = default__.safeIndex(991, (d_441_v0_).length(0))
                lhs56_ = globalState
                lhs54_[lhs55_] = rhs65_
                lhs56_.f2 = rhs66_
                d_473_v27_ = rhs67_
                index75_ = default__.safeIndex(991, (d_441_v0_).length(0))
                (d_441_v0_)[index75_] = (d_440_i0_) + (((d_441_v0_)[default__.safeIndex(991, (d_441_v0_).length(0))]) - (p3))
                r0 = (self).f15
                d_478_v32_ = d_478_v32_
            elif True:
                d_479_v33_: _dafny.Seq
                d_479_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfqkk"))
                d_479_v33_ = (self).f13
                d_480_v34_: _dafny.Set
                d_480_v34_ = _dafny.Set({489, (self).f12, (self).f12})
                (globalState).f2 = not((675) != (len((_dafny.Map({d_441_v0_: 286})).set(d_441_v0_, len(d_480_v34_)))))
                d_481_v35_: _dafny.Map
                d_481_v35_ = _dafny.Map({p2: p1})
                d_482_v36_: _dafny.Map
                d_482_v36_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(d_481_v35_)])): len(d_480_v34_)})
                index76_ = default__.safeIndex(479, (d_441_v0_).length(0))
                (d_441_v0_)[index76_] = len((d_482_v36_) | (d_482_v36_))
                index77_ = default__.safeIndex(479, (d_441_v0_).length(0))
                (d_441_v0_)[index77_] = (self).f12
                d_483_v37_: C0
                nw67_ = C0()
                nw67_.ctor__()
                d_483_v37_ = nw67_
                r1 = len(_dafny.SeqWithoutIsStrInference([d_440_i0_ for d_484_i3_ in range(default__.abs(390))]))
        hi2_ = -435
        for d_485_i4_ in range(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_498_i5_ in range(default__.abs(-31))])), hi2_):
            r2 = p3
            d_486_v38_: _dafny.Seq
            d_486_v38_ = _dafny.SeqWithoutIsStrInference([(self).f11])
            d_487_v39_: _dafny.MultiSet
            d_487_v39_ = _dafny.MultiSet([(d_486_v38_) + (d_486_v38_)])
            d_487_v39_ = _dafny.MultiSet([(default__.fm20((self).f15, (self).f15, not(p2), globalState)) + (d_486_v38_)])
            if (self).fm1(((d_486_v38_)[default__.safeIndex((self).f11, len(d_486_v38_))] if (D1_DC4(p2)).cf5 else (self).f12), (-979) == ((self).f12), p2, globalState):
                arr6_ = self.f24
                index78_ = default__.safeIndex(803, (self.f24).length(0))
                arr6_[index78_] = ((self).f15 if not((self).fm1((self).f14, (self).f15, (self).f15, globalState)) else p2)
                arr7_ = self.f24
                index79_ = default__.safeIndex(803, (self.f24).length(0))
                arr7_[index79_] = (((self).f14) + (d_485_i4_)) < ((self).f11)
                arr8_ = self.f24
                index80_ = default__.safeIndex(803, (self.f24).length(0))
                arr8_[index80_] = not (not ((self).fm1((self).f11, (self).f15, (self.f24)[default__.safeIndex(803, (self.f24).length(0))], globalState)) or ((self).f15)) or (p2)
                d_488_v40_: _dafny.Map
                d_488_v40_ = _dafny.Map({(self).f11: (self.f24)[default__.safeIndex(803, (self.f24).length(0))]})
                d_489_v41_: _dafny.Map
                d_489_v41_ = _dafny.Map({len(_dafny.Map({(self).f15: len(d_488_v40_)})): (self).f14})
                d_489_v41_ = (d_489_v41_ if p2 else d_489_v41_)
                d_490_v42_: _dafny.Array
                nw68_ = _dafny.Array(int(0), 13)
                d_490_v42_ = nw68_
                d_491_v43_: _dafny.Set
                d_491_v43_ = _dafny.Set({d_490_v42_, d_490_v42_})
                d_491_v43_ = d_491_v43_
                d_492_v44_: C0
                nw69_ = C0()
                nw69_.ctor__()
                d_492_v44_ = nw69_
            elif True:
                d_493_v45_: _dafny.Set
                d_493_v45_ = _dafny.Set({p3, (0) - (d_485_i4_)})
                (globalState).f6 = (0) - (default__.safeDivisionInt((0) - (len(d_493_v45_)), -77))
                d_494_v46_: _dafny.Set
                d_494_v46_ = _dafny.Set({p2})
                d_495_v47_: _dafny.Array
                d_496_v48_: _dafny.Set
                out25_: _dafny.Array
                out26_: _dafny.Set
                out25_, out26_ = (self).m13(d_494_v46_, globalState)
                d_495_v47_ = out25_
                d_496_v48_ = out26_
                index81_ = default__.safeIndex(925, (d_495_v47_).length(0))
                (d_495_v47_)[index81_] = (self).f14
                index82_ = default__.safeIndex(925, (d_495_v47_).length(0))
                (d_495_v47_)[index82_] = (0) - ((default__.safeDivisionInt(p3, p1)) * ((self).f11))
                d_495_v47_ = d_495_v47_
                d_495_v47_ = d_495_v47_
            d_497_v49_: C0
            nw70_ = C0()
            nw70_.ctor__()
            d_497_v49_ = nw70_
        d_499_v50_: _dafny.MultiSet
        d_499_v50_ = _dafny.MultiSet([p2, (self).f15, (self).f15, not (p2) or (False), not((p2) and (p2))])
        rhs68_ = (self).f11
        rhs69_ = d_499_v50_
        r2 = rhs68_
        d_499_v50_ = rhs69_
        (globalState).f7 = p2
        d_500_v51_: _dafny.Map
        d_500_v51_ = _dafny.Map({(self).f15: p2})
        d_501_v52_: _dafny.Map
        d_501_v52_ = _dafny.Map({p2: 371})
        rhs70_ = p2
        rhs71_ = ((self).f11) < ((len(d_500_v51_)) + ((self).f12))
        rhs72_ = (len(d_501_v52_) if not((self).f15) else (self).f14)
        lhs57_ = globalState
        lhs58_ = globalState
        lhs59_ = globalState
        lhs57_.f7 = rhs70_
        lhs58_.f7 = rhs71_
        lhs59_.f6 = rhs72_
        d_502_v53_: _dafny.Seq
        d_502_v53_ = _dafny.SeqWithoutIsStrInference([(self).f14, 500])
        d_503_v54_: _dafny.MultiSet
        d_503_v54_ = _dafny.MultiSet([d_502_v53_, d_502_v53_, d_502_v53_])
        if (((d_503_v54_)[d_502_v53_] if (d_502_v53_) in (d_503_v54_) else (self).f14)) < ((self).f14):
            arr9_ = self.f24
            index83_ = default__.safeIndex(182, (self.f24).length(0))
            arr9_[index83_] = p2
            d_504_v55_: _dafny.Array
            nw71_ = _dafny.Array(None, 22)
            d_504_v55_ = nw71_
            d_505_v56_: _dafny.Set
            d_505_v56_ = _dafny.Set({d_504_v55_, d_504_v55_})
            d_506_v57_: D1
            d_506_v57_ = D1_DC2(default__.fm21(p1, (self).f15, globalState))
            d_507_v58_: D2
            d_507_v58_ = D2_DC7(self.f24, d_505_v56_, d_506_v57_, (self).f12, (self).f12)
            d_508_v59_: _dafny.MultiSet
            d_508_v59_ = _dafny.MultiSet([(d_507_v58_).cf10, self.f24, self.f24, self.f24])
            arr10_ = self.f24
            index84_ = default__.safeIndex(182, (self.f24).length(0))
            arr10_[index84_] = (d_508_v59_).issubset(d_508_v59_)
            d_509_v60_: _dafny.Map
            d_509_v60_ = _dafny.Map({p3: p2})
            d_509_v60_ = d_509_v60_
            d_510_v61_: str
            d_510_v61_ = _dafny.CodePoint('k')
            d_500_v51_ = (d_500_v51_).set((d_510_v61_) in (p0), p2)
            if (self).f15:
                d_511_v62_: _dafny.Map
                d_511_v62_ = _dafny.Map({D1_DC4(p2): (self).f14})
                rhs73_ = p2
                rhs74_ = default__.fm0((self).f15, globalState)
                rhs75_ = not ((self).f15) or ((self).fm1((self).f11, (self.f24)[default__.safeIndex(182, (self.f24).length(0))], True, globalState))
                rhs76_ = len(d_511_v62_)
                rhs77_ = p2
                lhs60_ = globalState
                lhs61_ = globalState
                r0 = rhs73_
                r2 = rhs74_
                lhs60_.f2 = rhs75_
                r2 = rhs76_
                lhs61_.f2 = rhs77_
                (globalState).f6 = (self).f12
                d_512_v63_: _dafny.Seq
                d_512_v63_ = _dafny.SeqWithoutIsStrInference([((d_500_v51_)[False] if (False) in (d_500_v51_) else (self).f15), p2])
                d_513_v64_: _dafny.MultiSet
                d_513_v64_ = _dafny.MultiSet([43, (self).f14])
                arr11_ = self.f24
                index85_ = default__.safeIndex(182, (self.f24).length(0))
                arr11_[index85_] = (d_512_v63_)[default__.safeIndex(((d_513_v64_)[(0) - (len((self).f13))] if ((0) - (len((self).f13))) in (d_513_v64_) else -283), len(d_512_v63_))]
                d_514_v65_: D6
                d_514_v65_ = D6_DC20((self).f11, default__.fm0((self.f24)[default__.safeIndex(182, (self.f24).length(0))], globalState))
                d_515_v66_: D6
                d_515_v66_ = D6_DC22(d_514_v65_)
                d_516_v67_: _dafny.Map
                d_516_v67_ = _dafny.Map({(self).f15: d_515_v66_})
                rhs78_ = ((d_516_v67_) | (d_516_v67_)) | (d_516_v67_)
                rhs79_ = (self.f24)[default__.safeIndex(182, (self.f24).length(0))]
                lhs62_ = globalState
                d_516_v67_ = rhs78_
                lhs62_.f2 = rhs79_
                d_517_v68_: _dafny.Seq
                d_517_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcndqglo"))
                d_517_v68_ = (((self).f13) + (p0)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdyux")))
            elif True:
                d_518_v69_: D4
                d_518_v69_ = D4_DC13(d_502_v53_)
                d_519_v71_: _dafny.Set
                d_519_v71_ = _dafny.Set({p3})
                d_520_v72_: _dafny.Seq
                d_520_v72_ = _dafny.SeqWithoutIsStrInference([d_519_v71_])
                arr12_ = self.f24
                index86_ = default__.safeIndex(182, (self.f24).length(0))
                def iife34_():
                    coll22_ = _dafny.Set()
                    compr_22_: int
                    for compr_22_ in ((d_518_v69_).cf29).Elements:
                        d_521_v70_: int = compr_22_
                        if (d_521_v70_) in ((d_518_v69_).cf29):
                            coll22_ = coll22_.union(_dafny.Set([default__.safeDivisionInt(d_521_v70_, 870)]))
                    return _dafny.Set(coll22_)
                arr12_[index86_] = ((d_520_v72_)[default__.safeIndex(p3, len(d_520_v72_))]).ispropersubset(iife34_()
                )
                d_522_v73_: _dafny.Map
                d_522_v73_ = _dafny.Map({(self).f14: 280})
                d_522_v73_ = (d_522_v73_).set(default__.safeModuloInt((self).fm2(p3, False, p1, (self).f15, globalState), (d_502_v53_)[default__.safeIndex((0) - ((self).f11), len(d_502_v53_))]), p1)
                (globalState).f6 = p3
                d_523_v74_: _dafny.Array
                nw72_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_523_v74_ = nw72_
                d_524_v75_: _dafny.Array
                def lambda20_(d_525_v61_, d_526_p1_):
                    def lambda21_(d_527_i6_):
                        return D4_DC14(d_525_v61_, d_526_p1_)

                    return lambda21_

                init9_ = lambda20_(d_510_v61_, p1)
                nw73_ = _dafny.Array(None, 22)
                for i0_9_ in range(nw73_.length(0)):
                    nw73_[i0_9_] = init9_(i0_9_)
                d_524_v75_ = nw73_
                index87_ = default__.safeIndex(835, (d_523_v74_).length(0))
                (d_523_v74_)[index87_] = d_524_v75_
                index88_ = default__.safeIndex(835, (d_523_v74_).length(0))
                (d_523_v74_)[index88_] = d_524_v75_
                arr13_ = self.f24
                index89_ = default__.safeIndex(182, (self.f24).length(0))
                arr13_[index89_] = (self.f24)[default__.safeIndex(182, (self.f24).length(0))]
            d_528_v76_: _dafny.Array
            nw74_ = _dafny.Array(int(0), 23)
            d_528_v76_ = nw74_
            d_528_v76_ = d_528_v76_
        elif True:
            (self).f24 = self.f24
            (globalState).f6 = p1
            d_529_v77_: D4
            d_529_v77_ = D4_DC15(p3)
            d_530_v78_: str
            d_530_v78_ = _dafny.CodePoint('l')
            d_531_v79_: _dafny.Map
            d_531_v79_ = _dafny.Map({(p0) + ((p0).set(default__.safeIndex((d_529_v77_).cf32, len(p0)), d_530_v78_)): (self).f14})
            rhs80_ = self.f24
            rhs81_ = ((d_531_v79_)[(self).f13] if ((self).f13) in (d_531_v79_) else p3)
            lhs63_ = self
            lhs63_.f24 = rhs80_
            r2 = rhs81_
            if not (True) or ((self).f15):
                (globalState).f7 = False
                d_532_v80_: C0
                nw75_ = C0()
                nw75_.ctor__()
                d_532_v80_ = nw75_
                d_533_v81_: _dafny.Array
                nw76_ = _dafny.Array(int(0), 14)
                d_533_v81_ = nw76_
                d_534_v82_: _dafny.Set
                d_534_v82_ = _dafny.Set({(d_502_v53_)[default__.safeIndex(default__.fm0((self).f15, globalState), len(d_502_v53_))]})
                index90_ = default__.safeIndex(274, (d_533_v81_).length(0))
                (d_533_v81_)[index90_] = (0) - (default__.safeModuloInt((0) - (len(d_534_v82_)), 738))
                index91_ = default__.safeIndex(274, (d_533_v81_).length(0))
                rhs82_ = p3
                rhs83_ = (d_532_v80_).fm16(p2, p1, (self).f11, p1, globalState)
                lhs64_ = globalState
                lhs65_ = d_533_v81_
                lhs66_ = default__.safeIndex(274, (d_533_v81_).length(0))
                lhs64_.f6 = rhs82_
                lhs65_[lhs66_] = rhs83_
                def lambda22_(d_535_p2_):
                    def lambda23_(d_536_i7_):
                        return d_535_p2_

                    return lambda23_

                init10_ = lambda22_(p2)
                nw77_ = _dafny.Array(None, 3)
                for i0_10_ in range(nw77_.length(0)):
                    nw77_[i0_10_] = init10_(i0_10_)
                (self).f24 = nw77_
                d_537_v83_: _dafny.Map
                d_537_v83_ = _dafny.Map({(self).f13: (self).f15})
                d_538_v84_: _dafny.Map
                d_538_v84_ = _dafny.Map({len(_dafny.Map({(self).f15: d_502_v53_})): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcj"))})
                d_537_v83_ = (d_537_v83_).set((p0) + (((d_538_v84_)[(self).f12] if ((self).f12) in (d_538_v84_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juhvefvw")))), not((_dafny.SeqWithoutIsStrInference([d_530_v78_ for d_539_i8_ in range(default__.abs(788))])) < ((self).f13)))
            elif True:
                d_540_v85_: _dafny.Map
                d_540_v85_ = _dafny.Map({(self).f12: (self).f11})
                d_541_v86_: _dafny.Map
                d_541_v86_ = _dafny.Map({((d_540_v85_)[(self).f12] if ((self).f12) in (d_540_v85_) else (self).f14): (0) - (len(_dafny.SeqWithoutIsStrInference([True])))})
                d_542_v87_: _dafny.Map
                d_542_v87_ = _dafny.Map({d_530_v78_: (self).f15})
                d_543_v88_: _dafny.Set
                d_543_v88_ = _dafny.Set({774, len(d_542_v87_), (self).f11})
                d_544_v89_: _dafny.Map
                d_544_v89_ = _dafny.Map({d_543_v88_: (self).f11})
                d_545_v90_: C0
                nw78_ = C0()
                nw78_.ctor__()
                d_545_v90_ = nw78_
                d_546_v91_: _dafny.Map
                d_546_v91_ = _dafny.Map({d_545_v90_: (self).f13})
                d_541_v86_ = _dafny.Map({len((d_544_v89_) | (d_544_v89_)): (0) - ((default__.fm0((self).f15, globalState) if p2 else len(d_546_v91_)))})
                d_547_v92_: C0
                nw79_ = C0()
                nw79_.ctor__()
                d_547_v92_ = nw79_
                d_548_v93_: _dafny.Array
                nw80_ = _dafny.Array(int(0), 9)
                d_548_v93_ = nw80_
                d_549_v94_: _dafny.Map
                d_549_v94_ = _dafny.Map({d_548_v93_: (self).f11})
                arr14_ = self.f24
                index92_ = default__.safeIndex(235, (self.f24).length(0))
                arr14_[index92_] = (d_548_v93_) not in (d_549_v94_)
                arr15_ = self.f24
                index93_ = default__.safeIndex(235, (self.f24).length(0))
                arr15_[index93_] = (d_545_v90_).fm17((default__.fm20(p2, (self).f15, False, globalState)) + (d_502_v53_), globalState)
                d_550_v95_: _dafny.Seq
                d_550_v95_ = _dafny.SeqWithoutIsStrInference([(self.f24)[default__.safeIndex(235, (self.f24).length(0))]])
                (globalState).f6 = ((0) - ((self).f12) if (self).f15 else (0) - ((len(d_550_v95_)) + ((self).f14)))
                d_551_v96_: _dafny.Array
                nw81_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_551_v96_ = nw81_
                d_552_v97_: _dafny.Array
                nw82_ = _dafny.Array(_dafny.Map({}), 21)
                d_552_v97_ = nw82_
                index94_ = default__.safeIndex(715, (d_551_v96_).length(0))
                (d_551_v96_)[index94_] = d_552_v97_
                index95_ = default__.safeIndex(715, (d_551_v96_).length(0))
                (d_551_v96_)[index95_] = d_552_v97_
            (globalState).f7 = not(p2)
        r0 = (self).f15
        d_553_v98_: _dafny.MultiSet
        d_553_v98_ = _dafny.MultiSet([default__.fm0((self).f15, globalState), 66, p1, (self).f11])
        r1 = ((d_553_v98_)[default__.fm0(False, globalState)] if (default__.fm0(False, globalState)) in (d_553_v98_) else (self).f14)
        r2 = default__.safeModuloInt(default__.fm0((self).f15, globalState), p3)
        return r0, r1, r2

    def m2(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        d_554_i0_: int
        d_554_i0_ = 0
        with _dafny.label("3"):
            while (self).f15:
                with _dafny.c_label("3"):
                    if (d_554_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_554_i0_ = (d_554_i0_) + (1)
                    (globalState).f6 = (0) - (default__.safeModuloInt(((self).f14) + ((self).f14), 557))
                    (globalState).f6 = (self).f14
                    d_555_v0_: _dafny.Map
                    d_555_v0_ = _dafny.Map({self.f24: (self).f11})
                    d_555_v0_ = d_555_v0_
                    (globalState).f6 = (0) - (default__.safeDivisionInt((self).f11, (self).f12))
                    pass
            pass
        if (self).f15:
            d_556_v1_: C0
            nw83_ = C0()
            nw83_.ctor__()
            d_556_v1_ = nw83_
            arr16_ = self.f24
            index96_ = default__.safeIndex(246, (self.f24).length(0))
            arr16_[index96_] = ((self).f15 if (self).f15 else False)
            arr17_ = self.f24
            index97_ = default__.safeIndex(246, (self.f24).length(0))
            arr17_[index97_] = ((self).f14) > ((self).f11)
            if not((self).fm1((self).f11, (self.f24)[default__.safeIndex(246, (self.f24).length(0))], (self.f24)[default__.safeIndex(246, (self.f24).length(0))], globalState)):
                d_557_v2_: _dafny.Seq
                d_557_v2_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                d_558_v3_: _dafny.Seq
                d_558_v3_ = _dafny.SeqWithoutIsStrInference([len(d_557_v2_)])
                d_559_v4_: str
                d_559_v4_ = _dafny.CodePoint('a')
                arr18_ = self.f24
                index98_ = default__.safeIndex(246, (self.f24).length(0))
                arr18_[index98_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))).set(default__.safeIndex(len(d_558_v3_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")))), d_559_v4_)) < ((self).f13)
                (globalState).f6 = (self).f12
                d_560_v5_: _dafny.Set
                d_560_v5_ = _dafny.Set({(self.f24)[default__.safeIndex(246, (self.f24).length(0))], True})
                d_561_v6_: _dafny.Array
                d_562_v7_: _dafny.Set
                out27_: _dafny.Array
                out28_: _dafny.Set
                out27_, out28_ = (self).m13(d_560_v5_, globalState)
                d_561_v6_ = out27_
                d_562_v7_ = out28_
                (globalState).f6 = (default__.safeModuloInt((self).f14, (self).f11)) + (len(((self).f13) + ((self).f13)))
                d_563_v8_: C0
                nw84_ = C0()
                nw84_.ctor__()
                d_563_v8_ = nw84_
            elif True:
                (globalState).f6 = 951
                d_564_v9_: _dafny.Array
                nw85_ = _dafny.Array(_dafny.MultiSet({}), 19)
                d_564_v9_ = nw85_
                d_564_v9_ = d_564_v9_
                d_565_v10_: _dafny.MultiSet
                d_565_v10_ = _dafny.MultiSet([True])
                d_566_v11_: _dafny.Seq
                d_566_v11_ = _dafny.SeqWithoutIsStrInference([(self).f11, (d_565_v10_).cardinality])
                (globalState).f7 = (d_556_v1_).fm17((d_566_v11_) + (_dafny.SeqWithoutIsStrInference([(self).f11, 981])), globalState)
                d_567_v12_: _dafny.Array
                nw86_ = _dafny.Array(False, 5)
                d_567_v12_ = nw86_
                d_568_v13_: _dafny.MultiSet
                d_568_v13_ = _dafny.MultiSet([p0, p0, d_567_v12_])
                d_569_v14_: str
                d_569_v14_ = _dafny.CodePoint('q')
                d_570_v15_: _dafny.MultiSet
                d_570_v15_ = _dafny.MultiSet([d_569_v14_, d_569_v14_, default__.fm21((self).f14, (self).f15, globalState)])
                d_571_v16_: _dafny.Set
                d_571_v16_ = _dafny.Set({((self).f14) * (len(_dafny.SeqWithoutIsStrInference([285, (self).f14, (self).f11, ((d_570_v15_)[d_569_v14_] if (d_569_v14_) in (d_570_v15_) else (self).f12), (self).f14]))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryfouad"))), ((self).f11 if (d_556_v1_).fm17(_dafny.SeqWithoutIsStrInference([(self).f12]), globalState) else (self).f14), ((self).f14) + ((self).f12), (self).f12})
                rhs84_ = p0
                rhs85_ = d_568_v13_
                rhs86_ = d_571_v16_
                d_567_v12_ = rhs84_
                d_568_v13_ = rhs85_
                d_571_v16_ = rhs86_
                d_572_v17_: _dafny.Seq
                d_572_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyja"))
                d_573_v18_: _dafny.Seq
                d_573_v18_ = _dafny.SeqWithoutIsStrInference([((self).f13).set(default__.safeIndex((self).f14, len((self).f13)), d_569_v14_)])
                d_572_v17_ = ((d_573_v18_)[default__.safeIndex((self).f14, len(d_573_v18_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkid")))
            d_574_v19_: str
            d_574_v19_ = _dafny.CodePoint('n')
            d_575_v20_: _dafny.Array
            nw87_ = _dafny.Array(None, 13)
            nw87_[int(0)] = d_574_v19_
            nw87_[int(1)] = d_574_v19_
            nw87_[int(2)] = d_574_v19_
            nw87_[int(3)] = d_574_v19_
            nw87_[int(4)] = d_574_v19_
            nw87_[int(5)] = d_574_v19_
            nw87_[int(6)] = (d_574_v19_ if (self).f15 else d_574_v19_)
            nw87_[int(7)] = d_574_v19_
            nw87_[int(8)] = _dafny.CodePoint('p')
            nw87_[int(9)] = d_574_v19_
            nw87_[int(10)] = d_574_v19_
            nw87_[int(11)] = (d_574_v19_ if (self).f15 else d_574_v19_)
            nw87_[int(12)] = d_574_v19_
            d_575_v20_ = nw87_
            index99_ = default__.safeIndex(605, (d_575_v20_).length(0))
            (d_575_v20_)[index99_] = d_574_v19_
            index100_ = default__.safeIndex(605, (d_575_v20_).length(0))
            (d_575_v20_)[index100_] = d_574_v19_
            index101_ = default__.safeIndex(605, (d_575_v20_).length(0))
            (d_575_v20_)[index101_] = (d_575_v20_)[default__.safeIndex(605, (d_575_v20_).length(0))]
        elif True:
            d_576_v21_: _dafny.Map
            d_576_v21_ = _dafny.Map({(self).f14: (self).f11})
            d_577_v22_: _dafny.Map
            d_577_v22_ = _dafny.Map({(self).f15: (self).f13})
            d_578_v23_: str
            d_578_v23_ = _dafny.CodePoint('h')
            d_579_v24_: _dafny.Map
            d_579_v24_ = _dafny.Map({len(((self).f13).set(default__.safeIndex((_dafny.MultiSet(default__.fm22((self).f14, globalState))).cardinality, len((self).f13)), d_578_v23_)): (self).f13})
            d_580_v26_: _dafny.Map
            d_580_v26_ = _dafny.Map({(self).f15: (self).f15})
            d_581_v27_: D0
            d_581_v27_ = D0_DC1(((d_580_v26_)[(self).f15] if ((self).f15) in (d_580_v26_) else (self).f15), (self).f12)
            d_582_v28_: _dafny.Array
            nw88_ = _dafny.Array(None, 28)
            nw88_[int(0)] = (self).f12
            nw88_[int(1)] = (self).f12
            nw88_[int(2)] = ((self).f14) + ((self).f11)
            nw88_[int(3)] = len((self).f13)
            nw88_[int(4)] = ((d_576_v21_)[(self).f11] if ((self).f11) in (d_576_v21_) else (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avvm")))))
            nw88_[int(5)] = (self).f12
            nw88_[int(6)] = (self).f14
            nw88_[int(7)] = (self).f12
            nw88_[int(8)] = (self).f12
            nw88_[int(9)] = default__.fm0((self).f15, globalState)
            nw88_[int(10)] = (self).f14
            nw88_[int(11)] = 74
            nw88_[int(12)] = (0) - ((self).f11)
            nw88_[int(13)] = (len(d_577_v22_)) - (len(d_579_v24_))
            nw88_[int(14)] = (self).f12
            nw88_[int(15)] = (self).f14
            nw88_[int(16)] = (self).f14
            nw88_[int(17)] = (self).f14
            nw88_[int(18)] = (self).f12
            nw88_[int(19)] = ((self).f14) + (default__.fm0(True, globalState))
            def iife35_():
                coll23_ = _dafny.Set()
                compr_23_: int
                for compr_23_ in (_dafny.Map({(self).f12: (self).f13})).keys.Elements:
                    d_583_v25_: int = compr_23_
                    if (d_583_v25_) in (_dafny.Map({(self).f12: (self).f13})):
                        coll23_ = coll23_.union(_dafny.Set([default__.safeModuloInt(d_583_v25_, 443)]))
                return _dafny.Set(coll23_)
            nw88_[int(20)] = len(iife35_()
            )
            nw88_[int(21)] = (self).f12
            nw88_[int(22)] = len(((d_577_v22_)[(self).f15] if ((self).f15) in (d_577_v22_) else (self).f13))
            nw88_[int(23)] = (657 if (self).f15 else (self).f11)
            nw88_[int(24)] = default__.safeModuloInt((self).f12, (self).f11)
            nw88_[int(25)] = default__.safeDivisionInt(len(_dafny.Map({(self).f14: (self).f15})), (self).f11)
            nw88_[int(26)] = (d_581_v27_).cf2
            nw88_[int(27)] = (self).f11
            d_582_v28_ = nw88_
            index102_ = default__.safeIndex(177, (d_582_v28_).length(0))
            (d_582_v28_)[index102_] = (self).f14
            d_584_v29_: _dafny.Map
            d_584_v29_ = _dafny.Map({(self).f11: (self).f15})
            d_585_v30_: D6
            d_585_v30_ = D6_DC20((self).f14, len((self).f13))
            d_586_v31_: _dafny.Map
            d_586_v31_ = _dafny.Map({(self).f13: (self).f15})
            pat_let_tv6_ = d_586_v31_
            d_587_v32_: _dafny.Map
            def iife36_(_pat_let6_0):
                def iife37_(d_588_dt__update__tmp_h0_):
                    def iife38_(_pat_let7_0):
                        def iife39_(d_589_dt__update_hcf39_h0_):
                            return D6_DC20((d_588_dt__update__tmp_h0_).cf38, d_589_dt__update_hcf39_h0_)
                        return iife39_(_pat_let7_0)
                    return iife38_((0) - (len(pat_let_tv6_)))
                return iife37_(_pat_let6_0)
            d_587_v32_ = _dafny.Map({(self).f13: (iife36_(d_585_v30_)).cf38})
            d_590_v33_: _dafny.Map
            d_590_v33_ = _dafny.Map({len(d_584_v29_): (len(((self).f13).set(default__.safeIndex((p1).cardinality, len((self).f13)), d_578_v23_))) == (len(d_587_v32_))})
            d_591_v34_: _dafny.Seq
            d_591_v34_ = _dafny.SeqWithoutIsStrInference([(self).f14])
            d_592_v35_: D4
            d_592_v35_ = D4_DC13(d_591_v34_)
            d_593_v36_: _dafny.Map
            d_593_v36_ = _dafny.Map({not(False): (self).f12})
            d_594_v37_: _dafny.Array
            nw89_ = _dafny.Array(None, 9)
            nw89_[int(0)] = (d_591_v34_) + (d_591_v34_)
            nw89_[int(1)] = (d_591_v34_).set(default__.safeIndex((self).f14, len(d_591_v34_)), (self).fm2((self).f12, (self).f15, 959, False, globalState))
            nw89_[int(2)] = d_591_v34_
            nw89_[int(3)] = (d_591_v34_) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(D1_DC4((self).f15)).cf5: (self).f15})) for d_595_i1_ in range(default__.abs(42))]))
            nw89_[int(4)] = _dafny.SeqWithoutIsStrInference([((d_593_v36_)[(self).f15] if ((self).f15) in (d_593_v36_) else (self).f12), (self).f12])
            nw89_[int(5)] = d_591_v34_
            nw89_[int(6)] = d_591_v34_
            nw89_[int(7)] = d_591_v34_
            nw89_[int(8)] = default__.fm20(False, False, (self).f15, globalState)
            d_594_v37_ = nw89_
            index103_ = default__.safeIndex(689, (d_594_v37_).length(0))
            (d_594_v37_)[index103_] = _dafny.SeqWithoutIsStrInference([default__.fm0((self).f15, globalState)])
            index104_ = default__.safeIndex(177, (d_582_v28_).length(0))
            index105_ = default__.safeIndex(689, (d_594_v37_).length(0))
            rhs87_ = (self).f11
            rhs88_ = d_590_v33_
            rhs89_ = d_592_v35_
            rhs90_ = (self).f14
            rhs91_ = d_591_v34_
            lhs67_ = d_582_v28_
            lhs68_ = default__.safeIndex(177, (d_582_v28_).length(0))
            lhs69_ = globalState
            lhs70_ = d_594_v37_
            lhs71_ = default__.safeIndex(689, (d_594_v37_).length(0))
            lhs67_[lhs68_] = rhs87_
            d_584_v29_ = rhs88_
            d_592_v35_ = rhs89_
            lhs69_.f6 = rhs90_
            lhs70_[lhs71_] = rhs91_
            d_596_v38_: _dafny.Seq
            d_596_v38_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            d_597_v39_: _dafny.Map
            d_597_v39_ = _dafny.Map({(d_596_v38_)[default__.safeIndex((self).f11, len(d_596_v38_))]: d_578_v23_})
            d_597_v39_ = (d_597_v39_).set((self).f15, d_578_v23_)
            (globalState).f6 = (self).f12
            d_598_v40_: _dafny.Array
            nw90_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_598_v40_ = nw90_
            rhs92_ = (self).f15
            rhs93_ = (self).f15
            rhs94_ = d_598_v40_
            rhs95_ = (self).f15
            lhs72_ = globalState
            lhs73_ = globalState
            lhs72_.f2 = rhs92_
            r0 = rhs93_
            d_598_v40_ = rhs94_
            lhs73_.f7 = rhs95_
            d_580_v26_ = (d_580_v26_).set((self).f15, ((self).f12) >= ((self).f14))
        d_599_v41_: _dafny.Seq
        d_599_v41_ = _dafny.SeqWithoutIsStrInference([(self).f13])
        d_600_v42_: _dafny.Seq
        d_600_v42_ = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15, (self).f15])
        d_601_v43_: _dafny.MultiSet
        d_601_v43_ = _dafny.MultiSet([(d_600_v42_)[default__.safeIndex((self).f11, len(d_600_v42_))], False])
        d_602_v44_: _dafny.Set
        d_602_v44_ = _dafny.Set({(self).f15, (self).f15, (self).f15, (self).f15, (self).f15})
        d_603_v45_: _dafny.Map
        d_603_v45_ = _dafny.Map({(self).f15: (self).f12})
        d_604_v46_: _dafny.Seq
        d_604_v46_ = _dafny.SeqWithoutIsStrInference([(self).f11])
        d_605_v47_: _dafny.Map
        d_605_v47_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([831]): len(d_604_v46_)})
        d_606_v49_: _dafny.Map
        d_606_v49_ = _dafny.Map({(self).f15: not((self).fm1((self).f12, (self).f15, False, globalState))})
        d_607_v50_: _dafny.Array
        nw91_ = _dafny.Array(None, 29)
        nw91_[int(0)] = (self).f12
        nw91_[int(1)] = (0) - (len((d_599_v41_)[default__.safeIndex((d_601_v43_).cardinality, len(d_599_v41_))]))
        nw91_[int(2)] = -947
        nw91_[int(3)] = len(d_602_v44_)
        nw91_[int(4)] = (self).f12
        nw91_[int(5)] = len(d_602_v44_)
        nw91_[int(6)] = (D0_DC1(not(False), len(_dafny.SeqWithoutIsStrInference([(self).f15])))).cf2
        nw91_[int(7)] = (self).f11
        nw91_[int(8)] = (self).f11
        nw91_[int(9)] = (self).f11
        nw91_[int(10)] = len(_dafny.SeqWithoutIsStrInference([(self).fm2((0) - ((self).f12), False, len(d_600_v42_), (self).f15, globalState), (self).f12, (self).f12]))
        nw91_[int(11)] = len(d_603_v45_)
        nw91_[int(12)] = (0) - (((self).f11) * ((self).f14))
        nw91_[int(13)] = default__.safeModuloInt(((d_605_v47_)[_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f11, (0) - ((self).f12)])) for d_608_i2_ in range(default__.abs(255))])] if (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f11, (0) - ((self).f12)])) for d_609_i2_ in range(default__.abs(255))])) in (d_605_v47_) else (self).f12), (self).f11)
        nw91_[int(14)] = default__.safeDivisionInt((self).f14, 871)
        nw91_[int(15)] = (0) - (((self).f11) * ((self).f11))
        nw91_[int(16)] = default__.safeModuloInt(len(d_604_v46_), (self).f14)
        nw91_[int(17)] = (self).f14
        nw91_[int(18)] = (self).f14
        nw91_[int(19)] = default__.safeDivisionInt((self).f14, 992)
        nw91_[int(20)] = (self).f11
        nw91_[int(21)] = (self).f12
        nw91_[int(22)] = (self).f14
        nw91_[int(23)] = 990
        nw91_[int(24)] = (self).f11
        nw91_[int(25)] = (self).f14
        def iife40_():
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in _dafny.IntegerRange(527, 302):
                d_610_v48_: int = compr_24_
                if ((527) <= (d_610_v48_)) and ((d_610_v48_) < (302)):
                    coll24_[default__.safeModuloInt(d_610_v48_, (self).f12)] = (self).f11
            return _dafny.Map(coll24_)
        nw91_[int(26)] = (len(iife40_()
        )) + ((self).f11)
        nw91_[int(27)] = (len(d_606_v49_)) + ((self).f14)
        nw91_[int(28)] = (_dafny.MultiSet(d_600_v42_)).cardinality
        d_607_v50_ = nw91_
        d_607_v50_ = d_607_v50_
        d_611_v51_: _dafny.Map
        d_611_v51_ = _dafny.Map({d_602_v44_: (self).f15})
        d_611_v51_ = (d_611_v51_).set(_dafny.Set({not((self).f15), (self).f15}), (not((self).f15) if (d_600_v42_)[default__.safeIndex((self).f14, len(d_600_v42_))] else True))
        d_612_v52_: D2
        d_612_v52_ = D2_DC6(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_613_i3_ in range(default__.abs(316))]))
        d_614_v53_: _dafny.Seq
        d_614_v53_ = _dafny.SeqWithoutIsStrInference([d_612_v52_])
        d_615_v54_: _dafny.MultiSet
        d_615_v54_ = _dafny.MultiSet([d_614_v53_, d_614_v53_])
        d_616_v55_: _dafny.Map
        d_616_v55_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([((d_615_v54_)[_dafny.SeqWithoutIsStrInference([D2_DC6((self).f13) for d_617_i4_ in range(default__.abs(88))])] if (_dafny.SeqWithoutIsStrInference([D2_DC6((self).f13) for d_618_i4_ in range(default__.abs(88))])) in (d_615_v54_) else (self).f12), (self).f12]): d_607_v50_})
        d_619_v56_: _dafny.Map
        d_619_v56_ = _dafny.Map({(d_616_v55_) | (d_616_v55_): (self).f14})
        d_620_v57_: _dafny.Map
        d_620_v57_ = _dafny.Map({(self).f12: (self).f15})
        (globalState).f6 = ((d_619_v56_)[d_616_v55_] if (d_616_v55_) in (d_619_v56_) else (0) - (len(d_620_v57_)))
        d_621_v58_: _dafny.Array
        nw92_ = _dafny.Array(_dafny.Seq({}), 18)
        d_621_v58_ = nw92_
        index106_ = default__.safeIndex(78, (d_621_v58_).length(0))
        (d_621_v58_)[index106_] = (d_600_v42_).set(default__.safeIndex((self).f11, len(d_600_v42_)), (self).f15)
        index107_ = default__.safeIndex(78, (d_621_v58_).length(0))
        (d_621_v58_)[index107_] = d_600_v42_
        r0 = ((self).f15) == (not((self).f15))
        r1 = p1
        return r0, r1

    def m13(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Set = _dafny.Set({})
        (globalState).f6 = (0) - ((0) - ((self).f14))
        d_622_v0_: C0
        nw93_ = C0()
        nw93_.ctor__()
        d_622_v0_ = nw93_
        d_623_v1_: _dafny.Array
        def lambda24_(d_624_i0_):
            return (d_624_i0_) - (475)

        init11_ = lambda24_
        nw94_ = _dafny.Array(None, 25)
        for i0_11_ in range(nw94_.length(0)):
            nw94_[i0_11_] = init11_(i0_11_)
        d_623_v1_ = nw94_
        d_625_v2_: _dafny.Array
        d_625_v2_ = d_623_v1_
        d_626_v3_: _dafny.Array
        nw95_ = _dafny.Array(None, 19)
        nw95_[int(0)] = d_623_v1_
        nw95_[int(1)] = d_623_v1_
        nw95_[int(2)] = d_623_v1_
        nw95_[int(3)] = d_623_v1_
        nw95_[int(4)] = d_623_v1_
        nw95_[int(5)] = d_623_v1_
        nw95_[int(6)] = (d_625_v2_)
        nw95_[int(7)] = d_623_v1_
        nw95_[int(8)] = d_623_v1_
        nw95_[int(9)] = d_623_v1_
        nw95_[int(10)] = d_623_v1_
        nw95_[int(11)] = d_623_v1_
        nw95_[int(12)] = d_623_v1_
        nw95_[int(13)] = d_623_v1_
        nw95_[int(14)] = d_623_v1_
        nw95_[int(15)] = d_623_v1_
        nw95_[int(16)] = d_623_v1_
        nw95_[int(17)] = d_623_v1_
        nw95_[int(18)] = d_623_v1_
        d_626_v3_ = nw95_
        d_626_v3_ = (d_626_v3_ if ((self).f14) <= ((self).f11) else d_626_v3_)
        arr19_ = self.f24
        index108_ = default__.safeIndex(136, (self.f24).length(0))
        arr19_[index108_] = ((self).f15 if (self).f15 else (self).f15)
        arr20_ = self.f24
        index109_ = default__.safeIndex(136, (self.f24).length(0))
        arr20_[index109_] = (self).f15
        hi3_ = (self).f14
        for d_627_i1_ in range((self).fm2((self).f14, (self.f24)[default__.safeIndex(136, (self.f24).length(0))], (self).f14, (self).f15, globalState), hi3_):
            d_628_v4_: _dafny.Map
            d_628_v4_ = _dafny.Map({(self.f24)[default__.safeIndex(136, (self.f24).length(0))]: (self.f24)[default__.safeIndex(136, (self.f24).length(0))]})
            d_629_v5_: _dafny.MultiSet
            d_629_v5_ = _dafny.MultiSet([(d_628_v4_).set((self).f15, (self.f24)[default__.safeIndex(136, (self.f24).length(0))]), d_628_v4_])
            d_630_v6_: _dafny.MultiSet
            d_630_v6_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cewdik")), (self).f13])
            d_631_v7_: _dafny.Map
            d_631_v7_ = _dafny.Map({-587: (self).f14})
            d_632_v8_: _dafny.Map
            d_632_v8_ = _dafny.Map({(self.f24)[default__.safeIndex(136, (self.f24).length(0))]: d_631_v7_})
            (globalState).f2 = (((self).f12) + (((d_629_v5_)[d_628_v4_] if (d_628_v4_) in (d_629_v5_) else (0) - (((d_630_v6_)[_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_633_i2_ in range(default__.abs(704))])] if (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_634_i2_ in range(default__.abs(704))])) in (d_630_v6_) else (self).f14))))) < (len(d_632_v8_))
            index110_ = default__.safeIndex(711, (d_623_v1_).length(0))
            (d_623_v1_)[index110_] = (self).f11
            index111_ = default__.safeIndex(711, (d_623_v1_).length(0))
            (d_623_v1_)[index111_] = d_627_i1_
            (globalState).f6 = len(((self).f13) + ((self).f13))
            (globalState).f1 = _dafny.Map({(self).f15: (True if (self.f24)[default__.safeIndex(136, (self.f24).length(0))] else ((d_628_v4_)[(self).f15] if ((self).f15) in (d_628_v4_) else (self).f15))})
        if (self).f15:
            d_635_v9_: _dafny.Array
            def lambda25_(d_636_i3_):
                return _dafny.CodePoint('w')

            init12_ = lambda25_
            nw96_ = _dafny.Array(None, 6)
            for i0_12_ in range(nw96_.length(0)):
                nw96_[i0_12_] = init12_(i0_12_)
            d_635_v9_ = nw96_
            d_637_v10_: str
            d_637_v10_ = _dafny.CodePoint('s')
            index112_ = default__.safeIndex(36, (d_635_v9_).length(0))
            (d_635_v9_)[index112_] = d_637_v10_
            index113_ = default__.safeIndex(36, (d_635_v9_).length(0))
            (d_635_v9_)[index113_] = d_637_v10_
            (globalState).f2 = (self.f24)[default__.safeIndex(136, (self.f24).length(0))]
            d_638_v11_: C0
            nw97_ = C0()
            nw97_.ctor__()
            d_638_v11_ = nw97_
            (globalState).f6 = (self).f14
            d_639_v12_: _dafny.Seq
            d_639_v12_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f12])
            (globalState).f6 = (len(d_639_v12_)) + ((self).f11)
        elif True:
            if (self.f24)[default__.safeIndex(136, (self.f24).length(0))]:
                d_640_v13_: _dafny.Array
                def lambda26_(d_641_i4_):
                    return (self.f24)[default__.safeIndex(136, (self.f24).length(0))]

                init13_ = lambda26_
                nw98_ = _dafny.Array(None, 3)
                for i0_13_ in range(nw98_.length(0)):
                    nw98_[i0_13_] = init13_(i0_13_)
                d_640_v13_ = nw98_
                d_642_v14_: _dafny.Map
                d_642_v14_ = _dafny.Map({d_640_v13_: (self).f14})
                d_643_v15_: D5
                d_643_v15_ = D5_DC18(d_642_v14_, (self).f12)
                pat_let_tv7_ = d_642_v14_
                d_644_v16_: _dafny.Array
                nw99_ = _dafny.Array(None, 19)
                nw99_[int(0)] = d_643_v15_
                nw99_[int(1)] = d_643_v15_
                nw99_[int(2)] = d_643_v15_
                nw99_[int(3)] = d_643_v15_
                nw99_[int(4)] = d_643_v15_
                nw99_[int(5)] = d_643_v15_
                nw99_[int(6)] = d_643_v15_
                def iife41_(_pat_let8_0):
                    def iife42_(d_645_dt__update__tmp_h0_):
                        def iife43_(_pat_let9_0):
                            def iife44_(d_646_dt__update_hcf35_h0_):
                                return D5_DC18(d_646_dt__update_hcf35_h0_, (d_645_dt__update__tmp_h0_).cf36)
                            return iife44_(_pat_let9_0)
                        return iife43_(pat_let_tv7_)
                    return iife42_(_pat_let8_0)
                nw99_[int(7)] = iife41_(d_643_v15_)
                nw99_[int(8)] = D5_DC18(d_642_v14_, len(_dafny.Map({(self).f11: (self).f15})))
                nw99_[int(9)] = d_643_v15_
                nw99_[int(10)] = d_643_v15_
                nw99_[int(11)] = d_643_v15_
                nw99_[int(12)] = d_643_v15_
                nw99_[int(13)] = d_643_v15_
                nw99_[int(14)] = d_643_v15_
                nw99_[int(15)] = d_643_v15_
                nw99_[int(16)] = d_643_v15_
                nw99_[int(17)] = d_643_v15_
                nw99_[int(18)] = d_643_v15_
                d_644_v16_ = nw99_
                d_644_v16_ = d_644_v16_
                d_647_v17_: str
                d_647_v17_ = _dafny.CodePoint('w')
                d_647_v17_ = d_647_v17_
                d_648_v18_: _dafny.Map
                d_648_v18_ = _dafny.Map({d_623_v1_: (self).f15})
                d_648_v18_ = (d_648_v18_).set(d_623_v1_, (self.f24)[default__.safeIndex(136, (self.f24).length(0))])
                d_649_v19_: _dafny.Array
                nw100_ = _dafny.Array(_dafny.Set({}), 29)
                d_649_v19_ = nw100_
                index114_ = default__.safeIndex(761, (d_649_v19_).length(0))
                (d_649_v19_)[index114_] = p0
                index115_ = default__.safeIndex(761, (d_649_v19_).length(0))
                (d_649_v19_)[index115_] = ((p0) | (p0)) - (_dafny.Set({True}))
                d_650_v20_: _dafny.MultiSet
                d_650_v20_ = _dafny.MultiSet([(self).f11])
                d_651_v21_: _dafny.Seq
                d_651_v21_ = _dafny.SeqWithoutIsStrInference([((d_650_v20_)[(self).f11] if ((self).f11) in (d_650_v20_) else (self).f14), (self).f12, (self).f14, (self).f14])
                d_652_v22_: _dafny.Map
                d_652_v22_ = _dafny.Map({(self.f24)[default__.safeIndex(136, (self.f24).length(0))]: d_651_v21_})
                d_653_v23_: _dafny.Map
                d_653_v23_ = _dafny.Map({d_640_v13_: ((d_652_v22_)[(self).f15] if ((self).f15) in (d_652_v22_) else d_651_v21_)})
                d_653_v23_ = (d_653_v23_).set(d_640_v13_, (d_651_v21_) + (d_651_v21_))
            elif True:
                arr21_ = self.f24
                index116_ = default__.safeIndex(136, (self.f24).length(0))
                arr21_[index116_] = (self.f24)[default__.safeIndex(136, (self.f24).length(0))]
                d_654_v24_: _dafny.Array
                nw101_ = _dafny.Array(D5.default()(), 11)
                d_654_v24_ = nw101_
                d_655_v25_: _dafny.Array
                nw102_ = _dafny.Array(None, 17)
                nw102_[int(0)] = (self.f24)[default__.safeIndex(136, (self.f24).length(0))]
                nw102_[int(1)] = (self.f24)[default__.safeIndex(136, (self.f24).length(0))]
                nw102_[int(2)] = (self.f24)[default__.safeIndex(136, (self.f24).length(0))]
                nw102_[int(3)] = (self.f24)[default__.safeIndex(136, (self.f24).length(0))]
                nw102_[int(4)] = not((self.f24)[default__.safeIndex(136, (self.f24).length(0))])
                nw102_[int(5)] = True
                nw102_[int(6)] = (self).f15
                nw102_[int(7)] = (self.f24)[default__.safeIndex(136, (self.f24).length(0))]
                nw102_[int(8)] = (self).f15
                nw102_[int(9)] = (self).f15
                nw102_[int(10)] = (self).f15
                nw102_[int(11)] = (self.f24)[default__.safeIndex(136, (self.f24).length(0))]
                nw102_[int(12)] = (self.f24)[default__.safeIndex(136, (self.f24).length(0))]
                nw102_[int(13)] = (self).f15
                nw102_[int(14)] = (self.f24)[default__.safeIndex(136, (self.f24).length(0))]
                nw102_[int(15)] = False
                nw102_[int(16)] = (self).f15
                d_655_v25_ = nw102_
                d_656_v26_: _dafny.Map
                d_656_v26_ = _dafny.Map({d_655_v25_: (self).f11})
                d_657_v27_: D5
                d_657_v27_ = D5_DC18(d_656_v26_, (self).f12)
                index117_ = default__.safeIndex(530, (d_654_v24_).length(0))
                (d_654_v24_)[index117_] = d_657_v27_
                index118_ = default__.safeIndex(530, (d_654_v24_).length(0))
                (d_654_v24_)[index118_] = d_657_v27_
                d_658_v28_: _dafny.Array
                nw103_ = _dafny.Array(_dafny.Map({}), 5)
                d_658_v28_ = nw103_
                d_659_v29_: _dafny.Seq
                d_659_v29_ = _dafny.SeqWithoutIsStrInference([d_658_v28_, d_658_v28_])
                d_660_v30_: _dafny.Map
                d_660_v30_ = _dafny.Map({(self).f13: p0})
                d_661_v31_: _dafny.Map
                d_661_v31_ = _dafny.Map({not(False): len(((d_660_v30_)[(self).f13] if ((self).f13) in (d_660_v30_) else _dafny.Set({(self.f24)[default__.safeIndex(136, (self.f24).length(0))], (self).f15})))})
                d_662_v32_: _dafny.Array
                nw104_ = _dafny.Array(None, 1)
                nw104_[int(0)] = (d_659_v29_)[default__.safeIndex(len(d_661_v31_), len(d_659_v29_))]
                d_662_v32_ = nw104_
                index119_ = default__.safeIndex(461, (d_662_v32_).length(0))
                (d_662_v32_)[index119_] = d_658_v28_
                index120_ = default__.safeIndex(461, (d_662_v32_).length(0))
                (d_662_v32_)[index120_] = d_658_v28_
                d_663_v33_: _dafny.Seq
                d_663_v33_ = _dafny.SeqWithoutIsStrInference([(d_654_v24_)[default__.safeIndex(530, (d_654_v24_).length(0))]])
                d_664_v34_: bool
                d_665_v35_: _dafny.Seq
                out29_: bool
                out30_: _dafny.Seq
                out29_, out30_ = default__.m0((self.f24)[default__.safeIndex(136, (self.f24).length(0))], (self).f15, len((d_663_v33_) + (d_663_v33_)), 53, globalState)
                d_664_v34_ = out29_
                d_665_v35_ = out30_
                (globalState).f6 = (self).f14
            d_666_v36_: str
            d_666_v36_ = _dafny.CodePoint('o')
            d_667_v37_: _dafny.Set
            d_667_v37_ = _dafny.Set({(self).f12})
            d_668_v38_: _dafny.Array
            nw105_ = _dafny.Array(None, 11)
            nw105_[int(0)] = _dafny.CodePoint('p')
            nw105_[int(1)] = d_666_v36_
            nw105_[int(2)] = d_666_v36_
            nw105_[int(3)] = ((self).f13)[default__.safeIndex(len(d_667_v37_), len((self).f13))]
            nw105_[int(4)] = _dafny.CodePoint('v')
            nw105_[int(5)] = d_666_v36_
            nw105_[int(6)] = d_666_v36_
            nw105_[int(7)] = d_666_v36_
            nw105_[int(8)] = d_666_v36_
            nw105_[int(9)] = d_666_v36_
            nw105_[int(10)] = d_666_v36_
            d_668_v38_ = nw105_
            index121_ = default__.safeIndex(631, (d_668_v38_).length(0))
            (d_668_v38_)[index121_] = d_666_v36_
            index122_ = default__.safeIndex(631, (d_668_v38_).length(0))
            (d_668_v38_)[index122_] = d_666_v36_
            d_669_v39_: C0
            nw106_ = C0()
            nw106_.ctor__()
            d_669_v39_ = nw106_
            d_670_v40_: C0
            nw107_ = C0()
            nw107_.ctor__()
            d_670_v40_ = nw107_
            (globalState).f2 = not((self).f15)
        r0 = d_623_v1_
        d_671_v41_: _dafny.Seq
        d_671_v41_ = _dafny.SeqWithoutIsStrInference([p0])
        r1 = ((p0) | (p0)) | ((d_671_v41_)[default__.safeIndex((self).f14, len(d_671_v41_))])
        return r0, r1

    def m14(self, globalState):
        d_672_i0_: int
        d_672_i0_ = 0
        with _dafny.label("4"):
            while (self).f15:
                with _dafny.c_label("4"):
                    if (d_672_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_672_i0_ = (d_672_i0_) + (1)
                    d_673_v0_: D6
                    d_673_v0_ = D6_DC19(_dafny.Map({(self).f15: 418}))
                    d_674_v1_: _dafny.Seq
                    d_674_v1_ = _dafny.SeqWithoutIsStrInference([d_673_v0_, d_673_v0_, d_673_v0_, d_673_v0_])
                    d_675_v2_: _dafny.Set
                    d_675_v2_ = _dafny.Set({(self).f15, False, (self).f15, (self).f15, (self).f15})
                    d_676_v3_: _dafny.Map
                    d_676_v3_ = _dafny.Map({(d_674_v1_).set(default__.safeIndex(len(d_675_v2_), len(d_674_v1_)), d_673_v0_): self.f24})
                    d_677_v4_: _dafny.Map
                    d_677_v4_ = _dafny.Map({d_676_v3_: (self).f14})
                    d_678_v5_: _dafny.Map
                    d_678_v5_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ce")): (self).f15})
                    d_677_v4_ = (d_677_v4_).set(d_676_v3_, (0) - (len(d_678_v5_)))
                    (globalState).f6 = ((self).f14) + ((self).f12)
                    d_679_v6_: _dafny.Set
                    d_679_v6_ = _dafny.Set({(self).f14, (self).f11})
                    d_680_v7_: _dafny.Seq
                    d_680_v7_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                    (globalState).f7 = (((self).f14) in (d_679_v6_) if not(((self).f11) < (222)) else ((self).f11) < (len(d_680_v7_)))
                    d_681_v8_: str
                    d_681_v8_ = _dafny.CodePoint('d')
                    (globalState).f7 = not ((self).f15) or ((d_681_v8_) not in (_dafny.SeqWithoutIsStrInference([d_681_v8_ for d_682_i1_ in range(default__.abs(419))])))
                    pass
            pass
        d_683_v9_: D1
        d_683_v9_ = D1_DC4((self).f15)
        d_684_v10_: _dafny.Seq
        d_684_v10_ = _dafny.SeqWithoutIsStrInference([not((self).f15), not((d_683_v9_).cf5)])
        d_685_v11_: _dafny.Map
        d_685_v11_ = _dafny.Map({(self).f15: (self).f11})
        d_686_v12_: _dafny.MultiSet
        d_686_v12_ = _dafny.MultiSet([(self).f15, not((self).f15), not ((self).f15) or (False), (len(d_684_v10_)) <= (len(d_685_v11_)), (self).f15])
        d_687_v13_: _dafny.Seq
        d_687_v13_ = _dafny.SeqWithoutIsStrInference([d_686_v12_])
        d_686_v12_ = (((d_687_v13_)[default__.safeIndex((self).f14, len(d_687_v13_))]) - (d_686_v12_)) | (d_686_v12_)
        d_688_v14_: C0
        nw108_ = C0()
        nw108_.ctor__()
        d_688_v14_ = nw108_
        d_689_v15_: C0
        nw109_ = C0()
        nw109_.ctor__()
        d_689_v15_ = nw109_
        d_690_v16_: _dafny.Array
        def lambda27_(d_691_i2_):
            return (d_691_i2_) * ((self).f12)

        init14_ = lambda27_
        nw110_ = _dafny.Array(None, 15)
        for i0_14_ in range(nw110_.length(0)):
            nw110_[i0_14_] = init14_(i0_14_)
        d_690_v16_ = nw110_
        index123_ = default__.safeIndex(940, (d_690_v16_).length(0))
        (d_690_v16_)[index123_] = 916
        d_692_v17_: _dafny.Map
        d_692_v17_ = _dafny.Map({(self).f15: (self).f15})
        d_693_v18_: _dafny.Seq
        d_693_v18_ = _dafny.SeqWithoutIsStrInference([((d_689_v15_).fm16(((d_692_v17_)[(self).f15] if ((self).f15) in (d_692_v17_) else (self).f15), (self).f12, (self).f12, (self).f11, globalState) if (self).f15 else (self).f12)])
        index124_ = default__.safeIndex(940, (d_690_v16_).length(0))
        (d_690_v16_)[index124_] = (d_693_v18_)[default__.safeIndex((self).f11, len(d_693_v18_))]
        d_694_v19_: bool
        d_695_v20_: _dafny.Seq
        out31_: bool
        out32_: _dafny.Seq
        out31_, out32_ = (d_689_v15_).m12((self).f13, d_685_v11_, (self).f15, (self).f14, globalState)
        d_694_v19_ = out31_
        d_695_v20_ = out32_


class C2(T1, T0):
    def  __init__(self):
        self._f11: int = int(0)
        self._f12: int = int(0)
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f27: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f27, f13, f11, f12):
        (self).f27 = f27
        (self)._f13 = f13
        (self)._f11 = f11
        (self)._f12 = f12

    def fm1(self, p0, p1, p2, globalState):
        return (not (False) or (True)) or (((self).f12) != ((D11_DC34(True, False, (self).f11, True, True)).cf60))

    def fm40(self, p0, p1, p2, p3, globalState):
        return len((self).f13)

    def fm41(self, globalState):
        return (_dafny.CodePoint('d')) not in ((self).f13)

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        d_696_v0_: _dafny.Map
        d_696_v0_ = _dafny.Map({(self).f12: (self).f13})
        if (len(d_696_v0_)) < ((self).f12):
            (globalState).f2 = False
            (globalState).f6 = p3
            r1 = default__.safeModuloInt((0) - ((0) - (((self).f11) - (464))), self.f27)
            d_697_v1_: _dafny.Set
            d_697_v1_ = _dafny.Set({(0) - (p3)})
            rhs96_ = p2
            rhs97_ = (0) - (p1)
            rhs98_ = (d_697_v1_).ispropersubset((d_697_v1_) - (d_697_v1_))
            lhs74_ = globalState
            lhs75_ = globalState
            lhs76_ = globalState
            lhs74_.f7 = rhs96_
            lhs75_.f6 = rhs97_
            lhs76_.f2 = rhs98_
            d_698_v2_: _dafny.Array
            nw111_ = _dafny.Array(False, 21)
            d_698_v2_ = nw111_
            index125_ = default__.safeIndex(699, (d_698_v2_).length(0))
            (d_698_v2_)[index125_] = (len(_dafny.SeqWithoutIsStrInference([p2]))) <= (p1)
            index126_ = default__.safeIndex(699, (d_698_v2_).length(0))
            (d_698_v2_)[index126_] = (p0) != (p0)
        elif True:
            d_699_v3_: D10
            d_699_v3_ = D10_DC30(p2, p2, (self).f13, p2)
            d_700_v4_: _dafny.Array
            nw112_ = _dafny.Array(_dafny.Map({}), 11)
            d_700_v4_ = nw112_
            d_701_v5_: _dafny.Seq
            d_701_v5_ = _dafny.SeqWithoutIsStrInference([p1])
            d_702_v6_: _dafny.Map
            d_702_v6_ = _dafny.Map({len(d_701_v5_): False})
            index127_ = default__.safeIndex(773, (d_700_v4_).length(0))
            (d_700_v4_)[index127_] = d_702_v6_
            index128_ = default__.safeIndex(773, (d_700_v4_).length(0))
            def iife45_():
                coll25_ = _dafny.Map()
                compr_25_: int
                for compr_25_ in _dafny.IntegerRange(-647, 451):
                    d_703_v7_: int = compr_25_
                    if ((-647) <= (d_703_v7_)) and ((d_703_v7_) < (451)):
                        coll25_[(d_703_v7_) + ((self).f11)] = 157
                return _dafny.Map(coll25_)
            rhs99_ = d_699_v3_
            rhs100_ = p3
            rhs101_ = (default__.fm42(p3, p2, iife45_()
            , globalState)).set(default__.safeModuloInt(p1, p3), p2)
            lhs77_ = d_700_v4_
            lhs78_ = default__.safeIndex(773, (d_700_v4_).length(0))
            d_699_v3_ = rhs99_
            r1 = rhs100_
            lhs77_[lhs78_] = rhs101_
            (globalState).f7 = p2
            d_704_v8_: _dafny.Set
            d_704_v8_ = _dafny.Set({not(p2)})
            d_705_v9_: _dafny.Set
            d_705_v9_ = _dafny.Set({self.f27, (self).f11, p1})
            d_706_v10_: bool
            d_707_v11_: _dafny.Seq
            out33_: bool
            out34_: _dafny.Seq
            out33_, out34_ = default__.m0((D0_DC1((self).fm1(len(d_704_v8_), p2, p2, globalState), (self).f11)).cf1, not(p2), len(d_705_v9_), 747, globalState)
            d_706_v10_ = out33_
            d_707_v11_ = out34_
            d_708_v12_: _dafny.Map
            d_708_v12_ = _dafny.Map({(self).f12: p3})
            d_709_v13_: _dafny.Map
            d_709_v13_ = _dafny.Map({d_708_v12_: d_706_v10_})
            d_709_v13_ = ((d_709_v13_) | (d_709_v13_)) | (d_709_v13_)
            d_710_v14_: _dafny.Seq
            d_710_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvu"))
            rhs102_ = len(((d_710_v14_) + (p0)) + ((p0) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_711_i0_ in range(default__.abs(-176))]))))
            rhs103_ = p0
            lhs79_ = globalState
            lhs79_.f6 = rhs102_
            d_710_v14_ = rhs103_
        d_712_v15_: _dafny.MultiSet
        d_712_v15_ = _dafny.MultiSet([p1])
        d_713_v16_: _dafny.MultiSet
        d_713_v16_ = _dafny.MultiSet([True])
        d_714_v17_: _dafny.Set
        d_714_v17_ = _dafny.Set({p2, p2})
        r1 = ((d_712_v15_)[default__.safeDivisionInt((d_713_v16_).cardinality, 618)] if (default__.safeDivisionInt((d_713_v16_).cardinality, 618)) in (d_712_v15_) else (len(d_714_v17_)) * (len(p0)))
        d_715_v18_: _dafny.Array
        nw113_ = _dafny.Array(False, 2)
        d_715_v18_ = nw113_
        d_716_v19_: _dafny.Set
        d_716_v19_ = _dafny.Set({-875, p3, len(default__.fm43(p2, p2, globalState)), (self).f12})
        d_717_v20_: T2
        nw114_ = C1()
        nw114_.ctor__(d_715_v18_, len(d_716_v19_), 248, -666, p2, (self).f13)
        d_717_v20_ = nw114_
        d_718_v21_: D11
        d_718_v21_ = D11_DC33(d_717_v20_)
        source12_ = d_718_v21_
        if source12_.is_DC34:
            d_719___mcc_h0_ = source12_.cf58
            d_720___mcc_h1_ = source12_.cf59
            d_721___mcc_h2_ = source12_.cf60
            d_722___mcc_h3_ = source12_.cf61
            d_723___mcc_h4_ = source12_.cf62
            d_724_cf62_ = d_723___mcc_h4_
            d_725_cf61_ = d_722___mcc_h3_
            d_726_cf60_ = d_721___mcc_h2_
            d_727_cf59_ = d_720___mcc_h1_
            d_728_cf58_ = d_719___mcc_h0_
            d_729_v22_: _dafny.Map
            d_729_v22_ = _dafny.Map({(self).f11: (d_717_v20_).f12})
            d_730_v23_: _dafny.Array
            nw115_ = _dafny.Array(int(0), 11)
            d_730_v23_ = nw115_
            d_731_v24_: _dafny.MultiSet
            d_731_v24_ = _dafny.MultiSet([d_730_v23_])
            d_732_v25_: _dafny.Map
            d_732_v25_ = _dafny.Map({(d_717_v20_).f15: (d_731_v24_).cardinality})
            d_729_v22_ = (d_729_v22_).set(d_726_cf60_, len(((d_732_v25_).set(p2, (0) - ((d_717_v20_).f14))) | (d_732_v25_)))
            r1 = (self).f11
            d_733_v26_: _dafny.Seq
            d_733_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqqrgfmxc"))])
            d_734_v27_: _dafny.Map
            d_734_v27_ = _dafny.Map({p0: p3})
            d_735_v28_: _dafny.Map
            d_735_v28_ = _dafny.Map({d_729_v22_: (d_717_v20_).f15})
            d_736_v29_: _dafny.Seq
            d_736_v29_ = _dafny.SeqWithoutIsStrInference([-637])
            d_737_v30_: D3
            d_737_v30_ = D3_DC10(False, (d_717_v20_).f11, d_734_v27_, (self).f13, len(d_736_v29_))
            d_738_v31_: _dafny.Set
            d_738_v31_ = _dafny.Set({D3_DC10(d_728_cf58_, 739, d_734_v27_, p0, len(d_735_v28_)), d_737_v30_})
            rhs104_ = default__.safeDivisionInt((0) - ((d_717_v20_).f14), (d_717_v20_).f11)
            rhs105_ = not((d_738_v31_).issubset(d_738_v31_))
            rhs106_ = d_733_v26_
            lhs80_ = globalState
            lhs80_.f6 = rhs104_
            d_727_cf59_ = rhs105_
            d_733_v26_ = rhs106_
            if not(d_728_cf58_):
                d_715_v18_ = d_715_v18_
                index129_ = default__.safeIndex(292, (d_715_v18_).length(0))
                (d_715_v18_)[index129_] = d_727_cf59_
                index130_ = default__.safeIndex(292, (d_715_v18_).length(0))
                (d_715_v18_)[index130_] = d_728_cf58_
                d_713_v16_ = (d_713_v16_).set((d_717_v20_).f15, default__.abs(((d_713_v16_)[d_727_cf59_] if (d_727_cf59_) in (d_713_v16_) else (d_717_v20_).f14)))
                d_736_v29_ = d_736_v29_
                (self).f27 = len((self).f13)
            elif True:
                (globalState).f6 = (d_717_v20_).f12
                d_739_v32_: str
                d_739_v32_ = _dafny.CodePoint('p')
                d_739_v32_ = (d_739_v32_ if d_727_cf59_ else d_739_v32_)
                index131_ = default__.safeIndex(687, (d_715_v18_).length(0))
                (d_715_v18_)[index131_] = d_728_cf58_
                index132_ = default__.safeIndex(687, (d_715_v18_).length(0))
                (d_715_v18_)[index132_] = (default__.fm44((self).f11, not((d_717_v20_).f15), globalState)) == ((p0) + ((d_717_v20_).f13))
                d_740_v33_: _dafny.Seq
                d_740_v33_ = _dafny.SeqWithoutIsStrInference([(d_717_v20_).f15])
                d_741_v34_: _dafny.Array
                nw116_ = _dafny.Array(None, 7)
                nw116_[int(0)] = (d_717_v20_).f15
                nw116_[int(1)] = False
                nw116_[int(2)] = (d_717_v20_).f15
                nw116_[int(3)] = (d_740_v33_)[default__.safeIndex((self).f12, len(d_740_v33_))]
                nw116_[int(4)] = (self).fm1(p1, (d_717_v20_).f15, d_724_cf62_, globalState)
                nw116_[int(5)] = d_728_cf58_
                nw116_[int(6)] = (d_717_v20_).f15
                d_741_v34_ = nw116_
                d_742_v35_: C1
                nw117_ = C1()
                nw117_.ctor__(d_741_v34_, default__.safeModuloInt((d_717_v20_).f12, len(d_714_v17_)), (self).f12, (d_737_v30_).cf20, (d_717_v20_).f15, (d_717_v20_).f13)
                d_742_v35_ = nw117_
                d_743_v36_: D13
                d_743_v36_ = D13_DC37(d_712_v15_)
                d_744_v37_: _dafny.Map
                d_744_v37_ = _dafny.Map({((d_743_v36_).cf67).cardinality: (d_742_v35_).fm1((d_717_v20_).f11, d_728_cf58_, p2, globalState)})
                d_745_v38_: _dafny.Array
                def lambda28_(d_746_p0_):
                    def lambda29_(d_747_i1_):
                        return d_746_p0_

                    return lambda29_

                init15_ = lambda28_(p0)
                nw118_ = _dafny.Array(None, 14)
                for i0_15_ in range(nw118_.length(0)):
                    nw118_[i0_15_] = init15_(i0_15_)
                d_745_v38_ = nw118_
                d_748_v39_: _dafny.Map
                d_748_v39_ = _dafny.Map({len(d_744_v37_): d_745_v38_})
                d_748_v39_ = (d_748_v39_).set((d_717_v20_).f14, d_745_v38_)
        elif True:
            d_749___mcc_h5_ = source12_.cf57
            d_750_cf57_ = d_749___mcc_h5_
            r0 = p2
            d_751_v40_: _dafny.Map
            d_751_v40_ = _dafny.Map({(len(default__.fm44(p3, (d_717_v20_).f15, globalState))) - ((d_717_v20_).f14): (d_716_v19_ if (d_717_v20_).f15 else d_716_v19_)})
            d_752_v41_: str
            d_752_v41_ = _dafny.CodePoint('e')
            d_753_v43_: _dafny.Seq
            def iife46_():
                coll26_ = _dafny.Map()
                compr_26_: int
                for compr_26_ in _dafny.IntegerRange(716, -651):
                    d_754_v42_: int = compr_26_
                    if ((716) <= (d_754_v42_)) and ((d_754_v42_) < (-651)):
                        coll26_[(d_754_v42_) * ((d_750_cf57_).f11)] = _dafny.SeqWithoutIsStrInference([True])
                return _dafny.Map(coll26_)
            d_753_v43_ = _dafny.SeqWithoutIsStrInference([(d_751_v40_) | (d_751_v40_), (_dafny.Map({(self).fm40(67, (d_717_v20_).f15, d_752_v41_, p2, globalState): d_716_v19_})).set((d_717_v20_).f14, d_716_v19_), (default__.fm45(len(iife46_()
            ), (d_717_v20_).f15, p0, globalState)).set(len(d_714_v17_), d_716_v19_)])
            d_755_v44_: _dafny.Map
            d_755_v44_ = _dafny.Map({True: (d_717_v20_).f15})
            d_756_v45_: _dafny.Seq
            d_756_v45_ = _dafny.SeqWithoutIsStrInference([d_755_v44_, _dafny.Map({False: (d_717_v20_).f15})])
            d_751_v40_ = (d_753_v43_)[default__.safeIndex((222) - (len((d_756_v45_)[default__.safeIndex(414, len(d_756_v45_))])), len(d_753_v43_))]
            rhs107_ = (0) - (default__.safeDivisionInt((d_717_v20_).f14, p3))
            rhs108_ = ((0) - ((self).f11)) - ((self).f12)
            rhs109_ = (d_750_cf57_).f14
            lhs81_ = self
            r2 = rhs107_
            lhs81_.f27 = rhs108_
            r1 = rhs109_
            d_757_v46_: _dafny.Array
            nw119_ = _dafny.Array(int(0), 29)
            d_757_v46_ = nw119_
            d_758_v47_: _dafny.Map
            d_758_v47_ = _dafny.Map({(d_717_v20_).f15: d_757_v46_})
            d_758_v47_ = (d_758_v47_).set(not (False) or (False), d_757_v46_)
        hi4_ = (d_717_v20_).f14
        for d_759_i2_ in range((d_717_v20_).f14, hi4_):
            d_760_v48_: _dafny.Array
            nw120_ = _dafny.Array(int(0), 24)
            d_760_v48_ = nw120_
            d_760_v48_ = d_760_v48_
            d_761_v49_: _dafny.Seq
            d_761_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykyoyrt"))
            d_762_v50_: _dafny.MultiSet
            d_762_v50_ = _dafny.MultiSet([_dafny.CodePoint('b')])
            d_763_v51_: str
            d_763_v51_ = _dafny.CodePoint('i')
            index133_ = default__.safeIndex(510, (d_760_v48_).length(0))
            (d_760_v48_)[index133_] = ((d_762_v50_)[d_763_v51_] if (d_763_v51_) in (d_762_v50_) else 96)
            d_764_v52_: _dafny.Seq
            d_764_v52_ = _dafny.SeqWithoutIsStrInference([(d_717_v20_).f15])
            index134_ = default__.safeIndex(510, (d_760_v48_).length(0))
            rhs110_ = p0
            rhs111_ = len(((d_764_v52_) + (d_764_v52_)) + (d_764_v52_))
            rhs112_ = default__.safeDivisionInt((d_717_v20_).f14, (d_717_v20_).f11)
            rhs113_ = (0) - (len((d_717_v20_).f13))
            lhs82_ = d_760_v48_
            lhs83_ = default__.safeIndex(510, (d_760_v48_).length(0))
            lhs84_ = globalState
            lhs85_ = self
            d_761_v49_ = rhs110_
            lhs82_[lhs83_] = rhs111_
            lhs84_.f6 = rhs112_
            lhs85_.f27 = rhs113_
            d_765_v53_: C1
            nw121_ = C1()
            nw121_.ctor__(d_715_v18_, (d_717_v20_).f11, p1, (d_717_v20_).f11, p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjeqiw")))
            d_765_v53_ = nw121_
            d_766_v54_: _dafny.Map
            d_766_v54_ = _dafny.Map({(d_717_v20_).fm1((0) - ((d_717_v20_).f14), (d_717_v20_).f15, (d_717_v20_).f15, globalState): (d_717_v20_).f15})
            d_767_v55_: _dafny.Seq
            d_767_v55_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(True), p2])])
            d_768_v56_: D14
            d_768_v56_ = D14_DC40(d_764_v52_)
            d_769_v57_: _dafny.Seq
            d_769_v57_ = _dafny.SeqWithoutIsStrInference([(d_767_v55_)[default__.safeIndex(len((d_768_v56_).cf72), len(d_767_v55_))], _dafny.SeqWithoutIsStrInference([(d_765_v53_).fm1((default__.fm46(p2, (d_712_v15_).cardinality, d_759_i2_, 523, globalState)).cf24, (d_717_v20_).f15, p2, globalState), (d_717_v20_).f15])])
            d_770_v58_: _dafny.Map
            d_770_v58_ = _dafny.Map({d_763_v51_: len(d_764_v52_)})
            d_766_v54_ = (d_766_v54_).set((d_717_v20_).f15, ((d_717_v20_).f15) in ((d_767_v55_)[default__.safeIndex(((d_770_v58_)[d_763_v51_] if (d_763_v51_) in (d_770_v58_) else (d_717_v20_).f14), len(d_767_v55_))]))
        d_771_v59_: _dafny.Seq
        d_771_v59_ = _dafny.SeqWithoutIsStrInference([(self).f12, (self).f11, (self).f12])
        d_772_v60_: _dafny.Seq
        d_772_v60_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(d_717_v20_).f15, p2, (d_717_v20_).f15, False])).cardinality])): p1})) for d_773_i4_ in range(default__.abs(807))]), d_771_v59_])
        hi5_ = len(d_772_v60_)
        for d_774_i3_ in range(954, hi5_):
            index135_ = default__.safeIndex(674, (d_715_v18_).length(0))
            (d_715_v18_)[index135_] = (d_717_v20_).f15
            index136_ = default__.safeIndex(674, (d_715_v18_).length(0))
            (d_715_v18_)[index136_] = (d_717_v20_).f15
            d_775_v61_: C0
            nw122_ = C0()
            nw122_.ctor__()
            d_775_v61_ = nw122_
            d_776_v62_: str
            d_776_v62_ = _dafny.CodePoint('c')
            d_776_v62_ = (p0)[default__.safeIndex(872, len(p0))]
            d_777_v63_: C0
            nw123_ = C0()
            nw123_.ctor__()
            d_777_v63_ = nw123_
        if (d_717_v20_).f15:
            (globalState).f2 = p2
            (globalState).f2 = (d_717_v20_).f15
            d_778_v64_: _dafny.Seq
            d_778_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_779_v65_: str
            d_779_v65_ = _dafny.CodePoint('s')
            d_778_v64_ = (((self).f13) + ((self).f13)).set(default__.safeIndex(p3, len(((self).f13) + ((self).f13))), d_779_v65_)
            d_780_v66_: _dafny.Array
            nw124_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_780_v66_ = nw124_
            d_781_v67_: _dafny.Array
            nw125_ = _dafny.Array(_dafny.Seq({}), 2)
            d_781_v67_ = nw125_
            index137_ = default__.safeIndex(77, (d_780_v66_).length(0))
            (d_780_v66_)[index137_] = d_781_v67_
            index138_ = default__.safeIndex(77, (d_780_v66_).length(0))
            (d_780_v66_)[index138_] = d_781_v67_
            index139_ = default__.safeIndex(506, (d_715_v18_).length(0))
            (d_715_v18_)[index139_] = (d_717_v20_).f15
            index140_ = default__.safeIndex(506, (d_715_v18_).length(0))
            (d_715_v18_)[index140_] = ((self.f27) in (d_771_v59_)) == ((d_717_v20_).f15)
        elif True:
            (self).f27 = p3
            d_782_v68_: str
            d_782_v68_ = _dafny.CodePoint('a')
            d_783_v69_: _dafny.Array
            nw126_ = _dafny.Array(None, 20)
            nw126_[int(0)] = d_782_v68_
            nw126_[int(1)] = d_782_v68_
            nw126_[int(2)] = d_782_v68_
            nw126_[int(3)] = d_782_v68_
            nw126_[int(4)] = d_782_v68_
            nw126_[int(5)] = _dafny.CodePoint('d')
            nw126_[int(6)] = d_782_v68_
            nw126_[int(7)] = default__.fm47((0) - (-901), globalState)
            nw126_[int(8)] = d_782_v68_
            nw126_[int(9)] = d_782_v68_
            nw126_[int(10)] = (d_782_v68_ if (d_717_v20_).f15 else d_782_v68_)
            nw126_[int(11)] = d_782_v68_
            nw126_[int(12)] = d_782_v68_
            nw126_[int(13)] = d_782_v68_
            nw126_[int(14)] = d_782_v68_
            nw126_[int(15)] = d_782_v68_
            nw126_[int(16)] = d_782_v68_
            nw126_[int(17)] = d_782_v68_
            nw126_[int(18)] = d_782_v68_
            nw126_[int(19)] = d_782_v68_
            d_783_v69_ = nw126_
            index141_ = default__.safeIndex(592, (d_783_v69_).length(0))
            (d_783_v69_)[index141_] = d_782_v68_
            d_784_v70_: _dafny.Seq
            d_784_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
            index142_ = default__.safeIndex(592, (d_783_v69_).length(0))
            rhs114_ = ((d_714_v17_) | (d_714_v17_)).ispropersubset(d_714_v17_)
            rhs115_ = (((self).f11 if not(p2) else ((d_713_v16_)[(d_717_v20_).f15] if ((d_717_v20_).f15) in (d_713_v16_) else (d_717_v20_).f14))) <= ((d_717_v20_).fm2(self.f27, p2, (d_717_v20_).f12, (d_717_v20_).f15, globalState))
            rhs116_ = d_782_v68_
            rhs117_ = default__.fm44((0) - (((0) - ((0) - (-104))) + ((d_713_v16_).cardinality)), p2, globalState)
            rhs118_ = p1
            lhs86_ = globalState
            lhs87_ = globalState
            lhs88_ = d_783_v69_
            lhs89_ = default__.safeIndex(592, (d_783_v69_).length(0))
            lhs90_ = self
            lhs86_.f7 = rhs114_
            lhs87_.f7 = rhs115_
            lhs88_[lhs89_] = rhs116_
            d_784_v70_ = rhs117_
            lhs90_.f27 = rhs118_
            d_785_v71_: _dafny.Array
            nw127_ = _dafny.Array(int(0), 29)
            d_785_v71_ = nw127_
            index143_ = default__.safeIndex(528, (d_785_v71_).length(0))
            (d_785_v71_)[index143_] = (d_717_v20_).f12
            index144_ = default__.safeIndex(528, (d_785_v71_).length(0))
            (d_785_v71_)[index144_] = ((self).f12) * ((d_717_v20_).f11)
            index145_ = default__.safeIndex(958, (d_715_v18_).length(0))
            (d_715_v18_)[index145_] = False
            index146_ = default__.safeIndex(958, (d_715_v18_).length(0))
            (d_715_v18_)[index146_] = not (p2) or ((d_717_v20_).f15)
            d_786_v72_: _dafny.Array
            nw128_ = _dafny.Array(_dafny.Map({}), 9)
            d_786_v72_ = nw128_
            d_787_v73_: _dafny.Map
            d_787_v73_ = _dafny.Map({True: (d_717_v20_).f11})
            index147_ = default__.safeIndex(887, (d_786_v72_).length(0))
            (d_786_v72_)[index147_] = (d_787_v73_).set(True, (d_717_v20_).f11)
            index148_ = default__.safeIndex(958, (d_715_v18_).length(0))
            index149_ = default__.safeIndex(887, (d_786_v72_).length(0))
            rhs119_ = (d_714_v17_) == (d_714_v17_)
            rhs120_ = default__.safeModuloInt((d_717_v20_).f12, len(d_771_v59_))
            rhs121_ = d_787_v73_
            lhs91_ = d_715_v18_
            lhs92_ = default__.safeIndex(958, (d_715_v18_).length(0))
            lhs93_ = globalState
            lhs94_ = d_786_v72_
            lhs95_ = default__.safeIndex(887, (d_786_v72_).length(0))
            lhs91_[lhs92_] = rhs119_
            lhs93_.f6 = rhs120_
            lhs94_[lhs95_] = rhs121_
        r0 = p2
        r1 = ((d_717_v20_).f14 if (831) <= (p1) else 411)
        r2 = (d_713_v16_).cardinality
        return r0, r1, r2

    def m19(self, p0, p1, p2, p3, globalState):
        d_788_v0_: bool
        d_788_v0_ = False
        (globalState).f7 = not(d_788_v0_)
        d_789_v1_: str
        d_789_v1_ = _dafny.CodePoint('w')
        d_789_v1_ = d_789_v1_
        d_790_v2_: _dafny.Map
        d_790_v2_ = _dafny.Map({d_788_v0_: d_788_v0_})
        d_791_v3_: _dafny.Map
        d_791_v3_ = _dafny.Map({not(d_788_v0_): (self).f12})
        if ((d_790_v2_)[(True) not in (d_791_v3_)] if ((True) not in (d_791_v3_)) in (d_790_v2_) else not((self).fm41(globalState))):
            (globalState).f6 = (self).f12
            d_792_v4_: _dafny.Seq
            d_792_v4_ = _dafny.SeqWithoutIsStrInference([d_788_v0_, False])
            d_793_v5_: _dafny.Seq
            d_793_v5_ = _dafny.SeqWithoutIsStrInference([len(d_792_v4_)])
            d_794_v6_: _dafny.Seq
            d_794_v6_ = _dafny.SeqWithoutIsStrInference([default__.fm48(globalState)])
            d_795_v7_: _dafny.Set
            d_795_v7_ = _dafny.Set({(self).f11})
            d_796_v8_: _dafny.Map
            d_796_v8_ = _dafny.Map({len(d_795_v7_): d_789_v1_})
            d_797_v9_: _dafny.Array
            nw129_ = _dafny.Array(None, 16)
            nw129_[int(0)] = d_793_v5_
            nw129_[int(1)] = d_793_v5_
            nw129_[int(2)] = d_793_v5_
            nw129_[int(3)] = _dafny.SeqWithoutIsStrInference([self.f27 for d_798_i0_ in range(default__.abs(147))])
            nw129_[int(4)] = _dafny.SeqWithoutIsStrInference([self.f27])
            nw129_[int(5)] = d_793_v5_
            nw129_[int(6)] = (d_793_v5_).set(default__.safeIndex((self).f12, len(d_793_v5_)), 239)
            nw129_[int(7)] = _dafny.SeqWithoutIsStrInference([p0, (self).f12, 685, p0])
            nw129_[int(8)] = (d_794_v6_)[default__.safeIndex(len(d_796_v8_), len(d_794_v6_))]
            nw129_[int(9)] = d_793_v5_
            nw129_[int(10)] = d_793_v5_
            nw129_[int(11)] = d_793_v5_
            nw129_[int(12)] = _dafny.SeqWithoutIsStrInference([(self).f12 for d_799_i1_ in range(default__.abs(23))])
            nw129_[int(13)] = d_793_v5_
            nw129_[int(14)] = d_793_v5_
            nw129_[int(15)] = d_793_v5_
            d_797_v9_ = nw129_
            d_800_v10_: int
            out35_: int
            out35_ = (self).m20((self).fm41(globalState), _dafny.Set({d_788_v0_}), d_797_v9_, d_788_v0_, globalState)
            d_800_v10_ = out35_
            d_801_v11_: C0
            nw130_ = C0()
            nw130_.ctor__()
            d_801_v11_ = nw130_
            d_802_v12_: _dafny.Array
            nw131_ = _dafny.Array(int(0), 14)
            d_802_v12_ = nw131_
            index150_ = default__.safeIndex(42, (d_802_v12_).length(0))
            (d_802_v12_)[index150_] = (self).f11
            index151_ = default__.safeIndex(42, (d_802_v12_).length(0))
            (d_802_v12_)[index151_] = (self).f12
            (globalState).f7 = True
        elif True:
            (globalState).f7 = ((self).fm41(globalState)) == (d_788_v0_)
            d_803_v13_: _dafny.Seq
            d_803_v13_ = _dafny.SeqWithoutIsStrInference([p1])
            d_804_v14_: _dafny.Seq
            d_804_v14_ = _dafny.SeqWithoutIsStrInference([d_788_v0_])
            d_805_v15_: _dafny.Map
            d_805_v15_ = _dafny.Map({(d_803_v13_) + (d_803_v13_): default__.fm44(len(d_804_v14_), d_788_v0_, globalState)})
            rhs122_ = (len(d_803_v13_)) * (len(_dafny.Map({d_788_v0_: (self).fm41(globalState)})))
            rhs123_ = d_788_v0_
            rhs124_ = True
            rhs125_ = ((d_805_v15_ if d_788_v0_ else d_805_v15_)) | ((d_805_v15_) | (d_805_v15_))
            rhs126_ = not(not(d_788_v0_))
            lhs96_ = globalState
            lhs97_ = globalState
            lhs96_.f6 = rhs122_
            d_788_v0_ = rhs123_
            lhs97_.f2 = rhs124_
            d_805_v15_ = rhs125_
            d_788_v0_ = rhs126_
            d_806_v16_: _dafny.Array
            nw132_ = _dafny.Array(None, 9)
            nw132_[int(0)] = (self).f12
            nw132_[int(1)] = self.f27
            nw132_[int(2)] = self.f27
            nw132_[int(3)] = (p1 if d_788_v0_ else len((self).f13))
            nw132_[int(4)] = (self).f12
            nw132_[int(5)] = (self).f11
            nw132_[int(6)] = len((d_803_v13_) + (d_803_v13_))
            nw132_[int(7)] = (self).fm40((d_803_v13_)[default__.safeIndex(len(d_804_v14_), len(d_803_v13_))], d_788_v0_, d_789_v1_, d_788_v0_, globalState)
            nw132_[int(8)] = (len((self).f13)) * (p0)
            d_806_v16_ = nw132_
            d_807_v17_: _dafny.Map
            d_807_v17_ = _dafny.Map({(self).f13: D11_DC34(d_788_v0_, d_788_v0_, -590, d_788_v0_, d_788_v0_)})
            d_808_v18_: _dafny.Seq
            d_808_v18_ = _dafny.SeqWithoutIsStrInference([(self).f13])
            d_809_v19_: D11
            d_809_v19_ = D11_DC34(d_788_v0_, d_788_v0_, (self).f11, d_788_v0_, d_788_v0_)
            index152_ = default__.safeIndex(95, (d_806_v16_).length(0))
            (d_806_v16_)[index152_] = len((d_807_v17_ if False else _dafny.Map({(d_808_v18_)[default__.safeIndex(p0, len(d_808_v18_))]: d_809_v19_})))
            d_810_v20_: _dafny.Seq
            d_810_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kddujed"))
            index153_ = default__.safeIndex(95, (d_806_v16_).length(0))
            rhs127_ = p1
            rhs128_ = d_788_v0_
            rhs129_ = (d_804_v14_) + ((d_804_v14_) + (_dafny.SeqWithoutIsStrInference([d_788_v0_, not(False)])))
            rhs130_ = (((self).f13 if d_788_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hke")))) + (d_810_v20_)
            rhs131_ = ((self).f12) == ((self).fm40(864, d_788_v0_, d_789_v1_, d_788_v0_, globalState))
            lhs98_ = d_806_v16_
            lhs99_ = default__.safeIndex(95, (d_806_v16_).length(0))
            lhs100_ = globalState
            lhs98_[lhs99_] = rhs127_
            lhs100_.f2 = rhs128_
            d_804_v14_ = rhs129_
            d_810_v20_ = rhs130_
            d_788_v0_ = rhs131_
            d_811_v21_: D14
            d_811_v21_ = D14_DC41((self).fm1(p1, d_788_v0_, d_788_v0_, globalState), d_788_v0_, _dafny.SeqWithoutIsStrInference([d_789_v1_ for d_812_i2_ in range(default__.abs(309))]))
            d_813_v22_: D10
            d_813_v22_ = D10_DC30((True) or (d_788_v0_), (d_804_v14_) == (d_804_v14_), (d_811_v21_).cf75, not (d_788_v0_) or (d_788_v0_))
            source13_ = d_813_v22_
            if source13_.is_DC29:
                d_814_v23_: _dafny.MultiSet
                d_814_v23_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_789_v1_]), _dafny.SeqWithoutIsStrInference([d_789_v1_]), (self).f13, (self).f13, (self).f13])
                d_815_v24_: _dafny.Map
                d_815_v24_ = _dafny.Map({(self).f13: (d_806_v16_)[default__.safeIndex(95, (d_806_v16_).length(0))]})
                d_816_v25_: D10
                d_816_v25_ = D10_DC31((d_814_v23_).cardinality, d_789_v1_, (p3) + (p0), len(d_815_v24_))
                d_816_v25_ = (d_816_v25_ if (self).fm1(p0, d_788_v0_, True, globalState) else d_816_v25_)
                d_788_v0_ = d_788_v0_
                (self).f27 = 910
                d_817_v26_: C0
                nw133_ = C0()
                nw133_.ctor__()
                d_817_v26_ = nw133_
            elif source13_.is_DC30:
                d_818___mcc_h0_ = source13_.cf48
                d_819___mcc_h1_ = source13_.cf49
                d_820___mcc_h2_ = source13_.cf50
                d_821___mcc_h3_ = source13_.cf51
                d_822_cf51_ = d_821___mcc_h3_
                d_823_cf50_ = d_820___mcc_h2_
                d_824_cf49_ = d_819___mcc_h1_
                d_825_cf48_ = d_818___mcc_h0_
                (globalState).f7 = d_825_cf48_
                (globalState).f2 = d_825_cf48_
                d_826_v27_: _dafny.Array
                def lambda30_(d_827_cf48_):
                    def lambda31_(d_828_i3_):
                        return d_827_cf48_

                    return lambda31_

                init16_ = lambda30_(d_825_cf48_)
                nw134_ = _dafny.Array(None, 29)
                for i0_16_ in range(nw134_.length(0)):
                    nw134_[i0_16_] = init16_(i0_16_)
                d_826_v27_ = nw134_
                d_829_v28_: _dafny.Set
                d_829_v28_ = _dafny.Set({d_826_v27_})
                d_830_v29_: _dafny.Seq
                d_830_v29_ = _dafny.SeqWithoutIsStrInference([d_829_v28_, d_829_v28_, d_829_v28_, d_829_v28_])
                d_829_v28_ = ((_dafny.Set({d_826_v27_, d_826_v27_}) if d_822_cf51_ else d_829_v28_) if d_822_cf51_ else (d_830_v29_)[default__.safeIndex(p3, len(d_830_v29_))])
                d_831_v30_: _dafny.MultiSet
                d_831_v30_ = _dafny.MultiSet([(d_806_v16_)[default__.safeIndex(95, (d_806_v16_).length(0))], (_dafny.MultiSet([d_806_v16_])).cardinality])
                d_832_v31_: _dafny.Map
                d_832_v31_ = _dafny.Map({len(d_803_v13_): (self).f11})
                d_833_v32_: _dafny.Map
                d_833_v32_ = _dafny.Map({d_804_v14_: d_832_v31_})
                d_834_v33_: _dafny.Set
                d_834_v33_ = _dafny.Set({d_831_v30_, d_831_v30_, (d_831_v30_).set(len(d_833_v32_), default__.abs(len(d_803_v13_))), (d_831_v30_).set(len(d_804_v14_), default__.abs((d_806_v16_)[default__.safeIndex(95, (d_806_v16_).length(0))])), d_831_v30_})
                d_835_v34_: _dafny.Map
                d_835_v34_ = _dafny.Map({d_831_v30_: False})
                def iife47_():
                    coll27_ = _dafny.Set()
                    compr_27_: _dafny.MultiSet
                    for compr_27_ in (d_835_v34_).keys.Elements:
                        d_836_v35_: _dafny.MultiSet = compr_27_
                        if (d_836_v35_) in (d_835_v34_):
                            coll27_ = coll27_.union(_dafny.Set([d_836_v35_]))
                    return _dafny.Set(coll27_)
                rhs132_ = ((d_834_v33_ if d_788_v0_ else iife47_()
                )) - (d_834_v33_)
                rhs133_ = (self).fm1(-580, d_788_v0_, (238) < ((self).f11), globalState)
                d_834_v33_ = rhs132_
                d_822_cf51_ = rhs133_
            elif source13_.is_DC31:
                d_837___mcc_h4_ = source13_.cf52
                d_838___mcc_h5_ = source13_.cf53
                d_839___mcc_h6_ = source13_.cf54
                d_840___mcc_h7_ = source13_.cf55
                d_841_cf55_ = d_840___mcc_h7_
                d_842_cf54_ = d_839___mcc_h6_
                d_843_cf53_ = d_838___mcc_h5_
                d_844_cf52_ = d_837___mcc_h4_
                (globalState).f6 = (0) - (default__.safeDivisionInt(((d_806_v16_)[default__.safeIndex(95, (d_806_v16_).length(0))]) - (d_844_cf52_), d_844_cf52_))
                d_845_v36_: _dafny.MultiSet
                d_845_v36_ = _dafny.MultiSet([True])
                (self).f27 = default__.safeDivisionInt(p1, ((d_845_v36_)[d_788_v0_] if (d_788_v0_) in (d_845_v36_) else (0) - ((0) - (d_844_cf52_))))
                d_846_v37_: D9
                d_846_v37_ = D9_DC27(d_806_v16_)
                d_806_v16_ = (d_846_v37_).cf46
                index154_ = default__.safeIndex(95, (d_806_v16_).length(0))
                rhs134_ = d_806_v16_
                rhs135_ = (d_806_v16_)[default__.safeIndex(95, (d_806_v16_).length(0))]
                lhs101_ = d_806_v16_
                lhs102_ = default__.safeIndex(95, (d_806_v16_).length(0))
                d_806_v16_ = rhs134_
                lhs101_[lhs102_] = rhs135_
            elif source13_.is_DC28:
                d_847___mcc_h8_ = source13_.cf47
                d_848_cf47_ = d_847___mcc_h8_
                d_849_v38_: _dafny.Set
                d_849_v38_ = _dafny.Set({(self).f13, (self).f13, (self).f13})
                d_850_v39_: D5
                d_850_v39_ = D5_DC17(d_849_v38_)
                d_851_v40_: D1
                d_851_v40_ = D1_DC4(d_788_v0_)
                d_852_v41_: _dafny.Map
                d_852_v41_ = _dafny.Map({d_850_v39_: d_851_v40_})
                d_852_v41_ = (d_852_v41_) | ((d_852_v41_) | (_dafny.Map({d_850_v39_: d_851_v40_})))
                d_853_v42_: _dafny.Set
                d_853_v42_ = _dafny.Set({p0, (self).f11})
                d_854_v43_: _dafny.Array
                nw135_ = _dafny.Array(None, 9)
                nw135_[int(0)] = d_788_v0_
                nw135_[int(1)] = d_788_v0_
                nw135_[int(2)] = (self.f27) != ((self).fm40(p1, d_788_v0_, d_789_v1_, d_788_v0_, globalState))
                nw135_[int(3)] = d_788_v0_
                nw135_[int(4)] = d_788_v0_
                nw135_[int(5)] = d_788_v0_
                nw135_[int(6)] = False
                nw135_[int(7)] = d_788_v0_
                nw135_[int(8)] = (d_853_v42_).isdisjoint(d_853_v42_)
                d_854_v43_ = nw135_
                index155_ = default__.safeIndex(883, (d_854_v43_).length(0))
                (d_854_v43_)[index155_] = d_788_v0_
                index156_ = default__.safeIndex(883, (d_854_v43_).length(0))
                (d_854_v43_)[index156_] = (self).fm41(globalState)
                (globalState).f6 = (d_806_v16_)[default__.safeIndex(95, (d_806_v16_).length(0))]
                (globalState).f6 = ((self).f11) + ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxxopb"))) if d_788_v0_ else p3))
            elif True:
                d_855___mcc_h9_ = source13_.cf56
                d_856_cf56_ = d_855___mcc_h9_
                index157_ = default__.safeIndex(95, (d_806_v16_).length(0))
                (d_806_v16_)[index157_] = ((self).f11) - (p0)
                d_857_v44_: D6
                d_857_v44_ = D6_DC19(d_791_v3_)
                d_858_v45_: _dafny.Map
                d_858_v45_ = _dafny.Map({_dafny.MultiSet([d_857_v44_]): default__.fm43(d_788_v0_, d_788_v0_, globalState)})
                d_859_v46_: _dafny.Map
                d_859_v46_ = _dafny.Map({p1: (d_804_v14_).set(default__.safeIndex(p1, len(d_804_v14_)), d_788_v0_)})
                d_860_v47_: _dafny.MultiSet
                d_860_v47_ = _dafny.MultiSet([d_857_v44_])
                d_861_v48_: _dafny.Map
                d_861_v48_ = _dafny.Map({len(((d_859_v46_)[p0] if (p0) in (d_859_v46_) else d_804_v14_)): d_860_v47_})
                d_858_v45_ = ((d_858_v45_).set(_dafny.MultiSet([d_857_v44_]), d_804_v14_)).set((((d_861_v48_)[(d_806_v16_)[default__.safeIndex(95, (d_806_v16_).length(0))]] if ((d_806_v16_)[default__.safeIndex(95, (d_806_v16_).length(0))]) in (d_861_v48_) else d_860_v47_)).set(d_857_v44_, default__.abs(p1)), _dafny.SeqWithoutIsStrInference([d_788_v0_, d_788_v0_]))
                d_862_v49_: _dafny.Map
                d_862_v49_ = _dafny.Map({((self).f12) + ((self).f11): default__.safeModuloInt(len((self).f13), (self).f11)})
                d_862_v49_ = (d_862_v49_).set((len(_dafny.Map({d_788_v0_: d_788_v0_}))) * (self.f27), (d_806_v16_)[default__.safeIndex(95, (d_806_v16_).length(0))])
                d_863_v50_: D1
                d_863_v50_ = D1_DC2(d_789_v1_)
                pat_let_tv8_ = d_789_v1_
                pat_let_tv9_ = d_789_v1_
                d_864_v51_: _dafny.Set
                def iife48_(_pat_let10_0):
                    def iife49_(d_865_dt__update__tmp_h0_):
                        def iife50_(_pat_let11_0):
                            def iife51_(d_866_dt__update_hcf3_h0_):
                                return D1_DC2(d_866_dt__update_hcf3_h0_)
                            return iife51_(_pat_let11_0)
                        return iife50_(pat_let_tv8_)
                    return iife49_(_pat_let10_0)
                def iife52_(_pat_let12_0):
                    def iife53_(d_867_dt__update__tmp_h1_):
                        def iife54_(_pat_let13_0):
                            def iife55_(d_868_dt__update_hcf3_h1_):
                                return D1_DC2(d_868_dt__update_hcf3_h1_)
                            return iife55_(_pat_let13_0)
                        return iife54_(pat_let_tv9_)
                    return iife53_(_pat_let12_0)
                d_864_v51_ = _dafny.Set({D1_DC2(d_789_v1_), D1_DC2(d_789_v1_), iife48_(d_863_v50_), iife52_(d_863_v50_), d_863_v50_})
                d_869_v52_: _dafny.Seq
                d_869_v52_ = _dafny.SeqWithoutIsStrInference([default__.fm49(d_788_v0_, d_788_v0_, (self).f12, globalState), d_863_v50_])
                d_870_v53_: _dafny.Seq
                d_870_v53_ = _dafny.SeqWithoutIsStrInference([d_869_v52_])
                def iife56_():
                    coll28_ = _dafny.Set()
                    compr_28_: D1
                    for compr_28_ in ((d_870_v53_)[default__.safeIndex(p1, len(d_870_v53_))]).Elements:
                        d_871_v54_: D1 = compr_28_
                        if (d_871_v54_) in ((d_870_v53_)[default__.safeIndex(p1, len(d_870_v53_))]):
                            coll28_ = coll28_.union(_dafny.Set([d_871_v54_]))
                    return _dafny.Set(coll28_)
                (globalState).f2 = ((iife56_()
                 if d_788_v0_ else d_864_v51_)).issubset(d_864_v51_)
            d_872_v55_: C0
            nw136_ = C0()
            nw136_.ctor__()
            d_872_v55_ = nw136_
        d_873_v56_: _dafny.Array
        def lambda32_(d_874_v0_, d_875_v1_):
            def lambda33_(d_876_i4_):
                return _dafny.Map({d_874_v0_: d_875_v1_})

            return lambda33_

        init17_ = lambda32_(d_788_v0_, d_789_v1_)
        nw137_ = _dafny.Array(None, 11)
        for i0_17_ in range(nw137_.length(0)):
            nw137_[i0_17_] = init17_(i0_17_)
        d_873_v56_ = nw137_
        d_873_v56_ = d_873_v56_
        d_877_v57_: _dafny.Array
        nw138_ = _dafny.Array(None, 13)
        nw138_[int(0)] = d_788_v0_
        nw138_[int(1)] = d_788_v0_
        nw138_[int(2)] = d_788_v0_
        nw138_[int(3)] = d_788_v0_
        nw138_[int(4)] = not(d_788_v0_)
        nw138_[int(5)] = d_788_v0_
        nw138_[int(6)] = d_788_v0_
        nw138_[int(7)] = d_788_v0_
        nw138_[int(8)] = d_788_v0_
        nw138_[int(9)] = not(d_788_v0_)
        nw138_[int(10)] = d_788_v0_
        nw138_[int(11)] = d_788_v0_
        nw138_[int(12)] = d_788_v0_
        d_877_v57_ = nw138_
        d_878_v58_: C1
        nw139_ = C1()
        nw139_.ctor__(d_877_v57_, default__.fm0(d_788_v0_, globalState), p3, p3, not(d_788_v0_), (self).f13)
        d_878_v58_ = nw139_
        d_878_v58_ = d_878_v58_
        (globalState).f6 = (self).f12

    def m20(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_879_v0_: _dafny.MultiSet
        d_879_v0_ = _dafny.MultiSet([(self).f13, (self).f13])
        d_880_v1_: _dafny.MultiSet
        d_880_v1_ = _dafny.MultiSet([p0])
        d_881_v2_: _dafny.Map
        d_881_v2_ = _dafny.Map({p0: 736})
        d_882_v3_: _dafny.Seq
        d_882_v3_ = _dafny.SeqWithoutIsStrInference([True, p3])
        d_883_v4_: str
        d_883_v4_ = _dafny.CodePoint('h')
        d_884_v5_: _dafny.Array
        nw140_ = _dafny.Array(None, 26)
        nw140_[int(0)] = (self).f12
        nw140_[int(1)] = (self).f12
        nw140_[int(2)] = (self).f12
        nw140_[int(3)] = default__.fm0(p0, globalState)
        nw140_[int(4)] = self.f27
        nw140_[int(5)] = self.f27
        nw140_[int(6)] = (0) - ((self).f12)
        nw140_[int(7)] = (d_879_v0_).cardinality
        nw140_[int(8)] = (self).f11
        nw140_[int(9)] = 226
        nw140_[int(10)] = self.f27
        nw140_[int(11)] = ((d_880_v1_)[p3] if (p3) in (d_880_v1_) else len((self).f13))
        nw140_[int(12)] = len(d_881_v2_)
        nw140_[int(13)] = (self).f12
        nw140_[int(14)] = len(p1)
        nw140_[int(15)] = self.f27
        nw140_[int(16)] = self.f27
        nw140_[int(17)] = (d_880_v1_).cardinality
        nw140_[int(18)] = (self).fm40(self.f27, (d_882_v3_)[default__.safeIndex(303, len(d_882_v3_))], d_883_v4_, p0, globalState)
        nw140_[int(19)] = self.f27
        nw140_[int(20)] = len(p1)
        nw140_[int(21)] = (self).f12
        nw140_[int(22)] = self.f27
        nw140_[int(23)] = len(p1)
        nw140_[int(24)] = (d_880_v1_).cardinality
        nw140_[int(25)] = len((self).f13)
        d_884_v5_ = nw140_
        d_885_v6_: D9
        d_885_v6_ = D9_DC27(d_884_v5_)
        source14_ = (d_885_v6_ if p0 else D9_DC27(d_884_v5_))
        if source14_.is_DC27:
            d_886___mcc_h0_ = source14_.cf46
            d_887_cf46_ = d_886___mcc_h0_
            d_888_v7_: D4
            d_888_v7_ = D4_DC15(self.f27)
            d_889_v8_: D4
            d_889_v8_ = D4_DC16(d_888_v7_)
            pat_let_tv10_ = d_888_v7_
            def iife57_(_pat_let14_0):
                def iife58_(d_890_dt__update__tmp_h0_):
                    def iife59_(_pat_let15_0):
                        def iife60_(d_891_dt__update_hcf33_h0_):
                            return D4_DC16(d_891_dt__update_hcf33_h0_)
                        return iife60_(_pat_let15_0)
                    return iife59_(pat_let_tv10_)
                return iife58_(_pat_let14_0)
            source15_ = iife57_(d_889_v8_)
            if source15_.is_DC14:
                d_892___mcc_h2_ = source15_.cf30
                d_893___mcc_h3_ = source15_.cf31
                d_894_cf31_ = d_893___mcc_h3_
                d_895_cf30_ = d_892___mcc_h2_
                d_896_v9_: _dafny.Array
                def lambda34_(d_897_i0_):
                    return False

                init18_ = lambda34_
                nw141_ = _dafny.Array(None, 1)
                for i0_18_ in range(nw141_.length(0)):
                    nw141_[i0_18_] = init18_(i0_18_)
                d_896_v9_ = nw141_
                d_898_v10_: _dafny.Seq
                d_898_v10_ = _dafny.SeqWithoutIsStrInference([d_896_v9_])
                rhs136_ = len(d_898_v10_)
                rhs137_ = not (p3) or (p0)
                lhs103_ = self
                lhs104_ = globalState
                lhs103_.f27 = rhs136_
                lhs104_.f7 = rhs137_
                r0 = default__.safeDivisionInt((self).f12, d_894_cf31_)
                index158_ = default__.safeIndex(123, (d_896_v9_).length(0))
                (d_896_v9_)[index158_] = p0
                d_899_v11_: _dafny.Set
                d_899_v11_ = _dafny.Set({(self).f12})
                d_900_v12_: D14
                d_900_v12_ = D14_DC40(d_882_v3_)
                d_901_v13_: _dafny.Map
                d_901_v13_ = _dafny.Map({(self).f11: (self).fm41(globalState)})
                d_902_v14_: D0
                d_902_v14_ = D0_DC1(p3, 338)
                d_903_v15_: _dafny.MultiSet
                d_903_v15_ = _dafny.MultiSet([(self).f11])
                index159_ = default__.safeIndex(123, (d_896_v9_).length(0))
                rhs138_ = default__.safeDivisionInt(580, len(d_899_v11_))
                rhs139_ = ((self).fm41(globalState)) in ((d_900_v12_).cf72)
                rhs140_ = ((d_901_v13_)[(self).f11] if ((self).f11) in (d_901_v13_) else (d_902_v14_).cf1)
                rhs141_ = self.f27
                rhs142_ = (d_903_v15_).issubset(d_903_v15_)
                lhs105_ = d_896_v9_
                lhs106_ = default__.safeIndex(123, (d_896_v9_).length(0))
                lhs107_ = globalState
                lhs108_ = self
                lhs109_ = globalState
                r0 = rhs138_
                lhs105_[lhs106_] = rhs139_
                lhs107_.f2 = rhs140_
                lhs108_.f27 = rhs141_
                lhs109_.f2 = rhs142_
                d_901_v13_ = (d_901_v13_).set(self.f27, p3)
            elif source15_.is_DC15:
                d_904___mcc_h4_ = source15_.cf32
                d_905_cf32_ = d_904___mcc_h4_
                d_906_v16_: _dafny.Set
                d_906_v16_ = _dafny.Set({d_887_cf46_, d_884_v5_, d_884_v5_})
                rhs143_ = (d_906_v16_) | (d_906_v16_)
                rhs144_ = (self).f11
                lhs110_ = self
                d_906_v16_ = rhs143_
                lhs110_.f27 = rhs144_
                d_907_v17_: _dafny.Seq
                d_907_v17_ = _dafny.SeqWithoutIsStrInference([475])
                d_908_v18_: D2
                d_908_v18_ = D2_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))
                d_909_v19_: _dafny.Seq
                d_909_v19_ = _dafny.SeqWithoutIsStrInference([(d_907_v17_) + (_dafny.SeqWithoutIsStrInference([len((d_908_v18_).cf9), (self).f12])), (_dafny.SeqWithoutIsStrInference([(0) - ((self).f12) for d_910_i1_ in range(default__.abs(866))])) + (d_907_v17_), d_907_v17_, d_907_v17_])
                d_911_v20_: _dafny.Map
                d_911_v20_ = _dafny.Map({975: ((d_881_v2_)[p0] if (p0) in (d_881_v2_) else (0) - ((0) - (len(p1))))})
                d_912_v21_: _dafny.Map
                d_912_v21_ = _dafny.Map({default__.fm50(globalState): d_911_v20_})
                index160_ = default__.safeIndex(795, (d_887_cf46_).length(0))
                (d_887_cf46_)[index160_] = self.f27
                index161_ = default__.safeIndex(795, (d_887_cf46_).length(0))
                rhs145_ = _dafny.SeqWithoutIsStrInference([d_907_v17_])
                rhs146_ = d_912_v21_
                rhs147_ = ((self).fm40((self).f11, p0, d_883_v4_, (D14_DC41(p0, p3, (self).f13)).cf74, globalState)) + (((self).f11) - ((self).f12))
                rhs148_ = d_905_cf32_
                rhs149_ = (0) - (d_905_cf32_)
                lhs111_ = globalState
                lhs112_ = globalState
                lhs113_ = d_887_cf46_
                lhs114_ = default__.safeIndex(795, (d_887_cf46_).length(0))
                d_909_v19_ = rhs145_
                d_912_v21_ = rhs146_
                lhs111_.f6 = rhs147_
                lhs112_.f6 = rhs148_
                lhs113_[lhs114_] = rhs149_
                d_913_v22_: D6
                d_913_v22_ = D6_DC20((self).f12, (d_887_cf46_)[default__.safeIndex(795, (d_887_cf46_).length(0))])
                d_914_v23_: _dafny.MultiSet
                d_914_v23_ = _dafny.MultiSet([(self).f12])
                d_915_v24_: _dafny.Map
                d_915_v24_ = _dafny.Map({p3: p3})
                d_916_v25_: _dafny.Array
                nw142_ = _dafny.Array(None, 21)
                nw142_[int(0)] = p3
                nw142_[int(1)] = (d_882_v3_)[default__.safeIndex((d_913_v22_).cf38, len(d_882_v3_))]
                nw142_[int(2)] = p0
                nw142_[int(3)] = p3
                nw142_[int(4)] = p0
                nw142_[int(5)] = (_dafny.MultiSet([(d_887_cf46_)[default__.safeIndex(795, (d_887_cf46_).length(0))]])).ispropersubset(d_914_v23_)
                nw142_[int(6)] = p0
                nw142_[int(7)] = p0
                nw142_[int(8)] = p0
                nw142_[int(9)] = p3
                nw142_[int(10)] = p0
                nw142_[int(11)] = p0
                nw142_[int(12)] = p3
                nw142_[int(13)] = p3
                nw142_[int(14)] = p0
                nw142_[int(15)] = p3
                nw142_[int(16)] = (-613) != ((0) - (len((self).f13)))
                nw142_[int(17)] = p3
                nw142_[int(18)] = not(p0)
                nw142_[int(19)] = (p0) in (d_915_v24_)
                nw142_[int(20)] = p0
                d_916_v25_ = nw142_
                index162_ = default__.safeIndex(448, (d_916_v25_).length(0))
                (d_916_v25_)[index162_] = ((d_887_cf46_)[default__.safeIndex(795, (d_887_cf46_).length(0))]) == ((d_907_v17_)[default__.safeIndex(self.f27, len(d_907_v17_))])
                index163_ = default__.safeIndex(448, (d_916_v25_).length(0))
                (d_916_v25_)[index163_] = (self).fm41(globalState)
                (globalState).f2 = (d_905_cf32_) <= (len(d_907_v17_))
            elif source15_.is_DC13:
                d_917___mcc_h5_ = source15_.cf29
                d_918_cf29_ = d_917___mcc_h5_
                r0 = len(default__.fm51(p0, (self).f11, (self).f11, (self).f11, globalState))
                (globalState).f2 = p3
                d_919_v26_: _dafny.Map
                d_919_v26_ = _dafny.Map({(self).fm40(569, p0, _dafny.CodePoint('d'), p0, globalState): len((D4_DC13(d_918_cf29_)).cf29)})
                d_919_v26_ = (d_919_v26_).set((self).f12, default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([775])), (self).f11))
                d_920_v27_: _dafny.Map
                d_920_v27_ = _dafny.Map({p3: p0})
                d_921_v28_: D1
                d_921_v28_ = D1_DC5(self.f27, not(((d_920_v27_)[p3] if (p3) in (d_920_v27_) else not(p3))), p0)
                (globalState).f7 = (d_921_v28_).cf8
            elif True:
                d_922___mcc_h6_ = source15_.cf33
                d_923_cf33_ = d_922___mcc_h6_
                (globalState).f2 = ((self).f12) <= (default__.safeModuloInt((self).f11, (self).f11))
                d_924_v29_: _dafny.Array
                nw143_ = _dafny.Array(False, 27)
                d_924_v29_ = nw143_
                d_925_v30_: _dafny.Set
                d_925_v30_ = _dafny.Set({(self).f11, self.f27})
                d_926_v31_: _dafny.Map
                d_926_v31_ = _dafny.Map({p3: (self).f13})
                d_927_v32_: C1
                nw144_ = C1()
                nw144_.ctor__(d_924_v29_, len(d_925_v30_), (self).f12, 721, p3, ((d_926_v31_)[p3] if (p3) in (d_926_v31_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrtckefl"))))
                d_927_v32_ = nw144_
                (globalState).f7 = p3
                (globalState).f2 = not(p0)
            (globalState).f2 = not((self).fm41(globalState))
            d_928_v33_: _dafny.Array
            nw145_ = _dafny.Array(False, 5)
            d_928_v33_ = nw145_
            d_929_v34_: D5
            d_929_v34_ = D5_DC18(_dafny.Map({d_928_v33_: self.f27}), len(default__.fm48(globalState)))
            d_930_v35_: _dafny.Map
            d_930_v35_ = _dafny.Map({(0) - ((self).f12): d_929_v34_})
            d_930_v35_ = (d_930_v35_).set(len((self).f13), d_929_v34_)
            (globalState).f6 = len((((self).f13) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcrtaml"))) + ((self).f13))).set(default__.safeIndex((default__.fm0(p3, globalState)) - ((self).f11), len(((self).f13) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcrtaml"))) + ((self).f13)))), d_883_v4_))
        elif True:
            d_931___mcc_h1_ = source14_.cf45
            d_932_cf45_ = d_931___mcc_h1_
            d_933_v36_: _dafny.Map
            d_933_v36_ = _dafny.Map({p0: p3})
            d_934_v37_: _dafny.Map
            d_934_v37_ = _dafny.Map({True: d_933_v36_})
            d_881_v2_ = (d_881_v2_).set(p3, len((d_934_v37_) | (_dafny.Map({p0: d_933_v36_}))))
            (globalState).f6 = (self).f11
            d_935_v38_: _dafny.Set
            d_935_v38_ = _dafny.Set({len(d_881_v2_)})
            source16_ = default__.fm52(d_882_v3_, (p0 if p3 else p3), (d_935_v38_).issubset(d_935_v38_), self.f27, globalState)
            if source16_.is_DC20:
                d_936___mcc_h7_ = source16_.cf38
                d_937___mcc_h8_ = source16_.cf39
                d_938_cf39_ = d_937___mcc_h8_
                d_939_cf38_ = d_936___mcc_h7_
                (globalState).f7 = not(p3)
                (globalState).f7 = ((p3) == (p0) if (not(p3) if p3 else p3) else p0)
                (globalState).f2 = p3
                d_940_v39_: _dafny.Map
                d_940_v39_ = _dafny.Map({d_938_cf39_: p0})
                (self).f27 = (0) - (default__.safeDivisionInt(len(d_940_v39_), (d_938_cf39_) * (self.f27)))
            elif source16_.is_DC21:
                d_941___mcc_h9_ = source16_.cf40
                d_942_cf40_ = d_941___mcc_h9_
                d_880_v1_ = _dafny.MultiSet(((d_882_v3_) + (default__.fm43(p0, p3, globalState))) + (((d_882_v3_).set(default__.safeIndex((self).f11, len(d_882_v3_)), p0)) + (d_882_v3_)))
                (globalState).f7 = (d_879_v0_).ispropersubset((d_879_v0_).intersection(d_879_v0_))
                d_943_v40_: _dafny.Set
                d_943_v40_ = _dafny.Set({d_932_cf45_.f24})
                d_944_v41_: _dafny.Map
                d_944_v41_ = _dafny.Map({len((d_943_v40_) - (d_943_v40_)): 843})
                (globalState).f6 = ((d_944_v41_)[(self).f12] if ((self).f12) in (d_944_v41_) else (self).f11)
                r0 = (0) - (((self).f12) + (408))
            elif source16_.is_DC19:
                d_945___mcc_h10_ = source16_.cf37
                d_946_cf37_ = d_945___mcc_h10_
                d_947_v42_: _dafny.Set
                d_947_v42_ = _dafny.Set({(d_933_v36_).set(p3, p0), (d_933_v36_) | (d_933_v36_)})
                d_948_v43_: D2
                d_948_v43_ = D2_DC8(p0, p3, (self).f13)
                d_947_v42_ = _dafny.Set({d_933_v36_, d_933_v36_, d_933_v36_, _dafny.Map({(d_932_cf45_).fm1(-30, p3, p0, globalState): (d_948_v43_).cf16}), d_933_v36_})
                d_884_v5_ = d_884_v5_
                d_949_v44_: _dafny.MultiSet
                d_949_v44_ = _dafny.MultiSet([p1])
                d_950_v45_: _dafny.Set
                d_950_v45_ = _dafny.Set({(self).f13, (self).f13, (self).f13})
                d_951_v46_: _dafny.Map
                d_951_v46_ = _dafny.Map({p0: d_949_v44_})
                d_952_v47_: _dafny.Map
                d_952_v47_ = _dafny.Map({d_884_v5_: p3})
                d_953_v48_: _dafny.Seq
                d_953_v48_ = _dafny.SeqWithoutIsStrInference([d_952_v47_])
                d_954_v49_: _dafny.Map
                d_954_v49_ = _dafny.Map({len((d_953_v48_)[default__.safeIndex((self).f11, len(d_953_v48_))]): p3})
                d_955_v50_: _dafny.Seq
                d_955_v50_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p3, p3, p3}), p1, p1])
                rhs150_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ba"))) in (d_950_v45_)
                rhs151_ = d_882_v3_
                rhs152_ = p0
                rhs153_ = (default__.fm53(True, not(p0), (self).f12, p1, globalState)).intersection(((d_951_v46_)[((d_954_v49_)[self.f27] if (self.f27) in (d_954_v49_) else False)] if (((d_954_v49_)[self.f27] if (self.f27) in (d_954_v49_) else False)) in (d_951_v46_) else _dafny.MultiSet(d_955_v50_)))
                lhs115_ = globalState
                lhs116_ = globalState
                lhs115_.f2 = rhs150_
                d_882_v3_ = rhs151_
                lhs116_.f7 = rhs152_
                d_949_v44_ = rhs153_
                (globalState).f7 = p0
            elif True:
                d_956___mcc_h11_ = source16_.cf41
                d_957_cf41_ = d_956___mcc_h11_
                d_958_v51_: _dafny.Array
                def lambda35_(d_959_i2_):
                    return _dafny.Map({(self).f11: (self).f11})

                init19_ = lambda35_
                nw146_ = _dafny.Array(None, 6)
                for i0_19_ in range(nw146_.length(0)):
                    nw146_[i0_19_] = init19_(i0_19_)
                d_958_v51_ = nw146_
                d_960_v52_: _dafny.Map
                d_960_v52_ = _dafny.Map({self.f27: self.f27})
                index164_ = default__.safeIndex(147, (d_958_v51_).length(0))
                (d_958_v51_)[index164_] = d_960_v52_
                index165_ = default__.safeIndex(147, (d_958_v51_).length(0))
                (d_958_v51_)[index165_] = d_960_v52_
                d_961_v53_: D0
                d_961_v53_ = D0_DC0(_dafny.MultiSet((self).f13))
                d_962_v54_: _dafny.Map
                d_962_v54_ = _dafny.Map({self.f27: d_961_v53_})
                d_963_v55_: D1
                d_963_v55_ = D1_DC4(p0)
                d_962_v54_ = (d_962_v54_).set(default__.safeModuloInt(len((self).f13), (self).f12), (d_961_v53_ if (d_963_v55_).cf5 else d_961_v53_))
                index166_ = default__.safeIndex(492, (d_884_v5_).length(0))
                (d_884_v5_)[index166_] = (self).f12
                index167_ = default__.safeIndex(492, (d_884_v5_).length(0))
                (d_884_v5_)[index167_] = ((d_880_v1_)[p3] if (p3) in (d_880_v1_) else default__.safeModuloInt(self.f27, self.f27))
                (globalState).f7 = not(not (p3) or (not(p0)))
            arr22_ = d_932_cf45_.f24
            index168_ = default__.safeIndex(347, (d_932_cf45_.f24).length(0))
            arr22_[index168_] = not(p3)
            d_964_v56_: _dafny.Map
            d_964_v56_ = _dafny.Map({self.f27: d_881_v2_})
            arr23_ = d_932_cf45_.f24
            index169_ = default__.safeIndex(347, (d_932_cf45_.f24).length(0))
            arr23_[index169_] = (d_932_cf45_).fm1(len(d_964_v56_), False, True, globalState)
        (globalState).f6 = (self).f11
        d_965_v57_: _dafny.Array
        nw147_ = _dafny.Array(_dafny.Map({}), 29)
        d_965_v57_ = nw147_
        d_966_v58_: D6
        d_966_v58_ = D6_DC20((self).f12, len(_dafny.Map({p0: len(default__.fm44((d_880_v1_).cardinality, p0, globalState))})))
        d_967_v59_: D6
        d_967_v59_ = D6_DC22(d_966_v58_)
        d_968_v60_: _dafny.Map
        d_968_v60_ = _dafny.Map({d_967_v59_: p3})
        index170_ = default__.safeIndex(488, (d_965_v57_).length(0))
        (d_965_v57_)[index170_] = d_968_v60_
        index171_ = default__.safeIndex(488, (d_965_v57_).length(0))
        (d_965_v57_)[index171_] = ((d_968_v60_) | (d_968_v60_)) | ((d_968_v60_).set(d_967_v59_, p0))
        d_969_v61_: _dafny.MultiSet
        d_969_v61_ = _dafny.MultiSet([d_884_v5_])
        (globalState).f7 = (d_969_v61_).isdisjoint(d_969_v61_)
        d_970_v62_: _dafny.Map
        d_970_v62_ = _dafny.Map({(0) - (self.f27): -413})
        d_971_v63_: _dafny.Array
        nw148_ = _dafny.Array(None, 21)
        nw148_[int(0)] = p3
        nw148_[int(1)] = p0
        nw148_[int(2)] = p3
        nw148_[int(3)] = p0
        nw148_[int(4)] = False
        nw148_[int(5)] = not(p3)
        nw148_[int(6)] = p0
        nw148_[int(7)] = not(not(p0))
        nw148_[int(8)] = p0
        nw148_[int(9)] = p3
        nw148_[int(10)] = p3
        nw148_[int(11)] = (self).fm41(globalState)
        nw148_[int(12)] = p0
        nw148_[int(13)] = p0
        nw148_[int(14)] = p3
        nw148_[int(15)] = True
        nw148_[int(16)] = p0
        nw148_[int(17)] = p0
        nw148_[int(18)] = p0
        nw148_[int(19)] = p3
        nw148_[int(20)] = p0
        d_971_v63_ = nw148_
        (globalState).f6 = (self).fm40(((d_970_v62_)[(_dafny.MultiSet([d_971_v63_, d_971_v63_])).cardinality] if ((_dafny.MultiSet([d_971_v63_, d_971_v63_])).cardinality) in (d_970_v62_) else (self).f12), p0, _dafny.CodePoint('y'), p0, globalState)
        (globalState).f2 = ((self).f12) > ((self).f11)
        r0 = default__.safeModuloInt((self).fm40(56, p0, _dafny.CodePoint('e'), p0, globalState), ((self).f11 if p3 else -366))
        return r0


class C3(T3, T2, T1, T0):
    def  __init__(self):
        self._f16: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f11: int = int(0)
        self._f12: int = int(0)
        self._f14: int = int(0)
        self._f15: bool = False
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f25: bool = False
        self._f26: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
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
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f25, f26, f16, f14, f15, f13, f11, f12):
        (self)._f25 = f25
        (self)._f26 = f26
        (self).f16 = f16
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f13 = f13
        (self)._f11 = f11
        (self)._f12 = f12

    def fm2(self, p0, p1, p2, p3, globalState):
        if (self).f15:
            return len(_dafny.Map({False: (self).f12}))
        elif True:
            return default__.safeModuloInt((0) - ((self).f11), len(_dafny.Map({(self).f11: (self).f11})))

    def fm3(self, p0, p1, p2, globalState):
        if False:
            return (_dafny.Map({not((self).f25): (self).f14})) | (_dafny.Map({(self).f25: len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_972_i1_ in range(default__.abs(-566))]) for d_973_i0_ in range(default__.abs(403))]))}))
        elif True:
            return (_dafny.Map({(self).f25: (_dafny.MultiSet([(self).f12, (self).f14])).cardinality})) | (_dafny.Map({(self).f25: (self).f12}))

    def fm1(self, p0, p1, p2, globalState):
        def iife61_():
            coll29_ = _dafny.Set()
            compr_29_: _dafny.MultiSet
            for compr_29_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f11])])).Elements:
                d_974_v0_: _dafny.MultiSet = compr_29_
                if (d_974_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f11])])):
                    coll29_ = coll29_.union(_dafny.Set([d_974_v0_]))
            return _dafny.Set(coll29_)
        return (_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f11]))})).ispropersubset((iife61_()
        ) - (_dafny.Set({_dafny.MultiSet([(D3_DC11(len((self).f13), D2_DC8(True, False, (self).f13), _dafny.Set({D0_DC0(_dafny.MultiSet([_dafny.CodePoint('u')])), D0_DC0(_dafny.MultiSet([_dafny.CodePoint('f'), _dafny.CodePoint('t')]))}), (self).f25)).cf24])})))

    def fm26(self, p0, globalState):
        if (self).f15:
            return (self).f25
        elif True:
            return (self).f15

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: T0 = None
        d_975_v0_: _dafny.MultiSet
        d_975_v0_ = _dafny.MultiSet([(self).f25])
        d_976_v1_: _dafny.Seq
        d_976_v1_ = _dafny.SeqWithoutIsStrInference([d_975_v0_, d_975_v0_])
        d_977_v2_: _dafny.MultiSet
        d_977_v2_ = _dafny.MultiSet([(0) - (len(d_976_v1_)), (self).f11])
        d_978_v3_: _dafny.Seq
        d_978_v3_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        d_979_v4_: str
        d_979_v4_ = _dafny.CodePoint('n')
        d_980_v5_: D1
        d_980_v5_ = D1_DC3(False)
        d_981_v6_: _dafny.Set
        d_981_v6_ = _dafny.Set({(self).f12})
        d_982_v7_: _dafny.Array
        nw149_ = _dafny.Array(None, 19)
        nw149_[int(0)] = (self).f25
        nw149_[int(1)] = (d_977_v2_).isdisjoint(d_977_v2_)
        nw149_[int(2)] = not((d_978_v3_) != (_dafny.SeqWithoutIsStrInference([p1, True, not(not((self).f25)), (self).f15])))
        nw149_[int(3)] = (d_979_v4_) not in (self.f16)
        nw149_[int(4)] = (d_980_v5_).cf4
        nw149_[int(5)] = True
        nw149_[int(6)] = (d_978_v3_) != (d_978_v3_)
        nw149_[int(7)] = True
        nw149_[int(8)] = True
        nw149_[int(9)] = (d_981_v6_).isdisjoint(d_981_v6_)
        nw149_[int(10)] = p1
        nw149_[int(11)] = (self).f15
        nw149_[int(12)] = ((self).f15) == (not((self).f25))
        nw149_[int(13)] = ((self).f11) <= ((self).f14)
        nw149_[int(14)] = (self).f15
        nw149_[int(15)] = True
        nw149_[int(16)] = (self).fm1((self).f12, (self).f25, p1, globalState)
        nw149_[int(17)] = (self).f25
        nw149_[int(18)] = (self).fm1((self).f14, p1, p1, globalState)
        d_982_v7_ = nw149_
        index172_ = default__.safeIndex(55, (d_982_v7_).length(0))
        (d_982_v7_)[index172_] = (self).fm1((0) - ((self).f12), (self).f15, True, globalState)
        d_983_v9_: _dafny.Seq
        def iife62_():
            coll30_ = _dafny.Set()
            compr_30_: int
            for compr_30_ in _dafny.IntegerRange(-339, 906):
                d_984_v8_: int = compr_30_
                if ((-339) <= (d_984_v8_)) and ((d_984_v8_) < (906)):
                    coll30_ = coll30_.union(_dafny.Set([default__.safeModuloInt(d_984_v8_, (self).f11)]))
            return _dafny.Set(coll30_)
        d_983_v9_ = _dafny.SeqWithoutIsStrInference([iife62_()
        ])
        pat_let_tv11_ = p1
        pat_let_tv12_ = d_977_v2_
        pat_let_tv13_ = p1
        index173_ = default__.safeIndex(55, (d_982_v7_).length(0))
        def lambda36_(source17_):
            if source17_.is_DC10:
                d_985___mcc_h0_ = source17_.cf19
                d_986___mcc_h1_ = source17_.cf20
                d_987___mcc_h2_ = source17_.cf21
                d_988___mcc_h3_ = source17_.cf22
                d_989___mcc_h4_ = source17_.cf23
                d_990_cf23_ = d_989___mcc_h4_
                d_991_cf22_ = d_988___mcc_h3_
                d_992_cf21_ = d_987___mcc_h2_
                d_993_cf20_ = d_986___mcc_h1_
                d_994_cf19_ = d_985___mcc_h0_
                return pat_let_tv11_
            elif source17_.is_DC11:
                d_995___mcc_h5_ = source17_.cf24
                d_996___mcc_h6_ = source17_.cf25
                d_997___mcc_h7_ = source17_.cf26
                d_998___mcc_h8_ = source17_.cf27
                d_999_cf27_ = d_998___mcc_h8_
                d_1000_cf26_ = d_997___mcc_h7_
                d_1001_cf25_ = d_996___mcc_h6_
                d_1002_cf24_ = d_995___mcc_h5_
                return ((self).f11) not in (_dafny.Map({(0) - ((self).f11): _dafny.Map({(pat_let_tv12_).cardinality: (self).f12})}))
            elif source17_.is_DC9:
                d_1003___mcc_h9_ = source17_.cf18
                d_1004_cf18_ = d_1003___mcc_h9_
                return not ((self).f25) or ((self).f25)
            elif True:
                d_1005___mcc_h10_ = source17_.cf28
                d_1006_cf28_ = d_1005___mcc_h10_
                return not(pat_let_tv13_)

        (d_982_v7_)[index173_] = lambda36_(default__.fm27(len(d_983_v9_), (self).f15, globalState))
        d_1007_v10_: _dafny.Map
        d_1007_v10_ = _dafny.Map({d_978_v3_: p1})
        d_1008_i0_: int
        d_1008_i0_ = 0
        with _dafny.label("5"):
            while (882) <= (default__.fm0((d_978_v3_)[default__.safeIndex(len(d_1007_v10_), len(d_978_v3_))], globalState)):
                with _dafny.c_label("5"):
                    if (d_1008_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_1008_i0_ = (d_1008_i0_) + (1)
                    d_1009_v11_: C0
                    nw150_ = C0()
                    nw150_.ctor__()
                    d_1009_v11_ = nw150_
                    (globalState).f7 = (d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))]
                    d_1010_v12_: _dafny.Set
                    d_1010_v12_ = _dafny.Set({p1})
                    d_1011_v13_: _dafny.Array
                    nw151_ = _dafny.Array(None, 13)
                    nw151_[int(0)] = (self).f11
                    nw151_[int(1)] = (self).f14
                    nw151_[int(2)] = (self).f11
                    nw151_[int(3)] = (self).f12
                    nw151_[int(4)] = len((self).f13)
                    nw151_[int(5)] = (self).f11
                    nw151_[int(6)] = default__.safeModuloInt(len(d_1010_v12_), (self).f11)
                    nw151_[int(7)] = (self).f14
                    nw151_[int(8)] = (self).f12
                    nw151_[int(9)] = (self).f11
                    nw151_[int(10)] = len((self).fm3((self).f14, not((d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))]), p1, globalState))
                    nw151_[int(11)] = 189
                    nw151_[int(12)] = (self).f14
                    d_1011_v13_ = nw151_
                    index174_ = default__.safeIndex(425, (d_1011_v13_).length(0))
                    (d_1011_v13_)[index174_] = (self).f14
                    d_1012_v14_: D4
                    d_1012_v14_ = D4_DC14(d_979_v4_, (self).fm2((self).f12, p1, (self).f14, p1, globalState))
                    index175_ = default__.safeIndex(425, (d_1011_v13_).length(0))
                    (d_1011_v13_)[index175_] = (d_1012_v14_).cf31
                    d_1013_v15_: _dafny.Map
                    d_1013_v15_ = _dafny.Map({True: (self).f15})
                    d_1014_v16_: D6
                    d_1014_v16_ = D6_DC20((self).f12, len((d_1013_v15_) | (d_1013_v15_)))
                    source18_ = d_1014_v16_
                    if source18_.is_DC20:
                        d_1015___mcc_h11_ = source18_.cf38
                        d_1016___mcc_h12_ = source18_.cf39
                        d_1017_cf39_ = d_1016___mcc_h12_
                        d_1018_cf38_ = d_1015___mcc_h11_
                        d_1019_v17_: _dafny.Array
                        nw152_ = _dafny.Array(_dafny.Map({}), 2)
                        d_1019_v17_ = nw152_
                        d_1020_v18_: _dafny.MultiSet
                        d_1020_v18_ = _dafny.MultiSet([d_979_v4_])
                        d_1021_v19_: _dafny.Map
                        d_1021_v19_ = _dafny.Map({D1_DC4((self).f15): (d_1020_v18_).cardinality})
                        index176_ = default__.safeIndex(694, (d_1019_v17_).length(0))
                        (d_1019_v17_)[index176_] = d_1021_v19_
                        index177_ = default__.safeIndex(694, (d_1019_v17_).length(0))
                        (d_1019_v17_)[index177_] = d_1021_v19_
                        d_1022_v20_: D5
                        d_1022_v20_ = D5_DC18(p0, (_dafny.MultiSet([d_1018_cf38_])).cardinality)
                        d_1023_v21_: _dafny.MultiSet
                        d_1023_v21_ = _dafny.MultiSet([d_1022_v20_])
                        index178_ = default__.safeIndex(425, (d_1011_v13_).length(0))
                        index179_ = default__.safeIndex(55, (d_982_v7_).length(0))
                        rhs154_ = (d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))]
                        rhs155_ = d_1017_cf39_
                        rhs156_ = (d_1011_v13_)[default__.safeIndex(425, (d_1011_v13_).length(0))]
                        rhs157_ = (d_1022_v20_) in (d_1023_v21_)
                        lhs117_ = globalState
                        lhs118_ = d_1011_v13_
                        lhs119_ = default__.safeIndex(425, (d_1011_v13_).length(0))
                        lhs120_ = d_982_v7_
                        lhs121_ = default__.safeIndex(55, (d_982_v7_).length(0))
                        lhs117_.f7 = rhs154_
                        d_1018_cf38_ = rhs155_
                        lhs118_[lhs119_] = rhs156_
                        lhs120_[lhs121_] = rhs157_
                        def lambda37_(d_1024_cf38_):
                            def lambda38_(d_1025_i1_):
                                return default__.safeModuloInt(d_1025_i1_, d_1024_cf38_)

                            return lambda38_

                        init20_ = lambda37_(d_1018_cf38_)
                        nw153_ = _dafny.Array(None, 29)
                        for i0_20_ in range(nw153_.length(0)):
                            nw153_[i0_20_] = init20_(i0_20_)
                        d_1011_v13_ = nw153_
                        d_1026_v22_: _dafny.Map
                        d_1026_v22_ = _dafny.Map({len((self).f26): (d_975_v0_).ispropersubset(d_975_v0_)})
                        d_1026_v22_ = (d_1026_v22_).set((0) - (default__.safeDivisionInt(d_1017_cf39_, (d_1011_v13_)[default__.safeIndex(425, (d_1011_v13_).length(0))])), p1)
                    elif source18_.is_DC21:
                        d_1027___mcc_h13_ = source18_.cf40
                        d_1028_cf40_ = d_1027___mcc_h13_
                        index180_ = default__.safeIndex(425, (d_1011_v13_).length(0))
                        rhs158_ = (self).f14
                        rhs159_ = (self).fm2(23, False, (d_1011_v13_)[default__.safeIndex(425, (d_1011_v13_).length(0))], (self).f15, globalState)
                        lhs122_ = d_1011_v13_
                        lhs123_ = default__.safeIndex(425, (d_1011_v13_).length(0))
                        lhs124_ = globalState
                        lhs122_[lhs123_] = rhs158_
                        lhs124_.f6 = rhs159_
                        d_1029_v23_: _dafny.Seq
                        d_1029_v23_ = _dafny.SeqWithoutIsStrInference([433])
                        index181_ = default__.safeIndex(55, (d_982_v7_).length(0))
                        rhs160_ = d_1011_v13_
                        rhs161_ = _dafny.SeqWithoutIsStrInference([(d_1011_v13_)[default__.safeIndex(425, (d_1011_v13_).length(0))] for d_1030_i2_ in range(default__.abs(605))])
                        rhs162_ = d_982_v7_
                        rhs163_ = (d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))]
                        lhs125_ = d_982_v7_
                        lhs126_ = default__.safeIndex(55, (d_982_v7_).length(0))
                        d_1011_v13_ = rhs160_
                        d_1029_v23_ = rhs161_
                        d_982_v7_ = rhs162_
                        lhs125_[lhs126_] = rhs163_
                        d_1031_v24_: D6
                        d_1031_v24_ = D6_DC19((self).fm3((self).f12, False, (d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))], globalState))
                        d_1032_v25_: D6
                        d_1032_v25_ = D6_DC22(d_1031_v24_)
                        rhs164_ = (d_975_v0_ if d_1028_cf40_ else _dafny.MultiSet(d_978_v3_))
                        rhs165_ = d_1028_cf40_
                        rhs166_ = d_1032_v25_
                        lhs127_ = globalState
                        d_975_v0_ = rhs164_
                        lhs127_.f7 = rhs165_
                        d_1032_v25_ = rhs166_
                        d_1033_v26_: D1
                        d_1033_v26_ = D1_DC5((D5_DC18(p0, len(d_978_v3_))).cf36, (d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))], p1)
                        d_1034_v27_: D1
                        d_1034_v27_ = D1_DC2(_dafny.CodePoint('p'))
                        d_1035_v28_: _dafny.Array
                        nw154_ = _dafny.Array(None, 3)
                        nw154_[int(0)] = D1_DC2(d_979_v4_)
                        nw154_[int(1)] = default__.fm28(not((self).f15), d_1028_cf40_, (self).fm26(True, globalState), d_1033_v26_, globalState)
                        nw154_[int(2)] = d_1034_v27_
                        d_1035_v28_ = nw154_
                        index182_ = default__.safeIndex(491, (d_1035_v28_).length(0))
                        (d_1035_v28_)[index182_] = d_1034_v27_
                        index183_ = default__.safeIndex(491, (d_1035_v28_).length(0))
                        (d_1035_v28_)[index183_] = d_1034_v27_
                    elif source18_.is_DC19:
                        d_1036___mcc_h14_ = source18_.cf37
                        d_1037_cf37_ = d_1036___mcc_h14_
                        index184_ = default__.safeIndex(55, (d_982_v7_).length(0))
                        (d_982_v7_)[index184_] = ((101) * ((self).fm2((self).f12, (d_978_v3_)[default__.safeIndex((self).f11, len(d_978_v3_))], (d_1011_v13_)[default__.safeIndex(425, (d_1011_v13_).length(0))], (self).f15, globalState))) not in (d_981_v6_)
                        (globalState).f7 = (d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))]
                        (globalState).f7 = (True) in (d_1010_v12_)
                        (globalState).f6 = (self).f11
                    elif True:
                        d_1038___mcc_h15_ = source18_.cf41
                        d_1039_cf41_ = d_1038___mcc_h15_
                        r1 = not(not((self).f25))
                        d_1040_v29_: C1
                        nw155_ = C1()
                        nw155_.ctor__(d_982_v7_, (self).f14, (d_1011_v13_)[default__.safeIndex(425, (d_1011_v13_).length(0))], len(_dafny.Map({d_978_v3_: (self).f25})), True, ((self).f26).set(default__.safeIndex((d_1011_v13_)[default__.safeIndex(425, (d_1011_v13_).length(0))], len((self).f26)), _dafny.CodePoint('c')))
                        d_1040_v29_ = nw155_
                        d_1041_v30_: _dafny.Seq
                        d_1041_v30_ = _dafny.SeqWithoutIsStrInference([d_1040_v29_])
                        d_1041_v30_ = (d_1041_v30_) + (d_1041_v30_)
                        (globalState).f2 = (6) >= ((self).f12)
                        d_1042_v31_: _dafny.Seq
                        d_1042_v31_ = _dafny.SeqWithoutIsStrInference([d_978_v3_])
                        rhs167_ = d_1011_v13_
                        rhs168_ = d_1042_v31_
                        rhs169_ = d_1040_v29_.f24
                        rhs170_ = (self).f14
                        rhs171_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
                        lhs128_ = globalState
                        lhs129_ = self
                        d_1011_v13_ = rhs167_
                        d_1042_v31_ = rhs168_
                        d_982_v7_ = rhs169_
                        lhs128_.f6 = rhs170_
                        lhs129_.f16 = rhs171_
                    pass
            pass
        d_1043_v32_: _dafny.Array
        nw156_ = _dafny.Array(_dafny.MultiSet({}), 18)
        d_1043_v32_ = nw156_
        rhs172_ = (self).f14
        rhs173_ = (_dafny.MultiSet([(self).f15, not(p1), (self).f15, ((self).f12) == ((self).f12)])).cardinality
        rhs174_ = d_1043_v32_
        lhs130_ = globalState
        lhs130_.f6 = rhs172_
        r0 = rhs173_
        d_1043_v32_ = rhs174_
        if (82) > ((self).f14):
            d_1044_v33_: _dafny.Map
            d_1044_v33_ = _dafny.Map({p1: d_982_v7_})
            d_1045_v34_: _dafny.Seq
            d_1045_v34_ = _dafny.SeqWithoutIsStrInference([((d_1044_v33_)[(self).f25] if ((self).f25) in (d_1044_v33_) else d_982_v7_), d_982_v7_])
            (globalState).f7 = not ((d_1045_v34_) < (d_1045_v34_)) or (p1)
            if True:
                d_1046_v35_: _dafny.Array
                nw157_ = _dafny.Array(int(0), 12)
                d_1046_v35_ = nw157_
                d_1047_v36_: _dafny.Array
                d_1048_v37_: _dafny.Map
                d_1049_v38_: _dafny.Seq
                d_1050_v39_: str
                out36_: _dafny.Array
                out37_: _dafny.Map
                out38_: _dafny.Seq
                out39_: str
                out36_, out37_, out38_, out39_ = (self).m18(d_1046_v35_, (self).f14, globalState)
                d_1047_v36_ = out36_
                d_1048_v37_ = out37_
                d_1049_v38_ = out38_
                d_1050_v39_ = out39_
                (globalState).f2 = not((d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))])
                index185_ = default__.safeIndex(831, (d_1046_v35_).length(0))
                (d_1046_v35_)[index185_] = (self).f11
                index186_ = default__.safeIndex(831, (d_1046_v35_).length(0))
                rhs175_ = (d_978_v3_)[default__.safeIndex(default__.fm0(True, globalState), len(d_978_v3_))]
                rhs176_ = (default__.fm0((d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))], globalState)) * ((self).f11)
                rhs177_ = (712 if (self).f25 else default__.fm0((self).f15, globalState))
                lhs131_ = globalState
                lhs132_ = globalState
                lhs133_ = d_1046_v35_
                lhs134_ = default__.safeIndex(831, (d_1046_v35_).length(0))
                lhs131_.f7 = rhs175_
                lhs132_.f6 = rhs176_
                lhs133_[lhs134_] = rhs177_
                d_1049_v38_ = d_1049_v38_
                d_1051_v40_: C0
                nw158_ = C0()
                nw158_.ctor__()
                d_1051_v40_ = nw158_
            elif True:
                r0 = (self).f12
                d_1052_v41_: _dafny.Seq
                d_1052_v41_ = _dafny.SeqWithoutIsStrInference([len(((d_978_v3_).set(default__.safeIndex((self).f11, len(d_978_v3_)), (d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))])) + (_dafny.SeqWithoutIsStrInference([p1, (self).f25, (self).fm1((self).f12, False, (d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))], globalState)]))), default__.safeModuloInt((self).f14, (self).f12), (self).f14])
                (globalState).f6 = len(d_1052_v41_)
                index187_ = default__.safeIndex(55, (d_982_v7_).length(0))
                (d_982_v7_)[index187_] = p1
                (globalState).f7 = (self).f15
                (globalState).f6 = default__.safeModuloInt((0) - (len(_dafny.SeqWithoutIsStrInference([(self).f12]))), (self).f12)
            d_1053_v42_: _dafny.Array
            def lambda39_(d_1054_i3_):
                return (d_1054_i3_) * ((self).f12)

            init21_ = lambda39_
            nw159_ = _dafny.Array(None, 15)
            for i0_21_ in range(nw159_.length(0)):
                nw159_[i0_21_] = init21_(i0_21_)
            d_1053_v42_ = nw159_
            d_1055_v43_: _dafny.Map
            d_1055_v43_ = _dafny.Map({(self).f25: d_1053_v42_})
            d_1056_v44_: _dafny.Map
            d_1056_v44_ = _dafny.Map({(self).f15: _dafny.Set({(self).f12, len(d_1055_v43_)})})
            d_1057_v45_: _dafny.Array
            nw160_ = _dafny.Array(None, 24)
            nw160_[int(0)] = (d_981_v6_) - (d_981_v6_)
            nw160_[int(1)] = d_981_v6_
            nw160_[int(2)] = d_981_v6_
            nw160_[int(3)] = d_981_v6_
            nw160_[int(4)] = _dafny.Set({(self).f14})
            nw160_[int(5)] = _dafny.Set({(self).f12})
            nw160_[int(6)] = (d_981_v6_) | (_dafny.Set({(self).f14, (self).f12}))
            nw160_[int(7)] = d_981_v6_
            nw160_[int(8)] = (d_981_v6_) - (_dafny.Set({(d_975_v0_).cardinality}))
            nw160_[int(9)] = (_dafny.Set({(self).f14})) | (default__.fm29(globalState))
            nw160_[int(10)] = d_981_v6_
            nw160_[int(11)] = ((d_1056_v44_)[(d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))]] if ((d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))]) in (d_1056_v44_) else (d_983_v9_)[default__.safeIndex((self).f12, len(d_983_v9_))])
            nw160_[int(12)] = _dafny.Set({(self).f14})
            nw160_[int(13)] = d_981_v6_
            nw160_[int(14)] = _dafny.Set({(self).f11, (self).f11})
            nw160_[int(15)] = d_981_v6_
            nw160_[int(16)] = d_981_v6_
            nw160_[int(17)] = (d_981_v6_) | (d_981_v6_)
            nw160_[int(18)] = default__.fm29(globalState)
            nw160_[int(19)] = d_981_v6_
            nw160_[int(20)] = d_981_v6_
            nw160_[int(21)] = default__.fm29(globalState)
            nw160_[int(22)] = d_981_v6_
            nw160_[int(23)] = d_981_v6_
            d_1057_v45_ = nw160_
            index188_ = default__.safeIndex(867, (d_1057_v45_).length(0))
            def iife63_():
                coll31_ = _dafny.Set()
                compr_31_: int
                for compr_31_ in _dafny.IntegerRange(66, 693):
                    d_1058_v46_: int = compr_31_
                    if ((66) <= (d_1058_v46_)) and ((d_1058_v46_) < (693)):
                        coll31_ = coll31_.union(_dafny.Set([(d_1058_v46_) - ((self).f14)]))
                return _dafny.Set(coll31_)
            (d_1057_v45_)[index188_] = iife63_()
            
            index189_ = default__.safeIndex(867, (d_1057_v45_).length(0))
            def iife64_():
                coll32_ = _dafny.Set()
                compr_32_: int
                for compr_32_ in ((d_977_v2_) | (d_977_v2_)).Elements:
                    d_1059_v47_: int = compr_32_
                    if (d_1059_v47_) in ((d_977_v2_) | (d_977_v2_)):
                        coll32_ = coll32_.union(_dafny.Set([(d_1059_v47_) - (168)]))
                return _dafny.Set(coll32_)
            (d_1057_v45_)[index189_] = iife64_()
            
            (self).f16 = (self).f13
            d_1060_v48_: _dafny.Array
            d_1061_v49_: _dafny.Map
            d_1062_v50_: _dafny.Seq
            d_1063_v51_: str
            out40_: _dafny.Array
            out41_: _dafny.Map
            out42_: _dafny.Seq
            out43_: str
            out40_, out41_, out42_, out43_ = (self).m18(d_1053_v42_, default__.fm0(True, globalState), globalState)
            d_1060_v48_ = out40_
            d_1061_v49_ = out41_
            d_1062_v50_ = out42_
            d_1063_v51_ = out43_
        elif True:
            d_1064_v52_: _dafny.Map
            d_1064_v52_ = _dafny.Map({_dafny.CodePoint('f'): 598})
            d_1065_v53_: _dafny.Array
            nw161_ = _dafny.Array(None, 11)
            d_1065_v53_ = nw161_
            d_1066_v54_: _dafny.Set
            d_1066_v54_ = _dafny.Set({d_1065_v53_, d_1065_v53_})
            d_1067_v55_: D1
            d_1067_v55_ = D1_DC2(d_979_v4_)
            d_1068_v56_: _dafny.Map
            d_1068_v56_ = _dafny.Map({(d_982_v7_)[default__.safeIndex(55, (d_982_v7_).length(0))]: p1})
            d_1069_v57_: D2
            d_1069_v57_ = D2_DC7(d_982_v7_, d_1066_v54_, d_1067_v55_, len(d_1068_v56_), 714)
            d_1070_v58_: _dafny.Map
            d_1070_v58_ = _dafny.Map({d_1064_v52_: (self).fm2((d_1069_v57_).cf13, False, (self).f12, p1, globalState)})
            def iife65_():
                coll33_ = _dafny.Map()
                compr_33_: _dafny.Map
                for compr_33_ in (_dafny.SeqWithoutIsStrInference([d_1064_v52_ for d_1071_i4_ in range(default__.abs(396))])).Elements:
                    d_1072_v59_: _dafny.Map = compr_33_
                    if (d_1072_v59_) in (_dafny.SeqWithoutIsStrInference([d_1064_v52_ for d_1071_i4_ in range(default__.abs(396))])):
                        coll33_[d_1072_v59_] = (self).f12
                return _dafny.Map(coll33_)
            d_1070_v58_ = iife65_()
            
            r0 = (self).f14
            d_1073_v60_: _dafny.Map
            d_1073_v60_ = _dafny.Map({(self).f13: default__.safeDivisionInt((self).f14, (self).f11)})
            d_1073_v60_ = (d_1073_v60_).set(((self).f26).set(default__.safeIndex((self).f14, len((self).f26)), d_979_v4_), 450)
            d_1074_v61_: C0
            nw162_ = C0()
            nw162_.ctor__()
            d_1074_v61_ = nw162_
            d_1075_v62_: bool
            d_1076_v63_: _dafny.Seq
            out44_: bool
            out45_: _dafny.Seq
            out44_, out45_ = default__.m0(True, (d_978_v3_)[default__.safeIndex((self).f12, len(d_978_v3_))], (self).f14, default__.safeModuloInt((self).f12, (self).f14), globalState)
            d_1075_v62_ = out44_
            d_1076_v63_ = out45_
        d_1077_v64_: _dafny.Array
        nw163_ = _dafny.Array(int(0), 28)
        d_1077_v64_ = nw163_
        index190_ = default__.safeIndex(982, (d_1077_v64_).length(0))
        (d_1077_v64_)[index190_] = (self).f12
        index191_ = default__.safeIndex(982, (d_1077_v64_).length(0))
        (d_1077_v64_)[index191_] = (self).f12
        (globalState).f7 = (self).f15
        r0 = default__.safeDivisionInt((d_1077_v64_)[default__.safeIndex(982, (d_1077_v64_).length(0))], (self).f11)
        r1 = ((self).f12) > ((self).f11)
        d_1078_v65_: T0
        nw164_ = C1()
        nw164_.ctor__(d_982_v7_, 191, (self).f14, (self).f12, (self).f25, (self).f13)
        d_1078_v65_ = nw164_
        r2 = d_1078_v65_
        return r0, r1, r2

    def m4(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        d_1079_v0_: _dafny.Array
        nw165_ = _dafny.Array(False, 26)
        d_1079_v0_ = nw165_
        d_1079_v0_ = p0
        d_1080_i0_: int
        d_1080_i0_ = 0
        with _dafny.label("6"):
            while (self).f25:
                with _dafny.c_label("6"):
                    if (d_1080_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1080_i0_ = (d_1080_i0_) + (1)
                    d_1081_v1_: _dafny.Seq
                    d_1081_v1_ = _dafny.SeqWithoutIsStrInference([443])
                    d_1082_v2_: _dafny.Map
                    d_1082_v2_ = _dafny.Map({p0: (d_1081_v1_)[default__.safeIndex((self).f11, len(d_1081_v1_))]})
                    d_1083_v3_: D5
                    d_1083_v3_ = D5_DC18(d_1082_v2_, (self).f14)
                    d_1084_v4_: _dafny.MultiSet
                    d_1084_v4_ = _dafny.MultiSet([(self).f25, (self).fm1((self).f11, (self).f15, False, globalState)])
                    d_1085_v5_: str
                    d_1085_v5_ = _dafny.CodePoint('x')
                    d_1086_v6_: _dafny.Set
                    d_1086_v6_ = _dafny.Set({not((self).f15)})
                    d_1087_v7_: _dafny.MultiSet
                    d_1087_v7_ = _dafny.MultiSet([((d_1083_v3_).cf36) + ((0) - (((d_1084_v4_)[(self).f25] if ((self).f25) in (d_1084_v4_) else (0) - (len((self.f16).set(default__.safeIndex((self).f14, len(self.f16)), d_1085_v5_)))))), 3, ((self).f12) + (len(d_1086_v6_)), (self).f14])
                    d_1087_v7_ = (d_1087_v7_ if (self).f25 else d_1087_v7_)
                    (globalState).f6 = (self).f11
                    (globalState).f6 = (self).f11
                    if (self).fm26((self).f25, globalState):
                        (globalState).f2 = False
                        d_1088_v8_: _dafny.Map
                        d_1088_v8_ = _dafny.Map({(self).fm26((self).f15, globalState): (self).f25})
                        d_1088_v8_ = (d_1088_v8_).set((self).f15, (self).f15)
                        d_1089_v9_: _dafny.Set
                        d_1089_v9_ = _dafny.Set({(self).f12})
                        d_1090_v10_: _dafny.Map
                        d_1090_v10_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([True, (self).f25, False]): d_1089_v9_})
                        d_1091_v11_: _dafny.Map
                        d_1091_v11_ = _dafny.Map({len(((d_1090_v10_)[default__.fm30((self).f25, (self).f15, (0) - ((self).f12), (self).f25, globalState)] if (default__.fm30((self).f25, (self).f15, (0) - ((self).f12), (self).f25, globalState)) in (d_1090_v10_) else _dafny.Set({(_dafny.MultiSet([(self).f14, len(d_1081_v1_)])).cardinality}))): (self).f14})
                        d_1091_v11_ = (d_1091_v11_).set(default__.safeDivisionInt(-933, len(d_1086_v6_)), ((self).f11) + ((self).f11))
                        d_1092_v12_: _dafny.Map
                        d_1092_v12_ = _dafny.Map({False: 494})
                        d_1093_v13_: D6
                        d_1093_v13_ = D6_DC19(d_1092_v12_)
                        d_1094_v14_: _dafny.Array
                        nw166_ = _dafny.Array(int(0), 6)
                        d_1094_v14_ = nw166_
                        index192_ = default__.safeIndex(744, (d_1094_v14_).length(0))
                        (d_1094_v14_)[index192_] = (self).f12
                        d_1095_v15_: _dafny.Seq
                        d_1095_v15_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                        d_1096_v16_: _dafny.Map
                        d_1096_v16_ = _dafny.Map({d_1095_v15_: (self).f15})
                        d_1097_v17_: _dafny.MultiSet
                        d_1097_v17_ = _dafny.MultiSet([d_1085_v5_])
                        d_1098_v18_: D0
                        d_1098_v18_ = D0_DC0(d_1097_v17_)
                        index193_ = default__.safeIndex(744, (d_1094_v14_).length(0))
                        rhs178_ = (self).f25
                        rhs179_ = ((self).f15 if ((d_1096_v16_)[d_1095_v15_] if (d_1095_v15_) in (d_1096_v16_) else not((self).f25)) else False)
                        rhs180_ = d_1093_v13_
                        rhs181_ = (((self).f12 if (self).f15 else (self).f11)) + ((self).fm2((self).f11, (self).f15, (self).f11, (self).f15, globalState))
                        rhs182_ = (d_1085_v5_) not in (default__.fm31(d_1085_v5_, d_1098_v18_, globalState))
                        lhs135_ = globalState
                        lhs136_ = globalState
                        lhs137_ = d_1094_v14_
                        lhs138_ = default__.safeIndex(744, (d_1094_v14_).length(0))
                        lhs139_ = globalState
                        lhs135_.f2 = rhs178_
                        lhs136_.f7 = rhs179_
                        d_1093_v13_ = rhs180_
                        lhs137_[lhs138_] = rhs181_
                        lhs139_.f2 = rhs182_
                        (globalState).f7 = ((self).f14) == ((self).f12)
                    elif True:
                        d_1099_v19_: _dafny.Set
                        d_1099_v19_ = _dafny.Set({(self).f14, len((d_1081_v1_) + (d_1081_v1_)), (0) - ((self).f11), len(d_1086_v6_), (self).f14})
                        d_1099_v19_ = d_1099_v19_
                        d_1100_v20_: _dafny.Map
                        d_1100_v20_ = _dafny.Map({d_1081_v1_: len((self).f13)})
                        d_1100_v20_ = d_1100_v20_
                        (globalState).f6 = (0) - ((self).f14)
                        d_1101_v21_: _dafny.Seq
                        d_1101_v21_ = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f25])
                        d_1102_v22_: _dafny.Array
                        def lambda40_(d_1103_i1_):
                            return default__.safeModuloInt(d_1103_i1_, (self).f14)

                        init22_ = lambda40_
                        nw167_ = _dafny.Array(None, 21)
                        for i0_22_ in range(nw167_.length(0)):
                            nw167_[i0_22_] = init22_(i0_22_)
                        d_1102_v22_ = nw167_
                        d_1104_v23_: _dafny.Array
                        d_1104_v23_ = d_1102_v22_
                        d_1105_v24_: _dafny.Map
                        d_1105_v24_ = _dafny.Map({(d_1101_v21_)[default__.safeIndex(len(self.f16), len(d_1101_v21_))]: (d_1104_v23_)})
                        d_1106_v25_: _dafny.Seq
                        d_1106_v25_ = _dafny.SeqWithoutIsStrInference([d_1102_v22_])
                        d_1105_v24_ = (d_1105_v24_).set(((self).f15) == ((self).f15), (d_1106_v25_)[default__.safeIndex((self).f11, len(d_1106_v25_))])
                        (globalState).f2 = (self).f15
                    pass
            pass
        (globalState).f6 = (self).f12
        d_1107_v26_: _dafny.Seq
        d_1107_v26_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f15, True])
        d_1108_v27_: _dafny.MultiSet
        d_1108_v27_ = _dafny.MultiSet([(self).f15])
        (globalState).f7 = (d_1107_v26_)[default__.safeIndex(len(_dafny.Map({d_1108_v27_: (self).f14})), len(d_1107_v26_))]
        d_1109_v28_: _dafny.Array
        nw168_ = _dafny.Array(int(0), 8)
        d_1109_v28_ = nw168_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1109_v28_).length(0)):
            d_1110_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_1110_i2_)) and ((d_1110_i2_) < ((d_1109_v28_).length(0)))):
                (d_1109_v28_)[(d_1110_i2_)] = default__.safeModuloInt(d_1110_i2_, (self).f12)
        d_1111_v29_: str
        d_1111_v29_ = _dafny.CodePoint('w')
        d_1112_v30_: _dafny.Array
        nw169_ = _dafny.Array(None, 24)
        nw169_[int(0)] = d_1111_v29_
        nw169_[int(1)] = d_1111_v29_
        nw169_[int(2)] = d_1111_v29_
        nw169_[int(3)] = d_1111_v29_
        nw169_[int(4)] = d_1111_v29_
        nw169_[int(5)] = _dafny.CodePoint('q')
        nw169_[int(6)] = d_1111_v29_
        nw169_[int(7)] = _dafny.CodePoint('n')
        nw169_[int(8)] = d_1111_v29_
        nw169_[int(9)] = d_1111_v29_
        nw169_[int(10)] = d_1111_v29_
        nw169_[int(11)] = d_1111_v29_
        nw169_[int(12)] = d_1111_v29_
        nw169_[int(13)] = d_1111_v29_
        nw169_[int(14)] = d_1111_v29_
        nw169_[int(15)] = d_1111_v29_
        nw169_[int(16)] = d_1111_v29_
        nw169_[int(17)] = d_1111_v29_
        nw169_[int(18)] = d_1111_v29_
        nw169_[int(19)] = d_1111_v29_
        nw169_[int(20)] = d_1111_v29_
        nw169_[int(21)] = d_1111_v29_
        nw169_[int(22)] = d_1111_v29_
        nw169_[int(23)] = d_1111_v29_
        d_1112_v30_ = nw169_
        index194_ = default__.safeIndex(810, (d_1112_v30_).length(0))
        (d_1112_v30_)[index194_] = d_1111_v29_
        index195_ = default__.safeIndex(810, (d_1112_v30_).length(0))
        (d_1112_v30_)[index195_] = d_1111_v29_
        r0 = default__.safeDivisionInt((0) - ((0) - ((self).f14)), ((self).fm2(559, False, (self).f14, (self).f15, globalState)) + (-209))
        d_1113_v31_: _dafny.Seq
        d_1113_v31_ = _dafny.SeqWithoutIsStrInference([(self).f11])
        r1 = ((d_1113_v31_) + (d_1113_v31_) if (self).f25 else d_1113_v31_)
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        (globalState).f7 = not(not((self).f15))
        if not(((self).f25) and ((self).f15)):
            (globalState).f7 = (self).f15
            d_1114_v0_: _dafny.Map
            d_1114_v0_ = _dafny.Map({(self).f15: _dafny.MultiSet([False, not((self).f15)])})
            d_1115_v1_: _dafny.Seq
            d_1115_v1_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f14, (self).f14, 361, (self).f11])
            d_1116_v2_: _dafny.Map
            d_1116_v2_ = _dafny.Map({d_1115_v1_: p0})
            d_1117_v3_: _dafny.MultiSet
            d_1117_v3_ = _dafny.MultiSet([(self).f25, (self).f15, (self).f25, (self).f15])
            d_1114_v0_ = (d_1114_v0_).set((d_1116_v2_) != (_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f14 for d_1118_i0_ in range(default__.abs(844))]): p0})), d_1117_v3_)
            d_1119_v4_: _dafny.Map
            d_1119_v4_ = _dafny.Map({(self).f25: (self).f25})
            d_1120_v5_: _dafny.Seq
            d_1120_v5_ = _dafny.SeqWithoutIsStrInference([d_1119_v4_])
            d_1121_v6_: _dafny.Array
            nw170_ = _dafny.Array(None, 10)
            d_1121_v6_ = nw170_
            d_1122_v7_: _dafny.Set
            d_1122_v7_ = _dafny.Set({d_1121_v6_, d_1121_v6_, d_1121_v6_, d_1121_v6_, d_1121_v6_})
            d_1123_v8_: str
            d_1123_v8_ = _dafny.CodePoint('m')
            d_1124_v9_: D1
            d_1124_v9_ = D1_DC2(d_1123_v8_)
            d_1125_v10_: D2
            d_1125_v10_ = D2_DC7(p0, d_1122_v7_, d_1124_v9_, (self).f11, (self).f11)
            d_1126_v11_: _dafny.Seq
            d_1126_v11_ = _dafny.SeqWithoutIsStrInference([d_1125_v10_])
            d_1127_v12_: _dafny.Seq
            d_1127_v12_ = _dafny.SeqWithoutIsStrInference([d_1126_v11_])
            d_1128_v13_: _dafny.Seq
            d_1128_v13_ = _dafny.SeqWithoutIsStrInference([(d_1120_v5_)[default__.safeIndex(len(d_1127_v12_), len(d_1120_v5_))], d_1119_v4_])
            d_1120_v5_ = (_dafny.SeqWithoutIsStrInference([d_1119_v4_])) + (d_1120_v5_)
            (globalState).f6 = default__.safeDivisionInt((d_1125_v10_).cf14, (self).f14)
            d_1129_v14_: C0
            nw171_ = C0()
            nw171_.ctor__()
            d_1129_v14_ = nw171_
        elif True:
            d_1130_v15_: _dafny.Seq
            d_1130_v15_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            d_1131_v16_: T2
            nw172_ = C1()
            nw172_.ctor__(p0, (self).f11, (self).f12, len(d_1130_v15_), (self).f15, (self).f26)
            d_1131_v16_ = nw172_
            d_1131_v16_ = d_1131_v16_
            d_1132_v17_: _dafny.Seq
            d_1132_v17_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f25: (self).f14}))])
            d_1133_v18_: _dafny.Map
            d_1133_v18_ = _dafny.Map({True: len(d_1132_v17_)})
            d_1134_v19_: str
            d_1134_v19_ = _dafny.CodePoint('b')
            d_1135_v20_: D2
            d_1135_v20_ = D2_DC6((d_1131_v16_).f13)
            d_1136_v21_: _dafny.Map
            d_1136_v21_ = _dafny.Map({(self).f26: (self).f12})
            d_1137_v22_: D3
            d_1137_v22_ = D3_DC10(not((self).f25), len(_dafny.Map({True: (self).f25})), d_1136_v21_, (d_1131_v16_).f13, ((p1)[(self).f14] if ((self).f14) in (p1) else (d_1131_v16_).fm2((self).f11, (self).f25, (d_1131_v16_).f12, (self).f15, globalState)))
            d_1138_v23_: D2
            d_1138_v23_ = D2_DC8((self).f25, (self).f25, self.f16)
            d_1139_v24_: _dafny.MultiSet
            d_1139_v24_ = _dafny.MultiSet([d_1134_v19_, d_1134_v19_, d_1134_v19_])
            d_1140_v25_: D0
            d_1140_v25_ = D0_DC0(d_1139_v24_)
            d_1141_v26_: _dafny.Array
            nw173_ = _dafny.Array(None, 23)
            nw173_[int(0)] = (((self).f26).set(default__.safeIndex(len(d_1133_v18_), len((self).f26)), d_1134_v19_) if (d_1131_v16_).f15 else (d_1131_v16_).f13)
            nw173_[int(1)] = (d_1131_v16_).f13
            nw173_[int(2)] = (self).f13
            nw173_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpecquwlx"))
            nw173_[int(4)] = (d_1135_v20_).cf9
            nw173_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ql"))
            nw173_[int(6)] = (self).f26
            nw173_[int(7)] = (self).f13
            nw173_[int(8)] = self.f16
            nw173_[int(9)] = (self).f13
            nw173_[int(10)] = (self).f26
            nw173_[int(11)] = ((self).f26) + ((self).f13)
            nw173_[int(12)] = self.f16
            nw173_[int(13)] = (d_1131_v16_).f13
            nw173_[int(14)] = self.f16
            nw173_[int(15)] = (d_1135_v20_).cf9
            nw173_[int(16)] = ((self).f13) + ((d_1131_v16_).f13)
            nw173_[int(17)] = (d_1131_v16_).f13
            nw173_[int(18)] = (self.f16).set(default__.safeIndex((d_1131_v16_).f14, len(self.f16)), d_1134_v19_)
            nw173_[int(19)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xv"))) + (((d_1137_v22_).cf22).set(default__.safeIndex((self).f14, len((d_1137_v22_).cf22)), d_1134_v19_))
            nw173_[int(20)] = (d_1131_v16_).f13
            nw173_[int(21)] = (d_1138_v23_).cf17
            nw173_[int(22)] = (default__.fm31(d_1134_v19_, d_1140_v25_, globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bumgcuck")))
            d_1141_v26_ = nw173_
            index196_ = default__.safeIndex(598, (d_1141_v26_).length(0))
            (d_1141_v26_)[index196_] = (d_1138_v23_).cf17
            index197_ = default__.safeIndex(598, (d_1141_v26_).length(0))
            (d_1141_v26_)[index197_] = ((self).f26) + (((d_1131_v16_).f13) + ((self).f26))
            (globalState).f6 = (d_1131_v16_).f12
            d_1142_v27_: _dafny.Map
            d_1142_v27_ = _dafny.Map({p0: (self).f12})
            d_1143_v28_: D5
            d_1143_v28_ = D5_DC18(d_1142_v27_, (self).f11)
            d_1144_v29_: _dafny.Seq
            d_1144_v29_ = _dafny.SeqWithoutIsStrInference([d_1132_v17_])
            d_1145_v30_: _dafny.Array
            nw174_ = _dafny.Array(None, 14)
            nw174_[int(0)] = d_1143_v28_
            nw174_[int(1)] = d_1143_v28_
            nw174_[int(2)] = D5_DC18(d_1142_v27_, len(d_1144_v29_))
            nw174_[int(3)] = d_1143_v28_
            nw174_[int(4)] = d_1143_v28_
            nw174_[int(5)] = d_1143_v28_
            nw174_[int(6)] = d_1143_v28_
            nw174_[int(7)] = d_1143_v28_
            nw174_[int(8)] = d_1143_v28_
            nw174_[int(9)] = d_1143_v28_
            nw174_[int(10)] = d_1143_v28_
            nw174_[int(11)] = D5_DC18(d_1142_v27_, (self).f11)
            nw174_[int(12)] = d_1143_v28_
            nw174_[int(13)] = d_1143_v28_
            d_1145_v30_ = nw174_
            d_1145_v30_ = d_1145_v30_
            d_1146_v31_: _dafny.MultiSet
            d_1146_v31_ = _dafny.MultiSet([(self).f25])
            d_1147_v32_: _dafny.Map
            d_1147_v32_ = _dafny.Map({d_1134_v19_: p0})
            d_1148_v33_: C1
            nw175_ = C1()
            nw175_.ctor__(((d_1147_v32_)[d_1134_v19_] if (d_1134_v19_) in (d_1147_v32_) else p0), (d_1131_v16_).f12, (d_1131_v16_).f14, (d_1131_v16_).f11, (self).f15, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umu")))
            d_1148_v33_ = nw175_
            d_1149_v34_: _dafny.Map
            d_1149_v34_ = _dafny.Map({d_1148_v33_: (d_1131_v16_).f15})
            d_1150_v35_: _dafny.Seq
            d_1150_v35_ = _dafny.SeqWithoutIsStrInference([d_1148_v33_])
            d_1151_v36_: D9
            d_1151_v36_ = D9_DC26((d_1150_v35_)[default__.safeIndex((d_1131_v16_).f12, len(d_1150_v35_))])
            d_1152_v37_: _dafny.Map
            d_1152_v37_ = _dafny.Map({(self).f25: ((d_1149_v34_)[(d_1151_v36_).cf45] if ((d_1151_v36_).cf45) in (d_1149_v34_) else (self).f15)})
            d_1153_v38_: _dafny.Set
            d_1153_v38_ = _dafny.Set({(d_1131_v16_).f15, ((d_1152_v37_)[(self).f15] if ((self).f15) in (d_1152_v37_) else (d_1131_v16_).f15)})
            d_1154_v39_: _dafny.Map
            d_1154_v39_ = _dafny.Map({d_1134_v19_: (self).f14})
            d_1155_v40_: _dafny.Set
            d_1155_v40_ = _dafny.Set({(self).f11})
            d_1156_v41_: _dafny.Map
            d_1156_v41_ = _dafny.Map({(d_1131_v16_).f15: _dafny.MultiSet([402])})
            d_1157_v43_: _dafny.Map
            d_1157_v43_ = _dafny.Map({(d_1131_v16_).f11: d_1130_v15_})
            d_1158_v44_: D10
            d_1158_v44_ = D10_DC28(d_1157_v43_)
            d_1159_v45_: D4
            d_1159_v45_ = D4_DC13(d_1132_v17_)
            d_1160_v46_: _dafny.MultiSet
            d_1160_v46_ = _dafny.MultiSet([d_1159_v45_])
            d_1161_v48_: _dafny.Map
            def iife66_():
                coll34_ = _dafny.Set()
                compr_34_: int
                for compr_34_ in _dafny.IntegerRange(706, 204):
                    d_1162_v47_: int = compr_34_
                    if ((706) <= (d_1162_v47_)) and ((d_1162_v47_) < (204)):
                        coll34_ = coll34_.union(_dafny.Set([(d_1162_v47_) + ((self).f12)]))
                return _dafny.Set(coll34_)
            d_1161_v48_ = _dafny.Map({(self).f15: iife66_()
            })
            d_1163_v49_: _dafny.Array
            nw176_ = _dafny.Array(None, 27)
            nw176_[int(0)] = (d_1131_v16_).f15
            nw176_[int(1)] = (913) != ((self).f12)
            nw176_[int(2)] = (d_1146_v31_).isdisjoint((D8_DC24(d_1146_v31_)).cf43)
            nw176_[int(3)] = (default__.fm32((self).f25, d_1154_v39_, d_1155_v40_, (self).f25, globalState)).ispropersubset(d_1153_v38_)
            nw176_[int(4)] = (p1).ispropersubset((((d_1156_v41_)[(self).f25] if ((self).f25) in (d_1156_v41_) else _dafny.MultiSet([(self).f11, (d_1131_v16_).f14, (p1).cardinality]))).set(len(d_1132_v17_), default__.abs(-663)))
            nw176_[int(5)] = (self).f15
            nw176_[int(6)] = not((self).f25)
            nw176_[int(7)] = (d_1131_v16_).f15
            def iife67_():
                coll35_ = _dafny.Map()
                compr_35_: int
                for compr_35_ in _dafny.IntegerRange(-128, 126):
                    d_1164_v42_: int = compr_35_
                    if ((-128) <= (d_1164_v42_)) and ((d_1164_v42_) < (126)):
                        coll35_[(d_1164_v42_) * ((self).f14)] = _dafny.SeqWithoutIsStrInference([(d_1131_v16_).f15, (self).f15, (self).f25, (self).f15, (self).f25])
                return _dafny.Map(coll35_)
            nw176_[int(8)] = (iife67_()
            ) != ((d_1158_v44_).cf47)
            nw176_[int(9)] = (d_1131_v16_).f15
            nw176_[int(10)] = (d_1131_v16_).f15
            nw176_[int(11)] = (d_1160_v46_).ispropersubset(d_1160_v46_)
            nw176_[int(12)] = (self).f25
            nw176_[int(13)] = (self).f15
            nw176_[int(14)] = (((d_1161_v48_)[False] if (False) in (d_1161_v48_) else d_1155_v40_)) == (d_1155_v40_)
            nw176_[int(15)] = not(True)
            nw176_[int(16)] = (d_1131_v16_).f15
            nw176_[int(17)] = (self).fm1(len(self.f16), (self).f15, (D1_DC4((d_1131_v16_).f15)).cf5, globalState)
            nw176_[int(18)] = (self).f15
            nw176_[int(19)] = (self).f25
            nw176_[int(20)] = (self).f15
            nw176_[int(21)] = not(((0) - ((self).f12)) >= ((self).f14))
            nw176_[int(22)] = (self).f15
            nw176_[int(23)] = True
            nw176_[int(24)] = False
            nw176_[int(25)] = (d_1130_v15_)[default__.safeIndex((self).f11, len(d_1130_v15_))]
            nw176_[int(26)] = (self).f25
            d_1163_v49_ = nw176_
            d_1163_v49_ = d_1148_v33_.f24
        d_1165_v50_: str
        d_1165_v50_ = _dafny.CodePoint('s')
        d_1166_v51_: D10
        d_1166_v51_ = D10_DC31((0) - ((self).f11), d_1165_v50_, (self).f12, (self).f11)
        d_1167_v52_: D10
        d_1167_v52_ = D10_DC32(d_1166_v51_)
        source19_ = d_1167_v52_
        if source19_.is_DC29:
            d_1168_v53_: C0
            nw177_ = C0()
            nw177_.ctor__()
            d_1168_v53_ = nw177_
            d_1169_v54_: _dafny.MultiSet
            d_1169_v54_ = _dafny.MultiSet([(self).f15, (self).f15])
            d_1170_v55_: _dafny.Set
            d_1170_v55_ = _dafny.Set({(self).f15})
            d_1171_v56_: C1
            nw178_ = C1()
            nw178_.ctor__(p0, (d_1169_v54_).cardinality, (self).f12, default__.safeDivisionInt((self).f14, len(d_1170_v55_)), (self).f15, (self).f13)
            d_1171_v56_ = nw178_
            source20_ = D10_DC32(d_1166_v51_)
            if source20_.is_DC29:
                d_1172_v57_: T2
                nw179_ = C1()
                nw179_.ctor__(d_1171_v56_.f24, ((self).f12) * ((self).f11), -499, (self).f12, (self).f15, _dafny.SeqWithoutIsStrInference([d_1165_v50_ for d_1173_i1_ in range(default__.abs(685))]))
                d_1172_v57_ = nw179_
                d_1172_v57_ = d_1172_v57_
                d_1174_v58_: _dafny.Map
                d_1174_v58_ = _dafny.Map({(d_1172_v57_).f15: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccfbjwm")))})
                d_1175_v59_: C1
                nw180_ = C1()
                nw180_.ctor__(d_1171_v56_.f24, ((p1)[len(d_1174_v58_)] if (len(d_1174_v58_)) in (p1) else (d_1172_v57_).f14), (d_1172_v57_).f12, (d_1172_v57_).f11, (self).f15, ((self).f13) + ((self).f26))
                d_1175_v59_ = nw180_
                d_1176_v60_: _dafny.Map
                d_1176_v60_ = _dafny.Map({d_1175_v59_: d_1171_v56_.f24})
                (d_1175_v59_).f24 = ((d_1176_v60_)[d_1171_v56_] if (d_1171_v56_) in (d_1176_v60_) else d_1175_v59_.f24)
                (globalState).f2 = not((((self).f12) > ((self).f14) if (self).f15 else (self).f25))
            elif source20_.is_DC30:
                d_1177___mcc_h10_ = source20_.cf48
                d_1178___mcc_h11_ = source20_.cf49
                d_1179___mcc_h12_ = source20_.cf50
                d_1180___mcc_h13_ = source20_.cf51
                d_1181_cf51_ = d_1180___mcc_h13_
                d_1182_cf50_ = d_1179___mcc_h12_
                d_1183_cf49_ = d_1178___mcc_h11_
                d_1184_cf48_ = d_1177___mcc_h10_
                d_1185_v61_: _dafny.Map
                d_1185_v61_ = _dafny.Map({d_1184_cf48_: d_1171_v56_.f24})
                d_1185_v61_ = ((d_1185_v61_) | (d_1185_v61_)) | (d_1185_v61_)
                d_1186_v62_: D1
                d_1186_v62_ = D1_DC5(len(_dafny.SeqWithoutIsStrInference([d_1165_v50_ for d_1187_i2_ in range(default__.abs(617))])), d_1183_cf49_, (self).f25)
                d_1188_v63_: _dafny.Map
                d_1188_v63_ = _dafny.Map({d_1170_v55_: (self).f11})
                d_1189_v64_: _dafny.Map
                d_1189_v64_ = _dafny.Map({D1_DC5(len(d_1188_v63_), (self).f25, d_1181_cf51_): (self).f11})
                d_1190_v65_: _dafny.Map
                d_1190_v65_ = _dafny.Map({(d_1186_v62_) in (d_1189_v64_): (self).f13})
                d_1190_v65_ = (d_1190_v65_).set(False, self.f16)
                (globalState).f2 = True
                rhs183_ = 309
                rhs184_ = len(_dafny.SeqWithoutIsStrInference([D8_DC25((self).f14) for d_1191_i3_ in range(default__.abs(703))]))
                lhs140_ = globalState
                lhs141_ = globalState
                lhs140_.f6 = rhs183_
                lhs141_.f6 = rhs184_
            elif source20_.is_DC31:
                d_1192___mcc_h14_ = source20_.cf52
                d_1193___mcc_h15_ = source20_.cf53
                d_1194___mcc_h16_ = source20_.cf54
                d_1195___mcc_h17_ = source20_.cf55
                d_1196_cf55_ = d_1195___mcc_h17_
                d_1197_cf54_ = d_1194___mcc_h16_
                d_1198_cf53_ = d_1193___mcc_h15_
                d_1199_cf52_ = d_1192___mcc_h14_
                d_1200_v66_: _dafny.Map
                d_1200_v66_ = _dafny.Map({(self).f15: d_1171_v56_.f24})
                (globalState).f7 = (d_1200_v66_) == ((d_1200_v66_).set((self).f15, p0))
                (globalState).f6 = (self).f11
                d_1165_v50_ = d_1165_v50_
                r0 = not (((self).f11) > (-86)) or ((self).f15)
            elif source20_.is_DC28:
                d_1201___mcc_h18_ = source20_.cf47
                d_1202_cf47_ = d_1201___mcc_h18_
                d_1203_v67_: _dafny.Seq
                d_1203_v67_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                (globalState).f2 = (d_1203_v67_) < ((d_1203_v67_) + (_dafny.SeqWithoutIsStrInference([(self).f11])))
                d_1204_v68_: C1
                nw181_ = C1()
                nw181_.ctor__(d_1171_v56_.f24, ((self).f14) + ((self).f11), len(self.f16), len((self).f13), (self).f15, self.f16)
                d_1204_v68_ = nw181_
                (globalState).f6 = (0) - ((self).f12)
                d_1205_v69_: _dafny.Array
                def lambda41_(d_1206_i4_):
                    return (d_1206_i4_) * (len(_dafny.SeqWithoutIsStrInference([False])))

                init23_ = lambda41_
                nw182_ = _dafny.Array(None, 5)
                for i0_23_ in range(nw182_.length(0)):
                    nw182_[i0_23_] = init23_(i0_23_)
                d_1205_v69_ = nw182_
                index198_ = default__.safeIndex(801, (d_1205_v69_).length(0))
                (d_1205_v69_)[index198_] = ((p1 if (self).f15 else default__.fm33(globalState))).cardinality
                d_1207_v70_: _dafny.Seq
                d_1207_v70_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                index199_ = default__.safeIndex(801, (d_1205_v69_).length(0))
                rhs185_ = (d_1203_v67_)[default__.safeIndex(default__.safeDivisionInt((self).f14, len(d_1207_v70_)), len(d_1203_v67_))]
                rhs186_ = default__.safeDivisionInt((self).f14, (self).f14)
                lhs142_ = globalState
                lhs143_ = d_1205_v69_
                lhs144_ = default__.safeIndex(801, (d_1205_v69_).length(0))
                lhs142_.f6 = rhs185_
                lhs143_[lhs144_] = rhs186_
            elif True:
                d_1208___mcc_h19_ = source20_.cf56
                d_1209_cf56_ = d_1208___mcc_h19_
                (globalState).f2 = not(not((self).f15))
                arr24_ = d_1171_v56_.f24
                index200_ = default__.safeIndex(282, (d_1171_v56_.f24).length(0))
                arr24_[index200_] = True
                arr25_ = d_1171_v56_.f24
                index201_ = default__.safeIndex(282, (d_1171_v56_.f24).length(0))
                arr25_[index201_] = (((self).f12) - (len(self.f16))) != ((self).f12)
                (globalState).f7 = (d_1171_v56_.f24)[default__.safeIndex(282, (d_1171_v56_.f24).length(0))]
                (globalState).f7 = False
            d_1210_v71_: _dafny.Array
            nw183_ = _dafny.Array(_dafny.Seq({}), 12)
            d_1210_v71_ = nw183_
            d_1211_v72_: _dafny.Seq
            d_1211_v72_ = _dafny.SeqWithoutIsStrInference([False])
            index202_ = default__.safeIndex(18, (d_1210_v71_).length(0))
            (d_1210_v71_)[index202_] = (d_1211_v72_) + (d_1211_v72_)
            index203_ = default__.safeIndex(18, (d_1210_v71_).length(0))
            (d_1210_v71_)[index203_] = (d_1211_v72_) + ((d_1211_v72_) + (d_1211_v72_))
        elif source19_.is_DC30:
            d_1212___mcc_h0_ = source19_.cf48
            d_1213___mcc_h1_ = source19_.cf49
            d_1214___mcc_h2_ = source19_.cf50
            d_1215___mcc_h3_ = source19_.cf51
            d_1216_cf51_ = d_1215___mcc_h3_
            d_1217_cf50_ = d_1214___mcc_h2_
            d_1218_cf49_ = d_1213___mcc_h1_
            d_1219_cf48_ = d_1212___mcc_h0_
            if (self).fm26(((self).f12) < ((self).f11), globalState):
                d_1220_v73_: bool
                d_1221_v74_: _dafny.Seq
                out46_: bool
                out47_: _dafny.Seq
                out46_, out47_ = default__.m0(d_1219_cf48_, (self).f15, (default__.fm0(not((self).f15), globalState)) * ((self).f11), (self).f12, globalState)
                d_1220_v73_ = out46_
                d_1221_v74_ = out47_
                d_1222_v75_: D10
                d_1222_v75_ = D10_DC30(d_1219_cf48_, False, (self).f13, True)
                d_1223_v76_: _dafny.Map
                d_1223_v76_ = _dafny.Map({((d_1222_v75_).cf50) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kejwlatg"))): len(d_1217_cf50_)})
                d_1223_v76_ = (d_1223_v76_).set((self.f16) + (d_1217_cf50_), len(_dafny.Map({default__.fm33(globalState): self.f16})))
                d_1218_cf49_ = (self).f15
                d_1224_v77_: D1
                d_1224_v77_ = D1_DC5((self).f11, True, d_1220_v73_)
                d_1165_v50_ = (d_1165_v50_ if (d_1224_v77_).cf7 else default__.fm34(d_1165_v50_, globalState))
                d_1225_v78_: C0
                nw184_ = C0()
                nw184_.ctor__()
                d_1225_v78_ = nw184_
            elif True:
                d_1226_v79_: _dafny.Array
                nw185_ = _dafny.Array(None, 1)
                nw185_[int(0)] = (self).f12
                d_1226_v79_ = nw185_
                d_1227_v80_: _dafny.Seq
                d_1227_v80_ = _dafny.SeqWithoutIsStrInference([d_1226_v79_, d_1226_v79_])
                d_1228_v81_: D3
                d_1228_v81_ = D3_DC12(D3_DC9(d_1227_v80_))
                d_1228_v81_ = d_1228_v81_
                r0 = (self).f25
                d_1229_v82_: _dafny.Map
                d_1229_v82_ = _dafny.Map({len((self).f26): (self).f11})
                d_1230_v83_: C1
                nw186_ = C1()
                nw186_.ctor__(p0, (-884) - ((self).fm2((0) - ((self).f14), d_1216_cf51_, (self).f14, False, globalState)), ((self).f11 if (self).f25 else (self).f14), (self).f11, d_1216_cf51_, (_dafny.SeqWithoutIsStrInference([d_1165_v50_ for d_1231_i5_ in range(default__.abs(236))])).set(default__.safeIndex(len(d_1229_v82_), len(_dafny.SeqWithoutIsStrInference([d_1165_v50_ for d_1232_i5_ in range(default__.abs(236))]))), _dafny.CodePoint('v')))
                d_1230_v83_ = nw186_
                arr26_ = d_1230_v83_.f24
                index204_ = default__.safeIndex(30, (d_1230_v83_.f24).length(0))
                arr26_[index204_] = d_1218_cf49_
                d_1233_v84_: _dafny.Seq
                d_1233_v84_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                arr27_ = d_1230_v83_.f24
                index205_ = default__.safeIndex(30, (d_1230_v83_.f24).length(0))
                arr27_[index205_] = ((_dafny.MultiSet(d_1233_v84_)).cardinality) in (_dafny.Map({(self).f11: False}))
                d_1234_v85_: _dafny.Map
                d_1234_v85_ = _dafny.Map({default__.fm35((d_1230_v83_.f24)[default__.safeIndex(30, (d_1230_v83_.f24).length(0))], (d_1233_v84_).set(default__.safeIndex((self).f12, len(d_1233_v84_)), len((self).f13)), globalState): (self).f14})
                d_1235_v86_: _dafny.Set
                d_1235_v86_ = _dafny.Set({((self).f11) * ((0) - ((0) - (len(d_1234_v85_))))})
                d_1236_v87_: _dafny.MultiSet
                d_1236_v87_ = _dafny.MultiSet([(self).f25])
                d_1237_v88_: _dafny.Seq
                d_1237_v88_ = _dafny.SeqWithoutIsStrInference([d_1216_cf51_])
                d_1238_v89_: _dafny.Seq
                d_1238_v89_ = _dafny.SeqWithoutIsStrInference([d_1237_v88_, d_1237_v88_])
                d_1239_v90_: _dafny.Map
                d_1239_v90_ = _dafny.Map({d_1218_cf49_: d_1165_v50_})
                d_1240_v92_: _dafny.Seq
                def iife68_():
                    coll36_ = _dafny.Set()
                    compr_36_: int
                    for compr_36_ in (_dafny.MultiSet([((d_1236_v87_)[(d_1230_v83_.f24)[default__.safeIndex(30, (d_1230_v83_.f24).length(0))]] if ((d_1230_v83_.f24)[default__.safeIndex(30, (d_1230_v83_.f24).length(0))]) in (d_1236_v87_) else (self).f11), (self).f12, (_dafny.MultiSet((d_1238_v89_)[default__.safeIndex((self).f14, len(d_1238_v89_))])).cardinality, len(d_1239_v90_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cun")))])).Elements:
                        d_1241_v91_: int = compr_36_
                        if (d_1241_v91_) in (_dafny.MultiSet([((d_1236_v87_)[(d_1230_v83_.f24)[default__.safeIndex(30, (d_1230_v83_.f24).length(0))]] if ((d_1230_v83_.f24)[default__.safeIndex(30, (d_1230_v83_.f24).length(0))]) in (d_1236_v87_) else (self).f11), (self).f12, (_dafny.MultiSet((d_1238_v89_)[default__.safeIndex((self).f14, len(d_1238_v89_))])).cardinality, len(d_1239_v90_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cun")))])):
                            coll36_ = coll36_.union(_dafny.Set([default__.safeDivisionInt(d_1241_v91_, -289)]))
                    return _dafny.Set(coll36_)
                d_1240_v92_ = _dafny.SeqWithoutIsStrInference([iife68_()
                ])
                d_1235_v86_ = (d_1235_v86_) - ((d_1240_v92_)[default__.safeIndex((self).f14, len(d_1240_v92_))])
            d_1242_v93_: _dafny.Array
            nw187_ = _dafny.Array(None, 1)
            nw187_[int(0)] = (self).f14
            d_1242_v93_ = nw187_
            (self).m17(p0, d_1242_v93_, d_1242_v93_, globalState)
            d_1243_v94_: _dafny.Seq
            d_1243_v94_ = _dafny.SeqWithoutIsStrInference([(self).f11])
            d_1243_v94_ = _dafny.SeqWithoutIsStrInference([(self).f14, len((d_1243_v94_) + (d_1243_v94_))])
            index206_ = default__.safeIndex(977, (d_1242_v93_).length(0))
            (d_1242_v93_)[index206_] = (self).f12
            d_1244_v95_: _dafny.Array
            nw188_ = _dafny.Array(False, 28)
            d_1244_v95_ = nw188_
            index207_ = default__.safeIndex(977, (d_1242_v93_).length(0))
            rhs187_ = (((self).f14) + ((self).f12)) >= ((self).f11)
            rhs188_ = (0) - ((self).f12)
            rhs189_ = True
            rhs190_ = (d_1244_v95_ if False else d_1244_v95_)
            lhs145_ = d_1242_v93_
            lhs146_ = default__.safeIndex(977, (d_1242_v93_).length(0))
            d_1218_cf49_ = rhs187_
            lhs145_[lhs146_] = rhs188_
            d_1216_cf51_ = rhs189_
            d_1244_v95_ = rhs190_
        elif source19_.is_DC31:
            d_1245___mcc_h4_ = source19_.cf52
            d_1246___mcc_h5_ = source19_.cf53
            d_1247___mcc_h6_ = source19_.cf54
            d_1248___mcc_h7_ = source19_.cf55
            d_1249_cf55_ = d_1248___mcc_h7_
            d_1250_cf54_ = d_1247___mcc_h6_
            d_1251_cf53_ = d_1246___mcc_h5_
            d_1252_cf52_ = d_1245___mcc_h4_
            d_1253_v96_: D10
            d_1253_v96_ = D10_DC30((self).f25, not((self).f15), (self).f13, (self).f25)
            source21_ = d_1253_v96_
            if source21_.is_DC29:
                d_1254_v97_: _dafny.Map
                d_1254_v97_ = _dafny.Map({self.f16: (self).f12})
                d_1255_v98_: D3
                d_1255_v98_ = D3_DC10((self).f25, 850, d_1254_v97_, _dafny.SeqWithoutIsStrInference([d_1251_cf53_ for d_1256_i6_ in range(default__.abs(743))]), 190)
                (self).f16 = ((self).f13) + ((d_1255_v98_).cf22)
                (globalState).f6 = d_1252_cf52_
                d_1249_cf55_ = 780
                d_1257_v99_: _dafny.Array
                def lambda42_(d_1258_i7_):
                    return (self.f16)[default__.safeIndex(7, len(self.f16))]

                init24_ = lambda42_
                nw189_ = _dafny.Array(None, 6)
                for i0_24_ in range(nw189_.length(0)):
                    nw189_[i0_24_] = init24_(i0_24_)
                d_1257_v99_ = nw189_
                index208_ = default__.safeIndex(685, (d_1257_v99_).length(0))
                (d_1257_v99_)[index208_] = default__.fm34(d_1251_cf53_, globalState)
                index209_ = default__.safeIndex(685, (d_1257_v99_).length(0))
                (d_1257_v99_)[index209_] = _dafny.CodePoint('x')
            elif source21_.is_DC30:
                d_1259___mcc_h20_ = source21_.cf48
                d_1260___mcc_h21_ = source21_.cf49
                d_1261___mcc_h22_ = source21_.cf50
                d_1262___mcc_h23_ = source21_.cf51
                d_1263_cf51_ = d_1262___mcc_h23_
                d_1264_cf50_ = d_1261___mcc_h22_
                d_1265_cf49_ = d_1260___mcc_h21_
                d_1266_cf48_ = d_1259___mcc_h20_
                d_1267_v100_: _dafny.Seq
                d_1267_v100_ = _dafny.SeqWithoutIsStrInference([(0) - (d_1252_cf52_)])
                d_1268_v101_: _dafny.Set
                d_1268_v101_ = _dafny.Set({(self).fm2(d_1250_cf54_, (self).f25, d_1252_cf52_, (self).f15, globalState), len((self).f13), 981})
                d_1269_v103_: _dafny.Seq
                d_1269_v103_ = _dafny.SeqWithoutIsStrInference([not((self).fm1(34, d_1266_cf48_, False, globalState))])
                d_1270_v105_: _dafny.Array
                nw190_ = _dafny.Array(None, 21)
                nw190_[int(0)] = d_1268_v101_
                nw190_[int(1)] = d_1268_v101_
                nw190_[int(2)] = d_1268_v101_
                nw190_[int(3)] = d_1268_v101_
                nw190_[int(4)] = d_1268_v101_
                nw190_[int(5)] = _dafny.Set({(self).f12})
                nw190_[int(6)] = d_1268_v101_
                nw190_[int(7)] = d_1268_v101_
                nw190_[int(8)] = _dafny.Set({(self).f14})
                nw190_[int(9)] = _dafny.Set({104, (0) - ((self).f14), (self).f11, d_1249_cf55_})
                nw190_[int(10)] = d_1268_v101_
                def iife69_():
                    coll37_ = _dafny.Set()
                    compr_37_: int
                    for compr_37_ in _dafny.IntegerRange(-190, 933):
                        d_1271_v102_: int = compr_37_
                        if ((-190) <= (d_1271_v102_)) and ((d_1271_v102_) < (933)):
                            coll37_ = coll37_.union(_dafny.Set([default__.safeDivisionInt(d_1271_v102_, d_1252_cf52_)]))
                    return _dafny.Set(coll37_)
                nw190_[int(11)] = iife69_()
                
                nw190_[int(12)] = _dafny.Set({len(d_1269_v103_), d_1252_cf52_, d_1252_cf52_})
                nw190_[int(13)] = d_1268_v101_
                def iife70_():
                    coll38_ = _dafny.Set()
                    compr_38_: int
                    for compr_38_ in _dafny.IntegerRange(533, -695):
                        d_1272_v104_: int = compr_38_
                        if ((533) <= (d_1272_v104_)) and ((d_1272_v104_) < (-695)):
                            coll38_ = coll38_.union(_dafny.Set([(d_1272_v104_) + (len(d_1264_cf50_))]))
                    return _dafny.Set(coll38_)
                nw190_[int(14)] = iife70_()
                
                nw190_[int(15)] = (d_1268_v101_).intersection(_dafny.Set({d_1249_cf55_, d_1250_cf54_}))
                nw190_[int(16)] = d_1268_v101_
                nw190_[int(17)] = (d_1268_v101_).intersection(d_1268_v101_)
                nw190_[int(18)] = d_1268_v101_
                nw190_[int(19)] = d_1268_v101_
                nw190_[int(20)] = d_1268_v101_
                d_1270_v105_ = nw190_
                rhs191_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_1266_cf48_, globalState), d_1252_cf52_])
                rhs192_ = (((self).f12) * ((self).f14)) == (len(d_1268_v101_))
                rhs193_ = d_1270_v105_
                lhs147_ = globalState
                d_1267_v100_ = rhs191_
                lhs147_.f7 = rhs192_
                d_1270_v105_ = rhs193_
                (self).f16 = _dafny.SeqWithoutIsStrInference([d_1251_cf53_ for d_1273_i8_ in range(default__.abs(-88))])
                d_1274_v106_: C0
                nw191_ = C0()
                nw191_.ctor__()
                d_1274_v106_ = nw191_
                d_1267_v100_ = (d_1267_v100_) + (d_1267_v100_)
            elif source21_.is_DC31:
                d_1275___mcc_h24_ = source21_.cf52
                d_1276___mcc_h25_ = source21_.cf53
                d_1277___mcc_h26_ = source21_.cf54
                d_1278___mcc_h27_ = source21_.cf55
                d_1279_cf55_ = d_1278___mcc_h27_
                d_1280_cf54_ = d_1277___mcc_h26_
                d_1281_cf53_ = d_1276___mcc_h25_
                d_1282_cf52_ = d_1275___mcc_h24_
                d_1283_v107_: _dafny.Map
                d_1283_v107_ = _dafny.Map({(self).f25: d_1279_cf55_})
                d_1284_v108_: _dafny.Seq
                d_1284_v108_ = _dafny.SeqWithoutIsStrInference([((d_1283_v107_)[(self).f25] if ((self).f25) in (d_1283_v107_) else (self).f11), 231])
                d_1285_v109_: _dafny.MultiSet
                d_1285_v109_ = _dafny.MultiSet([p1, _dafny.MultiSet([(self).f14]), p1, (p1).intersection((p1).set((0) - ((self).f12), default__.abs((self).f12))), (_dafny.MultiSet(d_1284_v108_)) - ((_dafny.MultiSet([default__.fm0((self).f25, globalState), 376, d_1279_cf55_])).set(len((self).f13), default__.abs((self).f14)))])
                d_1285_v109_ = d_1285_v109_
                d_1286_v110_: _dafny.Seq
                d_1286_v110_ = _dafny.SeqWithoutIsStrInference([(self.f16) + (self.f16)])
                d_1286_v110_ = d_1286_v110_
                d_1287_v111_: C1
                nw192_ = C1()
                nw192_.ctor__(p0, d_1280_cf54_, (self).f14, (self).f11, (self).f15, (self.f16) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pma"))))
                d_1287_v111_ = nw192_
                arr28_ = d_1287_v111_.f24
                index210_ = default__.safeIndex(666, (d_1287_v111_.f24).length(0))
                arr28_[index210_] = (d_1287_v111_).fm1((self).f12, (d_1287_v111_).fm1((self).f14, (self).f15, (self).f25, globalState), False, globalState)
                d_1288_v113_: _dafny.Map
                def iife71_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in _dafny.IntegerRange(665, 237):
                        d_1289_v112_: int = compr_39_
                        if ((665) <= (d_1289_v112_)) and ((d_1289_v112_) < (237)):
                            coll39_[(d_1289_v112_) - (d_1279_cf55_)] = len(_dafny.Set({d_1282_cf52_, -394}))
                    return _dafny.Map(coll39_)
                d_1288_v113_ = _dafny.Map({iife71_()
                : (self).fm1(d_1279_cf55_, (self).f25, True, globalState)})
                arr29_ = d_1287_v111_.f24
                index211_ = default__.safeIndex(666, (d_1287_v111_.f24).length(0))
                rhs194_ = (self).f15
                rhs195_ = d_1253_v96_
                rhs196_ = (0) - (default__.safeDivisionInt((len(d_1288_v113_)) * ((self).f14), (self).fm2((self).f14, (self).f15, (self).f11, (self).f15, globalState)))
                lhs148_ = d_1287_v111_.f24
                lhs149_ = default__.safeIndex(666, (d_1287_v111_.f24).length(0))
                lhs150_ = globalState
                lhs148_[lhs149_] = rhs194_
                d_1253_v96_ = rhs195_
                lhs150_.f6 = rhs196_
            elif source21_.is_DC28:
                d_1290___mcc_h28_ = source21_.cf47
                d_1291_cf47_ = d_1290___mcc_h28_
                d_1292_v114_: C1
                nw193_ = C1()
                nw193_.ctor__(p0, len((self).f26), d_1252_cf52_, d_1250_cf54_, (self).f15, (self).f26)
                d_1292_v114_ = nw193_
                arr30_ = d_1292_v114_.f24
                index212_ = default__.safeIndex(915, (d_1292_v114_.f24).length(0))
                arr30_[index212_] = not((self).f15)
                d_1293_v115_: _dafny.MultiSet
                d_1293_v115_ = _dafny.MultiSet([d_1165_v50_])
                d_1294_v116_: D0
                d_1294_v116_ = D0_DC0(d_1293_v115_)
                d_1295_v117_: _dafny.Seq
                d_1295_v117_ = _dafny.SeqWithoutIsStrInference([d_1294_v116_])
                arr31_ = d_1292_v114_.f24
                index213_ = default__.safeIndex(915, (d_1292_v114_.f24).length(0))
                arr31_[index213_] = (_dafny.SeqWithoutIsStrInference([d_1294_v116_, D0_DC0(d_1293_v115_), d_1294_v116_])) < (d_1295_v117_)
                (d_1292_v114_).m14(globalState)
                (globalState).f6 = (0) - ((d_1249_cf55_) + ((self).fm2(d_1250_cf54_, (self).f25, d_1250_cf54_, (d_1292_v114_.f24)[default__.safeIndex(915, (d_1292_v114_.f24).length(0))], globalState)))
            elif True:
                d_1296___mcc_h29_ = source21_.cf56
                d_1297_cf56_ = d_1296___mcc_h29_
                d_1252_cf52_ = d_1252_cf52_
                d_1298_v118_: _dafny.MultiSet
                d_1298_v118_ = _dafny.MultiSet([d_1251_cf53_, _dafny.CodePoint('d'), d_1165_v50_, d_1251_cf53_, d_1165_v50_])
                d_1299_v119_: D0
                d_1299_v119_ = D0_DC0(d_1298_v118_)
                d_1300_v120_: C1
                nw194_ = C1()
                nw194_.ctor__(p0, (self).f12, (0) - (d_1249_cf55_), d_1250_cf54_, (self).f25, (default__.fm31(d_1251_cf53_, d_1299_v119_, globalState)) + (_dafny.SeqWithoutIsStrInference([d_1251_cf53_ for d_1301_i9_ in range(default__.abs(837))])))
                d_1300_v120_ = nw194_
                (globalState).f7 = (-587) <= ((0) - ((self).fm2((self).f12, (self).f15, (self).f14, (self).f15, globalState)))
                d_1302_v121_: _dafny.Seq
                d_1302_v121_ = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f25])
                d_1303_v122_: _dafny.MultiSet
                d_1303_v122_ = _dafny.MultiSet([(self).f15, (self).f25])
                d_1304_v123_: _dafny.Map
                d_1304_v123_ = _dafny.Map({(self).f15: (d_1303_v122_).cardinality})
                d_1305_v124_: _dafny.Seq
                d_1305_v124_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                d_1306_v125_: _dafny.Seq
                d_1306_v125_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len((d_1302_v121_).set(default__.safeIndex((self).f12, len(d_1302_v121_)), (self).f25)), len(d_1304_v123_)]), d_1305_v124_])
                rhs197_ = not(not((self).f15))
                rhs198_ = (_dafny.MultiSet(d_1306_v125_)).isdisjoint(_dafny.MultiSet(d_1306_v125_))
                rhs199_ = (d_1305_v124_) not in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1252_cf52_, d_1252_cf52_, (self).f14])]))
                rhs200_ = (self).f12
                lhs151_ = globalState
                lhs152_ = globalState
                lhs153_ = globalState
                lhs151_.f2 = rhs197_
                lhs152_.f7 = rhs198_
                lhs153_.f2 = rhs199_
                d_1252_cf52_ = rhs200_
            if (self).f25:
                d_1252_cf52_ = (0) - ((self).f14)
                d_1307_v126_: _dafny.Array
                def lambda43_(d_1308_i10_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gperypwu"))

                init25_ = lambda43_
                nw195_ = _dafny.Array(None, 14)
                for i0_25_ in range(nw195_.length(0)):
                    nw195_[i0_25_] = init25_(i0_25_)
                d_1307_v126_ = nw195_
                index214_ = default__.safeIndex(820, (d_1307_v126_).length(0))
                (d_1307_v126_)[index214_] = ((self).f13) + (self.f16)
                index215_ = default__.safeIndex(820, (d_1307_v126_).length(0))
                (d_1307_v126_)[index215_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wyovsxkv"))
                (globalState).f6 = default__.fm0((self).fm26((self).f25, globalState), globalState)
                d_1309_v127_: _dafny.Array
                def lambda44_(d_1310_i11_):
                    return (self).f25

                init26_ = lambda44_
                nw196_ = _dafny.Array(None, 8)
                for i0_26_ in range(nw196_.length(0)):
                    nw196_[i0_26_] = init26_(i0_26_)
                d_1309_v127_ = nw196_
                d_1311_v128_: D6
                d_1311_v128_ = D6_DC21((self).f25)
                index216_ = default__.safeIndex(224, (p0).length(0))
                (p0)[index216_] = (d_1165_v50_) not in ((self).f13)
                index217_ = default__.safeIndex(224, (p0).length(0))
                rhs201_ = d_1249_cf55_
                rhs202_ = p0
                rhs203_ = d_1311_v128_
                rhs204_ = (self).f25
                lhs154_ = p0
                lhs155_ = default__.safeIndex(224, (p0).length(0))
                d_1250_cf54_ = rhs201_
                d_1309_v127_ = rhs202_
                d_1311_v128_ = rhs203_
                lhs154_[lhs155_] = rhs204_
                d_1309_v127_ = d_1309_v127_
            elif True:
                d_1312_v129_: _dafny.Seq
                d_1312_v129_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1313_v130_: _dafny.Seq
                d_1313_v130_ = _dafny.SeqWithoutIsStrInference([(self).f14, len(d_1312_v129_)])
                d_1314_v131_: _dafny.Map
                d_1314_v131_ = _dafny.Map({(self).f15: d_1250_cf54_})
                d_1315_v132_: D6
                d_1315_v132_ = D6_DC19(d_1314_v131_)
                d_1316_v133_: _dafny.Map
                d_1316_v133_ = _dafny.Map({(self).f25: d_1315_v132_})
                d_1317_v134_: D6
                d_1317_v134_ = D6_DC22(((d_1316_v133_)[(self).f15] if ((self).f15) in (d_1316_v133_) else d_1315_v132_))
                d_1318_v135_: _dafny.Array
                nw197_ = _dafny.Array(None, 9)
                nw197_[int(0)] = default__.fm36(len(d_1313_v130_), (self).fm26((self).f25, globalState), (self).f26, d_1252_cf52_, globalState)
                nw197_[int(1)] = d_1317_v134_
                nw197_[int(2)] = d_1317_v134_
                nw197_[int(3)] = d_1317_v134_
                nw197_[int(4)] = d_1317_v134_
                nw197_[int(5)] = d_1317_v134_
                nw197_[int(6)] = d_1317_v134_
                nw197_[int(7)] = D6_DC22(d_1315_v132_)
                nw197_[int(8)] = d_1317_v134_
                d_1318_v135_ = nw197_
                rhs205_ = d_1318_v135_
                rhs206_ = _dafny.SeqWithoutIsStrInference([d_1165_v50_ for d_1319_i12_ in range(default__.abs(330))])
                rhs207_ = d_1251_cf53_
                lhs156_ = self
                d_1318_v135_ = rhs205_
                lhs156_.f16 = rhs206_
                d_1251_cf53_ = rhs207_
                d_1320_v136_: _dafny.Array
                nw198_ = _dafny.Array(D10.default()(), 3)
                d_1320_v136_ = nw198_
                d_1321_v137_: _dafny.Map
                d_1321_v137_ = _dafny.Map({default__.safeModuloInt((0) - ((self).f12), -626): (d_1320_v136_ if False else d_1320_v136_)})
                d_1321_v137_ = (d_1321_v137_).set(((self).f14) * (451), d_1320_v136_)
                index218_ = default__.safeIndex(824, (p0).length(0))
                (p0)[index218_] = ((self).f15 if (self).f15 else True)
                d_1322_v138_: _dafny.Set
                d_1322_v138_ = _dafny.Set({(self).f11, d_1249_cf55_, -763})
                index219_ = default__.safeIndex(824, (p0).length(0))
                (p0)[index219_] = ((self).f25 if not(not(((self).f25) or ((self).f25))) else (d_1322_v138_).isdisjoint(d_1322_v138_))
                d_1323_v139_: _dafny.Seq
                d_1323_v139_ = _dafny.SeqWithoutIsStrInference([default__.fm37((p0)[default__.safeIndex(824, (p0).length(0))], globalState)])
                d_1324_v140_: _dafny.Map
                d_1324_v140_ = _dafny.Map({(p0)[default__.safeIndex(824, (p0).length(0))]: d_1323_v139_})
                d_1324_v140_ = ((d_1324_v140_ if (self).f25 else d_1324_v140_)) | (_dafny.Map({(self).fm1(len(self.f16), (p0)[default__.safeIndex(824, (p0).length(0))], (self).f15, globalState): d_1323_v139_}))
                d_1325_v141_: _dafny.Array
                def lambda45_(d_1326_v138_):
                    def lambda46_(d_1327_i13_):
                        return d_1326_v138_

                    return lambda46_

                init27_ = lambda45_(d_1322_v138_)
                nw199_ = _dafny.Array(None, 27)
                for i0_27_ in range(nw199_.length(0)):
                    nw199_[i0_27_] = init27_(i0_27_)
                d_1325_v141_ = nw199_
                d_1325_v141_ = d_1325_v141_
            (globalState).f7 = (default__.fm0((self).f25, globalState)) <= (((0) - (d_1250_cf54_)) * (len((self).fm3(-922, (self).f15, (self).f25, globalState))))
            d_1165_v50_ = d_1251_cf53_
        elif source19_.is_DC28:
            d_1328___mcc_h8_ = source19_.cf47
            d_1329_cf47_ = d_1328___mcc_h8_
            d_1330_v142_: _dafny.MultiSet
            d_1330_v142_ = _dafny.MultiSet([(self).f15, (self).fm26(False, globalState)])
            if (d_1330_v142_).issubset(d_1330_v142_):
                d_1331_v143_: _dafny.Map
                d_1331_v143_ = _dafny.Map({(self).f11: default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvena"))), (self).f14)})
                d_1331_v143_ = (d_1331_v143_).set(-401, len(self.f16))
                d_1332_v144_: _dafny.Map
                d_1332_v144_ = _dafny.Map({(0) - ((len(_dafny.SeqWithoutIsStrInference([(self).f25]))) - ((self).f14)): d_1330_v142_})
                d_1333_v145_: _dafny.Seq
                d_1333_v145_ = _dafny.SeqWithoutIsStrInference([not((self).f25)])
                d_1332_v144_ = (d_1332_v144_).set((self).f11, _dafny.MultiSet(d_1333_v145_))
                d_1334_v146_: _dafny.Array
                nw200_ = _dafny.Array(int(0), 20)
                d_1334_v146_ = nw200_
                d_1335_v147_: _dafny.Array
                d_1336_v148_: _dafny.Map
                d_1337_v149_: _dafny.Seq
                d_1338_v150_: str
                out48_: _dafny.Array
                out49_: _dafny.Map
                out50_: _dafny.Seq
                out51_: str
                out48_, out49_, out50_, out51_ = (self).m18(d_1334_v146_, (self).f12, globalState)
                d_1335_v147_ = out48_
                d_1336_v148_ = out49_
                d_1337_v149_ = out50_
                d_1338_v150_ = out51_
                d_1339_v151_: _dafny.Map
                d_1339_v151_ = _dafny.Map({len(d_1333_v145_): (self).f25})
                d_1340_v152_: _dafny.Map
                d_1340_v152_ = _dafny.Map({d_1339_v151_: 941})
                (globalState).f2 = ((d_1340_v152_) | (d_1340_v152_)) != (d_1340_v152_)
                d_1341_v153_: D10
                d_1341_v153_ = D10_DC30(False, True, (self).f26, (self).f25)
                d_1339_v151_ = (d_1339_v151_).set((self).f14, (d_1341_v153_).cf49)
            elif True:
                d_1342_v154_: _dafny.Map
                d_1342_v154_ = _dafny.Map({default__.fm0((self).f15, globalState): (self).f15})
                index220_ = default__.safeIndex(480, (p0).length(0))
                (p0)[index220_] = (self).f25
                d_1343_v155_: _dafny.Set
                d_1343_v155_ = _dafny.Set({(self).f12})
                d_1344_v157_: _dafny.MultiSet
                d_1344_v157_ = _dafny.MultiSet([d_1342_v154_, d_1342_v154_, d_1342_v154_, _dafny.Map({(0) - ((self).f14): (self).f25}), d_1342_v154_])
                index221_ = default__.safeIndex(480, (p0).length(0))
                def iife72_():
                    coll40_ = _dafny.Map()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(446, 808):
                        d_1345_v156_: int = compr_40_
                        if ((446) <= (d_1345_v156_)) and ((d_1345_v156_) < (808)):
                            coll40_[default__.safeModuloInt(d_1345_v156_, (self).f11)] = (self).f25
                    return _dafny.Map(coll40_)
                rhs208_ = (d_1342_v154_) | (iife72_()
                )
                rhs209_ = (self).f15
                rhs210_ = (self).f25
                rhs211_ = d_1343_v155_
                rhs212_ = (d_1344_v157_).isdisjoint((d_1344_v157_) - (_dafny.MultiSet([d_1342_v154_, d_1342_v154_])))
                lhs157_ = p0
                lhs158_ = default__.safeIndex(480, (p0).length(0))
                lhs159_ = globalState
                lhs160_ = globalState
                d_1342_v154_ = rhs208_
                lhs157_[lhs158_] = rhs209_
                lhs159_.f7 = rhs210_
                d_1343_v155_ = rhs211_
                lhs160_.f7 = rhs212_
                d_1346_v158_: _dafny.Map
                d_1346_v158_ = _dafny.Map({True: (p0)[default__.safeIndex(480, (p0).length(0))]})
                (globalState).f1 = (_dafny.Map({(self).fm1((self).f12, (p0)[default__.safeIndex(480, (p0).length(0))], (p0)[default__.safeIndex(480, (p0).length(0))], globalState): (self).f25})) | (_dafny.Map({(self).f15: ((d_1346_v158_)[(self).f15] if ((self).f15) in (d_1346_v158_) else ((d_1346_v158_)[not(not(True))] if (not(not(True))) in (d_1346_v158_) else (self).f25))}))
                (globalState).f2 = False
                index222_ = default__.safeIndex(480, (p0).length(0))
                (p0)[index222_] = (self).f15
                d_1347_v159_: _dafny.Array
                nw201_ = _dafny.Array(None, 14)
                d_1347_v159_ = nw201_
                d_1348_v160_: C0
                nw202_ = C0()
                nw202_.ctor__()
                d_1348_v160_ = nw202_
                index223_ = default__.safeIndex(516, (d_1347_v159_).length(0))
                (d_1347_v159_)[index223_] = (d_1348_v160_ if (self).f15 else d_1348_v160_)
                index224_ = default__.safeIndex(516, (d_1347_v159_).length(0))
                (d_1347_v159_)[index224_] = d_1348_v160_
            index225_ = default__.safeIndex(515, (p0).length(0))
            (p0)[index225_] = True
            index226_ = default__.safeIndex(515, (p0).length(0))
            (p0)[index226_] = (self).fm1(481, (self).f15, (self).f15, globalState)
            d_1349_v161_: _dafny.Seq
            d_1349_v161_ = _dafny.SeqWithoutIsStrInference([(self).f15, (p0)[default__.safeIndex(515, (p0).length(0))]])
            d_1350_v162_: _dafny.Array
            nw203_ = _dafny.Array(None, 25)
            nw203_[int(0)] = (self).f25
            nw203_[int(1)] = (p0)[default__.safeIndex(515, (p0).length(0))]
            nw203_[int(2)] = (self).f25
            nw203_[int(3)] = (self).f15
            nw203_[int(4)] = (p0)[default__.safeIndex(515, (p0).length(0))]
            nw203_[int(5)] = (self).f15
            nw203_[int(6)] = (self).f25
            nw203_[int(7)] = (self).f25
            nw203_[int(8)] = (self).f15
            nw203_[int(9)] = (p0)[default__.safeIndex(515, (p0).length(0))]
            nw203_[int(10)] = (p0)[default__.safeIndex(515, (p0).length(0))]
            nw203_[int(11)] = (self).f15
            nw203_[int(12)] = (p0)[default__.safeIndex(515, (p0).length(0))]
            nw203_[int(13)] = (self).f15
            nw203_[int(14)] = not((self).f15)
            nw203_[int(15)] = (self).f25
            nw203_[int(16)] = (p0)[default__.safeIndex(515, (p0).length(0))]
            nw203_[int(17)] = False
            nw203_[int(18)] = not((self).fm1((self).f14, (p0)[default__.safeIndex(515, (p0).length(0))], (self).f25, globalState))
            nw203_[int(19)] = (self).f15
            nw203_[int(20)] = (self).fm1(len(d_1349_v161_), (self).f25, False, globalState)
            nw203_[int(21)] = (self).f15
            nw203_[int(22)] = (p0)[default__.safeIndex(515, (p0).length(0))]
            nw203_[int(23)] = (self).f25
            nw203_[int(24)] = (p0)[default__.safeIndex(515, (p0).length(0))]
            d_1350_v162_ = nw203_
            d_1351_v163_: _dafny.Array
            def lambda47_(d_1352_i14_):
                return (d_1352_i14_) - ((self).f14)

            init28_ = lambda47_
            nw204_ = _dafny.Array(None, 10)
            for i0_28_ in range(nw204_.length(0)):
                nw204_[i0_28_] = init28_(i0_28_)
            d_1351_v163_ = nw204_
            d_1353_v164_: D6
            d_1353_v164_ = D6_DC21((p0)[default__.safeIndex(515, (p0).length(0))])
            (self).m17((d_1350_v162_ if (p0)[default__.safeIndex(515, (p0).length(0))] else d_1350_v162_), d_1351_v163_, (d_1351_v163_ if (self).fm26((d_1353_v164_).cf40, globalState) else d_1351_v163_), globalState)
            index227_ = default__.safeIndex(515, (p0).length(0))
            (p0)[index227_] = (p0)[default__.safeIndex(515, (p0).length(0))]
        elif True:
            d_1354___mcc_h9_ = source19_.cf56
            d_1355_cf56_ = d_1354___mcc_h9_
            (globalState).f6 = 288
            (globalState).f2 = not (((p1).set((self).f14, default__.abs((self).f12))) == (p1)) or ((self).f15)
            d_1356_v165_: _dafny.MultiSet
            d_1356_v165_ = _dafny.MultiSet([False])
            d_1357_v166_: bool
            d_1358_v167_: _dafny.Seq
            out52_: bool
            out53_: _dafny.Seq
            out52_, out53_ = default__.m0((_dafny.MultiSet([(self).f15])) != (d_1356_v165_), (self).f15, (self).f11, ((self).f11) - ((self).f12), globalState)
            d_1357_v166_ = out52_
            d_1358_v167_ = out53_
            d_1359_v168_: C0
            nw205_ = C0()
            nw205_.ctor__()
            d_1359_v168_ = nw205_
        d_1360_v169_: _dafny.Seq
        d_1360_v169_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_1361_v170_: _dafny.Array
        nw206_ = _dafny.Array(None, 11)
        nw206_[int(0)] = (self).f12
        nw206_[int(1)] = (self).f14
        nw206_[int(2)] = (self).f14
        nw206_[int(3)] = (self).f12
        nw206_[int(4)] = (self).f11
        nw206_[int(5)] = (self).f11
        nw206_[int(6)] = (self).f14
        nw206_[int(7)] = len(d_1360_v169_)
        nw206_[int(8)] = (self).f11
        nw206_[int(9)] = (self).f11
        nw206_[int(10)] = (self).f14
        d_1361_v170_ = nw206_
        d_1362_v171_: _dafny.Seq
        d_1362_v171_ = _dafny.SeqWithoutIsStrInference([d_1361_v170_])
        d_1363_v172_: _dafny.Array
        nw207_ = _dafny.Array(None, 14)
        nw207_[int(0)] = d_1361_v170_
        nw207_[int(1)] = d_1361_v170_
        nw207_[int(2)] = d_1361_v170_
        nw207_[int(3)] = (d_1362_v171_)[default__.safeIndex((self).f11, len(d_1362_v171_))]
        nw207_[int(4)] = d_1361_v170_
        nw207_[int(5)] = d_1361_v170_
        nw207_[int(6)] = d_1361_v170_
        nw207_[int(7)] = d_1361_v170_
        nw207_[int(8)] = d_1361_v170_
        nw207_[int(9)] = d_1361_v170_
        nw207_[int(10)] = d_1361_v170_
        nw207_[int(11)] = d_1361_v170_
        nw207_[int(12)] = d_1361_v170_
        nw207_[int(13)] = d_1361_v170_
        d_1363_v172_ = nw207_
        d_1363_v172_ = d_1363_v172_
        d_1364_v173_: _dafny.Seq
        d_1364_v173_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])
        d_1365_v174_: D4
        d_1365_v174_ = D4_DC13(d_1364_v173_)
        pat_let_tv14_ = d_1364_v173_
        pat_let_tv15_ = d_1364_v173_
        def iife73_(_pat_let16_0):
            def iife74_(d_1366_dt__update__tmp_h0_):
                def iife75_(_pat_let17_0):
                    def iife76_(d_1367_dt__update_hcf29_h0_):
                        return D4_DC13(d_1367_dt__update_hcf29_h0_)
                    return iife76_(_pat_let17_0)
                return iife75_((pat_let_tv14_ if (self).f15 else pat_let_tv15_))
            return iife74_(_pat_let16_0)
        source22_ = iife73_(d_1365_v174_)
        if source22_.is_DC14:
            d_1368___mcc_h30_ = source22_.cf30
            d_1369___mcc_h31_ = source22_.cf31
            d_1370_cf31_ = d_1369___mcc_h31_
            d_1371_cf30_ = d_1368___mcc_h30_
            d_1372_v175_: _dafny.Array
            nw208_ = _dafny.Array(_dafny.MultiSet({}), 8)
            d_1372_v175_ = nw208_
            d_1373_v176_: _dafny.Seq
            d_1373_v176_ = _dafny.SeqWithoutIsStrInference([d_1372_v175_])
            d_1372_v175_ = (d_1373_v176_)[default__.safeIndex((self).f12, len(d_1373_v176_))]
            d_1374_v177_: _dafny.Map
            d_1374_v177_ = _dafny.Map({d_1370_cf31_: (self).f15})
            d_1375_v178_: _dafny.Map
            d_1375_v178_ = _dafny.Map({(len(_dafny.Map({d_1371_cf30_: True}))) > ((self).f11): (((d_1374_v177_)[-477] if (-477) in (d_1374_v177_) else (self).f15)) not in ((default__.fm30((self).f25, (self).f15, (0) - ((0) - (d_1370_cf31_)), (self).f15, globalState)).set(default__.safeIndex((self).f11, len(default__.fm30((self).f25, (self).f15, (0) - ((0) - (d_1370_cf31_)), (self).f15, globalState))), True))})
            d_1376_v179_: _dafny.MultiSet
            d_1376_v179_ = _dafny.MultiSet([(self).f25, False])
            d_1377_v180_: T2
            nw209_ = C1()
            nw209_.ctor__(p0, 247, d_1370_cf31_, (self).f14, (self).f15, _dafny.SeqWithoutIsStrInference([d_1165_v50_ for d_1378_i15_ in range(default__.abs(469))]))
            d_1377_v180_ = nw209_
            d_1379_v181_: D11
            d_1379_v181_ = D11_DC33(d_1377_v180_)
            d_1380_v182_: _dafny.Array
            nw210_ = _dafny.Array(None, 19)
            nw210_[int(0)] = d_1377_v180_
            nw210_[int(1)] = d_1377_v180_
            nw210_[int(2)] = d_1377_v180_
            nw210_[int(3)] = d_1377_v180_
            nw210_[int(4)] = d_1377_v180_
            nw210_[int(5)] = d_1377_v180_
            nw210_[int(6)] = d_1377_v180_
            nw210_[int(7)] = d_1377_v180_
            nw210_[int(8)] = d_1377_v180_
            nw210_[int(9)] = d_1377_v180_
            nw210_[int(10)] = d_1377_v180_
            nw210_[int(11)] = d_1377_v180_
            nw210_[int(12)] = (d_1379_v181_).cf57
            nw210_[int(13)] = d_1377_v180_
            nw210_[int(14)] = d_1377_v180_
            nw210_[int(15)] = d_1377_v180_
            nw210_[int(16)] = d_1377_v180_
            nw210_[int(17)] = d_1377_v180_
            nw210_[int(18)] = d_1377_v180_
            d_1380_v182_ = nw210_
            d_1381_v183_: _dafny.Set
            d_1381_v183_ = _dafny.Set({d_1380_v182_})
            d_1382_v184_: D1
            d_1382_v184_ = D1_DC2(_dafny.CodePoint('c'))
            d_1383_v185_: D2
            d_1383_v185_ = D2_DC7(p0, d_1381_v183_, d_1382_v184_, (self).f11, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wf"))))
            (globalState).f7 = ((d_1375_v178_)[not ((self).f25) or ((self).fm1(335, (self).f25, (self).f15, globalState))] if (not ((self).f25) or ((self).fm1(335, (self).f25, (self).f15, globalState))) in (d_1375_v178_) else ((d_1376_v179_).set((self).f15, default__.abs((d_1383_v185_).cf13))).ispropersubset(d_1376_v179_))
            d_1370_cf31_ = (d_1364_v173_)[default__.safeIndex((d_1377_v180_).f12, len(d_1364_v173_))]
            d_1384_v186_: D2
            d_1384_v186_ = D2_DC6(_dafny.SeqWithoutIsStrInference([d_1165_v50_ for d_1385_i16_ in range(default__.abs(265))]))
            d_1386_v187_: _dafny.Seq
            d_1386_v187_ = _dafny.SeqWithoutIsStrInference([(self).f13, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqkjge")), (d_1377_v180_).f13, (d_1384_v186_).cf9, (d_1377_v180_).f13])
            (self).f16 = ((self).f26) + (((d_1386_v187_)[default__.safeIndex((0) - ((d_1377_v180_).f12), len(d_1386_v187_))]).set(default__.safeIndex((d_1377_v180_).f14, len((d_1386_v187_)[default__.safeIndex((0) - ((d_1377_v180_).f12), len(d_1386_v187_))])), d_1165_v50_))
        elif source22_.is_DC15:
            d_1387___mcc_h32_ = source22_.cf32
            d_1388_cf32_ = d_1387___mcc_h32_
            rhs213_ = len(self.f16)
            rhs214_ = (self).f15
            lhs161_ = globalState
            lhs162_ = globalState
            lhs161_.f6 = rhs213_
            lhs162_.f7 = rhs214_
            d_1389_v188_: C0
            nw211_ = C0()
            nw211_.ctor__()
            d_1389_v188_ = nw211_
            d_1390_v189_: _dafny.Map
            d_1390_v189_ = _dafny.Map({(self).f15: (self).f25})
            d_1391_v190_: _dafny.Set
            d_1391_v190_ = _dafny.Set({_dafny.Map({(self).f25: (self).f25}), d_1390_v189_})
            d_1388_cf32_ = (0) - ((self).fm2(-207, (_dafny.Set({_dafny.Map({(self).f15: (self).f25})})).isdisjoint(d_1391_v190_), (d_1388_cf32_ if (self).f15 else (self).f12), ((self).f12) >= (len(d_1360_v169_)), globalState))
            d_1392_v191_: _dafny.Array
            nw212_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
            d_1392_v191_ = nw212_
            d_1393_v192_: _dafny.Map
            d_1393_v192_ = _dafny.Map({d_1392_v191_: (self).f15})
            d_1394_v193_: _dafny.Map
            d_1394_v193_ = _dafny.Map({d_1361_v170_: (self).f14})
            d_1395_v194_: _dafny.Set
            d_1395_v194_ = _dafny.Set({True})
            (globalState).f2 = (((d_1393_v192_)[d_1392_v191_] if (d_1392_v191_) in (d_1393_v192_) else (self).f15)) or ((d_1394_v193_) != ((d_1394_v193_).set(d_1361_v170_, len(_dafny.Map({(self).f14: len(d_1395_v194_)})))))
        elif source22_.is_DC13:
            d_1396___mcc_h33_ = source22_.cf29
            d_1397_cf29_ = d_1396___mcc_h33_
            index228_ = default__.safeIndex(606, (p0).length(0))
            (p0)[index228_] = (self).f15
            index229_ = default__.safeIndex(606, (p0).length(0))
            (p0)[index229_] = (self).f25
            (globalState).f6 = (self).f14
            d_1398_v195_: _dafny.Array
            nw213_ = _dafny.Array(D10.default()(), 22)
            d_1398_v195_ = nw213_
            d_1399_v196_: _dafny.MultiSet
            d_1399_v196_ = _dafny.MultiSet([d_1398_v195_])
            d_1400_v197_: _dafny.Map
            d_1400_v197_ = _dafny.Map({(self).f14: d_1399_v196_})
            d_1401_v198_: _dafny.Map
            d_1401_v198_ = _dafny.Map({(self).f11: (self).f12})
            d_1400_v197_ = (d_1400_v197_).set(((d_1401_v198_)[(self).f11] if ((self).f11) in (d_1401_v198_) else (self).f12), (d_1399_v196_).set(d_1398_v195_, default__.abs((self).f12)))
            d_1402_v199_: _dafny.Map
            d_1402_v199_ = _dafny.Map({928: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xw"))})
            d_1403_v200_: _dafny.Map
            d_1403_v200_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqso"))): _dafny.Set({_dafny.CodePoint('e'), d_1165_v50_, d_1165_v50_})})
            d_1402_v199_ = (d_1402_v199_).set(len((d_1403_v200_) | (d_1403_v200_)), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lp"))).set(default__.safeIndex((self).f12, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lp")))), d_1165_v50_))
        elif True:
            d_1404___mcc_h34_ = source22_.cf33
            d_1405_cf33_ = d_1404___mcc_h34_
            r0 = False
            index230_ = default__.safeIndex(542, (p0).length(0))
            (p0)[index230_] = (self).f15
            index231_ = default__.safeIndex(542, (p0).length(0))
            (p0)[index231_] = (self).f15
            d_1406_v201_: _dafny.Array
            nw214_ = _dafny.Array(None, 1)
            nw214_[int(0)] = d_1364_v173_
            d_1406_v201_ = nw214_
            index232_ = default__.safeIndex(506, (d_1406_v201_).length(0))
            (d_1406_v201_)[index232_] = d_1364_v173_
            d_1407_v202_: _dafny.Map
            d_1407_v202_ = _dafny.Map({not((self).f25): (self).fm26((self).f15, globalState)})
            index233_ = default__.safeIndex(542, (p0).length(0))
            index234_ = default__.safeIndex(506, (d_1406_v201_).length(0))
            rhs215_ = d_1361_v170_
            rhs216_ = (self).fm1(((self).f14) - (len(_dafny.Set({(self).f25}))), ((self).f11) <= ((self).f11), (self).f15, globalState)
            rhs217_ = (_dafny.SeqWithoutIsStrInference([(self).f12, len(self.f16)])) + (d_1364_v173_)
            rhs218_ = len(default__.fm38((self).f12, (self).f26, ((d_1407_v202_)[(self).f25] if ((self).f25) in (d_1407_v202_) else (p0)[default__.safeIndex(542, (p0).length(0))]), globalState))
            rhs219_ = (self).f14
            lhs163_ = p0
            lhs164_ = default__.safeIndex(542, (p0).length(0))
            lhs165_ = d_1406_v201_
            lhs166_ = default__.safeIndex(506, (d_1406_v201_).length(0))
            lhs167_ = globalState
            lhs168_ = globalState
            d_1361_v170_ = rhs215_
            lhs163_[lhs164_] = rhs216_
            lhs165_[lhs166_] = rhs217_
            lhs167_.f6 = rhs218_
            lhs168_.f6 = rhs219_
            rhs220_ = (d_1360_v169_)[default__.safeIndex(((self).f12) + ((self).f12), len(d_1360_v169_))]
            rhs221_ = False
            lhs169_ = globalState
            lhs170_ = globalState
            lhs169_.f7 = rhs220_
            lhs170_.f2 = rhs221_
        d_1408_v203_: _dafny.Map
        d_1408_v203_ = _dafny.Map({False: (self).f26})
        d_1409_v204_: _dafny.Array
        nw215_ = _dafny.Array(None, 23)
        nw215_[int(0)] = self.f16
        nw215_[int(1)] = (self).f26
        nw215_[int(2)] = ((self).f26) + (self.f16)
        nw215_[int(3)] = (self).f13
        nw215_[int(4)] = (self).f13
        nw215_[int(5)] = (self).f26
        nw215_[int(6)] = (((self).f26) + (((self).f13).set(default__.safeIndex((self).f12, len((self).f13)), d_1165_v50_))).set(default__.safeIndex(-868, len(((self).f26) + (((self).f13).set(default__.safeIndex((self).f12, len((self).f13)), d_1165_v50_)))), d_1165_v50_)
        nw215_[int(7)] = ((d_1408_v203_)[(self).f25] if ((self).f25) in (d_1408_v203_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dffjfo")))
        nw215_[int(8)] = (self).f13
        nw215_[int(9)] = (self).f13
        nw215_[int(10)] = (self).f26
        nw215_[int(11)] = (self).f13
        nw215_[int(12)] = (((self).f13).set(default__.safeIndex((self).f12, len((self).f13)), d_1165_v50_)) + (_dafny.SeqWithoutIsStrInference([d_1165_v50_ for d_1410_i17_ in range(default__.abs(791))]))
        nw215_[int(13)] = (self).f26
        nw215_[int(14)] = self.f16
        nw215_[int(15)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1411_i18_ in range(default__.abs(-584))])
        nw215_[int(16)] = (self).f26
        nw215_[int(17)] = self.f16
        nw215_[int(18)] = ((self).f13) + (self.f16)
        nw215_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwefadivg"))
        nw215_[int(20)] = (self).f26
        nw215_[int(21)] = (self).f13
        nw215_[int(22)] = ((self).f13) + ((self).f26)
        d_1409_v204_ = nw215_
        index235_ = default__.safeIndex(794, (d_1409_v204_).length(0))
        (d_1409_v204_)[index235_] = (self.f16) + ((self).f13)
        pat_let_tv16_ = d_1165_v50_
        pat_let_tv17_ = d_1165_v50_
        index236_ = default__.safeIndex(794, (d_1409_v204_).length(0))
        def lambda48_(source23_):
            if source23_.is_DC7:
                d_1412___mcc_h35_ = source23_.cf10
                d_1413___mcc_h36_ = source23_.cf11
                d_1414___mcc_h37_ = source23_.cf12
                d_1415___mcc_h38_ = source23_.cf13
                d_1416___mcc_h39_ = source23_.cf14
                d_1417_cf14_ = d_1416___mcc_h39_
                d_1418_cf13_ = d_1415___mcc_h38_
                d_1419_cf12_ = d_1414___mcc_h37_
                d_1420_cf11_ = d_1413___mcc_h36_
                d_1421_cf10_ = d_1412___mcc_h35_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmmrxafmt"))
            elif source23_.is_DC8:
                d_1422___mcc_h40_ = source23_.cf15
                d_1423___mcc_h41_ = source23_.cf16
                d_1424___mcc_h42_ = source23_.cf17
                d_1425_cf17_ = d_1424___mcc_h42_
                d_1426_cf16_ = d_1423___mcc_h41_
                d_1427_cf15_ = d_1422___mcc_h40_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfxans"))
            elif True:
                d_1428___mcc_h43_ = source23_.cf9
                d_1429_cf9_ = d_1428___mcc_h43_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv16_ for d_1430_i19_ in range(default__.abs(668))])

        def lambda49_(source24_):
            if source24_.is_DC7:
                d_1431___mcc_h44_ = source24_.cf10
                d_1432___mcc_h45_ = source24_.cf11
                d_1433___mcc_h46_ = source24_.cf12
                d_1434___mcc_h47_ = source24_.cf13
                d_1435___mcc_h48_ = source24_.cf14
                d_1436_cf14_ = d_1435___mcc_h48_
                d_1437_cf13_ = d_1434___mcc_h47_
                d_1438_cf12_ = d_1433___mcc_h46_
                d_1439_cf11_ = d_1432___mcc_h45_
                d_1440_cf10_ = d_1431___mcc_h44_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmmrxafmt"))
            elif source24_.is_DC8:
                d_1441___mcc_h49_ = source24_.cf15
                d_1442___mcc_h50_ = source24_.cf16
                d_1443___mcc_h51_ = source24_.cf17
                d_1444_cf17_ = d_1443___mcc_h51_
                d_1445_cf16_ = d_1442___mcc_h50_
                d_1446_cf15_ = d_1441___mcc_h49_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfxans"))
            elif True:
                d_1447___mcc_h52_ = source24_.cf9
                d_1448_cf9_ = d_1447___mcc_h52_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv17_ for d_1449_i19_ in range(default__.abs(668))])

        (d_1409_v204_)[index236_] = (lambda48_(default__.fm39(globalState))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f25, (self).f15, (self).f15, (self).f25, (self).fm26((self).f15, globalState)])), len(lambda49_(default__.fm39(globalState)))), d_1165_v50_)
        r0 = (self).f15
        d_1450_v205_: D2
        d_1450_v205_ = D2_DC6(self.f16)
        pat_let_tv18_ = p1
        pat_let_tv19_ = p1
        pat_let_tv20_ = p1
        def lambda50_(source25_):
            if source25_.is_DC7:
                d_1451___mcc_h53_ = source25_.cf10
                d_1452___mcc_h54_ = source25_.cf11
                d_1453___mcc_h55_ = source25_.cf12
                d_1454___mcc_h56_ = source25_.cf13
                d_1455___mcc_h57_ = source25_.cf14
                d_1456_cf14_ = d_1455___mcc_h57_
                d_1457_cf13_ = d_1454___mcc_h56_
                d_1458_cf12_ = d_1453___mcc_h55_
                d_1459_cf11_ = d_1452___mcc_h54_
                d_1460_cf10_ = d_1451___mcc_h53_
                return (_dafny.MultiSet([(self).f12, (self).f12, d_1456_cf14_, (self).f11, d_1457_cf13_])) | (pat_let_tv18_)
            elif source25_.is_DC8:
                d_1461___mcc_h58_ = source25_.cf15
                d_1462___mcc_h59_ = source25_.cf16
                d_1463___mcc_h60_ = source25_.cf17
                d_1464_cf17_ = d_1463___mcc_h60_
                d_1465_cf16_ = d_1462___mcc_h59_
                d_1466_cf15_ = d_1461___mcc_h58_
                return pat_let_tv19_
            elif True:
                d_1467___mcc_h61_ = source25_.cf9
                d_1468_cf9_ = d_1467___mcc_h61_
                return pat_let_tv20_

        r1 = lambda50_(d_1450_v205_)
        return r0, r1

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        d_1469_v0_: C0
        nw216_ = C0()
        nw216_.ctor__()
        d_1469_v0_ = nw216_
        d_1470_v1_: _dafny.Map
        d_1470_v1_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p1]): p2})
        d_1471_v2_: _dafny.Seq
        d_1471_v2_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1472_v3_: _dafny.Seq
        d_1472_v3_ = _dafny.SeqWithoutIsStrInference([len(d_1471_v2_)])
        d_1473_v4_: _dafny.Array
        nw217_ = _dafny.Array(None, 6)
        nw217_[int(0)] = not (p2) or ((self).f25)
        nw217_[int(1)] = ((self).f15) == ((self).f15)
        nw217_[int(2)] = True
        nw217_[int(3)] = p2
        nw217_[int(4)] = (p1) > ((self).f11)
        nw217_[int(5)] = ((d_1470_v1_)[d_1472_v3_] if (d_1472_v3_) in (d_1470_v1_) else (self).f25)
        d_1473_v4_ = nw217_
        d_1474_v5_: _dafny.MultiSet
        d_1474_v5_ = _dafny.MultiSet([(self).f15])
        d_1475_v6_: _dafny.Map
        d_1475_v6_ = _dafny.Map({p2: p2})
        d_1476_v7_: _dafny.Map
        d_1476_v7_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_1469_v0_).fm17(_dafny.SeqWithoutIsStrInference([(d_1474_v5_).cardinality]), globalState), ((d_1475_v6_)[(self).f25] if ((self).f25) in (d_1475_v6_) else (self).f15)])): (self).f14})
        index237_ = default__.safeIndex(465, (d_1473_v4_).length(0))
        (d_1473_v4_)[index237_] = (d_1476_v7_) == (d_1476_v7_)
        d_1477_v8_: _dafny.Array
        def lambda51_(d_1478_p1_):
            def lambda52_(d_1479_i0_):
                return default__.safeModuloInt(d_1479_i0_, d_1478_p1_)

            return lambda52_

        init29_ = lambda51_(p1)
        nw218_ = _dafny.Array(None, 27)
        for i0_29_ in range(nw218_.length(0)):
            nw218_[i0_29_] = init29_(i0_29_)
        d_1477_v8_ = nw218_
        d_1480_v9_: _dafny.Map
        d_1480_v9_ = _dafny.Map({p2: (self).f11})
        d_1481_v10_: _dafny.Seq
        d_1481_v10_ = _dafny.SeqWithoutIsStrInference([d_1480_v9_])
        d_1482_v11_: D6
        d_1482_v11_ = D6_DC19((d_1481_v10_)[default__.safeIndex((self).f12, len(d_1481_v10_))])
        d_1483_v12_: _dafny.Map
        d_1483_v12_ = _dafny.Map({(self).fm2((self).f14, True, 285, (self).f25, globalState): ((self).f14) > ((self).f11)})
        d_1484_v13_: _dafny.Seq
        d_1484_v13_ = _dafny.SeqWithoutIsStrInference([d_1477_v8_, d_1477_v8_, d_1477_v8_, d_1477_v8_, d_1477_v8_])
        index238_ = default__.safeIndex(465, (d_1473_v4_).length(0))
        rhs222_ = (self).f14
        rhs223_ = ((d_1483_v12_)[(len((d_1471_v2_).set(default__.safeIndex((self).f11, len(d_1471_v2_)), p2)) if False else p1)] if ((len((d_1471_v2_).set(default__.safeIndex((self).f11, len(d_1471_v2_)), p2)) if False else p1)) in (d_1483_v12_) else (self).f25)
        rhs224_ = (d_1484_v13_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference([(self).f15])) + (d_1471_v2_)), len(d_1484_v13_))]
        rhs225_ = d_1482_v11_
        lhs171_ = globalState
        lhs172_ = d_1473_v4_
        lhs173_ = default__.safeIndex(465, (d_1473_v4_).length(0))
        lhs171_.f6 = rhs222_
        lhs172_[lhs173_] = rhs223_
        d_1477_v8_ = rhs224_
        d_1482_v11_ = rhs225_
        d_1485_v14_: _dafny.Set
        d_1485_v14_ = _dafny.Set({(d_1471_v2_)[default__.safeIndex(704, len(d_1471_v2_))]})
        d_1486_v15_: D12
        d_1486_v15_ = D12_DC35(d_1485_v14_)
        d_1487_i1_: int
        d_1487_i1_ = 0
        with _dafny.label("7"):
            while ((_dafny.Set({(d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))]}) if (self).f15 else d_1485_v14_)).issubset((d_1486_v15_).cf63):
                with _dafny.c_label("7"):
                    if (d_1487_i1_) >= (100):
                        raise _dafny.Break("7")
                    d_1487_i1_ = (d_1487_i1_) + (1)
                    d_1488_v16_: str
                    d_1488_v16_ = _dafny.CodePoint('i')
                    d_1488_v16_ = d_1488_v16_
                    if (d_1471_v2_)[default__.safeIndex((-868) + ((0) - ((self).f12)), len(d_1471_v2_))]:
                        (globalState).f2 = (self).f15
                        index239_ = default__.safeIndex(556, (d_1477_v8_).length(0))
                        (d_1477_v8_)[index239_] = (self).f11
                        index240_ = default__.safeIndex(556, (d_1477_v8_).length(0))
                        (d_1477_v8_)[index240_] = (self).f11
                        (globalState).f6 = (default__.safeModuloInt((d_1477_v8_)[default__.safeIndex(556, (d_1477_v8_).length(0))], len(d_1471_v2_))) + ((self).f12)
                        (globalState).f7 = not((self).f25)
                        d_1489_v17_: _dafny.Array
                        nw219_ = _dafny.Array(_dafny.Array(None, 0), 24)
                        d_1489_v17_ = nw219_
                        d_1490_v18_: D4
                        d_1490_v18_ = D4_DC14(d_1488_v16_, (self).f14)
                        d_1491_v19_: _dafny.Set
                        d_1491_v19_ = _dafny.Set({(d_1477_v8_)[default__.safeIndex(556, (d_1477_v8_).length(0))]})
                        d_1492_v20_: _dafny.Seq
                        d_1492_v20_ = _dafny.SeqWithoutIsStrInference([d_1491_v19_])
                        d_1493_v21_: T1
                        nw220_ = C2()
                        nw220_.ctor__((self).f11, (p0).set(default__.safeIndex(len(_dafny.Set({d_1472_v3_})), len(p0)), (d_1490_v18_).cf30), p3, ((d_1474_v5_)[(d_1471_v2_)[default__.safeIndex(len((d_1492_v20_)[default__.safeIndex(len(d_1471_v2_), len(d_1492_v20_))]), len(d_1471_v2_))]] if ((d_1471_v2_)[default__.safeIndex(len((d_1492_v20_)[default__.safeIndex(len(d_1471_v2_), len(d_1492_v20_))]), len(d_1471_v2_))]) in (d_1474_v5_) else 435))
                        d_1493_v21_ = nw220_
                        d_1494_v22_: _dafny.Seq
                        d_1494_v22_ = _dafny.SeqWithoutIsStrInference([d_1493_v21_, (d_1493_v21_ if (d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))] else d_1493_v21_), d_1493_v21_, d_1493_v21_])
                        d_1495_v23_: _dafny.Map
                        d_1495_v23_ = _dafny.Map({(d_1477_v8_)[default__.safeIndex(556, (d_1477_v8_).length(0))]: d_1476_v7_})
                        rhs226_ = default__.safeModuloInt(len(d_1472_v3_), (d_1493_v21_).f12)
                        rhs227_ = d_1489_v17_
                        rhs228_ = (default__.safeModuloInt(len(d_1495_v23_), (0) - ((self).f11))) == ((self).f11)
                        rhs229_ = (d_1494_v22_) + (d_1494_v22_)
                        lhs174_ = globalState
                        lhs175_ = globalState
                        lhs174_.f6 = rhs226_
                        d_1489_v17_ = rhs227_
                        lhs175_.f7 = rhs228_
                        d_1494_v22_ = rhs229_
                    elif True:
                        d_1496_v24_: _dafny.Map
                        d_1496_v24_ = _dafny.Map({d_1471_v2_: not(p2)})
                        r0 = (p2) == (((d_1496_v24_)[d_1471_v2_] if (d_1471_v2_) in (d_1496_v24_) else (self).fm26((d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))], globalState)))
                        d_1497_v25_: _dafny.Map
                        d_1497_v25_ = _dafny.Map({p2: (d_1475_v6_).set((self).f25, False)})
                        d_1498_v26_: D15
                        d_1498_v26_ = D15_DC42(d_1497_v25_)
                        d_1497_v25_ = (d_1498_v26_).cf76
                        d_1499_v27_: _dafny.Seq
                        d_1499_v27_ = _dafny.SeqWithoutIsStrInference([d_1485_v14_, d_1485_v14_])
                        r0 = ((d_1485_v14_) | (d_1485_v14_)).issubset((d_1499_v27_)[default__.safeIndex((self).f12, len(d_1499_v27_))])
                        index241_ = default__.safeIndex(910, (d_1477_v8_).length(0))
                        (d_1477_v8_)[index241_] = 967
                        index242_ = default__.safeIndex(910, (d_1477_v8_).length(0))
                        (d_1477_v8_)[index242_] = ((self).f11) - (375)
                        index243_ = default__.safeIndex(465, (d_1473_v4_).length(0))
                        rhs230_ = not ((self).f25) or ((self).f25)
                        rhs231_ = (self).f11
                        rhs232_ = d_1488_v16_
                        rhs233_ = default__.fm47(default__.safeDivisionInt(len(self.f16), (self).f14), globalState)
                        rhs234_ = d_1473_v4_
                        lhs176_ = d_1473_v4_
                        lhs177_ = default__.safeIndex(465, (d_1473_v4_).length(0))
                        lhs178_ = globalState
                        lhs176_[lhs177_] = rhs230_
                        lhs178_.f6 = rhs231_
                        d_1488_v16_ = rhs232_
                        d_1488_v16_ = rhs233_
                        d_1473_v4_ = rhs234_
                    d_1500_v28_: C1
                    nw221_ = C1()
                    nw221_.ctor__(d_1473_v4_, -387, (self).f14, len(_dafny.Map({d_1473_v4_: (self).f25})), (self).f25, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iyyrfnvqk")))
                    d_1500_v28_ = nw221_
                    (globalState).f7 = (p3) < (((d_1472_v3_)[default__.safeIndex((d_1500_v28_).fm2(len((self).f26), (d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))], p3, p2, globalState), len(d_1472_v3_))]) * (len(p0)))
                    pass
            pass
        d_1501_v29_: _dafny.Set
        d_1501_v29_ = _dafny.Set({default__.fm44((self).f12, (self).f25, globalState)})
        d_1501_v29_ = d_1501_v29_
        d_1502_v30_: D2
        d_1502_v30_ = D2_DC8(False, (d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))], p0)
        d_1503_v31_: str
        d_1503_v31_ = _dafny.CodePoint('t')
        d_1504_v32_: _dafny.MultiSet
        d_1504_v32_ = _dafny.MultiSet([d_1503_v31_])
        d_1505_v33_: D0
        d_1505_v33_ = D0_DC0(d_1504_v32_)
        d_1506_v34_: _dafny.Set
        d_1506_v34_ = _dafny.Set({d_1505_v33_})
        d_1507_v35_: D3
        d_1507_v35_ = D3_DC12(D3_DC11((self).f14, d_1502_v30_, d_1506_v34_, p2))
        source26_ = d_1507_v35_
        if source26_.is_DC10:
            d_1508___mcc_h0_ = source26_.cf19
            d_1509___mcc_h1_ = source26_.cf20
            d_1510___mcc_h2_ = source26_.cf21
            d_1511___mcc_h3_ = source26_.cf22
            d_1512___mcc_h4_ = source26_.cf23
            d_1513_cf23_ = d_1512___mcc_h4_
            d_1514_cf22_ = d_1511___mcc_h3_
            d_1515_cf21_ = d_1510___mcc_h2_
            d_1516_cf20_ = d_1509___mcc_h1_
            d_1517_cf19_ = d_1508___mcc_h0_
            index244_ = default__.safeIndex(199, (d_1477_v8_).length(0))
            (d_1477_v8_)[index244_] = len(d_1514_cf22_)
            d_1518_v36_: _dafny.Map
            d_1518_v36_ = _dafny.Map({d_1503_v31_: (0) - (p3)})
            index245_ = default__.safeIndex(199, (d_1477_v8_).length(0))
            def iife77_():
                coll41_ = _dafny.Set()
                compr_41_: int
                for compr_41_ in _dafny.IntegerRange(558, 704):
                    d_1519_v37_: int = compr_41_
                    if ((558) <= (d_1519_v37_)) and ((d_1519_v37_) < (704)):
                        coll41_ = coll41_.union(_dafny.Set([default__.safeModuloInt(d_1519_v37_, d_1513_cf23_)]))
                return _dafny.Set(coll41_)
            (d_1477_v8_)[index245_] = (default__.safeModuloInt(len(default__.fm32((self).f25, d_1518_v36_, iife77_()
            , (self).f15, globalState)), 789)) - ((p1) - (d_1513_cf23_))
            d_1520_v38_: D8
            d_1520_v38_ = D8_DC24(d_1474_v5_)
            d_1521_v39_: _dafny.Seq
            d_1521_v39_ = _dafny.SeqWithoutIsStrInference([d_1520_v38_])
            pat_let_tv21_ = d_1474_v5_
            def iife78_(_pat_let18_0):
                def iife79_(d_1522_dt__update__tmp_h0_):
                    def iife80_(_pat_let19_0):
                        def iife81_(d_1523_dt__update_hcf43_h0_):
                            return D8_DC24(d_1523_dt__update_hcf43_h0_)
                        return iife81_(_pat_let19_0)
                    return iife80_(pat_let_tv21_)
                return iife79_(_pat_let18_0)
            rhs235_ = (((self).f12) * ((self).f14)) + (len((d_1521_v39_) + (_dafny.SeqWithoutIsStrInference([iife78_(d_1520_v38_), d_1520_v38_, d_1520_v38_, d_1520_v38_]))))
            rhs236_ = (0) - (d_1516_cf20_)
            rhs237_ = d_1503_v31_
            d_1516_cf20_ = rhs235_
            r2 = rhs236_
            d_1503_v31_ = rhs237_
            d_1524_v40_: _dafny.Set
            d_1524_v40_ = _dafny.Set({d_1516_cf20_, (self).f12})
            d_1525_v41_: _dafny.MultiSet
            d_1525_v41_ = _dafny.MultiSet([p3, len(d_1485_v14_)])
            d_1526_v42_: C1
            nw222_ = C1()
            nw222_.ctor__(d_1473_v4_, len((d_1524_v40_) | (d_1524_v40_)), (d_1525_v41_).cardinality, len(default__.fm43((d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))], p2, globalState)), d_1517_cf19_, (_dafny.SeqWithoutIsStrInference([d_1503_v31_ for d_1527_i2_ in range(default__.abs(-416))]) if (d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))] else default__.fm31(d_1503_v31_, d_1505_v33_, globalState)))
            d_1526_v42_ = nw222_
            d_1528_v43_: D10
            d_1528_v43_ = D10_DC31(p1, d_1503_v31_, p1, (self).f14)
            d_1529_v44_: _dafny.Array
            nw223_ = _dafny.Array(None, 21)
            nw223_[int(0)] = (_dafny.MultiSet([d_1503_v31_])).intersection(_dafny.MultiSet([d_1503_v31_]))
            nw223_[int(1)] = (_dafny.MultiSet([d_1503_v31_, (d_1528_v43_).cf53])) | (d_1504_v32_)
            nw223_[int(2)] = _dafny.MultiSet([d_1503_v31_])
            nw223_[int(3)] = ((d_1504_v32_).set(d_1503_v31_, default__.abs(d_1516_cf20_))) | (d_1504_v32_)
            nw223_[int(4)] = d_1504_v32_
            nw223_[int(5)] = d_1504_v32_
            nw223_[int(6)] = _dafny.MultiSet([_dafny.CodePoint('x')])
            nw223_[int(7)] = _dafny.MultiSet([d_1503_v31_, d_1503_v31_, d_1503_v31_])
            nw223_[int(8)] = (d_1504_v32_) | (d_1504_v32_)
            nw223_[int(9)] = d_1504_v32_
            nw223_[int(10)] = d_1504_v32_
            nw223_[int(11)] = d_1504_v32_
            nw223_[int(12)] = (_dafny.MultiSet([d_1503_v31_])) | (d_1504_v32_)
            nw223_[int(13)] = d_1504_v32_
            nw223_[int(14)] = d_1504_v32_
            nw223_[int(15)] = (d_1504_v32_ if (d_1469_v0_).fm17(d_1472_v3_, globalState) else (d_1504_v32_).set(d_1503_v31_, default__.abs(d_1516_cf20_)))
            nw223_[int(16)] = d_1504_v32_
            nw223_[int(17)] = (d_1504_v32_) - (_dafny.MultiSet(p0))
            nw223_[int(18)] = d_1504_v32_
            nw223_[int(19)] = (d_1504_v32_).intersection(d_1504_v32_)
            nw223_[int(20)] = d_1504_v32_
            d_1529_v44_ = nw223_
            index246_ = default__.safeIndex(737, (d_1529_v44_).length(0))
            (d_1529_v44_)[index246_] = (d_1504_v32_) - (d_1504_v32_)
            index247_ = default__.safeIndex(465, (d_1473_v4_).length(0))
            index248_ = default__.safeIndex(737, (d_1529_v44_).length(0))
            rhs238_ = True
            rhs239_ = (d_1504_v32_ if (self).f25 else d_1504_v32_)
            lhs179_ = d_1473_v4_
            lhs180_ = default__.safeIndex(465, (d_1473_v4_).length(0))
            lhs181_ = d_1529_v44_
            lhs182_ = default__.safeIndex(737, (d_1529_v44_).length(0))
            lhs179_[lhs180_] = rhs238_
            lhs181_[lhs182_] = rhs239_
        elif source26_.is_DC11:
            d_1530___mcc_h5_ = source26_.cf24
            d_1531___mcc_h6_ = source26_.cf25
            d_1532___mcc_h7_ = source26_.cf26
            d_1533___mcc_h8_ = source26_.cf27
            d_1534_cf27_ = d_1533___mcc_h8_
            d_1535_cf26_ = d_1532___mcc_h7_
            d_1536_cf25_ = d_1531___mcc_h6_
            d_1537_cf24_ = d_1530___mcc_h5_
            index249_ = default__.safeIndex(166, (d_1477_v8_).length(0))
            (d_1477_v8_)[index249_] = 761
            index250_ = default__.safeIndex(166, (d_1477_v8_).length(0))
            rhs240_ = (default__.fm0(p2, globalState)) - (-949)
            rhs241_ = ((self).f11) in (_dafny.Map({p1: p2}))
            rhs242_ = (self).f11
            rhs243_ = (d_1472_v3_)[default__.safeIndex(p3, len(d_1472_v3_))]
            lhs183_ = globalState
            lhs184_ = globalState
            lhs185_ = globalState
            lhs186_ = d_1477_v8_
            lhs187_ = default__.safeIndex(166, (d_1477_v8_).length(0))
            lhs183_.f6 = rhs240_
            lhs184_.f7 = rhs241_
            lhs185_.f6 = rhs242_
            lhs186_[lhs187_] = rhs243_
            r0 = (d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))]
            if (((self).f13) == ((self).f26)) and (d_1534_cf27_):
                (globalState).f6 = (0) - (default__.fm0(not((self).f15), globalState))
                index251_ = default__.safeIndex(166, (d_1477_v8_).length(0))
                rhs244_ = not(not((default__.fm0(p2, globalState)) < ((-174) - (len(d_1483_v12_)))))
                rhs245_ = len(((default__.fm31(d_1503_v31_, d_1505_v33_, globalState)) + (p0)) + ((self).f26))
                rhs246_ = (self).f12
                rhs247_ = ((d_1477_v8_)[default__.safeIndex(166, (d_1477_v8_).length(0))]) + ((p1) - ((self).f14))
                rhs248_ = True
                lhs188_ = globalState
                lhs189_ = globalState
                lhs190_ = d_1477_v8_
                lhs191_ = default__.safeIndex(166, (d_1477_v8_).length(0))
                lhs192_ = globalState
                lhs188_.f2 = rhs244_
                r1 = rhs245_
                lhs189_.f6 = rhs246_
                lhs190_[lhs191_] = rhs247_
                lhs192_.f7 = rhs248_
                d_1538_v45_: _dafny.MultiSet
                d_1538_v45_ = _dafny.MultiSet([(self).f14])
                (globalState).f6 = ((d_1538_v45_)[(d_1474_v5_).cardinality] if ((d_1474_v5_).cardinality) in (d_1538_v45_) else len(_dafny.Set({len(d_1472_v3_), d_1537_cf24_, p3})))
                r0 = not(d_1534_cf27_)
                d_1539_v46_: _dafny.Map
                d_1539_v46_ = _dafny.Map({p0: len(_dafny.SeqWithoutIsStrInference([(self).f15]))})
                d_1472_v3_ = (_dafny.SeqWithoutIsStrInference([(self).f12, len(d_1472_v3_), (0) - (len(d_1539_v46_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nl"))), (d_1477_v8_)[default__.safeIndex(166, (d_1477_v8_).length(0))]])) + (d_1472_v3_)
            elif True:
                d_1540_v47_: _dafny.Array
                nw224_ = _dafny.Array(None, 15)
                d_1540_v47_ = nw224_
                d_1541_v48_: _dafny.Map
                d_1541_v48_ = _dafny.Map({p2: d_1540_v47_})
                d_1542_v49_: _dafny.Set
                d_1542_v49_ = _dafny.Set({((d_1541_v48_)[p2] if (p2) in (d_1541_v48_) else d_1540_v47_)})
                d_1543_v50_: D1
                d_1543_v50_ = D1_DC2(d_1503_v31_)
                d_1544_v51_: D2
                d_1544_v51_ = D2_DC7(d_1473_v4_, d_1542_v49_, d_1543_v50_, (self).f12, (self).f14)
                d_1545_v52_: C1
                nw225_ = C1()
                nw225_.ctor__((d_1544_v51_).cf10, 210, d_1537_cf24_, (d_1537_cf24_ if (self).f15 else len(_dafny.Map({d_1537_cf24_: (self).f25}))), not ((self).f15) or ((self).f15), self.f16)
                d_1545_v52_ = nw225_
                d_1546_v53_: _dafny.Map
                d_1546_v53_ = _dafny.Map({(self).f12: d_1475_v6_})
                d_1547_v54_: D10
                d_1547_v54_ = D10_DC31(p1, d_1503_v31_, 566, default__.safeModuloInt(p3, (0) - (len(d_1546_v53_))))
                d_1547_v54_ = d_1547_v54_
                d_1548_v55_: _dafny.Array
                nw226_ = _dafny.Array(None, 1)
                nw226_[int(0)] = (_dafny.MultiSet([(d_1477_v8_)[default__.safeIndex(166, (d_1477_v8_).length(0))]])).cardinality
                d_1548_v55_ = nw226_
                d_1549_v56_: D9
                d_1549_v56_ = D9_DC27(d_1548_v55_)
                d_1550_v57_: _dafny.Map
                d_1550_v57_ = _dafny.Map({(d_1549_v56_).cf46: d_1474_v5_})
                d_1474_v5_ = ((((d_1550_v57_)[d_1548_v55_] if (d_1548_v55_) in (d_1550_v57_) else d_1474_v5_) if (d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))] else default__.fm54((d_1472_v3_)[default__.safeIndex(p3, len(d_1472_v3_))], globalState))).set(d_1534_cf27_, default__.abs(((self).f14) - (d_1537_cf24_)))
                r0 = False
                d_1551_v58_: T0
                nw227_ = C1()
                nw227_.ctor__(d_1545_v52_.f24, len(d_1483_v12_), p1, ((self).f11) + (p3), (self).f25, self.f16)
                d_1551_v58_ = nw227_
                d_1552_v59_: D16
                d_1552_v59_ = D16_DC44(d_1551_v58_)
                rhs249_ = (d_1544_v51_).cf14
                rhs250_ = ((d_1552_v59_).cf80 if not(not(p2)) else (d_1551_v58_ if True else d_1551_v58_))
                rhs251_ = d_1471_v2_
                d_1537_cf24_ = rhs249_
                d_1551_v58_ = rhs250_
                d_1471_v2_ = rhs251_
            (globalState).f6 = (0) - (-984)
        elif source26_.is_DC9:
            d_1553___mcc_h9_ = source26_.cf18
            d_1554_cf18_ = d_1553___mcc_h9_
            d_1555_v60_: _dafny.Seq
            d_1555_v60_ = _dafny.SeqWithoutIsStrInference([p0])
            pat_let_tv22_ = d_1476_v7_
            pat_let_tv23_ = d_1476_v7_
            pat_let_tv24_ = d_1555_v60_
            pat_let_tv25_ = p2
            pat_let_tv26_ = globalState
            def iife82_(_pat_let20_0):
                def iife83_(d_1556_dt__update__tmp_h1_):
                    def iife84_(_pat_let21_0):
                        def iife85_(d_1557_dt__update_hcf15_h0_):
                            return D2_DC8(d_1557_dt__update_hcf15_h0_, (d_1556_dt__update__tmp_h1_).cf16, (d_1556_dt__update__tmp_h1_).cf17)
                        return iife85_(_pat_let21_0)
                    return iife84_((self).fm1(((pat_let_tv22_)[(self).f12] if ((self).f12) in (pat_let_tv23_) else len(pat_let_tv24_)), (self).f25, pat_let_tv25_, pat_let_tv26_))
                return iife83_(_pat_let20_0)
            d_1502_v30_ = iife82_(d_1502_v30_)
            d_1558_v61_: D1
            d_1558_v61_ = D1_DC3((d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))])
            d_1559_v62_: _dafny.Array
            def lambda53_(d_1560_v31_):
                def lambda54_(d_1561_i3_):
                    return d_1560_v31_

                return lambda54_

            init30_ = lambda53_(d_1503_v31_)
            nw228_ = _dafny.Array(None, 10)
            for i0_30_ in range(nw228_.length(0)):
                nw228_[i0_30_] = init30_(i0_30_)
            d_1559_v62_ = nw228_
            d_1562_v63_: _dafny.Map
            d_1562_v63_ = _dafny.Map({(d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))]: d_1559_v62_})
            d_1563_v64_: D16
            d_1563_v64_ = D16_DC45(d_1473_v4_, d_1503_v31_, -868, len(d_1562_v63_), (self).f15)
            d_1564_v65_: bool
            d_1565_v66_: _dafny.Seq
            out54_: bool
            out55_: _dafny.Seq
            out54_, out55_ = default__.m0((d_1558_v61_).cf4, not((d_1563_v64_).cf85), (self).f12, default__.safeModuloInt((self).f12, len(_dafny.Map({(d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))]: d_1503_v31_}))), globalState)
            d_1564_v65_ = out54_
            d_1565_v66_ = out55_
            r2 = p3
            (self).f16 = self.f16
        elif True:
            d_1566___mcc_h10_ = source26_.cf28
            d_1567_cf28_ = d_1566___mcc_h10_
            d_1568_v67_: D1
            d_1568_v67_ = D1_DC5((self).f14, (d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))], (d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))])
            if (d_1568_v67_).cf8:
                d_1569_v68_: _dafny.MultiSet
                d_1569_v68_ = _dafny.MultiSet([(self).f11, (self).f12])
                d_1570_v69_: C1
                nw229_ = C1()
                nw229_.ctor__(d_1473_v4_, 176, p1, (p3) * ((self).f14), (d_1569_v68_).isdisjoint(d_1569_v68_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcmr")))
                d_1570_v69_ = nw229_
                d_1571_v70_: _dafny.Map
                d_1571_v70_ = _dafny.Map({(d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))]: d_1477_v8_})
                d_1571_v70_ = d_1571_v70_
                (globalState).f2 = (self).f25
                d_1572_v71_: C0
                nw230_ = C0()
                nw230_.ctor__()
                d_1572_v71_ = nw230_
                d_1573_v72_: D8
                d_1573_v72_ = D8_DC24(d_1474_v5_)
                index252_ = default__.safeIndex(465, (d_1473_v4_).length(0))
                rhs252_ = d_1572_v71_
                rhs253_ = ((d_1573_v72_).cf43).issubset(d_1474_v5_)
                lhs193_ = d_1473_v4_
                lhs194_ = default__.safeIndex(465, (d_1473_v4_).length(0))
                d_1572_v71_ = rhs252_
                lhs193_[lhs194_] = rhs253_
                d_1574_v73_: _dafny.Set
                d_1574_v73_ = _dafny.Set({(self).f12})
                index253_ = default__.safeIndex(488, (d_1477_v8_).length(0))
                (d_1477_v8_)[index253_] = (0) - (len((d_1574_v73_) | (d_1574_v73_)))
                index254_ = default__.safeIndex(488, (d_1477_v8_).length(0))
                index255_ = default__.safeIndex(465, (d_1473_v4_).length(0))
                rhs254_ = ((self).f14) * (174)
                rhs255_ = _dafny.CodePoint('e')
                rhs256_ = ((34) < (len(d_1574_v73_))) or ((self).f15)
                rhs257_ = (d_1570_v69_).fm1((self).f12, (d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))], True, globalState)
                rhs258_ = (self).fm26((d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))], globalState)
                lhs195_ = d_1477_v8_
                lhs196_ = default__.safeIndex(488, (d_1477_v8_).length(0))
                lhs197_ = d_1473_v4_
                lhs198_ = default__.safeIndex(465, (d_1473_v4_).length(0))
                lhs199_ = globalState
                lhs200_ = globalState
                lhs195_[lhs196_] = rhs254_
                d_1503_v31_ = rhs255_
                lhs197_[lhs198_] = rhs256_
                lhs199_.f7 = rhs257_
                lhs200_.f7 = rhs258_
            elif True:
                d_1575_v74_: _dafny.Array
                nw231_ = _dafny.Array(None, 7)
                d_1575_v74_ = nw231_
                d_1576_v75_: _dafny.Set
                d_1576_v75_ = _dafny.Set({d_1575_v74_, d_1575_v74_})
                d_1577_v76_: D1
                d_1577_v76_ = D1_DC2(d_1503_v31_)
                d_1578_v77_: D2
                d_1578_v77_ = D2_DC7(d_1473_v4_, d_1576_v75_, d_1577_v76_, p3, (self).f11)
                d_1579_v78_: _dafny.Seq
                d_1579_v78_ = _dafny.SeqWithoutIsStrInference([d_1473_v4_, d_1473_v4_])
                d_1578_v77_ = D2_DC7(d_1473_v4_, d_1576_v75_, D1_DC2(d_1503_v31_), (len(d_1471_v2_)) + (len(d_1476_v7_)), len((d_1579_v78_) + (d_1579_v78_)))
                d_1580_v79_: _dafny.Array
                nw232_ = _dafny.Array(_dafny.Seq({}), 20)
                d_1580_v79_ = nw232_
                index256_ = default__.safeIndex(626, (d_1580_v79_).length(0))
                (d_1580_v79_)[index256_] = d_1472_v3_
                index257_ = default__.safeIndex(626, (d_1580_v79_).length(0))
                (d_1580_v79_)[index257_] = d_1472_v3_
                d_1581_v80_: _dafny.Map
                d_1581_v80_ = _dafny.Map({False: d_1475_v6_})
                index258_ = default__.safeIndex(632, (d_1477_v8_).length(0))
                (d_1477_v8_)[index258_] = (p1) - (len(d_1581_v80_))
                index259_ = default__.safeIndex(465, (d_1473_v4_).length(0))
                index260_ = default__.safeIndex(632, (d_1477_v8_).length(0))
                def iife86_():
                    coll42_ = _dafny.Map()
                    compr_42_: int
                    for compr_42_ in _dafny.IntegerRange(22, 699):
                        d_1582_v81_: int = compr_42_
                        if ((22) <= (d_1582_v81_)) and ((d_1582_v81_) < (699)):
                            coll42_[(d_1582_v81_) * (len(self.f16))] = p3
                    return _dafny.Map(coll42_)
                rhs259_ = p2
                rhs260_ = ((self).f14) + (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxpno"))).set(default__.safeIndex(len(iife86_()
                ), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxpno")))), d_1503_v31_)))
                lhs201_ = d_1473_v4_
                lhs202_ = default__.safeIndex(465, (d_1473_v4_).length(0))
                lhs203_ = d_1477_v8_
                lhs204_ = default__.safeIndex(632, (d_1477_v8_).length(0))
                lhs201_[lhs202_] = rhs259_
                lhs203_[lhs204_] = rhs260_
                (globalState).f6 = (d_1474_v5_).cardinality
                d_1583_v82_: D11
                d_1583_v82_ = D11_DC34(not((d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))]), p2, (self).f12, p2, not((self).f25))
                d_1584_v83_: _dafny.Map
                d_1584_v83_ = _dafny.Map({p2: d_1583_v82_})
                d_1584_v83_ = (d_1584_v83_).set((d_1473_v4_)[default__.safeIndex(465, (d_1473_v4_).length(0))], D11_DC34((self).f15, p2, (self).f14, (self).f15, (self).f15))
            d_1585_v84_: D2
            d_1585_v84_ = D2_DC6(p0)
            (self).f16 = (d_1585_v84_).cf9
            d_1586_v85_: _dafny.Seq
            d_1586_v85_ = _dafny.SeqWithoutIsStrInference([d_1473_v4_])
            d_1587_v86_: _dafny.Map
            d_1587_v86_ = _dafny.Map({(d_1586_v85_) + (d_1586_v85_): p2})
            d_1587_v86_ = (d_1587_v86_).set(d_1586_v85_, (d_1472_v3_) == (d_1472_v3_))
            d_1588_v87_: _dafny.Set
            d_1588_v87_ = _dafny.Set({p3, (self).f11})
            def iife87_():
                coll43_ = _dafny.Set()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(-145, 290):
                    d_1589_v88_: int = compr_43_
                    if ((-145) <= (d_1589_v88_)) and ((d_1589_v88_) < (290)):
                        coll43_ = coll43_.union(_dafny.Set([(d_1589_v88_) - ((self).f11)]))
                return _dafny.Set(coll43_)
            (globalState).f7 = (d_1588_v87_).issubset((iife87_()
            ).intersection(d_1588_v87_))
        d_1590_v89_: _dafny.MultiSet
        d_1590_v89_ = _dafny.MultiSet([(self.f16).set(default__.safeIndex(271, len(self.f16)), d_1503_v31_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnoekch")), (self).f26, (((self).f26) + ((self).f13)).set(default__.safeIndex(len(d_1472_v3_), len(((self).f26) + ((self).f13))), d_1503_v31_)])
        d_1590_v89_ = d_1590_v89_
        r0 = (self).f25
        r1 = (0) - ((self).f14)
        r2 = ((d_1504_v32_)[_dafny.CodePoint('x')] if (_dafny.CodePoint('x')) in (d_1504_v32_) else default__.fm0(p2, globalState))
        return r0, r1, r2

    def m17(self, p0, p1, p2, globalState):
        d_1591_v0_: _dafny.Array
        def lambda55_(d_1592_i1_):
            return (d_1592_i1_) + ((self).f12)

        init31_ = lambda55_
        nw233_ = _dafny.Array(None, 5)
        for i0_31_ in range(nw233_.length(0)):
            nw233_[i0_31_] = init31_(i0_31_)
        d_1591_v0_ = nw233_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1591_v0_).length(0)):
            d_1593_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1593_i0_)) and ((d_1593_i0_) < ((d_1591_v0_).length(0)))):
                (d_1591_v0_)[(d_1593_i0_)] = default__.safeModuloInt(d_1593_i0_, (self).f12)
        d_1594_v1_: _dafny.Array
        nw234_ = _dafny.Array(_dafny.Set({}), 17)
        d_1594_v1_ = nw234_
        d_1595_v2_: _dafny.Seq
        d_1595_v2_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f11])
        d_1596_v3_: _dafny.Set
        d_1596_v3_ = _dafny.Set({(self).f11, (self).f11, len(d_1595_v2_)})
        index261_ = default__.safeIndex(231, (d_1594_v1_).length(0))
        (d_1594_v1_)[index261_] = d_1596_v3_
        d_1597_v4_: _dafny.Seq
        d_1597_v4_ = _dafny.SeqWithoutIsStrInference([(self).f25])
        d_1598_v5_: D14
        d_1598_v5_ = D14_DC40(d_1597_v4_)
        pat_let_tv27_ = d_1596_v3_
        pat_let_tv28_ = d_1595_v2_
        pat_let_tv29_ = d_1595_v2_
        index262_ = default__.safeIndex(231, (d_1594_v1_).length(0))
        def lambda56_(source27_):
            if source27_.is_DC41:
                d_1599___mcc_h0_ = source27_.cf73
                d_1600___mcc_h1_ = source27_.cf74
                d_1601___mcc_h2_ = source27_.cf75
                d_1602_cf75_ = d_1601___mcc_h2_
                d_1603_cf74_ = d_1600___mcc_h1_
                d_1604_cf73_ = d_1599___mcc_h0_
                return (_dafny.Set({(_dafny.MultiSet([(self).f14])).cardinality})).intersection(pat_let_tv27_)
            elif True:
                d_1605___mcc_h3_ = source27_.cf72
                d_1606_cf72_ = d_1605___mcc_h3_
                return _dafny.Set({(pat_let_tv28_)[default__.safeIndex(len((self).f13), len(pat_let_tv29_))]})

        (d_1594_v1_)[index262_] = lambda56_(d_1598_v5_)
        (globalState).f7 = (self).fm1(len(_dafny.SeqWithoutIsStrInference([d_1595_v2_, d_1595_v2_])), (self).fm26((self).f15, globalState), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1607_i2_ in range(default__.abs(221))])) <= (self.f16), globalState)
        if not ((self).f25) or (((self).f15 if (self).f15 else (self).f15)):
            (globalState).f6 = ((0) - ((self).f12) if (self).f25 else 107)
            if (((self).f12 if (self).f15 else (self).f14)) >= (len(self.f16)):
                index263_ = default__.safeIndex(231, (d_1594_v1_).length(0))
                (d_1594_v1_)[index263_] = (d_1594_v1_)[default__.safeIndex(231, (d_1594_v1_).length(0))]
                (globalState).f6 = (default__.safeDivisionInt((self).f12, (self).f11)) + ((self).f11)
                index264_ = default__.safeIndex(309, (p0).length(0))
                (p0)[index264_] = ((self).f12) <= ((self).f12)
                index265_ = default__.safeIndex(691, (p1).length(0))
                (p1)[index265_] = (self).f11
                d_1608_v6_: _dafny.Set
                d_1608_v6_ = _dafny.Set({(self).f15})
                d_1609_v7_: _dafny.MultiSet
                d_1609_v7_ = _dafny.MultiSet([d_1608_v6_])
                d_1610_v8_: str
                d_1610_v8_ = _dafny.CodePoint('m')
                d_1611_v9_: _dafny.Map
                d_1611_v9_ = _dafny.Map({d_1610_v8_: (self).f15})
                d_1612_v10_: _dafny.MultiSet
                d_1612_v10_ = _dafny.MultiSet([d_1611_v9_])
                index266_ = default__.safeIndex(309, (p0).length(0))
                index267_ = default__.safeIndex(691, (p1).length(0))
                rhs261_ = ((d_1612_v10_ if (self).f25 else d_1612_v10_)) != (d_1612_v10_)
                rhs262_ = (self).f11
                rhs263_ = _dafny.MultiSet(default__.fm55(globalState))
                rhs264_ = not((self).f25)
                lhs205_ = p0
                lhs206_ = default__.safeIndex(309, (p0).length(0))
                lhs207_ = p1
                lhs208_ = default__.safeIndex(691, (p1).length(0))
                lhs209_ = globalState
                lhs205_[lhs206_] = rhs261_
                lhs207_[lhs208_] = rhs262_
                d_1609_v7_ = rhs263_
                lhs209_.f7 = rhs264_
                (globalState).f6 = default__.safeModuloInt((self).f12, (self).f14)
                index268_ = default__.safeIndex(691, (p1).length(0))
                (p1)[index268_] = (p1)[default__.safeIndex(691, (p1).length(0))]
            elif True:
                rhs265_ = (self).f25
                rhs266_ = (self).f11
                lhs210_ = globalState
                lhs211_ = globalState
                lhs210_.f2 = rhs265_
                lhs211_.f6 = rhs266_
                d_1613_v11_: _dafny.Array
                nw235_ = _dafny.Array(False, 14)
                d_1613_v11_ = nw235_
                rhs267_ = (self).f11
                rhs268_ = not(not(not (((self).f25) and (False)) or (((self).f15 if (self).f25 else (self).f25))))
                rhs269_ = (self).f12
                rhs270_ = p0
                lhs212_ = globalState
                lhs213_ = globalState
                lhs214_ = globalState
                lhs212_.f6 = rhs267_
                lhs213_.f7 = rhs268_
                lhs214_.f6 = rhs269_
                d_1613_v11_ = rhs270_
                d_1614_v12_: str
                d_1614_v12_ = _dafny.CodePoint('x')
                d_1615_v13_: D0
                d_1615_v13_ = D0_DC1((d_1614_v12_) not in (self.f16), (self).f12)
                d_1615_v13_ = D0_DC1((self).f25, (0) - (default__.safeModuloInt((self).f11, (self).f12)))
                d_1591_v0_ = d_1591_v0_
                index269_ = default__.safeIndex(458, (p0).length(0))
                (p0)[index269_] = not ((self).f25) or (True)
                index270_ = default__.safeIndex(458, (p0).length(0))
                (p0)[index270_] = (self).f15
            (globalState).f2 = (self).f25
            (globalState).f6 = len((self).f13)
            if (self).f25:
                (globalState).f2 = (self).f15
                d_1616_v14_: _dafny.Array
                nw236_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_1616_v14_ = nw236_
                index271_ = default__.safeIndex(152, (d_1616_v14_).length(0))
                (d_1616_v14_)[index271_] = p1
                index272_ = default__.safeIndex(152, (d_1616_v14_).length(0))
                (d_1616_v14_)[index272_] = d_1591_v0_
                index273_ = default__.safeIndex(626, (p0).length(0))
                (p0)[index273_] = False
                index274_ = default__.safeIndex(626, (p0).length(0))
                (p0)[index274_] = (self).f25
                d_1617_v15_: _dafny.Array
                nw237_ = _dafny.Array(_dafny.Seq({}), 17)
                d_1617_v15_ = nw237_
                d_1617_v15_ = d_1617_v15_
                d_1618_v16_: _dafny.Array
                def lambda57_(d_1619_i3_):
                    return (self).f15

                init32_ = lambda57_
                nw238_ = _dafny.Array(None, 22)
                for i0_32_ in range(nw238_.length(0)):
                    nw238_[i0_32_] = init32_(i0_32_)
                d_1618_v16_ = nw238_
                d_1620_v17_: str
                d_1620_v17_ = _dafny.CodePoint('h')
                d_1621_v18_: C1
                nw239_ = C1()
                nw239_.ctor__(d_1618_v16_, default__.fm0((self).f15, globalState), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))).set(default__.safeIndex((self).f12, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))), d_1620_v17_)), ((self).f14) - ((self).f12), not(not(not((p0)[default__.safeIndex(626, (p0).length(0))]))), (self).f13)
                d_1621_v18_ = nw239_
                d_1621_v18_ = d_1621_v18_
            elif True:
                d_1622_v19_: _dafny.MultiSet
                d_1622_v19_ = _dafny.MultiSet([(self).f25])
                d_1623_v20_: str
                d_1623_v20_ = _dafny.CodePoint('j')
                d_1624_v21_: _dafny.Array
                nw240_ = _dafny.Array(None, 17)
                d_1624_v21_ = nw240_
                d_1625_v22_: _dafny.Set
                d_1625_v22_ = _dafny.Set({d_1624_v21_})
                d_1626_v23_: _dafny.Map
                d_1626_v23_ = _dafny.Map({(self).f25: d_1625_v22_})
                d_1627_v24_: D1
                d_1627_v24_ = D1_DC2(d_1623_v20_)
                d_1628_v25_: D2
                d_1628_v25_ = D2_DC7(p0, ((d_1626_v23_)[(d_1597_v4_)[default__.safeIndex((self).f12, len(d_1597_v4_))]] if ((d_1597_v4_)[default__.safeIndex((self).f12, len(d_1597_v4_))]) in (d_1626_v23_) else d_1625_v22_), d_1627_v24_, (self).f12, (self).f14)
                d_1629_v26_: C2
                nw241_ = C2()
                nw241_.ctor__(((d_1622_v19_) - (d_1622_v19_)).cardinality, ((self).f26) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1630_i4_ in range(default__.abs(378))])), (self).f14, len(_dafny.Map({_dafny.MultiSet([d_1623_v20_]): _dafny.Set({(d_1628_v25_).cf14})})))
                d_1629_v26_ = nw241_
                index275_ = default__.safeIndex(621, (p1).length(0))
                (p1)[index275_] = (d_1629_v26_.f27) - ((self).f12)
                index276_ = default__.safeIndex(206, (p0).length(0))
                (p0)[index276_] = (self).f15
                d_1631_v27_: _dafny.Set
                d_1631_v27_ = _dafny.Set({(self).f25, (self).f25})
                d_1632_v28_: _dafny.Map
                d_1632_v28_ = _dafny.Map({d_1623_v20_: p0})
                index277_ = default__.safeIndex(621, (p1).length(0))
                index278_ = default__.safeIndex(206, (p0).length(0))
                rhs271_ = 97
                rhs272_ = default__.safeModuloInt((len(d_1595_v2_) if not(not(True)) else 701), (self).f12)
                rhs273_ = (266) != (len((d_1631_v27_).intersection(d_1631_v27_)))
                rhs274_ = (self).f15
                rhs275_ = len(d_1632_v28_)
                lhs215_ = p1
                lhs216_ = default__.safeIndex(621, (p1).length(0))
                lhs217_ = globalState
                lhs218_ = p0
                lhs219_ = default__.safeIndex(206, (p0).length(0))
                lhs220_ = globalState
                lhs221_ = d_1629_v26_
                lhs215_[lhs216_] = rhs271_
                lhs217_.f6 = rhs272_
                lhs218_[lhs219_] = rhs273_
                lhs220_.f2 = rhs274_
                lhs221_.f27 = rhs275_
                index279_ = default__.safeIndex(206, (p0).length(0))
                (p0)[index279_] = not((p0)[default__.safeIndex(206, (p0).length(0))])
                d_1633_v29_: _dafny.Array
                nw242_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_1633_v29_ = nw242_
                d_1634_v30_: _dafny.Array
                nw243_ = _dafny.Array(False, 24)
                d_1634_v30_ = nw243_
                index280_ = default__.safeIndex(315, (d_1633_v29_).length(0))
                (d_1633_v29_)[index280_] = d_1634_v30_
                index281_ = default__.safeIndex(315, (d_1633_v29_).length(0))
                (d_1633_v29_)[index281_] = d_1634_v30_
                d_1635_v31_: _dafny.Map
                d_1635_v31_ = _dafny.Map({len(d_1597_v4_): p2})
                d_1636_v32_: _dafny.Map
                d_1636_v32_ = _dafny.Map({(d_1591_v0_ if (p0)[default__.safeIndex(206, (p0).length(0))] else d_1591_v0_): ((d_1635_v31_)[-658] if (-658) in (d_1635_v31_) else p2)})
                d_1636_v32_ = (d_1636_v32_).set(p2, d_1591_v0_)
        elif True:
            (globalState).f6 = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcgqgvidi")))) - ((self).f14)
            index282_ = default__.safeIndex(736, (p0).length(0))
            (p0)[index282_] = (self).f15
            index283_ = default__.safeIndex(736, (p0).length(0))
            (p0)[index283_] = (self).f25
            d_1637_v33_: _dafny.Array
            def lambda58_(d_1638_v4_):
                def lambda59_(d_1639_i5_):
                    return (_dafny.MultiSet([(d_1638_v4_)[default__.safeIndex((0) - ((self).f11), len(d_1638_v4_))]])).isdisjoint(_dafny.MultiSet([(self).f25, False]))

                return lambda59_

            init33_ = lambda58_(d_1597_v4_)
            nw244_ = _dafny.Array(None, 16)
            for i0_33_ in range(nw244_.length(0)):
                nw244_[i0_33_] = init33_(i0_33_)
            d_1637_v33_ = nw244_
            d_1637_v33_ = d_1637_v33_
            d_1640_v34_: _dafny.Map
            d_1640_v34_ = _dafny.Map({(p0)[default__.safeIndex(736, (p0).length(0))]: (self).f26})
            d_1641_v35_: C2
            nw245_ = C2()
            nw245_.ctor__((self).f12, ((d_1640_v34_)[(p0)[default__.safeIndex(736, (p0).length(0))]] if ((p0)[default__.safeIndex(736, (p0).length(0))]) in (d_1640_v34_) else (self).f26), (self).f14, (self).f12)
            d_1641_v35_ = nw245_
            (globalState).f7 = (self).f15
        d_1642_v36_: _dafny.Array
        nw246_ = _dafny.Array(False, 19)
        d_1642_v36_ = nw246_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1642_v36_).length(0)):
            d_1643_i6_: int = guard_loop_3_
            if (True) and (((0) <= (d_1643_i6_)) and ((d_1643_i6_) < ((d_1642_v36_).length(0)))):
                (d_1642_v36_)[(d_1643_i6_)] = True
        hi6_ = default__.fm0((self).f15, globalState)
        for d_1644_i7_ in range((self).f14, hi6_):
            (globalState).f6 = (self).f12
            if (self).f15:
                d_1645_v37_: _dafny.Array
                def lambda60_(d_1646_i8_):
                    return _dafny.SeqWithoutIsStrInference([not(True)])

                init34_ = lambda60_
                nw247_ = _dafny.Array(None, 1)
                for i0_34_ in range(nw247_.length(0)):
                    nw247_[i0_34_] = init34_(i0_34_)
                d_1645_v37_ = nw247_
                d_1647_v38_: _dafny.Map
                d_1647_v38_ = _dafny.Map({(self).f15: d_1597_v4_})
                d_1648_v39_: _dafny.Map
                d_1648_v39_ = _dafny.Map({False: (self).f25})
                index284_ = default__.safeIndex(632, (d_1645_v37_).length(0))
                (d_1645_v37_)[index284_] = (d_1597_v4_) + (((d_1647_v38_)[((d_1648_v39_)[(self).f15] if ((self).f15) in (d_1648_v39_) else (self).f25)] if (((d_1648_v39_)[(self).f15] if ((self).f15) in (d_1648_v39_) else (self).f25)) in (d_1647_v38_) else d_1597_v4_))
                index285_ = default__.safeIndex(632, (d_1645_v37_).length(0))
                (d_1645_v37_)[index285_] = ((d_1597_v4_ if not(not((self).f15)) else _dafny.SeqWithoutIsStrInference([(self).f25]))) + ((d_1597_v4_) + (d_1597_v4_))
                rhs276_ = (self).f14
                rhs277_ = not((self).f25)
                rhs278_ = (self).f15
                lhs222_ = globalState
                lhs223_ = globalState
                lhs224_ = globalState
                lhs222_.f6 = rhs276_
                lhs223_.f7 = rhs277_
                lhs224_.f7 = rhs278_
                (globalState).f7 = (self).f15
                d_1649_v40_: _dafny.Set
                d_1649_v40_ = _dafny.Set({(self).f15})
                d_1650_v41_: _dafny.MultiSet
                d_1650_v41_ = _dafny.MultiSet([d_1649_v40_, d_1649_v40_])
                d_1651_v42_: _dafny.MultiSet
                d_1651_v42_ = _dafny.MultiSet([(d_1650_v41_).cardinality, d_1644_i7_, (self).f12, (self).f11])
                d_1651_v42_ = (d_1651_v42_).intersection(((d_1651_v42_).set((self).f11, default__.abs(139))) | (d_1651_v42_))
                d_1652_v43_: str
                d_1652_v43_ = _dafny.CodePoint('m')
                d_1653_v44_: _dafny.MultiSet
                d_1653_v44_ = _dafny.MultiSet([d_1652_v43_, _dafny.CodePoint('y')])
                d_1654_v45_: D10
                d_1654_v45_ = D10_DC30((self).f25, not((self).f25), (self.f16).set(default__.safeIndex((d_1653_v44_).cardinality, len(self.f16)), d_1652_v43_), (d_1652_v43_) not in ((self).f13))
                rhs279_ = (self).f12
                rhs280_ = (self).f15
                rhs281_ = D10_DC30((self).f15, ((self).f25) and ((self).f25), self.f16, ((self).f12) != (d_1644_i7_))
                lhs225_ = globalState
                lhs226_ = globalState
                lhs225_.f6 = rhs279_
                lhs226_.f2 = rhs280_
                d_1654_v45_ = rhs281_
            elif True:
                d_1655_v46_: _dafny.Map
                d_1655_v46_ = _dafny.Map({(-902) - ((_dafny.MultiSet([(self).f15])).cardinality): (self).f15})
                d_1655_v46_ = (d_1655_v46_).set((self).f12, (self).f15)
                index286_ = default__.safeIndex(923, (d_1642_v36_).length(0))
                (d_1642_v36_)[index286_] = (self).f25
                d_1656_v47_: _dafny.Set
                d_1656_v47_ = _dafny.Set({not((self).fm1(d_1644_i7_, (self).f25, (self).f25, globalState))})
                d_1657_v48_: _dafny.Seq
                d_1657_v48_ = _dafny.SeqWithoutIsStrInference([d_1656_v47_, d_1656_v47_, d_1656_v47_])
                index287_ = default__.safeIndex(923, (d_1642_v36_).length(0))
                (d_1642_v36_)[index287_] = (d_1597_v4_)[default__.safeIndex(default__.safeModuloInt(len((d_1657_v48_)[default__.safeIndex(d_1644_i7_, len(d_1657_v48_))]), (self).f14), len(d_1597_v4_))]
                d_1658_v49_: _dafny.MultiSet
                d_1658_v49_ = _dafny.MultiSet([not((d_1642_v36_)[default__.safeIndex(923, (d_1642_v36_).length(0))])])
                index288_ = default__.safeIndex(524, (d_1591_v0_).length(0))
                (d_1591_v0_)[index288_] = (((d_1658_v49_)[(self).f15] if ((self).f15) in (d_1658_v49_) else len(d_1597_v4_))) - ((self).f14)
                index289_ = default__.safeIndex(524, (d_1591_v0_).length(0))
                (d_1591_v0_)[index289_] = (self).f11
                d_1659_v50_: _dafny.Map
                d_1659_v50_ = _dafny.Map({(self).f12: len((self).f26)})
                d_1660_v51_: C2
                nw248_ = C2()
                nw248_.ctor__((self).f12, (self).f26, len(d_1659_v50_), (d_1591_v0_)[default__.safeIndex(524, (d_1591_v0_).length(0))])
                d_1660_v51_ = nw248_
                (globalState).f6 = (_dafny.MultiSet([d_1660_v51_, d_1660_v51_, d_1660_v51_, d_1660_v51_])).cardinality
                index290_ = default__.safeIndex(923, (d_1642_v36_).length(0))
                (d_1642_v36_)[index290_] = ((self).f14) < (d_1660_v51_.f27)
            d_1661_v52_: _dafny.Map
            d_1661_v52_ = _dafny.Map({d_1644_i7_: 160})
            d_1662_v53_: _dafny.Map
            d_1662_v53_ = _dafny.Map({(self).f12: d_1661_v52_})
            d_1663_v54_: D1
            d_1663_v54_ = D1_DC5((self).f14, (self).f25, (self).f15)
            d_1664_v55_: D17
            d_1664_v55_ = D17_DC47(d_1661_v52_)
            d_1665_v56_: _dafny.MultiSet
            d_1665_v56_ = _dafny.MultiSet([((d_1662_v53_)[(d_1663_v54_).cf6] if ((d_1663_v54_).cf6) in (d_1662_v53_) else (d_1664_v55_).cf89), d_1661_v52_, d_1661_v52_])
            d_1666_v58_: _dafny.Map
            d_1666_v58_ = _dafny.Map({d_1644_i7_: (self).f25})
            d_1667_v59_: str
            d_1667_v59_ = _dafny.CodePoint('t')
            d_1668_v60_: _dafny.Map
            d_1668_v60_ = _dafny.Map({d_1667_v59_: (self).f25})
            d_1669_v61_: _dafny.MultiSet
            d_1669_v61_ = _dafny.MultiSet([d_1668_v60_])
            d_1670_v62_: _dafny.Map
            def iife88_():
                coll44_ = _dafny.Map()
                compr_44_: int
                for compr_44_ in (d_1666_v58_).keys.Elements:
                    d_1671_v57_: int = compr_44_
                    if (d_1671_v57_) in (d_1666_v58_):
                        coll44_[default__.safeModuloInt(d_1671_v57_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwekm"))))] = (self).f11
                return _dafny.Map(coll44_)
            d_1670_v62_ = _dafny.Map({d_1665_v56_: (len((iife88_()
            ).set(((d_1661_v52_)[(self).f12] if ((self).f12) in (d_1661_v52_) else ((d_1669_v61_)[d_1668_v60_] if (d_1668_v60_) in (d_1669_v61_) else d_1644_i7_)), (0) - (len(d_1595_v2_))))) + (d_1644_i7_)})
            d_1672_v63_: _dafny.MultiSet
            d_1672_v63_ = _dafny.MultiSet([(self).f25])
            d_1670_v62_ = (d_1670_v62_).set(_dafny.MultiSet([_dafny.Map({len(((d_1597_v4_).set(default__.safeIndex((self).f12, len(d_1597_v4_)), (self).f15)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1667_v59_, d_1667_v59_])), len(d_1597_v4_)])), len((d_1597_v4_).set(default__.safeIndex((self).f12, len(d_1597_v4_)), (self).f15))), (self).f25)): (d_1672_v63_).cardinality})]), (self).f14)
            d_1673_v64_: _dafny.Set
            d_1673_v64_ = _dafny.Set({(self).f15, (self).f25})
            d_1674_v65_: C2
            nw249_ = C2()
            nw249_.ctor__(default__.safeDivisionInt(len(d_1673_v64_), (self).f12), _dafny.SeqWithoutIsStrInference([d_1667_v59_ for d_1675_i9_ in range(default__.abs(485))]), 187, (-745) + (len(d_1595_v2_)))
            d_1674_v65_ = nw249_

    def m18(self, p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Map = _dafny.Map({})
        r2: _dafny.Seq = _dafny.Seq({})
        r3: str = _dafny.CodePoint('D')
        d_1676_v0_: _dafny.MultiSet
        d_1676_v0_ = _dafny.MultiSet([not ((self).f25) or ((self).f25)])
        d_1676_v0_ = d_1676_v0_
        source28_ = D17_DC48(((self).f25) or ((self).f15), not((p1) <= ((self).f11)), (self).f15)
        if source28_.is_DC48:
            d_1677___mcc_h0_ = source28_.cf90
            d_1678___mcc_h1_ = source28_.cf91
            d_1679___mcc_h2_ = source28_.cf92
            d_1680_cf92_ = d_1679___mcc_h2_
            d_1681_cf91_ = d_1678___mcc_h1_
            d_1682_cf90_ = d_1677___mcc_h0_
            rhs282_ = (0) - ((self).f14)
            rhs283_ = default__.fm0(d_1680_cf92_, globalState)
            rhs284_ = (0) - ((self).f11)
            rhs285_ = d_1681_cf91_
            rhs286_ = not (d_1682_cf90_) or (d_1680_cf92_)
            lhs227_ = globalState
            lhs228_ = globalState
            lhs229_ = globalState
            lhs230_ = globalState
            lhs231_ = globalState
            lhs227_.f6 = rhs282_
            lhs228_.f6 = rhs283_
            lhs229_.f6 = rhs284_
            lhs230_.f2 = rhs285_
            lhs231_.f2 = rhs286_
            d_1683_v1_: _dafny.Array
            nw250_ = _dafny.Array(False, 26)
            d_1683_v1_ = nw250_
            index291_ = default__.safeIndex(451, (d_1683_v1_).length(0))
            (d_1683_v1_)[index291_] = d_1681_cf91_
            index292_ = default__.safeIndex(451, (d_1683_v1_).length(0))
            (d_1683_v1_)[index292_] = (self).f15
            d_1684_v2_: _dafny.Array
            def lambda61_(d_1685_i0_):
                return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f12]))

            init35_ = lambda61_
            nw251_ = _dafny.Array(None, 21)
            for i0_35_ in range(nw251_.length(0)):
                nw251_[i0_35_] = init35_(i0_35_)
            d_1684_v2_ = nw251_
            d_1686_v3_: _dafny.MultiSet
            d_1686_v3_ = _dafny.MultiSet([(self).f14])
            index293_ = default__.safeIndex(216, (d_1684_v2_).length(0))
            (d_1684_v2_)[index293_] = d_1686_v3_
            d_1687_v4_: _dafny.Seq
            d_1687_v4_ = _dafny.SeqWithoutIsStrInference([p1, (self).f14])
            d_1688_v5_: _dafny.Seq
            d_1688_v5_ = _dafny.SeqWithoutIsStrInference([d_1687_v4_])
            d_1689_v6_: _dafny.Seq
            d_1689_v6_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            d_1690_v7_: _dafny.Seq
            d_1690_v7_ = _dafny.SeqWithoutIsStrInference([(d_1689_v6_)[default__.safeIndex((self).f12, len(d_1689_v6_))]])
            index294_ = default__.safeIndex(216, (d_1684_v2_).length(0))
            (d_1684_v2_)[index294_] = (_dafny.MultiSet((d_1688_v5_)[default__.safeIndex((self).fm2((self).f14, d_1682_cf90_, len(d_1690_v7_), (d_1683_v1_)[default__.safeIndex(451, (d_1683_v1_).length(0))], globalState), len(d_1688_v5_))])) | (_dafny.MultiSet([len(d_1689_v6_), default__.fm0(True, globalState), p1]))
            d_1691_v8_: _dafny.Set
            d_1691_v8_ = _dafny.Set({len(d_1690_v7_)})
            d_1692_v9_: _dafny.Map
            d_1692_v9_ = _dafny.Map({(0) - (len(d_1691_v8_)): (self).f12})
            d_1693_v10_: _dafny.Map
            d_1693_v10_ = _dafny.Map({p1: d_1692_v9_})
            (globalState).f6 = (p1) * (len(d_1693_v10_))
        elif source28_.is_DC49:
            d_1694___mcc_h3_ = source28_.cf93
            d_1695_cf93_ = d_1694___mcc_h3_
            index295_ = default__.safeIndex(248, (p0).length(0))
            (p0)[index295_] = p1
            index296_ = default__.safeIndex(248, (p0).length(0))
            (p0)[index296_] = (self).f11
            d_1696_v11_: str
            d_1696_v11_ = _dafny.CodePoint('t')
            d_1697_v12_: _dafny.MultiSet
            d_1697_v12_ = _dafny.MultiSet([d_1696_v11_])
            d_1698_v13_: D0
            d_1698_v13_ = D0_DC0(d_1697_v12_)
            pat_let_tv30_ = d_1697_v12_
            d_1699_v14_: _dafny.Array
            nw252_ = _dafny.Array(None, 25)
            nw252_[int(0)] = d_1698_v13_
            nw252_[int(1)] = D0_DC0(_dafny.MultiSet([d_1696_v11_, d_1696_v11_, _dafny.CodePoint('m')]))
            nw252_[int(2)] = default__.fm56((p0)[default__.safeIndex(248, (p0).length(0))], globalState)
            nw252_[int(3)] = D0_DC0(d_1697_v12_)
            nw252_[int(4)] = d_1698_v13_
            nw252_[int(5)] = d_1698_v13_
            nw252_[int(6)] = D0_DC0(d_1697_v12_)
            nw252_[int(7)] = (d_1698_v13_ if (self).fm1(d_1695_cf93_, (self).fm26((self).f25, globalState), (self).f15, globalState) else default__.fm56((self).f12, globalState))
            nw252_[int(8)] = d_1698_v13_
            nw252_[int(9)] = d_1698_v13_
            nw252_[int(10)] = d_1698_v13_
            nw252_[int(11)] = d_1698_v13_
            nw252_[int(12)] = d_1698_v13_
            nw252_[int(13)] = d_1698_v13_
            nw252_[int(14)] = D0_DC0(d_1697_v12_)
            nw252_[int(15)] = D0_DC0(_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('l')]))
            nw252_[int(16)] = d_1698_v13_
            nw252_[int(17)] = d_1698_v13_
            nw252_[int(18)] = d_1698_v13_
            nw252_[int(19)] = d_1698_v13_
            nw252_[int(20)] = d_1698_v13_
            nw252_[int(21)] = d_1698_v13_
            def iife89_(_pat_let22_0):
                def iife90_(d_1700_dt__update__tmp_h0_):
                    def iife91_(_pat_let23_0):
                        def iife92_(d_1701_dt__update_hcf0_h0_):
                            return D0_DC0(d_1701_dt__update_hcf0_h0_)
                        return iife92_(_pat_let23_0)
                    return iife91_(_dafny.MultiSet([_dafny.CodePoint('m')]))
                return iife90_(_pat_let22_0)
            nw252_[int(22)] = iife89_(d_1698_v13_)
            nw252_[int(23)] = D0_DC0(d_1697_v12_)
            def iife93_(_pat_let24_0):
                def iife94_(d_1702_dt__update__tmp_h1_):
                    def iife95_(_pat_let25_0):
                        def iife96_(d_1703_dt__update_hcf0_h1_):
                            return D0_DC0(d_1703_dt__update_hcf0_h1_)
                        return iife96_(_pat_let25_0)
                    return iife95_(pat_let_tv30_)
                return iife94_(_pat_let24_0)
            nw252_[int(24)] = iife93_(d_1698_v13_)
            d_1699_v14_ = nw252_
            index297_ = default__.safeIndex(90, (d_1699_v14_).length(0))
            (d_1699_v14_)[index297_] = d_1698_v13_
            index298_ = default__.safeIndex(90, (d_1699_v14_).length(0))
            (d_1699_v14_)[index298_] = d_1698_v13_
            d_1704_v15_: _dafny.Array
            nw253_ = _dafny.Array(int(0), 17)
            d_1704_v15_ = nw253_
            d_1705_v16_: _dafny.Array
            d_1705_v16_ = d_1704_v15_
            d_1706_v17_: _dafny.Array
            nw254_ = _dafny.Array(None, 9)
            nw254_[int(0)] = d_1704_v15_
            nw254_[int(1)] = d_1704_v15_
            nw254_[int(2)] = d_1704_v15_
            nw254_[int(3)] = (d_1704_v15_)
            nw254_[int(4)] = d_1704_v15_
            nw254_[int(5)] = d_1704_v15_
            nw254_[int(6)] = d_1704_v15_
            nw254_[int(7)] = d_1704_v15_
            nw254_[int(8)] = d_1704_v15_
            d_1706_v17_ = nw254_
            index299_ = default__.safeIndex(374, (d_1706_v17_).length(0))
            (d_1706_v17_)[index299_] = (d_1704_v15_ if (self).f15 else d_1704_v15_)
            index300_ = default__.safeIndex(374, (d_1706_v17_).length(0))
            (d_1706_v17_)[index300_] = d_1704_v15_
            (globalState).f7 = (((0) - ((p0)[default__.safeIndex(248, (p0).length(0))])) - ((self).f12)) > ((self).f11)
        elif source28_.is_DC50:
            d_1707___mcc_h4_ = source28_.cf94
            d_1708_cf94_ = d_1707___mcc_h4_
            d_1709_v18_: _dafny.Seq
            d_1709_v18_ = _dafny.SeqWithoutIsStrInference([(self).f25, (self).f15])
            d_1710_v19_: _dafny.Array
            nw255_ = _dafny.Array(None, 6)
            nw255_[int(0)] = (d_1709_v18_) + (d_1709_v18_)
            nw255_[int(1)] = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f25])
            nw255_[int(2)] = d_1709_v18_
            nw255_[int(3)] = d_1709_v18_
            nw255_[int(4)] = d_1709_v18_
            nw255_[int(5)] = d_1709_v18_
            d_1710_v19_ = nw255_
            def lambda62_(d_1711_v18_):
                def lambda63_(d_1712_i1_):
                    return d_1711_v18_

                return lambda63_

            init36_ = lambda62_(d_1709_v18_)
            nw256_ = _dafny.Array(None, 12)
            for i0_36_ in range(nw256_.length(0)):
                nw256_[i0_36_] = init36_(i0_36_)
            d_1710_v19_ = nw256_
            (globalState).f2 = (self).f25
            (globalState).f6 = (self).f12
            (globalState).f6 = len((self).f26)
        elif True:
            d_1713___mcc_h5_ = source28_.cf89
            d_1714_cf89_ = d_1713___mcc_h5_
            rhs287_ = (self).f14
            rhs288_ = (0) - (default__.safeModuloInt((p1 if (self).f25 else 956), (self).f12))
            lhs232_ = globalState
            lhs233_ = globalState
            lhs232_.f6 = rhs287_
            lhs233_.f6 = rhs288_
            index301_ = default__.safeIndex(47, (p0).length(0))
            (p0)[index301_] = (self).f11
            d_1715_v20_: _dafny.Seq
            d_1715_v20_ = _dafny.SeqWithoutIsStrInference([(self).f25])
            index302_ = default__.safeIndex(47, (p0).length(0))
            (p0)[index302_] = len((d_1715_v20_) + ((d_1715_v20_) + (d_1715_v20_)))
            d_1716_v21_: str
            d_1716_v21_ = _dafny.CodePoint('n')
            d_1717_v22_: _dafny.Map
            d_1717_v22_ = _dafny.Map({d_1716_v21_: (0) - ((0) - (p1))})
            d_1718_v23_: _dafny.Map
            d_1718_v23_ = _dafny.Map({p1: (self).fm1((self).f11, (self).f15, (self).f25, globalState)})
            d_1719_v24_: _dafny.Seq
            d_1719_v24_ = _dafny.SeqWithoutIsStrInference([len(d_1718_v23_)])
            d_1720_v25_: _dafny.Set
            d_1720_v25_ = _dafny.Set({(self).f11, len(d_1719_v24_)})
            d_1721_v26_: _dafny.Set
            d_1721_v26_ = _dafny.Set({((0) - (p1)) + (861), ((p0)[default__.safeIndex(47, (p0).length(0))] if (self).f25 else (0) - (len(default__.fm32((self).f15, d_1717_v22_, d_1720_v25_, (self).f25, globalState)))), p1, p1})
            d_1721_v26_ = d_1721_v26_
            if (self).f15:
                d_1716_v21_ = d_1716_v21_
                (globalState).f6 = (self).f11
                (globalState).f2 = True
                d_1722_v27_: _dafny.Array
                nw257_ = _dafny.Array(None, 2)
                nw257_[int(0)] = ((self).f11) - ((self).f12)
                nw257_[int(1)] = (self).f14
                d_1722_v27_ = nw257_
                r0 = d_1722_v27_
                d_1723_v28_: _dafny.Array
                def lambda64_(d_1724_i2_):
                    return (self).f25

                init37_ = lambda64_
                nw258_ = _dafny.Array(None, 15)
                for i0_37_ in range(nw258_.length(0)):
                    nw258_[i0_37_] = init37_(i0_37_)
                d_1723_v28_ = nw258_
                rhs289_ = True
                rhs290_ = d_1723_v28_
                lhs234_ = globalState
                lhs234_.f7 = rhs289_
                d_1723_v28_ = rhs290_
            elif True:
                d_1725_v29_: D17
                d_1725_v29_ = D17_DC49(len((self).f13))
                d_1726_v30_: _dafny.Set
                def iife97_(_pat_let26_0):
                    def iife98_(d_1727_dt__update__tmp_h3_):
                        def iife99_(_pat_let27_0):
                            def iife100_(d_1728_dt__update_hcf93_h0_):
                                return D17_DC49(d_1728_dt__update_hcf93_h0_)
                            return iife100_(_pat_let27_0)
                        return iife99_(-683)
                    return iife98_(_pat_let26_0)
                d_1726_v30_ = _dafny.Set({D17_DC49((self).fm2((self).f11, (self).f25, len(d_1719_v24_), (self).f15, globalState)), iife97_(d_1725_v29_)})
                d_1729_v31_: C2
                nw259_ = C2()
                nw259_.ctor__((self).f14, (self).f26, 48, len(d_1726_v30_))
                d_1729_v31_ = nw259_
                d_1730_v32_: _dafny.Array
                nw260_ = _dafny.Array(None, 26)
                nw260_[int(0)] = (self).f25
                nw260_[int(1)] = (self).f25
                nw260_[int(2)] = (self).f25
                nw260_[int(3)] = (self).f25
                nw260_[int(4)] = False
                nw260_[int(5)] = (self).f25
                nw260_[int(6)] = (self).f25
                nw260_[int(7)] = (self).f25
                nw260_[int(8)] = (self).f15
                nw260_[int(9)] = (self).f15
                nw260_[int(10)] = (self).f15
                nw260_[int(11)] = (self).f15
                nw260_[int(12)] = (self).f25
                nw260_[int(13)] = False
                nw260_[int(14)] = (self).f25
                nw260_[int(15)] = (self).f25
                nw260_[int(16)] = (d_1729_v31_).fm41(globalState)
                nw260_[int(17)] = True
                nw260_[int(18)] = (self).f15
                nw260_[int(19)] = (self).f15
                nw260_[int(20)] = (self).f25
                nw260_[int(21)] = False
                nw260_[int(22)] = True
                nw260_[int(23)] = (self).f15
                nw260_[int(24)] = not((self).f25)
                nw260_[int(25)] = (self).f25
                d_1730_v32_ = nw260_
                (d_1729_v31_).f27 = len(_dafny.SeqWithoutIsStrInference([d_1730_v32_, d_1730_v32_, d_1730_v32_]))
                d_1731_v33_: _dafny.Array
                nw261_ = _dafny.Array(None, 13)
                nw261_[int(0)] = (self).f11
                nw261_[int(1)] = d_1729_v31_.f27
                nw261_[int(2)] = len(self.f16)
                nw261_[int(3)] = (0) - ((self).f11)
                nw261_[int(4)] = d_1729_v31_.f27
                nw261_[int(5)] = 351
                nw261_[int(6)] = d_1729_v31_.f27
                nw261_[int(7)] = p1
                nw261_[int(8)] = d_1729_v31_.f27
                nw261_[int(9)] = p1
                nw261_[int(10)] = d_1729_v31_.f27
                nw261_[int(11)] = d_1729_v31_.f27
                nw261_[int(12)] = (0) - ((self).f12)
                d_1731_v33_ = nw261_
                (self).m17(d_1730_v32_, d_1731_v33_, d_1731_v33_, globalState)
                d_1732_v34_: D12
                d_1732_v34_ = D12_DC36((p0)[default__.safeIndex(47, (p0).length(0))], (self).f11, (self).f11)
                d_1733_v35_: _dafny.Set
                d_1733_v35_ = _dafny.Set({d_1732_v34_})
                d_1734_v36_: _dafny.Seq
                d_1734_v36_ = _dafny.SeqWithoutIsStrInference([d_1733_v35_, d_1733_v35_, d_1733_v35_, _dafny.Set({D12_DC36(d_1729_v31_.f27, len(d_1719_v24_), d_1729_v31_.f27)}), d_1733_v35_])
                d_1735_v37_: _dafny.Seq
                d_1735_v37_ = _dafny.SeqWithoutIsStrInference([D12_DC36(p1, (self).f12, len(d_1718_v23_)), d_1732_v34_, d_1732_v34_, d_1732_v34_, d_1732_v34_])
                def iife101_():
                    coll45_ = _dafny.Set()
                    compr_45_: D12
                    for compr_45_ in (d_1735_v37_).Elements:
                        d_1736_v38_: D12 = compr_45_
                        if (d_1736_v38_) in (d_1735_v37_):
                            coll45_ = coll45_.union(_dafny.Set([d_1736_v38_]))
                    return _dafny.Set(coll45_)
                d_1734_v36_ = _dafny.SeqWithoutIsStrInference([iife101_()
                , d_1733_v35_, (d_1733_v35_) - (d_1733_v35_)])
                d_1737_v39_: _dafny.Map
                d_1737_v39_ = _dafny.Map({len((self).f26): self.f16})
                d_1738_v41_: _dafny.Array
                nw262_ = _dafny.Array(None, 7)
                nw262_[int(0)] = d_1737_v39_
                nw262_[int(1)] = d_1737_v39_
                nw262_[int(2)] = default__.fm57((self).f25, d_1729_v31_.f27, globalState)
                nw262_[int(3)] = d_1737_v39_
                def iife102_():
                    coll46_ = _dafny.Map()
                    compr_46_: int
                    for compr_46_ in (d_1718_v23_).keys.Elements:
                        d_1739_v40_: int = compr_46_
                        if (d_1739_v40_) in (d_1718_v23_):
                            coll46_[default__.safeModuloInt(d_1739_v40_, len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_1729_v31_.f27])).cardinality for d_1740_i3_ in range(default__.abs(884))])))] = self.f16
                    return _dafny.Map(coll46_)
                nw262_[int(4)] = (d_1737_v39_) | ((iife102_()
                ).set(p1, (self).f26))
                nw262_[int(5)] = (d_1737_v39_) | (d_1737_v39_)
                nw262_[int(6)] = _dafny.Map({(p0)[default__.safeIndex(47, (p0).length(0))]: (self).f13})
                d_1738_v41_ = nw262_
                index303_ = default__.safeIndex(102, (d_1738_v41_).length(0))
                (d_1738_v41_)[index303_] = (_dafny.Map({(self).f14: self.f16})) | (d_1737_v39_)
                index304_ = default__.safeIndex(102, (d_1738_v41_).length(0))
                (d_1738_v41_)[index304_] = _dafny.Map({(d_1729_v31_).fm40(len((self).f13), (self).f15, d_1716_v21_, not((self).f15), globalState): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhjdtwyrw"))})
        (globalState).f7 = (((self).f14 if (self).f25 else (self).f14)) > ((self).f12)
        d_1741_v42_: _dafny.Array
        nw263_ = _dafny.Array(False, 25)
        d_1741_v42_ = nw263_
        d_1742_v43_: C1
        nw264_ = C1()
        nw264_.ctor__(d_1741_v42_, (self).f14, p1, (self).f11, (self).f15, (self).f26)
        d_1742_v43_ = nw264_
        d_1743_v44_: _dafny.Map
        d_1743_v44_ = _dafny.Map({d_1741_v42_: 396})
        d_1744_v45_: D5
        d_1744_v45_ = D5_DC18((d_1743_v44_).set(d_1742_v43_.f24, (self).f12), (self).f11)
        d_1745_v46_: D5
        d_1745_v46_ = D5_DC18((d_1744_v45_).cf35, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))))
        (globalState).f6 = (d_1745_v46_).cf36
        rhs291_ = (self).f14
        rhs292_ = (self).f15
        lhs235_ = globalState
        lhs236_ = globalState
        lhs235_.f6 = rhs291_
        lhs236_.f2 = rhs292_
        d_1746_v47_: _dafny.Map
        d_1746_v47_ = _dafny.Map({(self).f15: (self).f25})
        d_1747_v48_: _dafny.Map
        d_1747_v48_ = _dafny.Map({_dafny.MultiSet([d_1746_v47_, d_1746_v47_, _dafny.Map({((d_1746_v47_)[(self).f25] if ((self).f25) in (d_1746_v47_) else (self).f15): (self).f25}), d_1746_v47_, d_1746_v47_]): p0})
        d_1748_v49_: _dafny.Map
        d_1748_v49_ = _dafny.Map({(self).f25: d_1746_v47_})
        d_1749_v50_: _dafny.MultiSet
        d_1749_v50_ = _dafny.MultiSet([((d_1748_v49_)[True] if (True) in (d_1748_v49_) else _dafny.Map({(self).f15: (self).f25}))])
        d_1750_v51_: _dafny.Seq
        d_1750_v51_ = _dafny.SeqWithoutIsStrInference([p1, len((self).f26), p1])
        d_1751_v52_: _dafny.Seq
        d_1751_v52_ = _dafny.SeqWithoutIsStrInference([(d_1749_v50_).set(d_1746_v47_, default__.abs(len(d_1750_v51_))), d_1749_v50_, d_1749_v50_])
        r0 = ((d_1747_v48_)[(d_1751_v52_)[default__.safeIndex(p1, len(d_1751_v52_))]] if ((d_1751_v52_)[default__.safeIndex(p1, len(d_1751_v52_))]) in (d_1747_v48_) else p0)
        d_1752_v53_: _dafny.Map
        d_1752_v53_ = _dafny.Map({(self).f14: d_1750_v51_})
        r1 = d_1752_v53_
        r2 = d_1750_v51_
        d_1753_v54_: str
        d_1753_v54_ = _dafny.CodePoint('u')
        r3 = d_1753_v54_
        return r0, r1, r2, r3

    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26

class C4(T3, T2, T1, T0):
    def  __init__(self):
        self._f16: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f11: int = int(0)
        self._f12: int = int(0)
        self._f14: int = int(0)
        self._f15: bool = False
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
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
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f16, f14, f15, f13, f11, f12):
        (self).f16 = f16
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f13 = f13
        (self)._f11 = f11
        (self)._f12 = f12

    def fm2(self, p0, p1, p2, p3, globalState):
        def iife103_():
            coll47_ = _dafny.Map()
            compr_47_: str
            for compr_47_ in (_dafny.Map({_dafny.CodePoint('d'): 756})).keys.Elements:
                d_1754_v0_: str = compr_47_
                if (d_1754_v0_) in (_dafny.Map({_dafny.CodePoint('d'): 756})):
                    coll47_[d_1754_v0_] = _dafny.Map({(self).f15: (self).f15})
            return _dafny.Map(coll47_)
        return len(iife103_()
        )

    def fm3(self, p0, p1, p2, globalState):
        return (_dafny.Map({(self).f15: 362})) | (_dafny.Map({(self).f15: (_dafny.MultiSet([(self).f15])).cardinality}))

    def fm1(self, p0, p1, p2, globalState):
        return False

    def fm24(self, p0, p1, globalState):
        if (self).f15:
            return not ((self).f15) or ((self).f15)
        elif True:
            return (_dafny.MultiSet([(self).f15, (self).f15])).ispropersubset(_dafny.MultiSet([(self).f15]))

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: T0 = None
        d_1755_v0_: _dafny.Seq
        d_1755_v0_ = _dafny.SeqWithoutIsStrInference([self.f16, (self).f13])
        d_1756_v2_: T3
        nw265_ = C3()
        def iife104_():
            coll48_ = _dafny.Map()
            compr_48_: int
            for compr_48_ in _dafny.IntegerRange(128, 60):
                d_1757_v1_: int = compr_48_
                if ((128) <= (d_1757_v1_)) and ((d_1757_v1_) < (60)):
                    coll48_[default__.safeModuloInt(d_1757_v1_, (self).f11)] = p1
            return _dafny.Map(coll48_)
        nw265_.ctor__((self).f15, (d_1755_v0_)[default__.safeIndex(len(iife104_()
        ), len(d_1755_v0_))], self.f16, (self).f14, p1, (self).f13, (self).f14, (self).f12)
        d_1756_v2_ = nw265_
        d_1758_v3_: _dafny.Seq
        d_1758_v3_ = _dafny.SeqWithoutIsStrInference([d_1756_v2_, d_1756_v2_, d_1756_v2_])
        d_1759_v4_: _dafny.Map
        d_1759_v4_ = _dafny.Map({(self).f15: p1})
        d_1760_v5_: _dafny.Map
        d_1760_v5_ = _dafny.Map({len(d_1759_v4_): (self).f15})
        d_1761_v6_: _dafny.Seq
        d_1761_v6_ = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15, (d_1756_v2_).f15, (d_1756_v2_).f15, (d_1756_v2_).f15])
        d_1762_v7_: str
        d_1762_v7_ = _dafny.CodePoint('n')
        d_1763_v8_: _dafny.Set
        d_1763_v8_ = _dafny.Set({p1, (d_1756_v2_).f15})
        d_1764_v9_: _dafny.Map
        d_1764_v9_ = _dafny.Map({d_1762_v7_: d_1763_v8_})
        d_1765_v10_: _dafny.Array
        def lambda65_(d_1766_v2_):
            def lambda66_(d_1767_i0_):
                return (d_1766_v2_).f15

            return lambda66_

        init38_ = lambda65_(d_1756_v2_)
        nw266_ = _dafny.Array(None, 18)
        for i0_38_ in range(nw266_.length(0)):
            nw266_[i0_38_] = init38_(i0_38_)
        d_1765_v10_ = nw266_
        d_1768_v11_: _dafny.Map
        d_1768_v11_ = _dafny.Map({len(((d_1764_v9_)[d_1762_v7_] if (d_1762_v7_) in (d_1764_v9_) else d_1763_v8_)): d_1765_v10_})
        d_1769_v12_: _dafny.Array
        nw267_ = _dafny.Array(None, 18)
        nw267_[int(0)] = (self).f15
        nw267_[int(1)] = p1
        nw267_[int(2)] = (self).fm24((default__.fm25((self).f12, p1, globalState)).cardinality, (self).f15, globalState)
        nw267_[int(3)] = p1
        nw267_[int(4)] = (self).f15
        nw267_[int(5)] = not((d_1758_v3_) < (d_1758_v3_))
        nw267_[int(6)] = p1
        nw267_[int(7)] = p1
        nw267_[int(8)] = ((self).f14) <= ((d_1756_v2_).f11)
        nw267_[int(9)] = ((self).f15 if ((d_1760_v5_)[983] if (983) in (d_1760_v5_) else (self).f15) else False)
        nw267_[int(10)] = (d_1761_v6_)[default__.safeIndex((d_1756_v2_).f14, len(d_1761_v6_))]
        nw267_[int(11)] = (self).f15
        nw267_[int(12)] = (d_1756_v2_).f15
        nw267_[int(13)] = (len(d_1756_v2_.f16)) not in (d_1768_v11_)
        nw267_[int(14)] = (d_1756_v2_).f15
        nw267_[int(15)] = p1
        nw267_[int(16)] = not((len(self.f16)) >= ((self).f14))
        nw267_[int(17)] = p1
        d_1769_v12_ = nw267_
        d_1770_v13_: _dafny.MultiSet
        d_1770_v13_ = _dafny.MultiSet([False, p1])
        index305_ = default__.safeIndex(96, (d_1769_v12_).length(0))
        (d_1769_v12_)[index305_] = (d_1770_v13_).issubset(d_1770_v13_)
        index306_ = default__.safeIndex(96, (d_1769_v12_).length(0))
        (d_1769_v12_)[index306_] = (self).f15
        d_1771_v14_: _dafny.Map
        d_1771_v14_ = _dafny.Map({310: (d_1756_v2_).f14})
        d_1772_v15_: _dafny.Map
        d_1772_v15_ = _dafny.Map({d_1756_v2_.f16: (d_1756_v2_).f12})
        d_1773_v16_: D3
        d_1773_v16_ = D3_DC10(False, len(d_1756_v2_.f16), d_1772_v15_, d_1756_v2_.f16, len(d_1756_v2_.f16))
        d_1774_v17_: D3
        d_1774_v17_ = D3_DC10(p1, len(default__.fm42((d_1756_v2_).f12, (d_1756_v2_).f15, d_1771_v14_, globalState)), (d_1773_v16_).cf21, (d_1756_v2_).f13, (d_1756_v2_).f14)
        d_1775_i1_: int
        d_1775_i1_ = 0
        with _dafny.label("8"):
            while (d_1761_v6_)[default__.safeIndex(default__.fm0((d_1773_v16_).cf19, globalState), len(d_1761_v6_))]:
                with _dafny.c_label("8"):
                    if (d_1775_i1_) >= (100):
                        raise _dafny.Break("8")
                    d_1775_i1_ = (d_1775_i1_) + (1)
                    (globalState).f6 = (d_1756_v2_).f12
                    d_1776_v18_: _dafny.Seq
                    d_1776_v18_ = _dafny.SeqWithoutIsStrInference([(d_1756_v2_).f14, (self).f11, (d_1756_v2_).f12])
                    rhs293_ = (d_1776_v18_) + (_dafny.SeqWithoutIsStrInference([(d_1756_v2_).f14 for d_1777_i2_ in range(default__.abs(809))]))
                    rhs294_ = not ((d_1761_v6_)[default__.safeIndex((self).f12, len(d_1761_v6_))]) or (((d_1756_v2_).f14) != ((d_1756_v2_).f11))
                    rhs295_ = p1
                    lhs237_ = globalState
                    lhs238_ = globalState
                    d_1776_v18_ = rhs293_
                    lhs237_.f7 = rhs294_
                    lhs238_.f7 = rhs295_
                    (globalState).f7 = ((d_1756_v2_).f11) > ((d_1756_v2_).f11)
                    d_1778_v19_: C0
                    nw268_ = C0()
                    nw268_.ctor__()
                    d_1778_v19_ = nw268_
                    pass
            pass
        d_1779_v20_: _dafny.Array
        nw269_ = _dafny.Array(int(0), 15)
        d_1779_v20_ = nw269_
        index307_ = default__.safeIndex(616, (d_1779_v20_).length(0))
        (d_1779_v20_)[index307_] = (0) - ((len(d_1759_v4_)) * ((self).f12))
        d_1780_v21_: _dafny.Set
        d_1780_v21_ = _dafny.Set({self.f16})
        index308_ = default__.safeIndex(616, (d_1779_v20_).length(0))
        (d_1779_v20_)[index308_] = (self).fm2((self).f11, (_dafny.Set({_dafny.SeqWithoutIsStrInference([d_1762_v7_ for d_1781_i3_ in range(default__.abs(928))])})).isdisjoint(d_1780_v21_), (self).f11, p1, globalState)
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1779_v20_).length(0)):
            d_1782_i4_: int = guard_loop_4_
            if (True) and (((0) <= (d_1782_i4_)) and ((d_1782_i4_) < ((d_1779_v20_).length(0)))):
                (d_1779_v20_)[(d_1782_i4_)] = (d_1782_i4_) * ((d_1756_v2_).f14)
        d_1783_i5_: int
        d_1783_i5_ = 0
        with _dafny.label("9"):
            while p1:
                with _dafny.c_label("9"):
                    if (d_1783_i5_) >= (100):
                        raise _dafny.Break("9")
                    d_1783_i5_ = (d_1783_i5_) + (1)
                    if not((d_1769_v12_)[default__.safeIndex(96, (d_1769_v12_).length(0))]):
                        d_1784_v22_: D2
                        d_1784_v22_ = D2_DC8(p1, (d_1769_v12_)[default__.safeIndex(96, (d_1769_v12_).length(0))], d_1756_v2_.f16)
                        d_1785_v23_: _dafny.MultiSet
                        d_1785_v23_ = _dafny.MultiSet([d_1762_v7_, d_1762_v7_])
                        d_1786_v24_: D0
                        d_1786_v24_ = D0_DC0(d_1785_v23_)
                        d_1787_v25_: _dafny.Set
                        d_1787_v25_ = _dafny.Set({default__.fm56(len((d_1756_v2_).f13), globalState), d_1786_v24_})
                        d_1788_v26_: D3
                        d_1788_v26_ = D3_DC11(647, d_1784_v22_, d_1787_v25_, (d_1756_v2_).f15)
                        pat_let_tv31_ = d_1756_v2_
                        d_1789_v27_: _dafny.Set
                        def iife105_(_pat_let28_0):
                            def iife106_(d_1790_dt__update__tmp_h0_):
                                def iife107_(_pat_let29_0):
                                    def iife108_(d_1791_dt__update_hcf27_h0_):
                                        def iife109_(_pat_let30_0):
                                            def iife110_(d_1792_dt__update_hcf24_h0_):
                                                return D3_DC11(d_1792_dt__update_hcf24_h0_, (d_1790_dt__update__tmp_h0_).cf25, (d_1790_dt__update__tmp_h0_).cf26, d_1791_dt__update_hcf27_h0_)
                                            return iife110_(_pat_let30_0)
                                        return iife109_((self).f12)
                                    return iife108_(_pat_let29_0)
                                return iife107_((pat_let_tv31_).f15)
                            return iife106_(_pat_let28_0)
                        d_1789_v27_ = _dafny.Set({d_1788_v26_, iife105_(default__.fm46((d_1756_v2_).f15, -588, (self).f12, (d_1756_v2_).f14, globalState)), D3_DC11(344, d_1784_v22_, d_1787_v25_, p1), d_1788_v26_, d_1788_v26_})
                        pat_let_tv32_ = d_1784_v22_
                        pat_let_tv33_ = d_1756_v2_
                        def iife111_(_pat_let31_0):
                            def iife112_(d_1793_dt__update__tmp_h1_):
                                def iife113_(_pat_let32_0):
                                    def iife114_(d_1794_dt__update_hcf25_h0_):
                                        def iife115_(_pat_let33_0):
                                            def iife116_(d_1795_dt__update_hcf27_h1_):
                                                return D3_DC11((d_1793_dt__update__tmp_h1_).cf24, d_1794_dt__update_hcf25_h0_, (d_1793_dt__update__tmp_h1_).cf26, d_1795_dt__update_hcf27_h1_)
                                            return iife116_(_pat_let33_0)
                                        return iife115_((pat_let_tv33_).f15)
                                    return iife114_(_pat_let32_0)
                                return iife113_(pat_let_tv32_)
                            return iife112_(_pat_let31_0)
                        d_1789_v27_ = (d_1789_v27_) - (_dafny.Set({d_1788_v26_, d_1788_v26_, iife111_(d_1788_v26_), d_1788_v26_, d_1788_v26_}))
                        index309_ = default__.safeIndex(616, (d_1779_v20_).length(0))
                        (d_1779_v20_)[index309_] = ((d_1756_v2_).f11) - (((self).f12 if (d_1756_v2_).f15 else 253))
                        d_1796_v28_: _dafny.Array
                        nw270_ = _dafny.Array(_dafny.CodePoint('D'), 2)
                        d_1796_v28_ = nw270_
                        index310_ = default__.safeIndex(658, (d_1796_v28_).length(0))
                        (d_1796_v28_)[index310_] = d_1762_v7_
                        d_1797_v29_: _dafny.Map
                        d_1797_v29_ = _dafny.Map({d_1761_v6_: False})
                        index311_ = default__.safeIndex(658, (d_1796_v28_).length(0))
                        index312_ = default__.safeIndex(616, (d_1779_v20_).length(0))
                        rhs296_ = (self).f11
                        rhs297_ = d_1762_v7_
                        rhs298_ = len(((d_1797_v29_) | (d_1797_v29_)).set(d_1761_v6_, (d_1770_v13_).isdisjoint(d_1770_v13_)))
                        lhs239_ = globalState
                        lhs240_ = d_1796_v28_
                        lhs241_ = default__.safeIndex(658, (d_1796_v28_).length(0))
                        lhs242_ = d_1779_v20_
                        lhs243_ = default__.safeIndex(616, (d_1779_v20_).length(0))
                        lhs239_.f6 = rhs296_
                        lhs240_[lhs241_] = rhs297_
                        lhs242_[lhs243_] = rhs298_
                        d_1798_v30_: T2
                        nw271_ = C1()
                        nw271_.ctor__(d_1769_v12_, (d_1756_v2_).f14, (d_1779_v20_)[default__.safeIndex(616, (d_1779_v20_).length(0))], ((d_1756_v2_).f12) * (len((d_1756_v2_).f13)), (default__.fm54((self).f12, globalState)).ispropersubset(_dafny.MultiSet([(d_1756_v2_).f15, False])), d_1756_v2_.f16)
                        d_1798_v30_ = nw271_
                        d_1798_v30_ = d_1798_v30_
                        d_1799_v31_: _dafny.Map
                        d_1799_v31_ = _dafny.Map({(d_1779_v20_)[default__.safeIndex(616, (d_1779_v20_).length(0))]: default__.fm34((d_1796_v28_)[default__.safeIndex(658, (d_1796_v28_).length(0))], globalState)})
                        (globalState).f2 = (((self).f14) + (len(d_1799_v31_))) >= (597)
                    elif True:
                        (globalState).f6 = default__.safeModuloInt(((d_1756_v2_).f11) + ((d_1756_v2_).f12), (self).fm2(291, (d_1756_v2_).f15, len((d_1756_v2_).f13), (d_1756_v2_).f15, globalState))
                        d_1800_v33_: C3
                        nw272_ = C3()
                        def iife117_():
                            coll49_ = _dafny.Map()
                            compr_49_: int
                            for compr_49_ in _dafny.IntegerRange(196, 54):
                                d_1801_v32_: int = compr_49_
                                if ((196) <= (d_1801_v32_)) and ((d_1801_v32_) < (54)):
                                    coll49_[default__.safeDivisionInt(d_1801_v32_, len((d_1756_v2_).f13))] = (d_1756_v2_).f11
                            return _dafny.Map(coll49_)
                        nw272_.ctor__((d_1756_v2_).f15, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psqvohow")), d_1756_v2_.f16, (d_1756_v2_).f14, p1, (self).f13, len(iife117_()
                        ), (d_1756_v2_).f11)
                        d_1800_v33_ = nw272_
                        (globalState).f2 = (default__.fm43((d_1756_v2_).f15, (self).f15, globalState)) < (d_1761_v6_)
                        d_1802_v34_: D4
                        d_1802_v34_ = D4_DC15((d_1756_v2_).f14)
                        d_1803_v35_: _dafny.Map
                        d_1803_v35_ = _dafny.Map({d_1802_v34_: (self).f15})
                        d_1803_v35_ = d_1803_v35_
                        (globalState).f7 = ((d_1770_v13_ if (d_1756_v2_).f15 else d_1770_v13_)).isdisjoint((d_1770_v13_) - (default__.fm54((d_1756_v2_).f14, globalState)))
                    d_1804_v36_: _dafny.Map
                    d_1804_v36_ = _dafny.Map({((self).f14 if (d_1756_v2_).f15 else (d_1756_v2_).f12): d_1762_v7_})
                    d_1805_v37_: _dafny.Map
                    d_1805_v37_ = _dafny.Map({(d_1756_v2_).f14: _dafny.Map({(d_1779_v20_)[default__.safeIndex(616, (d_1779_v20_).length(0))]: _dafny.CodePoint('o')})})
                    d_1804_v36_ = ((d_1804_v36_) | (d_1804_v36_)).set(len(d_1805_v37_), d_1762_v7_)
                    (globalState).f6 = (0) - ((d_1756_v2_).f12)
                    d_1806_v38_: _dafny.Map
                    d_1806_v38_ = _dafny.Map({d_1762_v7_: p1})
                    d_1806_v38_ = (d_1806_v38_).set(d_1762_v7_, (d_1756_v2_).f15)
                    pass
            pass
        hi7_ = (0) - ((-809) + ((d_1756_v2_).f12))
        for d_1807_i6_ in range(879, hi7_):
            d_1808_v39_: _dafny.Map
            d_1808_v39_ = _dafny.Map({(d_1756_v2_).f14: d_1756_v2_.f16})
            d_1809_v40_: T2
            nw273_ = C1()
            nw273_.ctor__(d_1765_v10_, (d_1756_v2_).f12, (self).f11, 282, (d_1756_v2_).f15, ((d_1808_v39_)[d_1807_i6_] if (d_1807_i6_) in (d_1808_v39_) else _dafny.SeqWithoutIsStrInference([d_1762_v7_ for d_1810_i7_ in range(default__.abs(-986))])))
            d_1809_v40_ = nw273_
            d_1811_v41_: D11
            d_1811_v41_ = D11_DC33(d_1809_v40_)
            d_1811_v41_ = d_1811_v41_
            d_1812_v42_: _dafny.Seq
            d_1812_v42_ = _dafny.SeqWithoutIsStrInference([(self).f12])
            d_1813_v43_: D13
            d_1813_v43_ = D13_DC37(_dafny.MultiSet([(d_1809_v40_).f11]))
            (globalState).f2 = (len((((d_1812_v42_).set(default__.safeIndex((d_1779_v20_)[default__.safeIndex(616, (d_1779_v20_).length(0))], len(d_1812_v42_)), ((d_1813_v43_).cf67).cardinality)) + (d_1812_v42_)).set(default__.safeIndex((d_1809_v40_).f14, len(((d_1812_v42_).set(default__.safeIndex((d_1779_v20_)[default__.safeIndex(616, (d_1779_v20_).length(0))], len(d_1812_v42_)), ((d_1813_v43_).cf67).cardinality)) + (d_1812_v42_))), (d_1809_v40_).f11))) in (d_1812_v42_)
            d_1814_v44_: _dafny.Map
            d_1814_v44_ = _dafny.Map({(d_1756_v2_).f14: d_1771_v14_})
            (globalState).f2 = (d_1809_v40_).fm1(len(d_1814_v44_), (not((d_1769_v12_)[default__.safeIndex(96, (d_1769_v12_).length(0))])) == ((d_1756_v2_).f15), p1, globalState)
            (globalState).f7 = (_dafny.SeqWithoutIsStrInference([d_1762_v7_ for d_1815_i8_ in range(default__.abs(219))])) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ie")))
        r0 = (d_1779_v20_)[default__.safeIndex(616, (d_1779_v20_).length(0))]
        r1 = not ((d_1761_v6_) <= (d_1761_v6_)) or (((d_1756_v2_).f13) == ((self).f13))
        nw274_ = C1()
        nw274_.ctor__(d_1769_v12_, (0) - ((d_1779_v20_)[default__.safeIndex(616, (d_1779_v20_).length(0))]), len(d_1771_v14_), (self).f14, (d_1769_v12_)[default__.safeIndex(96, (d_1769_v12_).length(0))], _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1816_i9_ in range(default__.abs(421))]))
        r2 = nw274_
        return r0, r1, r2

    def m4(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        d_1817_v0_: _dafny.Map
        d_1817_v0_ = _dafny.Map({(self).f11: (self).f14})
        d_1818_v1_: _dafny.Seq
        d_1818_v1_ = _dafny.SeqWithoutIsStrInference([348])
        d_1819_v2_: D12
        d_1819_v2_ = D12_DC36((self).f14, (self).f11, (self).f11)
        d_1820_v3_: D17
        d_1820_v3_ = D17_DC47(d_1817_v0_)
        d_1821_v4_: _dafny.Seq
        d_1821_v4_ = _dafny.SeqWithoutIsStrInference([False, (self).f15, (self).f15, (self).f15, (self).f15])
        d_1822_v5_: _dafny.Array
        nw275_ = _dafny.Array(None, 18)
        nw275_[int(0)] = default__.safeDivisionInt((self).f14, (self).f11)
        nw275_[int(1)] = 659
        nw275_[int(2)] = ((self).f11) - (len(_dafny.Map({(self).f12: (self).f15})))
        nw275_[int(3)] = default__.safeDivisionInt((self).f14, (self).f14)
        nw275_[int(4)] = (len((d_1817_v0_).set((d_1818_v1_)[default__.safeIndex((self).f14, len(d_1818_v1_))], (self).f14))) * ((self).f14)
        nw275_[int(5)] = (0) - (((self).f11) + ((0) - ((self).f14)))
        nw275_[int(6)] = ((self).f11) + ((self).f11)
        nw275_[int(7)] = (self).f12
        nw275_[int(8)] = len(d_1817_v0_)
        nw275_[int(9)] = (0) - (default__.safeDivisionInt((0) - ((self).f14), (d_1819_v2_).cf66))
        nw275_[int(10)] = (self).f12
        nw275_[int(11)] = default__.safeModuloInt((self).f11, (0) - (len((d_1820_v3_).cf89)))
        nw275_[int(12)] = 61
        nw275_[int(13)] = default__.safeModuloInt((self).f11, default__.fm0((self).f15, globalState))
        nw275_[int(14)] = len(d_1821_v4_)
        nw275_[int(15)] = ((self).f12) * ((self).f14)
        nw275_[int(16)] = (0) - ((self).f11)
        nw275_[int(17)] = (self).f14
        d_1822_v5_ = nw275_
        d_1823_v6_: _dafny.MultiSet
        d_1823_v6_ = _dafny.MultiSet([(self).f15, not(not((self).f15))])
        rhs299_ = d_1822_v5_
        rhs300_ = (d_1822_v5_ if (d_1821_v4_)[default__.safeIndex(((d_1823_v6_)[not((self).f15)] if (not((self).f15)) in (d_1823_v6_) else (self).f12), len(d_1821_v4_))] else d_1822_v5_)
        rhs301_ = (self).f15
        lhs244_ = globalState
        d_1822_v5_ = rhs299_
        d_1822_v5_ = rhs300_
        lhs244_.f2 = rhs301_
        d_1824_v7_: D10
        d_1824_v7_ = D10_DC29()
        d_1825_v8_: D10
        d_1825_v8_ = D10_DC32(d_1824_v7_)
        source29_ = d_1825_v8_
        if source29_.is_DC29:
            r0 = ((self).f12) + ((self).f12)
            index313_ = default__.safeIndex(168, (d_1822_v5_).length(0))
            (d_1822_v5_)[index313_] = (self).f12
            d_1826_v9_: _dafny.MultiSet
            d_1826_v9_ = _dafny.MultiSet([(_dafny.MultiSet((self).f13)).cardinality])
            d_1827_v10_: D15
            d_1827_v10_ = D15_DC43(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1828_i0_ in range(default__.abs(780))])), (self).f15, (self).f15)
            index314_ = default__.safeIndex(168, (d_1822_v5_).length(0))
            (d_1822_v5_)[index314_] = default__.safeDivisionInt(((d_1826_v9_)[(self).f12] if ((self).f12) in (d_1826_v9_) else (self).f14), (d_1827_v10_).cf77)
            d_1829_v11_: _dafny.Array
            def lambda67_(d_1830_v9_):
                def lambda68_(d_1831_i1_):
                    return d_1830_v9_

                return lambda68_

            init39_ = lambda67_(d_1826_v9_)
            nw276_ = _dafny.Array(None, 28)
            for i0_39_ in range(nw276_.length(0)):
                nw276_[i0_39_] = init39_(i0_39_)
            d_1829_v11_ = nw276_
            d_1832_v12_: _dafny.Map
            d_1832_v12_ = _dafny.Map({(self).f15: d_1829_v11_})
            d_1833_v13_: _dafny.Seq
            d_1833_v13_ = _dafny.SeqWithoutIsStrInference([d_1829_v11_, d_1829_v11_, d_1829_v11_])
            d_1834_v14_: D18
            d_1834_v14_ = D18_DC51((d_1833_v13_)[default__.safeIndex(684, len(d_1833_v13_))])
            d_1835_v15_: _dafny.Array
            nw277_ = _dafny.Array(None, 6)
            nw277_[int(0)] = d_1829_v11_
            nw277_[int(1)] = d_1829_v11_
            nw277_[int(2)] = ((d_1832_v12_)[(self).f15] if ((self).f15) in (d_1832_v12_) else d_1829_v11_)
            nw277_[int(3)] = (d_1834_v14_).cf95
            nw277_[int(4)] = d_1829_v11_
            nw277_[int(5)] = d_1829_v11_
            d_1835_v15_ = nw277_
            index315_ = default__.safeIndex(374, (d_1835_v15_).length(0))
            (d_1835_v15_)[index315_] = d_1829_v11_
            index316_ = default__.safeIndex(374, (d_1835_v15_).length(0))
            (d_1835_v15_)[index316_] = d_1829_v11_
            if (d_1821_v4_)[default__.safeIndex(len(_dafny.Set({(self).f15, (self).f15, (self).f15})), len(d_1821_v4_))]:
                d_1836_v16_: _dafny.Map
                d_1836_v16_ = _dafny.Map({(self).f14: (len(d_1818_v1_)) > ((self).f14)})
                d_1837_v17_: _dafny.Array
                def lambda69_(d_1838_i2_):
                    return _dafny.SeqWithoutIsStrInference([D5_DC17(_dafny.Set({(self).f13})), D5_DC17(_dafny.Set({(self).f13}))])

                init40_ = lambda69_
                nw278_ = _dafny.Array(None, 25)
                for i0_40_ in range(nw278_.length(0)):
                    nw278_[i0_40_] = init40_(i0_40_)
                d_1837_v17_ = nw278_
                d_1839_v18_: _dafny.Set
                d_1839_v18_ = _dafny.Set({self.f16, (self).f13})
                d_1840_v19_: D5
                d_1840_v19_ = D5_DC17(d_1839_v18_)
                d_1841_v20_: _dafny.Seq
                d_1841_v20_ = _dafny.SeqWithoutIsStrInference([d_1840_v19_])
                index317_ = default__.safeIndex(257, (d_1837_v17_).length(0))
                (d_1837_v17_)[index317_] = d_1841_v20_
                index318_ = default__.safeIndex(168, (d_1822_v5_).length(0))
                index319_ = default__.safeIndex(257, (d_1837_v17_).length(0))
                rhs302_ = (default__.safeModuloInt(len(d_1817_v0_), (d_1822_v5_)[default__.safeIndex(168, (d_1822_v5_).length(0))])) - ((self).f12)
                rhs303_ = _dafny.Map({default__.safeModuloInt((self).f11, (self).f14): not((self).f15)})
                rhs304_ = (self).f11
                rhs305_ = _dafny.SeqWithoutIsStrInference([d_1840_v19_])
                lhs245_ = globalState
                lhs246_ = d_1822_v5_
                lhs247_ = default__.safeIndex(168, (d_1822_v5_).length(0))
                lhs248_ = d_1837_v17_
                lhs249_ = default__.safeIndex(257, (d_1837_v17_).length(0))
                lhs245_.f6 = rhs302_
                d_1836_v16_ = rhs303_
                lhs246_[lhs247_] = rhs304_
                lhs248_[lhs249_] = rhs305_
                r0 = (0) - (((d_1826_v9_)[(self).f11] if ((self).f11) in (d_1826_v9_) else len(self.f16)))
                d_1842_v21_: _dafny.Set
                d_1842_v21_ = _dafny.Set({d_1821_v4_, d_1821_v4_, _dafny.SeqWithoutIsStrInference([True])})
                (globalState).f2 = not ((self).f15) or ((d_1842_v21_) == (d_1842_v21_))
                (globalState).f2 = (self).f15
                (globalState).f7 = not((self).f15)
            elif True:
                d_1843_v22_: _dafny.Set
                d_1843_v22_ = _dafny.Set({True})
                d_1844_v23_: _dafny.Map
                d_1844_v23_ = _dafny.Map({((d_1826_v9_).set((self).f12, default__.abs((self).f12))).cardinality: d_1843_v22_})
                d_1845_v24_: _dafny.Map
                d_1845_v24_ = _dafny.Map({((d_1822_v5_)[default__.safeIndex(168, (d_1822_v5_).length(0))]) <= (len(d_1818_v1_)): ((d_1844_v23_)[(self).f12] if ((self).f12) in (d_1844_v23_) else d_1843_v22_)})
                d_1845_v24_ = (d_1845_v24_).set((self).f15, d_1843_v22_)
                d_1846_v25_: C0
                nw279_ = C0()
                nw279_.ctor__()
                d_1846_v25_ = nw279_
                d_1847_v26_: str
                d_1847_v26_ = _dafny.CodePoint('o')
                d_1848_v27_: _dafny.Map
                d_1848_v27_ = _dafny.Map({d_1847_v26_: (self).f15})
                d_1849_v28_: _dafny.Map
                d_1849_v28_ = _dafny.Map({(d_1822_v5_)[default__.safeIndex(168, (d_1822_v5_).length(0))]: (self).f15})
                d_1850_v29_: _dafny.Map
                d_1850_v29_ = _dafny.Map({d_1849_v28_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsvhdg"))).set(default__.safeIndex((d_1822_v5_)[default__.safeIndex(168, (d_1822_v5_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsvhdg")))), d_1847_v26_)})
                d_1851_v30_: _dafny.Array
                nw280_ = _dafny.Array(None, 16)
                nw280_[int(0)] = len(_dafny.SeqWithoutIsStrInference([d_1847_v26_ for d_1852_i3_ in range(default__.abs(561))]))
                nw280_[int(1)] = (d_1822_v5_)[default__.safeIndex(168, (d_1822_v5_).length(0))]
                nw280_[int(2)] = (0) - ((d_1822_v5_)[default__.safeIndex(168, (d_1822_v5_).length(0))])
                nw280_[int(3)] = (self).f14
                nw280_[int(4)] = (self).f11
                nw280_[int(5)] = (self).f11
                nw280_[int(6)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))
                nw280_[int(7)] = len(d_1818_v1_)
                nw280_[int(8)] = (d_1822_v5_)[default__.safeIndex(168, (d_1822_v5_).length(0))]
                nw280_[int(9)] = len(d_1850_v29_)
                nw280_[int(10)] = (d_1822_v5_)[default__.safeIndex(168, (d_1822_v5_).length(0))]
                nw280_[int(11)] = (self).f14
                nw280_[int(12)] = (0) - ((self).f12)
                nw280_[int(13)] = (self).f14
                nw280_[int(14)] = (self).f14
                nw280_[int(15)] = (self).f14
                d_1851_v30_ = nw280_
                d_1853_v32_: C3
                nw281_ = C3()
                def iife118_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in _dafny.IntegerRange(-953, 262):
                        d_1855_v31_: int = compr_50_
                        if ((-953) <= (d_1855_v31_)) and ((d_1855_v31_) < (262)):
                            coll50_[(d_1855_v31_) * (len(_dafny.Map({(self).f15: (self).f15})))] = (self).f15
                    return _dafny.Map(coll50_)
                nw281_.ctor__((self).f15, self.f16, _dafny.SeqWithoutIsStrInference([d_1847_v26_ for d_1854_i4_ in range(default__.abs(54))]), (self).f12, (self).f15, (self).f13, len(iife118_()
                ), (self).fm2(-827, False, (self).fm2((d_1822_v5_)[default__.safeIndex(168, (d_1822_v5_).length(0))], (self).f15, len(d_1821_v4_), (self).f15, globalState), (self).f15, globalState))
                d_1853_v32_ = nw281_
                d_1856_v33_: _dafny.Map
                d_1856_v33_ = _dafny.Map({d_1851_v30_: d_1853_v32_})
                d_1817_v0_ = (d_1817_v0_).set(((self).f11 if ((d_1848_v27_)[d_1847_v26_] if (d_1847_v26_) in (d_1848_v27_) else (self).f15) else (self).f11), len(d_1856_v33_))
                (globalState).f6 = default__.safeModuloInt((d_1822_v5_)[default__.safeIndex(168, (d_1822_v5_).length(0))], (self).f11)
                d_1857_v34_: _dafny.Array
                nw282_ = _dafny.Array(D5.default()(), 13)
                d_1857_v34_ = nw282_
                d_1858_v35_: _dafny.Map
                d_1858_v35_ = _dafny.Map({p0: (self).f14})
                d_1859_v36_: D5
                d_1859_v36_ = D5_DC18(d_1858_v35_, (self).f11)
                index320_ = default__.safeIndex(351, (d_1857_v34_).length(0))
                (d_1857_v34_)[index320_] = d_1859_v36_
                d_1860_v37_: _dafny.Map
                d_1860_v37_ = _dafny.Map({not((d_1847_v26_) not in (self.f16)): (self).f11})
                d_1861_v38_: D6
                d_1861_v38_ = D6_DC19(d_1860_v37_)
                index321_ = default__.safeIndex(351, (d_1857_v34_).length(0))
                index322_ = default__.safeIndex(168, (d_1822_v5_).length(0))
                rhs306_ = d_1859_v36_
                rhs307_ = (self).f11
                rhs308_ = (d_1861_v38_).cf37
                lhs250_ = d_1857_v34_
                lhs251_ = default__.safeIndex(351, (d_1857_v34_).length(0))
                lhs252_ = d_1822_v5_
                lhs253_ = default__.safeIndex(168, (d_1822_v5_).length(0))
                lhs250_[lhs251_] = rhs306_
                lhs252_[lhs253_] = rhs307_
                d_1860_v37_ = rhs308_
        elif source29_.is_DC30:
            d_1862___mcc_h0_ = source29_.cf48
            d_1863___mcc_h1_ = source29_.cf49
            d_1864___mcc_h2_ = source29_.cf50
            d_1865___mcc_h3_ = source29_.cf51
            d_1866_cf51_ = d_1865___mcc_h3_
            d_1867_cf50_ = d_1864___mcc_h2_
            d_1868_cf49_ = d_1863___mcc_h1_
            d_1869_cf48_ = d_1862___mcc_h0_
            d_1870_v39_: _dafny.Map
            d_1870_v39_ = _dafny.Map({d_1867_cf50_: d_1868_cf49_})
            d_1871_v40_: D11
            d_1871_v40_ = D11_DC34(not(d_1868_cf49_), d_1869_cf48_, len(d_1870_v39_), d_1868_cf49_, d_1868_cf49_)
            (globalState).f6 = (0) - ((0) - ((d_1871_v40_).cf60))
            source30_ = d_1820_v3_
            if source30_.is_DC48:
                d_1872___mcc_h10_ = source30_.cf90
                d_1873___mcc_h11_ = source30_.cf91
                d_1874___mcc_h12_ = source30_.cf92
                d_1875_cf92_ = d_1874___mcc_h12_
                d_1876_cf91_ = d_1873___mcc_h11_
                d_1877_cf90_ = d_1872___mcc_h10_
                d_1878_v41_: C0
                nw283_ = C0()
                nw283_.ctor__()
                d_1878_v41_ = nw283_
                d_1879_v42_: _dafny.Map
                d_1879_v42_ = _dafny.Map({d_1876_cf91_: (self).f14})
                d_1880_v43_: D0
                d_1880_v43_ = D0_DC1((self).f15, len(d_1879_v42_))
                (globalState).f7 = ((d_1869_cf48_ if d_1866_cf51_ else True) if d_1868_cf49_ else (d_1880_v43_).cf1)
                d_1881_v44_: C0
                nw284_ = C0()
                nw284_.ctor__()
                d_1881_v44_ = nw284_
                d_1882_v45_: _dafny.Seq
                d_1882_v45_ = _dafny.SeqWithoutIsStrInference([d_1821_v4_])
                d_1883_v46_: _dafny.Map
                d_1883_v46_ = _dafny.Map({False: _dafny.SeqWithoutIsStrInference([d_1821_v4_ for d_1884_i5_ in range(default__.abs(-426))])})
                d_1885_v47_: _dafny.Array
                nw285_ = _dafny.Array(None, 22)
                nw285_[int(0)] = d_1882_v45_
                nw285_[int(1)] = (((d_1883_v46_)[d_1869_cf48_] if (d_1869_cf48_) in (d_1883_v46_) else d_1882_v45_)) + (d_1882_v45_)
                nw285_[int(2)] = (d_1882_v45_).set(default__.safeIndex((self).f14, len(d_1882_v45_)), d_1821_v4_)
                nw285_[int(3)] = (d_1882_v45_) + (d_1882_v45_)
                nw285_[int(4)] = d_1882_v45_
                nw285_[int(5)] = d_1882_v45_
                nw285_[int(6)] = d_1882_v45_
                nw285_[int(7)] = default__.fm58(len((self).f13), globalState)
                nw285_[int(8)] = d_1882_v45_
                nw285_[int(9)] = _dafny.SeqWithoutIsStrInference([d_1821_v4_])
                nw285_[int(10)] = d_1882_v45_
                nw285_[int(11)] = d_1882_v45_
                nw285_[int(12)] = (d_1882_v45_) + (_dafny.SeqWithoutIsStrInference([d_1821_v4_]))
                nw285_[int(13)] = _dafny.SeqWithoutIsStrInference([d_1821_v4_ for d_1886_i6_ in range(default__.abs(774))])
                nw285_[int(14)] = d_1882_v45_
                nw285_[int(15)] = (_dafny.SeqWithoutIsStrInference([d_1821_v4_ for d_1887_i7_ in range(default__.abs(194))])) + (default__.fm58(853, globalState))
                nw285_[int(16)] = default__.fm58((self).f11, globalState)
                nw285_[int(17)] = d_1882_v45_
                nw285_[int(18)] = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1869_cf48_]), d_1821_v4_])) + (_dafny.SeqWithoutIsStrInference([d_1821_v4_, d_1821_v4_]))
                nw285_[int(19)] = (((d_1883_v46_)[(self).f15] if ((self).f15) in (d_1883_v46_) else _dafny.SeqWithoutIsStrInference([d_1821_v4_, d_1821_v4_]))) + (_dafny.SeqWithoutIsStrInference([(d_1821_v4_).set(default__.safeIndex((self).f12, len(d_1821_v4_)), d_1868_cf49_) for d_1888_i8_ in range(default__.abs(988))]))
                nw285_[int(20)] = d_1882_v45_
                nw285_[int(21)] = d_1882_v45_
                d_1885_v47_ = nw285_
                index323_ = default__.safeIndex(954, (d_1885_v47_).length(0))
                (d_1885_v47_)[index323_] = _dafny.SeqWithoutIsStrInference([d_1821_v4_])
                index324_ = default__.safeIndex(954, (d_1885_v47_).length(0))
                (d_1885_v47_)[index324_] = ((d_1882_v45_) + (d_1882_v45_)) + (_dafny.SeqWithoutIsStrInference([d_1821_v4_]))
            elif source30_.is_DC49:
                d_1889___mcc_h13_ = source30_.cf93
                d_1890_cf93_ = d_1889___mcc_h13_
                (globalState).f2 = not((self).f15)
                index325_ = default__.safeIndex(493, (p0).length(0))
                (p0)[index325_] = (d_1868_cf49_ if (self).f15 else d_1869_cf48_)
                index326_ = default__.safeIndex(493, (p0).length(0))
                (p0)[index326_] = d_1869_cf48_
                d_1891_v48_: C3
                nw286_ = C3()
                nw286_.ctor__(False, d_1867_cf50_, self.f16, len((self.f16) + (self.f16)), (self).fm1((self).f14, d_1869_cf48_, (self).fm24((d_1823_v6_).cardinality, (self).f15, globalState), globalState), d_1867_cf50_, d_1890_cf93_, (self).f14)
                d_1891_v48_ = nw286_
                d_1823_v6_ = _dafny.MultiSet(default__.fm43(d_1866_cf51_, not(not (d_1866_cf51_) or (False)), globalState))
            elif source30_.is_DC50:
                d_1892___mcc_h14_ = source30_.cf94
                d_1893_cf94_ = d_1892___mcc_h14_
                d_1894_v49_: C1
                nw287_ = C1()
                nw287_.ctor__(p0, (self).f14, (self).f11, (self).f12, (self).f15, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1895_i9_ in range(default__.abs(475))]))
                d_1894_v49_ = nw287_
                d_1822_v5_ = d_1822_v5_
                r0 = (self).f14
                d_1896_v50_: C0
                nw288_ = C0()
                nw288_.ctor__()
                d_1896_v50_ = nw288_
            elif True:
                d_1897___mcc_h15_ = source30_.cf89
                d_1898_cf89_ = d_1897___mcc_h15_
                d_1899_v51_: _dafny.Array
                nw289_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
                d_1899_v51_ = nw289_
                index327_ = default__.safeIndex(188, (d_1899_v51_).length(0))
                (d_1899_v51_)[index327_] = self.f16
                index328_ = default__.safeIndex(269, (d_1822_v5_).length(0))
                (d_1822_v5_)[index328_] = default__.safeDivisionInt((self).f12, (self).f12)
                d_1900_v52_: D4
                d_1900_v52_ = D4_DC13(d_1818_v1_)
                d_1901_v53_: _dafny.Seq
                d_1901_v53_ = _dafny.SeqWithoutIsStrInference([d_1900_v52_])
                d_1902_v54_: str
                d_1902_v54_ = _dafny.CodePoint('w')
                d_1903_v55_: _dafny.Map
                d_1903_v55_ = _dafny.Map({d_1902_v54_: (self).f15})
                d_1904_v56_: _dafny.Map
                d_1904_v56_ = _dafny.Map({((d_1903_v55_)[d_1902_v54_] if (d_1902_v54_) in (d_1903_v55_) else d_1869_cf48_): d_1869_cf48_})
                index329_ = default__.safeIndex(188, (d_1899_v51_).length(0))
                index330_ = default__.safeIndex(269, (d_1822_v5_).length(0))
                rhs309_ = ((self).f11) != ((-580) - ((_dafny.MultiSet([d_1868_cf49_, d_1868_cf49_])).cardinality))
                rhs310_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttxre"))
                rhs311_ = ((self).f12) * ((len((self).fm3((self).f11, d_1866_cf51_, d_1869_cf48_, globalState)) if d_1866_cf51_ else (self).f14))
                rhs312_ = (len((d_1901_v53_).set(default__.safeIndex(len(d_1904_v56_), len(d_1901_v53_)), D4_DC13(d_1818_v1_)))) - ((0) - (len(d_1818_v1_)))
                lhs254_ = globalState
                lhs255_ = d_1899_v51_
                lhs256_ = default__.safeIndex(188, (d_1899_v51_).length(0))
                lhs257_ = globalState
                lhs258_ = d_1822_v5_
                lhs259_ = default__.safeIndex(269, (d_1822_v5_).length(0))
                lhs254_.f2 = rhs309_
                lhs255_[lhs256_] = rhs310_
                lhs257_.f6 = rhs311_
                lhs258_[lhs259_] = rhs312_
                (globalState).f2 = not((self).f15)
                d_1905_v57_: _dafny.Map
                d_1905_v57_ = _dafny.Map({(self).f14: d_1866_cf51_})
                d_1906_v58_: _dafny.MultiSet
                d_1906_v58_ = _dafny.MultiSet([(self).f12])
                rhs313_ = ((d_1905_v57_) | (d_1905_v57_) if (self).f15 else d_1905_v57_)
                rhs314_ = ((d_1905_v57_)[(len(_dafny.SeqWithoutIsStrInference([d_1902_v54_ for d_1907_i10_ in range(default__.abs(434))]))) * ((d_1822_v5_)[default__.safeIndex(269, (d_1822_v5_).length(0))])] if ((len(_dafny.SeqWithoutIsStrInference([d_1902_v54_ for d_1908_i10_ in range(default__.abs(434))]))) * ((d_1822_v5_)[default__.safeIndex(269, (d_1822_v5_).length(0))])) in (d_1905_v57_) else (d_1906_v58_).issubset(d_1906_v58_))
                rhs315_ = (((d_1821_v4_).set(default__.safeIndex((self).f12, len(d_1821_v4_)), d_1869_cf48_)) + (_dafny.SeqWithoutIsStrInference([d_1866_cf51_, not(d_1866_cf51_), d_1869_cf48_]))) + (d_1821_v4_)
                rhs316_ = ((default__.fm0(d_1869_cf48_, globalState)) < ((self).f14) if ((self).f13) <= ((self).f13) else d_1868_cf49_)
                lhs260_ = globalState
                d_1905_v57_ = rhs313_
                lhs260_.f7 = rhs314_
                d_1821_v4_ = rhs315_
                d_1868_cf49_ = rhs316_
                (globalState).f6 = (d_1822_v5_)[default__.safeIndex(269, (d_1822_v5_).length(0))]
            (globalState).f7 = not((self).f15)
            index331_ = default__.safeIndex(45, (d_1822_v5_).length(0))
            (d_1822_v5_)[index331_] = (self).f11
            index332_ = default__.safeIndex(45, (d_1822_v5_).length(0))
            (d_1822_v5_)[index332_] = (self).f12
        elif source29_.is_DC31:
            d_1909___mcc_h4_ = source29_.cf52
            d_1910___mcc_h5_ = source29_.cf53
            d_1911___mcc_h6_ = source29_.cf54
            d_1912___mcc_h7_ = source29_.cf55
            d_1913_cf55_ = d_1912___mcc_h7_
            d_1914_cf54_ = d_1911___mcc_h6_
            d_1915_cf53_ = d_1910___mcc_h5_
            d_1916_cf52_ = d_1909___mcc_h4_
            index333_ = default__.safeIndex(149, (p0).length(0))
            (p0)[index333_] = True
            index334_ = default__.safeIndex(149, (p0).length(0))
            (p0)[index334_] = ((d_1823_v6_).intersection(_dafny.MultiSet([(self).f15]))).issubset(d_1823_v6_)
            d_1917_v59_: _dafny.Map
            d_1917_v59_ = _dafny.Map({(self).f14: not ((self).f15) or (not((self).f15))})
            if ((d_1917_v59_)[(self).fm2(len(d_1821_v4_), False, d_1913_cf55_, False, globalState)] if ((self).fm2(len(d_1821_v4_), False, d_1913_cf55_, False, globalState)) in (d_1917_v59_) else (self).f15):
                index335_ = default__.safeIndex(982, (d_1822_v5_).length(0))
                (d_1822_v5_)[index335_] = (d_1914_cf54_) + ((self).f12)
                index336_ = default__.safeIndex(982, (d_1822_v5_).length(0))
                (d_1822_v5_)[index336_] = (self).f12
                index337_ = default__.safeIndex(982, (d_1822_v5_).length(0))
                (d_1822_v5_)[index337_] = (d_1914_cf54_) * ((0) - ((self).f12))
                d_1918_v60_: _dafny.Set
                d_1918_v60_ = _dafny.Set({(p0)[default__.safeIndex(149, (p0).length(0))]})
                d_1919_v61_: bool
                d_1920_v62_: _dafny.Seq
                out56_: bool
                out57_: _dafny.Seq
                out56_, out57_ = default__.m0((_dafny.MultiSet([(p0)[default__.safeIndex(149, (p0).length(0))]])).isdisjoint(_dafny.MultiSet([(self).f15])), (d_1918_v60_).isdisjoint(d_1918_v60_), d_1913_cf55_, 562, globalState)
                d_1919_v61_ = out56_
                d_1920_v62_ = out57_
                d_1921_v63_: D1
                d_1921_v63_ = D1_DC5(d_1916_cf52_, (self).f15, (p0)[default__.safeIndex(149, (p0).length(0))])
                d_1922_v64_: _dafny.Map
                d_1922_v64_ = _dafny.Map({(self).f15: not((p0)[default__.safeIndex(149, (p0).length(0))])})
                d_1923_v65_: _dafny.Array
                nw290_ = _dafny.Array(None, 27)
                nw290_[int(0)] = (self).f15
                nw290_[int(1)] = (p0)[default__.safeIndex(149, (p0).length(0))]
                nw290_[int(2)] = d_1919_v61_
                nw290_[int(3)] = (p0)[default__.safeIndex(149, (p0).length(0))]
                nw290_[int(4)] = (p0)[default__.safeIndex(149, (p0).length(0))]
                nw290_[int(5)] = (self).f15
                nw290_[int(6)] = True
                nw290_[int(7)] = (p0)[default__.safeIndex(149, (p0).length(0))]
                nw290_[int(8)] = (self).f15
                nw290_[int(9)] = ((d_1917_v59_)[(self).f14] if ((self).f14) in (d_1917_v59_) else (p0)[default__.safeIndex(149, (p0).length(0))])
                nw290_[int(10)] = not(d_1919_v61_)
                nw290_[int(11)] = d_1919_v61_
                nw290_[int(12)] = d_1919_v61_
                nw290_[int(13)] = (p0)[default__.safeIndex(149, (p0).length(0))]
                nw290_[int(14)] = d_1919_v61_
                nw290_[int(15)] = (d_1921_v63_).cf8
                nw290_[int(16)] = True
                nw290_[int(17)] = d_1919_v61_
                nw290_[int(18)] = not(d_1919_v61_)
                nw290_[int(19)] = True
                nw290_[int(20)] = (p0)[default__.safeIndex(149, (p0).length(0))]
                nw290_[int(21)] = (self).f15
                nw290_[int(22)] = (p0)[default__.safeIndex(149, (p0).length(0))]
                nw290_[int(23)] = (p0)[default__.safeIndex(149, (p0).length(0))]
                nw290_[int(24)] = ((d_1922_v64_)[((d_1922_v64_)[(self).f15] if ((self).f15) in (d_1922_v64_) else True)] if (((d_1922_v64_)[(self).f15] if ((self).f15) in (d_1922_v64_) else True)) in (d_1922_v64_) else d_1919_v61_)
                nw290_[int(25)] = (self).f15
                nw290_[int(26)] = d_1919_v61_
                d_1923_v65_ = nw290_
                d_1924_v66_: _dafny.Map
                d_1924_v66_ = _dafny.Map({(self).f15: d_1923_v65_})
                index338_ = default__.safeIndex(149, (p0).length(0))
                (p0)[index338_] = ((d_1924_v66_) | (d_1924_v66_)) != ((d_1924_v66_).set((p0)[default__.safeIndex(149, (p0).length(0))], d_1923_v65_))
                r0 = 634
            elif True:
                (globalState).f6 = d_1916_cf52_
                d_1925_v67_: C0
                nw291_ = C0()
                nw291_.ctor__()
                d_1925_v67_ = nw291_
                index339_ = default__.safeIndex(851, (d_1822_v5_).length(0))
                (d_1822_v5_)[index339_] = d_1913_cf55_
                index340_ = default__.safeIndex(851, (d_1822_v5_).length(0))
                rhs317_ = (self).f11
                rhs318_ = d_1821_v4_
                lhs261_ = d_1822_v5_
                lhs262_ = default__.safeIndex(851, (d_1822_v5_).length(0))
                lhs261_[lhs262_] = rhs317_
                d_1821_v4_ = rhs318_
                d_1926_v68_: C0
                nw292_ = C0()
                nw292_.ctor__()
                d_1926_v68_ = nw292_
                index341_ = default__.safeIndex(149, (p0).length(0))
                (p0)[index341_] = (self).f15
            d_1927_v69_: _dafny.Seq
            d_1927_v69_ = _dafny.SeqWithoutIsStrInference([(d_1817_v0_).set(d_1914_cf54_, 811)])
            d_1928_v70_: _dafny.Set
            d_1928_v70_ = _dafny.Set({(d_1927_v69_)[default__.safeIndex(743, len(d_1927_v69_))], d_1817_v0_})
            rhs319_ = 979
            rhs320_ = not (not((not((p0)[default__.safeIndex(149, (p0).length(0))])) == ((p0)[default__.safeIndex(149, (p0).length(0))]))) or ((self).f15)
            rhs321_ = d_1928_v70_
            rhs322_ = default__.safeDivisionInt(d_1913_cf55_, d_1916_cf52_)
            rhs323_ = (self).f15
            lhs263_ = globalState
            lhs264_ = globalState
            lhs265_ = globalState
            lhs263_.f6 = rhs319_
            lhs264_.f2 = rhs320_
            d_1928_v70_ = rhs321_
            d_1916_cf52_ = rhs322_
            lhs265_.f2 = rhs323_
            d_1915_cf53_ = _dafny.CodePoint('v')
        elif source29_.is_DC28:
            d_1929___mcc_h8_ = source29_.cf47
            d_1930_cf47_ = d_1929___mcc_h8_
            index342_ = default__.safeIndex(979, (d_1822_v5_).length(0))
            (d_1822_v5_)[index342_] = (self).f11
            index343_ = default__.safeIndex(979, (d_1822_v5_).length(0))
            (d_1822_v5_)[index343_] = ((self).f11) - (-148)
            d_1931_v71_: _dafny.Map
            d_1931_v71_ = _dafny.Map({(self).f15: not(not((self).f15))})
            d_1932_v72_: _dafny.Map
            d_1932_v72_ = _dafny.Map({(self).f11: False})
            d_1933_v73_: D0
            d_1933_v73_ = D0_DC1((self).f15, (self).f11)
            d_1931_v71_ = (d_1931_v71_).set((((d_1932_v72_)[(self).f12] if ((self).f12) in (d_1932_v72_) else (self).f15)) and ((d_1933_v73_).cf1), (self).f15)
            d_1934_v74_: _dafny.Map
            d_1934_v74_ = _dafny.Map({p0: (self).f15})
            d_1935_v75_: C3
            nw293_ = C3()
            nw293_.ctor__((self).f15, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "henmwywpg")), self.f16, 390, (self).f15, (self).f13, (self).f14, (self).f14)
            d_1935_v75_ = nw293_
            d_1936_v76_: _dafny.Seq
            d_1936_v76_ = _dafny.SeqWithoutIsStrInference([d_1935_v75_])
            d_1937_v77_: _dafny.Map
            d_1937_v77_ = _dafny.Map({(0) - ((self).f11): d_1936_v76_})
            d_1934_v74_ = (d_1934_v74_).set(p0, ((d_1822_v5_)[default__.safeIndex(979, (d_1822_v5_).length(0))]) > ((0) - (len(d_1937_v77_))))
            d_1938_v78_: _dafny.MultiSet
            d_1938_v78_ = _dafny.MultiSet([_dafny.Set({(self).f12, 580, (self).f11})])
            d_1939_v79_: _dafny.Map
            d_1939_v79_ = _dafny.Map({(d_1821_v4_)[default__.safeIndex((d_1822_v5_)[default__.safeIndex(979, (d_1822_v5_).length(0))], len(d_1821_v4_))]: d_1821_v4_})
            d_1940_v80_: str
            d_1940_v80_ = _dafny.CodePoint('y')
            d_1941_v81_: _dafny.Set
            d_1941_v81_ = _dafny.Set({(self).f11})
            d_1942_v82_: _dafny.Seq
            d_1942_v82_ = _dafny.SeqWithoutIsStrInference([d_1930_cf47_])
            index344_ = default__.safeIndex(979, (d_1822_v5_).length(0))
            rhs324_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1943_i11_ in range(default__.abs(324))])).set(default__.safeIndex(default__.safeDivisionInt((self).f14, len(((d_1939_v79_)[False] if (False) in (d_1939_v79_) else _dafny.SeqWithoutIsStrInference([(self).f15, (d_1935_v75_).f25])))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1944_i11_ in range(default__.abs(324))]))), d_1940_v80_)
            rhs325_ = ((d_1822_v5_)[default__.safeIndex(979, (d_1822_v5_).length(0))]) * (len(self.f16))
            rhs326_ = ((d_1938_v78_) - (_dafny.MultiSet([d_1941_v81_]))).set(d_1941_v81_, default__.abs(len((d_1942_v82_)[default__.safeIndex((self).f14, len(d_1942_v82_))])))
            lhs266_ = self
            lhs267_ = d_1822_v5_
            lhs268_ = default__.safeIndex(979, (d_1822_v5_).length(0))
            lhs266_.f16 = rhs324_
            lhs267_[lhs268_] = rhs325_
            d_1938_v78_ = rhs326_
        elif True:
            d_1945___mcc_h9_ = source29_.cf56
            d_1946_cf56_ = d_1945___mcc_h9_
            d_1822_v5_ = d_1822_v5_
            d_1947_v83_: _dafny.Array
            nw294_ = _dafny.Array(_dafny.Map({}), 15)
            d_1947_v83_ = nw294_
            d_1948_v85_: _dafny.Map
            d_1948_v85_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f15]): (self).f12})
            d_1949_v86_: _dafny.Map
            def iife119_():
                coll51_ = _dafny.Map()
                compr_51_: _dafny.Seq
                for compr_51_ in (d_1948_v85_).keys.Elements:
                    d_1950_v84_: _dafny.Seq = compr_51_
                    if (d_1950_v84_) in (d_1948_v85_):
                        coll51_[d_1950_v84_] = d_1821_v4_
                return _dafny.Map(coll51_)
            d_1949_v86_ = _dafny.Map({(self).f15: iife119_()
            })
            d_1951_v87_: _dafny.Map
            d_1951_v87_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f15, (self).f15, (self).f15]): d_1821_v4_})
            index345_ = default__.safeIndex(974, (d_1947_v83_).length(0))
            (d_1947_v83_)[index345_] = (((d_1949_v86_)[not((self).f15)] if (not((self).f15)) in (d_1949_v86_) else d_1951_v87_)) | (d_1951_v87_)
            index346_ = default__.safeIndex(974, (d_1947_v83_).length(0))
            (d_1947_v83_)[index346_] = d_1951_v87_
            d_1952_v88_: str
            d_1952_v88_ = _dafny.CodePoint('e')
            d_1953_v89_: _dafny.Map
            d_1953_v89_ = _dafny.Map({d_1818_v1_: (self.f16).set(default__.safeIndex((self).f11, len(self.f16)), d_1952_v88_)})
            d_1954_v90_: _dafny.Set
            d_1954_v90_ = _dafny.Set({(d_1821_v4_)[default__.safeIndex((0) - (len(d_1821_v4_)), len(d_1821_v4_))], (self).f15})
            d_1953_v89_ = (d_1953_v89_).set(_dafny.SeqWithoutIsStrInference([93, len(d_1954_v90_), (self).f12]), self.f16)
            d_1955_v91_: C2
            nw295_ = C2()
            nw295_.ctor__((self).f12, self.f16, ((self).f11) + ((self).f12), (self).f11)
            d_1955_v91_ = nw295_
        (globalState).f6 = (self).f11
        (globalState).f6 = (self).f12
        d_1956_i12_: int
        d_1956_i12_ = 0
        with _dafny.label("10"):
            while (len((d_1818_v1_) + (d_1818_v1_))) != ((self).f14):
                with _dafny.c_label("10"):
                    if (d_1956_i12_) >= (100):
                        raise _dafny.Break("10")
                    d_1956_i12_ = (d_1956_i12_) + (1)
                    d_1957_v92_: str
                    d_1957_v92_ = _dafny.CodePoint('d')
                    rhs327_ = ((self).f11) * ((self).f11)
                    rhs328_ = d_1957_v92_
                    lhs269_ = globalState
                    lhs269_.f6 = rhs327_
                    d_1957_v92_ = rhs328_
                    (globalState).f6 = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))) + ((0) - ((self).f12))
                    (globalState).f2 = (self).f15
                    d_1958_v93_: _dafny.MultiSet
                    d_1958_v93_ = _dafny.MultiSet([p0])
                    d_1958_v93_ = d_1958_v93_
                    pass
            pass
        (globalState).f7 = (self).f15
        r0 = (self).f14
        r1 = d_1818_v1_
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        d_1959_v1_: _dafny.Map
        d_1959_v1_ = _dafny.Map({(self).f14: (self).f12})
        d_1960_v2_: D17
        def iife120_():
            coll52_ = _dafny.Map()
            compr_52_: int
            for compr_52_ in _dafny.IntegerRange(286, -19):
                d_1961_v0_: int = compr_52_
                if ((286) <= (d_1961_v0_)) and ((d_1961_v0_) < (-19)):
                    coll52_[(d_1961_v0_) + ((self).f11)] = (self).f14
            return _dafny.Map(coll52_)
        d_1960_v2_ = D17_DC47((iife120_()
) | (d_1959_v1_))
        d_1960_v2_ = d_1960_v2_
        d_1962_i0_: int
        d_1962_i0_ = 0
        with _dafny.label("11"):
            while (self).f15:
                with _dafny.c_label("11"):
                    if (d_1962_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_1962_i0_ = (d_1962_i0_) + (1)
                    d_1963_v3_: C3
                    nw296_ = C3()
                    nw296_.ctor__(((self).f12) == (len(d_1959_v1_)), self.f16, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1964_i1_ in range(default__.abs(-415))])) + (self.f16), (self).f11, False, (self).f13, (self).f12, 855)
                    d_1963_v3_ = nw296_
                    d_1965_v4_: D18
                    d_1965_v4_ = D18_DC52((d_1963_v3_).f25, (d_1963_v3_).f25, (d_1963_v3_).f26)
                    d_1966_v5_: _dafny.Seq
                    d_1966_v5_ = _dafny.SeqWithoutIsStrInference([(d_1963_v3_).f25, (d_1963_v3_).f25, (d_1965_v4_).cf96, (self).f15])
                    (globalState).f7 = (d_1966_v5_)[default__.safeIndex((self).f12, len(d_1966_v5_))]
                    d_1967_v6_: D17
                    d_1967_v6_ = D17_DC48((d_1963_v3_).f25, (self).f15, True)
                    d_1968_v7_: D18
                    d_1968_v7_ = D18_DC52((d_1967_v6_).cf92, (d_1963_v3_).f25, (self).f13)
                    d_1969_v8_: D18
                    d_1969_v8_ = D18_DC54(d_1968_v7_)
                    d_1970_v9_: _dafny.Map
                    d_1970_v9_ = _dafny.Map({len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xgs")) if (self).f15 else (d_1963_v3_).f26)): ((self).f13) + ((d_1963_v3_).f26)})
                    d_1971_v10_: D6
                    d_1971_v10_ = D6_DC20(len(_dafny.SeqWithoutIsStrInference([(self).f11, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lh")))])), (self).f12)
                    rhs329_ = ((self).f12) + ((d_1971_v10_).cf38)
                    rhs330_ = d_1969_v8_
                    rhs331_ = d_1970_v9_
                    lhs270_ = globalState
                    lhs270_.f6 = rhs329_
                    d_1969_v8_ = rhs330_
                    d_1970_v9_ = rhs331_
                    if (d_1963_v3_).f25:
                        d_1972_v11_: _dafny.Array
                        def lambda70_(d_1973_v3_, d_1974_v5_):
                            def lambda71_(d_1975_i2_):
                                return (d_1974_v5_ if (d_1973_v3_).f25 else d_1974_v5_)

                            return lambda71_

                        init41_ = lambda70_(d_1963_v3_, d_1966_v5_)
                        nw297_ = _dafny.Array(None, 3)
                        for i0_41_ in range(nw297_.length(0)):
                            nw297_[i0_41_] = init41_(i0_41_)
                        d_1972_v11_ = nw297_
                        nw298_ = _dafny.Array(_dafny.Seq({}), 2)
                        d_1972_v11_ = nw298_
                        (globalState).f7 = ((self).f12) == ((self).f12)
                        d_1976_v12_: D6
                        d_1976_v12_ = D6_DC21((self).f15)
                        def iife121_(_pat_let34_0):
                            def iife122_(d_1977_dt__update__tmp_h0_):
                                def iife123_(_pat_let35_0):
                                    def iife124_(d_1978_dt__update_hcf40_h0_):
                                        return D6_DC21(d_1978_dt__update_hcf40_h0_)
                                    return iife124_(_pat_let35_0)
                                return iife123_((self).f15)
                            return iife122_(_pat_let34_0)
                        d_1976_v12_ = iife121_(d_1976_v12_)
                        index347_ = default__.safeIndex(15, (p0).length(0))
                        (p0)[index347_] = not ((d_1966_v5_)[default__.safeIndex((self).f11, len(d_1966_v5_))]) or ((d_1963_v3_).f25)
                        index348_ = default__.safeIndex(15, (p0).length(0))
                        (p0)[index348_] = (not(((self).f14) > ((self).f12))) or ((d_1963_v3_).f25)
                        (globalState).f6 = ((self).f11) * ((self).f14)
                    elif True:
                        d_1979_v13_: _dafny.Set
                        d_1979_v13_ = _dafny.Set({p0, (p0 if (self).f15 else p0), p0, p0, p0})
                        d_1979_v13_ = (_dafny.Set({p0})).intersection((d_1979_v13_).intersection(_dafny.Set({p0})))
                        index349_ = default__.safeIndex(392, (p0).length(0))
                        (p0)[index349_] = (self).f15
                        d_1980_v14_: _dafny.Map
                        d_1980_v14_ = _dafny.Map({(d_1963_v3_).f25: (self).f15})
                        d_1981_v15_: _dafny.Set
                        d_1981_v15_ = _dafny.Set({len(d_1980_v14_)})
                        d_1982_v16_: _dafny.Seq
                        d_1982_v16_ = _dafny.SeqWithoutIsStrInference([(self).f11, len(d_1981_v15_), (self).f12, 493])
                        d_1983_v17_: _dafny.Map
                        d_1983_v17_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1984_i3_ in range(default__.abs(561))]): len((d_1982_v16_) + (d_1982_v16_))})
                        index350_ = default__.safeIndex(392, (p0).length(0))
                        rhs332_ = False
                        rhs333_ = not(not((p1).issubset(p1)))
                        rhs334_ = d_1983_v17_
                        lhs271_ = globalState
                        lhs272_ = p0
                        lhs273_ = default__.safeIndex(392, (p0).length(0))
                        lhs271_.f7 = rhs332_
                        lhs272_[lhs273_] = rhs333_
                        d_1983_v17_ = rhs334_
                        d_1985_v18_: _dafny.Array
                        nw299_ = _dafny.Array(None, 22)
                        nw299_[int(0)] = default__.fm30(True, (d_1963_v3_).f25, (self).f11, not((self).f15), globalState)
                        nw299_[int(1)] = (_dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(392, (p0).length(0))]])) + (d_1966_v5_)
                        nw299_[int(2)] = (d_1966_v5_) + (d_1966_v5_)
                        nw299_[int(3)] = d_1966_v5_
                        nw299_[int(4)] = d_1966_v5_
                        nw299_[int(5)] = (d_1966_v5_) + (d_1966_v5_)
                        nw299_[int(6)] = d_1966_v5_
                        nw299_[int(7)] = _dafny.SeqWithoutIsStrInference([not(False), (self).f15])
                        nw299_[int(8)] = d_1966_v5_
                        nw299_[int(9)] = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(392, (p0).length(0))]])
                        nw299_[int(10)] = d_1966_v5_
                        nw299_[int(11)] = d_1966_v5_
                        nw299_[int(12)] = d_1966_v5_
                        nw299_[int(13)] = (d_1966_v5_).set(default__.safeIndex((self).f14, len(d_1966_v5_)), not((self).f15))
                        nw299_[int(14)] = d_1966_v5_
                        nw299_[int(15)] = d_1966_v5_
                        nw299_[int(16)] = d_1966_v5_
                        nw299_[int(17)] = d_1966_v5_
                        nw299_[int(18)] = d_1966_v5_
                        nw299_[int(19)] = d_1966_v5_
                        nw299_[int(20)] = (_dafny.SeqWithoutIsStrInference([(d_1963_v3_).f25])) + (d_1966_v5_)
                        nw299_[int(21)] = d_1966_v5_
                        d_1985_v18_ = nw299_
                        index351_ = default__.safeIndex(60, (d_1985_v18_).length(0))
                        (d_1985_v18_)[index351_] = d_1966_v5_
                        index352_ = default__.safeIndex(60, (d_1985_v18_).length(0))
                        rhs335_ = _dafny.SeqWithoutIsStrInference([(d_1963_v3_).f25])
                        rhs336_ = (p0)[default__.safeIndex(392, (p0).length(0))]
                        lhs274_ = d_1985_v18_
                        lhs275_ = default__.safeIndex(60, (d_1985_v18_).length(0))
                        lhs276_ = globalState
                        lhs274_[lhs275_] = rhs335_
                        lhs276_.f7 = rhs336_
                        d_1980_v14_ = (d_1980_v14_).set((not((p0)[default__.safeIndex(392, (p0).length(0))]) if True else not((p0)[default__.safeIndex(392, (p0).length(0))])), False)
                        d_1986_v19_: _dafny.MultiSet
                        d_1986_v19_ = _dafny.MultiSet([(self).f15, (p0)[default__.safeIndex(392, (p0).length(0))]])
                        d_1986_v19_ = (d_1986_v19_).intersection(_dafny.MultiSet((d_1985_v18_)[default__.safeIndex(60, (d_1985_v18_).length(0))]))
                    pass
            pass
        d_1987_v20_: _dafny.Seq
        d_1987_v20_ = _dafny.SeqWithoutIsStrInference([True, not((self).f15)])
        d_1988_v21_: _dafny.Map
        d_1988_v21_ = _dafny.Map({self.f16: D14_DC40(d_1987_v20_)})
        d_1989_v22_: str
        d_1989_v22_ = _dafny.CodePoint('u')
        d_1990_v23_: D14
        d_1990_v23_ = D14_DC40(d_1987_v20_)
        d_1988_v21_ = (d_1988_v21_).set(((self.f16) + ((self).f13)).set(default__.safeIndex(len(d_1959_v1_), len((self.f16) + ((self).f13))), d_1989_v22_), d_1990_v23_)
        d_1991_v24_: _dafny.Array
        nw300_ = _dafny.Array(D0.default()(), 5)
        d_1991_v24_ = nw300_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1991_v24_).length(0)):
            d_1992_i4_: int = guard_loop_5_
            if (True) and (((0) <= (d_1992_i4_)) and ((d_1992_i4_) < ((d_1991_v24_).length(0)))):
                (d_1991_v24_)[(d_1992_i4_)] = D0_DC0((_dafny.MultiSet([d_1989_v22_])) - (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_1989_v22_ for d_1993_i5_ in range(default__.abs(123))])).set(default__.safeIndex((0) - ((self).f12), len(_dafny.SeqWithoutIsStrInference([d_1989_v22_ for d_1994_i5_ in range(default__.abs(123))]))), d_1989_v22_))))
        d_1995_v25_: _dafny.Array
        nw301_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
        d_1995_v25_ = nw301_
        d_1996_v26_: _dafny.Array
        def lambda72_(d_1997_i6_):
            return (d_1997_i6_) * ((self).f11)

        init42_ = lambda72_
        nw302_ = _dafny.Array(None, 17)
        for i0_42_ in range(nw302_.length(0)):
            nw302_[i0_42_] = init42_(i0_42_)
        d_1996_v26_ = nw302_
        d_1998_v27_: _dafny.Set
        d_1998_v27_ = _dafny.Set({d_1996_v26_})
        d_1999_v28_: _dafny.Map
        d_1999_v28_ = _dafny.Map({d_1995_v25_: d_1998_v27_})
        d_1999_v28_ = (d_1999_v28_).set(d_1995_v25_, _dafny.Set({d_1996_v26_, d_1996_v26_}))
        d_2000_v29_: _dafny.Seq
        d_2000_v29_ = _dafny.SeqWithoutIsStrInference([(self).f12])
        d_2001_i7_: int
        d_2001_i7_ = 0
        with _dafny.label("12"):
            while ((self).f12) in (d_2000_v29_):
                with _dafny.c_label("12"):
                    if (d_2001_i7_) >= (100):
                        raise _dafny.Break("12")
                    d_2001_i7_ = (d_2001_i7_) + (1)
                    if False:
                        d_2002_v30_: _dafny.Array
                        nw303_ = _dafny.Array(False, 1)
                        d_2002_v30_ = nw303_
                        d_2002_v30_ = p0
                        d_2003_v31_: _dafny.MultiSet
                        d_2003_v31_ = _dafny.MultiSet([False])
                        d_2004_v32_: _dafny.Set
                        d_2005_v33_: int
                        d_2006_v34_: bool
                        d_2007_v35_: int
                        out58_: _dafny.Set
                        out59_: int
                        out60_: bool
                        out61_: int
                        out58_, out59_, out60_, out61_ = (self).m16(d_1996_v26_, (self).fm1((self).f11, (self).f15, (self).fm24((self).f14, (self).f15, globalState), globalState), ((self).f13) + ((self).f13), d_2003_v31_, globalState)
                        d_2004_v32_ = out58_
                        d_2005_v33_ = out59_
                        d_2006_v34_ = out60_
                        d_2007_v35_ = out61_
                        d_1996_v26_ = d_1996_v26_
                        (self).f16 = self.f16
                        d_2008_v36_: C3
                        nw304_ = C3()
                        nw304_.ctor__(d_2006_v34_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iksrlv"))) + ((self).f13), (self).f13, d_2007_v35_, (self).f15, self.f16, ((self).f14) * ((self).f12), (self).f14)
                        d_2008_v36_ = nw304_
                    elif True:
                        d_2009_v37_: C3
                        nw305_ = C3()
                        nw305_.ctor__((d_1987_v20_) < (d_1987_v20_), self.f16, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "guxkleia")), ((self).f12) - ((self).f11), (self).f15, self.f16, (self).f14, (self).f14)
                        d_2009_v37_ = nw305_
                        d_1987_v20_ = d_1987_v20_
                        (globalState).f6 = default__.safeDivisionInt((self).f11, (self).f14)
                        d_2010_v38_: D17
                        d_2010_v38_ = D17_DC48(True, (self).f15, False)
                        d_2011_v39_: _dafny.Map
                        d_2011_v39_ = _dafny.Map({d_2010_v38_: (self).f14})
                        (globalState).f6 = ((d_2011_v39_)[d_2010_v38_] if (d_2010_v38_) in (d_2011_v39_) else len(((self).f13) + ((d_2009_v37_).f26)))
                        d_2012_v40_: _dafny.Map
                        d_2012_v40_ = _dafny.Map({(self).f15: (d_2009_v37_).f25})
                        d_2013_v41_: _dafny.Map
                        d_2013_v41_ = _dafny.Map({(self).f15: d_2012_v40_})
                        d_2014_v42_: T0
                        nw306_ = C1()
                        nw306_.ctor__(p0, (self).f12, (self).f14, (self).f14, (d_2009_v37_).f25, _dafny.SeqWithoutIsStrInference([d_1989_v22_ for d_2015_i8_ in range(default__.abs(113))]))
                        d_2014_v42_ = nw306_
                        d_2016_v43_: _dafny.Map
                        d_2016_v43_ = _dafny.Map({((d_2013_v41_)[(d_2009_v37_).fm1(len(_dafny.SeqWithoutIsStrInference([(d_2009_v37_).f25, (self).f15])), (d_2009_v37_).f25, (self).f15, globalState)] if ((d_2009_v37_).fm1(len(_dafny.SeqWithoutIsStrInference([(d_2009_v37_).f25, (self).f15])), (d_2009_v37_).f25, (self).f15, globalState)) in (d_2013_v41_) else d_2012_v40_): d_2014_v42_})
                        d_2016_v43_ = (d_2016_v43_).set(d_2012_v40_, d_2014_v42_)
                    d_2017_v44_: _dafny.Map
                    d_2017_v44_ = _dafny.Map({True: _dafny.Map({(self).f15: d_1989_v22_})})
                    d_2017_v44_ = d_2017_v44_
                    index353_ = default__.safeIndex(585, (p0).length(0))
                    (p0)[index353_] = (self).f15
                    index354_ = default__.safeIndex(585, (p0).length(0))
                    (p0)[index354_] = (self).fm24(347, (self).f15, globalState)
                    index355_ = default__.safeIndex(585, (p0).length(0))
                    rhs337_ = (p0)[default__.safeIndex(585, (p0).length(0))]
                    rhs338_ = (p0)[default__.safeIndex(585, (p0).length(0))]
                    lhs277_ = globalState
                    lhs278_ = p0
                    lhs279_ = default__.safeIndex(585, (p0).length(0))
                    lhs277_.f7 = rhs337_
                    lhs278_[lhs279_] = rhs338_
                    pass
            pass
        d_2018_v45_: D2
        d_2018_v45_ = D2_DC8((self).f15, (self).f15, (self).f13)
        d_2019_v46_: _dafny.MultiSet
        d_2019_v46_ = _dafny.MultiSet([d_1989_v22_, d_1989_v22_])
        d_2020_v47_: D0
        d_2020_v47_ = D0_DC0(_dafny.MultiSet([d_1989_v22_, d_1989_v22_, _dafny.CodePoint('o')]))
        d_2021_v49_: D3
        def iife125_():
            coll53_ = _dafny.Set()
            compr_53_: D0
            for compr_53_ in (_dafny.MultiSet([D0_DC0(d_2019_v46_), d_2020_v47_])).Elements:
                d_2022_v48_: D0 = compr_53_
                if (d_2022_v48_) in (_dafny.MultiSet([D0_DC0(d_2019_v46_), d_2020_v47_])):
                    coll53_ = coll53_.union(_dafny.Set([d_2022_v48_]))
            return _dafny.Set(coll53_)
        d_2021_v49_ = D3_DC11((self).f14, d_2018_v45_, iife125_()
, (self).f15)
        d_2023_v50_: _dafny.Set
        d_2023_v50_ = _dafny.Set({d_2020_v47_})
        pat_let_tv34_ = d_2023_v50_
        pat_let_tv35_ = d_2018_v45_
        def iife126_(_pat_let36_0):
            def iife127_(d_2024_dt__update__tmp_h1_):
                def iife128_(_pat_let37_0):
                    def iife129_(d_2025_dt__update_hcf26_h0_):
                        def iife130_(_pat_let38_0):
                            def iife131_(d_2026_dt__update_hcf25_h0_):
                                return D3_DC11((d_2024_dt__update__tmp_h1_).cf24, d_2026_dt__update_hcf25_h0_, d_2025_dt__update_hcf26_h0_, (d_2024_dt__update__tmp_h1_).cf27)
                            return iife131_(_pat_let38_0)
                        return iife130_(pat_let_tv35_)
                    return iife129_(_pat_let37_0)
                return iife128_(pat_let_tv34_)
            return iife127_(_pat_let36_0)
        r0 = (iife126_(d_2021_v49_)).cf27
        r1 = p1
        return r0, r1

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        if p2:
            d_2027_v0_: C2
            nw307_ = C2()
            nw307_.ctor__((self).f14, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnqct"))) + ((self).f13), p3, p3)
            d_2027_v0_ = nw307_
            d_2028_v1_: _dafny.Array
            nw308_ = _dafny.Array(None, 9)
            nw308_[int(0)] = True
            nw308_[int(1)] = (self).f15
            nw308_[int(2)] = p2
            nw308_[int(3)] = p2
            nw308_[int(4)] = p2
            nw308_[int(5)] = p2
            nw308_[int(6)] = p2
            nw308_[int(7)] = (self).f15
            nw308_[int(8)] = (self).f15
            d_2028_v1_ = nw308_
            d_2029_v2_: _dafny.Seq
            d_2029_v2_ = _dafny.SeqWithoutIsStrInference([d_2028_v1_, d_2028_v1_, d_2028_v1_])
            (d_2027_v0_).f27 = default__.safeModuloInt(len((d_2029_v2_).set(default__.safeIndex((0) - (p1), len(d_2029_v2_)), d_2028_v1_)), (self).f14)
            d_2030_v3_: str
            d_2030_v3_ = _dafny.CodePoint('m')
            d_2031_v4_: _dafny.Set
            d_2031_v4_ = _dafny.Set({d_2030_v3_})
            d_2032_v5_: _dafny.Seq
            d_2032_v5_ = _dafny.SeqWithoutIsStrInference([p1, (self).f11, (self).f12, len(d_2031_v4_)])
            r0 = ((d_2032_v5_)[default__.safeIndex(p3, len(d_2032_v5_))]) > ((d_2027_v0_.f27 if False else (self).f12))
            (globalState).f7 = p2
            (globalState).f2 = not(p2)
        elif True:
            d_2033_v6_: _dafny.Seq
            d_2033_v6_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            d_2034_v7_: _dafny.Seq
            d_2034_v7_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p2, globalState)])
            (globalState).f2 = (self).fm1(len((_dafny.SeqWithoutIsStrInference([p3, len(d_2033_v6_)])) + (d_2034_v7_)), (False) not in (d_2033_v6_), p2, globalState)
            (globalState).f2 = (self).f15
            d_2035_v8_: _dafny.Array
            nw309_ = _dafny.Array(False, 15)
            d_2035_v8_ = nw309_
            d_2036_v9_: C1
            nw310_ = C1()
            nw310_.ctor__(d_2035_v8_, ((self).fm2((self).f12, False, p1, p2, globalState)) * ((self).f12), p1, p3, (self).f15, p0)
            d_2036_v9_ = nw310_
            (globalState).f7 = (d_2033_v6_)[default__.safeIndex((self).f14, len(d_2033_v6_))]
            d_2037_v10_: _dafny.Map
            d_2037_v10_ = _dafny.Map({p2: d_2035_v8_})
            d_2037_v10_ = d_2037_v10_
        if True:
            d_2038_v11_: _dafny.Array
            nw311_ = _dafny.Array(False, 19)
            d_2038_v11_ = nw311_
            d_2039_v12_: _dafny.Map
            d_2039_v12_ = _dafny.Map({p2: p2})
            d_2040_v13_: _dafny.Set
            d_2040_v13_ = _dafny.Set({len(d_2039_v12_)})
            d_2041_v14_: C1
            nw312_ = C1()
            nw312_.ctor__(d_2038_v11_, p3, len(d_2040_v13_), p1, (self).fm24((self).f11, (self).f15, globalState), p0)
            d_2041_v14_ = nw312_
            d_2042_v15_: _dafny.Seq
            d_2042_v15_ = _dafny.SeqWithoutIsStrInference([d_2041_v14_])
            (globalState).f6 = (((0) - (len(_dafny.SeqWithoutIsStrInference([d_2042_v15_, d_2042_v15_, d_2042_v15_, d_2042_v15_])))) - ((self).f12)) + ((0) - (p3))
            d_2043_v16_: _dafny.Set
            d_2043_v16_ = _dafny.Set({not(not((self).f15)), not(p2)})
            d_2044_v17_: _dafny.Map
            d_2044_v17_ = _dafny.Map({d_2043_v16_: p2})
            rhs339_ = (self).fm1((self).f14, ((self).f13) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgbxbuka"))), (self).f15, globalState)
            rhs340_ = (len(self.f16)) * (len(d_2043_v16_))
            rhs341_ = len((((d_2044_v17_).set(d_2043_v16_, p2)) | (d_2044_v17_)).set(d_2043_v16_, False))
            rhs342_ = (p3) + ((self).f14)
            rhs343_ = p3
            lhs280_ = globalState
            lhs281_ = globalState
            lhs280_.f2 = rhs339_
            r1 = rhs340_
            lhs281_.f6 = rhs341_
            r2 = rhs342_
            r1 = rhs343_
            d_2045_v18_: _dafny.Map
            d_2045_v18_ = _dafny.Map({d_2038_v11_: p2})
            d_2046_v19_: _dafny.Map
            d_2046_v19_ = _dafny.Map({len(d_2045_v18_): p2})
            d_2046_v19_ = (d_2046_v19_).set((self).f11, p2)
            d_2047_v20_: _dafny.Array
            nw313_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_2047_v20_ = nw313_
            index356_ = default__.safeIndex(939, (d_2047_v20_).length(0))
            (d_2047_v20_)[index356_] = d_2038_v11_
            index357_ = default__.safeIndex(939, (d_2047_v20_).length(0))
            (d_2047_v20_)[index357_] = d_2041_v14_.f24
            d_2048_v21_: str
            d_2048_v21_ = _dafny.CodePoint('x')
            d_2049_v22_: _dafny.Map
            d_2049_v22_ = _dafny.Map({d_2048_v21_: (self).f15})
            r0 = ((_dafny.Map({d_2048_v21_: p2})) | (d_2049_v22_)) != (d_2049_v22_)
        elif True:
            d_2050_v23_: _dafny.Map
            d_2050_v23_ = _dafny.Map({(p0) + (self.f16): (0) - (default__.safeDivisionInt(p1, (self).f14))})
            d_2050_v23_ = d_2050_v23_
            d_2051_v24_: _dafny.Set
            d_2051_v24_ = _dafny.Set({(self).f14})
            (globalState).f7 = (self).fm24((self).f14, (310) not in (d_2051_v24_), globalState)
            if (p1) < (len(self.f16)):
                d_2052_v25_: _dafny.Array
                nw314_ = _dafny.Array(int(0), 10)
                d_2052_v25_ = nw314_
                index358_ = default__.safeIndex(532, (d_2052_v25_).length(0))
                (d_2052_v25_)[index358_] = (self).f11
                index359_ = default__.safeIndex(532, (d_2052_v25_).length(0))
                (d_2052_v25_)[index359_] = default__.fm0(p2, globalState)
                d_2053_v26_: _dafny.Array
                nw315_ = _dafny.Array(False, 18)
                d_2053_v26_ = nw315_
                index360_ = default__.safeIndex(40, (d_2053_v26_).length(0))
                (d_2053_v26_)[index360_] = True
                index361_ = default__.safeIndex(40, (d_2053_v26_).length(0))
                (d_2053_v26_)[index361_] = p2
                d_2054_v27_: C2
                nw316_ = C2()
                nw316_.ctor__((self).f12, self.f16, p3, (d_2052_v25_)[default__.safeIndex(532, (d_2052_v25_).length(0))])
                d_2054_v27_ = nw316_
                d_2055_v28_: _dafny.MultiSet
                d_2055_v28_ = _dafny.MultiSet([d_2054_v27_, d_2054_v27_, d_2054_v27_, d_2054_v27_])
                d_2056_v29_: _dafny.Seq
                d_2056_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p3: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2057_i0_ in range(default__.abs(430))]))})])
                d_2058_v30_: D2
                d_2058_v30_ = D2_DC8((self).f15, not(False), self.f16)
                d_2059_v31_: D0
                d_2059_v31_ = D0_DC0(_dafny.MultiSet([_dafny.CodePoint('u')]))
                d_2060_v32_: _dafny.Set
                d_2060_v32_ = _dafny.Set({d_2059_v31_, d_2059_v31_})
                index362_ = default__.safeIndex(40, (d_2053_v26_).length(0))
                rhs344_ = (self).f15
                rhs345_ = ((self).f11) * ((len(_dafny.Map({(self).f12: (self).f15})) if p2 else (D3_DC11(((d_2055_v28_)[d_2054_v27_] if (d_2054_v27_) in (d_2055_v28_) else len(d_2056_v29_)), d_2058_v30_, d_2060_v32_, (self).f15)).cf24))
                rhs346_ = (self).f11
                rhs347_ = (self).f15
                lhs282_ = globalState
                lhs283_ = globalState
                lhs284_ = d_2053_v26_
                lhs285_ = default__.safeIndex(40, (d_2053_v26_).length(0))
                lhs282_.f7 = rhs344_
                r1 = rhs345_
                lhs283_.f6 = rhs346_
                lhs284_[lhs285_] = rhs347_
                d_2061_v33_: _dafny.Seq
                d_2061_v33_ = _dafny.SeqWithoutIsStrInference([self.f16])
                d_2062_v34_: _dafny.Map
                d_2062_v34_ = _dafny.Map({d_2054_v27_.f27: (self).f12})
                (self).f16 = (d_2061_v33_)[default__.safeIndex(len((d_2062_v34_ if p2 else _dafny.Map({-234: d_2054_v27_.f27}))), len(d_2061_v33_))]
                r0 = not((self).fm24(d_2054_v27_.f27, (d_2053_v26_)[default__.safeIndex(40, (d_2053_v26_).length(0))], globalState))
            elif True:
                d_2063_v35_: _dafny.Seq
                d_2063_v35_ = _dafny.SeqWithoutIsStrInference([p2])
                d_2064_v36_: D8
                d_2064_v36_ = D8_DC24(_dafny.MultiSet(d_2063_v35_))
                d_2065_v37_: _dafny.Map
                d_2065_v37_ = _dafny.Map({d_2064_v36_: p2})
                d_2065_v37_ = d_2065_v37_
                d_2066_v38_: D4
                d_2066_v38_ = D4_DC14(_dafny.CodePoint('j'), (self).f12)
                d_2067_v39_: _dafny.Set
                d_2067_v39_ = _dafny.Set({False})
                d_2068_v40_: D8
                d_2068_v40_ = D8_DC25(p3)
                d_2069_v41_: _dafny.Map
                d_2069_v41_ = _dafny.Map({p2: (self).f11})
                d_2070_v42_: _dafny.Map
                d_2070_v42_ = _dafny.Map({470: True})
                rhs348_ = (((self).f15) in (d_2067_v39_)) or ((len(self.f16)) <= (p1))
                rhs349_ = default__.fm59((_dafny.Map({p2: (d_2068_v40_).cf44}) if p2 else d_2069_v41_), (self).f14, not (p2) or (p2), globalState)
                rhs350_ = ((default__.fm0(((d_2070_v42_)[(self).f11] if ((self).f11) in (d_2070_v42_) else (self).f15), globalState)) * (p1)) - (p3)
                rhs351_ = (self).fm1(p3, (self).f15, (self).f15, globalState)
                lhs286_ = globalState
                lhs287_ = globalState
                lhs286_.f2 = rhs348_
                d_2066_v38_ = rhs349_
                r2 = rhs350_
                lhs287_.f7 = rhs351_
                d_2071_v43_: _dafny.Map
                d_2071_v43_ = _dafny.Map({(self).f14: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yg"))})
                d_2071_v43_ = d_2071_v43_
                d_2072_v44_: _dafny.Map
                d_2072_v44_ = _dafny.Map({(self).f12: (self).f14})
                d_2073_v45_: _dafny.Map
                d_2073_v45_ = _dafny.Map({True: d_2072_v44_})
                d_2074_v46_: _dafny.Array
                nw317_ = _dafny.Array(None, 9)
                nw317_[int(0)] = d_2072_v44_
                nw317_[int(1)] = d_2072_v44_
                nw317_[int(2)] = d_2072_v44_
                nw317_[int(3)] = default__.fm60(globalState)
                nw317_[int(4)] = d_2072_v44_
                nw317_[int(5)] = ((d_2073_v45_)[False] if (False) in (d_2073_v45_) else d_2072_v44_)
                nw317_[int(6)] = d_2072_v44_
                nw317_[int(7)] = d_2072_v44_
                nw317_[int(8)] = d_2072_v44_
                d_2074_v46_ = nw317_
                d_2075_v47_: _dafny.MultiSet
                d_2075_v47_ = _dafny.MultiSet([(self).f14])
                d_2076_v48_: _dafny.Map
                d_2076_v48_ = _dafny.Map({(d_2075_v47_).cardinality: d_2072_v44_})
                index363_ = default__.safeIndex(481, (d_2074_v46_).length(0))
                (d_2074_v46_)[index363_] = ((d_2076_v48_)[(self).f12] if ((self).f12) in (d_2076_v48_) else d_2072_v44_)
                index364_ = default__.safeIndex(481, (d_2074_v46_).length(0))
                (d_2074_v46_)[index364_] = ((_dafny.Map({default__.fm0((self).f15, globalState): (0) - (p3)})) | (d_2072_v44_)) | ((d_2072_v44_ if p2 else d_2072_v44_))
                d_2077_v49_: _dafny.Array
                def lambda73_(d_2078_i1_):
                    return (868) >= (570)

                init43_ = lambda73_
                nw318_ = _dafny.Array(None, 10)
                for i0_43_ in range(nw318_.length(0)):
                    nw318_[i0_43_] = init43_(i0_43_)
                d_2077_v49_ = nw318_
                d_2077_v49_ = d_2077_v49_
            (globalState).f2 = ((self).f15) or (not (p2) or ((self).f15))
            d_2079_v50_: _dafny.Map
            d_2079_v50_ = _dafny.Map({-871: (0) - (len(self.f16))})
            d_2079_v50_ = ((d_2079_v50_) | (d_2079_v50_)) | (d_2079_v50_)
        d_2080_v51_: _dafny.Seq
        d_2080_v51_ = _dafny.SeqWithoutIsStrInference([False, (self).fm1(len(_dafny.SeqWithoutIsStrInference([_dafny.Set({p3}) for d_2081_i2_ in range(default__.abs(99))])), False, p2, globalState)])
        d_2080_v51_ = d_2080_v51_
        d_2082_v52_: str
        d_2082_v52_ = _dafny.CodePoint('p')
        d_2082_v52_ = d_2082_v52_
        d_2083_v53_: D10
        d_2083_v53_ = D10_DC31((self).f12, d_2082_v52_, (self).f11, 330)
        d_2084_i3_: int
        d_2084_i3_ = 0
        with _dafny.label("13"):
            pat_let_tv36_ = d_2082_v52_
            pat_let_tv37_ = d_2082_v52_
            def lambda77_(source31_):
                if source31_.is_DC29:
                    return True
                elif source31_.is_DC30:
                    d_2111___mcc_h0_ = source31_.cf48
                    d_2112___mcc_h1_ = source31_.cf49
                    d_2113___mcc_h2_ = source31_.cf50
                    d_2114___mcc_h3_ = source31_.cf51
                    d_2115_cf51_ = d_2114___mcc_h3_
                    d_2116_cf50_ = d_2113___mcc_h2_
                    d_2117_cf49_ = d_2112___mcc_h1_
                    d_2118_cf48_ = d_2111___mcc_h0_
                    return (_dafny.Map({pat_let_tv36_: (self).f15})) == (_dafny.Map({pat_let_tv37_: d_2117_cf49_}))
                elif source31_.is_DC31:
                    d_2119___mcc_h4_ = source31_.cf52
                    d_2120___mcc_h5_ = source31_.cf53
                    d_2121___mcc_h6_ = source31_.cf54
                    d_2122___mcc_h7_ = source31_.cf55
                    d_2123_cf55_ = d_2122___mcc_h7_
                    d_2124_cf54_ = d_2121___mcc_h6_
                    d_2125_cf53_ = d_2120___mcc_h5_
                    d_2126_cf52_ = d_2119___mcc_h4_
                    return (self).f15
                elif source31_.is_DC28:
                    d_2127___mcc_h8_ = source31_.cf47
                    d_2128_cf47_ = d_2127___mcc_h8_
                    return (self).f15
                elif True:
                    d_2129___mcc_h9_ = source31_.cf56
                    d_2130_cf56_ = d_2129___mcc_h9_
                    return (self).f15

            while lambda77_(d_2083_v53_):
                with _dafny.c_label("13"):
                    if (d_2084_i3_) >= (100):
                        raise _dafny.Break("13")
                    d_2084_i3_ = (d_2084_i3_) + (1)
                    d_2085_v54_: _dafny.Array
                    def lambda74_(d_2086_i4_):
                        return (self).f15

                    init44_ = lambda74_
                    nw319_ = _dafny.Array(None, 5)
                    for i0_44_ in range(nw319_.length(0)):
                        nw319_[i0_44_] = init44_(i0_44_)
                    d_2085_v54_ = nw319_
                    d_2087_v55_: _dafny.Set
                    d_2087_v55_ = _dafny.Set({d_2085_v54_})
                    d_2088_v56_: D19
                    d_2088_v56_ = D19_DC55(d_2087_v55_)
                    (globalState).f2 = (((d_2088_v56_).cf104) | (d_2087_v55_)).ispropersubset((d_2087_v55_) | (d_2087_v55_))
                    d_2089_v57_: _dafny.MultiSet
                    d_2089_v57_ = _dafny.MultiSet([d_2082_v52_])
                    d_2090_v58_: D0
                    d_2090_v58_ = D0_DC0(d_2089_v57_)
                    rhs352_ = default__.safeModuloInt((self).f14, (self).f11)
                    rhs353_ = (self.f16) + (self.f16)
                    rhs354_ = (d_2082_v52_) not in (default__.fm31(d_2082_v52_, d_2090_v58_, globalState))
                    rhs355_ = (self).f12
                    lhs288_ = self
                    lhs289_ = globalState
                    lhs290_ = globalState
                    r1 = rhs352_
                    lhs288_.f16 = rhs353_
                    lhs289_.f7 = rhs354_
                    lhs290_.f6 = rhs355_
                    d_2091_v59_: C3
                    nw320_ = C3()
                    nw320_.ctor__(p2, default__.fm44(len(self.f16), (self).f15, globalState), p0, len(d_2080_v51_), p2, default__.fm31(d_2082_v52_, d_2090_v58_, globalState), (self).f11, len((self).f13))
                    d_2091_v59_ = nw320_
                    d_2092_v60_: _dafny.Array
                    nw321_ = _dafny.Array(None, 12)
                    nw321_[int(0)] = d_2091_v59_
                    nw321_[int(1)] = d_2091_v59_
                    nw321_[int(2)] = d_2091_v59_
                    nw321_[int(3)] = d_2091_v59_
                    nw321_[int(4)] = d_2091_v59_
                    nw321_[int(5)] = d_2091_v59_
                    nw321_[int(6)] = d_2091_v59_
                    nw321_[int(7)] = d_2091_v59_
                    nw321_[int(8)] = d_2091_v59_
                    nw321_[int(9)] = d_2091_v59_
                    nw321_[int(10)] = d_2091_v59_
                    nw321_[int(11)] = d_2091_v59_
                    d_2092_v60_ = nw321_
                    index365_ = default__.safeIndex(937, (d_2092_v60_).length(0))
                    (d_2092_v60_)[index365_] = d_2091_v59_
                    index366_ = default__.safeIndex(937, (d_2092_v60_).length(0))
                    (d_2092_v60_)[index366_] = d_2091_v59_
                    if ((d_2091_v59_).fm1((self).f14, (self).f15, False, globalState)) == (True):
                        d_2082_v52_ = _dafny.CodePoint('l')
                        r2 = p3
                        (self).f16 = ((self).f13) + ((default__.fm31(d_2082_v52_, d_2090_v58_, globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ingo"))))
                        d_2093_v61_: _dafny.Map
                        d_2093_v61_ = _dafny.Map({p1: (self).fm1(len((d_2091_v59_).f26), (d_2091_v59_).f25, p2, globalState)})
                        d_2094_v62_: _dafny.Map
                        d_2094_v62_ = _dafny.Map({(d_2080_v51_)[default__.safeIndex(len(d_2093_v61_), len(d_2080_v51_))]: p1})
                        d_2095_v63_: _dafny.MultiSet
                        d_2095_v63_ = _dafny.MultiSet([((d_2094_v62_)[(d_2091_v59_).f25] if ((d_2091_v59_).f25) in (d_2094_v62_) else default__.fm0((d_2091_v59_).f25, globalState)), (self).f14])
                        (globalState).f7 = (d_2080_v51_)[default__.safeIndex(((default__.fm33(globalState)) | (d_2095_v63_)).cardinality, len(d_2080_v51_))]
                        d_2096_v64_: bool
                        d_2097_v65_: _dafny.MultiSet
                        out62_: bool
                        out63_: _dafny.MultiSet
                        out62_, out63_ = (d_2091_v59_).m2(d_2085_v54_, (d_2095_v63_).intersection(d_2095_v63_), globalState)
                        d_2096_v64_ = out62_
                        d_2097_v65_ = out63_
                    elif True:
                        d_2098_v66_: D10
                        d_2098_v66_ = D10_DC30((d_2091_v59_).f25, p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ken")), not((d_2091_v59_).f25))
                        d_2099_v67_: _dafny.Map
                        d_2099_v67_ = _dafny.Map({(self).f15: 51})
                        d_2100_v68_: _dafny.Seq
                        d_2100_v68_ = _dafny.SeqWithoutIsStrInference([len((d_2091_v59_).f26)])
                        d_2101_v69_: _dafny.Seq
                        d_2101_v69_ = _dafny.SeqWithoutIsStrInference([p1, len(d_2100_v68_)])
                        d_2102_v70_: _dafny.Array
                        nw322_ = _dafny.Array(None, 11)
                        nw322_[int(0)] = p3
                        nw322_[int(1)] = (self).f14
                        nw322_[int(2)] = len(_dafny.SeqWithoutIsStrInference([(d_2091_v59_).f25, (d_2098_v66_).cf49]))
                        nw322_[int(3)] = 617
                        nw322_[int(4)] = len(d_2099_v67_)
                        nw322_[int(5)] = default__.safeModuloInt((self).f14, p1)
                        nw322_[int(6)] = (p3 if (d_2091_v59_).f25 else default__.fm0((self).f15, globalState))
                        nw322_[int(7)] = (d_2101_v69_)[default__.safeIndex(696, len(d_2101_v69_))]
                        nw322_[int(8)] = default__.fm0((d_2091_v59_).f25, globalState)
                        nw322_[int(9)] = default__.safeModuloInt((self).f12, p3)
                        nw322_[int(10)] = default__.safeDivisionInt((0) - ((self).f14), (self).f11)
                        d_2102_v70_ = nw322_
                        index367_ = default__.safeIndex(774, (d_2102_v70_).length(0))
                        (d_2102_v70_)[index367_] = ((self).f11) + (p1)
                        index368_ = default__.safeIndex(774, (d_2102_v70_).length(0))
                        (d_2102_v70_)[index368_] = (((self).f14 if not((d_2091_v59_).f25) else p3)) + (280)
                        d_2103_v71_: _dafny.Set
                        d_2103_v71_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udycj")))})
                        index369_ = default__.safeIndex(101, (d_2085_v54_).length(0))
                        def iife132_():
                            coll54_ = _dafny.Set()
                            compr_54_: int
                            for compr_54_ in (d_2100_v68_).Elements:
                                d_2104_v72_: int = compr_54_
                                if (d_2104_v72_) in (d_2100_v68_):
                                    coll54_ = coll54_.union(_dafny.Set([default__.safeDivisionInt(d_2104_v72_, len(_dafny.SeqWithoutIsStrInference([True, False])))]))
                            return _dafny.Set(coll54_)
                        (d_2085_v54_)[index369_] = (iife132_()
                        ).issubset(d_2103_v71_)
                        index370_ = default__.safeIndex(101, (d_2085_v54_).length(0))
                        (d_2085_v54_)[index370_] = (d_2091_v59_).f25
                        index371_ = default__.safeIndex(774, (d_2102_v70_).length(0))
                        (d_2102_v70_)[index371_] = 798
                        (globalState).f6 = (self).f11
                        d_2105_v73_: _dafny.Array
                        def lambda75_(d_2106_p1_):
                            def lambda76_(d_2107_i5_):
                                return _dafny.MultiSet([d_2106_p1_])

                            return lambda76_

                        init45_ = lambda75_(p1)
                        nw323_ = _dafny.Array(None, 4)
                        for i0_45_ in range(nw323_.length(0)):
                            nw323_[i0_45_] = init45_(i0_45_)
                        d_2105_v73_ = nw323_
                        d_2108_v74_: D18
                        d_2108_v74_ = D18_DC54(D18_DC51(d_2105_v73_))
                        d_2109_v75_: D18
                        d_2109_v75_ = D18_DC54(d_2108_v74_)
                        d_2110_v76_: _dafny.Map
                        d_2110_v76_ = _dafny.Map({p1: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2109_v75_]))})
                        r1 = (default__.safeModuloInt(646, p1)) * (len((d_2110_v76_) | (d_2110_v76_)))
                    pass
            pass
        d_2131_i6_: int
        d_2131_i6_ = 0
        with _dafny.label("14"):
            while not(not((len(p0)) < (p3))):
                with _dafny.c_label("14"):
                    if (d_2131_i6_) >= (100):
                        raise _dafny.Break("14")
                    d_2131_i6_ = (d_2131_i6_) + (1)
                    d_2132_v77_: _dafny.Array
                    def lambda78_(d_2133_p3_):
                        def lambda79_(d_2134_i7_):
                            return D6_DC22(D6_DC20((self).f11, d_2133_p3_))

                        return lambda79_

                    init46_ = lambda78_(p3)
                    nw324_ = _dafny.Array(None, 12)
                    for i0_46_ in range(nw324_.length(0)):
                        nw324_[i0_46_] = init46_(i0_46_)
                    d_2132_v77_ = nw324_
                    d_2135_v78_: _dafny.Array
                    nw325_ = _dafny.Array(False, 13)
                    d_2135_v78_ = nw325_
                    d_2136_v79_: _dafny.Map
                    d_2136_v79_ = _dafny.Map({(self).f11: p2})
                    d_2137_v80_: _dafny.Map
                    d_2137_v80_ = _dafny.Map({d_2135_v78_: d_2136_v79_})
                    d_2138_v81_: _dafny.Map
                    d_2138_v81_ = _dafny.Map({d_2132_v77_: (d_2137_v80_).set(d_2135_v78_, d_2136_v79_)})
                    d_2138_v81_ = (d_2138_v81_).set(d_2132_v77_, (d_2137_v80_) | (d_2137_v80_))
                    rhs356_ = d_2135_v78_
                    rhs357_ = (default__.safeModuloInt((self).f11, p1)) >= ((self).f14)
                    rhs358_ = (p0) + (default__.fm44(p3, p2, globalState))
                    lhs291_ = globalState
                    lhs292_ = self
                    d_2135_v78_ = rhs356_
                    lhs291_.f2 = rhs357_
                    lhs292_.f16 = rhs358_
                    d_2139_v82_: _dafny.Map
                    d_2139_v82_ = _dafny.Map({p2: len(d_2080_v51_)})
                    d_2140_v83_: _dafny.Array
                    nw326_ = _dafny.Array(None, 6)
                    nw326_[int(0)] = (self).f11
                    nw326_[int(1)] = len(d_2139_v82_)
                    nw326_[int(2)] = len(p0)
                    nw326_[int(3)] = (0) - ((0) - ((self).f11))
                    nw326_[int(4)] = (self).f12
                    nw326_[int(5)] = len(p0)
                    d_2140_v83_ = nw326_
                    d_2141_v84_: _dafny.Map
                    d_2141_v84_ = _dafny.Map({d_2140_v83_: (self).f15})
                    d_2141_v84_ = (d_2141_v84_).set(d_2140_v83_, ((self).f12) != (p3))
                    (globalState).f2 = (self).f15
                    pass
            pass
        d_2142_v85_: D0
        d_2142_v85_ = D0_DC0(_dafny.MultiSet([d_2082_v52_, d_2082_v52_]))
        pat_let_tv38_ = p2
        def lambda80_(source32_):
            if source32_.is_DC1:
                d_2143___mcc_h10_ = source32_.cf1
                d_2144___mcc_h11_ = source32_.cf2
                d_2145_cf2_ = d_2144___mcc_h11_
                d_2146_cf1_ = d_2143___mcc_h10_
                return (self).f15
            elif True:
                d_2147___mcc_h12_ = source32_.cf0
                d_2148_cf0_ = d_2147___mcc_h12_
                return pat_let_tv38_

        r0 = lambda80_(d_2142_v85_)
        d_2149_v86_: _dafny.Map
        d_2149_v86_ = _dafny.Map({(self).f15: (self).f14})
        r1 = default__.safeModuloInt(len((d_2149_v86_) | (d_2149_v86_)), (self).f11)
        d_2150_v87_: D6
        d_2150_v87_ = D6_DC20(p1, (self).f11)
        d_2151_v88_: _dafny.Set
        d_2151_v88_ = _dafny.Set({d_2150_v87_, D6_DC20((0) - (p3), p1)})
        d_2152_v89_: _dafny.MultiSet
        d_2152_v89_ = _dafny.MultiSet([d_2082_v52_])
        r2 = default__.fm0((d_2152_v89_).ispropersubset((_dafny.MultiSet([d_2082_v52_])).set(d_2082_v52_, default__.abs(len(d_2151_v88_)))), globalState)
        return r0, r1, r2

    def m15(self, p0, p1, p2, p3, globalState):
        (globalState).f6 = (self).f12
        d_2153_v0_: _dafny.Array
        nw327_ = _dafny.Array(int(0), 16)
        d_2153_v0_ = nw327_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_2153_v0_).length(0)):
            d_2154_i0_: int = guard_loop_6_
            if (True) and (((0) <= (d_2154_i0_)) and ((d_2154_i0_) < ((d_2153_v0_).length(0)))):
                (d_2153_v0_)[(d_2154_i0_)] = (d_2154_i0_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2155_i1_ in range(default__.abs(445))])))
        if (self).f15:
            (globalState).f2 = (self).f15
            d_2156_v1_: _dafny.Seq
            d_2156_v1_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_2157_v2_: _dafny.Seq
            d_2157_v2_ = _dafny.SeqWithoutIsStrInference([(self).f15, p3, p3])
            d_2158_v3_: _dafny.Seq
            d_2158_v3_ = _dafny.SeqWithoutIsStrInference([d_2157_v2_])
            d_2159_v4_: _dafny.Map
            d_2159_v4_ = _dafny.Map({p3: (d_2158_v3_)[default__.safeIndex((self).f12, len(d_2158_v3_))]})
            (self).f16 = (d_2156_v1_)[default__.safeIndex((0) - ((len((d_2159_v4_).set(p3, _dafny.SeqWithoutIsStrInference([(self).f15]))) if p3 else (self).f12)), len(d_2156_v1_))]
            d_2160_v5_: _dafny.Array
            def lambda81_(d_2161_i2_):
                return not(False)

            init47_ = lambda81_
            nw328_ = _dafny.Array(None, 26)
            for i0_47_ in range(nw328_.length(0)):
                nw328_[i0_47_] = init47_(i0_47_)
            d_2160_v5_ = nw328_
            d_2160_v5_ = d_2160_v5_
            d_2162_v6_: _dafny.Seq
            d_2162_v6_ = _dafny.SeqWithoutIsStrInference([(self).f11])
            d_2163_v7_: _dafny.Seq
            d_2163_v7_ = _dafny.SeqWithoutIsStrInference([(self).f12, len(d_2162_v6_), len(_dafny.Map({(self).f11: (0) - ((self).f11)}))])
            d_2164_v8_: _dafny.Map
            d_2164_v8_ = _dafny.Map({(self).f11: d_2162_v6_})
            d_2165_v9_: _dafny.Map
            d_2165_v9_ = _dafny.Map({(self).f15: (self).f15})
            d_2166_v10_: _dafny.Map
            d_2166_v10_ = _dafny.Map({156: (D19_DC56(d_2153_v0_, (self).f15, (self).f12, _dafny.SeqWithoutIsStrInference([d_2165_v9_, d_2165_v9_]))).cf107})
            d_2167_v11_: T0
            nw329_ = C1()
            nw329_.ctor__(d_2160_v5_, (self).f12, (0) - ((self).f12), len((d_2166_v10_).set(626, 55)), (self).f15, (self).f13)
            d_2167_v11_ = nw329_
            d_2168_v12_: D16
            d_2168_v12_ = D16_DC44(d_2167_v11_)
            d_2169_v13_: _dafny.Map
            d_2169_v13_ = _dafny.Map({d_2168_v12_: _dafny.SeqWithoutIsStrInference([(self).f12 for d_2170_i3_ in range(default__.abs(308))])})
            d_2171_v14_: D19
            d_2171_v14_ = D19_DC56(d_2153_v0_, (d_2157_v2_)[default__.safeIndex((self).f14, len(d_2157_v2_))], len(((d_2164_v8_)[-28] if (-28) in (d_2164_v8_) else ((d_2169_v13_)[d_2168_v12_] if (d_2168_v12_) in (d_2169_v13_) else d_2163_v7_))), _dafny.SeqWithoutIsStrInference([(d_2165_v9_).set((self).f15, p3)]))
            source33_ = d_2171_v14_
            if source33_.is_DC56:
                d_2172___mcc_h0_ = source33_.cf105
                d_2173___mcc_h1_ = source33_.cf106
                d_2174___mcc_h2_ = source33_.cf107
                d_2175___mcc_h3_ = source33_.cf108
                d_2176_cf108_ = d_2175___mcc_h3_
                d_2177_cf107_ = d_2174___mcc_h2_
                d_2178_cf106_ = d_2173___mcc_h1_
                d_2179_cf105_ = d_2172___mcc_h0_
                d_2180_v15_: _dafny.MultiSet
                d_2180_v15_ = _dafny.MultiSet([len(d_2157_v2_), len(d_2163_v7_)])
                d_2181_v16_: C3
                nw330_ = C3()
                nw330_.ctor__(True, self.f16, (self.f16 if (self).f15 else (self).f13), ((d_2180_v15_)[(self).f14] if ((self).f14) in (d_2180_v15_) else p1), p3, (self).f13, (d_2167_v11_).f12, (0) - (d_2177_cf107_))
                d_2181_v16_ = nw330_
                d_2182_v17_: str
                d_2182_v17_ = _dafny.CodePoint('p')
                d_2183_v18_: D16
                d_2183_v18_ = D16_DC45(d_2160_v5_, d_2182_v17_, (d_2167_v11_).f12, (d_2181_v16_).fm2((d_2167_v11_).f11, ((d_2165_v9_)[p3] if (p3) in (d_2165_v9_) else True), p1, d_2178_cf106_, globalState), p3)
                d_2183_v18_ = d_2183_v18_
                d_2184_v19_: _dafny.Map
                d_2184_v19_ = _dafny.Map({d_2177_cf107_: (d_2181_v16_).f25})
                index372_ = default__.safeIndex(855, (d_2179_cf105_).length(0))
                (d_2179_cf105_)[index372_] = default__.safeDivisionInt(len(p0), len(d_2184_v19_))
                index373_ = default__.safeIndex(855, (d_2179_cf105_).length(0))
                (d_2179_cf105_)[index373_] = ((d_2167_v11_).f11) + ((self).f11)
                d_2185_v20_: _dafny.Map
                d_2185_v20_ = _dafny.Map({((d_2184_v19_)[(d_2167_v11_).f12] if ((d_2167_v11_).f12) in (d_2184_v19_) else d_2178_cf106_): (self).f11})
                d_2186_v21_: _dafny.Map
                d_2186_v21_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_2167_v11_).f11 for d_2187_i4_ in range(default__.abs(258))]): (d_2185_v20_ if p3 else (d_2185_v20_).set(not((d_2181_v16_).f25), p1))})
                d_2186_v21_ = d_2186_v21_
            elif True:
                d_2188___mcc_h4_ = source33_.cf104
                d_2189_cf104_ = d_2188___mcc_h4_
                d_2190_v22_: _dafny.Map
                d_2190_v22_ = _dafny.Map({(d_2158_v3_)[default__.safeIndex((self).f11, len(d_2158_v3_))]: False})
                d_2190_v22_ = (d_2190_v22_).set((d_2157_v2_) + (_dafny.SeqWithoutIsStrInference([(self).fm24((d_2167_v11_).f12, not((self).f15), globalState)])), (self).f15)
                d_2191_v23_: _dafny.MultiSet
                d_2191_v23_ = _dafny.MultiSet([(self).f15])
                d_2192_v24_: _dafny.Map
                d_2192_v24_ = _dafny.Map({p3: (622) + ((d_2191_v23_).cardinality)})
                d_2192_v24_ = (d_2192_v24_).set(((d_2167_v11_).f12) != (543), (0) - (default__.safeModuloInt(len(d_2162_v6_), (d_2167_v11_).f11)))
                d_2192_v24_ = (d_2192_v24_ if (self).f15 else (d_2192_v24_) | (d_2192_v24_))
                (globalState).f6 = (d_2167_v11_).f12
            d_2193_v25_: str
            d_2193_v25_ = _dafny.CodePoint('n')
            d_2193_v25_ = (_dafny.CodePoint('t') if (self).f15 else _dafny.CodePoint('x'))
        elif True:
            d_2194_v26_: _dafny.Set
            d_2194_v26_ = _dafny.Set({(self).f14, -164, (self).f12})
            d_2195_v27_: _dafny.Seq
            d_2195_v27_ = _dafny.SeqWithoutIsStrInference([d_2194_v26_])
            d_2196_v28_: _dafny.Seq
            d_2196_v28_ = _dafny.SeqWithoutIsStrInference([p3])
            d_2197_v29_: _dafny.MultiSet
            d_2197_v29_ = _dafny.MultiSet([(d_2194_v26_).intersection(d_2194_v26_), (d_2195_v27_)[default__.safeIndex(len(d_2196_v28_), len(d_2195_v27_))], default__.fm29(globalState)])
            d_2198_v30_: _dafny.Seq
            d_2198_v30_ = _dafny.SeqWithoutIsStrInference([d_2197_v29_])
            d_2199_v31_: _dafny.Seq
            d_2199_v31_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2200_v32_: C2
            nw331_ = C2()
            nw331_.ctor__(167, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehkwmbwj")), len(d_2199_v31_), -273)
            d_2200_v32_ = nw331_
            d_2201_v33_: D18
            d_2201_v33_ = D18_DC53((self).f14, d_2194_v26_, (self).f11, d_2200_v32_)
            d_2197_v29_ = (d_2197_v29_).intersection(((d_2198_v30_)[default__.safeIndex(len((d_2201_v33_).cf100), len(d_2198_v30_))]).intersection(d_2197_v29_))
            d_2202_v34_: _dafny.Set
            d_2202_v34_ = _dafny.Set({p3})
            d_2203_v35_: _dafny.Array
            def lambda82_(d_2204_i5_):
                return _dafny.SeqWithoutIsStrInference([(self).f11 for d_2205_i6_ in range(default__.abs(134))])

            init48_ = lambda82_
            nw332_ = _dafny.Array(None, 14)
            for i0_48_ in range(nw332_.length(0)):
                nw332_[i0_48_] = init48_(i0_48_)
            d_2203_v35_ = nw332_
            d_2206_v36_: int
            out64_: int
            out64_ = (d_2200_v32_).m20(p3, d_2202_v34_, d_2203_v35_, not(p3), globalState)
            d_2206_v36_ = out64_
            (d_2200_v32_).f27 = default__.safeModuloInt(((self).f12) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdjyeq")))), d_2206_v36_)
            d_2207_v37_: _dafny.Array
            nw333_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
            d_2207_v37_ = nw333_
            index374_ = default__.safeIndex(36, (d_2207_v37_).length(0))
            (d_2207_v37_)[index374_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmgyjqauo"))
            d_2208_v38_: _dafny.Map
            d_2208_v38_ = _dafny.Map({(self).f15: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwhmi"))) + ((self).f13)})
            d_2209_v39_: str
            d_2209_v39_ = _dafny.CodePoint('o')
            index375_ = default__.safeIndex(36, (d_2207_v37_).length(0))
            (d_2207_v37_)[index375_] = ((d_2208_v38_)[(self).f15] if ((self).f15) in (d_2208_v38_) else (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2210_i7_ in range(default__.abs(936))])).set(default__.safeIndex((self).f12, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2211_i7_ in range(default__.abs(936))]))), d_2209_v39_))
            d_2199_v31_ = (d_2199_v31_) + ((d_2199_v31_) + (d_2199_v31_))
        d_2212_v40_: _dafny.Seq
        d_2212_v40_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2213_v41_: D1
        d_2213_v41_ = D1_DC3(not(not(((self).f11) not in (d_2212_v40_))))
        d_2213_v41_ = d_2213_v41_
        if p3:
            d_2214_v42_: str
            d_2214_v42_ = _dafny.CodePoint('u')
            d_2215_v43_: _dafny.Map
            d_2215_v43_ = _dafny.Map({(self).f15: len(_dafny.Set({p0, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcvpfgnfj"))).set(default__.safeIndex((self).f14, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcvpfgnfj")))), d_2214_v42_), self.f16, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oewnblx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yitlhryc"))}))})
            d_2216_v44_: D6
            d_2216_v44_ = D6_DC19(d_2215_v43_)
            source34_ = d_2216_v44_
            if source34_.is_DC20:
                d_2217___mcc_h5_ = source34_.cf38
                d_2218___mcc_h6_ = source34_.cf39
                d_2219_cf39_ = d_2218___mcc_h6_
                d_2220_cf38_ = d_2217___mcc_h5_
                d_2221_v45_: _dafny.Set
                d_2221_v45_ = _dafny.Set({p3, (self).f15})
                d_2222_v46_: _dafny.MultiSet
                d_2222_v46_ = _dafny.MultiSet([d_2214_v42_])
                d_2223_v47_: _dafny.MultiSet
                d_2223_v47_ = _dafny.MultiSet([d_2219_cf39_])
                d_2224_v48_: _dafny.Array
                nw334_ = _dafny.Array(None, 24)
                nw334_[int(0)] = not (p3) or (p3)
                nw334_[int(1)] = p3
                nw334_[int(2)] = not ((self).f15) or ((self).f15)
                nw334_[int(3)] = True
                nw334_[int(4)] = not (p3) or (p3)
                nw334_[int(5)] = True
                nw334_[int(6)] = not(not((True if p3 else (self).f15)))
                nw334_[int(7)] = ((self).f15 if p3 else (self).f15)
                nw334_[int(8)] = (self).f15
                nw334_[int(9)] = True
                nw334_[int(10)] = p3
                nw334_[int(11)] = (self).f15
                nw334_[int(12)] = (d_2221_v45_).ispropersubset(_dafny.Set({p3}))
                nw334_[int(13)] = (len(_dafny.SeqWithoutIsStrInference([(self).f14 for d_2225_i8_ in range(default__.abs(403))]))) > (((d_2222_v46_)[d_2214_v42_] if (d_2214_v42_) in (d_2222_v46_) else (self).f12))
                nw334_[int(14)] = ((self).f15 if (self).f15 else p3)
                nw334_[int(15)] = ((self).f15 if not(p3) else (self).f15)
                nw334_[int(16)] = ((self).f12) >= (d_2219_cf39_)
                nw334_[int(17)] = (self).f15
                nw334_[int(18)] = (self).f15
                nw334_[int(19)] = p3
                nw334_[int(20)] = p3
                nw334_[int(21)] = (self).f15
                nw334_[int(22)] = (self).f15
                nw334_[int(23)] = (_dafny.MultiSet([d_2219_cf39_])) != (d_2223_v47_)
                d_2224_v48_ = nw334_
                index376_ = default__.safeIndex(196, (d_2224_v48_).length(0))
                (d_2224_v48_)[index376_] = (p3 if (self).f15 else not(True))
                index377_ = default__.safeIndex(196, (d_2224_v48_).length(0))
                (d_2224_v48_)[index377_] = (self).f15
                (globalState).f6 = (self).f12
                d_2226_v49_: _dafny.Set
                d_2226_v49_ = _dafny.Set({d_2220_cf38_})
                d_2227_v50_: _dafny.Map
                d_2227_v50_ = _dafny.Map({(d_2226_v49_).isdisjoint(_dafny.Set({d_2219_cf39_})): d_2224_v48_})
                d_2227_v50_ = (d_2227_v50_).set((p3) or (p3), d_2224_v48_)
                (globalState).f6 = p1
            elif source34_.is_DC21:
                d_2228___mcc_h7_ = source34_.cf40
                d_2229_cf40_ = d_2228___mcc_h7_
                d_2230_v51_: _dafny.Seq
                d_2230_v51_ = _dafny.SeqWithoutIsStrInference([d_2153_v0_, d_2153_v0_, d_2153_v0_, d_2153_v0_, d_2153_v0_])
                d_2231_v52_: D3
                d_2231_v52_ = D3_DC9(d_2230_v51_)
                d_2232_v53_: _dafny.Array
                nw335_ = _dafny.Array(None, 11)
                nw335_[int(0)] = d_2231_v52_
                nw335_[int(1)] = d_2231_v52_
                nw335_[int(2)] = D3_DC9(d_2230_v51_)
                nw335_[int(3)] = d_2231_v52_
                nw335_[int(4)] = d_2231_v52_
                nw335_[int(5)] = d_2231_v52_
                nw335_[int(6)] = d_2231_v52_
                nw335_[int(7)] = d_2231_v52_
                nw335_[int(8)] = d_2231_v52_
                nw335_[int(9)] = d_2231_v52_
                nw335_[int(10)] = d_2231_v52_
                d_2232_v53_ = nw335_
                index378_ = default__.safeIndex(680, (d_2232_v53_).length(0))
                (d_2232_v53_)[index378_] = d_2231_v52_
                index379_ = default__.safeIndex(680, (d_2232_v53_).length(0))
                (d_2232_v53_)[index379_] = d_2231_v52_
                d_2233_v54_: _dafny.Set
                d_2233_v54_ = _dafny.Set({(self).f11, (self).f12})
                d_2234_v55_: _dafny.Set
                d_2234_v55_ = _dafny.Set({d_2233_v54_})
                d_2235_v56_: _dafny.Seq
                d_2235_v56_ = _dafny.SeqWithoutIsStrInference([(self).f15, d_2229_cf40_, p3, (self).f15, (self).f15])
                d_2236_v57_: _dafny.MultiSet
                d_2236_v57_ = _dafny.MultiSet([(self).fm24(208, d_2229_cf40_, globalState), (self).f15])
                d_2234_v55_ = default__.fm61((p3) and (not(p3)), (d_2236_v57_).issubset(_dafny.MultiSet(d_2235_v56_)), d_2229_cf40_, (self).f11, globalState)
                (globalState).f6 = (0) - ((self).f11)
                d_2237_v58_: _dafny.Array
                nw336_ = _dafny.Array(False, 14)
                d_2237_v58_ = nw336_
                d_2238_v59_: C1
                nw337_ = C1()
                nw337_.ctor__(d_2237_v58_, (self).f11, p1, (self).f14, p3, p0)
                d_2238_v59_ = nw337_
            elif source34_.is_DC19:
                d_2239___mcc_h8_ = source34_.cf37
                d_2240_cf37_ = d_2239___mcc_h8_
                d_2241_v60_: _dafny.Map
                d_2241_v60_ = _dafny.Map({(self).f15: p3})
                d_2242_v61_: _dafny.Map
                d_2242_v61_ = _dafny.Map({(self).f15: d_2241_v60_})
                d_2243_v62_: D15
                d_2243_v62_ = D15_DC42((d_2242_v61_) | (d_2242_v61_))
                d_2244_v63_: _dafny.Set
                d_2244_v63_ = _dafny.Set({not(p3)})
                rhs359_ = d_2243_v62_
                rhs360_ = default__.safeDivisionInt((0) - (len(d_2244_v63_)), (self).f12)
                lhs293_ = globalState
                d_2243_v62_ = rhs359_
                lhs293_.f6 = rhs360_
                d_2245_v64_: _dafny.Seq
                d_2245_v64_ = _dafny.SeqWithoutIsStrInference([p3])
                d_2246_v65_: _dafny.MultiSet
                d_2246_v65_ = _dafny.MultiSet([(d_2245_v64_)[default__.safeIndex((self).f12, len(d_2245_v64_))], not((self).f15), p3, p3])
                d_2247_v66_: _dafny.Array
                nw338_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
                d_2247_v66_ = nw338_
                index380_ = default__.safeIndex(807, (d_2247_v66_).length(0))
                (d_2247_v66_)[index380_] = (self).f13
                d_2248_v67_: _dafny.Seq
                d_2248_v67_ = _dafny.SeqWithoutIsStrInference([(self).f13])
                index381_ = default__.safeIndex(807, (d_2247_v66_).length(0))
                rhs361_ = p3
                rhs362_ = (self).fm2(default__.safeDivisionInt(p1, (0) - (p1)), (self).fm1((0) - (len(d_2245_v64_)), p3, p3, globalState), (self).f11, (self).f15, globalState)
                rhs363_ = d_2246_v65_
                rhs364_ = ((d_2248_v67_)[default__.safeIndex((self).f14, len(d_2248_v67_))] if ((self).f15 if (self).f15 else p3) else self.f16)
                rhs365_ = p1
                lhs294_ = globalState
                lhs295_ = globalState
                lhs296_ = d_2247_v66_
                lhs297_ = default__.safeIndex(807, (d_2247_v66_).length(0))
                lhs298_ = globalState
                lhs294_.f7 = rhs361_
                lhs295_.f6 = rhs362_
                d_2246_v65_ = rhs363_
                lhs296_[lhs297_] = rhs364_
                lhs298_.f6 = rhs365_
                d_2212_v40_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f11, default__.safeDivisionInt((self).f11, (self).f12)])
                (globalState).f2 = p3
            elif True:
                d_2249___mcc_h9_ = source34_.cf41
                d_2250_cf41_ = d_2249___mcc_h9_
                d_2153_v0_ = d_2153_v0_
                d_2251_v68_: _dafny.Seq
                d_2251_v68_ = _dafny.SeqWithoutIsStrInference([self.f16, (p0) + ((self).f13), p0])
                d_2251_v68_ = d_2251_v68_
                d_2252_v69_: _dafny.Array
                def lambda83_(d_2253_i9_):
                    return _dafny.Map({(self).f12: (self).f13})

                init49_ = lambda83_
                nw339_ = _dafny.Array(None, 28)
                for i0_49_ in range(nw339_.length(0)):
                    nw339_[i0_49_] = init49_(i0_49_)
                d_2252_v69_ = nw339_
                d_2254_v71_: _dafny.Set
                d_2254_v71_ = _dafny.Set({503})
                index382_ = default__.safeIndex(263, (d_2252_v69_).length(0))
                def iife133_():
                    coll55_ = _dafny.Map()
                    compr_55_: int
                    for compr_55_ in (d_2254_v71_).Elements:
                        d_2255_v70_: int = compr_55_
                        if (d_2255_v70_) in (d_2254_v71_):
                            coll55_[(d_2255_v70_) - (len(_dafny.Map({(self).f15: d_2212_v40_})))] = _dafny.SeqWithoutIsStrInference([d_2214_v42_ for d_2256_i10_ in range(default__.abs(-19))])
                    return _dafny.Map(coll55_)
                (d_2252_v69_)[index382_] = iife133_()
                
                d_2257_v72_: D4
                d_2257_v72_ = D4_DC15((self).f14)
                d_2258_v73_: _dafny.Map
                d_2258_v73_ = _dafny.Map({D4_DC16(D4_DC16(d_2257_v72_)): p3})
                d_2259_v74_: _dafny.Map
                d_2259_v74_ = _dafny.Map({-772: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oeptsqduf"))})
                d_2260_v75_: _dafny.Map
                d_2260_v75_ = _dafny.Map({(0) - ((0) - (len(d_2258_v73_))): (d_2259_v74_) | (d_2259_v74_)})
                d_2261_v76_: _dafny.Array
                def lambda84_(d_2262_i11_):
                    return (self).f15

                init50_ = lambda84_
                nw340_ = _dafny.Array(None, 27)
                for i0_50_ in range(nw340_.length(0)):
                    nw340_[i0_50_] = init50_(i0_50_)
                d_2261_v76_ = nw340_
                d_2263_v77_: _dafny.Array
                nw341_ = _dafny.Array(None, 5)
                d_2263_v77_ = nw341_
                d_2264_v78_: _dafny.Set
                d_2264_v78_ = _dafny.Set({d_2263_v77_, d_2263_v77_})
                d_2265_v79_: D1
                d_2265_v79_ = D1_DC2(d_2214_v42_)
                index383_ = default__.safeIndex(263, (d_2252_v69_).length(0))
                (d_2252_v69_)[index383_] = ((d_2260_v75_)[default__.safeModuloInt((D2_DC7(d_2261_v76_, d_2264_v78_, d_2265_v79_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "os"))), p1)).cf13, p1)] if (default__.safeModuloInt((D2_DC7(d_2261_v76_, d_2264_v78_, d_2265_v79_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "os"))), p1)).cf13, p1)) in (d_2260_v75_) else d_2259_v74_)
                d_2266_v80_: _dafny.MultiSet
                d_2266_v80_ = _dafny.MultiSet([(self).f15])
                d_2267_v81_: _dafny.Seq
                d_2267_v81_ = _dafny.SeqWithoutIsStrInference([p3])
                d_2268_v82_: _dafny.Map
                d_2268_v82_ = _dafny.Map({(d_2266_v80_).issubset(_dafny.MultiSet(d_2267_v81_)): d_2261_v76_})
                d_2268_v82_ = (d_2268_v82_).set((self).fm1((0) - ((self).f11), (self).f15, (self).f15, globalState), d_2261_v76_)
            d_2269_v83_: _dafny.Seq
            d_2269_v83_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            (globalState).f2 = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, (self).f15]))).issubset(default__.fm54(len(d_2269_v83_), globalState))
            source35_ = d_2216_v44_
            if source35_.is_DC20:
                d_2270___mcc_h10_ = source35_.cf38
                d_2271___mcc_h11_ = source35_.cf39
                d_2272_cf39_ = d_2271___mcc_h11_
                d_2273_cf38_ = d_2270___mcc_h10_
                d_2274_v84_: _dafny.Map
                d_2274_v84_ = _dafny.Map({(self).f11: (self).f11})
                d_2275_v85_: D17
                d_2275_v85_ = D17_DC47(d_2274_v84_)
                d_2276_v86_: _dafny.Seq
                d_2276_v86_ = _dafny.SeqWithoutIsStrInference([(d_2275_v85_).cf89, d_2274_v84_])
                d_2277_v87_: _dafny.MultiSet
                d_2277_v87_ = _dafny.MultiSet([p3])
                (globalState).f6 = default__.safeModuloInt(len(d_2276_v86_), ((d_2277_v87_)[p3] if (p3) in (d_2277_v87_) else d_2272_cf39_))
                d_2273_cf38_ = default__.fm0((self).f15, globalState)
                d_2278_v88_: _dafny.Array
                nw342_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                d_2278_v88_ = nw342_
                index384_ = default__.safeIndex(330, (d_2278_v88_).length(0))
                (d_2278_v88_)[index384_] = _dafny.CodePoint('h')
                index385_ = default__.safeIndex(330, (d_2278_v88_).length(0))
                (d_2278_v88_)[index385_] = d_2214_v42_
                d_2279_v89_: _dafny.Array
                def lambda85_(d_2280_p3_):
                    def lambda86_(d_2281_i12_):
                        return d_2280_p3_

                    return lambda86_

                init51_ = lambda85_(p3)
                nw343_ = _dafny.Array(None, 24)
                for i0_51_ in range(nw343_.length(0)):
                    nw343_[i0_51_] = init51_(i0_51_)
                d_2279_v89_ = nw343_
                d_2282_v90_: C1
                nw344_ = C1()
                nw344_.ctor__(d_2279_v89_, (self).f14, ((0) - ((0) - ((self).f14))) + (d_2273_cf38_), (0) - (d_2273_cf38_), (p1) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({671})) for d_2283_i13_ in range(default__.abs(650))])), ((self).f13 if p3 else (self).f13))
                d_2282_v90_ = nw344_
            elif source35_.is_DC21:
                d_2284___mcc_h12_ = source35_.cf40
                d_2285_cf40_ = d_2284___mcc_h12_
                d_2269_v83_ = d_2269_v83_
                rhs366_ = ((p0) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymcpfs")))) == (p3)
                rhs367_ = (self).f14
                lhs299_ = globalState
                lhs300_ = globalState
                lhs299_.f7 = rhs366_
                lhs300_.f6 = rhs367_
                d_2212_v40_ = d_2212_v40_
                d_2286_v91_: _dafny.Array
                nw345_ = _dafny.Array(D0.default()(), 11)
                d_2286_v91_ = nw345_
                d_2287_v92_: _dafny.Seq
                d_2287_v92_ = _dafny.SeqWithoutIsStrInference([d_2286_v91_])
                d_2288_v93_: _dafny.Array
                nw346_ = _dafny.Array(None, 9)
                nw346_[int(0)] = d_2286_v91_
                nw346_[int(1)] = (d_2287_v92_)[default__.safeIndex((self).f11, len(d_2287_v92_))]
                nw346_[int(2)] = d_2286_v91_
                nw346_[int(3)] = d_2286_v91_
                nw346_[int(4)] = d_2286_v91_
                nw346_[int(5)] = d_2286_v91_
                nw346_[int(6)] = d_2286_v91_
                nw346_[int(7)] = (d_2286_v91_ if (self).f15 else d_2286_v91_)
                nw346_[int(8)] = d_2286_v91_
                d_2288_v93_ = nw346_
                d_2288_v93_ = d_2288_v93_
            elif source35_.is_DC19:
                d_2289___mcc_h13_ = source35_.cf37
                d_2290_cf37_ = d_2289___mcc_h13_
                d_2214_v42_ = d_2214_v42_
                (globalState).f7 = p3
                (globalState).f2 = (self).f15
                d_2291_v94_: _dafny.Array
                def lambda87_(d_2292_i14_):
                    return (self).f15

                init52_ = lambda87_
                nw347_ = _dafny.Array(None, 11)
                for i0_52_ in range(nw347_.length(0)):
                    nw347_[i0_52_] = init52_(i0_52_)
                d_2291_v94_ = nw347_
                d_2293_v95_: _dafny.Set
                d_2293_v95_ = _dafny.Set({d_2291_v94_})
                d_2294_v96_: D19
                d_2294_v96_ = D19_DC55(d_2293_v95_)
                d_2295_v97_: _dafny.Seq
                d_2295_v97_ = _dafny.SeqWithoutIsStrInference([d_2294_v96_])
                rhs368_ = ((self).f14) - ((self).f12)
                rhs369_ = (d_2295_v97_) + (d_2295_v97_)
                lhs301_ = globalState
                lhs301_.f6 = rhs368_
                d_2295_v97_ = rhs369_
            elif True:
                d_2296___mcc_h14_ = source35_.cf41
                d_2297_cf41_ = d_2296___mcc_h14_
                (globalState).f7 = ((_dafny.SeqWithoutIsStrInference([d_2214_v42_ for d_2298_i15_ in range(default__.abs(803))])) + (p0)) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tiwiwo")) if p3 else (self).f13))
                (globalState).f6 = (self).f12
                d_2299_v98_: _dafny.Map
                d_2299_v98_ = _dafny.Map({default__.fm0((self).f15, globalState): _dafny.CodePoint('n')})
                d_2300_v99_: D12
                d_2300_v99_ = D12_DC36(p1, p1, (0) - ((self).f14))
                (globalState).f6 = default__.safeDivisionInt((len((d_2299_v98_).set(917, d_2214_v42_))) + (788), (d_2300_v99_).cf64)
                index386_ = default__.safeIndex(511, (d_2153_v0_).length(0))
                (d_2153_v0_)[index386_] = default__.fm0(p3, globalState)
                index387_ = default__.safeIndex(511, (d_2153_v0_).length(0))
                (d_2153_v0_)[index387_] = (self).f12
            d_2301_v100_: _dafny.Array
            def lambda88_(d_2302_i16_):
                return (self).f15

            init53_ = lambda88_
            nw348_ = _dafny.Array(None, 26)
            for i0_53_ in range(nw348_.length(0)):
                nw348_[i0_53_] = init53_(i0_53_)
            d_2301_v100_ = nw348_
            d_2303_v101_: _dafny.Map
            d_2303_v101_ = _dafny.Map({p3: p3})
            d_2304_v102_: _dafny.Map
            d_2304_v102_ = _dafny.Map({p3: d_2213_v41_})
            d_2305_v103_: C1
            nw349_ = C1()
            nw349_.ctor__(d_2301_v100_, len(d_2303_v101_), len(d_2304_v102_), p1, p3, p0)
            d_2305_v103_ = nw349_
            d_2306_v104_: _dafny.Seq
            d_2306_v104_ = _dafny.SeqWithoutIsStrInference([d_2305_v103_, d_2305_v103_, d_2305_v103_, d_2305_v103_])
            d_2307_v105_: _dafny.MultiSet
            d_2307_v105_ = _dafny.MultiSet([d_2306_v104_, d_2306_v104_])
            rhs370_ = (d_2307_v105_) | (_dafny.MultiSet([d_2306_v104_]))
            rhs371_ = p1
            rhs372_ = d_2305_v103_.f24
            lhs302_ = globalState
            lhs303_ = d_2305_v103_
            d_2307_v105_ = rhs370_
            lhs302_.f6 = rhs371_
            lhs303_.f24 = rhs372_
            d_2214_v42_ = d_2214_v42_
        elif True:
            if not ((self).f15) or ((p1) > (p1)):
                d_2308_v106_: bool
                d_2309_v107_: _dafny.Seq
                out65_: bool
                out66_: _dafny.Seq
                out65_, out66_ = default__.m0((self).f15, not((self).f15), p1, (self).f11, globalState)
                d_2308_v106_ = out65_
                d_2309_v107_ = out66_
                d_2310_v108_: _dafny.Set
                d_2311_v109_: int
                d_2312_v110_: bool
                d_2313_v111_: int
                out67_: _dafny.Set
                out68_: int
                out69_: bool
                out70_: int
                out67_, out68_, out69_, out70_ = (self).m16((d_2153_v0_ if p3 else d_2153_v0_), not((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_2314_i17_ in range(default__.abs(743))])) < ((self).f13)), p0, _dafny.MultiSet([p3, (self).fm1((self).f11, (self).f15, not(not((self).f15)), globalState), d_2308_v106_]), globalState)
                d_2310_v108_ = out67_
                d_2311_v109_ = out68_
                d_2312_v110_ = out69_
                d_2313_v111_ = out70_
                d_2315_v112_: str
                d_2315_v112_ = _dafny.CodePoint('g')
                d_2316_v113_: _dafny.Map
                d_2316_v113_ = _dafny.Map({((self).f12) == ((self).f11): d_2315_v112_})
                d_2316_v113_ = (d_2316_v113_).set(d_2312_v110_, d_2315_v112_)
                (globalState).f7 = d_2308_v106_
                d_2317_v114_: _dafny.MultiSet
                d_2317_v114_ = _dafny.MultiSet([d_2312_v110_])
                d_2318_v115_: D17
                d_2318_v115_ = D17_DC50((d_2317_v114_).set(d_2312_v110_, default__.abs(102)))
                d_2318_v115_ = d_2318_v115_
            elif True:
                d_2319_v116_: _dafny.Seq
                d_2319_v116_ = _dafny.SeqWithoutIsStrInference([d_2153_v0_])
                d_2320_v117_: D3
                d_2320_v117_ = D3_DC9(d_2319_v116_)
                d_2321_v118_: D13
                d_2321_v118_ = D13_DC38(d_2320_v117_, (self).f15, p3)
                d_2322_v119_: D13
                d_2322_v119_ = D13_DC39(d_2321_v118_)
                d_2322_v119_ = d_2322_v119_
                (globalState).f6 = default__.safeDivisionInt((self).f14, default__.fm0(p3, globalState))
                d_2323_v120_: str
                d_2323_v120_ = _dafny.CodePoint('u')
                d_2324_v121_: _dafny.Map
                d_2324_v121_ = _dafny.Map({(self).f14: d_2323_v120_})
                d_2325_v122_: _dafny.Set
                d_2325_v122_ = _dafny.Set({len(p0)})
                d_2324_v121_ = (d_2324_v121_).set(len(((_dafny.SeqWithoutIsStrInference([(self).f12 for d_2326_i18_ in range(default__.abs(578))])).set(default__.safeIndex(-247, len(_dafny.SeqWithoutIsStrInference([(self).f12 for d_2327_i18_ in range(default__.abs(578))]))), len(d_2325_v122_))) + (_dafny.SeqWithoutIsStrInference([(self).f12 for d_2328_i19_ in range(default__.abs(337))]))), _dafny.CodePoint('u'))
                (globalState).f6 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbpofg")))
                d_2329_v123_: _dafny.Set
                d_2329_v123_ = _dafny.Set({(D6_DC19(_dafny.Map({(self).f15: p1}))).cf37})
                d_2330_v124_: D1
                d_2330_v124_ = D1_DC5((self).f14, p3, (self).f15)
                d_2331_v125_: _dafny.Map
                d_2331_v125_ = _dafny.Map({(self).f15: (self).f14})
                d_2332_v126_: _dafny.Map
                d_2332_v126_ = _dafny.Map({(d_2330_v124_).cf8: ((d_2331_v125_)[(self).f15] if ((self).f15) in (d_2331_v125_) else (self).f11)})
                d_2329_v123_ = _dafny.Set({(d_2332_v126_ if p3 else d_2331_v125_), (d_2331_v125_) | (d_2332_v126_), d_2331_v125_, d_2331_v125_})
            d_2333_v127_: str
            d_2333_v127_ = _dafny.CodePoint('b')
            d_2334_v128_: _dafny.MultiSet
            d_2334_v128_ = _dafny.MultiSet([d_2333_v127_, d_2333_v127_])
            d_2335_v129_: _dafny.Map
            d_2335_v129_ = _dafny.Map({(self).f15: p3})
            d_2336_v130_: _dafny.Array
            nw350_ = _dafny.Array(None, 12)
            nw350_[int(0)] = (self).f15
            nw350_[int(1)] = ((self).f12) >= ((self).f11)
            nw350_[int(2)] = not(p3)
            nw350_[int(3)] = (self).f15
            nw350_[int(4)] = (d_2334_v128_).ispropersubset(d_2334_v128_)
            nw350_[int(5)] = (self).f15
            nw350_[int(6)] = p3
            nw350_[int(7)] = (self).f15
            nw350_[int(8)] = ((d_2335_v129_)[p3] if (p3) in (d_2335_v129_) else (self).f15)
            nw350_[int(9)] = (self).f15
            nw350_[int(10)] = p3
            nw350_[int(11)] = (self).f15
            d_2336_v130_ = nw350_
            index388_ = default__.safeIndex(483, (d_2336_v130_).length(0))
            (d_2336_v130_)[index388_] = p3
            index389_ = default__.safeIndex(156, (d_2153_v0_).length(0))
            (d_2153_v0_)[index389_] = ((self).f12) - (171)
            d_2337_v131_: _dafny.Set
            d_2337_v131_ = _dafny.Set({157})
            d_2338_v132_: _dafny.Seq
            d_2338_v132_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            d_2339_v133_: _dafny.Map
            d_2339_v133_ = _dafny.Map({(self).f14: (self).f13})
            index390_ = default__.safeIndex(483, (d_2336_v130_).length(0))
            index391_ = default__.safeIndex(156, (d_2153_v0_).length(0))
            rhs373_ = (_dafny.Set({p1})).ispropersubset(d_2337_v131_)
            rhs374_ = (0) - (((p1) + (319)) * ((0) - ((self).f14)))
            rhs375_ = ((d_2338_v132_).set(default__.safeIndex((self).f14, len(d_2338_v132_)), (self).f15)) <= (d_2338_v132_)
            rhs376_ = (p0) <= ((((d_2339_v133_)[(self).f12] if ((self).f12) in (d_2339_v133_) else (self).f13)) + (self.f16))
            lhs304_ = d_2336_v130_
            lhs305_ = default__.safeIndex(483, (d_2336_v130_).length(0))
            lhs306_ = d_2153_v0_
            lhs307_ = default__.safeIndex(156, (d_2153_v0_).length(0))
            lhs308_ = globalState
            lhs309_ = globalState
            lhs304_[lhs305_] = rhs373_
            lhs306_[lhs307_] = rhs374_
            lhs308_.f2 = rhs375_
            lhs309_.f7 = rhs376_
            d_2340_v134_: _dafny.Array
            nw351_ = _dafny.Array(_dafny.CodePoint('D'), 13)
            d_2340_v134_ = nw351_
            d_2340_v134_ = (d_2340_v134_ if ((d_2153_v0_)[default__.safeIndex(156, (d_2153_v0_).length(0))]) != ((d_2212_v40_)[default__.safeIndex((self).f12, len(d_2212_v40_))]) else d_2340_v134_)
            d_2341_v135_: _dafny.Map
            d_2341_v135_ = _dafny.Map({default__.fm30(p3, (self).f15, p1, (d_2336_v130_)[default__.safeIndex(483, (d_2336_v130_).length(0))], globalState): (p1) != (default__.fm0(not(p3), globalState))})
            d_2341_v135_ = d_2341_v135_
            d_2342_v136_: _dafny.Array
            nw352_ = _dafny.Array(_dafny.Map({}), 1)
            d_2342_v136_ = nw352_
            d_2343_v137_: _dafny.Map
            d_2343_v137_ = _dafny.Map({not((self).f15): (self).f12})
            index392_ = default__.safeIndex(854, (d_2342_v136_).length(0))
            (d_2342_v136_)[index392_] = (d_2343_v137_) | (d_2343_v137_)
            index393_ = default__.safeIndex(854, (d_2342_v136_).length(0))
            (d_2342_v136_)[index393_] = ((d_2343_v137_).set((self).f15, (self).f11)) | ((d_2343_v137_) | (d_2343_v137_))
        if (self).f15:
            d_2344_v138_: _dafny.Seq
            d_2344_v138_ = _dafny.SeqWithoutIsStrInference([d_2153_v0_, d_2153_v0_])
            d_2345_v139_: D3
            d_2345_v139_ = D3_DC9(d_2344_v138_)
            d_2346_v140_: D13
            d_2346_v140_ = D13_DC38(d_2345_v139_, (self).f15, (self).f15)
            d_2347_v141_: str
            d_2347_v141_ = _dafny.CodePoint('o')
            d_2348_v142_: D2
            d_2348_v142_ = D2_DC8(True, (self).f15, p0)
            d_2349_v143_: _dafny.MultiSet
            d_2349_v143_ = _dafny.MultiSet([d_2347_v141_, d_2347_v141_])
            d_2350_v144_: D0
            d_2350_v144_ = D0_DC0(d_2349_v143_)
            d_2351_v145_: _dafny.Set
            d_2351_v145_ = _dafny.Set({d_2350_v144_})
            d_2352_v146_: D3
            d_2352_v146_ = D3_DC11((self).f11, d_2348_v142_, d_2351_v145_, (self).f15)
            d_2353_v147_: _dafny.Set
            d_2353_v147_ = _dafny.Set({(self).f11})
            d_2354_v148_: _dafny.MultiSet
            d_2354_v148_ = _dafny.MultiSet([(self).f12])
            d_2355_v149_: _dafny.Array
            nw353_ = _dafny.Array(None, 15)
            nw353_[int(0)] = p3
            nw353_[int(1)] = (self).f15
            nw353_[int(2)] = (d_2346_v140_).cf70
            nw353_[int(3)] = (d_2347_v141_) not in (p0)
            nw353_[int(4)] = ((self).f15 if p3 else (d_2352_v146_).cf27)
            nw353_[int(5)] = not(p3)
            nw353_[int(6)] = p3
            nw353_[int(7)] = p3
            nw353_[int(8)] = (d_2353_v147_) != (d_2353_v147_)
            nw353_[int(9)] = (d_2354_v148_).isdisjoint(d_2354_v148_)
            nw353_[int(10)] = ((self).f15 if p3 else p3)
            nw353_[int(11)] = not((self).f15)
            nw353_[int(12)] = (self).f15
            nw353_[int(13)] = False
            nw353_[int(14)] = not(((self).f13) <= (p0))
            d_2355_v149_ = nw353_
            index394_ = default__.safeIndex(34, (d_2355_v149_).length(0))
            (d_2355_v149_)[index394_] = (self).f15
            d_2356_v150_: D2
            d_2356_v150_ = D2_DC6((self).f13)
            d_2357_v151_: _dafny.Seq
            d_2357_v151_ = _dafny.SeqWithoutIsStrInference([d_2355_v149_, d_2355_v149_, d_2355_v149_, d_2355_v149_])
            index395_ = default__.safeIndex(34, (d_2355_v149_).length(0))
            rhs377_ = (len(((self).f13) + ((self).f13))) in (d_2354_v148_)
            rhs378_ = d_2356_v150_
            rhs379_ = _dafny.SeqWithoutIsStrInference([d_2355_v149_, d_2355_v149_, d_2355_v149_, d_2355_v149_])
            lhs310_ = d_2355_v149_
            lhs311_ = default__.safeIndex(34, (d_2355_v149_).length(0))
            lhs310_[lhs311_] = rhs377_
            d_2356_v150_ = rhs378_
            d_2357_v151_ = rhs379_
            (globalState).f6 = (self).f12
            (globalState).f6 = default__.fm0((d_2355_v149_)[default__.safeIndex(34, (d_2355_v149_).length(0))], globalState)
            d_2358_v152_: _dafny.MultiSet
            d_2358_v152_ = _dafny.MultiSet([(self).f15])
            d_2359_v153_: C3
            nw354_ = C3()
            nw354_.ctor__(p3, _dafny.SeqWithoutIsStrInference([d_2347_v141_ for d_2360_i20_ in range(default__.abs(791))]), p0, (self).f11, (d_2355_v149_)[default__.safeIndex(34, (d_2355_v149_).length(0))], p0, (self).f14, (d_2358_v152_).cardinality)
            d_2359_v153_ = nw354_
            d_2361_v154_: _dafny.Map
            d_2361_v154_ = _dafny.Map({d_2153_v0_: d_2213_v41_})
            index396_ = default__.safeIndex(34, (d_2355_v149_).length(0))
            index397_ = default__.safeIndex(34, (d_2355_v149_).length(0))
            rhs380_ = (len(d_2361_v154_)) == ((self).f11)
            rhs381_ = ((self).f14) <= ((111) - ((self).f12))
            rhs382_ = p3
            rhs383_ = d_2359_v153_
            rhs384_ = (d_2355_v149_)[default__.safeIndex(34, (d_2355_v149_).length(0))]
            lhs312_ = d_2355_v149_
            lhs313_ = default__.safeIndex(34, (d_2355_v149_).length(0))
            lhs314_ = globalState
            lhs315_ = globalState
            lhs316_ = d_2355_v149_
            lhs317_ = default__.safeIndex(34, (d_2355_v149_).length(0))
            lhs312_[lhs313_] = rhs380_
            lhs314_.f2 = rhs381_
            lhs315_.f2 = rhs382_
            d_2359_v153_ = rhs383_
            lhs316_[lhs317_] = rhs384_
            index398_ = default__.safeIndex(23, (d_2153_v0_).length(0))
            (d_2153_v0_)[index398_] = (self).f14
            d_2362_v155_: _dafny.Array
            nw355_ = _dafny.Array(_dafny.Array(None, 0), 29)
            d_2362_v155_ = nw355_
            d_2363_v156_: _dafny.Map
            d_2363_v156_ = _dafny.Map({d_2362_v155_: ((self).f12) - (284)})
            d_2364_v157_: _dafny.Map
            d_2364_v157_ = _dafny.Map({self.f16: d_2362_v155_})
            index399_ = default__.safeIndex(23, (d_2153_v0_).length(0))
            (d_2153_v0_)[index399_] = ((d_2363_v156_)[((d_2364_v157_)[(self).f13] if ((self).f13) in (d_2364_v157_) else d_2362_v155_)] if (((d_2364_v157_)[(self).f13] if ((self).f13) in (d_2364_v157_) else d_2362_v155_)) in (d_2363_v156_) else p1)
        elif True:
            d_2365_v158_: _dafny.Map
            d_2365_v158_ = _dafny.Map({(self).f11: p3})
            (globalState).f6 = default__.safeDivisionInt(p1, len(_dafny.Map({902: len(d_2365_v158_)})))
            rhs385_ = not((self).f15)
            rhs386_ = False
            rhs387_ = p3
            rhs388_ = p3
            lhs318_ = globalState
            lhs319_ = globalState
            lhs320_ = globalState
            lhs321_ = globalState
            lhs318_.f2 = rhs385_
            lhs319_.f2 = rhs386_
            lhs320_.f2 = rhs387_
            lhs321_.f7 = rhs388_
            d_2366_v159_: _dafny.Seq
            d_2366_v159_ = _dafny.SeqWithoutIsStrInference([self.f16])
            index400_ = default__.safeIndex(665, (d_2153_v0_).length(0))
            (d_2153_v0_)[index400_] = len(d_2366_v159_)
            d_2367_v160_: str
            d_2367_v160_ = _dafny.CodePoint('g')
            d_2368_v161_: _dafny.Map
            d_2368_v161_ = _dafny.Map({d_2153_v0_: d_2367_v160_})
            index401_ = default__.safeIndex(665, (d_2153_v0_).length(0))
            (d_2153_v0_)[index401_] = len(d_2368_v161_)
            d_2369_v162_: _dafny.Map
            d_2369_v162_ = _dafny.Map({p3: (self).f15})
            d_2370_v163_: _dafny.Seq
            d_2370_v163_ = _dafny.SeqWithoutIsStrInference([(True if ((d_2369_v162_)[(self).f15] if ((self).f15) in (d_2369_v162_) else False) else not(not(p3)))])
            d_2371_v164_: _dafny.Array
            nw356_ = _dafny.Array(None, 9)
            nw356_[int(0)] = d_2365_v158_
            nw356_[int(1)] = _dafny.Map({p1: not((self).f15)})
            nw356_[int(2)] = (d_2365_v158_).set(len(d_2369_v162_), p3)
            nw356_[int(3)] = d_2365_v158_
            nw356_[int(4)] = d_2365_v158_
            nw356_[int(5)] = _dafny.Map({(self).f14: p3})
            nw356_[int(6)] = d_2365_v158_
            nw356_[int(7)] = d_2365_v158_
            nw356_[int(8)] = (d_2365_v158_).set((self).f11, (self).f15)
            d_2371_v164_ = nw356_
            d_2372_v165_: _dafny.Array
            nw357_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_2372_v165_ = nw357_
            d_2373_v166_: _dafny.Map
            d_2373_v166_ = _dafny.Map({d_2372_v165_: d_2371_v164_})
            index402_ = default__.safeIndex(665, (d_2153_v0_).length(0))
            index403_ = default__.safeIndex(665, (d_2153_v0_).length(0))
            rhs389_ = (self.f16).set(default__.safeIndex(len((d_2365_v158_ if not(p3) else d_2365_v158_)), len(self.f16)), (d_2367_v160_ if p3 else d_2367_v160_))
            rhs390_ = default__.safeModuloInt(default__.safeDivisionInt((d_2212_v40_)[default__.safeIndex(-719, len(d_2212_v40_))], p1), (d_2153_v0_)[default__.safeIndex(665, (d_2153_v0_).length(0))])
            rhs391_ = (d_2370_v163_) + (d_2370_v163_)
            rhs392_ = ((d_2373_v166_)[d_2372_v165_] if (d_2372_v165_) in (d_2373_v166_) else d_2371_v164_)
            rhs393_ = 611
            lhs322_ = self
            lhs323_ = d_2153_v0_
            lhs324_ = default__.safeIndex(665, (d_2153_v0_).length(0))
            lhs325_ = d_2153_v0_
            lhs326_ = default__.safeIndex(665, (d_2153_v0_).length(0))
            lhs322_.f16 = rhs389_
            lhs323_[lhs324_] = rhs390_
            d_2370_v163_ = rhs391_
            d_2371_v164_ = rhs392_
            lhs325_[lhs326_] = rhs393_
            d_2374_v167_: C2
            nw358_ = C2()
            nw358_.ctor__(487, _dafny.SeqWithoutIsStrInference([d_2367_v160_ for d_2375_i21_ in range(default__.abs(568))]), p1, (p1) - (p1))
            d_2374_v167_ = nw358_

    def m16(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: int = int(0)
        r2: bool = False
        r3: int = int(0)
        d_2376_v0_: _dafny.Array
        nw359_ = _dafny.Array(False, 20)
        d_2376_v0_ = nw359_
        d_2377_v1_: D19
        d_2377_v1_ = D19_DC55(_dafny.Set({d_2376_v0_, d_2376_v0_}))
        source36_ = d_2377_v1_
        if source36_.is_DC56:
            d_2378___mcc_h0_ = source36_.cf105
            d_2379___mcc_h1_ = source36_.cf106
            d_2380___mcc_h2_ = source36_.cf107
            d_2381___mcc_h3_ = source36_.cf108
            d_2382_cf108_ = d_2381___mcc_h3_
            d_2383_cf107_ = d_2380___mcc_h2_
            d_2384_cf106_ = d_2379___mcc_h1_
            d_2385_cf105_ = d_2378___mcc_h0_
            d_2386_v2_: _dafny.Seq
            d_2386_v2_ = _dafny.SeqWithoutIsStrInference([p2, (self).f13])
            d_2386_v2_ = d_2386_v2_
            r2 = p1
            d_2387_v3_: D5
            d_2387_v3_ = D5_DC18(_dafny.Map({d_2376_v0_: 528}), (self).f11)
            d_2388_v4_: _dafny.MultiSet
            d_2388_v4_ = _dafny.MultiSet([d_2383_cf107_, (self).f12, default__.fm0(d_2384_cf106_, globalState), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1, p1]))).cardinality, (d_2387_v3_).cf36])
            d_2389_v5_: _dafny.Seq
            d_2389_v5_ = _dafny.SeqWithoutIsStrInference([(self).f14, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdu"))), (self).f12, 151])
            d_2390_v6_: _dafny.Set
            d_2390_v6_ = _dafny.Set({d_2383_cf107_})
            (globalState).f7 = (self).fm24(d_2383_cf107_, (_dafny.MultiSet([len(d_2389_v5_), len(d_2390_v6_)])).ispropersubset(d_2388_v4_), globalState)
            d_2391_v7_: _dafny.Map
            d_2391_v7_ = _dafny.Map({d_2384_cf106_: (self).f12})
            (globalState).f6 = ((d_2391_v7_)[(p1) == (p1)] if ((p1) == (p1)) in (d_2391_v7_) else len(d_2390_v6_))
        elif True:
            d_2392___mcc_h4_ = source36_.cf104
            d_2393_cf104_ = d_2392___mcc_h4_
            if not((p1 if ((self).fm2(366, (self).f15, 750, False, globalState)) > (120) else (self).f15)):
                d_2394_v8_: _dafny.Map
                d_2394_v8_ = _dafny.Map({(0) - ((self).f12): (self).f12})
                d_2394_v8_ = (d_2394_v8_).set((self).f14, 712)
                (globalState).f7 = p1
                d_2395_v9_: str
                d_2395_v9_ = _dafny.CodePoint('u')
                d_2396_v10_: _dafny.Seq
                d_2396_v10_ = _dafny.SeqWithoutIsStrInference([False, (self).f15, ((self).f13) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbhdeamr"))).set(default__.safeIndex((self).f11, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbhdeamr")))), d_2395_v9_))])
                d_2396_v10_ = _dafny.SeqWithoutIsStrInference([True, p1, p1])
                d_2397_v11_: _dafny.Array
                def lambda89_(d_2398_i0_):
                    return (d_2398_i0_) - ((0) - ((self).f14))

                init54_ = lambda89_
                nw360_ = _dafny.Array(None, 4)
                for i0_54_ in range(nw360_.length(0)):
                    nw360_[i0_54_] = init54_(i0_54_)
                d_2397_v11_ = nw360_
                d_2397_v11_ = p0
                d_2399_v13_: _dafny.Array
                def lambda90_(d_2400_v8_):
                    def lambda91_(d_2401_i1_):
                        def iife134_():
                            coll56_ = _dafny.Set()
                            compr_56_: int
                            for compr_56_ in _dafny.IntegerRange(-112, 873):
                                d_2402_v12_: int = compr_56_
                                if ((-112) <= (d_2402_v12_)) and ((d_2402_v12_) < (873)):
                                    coll56_ = coll56_.union(_dafny.Set([(d_2402_v12_) + ((self).f11)]))
                            return _dafny.Set(coll56_)
                        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([len((self).f13), len(iife134_()
                        ), (self).f11, (self).f11, (self).f14])) + (_dafny.SeqWithoutIsStrInference([((d_2400_v8_)[(self).f11] if ((self).f11) in (d_2400_v8_) else (self).f14) for d_2403_i2_ in range(default__.abs(-337))])))

                    return lambda91_

                init55_ = lambda90_(d_2394_v8_)
                nw361_ = _dafny.Array(None, 25)
                for i0_55_ in range(nw361_.length(0)):
                    nw361_[i0_55_] = init55_(i0_55_)
                d_2399_v13_ = nw361_
                d_2404_v14_: _dafny.Map
                d_2404_v14_ = _dafny.Map({not(p1): (self).f14})
                d_2405_v15_: _dafny.MultiSet
                d_2405_v15_ = _dafny.MultiSet([865, ((d_2404_v14_)[True] if (True) in (d_2404_v14_) else (self).f11), len(p2)])
                index404_ = default__.safeIndex(675, (d_2399_v13_).length(0))
                (d_2399_v13_)[index404_] = d_2405_v15_
                d_2406_v16_: _dafny.Map
                d_2406_v16_ = _dafny.Map({(self).f15: (self).f15})
                index405_ = default__.safeIndex(675, (d_2399_v13_).length(0))
                rhs394_ = d_2405_v15_
                rhs395_ = (self).fm2(default__.safeDivisionInt((self).f11, (self).f12), ((d_2406_v16_)[p1] if (p1) in (d_2406_v16_) else p1), ((self).f14) + ((self).f12), False, globalState)
                rhs396_ = (self).f14
                lhs327_ = d_2399_v13_
                lhs328_ = default__.safeIndex(675, (d_2399_v13_).length(0))
                lhs329_ = globalState
                lhs330_ = globalState
                lhs327_[lhs328_] = rhs394_
                lhs329_.f6 = rhs395_
                lhs330_.f6 = rhs396_
            elif True:
                d_2407_v17_: _dafny.Map
                d_2407_v17_ = _dafny.Map({p1: (self).f12})
                d_2407_v17_ = (d_2407_v17_) | ((d_2407_v17_).set(p1, len(_dafny.SeqWithoutIsStrInference([d_2407_v17_, d_2407_v17_]))))
                (globalState).f7 = (self).f15
                (globalState).f2 = not(p1)
                (globalState).f7 = (self).f15
                (globalState).f7 = p1
            d_2408_v18_: _dafny.Seq
            d_2408_v18_ = _dafny.SeqWithoutIsStrInference([((self).fm1((self).f14, (self).f15, True, globalState)) or (p1)])
            d_2408_v18_ = (d_2408_v18_) + ((_dafny.SeqWithoutIsStrInference([(self).f15])) + (d_2408_v18_))
            r3 = (self).f11
            d_2409_v19_: _dafny.Seq
            d_2409_v19_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_2410_v20_: str
            d_2410_v20_ = _dafny.CodePoint('l')
            d_2411_v21_: _dafny.MultiSet
            d_2411_v21_ = _dafny.MultiSet([p2, ((d_2409_v19_)[default__.safeIndex((_dafny.MultiSet([(self).f12])).cardinality, len(d_2409_v19_))]) + ((self).f13), ((self).f13 if p1 else (self).f13), default__.fm31(d_2410_v20_, D0_DC0(default__.fm62(p1, globalState)), globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmr"))])
            d_2412_v22_: _dafny.Array
            nw362_ = _dafny.Array(_dafny.Seq({}), 28)
            d_2412_v22_ = nw362_
            index406_ = default__.safeIndex(361, (d_2412_v22_).length(0))
            (d_2412_v22_)[index406_] = d_2408_v18_
            d_2413_v23_: _dafny.Map
            d_2413_v23_ = _dafny.Map({p1: (self).f15})
            d_2414_v24_: _dafny.Map
            d_2414_v24_ = _dafny.Map({len(d_2413_v23_): ((self).f15) or (p1)})
            index407_ = default__.safeIndex(361, (d_2412_v22_).length(0))
            rhs397_ = d_2411_v21_
            rhs398_ = d_2408_v18_
            rhs399_ = not(not((default__.fm39(globalState)).cf16))
            rhs400_ = ((d_2414_v24_)[(self).f12] if ((self).f12) in (d_2414_v24_) else (self).fm1((self).f14, not(p1), (self).f15, globalState))
            lhs331_ = d_2412_v22_
            lhs332_ = default__.safeIndex(361, (d_2412_v22_).length(0))
            lhs333_ = globalState
            d_2411_v21_ = rhs397_
            lhs331_[lhs332_] = rhs398_
            lhs333_.f7 = rhs399_
            r2 = rhs400_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_2376_v0_).length(0)):
            d_2415_i3_: int = guard_loop_7_
            if (True) and (((0) <= (d_2415_i3_)) and ((d_2415_i3_) < ((d_2376_v0_).length(0)))):
                (d_2376_v0_)[(d_2415_i3_)] = False
        d_2416_i4_: int
        d_2416_i4_ = 0
        with _dafny.label("15"):
            while p1:
                with _dafny.c_label("15"):
                    if (d_2416_i4_) >= (100):
                        raise _dafny.Break("15")
                    d_2416_i4_ = (d_2416_i4_) + (1)
                    d_2417_v25_: str
                    d_2417_v25_ = _dafny.CodePoint('r')
                    d_2418_v26_: _dafny.Set
                    d_2418_v26_ = _dafny.Set({(((self).f13) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whv")))).set(default__.safeIndex((self).f12, len(((self).f13) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whv"))))), d_2417_v25_), (p2 if p1 else self.f16), ((self).f13 if (self).f15 else self.f16)})
                    d_2419_v27_: _dafny.Map
                    d_2419_v27_ = _dafny.Map({(0) - ((self).f11): d_2418_v26_})
                    d_2418_v26_ = ((d_2419_v27_)[(len(p2)) + ((self).f12)] if ((len(p2)) + ((self).f12)) in (d_2419_v27_) else d_2418_v26_)
                    d_2420_v28_: _dafny.Seq
                    d_2420_v28_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14])
                    d_2421_v29_: _dafny.Set
                    d_2421_v29_ = _dafny.Set({(d_2420_v28_)[default__.safeIndex((self).f11, len(d_2420_v28_))], (self).f12, (self).f14})
                    d_2422_v30_: _dafny.Map
                    d_2422_v30_ = _dafny.Map({p2: (d_2421_v29_) == (d_2421_v29_)})
                    d_2422_v30_ = (d_2422_v30_) | (d_2422_v30_)
                    pat_let_tv39_ = p1
                    index408_ = default__.safeIndex(55, (d_2376_v0_).length(0))
                    def iife135_(_pat_let39_0):
                        def iife136_(d_2423_dt__update__tmp_h0_):
                            def iife137_(_pat_let40_0):
                                def iife138_(d_2424_dt__update_hcf97_h0_):
                                    def iife139_(_pat_let41_0):
                                        def iife140_(d_2425_dt__update_hcf96_h0_):
                                            return D18_DC52(d_2425_dt__update_hcf96_h0_, d_2424_dt__update_hcf97_h0_, (d_2423_dt__update__tmp_h0_).cf98)
                                        return iife140_(_pat_let41_0)
                                    return iife139_(False)
                                return iife138_(_pat_let40_0)
                            return iife137_(pat_let_tv39_)
                        return iife136_(_pat_let39_0)
                    (d_2376_v0_)[index408_] = (self.f16) <= ((iife135_(D18_DC52((self).f15, p1, (self).f13))).cf98)
                    index409_ = default__.safeIndex(55, (d_2376_v0_).length(0))
                    (d_2376_v0_)[index409_] = (self).fm24((0) - (default__.safeModuloInt((self).f11, (self).f12)), True, globalState)
                    d_2426_v31_: _dafny.Set
                    d_2426_v31_ = _dafny.Set({(self).f15})
                    d_2427_v32_: _dafny.Seq
                    d_2427_v32_ = _dafny.SeqWithoutIsStrInference([True])
                    d_2428_v33_: _dafny.MultiSet
                    d_2428_v33_ = _dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f14, len(d_2426_v31_), len(d_2427_v32_), (self).f12]))).cardinality, (self).f12, (self).f11, (self).f12, (self).f14])
                    d_2429_v34_: _dafny.Map
                    d_2429_v34_ = _dafny.Map({(d_2428_v33_).ispropersubset((_dafny.MultiSet(d_2420_v28_)).set((self).f11, default__.abs((self).f14))): (self).f11})
                    d_2429_v34_ = (d_2429_v34_).set(p1, (self).f11)
                    pass
            pass
        hi8_ = (self).f12
        for d_2430_i5_ in range(((self).f12) - (435), hi8_):
            d_2431_v35_: bool
            d_2432_v36_: _dafny.Seq
            out71_: bool
            out72_: _dafny.Seq
            out71_, out72_ = default__.m0((default__.fm0((self).f15, globalState)) >= ((self).f12), p1, (self).f14, len(p2), globalState)
            d_2431_v35_ = out71_
            d_2432_v36_ = out72_
            (globalState).f7 = d_2431_v35_
            r3 = ((self).f14 if p1 else (self).f14)
            (self).f16 = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2433_i6_ in range(default__.abs(312))])) + (p2)
        d_2434_v37_: _dafny.Set
        d_2434_v37_ = _dafny.Set({default__.fm0(p1, globalState)})
        if (d_2434_v37_).isdisjoint((d_2434_v37_) | (_dafny.Set({(self).f12, (self).f12, len((_dafny.SeqWithoutIsStrInference([p0])).set(default__.safeIndex((self).f11, len(_dafny.SeqWithoutIsStrInference([p0]))), p0))}))):
            if (d_2434_v37_).issubset(d_2434_v37_):
                d_2435_v38_: _dafny.Seq
                d_2435_v38_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f12, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2436_i7_ in range(default__.abs(537))])) + ((self).f13)), default__.safeModuloInt((self).f12, (self).f11), (self).f11])
                (globalState).f6 = (d_2435_v38_)[default__.safeIndex((self).f14, len(d_2435_v38_))]
                d_2437_v39_: str
                d_2437_v39_ = _dafny.CodePoint('p')
                d_2438_v40_: _dafny.Map
                d_2438_v40_ = _dafny.Map({d_2437_v39_: (self).f14})
                d_2439_v41_: C1
                nw363_ = C1()
                nw363_.ctor__(d_2376_v0_, ((d_2438_v40_)[d_2437_v39_] if (d_2437_v39_) in (d_2438_v40_) else (self).f11), 93, (self).f14, (self).fm1((0) - ((self).f11), p1, p1, globalState), p2)
                d_2439_v41_ = nw363_
                d_2440_v42_: _dafny.Array
                def lambda92_(d_2441_i8_):
                    return _dafny.SeqWithoutIsStrInference([(self).f13, (self).f13])

                init56_ = lambda92_
                nw364_ = _dafny.Array(None, 23)
                for i0_56_ in range(nw364_.length(0)):
                    nw364_[i0_56_] = init56_(i0_56_)
                d_2440_v42_ = nw364_
                d_2442_v43_: _dafny.Seq
                d_2442_v43_ = _dafny.SeqWithoutIsStrInference([(self).f13, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eab"))])
                index410_ = default__.safeIndex(913, (d_2440_v42_).length(0))
                (d_2440_v42_)[index410_] = d_2442_v43_
                index411_ = default__.safeIndex(913, (d_2440_v42_).length(0))
                (d_2440_v42_)[index411_] = d_2442_v43_
                d_2443_v44_: _dafny.Set
                d_2443_v44_ = _dafny.Set({p1})
                rhs401_ = (d_2443_v44_ if p1 else d_2443_v44_)
                rhs402_ = ((self).f11) * ((self).f14)
                rhs403_ = p1
                r0 = rhs401_
                r1 = rhs402_
                r2 = rhs403_
                d_2444_v45_: _dafny.Map
                d_2444_v45_ = _dafny.Map({((self).f14) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))): len(p2)})
                (globalState).f6 = ((d_2444_v45_)[(self).f14] if ((self).f14) in (d_2444_v45_) else -49)
            elif True:
                d_2445_v46_: _dafny.Array
                nw365_ = _dafny.Array(_dafny.Seq({}), 14)
                d_2445_v46_ = nw365_
                d_2446_v47_: _dafny.Seq
                d_2446_v47_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])
                index412_ = default__.safeIndex(621, (d_2445_v46_).length(0))
                (d_2445_v46_)[index412_] = d_2446_v47_
                d_2447_v48_: _dafny.Seq
                d_2447_v48_ = _dafny.SeqWithoutIsStrInference([(d_2446_v47_) + (d_2446_v47_), d_2446_v47_])
                index413_ = default__.safeIndex(621, (d_2445_v46_).length(0))
                rhs404_ = d_2376_v0_
                rhs405_ = (self).f11
                rhs406_ = len((d_2447_v48_)[default__.safeIndex((0) - ((self).f11), len(d_2447_v48_))])
                rhs407_ = d_2446_v47_
                rhs408_ = 945
                lhs334_ = globalState
                lhs335_ = globalState
                lhs336_ = d_2445_v46_
                lhs337_ = default__.safeIndex(621, (d_2445_v46_).length(0))
                lhs338_ = globalState
                d_2376_v0_ = rhs404_
                lhs334_.f6 = rhs405_
                lhs335_.f6 = rhs406_
                lhs336_[lhs337_] = rhs407_
                lhs338_.f6 = rhs408_
                d_2376_v0_ = d_2376_v0_
                d_2448_v49_: D18
                d_2448_v49_ = D18_DC52((self).f15, (self).f15, self.f16)
                d_2449_v50_: _dafny.Array
                nw366_ = _dafny.Array(None, 1)
                nw366_[int(0)] = d_2448_v49_
                d_2449_v50_ = nw366_
                d_2449_v50_ = d_2449_v50_
                d_2450_v51_: _dafny.Seq
                d_2450_v51_ = _dafny.SeqWithoutIsStrInference([False, p1, p1])
                d_2451_v52_: _dafny.Array
                def lambda93_(d_2452_p1_):
                    def lambda94_(d_2453_i9_):
                        return D6_DC22(D6_DC21(d_2452_p1_))

                    return lambda94_

                init57_ = lambda93_(p1)
                nw367_ = _dafny.Array(None, 4)
                for i0_57_ in range(nw367_.length(0)):
                    nw367_[i0_57_] = init57_(i0_57_)
                d_2451_v52_ = nw367_
                d_2454_v53_: _dafny.Map
                d_2454_v53_ = _dafny.Map({d_2450_v51_: d_2451_v52_})
                d_2454_v53_ = (d_2454_v53_) | (d_2454_v53_)
                r1 = (self).f14
            d_2455_v54_: _dafny.Array
            nw368_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
            d_2455_v54_ = nw368_
            d_2455_v54_ = d_2455_v54_
            d_2456_v55_: _dafny.Array
            nw369_ = _dafny.Array(None, 10)
            nw369_[int(0)] = d_2376_v0_
            nw369_[int(1)] = d_2376_v0_
            nw369_[int(2)] = (d_2376_v0_ if (self).f15 else d_2376_v0_)
            nw369_[int(3)] = d_2376_v0_
            nw369_[int(4)] = d_2376_v0_
            nw369_[int(5)] = d_2376_v0_
            nw369_[int(6)] = d_2376_v0_
            nw369_[int(7)] = d_2376_v0_
            nw369_[int(8)] = d_2376_v0_
            nw369_[int(9)] = d_2376_v0_
            d_2456_v55_ = nw369_
            index414_ = default__.safeIndex(890, (d_2456_v55_).length(0))
            (d_2456_v55_)[index414_] = d_2376_v0_
            index415_ = default__.safeIndex(890, (d_2456_v55_).length(0))
            (d_2456_v55_)[index415_] = d_2376_v0_
            (globalState).f7 = p1
            d_2457_v56_: _dafny.Seq
            d_2457_v56_ = _dafny.SeqWithoutIsStrInference([p1, not(p1)])
            d_2458_v57_: _dafny.Seq
            d_2458_v57_ = _dafny.SeqWithoutIsStrInference([p1, (d_2457_v56_)[default__.safeIndex((self).f14, len(d_2457_v56_))]])
            d_2459_v58_: D14
            d_2459_v58_ = D14_DC40(d_2458_v57_)
            if (((self).f11) + ((self).f14)) != ((0) - (default__.safeModuloInt((0) - ((self).f12), (0) - (len((d_2459_v58_).cf72))))):
                d_2460_v59_: _dafny.MultiSet
                d_2460_v59_ = _dafny.MultiSet([p2])
                d_2461_v60_: D20
                d_2461_v60_ = D20_DC57(d_2460_v59_)
                r2 = ((d_2461_v60_).cf109).isdisjoint(d_2460_v59_)
                (globalState).f2 = p1
                (globalState).f6 = ((self).f11) * (((self).f14) + ((self).f12))
                index416_ = default__.safeIndex(449, (p0).length(0))
                (p0)[index416_] = (self).f11
                index417_ = default__.safeIndex(449, (p0).length(0))
                (p0)[index417_] = (self).f12
                d_2456_v55_ = d_2456_v55_
            elif True:
                d_2462_v61_: _dafny.Map
                d_2462_v61_ = _dafny.Map({len((self.f16) + (self.f16)): (self).f15})
                d_2463_v62_: _dafny.Map
                d_2463_v62_ = _dafny.Map({(self).f12: 134})
                d_2464_v63_: _dafny.Seq
                d_2464_v63_ = _dafny.SeqWithoutIsStrInference([367])
                d_2462_v61_ = (d_2462_v61_).set(((d_2463_v62_)[len(default__.fm35(not((self).f15), d_2464_v63_, globalState))] if (len(default__.fm35(not((self).f15), d_2464_v63_, globalState))) in (d_2463_v62_) else (self).f14), (len((d_2463_v62_).set((self).f14, (self).f14))) == (len(d_2464_v63_)))
                (globalState).f2 = ((self).f13) == (self.f16)
                d_2465_v64_: _dafny.Array
                nw370_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_2465_v64_ = nw370_
                d_2466_v65_: _dafny.Seq
                d_2466_v65_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2467_v66_: _dafny.Array
                d_2467_v66_ = (d_2466_v65_)[default__.safeIndex((self).f14, len(d_2466_v65_))]
                index418_ = default__.safeIndex(749, (d_2465_v64_).length(0))
                (d_2465_v64_)[index418_] = d_2467_v66_
                index419_ = default__.safeIndex(749, (d_2465_v64_).length(0))
                (d_2465_v64_)[index419_] = d_2467_v66_
                (globalState).f6 = len(p2)
                d_2468_v67_: _dafny.Array
                nw371_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_2468_v67_ = nw371_
                d_2469_v68_: D10
                d_2469_v68_ = D10_DC30(p1, p1, (self).f13, p1)
                d_2470_v69_: D20
                d_2470_v69_ = D20_DC58((d_2469_v68_).cf51)
                d_2471_v70_: _dafny.Array
                nw372_ = _dafny.Array(None, 5)
                nw372_[int(0)] = (p3).set((self).f15, default__.abs((self).f14))
                nw372_[int(1)] = p3
                nw372_[int(2)] = _dafny.MultiSet([True, (d_2470_v69_).cf110])
                nw372_[int(3)] = p3
                nw372_[int(4)] = _dafny.MultiSet([p1, (self).f15])
                d_2471_v70_ = nw372_
                index420_ = default__.safeIndex(502, (d_2468_v67_).length(0))
                (d_2468_v67_)[index420_] = d_2471_v70_
                index421_ = default__.safeIndex(502, (d_2468_v67_).length(0))
                (d_2468_v67_)[index421_] = d_2471_v70_
        elif True:
            (globalState).f6 = (self).f12
            d_2472_v71_: str
            d_2472_v71_ = _dafny.CodePoint('h')
            d_2473_v72_: _dafny.MultiSet
            d_2473_v72_ = _dafny.MultiSet([(self.f16)[default__.safeIndex((self).f14, len(self.f16))], d_2472_v71_])
            d_2473_v72_ = _dafny.MultiSet((self.f16) + (_dafny.SeqWithoutIsStrInference([d_2472_v71_])))
            d_2474_v73_: _dafny.Array
            nw373_ = _dafny.Array(_dafny.Map({}), 28)
            d_2474_v73_ = nw373_
            d_2475_v74_: _dafny.Map
            d_2475_v74_ = _dafny.Map({d_2474_v73_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjxwipci")))})
            d_2475_v74_ = (d_2475_v74_).set(d_2474_v73_, (self).f12)
            d_2476_v75_: D6
            d_2476_v75_ = D6_DC21((self).f15)
            index422_ = default__.safeIndex(649, (d_2376_v0_).length(0))
            (d_2376_v0_)[index422_] = (d_2476_v75_).cf40
            d_2477_v77_: _dafny.Seq
            d_2477_v77_ = _dafny.SeqWithoutIsStrInference([(self).f12])
            d_2478_v78_: _dafny.Seq
            d_2478_v78_ = _dafny.SeqWithoutIsStrInference([len(d_2477_v77_)])
            d_2479_v79_: _dafny.Map
            d_2479_v79_ = _dafny.Map({d_2478_v78_: (self).f12})
            index423_ = default__.safeIndex(649, (d_2376_v0_).length(0))
            def iife141_():
                coll57_ = _dafny.Map()
                compr_57_: _dafny.Seq
                for compr_57_ in (d_2479_v79_).keys.Elements:
                    d_2480_v76_: _dafny.Seq = compr_57_
                    if (d_2480_v76_) in (d_2479_v79_):
                        coll57_[d_2480_v76_] = ((p3)[False] if (False) in (p3) else 45)
                return _dafny.Map(coll57_)
            (d_2376_v0_)[index423_] = (len(iife141_()
            )) == ((self).f14)
            (globalState).f6 = default__.safeDivisionInt(len(self.f16), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxurtcjh"))))
        d_2481_v80_: _dafny.Array
        nw374_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_2481_v80_ = nw374_
        index424_ = default__.safeIndex(318, (d_2481_v80_).length(0))
        (d_2481_v80_)[index424_] = p0
        index425_ = default__.safeIndex(318, (d_2481_v80_).length(0))
        nw375_ = _dafny.Array(int(0), 21)
        (d_2481_v80_)[index425_] = nw375_
        d_2482_v81_: _dafny.Set
        d_2482_v81_ = _dafny.Set({p1})
        d_2483_v82_: D12
        d_2483_v82_ = D12_DC35(d_2482_v81_)
        r0 = (d_2483_v82_).cf63
        r1 = (default__.fm63(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vei")), p1, 296, globalState)).cf93
        r2 = (self).fm24(default__.safeDivisionInt((self).f12, (self).f14), ((self).f14) >= ((self).f12), globalState)
        r3 = (default__.fm64((self).f14, globalState)).cardinality
        return r0, r1, r2, r3


class C5:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self):
        pass
        pass

    def fm13(self, p0, p1, p2, globalState):
        if False:
            return (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))]))) > ((_dafny.MultiSet([not(True)])).cardinality)
        elif True:
            return False

    def m10(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        d_2484_v0_: str
        d_2484_v0_ = _dafny.CodePoint('m')
        d_2485_v1_: _dafny.Seq
        d_2485_v1_ = _dafny.SeqWithoutIsStrInference([d_2484_v0_])
        r1 = ((d_2485_v1_) + (d_2485_v1_)) != (d_2485_v1_)
        if (self).fm13(p0, d_2484_v0_, d_2484_v0_, globalState):
            (globalState).f6 = (0) - (default__.fm0(not(not(p2)), globalState))
            d_2486_v2_: D1
            d_2486_v2_ = D1_DC2(d_2484_v0_)
            d_2487_v3_: _dafny.Map
            d_2487_v3_ = _dafny.Map({p0: d_2484_v0_})
            d_2488_v4_: _dafny.MultiSet
            d_2488_v4_ = _dafny.MultiSet([False])
            d_2489_v5_: _dafny.Map
            d_2489_v5_ = _dafny.Map({p2: d_2484_v0_})
            d_2490_v6_: _dafny.Array
            nw376_ = _dafny.Array(None, 17)
            nw376_[int(0)] = d_2484_v0_
            nw376_[int(1)] = (d_2484_v0_ if p2 else _dafny.CodePoint('b'))
            nw376_[int(2)] = (d_2484_v0_ if True else d_2484_v0_)
            nw376_[int(3)] = d_2484_v0_
            nw376_[int(4)] = (d_2486_v2_).cf3
            nw376_[int(5)] = ((d_2487_v3_)[p1] if (p1) in (d_2487_v3_) else d_2484_v0_)
            nw376_[int(6)] = d_2484_v0_
            nw376_[int(7)] = d_2484_v0_
            nw376_[int(8)] = (d_2485_v1_)[default__.safeIndex((d_2488_v4_).cardinality, len(d_2485_v1_))]
            nw376_[int(9)] = d_2484_v0_
            nw376_[int(10)] = d_2484_v0_
            nw376_[int(11)] = default__.fm14(p2, d_2484_v0_, p1, p2, globalState)
            nw376_[int(12)] = default__.fm14(p2, d_2484_v0_, p0, p2, globalState)
            nw376_[int(13)] = d_2484_v0_
            nw376_[int(14)] = d_2484_v0_
            nw376_[int(15)] = ((d_2489_v5_)[p2] if (p2) in (d_2489_v5_) else d_2484_v0_)
            nw376_[int(16)] = d_2484_v0_
            d_2490_v6_ = nw376_
            index426_ = default__.safeIndex(29, (d_2490_v6_).length(0))
            (d_2490_v6_)[index426_] = d_2484_v0_
            d_2491_v7_: _dafny.Array
            def lambda95_(d_2492_v1_, d_2493_p2_):
                def lambda96_(d_2494_i0_):
                    return (_dafny.Map({d_2492_v1_: _dafny.Set({d_2493_p2_})})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yik")): _dafny.Set({d_2493_p2_})}))

                return lambda96_

            init58_ = lambda95_(d_2485_v1_, p2)
            nw377_ = _dafny.Array(None, 28)
            for i0_58_ in range(nw377_.length(0)):
                nw377_[i0_58_] = init58_(i0_58_)
            d_2491_v7_ = nw377_
            d_2495_v8_: _dafny.Map
            d_2495_v8_ = _dafny.Map({p2: (self).fm13(p0, _dafny.CodePoint('t'), d_2484_v0_, globalState)})
            d_2496_v9_: _dafny.Set
            d_2496_v9_ = _dafny.Set({p2})
            index427_ = default__.safeIndex(667, (d_2491_v7_).length(0))
            (d_2491_v7_)[index427_] = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yg"))).set(default__.safeIndex(len(d_2495_v8_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yg")))), d_2484_v0_): d_2496_v9_})
            d_2497_v10_: D0
            d_2497_v10_ = D0_DC1(p2, p0)
            index428_ = default__.safeIndex(29, (d_2490_v6_).length(0))
            index429_ = default__.safeIndex(667, (d_2491_v7_).length(0))
            rhs409_ = d_2484_v0_
            rhs410_ = default__.fm15(p0, (d_2497_v10_).cf2, globalState)
            lhs339_ = d_2490_v6_
            lhs340_ = default__.safeIndex(29, (d_2490_v6_).length(0))
            lhs341_ = d_2491_v7_
            lhs342_ = default__.safeIndex(667, (d_2491_v7_).length(0))
            lhs339_[lhs340_] = rhs409_
            lhs341_[lhs342_] = rhs410_
            (globalState).f2 = not((p1) > (p0))
            d_2498_v11_: _dafny.Set
            d_2498_v11_ = _dafny.Set({d_2484_v0_, (d_2490_v6_)[default__.safeIndex(29, (d_2490_v6_).length(0))]})
            d_2499_v12_: _dafny.Array
            nw378_ = _dafny.Array(int(0), 23)
            d_2499_v12_ = nw378_
            d_2500_v13_: _dafny.Map
            d_2500_v13_ = _dafny.Map({(d_2498_v11_).issubset(_dafny.Set({d_2484_v0_, d_2484_v0_, (d_2490_v6_)[default__.safeIndex(29, (d_2490_v6_).length(0))]})): d_2499_v12_})
            d_2500_v13_ = ((d_2500_v13_) | (d_2500_v13_)) | (_dafny.Map({p2: d_2499_v12_}))
            if p2:
                d_2501_v14_: _dafny.Map
                d_2501_v14_ = _dafny.Map({(0) - (p1): p0})
                (globalState).f7 = not((p1) >= (((d_2501_v14_)[p1] if (p1) in (d_2501_v14_) else p1)))
                d_2502_v15_: bool
                d_2503_v16_: _dafny.Array
                d_2504_v17_: _dafny.MultiSet
                out73_: bool
                out74_: _dafny.Array
                out75_: _dafny.MultiSet
                out73_, out74_, out75_ = (self).m11(globalState)
                d_2502_v15_ = out73_
                d_2503_v16_ = out74_
                d_2504_v17_ = out75_
                d_2505_v18_: C0
                nw379_ = C0()
                nw379_.ctor__()
                d_2505_v18_ = nw379_
                d_2485_v1_ = ((d_2485_v1_) + (d_2485_v1_)) + (d_2485_v1_)
                (globalState).f6 = p0
            elif True:
                d_2506_v19_: _dafny.Map
                d_2506_v19_ = _dafny.Map({len(d_2485_v1_): d_2485_v1_})
                d_2506_v19_ = d_2506_v19_
                d_2485_v1_ = d_2485_v1_
                (globalState).f2 = p2
                d_2507_v20_: D0
                d_2507_v20_ = D0_DC0(_dafny.MultiSet([(d_2490_v6_)[default__.safeIndex(29, (d_2490_v6_).length(0))]]))
                d_2508_v21_: _dafny.Set
                d_2508_v21_ = _dafny.Set({d_2507_v20_, d_2507_v20_})
                d_2509_v22_: D3
                d_2509_v22_ = D3_DC11((0) - (p0), D2_DC8(p2, p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lc"))), d_2508_v21_, p2)
                d_2510_v23_: _dafny.Seq
                d_2510_v23_ = _dafny.SeqWithoutIsStrInference([d_2489_v5_, _dafny.Map({p2: _dafny.CodePoint('k')}), d_2489_v5_])
                d_2511_v24_: _dafny.Map
                d_2511_v24_ = _dafny.Map({d_2509_v22_: d_2510_v23_})
                d_2511_v24_ = (d_2511_v24_).set(d_2509_v22_, _dafny.SeqWithoutIsStrInference([d_2489_v5_ for d_2512_i1_ in range(default__.abs(897))]))
                (globalState).f2 = p2
        elif True:
            (globalState).f6 = (909 if p2 else len(d_2485_v1_))
            (globalState).f2 = p2
            d_2513_v25_: _dafny.Map
            d_2513_v25_ = _dafny.Map({False: p1})
            d_2514_v26_: _dafny.Set
            d_2514_v26_ = _dafny.Set({p2})
            d_2515_v27_: _dafny.Map
            d_2515_v27_ = _dafny.Map({d_2484_v0_: p2})
            d_2516_v28_: _dafny.Array
            nw380_ = _dafny.Array(None, 15)
            nw380_[int(0)] = len((_dafny.Map({p1: p2})) | (_dafny.Map({len(d_2513_v25_): p2})))
            nw380_[int(1)] = (len(d_2485_v1_)) + (p0)
            nw380_[int(2)] = p0
            nw380_[int(3)] = default__.safeModuloInt(p1, p0)
            nw380_[int(4)] = (0) - (default__.fm0(p2, globalState))
            nw380_[int(5)] = (0) - (len(d_2514_v26_))
            nw380_[int(6)] = default__.safeDivisionInt((0) - (p1), default__.fm0(p2, globalState))
            nw380_[int(7)] = (0) - (len((_dafny.Map({p2: len(d_2515_v27_)})) | (_dafny.Map({False: p1}))))
            nw380_[int(8)] = p1
            nw380_[int(9)] = p1
            nw380_[int(10)] = p0
            nw380_[int(11)] = len(d_2485_v1_)
            nw380_[int(12)] = p1
            nw380_[int(13)] = p0
            nw380_[int(14)] = default__.safeDivisionInt(p1, 197)
            d_2516_v28_ = nw380_
            index430_ = default__.safeIndex(547, (d_2516_v28_).length(0))
            (d_2516_v28_)[index430_] = p1
            index431_ = default__.safeIndex(547, (d_2516_v28_).length(0))
            (d_2516_v28_)[index431_] = p0
            d_2517_v29_: _dafny.Seq
            d_2517_v29_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(225, -413)])
            d_2518_v30_: _dafny.Set
            d_2518_v30_ = _dafny.Set({default__.safeDivisionInt(300, (0) - (p1))})
            d_2519_v31_: D1
            d_2519_v31_ = D1_DC2((d_2485_v1_)[default__.safeIndex(p0, len(d_2485_v1_))])
            index432_ = default__.safeIndex(547, (d_2516_v28_).length(0))
            index433_ = default__.safeIndex(547, (d_2516_v28_).length(0))
            rhs411_ = len(d_2518_v30_)
            rhs412_ = (self).fm13((p0) - (p1), (d_2519_v31_).cf3, d_2484_v0_, globalState)
            rhs413_ = (0) - ((d_2516_v28_)[default__.safeIndex(547, (d_2516_v28_).length(0))])
            rhs414_ = (D4_DC13(d_2517_v29_)).cf29
            lhs343_ = d_2516_v28_
            lhs344_ = default__.safeIndex(547, (d_2516_v28_).length(0))
            lhs345_ = d_2516_v28_
            lhs346_ = default__.safeIndex(547, (d_2516_v28_).length(0))
            lhs343_[lhs344_] = rhs411_
            r1 = rhs412_
            lhs345_[lhs346_] = rhs413_
            d_2517_v29_ = rhs414_
            d_2485_v1_ = default__.fm18(not (p2) or (p2), (d_2516_v28_)[default__.safeIndex(547, (d_2516_v28_).length(0))], globalState)
        if p2:
            d_2520_v32_: _dafny.Array
            nw381_ = _dafny.Array(False, 19)
            d_2520_v32_ = nw381_
            index434_ = default__.safeIndex(363, (d_2520_v32_).length(0))
            (d_2520_v32_)[index434_] = (self).fm13((0) - (p1), (d_2485_v1_)[default__.safeIndex(91, len(d_2485_v1_))], _dafny.CodePoint('s'), globalState)
            d_2521_v33_: _dafny.Set
            d_2521_v33_ = _dafny.Set({d_2485_v1_})
            d_2522_v34_: D5
            d_2522_v34_ = D5_DC17(d_2521_v33_)
            index435_ = default__.safeIndex(363, (d_2520_v32_).length(0))
            (d_2520_v32_)[index435_] = ((d_2522_v34_).cf34) != ((d_2521_v33_) | (d_2521_v33_))
            d_2523_v35_: T2
            nw382_ = C1()
            nw382_.ctor__(d_2520_v32_, p0, p0, p0, (d_2520_v32_)[default__.safeIndex(363, (d_2520_v32_).length(0))], d_2485_v1_)
            d_2523_v35_ = nw382_
            d_2524_v36_: _dafny.Array
            nw383_ = _dafny.Array(None, 11)
            nw383_[int(0)] = d_2523_v35_
            nw383_[int(1)] = d_2523_v35_
            nw383_[int(2)] = d_2523_v35_
            nw383_[int(3)] = d_2523_v35_
            nw383_[int(4)] = d_2523_v35_
            nw383_[int(5)] = d_2523_v35_
            nw383_[int(6)] = d_2523_v35_
            nw383_[int(7)] = d_2523_v35_
            nw383_[int(8)] = d_2523_v35_
            nw383_[int(9)] = d_2523_v35_
            nw383_[int(10)] = d_2523_v35_
            d_2524_v36_ = nw383_
            d_2525_v37_: _dafny.Set
            d_2525_v37_ = _dafny.Set({d_2524_v36_})
            d_2526_v38_: D1
            d_2526_v38_ = D1_DC2(d_2484_v0_)
            d_2527_v39_: D2
            d_2527_v39_ = D2_DC7(d_2520_v32_, d_2525_v37_, d_2526_v38_, 102, 506)
            d_2528_v40_: D2
            d_2528_v40_ = D2_DC7(d_2520_v32_, (d_2527_v39_).cf11, d_2526_v38_, (d_2523_v35_).f12, 855)
            d_2529_v41_: _dafny.Map
            d_2529_v41_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puiuopita"))): d_2520_v32_})
            pat_let_tv40_ = d_2526_v38_
            pat_let_tv41_ = d_2525_v37_
            d_2530_v42_: _dafny.Array
            nw384_ = _dafny.Array(None, 23)
            def iife142_(_pat_let42_0):
                def iife143_(d_2531_dt__update__tmp_h0_):
                    def iife144_(_pat_let43_0):
                        def iife145_(d_2532_dt__update_hcf12_h0_):
                            def iife146_(_pat_let44_0):
                                def iife147_(d_2533_dt__update_hcf11_h0_):
                                    return D2_DC7((d_2531_dt__update__tmp_h0_).cf10, d_2533_dt__update_hcf11_h0_, d_2532_dt__update_hcf12_h0_, (d_2531_dt__update__tmp_h0_).cf13, (d_2531_dt__update__tmp_h0_).cf14)
                                return iife147_(_pat_let44_0)
                            return iife146_(pat_let_tv41_)
                        return iife145_(_pat_let43_0)
                    return iife144_(pat_let_tv40_)
                return iife143_(_pat_let42_0)
            nw384_[int(0)] = (iife142_(d_2528_v40_)).cf10
            nw384_[int(1)] = d_2520_v32_
            nw384_[int(2)] = d_2520_v32_
            nw384_[int(3)] = d_2520_v32_
            nw384_[int(4)] = d_2520_v32_
            nw384_[int(5)] = d_2520_v32_
            nw384_[int(6)] = d_2520_v32_
            nw384_[int(7)] = d_2520_v32_
            nw384_[int(8)] = d_2520_v32_
            nw384_[int(9)] = ((d_2529_v41_)[(0) - (p0)] if ((0) - (p0)) in (d_2529_v41_) else d_2520_v32_)
            nw384_[int(10)] = d_2520_v32_
            nw384_[int(11)] = d_2520_v32_
            nw384_[int(12)] = d_2520_v32_
            nw384_[int(13)] = d_2520_v32_
            nw384_[int(14)] = d_2520_v32_
            nw384_[int(15)] = d_2520_v32_
            nw384_[int(16)] = d_2520_v32_
            nw384_[int(17)] = d_2520_v32_
            nw384_[int(18)] = d_2520_v32_
            nw384_[int(19)] = d_2520_v32_
            nw384_[int(20)] = d_2520_v32_
            nw384_[int(21)] = d_2520_v32_
            nw384_[int(22)] = d_2520_v32_
            d_2530_v42_ = nw384_
            d_2530_v42_ = d_2530_v42_
            if (p1) != (-868):
                d_2534_v43_: bool
                d_2535_v44_: _dafny.Array
                d_2536_v45_: _dafny.MultiSet
                out76_: bool
                out77_: _dafny.Array
                out78_: _dafny.MultiSet
                out76_, out77_, out78_ = (self).m11(globalState)
                d_2534_v43_ = out76_
                d_2535_v44_ = out77_
                d_2536_v45_ = out78_
                d_2537_v46_: _dafny.Seq
                d_2537_v46_ = _dafny.SeqWithoutIsStrInference([d_2535_v44_])
                d_2538_v47_: _dafny.Map
                d_2538_v47_ = _dafny.Map({(d_2523_v35_).f15: len(d_2537_v46_)})
                d_2539_v48_: _dafny.Set
                d_2539_v48_ = _dafny.Set({d_2538_v47_})
                (globalState).f6 = len((d_2539_v48_) - (d_2539_v48_))
                r1 = ((d_2523_v35_).f12) != (p1)
                d_2540_v49_: _dafny.Map
                d_2540_v49_ = _dafny.Map({p0: (self).fm13((d_2523_v35_).f12, d_2484_v0_, d_2484_v0_, globalState)})
                d_2541_v50_: _dafny.Map
                d_2541_v50_ = _dafny.Map({d_2540_v49_: False})
                (globalState).f7 = ((d_2541_v50_)[_dafny.Map({p0: (d_2520_v32_)[default__.safeIndex(363, (d_2520_v32_).length(0))]})] if (_dafny.Map({p0: (d_2520_v32_)[default__.safeIndex(363, (d_2520_v32_).length(0))]})) in (d_2541_v50_) else d_2534_v43_)
                (globalState).f7 = (d_2484_v0_) not in ((d_2485_v1_) + ((d_2523_v35_).f13))
            elif True:
                d_2542_v51_: _dafny.Map
                d_2542_v51_ = _dafny.Map({(d_2523_v35_).f12: False})
                d_2543_v52_: _dafny.MultiSet
                d_2543_v52_ = _dafny.MultiSet([_dafny.Map({-333: (d_2523_v35_).f15}), d_2542_v51_, d_2542_v51_])
                index436_ = default__.safeIndex(363, (d_2520_v32_).length(0))
                (d_2520_v32_)[index436_] = (d_2543_v52_).issubset(d_2543_v52_)
                d_2544_v53_: _dafny.Seq
                d_2544_v53_ = _dafny.SeqWithoutIsStrInference([(d_2520_v32_)[default__.safeIndex(363, (d_2520_v32_).length(0))]])
                index437_ = default__.safeIndex(363, (d_2520_v32_).length(0))
                (d_2520_v32_)[index437_] = (((d_2544_v53_)[default__.safeIndex((d_2523_v35_).f12, len(d_2544_v53_))]) == (False)) == (not((d_2523_v35_).f15))
                index438_ = default__.safeIndex(363, (d_2520_v32_).length(0))
                (d_2520_v32_)[index438_] = (d_2523_v35_).f15
                (globalState).f2 = (d_2520_v32_)[default__.safeIndex(363, (d_2520_v32_).length(0))]
                d_2520_v32_ = d_2520_v32_
            d_2545_v54_: _dafny.MultiSet
            d_2545_v54_ = _dafny.MultiSet([560])
            d_2545_v54_ = d_2545_v54_
            if p2:
                d_2530_v42_ = d_2530_v42_
                (globalState).f2 = (d_2523_v35_).f15
                d_2484_v0_ = d_2484_v0_
                d_2546_v56_: _dafny.Seq
                d_2546_v56_ = _dafny.SeqWithoutIsStrInference([(d_2520_v32_)[default__.safeIndex(363, (d_2520_v32_).length(0))]])
                def iife148_():
                    coll58_ = _dafny.Set()
                    compr_58_: int
                    for compr_58_ in _dafny.IntegerRange(971, 685):
                        d_2547_v55_: int = compr_58_
                        if ((971) <= (d_2547_v55_)) and ((d_2547_v55_) < (685)):
                            coll58_ = coll58_.union(_dafny.Set([(d_2547_v55_) + ((d_2523_v35_).f11)]))
                    return _dafny.Set(coll58_)
                (globalState).f6 = ((0) - (((d_2523_v35_).f11 if p2 else len(iife148_()
                ))) if ((d_2523_v35_).f13) == ((d_2523_v35_).f13) else default__.safeModuloInt((d_2523_v35_).fm2((d_2523_v35_).f11, (d_2520_v32_)[default__.safeIndex(363, (d_2520_v32_).length(0))], len(d_2546_v56_), (d_2520_v32_)[default__.safeIndex(363, (d_2520_v32_).length(0))], globalState), (d_2523_v35_).f11))
                d_2520_v32_ = (d_2520_v32_ if (d_2520_v32_)[default__.safeIndex(363, (d_2520_v32_).length(0))] else d_2520_v32_)
            elif True:
                d_2548_v57_: _dafny.Array
                nw385_ = _dafny.Array(None, 2)
                nw385_[int(0)] = (d_2523_v35_).f13
                nw385_[int(1)] = (d_2523_v35_).f13
                d_2548_v57_ = nw385_
                index439_ = default__.safeIndex(952, (d_2548_v57_).length(0))
                (d_2548_v57_)[index439_] = _dafny.SeqWithoutIsStrInference([d_2484_v0_ for d_2549_i2_ in range(default__.abs(127))])
                index440_ = default__.safeIndex(952, (d_2548_v57_).length(0))
                (d_2548_v57_)[index440_] = (d_2485_v1_).set(default__.safeIndex((d_2523_v35_).f12, len(d_2485_v1_)), d_2484_v0_)
                d_2550_v58_: _dafny.Set
                d_2550_v58_ = _dafny.Set({(d_2523_v35_).f15})
                d_2551_v59_: _dafny.Array
                nw386_ = _dafny.Array(None, 28)
                nw386_[int(0)] = p1
                nw386_[int(1)] = 744
                nw386_[int(2)] = 148
                nw386_[int(3)] = p0
                nw386_[int(4)] = (d_2523_v35_).f11
                nw386_[int(5)] = p1
                nw386_[int(6)] = (d_2523_v35_).f14
                nw386_[int(7)] = 828
                nw386_[int(8)] = (d_2523_v35_).fm2(len(d_2485_v1_), p2, (d_2523_v35_).f14, (d_2523_v35_).f15, globalState)
                nw386_[int(9)] = p0
                nw386_[int(10)] = p1
                nw386_[int(11)] = p0
                nw386_[int(12)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vahvy"))).set(default__.safeIndex(default__.fm0((d_2523_v35_).f15, globalState), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vahvy")))), d_2484_v0_))
                nw386_[int(13)] = (d_2523_v35_).f12
                nw386_[int(14)] = p0
                nw386_[int(15)] = (d_2523_v35_).f12
                nw386_[int(16)] = (d_2523_v35_).f12
                nw386_[int(17)] = (d_2523_v35_).f14
                nw386_[int(18)] = (d_2523_v35_).f12
                nw386_[int(19)] = p1
                nw386_[int(20)] = (d_2523_v35_).f12
                nw386_[int(21)] = p1
                nw386_[int(22)] = (d_2523_v35_).f14
                nw386_[int(23)] = default__.fm0(p2, globalState)
                nw386_[int(24)] = 866
                nw386_[int(25)] = (d_2523_v35_).f11
                nw386_[int(26)] = 365
                nw386_[int(27)] = len(d_2550_v58_)
                d_2551_v59_ = nw386_
                d_2552_v60_: _dafny.Set
                d_2552_v60_ = _dafny.Set({d_2551_v59_, d_2551_v59_})
                d_2552_v60_ = d_2552_v60_
                (globalState).f6 = (d_2523_v35_).f14
                d_2553_v61_: _dafny.Seq
                d_2553_v61_ = _dafny.SeqWithoutIsStrInference([(d_2523_v35_).f13, (d_2523_v35_).f13])
                d_2554_v62_: _dafny.Map
                d_2554_v62_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")): (d_2523_v35_).f11})
                d_2555_v63_: D3
                d_2555_v63_ = D3_DC10(p2, p0, d_2554_v62_, _dafny.SeqWithoutIsStrInference([d_2484_v0_ for d_2556_i3_ in range(default__.abs(99))]), p0)
                d_2557_v64_: _dafny.Seq
                d_2557_v64_ = _dafny.SeqWithoutIsStrInference([d_2555_v63_, d_2555_v63_, d_2555_v63_, d_2555_v63_])
                d_2553_v61_ = default__.fm23((d_2523_v35_).fm2((d_2523_v35_).f11, p2, (0) - (len(d_2557_v64_)), (d_2523_v35_).f15, globalState), (d_2523_v35_).f11, globalState)
                d_2558_v65_: _dafny.Map
                d_2558_v65_ = _dafny.Map({(0) - (((d_2523_v35_).f14) * ((d_2523_v35_).f11)): default__.fm0((d_2523_v35_).f15, globalState)})
                d_2558_v65_ = (d_2558_v65_) | (d_2558_v65_)
        elif True:
            if (d_2485_v1_) == (d_2485_v1_):
                (globalState).f2 = p2
                d_2559_v66_: _dafny.MultiSet
                d_2559_v66_ = _dafny.MultiSet([p2])
                d_2560_v67_: _dafny.MultiSet
                d_2560_v67_ = _dafny.MultiSet([default__.fm0(p2, globalState), p0, (d_2559_v66_).cardinality, p1])
                d_2561_v68_: _dafny.Seq
                d_2561_v68_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt((0) - (p1), p0), ((d_2560_v67_).intersection(d_2560_v67_)).cardinality])
                (globalState).f6 = (d_2561_v68_)[default__.safeIndex(p1, len(d_2561_v68_))]
                d_2562_v69_: _dafny.Array
                nw387_ = _dafny.Array(False, 29)
                d_2562_v69_ = nw387_
                d_2563_v70_: _dafny.Seq
                d_2563_v70_ = _dafny.SeqWithoutIsStrInference([d_2562_v69_, d_2562_v69_])
                d_2563_v70_ = (d_2563_v70_) + (d_2563_v70_)
                (globalState).f7 = p2
                d_2564_v71_: C0
                nw388_ = C0()
                nw388_.ctor__()
                d_2564_v71_ = nw388_
            elif True:
                d_2565_v72_: _dafny.Array
                def lambda97_(d_2566_p2_):
                    def lambda98_(d_2567_i4_):
                        return d_2566_p2_

                    return lambda98_

                init59_ = lambda97_(p2)
                nw389_ = _dafny.Array(None, 13)
                for i0_59_ in range(nw389_.length(0)):
                    nw389_[i0_59_] = init59_(i0_59_)
                d_2565_v72_ = nw389_
                d_2568_v73_: _dafny.Map
                d_2568_v73_ = _dafny.Map({p2: d_2565_v72_})
                d_2568_v73_ = (d_2568_v73_).set(p2, d_2565_v72_)
                d_2569_v74_: C1
                nw390_ = C1()
                nw390_.ctor__(d_2565_v72_, default__.fm0(p2, globalState), p0, (476) - (p0), (p1) != (-232), d_2485_v1_)
                d_2569_v74_ = nw390_
                d_2570_v75_: _dafny.Set
                d_2570_v75_ = _dafny.Set({d_2484_v0_, d_2484_v0_})
                d_2571_v76_: _dafny.Seq
                d_2571_v76_ = _dafny.SeqWithoutIsStrInference([len(d_2570_v75_)])
                d_2572_v77_: _dafny.Seq
                d_2572_v77_ = _dafny.SeqWithoutIsStrInference([(d_2571_v76_)[default__.safeIndex(843, len(d_2571_v76_))], len((d_2485_v1_).set(default__.safeIndex(p1, len(d_2485_v1_)), d_2484_v0_)), (0) - (p1), p1, p0])
                (globalState).f6 = len(d_2572_v77_)
                d_2573_v78_: C0
                nw391_ = C0()
                nw391_.ctor__()
                d_2573_v78_ = nw391_
                d_2574_v79_: _dafny.Set
                d_2574_v79_ = _dafny.Set({d_2485_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crpofqitv"))})
                d_2575_v80_: D5
                d_2575_v80_ = D5_DC17(d_2574_v79_)
                d_2575_v80_ = d_2575_v80_
            r0 = (_dafny.CodePoint('x')) not in ((d_2485_v1_) + (d_2485_v1_))
            d_2576_v81_: D2
            d_2576_v81_ = D2_DC8(p2, p2, d_2485_v1_)
            d_2577_v82_: _dafny.MultiSet
            d_2577_v82_ = _dafny.MultiSet([d_2484_v0_])
            d_2578_v83_: D0
            d_2578_v83_ = D0_DC0(d_2577_v82_)
            d_2579_v84_: _dafny.Map
            d_2579_v84_ = _dafny.Map({d_2578_v83_: p2})
            d_2580_v86_: D3
            def iife149_():
                coll59_ = _dafny.Set()
                compr_59_: D0
                for compr_59_ in (d_2579_v84_).keys.Elements:
                    d_2581_v85_: D0 = compr_59_
                    if (d_2581_v85_) in (d_2579_v84_):
                        coll59_ = coll59_.union(_dafny.Set([d_2581_v85_]))
                return _dafny.Set(coll59_)
            d_2580_v86_ = D3_DC11((0) - (p0), d_2576_v81_, iife149_()
, p2)
            d_2582_v87_: D3
            d_2582_v87_ = D3_DC12(d_2580_v86_)
            d_2582_v87_ = d_2582_v87_
            d_2583_v88_: _dafny.Array
            nw392_ = _dafny.Array(int(0), 4)
            d_2583_v88_ = nw392_
            d_2584_v89_: _dafny.Seq
            d_2584_v89_ = _dafny.SeqWithoutIsStrInference([d_2583_v88_, d_2583_v88_])
            d_2583_v88_ = (d_2584_v89_)[default__.safeIndex(p1, len(d_2584_v89_))]
            d_2585_v90_: _dafny.Array
            def lambda99_(d_2586_i5_):
                return True

            init60_ = lambda99_
            nw393_ = _dafny.Array(None, 26)
            for i0_60_ in range(nw393_.length(0)):
                nw393_[i0_60_] = init60_(i0_60_)
            d_2585_v90_ = nw393_
            d_2587_v91_: C1
            nw394_ = C1()
            nw394_.ctor__(d_2585_v90_, p0, p1, default__.safeModuloInt(p0, p1), (self).fm13(p0, d_2484_v0_, d_2484_v0_, globalState), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsuvd"))) + (d_2485_v1_))
            d_2587_v91_ = nw394_
        d_2588_v92_: D6
        d_2588_v92_ = D6_DC20(p1, p0)
        d_2589_v93_: D6
        d_2589_v93_ = D6_DC22(d_2588_v92_)
        source37_ = d_2589_v93_
        if source37_.is_DC20:
            d_2590___mcc_h0_ = source37_.cf38
            d_2591___mcc_h1_ = source37_.cf39
            d_2592_cf39_ = d_2591___mcc_h1_
            d_2593_cf38_ = d_2590___mcc_h0_
            d_2594_v94_: _dafny.Array
            nw395_ = _dafny.Array(False, 2)
            d_2594_v94_ = nw395_
            d_2595_v95_: _dafny.Set
            d_2595_v95_ = _dafny.Set({d_2594_v94_, d_2594_v94_})
            (globalState).f2 = ((d_2595_v95_).ispropersubset(_dafny.Set({d_2594_v94_})) if p2 else p2)
            d_2596_v96_: _dafny.Set
            d_2596_v96_ = _dafny.Set({d_2592_cf39_})
            rhs415_ = default__.fm0((self).fm13((0) - (p1), d_2484_v0_, d_2484_v0_, globalState), globalState)
            rhs416_ = ((d_2596_v96_).intersection(d_2596_v96_)).isdisjoint(_dafny.Set({(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_2593_cf38_ for d_2597_i6_ in range(default__.abs(764))])).set(default__.safeIndex(845, len(_dafny.SeqWithoutIsStrInference([d_2593_cf38_ for d_2598_i6_ in range(default__.abs(764))]))), len(d_2596_v96_)))).cardinality, -444}))
            rhs417_ = p1
            lhs347_ = globalState
            d_2592_cf39_ = rhs415_
            lhs347_.f7 = rhs416_
            d_2592_cf39_ = rhs417_
            if False:
                d_2599_v97_: _dafny.Set
                d_2599_v97_ = _dafny.Set({not(p2)})
                (globalState).f2 = ((d_2599_v97_) | (d_2599_v97_)).isdisjoint(d_2599_v97_)
                d_2600_v98_: _dafny.Array
                def lambda100_(d_2601_i7_):
                    return default__.safeDivisionInt(d_2601_i7_, (0) - (-270))

                init61_ = lambda100_
                nw396_ = _dafny.Array(None, 7)
                for i0_61_ in range(nw396_.length(0)):
                    nw396_[i0_61_] = init61_(i0_61_)
                d_2600_v98_ = nw396_
                d_2600_v98_ = d_2600_v98_
                index441_ = default__.safeIndex(252, (d_2594_v94_).length(0))
                (d_2594_v94_)[index441_] = p2
                index442_ = default__.safeIndex(252, (d_2594_v94_).length(0))
                (d_2594_v94_)[index442_] = p2
                d_2485_v1_ = _dafny.SeqWithoutIsStrInference([(d_2484_v0_ if p2 else d_2484_v0_) for d_2602_i8_ in range(default__.abs(69))])
                d_2603_v99_: _dafny.Seq
                d_2603_v99_ = _dafny.SeqWithoutIsStrInference([(d_2594_v94_)[default__.safeIndex(252, (d_2594_v94_).length(0))]])
                d_2604_v100_: _dafny.Map
                d_2604_v100_ = _dafny.Map({d_2603_v99_: (_dafny.SeqWithoutIsStrInference([d_2484_v0_ for d_2605_i9_ in range(default__.abs(418))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "co")))})
                d_2604_v100_ = (d_2604_v100_).set(_dafny.SeqWithoutIsStrInference([p2]), (d_2485_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vj"))))
            elif True:
                r0 = p2
                d_2606_v101_: _dafny.Seq
                d_2606_v101_ = _dafny.SeqWithoutIsStrInference([d_2592_cf39_])
                d_2606_v101_ = d_2606_v101_
                d_2607_v102_: T1
                nw397_ = C4()
                nw397_.ctor__(d_2485_v1_, p0, p2, d_2485_v1_, len(d_2596_v96_), d_2593_cf38_)
                d_2607_v102_ = nw397_
                d_2608_v103_: _dafny.Set
                d_2608_v103_ = _dafny.Set({d_2607_v102_})
                (globalState).f2 = ((d_2607_v102_) not in (d_2608_v103_)) and (p2)
                (globalState).f7 = (p2 if p2 else (p2) == (not(p2)))
                d_2609_v104_: _dafny.Map
                d_2609_v104_ = _dafny.Map({p2: p2})
                d_2593_cf38_ = len(d_2609_v104_)
            r0 = p2
        elif source37_.is_DC21:
            d_2610___mcc_h2_ = source37_.cf40
            d_2611_cf40_ = d_2610___mcc_h2_
            d_2612_v105_: D15
            d_2612_v105_ = D15_DC43(p1, d_2611_cf40_, d_2611_cf40_)
            (globalState).f2 = (d_2612_v105_).cf78
            d_2484_v0_ = d_2484_v0_
            d_2613_v106_: _dafny.Array
            nw398_ = _dafny.Array(int(0), 20)
            d_2613_v106_ = nw398_
            index443_ = default__.safeIndex(989, (d_2613_v106_).length(0))
            (d_2613_v106_)[index443_] = default__.safeDivisionInt(p0, p1)
            index444_ = default__.safeIndex(989, (d_2613_v106_).length(0))
            (d_2613_v106_)[index444_] = default__.safeModuloInt(p1, p1)
            d_2614_v107_: _dafny.Array
            def lambda101_(d_2615_p0_):
                def lambda102_(d_2616_i10_):
                    return D17_DC49(d_2615_p0_)

                return lambda102_

            init62_ = lambda101_(p0)
            nw399_ = _dafny.Array(None, 28)
            for i0_62_ in range(nw399_.length(0)):
                nw399_[i0_62_] = init62_(i0_62_)
            d_2614_v107_ = nw399_
            d_2617_v108_: D17
            d_2617_v108_ = D17_DC49((d_2613_v106_)[default__.safeIndex(989, (d_2613_v106_).length(0))])
            index445_ = default__.safeIndex(353, (d_2614_v107_).length(0))
            (d_2614_v107_)[index445_] = d_2617_v108_
            index446_ = default__.safeIndex(353, (d_2614_v107_).length(0))
            (d_2614_v107_)[index446_] = d_2617_v108_
        elif source37_.is_DC19:
            d_2618___mcc_h3_ = source37_.cf37
            d_2619_cf37_ = d_2618___mcc_h3_
            d_2620_v109_: _dafny.Array
            nw400_ = _dafny.Array(False, 26)
            d_2620_v109_ = nw400_
            index447_ = default__.safeIndex(494, (d_2620_v109_).length(0))
            (d_2620_v109_)[index447_] = False
            index448_ = default__.safeIndex(494, (d_2620_v109_).length(0))
            (d_2620_v109_)[index448_] = (p1) > (p1)
            if (d_2620_v109_)[default__.safeIndex(494, (d_2620_v109_).length(0))]:
                r1 = (d_2620_v109_)[default__.safeIndex(494, (d_2620_v109_).length(0))]
                d_2621_v110_: C0
                nw401_ = C0()
                nw401_.ctor__()
                d_2621_v110_ = nw401_
                d_2622_v111_: C4
                nw402_ = C4()
                nw402_.ctor__(d_2485_v1_, (len(d_2485_v1_) if not(False) else (D16_DC45(d_2620_v109_, d_2484_v0_, p1, p1, (d_2620_v109_)[default__.safeIndex(494, (d_2620_v109_).length(0))])).cf84), True, d_2485_v1_, p1, p0)
                d_2622_v111_ = nw402_
                d_2623_v112_: _dafny.MultiSet
                d_2623_v112_ = _dafny.MultiSet([631, p1])
                d_2624_v113_: _dafny.Set
                d_2624_v113_ = _dafny.Set({(d_2620_v109_)[default__.safeIndex(494, (d_2620_v109_).length(0))], (d_2620_v109_)[default__.safeIndex(494, (d_2620_v109_).length(0))], p2, p2})
                d_2625_v114_: C2
                nw403_ = C2()
                nw403_.ctor__(((d_2623_v112_)[len(default__.fm50(globalState))] if (len(default__.fm50(globalState))) in (d_2623_v112_) else p1), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elg")), -62, len(_dafny.Map({d_2624_v113_: p0})))
                d_2625_v114_ = nw403_
                (globalState).f6 = (-132) * (p0)
            elif True:
                d_2620_v109_ = d_2620_v109_
                index449_ = default__.safeIndex(494, (d_2620_v109_).length(0))
                (d_2620_v109_)[index449_] = p2
                (globalState).f6 = 438
                d_2626_v115_: _dafny.Array
                nw404_ = _dafny.Array(D11.default()(), 28)
                d_2626_v115_ = nw404_
                d_2627_v116_: _dafny.Array
                nw405_ = _dafny.Array(int(0), 11)
                d_2627_v116_ = nw405_
                d_2628_v117_: _dafny.Map
                d_2628_v117_ = _dafny.Map({(d_2485_v1_)[default__.safeIndex((0) - ((0) - (p0)), len(d_2485_v1_))]: default__.safeDivisionInt(p1, 665)})
                rhs418_ = (0) - ((0) - (((d_2628_v117_)[d_2484_v0_] if (d_2484_v0_) in (d_2628_v117_) else (0) - (len(_dafny.Map({(d_2620_v109_)[default__.safeIndex(494, (d_2620_v109_).length(0))]: d_2484_v0_}))))))
                rhs419_ = d_2626_v115_
                rhs420_ = d_2627_v116_
                lhs348_ = globalState
                lhs348_.f6 = rhs418_
                d_2626_v115_ = rhs419_
                d_2627_v116_ = rhs420_
                d_2629_v118_: _dafny.Seq
                d_2629_v118_ = _dafny.SeqWithoutIsStrInference([p0, p1])
                d_2630_v119_: D4
                d_2630_v119_ = D4_DC13(d_2629_v118_)
                d_2631_v120_: _dafny.Seq
                d_2631_v120_ = _dafny.SeqWithoutIsStrInference([d_2629_v118_, (d_2630_v119_).cf29, d_2629_v118_])
                index450_ = default__.safeIndex(814, (d_2627_v116_).length(0))
                (d_2627_v116_)[index450_] = len((d_2631_v120_)[default__.safeIndex(p0, len(d_2631_v120_))])
                index451_ = default__.safeIndex(814, (d_2627_v116_).length(0))
                rhs421_ = default__.fm0((_dafny.Set({p2})).isdisjoint(_dafny.Set({(self).fm13(p0, d_2484_v0_, _dafny.CodePoint('c'), globalState), not(p2), (d_2620_v109_)[default__.safeIndex(494, (d_2620_v109_).length(0))], (d_2620_v109_)[default__.safeIndex(494, (d_2620_v109_).length(0))], p2})), globalState)
                rhs422_ = p0
                rhs423_ = p1
                lhs349_ = globalState
                lhs350_ = d_2627_v116_
                lhs351_ = default__.safeIndex(814, (d_2627_v116_).length(0))
                lhs352_ = globalState
                lhs349_.f6 = rhs421_
                lhs350_[lhs351_] = rhs422_
                lhs352_.f6 = rhs423_
            (globalState).f6 = p0
            d_2632_v121_: _dafny.Map
            d_2632_v121_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhcv")): p0})
            d_2633_v122_: D3
            d_2633_v122_ = D3_DC10(p2, p1, d_2632_v121_, d_2485_v1_, -146)
            d_2634_v123_: _dafny.Seq
            d_2634_v123_ = _dafny.SeqWithoutIsStrInference([(d_2633_v122_).cf22])
            d_2635_v124_: _dafny.Map
            d_2635_v124_ = _dafny.Map({p0: p2})
            d_2636_v125_: _dafny.Seq
            d_2636_v125_ = _dafny.SeqWithoutIsStrInference([p0, len(d_2635_v124_), p1])
            d_2637_v126_: C1
            nw406_ = C1()
            nw406_.ctor__(d_2620_v109_, len((d_2634_v123_) + (d_2634_v123_)), len(_dafny.SeqWithoutIsStrInference([p1])), (d_2636_v125_)[default__.safeIndex(p0, len(d_2636_v125_))], p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbxl")))
            d_2637_v126_ = nw406_
        elif True:
            d_2638___mcc_h4_ = source37_.cf41
            d_2639_cf41_ = d_2638___mcc_h4_
            d_2640_v127_: C2
            nw407_ = C2()
            nw407_.ctor__(p1, d_2485_v1_, p1, p1)
            d_2640_v127_ = nw407_
            (globalState).f6 = default__.safeModuloInt(default__.safeDivisionInt(d_2640_v127_.f27, d_2640_v127_.f27), d_2640_v127_.f27)
            d_2641_v128_: _dafny.Seq
            d_2641_v128_ = _dafny.SeqWithoutIsStrInference([p2])
            d_2642_v129_: _dafny.Set
            d_2642_v129_ = _dafny.Set({(d_2640_v127_).fm41(globalState), (d_2640_v127_).fm1(p0, False, (d_2641_v128_)[default__.safeIndex(d_2640_v127_.f27, len(d_2641_v128_))], globalState), p2})
            d_2642_v129_ = ((_dafny.Set({p2})) - (d_2642_v129_)) | (d_2642_v129_)
            d_2643_v130_: _dafny.Map
            d_2643_v130_ = _dafny.Map({default__.fm18((d_2641_v128_)[default__.safeIndex(d_2640_v127_.f27, len(d_2641_v128_))], d_2640_v127_.f27, globalState): (True) or (p2)})
            d_2643_v130_ = (d_2643_v130_).set(default__.fm44(p1, p2, globalState), p2)
        (globalState).f6 = p1
        r1 = not(not(not(p2)))
        d_2644_v131_: _dafny.Set
        d_2644_v131_ = _dafny.Set({p0})
        d_2645_v132_: _dafny.Map
        d_2645_v132_ = _dafny.Map({False: (0) - (len(d_2644_v131_))})
        r0 = (p2) == ((len(_dafny.SeqWithoutIsStrInference([((d_2645_v132_)[p2] if (p2) in (d_2645_v132_) else p1)]))) < (p0))
        d_2646_v133_: _dafny.Set
        d_2646_v133_ = _dafny.Set({p2})
        r1 = not(((d_2646_v133_).isdisjoint(_dafny.Set({p2}))) == (True))
        return r0, r1

    def m11(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        d_2647_v0_: bool
        d_2647_v0_ = False
        d_2648_v1_: int
        d_2648_v1_ = 411
        d_2649_v2_: _dafny.Seq
        d_2649_v2_ = _dafny.SeqWithoutIsStrInference([d_2648_v1_])
        d_2650_v3_: _dafny.Map
        d_2650_v3_ = _dafny.Map({len(d_2649_v2_): d_2648_v1_})
        d_2651_v4_: _dafny.Seq
        d_2651_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntg"))
        d_2652_v5_: _dafny.Map
        d_2652_v5_ = _dafny.Map({d_2651_v4_: d_2648_v1_})
        d_2653_v6_: D3
        d_2653_v6_ = D3_DC10(d_2647_v0_, len(d_2650_v3_), d_2652_v5_, d_2651_v4_, d_2648_v1_)
        d_2654_v7_: _dafny.Seq
        d_2654_v7_ = _dafny.SeqWithoutIsStrInference([d_2651_v4_, d_2651_v4_, d_2651_v4_])
        rhs424_ = (d_2653_v6_ if d_2647_v0_ else d_2653_v6_)
        rhs425_ = (default__.safeDivisionInt(d_2648_v1_, d_2648_v1_)) >= (default__.safeDivisionInt(len(d_2651_v4_), d_2648_v1_))
        rhs426_ = d_2654_v7_
        lhs353_ = globalState
        d_2653_v6_ = rhs424_
        lhs353_.f7 = rhs425_
        d_2654_v7_ = rhs426_
        d_2655_v8_: _dafny.Map
        d_2655_v8_ = _dafny.Map({d_2647_v0_: len(default__.fm50(globalState))})
        d_2656_v10_: _dafny.Set
        def iife150_():
            coll60_ = _dafny.Set()
            compr_60_: int
            for compr_60_ in (_dafny.SeqWithoutIsStrInference([d_2648_v1_, ((d_2655_v8_)[d_2647_v0_] if (d_2647_v0_) in (d_2655_v8_) else d_2648_v1_)])).Elements:
                d_2657_v9_: int = compr_60_
                if (d_2657_v9_) in (_dafny.SeqWithoutIsStrInference([d_2648_v1_, ((d_2655_v8_)[d_2647_v0_] if (d_2647_v0_) in (d_2655_v8_) else d_2648_v1_)])):
                    coll60_ = coll60_.union(_dafny.Set([(d_2657_v9_) - (len(_dafny.Map({247: _dafny.CodePoint('s')})))]))
            return _dafny.Set(coll60_)
        d_2656_v10_ = _dafny.Set({iife150_()
        })
        d_2658_v11_: _dafny.Map
        d_2658_v11_ = _dafny.Map({d_2648_v1_: d_2647_v0_})
        d_2659_v12_: _dafny.Array
        nw408_ = _dafny.Array(None, 8)
        nw408_[int(0)] = d_2648_v1_
        nw408_[int(1)] = default__.fm0(d_2647_v0_, globalState)
        nw408_[int(2)] = len((d_2649_v2_) + (d_2649_v2_))
        nw408_[int(3)] = d_2648_v1_
        nw408_[int(4)] = d_2648_v1_
        nw408_[int(5)] = len((d_2656_v10_ if d_2647_v0_ else d_2656_v10_))
        nw408_[int(6)] = len(d_2658_v11_)
        nw408_[int(7)] = (d_2648_v1_) + (d_2648_v1_)
        d_2659_v12_ = nw408_
        index452_ = default__.safeIndex(233, (d_2659_v12_).length(0))
        (d_2659_v12_)[index452_] = d_2648_v1_
        d_2660_v13_: _dafny.MultiSet
        d_2660_v13_ = _dafny.MultiSet([d_2647_v0_, d_2647_v0_, d_2647_v0_])
        d_2661_v14_: D8
        d_2661_v14_ = D8_DC24(d_2660_v13_)
        pat_let_tv42_ = d_2650_v3_
        pat_let_tv43_ = d_2647_v0_
        index453_ = default__.safeIndex(233, (d_2659_v12_).length(0))
        def lambda103_(source38_):
            if source38_.is_DC25:
                d_2662___mcc_h0_ = source38_.cf44
                d_2663_cf44_ = d_2662___mcc_h0_
                return (_dafny.Set({d_2663_cf44_, d_2663_cf44_})).ispropersubset(_dafny.Set({len(pat_let_tv42_)}))
            elif True:
                d_2664___mcc_h1_ = source38_.cf43
                d_2665_cf43_ = d_2664___mcc_h1_
                return not(pat_let_tv43_)

        rhs427_ = ((d_2651_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iokci")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "if")))
        rhs428_ = lambda103_(d_2661_v14_)
        rhs429_ = d_2651_v4_
        rhs430_ = d_2648_v1_
        rhs431_ = d_2647_v0_
        lhs354_ = globalState
        lhs355_ = d_2659_v12_
        lhs356_ = default__.safeIndex(233, (d_2659_v12_).length(0))
        d_2651_v4_ = rhs427_
        lhs354_.f7 = rhs428_
        d_2651_v4_ = rhs429_
        lhs355_[lhs356_] = rhs430_
        r0 = rhs431_
        r0 = d_2647_v0_
        index454_ = default__.safeIndex(233, (d_2659_v12_).length(0))
        (d_2659_v12_)[index454_] = default__.fm0(d_2647_v0_, globalState)
        (globalState).f7 = (d_2660_v13_).ispropersubset(d_2660_v13_)
        if d_2647_v0_:
            d_2666_v15_: D17
            d_2666_v15_ = D17_DC50(_dafny.MultiSet([d_2647_v0_]))
            d_2667_v16_: _dafny.Map
            d_2667_v16_ = _dafny.Map({820: d_2666_v15_})
            d_2667_v16_ = (d_2667_v16_) | ((d_2667_v16_) | (d_2667_v16_))
            if ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]) != ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]):
                d_2668_v17_: bool
                d_2669_v18_: _dafny.Seq
                out79_: bool
                out80_: _dafny.Seq
                out79_, out80_ = default__.m0((702) <= ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]), d_2647_v0_, d_2648_v1_, d_2648_v1_, globalState)
                d_2668_v17_ = out79_
                d_2669_v18_ = out80_
                d_2670_v19_: D17
                d_2670_v19_ = D17_DC47(d_2650_v3_)
                d_2670_v19_ = d_2670_v19_
                d_2655_v8_ = d_2655_v8_
                d_2671_v20_: _dafny.Map
                d_2671_v20_ = _dafny.Map({d_2647_v0_: d_2647_v0_})
                d_2672_v21_: _dafny.MultiSet
                d_2672_v21_ = _dafny.MultiSet([len(d_2671_v20_), 850, default__.fm0(d_2668_v17_, globalState), (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], d_2648_v1_])
                index455_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                (d_2659_v12_)[index455_] = (0) - (default__.safeModuloInt(((d_2672_v21_)[(d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]] if ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]) in (d_2672_v21_) else d_2648_v1_), (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]))
                d_2650_v3_ = (d_2650_v3_).set(d_2648_v1_, default__.safeModuloInt(len(d_2669_v18_), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))))))
            elif True:
                d_2673_v22_: _dafny.Seq
                d_2673_v22_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2647_v0_: d_2647_v0_})])
                d_2674_v23_: D19
                d_2674_v23_ = D19_DC56(d_2659_v12_, not(d_2647_v0_), (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], d_2673_v22_)
                d_2648_v1_ = (d_2674_v23_).cf107
                (globalState).f6 = d_2648_v1_
                d_2675_v24_: str
                d_2675_v24_ = _dafny.CodePoint('g')
                d_2676_v25_: _dafny.MultiSet
                d_2676_v25_ = _dafny.MultiSet([d_2675_v24_, default__.fm34(d_2675_v24_, globalState)])
                d_2677_v26_: D0
                d_2677_v26_ = D0_DC0(d_2676_v25_)
                d_2678_v27_: _dafny.Array
                nw409_ = _dafny.Array(None, 17)
                nw409_[int(0)] = d_2651_v4_
                nw409_[int(1)] = d_2651_v4_
                nw409_[int(2)] = d_2651_v4_
                nw409_[int(3)] = (d_2651_v4_) + (d_2651_v4_)
                nw409_[int(4)] = d_2651_v4_
                nw409_[int(5)] = default__.fm18(d_2647_v0_, (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], globalState)
                nw409_[int(6)] = d_2651_v4_
                nw409_[int(7)] = d_2651_v4_
                nw409_[int(8)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbogkbui"))).set(default__.safeIndex((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbogkbui")))), d_2675_v24_)) + (_dafny.SeqWithoutIsStrInference([d_2675_v24_ for d_2679_i0_ in range(default__.abs(355))]))
                nw409_[int(9)] = default__.fm31(d_2675_v24_, d_2677_v26_, globalState)
                nw409_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uan"))
                nw409_[int(11)] = _dafny.SeqWithoutIsStrInference([d_2675_v24_ for d_2680_i1_ in range(default__.abs(783))])
                nw409_[int(12)] = _dafny.SeqWithoutIsStrInference([d_2675_v24_ for d_2681_i2_ in range(default__.abs(689))])
                nw409_[int(13)] = d_2651_v4_
                nw409_[int(14)] = d_2651_v4_
                nw409_[int(15)] = d_2651_v4_
                nw409_[int(16)] = d_2651_v4_
                d_2678_v27_ = nw409_
                index456_ = default__.safeIndex(847, (d_2678_v27_).length(0))
                (d_2678_v27_)[index456_] = d_2651_v4_
                index457_ = default__.safeIndex(847, (d_2678_v27_).length(0))
                (d_2678_v27_)[index457_] = (d_2651_v4_) + (d_2651_v4_)
                d_2682_v28_: _dafny.Array
                def lambda104_(d_2683_v0_):
                    def lambda105_(d_2684_i3_):
                        return not(d_2683_v0_)

                    return lambda105_

                init63_ = lambda104_(d_2647_v0_)
                nw410_ = _dafny.Array(None, 21)
                for i0_63_ in range(nw410_.length(0)):
                    nw410_[i0_63_] = init63_(i0_63_)
                d_2682_v28_ = nw410_
                d_2685_v29_: _dafny.Map
                d_2685_v29_ = _dafny.Map({d_2682_v28_: _dafny.Map({d_2647_v0_: d_2648_v1_})})
                d_2685_v29_ = (d_2685_v29_).set(d_2682_v28_, d_2655_v8_)
                d_2686_v30_: _dafny.Map
                d_2686_v30_ = _dafny.Map({D4_DC13(d_2649_v2_): d_2647_v0_})
                d_2687_v31_: D4
                d_2687_v31_ = D4_DC13(d_2649_v2_)
                d_2686_v30_ = (d_2686_v30_).set(d_2687_v31_, not(d_2647_v0_))
            d_2655_v8_ = (d_2655_v8_).set(d_2647_v0_, (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))])
            d_2688_v32_: D0
            d_2688_v32_ = D0_DC1(d_2647_v0_, -521)
            d_2689_v33_: _dafny.Array
            nw411_ = _dafny.Array(False, 16)
            d_2689_v33_ = nw411_
            index458_ = default__.safeIndex(233, (d_2659_v12_).length(0))
            rhs432_ = d_2688_v32_
            rhs433_ = d_2689_v33_
            rhs434_ = (((0) - ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]) if d_2647_v0_ else len(d_2649_v2_))) + ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))])
            rhs435_ = (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]
            lhs357_ = globalState
            lhs358_ = d_2659_v12_
            lhs359_ = default__.safeIndex(233, (d_2659_v12_).length(0))
            d_2688_v32_ = rhs432_
            d_2689_v33_ = rhs433_
            lhs357_.f6 = rhs434_
            lhs358_[lhs359_] = rhs435_
            d_2690_v34_: _dafny.Array
            nw412_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_2690_v34_ = nw412_
            d_2691_v35_: _dafny.Array
            nw413_ = _dafny.Array(None, 14)
            nw413_[int(0)] = d_2649_v2_
            nw413_[int(1)] = d_2649_v2_
            nw413_[int(2)] = d_2649_v2_
            nw413_[int(3)] = d_2649_v2_
            nw413_[int(4)] = d_2649_v2_
            nw413_[int(5)] = d_2649_v2_
            nw413_[int(6)] = d_2649_v2_
            nw413_[int(7)] = d_2649_v2_
            nw413_[int(8)] = d_2649_v2_
            nw413_[int(9)] = d_2649_v2_
            nw413_[int(10)] = d_2649_v2_
            nw413_[int(11)] = d_2649_v2_
            nw413_[int(12)] = _dafny.SeqWithoutIsStrInference([d_2648_v1_, d_2648_v1_])
            nw413_[int(13)] = d_2649_v2_
            d_2691_v35_ = nw413_
            index459_ = default__.safeIndex(745, (d_2690_v34_).length(0))
            (d_2690_v34_)[index459_] = d_2691_v35_
            index460_ = default__.safeIndex(745, (d_2690_v34_).length(0))
            (d_2690_v34_)[index460_] = d_2691_v35_
        elif True:
            d_2692_v36_: _dafny.Array
            def lambda106_(d_2693_v3_):
                def lambda107_(d_2694_i4_):
                    return d_2693_v3_

                return lambda107_

            init64_ = lambda106_(d_2650_v3_)
            nw414_ = _dafny.Array(None, 19)
            for i0_64_ in range(nw414_.length(0)):
                nw414_[i0_64_] = init64_(i0_64_)
            d_2692_v36_ = nw414_
            d_2695_v38_: _dafny.Set
            d_2695_v38_ = _dafny.Set({d_2648_v1_, len(_dafny.SeqWithoutIsStrInference([d_2648_v1_ for d_2696_i5_ in range(default__.abs(229))]))})
            index461_ = default__.safeIndex(552, (d_2692_v36_).length(0))
            def iife151_():
                coll61_ = _dafny.Map()
                compr_61_: int
                for compr_61_ in (d_2695_v38_).Elements:
                    d_2697_v37_: int = compr_61_
                    if (d_2697_v37_) in (d_2695_v38_):
                        coll61_[(d_2697_v37_) + (d_2648_v1_)] = d_2648_v1_
                return _dafny.Map(coll61_)
            (d_2692_v36_)[index461_] = iife151_()
            
            d_2698_v39_: str
            d_2698_v39_ = _dafny.CodePoint('q')
            index462_ = default__.safeIndex(552, (d_2692_v36_).length(0))
            rhs436_ = not((self).fm13((-391) * ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]), d_2698_v39_, d_2698_v39_, globalState))
            rhs437_ = d_2650_v3_
            rhs438_ = d_2647_v0_
            rhs439_ = (d_2647_v0_) or (d_2647_v0_)
            lhs360_ = globalState
            lhs361_ = d_2692_v36_
            lhs362_ = default__.safeIndex(552, (d_2692_v36_).length(0))
            lhs363_ = globalState
            lhs360_.f7 = rhs436_
            lhs361_[lhs362_] = rhs437_
            lhs363_.f2 = rhs438_
            d_2647_v0_ = rhs439_
            d_2699_v40_: D4
            d_2699_v40_ = D4_DC13(d_2649_v2_)
            d_2700_v41_: _dafny.Seq
            d_2700_v41_ = _dafny.SeqWithoutIsStrInference([(self).fm13(d_2648_v1_, d_2698_v39_, d_2698_v39_, globalState), d_2647_v0_])
            source39_ = (d_2699_v40_ if not((d_2647_v0_) == (d_2647_v0_)) else default__.fm65(d_2647_v0_, (d_2655_v8_).set((d_2700_v41_)[default__.safeIndex((d_2660_v13_).cardinality, len(d_2700_v41_))], 441), globalState))
            if source39_.is_DC14:
                d_2701___mcc_h2_ = source39_.cf30
                d_2702___mcc_h3_ = source39_.cf31
                d_2703_cf31_ = d_2702___mcc_h3_
                d_2704_cf30_ = d_2701___mcc_h2_
                d_2658_v11_ = (d_2658_v11_).set(len((d_2655_v8_).set(not(d_2647_v0_), (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))])), (d_2695_v38_).issubset(d_2695_v38_))
                d_2705_v42_: D2
                d_2705_v42_ = D2_DC8(d_2647_v0_, d_2647_v0_, d_2651_v4_)
                d_2706_v43_: _dafny.MultiSet
                d_2706_v43_ = _dafny.MultiSet([d_2704_cf30_])
                d_2707_v44_: D0
                d_2707_v44_ = D0_DC0(d_2706_v43_)
                d_2708_v45_: _dafny.Set
                d_2708_v45_ = _dafny.Set({d_2707_v44_, d_2707_v44_})
                d_2709_v46_: D3
                d_2709_v46_ = D3_DC11((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], d_2705_v42_, d_2708_v45_, d_2647_v0_)
                d_2710_v47_: D3
                d_2710_v47_ = D3_DC11(default__.fm0(d_2647_v0_, globalState), D2_DC8(d_2647_v0_, d_2647_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqtltn"))), (d_2709_v46_).cf26, not(d_2647_v0_))
                (globalState).f6 = (((d_2709_v46_).cf24) + (d_2703_cf31_) if d_2647_v0_ else (d_2649_v2_)[default__.safeIndex(d_2648_v1_, len(d_2649_v2_))])
                d_2711_v48_: _dafny.Map
                d_2711_v48_ = _dafny.Map({d_2648_v1_: d_2658_v11_})
                def iife152_():
                    coll62_ = _dafny.Set()
                    compr_62_: int
                    for compr_62_ in ((((d_2711_v48_)[(d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]] if ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]) in (d_2711_v48_) else d_2658_v11_) if d_2647_v0_ else d_2658_v11_)).keys.Elements:
                        d_2712_v49_: int = compr_62_
                        if (d_2712_v49_) in ((((d_2711_v48_)[(d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]] if ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]) in (d_2711_v48_) else d_2658_v11_) if d_2647_v0_ else d_2658_v11_)):
                            coll62_ = coll62_.union(_dafny.Set([default__.safeModuloInt(d_2712_v49_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsn"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2713_i6_ in range(default__.abs(935))]))))]))
                    return _dafny.Set(coll62_)
                d_2695_v38_ = iife152_()
                
                (globalState).f2 = (default__.fm0(True, globalState)) == ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))])
            elif source39_.is_DC15:
                d_2714___mcc_h4_ = source39_.cf32
                d_2715_cf32_ = d_2714___mcc_h4_
                index463_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                (d_2659_v12_)[index463_] = (d_2649_v2_)[default__.safeIndex(d_2715_cf32_, len(d_2649_v2_))]
                d_2716_v50_: _dafny.Array
                nw415_ = _dafny.Array(False, 19)
                d_2716_v50_ = nw415_
                index464_ = default__.safeIndex(686, (d_2716_v50_).length(0))
                (d_2716_v50_)[index464_] = d_2647_v0_
                index465_ = default__.safeIndex(686, (d_2716_v50_).length(0))
                rhs440_ = d_2647_v0_
                rhs441_ = d_2647_v0_
                lhs364_ = d_2716_v50_
                lhs365_ = default__.safeIndex(686, (d_2716_v50_).length(0))
                lhs366_ = globalState
                lhs364_[lhs365_] = rhs440_
                lhs366_.f7 = rhs441_
                d_2717_v51_: _dafny.Array
                nw416_ = _dafny.Array(D12.default()(), 15)
                d_2717_v51_ = nw416_
                d_2718_v52_: D12
                d_2718_v52_ = D12_DC36(len(d_2651_v4_), d_2648_v1_, (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))])
                index466_ = default__.safeIndex(408, (d_2717_v51_).length(0))
                (d_2717_v51_)[index466_] = d_2718_v52_
                index467_ = default__.safeIndex(408, (d_2717_v51_).length(0))
                (d_2717_v51_)[index467_] = (d_2718_v52_ if ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]) >= ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]) else d_2718_v52_)
                (globalState).f2 = (487) != ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))])
            elif source39_.is_DC13:
                d_2719___mcc_h5_ = source39_.cf29
                d_2720_cf29_ = d_2719___mcc_h5_
                d_2721_v53_: _dafny.Seq
                d_2721_v53_ = _dafny.SeqWithoutIsStrInference([d_2649_v2_, d_2720_cf29_, d_2720_cf29_, _dafny.SeqWithoutIsStrInference([d_2648_v1_ for d_2722_i7_ in range(default__.abs(220))])])
                d_2723_v54_: _dafny.Map
                d_2723_v54_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2647_v0_])): d_2721_v53_})
                d_2724_v55_: D15
                d_2724_v55_ = D15_DC43(d_2648_v1_, d_2647_v0_, d_2647_v0_)
                d_2725_v56_: _dafny.Array
                nw417_ = _dafny.Array(None, 18)
                nw417_[int(0)] = d_2721_v53_
                nw417_[int(1)] = d_2721_v53_
                nw417_[int(2)] = d_2721_v53_
                nw417_[int(3)] = (d_2721_v53_).set(default__.safeIndex(((d_2655_v8_)[d_2647_v0_] if (d_2647_v0_) in (d_2655_v8_) else (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]), len(d_2721_v53_)), d_2720_cf29_)
                nw417_[int(4)] = d_2721_v53_
                nw417_[int(5)] = ((d_2723_v54_)[215] if (215) in (d_2723_v54_) else d_2721_v53_)
                nw417_[int(6)] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], d_2648_v1_]), default__.fm20(d_2647_v0_, d_2647_v0_, False, globalState)])
                nw417_[int(7)] = d_2721_v53_
                nw417_[int(8)] = (d_2721_v53_ if not(d_2647_v0_) else d_2721_v53_)
                nw417_[int(9)] = (d_2721_v53_) + (d_2721_v53_)
                nw417_[int(10)] = default__.fm66(d_2647_v0_, d_2647_v0_, D15_DC43((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], d_2647_v0_, d_2647_v0_), d_2700_v41_, globalState)
                nw417_[int(11)] = d_2721_v53_
                nw417_[int(12)] = (d_2721_v53_) + (d_2721_v53_)
                nw417_[int(13)] = d_2721_v53_
                nw417_[int(14)] = ((d_2721_v53_ if True else d_2721_v53_)).set(default__.safeIndex(len(d_2651_v4_), len((d_2721_v53_ if True else d_2721_v53_))), d_2649_v2_)
                nw417_[int(15)] = (d_2721_v53_) + (d_2721_v53_)
                nw417_[int(16)] = d_2721_v53_
                nw417_[int(17)] = (default__.fm66(d_2647_v0_, d_2647_v0_, d_2724_v55_, d_2700_v41_, globalState)).set(default__.safeIndex((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], len(default__.fm66(d_2647_v0_, d_2647_v0_, d_2724_v55_, d_2700_v41_, globalState))), d_2649_v2_)
                d_2725_v56_ = nw417_
                index468_ = default__.safeIndex(880, (d_2725_v56_).length(0))
                (d_2725_v56_)[index468_] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_2700_v41_), len(d_2700_v41_)]) for d_2726_i8_ in range(default__.abs(915))])
                index469_ = default__.safeIndex(405, (d_2725_v56_).length(0))
                (d_2725_v56_)[index469_] = (d_2721_v53_) + (d_2721_v53_)
                d_2727_v57_: _dafny.Array
                nw418_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_2727_v57_ = nw418_
                d_2728_v58_: _dafny.Array
                nw419_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
                d_2728_v58_ = nw419_
                index470_ = default__.safeIndex(409, (d_2727_v57_).length(0))
                (d_2727_v57_)[index470_] = d_2728_v58_
                d_2729_v59_: _dafny.Array
                nw420_ = _dafny.Array(None, 7)
                nw420_[int(0)] = d_2698_v39_
                nw420_[int(1)] = d_2698_v39_
                nw420_[int(2)] = d_2698_v39_
                nw420_[int(3)] = _dafny.CodePoint('q')
                nw420_[int(4)] = _dafny.CodePoint('q')
                nw420_[int(5)] = d_2698_v39_
                nw420_[int(6)] = d_2698_v39_
                d_2729_v59_ = nw420_
                d_2730_v60_: _dafny.Array
                nw421_ = _dafny.Array(None, 24)
                nw421_[int(0)] = d_2729_v59_
                nw421_[int(1)] = d_2729_v59_
                nw421_[int(2)] = d_2729_v59_
                nw421_[int(3)] = d_2729_v59_
                nw421_[int(4)] = d_2729_v59_
                nw421_[int(5)] = d_2729_v59_
                nw421_[int(6)] = d_2729_v59_
                nw421_[int(7)] = d_2729_v59_
                nw421_[int(8)] = d_2729_v59_
                nw421_[int(9)] = d_2729_v59_
                nw421_[int(10)] = d_2729_v59_
                nw421_[int(11)] = d_2729_v59_
                nw421_[int(12)] = d_2729_v59_
                nw421_[int(13)] = d_2729_v59_
                nw421_[int(14)] = d_2729_v59_
                nw421_[int(15)] = d_2729_v59_
                nw421_[int(16)] = d_2729_v59_
                nw421_[int(17)] = d_2729_v59_
                nw421_[int(18)] = d_2729_v59_
                nw421_[int(19)] = d_2729_v59_
                nw421_[int(20)] = d_2729_v59_
                nw421_[int(21)] = d_2729_v59_
                nw421_[int(22)] = d_2729_v59_
                nw421_[int(23)] = d_2729_v59_
                d_2730_v60_ = nw421_
                d_2731_v61_: _dafny.Map
                d_2731_v61_ = _dafny.Map({d_2730_v60_: d_2728_v58_})
                index471_ = default__.safeIndex(880, (d_2725_v56_).length(0))
                index472_ = default__.safeIndex(405, (d_2725_v56_).length(0))
                index473_ = default__.safeIndex(409, (d_2727_v57_).length(0))
                index474_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                rhs442_ = ((d_2660_v13_)[(d_2647_v0_) in (d_2700_v41_)] if ((d_2647_v0_) in (d_2700_v41_)) in (d_2660_v13_) else (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))])
                rhs443_ = (d_2721_v53_) + (d_2721_v53_)
                rhs444_ = default__.fm66(d_2647_v0_, not(d_2647_v0_), D15_DC43(len(d_2720_cf29_), d_2647_v0_, d_2647_v0_), (d_2700_v41_) + (d_2700_v41_), globalState)
                rhs445_ = ((d_2731_v61_)[d_2730_v60_] if (d_2730_v60_) in (d_2731_v61_) else d_2728_v58_)
                rhs446_ = default__.fm0(d_2647_v0_, globalState)
                lhs367_ = d_2725_v56_
                lhs368_ = default__.safeIndex(880, (d_2725_v56_).length(0))
                lhs369_ = d_2725_v56_
                lhs370_ = default__.safeIndex(405, (d_2725_v56_).length(0))
                lhs371_ = d_2727_v57_
                lhs372_ = default__.safeIndex(409, (d_2727_v57_).length(0))
                lhs373_ = d_2659_v12_
                lhs374_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                d_2648_v1_ = rhs442_
                lhs367_[lhs368_] = rhs443_
                lhs369_[lhs370_] = rhs444_
                lhs371_[lhs372_] = rhs445_
                lhs373_[lhs374_] = rhs446_
                d_2732_v62_: _dafny.Array
                nw422_ = _dafny.Array(False, 10)
                d_2732_v62_ = nw422_
                d_2733_v63_: D20
                d_2733_v63_ = D20_DC58(d_2647_v0_)
                index475_ = default__.safeIndex(702, (d_2732_v62_).length(0))
                (d_2732_v62_)[index475_] = (d_2733_v63_).cf110
                index476_ = default__.safeIndex(128, (d_2732_v62_).length(0))
                (d_2732_v62_)[index476_] = d_2647_v0_
                d_2734_v64_: _dafny.MultiSet
                d_2734_v64_ = _dafny.MultiSet([d_2648_v1_, d_2648_v1_, (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]])
                d_2735_v65_: _dafny.Seq
                d_2735_v65_ = _dafny.SeqWithoutIsStrInference([d_2700_v41_, _dafny.SeqWithoutIsStrInference([d_2647_v0_, d_2647_v0_])])
                index477_ = default__.safeIndex(702, (d_2732_v62_).length(0))
                index478_ = default__.safeIndex(128, (d_2732_v62_).length(0))
                index479_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                index480_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                rhs447_ = d_2647_v0_
                rhs448_ = d_2647_v0_
                rhs449_ = d_2648_v1_
                rhs450_ = ((d_2734_v64_)[453] if (453) in (d_2734_v64_) else (((d_2692_v36_)[default__.safeIndex(552, (d_2692_v36_).length(0))])[(d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]] if ((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]) in ((d_2692_v36_)[default__.safeIndex(552, (d_2692_v36_).length(0))]) else 273))
                rhs451_ = ((d_2735_v65_)[default__.safeIndex((d_2660_v13_).cardinality, len(d_2735_v65_))]) + (d_2700_v41_)
                lhs375_ = d_2732_v62_
                lhs376_ = default__.safeIndex(702, (d_2732_v62_).length(0))
                lhs377_ = d_2732_v62_
                lhs378_ = default__.safeIndex(128, (d_2732_v62_).length(0))
                lhs379_ = d_2659_v12_
                lhs380_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                lhs381_ = d_2659_v12_
                lhs382_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                lhs375_[lhs376_] = rhs447_
                lhs377_[lhs378_] = rhs448_
                lhs379_[lhs380_] = rhs449_
                lhs381_[lhs382_] = rhs450_
                d_2700_v41_ = rhs451_
                d_2736_v67_: _dafny.Array
                def lambda108_(d_2737_v1_, d_2738_v12_, d_2739_v0_, d_2740_v11_, d_2741_v4_):
                    def lambda109_(d_2742_i9_):
                        def iife153_():
                            coll63_ = _dafny.Map()
                            compr_63_: int
                            for compr_63_ in _dafny.IntegerRange(-763, 590):
                                d_2743_v66_: int = compr_63_
                                if ((-763) <= (d_2743_v66_)) and ((d_2743_v66_) < (590)):
                                    coll63_[default__.safeModuloInt(d_2743_v66_, (d_2738_v12_)[default__.safeIndex(233, (d_2738_v12_).length(0))])] = d_2739_v0_
                            return _dafny.Map(coll63_)
                        return (_dafny.Map({iife153_()
                        : d_2737_v1_})) | (_dafny.Map({d_2740_v11_: len(d_2741_v4_)}))

                    return lambda109_

                init65_ = lambda108_(d_2648_v1_, d_2659_v12_, d_2647_v0_, d_2658_v11_, d_2651_v4_)
                nw423_ = _dafny.Array(None, 1)
                for i0_65_ in range(nw423_.length(0)):
                    nw423_[i0_65_] = init65_(i0_65_)
                d_2736_v67_ = nw423_
                d_2744_v69_: _dafny.Map
                d_2744_v69_ = _dafny.Map({d_2658_v11_: d_2648_v1_})
                index481_ = default__.safeIndex(21, (d_2736_v67_).length(0))
                def iife154_():
                    coll64_ = _dafny.Map()
                    compr_64_: _dafny.Map
                    for compr_64_ in (default__.fm67(d_2698_v39_, globalState)).keys.Elements:
                        d_2745_v68_: _dafny.Map = compr_64_
                        if (d_2745_v68_) in (default__.fm67(d_2698_v39_, globalState)):
                            coll64_[d_2745_v68_] = 93
                    return _dafny.Map(coll64_)
                (d_2736_v67_)[index481_] = (iife154_()
                 if d_2647_v0_ else d_2744_v69_)
                index482_ = default__.safeIndex(21, (d_2736_v67_).length(0))
                (d_2736_v67_)[index482_] = d_2744_v69_
                d_2746_v70_: _dafny.Map
                d_2746_v70_ = _dafny.Map({(d_2732_v62_)[default__.safeIndex(702, (d_2732_v62_).length(0))]: d_2647_v0_})
                d_2747_v71_: _dafny.Map
                d_2747_v71_ = _dafny.Map({(self).fm13(len(d_2700_v41_), d_2698_v39_, d_2698_v39_, globalState): d_2651_v4_})
                d_2748_v72_: C3
                nw424_ = C3()
                nw424_.ctor__(((d_2746_v70_)[(d_2732_v62_)[default__.safeIndex(702, (d_2732_v62_).length(0))]] if ((d_2732_v62_)[default__.safeIndex(702, (d_2732_v62_).length(0))]) in (d_2746_v70_) else (d_2732_v62_)[default__.safeIndex(702, (d_2732_v62_).length(0))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glah")), ((d_2747_v71_)[(d_2732_v62_)[default__.safeIndex(702, (d_2732_v62_).length(0))]] if ((d_2732_v62_)[default__.safeIndex(702, (d_2732_v62_).length(0))]) in (d_2747_v71_) else d_2651_v4_), d_2648_v1_, (d_2732_v62_)[default__.safeIndex(702, (d_2732_v62_).length(0))], d_2651_v4_, len(d_2700_v41_), d_2648_v1_)
                d_2748_v72_ = nw424_
            elif True:
                d_2749___mcc_h6_ = source39_.cf33
                d_2750_cf33_ = d_2749___mcc_h6_
                d_2751_v73_: _dafny.Array
                def lambda110_(d_2752_v41_, d_2753_v1_):
                    def lambda111_(d_2754_i10_):
                        return D13_DC37(_dafny.MultiSet([len(d_2752_v41_), d_2753_v1_]))

                    return lambda111_

                init66_ = lambda110_(d_2700_v41_, d_2648_v1_)
                nw425_ = _dafny.Array(None, 27)
                for i0_66_ in range(nw425_.length(0)):
                    nw425_[i0_66_] = init66_(i0_66_)
                d_2751_v73_ = nw425_
                d_2755_v74_: _dafny.MultiSet
                d_2755_v74_ = _dafny.MultiSet([(d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]])
                d_2756_v75_: D13
                d_2756_v75_ = D13_DC37((d_2755_v74_).set(len(d_2651_v4_), default__.abs(((d_2755_v74_)[(0) - (d_2648_v1_)] if ((0) - (d_2648_v1_)) in (d_2755_v74_) else d_2648_v1_))))
                index483_ = default__.safeIndex(888, (d_2751_v73_).length(0))
                (d_2751_v73_)[index483_] = d_2756_v75_
                d_2757_v76_: C0
                nw426_ = C0()
                nw426_.ctor__()
                d_2757_v76_ = nw426_
                d_2758_v77_: _dafny.Set
                d_2758_v77_ = _dafny.Set({False})
                d_2759_v78_: _dafny.Seq
                d_2759_v78_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2647_v0_, d_2647_v0_])])
                d_2760_v79_: _dafny.Map
                d_2760_v79_ = _dafny.Map({d_2647_v0_: d_2647_v0_})
                d_2761_v80_: C2
                nw427_ = C2()
                nw427_.ctor__(-859, d_2651_v4_, d_2648_v1_, (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))])
                d_2761_v80_ = nw427_
                d_2762_v81_: _dafny.Array
                nw428_ = _dafny.Array(None, 29)
                nw428_[int(0)] = (d_2700_v41_) + (((_dafny.SeqWithoutIsStrInference([d_2647_v0_])).set(default__.safeIndex(len(d_2758_v77_), len(_dafny.SeqWithoutIsStrInference([d_2647_v0_]))), d_2647_v0_)).set(default__.safeIndex(d_2648_v1_, len((_dafny.SeqWithoutIsStrInference([d_2647_v0_])).set(default__.safeIndex(len(d_2758_v77_), len(_dafny.SeqWithoutIsStrInference([d_2647_v0_]))), d_2647_v0_))), d_2647_v0_))
                nw428_[int(1)] = (d_2759_v78_)[default__.safeIndex((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], len(d_2759_v78_))]
                nw428_[int(2)] = _dafny.SeqWithoutIsStrInference([not(d_2647_v0_)])
                nw428_[int(3)] = (d_2700_v41_).set(default__.safeIndex(len(d_2695_v38_), len(d_2700_v41_)), d_2647_v0_)
                nw428_[int(4)] = d_2700_v41_
                nw428_[int(5)] = d_2700_v41_
                nw428_[int(6)] = d_2700_v41_
                nw428_[int(7)] = (default__.fm30(d_2647_v0_, d_2647_v0_, d_2648_v1_, (d_2700_v41_)[default__.safeIndex((d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], len(d_2700_v41_))], globalState)) + (d_2700_v41_)
                nw428_[int(8)] = _dafny.SeqWithoutIsStrInference([((d_2760_v79_)[d_2647_v0_] if (d_2647_v0_) in (d_2760_v79_) else d_2647_v0_)])
                nw428_[int(9)] = d_2700_v41_
                nw428_[int(10)] = ((_dafny.SeqWithoutIsStrInference([d_2647_v0_, d_2647_v0_])).set(default__.safeIndex((0) - ((d_2649_v2_)[default__.safeIndex(979, len(d_2649_v2_))]), len(_dafny.SeqWithoutIsStrInference([d_2647_v0_, d_2647_v0_]))), d_2647_v0_)) + (d_2700_v41_)
                nw428_[int(11)] = (_dafny.SeqWithoutIsStrInference([d_2647_v0_, d_2647_v0_])) + (_dafny.SeqWithoutIsStrInference([d_2647_v0_]))
                nw428_[int(12)] = d_2700_v41_
                nw428_[int(13)] = d_2700_v41_
                nw428_[int(14)] = (d_2700_v41_) + (d_2700_v41_)
                nw428_[int(15)] = d_2700_v41_
                nw428_[int(16)] = default__.fm30(((d_2760_v79_)[d_2647_v0_] if (d_2647_v0_) in (d_2760_v79_) else (d_2700_v41_)[default__.safeIndex(d_2648_v1_, len(d_2700_v41_))]), d_2647_v0_, d_2648_v1_, d_2647_v0_, globalState)
                nw428_[int(17)] = (d_2700_v41_).set(default__.safeIndex((0) - (d_2648_v1_), len(d_2700_v41_)), d_2647_v0_)
                nw428_[int(18)] = ((d_2700_v41_).set(default__.safeIndex(236, len(d_2700_v41_)), d_2647_v0_)) + (_dafny.SeqWithoutIsStrInference([d_2647_v0_]))
                nw428_[int(19)] = (d_2700_v41_) + (d_2700_v41_)
                nw428_[int(20)] = d_2700_v41_
                nw428_[int(21)] = d_2700_v41_
                nw428_[int(22)] = d_2700_v41_
                nw428_[int(23)] = d_2700_v41_
                nw428_[int(24)] = d_2700_v41_
                nw428_[int(25)] = d_2700_v41_
                nw428_[int(26)] = ((d_2759_v78_)[default__.safeIndex((_dafny.MultiSet([d_2761_v80_, d_2761_v80_])).cardinality, len(d_2759_v78_))]) + (d_2700_v41_)
                nw428_[int(27)] = d_2700_v41_
                nw428_[int(28)] = (d_2700_v41_) + (d_2700_v41_)
                d_2762_v81_ = nw428_
                d_2763_v82_: _dafny.Set
                d_2763_v82_ = _dafny.Set({d_2698_v39_})
                index484_ = default__.safeIndex(888, (d_2751_v73_).length(0))
                rhs452_ = (d_2647_v0_ if (d_2695_v38_).ispropersubset(d_2695_v38_) else d_2647_v0_)
                rhs453_ = len((d_2763_v82_) | (d_2763_v82_))
                rhs454_ = D13_DC37((d_2755_v74_).set(d_2761_v80_.f27, default__.abs(702)))
                rhs455_ = (d_2757_v76_ if (d_2660_v13_).isdisjoint(_dafny.MultiSet([d_2647_v0_])) else d_2757_v76_)
                rhs456_ = d_2762_v81_
                lhs383_ = globalState
                lhs384_ = globalState
                lhs385_ = d_2751_v73_
                lhs386_ = default__.safeIndex(888, (d_2751_v73_).length(0))
                lhs383_.f2 = rhs452_
                lhs384_.f6 = rhs453_
                lhs385_[lhs386_] = rhs454_
                d_2757_v76_ = rhs455_
                d_2762_v81_ = rhs456_
                index485_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                (d_2659_v12_)[index485_] = len(d_2649_v2_)
                d_2698_v39_ = d_2698_v39_
                d_2764_v83_: _dafny.Map
                d_2764_v83_ = _dafny.Map({(0) - (d_2761_v80_.f27): d_2700_v41_})
                d_2765_v84_: D10
                d_2765_v84_ = D10_DC28((d_2764_v83_) | (d_2764_v83_))
                d_2765_v84_ = d_2765_v84_
            index486_ = default__.safeIndex(233, (d_2659_v12_).length(0))
            (d_2659_v12_)[index486_] = (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]
            d_2766_v85_: _dafny.Map
            d_2766_v85_ = _dafny.Map({(True) == (True): d_2647_v0_})
            d_2766_v85_ = (d_2766_v85_).set(True, d_2647_v0_)
            d_2767_v86_: D14
            d_2767_v86_ = D14_DC40(d_2700_v41_)
            pat_let_tv44_ = d_2648_v1_
            pat_let_tv45_ = globalState
            def iife155_(_pat_let45_0):
                def iife156_(d_2768_dt__update__tmp_h0_):
                    def iife157_(_pat_let46_0):
                        def iife158_(d_2769_dt__update_hcf72_h0_):
                            return D14_DC40(d_2769_dt__update_hcf72_h0_)
                        return iife158_(_pat_let46_0)
                    return iife157_(default__.fm22(pat_let_tv44_, pat_let_tv45_))
                return iife156_(_pat_let45_0)
            source40_ = iife155_(d_2767_v86_)
            if source40_.is_DC41:
                d_2770___mcc_h7_ = source40_.cf73
                d_2771___mcc_h8_ = source40_.cf74
                d_2772___mcc_h9_ = source40_.cf75
                d_2773_cf75_ = d_2772___mcc_h9_
                d_2774_cf74_ = d_2771___mcc_h8_
                d_2775_cf73_ = d_2770___mcc_h7_
                d_2776_v87_: _dafny.Set
                d_2776_v87_ = _dafny.Set({True})
                d_2777_v88_: D4
                d_2777_v88_ = D4_DC14(d_2698_v39_, len(d_2766_v85_))
                d_2778_v90_: _dafny.Map
                d_2778_v90_ = _dafny.Map({_dafny.CodePoint('w'): d_2698_v39_})
                d_2779_v92_: _dafny.Map
                def iife159_():
                    def iife161_():
                        coll67_ = _dafny.Map()
                        compr_67_: str
                        for compr_67_ in (d_2778_v90_).keys.Elements:
                            d_2780_v89_: str = compr_67_
                            if (d_2780_v89_) in (d_2778_v90_):
                                coll67_[d_2780_v89_] = (d_2649_v2_)[default__.safeIndex(d_2648_v1_, len(d_2649_v2_))]
                        return _dafny.Map(coll67_)
                    coll65_ = _dafny.Set()
                    def iife160_():
                        coll66_ = _dafny.Map()
                        compr_66_: str
                        for compr_66_ in (d_2778_v90_).keys.Elements:
                            d_2780_v89_: str = compr_66_
                            if (d_2780_v89_) in (d_2778_v90_):
                                coll66_[d_2780_v89_] = (d_2649_v2_)[default__.safeIndex(d_2648_v1_, len(d_2649_v2_))]
                        return _dafny.Map(coll66_)
                    compr_65_: str
                    for compr_65_ in (iife160_()
                    ).keys.Elements:
                        d_2781_v91_: str = compr_65_
                        if (d_2781_v91_) in (iife161_()
                        ):
                            coll65_ = coll65_.union(_dafny.Set([d_2781_v91_]))
                    return _dafny.Set(coll65_)
                d_2779_v92_ = _dafny.Map({d_2698_v39_: (0) - (len(iife159_()
                ))})
                d_2782_v94_: _dafny.Array
                nw429_ = _dafny.Array(None, 24)
                nw429_[int(0)] = d_2776_v87_
                nw429_[int(1)] = d_2776_v87_
                def iife162_():
                    coll68_ = _dafny.Set()
                    compr_68_: int
                    for compr_68_ in _dafny.IntegerRange(-758, 853):
                        d_2783_v93_: int = compr_68_
                        if ((-758) <= (d_2783_v93_)) and ((d_2783_v93_) < (853)):
                            coll68_ = coll68_.union(_dafny.Set([default__.safeDivisionInt(d_2783_v93_, 645)]))
                    return _dafny.Set(coll68_)
                nw429_[int(2)] = (default__.fm32((self).fm13(len(d_2773_cf75_), (d_2777_v88_).cf30, d_2698_v39_, globalState), d_2779_v92_, iife162_()
                , True, globalState)) - (d_2776_v87_)
                nw429_[int(3)] = d_2776_v87_
                nw429_[int(4)] = d_2776_v87_
                nw429_[int(5)] = (d_2776_v87_).intersection(default__.fm32(d_2774_cf74_, d_2779_v92_, d_2695_v38_, d_2647_v0_, globalState))
                nw429_[int(6)] = d_2776_v87_
                nw429_[int(7)] = d_2776_v87_
                nw429_[int(8)] = d_2776_v87_
                nw429_[int(9)] = d_2776_v87_
                nw429_[int(10)] = d_2776_v87_
                nw429_[int(11)] = d_2776_v87_
                nw429_[int(12)] = d_2776_v87_
                nw429_[int(13)] = (d_2776_v87_).intersection(_dafny.Set({d_2775_cf73_}))
                nw429_[int(14)] = d_2776_v87_
                nw429_[int(15)] = (d_2776_v87_) | (d_2776_v87_)
                nw429_[int(16)] = (d_2776_v87_) | (d_2776_v87_)
                nw429_[int(17)] = d_2776_v87_
                nw429_[int(18)] = d_2776_v87_
                nw429_[int(19)] = d_2776_v87_
                nw429_[int(20)] = d_2776_v87_
                nw429_[int(21)] = d_2776_v87_
                nw429_[int(22)] = (d_2776_v87_) | (d_2776_v87_)
                nw429_[int(23)] = d_2776_v87_
                d_2782_v94_ = nw429_
                index487_ = default__.safeIndex(142, (d_2782_v94_).length(0))
                (d_2782_v94_)[index487_] = _dafny.Set({d_2647_v0_})
                d_2784_v95_: _dafny.Array
                nw430_ = _dafny.Array(None, 10)
                nw430_[int(0)] = d_2775_cf73_
                nw430_[int(1)] = d_2774_cf74_
                nw430_[int(2)] = d_2774_cf74_
                nw430_[int(3)] = d_2647_v0_
                nw430_[int(4)] = d_2774_cf74_
                nw430_[int(5)] = d_2774_cf74_
                nw430_[int(6)] = d_2774_cf74_
                nw430_[int(7)] = False
                nw430_[int(8)] = d_2775_cf73_
                nw430_[int(9)] = d_2774_cf74_
                d_2784_v95_ = nw430_
                d_2785_v96_: _dafny.MultiSet
                d_2785_v96_ = _dafny.MultiSet([_dafny.CodePoint('w')])
                d_2786_v97_: C1
                nw431_ = C1()
                nw431_.ctor__(d_2784_v95_, d_2648_v1_, (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], ((d_2785_v96_)[d_2698_v39_] if (d_2698_v39_) in (d_2785_v96_) else (d_2649_v2_)[default__.safeIndex(d_2648_v1_, len(d_2649_v2_))]), d_2774_cf74_, d_2773_cf75_)
                d_2786_v97_ = nw431_
                d_2787_v98_: D9
                d_2787_v98_ = D9_DC26(d_2786_v97_)
                d_2788_v99_: _dafny.Map
                d_2788_v99_ = _dafny.Map({(d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]: d_2787_v98_})
                index488_ = default__.safeIndex(142, (d_2782_v94_).length(0))
                index489_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                rhs457_ = d_2776_v87_
                rhs458_ = d_2788_v99_
                rhs459_ = -589
                rhs460_ = d_2784_v95_
                lhs387_ = d_2782_v94_
                lhs388_ = default__.safeIndex(142, (d_2782_v94_).length(0))
                lhs389_ = d_2659_v12_
                lhs390_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                lhs387_[lhs388_] = rhs457_
                d_2788_v99_ = rhs458_
                lhs389_[lhs390_] = rhs459_
                d_2784_v95_ = rhs460_
                d_2789_v100_: _dafny.Map
                d_2789_v100_ = _dafny.Map({not(d_2647_v0_): d_2784_v95_})
                d_2790_v101_: T0
                nw432_ = C1()
                nw432_.ctor__(((d_2789_v100_)[d_2774_cf74_] if (d_2774_cf74_) in (d_2789_v100_) else d_2784_v95_), (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], d_2774_cf74_, d_2651_v4_)
                d_2790_v101_ = nw432_
                d_2791_v102_: _dafny.Map
                d_2791_v102_ = _dafny.Map({d_2774_cf74_: d_2790_v101_})
                d_2792_v103_: _dafny.Array
                nw433_ = _dafny.Array(None, 19)
                nw433_[int(0)] = d_2790_v101_
                nw433_[int(1)] = d_2790_v101_
                nw433_[int(2)] = ((d_2791_v102_)[d_2647_v0_] if (d_2647_v0_) in (d_2791_v102_) else d_2790_v101_)
                nw433_[int(3)] = d_2790_v101_
                nw433_[int(4)] = d_2790_v101_
                nw433_[int(5)] = d_2790_v101_
                nw433_[int(6)] = d_2790_v101_
                nw433_[int(7)] = d_2790_v101_
                nw433_[int(8)] = d_2790_v101_
                nw433_[int(9)] = d_2790_v101_
                nw433_[int(10)] = d_2790_v101_
                nw433_[int(11)] = d_2790_v101_
                nw433_[int(12)] = d_2790_v101_
                nw433_[int(13)] = d_2790_v101_
                nw433_[int(14)] = d_2790_v101_
                nw433_[int(15)] = (d_2790_v101_ if d_2775_cf73_ else d_2790_v101_)
                nw433_[int(16)] = d_2790_v101_
                nw433_[int(17)] = d_2790_v101_
                nw433_[int(18)] = d_2790_v101_
                d_2792_v103_ = nw433_
                index490_ = default__.safeIndex(429, (d_2792_v103_).length(0))
                (d_2792_v103_)[index490_] = d_2790_v101_
                d_2793_v104_: _dafny.Seq
                d_2793_v104_ = _dafny.SeqWithoutIsStrInference([(d_2790_v101_ if d_2775_cf73_ else d_2790_v101_)])
                index491_ = default__.safeIndex(429, (d_2792_v103_).length(0))
                (d_2792_v103_)[index491_] = (d_2793_v104_)[default__.safeIndex((d_2790_v101_).f11, len(d_2793_v104_))]
                (globalState).f6 = (d_2790_v101_).f11
                index492_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                (d_2659_v12_)[index492_] = default__.fm0(not(d_2774_cf74_), globalState)
            elif True:
                d_2794___mcc_h10_ = source40_.cf72
                d_2795_cf72_ = d_2794___mcc_h10_
                d_2796_v105_: _dafny.Map
                d_2796_v105_ = _dafny.Map({(_dafny.CodePoint('a') if True else d_2698_v39_): default__.fm0(d_2647_v0_, globalState)})
                d_2796_v105_ = (d_2796_v105_) | (d_2796_v105_)
                d_2797_v106_: bool
                d_2798_v107_: _dafny.Seq
                out81_: bool
                out82_: _dafny.Seq
                out81_, out82_ = default__.m0(d_2647_v0_, d_2647_v0_, (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))], 283, globalState)
                d_2797_v106_ = out81_
                d_2798_v107_ = out82_
                index493_ = default__.safeIndex(233, (d_2659_v12_).length(0))
                (d_2659_v12_)[index493_] = (d_2659_v12_)[default__.safeIndex(233, (d_2659_v12_).length(0))]
                d_2799_v108_: _dafny.Map
                d_2799_v108_ = _dafny.Map({_dafny.Map({d_2647_v0_: d_2797_v106_}): not (d_2797_v106_) or (d_2647_v0_)})
                d_2800_v109_: _dafny.Set
                d_2800_v109_ = _dafny.Set({d_2797_v106_, d_2797_v106_, d_2797_v106_, False, d_2797_v106_})
                d_2799_v108_ = (d_2799_v108_).set(_dafny.Map({d_2797_v106_: d_2647_v0_}), (_dafny.Set({d_2797_v106_})).isdisjoint(d_2800_v109_))
        r0 = not(d_2647_v0_)
        r1 = d_2659_v12_
        d_2801_v110_: _dafny.MultiSet
        d_2801_v110_ = _dafny.MultiSet([(d_2649_v2_) + (d_2649_v2_), d_2649_v2_])
        r2 = d_2801_v110_
        return r0, r1, r2


class C6(T2, T1, T0):
    def  __init__(self):
        self._f11: int = int(0)
        self._f12: int = int(0)
        self._f14: int = int(0)
        self._f15: bool = False
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f23: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
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
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f23, f14, f15, f13, f11, f12):
        (self)._f23 = f23
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f13 = f13
        (self)._f11 = f11
        (self)._f12 = f12

    def fm2(self, p0, p1, p2, p3, globalState):
        return len(((self).f13 if True else ((self).f13).set(default__.safeIndex((self).f11, len((self).f13)), _dafny.CodePoint('h'))))

    def fm3(self, p0, p1, p2, globalState):
        return ((_dafny.Map({(self).f15: 426})) | (_dafny.Map({True: (self).f14}))) | ((_dafny.Map({(self).f15: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "niiwk")))})) | (_dafny.Map({(self).f15: (self).f14})))

    def fm1(self, p0, p1, p2, globalState):
        return (802) < ((len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oaa"))]))) - ((self).f12))

    def fm12(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Map({len(_dafny.Set({(self).f12})): (self).f14})) | (_dafny.Map({-16: (self).f12}))) | (_dafny.Map({(self).f11: (self).f14}))

    def m2(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        d_2802_v0_: str
        d_2802_v0_ = _dafny.CodePoint('s')
        d_2803_v1_: _dafny.Map
        d_2803_v1_ = _dafny.Map({(self).f15: (self).f14})
        d_2804_v2_: _dafny.MultiSet
        d_2804_v2_ = _dafny.MultiSet([d_2802_v0_])
        d_2805_v3_: C3
        nw434_ = C3()
        nw434_.ctor__((d_2802_v0_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnbi"))), _dafny.SeqWithoutIsStrInference([d_2802_v0_ for d_2806_i0_ in range(default__.abs(814))]), (self).f13, default__.safeDivisionInt((self).f14, (self).f12), (self).f15, (self).f13, (((d_2803_v1_)[not((self).f15)] if (not((self).f15)) in (d_2803_v1_) else (self).f11)) + ((d_2804_v2_).cardinality), (self).f12)
        d_2805_v3_ = nw434_
        d_2807_v4_: D14
        d_2807_v4_ = D14_DC41((d_2805_v3_).f25, True, (self).f13)
        def iife163_(_pat_let47_0):
            def iife164_(d_2808_dt__update__tmp_h0_):
                def iife165_(_pat_let48_0):
                    def iife166_(d_2809_dt__update_hcf74_h0_):
                        return D14_DC41((d_2808_dt__update__tmp_h0_).cf73, d_2809_dt__update_hcf74_h0_, (d_2808_dt__update__tmp_h0_).cf75)
                    return iife166_(_pat_let48_0)
                return iife165_((self).f15)
            return iife164_(_pat_let47_0)
        r0 = not((iife163_(d_2807_v4_)).cf73)
        hi9_ = (self).f14
        for d_2810_i1_ in range((self).f11, hi9_):
            d_2811_v5_: _dafny.Set
            d_2811_v5_ = _dafny.Set({(d_2805_v3_).f25})
            d_2812_v6_: _dafny.Map
            d_2812_v6_ = _dafny.Map({(self).f14: (self).f15})
            d_2811_v5_ = ((d_2811_v5_) | (d_2811_v5_)) - (_dafny.Set({((d_2812_v6_)[(self).f12] if ((self).f12) in (d_2812_v6_) else not((self).f15)), (d_2805_v3_).f25, (self).f15, (self).f15}))
            (globalState).f2 = not((d_2805_v3_).f25)
            d_2813_v7_: _dafny.Seq
            d_2813_v7_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2814_v8_: _dafny.Array
            nw435_ = _dafny.Array(None, 3)
            nw435_[int(0)] = (d_2813_v7_)[default__.safeIndex((self).f12, len(d_2813_v7_))]
            nw435_[int(1)] = p0
            nw435_[int(2)] = p0
            d_2814_v8_ = nw435_
            nw436_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_2814_v8_ = nw436_
            d_2815_v9_: _dafny.Array
            nw437_ = _dafny.Array(False, 4)
            d_2815_v9_ = nw437_
            d_2815_v9_ = d_2815_v9_
        d_2816_v10_: D10
        d_2816_v10_ = D10_DC30((self).f15, (d_2805_v3_).f25, ((self).f13).set(default__.safeIndex((self).f12, len((self).f13)), d_2802_v0_), (d_2805_v3_).f25)
        def lambda112_(source41_):
            if source41_.is_DC29:
                return False
            elif source41_.is_DC30:
                d_2817___mcc_h0_ = source41_.cf48
                d_2818___mcc_h1_ = source41_.cf49
                d_2819___mcc_h2_ = source41_.cf50
                d_2820___mcc_h3_ = source41_.cf51
                d_2821_cf51_ = d_2820___mcc_h3_
                d_2822_cf50_ = d_2819___mcc_h2_
                d_2823_cf49_ = d_2818___mcc_h1_
                d_2824_cf48_ = d_2817___mcc_h0_
                return ((self).f12) > (797)
            elif source41_.is_DC31:
                d_2825___mcc_h4_ = source41_.cf52
                d_2826___mcc_h5_ = source41_.cf53
                d_2827___mcc_h6_ = source41_.cf54
                d_2828___mcc_h7_ = source41_.cf55
                d_2829_cf55_ = d_2828___mcc_h7_
                d_2830_cf54_ = d_2827___mcc_h6_
                d_2831_cf53_ = d_2826___mcc_h5_
                d_2832_cf52_ = d_2825___mcc_h4_
                return (self).f15
            elif source41_.is_DC28:
                d_2833___mcc_h8_ = source41_.cf47
                d_2834_cf47_ = d_2833___mcc_h8_
                return True
            elif True:
                d_2835___mcc_h9_ = source41_.cf56
                d_2836_cf56_ = d_2835___mcc_h9_
                return not(((self).f11) >= ((self).f12))

        (globalState).f2 = lambda112_(d_2816_v10_)
        d_2837_v11_: _dafny.Map
        d_2837_v11_ = _dafny.Map({(d_2805_v3_).f26: 14})
        d_2838_v12_: _dafny.Map
        d_2838_v12_ = _dafny.Map({(d_2805_v3_).f25: (d_2805_v3_).f25})
        hi10_ = len(_dafny.Map({(self).f11: D3_DC10((self).f15, (self).f14, d_2837_v11_, (d_2805_v3_).f26, len(d_2838_v12_))}))
        for d_2839_i2_ in range((self).f14, hi10_):
            if (d_2805_v3_).f25:
                d_2840_v13_: C1
                nw438_ = C1()
                nw438_.ctor__(p0, d_2839_i2_, d_2839_i2_, len(_dafny.Map({(0) - ((self).f14): (self).f11})), ((d_2838_v12_)[False] if (False) in (d_2838_v12_) else (d_2805_v3_).f25), (d_2805_v3_).f26)
                d_2840_v13_ = nw438_
                d_2841_v14_: _dafny.Array
                def lambda113_(d_2842_i3_):
                    return default__.safeModuloInt(d_2842_i3_, (self).f11)

                init67_ = lambda113_
                nw439_ = _dafny.Array(None, 24)
                for i0_67_ in range(nw439_.length(0)):
                    nw439_[i0_67_] = init67_(i0_67_)
                d_2841_v14_ = nw439_
                d_2843_v15_: _dafny.Seq
                d_2843_v15_ = _dafny.SeqWithoutIsStrInference([d_2841_v14_])
                d_2844_v16_: D3
                d_2844_v16_ = D3_DC9(d_2843_v15_)
                d_2845_v17_: _dafny.Set
                d_2845_v17_ = _dafny.Set({_dafny.Map({d_2840_v13_: d_2844_v16_})})
                d_2846_v18_: _dafny.Map
                d_2846_v18_ = _dafny.Map({(self).f11: (d_2845_v17_ if (d_2805_v3_).f25 else d_2845_v17_)})
                d_2846_v18_ = (d_2846_v18_).set((self).f11, d_2845_v17_)
                d_2802_v0_ = d_2802_v0_
                (globalState).f2 = True
                d_2847_v19_: _dafny.Seq
                d_2847_v19_ = _dafny.SeqWithoutIsStrInference([(d_2805_v3_).f25, (d_2805_v3_).f25])
                d_2847_v19_ = ((d_2847_v19_) + (d_2847_v19_)) + ((_dafny.SeqWithoutIsStrInference([(d_2805_v3_).f25])) + (_dafny.SeqWithoutIsStrInference([(d_2805_v3_).f25])))
                d_2802_v0_ = d_2802_v0_
            elif True:
                d_2848_v20_: _dafny.Set
                d_2848_v20_ = _dafny.Set({(self).f11})
                d_2848_v20_ = d_2848_v20_
                d_2849_v21_: C3
                nw440_ = C3()
                nw440_.ctor__(not((self).f15), (self).f13, (d_2805_v3_).f26, (self).f14, (d_2805_v3_).f25, (self).f13, (0) - ((0) - ((self).f11)), d_2839_i2_)
                d_2849_v21_ = nw440_
                index494_ = default__.safeIndex(600, (p0).length(0))
                (p0)[index494_] = not((d_2805_v3_).f25)
                index495_ = default__.safeIndex(407, (p0).length(0))
                (p0)[index495_] = ((self).f14) >= (d_2839_i2_)
                index496_ = default__.safeIndex(600, (p0).length(0))
                index497_ = default__.safeIndex(407, (p0).length(0))
                rhs461_ = default__.safeDivisionInt(((self).f14) * ((self).f14), d_2839_i2_)
                rhs462_ = not((d_2805_v3_).f25)
                rhs463_ = (d_2805_v3_).f25
                rhs464_ = (d_2805_v3_).f25
                rhs465_ = (d_2849_v21_).f25
                lhs391_ = globalState
                lhs392_ = globalState
                lhs393_ = p0
                lhs394_ = default__.safeIndex(600, (p0).length(0))
                lhs395_ = p0
                lhs396_ = default__.safeIndex(407, (p0).length(0))
                lhs397_ = globalState
                lhs391_.f6 = rhs461_
                lhs392_.f2 = rhs462_
                lhs393_[lhs394_] = rhs463_
                lhs395_[lhs396_] = rhs464_
                lhs397_.f2 = rhs465_
                d_2850_v22_: _dafny.Array
                def lambda114_(d_2851_i4_):
                    return (self).f15

                init68_ = lambda114_
                nw441_ = _dafny.Array(None, 5)
                for i0_68_ in range(nw441_.length(0)):
                    nw441_[i0_68_] = init68_(i0_68_)
                d_2850_v22_ = nw441_
                d_2852_v23_: _dafny.Set
                d_2852_v23_ = _dafny.Set({d_2850_v22_, d_2850_v22_, d_2850_v22_})
                d_2853_v24_: D19
                d_2853_v24_ = D19_DC55(d_2852_v23_)
                d_2854_v25_: _dafny.Map
                d_2854_v25_ = _dafny.Map({866: (d_2805_v3_).f25})
                d_2855_v26_: bool
                d_2856_v27_: _dafny.Seq
                out83_: bool
                out84_: _dafny.Seq
                out83_, out84_ = default__.m0(((d_2849_v21_).f25) or ((d_2805_v3_).f25), True, len((d_2853_v24_).cf104), (0) - (((self).f12) + (len(d_2854_v25_))), globalState)
                d_2855_v26_ = out83_
                d_2856_v27_ = out84_
                rhs466_ = (0) - ((self).f11)
                rhs467_ = (self).f11
                lhs398_ = globalState
                lhs399_ = globalState
                lhs398_.f6 = rhs466_
                lhs399_.f6 = rhs467_
            d_2857_v28_: D17
            d_2857_v28_ = D17_DC47(_dafny.Map({d_2839_i2_: len(_dafny.Map({(self).f15: len(_dafny.SeqWithoutIsStrInference([d_2802_v0_ for d_2858_i5_ in range(default__.abs(211))]))}))}))
            d_2859_v29_: _dafny.Map
            d_2859_v29_ = _dafny.Map({(self).f14: len(d_2803_v1_)})
            pat_let_tv46_ = d_2859_v29_
            pat_let_tv47_ = d_2859_v29_
            d_2860_v30_: _dafny.MultiSet
            def iife167_(_pat_let49_0):
                def iife168_(d_2861_dt__update__tmp_h1_):
                    def iife169_(_pat_let50_0):
                        def iife170_(d_2862_dt__update_hcf89_h0_):
                            return D17_DC47(d_2862_dt__update_hcf89_h0_)
                        return iife170_(_pat_let50_0)
                    return iife169_(pat_let_tv46_)
                return iife168_(_pat_let49_0)
            def iife171_(_pat_let51_0):
                def iife172_(d_2863_dt__update__tmp_h2_):
                    def iife173_(_pat_let52_0):
                        def iife174_(d_2864_dt__update_hcf89_h1_):
                            return D17_DC47(d_2864_dt__update_hcf89_h1_)
                        return iife174_(_pat_let52_0)
                    return iife173_(pat_let_tv47_)
                return iife172_(_pat_let51_0)
            d_2860_v30_ = _dafny.MultiSet([iife167_(d_2857_v28_), iife171_(d_2857_v28_), D17_DC47(_dafny.Map({(self).f11: len(_dafny.SeqWithoutIsStrInference([(self).f14 for d_2865_i6_ in range(default__.abs(-61))]))}))])
            pat_let_tv48_ = d_2859_v29_
            pat_let_tv49_ = d_2859_v29_
            def iife175_(_pat_let53_0):
                def iife176_(d_2866_dt__update__tmp_h4_):
                    def iife177_(_pat_let54_0):
                        def iife178_(d_2867_dt__update_hcf89_h3_):
                            return D17_DC47(d_2867_dt__update_hcf89_h3_)
                        return iife178_(_pat_let54_0)
                    return iife177_(pat_let_tv48_)
                return iife176_(_pat_let53_0)
            def iife179_(_pat_let55_0):
                def iife180_(d_2868_dt__update__tmp_h3_):
                    def iife181_(_pat_let56_0):
                        def iife182_(d_2869_dt__update_hcf89_h2_):
                            return D17_DC47(d_2869_dt__update_hcf89_h2_)
                        return iife182_(_pat_let56_0)
                    return iife181_(pat_let_tv49_)
                return iife180_(_pat_let55_0)
            (globalState).f2 = (((d_2860_v30_)[iife175_(d_2857_v28_)] if (iife179_(d_2857_v28_)) in (d_2860_v30_) else (self).f14)) < ((self).f11)
            if (d_2805_v3_).f25:
                index498_ = default__.safeIndex(441, (p0).length(0))
                (p0)[index498_] = (self).f15
                d_2870_v31_: _dafny.Array
                nw442_ = _dafny.Array(_dafny.Map({}), 13)
                d_2870_v31_ = nw442_
                d_2871_v32_: _dafny.Seq
                d_2871_v32_ = _dafny.SeqWithoutIsStrInference([(self).f12, 546])
                d_2872_v33_: _dafny.Set
                d_2872_v33_ = _dafny.Set({default__.fm48(globalState), d_2871_v32_})
                d_2873_v34_: _dafny.Seq
                d_2873_v34_ = _dafny.SeqWithoutIsStrInference([(d_2805_v3_).f25])
                index499_ = default__.safeIndex(441, (p0).length(0))
                rhs468_ = not((d_2872_v33_).issubset(d_2872_v33_))
                rhs469_ = d_2870_v31_
                rhs470_ = ((self).f12 if ((d_2805_v3_).f25) or ((self).f15) else (0) - (len((d_2873_v34_) + (d_2873_v34_))))
                lhs400_ = p0
                lhs401_ = default__.safeIndex(441, (p0).length(0))
                lhs402_ = globalState
                lhs400_[lhs401_] = rhs468_
                d_2870_v31_ = rhs469_
                lhs402_.f6 = rhs470_
                d_2874_v35_: _dafny.Array
                nw443_ = _dafny.Array(None, 13)
                nw443_[int(0)] = (p0)[default__.safeIndex(441, (p0).length(0))]
                nw443_[int(1)] = (self).f15
                nw443_[int(2)] = (p0)[default__.safeIndex(441, (p0).length(0))]
                nw443_[int(3)] = (self).f15
                nw443_[int(4)] = (d_2805_v3_).f25
                nw443_[int(5)] = not(True)
                nw443_[int(6)] = (p0)[default__.safeIndex(441, (p0).length(0))]
                nw443_[int(7)] = (d_2805_v3_).f25
                nw443_[int(8)] = (self).f15
                nw443_[int(9)] = (d_2805_v3_).f25
                nw443_[int(10)] = (self).f15
                nw443_[int(11)] = not((d_2805_v3_).f25)
                nw443_[int(12)] = (d_2805_v3_).f25
                d_2874_v35_ = nw443_
                d_2875_v36_: _dafny.Seq
                d_2875_v36_ = _dafny.SeqWithoutIsStrInference([d_2874_v35_])
                d_2876_v37_: C1
                nw444_ = C1()
                nw444_.ctor__(d_2874_v35_, d_2839_i2_, len(d_2875_v36_), 518, (d_2805_v3_).f25, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))
                d_2876_v37_ = nw444_
                d_2876_v37_ = d_2876_v37_
                d_2877_v38_: C5
                nw445_ = C5()
                nw445_.ctor__()
                d_2877_v38_ = nw445_
                d_2878_v39_: _dafny.Array
                nw446_ = _dafny.Array(None, 8)
                nw446_[int(0)] = (_dafny.MultiSet([len((self).f13), len(d_2871_v32_)])).cardinality
                nw446_[int(1)] = d_2839_i2_
                nw446_[int(2)] = (self).f11
                nw446_[int(3)] = (self).f14
                nw446_[int(4)] = (self).f11
                nw446_[int(5)] = (self).f12
                nw446_[int(6)] = d_2839_i2_
                nw446_[int(7)] = (0) - ((self).f12)
                d_2878_v39_ = nw446_
                d_2879_v40_: _dafny.Map
                d_2879_v40_ = _dafny.Map({(d_2805_v3_).f25: d_2878_v39_})
                d_2880_v41_: _dafny.Map
                d_2880_v41_ = _dafny.Map({d_2877_v38_: d_2873_v34_})
                d_2881_v42_: _dafny.Array
                nw447_ = _dafny.Array(None, 19)
                nw447_[int(0)] = (d_2873_v34_) + (_dafny.SeqWithoutIsStrInference([(default__.fm39(globalState)).cf15, (d_2805_v3_).f25]))
                nw447_[int(1)] = d_2873_v34_
                nw447_[int(2)] = _dafny.SeqWithoutIsStrInference([False, (d_2805_v3_).f25, False])
                nw447_[int(3)] = d_2873_v34_
                nw447_[int(4)] = d_2873_v34_
                nw447_[int(5)] = (d_2873_v34_) + (d_2873_v34_)
                nw447_[int(6)] = d_2873_v34_
                nw447_[int(7)] = _dafny.SeqWithoutIsStrInference([not((self).f15)])
                nw447_[int(8)] = d_2873_v34_
                nw447_[int(9)] = d_2873_v34_
                nw447_[int(10)] = (d_2873_v34_) + (d_2873_v34_)
                nw447_[int(11)] = d_2873_v34_
                nw447_[int(12)] = d_2873_v34_
                nw447_[int(13)] = d_2873_v34_
                nw447_[int(14)] = _dafny.SeqWithoutIsStrInference([False, not((p0)[default__.safeIndex(441, (p0).length(0))])])
                nw447_[int(15)] = (d_2873_v34_).set(default__.safeIndex(len(d_2879_v40_), len(d_2873_v34_)), (d_2805_v3_).f25)
                nw447_[int(16)] = d_2873_v34_
                nw447_[int(17)] = ((d_2880_v41_)[d_2877_v38_] if (d_2877_v38_) in (d_2880_v41_) else d_2873_v34_)
                nw447_[int(18)] = d_2873_v34_
                d_2881_v42_ = nw447_
                d_2881_v42_ = d_2881_v42_
                (globalState).f6 = default__.safeDivisionInt(default__.safeModuloInt((self).f11, len(d_2838_v12_)), d_2839_i2_)
            elif True:
                d_2882_v43_: _dafny.Set
                d_2882_v43_ = _dafny.Set({(self).f12, 120})
                d_2882_v43_ = d_2882_v43_
                (globalState).f6 = d_2839_i2_
                rhs471_ = (0) - (d_2839_i2_)
                rhs472_ = ((d_2803_v1_)[(self).fm1(d_2839_i2_, (self).f15, (d_2805_v3_).fm26((self).f15, globalState), globalState)] if ((self).fm1(d_2839_i2_, (self).f15, (d_2805_v3_).fm26((self).f15, globalState), globalState)) in (d_2803_v1_) else d_2839_i2_)
                lhs403_ = globalState
                lhs404_ = globalState
                lhs403_.f6 = rhs471_
                lhs404_.f6 = rhs472_
                d_2883_v44_: _dafny.Seq
                d_2883_v44_ = _dafny.SeqWithoutIsStrInference([(d_2839_i2_ if (self).f15 else (self).f12), (self).f14])
                d_2884_v45_: _dafny.Seq
                d_2884_v45_ = _dafny.SeqWithoutIsStrInference([d_2883_v44_, d_2883_v44_, (_dafny.SeqWithoutIsStrInference([271 for d_2885_i7_ in range(default__.abs(492))])).set(default__.safeIndex((self).f11, len(_dafny.SeqWithoutIsStrInference([271 for d_2886_i7_ in range(default__.abs(492))]))), (self).f11), d_2883_v44_])
                d_2883_v44_ = ((d_2884_v45_)[default__.safeIndex(d_2839_i2_, len(d_2884_v45_))]) + (d_2883_v44_)
                (globalState).f2 = (self).f15
            if (self).f15:
                d_2887_v46_: _dafny.Map
                d_2887_v46_ = _dafny.Map({p0: d_2839_i2_})
                d_2888_v47_: D16
                d_2888_v47_ = D16_DC45(p0, d_2802_v0_, (0) - ((self).f12), ((d_2887_v46_)[p0] if (p0) in (d_2887_v46_) else (self).f14), not((self).f15))
                (globalState).f7 = (d_2805_v3_).fm1((self).f14, not((d_2805_v3_).f25), (d_2888_v47_).cf85, globalState)
                d_2889_v48_: C2
                nw448_ = C2()
                nw448_.ctor__((self).f11, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdxmf")), (self).f11, (self).f12)
                d_2889_v48_ = nw448_
                d_2890_v49_: _dafny.Map
                d_2890_v49_ = _dafny.Map({d_2802_v0_: (self).f12})
                d_2891_v50_: _dafny.Map
                d_2891_v50_ = _dafny.Map({d_2889_v48_: (0) - (len(d_2890_v49_))})
                d_2892_v51_: _dafny.Array
                nw449_ = _dafny.Array(None, 2)
                nw449_[int(0)] = (d_2839_i2_) - ((self).f12)
                nw449_[int(1)] = len((_dafny.Map({d_2889_v48_: d_2889_v48_.f27})) | (d_2891_v50_))
                d_2892_v51_ = nw449_
                index500_ = default__.safeIndex(871, (d_2892_v51_).length(0))
                (d_2892_v51_)[index500_] = ((self).f11) * ((self).f14)
                d_2893_v52_: _dafny.Map
                d_2893_v52_ = _dafny.Map({d_2839_i2_: (self).f15})
                d_2894_v53_: _dafny.Seq
                d_2894_v53_ = _dafny.SeqWithoutIsStrInference([809])
                d_2895_v54_: D11
                d_2895_v54_ = D11_DC34(((d_2893_v52_)[len(d_2894_v53_)] if (len(d_2894_v53_)) in (d_2893_v52_) else (d_2805_v3_).f25), (d_2805_v3_).f25, d_2839_i2_, not((d_2805_v3_).f25), (self).f15)
                index501_ = default__.safeIndex(871, (d_2892_v51_).length(0))
                (d_2892_v51_)[index501_] = (d_2895_v54_).cf60
                index502_ = default__.safeIndex(871, (d_2892_v51_).length(0))
                (d_2892_v51_)[index502_] = len((self).f13)
                (d_2889_v48_).f27 = (self).f14
                d_2896_v55_: C4
                nw450_ = C4()
                nw450_.ctor__((d_2805_v3_).f26, (0) - (d_2839_i2_), ((self).f12) >= (d_2839_i2_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pesc")), (597 if (d_2805_v3_).f25 else d_2889_v48_.f27), (d_2892_v51_)[default__.safeIndex(871, (d_2892_v51_).length(0))])
                d_2896_v55_ = nw450_
            elif True:
                d_2897_v56_: _dafny.Seq
                d_2897_v56_ = _dafny.SeqWithoutIsStrInference([941, default__.safeDivisionInt((self).f11, default__.fm0((d_2805_v3_).f25, globalState)), (self).f12])
                (globalState).f6 = (d_2897_v56_)[default__.safeIndex((self).f12, len(d_2897_v56_))]
                (globalState).f7 = (d_2805_v3_).f25
                index503_ = default__.safeIndex(656, (p0).length(0))
                (p0)[index503_] = (len(d_2897_v56_)) not in (d_2897_v56_)
                index504_ = default__.safeIndex(656, (p0).length(0))
                (p0)[index504_] = True
                (globalState).f2 = (False if (d_2805_v3_).f25 else (p0)[default__.safeIndex(656, (p0).length(0))])
                (globalState).f6 = default__.fm0((d_2805_v3_).f25, globalState)
        d_2898_v57_: _dafny.Set
        d_2898_v57_ = _dafny.Set({d_2802_v0_, d_2802_v0_, d_2802_v0_})
        d_2899_v58_: _dafny.Map
        d_2899_v58_ = _dafny.Map({d_2802_v0_: (self).f15})
        d_2900_v60_: _dafny.Array
        nw451_ = _dafny.Array(None, 16)
        nw451_[int(0)] = d_2898_v57_
        nw451_[int(1)] = d_2898_v57_
        nw451_[int(2)] = d_2898_v57_
        nw451_[int(3)] = d_2898_v57_
        nw451_[int(4)] = d_2898_v57_
        nw451_[int(5)] = d_2898_v57_
        nw451_[int(6)] = default__.fm68((self).f12, ((self).f13).set(default__.safeIndex((d_2805_v3_).fm2((0) - ((self).f12), (d_2805_v3_).f25, (self).f14, (self).f15, globalState), len((self).f13)), d_2802_v0_), globalState)
        nw451_[int(7)] = d_2898_v57_
        nw451_[int(8)] = d_2898_v57_
        nw451_[int(9)] = (d_2898_v57_).intersection(d_2898_v57_)
        nw451_[int(10)] = _dafny.Set({d_2802_v0_})
        nw451_[int(11)] = (d_2898_v57_).intersection(d_2898_v57_)
        nw451_[int(12)] = (d_2898_v57_).intersection(d_2898_v57_)
        nw451_[int(13)] = (d_2898_v57_ if (d_2805_v3_).f25 else d_2898_v57_)
        def iife183_():
            coll69_ = _dafny.Set()
            compr_69_: str
            for compr_69_ in (d_2899_v58_).keys.Elements:
                d_2901_v59_: str = compr_69_
                if (d_2901_v59_) in (d_2899_v58_):
                    coll69_ = coll69_.union(_dafny.Set([d_2901_v59_]))
            return _dafny.Set(coll69_)
        nw451_[int(14)] = iife183_()
        
        nw451_[int(15)] = _dafny.Set({d_2802_v0_})
        d_2900_v60_ = nw451_
        d_2902_v61_: _dafny.Array
        nw452_ = _dafny.Array(int(0), 13)
        d_2902_v61_ = nw452_
        d_2903_v62_: _dafny.Map
        d_2903_v62_ = _dafny.Map({d_2902_v61_: (self).f12})
        d_2904_v63_: _dafny.Map
        d_2904_v63_ = _dafny.Map({default__.safeDivisionInt(((d_2903_v62_)[d_2902_v61_] if (d_2902_v61_) in (d_2903_v62_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yajajypjw")))), (d_2805_v3_).fm2((self).f11, (d_2805_v3_).f25, len(d_2838_v12_), (self).f15, globalState)): d_2900_v60_})
        rhs473_ = (self).f11
        rhs474_ = (0) - ((self).f14)
        rhs475_ = ((d_2904_v63_)[(self).f11] if ((self).f11) in (d_2904_v63_) else d_2900_v60_)
        rhs476_ = (self).f12
        lhs405_ = globalState
        lhs406_ = globalState
        lhs407_ = globalState
        lhs405_.f6 = rhs473_
        lhs406_.f6 = rhs474_
        d_2900_v60_ = rhs475_
        lhs407_.f6 = rhs476_
        d_2905_v64_: _dafny.Set
        d_2905_v64_ = _dafny.Set({not((d_2805_v3_).f25), (d_2805_v3_).f25, (d_2805_v3_).f25, (d_2805_v3_).f25, (self).f15})
        d_2906_v65_: _dafny.Set
        d_2906_v65_ = _dafny.Set({d_2905_v64_, d_2905_v64_, d_2905_v64_, d_2905_v64_})
        def iife184_():
            coll70_ = _dafny.Set()
            compr_70_: _dafny.Set
            for compr_70_ in (_dafny.SeqWithoutIsStrInference([d_2905_v64_])).Elements:
                d_2907_v66_: _dafny.Set = compr_70_
                if (d_2907_v66_) in (_dafny.SeqWithoutIsStrInference([d_2905_v64_])):
                    coll70_ = coll70_.union(_dafny.Set([d_2907_v66_]))
            return _dafny.Set(coll70_)
        r0 = (iife184_()
        ).ispropersubset((d_2906_v65_) | (_dafny.Set({d_2905_v64_, d_2905_v64_})))
        r1 = p1
        return r0, r1

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        d_2908_v0_: _dafny.Seq
        d_2908_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        hi11_ = len(d_2908_v0_)
        for d_2909_i0_ in range(p1, hi11_):
            d_2910_v1_: _dafny.Seq
            d_2910_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ti"))
            d_2910_v1_ = (((self).f13).set(default__.safeIndex(p3, len((self).f13)), ((self).f13)[default__.safeIndex((self).f12, len((self).f13))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tx")))
            (globalState).f7 = (default__.safeDivisionInt(p1, (self).f11)) != (d_2909_i0_)
            d_2911_v2_: _dafny.Map
            d_2911_v2_ = _dafny.Map({p2: (self).f14})
            d_2912_v3_: _dafny.Array
            def lambda115_(d_2913_p3_):
                def lambda116_(d_2914_i1_):
                    return (d_2914_i1_) - (d_2913_p3_)

                return lambda116_

            init69_ = lambda115_(p3)
            nw453_ = _dafny.Array(None, 6)
            for i0_69_ in range(nw453_.length(0)):
                nw453_[i0_69_] = init69_(i0_69_)
            d_2912_v3_ = nw453_
            d_2915_v4_: _dafny.Map
            d_2915_v4_ = _dafny.Map({d_2911_v2_: d_2912_v3_})
            d_2916_v5_: _dafny.Array
            nw454_ = _dafny.Array(False, 14)
            d_2916_v5_ = nw454_
            d_2917_v6_: _dafny.Seq
            d_2917_v6_ = _dafny.SeqWithoutIsStrInference([d_2911_v2_, _dafny.Map({p2: p3})])
            d_2918_v7_: _dafny.Seq
            d_2918_v7_ = _dafny.SeqWithoutIsStrInference([p2, (self).f15])
            d_2919_v8_: _dafny.Seq
            d_2919_v8_ = _dafny.SeqWithoutIsStrInference([d_2912_v3_])
            rhs477_ = (d_2915_v4_) | ((_dafny.Map({(d_2917_v6_)[default__.safeIndex(len(d_2918_v7_), len(d_2917_v6_))]: d_2912_v3_})) | (d_2915_v4_))
            rhs478_ = default__.safeDivisionInt(len((d_2919_v8_) + (d_2919_v8_)), (self).f14)
            rhs479_ = d_2916_v5_
            rhs480_ = (default__.safeDivisionInt(p1, 559)) + (default__.safeModuloInt((self).f14, (self).f12))
            rhs481_ = (p1) - (906)
            lhs408_ = globalState
            d_2915_v4_ = rhs477_
            r2 = rhs478_
            d_2916_v5_ = rhs479_
            lhs408_.f6 = rhs480_
            r1 = rhs481_
            d_2920_v9_: _dafny.Map
            d_2920_v9_ = _dafny.Map({False: (self).f15})
            index505_ = default__.safeIndex(62, (d_2912_v3_).length(0))
            (d_2912_v3_)[index505_] = len(d_2920_v9_)
            index506_ = default__.safeIndex(62, (d_2912_v3_).length(0))
            (d_2912_v3_)[index506_] = (0) - ((0) - (p1))
        if (self).f15:
            d_2921_v10_: _dafny.Map
            d_2921_v10_ = _dafny.Map({default__.fm0(False, globalState): (self).f15})
            d_2921_v10_ = (d_2921_v10_).set(604, (p2) and ((self).fm1((self).f14, p2, p2, globalState)))
            if p2:
                (globalState).f2 = (self).f15
                d_2922_v11_: _dafny.Array
                def lambda117_(d_2923_i2_):
                    return (self).f15

                init70_ = lambda117_
                nw455_ = _dafny.Array(None, 21)
                for i0_70_ in range(nw455_.length(0)):
                    nw455_[i0_70_] = init70_(i0_70_)
                d_2922_v11_ = nw455_
                d_2924_v12_: _dafny.Array
                nw456_ = _dafny.Array(None, 23)
                nw456_[int(0)] = (d_2922_v11_ if (self).f15 else d_2922_v11_)
                nw456_[int(1)] = d_2922_v11_
                nw456_[int(2)] = d_2922_v11_
                nw456_[int(3)] = d_2922_v11_
                nw456_[int(4)] = d_2922_v11_
                nw456_[int(5)] = d_2922_v11_
                nw456_[int(6)] = d_2922_v11_
                nw456_[int(7)] = (d_2922_v11_ if (self).f15 else d_2922_v11_)
                nw456_[int(8)] = d_2922_v11_
                nw456_[int(9)] = d_2922_v11_
                nw456_[int(10)] = d_2922_v11_
                nw456_[int(11)] = d_2922_v11_
                nw456_[int(12)] = d_2922_v11_
                nw456_[int(13)] = (d_2922_v11_ if (self).fm1((self).f12, False, p2, globalState) else d_2922_v11_)
                nw456_[int(14)] = d_2922_v11_
                nw456_[int(15)] = d_2922_v11_
                nw456_[int(16)] = d_2922_v11_
                nw456_[int(17)] = d_2922_v11_
                nw456_[int(18)] = d_2922_v11_
                nw456_[int(19)] = d_2922_v11_
                nw456_[int(20)] = d_2922_v11_
                nw456_[int(21)] = d_2922_v11_
                nw456_[int(22)] = d_2922_v11_
                d_2924_v12_ = nw456_
                index507_ = default__.safeIndex(243, (d_2924_v12_).length(0))
                (d_2924_v12_)[index507_] = d_2922_v11_
                index508_ = default__.safeIndex(243, (d_2924_v12_).length(0))
                (d_2924_v12_)[index508_] = d_2922_v11_
                r0 = p2
                rhs482_ = (d_2924_v12_)[default__.safeIndex(243, (d_2924_v12_).length(0))]
                rhs483_ = (self).f15
                lhs409_ = globalState
                d_2922_v11_ = rhs482_
                lhs409_.f7 = rhs483_
                d_2925_v13_: D1
                d_2925_v13_ = D1_DC5(p1, p2, p2)
                d_2926_v14_: _dafny.Map
                d_2926_v14_ = _dafny.Map({(self).f15: (self).f15})
                (globalState).f2 = (self).fm1((self).f11, (d_2925_v13_).cf7, (not(p2) if ((d_2926_v14_)[p2] if (p2) in (d_2926_v14_) else True) else (self).f15), globalState)
            elif True:
                d_2927_v15_: _dafny.Seq
                d_2927_v15_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                d_2927_v15_ = d_2927_v15_
                r1 = p3
                d_2928_v16_: _dafny.Map
                d_2928_v16_ = _dafny.Map({(self).f11: (self).f14})
                d_2929_v17_: _dafny.Map
                d_2929_v17_ = _dafny.Map({(len(d_2927_v15_)) != ((self).f14): (d_2928_v16_) | (d_2928_v16_)})
                d_2930_v19_: _dafny.Seq
                def iife185_():
                    coll71_ = _dafny.Map()
                    compr_71_: int
                    for compr_71_ in _dafny.IntegerRange(60, 653):
                        d_2931_v18_: int = compr_71_
                        if ((60) <= (d_2931_v18_)) and ((d_2931_v18_) < (653)):
                            coll71_[(d_2931_v18_) + ((self).f14)] = (self).f12
                    return _dafny.Map(coll71_)
                d_2930_v19_ = _dafny.SeqWithoutIsStrInference([iife185_()
                ])
                d_2932_v20_: _dafny.Map
                d_2932_v20_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flhmxyvdb")): d_2929_v17_})
                d_2929_v17_ = ((_dafny.Map({p2: ((d_2930_v19_)[default__.safeIndex(len(d_2928_v16_), len(d_2930_v19_))]).set(p3, p3)})) | (d_2929_v17_)) | ((d_2929_v17_) | (((d_2932_v20_)[(self).f13] if ((self).f13) in (d_2932_v20_) else _dafny.Map({True: d_2928_v16_}))))
                (globalState).f2 = (self).f15
                d_2933_v21_: C5
                nw457_ = C5()
                nw457_.ctor__()
                d_2933_v21_ = nw457_
            (globalState).f6 = p1
            (globalState).f6 = default__.fm0((self).f15, globalState)
            if p2:
                d_2934_v22_: _dafny.Set
                d_2934_v22_ = _dafny.Set({p1, p3, (self).fm2((self).f11, False, (self).f14, (self).f15, globalState)})
                def iife186_():
                    coll72_ = _dafny.Set()
                    compr_72_: int
                    for compr_72_ in _dafny.IntegerRange(750, -196):
                        d_2935_v23_: int = compr_72_
                        if ((750) <= (d_2935_v23_)) and ((d_2935_v23_) < (-196)):
                            coll72_ = coll72_.union(_dafny.Set([(d_2935_v23_) + ((self).f12)]))
                    return _dafny.Set(coll72_)
                d_2934_v22_ = (iife186_()
                ) | (d_2934_v22_)
                d_2936_v24_: _dafny.Map
                d_2936_v24_ = _dafny.Map({(self).f14: p1})
                d_2937_v25_: _dafny.Set
                d_2937_v25_ = _dafny.Set({(self).f15})
                d_2938_v26_: D15
                d_2938_v26_ = D15_DC43((self).f11, p2, p2)
                d_2939_v27_: _dafny.Map
                d_2939_v27_ = _dafny.Map({not((self).f15): (self).f15})
                d_2940_v28_: _dafny.Seq
                d_2940_v28_ = _dafny.SeqWithoutIsStrInference([(self).f15, False, p2])
                d_2941_v29_: _dafny.Array
                nw458_ = _dafny.Array(None, 19)
                nw458_[int(0)] = len(d_2936_v24_)
                nw458_[int(1)] = ((self).f11 if not((self).fm1(p1, (self).f15, p2, globalState)) else len(d_2937_v25_))
                nw458_[int(2)] = ((self).f11) + ((self).f12)
                nw458_[int(3)] = (self).f11
                nw458_[int(4)] = 238
                nw458_[int(5)] = p3
                nw458_[int(6)] = p1
                nw458_[int(7)] = (d_2938_v26_).cf77
                nw458_[int(8)] = 391
                nw458_[int(9)] = len(d_2939_v27_)
                nw458_[int(10)] = len(d_2940_v28_)
                nw458_[int(11)] = (self).f12
                nw458_[int(12)] = len(default__.fm48(globalState))
                nw458_[int(13)] = p3
                nw458_[int(14)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubupnkhmv")))
                nw458_[int(15)] = len(default__.fm42(p3, (self).f15, (self).fm12(d_2934_v22_, p0, (self).f15, len((self).f13), globalState), globalState))
                nw458_[int(16)] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f11])])), default__.fm0(p2, globalState))
                nw458_[int(17)] = p1
                nw458_[int(18)] = ((self).f11) + ((self).f12)
                d_2941_v29_ = nw458_
                index509_ = default__.safeIndex(760, (d_2941_v29_).length(0))
                (d_2941_v29_)[index509_] = p1
                d_2942_v30_: _dafny.Array
                nw459_ = _dafny.Array(None, 5)
                nw459_[int(0)] = p2
                nw459_[int(1)] = not((self).f15)
                nw459_[int(2)] = p2
                nw459_[int(3)] = (self).f15
                nw459_[int(4)] = (self).f15
                d_2942_v30_ = nw459_
                d_2943_v31_: _dafny.Seq
                d_2943_v31_ = _dafny.SeqWithoutIsStrInference([d_2942_v30_, d_2942_v30_])
                index510_ = default__.safeIndex(760, (d_2941_v29_).length(0))
                rhs484_ = ((0) - ((0) - (len(d_2943_v31_)))) - (len(default__.fm60(globalState)))
                rhs485_ = (self).fm1(len((d_2934_v22_ if True else d_2934_v22_)), True, (self).f15, globalState)
                rhs486_ = p1
                rhs487_ = not(p2)
                lhs410_ = d_2941_v29_
                lhs411_ = default__.safeIndex(760, (d_2941_v29_).length(0))
                lhs412_ = globalState
                lhs413_ = globalState
                lhs410_[lhs411_] = rhs484_
                lhs412_.f7 = rhs485_
                r1 = rhs486_
                lhs413_.f7 = rhs487_
                d_2944_v32_: _dafny.Seq
                d_2944_v32_ = _dafny.SeqWithoutIsStrInference([p0, default__.fm18(p2, -917, globalState), (self).f13])
                d_2945_v33_: _dafny.MultiSet
                d_2945_v33_ = _dafny.MultiSet([p1])
                d_2946_v34_: C1
                nw460_ = C1()
                nw460_.ctor__(d_2942_v30_, len(p0), len((d_2944_v32_)[default__.safeIndex((d_2945_v33_).cardinality, len(d_2944_v32_))]), (d_2945_v33_).cardinality, (self).f15, p0)
                d_2946_v34_ = nw460_
                d_2947_v35_: _dafny.Map
                d_2947_v35_ = _dafny.Map({d_2946_v34_: (self).f15})
                d_2947_v35_ = (d_2947_v35_).set(d_2946_v34_, ((self).f15) or (p2))
                r0 = (-667) not in (d_2945_v33_)
                d_2948_v36_: _dafny.Set
                d_2948_v36_ = _dafny.Set({d_2946_v34_.f24, d_2942_v30_})
                rhs488_ = (self).f15
                rhs489_ = (self).f15
                rhs490_ = not((127) == (p3))
                rhs491_ = (d_2948_v36_).issubset(d_2948_v36_)
                rhs492_ = (default__.fm23(len(d_2908_v0_), p3, globalState) if p2 else _dafny.SeqWithoutIsStrInference([p0]))
                lhs414_ = globalState
                lhs415_ = globalState
                r0 = rhs488_
                r0 = rhs489_
                lhs414_.f2 = rhs490_
                lhs415_.f2 = rhs491_
                d_2944_v32_ = rhs492_
            elif True:
                d_2949_v37_: _dafny.Array
                def lambda118_(d_2950_v10_):
                    def lambda119_(d_2951_i3_):
                        return (d_2951_i3_) * (len(d_2950_v10_))

                    return lambda119_

                init71_ = lambda118_(d_2921_v10_)
                nw461_ = _dafny.Array(None, 21)
                for i0_71_ in range(nw461_.length(0)):
                    nw461_[i0_71_] = init71_(i0_71_)
                d_2949_v37_ = nw461_
                index511_ = default__.safeIndex(396, (d_2949_v37_).length(0))
                (d_2949_v37_)[index511_] = p1
                index512_ = default__.safeIndex(396, (d_2949_v37_).length(0))
                rhs493_ = (self).f11
                rhs494_ = default__.safeDivisionInt(643, (_dafny.MultiSet([(self).f15])).cardinality)
                lhs416_ = d_2949_v37_
                lhs417_ = default__.safeIndex(396, (d_2949_v37_).length(0))
                r1 = rhs493_
                lhs416_[lhs417_] = rhs494_
                d_2952_v38_: str
                d_2952_v38_ = _dafny.CodePoint('o')
                d_2952_v38_ = d_2952_v38_
                d_2953_v39_: _dafny.Array
                nw462_ = _dafny.Array(_dafny.Seq({}), 21)
                d_2953_v39_ = nw462_
                d_2953_v39_ = d_2953_v39_
                d_2954_v40_: _dafny.Map
                d_2954_v40_ = _dafny.Map({p2: p2})
                d_2955_v41_: _dafny.Map
                d_2955_v41_ = _dafny.Map({(self).f15: d_2954_v40_})
                d_2956_v42_: D15
                d_2956_v42_ = D15_DC42(d_2955_v41_)
                d_2957_v43_: _dafny.Set
                d_2957_v43_ = _dafny.Set({d_2956_v42_, d_2956_v42_})
                d_2958_v44_: _dafny.Map
                d_2958_v44_ = _dafny.Map({p2: d_2957_v43_})
                d_2958_v44_ = (d_2958_v44_).set(p2, _dafny.Set({d_2956_v42_}))
                d_2959_v45_: _dafny.Map
                d_2959_v45_ = _dafny.Map({(self).f12: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahto"))})
                r1 = len(((d_2959_v45_) | (d_2959_v45_)).set((d_2949_v37_)[default__.safeIndex(396, (d_2949_v37_).length(0))], (p0) + (p0)))
        elif True:
            (globalState).f2 = p2
            d_2960_v46_: str
            d_2960_v46_ = _dafny.CodePoint('w')
            d_2961_v47_: _dafny.Seq
            d_2961_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_2960_v46_})])
            d_2962_v48_: _dafny.Array
            nw463_ = _dafny.Array(None, 11)
            nw463_[int(0)] = p2
            nw463_[int(1)] = (self).f15
            nw463_[int(2)] = (self).f15
            nw463_[int(3)] = (self).f15
            nw463_[int(4)] = False
            nw463_[int(5)] = p2
            nw463_[int(6)] = ((d_2961_v47_)[default__.safeIndex((self).f12, len(d_2961_v47_))]).ispropersubset(_dafny.Set({d_2960_v46_, d_2960_v46_}))
            nw463_[int(7)] = p2
            nw463_[int(8)] = (self).f15
            nw463_[int(9)] = p2
            nw463_[int(10)] = (self).f15
            d_2962_v48_ = nw463_
            index513_ = default__.safeIndex(350, (d_2962_v48_).length(0))
            (d_2962_v48_)[index513_] = p2
            d_2963_v49_: D16
            d_2963_v49_ = D16_DC45(d_2962_v48_, d_2960_v46_, (0) - (p1), 382, (self).f15)
            index514_ = default__.safeIndex(350, (d_2962_v48_).length(0))
            (d_2962_v48_)[index514_] = (d_2960_v46_) not in (((self).f13) + ((p0).set(default__.safeIndex(p1, len(p0)), (d_2963_v49_).cf82)))
            if (d_2962_v48_)[default__.safeIndex(350, (d_2962_v48_).length(0))]:
                d_2964_v50_: _dafny.Array
                nw464_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_2964_v50_ = nw464_
                nw465_ = _dafny.Array(None, 13)
                nw465_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rte"))
                nw465_[int(1)] = _dafny.SeqWithoutIsStrInference([d_2960_v46_ for d_2965_i4_ in range(default__.abs(294))])
                nw465_[int(2)] = _dafny.SeqWithoutIsStrInference([d_2960_v46_ for d_2966_i5_ in range(default__.abs(125))])
                nw465_[int(3)] = p0
                nw465_[int(4)] = (self).f13
                nw465_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2967_i6_ in range(default__.abs(-921))])
                nw465_[int(6)] = p0
                nw465_[int(7)] = (self).f13
                nw465_[int(8)] = ((self).f13) + ((self).f13)
                nw465_[int(9)] = p0
                nw465_[int(10)] = (self).f13
                nw465_[int(11)] = p0
                nw465_[int(12)] = ((self).f13) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhxtce")))
                d_2964_v50_ = nw465_
                d_2968_v51_: _dafny.Array
                nw466_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                d_2968_v51_ = nw466_
                d_2968_v51_ = d_2968_v51_
                d_2969_v52_: D3
                d_2969_v52_ = D3_DC10(p2, (self).f12, _dafny.Map({(self).f13: (self).f12}), p0, (self).f11)
                d_2970_v53_: _dafny.Map
                d_2970_v53_ = _dafny.Map({d_2969_v52_: (self).f14})
                d_2971_v54_: bool
                d_2972_v55_: _dafny.Seq
                out85_: bool
                out86_: _dafny.Seq
                out85_, out86_ = default__.m0(((0) - ((self).f11)) < (p1), (d_2962_v48_)[default__.safeIndex(350, (d_2962_v48_).length(0))], (self).f12, ((self).f14) * (((d_2970_v53_)[d_2969_v52_] if (d_2969_v52_) in (d_2970_v53_) else (self).f14)), globalState)
                d_2971_v54_ = out85_
                d_2972_v55_ = out86_
                r1 = (d_2963_v49_).cf83
                d_2973_v56_: _dafny.Set
                d_2973_v56_ = _dafny.Set({d_2962_v48_})
                d_2974_v57_: _dafny.Set
                d_2974_v57_ = _dafny.Set({(self).f15})
                d_2975_v58_: C2
                nw467_ = C2()
                nw467_.ctor__(len(d_2973_v56_), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgmsv"))).set(default__.safeIndex((self).f14, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgmsv")))), d_2960_v46_), (p1 if p2 else p1), (0) - (((self).f12) * ((0) - ((self).fm2(p1, d_2971_v54_, len(d_2974_v57_), d_2971_v54_, globalState)))))
                d_2975_v58_ = nw467_
                rhs495_ = d_2975_v58_
                rhs496_ = (self).f15
                rhs497_ = True
                lhs418_ = globalState
                lhs419_ = globalState
                d_2975_v58_ = rhs495_
                lhs418_.f7 = rhs496_
                lhs419_.f7 = rhs497_
            elif True:
                d_2976_v59_: bool
                d_2977_v60_: _dafny.Seq
                out87_: bool
                out88_: _dafny.Seq
                out87_, out88_ = default__.m0((d_2962_v48_)[default__.safeIndex(350, (d_2962_v48_).length(0))], False, ((self).f14) - ((self).f11), (self).f11, globalState)
                d_2976_v59_ = out87_
                d_2977_v60_ = out88_
                d_2978_v61_: _dafny.Array
                nw468_ = _dafny.Array(int(0), 17)
                d_2978_v61_ = nw468_
                d_2979_v62_: _dafny.MultiSet
                d_2979_v62_ = _dafny.MultiSet([p1, len(p0), p3, (self).f14, (self).f12])
                index515_ = default__.safeIndex(456, (d_2978_v61_).length(0))
                (d_2978_v61_)[index515_] = (d_2979_v62_).cardinality
                index516_ = default__.safeIndex(456, (d_2978_v61_).length(0))
                (d_2978_v61_)[index516_] = len(default__.fm22(488, globalState))
                (globalState).f2 = d_2976_v59_
                index517_ = default__.safeIndex(456, (d_2978_v61_).length(0))
                (d_2978_v61_)[index517_] = len(_dafny.Set({(self).f15}))
                d_2980_v63_: C4
                nw469_ = C4()
                nw469_.ctor__((self).f13, (d_2978_v61_)[default__.safeIndex(456, (d_2978_v61_).length(0))], not(p2), p0, p3, (self).f11)
                d_2980_v63_ = nw469_
                index518_ = default__.safeIndex(456, (d_2978_v61_).length(0))
                index519_ = default__.safeIndex(456, (d_2978_v61_).length(0))
                rhs498_ = p2
                rhs499_ = (d_2962_v48_)[default__.safeIndex(350, (d_2962_v48_).length(0))]
                rhs500_ = ((self).f12 if not(d_2976_v59_) else len((D21_DC60(_dafny.SeqWithoutIsStrInference([d_2980_v63_, d_2980_v63_, d_2980_v63_, d_2980_v63_, d_2980_v63_]))).cf112))
                rhs501_ = d_2976_v59_
                rhs502_ = (d_2978_v61_)[default__.safeIndex(456, (d_2978_v61_).length(0))]
                lhs420_ = globalState
                lhs421_ = d_2978_v61_
                lhs422_ = default__.safeIndex(456, (d_2978_v61_).length(0))
                lhs423_ = globalState
                lhs424_ = d_2978_v61_
                lhs425_ = default__.safeIndex(456, (d_2978_v61_).length(0))
                r0 = rhs498_
                lhs420_.f2 = rhs499_
                lhs421_[lhs422_] = rhs500_
                lhs423_.f7 = rhs501_
                lhs424_[lhs425_] = rhs502_
            (globalState).f2 = ((self).f15) or (p2)
            d_2981_v64_: C0
            nw470_ = C0()
            nw470_.ctor__()
            d_2981_v64_ = nw470_
        d_2982_v65_: str
        d_2982_v65_ = _dafny.CodePoint('h')
        d_2982_v65_ = d_2982_v65_
        d_2983_v66_: _dafny.Array
        nw471_ = _dafny.Array(_dafny.MultiSet({}), 25)
        d_2983_v66_ = nw471_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_2983_v66_).length(0)):
            d_2984_i7_: int = guard_loop_8_
            if (True) and (((0) <= (d_2984_i7_)) and ((d_2984_i7_) < ((d_2983_v66_).length(0)))):
                (d_2983_v66_)[(d_2984_i7_)] = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p2])) if (self).f15 else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, (self).f15])))) - ((_dafny.MultiSet([not(p2), (self).f15])) - (_dafny.MultiSet([False])))
        d_2985_v67_: _dafny.Map
        d_2985_v67_ = _dafny.Map({625: p3})
        d_2986_v68_: _dafny.Set
        d_2986_v68_ = _dafny.Set({(self).fm1(((d_2985_v67_)[p3] if (p3) in (d_2985_v67_) else (0) - ((0) - (p1))), (self).f15, p2, globalState)})
        if (p3) != (len((d_2986_v68_) - (d_2986_v68_))):
            d_2987_v69_: _dafny.MultiSet
            d_2987_v69_ = _dafny.MultiSet([(self).f15])
            index520_ = default__.safeIndex(318, (d_2983_v66_).length(0))
            (d_2983_v66_)[index520_] = d_2987_v69_
            index521_ = default__.safeIndex(318, (d_2983_v66_).length(0))
            (d_2983_v66_)[index521_] = (d_2987_v69_).intersection((d_2987_v69_).intersection(d_2987_v69_))
            (globalState).f6 = (self).f11
            d_2988_v70_: _dafny.Map
            d_2988_v70_ = _dafny.Map({(self).f15: (self).fm1(p1, p2, False, globalState)})
            d_2989_v71_: _dafny.Seq
            d_2989_v71_ = _dafny.SeqWithoutIsStrInference([p2])
            d_2988_v70_ = (d_2988_v70_).set((d_2989_v71_)[default__.safeIndex((self).f12, len(d_2989_v71_))], (self).f15)
            r0 = p2
            d_2990_v72_: C2
            nw472_ = C2()
            nw472_.ctor__(808, (p0) + (p0), (814) * ((self).f12), (self).f12)
            d_2990_v72_ = nw472_
        elif True:
            r2 = default__.safeDivisionInt(p1, p3)
            if (self).f15:
                (globalState).f6 = (d_2908_v0_)[default__.safeIndex((self).f14, len(d_2908_v0_))]
                d_2991_v73_: _dafny.Map
                d_2991_v73_ = _dafny.Map({(self).f14: _dafny.SeqWithoutIsStrInference([(self).f13])})
                d_2992_v74_: _dafny.Seq
                d_2992_v74_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntli"))])
                d_2991_v73_ = (d_2991_v73_).set(default__.safeModuloInt((self).f12, (d_2908_v0_)[default__.safeIndex(p1, len(d_2908_v0_))]), d_2992_v74_)
                d_2993_v75_: _dafny.Array
                def lambda120_(d_2994_p1_):
                    def lambda121_(d_2995_i8_):
                        return (d_2995_i8_) * (d_2994_p1_)

                    return lambda121_

                init72_ = lambda120_(p1)
                nw473_ = _dafny.Array(None, 3)
                for i0_72_ in range(nw473_.length(0)):
                    nw473_[i0_72_] = init72_(i0_72_)
                d_2993_v75_ = nw473_
                index522_ = default__.safeIndex(922, (d_2993_v75_).length(0))
                (d_2993_v75_)[index522_] = len(d_2985_v67_)
                d_2996_v76_: D12
                d_2996_v76_ = D12_DC35(d_2986_v68_)
                d_2997_v77_: _dafny.Array
                nw474_ = _dafny.Array(False, 28)
                d_2997_v77_ = nw474_
                d_2998_v78_: _dafny.Map
                d_2998_v78_ = _dafny.Map({d_2997_v77_: p3})
                d_2999_v79_: _dafny.MultiSet
                d_2999_v79_ = _dafny.MultiSet([(self).f15, (self).f15])
                d_3000_v80_: D5
                d_3000_v80_ = D5_DC18(d_2998_v78_, ((d_2999_v79_)[(self).f15] if ((self).f15) in (d_2999_v79_) else (self).f11))
                d_3001_v81_: _dafny.Map
                d_3001_v81_ = _dafny.Map({d_2996_v76_: d_3000_v80_})
                d_3002_v82_: _dafny.Seq
                d_3002_v82_ = _dafny.SeqWithoutIsStrInference([d_3001_v81_, d_3001_v81_])
                index523_ = default__.safeIndex(922, (d_2993_v75_).length(0))
                (d_2993_v75_)[index523_] = len(d_3002_v82_)
                r0 = p2
                d_3003_v83_: _dafny.Map
                d_3003_v83_ = _dafny.Map({p2: False})
                d_3004_v84_: _dafny.Map
                d_3004_v84_ = _dafny.Map({d_2982_v65_: ((d_3003_v83_)[p2] if (p2) in (d_3003_v83_) else False)})
                d_3005_v85_: _dafny.Seq
                d_3005_v85_ = _dafny.SeqWithoutIsStrInference([p2, (self).f15])
                d_3006_v86_: _dafny.Map
                d_3006_v86_ = _dafny.Map({((d_3004_v84_)[_dafny.CodePoint('d')] if (_dafny.CodePoint('d')) in (d_3004_v84_) else (d_3005_v85_)[default__.safeIndex((self).f11, len(d_3005_v85_))]): (self).f14})
                rhs503_ = (self).f15
                rhs504_ = (self).f14
                rhs505_ = ((d_3006_v86_)[(self).f15] if ((self).f15) in (d_3006_v86_) else p3)
                lhs426_ = globalState
                lhs427_ = globalState
                lhs426_.f2 = rhs503_
                lhs427_.f6 = rhs504_
                r2 = rhs505_
            elif True:
                (globalState).f6 = (len(p0)) - ((self).f11)
                d_3007_v87_: _dafny.Map
                d_3007_v87_ = _dafny.Map({p2: -825})
                d_3008_v88_: _dafny.Map
                d_3008_v88_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(D6_DC19(d_3007_v87_)).cf37]): p0})
                d_3009_v89_: _dafny.Seq
                d_3009_v89_ = _dafny.SeqWithoutIsStrInference([d_3007_v87_])
                d_3008_v88_ = (d_3008_v88_).set(d_3009_v89_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dysfagln")) if p2 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epqewq"))))
                r1 = (self).f14
                d_3010_v90_: _dafny.Array
                nw475_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_3010_v90_ = nw475_
                d_3010_v90_ = d_3010_v90_
                d_3011_v91_: _dafny.Array
                nw476_ = _dafny.Array(int(0), 24)
                d_3011_v91_ = nw476_
                index524_ = default__.safeIndex(953, (d_3011_v91_).length(0))
                (d_3011_v91_)[index524_] = (self).f11
                index525_ = default__.safeIndex(953, (d_3011_v91_).length(0))
                (d_3011_v91_)[index525_] = len(p0)
            (globalState).f7 = p2
            d_3012_v92_: _dafny.Seq
            d_3012_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ya"))
            d_3012_v92_ = p0
            d_3013_v93_: _dafny.MultiSet
            d_3013_v93_ = _dafny.MultiSet([p1])
            (globalState).f2 = not((d_3013_v93_).issubset(d_3013_v93_))
        d_3014_v94_: D10
        d_3014_v94_ = D10_DC30(p2, False, (self).f13, (self).f15)
        pat_let_tv50_ = p2
        pat_let_tv51_ = p2
        def lambda122_(source42_):
            if source42_.is_DC29:
                return pat_let_tv50_
            elif source42_.is_DC30:
                d_3015___mcc_h0_ = source42_.cf48
                d_3016___mcc_h1_ = source42_.cf49
                d_3017___mcc_h2_ = source42_.cf50
                d_3018___mcc_h3_ = source42_.cf51
                d_3019_cf51_ = d_3018___mcc_h3_
                d_3020_cf50_ = d_3017___mcc_h2_
                d_3021_cf49_ = d_3016___mcc_h1_
                d_3022_cf48_ = d_3015___mcc_h0_
                return not (pat_let_tv51_) or ((self).f15)
            elif source42_.is_DC31:
                d_3023___mcc_h4_ = source42_.cf52
                d_3024___mcc_h5_ = source42_.cf53
                d_3025___mcc_h6_ = source42_.cf54
                d_3026___mcc_h7_ = source42_.cf55
                d_3027_cf55_ = d_3026___mcc_h7_
                d_3028_cf54_ = d_3025___mcc_h6_
                d_3029_cf53_ = d_3024___mcc_h5_
                d_3030_cf52_ = d_3023___mcc_h4_
                return False
            elif source42_.is_DC28:
                d_3031___mcc_h8_ = source42_.cf47
                d_3032_cf47_ = d_3031___mcc_h8_
                return ((self).f12) <= ((self).f14)
            elif True:
                d_3033___mcc_h9_ = source42_.cf56
                d_3034_cf56_ = d_3033___mcc_h9_
                return (_dafny.MultiSet([not((self).f15), (self).f15])).ispropersubset(_dafny.MultiSet([False]))

        (globalState).f2 = lambda122_(d_3014_v94_)
        r0 = p2
        r1 = (self).f11
        r2 = (self).f14
        return r0, r1, r2

    @property
    def f23(self):
        return self._f23

class C7:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def fm10(self, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncex"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flywa")))

    def m9(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_3035_v0_: _dafny.Array
        def lambda123_(d_3036_i0_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apfdd"))).set(default__.safeIndex(694, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apfdd")))), _dafny.CodePoint('c'))

        init73_ = lambda123_
        nw477_ = _dafny.Array(None, 10)
        for i0_73_ in range(nw477_.length(0)):
            nw477_[i0_73_] = init73_(i0_73_)
        d_3035_v0_ = nw477_
        index526_ = default__.safeIndex(293, (d_3035_v0_).length(0))
        (d_3035_v0_)[index526_] = p1
        d_3037_v1_: bool
        d_3037_v1_ = True
        index527_ = default__.safeIndex(293, (d_3035_v0_).length(0))
        (d_3035_v0_)[index527_] = (p1).set(default__.safeIndex((p0).f12, len(p1)), default__.fm11(d_3037_v1_, (0) - ((0) - ((p0).f11)), globalState))
        d_3038_v2_: _dafny.MultiSet
        d_3038_v2_ = _dafny.MultiSet([_dafny.CodePoint('y')])
        d_3039_v3_: _dafny.Array
        nw478_ = _dafny.Array(None, 14)
        nw478_[int(0)] = default__.fm69(not(d_3037_v1_), (0) - ((p0).f11), d_3038_v2_, (p0).f12, globalState)
        nw478_[int(1)] = not(d_3037_v1_)
        nw478_[int(2)] = d_3037_v1_
        nw478_[int(3)] = d_3037_v1_
        nw478_[int(4)] = d_3037_v1_
        nw478_[int(5)] = d_3037_v1_
        nw478_[int(6)] = d_3037_v1_
        nw478_[int(7)] = d_3037_v1_
        nw478_[int(8)] = d_3037_v1_
        nw478_[int(9)] = d_3037_v1_
        nw478_[int(10)] = not(False)
        nw478_[int(11)] = d_3037_v1_
        nw478_[int(12)] = d_3037_v1_
        nw478_[int(13)] = d_3037_v1_
        d_3039_v3_ = nw478_
        d_3040_v4_: C1
        nw479_ = C1()
        nw479_.ctor__(d_3039_v3_, ((p0).f11) * (350), (p0).f11, default__.safeDivisionInt((p0).f11, (p0).f11), d_3037_v1_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))) + (p1))
        d_3040_v4_ = nw479_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_3039_v3_).length(0)):
            d_3041_i1_: int = guard_loop_9_
            if (True) and (((0) <= (d_3041_i1_)) and ((d_3041_i1_) < ((d_3039_v3_).length(0)))):
                (d_3039_v3_)[(d_3041_i1_)] = d_3037_v1_
        d_3037_v1_ = d_3037_v1_
        d_3042_v5_: _dafny.Seq
        d_3042_v5_ = _dafny.SeqWithoutIsStrInference([d_3037_v1_])
        d_3042_v5_ = d_3042_v5_
        r2 = not(((p0).f12) < ((p0).f12))
        r0 = (p0).f11
        d_3043_v6_: _dafny.Map
        d_3043_v6_ = _dafny.Map({d_3040_v4_.f24: d_3037_v1_})
        d_3044_v7_: _dafny.Seq
        d_3044_v7_ = _dafny.SeqWithoutIsStrInference([len(d_3043_v6_)])
        r1 = len((d_3044_v7_ if d_3037_v1_ else _dafny.SeqWithoutIsStrInference([default__.fm0(d_3037_v1_, globalState)])))
        r2 = d_3037_v1_
        return r0, r1, r2


class C8:
    def  __init__(self):
        self.f21: bool = False
        self._f22: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self, f21, f22):
        (self).f21 = f21
        (self)._f22 = f22

    def fm6(self, p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([311])) | (_dafny.MultiSet([412]))) | ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([218]))) | (_dafny.MultiSet([-376, -491])))

    def m7(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: T1 = None
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        if not (self.f21) or (((self).f22 if (self).f22 else default__.fm7(self.f21, p0, globalState))):
            d_3045_v0_: _dafny.Seq
            d_3045_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksw"))
            d_3046_v1_: _dafny.Seq
            d_3046_v1_ = _dafny.SeqWithoutIsStrInference([default__.fm7(not((self).f22), len(d_3045_v0_), globalState)])
            (globalState).f7 = not (not((p0) >= (p0))) or ((d_3046_v1_)[default__.safeIndex(p0, len(d_3046_v1_))])
            (globalState).f6 = p0
            if self.f21:
                (globalState).f2 = (self).f22
                (globalState).f6 = (0) - (default__.safeModuloInt(-450, p0))
                d_3047_v2_: D0
                d_3047_v2_ = D0_DC1((self).f22, p0)
                d_3048_v3_: _dafny.MultiSet
                d_3048_v3_ = _dafny.MultiSet([self.f21, self.f21, self.f21, p2, p2])
                d_3049_v4_: _dafny.Map
                d_3049_v4_ = _dafny.Map({p0: d_3046_v1_})
                d_3050_v5_: _dafny.Set
                d_3050_v5_ = _dafny.Set({len(d_3045_v0_)})
                d_3051_v6_: D3
                d_3051_v6_ = D3_DC10(self.f21, p0, default__.fm9((self).f22, globalState), d_3045_v0_, len(d_3050_v5_))
                d_3052_v7_: _dafny.Map
                d_3052_v7_ = _dafny.Map({d_3051_v6_: d_3046_v1_})
                rhs506_ = (0) - (p0)
                rhs507_ = (p0) - (p0)
                rhs508_ = (d_3046_v1_ if self.f21 else (_dafny.SeqWithoutIsStrInference([self.f21, self.f21])) + (default__.fm8((self).f22, self.f21, (d_3047_v2_).cf1, p0, globalState)))
                rhs509_ = ((d_3048_v3_).intersection(_dafny.MultiSet(((d_3049_v4_)[p0] if (p0) in (d_3049_v4_) else ((d_3052_v7_)[d_3051_v6_] if (d_3051_v6_) in (d_3052_v7_) else d_3046_v1_))))) - (_dafny.MultiSet([False, True, True]))
                rhs510_ = (self).f22
                lhs428_ = globalState
                lhs429_ = globalState
                lhs430_ = self
                lhs428_.f6 = rhs506_
                lhs429_.f6 = rhs507_
                d_3046_v1_ = rhs508_
                r2 = rhs509_
                lhs430_.f21 = rhs510_
                d_3053_v8_: _dafny.Array
                nw480_ = _dafny.Array(None, 10)
                nw480_[int(0)] = not(self.f21)
                nw480_[int(1)] = False
                nw480_[int(2)] = p2
                nw480_[int(3)] = (self).f22
                nw480_[int(4)] = (self).f22
                nw480_[int(5)] = True
                nw480_[int(6)] = (self).f22
                nw480_[int(7)] = self.f21
                nw480_[int(8)] = (self).f22
                nw480_[int(9)] = not((self).f22)
                d_3053_v8_ = nw480_
                d_3054_v9_: _dafny.Map
                d_3054_v9_ = _dafny.Map({d_3053_v8_: p2})
                d_3055_v10_: _dafny.Seq
                d_3055_v10_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_3054_v9_))])
                d_3056_v11_: _dafny.Map
                d_3056_v11_ = _dafny.Map({d_3055_v10_: d_3045_v0_})
                (globalState).f7 = not(default__.fm7((d_3055_v10_) in (d_3056_v11_), (p1).cf24, globalState))
                d_3057_v12_: _dafny.MultiSet
                d_3057_v12_ = _dafny.MultiSet([p0])
                (globalState).f2 = (p0) != (((d_3057_v12_).cardinality) * (p0))
            elif True:
                d_3058_v13_: _dafny.Array
                nw481_ = _dafny.Array(_dafny.Seq({}), 15)
                d_3058_v13_ = nw481_
                d_3059_v14_: _dafny.Array
                nw482_ = _dafny.Array(int(0), 11)
                d_3059_v14_ = nw482_
                d_3060_v15_: _dafny.Seq
                d_3060_v15_ = _dafny.SeqWithoutIsStrInference([d_3059_v14_])
                index528_ = default__.safeIndex(398, (d_3058_v13_).length(0))
                (d_3058_v13_)[index528_] = d_3060_v15_
                index529_ = default__.safeIndex(398, (d_3058_v13_).length(0))
                (d_3058_v13_)[index529_] = d_3060_v15_
                d_3061_v16_: C0
                nw483_ = C0()
                nw483_.ctor__()
                d_3061_v16_ = nw483_
                d_3062_v17_: C7
                nw484_ = C7()
                nw484_.ctor__()
                d_3062_v17_ = nw484_
                r0 = p0
                d_3063_v18_: _dafny.Array
                nw485_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                d_3063_v18_ = nw485_
                index530_ = default__.safeIndex(720, (d_3063_v18_).length(0))
                (d_3063_v18_)[index530_] = default__.fm34(_dafny.CodePoint('t'), globalState)
                d_3064_v19_: str
                d_3064_v19_ = _dafny.CodePoint('u')
                index531_ = default__.safeIndex(720, (d_3063_v18_).length(0))
                (d_3063_v18_)[index531_] = d_3064_v19_
            d_3065_v20_: _dafny.Array
            nw486_ = _dafny.Array(False, 18)
            d_3065_v20_ = nw486_
            d_3066_v21_: _dafny.Set
            d_3066_v21_ = _dafny.Set({d_3065_v20_})
            d_3066_v21_ = d_3066_v21_
            (globalState).f6 = default__.safeDivisionInt(p0, (p0) * (p0))
        elif True:
            d_3067_v22_: str
            d_3067_v22_ = _dafny.CodePoint('l')
            d_3068_v23_: _dafny.MultiSet
            d_3068_v23_ = _dafny.MultiSet([d_3067_v22_, d_3067_v22_, d_3067_v22_])
            (globalState).f7 = not(default__.fm69(self.f21, p0, (d_3068_v23_ if p2 else d_3068_v23_), (p0) + (p0), globalState))
            if ((p0 if (self).f22 else p0)) < (default__.fm0(self.f21, globalState)):
                d_3069_v25_: _dafny.Map
                d_3069_v25_ = _dafny.Map({243: p2})
                d_3070_v26_: _dafny.Seq
                def iife187_():
                    coll73_ = _dafny.Map()
                    compr_73_: int
                    for compr_73_ in _dafny.IntegerRange(326, -130):
                        d_3071_v24_: int = compr_73_
                        if ((326) <= (d_3071_v24_)) and ((d_3071_v24_) < (-130)):
                            coll73_[(d_3071_v24_) - (851)] = self.f21
                    return _dafny.Map(coll73_)
                d_3070_v26_ = _dafny.SeqWithoutIsStrInference([iife187_()
                , (d_3069_v25_).set(default__.fm0((self).f22, globalState), False)])
                d_3072_v27_: D22
                d_3072_v27_ = D22_DC62(d_3069_v25_)
                d_3073_v28_: _dafny.Seq
                d_3073_v28_ = _dafny.SeqWithoutIsStrInference([(d_3070_v26_)[default__.safeIndex(p0, len(d_3070_v26_))], d_3069_v25_, (d_3072_v27_).cf117])
                d_3074_v29_: _dafny.Seq
                d_3074_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sijb"))
                d_3075_v30_: _dafny.Seq
                d_3075_v30_ = _dafny.SeqWithoutIsStrInference([True])
                d_3076_v31_: _dafny.Seq
                d_3076_v31_ = _dafny.SeqWithoutIsStrInference([408, len(d_3074_v29_)])
                d_3077_v32_: D15
                d_3077_v32_ = D15_DC43(p0, (self).f22, p2)
                d_3078_v33_: _dafny.Map
                d_3078_v33_ = _dafny.Map({p2: p0})
                d_3079_v34_: C6
                nw487_ = C6()
                nw487_.ctor__((d_3070_v26_) + (_dafny.SeqWithoutIsStrInference([d_3069_v25_ for d_3080_i0_ in range(default__.abs(563))])), p0, p2, d_3074_v29_, (len(d_3075_v30_)) - (len(d_3076_v31_)), ((d_3077_v32_).cf77) * (((d_3078_v33_)[(self).f22] if ((self).f22) in (d_3078_v33_) else p0)))
                d_3079_v34_ = nw487_
                d_3081_v36_: _dafny.Array
                def lambda124_(d_3082_p0_, d_3083_v22_):
                    def lambda125_(d_3084_i1_):
                        def iife188_():
                            coll74_ = _dafny.Map()
                            compr_74_: int
                            for compr_74_ in _dafny.IntegerRange(288, 187):
                                d_3085_v35_: int = compr_74_
                                if ((288) <= (d_3085_v35_)) and ((d_3085_v35_) < (187)):
                                    coll74_[(d_3085_v35_) - (d_3082_p0_)] = d_3083_v22_
                            return _dafny.Map(coll74_)
                        return (d_3084_i1_) * (len(iife188_()
                        ))

                    return lambda125_

                init74_ = lambda124_(p0, d_3067_v22_)
                nw488_ = _dafny.Array(None, 15)
                for i0_74_ in range(nw488_.length(0)):
                    nw488_[i0_74_] = init74_(i0_74_)
                d_3081_v36_ = nw488_
                index532_ = default__.safeIndex(416, (d_3081_v36_).length(0))
                (d_3081_v36_)[index532_] = p0
                index533_ = default__.safeIndex(416, (d_3081_v36_).length(0))
                (d_3081_v36_)[index533_] = p0
                d_3086_v37_: _dafny.Map
                d_3086_v37_ = _dafny.Map({(self).f22: self.f21})
                d_3087_v38_: _dafny.Array
                nw489_ = _dafny.Array(None, 18)
                nw489_[int(0)] = self.f21
                nw489_[int(1)] = ((d_3081_v36_)[default__.safeIndex(416, (d_3081_v36_).length(0))]) != ((d_3081_v36_)[default__.safeIndex(416, (d_3081_v36_).length(0))])
                nw489_[int(2)] = p2
                nw489_[int(3)] = p2
                nw489_[int(4)] = self.f21
                nw489_[int(5)] = self.f21
                nw489_[int(6)] = p2
                nw489_[int(7)] = ((self).f22) or (self.f21)
                nw489_[int(8)] = (self).f22
                nw489_[int(9)] = self.f21
                nw489_[int(10)] = p2
                nw489_[int(11)] = p2
                nw489_[int(12)] = (self).f22
                nw489_[int(13)] = self.f21
                nw489_[int(14)] = (self).f22
                nw489_[int(15)] = self.f21
                nw489_[int(16)] = True
                nw489_[int(17)] = ((d_3086_v37_)[True] if (True) in (d_3086_v37_) else p2)
                d_3087_v38_ = nw489_
                index534_ = default__.safeIndex(316, (d_3087_v38_).length(0))
                (d_3087_v38_)[index534_] = (self).f22
                d_3088_v39_: _dafny.MultiSet
                d_3088_v39_ = _dafny.MultiSet([(self).f22])
                index535_ = default__.safeIndex(316, (d_3087_v38_).length(0))
                (d_3087_v38_)[index535_] = (d_3088_v39_).issubset(_dafny.MultiSet([True, not(p2)]))
                d_3089_v40_: _dafny.Seq
                d_3089_v40_ = _dafny.SeqWithoutIsStrInference([d_3081_v36_, d_3081_v36_, d_3081_v36_, d_3081_v36_])
                d_3089_v40_ = d_3089_v40_
                (globalState).f6 = (0) - ((default__.safeModuloInt((d_3081_v36_)[default__.safeIndex(416, (d_3081_v36_).length(0))], p0)) * (default__.safeDivisionInt(len(_dafny.Set({not(self.f21)})), (d_3079_v34_).fm2((d_3081_v36_)[default__.safeIndex(416, (d_3081_v36_).length(0))], (self).f22, (d_3081_v36_)[default__.safeIndex(416, (d_3081_v36_).length(0))], self.f21, globalState))))
            elif True:
                d_3090_v41_: _dafny.Array
                def lambda126_(d_3091_p2_):
                    def lambda127_(d_3092_i2_):
                        return d_3091_p2_

                    return lambda127_

                init75_ = lambda126_(p2)
                nw490_ = _dafny.Array(None, 29)
                for i0_75_ in range(nw490_.length(0)):
                    nw490_[i0_75_] = init75_(i0_75_)
                d_3090_v41_ = nw490_
                d_3090_v41_ = d_3090_v41_
                d_3093_v42_: D2
                d_3093_v42_ = D2_DC6(_dafny.SeqWithoutIsStrInference([d_3067_v22_ for d_3094_i3_ in range(default__.abs(596))]))
                d_3093_v42_ = default__.fm70(globalState)
                d_3095_v43_: _dafny.Array
                nw491_ = _dafny.Array(int(0), 12)
                d_3095_v43_ = nw491_
                d_3096_v44_: _dafny.Seq
                d_3096_v44_ = _dafny.SeqWithoutIsStrInference([p0])
                index536_ = default__.safeIndex(941, (d_3095_v43_).length(0))
                (d_3095_v43_)[index536_] = len((d_3096_v44_ if self.f21 else d_3096_v44_))
                index537_ = default__.safeIndex(941, (d_3095_v43_).length(0))
                (d_3095_v43_)[index537_] = p0
                d_3097_v45_: _dafny.Array
                nw492_ = _dafny.Array(D3.default()(), 6)
                d_3097_v45_ = nw492_
                d_3098_v46_: _dafny.Seq
                d_3098_v46_ = _dafny.SeqWithoutIsStrInference([d_3097_v45_])
                d_3099_v47_: _dafny.MultiSet
                d_3099_v47_ = _dafny.MultiSet([(d_3095_v43_)[default__.safeIndex(941, (d_3095_v43_).length(0))]])
                rhs511_ = ((d_3098_v46_) != (_dafny.SeqWithoutIsStrInference([d_3097_v45_])) if (d_3096_v44_) <= (_dafny.SeqWithoutIsStrInference([p0])) else (((d_3099_v47_)[(d_3095_v43_)[default__.safeIndex(941, (d_3095_v43_).length(0))]] if ((d_3095_v43_)[default__.safeIndex(941, (d_3095_v43_).length(0))]) in (d_3099_v47_) else p0)) in (default__.fm29(globalState)))
                rhs512_ = self.f21
                rhs513_ = d_3067_v22_
                rhs514_ = (p2) or (not (p2) or (self.f21))
                lhs431_ = globalState
                lhs432_ = globalState
                lhs433_ = globalState
                lhs431_.f2 = rhs511_
                lhs432_.f7 = rhs512_
                d_3067_v22_ = rhs513_
                lhs433_.f2 = rhs514_
                index538_ = default__.safeIndex(587, (d_3090_v41_).length(0))
                (d_3090_v41_)[index538_] = not((self).f22)
                d_3100_v48_: _dafny.MultiSet
                d_3100_v48_ = _dafny.MultiSet([p2, not(not((self).f22))])
                index539_ = default__.safeIndex(587, (d_3090_v41_).length(0))
                rhs515_ = ((d_3095_v43_)[default__.safeIndex(941, (d_3095_v43_).length(0))]) >= ((d_3100_v48_).cardinality)
                rhs516_ = p0
                rhs517_ = p0
                lhs434_ = d_3090_v41_
                lhs435_ = default__.safeIndex(587, (d_3090_v41_).length(0))
                lhs436_ = globalState
                lhs437_ = globalState
                lhs434_[lhs435_] = rhs515_
                lhs436_.f6 = rhs516_
                lhs437_.f6 = rhs517_
            d_3101_v49_: _dafny.MultiSet
            d_3101_v49_ = _dafny.MultiSet([p0])
            d_3102_v50_: _dafny.Map
            d_3102_v50_ = _dafny.Map({False: (d_3101_v49_).cardinality})
            d_3103_v51_: _dafny.Seq
            d_3103_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kypvkm"))
            d_3104_v52_: _dafny.Array
            nw493_ = _dafny.Array(None, 20)
            nw493_[int(0)] = self.f21
            nw493_[int(1)] = (self).f22
            nw493_[int(2)] = self.f21
            nw493_[int(3)] = not(self.f21)
            nw493_[int(4)] = self.f21
            nw493_[int(5)] = p2
            nw493_[int(6)] = (self).f22
            nw493_[int(7)] = p2
            nw493_[int(8)] = (self).f22
            nw493_[int(9)] = self.f21
            nw493_[int(10)] = (self).f22
            nw493_[int(11)] = False
            nw493_[int(12)] = not(self.f21)
            nw493_[int(13)] = p2
            nw493_[int(14)] = True
            nw493_[int(15)] = (self).f22
            nw493_[int(16)] = p2
            nw493_[int(17)] = False
            nw493_[int(18)] = p2
            nw493_[int(19)] = default__.fm69(False, len(d_3102_v50_), d_3068_v23_, len(d_3103_v51_), globalState)
            d_3104_v52_ = nw493_
            d_3105_v53_: _dafny.Seq
            d_3105_v53_ = _dafny.SeqWithoutIsStrInference([d_3104_v52_])
            d_3106_v54_: _dafny.Seq
            d_3106_v54_ = _dafny.SeqWithoutIsStrInference([False, p2, self.f21, p2])
            d_3107_v55_: _dafny.Array
            nw494_ = _dafny.Array(None, 26)
            nw494_[int(0)] = d_3104_v52_
            nw494_[int(1)] = d_3104_v52_
            nw494_[int(2)] = (d_3105_v53_)[default__.safeIndex(len(d_3106_v54_), len(d_3105_v53_))]
            nw494_[int(3)] = d_3104_v52_
            nw494_[int(4)] = d_3104_v52_
            nw494_[int(5)] = d_3104_v52_
            nw494_[int(6)] = d_3104_v52_
            nw494_[int(7)] = d_3104_v52_
            nw494_[int(8)] = d_3104_v52_
            nw494_[int(9)] = d_3104_v52_
            nw494_[int(10)] = d_3104_v52_
            nw494_[int(11)] = d_3104_v52_
            nw494_[int(12)] = d_3104_v52_
            nw494_[int(13)] = d_3104_v52_
            nw494_[int(14)] = d_3104_v52_
            nw494_[int(15)] = d_3104_v52_
            nw494_[int(16)] = d_3104_v52_
            nw494_[int(17)] = d_3104_v52_
            nw494_[int(18)] = d_3104_v52_
            nw494_[int(19)] = (d_3105_v53_)[default__.safeIndex(704, len(d_3105_v53_))]
            nw494_[int(20)] = d_3104_v52_
            nw494_[int(21)] = d_3104_v52_
            nw494_[int(22)] = d_3104_v52_
            nw494_[int(23)] = d_3104_v52_
            nw494_[int(24)] = d_3104_v52_
            nw494_[int(25)] = d_3104_v52_
            d_3107_v55_ = nw494_
            index540_ = default__.safeIndex(376, (d_3107_v55_).length(0))
            (d_3107_v55_)[index540_] = d_3104_v52_
            index541_ = default__.safeIndex(376, (d_3107_v55_).length(0))
            (d_3107_v55_)[index541_] = ((d_3104_v52_ if (self).f22 else (d_3105_v53_)[default__.safeIndex(p0, len(d_3105_v53_))]) if (self).f22 else d_3104_v52_)
            (globalState).f6 = (0) - (default__.safeModuloInt(p0, default__.safeDivisionInt(p0, (0) - (p0))))
            (globalState).f7 = (self).f22
        (globalState).f6 = (0) - (default__.safeDivisionInt(default__.fm0(p2, globalState), 124))
        if not ((self).f22) or ((self).f22):
            d_3108_v56_: _dafny.Seq
            d_3108_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbxauraqu"))
            d_3109_v57_: D18
            d_3109_v57_ = D18_DC52(True, (self).f22, d_3108_v56_)
            (globalState).f2 = (d_3109_v57_).cf96
            (globalState).f6 = p0
            (globalState).f6 = (0) - (default__.fm0((d_3108_v56_) < (default__.fm18(p2, -815, globalState)), globalState))
            if (self).f22:
                d_3110_v58_: _dafny.Array
                def lambda128_(d_3111_i4_):
                    return (self).f22

                init76_ = lambda128_
                nw495_ = _dafny.Array(None, 10)
                for i0_76_ in range(nw495_.length(0)):
                    nw495_[i0_76_] = init76_(i0_76_)
                d_3110_v58_ = nw495_
                d_3112_v59_: _dafny.Map
                d_3112_v59_ = _dafny.Map({True: d_3110_v58_})
                d_3113_v60_: _dafny.Array
                nw496_ = _dafny.Array(None, 5)
                d_3113_v60_ = nw496_
                d_3114_v61_: _dafny.Set
                d_3114_v61_ = _dafny.Set({d_3113_v60_})
                d_3115_v62_: D1
                d_3115_v62_ = D1_DC2(_dafny.CodePoint('g'))
                d_3116_v63_: D2
                d_3116_v63_ = D2_DC7(d_3110_v58_, d_3114_v61_, d_3115_v62_, p0, 930)
                d_3117_v64_: str
                d_3117_v64_ = _dafny.CodePoint('c')
                d_3118_v65_: D16
                d_3118_v65_ = D16_DC45((d_3116_v63_).cf10, d_3117_v64_, p0, p0, not(self.f21))
                d_3119_v66_: _dafny.Seq
                d_3119_v66_ = _dafny.SeqWithoutIsStrInference([d_3110_v58_])
                d_3120_v67_: _dafny.Array
                nw497_ = _dafny.Array(None, 28)
                nw497_[int(0)] = d_3110_v58_
                nw497_[int(1)] = d_3110_v58_
                nw497_[int(2)] = d_3110_v58_
                nw497_[int(3)] = d_3110_v58_
                nw497_[int(4)] = d_3110_v58_
                nw497_[int(5)] = d_3110_v58_
                nw497_[int(6)] = d_3110_v58_
                nw497_[int(7)] = d_3110_v58_
                nw497_[int(8)] = ((d_3112_v59_)[True] if (True) in (d_3112_v59_) else d_3110_v58_)
                nw497_[int(9)] = d_3110_v58_
                nw497_[int(10)] = d_3110_v58_
                nw497_[int(11)] = d_3110_v58_
                nw497_[int(12)] = d_3110_v58_
                nw497_[int(13)] = d_3110_v58_
                nw497_[int(14)] = d_3110_v58_
                nw497_[int(15)] = d_3110_v58_
                nw497_[int(16)] = d_3110_v58_
                nw497_[int(17)] = d_3110_v58_
                nw497_[int(18)] = d_3110_v58_
                nw497_[int(19)] = d_3110_v58_
                nw497_[int(20)] = (d_3118_v65_).cf81
                nw497_[int(21)] = (d_3119_v66_)[default__.safeIndex(0, len(d_3119_v66_))]
                nw497_[int(22)] = d_3110_v58_
                nw497_[int(23)] = d_3110_v58_
                nw497_[int(24)] = d_3110_v58_
                nw497_[int(25)] = d_3110_v58_
                nw497_[int(26)] = d_3110_v58_
                nw497_[int(27)] = d_3110_v58_
                d_3120_v67_ = nw497_
                index542_ = default__.safeIndex(373, (d_3120_v67_).length(0))
                (d_3120_v67_)[index542_] = d_3110_v58_
                index543_ = default__.safeIndex(373, (d_3120_v67_).length(0))
                def lambda129_(d_3121_i5_):
                    return (self).f22

                init77_ = lambda129_
                nw498_ = _dafny.Array(None, 9)
                for i0_77_ in range(nw498_.length(0)):
                    nw498_[i0_77_] = init77_(i0_77_)
                (d_3120_v67_)[index543_] = nw498_
                (globalState).f6 = p0
                d_3117_v64_ = d_3117_v64_
                d_3122_v68_: _dafny.Seq
                d_3122_v68_ = _dafny.SeqWithoutIsStrInference([self.f21])
                d_3123_v69_: _dafny.Map
                d_3123_v69_ = _dafny.Map({d_3122_v68_: p2})
                d_3124_v70_: D17
                d_3124_v70_ = D17_DC48(p2, p2, (self).f22)
                d_3125_v71_: _dafny.Seq
                d_3125_v71_ = _dafny.SeqWithoutIsStrInference([d_3124_v70_])
                d_3126_v72_: _dafny.Array
                nw499_ = _dafny.Array(int(0), 29)
                d_3126_v72_ = nw499_
                d_3127_v73_: _dafny.Seq
                d_3127_v73_ = _dafny.SeqWithoutIsStrInference([d_3126_v72_])
                d_3128_v74_: D3
                d_3128_v74_ = D3_DC9(d_3127_v73_)
                d_3129_v75_: D13
                d_3129_v75_ = D13_DC38(d_3128_v74_, False, p2)
                d_3130_v76_: _dafny.Map
                d_3130_v76_ = _dafny.Map({p0: (d_3129_v75_).cf69})
                d_3123_v69_ = default__.fm71((_dafny.SeqWithoutIsStrInference([d_3125_v71_, d_3125_v71_, d_3125_v71_])) + ((_dafny.SeqWithoutIsStrInference([d_3125_v71_ for d_3131_i6_ in range(default__.abs(-413))])).set(default__.safeIndex(len(d_3130_v76_), len(_dafny.SeqWithoutIsStrInference([d_3125_v71_ for d_3132_i6_ in range(default__.abs(-413))]))), (_dafny.SeqWithoutIsStrInference([d_3124_v70_, d_3124_v70_])).set(default__.safeIndex(len(d_3108_v56_), len(_dafny.SeqWithoutIsStrInference([d_3124_v70_, d_3124_v70_]))), D17_DC48(self.f21, p2, (self).f22)))), p0, False, globalState)
                d_3133_v77_: D4
                d_3133_v77_ = D4_DC13(_dafny.SeqWithoutIsStrInference([(0) - (p0), 186]))
                (globalState).f1 = default__.fm35((self).f22, (d_3133_v77_).cf29, globalState)
            elif True:
                d_3134_v78_: _dafny.Map
                d_3134_v78_ = _dafny.Map({(self).f22: p0})
                (globalState).f6 = (p0) + (default__.safeModuloInt(((d_3134_v78_)[p2] if (p2) in (d_3134_v78_) else p0), p0))
                d_3135_v79_: D10
                d_3135_v79_ = D10_DC29()
                d_3135_v79_ = d_3135_v79_
                d_3136_v80_: _dafny.Map
                d_3136_v80_ = _dafny.Map({p0: self.f21})
                (globalState).f7 = not (((d_3136_v80_)[(0) - (p0)] if ((0) - (p0)) in (d_3136_v80_) else False)) or (p2)
                rhs518_ = False
                rhs519_ = not(((self).f22 if not((self).f22) else self.f21))
                rhs520_ = p0
                lhs438_ = globalState
                lhs439_ = globalState
                lhs440_ = globalState
                lhs438_.f7 = rhs518_
                lhs439_.f2 = rhs519_
                lhs440_.f6 = rhs520_
                d_3137_v81_: D3
                d_3137_v81_ = D3_DC10(not((self).f22), p0, default__.fm9(self.f21, globalState), d_3108_v56_, len(_dafny.SeqWithoutIsStrInference([p0 for d_3138_i7_ in range(default__.abs(964))])))
                d_3139_v82_: _dafny.Map
                d_3139_v82_ = _dafny.Map({p0: p0})
                d_3140_v83_: _dafny.Array
                nw500_ = _dafny.Array(None, 14)
                nw500_[int(0)] = self.f21
                nw500_[int(1)] = (self).f22
                nw500_[int(2)] = not (p2) or (p2)
                nw500_[int(3)] = p2
                nw500_[int(4)] = ((d_3137_v81_).cf20) <= (p0)
                nw500_[int(5)] = p2
                nw500_[int(6)] = (self).f22
                nw500_[int(7)] = (d_3139_v82_) == (d_3139_v82_)
                nw500_[int(8)] = self.f21
                nw500_[int(9)] = (self).f22
                nw500_[int(10)] = (self).f22
                nw500_[int(11)] = (self).f22
                nw500_[int(12)] = not((494) < (p0))
                nw500_[int(13)] = not((p0) != (p0))
                d_3140_v83_ = nw500_
                d_3141_v84_: _dafny.Map
                d_3141_v84_ = _dafny.Map({self.f21: self.f21})
                index544_ = default__.safeIndex(277, (d_3140_v83_).length(0))
                (d_3140_v83_)[index544_] = (p0) >= (len(_dafny.Map({d_3141_v84_: p2})))
                index545_ = default__.safeIndex(277, (d_3140_v83_).length(0))
                (d_3140_v83_)[index545_] = ((p0) * (p0)) != (p0)
            d_3142_v85_: _dafny.MultiSet
            d_3142_v85_ = _dafny.MultiSet([(self).f22])
            d_3143_v86_: _dafny.Seq
            d_3143_v86_ = _dafny.SeqWithoutIsStrInference([((d_3142_v85_).set((self).f22, default__.abs(153))) - (d_3142_v85_)])
            d_3143_v86_ = (d_3143_v86_) + (d_3143_v86_)
        elif True:
            d_3144_v87_: _dafny.Array
            def lambda130_(d_3145_p2_):
                def lambda131_(d_3146_i8_):
                    return d_3145_p2_

                return lambda131_

            init78_ = lambda130_(p2)
            nw501_ = _dafny.Array(None, 25)
            for i0_78_ in range(nw501_.length(0)):
                nw501_[i0_78_] = init78_(i0_78_)
            d_3144_v87_ = nw501_
            index546_ = default__.safeIndex(559, (d_3144_v87_).length(0))
            (d_3144_v87_)[index546_] = (self).f22
            index547_ = default__.safeIndex(559, (d_3144_v87_).length(0))
            (d_3144_v87_)[index547_] = p2
            r0 = p0
            d_3147_v88_: _dafny.Array
            nw502_ = _dafny.Array(_dafny.MultiSet({}), 27)
            d_3147_v88_ = nw502_
            d_3148_v89_: _dafny.Seq
            d_3148_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyhf"))
            d_3149_v90_: _dafny.Map
            d_3149_v90_ = _dafny.Map({(self).f22: (d_3144_v87_)[default__.safeIndex(559, (d_3144_v87_).length(0))]})
            d_3150_v91_: T0
            nw503_ = C1()
            nw503_.ctor__(d_3144_v87_, len(d_3148_v89_), (0) - (len(d_3149_v90_)), p0, (d_3144_v87_)[default__.safeIndex(559, (d_3144_v87_).length(0))], d_3148_v89_)
            d_3150_v91_ = nw503_
            d_3151_v92_: _dafny.MultiSet
            d_3151_v92_ = _dafny.MultiSet([d_3150_v91_])
            index548_ = default__.safeIndex(914, (d_3147_v88_).length(0))
            (d_3147_v88_)[index548_] = d_3151_v92_
            d_3152_v93_: _dafny.Seq
            d_3152_v93_ = _dafny.SeqWithoutIsStrInference([d_3150_v91_, d_3150_v91_, d_3150_v91_, d_3150_v91_, d_3150_v91_])
            index549_ = default__.safeIndex(914, (d_3147_v88_).length(0))
            (d_3147_v88_)[index549_] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_3150_v91_])) + (d_3152_v93_))
            if (self).f22:
                d_3153_v94_: str
                d_3153_v94_ = _dafny.CodePoint('o')
                d_3154_v95_: _dafny.Map
                d_3154_v95_ = _dafny.Map({self.f21: d_3153_v94_})
                d_3154_v95_ = (_dafny.Map({not((self).f22): d_3153_v94_})).set((d_3148_v89_) <= (_dafny.SeqWithoutIsStrInference([d_3153_v94_ for d_3155_i9_ in range(default__.abs(952))])), d_3153_v94_)
                d_3156_v96_: D1
                d_3156_v96_ = D1_DC5((d_3150_v91_).f11, p2, self.f21)
                index550_ = default__.safeIndex(559, (d_3144_v87_).length(0))
                (d_3144_v87_)[index550_] = (d_3156_v96_).cf8
                d_3157_v97_: _dafny.Array
                nw504_ = _dafny.Array(None, 1)
                nw504_[int(0)] = (d_3150_v91_).f11
                d_3157_v97_ = nw504_
                index551_ = default__.safeIndex(85, (d_3157_v97_).length(0))
                (d_3157_v97_)[index551_] = 864
                index552_ = default__.safeIndex(85, (d_3157_v97_).length(0))
                rhs521_ = d_3153_v94_
                rhs522_ = (d_3150_v91_).f11
                rhs523_ = (d_3148_v89_)[default__.safeIndex(((d_3150_v91_).f12) * ((d_3150_v91_).f12), len(d_3148_v89_))]
                rhs524_ = default__.fm21((d_3150_v91_).f11, self.f21, globalState)
                lhs441_ = d_3157_v97_
                lhs442_ = default__.safeIndex(85, (d_3157_v97_).length(0))
                d_3153_v94_ = rhs521_
                lhs441_[lhs442_] = rhs522_
                d_3153_v94_ = rhs523_
                d_3153_v94_ = rhs524_
                d_3158_v98_: _dafny.Array
                nw505_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                d_3158_v98_ = nw505_
                d_3158_v98_ = d_3158_v98_
                d_3159_v99_: _dafny.Seq
                d_3159_v99_ = _dafny.SeqWithoutIsStrInference([d_3148_v89_, d_3148_v89_, (d_3148_v89_) + (d_3148_v89_), d_3148_v89_, _dafny.SeqWithoutIsStrInference([d_3153_v94_ for d_3160_i10_ in range(default__.abs(775))])])
                d_3159_v99_ = d_3159_v99_
            elif True:
                (globalState).f2 = self.f21
                d_3161_v100_: C7
                nw506_ = C7()
                nw506_.ctor__()
                d_3161_v100_ = nw506_
                d_3162_v101_: _dafny.Array
                nw507_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_3162_v101_ = nw507_
                d_3163_v102_: str
                d_3163_v102_ = _dafny.CodePoint('d')
                d_3164_v103_: D1
                d_3164_v103_ = D1_DC3(self.f21)
                d_3165_v104_: _dafny.Map
                d_3165_v104_ = _dafny.Map({d_3164_v103_: d_3162_v101_})
                d_3166_v105_: _dafny.Array
                d_3166_v105_ = ((d_3165_v104_)[d_3164_v103_] if (d_3164_v103_) in (d_3165_v104_) else d_3162_v101_)
                d_3167_v106_: _dafny.Array
                nw508_ = _dafny.Array(_dafny.MultiSet({}), 10)
                d_3167_v106_ = nw508_
                d_3168_v107_: D18
                d_3168_v107_ = D18_DC54(D18_DC51(d_3167_v106_))
                d_3169_v108_: _dafny.Array
                nw509_ = _dafny.Array(None, 9)
                nw509_[int(0)] = d_3168_v107_
                nw509_[int(1)] = d_3168_v107_
                nw509_[int(2)] = d_3168_v107_
                nw509_[int(3)] = d_3168_v107_
                nw509_[int(4)] = d_3168_v107_
                nw509_[int(5)] = d_3168_v107_
                nw509_[int(6)] = d_3168_v107_
                nw509_[int(7)] = d_3168_v107_
                nw509_[int(8)] = d_3168_v107_
                d_3169_v108_ = nw509_
                d_3170_v109_: _dafny.Map
                d_3170_v109_ = _dafny.Map({d_3169_v108_: (self).f22})
                d_3171_v110_: _dafny.Set
                d_3171_v110_ = _dafny.Set({self.f21, ((d_3170_v109_)[d_3169_v108_] if (d_3169_v108_) in (d_3170_v109_) else (self).f22), self.f21, not((d_3144_v87_)[default__.safeIndex(559, (d_3144_v87_).length(0))]), self.f21})
                index553_ = default__.safeIndex(559, (d_3144_v87_).length(0))
                rhs525_ = (0) - ((d_3150_v91_).f12)
                rhs526_ = (d_3150_v91_).f12
                rhs527_ = (d_3166_v105_)
                rhs528_ = d_3163_v102_
                rhs529_ = ((d_3171_v110_).intersection(d_3171_v110_)).isdisjoint(d_3171_v110_)
                lhs443_ = globalState
                lhs444_ = d_3144_v87_
                lhs445_ = default__.safeIndex(559, (d_3144_v87_).length(0))
                r0 = rhs525_
                lhs443_.f6 = rhs526_
                d_3162_v101_ = rhs527_
                d_3163_v102_ = rhs528_
                lhs444_[lhs445_] = rhs529_
                index554_ = default__.safeIndex(559, (d_3144_v87_).length(0))
                (d_3144_v87_)[index554_] = self.f21
                (globalState).f6 = default__.safeDivisionInt(default__.safeModuloInt((d_3150_v91_).f12, (d_3150_v91_).f11), ((0) - (default__.fm0(self.f21, globalState))) + ((d_3150_v91_).f11))
            d_3172_v111_: _dafny.Set
            d_3172_v111_ = _dafny.Set({self.f21})
            d_3173_v112_: D12
            d_3173_v112_ = D12_DC35(_dafny.Set({(self).f22}))
            d_3174_v113_: D1
            d_3174_v113_ = D1_DC5(len((d_3172_v111_).intersection((d_3173_v112_).cf63)), p2, True)
            d_3175_v114_: str
            d_3175_v114_ = _dafny.CodePoint('m')
            d_3176_v115_: _dafny.Array
            nw510_ = _dafny.Array(int(0), 1)
            d_3176_v115_ = nw510_
            d_3177_v116_: _dafny.Map
            d_3177_v116_ = _dafny.Map({(d_3150_v91_).f11: (self).f22})
            index555_ = default__.safeIndex(271, (d_3176_v115_).length(0))
            (d_3176_v115_)[index555_] = len(d_3177_v116_)
            d_3178_v117_: _dafny.Seq
            d_3178_v117_ = _dafny.SeqWithoutIsStrInference([d_3148_v89_])
            index556_ = default__.safeIndex(271, (d_3176_v115_).length(0))
            rhs530_ = d_3174_v113_
            rhs531_ = len(d_3178_v117_)
            rhs532_ = d_3175_v114_
            rhs533_ = default__.safeModuloInt((d_3150_v91_).f11, (d_3150_v91_).f11)
            lhs446_ = d_3176_v115_
            lhs447_ = default__.safeIndex(271, (d_3176_v115_).length(0))
            d_3174_v113_ = rhs530_
            r0 = rhs531_
            d_3175_v114_ = rhs532_
            lhs446_[lhs447_] = rhs533_
        d_3179_v118_: _dafny.Seq
        d_3179_v118_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsx"))
        d_3180_v119_: _dafny.Map
        d_3180_v119_ = _dafny.Map({p0: (self).f22})
        d_3181_v120_: str
        d_3181_v120_ = _dafny.CodePoint('x')
        d_3182_v121_: _dafny.Set
        d_3182_v121_ = _dafny.Set({p0})
        d_3183_v122_: _dafny.Seq
        d_3183_v122_ = _dafny.SeqWithoutIsStrInference([False])
        d_3184_v123_: _dafny.Set
        d_3184_v123_ = _dafny.Set({(d_3183_v122_)[default__.safeIndex(p0, len(d_3183_v122_))]})
        d_3185_v124_: _dafny.Array
        nw511_ = _dafny.Array(None, 27)
        nw511_[int(0)] = p0
        nw511_[int(1)] = p0
        nw511_[int(2)] = p0
        nw511_[int(3)] = (p0) + (len((d_3179_v118_).set(default__.safeIndex(len(d_3180_v119_), len(d_3179_v118_)), d_3181_v120_)))
        nw511_[int(4)] = p0
        nw511_[int(5)] = len(d_3179_v118_)
        nw511_[int(6)] = p0
        nw511_[int(7)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))
        nw511_[int(8)] = p0
        nw511_[int(9)] = p0
        nw511_[int(10)] = len(d_3182_v121_)
        nw511_[int(11)] = p0
        nw511_[int(12)] = p0
        nw511_[int(13)] = (p0 if p2 else p0)
        nw511_[int(14)] = p0
        nw511_[int(15)] = default__.safeDivisionInt(p0, p0)
        nw511_[int(16)] = len(default__.fm48(globalState))
        nw511_[int(17)] = len((d_3184_v123_).intersection(d_3184_v123_))
        nw511_[int(18)] = p0
        nw511_[int(19)] = p0
        nw511_[int(20)] = p0
        nw511_[int(21)] = default__.safeModuloInt(p0, 761)
        nw511_[int(22)] = p0
        nw511_[int(23)] = (len(default__.fm18(self.f21, p0, globalState))) - (p0)
        nw511_[int(24)] = p0
        nw511_[int(25)] = len((((d_3179_v118_).set(default__.safeIndex((0) - (len(d_3183_v122_)), len(d_3179_v118_)), d_3181_v120_)).set(default__.safeIndex(p0, len((d_3179_v118_).set(default__.safeIndex((0) - (len(d_3183_v122_)), len(d_3179_v118_)), d_3181_v120_))), d_3181_v120_)) + (d_3179_v118_))
        nw511_[int(26)] = p0
        d_3185_v124_ = nw511_
        d_3186_v125_: _dafny.Seq
        d_3186_v125_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_3187_v126_: _dafny.MultiSet
        d_3187_v126_ = _dafny.MultiSet([p0, p0, (d_3186_v125_)[default__.safeIndex(p0, len(d_3186_v125_))], len(d_3179_v118_)])
        index557_ = default__.safeIndex(197, (d_3185_v124_).length(0))
        (d_3185_v124_)[index557_] = ((d_3187_v126_)[p0] if (p0) in (d_3187_v126_) else 974)
        index558_ = default__.safeIndex(197, (d_3185_v124_).length(0))
        rhs534_ = (0) - ((0) - ((p0) + ((p0) + (p0))))
        rhs535_ = (d_3183_v122_)[default__.safeIndex(703, len(d_3183_v122_))]
        lhs448_ = d_3185_v124_
        lhs449_ = default__.safeIndex(197, (d_3185_v124_).length(0))
        lhs450_ = globalState
        lhs448_[lhs449_] = rhs534_
        lhs450_.f2 = rhs535_
        d_3188_v127_: _dafny.Array
        nw512_ = _dafny.Array(_dafny.Set({}), 15)
        d_3188_v127_ = nw512_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_3188_v127_).length(0)):
            d_3189_i11_: int = guard_loop_10_
            if (True) and (((0) <= (d_3189_i11_)) and ((d_3189_i11_) < ((d_3188_v127_).length(0)))):
                (d_3188_v127_)[(d_3189_i11_)] = (d_3184_v123_) - (d_3184_v123_)
        d_3190_v128_: D14
        d_3190_v128_ = D14_DC41(not (not(self.f21)) or ((self).f22), p2, (d_3179_v118_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpiljk"))))
        source43_ = d_3190_v128_
        if source43_.is_DC41:
            d_3191___mcc_h0_ = source43_.cf73
            d_3192___mcc_h1_ = source43_.cf74
            d_3193___mcc_h2_ = source43_.cf75
            d_3194_cf75_ = d_3193___mcc_h2_
            d_3195_cf74_ = d_3192___mcc_h1_
            d_3196_cf73_ = d_3191___mcc_h0_
            d_3197_v129_: _dafny.Map
            d_3197_v129_ = _dafny.Map({self.f21: d_3181_v120_})
            d_3198_v130_: _dafny.Seq
            d_3198_v130_ = _dafny.SeqWithoutIsStrInference([d_3197_v129_])
            d_3195_cf74_ = (((d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]) + (len((d_3198_v130_)[default__.safeIndex((d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))], len(d_3198_v130_))]))) >= (default__.safeDivisionInt(len(d_3186_v125_), (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]))
            d_3199_v131_: _dafny.MultiSet
            d_3199_v131_ = _dafny.MultiSet([d_3181_v120_, d_3181_v120_])
            d_3200_v132_: D0
            d_3200_v132_ = D0_DC0(d_3199_v131_)
            d_3201_v133_: _dafny.Set
            d_3201_v133_ = _dafny.Set({d_3200_v132_})
            pat_let_tv52_ = d_3196_cf73_
            pat_let_tv53_ = d_3201_v133_
            def iife189_(_pat_let57_0):
                def iife190_(d_3202_dt__update__tmp_h0_):
                    def iife191_(_pat_let58_0):
                        def iife192_(d_3203_dt__update_hcf27_h0_):
                            def iife193_(_pat_let59_0):
                                def iife194_(d_3204_dt__update_hcf26_h0_):
                                    return D3_DC11((d_3202_dt__update__tmp_h0_).cf24, (d_3202_dt__update__tmp_h0_).cf25, d_3204_dt__update_hcf26_h0_, d_3203_dt__update_hcf27_h0_)
                                return iife194_(_pat_let59_0)
                            return iife193_(pat_let_tv53_)
                        return iife192_(_pat_let58_0)
                    return iife191_(pat_let_tv52_)
                return iife190_(_pat_let57_0)
            source44_ = iife189_(p1)
            if source44_.is_DC10:
                d_3205___mcc_h4_ = source44_.cf19
                d_3206___mcc_h5_ = source44_.cf20
                d_3207___mcc_h6_ = source44_.cf21
                d_3208___mcc_h7_ = source44_.cf22
                d_3209___mcc_h8_ = source44_.cf23
                d_3210_cf23_ = d_3209___mcc_h8_
                d_3211_cf22_ = d_3208___mcc_h7_
                d_3212_cf21_ = d_3207___mcc_h6_
                d_3213_cf20_ = d_3206___mcc_h5_
                d_3214_cf19_ = d_3205___mcc_h4_
                (globalState).f2 = d_3196_cf73_
                d_3215_v134_: C5
                nw513_ = C5()
                nw513_.ctor__()
                d_3215_v134_ = nw513_
                rhs536_ = d_3214_cf19_
                rhs537_ = d_3215_v134_
                lhs451_ = globalState
                lhs451_.f7 = rhs536_
                d_3215_v134_ = rhs537_
                d_3216_v135_: D12
                d_3216_v135_ = D12_DC35(_dafny.Set({d_3195_cf74_, (self).f22}))
                pat_let_tv54_ = d_3184_v123_
                def iife195_(_pat_let60_0):
                    def iife196_(d_3217_dt__update__tmp_h1_):
                        def iife197_(_pat_let61_0):
                            def iife198_(d_3218_dt__update_hcf63_h0_):
                                return D12_DC35(d_3218_dt__update_hcf63_h0_)
                            return iife198_(_pat_let61_0)
                        return iife197_(pat_let_tv54_)
                    return iife196_(_pat_let60_0)
                d_3216_v135_ = iife195_(d_3216_v135_)
                d_3219_v136_: _dafny.Array
                nw514_ = _dafny.Array(None, 6)
                nw514_[int(0)] = ((d_3180_v119_)[d_3210_cf23_] if (d_3210_cf23_) in (d_3180_v119_) else d_3196_cf73_)
                nw514_[int(1)] = d_3214_cf19_
                nw514_[int(2)] = p2
                nw514_[int(3)] = d_3196_cf73_
                nw514_[int(4)] = not(d_3196_cf73_)
                nw514_[int(5)] = (_dafny.Set({p0, p0, d_3213_cf20_})).ispropersubset(d_3182_v121_)
                d_3219_v136_ = nw514_
                index559_ = default__.safeIndex(638, (d_3219_v136_).length(0))
                (d_3219_v136_)[index559_] = default__.fm69(p2, len(d_3186_v125_), (d_3199_v131_).set(_dafny.CodePoint('x'), default__.abs((d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))])), (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))], globalState)
                index560_ = default__.safeIndex(638, (d_3219_v136_).length(0))
                rhs538_ = (self).f22
                rhs539_ = not (d_3214_cf19_) or ((self).f22)
                lhs452_ = globalState
                lhs453_ = d_3219_v136_
                lhs454_ = default__.safeIndex(638, (d_3219_v136_).length(0))
                lhs452_.f2 = rhs538_
                lhs453_[lhs454_] = rhs539_
            elif source44_.is_DC11:
                d_3220___mcc_h9_ = source44_.cf24
                d_3221___mcc_h10_ = source44_.cf25
                d_3222___mcc_h11_ = source44_.cf26
                d_3223___mcc_h12_ = source44_.cf27
                d_3224_cf27_ = d_3223___mcc_h12_
                d_3225_cf26_ = d_3222___mcc_h11_
                d_3226_cf25_ = d_3221___mcc_h10_
                d_3227_cf24_ = d_3220___mcc_h9_
                d_3185_v124_ = d_3185_v124_
                d_3228_v137_: _dafny.Array
                nw515_ = _dafny.Array(None, 12)
                nw515_[int(0)] = not((p0) >= (d_3227_cf24_))
                nw515_[int(1)] = d_3196_cf73_
                nw515_[int(2)] = p2
                nw515_[int(3)] = not(not (self.f21) or (d_3196_cf73_))
                nw515_[int(4)] = (d_3182_v121_) != (d_3182_v121_)
                nw515_[int(5)] = ((d_3180_v119_)[(0) - ((d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))])] if ((0) - ((d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))])) in (d_3180_v119_) else (self).f22)
                nw515_[int(6)] = self.f21
                nw515_[int(7)] = (d_3224_cf27_ if d_3224_cf27_ else d_3224_cf27_)
                nw515_[int(8)] = default__.fm69(False, d_3227_cf24_, d_3199_v131_, len(d_3179_v118_), globalState)
                nw515_[int(9)] = d_3195_cf74_
                nw515_[int(10)] = (d_3190_v128_).cf74
                nw515_[int(11)] = d_3195_cf74_
                d_3228_v137_ = nw515_
                d_3228_v137_ = d_3228_v137_
                d_3229_v138_: _dafny.Map
                d_3229_v138_ = _dafny.Map({(self).f22: d_3224_cf27_})
                d_3230_v139_: _dafny.Array
                nw516_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_3230_v139_ = nw516_
                d_3231_v140_: _dafny.Map
                d_3231_v140_ = _dafny.Map({d_3230_v139_: d_3229_v138_})
                d_3232_v141_: _dafny.Set
                d_3232_v141_ = _dafny.Set({d_3229_v138_, (d_3229_v138_).set(not(d_3196_cf73_), p2), (_dafny.Map({default__.fm69((self).f22, (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))], d_3199_v131_, p0, globalState): d_3224_cf27_})) | (((d_3231_v140_)[d_3230_v139_] if (d_3230_v139_) in (d_3231_v140_) else d_3229_v138_)), d_3229_v138_})
                d_3232_v141_ = d_3232_v141_
                (globalState).f7 = d_3195_cf74_
            elif source44_.is_DC9:
                d_3233___mcc_h13_ = source44_.cf18
                d_3234_cf18_ = d_3233___mcc_h13_
                (globalState).f2 = (self).f22
                d_3235_v142_: bool
                d_3236_v143_: bool
                d_3237_v144_: bool
                out89_: bool
                out90_: bool
                out91_: bool
                out89_, out90_, out91_ = (self).m8(globalState)
                d_3235_v142_ = out89_
                d_3236_v143_ = out90_
                d_3237_v144_ = out91_
                d_3238_v145_: _dafny.Map
                d_3238_v145_ = _dafny.Map({d_3181_v120_: (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]})
                d_3239_v146_: _dafny.Array
                nw517_ = _dafny.Array(None, 12)
                nw517_[int(0)] = d_3236_v143_
                nw517_[int(1)] = d_3196_cf73_
                nw517_[int(2)] = p2
                nw517_[int(3)] = (self).f22
                nw517_[int(4)] = (self).f22
                nw517_[int(5)] = d_3237_v144_
                nw517_[int(6)] = d_3235_v142_
                nw517_[int(7)] = (d_3183_v122_)[default__.safeIndex(((d_3238_v145_)[d_3181_v120_] if (d_3181_v120_) in (d_3238_v145_) else p0), len(d_3183_v122_))]
                nw517_[int(8)] = d_3237_v144_
                nw517_[int(9)] = d_3237_v144_
                nw517_[int(10)] = p2
                nw517_[int(11)] = d_3235_v142_
                d_3239_v146_ = nw517_
                d_3240_v147_: T2
                nw518_ = C1()
                nw518_.ctor__(d_3239_v146_, (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))], (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))], p0, d_3195_cf74_, d_3179_v118_)
                d_3240_v147_ = nw518_
                d_3241_v148_: _dafny.Map
                d_3241_v148_ = _dafny.Map({_dafny.Set({d_3240_v147_}): d_3239_v146_})
                d_3241_v148_ = d_3241_v148_
                d_3242_v149_: _dafny.Map
                d_3242_v149_ = _dafny.Map({d_3236_v143_: d_3195_cf74_})
                d_3243_v150_: C4
                nw519_ = C4()
                nw519_.ctor__((d_3240_v147_).f13, 734, d_3195_cf74_, (d_3240_v147_).f13, (d_3240_v147_).f12, default__.fm0(d_3196_cf73_, globalState))
                d_3243_v150_ = nw519_
                d_3244_v151_: _dafny.Map
                d_3244_v151_ = _dafny.Map({(d_3240_v147_).f12: (d_3243_v150_ if ((d_3242_v149_)[self.f21] if (self.f21) in (d_3242_v149_) else d_3237_v144_) else d_3243_v150_)})
                d_3244_v151_ = (d_3244_v151_).set(51, (d_3243_v150_ if d_3195_cf74_ else d_3243_v150_))
            elif True:
                d_3245___mcc_h14_ = source44_.cf28
                d_3246_cf28_ = d_3245___mcc_h14_
                d_3247_v152_: _dafny.Set
                d_3247_v152_ = _dafny.Set({d_3181_v120_, d_3181_v120_})
                d_3248_v153_: C7
                nw520_ = C7()
                nw520_.ctor__()
                d_3248_v153_ = nw520_
                d_3249_v154_: _dafny.Seq
                d_3249_v154_ = _dafny.SeqWithoutIsStrInference([d_3248_v153_])
                d_3250_v155_: _dafny.Map
                d_3250_v155_ = _dafny.Map({d_3181_v120_: not(default__.fm7(True, p0, globalState))})
                d_3251_v156_: D17
                d_3251_v156_ = D17_DC48(not(d_3195_cf74_), False, False)
                d_3252_v157_: _dafny.Array
                nw521_ = _dafny.Array(None, 14)
                nw521_[int(0)] = self.f21
                nw521_[int(1)] = (self).f22
                nw521_[int(2)] = ((d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]) not in ((_dafny.Map({len(d_3247_v152_): (d_3249_v154_)[default__.safeIndex((d_3186_v125_)[default__.safeIndex(len(d_3250_v155_), len(d_3186_v125_))], len(d_3249_v154_))]})).set(p0, d_3248_v153_))
                nw521_[int(3)] = (self).f22
                nw521_[int(4)] = (d_3183_v122_)[default__.safeIndex(p0, len(d_3183_v122_))]
                nw521_[int(5)] = (d_3251_v156_).cf92
                nw521_[int(6)] = d_3195_cf74_
                nw521_[int(7)] = self.f21
                nw521_[int(8)] = self.f21
                nw521_[int(9)] = self.f21
                nw521_[int(10)] = self.f21
                nw521_[int(11)] = d_3195_cf74_
                nw521_[int(12)] = (self).f22
                nw521_[int(13)] = (d_3183_v122_)[default__.safeIndex(p0, len(d_3183_v122_))]
                d_3252_v157_ = nw521_
                index561_ = default__.safeIndex(430, (d_3252_v157_).length(0))
                (d_3252_v157_)[index561_] = d_3195_cf74_
                d_3253_v158_: _dafny.Map
                d_3253_v158_ = _dafny.Map({d_3185_v124_: (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]})
                index562_ = default__.safeIndex(430, (d_3252_v157_).length(0))
                rhs540_ = (d_3185_v124_) in (d_3253_v158_)
                rhs541_ = (default__.fm18(d_3195_cf74_, (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))], globalState)) + (d_3179_v118_)
                lhs455_ = d_3252_v157_
                lhs456_ = default__.safeIndex(430, (d_3252_v157_).length(0))
                lhs455_[lhs456_] = rhs540_
                d_3194_cf75_ = rhs541_
                (globalState).f6 = (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]
                d_3187_v126_ = d_3187_v126_
                (globalState).f7 = default__.fm7(d_3196_cf73_, (0) - (p0), globalState)
            d_3254_v159_: _dafny.Seq
            d_3254_v159_ = _dafny.SeqWithoutIsStrInference([d_3185_v124_, d_3185_v124_, d_3185_v124_])
            d_3255_v160_: _dafny.Array
            nw522_ = _dafny.Array(None, 19)
            nw522_[int(0)] = d_3185_v124_
            nw522_[int(1)] = d_3185_v124_
            nw522_[int(2)] = (d_3185_v124_ if d_3195_cf74_ else (d_3254_v159_)[default__.safeIndex(28, len(d_3254_v159_))])
            nw522_[int(3)] = d_3185_v124_
            nw522_[int(4)] = d_3185_v124_
            nw522_[int(5)] = d_3185_v124_
            nw522_[int(6)] = d_3185_v124_
            nw522_[int(7)] = d_3185_v124_
            nw522_[int(8)] = d_3185_v124_
            nw522_[int(9)] = d_3185_v124_
            nw522_[int(10)] = d_3185_v124_
            nw522_[int(11)] = d_3185_v124_
            nw522_[int(12)] = d_3185_v124_
            nw522_[int(13)] = (d_3185_v124_ if default__.fm69(d_3196_cf73_, p0, d_3199_v131_, (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))], globalState) else d_3185_v124_)
            nw522_[int(14)] = d_3185_v124_
            nw522_[int(15)] = d_3185_v124_
            nw522_[int(16)] = d_3185_v124_
            nw522_[int(17)] = d_3185_v124_
            nw522_[int(18)] = (d_3185_v124_ if d_3195_cf74_ else d_3185_v124_)
            d_3255_v160_ = nw522_
            d_3256_v161_: _dafny.Map
            d_3256_v161_ = _dafny.Map({(d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]: d_3185_v124_})
            d_3257_v162_: _dafny.Map
            d_3257_v162_ = _dafny.Map({d_3195_cf74_: ((d_3256_v161_)[len(d_3186_v125_)] if (len(d_3186_v125_)) in (d_3256_v161_) else d_3185_v124_)})
            d_3258_v163_: _dafny.Map
            d_3258_v163_ = _dafny.Map({d_3195_cf74_: p2})
            d_3259_v164_: _dafny.Seq
            d_3259_v164_ = _dafny.SeqWithoutIsStrInference([d_3258_v163_, _dafny.Map({d_3195_cf74_: False})])
            d_3260_v165_: D19
            d_3260_v165_ = D19_DC56(((d_3257_v162_)[self.f21] if (self.f21) in (d_3257_v162_) else d_3185_v124_), d_3196_cf73_, (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))], d_3259_v164_)
            d_3261_v166_: _dafny.Map
            d_3261_v166_ = _dafny.Map({(d_3260_v165_).cf107: (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]})
            d_3262_v167_: _dafny.Map
            d_3262_v167_ = _dafny.Map({((d_3186_v125_)[default__.safeIndex((d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))], len(d_3186_v125_))]) in (d_3261_v166_): (self).f22})
            rhs542_ = len(d_3183_v122_)
            rhs543_ = ((d_3258_v163_)[d_3195_cf74_] if (d_3195_cf74_) in (d_3258_v163_) else self.f21)
            rhs544_ = d_3255_v160_
            rhs545_ = ((default__.fm9(d_3196_cf73_, globalState)) | (_dafny.Map({d_3194_cf75_: (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]}))) != (default__.fm9((self).f22, globalState))
            rhs546_ = self.f21
            lhs457_ = globalState
            lhs458_ = globalState
            lhs459_ = self
            r0 = rhs542_
            lhs457_.f2 = rhs543_
            d_3255_v160_ = rhs544_
            lhs458_.f2 = rhs545_
            lhs459_.f21 = rhs546_
            (globalState).f7 = self.f21
        elif True:
            d_3263___mcc_h3_ = source43_.cf72
            d_3264_cf72_ = d_3263___mcc_h3_
            d_3265_v168_: _dafny.Map
            d_3265_v168_ = _dafny.Map({self.f21: (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]})
            d_3266_v169_: _dafny.Map
            d_3266_v169_ = _dafny.Map({True: self.f21})
            d_3267_v170_: _dafny.Array
            nw523_ = _dafny.Array(_dafny.Map({}), 20)
            d_3267_v170_ = nw523_
            d_3268_v171_: _dafny.Map
            d_3268_v171_ = _dafny.Map({d_3267_v170_: self.f21})
            d_3269_v172_: _dafny.Array
            nw524_ = _dafny.Array(None, 28)
            nw524_[int(0)] = default__.fm7(not(self.f21), (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))], globalState)
            nw524_[int(1)] = False
            nw524_[int(2)] = not(default__.fm7((self).f22, 638, globalState))
            nw524_[int(3)] = self.f21
            nw524_[int(4)] = p2
            nw524_[int(5)] = (952) >= (((d_3265_v168_)[default__.fm7(False, len(d_3184_v123_), globalState)] if (default__.fm7(False, len(d_3184_v123_), globalState)) in (d_3265_v168_) else (d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]))
            nw524_[int(6)] = p2
            nw524_[int(7)] = self.f21
            nw524_[int(8)] = self.f21
            nw524_[int(9)] = True
            nw524_[int(10)] = p2
            nw524_[int(11)] = self.f21
            nw524_[int(12)] = p2
            nw524_[int(13)] = ((d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))]) > (p0)
            nw524_[int(14)] = ((d_3266_v169_)[True] if (True) in (d_3266_v169_) else default__.fm7((self).f22, -56, globalState))
            nw524_[int(15)] = default__.fm7((self).f22, (d_3186_v125_)[default__.safeIndex(789, len(d_3186_v125_))], globalState)
            nw524_[int(16)] = (p2) == (p2)
            nw524_[int(17)] = self.f21
            nw524_[int(18)] = (self).f22
            nw524_[int(19)] = (self).f22
            nw524_[int(20)] = ((d_3268_v171_)[d_3267_v170_] if (d_3267_v170_) in (d_3268_v171_) else self.f21)
            nw524_[int(21)] = self.f21
            nw524_[int(22)] = (p2) or (self.f21)
            nw524_[int(23)] = p2
            nw524_[int(24)] = (d_3183_v122_) != (d_3183_v122_)
            nw524_[int(25)] = not ((self).f22) or (p2)
            nw524_[int(26)] = p2
            nw524_[int(27)] = p2
            d_3269_v172_ = nw524_
            d_3269_v172_ = d_3269_v172_
            index563_ = default__.safeIndex(993, (p3).length(0))
            (p3)[index563_] = d_3179_v118_
            index564_ = default__.safeIndex(993, (p3).length(0))
            (p3)[index564_] = (d_3179_v118_) + (d_3179_v118_)
            d_3270_v173_: _dafny.Map
            d_3270_v173_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuxngc"))) + (d_3179_v118_): d_3184_v123_})
            d_3270_v173_ = (d_3270_v173_).set((d_3179_v118_) + ((p3)[default__.safeIndex(993, (p3).length(0))]), d_3184_v123_)
            (globalState).f7 = p2
        r0 = ((0) - (default__.safeModuloInt(p0, p0))) - ((d_3185_v124_)[default__.safeIndex(197, (d_3185_v124_).length(0))])
        d_3271_v174_: _dafny.Seq
        d_3271_v174_ = _dafny.SeqWithoutIsStrInference([d_3180_v119_, d_3180_v119_])
        d_3272_v175_: D2
        d_3272_v175_ = D2_DC6(d_3179_v118_)
        d_3273_v176_: T1
        nw525_ = C6()
        nw525_.ctor__(d_3271_v174_, p0, self.f21, (d_3272_v175_).cf9, (p0) * (p0), (0) - (p0))
        d_3273_v176_ = nw525_
        r1 = d_3273_v176_
        d_3274_v177_: _dafny.MultiSet
        d_3274_v177_ = _dafny.MultiSet([d_3181_v120_])
        d_3275_v178_: _dafny.MultiSet
        d_3275_v178_ = _dafny.MultiSet([(self).f22, p2])
        r2 = (_dafny.MultiSet([self.f21, default__.fm69(True, -733, d_3274_v177_, (d_3273_v176_).f12, globalState)])).intersection(d_3275_v178_)
        return r0, r1, r2

    def m8(self, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        d_3276_i0_: int
        d_3276_i0_ = 0
        with _dafny.label("16"):
            while (self).f22:
                with _dafny.c_label("16"):
                    if (d_3276_i0_) >= (100):
                        raise _dafny.Break("16")
                    d_3276_i0_ = (d_3276_i0_) + (1)
                    d_3277_v0_: int
                    d_3277_v0_ = 176
                    d_3278_v1_: _dafny.Map
                    d_3278_v1_ = _dafny.Map({self.f21: d_3277_v0_})
                    d_3279_v2_: _dafny.Map
                    d_3279_v2_ = _dafny.Map({d_3277_v0_: (self).f22})
                    d_3278_v1_ = (d_3278_v1_).set((d_3279_v2_) != (d_3279_v2_), d_3277_v0_)
                    d_3280_v3_: _dafny.Array
                    nw526_ = _dafny.Array(False, 17)
                    d_3280_v3_ = nw526_
                    d_3281_v4_: _dafny.Seq
                    d_3281_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cidh"))
                    d_3282_v5_: str
                    d_3282_v5_ = _dafny.CodePoint('x')
                    d_3283_v6_: _dafny.MultiSet
                    d_3283_v6_ = _dafny.MultiSet([d_3282_v5_])
                    d_3284_v7_: C1
                    nw527_ = C1()
                    nw527_.ctor__(d_3280_v3_, (0) - ((len(d_3281_v4_)) * (d_3277_v0_)), d_3277_v0_, d_3277_v0_, default__.fm69((self).f22, d_3277_v0_, d_3283_v6_, 93, globalState), d_3281_v4_)
                    d_3284_v7_ = nw527_
                    if (self).f22:
                        index565_ = default__.safeIndex(69, (d_3280_v3_).length(0))
                        (d_3280_v3_)[index565_] = (self).f22
                        index566_ = default__.safeIndex(69, (d_3280_v3_).length(0))
                        (d_3280_v3_)[index566_] = ((d_3279_v2_)[631] if (631) in (d_3279_v2_) else (self).f22)
                        d_3285_v8_: _dafny.MultiSet
                        d_3285_v8_ = _dafny.MultiSet([self.f21])
                        d_3286_v9_: _dafny.Map
                        d_3286_v9_ = _dafny.Map({(d_3280_v3_)[default__.safeIndex(69, (d_3280_v3_).length(0))]: d_3285_v8_})
                        d_3286_v9_ = ((_dafny.Map({(d_3280_v3_)[default__.safeIndex(69, (d_3280_v3_).length(0))]: d_3285_v8_})) | (d_3286_v9_)) | (d_3286_v9_)
                        d_3287_v10_: _dafny.Map
                        d_3287_v10_ = _dafny.Map({d_3277_v0_: default__.fm47(d_3277_v0_, globalState)})
                        d_3287_v10_ = (d_3287_v10_).set(len(d_3281_v4_), d_3282_v5_)
                        (globalState).f6 = d_3277_v0_
                        index567_ = default__.safeIndex(69, (d_3280_v3_).length(0))
                        (d_3280_v3_)[index567_] = (d_3280_v3_)[default__.safeIndex(69, (d_3280_v3_).length(0))]
                    elif True:
                        d_3288_v11_: _dafny.Set
                        d_3288_v11_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([True])})
                        d_3289_v12_: _dafny.Map
                        d_3289_v12_ = _dafny.Map({False: self.f21})
                        d_3290_v13_: _dafny.Map
                        d_3290_v13_ = _dafny.Map({d_3288_v11_: (False if ((d_3289_v12_)[(self).f22] if ((self).f22) in (d_3289_v12_) else self.f21) else self.f21)})
                        rhs547_ = (d_3284_v7_).fm1((d_3277_v0_) + (d_3277_v0_), self.f21, self.f21, globalState)
                        rhs548_ = ((d_3290_v13_)[(d_3288_v11_).intersection(d_3288_v11_)] if ((d_3288_v11_).intersection(d_3288_v11_)) in (d_3290_v13_) else not((self).f22))
                        lhs460_ = globalState
                        lhs461_ = globalState
                        lhs460_.f7 = rhs547_
                        lhs461_.f2 = rhs548_
                        arr32_ = d_3284_v7_.f24
                        index568_ = default__.safeIndex(853, (d_3284_v7_.f24).length(0))
                        arr32_[index568_] = True
                        d_3291_v14_: _dafny.Map
                        d_3291_v14_ = _dafny.Map({d_3277_v0_: 557})
                        d_3292_v15_: _dafny.Set
                        d_3292_v15_ = _dafny.Set({False})
                        d_3293_v16_: _dafny.Seq
                        d_3293_v16_ = _dafny.SeqWithoutIsStrInference([len(d_3292_v15_)])
                        d_3294_v17_: _dafny.Set
                        d_3294_v17_ = _dafny.Set({d_3277_v0_})
                        d_3295_v18_: _dafny.Array
                        nw528_ = _dafny.Array(None, 22)
                        nw528_[int(0)] = d_3277_v0_
                        nw528_[int(1)] = d_3277_v0_
                        nw528_[int(2)] = d_3277_v0_
                        nw528_[int(3)] = d_3277_v0_
                        nw528_[int(4)] = d_3277_v0_
                        nw528_[int(5)] = d_3277_v0_
                        nw528_[int(6)] = d_3277_v0_
                        nw528_[int(7)] = len(d_3281_v4_)
                        nw528_[int(8)] = len(d_3291_v14_)
                        nw528_[int(9)] = len(default__.fm44(len(d_3293_v16_), (self).f22, globalState))
                        nw528_[int(10)] = (0) - (d_3277_v0_)
                        nw528_[int(11)] = -787
                        nw528_[int(12)] = d_3277_v0_
                        nw528_[int(13)] = d_3277_v0_
                        nw528_[int(14)] = len(d_3294_v17_)
                        nw528_[int(15)] = d_3277_v0_
                        nw528_[int(16)] = d_3277_v0_
                        nw528_[int(17)] = len(d_3293_v16_)
                        nw528_[int(18)] = len(d_3292_v15_)
                        nw528_[int(19)] = d_3277_v0_
                        nw528_[int(20)] = d_3277_v0_
                        nw528_[int(21)] = d_3277_v0_
                        d_3295_v18_ = nw528_
                        d_3296_v19_: D19
                        d_3296_v19_ = D19_DC56(d_3295_v18_, True, d_3277_v0_, _dafny.SeqWithoutIsStrInference([d_3289_v12_, d_3289_v12_, d_3289_v12_]))
                        d_3297_v20_: D20
                        d_3297_v20_ = D20_DC58((d_3296_v19_).cf106)
                        d_3298_v21_: D20
                        d_3298_v21_ = D20_DC59(d_3297_v20_)
                        d_3299_v22_: C3
                        nw529_ = C3()
                        nw529_.ctor__((D16_DC45(d_3284_v7_.f24, _dafny.CodePoint('p'), d_3277_v0_, d_3277_v0_, (self).f22)).cf85, d_3281_v4_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgjk")), d_3277_v0_, (self).f22, d_3281_v4_, d_3277_v0_, len(d_3293_v16_))
                        d_3299_v22_ = nw529_
                        d_3300_v23_: _dafny.Seq
                        d_3300_v23_ = _dafny.SeqWithoutIsStrInference([d_3299_v22_])
                        d_3301_v24_: _dafny.Map
                        d_3301_v24_ = _dafny.Map({d_3281_v4_: d_3299_v22_})
                        d_3302_v25_: D24
                        d_3302_v25_ = D24_DC65(d_3299_v22_)
                        pat_let_tv55_ = d_3299_v22_
                        d_3303_v26_: _dafny.Array
                        nw530_ = _dafny.Array(None, 19)
                        nw530_[int(0)] = d_3299_v22_
                        nw530_[int(1)] = d_3299_v22_
                        nw530_[int(2)] = d_3299_v22_
                        nw530_[int(3)] = d_3299_v22_
                        nw530_[int(4)] = d_3299_v22_
                        nw530_[int(5)] = d_3299_v22_
                        nw530_[int(6)] = d_3299_v22_
                        nw530_[int(7)] = d_3299_v22_
                        nw530_[int(8)] = ((d_3300_v23_)[default__.safeIndex(d_3277_v0_, len(d_3300_v23_))] if (self).f22 else d_3299_v22_)
                        nw530_[int(9)] = ((d_3301_v24_)[_dafny.SeqWithoutIsStrInference([d_3282_v5_ for d_3304_i1_ in range(default__.abs(569))])] if (_dafny.SeqWithoutIsStrInference([d_3282_v5_ for d_3305_i1_ in range(default__.abs(569))])) in (d_3301_v24_) else d_3299_v22_)
                        nw530_[int(10)] = d_3299_v22_
                        nw530_[int(11)] = d_3299_v22_
                        nw530_[int(12)] = d_3299_v22_
                        nw530_[int(13)] = d_3299_v22_
                        nw530_[int(14)] = d_3299_v22_
                        nw530_[int(15)] = d_3299_v22_
                        nw530_[int(16)] = d_3299_v22_
                        def iife199_(_pat_let62_0):
                            def iife200_(d_3306_dt__update__tmp_h0_):
                                def iife201_(_pat_let63_0):
                                    def iife202_(d_3307_dt__update_hcf122_h0_):
                                        return D24_DC65(d_3307_dt__update_hcf122_h0_)
                                    return iife202_(_pat_let63_0)
                                return iife201_(pat_let_tv55_)
                            return iife200_(_pat_let62_0)
                        nw530_[int(17)] = (iife199_(d_3302_v25_)).cf122
                        nw530_[int(18)] = d_3299_v22_
                        d_3303_v26_ = nw530_
                        index569_ = default__.safeIndex(280, (d_3303_v26_).length(0))
                        (d_3303_v26_)[index569_] = d_3299_v22_
                        d_3308_v27_: _dafny.Array
                        nw531_ = _dafny.Array(_dafny.CodePoint('D'), 18)
                        d_3308_v27_ = nw531_
                        index570_ = default__.safeIndex(796, (d_3308_v27_).length(0))
                        (d_3308_v27_)[index570_] = default__.fm14((d_3299_v22_).f25, d_3282_v5_, d_3277_v0_, (self).f22, globalState)
                        d_3309_v28_: D10
                        d_3309_v28_ = D10_DC29()
                        arr33_ = d_3284_v7_.f24
                        index571_ = default__.safeIndex(853, (d_3284_v7_.f24).length(0))
                        index572_ = default__.safeIndex(280, (d_3303_v26_).length(0))
                        index573_ = default__.safeIndex(796, (d_3308_v27_).length(0))
                        rhs549_ = (self).f22
                        rhs550_ = default__.fm72(d_3309_v28_, d_3277_v0_, default__.safeModuloInt(d_3277_v0_, d_3277_v0_), d_3279_v2_, globalState)
                        rhs551_ = d_3299_v22_
                        rhs552_ = d_3282_v5_
                        rhs553_ = d_3281_v4_
                        lhs462_ = d_3284_v7_.f24
                        lhs463_ = default__.safeIndex(853, (d_3284_v7_.f24).length(0))
                        lhs464_ = d_3303_v26_
                        lhs465_ = default__.safeIndex(280, (d_3303_v26_).length(0))
                        lhs466_ = d_3308_v27_
                        lhs467_ = default__.safeIndex(796, (d_3308_v27_).length(0))
                        lhs462_[lhs463_] = rhs549_
                        d_3298_v21_ = rhs550_
                        lhs464_[lhs465_] = rhs551_
                        lhs466_[lhs467_] = rhs552_
                        d_3281_v4_ = rhs553_
                        d_3293_v16_ = (d_3293_v16_) + (d_3293_v16_)
                        index574_ = default__.safeIndex(37, (d_3295_v18_).length(0))
                        (d_3295_v18_)[index574_] = d_3277_v0_
                        index575_ = default__.safeIndex(37, (d_3295_v18_).length(0))
                        (d_3295_v18_)[index575_] = d_3277_v0_
                        arr34_ = d_3284_v7_.f24
                        index576_ = default__.safeIndex(853, (d_3284_v7_.f24).length(0))
                        arr34_[index576_] = self.f21
                    d_3310_v29_: _dafny.Seq
                    d_3310_v29_ = _dafny.SeqWithoutIsStrInference([d_3279_v2_])
                    d_3311_v30_: C6
                    nw532_ = C6()
                    nw532_.ctor__((d_3310_v29_) + (d_3310_v29_), d_3277_v0_, (self).f22, d_3281_v4_, d_3277_v0_, d_3277_v0_)
                    d_3311_v30_ = nw532_
                    pass
            pass
        d_3312_v31_: _dafny.Seq
        d_3312_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bife"))
        d_3313_v32_: int
        d_3313_v32_ = 293
        hi12_ = d_3313_v32_
        for d_3314_i2_ in range(len(d_3312_v31_), hi12_):
            d_3315_v33_: _dafny.Seq
            d_3315_v33_ = _dafny.SeqWithoutIsStrInference([self.f21, self.f21])
            d_3316_v34_: D6
            d_3316_v34_ = D6_DC21((self).f22)
            d_3317_v35_: _dafny.Array
            nw533_ = _dafny.Array(None, 16)
            nw533_[int(0)] = self.f21
            nw533_[int(1)] = (d_3315_v33_)[default__.safeIndex(default__.fm0((self).f22, globalState), len(d_3315_v33_))]
            nw533_[int(2)] = self.f21
            nw533_[int(3)] = (self).f22
            nw533_[int(4)] = not(self.f21)
            nw533_[int(5)] = (self).f22
            nw533_[int(6)] = self.f21
            nw533_[int(7)] = self.f21
            nw533_[int(8)] = (d_3315_v33_)[default__.safeIndex((0) - (d_3314_i2_), len(d_3315_v33_))]
            nw533_[int(9)] = (self).f22
            nw533_[int(10)] = self.f21
            nw533_[int(11)] = (d_3316_v34_).cf40
            nw533_[int(12)] = self.f21
            nw533_[int(13)] = self.f21
            nw533_[int(14)] = not(self.f21)
            nw533_[int(15)] = self.f21
            d_3317_v35_ = nw533_
            d_3318_v36_: _dafny.Map
            d_3318_v36_ = _dafny.Map({d_3313_v32_: 299})
            d_3319_v37_: T2
            nw534_ = C1()
            nw534_.ctor__(d_3317_v35_, len(d_3318_v36_), d_3314_i2_, 824, self.f21, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_3320_i3_ in range(default__.abs(477))]))
            d_3319_v37_ = nw534_
            d_3321_v38_: _dafny.Map
            d_3321_v38_ = _dafny.Map({d_3319_v37_: (d_3315_v33_) + (d_3315_v33_)})
            d_3321_v38_ = (d_3321_v38_).set(d_3319_v37_, _dafny.SeqWithoutIsStrInference([(self).f22, (self).f22]))
            d_3317_v35_ = d_3317_v35_
            d_3322_v39_: str
            d_3322_v39_ = _dafny.CodePoint('g')
            d_3323_v40_: _dafny.Map
            d_3323_v40_ = _dafny.Map({(d_3319_v37_).f12: d_3322_v39_})
            d_3324_v42_: _dafny.Set
            d_3324_v42_ = _dafny.Set({(d_3319_v37_).f12})
            d_3325_v43_: _dafny.Array
            nw535_ = _dafny.Array(None, 4)
            nw535_[int(0)] = (d_3323_v40_) | (d_3323_v40_)
            nw535_[int(1)] = d_3323_v40_
            def iife203_():
                coll75_ = _dafny.Map()
                compr_75_: int
                for compr_75_ in (d_3324_v42_).Elements:
                    d_3326_v41_: int = compr_75_
                    if (d_3326_v41_) in (d_3324_v42_):
                        coll75_[(d_3326_v41_) - ((d_3319_v37_).f14)] = d_3322_v39_
                return _dafny.Map(coll75_)
            nw535_[int(2)] = (d_3323_v40_) | (iife203_()
            )
            nw535_[int(3)] = _dafny.Map({(d_3319_v37_).f14: d_3322_v39_})
            d_3325_v43_ = nw535_
            d_3325_v43_ = d_3325_v43_
            d_3327_v44_: _dafny.Array
            nw536_ = _dafny.Array(int(0), 23)
            d_3327_v44_ = nw536_
            d_3328_v45_: _dafny.Map
            d_3328_v45_ = _dafny.Map({(d_3319_v37_).f15: (d_3319_v37_).f15})
            index577_ = default__.safeIndex(489, (d_3327_v44_).length(0))
            (d_3327_v44_)[index577_] = (0) - (default__.safeDivisionInt(len((D25_DC67(d_3328_v45_)).cf125), (d_3319_v37_).f11))
            index578_ = default__.safeIndex(489, (d_3327_v44_).length(0))
            (d_3327_v44_)[index578_] = d_3313_v32_
        (globalState).f6 = default__.fm0((self).f22, globalState)
        d_3329_v46_: _dafny.MultiSet
        d_3329_v46_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f22, self.f21, self.f21])])
        d_3330_v47_: _dafny.Seq
        d_3330_v47_ = _dafny.SeqWithoutIsStrInference([self.f21])
        d_3331_v48_: _dafny.Array
        nw537_ = _dafny.Array(None, 9)
        nw537_[int(0)] = d_3329_v46_
        nw537_[int(1)] = d_3329_v46_
        nw537_[int(2)] = d_3329_v46_
        nw537_[int(3)] = (default__.fm73(d_3313_v32_, 235, globalState)) - (d_3329_v46_)
        nw537_[int(4)] = _dafny.MultiSet([d_3330_v47_])
        nw537_[int(5)] = (d_3329_v46_).intersection(d_3329_v46_)
        nw537_[int(6)] = d_3329_v46_
        nw537_[int(7)] = (d_3329_v46_) - (d_3329_v46_)
        nw537_[int(8)] = d_3329_v46_
        d_3331_v48_ = nw537_
        index579_ = default__.safeIndex(65, (d_3331_v48_).length(0))
        (d_3331_v48_)[index579_] = d_3329_v46_
        pat_let_tv56_ = d_3329_v46_
        pat_let_tv57_ = d_3329_v46_
        pat_let_tv58_ = d_3329_v46_
        index580_ = default__.safeIndex(65, (d_3331_v48_).length(0))
        def lambda132_(source45_):
            if source45_.is_DC3:
                d_3332___mcc_h0_ = source45_.cf4
                d_3333_cf4_ = d_3332___mcc_h0_
                return pat_let_tv56_
            elif source45_.is_DC4:
                d_3334___mcc_h1_ = source45_.cf5
                d_3335_cf5_ = d_3334___mcc_h1_
                return pat_let_tv57_
            elif source45_.is_DC5:
                d_3336___mcc_h2_ = source45_.cf6
                d_3337___mcc_h3_ = source45_.cf7
                d_3338___mcc_h4_ = source45_.cf8
                d_3339_cf8_ = d_3338___mcc_h4_
                d_3340_cf7_ = d_3337___mcc_h3_
                d_3341_cf6_ = d_3336___mcc_h2_
                return pat_let_tv58_
            elif True:
                d_3342___mcc_h5_ = source45_.cf3
                d_3343_cf3_ = d_3342___mcc_h5_
                return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f22])])

        (d_3331_v48_)[index580_] = lambda132_(D1_DC4(self.f21))
        d_3344_v49_: _dafny.Map
        d_3344_v49_ = _dafny.Map({d_3312_v31_: d_3313_v32_})
        d_3345_v50_: D3
        d_3345_v50_ = D3_DC10((self).f22, (0) - (d_3313_v32_), d_3344_v49_, d_3312_v31_, d_3313_v32_)
        pat_let_tv59_ = d_3313_v32_
        pat_let_tv60_ = d_3313_v32_
        def lambda133_(source46_):
            if source46_.is_DC10:
                d_3346___mcc_h6_ = source46_.cf19
                d_3347___mcc_h7_ = source46_.cf20
                d_3348___mcc_h8_ = source46_.cf21
                d_3349___mcc_h9_ = source46_.cf22
                d_3350___mcc_h10_ = source46_.cf23
                d_3351_cf23_ = d_3350___mcc_h10_
                d_3352_cf22_ = d_3349___mcc_h9_
                d_3353_cf21_ = d_3348___mcc_h8_
                d_3354_cf20_ = d_3347___mcc_h7_
                d_3355_cf19_ = d_3346___mcc_h6_
                return default__.safeModuloInt(d_3351_cf23_, d_3354_cf20_)
            elif source46_.is_DC11:
                d_3356___mcc_h11_ = source46_.cf24
                d_3357___mcc_h12_ = source46_.cf25
                d_3358___mcc_h13_ = source46_.cf26
                d_3359___mcc_h14_ = source46_.cf27
                d_3360_cf27_ = d_3359___mcc_h14_
                d_3361_cf26_ = d_3358___mcc_h13_
                d_3362_cf25_ = d_3357___mcc_h12_
                d_3363_cf24_ = d_3356___mcc_h11_
                return pat_let_tv59_
            elif source46_.is_DC9:
                d_3364___mcc_h15_ = source46_.cf18
                d_3365_cf18_ = d_3364___mcc_h15_
                return pat_let_tv60_
            elif True:
                d_3366___mcc_h16_ = source46_.cf28
                d_3367_cf28_ = d_3366___mcc_h16_
                return 735

        (globalState).f6 = lambda133_(d_3345_v50_)
        d_3368_v51_: _dafny.Array
        nw538_ = _dafny.Array(int(0), 13)
        d_3368_v51_ = nw538_
        index581_ = default__.safeIndex(251, (d_3368_v51_).length(0))
        (d_3368_v51_)[index581_] = d_3313_v32_
        d_3369_v52_: _dafny.MultiSet
        d_3369_v52_ = _dafny.MultiSet([d_3313_v32_])
        index582_ = default__.safeIndex(251, (d_3368_v51_).length(0))
        (d_3368_v51_)[index582_] = (0) - (((_dafny.MultiSet([default__.fm0(self.f21, globalState)])) | ((d_3369_v52_).intersection(d_3369_v52_))).cardinality)
        d_3370_v53_: _dafny.Seq
        d_3370_v53_ = _dafny.SeqWithoutIsStrInference([d_3313_v32_, (d_3368_v51_)[default__.safeIndex(251, (d_3368_v51_).length(0))], 564])
        r0 = (_dafny.SeqWithoutIsStrInference([d_3313_v32_ for d_3371_i4_ in range(default__.abs(537))])) < (d_3370_v53_)
        d_3372_v54_: D1
        d_3372_v54_ = D1_DC5(203, self.f21, (self).f22)
        r1 = (d_3372_v54_).cf7
        r2 = self.f21
        return r0, r1, r2

    @property
    def f22(self):
        return self._f22

class C9(T0):
    def  __init__(self):
        self._f11: int = int(0)
        self._f12: int = int(0)
        self._f19: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f20: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12
    def ctor__(self, f19, f20, f11, f12):
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f11 = f11
        (self)._f12 = f12

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        d_3373_v0_: str
        d_3373_v0_ = _dafny.CodePoint('v')
        d_3374_v1_: _dafny.Seq
        d_3374_v1_ = _dafny.SeqWithoutIsStrInference([((self).f19).set(default__.safeIndex((self).f12, len((self).f19)), d_3373_v0_), _dafny.SeqWithoutIsStrInference([d_3373_v0_])])
        d_3375_v2_: _dafny.Array
        nw539_ = _dafny.Array(None, 28)
        nw539_[int(0)] = ((self).f19 if p2 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')]))
        nw539_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_3376_i0_ in range(default__.abs(101))])
        nw539_[int(2)] = ((p0) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_3377_i1_ in range(default__.abs(-302))]))).set(default__.safeIndex(707, len((p0) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_3378_i1_ in range(default__.abs(-302))])))), _dafny.CodePoint('x'))
        nw539_[int(3)] = p0
        nw539_[int(4)] = p0
        nw539_[int(5)] = (_dafny.SeqWithoutIsStrInference([((self).f19)[default__.safeIndex(p3, len((self).f19))] for d_3379_i2_ in range(default__.abs(514))])).set(default__.safeIndex(len((self).f19), len(_dafny.SeqWithoutIsStrInference([((self).f19)[default__.safeIndex(p3, len((self).f19))] for d_3380_i2_ in range(default__.abs(514))]))), d_3373_v0_)
        nw539_[int(6)] = (d_3374_v1_)[default__.safeIndex((self).f12, len(d_3374_v1_))]
        nw539_[int(7)] = (self).f19
        nw539_[int(8)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u'), d_3373_v0_, d_3373_v0_])) + (_dafny.SeqWithoutIsStrInference([d_3373_v0_ for d_3381_i3_ in range(default__.abs(62))]))
        nw539_[int(9)] = ((d_3374_v1_)[default__.safeIndex((self).f11, len(d_3374_v1_))]) + (p0)
        nw539_[int(10)] = (self).f19
        nw539_[int(11)] = (self).f19
        nw539_[int(12)] = (p0) + ((self).f19)
        nw539_[int(13)] = ((self).f19) + ((self).f19)
        nw539_[int(14)] = p0
        nw539_[int(15)] = (self).f19
        nw539_[int(16)] = _dafny.SeqWithoutIsStrInference([d_3373_v0_ for d_3382_i4_ in range(default__.abs(253))])
        nw539_[int(17)] = (self).f19
        nw539_[int(18)] = p0
        nw539_[int(19)] = _dafny.SeqWithoutIsStrInference([d_3373_v0_, d_3373_v0_, d_3373_v0_])
        nw539_[int(20)] = (_dafny.SeqWithoutIsStrInference([d_3373_v0_])) + ((self).f19)
        nw539_[int(21)] = _dafny.SeqWithoutIsStrInference([d_3373_v0_, d_3373_v0_])
        nw539_[int(22)] = _dafny.SeqWithoutIsStrInference([d_3373_v0_])
        nw539_[int(23)] = (self).f19
        nw539_[int(24)] = (self).f19
        nw539_[int(25)] = ((self).f19) + (p0)
        nw539_[int(26)] = p0
        nw539_[int(27)] = (p0 if (self).f20 else (self).f19)
        d_3375_v2_ = nw539_
        d_3375_v2_ = d_3375_v2_
        d_3383_v3_: _dafny.Array
        def lambda134_(d_3384_i6_):
            return (self).f20

        init79_ = lambda134_
        nw540_ = _dafny.Array(None, 20)
        for i0_79_ in range(nw540_.length(0)):
            nw540_[i0_79_] = init79_(i0_79_)
        d_3383_v3_ = nw540_
        guard_loop_11_: int
        for guard_loop_11_ in _dafny.IntegerRange(0, (d_3383_v3_).length(0)):
            d_3385_i5_: int = guard_loop_11_
            if (True) and (((0) <= (d_3385_i5_)) and ((d_3385_i5_) < ((d_3383_v3_).length(0)))):
                (d_3383_v3_)[(d_3385_i5_)] = (self).f20
        d_3386_v4_: _dafny.Seq
        d_3386_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncq"))
        d_3387_v5_: _dafny.Seq
        d_3387_v5_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({p2})), p1, p3, (0) - ((self).f12), len(_dafny.Map({(self).f20: (self).f20}))])
        d_3388_v6_: _dafny.Map
        d_3388_v6_ = _dafny.Map({p2: p2})
        d_3389_v7_: _dafny.Map
        d_3389_v7_ = _dafny.Map({(d_3383_v3_ if p2 else d_3383_v3_): ((d_3388_v6_)[((d_3388_v6_)[(self).f20] if ((self).f20) in (d_3388_v6_) else (self).f20)] if (((d_3388_v6_)[(self).f20] if ((self).f20) in (d_3388_v6_) else (self).f20)) in (d_3388_v6_) else p2)})
        rhs554_ = ((d_3387_v5_)[default__.safeIndex((self).f11, len(d_3387_v5_))]) == (default__.safeDivisionInt(p1, p1))
        rhs555_ = ((d_3389_v7_)[d_3383_v3_] if (d_3383_v3_) in (d_3389_v7_) else p2)
        rhs556_ = _dafny.SeqWithoutIsStrInference([d_3373_v0_ for d_3390_i7_ in range(default__.abs(149))])
        lhs468_ = globalState
        lhs469_ = globalState
        lhs468_.f2 = rhs554_
        lhs469_.f7 = rhs555_
        d_3386_v4_ = rhs556_
        d_3391_v8_: _dafny.Seq
        d_3391_v8_ = _dafny.SeqWithoutIsStrInference([(self).f20, (self).f20, p2, (self).f20])
        d_3392_i8_: int
        d_3392_i8_ = 0
        with _dafny.label("17"):
            while (p3) < (len(d_3391_v8_)):
                with _dafny.c_label("17"):
                    if (d_3392_i8_) >= (100):
                        raise _dafny.Break("17")
                    d_3392_i8_ = (d_3392_i8_) + (1)
                    d_3393_v9_: _dafny.Array
                    nw541_ = _dafny.Array(None, 6)
                    d_3393_v9_ = nw541_
                    d_3394_v10_: _dafny.Set
                    d_3394_v10_ = _dafny.Set({d_3393_v9_, d_3393_v9_, d_3393_v9_, d_3393_v9_})
                    d_3395_v11_: D2
                    d_3395_v11_ = D2_DC7(d_3383_v3_, d_3394_v10_, D1_DC2(d_3373_v0_), p1, len(d_3391_v8_))
                    rhs557_ = (self).f20
                    rhs558_ = (self).f20
                    rhs559_ = ((d_3395_v11_).cf10 if (40) < (len(d_3387_v5_)) else d_3383_v3_)
                    lhs470_ = globalState
                    lhs470_.f2 = rhs557_
                    r0 = rhs558_
                    d_3383_v3_ = rhs559_
                    d_3396_v12_: _dafny.Map
                    d_3396_v12_ = _dafny.Map({(self).f11: not(p2)})
                    d_3397_v13_: _dafny.Seq
                    d_3397_v13_ = _dafny.SeqWithoutIsStrInference([d_3396_v12_])
                    d_3398_v14_: C6
                    nw542_ = C6()
                    nw542_.ctor__((d_3397_v13_) + (_dafny.SeqWithoutIsStrInference([d_3396_v12_])), (p1) + ((self).f11), p2, p0, (self).f11, len(p0))
                    d_3398_v14_ = nw542_
                    (globalState).f6 = len(((self).f19) + (d_3386_v4_))
                    d_3399_v15_: _dafny.MultiSet
                    d_3399_v15_ = _dafny.MultiSet([((d_3396_v12_)[(self).f12] if ((self).f12) in (d_3396_v12_) else (self).f20)])
                    d_3400_v16_: C3
                    nw543_ = C3()
                    nw543_.ctor__((self).f20, (p0) + ((self).f19), default__.fm44(646, (self).f20, globalState), (d_3399_v15_).cardinality, p2, d_3386_v4_, p3, (self).f11)
                    d_3400_v16_ = nw543_
                    pass
            pass
        d_3401_v17_: _dafny.Set
        d_3401_v17_ = _dafny.Set({p2})
        r2 = len(_dafny.SeqWithoutIsStrInference([_dafny.Set({p2}), _dafny.Set({p2, (self).f20, (self).f20, p2, p2}), (d_3401_v17_) - (d_3401_v17_)]))
        d_3402_v18_: _dafny.MultiSet
        d_3402_v18_ = _dafny.MultiSet([not(p2), False])
        d_3403_v19_: T2
        nw544_ = C1()
        nw544_.ctor__(d_3383_v3_, p1, len(_dafny.Map({(self).f20: 32})), (0) - ((self).f11), p2, d_3386_v4_)
        d_3403_v19_ = nw544_
        d_3404_v20_: _dafny.Map
        d_3404_v20_ = _dafny.Map({d_3403_v19_: -797})
        r2 = ((d_3402_v18_)[(d_3403_v19_) in (d_3404_v20_)] if ((d_3403_v19_) in (d_3404_v20_)) in (d_3402_v18_) else p3)
        r0 = (self).f20
        r1 = len((d_3403_v19_).f13)
        r2 = (0) - ((d_3403_v19_).f12)
        return r0, r1, r2

    def m6(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        r2: bool = False
        r3: int = int(0)
        d_3405_v0_: _dafny.Set
        d_3405_v0_ = _dafny.Set({p2})
        d_3406_v1_: _dafny.Seq
        d_3406_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f20})])
        d_3407_v2_: _dafny.Seq
        d_3407_v2_ = _dafny.SeqWithoutIsStrInference([(d_3406_v1_)[default__.safeIndex(p3, len(d_3406_v1_))], _dafny.Set({(self).f20}), d_3405_v0_])
        (globalState).f2 = ((d_3405_v0_).intersection(d_3405_v0_)) in (d_3406_v1_)
        d_3408_v3_: _dafny.Array
        nw545_ = _dafny.Array(None, 28)
        d_3408_v3_ = nw545_
        d_3409_v4_: D4
        d_3409_v4_ = D4_DC14(_dafny.CodePoint('w'), p0)
        d_3410_v5_: _dafny.Set
        d_3410_v5_ = _dafny.Set({d_3409_v4_})
        d_3411_v6_: D2
        d_3411_v6_ = D2_DC8(p2, p2, (self).f19)
        d_3412_v7_: _dafny.Map
        d_3412_v7_ = _dafny.Map({p2: not(default__.fm7((self).f20, (self).f12, globalState))})
        d_3413_v8_: _dafny.Map
        d_3413_v8_ = _dafny.Map({p2: d_3412_v7_})
        d_3414_v9_: D15
        d_3414_v9_ = D15_DC42(d_3413_v8_)
        pat_let_tv61_ = p2
        pat_let_tv62_ = p2
        pat_let_tv63_ = p2
        pat_let_tv64_ = p2
        pat_let_tv65_ = p2
        def lambda135_(source47_):
            if source47_.is_DC43:
                d_3415___mcc_h0_ = source47_.cf77
                d_3416___mcc_h1_ = source47_.cf78
                d_3417___mcc_h2_ = source47_.cf79
                d_3418_cf79_ = d_3417___mcc_h2_
                d_3419_cf78_ = d_3416___mcc_h1_
                d_3420_cf77_ = d_3415___mcc_h0_
                return (_dafny.MultiSet([True])).issubset(_dafny.MultiSet([not(pat_let_tv61_), pat_let_tv62_]))
            elif True:
                d_3421___mcc_h3_ = source47_.cf76
                d_3422_cf76_ = d_3421___mcc_h3_
                return (_dafny.MultiSet([pat_let_tv63_, (D14_DC41((self).f20, pat_let_tv64_, (self).f19)).cf74])).ispropersubset(_dafny.MultiSet([True, pat_let_tv65_]))

        rhs560_ = (self).f11
        rhs561_ = len((d_3410_v5_) - (d_3410_v5_))
        rhs562_ = (d_3411_v6_).cf16
        rhs563_ = not(lambda135_(d_3414_v9_))
        rhs564_ = d_3408_v3_
        lhs471_ = globalState
        r3 = rhs560_
        r3 = rhs561_
        lhs471_.f2 = rhs562_
        r1 = rhs563_
        d_3408_v3_ = rhs564_
        d_3423_v10_: _dafny.Seq
        d_3423_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tcvc"))
        d_3423_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjdwohj"))
        d_3424_i0_: int
        d_3424_i0_ = 0
        with _dafny.label("18"):
            while not((self).f20):
                with _dafny.c_label("18"):
                    if (d_3424_i0_) >= (100):
                        raise _dafny.Break("18")
                    d_3424_i0_ = (d_3424_i0_) + (1)
                    d_3425_v11_: _dafny.MultiSet
                    d_3425_v11_ = _dafny.MultiSet([(self).f12])
                    d_3426_v12_: _dafny.Set
                    d_3426_v12_ = _dafny.Set({(self).f12, p3})
                    d_3427_v13_: _dafny.Array
                    nw546_ = _dafny.Array(None, 13)
                    nw546_[int(0)] = 924
                    nw546_[int(1)] = default__.fm0((self).f20, globalState)
                    nw546_[int(2)] = p3
                    nw546_[int(3)] = p0
                    nw546_[int(4)] = (self).f11
                    nw546_[int(5)] = ((d_3425_v11_).set(len(_dafny.Map({726: p2})), default__.abs(p0))).cardinality
                    nw546_[int(6)] = (self).f11
                    nw546_[int(7)] = len((self).f19)
                    nw546_[int(8)] = (self).f12
                    nw546_[int(9)] = p3
                    nw546_[int(10)] = len(d_3426_v12_)
                    nw546_[int(11)] = len(d_3423_v10_)
                    nw546_[int(12)] = (self).f12
                    d_3427_v13_ = nw546_
                    d_3428_v14_: _dafny.Map
                    d_3428_v14_ = _dafny.Map({(self).f11: d_3427_v13_})
                    d_3429_v15_: C2
                    nw547_ = C2()
                    nw547_.ctor__((self).f12, d_3423_v10_, len(_dafny.SeqWithoutIsStrInference([p3 for d_3430_i1_ in range(default__.abs(984))])), p3)
                    d_3429_v15_ = nw547_
                    d_3431_v16_: D18
                    d_3431_v16_ = D18_DC53(p0, d_3426_v12_, (0) - (p3), d_3429_v15_)
                    d_3432_v17_: D0
                    d_3432_v17_ = D0_DC1(p2, 886)
                    pat_let_tv66_ = d_3426_v12_
                    pat_let_tv67_ = d_3429_v15_
                    d_3433_v18_: _dafny.Array
                    nw548_ = _dafny.Array(None, 20)
                    nw548_[int(0)] = D18_DC53(len(d_3428_v14_), d_3426_v12_, p0, d_3429_v15_)
                    nw548_[int(1)] = d_3431_v16_
                    nw548_[int(2)] = d_3431_v16_
                    nw548_[int(3)] = (d_3431_v16_ if True else d_3431_v16_)
                    nw548_[int(4)] = d_3431_v16_
                    nw548_[int(5)] = d_3431_v16_
                    nw548_[int(6)] = (d_3431_v16_ if (d_3432_v17_).cf1 else d_3431_v16_)
                    nw548_[int(7)] = d_3431_v16_
                    def iife204_(_pat_let64_0):
                        def iife205_(d_3434_dt__update__tmp_h0_):
                            def iife206_(_pat_let65_0):
                                def iife207_(d_3435_dt__update_hcf99_h0_):
                                    def iife208_(_pat_let66_0):
                                        def iife209_(d_3436_dt__update_hcf100_h0_):
                                            return D18_DC53(d_3435_dt__update_hcf99_h0_, d_3436_dt__update_hcf100_h0_, (d_3434_dt__update__tmp_h0_).cf101, (d_3434_dt__update__tmp_h0_).cf102)
                                        return iife209_(_pat_let66_0)
                                    return iife208_(pat_let_tv66_)
                                return iife207_(_pat_let65_0)
                            return iife206_((self).f12)
                        return iife205_(_pat_let64_0)
                    nw548_[int(8)] = iife204_(d_3431_v16_)
                    nw548_[int(9)] = d_3431_v16_
                    nw548_[int(10)] = d_3431_v16_
                    nw548_[int(11)] = d_3431_v16_
                    nw548_[int(12)] = d_3431_v16_
                    nw548_[int(13)] = d_3431_v16_
                    nw548_[int(14)] = d_3431_v16_
                    def iife210_(_pat_let67_0):
                        def iife211_(d_3437_dt__update__tmp_h1_):
                            def iife212_(_pat_let68_0):
                                def iife213_(d_3438_dt__update_hcf102_h0_):
                                    return D18_DC53((d_3437_dt__update__tmp_h1_).cf99, (d_3437_dt__update__tmp_h1_).cf100, (d_3437_dt__update__tmp_h1_).cf101, d_3438_dt__update_hcf102_h0_)
                                return iife213_(_pat_let68_0)
                            return iife212_(pat_let_tv67_)
                        return iife211_(_pat_let67_0)
                    nw548_[int(15)] = iife210_(D18_DC53(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmmc"))), d_3426_v12_, len((d_3423_v10_).set(default__.safeIndex(p0, len(d_3423_v10_)), p1)), d_3429_v15_))
                    nw548_[int(16)] = (d_3431_v16_ if False else d_3431_v16_)
                    nw548_[int(17)] = d_3431_v16_
                    nw548_[int(18)] = d_3431_v16_
                    nw548_[int(19)] = d_3431_v16_
                    d_3433_v18_ = nw548_
                    index583_ = default__.safeIndex(338, (d_3433_v18_).length(0))
                    (d_3433_v18_)[index583_] = d_3431_v16_
                    d_3439_v19_: _dafny.Seq
                    d_3439_v19_ = _dafny.SeqWithoutIsStrInference([p3, p3])
                    pat_let_tv68_ = d_3429_v15_
                    pat_let_tv69_ = p0
                    index584_ = default__.safeIndex(338, (d_3433_v18_).length(0))
                    def iife214_(_pat_let69_0):
                        def iife215_(d_3440_dt__update__tmp_h2_):
                            def iife216_(_pat_let70_0):
                                def iife217_(d_3441_dt__update_hcf102_h1_):
                                    def iife218_(_pat_let71_0):
                                        def iife219_(d_3442_dt__update_hcf100_h1_):
                                            return D18_DC53((d_3440_dt__update__tmp_h2_).cf99, d_3442_dt__update_hcf100_h1_, (d_3440_dt__update__tmp_h2_).cf101, d_3441_dt__update_hcf102_h1_)
                                        return iife219_(_pat_let71_0)
                                    return iife218_(_dafny.Set({pat_let_tv69_}))
                                return iife217_(_pat_let70_0)
                            return iife216_(pat_let_tv68_)
                        return iife215_(_pat_let69_0)
                    rhs565_ = p3
                    rhs566_ = (0) - ((0) - ((len((d_3439_v19_) + (_dafny.SeqWithoutIsStrInference([p3])))) * (d_3429_v15_.f27)))
                    rhs567_ = iife214_(d_3431_v16_)
                    rhs568_ = (55) != (503)
                    lhs472_ = globalState
                    lhs473_ = d_3433_v18_
                    lhs474_ = default__.safeIndex(338, (d_3433_v18_).length(0))
                    lhs472_.f6 = rhs565_
                    r3 = rhs566_
                    lhs473_[lhs474_] = rhs567_
                    r1 = rhs568_
                    (globalState).f6 = (0) - (d_3429_v15_.f27)
                    d_3443_v20_: _dafny.Array
                    def lambda136_(d_3444_i2_):
                        return (self).f20

                    init80_ = lambda136_
                    nw549_ = _dafny.Array(None, 27)
                    for i0_80_ in range(nw549_.length(0)):
                        nw549_[i0_80_] = init80_(i0_80_)
                    d_3443_v20_ = nw549_
                    d_3445_v21_: _dafny.Map
                    d_3445_v21_ = _dafny.Map({-48: d_3443_v20_})
                    d_3445_v21_ = (d_3445_v21_).set((d_3425_v11_).cardinality, d_3443_v20_)
                    source48_ = d_3411_v6_
                    if source48_.is_DC7:
                        d_3446___mcc_h4_ = source48_.cf10
                        d_3447___mcc_h5_ = source48_.cf11
                        d_3448___mcc_h6_ = source48_.cf12
                        d_3449___mcc_h7_ = source48_.cf13
                        d_3450___mcc_h8_ = source48_.cf14
                        d_3451_cf14_ = d_3450___mcc_h8_
                        d_3452_cf13_ = d_3449___mcc_h7_
                        d_3453_cf12_ = d_3448___mcc_h6_
                        d_3454_cf11_ = d_3447___mcc_h5_
                        d_3455_cf10_ = d_3446___mcc_h4_
                        (globalState).f6 = len(default__.fm18((self).f20, (self).f12, globalState))
                        d_3423_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlal"))
                        d_3456_v22_: _dafny.Array
                        nw550_ = _dafny.Array(_dafny.Seq({}), 17)
                        d_3456_v22_ = nw550_
                        index585_ = default__.safeIndex(474, (d_3456_v22_).length(0))
                        (d_3456_v22_)[index585_] = (d_3439_v19_) + (d_3439_v19_)
                        index586_ = default__.safeIndex(474, (d_3456_v22_).length(0))
                        (d_3456_v22_)[index586_] = default__.fm48(globalState)
                        d_3457_v23_: str
                        d_3457_v23_ = _dafny.CodePoint('f')
                        d_3458_v24_: C1
                        nw551_ = C1()
                        nw551_.ctor__(d_3455_cf10_, d_3451_cf14_, (self).f12, d_3429_v15_.f27, p2, (self).f19)
                        d_3458_v24_ = nw551_
                        d_3459_v25_: _dafny.MultiSet
                        d_3459_v25_ = _dafny.MultiSet([d_3458_v24_, d_3458_v24_])
                        d_3460_v26_: _dafny.Map
                        d_3460_v26_ = _dafny.Map({(self).f20: d_3459_v25_})
                        rhs569_ = (True) and (p2)
                        rhs570_ = p2
                        rhs571_ = not((d_3459_v25_).isdisjoint((d_3459_v25_) | (((d_3460_v26_)[p2] if (p2) in (d_3460_v26_) else d_3459_v25_))))
                        rhs572_ = default__.safeModuloInt(default__.safeModuloInt(len(d_3405_v0_), d_3451_cf14_), ((d_3456_v22_)[default__.safeIndex(474, (d_3456_v22_).length(0))])[default__.safeIndex(len(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_3461_i3_ in range(default__.abs(-546))])).set(default__.safeIndex(d_3452_cf13_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_3462_i3_ in range(default__.abs(-546))]))), p1)).set(default__.safeIndex((self).f11, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_3463_i3_ in range(default__.abs(-546))])).set(default__.safeIndex(d_3452_cf13_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_3464_i3_ in range(default__.abs(-546))]))), p1))), d_3457_v23_)), len((d_3456_v22_)[default__.safeIndex(474, (d_3456_v22_).length(0))]))])
                        rhs573_ = ((self).f19)[default__.safeIndex(d_3429_v15_.f27, len((self).f19))]
                        lhs475_ = globalState
                        lhs476_ = globalState
                        lhs477_ = globalState
                        lhs478_ = globalState
                        lhs475_.f2 = rhs569_
                        lhs476_.f2 = rhs570_
                        lhs477_.f7 = rhs571_
                        lhs478_.f6 = rhs572_
                        d_3457_v23_ = rhs573_
                    elif source48_.is_DC8:
                        d_3465___mcc_h9_ = source48_.cf15
                        d_3466___mcc_h10_ = source48_.cf16
                        d_3467___mcc_h11_ = source48_.cf17
                        d_3468_cf17_ = d_3467___mcc_h11_
                        d_3469_cf16_ = d_3466___mcc_h10_
                        d_3470_cf15_ = d_3465___mcc_h9_
                        d_3471_v27_: str
                        d_3471_v27_ = _dafny.CodePoint('d')
                        d_3471_v27_ = p1
                        d_3472_v28_: _dafny.Array
                        nw552_ = _dafny.Array(None, 24)
                        d_3472_v28_ = nw552_
                        d_3473_v29_: _dafny.Set
                        d_3473_v29_ = _dafny.Set({d_3472_v28_, d_3472_v28_})
                        d_3474_v30_: D1
                        d_3474_v30_ = D1_DC2(p1)
                        d_3475_v31_: D2
                        d_3475_v31_ = D2_DC7(d_3443_v20_, d_3473_v29_, d_3474_v30_, d_3429_v15_.f27, d_3429_v15_.f27)
                        d_3443_v20_ = (d_3475_v31_).cf10
                        d_3476_v32_: C7
                        nw553_ = C7()
                        nw553_.ctor__()
                        d_3476_v32_ = nw553_
                        d_3477_v33_: _dafny.Seq
                        d_3477_v33_ = _dafny.SeqWithoutIsStrInference([d_3412_v7_])
                        d_3477_v33_ = _dafny.SeqWithoutIsStrInference([d_3412_v7_, _dafny.Map({d_3469_cf16_: not((self).f20)})])
                    elif True:
                        d_3478___mcc_h12_ = source48_.cf9
                        d_3479_cf9_ = d_3478___mcc_h12_
                        d_3480_v34_: _dafny.Array
                        nw554_ = _dafny.Array(None, 29)
                        d_3480_v34_ = nw554_
                        d_3480_v34_ = d_3480_v34_
                        d_3481_v35_: _dafny.Array
                        nw555_ = _dafny.Array(_dafny.MultiSet({}), 11)
                        d_3481_v35_ = nw555_
                        d_3482_v36_: _dafny.Map
                        d_3482_v36_ = _dafny.Map({p2: len(d_3479_cf9_)})
                        d_3483_v37_: _dafny.MultiSet
                        d_3483_v37_ = _dafny.MultiSet([d_3482_v36_, d_3482_v36_, d_3482_v36_, d_3482_v36_, d_3482_v36_])
                        index587_ = default__.safeIndex(733, (d_3481_v35_).length(0))
                        (d_3481_v35_)[index587_] = d_3483_v37_
                        index588_ = default__.safeIndex(733, (d_3481_v35_).length(0))
                        rhs574_ = d_3483_v37_
                        rhs575_ = (not ((self).f20) or ((self).f20)) and ((self).f20)
                        rhs576_ = (default__.safeDivisionInt(len(d_3423_v10_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otjwwaac"))))) > ((self).f12)
                        lhs479_ = d_3481_v35_
                        lhs480_ = default__.safeIndex(733, (d_3481_v35_).length(0))
                        lhs481_ = globalState
                        lhs482_ = globalState
                        lhs479_[lhs480_] = rhs574_
                        lhs481_.f2 = rhs575_
                        lhs482_.f7 = rhs576_
                        d_3484_v38_: _dafny.Seq
                        d_3484_v38_ = _dafny.SeqWithoutIsStrInference([not(p2), (self).f20, p2, p2, p2])
                        d_3485_v39_: C4
                        nw556_ = C4()
                        nw556_.ctor__(d_3423_v10_, p3, (self).f20, (self).f19, (0) - (len(d_3484_v38_)), d_3429_v15_.f27)
                        d_3485_v39_ = nw556_
                        (d_3429_v15_).f27 = d_3429_v15_.f27
                    pass
            pass
        d_3486_v40_: _dafny.Seq
        d_3486_v40_ = _dafny.SeqWithoutIsStrInference([p3])
        d_3487_v41_: _dafny.Array
        nw557_ = _dafny.Array(None, 9)
        nw557_[int(0)] = p0
        nw557_[int(1)] = len((self).f19)
        nw557_[int(2)] = len(d_3486_v40_)
        nw557_[int(3)] = (self).f11
        nw557_[int(4)] = (self).f12
        nw557_[int(5)] = (self).f12
        nw557_[int(6)] = ((self).f11) - ((self).f11)
        nw557_[int(7)] = (p3) - ((self).f12)
        nw557_[int(8)] = (len(d_3423_v10_)) * ((self).f11)
        d_3487_v41_ = nw557_
        guard_loop_12_: int
        for guard_loop_12_ in _dafny.IntegerRange(0, (d_3487_v41_).length(0)):
            d_3488_i4_: int = guard_loop_12_
            if (True) and (((0) <= (d_3488_i4_)) and ((d_3488_i4_) < ((d_3487_v41_).length(0)))):
                (d_3487_v41_)[(d_3488_i4_)] = default__.safeModuloInt(d_3488_i4_, default__.safeDivisionInt(262, (self).f12))
        d_3489_i5_: int
        d_3489_i5_ = 0
        with _dafny.label("19"):
            while p2:
                with _dafny.c_label("19"):
                    if (d_3489_i5_) >= (100):
                        raise _dafny.Break("19")
                    d_3489_i5_ = (d_3489_i5_) + (1)
                    d_3490_v42_: _dafny.Set
                    d_3490_v42_ = _dafny.Set({(self).f12})
                    d_3491_v43_: _dafny.MultiSet
                    d_3491_v43_ = _dafny.MultiSet([p1])
                    d_3492_v44_: _dafny.Map
                    d_3492_v44_ = _dafny.Map({d_3490_v42_: d_3491_v43_})
                    d_3493_v45_: _dafny.Array
                    nw558_ = _dafny.Array(None, 5)
                    nw558_[int(0)] = default__.fm69((self).f20, 373, ((d_3492_v44_)[d_3490_v42_] if (d_3490_v42_) in (d_3492_v44_) else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1]))), len(_dafny.SeqWithoutIsStrInference([p2, (self).f20])), globalState)
                    nw558_[int(1)] = ((self).f11) <= (p0)
                    nw558_[int(2)] = (True if (self).f20 else p2)
                    nw558_[int(3)] = (self).f20
                    nw558_[int(4)] = True
                    d_3493_v45_ = nw558_
                    index589_ = default__.safeIndex(407, (d_3493_v45_).length(0))
                    (d_3493_v45_)[index589_] = p2
                    d_3494_v46_: _dafny.MultiSet
                    d_3494_v46_ = _dafny.MultiSet([(self).f12])
                    index590_ = default__.safeIndex(407, (d_3493_v45_).length(0))
                    rhs577_ = (0) - ((self).f12)
                    rhs578_ = (self).f19
                    rhs579_ = (d_3494_v46_).isdisjoint(d_3494_v46_)
                    lhs483_ = globalState
                    lhs484_ = d_3493_v45_
                    lhs485_ = default__.safeIndex(407, (d_3493_v45_).length(0))
                    lhs483_.f6 = rhs577_
                    d_3423_v10_ = rhs578_
                    lhs484_[lhs485_] = rhs579_
                    (globalState).f7 = not((p2) or ((self).f20))
                    index591_ = default__.safeIndex(45, (d_3487_v41_).length(0))
                    (d_3487_v41_)[index591_] = (self).f12
                    index592_ = default__.safeIndex(45, (d_3487_v41_).length(0))
                    (d_3487_v41_)[index592_] = (0) - (p3)
                    d_3495_v47_: _dafny.MultiSet
                    d_3495_v47_ = _dafny.MultiSet([(self).f20, True])
                    d_3496_v48_: C2
                    nw559_ = C2()
                    nw559_.ctor__((d_3495_v47_).cardinality, (self).f19, (d_3487_v41_)[default__.safeIndex(45, (d_3487_v41_).length(0))], len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([p3, p0, 296])).cardinality for d_3497_i6_ in range(default__.abs(-99))])))
                    d_3496_v48_ = nw559_
                    d_3498_v49_: _dafny.Map
                    d_3498_v49_ = _dafny.Map({(d_3487_v41_)[default__.safeIndex(45, (d_3487_v41_).length(0))]: d_3496_v48_})
                    d_3499_v50_: _dafny.Array
                    nw560_ = _dafny.Array(D3.default()(), 9)
                    d_3499_v50_ = nw560_
                    d_3500_v51_: _dafny.Map
                    d_3500_v51_ = _dafny.Map({d_3499_v50_: d_3498_v49_})
                    d_3501_v52_: _dafny.Map
                    d_3501_v52_ = _dafny.Map({default__.safeDivisionInt(199, p3): (d_3498_v49_) | (((d_3500_v51_)[d_3499_v50_] if (d_3499_v50_) in (d_3500_v51_) else d_3498_v49_))})
                    d_3501_v52_ = (d_3501_v52_).set((self).f12, (_dafny.Map({218: d_3496_v48_})).set(d_3496_v48_.f27, d_3496_v48_))
                    pass
            pass
        d_3502_v54_: _dafny.Set
        d_3502_v54_ = _dafny.Set({p3})
        d_3503_v55_: _dafny.Seq
        def iife220_():
            coll76_ = _dafny.Set()
            compr_76_: int
            for compr_76_ in (d_3486_v40_).Elements:
                d_3504_v53_: int = compr_76_
                if (d_3504_v53_) in (d_3486_v40_):
                    coll76_ = coll76_.union(_dafny.Set([(d_3504_v53_) * (-695)]))
            return _dafny.Set(coll76_)
        d_3503_v55_ = _dafny.SeqWithoutIsStrInference([p2, True, (iife220_()
        ).isdisjoint(d_3502_v54_), (p1) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))), p2])
        r0 = d_3503_v55_
        d_3505_v56_: _dafny.Map
        d_3505_v56_ = _dafny.Map({(self).f20: (self).f12})
        r1 = (len(d_3505_v56_)) >= (p0)
        r2 = p2
        r3 = default__.safeDivisionInt(p3, p0)
        return r0, r1, r2, r3

    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20

class C10(T3, T2, T1, T0):
    def  __init__(self):
        self._f16: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f11: int = int(0)
        self._f12: int = int(0)
        self._f14: int = int(0)
        self._f15: bool = False
        self._f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f17: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f18: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
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
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f17, f18, f16, f14, f15, f13, f11, f12):
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f16 = f16
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f13 = f13
        (self)._f11 = f11
        (self)._f12 = f12

    def fm2(self, p0, p1, p2, p3, globalState):
        return (self).f11

    def fm3(self, p0, p1, p2, globalState):
        return (_dafny.Map({(self).f15: 990})) | (_dafny.Map({(self).f18: (self).f12}))

    def fm1(self, p0, p1, p2, globalState):
        return (self).f18

    def fm4(self, p0, p1, p2, p3, globalState):
        return (self).f11

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: T0 = None
        (globalState).f6 = (self).f12
        d_3506_v0_: _dafny.Array
        def lambda137_(d_3507_i0_):
            return (d_3507_i0_) * ((self).f14)

        init81_ = lambda137_
        nw561_ = _dafny.Array(None, 4)
        for i0_81_ in range(nw561_.length(0)):
            nw561_[i0_81_] = init81_(i0_81_)
        d_3506_v0_ = nw561_
        d_3508_v1_: _dafny.Seq
        d_3508_v1_ = _dafny.SeqWithoutIsStrInference([False])
        d_3509_v2_: _dafny.Map
        d_3509_v2_ = _dafny.Map({d_3506_v0_: _dafny.Set({(d_3508_v1_)[default__.safeIndex((0) - (-76), len(d_3508_v1_))], (self).f15, (self).f15, (self).f18})})
        rhs580_ = (d_3509_v2_) | (d_3509_v2_)
        rhs581_ = (self).f15
        d_3509_v2_ = rhs580_
        r1 = rhs581_
        d_3510_v3_: D0
        d_3510_v3_ = D0_DC1(((self).f15) == (p1), (self).f12)
        source49_ = d_3510_v3_
        if source49_.is_DC1:
            d_3511___mcc_h0_ = source49_.cf1
            d_3512___mcc_h1_ = source49_.cf2
            d_3513_cf2_ = d_3512___mcc_h1_
            d_3514_cf1_ = d_3511___mcc_h0_
            nw562_ = _dafny.Array(int(0), 2)
            d_3506_v0_ = nw562_
            d_3515_v4_: _dafny.Array
            nw563_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
            d_3515_v4_ = nw563_
            index593_ = default__.safeIndex(280, (d_3515_v4_).length(0))
            (d_3515_v4_)[index593_] = (self).f13
            index594_ = default__.safeIndex(280, (d_3515_v4_).length(0))
            (d_3515_v4_)[index594_] = (self).f13
            d_3516_v5_: _dafny.Seq
            d_3516_v5_ = _dafny.SeqWithoutIsStrInference([d_3513_cf2_, (self).f11, (self).f14, (self).f11])
            d_3516_v5_ = d_3516_v5_
            d_3517_v6_: D2
            d_3517_v6_ = D2_DC8(p1, (self).f18, self.f16)
            index595_ = default__.safeIndex(280, (d_3515_v4_).length(0))
            (d_3515_v4_)[index595_] = ((d_3517_v6_).cf17).set(default__.safeIndex(-472, len((d_3517_v6_).cf17)), _dafny.CodePoint('g'))
        elif True:
            d_3518___mcc_h2_ = source49_.cf0
            d_3519_cf0_ = d_3518___mcc_h2_
            d_3520_v7_: _dafny.Array
            nw564_ = _dafny.Array(False, 7)
            d_3520_v7_ = nw564_
            d_3521_v8_: _dafny.Map
            d_3521_v8_ = _dafny.Map({((self).f14) - ((self).f12): d_3520_v7_})
            d_3521_v8_ = (d_3521_v8_).set(-228, d_3520_v7_)
            index596_ = default__.safeIndex(326, (d_3506_v0_).length(0))
            (d_3506_v0_)[index596_] = (self).f14
            d_3522_v9_: _dafny.MultiSet
            d_3522_v9_ = _dafny.MultiSet([p1])
            d_3523_v10_: str
            d_3523_v10_ = _dafny.CodePoint('o')
            d_3524_v11_: D1
            d_3524_v11_ = D1_DC4((self).f18)
            d_3525_v12_: _dafny.Set
            d_3525_v12_ = _dafny.Set({(d_3524_v11_).cf5})
            d_3526_v13_: _dafny.Map
            d_3526_v13_ = _dafny.Map({(self).f14: len(d_3525_v12_)})
            d_3527_v14_: _dafny.Map
            d_3527_v14_ = _dafny.Map({d_3523_v10_: d_3526_v13_})
            index597_ = default__.safeIndex(326, (d_3506_v0_).length(0))
            (d_3506_v0_)[index597_] = (((d_3522_v9_)[not((self).f18)] if (not((self).f18)) in (d_3522_v9_) else (0) - (default__.fm0((self).fm1((self).f14, p1, False, globalState), globalState)))) + (len(d_3527_v14_))
            index598_ = default__.safeIndex(326, (d_3506_v0_).length(0))
            (d_3506_v0_)[index598_] = (self).f14
            d_3528_v15_: _dafny.Map
            d_3528_v15_ = _dafny.Map({self.f16: len(self.f16)})
            d_3529_v16_: _dafny.Map
            d_3529_v16_ = _dafny.Map({len(d_3528_v15_): (self).f15})
            (globalState).f6 = ((d_3506_v0_)[default__.safeIndex(326, (d_3506_v0_).length(0))] if (self).f15 else len((d_3529_v16_) | ((d_3529_v16_).set(-251, (self).f18))))
        d_3530_v17_: _dafny.Array
        nw565_ = _dafny.Array(False, 24)
        d_3530_v17_ = nw565_
        index599_ = default__.safeIndex(701, (d_3530_v17_).length(0))
        (d_3530_v17_)[index599_] = (self).f18
        d_3531_v18_: _dafny.Seq
        d_3531_v18_ = _dafny.SeqWithoutIsStrInference([d_3506_v0_, d_3506_v0_])
        d_3532_v19_: _dafny.Seq
        d_3532_v19_ = _dafny.SeqWithoutIsStrInference([d_3531_v18_])
        d_3533_v20_: _dafny.Array
        nw566_ = _dafny.Array(None, 20)
        nw566_[int(0)] = d_3531_v18_
        nw566_[int(1)] = d_3531_v18_
        nw566_[int(2)] = d_3531_v18_
        nw566_[int(3)] = ((d_3532_v19_)[default__.safeIndex((self).f11, len(d_3532_v19_))]) + ((d_3531_v18_).set(default__.safeIndex(552, len(d_3531_v18_)), d_3506_v0_))
        nw566_[int(4)] = d_3531_v18_
        nw566_[int(5)] = (d_3531_v18_) + (d_3531_v18_)
        nw566_[int(6)] = (d_3531_v18_).set(default__.safeIndex((self).f12, len(d_3531_v18_)), d_3506_v0_)
        nw566_[int(7)] = (d_3531_v18_ if p1 else (_dafny.SeqWithoutIsStrInference([d_3506_v0_, d_3506_v0_])).set(default__.safeIndex((self).f12, len(_dafny.SeqWithoutIsStrInference([d_3506_v0_, d_3506_v0_]))), d_3506_v0_))
        nw566_[int(8)] = (d_3531_v18_) + (d_3531_v18_)
        nw566_[int(9)] = d_3531_v18_
        nw566_[int(10)] = _dafny.SeqWithoutIsStrInference([d_3506_v0_, d_3506_v0_, d_3506_v0_])
        nw566_[int(11)] = d_3531_v18_
        nw566_[int(12)] = d_3531_v18_
        nw566_[int(13)] = (d_3531_v18_ if (self).f15 else _dafny.SeqWithoutIsStrInference([d_3506_v0_, d_3506_v0_]))
        nw566_[int(14)] = (d_3531_v18_).set(default__.safeIndex(811, len(d_3531_v18_)), d_3506_v0_)
        nw566_[int(15)] = d_3531_v18_
        nw566_[int(16)] = d_3531_v18_
        nw566_[int(17)] = d_3531_v18_
        nw566_[int(18)] = _dafny.SeqWithoutIsStrInference([d_3506_v0_])
        nw566_[int(19)] = ((d_3532_v19_)[default__.safeIndex((self).f11, len(d_3532_v19_))]) + (d_3531_v18_)
        d_3533_v20_ = nw566_
        index600_ = default__.safeIndex(166, (d_3533_v20_).length(0))
        (d_3533_v20_)[index600_] = (d_3531_v18_).set(default__.safeIndex(531, len(d_3531_v18_)), d_3506_v0_)
        d_3534_v21_: D3
        d_3534_v21_ = D3_DC9(_dafny.SeqWithoutIsStrInference([d_3506_v0_]))
        index601_ = default__.safeIndex(701, (d_3530_v17_).length(0))
        index602_ = default__.safeIndex(166, (d_3533_v20_).length(0))
        rhs582_ = p1
        rhs583_ = 163
        rhs584_ = (self).f18
        rhs585_ = (d_3534_v21_).cf18
        rhs586_ = (self).f12
        lhs486_ = d_3530_v17_
        lhs487_ = default__.safeIndex(701, (d_3530_v17_).length(0))
        lhs488_ = globalState
        lhs489_ = d_3533_v20_
        lhs490_ = default__.safeIndex(166, (d_3533_v20_).length(0))
        lhs491_ = globalState
        lhs486_[lhs487_] = rhs582_
        r0 = rhs583_
        lhs488_.f2 = rhs584_
        lhs489_[lhs490_] = rhs585_
        lhs491_.f6 = rhs586_
        d_3535_v23_: _dafny.Set
        d_3535_v23_ = _dafny.Set({641, (self).f11})
        d_3536_v24_: _dafny.Map
        def iife221_():
            coll77_ = _dafny.Set()
            compr_77_: int
            for compr_77_ in _dafny.IntegerRange(327, 169):
                d_3537_v22_: int = compr_77_
                if ((327) <= (d_3537_v22_)) and ((d_3537_v22_) < (169)):
                    coll77_ = coll77_.union(_dafny.Set([default__.safeDivisionInt(d_3537_v22_, (self).f11)]))
            return _dafny.Set(coll77_)
        d_3536_v24_ = _dafny.Map({(d_3535_v23_).issubset(iife221_()
        ): default__.fm5(globalState)})
        d_3538_v25_: str
        d_3538_v25_ = _dafny.CodePoint('e')
        d_3536_v24_ = (d_3536_v24_).set(((self).f15 if p1 else False), d_3538_v25_)
        index603_ = default__.safeIndex(997, (d_3506_v0_).length(0))
        (d_3506_v0_)[index603_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouawvkk")))
        index604_ = default__.safeIndex(997, (d_3506_v0_).length(0))
        (d_3506_v0_)[index604_] = ((self).f11 if p1 else 525)
        r0 = (self).f14
        r1 = (d_3508_v1_)[default__.safeIndex(len(d_3535_v23_), len(d_3508_v1_))]
        d_3539_v26_: T0
        nw567_ = C9()
        nw567_.ctor__((self).f17, True, (self).f12, ((self).f12) + ((self).f11))
        d_3539_v26_ = nw567_
        r2 = d_3539_v26_
        return r0, r1, r2

    def m4(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        d_3540_v0_: str
        d_3540_v0_ = _dafny.CodePoint('i')
        d_3540_v0_ = d_3540_v0_
        d_3541_v1_: _dafny.Set
        d_3542_v2_: bool
        d_3543_v3_: bool
        out92_: _dafny.Set
        out93_: bool
        out94_: bool
        out92_, out93_, out94_ = (self).m5(globalState)
        d_3541_v1_ = out92_
        d_3542_v2_ = out93_
        d_3543_v3_ = out94_
        d_3544_v4_: _dafny.Array
        nw568_ = _dafny.Array(None, 1)
        nw568_[int(0)] = (self).f18
        d_3544_v4_ = nw568_
        guard_loop_13_: int
        for guard_loop_13_ in _dafny.IntegerRange(0, (d_3544_v4_).length(0)):
            d_3545_i0_: int = guard_loop_13_
            if (True) and (((0) <= (d_3545_i0_)) and ((d_3545_i0_) < ((d_3544_v4_).length(0)))):
                (d_3544_v4_)[(d_3545_i0_)] = ((self).f11) != ((self).f12)
        r0 = (self).f12
        d_3546_v5_: _dafny.MultiSet
        d_3546_v5_ = _dafny.MultiSet([d_3540_v0_, d_3540_v0_, d_3540_v0_])
        if not(default__.fm69(d_3542_v2_, ((self).f14) * ((self).f11), d_3546_v5_, (self).f12, globalState)):
            if (self).f15:
                d_3547_v6_: _dafny.Array
                nw569_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                d_3547_v6_ = nw569_
                d_3548_v7_: _dafny.Map
                d_3548_v7_ = _dafny.Map({d_3547_v6_: p0})
                d_3544_v4_ = ((d_3548_v7_)[d_3547_v6_] if (d_3547_v6_) in (d_3548_v7_) else p0)
                d_3549_v8_: _dafny.Array
                def lambda138_(d_3550_i1_):
                    return (d_3550_i1_) * (len(_dafny.SeqWithoutIsStrInference([(self).f15])))

                init82_ = lambda138_
                nw570_ = _dafny.Array(None, 12)
                for i0_82_ in range(nw570_.length(0)):
                    nw570_[i0_82_] = init82_(i0_82_)
                d_3549_v8_ = nw570_
                d_3551_v9_: _dafny.Map
                d_3551_v9_ = _dafny.Map({(self).f14: (self).f14})
                d_3552_v10_: _dafny.Set
                d_3552_v10_ = _dafny.Set({(self).f13, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfypunwc")), (self).f13, (self).f13, (self).f13})
                index605_ = default__.safeIndex(851, (d_3549_v8_).length(0))
                (d_3549_v8_)[index605_] = default__.safeModuloInt((self).fm4(((d_3551_v9_)[(self).f11] if ((self).f11) in (d_3551_v9_) else -971), (self).f11, d_3542_v2_, d_3543_v3_, globalState), len(d_3552_v10_))
                index606_ = default__.safeIndex(851, (d_3549_v8_).length(0))
                (d_3549_v8_)[index606_] = (self).f11
                d_3553_v11_: _dafny.Set
                d_3554_v12_: bool
                d_3555_v13_: bool
                out95_: _dafny.Set
                out96_: bool
                out97_: bool
                out95_, out96_, out97_ = (self).m5(globalState)
                d_3553_v11_ = out95_
                d_3554_v12_ = out96_
                d_3555_v13_ = out97_
                d_3556_v14_: _dafny.MultiSet
                d_3556_v14_ = _dafny.MultiSet([p0])
                d_3556_v14_ = d_3556_v14_
                r0 = (default__.safeDivisionInt((0) - ((self).f11), (self).f14)) * (((self).f11 if d_3542_v2_ else (d_3549_v8_)[default__.safeIndex(851, (d_3549_v8_).length(0))]))
            elif True:
                d_3557_v15_: C4
                nw571_ = C4()
                nw571_.ctor__((self).f17, (self).f12, d_3543_v3_, (self).f17, default__.safeDivisionInt((self).f14, (0) - ((0) - ((self).f14))), ((self).f12 if (self).f18 else (self).f14))
                d_3557_v15_ = nw571_
                d_3557_v15_ = d_3557_v15_
                (globalState).f7 = d_3543_v3_
                (globalState).f6 = (self).f12
                d_3558_v16_: _dafny.Map
                d_3558_v16_ = _dafny.Map({(self).f18: (self).f11})
                d_3558_v16_ = d_3558_v16_
                d_3559_v17_: _dafny.Map
                d_3559_v17_ = _dafny.Map({d_3544_v4_: d_3543_v3_})
                d_3560_v18_: _dafny.Seq
                d_3560_v18_ = _dafny.SeqWithoutIsStrInference([d_3543_v3_, (len(d_3559_v17_)) <= ((self).f12), (self).f15, False, (d_3557_v15_).fm24((self).f12, (self).f18, globalState)])
                rhs587_ = default__.fm50(globalState)
                rhs588_ = (self).f14
                rhs589_ = 314
                rhs590_ = default__.fm22((self).f11, globalState)
                lhs492_ = globalState
                d_3558_v16_ = rhs587_
                lhs492_.f6 = rhs588_
                r0 = rhs589_
                d_3560_v18_ = rhs590_
            index607_ = default__.safeIndex(615, (p0).length(0))
            (p0)[index607_] = d_3542_v2_
            d_3561_v19_: _dafny.MultiSet
            d_3561_v19_ = _dafny.MultiSet([d_3542_v2_])
            d_3562_v20_: _dafny.Set
            d_3562_v20_ = _dafny.Set({_dafny.MultiSet([len(_dafny.Map({(self).f11: (self).f18}))])})
            d_3563_v21_: _dafny.Map
            d_3563_v21_ = _dafny.Map({(self).f18: d_3543_v3_})
            d_3564_v22_: _dafny.Seq
            d_3564_v22_ = _dafny.SeqWithoutIsStrInference([len(d_3563_v21_)])
            d_3565_v23_: _dafny.MultiSet
            d_3565_v23_ = _dafny.MultiSet([(self).f11, len(d_3564_v22_), (self).f11])
            d_3566_v24_: _dafny.Seq
            d_3566_v24_ = _dafny.SeqWithoutIsStrInference([d_3565_v23_])
            index608_ = default__.safeIndex(615, (p0).length(0))
            rhs591_ = ((d_3561_v19_)[d_3543_v3_] if (d_3543_v3_) in (d_3561_v19_) else (self).f14)
            rhs592_ = ((_dafny.Set({d_3565_v23_, d_3565_v23_, (d_3566_v24_)[default__.safeIndex((0) - ((self).f12), len(d_3566_v24_))]})) | (d_3562_v20_)).issubset(d_3562_v20_)
            rhs593_ = ((self).f13 if d_3542_v2_ else (self).f17)
            lhs493_ = globalState
            lhs494_ = p0
            lhs495_ = default__.safeIndex(615, (p0).length(0))
            lhs496_ = self
            lhs493_.f6 = rhs591_
            lhs494_[lhs495_] = rhs592_
            lhs496_.f16 = rhs593_
            d_3567_v25_: _dafny.Set
            d_3567_v25_ = _dafny.Set({(self).f12, (self).f14})
            d_3568_v26_: bool
            d_3569_v27_: _dafny.Seq
            out98_: bool
            out99_: _dafny.Seq
            out98_, out99_ = default__.m0((self).f18, d_3543_v3_, len(d_3567_v25_), (self).f14, globalState)
            d_3568_v26_ = out98_
            d_3569_v27_ = out99_
            if d_3568_v26_:
                (globalState).f7 = d_3543_v3_
                (globalState).f6 = default__.safeDivisionInt(len(d_3541_v1_), ((self).f14) - ((0) - ((self).f14)))
                (globalState).f7 = d_3568_v26_
                d_3570_v28_: _dafny.Map
                d_3570_v28_ = _dafny.Map({(p0)[default__.safeIndex(615, (p0).length(0))]: d_3544_v4_})
                d_3571_v29_: _dafny.Map
                d_3571_v29_ = _dafny.Map({(self).f12: d_3544_v4_})
                d_3570_v28_ = (d_3570_v28_).set(d_3542_v2_, ((d_3571_v29_)[(self).f11] if ((self).f11) in (d_3571_v29_) else d_3544_v4_))
                d_3572_v30_: C7
                nw572_ = C7()
                nw572_.ctor__()
                d_3572_v30_ = nw572_
            elif True:
                index609_ = default__.safeIndex(615, (p0).length(0))
                (p0)[index609_] = ((self).f11) > ((len(d_3564_v22_)) + ((self).f12))
                d_3573_v31_: _dafny.Array
                def lambda139_(d_3574_v23_, d_3575_v3_):
                    def lambda140_(d_3576_i2_):
                        return (_dafny.MultiSet([d_3574_v23_])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.Map({(self).f14: d_3575_v3_}))]), d_3574_v23_])))

                    return lambda140_

                init83_ = lambda139_(d_3565_v23_, d_3543_v3_)
                nw573_ = _dafny.Array(None, 14)
                for i0_83_ in range(nw573_.length(0)):
                    nw573_[i0_83_] = init83_(i0_83_)
                d_3573_v31_ = nw573_
                d_3577_v32_: _dafny.MultiSet
                d_3577_v32_ = _dafny.MultiSet([d_3565_v23_, _dafny.MultiSet([(0) - (len((self).f13)), (self).f14]), _dafny.MultiSet((D4_DC13(d_3564_v22_)).cf29), d_3565_v23_])
                index610_ = default__.safeIndex(997, (d_3573_v31_).length(0))
                (d_3573_v31_)[index610_] = d_3577_v32_
                d_3578_v33_: _dafny.Map
                d_3578_v33_ = _dafny.Map({(self).f14: _dafny.MultiSet([d_3565_v23_])})
                index611_ = default__.safeIndex(615, (p0).length(0))
                index612_ = default__.safeIndex(997, (d_3573_v31_).length(0))
                rhs594_ = ((self).f11) >= ((self).f11)
                rhs595_ = ((d_3577_v32_).intersection(d_3577_v32_) if d_3542_v2_ else ((d_3578_v33_)[(self).fm2((self).f14, (self).f18, (self).f11, (self).f15, globalState)] if ((self).fm2((self).f14, (self).f18, (self).f11, (self).f15, globalState)) in (d_3578_v33_) else _dafny.MultiSet(d_3566_v24_)))
                lhs497_ = p0
                lhs498_ = default__.safeIndex(615, (p0).length(0))
                lhs499_ = d_3573_v31_
                lhs500_ = default__.safeIndex(997, (d_3573_v31_).length(0))
                lhs497_[lhs498_] = rhs594_
                lhs499_[lhs500_] = rhs595_
                (globalState).f6 = len((d_3541_v1_).intersection((d_3541_v1_) | (d_3541_v1_)))
                d_3579_v34_: _dafny.Array
                nw574_ = _dafny.Array(_dafny.Set({}), 4)
                d_3579_v34_ = nw574_
                d_3580_v35_: T1
                nw575_ = C6()
                nw575_.ctor__(_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f11: (self).f15}) for d_3581_i3_ in range(default__.abs(874))]), (self).f12, not(d_3543_v3_), (self).f17, (self).f14, (self).f11)
                d_3580_v35_ = nw575_
                d_3582_v36_: _dafny.Map
                d_3582_v36_ = _dafny.Map({(self).f15: (d_3580_v35_).f11})
                d_3583_v37_: _dafny.Map
                d_3583_v37_ = _dafny.Map({d_3580_v35_: ((d_3582_v36_)[d_3568_v26_] if (d_3568_v26_) in (d_3582_v36_) else (d_3580_v35_).f11)})
                d_3584_v38_: C2
                nw576_ = C2()
                nw576_.ctor__(len(d_3569_v27_), self.f16, (self).f14, len(d_3583_v37_))
                d_3584_v38_ = nw576_
                d_3585_v39_: _dafny.Set
                d_3585_v39_ = _dafny.Set({d_3584_v38_})
                index613_ = default__.safeIndex(707, (d_3579_v34_).length(0))
                (d_3579_v34_)[index613_] = d_3585_v39_
                d_3586_v40_: _dafny.Seq
                d_3586_v40_ = _dafny.SeqWithoutIsStrInference([d_3585_v39_, d_3585_v39_])
                index614_ = default__.safeIndex(707, (d_3579_v34_).length(0))
                index615_ = default__.safeIndex(615, (p0).length(0))
                rhs596_ = ((d_3586_v40_)[default__.safeIndex(76, len(d_3586_v40_))]).intersection(d_3585_v39_)
                rhs597_ = default__.fm0((self).f15, globalState)
                rhs598_ = d_3568_v26_
                rhs599_ = (self).f15
                rhs600_ = ((d_3580_v35_).f12) < ((d_3580_v35_).f11)
                lhs501_ = d_3579_v34_
                lhs502_ = default__.safeIndex(707, (d_3579_v34_).length(0))
                lhs503_ = d_3584_v38_
                lhs504_ = p0
                lhs505_ = default__.safeIndex(615, (p0).length(0))
                lhs506_ = globalState
                lhs507_ = globalState
                lhs501_[lhs502_] = rhs596_
                lhs503_.f27 = rhs597_
                lhs504_[lhs505_] = rhs598_
                lhs506_.f7 = rhs599_
                lhs507_.f2 = rhs600_
                d_3587_v41_: _dafny.Seq
                d_3587_v41_ = _dafny.SeqWithoutIsStrInference([d_3544_v4_, d_3544_v4_, d_3544_v4_, d_3544_v4_, d_3544_v4_])
                d_3588_v42_: _dafny.MultiSet
                d_3588_v42_ = _dafny.MultiSet([(d_3587_v41_)[default__.safeIndex((self).f14, len(d_3587_v41_))]])
                d_3589_v43_: _dafny.Map
                d_3589_v43_ = _dafny.Map({(d_3580_v35_).f11: d_3588_v42_})
                d_3589_v43_ = (d_3589_v43_).set((self).f12, (d_3588_v42_).intersection(d_3588_v42_))
            d_3590_v44_: _dafny.Map
            d_3590_v44_ = _dafny.Map({(self).f14: d_3544_v4_})
            d_3591_v45_: _dafny.Map
            d_3591_v45_ = _dafny.Map({(self).f12: (0) - (len(d_3569_v27_))})
            d_3592_v46_: _dafny.Set
            d_3592_v46_ = _dafny.Set({d_3567_v25_})
            d_3593_v47_: D14
            d_3593_v47_ = D14_DC41((self).f15, False, (self).f17)
            d_3594_v48_: T0
            nw577_ = C1()
            nw577_.ctor__(((d_3590_v44_)[(self).f11] if ((self).f11) in (d_3590_v44_) else d_3544_v4_), default__.safeModuloInt(len(d_3591_v45_), (self).f14), (self).f11, 337, (d_3592_v46_).ispropersubset(d_3592_v46_), (d_3593_v47_).cf75)
            d_3594_v48_ = nw577_
            d_3594_v48_ = d_3594_v48_
        elif True:
            (self).f16 = (self).f13
            (self).f16 = ((self).f13) + (((self).f13) + (self.f16))
            d_3595_v49_: _dafny.Seq
            d_3595_v49_ = _dafny.SeqWithoutIsStrInference([(self).f14])
            d_3596_v50_: _dafny.Set
            d_3596_v50_ = _dafny.Set({((self).f11) * ((self).f12), len(d_3595_v49_), (self).f14})
            d_3596_v50_ = (d_3596_v50_) | (d_3596_v50_)
            d_3597_v51_: C4
            nw578_ = C4()
            nw578_.ctor__(self.f16, (self).f11, d_3542_v2_, (self).f17, (self).f11, (self).f14)
            d_3597_v51_ = nw578_
            d_3598_v52_: _dafny.Map
            d_3598_v52_ = _dafny.Map({d_3542_v2_: d_3597_v51_})
            d_3599_v53_: _dafny.Array
            nw579_ = _dafny.Array(None, 22)
            nw579_[int(0)] = (d_3597_v51_ if d_3543_v3_ else d_3597_v51_)
            nw579_[int(1)] = d_3597_v51_
            nw579_[int(2)] = ((d_3598_v52_)[d_3542_v2_] if (d_3542_v2_) in (d_3598_v52_) else d_3597_v51_)
            nw579_[int(3)] = d_3597_v51_
            nw579_[int(4)] = d_3597_v51_
            nw579_[int(5)] = d_3597_v51_
            nw579_[int(6)] = (D26_DC69(d_3597_v51_)).cf126
            nw579_[int(7)] = d_3597_v51_
            nw579_[int(8)] = d_3597_v51_
            nw579_[int(9)] = d_3597_v51_
            nw579_[int(10)] = d_3597_v51_
            nw579_[int(11)] = d_3597_v51_
            nw579_[int(12)] = d_3597_v51_
            nw579_[int(13)] = d_3597_v51_
            nw579_[int(14)] = d_3597_v51_
            nw579_[int(15)] = d_3597_v51_
            nw579_[int(16)] = d_3597_v51_
            nw579_[int(17)] = d_3597_v51_
            nw579_[int(18)] = d_3597_v51_
            nw579_[int(19)] = d_3597_v51_
            nw579_[int(20)] = d_3597_v51_
            nw579_[int(21)] = d_3597_v51_
            d_3599_v53_ = nw579_
            index616_ = default__.safeIndex(179, (d_3599_v53_).length(0))
            (d_3599_v53_)[index616_] = d_3597_v51_
            index617_ = default__.safeIndex(179, (d_3599_v53_).length(0))
            (d_3599_v53_)[index617_] = d_3597_v51_
            d_3600_v54_: _dafny.Map
            d_3600_v54_ = _dafny.Map({default__.fm0(d_3543_v3_, globalState): (self).f15})
            d_3601_v55_: _dafny.Seq
            d_3601_v55_ = _dafny.SeqWithoutIsStrInference([d_3600_v54_])
            d_3602_v56_: _dafny.MultiSet
            d_3602_v56_ = _dafny.MultiSet([True, (self).f18])
            d_3603_v57_: T2
            nw580_ = C6()
            nw580_.ctor__(d_3601_v55_, (d_3602_v56_).cardinality, (self).f18, (self).f13, (self).f11, (self).f12)
            d_3603_v57_ = nw580_
            d_3604_v58_: _dafny.Map
            d_3604_v58_ = _dafny.Map({p0: d_3603_v57_})
            d_3605_v59_: _dafny.Array
            nw581_ = _dafny.Array(None, 6)
            nw581_[int(0)] = _dafny.Map({p0: d_3603_v57_})
            nw581_[int(1)] = d_3604_v58_
            nw581_[int(2)] = (d_3604_v58_) | (d_3604_v58_)
            nw581_[int(3)] = d_3604_v58_
            nw581_[int(4)] = d_3604_v58_
            nw581_[int(5)] = d_3604_v58_
            d_3605_v59_ = nw581_
            index618_ = default__.safeIndex(469, (d_3605_v59_).length(0))
            (d_3605_v59_)[index618_] = (d_3604_v58_).set(d_3544_v4_, d_3603_v57_)
            d_3606_v60_: _dafny.Map
            d_3606_v60_ = _dafny.Map({d_3543_v3_: d_3540_v0_})
            index619_ = default__.safeIndex(469, (d_3605_v59_).length(0))
            rhs601_ = d_3604_v58_
            rhs602_ = (self).f18
            rhs603_ = default__.fm21(len((d_3606_v60_) | (d_3606_v60_)), (self).f18, globalState)
            lhs508_ = d_3605_v59_
            lhs509_ = default__.safeIndex(469, (d_3605_v59_).length(0))
            lhs510_ = globalState
            lhs508_[lhs509_] = rhs601_
            lhs510_.f2 = rhs602_
            d_3540_v0_ = rhs603_
        d_3607_v61_: _dafny.Seq
        d_3607_v61_ = _dafny.SeqWithoutIsStrInference([(self).f13, (self).f13])
        d_3608_v62_: _dafny.Map
        d_3608_v62_ = _dafny.Map({(self).f18: False})
        d_3609_v63_: _dafny.MultiSet
        d_3609_v63_ = _dafny.MultiSet([(0) - (len(_dafny.Map({(self).f14: d_3608_v62_})))])
        d_3610_v64_: D3
        d_3610_v64_ = D3_DC10(d_3543_v3_, (self).f12, _dafny.Map({self.f16: (self).f14}), (d_3607_v61_)[default__.safeIndex((0) - (len(d_3608_v62_)), len(d_3607_v61_))], (d_3609_v63_).cardinality)
        d_3611_v65_: D0
        d_3611_v65_ = D0_DC0(d_3546_v5_)
        d_3612_v66_: _dafny.Array
        nw582_ = _dafny.Array(None, 23)
        nw582_[int(0)] = (d_3610_v64_).cf22
        nw582_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qytdjteat"))
        nw582_[int(2)] = (self.f16) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubnvsov")))
        nw582_[int(3)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fphrhl"))) + ((self).f17)
        nw582_[int(4)] = self.f16
        nw582_[int(5)] = ((self).f17) + (self.f16)
        nw582_[int(6)] = self.f16
        nw582_[int(7)] = (self).f17
        nw582_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrv"))) + ((self).f17)
        nw582_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wuoqrwen"))
        nw582_[int(10)] = (self).f13
        nw582_[int(11)] = _dafny.SeqWithoutIsStrInference([d_3540_v0_ for d_3613_i5_ in range(default__.abs(255))])
        nw582_[int(12)] = self.f16
        nw582_[int(13)] = (self).f17
        nw582_[int(14)] = default__.fm44(len(_dafny.SeqWithoutIsStrInference([(self).f11 for d_3614_i6_ in range(default__.abs(875))])), d_3543_v3_, globalState)
        nw582_[int(15)] = (self).f17
        nw582_[int(16)] = (self).f17
        nw582_[int(17)] = self.f16
        nw582_[int(18)] = ((D2_DC8(d_3542_v2_, (self).f18, self.f16)).cf17) + ((self).f17)
        nw582_[int(19)] = (((self).f13).set(default__.safeIndex((self).f14, len((self).f13)), d_3540_v0_)) + (default__.fm31(d_3540_v0_, d_3611_v65_, globalState))
        nw582_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghpfxsjx"))
        nw582_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fx"))
        nw582_[int(22)] = (self.f16) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))
        d_3612_v66_ = nw582_
        guard_loop_14_: int
        for guard_loop_14_ in _dafny.IntegerRange(0, (d_3612_v66_).length(0)):
            d_3615_i4_: int = guard_loop_14_
            if (True) and (((0) <= (d_3615_i4_)) and ((d_3615_i4_) < ((d_3612_v66_).length(0)))):
                (d_3612_v66_)[(d_3615_i4_)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abjfwpiw"))) + ((d_3607_v61_)[default__.safeIndex((self).f12, len(d_3607_v61_))])
        r0 = (self).f11
        d_3616_v67_: _dafny.Seq
        d_3616_v67_ = _dafny.SeqWithoutIsStrInference([len((self).f17), 2])
        r1 = d_3616_v67_
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        index620_ = default__.safeIndex(960, (p0).length(0))
        (p0)[index620_] = (self).f15
        index621_ = default__.safeIndex(960, (p0).length(0))
        (p0)[index621_] = ((self).f15 if False else (self).f18)
        hi13_ = (self).f14
        for d_3617_i0_ in range(default__.fm0((p0)[default__.safeIndex(960, (p0).length(0))], globalState), hi13_):
            (globalState).f6 = (self).f14
            d_3618_v0_: str
            d_3618_v0_ = _dafny.CodePoint('j')
            (self).f16 = (((self).f13) + ((self).f17)).set(default__.safeIndex(((self).f11) * (963), len(((self).f13) + ((self).f17))), d_3618_v0_)
            d_3619_v1_: _dafny.Array
            def lambda141_(d_3620_i1_):
                return default__.safeModuloInt(d_3620_i1_, (self).f12)

            init84_ = lambda141_
            nw583_ = _dafny.Array(None, 9)
            for i0_84_ in range(nw583_.length(0)):
                nw583_[i0_84_] = init84_(i0_84_)
            d_3619_v1_ = nw583_
            d_3621_v2_: _dafny.Map
            d_3621_v2_ = _dafny.Map({d_3619_v1_: 66})
            d_3622_v3_: D22
            d_3622_v3_ = D22_DC63((_dafny.MultiSet([len(d_3621_v2_)])).cardinality, (self).f15, (p0)[default__.safeIndex(960, (p0).length(0))])
            d_3623_v4_: _dafny.Map
            d_3623_v4_ = _dafny.Map({(self).f14: (self).f15})
            d_3624_v5_: _dafny.Array
            nw584_ = _dafny.Array(None, 5)
            nw584_[int(0)] = (self).f14
            nw584_[int(1)] = (self).f14
            nw584_[int(2)] = (d_3622_v3_).cf118
            nw584_[int(3)] = (self).f14
            nw584_[int(4)] = (self).fm4(len(d_3623_v4_), 401, (p0)[default__.safeIndex(960, (p0).length(0))], (self).f15, globalState)
            d_3624_v5_ = nw584_
            d_3625_v6_: _dafny.Map
            d_3625_v6_ = _dafny.Map({(p0)[default__.safeIndex(960, (p0).length(0))]: d_3619_v1_})
            d_3626_v7_: _dafny.Array
            nw585_ = _dafny.Array(None, 11)
            nw585_[int(0)] = d_3625_v6_
            nw585_[int(1)] = d_3625_v6_
            nw585_[int(2)] = _dafny.Map({(self).f18: d_3624_v5_})
            nw585_[int(3)] = d_3625_v6_
            nw585_[int(4)] = _dafny.Map({(p0)[default__.safeIndex(960, (p0).length(0))]: d_3619_v1_})
            nw585_[int(5)] = d_3625_v6_
            nw585_[int(6)] = d_3625_v6_
            nw585_[int(7)] = _dafny.Map({(self).f18: d_3619_v1_})
            nw585_[int(8)] = d_3625_v6_
            nw585_[int(9)] = d_3625_v6_
            nw585_[int(10)] = d_3625_v6_
            d_3626_v7_ = nw585_
            index622_ = default__.safeIndex(216, (d_3626_v7_).length(0))
            (d_3626_v7_)[index622_] = (d_3625_v6_) | (d_3625_v6_)
            index623_ = default__.safeIndex(126, (d_3624_v5_).length(0))
            (d_3624_v5_)[index623_] = (self).f12
            d_3627_v8_: D4
            d_3627_v8_ = D4_DC15((0) - ((self).f11))
            d_3628_v9_: D2
            d_3628_v9_ = D2_DC8(False, (self).f15, self.f16)
            d_3629_v10_: _dafny.Map
            d_3629_v10_ = _dafny.Map({d_3628_v9_: d_3625_v6_})
            index624_ = default__.safeIndex(216, (d_3626_v7_).length(0))
            index625_ = default__.safeIndex(126, (d_3624_v5_).length(0))
            index626_ = default__.safeIndex(960, (p0).length(0))
            rhs604_ = (d_3627_v8_).cf32
            rhs605_ = d_3617_i0_
            rhs606_ = ((d_3629_v10_)[default__.fm39(globalState)] if (default__.fm39(globalState)) in (d_3629_v10_) else d_3625_v6_)
            rhs607_ = (self).f11
            rhs608_ = (self).f15
            lhs511_ = globalState
            lhs512_ = globalState
            lhs513_ = d_3626_v7_
            lhs514_ = default__.safeIndex(216, (d_3626_v7_).length(0))
            lhs515_ = d_3624_v5_
            lhs516_ = default__.safeIndex(126, (d_3624_v5_).length(0))
            lhs517_ = p0
            lhs518_ = default__.safeIndex(960, (p0).length(0))
            lhs511_.f6 = rhs604_
            lhs512_.f6 = rhs605_
            lhs513_[lhs514_] = rhs606_
            lhs515_[lhs516_] = rhs607_
            lhs517_[lhs518_] = rhs608_
            (globalState).f6 = (328) * (d_3617_i0_)
        r0 = (self).f15
        d_3630_v11_: D0
        d_3630_v11_ = D0_DC1(not(((self).f12) <= (default__.fm0(not((self).f18), globalState))), (self).f12)
        source50_ = d_3630_v11_
        if source50_.is_DC1:
            d_3631___mcc_h0_ = source50_.cf1
            d_3632___mcc_h1_ = source50_.cf2
            d_3633_cf2_ = d_3632___mcc_h1_
            d_3634_cf1_ = d_3631___mcc_h0_
            d_3635_v12_: _dafny.Set
            d_3635_v12_ = _dafny.Set({d_3633_cf2_})
            def iife222_():
                coll78_ = _dafny.Set()
                compr_78_: int
                for compr_78_ in (p1).Elements:
                    d_3636_v13_: int = compr_78_
                    if (d_3636_v13_) in (p1):
                        coll78_ = coll78_.union(_dafny.Set([default__.safeDivisionInt(d_3636_v13_, 307)]))
                return _dafny.Set(coll78_)
            (globalState).f7 = (d_3635_v12_).isdisjoint((d_3635_v12_) | (iife222_()
            ))
            d_3637_v14_: _dafny.Array
            nw586_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_3637_v14_ = nw586_
            d_3638_v15_: _dafny.Array
            nw587_ = _dafny.Array(None, 5)
            nw587_[int(0)] = d_3637_v14_
            nw587_[int(1)] = d_3637_v14_
            nw587_[int(2)] = d_3637_v14_
            nw587_[int(3)] = d_3637_v14_
            nw587_[int(4)] = d_3637_v14_
            d_3638_v15_ = nw587_
            d_3639_v16_: _dafny.Seq
            d_3639_v16_ = _dafny.SeqWithoutIsStrInference([d_3638_v15_, d_3638_v15_, d_3638_v15_])
            d_3640_v17_: _dafny.Map
            d_3640_v17_ = _dafny.Map({d_3633_cf2_: d_3638_v15_})
            d_3641_v18_: _dafny.Array
            nw588_ = _dafny.Array(None, 25)
            nw588_[int(0)] = (d_3638_v15_ if (self).f18 else d_3638_v15_)
            nw588_[int(1)] = d_3638_v15_
            nw588_[int(2)] = d_3638_v15_
            nw588_[int(3)] = d_3638_v15_
            nw588_[int(4)] = d_3638_v15_
            nw588_[int(5)] = (d_3639_v16_)[default__.safeIndex((self).f14, len(d_3639_v16_))]
            nw588_[int(6)] = d_3638_v15_
            nw588_[int(7)] = d_3638_v15_
            nw588_[int(8)] = d_3638_v15_
            nw588_[int(9)] = d_3638_v15_
            nw588_[int(10)] = (d_3638_v15_ if d_3634_cf1_ else d_3638_v15_)
            nw588_[int(11)] = d_3638_v15_
            nw588_[int(12)] = ((d_3640_v17_)[d_3633_cf2_] if (d_3633_cf2_) in (d_3640_v17_) else d_3638_v15_)
            nw588_[int(13)] = d_3638_v15_
            nw588_[int(14)] = d_3638_v15_
            nw588_[int(15)] = d_3638_v15_
            nw588_[int(16)] = d_3638_v15_
            nw588_[int(17)] = d_3638_v15_
            nw588_[int(18)] = d_3638_v15_
            nw588_[int(19)] = d_3638_v15_
            nw588_[int(20)] = d_3638_v15_
            nw588_[int(21)] = d_3638_v15_
            nw588_[int(22)] = d_3638_v15_
            nw588_[int(23)] = d_3638_v15_
            nw588_[int(24)] = (d_3638_v15_ if d_3634_cf1_ else d_3638_v15_)
            d_3641_v18_ = nw588_
            index627_ = default__.safeIndex(583, (d_3641_v18_).length(0))
            (d_3641_v18_)[index627_] = ((d_3640_v17_)[d_3633_cf2_] if (d_3633_cf2_) in (d_3640_v17_) else d_3638_v15_)
            d_3642_v19_: _dafny.Array
            def lambda142_(d_3643_i2_):
                return (self).f15

            init85_ = lambda142_
            nw589_ = _dafny.Array(None, 13)
            for i0_85_ in range(nw589_.length(0)):
                nw589_[i0_85_] = init85_(i0_85_)
            d_3642_v19_ = nw589_
            index628_ = default__.safeIndex(583, (d_3641_v18_).length(0))
            rhs609_ = (d_3638_v15_ if (self).f18 else d_3638_v15_)
            rhs610_ = d_3642_v19_
            lhs519_ = d_3641_v18_
            lhs520_ = default__.safeIndex(583, (d_3641_v18_).length(0))
            lhs519_[lhs520_] = rhs609_
            d_3642_v19_ = rhs610_
            d_3644_v20_: _dafny.Array
            nw590_ = _dafny.Array(int(0), 28)
            d_3644_v20_ = nw590_
            d_3645_v21_: _dafny.Array
            d_3645_v21_ = d_3644_v20_
            d_3646_v22_: _dafny.Set
            d_3646_v22_ = _dafny.Set({d_3645_v21_})
            d_3646_v22_ = (d_3646_v22_) - (_dafny.Set({d_3644_v20_, d_3645_v21_}))
            d_3647_v23_: D18
            d_3647_v23_ = D18_DC54(D18_DC52(d_3634_cf1_, (self).f18, (self).f17))
            d_3647_v23_ = d_3647_v23_
        elif True:
            d_3648___mcc_h2_ = source50_.cf0
            d_3649_cf0_ = d_3648___mcc_h2_
            if (self).f18:
                d_3650_v24_: _dafny.Set
                d_3650_v24_ = _dafny.Set({(self).f15, (self).f18, (self).f15})
                d_3651_v25_: _dafny.Map
                d_3651_v25_ = _dafny.Map({(self).f18: d_3650_v24_})
                d_3652_v26_: _dafny.Map
                d_3652_v26_ = _dafny.Map({default__.safeModuloInt((self).fm4((0) - ((self).f11), -729, (self).f18, True, globalState), len(d_3651_v25_)): (self).f12})
                d_3653_v27_: _dafny.Array
                nw591_ = _dafny.Array(int(0), 27)
                d_3653_v27_ = nw591_
                d_3654_v28_: _dafny.Seq
                d_3654_v28_ = _dafny.SeqWithoutIsStrInference([d_3653_v27_])
                d_3655_v29_: _dafny.Map
                d_3655_v29_ = _dafny.Map({d_3654_v28_: _dafny.Map({(self).f14: (self).f12})})
                d_3652_v26_ = ((d_3655_v29_)[(d_3654_v28_) + (d_3654_v28_)] if ((d_3654_v28_) + (d_3654_v28_)) in (d_3655_v29_) else d_3652_v26_)
                d_3656_v30_: str
                d_3656_v30_ = _dafny.CodePoint('f')
                d_3656_v30_ = d_3656_v30_
                d_3657_v31_: _dafny.Array
                def lambda143_(d_3658_v30_):
                    def lambda144_(d_3659_i3_):
                        return d_3658_v30_

                    return lambda144_

                init86_ = lambda143_(d_3656_v30_)
                nw592_ = _dafny.Array(None, 27)
                for i0_86_ in range(nw592_.length(0)):
                    nw592_[i0_86_] = init86_(i0_86_)
                d_3657_v31_ = nw592_
                index629_ = default__.safeIndex(645, (d_3657_v31_).length(0))
                (d_3657_v31_)[index629_] = d_3656_v30_
                index630_ = default__.safeIndex(645, (d_3657_v31_).length(0))
                (d_3657_v31_)[index630_] = d_3656_v30_
                d_3660_v32_: C3
                nw593_ = C3()
                nw593_.ctor__((self).f15, self.f16, self.f16, (self).f14, (_dafny.SeqWithoutIsStrInference([(d_3657_v31_)[default__.safeIndex(645, (d_3657_v31_).length(0))] for d_3661_i4_ in range(default__.abs(278))])) < ((self).f17), (self).f17, (0) - ((self).f12), (self).f12)
                d_3660_v32_ = nw593_
                d_3662_v33_: _dafny.Seq
                d_3662_v33_ = _dafny.SeqWithoutIsStrInference([d_3660_v32_])
                index631_ = default__.safeIndex(960, (p0).length(0))
                (p0)[index631_] = (d_3660_v32_).fm1(((self).f14) + (272), (self).f18, (d_3660_v32_) not in (d_3662_v33_), globalState)
            elif True:
                d_3663_v34_: _dafny.Seq
                d_3663_v34_ = _dafny.SeqWithoutIsStrInference([(self).f18])
                d_3664_v35_: D14
                d_3664_v35_ = D14_DC40(d_3663_v34_)
                d_3665_v36_: _dafny.Seq
                d_3665_v36_ = _dafny.SeqWithoutIsStrInference([d_3664_v35_, d_3664_v35_, D14_DC40(d_3663_v34_)])
                d_3665_v36_ = d_3665_v36_
                d_3666_v37_: C3
                nw594_ = C3()
                nw594_.ctor__(not((self).f15), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jm")), (self).f17, (self).f14, (p0)[default__.safeIndex(960, (p0).length(0))], (self).f17, (self).f11, (self).f12)
                d_3666_v37_ = nw594_
                d_3667_v38_: _dafny.Set
                d_3667_v38_ = _dafny.Set({d_3666_v37_})
                d_3668_v39_: _dafny.MultiSet
                d_3668_v39_ = _dafny.MultiSet([not(not((self).f15))])
                d_3669_v40_: C9
                nw595_ = C9()
                nw595_.ctor__(self.f16, not((self).f15), (320) * ((self).f14), (len(d_3667_v38_)) - ((d_3668_v39_).cardinality))
                d_3669_v40_ = nw595_
                d_3669_v40_ = d_3669_v40_
                (globalState).f6 = (0) - ((self).f11)
                d_3670_v41_: _dafny.Seq
                d_3670_v41_ = _dafny.SeqWithoutIsStrInference([(self).f12, (self).f11, (self).f14])
                (globalState).f7 = ((self).f12) < (len(d_3670_v41_))
                d_3671_v42_: D14
                d_3671_v42_ = D14_DC41(not(not(True)), (d_3666_v37_).f25, self.f16)
                (globalState).f2 = not (((self).f14) != ((self).f14)) or (((self).f12) == (len((d_3671_v42_).cf75)))
            d_3672_v43_: _dafny.Array
            nw596_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_3672_v43_ = nw596_
            d_3672_v43_ = d_3672_v43_
            (globalState).f2 = (p0)[default__.safeIndex(960, (p0).length(0))]
            d_3673_v45_: _dafny.Seq
            def iife223_():
                coll79_ = _dafny.Map()
                compr_79_: int
                for compr_79_ in _dafny.IntegerRange(923, 853):
                    d_3674_v44_: int = compr_79_
                    if ((923) <= (d_3674_v44_)) and ((d_3674_v44_) < (853)):
                        coll79_[default__.safeModuloInt(d_3674_v44_, (self).f11)] = True
                return _dafny.Map(coll79_)
            d_3673_v45_ = _dafny.SeqWithoutIsStrInference([iife223_()
            ])
            d_3675_v46_: C6
            nw597_ = C6()
            nw597_.ctor__(d_3673_v45_, (self).f12, (self).f15, (self).f13, 502, 569)
            d_3675_v46_ = nw597_
            d_3676_v47_: D24
            d_3676_v47_ = D24_DC66(d_3675_v46_, (self).f15)
            d_3677_v48_: _dafny.Map
            d_3677_v48_ = _dafny.Map({(d_3676_v47_).cf124: (self).f11})
            d_3678_v49_: _dafny.Map
            d_3678_v49_ = _dafny.Map({(self).f14: (d_3677_v48_).set((self).f18, -602)})
            d_3679_v50_: _dafny.Seq
            d_3679_v50_ = _dafny.SeqWithoutIsStrInference([-357, 847])
            d_3680_v51_: D4
            d_3680_v51_ = D4_DC13(d_3679_v50_)
            d_3678_v49_ = (d_3678_v49_).set(default__.safeDivisionInt((self).f11, (self).f14), (d_3677_v48_) | (_dafny.Map({(self).f15: len((d_3680_v51_).cf29)})))
        d_3681_v52_: _dafny.Map
        d_3681_v52_ = _dafny.Map({(self).f15: (self).f12})
        d_3682_v53_: _dafny.Seq
        d_3682_v53_ = _dafny.SeqWithoutIsStrInference([d_3681_v52_])
        d_3682_v53_ = d_3682_v53_
        d_3683_v54_: str
        d_3683_v54_ = _dafny.CodePoint('s')
        d_3684_v55_: _dafny.Seq
        d_3684_v55_ = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15, (self).f15])
        d_3685_v56_: _dafny.Map
        d_3685_v56_ = _dafny.Map({962: (self).f11})
        d_3686_v57_: C9
        nw598_ = C9()
        nw598_.ctor__(((self.f16) + ((self).f13)).set(default__.safeIndex(-713, len((self.f16) + ((self).f13))), d_3683_v54_), ((d_3684_v55_)[default__.safeIndex((self).f12, len(d_3684_v55_))]) and ((d_3684_v55_)[default__.safeIndex(191, len(d_3684_v55_))]), ((self).f11) + ((self).f12), len(d_3685_v56_))
        d_3686_v57_ = nw598_
        r0 = (((self).f11) + ((self).f14)) > ((self).f14)
        r1 = (p1).intersection(p1)
        return r0, r1

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: int = int(0)
        d_3687_v0_: _dafny.Map
        d_3687_v0_ = _dafny.Map({p0: (self).f12})
        d_3688_v1_: C5
        nw599_ = C5()
        nw599_.ctor__()
        d_3688_v1_ = nw599_
        d_3689_v2_: _dafny.Seq
        d_3689_v2_ = _dafny.SeqWithoutIsStrInference([d_3688_v1_, d_3688_v1_])
        d_3690_v3_: D3
        d_3690_v3_ = D3_DC10(p2, (0) - ((self).f12), d_3687_v0_, (self).f17, len(d_3689_v2_))
        hi14_ = (len((d_3690_v3_).cf21)) - (p3)
        for d_3691_i0_ in range(len(self.f16), hi14_):
            d_3692_v4_: C0
            nw600_ = C0()
            nw600_.ctor__()
            d_3692_v4_ = nw600_
            d_3693_v5_: _dafny.MultiSet
            d_3693_v5_ = _dafny.MultiSet([(self).f15])
            d_3694_v6_: _dafny.Seq
            d_3694_v6_ = _dafny.SeqWithoutIsStrInference([True])
            rhs611_ = p1
            rhs612_ = p2
            rhs613_ = _dafny.MultiSet(d_3694_v6_)
            rhs614_ = (self).f14
            lhs521_ = globalState
            lhs522_ = globalState
            r2 = rhs611_
            lhs521_.f2 = rhs612_
            d_3693_v5_ = rhs613_
            lhs522_.f6 = rhs614_
            d_3695_v7_: _dafny.Array
            def lambda145_(d_3696_p2_):
                def lambda146_(d_3697_i1_):
                    return not(d_3696_p2_)

                return lambda146_

            init87_ = lambda145_(p2)
            nw601_ = _dafny.Array(None, 13)
            for i0_87_ in range(nw601_.length(0)):
                nw601_[i0_87_] = init87_(i0_87_)
            d_3695_v7_ = nw601_
            d_3698_v8_: _dafny.Array
            nw602_ = _dafny.Array(None, 10)
            nw602_[int(0)] = d_3695_v7_
            nw602_[int(1)] = d_3695_v7_
            nw602_[int(2)] = d_3695_v7_
            nw602_[int(3)] = d_3695_v7_
            nw602_[int(4)] = d_3695_v7_
            nw602_[int(5)] = d_3695_v7_
            nw602_[int(6)] = d_3695_v7_
            nw602_[int(7)] = d_3695_v7_
            nw602_[int(8)] = d_3695_v7_
            nw602_[int(9)] = d_3695_v7_
            d_3698_v8_ = nw602_
            d_3699_v9_: _dafny.Array
            d_3699_v9_ = d_3698_v8_
            d_3699_v9_ = d_3699_v9_
            d_3700_v10_: _dafny.Map
            d_3700_v10_ = _dafny.Map({(self).f11: p2})
            d_3701_v11_: _dafny.Map
            d_3701_v11_ = _dafny.Map({(self).f14: len(d_3700_v10_)})
            r2 = ((d_3701_v11_)[(self).f14] if ((self).f14) in (d_3701_v11_) else ((self).f14 if (self).f18 else (self).f12))
        (globalState).f6 = ((0) - (default__.safeModuloInt(p1, len(_dafny.Set({p1}))))) + ((self).f12)
        d_3702_v12_: str
        d_3702_v12_ = _dafny.CodePoint('f')
        d_3703_v13_: _dafny.MultiSet
        d_3703_v13_ = _dafny.MultiSet([d_3702_v12_])
        d_3704_v14_: _dafny.MultiSet
        d_3704_v14_ = _dafny.MultiSet([p2])
        d_3705_v15_: _dafny.Map
        d_3705_v15_ = _dafny.Map({(self).f11: (self).f15})
        d_3706_v16_: _dafny.Seq
        d_3706_v16_ = _dafny.SeqWithoutIsStrInference([p2, ((d_3705_v15_)[652] if (652) in (d_3705_v15_) else p2)])
        d_3707_v17_: T3
        nw603_ = C3()
        nw603_.ctor__((self).f15, ((self).f13).set(default__.safeIndex(len(d_3706_v16_), len((self).f13)), d_3702_v12_), (self).f13, (self).f11, p2, p0, len((self).f17), len(p0))
        d_3707_v17_ = nw603_
        d_3708_v18_: _dafny.MultiSet
        d_3708_v18_ = _dafny.MultiSet([d_3707_v17_])
        d_3709_v19_: _dafny.Map
        d_3709_v19_ = _dafny.Map({(self).f11: (d_3707_v17_).f14})
        d_3710_v20_: _dafny.MultiSet
        d_3710_v20_ = _dafny.MultiSet([len(d_3709_v19_), (d_3707_v17_).f14])
        d_3711_v21_: _dafny.Seq
        d_3711_v21_ = _dafny.SeqWithoutIsStrInference([(self).f14])
        d_3712_v22_: _dafny.Map
        d_3712_v22_ = _dafny.Map({(self).f18: default__.fm44(len(_dafny.SeqWithoutIsStrInference([p1])), (d_3707_v17_).f15, globalState)})
        d_3713_v23_: _dafny.Map
        d_3713_v23_ = _dafny.Map({p2: ((d_3712_v22_)[p2] if (p2) in (d_3712_v22_) else _dafny.SeqWithoutIsStrInference([d_3702_v12_ for d_3714_i2_ in range(default__.abs(979))]))})
        d_3715_v24_: _dafny.Set
        d_3715_v24_ = _dafny.Set({(0) - ((d_3707_v17_).f12), (d_3707_v17_).f12, (0) - (len(d_3706_v16_)), len(_dafny.SeqWithoutIsStrInference([d_3702_v12_ for d_3716_i3_ in range(default__.abs(424))]))})
        d_3717_v25_: _dafny.Array
        nw604_ = _dafny.Array(None, 14)
        nw604_[int(0)] = (self).f18
        nw604_[int(1)] = not(p2)
        nw604_[int(2)] = (default__.fm69((self).f15, p3, d_3703_v13_, (self).f11, globalState)) and ((self).f15)
        nw604_[int(3)] = (self).f15
        nw604_[int(4)] = (self).f18
        nw604_[int(5)] = (d_3704_v14_).issubset(d_3704_v14_)
        nw604_[int(6)] = p2
        nw604_[int(7)] = (d_3708_v18_).ispropersubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_3707_v17_, d_3707_v17_, d_3707_v17_])))
        nw604_[int(8)] = (_dafny.MultiSet(d_3711_v21_)).ispropersubset(d_3710_v20_)
        nw604_[int(9)] = (d_3712_v22_) == (d_3712_v22_)
        nw604_[int(10)] = (self).f18
        nw604_[int(11)] = not(((d_3705_v15_)[(self).f14] if ((self).f14) in (d_3705_v15_) else (self).f18))
        nw604_[int(12)] = (d_3715_v24_).isdisjoint(d_3715_v24_)
        nw604_[int(13)] = p2
        d_3717_v25_ = nw604_
        index632_ = default__.safeIndex(495, (d_3717_v25_).length(0))
        (d_3717_v25_)[index632_] = ((d_3704_v14_).cardinality) != ((d_3707_v17_).f14)
        d_3718_v26_: _dafny.Set
        d_3718_v26_ = _dafny.Set({p2})
        index633_ = default__.safeIndex(495, (d_3717_v25_).length(0))
        (d_3717_v25_)[index633_] = not (True) or ((d_3718_v26_).ispropersubset(d_3718_v26_))
        d_3719_v27_: _dafny.Set
        d_3719_v27_ = _dafny.Set({d_3717_v25_})
        d_3720_v28_: D19
        d_3720_v28_ = D19_DC55(d_3719_v27_)
        source51_ = d_3720_v28_
        if source51_.is_DC56:
            d_3721___mcc_h0_ = source51_.cf105
            d_3722___mcc_h1_ = source51_.cf106
            d_3723___mcc_h2_ = source51_.cf107
            d_3724___mcc_h3_ = source51_.cf108
            d_3725_cf108_ = d_3724___mcc_h3_
            d_3726_cf107_ = d_3723___mcc_h2_
            d_3727_cf106_ = d_3722___mcc_h1_
            d_3728_cf105_ = d_3721___mcc_h0_
            (globalState).f2 = (p2) and ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrj"))) <= ((self).f13))
            if (d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))]:
                index634_ = default__.safeIndex(391, (d_3728_cf105_).length(0))
                (d_3728_cf105_)[index634_] = (d_3707_v17_).f11
                index635_ = default__.safeIndex(391, (d_3728_cf105_).length(0))
                (d_3728_cf105_)[index635_] = len(default__.fm74(False, d_3727_cf106_, (d_3707_v17_).f11, globalState))
                d_3727_cf106_ = (d_3706_v16_)[default__.safeIndex(default__.safeModuloInt(552, (d_3707_v17_).f11), len(d_3706_v16_))]
                (d_3707_v17_).f16 = ((p0) + (self.f16)) + ((self).f13)
                index636_ = default__.safeIndex(391, (d_3728_cf105_).length(0))
                (d_3728_cf105_)[index636_] = (self).f11
                r0 = not((d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))])
            elif True:
                d_3729_v29_: _dafny.MultiSet
                d_3729_v29_ = _dafny.MultiSet([d_3728_cf105_, d_3728_cf105_])
                d_3727_cf106_ = (len(d_3709_v19_)) < (((d_3729_v29_).cardinality if (self).f15 else (d_3707_v17_).f14))
                d_3730_v30_: D12
                d_3730_v30_ = D12_DC35(d_3718_v26_)
                d_3731_v31_: C2
                nw605_ = C2()
                nw605_.ctor__(len((d_3730_v30_).cf63), (d_3707_v17_).f13, (d_3707_v17_).f14, len(d_3711_v21_))
                d_3731_v31_ = nw605_
                nw606_ = C2()
                nw606_.ctor__((len(_dafny.SeqWithoutIsStrInference([d_3702_v12_ for d_3732_i4_ in range(default__.abs(815))]))) * (-68), ((d_3707_v17_).f13) + (self.f16), (d_3707_v17_).f14, 827)
                d_3731_v31_ = nw606_
                d_3733_v32_: _dafny.Set
                d_3733_v32_ = _dafny.Set({d_3710_v20_})
                index637_ = default__.safeIndex(495, (d_3717_v25_).length(0))
                rhs615_ = d_3733_v32_
                rhs616_ = d_3717_v25_
                rhs617_ = (d_3707_v17_).f15
                rhs618_ = default__.safeModuloInt((d_3707_v17_).f12, (d_3707_v17_).f12)
                lhs523_ = d_3717_v25_
                lhs524_ = default__.safeIndex(495, (d_3717_v25_).length(0))
                lhs525_ = d_3731_v31_
                d_3733_v32_ = rhs615_
                d_3717_v25_ = rhs616_
                lhs523_[lhs524_] = rhs617_
                lhs525_.f27 = rhs618_
                d_3734_v33_: _dafny.Map
                d_3734_v33_ = _dafny.Map({(self).f15: (d_3707_v17_).f15})
                d_3735_v34_: _dafny.Seq
                d_3735_v34_ = _dafny.SeqWithoutIsStrInference([default__.fm32(((d_3734_v33_)[(d_3707_v17_).f15] if ((d_3707_v17_).f15) in (d_3734_v33_) else p2), _dafny.Map({d_3702_v12_: (d_3707_v17_).f11}), d_3715_v24_, (self).f15, globalState), d_3718_v26_, (d_3718_v26_) - (d_3718_v26_), d_3718_v26_])
                d_3735_v34_ = (_dafny.SeqWithoutIsStrInference([d_3718_v26_, d_3718_v26_])) + (d_3735_v34_)
                d_3702_v12_ = d_3702_v12_
            r1 = 775
            (globalState).f6 = default__.safeModuloInt((d_3707_v17_).f11, (d_3707_v17_).f14)
        elif True:
            d_3736___mcc_h4_ = source51_.cf104
            d_3737_cf104_ = d_3736___mcc_h4_
            d_3738_v35_: D17
            d_3738_v35_ = D17_DC49(default__.fm0((self).f15, globalState))
            source52_ = d_3738_v35_
            if source52_.is_DC48:
                d_3739___mcc_h5_ = source52_.cf90
                d_3740___mcc_h6_ = source52_.cf91
                d_3741___mcc_h7_ = source52_.cf92
                d_3742_cf92_ = d_3741___mcc_h7_
                d_3743_cf91_ = d_3740___mcc_h6_
                d_3744_cf90_ = d_3739___mcc_h5_
                d_3745_v36_: C8
                nw607_ = C8()
                nw607_.ctor__((d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))], d_3744_cf90_)
                d_3745_v36_ = nw607_
                d_3746_v37_: D18
                d_3746_v37_ = D18_DC52(d_3743_cf91_, (d_3707_v17_).f15, (_dafny.SeqWithoutIsStrInference([d_3702_v12_ for d_3747_i5_ in range(default__.abs(199))])).set(default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference([(self).f12]))), len(_dafny.SeqWithoutIsStrInference([d_3702_v12_ for d_3748_i5_ in range(default__.abs(199))]))), d_3702_v12_))
                d_3749_v38_: _dafny.Map
                d_3749_v38_ = _dafny.Map({d_3746_v37_: False})
                d_3749_v38_ = (d_3749_v38_) | (d_3749_v38_)
                d_3750_v39_: _dafny.Map
                d_3750_v39_ = _dafny.Map({(self).f15: p2})
                d_3751_v40_: D14
                d_3751_v40_ = D14_DC41(((d_3750_v39_)[(d_3707_v17_).f15] if ((d_3707_v17_).f15) in (d_3750_v39_) else d_3745_v36_.f21), (d_3745_v36_).f22, (d_3707_v17_).f13)
                d_3687_v0_ = (d_3687_v0_).set((d_3751_v40_).cf75, p1)
                d_3752_v41_: D1
                d_3752_v41_ = D1_DC5((self).f12, (self).f18, p2)
                (globalState).f7 = (d_3752_v41_).cf8
            elif source52_.is_DC49:
                d_3753___mcc_h8_ = source52_.cf93
                d_3754_cf93_ = d_3753___mcc_h8_
                d_3755_v42_: _dafny.Map
                d_3755_v42_ = _dafny.Map({(d_3707_v17_).fm1((d_3707_v17_).f12, False, (d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))], globalState): (d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))]})
                d_3755_v42_ = (d_3755_v42_).set(default__.fm69((d_3707_v17_).f15, (self).f12, d_3703_v13_, len(d_3718_v26_), globalState), (d_3707_v17_).f15)
                d_3709_v19_ = (d_3709_v19_).set(len(_dafny.Map({(d_3707_v17_).f14: (d_3707_v17_).f15})), (d_3707_v17_).f14)
                (globalState).f2 = not(not((p1) > (d_3754_cf93_)))
                index638_ = default__.safeIndex(495, (d_3717_v25_).length(0))
                (d_3717_v25_)[index638_] = (d_3707_v17_).f15
            elif source52_.is_DC50:
                d_3756___mcc_h9_ = source52_.cf94
                d_3757_cf94_ = d_3756___mcc_h9_
                (globalState).f6 = (d_3707_v17_).f12
                d_3758_v43_: _dafny.Map
                d_3758_v43_ = _dafny.Map({(d_3715_v24_) - (d_3715_v24_): (self).f15})
                d_3758_v43_ = (d_3758_v43_).set((d_3715_v24_).intersection(d_3715_v24_), True)
                d_3759_v44_: C3
                nw608_ = C3()
                nw608_.ctor__(p2, (d_3707_v17_).f13, d_3707_v17_.f16, (0) - ((0) - ((self).f14)), (d_3707_v17_).f15, _dafny.SeqWithoutIsStrInference([d_3702_v12_ for d_3760_i6_ in range(default__.abs(212))]), (d_3707_v17_).f11, (self).f12)
                d_3759_v44_ = nw608_
                d_3761_v45_: D16
                d_3761_v45_ = D16_DC46((d_3718_v26_) | (d_3718_v26_), default__.fm8(p2, (d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))], (d_3707_v17_).f15, len(((self).f17).set(default__.safeIndex((d_3707_v17_).f12, len((self).f17)), d_3702_v12_)), globalState), d_3757_cf94_)
                d_3761_v45_ = d_3761_v45_
            elif True:
                d_3762___mcc_h10_ = source52_.cf89
                d_3763_cf89_ = d_3762___mcc_h10_
                index639_ = default__.safeIndex(495, (d_3717_v25_).length(0))
                (d_3717_v25_)[index639_] = False
                d_3764_v46_: D6
                d_3764_v46_ = D6_DC21((self).f18)
                d_3765_v47_: _dafny.Map
                d_3765_v47_ = _dafny.Map({p3: d_3702_v12_})
                d_3764_v46_ = D6_DC21((d_3707_v17_).fm1(len(d_3765_v47_), (d_3707_v17_).f15, (d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))], globalState))
                d_3766_v48_: int
                d_3767_v49_: _dafny.Seq
                out100_: int
                out101_: _dafny.Seq
                out100_, out101_ = (d_3707_v17_).m4(d_3717_v25_, globalState)
                d_3766_v48_ = out100_
                d_3767_v49_ = out101_
                (globalState).f6 = (d_3707_v17_).f11
            (globalState).f2 = (d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))]
            if not(((0) - ((d_3707_v17_).f14)) <= ((self).f12)):
                d_3768_v50_: D15
                d_3768_v50_ = D15_DC43((d_3707_v17_).f12, (self).f18, p2)
                rhs619_ = 247
                rhs620_ = d_3709_v19_
                rhs621_ = (d_3768_v50_).cf77
                lhs526_ = globalState
                lhs526_.f6 = rhs619_
                d_3709_v19_ = rhs620_
                r1 = rhs621_
                d_3769_v51_: _dafny.Array
                nw609_ = _dafny.Array(_dafny.CodePoint('D'), 23)
                d_3769_v51_ = nw609_
                index640_ = default__.safeIndex(668, (d_3769_v51_).length(0))
                (d_3769_v51_)[index640_] = d_3702_v12_
                index641_ = default__.safeIndex(668, (d_3769_v51_).length(0))
                (d_3769_v51_)[index641_] = d_3702_v12_
                d_3770_v52_: _dafny.Map
                d_3770_v52_ = _dafny.Map({(d_3707_v17_).f15: (d_3707_v17_).f12})
                d_3771_v53_: _dafny.Map
                d_3771_v53_ = _dafny.Map({default__.fm75((self).f15, globalState): (d_3707_v17_).f15})
                d_3772_v54_: _dafny.Map
                d_3772_v54_ = _dafny.Map({d_3717_v25_: (d_3707_v17_).f14})
                d_3773_v55_: _dafny.Map
                d_3773_v55_ = _dafny.Map({d_3717_v25_: (d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))]})
                index642_ = default__.safeIndex(495, (d_3717_v25_).length(0))
                rhs622_ = d_3770_v52_
                rhs623_ = (default__.safeModuloInt(len(d_3770_v52_), (0) - ((d_3710_v20_).cardinality))) + ((self).fm2((d_3707_v17_).f11, (d_3707_v17_).f15, (d_3707_v17_).f11, True, globalState))
                rhs624_ = default__.safeModuloInt(len((d_3772_v54_) | (d_3772_v54_)), p1)
                rhs625_ = ((d_3773_v55_)[d_3717_v25_] if (d_3717_v25_) in (d_3773_v55_) else p2)
                rhs626_ = default__.fm76(len((d_3709_v19_ if (d_3707_v17_).f15 else d_3709_v19_)), globalState)
                lhs527_ = globalState
                lhs528_ = globalState
                lhs529_ = d_3717_v25_
                lhs530_ = default__.safeIndex(495, (d_3717_v25_).length(0))
                d_3770_v52_ = rhs622_
                lhs527_.f6 = rhs623_
                lhs528_.f6 = rhs624_
                lhs529_[lhs530_] = rhs625_
                d_3771_v53_ = rhs626_
                d_3711_v21_ = (d_3711_v21_) + (d_3711_v21_)
                (globalState).f2 = (d_3707_v17_).f15
            elif True:
                d_3774_v56_: _dafny.Map
                d_3774_v56_ = _dafny.Map({not(p2): (d_3707_v17_).f15})
                (globalState).f2 = not (((d_3774_v56_)[(d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))]] if ((d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))]) in (d_3774_v56_) else (d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))])) or (((d_3705_v15_)[(d_3707_v17_).f12] if ((d_3707_v17_).f12) in (d_3705_v15_) else (d_3707_v17_).f15))
                d_3775_v57_: _dafny.Map
                d_3775_v57_ = _dafny.Map({(self).f18: (d_3707_v17_).f11})
                d_3775_v57_ = (d_3775_v57_).set((d_3715_v24_).isdisjoint(d_3715_v24_), (d_3707_v17_).f12)
                d_3776_v58_: _dafny.Array
                def lambda147_(d_3777_v17_):
                    def lambda148_(d_3778_i7_):
                        return (d_3778_i7_) + ((d_3777_v17_).f14)

                    return lambda148_

                init88_ = lambda147_(d_3707_v17_)
                nw610_ = _dafny.Array(None, 29)
                for i0_88_ in range(nw610_.length(0)):
                    nw610_[i0_88_] = init88_(i0_88_)
                d_3776_v58_ = nw610_
                index643_ = default__.safeIndex(300, (d_3776_v58_).length(0))
                (d_3776_v58_)[index643_] = 880
                d_3779_v59_: D4
                d_3779_v59_ = D4_DC15(((d_3704_v14_)[(d_3707_v17_).f15] if ((d_3707_v17_).f15) in (d_3704_v14_) else -979))
                index644_ = default__.safeIndex(300, (d_3776_v58_).length(0))
                (d_3776_v58_)[index644_] = (d_3779_v59_).cf32
                d_3780_v60_: _dafny.Map
                d_3780_v60_ = _dafny.Map({(self).f18: d_3702_v12_})
                d_3780_v60_ = d_3780_v60_
                d_3781_v61_: _dafny.Set
                d_3781_v61_ = _dafny.Set({d_3704_v14_})
                d_3782_v62_: _dafny.Seq
                d_3782_v62_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_3706_v16_)])
                index645_ = default__.safeIndex(300, (d_3776_v58_).length(0))
                index646_ = default__.safeIndex(300, (d_3776_v58_).length(0))
                def iife224_():
                    coll80_ = _dafny.Set()
                    compr_80_: _dafny.MultiSet
                    for compr_80_ in (d_3782_v62_).Elements:
                        d_3783_v63_: _dafny.MultiSet = compr_80_
                        if (d_3783_v63_) in (d_3782_v62_):
                            coll80_ = coll80_.union(_dafny.Set([d_3783_v63_]))
                    return _dafny.Set(coll80_)
                rhs627_ = (0) - (p1)
                rhs628_ = (iife224_()
                ) - ((d_3781_v61_).intersection(d_3781_v61_))
                rhs629_ = -994
                lhs531_ = d_3776_v58_
                lhs532_ = default__.safeIndex(300, (d_3776_v58_).length(0))
                lhs533_ = d_3776_v58_
                lhs534_ = default__.safeIndex(300, (d_3776_v58_).length(0))
                lhs531_[lhs532_] = rhs627_
                d_3781_v61_ = rhs628_
                lhs533_[lhs534_] = rhs629_
            d_3784_v64_: _dafny.Array
            nw611_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_3784_v64_ = nw611_
            d_3785_v65_: _dafny.Array
            nw612_ = _dafny.Array(None, 13)
            d_3785_v65_ = nw612_
            index647_ = default__.safeIndex(385, (d_3784_v64_).length(0))
            (d_3784_v64_)[index647_] = d_3785_v65_
            index648_ = default__.safeIndex(385, (d_3784_v64_).length(0))
            rhs630_ = (0) - (((d_3710_v20_)[(self).f11] if ((self).f11) in (d_3710_v20_) else 269))
            rhs631_ = (d_3707_v17_).f14
            rhs632_ = d_3785_v65_
            lhs535_ = globalState
            lhs536_ = d_3784_v64_
            lhs537_ = default__.safeIndex(385, (d_3784_v64_).length(0))
            r2 = rhs630_
            lhs535_.f6 = rhs631_
            lhs536_[lhs537_] = rhs632_
        d_3786_v66_: C4
        nw613_ = C4()
        nw613_.ctor__((self).f13, (self).f12, p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")), (d_3707_v17_).f11, len(_dafny.Set({d_3702_v12_, default__.fm14((self).f15, _dafny.CodePoint('g'), (self).f12, (d_3707_v17_).f15, globalState), d_3702_v12_, d_3702_v12_})))
        d_3786_v66_ = nw613_
        d_3787_v67_: _dafny.Set
        d_3787_v67_ = _dafny.Set({d_3786_v66_, d_3786_v66_})
        hi15_ = 784
        for d_3788_i8_ in range((0) - (((0) - ((d_3707_v17_).f12) if not(default__.fm7((d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))], ((d_3704_v14_)[(d_3707_v17_).f15] if ((d_3707_v17_).f15) in (d_3704_v14_) else (self).f11), globalState)) else len(d_3787_v67_))), hi15_):
            r0 = (default__.fm0((self).f15, globalState)) == (default__.safeDivisionInt((self).f11, p1))
            d_3789_v68_: C2
            nw614_ = C2()
            nw614_.ctor__((self).f12, p0, ((d_3707_v17_).f12) + ((self).f11), (d_3707_v17_).f14)
            d_3789_v68_ = nw614_
            d_3789_v68_ = d_3789_v68_
            d_3790_v69_: _dafny.Array
            nw615_ = _dafny.Array(int(0), 4)
            d_3790_v69_ = nw615_
            d_3791_v70_: _dafny.Map
            d_3791_v70_ = _dafny.Map({d_3718_v26_: d_3790_v69_})
            d_3792_v71_: _dafny.Map
            d_3792_v71_ = _dafny.Map({p1: d_3790_v69_})
            d_3793_v72_: _dafny.Map
            d_3793_v72_ = _dafny.Map({(d_3707_v17_).f15: d_3718_v26_})
            d_3794_v73_: _dafny.Seq
            d_3794_v73_ = _dafny.SeqWithoutIsStrInference([d_3791_v70_, _dafny.Map({d_3718_v26_: d_3790_v69_}), d_3791_v70_, d_3791_v70_])
            d_3795_v74_: _dafny.Array
            nw616_ = _dafny.Array(None, 22)
            nw616_[int(0)] = (d_3791_v70_) | ((d_3791_v70_).set(d_3718_v26_, d_3790_v69_))
            nw616_[int(1)] = (d_3791_v70_) | (_dafny.Map({d_3718_v26_: ((d_3792_v71_)[len(d_3718_v26_)] if (len(d_3718_v26_)) in (d_3792_v71_) else d_3790_v69_)}))
            nw616_[int(2)] = d_3791_v70_
            nw616_[int(3)] = (_dafny.Map({((d_3793_v72_)[(d_3707_v17_).f15] if ((d_3707_v17_).f15) in (d_3793_v72_) else d_3718_v26_): d_3790_v69_})) | (d_3791_v70_)
            nw616_[int(4)] = d_3791_v70_
            nw616_[int(5)] = d_3791_v70_
            nw616_[int(6)] = d_3791_v70_
            nw616_[int(7)] = d_3791_v70_
            nw616_[int(8)] = d_3791_v70_
            nw616_[int(9)] = d_3791_v70_
            nw616_[int(10)] = (d_3791_v70_) | (d_3791_v70_)
            nw616_[int(11)] = (d_3791_v70_) | (d_3791_v70_)
            nw616_[int(12)] = (D27_DC71(_dafny.Map({d_3718_v26_: d_3790_v69_}))).cf130
            nw616_[int(13)] = d_3791_v70_
            nw616_[int(14)] = d_3791_v70_
            nw616_[int(15)] = d_3791_v70_
            nw616_[int(16)] = (d_3791_v70_) | (d_3791_v70_)
            nw616_[int(17)] = _dafny.Map({_dafny.Set({(d_3707_v17_).f15}): d_3790_v69_})
            nw616_[int(18)] = (_dafny.Map({d_3718_v26_: d_3790_v69_})) | ((d_3794_v73_)[default__.safeIndex((self).f14, len(d_3794_v73_))])
            nw616_[int(19)] = d_3791_v70_
            nw616_[int(20)] = (d_3791_v70_) | (d_3791_v70_)
            nw616_[int(21)] = d_3791_v70_
            d_3795_v74_ = nw616_
            index649_ = default__.safeIndex(772, (d_3795_v74_).length(0))
            (d_3795_v74_)[index649_] = d_3791_v70_
            index650_ = default__.safeIndex(772, (d_3795_v74_).length(0))
            (d_3795_v74_)[index650_] = (d_3791_v70_) | ((d_3791_v70_).set(d_3718_v26_, d_3790_v69_))
            d_3796_v75_: _dafny.Set
            d_3796_v75_ = _dafny.Set({(d_3707_v17_).f13, self.f16})
            d_3797_v76_: C8
            nw617_ = C8()
            nw617_.ctor__((_dafny.Set({(d_3707_v17_).f13})).isdisjoint(d_3796_v75_), not((d_3707_v17_).f15))
            d_3797_v76_ = nw617_
        hi16_ = p1
        for d_3798_i9_ in range(p3, hi16_):
            d_3799_v77_: C2
            nw618_ = C2()
            nw618_.ctor__((self).f14, (d_3707_v17_).f13, 229, (d_3707_v17_).f11)
            d_3799_v77_ = nw618_
            d_3800_v78_: _dafny.Map
            d_3800_v78_ = _dafny.Map({d_3799_v77_: (self).f13})
            d_3801_v79_: D0
            d_3801_v79_ = D0_DC0(d_3703_v13_)
            d_3802_v80_: _dafny.Map
            d_3802_v80_ = _dafny.Map({(d_3703_v13_).cardinality: d_3717_v25_})
            d_3803_v81_: D4
            d_3803_v81_ = D4_DC14(d_3702_v12_, (d_3707_v17_).f12)
            d_3804_v82_: _dafny.Array
            nw619_ = _dafny.Array(None, 15)
            nw619_[int(0)] = len(d_3705_v15_)
            nw619_[int(1)] = d_3798_i9_
            nw619_[int(2)] = default__.fm0(not((self).f18), globalState)
            nw619_[int(3)] = (0) - (len((((d_3800_v78_)[d_3799_v77_] if (d_3799_v77_) in (d_3800_v78_) else p0)).set(default__.safeIndex((0) - ((d_3707_v17_).f14), len(((d_3800_v78_)[d_3799_v77_] if (d_3799_v77_) in (d_3800_v78_) else p0))), d_3702_v12_)))
            nw619_[int(4)] = (d_3707_v17_).f11
            nw619_[int(5)] = ((d_3707_v17_).f14) - ((self).f14)
            nw619_[int(6)] = p1
            nw619_[int(7)] = ((d_3687_v0_)[default__.fm31(d_3702_v12_, d_3801_v79_, globalState)] if (default__.fm31(d_3702_v12_, d_3801_v79_, globalState)) in (d_3687_v0_) else len(d_3718_v26_))
            nw619_[int(8)] = ((d_3707_v17_).f12) * (d_3799_v77_.f27)
            nw619_[int(9)] = (d_3707_v17_).f11
            nw619_[int(10)] = default__.safeModuloInt((0) - ((self).f11), (d_3707_v17_).f12)
            nw619_[int(11)] = (d_3707_v17_).f12
            nw619_[int(12)] = 13
            nw619_[int(13)] = (self).f14
            nw619_[int(14)] = (self).fm4(len(d_3802_v80_), (d_3803_v81_).cf31, False, False, globalState)
            d_3804_v82_ = nw619_
            d_3804_v82_ = d_3804_v82_
            (globalState).f6 = 609
            (globalState).f2 = (d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))]
            if (self).f18:
                (globalState).f2 = ((d_3707_v17_).f12) > ((0) - (((self).f14) + ((self).f12)))
                d_3805_v83_: _dafny.Map
                d_3805_v83_ = _dafny.Map({d_3709_v19_: d_3799_v77_.f27})
                (globalState).f2 = (len(d_3805_v83_)) >= ((d_3707_v17_).f11)
                (globalState).f6 = (d_3707_v17_).f12
                (globalState).f6 = p1
                d_3806_v84_: D2
                d_3806_v84_ = D2_DC8(not((self).f15), p2, _dafny.SeqWithoutIsStrInference([d_3702_v12_ for d_3807_i10_ in range(default__.abs(364))]))
                d_3808_v85_: _dafny.Map
                d_3808_v85_ = _dafny.Map({p2: (d_3707_v17_).f14})
                d_3809_v86_: _dafny.Map
                d_3809_v86_ = _dafny.Map({d_3808_v85_: 483})
                d_3810_v87_: C2
                nw620_ = C2()
                nw620_.ctor__(d_3798_i9_, (d_3806_v84_).cf17, len(d_3809_v86_), (self).f11)
                d_3810_v87_ = nw620_
            elif True:
                r1 = d_3799_v77_.f27
                (globalState).f2 = ((d_3707_v17_).f15 if ((d_3707_v17_).f15) or ((d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))]) else True)
                (d_3799_v77_).f27 = (0) - (p3)
                d_3811_v88_: _dafny.Map
                d_3811_v88_ = _dafny.Map({d_3704_v14_: p1})
                d_3811_v88_ = (d_3811_v88_).set(d_3704_v14_, d_3798_i9_)
                d_3812_v89_: _dafny.Seq
                d_3812_v89_ = _dafny.SeqWithoutIsStrInference([d_3786_v66_])
                rhs633_ = (_dafny.Set({False, (d_3707_v17_).f15, False})) | (d_3718_v26_)
                rhs634_ = _dafny.MultiSet([(d_3812_v89_) < (d_3812_v89_), p2, (d_3707_v17_).f15, (d_3717_v25_)[default__.safeIndex(495, (d_3717_v25_).length(0))], (self).f18])
                rhs635_ = ((d_3707_v17_).f12) * ((d_3707_v17_).f12)
                lhs538_ = d_3799_v77_
                d_3718_v26_ = rhs633_
                d_3704_v14_ = rhs634_
                lhs538_.f27 = rhs635_
        d_3813_v90_: _dafny.MultiSet
        d_3813_v90_ = _dafny.MultiSet([(d_3707_v17_).f13])
        d_3814_v91_: _dafny.Seq
        d_3814_v91_ = _dafny.SeqWithoutIsStrInference([d_3813_v90_, d_3813_v90_])
        r0 = not (((d_3814_v91_)[default__.safeIndex((d_3707_v17_).f14, len(d_3814_v91_))]).isdisjoint(d_3813_v90_)) or ((self).f18)
        r1 = (d_3707_v17_).f11
        r2 = 833
        return r0, r1, r2

    def m5(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: bool = False
        d_3815_v0_: str
        d_3815_v0_ = _dafny.CodePoint('k')
        d_3816_v1_: _dafny.MultiSet
        d_3816_v1_ = _dafny.MultiSet([d_3815_v0_, d_3815_v0_])
        d_3817_v2_: D0
        d_3817_v2_ = D0_DC0(d_3816_v1_)
        rhs636_ = default__.fm31(_dafny.CodePoint('g'), d_3817_v2_, globalState)
        rhs637_ = default__.safeDivisionInt((self).f12, (self).f11)
        lhs539_ = self
        lhs540_ = globalState
        lhs539_.f16 = rhs636_
        lhs540_.f6 = rhs637_
        (globalState).f6 = (self).f11
        d_3818_v3_: _dafny.MultiSet
        d_3818_v3_ = _dafny.MultiSet([(self).f12, (self).f12, (self).f12])
        r2 = (_dafny.MultiSet([570])).issubset(d_3818_v3_)
        d_3819_v4_: _dafny.Array
        def lambda149_(d_3820_i1_):
            return (_dafny.Set({(self).f12, (self).f14, (self).f11, (self).f14})).isdisjoint(_dafny.Set({(self).f14}))

        init89_ = lambda149_
        nw621_ = _dafny.Array(None, 21)
        for i0_89_ in range(nw621_.length(0)):
            nw621_[i0_89_] = init89_(i0_89_)
        d_3819_v4_ = nw621_
        guard_loop_15_: int
        for guard_loop_15_ in _dafny.IntegerRange(0, (d_3819_v4_).length(0)):
            d_3821_i0_: int = guard_loop_15_
            if (True) and (((0) <= (d_3821_i0_)) and ((d_3821_i0_) < ((d_3819_v4_).length(0)))):
                (d_3819_v4_)[(d_3821_i0_)] = ((self).f15 if ((self).f15) or ((self).f18) else True)
        d_3822_v5_: _dafny.Array
        nw622_ = _dafny.Array(_dafny.Array(None, 0), 16)
        d_3822_v5_ = nw622_
        d_3823_v6_: _dafny.Array
        d_3823_v6_ = d_3822_v5_
        source53_ = d_3823_v6_
        d_3824___mcc_h0_ = source53_
        d_3825_cf121_ = d_3824___mcc_h0_
        d_3826_v7_: _dafny.Array
        nw623_ = _dafny.Array(_dafny.Map({}), 9)
        d_3826_v7_ = nw623_
        d_3827_v8_: _dafny.Map
        d_3827_v8_ = _dafny.Map({d_3815_v0_: d_3826_v7_})
        d_3827_v8_ = (d_3827_v8_).set(d_3815_v0_, d_3826_v7_)
        d_3828_v9_: _dafny.Map
        d_3828_v9_ = _dafny.Map({(self).f15: default__.fm77(globalState)})
        d_3829_v10_: _dafny.Map
        d_3829_v10_ = _dafny.Map({(self).f11: (self.f16)[default__.safeIndex((self).f12, len(self.f16))]})
        d_3828_v9_ = (default__.fm78(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxjh")), (self).f15, globalState)).set(True, d_3829_v10_)
        d_3830_v11_: _dafny.Map
        d_3830_v11_ = _dafny.Map({(self).f11: (self).f11})
        d_3831_v12_: _dafny.Map
        d_3831_v12_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gyr")): (self).f11})
        d_3832_v13_: _dafny.Map
        d_3832_v13_ = _dafny.Map({(self).f11: d_3831_v12_})
        d_3833_v14_: D3
        d_3833_v14_ = D3_DC10((self).f15, (self).f11, ((d_3832_v13_)[len(self.f16)] if (len(self.f16)) in (d_3832_v13_) else d_3831_v12_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajpcub")), (self).f11)
        d_3834_v15_: _dafny.Array
        nw624_ = _dafny.Array(None, 15)
        nw624_[int(0)] = D3_DC10(default__.fm7((self).f18, (self).f12, globalState), (self).fm4((0) - ((self).f14), (self).f14, (self).f18, (self).f15, globalState), _dafny.Map({(self).f17: (self).f14}), self.f16, len(d_3830_v11_))
        nw624_[int(1)] = d_3833_v14_
        nw624_[int(2)] = d_3833_v14_
        nw624_[int(3)] = d_3833_v14_
        nw624_[int(4)] = d_3833_v14_
        nw624_[int(5)] = d_3833_v14_
        nw624_[int(6)] = D3_DC10((self).f18, (self).f14, d_3831_v12_, (self).f17, (self).f11)
        nw624_[int(7)] = d_3833_v14_
        nw624_[int(8)] = d_3833_v14_
        nw624_[int(9)] = d_3833_v14_
        nw624_[int(10)] = d_3833_v14_
        nw624_[int(11)] = (d_3833_v14_ if not((self).f18) else d_3833_v14_)
        nw624_[int(12)] = d_3833_v14_
        nw624_[int(13)] = d_3833_v14_
        nw624_[int(14)] = d_3833_v14_
        d_3834_v15_ = nw624_
        index651_ = default__.safeIndex(132, (d_3834_v15_).length(0))
        (d_3834_v15_)[index651_] = d_3833_v14_
        index652_ = default__.safeIndex(132, (d_3834_v15_).length(0))
        rhs638_ = (self).f18
        rhs639_ = not((self).f18)
        rhs640_ = (self).f18
        rhs641_ = d_3833_v14_
        lhs541_ = globalState
        lhs542_ = globalState
        lhs543_ = d_3834_v15_
        lhs544_ = default__.safeIndex(132, (d_3834_v15_).length(0))
        lhs541_.f7 = rhs638_
        lhs542_.f7 = rhs639_
        r2 = rhs640_
        lhs543_[lhs544_] = rhs641_
        d_3835_v16_: _dafny.Seq
        d_3835_v16_ = _dafny.SeqWithoutIsStrInference([(self).f11])
        d_3836_v17_: _dafny.MultiSet
        d_3836_v17_ = _dafny.MultiSet([(((self).f17).set(default__.safeIndex(len((self).f17), len((self).f17)), d_3815_v0_)).set(default__.safeIndex(len(d_3835_v16_), len(((self).f17).set(default__.safeIndex(len((self).f17), len((self).f17)), d_3815_v0_))), d_3815_v0_)])
        d_3837_v18_: D20
        d_3837_v18_ = D20_DC57(d_3836_v17_)
        d_3837_v18_ = (d_3837_v18_ if (self).f18 else d_3837_v18_)
        d_3838_i2_: int
        d_3838_i2_ = 0
        with _dafny.label("20"):
            while (self).fm1((self).f11, (self).f15, not((self).f15), globalState):
                with _dafny.c_label("20"):
                    if (d_3838_i2_) >= (100):
                        raise _dafny.Break("20")
                    d_3838_i2_ = (d_3838_i2_) + (1)
                    (globalState).f6 = (self).f11
                    rhs642_ = (self).f18
                    rhs643_ = ((self).f14) * (648)
                    lhs545_ = globalState
                    lhs546_ = globalState
                    lhs545_.f7 = rhs642_
                    lhs546_.f6 = rhs643_
                    d_3839_v19_: _dafny.Map
                    d_3839_v19_ = _dafny.Map({(self).f11: (self).f15})
                    (globalState).f6 = len(d_3839_v19_)
                    d_3840_v20_: _dafny.Map
                    d_3840_v20_ = _dafny.Map({(self).f15: (self).f14})
                    d_3841_v21_: _dafny.Set
                    d_3841_v21_ = _dafny.Set({(self).fm2((self).f14, not((self).f18), len(_dafny.Map({(self).f12: (self).f18})), (self).f18, globalState), (self).f14, len(d_3840_v20_), ((d_3840_v20_)[(self).f18] if ((self).f18) in (d_3840_v20_) else (self).f12)})
                    d_3842_v22_: _dafny.Map
                    d_3842_v22_ = _dafny.Map({((self).fm2((0) - (-311), (self).f18, (self).f14, (self).f18, globalState)) * ((self).f11): len(d_3841_v21_)})
                    (globalState).f6 = len(d_3842_v22_)
                    pass
            pass
        d_3843_v23_: _dafny.Set
        d_3843_v23_ = _dafny.Set({(self).f18})
        r0 = d_3843_v23_
        r1 = (self).f15
        d_3844_v24_: _dafny.Map
        d_3844_v24_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_3815_v0_ for d_3845_i3_ in range(default__.abs(420))]): d_3815_v0_})
        r2 = (self.f16) not in (d_3844_v24_)
        return r0, r1, r2

    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
