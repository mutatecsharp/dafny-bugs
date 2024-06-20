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
        return (D21_DC62(_dafny.Map({len(_dafny.Map({False: True})): True}), False)).cf105

    @staticmethod
    def fm1(globalState):
        return 969

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_0_i0_ in range(default__.abs(183))])).Elements:
                d_1_v0_: _dafny.Seq = compr_0_
                if (d_1_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_0_i0_ in range(default__.abs(183))])):
                    coll0_[d_1_v0_] = _dafny.Map({not(True): len(_dafny.Map({len(_dafny.Set({len(_dafny.Set({_dafny.CodePoint('h')}))})): (_dafny.MultiSet([_dafny.Map({False: True}), _dafny.Map({True: True}), _dafny.Map({True: False})])).cardinality}))})
            return _dafny.Map(coll0_)
        return (iife0_()
        ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, False]): _dafny.Map({False: 760})}))

    @staticmethod
    def fm3(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfo"))): -676})).keys.Elements:
                d_3_v0_: int = compr_1_
                if (d_3_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfo"))): -676})):
                    coll1_[(d_3_v0_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_4_i1_ in range(default__.abs(141))])))] = True
            return _dafny.Map(coll1_)
        return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2_i0_ in range(default__.abs(409))])), len((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), _dafny.CodePoint('m')])): not(False)})) | (iife1_()
        )), (0) - ((765) + (len(_dafny.Map({not(True): 129})))), default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ny"))), 382), -840])

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(33, 380):
                d_5_v0_: int = compr_2_
                if ((33) <= (d_5_v0_)) and ((d_5_v0_) < (380)):
                    coll2_[(d_5_v0_) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), True, True, True]))).cardinality)] = len(_dafny.Set({839, len(_dafny.Map({True: not(True)}))}))
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: _dafny.Seq
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, not(True)])])).Elements:
                d_6_v1_: _dafny.Seq = compr_3_
                if (d_6_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, not(True)])])):
                    coll3_ = coll3_.union(_dafny.Set([d_6_v1_]))
            return _dafny.Set(coll3_)
        if (_dafny.MultiSet([D7_DC22(_dafny.Map({len(_dafny.Set({449})): 478}))])).issubset(_dafny.MultiSet([D7_DC22(_dafny.Map({len((D27_DC72(_dafny.SeqWithoutIsStrInference([405]))).cf123): 445})), D7_DC22(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, False, False, False]), _dafny.SeqWithoutIsStrInference([False, True])])): False})): len(_dafny.SeqWithoutIsStrInference([52]))})), D7_DC22(iife2_()
), D7_DC22(_dafny.Map({len(_dafny.Map({False: False})): (0) - (len(_dafny.SeqWithoutIsStrInference([257, len(_dafny.Set({len(_dafny.Set({False, True})), len(_dafny.SeqWithoutIsStrInference([True])), len(iife3_()
)}))])))}))])):
            return _dafny.CodePoint('f')
        elif True:
            return _dafny.CodePoint('m')

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return ((_dafny.Map({True: False}) if True else _dafny.Map({True: False}))) | ((_dafny.Map({False: False})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txrnv")) if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvkpmu"))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwry"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_7_i0_ in range(default__.abs(-302))])), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_8_i1_ in range(default__.abs(595))]))])

    @staticmethod
    def fm13(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yuwq"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_9_i0_ in range(default__.abs(571))]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return _dafny.MultiSet([(len(_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference([not(False)]))}))) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "taessl")))), 129, (125) + (-426), (736) + (len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True]))})))])

    @staticmethod
    def fm19(p0, p1, p2, p3, globalState):
        source0_ = D18_DC56(_dafny.SeqWithoutIsStrInference([True]))
        if source0_.is_DC54:
            d_10___mcc_h0_ = source0_.cf91
            d_11___mcc_h1_ = source0_.cf92
            d_12___mcc_h2_ = source0_.cf93
            d_13___mcc_h3_ = source0_.cf94
            d_14_cf94_ = d_13___mcc_h3_
            d_15_cf93_ = d_12___mcc_h2_
            d_16_cf92_ = d_11___mcc_h1_
            d_17_cf91_ = d_10___mcc_h0_
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])]))).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(D27_DC73(d_15_cf93_, d_16_cf92_, _dafny.Set({d_15_cf93_, True}))).cf124, d_15_cf93_])]))
        elif source0_.is_DC55:
            d_18___mcc_h4_ = source0_.cf95
            d_19___mcc_h5_ = source0_.cf96
            d_20_cf96_ = d_19___mcc_h5_
            d_21_cf95_ = d_18___mcc_h4_
            if False:
                return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, False])])
            elif True:
                return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True, True]), _dafny.SeqWithoutIsStrInference([True])])
        elif source0_.is_DC56:
            d_22___mcc_h6_ = source0_.cf97
            d_23_cf97_ = d_22___mcc_h6_
            return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False]), d_23_cf97_, d_23_cf97_])
        elif True:
            d_24___mcc_h7_ = source0_.cf90
            d_25_cf90_ = d_24___mcc_h7_
            return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])]))

    @staticmethod
    def fm24(globalState):
        return (_dafny.MultiSet([True])) | ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True)]))) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True]))))

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))

    @staticmethod
    def fm26(globalState):
        return ((_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, True])).cardinality])))})) | (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([-77, 196]))}))) - ((_dafny.Set({382})) - (_dafny.Set({-244})))

    @staticmethod
    def fm28(p0, p1, globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: _dafny.Seq
            for compr_4_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxusns"))})).Elements:
                d_26_v0_: _dafny.Seq = compr_4_
                if (d_26_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxusns"))})):
                    coll4_ = coll4_.union(_dafny.Set([d_26_v0_]))
            return _dafny.Set(coll4_)
        return iife4_()
        

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        return D3_DC11((len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqsfuxsch"))) for d_27_i0_ in range(default__.abs(916))]))) - (len(_dafny.Set({not(not(True))}))), (_dafny.Map({True: False})) | (_dafny.Map({False: False})))

    @staticmethod
    def fm30(globalState):
        return _dafny.Set({not(not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtved"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhtg"))))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hd")))})

    @staticmethod
    def fm31(p0, globalState):
        return (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([(D19_DC58(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hekts"))), not(True))).cf100])), 275})).intersection(_dafny.Set({len((D28_DC76(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([344]))}))).cf134)}))

    @staticmethod
    def fm32(p0, p1, p2, globalState):
        return D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mijxeu")))

    @staticmethod
    def fm33(p0, p1, p2, p3, globalState):
        return (_dafny.Map({False: False})) | ((_dafny.Map({not(False): True})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm34(p0, globalState):
        return D2_DC8(D2_DC7(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_28_i0_ in range(default__.abs(914))])), True, len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjqrabqx")))), len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({789: not(True)}))])), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([False, False]))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({111, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l'), _dafny.CodePoint('q')]))}))]))]))])), -233, -262]))))

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        return _dafny.Set({451, (len(_dafny.Set({(0) - ((_dafny.MultiSet([False, not(not(False))])).cardinality)}))) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_29_i0_ in range(default__.abs(-630))])))})

    @staticmethod
    def fm38(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_30_i0_ in range(default__.abs(780))]) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrwuvm")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbhrr"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctmp"))))

    @staticmethod
    def fm39(p0, p1, p2, globalState):
        return D2_DC6(_dafny.CodePoint('k'))

    @staticmethod
    def fm40(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm41(globalState):
        return (_dafny.MultiSet([False])).intersection(_dafny.MultiSet([False, True, True, False, False]))

    @staticmethod
    def fm42(globalState):
        source1_ = D7_DC24(-767)
        if source1_.is_DC23:
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(-528, -312):
                    d_31_v0_: int = compr_5_
                    if ((-528) <= (d_31_v0_)) and ((d_31_v0_) < (-312)):
                        coll5_[default__.safeModuloInt(d_31_v0_, 66)] = not(True)
                return _dafny.Map(coll5_)
            return D11_DC32(iife5_()
)
        elif source1_.is_DC24:
            d_32___mcc_h0_ = source1_.cf41
            d_33_cf41_ = d_32___mcc_h0_
            return D11_DC32(_dafny.Map({d_33_cf41_: True}))
        elif True:
            d_34___mcc_h1_ = source1_.cf40
            d_35_cf40_ = d_34___mcc_h1_
            return D11_DC32(_dafny.Map({len(_dafny.Set({True, True})): not(True)}))

    @staticmethod
    def fm43(p0, p1, p2, globalState):
        source2_ = D4_DC13(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): True}))
        if source2_.is_DC14:
            d_36___mcc_h0_ = source2_.cf21
            d_37___mcc_h1_ = source2_.cf22
            d_38_cf22_ = d_37___mcc_h1_
            d_39_cf21_ = d_36___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([D5_DC18(D5_DC17(False, _dafny.CodePoint('v'))), D5_DC18(D5_DC18(D5_DC17(False, _dafny.CodePoint('k'))))])) + (_dafny.SeqWithoutIsStrInference([D5_DC18(D5_DC17(True, _dafny.CodePoint('e'))), D5_DC18(D5_DC16(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])))), D5_DC18(D5_DC18(D5_DC16(_dafny.MultiSet([True]))))]))
        elif source2_.is_DC15:
            d_40___mcc_h2_ = source2_.cf23
            d_41___mcc_h3_ = source2_.cf24
            d_42___mcc_h4_ = source2_.cf25
            d_43___mcc_h5_ = source2_.cf26
            d_44_cf26_ = d_43___mcc_h5_
            d_45_cf25_ = d_42___mcc_h4_
            d_46_cf24_ = d_41___mcc_h3_
            d_47_cf23_ = d_40___mcc_h2_
            return _dafny.SeqWithoutIsStrInference([D5_DC18(D5_DC18(D5_DC16(_dafny.MultiSet([True]))))])
        elif True:
            d_48___mcc_h6_ = source2_.cf20
            d_49_cf20_ = d_48___mcc_h6_
            return (_dafny.SeqWithoutIsStrInference([D5_DC18(D5_DC18(D5_DC16(_dafny.MultiSet([False]))))])) + (_dafny.SeqWithoutIsStrInference([D5_DC18(D5_DC18(D5_DC18(D5_DC16(_dafny.MultiSet([False, not(False)]))))) for d_50_i0_ in range(default__.abs(22))]))

    @staticmethod
    def fm44(p0, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: _dafny.Seq
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajwlwtmnx"))])).Elements:
                d_52_v0_: _dafny.Seq = compr_6_
                if (d_52_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajwlwtmnx"))])):
                    coll6_[d_52_v0_] = 264
            return _dafny.Map(coll6_)
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_51_i0_ in range(default__.abs(459))]): 153})) | (iife6_()
        )

    @staticmethod
    def fm45(p0, p1, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tp"))): -214})

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        if (D2_DC7(len(_dafny.SeqWithoutIsStrInference([False, True])), True, len(_dafny.Map({_dafny.CodePoint('r'): True})))).cf10:
            return D10_DC31(not(False), False, len(_dafny.Map({True: True})))
        elif True:
            return D10_DC31(True, True, -559)

    @staticmethod
    def fm47(p0, p1, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({False: False})))]), _dafny.SeqWithoutIsStrInference([884]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({True, False}))]), _dafny.SeqWithoutIsStrInference([(D1_DC3((_dafny.MultiSet([770])).cardinality)).cf4])])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_53_i0_ in range(default__.abs(110))])), 900])]))) | ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(False)]))])])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([535 for d_54_i1_ in range(default__.abs(143))])])))

    @staticmethod
    def fm48(globalState):
        return (_dafny.MultiSet([D2_DC8(D2_DC8(D2_DC7(len(_dafny.SeqWithoutIsStrInference([True])), False, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False, False])]))))), D2_DC8(D2_DC6(_dafny.CodePoint('a')))])) - (_dafny.MultiSet([D2_DC8(D2_DC6(_dafny.CodePoint('o')))]))

    @staticmethod
    def fm49(p0, p1, globalState):
        return D3_DC9((_dafny.Set({_dafny.Set({not(not(not(True)))}), _dafny.Set({True})})) | (_dafny.Set({_dafny.Set({True}), _dafny.Set({True})})))

    @staticmethod
    def fm50(p0, p1, globalState):
        if True:
            def iife7_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in (_dafny.SeqWithoutIsStrInference([631])).Elements:
                    d_56_v0_: int = compr_7_
                    if (d_56_v0_) in (_dafny.SeqWithoutIsStrInference([631])):
                        coll7_[(d_56_v0_) + (-154)] = _dafny.CodePoint('n')
                return _dafny.Map(coll7_)
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_55_i0_ in range(default__.abs(419))])): _dafny.CodePoint('a')})) | (iife7_()
            )
        elif True:
            return _dafny.Map({827: _dafny.CodePoint('q')})

    @staticmethod
    def fm51(globalState):
        source3_ = D5_DC18(D5_DC18(D5_DC18(D5_DC18(D5_DC16(_dafny.MultiSet([True]))))))
        if source3_.is_DC17:
            d_57___mcc_h0_ = source3_.cf28
            d_58___mcc_h1_ = source3_.cf29
            d_59_cf29_ = d_58___mcc_h1_
            d_60_cf28_ = d_57___mcc_h0_
            return (_dafny.Map({True: 120})) | (_dafny.Map({d_60_cf28_: -616}))
        elif source3_.is_DC16:
            d_61___mcc_h2_ = source3_.cf27
            d_62_cf27_ = d_61___mcc_h2_
            return (D28_DC76(_dafny.Map({not(True): len(_dafny.Set({812}))}))).cf134
        elif True:
            d_63___mcc_h3_ = source3_.cf30
            d_64_cf30_ = d_63___mcc_h3_
            return _dafny.Map({False: (0) - (-316)})

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, True]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True, True]), _dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([False, True])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False, False]), _dafny.SeqWithoutIsStrInference([True, False]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([not(not(not(True))), not(True)]), _dafny.SeqWithoutIsStrInference([False, True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])])))

    @staticmethod
    def fm53(p0, p1, globalState):
        return D14_DC43(_dafny.CodePoint('t'), 487, 304, (-702) * (-457))

    @staticmethod
    def fm54(globalState):
        source4_ = D10_DC31(False, False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yiqmqaeh"))))
        if source4_.is_DC31:
            d_65___mcc_h0_ = source4_.cf50
            d_66___mcc_h1_ = source4_.cf51
            d_67___mcc_h2_ = source4_.cf52
            d_68_cf52_ = d_67___mcc_h2_
            d_69_cf51_ = d_66___mcc_h1_
            d_70_cf50_ = d_65___mcc_h0_
            return D4_DC15(d_68_cf52_, _dafny.MultiSet([d_68_cf52_]), _dafny.CodePoint('r'), d_68_cf52_)
        elif True:
            d_71___mcc_h3_ = source4_.cf49
            d_72_cf49_ = d_71___mcc_h3_
            return D4_DC15(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False])), (_dafny.MultiSet([_dafny.CodePoint('f'), _dafny.CodePoint('m')])).cardinality])), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, False]))]), _dafny.CodePoint('f'), 677)

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        source5_ = (D20_DC60((_dafny.MultiSet([True])).cardinality) if False else D20_DC60(589))
        if source5_.is_DC60:
            d_73___mcc_h0_ = source5_.cf102
            d_74_cf102_ = d_73___mcc_h0_
            return D18_DC53(_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('b'): not(False)}), _dafny.Map({_dafny.CodePoint('p'): False})]))
        elif True:
            d_75___mcc_h1_ = source5_.cf101
            d_76_cf101_ = d_75___mcc_h1_
            return D18_DC53(_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('b'): True})]))

    @staticmethod
    def fm56(globalState):
        return D16_DC49(394)

    @staticmethod
    def fm57(globalState):
        return D15_DC45(-864, (D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsgxfjmq")), False, True)).cf16, False)

    @staticmethod
    def fm58(globalState):
        return D19_DC58(((_dafny.MultiSet([len(_dafny.Set({True})), len(_dafny.Set({not(False)})), 814])).cardinality) + (len(_dafny.SeqWithoutIsStrInference([False]))), not((757) <= (805)))

    @staticmethod
    def fm59(p0, p1, p2, p3, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(464, 352):
                d_78_v0_: int = compr_8_
                if ((464) <= (d_78_v0_)) and ((d_78_v0_) < (352)):
                    coll8_[default__.safeModuloInt(d_78_v0_, len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.Map({False: True}))
            return _dafny.Map(coll8_)
        def iife9_():
            coll9_ = _dafny.Map()
            compr_9_: str
            for compr_9_ in (_dafny.Map({_dafny.CodePoint('m'): len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doi"))))}))})).keys.Elements:
                d_79_v1_: str = compr_9_
                if (d_79_v1_) in (_dafny.Map({_dafny.CodePoint('m'): len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doi"))))}))})):
                    coll9_[d_79_v1_] = True
            return _dafny.Map(coll9_)
        return (((D29_DC79((D29_DC79(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({990: True})) for d_77_i0_ in range(default__.abs(9))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True]))]), _dafny.SeqWithoutIsStrInference([len(iife8_()
)])]))).cf141)).cf141) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([952, len(_dafny.SeqWithoutIsStrInference([D18_DC53(_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('m'): False})])), D18_DC53(_dafny.SeqWithoutIsStrInference([iife9_()
, _dafny.Map({_dafny.CodePoint('t'): False})]))]))]) for d_80_i1_ in range(default__.abs(500))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([485, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D6_DC21(True, True, True, 400)).cf39 for d_81_i2_ in range(default__.abs(-386))]))).cardinality]), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({925: False})) for d_82_i3_ in range(default__.abs(311))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.Set({True})), 250})), len(_dafny.Map({345: False}))])]))

    @staticmethod
    def fm60(p0, p1, p2, p3, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: str
            for compr_10_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('n'), _dafny.CodePoint('w'), _dafny.CodePoint('f')])).Elements:
                d_83_v0_: str = compr_10_
                if (d_83_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('n'), _dafny.CodePoint('w'), _dafny.CodePoint('f')])):
                    coll10_[d_83_v0_] = True
            return _dafny.Map(coll10_)
        return _dafny.Map({default__.safeDivisionInt((0) - ((_dafny.MultiSet([not(not(True)), False])).cardinality), 141): iife10_()
        })

    @staticmethod
    def fm61(p0, p1, p2, p3, globalState):
        return D7_DC22((_dafny.Map({-6: -792})) | (_dafny.Map({len(_dafny.Set({False})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "redxim")))})))

    @staticmethod
    def fm62(p0, globalState):
        return _dafny.Map({(_dafny.MultiSet([False, True])) - (_dafny.MultiSet([not(False), not(not(False))])): (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({129: False})), len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, not(True)])).cardinality for d_84_i0_ in range(default__.abs(236))])), len(_dafny.Map({not(False): False}))])])).issubset(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({230: 483})))])]))})

    @staticmethod
    def fm63(globalState):
        source6_ = D18_DC56(_dafny.SeqWithoutIsStrInference([False]))
        if source6_.is_DC54:
            d_85___mcc_h0_ = source6_.cf91
            d_86___mcc_h1_ = source6_.cf92
            d_87___mcc_h2_ = source6_.cf93
            d_88___mcc_h3_ = source6_.cf94
            d_89_cf94_ = d_88___mcc_h3_
            d_90_cf93_ = d_87___mcc_h2_
            d_91_cf92_ = d_86___mcc_h1_
            d_92_cf91_ = d_85___mcc_h0_
            return D5_DC16(_dafny.MultiSet([d_90_cf93_, d_90_cf93_, d_90_cf93_]))
        elif source6_.is_DC55:
            d_93___mcc_h4_ = source6_.cf95
            d_94___mcc_h5_ = source6_.cf96
            d_95_cf96_ = d_94___mcc_h5_
            d_96_cf95_ = d_93___mcc_h4_
            return D5_DC16(_dafny.MultiSet([False]))
        elif source6_.is_DC56:
            d_97___mcc_h6_ = source6_.cf97
            d_98_cf97_ = d_97___mcc_h6_
            return D5_DC16((D5_DC16(_dafny.MultiSet([False]))).cf27)
        elif True:
            d_99___mcc_h7_ = source6_.cf90
            d_100_cf90_ = d_99___mcc_h7_
            return D5_DC16(_dafny.MultiSet([True, not(True)]))

    @staticmethod
    def fm64(globalState):
        source7_ = D5_DC16(_dafny.MultiSet([not(False), False, False, True]))
        if source7_.is_DC17:
            d_101___mcc_h0_ = source7_.cf28
            d_102___mcc_h1_ = source7_.cf29
            d_103_cf29_ = d_102___mcc_h1_
            d_104_cf28_ = d_101___mcc_h0_
            return D10_DC30(_dafny.Set({d_104_cf28_, not(d_104_cf28_), d_104_cf28_}))
        elif source7_.is_DC16:
            d_105___mcc_h2_ = source7_.cf27
            d_106_cf27_ = d_105___mcc_h2_
            return D10_DC30(_dafny.Set({not(True)}))
        elif True:
            d_107___mcc_h3_ = source7_.cf30
            d_108_cf30_ = d_107___mcc_h3_
            return D10_DC30(_dafny.Set({not(True)}))

    @staticmethod
    def fm65(p0, p1, p2, globalState):
        source8_ = D28_DC77((0) - (len(_dafny.Map({not(not(True)): _dafny.Set({False})}))), _dafny.Set({True}), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_109_i0_ in range(default__.abs(782))])), True, _dafny.Map({_dafny.MultiSet([448, len(_dafny.Set({True})), (0) - ((_dafny.MultiSet([-690])).cardinality)]): True}))
        if source8_.is_DC77:
            d_110___mcc_h0_ = source8_.cf135
            d_111___mcc_h1_ = source8_.cf136
            d_112___mcc_h2_ = source8_.cf137
            d_113___mcc_h3_ = source8_.cf138
            d_114___mcc_h4_ = source8_.cf139
            d_115_cf139_ = d_114___mcc_h4_
            d_116_cf138_ = d_113___mcc_h3_
            d_117_cf137_ = d_112___mcc_h2_
            d_118_cf136_ = d_111___mcc_h1_
            d_119_cf135_ = d_110___mcc_h0_
            return D3_DC10(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_120_i1_ in range(default__.abs(-969))]), d_116_cf138_, not(d_116_cf138_))
        elif source8_.is_DC76:
            d_121___mcc_h5_ = source8_.cf134
            d_122_cf134_ = d_121___mcc_h5_
            return D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eicihmi")), not(False), False)
        elif True:
            d_123___mcc_h6_ = source8_.cf140
            d_124_cf140_ = d_123___mcc_h6_
            return D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmdp")), not(False), False)

    @staticmethod
    def fm66(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True)]): True})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): True}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): (D27_DC73(True, -640, _dafny.Set({True}))).cf124}))

    @staticmethod
    def fm67(p0, p1, p2, globalState):
        return D21_DC63(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swsbxe"))), (True) and (False), not((_dafny.Set({862, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality, len(_dafny.Map({True: False})), len(_dafny.SeqWithoutIsStrInference([not(True), not(False), False])), len(_dafny.Set({True}))})).ispropersubset(_dafny.Set({(0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality)}))), (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-712, 727, 993]))]))) == (658))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_125_v0_: bool
        d_125_v0_ = False
        d_126_v1_: _dafny.Map
        d_126_v1_ = _dafny.Map({p1: d_125_v0_})
        d_127_v2_: _dafny.Array
        def lambda0_(d_128_p0_):
            def lambda1_(d_129_i0_):
                return (d_129_i0_) * (d_128_p0_)

            return lambda1_

        init0_ = lambda0_(p0)
        nw0_ = _dafny.Array(None, 21)
        for i0_0_ in range(nw0_.length(0)):
            nw0_[i0_0_] = init0_(i0_0_)
        d_127_v2_ = nw0_
        index0_ = default__.safeIndex(173, (d_127_v2_).length(0))
        (d_127_v2_)[index0_] = p0
        d_130_v3_: _dafny.Seq
        d_130_v3_ = _dafny.SeqWithoutIsStrInference([d_125_v0_])
        d_131_v4_: _dafny.Seq
        d_131_v4_ = _dafny.SeqWithoutIsStrInference([114, p0])
        d_132_v5_: _dafny.MultiSet
        d_132_v5_ = _dafny.MultiSet([d_125_v0_, d_125_v0_])
        d_133_v6_: str
        d_133_v6_ = _dafny.CodePoint('h')
        d_134_v7_: _dafny.Seq
        d_134_v7_ = _dafny.SeqWithoutIsStrInference([d_133_v6_, d_133_v6_])
        d_135_v8_: D11
        d_135_v8_ = D11_DC34((0) - (p0), not(default__.fm0((d_132_v5_).cardinality, globalState)), len(d_134_v7_))
        pat_let_tv0_ = d_133_v6_
        pat_let_tv1_ = d_134_v7_
        pat_let_tv2_ = d_133_v6_
        pat_let_tv3_ = d_134_v7_
        index1_ = default__.safeIndex(173, (d_127_v2_).length(0))
        def lambda2_(source9_):
            if source9_.is_DC33:
                d_136___mcc_h0_ = source9_.cf54
                d_137___mcc_h1_ = source9_.cf55
                d_138_cf55_ = d_137___mcc_h1_
                d_139_cf54_ = d_136___mcc_h0_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwc"))
            elif source9_.is_DC34:
                d_140___mcc_h2_ = source9_.cf56
                d_141___mcc_h3_ = source9_.cf57
                d_142___mcc_h4_ = source9_.cf58
                d_143_cf58_ = d_142___mcc_h4_
                d_144_cf57_ = d_141___mcc_h3_
                d_145_cf56_ = d_140___mcc_h2_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv0_ for d_146_i1_ in range(default__.abs(-397))])
            elif True:
                d_147___mcc_h5_ = source9_.cf53
                d_148_cf53_ = d_147___mcc_h5_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "duyjggi"))) + (pat_let_tv1_)

        def lambda3_(source10_):
            if source10_.is_DC33:
                d_149___mcc_h6_ = source10_.cf54
                d_150___mcc_h7_ = source10_.cf55
                d_151_cf55_ = d_150___mcc_h7_
                d_152_cf54_ = d_149___mcc_h6_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwc"))
            elif source10_.is_DC34:
                d_153___mcc_h8_ = source10_.cf56
                d_154___mcc_h9_ = source10_.cf57
                d_155___mcc_h10_ = source10_.cf58
                d_156_cf58_ = d_155___mcc_h10_
                d_157_cf57_ = d_154___mcc_h9_
                d_158_cf56_ = d_153___mcc_h8_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv2_ for d_159_i1_ in range(default__.abs(-397))])
            elif True:
                d_160___mcc_h11_ = source10_.cf53
                d_161_cf53_ = d_160___mcc_h11_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "duyjggi"))) + (pat_let_tv3_)

        rhs0_ = default__.fm0((_dafny.MultiSet(((d_130_v3_ if d_125_v0_ else d_130_v3_)).set(default__.safeIndex(p0, len((d_130_v3_ if d_125_v0_ else d_130_v3_))), d_125_v0_))).cardinality, globalState)
        rhs1_ = True
        rhs2_ = _dafny.Map({p1: (d_131_v4_) == (d_131_v4_)})
        rhs3_ = (0) - (len((lambda2_(d_135_v8_)).set(default__.safeIndex(p0, len(lambda3_(d_135_v8_))), d_133_v6_)))
        lhs0_ = d_127_v2_
        lhs1_ = default__.safeIndex(173, (d_127_v2_).length(0))
        d_125_v0_ = rhs0_
        d_125_v0_ = rhs1_
        d_126_v1_ = rhs2_
        lhs0_[lhs1_] = rhs3_
        d_162_i2_: int
        d_162_i2_ = 0
        with _dafny.label("0"):
            while ((d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]) == (default__.safeDivisionInt(p0, -473)):
                with _dafny.c_label("0"):
                    if (d_162_i2_) >= (100):
                        raise _dafny.Break("0")
                    d_162_i2_ = (d_162_i2_) + (1)
                    d_163_v9_: _dafny.Array
                    nw1_ = _dafny.Array(_dafny.Array(None, 0), 5)
                    d_163_v9_ = nw1_
                    index2_ = default__.safeIndex(922, (d_163_v9_).length(0))
                    (d_163_v9_)[index2_] = p1
                    index3_ = default__.safeIndex(922, (d_163_v9_).length(0))
                    (d_163_v9_)[index3_] = (p1 if (not(d_125_v0_) if d_125_v0_ else d_125_v0_) else p1)
                    r1 = ((d_134_v7_) + (d_134_v7_)) + (d_134_v7_)
                    (globalState).f7 = (d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]
                    if d_125_v0_:
                        d_164_v10_: _dafny.Map
                        d_164_v10_ = _dafny.Map({d_125_v0_: d_125_v0_})
                        d_165_v11_: T0
                        nw2_ = C12()
                        nw2_.ctor__(d_125_v0_, d_125_v0_, d_125_v0_)
                        d_165_v11_ = nw2_
                        d_166_v12_: _dafny.Map
                        d_166_v12_ = _dafny.Map({len(d_164_v10_): d_165_v11_})
                        d_167_v13_: _dafny.Seq
                        d_167_v13_ = _dafny.SeqWithoutIsStrInference([((d_166_v12_)[(d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]] if ((d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]) in (d_166_v12_) else d_165_v11_), d_165_v11_])
                        d_168_v14_: _dafny.Seq
                        d_168_v14_ = _dafny.SeqWithoutIsStrInference([(d_130_v3_).set(default__.safeIndex(p0, len(d_130_v3_)), (d_165_v11_).f10), d_130_v3_, d_130_v3_, d_130_v3_, _dafny.SeqWithoutIsStrInference([(d_165_v11_).f10])])
                        d_167_v13_ = (((d_167_v13_) + (d_167_v13_)).set(default__.safeIndex(len(d_168_v14_), len((d_167_v13_) + (d_167_v13_))), d_165_v11_)) + (d_167_v13_)
                        d_169_v15_: _dafny.Set
                        d_169_v15_ = _dafny.Set({p0})
                        d_170_v16_: _dafny.Set
                        d_170_v16_ = _dafny.Set({(d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))], p0, (0) - (len(d_169_v15_)), p0})
                        d_125_v0_ = ((d_169_v15_).intersection(d_169_v15_)).issubset(_dafny.Set({-157}))
                        d_127_v2_ = d_127_v2_
                        d_131_v4_ = default__.fm3(142, (p0 if (d_165_v11_).f10 else 973), globalState)
                        r0 = d_134_v7_
                    elif True:
                        (globalState).f0 = default__.safeDivisionInt(p0, default__.safeDivisionInt(p0, -689))
                        d_171_v17_: _dafny.Map
                        d_171_v17_ = _dafny.Map({(d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]: p0})
                        d_171_v17_ = (d_171_v17_).set((0) - (p0), default__.safeModuloInt(p0, len(_dafny.Set({p0, len(d_134_v7_)}))))
                        d_125_v0_ = default__.fm0(default__.fm1(globalState), globalState)
                        d_172_v18_: _dafny.Array
                        nw3_ = _dafny.Array(D1.default()(), 2)
                        d_172_v18_ = nw3_
                        d_173_v19_: D8
                        d_173_v19_ = D8_DC27(d_125_v0_, _dafny.SeqWithoutIsStrInference([d_172_v18_, d_172_v18_]), d_131_v4_, (d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))])
                        d_174_v20_: D21
                        d_174_v20_ = D21_DC63((d_173_v19_).cf46, d_125_v0_, d_125_v0_, d_125_v0_)
                        d_125_v0_ = not((((0) - (p0)) * (((d_132_v5_)[d_125_v0_] if (d_125_v0_) in (d_132_v5_) else 938))) > ((0) - ((d_174_v20_).cf106)))
                        (globalState).f1 = default__.safeDivisionInt(355, p0)
                    pass
            pass
        if not (not (d_125_v0_) or (not(d_125_v0_))) or (d_125_v0_):
            d_175_v21_: _dafny.Seq
            d_175_v21_ = _dafny.SeqWithoutIsStrInference([d_134_v7_])
            d_176_v22_: _dafny.Map
            d_176_v22_ = _dafny.Map({d_175_v21_: not(d_125_v0_)})
            d_176_v22_ = (d_176_v22_).set(_dafny.SeqWithoutIsStrInference([d_134_v7_]), d_125_v0_)
            d_177_v23_: _dafny.MultiSet
            d_177_v23_ = _dafny.MultiSet([(0) - ((d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))])])
            d_178_v24_: _dafny.Array
            nw4_ = _dafny.Array(None, 2)
            nw4_[int(0)] = d_177_v23_
            nw4_[int(1)] = d_177_v23_
            d_178_v24_ = nw4_
            d_178_v24_ = d_178_v24_
            (globalState).f1 = p0
            d_125_v0_ = (d_133_v6_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nblvw")))
            d_179_v25_: C9
            nw5_ = C9()
            nw5_.ctor__(d_125_v0_, default__.fm0(len(default__.fm33(d_125_v0_, d_125_v0_, d_134_v7_, d_125_v0_, globalState)), globalState))
            d_179_v25_ = nw5_
        elif True:
            (globalState).f0 = p0
            d_180_v26_: _dafny.Map
            d_180_v26_ = _dafny.Map({p0: 801})
            if (p0) > ((((d_132_v5_)[d_125_v0_] if (d_125_v0_) in (d_132_v5_) else len(d_180_v26_))) - ((d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))])):
                d_181_v27_: _dafny.Seq
                d_181_v27_ = _dafny.SeqWithoutIsStrInference([d_127_v2_, d_127_v2_, d_127_v2_, d_127_v2_, d_127_v2_])
                d_182_v28_: T1
                nw6_ = C2()
                nw6_.ctor__(d_125_v0_, (d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))])
                d_182_v28_ = nw6_
                d_183_v29_: _dafny.MultiSet
                d_183_v29_ = _dafny.MultiSet([d_182_v28_])
                d_184_v30_: _dafny.Map
                d_184_v30_ = _dafny.Map({(d_181_v27_)[default__.safeIndex((0) - ((d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]), len(d_181_v27_))]: (d_183_v29_) == (d_183_v29_)})
                d_184_v30_ = (d_184_v30_) | (d_184_v30_)
                d_185_v31_: T0
                nw7_ = C3()
                nw7_.ctor__(d_125_v0_, d_125_v0_)
                d_185_v31_ = nw7_
                d_185_v31_ = d_185_v31_
                index4_ = default__.safeIndex(173, (d_127_v2_).length(0))
                (d_127_v2_)[index4_] = (0) - (p0)
                nw8_ = _dafny.Array(int(0), 2)
                d_127_v2_ = nw8_
                d_125_v0_ = (d_185_v31_).f10
            elif True:
                d_186_v32_: _dafny.Array
                nw9_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
                d_186_v32_ = nw9_
                index5_ = default__.safeIndex(118, (d_186_v32_).length(0))
                (d_186_v32_)[index5_] = d_134_v7_
                index6_ = default__.safeIndex(118, (d_186_v32_).length(0))
                (d_186_v32_)[index6_] = ((d_134_v7_).set(default__.safeIndex(p0, len(d_134_v7_)), d_133_v6_)) + (d_134_v7_)
                d_187_v33_: _dafny.MultiSet
                d_187_v33_ = _dafny.MultiSet([len((d_186_v32_)[default__.safeIndex(118, (d_186_v32_).length(0))]), (d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]])
                index7_ = default__.safeIndex(988, (p1).length(0))
                (p1)[index7_] = (d_187_v33_).isdisjoint(d_187_v33_)
                index8_ = default__.safeIndex(988, (p1).length(0))
                index9_ = default__.safeIndex(173, (d_127_v2_).length(0))
                rhs4_ = p0
                rhs5_ = (d_134_v7_) < (_dafny.SeqWithoutIsStrInference([d_133_v6_ for d_188_i3_ in range(default__.abs(326))]))
                rhs6_ = p0
                lhs2_ = globalState
                lhs3_ = p1
                lhs4_ = default__.safeIndex(988, (p1).length(0))
                lhs5_ = d_127_v2_
                lhs6_ = default__.safeIndex(173, (d_127_v2_).length(0))
                lhs2_.f7 = rhs4_
                lhs3_[lhs4_] = rhs5_
                lhs5_[lhs6_] = rhs6_
                d_189_v34_: _dafny.Array
                nw10_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_189_v34_ = nw10_
                index10_ = default__.safeIndex(529, (d_189_v34_).length(0))
                (d_189_v34_)[index10_] = d_127_v2_
                index11_ = default__.safeIndex(529, (d_189_v34_).length(0))
                nw11_ = _dafny.Array(int(0), 26)
                (d_189_v34_)[index11_] = nw11_
                d_190_v35_: _dafny.MultiSet
                d_190_v35_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(988, (p1).length(0))]])) + (d_130_v3_)])
                d_191_v36_: _dafny.Map
                d_191_v36_ = _dafny.Map({p0: d_125_v0_})
                d_190_v35_ = default__.fm19((p1)[default__.safeIndex(988, (p1).length(0))], ((0) - (len(d_191_v36_))) > ((d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]), (p1)[default__.safeIndex(988, (p1).length(0))], (d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))], globalState)
                d_192_v37_: _dafny.Map
                d_192_v37_ = _dafny.Map({(p1)[default__.safeIndex(988, (p1).length(0))]: (p1)[default__.safeIndex(988, (p1).length(0))]})
                d_193_v38_: _dafny.Array
                nw12_ = _dafny.Array(None, 21)
                nw12_[int(0)] = (p1)[default__.safeIndex(988, (p1).length(0))]
                nw12_[int(1)] = False
                nw12_[int(2)] = d_125_v0_
                nw12_[int(3)] = default__.fm0((d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))], globalState)
                nw12_[int(4)] = (p1)[default__.safeIndex(988, (p1).length(0))]
                nw12_[int(5)] = (p1)[default__.safeIndex(988, (p1).length(0))]
                nw12_[int(6)] = (p1)[default__.safeIndex(988, (p1).length(0))]
                nw12_[int(7)] = ((d_192_v37_)[(p1)[default__.safeIndex(988, (p1).length(0))]] if ((p1)[default__.safeIndex(988, (p1).length(0))]) in (d_192_v37_) else (p1)[default__.safeIndex(988, (p1).length(0))])
                nw12_[int(8)] = d_125_v0_
                nw12_[int(9)] = d_125_v0_
                nw12_[int(10)] = (p1)[default__.safeIndex(988, (p1).length(0))]
                nw12_[int(11)] = d_125_v0_
                nw12_[int(12)] = (p1)[default__.safeIndex(988, (p1).length(0))]
                nw12_[int(13)] = (p1)[default__.safeIndex(988, (p1).length(0))]
                nw12_[int(14)] = (p1)[default__.safeIndex(988, (p1).length(0))]
                nw12_[int(15)] = (p1)[default__.safeIndex(988, (p1).length(0))]
                nw12_[int(16)] = False
                nw12_[int(17)] = d_125_v0_
                nw12_[int(18)] = d_125_v0_
                nw12_[int(19)] = d_125_v0_
                nw12_[int(20)] = not((p1)[default__.safeIndex(988, (p1).length(0))])
                d_193_v38_ = nw12_
                d_194_v39_: _dafny.Array
                nw13_ = _dafny.Array(None, 2)
                nw13_[int(0)] = d_193_v38_
                nw13_[int(1)] = d_193_v38_
                d_194_v39_ = nw13_
                index12_ = default__.safeIndex(784, (d_194_v39_).length(0))
                (d_194_v39_)[index12_] = d_193_v38_
                index13_ = default__.safeIndex(784, (d_194_v39_).length(0))
                (d_194_v39_)[index13_] = d_193_v38_
            d_195_v40_: C8
            nw14_ = C8()
            nw14_.ctor__(default__.fm4(d_125_v0_, True, d_125_v0_, globalState))
            d_195_v40_ = nw14_
            d_196_v41_: C8
            d_196_v41_ = d_195_v40_
            d_197_v42_: _dafny.Array
            nw15_ = _dafny.Array(None, 2)
            nw15_[int(0)] = ((d_196_v41_) if True else d_195_v40_)
            nw15_[int(1)] = d_195_v40_
            d_197_v42_ = nw15_
            d_197_v42_ = d_197_v42_
            d_125_v0_ = d_125_v0_
            d_198_v43_: _dafny.MultiSet
            d_198_v43_ = _dafny.MultiSet([len(d_131_v4_)])
            d_199_v44_: _dafny.Map
            d_199_v44_ = _dafny.Map({not(True): d_198_v43_})
            (globalState).f1 = ((0) - ((((d_199_v44_)[d_125_v0_] if (d_125_v0_) in (d_199_v44_) else _dafny.MultiSet([len(_dafny.Map({p0: True})), p0, p0, (_dafny.MultiSet(d_131_v4_)).cardinality, p0]))).cardinality)) * (p0)
        hi0_ = len(d_130_v3_)
        for d_200_i4_ in range((d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))], hi0_):
            d_201_v45_: _dafny.Set
            d_201_v45_ = _dafny.Set({_dafny.CodePoint('a'), d_133_v6_, d_133_v6_})
            index14_ = default__.safeIndex(173, (d_127_v2_).length(0))
            def iife11_():
                coll11_ = _dafny.Set()
                compr_11_: str
                for compr_11_ in ((d_134_v7_).set(default__.safeIndex(p0, len(d_134_v7_)), d_133_v6_)).Elements:
                    d_202_v46_: str = compr_11_
                    if (d_202_v46_) in ((d_134_v7_).set(default__.safeIndex(p0, len(d_134_v7_)), d_133_v6_)):
                        coll11_ = coll11_.union(_dafny.Set([d_202_v46_]))
                return _dafny.Set(coll11_)
            (d_127_v2_)[index14_] = ((d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))] if (d_201_v45_).issubset(iife11_()
            ) else d_200_i4_)
            d_203_v47_: _dafny.Array
            nw16_ = _dafny.Array(None, 6)
            nw16_[int(0)] = d_134_v7_
            nw16_[int(1)] = (d_134_v7_) + (d_134_v7_)
            nw16_[int(2)] = _dafny.SeqWithoutIsStrInference([d_133_v6_ for d_204_i5_ in range(default__.abs(-211))])
            nw16_[int(3)] = (d_134_v7_) + (d_134_v7_)
            nw16_[int(4)] = (d_134_v7_) + (d_134_v7_)
            nw16_[int(5)] = d_134_v7_
            d_203_v47_ = nw16_
            index15_ = default__.safeIndex(927, (d_203_v47_).length(0))
            (d_203_v47_)[index15_] = d_134_v7_
            index16_ = default__.safeIndex(927, (d_203_v47_).length(0))
            (d_203_v47_)[index16_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwrlipv"))
            if d_125_v0_:
                d_134_v7_ = ((d_203_v47_)[default__.safeIndex(927, (d_203_v47_).length(0))]) + (_dafny.SeqWithoutIsStrInference([d_133_v6_ for d_205_i6_ in range(default__.abs(657))]))
                index17_ = default__.safeIndex(173, (d_127_v2_).length(0))
                (d_127_v2_)[index17_] = len(_dafny.Set({not (d_125_v0_) or (d_125_v0_)}))
                d_125_v0_ = d_125_v0_
                d_206_v48_: _dafny.Map
                d_206_v48_ = _dafny.Map({d_125_v0_: d_125_v0_})
                d_207_v49_: _dafny.Map
                d_207_v49_ = _dafny.Map({d_125_v0_: d_206_v48_})
                d_207_v49_ = (d_207_v49_).set((d_130_v3_) != (_dafny.SeqWithoutIsStrInference([False])), d_206_v48_)
                d_208_v50_: T0
                nw17_ = C9()
                nw17_.ctor__(d_125_v0_, d_125_v0_)
                d_208_v50_ = nw17_
                d_209_v51_: _dafny.Map
                d_209_v51_ = _dafny.Map({d_133_v6_: d_208_v50_})
                d_209_v51_ = ((d_209_v51_).set(d_133_v6_, d_208_v50_)) | ((d_209_v51_).set(d_133_v6_, d_208_v50_))
            elif True:
                d_125_v0_ = default__.fm0(825, globalState)
                (globalState).f1 = (d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]
                d_210_v52_: C10
                nw18_ = C10()
                nw18_.ctor__(d_200_i4_, d_125_v0_)
                d_210_v52_ = nw18_
                d_125_v0_ = (d_210_v52_).f16
                index18_ = default__.safeIndex(979, (p1).length(0))
                (p1)[index18_] = d_125_v0_
                index19_ = default__.safeIndex(979, (p1).length(0))
                (p1)[index19_] = default__.fm0(185, globalState)
            index20_ = default__.safeIndex(173, (d_127_v2_).length(0))
            (d_127_v2_)[index20_] = (d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]
        index21_ = default__.safeIndex(173, (d_127_v2_).length(0))
        rhs7_ = -290
        rhs8_ = (d_127_v2_)[default__.safeIndex(173, (d_127_v2_).length(0))]
        rhs9_ = d_125_v0_
        lhs7_ = d_127_v2_
        lhs8_ = default__.safeIndex(173, (d_127_v2_).length(0))
        lhs9_ = globalState
        lhs7_[lhs8_] = rhs7_
        lhs9_.f7 = rhs8_
        d_125_v0_ = rhs9_
        d_211_v53_: C2
        nw19_ = C2()
        nw19_.ctor__(d_125_v0_, 936)
        d_211_v53_ = nw19_
        d_212_v54_: _dafny.MultiSet
        d_212_v54_ = _dafny.MultiSet([d_211_v53_])
        (globalState).f1 = (-876) + ((d_212_v54_).cardinality)
        r0 = d_134_v7_
        r1 = d_134_v7_
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_213_v0_: bool
        d_213_v0_ = True
        d_214_v1_: _dafny.Set
        d_214_v1_ = _dafny.Set({d_213_v0_})
        d_215_v2_: str
        d_215_v2_ = _dafny.CodePoint('c')
        d_216_v3_: _dafny.MultiSet
        d_216_v3_ = _dafny.MultiSet([not(False)])
        d_217_v4_: _dafny.Seq
        d_217_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxfbexr"))
        d_218_globalState_: GlobalState
        nw20_ = GlobalState()
        nw20_.ctor__(546, 521, True, d_214_v1_, _dafny.Set({d_215_v2_}), d_216_v3_, (d_217_v4_) + (d_217_v4_), 359)
        d_218_globalState_ = nw20_
        d_213_v0_ = d_213_v0_
        d_219_v5_: _dafny.Array
        def lambda4_(d_220_v0_):
            def lambda5_(d_221_i1_):
                return d_220_v0_

            return lambda5_

        init1_ = lambda4_(d_213_v0_)
        nw21_ = _dafny.Array(None, 25)
        for i0_1_ in range(nw21_.length(0)):
            nw21_[i0_1_] = init1_(i0_1_)
        d_219_v5_ = nw21_
        d_222_v6_: _dafny.Seq
        d_222_v6_ = _dafny.SeqWithoutIsStrInference([d_219_v5_])
        d_223_i0_: int
        d_223_i0_ = 0
        with _dafny.label("1"):
            while ((_dafny.SeqWithoutIsStrInference([d_219_v5_])) + (d_222_v6_)) <= ((d_222_v6_) + (_dafny.SeqWithoutIsStrInference([d_219_v5_, d_219_v5_]))):
                with _dafny.c_label("1"):
                    if (d_223_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_223_i0_ = (d_223_i0_) + (1)
                    d_215_v2_ = d_215_v2_
                    d_224_v7_: int
                    d_224_v7_ = 52
                    d_225_v8_: _dafny.Seq
                    d_226_v9_: _dafny.Seq
                    out0_: _dafny.Seq
                    out1_: _dafny.Seq
                    out0_, out1_ = default__.m0((0) - (d_224_v7_), d_219_v5_, d_218_globalState_)
                    d_225_v8_ = out0_
                    d_226_v9_ = out1_
                    d_227_v10_: _dafny.Seq
                    d_227_v10_ = _dafny.SeqWithoutIsStrInference([d_213_v0_, d_213_v0_])
                    d_228_v11_: _dafny.Map
                    d_228_v11_ = _dafny.Map({(d_227_v10_) <= (d_227_v10_): d_224_v7_})
                    d_228_v11_ = (d_228_v11_).set(d_213_v0_, len(d_225_v8_))
                    d_229_v12_: _dafny.Map
                    d_229_v12_ = _dafny.Map({d_215_v2_: d_213_v0_})
                    d_229_v12_ = (d_229_v12_).set(d_215_v2_, d_213_v0_)
                    pass
            pass
        d_230_v13_: _dafny.Map
        d_230_v13_ = _dafny.Map({d_213_v0_: d_213_v0_})
        d_231_v14_: _dafny.Map
        d_231_v14_ = _dafny.Map({len(d_230_v13_): (d_213_v0_) or (d_213_v0_)})
        index22_ = default__.safeIndex(318, (d_219_v5_).length(0))
        (d_219_v5_)[index22_] = d_213_v0_
        index23_ = default__.safeIndex(318, (d_219_v5_).length(0))
        rhs10_ = d_231_v14_
        rhs11_ = (d_217_v4_) <= (d_217_v4_)
        lhs10_ = d_219_v5_
        lhs11_ = default__.safeIndex(318, (d_219_v5_).length(0))
        d_231_v14_ = rhs10_
        lhs10_[lhs11_] = rhs11_
        d_232_v15_: _dafny.Seq
        d_232_v15_ = _dafny.SeqWithoutIsStrInference([(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsrrqfxka")))) - (435)])
        (d_218_globalState_).f1 = (0) - (len(d_232_v15_))
        index24_ = default__.safeIndex(318, (d_219_v5_).length(0))
        (d_219_v5_)[index24_] = d_213_v0_
        d_233_v16_: _dafny.Map
        d_233_v16_ = _dafny.Map({len(d_217_v4_): 901})
        d_234_v17_: int
        d_234_v17_ = 470
        d_235_v18_: D0
        d_235_v18_ = D0_DC2((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], (d_222_v6_)[default__.safeIndex(((d_233_v16_)[803] if (803) in (d_233_v16_) else d_234_v17_), len(d_222_v6_))])
        source11_ = d_235_v18_
        if source11_.is_DC0:
            d_236___mcc_h0_ = source11_.cf0
            d_237_cf0_ = d_236___mcc_h0_
            d_238_v19_: _dafny.Seq
            d_238_v19_ = _dafny.SeqWithoutIsStrInference([d_213_v0_, d_213_v0_])
            d_239_v20_: _dafny.Array
            def lambda6_(d_240_v17_):
                def lambda7_(d_241_i2_):
                    return (d_241_i2_) + (d_240_v17_)

                return lambda7_

            init2_ = lambda6_(d_234_v17_)
            nw22_ = _dafny.Array(None, 23)
            for i0_2_ in range(nw22_.length(0)):
                nw22_[i0_2_] = init2_(i0_2_)
            d_239_v20_ = nw22_
            d_242_v21_: D1
            d_242_v21_ = D1_DC4(len(d_238_v19_), d_239_v20_)
            index25_ = default__.safeIndex(318, (d_219_v5_).length(0))
            rhs12_ = d_234_v17_
            rhs13_ = not(((0) - (d_234_v17_)) > ((0) - (d_234_v17_)))
            rhs14_ = (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]
            rhs15_ = (d_234_v17_) != ((d_242_v21_).cf5)
            rhs16_ = (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]
            lhs12_ = d_218_globalState_
            lhs13_ = d_219_v5_
            lhs14_ = default__.safeIndex(318, (d_219_v5_).length(0))
            lhs12_.f1 = rhs12_
            lhs13_[lhs14_] = rhs13_
            d_213_v0_ = rhs14_
            d_213_v0_ = rhs15_
            d_213_v0_ = rhs16_
            index26_ = default__.safeIndex(318, (d_219_v5_).length(0))
            (d_219_v5_)[index26_] = (_dafny.SeqWithoutIsStrInference([d_213_v0_, not(False), (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]])) < (((_dafny.SeqWithoutIsStrInference([default__.fm0(d_234_v17_, d_218_globalState_)])).set(default__.safeIndex(d_234_v17_, len(_dafny.SeqWithoutIsStrInference([default__.fm0(d_234_v17_, d_218_globalState_)]))), (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))])) + (_dafny.SeqWithoutIsStrInference([True])))
            rhs17_ = default__.fm1(d_218_globalState_)
            rhs18_ = (d_235_v18_).cf2
            d_234_v17_ = rhs17_
            d_213_v0_ = rhs18_
            rhs19_ = default__.fm1(d_218_globalState_)
            rhs20_ = d_213_v0_
            rhs21_ = d_217_v4_
            d_234_v17_ = rhs19_
            d_213_v0_ = rhs20_
            d_217_v4_ = rhs21_
        elif source11_.is_DC1:
            d_243___mcc_h1_ = source11_.cf1
            d_244_cf1_ = d_243___mcc_h1_
            d_245_v22_: _dafny.Map
            d_245_v22_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncsyqp")): False})
            index27_ = default__.safeIndex(318, (d_219_v5_).length(0))
            (d_219_v5_)[index27_] = ((d_245_v22_)[d_244_cf1_] if (d_244_cf1_) in (d_245_v22_) else not((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]))
            d_246_v23_: _dafny.Seq
            d_246_v23_ = _dafny.SeqWithoutIsStrInference([True])
            d_247_v24_: D1
            d_247_v24_ = D1_DC3(len(_dafny.Set({d_235_v18_})))
            (d_218_globalState_).f1 = (len(d_246_v23_)) - ((d_247_v24_).cf4)
            d_248_v25_: _dafny.Map
            d_248_v25_ = _dafny.Map({d_219_v5_: d_234_v17_})
            (d_218_globalState_).f7 = ((d_248_v25_)[d_219_v5_] if (d_219_v5_) in (d_248_v25_) else d_234_v17_)
            d_249_v26_: _dafny.Map
            d_249_v26_ = _dafny.Map({(d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]: (d_216_v3_).cardinality})
            d_250_v27_: _dafny.Map
            d_250_v27_ = _dafny.Map({d_246_v23_: d_249_v26_})
            d_251_v28_: _dafny.Seq
            d_251_v28_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({d_246_v23_: d_249_v26_})).set(d_246_v23_, d_249_v26_)])
            d_252_v29_: _dafny.Array
            nw23_ = _dafny.Array(None, 7)
            nw23_[int(0)] = (d_250_v27_ if False else _dafny.Map({d_246_v23_: d_249_v26_}))
            nw23_[int(1)] = (d_250_v27_ if d_213_v0_ else d_250_v27_)
            nw23_[int(2)] = d_250_v27_
            nw23_[int(3)] = default__.fm2(d_234_v17_, _dafny.CodePoint('o'), d_234_v17_, d_234_v17_, d_218_globalState_)
            nw23_[int(4)] = _dafny.Map({d_246_v23_: _dafny.Map({default__.fm0(d_234_v17_, d_218_globalState_): d_234_v17_})})
            nw23_[int(5)] = (d_251_v28_)[default__.safeIndex(d_234_v17_, len(d_251_v28_))]
            nw23_[int(6)] = _dafny.Map({d_246_v23_: d_249_v26_})
            d_252_v29_ = nw23_
            index28_ = default__.safeIndex(568, (d_252_v29_).length(0))
            (d_252_v29_)[index28_] = d_250_v27_
            index29_ = default__.safeIndex(568, (d_252_v29_).length(0))
            (d_252_v29_)[index29_] = d_250_v27_
        elif True:
            d_253___mcc_h2_ = source11_.cf2
            d_254___mcc_h3_ = source11_.cf3
            d_255_cf3_ = d_254___mcc_h3_
            d_256_cf2_ = d_253___mcc_h2_
            d_257_v30_: _dafny.Seq
            d_258_v31_: _dafny.Seq
            out2_: _dafny.Seq
            out3_: _dafny.Seq
            out2_, out3_ = default__.m0(d_234_v17_, d_219_v5_, d_218_globalState_)
            d_257_v30_ = out2_
            d_258_v31_ = out3_
            d_234_v17_ = 544
            if True:
                d_259_v32_: D0
                d_259_v32_ = D0_DC0(d_219_v5_)
                index30_ = default__.safeIndex(318, (d_219_v5_).length(0))
                rhs22_ = ((d_235_v18_).cf2 if default__.fm0(d_234_v17_, d_218_globalState_) else False)
                rhs23_ = default__.fm1(d_218_globalState_)
                rhs24_ = d_256_cf2_
                rhs25_ = d_259_v32_
                lhs15_ = d_219_v5_
                lhs16_ = default__.safeIndex(318, (d_219_v5_).length(0))
                lhs17_ = d_218_globalState_
                lhs15_[lhs16_] = rhs22_
                lhs17_.f1 = rhs23_
                d_256_cf2_ = rhs24_
                d_259_v32_ = rhs25_
                d_260_v33_: _dafny.Seq
                d_261_v34_: _dafny.Seq
                out4_: _dafny.Seq
                out5_: _dafny.Seq
                out4_, out5_ = default__.m0(default__.fm1(d_218_globalState_), d_219_v5_, d_218_globalState_)
                d_260_v33_ = out4_
                d_261_v34_ = out5_
                d_262_v35_: _dafny.Seq
                d_263_v36_: _dafny.Seq
                out6_: _dafny.Seq
                out7_: _dafny.Seq
                out6_, out7_ = default__.m0((len(d_258_v31_)) - (d_234_v17_), d_255_cf3_, d_218_globalState_)
                d_262_v35_ = out6_
                d_263_v36_ = out7_
                index31_ = default__.safeIndex(318, (d_219_v5_).length(0))
                (d_219_v5_)[index31_] = (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]
                (d_218_globalState_).f0 = (len(d_262_v35_)) - (d_234_v17_)
            elif True:
                d_213_v0_ = (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]
                d_264_v37_: _dafny.Seq
                d_264_v37_ = _dafny.SeqWithoutIsStrInference([False])
                d_265_v38_: _dafny.Seq
                d_266_v39_: _dafny.Seq
                out8_: _dafny.Seq
                out9_: _dafny.Seq
                out8_, out9_ = default__.m0(len(d_264_v37_), d_219_v5_, d_218_globalState_)
                d_265_v38_ = out8_
                d_266_v39_ = out9_
                d_267_v40_: D0
                d_267_v40_ = D0_DC0(d_255_cf3_)
                d_267_v40_ = D0_DC0(d_255_cf3_)
                d_268_v41_: _dafny.Seq
                d_269_v42_: _dafny.Seq
                out10_: _dafny.Seq
                out11_: _dafny.Seq
                out10_, out11_ = default__.m0(d_234_v17_, d_219_v5_, d_218_globalState_)
                d_268_v41_ = out10_
                d_269_v42_ = out11_
                d_270_v43_: _dafny.Array
                nw24_ = _dafny.Array(None, 10)
                nw24_[int(0)] = len(_dafny.SeqWithoutIsStrInference([d_215_v2_ for d_271_i3_ in range(default__.abs(917))]))
                nw24_[int(1)] = d_234_v17_
                nw24_[int(2)] = d_234_v17_
                nw24_[int(3)] = d_234_v17_
                nw24_[int(4)] = d_234_v17_
                nw24_[int(5)] = d_234_v17_
                nw24_[int(6)] = len(d_222_v6_)
                nw24_[int(7)] = (_dafny.MultiSet([d_256_cf2_])).cardinality
                nw24_[int(8)] = len(d_233_v16_)
                nw24_[int(9)] = d_234_v17_
                d_270_v43_ = nw24_
                (d_218_globalState_).f1 = (D1_DC4(620, d_270_v43_)).cf5
            d_272_v44_: D0
            d_272_v44_ = D0_DC1((d_257_v30_) + (d_258_v31_))
            source12_ = d_272_v44_
            if source12_.is_DC0:
                d_273___mcc_h4_ = source12_.cf0
                d_274_cf0_ = d_273___mcc_h4_
                d_275_v45_: _dafny.Seq
                d_276_v46_: _dafny.Seq
                out12_: _dafny.Seq
                out13_: _dafny.Seq
                out12_, out13_ = default__.m0(default__.safeModuloInt((0) - (d_234_v17_), d_234_v17_), d_255_cf3_, d_218_globalState_)
                d_275_v45_ = out12_
                d_276_v46_ = out13_
                d_256_cf2_ = not(default__.fm0(d_234_v17_, d_218_globalState_))
                d_277_v47_: _dafny.Array
                nw25_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_277_v47_ = nw25_
                index32_ = default__.safeIndex(465, (d_277_v47_).length(0))
                (d_277_v47_)[index32_] = d_219_v5_
                d_278_v48_: _dafny.MultiSet
                d_278_v48_ = _dafny.MultiSet([d_234_v17_, d_234_v17_, d_234_v17_])
                d_279_v49_: _dafny.Set
                d_279_v49_ = _dafny.Set({d_275_v45_, _dafny.SeqWithoutIsStrInference([d_215_v2_ for d_280_i4_ in range(default__.abs(476))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ig"))})
                index33_ = default__.safeIndex(318, (d_219_v5_).length(0))
                index34_ = default__.safeIndex(465, (d_277_v47_).length(0))
                rhs26_ = ((d_275_v45_) < (d_217_v4_) if d_256_cf2_ else (d_278_v48_).isdisjoint(d_278_v48_))
                rhs27_ = d_217_v4_
                rhs28_ = ((d_258_v31_) not in (d_279_v49_)) or (((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]) == (d_256_cf2_))
                rhs29_ = d_234_v17_
                rhs30_ = d_255_cf3_
                lhs18_ = d_219_v5_
                lhs19_ = default__.safeIndex(318, (d_219_v5_).length(0))
                lhs20_ = d_218_globalState_
                lhs21_ = d_277_v47_
                lhs22_ = default__.safeIndex(465, (d_277_v47_).length(0))
                d_256_cf2_ = rhs26_
                d_258_v31_ = rhs27_
                lhs18_[lhs19_] = rhs28_
                lhs20_.f7 = rhs29_
                lhs21_[lhs22_] = rhs30_
                d_281_v50_: _dafny.Seq
                d_282_v51_: _dafny.Seq
                out14_: _dafny.Seq
                out15_: _dafny.Seq
                out14_, out15_ = default__.m0(d_234_v17_, (d_222_v6_)[default__.safeIndex(d_234_v17_, len(d_222_v6_))], d_218_globalState_)
                d_281_v50_ = out14_
                d_282_v51_ = out15_
            elif source12_.is_DC1:
                d_283___mcc_h5_ = source12_.cf1
                d_284_cf1_ = d_283___mcc_h5_
                d_285_v52_: _dafny.Array
                nw26_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_285_v52_ = nw26_
                index35_ = default__.safeIndex(845, (d_285_v52_).length(0))
                (d_285_v52_)[index35_] = d_255_cf3_
                d_286_v53_: _dafny.Map
                d_286_v53_ = _dafny.Map({False: d_217_v4_})
                d_287_v54_: _dafny.Seq
                d_287_v54_ = _dafny.SeqWithoutIsStrInference([d_213_v0_])
                d_288_v55_: _dafny.MultiSet
                d_288_v55_ = _dafny.MultiSet([752, d_234_v17_, -36, d_234_v17_, 2])
                d_289_v56_: D0
                d_289_v56_ = D0_DC0(d_255_cf3_)
                d_290_v57_: _dafny.MultiSet
                d_290_v57_ = _dafny.MultiSet([D0_DC0(d_255_cf3_), d_289_v56_, d_289_v56_, D0_DC0(d_219_v5_)])
                d_291_v58_: _dafny.Map
                d_291_v58_ = _dafny.Map({d_234_v17_: _dafny.SeqWithoutIsStrInference([d_234_v17_ for d_292_i5_ in range(default__.abs(549))])})
                index36_ = default__.safeIndex(845, (d_285_v52_).length(0))
                rhs31_ = ((d_286_v53_)[not(((d_230_v13_)[not(d_256_cf2_)] if (not(d_256_cf2_)) in (d_230_v13_) else (d_287_v54_)[default__.safeIndex(0, len(d_287_v54_))]))] if (not(((d_230_v13_)[not(d_256_cf2_)] if (not(d_256_cf2_)) in (d_230_v13_) else (d_287_v54_)[default__.safeIndex(0, len(d_287_v54_))]))) in (d_286_v53_) else d_284_cf1_)
                rhs32_ = not((d_288_v55_).ispropersubset(d_288_v55_))
                rhs33_ = ((default__.fm3(d_234_v17_, d_234_v17_, d_218_globalState_)) + (((d_291_v58_)[((d_216_v3_)[d_213_v0_] if (d_213_v0_) in (d_216_v3_) else d_234_v17_)] if (((d_216_v3_)[d_213_v0_] if (d_213_v0_) in (d_216_v3_) else d_234_v17_)) in (d_291_v58_) else _dafny.SeqWithoutIsStrInference([d_234_v17_ for d_293_i6_ in range(default__.abs(75))]))) if (_dafny.MultiSet([d_289_v56_])).issubset(d_290_v57_) else d_232_v15_)
                rhs34_ = d_255_cf3_
                lhs23_ = d_285_v52_
                lhs24_ = default__.safeIndex(845, (d_285_v52_).length(0))
                d_257_v30_ = rhs31_
                d_213_v0_ = rhs32_
                d_232_v15_ = rhs33_
                lhs23_[lhs24_] = rhs34_
                d_294_v59_: _dafny.Seq
                d_295_v60_: _dafny.Seq
                out16_: _dafny.Seq
                out17_: _dafny.Seq
                out16_, out17_ = default__.m0(340, d_219_v5_, d_218_globalState_)
                d_294_v59_ = out16_
                d_295_v60_ = out17_
                d_296_v61_: _dafny.Array
                nw27_ = _dafny.Array(None, 20)
                nw27_[int(0)] = d_215_v2_
                nw27_[int(1)] = d_215_v2_
                nw27_[int(2)] = d_215_v2_
                nw27_[int(3)] = d_215_v2_
                nw27_[int(4)] = d_215_v2_
                nw27_[int(5)] = d_215_v2_
                nw27_[int(6)] = d_215_v2_
                nw27_[int(7)] = (D2_DC6(default__.fm4(d_213_v0_, True, (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_218_globalState_))).cf8
                nw27_[int(8)] = d_215_v2_
                nw27_[int(9)] = d_215_v2_
                nw27_[int(10)] = d_215_v2_
                nw27_[int(11)] = default__.fm4(d_256_cf2_, d_213_v0_, (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_218_globalState_)
                nw27_[int(12)] = d_215_v2_
                nw27_[int(13)] = (d_215_v2_ if d_213_v0_ else _dafny.CodePoint('y'))
                nw27_[int(14)] = d_215_v2_
                nw27_[int(15)] = d_215_v2_
                nw27_[int(16)] = d_215_v2_
                nw27_[int(17)] = (d_215_v2_ if not(d_256_cf2_) else d_215_v2_)
                nw27_[int(18)] = d_215_v2_
                nw27_[int(19)] = d_215_v2_
                d_296_v61_ = nw27_
                index37_ = default__.safeIndex(301, (d_296_v61_).length(0))
                (d_296_v61_)[index37_] = d_215_v2_
                index38_ = default__.safeIndex(301, (d_296_v61_).length(0))
                (d_296_v61_)[index38_] = default__.fm4((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], not(default__.fm0((0) - (-842), d_218_globalState_)), (_dafny.MultiSet(d_287_v54_)).issubset((d_216_v3_).set(d_213_v0_, default__.abs(d_234_v17_))), d_218_globalState_)
                (d_218_globalState_).f3 = ((d_214_v1_) | (d_214_v1_) if not((D0_DC2(d_256_cf2_, d_219_v5_)).cf2) else d_214_v1_)
            elif True:
                d_297___mcc_h6_ = source12_.cf2
                d_298___mcc_h7_ = source12_.cf3
                d_299_cf3_ = d_298___mcc_h7_
                d_300_cf2_ = d_297___mcc_h6_
                d_301_v62_: D2
                d_301_v62_ = D2_DC6(d_215_v2_)
                pat_let_tv4_ = d_215_v2_
                d_302_v63_: _dafny.MultiSet
                def iife12_(_pat_let0_0):
                    def iife13_(d_303_dt__update__tmp_h0_):
                        def iife14_(_pat_let1_0):
                            def iife15_(d_304_dt__update_hcf8_h0_):
                                return D2_DC6(d_304_dt__update_hcf8_h0_)
                            return iife15_(_pat_let1_0)
                        return iife14_(pat_let_tv4_)
                    return iife13_(_pat_let0_0)
                d_302_v63_ = _dafny.MultiSet([(d_257_v30_)[default__.safeIndex(d_234_v17_, len(d_257_v30_))], (iife12_(d_301_v62_)).cf8])
                index39_ = default__.safeIndex(318, (d_219_v5_).length(0))
                (d_219_v5_)[index39_] = (925) != (((d_302_v63_)[d_215_v2_] if (d_215_v2_) in (d_302_v63_) else default__.fm1(d_218_globalState_)))
                (d_218_globalState_).f1 = default__.fm1(d_218_globalState_)
                d_305_v64_: _dafny.Seq
                d_305_v64_ = _dafny.SeqWithoutIsStrInference([(d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]])
                d_306_v65_: _dafny.Set
                d_306_v65_ = _dafny.Set({d_305_v64_})
                index40_ = default__.safeIndex(318, (d_219_v5_).length(0))
                (d_219_v5_)[index40_] = (d_306_v65_).ispropersubset((d_306_v65_) | (_dafny.Set({d_305_v64_})))
                d_307_v66_: _dafny.Array
                nw28_ = _dafny.Array(int(0), 6)
                d_307_v66_ = nw28_
                d_308_v67_: D1
                d_308_v67_ = D1_DC4(d_234_v17_, d_307_v66_)
                index41_ = default__.safeIndex(423, (d_307_v66_).length(0))
                (d_307_v66_)[index41_] = (d_234_v17_) * ((d_308_v67_).cf5)
                index42_ = default__.safeIndex(423, (d_307_v66_).length(0))
                rhs35_ = 984
                rhs36_ = d_234_v17_
                lhs25_ = d_307_v66_
                lhs26_ = default__.safeIndex(423, (d_307_v66_).length(0))
                lhs27_ = d_218_globalState_
                lhs25_[lhs26_] = rhs35_
                lhs27_.f0 = rhs36_
        d_309_v68_: _dafny.Seq
        d_309_v68_ = _dafny.SeqWithoutIsStrInference([d_214_v1_])
        d_310_v69_: _dafny.Set
        d_310_v69_ = _dafny.Set({(d_309_v68_)[default__.safeIndex(d_234_v17_, len(d_309_v68_))]})
        d_311_v70_: D3
        d_311_v70_ = D3_DC9(d_310_v69_)
        index43_ = default__.safeIndex(318, (d_219_v5_).length(0))
        (d_219_v5_)[index43_] = not(((d_311_v70_).cf13).issubset(d_310_v69_))
        hi1_ = d_234_v17_
        for d_312_i7_ in range(d_234_v17_, hi1_):
            (d_218_globalState_).f7 = d_234_v17_
            d_313_v71_: _dafny.Seq
            d_313_v71_ = _dafny.SeqWithoutIsStrInference([not(d_213_v0_)])
            d_314_v72_: _dafny.MultiSet
            d_314_v72_ = _dafny.MultiSet([d_234_v17_])
            d_315_v73_: _dafny.Map
            d_315_v73_ = _dafny.Map({d_314_v72_: 664})
            d_316_v74_: _dafny.Array
            nw29_ = _dafny.Array(None, 19)
            nw29_[int(0)] = d_312_i7_
            nw29_[int(1)] = d_312_i7_
            nw29_[int(2)] = (d_234_v17_) + (d_234_v17_)
            nw29_[int(3)] = 368
            nw29_[int(4)] = d_312_i7_
            nw29_[int(5)] = len(d_313_v71_)
            nw29_[int(6)] = (default__.fm1(d_218_globalState_)) * (409)
            nw29_[int(7)] = (len(_dafny.Map({(d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]: len(d_214_v1_)}))) - (344)
            nw29_[int(8)] = len(d_313_v71_)
            nw29_[int(9)] = ((d_314_v72_).cardinality) + (d_312_i7_)
            nw29_[int(10)] = (len(d_313_v71_) if (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))] else d_312_i7_)
            nw29_[int(11)] = d_234_v17_
            nw29_[int(12)] = d_234_v17_
            nw29_[int(13)] = d_312_i7_
            nw29_[int(14)] = default__.fm1(d_218_globalState_)
            nw29_[int(15)] = len(d_315_v73_)
            nw29_[int(16)] = default__.safeModuloInt(d_234_v17_, d_234_v17_)
            nw29_[int(17)] = d_312_i7_
            nw29_[int(18)] = (d_234_v17_) * (d_312_i7_)
            d_316_v74_ = nw29_
            d_316_v74_ = d_316_v74_
            (d_218_globalState_).f1 = d_234_v17_
            index44_ = default__.safeIndex(318, (d_219_v5_).length(0))
            (d_219_v5_)[index44_] = d_213_v0_
        if (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]:
            d_213_v0_ = not(d_213_v0_)
            d_317_v75_: C1
            nw30_ = C1()
            nw30_.ctor__((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))])
            d_317_v75_ = nw30_
            d_318_v76_: _dafny.Array
            nw31_ = _dafny.Array(int(0), 6)
            d_318_v76_ = nw31_
            index45_ = default__.safeIndex(895, (d_318_v76_).length(0))
            (d_318_v76_)[index45_] = d_234_v17_
            index46_ = default__.safeIndex(895, (d_318_v76_).length(0))
            (d_318_v76_)[index46_] = default__.safeModuloInt(default__.safeModuloInt(d_234_v17_, d_234_v17_), (438) - ((d_317_v75_).fm9(d_217_v4_, _dafny.Set({_dafny.Map({d_213_v0_: (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]})}), d_218_globalState_)))
            d_319_v77_: C5
            nw32_ = C5()
            nw32_.ctor__()
            d_319_v77_ = nw32_
            index47_ = default__.safeIndex(318, (d_219_v5_).length(0))
            (d_219_v5_)[index47_] = not((d_317_v75_).fm6(445, (d_318_v76_)[default__.safeIndex(895, (d_318_v76_).length(0))], (0) - ((d_234_v17_) - ((d_318_v76_)[default__.safeIndex(895, (d_318_v76_).length(0))])), (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_218_globalState_))
        elif True:
            index48_ = default__.safeIndex(318, (d_219_v5_).length(0))
            rhs37_ = _dafny.SeqWithoutIsStrInference([d_215_v2_ for d_320_i8_ in range(default__.abs(-10))])
            rhs38_ = d_213_v0_
            lhs28_ = d_219_v5_
            lhs29_ = default__.safeIndex(318, (d_219_v5_).length(0))
            d_217_v4_ = rhs37_
            lhs28_[lhs29_] = rhs38_
            (d_218_globalState_).f7 = (len(d_217_v4_)) - ((372) * (d_234_v17_))
            d_321_v78_: _dafny.Map
            d_321_v78_ = _dafny.Map({(d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]: d_217_v4_})
            d_322_v79_: _dafny.Seq
            d_322_v79_ = _dafny.SeqWithoutIsStrInference([d_213_v0_, False])
            d_321_v78_ = (d_321_v78_).set(((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]) in (_dafny.MultiSet([(d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], (default__.fm65(d_213_v0_, d_322_v79_, d_234_v17_, d_218_globalState_)).cf16])), d_217_v4_)
            d_323_v80_: _dafny.Array
            nw33_ = _dafny.Array(None, 19)
            nw33_[int(0)] = d_215_v2_
            nw33_[int(1)] = d_215_v2_
            nw33_[int(2)] = d_215_v2_
            nw33_[int(3)] = default__.fm4((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_213_v0_, d_213_v0_, d_218_globalState_)
            nw33_[int(4)] = (d_215_v2_ if not((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]) else _dafny.CodePoint('v'))
            nw33_[int(5)] = d_215_v2_
            nw33_[int(6)] = d_215_v2_
            nw33_[int(7)] = (default__.fm4(not(not(d_213_v0_)), (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_213_v0_, d_218_globalState_) if (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))] else d_215_v2_)
            nw33_[int(8)] = d_215_v2_
            nw33_[int(9)] = _dafny.CodePoint('u')
            nw33_[int(10)] = d_215_v2_
            nw33_[int(11)] = d_215_v2_
            nw33_[int(12)] = d_215_v2_
            nw33_[int(13)] = d_215_v2_
            nw33_[int(14)] = _dafny.CodePoint('w')
            nw33_[int(15)] = d_215_v2_
            nw33_[int(16)] = d_215_v2_
            nw33_[int(17)] = d_215_v2_
            nw33_[int(18)] = d_215_v2_
            d_323_v80_ = nw33_
            d_324_v81_: D12
            d_324_v81_ = D12_DC37(d_234_v17_, d_323_v80_, d_215_v2_)
            d_325_v82_: _dafny.Map
            d_325_v82_ = _dafny.Map({d_234_v17_: (d_324_v81_).cf62})
            d_326_v83_: _dafny.MultiSet
            d_326_v83_ = _dafny.MultiSet([d_234_v17_, default__.fm1(d_218_globalState_), 705, d_234_v17_, d_234_v17_])
            rhs39_ = ((d_325_v82_)[(d_326_v83_).cardinality] if ((d_326_v83_).cardinality) in (d_325_v82_) else ((d_325_v82_)[d_234_v17_] if (d_234_v17_) in (d_325_v82_) else d_323_v80_))
            rhs40_ = d_215_v2_
            rhs41_ = (0) - ((0) - ((d_234_v17_) - (d_234_v17_)))
            lhs30_ = d_218_globalState_
            d_323_v80_ = rhs39_
            d_215_v2_ = rhs40_
            lhs30_.f7 = rhs41_
            (d_218_globalState_).f1 = d_234_v17_
        d_213_v0_ = d_213_v0_
        d_327_v84_: D13
        d_327_v84_ = D13_DC39(d_234_v17_, True)
        pat_let_tv5_ = d_233_v16_
        pat_let_tv6_ = d_233_v16_
        pat_let_tv7_ = d_233_v16_
        def lambda8_(source14_):
            if source14_.is_DC39:
                d_328___mcc_h10_ = source14_.cf65
                d_329___mcc_h11_ = source14_.cf66
                d_330_cf66_ = d_329___mcc_h11_
                d_331_cf65_ = d_328___mcc_h10_
                return D7_DC22(pat_let_tv5_)
            elif source14_.is_DC40:
                d_332___mcc_h12_ = source14_.cf67
                d_333___mcc_h13_ = source14_.cf68
                d_334___mcc_h14_ = source14_.cf69
                d_335___mcc_h15_ = source14_.cf70
                d_336_cf70_ = d_335___mcc_h15_
                d_337_cf69_ = d_334___mcc_h14_
                d_338_cf68_ = d_333___mcc_h13_
                d_339_cf67_ = d_332___mcc_h12_
                return D7_DC22(pat_let_tv6_)
            elif True:
                d_340___mcc_h16_ = source14_.cf64
                d_341_cf64_ = d_340___mcc_h16_
                return D7_DC22(pat_let_tv7_)

        source13_ = lambda8_(d_327_v84_)
        if source13_.is_DC23:
            index49_ = default__.safeIndex(318, (d_219_v5_).length(0))
            (d_219_v5_)[index49_] = (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]
            if d_213_v0_:
                (d_218_globalState_).f7 = d_234_v17_
                d_342_v85_: _dafny.Array
                nw34_ = _dafny.Array(D8.default()(), 3)
                d_342_v85_ = nw34_
                d_343_v86_: _dafny.Array
                nw35_ = _dafny.Array(D1.default()(), 25)
                d_343_v86_ = nw35_
                d_344_v87_: _dafny.Seq
                d_344_v87_ = _dafny.SeqWithoutIsStrInference([d_343_v86_, d_343_v86_])
                d_345_v88_: D8
                d_345_v88_ = D8_DC27(True, d_344_v87_, d_232_v15_, d_234_v17_)
                index50_ = default__.safeIndex(811, (d_342_v85_).length(0))
                (d_342_v85_)[index50_] = d_345_v88_
                d_346_v89_: C7
                nw36_ = C7()
                nw36_.ctor__((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_213_v0_)
                d_346_v89_ = nw36_
                d_347_v90_: _dafny.Map
                d_347_v90_ = _dafny.Map({d_346_v89_: d_215_v2_})
                index51_ = default__.safeIndex(811, (d_342_v85_).length(0))
                rhs42_ = d_345_v88_
                rhs43_ = (_dafny.Map({d_346_v89_: d_215_v2_})) | (_dafny.Map({d_346_v89_: d_215_v2_}))
                rhs44_ = (0) - (d_234_v17_)
                lhs31_ = d_342_v85_
                lhs32_ = default__.safeIndex(811, (d_342_v85_).length(0))
                lhs33_ = d_218_globalState_
                lhs31_[lhs32_] = rhs42_
                d_347_v90_ = rhs43_
                lhs33_.f0 = rhs44_
                (d_218_globalState_).f1 = d_234_v17_
                d_348_v91_: C6
                nw37_ = C6()
                nw37_.ctor__()
                d_348_v91_ = nw37_
                d_348_v91_ = d_348_v91_
                d_349_v92_: _dafny.Array
                nw38_ = _dafny.Array(None, 27)
                d_349_v92_ = nw38_
                d_350_v93_: _dafny.Array
                def lambda9_(d_351_v4_):
                    def lambda10_(d_352_i9_):
                        return (d_351_v4_) + (d_351_v4_)

                    return lambda10_

                init3_ = lambda9_(d_217_v4_)
                nw39_ = _dafny.Array(None, 25)
                for i0_3_ in range(nw39_.length(0)):
                    nw39_[i0_3_] = init3_(i0_3_)
                d_350_v93_ = nw39_
                index52_ = default__.safeIndex(628, (d_350_v93_).length(0))
                (d_350_v93_)[index52_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqicmgfh"))
                index53_ = default__.safeIndex(628, (d_350_v93_).length(0))
                rhs45_ = d_349_v92_
                rhs46_ = (0) - ((-265) + (d_234_v17_))
                rhs47_ = (d_217_v4_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsyncjph"))) + (d_217_v4_))
                lhs34_ = d_218_globalState_
                lhs35_ = d_350_v93_
                lhs36_ = default__.safeIndex(628, (d_350_v93_).length(0))
                d_349_v92_ = rhs45_
                lhs34_.f0 = rhs46_
                lhs35_[lhs36_] = rhs47_
            elif True:
                d_353_v94_: C2
                nw40_ = C2()
                nw40_.ctor__((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], (d_234_v17_) - (d_234_v17_))
                d_353_v94_ = nw40_
                d_354_v95_: _dafny.Seq
                d_355_v96_: _dafny.Seq
                out18_: _dafny.Seq
                out19_: _dafny.Seq
                out18_, out19_ = default__.m0(d_234_v17_, d_219_v5_, d_218_globalState_)
                d_354_v95_ = out18_
                d_355_v96_ = out19_
                d_215_v2_ = d_215_v2_
                d_356_v97_: C8
                nw41_ = C8()
                nw41_.ctor__((d_354_v95_)[default__.safeIndex((d_353_v94_).f22, len(d_354_v95_))])
                d_356_v97_ = nw41_
                d_357_v98_: _dafny.Array
                def lambda11_(d_358_v17_):
                    def lambda12_(d_359_i10_):
                        return default__.safeModuloInt(d_359_i10_, d_358_v17_)

                    return lambda12_

                init4_ = lambda11_(d_234_v17_)
                nw42_ = _dafny.Array(None, 6)
                for i0_4_ in range(nw42_.length(0)):
                    nw42_[i0_4_] = init4_(i0_4_)
                d_357_v98_ = nw42_
                d_360_v99_: _dafny.Seq
                d_360_v99_ = _dafny.SeqWithoutIsStrInference([d_357_v98_])
                d_361_v100_: D24
                d_361_v100_ = D24_DC68((d_360_v99_).set(default__.safeIndex(len(d_231_v14_), len(d_360_v99_)), d_357_v98_))
                d_362_v101_: _dafny.Map
                d_362_v101_ = _dafny.Map({(d_361_v100_).cf117: ((d_353_v94_).f21 if (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))] else (d_353_v94_).f21)})
                d_363_v102_: _dafny.Seq
                d_363_v102_ = _dafny.SeqWithoutIsStrInference([d_360_v99_])
                d_364_v103_: D4
                d_364_v103_ = D4_DC13(default__.fm66((d_353_v94_).f21, d_215_v2_, d_234_v17_, (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_218_globalState_))
                d_365_v104_: _dafny.Seq
                d_365_v104_ = _dafny.SeqWithoutIsStrInference([d_364_v103_])
                d_362_v101_ = (d_362_v101_).set((d_363_v102_)[default__.safeIndex(d_234_v17_, len(d_363_v102_))], (d_356_v97_).fm17((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_234_v17_, d_365_v104_, (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_218_globalState_))
            if not((d_234_v17_) > (d_234_v17_)):
                d_366_v105_: _dafny.Seq
                d_367_v106_: _dafny.Seq
                out20_: _dafny.Seq
                out21_: _dafny.Seq
                out20_, out21_ = default__.m0((-56) * (445), d_219_v5_, d_218_globalState_)
                d_366_v105_ = out20_
                d_367_v106_ = out21_
                d_368_v107_: _dafny.Map
                d_368_v107_ = _dafny.Map({d_219_v5_: d_234_v17_})
                d_369_v108_: _dafny.Seq
                d_370_v109_: _dafny.Seq
                out22_: _dafny.Seq
                out23_: _dafny.Seq
                out22_, out23_ = default__.m0(((d_368_v107_)[d_219_v5_] if (d_219_v5_) in (d_368_v107_) else 425), d_219_v5_, d_218_globalState_)
                d_369_v108_ = out22_
                d_370_v109_ = out23_
                d_219_v5_ = ((d_219_v5_ if d_213_v0_ else d_219_v5_) if d_213_v0_ else d_219_v5_)
                d_371_v110_: _dafny.Seq
                d_372_v111_: _dafny.Seq
                out24_: _dafny.Seq
                out25_: _dafny.Seq
                out24_, out25_ = default__.m0(d_234_v17_, d_219_v5_, d_218_globalState_)
                d_371_v110_ = out24_
                d_372_v111_ = out25_
                d_213_v0_ = default__.fm0(d_234_v17_, d_218_globalState_)
            elif True:
                d_373_v112_: _dafny.Seq
                d_374_v113_: _dafny.Seq
                out26_: _dafny.Seq
                out27_: _dafny.Seq
                out26_, out27_ = default__.m0(d_234_v17_, (D0_DC2((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_219_v5_)).cf3, d_218_globalState_)
                d_373_v112_ = out26_
                d_374_v113_ = out27_
                d_375_v114_: _dafny.Array
                nw43_ = _dafny.Array(int(0), 13)
                d_375_v114_ = nw43_
                index54_ = default__.safeIndex(955, (d_375_v114_).length(0))
                (d_375_v114_)[index54_] = d_234_v17_
                index55_ = default__.safeIndex(955, (d_375_v114_).length(0))
                (d_375_v114_)[index55_] = (0) - ((0) - (len(d_214_v1_)))
                d_376_v115_: D16
                d_376_v115_ = D16_DC49(d_234_v17_)
                d_376_v115_ = d_376_v115_
                d_377_v116_: T1
                nw44_ = C11()
                nw44_.ctor__((d_375_v114_)[default__.safeIndex(955, (d_375_v114_).length(0))], d_373_v112_, (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))])
                d_377_v116_ = nw44_
                d_377_v116_ = d_377_v116_
                d_378_v117_: C2
                nw45_ = C2()
                nw45_.ctor__((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_234_v17_)
                d_378_v117_ = nw45_
            d_379_v118_: C11
            nw46_ = C11()
            nw46_.ctor__(default__.fm1(d_218_globalState_), (d_217_v4_ if d_213_v0_ else d_217_v4_), (default__.fm67((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_234_v17_, d_234_v17_, d_218_globalState_)).cf109)
            d_379_v118_ = nw46_
        elif source13_.is_DC24:
            d_380___mcc_h8_ = source13_.cf41
            d_381_cf41_ = d_380___mcc_h8_
            if (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]:
                d_382_v119_: C12
                nw47_ = C12()
                nw47_.ctor__(d_213_v0_, d_213_v0_, d_213_v0_)
                d_382_v119_ = nw47_
                d_383_v120_: _dafny.Seq
                d_383_v120_ = _dafny.SeqWithoutIsStrInference([d_382_v119_])
                index56_ = default__.safeIndex(318, (d_219_v5_).length(0))
                (d_219_v5_)[index56_] = (len(d_383_v120_)) <= (d_234_v17_)
                d_384_v121_: C0
                nw48_ = C0()
                nw48_.ctor__()
                d_384_v121_ = nw48_
                d_215_v2_ = (_dafny.CodePoint('e') if d_382_v119_.f12 else d_215_v2_)
                d_385_v122_: _dafny.Seq
                d_385_v122_ = _dafny.SeqWithoutIsStrInference([d_217_v4_, d_217_v4_])
                d_386_v123_: _dafny.Map
                d_386_v123_ = _dafny.Map({_dafny.MultiSet([_dafny.CodePoint('p'), d_215_v2_]): True})
                d_387_v124_: T1
                nw49_ = C1()
                nw49_.ctor__(((d_386_v123_)[_dafny.MultiSet([_dafny.CodePoint('a')])] if (_dafny.MultiSet([_dafny.CodePoint('a')])) in (d_386_v123_) else (d_384_v121_).fm37(d_217_v4_, d_381_cf41_, (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], d_382_v119_.f12, d_218_globalState_)))
                d_387_v124_ = nw49_
                d_388_v125_: _dafny.Array
                nw50_ = _dafny.Array(int(0), 1)
                d_388_v125_ = nw50_
                index57_ = default__.safeIndex(732, (d_388_v125_).length(0))
                (d_388_v125_)[index57_] = d_381_cf41_
                d_389_v126_: _dafny.Array
                nw51_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                d_389_v126_ = nw51_
                d_390_v127_: D12
                d_390_v127_ = D12_DC37(-799, d_389_v126_, d_215_v2_)
                d_391_v128_: _dafny.MultiSet
                d_391_v128_ = _dafny.MultiSet([d_381_cf41_, d_234_v17_])
                d_392_v131_: D7
                d_392_v131_ = D7_DC22(_dafny.Map({d_381_cf41_: d_381_cf41_}))
                d_393_v132_: _dafny.Seq
                def iife16_():
                    coll12_ = _dafny.Map()
                    compr_12_: int
                    for compr_12_ in _dafny.IntegerRange(547, 470):
                        d_394_v129_: int = compr_12_
                        if ((547) <= (d_394_v129_)) and ((d_394_v129_) < (470)):
                            coll12_[default__.safeModuloInt(d_394_v129_, d_381_cf41_)] = d_381_cf41_
                    return _dafny.Map(coll12_)
                def iife17_():
                    coll13_ = _dafny.Map()
                    compr_13_: int
                    for compr_13_ in _dafny.IntegerRange(975, -989):
                        d_395_v130_: int = compr_13_
                        if ((975) <= (d_395_v130_)) and ((d_395_v130_) < (-989)):
                            coll13_[(d_395_v130_) - ((D4_DC15(((d_391_v128_)[len(d_231_v14_)] if (len(d_231_v14_)) in (d_391_v128_) else 341), d_391_v128_, d_215_v2_, d_381_cf41_)).cf23)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tolfmsbf")))
                    return _dafny.Map(coll13_)
                d_393_v132_ = _dafny.SeqWithoutIsStrInference([d_233_v16_, (d_233_v16_).set(d_234_v17_, d_381_cf41_), iife16_()
                , iife17_()
                , (d_392_v131_).cf40])
                d_396_v133_: D13
                d_396_v133_ = D13_DC40(d_390_v127_, True, ((d_391_v128_)[len(d_214_v1_)] if (len(d_214_v1_)) in (d_391_v128_) else d_234_v17_), len((d_393_v132_)[default__.safeIndex(d_381_cf41_, len(d_393_v132_))]))
                d_397_v134_: _dafny.Seq
                d_397_v134_ = _dafny.SeqWithoutIsStrInference([d_396_v133_])
                d_398_v135_: _dafny.MultiSet
                d_398_v135_ = _dafny.MultiSet([d_397_v134_])
                index58_ = default__.safeIndex(732, (d_388_v125_).length(0))
                rhs48_ = d_385_v122_
                rhs49_ = d_381_cf41_
                rhs50_ = d_387_v124_
                rhs51_ = ((d_398_v135_)[(d_397_v134_) + (d_397_v134_)] if ((d_397_v134_) + (d_397_v134_)) in (d_398_v135_) else d_381_cf41_)
                lhs37_ = d_218_globalState_
                lhs38_ = d_388_v125_
                lhs39_ = default__.safeIndex(732, (d_388_v125_).length(0))
                d_385_v122_ = rhs48_
                lhs37_.f0 = rhs49_
                d_387_v124_ = rhs50_
                lhs38_[lhs39_] = rhs51_
                d_399_v136_: _dafny.Array
                nw52_ = _dafny.Array(_dafny.Seq({}), 18)
                d_399_v136_ = nw52_
                d_399_v136_ = d_399_v136_
            elif True:
                d_400_v137_: D11
                d_400_v137_ = D11_DC32((d_231_v14_) | (d_231_v14_))
                d_400_v137_ = d_400_v137_
                d_401_v138_: D19
                d_401_v138_ = D19_DC58(len(d_217_v4_), (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))])
                d_402_v139_: _dafny.Map
                d_402_v139_ = _dafny.Map({d_401_v138_: (d_217_v4_) + (d_217_v4_)})
                pat_let_tv8_ = d_219_v5_
                pat_let_tv9_ = d_219_v5_
                def iife18_(_pat_let2_0):
                    def iife19_(d_403_dt__update__tmp_h1_):
                        def iife20_(_pat_let3_0):
                            def iife21_(d_404_dt__update_hcf100_h0_):
                                return D19_DC58((d_403_dt__update__tmp_h1_).cf99, d_404_dt__update_hcf100_h0_)
                            return iife21_(_pat_let3_0)
                        return iife20_((pat_let_tv9_)[default__.safeIndex(318, (pat_let_tv8_).length(0))])
                    return iife19_(_pat_let2_0)
                d_402_v139_ = (d_402_v139_).set(iife18_(d_401_v138_), (d_217_v4_) + (d_217_v4_))
                d_405_v140_: D20
                d_405_v140_ = D20_DC60(d_234_v17_)
                pat_let_tv10_ = d_381_cf41_
                def iife22_(_pat_let4_0):
                    def iife23_(d_406_dt__update__tmp_h2_):
                        def iife24_(_pat_let5_0):
                            def iife25_(d_407_dt__update_hcf102_h0_):
                                return D20_DC60(d_407_dt__update_hcf102_h0_)
                            return iife25_(_pat_let5_0)
                        return iife24_(pat_let_tv10_)
                    return iife23_(_pat_let4_0)
                d_405_v140_ = iife22_(D20_DC60((d_232_v15_)[default__.safeIndex(len(d_214_v1_), len(d_232_v15_))]))
                d_408_v141_: D3
                d_408_v141_ = D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrso")), d_213_v0_, d_213_v0_)
                d_409_v142_: _dafny.Map
                d_409_v142_ = _dafny.Map({(d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]: (0) - (d_234_v17_)})
                d_410_v143_: _dafny.Set
                d_410_v143_ = _dafny.Set({d_215_v2_})
                d_411_v144_: _dafny.Map
                d_411_v144_ = _dafny.Map({d_408_v141_: (default__.fm18(d_234_v17_, d_409_v142_, d_213_v0_, d_218_globalState_)).intersection(_dafny.MultiSet([len(d_410_v143_), d_381_cf41_, d_234_v17_]))})
                d_412_v145_: _dafny.MultiSet
                d_412_v145_ = _dafny.MultiSet([d_381_cf41_, d_381_cf41_])
                d_411_v144_ = (d_411_v144_).set(d_408_v141_, d_412_v145_)
                d_413_v146_: _dafny.Array
                nw53_ = _dafny.Array(None, 9)
                nw53_[int(0)] = d_215_v2_
                nw53_[int(1)] = d_215_v2_
                nw53_[int(2)] = d_215_v2_
                nw53_[int(3)] = d_215_v2_
                nw53_[int(4)] = d_215_v2_
                nw53_[int(5)] = d_215_v2_
                nw53_[int(6)] = d_215_v2_
                nw53_[int(7)] = d_215_v2_
                nw53_[int(8)] = (d_215_v2_ if (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))] else d_215_v2_)
                d_413_v146_ = nw53_
                index59_ = default__.safeIndex(230, (d_413_v146_).length(0))
                (d_413_v146_)[index59_] = d_215_v2_
                index60_ = default__.safeIndex(230, (d_413_v146_).length(0))
                (d_413_v146_)[index60_] = d_215_v2_
            d_414_v147_: _dafny.Seq
            d_414_v147_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpdogu")), d_217_v4_])
            d_415_v148_: _dafny.Map
            d_415_v148_ = _dafny.Map({d_234_v17_: _dafny.CodePoint('m')})
            rhs52_ = len(((d_217_v4_) + (d_217_v4_)) + ((d_414_v147_)[default__.safeIndex((0) - (d_234_v17_), len(d_414_v147_))]))
            rhs53_ = ((d_415_v148_)[d_381_cf41_] if (d_381_cf41_) in (d_415_v148_) else (d_215_v2_ if ((d_231_v14_)[d_381_cf41_] if (d_381_cf41_) in (d_231_v14_) else (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]) else d_215_v2_))
            lhs40_ = d_218_globalState_
            lhs40_.f1 = rhs52_
            d_215_v2_ = rhs53_
            d_217_v4_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfu")) if d_213_v0_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))
            d_416_v149_: _dafny.Array
            nw54_ = _dafny.Array(None, 18)
            nw54_[int(0)] = _dafny.CodePoint('v')
            nw54_[int(1)] = d_215_v2_
            nw54_[int(2)] = d_215_v2_
            nw54_[int(3)] = d_215_v2_
            nw54_[int(4)] = d_215_v2_
            nw54_[int(5)] = d_215_v2_
            nw54_[int(6)] = d_215_v2_
            nw54_[int(7)] = d_215_v2_
            nw54_[int(8)] = _dafny.CodePoint('x')
            nw54_[int(9)] = _dafny.CodePoint('a')
            nw54_[int(10)] = d_215_v2_
            nw54_[int(11)] = d_215_v2_
            nw54_[int(12)] = d_215_v2_
            nw54_[int(13)] = d_215_v2_
            nw54_[int(14)] = d_215_v2_
            nw54_[int(15)] = d_215_v2_
            nw54_[int(16)] = _dafny.CodePoint('o')
            nw54_[int(17)] = d_215_v2_
            d_416_v149_ = nw54_
            d_417_v150_: D12
            d_417_v150_ = D12_DC37(d_234_v17_, d_416_v149_, d_215_v2_)
            d_418_v151_: _dafny.Array
            nw55_ = _dafny.Array(None, 28)
            nw55_[int(0)] = d_215_v2_
            nw55_[int(1)] = d_215_v2_
            nw55_[int(2)] = d_215_v2_
            nw55_[int(3)] = d_215_v2_
            nw55_[int(4)] = d_215_v2_
            nw55_[int(5)] = d_215_v2_
            nw55_[int(6)] = d_215_v2_
            nw55_[int(7)] = d_215_v2_
            nw55_[int(8)] = d_215_v2_
            nw55_[int(9)] = d_215_v2_
            nw55_[int(10)] = (d_217_v4_)[default__.safeIndex(d_381_cf41_, len(d_217_v4_))]
            nw55_[int(11)] = d_215_v2_
            nw55_[int(12)] = d_215_v2_
            nw55_[int(13)] = d_215_v2_
            nw55_[int(14)] = _dafny.CodePoint('e')
            nw55_[int(15)] = d_215_v2_
            nw55_[int(16)] = d_215_v2_
            nw55_[int(17)] = d_215_v2_
            nw55_[int(18)] = d_215_v2_
            nw55_[int(19)] = _dafny.CodePoint('k')
            nw55_[int(20)] = d_215_v2_
            nw55_[int(21)] = d_215_v2_
            nw55_[int(22)] = d_215_v2_
            nw55_[int(23)] = (d_417_v150_).cf63
            nw55_[int(24)] = d_215_v2_
            nw55_[int(25)] = (d_215_v2_ if not((d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]) else d_215_v2_)
            nw55_[int(26)] = d_215_v2_
            nw55_[int(27)] = d_215_v2_
            d_418_v151_ = nw55_
            index61_ = default__.safeIndex(428, (d_416_v149_).length(0))
            (d_416_v149_)[index61_] = d_215_v2_
            index62_ = default__.safeIndex(428, (d_416_v149_).length(0))
            (d_416_v149_)[index62_] = (d_215_v2_ if True else _dafny.CodePoint('k'))
        elif True:
            d_419___mcc_h9_ = source13_.cf40
            d_420_cf40_ = d_419___mcc_h9_
            d_213_v0_ = (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]
            (d_218_globalState_).f7 = d_234_v17_
            (d_218_globalState_).f7 = d_234_v17_
            d_421_v152_: _dafny.Array
            nw56_ = _dafny.Array(_dafny.Map({}), 24)
            d_421_v152_ = nw56_
            d_422_v153_: _dafny.Map
            d_422_v153_ = _dafny.Map({d_213_v0_: d_215_v2_})
            d_423_v154_: _dafny.Map
            d_423_v154_ = _dafny.Map({d_422_v153_: (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]})
            d_424_v155_: _dafny.Map
            d_424_v155_ = d_423_v154_
            index63_ = default__.safeIndex(103, (d_421_v152_).length(0))
            (d_421_v152_)[index63_] = d_424_v155_
            d_425_v156_: _dafny.Array
            nw57_ = _dafny.Array(int(0), 14)
            d_425_v156_ = nw57_
            d_426_v157_: D2
            d_426_v157_ = D2_DC6(d_215_v2_)
            d_427_v158_: _dafny.Set
            d_427_v158_ = _dafny.Set({_dafny.Map({len(d_231_v14_): (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]})})
            d_428_v159_: _dafny.Seq
            d_428_v159_ = _dafny.SeqWithoutIsStrInference([(d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))], (d_219_v5_)[default__.safeIndex(318, (d_219_v5_).length(0))]])
            d_429_v161_: _dafny.Map
            def iife26_():
                coll14_ = _dafny.Set()
                compr_14_: _dafny.Map
                for compr_14_ in (_dafny.MultiSet([d_231_v14_, d_231_v14_, d_231_v14_, d_231_v14_])).Elements:
                    d_430_v160_: _dafny.Map = compr_14_
                    if (d_430_v160_) in (_dafny.MultiSet([d_231_v14_, d_231_v14_, d_231_v14_, d_231_v14_])):
                        coll14_ = coll14_.union(_dafny.Set([d_430_v160_]))
                return _dafny.Set(coll14_)
            d_429_v161_ = _dafny.Map({d_234_v17_: iife26_()
            })
            d_431_v162_: _dafny.Seq
            d_431_v162_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_234_v17_: d_213_v0_})])
            d_432_v164_: _dafny.Seq
            d_432_v164_ = _dafny.SeqWithoutIsStrInference([d_427_v158_, d_427_v158_])
            d_433_v165_: _dafny.Array
            nw58_ = _dafny.Array(None, 13)
            nw58_[int(0)] = (d_427_v158_).intersection(d_427_v158_)
            nw58_[int(1)] = _dafny.Set({d_231_v14_, _dafny.Map({len(d_217_v4_): False}), d_231_v14_, d_231_v14_, _dafny.Map({len(d_233_v16_): (d_428_v159_)[default__.safeIndex(d_234_v17_, len(d_428_v159_))]})})
            nw58_[int(2)] = d_427_v158_
            nw58_[int(3)] = d_427_v158_
            nw58_[int(4)] = ((d_429_v161_)[d_234_v17_] if (d_234_v17_) in (d_429_v161_) else d_427_v158_)
            nw58_[int(5)] = _dafny.Set({_dafny.Map({d_234_v17_: d_213_v0_}), d_231_v14_, d_231_v14_})
            def iife27_():
                coll15_ = _dafny.Set()
                compr_15_: _dafny.Map
                for compr_15_ in (d_431_v162_).Elements:
                    d_434_v163_: _dafny.Map = compr_15_
                    if (d_434_v163_) in (d_431_v162_):
                        coll15_ = coll15_.union(_dafny.Set([d_434_v163_]))
                return _dafny.Set(coll15_)
            nw58_[int(6)] = (iife27_()
            ).intersection((d_432_v164_)[default__.safeIndex(d_234_v17_, len(d_432_v164_))])
            nw58_[int(7)] = d_427_v158_
            nw58_[int(8)] = d_427_v158_
            nw58_[int(9)] = (d_427_v158_ if d_213_v0_ else (d_427_v158_))
            nw58_[int(10)] = d_427_v158_
            nw58_[int(11)] = (d_427_v158_) - (d_427_v158_)
            nw58_[int(12)] = (_dafny.Set({d_231_v14_})) - (d_427_v158_)
            d_433_v165_ = nw58_
            index64_ = default__.safeIndex(314, (d_433_v165_).length(0))
            (d_433_v165_)[index64_] = d_427_v158_
            index65_ = default__.safeIndex(103, (d_421_v152_).length(0))
            index66_ = default__.safeIndex(314, (d_433_v165_).length(0))
            rhs54_ = (d_217_v4_) + (default__.fm38(d_213_v0_, d_234_v17_, d_218_globalState_))
            rhs55_ = d_424_v155_
            rhs56_ = d_425_v156_
            rhs57_ = d_426_v157_
            rhs58_ = (d_427_v158_ if default__.fm0(d_234_v17_, d_218_globalState_) else d_427_v158_)
            lhs41_ = d_421_v152_
            lhs42_ = default__.safeIndex(103, (d_421_v152_).length(0))
            lhs43_ = d_433_v165_
            lhs44_ = default__.safeIndex(314, (d_433_v165_).length(0))
            d_217_v4_ = rhs54_
            lhs41_[lhs42_] = rhs55_
            d_425_v156_ = rhs56_
            d_426_v157_ = rhs57_
            lhs43_[lhs44_] = rhs58_
        d_435_v166_: C6
        nw59_ = C6()
        nw59_.ctor__()
        d_435_v166_ = nw59_
        d_436_v167_: _dafny.Map
        d_436_v167_ = _dafny.Map({d_234_v17_: d_435_v166_})
        d_437_v168_: _dafny.Set
        d_437_v168_ = _dafny.Set({len(d_232_v15_)})
        d_436_v167_ = (d_436_v167_).set(len(d_437_v168_), d_435_v166_)
        index67_ = default__.safeIndex(318, (d_219_v5_).length(0))
        rhs59_ = d_234_v17_
        rhs60_ = ((d_234_v17_) - (d_234_v17_)) <= (d_234_v17_)
        lhs45_ = d_218_globalState_
        lhs46_ = d_219_v5_
        lhs47_ = default__.safeIndex(318, (d_219_v5_).length(0))
        lhs45_.f7 = rhs59_
        lhs46_[lhs47_] = rhs60_
        d_438_v169_: _dafny.Array
        nw60_ = _dafny.Array(_dafny.Array(None, 0), 20)
        d_438_v169_ = nw60_
        d_439_v170_: _dafny.Array
        def lambda13_(d_440_v2_):
            def lambda14_(d_441_i11_):
                return d_440_v2_

            return lambda14_

        init5_ = lambda13_(d_215_v2_)
        nw61_ = _dafny.Array(None, 1)
        for i0_5_ in range(nw61_.length(0)):
            nw61_[i0_5_] = init5_(i0_5_)
        d_439_v170_ = nw61_
        index68_ = default__.safeIndex(483, (d_438_v169_).length(0))
        (d_438_v169_)[index68_] = d_439_v170_
        index69_ = default__.safeIndex(483, (d_438_v169_).length(0))
        (d_438_v169_)[index69_] = d_439_v170_
        d_442_v171_: D24
        d_442_v171_ = D24_DC69(d_435_v166_, -344, d_234_v17_)
        d_435_v166_ = (d_442_v171_).cf118
        d_443_v172_: C11
        nw62_ = C11()
        nw62_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwudkt"))), d_217_v4_, (d_234_v17_) != (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_444_i12_ in range(default__.abs(420))]))))
        d_443_v172_ = nw62_
        _dafny.print(_dafny.string_of(d_213_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_214_v1_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_215_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v3_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_217_v4_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_218_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_218_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_globalState_.f3) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f4) == (_dafny.Set({_dafny.CodePoint('c')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_218_globalState_).f5) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_218_globalState_).f6).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_218_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v5_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_222_v6_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_223_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_230_v13_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v14_) == (_dafny.Map({1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_v15_) == (_dafny.SeqWithoutIsStrInference([409, 2, -766, 0, -840, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544, 544]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v16_) == (_dafny.Map({7: 901}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_234_v17_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_235_v18_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v18_).cf3)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_v68_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_310_v69_) == (_dafny.Set({_dafny.Set({True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_311_v70_).cf13) == (_dafny.Set({_dafny.Set({True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v84_).cf65))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v84_).cf66))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_436_v167_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_437_v168_) == (_dafny.Set({80}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_438_v169_)[3])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_439_v170_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_442_v171_).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_442_v171_).cf120))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_443_v172_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_443_v172_).f14).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({self.cf1.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(int(0), _dafny.Array(None, 0))
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

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC7(D2, NamedTuple('DC7', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, False)
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

class D3_DC10(D3, NamedTuple('DC10', [('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({self.cf14.VerbatimString(True)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14(int(0), int(0))
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

class D4_DC14(D4, NamedTuple('DC14', [('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf23', Any), ('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC17(False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)

class D5_DC17(D5, NamedTuple('DC17', [('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC20(False, False, False, int(0))
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

class D6_DC20(D6, NamedTuple('DC20', [('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC21(D6, NamedTuple('DC21', [('cf36', Any), ('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC23()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D7_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D7_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)

class D7_DC23(D7, NamedTuple('DC23', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC23'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC23)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC24(D7, NamedTuple('DC24', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC24({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC24) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC22(D7, NamedTuple('DC22', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC26()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D8_DC27)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D8_DC28)

class D8_DC26(D8, NamedTuple('DC26', [])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26)
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC27(D8, NamedTuple('DC27', [('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC27({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC27) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC28(D8, NamedTuple('DC28', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC28({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC28) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D9_DC29)

class D9_DC29(D9, NamedTuple('DC29', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC29({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC29) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC31(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D10_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D10_DC30)

class D10_DC31(D10, NamedTuple('DC31', [('cf50', Any), ('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC31({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC31) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC30(D10, NamedTuple('DC30', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC33(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D11_DC34)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)

class D11_DC33(D11, NamedTuple('DC33', [('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC34(D11, NamedTuple('DC34', [('cf56', Any), ('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC34({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC34) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC36(_dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D12_DC37)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)

class D12_DC36(D12, NamedTuple('DC36', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC37(D12, NamedTuple('DC37', [('cf61', Any), ('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37({_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37) and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC35(D12, NamedTuple('DC35', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC39(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D13_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D13_DC40)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)

class D13_DC39(D13, NamedTuple('DC39', [('cf65', Any), ('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC39({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC39) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC40(D13, NamedTuple('DC40', [('cf67', Any), ('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC40({_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC40) and self.cf67 == __o.cf67 and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC38(D13, NamedTuple('DC38', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC42(D1.default()(), _dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D14_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D14_DC43)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)

class D14_DC42(D14, NamedTuple('DC42', [('cf72', Any), ('cf73', Any), ('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC42({_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC42) and self.cf72 == __o.cf72 and self.cf73 == __o.cf73 and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC43(D14, NamedTuple('DC43', [('cf75', Any), ('cf76', Any), ('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC43({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC43) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76 and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC41(D14, NamedTuple('DC41', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC45(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D15_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D15_DC46)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D15_DC44)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D15_DC47)

class D15_DC45(D15, NamedTuple('DC45', [('cf80', Any), ('cf81', Any), ('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC45({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC45) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81 and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC46(D15, NamedTuple('DC46', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC46'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC46)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC44(D15, NamedTuple('DC44', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC44({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC44) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC47(D15, NamedTuple('DC47', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC47({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC47) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC49(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D16_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D16_DC48)

class D16_DC49(D16, NamedTuple('DC49', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC49({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC49) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC48(D16, NamedTuple('DC48', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC48({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC48) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC51(_dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D17_DC51)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D17_DC50)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D17_DC52)

class D17_DC51(D17, NamedTuple('DC51', [('cf87', Any), ('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC51({_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC51) and self.cf87 == __o.cf87 and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC50(D17, NamedTuple('DC50', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC50({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC50) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC52(D17, NamedTuple('DC52', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC52({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC52) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC54(False, int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D18_DC54)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D18_DC55)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D18_DC56)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D18_DC53)

class D18_DC54(D18, NamedTuple('DC54', [('cf91', Any), ('cf92', Any), ('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC54({_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC54) and self.cf91 == __o.cf91 and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC55(D18, NamedTuple('DC55', [('cf95', Any), ('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC55({_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC55) and self.cf95 == __o.cf95 and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC56(D18, NamedTuple('DC56', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC56({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC56) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC53(D18, NamedTuple('DC53', [('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC53({_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC53) and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC58(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D19_DC58)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D19_DC57)

class D19_DC58(D19, NamedTuple('DC58', [('cf99', Any), ('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC58({_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC58) and self.cf99 == __o.cf99 and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC57(D19, NamedTuple('DC57', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC57({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC57) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC60(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D20_DC60)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D20_DC59)

class D20_DC60(D20, NamedTuple('DC60', [('cf102', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC60({_dafny.string_of(self.cf102)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC60) and self.cf102 == __o.cf102
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC59(D20, NamedTuple('DC59', [('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC59({_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC59) and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC62(_dafny.Map({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D21_DC62)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D21_DC63)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D21_DC61)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D21_DC64)

class D21_DC62(D21, NamedTuple('DC62', [('cf104', Any), ('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC62({_dafny.string_of(self.cf104)}, {_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC62) and self.cf104 == __o.cf104 and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC63(D21, NamedTuple('DC63', [('cf106', Any), ('cf107', Any), ('cf108', Any), ('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC63({_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)}, {_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC63) and self.cf106 == __o.cf106 and self.cf107 == __o.cf107 and self.cf108 == __o.cf108 and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC61(D21, NamedTuple('DC61', [('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC61({_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC61) and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC64(D21, NamedTuple('DC64', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC64({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC64) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC66(_dafny.CodePoint('D'), False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D22_DC66)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D22_DC65)

class D22_DC66(D22, NamedTuple('DC66', [('cf112', Any), ('cf113', Any), ('cf114', Any), ('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC66({_dafny.string_of(self.cf112)}, {_dafny.string_of(self.cf113)}, {_dafny.string_of(self.cf114)}, {self.cf115.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC66) and self.cf112 == __o.cf112 and self.cf113 == __o.cf113 and self.cf114 == __o.cf114 and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC65(D22, NamedTuple('DC65', [('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC65({_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC65) and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D23_DC67)

class D23_DC67(D23, NamedTuple('DC67', [('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC67({_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC67) and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC69(None, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D24_DC69)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D24_DC68)

class D24_DC69(D24, NamedTuple('DC69', [('cf118', Any), ('cf119', Any), ('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC69({_dafny.string_of(self.cf118)}, {_dafny.string_of(self.cf119)}, {_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC69) and self.cf118 == __o.cf118 and self.cf119 == __o.cf119 and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC68(D24, NamedTuple('DC68', [('cf117', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC68({_dafny.string_of(self.cf117)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC68) and self.cf117 == __o.cf117
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D25_DC70)

class D25_DC70(D25, NamedTuple('DC70', [('cf121', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC70({_dafny.string_of(self.cf121)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC70) and self.cf121 == __o.cf121
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D26_DC71)

class D26_DC71(D26, NamedTuple('DC71', [('cf122', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC71({_dafny.string_of(self.cf122)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC71) and self.cf122 == __o.cf122
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC73(False, int(0), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D27_DC73)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D27_DC74)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D27_DC75)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D27_DC72)

class D27_DC73(D27, NamedTuple('DC73', [('cf124', Any), ('cf125', Any), ('cf126', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC73({_dafny.string_of(self.cf124)}, {_dafny.string_of(self.cf125)}, {_dafny.string_of(self.cf126)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC73) and self.cf124 == __o.cf124 and self.cf125 == __o.cf125 and self.cf126 == __o.cf126
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC74(D27, NamedTuple('DC74', [('cf127', Any), ('cf128', Any), ('cf129', Any), ('cf130', Any), ('cf131', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC74({_dafny.string_of(self.cf127)}, {_dafny.string_of(self.cf128)}, {_dafny.string_of(self.cf129)}, {_dafny.string_of(self.cf130)}, {_dafny.string_of(self.cf131)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC74) and self.cf127 == __o.cf127 and self.cf128 == __o.cf128 and self.cf129 == __o.cf129 and self.cf130 == __o.cf130 and self.cf131 == __o.cf131
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC75(D27, NamedTuple('DC75', [('cf132', Any), ('cf133', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC75({_dafny.string_of(self.cf132)}, {_dafny.string_of(self.cf133)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC75) and self.cf132 == __o.cf132 and self.cf133 == __o.cf133
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC72(D27, NamedTuple('DC72', [('cf123', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC72({_dafny.string_of(self.cf123)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC72) and self.cf123 == __o.cf123
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC77(int(0), _dafny.Set({}), int(0), False, _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D28_DC77)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D28_DC76)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D28_DC78)

class D28_DC77(D28, NamedTuple('DC77', [('cf135', Any), ('cf136', Any), ('cf137', Any), ('cf138', Any), ('cf139', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC77({_dafny.string_of(self.cf135)}, {_dafny.string_of(self.cf136)}, {_dafny.string_of(self.cf137)}, {_dafny.string_of(self.cf138)}, {_dafny.string_of(self.cf139)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC77) and self.cf135 == __o.cf135 and self.cf136 == __o.cf136 and self.cf137 == __o.cf137 and self.cf138 == __o.cf138 and self.cf139 == __o.cf139
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC76(D28, NamedTuple('DC76', [('cf134', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC76({_dafny.string_of(self.cf134)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC76) and self.cf134 == __o.cf134
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC78(D28, NamedTuple('DC78', [('cf140', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC78({_dafny.string_of(self.cf140)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC78) and self.cf140 == __o.cf140
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC80(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D29_DC80)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D29_DC81)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D29_DC79)

class D29_DC80(D29, NamedTuple('DC80', [('cf142', Any), ('cf143', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC80({_dafny.string_of(self.cf142)}, {_dafny.string_of(self.cf143)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC80) and self.cf142 == __o.cf142 and self.cf143 == __o.cf143
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC81(D29, NamedTuple('DC81', [('cf144', Any), ('cf145', Any), ('cf146', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC81({_dafny.string_of(self.cf144)}, {_dafny.string_of(self.cf145)}, {_dafny.string_of(self.cf146)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC81) and self.cf144 == __o.cf144 and self.cf145 == __o.cf145 and self.cf146 == __o.cf146
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC79(D29, NamedTuple('DC79', [('cf141', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC79({_dafny.string_of(self.cf141)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC79) and self.cf141 == __o.cf141
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m3(self, p0, p1, globalState):
        pass


class T1:
    pass
    def m4(self, globalState):
        pass

    def m5(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f1: int = int(0)
        self.f3: _dafny.Set = _dafny.Set({})
        self.f7: int = int(0)
        self._f2: bool = False
        self._f4: _dafny.Set = _dafny.Set({})
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        self._f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7):
        (self).f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7

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

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm36(self, globalState):
        return (default__.safeModuloInt(len(_dafny.Map({_dafny.CodePoint('e'): 744})), (0) - (-290))) - (len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([306])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-286, len(_dafny.SeqWithoutIsStrInference([304, -160]))])]))))

    def fm37(self, p0, p1, p2, p3, globalState):
        return False


class C1(T1, T0):
    def  __init__(self):
        self._f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f10):
        (self)._f10 = f10

    def fm8(self, p0, p1, p2, globalState):
        return (_dafny.MultiSet([(self).f10, (self).f10, False])).ispropersubset(_dafny.MultiSet([not((self).f10)]))

    def fm9(self, p0, p1, globalState):
        return (((0) - (-283)) - (-605)) * (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_445_i0_ in range(default__.abs(622))]) if (self).f10 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l')]))))

    def fm6(self, p0, p1, p2, p3, globalState):
        return not(((_dafny.Map({(self).f10: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mljbys"))}) if (self).f10 else _dafny.Map({(self).f10: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gekebnpd"))}))) == (_dafny.Map({(self).f10: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vup"))})))

    def fm7(self, p0, globalState):
        return not((_dafny.MultiSet([(0) - (-457)])) == ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(self).f10]))])) | (_dafny.MultiSet([len(_dafny.Map({len(_dafny.Set({(self).f10})): (0) - (len(_dafny.Map({(self).f10: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xgpysmdex")))})))})), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mifruf")))), (0) - (len(_dafny.SeqWithoutIsStrInference([125 for d_446_i0_ in range(default__.abs(55))])))]))))

    def m4(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: int = int(0)
        d_447_v0_: int
        d_447_v0_ = 747
        pat_let_tv11_ = d_447_v0_
        pat_let_tv12_ = d_447_v0_
        def lambda15_(source15_):
            if source15_.is_DC4:
                d_448___mcc_h0_ = source15_.cf5
                d_449___mcc_h1_ = source15_.cf6
                d_450_cf6_ = d_449___mcc_h1_
                d_451_cf5_ = d_448___mcc_h0_
                return len(_dafny.SeqWithoutIsStrInference([(self).f10, not((self).f10)]))
            elif source15_.is_DC3:
                d_452___mcc_h2_ = source15_.cf4
                d_453_cf4_ = d_452___mcc_h2_
                return 163
            elif True:
                d_454___mcc_h3_ = source15_.cf7
                d_455_cf7_ = d_454___mcc_h3_
                return default__.safeDivisionInt(pat_let_tv11_, pat_let_tv12_)

        (globalState).f7 = lambda15_(D1_DC3(d_447_v0_))
        d_456_v1_: _dafny.Array
        def lambda16_(d_457_i0_):
            return default__.safeDivisionInt(d_457_i0_, 287)

        init6_ = lambda16_
        nw63_ = _dafny.Array(None, 23)
        for i0_6_ in range(nw63_.length(0)):
            nw63_[i0_6_] = init6_(i0_6_)
        d_456_v1_ = nw63_
        index70_ = default__.safeIndex(330, (d_456_v1_).length(0))
        (d_456_v1_)[index70_] = d_447_v0_
        d_458_v2_: str
        d_458_v2_ = _dafny.CodePoint('r')
        index71_ = default__.safeIndex(330, (d_456_v1_).length(0))
        (d_456_v1_)[index71_] = default__.safeDivisionInt(default__.safeDivisionInt((_dafny.MultiSet([d_458_v2_])).cardinality, d_447_v0_), (d_447_v0_) * (d_447_v0_))
        d_459_v3_: _dafny.Array
        def lambda17_(d_460_i1_):
            return (self).f10

        init7_ = lambda17_
        nw64_ = _dafny.Array(None, 25)
        for i0_7_ in range(nw64_.length(0)):
            nw64_[i0_7_] = init7_(i0_7_)
        d_459_v3_ = nw64_
        d_461_v4_: _dafny.Seq
        d_461_v4_ = _dafny.SeqWithoutIsStrInference([d_459_v3_])
        d_462_v5_: _dafny.Map
        d_462_v5_ = _dafny.Map({not((self).f10): d_461_v4_})
        d_461_v4_ = ((d_462_v5_)[(self).f10] if ((self).f10) in (d_462_v5_) else _dafny.SeqWithoutIsStrInference([d_459_v3_]))
        r1 = (self).f10
        d_463_v6_: _dafny.Seq
        d_463_v6_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10, False])
        d_464_v7_: _dafny.Seq
        d_464_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "raqvjn"))
        d_465_v8_: _dafny.Map
        d_465_v8_ = _dafny.Map({d_464_v7_: d_458_v2_})
        index72_ = default__.safeIndex(330, (d_456_v1_).length(0))
        (d_456_v1_)[index72_] = (((d_456_v1_)[default__.safeIndex(330, (d_456_v1_).length(0))]) + (len(d_465_v8_)) if (d_463_v6_) <= (d_463_v6_) else d_447_v0_)
        index73_ = default__.safeIndex(330, (d_456_v1_).length(0))
        (d_456_v1_)[index73_] = d_447_v0_
        d_466_v9_: _dafny.Map
        d_466_v9_ = _dafny.Map({d_464_v7_: (self).f10})
        d_467_v11_: _dafny.Set
        d_467_v11_ = _dafny.Set({d_464_v7_, d_464_v7_})
        def iife28_():
            coll16_ = _dafny.Set()
            compr_16_: _dafny.Seq
            for compr_16_ in (d_466_v9_).keys.Elements:
                d_468_v10_: _dafny.Seq = compr_16_
                if (d_468_v10_) in (d_466_v9_):
                    coll16_ = coll16_.union(_dafny.Set([d_468_v10_]))
            return _dafny.Set(coll16_)
        r0 = ((iife28_()
        ).intersection(d_467_v11_)).intersection(d_467_v11_)
        d_469_v12_: _dafny.Set
        d_469_v12_ = _dafny.Set({(d_456_v1_)[default__.safeIndex(330, (d_456_v1_).length(0))], (d_456_v1_)[default__.safeIndex(330, (d_456_v1_).length(0))], (d_456_v1_)[default__.safeIndex(330, (d_456_v1_).length(0))], 796, d_447_v0_})
        r1 = (not((self).f10) if ((self).f10) or (not((self).f10)) else (d_469_v12_).isdisjoint(d_469_v12_))
        r2 = d_447_v0_
        return r0, r1, r2

    def m5(self, p0, globalState):
        d_470_v0_: _dafny.Seq
        d_470_v0_ = _dafny.SeqWithoutIsStrInference([not((self).f10)])
        d_470_v0_ = ((d_470_v0_) + (_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10, (self).f10])) if (self).f10 else d_470_v0_)
        d_471_v1_: _dafny.Set
        d_471_v1_ = _dafny.Set({(self).f10, (self).f10})
        d_472_v2_: _dafny.Seq
        d_472_v2_ = _dafny.SeqWithoutIsStrInference([p0, p0, len(d_471_v1_)])
        d_473_v3_: _dafny.Array
        nw65_ = _dafny.Array(int(0), 24)
        d_473_v3_ = nw65_
        d_474_v4_: _dafny.MultiSet
        d_474_v4_ = _dafny.MultiSet([d_473_v3_])
        d_475_v5_: _dafny.Map
        d_475_v5_ = _dafny.Map({(self).f10: (self).f10})
        d_476_v6_: _dafny.Array
        nw66_ = _dafny.Array(None, 14)
        nw66_[int(0)] = (d_472_v2_) < (d_472_v2_)
        nw66_[int(1)] = (self).f10
        nw66_[int(2)] = (d_470_v0_) not in (_dafny.SeqWithoutIsStrInference([d_470_v0_, d_470_v0_, d_470_v0_, (d_470_v0_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0 for d_477_i0_ in range(default__.abs(676))])), len(d_470_v0_)), (self).f10), d_470_v0_]))
        nw66_[int(3)] = (_dafny.MultiSet([d_473_v3_])).isdisjoint(d_474_v4_)
        nw66_[int(4)] = (p0) == (p0)
        nw66_[int(5)] = (d_471_v1_).isdisjoint(d_471_v1_)
        nw66_[int(6)] = not(not((self).fm6((0) - (p0), default__.fm1(globalState), p0, (self).f10, globalState)))
        nw66_[int(7)] = (((d_475_v5_)[(d_470_v0_)[default__.safeIndex(p0, len(d_470_v0_))]] if ((d_470_v0_)[default__.safeIndex(p0, len(d_470_v0_))]) in (d_475_v5_) else (self).f10)) and ((self).f10)
        nw66_[int(8)] = (self).f10
        nw66_[int(9)] = not((self).f10)
        nw66_[int(10)] = (self).f10
        nw66_[int(11)] = (self).f10
        nw66_[int(12)] = not ((self).f10) or (not((self).f10))
        nw66_[int(13)] = (self).f10
        d_476_v6_ = nw66_
        d_478_v7_: _dafny.Seq
        d_478_v7_ = _dafny.SeqWithoutIsStrInference([d_476_v6_, d_476_v6_, d_476_v6_, d_476_v6_, d_476_v6_])
        d_476_v6_ = (d_478_v7_)[default__.safeIndex(p0, len(d_478_v7_))]
        d_479_v8_: _dafny.Array
        nw67_ = _dafny.Array(D2.default()(), 7)
        d_479_v8_ = nw67_
        d_480_v9_: _dafny.Map
        d_480_v9_ = _dafny.Map({p0: d_479_v8_})
        d_481_v10_: D6
        d_481_v10_ = D6_DC19(d_480_v9_)
        d_482_v11_: _dafny.Map
        d_482_v11_ = _dafny.Map({d_481_v10_: not((self).f10)})
        d_483_v12_: _dafny.Map
        d_483_v12_ = _dafny.Map({(self).f10: p0})
        pat_let_tv13_ = d_480_v9_
        def iife29_(_pat_let6_0):
            def iife30_(d_484_dt__update__tmp_h0_):
                def iife31_(_pat_let7_0):
                    def iife32_(d_485_dt__update_hcf31_h0_):
                        return D6_DC19(d_485_dt__update_hcf31_h0_)
                    return iife32_(_pat_let7_0)
                return iife31_(pat_let_tv13_)
            return iife30_(_pat_let6_0)
        d_482_v11_ = (d_482_v11_).set(iife29_(d_481_v10_), (len((d_483_v12_).set((self).f10, p0))) > (-462))
        d_486_v13_: bool
        d_486_v13_ = True
        d_487_v14_: _dafny.Array
        nw68_ = _dafny.Array(_dafny.Array(None, 0), 22)
        d_487_v14_ = nw68_
        index74_ = default__.safeIndex(487, (d_487_v14_).length(0))
        (d_487_v14_)[index74_] = d_476_v6_
        d_488_v15_: _dafny.Map
        d_488_v15_ = _dafny.Map({p0: (self).f10})
        d_489_v16_: _dafny.Map
        d_489_v16_ = _dafny.Map({(0) - (p0): ((d_488_v15_)[240] if (240) in (d_488_v15_) else False)})
        d_490_v17_: _dafny.Map
        d_490_v17_ = _dafny.Map({default__.fm0(p0, globalState): d_473_v3_})
        index75_ = default__.safeIndex(487, (d_487_v14_).length(0))
        rhs61_ = ((d_489_v16_)[len((d_490_v17_) | (d_490_v17_))] if (len((d_490_v17_) | (d_490_v17_))) in (d_489_v16_) else (self).f10)
        rhs62_ = (d_476_v6_ if d_486_v13_ else d_476_v6_)
        lhs48_ = d_487_v14_
        lhs49_ = default__.safeIndex(487, (d_487_v14_).length(0))
        d_486_v13_ = rhs61_
        lhs48_[lhs49_] = rhs62_
        d_491_v18_: _dafny.Set
        d_491_v18_ = _dafny.Set({p0})
        d_492_v19_: _dafny.Seq
        d_492_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmpiyma"))
        d_493_v20_: _dafny.Map
        d_493_v20_ = _dafny.Map({(d_492_v19_)[default__.safeIndex(p0, len(d_492_v19_))]: (self).f10})
        d_494_v21_: str
        d_494_v21_ = _dafny.CodePoint('t')
        d_495_v22_: _dafny.Map
        d_495_v22_ = _dafny.Map({(self).f10: d_494_v21_})
        d_486_v13_ = (d_491_v18_).isdisjoint((default__.fm35(d_472_v2_, _dafny.SeqWithoutIsStrInference([p0, len(d_493_v20_), 195]), d_486_v13_, globalState)) - (_dafny.Set({len(d_495_v22_), p0})))
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_476_v6_).length(0)):
            d_496_i1_: int = guard_loop_0_
            if (True) and (((0) <= (d_496_i1_)) and ((d_496_i1_) < ((d_476_v6_).length(0)))):
                (d_476_v6_)[(d_496_i1_)] = not ((not(d_486_v13_)) and (d_486_v13_)) or ((205) != (-815))

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Map = _dafny.Map({})
        d_497_i0_: int
        d_497_i0_ = 0
        with _dafny.label("2"):
            while not((self).f10):
                with _dafny.c_label("2"):
                    if (d_497_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_497_i0_ = (d_497_i0_) + (1)
                    d_498_v0_: _dafny.Array
                    nw69_ = _dafny.Array(_dafny.Map({}), 22)
                    d_498_v0_ = nw69_
                    d_499_v2_: str
                    d_499_v2_ = _dafny.CodePoint('m')
                    d_500_v3_: _dafny.Map
                    d_500_v3_ = _dafny.Map({p1: d_499_v2_})
                    d_501_v4_: _dafny.Seq
                    d_501_v4_ = _dafny.SeqWithoutIsStrInference([d_500_v3_, d_500_v3_, _dafny.Map({(self).f10: d_499_v2_})])
                    index76_ = default__.safeIndex(668, (d_498_v0_).length(0))
                    def iife33_():
                        coll17_ = _dafny.Map()
                        compr_17_: _dafny.Map
                        for compr_17_ in (d_501_v4_).Elements:
                            d_502_v1_: _dafny.Map = compr_17_
                            if (d_502_v1_) in (d_501_v4_):
                                coll17_[d_502_v1_] = not(not((self).f10))
                        return _dafny.Map(coll17_)
                    (d_498_v0_)[index76_] = iife33_()
                    
                    d_503_v5_: _dafny.Map
                    d_503_v5_ = _dafny.Map({d_500_v3_: (self).f10})
                    index77_ = default__.safeIndex(668, (d_498_v0_).length(0))
                    (d_498_v0_)[index77_] = (d_503_v5_)
                    d_504_v6_: int
                    d_504_v6_ = 585
                    r0 = (d_504_v6_) == (d_504_v6_)
                    d_505_v7_: D4
                    d_505_v7_ = D4_DC14(d_504_v6_, 699)
                    d_506_v8_: _dafny.Map
                    d_506_v8_ = _dafny.Map({d_504_v6_: d_505_v7_})
                    d_506_v8_ = (_dafny.Map({d_504_v6_: D4_DC14(d_504_v6_, (0) - (d_504_v6_))})) | (d_506_v8_)
                    d_507_v9_: _dafny.Map
                    d_507_v9_ = _dafny.Map({(self).f10: d_504_v6_})
                    if default__.fm0(((d_507_v9_)[p1] if (p1) in (d_507_v9_) else d_504_v6_), globalState):
                        d_508_v10_: _dafny.Array
                        nw70_ = _dafny.Array(False, 28)
                        d_508_v10_ = nw70_
                        index78_ = default__.safeIndex(30, (d_508_v10_).length(0))
                        (d_508_v10_)[index78_] = (self).fm8(d_504_v6_, (self).f10, p1, globalState)
                        d_509_v11_: _dafny.Seq
                        d_509_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojm"))
                        d_510_v12_: _dafny.Map
                        d_510_v12_ = _dafny.Map({(self).f10: (self).f10})
                        d_511_v13_: _dafny.MultiSet
                        d_511_v13_ = _dafny.MultiSet([d_507_v9_])
                        index79_ = default__.safeIndex(30, (d_508_v10_).length(0))
                        rhs63_ = (self).fm9(d_509_v11_, _dafny.Set({d_510_v12_}), globalState)
                        rhs64_ = (self).f10
                        rhs65_ = ((d_511_v13_).ispropersubset(d_511_v13_)) not in (d_507_v9_)
                        rhs66_ = default__.safeDivisionInt(d_504_v6_, d_504_v6_)
                        lhs50_ = globalState
                        lhs51_ = d_508_v10_
                        lhs52_ = default__.safeIndex(30, (d_508_v10_).length(0))
                        lhs53_ = globalState
                        lhs50_.f1 = rhs63_
                        r0 = rhs64_
                        lhs51_[lhs52_] = rhs65_
                        lhs53_.f1 = rhs66_
                        d_512_v14_: _dafny.Array
                        def lambda18_(d_513_v6_, d_514_v12_):
                            def lambda19_(d_515_i1_):
                                return default__.safeModuloInt(d_515_i1_, (D3_DC11(d_513_v6_, d_514_v12_)).cf17)

                            return lambda19_

                        init8_ = lambda18_(d_504_v6_, d_510_v12_)
                        nw71_ = _dafny.Array(None, 1)
                        for i0_8_ in range(nw71_.length(0)):
                            nw71_[i0_8_] = init8_(i0_8_)
                        d_512_v14_ = nw71_
                        d_516_v15_: _dafny.MultiSet
                        d_516_v15_ = _dafny.MultiSet([d_504_v6_])
                        index80_ = default__.safeIndex(378, (d_512_v14_).length(0))
                        (d_512_v14_)[index80_] = ((d_516_v15_)[d_504_v6_] if (d_504_v6_) in (d_516_v15_) else d_504_v6_)
                        index81_ = default__.safeIndex(378, (d_512_v14_).length(0))
                        (d_512_v14_)[index81_] = d_504_v6_
                        d_517_v16_: C0
                        nw72_ = C0()
                        nw72_.ctor__()
                        d_517_v16_ = nw72_
                        d_518_v17_: _dafny.Seq
                        d_518_v17_ = _dafny.SeqWithoutIsStrInference([(d_512_v14_)[default__.safeIndex(378, (d_512_v14_).length(0))]])
                        d_518_v17_ = d_518_v17_
                        d_519_v18_: _dafny.MultiSet
                        d_519_v18_ = _dafny.MultiSet([p1, (self).f10, False, (self).f10, (d_508_v10_)[default__.safeIndex(30, (d_508_v10_).length(0))]])
                        d_520_v19_: _dafny.Seq
                        d_520_v19_ = _dafny.SeqWithoutIsStrInference([(d_508_v10_)[default__.safeIndex(30, (d_508_v10_).length(0))]])
                        r0 = (_dafny.MultiSet([p1])).isdisjoint((d_519_v18_).set((d_520_v19_)[default__.safeIndex((0) - (len(d_518_v17_)), len(d_520_v19_))], default__.abs(-324)))
                    elif True:
                        d_521_v20_: C0
                        nw73_ = C0()
                        nw73_.ctor__()
                        d_521_v20_ = nw73_
                        r0 = (self).f10
                        d_522_v21_: _dafny.Map
                        d_522_v21_ = _dafny.Map({default__.fm0(d_504_v6_, globalState): p1})
                        d_523_v22_: _dafny.Set
                        d_523_v22_ = _dafny.Set({not(((d_522_v21_)[p1] if (p1) in (d_522_v21_) else p1)), p1, (self).f10, p1})
                        d_524_v23_: D10
                        d_524_v23_ = D10_DC30(_dafny.Set({(self).f10, (self).f10, (self).f10, p1}))
                        r0 = ((d_524_v23_).cf49).issubset((d_523_v22_) - (d_523_v22_))
                        d_525_v24_: _dafny.Seq
                        d_525_v24_ = _dafny.SeqWithoutIsStrInference([523])
                        d_526_v25_: _dafny.Seq
                        d_526_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
                        rhs67_ = ((d_504_v6_) * (d_504_v6_)) - (d_504_v6_)
                        rhs68_ = d_525_v24_
                        rhs69_ = (d_526_v25_) + (d_526_v25_)
                        lhs54_ = globalState
                        lhs54_.f1 = rhs67_
                        d_525_v24_ = rhs68_
                        r1 = rhs69_
                        r0 = (self).f10
                    pass
            pass
        d_527_v26_: str
        d_527_v26_ = _dafny.CodePoint('p')
        d_527_v26_ = d_527_v26_
        d_528_v27_: int
        d_528_v27_ = 380
        d_529_v28_: _dafny.Seq
        d_529_v28_ = _dafny.SeqWithoutIsStrInference([d_528_v27_])
        d_530_v29_: _dafny.Seq
        d_530_v29_ = _dafny.SeqWithoutIsStrInference([(D6_DC20((self).f10, p1, (self).f10, len(d_529_v28_))).cf33, p1])
        d_531_v30_: _dafny.Seq
        d_531_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obshd"))
        d_532_v31_: _dafny.Set
        d_532_v31_ = _dafny.Set({d_531_v30_})
        d_533_v32_: _dafny.MultiSet
        d_533_v32_ = _dafny.MultiSet([p1])
        d_534_v33_: D5
        d_534_v33_ = D5_DC17((self).fm8(d_528_v27_, (self).f10, (self).f10, globalState), d_527_v26_)
        d_535_v34_: _dafny.Map
        d_535_v34_ = _dafny.Map({(d_533_v32_).cardinality: (d_534_v33_).cf28})
        d_536_v35_: D6
        d_536_v35_ = D6_DC20(p1, p1, (self).f10, d_528_v27_)
        d_537_v36_: _dafny.Array
        nw74_ = _dafny.Array(None, 24)
        nw74_[int(0)] = (self).f10
        nw74_[int(1)] = (self).f10
        nw74_[int(2)] = (d_530_v29_)[default__.safeIndex(d_528_v27_, len(d_530_v29_))]
        nw74_[int(3)] = (d_528_v27_) == (728)
        nw74_[int(4)] = (self).f10
        nw74_[int(5)] = (self).f10
        nw74_[int(6)] = (d_528_v27_) != (d_528_v27_)
        nw74_[int(7)] = (self).f10
        nw74_[int(8)] = (d_532_v31_).issubset(d_532_v31_)
        nw74_[int(9)] = ((d_535_v34_)[d_528_v27_] if (d_528_v27_) in (d_535_v34_) else (self).f10)
        nw74_[int(10)] = (not(not((self).f10)) if (self).f10 else False)
        nw74_[int(11)] = (d_536_v35_).cf32
        nw74_[int(12)] = ((d_535_v34_)[(0) - (default__.fm1(globalState))] if ((0) - (default__.fm1(globalState))) in (d_535_v34_) else p1)
        nw74_[int(13)] = (self).f10
        nw74_[int(14)] = (self).f10
        nw74_[int(15)] = (d_528_v27_) < (d_528_v27_)
        nw74_[int(16)] = p1
        nw74_[int(17)] = (self).f10
        nw74_[int(18)] = p1
        nw74_[int(19)] = p1
        nw74_[int(20)] = (self).f10
        nw74_[int(21)] = (d_528_v27_) > (len(_dafny.Map({d_528_v27_: True})))
        nw74_[int(22)] = (self).f10
        nw74_[int(23)] = (self).f10
        d_537_v36_ = nw74_
        index82_ = default__.safeIndex(738, (d_537_v36_).length(0))
        (d_537_v36_)[index82_] = (not(not(p1)) if True else p1)
        d_538_v37_: _dafny.Array
        def lambda20_(d_539_v27_):
            def lambda21_(d_540_i2_):
                return default__.safeDivisionInt(d_540_i2_, d_539_v27_)

            return lambda21_

        init9_ = lambda20_(d_528_v27_)
        nw75_ = _dafny.Array(None, 28)
        for i0_9_ in range(nw75_.length(0)):
            nw75_[i0_9_] = init9_(i0_9_)
        d_538_v37_ = nw75_
        d_541_v38_: D1
        d_541_v38_ = D1_DC4(d_528_v27_, d_538_v37_)
        d_542_v39_: D2
        d_542_v39_ = D2_DC6(d_527_v26_)
        d_543_v40_: _dafny.Map
        d_543_v40_ = _dafny.Map({d_542_v39_: d_538_v37_})
        d_544_v41_: _dafny.Map
        d_544_v41_ = _dafny.Map({p1: d_538_v37_})
        pat_let_tv14_ = d_538_v37_
        d_545_v42_: _dafny.Array
        nw76_ = _dafny.Array(None, 29)
        nw76_[int(0)] = d_538_v37_
        nw76_[int(1)] = d_538_v37_
        nw76_[int(2)] = d_538_v37_
        nw76_[int(3)] = (d_538_v37_ if (self).f10 else d_538_v37_)
        nw76_[int(4)] = d_538_v37_
        nw76_[int(5)] = d_538_v37_
        nw76_[int(6)] = d_538_v37_
        nw76_[int(7)] = d_538_v37_
        nw76_[int(8)] = d_538_v37_
        nw76_[int(9)] = d_538_v37_
        nw76_[int(10)] = (d_541_v38_).cf6
        nw76_[int(11)] = d_538_v37_
        nw76_[int(12)] = d_538_v37_
        nw76_[int(13)] = d_538_v37_
        nw76_[int(14)] = d_538_v37_
        nw76_[int(15)] = d_538_v37_
        nw76_[int(16)] = d_538_v37_
        nw76_[int(17)] = (D1_DC4((0) - (d_528_v27_), d_538_v37_)).cf6
        nw76_[int(18)] = d_538_v37_
        nw76_[int(19)] = d_538_v37_
        nw76_[int(20)] = d_538_v37_
        def iife34_(_pat_let8_0):
            def iife35_(d_546_dt__update__tmp_h0_):
                def iife36_(_pat_let9_0):
                    def iife37_(d_547_dt__update_hcf6_h0_):
                        return D1_DC4((d_546_dt__update__tmp_h0_).cf5, d_547_dt__update_hcf6_h0_)
                    return iife37_(_pat_let9_0)
                return iife36_(pat_let_tv14_)
            return iife35_(_pat_let8_0)
        nw76_[int(21)] = (iife34_(d_541_v38_)).cf6
        nw76_[int(22)] = ((d_543_v40_)[d_542_v39_] if (d_542_v39_) in (d_543_v40_) else d_538_v37_)
        nw76_[int(23)] = d_538_v37_
        nw76_[int(24)] = ((d_544_v41_)[(self).f10] if ((self).f10) in (d_544_v41_) else d_538_v37_)
        nw76_[int(25)] = d_538_v37_
        nw76_[int(26)] = d_538_v37_
        nw76_[int(27)] = d_538_v37_
        nw76_[int(28)] = d_538_v37_
        d_545_v42_ = nw76_
        d_548_v43_: _dafny.Map
        d_548_v43_ = _dafny.Map({True: p1})
        index83_ = default__.safeIndex(738, (d_537_v36_).length(0))
        rhs70_ = (_dafny.SeqWithoutIsStrInference([(self).f10, p1])) < (d_530_v29_)
        rhs71_ = (d_545_v42_ if p1 else d_545_v42_)
        rhs72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xk"))
        rhs73_ = not (p1) or (((d_548_v43_)[p1] if (p1) in (d_548_v43_) else (self).f10))
        lhs55_ = d_537_v36_
        lhs56_ = default__.safeIndex(738, (d_537_v36_).length(0))
        lhs55_[lhs56_] = rhs70_
        d_545_v42_ = rhs71_
        r1 = rhs72_
        r0 = rhs73_
        index84_ = default__.safeIndex(738, (d_537_v36_).length(0))
        rhs74_ = ((True) == ((self).fm7((d_537_v36_)[default__.safeIndex(738, (d_537_v36_).length(0))], globalState)) if True else (d_537_v36_)[default__.safeIndex(738, (d_537_v36_).length(0))])
        rhs75_ = (0) - (d_528_v27_)
        rhs76_ = d_528_v27_
        lhs57_ = d_537_v36_
        lhs58_ = default__.safeIndex(738, (d_537_v36_).length(0))
        lhs59_ = globalState
        lhs57_[lhs58_] = rhs74_
        d_528_v27_ = rhs75_
        lhs59_.f7 = rhs76_
        d_549_v44_: _dafny.Seq
        d_549_v44_ = _dafny.SeqWithoutIsStrInference([d_531_v30_, default__.fm38((self).f10, d_528_v27_, globalState)])
        d_550_v45_: _dafny.Map
        d_550_v45_ = _dafny.Map({d_528_v27_: d_549_v44_})
        d_550_v45_ = (d_550_v45_).set(d_528_v27_, d_549_v44_)
        index85_ = default__.safeIndex(738, (d_537_v36_).length(0))
        (d_537_v36_)[index85_] = (self).f10
        r0 = ((self).f10) and (p1)
        r1 = d_531_v30_
        d_551_v46_: _dafny.Map
        d_551_v46_ = _dafny.Map({d_528_v27_: 715})
        r2 = (_dafny.Map({len(d_535_v34_): d_528_v27_})) | ((d_551_v46_) | (_dafny.Map({d_528_v27_: d_528_v27_})))
        return r0, r1, r2


class C2(T1):
    def  __init__(self):
        self._f21: bool = False
        self._f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f21, f22):
        (self)._f21 = f21
        (self)._f22 = f22

    def fm8(self, p0, p1, p2, globalState):
        return (self).f21

    def fm9(self, p0, p1, globalState):
        return (self).f22

    def fm27(self, p0, p1, globalState):
        return _dafny.Map({_dafny.CodePoint('n'): (self).f22})

    def m4(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: int = int(0)
        d_552_v0_: str
        d_552_v0_ = _dafny.CodePoint('i')
        d_553_v1_: _dafny.Map
        d_553_v1_ = _dafny.Map({d_552_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvocmt"))})
        d_554_v2_: _dafny.Seq
        d_554_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
        d_553_v1_ = ((d_553_v1_) | (d_553_v1_)) | (_dafny.Map({d_552_v0_: d_554_v2_}))
        if (self).f21:
            d_555_v3_: _dafny.Set
            d_555_v3_ = _dafny.Set({d_554_v2_, d_554_v2_, d_554_v2_, d_554_v2_})
            if ((d_555_v3_ if False else d_555_v3_)).issubset((d_555_v3_) | (default__.fm28((self).f21, (self).f21, globalState))):
                d_556_v4_: _dafny.Array
                def lambda22_(d_557_i0_):
                    return default__.safeModuloInt(d_557_i0_, 201)

                init10_ = lambda22_
                nw77_ = _dafny.Array(None, 22)
                for i0_10_ in range(nw77_.length(0)):
                    nw77_[i0_10_] = init10_(i0_10_)
                d_556_v4_ = nw77_
                d_556_v4_ = d_556_v4_
                d_558_v5_: _dafny.Map
                d_558_v5_ = _dafny.Map({(self).f22: (self).f21})
                d_558_v5_ = (d_558_v5_).set((self).f22, (self).f21)
                d_559_v6_: _dafny.Map
                d_559_v6_ = _dafny.Map({(self).f21: (self).f21})
                d_560_v7_: D3
                d_560_v7_ = D3_DC11((self).f22, (d_559_v6_) | (d_559_v6_))
                d_561_v8_: D6
                d_561_v8_ = D6_DC20((self).f21, (self).f21, (self).f21, (self).f22)
                rhs77_ = d_560_v7_
                rhs78_ = (d_561_v8_).cf33
                rhs79_ = (self).f22
                rhs80_ = (self).f22
                lhs60_ = globalState
                lhs61_ = globalState
                d_560_v7_ = rhs77_
                r1 = rhs78_
                lhs60_.f7 = rhs79_
                lhs61_.f0 = rhs80_
                (globalState).f7 = ((self).f22) * (((self).f22) * ((self).f22))
                d_552_v0_ = d_552_v0_
            elif True:
                d_562_v9_: D7
                d_562_v9_ = D7_DC23()
                d_562_v9_ = d_562_v9_
                d_563_v10_: _dafny.Array
                nw78_ = _dafny.Array(_dafny.Seq({}), 3)
                d_563_v10_ = nw78_
                d_564_v11_: _dafny.Seq
                d_564_v11_ = _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21])
                nw79_ = _dafny.Array(None, 1)
                nw79_[int(0)] = (d_564_v11_) + (_dafny.SeqWithoutIsStrInference([(self).f21]))
                d_563_v10_ = nw79_
                d_565_v12_: _dafny.Seq
                d_565_v12_ = _dafny.SeqWithoutIsStrInference([(self).f22, len(_dafny.SeqWithoutIsStrInference([(self).f21]))])
                d_566_v13_: _dafny.Array
                nw80_ = _dafny.Array(None, 8)
                nw80_[int(0)] = True
                nw80_[int(1)] = (self).f21
                nw80_[int(2)] = True
                nw80_[int(3)] = (self).f21
                nw80_[int(4)] = not((self).f21)
                nw80_[int(5)] = (self).f21
                nw80_[int(6)] = (self).f21
                nw80_[int(7)] = (self).f21
                d_566_v13_ = nw80_
                d_567_v14_: _dafny.Map
                d_567_v14_ = _dafny.Map({d_565_v12_: d_566_v13_})
                d_568_v15_: _dafny.Array
                out28_: _dafny.Array
                out28_ = (self).m15(d_567_v14_, (self).f21, False, d_552_v0_, globalState)
                d_568_v15_ = out28_
                d_569_v16_: _dafny.Map
                d_569_v16_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(self).f22 for d_570_i1_ in range(default__.abs(113))])) <= (d_565_v12_): default__.safeDivisionInt(670, default__.fm1(globalState))})
                d_569_v16_ = (d_569_v16_).set((self).f21, (self).f22)
                r1 = (self).f21
            (globalState).f7 = ((self).f22) - (len(d_554_v2_))
            d_571_v17_: _dafny.MultiSet
            d_571_v17_ = _dafny.MultiSet([False, (self).f21])
            d_572_v18_: _dafny.Seq
            d_572_v18_ = _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21, True, (self).f21])
            d_573_v19_: _dafny.Seq
            d_573_v19_ = _dafny.SeqWithoutIsStrInference([(self).f22, len(d_572_v18_)])
            d_574_v20_: _dafny.Seq
            d_574_v20_ = _dafny.SeqWithoutIsStrInference([d_573_v19_])
            d_575_v21_: D2
            d_575_v21_ = D2_DC7(((d_571_v17_)[(self).f21] if ((self).f21) in (d_571_v17_) else (self).f22), default__.fm0(len(d_574_v20_), globalState), -225)
            d_576_v22_: _dafny.Seq
            d_576_v22_ = _dafny.SeqWithoutIsStrInference([(d_575_v21_).cf10])
            d_576_v22_ = ((d_572_v18_) + (d_576_v22_)) + ((d_576_v22_) + (d_576_v22_))
            rhs81_ = (((d_571_v17_)[(self).f21] if ((self).f21) in (d_571_v17_) else (d_573_v19_)[default__.safeIndex(964, len(d_573_v19_))])) - ((0) - (((self).f22) + (49)))
            rhs82_ = d_571_v17_
            lhs62_ = globalState
            lhs62_.f0 = rhs81_
            d_571_v17_ = rhs82_
            d_577_v23_: _dafny.Map
            d_577_v23_ = _dafny.Map({(self).f22: (self).f22})
            d_578_v24_: _dafny.Map
            d_578_v24_ = _dafny.Map({d_577_v23_: (self).f22})
            source16_ = default__.fm29(len(d_578_v24_), (self).f21, _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21, (self).fm8((self).f22, (self).f21, (self).fm8((self).f22, (self).f21, (self).f21, globalState), globalState), (self).f21, True]), globalState)
            if source16_.is_DC10:
                d_579___mcc_h0_ = source16_.cf14
                d_580___mcc_h1_ = source16_.cf15
                d_581___mcc_h2_ = source16_.cf16
                d_582_cf16_ = d_581___mcc_h2_
                d_583_cf15_ = d_580___mcc_h1_
                d_584_cf14_ = d_579___mcc_h0_
                d_585_v25_: _dafny.Array
                nw81_ = _dafny.Array(None, 13)
                nw81_[int(0)] = d_552_v0_
                nw81_[int(1)] = d_552_v0_
                nw81_[int(2)] = d_552_v0_
                nw81_[int(3)] = d_552_v0_
                nw81_[int(4)] = d_552_v0_
                nw81_[int(5)] = d_552_v0_
                nw81_[int(6)] = d_552_v0_
                nw81_[int(7)] = default__.fm4((self).f21, d_583_cf15_, d_583_cf15_, globalState)
                nw81_[int(8)] = d_552_v0_
                nw81_[int(9)] = (d_554_v2_)[default__.safeIndex((self).f22, len(d_554_v2_))]
                nw81_[int(10)] = d_552_v0_
                nw81_[int(11)] = d_552_v0_
                nw81_[int(12)] = d_552_v0_
                d_585_v25_ = nw81_
                d_586_v26_: D5
                d_586_v26_ = D5_DC17((self).f21, d_552_v0_)
                index86_ = default__.safeIndex(670, (d_585_v25_).length(0))
                (d_585_v25_)[index86_] = (d_586_v26_).cf29
                d_587_v27_: D2
                d_587_v27_ = D2_DC6(_dafny.CodePoint('b'))
                index87_ = default__.safeIndex(670, (d_585_v25_).length(0))
                (d_585_v25_)[index87_] = (d_587_v27_).cf8
                r1 = True
                d_588_v28_: _dafny.Seq
                d_588_v28_ = _dafny.SeqWithoutIsStrInference([default__.fm30(globalState)])
                d_589_v29_: _dafny.Map
                d_589_v29_ = _dafny.Map({(self).f22: _dafny.Set({d_583_cf15_})})
                d_590_v30_: _dafny.Set
                d_590_v30_ = _dafny.Set({(self).f21, False, True, not((self).f21)})
                d_583_cf15_ = ((d_588_v28_).set(default__.safeIndex(len(((d_589_v29_)[220] if (220) in (d_589_v29_) else d_590_v30_)), len(d_588_v28_)), d_590_v30_)) == (d_588_v28_)
                d_591_v31_: _dafny.Array
                nw82_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_591_v31_ = nw82_
                d_592_v32_: _dafny.Array
                def lambda23_(d_593_i2_):
                    return (d_593_i2_) - ((self).f22)

                init11_ = lambda23_
                nw83_ = _dafny.Array(None, 4)
                for i0_11_ in range(nw83_.length(0)):
                    nw83_[i0_11_] = init11_(i0_11_)
                d_592_v32_ = nw83_
                d_594_v33_: _dafny.Seq
                d_594_v33_ = _dafny.SeqWithoutIsStrInference([d_592_v32_, d_592_v32_, d_592_v32_, d_592_v32_])
                index88_ = default__.safeIndex(349, (d_591_v31_).length(0))
                (d_591_v31_)[index88_] = (d_594_v33_)[default__.safeIndex((self).f22, len(d_594_v33_))]
                index89_ = default__.safeIndex(349, (d_591_v31_).length(0))
                rhs83_ = (d_554_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))
                rhs84_ = d_592_v32_
                lhs63_ = d_591_v31_
                lhs64_ = default__.safeIndex(349, (d_591_v31_).length(0))
                d_554_v2_ = rhs83_
                lhs63_[lhs64_] = rhs84_
            elif source16_.is_DC11:
                d_595___mcc_h3_ = source16_.cf17
                d_596___mcc_h4_ = source16_.cf18
                d_597_cf18_ = d_596___mcc_h4_
                d_598_cf17_ = d_595___mcc_h3_
                d_599_v34_: _dafny.Array
                nw84_ = _dafny.Array(False, 26)
                d_599_v34_ = nw84_
                index90_ = default__.safeIndex(753, (d_599_v34_).length(0))
                (d_599_v34_)[index90_] = not ((self).f21) or (default__.fm0(d_598_cf17_, globalState))
                index91_ = default__.safeIndex(753, (d_599_v34_).length(0))
                (d_599_v34_)[index91_] = (self).f21
                d_600_v35_: _dafny.Array
                nw85_ = _dafny.Array(int(0), 22)
                d_600_v35_ = nw85_
                d_601_v36_: _dafny.Set
                d_601_v36_ = _dafny.Set({not((self).f21), True})
                d_602_v37_: _dafny.Set
                d_602_v37_ = _dafny.Set({len(d_601_v36_), (0) - (len(d_573_v19_)), d_598_cf17_, 595})
                index92_ = default__.safeIndex(956, (d_600_v35_).length(0))
                (d_600_v35_)[index92_] = len((d_602_v37_).intersection(d_602_v37_))
                d_603_v38_: _dafny.Map
                d_603_v38_ = _dafny.Map({(0) - ((self).f22): (self).f21})
                index93_ = default__.safeIndex(956, (d_600_v35_).length(0))
                (d_600_v35_)[index93_] = (d_573_v19_)[default__.safeIndex(len((d_603_v38_) | (d_603_v38_)), len(d_573_v19_))]
                d_603_v38_ = (d_603_v38_).set(10, (self).f21)
                index94_ = default__.safeIndex(753, (d_599_v34_).length(0))
                (d_599_v34_)[index94_] = (d_599_v34_)[default__.safeIndex(753, (d_599_v34_).length(0))]
            elif source16_.is_DC9:
                d_604___mcc_h5_ = source16_.cf13
                d_605_cf13_ = d_604___mcc_h5_
                d_606_v39_: _dafny.Array
                nw86_ = _dafny.Array(int(0), 6)
                d_606_v39_ = nw86_
                d_607_v40_: _dafny.Set
                d_607_v40_ = _dafny.Set({d_606_v39_})
                d_554_v2_ = (d_554_v2_ if (d_607_v40_).ispropersubset(d_607_v40_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuqiag")))
                d_608_v41_: _dafny.Set
                d_608_v41_ = _dafny.Set({not((self).f21)})
                rhs85_ = False
                rhs86_ = ((False) or ((self).f21) if (self).f21 else (self).f21)
                rhs87_ = (self).f22
                rhs88_ = not((((self).f21) or ((self).f21) if (d_608_v41_).ispropersubset(d_608_v41_) else (self).f21))
                lhs65_ = globalState
                r1 = rhs85_
                r1 = rhs86_
                lhs65_.f7 = rhs87_
                r1 = rhs88_
                d_609_v42_: _dafny.Set
                d_609_v42_ = _dafny.Set({(self).f22})
                d_610_v43_: _dafny.Map
                d_610_v43_ = _dafny.Map({(self).f22: (self).f21})
                d_611_v44_: _dafny.Array
                nw87_ = _dafny.Array(None, 27)
                nw87_[int(0)] = (d_554_v2_) <= (d_554_v2_)
                nw87_[int(1)] = (self).f21
                nw87_[int(2)] = (self).f21
                nw87_[int(3)] = (True) and ((self).f21)
                nw87_[int(4)] = not((self).fm8((self).f22, (self).f21, (self).f21, globalState))
                nw87_[int(5)] = default__.fm0(len(d_609_v42_), globalState)
                nw87_[int(6)] = (self).f21
                nw87_[int(7)] = ((self).f22) in (d_573_v19_)
                nw87_[int(8)] = (self).f21
                nw87_[int(9)] = not((len(d_577_v23_)) > ((0) - ((self).f22)))
                nw87_[int(10)] = (self).f21
                nw87_[int(11)] = False
                nw87_[int(12)] = not((self).f21)
                nw87_[int(13)] = (self).f21
                nw87_[int(14)] = ((d_610_v43_)[(self).f22] if ((self).f22) in (d_610_v43_) else (self).f21)
                nw87_[int(15)] = (self).f21
                nw87_[int(16)] = (self).f21
                nw87_[int(17)] = (self).f21
                nw87_[int(18)] = (self).f21
                nw87_[int(19)] = (self).f21
                nw87_[int(20)] = not((self).f21)
                nw87_[int(21)] = (self).fm8((self).f22, (self).f21, (self).f21, globalState)
                nw87_[int(22)] = not((931) != ((self).f22))
                nw87_[int(23)] = (self).f21
                nw87_[int(24)] = (self).f21
                nw87_[int(25)] = (d_609_v42_).issubset(default__.fm31(82, globalState))
                nw87_[int(26)] = (self).f21
                d_611_v44_ = nw87_
                index95_ = default__.safeIndex(741, (d_611_v44_).length(0))
                (d_611_v44_)[index95_] = (self).f21
                d_612_v45_: _dafny.Map
                d_612_v45_ = _dafny.Map({(self).f21: (self).f21})
                index96_ = default__.safeIndex(741, (d_611_v44_).length(0))
                rhs89_ = (D3_DC11((self).f22, d_612_v45_)).cf17
                rhs90_ = not ((134) != (-995)) or ((self).f21)
                rhs91_ = (self).f21
                lhs66_ = globalState
                lhs67_ = d_611_v44_
                lhs68_ = default__.safeIndex(741, (d_611_v44_).length(0))
                lhs66_.f1 = rhs89_
                r1 = rhs90_
                lhs67_[lhs68_] = rhs91_
                index97_ = default__.safeIndex(827, (d_606_v39_).length(0))
                (d_606_v39_)[index97_] = ((self).f22) * ((self).f22)
                index98_ = default__.safeIndex(827, (d_606_v39_).length(0))
                (d_606_v39_)[index98_] = (0) - (-859)
            elif True:
                d_613___mcc_h6_ = source16_.cf19
                d_614_cf19_ = d_613___mcc_h6_
                d_615_v46_: _dafny.Set
                d_615_v46_ = _dafny.Set({(self).f21})
                (globalState).f3 = d_615_v46_
                (globalState).f7 = ((self).f22) + ((self).f22)
                r1 = (((self).f22) < ((self).f22)) in (d_572_v18_)
                d_616_v47_: _dafny.Array
                def lambda24_(d_617_i3_):
                    return True

                init12_ = lambda24_
                nw88_ = _dafny.Array(None, 25)
                for i0_12_ in range(nw88_.length(0)):
                    nw88_[i0_12_] = init12_(i0_12_)
                d_616_v47_ = nw88_
                d_618_v48_: _dafny.Set
                d_618_v48_ = _dafny.Set({(self).f22})
                index99_ = default__.safeIndex(153, (d_616_v47_).length(0))
                (d_616_v47_)[index99_] = (d_618_v48_).issubset(d_618_v48_)
                index100_ = default__.safeIndex(153, (d_616_v47_).length(0))
                (d_616_v47_)[index100_] = (self).f21
        elif True:
            d_619_v49_: _dafny.Seq
            d_619_v49_ = _dafny.SeqWithoutIsStrInference([((self).f21) and ((self).f21)])
            d_619_v49_ = _dafny.SeqWithoutIsStrInference([(self).f21, (False) == ((self).f21), (self).f21])
            rhs92_ = ((self).f22) >= (default__.safeDivisionInt((self).f22, (self).f22))
            rhs93_ = (self).f21
            r1 = rhs92_
            r1 = rhs93_
            d_620_v50_: _dafny.Array
            def lambda25_(d_621_v0_, d_622_v2_):
                def lambda26_(d_623_i4_):
                    return (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esg")), _dafny.SeqWithoutIsStrInference([d_621_v0_ for d_624_i5_ in range(default__.abs(537))]), d_622_v2_})).isdisjoint(_dafny.Set({d_622_v2_, _dafny.SeqWithoutIsStrInference([d_621_v0_ for d_625_i6_ in range(default__.abs(710))])}))

                return lambda26_

            init13_ = lambda25_(d_552_v0_, d_554_v2_)
            nw89_ = _dafny.Array(None, 18)
            for i0_13_ in range(nw89_.length(0)):
                nw89_[i0_13_] = init13_(i0_13_)
            d_620_v50_ = nw89_
            d_626_v51_: _dafny.MultiSet
            d_626_v51_ = _dafny.MultiSet([(self).f22, -467])
            index101_ = default__.safeIndex(344, (d_620_v50_).length(0))
            (d_620_v50_)[index101_] = not(default__.fm0((d_626_v51_).cardinality, globalState))
            d_627_v52_: D0
            d_627_v52_ = D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imswxpkga")))
            pat_let_tv15_ = d_554_v2_
            d_628_v53_: _dafny.Array
            nw90_ = _dafny.Array(None, 28)
            nw90_[int(0)] = d_627_v52_
            nw90_[int(1)] = D0_DC1(d_554_v2_)
            nw90_[int(2)] = default__.fm32((self).f21, (self).f22, False, globalState)
            nw90_[int(3)] = d_627_v52_
            nw90_[int(4)] = d_627_v52_
            nw90_[int(5)] = d_627_v52_
            nw90_[int(6)] = d_627_v52_
            nw90_[int(7)] = d_627_v52_
            nw90_[int(8)] = d_627_v52_
            nw90_[int(9)] = d_627_v52_
            nw90_[int(10)] = d_627_v52_
            nw90_[int(11)] = d_627_v52_
            nw90_[int(12)] = d_627_v52_
            nw90_[int(13)] = d_627_v52_
            nw90_[int(14)] = d_627_v52_
            nw90_[int(15)] = d_627_v52_
            nw90_[int(16)] = D0_DC1(_dafny.SeqWithoutIsStrInference([d_552_v0_ for d_629_i7_ in range(default__.abs(823))]))
            nw90_[int(17)] = D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuajlyoaj")))
            nw90_[int(18)] = d_627_v52_
            nw90_[int(19)] = d_627_v52_
            nw90_[int(20)] = d_627_v52_
            nw90_[int(21)] = D0_DC1(_dafny.SeqWithoutIsStrInference([d_552_v0_ for d_630_i8_ in range(default__.abs(6))]))
            nw90_[int(22)] = d_627_v52_
            nw90_[int(23)] = d_627_v52_
            def iife38_(_pat_let10_0):
                def iife39_(d_631_dt__update__tmp_h0_):
                    def iife40_(_pat_let11_0):
                        def iife41_(d_632_dt__update_hcf1_h0_):
                            return D0_DC1(d_632_dt__update_hcf1_h0_)
                        return iife41_(_pat_let11_0)
                    return iife40_(pat_let_tv15_)
                return iife39_(_pat_let10_0)
            nw90_[int(24)] = iife38_(d_627_v52_)
            nw90_[int(25)] = (d_627_v52_ if not((self).f21) else d_627_v52_)
            nw90_[int(26)] = D0_DC1(d_554_v2_)
            nw90_[int(27)] = D0_DC1(_dafny.SeqWithoutIsStrInference([d_552_v0_ for d_633_i9_ in range(default__.abs(139))]))
            d_628_v53_ = nw90_
            index102_ = default__.safeIndex(663, (d_628_v53_).length(0))
            (d_628_v53_)[index102_] = d_627_v52_
            index103_ = default__.safeIndex(344, (d_620_v50_).length(0))
            index104_ = default__.safeIndex(663, (d_628_v53_).length(0))
            rhs94_ = (self).f21
            rhs95_ = 180
            rhs96_ = d_627_v52_
            lhs69_ = d_620_v50_
            lhs70_ = default__.safeIndex(344, (d_620_v50_).length(0))
            lhs71_ = globalState
            lhs72_ = d_628_v53_
            lhs73_ = default__.safeIndex(663, (d_628_v53_).length(0))
            lhs69_[lhs70_] = rhs94_
            lhs71_.f1 = rhs95_
            lhs72_[lhs73_] = rhs96_
            d_634_v54_: _dafny.MultiSet
            d_634_v54_ = _dafny.MultiSet([d_552_v0_, d_552_v0_, d_552_v0_])
            d_635_v55_: _dafny.Set
            d_635_v55_ = _dafny.Set({(d_620_v50_)[default__.safeIndex(344, (d_620_v50_).length(0))]})
            d_634_v54_ = (d_634_v54_) | ((d_634_v54_).set(d_552_v0_, default__.abs(((d_626_v51_)[(self).f22] if ((self).f22) in (d_626_v51_) else len(d_635_v55_)))))
            d_636_v56_: _dafny.Map
            d_636_v56_ = _dafny.Map({(d_619_v49_)[default__.safeIndex((self).f22, len(d_619_v49_))]: (self).f21})
            d_637_v57_: _dafny.Seq
            d_637_v57_ = _dafny.SeqWithoutIsStrInference([((d_636_v56_).set((d_620_v50_)[default__.safeIndex(344, (d_620_v50_).length(0))], (d_620_v50_)[default__.safeIndex(344, (d_620_v50_).length(0))])) | (default__.fm33((d_620_v50_)[default__.safeIndex(344, (d_620_v50_).length(0))], (self).f21, _dafny.SeqWithoutIsStrInference([d_552_v0_ for d_638_i10_ in range(default__.abs(-744))]), (d_620_v50_)[default__.safeIndex(344, (d_620_v50_).length(0))], globalState))])
            (globalState).f0 = len(d_637_v57_)
        source17_ = default__.fm34((self).f22, globalState)
        if source17_.is_DC7:
            d_639___mcc_h7_ = source17_.cf9
            d_640___mcc_h8_ = source17_.cf10
            d_641___mcc_h9_ = source17_.cf11
            d_642_cf11_ = d_641___mcc_h9_
            d_643_cf10_ = d_640___mcc_h8_
            d_644_cf9_ = d_639___mcc_h7_
            d_554_v2_ = d_554_v2_
            d_645_v59_: _dafny.Seq
            d_645_v59_ = _dafny.SeqWithoutIsStrInference([d_643_cf10_])
            d_646_v60_: _dafny.Set
            d_646_v60_ = _dafny.Set({(self).f22, 4})
            d_647_v61_: _dafny.Map
            d_647_v61_ = _dafny.Map({d_645_v59_: len(d_646_v60_)})
            d_648_v62_: D4
            def iife42_():
                coll18_ = _dafny.Map()
                compr_18_: _dafny.Seq
                for compr_18_ in (d_647_v61_).keys.Elements:
                    d_649_v58_: _dafny.Seq = compr_18_
                    if (d_649_v58_) in (d_647_v61_):
                        coll18_[d_649_v58_] = d_643_cf10_
                return _dafny.Map(coll18_)
            d_648_v62_ = D4_DC13((iife42_()
) | (_dafny.Map({d_645_v59_: d_643_cf10_})))
            source18_ = d_648_v62_
            if source18_.is_DC14:
                d_650___mcc_h12_ = source18_.cf21
                d_651___mcc_h13_ = source18_.cf22
                d_652_cf22_ = d_651___mcc_h13_
                d_653_cf21_ = d_650___mcc_h12_
                d_654_v63_: _dafny.Array
                nw91_ = _dafny.Array(None, 7)
                nw91_[int(0)] = d_643_cf10_
                nw91_[int(1)] = (self).f21
                nw91_[int(2)] = (self).f21
                nw91_[int(3)] = (self).f21
                nw91_[int(4)] = d_643_cf10_
                nw91_[int(5)] = (self).f21
                nw91_[int(6)] = (self).f21
                d_654_v63_ = nw91_
                d_655_v64_: _dafny.MultiSet
                d_655_v64_ = _dafny.MultiSet([d_654_v63_])
                d_656_v65_: _dafny.Set
                d_656_v65_ = _dafny.Set({(self).f21, not(True)})
                rhs97_ = (_dafny.MultiSet([d_654_v63_]) if (self).f21 else d_655_v64_)
                rhs98_ = (d_656_v65_).ispropersubset(d_656_v65_)
                d_655_v64_ = rhs97_
                r1 = rhs98_
                d_657_v66_: _dafny.Map
                d_657_v66_ = _dafny.Map({d_642_cf11_: (d_645_v59_) + (d_645_v59_)})
                d_657_v66_ = (d_657_v66_).set(d_644_cf9_, ((_dafny.SeqWithoutIsStrInference([d_643_cf10_])) + (_dafny.SeqWithoutIsStrInference([(self).f21]))).set(default__.safeIndex(len(_dafny.Set({d_643_cf10_, (self).f21})), len((_dafny.SeqWithoutIsStrInference([d_643_cf10_])) + (_dafny.SeqWithoutIsStrInference([(self).f21])))), d_643_cf10_))
                d_658_v67_: _dafny.Array
                nw92_ = _dafny.Array(int(0), 16)
                d_658_v67_ = nw92_
                index105_ = default__.safeIndex(499, (d_658_v67_).length(0))
                (d_658_v67_)[index105_] = d_652_cf22_
                d_659_v68_: _dafny.Map
                d_659_v68_ = _dafny.Map({(self).f21: d_644_cf9_})
                index106_ = default__.safeIndex(499, (d_658_v67_).length(0))
                (d_658_v67_)[index106_] = (len(d_659_v68_) if d_643_cf10_ else d_644_cf9_)
                r1 = (self).fm8((0) - (d_642_cf11_), d_643_cf10_, (self).f21, globalState)
            elif source18_.is_DC15:
                d_660___mcc_h14_ = source18_.cf23
                d_661___mcc_h15_ = source18_.cf24
                d_662___mcc_h16_ = source18_.cf25
                d_663___mcc_h17_ = source18_.cf26
                d_664_cf26_ = d_663___mcc_h17_
                d_665_cf25_ = d_662___mcc_h16_
                d_666_cf24_ = d_661___mcc_h15_
                d_667_cf23_ = d_660___mcc_h14_
                d_668_v69_: _dafny.Array
                nw93_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_668_v69_ = nw93_
                d_668_v69_ = (D8_DC25(d_668_v69_)).cf42
                d_554_v2_ = _dafny.SeqWithoutIsStrInference([(d_554_v2_)[default__.safeIndex((D4_DC14(d_642_cf11_, d_644_cf9_)).cf22, len(d_554_v2_))] for d_669_i11_ in range(default__.abs(-601))])
                r1 = not((d_643_cf10_) and ((self).f21))
                rhs99_ = True
                rhs100_ = d_643_cf10_
                r1 = rhs99_
                r1 = rhs100_
            elif True:
                d_670___mcc_h18_ = source18_.cf20
                d_671_cf20_ = d_670___mcc_h18_
                d_672_v70_: _dafny.Set
                d_672_v70_ = _dafny.Set({d_643_cf10_})
                (globalState).f3 = (d_672_v70_) - (d_672_v70_)
                d_673_v71_: _dafny.Array
                def lambda27_(d_674_v2_, d_675_v0_):
                    def lambda28_(d_676_i12_):
                        return (d_674_v2_) != (_dafny.SeqWithoutIsStrInference([d_675_v0_ for d_677_i13_ in range(default__.abs(793))]))

                    return lambda28_

                init14_ = lambda27_(d_554_v2_, d_552_v0_)
                nw94_ = _dafny.Array(None, 9)
                for i0_14_ in range(nw94_.length(0)):
                    nw94_[i0_14_] = init14_(i0_14_)
                d_673_v71_ = nw94_
                d_673_v71_ = d_673_v71_
                r1 = not((d_554_v2_) < (d_554_v2_))
                d_673_v71_ = d_673_v71_
            d_678_v72_: _dafny.Map
            d_678_v72_ = _dafny.Map({not(d_643_cf10_): d_644_cf9_})
            r1 = (d_643_cf10_ if (d_646_v60_).issubset(d_646_v60_) else (d_642_cf11_) > (((d_678_v72_)[(self).f21] if ((self).f21) in (d_678_v72_) else (self).f22)))
            d_679_v73_: C1
            nw95_ = C1()
            nw95_.ctor__((self).f21)
            d_679_v73_ = nw95_
        elif source17_.is_DC6:
            d_680___mcc_h10_ = source17_.cf8
            d_681_cf8_ = d_680___mcc_h10_
            if (self).f21:
                r1 = (self).f21
                d_682_v74_: _dafny.Seq
                d_682_v74_ = _dafny.SeqWithoutIsStrInference([True, (self).f21])
                d_683_v75_: _dafny.MultiSet
                d_683_v75_ = _dafny.MultiSet([d_682_v74_])
                d_684_v76_: _dafny.Map
                d_684_v76_ = _dafny.Map({d_683_v75_: (self).f22})
                d_685_v77_: _dafny.Map
                d_685_v77_ = _dafny.Map({(self).f22: (self).f21})
                d_686_v78_: _dafny.Map
                d_686_v78_ = _dafny.Map({((d_685_v77_)[(self).f22] if ((self).f22) in (d_685_v77_) else (self).f21): 61})
                d_687_v79_: _dafny.Map
                d_687_v79_ = _dafny.Map({(self).f22: d_686_v78_})
                (globalState).f0 = ((d_684_v76_)[(d_683_v75_) - (_dafny.MultiSet([d_682_v74_]))] if ((d_683_v75_) - (_dafny.MultiSet([d_682_v74_]))) in (d_684_v76_) else len(d_687_v79_))
                (globalState).f1 = (self).f22
                d_688_v80_: _dafny.Array
                nw96_ = _dafny.Array(int(0), 8)
                d_688_v80_ = nw96_
                d_689_v81_: _dafny.Map
                d_689_v81_ = _dafny.Map({d_688_v80_: d_554_v2_})
                d_690_v82_: _dafny.Map
                d_690_v82_ = _dafny.Map({(self).f21: (self).f21})
                d_691_v83_: _dafny.Set
                d_691_v83_ = _dafny.Set({(self).f21})
                d_689_v81_ = (_dafny.Map({d_688_v80_: d_554_v2_}) if ((d_690_v82_)[(self).f21] if ((self).f21) in (d_690_v82_) else (self).f21) else _dafny.Map({d_688_v80_: (((d_554_v2_).set(default__.safeIndex(len(d_691_v83_), len(d_554_v2_)), d_681_cf8_)).set(default__.safeIndex((self).f22, len((d_554_v2_).set(default__.safeIndex(len(d_691_v83_), len(d_554_v2_)), d_681_cf8_))), d_552_v0_)).set(default__.safeIndex((self).f22, len(((d_554_v2_).set(default__.safeIndex(len(d_691_v83_), len(d_554_v2_)), d_681_cf8_)).set(default__.safeIndex((self).f22, len((d_554_v2_).set(default__.safeIndex(len(d_691_v83_), len(d_554_v2_)), d_681_cf8_))), d_552_v0_))), d_681_cf8_)}))
                d_692_v84_: _dafny.MultiSet
                d_692_v84_ = _dafny.MultiSet([d_681_cf8_])
                d_693_v85_: D6
                d_693_v85_ = D6_DC21((self).f21, (self).f21, (self).f21, ((d_692_v84_)[d_681_cf8_] if (d_681_cf8_) in (d_692_v84_) else 918))
                r1 = (d_693_v85_).cf38
            elif True:
                d_694_v86_: _dafny.Array
                nw97_ = _dafny.Array(int(0), 24)
                d_694_v86_ = nw97_
                index107_ = default__.safeIndex(455, (d_694_v86_).length(0))
                (d_694_v86_)[index107_] = len((d_554_v2_) + (_dafny.SeqWithoutIsStrInference([d_681_cf8_ for d_695_i14_ in range(default__.abs(777))])))
                index108_ = default__.safeIndex(455, (d_694_v86_).length(0))
                (d_694_v86_)[index108_] = default__.safeModuloInt((self).f22, (self).f22)
                d_696_v87_: _dafny.Map
                d_696_v87_ = _dafny.Map({(self).f21: (self).f21})
                d_696_v87_ = (d_696_v87_) | ((d_696_v87_) | (_dafny.Map({(self).f21: False})))
                d_697_v89_: _dafny.Seq
                d_697_v89_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_552_v0_ for d_698_i15_ in range(default__.abs(-580))]), d_554_v2_])
                d_699_v90_: _dafny.Map
                def iife43_():
                    coll19_ = _dafny.Map()
                    compr_19_: _dafny.Seq
                    for compr_19_ in (d_697_v89_).Elements:
                        d_700_v88_: _dafny.Seq = compr_19_
                        if (d_700_v88_) in (d_697_v89_):
                            coll19_[d_700_v88_] = (self).f21
                    return _dafny.Map(coll19_)
                d_699_v90_ = _dafny.Map({iife43_()
                : ((self).f21 if (self).f21 else (self).f21)})
                d_701_v91_: _dafny.Map
                d_701_v91_ = _dafny.Map({d_554_v2_: (self).fm8((self).f22, (self).f21, (self).f21, globalState)})
                d_699_v90_ = (d_699_v90_).set(d_701_v91_, (self).f21)
                d_702_v92_: _dafny.Array
                def lambda29_(d_703_i16_):
                    return D2_DC6(_dafny.CodePoint('t'))

                init15_ = lambda29_
                nw98_ = _dafny.Array(None, 19)
                for i0_15_ in range(nw98_.length(0)):
                    nw98_[i0_15_] = init15_(i0_15_)
                d_702_v92_ = nw98_
                d_704_v93_: D2
                d_704_v93_ = D2_DC6(d_552_v0_)
                index109_ = default__.safeIndex(7, (d_702_v92_).length(0))
                (d_702_v92_)[index109_] = d_704_v93_
                index110_ = default__.safeIndex(7, (d_702_v92_).length(0))
                (d_702_v92_)[index110_] = default__.fm39(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_705_i17_ in range(default__.abs(71))])), _dafny.SeqWithoutIsStrInference([d_552_v0_ for d_706_i18_ in range(default__.abs(625))]), False, globalState)
                r1 = default__.fm0(((self).f22 if (self).f21 else (self).f22), globalState)
            d_707_v94_: _dafny.MultiSet
            d_707_v94_ = _dafny.MultiSet([(self).f21])
            d_708_v95_: D5
            d_708_v95_ = D5_DC16((d_707_v94_).set((self).f21, default__.abs((self).f22)))
            r1 = not (((d_708_v95_).cf27).isdisjoint(d_707_v94_)) or (not (True) or ((self).f21))
            r1 = default__.fm0(default__.safeDivisionInt((self).f22, (self).f22), globalState)
            d_709_v96_: _dafny.Map
            d_709_v96_ = _dafny.Map({(self).f21: (self).fm8((self).f22, (self).f21, (self).f21, globalState)})
            d_710_v97_: _dafny.Map
            d_710_v97_ = _dafny.Map({len(d_709_v96_): (self).f21})
            d_711_v98_: _dafny.Set
            d_711_v98_ = _dafny.Set({(self).f22, ((self).f22) * (len(_dafny.Map({(self).f21: (self).f22}))), len((d_710_v97_) | (d_710_v97_))})
            d_711_v98_ = default__.fm31((self).f22, globalState)
        elif True:
            d_712___mcc_h11_ = source17_.cf12
            d_713_cf12_ = d_712___mcc_h11_
            r1 = (self).f21
            d_714_v99_: _dafny.Array
            nw99_ = _dafny.Array(False, 23)
            d_714_v99_ = nw99_
            d_715_v100_: D0
            d_715_v100_ = D0_DC0(d_714_v99_)
            d_716_v101_: _dafny.Array
            nw100_ = _dafny.Array(int(0), 8)
            d_716_v101_ = nw100_
            d_717_v102_: _dafny.Array
            nw101_ = _dafny.Array(None, 12)
            nw101_[int(0)] = d_716_v101_
            nw101_[int(1)] = d_716_v101_
            nw101_[int(2)] = d_716_v101_
            nw101_[int(3)] = d_716_v101_
            nw101_[int(4)] = d_716_v101_
            nw101_[int(5)] = d_716_v101_
            nw101_[int(6)] = d_716_v101_
            nw101_[int(7)] = d_716_v101_
            nw101_[int(8)] = d_716_v101_
            nw101_[int(9)] = d_716_v101_
            nw101_[int(10)] = d_716_v101_
            nw101_[int(11)] = d_716_v101_
            d_717_v102_ = nw101_
            index111_ = default__.safeIndex(113, (d_717_v102_).length(0))
            (d_717_v102_)[index111_] = d_716_v101_
            d_718_v103_: _dafny.Map
            d_718_v103_ = _dafny.Map({d_552_v0_: ((self).f21) and ((self).f21)})
            d_719_v105_: _dafny.Seq
            d_719_v105_ = _dafny.SeqWithoutIsStrInference([False, (self).f21])
            d_720_v106_: _dafny.Map
            d_720_v106_ = _dafny.Map({d_719_v105_: (self).f21})
            d_721_v107_: D4
            def iife44_():
                coll20_ = _dafny.Map()
                compr_20_: _dafny.Seq
                for compr_20_ in (d_720_v106_).keys.Elements:
                    d_722_v104_: _dafny.Seq = compr_20_
                    if (d_722_v104_) in (d_720_v106_):
                        coll20_[d_722_v104_] = (self).f21
                return _dafny.Map(coll20_)
            d_721_v107_ = D4_DC13(iife44_()
)
            index112_ = default__.safeIndex(113, (d_717_v102_).length(0))
            rhs101_ = d_715_v100_
            rhs102_ = d_716_v101_
            rhs103_ = d_718_v103_
            rhs104_ = (d_554_v2_ if (self).fm8((self).f22, (self).f21, (self).f21, globalState) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgl")))
            rhs105_ = d_721_v107_
            lhs74_ = d_717_v102_
            lhs75_ = default__.safeIndex(113, (d_717_v102_).length(0))
            d_715_v100_ = rhs101_
            lhs74_[lhs75_] = rhs102_
            d_718_v103_ = rhs103_
            d_554_v2_ = rhs104_
            d_721_v107_ = rhs105_
            r1 = not(False)
            d_723_v108_: D8
            d_723_v108_ = D8_DC26()
            source19_ = d_723_v108_
            if source19_.is_DC26:
                d_724_v109_: _dafny.Map
                d_724_v109_ = _dafny.Map({True: (self).f21})
                d_725_v110_: _dafny.Map
                d_725_v110_ = _dafny.Map({len(d_724_v109_): (self).f21})
                rhs106_ = d_554_v2_
                rhs107_ = (d_719_v105_) + ((_dafny.SeqWithoutIsStrInference([(self).f21, (self).f21, (self).f21, (d_719_v105_)[default__.safeIndex(len(d_725_v110_), len(d_719_v105_))]])) + (default__.fm40((self).f21, False, globalState)))
                rhs108_ = (self).f22
                rhs109_ = default__.safeModuloInt((self).f22, default__.safeDivisionInt((self).f22, (self).f22))
                lhs76_ = globalState
                lhs77_ = globalState
                d_554_v2_ = rhs106_
                d_719_v105_ = rhs107_
                lhs76_.f0 = rhs108_
                lhs77_.f0 = rhs109_
                r1 = not((self).f21)
                d_726_v111_: D3
                d_726_v111_ = D3_DC11((self).f22, d_724_v109_)
                d_727_v112_: _dafny.Seq
                d_728_v113_: _dafny.Seq
                out29_: _dafny.Seq
                out30_: _dafny.Seq
                out29_, out30_ = default__.m0((d_726_v111_).cf17, (d_714_v99_ if not((self).f21) else d_714_v99_), globalState)
                d_727_v112_ = out29_
                d_728_v113_ = out30_
                d_729_v114_: _dafny.Array
                def lambda30_(d_730_i19_):
                    return _dafny.MultiSet([(self).f21])

                init16_ = lambda30_
                nw102_ = _dafny.Array(None, 17)
                for i0_16_ in range(nw102_.length(0)):
                    nw102_[i0_16_] = init16_(i0_16_)
                d_729_v114_ = nw102_
                d_731_v115_: _dafny.MultiSet
                d_731_v115_ = _dafny.MultiSet([(self).f21, (self).f21, (self).f21])
                d_732_v116_: _dafny.Seq
                d_732_v116_ = _dafny.SeqWithoutIsStrInference([d_731_v115_])
                d_733_v117_: D2
                d_733_v117_ = D2_DC7((self).f22, (self).f21, (0) - (default__.fm1(globalState)))
                d_734_v118_: _dafny.Map
                d_734_v118_ = _dafny.Map({(self).f21: _dafny.MultiSet(d_719_v105_)})
                nw103_ = _dafny.Array(None, 23)
                nw103_[int(0)] = d_731_v115_
                nw103_[int(1)] = _dafny.MultiSet(default__.fm40((self).f21, (self).f21, globalState))
                nw103_[int(2)] = _dafny.MultiSet([(self).f21, False])
                nw103_[int(3)] = (d_732_v116_)[default__.safeIndex((self).f22, len(d_732_v116_))]
                nw103_[int(4)] = (d_731_v115_) - (d_731_v115_)
                nw103_[int(5)] = _dafny.MultiSet((d_719_v105_) + (_dafny.SeqWithoutIsStrInference([(self).f21, (self).f21])))
                nw103_[int(6)] = (_dafny.MultiSet(d_719_v105_)).set((self).f21, default__.abs((self).f22))
                nw103_[int(7)] = (_dafny.MultiSet([True]) if (self).f21 else d_731_v115_)
                nw103_[int(8)] = d_731_v115_
                nw103_[int(9)] = (d_731_v115_) - (d_731_v115_)
                nw103_[int(10)] = default__.fm41(globalState)
                nw103_[int(11)] = d_731_v115_
                nw103_[int(12)] = (d_731_v115_).set((d_733_v117_).cf10, default__.abs((self).f22))
                nw103_[int(13)] = (((d_734_v118_)[(self).f21] if ((self).f21) in (d_734_v118_) else d_731_v115_)).intersection(_dafny.MultiSet(d_719_v105_))
                nw103_[int(14)] = d_731_v115_
                nw103_[int(15)] = (d_731_v115_).intersection(d_731_v115_)
                nw103_[int(16)] = _dafny.MultiSet((d_719_v105_) + (d_719_v105_))
                nw103_[int(17)] = d_731_v115_
                nw103_[int(18)] = _dafny.MultiSet([(self).f21])
                nw103_[int(19)] = d_731_v115_
                nw103_[int(20)] = d_731_v115_
                nw103_[int(21)] = d_731_v115_
                nw103_[int(22)] = d_731_v115_
                d_729_v114_ = nw103_
            elif source19_.is_DC27:
                d_735___mcc_h19_ = source19_.cf43
                d_736___mcc_h20_ = source19_.cf44
                d_737___mcc_h21_ = source19_.cf45
                d_738___mcc_h22_ = source19_.cf46
                d_739_cf46_ = d_738___mcc_h22_
                d_740_cf45_ = d_737___mcc_h21_
                d_741_cf44_ = d_736___mcc_h20_
                d_742_cf43_ = d_735___mcc_h19_
                index113_ = default__.safeIndex(493, (d_714_v99_).length(0))
                (d_714_v99_)[index113_] = d_742_cf43_
                index114_ = default__.safeIndex(493, (d_714_v99_).length(0))
                rhs110_ = d_742_cf43_
                rhs111_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nungwwqg")))
                lhs78_ = d_714_v99_
                lhs79_ = default__.safeIndex(493, (d_714_v99_).length(0))
                lhs80_ = globalState
                lhs78_[lhs79_] = rhs110_
                lhs80_.f0 = rhs111_
                d_742_cf43_ = (d_742_cf43_) and (False)
                arr0_ = (d_717_v102_)[default__.safeIndex(113, (d_717_v102_).length(0))]
                index115_ = default__.safeIndex(972, ((d_717_v102_)[default__.safeIndex(113, (d_717_v102_).length(0))]).length(0))
                arr0_[index115_] = default__.safeModuloInt(len(d_554_v2_), d_739_cf46_)
                arr1_ = (d_717_v102_)[default__.safeIndex(113, (d_717_v102_).length(0))]
                index116_ = default__.safeIndex(972, ((d_717_v102_)[default__.safeIndex(113, (d_717_v102_).length(0))]).length(0))
                arr1_[index116_] = default__.safeModuloInt(-332, d_739_cf46_)
                (globalState).f7 = (0) - ((self).f22)
            elif source19_.is_DC25:
                d_743___mcc_h23_ = source19_.cf42
                d_744_cf42_ = d_743___mcc_h23_
                d_745_v120_: _dafny.MultiSet
                d_745_v120_ = _dafny.MultiSet([(self).f22, (self).f22, (self).f22])
                d_746_v121_: _dafny.Seq
                d_746_v121_ = _dafny.SeqWithoutIsStrInference([((d_745_v120_)[(self).f22] if ((self).f22) in (d_745_v120_) else (self).f22), len(d_719_v105_)])
                d_747_v122_: _dafny.MultiSet
                d_747_v122_ = _dafny.MultiSet([(self).f22, len(d_746_v121_)])
                d_748_v123_: _dafny.Seq
                d_748_v123_ = _dafny.SeqWithoutIsStrInference([d_747_v122_, d_747_v122_])
                def iife45_():
                    coll21_ = _dafny.Map()
                    compr_21_: _dafny.MultiSet
                    for compr_21_ in (d_748_v123_).Elements:
                        d_749_v119_: _dafny.MultiSet = compr_21_
                        if (d_749_v119_) in (d_748_v123_):
                            coll21_[d_749_v119_] = d_746_v121_
                    return _dafny.Map(coll21_)
                (globalState).f0 = (((self).f22) * (len(iife45_()
                ))) + ((self).f22)
                d_750_v124_: _dafny.Seq
                d_750_v124_ = _dafny.SeqWithoutIsStrInference([d_746_v121_])
                d_751_v125_: _dafny.Array
                nw104_ = _dafny.Array(D1.default()(), 11)
                d_751_v125_ = nw104_
                d_752_v126_: _dafny.Seq
                d_752_v126_ = _dafny.SeqWithoutIsStrInference([d_751_v125_])
                d_753_v127_: D8
                d_753_v127_ = D8_DC27((self).f21, d_752_v126_, d_746_v121_, (self).f22)
                d_746_v121_ = ((d_750_v124_)[default__.safeIndex(len(d_554_v2_), len(d_750_v124_))]) + ((d_753_v127_).cf45)
                d_754_v128_: C1
                nw105_ = C1()
                nw105_.ctor__((self).f21)
                d_754_v128_ = nw105_
                index117_ = default__.safeIndex(695, (d_714_v99_).length(0))
                (d_714_v99_)[index117_] = (self).f21
                index118_ = default__.safeIndex(695, (d_714_v99_).length(0))
                rhs112_ = (self).f21
                rhs113_ = ((self).f22) - (-525)
                rhs114_ = (self).f21
                lhs81_ = globalState
                lhs82_ = d_714_v99_
                lhs83_ = default__.safeIndex(695, (d_714_v99_).length(0))
                r1 = rhs112_
                lhs81_.f0 = rhs113_
                lhs82_[lhs83_] = rhs114_
            elif True:
                d_755___mcc_h24_ = source19_.cf47
                d_756_cf47_ = d_755___mcc_h24_
                index119_ = default__.safeIndex(888, (d_716_v101_).length(0))
                (d_716_v101_)[index119_] = (self).f22
                index120_ = default__.safeIndex(888, (d_716_v101_).length(0))
                (d_716_v101_)[index120_] = (self).f22
                d_757_v129_: _dafny.Map
                d_757_v129_ = _dafny.Map({(d_716_v101_)[default__.safeIndex(888, (d_716_v101_).length(0))]: (self).fm8((d_716_v101_)[default__.safeIndex(888, (d_716_v101_).length(0))], (self).f21, (self).f21, globalState)})
                d_758_v130_: _dafny.Map
                d_758_v130_ = _dafny.Map({(d_757_v129_) | (d_757_v129_): (self).f22})
                pat_let_tv16_ = d_757_v129_
                pat_let_tv17_ = d_757_v129_
                index121_ = default__.safeIndex(888, (d_716_v101_).length(0))
                def iife46_(_pat_let12_0):
                    def iife47_(d_759_dt__update__tmp_h2_):
                        def iife48_(_pat_let13_0):
                            def iife49_(d_760_dt__update_hcf53_h1_):
                                return D11_DC32(d_760_dt__update_hcf53_h1_)
                            return iife49_(_pat_let13_0)
                        return iife48_(pat_let_tv16_)
                    return iife47_(_pat_let12_0)
                def iife50_(_pat_let14_0):
                    def iife51_(d_761_dt__update__tmp_h1_):
                        def iife52_(_pat_let15_0):
                            def iife53_(d_762_dt__update_hcf53_h0_):
                                return D11_DC32(d_762_dt__update_hcf53_h0_)
                            return iife53_(_pat_let15_0)
                        return iife52_(pat_let_tv17_)
                    return iife51_(_pat_let14_0)
                (d_716_v101_)[index121_] = ((d_758_v130_)[(iife46_(default__.fm42(globalState))).cf53] if ((iife50_(default__.fm42(globalState))).cf53) in (d_758_v130_) else ((d_716_v101_)[default__.safeIndex(888, (d_716_v101_).length(0))]) + ((self).f22))
                d_763_v131_: _dafny.Seq
                d_763_v131_ = _dafny.SeqWithoutIsStrInference([(self).f22])
                d_764_v132_: _dafny.MultiSet
                d_764_v132_ = _dafny.MultiSet([(self).f21])
                d_765_v133_: _dafny.MultiSet
                d_765_v133_ = _dafny.MultiSet([(d_764_v132_).cardinality, (self).f22, (self).f22, (d_716_v101_)[default__.safeIndex(888, (d_716_v101_).length(0))]])
                d_766_v134_: _dafny.Set
                d_766_v134_ = _dafny.Set({d_765_v133_})
                rhs115_ = d_714_v99_
                rhs116_ = d_552_v0_
                rhs117_ = not(not((self).f21))
                rhs118_ = (d_763_v131_) + (d_763_v131_)
                rhs119_ = d_766_v134_
                d_714_v99_ = rhs115_
                d_552_v0_ = rhs116_
                r1 = rhs117_
                d_763_v131_ = rhs118_
                d_766_v134_ = rhs119_
                d_767_v136_: _dafny.Array
                def lambda31_(d_768_v0_):
                    def lambda32_(d_769_i20_):
                        def iife54_():
                            coll22_ = _dafny.Map()
                            compr_22_: int
                            for compr_22_ in _dafny.IntegerRange(712, 334):
                                d_770_v135_: int = compr_22_
                                if ((712) <= (d_770_v135_)) and ((d_770_v135_) < (334)):
                                    coll22_[(d_770_v135_) - ((self).f22)] = d_768_v0_
                            return _dafny.Map(coll22_)
                        return (_dafny.Map({d_768_v0_: iife54_()
                        })) | (_dafny.Map({d_768_v0_: _dafny.Map({-560: d_768_v0_})}))

                    return lambda32_

                init17_ = lambda31_(d_552_v0_)
                nw106_ = _dafny.Array(None, 24)
                for i0_17_ in range(nw106_.length(0)):
                    nw106_[i0_17_] = init17_(i0_17_)
                d_767_v136_ = nw106_
                d_771_v137_: _dafny.Map
                d_771_v137_ = _dafny.Map({(d_716_v101_)[default__.safeIndex(888, (d_716_v101_).length(0))]: d_552_v0_})
                d_772_v138_: _dafny.Map
                d_772_v138_ = _dafny.Map({d_552_v0_: d_771_v137_})
                d_773_v139_: D5
                d_773_v139_ = D5_DC17((self).f21, d_552_v0_)
                index122_ = default__.safeIndex(493, (d_767_v136_).length(0))
                (d_767_v136_)[index122_] = (d_772_v138_) | (_dafny.Map({default__.fm4((self).f21, (self).f21, (d_773_v139_).cf28, globalState): d_771_v137_}))
                index123_ = default__.safeIndex(493, (d_767_v136_).length(0))
                (d_767_v136_)[index123_] = _dafny.Map({d_552_v0_: d_771_v137_})
        d_774_v140_: _dafny.Array
        nw107_ = _dafny.Array(False, 26)
        d_774_v140_ = nw107_
        index124_ = default__.safeIndex(691, (d_774_v140_).length(0))
        (d_774_v140_)[index124_] = (self).f21
        d_775_v141_: _dafny.Map
        d_775_v141_ = _dafny.Map({(self).f21: (self).f21})
        d_776_v142_: _dafny.MultiSet
        d_776_v142_ = _dafny.MultiSet([(self).f21])
        d_777_v143_: _dafny.Map
        d_777_v143_ = _dafny.Map({(_dafny.MultiSet([((d_775_v141_)[(self).f21] if ((self).f21) in (d_775_v141_) else (self).f21)])).intersection(d_776_v142_): d_554_v2_})
        d_778_v144_: _dafny.Map
        d_778_v144_ = _dafny.Map({(self).f22: (self).f21})
        d_779_v145_: _dafny.MultiSet
        d_779_v145_ = _dafny.MultiSet([d_778_v144_])
        d_780_v146_: _dafny.Map
        d_780_v146_ = _dafny.Map({(self).f21: (self).f22})
        d_781_v147_: _dafny.Map
        d_781_v147_ = _dafny.Map({d_780_v146_: (self).f22})
        d_782_v148_: D6
        d_782_v148_ = D6_DC20((self).f21, (self).f21, (self).f21, (self).f22)
        d_783_v149_: _dafny.Map
        d_783_v149_ = _dafny.Map({d_782_v148_: (self).f21})
        d_784_v150_: _dafny.Seq
        d_784_v150_ = _dafny.SeqWithoutIsStrInference([True])
        d_785_v152_: _dafny.Map
        def iife55_():
            coll23_ = _dafny.Map()
            compr_23_: _dafny.MultiSet
            for compr_23_ in (_dafny.SeqWithoutIsStrInference([d_776_v142_ for d_786_i21_ in range(default__.abs(692))])).Elements:
                d_787_v151_: _dafny.MultiSet = compr_23_
                if (d_787_v151_) in (_dafny.SeqWithoutIsStrInference([d_776_v142_ for d_786_i21_ in range(default__.abs(692))])):
                    coll23_[d_787_v151_] = d_554_v2_
            return _dafny.Map(coll23_)
        d_785_v152_ = _dafny.Map({(self).f22: (d_777_v143_) | (iife55_()
        )})
        index125_ = default__.safeIndex(691, (d_774_v140_).length(0))
        rhs120_ = (self).fm8((((d_779_v145_)[d_778_v144_] if (d_778_v144_) in (d_779_v145_) else (0) - (((d_781_v147_)[d_780_v146_] if (d_780_v146_) in (d_781_v147_) else -78))) if default__.fm0((self).f22, globalState) else (self).f22), ((self).f21 if ((d_783_v149_)[d_782_v148_] if (d_782_v148_) in (d_783_v149_) else (self).f21) else (self).f21), ((self).f21 if (self).f21 else (self).f21), globalState)
        rhs121_ = ((D11_DC34((self).f22, (self).f21, len(d_784_v150_))).cf58) > ((self).f22)
        rhs122_ = d_552_v0_
        rhs123_ = ((d_785_v152_)[(self).f22] if ((self).f22) in (d_785_v152_) else d_777_v143_)
        lhs84_ = d_774_v140_
        lhs85_ = default__.safeIndex(691, (d_774_v140_).length(0))
        lhs84_[lhs85_] = rhs120_
        r1 = rhs121_
        d_552_v0_ = rhs122_
        d_777_v143_ = rhs123_
        d_788_v153_: D0
        d_788_v153_ = D0_DC0(d_774_v140_)
        d_789_v154_: _dafny.Map
        d_789_v154_ = _dafny.Map({d_788_v153_: (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))]})
        r1 = default__.fm0(len(d_789_v154_), globalState)
        d_790_v155_: D7
        d_790_v155_ = D7_DC24((self).f22)
        source20_ = d_790_v155_
        if source20_.is_DC23:
            d_791_v157_: _dafny.Array
            def lambda33_(d_792_i22_):
                def iife56_():
                    coll24_ = _dafny.Map()
                    compr_24_: int
                    for compr_24_ in _dafny.IntegerRange(128, 543):
                        d_793_v156_: int = compr_24_
                        if ((128) <= (d_793_v156_)) and ((d_793_v156_) < (543)):
                            coll24_[(d_793_v156_) + ((self).f22)] = (self).f22
                    return _dafny.Map(coll24_)
                return D1_DC3(len(iife56_()
))

            init18_ = lambda33_
            nw108_ = _dafny.Array(None, 16)
            for i0_18_ in range(nw108_.length(0)):
                nw108_[i0_18_] = init18_(i0_18_)
            d_791_v157_ = nw108_
            d_794_v158_: D8
            d_794_v158_ = D8_DC28(D8_DC27((d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))], _dafny.SeqWithoutIsStrInference([d_791_v157_]), (_dafny.SeqWithoutIsStrInference([(self).f22 for d_795_i23_ in range(default__.abs(61))])).set(default__.safeIndex((self).f22, len(_dafny.SeqWithoutIsStrInference([(self).f22 for d_796_i23_ in range(default__.abs(61))]))), -2), 791))
            d_797_v159_: D8
            d_797_v159_ = D8_DC28(d_794_v158_)
            rhs124_ = d_797_v159_
            rhs125_ = (self).f22
            d_797_v159_ = rhs124_
            r2 = rhs125_
            d_778_v144_ = (d_778_v144_).set((0) - ((self).f22), (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))])
            d_798_v160_: _dafny.Map
            d_798_v160_ = _dafny.Map({(self).f21: _dafny.MultiSet([(self).f22])})
            if (len(d_798_v160_)) < ((self).f22):
                d_799_v161_: C1
                nw109_ = C1()
                nw109_.ctor__((self).f21)
                d_799_v161_ = nw109_
                d_800_v162_: _dafny.Map
                d_800_v162_ = _dafny.Map({d_799_v161_: default__.fm1(globalState)})
                d_801_v163_: _dafny.Seq
                d_801_v163_ = _dafny.SeqWithoutIsStrInference([len(d_554_v2_), len(d_800_v162_), (self).f22, (self).f22, (self).f22])
                d_802_v164_: _dafny.Map
                d_802_v164_ = _dafny.Map({d_801_v163_: d_774_v140_})
                d_803_v165_: _dafny.Array
                out31_: _dafny.Array
                out31_ = (self).m15(d_802_v164_, (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))], (self).f21, d_552_v0_, globalState)
                d_803_v165_ = out31_
                (globalState).f7 = (self).f22
                d_784_v150_ = d_784_v150_
                d_803_v165_ = d_803_v165_
                d_804_v166_: D5
                d_804_v166_ = D5_DC16(_dafny.MultiSet([(self).f21]))
                d_805_v167_: D5
                d_805_v167_ = D5_DC18(d_804_v166_)
                d_806_v168_: D5
                d_806_v168_ = D5_DC18(d_804_v166_)
                d_807_v169_: _dafny.Seq
                d_807_v169_ = _dafny.SeqWithoutIsStrInference([d_806_v168_])
                d_808_v170_: _dafny.Map
                d_808_v170_ = _dafny.Map({d_807_v169_: (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))]})
                d_809_v171_: _dafny.Map
                d_809_v171_ = _dafny.Map({(0) - ((self).f22): d_807_v169_})
                d_808_v170_ = (d_808_v170_).set(((d_809_v171_)[(self).f22] if ((self).f22) in (d_809_v171_) else default__.fm43((self).f22, (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))], (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))], globalState)), (d_784_v150_)[default__.safeIndex((self).f22, len(d_784_v150_))])
            elif True:
                d_810_v172_: D11
                d_810_v172_ = D11_DC34((self).f22, (self).f21, len(d_784_v150_))
                d_811_v173_: _dafny.Set
                d_811_v173_ = _dafny.Set({(self).f22, ((self).f22) - ((d_810_v172_).cf58)})
                d_811_v173_ = ((d_811_v173_) | (d_811_v173_)) - ((d_811_v173_) - (d_811_v173_))
                d_812_v174_: _dafny.Map
                d_812_v174_ = _dafny.Map({(self).f21: d_774_v140_})
                d_813_v175_: D0
                d_813_v175_ = D0_DC2(((self).f21) and ((self).f21), ((d_812_v174_)[(self).f21] if ((self).f21) in (d_812_v174_) else d_774_v140_))
                d_813_v175_ = D0_DC2(False, d_774_v140_)
                index126_ = default__.safeIndex(691, (d_774_v140_).length(0))
                (d_774_v140_)[index126_] = (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))]
                r1 = (self).f21
                d_814_v176_: C0
                nw110_ = C0()
                nw110_.ctor__()
                d_814_v176_ = nw110_
            d_815_v177_: C0
            nw111_ = C0()
            nw111_.ctor__()
            d_815_v177_ = nw111_
        elif source20_.is_DC24:
            d_816___mcc_h25_ = source20_.cf41
            d_817_cf41_ = d_816___mcc_h25_
            d_818_v178_: _dafny.Array
            nw112_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_818_v178_ = nw112_
            d_819_v179_: _dafny.Seq
            d_819_v179_ = _dafny.SeqWithoutIsStrInference([d_776_v142_])
            d_820_v180_: _dafny.Array
            nw113_ = _dafny.Array(None, 9)
            nw113_[int(0)] = (self).f22
            nw113_[int(1)] = (self).f22
            nw113_[int(2)] = d_817_cf41_
            nw113_[int(3)] = (self).f22
            nw113_[int(4)] = (self).f22
            nw113_[int(5)] = d_817_cf41_
            nw113_[int(6)] = len(_dafny.Set({default__.fm0(len(_dafny.SeqWithoutIsStrInference([(d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))]])), globalState), (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))], (self).f21, (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))], (self).f21}))
            nw113_[int(7)] = (0) - (d_817_cf41_)
            nw113_[int(8)] = len(d_819_v179_)
            d_820_v180_ = nw113_
            index127_ = default__.safeIndex(875, (d_818_v178_).length(0))
            (d_818_v178_)[index127_] = d_820_v180_
            index128_ = default__.safeIndex(875, (d_818_v178_).length(0))
            (d_818_v178_)[index128_] = d_820_v180_
            d_821_v181_: D3
            d_821_v181_ = D3_DC10(d_554_v2_, not((self).f21), (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))])
            d_822_v182_: _dafny.Set
            d_822_v182_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_552_v0_]), (d_821_v181_).cf14})
            r0 = ((d_822_v182_).intersection(_dafny.Set({d_554_v2_, (d_554_v2_).set(default__.safeIndex((self).f22, len(d_554_v2_)), d_552_v0_)}))) | (d_822_v182_)
            d_817_cf41_ = len(default__.fm3(len((d_554_v2_) + (d_554_v2_)), d_817_cf41_, globalState))
            index129_ = default__.safeIndex(691, (d_774_v140_).length(0))
            (d_774_v140_)[index129_] = (self).f21
        elif True:
            d_823___mcc_h26_ = source20_.cf40
            d_824_cf40_ = d_823___mcc_h26_
            index130_ = default__.safeIndex(691, (d_774_v140_).length(0))
            (d_774_v140_)[index130_] = not (((self).f22) <= (len(_dafny.SeqWithoutIsStrInference([d_774_v140_])))) or (not ((d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))]) or ((self).f21))
            d_825_v183_: _dafny.Set
            d_825_v183_ = _dafny.Set({(self).f22, (self).f22, (self).f22})
            index131_ = default__.safeIndex(691, (d_774_v140_).length(0))
            rhs126_ = (d_784_v150_)[default__.safeIndex(len(d_825_v183_), len(d_784_v150_))]
            rhs127_ = 169
            rhs128_ = (self).f21
            lhs86_ = globalState
            lhs87_ = d_774_v140_
            lhs88_ = default__.safeIndex(691, (d_774_v140_).length(0))
            r1 = rhs126_
            lhs86_.f0 = rhs127_
            lhs87_[lhs88_] = rhs128_
            d_826_v184_: _dafny.Seq
            d_827_v185_: _dafny.Seq
            out32_: _dafny.Seq
            out33_: _dafny.Seq
            out32_, out33_ = default__.m0((self).f22, d_774_v140_, globalState)
            d_826_v184_ = out32_
            d_827_v185_ = out33_
            d_828_v186_: C1
            nw114_ = C1()
            nw114_.ctor__(not((self).f21))
            d_828_v186_ = nw114_
        d_829_v187_: _dafny.Set
        d_829_v187_ = _dafny.Set({d_554_v2_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_830_i24_ in range(default__.abs(206))])})
        r0 = (_dafny.Set({d_554_v2_, d_554_v2_, d_554_v2_})) - (d_829_v187_)
        r1 = (d_774_v140_)[default__.safeIndex(691, (d_774_v140_).length(0))]
        r2 = (0) - (((self).f22) * ((-75) - ((self).f22)))
        return r0, r1, r2

    def m5(self, p0, globalState):
        d_831_v0_: _dafny.Seq
        d_831_v0_ = _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21, (self).f21])
        (globalState).f1 = (len(d_831_v0_)) - (default__.fm1(globalState))
        d_832_v1_: _dafny.Array
        nw115_ = _dafny.Array(False, 24)
        d_832_v1_ = nw115_
        d_832_v1_ = d_832_v1_
        d_833_v2_: bool
        d_833_v2_ = False
        d_834_v3_: _dafny.Array
        nw116_ = _dafny.Array(int(0), 11)
        d_834_v3_ = nw116_
        d_835_v4_: _dafny.Set
        d_835_v4_ = _dafny.Set({p0})
        d_836_v5_: _dafny.Map
        d_836_v5_ = _dafny.Map({len(d_835_v4_): (self).f22})
        index132_ = default__.safeIndex(341, (d_834_v3_).length(0))
        (d_834_v3_)[index132_] = ((d_836_v5_)[874] if (874) in (d_836_v5_) else p0)
        index133_ = default__.safeIndex(684, (d_834_v3_).length(0))
        (d_834_v3_)[index133_] = p0
        d_837_v6_: D10
        d_837_v6_ = D10_DC31((self).f21, d_833_v2_, (self).f22)
        d_838_v7_: _dafny.Seq
        d_838_v7_ = _dafny.SeqWithoutIsStrInference([d_837_v6_])
        d_839_v8_: _dafny.Seq
        d_839_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
        d_840_v9_: str
        d_840_v9_ = _dafny.CodePoint('p')
        d_841_v10_: _dafny.Map
        d_841_v10_ = _dafny.Map({(self).f21: d_833_v2_})
        d_842_v11_: _dafny.Set
        d_842_v11_ = _dafny.Set({d_841_v10_, d_841_v10_})
        d_843_v12_: D12
        d_843_v12_ = D12_DC35(_dafny.SeqWithoutIsStrInference([d_837_v6_ for d_844_i1_ in range(default__.abs(529))]))
        d_845_v13_: _dafny.Seq
        d_845_v13_ = _dafny.SeqWithoutIsStrInference([((d_843_v12_).cf59) + (d_838_v7_), (d_838_v7_) + ((_dafny.SeqWithoutIsStrInference([d_837_v6_ for d_846_i2_ in range(default__.abs(685))])).set(default__.safeIndex((self).f22, len(_dafny.SeqWithoutIsStrInference([d_837_v6_ for d_847_i2_ in range(default__.abs(685))]))), d_837_v6_))])
        d_848_v14_: _dafny.MultiSet
        d_848_v14_ = _dafny.MultiSet([(self).f22])
        index134_ = default__.safeIndex(341, (d_834_v3_).length(0))
        index135_ = default__.safeIndex(684, (d_834_v3_).length(0))
        rhs129_ = d_833_v2_
        rhs130_ = (self).f22
        rhs131_ = (self).fm9((((d_839_v8_).set(default__.safeIndex(465, len(d_839_v8_)), _dafny.CodePoint('m'))) + (d_839_v8_)).set(default__.safeIndex(len(default__.fm38(False, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_849_i0_ in range(default__.abs(-606))])), globalState)), len(((d_839_v8_).set(default__.safeIndex(465, len(d_839_v8_)), _dafny.CodePoint('m'))) + (d_839_v8_))), d_840_v9_), (d_842_v11_) - (_dafny.Set({d_841_v10_})), globalState)
        rhs132_ = (d_845_v13_)[default__.safeIndex(default__.safeModuloInt(default__.fm1(globalState), ((d_848_v14_)[(self).f22] if ((self).f22) in (d_848_v14_) else p0)), len(d_845_v13_))]
        lhs89_ = d_834_v3_
        lhs90_ = default__.safeIndex(341, (d_834_v3_).length(0))
        lhs91_ = d_834_v3_
        lhs92_ = default__.safeIndex(684, (d_834_v3_).length(0))
        d_833_v2_ = rhs129_
        lhs89_[lhs90_] = rhs130_
        lhs91_[lhs92_] = rhs131_
        d_838_v7_ = rhs132_
        d_850_v15_: _dafny.MultiSet
        d_850_v15_ = _dafny.MultiSet([d_839_v8_])
        rhs133_ = not((d_850_v15_).ispropersubset(d_850_v15_))
        rhs134_ = default__.fm4((d_833_v2_ if (self).f21 else (self).f21), (self).f21, (self).f21, globalState)
        d_833_v2_ = rhs133_
        d_840_v9_ = rhs134_
        d_832_v1_ = d_832_v1_
        d_851_v16_: _dafny.Map
        d_851_v16_ = _dafny.Map({p0: not(d_833_v2_)})
        d_852_v17_: _dafny.MultiSet
        d_852_v17_ = _dafny.MultiSet([d_851_v16_])
        (globalState).f1 = ((d_836_v5_)[(d_852_v17_).cardinality] if ((d_852_v17_).cardinality) in (d_836_v5_) else (d_834_v3_)[default__.safeIndex(341, (d_834_v3_).length(0))])

    def m15(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_853_v0_: _dafny.Seq
        d_853_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        hi2_ = (self).f22
        for d_854_i0_ in range(len(d_853_v0_), hi2_):
            d_855_v1_: _dafny.Seq
            d_855_v1_ = _dafny.SeqWithoutIsStrInference([(0) - (d_854_i0_)])
            d_856_v2_: _dafny.Set
            d_856_v2_ = _dafny.Set({p2})
            d_857_v3_: _dafny.Array
            nw117_ = _dafny.Array(False, 10)
            d_857_v3_ = nw117_
            d_858_v4_: _dafny.Seq
            d_859_v5_: _dafny.Seq
            out34_: _dafny.Seq
            out35_: _dafny.Seq
            out34_, out35_ = default__.m0(len((d_855_v1_) + (default__.fm3((self).f22, len(d_856_v2_), globalState))), d_857_v3_, globalState)
            d_858_v4_ = out34_
            d_859_v5_ = out35_
            index136_ = default__.safeIndex(603, (d_857_v3_).length(0))
            (d_857_v3_)[index136_] = p1
            index137_ = default__.safeIndex(603, (d_857_v3_).length(0))
            rhs135_ = (d_853_v0_)[default__.safeIndex((d_854_i0_) + (d_854_i0_), len(d_853_v0_))]
            rhs136_ = (self).f22
            lhs93_ = d_857_v3_
            lhs94_ = default__.safeIndex(603, (d_857_v3_).length(0))
            lhs95_ = globalState
            lhs93_[lhs94_] = rhs135_
            lhs95_.f7 = rhs136_
            d_860_v6_: _dafny.MultiSet
            d_860_v6_ = _dafny.MultiSet([d_854_i0_, (_dafny.MultiSet([p2, p2, (self).f21, p1])).cardinality, (0) - (d_854_i0_)])
            d_860_v6_ = (d_860_v6_).intersection((d_860_v6_) | (d_860_v6_))
            d_861_v7_: _dafny.Map
            d_861_v7_ = _dafny.Map({d_854_i0_: d_854_i0_})
            d_861_v7_ = d_861_v7_
        d_862_v8_: _dafny.Map
        d_862_v8_ = _dafny.Map({default__.fm40(p2, (self).f21, globalState): p1})
        d_863_v9_: D4
        d_863_v9_ = D4_DC13((d_862_v8_) | (d_862_v8_))
        d_863_v9_ = d_863_v9_
        d_864_v10_: _dafny.Array
        def lambda34_(d_865_p1_):
            def lambda35_(d_866_i1_):
                return d_865_p1_

            return lambda35_

        init19_ = lambda34_(p1)
        nw118_ = _dafny.Array(None, 23)
        for i0_19_ in range(nw118_.length(0)):
            nw118_[i0_19_] = init19_(i0_19_)
        d_864_v10_ = nw118_
        d_867_v11_: _dafny.Map
        d_867_v11_ = _dafny.Map({(self).f21: (self).f22})
        d_868_v12_: _dafny.Seq
        d_868_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
        d_869_v13_: _dafny.Seq
        d_869_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bv"))])
        d_870_v14_: _dafny.Set
        d_870_v14_ = _dafny.Set({d_868_v12_, (d_869_v13_)[default__.safeIndex((self).f22, len(d_869_v13_))]})
        d_871_v15_: _dafny.MultiSet
        d_871_v15_ = _dafny.MultiSet([(self).f22])
        d_872_v16_: _dafny.Array
        nw119_ = _dafny.Array(None, 11)
        nw119_[int(0)] = (self).f22
        nw119_[int(1)] = (self).f22
        nw119_[int(2)] = (self).f22
        nw119_[int(3)] = (self).f22
        nw119_[int(4)] = (806) - (((d_867_v11_)[p1] if (p1) in (d_867_v11_) else len(d_870_v14_)))
        nw119_[int(5)] = (d_871_v15_).cardinality
        nw119_[int(6)] = (self).f22
        nw119_[int(7)] = (self).f22
        nw119_[int(8)] = (self).f22
        nw119_[int(9)] = (self).f22
        nw119_[int(10)] = (self).f22
        d_872_v16_ = nw119_
        d_873_v17_: _dafny.Map
        d_873_v17_ = _dafny.Map({p2: False})
        d_874_v18_: _dafny.Set
        d_874_v18_ = _dafny.Set({(d_873_v17_).set(p2, (self).f21), d_873_v17_, (d_873_v17_).set(p2, p1), d_873_v17_})
        index138_ = default__.safeIndex(364, (d_872_v16_).length(0))
        (d_872_v16_)[index138_] = ((self).f22) + ((self).fm9(d_868_v12_, d_874_v18_, globalState))
        d_875_v19_: bool
        d_875_v19_ = False
        d_876_v20_: _dafny.MultiSet
        d_876_v20_ = _dafny.MultiSet([p3, default__.fm4(p2, (self).f21, (self).f21, globalState), p3])
        d_877_v21_: _dafny.Seq
        d_877_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_868_v12_), d_876_v20_, d_876_v20_])
        d_878_v22_: _dafny.MultiSet
        d_878_v22_ = _dafny.MultiSet([d_872_v16_, d_872_v16_, d_872_v16_])
        index139_ = default__.safeIndex(364, (d_872_v16_).length(0))
        rhs137_ = d_864_v10_
        rhs138_ = (default__.fm1(globalState)) * ((self).f22)
        rhs139_ = not(((d_877_v21_)[default__.safeIndex((self).f22, len(d_877_v21_))]).issubset(d_876_v20_))
        rhs140_ = (d_878_v22_).issubset(d_878_v22_)
        lhs96_ = d_872_v16_
        lhs97_ = default__.safeIndex(364, (d_872_v16_).length(0))
        d_864_v10_ = rhs137_
        lhs96_[lhs97_] = rhs138_
        d_875_v19_ = rhs139_
        d_875_v19_ = rhs140_
        d_879_v23_: _dafny.MultiSet
        d_879_v23_ = _dafny.MultiSet([d_864_v10_])
        d_880_v24_: D13
        d_880_v24_ = D13_DC38(d_879_v23_)
        d_881_v25_: _dafny.Map
        d_881_v25_ = _dafny.Map({d_875_v19_: d_879_v23_})
        d_882_v26_: _dafny.Array
        nw120_ = _dafny.Array(None, 22)
        nw120_[int(0)] = ((d_880_v24_).cf64) | (((d_881_v25_)[d_875_v19_] if (d_875_v19_) in (d_881_v25_) else d_879_v23_))
        nw120_[int(1)] = d_879_v23_
        nw120_[int(2)] = d_879_v23_
        nw120_[int(3)] = _dafny.MultiSet([d_864_v10_])
        nw120_[int(4)] = d_879_v23_
        nw120_[int(5)] = d_879_v23_
        nw120_[int(6)] = d_879_v23_
        nw120_[int(7)] = d_879_v23_
        nw120_[int(8)] = d_879_v23_
        nw120_[int(9)] = d_879_v23_
        nw120_[int(10)] = ((d_880_v24_).cf64).intersection((_dafny.MultiSet([d_864_v10_])).set(d_864_v10_, default__.abs((0) - ((self).f22))))
        nw120_[int(11)] = d_879_v23_
        nw120_[int(12)] = d_879_v23_
        nw120_[int(13)] = _dafny.MultiSet([d_864_v10_, d_864_v10_, d_864_v10_])
        nw120_[int(14)] = d_879_v23_
        nw120_[int(15)] = d_879_v23_
        nw120_[int(16)] = d_879_v23_
        nw120_[int(17)] = d_879_v23_
        nw120_[int(18)] = (d_879_v23_).set(d_864_v10_, default__.abs((d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))]))
        nw120_[int(19)] = (d_879_v23_) - (_dafny.MultiSet([d_864_v10_]))
        nw120_[int(20)] = d_879_v23_
        nw120_[int(21)] = d_879_v23_
        d_882_v26_ = nw120_
        index140_ = default__.safeIndex(245, (d_882_v26_).length(0))
        (d_882_v26_)[index140_] = _dafny.MultiSet([d_864_v10_])
        index141_ = default__.safeIndex(245, (d_882_v26_).length(0))
        (d_882_v26_)[index141_] = d_879_v23_
        d_883_v27_: _dafny.Seq
        d_883_v27_ = _dafny.SeqWithoutIsStrInference([(d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))], -928])
        d_884_v28_: _dafny.Seq
        d_884_v28_ = _dafny.SeqWithoutIsStrInference([len(d_873_v17_), len(d_883_v27_)])
        d_875_v19_ = (65) > ((d_884_v28_)[default__.safeIndex((d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))], len(d_884_v28_))])
        hi3_ = (d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))]
        for d_885_i2_ in range(default__.safeModuloInt(25, (d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))]), hi3_):
            d_886_v29_: _dafny.Map
            d_886_v29_ = _dafny.Map({d_885_i2_: d_870_v14_})
            rhs141_ = default__.fm40((self).f21, (d_870_v14_).issubset(((d_886_v29_)[(d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))]] if ((d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))]) in (d_886_v29_) else _dafny.Set({d_868_v12_, d_868_v12_}))), globalState)
            rhs142_ = not((p3) in ((d_868_v12_ if not(p1) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "co")))))
            d_853_v0_ = rhs141_
            d_875_v19_ = rhs142_
            d_887_v30_: D6
            d_887_v30_ = D6_DC21((self).f21, d_875_v19_, (self).f21, ((d_871_v15_)[262] if (262) in (d_871_v15_) else d_885_i2_))
            d_875_v19_ = ((self).f22) > (((d_887_v30_).cf39) + (len(_dafny.SeqWithoutIsStrInference([(self).f22 for d_888_i3_ in range(default__.abs(417))]))))
            d_889_v31_: D3
            d_889_v31_ = D3_DC10((d_868_v12_).set(default__.safeIndex((d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))], len(d_868_v12_)), p3), p2, True)
            d_890_v32_: D13
            d_890_v32_ = D13_DC39(((d_867_v11_)[not(True)] if (not(True)) in (d_867_v11_) else (self).f22), d_875_v19_)
            d_891_v33_: _dafny.Array
            nw121_ = _dafny.Array(None, 28)
            nw121_[int(0)] = d_868_v12_
            nw121_[int(1)] = d_868_v12_
            nw121_[int(2)] = (d_868_v12_) + (d_868_v12_)
            nw121_[int(3)] = (d_868_v12_) + (d_868_v12_)
            nw121_[int(4)] = d_868_v12_
            nw121_[int(5)] = d_868_v12_
            nw121_[int(6)] = (d_869_v13_)[default__.safeIndex((d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))], len(d_869_v13_))]
            nw121_[int(7)] = d_868_v12_
            nw121_[int(8)] = d_868_v12_
            nw121_[int(9)] = (default__.fm38((self).f21, len(d_868_v12_), globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gotocugl")))
            nw121_[int(10)] = d_868_v12_
            nw121_[int(11)] = d_868_v12_
            nw121_[int(12)] = (d_868_v12_) + (d_868_v12_)
            nw121_[int(13)] = d_868_v12_
            nw121_[int(14)] = d_868_v12_
            nw121_[int(15)] = d_868_v12_
            nw121_[int(16)] = (d_889_v31_).cf14
            nw121_[int(17)] = d_868_v12_
            nw121_[int(18)] = (_dafny.SeqWithoutIsStrInference([p3 for d_892_i4_ in range(default__.abs(77))])) + (_dafny.SeqWithoutIsStrInference([p3 for d_893_i5_ in range(default__.abs(106))]))
            nw121_[int(19)] = d_868_v12_
            nw121_[int(20)] = (d_868_v12_ if True else d_868_v12_)
            nw121_[int(21)] = d_868_v12_
            nw121_[int(22)] = d_868_v12_
            nw121_[int(23)] = d_868_v12_
            nw121_[int(24)] = (d_868_v12_) + ((d_869_v13_)[default__.safeIndex(302, len(d_869_v13_))])
            nw121_[int(25)] = d_868_v12_
            nw121_[int(26)] = ((d_868_v12_) + (d_868_v12_)).set(default__.safeIndex((d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))], len((d_868_v12_) + (d_868_v12_))), default__.fm4(p1, (self).f21, (self).f21, globalState))
            nw121_[int(27)] = (d_868_v12_).set(default__.safeIndex((d_890_v32_).cf65, len(d_868_v12_)), p3)
            d_891_v33_ = nw121_
            index142_ = default__.safeIndex(442, (d_891_v33_).length(0))
            (d_891_v33_)[index142_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mu"))
            index143_ = default__.safeIndex(442, (d_891_v33_).length(0))
            (d_891_v33_)[index143_] = ((d_868_v12_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ussllkr"))) if p1 else d_868_v12_)
            d_894_v34_: _dafny.Set
            d_894_v34_ = _dafny.Set({(self).f22})
            d_895_v35_: _dafny.Map
            d_895_v35_ = _dafny.Map({len(d_894_v34_): (self).f22})
            d_896_v36_: D11
            d_896_v36_ = D11_DC34((d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))], p1, ((d_895_v35_)[(d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))]] if ((d_872_v16_)[default__.safeIndex(364, (d_872_v16_).length(0))]) in (d_895_v35_) else len(d_894_v34_)))
            d_896_v36_ = d_896_v36_
        r0 = d_872_v16_
        return r0

    @property
    def f21(self):
        return self._f21
    @property
    def f22(self):
        return self._f22

class C3(T0, T1):
    def  __init__(self):
        self._f10: bool = False
        self._f20: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f20, f10):
        (self)._f20 = f20
        (self)._f10 = f10

    def fm6(self, p0, p1, p2, p3, globalState):
        return (self).f20

    def fm7(self, p0, globalState):
        return not((self).f10)

    def fm8(self, p0, p1, p2, globalState):
        return (self).f20

    def fm9(self, p0, p1, globalState):
        return len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmomwxle"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nd")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lltikvrj"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_897_i0_ in range(default__.abs(259))]))))

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Map = _dafny.Map({})
        d_898_v0_: _dafny.Set
        d_898_v0_ = _dafny.Set({False, True})
        d_899_v1_: str
        d_899_v1_ = _dafny.CodePoint('l')
        d_900_v2_: D5
        d_900_v2_ = D5_DC17((d_898_v0_).issubset(_dafny.Set({(self).f10})), d_899_v1_)
        source21_ = d_900_v2_
        if source21_.is_DC17:
            d_901___mcc_h0_ = source21_.cf28
            d_902___mcc_h1_ = source21_.cf29
            d_903_cf29_ = d_902___mcc_h1_
            d_904_cf28_ = d_901___mcc_h0_
            d_905_v3_: int
            d_905_v3_ = 525
            (globalState).f7 = default__.safeDivisionInt(d_905_v3_, d_905_v3_)
            d_906_v4_: _dafny.Array
            nw122_ = _dafny.Array(int(0), 5)
            d_906_v4_ = nw122_
            d_906_v4_ = d_906_v4_
            r0 = p1
            d_907_v5_: _dafny.Seq
            d_907_v5_ = _dafny.SeqWithoutIsStrInference([(self).f10])
            index144_ = default__.safeIndex(591, (d_906_v4_).length(0))
            (d_906_v4_)[index144_] = (0) - ((d_905_v3_ if (d_907_v5_)[default__.safeIndex(d_905_v3_, len(d_907_v5_))] else d_905_v3_))
            index145_ = default__.safeIndex(591, (d_906_v4_).length(0))
            (d_906_v4_)[index145_] = d_905_v3_
        elif source21_.is_DC16:
            d_908___mcc_h2_ = source21_.cf27
            d_909_cf27_ = d_908___mcc_h2_
            d_910_v6_: _dafny.Array
            nw123_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
            d_910_v6_ = nw123_
            d_911_v7_: _dafny.Seq
            d_911_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifnhqam"))
            index146_ = default__.safeIndex(331, (d_910_v6_).length(0))
            (d_910_v6_)[index146_] = d_911_v7_
            d_912_v8_: int
            d_912_v8_ = 967
            index147_ = default__.safeIndex(331, (d_910_v6_).length(0))
            rhs143_ = 892
            rhs144_ = (d_911_v7_).set(default__.safeIndex(d_912_v8_, len(d_911_v7_)), d_899_v1_)
            lhs98_ = globalState
            lhs99_ = d_910_v6_
            lhs100_ = default__.safeIndex(331, (d_910_v6_).length(0))
            lhs98_.f1 = rhs143_
            lhs99_[lhs100_] = rhs144_
            d_913_v9_: _dafny.Seq
            d_913_v9_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10, False])
            d_914_v10_: _dafny.Array
            nw124_ = _dafny.Array(None, 19)
            nw124_[int(0)] = d_913_v9_
            nw124_[int(1)] = d_913_v9_
            nw124_[int(2)] = _dafny.SeqWithoutIsStrInference([(self).f20, (self).f20, (self).f20, p1, True])
            nw124_[int(3)] = (d_913_v9_).set(default__.safeIndex(d_912_v8_, len(d_913_v9_)), (self).f10)
            nw124_[int(4)] = d_913_v9_
            nw124_[int(5)] = _dafny.SeqWithoutIsStrInference([p1])
            nw124_[int(6)] = d_913_v9_
            nw124_[int(7)] = d_913_v9_
            nw124_[int(8)] = (d_913_v9_) + (d_913_v9_)
            nw124_[int(9)] = d_913_v9_
            nw124_[int(10)] = d_913_v9_
            nw124_[int(11)] = d_913_v9_
            nw124_[int(12)] = d_913_v9_
            nw124_[int(13)] = (d_913_v9_) + (_dafny.SeqWithoutIsStrInference([(self).f10]))
            nw124_[int(14)] = d_913_v9_
            nw124_[int(15)] = d_913_v9_
            nw124_[int(16)] = d_913_v9_
            nw124_[int(17)] = d_913_v9_
            nw124_[int(18)] = (d_913_v9_) + (d_913_v9_)
            d_914_v10_ = nw124_
            d_914_v10_ = d_914_v10_
            d_915_v11_: _dafny.Array
            def lambda36_(d_916_i0_):
                return not((self).f10)

            init20_ = lambda36_
            nw125_ = _dafny.Array(None, 24)
            for i0_20_ in range(nw125_.length(0)):
                nw125_[i0_20_] = init20_(i0_20_)
            d_915_v11_ = nw125_
            index148_ = default__.safeIndex(528, (d_915_v11_).length(0))
            (d_915_v11_)[index148_] = (self).f20
            d_917_v12_: _dafny.Map
            d_917_v12_ = _dafny.Map({23: not((self).f20)})
            index149_ = default__.safeIndex(528, (d_915_v11_).length(0))
            (d_915_v11_)[index149_] = (self).fm6(d_912_v8_, default__.fm1(globalState), d_912_v8_, ((d_917_v12_)[len(_dafny.SeqWithoutIsStrInference([d_899_v1_ for d_918_i1_ in range(default__.abs(480))]))] if (len(_dafny.SeqWithoutIsStrInference([d_899_v1_ for d_919_i1_ in range(default__.abs(480))]))) in (d_917_v12_) else (self).f10), globalState)
            d_920_v13_: _dafny.MultiSet
            d_920_v13_ = _dafny.MultiSet([d_912_v8_])
            d_921_v14_: D4
            d_921_v14_ = D4_DC15(d_912_v8_, d_920_v13_, d_899_v1_, d_912_v8_)
            d_922_v15_: _dafny.Array
            def lambda37_(d_923_v8_):
                def lambda38_(d_924_i2_):
                    return default__.safeDivisionInt(d_924_i2_, (0) - (d_923_v8_))

                return lambda38_

            init21_ = lambda37_(d_912_v8_)
            nw126_ = _dafny.Array(None, 21)
            for i0_21_ in range(nw126_.length(0)):
                nw126_[i0_21_] = init21_(i0_21_)
            d_922_v15_ = nw126_
            index150_ = default__.safeIndex(723, (d_922_v15_).length(0))
            (d_922_v15_)[index150_] = d_912_v8_
            d_925_v16_: _dafny.Set
            d_925_v16_ = _dafny.Set({508})
            pat_let_tv18_ = d_920_v13_
            pat_let_tv19_ = d_920_v13_
            pat_let_tv20_ = d_912_v8_
            index151_ = default__.safeIndex(723, (d_922_v15_).length(0))
            def iife57_(_pat_let16_0):
                def iife58_(d_926_dt__update__tmp_h0_):
                    def iife59_(_pat_let17_0):
                        def iife60_(d_927_dt__update_hcf24_h0_):
                            def iife61_(_pat_let18_0):
                                def iife62_(d_928_dt__update_hcf23_h0_):
                                    return D4_DC15(d_928_dt__update_hcf23_h0_, d_927_dt__update_hcf24_h0_, (d_926_dt__update__tmp_h0_).cf25, (d_926_dt__update__tmp_h0_).cf26)
                                return iife62_(_pat_let18_0)
                            return iife61_(pat_let_tv20_)
                        return iife60_(_pat_let17_0)
                    return iife59_((pat_let_tv18_) | (pat_let_tv19_))
                return iife58_(_pat_let16_0)
            rhs145_ = len(d_925_v16_)
            rhs146_ = iife57_(D4_DC15(len(d_925_v16_), _dafny.MultiSet([d_912_v8_]), d_899_v1_, d_912_v8_))
            rhs147_ = len((d_910_v6_)[default__.safeIndex(331, (d_910_v6_).length(0))])
            rhs148_ = d_912_v8_
            lhs101_ = d_922_v15_
            lhs102_ = default__.safeIndex(723, (d_922_v15_).length(0))
            lhs103_ = globalState
            d_912_v8_ = rhs145_
            d_921_v14_ = rhs146_
            lhs101_[lhs102_] = rhs147_
            lhs103_.f1 = rhs148_
        elif True:
            d_929___mcc_h3_ = source21_.cf30
            d_930_cf30_ = d_929___mcc_h3_
            d_931_v17_: _dafny.Array
            nw127_ = _dafny.Array(int(0), 28)
            d_931_v17_ = nw127_
            d_932_v18_: int
            d_932_v18_ = 590
            index152_ = default__.safeIndex(464, (d_931_v17_).length(0))
            (d_931_v17_)[index152_] = default__.safeDivisionInt(d_932_v18_, d_932_v18_)
            index153_ = default__.safeIndex(464, (d_931_v17_).length(0))
            rhs149_ = (0) - (default__.safeModuloInt(d_932_v18_, (d_932_v18_) - (d_932_v18_)))
            rhs150_ = not ((self).f20) or ((self).f10)
            rhs151_ = not (False) or ((self).f20)
            rhs152_ = d_899_v1_
            rhs153_ = ((0) - ((0) - (d_932_v18_))) + (d_932_v18_)
            lhs104_ = d_931_v17_
            lhs105_ = default__.safeIndex(464, (d_931_v17_).length(0))
            lhs106_ = globalState
            lhs104_[lhs105_] = rhs149_
            r0 = rhs150_
            r0 = rhs151_
            d_899_v1_ = rhs152_
            lhs106_.f0 = rhs153_
            d_899_v1_ = d_899_v1_
            r0 = False
            d_933_v19_: _dafny.Map
            d_933_v19_ = _dafny.Map({p1: d_932_v18_})
            d_934_v20_: _dafny.MultiSet
            d_934_v20_ = _dafny.MultiSet([((d_933_v19_)[(self).f20] if ((self).f20) in (d_933_v19_) else d_932_v18_), d_932_v18_])
            r0 = ((self).f20 if ((d_934_v20_).cardinality) > ((0) - ((d_931_v17_)[default__.safeIndex(464, (d_931_v17_).length(0))])) else False)
        d_935_v21_: int
        d_935_v21_ = -136
        d_936_v22_: _dafny.Map
        d_936_v22_ = _dafny.Map({(d_935_v21_) <= (d_935_v21_): (699) - (default__.fm1(globalState))})
        d_937_v23_: _dafny.Seq
        d_937_v23_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f10: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjtarsvry")))}), d_936_v22_])
        d_938_v24_: _dafny.Seq
        d_938_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjgi"))
        d_936_v22_ = (d_936_v22_).set((self).f10, len((d_937_v23_)[default__.safeIndex(len(d_938_v24_), len(d_937_v23_))]))
        d_939_v25_: _dafny.Array
        nw128_ = _dafny.Array(_dafny.Array(None, 0), 29)
        d_939_v25_ = nw128_
        d_940_v26_: D6
        d_940_v26_ = D6_DC20((self).f20, False, (self).f20, (0) - (d_935_v21_))
        d_941_v27_: _dafny.Array
        nw129_ = _dafny.Array(None, 21)
        nw129_[int(0)] = p1
        nw129_[int(1)] = p1
        nw129_[int(2)] = p1
        nw129_[int(3)] = False
        nw129_[int(4)] = False
        nw129_[int(5)] = (self).f20
        nw129_[int(6)] = (self).f10
        nw129_[int(7)] = False
        nw129_[int(8)] = (d_940_v26_).cf32
        nw129_[int(9)] = True
        nw129_[int(10)] = (self).f20
        nw129_[int(11)] = (self).f20
        nw129_[int(12)] = (self).f20
        nw129_[int(13)] = (self).f20
        nw129_[int(14)] = (self).f20
        nw129_[int(15)] = (self).f10
        nw129_[int(16)] = (self).f10
        nw129_[int(17)] = (self).f20
        nw129_[int(18)] = True
        nw129_[int(19)] = (self).f10
        nw129_[int(20)] = True
        d_941_v27_ = nw129_
        index154_ = default__.safeIndex(810, (d_939_v25_).length(0))
        (d_939_v25_)[index154_] = d_941_v27_
        index155_ = default__.safeIndex(810, (d_939_v25_).length(0))
        (d_939_v25_)[index155_] = d_941_v27_
        r0 = not (p1) or (p1)
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, ((d_939_v25_)[default__.safeIndex(810, (d_939_v25_).length(0))]).length(0)):
            d_942_i3_: int = guard_loop_1_
            if (True) and (((0) <= (d_942_i3_)) and ((d_942_i3_) < (((d_939_v25_)[default__.safeIndex(810, (d_939_v25_).length(0))]).length(0)))):
                _ingredientsd_0.append(((d_939_v25_)[default__.safeIndex(810, (d_939_v25_).length(0))], int(d_942_i3_), (self).f20))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        d_943_v28_: _dafny.Array
        def lambda39_(d_944_v21_):
            def lambda40_(d_945_i4_):
                return default__.safeDivisionInt(d_945_i4_, d_944_v21_)

            return lambda40_

        init22_ = lambda39_(d_935_v21_)
        nw130_ = _dafny.Array(None, 5)
        for i0_22_ in range(nw130_.length(0)):
            nw130_[i0_22_] = init22_(i0_22_)
        d_943_v28_ = nw130_
        d_946_v29_: _dafny.Set
        d_946_v29_ = _dafny.Set({d_943_v28_, (d_943_v28_ if p1 else d_943_v28_)})
        d_946_v29_ = d_946_v29_
        r0 = not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "metlvgmn"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbiudyvt"))))
        d_947_v30_: _dafny.Seq
        d_947_v30_ = _dafny.SeqWithoutIsStrInference([d_938_v24_, d_938_v24_, d_938_v24_])
        r1 = (d_947_v30_)[default__.safeIndex(d_935_v21_, len(d_947_v30_))]
        d_948_v31_: _dafny.Map
        d_948_v31_ = _dafny.Map({d_935_v21_: d_935_v21_})
        d_949_v32_: D7
        d_949_v32_ = D7_DC22(d_948_v31_)
        r2 = (d_949_v32_).cf40
        return r0, r1, r2

    def m4(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: int = int(0)
        d_950_v0_: _dafny.Seq
        d_950_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vetxc"))
        d_951_v1_: int
        d_951_v1_ = 208
        d_952_v2_: str
        d_952_v2_ = _dafny.CodePoint('l')
        d_953_v3_: _dafny.MultiSet
        d_953_v3_ = _dafny.MultiSet([(self).f10, True])
        d_954_v4_: _dafny.Seq
        d_954_v4_ = _dafny.SeqWithoutIsStrInference([(self).f10])
        d_955_v5_: _dafny.Map
        d_955_v5_ = _dafny.Map({not((d_954_v4_)[default__.safeIndex(605, len(d_954_v4_))]): d_953_v3_})
        if (len(((d_950_v0_) + (d_950_v0_)).set(default__.safeIndex(d_951_v1_, len((d_950_v0_) + (d_950_v0_))), d_952_v2_))) >= (((d_953_v3_) | (((d_955_v5_)[(self).f20] if ((self).f20) in (d_955_v5_) else d_953_v3_))).cardinality):
            d_956_v6_: _dafny.Map
            d_956_v6_ = _dafny.Map({False: d_951_v1_})
            d_951_v1_ = ((d_956_v6_)[(self).f10] if ((self).f10) in (d_956_v6_) else len((d_950_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qssp")))))
            r1 = (self).fm6(d_951_v1_, d_951_v1_, (d_951_v1_) * (d_951_v1_), (self).f20, globalState)
            d_957_v7_: _dafny.Array
            def lambda41_(d_958_i0_):
                return (self).f10

            init23_ = lambda41_
            nw131_ = _dafny.Array(None, 26)
            for i0_23_ in range(nw131_.length(0)):
                nw131_[i0_23_] = init23_(i0_23_)
            d_957_v7_ = nw131_
            d_957_v7_ = d_957_v7_
            if not((self).f10):
                d_959_v9_: _dafny.Array
                def lambda42_(d_960_v1_):
                    def lambda43_(d_961_i1_):
                        def iife63_():
                            coll25_ = _dafny.Map()
                            compr_25_: int
                            for compr_25_ in _dafny.IntegerRange(763, -353):
                                d_962_v8_: int = compr_25_
                                if ((763) <= (d_962_v8_)) and ((d_962_v8_) < (-353)):
                                    coll25_[(d_962_v8_) * (d_960_v1_)] = (self).f20
                            return _dafny.Map(coll25_)
                        return iife63_()
                        

                    return lambda43_

                init24_ = lambda42_(d_951_v1_)
                nw132_ = _dafny.Array(None, 19)
                for i0_24_ in range(nw132_.length(0)):
                    nw132_[i0_24_] = init24_(i0_24_)
                d_959_v9_ = nw132_
                d_963_v10_: _dafny.Array
                nw133_ = _dafny.Array(None, 24)
                nw133_[int(0)] = d_959_v9_
                nw133_[int(1)] = d_959_v9_
                nw133_[int(2)] = d_959_v9_
                nw133_[int(3)] = d_959_v9_
                nw133_[int(4)] = d_959_v9_
                nw133_[int(5)] = d_959_v9_
                nw133_[int(6)] = (d_959_v9_ if (self).f20 else d_959_v9_)
                nw133_[int(7)] = d_959_v9_
                nw133_[int(8)] = d_959_v9_
                nw133_[int(9)] = d_959_v9_
                nw133_[int(10)] = d_959_v9_
                nw133_[int(11)] = d_959_v9_
                nw133_[int(12)] = d_959_v9_
                nw133_[int(13)] = (d_959_v9_ if (self).f10 else d_959_v9_)
                nw133_[int(14)] = d_959_v9_
                nw133_[int(15)] = d_959_v9_
                nw133_[int(16)] = d_959_v9_
                nw133_[int(17)] = d_959_v9_
                nw133_[int(18)] = d_959_v9_
                nw133_[int(19)] = d_959_v9_
                nw133_[int(20)] = d_959_v9_
                nw133_[int(21)] = d_959_v9_
                nw133_[int(22)] = d_959_v9_
                nw133_[int(23)] = d_959_v9_
                d_963_v10_ = nw133_
                index156_ = default__.safeIndex(299, (d_963_v10_).length(0))
                (d_963_v10_)[index156_] = d_959_v9_
                index157_ = default__.safeIndex(299, (d_963_v10_).length(0))
                (d_963_v10_)[index157_] = d_959_v9_
                d_964_v11_: T1
                nw134_ = C2()
                nw134_.ctor__((self).f20, d_951_v1_)
                d_964_v11_ = nw134_
                d_965_v12_: _dafny.Map
                d_965_v12_ = _dafny.Map({d_964_v11_: d_950_v0_})
                d_966_v13_: _dafny.Map
                d_966_v13_ = _dafny.Map({(self).f10: (d_965_v12_).set(d_964_v11_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sywu")))})
                d_967_v14_: _dafny.Array
                nw135_ = _dafny.Array(None, 4)
                nw135_[int(0)] = ((d_966_v13_)[(self).f10] if ((self).f10) in (d_966_v13_) else d_965_v12_)
                nw135_[int(1)] = d_965_v12_
                nw135_[int(2)] = (_dafny.Map({d_964_v11_: d_950_v0_})) | (d_965_v12_)
                nw135_[int(3)] = d_965_v12_
                d_967_v14_ = nw135_
                index158_ = default__.safeIndex(490, (d_967_v14_).length(0))
                (d_967_v14_)[index158_] = d_965_v12_
                d_968_v15_: _dafny.Seq
                d_968_v15_ = _dafny.SeqWithoutIsStrInference([d_965_v12_, d_965_v12_])
                index159_ = default__.safeIndex(490, (d_967_v14_).length(0))
                (d_967_v14_)[index159_] = ((d_965_v12_) | (d_965_v12_)) | ((d_968_v15_)[default__.safeIndex(d_951_v1_, len(d_968_v15_))])
                (globalState).f1 = d_951_v1_
                d_969_v16_: _dafny.Map
                d_969_v16_ = _dafny.Map({(self).f20: d_950_v0_})
                d_970_v17_: _dafny.Map
                d_970_v17_ = _dafny.Map({len(d_969_v16_): d_951_v1_})
                d_971_v18_: _dafny.Seq
                d_971_v18_ = _dafny.SeqWithoutIsStrInference([(d_970_v17_) | ((d_970_v17_).set(d_951_v1_, d_951_v1_))])
                d_971_v18_ = _dafny.SeqWithoutIsStrInference([d_970_v17_])
                d_972_v19_: C2
                nw136_ = C2()
                nw136_.ctor__((d_954_v4_)[default__.safeIndex(d_951_v1_, len(d_954_v4_))], d_951_v1_)
                d_972_v19_ = nw136_
                d_972_v19_ = d_972_v19_
            elif True:
                d_973_v20_: _dafny.Array
                nw137_ = _dafny.Array(_dafny.Map({}), 2)
                d_973_v20_ = nw137_
                d_974_v21_: D4
                d_974_v21_ = D4_DC14(-209, 678)
                d_975_v22_: _dafny.Map
                d_975_v22_ = _dafny.Map({d_951_v1_: d_974_v21_})
                d_976_v23_: _dafny.Map
                d_976_v23_ = _dafny.Map({d_951_v1_: d_975_v22_})
                index160_ = default__.safeIndex(215, (d_973_v20_).length(0))
                (d_973_v20_)[index160_] = ((d_976_v23_)[d_951_v1_] if (d_951_v1_) in (d_976_v23_) else d_975_v22_)
                index161_ = default__.safeIndex(215, (d_973_v20_).length(0))
                (d_973_v20_)[index161_] = d_975_v22_
                d_977_v24_: _dafny.Array
                def lambda44_(d_978_v1_):
                    def lambda45_(d_979_i2_):
                        return D7_DC24(d_978_v1_)

                    return lambda45_

                init25_ = lambda44_(d_951_v1_)
                nw138_ = _dafny.Array(None, 26)
                for i0_25_ in range(nw138_.length(0)):
                    nw138_[i0_25_] = init25_(i0_25_)
                d_977_v24_ = nw138_
                d_980_v25_: D7
                d_980_v25_ = D7_DC24(len(d_950_v0_))
                index162_ = default__.safeIndex(707, (d_977_v24_).length(0))
                (d_977_v24_)[index162_] = d_980_v25_
                pat_let_tv21_ = d_951_v1_
                index163_ = default__.safeIndex(707, (d_977_v24_).length(0))
                def iife64_(_pat_let19_0):
                    def iife65_(d_981_dt__update__tmp_h0_):
                        def iife66_(_pat_let20_0):
                            def iife67_(d_982_dt__update_hcf41_h0_):
                                return D7_DC24(d_982_dt__update_hcf41_h0_)
                            return iife67_(_pat_let20_0)
                        return iife66_(pat_let_tv21_)
                    return iife65_(_pat_let19_0)
                (d_977_v24_)[index163_] = iife64_(d_980_v25_)
                d_983_v26_: _dafny.Array
                nw139_ = _dafny.Array(_dafny.MultiSet({}), 15)
                d_983_v26_ = nw139_
                d_984_v27_: _dafny.MultiSet
                d_984_v27_ = _dafny.MultiSet([d_950_v0_, d_950_v0_])
                index164_ = default__.safeIndex(751, (d_983_v26_).length(0))
                (d_983_v26_)[index164_] = d_984_v27_
                d_985_v28_: D5
                d_985_v28_ = D5_DC16(_dafny.MultiSet([(self).f10, (self).f20, (self).f20, (self).f10]))
                d_986_v29_: _dafny.Set
                d_986_v29_ = _dafny.Set({d_985_v28_})
                index165_ = default__.safeIndex(751, (d_983_v26_).length(0))
                rhs154_ = 544
                rhs155_ = d_984_v27_
                rhs156_ = d_952_v2_
                rhs157_ = d_951_v1_
                rhs158_ = default__.fm4((d_986_v29_).isdisjoint(d_986_v29_), (self).f20, (self).f10, globalState)
                lhs107_ = d_983_v26_
                lhs108_ = default__.safeIndex(751, (d_983_v26_).length(0))
                lhs109_ = globalState
                d_951_v1_ = rhs154_
                lhs107_[lhs108_] = rhs155_
                d_952_v2_ = rhs156_
                lhs109_.f0 = rhs157_
                d_952_v2_ = rhs158_
                d_987_v30_: _dafny.Array
                nw140_ = _dafny.Array(int(0), 6)
                d_987_v30_ = nw140_
                index166_ = default__.safeIndex(860, (d_987_v30_).length(0))
                (d_987_v30_)[index166_] = d_951_v1_
                d_988_v31_: _dafny.Seq
                d_988_v31_ = _dafny.SeqWithoutIsStrInference([d_951_v1_])
                d_989_v32_: _dafny.Map
                d_989_v32_ = _dafny.Map({d_950_v0_: len(d_988_v31_)})
                index167_ = default__.safeIndex(860, (d_987_v30_).length(0))
                (d_987_v30_)[index167_] = len((d_989_v32_) | ((default__.fm44(d_951_v1_, globalState)) | (d_989_v32_)))
                r1 = (self).fm6((d_987_v30_)[default__.safeIndex(860, (d_987_v30_).length(0))], (d_987_v30_)[default__.safeIndex(860, (d_987_v30_).length(0))], ((d_984_v27_)[_dafny.SeqWithoutIsStrInference([d_952_v2_ for d_990_i3_ in range(default__.abs(747))])] if (_dafny.SeqWithoutIsStrInference([d_952_v2_ for d_991_i3_ in range(default__.abs(747))])) in (d_984_v27_) else d_951_v1_), (self).f10, globalState)
            r1 = (self).f20
        elif True:
            d_992_v33_: _dafny.Map
            d_992_v33_ = _dafny.Map({(self).f20: d_950_v0_})
            d_992_v33_ = (d_992_v33_).set((self).f10, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsvlui")))
            d_993_v34_: _dafny.Map
            d_993_v34_ = _dafny.Map({-146: d_951_v1_})
            d_994_v35_: _dafny.Seq
            d_994_v35_ = _dafny.SeqWithoutIsStrInference([d_951_v1_])
            d_995_v36_: _dafny.Map
            d_995_v36_ = _dafny.Map({d_951_v1_: (self).f20})
            rhs159_ = (self).f20
            rhs160_ = (default__.fm45(_dafny.SeqWithoutIsStrInference([d_994_v35_]), d_951_v1_, globalState)).set(len((_dafny.Map({d_951_v1_: (self).f20})) | (d_995_v36_)), (d_951_v1_) * (d_951_v1_))
            r1 = rhs159_
            d_993_v34_ = rhs160_
            r2 = ((d_951_v1_) * (d_951_v1_)) * (default__.fm1(globalState))
            d_994_v35_ = d_994_v35_
            d_996_v37_: _dafny.Seq
            d_996_v37_ = _dafny.SeqWithoutIsStrInference([d_950_v0_])
            d_997_v38_: _dafny.MultiSet
            d_997_v38_ = _dafny.MultiSet([864, (0) - ((d_994_v35_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_952_v2_ for d_998_i4_ in range(default__.abs(359))])), len(d_994_v35_))]), d_951_v1_, len(d_996_v37_), d_951_v1_])
            d_999_v39_: _dafny.Array
            nw141_ = _dafny.Array(None, 12)
            nw141_[int(0)] = d_951_v1_
            nw141_[int(1)] = d_951_v1_
            nw141_[int(2)] = ((d_997_v38_).cardinality) * (d_951_v1_)
            nw141_[int(3)] = d_951_v1_
            nw141_[int(4)] = (d_951_v1_) * (d_951_v1_)
            nw141_[int(5)] = (293) * (d_951_v1_)
            nw141_[int(6)] = d_951_v1_
            nw141_[int(7)] = (0) - (default__.fm1(globalState))
            nw141_[int(8)] = d_951_v1_
            nw141_[int(9)] = default__.fm1(globalState)
            nw141_[int(10)] = d_951_v1_
            nw141_[int(11)] = (default__.fm46((self).f10, (self).f20, (self).f10, globalState)).cf52
            d_999_v39_ = nw141_
            d_999_v39_ = (d_999_v39_ if (self).f20 else d_999_v39_)
        d_1000_v40_: D11
        d_1000_v40_ = D11_DC33((self).f20, d_951_v1_)
        d_1000_v40_ = d_1000_v40_
        d_1001_v41_: D7
        d_1001_v41_ = D7_DC23()
        d_1002_v42_: _dafny.Seq
        d_1002_v42_ = _dafny.SeqWithoutIsStrInference([d_1001_v41_, d_1001_v41_])
        d_1003_v43_: D12
        d_1003_v43_ = D12_DC35(_dafny.SeqWithoutIsStrInference([D10_DC31((self).f20, (self).f10, -733) for d_1004_i5_ in range(default__.abs(397))]))
        d_1005_v44_: _dafny.Map
        d_1005_v44_ = _dafny.Map({d_1003_v43_: d_950_v0_})
        d_951_v1_ = len((d_1002_v42_).set(default__.safeIndex(len((d_1005_v44_).set(d_1003_v43_, d_950_v0_)), len(d_1002_v42_)), d_1001_v41_))
        r1 = True
        d_1006_v45_: _dafny.Array
        def lambda46_(d_1007_v1_, d_1008_v0_):
            def lambda47_(d_1009_i6_):
                return (_dafny.SeqWithoutIsStrInference([d_1007_v1_, d_1007_v1_])) + (_dafny.SeqWithoutIsStrInference([d_1007_v1_, len(d_1008_v0_)]))

            return lambda47_

        init26_ = lambda46_(d_951_v1_, d_950_v0_)
        nw142_ = _dafny.Array(None, 16)
        for i0_26_ in range(nw142_.length(0)):
            nw142_[i0_26_] = init26_(i0_26_)
        d_1006_v45_ = nw142_
        d_1006_v45_ = (d_1006_v45_ if (self).f20 else d_1006_v45_)
        d_1010_v46_: _dafny.Array
        nw143_ = _dafny.Array(False, 2)
        d_1010_v46_ = nw143_
        d_1011_v47_: D0
        d_1011_v47_ = D0_DC2((self).f10, d_1010_v46_)
        source22_ = d_1011_v47_
        if source22_.is_DC0:
            d_1012___mcc_h0_ = source22_.cf0
            d_1013_cf0_ = d_1012___mcc_h0_
            d_1014_v48_: _dafny.Set
            d_1014_v48_ = _dafny.Set({d_951_v1_, d_951_v1_, d_951_v1_, d_951_v1_, d_951_v1_})
            d_1015_v49_: _dafny.Seq
            d_1015_v49_ = _dafny.SeqWithoutIsStrInference([len(d_950_v0_)])
            d_1016_v50_: _dafny.Map
            d_1016_v50_ = _dafny.Map({(self).fm8(len(d_950_v0_), (self).f10, (self).f10, globalState): (d_1014_v48_) in (_dafny.Set({d_1014_v48_, default__.fm35(d_1015_v49_, d_1015_v49_, (self).f10, globalState)}))})
            d_1017_v51_: _dafny.Array
            def lambda48_(d_1018_v1_):
                def lambda49_(d_1019_i7_):
                    return _dafny.Map({(self).f20: not(not((D6_DC20((self).f10, (self).f10, (self).f10, d_1018_v1_)).cf33))})

                return lambda49_

            init27_ = lambda48_(d_951_v1_)
            nw144_ = _dafny.Array(None, 22)
            for i0_27_ in range(nw144_.length(0)):
                nw144_[i0_27_] = init27_(i0_27_)
            d_1017_v51_ = nw144_
            rhs161_ = d_1006_v45_
            rhs162_ = (self).f10
            rhs163_ = (_dafny.Map({False: True})) | ((d_1016_v50_) | (_dafny.Map({(self).f10: (self).f10})))
            rhs164_ = d_1017_v51_
            d_1006_v45_ = rhs161_
            r1 = rhs162_
            d_1016_v50_ = rhs163_
            d_1017_v51_ = rhs164_
            d_1020_v52_: _dafny.Set
            d_1020_v52_ = _dafny.Set({(self).f20, not((self).f20)})
            d_1021_v53_: _dafny.Set
            d_1021_v53_ = _dafny.Set({d_1020_v52_})
            d_1022_v54_: D3
            d_1022_v54_ = D3_DC9(d_1021_v53_)
            d_1023_v55_: D3
            d_1023_v55_ = D3_DC12(d_1022_v54_)
            d_1024_v56_: D3
            d_1024_v56_ = D3_DC12(d_1023_v55_)
            d_1024_v56_ = d_1024_v56_
            d_1025_v57_: _dafny.Array
            nw145_ = _dafny.Array(int(0), 1)
            d_1025_v57_ = nw145_
            index168_ = default__.safeIndex(118, (d_1025_v57_).length(0))
            (d_1025_v57_)[index168_] = d_951_v1_
            index169_ = default__.safeIndex(118, (d_1025_v57_).length(0))
            (d_1025_v57_)[index169_] = d_951_v1_
            if ((self).f10) in (d_954_v4_):
                d_1026_v58_: _dafny.Map
                d_1026_v58_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_952_v2_ for d_1027_i8_ in range(default__.abs(911))]): (self).f10})
                index170_ = default__.safeIndex(519, (d_1013_cf0_).length(0))
                (d_1013_cf0_)[index170_] = not(((d_1026_v58_)[(d_950_v0_).set(default__.safeIndex(d_951_v1_, len(d_950_v0_)), d_952_v2_)] if ((d_950_v0_).set(default__.safeIndex(d_951_v1_, len(d_950_v0_)), d_952_v2_)) in (d_1026_v58_) else (self).f20))
                index171_ = default__.safeIndex(519, (d_1013_cf0_).length(0))
                (d_1013_cf0_)[index171_] = not ((self).f20) or ((self).f10)
                d_1028_v59_: _dafny.Array
                nw146_ = _dafny.Array(None, 1)
                d_1028_v59_ = nw146_
                d_1029_v60_: C1
                nw147_ = C1()
                nw147_.ctor__((self).f10)
                d_1029_v60_ = nw147_
                index172_ = default__.safeIndex(315, (d_1028_v59_).length(0))
                (d_1028_v59_)[index172_] = d_1029_v60_
                index173_ = default__.safeIndex(315, (d_1028_v59_).length(0))
                nw148_ = C1()
                nw148_.ctor__((d_1020_v52_).ispropersubset(d_1020_v52_))
                (d_1028_v59_)[index173_] = nw148_
                (globalState).f7 = (d_1025_v57_)[default__.safeIndex(118, (d_1025_v57_).length(0))]
                (d_1029_v60_).m5(len(d_1015_v49_), globalState)
                r1 = default__.fm0(d_951_v1_, globalState)
            elif True:
                d_1030_v61_: C0
                nw149_ = C0()
                nw149_.ctor__()
                d_1030_v61_ = nw149_
                r1 = (((d_1025_v57_)[default__.safeIndex(118, (d_1025_v57_).length(0))]) + ((d_1025_v57_)[default__.safeIndex(118, (d_1025_v57_).length(0))])) < ((d_1025_v57_)[default__.safeIndex(118, (d_1025_v57_).length(0))])
                d_1031_v62_: _dafny.Map
                d_1031_v62_ = _dafny.Map({(d_1025_v57_)[default__.safeIndex(118, (d_1025_v57_).length(0))]: 143})
                d_1032_v63_: _dafny.Map
                d_1032_v63_ = _dafny.Map({not((908) == (((d_1031_v62_)[(d_1025_v57_)[default__.safeIndex(118, (d_1025_v57_).length(0))]] if ((d_1025_v57_)[default__.safeIndex(118, (d_1025_v57_).length(0))]) in (d_1031_v62_) else d_951_v1_))): d_951_v1_})
                d_1033_v64_: D3
                d_1033_v64_ = D3_DC10(d_950_v0_, (self).f20, (self).f20)
                index174_ = default__.safeIndex(118, (d_1025_v57_).length(0))
                rhs165_ = ((d_1032_v63_)[(self).f20] if ((self).f20) in (d_1032_v63_) else (D2_DC7((d_1025_v57_)[default__.safeIndex(118, (d_1025_v57_).length(0))], (self).f10, (d_1025_v57_)[default__.safeIndex(118, (d_1025_v57_).length(0))])).cf11)
                rhs166_ = not((self).f10)
                rhs167_ = ((d_1033_v64_).cf14) + (d_950_v0_)
                lhs110_ = d_1025_v57_
                lhs111_ = default__.safeIndex(118, (d_1025_v57_).length(0))
                lhs110_[lhs111_] = rhs165_
                r1 = rhs166_
                d_950_v0_ = rhs167_
                r1 = (self).f10
                r1 = (self).f10
        elif source22_.is_DC1:
            d_1034___mcc_h1_ = source22_.cf1
            d_1035_cf1_ = d_1034___mcc_h1_
            d_950_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvlqw"))
            d_1036_v65_: _dafny.Seq
            d_1036_v65_ = _dafny.SeqWithoutIsStrInference([d_951_v1_])
            d_1037_v66_: C2
            nw150_ = C2()
            nw150_.ctor__((d_951_v1_) < (-760), len(d_1036_v65_))
            d_1037_v66_ = nw150_
            d_1037_v66_ = d_1037_v66_
            d_1038_v67_: _dafny.Map
            d_1038_v67_ = _dafny.Map({(self).f10: (self).f20})
            r1 = ((d_1038_v67_)[(d_1037_v66_).f21] if ((d_1037_v66_).f21) in (d_1038_v67_) else (d_1037_v66_).f21)
            (d_1037_v66_).m5(default__.safeDivisionInt((d_1037_v66_).f22, (d_1037_v66_).f22), globalState)
        elif True:
            d_1039___mcc_h2_ = source22_.cf2
            d_1040___mcc_h3_ = source22_.cf3
            d_1041_cf3_ = d_1040___mcc_h3_
            d_1042_cf2_ = d_1039___mcc_h2_
            d_1043_v68_: _dafny.Seq
            d_1043_v68_ = _dafny.SeqWithoutIsStrInference([324, d_951_v1_])
            d_1044_v69_: _dafny.Seq
            d_1044_v69_ = _dafny.SeqWithoutIsStrInference([d_1043_v68_, d_1043_v68_, d_1043_v68_, (d_1043_v68_) + (d_1043_v68_), _dafny.SeqWithoutIsStrInference([d_951_v1_ for d_1045_i9_ in range(default__.abs(884))])])
            d_1046_v70_: _dafny.Map
            d_1046_v70_ = _dafny.Map({(self).f10: d_1044_v69_})
            d_1044_v69_ = (((d_1046_v70_)[False] if (False) in (d_1046_v70_) else d_1044_v69_)) + (d_1044_v69_)
            d_1047_v71_: _dafny.Array
            nw151_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_1047_v71_ = nw151_
            d_1048_v72_: _dafny.Array
            nw152_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
            d_1048_v72_ = nw152_
            index175_ = default__.safeIndex(828, (d_1047_v71_).length(0))
            (d_1047_v71_)[index175_] = d_1048_v72_
            index176_ = default__.safeIndex(828, (d_1047_v71_).length(0))
            rhs168_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwtb"))) + (default__.fm38((self).f10, d_951_v1_, globalState))
            rhs169_ = d_1048_v72_
            lhs112_ = d_1047_v71_
            lhs113_ = default__.safeIndex(828, (d_1047_v71_).length(0))
            d_950_v0_ = rhs168_
            lhs112_[lhs113_] = rhs169_
            d_1049_v73_: _dafny.Map
            d_1049_v73_ = _dafny.Map({(d_950_v0_) + (d_950_v0_): False})
            d_1049_v73_ = (d_1049_v73_).set(d_950_v0_, (d_951_v1_) <= (d_951_v1_))
            r1 = not (d_1042_cf2_) or ((self).f10)
        d_1050_v74_: _dafny.Set
        d_1050_v74_ = _dafny.Set({(d_950_v0_) + (d_950_v0_), ((d_950_v0_).set(default__.safeIndex((_dafny.MultiSet([d_951_v1_])).cardinality, len(d_950_v0_)), d_952_v2_)) + (d_950_v0_), d_950_v0_})
        r0 = d_1050_v74_
        r1 = (self).f10
        r2 = (d_951_v1_ if (self).f20 else d_951_v1_)
        return r0, r1, r2

    def m5(self, p0, globalState):
        d_1051_v0_: bool
        d_1051_v0_ = False
        d_1052_v1_: _dafny.Array
        nw153_ = _dafny.Array(False, 18)
        d_1052_v1_ = nw153_
        d_1053_v2_: _dafny.Map
        d_1053_v2_ = _dafny.Map({113: d_1052_v1_})
        d_1054_v3_: _dafny.MultiSet
        d_1054_v3_ = _dafny.MultiSet([381])
        d_1055_v4_: _dafny.Map
        d_1055_v4_ = _dafny.Map({d_1051_v0_: d_1052_v1_})
        d_1056_v5_: _dafny.Array
        nw154_ = _dafny.Array(None, 15)
        nw154_[int(0)] = d_1052_v1_
        nw154_[int(1)] = d_1052_v1_
        nw154_[int(2)] = d_1052_v1_
        nw154_[int(3)] = ((d_1053_v2_)[(d_1054_v3_).cardinality] if ((d_1054_v3_).cardinality) in (d_1053_v2_) else d_1052_v1_)
        nw154_[int(4)] = d_1052_v1_
        nw154_[int(5)] = d_1052_v1_
        nw154_[int(6)] = d_1052_v1_
        nw154_[int(7)] = d_1052_v1_
        nw154_[int(8)] = d_1052_v1_
        nw154_[int(9)] = d_1052_v1_
        nw154_[int(10)] = d_1052_v1_
        nw154_[int(11)] = d_1052_v1_
        nw154_[int(12)] = ((d_1055_v4_)[(self).f20] if ((self).f20) in (d_1055_v4_) else d_1052_v1_)
        nw154_[int(13)] = (d_1052_v1_ if (self).f20 else d_1052_v1_)
        nw154_[int(14)] = d_1052_v1_
        d_1056_v5_ = nw154_
        index177_ = default__.safeIndex(51, (d_1056_v5_).length(0))
        (d_1056_v5_)[index177_] = d_1052_v1_
        d_1057_v6_: _dafny.Seq
        d_1057_v6_ = _dafny.SeqWithoutIsStrInference([(self).f20])
        d_1058_v7_: _dafny.Seq
        d_1058_v7_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0, -966])
        d_1059_v8_: _dafny.Set
        d_1059_v8_ = _dafny.Set({len(d_1057_v6_), 908, (d_1058_v7_)[default__.safeIndex(p0, len(d_1058_v7_))]})
        d_1060_v9_: _dafny.Map
        d_1060_v9_ = _dafny.Map({p0: ((0) - (len(d_1059_v8_))) >= (p0)})
        d_1061_v10_: _dafny.Map
        d_1061_v10_ = _dafny.Map({d_1051_v0_: ((d_1060_v9_)[p0] if (p0) in (d_1060_v9_) else (self).f20)})
        d_1062_v11_: _dafny.Map
        d_1062_v11_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftlkutv"))): d_1061_v10_})
        d_1063_v12_: D14
        d_1063_v12_ = D14_DC41(d_1062_v11_)
        index178_ = default__.safeIndex(51, (d_1056_v5_).length(0))
        rhs170_ = ((d_1060_v9_)[len((d_1061_v10_) | (d_1061_v10_))] if (len((d_1061_v10_) | (d_1061_v10_))) in (d_1060_v9_) else (self).fm8(-556, (self).f10, True, globalState))
        rhs171_ = d_1052_v1_
        rhs172_ = (len(_dafny.Set({d_1060_v9_}))) == ((p0) + (p0))
        rhs173_ = len((d_1063_v12_).cf71)
        lhs114_ = d_1056_v5_
        lhs115_ = default__.safeIndex(51, (d_1056_v5_).length(0))
        lhs116_ = globalState
        d_1051_v0_ = rhs170_
        lhs114_[lhs115_] = rhs171_
        d_1051_v0_ = rhs172_
        lhs116_.f7 = rhs173_
        def iife68_():
            coll26_ = _dafny.Map()
            compr_26_: int
            for compr_26_ in _dafny.IntegerRange(65, 828):
                d_1064_v13_: int = compr_26_
                if ((65) <= (d_1064_v13_)) and ((d_1064_v13_) < (828)):
                    coll26_[(d_1064_v13_) * (p0)] = d_1051_v0_
            return _dafny.Map(coll26_)
        d_1051_v0_ = (len(iife68_()
        )) <= (p0)
        d_1065_v14_: _dafny.Seq
        d_1065_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gip"))
        d_1066_v15_: _dafny.MultiSet
        d_1066_v15_ = _dafny.MultiSet([(d_1057_v6_)[default__.safeIndex(p0, len(d_1057_v6_))], False])
        d_1067_v16_: _dafny.Array
        nw155_ = _dafny.Array(None, 10)
        nw155_[int(0)] = default__.safeDivisionInt(len(d_1065_v14_), p0)
        nw155_[int(1)] = (p0) * ((d_1066_v15_).cardinality)
        nw155_[int(2)] = ((d_1066_v15_).intersection(_dafny.MultiSet([d_1051_v0_]))).cardinality
        nw155_[int(3)] = p0
        nw155_[int(4)] = p0
        nw155_[int(5)] = (231) - (p0)
        nw155_[int(6)] = p0
        nw155_[int(7)] = default__.fm1(globalState)
        nw155_[int(8)] = p0
        nw155_[int(9)] = p0
        d_1067_v16_ = nw155_
        d_1068_v17_: str
        d_1068_v17_ = _dafny.CodePoint('v')
        rhs174_ = d_1067_v16_
        rhs175_ = (0) - ((p0) * (default__.safeModuloInt(default__.fm1(globalState), p0)))
        rhs176_ = ((D4_DC15(p0, _dafny.MultiSet(d_1058_v7_), d_1068_v17_, p0)).cf26) - (len(d_1057_v6_))
        rhs177_ = len(d_1057_v6_)
        lhs117_ = globalState
        lhs118_ = globalState
        lhs119_ = globalState
        d_1067_v16_ = rhs174_
        lhs117_.f0 = rhs175_
        lhs118_.f1 = rhs176_
        lhs119_.f0 = rhs177_
        d_1069_i0_: int
        d_1069_i0_ = 0
        with _dafny.label("3"):
            while not(not((self).f10)):
                with _dafny.c_label("3"):
                    if (d_1069_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_1069_i0_ = (d_1069_i0_) + (1)
                    d_1051_v0_ = (self).f10
                    d_1055_v4_ = ((d_1055_v4_) | (d_1055_v4_)) | (_dafny.Map({d_1051_v0_: d_1052_v1_}))
                    d_1051_v0_ = True
                    d_1070_v18_: _dafny.Array
                    def lambda50_(d_1071_v15_, d_1072_v14_, d_1073_v9_, d_1074_v0_):
                        def lambda51_(d_1075_i1_):
                            return (d_1071_v15_).set(((d_1073_v9_)[len(d_1072_v14_)] if (len(d_1072_v14_)) in (d_1073_v9_) else d_1074_v0_), default__.abs(len(d_1072_v14_)))

                        return lambda51_

                    init28_ = lambda50_(d_1066_v15_, d_1065_v14_, d_1060_v9_, d_1051_v0_)
                    nw156_ = _dafny.Array(None, 18)
                    for i0_28_ in range(nw156_.length(0)):
                        nw156_[i0_28_] = init28_(i0_28_)
                    d_1070_v18_ = nw156_
                    index179_ = default__.safeIndex(244, (d_1070_v18_).length(0))
                    (d_1070_v18_)[index179_] = (_dafny.MultiSet([d_1051_v0_])) - (d_1066_v15_)
                    index180_ = default__.safeIndex(244, (d_1070_v18_).length(0))
                    (d_1070_v18_)[index180_] = d_1066_v15_
                    pass
            pass
        d_1076_v19_: _dafny.MultiSet
        d_1076_v19_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0]), (d_1058_v7_) + (_dafny.SeqWithoutIsStrInference([p0 for d_1077_i2_ in range(default__.abs(754))])), d_1058_v7_])
        d_1076_v19_ = default__.fm47(p0, (self).f10, globalState)
        index181_ = default__.safeIndex(467, (d_1067_v16_).length(0))
        (d_1067_v16_)[index181_] = p0
        index182_ = default__.safeIndex(467, (d_1067_v16_).length(0))
        (d_1067_v16_)[index182_] = p0

    @property
    def f20(self):
        return self._f20

class C4(T1):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self):
        pass
        pass

    def fm8(self, p0, p1, p2, globalState):
        return not ((not(False)) and (False)) or (not(False))

    def fm9(self, p0, p1, globalState):
        return (0) - (-916)

    def m4(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: int = int(0)
        d_1078_v0_: _dafny.Array
        nw157_ = _dafny.Array(_dafny.Map({}), 12)
        d_1078_v0_ = nw157_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1078_v0_).length(0)):
            d_1079_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1079_i0_)) and ((d_1079_i0_) < ((d_1078_v0_).length(0)))):
                (d_1078_v0_)[(d_1079_i0_)] = _dafny.Map({(not(False) if False else True): D6_DC21(True, False, True, 970)})
        d_1080_v1_: bool
        d_1080_v1_ = False
        d_1081_v2_: int
        d_1081_v2_ = -545
        d_1082_v3_: C2
        nw158_ = C2()
        nw158_.ctor__(d_1080_v1_, d_1081_v2_)
        d_1082_v3_ = nw158_
        d_1083_v4_: _dafny.MultiSet
        d_1083_v4_ = _dafny.MultiSet([d_1080_v1_, (d_1082_v3_).f21])
        d_1083_v4_ = d_1083_v4_
        d_1084_v5_: _dafny.Seq
        d_1084_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tiojau"))
        d_1085_v6_: D10
        d_1085_v6_ = D10_DC30(_dafny.Set({d_1080_v1_, d_1080_v1_, (d_1082_v3_).f21, (d_1082_v3_).f21, (d_1082_v3_).f21}))
        pat_let_tv22_ = d_1084_v5_
        def lambda52_(source23_):
            if source23_.is_DC31:
                d_1086___mcc_h0_ = source23_.cf50
                d_1087___mcc_h1_ = source23_.cf51
                d_1088___mcc_h2_ = source23_.cf52
                d_1089_cf52_ = d_1088___mcc_h2_
                d_1090_cf51_ = d_1087___mcc_h1_
                d_1091_cf50_ = d_1086___mcc_h0_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydrrviwd"))
            elif True:
                d_1092___mcc_h3_ = source23_.cf49
                d_1093_cf49_ = d_1092___mcc_h3_
                return pat_let_tv22_

        d_1084_v5_ = lambda52_(d_1085_v6_)
        r1 = (d_1080_v1_) == (d_1080_v1_)
        d_1094_v7_: _dafny.Map
        d_1094_v7_ = _dafny.Map({d_1080_v1_: (d_1082_v3_).f21})
        d_1095_v8_: _dafny.Set
        d_1095_v8_ = _dafny.Set({d_1094_v7_, d_1094_v7_})
        (globalState).f7 = default__.safeModuloInt((113) + (d_1081_v2_), (self).fm9((d_1084_v5_).set(default__.safeIndex(98, len(d_1084_v5_)), _dafny.CodePoint('m')), d_1095_v8_, globalState))
        d_1096_v9_: _dafny.Map
        d_1096_v9_ = _dafny.Map({not((d_1082_v3_).f21): d_1084_v5_})
        d_1097_v10_: _dafny.Set
        d_1097_v10_ = _dafny.Set({d_1084_v5_})
        r0 = ((_dafny.Set({((d_1096_v9_)[(d_1082_v3_).f21] if ((d_1082_v3_).f21) in (d_1096_v9_) else d_1084_v5_)})).intersection(d_1097_v10_)) | (d_1097_v10_)
        r1 = (d_1082_v3_).f21
        r2 = (d_1082_v3_).f22
        return r0, r1, r2

    def m5(self, p0, globalState):
        d_1098_v0_: bool
        d_1098_v0_ = True
        rhs178_ = d_1098_v0_
        rhs179_ = (default__.fm48(globalState)).cardinality
        rhs180_ = p0
        lhs120_ = globalState
        lhs121_ = globalState
        d_1098_v0_ = rhs178_
        lhs120_.f7 = rhs179_
        lhs121_.f7 = rhs180_
        d_1099_v1_: _dafny.Array
        nw159_ = _dafny.Array(False, 6)
        d_1099_v1_ = nw159_
        index183_ = default__.safeIndex(680, (d_1099_v1_).length(0))
        (d_1099_v1_)[index183_] = True
        index184_ = default__.safeIndex(680, (d_1099_v1_).length(0))
        (d_1099_v1_)[index184_] = (p0) < (p0)
        d_1100_v2_: _dafny.MultiSet
        d_1100_v2_ = _dafny.MultiSet([p0])
        d_1101_v3_: _dafny.Seq
        d_1101_v3_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
        d_1102_i0_: int
        d_1102_i0_ = 0
        with _dafny.label("4"):
            pat_let_tv23_ = d_1098_v0_
            pat_let_tv24_ = d_1098_v0_
            pat_let_tv25_ = d_1098_v0_
            pat_let_tv26_ = d_1099_v1_
            pat_let_tv27_ = d_1099_v1_
            def lambda54_(source24_):
                if source24_.is_DC10:
                    d_1111___mcc_h0_ = source24_.cf14
                    d_1112___mcc_h1_ = source24_.cf15
                    d_1113___mcc_h2_ = source24_.cf16
                    d_1114_cf16_ = d_1113___mcc_h2_
                    d_1115_cf15_ = d_1112___mcc_h1_
                    d_1116_cf14_ = d_1111___mcc_h0_
                    return d_1115_cf15_
                elif source24_.is_DC11:
                    d_1117___mcc_h3_ = source24_.cf17
                    d_1118___mcc_h4_ = source24_.cf18
                    d_1119_cf18_ = d_1118___mcc_h4_
                    d_1120_cf17_ = d_1117___mcc_h3_
                    return (pat_let_tv23_) and (pat_let_tv24_)
                elif source24_.is_DC9:
                    d_1121___mcc_h5_ = source24_.cf13
                    d_1122_cf13_ = d_1121___mcc_h5_
                    return pat_let_tv25_
                elif True:
                    d_1123___mcc_h6_ = source24_.cf19
                    d_1124_cf19_ = d_1123___mcc_h6_
                    return (pat_let_tv27_)[default__.safeIndex(680, (pat_let_tv26_).length(0))]

            while lambda54_(default__.fm49(((d_1100_v2_)[p0] if (p0) in (d_1100_v2_) else (d_1101_v3_)[default__.safeIndex(p0, len(d_1101_v3_))]), p0, globalState)):
                with _dafny.c_label("4"):
                    if (d_1102_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_1102_i0_ = (d_1102_i0_) + (1)
                    d_1103_v4_: _dafny.Seq
                    d_1103_v4_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p0, globalState)])
                    index185_ = default__.safeIndex(680, (d_1099_v1_).length(0))
                    rhs181_ = (len((_dafny.Map({True: False})).set((d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))], (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]))) <= (len(d_1103_v4_))
                    rhs182_ = len((d_1103_v4_).set(default__.safeIndex(962, len(d_1103_v4_)), (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]))
                    lhs122_ = d_1099_v1_
                    lhs123_ = default__.safeIndex(680, (d_1099_v1_).length(0))
                    lhs124_ = globalState
                    lhs122_[lhs123_] = rhs181_
                    lhs124_.f1 = rhs182_
                    d_1104_v5_: _dafny.Array
                    nw160_ = _dafny.Array(_dafny.Array(None, 0), 18)
                    d_1104_v5_ = nw160_
                    d_1105_v6_: _dafny.Array
                    def lambda53_(d_1106_i1_):
                        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmyxcurb"))

                    init29_ = lambda53_
                    nw161_ = _dafny.Array(None, 7)
                    for i0_29_ in range(nw161_.length(0)):
                        nw161_[i0_29_] = init29_(i0_29_)
                    d_1105_v6_ = nw161_
                    index186_ = default__.safeIndex(256, (d_1104_v5_).length(0))
                    (d_1104_v5_)[index186_] = d_1105_v6_
                    index187_ = default__.safeIndex(680, (d_1099_v1_).length(0))
                    index188_ = default__.safeIndex(256, (d_1104_v5_).length(0))
                    rhs183_ = d_1098_v0_
                    rhs184_ = p0
                    rhs185_ = d_1105_v6_
                    rhs186_ = d_1103_v4_
                    lhs125_ = d_1099_v1_
                    lhs126_ = default__.safeIndex(680, (d_1099_v1_).length(0))
                    lhs127_ = globalState
                    lhs128_ = d_1104_v5_
                    lhs129_ = default__.safeIndex(256, (d_1104_v5_).length(0))
                    lhs125_[lhs126_] = rhs183_
                    lhs127_.f7 = rhs184_
                    lhs128_[lhs129_] = rhs185_
                    d_1103_v4_ = rhs186_
                    d_1107_v7_: _dafny.Array
                    nw162_ = _dafny.Array(_dafny.Map({}), 7)
                    d_1107_v7_ = nw162_
                    d_1108_v8_: _dafny.Map
                    d_1108_v8_ = _dafny.Map({False: len(_dafny.Map({True: (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]}))})
                    index189_ = default__.safeIndex(219, (d_1107_v7_).length(0))
                    (d_1107_v7_)[index189_] = d_1108_v8_
                    index190_ = default__.safeIndex(219, (d_1107_v7_).length(0))
                    (d_1107_v7_)[index190_] = d_1108_v8_
                    d_1109_v9_: _dafny.Map
                    d_1109_v9_ = _dafny.Map({d_1098_v0_: d_1099_v1_})
                    d_1110_v10_: _dafny.Set
                    d_1110_v10_ = _dafny.Set({p0, len(d_1109_v9_)})
                    d_1098_v0_ = (-52) < ((p0) * (len(d_1110_v10_)))
                    pass
            pass
        hi4_ = p0
        for d_1125_i2_ in range(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1136_i3_ in range(default__.abs(766))])), hi4_):
            d_1126_v11_: _dafny.Seq
            d_1126_v11_ = _dafny.SeqWithoutIsStrInference([(d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]])
            index191_ = default__.safeIndex(680, (d_1099_v1_).length(0))
            (d_1099_v1_)[index191_] = not((d_1126_v11_)[default__.safeIndex(d_1125_i2_, len(d_1126_v11_))])
            d_1127_v12_: str
            d_1127_v12_ = _dafny.CodePoint('b')
            d_1128_v13_: _dafny.Map
            d_1128_v13_ = _dafny.Map({d_1126_v11_: p0})
            index192_ = default__.safeIndex(680, (d_1099_v1_).length(0))
            index193_ = default__.safeIndex(680, (d_1099_v1_).length(0))
            rhs187_ = (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]
            rhs188_ = d_1098_v0_
            rhs189_ = (d_1127_v12_ if (p0) >= (p0) else d_1127_v12_)
            rhs190_ = ((len(d_1128_v13_) if True else p0) if (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))] else d_1125_i2_)
            lhs130_ = d_1099_v1_
            lhs131_ = default__.safeIndex(680, (d_1099_v1_).length(0))
            lhs132_ = d_1099_v1_
            lhs133_ = default__.safeIndex(680, (d_1099_v1_).length(0))
            lhs134_ = globalState
            lhs130_[lhs131_] = rhs187_
            lhs132_[lhs133_] = rhs188_
            d_1127_v12_ = rhs189_
            lhs134_.f0 = rhs190_
            d_1129_v14_: _dafny.Array
            def lambda55_(d_1130_p0_):
                def lambda56_(d_1131_i4_):
                    return (d_1131_i4_) + (d_1130_p0_)

                return lambda56_

            init30_ = lambda55_(p0)
            nw163_ = _dafny.Array(None, 24)
            for i0_30_ in range(nw163_.length(0)):
                nw163_[i0_30_] = init30_(i0_30_)
            d_1129_v14_ = nw163_
            d_1129_v14_ = d_1129_v14_
            d_1132_v15_: _dafny.Seq
            d_1132_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evnnhnsnk"))
            d_1133_v16_: D1
            d_1133_v16_ = D1_DC4((0) - (p0), d_1129_v14_)
            d_1134_v17_: D14
            d_1134_v17_ = D14_DC42(d_1133_v16_, d_1126_v11_, False)
            d_1135_v18_: _dafny.Map
            d_1135_v18_ = _dafny.Map({d_1134_v17_: (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]})
            rhs191_ = (default__.safeDivisionInt(p0, len((d_1126_v11_).set(default__.safeIndex(p0, len(d_1126_v11_)), (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))])))) == ((len(d_1132_v15_)) * (p0))
            rhs192_ = (p0) - (default__.safeModuloInt(len(d_1135_v18_), d_1125_i2_))
            lhs135_ = globalState
            d_1098_v0_ = rhs191_
            lhs135_.f7 = rhs192_
        d_1137_v19_: D11
        d_1137_v19_ = D11_DC33((d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))], p0)
        d_1138_i5_: int
        d_1138_i5_ = 0
        with _dafny.label("5"):
            while ((d_1137_v19_).cf54) or ((d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]):
                with _dafny.c_label("5"):
                    if (d_1138_i5_) >= (100):
                        raise _dafny.Break("5")
                    d_1138_i5_ = (d_1138_i5_) + (1)
                    d_1139_v20_: D13
                    d_1139_v20_ = D13_DC38(_dafny.MultiSet([d_1099_v1_, d_1099_v1_, d_1099_v1_]))
                    d_1140_v21_: _dafny.Map
                    d_1140_v21_ = _dafny.Map({_dafny.Map({p0: (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]}): d_1139_v20_})
                    d_1140_v21_ = d_1140_v21_
                    d_1141_v22_: D7
                    d_1141_v22_ = D7_DC23()
                    d_1141_v22_ = d_1141_v22_
                    d_1142_v23_: _dafny.Map
                    d_1142_v23_ = _dafny.Map({p0: p0})
                    d_1143_v24_: _dafny.Map
                    d_1143_v24_ = _dafny.Map({False: p0})
                    d_1144_v25_: _dafny.Seq
                    d_1144_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovsrjxkbl"))
                    d_1145_v26_: str
                    d_1145_v26_ = _dafny.CodePoint('w')
                    d_1146_v27_: _dafny.Map
                    d_1146_v27_ = _dafny.Map({not(d_1098_v0_): d_1098_v0_})
                    d_1147_v28_: _dafny.MultiSet
                    d_1147_v28_ = _dafny.MultiSet([(d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]])
                    d_1148_v29_: _dafny.Map
                    d_1148_v29_ = _dafny.Map({d_1145_v26_: d_1098_v0_})
                    d_1149_v30_: _dafny.Set
                    d_1149_v30_ = _dafny.Set({_dafny.Map({not(((d_1148_v29_)[d_1145_v26_] if (d_1145_v26_) in (d_1148_v29_) else True)): d_1098_v0_}), d_1146_v27_})
                    d_1150_v31_: _dafny.Seq
                    d_1150_v31_ = _dafny.SeqWithoutIsStrInference([default__.fm3(len(_dafny.SeqWithoutIsStrInference([d_1101_v3_ for d_1151_i6_ in range(default__.abs(840))])), -277, globalState)])
                    d_1152_v32_: _dafny.Array
                    nw164_ = _dafny.Array(None, 27)
                    nw164_[int(0)] = ((d_1142_v23_)[698] if (698) in (d_1142_v23_) else 485)
                    nw164_[int(1)] = (0) - (len((d_1143_v24_) | (d_1143_v24_)))
                    nw164_[int(2)] = p0
                    nw164_[int(3)] = (self).fm9(((d_1144_v25_).set(default__.safeIndex(p0, len(d_1144_v25_)), d_1145_v26_)).set(default__.safeIndex(p0, len((d_1144_v25_).set(default__.safeIndex(p0, len(d_1144_v25_)), d_1145_v26_))), d_1145_v26_), _dafny.Set({d_1146_v27_, d_1146_v27_}), globalState)
                    nw164_[int(4)] = (d_1147_v28_).cardinality
                    nw164_[int(5)] = 802
                    nw164_[int(6)] = p0
                    nw164_[int(7)] = p0
                    nw164_[int(8)] = p0
                    nw164_[int(9)] = (self).fm9(d_1144_v25_, d_1149_v30_, globalState)
                    nw164_[int(10)] = p0
                    nw164_[int(11)] = len(d_1143_v24_)
                    nw164_[int(12)] = ((self).fm9(d_1144_v25_, _dafny.Set({_dafny.Map({(d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]: d_1098_v0_})}), globalState)) - (p0)
                    nw164_[int(13)] = p0
                    nw164_[int(14)] = p0
                    nw164_[int(15)] = 204
                    nw164_[int(16)] = (p0 if not(d_1098_v0_) else p0)
                    nw164_[int(17)] = 867
                    nw164_[int(18)] = (0) - ((0) - (p0))
                    nw164_[int(19)] = p0
                    nw164_[int(20)] = p0
                    nw164_[int(21)] = default__.safeModuloInt(p0, p0)
                    nw164_[int(22)] = (14) + (p0)
                    nw164_[int(23)] = p0
                    nw164_[int(24)] = default__.safeDivisionInt(p0, (self).fm9(d_1144_v25_, d_1149_v30_, globalState))
                    nw164_[int(25)] = (p0) - (len((d_1150_v31_)[default__.safeIndex(p0, len(d_1150_v31_))]))
                    nw164_[int(26)] = (0) - (p0)
                    d_1152_v32_ = nw164_
                    index194_ = default__.safeIndex(380, (d_1152_v32_).length(0))
                    (d_1152_v32_)[index194_] = len((d_1143_v24_) | (d_1143_v24_))
                    d_1153_v33_: D4
                    d_1153_v33_ = D4_DC15(p0, d_1100_v2_, d_1145_v26_, len(d_1146_v27_))
                    index195_ = default__.safeIndex(680, (d_1099_v1_).length(0))
                    index196_ = default__.safeIndex(380, (d_1152_v32_).length(0))
                    rhs193_ = (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]
                    rhs194_ = p0
                    rhs195_ = d_1153_v33_
                    rhs196_ = p0
                    lhs136_ = d_1099_v1_
                    lhs137_ = default__.safeIndex(680, (d_1099_v1_).length(0))
                    lhs138_ = d_1152_v32_
                    lhs139_ = default__.safeIndex(380, (d_1152_v32_).length(0))
                    lhs140_ = globalState
                    lhs136_[lhs137_] = rhs193_
                    lhs138_[lhs139_] = rhs194_
                    d_1153_v33_ = rhs195_
                    lhs140_.f7 = rhs196_
                    d_1154_v34_: _dafny.Array
                    def lambda57_(d_1155_v25_):
                        def lambda58_(d_1156_i7_):
                            return _dafny.MultiSet(d_1155_v25_)

                        return lambda58_

                    init31_ = lambda57_(d_1144_v25_)
                    nw165_ = _dafny.Array(None, 12)
                    for i0_31_ in range(nw165_.length(0)):
                        nw165_[i0_31_] = init31_(i0_31_)
                    d_1154_v34_ = nw165_
                    d_1157_v35_: _dafny.Map
                    d_1157_v35_ = _dafny.Map({d_1154_v34_: (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]})
                    d_1157_v35_ = (d_1157_v35_).set(d_1154_v34_, d_1098_v0_)
                    pass
            pass
        hi5_ = (p0 if default__.fm0(216, globalState) else p0)
        for d_1158_i8_ in range(p0, hi5_):
            d_1159_v36_: _dafny.Set
            d_1159_v36_ = _dafny.Set({p0, default__.fm1(globalState)})
            (globalState).f1 = (default__.safeModuloInt(len(d_1159_v36_), len(d_1101_v3_))) + (701)
            d_1160_v37_: _dafny.Seq
            d_1160_v37_ = _dafny.SeqWithoutIsStrInference([d_1098_v0_, default__.fm0((0) - (p0), globalState), ((d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))] if (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))] else (d_1099_v1_)[default__.safeIndex(680, (d_1099_v1_).length(0))]), ((_dafny.MultiSet(d_1101_v3_)).cardinality) < (d_1158_i8_)])
            index197_ = default__.safeIndex(680, (d_1099_v1_).length(0))
            (d_1099_v1_)[index197_] = (d_1160_v37_)[default__.safeIndex((_dafny.MultiSet(d_1160_v37_)).cardinality, len(d_1160_v37_))]
            d_1161_v38_: _dafny.Seq
            d_1161_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tleyvejo"))
            d_1162_v39_: str
            d_1162_v39_ = _dafny.CodePoint('u')
            d_1163_v40_: T1
            nw166_ = C1()
            nw166_.ctor__(d_1098_v0_)
            d_1163_v40_ = nw166_
            d_1164_v41_: _dafny.Set
            d_1164_v41_ = _dafny.Set({d_1163_v40_})
            (globalState).f7 = (len(((d_1161_v38_).set(default__.safeIndex(d_1158_i8_, len(d_1161_v38_)), d_1162_v39_)).set(default__.safeIndex(default__.fm1(globalState), len((d_1161_v38_).set(default__.safeIndex(d_1158_i8_, len(d_1161_v38_)), d_1162_v39_))), d_1162_v39_))) + (len(d_1164_v41_))
            d_1165_v42_: int
            d_1166_v43_: _dafny.Set
            out36_: int
            out37_: _dafny.Set
            out36_, out37_ = (self).m14(globalState)
            d_1165_v42_ = out36_
            d_1166_v43_ = out37_

    def m14(self, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        d_1167_v0_: int
        d_1167_v0_ = 154
        d_1168_v1_: _dafny.Array
        nw167_ = _dafny.Array(False, 24)
        d_1168_v1_ = nw167_
        d_1169_v2_: _dafny.Seq
        d_1170_v3_: _dafny.Seq
        out38_: _dafny.Seq
        out39_: _dafny.Seq
        out38_, out39_ = default__.m0(d_1167_v0_, d_1168_v1_, globalState)
        d_1169_v2_ = out38_
        d_1170_v3_ = out39_
        d_1171_v4_: bool
        d_1171_v4_ = True
        d_1172_v5_: _dafny.Set
        d_1172_v5_ = _dafny.Set({d_1171_v4_, d_1171_v4_})
        d_1173_v6_: _dafny.Map
        d_1173_v6_ = _dafny.Map({len(d_1172_v5_): d_1167_v0_})
        if (d_1167_v0_) not in (d_1173_v6_):
            d_1171_v4_ = True
            index198_ = default__.safeIndex(554, (d_1168_v1_).length(0))
            (d_1168_v1_)[index198_] = d_1171_v4_
            index199_ = default__.safeIndex(554, (d_1168_v1_).length(0))
            (d_1168_v1_)[index199_] = d_1171_v4_
            d_1174_v7_: _dafny.Array
            nw168_ = _dafny.Array(False, 29)
            d_1174_v7_ = nw168_
            d_1175_v8_: _dafny.Seq
            d_1176_v9_: _dafny.Seq
            out40_: _dafny.Seq
            out41_: _dafny.Seq
            out40_, out41_ = default__.m0(default__.safeModuloInt(d_1167_v0_, (0) - (len(d_1169_v2_))), d_1174_v7_, globalState)
            d_1175_v8_ = out40_
            d_1176_v9_ = out41_
            d_1177_v10_: C3
            nw169_ = C3()
            nw169_.ctor__(((d_1168_v1_)[default__.safeIndex(554, (d_1168_v1_).length(0))]) == ((d_1168_v1_)[default__.safeIndex(554, (d_1168_v1_).length(0))]), default__.fm0(366, globalState))
            d_1177_v10_ = nw169_
            d_1178_v11_: D7
            d_1178_v11_ = D7_DC22(d_1173_v6_)
            source25_ = d_1178_v11_
            if source25_.is_DC23:
                d_1179_v12_: C2
                nw170_ = C2()
                nw170_.ctor__(d_1171_v4_, (0) - (d_1167_v0_))
                d_1179_v12_ = nw170_
                d_1171_v4_ = d_1171_v4_
                d_1171_v4_ = (d_1172_v5_).issubset((d_1172_v5_) - (default__.fm30(globalState)))
                d_1180_v13_: _dafny.Map
                d_1180_v13_ = _dafny.Map({(d_1172_v5_).ispropersubset(d_1172_v5_): d_1174_v7_})
                d_1174_v7_ = ((d_1180_v13_)[(d_1179_v12_).f21] if ((d_1179_v12_).f21) in (d_1180_v13_) else d_1174_v7_)
            elif source25_.is_DC24:
                d_1181___mcc_h0_ = source25_.cf41
                d_1182_cf41_ = d_1181___mcc_h0_
                d_1183_v14_: C1
                nw171_ = C1()
                nw171_.ctor__((d_1168_v1_)[default__.safeIndex(554, (d_1168_v1_).length(0))])
                d_1183_v14_ = nw171_
                index200_ = default__.safeIndex(554, (d_1168_v1_).length(0))
                rhs197_ = d_1183_v14_
                rhs198_ = (d_1177_v10_).f20
                lhs141_ = d_1168_v1_
                lhs142_ = default__.safeIndex(554, (d_1168_v1_).length(0))
                d_1183_v14_ = rhs197_
                lhs141_[lhs142_] = rhs198_
                d_1184_v15_: _dafny.Map
                d_1184_v15_ = _dafny.Map({962: (d_1177_v10_).f20})
                d_1185_v16_: _dafny.Map
                d_1185_v16_ = _dafny.Map({(d_1177_v10_).f20: (d_1168_v1_)[default__.safeIndex(554, (d_1168_v1_).length(0))]})
                d_1186_v17_: _dafny.MultiSet
                d_1186_v17_ = _dafny.MultiSet([d_1185_v16_])
                (globalState).f1 = len((d_1184_v15_).set((d_1186_v17_).cardinality, (d_1183_v14_).fm6(242, d_1167_v0_, d_1182_cf41_, d_1171_v4_, globalState)))
                index201_ = default__.safeIndex(554, (d_1168_v1_).length(0))
                (d_1168_v1_)[index201_] = d_1171_v4_
                d_1171_v4_ = (d_1182_cf41_) <= (62)
            elif True:
                d_1187___mcc_h1_ = source25_.cf40
                d_1188_cf40_ = d_1187___mcc_h1_
                d_1171_v4_ = not(True)
                d_1189_v18_: _dafny.Array
                nw172_ = _dafny.Array(_dafny.Map({}), 3)
                d_1189_v18_ = nw172_
                d_1190_v19_: _dafny.Map
                d_1190_v19_ = _dafny.Map({(d_1168_v1_)[default__.safeIndex(554, (d_1168_v1_).length(0))]: not(d_1171_v4_)})
                index202_ = default__.safeIndex(928, (d_1189_v18_).length(0))
                (d_1189_v18_)[index202_] = d_1190_v19_
                index203_ = default__.safeIndex(928, (d_1189_v18_).length(0))
                (d_1189_v18_)[index203_] = _dafny.Map({(d_1168_v1_)[default__.safeIndex(554, (d_1168_v1_).length(0))]: (d_1177_v10_).f20})
                d_1191_v20_: _dafny.Map
                d_1191_v20_ = _dafny.Map({not(d_1171_v4_): default__.safeModuloInt(d_1167_v0_, len(d_1176_v9_))})
                d_1191_v20_ = (d_1191_v20_).set((d_1177_v10_).f20, d_1167_v0_)
                d_1192_v21_: str
                d_1192_v21_ = _dafny.CodePoint('l')
                d_1192_v21_ = d_1192_v21_
        elif True:
            d_1193_v22_: C0
            nw173_ = C0()
            nw173_.ctor__()
            d_1193_v22_ = nw173_
            d_1169_v2_ = d_1170_v3_
            d_1194_v23_: _dafny.Set
            d_1194_v23_ = _dafny.Set({d_1168_v1_, d_1168_v1_})
            d_1195_v24_: _dafny.Seq
            d_1195_v24_ = _dafny.SeqWithoutIsStrInference([True, (d_1194_v23_).ispropersubset(d_1194_v23_), (d_1172_v5_) != (d_1172_v5_), d_1171_v4_, (d_1167_v0_) > (d_1167_v0_)])
            d_1196_v25_: _dafny.MultiSet
            d_1196_v25_ = _dafny.MultiSet([d_1171_v4_])
            rhs199_ = (d_1195_v24_)[default__.safeIndex(523, len(d_1195_v24_))]
            rhs200_ = (((d_1196_v25_).set(True, default__.abs(d_1167_v0_))) | (d_1196_v25_)).cardinality
            lhs143_ = globalState
            d_1171_v4_ = rhs199_
            lhs143_.f1 = rhs200_
            index204_ = default__.safeIndex(88, (d_1168_v1_).length(0))
            (d_1168_v1_)[index204_] = default__.fm0(len(d_1170_v3_), globalState)
            index205_ = default__.safeIndex(88, (d_1168_v1_).length(0))
            (d_1168_v1_)[index205_] = (d_1193_v22_).fm37(d_1169_v2_, (0) - (d_1167_v0_), d_1171_v4_, d_1171_v4_, globalState)
            d_1197_v26_: _dafny.Array
            def lambda59_(d_1198_v4_):
                def lambda60_(d_1199_i0_):
                    return d_1198_v4_

                return lambda60_

            init32_ = lambda59_(d_1171_v4_)
            nw174_ = _dafny.Array(None, 17)
            for i0_32_ in range(nw174_.length(0)):
                nw174_[i0_32_] = init32_(i0_32_)
            d_1197_v26_ = nw174_
            d_1200_v27_: _dafny.Seq
            d_1201_v28_: _dafny.Seq
            out42_: _dafny.Seq
            out43_: _dafny.Seq
            out42_, out43_ = default__.m0(d_1167_v0_, d_1197_v26_, globalState)
            d_1200_v27_ = out42_
            d_1201_v28_ = out43_
        r0 = (-298) - (d_1167_v0_)
        d_1202_v29_: C3
        nw175_ = C3()
        nw175_.ctor__(d_1171_v4_, not(d_1171_v4_))
        d_1202_v29_ = nw175_
        if (d_1202_v29_).f20:
            d_1203_v30_: _dafny.Map
            d_1203_v30_ = _dafny.Map({(d_1202_v29_).f20: not((d_1167_v0_) > (d_1167_v0_))})
            d_1204_v31_: _dafny.MultiSet
            d_1204_v31_ = _dafny.MultiSet([d_1167_v0_, d_1167_v0_])
            d_1205_v32_: D3
            d_1205_v32_ = D3_DC11(d_1167_v0_, d_1203_v30_)
            d_1206_v33_: _dafny.Map
            d_1206_v33_ = _dafny.Map({((d_1204_v31_)[(d_1202_v29_).fm9(d_1169_v2_, _dafny.Set({d_1203_v30_, d_1203_v30_, d_1203_v30_, (d_1205_v32_).cf18}), globalState)] if ((d_1202_v29_).fm9(d_1169_v2_, _dafny.Set({d_1203_v30_, d_1203_v30_, d_1203_v30_, (d_1205_v32_).cf18}), globalState)) in (d_1204_v31_) else d_1167_v0_): _dafny.Map({(d_1202_v29_).f20: not((d_1202_v29_).f20)})})
            rhs201_ = d_1168_v1_
            rhs202_ = default__.safeDivisionInt(len((d_1169_v2_) + (d_1170_v3_)), d_1167_v0_)
            rhs203_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1207_i1_ in range(default__.abs(486))])), d_1167_v0_)
            rhs204_ = (d_1202_v29_).f20
            rhs205_ = (((d_1206_v33_)[d_1167_v0_] if (d_1167_v0_) in (d_1206_v33_) else d_1203_v30_)).set((d_1202_v29_).f20, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkc"))) != (d_1169_v2_))
            lhs144_ = globalState
            lhs145_ = globalState
            d_1168_v1_ = rhs201_
            lhs144_.f7 = rhs202_
            lhs145_.f1 = rhs203_
            d_1171_v4_ = rhs204_
            d_1203_v30_ = rhs205_
            d_1208_v34_: _dafny.Array
            def lambda61_(d_1209_v0_):
                def lambda62_(d_1210_i2_):
                    return (d_1210_i2_) * (d_1209_v0_)

                return lambda62_

            init33_ = lambda61_(d_1167_v0_)
            nw176_ = _dafny.Array(None, 6)
            for i0_33_ in range(nw176_.length(0)):
                nw176_[i0_33_] = init33_(i0_33_)
            d_1208_v34_ = nw176_
            index206_ = default__.safeIndex(11, (d_1208_v34_).length(0))
            (d_1208_v34_)[index206_] = 715
            d_1211_v35_: _dafny.Map
            d_1211_v35_ = _dafny.Map({d_1171_v4_: d_1167_v0_})
            d_1212_v36_: _dafny.Set
            d_1212_v36_ = _dafny.Set({d_1211_v35_})
            d_1213_v37_: _dafny.Map
            d_1213_v37_ = _dafny.Map({d_1167_v0_: d_1212_v36_})
            index207_ = default__.safeIndex(11, (d_1208_v34_).length(0))
            (d_1208_v34_)[index207_] = (0) - (len(((d_1213_v37_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqjtt")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqjtt")))) in (d_1213_v37_) else d_1212_v36_)))
            d_1171_v4_ = not(False)
            d_1214_v38_: D10
            d_1214_v38_ = D10_DC30(d_1172_v5_)
            def iife69_(_pat_let21_0):
                def iife70_(d_1215_dt__update__tmp_h0_):
                    def iife71_(_pat_let22_0):
                        def iife72_(d_1216_dt__update_hcf49_h0_):
                            return D10_DC30(d_1216_dt__update_hcf49_h0_)
                        return iife72_(_pat_let22_0)
                    return iife71_(_dafny.Set({False}))
                return iife70_(_pat_let21_0)
            d_1214_v38_ = iife69_(D10_DC30(d_1172_v5_))
            index208_ = default__.safeIndex(538, (d_1168_v1_).length(0))
            (d_1168_v1_)[index208_] = (d_1167_v0_) == (d_1167_v0_)
            index209_ = default__.safeIndex(59, (d_1168_v1_).length(0))
            (d_1168_v1_)[index209_] = (d_1202_v29_).f20
            d_1217_v39_: _dafny.Array
            def lambda63_(d_1218_i3_):
                return _dafny.MultiSet([_dafny.Set({_dafny.CodePoint('i')})])

            init34_ = lambda63_
            nw177_ = _dafny.Array(None, 20)
            for i0_34_ in range(nw177_.length(0)):
                nw177_[i0_34_] = init34_(i0_34_)
            d_1217_v39_ = nw177_
            d_1219_v40_: str
            d_1219_v40_ = _dafny.CodePoint('v')
            d_1220_v41_: _dafny.Set
            d_1220_v41_ = _dafny.Set({d_1219_v40_, d_1219_v40_, d_1219_v40_, d_1219_v40_})
            d_1221_v42_: _dafny.MultiSet
            d_1221_v42_ = _dafny.MultiSet([d_1220_v41_, d_1220_v41_])
            index210_ = default__.safeIndex(712, (d_1217_v39_).length(0))
            (d_1217_v39_)[index210_] = d_1221_v42_
            index211_ = default__.safeIndex(538, (d_1168_v1_).length(0))
            index212_ = default__.safeIndex(59, (d_1168_v1_).length(0))
            index213_ = default__.safeIndex(712, (d_1217_v39_).length(0))
            rhs206_ = not(((d_1208_v34_)[default__.safeIndex(11, (d_1208_v34_).length(0))]) > ((d_1167_v0_) + (273)))
            rhs207_ = d_1167_v0_
            rhs208_ = (((d_1169_v2_).set(default__.safeIndex(len(_dafny.Map({(d_1202_v29_).f20: d_1171_v4_})), len(d_1169_v2_)), _dafny.CodePoint('v'))) + (d_1170_v3_)) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhmbab"))) + (d_1169_v2_))
            rhs209_ = d_1221_v42_
            lhs146_ = d_1168_v1_
            lhs147_ = default__.safeIndex(538, (d_1168_v1_).length(0))
            lhs148_ = globalState
            lhs149_ = d_1168_v1_
            lhs150_ = default__.safeIndex(59, (d_1168_v1_).length(0))
            lhs151_ = d_1217_v39_
            lhs152_ = default__.safeIndex(712, (d_1217_v39_).length(0))
            lhs146_[lhs147_] = rhs206_
            lhs148_.f1 = rhs207_
            lhs149_[lhs150_] = rhs208_
            lhs151_[lhs152_] = rhs209_
        elif True:
            d_1222_v43_: _dafny.Map
            d_1222_v43_ = _dafny.Map({d_1171_v4_: d_1167_v0_})
            d_1223_v44_: str
            d_1223_v44_ = _dafny.CodePoint('n')
            d_1224_v45_: _dafny.Map
            d_1224_v45_ = _dafny.Map({d_1223_v44_: False})
            d_1225_v46_: _dafny.Map
            d_1225_v46_ = _dafny.Map({len(d_1222_v43_): d_1224_v45_})
            d_1226_v47_: bool
            d_1227_v48_: _dafny.Seq
            d_1228_v49_: _dafny.Map
            out44_: bool
            out45_: _dafny.Seq
            out46_: _dafny.Map
            out44_, out45_, out46_ = (d_1202_v29_).m3(d_1225_v46_, d_1171_v4_, globalState)
            d_1226_v47_ = out44_
            d_1227_v48_ = out45_
            d_1228_v49_ = out46_
            d_1229_v50_: _dafny.Map
            d_1229_v50_ = _dafny.Map({d_1167_v0_: (d_1202_v29_).f20})
            d_1230_v51_: _dafny.Map
            d_1230_v51_ = _dafny.Map({d_1229_v50_: d_1171_v4_})
            (globalState).f1 = (len((d_1230_v51_) | (d_1230_v51_))) - (d_1167_v0_)
            d_1231_v52_: _dafny.Seq
            d_1231_v52_ = _dafny.SeqWithoutIsStrInference([(d_1202_v29_).f20])
            d_1232_v53_: _dafny.Map
            d_1232_v53_ = _dafny.Map({d_1167_v0_: d_1231_v52_})
            d_1226_v47_ = (d_1167_v0_) > ((_dafny.MultiSet(((d_1231_v52_) + (((d_1232_v53_)[d_1167_v0_] if (d_1167_v0_) in (d_1232_v53_) else d_1231_v52_))).set(default__.safeIndex(d_1167_v0_, len((d_1231_v52_) + (((d_1232_v53_)[d_1167_v0_] if (d_1167_v0_) in (d_1232_v53_) else d_1231_v52_)))), (d_1231_v52_)[default__.safeIndex(-258, len(d_1231_v52_))]))).cardinality)
            d_1233_v54_: _dafny.Set
            d_1233_v54_ = _dafny.Set({d_1222_v43_})
            d_1226_v47_ = not ((d_1222_v43_) in (d_1233_v54_)) or ((d_1167_v0_) >= (d_1167_v0_))
            d_1234_v55_: _dafny.Array
            nw178_ = _dafny.Array(None, 28)
            nw178_[int(0)] = d_1223_v44_
            nw178_[int(1)] = d_1223_v44_
            nw178_[int(2)] = d_1223_v44_
            nw178_[int(3)] = d_1223_v44_
            nw178_[int(4)] = d_1223_v44_
            nw178_[int(5)] = d_1223_v44_
            nw178_[int(6)] = d_1223_v44_
            nw178_[int(7)] = d_1223_v44_
            nw178_[int(8)] = _dafny.CodePoint('a')
            nw178_[int(9)] = d_1223_v44_
            nw178_[int(10)] = d_1223_v44_
            nw178_[int(11)] = d_1223_v44_
            nw178_[int(12)] = d_1223_v44_
            nw178_[int(13)] = d_1223_v44_
            nw178_[int(14)] = d_1223_v44_
            nw178_[int(15)] = d_1223_v44_
            nw178_[int(16)] = _dafny.CodePoint('e')
            nw178_[int(17)] = _dafny.CodePoint('h')
            nw178_[int(18)] = d_1223_v44_
            nw178_[int(19)] = d_1223_v44_
            nw178_[int(20)] = d_1223_v44_
            nw178_[int(21)] = d_1223_v44_
            nw178_[int(22)] = d_1223_v44_
            nw178_[int(23)] = d_1223_v44_
            nw178_[int(24)] = d_1223_v44_
            nw178_[int(25)] = d_1223_v44_
            nw178_[int(26)] = _dafny.CodePoint('b')
            nw178_[int(27)] = d_1223_v44_
            d_1234_v55_ = nw178_
            d_1235_v56_: D12
            d_1235_v56_ = D12_DC37((0) - (d_1167_v0_), d_1234_v55_, d_1223_v44_)
            d_1236_v57_: D13
            d_1236_v57_ = D13_DC40(d_1235_v56_, not(d_1171_v4_), len(_dafny.SeqWithoutIsStrInference([-227 for d_1237_i4_ in range(default__.abs(628))])), d_1167_v0_)
            (globalState).f1 = (d_1236_v57_).cf69
        d_1238_v58_: _dafny.Array
        nw179_ = _dafny.Array(_dafny.Array(None, 0), 19)
        d_1238_v58_ = nw179_
        d_1238_v58_ = d_1238_v58_
        r0 = d_1167_v0_
        r1 = (_dafny.Set({default__.fm0((0) - (d_1167_v0_), globalState)})).intersection(default__.fm30(globalState))
        return r0, r1


class C5:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self):
        pass
        pass

    def m13(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_1239_v0_: bool
        d_1239_v0_ = True
        if d_1239_v0_:
            d_1240_v1_: _dafny.Array
            def lambda64_(d_1241_v0_):
                def lambda65_(d_1242_i0_):
                    return D3_DC12(D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), d_1241_v0_, d_1241_v0_))

                return lambda65_

            init35_ = lambda64_(d_1239_v0_)
            nw180_ = _dafny.Array(None, 8)
            for i0_35_ in range(nw180_.length(0)):
                nw180_[i0_35_] = init35_(i0_35_)
            d_1240_v1_ = nw180_
            d_1243_v2_: _dafny.Seq
            d_1243_v2_ = _dafny.SeqWithoutIsStrInference([d_1240_v1_])
            d_1244_v3_: _dafny.Map
            d_1244_v3_ = _dafny.Map({897: d_1243_v2_})
            d_1245_v4_: int
            d_1245_v4_ = 303
            d_1244_v3_ = (d_1244_v3_).set((d_1245_v4_) + (d_1245_v4_), (d_1243_v2_).set(default__.safeIndex((0) - (d_1245_v4_), len(d_1243_v2_)), d_1240_v1_))
            d_1246_v5_: C0
            nw181_ = C0()
            nw181_.ctor__()
            d_1246_v5_ = nw181_
            d_1247_v6_: _dafny.Array
            nw182_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
            d_1247_v6_ = nw182_
            d_1247_v6_ = d_1247_v6_
            d_1248_v7_: _dafny.Array
            nw183_ = _dafny.Array(_dafny.CodePoint('D'), 18)
            d_1248_v7_ = nw183_
            r0 = d_1248_v7_
            d_1249_v8_: C0
            nw184_ = C0()
            nw184_.ctor__()
            d_1249_v8_ = nw184_
        elif True:
            d_1250_v9_: int
            d_1250_v9_ = -401
            d_1251_v10_: _dafny.Map
            d_1251_v10_ = _dafny.Map({d_1250_v9_: d_1239_v0_})
            d_1252_v11_: D11
            d_1252_v11_ = D11_DC32(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "em"))): d_1239_v0_}))
            d_1253_v12_: _dafny.Set
            d_1253_v12_ = _dafny.Set({len((d_1251_v10_) | ((d_1252_v11_).cf53))})
            (globalState).f7 = len(d_1253_v12_)
            d_1254_v13_: str
            d_1254_v13_ = _dafny.CodePoint('j')
            d_1254_v13_ = d_1254_v13_
            d_1255_v14_: _dafny.Seq
            d_1255_v14_ = _dafny.SeqWithoutIsStrInference([len(default__.fm40(d_1239_v0_, True, globalState))])
            d_1256_v15_: _dafny.Map
            d_1256_v15_ = _dafny.Map({d_1239_v0_: True})
            d_1257_v16_: D3
            d_1257_v16_ = D3_DC11(len(d_1255_v14_), d_1256_v15_)
            d_1258_v17_: D3
            d_1258_v17_ = D3_DC12(d_1257_v16_)
            d_1259_v18_: D3
            d_1259_v18_ = D3_DC12(d_1258_v17_)
            d_1259_v18_ = D3_DC12(d_1257_v16_)
            d_1260_v19_: _dafny.Seq
            d_1260_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vawvwkwoa"))
            d_1260_v19_ = ((d_1260_v19_) + (d_1260_v19_)) + (d_1260_v19_)
            d_1261_v20_: _dafny.MultiSet
            d_1261_v20_ = _dafny.MultiSet([d_1239_v0_, d_1239_v0_])
            d_1239_v0_ = ((default__.fm41(globalState)).set(not(d_1239_v0_), default__.abs(d_1250_v9_))).ispropersubset(d_1261_v20_)
        d_1262_v21_: _dafny.Seq
        d_1262_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
        d_1263_v22_: _dafny.Seq
        d_1263_v22_ = _dafny.SeqWithoutIsStrInference([d_1239_v0_, d_1239_v0_])
        d_1264_v23_: int
        d_1264_v23_ = 201
        d_1265_v24_: _dafny.Array
        nw185_ = _dafny.Array(int(0), 15)
        d_1265_v24_ = nw185_
        d_1266_v25_: D1
        d_1266_v25_ = D1_DC4(d_1264_v23_, d_1265_v24_)
        d_1267_v26_: D14
        d_1267_v26_ = D14_DC42(d_1266_v25_, d_1263_v22_, d_1239_v0_)
        d_1268_v27_: _dafny.MultiSet
        d_1268_v27_ = _dafny.MultiSet([d_1263_v22_, (d_1263_v22_) + (d_1263_v22_), default__.fm40(d_1239_v0_, d_1239_v0_, globalState), (d_1263_v22_) + (d_1263_v22_), _dafny.SeqWithoutIsStrInference([(d_1267_v26_).cf74])])
        index214_ = default__.safeIndex(587, (d_1265_v24_).length(0))
        (d_1265_v24_)[index214_] = d_1264_v23_
        d_1269_v28_: _dafny.Map
        d_1269_v28_ = _dafny.Map({d_1239_v0_: default__.fm0(d_1264_v23_, globalState)})
        index215_ = default__.safeIndex(587, (d_1265_v24_).length(0))
        rhs210_ = (d_1262_v21_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ys")))
        rhs211_ = d_1268_v27_
        rhs212_ = d_1264_v23_
        rhs213_ = (d_1263_v22_) < ((d_1263_v22_).set(default__.safeIndex((0) - (d_1264_v23_), len(d_1263_v22_)), ((d_1269_v28_)[d_1239_v0_] if (d_1239_v0_) in (d_1269_v28_) else d_1239_v0_)))
        rhs214_ = default__.fm0(d_1264_v23_, globalState)
        lhs153_ = d_1265_v24_
        lhs154_ = default__.safeIndex(587, (d_1265_v24_).length(0))
        d_1262_v21_ = rhs210_
        d_1268_v27_ = rhs211_
        lhs153_[lhs154_] = rhs212_
        d_1239_v0_ = rhs213_
        d_1239_v0_ = rhs214_
        d_1270_v29_: _dafny.Set
        d_1270_v29_ = _dafny.Set({_dafny.Map({d_1239_v0_: True})})
        d_1271_v30_: _dafny.Map
        d_1271_v30_ = _dafny.Map({(p0).fm9(d_1262_v21_, d_1270_v29_, globalState): d_1239_v0_})
        d_1271_v30_ = (d_1271_v30_).set(d_1264_v23_, d_1239_v0_)
        hi6_ = d_1264_v23_
        for d_1272_i1_ in range(d_1264_v23_, hi6_):
            d_1273_v31_: _dafny.Array
            nw186_ = _dafny.Array(D1.default()(), 12)
            d_1273_v31_ = nw186_
            d_1274_v32_: _dafny.Seq
            d_1274_v32_ = _dafny.SeqWithoutIsStrInference([d_1273_v31_, d_1273_v31_, d_1273_v31_])
            d_1275_v33_: _dafny.Seq
            d_1275_v33_ = _dafny.SeqWithoutIsStrInference([d_1272_i1_, d_1264_v23_, d_1272_i1_, d_1264_v23_])
            d_1276_v34_: D8
            d_1276_v34_ = D8_DC27(d_1239_v0_, d_1274_v32_, d_1275_v33_, d_1272_i1_)
            index216_ = default__.safeIndex(587, (d_1265_v24_).length(0))
            (d_1265_v24_)[index216_] = (d_1276_v34_).cf46
            (globalState).f1 = ((d_1275_v33_)[default__.safeIndex(-856, len(d_1275_v33_))]) - (96)
            d_1277_v35_: _dafny.Array
            nw187_ = _dafny.Array(False, 17)
            d_1277_v35_ = nw187_
            index217_ = default__.safeIndex(663, (d_1277_v35_).length(0))
            (d_1277_v35_)[index217_] = d_1239_v0_
            index218_ = default__.safeIndex(663, (d_1277_v35_).length(0))
            (d_1277_v35_)[index218_] = d_1239_v0_
            d_1278_v36_: str
            d_1278_v36_ = _dafny.CodePoint('l')
            d_1279_v37_: _dafny.Array
            nw188_ = _dafny.Array(None, 18)
            nw188_[int(0)] = _dafny.CodePoint('r')
            nw188_[int(1)] = d_1278_v36_
            nw188_[int(2)] = d_1278_v36_
            nw188_[int(3)] = _dafny.CodePoint('l')
            nw188_[int(4)] = _dafny.CodePoint('k')
            nw188_[int(5)] = d_1278_v36_
            nw188_[int(6)] = d_1278_v36_
            nw188_[int(7)] = d_1278_v36_
            nw188_[int(8)] = d_1278_v36_
            nw188_[int(9)] = d_1278_v36_
            nw188_[int(10)] = d_1278_v36_
            nw188_[int(11)] = d_1278_v36_
            nw188_[int(12)] = d_1278_v36_
            nw188_[int(13)] = d_1278_v36_
            nw188_[int(14)] = d_1278_v36_
            nw188_[int(15)] = d_1278_v36_
            nw188_[int(16)] = d_1278_v36_
            nw188_[int(17)] = d_1278_v36_
            d_1279_v37_ = nw188_
            d_1280_v38_: _dafny.Seq
            d_1280_v38_ = _dafny.SeqWithoutIsStrInference([d_1279_v37_, d_1279_v37_, d_1279_v37_, d_1279_v37_, d_1279_v37_])
            d_1281_v39_: _dafny.Map
            d_1281_v39_ = _dafny.Map({d_1239_v0_: default__.fm50(d_1239_v0_, d_1272_i1_, globalState)})
            d_1282_v40_: _dafny.Set
            d_1282_v40_ = _dafny.Set({(d_1277_v35_)[default__.safeIndex(663, (d_1277_v35_).length(0))], (d_1277_v35_)[default__.safeIndex(663, (d_1277_v35_).length(0))]})
            index219_ = default__.safeIndex(663, (d_1277_v35_).length(0))
            rhs215_ = not (not((d_1239_v0_) not in (d_1281_v39_))) or (((p0).fm8(len(d_1280_v38_), not(d_1239_v0_), d_1239_v0_, globalState) if (d_1277_v35_)[default__.safeIndex(663, (d_1277_v35_).length(0))] else (d_1277_v35_)[default__.safeIndex(663, (d_1277_v35_).length(0))]))
            rhs216_ = (d_1282_v40_) | (_dafny.Set({(d_1277_v35_)[default__.safeIndex(663, (d_1277_v35_).length(0))], d_1239_v0_}))
            lhs155_ = d_1277_v35_
            lhs156_ = default__.safeIndex(663, (d_1277_v35_).length(0))
            lhs157_ = globalState
            lhs155_[lhs156_] = rhs215_
            lhs157_.f3 = rhs216_
        d_1283_v41_: str
        d_1283_v41_ = _dafny.CodePoint('x')
        d_1284_v42_: _dafny.Array
        nw189_ = _dafny.Array(None, 9)
        nw189_[int(0)] = d_1283_v41_
        nw189_[int(1)] = d_1283_v41_
        nw189_[int(2)] = d_1283_v41_
        nw189_[int(3)] = d_1283_v41_
        nw189_[int(4)] = d_1283_v41_
        nw189_[int(5)] = d_1283_v41_
        nw189_[int(6)] = d_1283_v41_
        nw189_[int(7)] = d_1283_v41_
        nw189_[int(8)] = d_1283_v41_
        d_1284_v42_ = nw189_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1284_v42_).length(0)):
            d_1285_i2_: int = guard_loop_3_
            if (True) and (((0) <= (d_1285_i2_)) and ((d_1285_i2_) < ((d_1284_v42_).length(0)))):
                (d_1284_v42_)[(d_1285_i2_)] = d_1283_v41_
        d_1286_v43_: C1
        nw190_ = C1()
        nw190_.ctor__(d_1239_v0_)
        d_1286_v43_ = nw190_
        d_1287_v44_: _dafny.Seq
        d_1287_v44_ = _dafny.SeqWithoutIsStrInference([d_1286_v43_, d_1286_v43_])
        d_1288_v45_: _dafny.Array
        nw191_ = _dafny.Array(None, 18)
        nw191_[int(0)] = d_1286_v43_
        nw191_[int(1)] = (d_1287_v44_)[default__.safeIndex(d_1264_v23_, len(d_1287_v44_))]
        nw191_[int(2)] = d_1286_v43_
        nw191_[int(3)] = d_1286_v43_
        nw191_[int(4)] = d_1286_v43_
        nw191_[int(5)] = d_1286_v43_
        nw191_[int(6)] = d_1286_v43_
        nw191_[int(7)] = d_1286_v43_
        nw191_[int(8)] = d_1286_v43_
        nw191_[int(9)] = d_1286_v43_
        nw191_[int(10)] = d_1286_v43_
        nw191_[int(11)] = d_1286_v43_
        nw191_[int(12)] = d_1286_v43_
        nw191_[int(13)] = d_1286_v43_
        nw191_[int(14)] = d_1286_v43_
        nw191_[int(15)] = d_1286_v43_
        nw191_[int(16)] = d_1286_v43_
        nw191_[int(17)] = d_1286_v43_
        d_1288_v45_ = nw191_
        d_1289_v46_: _dafny.Map
        d_1289_v46_ = _dafny.Map({default__.safeDivisionInt((d_1265_v24_)[default__.safeIndex(587, (d_1265_v24_).length(0))], 884): d_1288_v45_})
        d_1289_v46_ = (d_1289_v46_).set(d_1264_v23_, (d_1288_v45_ if d_1239_v0_ else d_1288_v45_))
        r0 = d_1284_v42_
        r1 = d_1239_v0_
        return r0, r1


class C6:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self):
        pass
        pass

    def fm22(self, p0, p1, globalState):
        return ((202 if True else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))))) != (133)

    def fm23(self, p0, p1, p2, p3, globalState):
        return (True) and ((_dafny.MultiSet([520, len(_dafny.Map({-300: False}))])).issubset(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: True}), _dafny.Map({False: True})]))])))

    def m11(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_1290_v0_: int
        d_1290_v0_ = -465
        d_1291_v1_: _dafny.Seq
        d_1291_v1_ = _dafny.SeqWithoutIsStrInference([d_1290_v0_])
        d_1292_v2_: _dafny.Seq
        d_1292_v2_ = _dafny.SeqWithoutIsStrInference([d_1290_v0_, len(d_1291_v1_)])
        d_1293_v3_: _dafny.MultiSet
        d_1293_v3_ = _dafny.MultiSet([len(d_1291_v1_), d_1290_v0_, 962])
        r0 = (d_1293_v3_).ispropersubset((d_1293_v3_) | ((_dafny.MultiSet([d_1290_v0_])).set((d_1292_v2_)[default__.safeIndex(d_1290_v0_, len(d_1292_v2_))], default__.abs(d_1290_v0_))))
        d_1294_v4_: _dafny.Seq
        d_1294_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rcvrbcdrn"))
        d_1295_v5_: _dafny.Seq
        d_1295_v5_ = _dafny.SeqWithoutIsStrInference([D3_DC10(d_1294_v4_, False, p0)])
        d_1296_v6_: D3
        d_1296_v6_ = D3_DC12((d_1295_v5_)[default__.safeIndex(d_1290_v0_, len(d_1295_v5_))])
        d_1297_i0_: int
        d_1297_i0_ = 0
        with _dafny.label("6"):
            pat_let_tv28_ = p0
            pat_let_tv29_ = p0
            pat_let_tv30_ = p0
            pat_let_tv31_ = p0
            def lambda66_(source26_):
                if source26_.is_DC10:
                    d_1303___mcc_h0_ = source26_.cf14
                    d_1304___mcc_h1_ = source26_.cf15
                    d_1305___mcc_h2_ = source26_.cf16
                    d_1306_cf16_ = d_1305___mcc_h2_
                    d_1307_cf15_ = d_1304___mcc_h1_
                    d_1308_cf14_ = d_1303___mcc_h0_
                    return False
                elif source26_.is_DC11:
                    d_1309___mcc_h3_ = source26_.cf17
                    d_1310___mcc_h4_ = source26_.cf18
                    d_1311_cf18_ = d_1310___mcc_h4_
                    d_1312_cf17_ = d_1309___mcc_h3_
                    return (pat_let_tv28_) == (True)
                elif source26_.is_DC9:
                    d_1313___mcc_h5_ = source26_.cf13
                    d_1314_cf13_ = d_1313___mcc_h5_
                    return pat_let_tv29_
                elif True:
                    d_1315___mcc_h6_ = source26_.cf19
                    d_1316_cf19_ = d_1315___mcc_h6_
                    return not (pat_let_tv30_) or (pat_let_tv31_)

            while lambda66_(d_1296_v6_):
                with _dafny.c_label("6"):
                    if (d_1297_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1297_i0_ = (d_1297_i0_) + (1)
                    d_1298_v7_: _dafny.Array
                    nw192_ = _dafny.Array(False, 18)
                    d_1298_v7_ = nw192_
                    index220_ = default__.safeIndex(144, (d_1298_v7_).length(0))
                    (d_1298_v7_)[index220_] = (d_1294_v4_) < (d_1294_v4_)
                    index221_ = default__.safeIndex(144, (d_1298_v7_).length(0))
                    (d_1298_v7_)[index221_] = p0
                    d_1299_v8_: _dafny.Seq
                    d_1299_v8_ = _dafny.SeqWithoutIsStrInference([(d_1298_v7_)[default__.safeIndex(144, (d_1298_v7_).length(0))]])
                    d_1300_v9_: _dafny.Array
                    nw193_ = _dafny.Array(D2.default()(), 11)
                    d_1300_v9_ = nw193_
                    d_1301_v10_: _dafny.Map
                    d_1301_v10_ = _dafny.Map({len((d_1299_v8_) + (d_1299_v8_)): d_1300_v9_})
                    d_1302_v11_: D6
                    d_1302_v11_ = D6_DC19(d_1301_v10_)
                    d_1301_v10_ = (d_1301_v10_) | ((d_1302_v11_).cf31)
                    index222_ = default__.safeIndex(144, (d_1298_v7_).length(0))
                    (d_1298_v7_)[index222_] = (d_1298_v7_)[default__.safeIndex(144, (d_1298_v7_).length(0))]
                    r0 = not (((d_1298_v7_)[default__.safeIndex(144, (d_1298_v7_).length(0))]) and ((d_1298_v7_)[default__.safeIndex(144, (d_1298_v7_).length(0))])) or (p0)
                    pass
            pass
        d_1317_v12_: D6
        d_1317_v12_ = D6_DC21((self).fm23(p0, d_1290_v0_, p0, d_1290_v0_, globalState), p0, p0, len(d_1291_v1_))
        source27_ = d_1317_v12_
        if source27_.is_DC20:
            d_1318___mcc_h7_ = source27_.cf32
            d_1319___mcc_h8_ = source27_.cf33
            d_1320___mcc_h9_ = source27_.cf34
            d_1321___mcc_h10_ = source27_.cf35
            d_1322_cf35_ = d_1321___mcc_h10_
            d_1323_cf34_ = d_1320___mcc_h9_
            d_1324_cf33_ = d_1319___mcc_h8_
            d_1325_cf32_ = d_1318___mcc_h7_
            d_1326_v13_: _dafny.Array
            def lambda67_(d_1327_cf33_, d_1328_p0_, d_1329_cf32_):
                def lambda68_(d_1330_i1_):
                    return D4_DC13(_dafny.Map({_dafny.SeqWithoutIsStrInference([d_1327_cf33_, d_1328_p0_]): d_1329_cf32_}))

                return lambda68_

            init36_ = lambda67_(d_1324_cf33_, p0, d_1325_cf32_)
            nw194_ = _dafny.Array(None, 9)
            for i0_36_ in range(nw194_.length(0)):
                nw194_[i0_36_] = init36_(i0_36_)
            d_1326_v13_ = nw194_
            d_1331_v14_: _dafny.Map
            d_1331_v14_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1324_cf33_]): d_1323_cf34_})
            index223_ = default__.safeIndex(663, (d_1326_v13_).length(0))
            (d_1326_v13_)[index223_] = D4_DC13(d_1331_v14_)
            d_1332_v15_: D4
            d_1332_v15_ = D4_DC13(d_1331_v14_)
            d_1333_v16_: _dafny.Array
            nw195_ = _dafny.Array(False, 21)
            d_1333_v16_ = nw195_
            d_1334_v17_: D0
            d_1334_v17_ = D0_DC2(d_1325_cf32_, d_1333_v16_)
            index224_ = default__.safeIndex(663, (d_1326_v13_).length(0))
            rhs217_ = d_1332_v15_
            rhs218_ = d_1290_v0_
            rhs219_ = (d_1291_v1_)[default__.safeIndex(d_1290_v0_, len(d_1291_v1_))]
            rhs220_ = (d_1334_v17_).cf2
            lhs158_ = d_1326_v13_
            lhs159_ = default__.safeIndex(663, (d_1326_v13_).length(0))
            lhs160_ = globalState
            lhs161_ = globalState
            lhs158_[lhs159_] = rhs217_
            lhs160_.f1 = rhs218_
            lhs161_.f7 = rhs219_
            d_1323_cf34_ = rhs220_
            d_1335_v18_: _dafny.Seq
            d_1335_v18_ = _dafny.SeqWithoutIsStrInference([d_1292_v2_, d_1291_v1_])
            d_1336_v19_: _dafny.Array
            def lambda69_(d_1337_cf35_):
                def lambda70_(d_1338_i2_):
                    return (d_1338_i2_) * (d_1337_cf35_)

                return lambda70_

            init37_ = lambda69_(d_1322_cf35_)
            nw196_ = _dafny.Array(None, 29)
            for i0_37_ in range(nw196_.length(0)):
                nw196_[i0_37_] = init37_(i0_37_)
            d_1336_v19_ = nw196_
            d_1339_v20_: _dafny.Map
            d_1339_v20_ = _dafny.Map({len(d_1335_v18_): d_1336_v19_})
            d_1340_v21_: _dafny.Map
            d_1340_v21_ = _dafny.Map({193: d_1339_v20_})
            d_1341_v22_: _dafny.Seq
            d_1341_v22_ = _dafny.SeqWithoutIsStrInference([d_1339_v20_, d_1339_v20_, d_1339_v20_, d_1339_v20_])
            d_1340_v21_ = (d_1340_v21_).set(d_1322_cf35_, (d_1339_v20_ if d_1325_cf32_ else (d_1341_v22_)[default__.safeIndex((0) - (d_1290_v0_), len(d_1341_v22_))]))
            d_1342_v23_: _dafny.Map
            d_1342_v23_ = _dafny.Map({d_1333_v16_: d_1296_v6_})
            d_1342_v23_ = (d_1342_v23_).set(d_1333_v16_, d_1296_v6_)
            d_1343_v24_: _dafny.Map
            d_1343_v24_ = _dafny.Map({d_1324_cf33_: (d_1322_cf35_) * (len(_dafny.Map({d_1290_v0_: d_1324_cf33_})))})
            d_1344_v25_: _dafny.Map
            d_1344_v25_ = _dafny.Map({d_1322_cf35_: p0})
            d_1343_v24_ = (d_1343_v24_).set(((d_1344_v25_)[993] if (993) in (d_1344_v25_) else default__.fm0(d_1290_v0_, globalState)), d_1322_cf35_)
        elif source27_.is_DC21:
            d_1345___mcc_h11_ = source27_.cf36
            d_1346___mcc_h12_ = source27_.cf37
            d_1347___mcc_h13_ = source27_.cf38
            d_1348___mcc_h14_ = source27_.cf39
            d_1349_cf39_ = d_1348___mcc_h14_
            d_1350_cf38_ = d_1347___mcc_h13_
            d_1351_cf37_ = d_1346___mcc_h12_
            d_1352_cf36_ = d_1345___mcc_h11_
            d_1353_v26_: _dafny.Seq
            d_1353_v26_ = _dafny.SeqWithoutIsStrInference([not(d_1350_cf38_)])
            d_1352_cf36_ = (((0) - ((0) - (d_1349_cf39_))) <= (d_1290_v0_) if (self).fm23(p0, d_1290_v0_, d_1350_cf38_, (default__.fm24(globalState)).cardinality, globalState) else (d_1353_v26_)[default__.safeIndex(d_1349_cf39_, len(d_1353_v26_))])
            d_1354_v27_: str
            d_1354_v27_ = _dafny.CodePoint('d')
            d_1355_v28_: _dafny.Array
            nw197_ = _dafny.Array(None, 3)
            nw197_[int(0)] = default__.fm4(d_1351_cf37_, True, d_1352_cf36_, globalState)
            nw197_[int(1)] = d_1354_v27_
            nw197_[int(2)] = d_1354_v27_
            d_1355_v28_ = nw197_
            index225_ = default__.safeIndex(230, (d_1355_v28_).length(0))
            (d_1355_v28_)[index225_] = d_1354_v27_
            index226_ = default__.safeIndex(230, (d_1355_v28_).length(0))
            rhs221_ = d_1350_cf38_
            rhs222_ = (d_1294_v4_)[default__.safeIndex(-412, len(d_1294_v4_))]
            lhs162_ = d_1355_v28_
            lhs163_ = default__.safeIndex(230, (d_1355_v28_).length(0))
            d_1351_cf37_ = rhs221_
            lhs162_[lhs163_] = rhs222_
            d_1356_v29_: _dafny.Array
            nw198_ = _dafny.Array(_dafny.Map({}), 8)
            d_1356_v29_ = nw198_
            d_1357_v30_: _dafny.Map
            d_1357_v30_ = _dafny.Map({d_1350_cf38_: d_1350_cf38_})
            index227_ = default__.safeIndex(642, (d_1356_v29_).length(0))
            (d_1356_v29_)[index227_] = d_1357_v30_
            index228_ = default__.safeIndex(642, (d_1356_v29_).length(0))
            rhs223_ = (d_1357_v30_) | ((d_1357_v30_).set(False, p0))
            rhs224_ = not (False) or (d_1352_cf36_)
            rhs225_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "echwct"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fckgecd")))
            lhs164_ = d_1356_v29_
            lhs165_ = default__.safeIndex(642, (d_1356_v29_).length(0))
            lhs164_[lhs165_] = rhs223_
            d_1351_cf37_ = rhs224_
            r0 = rhs225_
            d_1358_v31_: _dafny.Array
            nw199_ = _dafny.Array(None, 4)
            nw199_[int(0)] = d_1350_cf38_
            nw199_[int(1)] = d_1351_cf37_
            nw199_[int(2)] = p0
            nw199_[int(3)] = p0
            d_1358_v31_ = nw199_
            d_1359_v32_: _dafny.Seq
            d_1359_v32_ = _dafny.SeqWithoutIsStrInference([d_1358_v31_, d_1358_v31_])
            d_1360_v33_: _dafny.Array
            nw200_ = _dafny.Array(None, 27)
            nw200_[int(0)] = d_1358_v31_
            nw200_[int(1)] = d_1358_v31_
            nw200_[int(2)] = (d_1358_v31_ if d_1351_cf37_ else d_1358_v31_)
            nw200_[int(3)] = d_1358_v31_
            nw200_[int(4)] = d_1358_v31_
            nw200_[int(5)] = d_1358_v31_
            nw200_[int(6)] = d_1358_v31_
            nw200_[int(7)] = d_1358_v31_
            nw200_[int(8)] = d_1358_v31_
            nw200_[int(9)] = d_1358_v31_
            nw200_[int(10)] = d_1358_v31_
            nw200_[int(11)] = (d_1359_v32_)[default__.safeIndex(d_1349_cf39_, len(d_1359_v32_))]
            nw200_[int(12)] = d_1358_v31_
            nw200_[int(13)] = d_1358_v31_
            nw200_[int(14)] = d_1358_v31_
            nw200_[int(15)] = d_1358_v31_
            nw200_[int(16)] = d_1358_v31_
            nw200_[int(17)] = d_1358_v31_
            nw200_[int(18)] = d_1358_v31_
            nw200_[int(19)] = d_1358_v31_
            nw200_[int(20)] = d_1358_v31_
            nw200_[int(21)] = d_1358_v31_
            nw200_[int(22)] = d_1358_v31_
            nw200_[int(23)] = d_1358_v31_
            nw200_[int(24)] = d_1358_v31_
            nw200_[int(25)] = d_1358_v31_
            nw200_[int(26)] = d_1358_v31_
            d_1360_v33_ = nw200_
            d_1361_v34_: _dafny.Seq
            d_1361_v34_ = _dafny.SeqWithoutIsStrInference([d_1360_v33_])
            d_1362_v35_: _dafny.MultiSet
            d_1362_v35_ = _dafny.MultiSet([not(p0)])
            rhs226_ = (d_1361_v34_)[default__.safeIndex((0) - (default__.safeModuloInt(d_1290_v0_, d_1349_cf39_)), len(d_1361_v34_))]
            rhs227_ = p0
            rhs228_ = len((((d_1294_v4_) + ((d_1294_v4_ if d_1351_cf37_ else d_1294_v4_))).set(default__.safeIndex(((d_1362_v35_).intersection(_dafny.MultiSet([d_1351_cf37_]))).cardinality, len((d_1294_v4_) + ((d_1294_v4_ if d_1351_cf37_ else d_1294_v4_)))), _dafny.CodePoint('m'))).set(default__.safeIndex(d_1290_v0_, len(((d_1294_v4_) + ((d_1294_v4_ if d_1351_cf37_ else d_1294_v4_))).set(default__.safeIndex(((d_1362_v35_).intersection(_dafny.MultiSet([d_1351_cf37_]))).cardinality, len((d_1294_v4_) + ((d_1294_v4_ if d_1351_cf37_ else d_1294_v4_)))), _dafny.CodePoint('m')))), (d_1355_v28_)[default__.safeIndex(230, (d_1355_v28_).length(0))]))
            lhs166_ = globalState
            d_1360_v33_ = rhs226_
            d_1351_cf37_ = rhs227_
            lhs166_.f1 = rhs228_
        elif True:
            d_1363___mcc_h15_ = source27_.cf31
            d_1364_cf31_ = d_1363___mcc_h15_
            d_1365_v36_: D0
            d_1365_v36_ = D0_DC1(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1366_i3_ in range(default__.abs(-533))]))
            d_1367_v37_: _dafny.Map
            d_1367_v37_ = _dafny.Map({False: d_1365_v36_})
            d_1368_v38_: _dafny.Map
            d_1368_v38_ = _dafny.Map({False: d_1367_v37_})
            pat_let_tv32_ = d_1294_v4_
            def iife73_(_pat_let23_0):
                def iife74_(d_1369_dt__update__tmp_h0_):
                    def iife75_(_pat_let24_0):
                        def iife76_(d_1370_dt__update_hcf1_h0_):
                            return D0_DC1(d_1370_dt__update_hcf1_h0_)
                        return iife76_(_pat_let24_0)
                    return iife75_(pat_let_tv32_)
                return iife74_(_pat_let23_0)
            if ((_dafny.Map({p0: iife73_(d_1365_v36_)})) | (_dafny.Map({p0: d_1365_v36_}))) != (((((d_1368_v38_)[p0] if (p0) in (d_1368_v38_) else d_1367_v37_)).set(p0, d_1365_v36_) if p0 else d_1367_v37_)):
                d_1371_v39_: _dafny.Map
                d_1371_v39_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1372_i4_ in range(default__.abs(82))])) + (d_1294_v4_): d_1290_v0_})
                d_1371_v39_ = (d_1371_v39_).set(d_1294_v4_, default__.fm1(globalState))
                d_1373_v40_: _dafny.Array
                nw201_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_1373_v40_ = nw201_
                d_1374_v41_: _dafny.Array
                nw202_ = _dafny.Array(int(0), 11)
                d_1374_v41_ = nw202_
                index229_ = default__.safeIndex(902, (d_1373_v40_).length(0))
                (d_1373_v40_)[index229_] = d_1374_v41_
                d_1375_v42_: _dafny.Map
                d_1375_v42_ = _dafny.Map({987: d_1374_v41_})
                d_1376_v43_: _dafny.Seq
                d_1376_v43_ = _dafny.SeqWithoutIsStrInference([d_1374_v41_])
                index230_ = default__.safeIndex(902, (d_1373_v40_).length(0))
                (d_1373_v40_)[index230_] = ((d_1375_v42_)[(d_1290_v0_) * (d_1290_v0_)] if ((d_1290_v0_) * (d_1290_v0_)) in (d_1375_v42_) else (d_1376_v43_)[default__.safeIndex(-802, len(d_1376_v43_))])
                d_1377_v44_: _dafny.Array
                nw203_ = _dafny.Array(False, 19)
                d_1377_v44_ = nw203_
                d_1377_v44_ = d_1377_v44_
                d_1378_v45_: _dafny.Map
                d_1378_v45_ = _dafny.Map({p0: d_1290_v0_})
                d_1379_v46_: D2
                d_1379_v46_ = D2_DC7(len(d_1378_v45_), p0, d_1290_v0_)
                r0 = not (p0) or ((d_1379_v46_).cf10)
                d_1380_v47_: _dafny.Seq
                d_1380_v47_ = _dafny.SeqWithoutIsStrInference([True])
                arr2_ = (d_1373_v40_)[default__.safeIndex(902, (d_1373_v40_).length(0))]
                index231_ = default__.safeIndex(870, ((d_1373_v40_)[default__.safeIndex(902, (d_1373_v40_).length(0))]).length(0))
                arr2_[index231_] = len(d_1380_v47_)
                d_1381_v48_: _dafny.Set
                d_1381_v48_ = _dafny.Set({d_1380_v47_, d_1380_v47_, (d_1380_v47_) + (_dafny.SeqWithoutIsStrInference([p0]))})
                arr3_ = (d_1373_v40_)[default__.safeIndex(902, (d_1373_v40_).length(0))]
                index232_ = default__.safeIndex(870, ((d_1373_v40_)[default__.safeIndex(902, (d_1373_v40_).length(0))]).length(0))
                arr3_[index232_] = (0) - ((0) - (len(d_1381_v48_)))
            elif True:
                d_1382_v49_: _dafny.Array
                nw204_ = _dafny.Array(_dafny.Seq({}), 28)
                d_1382_v49_ = nw204_
                d_1383_v50_: _dafny.Seq
                d_1383_v50_ = _dafny.SeqWithoutIsStrInference([p0])
                index233_ = default__.safeIndex(715, (d_1382_v49_).length(0))
                (d_1382_v49_)[index233_] = (d_1383_v50_) + (_dafny.SeqWithoutIsStrInference([p0, p0]))
                index234_ = default__.safeIndex(715, (d_1382_v49_).length(0))
                (d_1382_v49_)[index234_] = d_1383_v50_
                d_1384_v51_: _dafny.Seq
                d_1384_v51_ = _dafny.SeqWithoutIsStrInference([d_1294_v4_, default__.fm25(410, True, (_dafny.MultiSet([p0])).cardinality, globalState)])
                d_1385_v52_: _dafny.Set
                d_1385_v52_ = _dafny.Set({d_1290_v0_, d_1290_v0_, d_1290_v0_, d_1290_v0_})
                rhs229_ = p0
                rhs230_ = default__.safeModuloInt((0) - (d_1290_v0_), d_1290_v0_)
                rhs231_ = _dafny.SeqWithoutIsStrInference([d_1294_v4_ for d_1386_i5_ in range(default__.abs(395))])
                rhs232_ = ((d_1294_v4_) != (d_1294_v4_)) == (not(p0))
                rhs233_ = (_dafny.Set({d_1290_v0_, 648, 510})).isdisjoint((d_1385_v52_) - (default__.fm26(globalState)))
                lhs167_ = globalState
                r0 = rhs229_
                lhs167_.f0 = rhs230_
                d_1384_v51_ = rhs231_
                r0 = rhs232_
                r0 = rhs233_
                d_1387_v53_: D4
                d_1387_v53_ = D4_DC14(d_1290_v0_, d_1290_v0_)
                (self).m12(default__.safeModuloInt(d_1290_v0_, (d_1387_v53_).cf21), globalState)
                d_1388_v54_: _dafny.Map
                d_1388_v54_ = _dafny.Map({True: _dafny.SeqWithoutIsStrInference([p0, p0])})
                d_1383_v50_ = (((d_1388_v54_)[p0] if (p0) in (d_1388_v54_) else (d_1382_v49_)[default__.safeIndex(715, (d_1382_v49_).length(0))])) + (_dafny.SeqWithoutIsStrInference([False]))
                r0 = not (p0) or (p0)
            (self).m12(d_1290_v0_, globalState)
            d_1389_v55_: D6
            d_1389_v55_ = D6_DC20(p0, p0, p0, (((d_1293_v3_)[d_1290_v0_] if (d_1290_v0_) in (d_1293_v3_) else d_1290_v0_)) - (d_1290_v0_))
            source28_ = d_1389_v55_
            if source28_.is_DC20:
                d_1390___mcc_h16_ = source28_.cf32
                d_1391___mcc_h17_ = source28_.cf33
                d_1392___mcc_h18_ = source28_.cf34
                d_1393___mcc_h19_ = source28_.cf35
                d_1394_cf35_ = d_1393___mcc_h19_
                d_1395_cf34_ = d_1392___mcc_h18_
                d_1396_cf33_ = d_1391___mcc_h17_
                d_1397_cf32_ = d_1390___mcc_h16_
                d_1398_v56_: str
                d_1398_v56_ = _dafny.CodePoint('v')
                d_1399_v57_: _dafny.Map
                d_1399_v57_ = _dafny.Map({(d_1291_v1_) == (d_1292_v2_): d_1398_v56_})
                d_1399_v57_ = (d_1399_v57_).set(True, _dafny.CodePoint('b'))
                d_1400_v58_: C3
                nw205_ = C3()
                nw205_.ctor__(p0, p0)
                d_1400_v58_ = nw205_
                d_1401_v59_: _dafny.Array
                nw206_ = _dafny.Array(D4.default()(), 21)
                d_1401_v59_ = nw206_
                d_1402_v60_: D4
                d_1402_v60_ = D4_DC15((d_1291_v1_)[default__.safeIndex(d_1394_cf35_, len(d_1291_v1_))], d_1293_v3_, d_1398_v56_, d_1394_cf35_)
                index235_ = default__.safeIndex(79, (d_1401_v59_).length(0))
                (d_1401_v59_)[index235_] = d_1402_v60_
                index236_ = default__.safeIndex(79, (d_1401_v59_).length(0))
                (d_1401_v59_)[index236_] = d_1402_v60_
                d_1294_v4_ = d_1294_v4_
            elif source28_.is_DC21:
                d_1403___mcc_h20_ = source28_.cf36
                d_1404___mcc_h21_ = source28_.cf37
                d_1405___mcc_h22_ = source28_.cf38
                d_1406___mcc_h23_ = source28_.cf39
                d_1407_cf39_ = d_1406___mcc_h23_
                d_1408_cf38_ = d_1405___mcc_h22_
                d_1409_cf37_ = d_1404___mcc_h21_
                d_1410_cf36_ = d_1403___mcc_h20_
                d_1410_cf36_ = p0
                d_1411_v61_: _dafny.Seq
                d_1411_v61_ = _dafny.SeqWithoutIsStrInference([d_1292_v2_])
                d_1412_v62_: D13
                d_1412_v62_ = D13_DC39(d_1290_v0_, p0)
                d_1413_v63_: _dafny.Array
                nw207_ = _dafny.Array(None, 14)
                nw207_[int(0)] = d_1292_v2_
                nw207_[int(1)] = (_dafny.SeqWithoutIsStrInference([d_1407_cf39_, d_1407_cf39_, (0) - (d_1290_v0_)]) if p0 else d_1292_v2_)
                nw207_[int(2)] = (d_1411_v61_)[default__.safeIndex(len(d_1294_v4_), len(d_1411_v61_))]
                nw207_[int(3)] = d_1291_v1_
                nw207_[int(4)] = d_1291_v1_
                nw207_[int(5)] = ((_dafny.SeqWithoutIsStrInference([d_1407_cf39_ for d_1414_i6_ in range(default__.abs(-211))])).set(default__.safeIndex(d_1290_v0_, len(_dafny.SeqWithoutIsStrInference([d_1407_cf39_ for d_1415_i6_ in range(default__.abs(-211))]))), d_1407_cf39_)) + (d_1292_v2_)
                nw207_[int(6)] = d_1292_v2_
                nw207_[int(7)] = d_1291_v1_
                nw207_[int(8)] = ((d_1292_v2_).set(default__.safeIndex(752, len(d_1292_v2_)), d_1290_v0_)) + (_dafny.SeqWithoutIsStrInference([774 for d_1416_i7_ in range(default__.abs(-687))]))
                nw207_[int(9)] = _dafny.SeqWithoutIsStrInference([default__.fm1(globalState)])
                nw207_[int(10)] = (_dafny.SeqWithoutIsStrInference([(0) - ((0) - (-811))])) + (d_1292_v2_)
                nw207_[int(11)] = d_1292_v2_
                nw207_[int(12)] = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1294_v4_)) for d_1417_i8_ in range(default__.abs(763))])
                nw207_[int(13)] = _dafny.SeqWithoutIsStrInference([(d_1412_v62_).cf65])
                d_1413_v63_ = nw207_
                d_1413_v63_ = d_1413_v63_
                d_1418_v64_: _dafny.Seq
                d_1418_v64_ = _dafny.SeqWithoutIsStrInference([d_1409_cf37_])
                d_1419_v65_: _dafny.Map
                d_1419_v65_ = _dafny.Map({d_1293_v3_: (d_1418_v64_) + (d_1418_v64_)})
                d_1419_v65_ = d_1419_v65_
                d_1409_cf37_ = not(d_1410_cf36_)
            elif True:
                d_1420___mcc_h24_ = source28_.cf31
                d_1421_cf31_ = d_1420___mcc_h24_
                r0 = p0
                d_1422_v66_: _dafny.Array
                def lambda71_(d_1423_i9_):
                    return default__.safeModuloInt(d_1423_i9_, 22)

                init38_ = lambda71_
                nw208_ = _dafny.Array(None, 21)
                for i0_38_ in range(nw208_.length(0)):
                    nw208_[i0_38_] = init38_(i0_38_)
                d_1422_v66_ = nw208_
                d_1424_v67_: _dafny.Map
                d_1424_v67_ = _dafny.Map({d_1293_v3_: d_1422_v66_})
                d_1424_v67_ = (d_1424_v67_).set(_dafny.MultiSet(d_1291_v1_), d_1422_v66_)
                d_1425_v68_: _dafny.Map
                d_1425_v68_ = _dafny.Map({not(p0): default__.fm1(globalState)})
                d_1425_v68_ = (d_1425_v68_).set(p0, d_1290_v0_)
                r0 = p0
            r0 = p0
        d_1426_v69_: _dafny.Array
        def lambda72_(d_1427_p0_):
            def lambda73_(d_1428_i10_):
                return d_1427_p0_

            return lambda73_

        init39_ = lambda72_(p0)
        nw209_ = _dafny.Array(None, 27)
        for i0_39_ in range(nw209_.length(0)):
            nw209_[i0_39_] = init39_(i0_39_)
        d_1426_v69_ = nw209_
        index237_ = default__.safeIndex(201, (d_1426_v69_).length(0))
        (d_1426_v69_)[index237_] = p0
        d_1429_v70_: D2
        d_1429_v70_ = D2_DC7(d_1290_v0_, p0, d_1290_v0_)
        d_1430_v71_: D2
        d_1430_v71_ = D2_DC8(d_1429_v70_)
        pat_let_tv33_ = p0
        pat_let_tv34_ = p0
        pat_let_tv35_ = p0
        pat_let_tv36_ = p0
        index238_ = default__.safeIndex(201, (d_1426_v69_).length(0))
        def lambda74_(source29_):
            if source29_.is_DC7:
                d_1431___mcc_h25_ = source29_.cf9
                d_1432___mcc_h26_ = source29_.cf10
                d_1433___mcc_h27_ = source29_.cf11
                d_1434_cf11_ = d_1433___mcc_h27_
                d_1435_cf10_ = d_1432___mcc_h26_
                d_1436_cf9_ = d_1431___mcc_h25_
                return pat_let_tv33_
            elif source29_.is_DC6:
                d_1437___mcc_h28_ = source29_.cf8
                d_1438_cf8_ = d_1437___mcc_h28_
                return (pat_let_tv34_) or (pat_let_tv35_)
            elif True:
                d_1439___mcc_h29_ = source29_.cf12
                d_1440_cf12_ = d_1439___mcc_h29_
                return pat_let_tv36_

        (d_1426_v69_)[index238_] = lambda74_(d_1430_v71_)
        d_1441_i11_: int
        d_1441_i11_ = 0
        with _dafny.label("7"):
            while (d_1426_v69_)[default__.safeIndex(201, (d_1426_v69_).length(0))]:
                with _dafny.c_label("7"):
                    if (d_1441_i11_) >= (100):
                        raise _dafny.Break("7")
                    d_1441_i11_ = (d_1441_i11_) + (1)
                    d_1442_v72_: _dafny.Set
                    d_1442_v72_ = _dafny.Set({d_1290_v0_, d_1290_v0_})
                    d_1443_v73_: _dafny.Map
                    d_1443_v73_ = _dafny.Map({d_1290_v0_: p0})
                    d_1444_v75_: _dafny.Seq
                    def iife77_():
                        coll27_ = _dafny.Set()
                        compr_27_: int
                        for compr_27_ in (d_1443_v73_).keys.Elements:
                            d_1445_v74_: int = compr_27_
                            if (d_1445_v74_) in (d_1443_v73_):
                                coll27_ = coll27_.union(_dafny.Set([(d_1445_v74_) + (335)]))
                        return _dafny.Set(coll27_)
                    d_1444_v75_ = _dafny.SeqWithoutIsStrInference([d_1442_v72_, d_1442_v72_, (iife77_()
                    ).intersection(d_1442_v72_)])
                    d_1446_v76_: _dafny.Map
                    d_1446_v76_ = _dafny.Map({p0: False})
                    d_1444_v75_ = (d_1444_v75_ if ((d_1446_v76_)[(d_1426_v69_)[default__.safeIndex(201, (d_1426_v69_).length(0))]] if ((d_1426_v69_)[default__.safeIndex(201, (d_1426_v69_).length(0))]) in (d_1446_v76_) else (d_1426_v69_)[default__.safeIndex(201, (d_1426_v69_).length(0))]) else d_1444_v75_)
                    r0 = (d_1294_v4_) != (d_1294_v4_)
                    d_1447_v77_: C1
                    nw210_ = C1()
                    nw210_.ctor__(((d_1426_v69_)[default__.safeIndex(201, (d_1426_v69_).length(0))]) == ((d_1426_v69_)[default__.safeIndex(201, (d_1426_v69_).length(0))]))
                    d_1447_v77_ = nw210_
                    d_1448_v78_: _dafny.Seq
                    d_1448_v78_ = _dafny.SeqWithoutIsStrInference([False])
                    d_1449_v79_: _dafny.MultiSet
                    d_1449_v79_ = _dafny.MultiSet([(d_1448_v78_)[default__.safeIndex(d_1290_v0_, len(d_1448_v78_))], (len(d_1294_v4_)) in (d_1443_v73_)])
                    index239_ = default__.safeIndex(201, (d_1426_v69_).length(0))
                    rhs234_ = default__.fm41(globalState)
                    rhs235_ = len((d_1443_v73_) | (d_1443_v73_))
                    rhs236_ = (d_1426_v69_)[default__.safeIndex(201, (d_1426_v69_).length(0))]
                    rhs237_ = (d_1442_v72_).ispropersubset(d_1442_v72_)
                    lhs168_ = d_1426_v69_
                    lhs169_ = default__.safeIndex(201, (d_1426_v69_).length(0))
                    d_1449_v79_ = rhs234_
                    d_1290_v0_ = rhs235_
                    lhs168_[lhs169_] = rhs236_
                    r0 = rhs237_
                    pass
            pass
        index240_ = default__.safeIndex(201, (d_1426_v69_).length(0))
        index241_ = default__.safeIndex(201, (d_1426_v69_).length(0))
        rhs238_ = default__.fm0(d_1290_v0_, globalState)
        rhs239_ = p0
        lhs170_ = d_1426_v69_
        lhs171_ = default__.safeIndex(201, (d_1426_v69_).length(0))
        lhs172_ = d_1426_v69_
        lhs173_ = default__.safeIndex(201, (d_1426_v69_).length(0))
        lhs170_[lhs171_] = rhs238_
        lhs172_[lhs173_] = rhs239_
        r0 = not((self).fm22(default__.safeDivisionInt(d_1290_v0_, 739), p0, globalState))
        r1 = d_1290_v0_
        d_1450_v80_: _dafny.Array
        nw211_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
        d_1450_v80_ = nw211_
        r2 = d_1450_v80_
        return r0, r1, r2

    def m12(self, p0, globalState):
        d_1451_v0_: _dafny.Array
        nw212_ = _dafny.Array(False, 22)
        d_1451_v0_ = nw212_
        d_1452_v1_: bool
        d_1452_v1_ = False
        index242_ = default__.safeIndex(395, (d_1451_v0_).length(0))
        (d_1451_v0_)[index242_] = d_1452_v1_
        d_1453_v2_: C0
        nw213_ = C0()
        nw213_.ctor__()
        d_1453_v2_ = nw213_
        d_1454_v3_: _dafny.MultiSet
        d_1454_v3_ = _dafny.MultiSet([d_1453_v2_, d_1453_v2_])
        d_1455_v4_: _dafny.Map
        d_1455_v4_ = _dafny.Map({d_1452_v1_: _dafny.MultiSet([d_1453_v2_])})
        index243_ = default__.safeIndex(395, (d_1451_v0_).length(0))
        (d_1451_v0_)[index243_] = (d_1454_v3_) != (((d_1455_v4_)[d_1452_v1_] if (d_1452_v1_) in (d_1455_v4_) else d_1454_v3_))
        d_1456_v5_: _dafny.Seq
        d_1457_v6_: _dafny.Seq
        out47_: _dafny.Seq
        out48_: _dafny.Seq
        out47_, out48_ = default__.m0((p0) - (p0), d_1451_v0_, globalState)
        d_1456_v5_ = out47_
        d_1457_v6_ = out48_
        d_1458_v7_: D1
        d_1458_v7_ = D1_DC3(p0)
        source30_ = d_1458_v7_
        if source30_.is_DC4:
            d_1459___mcc_h0_ = source30_.cf5
            d_1460___mcc_h1_ = source30_.cf6
            d_1461_cf6_ = d_1460___mcc_h1_
            d_1462_cf5_ = d_1459___mcc_h0_
            d_1463_v8_: _dafny.Seq
            d_1463_v8_ = _dafny.SeqWithoutIsStrInference([d_1451_v0_])
            d_1464_v9_: _dafny.Map
            d_1464_v9_ = _dafny.Map({d_1462_cf5_: (d_1463_v8_)[default__.safeIndex(p0, len(d_1463_v8_))]})
            d_1464_v9_ = _dafny.Map({d_1462_cf5_: d_1451_v0_})
            d_1465_v10_: D5
            d_1465_v10_ = D5_DC16(default__.fm41(globalState))
            d_1465_v10_ = d_1465_v10_
            d_1466_v11_: _dafny.Map
            d_1466_v11_ = _dafny.Map({(d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]: (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]})
            d_1467_v12_: _dafny.Set
            d_1467_v12_ = _dafny.Set({p0})
            d_1468_v13_: _dafny.Seq
            d_1468_v13_ = _dafny.SeqWithoutIsStrInference([d_1467_v12_])
            d_1469_v14_: _dafny.Set
            d_1469_v14_ = _dafny.Set({len(d_1468_v13_)})
            d_1470_v15_: _dafny.MultiSet
            d_1470_v15_ = _dafny.MultiSet([d_1467_v12_])
            d_1471_v16_: _dafny.Seq
            d_1471_v16_ = _dafny.SeqWithoutIsStrInference([len(d_1466_v11_), p0, p0, (d_1470_v15_).cardinality])
            (globalState).f7 = ((d_1471_v16_)[default__.safeIndex(p0, len(d_1471_v16_))]) * (d_1462_cf5_)
            def iife78_():
                coll28_ = _dafny.Set()
                compr_28_: int
                for compr_28_ in _dafny.IntegerRange(334, 780):
                    d_1472_v17_: int = compr_28_
                    if ((334) <= (d_1472_v17_)) and ((d_1472_v17_) < (780)):
                        coll28_ = coll28_.union(_dafny.Set([(d_1472_v17_) * (84)]))
                return _dafny.Set(coll28_)
            (globalState).f7 = len((iife78_()
             if False else _dafny.Set({p0})))
        elif source30_.is_DC3:
            d_1473___mcc_h2_ = source30_.cf4
            d_1474_cf4_ = d_1473___mcc_h2_
            d_1475_v18_: _dafny.Array
            nw214_ = _dafny.Array(_dafny.MultiSet({}), 9)
            d_1475_v18_ = nw214_
            d_1476_v19_: C4
            nw215_ = C4()
            nw215_.ctor__()
            d_1476_v19_ = nw215_
            index244_ = default__.safeIndex(277, (d_1475_v18_).length(0))
            (d_1475_v18_)[index244_] = _dafny.MultiSet([d_1476_v19_])
            d_1477_v20_: _dafny.MultiSet
            d_1477_v20_ = _dafny.MultiSet([d_1476_v19_, d_1476_v19_, (d_1476_v19_ if d_1452_v1_ else d_1476_v19_)])
            index245_ = default__.safeIndex(277, (d_1475_v18_).length(0))
            rhs240_ = len(d_1456_v5_)
            rhs241_ = d_1477_v20_
            lhs174_ = d_1475_v18_
            lhs175_ = default__.safeIndex(277, (d_1475_v18_).length(0))
            d_1474_cf4_ = rhs240_
            lhs174_[lhs175_] = rhs241_
            index246_ = default__.safeIndex(395, (d_1451_v0_).length(0))
            (d_1451_v0_)[index246_] = not(d_1452_v1_)
            d_1478_v21_: str
            d_1478_v21_ = _dafny.CodePoint('b')
            d_1479_v22_: D14
            d_1479_v22_ = D14_DC43(d_1478_v21_, p0, p0, d_1474_cf4_)
            d_1480_v23_: _dafny.Set
            d_1480_v23_ = _dafny.Set({D14_DC43(d_1478_v21_, d_1474_cf4_, len(d_1457_v6_), d_1474_cf4_), d_1479_v22_})
            d_1481_v24_: _dafny.Map
            d_1481_v24_ = _dafny.Map({(d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]: d_1452_v1_})
            d_1482_v25_: D3
            d_1482_v25_ = D3_DC11(d_1474_cf4_, d_1481_v24_)
            d_1483_v26_: _dafny.Array
            nw216_ = _dafny.Array(None, 4)
            nw216_[int(0)] = (d_1453_v2_).fm36(globalState)
            nw216_[int(1)] = len(d_1480_v23_)
            nw216_[int(2)] = (d_1482_v25_).cf17
            nw216_[int(3)] = d_1474_cf4_
            d_1483_v26_ = nw216_
            d_1484_v27_: _dafny.Seq
            d_1484_v27_ = _dafny.SeqWithoutIsStrInference([len(d_1457_v6_)])
            d_1485_v28_: _dafny.Map
            d_1485_v28_ = _dafny.Map({d_1478_v21_: d_1474_cf4_})
            d_1486_v29_: _dafny.Map
            d_1486_v29_ = _dafny.Map({p0: d_1474_cf4_})
            nw217_ = _dafny.Array(None, 13)
            nw217_[int(0)] = (d_1474_cf4_) - (d_1474_cf4_)
            nw217_[int(1)] = 320
            nw217_[int(2)] = p0
            nw217_[int(3)] = default__.safeDivisionInt(d_1474_cf4_, d_1474_cf4_)
            nw217_[int(4)] = default__.safeModuloInt(len(d_1484_v27_), p0)
            nw217_[int(5)] = p0
            nw217_[int(6)] = d_1474_cf4_
            nw217_[int(7)] = 937
            nw217_[int(8)] = (0) - (len(d_1457_v6_))
            nw217_[int(9)] = default__.safeModuloInt(p0, len(d_1485_v28_))
            nw217_[int(10)] = d_1474_cf4_
            nw217_[int(11)] = d_1474_cf4_
            nw217_[int(12)] = len(d_1486_v29_)
            d_1483_v26_ = nw217_
            d_1487_v30_: _dafny.Array
            nw218_ = _dafny.Array(_dafny.Seq({}), 3)
            d_1487_v30_ = nw218_
            d_1488_v31_: _dafny.Seq
            d_1488_v31_ = _dafny.SeqWithoutIsStrInference([d_1452_v1_])
            index247_ = default__.safeIndex(766, (d_1487_v30_).length(0))
            (d_1487_v30_)[index247_] = (d_1488_v31_) + (_dafny.SeqWithoutIsStrInference([d_1452_v1_]))
            index248_ = default__.safeIndex(766, (d_1487_v30_).length(0))
            (d_1487_v30_)[index248_] = (d_1488_v31_) + ((d_1488_v31_) + (_dafny.SeqWithoutIsStrInference([d_1452_v1_])))
        elif True:
            d_1489___mcc_h3_ = source30_.cf7
            d_1490_cf7_ = d_1489___mcc_h3_
            if (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]:
                (globalState).f1 = (0) - (((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1491_i0_ in range(default__.abs(906))])))) * (default__.safeModuloInt(p0, p0)))
                (globalState).f7 = 740
                d_1452_v1_ = (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]
                d_1492_v32_: _dafny.Array
                nw219_ = _dafny.Array(int(0), 16)
                d_1492_v32_ = nw219_
                d_1493_v33_: _dafny.Seq
                d_1493_v33_ = _dafny.SeqWithoutIsStrInference([d_1452_v1_, (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]])
                index249_ = default__.safeIndex(301, (d_1492_v32_).length(0))
                (d_1492_v32_)[index249_] = len(d_1493_v33_)
                d_1494_v34_: _dafny.Map
                d_1494_v34_ = _dafny.Map({p0: p0})
                d_1495_v35_: _dafny.Map
                d_1495_v35_ = _dafny.Map({d_1494_v34_: 888})
                d_1496_v36_: _dafny.Map
                d_1496_v36_ = _dafny.Map({((d_1495_v35_)[d_1494_v34_] if (d_1494_v34_) in (d_1495_v35_) else (_dafny.MultiSet([len(d_1457_v6_)])).cardinality): p0})
                d_1497_v37_: _dafny.MultiSet
                d_1497_v37_ = _dafny.MultiSet([True, d_1452_v1_])
                index250_ = default__.safeIndex(301, (d_1492_v32_).length(0))
                index251_ = default__.safeIndex(395, (d_1451_v0_).length(0))
                rhs242_ = (d_1453_v2_).fm37((d_1457_v6_) + (d_1456_v5_), ((d_1496_v36_)[p0] if (p0) in (d_1496_v36_) else p0), (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))], (d_1452_v1_) in (d_1497_v37_), globalState)
                rhs243_ = p0
                rhs244_ = d_1456_v5_
                rhs245_ = (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]
                lhs176_ = d_1492_v32_
                lhs177_ = default__.safeIndex(301, (d_1492_v32_).length(0))
                lhs178_ = d_1451_v0_
                lhs179_ = default__.safeIndex(395, (d_1451_v0_).length(0))
                d_1452_v1_ = rhs242_
                lhs176_[lhs177_] = rhs243_
                d_1457_v6_ = rhs244_
                lhs178_[lhs179_] = rhs245_
                d_1498_v38_: str
                d_1498_v38_ = _dafny.CodePoint('q')
                d_1498_v38_ = d_1498_v38_
            elif True:
                d_1499_v40_: _dafny.Set
                def iife79_():
                    coll29_ = _dafny.Map()
                    compr_29_: int
                    for compr_29_ in _dafny.IntegerRange(-89, 883):
                        d_1500_v39_: int = compr_29_
                        if ((-89) <= (d_1500_v39_)) and ((d_1500_v39_) < (883)):
                            coll29_[(d_1500_v39_) - (p0)] = p0
                    return _dafny.Map(coll29_)
                d_1499_v40_ = _dafny.Set({iife79_()
                , _dafny.Map({p0: p0})})
                d_1501_v41_: _dafny.Map
                d_1501_v41_ = _dafny.Map({p0: p0})
                d_1502_v43_: _dafny.Seq
                d_1502_v43_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0 for d_1503_i1_ in range(default__.abs(134))])])
                d_1504_v44_: _dafny.Seq
                d_1504_v44_ = _dafny.SeqWithoutIsStrInference([len(d_1502_v43_)])
                def iife80_():
                    coll30_ = _dafny.Map()
                    compr_30_: int
                    for compr_30_ in (d_1504_v44_).Elements:
                        d_1505_v42_: int = compr_30_
                        if (d_1505_v42_) in (d_1504_v44_):
                            coll30_[(d_1505_v42_) + (p0)] = 325
                    return _dafny.Map(coll30_)
                d_1499_v40_ = _dafny.Set({(d_1501_v41_) | ((iife80_()
                ).set(p0, p0)), d_1501_v41_})
                d_1506_v45_: _dafny.Array
                def lambda75_(d_1507_p0_):
                    def lambda76_(d_1508_i2_):
                        return default__.safeDivisionInt(d_1508_i2_, d_1507_p0_)

                    return lambda76_

                init40_ = lambda75_(p0)
                nw220_ = _dafny.Array(None, 9)
                for i0_40_ in range(nw220_.length(0)):
                    nw220_[i0_40_] = init40_(i0_40_)
                d_1506_v45_ = nw220_
                d_1509_v46_: _dafny.Array
                nw221_ = _dafny.Array(None, 15)
                nw221_[int(0)] = d_1506_v45_
                nw221_[int(1)] = d_1506_v45_
                nw221_[int(2)] = d_1506_v45_
                nw221_[int(3)] = d_1506_v45_
                nw221_[int(4)] = d_1506_v45_
                nw221_[int(5)] = d_1506_v45_
                nw221_[int(6)] = d_1506_v45_
                nw221_[int(7)] = d_1506_v45_
                nw221_[int(8)] = (d_1506_v45_ if not((d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]) else d_1506_v45_)
                nw221_[int(9)] = d_1506_v45_
                nw221_[int(10)] = d_1506_v45_
                nw221_[int(11)] = d_1506_v45_
                nw221_[int(12)] = d_1506_v45_
                nw221_[int(13)] = d_1506_v45_
                nw221_[int(14)] = d_1506_v45_
                d_1509_v46_ = nw221_
                index252_ = default__.safeIndex(781, (d_1509_v46_).length(0))
                (d_1509_v46_)[index252_] = d_1506_v45_
                d_1510_v47_: _dafny.MultiSet
                d_1510_v47_ = _dafny.MultiSet([p0])
                d_1511_v48_: _dafny.Set
                d_1511_v48_ = _dafny.Set({(0) - (p0), p0})
                index253_ = default__.safeIndex(395, (d_1451_v0_).length(0))
                index254_ = default__.safeIndex(781, (d_1509_v46_).length(0))
                rhs246_ = (((d_1510_v47_) - (_dafny.MultiSet([p0]))) | (d_1510_v47_)).cardinality
                rhs247_ = (p0) < ((0) - (len(d_1511_v48_)))
                rhs248_ = p0
                rhs249_ = d_1506_v45_
                lhs180_ = globalState
                lhs181_ = d_1451_v0_
                lhs182_ = default__.safeIndex(395, (d_1451_v0_).length(0))
                lhs183_ = globalState
                lhs184_ = d_1509_v46_
                lhs185_ = default__.safeIndex(781, (d_1509_v46_).length(0))
                lhs180_.f7 = rhs246_
                lhs181_[lhs182_] = rhs247_
                lhs183_.f0 = rhs248_
                lhs184_[lhs185_] = rhs249_
                index255_ = default__.safeIndex(395, (d_1451_v0_).length(0))
                (d_1451_v0_)[index255_] = (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]
                d_1452_v1_ = d_1452_v1_
                index256_ = default__.safeIndex(730, (d_1506_v45_).length(0))
                (d_1506_v45_)[index256_] = default__.safeModuloInt(p0, p0)
                index257_ = default__.safeIndex(730, (d_1506_v45_).length(0))
                (d_1506_v45_)[index257_] = p0
            d_1456_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhffhwbq"))
            d_1512_v49_: _dafny.Map
            d_1512_v49_ = _dafny.Map({838: _dafny.SeqWithoutIsStrInference([p0])})
            d_1513_v50_: _dafny.Map
            d_1513_v50_ = _dafny.Map({len(d_1512_v49_): (d_1453_v2_).fm36(globalState)})
            d_1514_v51_: _dafny.MultiSet
            d_1514_v51_ = _dafny.MultiSet([p0])
            d_1515_v52_: _dafny.Map
            d_1515_v52_ = _dafny.Map({(d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]: (d_1514_v51_).cardinality})
            d_1516_v53_: _dafny.Map
            d_1516_v53_ = _dafny.Map({d_1515_v52_: 647})
            d_1517_v54_: _dafny.Map
            d_1517_v54_ = _dafny.Map({((d_1513_v50_)[p0] if (p0) in (d_1513_v50_) else p0): len(d_1516_v53_)})
            d_1518_v55_: _dafny.Array
            nw222_ = _dafny.Array(None, 14)
            d_1518_v55_ = nw222_
            d_1519_v56_: _dafny.Map
            d_1519_v56_ = _dafny.Map({d_1452_v1_: d_1518_v55_})
            d_1520_v57_: str
            d_1520_v57_ = _dafny.CodePoint('n')
            d_1521_v58_: _dafny.Map
            d_1521_v58_ = _dafny.Map({d_1452_v1_: d_1520_v57_})
            (globalState).f0 = (len((d_1517_v54_).set(len(d_1519_v56_), (0) - (p0)))) - (len(d_1521_v58_))
            d_1522_v59_: C0
            nw223_ = C0()
            nw223_.ctor__()
            d_1522_v59_ = nw223_
        d_1523_v60_: _dafny.Seq
        d_1523_v60_ = _dafny.SeqWithoutIsStrInference([p0, -978])
        d_1524_v61_: _dafny.Seq
        d_1524_v61_ = _dafny.SeqWithoutIsStrInference([(d_1523_v60_)[default__.safeIndex(81, len(d_1523_v60_))]])
        d_1525_v62_: _dafny.Map
        d_1525_v62_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0]): d_1457_v6_})
        d_1526_i3_: int
        d_1526_i3_ = 0
        with _dafny.label("8"):
            while (d_1524_v61_) not in (((d_1525_v62_).set(_dafny.SeqWithoutIsStrInference([p0, p0]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1544_i4_ in range(default__.abs(477))]))) | (d_1525_v62_)):
                with _dafny.c_label("8"):
                    if (d_1526_i3_) >= (100):
                        raise _dafny.Break("8")
                    d_1526_i3_ = (d_1526_i3_) + (1)
                    d_1527_v63_: str
                    d_1527_v63_ = _dafny.CodePoint('y')
                    d_1528_v64_: _dafny.Map
                    d_1528_v64_ = _dafny.Map({p0: (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]})
                    d_1529_v65_: _dafny.Map
                    d_1529_v65_ = _dafny.Map({(_dafny.MultiSet([p0])).cardinality: len(d_1528_v64_)})
                    d_1530_v66_: _dafny.Array
                    nw224_ = _dafny.Array(_dafny.CodePoint('D'), 21)
                    d_1530_v66_ = nw224_
                    d_1531_v67_: D12
                    d_1531_v67_ = D12_DC37(p0, d_1530_v66_, _dafny.CodePoint('j'))
                    d_1532_v68_: _dafny.Map
                    d_1532_v68_ = _dafny.Map({((d_1457_v6_) + (d_1456_v5_)).set(default__.safeIndex(p0, len((d_1457_v6_) + (d_1456_v5_))), d_1527_v63_): ((0) - (((d_1529_v65_)[len(_dafny.SeqWithoutIsStrInference([d_1527_v63_ for d_1533_i5_ in range(default__.abs(271))]))] if (len(_dafny.SeqWithoutIsStrInference([d_1527_v63_ for d_1534_i5_ in range(default__.abs(271))]))) in (d_1529_v65_) else (d_1531_v67_).cf61))) * (p0)})
                    d_1535_v70_: _dafny.Seq
                    d_1535_v70_ = _dafny.SeqWithoutIsStrInference([d_1456_v5_])
                    def iife81_():
                        coll31_ = _dafny.Map()
                        compr_31_: _dafny.Seq
                        for compr_31_ in (d_1535_v70_).Elements:
                            d_1536_v69_: _dafny.Seq = compr_31_
                            if (d_1536_v69_) in (d_1535_v70_):
                                coll31_[d_1536_v69_] = (0) - (p0)
                        return _dafny.Map(coll31_)
                    d_1532_v68_ = (iife81_()
                    ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "budcfif")): p0}))
                    (globalState).f1 = default__.safeModuloInt(p0, p0)
                    d_1537_v71_: _dafny.Map
                    d_1537_v71_ = _dafny.Map({default__.fm0(p0, globalState): d_1527_v63_})
                    d_1538_v72_: _dafny.Set
                    d_1538_v72_ = _dafny.Set({d_1452_v1_, d_1452_v1_, (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]})
                    d_1539_v73_: _dafny.Map
                    d_1539_v73_ = _dafny.Map({(d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]: d_1538_v72_})
                    d_1537_v71_ = (d_1537_v71_).set((d_1538_v72_).issubset(((d_1539_v73_)[(d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]] if ((d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]) in (d_1539_v73_) else d_1538_v72_)), d_1527_v63_)
                    d_1540_v74_: D10
                    d_1540_v74_ = D10_DC31((d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))], False, 930)
                    d_1541_v75_: _dafny.Map
                    d_1541_v75_ = _dafny.Map({d_1540_v74_: (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]})
                    d_1542_v76_: _dafny.MultiSet
                    d_1542_v76_ = _dafny.MultiSet([d_1541_v75_, d_1541_v75_, d_1541_v75_, d_1541_v75_])
                    d_1543_v77_: D15
                    d_1543_v77_ = D15_DC44(d_1542_v76_)
                    d_1452_v1_ = (d_1541_v75_) in ((d_1543_v77_).cf79)
                    pass
            pass
        index258_ = default__.safeIndex(395, (d_1451_v0_).length(0))
        (d_1451_v0_)[index258_] = (d_1451_v0_)[default__.safeIndex(395, (d_1451_v0_).length(0))]
        d_1545_v78_: _dafny.Array
        nw225_ = _dafny.Array(int(0), 13)
        d_1545_v78_ = nw225_
        index259_ = default__.safeIndex(684, (d_1545_v78_).length(0))
        (d_1545_v78_)[index259_] = p0
        index260_ = default__.safeIndex(684, (d_1545_v78_).length(0))
        (d_1545_v78_)[index260_] = (p0) * ((len(d_1457_v6_)) + (p0))


class C7(T0):
    def  __init__(self):
        self._f10: bool = False
        self._f19: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f19, f10):
        (self)._f19 = f19
        (self)._f10 = f10

    def fm6(self, p0, p1, p2, p3, globalState):
        return not ((-377) < (-864)) or ((self).f10)

    def fm7(self, p0, globalState):
        return (self).f19

    def fm20(self, p0, p1, p2, globalState):
        return (_dafny.MultiSet([(self).f19])).intersection((D5_DC16(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f19])))).cf27)

    def fm21(self, p0, globalState):
        return (self).f19

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Map = _dafny.Map({})
        d_1546_v0_: int
        d_1546_v0_ = 733
        (globalState).f0 = d_1546_v0_
        d_1547_v1_: C5
        nw226_ = C5()
        nw226_.ctor__()
        d_1547_v1_ = nw226_
        if p1:
            d_1548_v2_: _dafny.Map
            d_1548_v2_ = _dafny.Map({(self).f10: (self).f10})
            d_1549_v3_: _dafny.Map
            d_1549_v3_ = _dafny.Map({len((d_1548_v2_).set((self).f19, p1)): _dafny.SeqWithoutIsStrInference([False, (self).f10])})
            (globalState).f0 = (0) - ((d_1546_v0_ if p1 else (d_1546_v0_) - (len(d_1549_v3_))))
            d_1550_v4_: _dafny.Array
            nw227_ = _dafny.Array(False, 11)
            d_1550_v4_ = nw227_
            d_1551_v5_: _dafny.Seq
            d_1551_v5_ = _dafny.SeqWithoutIsStrInference([d_1550_v4_, d_1550_v4_, d_1550_v4_])
            d_1552_v6_: _dafny.Map
            d_1552_v6_ = _dafny.Map({d_1546_v0_: (self).f19})
            d_1553_v7_: _dafny.Seq
            d_1553_v7_ = _dafny.SeqWithoutIsStrInference([d_1546_v0_, len(d_1552_v6_), d_1546_v0_])
            d_1554_v8_: _dafny.Seq
            d_1554_v8_ = _dafny.SeqWithoutIsStrInference([d_1553_v7_])
            d_1550_v4_ = (d_1551_v5_)[default__.safeIndex(len((d_1554_v8_)[default__.safeIndex((d_1553_v7_)[default__.safeIndex(d_1546_v0_, len(d_1553_v7_))], len(d_1554_v8_))]), len(d_1551_v5_))]
            (globalState).f1 = default__.safeModuloInt(d_1546_v0_, d_1546_v0_)
            d_1555_v10_: _dafny.Set
            def iife82_():
                coll32_ = _dafny.Set()
                compr_32_: int
                for compr_32_ in _dafny.IntegerRange(930, 386):
                    d_1556_v9_: int = compr_32_
                    if ((930) <= (d_1556_v9_)) and ((d_1556_v9_) < (386)):
                        coll32_ = coll32_.union(_dafny.Set([(d_1556_v9_) + (d_1546_v0_)]))
                return _dafny.Set(coll32_)
            d_1555_v10_ = _dafny.Set({len(iife82_()
            )})
            d_1557_v11_: _dafny.Set
            d_1557_v11_ = _dafny.Set({len(d_1555_v10_)})
            rhs250_ = ((d_1557_v11_) - (d_1557_v11_)).issubset((d_1555_v10_) | (d_1557_v11_))
            rhs251_ = d_1546_v0_
            lhs186_ = globalState
            r0 = rhs250_
            lhs186_.f1 = rhs251_
            d_1558_v12_: str
            d_1558_v12_ = _dafny.CodePoint('f')
            d_1559_v13_: _dafny.Seq
            d_1559_v13_ = _dafny.SeqWithoutIsStrInference([(self).f10, not(p1), p1, p1])
            d_1560_v14_: _dafny.Seq
            d_1560_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmux"))
            d_1558_v12_ = default__.fm4((d_1559_v13_)[default__.safeIndex(len(d_1560_v14_), len(d_1559_v13_))], (self).f10, (self).fm7(False, globalState), globalState)
        elif True:
            (globalState).f1 = d_1546_v0_
            d_1561_v15_: _dafny.Seq
            d_1561_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsp"))
            d_1562_v16_: _dafny.Map
            d_1562_v16_ = _dafny.Map({(d_1546_v0_) < (d_1546_v0_): d_1561_v15_})
            d_1563_v17_: str
            d_1563_v17_ = _dafny.CodePoint('n')
            d_1564_v18_: _dafny.Array
            nw228_ = _dafny.Array(None, 2)
            nw228_[int(0)] = (d_1561_v15_).set(default__.safeIndex(d_1546_v0_, len(d_1561_v15_)), d_1563_v17_)
            nw228_[int(1)] = d_1561_v15_
            d_1564_v18_ = nw228_
            index261_ = default__.safeIndex(775, (d_1564_v18_).length(0))
            (d_1564_v18_)[index261_] = _dafny.SeqWithoutIsStrInference([d_1563_v17_ for d_1565_i0_ in range(default__.abs(627))])
            d_1566_v19_: _dafny.Array
            nw229_ = _dafny.Array(int(0), 15)
            d_1566_v19_ = nw229_
            d_1567_v20_: _dafny.Seq
            d_1567_v20_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1568_v21_: _dafny.Set
            d_1568_v21_ = _dafny.Set({d_1567_v20_})
            d_1569_v22_: _dafny.Map
            d_1569_v22_ = _dafny.Map({len(d_1568_v21_): default__.fm51(globalState)})
            d_1570_v23_: _dafny.Map
            d_1570_v23_ = _dafny.Map({len(d_1567_v20_): d_1562_v16_})
            index262_ = default__.safeIndex(775, (d_1564_v18_).length(0))
            rhs252_ = (d_1562_v16_) | (((d_1570_v23_)[-866] if (-866) in (d_1570_v23_) else _dafny.Map({(self).f19: d_1561_v15_})))
            rhs253_ = d_1561_v15_
            rhs254_ = d_1566_v19_
            rhs255_ = d_1569_v22_
            rhs256_ = p1
            lhs187_ = d_1564_v18_
            lhs188_ = default__.safeIndex(775, (d_1564_v18_).length(0))
            d_1562_v16_ = rhs252_
            lhs187_[lhs188_] = rhs253_
            d_1566_v19_ = rhs254_
            d_1569_v22_ = rhs255_
            r0 = rhs256_
            d_1571_v24_: D5
            d_1571_v24_ = D5_DC17(not(False), d_1563_v17_)
            d_1572_v25_: D5
            d_1572_v25_ = D5_DC18(d_1571_v24_)
            d_1573_v26_: D5
            d_1573_v26_ = D5_DC18(d_1572_v25_)
            source31_ = d_1573_v26_
            if source31_.is_DC17:
                d_1574___mcc_h0_ = source31_.cf28
                d_1575___mcc_h1_ = source31_.cf29
                d_1576_cf29_ = d_1575___mcc_h1_
                d_1577_cf28_ = d_1574___mcc_h0_
                d_1578_v27_: _dafny.MultiSet
                d_1578_v27_ = _dafny.MultiSet([True, d_1577_cf28_, d_1577_cf28_])
                d_1579_v28_: _dafny.Map
                d_1579_v28_ = _dafny.Map({d_1546_v0_: ((d_1578_v27_)[(self).f19] if ((self).f19) in (d_1578_v27_) else (0) - (d_1546_v0_))})
                d_1580_v29_: C2
                nw230_ = C2()
                nw230_.ctor__(d_1577_cf28_, len((d_1579_v28_) | (d_1579_v28_)))
                d_1580_v29_ = nw230_
                d_1581_v30_: C3
                nw231_ = C3()
                nw231_.ctor__(d_1577_cf28_, d_1577_cf28_)
                d_1581_v30_ = nw231_
                (globalState).f0 = (d_1580_v29_).f22
                d_1582_v31_: _dafny.Map
                d_1582_v31_ = _dafny.Map({d_1577_cf28_: (self).f19})
                d_1583_v32_: D3
                d_1583_v32_ = D3_DC11(788, d_1582_v31_)
                d_1584_v33_: _dafny.Set
                d_1584_v33_ = _dafny.Set({p1})
                r0 = ((d_1580_v29_).f21 if (d_1582_v31_) != (((d_1583_v32_).cf18).set((self).f10, (self).f10)) else (d_1584_v33_).issubset(d_1584_v33_))
            elif source31_.is_DC16:
                d_1585___mcc_h2_ = source31_.cf27
                d_1586_cf27_ = d_1585___mcc_h2_
                r0 = (p1) == ((self).f19)
                (globalState).f0 = (d_1546_v0_) - ((d_1546_v0_) * (d_1546_v0_))
                d_1587_v34_: D11
                d_1587_v34_ = D11_DC34(d_1546_v0_, (self).fm21(d_1546_v0_, globalState), d_1546_v0_)
                (globalState).f1 = ((d_1587_v34_).cf56) + (d_1546_v0_)
                d_1588_v35_: _dafny.Array
                nw232_ = _dafny.Array(False, 18)
                d_1588_v35_ = nw232_
                d_1589_v36_: _dafny.Set
                d_1589_v36_ = _dafny.Set({(self).f19})
                d_1590_v37_: _dafny.Map
                d_1590_v37_ = _dafny.Map({d_1588_v35_: (d_1589_v36_).ispropersubset(d_1589_v36_)})
                d_1590_v37_ = (d_1590_v37_).set(d_1588_v35_, (self).f10)
            elif True:
                d_1591___mcc_h3_ = source31_.cf30
                d_1592_cf30_ = d_1591___mcc_h3_
                d_1593_v38_: _dafny.Set
                d_1593_v38_ = _dafny.Set({(self).f10})
                d_1594_v39_: _dafny.Set
                d_1594_v39_ = _dafny.Set({d_1593_v38_, d_1593_v38_})
                d_1595_v40_: D3
                d_1595_v40_ = D3_DC9((d_1594_v39_).intersection(d_1594_v39_))
                d_1595_v40_ = (d_1595_v40_ if p1 else d_1595_v40_)
                d_1596_v41_: C0
                nw233_ = C0()
                nw233_.ctor__()
                d_1596_v41_ = nw233_
                d_1597_v42_: C0
                nw234_ = C0()
                nw234_.ctor__()
                d_1597_v42_ = nw234_
                d_1598_v43_: T1
                nw235_ = C3()
                nw235_.ctor__(p1, True)
                d_1598_v43_ = nw235_
                d_1599_v44_: _dafny.Array
                d_1600_v45_: bool
                out49_: _dafny.Array
                out50_: bool
                out49_, out50_ = (d_1547_v1_).m13((d_1598_v43_ if (self).f10 else d_1598_v43_), globalState)
                d_1599_v44_ = out49_
                d_1600_v45_ = out50_
            rhs257_ = not(((0) - (d_1546_v0_)) > (d_1546_v0_))
            rhs258_ = d_1546_v0_
            lhs189_ = globalState
            r0 = rhs257_
            lhs189_.f0 = rhs258_
            rhs259_ = p1
            rhs260_ = d_1546_v0_
            rhs261_ = (self).f10
            rhs262_ = d_1563_v17_
            rhs263_ = d_1564_v18_
            lhs190_ = globalState
            r0 = rhs259_
            lhs190_.f1 = rhs260_
            r0 = rhs261_
            d_1563_v17_ = rhs262_
            d_1564_v18_ = rhs263_
        d_1601_v46_: _dafny.Seq
        d_1601_v46_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({p1: d_1546_v0_}))])
        d_1602_v47_: _dafny.Map
        d_1602_v47_ = _dafny.Map({(self).f10: d_1601_v46_})
        d_1602_v47_ = (d_1602_v47_).set((self).f10, d_1601_v46_)
        d_1603_v48_: D10
        d_1603_v48_ = D10_DC31((self).f19, p1, d_1546_v0_)
        d_1604_v49_: str
        d_1604_v49_ = _dafny.CodePoint('u')
        d_1605_v50_: _dafny.Map
        d_1605_v50_ = _dafny.Map({(self).f10: d_1546_v0_})
        d_1606_v51_: _dafny.Seq
        d_1606_v51_ = _dafny.SeqWithoutIsStrInference([(self).f19, p1])
        d_1607_v52_: _dafny.Map
        d_1607_v52_ = _dafny.Map({d_1604_v49_: not((self).f19)})
        d_1608_v53_: _dafny.Array
        nw236_ = _dafny.Array(None, 29)
        nw236_[int(0)] = (d_1603_v48_).cf50
        nw236_[int(1)] = (d_1604_v49_) not in (_dafny.Map({d_1604_v49_: (self).f10}))
        nw236_[int(2)] = (self).f19
        nw236_[int(3)] = ((self).fm6(d_1546_v0_, d_1546_v0_, d_1546_v0_, (self).f19, globalState)) in (d_1605_v50_)
        nw236_[int(4)] = p1
        nw236_[int(5)] = (self).f10
        nw236_[int(6)] = (self).f10
        nw236_[int(7)] = (d_1546_v0_) > (d_1546_v0_)
        nw236_[int(8)] = (self).f19
        nw236_[int(9)] = (self).f10
        nw236_[int(10)] = (d_1546_v0_) >= (d_1546_v0_)
        nw236_[int(11)] = not(True)
        nw236_[int(12)] = (self).f19
        nw236_[int(13)] = (self).f19
        nw236_[int(14)] = (self).f19
        nw236_[int(15)] = False
        nw236_[int(16)] = (self).f19
        nw236_[int(17)] = p1
        nw236_[int(18)] = (p1) in (d_1605_v50_)
        nw236_[int(19)] = (self).f10
        nw236_[int(20)] = (self).f19
        nw236_[int(21)] = (self).f10
        nw236_[int(22)] = p1
        nw236_[int(23)] = (_dafny.SeqWithoutIsStrInference([True])) != (d_1606_v51_)
        nw236_[int(24)] = ((d_1607_v52_)[d_1604_v49_] if (d_1604_v49_) in (d_1607_v52_) else (self).f19)
        nw236_[int(25)] = (self).f19
        nw236_[int(26)] = not(p1)
        nw236_[int(27)] = (self).f19
        nw236_[int(28)] = (self).f19
        d_1608_v53_ = nw236_
        index263_ = default__.safeIndex(678, (d_1608_v53_).length(0))
        (d_1608_v53_)[index263_] = (self).f10
        index264_ = default__.safeIndex(678, (d_1608_v53_).length(0))
        rhs264_ = (self).f10
        rhs265_ = (d_1546_v0_) + (-186)
        rhs266_ = ((self).f10) == (False)
        rhs267_ = not((self).f19)
        rhs268_ = d_1546_v0_
        lhs191_ = globalState
        lhs192_ = d_1608_v53_
        lhs193_ = default__.safeIndex(678, (d_1608_v53_).length(0))
        lhs194_ = globalState
        r0 = rhs264_
        lhs191_.f0 = rhs265_
        lhs192_[lhs193_] = rhs266_
        r0 = rhs267_
        lhs194_.f1 = rhs268_
        d_1609_v54_: _dafny.MultiSet
        d_1609_v54_ = _dafny.MultiSet([d_1604_v49_, d_1604_v49_, d_1604_v49_, d_1604_v49_, d_1604_v49_])
        d_1610_v55_: _dafny.Seq
        d_1610_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebbclmx"))
        d_1611_v56_: _dafny.Set
        d_1611_v56_ = _dafny.Set({d_1610_v55_})
        d_1612_v57_: _dafny.Seq
        d_1612_v57_ = _dafny.SeqWithoutIsStrInference([d_1606_v51_, _dafny.SeqWithoutIsStrInference([(self).f10, (self).f19])])
        d_1613_v58_: _dafny.MultiSet
        d_1613_v58_ = _dafny.MultiSet([592])
        d_1614_v59_: _dafny.Map
        d_1614_v59_ = _dafny.Map({len((d_1612_v57_)[default__.safeIndex(((d_1613_v58_)[-847] if (-847) in (d_1613_v58_) else d_1546_v0_), len(d_1612_v57_))]): len(d_1605_v50_)})
        d_1615_v60_: _dafny.Map
        d_1615_v60_ = _dafny.Map({len(d_1601_v46_): True})
        d_1616_v61_: _dafny.Array
        nw237_ = _dafny.Array(None, 22)
        nw237_[int(0)] = 190
        nw237_[int(1)] = ((d_1609_v54_).intersection(d_1609_v54_)).cardinality
        nw237_[int(2)] = d_1546_v0_
        nw237_[int(3)] = d_1546_v0_
        nw237_[int(4)] = len(d_1611_v56_)
        nw237_[int(5)] = d_1546_v0_
        nw237_[int(6)] = (0) - (d_1546_v0_)
        nw237_[int(7)] = d_1546_v0_
        nw237_[int(8)] = (((d_1614_v59_)[d_1546_v0_] if (d_1546_v0_) in (d_1614_v59_) else -693)) * (684)
        nw237_[int(9)] = (len(_dafny.Set({(d_1608_v53_)[default__.safeIndex(678, (d_1608_v53_).length(0))]}))) - (len(d_1610_v55_))
        nw237_[int(10)] = len(d_1615_v60_)
        nw237_[int(11)] = d_1546_v0_
        nw237_[int(12)] = len(d_1610_v55_)
        nw237_[int(13)] = d_1546_v0_
        nw237_[int(14)] = default__.fm1(globalState)
        nw237_[int(15)] = (-342) - (len(d_1610_v55_))
        nw237_[int(16)] = d_1546_v0_
        nw237_[int(17)] = d_1546_v0_
        nw237_[int(18)] = d_1546_v0_
        nw237_[int(19)] = d_1546_v0_
        nw237_[int(20)] = d_1546_v0_
        nw237_[int(21)] = d_1546_v0_
        d_1616_v61_ = nw237_
        index265_ = default__.safeIndex(976, (d_1616_v61_).length(0))
        (d_1616_v61_)[index265_] = default__.fm1(globalState)
        index266_ = default__.safeIndex(976, (d_1616_v61_).length(0))
        (d_1616_v61_)[index266_] = default__.fm1(globalState)
        r0 = (self).f10
        r1 = d_1610_v55_
        d_1617_v62_: _dafny.Map
        d_1617_v62_ = _dafny.Map({len(d_1606_v51_): d_1614_v59_})
        r2 = (d_1614_v59_) | (((d_1617_v62_)[852] if (852) in (d_1617_v62_) else d_1614_v59_))
        return r0, r1, r2

    @property
    def f19(self):
        return self._f19

class C8:
    def  __init__(self):
        self.f18: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self, f18):
        (self).f18 = f18

    def fm16(self, p0, p1, p2, p3, globalState):
        return D4_DC13((_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True)]): False}) if False else _dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False})))

    def fm17(self, p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([-426])).intersection(_dafny.MultiSet([-269]))).issubset(_dafny.MultiSet([207, -989, len(_dafny.SeqWithoutIsStrInference([False, False]))]))

    def m10(self, p0, p1, globalState):
        r0: bool = False
        d_1618_v0_: _dafny.MultiSet
        d_1618_v0_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([self.f18]))])
        d_1619_v1_: bool
        d_1619_v1_ = True
        d_1620_v2_: _dafny.Map
        d_1620_v2_ = _dafny.Map({d_1619_v1_: p0})
        r0 = ((d_1618_v0_).intersection(d_1618_v0_)).isdisjoint(default__.fm18((0) - (p0), d_1620_v2_, not(not(d_1619_v1_)), globalState))
        d_1621_v3_: _dafny.Map
        d_1621_v3_ = _dafny.Map({p0: (p0 if d_1619_v1_ else p0)})
        d_1621_v3_ = (d_1621_v3_).set((p0) + (p0), (p0) - ((default__.fm19(d_1619_v1_, d_1619_v1_, d_1619_v1_, p0, globalState)).cardinality))
        d_1622_v4_: _dafny.Array
        nw238_ = _dafny.Array(int(0), 2)
        d_1622_v4_ = nw238_
        index267_ = default__.safeIndex(703, (d_1622_v4_).length(0))
        (d_1622_v4_)[index267_] = (0) - (p0)
        index268_ = default__.safeIndex(703, (d_1622_v4_).length(0))
        (d_1622_v4_)[index268_] = (D1_DC3(default__.fm1(globalState))).cf4
        d_1623_v5_: C6
        nw239_ = C6()
        nw239_.ctor__()
        d_1623_v5_ = nw239_
        (globalState).f7 = default__.safeDivisionInt(((d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))]) * ((d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))]), (d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))])
        d_1624_i0_: int
        d_1624_i0_ = 0
        with _dafny.label("9"):
            while d_1619_v1_:
                with _dafny.c_label("9"):
                    if (d_1624_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1624_i0_ = (d_1624_i0_) + (1)
                    d_1619_v1_ = (_dafny.MultiSet([316, (d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))]])).isdisjoint(d_1618_v0_)
                    d_1625_v6_: D13
                    d_1625_v6_ = D13_DC39(247, d_1619_v1_)
                    d_1619_v1_ = (d_1625_v6_).cf66
                    d_1626_v7_: _dafny.Seq
                    d_1626_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elrwfcc"))
                    d_1627_v9_: _dafny.Seq
                    d_1627_v9_ = _dafny.SeqWithoutIsStrInference([False, d_1619_v1_])
                    d_1628_v10_: _dafny.Seq
                    d_1628_v10_ = _dafny.SeqWithoutIsStrInference([d_1620_v2_])
                    d_1629_v11_: D11
                    d_1629_v11_ = D11_DC34(len(d_1627_v9_), d_1619_v1_, len((d_1628_v10_).set(default__.safeIndex((d_1618_v0_).cardinality, len(d_1628_v10_)), default__.fm51(globalState))))
                    d_1630_v12_: D4
                    d_1630_v12_ = D4_DC13((_dafny.Map({_dafny.SeqWithoutIsStrInference([(d_1629_v11_).cf57, d_1619_v1_]): d_1619_v1_})).set(d_1627_v9_, d_1619_v1_))
                    d_1631_v13_: _dafny.Seq
                    d_1631_v13_ = _dafny.SeqWithoutIsStrInference([d_1630_v12_])
                    d_1632_v14_: D0
                    d_1632_v14_ = D0_DC1(d_1626_v7_)
                    d_1633_v15_: _dafny.Seq
                    d_1633_v15_ = _dafny.SeqWithoutIsStrInference([d_1632_v14_, d_1632_v14_, d_1632_v14_])
                    def iife83_():
                        coll33_ = _dafny.Map()
                        compr_33_: _dafny.Seq
                        for compr_33_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1619_v1_])])).Elements:
                            d_1634_v8_: _dafny.Seq = compr_33_
                            if (d_1634_v8_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1619_v1_])])):
                                coll33_[d_1634_v8_] = d_1619_v1_
                        return _dafny.Map(coll33_)
                    if (self).fm17((d_1626_v7_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "js"))), 448, (_dafny.SeqWithoutIsStrInference([D4_DC13(iife83_()
) for d_1635_i1_ in range(default__.abs(75))])) + (d_1631_v13_), (d_1632_v14_) in (d_1633_v15_), globalState):
                        d_1636_v16_: _dafny.Array
                        def lambda77_(d_1637_v1_):
                            def lambda78_(d_1638_i2_):
                                return (_dafny.Set({d_1637_v1_, d_1637_v1_, d_1637_v1_, d_1637_v1_})) != (_dafny.Set({True}))

                            return lambda78_

                        init41_ = lambda77_(d_1619_v1_)
                        nw240_ = _dafny.Array(None, 15)
                        for i0_41_ in range(nw240_.length(0)):
                            nw240_[i0_41_] = init41_(i0_41_)
                        d_1636_v16_ = nw240_
                        d_1636_v16_ = p1
                        d_1639_v17_: _dafny.Seq
                        d_1639_v17_ = _dafny.SeqWithoutIsStrInference([(d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))]])
                        d_1640_v18_: C4
                        nw241_ = C4()
                        nw241_.ctor__()
                        d_1640_v18_ = nw241_
                        rhs269_ = _dafny.SeqWithoutIsStrInference([(p0) - ((d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))]), (d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))], (len(_dafny.Map({d_1640_v18_: d_1619_v1_}))) + (p0)])
                        rhs270_ = (_dafny.CodePoint('a')) in (d_1626_v7_)
                        d_1639_v17_ = rhs269_
                        r0 = rhs270_
                        d_1641_v19_: _dafny.Map
                        d_1641_v19_ = _dafny.Map({(d_1623_v5_).fm23(False, (d_1618_v0_).cardinality, d_1619_v1_, 406, globalState): d_1619_v1_})
                        d_1642_v20_: _dafny.Map
                        d_1642_v20_ = _dafny.Map({len(d_1641_v19_): d_1619_v1_})
                        index269_ = default__.safeIndex(749, (d_1636_v16_).length(0))
                        (d_1636_v16_)[index269_] = ((d_1642_v20_)[(d_1618_v0_).cardinality] if ((d_1618_v0_).cardinality) in (d_1642_v20_) else default__.fm0((d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))], globalState))
                        index270_ = default__.safeIndex(749, (d_1636_v16_).length(0))
                        (d_1636_v16_)[index270_] = d_1619_v1_
                        (globalState).f1 = 370
                        (globalState).f7 = (((d_1618_v0_)[(d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))]] if ((d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))]) in (d_1618_v0_) else (d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))])) + ((p0 if d_1619_v1_ else (d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))]))
                    elif True:
                        d_1643_v21_: _dafny.Map
                        d_1643_v21_ = _dafny.Map({default__.fm0(p0, globalState): d_1619_v1_})
                        d_1644_v22_: _dafny.Map
                        d_1644_v22_ = _dafny.Map({p1: D3_DC11((d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))], d_1643_v21_)})
                        d_1645_v23_: D3
                        d_1645_v23_ = D3_DC10(d_1626_v7_, True, d_1619_v1_)
                        d_1646_v24_: D3
                        d_1646_v24_ = D3_DC12(((d_1644_v22_)[p1] if (p1) in (d_1644_v22_) else d_1645_v23_))
                        d_1646_v24_ = d_1646_v24_
                        index271_ = default__.safeIndex(703, (d_1622_v4_).length(0))
                        (d_1622_v4_)[index271_] = ((294 if True else len(d_1627_v9_))) + ((d_1622_v4_)[default__.safeIndex(703, (d_1622_v4_).length(0))])
                        d_1647_v26_: _dafny.Array
                        def lambda79_(d_1648_p0_, d_1649_v1_, d_1650_v4_):
                            def lambda80_(d_1651_i3_):
                                def iife84_():
                                    coll34_ = _dafny.Map()
                                    compr_34_: int
                                    for compr_34_ in _dafny.IntegerRange(195, -108):
                                        d_1652_v25_: int = compr_34_
                                        if ((195) <= (d_1652_v25_)) and ((d_1652_v25_) < (-108)):
                                            coll34_[(d_1652_v25_) * (d_1648_p0_)] = D11_DC33(d_1649_v1_, (d_1650_v4_)[default__.safeIndex(703, (d_1650_v4_).length(0))])
                                    return _dafny.Map(coll34_)
                                return iife84_()
                                

                            return lambda80_

                        init42_ = lambda79_(p0, d_1619_v1_, d_1622_v4_)
                        nw242_ = _dafny.Array(None, 8)
                        for i0_42_ in range(nw242_.length(0)):
                            nw242_[i0_42_] = init42_(i0_42_)
                        d_1647_v26_ = nw242_
                        d_1653_v27_: _dafny.Map
                        d_1653_v27_ = _dafny.Map({d_1647_v26_: d_1619_v1_})
                        d_1653_v27_ = (d_1653_v27_).set(d_1647_v26_, True)
                        d_1654_v28_: C4
                        nw243_ = C4()
                        nw243_.ctor__()
                        d_1654_v28_ = nw243_
                        rhs271_ = d_1619_v1_
                        rhs272_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qe"))) + (d_1626_v7_)) + (d_1626_v7_)
                        rhs273_ = p0
                        rhs274_ = d_1619_v1_
                        lhs195_ = globalState
                        d_1619_v1_ = rhs271_
                        d_1626_v7_ = rhs272_
                        lhs195_.f1 = rhs273_
                        d_1619_v1_ = rhs274_
                    (globalState).f0 = (0) - (default__.fm1(globalState))
                    pass
            pass
        r0 = not((not(d_1619_v1_)) or (d_1619_v1_))
        return r0


class C9(T0):
    def  __init__(self):
        self._f10: bool = False
        self._f17: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f17, f10):
        (self)._f17 = f17
        (self)._f10 = f10

    def fm6(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cepp"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onfrv")))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))

    def fm7(self, p0, globalState):
        source32_ = D2_DC6(_dafny.CodePoint('b'))
        if source32_.is_DC7:
            d_1655___mcc_h0_ = source32_.cf9
            d_1656___mcc_h1_ = source32_.cf10
            d_1657___mcc_h2_ = source32_.cf11
            d_1658_cf11_ = d_1657___mcc_h2_
            d_1659_cf10_ = d_1656___mcc_h1_
            d_1660_cf9_ = d_1655___mcc_h0_
            return d_1659_cf10_
        elif source32_.is_DC6:
            d_1661___mcc_h3_ = source32_.cf8
            d_1662_cf8_ = d_1661___mcc_h3_
            return (self).f10
        elif True:
            d_1663___mcc_h4_ = source32_.cf12
            d_1664_cf12_ = d_1663___mcc_h4_
            return False

    def fm14(self, p0, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jn"))])).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbvekb"))]))

    def fm15(self, globalState):
        return len(_dafny.Map({(_dafny.Set({(0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f17, (self).f10, (self).f17]))).cardinality)}) if (self).f10 else _dafny.Set({529})): (0) - ((len(_dafny.Map({(self).f10: 774}))) * (660))}))

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Map = _dafny.Map({})
        d_1665_v0_: _dafny.Array
        nw244_ = _dafny.Array(_dafny.Map({}), 27)
        d_1665_v0_ = nw244_
        d_1666_v1_: str
        d_1666_v1_ = _dafny.CodePoint('g')
        d_1667_v2_: int
        d_1667_v2_ = 260
        index272_ = default__.safeIndex(64, (d_1665_v0_).length(0))
        (d_1665_v0_)[index272_] = _dafny.Map({d_1666_v1_: (0) - (d_1667_v2_)})
        d_1668_v3_: _dafny.Map
        d_1668_v3_ = _dafny.Map({d_1666_v1_: d_1667_v2_})
        index273_ = default__.safeIndex(64, (d_1665_v0_).length(0))
        (d_1665_v0_)[index273_] = (d_1668_v3_) | (d_1668_v3_)
        d_1669_v4_: _dafny.Array
        def lambda81_(d_1670_i0_):
            return (d_1670_i0_) + (-728)

        init43_ = lambda81_
        nw245_ = _dafny.Array(None, 11)
        for i0_43_ in range(nw245_.length(0)):
            nw245_[i0_43_] = init43_(i0_43_)
        d_1669_v4_ = nw245_
        d_1671_v5_: _dafny.Set
        d_1671_v5_ = _dafny.Set({d_1669_v4_, d_1669_v4_, d_1669_v4_})
        if (d_1671_v5_).isdisjoint(d_1671_v5_):
            d_1672_v6_: _dafny.Array
            nw246_ = _dafny.Array(False, 12)
            d_1672_v6_ = nw246_
            d_1672_v6_ = d_1672_v6_
            d_1672_v6_ = (d_1672_v6_ if (self).fm7((self).f10, globalState) else d_1672_v6_)
            d_1673_v7_: _dafny.Seq
            d_1674_v8_: _dafny.Seq
            out51_: _dafny.Seq
            out52_: _dafny.Seq
            out51_, out52_ = default__.m0(len(_dafny.SeqWithoutIsStrInference([d_1666_v1_, d_1666_v1_, d_1666_v1_])), d_1672_v6_, globalState)
            d_1673_v7_ = out51_
            d_1674_v8_ = out52_
            d_1666_v1_ = _dafny.CodePoint('t')
            d_1675_v9_: C1
            nw247_ = C1()
            nw247_.ctor__((self).f17)
            d_1675_v9_ = nw247_
        elif True:
            d_1676_v10_: _dafny.Seq
            d_1676_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbye"))
            d_1677_v11_: _dafny.MultiSet
            d_1677_v11_ = _dafny.MultiSet([(0) - ((d_1667_v2_) - ((self).fm15(globalState))), d_1667_v2_, d_1667_v2_, d_1667_v2_, len(d_1676_v10_)])
            d_1678_v12_: _dafny.Map
            d_1678_v12_ = _dafny.Map({d_1677_v11_: (self).f10})
            (globalState).f1 = ((d_1677_v11_)[d_1667_v2_] if (d_1667_v2_) in (d_1677_v11_) else len(d_1678_v12_))
            d_1679_v13_: _dafny.Array
            nw248_ = _dafny.Array(_dafny.CodePoint('D'), 29)
            d_1679_v13_ = nw248_
            d_1679_v13_ = d_1679_v13_
            d_1680_v14_: _dafny.Map
            d_1680_v14_ = _dafny.Map({p1: d_1667_v2_})
            r0 = (((self).f10 if p1 else p1)) not in (d_1680_v14_)
            d_1681_v15_: _dafny.Array
            nw249_ = _dafny.Array(False, 15)
            d_1681_v15_ = nw249_
            d_1682_v16_: _dafny.Map
            d_1682_v16_ = _dafny.Map({p1: d_1676_v10_})
            d_1683_v17_: _dafny.Map
            d_1683_v17_ = _dafny.Map({len(d_1682_v16_): d_1667_v2_})
            d_1684_v18_: _dafny.Seq
            d_1684_v18_ = _dafny.SeqWithoutIsStrInference([d_1683_v17_, d_1683_v17_])
            d_1685_v19_: _dafny.Map
            d_1685_v19_ = _dafny.Map({(self).f10: d_1683_v17_})
            d_1686_v20_: _dafny.Array
            nw250_ = _dafny.Array(None, 10)
            nw250_[int(0)] = (d_1684_v18_)[default__.safeIndex(d_1667_v2_, len(d_1684_v18_))]
            nw250_[int(1)] = (d_1683_v17_) | (d_1683_v17_)
            nw250_[int(2)] = ((d_1685_v19_)[(self).f17] if ((self).f17) in (d_1685_v19_) else d_1683_v17_)
            nw250_[int(3)] = d_1683_v17_
            nw250_[int(4)] = d_1683_v17_
            nw250_[int(5)] = d_1683_v17_
            nw250_[int(6)] = _dafny.Map({d_1667_v2_: (0) - (d_1667_v2_)})
            nw250_[int(7)] = (d_1683_v17_) | (d_1683_v17_)
            nw250_[int(8)] = _dafny.Map({d_1667_v2_: d_1667_v2_})
            nw250_[int(9)] = _dafny.Map({d_1667_v2_: d_1667_v2_})
            d_1686_v20_ = nw250_
            index274_ = default__.safeIndex(507, (d_1686_v20_).length(0))
            (d_1686_v20_)[index274_] = _dafny.Map({109: default__.fm1(globalState)})
            d_1687_v21_: _dafny.Seq
            d_1687_v21_ = _dafny.SeqWithoutIsStrInference([d_1676_v10_, d_1676_v10_])
            index275_ = default__.safeIndex(507, (d_1686_v20_).length(0))
            rhs275_ = (d_1681_v15_ if ((d_1687_v21_).set(default__.safeIndex(d_1667_v2_, len(d_1687_v21_)), d_1676_v10_)) <= (d_1687_v21_) else d_1681_v15_)
            rhs276_ = d_1683_v17_
            lhs196_ = d_1686_v20_
            lhs197_ = default__.safeIndex(507, (d_1686_v20_).length(0))
            d_1681_v15_ = rhs275_
            lhs196_[lhs197_] = rhs276_
            d_1667_v2_ = d_1667_v2_
        d_1688_v22_: _dafny.Seq
        d_1688_v22_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(d_1667_v2_, len(d_1671_v5_))])
        rhs277_ = 964
        rhs278_ = d_1667_v2_
        rhs279_ = d_1688_v22_
        rhs280_ = (0) - (d_1667_v2_)
        lhs198_ = globalState
        lhs199_ = globalState
        lhs198_.f0 = rhs277_
        d_1667_v2_ = rhs278_
        d_1688_v22_ = rhs279_
        lhs199_.f0 = rhs280_
        d_1689_v23_: _dafny.Seq
        d_1689_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfanlsw"))
        d_1690_v24_: D3
        d_1690_v24_ = D3_DC10(d_1689_v23_, p1, p1)
        d_1691_i1_: int
        d_1691_i1_ = 0
        with _dafny.label("10"):
            while (d_1690_v24_).cf15:
                with _dafny.c_label("10"):
                    if (d_1691_i1_) >= (100):
                        raise _dafny.Break("10")
                    d_1691_i1_ = (d_1691_i1_) + (1)
                    d_1692_v25_: _dafny.Set
                    d_1692_v25_ = _dafny.Set({default__.fm0(477, globalState)})
                    (globalState).f1 = len(d_1692_v25_)
                    d_1693_v26_: D10
                    d_1693_v26_ = D10_DC31(p1, p1, d_1667_v2_)
                    d_1694_v27_: _dafny.Map
                    d_1694_v27_ = _dafny.Map({d_1693_v26_: False})
                    d_1695_v29_: _dafny.MultiSet
                    d_1695_v29_ = _dafny.MultiSet([d_1693_v26_])
                    d_1696_v30_: _dafny.MultiSet
                    def iife85_():
                        coll35_ = _dafny.Map()
                        compr_35_: D10
                        for compr_35_ in (d_1695_v29_).Elements:
                            d_1697_v28_: D10 = compr_35_
                            if (d_1697_v28_) in (d_1695_v29_):
                                coll35_[d_1697_v28_] = p1
                        return _dafny.Map(coll35_)
                    d_1696_v30_ = _dafny.MultiSet([d_1694_v27_, iife85_()
                    ])
                    d_1698_v31_: _dafny.Map
                    d_1698_v31_ = _dafny.Map({p1: D15_DC44(d_1696_v30_)})
                    d_1699_v32_: _dafny.Map
                    d_1699_v32_ = _dafny.Map({d_1688_v22_: d_1698_v31_})
                    d_1700_v33_: _dafny.Map
                    d_1700_v33_ = _dafny.Map({d_1699_v32_: (self).fm6(d_1667_v2_, d_1667_v2_, d_1667_v2_, False, globalState)})
                    if not(((d_1700_v33_)[d_1699_v32_] if (d_1699_v32_) in (d_1700_v33_) else p1)):
                        d_1701_v34_: D1
                        d_1701_v34_ = D1_DC4(default__.safeDivisionInt(d_1667_v2_, d_1667_v2_), d_1669_v4_)
                        rhs281_ = not(not(p1))
                        rhs282_ = d_1666_v1_
                        rhs283_ = default__.safeModuloInt(d_1667_v2_, (d_1667_v2_) * (d_1667_v2_))
                        rhs284_ = d_1701_v34_
                        rhs285_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nn"))) + (d_1689_v23_)).set(default__.safeIndex((0) - (d_1667_v2_), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nn"))) + (d_1689_v23_))), d_1666_v1_)
                        lhs200_ = globalState
                        r0 = rhs281_
                        d_1666_v1_ = rhs282_
                        lhs200_.f1 = rhs283_
                        d_1701_v34_ = rhs284_
                        r1 = rhs285_
                        d_1702_v35_: _dafny.Seq
                        d_1702_v35_ = _dafny.SeqWithoutIsStrInference([d_1692_v25_, d_1692_v25_, d_1692_v25_, _dafny.Set({default__.fm0(len(d_1689_v23_), globalState), p1, (self).f10}), d_1692_v25_])
                        r0 = (d_1692_v25_) in ((d_1702_v35_) + (d_1702_v35_))
                        (globalState).f0 = d_1667_v2_
                        d_1703_v36_: _dafny.Array
                        def lambda82_(d_1704_p1_):
                            def lambda83_(d_1705_i2_):
                                return d_1704_p1_

                            return lambda83_

                        init44_ = lambda82_(p1)
                        nw251_ = _dafny.Array(None, 4)
                        for i0_44_ in range(nw251_.length(0)):
                            nw251_[i0_44_] = init44_(i0_44_)
                        d_1703_v36_ = nw251_
                        index276_ = default__.safeIndex(93, (d_1703_v36_).length(0))
                        (d_1703_v36_)[index276_] = (self).f17
                        index277_ = default__.safeIndex(93, (d_1703_v36_).length(0))
                        (d_1703_v36_)[index277_] = not(not((d_1692_v25_).issubset(d_1692_v25_)))
                        (globalState).f0 = (0) - (d_1667_v2_)
                    elif True:
                        d_1706_v37_: _dafny.Map
                        d_1706_v37_ = _dafny.Map({(self).f10: (self).f17})
                        d_1707_v38_: _dafny.Array
                        nw252_ = _dafny.Array(None, 24)
                        nw252_[int(0)] = d_1669_v4_
                        nw252_[int(1)] = (d_1669_v4_ if ((d_1706_v37_)[(self).f17] if ((self).f17) in (d_1706_v37_) else (self).f17) else d_1669_v4_)
                        nw252_[int(2)] = d_1669_v4_
                        nw252_[int(3)] = d_1669_v4_
                        nw252_[int(4)] = d_1669_v4_
                        nw252_[int(5)] = d_1669_v4_
                        nw252_[int(6)] = d_1669_v4_
                        nw252_[int(7)] = d_1669_v4_
                        nw252_[int(8)] = d_1669_v4_
                        nw252_[int(9)] = d_1669_v4_
                        nw252_[int(10)] = d_1669_v4_
                        nw252_[int(11)] = d_1669_v4_
                        nw252_[int(12)] = d_1669_v4_
                        nw252_[int(13)] = (d_1669_v4_ if (self).f10 else d_1669_v4_)
                        nw252_[int(14)] = d_1669_v4_
                        nw252_[int(15)] = d_1669_v4_
                        nw252_[int(16)] = d_1669_v4_
                        nw252_[int(17)] = d_1669_v4_
                        nw252_[int(18)] = d_1669_v4_
                        nw252_[int(19)] = d_1669_v4_
                        nw252_[int(20)] = d_1669_v4_
                        nw252_[int(21)] = d_1669_v4_
                        nw252_[int(22)] = d_1669_v4_
                        nw252_[int(23)] = d_1669_v4_
                        d_1707_v38_ = nw252_
                        index278_ = default__.safeIndex(587, (d_1707_v38_).length(0))
                        (d_1707_v38_)[index278_] = d_1669_v4_
                        index279_ = default__.safeIndex(587, (d_1707_v38_).length(0))
                        (d_1707_v38_)[index279_] = d_1669_v4_
                        d_1708_v39_: _dafny.MultiSet
                        d_1708_v39_ = _dafny.MultiSet([d_1667_v2_])
                        d_1709_v40_: _dafny.Map
                        d_1709_v40_ = _dafny.Map({d_1693_v26_: d_1667_v2_})
                        rhs286_ = not ((d_1709_v40_) != (d_1709_v40_)) or ((self).f17)
                        rhs287_ = (d_1708_v39_).intersection(d_1708_v39_)
                        rhs288_ = len(d_1689_v23_)
                        lhs201_ = globalState
                        r0 = rhs286_
                        d_1708_v39_ = rhs287_
                        lhs201_.f7 = rhs288_
                        d_1710_v41_: _dafny.Array
                        nw253_ = _dafny.Array(None, 7)
                        nw253_[int(0)] = (self).f10
                        nw253_[int(1)] = (self).f10
                        nw253_[int(2)] = (self).f10
                        nw253_[int(3)] = True
                        nw253_[int(4)] = p1
                        nw253_[int(5)] = False
                        nw253_[int(6)] = (self).f10
                        d_1710_v41_ = nw253_
                        d_1711_v42_: D0
                        d_1711_v42_ = D0_DC2((self).f17, d_1710_v41_)
                        d_1712_v43_: _dafny.Map
                        d_1712_v43_ = _dafny.Map({(0) - (d_1667_v2_): d_1711_v42_})
                        d_1712_v43_ = d_1712_v43_
                        d_1713_v44_: _dafny.Array
                        def lambda84_(d_1714_v2_):
                            def lambda85_(d_1715_i3_):
                                return D2_DC8(D2_DC7(d_1714_v2_, (self).f17, len(_dafny.Set({d_1714_v2_}))))

                            return lambda85_

                        init45_ = lambda84_(d_1667_v2_)
                        nw254_ = _dafny.Array(None, 16)
                        for i0_45_ in range(nw254_.length(0)):
                            nw254_[i0_45_] = init45_(i0_45_)
                        d_1713_v44_ = nw254_
                        d_1716_v45_: D2
                        d_1716_v45_ = D2_DC6(_dafny.CodePoint('m'))
                        d_1717_v46_: D2
                        d_1717_v46_ = D2_DC8(d_1716_v45_)
                        index280_ = default__.safeIndex(755, (d_1713_v44_).length(0))
                        (d_1713_v44_)[index280_] = d_1717_v46_
                        index281_ = default__.safeIndex(755, (d_1713_v44_).length(0))
                        (d_1713_v44_)[index281_] = default__.fm34(d_1667_v2_, globalState)
                        r0 = (d_1667_v2_) <= (default__.safeModuloInt(d_1667_v2_, d_1667_v2_))
                    (globalState).f1 = (0) - ((0) - ((d_1667_v2_) + (d_1667_v2_)))
                    d_1718_v47_: _dafny.Array
                    def lambda86_(d_1719_v2_, d_1720_v23_):
                        def lambda87_(d_1721_i4_):
                            return (D1_DC3(d_1719_v2_)) in (_dafny.Set({D1_DC3(644), D1_DC3(d_1719_v2_), D1_DC3((_dafny.MultiSet([d_1719_v2_])).cardinality), D1_DC3(len(d_1720_v23_)), D1_DC3(d_1719_v2_)}))

                        return lambda87_

                    init46_ = lambda86_(d_1667_v2_, d_1689_v23_)
                    nw255_ = _dafny.Array(None, 27)
                    for i0_46_ in range(nw255_.length(0)):
                        nw255_[i0_46_] = init46_(i0_46_)
                    d_1718_v47_ = nw255_
                    d_1722_v48_: _dafny.Map
                    d_1722_v48_ = _dafny.Map({d_1667_v2_: (self).f17})
                    d_1723_v49_: _dafny.Map
                    d_1723_v49_ = _dafny.Map({((d_1722_v48_)[d_1667_v2_] if (d_1667_v2_) in (d_1722_v48_) else (self).f10): (self).f17})
                    d_1724_v50_: _dafny.Seq
                    d_1724_v50_ = _dafny.SeqWithoutIsStrInference([(self).f17, p1])
                    index282_ = default__.safeIndex(903, (d_1718_v47_).length(0))
                    (d_1718_v47_)[index282_] = ((d_1723_v49_)[not((self).f17)] if (not((self).f17)) in (d_1723_v49_) else (d_1724_v50_)[default__.safeIndex(len(d_1689_v23_), len(d_1724_v50_))])
                    index283_ = default__.safeIndex(903, (d_1718_v47_).length(0))
                    (d_1718_v47_)[index283_] = ((d_1723_v49_)[not(p1)] if (not(p1)) in (d_1723_v49_) else (self).fm6(d_1667_v2_, default__.fm1(globalState), d_1667_v2_, p1, globalState))
                    pass
            pass
        d_1725_v51_: _dafny.Array
        nw256_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
        d_1725_v51_ = nw256_
        index284_ = default__.safeIndex(404, (d_1725_v51_).length(0))
        (d_1725_v51_)[index284_] = d_1689_v23_
        index285_ = default__.safeIndex(404, (d_1725_v51_).length(0))
        (d_1725_v51_)[index285_] = ((d_1689_v23_).set(default__.safeIndex(d_1667_v2_, len(d_1689_v23_)), d_1666_v1_)) + ((d_1689_v23_) + (d_1689_v23_))
        d_1726_v52_: _dafny.Seq
        d_1726_v52_ = _dafny.SeqWithoutIsStrInference([(self).f17, p1])
        r0 = (d_1726_v52_)[default__.safeIndex(d_1667_v2_, len(d_1726_v52_))]
        r0 = (not (p1) or (True)) or ((d_1667_v2_) != ((0) - (d_1667_v2_)))
        r1 = (d_1725_v51_)[default__.safeIndex(404, (d_1725_v51_).length(0))]
        d_1727_v53_: _dafny.Seq
        d_1727_v53_ = _dafny.SeqWithoutIsStrInference([d_1688_v22_, _dafny.SeqWithoutIsStrInference([d_1667_v2_])])
        r2 = default__.fm45(d_1727_v53_, (0) - ((d_1688_v22_)[default__.safeIndex(d_1667_v2_, len(d_1688_v22_))]), globalState)
        return r0, r1, r2

    def m8(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Set = _dafny.Set({})
        d_1728_v0_: int
        d_1728_v0_ = 690
        d_1729_v1_: _dafny.Seq
        d_1729_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hshhrwt"))
        d_1730_i0_: int
        d_1730_i0_ = 0
        with _dafny.label("11"):
            while (default__.fm25(d_1728_v0_, (self).f17, d_1728_v0_, globalState)) == (d_1729_v1_):
                with _dafny.c_label("11"):
                    if (d_1730_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_1730_i0_ = (d_1730_i0_) + (1)
                    d_1731_v2_: _dafny.Array
                    nw257_ = _dafny.Array(_dafny.MultiSet({}), 19)
                    d_1731_v2_ = nw257_
                    d_1731_v2_ = d_1731_v2_
                    d_1732_v3_: _dafny.Array
                    nw258_ = _dafny.Array(False, 28)
                    d_1732_v3_ = nw258_
                    d_1733_v4_: _dafny.Seq
                    d_1734_v5_: _dafny.Seq
                    out53_: _dafny.Seq
                    out54_: _dafny.Seq
                    out53_, out54_ = default__.m0((self).fm15(globalState), d_1732_v3_, globalState)
                    d_1733_v4_ = out53_
                    d_1734_v5_ = out54_
                    d_1735_v6_: bool
                    d_1735_v6_ = False
                    d_1735_v6_ = default__.fm0(d_1728_v0_, globalState)
                    d_1736_v7_: D10
                    d_1736_v7_ = D10_DC30(p3)
                    d_1737_v8_: _dafny.Map
                    d_1737_v8_ = _dafny.Map({(d_1736_v7_ if p1 else d_1736_v7_): d_1732_v3_})
                    d_1737_v8_ = (d_1737_v8_).set(d_1736_v7_, d_1732_v3_)
                    pass
            pass
        d_1738_v9_: D7
        d_1738_v9_ = D7_DC23()
        d_1739_v10_: _dafny.Map
        d_1739_v10_ = _dafny.Map({d_1728_v0_: (self).f10})
        d_1740_v11_: _dafny.Map
        d_1740_v11_ = _dafny.Map({d_1738_v9_: len(d_1739_v10_)})
        d_1741_v12_: _dafny.Seq
        d_1741_v12_ = _dafny.SeqWithoutIsStrInference([d_1728_v0_, ((d_1740_v11_)[d_1738_v9_] if (d_1738_v9_) in (d_1740_v11_) else d_1728_v0_), d_1728_v0_])
        d_1742_v13_: D1
        d_1742_v13_ = D1_DC3(len(d_1741_v12_))
        d_1742_v13_ = (d_1742_v13_ if p1 else D1_DC3(d_1728_v0_))
        d_1743_v14_: _dafny.Map
        d_1743_v14_ = _dafny.Map({len(d_1729_v1_): d_1728_v0_})
        d_1743_v14_ = (d_1743_v14_).set(d_1728_v0_, d_1728_v0_)
        d_1744_v15_: _dafny.Array
        nw259_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
        d_1744_v15_ = nw259_
        d_1745_v16_: _dafny.Seq
        d_1745_v16_ = _dafny.SeqWithoutIsStrInference([d_1744_v15_, d_1744_v15_, d_1744_v15_])
        d_1744_v15_ = (d_1745_v16_)[default__.safeIndex(d_1728_v0_, len(d_1745_v16_))]
        hi7_ = d_1728_v0_
        for d_1746_i1_ in range(d_1728_v0_, hi7_):
            d_1747_v17_: _dafny.Array
            nw260_ = _dafny.Array(False, 7)
            d_1747_v17_ = nw260_
            d_1747_v17_ = d_1747_v17_
            d_1748_v19_: _dafny.Array
            def lambda88_(d_1749_p2_, d_1750_p1_, d_1751_p0_):
                def lambda89_(d_1752_i2_):
                    def iife86_():
                        coll36_ = _dafny.Map()
                        compr_36_: _dafny.Seq
                        for compr_36_ in (_dafny.SeqWithoutIsStrInference([d_1751_p0_])).Elements:
                            d_1753_v18_: _dafny.Seq = compr_36_
                            if (d_1753_v18_) in (_dafny.SeqWithoutIsStrInference([d_1751_p0_])):
                                coll36_[d_1753_v18_] = d_1749_p2_
                        return _dafny.Map(coll36_)
                    return (_dafny.MultiSet([D4_DC13((D4_DC13(_dafny.Map({_dafny.SeqWithoutIsStrInference([d_1749_p2_]): d_1750_p1_}))).cf20), D4_DC13(_dafny.Map({(d_1751_p0_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "neft"))), len(d_1751_p0_)), (self).f10): (D6_DC21(True, (self).f10, (self).f17, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjrt"))))).cf36})), D4_DC13(iife86_()
), D4_DC13(_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10, True]): (self).f17}))])) - (_dafny.MultiSet([D4_DC13(_dafny.Map({d_1751_p0_: (self).f10})), D4_DC13(_dafny.Map({d_1751_p0_: (self).f17}))]))

                return lambda89_

            init47_ = lambda88_(p2, p1, p0)
            nw261_ = _dafny.Array(None, 22)
            for i0_47_ in range(nw261_.length(0)):
                nw261_[i0_47_] = init47_(i0_47_)
            d_1748_v19_ = nw261_
            d_1754_v20_: D4
            d_1754_v20_ = D4_DC13(_dafny.Map({p0: p1}))
            index286_ = default__.safeIndex(659, (d_1748_v19_).length(0))
            (d_1748_v19_)[index286_] = (_dafny.MultiSet([d_1754_v20_])).intersection(_dafny.MultiSet([d_1754_v20_]))
            d_1755_v21_: _dafny.MultiSet
            d_1755_v21_ = _dafny.MultiSet([d_1754_v20_, d_1754_v20_])
            index287_ = default__.safeIndex(659, (d_1748_v19_).length(0))
            (d_1748_v19_)[index287_] = ((d_1755_v21_) | (d_1755_v21_)).set(d_1754_v20_, default__.abs(default__.safeModuloInt((d_1741_v12_)[default__.safeIndex(d_1746_i1_, len(d_1741_v12_))], d_1746_i1_)))
            d_1756_v22_: bool
            d_1756_v22_ = True
            d_1757_v23_: str
            d_1757_v23_ = _dafny.CodePoint('l')
            d_1758_v24_: _dafny.Map
            d_1758_v24_ = _dafny.Map({d_1757_v23_: (self).f17})
            d_1756_v22_ = (not(d_1756_v22_) if ((d_1758_v24_)[d_1757_v23_] if (d_1757_v23_) in (d_1758_v24_) else False) else (self).f17)
            d_1759_v25_: _dafny.Map
            d_1759_v25_ = _dafny.Map({(d_1746_i1_) > (d_1728_v0_): (0) - ((0) - (d_1746_i1_))})
            d_1759_v25_ = (d_1759_v25_).set((self).fm7((self).f10, globalState), d_1746_i1_)
        pat_let_tv37_ = d_1728_v0_
        pat_let_tv38_ = d_1728_v0_
        def lambda90_(source33_):
            if source33_.is_DC23:
                return pat_let_tv37_
            elif source33_.is_DC24:
                d_1760___mcc_h0_ = source33_.cf41
                d_1761_cf41_ = d_1760___mcc_h0_
                return -122
            elif True:
                d_1762___mcc_h1_ = source33_.cf40
                d_1763_cf40_ = d_1762___mcc_h1_
                return (pat_let_tv38_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1764_i3_ in range(default__.abs(519))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1765_i4_ in range(default__.abs(114))])])))

        (globalState).f7 = lambda90_(D7_DC24(d_1728_v0_))
        d_1766_v26_: _dafny.Set
        d_1766_v26_ = _dafny.Set({d_1728_v0_})
        r0 = d_1766_v26_
        return r0

    def m9(self, p0, p1, p2, globalState):
        r0: bool = False
        d_1767_v0_: _dafny.Map
        d_1767_v0_ = _dafny.Map({not(not(not((self).f17))): (self).f17})
        d_1768_v1_: D10
        d_1768_v1_ = D10_DC31(p1, (self).f10, p0)
        pat_let_tv39_ = p1
        pat_let_tv40_ = p0
        pat_let_tv41_ = p0
        pat_let_tv42_ = p0
        pat_let_tv43_ = p0
        pat_let_tv44_ = p1
        def lambda91_(source34_):
            if source34_.is_DC36:
                d_1769___mcc_h0_ = source34_.cf60
                d_1770_cf60_ = d_1769___mcc_h0_
                if pat_let_tv39_:
                    return len(_dafny.SeqWithoutIsStrInference([pat_let_tv40_ for d_1771_i0_ in range(default__.abs(812))]))
                elif True:
                    return pat_let_tv41_
            elif source34_.is_DC37:
                d_1772___mcc_h1_ = source34_.cf61
                d_1773___mcc_h2_ = source34_.cf62
                d_1774___mcc_h3_ = source34_.cf63
                d_1775_cf63_ = d_1774___mcc_h3_
                d_1776_cf62_ = d_1773___mcc_h2_
                d_1777_cf61_ = d_1772___mcc_h1_
                return pat_let_tv42_
            elif True:
                d_1778___mcc_h4_ = source34_.cf59
                d_1779_cf59_ = d_1778___mcc_h4_
                return len((_dafny.Map({(self).f10: pat_let_tv43_})) | (_dafny.Map({pat_let_tv44_: -149})))

        rhs289_ = lambda91_(D12_DC35(_dafny.SeqWithoutIsStrInference([d_1768_v1_, default__.fm46((self).f17, (self).f17, False, globalState), d_1768_v1_, d_1768_v1_, d_1768_v1_])))
        rhs290_ = _dafny.Map({(self).f10: (self).f10})
        lhs202_ = globalState
        lhs202_.f7 = rhs289_
        d_1767_v0_ = rhs290_
        d_1780_v2_: str
        d_1780_v2_ = _dafny.CodePoint('j')
        d_1781_v3_: _dafny.Seq
        d_1781_v3_ = _dafny.SeqWithoutIsStrInference([default__.fm4((self).f10, (self).f17, p1, globalState)])
        rhs291_ = (self).f17
        rhs292_ = d_1780_v2_
        rhs293_ = (d_1781_v3_) < (d_1781_v3_)
        rhs294_ = p0
        lhs203_ = globalState
        r0 = rhs291_
        d_1780_v2_ = rhs292_
        r0 = rhs293_
        lhs203_.f7 = rhs294_
        d_1782_v4_: _dafny.MultiSet
        d_1782_v4_ = _dafny.MultiSet([p0, p0, (p0) + (p0)])
        d_1782_v4_ = d_1782_v4_
        r0 = not((((self).f17) or (not((self).f17)) if (p2)[default__.safeIndex(423, len(p2))] else (self).f10))
        if p1:
            d_1783_v5_: _dafny.Map
            d_1783_v5_ = _dafny.Map({p1: d_1767_v0_})
            d_1767_v0_ = (((d_1783_v5_)[(self).f10] if ((self).f10) in (d_1783_v5_) else d_1767_v0_)).set((self).f17, p1)
            d_1781_v3_ = d_1781_v3_
            d_1784_v6_: _dafny.Array
            def lambda92_(d_1785_p1_, d_1786_p0_):
                def lambda93_(d_1787_i1_):
                    return (_dafny.Map({d_1785_p1_: d_1786_p0_})) == (_dafny.Map({(self).f10: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okyayl")))}))

                return lambda93_

            init48_ = lambda92_(p1, p0)
            nw262_ = _dafny.Array(None, 19)
            for i0_48_ in range(nw262_.length(0)):
                nw262_[i0_48_] = init48_(i0_48_)
            d_1784_v6_ = nw262_
            index288_ = default__.safeIndex(72, (d_1784_v6_).length(0))
            (d_1784_v6_)[index288_] = p1
            d_1788_v8_: _dafny.MultiSet
            d_1788_v8_ = _dafny.MultiSet([p2])
            d_1789_v9_: _dafny.Seq
            def iife87_():
                coll37_ = _dafny.Map()
                compr_37_: _dafny.Seq
                for compr_37_ in (d_1788_v8_).Elements:
                    d_1790_v7_: _dafny.Seq = compr_37_
                    if (d_1790_v7_) in (d_1788_v8_):
                        coll37_[d_1790_v7_] = (self).f10
                return _dafny.Map(coll37_)
            d_1789_v9_ = _dafny.SeqWithoutIsStrInference([p0, len(d_1781_v3_), len(iife87_()
            )])
            index289_ = default__.safeIndex(72, (d_1784_v6_).length(0))
            (d_1784_v6_)[index289_] = (p0) not in (d_1789_v9_)
            d_1791_v10_: _dafny.Seq
            d_1791_v10_ = _dafny.SeqWithoutIsStrInference([default__.fm40((self).f17, p1, globalState)])
            index290_ = default__.safeIndex(72, (d_1784_v6_).length(0))
            rhs295_ = ((default__.fm52(p0, p1, p0, len(d_1789_v9_), globalState)) + (d_1791_v10_)) + ((_dafny.SeqWithoutIsStrInference([p2, p2, p2])) + (d_1791_v10_))
            rhs296_ = (self).f10
            lhs204_ = d_1784_v6_
            lhs205_ = default__.safeIndex(72, (d_1784_v6_).length(0))
            d_1791_v10_ = rhs295_
            lhs204_[lhs205_] = rhs296_
            d_1792_v11_: _dafny.Map
            d_1792_v11_ = _dafny.Map({((d_1784_v6_)[default__.safeIndex(72, (d_1784_v6_).length(0))] if (self).f17 else (d_1784_v6_)[default__.safeIndex(72, (d_1784_v6_).length(0))]): (d_1781_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "meguffaj")))})
            rhs297_ = ((d_1792_v11_)[False] if (False) in (d_1792_v11_) else _dafny.SeqWithoutIsStrInference([d_1780_v2_ for d_1793_i2_ in range(default__.abs(160))]))
            rhs298_ = ((p0) != (p0) if (d_1784_v6_)[default__.safeIndex(72, (d_1784_v6_).length(0))] else ((self).f17 if (d_1784_v6_)[default__.safeIndex(72, (d_1784_v6_).length(0))] else p1))
            rhs299_ = default__.fm1(globalState)
            lhs206_ = globalState
            d_1781_v3_ = rhs297_
            r0 = rhs298_
            lhs206_.f7 = rhs299_
        elif True:
            d_1794_v12_: _dafny.Array
            nw263_ = _dafny.Array(False, 22)
            d_1794_v12_ = nw263_
            d_1795_v13_: D13
            d_1795_v13_ = D13_DC38(_dafny.MultiSet([d_1794_v12_]))
            d_1796_v14_: _dafny.Seq
            d_1796_v14_ = _dafny.SeqWithoutIsStrInference([d_1781_v3_])
            d_1797_v15_: _dafny.Seq
            d_1797_v15_ = _dafny.SeqWithoutIsStrInference([p0, len((d_1796_v14_)[default__.safeIndex(-690, len(d_1796_v14_))]), len(d_1767_v0_), 638, len(d_1781_v3_)])
            pat_let_tv45_ = d_1795_v13_
            pat_let_tv46_ = d_1794_v12_
            pat_let_tv47_ = d_1797_v15_
            pat_let_tv48_ = p0
            pat_let_tv49_ = d_1797_v15_
            def iife88_(_pat_let25_0):
                def iife89_(d_1798_dt__update__tmp_h0_):
                    def iife90_(_pat_let26_0):
                        def iife91_(d_1799_dt__update_hcf64_h0_):
                            return D13_DC38(d_1799_dt__update_hcf64_h0_)
                        return iife91_(_pat_let26_0)
                    return iife90_(((pat_let_tv45_).cf64).set(pat_let_tv46_, default__.abs((pat_let_tv47_)[default__.safeIndex((0) - (pat_let_tv48_), len(pat_let_tv49_))])))
                return iife89_(_pat_let25_0)
            d_1795_v13_ = iife88_(d_1795_v13_)
            d_1800_v16_: _dafny.Array
            nw264_ = _dafny.Array(int(0), 25)
            d_1800_v16_ = nw264_
            d_1800_v16_ = d_1800_v16_
            index291_ = default__.safeIndex(545, (d_1800_v16_).length(0))
            (d_1800_v16_)[index291_] = default__.safeModuloInt(len(d_1781_v3_), p0)
            index292_ = default__.safeIndex(545, (d_1800_v16_).length(0))
            (d_1800_v16_)[index292_] = p0
            (globalState).f7 = default__.safeModuloInt(((d_1782_v4_)[p0] if (p0) in (d_1782_v4_) else -131), p0)
            index293_ = default__.safeIndex(545, (d_1800_v16_).length(0))
            (d_1800_v16_)[index293_] = (0) - ((d_1800_v16_)[default__.safeIndex(545, (d_1800_v16_).length(0))])
        d_1801_v17_: _dafny.Map
        d_1801_v17_ = _dafny.Map({(self).f17: (0) - (p0)})
        d_1802_v18_: _dafny.Set
        d_1802_v18_ = _dafny.Set({(((d_1801_v17_)[True] if (True) in (d_1801_v17_) else len(d_1781_v3_))) * (p0), p0, 565, ((d_1782_v4_).intersection(d_1782_v4_)).cardinality})
        d_1803_v19_: _dafny.Set
        d_1803_v19_ = _dafny.Set({d_1780_v2_, d_1780_v2_})
        d_1804_v21_: _dafny.Map
        d_1804_v21_ = _dafny.Map({p0: -207})
        d_1805_v22_: C8
        nw265_ = C8()
        nw265_.ctor__(d_1780_v2_)
        d_1805_v22_ = nw265_
        d_1806_v23_: _dafny.Seq
        d_1806_v23_ = _dafny.SeqWithoutIsStrInference([default__.fm1(globalState)])
        d_1807_v24_: _dafny.Map
        d_1807_v24_ = _dafny.Map({d_1805_v22_: (d_1806_v23_).set(default__.safeIndex(p0, len(d_1806_v23_)), p0)})
        d_1808_v25_: _dafny.Seq
        d_1808_v25_ = _dafny.SeqWithoutIsStrInference([d_1806_v23_, d_1806_v23_])
        d_1809_v26_: _dafny.MultiSet
        d_1809_v26_ = _dafny.MultiSet([((d_1807_v24_)[d_1805_v22_] if (d_1805_v22_) in (d_1807_v24_) else (d_1808_v25_)[default__.safeIndex(p0, len(d_1808_v25_))]), d_1806_v23_, _dafny.SeqWithoutIsStrInference([587, p0])])
        d_1810_v28_: _dafny.Map
        d_1810_v28_ = _dafny.Map({default__.fm1(globalState): d_1781_v3_})
        d_1811_v29_: _dafny.MultiSet
        d_1811_v29_ = _dafny.MultiSet([p1])
        d_1812_v30_: _dafny.Array
        nw266_ = _dafny.Array(None, 27)
        nw266_[int(0)] = (self).f10
        def iife92_():
            coll38_ = _dafny.Set()
            compr_38_: str
            for compr_38_ in (d_1803_v19_).Elements:
                d_1813_v20_: str = compr_38_
                if (d_1813_v20_) in (d_1803_v19_):
                    coll38_ = coll38_.union(_dafny.Set([d_1813_v20_]))
            return _dafny.Set(coll38_)
        nw266_[int(1)] = (d_1803_v19_).ispropersubset(iife92_()
        )
        nw266_[int(2)] = (len(d_1804_v21_)) < (((d_1809_v26_)[_dafny.SeqWithoutIsStrInference([p0])] if (_dafny.SeqWithoutIsStrInference([p0])) in (d_1809_v26_) else p0))
        nw266_[int(3)] = False
        nw266_[int(4)] = ((self).f17 if (self).f17 else (self).f10)
        nw266_[int(5)] = p1
        def iife93_():
            coll39_ = _dafny.Map()
            compr_39_: int
            for compr_39_ in (d_1810_v28_).keys.Elements:
                d_1814_v27_: int = compr_39_
                if (d_1814_v27_) in (d_1810_v28_):
                    coll39_[(d_1814_v27_) + (p0)] = p0
            return _dafny.Map(coll39_)
        nw266_[int(6)] = (p0) in (iife93_()
        )
        nw266_[int(7)] = (p0) == (p0)
        nw266_[int(8)] = (p0) != (len(_dafny.SeqWithoutIsStrInference([_dafny.Set({p1, p1, (self).f17, not((self).f10)}) for d_1815_i3_ in range(default__.abs(430))])))
        nw266_[int(9)] = False
        nw266_[int(10)] = True
        nw266_[int(11)] = (p0) == (p0)
        nw266_[int(12)] = not(True)
        nw266_[int(13)] = ((self).f17) == ((self).f10)
        nw266_[int(14)] = (d_1811_v29_).isdisjoint(d_1811_v29_)
        nw266_[int(15)] = False
        nw266_[int(16)] = ((d_1767_v0_)[(self).f17] if ((self).f17) in (d_1767_v0_) else (self).f10)
        nw266_[int(17)] = (self).f17
        nw266_[int(18)] = (True) == ((self).f10)
        nw266_[int(19)] = p1
        nw266_[int(20)] = (p0) >= (len(default__.fm38((self).f17, p0, globalState)))
        nw266_[int(21)] = (p0) == (p0)
        nw266_[int(22)] = (self).f17
        nw266_[int(23)] = (len(d_1781_v3_)) != (p0)
        nw266_[int(24)] = p1
        nw266_[int(25)] = False
        nw266_[int(26)] = (self).f17
        d_1812_v30_ = nw266_
        d_1816_v31_: _dafny.Map
        d_1816_v31_ = _dafny.Map({d_1812_v30_: d_1812_v30_})
        rhs300_ = ((d_1802_v18_) | (d_1802_v18_)) | (d_1802_v18_)
        rhs301_ = (self).f10
        rhs302_ = ((d_1816_v31_)[d_1812_v30_] if (d_1812_v30_) in (d_1816_v31_) else d_1812_v30_)
        d_1802_v18_ = rhs300_
        r0 = rhs301_
        d_1812_v30_ = rhs302_
        r0 = not (p1) or ((self).f10)
        return r0

    @property
    def f17(self):
        return self._f17

class C10(T1):
    def  __init__(self):
        self._f15: int = int(0)
        self._f16: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    def ctor__(self, f15, f16):
        (self)._f15 = f15
        (self)._f16 = f16

    def fm8(self, p0, p1, p2, globalState):
        return not((default__.safeDivisionInt((self).f15, len(_dafny.SeqWithoutIsStrInference([(self).f15])))) == (default__.safeDivisionInt((self).f15, (self).f15)))

    def fm9(self, p0, p1, globalState):
        return len((_dafny.Set({516})) - ((_dafny.Set({len(_dafny.Map({(self).f16: (self).f16}))})) - (_dafny.Set({(_dafny.MultiSet([True, False])).cardinality, 352}))))

    def m4(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: int = int(0)
        hi8_ = default__.safeModuloInt(339, (self).f15)
        for d_1817_i0_ in range((self).f15, hi8_):
            if (self).f16:
                d_1818_v0_: _dafny.Seq
                d_1818_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktduqgy"))
                d_1819_v1_: _dafny.Set
                d_1819_v1_ = _dafny.Set({(self).f16})
                d_1820_v2_: _dafny.Set
                d_1820_v2_ = _dafny.Set({d_1819_v1_})
                d_1821_v3_: D3
                d_1821_v3_ = D3_DC9((d_1820_v2_) | (d_1820_v2_))
                d_1822_v4_: D0
                d_1822_v4_ = D0_DC1(default__.fm13((self).f16, not((self).f16), globalState))
                d_1823_v5_: _dafny.Seq
                d_1823_v5_ = _dafny.SeqWithoutIsStrInference([d_1820_v2_])
                pat_let_tv50_ = d_1823_v5_
                pat_let_tv51_ = d_1823_v5_
                pat_let_tv52_ = d_1820_v2_
                def iife94_(_pat_let27_0):
                    def iife95_(d_1824_dt__update__tmp_h0_):
                        def iife96_(_pat_let28_0):
                            def iife97_(d_1825_dt__update_hcf13_h0_):
                                return D3_DC9(d_1825_dt__update_hcf13_h0_)
                            return iife97_(_pat_let28_0)
                        return iife96_(((pat_let_tv50_)[default__.safeIndex((0) - ((self).f15), len(pat_let_tv51_))]) - (pat_let_tv52_))
                    return iife95_(_pat_let27_0)
                rhs303_ = (d_1818_v0_ if (self).f16 else ((d_1822_v4_).cf1) + (d_1818_v0_))
                rhs304_ = iife94_(d_1821_v3_)
                rhs305_ = default__.safeDivisionInt((self).f15, (d_1817_i0_) * (d_1817_i0_))
                lhs207_ = globalState
                d_1818_v0_ = rhs303_
                d_1821_v3_ = rhs304_
                lhs207_.f7 = rhs305_
                d_1826_v6_: _dafny.MultiSet
                d_1826_v6_ = _dafny.MultiSet([(self).f16])
                d_1827_v7_: _dafny.Map
                d_1827_v7_ = _dafny.Map({(self).f16: (d_1826_v6_).cardinality})
                d_1828_v8_: D4
                d_1828_v8_ = D4_DC14((self).f15, (self).f15)
                d_1829_v9_: _dafny.Seq
                d_1829_v9_ = _dafny.SeqWithoutIsStrInference([False])
                d_1830_v10_: _dafny.Array
                nw267_ = _dafny.Array(int(0), 20)
                d_1830_v10_ = nw267_
                d_1831_v11_: _dafny.Map
                d_1831_v11_ = _dafny.Map({d_1830_v10_: (self).f15})
                d_1832_v12_: _dafny.Seq
                d_1832_v12_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                d_1833_v13_: _dafny.Array
                nw268_ = _dafny.Array(None, 26)
                nw268_[int(0)] = (self).f15
                nw268_[int(1)] = (0) - ((self).f15)
                nw268_[int(2)] = d_1817_i0_
                nw268_[int(3)] = len(d_1819_v1_)
                nw268_[int(4)] = (d_1828_v8_).cf22
                nw268_[int(5)] = len(d_1829_v9_)
                nw268_[int(6)] = (self).f15
                nw268_[int(7)] = d_1817_i0_
                nw268_[int(8)] = (0) - (d_1817_i0_)
                nw268_[int(9)] = d_1817_i0_
                nw268_[int(10)] = 244
                nw268_[int(11)] = (self).f15
                nw268_[int(12)] = (self).f15
                nw268_[int(13)] = (self).f15
                nw268_[int(14)] = 435
                nw268_[int(15)] = 526
                nw268_[int(16)] = d_1817_i0_
                nw268_[int(17)] = 484
                nw268_[int(18)] = len(d_1831_v11_)
                nw268_[int(19)] = -955
                nw268_[int(20)] = 496
                nw268_[int(21)] = d_1817_i0_
                nw268_[int(22)] = d_1817_i0_
                nw268_[int(23)] = d_1817_i0_
                nw268_[int(24)] = (d_1832_v12_)[default__.safeIndex(len(d_1818_v0_), len(d_1832_v12_))]
                nw268_[int(25)] = (self).f15
                d_1833_v13_ = nw268_
                d_1834_v14_: D1
                d_1834_v14_ = D1_DC4(((d_1827_v7_)[not((self).fm8(len(d_1818_v0_), (self).f16, (self).f16, globalState))] if (not((self).fm8(len(d_1818_v0_), (self).f16, (self).f16, globalState))) in (d_1827_v7_) else d_1817_i0_), d_1830_v10_)
                d_1835_v15_: _dafny.Map
                d_1835_v15_ = _dafny.Map({(self).f15: (self).f15})
                d_1836_v16_: T0
                nw269_ = C1()
                nw269_.ctor__((self).f16)
                d_1836_v16_ = nw269_
                pat_let_tv53_ = d_1830_v10_
                d_1837_v17_: _dafny.Array
                nw270_ = _dafny.Array(None, 12)
                nw270_[int(0)] = d_1834_v14_
                nw270_[int(1)] = d_1834_v14_
                nw270_[int(2)] = d_1834_v14_
                nw270_[int(3)] = d_1834_v14_
                nw270_[int(4)] = d_1834_v14_
                def iife98_(_pat_let29_0):
                    def iife99_(d_1838_dt__update__tmp_h1_):
                        def iife100_(_pat_let30_0):
                            def iife101_(d_1839_dt__update_hcf6_h0_):
                                return D1_DC4((d_1838_dt__update__tmp_h1_).cf5, d_1839_dt__update_hcf6_h0_)
                            return iife101_(_pat_let30_0)
                        return iife100_(pat_let_tv53_)
                    return iife99_(_pat_let29_0)
                nw270_[int(5)] = iife98_(d_1834_v14_)
                nw270_[int(6)] = D1_DC4((self).f15, d_1830_v10_)
                nw270_[int(7)] = D1_DC4(((d_1835_v15_)[len(_dafny.SeqWithoutIsStrInference([652, (0) - ((self).f15), d_1817_i0_]))] if (len(_dafny.SeqWithoutIsStrInference([652, (0) - ((self).f15), d_1817_i0_]))) in (d_1835_v15_) else len(_dafny.Map({len(d_1818_v0_): d_1836_v16_}))), d_1833_v13_)
                nw270_[int(8)] = D1_DC4(d_1817_i0_, d_1830_v10_)
                nw270_[int(9)] = d_1834_v14_
                nw270_[int(10)] = d_1834_v14_
                nw270_[int(11)] = d_1834_v14_
                d_1837_v17_ = nw270_
                d_1840_v18_: _dafny.Array
                nw271_ = _dafny.Array(None, 2)
                nw271_[int(0)] = d_1837_v17_
                nw271_[int(1)] = d_1837_v17_
                d_1840_v18_ = nw271_
                index294_ = default__.safeIndex(105, (d_1840_v18_).length(0))
                (d_1840_v18_)[index294_] = d_1837_v17_
                index295_ = default__.safeIndex(105, (d_1840_v18_).length(0))
                (d_1840_v18_)[index295_] = d_1837_v17_
                r1 = (d_1836_v16_).f10
                d_1841_v19_: str
                d_1841_v19_ = _dafny.CodePoint('h')
                d_1842_v20_: D14
                d_1842_v20_ = D14_DC43(d_1841_v19_, (0) - (len(d_1818_v0_)), default__.fm1(globalState), (self).f15)
                d_1843_v21_: _dafny.MultiSet
                d_1843_v21_ = _dafny.MultiSet([d_1833_v13_, d_1833_v13_, d_1833_v13_, d_1833_v13_, d_1830_v10_])
                d_1844_v22_: _dafny.Map
                d_1844_v22_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f16])): d_1843_v21_})
                d_1845_v23_: _dafny.Seq
                d_1845_v23_ = _dafny.SeqWithoutIsStrInference([d_1843_v21_])
                d_1846_v24_: D12
                d_1846_v24_ = D12_DC36(d_1843_v21_)
                d_1847_v25_: _dafny.Array
                nw272_ = _dafny.Array(None, 18)
                nw272_[int(0)] = _dafny.MultiSet([d_1830_v10_])
                nw272_[int(1)] = d_1843_v21_
                nw272_[int(2)] = (d_1843_v21_) | (d_1843_v21_)
                nw272_[int(3)] = ((d_1844_v22_)[d_1817_i0_] if (d_1817_i0_) in (d_1844_v22_) else _dafny.MultiSet([d_1830_v10_]))
                nw272_[int(4)] = d_1843_v21_
                nw272_[int(5)] = d_1843_v21_
                nw272_[int(6)] = d_1843_v21_
                nw272_[int(7)] = (d_1843_v21_ if (self).f16 else ((d_1844_v22_)[(self).f15] if ((self).f15) in (d_1844_v22_) else d_1843_v21_))
                nw272_[int(8)] = d_1843_v21_
                nw272_[int(9)] = _dafny.MultiSet([d_1830_v10_, d_1833_v13_, d_1830_v10_])
                nw272_[int(10)] = (d_1845_v23_)[default__.safeIndex(d_1817_i0_, len(d_1845_v23_))]
                nw272_[int(11)] = ((d_1843_v21_).set((D1_DC4((self).f15, d_1830_v10_)).cf6, default__.abs(d_1817_i0_))) - (d_1843_v21_)
                nw272_[int(12)] = (d_1843_v21_).intersection(d_1843_v21_)
                nw272_[int(13)] = d_1843_v21_
                nw272_[int(14)] = (d_1843_v21_).set(d_1833_v13_, default__.abs(d_1817_i0_))
                nw272_[int(15)] = (d_1843_v21_) | (d_1843_v21_)
                nw272_[int(16)] = (d_1846_v24_).cf60
                nw272_[int(17)] = d_1843_v21_
                d_1847_v25_ = nw272_
                index296_ = default__.safeIndex(463, (d_1847_v25_).length(0))
                (d_1847_v25_)[index296_] = (d_1843_v21_) | (d_1843_v21_)
                index297_ = default__.safeIndex(463, (d_1847_v25_).length(0))
                rhs306_ = d_1836_v16_
                rhs307_ = default__.fm53(d_1841_v19_, (self).f15, globalState)
                rhs308_ = (len(default__.fm13((self).f16, (self).f16, globalState)) if (self).f16 else (self).f15)
                rhs309_ = (d_1843_v21_).intersection(_dafny.MultiSet([d_1830_v10_, d_1830_v10_, d_1830_v10_, d_1830_v10_]))
                lhs208_ = globalState
                lhs209_ = d_1847_v25_
                lhs210_ = default__.safeIndex(463, (d_1847_v25_).length(0))
                d_1836_v16_ = rhs306_
                d_1842_v20_ = rhs307_
                lhs208_.f0 = rhs308_
                lhs209_[lhs210_] = rhs309_
                d_1848_v26_: _dafny.Map
                d_1848_v26_ = _dafny.Map({(self).f16: (self).f16})
                d_1849_v27_: _dafny.Map
                d_1849_v27_ = _dafny.Map({(self).f15: d_1832_v12_})
                d_1850_v28_: D13
                d_1850_v28_ = D13_DC39(len(((d_1849_v27_)[(self).f15] if ((self).f15) in (d_1849_v27_) else d_1832_v12_)), (d_1836_v16_).f10)
                nw273_ = _dafny.Array(None, 29)
                nw273_[int(0)] = (self).f15
                nw273_[int(1)] = (self).f15
                nw273_[int(2)] = (self).f15
                nw273_[int(3)] = (self).f15
                nw273_[int(4)] = (0) - (len(d_1819_v1_))
                nw273_[int(5)] = default__.safeModuloInt(d_1817_i0_, len(d_1848_v26_))
                nw273_[int(6)] = (default__.fm54(globalState)).cf23
                nw273_[int(7)] = (len(((d_1818_v0_).set(default__.safeIndex(840, len(d_1818_v0_)), d_1841_v19_)).set(default__.safeIndex((self).f15, len((d_1818_v0_).set(default__.safeIndex(840, len(d_1818_v0_)), d_1841_v19_))), d_1841_v19_))) * (d_1817_i0_)
                nw273_[int(8)] = d_1817_i0_
                nw273_[int(9)] = (self).f15
                nw273_[int(10)] = d_1817_i0_
                nw273_[int(11)] = (self).f15
                nw273_[int(12)] = ((d_1826_v6_)[(d_1836_v16_).f10] if ((d_1836_v16_).f10) in (d_1826_v6_) else len(d_1818_v0_))
                nw273_[int(13)] = ((self).f15) * (len(d_1848_v26_))
                nw273_[int(14)] = d_1817_i0_
                nw273_[int(15)] = d_1817_i0_
                nw273_[int(16)] = (d_1850_v28_).cf65
                nw273_[int(17)] = (d_1817_i0_) - ((self).f15)
                nw273_[int(18)] = default__.fm1(globalState)
                nw273_[int(19)] = len(_dafny.SeqWithoutIsStrInference([(d_1836_v16_).f10]))
                nw273_[int(20)] = default__.safeDivisionInt((0) - ((self).f15), d_1817_i0_)
                nw273_[int(21)] = 64
                nw273_[int(22)] = 579
                nw273_[int(23)] = ((self).f15) * (len(d_1818_v0_))
                nw273_[int(24)] = (self).f15
                nw273_[int(25)] = (self).f15
                nw273_[int(26)] = d_1817_i0_
                nw273_[int(27)] = d_1817_i0_
                nw273_[int(28)] = (self).f15
                d_1833_v13_ = nw273_
            elif True:
                (globalState).f0 = default__.fm1(globalState)
                d_1851_v29_: _dafny.Array
                nw274_ = _dafny.Array(None, 14)
                nw274_[int(0)] = not(True)
                nw274_[int(1)] = (self).f16
                nw274_[int(2)] = (self).f16
                nw274_[int(3)] = (self).f16
                nw274_[int(4)] = (self).f16
                nw274_[int(5)] = (self).f16
                nw274_[int(6)] = (self).f16
                nw274_[int(7)] = (self).f16
                nw274_[int(8)] = (self).f16
                nw274_[int(9)] = (self).f16
                nw274_[int(10)] = (self).f16
                nw274_[int(11)] = False
                nw274_[int(12)] = False
                nw274_[int(13)] = (self).f16
                d_1851_v29_ = nw274_
                d_1852_v30_: _dafny.Set
                d_1852_v30_ = _dafny.Set({d_1851_v29_, d_1851_v29_})
                d_1853_v31_: _dafny.Array
                nw275_ = _dafny.Array(int(0), 6)
                d_1853_v31_ = nw275_
                d_1854_v32_: _dafny.Map
                d_1854_v32_ = _dafny.Map({(self).f15: d_1853_v31_})
                d_1855_v33_: _dafny.Seq
                d_1855_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwfviquy"))
                d_1856_v34_: _dafny.Seq
                d_1856_v34_ = _dafny.SeqWithoutIsStrInference([d_1853_v31_])
                d_1857_v35_: _dafny.Seq
                d_1857_v35_ = _dafny.SeqWithoutIsStrInference([((d_1854_v32_)[len(d_1855_v33_)] if (len(d_1855_v33_)) in (d_1854_v32_) else (d_1856_v34_)[default__.safeIndex((self).f15, len(d_1856_v34_))]), d_1853_v31_])
                d_1858_v36_: _dafny.Set
                d_1858_v36_ = _dafny.Set({d_1817_i0_, len(_dafny.Set({(d_1857_v35_)[default__.safeIndex(d_1817_i0_, len(d_1857_v35_))], d_1853_v31_}))})
                d_1859_v37_: _dafny.Map
                d_1859_v37_ = _dafny.Map({d_1852_v30_: d_1858_v36_})
                d_1859_v37_ = (d_1859_v37_).set(d_1852_v30_, d_1858_v36_)
                d_1860_v38_: C1
                nw276_ = C1()
                nw276_.ctor__((self).f16)
                d_1860_v38_ = nw276_
                r1 = (self).f16
                d_1861_v39_: _dafny.Seq
                d_1861_v39_ = _dafny.SeqWithoutIsStrInference([38])
                d_1862_v40_: _dafny.Seq
                d_1862_v40_ = _dafny.SeqWithoutIsStrInference([(self).f15, (d_1861_v39_)[default__.safeIndex(len(d_1861_v39_), len(d_1861_v39_))]])
                index298_ = default__.safeIndex(795, (d_1853_v31_).length(0))
                (d_1853_v31_)[index298_] = (0) - ((d_1861_v39_)[default__.safeIndex(d_1817_i0_, len(d_1861_v39_))])
                d_1863_v41_: _dafny.Map
                d_1863_v41_ = _dafny.Map({default__.fm4((self).f16, (self).f16, (self).f16, globalState): (self).f16})
                d_1864_v42_: _dafny.Array
                nw277_ = _dafny.Array(_dafny.Map({}), 27)
                d_1864_v42_ = nw277_
                d_1865_v43_: _dafny.Map
                d_1865_v43_ = _dafny.Map({d_1817_i0_: d_1864_v42_})
                d_1866_v44_: _dafny.Array
                nw278_ = _dafny.Array(None, 23)
                nw278_[int(0)] = d_1864_v42_
                nw278_[int(1)] = d_1864_v42_
                nw278_[int(2)] = d_1864_v42_
                nw278_[int(3)] = d_1864_v42_
                nw278_[int(4)] = ((d_1865_v43_)[(self).f15] if ((self).f15) in (d_1865_v43_) else d_1864_v42_)
                nw278_[int(5)] = d_1864_v42_
                nw278_[int(6)] = d_1864_v42_
                nw278_[int(7)] = d_1864_v42_
                nw278_[int(8)] = d_1864_v42_
                nw278_[int(9)] = d_1864_v42_
                nw278_[int(10)] = d_1864_v42_
                nw278_[int(11)] = d_1864_v42_
                nw278_[int(12)] = d_1864_v42_
                nw278_[int(13)] = d_1864_v42_
                nw278_[int(14)] = d_1864_v42_
                nw278_[int(15)] = d_1864_v42_
                nw278_[int(16)] = d_1864_v42_
                nw278_[int(17)] = d_1864_v42_
                nw278_[int(18)] = d_1864_v42_
                nw278_[int(19)] = d_1864_v42_
                nw278_[int(20)] = d_1864_v42_
                nw278_[int(21)] = d_1864_v42_
                nw278_[int(22)] = d_1864_v42_
                d_1866_v44_ = nw278_
                index299_ = default__.safeIndex(983, (d_1866_v44_).length(0))
                (d_1866_v44_)[index299_] = d_1864_v42_
                index300_ = default__.safeIndex(795, (d_1853_v31_).length(0))
                index301_ = default__.safeIndex(983, (d_1866_v44_).length(0))
                rhs310_ = (default__.safeModuloInt(d_1817_i0_, default__.fm1(globalState))) + (d_1817_i0_)
                rhs311_ = (d_1863_v41_ if (self).f16 else d_1863_v41_)
                rhs312_ = d_1864_v42_
                lhs211_ = d_1853_v31_
                lhs212_ = default__.safeIndex(795, (d_1853_v31_).length(0))
                lhs213_ = d_1866_v44_
                lhs214_ = default__.safeIndex(983, (d_1866_v44_).length(0))
                lhs211_[lhs212_] = rhs310_
                d_1863_v41_ = rhs311_
                lhs213_[lhs214_] = rhs312_
            r1 = (self).f16
            (globalState).f0 = d_1817_i0_
            d_1867_v45_: C1
            nw279_ = C1()
            nw279_.ctor__(False)
            d_1867_v45_ = nw279_
        hi9_ = (self).f15
        for d_1868_i1_ in range(9, hi9_):
            d_1869_v46_: _dafny.MultiSet
            d_1869_v46_ = _dafny.MultiSet([(self).f16])
            d_1870_v47_: _dafny.Map
            d_1870_v47_ = _dafny.Map({(d_1869_v46_) - (d_1869_v46_): (default__.fm42(globalState)).cf53})
            d_1871_v48_: _dafny.Seq
            d_1871_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbiish"))
            d_1872_v49_: _dafny.Seq
            d_1872_v49_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_1868_i1_, globalState)])
            d_1873_v50_: _dafny.Map
            d_1873_v50_ = _dafny.Map({len(d_1871_v48_): (d_1872_v49_)[default__.safeIndex((self).f15, len(d_1872_v49_))]})
            d_1870_v47_ = (d_1870_v47_).set(d_1869_v46_, d_1873_v50_)
            d_1874_v51_: C1
            nw280_ = C1()
            nw280_.ctor__((self).f16)
            d_1874_v51_ = nw280_
            r1 = (self).f16
            d_1875_v52_: str
            d_1875_v52_ = _dafny.CodePoint('g')
            d_1876_v53_: _dafny.Map
            d_1876_v53_ = _dafny.Map({(self).f16: d_1875_v52_})
            d_1877_v54_: _dafny.Map
            d_1877_v54_ = _dafny.Map({d_1868_i1_: (d_1876_v53_) | (d_1876_v53_)})
            d_1877_v54_ = (d_1877_v54_).set(d_1868_i1_, d_1876_v53_)
        d_1878_v55_: _dafny.Array
        nw281_ = _dafny.Array(_dafny.Array(None, 0), 3)
        d_1878_v55_ = nw281_
        d_1879_v56_: _dafny.Array
        nw282_ = _dafny.Array(False, 26)
        d_1879_v56_ = nw282_
        d_1880_v57_: _dafny.Array
        nw283_ = _dafny.Array(int(0), 12)
        d_1880_v57_ = nw283_
        d_1881_v58_: D1
        d_1881_v58_ = D1_DC4((self).f15, d_1880_v57_)
        d_1882_v59_: _dafny.Seq
        d_1882_v59_ = _dafny.SeqWithoutIsStrInference([(self).f16])
        d_1883_v60_: D14
        d_1883_v60_ = D14_DC42(d_1881_v58_, d_1882_v59_, (self).f16)
        pat_let_tv54_ = d_1881_v58_
        d_1884_v61_: _dafny.Array
        nw284_ = _dafny.Array(None, 19)
        nw284_[int(0)] = d_1883_v60_
        def iife102_(_pat_let31_0):
            def iife103_(d_1885_dt__update__tmp_h2_):
                def iife104_(_pat_let32_0):
                    def iife105_(d_1886_dt__update_hcf74_h0_):
                        def iife106_(_pat_let33_0):
                            def iife107_(d_1887_dt__update_hcf72_h0_):
                                return D14_DC42(d_1887_dt__update_hcf72_h0_, (d_1885_dt__update__tmp_h2_).cf73, d_1886_dt__update_hcf74_h0_)
                            return iife107_(_pat_let33_0)
                        return iife106_(pat_let_tv54_)
                    return iife105_(_pat_let32_0)
                return iife104_((self).f16)
            return iife103_(_pat_let31_0)
        nw284_[int(1)] = iife102_(d_1883_v60_)
        nw284_[int(2)] = d_1883_v60_
        nw284_[int(3)] = d_1883_v60_
        nw284_[int(4)] = d_1883_v60_
        nw284_[int(5)] = d_1883_v60_
        nw284_[int(6)] = d_1883_v60_
        nw284_[int(7)] = d_1883_v60_
        nw284_[int(8)] = d_1883_v60_
        nw284_[int(9)] = d_1883_v60_
        nw284_[int(10)] = d_1883_v60_
        nw284_[int(11)] = d_1883_v60_
        nw284_[int(12)] = D14_DC42(d_1881_v58_, _dafny.SeqWithoutIsStrInference([(self).f16, (self).f16]), True)
        nw284_[int(13)] = d_1883_v60_
        nw284_[int(14)] = d_1883_v60_
        nw284_[int(15)] = d_1883_v60_
        nw284_[int(16)] = d_1883_v60_
        nw284_[int(17)] = d_1883_v60_
        nw284_[int(18)] = d_1883_v60_
        d_1884_v61_ = nw284_
        d_1888_v62_: _dafny.Map
        d_1888_v62_ = _dafny.Map({d_1879_v56_: (D16_DC48(d_1884_v61_)).cf84})
        d_1889_v63_: _dafny.Map
        d_1889_v63_ = _dafny.Map({(self).f15: (self).f15})
        d_1890_v64_: _dafny.Map
        d_1890_v64_ = _dafny.Map({(self).f16: d_1878_v55_})
        d_1891_v65_: D17
        d_1891_v65_ = D17_DC50(((d_1890_v64_)[(self).f16] if ((self).f16) in (d_1890_v64_) else d_1878_v55_))
        d_1892_v66_: _dafny.Seq
        d_1892_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdcyl"))
        rhs313_ = (d_1891_v65_).cf86
        rhs314_ = d_1888_v62_
        rhs315_ = (self).f15
        rhs316_ = d_1889_v63_
        rhs317_ = (d_1892_v66_) != (d_1892_v66_)
        lhs215_ = globalState
        d_1878_v55_ = rhs313_
        d_1888_v62_ = rhs314_
        lhs215_.f0 = rhs315_
        d_1889_v63_ = rhs316_
        r1 = rhs317_
        d_1893_v67_: str
        d_1893_v67_ = _dafny.CodePoint('r')
        d_1893_v67_ = d_1893_v67_
        index302_ = default__.safeIndex(165, (d_1878_v55_).length(0))
        (d_1878_v55_)[index302_] = d_1879_v56_
        d_1894_v68_: _dafny.Map
        d_1894_v68_ = _dafny.Map({default__.fm0((self).f15, globalState): d_1879_v56_})
        index303_ = default__.safeIndex(165, (d_1878_v55_).length(0))
        (d_1878_v55_)[index303_] = (((d_1894_v68_)[(self).f16] if ((self).f16) in (d_1894_v68_) else d_1879_v56_) if (self).f16 else d_1879_v56_)
        arr4_ = (d_1878_v55_)[default__.safeIndex(165, (d_1878_v55_).length(0))]
        index304_ = default__.safeIndex(495, ((d_1878_v55_)[default__.safeIndex(165, (d_1878_v55_).length(0))]).length(0))
        arr4_[index304_] = (self).f16
        d_1895_v69_: _dafny.Map
        d_1895_v69_ = _dafny.Map({(self).f16: d_1892_v66_})
        arr5_ = (d_1878_v55_)[default__.safeIndex(165, (d_1878_v55_).length(0))]
        index305_ = default__.safeIndex(495, ((d_1878_v55_)[default__.safeIndex(165, (d_1878_v55_).length(0))]).length(0))
        arr5_[index305_] = not((((d_1895_v69_)[(self).f16] if ((self).f16) in (d_1895_v69_) else d_1892_v66_)) != ((d_1892_v66_) + (d_1892_v66_)))
        d_1896_v70_: _dafny.Set
        d_1896_v70_ = _dafny.Set({d_1892_v66_})
        r0 = (d_1896_v70_).intersection(d_1896_v70_)
        r1 = (self).f16
        r2 = 299
        return r0, r1, r2

    def m5(self, p0, globalState):
        d_1897_v0_: str
        d_1897_v0_ = _dafny.CodePoint('r')
        d_1898_v1_: _dafny.Seq
        d_1898_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbhekjmp"))
        d_1897_v0_ = (d_1898_v1_)[default__.safeIndex(p0, len(d_1898_v1_))]
        d_1899_v2_: _dafny.Array
        nw285_ = _dafny.Array(_dafny.Map({}), 27)
        d_1899_v2_ = nw285_
        d_1899_v2_ = d_1899_v2_
        d_1900_v3_: _dafny.Array
        nw286_ = _dafny.Array(_dafny.MultiSet({}), 3)
        d_1900_v3_ = nw286_
        d_1901_v4_: _dafny.Seq
        d_1901_v4_ = _dafny.SeqWithoutIsStrInference([(self).f16, (self).f16, (self).f16])
        index306_ = default__.safeIndex(281, (d_1900_v3_).length(0))
        (d_1900_v3_)[index306_] = _dafny.MultiSet((d_1901_v4_) + (d_1901_v4_))
        d_1902_v5_: C9
        nw287_ = C9()
        nw287_.ctor__((self).f16, (self).f16)
        d_1902_v5_ = nw287_
        d_1903_v6_: _dafny.Set
        d_1903_v6_ = _dafny.Set({d_1902_v5_})
        d_1904_v7_: _dafny.MultiSet
        d_1904_v7_ = _dafny.MultiSet([(self).f16, (d_1903_v6_).ispropersubset(d_1903_v6_)])
        index307_ = default__.safeIndex(281, (d_1900_v3_).length(0))
        (d_1900_v3_)[index307_] = d_1904_v7_
        d_1905_v8_: _dafny.Array
        nw288_ = _dafny.Array(False, 1)
        d_1905_v8_ = nw288_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1905_v8_).length(0)):
            d_1906_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_1906_i0_)) and ((d_1906_i0_) < ((d_1905_v8_).length(0)))):
                (d_1905_v8_)[(d_1906_i0_)] = not((_dafny.Map({_dafny.Map({(d_1902_v5_).f17: (self).f16}): p0})) != (_dafny.Map({_dafny.Map({True: (d_1902_v5_).f17}): p0})))
        d_1907_v9_: _dafny.Set
        d_1907_v9_ = _dafny.Set({not(False), (self).f16})
        d_1908_v10_: _dafny.Map
        d_1908_v10_ = _dafny.Map({len(d_1907_v9_): (d_1902_v5_).f17})
        d_1909_v12_: _dafny.Map
        d_1909_v12_ = _dafny.Map({True: (d_1902_v5_).f17})
        d_1910_i1_: int
        d_1910_i1_ = 0
        with _dafny.label("12"):
            def iife133_():
                coll49_ = _dafny.Set()
                compr_49_: int
                for compr_49_ in (d_1908_v10_).keys.Elements:
                    d_1981_v11_: int = compr_49_
                    if (d_1981_v11_) in (d_1908_v10_):
                        coll49_ = coll49_.union(_dafny.Set([default__.safeDivisionInt(d_1981_v11_, 340)]))
                return _dafny.Set(coll49_)
            while (len(iife133_()
            )) > ((len(d_1909_v12_)) - (p0)):
                with _dafny.c_label("12"):
                    if (d_1910_i1_) >= (100):
                        raise _dafny.Break("12")
                    d_1910_i1_ = (d_1910_i1_) + (1)
                    if (self).f16:
                        d_1911_v13_: C7
                        nw289_ = C7()
                        nw289_.ctor__(False, not(not((d_1902_v5_).f17)))
                        d_1911_v13_ = nw289_
                        d_1912_v14_: bool
                        d_1912_v14_ = False
                        d_1912_v14_ = (d_1911_v13_).fm7(((self).f15) < ((self).f15), globalState)
                        d_1913_v15_: _dafny.Set
                        d_1913_v15_ = _dafny.Set({(self).f15, (self).f15})
                        d_1914_v16_: _dafny.Seq
                        d_1914_v16_ = _dafny.SeqWithoutIsStrInference([d_1913_v15_, d_1913_v15_])
                        d_1915_v19_: _dafny.Seq
                        d_1915_v19_ = _dafny.SeqWithoutIsStrInference([737])
                        d_1916_v20_: _dafny.Seq
                        d_1916_v20_ = _dafny.SeqWithoutIsStrInference([(self).f15, (d_1915_v19_)[default__.safeIndex(p0, len(d_1915_v19_))]])
                        d_1917_v23_: _dafny.MultiSet
                        d_1917_v23_ = _dafny.MultiSet([D15_DC46()])
                        d_1918_v26_: _dafny.Array
                        nw290_ = _dafny.Array(None, 19)
                        nw290_[int(0)] = d_1913_v15_
                        nw290_[int(1)] = _dafny.Set({(self).f15})
                        nw290_[int(2)] = d_1913_v15_
                        nw290_[int(3)] = d_1913_v15_
                        nw290_[int(4)] = d_1913_v15_
                        nw290_[int(5)] = (d_1914_v16_)[default__.safeIndex(-357, len(d_1914_v16_))]
                        def iife108_():
                            coll40_ = _dafny.Set()
                            compr_40_: int
                            for compr_40_ in _dafny.IntegerRange(993, 733):
                                d_1919_v17_: int = compr_40_
                                if ((993) <= (d_1919_v17_)) and ((d_1919_v17_) < (733)):
                                    coll40_ = coll40_.union(_dafny.Set([(d_1919_v17_) * ((self).f15)]))
                            return _dafny.Set(coll40_)
                        nw290_[int(6)] = iife108_()
                        
                        def iife109_():
                            coll41_ = _dafny.Set()
                            compr_41_: int
                            for compr_41_ in _dafny.IntegerRange(113, 960):
                                d_1920_v18_: int = compr_41_
                                if ((113) <= (d_1920_v18_)) and ((d_1920_v18_) < (960)):
                                    coll41_ = coll41_.union(_dafny.Set([(d_1920_v18_) * (len(d_1901_v4_))]))
                            return _dafny.Set(coll41_)
                        nw290_[int(7)] = iife109_()
                        
                        def iife110_():
                            coll42_ = _dafny.Set()
                            compr_42_: int
                            for compr_42_ in (d_1916_v20_).Elements:
                                d_1921_v21_: int = compr_42_
                                if (d_1921_v21_) in (d_1916_v20_):
                                    coll42_ = coll42_.union(_dafny.Set([(d_1921_v21_) * ((0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))).cardinality))]))
                            return _dafny.Set(coll42_)
                        nw290_[int(8)] = iife110_()
                        
                        nw290_[int(9)] = d_1913_v15_
                        nw290_[int(10)] = d_1913_v15_
                        nw290_[int(11)] = d_1913_v15_
                        def iife111_():
                            coll43_ = _dafny.Set()
                            compr_43_: int
                            for compr_43_ in _dafny.IntegerRange(155, 775):
                                d_1922_v22_: int = compr_43_
                                if ((155) <= (d_1922_v22_)) and ((d_1922_v22_) < (775)):
                                    coll43_ = coll43_.union(_dafny.Set([(d_1922_v22_) * ((self).f15)]))
                            return _dafny.Set(coll43_)
                        nw290_[int(12)] = iife111_()
                        
                        nw290_[int(13)] = _dafny.Set({(self).f15, p0, (d_1917_v23_).cardinality})
                        nw290_[int(14)] = d_1913_v15_
                        nw290_[int(15)] = d_1913_v15_
                        def iife112_():
                            coll44_ = _dafny.Set()
                            compr_44_: int
                            for compr_44_ in (d_1908_v10_).keys.Elements:
                                d_1923_v24_: int = compr_44_
                                if (d_1923_v24_) in (d_1908_v10_):
                                    coll44_ = coll44_.union(_dafny.Set([default__.safeDivisionInt(d_1923_v24_, -723)]))
                            return _dafny.Set(coll44_)
                        nw290_[int(16)] = iife112_()
                        
                        def iife113_():
                            coll45_ = _dafny.Set()
                            compr_45_: int
                            for compr_45_ in _dafny.IntegerRange(405, 863):
                                d_1924_v25_: int = compr_45_
                                if ((405) <= (d_1924_v25_)) and ((d_1924_v25_) < (863)):
                                    coll45_ = coll45_.union(_dafny.Set([default__.safeModuloInt(d_1924_v25_, len(_dafny.Set({D16_DC49((0) - (p0)), D16_DC49((self).f15)})))]))
                            return _dafny.Set(coll45_)
                        nw290_[int(17)] = iife113_()
                        
                        nw290_[int(18)] = d_1913_v15_
                        d_1918_v26_ = nw290_
                        d_1925_v27_: _dafny.Map
                        d_1925_v27_ = _dafny.Map({d_1908_v10_: d_1918_v26_})
                        d_1925_v27_ = (d_1925_v27_).set(d_1908_v10_, d_1918_v26_)
                        d_1901_v4_ = (d_1901_v4_) + (d_1901_v4_)
                        d_1912_v14_ = True
                    elif True:
                        d_1926_v28_: bool
                        d_1926_v28_ = False
                        d_1927_v29_: _dafny.MultiSet
                        d_1927_v29_ = _dafny.MultiSet([d_1907_v9_])
                        d_1926_v28_ = (_dafny.MultiSet([d_1907_v9_])).isdisjoint((d_1927_v29_).intersection(d_1927_v29_))
                        d_1928_v30_: _dafny.Seq
                        d_1928_v30_ = _dafny.SeqWithoutIsStrInference([d_1905_v8_, d_1905_v8_, d_1905_v8_, d_1905_v8_, d_1905_v8_])
                        d_1929_v31_: _dafny.Seq
                        d_1929_v31_ = _dafny.SeqWithoutIsStrInference([d_1905_v8_, d_1905_v8_, (d_1928_v30_)[default__.safeIndex(p0, len(d_1928_v30_))], d_1905_v8_, d_1905_v8_])
                        d_1930_v32_: _dafny.Map
                        d_1930_v32_ = _dafny.Map({(d_1929_v31_).set(default__.safeIndex((self).f15, len(d_1929_v31_)), d_1905_v8_): (self).f16})
                        d_1930_v32_ = (d_1930_v32_).set(d_1928_v30_, ((self).f15) >= ((self).f15))
                        d_1898_v1_ = (d_1898_v1_) + (d_1898_v1_)
                        d_1931_v33_: _dafny.Map
                        d_1931_v33_ = _dafny.Map({(d_1902_v5_).f17: p0})
                        d_1932_v34_: _dafny.Array
                        nw291_ = _dafny.Array(None, 24)
                        nw291_[int(0)] = d_1931_v33_
                        nw291_[int(1)] = d_1931_v33_
                        nw291_[int(2)] = d_1931_v33_
                        nw291_[int(3)] = (d_1931_v33_) | (d_1931_v33_)
                        nw291_[int(4)] = (d_1931_v33_) | (_dafny.Map({d_1926_v28_: (self).f15}))
                        nw291_[int(5)] = ((d_1931_v33_).set(True, 57)) | (d_1931_v33_)
                        nw291_[int(6)] = d_1931_v33_
                        nw291_[int(7)] = d_1931_v33_
                        nw291_[int(8)] = _dafny.Map({(d_1902_v5_).f17: (self).f15})
                        nw291_[int(9)] = (d_1931_v33_).set(d_1926_v28_, p0)
                        nw291_[int(10)] = d_1931_v33_
                        nw291_[int(11)] = (d_1931_v33_) | (d_1931_v33_)
                        nw291_[int(12)] = (d_1931_v33_) | (_dafny.Map({(d_1902_v5_).f17: p0}))
                        nw291_[int(13)] = (d_1931_v33_) | ((d_1931_v33_).set((d_1902_v5_).f17, (self).f15))
                        nw291_[int(14)] = (d_1931_v33_) | (d_1931_v33_)
                        nw291_[int(15)] = (d_1931_v33_).set((self).f16, (self).f15)
                        nw291_[int(16)] = (d_1931_v33_ if (d_1902_v5_).f17 else d_1931_v33_)
                        nw291_[int(17)] = d_1931_v33_
                        nw291_[int(18)] = _dafny.Map({(d_1902_v5_).f17: default__.fm1(globalState)})
                        nw291_[int(19)] = d_1931_v33_
                        nw291_[int(20)] = d_1931_v33_
                        nw291_[int(21)] = (d_1931_v33_) | (d_1931_v33_)
                        nw291_[int(22)] = d_1931_v33_
                        nw291_[int(23)] = d_1931_v33_
                        d_1932_v34_ = nw291_
                        d_1932_v34_ = d_1932_v34_
                        d_1897_v0_ = d_1897_v0_
                    d_1933_v35_: D17
                    d_1933_v35_ = D17_DC51(d_1897_v0_, (self).f16)
                    source35_ = d_1933_v35_
                    if source35_.is_DC51:
                        d_1934___mcc_h0_ = source35_.cf87
                        d_1935___mcc_h1_ = source35_.cf88
                        d_1936_cf88_ = d_1935___mcc_h1_
                        d_1937_cf87_ = d_1934___mcc_h0_
                        (globalState).f0 = (self).f15
                        d_1936_cf88_ = (((self).f15) + (p0)) != (p0)
                        d_1938_v36_: D6
                        d_1938_v36_ = D6_DC21((d_1902_v5_).f17, d_1936_cf88_, (self).f16, (self).f15)
                        d_1939_v37_: _dafny.Set
                        out55_: _dafny.Set
                        out55_ = (d_1902_v5_).m8(d_1901_v4_, d_1936_cf88_, (False if (d_1938_v36_).cf36 else (d_1902_v5_).fm6(p0, (self).f15, (self).f15, (d_1902_v5_).f17, globalState)), d_1907_v9_, globalState)
                        d_1939_v37_ = out55_
                        (globalState).f0 = default__.safeDivisionInt(p0, (self).f15)
                    elif source35_.is_DC50:
                        d_1940___mcc_h2_ = source35_.cf86
                        d_1941_cf86_ = d_1940___mcc_h2_
                        index308_ = default__.safeIndex(195, (d_1905_v8_).length(0))
                        (d_1905_v8_)[index308_] = (d_1902_v5_).f17
                        index309_ = default__.safeIndex(195, (d_1905_v8_).length(0))
                        (d_1905_v8_)[index309_] = not(not((self).f16))
                        d_1942_v38_: _dafny.Map
                        d_1942_v38_ = _dafny.Map({(d_1902_v5_).f17: d_1897_v0_})
                        d_1942_v38_ = (d_1942_v38_).set(not ((self).f16) or (not((d_1902_v5_).f17)), (d_1898_v1_)[default__.safeIndex(p0, len(d_1898_v1_))])
                        d_1943_v39_: C8
                        nw292_ = C8()
                        nw292_.ctor__(d_1897_v0_)
                        d_1943_v39_ = nw292_
                        d_1944_v40_: D14
                        d_1944_v40_ = D14_DC43(d_1897_v0_, len(_dafny.Set({(self).f15, (self).f15})), (self).f15, 994)
                        pat_let_tv55_ = p0
                        pat_let_tv56_ = p0
                        pat_let_tv57_ = d_1897_v0_
                        pat_let_tv58_ = p0
                        d_1945_v41_: _dafny.Array
                        nw293_ = _dafny.Array(None, 29)
                        nw293_[int(0)] = d_1944_v40_
                        nw293_[int(1)] = (d_1944_v40_ if (d_1902_v5_).f17 else d_1944_v40_)
                        nw293_[int(2)] = d_1944_v40_
                        def iife115_(_pat_let35_0):
                            def iife116_(d_1946_dt__update__tmp_h0_):
                                def iife117_(_pat_let36_0):
                                    def iife118_(d_1947_dt__update_hcf76_h0_):
                                        def iife119_(_pat_let37_0):
                                            def iife120_(d_1948_dt__update_hcf78_h0_):
                                                return D14_DC43((d_1946_dt__update__tmp_h0_).cf75, d_1947_dt__update_hcf76_h0_, (d_1946_dt__update__tmp_h0_).cf77, d_1948_dt__update_hcf78_h0_)
                                            return iife120_(_pat_let37_0)
                                        return iife119_(pat_let_tv56_)
                                    return iife118_(_pat_let36_0)
                                return iife117_(pat_let_tv55_)
                            return iife116_(_pat_let35_0)
                        def iife114_(_pat_let34_0):
                            def iife121_(d_1949_dt__update__tmp_h1_):
                                def iife122_(_pat_let38_0):
                                    def iife123_(d_1950_dt__update_hcf75_h0_):
                                        def iife124_(_pat_let39_0):
                                            def iife125_(d_1951_dt__update_hcf77_h0_):
                                                return D14_DC43(d_1950_dt__update_hcf75_h0_, (d_1949_dt__update__tmp_h1_).cf76, d_1951_dt__update_hcf77_h0_, (d_1949_dt__update__tmp_h1_).cf78)
                                            return iife125_(_pat_let39_0)
                                        return iife124_((self).f15)
                                    return iife123_(_pat_let38_0)
                                return iife122_(pat_let_tv57_)
                            return iife121_(_pat_let34_0)
                        nw293_[int(3)] = iife114_(iife115_(d_1944_v40_))
                        nw293_[int(4)] = d_1944_v40_
                        nw293_[int(5)] = d_1944_v40_
                        nw293_[int(6)] = d_1944_v40_
                        nw293_[int(7)] = d_1944_v40_
                        nw293_[int(8)] = d_1944_v40_
                        nw293_[int(9)] = d_1944_v40_
                        nw293_[int(10)] = default__.fm53(d_1943_v39_.f18, p0, globalState)
                        nw293_[int(11)] = d_1944_v40_
                        nw293_[int(12)] = d_1944_v40_
                        nw293_[int(13)] = d_1944_v40_
                        nw293_[int(14)] = d_1944_v40_
                        nw293_[int(15)] = D14_DC43(d_1897_v0_, 14, p0, (0) - ((self).f15))
                        nw293_[int(16)] = d_1944_v40_
                        nw293_[int(17)] = D14_DC43(d_1897_v0_, p0, p0, p0)
                        nw293_[int(18)] = default__.fm53(d_1943_v39_.f18, (self).f15, globalState)
                        nw293_[int(19)] = d_1944_v40_
                        nw293_[int(20)] = D14_DC43(d_1943_v39_.f18, (self).f15, len(d_1898_v1_), (self).f15)
                        nw293_[int(21)] = D14_DC43(_dafny.CodePoint('d'), (self).f15, p0, p0)
                        nw293_[int(22)] = D14_DC43(d_1943_v39_.f18, 380, (self).f15, len(d_1898_v1_))
                        nw293_[int(23)] = d_1944_v40_
                        nw293_[int(24)] = d_1944_v40_
                        def iife126_(_pat_let40_0):
                            def iife127_(d_1952_dt__update__tmp_h2_):
                                def iife128_(_pat_let41_0):
                                    def iife129_(d_1953_dt__update_hcf76_h1_):
                                        return D14_DC43((d_1952_dt__update__tmp_h2_).cf75, d_1953_dt__update_hcf76_h1_, (d_1952_dt__update__tmp_h2_).cf77, (d_1952_dt__update__tmp_h2_).cf78)
                                    return iife129_(_pat_let41_0)
                                return iife128_(pat_let_tv58_)
                            return iife127_(_pat_let40_0)
                        nw293_[int(25)] = iife126_(d_1944_v40_)
                        nw293_[int(26)] = d_1944_v40_
                        nw293_[int(27)] = d_1944_v40_
                        nw293_[int(28)] = d_1944_v40_
                        d_1945_v41_ = nw293_
                        index310_ = default__.safeIndex(355, (d_1945_v41_).length(0))
                        (d_1945_v41_)[index310_] = d_1944_v40_
                        index311_ = default__.safeIndex(355, (d_1945_v41_).length(0))
                        rhs318_ = default__.safeModuloInt((self).f15, (self).f15)
                        rhs319_ = d_1944_v40_
                        lhs216_ = globalState
                        lhs217_ = d_1945_v41_
                        lhs218_ = default__.safeIndex(355, (d_1945_v41_).length(0))
                        lhs216_.f1 = rhs318_
                        lhs217_[lhs218_] = rhs319_
                    elif True:
                        d_1954___mcc_h3_ = source35_.cf89
                        d_1955_cf89_ = d_1954___mcc_h3_
                        d_1956_v42_: _dafny.Array
                        nw294_ = _dafny.Array(int(0), 19)
                        d_1956_v42_ = nw294_
                        d_1957_v43_: _dafny.Array
                        def lambda94_(d_1958_v4_, d_1959_v5_):
                            def lambda95_(d_1960_i2_):
                                return D4_DC13(_dafny.Map({d_1958_v4_: (d_1959_v5_).f17}))

                            return lambda95_

                        init49_ = lambda94_(d_1901_v4_, d_1902_v5_)
                        nw295_ = _dafny.Array(None, 13)
                        for i0_49_ in range(nw295_.length(0)):
                            nw295_[i0_49_] = init49_(i0_49_)
                        d_1957_v43_ = nw295_
                        d_1961_v44_: D4
                        d_1961_v44_ = D4_DC13(_dafny.Map({(d_1901_v4_).set(default__.safeIndex((self).f15, len(d_1901_v4_)), (d_1902_v5_).f17): (d_1902_v5_).f17}))
                        index312_ = default__.safeIndex(289, (d_1957_v43_).length(0))
                        (d_1957_v43_)[index312_] = d_1961_v44_
                        index313_ = default__.safeIndex(289, (d_1957_v43_).length(0))
                        rhs320_ = d_1956_v42_
                        rhs321_ = d_1961_v44_
                        lhs219_ = d_1957_v43_
                        lhs220_ = default__.safeIndex(289, (d_1957_v43_).length(0))
                        d_1956_v42_ = rhs320_
                        lhs219_[lhs220_] = rhs321_
                        d_1962_v45_: _dafny.Map
                        d_1962_v45_ = _dafny.Map({(d_1909_v12_).set((d_1902_v5_).f17, (d_1902_v5_).f17): d_1897_v0_})
                        def iife130_():
                            coll46_ = _dafny.Set()
                            compr_46_: _dafny.Map
                            for compr_46_ in (d_1962_v45_).keys.Elements:
                                d_1963_v46_: _dafny.Map = compr_46_
                                if (d_1963_v46_) in (d_1962_v45_):
                                    coll46_ = coll46_.union(_dafny.Set([d_1963_v46_]))
                            return _dafny.Set(coll46_)
                        (globalState).f7 = (0) - ((0) - ((self).fm9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxqtm")), iife130_()
                        , globalState)))
                        d_1964_v47_: _dafny.Map
                        d_1964_v47_ = _dafny.Map({d_1898_v1_: p0})
                        d_1965_v48_: _dafny.Map
                        d_1965_v48_ = _dafny.Map({(self).f16: len(d_1964_v47_)})
                        d_1966_v49_: _dafny.MultiSet
                        d_1966_v49_ = _dafny.MultiSet([len(d_1965_v48_), (self).f15])
                        d_1967_v50_: C5
                        nw296_ = C5()
                        nw296_.ctor__()
                        d_1967_v50_ = nw296_
                        d_1968_v51_: _dafny.Map
                        d_1968_v51_ = _dafny.Map({(self).f15: (self).f15})
                        d_1969_v52_: _dafny.Seq
                        d_1969_v52_ = _dafny.SeqWithoutIsStrInference([((d_1968_v51_)[len(d_1968_v51_)] if (len(d_1968_v51_)) in (d_1968_v51_) else p0)])
                        d_1970_v53_: _dafny.Map
                        d_1970_v53_ = _dafny.Map({(d_1966_v49_) - (_dafny.MultiSet([p0, len(_dafny.Set({d_1967_v50_})), len(d_1898_v1_)])): (d_1969_v52_) + (d_1969_v52_)})
                        d_1970_v53_ = (d_1970_v53_) | (d_1970_v53_)
                        d_1971_v54_: _dafny.Array
                        nw297_ = _dafny.Array(_dafny.Map({}), 4)
                        d_1971_v54_ = nw297_
                        index314_ = default__.safeIndex(263, (d_1971_v54_).length(0))
                        (d_1971_v54_)[index314_] = _dafny.Map({d_1909_v12_: (d_1902_v5_).f17})
                        d_1972_v55_: _dafny.Map
                        d_1972_v55_ = _dafny.Map({(d_1909_v12_) | (d_1909_v12_): True})
                        d_1973_v56_: _dafny.Seq
                        d_1973_v56_ = _dafny.SeqWithoutIsStrInference([d_1908_v10_, _dafny.Map({(self).fm9(d_1898_v1_, _dafny.Set({d_1909_v12_}), globalState): (d_1902_v5_).f17}), d_1908_v10_, d_1908_v10_])
                        index315_ = default__.safeIndex(263, (d_1971_v54_).length(0))
                        def iife131_():
                            coll47_ = _dafny.Map()
                            compr_47_: _dafny.Map
                            for compr_47_ in (d_1972_v55_).keys.Elements:
                                d_1974_v57_: _dafny.Map = compr_47_
                                if (d_1974_v57_) in (d_1972_v55_):
                                    coll47_[d_1974_v57_] = (d_1902_v5_).f17
                            return _dafny.Map(coll47_)
                        rhs322_ = d_1897_v0_
                        rhs323_ = len(((d_1969_v52_) + (d_1969_v52_)).set(default__.safeIndex(default__.safeDivisionInt(len(d_1907_v9_), (self).f15), len((d_1969_v52_) + (d_1969_v52_))), len(d_1973_v56_)))
                        rhs324_ = d_1972_v55_
                        rhs325_ = (d_1972_v55_ if (d_1902_v5_).f17 else (iife131_()
                        ).set(d_1909_v12_, (d_1902_v5_).f17))
                        lhs221_ = globalState
                        lhs222_ = d_1971_v54_
                        lhs223_ = default__.safeIndex(263, (d_1971_v54_).length(0))
                        d_1897_v0_ = rhs322_
                        lhs221_.f1 = rhs323_
                        lhs222_[lhs223_] = rhs324_
                        d_1972_v55_ = rhs325_
                    d_1975_v58_: C6
                    nw298_ = C6()
                    nw298_.ctor__()
                    d_1975_v58_ = nw298_
                    d_1976_v59_: _dafny.Map
                    d_1976_v59_ = _dafny.Map({(self).f16: _dafny.Set({d_1897_v0_})})
                    d_1977_v61_: _dafny.Map
                    def iife132_():
                        coll48_ = _dafny.Set()
                        compr_48_: str
                        for compr_48_ in (_dafny.Map({d_1897_v0_: (self).f15})).keys.Elements:
                            d_1978_v60_: str = compr_48_
                            if (d_1978_v60_) in (_dafny.Map({d_1897_v0_: (self).f15})):
                                coll48_ = coll48_.union(_dafny.Set([d_1978_v60_]))
                        return _dafny.Set(coll48_)
                    d_1977_v61_ = _dafny.Map({(self).f16: len(((d_1976_v59_)[True] if (True) in (d_1976_v59_) else iife132_()
                    ))})
                    d_1979_v62_: _dafny.Map
                    d_1979_v62_ = _dafny.Map({(self).f16: d_1977_v61_})
                    d_1980_v63_: _dafny.Map
                    d_1980_v63_ = _dafny.Map({((d_1979_v62_)[True] if (True) in (d_1979_v62_) else d_1977_v61_): (_dafny.Set({True, (d_1902_v5_).f17, (self).f16})) | (d_1907_v9_)})
                    d_1980_v63_ = (d_1980_v63_).set(d_1977_v61_, _dafny.Set({(d_1902_v5_).f17, (self).f16, (self).f16, (self).f16}))
                    pass
            pass
        if (self).f16:
            (globalState).f7 = 449
            d_1982_v64_: _dafny.Array
            nw299_ = _dafny.Array(None, 2)
            nw299_[int(0)] = d_1897_v0_
            nw299_[int(1)] = d_1897_v0_
            d_1982_v64_ = nw299_
            d_1983_v65_: D17
            d_1983_v65_ = D17_DC51(d_1897_v0_, (d_1902_v5_).f17)
            index316_ = default__.safeIndex(851, (d_1982_v64_).length(0))
            (d_1982_v64_)[index316_] = (_dafny.CodePoint('i') if (d_1902_v5_).f17 else (d_1983_v65_).cf87)
            d_1984_v66_: bool
            d_1984_v66_ = False
            d_1985_v67_: D5
            d_1985_v67_ = D5_DC17((d_1902_v5_).f17, _dafny.CodePoint('o'))
            d_1986_v68_: _dafny.Map
            d_1986_v68_ = _dafny.Map({(d_1902_v5_).f17: (d_1985_v67_).cf29})
            d_1987_v69_: _dafny.Map
            d_1987_v69_ = _dafny.Map({((d_1986_v68_)[(d_1902_v5_).f17] if ((d_1902_v5_).f17) in (d_1986_v68_) else d_1897_v0_): d_1984_v66_})
            d_1988_v70_: _dafny.Seq
            d_1988_v70_ = _dafny.SeqWithoutIsStrInference([d_1987_v69_, d_1987_v69_, d_1987_v69_])
            d_1989_v72_: D18
            def iife134_():
                coll50_ = _dafny.Map()
                compr_50_: str
                for compr_50_ in (d_1898_v1_).Elements:
                    d_1990_v71_: str = compr_50_
                    if (d_1990_v71_) in (d_1898_v1_):
                        coll50_[d_1990_v71_] = True
                return _dafny.Map(coll50_)
            d_1989_v72_ = D18_DC53(_dafny.SeqWithoutIsStrInference([iife134_()
, d_1987_v69_]))
            index317_ = default__.safeIndex(851, (d_1982_v64_).length(0))
            rhs326_ = d_1897_v0_
            rhs327_ = (self).f16
            rhs328_ = ((d_1989_v72_).cf90).set(default__.safeIndex(p0, len((d_1989_v72_).cf90)), d_1987_v69_)
            lhs224_ = d_1982_v64_
            lhs225_ = default__.safeIndex(851, (d_1982_v64_).length(0))
            lhs224_[lhs225_] = rhs326_
            d_1984_v66_ = rhs327_
            d_1988_v70_ = rhs328_
            if (d_1902_v5_).f17:
                (globalState).f1 = p0
                d_1991_v73_: _dafny.Array
                nw300_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_1991_v73_ = nw300_
                d_1991_v73_ = d_1991_v73_
                d_1992_v74_: _dafny.Map
                d_1992_v74_ = _dafny.Map({852: d_1905_v8_})
                d_1992_v74_ = (d_1992_v74_).set((self).f15, d_1905_v8_)
                (globalState).f1 = 239
                d_1898_v1_ = (d_1898_v1_) + ((d_1898_v1_) + (d_1898_v1_))
            elif True:
                d_1993_v75_: _dafny.MultiSet
                d_1993_v75_ = _dafny.MultiSet([(self).f15])
                d_1994_v76_: _dafny.Set
                d_1994_v76_ = _dafny.Set({(((d_1993_v75_)[419] if (419) in (d_1993_v75_) else (self).f15)) - ((self).f15)})
                def iife135_():
                    coll51_ = _dafny.Set()
                    compr_51_: int
                    for compr_51_ in _dafny.IntegerRange(614, 831):
                        d_1995_v77_: int = compr_51_
                        if ((614) <= (d_1995_v77_)) and ((d_1995_v77_) < (831)):
                            coll51_ = coll51_.union(_dafny.Set([default__.safeDivisionInt(d_1995_v77_, (self).f15)]))
                    return _dafny.Set(coll51_)
                def iife136_():
                    coll52_ = _dafny.Set()
                    compr_52_: int
                    for compr_52_ in (d_1908_v10_).keys.Elements:
                        d_1996_v78_: int = compr_52_
                        if (d_1996_v78_) in (d_1908_v10_):
                            coll52_ = coll52_.union(_dafny.Set([default__.safeModuloInt(d_1996_v78_, 186)]))
                    return _dafny.Set(coll52_)
                d_1994_v76_ = ((d_1994_v76_).intersection(iife135_()
                )).intersection((d_1994_v76_) | (iife136_()
                ))
                index318_ = default__.safeIndex(871, (d_1905_v8_).length(0))
                (d_1905_v8_)[index318_] = (d_1902_v5_).f17
                index319_ = default__.safeIndex(871, (d_1905_v8_).length(0))
                rhs329_ = d_1984_v66_
                rhs330_ = (d_1904_v7_).issubset(_dafny.MultiSet([(d_1902_v5_).f17, (self).f16]))
                rhs331_ = p0
                rhs332_ = d_1898_v1_
                rhs333_ = (d_1902_v5_).f17
                lhs226_ = globalState
                lhs227_ = d_1905_v8_
                lhs228_ = default__.safeIndex(871, (d_1905_v8_).length(0))
                d_1984_v66_ = rhs329_
                d_1984_v66_ = rhs330_
                lhs226_.f0 = rhs331_
                d_1898_v1_ = rhs332_
                lhs227_[lhs228_] = rhs333_
                d_1997_v79_: _dafny.Map
                d_1997_v79_ = _dafny.Map({d_1898_v1_: (p0) - (p0)})
                d_1997_v79_ = d_1997_v79_
                d_1998_v80_: D3
                d_1998_v80_ = D3_DC11((self).f15, d_1909_v12_)
                d_1909_v12_ = (d_1998_v80_).cf18
                d_1999_v81_: bool
                d_2000_v82_: bool
                d_2001_v83_: D0
                d_2002_v84_: int
                out56_: bool
                out57_: bool
                out58_: D0
                out59_: int
                out56_, out57_, out58_, out59_ = (self).m7((d_1902_v5_).f17, True, globalState)
                d_1999_v81_ = out56_
                d_2000_v82_ = out57_
                d_2001_v83_ = out58_
                d_2002_v84_ = out59_
            d_1984_v66_ = not ((d_1902_v5_).f17) or ((d_1898_v1_) < (d_1898_v1_))
            d_2003_v85_: _dafny.Map
            d_2003_v85_ = _dafny.Map({(d_1902_v5_).f17: d_1898_v1_})
            (globalState).f7 = default__.safeDivisionInt(len(d_2003_v85_), p0)
        elif True:
            d_2004_v86_: C8
            nw301_ = C8()
            nw301_.ctor__(d_1897_v0_)
            d_2004_v86_ = nw301_
            d_2005_v87_: C0
            nw302_ = C0()
            nw302_.ctor__()
            d_2005_v87_ = nw302_
            if (default__.safeDivisionInt((_dafny.MultiSet([(d_1902_v5_).f17, (d_1902_v5_).f17])).cardinality, p0)) <= ((0) - (p0)):
                d_2006_v88_: bool
                d_2006_v88_ = True
                d_2007_v89_: _dafny.Seq
                d_2007_v89_ = _dafny.SeqWithoutIsStrInference([164, (0) - (len(d_1898_v1_))])
                d_2008_v90_: _dafny.Seq
                d_2008_v90_ = _dafny.SeqWithoutIsStrInference([len(d_1898_v1_), len(d_2007_v89_)])
                d_2009_v91_: _dafny.Seq
                d_2009_v91_ = _dafny.SeqWithoutIsStrInference([d_2007_v89_])
                d_2006_v88_ = (d_2009_v91_) < (d_2009_v91_)
                d_2010_v92_: _dafny.Array
                nw303_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_2010_v92_ = nw303_
                d_2011_v93_: _dafny.Map
                d_2011_v93_ = _dafny.Map({(d_1902_v5_).f17: (d_2010_v92_ if (d_1902_v5_).f17 else d_2010_v92_)})
                d_2012_v94_: T0
                nw304_ = C7()
                nw304_.ctor__((d_1902_v5_).f17, not((self).f16))
                d_2012_v94_ = nw304_
                d_2013_v95_: _dafny.Map
                d_2013_v95_ = _dafny.Map({d_2012_v94_: (0) - ((self).f15)})
                d_2011_v93_ = (d_2011_v93_).set((p0) == (len(d_2013_v95_)), d_2010_v92_)
                d_2014_v96_: C4
                nw305_ = C4()
                nw305_.ctor__()
                d_2014_v96_ = nw305_
                d_2015_v98_: D19
                def iife137_():
                    coll53_ = _dafny.Set()
                    compr_53_: int
                    for compr_53_ in _dafny.IntegerRange(74, 436):
                        d_2016_v97_: int = compr_53_
                        if ((74) <= (d_2016_v97_)) and ((d_2016_v97_) < (436)):
                            coll53_ = coll53_.union(_dafny.Set([(d_2016_v97_) * (p0)]))
                    return _dafny.Set(coll53_)
                d_2015_v98_ = D19_DC57(iife137_()
)
                d_2017_v99_: _dafny.Set
                d_2017_v99_ = _dafny.Set({(d_2015_v98_).cf98})
                d_2018_v100_: _dafny.MultiSet
                d_2018_v100_ = _dafny.MultiSet([(self).f15, p0, (len(d_2017_v99_)) * (p0)])
                d_2018_v100_ = ((d_2018_v100_ if (self).f16 else d_2018_v100_)) | (d_2018_v100_)
                (globalState).f7 = default__.safeModuloInt((d_2007_v89_)[default__.safeIndex(len(_dafny.Set({(d_2012_v94_).f10})), len(d_2007_v89_))], default__.safeDivisionInt((self).f15, len(d_1907_v9_)))
            elif True:
                d_2019_v101_: _dafny.Map
                d_2019_v101_ = _dafny.Map({d_2004_v86_.f18: (self).f16})
                d_2020_v102_: T1
                nw306_ = C2()
                nw306_.ctor__((d_1902_v5_).f17, (self).f15)
                d_2020_v102_ = nw306_
                d_2021_v103_: _dafny.Map
                d_2021_v103_ = _dafny.Map({((d_2019_v101_)[d_2004_v86_.f18] if (d_2004_v86_.f18) in (d_2019_v101_) else (d_1902_v5_).f17): d_2020_v102_})
                d_2021_v103_ = (d_2021_v103_).set((d_1902_v5_).f17, d_2020_v102_)
                rhs334_ = ((self).f15) - ((self).f15)
                rhs335_ = d_2004_v86_.f18
                rhs336_ = (0) - ((p0) - ((p0 if (d_1902_v5_).f17 else len(_dafny.SeqWithoutIsStrInference([d_2004_v86_.f18 for d_2022_i3_ in range(default__.abs(334))])))))
                lhs229_ = globalState
                lhs230_ = d_2004_v86_
                lhs231_ = globalState
                lhs229_.f1 = rhs334_
                lhs230_.f18 = rhs335_
                lhs231_.f0 = rhs336_
                d_2023_v104_: _dafny.Array
                def lambda96_(d_2024_v1_):
                    def lambda97_(d_2025_i4_):
                        return (_dafny.Map({d_2024_v1_: d_2024_v1_})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pabfkukml")): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2026_i5_ in range(default__.abs(102))])}))

                    return lambda97_

                init50_ = lambda96_(d_1898_v1_)
                nw307_ = _dafny.Array(None, 20)
                for i0_50_ in range(nw307_.length(0)):
                    nw307_[i0_50_] = init50_(i0_50_)
                d_2023_v104_ = nw307_
                d_2027_v105_: _dafny.Map
                d_2027_v105_ = _dafny.Map({d_1898_v1_: d_1898_v1_})
                index320_ = default__.safeIndex(125, (d_2023_v104_).length(0))
                (d_2023_v104_)[index320_] = (d_2027_v105_) | (d_2027_v105_)
                d_2028_v106_: _dafny.Seq
                d_2028_v106_ = _dafny.SeqWithoutIsStrInference([(d_2027_v105_) | (_dafny.Map({d_1898_v1_: d_1898_v1_}))])
                d_2029_v107_: _dafny.Set
                d_2029_v107_ = _dafny.Set({d_1909_v12_, d_1909_v12_})
                d_2030_v108_: _dafny.Map
                d_2030_v108_ = _dafny.Map({default__.fm55((self).f15, (self).f15, (d_1902_v5_).f17, globalState): (self).fm9(d_1898_v1_, d_2029_v107_, globalState)})
                index321_ = default__.safeIndex(125, (d_2023_v104_).length(0))
                (d_2023_v104_)[index321_] = (d_2028_v106_)[default__.safeIndex(((self).f15 if (d_1902_v5_).f17 else len(d_2030_v108_)), len(d_2028_v106_))]
                d_2031_v109_: bool
                d_2032_v110_: bool
                d_2033_v111_: D0
                d_2034_v112_: int
                out60_: bool
                out61_: bool
                out62_: D0
                out63_: int
                out60_, out61_, out62_, out63_ = (self).m7((d_1902_v5_).f17, (d_1902_v5_).f17, globalState)
                d_2031_v109_ = out60_
                d_2032_v110_ = out61_
                d_2033_v111_ = out62_
                d_2034_v112_ = out63_
                d_2035_v113_: _dafny.Array
                nw308_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_2035_v113_ = nw308_
                d_2036_v114_: D17
                d_2036_v114_ = D17_DC50(d_2035_v113_)
                pat_let_tv59_ = d_2035_v113_
                d_2037_v115_: _dafny.Seq
                def iife138_(_pat_let42_0):
                    def iife139_(d_2038_dt__update__tmp_h3_):
                        def iife140_(_pat_let43_0):
                            def iife141_(d_2039_dt__update_hcf86_h0_):
                                return D17_DC50(d_2039_dt__update_hcf86_h0_)
                            return iife141_(_pat_let43_0)
                        return iife140_(pat_let_tv59_)
                    return iife139_(_pat_let42_0)
                d_2037_v115_ = _dafny.SeqWithoutIsStrInference([d_2036_v114_, d_2036_v114_, iife138_(d_2036_v114_)])
                d_2040_v116_: _dafny.Map
                d_2040_v116_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnd"))): (_dafny.SeqWithoutIsStrInference([d_2036_v114_, d_2036_v114_, d_2036_v114_, D17_DC50(d_2035_v113_)]) if not((d_1902_v5_).f17) else d_2037_v115_)})
                index322_ = default__.safeIndex(781, (d_1905_v8_).length(0))
                (d_1905_v8_)[index322_] = (d_1902_v5_).f17
                index323_ = default__.safeIndex(813, (d_1905_v8_).length(0))
                (d_1905_v8_)[index323_] = (d_1902_v5_).f17
                d_2041_v117_: D3
                d_2041_v117_ = D3_DC9(_dafny.Set({_dafny.Set({(d_1902_v5_).f17}), d_1907_v9_, d_1907_v9_, d_1907_v9_}))
                d_2042_v118_: D20
                d_2042_v118_ = D20_DC59(_dafny.Map({388: d_2037_v115_}))
                d_2043_v119_: _dafny.Array
                nw309_ = _dafny.Array(D13.default()(), 6)
                d_2043_v119_ = nw309_
                d_2044_v120_: _dafny.Map
                d_2044_v120_ = _dafny.Map({d_2031_v109_: d_2043_v119_})
                index324_ = default__.safeIndex(781, (d_1905_v8_).length(0))
                index325_ = default__.safeIndex(813, (d_1905_v8_).length(0))
                rhs337_ = (((d_2042_v118_).cf101) | (d_2040_v116_)) | (d_2040_v116_)
                rhs338_ = len((_dafny.Map({(d_1902_v5_).f17: d_2043_v119_})) | (d_2044_v120_))
                rhs339_ = d_2032_v110_
                rhs340_ = False
                rhs341_ = d_2041_v117_
                lhs232_ = globalState
                lhs233_ = d_1905_v8_
                lhs234_ = default__.safeIndex(781, (d_1905_v8_).length(0))
                lhs235_ = d_1905_v8_
                lhs236_ = default__.safeIndex(813, (d_1905_v8_).length(0))
                d_2040_v116_ = rhs337_
                lhs232_.f1 = rhs338_
                lhs233_[lhs234_] = rhs339_
                lhs235_[lhs236_] = rhs340_
                d_2041_v117_ = rhs341_
            index326_ = default__.safeIndex(792, (d_1905_v8_).length(0))
            (d_1905_v8_)[index326_] = (d_1902_v5_).f17
            index327_ = default__.safeIndex(792, (d_1905_v8_).length(0))
            (d_1905_v8_)[index327_] = ((d_1902_v5_).f17 if (d_1902_v5_).f17 else not((self).f16))
            (globalState).f1 = (self).f15

    def m7(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: D0 = D0.default()()
        r3: int = int(0)
        hi10_ = (self).f15
        for d_2045_i0_ in range(((self).f15) + ((self).f15), hi10_):
            d_2046_v0_: _dafny.Map
            d_2046_v0_ = _dafny.Map({d_2045_i0_: ((self).f15) >= (default__.fm1(globalState))})
            d_2046_v0_ = (d_2046_v0_).set((self).f15, True)
            d_2047_v1_: C5
            nw310_ = C5()
            nw310_.ctor__()
            d_2047_v1_ = nw310_
            d_2047_v1_ = d_2047_v1_
            d_2048_v2_: _dafny.Array
            nw311_ = _dafny.Array(False, 28)
            d_2048_v2_ = nw311_
            index328_ = default__.safeIndex(13, (d_2048_v2_).length(0))
            (d_2048_v2_)[index328_] = p0
            d_2049_v3_: _dafny.Array
            def lambda98_(d_2050_i1_):
                return D1_DC3((self).f15)

            init51_ = lambda98_
            nw312_ = _dafny.Array(None, 14)
            for i0_51_ in range(nw312_.length(0)):
                nw312_[i0_51_] = init51_(i0_51_)
            d_2049_v3_ = nw312_
            d_2051_v4_: _dafny.Seq
            d_2051_v4_ = _dafny.SeqWithoutIsStrInference([d_2049_v3_])
            d_2052_v5_: _dafny.Seq
            d_2052_v5_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            d_2053_v6_: D8
            d_2053_v6_ = D8_DC27(p1, d_2051_v4_, d_2052_v5_, d_2045_i0_)
            d_2054_v7_: _dafny.Seq
            d_2054_v7_ = _dafny.SeqWithoutIsStrInference([(d_2053_v6_).cf43, (self).f16, p0])
            index329_ = default__.safeIndex(13, (d_2048_v2_).length(0))
            (d_2048_v2_)[index329_] = (p1 if not(p1) else (d_2054_v7_)[default__.safeIndex((self).f15, len(d_2054_v7_))])
            index330_ = default__.safeIndex(13, (d_2048_v2_).length(0))
            (d_2048_v2_)[index330_] = (d_2054_v7_)[default__.safeIndex((self).f15, len(d_2054_v7_))]
        d_2055_v8_: _dafny.Array
        nw313_ = _dafny.Array(_dafny.Array(None, 0), 5)
        d_2055_v8_ = nw313_
        d_2055_v8_ = d_2055_v8_
        d_2056_v9_: _dafny.Set
        d_2056_v9_ = _dafny.Set({(self).f15})
        if (d_2056_v9_).issubset(d_2056_v9_):
            d_2057_v10_: str
            d_2057_v10_ = _dafny.CodePoint('l')
            d_2058_v11_: _dafny.Set
            d_2058_v11_ = _dafny.Set({p1, (self).f16})
            d_2059_v13_: _dafny.Map
            d_2059_v13_ = _dafny.Map({(self).f15: d_2057_v10_})
            d_2060_v14_: D5
            d_2060_v14_ = D5_DC17((self).f16, _dafny.CodePoint('y'))
            d_2061_v15_: _dafny.Seq
            d_2061_v15_ = _dafny.SeqWithoutIsStrInference([d_2060_v14_, D5_DC17(p0, d_2057_v10_)])
            d_2062_v17_: _dafny.Array
            nw314_ = _dafny.Array(None, 15)
            nw314_[int(0)] = (_dafny.Map({(self).f15: d_2057_v10_})).set(len(d_2058_v11_), d_2057_v10_)
            def iife142_():
                coll54_ = _dafny.Map()
                compr_54_: int
                for compr_54_ in _dafny.IntegerRange(145, 430):
                    d_2063_v12_: int = compr_54_
                    if ((145) <= (d_2063_v12_)) and ((d_2063_v12_) < (430)):
                        coll54_[default__.safeModuloInt(d_2063_v12_, (self).f15)] = d_2057_v10_
                return _dafny.Map(coll54_)
            nw314_[int(1)] = iife142_()
            
            nw314_[int(2)] = _dafny.Map({(self).f15: d_2057_v10_})
            nw314_[int(3)] = d_2059_v13_
            nw314_[int(4)] = d_2059_v13_
            nw314_[int(5)] = (d_2059_v13_).set((self).f15, d_2057_v10_)
            nw314_[int(6)] = d_2059_v13_
            nw314_[int(7)] = d_2059_v13_
            nw314_[int(8)] = _dafny.Map({(self).f15: d_2057_v10_})
            nw314_[int(9)] = default__.fm50(p0, len(d_2061_v15_), globalState)
            nw314_[int(10)] = d_2059_v13_
            def iife143_():
                coll55_ = _dafny.Map()
                compr_55_: int
                for compr_55_ in _dafny.IntegerRange(12, -997):
                    d_2064_v16_: int = compr_55_
                    if ((12) <= (d_2064_v16_)) and ((d_2064_v16_) < (-997)):
                        coll55_[(d_2064_v16_) + ((self).f15)] = d_2057_v10_
                return _dafny.Map(coll55_)
            nw314_[int(11)] = iife143_()
            
            nw314_[int(12)] = _dafny.Map({(self).f15: d_2057_v10_})
            nw314_[int(13)] = d_2059_v13_
            nw314_[int(14)] = default__.fm50((D5_DC17((self).f16, d_2057_v10_)).cf28, 909, globalState)
            d_2062_v17_ = nw314_
            d_2065_v18_: _dafny.Map
            d_2065_v18_ = _dafny.Map({not(p0): d_2062_v17_})
            d_2066_v19_: _dafny.Array
            nw315_ = _dafny.Array(None, 15)
            nw315_[int(0)] = d_2062_v17_
            nw315_[int(1)] = d_2062_v17_
            nw315_[int(2)] = d_2062_v17_
            nw315_[int(3)] = (d_2062_v17_ if (self).f16 else d_2062_v17_)
            nw315_[int(4)] = (d_2062_v17_ if p0 else d_2062_v17_)
            nw315_[int(5)] = d_2062_v17_
            nw315_[int(6)] = d_2062_v17_
            nw315_[int(7)] = d_2062_v17_
            nw315_[int(8)] = d_2062_v17_
            nw315_[int(9)] = ((d_2065_v18_)[p0] if (p0) in (d_2065_v18_) else d_2062_v17_)
            nw315_[int(10)] = d_2062_v17_
            nw315_[int(11)] = d_2062_v17_
            nw315_[int(12)] = d_2062_v17_
            nw315_[int(13)] = d_2062_v17_
            nw315_[int(14)] = d_2062_v17_
            d_2066_v19_ = nw315_
            index331_ = default__.safeIndex(10, (d_2066_v19_).length(0))
            (d_2066_v19_)[index331_] = d_2062_v17_
            index332_ = default__.safeIndex(10, (d_2066_v19_).length(0))
            (d_2066_v19_)[index332_] = d_2062_v17_
            d_2067_v20_: D17
            d_2067_v20_ = D17_DC50(d_2055_v8_)
            d_2068_v21_: D17
            d_2068_v21_ = D17_DC52(d_2067_v20_)
            d_2069_v22_: D17
            d_2069_v22_ = D17_DC52((d_2067_v20_ if p1 else d_2067_v20_))
            d_2069_v22_ = d_2069_v22_
            d_2070_v23_: _dafny.Array
            nw316_ = _dafny.Array(None, 18)
            nw316_[int(0)] = d_2058_v11_
            nw316_[int(1)] = _dafny.Set({p1})
            nw316_[int(2)] = d_2058_v11_
            nw316_[int(3)] = d_2058_v11_
            nw316_[int(4)] = d_2058_v11_
            nw316_[int(5)] = d_2058_v11_
            nw316_[int(6)] = d_2058_v11_
            nw316_[int(7)] = d_2058_v11_
            nw316_[int(8)] = d_2058_v11_
            nw316_[int(9)] = d_2058_v11_
            nw316_[int(10)] = d_2058_v11_
            nw316_[int(11)] = d_2058_v11_
            nw316_[int(12)] = d_2058_v11_
            nw316_[int(13)] = d_2058_v11_
            nw316_[int(14)] = _dafny.Set({(self).f16})
            nw316_[int(15)] = d_2058_v11_
            nw316_[int(16)] = d_2058_v11_
            nw316_[int(17)] = d_2058_v11_
            d_2070_v23_ = nw316_
            d_2071_v24_: _dafny.Array
            nw317_ = _dafny.Array(int(0), 2)
            d_2071_v24_ = nw317_
            d_2072_v25_: _dafny.Set
            d_2072_v25_ = _dafny.Set({d_2071_v24_, d_2071_v24_, d_2071_v24_})
            d_2073_v26_: _dafny.Seq
            d_2073_v26_ = _dafny.SeqWithoutIsStrInference([d_2072_v25_, _dafny.Set({d_2071_v24_, d_2071_v24_}), d_2072_v25_])
            d_2074_v27_: _dafny.Map
            d_2074_v27_ = _dafny.Map({d_2070_v23_: (d_2073_v26_)[default__.safeIndex((self).f15, len(d_2073_v26_))]})
            r0 = (d_2072_v25_).ispropersubset(((d_2074_v27_)[d_2070_v23_] if (d_2070_v23_) in (d_2074_v27_) else d_2072_v25_))
            d_2075_v28_: _dafny.Seq
            d_2075_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtpb"))
            d_2076_v29_: _dafny.Map
            d_2076_v29_ = _dafny.Map({p1: (self).f16})
            d_2077_v30_: _dafny.Set
            d_2077_v30_ = _dafny.Set({d_2076_v29_})
            (globalState).f7 = (self).fm9(d_2075_v28_, d_2077_v30_, globalState)
            (globalState).f1 = (self).f15
        elif True:
            d_2078_v31_: _dafny.Map
            d_2078_v31_ = _dafny.Map({479: (self).f15})
            d_2078_v31_ = (d_2078_v31_).set((self).f15, default__.fm1(globalState))
            r1 = p1
            r3 = (self).f15
            d_2079_v32_: _dafny.Array
            nw318_ = _dafny.Array(D1.default()(), 27)
            d_2079_v32_ = nw318_
            d_2080_v33_: _dafny.Seq
            d_2080_v33_ = _dafny.SeqWithoutIsStrInference([d_2079_v32_])
            d_2081_v34_: _dafny.Seq
            d_2081_v34_ = _dafny.SeqWithoutIsStrInference([(self).f15, len(d_2056_v9_), (self).f15])
            d_2082_v35_: D8
            d_2082_v35_ = D8_DC27(p1, d_2080_v33_, d_2081_v34_, 726)
            d_2083_v36_: _dafny.Seq
            d_2083_v36_ = _dafny.SeqWithoutIsStrInference([(self).f16, default__.fm0((self).f15, globalState), ((d_2082_v35_).cf43) and (p0), ((self).f15) == ((self).f15), p1])
            r0 = (d_2083_v36_)[default__.safeIndex((self).f15, len(d_2083_v36_))]
            d_2084_v37_: _dafny.Array
            def lambda99_(d_2085_p1_):
                def lambda100_(d_2086_i2_):
                    return d_2085_p1_

                return lambda100_

            init52_ = lambda99_(p1)
            nw319_ = _dafny.Array(None, 4)
            for i0_52_ in range(nw319_.length(0)):
                nw319_[i0_52_] = init52_(i0_52_)
            d_2084_v37_ = nw319_
            d_2087_v38_: _dafny.Map
            d_2087_v38_ = _dafny.Map({p0: (D0_DC0(d_2084_v37_)).cf0})
            d_2087_v38_ = (d_2087_v38_).set(p1, d_2084_v37_)
        d_2088_v39_: C0
        nw320_ = C0()
        nw320_.ctor__()
        d_2088_v39_ = nw320_
        r1 = (self).f16
        d_2089_v40_: _dafny.Map
        d_2089_v40_ = _dafny.Map({not((d_2088_v39_).fm37(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")), (self).f15, p1, (self).f16, globalState)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icicg"))})
        d_2089_v40_ = (d_2089_v40_).set((self).f16, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfbbxc")))
        r0 = (self).f16
        d_2090_v41_: _dafny.Map
        d_2090_v41_ = _dafny.Map({(self).f15: not(not(p1))})
        r1 = ((d_2090_v41_)[(self).f15] if ((self).f15) in (d_2090_v41_) else False)
        d_2091_v42_: _dafny.Map
        d_2091_v42_ = _dafny.Map({(self).f15: (self).f15})
        d_2092_v43_: str
        d_2092_v43_ = _dafny.CodePoint('x')
        d_2093_v44_: D0
        d_2093_v44_ = D0_DC1((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwthpjsf"))).set(default__.safeIndex(len(d_2091_v42_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwthpjsf")))), d_2092_v43_))
        r2 = d_2093_v44_
        r3 = ((self).f15) - ((self).f15)
        return r0, r1, r2, r3

    @property
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16

class C11(T0, T1):
    def  __init__(self):
        self._f10: bool = False
        self._f13: int = int(0)
        self._f14: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f13, f14, f10):
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f10 = f10

    def fm6(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Set({(self).f10, (self).f10})) - (_dafny.Set({(self).f10, (self).f10}))).isdisjoint((_dafny.Set({(self).f10})) - (_dafny.Set({not(True)})))

    def fm7(self, p0, globalState):
        return (self).f10

    def fm8(self, p0, p1, p2, globalState):
        source36_ = D3_DC10((self).f14, (self).f10, (self).f10)
        if source36_.is_DC10:
            d_2094___mcc_h0_ = source36_.cf14
            d_2095___mcc_h1_ = source36_.cf15
            d_2096___mcc_h2_ = source36_.cf16
            d_2097_cf16_ = d_2096___mcc_h2_
            d_2098_cf15_ = d_2095___mcc_h1_
            d_2099_cf14_ = d_2094___mcc_h0_
            return (self).f10
        elif source36_.is_DC11:
            d_2100___mcc_h3_ = source36_.cf17
            d_2101___mcc_h4_ = source36_.cf18
            d_2102_cf18_ = d_2101___mcc_h4_
            d_2103_cf17_ = d_2100___mcc_h3_
            return (self).f10
        elif source36_.is_DC9:
            d_2104___mcc_h5_ = source36_.cf13
            d_2105_cf13_ = d_2104___mcc_h5_
            return True
        elif True:
            d_2106___mcc_h6_ = source36_.cf19
            d_2107_cf19_ = d_2106___mcc_h6_
            return (self).f10

    def fm9(self, p0, p1, globalState):
        source37_ = D2_DC6(_dafny.CodePoint('b'))
        if source37_.is_DC7:
            d_2108___mcc_h0_ = source37_.cf9
            d_2109___mcc_h1_ = source37_.cf10
            d_2110___mcc_h2_ = source37_.cf11
            d_2111_cf11_ = d_2110___mcc_h2_
            d_2112_cf10_ = d_2109___mcc_h1_
            d_2113_cf9_ = d_2108___mcc_h0_
            return (len(_dafny.SeqWithoutIsStrInference([(self).f10]))) - (d_2111_cf11_)
        elif source37_.is_DC6:
            d_2114___mcc_h3_ = source37_.cf8
            d_2115_cf8_ = d_2114___mcc_h3_
            return (self).f13
        elif True:
            d_2116___mcc_h4_ = source37_.cf12
            d_2117_cf12_ = d_2116___mcc_h4_
            return default__.safeModuloInt(627, (self).f13)

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Map = _dafny.Map({})
        r0 = p1
        d_2118_v0_: _dafny.Array
        def lambda101_(d_2119_i0_):
            return _dafny.Map({D0_DC1((self).f14): (self).f14})

        init53_ = lambda101_
        nw321_ = _dafny.Array(None, 14)
        for i0_53_ in range(nw321_.length(0)):
            nw321_[i0_53_] = init53_(i0_53_)
        d_2118_v0_ = nw321_
        d_2120_v1_: D0
        d_2120_v1_ = D0_DC1((self).f14)
        d_2121_v2_: _dafny.Map
        d_2121_v2_ = _dafny.Map({d_2120_v1_: (self).f14})
        index333_ = default__.safeIndex(289, (d_2118_v0_).length(0))
        (d_2118_v0_)[index333_] = (d_2121_v2_) | (d_2121_v2_)
        index334_ = default__.safeIndex(289, (d_2118_v0_).length(0))
        (d_2118_v0_)[index334_] = d_2121_v2_
        d_2122_v3_: _dafny.Array
        nw322_ = _dafny.Array(int(0), 2)
        d_2122_v3_ = nw322_
        d_2123_v4_: D3
        d_2123_v4_ = D3_DC10((self).f14, p1, (self).f10)
        d_2124_v5_: D3
        d_2124_v5_ = D3_DC12(d_2123_v4_)
        d_2125_v6_: _dafny.Map
        d_2125_v6_ = _dafny.Map({(self).f10: p1})
        d_2126_v7_: _dafny.Set
        d_2126_v7_ = _dafny.Set({d_2125_v6_, d_2125_v6_})
        d_2127_v8_: _dafny.Map
        d_2127_v8_ = _dafny.Map({(self).fm9((self).f14, d_2126_v7_, globalState): (self).f10})
        d_2128_v9_: _dafny.Map
        d_2128_v9_ = _dafny.Map({D3_DC12(d_2123_v4_): (d_2127_v8_).set((0) - ((self).f13), p1)})
        d_2129_v10_: D3
        d_2129_v10_ = D3_DC12(d_2124_v5_)
        index335_ = default__.safeIndex(539, (d_2122_v3_).length(0))
        (d_2122_v3_)[index335_] = (len((d_2128_v9_).set(d_2129_v10_, d_2127_v8_)) if p1 else (self).f13)
        index336_ = default__.safeIndex(539, (d_2122_v3_).length(0))
        (d_2122_v3_)[index336_] = (self).f13
        hi11_ = (0) - ((0) - (((self).f13) + (135)))
        for d_2130_i1_ in range(((d_2122_v3_)[default__.safeIndex(539, (d_2122_v3_).length(0))]) * ((0) - ((self).f13)), hi11_):
            d_2131_v11_: _dafny.Set
            d_2131_v11_ = _dafny.Set({(d_2122_v3_)[default__.safeIndex(539, (d_2122_v3_).length(0))]})
            d_2131_v11_ = ((d_2131_v11_) - (_dafny.Set({(D2_DC7(d_2130_i1_, True, (d_2122_v3_)[default__.safeIndex(539, (d_2122_v3_).length(0))])).cf11, 0, 241, d_2130_i1_, (self).f13}))) - (d_2131_v11_)
            d_2132_v12_: C2
            nw323_ = C2()
            nw323_.ctor__(p1, d_2130_i1_)
            d_2132_v12_ = nw323_
            r0 = not (p1) or (p1)
            if (self).f10:
                d_2133_v13_: _dafny.MultiSet
                d_2133_v13_ = _dafny.MultiSet([(d_2132_v12_).f21])
                d_2134_v14_: D5
                d_2134_v14_ = D5_DC16(d_2133_v13_)
                d_2135_v15_: _dafny.Array
                nw324_ = _dafny.Array(None, 11)
                nw324_[int(0)] = d_2134_v14_
                nw324_[int(1)] = d_2134_v14_
                nw324_[int(2)] = d_2134_v14_
                nw324_[int(3)] = D5_DC16(d_2133_v13_)
                nw324_[int(4)] = D5_DC16(d_2133_v13_)
                nw324_[int(5)] = d_2134_v14_
                nw324_[int(6)] = d_2134_v14_
                nw324_[int(7)] = d_2134_v14_
                nw324_[int(8)] = D5_DC16(d_2133_v13_)
                nw324_[int(9)] = d_2134_v14_
                nw324_[int(10)] = d_2134_v14_
                d_2135_v15_ = nw324_
                index337_ = default__.safeIndex(117, (d_2135_v15_).length(0))
                (d_2135_v15_)[index337_] = d_2134_v14_
                index338_ = default__.safeIndex(117, (d_2135_v15_).length(0))
                (d_2135_v15_)[index338_] = d_2134_v14_
                d_2136_v16_: _dafny.Array
                nw325_ = _dafny.Array(D16.default()(), 9)
                d_2136_v16_ = nw325_
                d_2137_v17_: _dafny.Array
                nw326_ = _dafny.Array(D14.default()(), 23)
                d_2137_v17_ = nw326_
                d_2138_v18_: D16
                d_2138_v18_ = D16_DC48(d_2137_v17_)
                index339_ = default__.safeIndex(423, (d_2136_v16_).length(0))
                (d_2136_v16_)[index339_] = d_2138_v18_
                index340_ = default__.safeIndex(423, (d_2136_v16_).length(0))
                (d_2136_v16_)[index340_] = d_2138_v18_
                d_2139_v19_: _dafny.Seq
                d_2139_v19_ = _dafny.SeqWithoutIsStrInference([(d_2132_v12_).f21, ((d_2127_v8_)[(d_2132_v12_).f22] if ((d_2132_v12_).f22) in (d_2127_v8_) else (d_2132_v12_).f21)])
                d_2140_v20_: _dafny.Map
                d_2140_v20_ = _dafny.Map({d_2139_v19_: d_2131_v11_})
                d_2140_v20_ = (d_2140_v20_).set((_dafny.SeqWithoutIsStrInference([p1, True, (self).f10, (d_2132_v12_).f21])) + (_dafny.SeqWithoutIsStrInference([(d_2132_v12_).f21])), (d_2131_v11_).intersection(_dafny.Set({d_2130_i1_, (self).f13, (d_2122_v3_)[default__.safeIndex(539, (d_2122_v3_).length(0))]})))
                d_2141_v21_: D19
                d_2141_v21_ = D19_DC58((d_2130_i1_) - ((self).f13), (d_2132_v12_).f21)
                d_2141_v21_ = d_2141_v21_
                (globalState).f1 = (self).f13
            elif True:
                d_2142_v22_: D19
                d_2142_v22_ = D19_DC57(d_2131_v11_)
                d_2143_v23_: _dafny.Seq
                d_2143_v23_ = _dafny.SeqWithoutIsStrInference([d_2142_v22_])
                d_2144_v24_: D2
                d_2144_v24_ = D2_DC7((self).f13, (d_2132_v12_).f21, 245)
                d_2145_v25_: D2
                d_2145_v25_ = D2_DC8(d_2144_v24_)
                d_2146_v26_: _dafny.Map
                d_2146_v26_ = _dafny.Map({len(d_2143_v23_): d_2145_v25_})
                d_2147_v27_: str
                d_2147_v27_ = _dafny.CodePoint('r')
                r1 = ((self).f14).set(default__.safeIndex((len(d_2146_v26_)) * ((self).f13), len((self).f14)), d_2147_v27_)
                d_2148_v28_: _dafny.Array
                def lambda102_(d_2149_i2_):
                    return (self).f10

                init54_ = lambda102_
                nw327_ = _dafny.Array(None, 22)
                for i0_54_ in range(nw327_.length(0)):
                    nw327_[i0_54_] = init54_(i0_54_)
                d_2148_v28_ = nw327_
                index341_ = default__.safeIndex(395, (d_2148_v28_).length(0))
                (d_2148_v28_)[index341_] = p1
                index342_ = default__.safeIndex(395, (d_2148_v28_).length(0))
                (d_2148_v28_)[index342_] = p1
                (globalState).f0 = d_2130_i1_
                def iife144_():
                    def iife146_():
                        coll58_ = _dafny.Map()
                        compr_58_: int
                        for compr_58_ in (d_2131_v11_).Elements:
                            d_2150_v30_: int = compr_58_
                            if (d_2150_v30_) in (d_2131_v11_):
                                coll58_[(d_2150_v30_) + ((d_2132_v12_).f22)] = _dafny.SeqWithoutIsStrInference([(self).f10])
                        return _dafny.Map(coll58_)
                    coll56_ = _dafny.Map()
                    def iife145_():
                        coll57_ = _dafny.Map()
                        compr_57_: int
                        for compr_57_ in (d_2131_v11_).Elements:
                            d_2150_v30_: int = compr_57_
                            if (d_2150_v30_) in (d_2131_v11_):
                                coll57_[(d_2150_v30_) + ((d_2132_v12_).f22)] = _dafny.SeqWithoutIsStrInference([(self).f10])
                        return _dafny.Map(coll57_)
                    compr_56_: int
                    for compr_56_ in (iife145_()
                    ).keys.Elements:
                        d_2151_v29_: int = compr_56_
                        if (d_2151_v29_) in (iife146_()
                        ):
                            coll56_[default__.safeModuloInt(d_2151_v29_, (0) - (len((self).f14)))] = 578
                    return _dafny.Map(coll56_)
                (globalState).f1 = len(iife144_()
                )
                r0 = not (p1) or (not(False))
        r0 = p1
        hi12_ = default__.safeModuloInt((0) - ((d_2122_v3_)[default__.safeIndex(539, (d_2122_v3_).length(0))]), (0) - (len((self).f14)))
        for d_2152_i3_ in range((d_2122_v3_)[default__.safeIndex(539, (d_2122_v3_).length(0))], hi12_):
            index343_ = default__.safeIndex(539, (d_2122_v3_).length(0))
            (d_2122_v3_)[index343_] = ((self).f13) * (default__.safeDivisionInt((default__.fm41(globalState)).cardinality, (d_2122_v3_)[default__.safeIndex(539, (d_2122_v3_).length(0))]))
            d_2153_v31_: _dafny.Array
            def lambda103_(d_2154_p1_):
                def lambda104_(d_2155_i4_):
                    return _dafny.Set({(self).f10, d_2154_p1_})

                return lambda104_

            init55_ = lambda103_(p1)
            nw328_ = _dafny.Array(None, 16)
            for i0_55_ in range(nw328_.length(0)):
                nw328_[i0_55_] = init55_(i0_55_)
            d_2153_v31_ = nw328_
            d_2156_v32_: _dafny.Map
            d_2156_v32_ = _dafny.Map({d_2153_v31_: d_2152_i3_})
            d_2156_v32_ = (d_2156_v32_).set(d_2153_v31_, (d_2122_v3_)[default__.safeIndex(539, (d_2122_v3_).length(0))])
            d_2157_v33_: _dafny.MultiSet
            d_2157_v33_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([p1])) + ((_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])).set(default__.safeIndex((self).f13, len(_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10]))), (self).f10))])
            d_2157_v33_ = d_2157_v33_
            d_2127_v8_ = (d_2127_v8_).set(d_2152_i3_, p1)
        r0 = p1
        r1 = ((self).f14) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_2158_i5_ in range(default__.abs(961))])) + ((self).f14))
        d_2159_v34_: _dafny.Map
        d_2159_v34_ = _dafny.Map({212: (d_2122_v3_)[default__.safeIndex(539, (d_2122_v3_).length(0))]})
        r2 = d_2159_v34_
        return r0, r1, r2

    def m4(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: int = int(0)
        hi13_ = ((self).f13 if (self).f10 else (self).f13)
        for d_2160_i0_ in range((self).f13, hi13_):
            d_2161_v0_: _dafny.Seq
            d_2161_v0_ = _dafny.SeqWithoutIsStrInference([(self).f13, d_2160_i0_])
            d_2162_v1_: _dafny.Set
            d_2162_v1_ = _dafny.Set({len(d_2161_v0_), (self).f13})
            d_2163_v2_: _dafny.Set
            d_2163_v2_ = _dafny.Set({(self).f10})
            d_2164_v3_: _dafny.Map
            d_2164_v3_ = _dafny.Map({(self).f10: len(d_2163_v2_)})
            d_2165_v4_: _dafny.Array
            nw329_ = _dafny.Array(D2.default()(), 4)
            d_2165_v4_ = nw329_
            d_2166_v5_: _dafny.Map
            d_2166_v5_ = _dafny.Map({(0) - ((self).f13): d_2165_v4_})
            d_2167_v6_: D6
            d_2167_v6_ = D6_DC19(d_2166_v5_)
            d_2168_v7_: _dafny.Map
            d_2168_v7_ = _dafny.Map({(_dafny.MultiSet([(self).f13, len(d_2162_v1_)])).issubset(default__.fm18((self).f13, d_2164_v3_, (self).f10, globalState)): d_2167_v6_})
            d_2168_v7_ = (d_2168_v7_).set(((self).f13) > (d_2160_i0_), d_2167_v6_)
            d_2169_v8_: _dafny.Array
            def lambda105_(d_2170_i0_):
                def lambda106_(d_2171_i1_):
                    return (d_2171_i1_) + (d_2170_i0_)

                return lambda106_

            init56_ = lambda105_(d_2160_i0_)
            nw330_ = _dafny.Array(None, 13)
            for i0_56_ in range(nw330_.length(0)):
                nw330_[i0_56_] = init56_(i0_56_)
            d_2169_v8_ = nw330_
            index344_ = default__.safeIndex(482, (d_2169_v8_).length(0))
            (d_2169_v8_)[index344_] = d_2160_i0_
            index345_ = default__.safeIndex(482, (d_2169_v8_).length(0))
            (d_2169_v8_)[index345_] = ((self).f13) - (d_2160_i0_)
            r1 = (self).f10
            d_2172_v9_: _dafny.MultiSet
            d_2172_v9_ = _dafny.MultiSet([(self).f13])
            d_2173_v10_: _dafny.Array
            nw331_ = _dafny.Array(None, 11)
            nw331_[int(0)] = False
            nw331_[int(1)] = (self).f10
            nw331_[int(2)] = (self).f10
            nw331_[int(3)] = (self).f10
            nw331_[int(4)] = not((d_2160_i0_) > (d_2160_i0_))
            nw331_[int(5)] = False
            nw331_[int(6)] = (d_2172_v9_).issubset(d_2172_v9_)
            nw331_[int(7)] = False
            nw331_[int(8)] = (self).f10
            nw331_[int(9)] = (self).f10
            nw331_[int(10)] = (self).f10
            d_2173_v10_ = nw331_
            index346_ = default__.safeIndex(690, (d_2173_v10_).length(0))
            (d_2173_v10_)[index346_] = (self).f10
            d_2174_v11_: D13
            d_2174_v11_ = D13_DC39(((d_2164_v3_)[(self).f10] if ((self).f10) in (d_2164_v3_) else (d_2169_v8_)[default__.safeIndex(482, (d_2169_v8_).length(0))]), not((self).f10))
            index347_ = default__.safeIndex(690, (d_2173_v10_).length(0))
            (d_2173_v10_)[index347_] = not(((d_2174_v11_ if (self).f10 else d_2174_v11_)).cf66)
        d_2175_v12_: _dafny.Array
        nw332_ = _dafny.Array(_dafny.Array(None, 0), 16)
        d_2175_v12_ = nw332_
        d_2176_v13_: _dafny.Array
        nw333_ = _dafny.Array(None, 15)
        nw333_[int(0)] = (self).f10
        nw333_[int(1)] = (self).f10
        nw333_[int(2)] = True
        nw333_[int(3)] = (self).f10
        nw333_[int(4)] = (self).f10
        nw333_[int(5)] = (self).f10
        nw333_[int(6)] = (self).f10
        nw333_[int(7)] = (self).f10
        nw333_[int(8)] = (self).f10
        nw333_[int(9)] = (self).f10
        nw333_[int(10)] = (self).f10
        nw333_[int(11)] = (self).f10
        nw333_[int(12)] = (self).f10
        nw333_[int(13)] = (self).f10
        nw333_[int(14)] = (self).f10
        d_2176_v13_ = nw333_
        index348_ = default__.safeIndex(909, (d_2175_v12_).length(0))
        (d_2175_v12_)[index348_] = d_2176_v13_
        index349_ = default__.safeIndex(909, (d_2175_v12_).length(0))
        (d_2175_v12_)[index349_] = d_2176_v13_
        d_2177_v14_: D10
        d_2177_v14_ = D10_DC30(_dafny.Set({(self).f10, False}))
        source38_ = d_2177_v14_
        if source38_.is_DC31:
            d_2178___mcc_h0_ = source38_.cf50
            d_2179___mcc_h1_ = source38_.cf51
            d_2180___mcc_h2_ = source38_.cf52
            d_2181_cf52_ = d_2180___mcc_h2_
            d_2182_cf51_ = d_2179___mcc_h1_
            d_2183_cf50_ = d_2178___mcc_h0_
            d_2184_v15_: _dafny.Array
            nw334_ = _dafny.Array(int(0), 25)
            d_2184_v15_ = nw334_
            index350_ = default__.safeIndex(116, (d_2184_v15_).length(0))
            (d_2184_v15_)[index350_] = (self).f13
            index351_ = default__.safeIndex(116, (d_2184_v15_).length(0))
            (d_2184_v15_)[index351_] = d_2181_cf52_
            d_2185_v16_: D0
            d_2185_v16_ = D0_DC0((D0_DC2(False, (d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))])).cf3)
            d_2186_v17_: _dafny.Seq
            d_2186_v17_ = _dafny.SeqWithoutIsStrInference([(d_2185_v16_).cf0, d_2176_v13_, d_2176_v13_])
            d_2187_v18_: C10
            nw335_ = C10()
            nw335_.ctor__(len(d_2186_v17_), d_2183_cf50_)
            d_2187_v18_ = nw335_
            (globalState).f7 = default__.safeModuloInt((d_2187_v18_).f15, default__.safeModuloInt((d_2187_v18_).f15, (self).f13))
            if (d_2187_v18_).f16:
                d_2188_v19_: _dafny.Array
                nw336_ = _dafny.Array(_dafny.CodePoint('D'), 16)
                d_2188_v19_ = nw336_
                d_2189_v20_: str
                d_2189_v20_ = _dafny.CodePoint('k')
                index352_ = default__.safeIndex(138, (d_2188_v19_).length(0))
                (d_2188_v19_)[index352_] = d_2189_v20_
                index353_ = default__.safeIndex(138, (d_2188_v19_).length(0))
                (d_2188_v19_)[index353_] = d_2189_v20_
                d_2190_v21_: C5
                nw337_ = C5()
                nw337_.ctor__()
                d_2190_v21_ = nw337_
                index354_ = default__.safeIndex(614, (d_2176_v13_).length(0))
                (d_2176_v13_)[index354_] = True
                index355_ = default__.safeIndex(614, (d_2176_v13_).length(0))
                (d_2176_v13_)[index355_] = d_2182_cf51_
                d_2191_v22_: _dafny.Map
                d_2191_v22_ = _dafny.Map({d_2189_v20_: (d_2187_v18_).f15})
                d_2192_v23_: _dafny.Set
                d_2192_v23_ = _dafny.Set({(d_2187_v18_).f16, d_2182_cf51_})
                d_2191_v22_ = (d_2191_v22_).set((d_2188_v19_)[default__.safeIndex(138, (d_2188_v19_).length(0))], len(d_2192_v23_))
                index356_ = default__.safeIndex(138, (d_2188_v19_).length(0))
                (d_2188_v19_)[index356_] = d_2189_v20_
            elif True:
                r1 = (d_2187_v18_).f16
                d_2193_v24_: _dafny.Array
                def lambda107_(d_2194_cf50_, d_2195_v18_):
                    def lambda108_(d_2196_i2_):
                        return (_dafny.Map({(self).f10: _dafny.Map({False: _dafny.Map({_dafny.Map({d_2194_cf50_: _dafny.CodePoint('h')}): (d_2195_v18_).f16})})})) | (_dafny.Map({(d_2195_v18_).f16: _dafny.Map({(d_2195_v18_).f16: _dafny.Map({_dafny.Map({True: _dafny.CodePoint('c')}): (d_2195_v18_).f16})})}))

                    return lambda108_

                init57_ = lambda107_(d_2183_cf50_, d_2187_v18_)
                nw338_ = _dafny.Array(None, 12)
                for i0_57_ in range(nw338_.length(0)):
                    nw338_[i0_57_] = init57_(i0_57_)
                d_2193_v24_ = nw338_
                d_2197_v25_: _dafny.Set
                d_2197_v25_ = _dafny.Set({(self).f10, (self).f10})
                d_2198_v26_: str
                d_2198_v26_ = _dafny.CodePoint('p')
                d_2199_v27_: _dafny.Map
                d_2199_v27_ = _dafny.Map({len(d_2197_v25_): d_2198_v26_})
                d_2200_v28_: _dafny.Map
                d_2200_v28_ = _dafny.Map({(self).f10: ((d_2199_v27_)[(d_2187_v18_).f15] if ((d_2187_v18_).f15) in (d_2199_v27_) else d_2198_v26_)})
                d_2201_v29_: _dafny.Map
                d_2201_v29_ = _dafny.Map({d_2200_v28_: d_2182_cf51_})
                d_2202_v30_: _dafny.Map
                d_2202_v30_ = d_2201_v29_
                d_2203_v31_: _dafny.Map
                d_2203_v31_ = _dafny.Map({(d_2187_v18_).f16: d_2202_v30_})
                d_2204_v32_: _dafny.Map
                d_2204_v32_ = _dafny.Map({(d_2187_v18_).f16: d_2203_v31_})
                index357_ = default__.safeIndex(437, (d_2193_v24_).length(0))
                (d_2193_v24_)[index357_] = d_2204_v32_
                d_2205_v33_: _dafny.MultiSet
                d_2205_v33_ = _dafny.MultiSet([(d_2184_v15_)[default__.safeIndex(116, (d_2184_v15_).length(0))]])
                d_2206_v34_: _dafny.MultiSet
                d_2206_v34_ = _dafny.MultiSet([(d_2205_v33_).cardinality, (d_2187_v18_).f15])
                d_2207_v35_: D4
                d_2207_v35_ = D4_DC15((d_2187_v18_).f15, d_2205_v33_, _dafny.CodePoint('u'), (d_2187_v18_).f15)
                d_2208_v36_: _dafny.Array
                nw339_ = _dafny.Array(None, 7)
                nw339_[int(0)] = _dafny.CodePoint('i')
                nw339_[int(1)] = d_2198_v26_
                nw339_[int(2)] = d_2198_v26_
                nw339_[int(3)] = (d_2207_v35_).cf25
                nw339_[int(4)] = d_2198_v26_
                nw339_[int(5)] = d_2198_v26_
                nw339_[int(6)] = (_dafny.CodePoint('j') if True else d_2198_v26_)
                d_2208_v36_ = nw339_
                d_2209_v37_: _dafny.Map
                d_2209_v37_ = _dafny.Map({d_2184_v15_: (d_2187_v18_).f16})
                index358_ = default__.safeIndex(437, (d_2193_v24_).length(0))
                rhs342_ = ((d_2209_v37_)[d_2184_v15_] if (d_2184_v15_) in (d_2209_v37_) else d_2182_cf51_)
                rhs343_ = (self).fm8((d_2187_v18_).f15, (d_2187_v18_).f16, not (d_2182_cf51_) or (d_2182_cf51_), globalState)
                rhs344_ = (d_2204_v32_) | (d_2204_v32_)
                rhs345_ = d_2208_v36_
                lhs237_ = d_2193_v24_
                lhs238_ = default__.safeIndex(437, (d_2193_v24_).length(0))
                d_2183_cf50_ = rhs342_
                d_2183_cf50_ = rhs343_
                lhs237_[lhs238_] = rhs344_
                d_2208_v36_ = rhs345_
                (globalState).f1 = d_2181_cf52_
                d_2210_v38_: _dafny.Seq
                d_2210_v38_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                d_2211_v39_: _dafny.Map
                d_2211_v39_ = _dafny.Map({d_2183_cf50_: d_2183_cf50_})
                d_2212_v40_: _dafny.Set
                d_2212_v40_ = _dafny.Set({_dafny.Map({d_2182_cf51_: d_2182_cf51_}), _dafny.Map({True: (d_2187_v18_).f16}), d_2211_v39_, _dafny.Map({(d_2187_v18_).f16: d_2182_cf51_}), d_2211_v39_})
                index359_ = default__.safeIndex(116, (d_2184_v15_).length(0))
                (d_2184_v15_)[index359_] = (self).fm9((d_2210_v38_)[default__.safeIndex(d_2181_cf52_, len(d_2210_v38_))], (d_2212_v40_).intersection(_dafny.Set({d_2211_v39_, (d_2211_v39_).set(d_2183_cf50_, (d_2187_v18_).f16), _dafny.Map({d_2182_cf51_: (self).f10}), d_2211_v39_})), globalState)
                arr6_ = (d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]
                index360_ = default__.safeIndex(590, ((d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]).length(0))
                arr6_[index360_] = d_2182_cf51_
                d_2213_v41_: _dafny.Array
                nw340_ = _dafny.Array(_dafny.Seq({}), 19)
                d_2213_v41_ = nw340_
                d_2214_v42_: _dafny.Seq
                d_2214_v42_ = _dafny.SeqWithoutIsStrInference([(d_2187_v18_).f16])
                index361_ = default__.safeIndex(925, (d_2213_v41_).length(0))
                (d_2213_v41_)[index361_] = (d_2214_v42_) + (d_2214_v42_)
                arr7_ = (d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]
                index362_ = default__.safeIndex(220, ((d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]).length(0))
                arr7_[index362_] = d_2182_cf51_
                d_2215_v43_: D11
                d_2215_v43_ = D11_DC34((d_2187_v18_).f15, True, d_2181_cf52_)
                d_2216_v44_: _dafny.Map
                d_2216_v44_ = _dafny.Map({(d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]: d_2183_cf50_})
                d_2217_v45_: C2
                nw341_ = C2()
                nw341_.ctor__(d_2182_cf51_, (d_2187_v18_).f15)
                d_2217_v45_ = nw341_
                d_2218_v46_: _dafny.Seq
                d_2218_v46_ = _dafny.SeqWithoutIsStrInference([d_2217_v45_])
                pat_let_tv60_ = d_2175_v12_
                pat_let_tv61_ = d_2175_v12_
                pat_let_tv62_ = d_2176_v13_
                pat_let_tv63_ = d_2187_v18_
                arr8_ = (d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]
                index363_ = default__.safeIndex(590, ((d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]).length(0))
                index364_ = default__.safeIndex(925, (d_2213_v41_).length(0))
                arr9_ = (d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]
                index365_ = default__.safeIndex(220, ((d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]).length(0))
                def iife148_(_pat_let45_0):
                    def iife149_(d_2219_dt__update__tmp_h0_):
                        def iife150_(_pat_let46_0):
                            def iife151_(d_2220_dt__update_hcf58_h0_):
                                return D11_DC34((d_2219_dt__update__tmp_h0_).cf56, (d_2219_dt__update__tmp_h0_).cf57, d_2220_dt__update_hcf58_h0_)
                            return iife151_(_pat_let46_0)
                        return iife150_(len(_dafny.Set({(pat_let_tv61_)[default__.safeIndex(909, (pat_let_tv60_).length(0))], pat_let_tv62_})))
                    return iife149_(_pat_let45_0)
                def iife147_(_pat_let44_0):
                    def iife152_(d_2221_dt__update__tmp_h1_):
                        def iife153_(_pat_let47_0):
                            def iife154_(d_2222_dt__update_hcf57_h0_):
                                def iife155_(_pat_let48_0):
                                    def iife156_(d_2223_dt__update_hcf56_h0_):
                                        return D11_DC34(d_2223_dt__update_hcf56_h0_, d_2222_dt__update_hcf57_h0_, (d_2221_dt__update__tmp_h1_).cf58)
                                    return iife156_(_pat_let48_0)
                                return iife155_((self).f13)
                            return iife154_(_pat_let47_0)
                        return iife153_((pat_let_tv63_).f16)
                    return iife152_(_pat_let44_0)
                rhs346_ = (self).f10
                rhs347_ = ((d_2214_v42_) + (d_2214_v42_)) + (_dafny.SeqWithoutIsStrInference([(iife147_(iife148_(d_2215_v43_))).cf57, ((d_2216_v44_)[d_2176_v13_] if (d_2176_v13_) in (d_2216_v44_) else (self).f10)]))
                rhs348_ = (d_2187_v18_).f16
                rhs349_ = len(((d_2218_v46_) + (_dafny.SeqWithoutIsStrInference([d_2217_v45_]))) + (d_2218_v46_))
                rhs350_ = False
                lhs239_ = (d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]
                lhs240_ = default__.safeIndex(590, ((d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]).length(0))
                lhs241_ = d_2213_v41_
                lhs242_ = default__.safeIndex(925, (d_2213_v41_).length(0))
                lhs243_ = (d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]
                lhs244_ = default__.safeIndex(220, ((d_2175_v12_)[default__.safeIndex(909, (d_2175_v12_).length(0))]).length(0))
                lhs245_ = globalState
                lhs239_[lhs240_] = rhs346_
                lhs241_[lhs242_] = rhs347_
                lhs243_[lhs244_] = rhs348_
                lhs245_.f1 = rhs349_
                r1 = rhs350_
        elif True:
            d_2224___mcc_h3_ = source38_.cf49
            d_2225_cf49_ = d_2224___mcc_h3_
            d_2226_v47_: _dafny.MultiSet
            d_2226_v47_ = _dafny.MultiSet([(self).f10])
            d_2227_v48_: D5
            d_2227_v48_ = D5_DC16(d_2226_v47_)
            source39_ = d_2227_v48_
            if source39_.is_DC17:
                d_2228___mcc_h4_ = source39_.cf28
                d_2229___mcc_h5_ = source39_.cf29
                d_2230_cf29_ = d_2229___mcc_h5_
                d_2231_cf28_ = d_2228___mcc_h4_
                index366_ = default__.safeIndex(170, (d_2176_v13_).length(0))
                (d_2176_v13_)[index366_] = d_2231_cf28_
                index367_ = default__.safeIndex(170, (d_2176_v13_).length(0))
                rhs351_ = (self).f10
                rhs352_ = d_2230_cf29_
                rhs353_ = (self).f13
                lhs246_ = d_2176_v13_
                lhs247_ = default__.safeIndex(170, (d_2176_v13_).length(0))
                lhs246_[lhs247_] = rhs351_
                d_2230_cf29_ = rhs352_
                r2 = rhs353_
                d_2232_v49_: _dafny.Array
                nw342_ = _dafny.Array(_dafny.Set({}), 1)
                d_2232_v49_ = nw342_
                d_2233_v50_: _dafny.Set
                d_2233_v50_ = _dafny.Set({(self).f13})
                index368_ = default__.safeIndex(314, (d_2232_v49_).length(0))
                (d_2232_v49_)[index368_] = d_2233_v50_
                index369_ = default__.safeIndex(314, (d_2232_v49_).length(0))
                (d_2232_v49_)[index369_] = (_dafny.Set({(self).f13, (self).f13})) | ((d_2233_v50_) - (d_2233_v50_))
                d_2234_v51_: _dafny.Array
                nw343_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_2234_v51_ = nw343_
                d_2235_v52_: _dafny.Array
                nw344_ = _dafny.Array(int(0), 26)
                d_2235_v52_ = nw344_
                index370_ = default__.safeIndex(834, (d_2234_v51_).length(0))
                (d_2234_v51_)[index370_] = d_2235_v52_
                index371_ = default__.safeIndex(834, (d_2234_v51_).length(0))
                (d_2234_v51_)[index371_] = d_2235_v52_
                r1 = not((d_2176_v13_)[default__.safeIndex(170, (d_2176_v13_).length(0))])
            elif source39_.is_DC16:
                d_2236___mcc_h6_ = source39_.cf27
                d_2237_cf27_ = d_2236___mcc_h6_
                d_2238_v53_: _dafny.Seq
                d_2238_v53_ = _dafny.SeqWithoutIsStrInference([-535])
                d_2239_v55_: _dafny.MultiSet
                def iife157_():
                    coll59_ = _dafny.Map()
                    compr_59_: int
                    for compr_59_ in (d_2238_v53_).Elements:
                        d_2240_v54_: int = compr_59_
                        if (d_2240_v54_) in (d_2238_v53_):
                            coll59_[default__.safeDivisionInt(d_2240_v54_, (self).f13)] = (self).f13
                    return _dafny.Map(coll59_)
                d_2239_v55_ = _dafny.MultiSet([((d_2237_cf27_)[(self).f10] if ((self).f10) in (d_2237_cf27_) else (self).f13), len(iife157_()
                )])
                d_2241_v56_: _dafny.Map
                d_2241_v56_ = _dafny.Map({False: (self).f10})
                d_2242_v57_: _dafny.Array
                nw345_ = _dafny.Array(None, 29)
                nw345_[int(0)] = (self).f13
                nw345_[int(1)] = (0) - (((self).f13) * ((self).f13))
                nw345_[int(2)] = (self).f13
                nw345_[int(3)] = len((self).f14)
                nw345_[int(4)] = (self).f13
                nw345_[int(5)] = (self).f13
                nw345_[int(6)] = ((0) - ((d_2238_v53_)[default__.safeIndex((self).f13, len(d_2238_v53_))])) + ((self).f13)
                nw345_[int(7)] = 658
                nw345_[int(8)] = default__.safeDivisionInt((self).f13, (self).f13)
                nw345_[int(9)] = (0) - (-834)
                nw345_[int(10)] = (self).f13
                nw345_[int(11)] = (self).f13
                nw345_[int(12)] = ((d_2239_v55_).intersection(d_2239_v55_)).cardinality
                nw345_[int(13)] = (self).f13
                nw345_[int(14)] = (self).f13
                nw345_[int(15)] = len((_dafny.Map({(self).f10: (self).f10})) | (d_2241_v56_))
                nw345_[int(16)] = (self).f13
                nw345_[int(17)] = ((self).f13) - (default__.fm1(globalState))
                nw345_[int(18)] = (self).f13
                nw345_[int(19)] = default__.safeModuloInt((self).f13, len((self).f14))
                nw345_[int(20)] = (self).f13
                nw345_[int(21)] = (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffvvncjhg")))) + ((self).f13))
                nw345_[int(22)] = (0) - (len(default__.fm3((0) - ((self).f13), (self).f13, globalState)))
                nw345_[int(23)] = default__.safeDivisionInt((self).f13, 626)
                nw345_[int(24)] = (self).f13
                nw345_[int(25)] = (self).f13
                nw345_[int(26)] = (self).f13
                nw345_[int(27)] = (self).f13
                nw345_[int(28)] = (self).f13
                d_2242_v57_ = nw345_
                index372_ = default__.safeIndex(865, (d_2242_v57_).length(0))
                (d_2242_v57_)[index372_] = (default__.fm41(globalState)).cardinality
                index373_ = default__.safeIndex(865, (d_2242_v57_).length(0))
                (d_2242_v57_)[index373_] = (self).f13
                d_2243_v58_: C6
                nw346_ = C6()
                nw346_.ctor__()
                d_2243_v58_ = nw346_
                r1 = (self).f10
                d_2244_v59_: D19
                d_2244_v59_ = D19_DC58((self).f13, (self).f10)
                d_2245_v60_: C9
                nw347_ = C9()
                nw347_.ctor__((self).f10, (d_2244_v59_).cf100)
                d_2245_v60_ = nw347_
                d_2246_v61_: _dafny.Seq
                d_2246_v61_ = _dafny.SeqWithoutIsStrInference([d_2245_v60_, d_2245_v60_])
                r1 = ((d_2246_v61_ if (self).f10 else (d_2246_v61_).set(default__.safeIndex((d_2242_v57_)[default__.safeIndex(865, (d_2242_v57_).length(0))], len(d_2246_v61_)), d_2245_v60_))) == (d_2246_v61_)
            elif True:
                d_2247___mcc_h7_ = source39_.cf30
                d_2248_cf30_ = d_2247___mcc_h7_
                d_2249_v62_: _dafny.Seq
                d_2249_v62_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])
                d_2250_v63_: _dafny.Map
                d_2250_v63_ = _dafny.Map({(d_2249_v62_) + (d_2249_v62_): (668) - ((0) - ((self).f13))})
                d_2250_v63_ = (d_2250_v63_).set(d_2249_v62_, (0) - ((self).f13))
                d_2251_v64_: str
                d_2251_v64_ = _dafny.CodePoint('k')
                d_2252_v65_: _dafny.Map
                d_2252_v65_ = _dafny.Map({(self).f10: (self).f10})
                d_2253_v66_: _dafny.Map
                d_2253_v66_ = _dafny.Map({d_2251_v64_: d_2252_v65_})
                d_2253_v66_ = (d_2253_v66_) | (d_2253_v66_)
                d_2254_v67_: _dafny.Seq
                d_2254_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmkyku"))
                d_2254_v67_ = _dafny.SeqWithoutIsStrInference([d_2251_v64_ for d_2255_i3_ in range(default__.abs(-662))])
                (globalState).f7 = ((self).f13) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpyfyh"))))
            d_2256_v68_: C7
            nw348_ = C7()
            nw348_.ctor__((self).f10, ((self).f10 if not((self).f10) else (self).f10))
            d_2256_v68_ = nw348_
            d_2257_v69_: _dafny.Array
            nw349_ = _dafny.Array(int(0), 25)
            d_2257_v69_ = nw349_
            index374_ = default__.safeIndex(589, (d_2257_v69_).length(0))
            (d_2257_v69_)[index374_] = len(d_2225_cf49_)
            d_2258_v70_: _dafny.Array
            nw350_ = _dafny.Array(_dafny.CodePoint('D'), 7)
            d_2258_v70_ = nw350_
            d_2259_v71_: str
            d_2259_v71_ = _dafny.CodePoint('t')
            index375_ = default__.safeIndex(261, (d_2258_v70_).length(0))
            (d_2258_v70_)[index375_] = d_2259_v71_
            index376_ = default__.safeIndex(589, (d_2257_v69_).length(0))
            index377_ = default__.safeIndex(261, (d_2258_v70_).length(0))
            rhs354_ = d_2226_v47_
            rhs355_ = (d_2225_cf49_).ispropersubset((d_2225_cf49_) | (d_2225_cf49_))
            rhs356_ = ((self).f13) - ((self).f13)
            rhs357_ = d_2259_v71_
            lhs248_ = d_2257_v69_
            lhs249_ = default__.safeIndex(589, (d_2257_v69_).length(0))
            lhs250_ = d_2258_v70_
            lhs251_ = default__.safeIndex(261, (d_2258_v70_).length(0))
            d_2226_v47_ = rhs354_
            r1 = rhs355_
            lhs248_[lhs249_] = rhs356_
            lhs250_[lhs251_] = rhs357_
            d_2260_v72_: _dafny.Map
            d_2260_v72_ = _dafny.Map({(self).f13: (self).f13})
            d_2260_v72_ = (d_2260_v72_).set((default__.fm18((self).f13, _dafny.Map({(d_2256_v68_).f19: (self).f13}), (d_2256_v68_).f19, globalState)).cardinality, (self).f13)
        d_2261_v73_: C9
        nw351_ = C9()
        nw351_.ctor__((self).f10, (self).f10)
        d_2261_v73_ = nw351_
        d_2262_v74_: _dafny.Set
        d_2262_v74_ = _dafny.Set({d_2261_v73_, d_2261_v73_})
        hi14_ = len(d_2262_v74_)
        for d_2263_i4_ in range((self).f13, hi14_):
            d_2264_v75_: _dafny.Map
            d_2264_v75_ = _dafny.Map({d_2263_i4_: ((self).f13) + ((self).f13)})
            d_2264_v75_ = (d_2264_v75_).set(d_2263_i4_, len((d_2264_v75_) | (d_2264_v75_)))
            d_2265_v76_: _dafny.MultiSet
            d_2265_v76_ = _dafny.MultiSet([not((self).f10), (d_2261_v73_).f17, (d_2261_v73_).f17])
            d_2266_v77_: D11
            d_2266_v77_ = D11_DC34((d_2265_v76_).cardinality, (self).f10, len((self).f14))
            d_2267_v78_: _dafny.Map
            d_2267_v78_ = _dafny.Map({(d_2261_v73_).f17: default__.safeModuloInt((d_2266_v77_).cf56, d_2263_i4_)})
            d_2268_v79_: _dafny.Seq
            d_2268_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gijrn"))
            d_2269_v80_: str
            d_2269_v80_ = _dafny.CodePoint('j')
            d_2270_v81_: _dafny.Map
            d_2270_v81_ = _dafny.Map({d_2269_v80_: default__.safeDivisionInt((self).f13, d_2263_i4_)})
            d_2271_v82_: _dafny.Seq
            d_2271_v82_ = _dafny.SeqWithoutIsStrInference([(self).f13])
            d_2272_v83_: _dafny.Seq
            d_2272_v83_ = _dafny.SeqWithoutIsStrInference([(d_2261_v73_).f17, False])
            rhs358_ = (((d_2267_v78_).set((self).f10, (self).f13)).set((d_2261_v73_).f17, d_2263_i4_)).set(not ((d_2261_v73_).f17) or ((d_2261_v73_).f17), (d_2271_v82_)[default__.safeIndex((self).f13, len(d_2271_v82_))])
            rhs359_ = ((self).f14) + ((self).f14)
            rhs360_ = ((d_2270_v81_).set(default__.fm4((d_2261_v73_).f17, (d_2261_v73_).f17, (d_2272_v83_)[default__.safeIndex(443, len(d_2272_v83_))], globalState), d_2263_i4_)).set(_dafny.CodePoint('i'), (0) - (d_2263_i4_))
            d_2267_v78_ = rhs358_
            d_2268_v79_ = rhs359_
            d_2270_v81_ = rhs360_
            d_2273_v84_: _dafny.Array
            nw352_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_2273_v84_ = nw352_
            d_2273_v84_ = d_2273_v84_
            d_2272_v83_ = (d_2272_v83_) + ((D18_DC56(d_2272_v83_)).cf97)
        (globalState).f0 = default__.safeDivisionInt((self).f13, (self).f13)
        d_2274_v85_: _dafny.Map
        d_2274_v85_ = _dafny.Map({not((d_2261_v73_).f17): d_2176_v13_})
        d_2274_v85_ = (d_2274_v85_) | (d_2274_v85_)
        d_2275_v86_: _dafny.Set
        d_2275_v86_ = _dafny.Set({(self).f14})
        r0 = (d_2275_v86_) - ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))})).intersection(d_2275_v86_))
        r1 = ((self).f13) >= (720)
        r2 = default__.safeModuloInt((self).f13, (0) - ((self).f13))
        return r0, r1, r2

    def m5(self, p0, globalState):
        d_2276_v0_: bool
        d_2276_v0_ = False
        d_2277_v1_: _dafny.Seq
        d_2277_v1_ = _dafny.SeqWithoutIsStrInference([len((self).f14)])
        d_2278_v2_: _dafny.Array
        nw353_ = _dafny.Array(None, 3)
        nw353_[int(0)] = d_2277_v1_
        nw353_[int(1)] = (d_2277_v1_) + (d_2277_v1_)
        nw353_[int(2)] = d_2277_v1_
        d_2278_v2_ = nw353_
        d_2279_v4_: _dafny.Seq
        def iife158_():
            coll60_ = _dafny.Map()
            compr_60_: int
            for compr_60_ in (_dafny.Map({(self).f13: (self).f13})).keys.Elements:
                d_2280_v3_: int = compr_60_
                if (d_2280_v3_) in (_dafny.Map({(self).f13: (self).f13})):
                    coll60_[(d_2280_v3_) * (len((self).f14))] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2281_i1_ in range(default__.abs(239))])
            return _dafny.Map(coll60_)
        d_2279_v4_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(iife158_()
) for d_2282_i0_ in range(default__.abs(517))]), _dafny.SeqWithoutIsStrInference([(self).f13])])
        d_2283_v5_: _dafny.Seq
        d_2283_v5_ = _dafny.SeqWithoutIsStrInference([(self).fm7((self).f10, globalState), d_2276_v0_])
        index378_ = default__.safeIndex(670, (d_2278_v2_).length(0))
        (d_2278_v2_)[index378_] = (d_2279_v4_)[default__.safeIndex(len(d_2283_v5_), len(d_2279_v4_))]
        d_2284_v6_: D16
        d_2284_v6_ = D16_DC49((D20_DC60(627)).cf102)
        pat_let_tv64_ = p0
        d_2285_v7_: _dafny.Array
        nw354_ = _dafny.Array(None, 27)
        nw354_[int(0)] = d_2284_v6_
        nw354_[int(1)] = d_2284_v6_
        def iife159_(_pat_let49_0):
            def iife160_(d_2286_dt__update__tmp_h0_):
                def iife161_(_pat_let50_0):
                    def iife162_(d_2287_dt__update_hcf85_h0_):
                        return D16_DC49(d_2287_dt__update_hcf85_h0_)
                    return iife162_(_pat_let50_0)
                return iife161_((0) - ((self).f13))
            return iife160_(_pat_let49_0)
        nw354_[int(2)] = iife159_(d_2284_v6_)
        def iife163_(_pat_let51_0):
            def iife164_(d_2288_dt__update__tmp_h1_):
                def iife165_(_pat_let52_0):
                    def iife166_(d_2289_dt__update_hcf85_h1_):
                        return D16_DC49(d_2289_dt__update_hcf85_h1_)
                    return iife166_(_pat_let52_0)
                return iife165_((self).f13)
            return iife164_(_pat_let51_0)
        nw354_[int(3)] = (iife163_(d_2284_v6_) if d_2276_v0_ else d_2284_v6_)
        nw354_[int(4)] = d_2284_v6_
        nw354_[int(5)] = d_2284_v6_
        nw354_[int(6)] = d_2284_v6_
        nw354_[int(7)] = d_2284_v6_
        nw354_[int(8)] = (d_2284_v6_ if not(default__.fm0(p0, globalState)) else d_2284_v6_)
        nw354_[int(9)] = d_2284_v6_
        nw354_[int(10)] = default__.fm56(globalState)
        nw354_[int(11)] = d_2284_v6_
        nw354_[int(12)] = D16_DC49((self).f13)
        nw354_[int(13)] = D16_DC49(711)
        nw354_[int(14)] = default__.fm56(globalState)
        nw354_[int(15)] = D16_DC49(p0)
        def iife167_(_pat_let53_0):
            def iife168_(d_2290_dt__update__tmp_h2_):
                def iife169_(_pat_let54_0):
                    def iife170_(d_2291_dt__update_hcf85_h2_):
                        return D16_DC49(d_2291_dt__update_hcf85_h2_)
                    return iife170_(_pat_let54_0)
                return iife169_((0) - (pat_let_tv64_))
            return iife168_(_pat_let53_0)
        nw354_[int(16)] = iife167_(d_2284_v6_)
        nw354_[int(17)] = d_2284_v6_
        nw354_[int(18)] = d_2284_v6_
        nw354_[int(19)] = d_2284_v6_
        nw354_[int(20)] = d_2284_v6_
        nw354_[int(21)] = d_2284_v6_
        nw354_[int(22)] = d_2284_v6_
        nw354_[int(23)] = d_2284_v6_
        nw354_[int(24)] = d_2284_v6_
        nw354_[int(25)] = D16_DC49(len(default__.fm13((self).f10, d_2276_v0_, globalState)))
        nw354_[int(26)] = d_2284_v6_
        d_2285_v7_ = nw354_
        index379_ = default__.safeIndex(667, (d_2285_v7_).length(0))
        (d_2285_v7_)[index379_] = d_2284_v6_
        d_2292_v8_: _dafny.Set
        d_2292_v8_ = _dafny.Set({(self).f13, p0})
        pat_let_tv65_ = p0
        index380_ = default__.safeIndex(670, (d_2278_v2_).length(0))
        index381_ = default__.safeIndex(667, (d_2285_v7_).length(0))
        def iife171_(_pat_let55_0):
            def iife172_(d_2293_dt__update__tmp_h3_):
                def iife173_(_pat_let56_0):
                    def iife174_(d_2294_dt__update_hcf85_h3_):
                        return D16_DC49(d_2294_dt__update_hcf85_h3_)
                    return iife174_(_pat_let56_0)
                return iife173_(pat_let_tv65_)
            return iife172_(_pat_let55_0)
        rhs361_ = (d_2292_v8_) != ((d_2292_v8_) | (d_2292_v8_))
        rhs362_ = (d_2277_v1_) + (d_2277_v1_)
        rhs363_ = iife171_(d_2284_v6_)
        lhs252_ = d_2278_v2_
        lhs253_ = default__.safeIndex(670, (d_2278_v2_).length(0))
        lhs254_ = d_2285_v7_
        lhs255_ = default__.safeIndex(667, (d_2285_v7_).length(0))
        d_2276_v0_ = rhs361_
        lhs252_[lhs253_] = rhs362_
        lhs254_[lhs255_] = rhs363_
        d_2295_v9_: _dafny.Map
        d_2295_v9_ = _dafny.Map({default__.fm4((self).f10, d_2276_v0_, (self).f10, globalState): default__.safeDivisionInt((d_2277_v1_)[default__.safeIndex(default__.fm1(globalState), len(d_2277_v1_))], 916)})
        d_2296_v10_: str
        d_2296_v10_ = _dafny.CodePoint('x')
        d_2295_v9_ = (d_2295_v9_).set(d_2296_v10_, p0)
        d_2295_v9_ = (d_2295_v9_).set(d_2296_v10_, len((self).f14))
        d_2297_v11_: _dafny.Map
        d_2297_v11_ = _dafny.Map({d_2296_v10_: d_2283_v5_})
        d_2297_v11_ = (d_2297_v11_).set(d_2296_v10_, d_2283_v5_)
        hi15_ = default__.fm1(globalState)
        for d_2298_i2_ in range(p0, hi15_):
            if (self).f10:
                d_2299_v12_: _dafny.Seq
                d_2299_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqev"))])
                d_2299_v12_ = d_2299_v12_
                d_2300_v13_: _dafny.Map
                d_2300_v13_ = _dafny.Map({d_2276_v0_: d_2296_v10_})
                d_2300_v13_ = (d_2300_v13_).set(d_2276_v0_, d_2296_v10_)
                d_2301_v14_: _dafny.Map
                d_2301_v14_ = _dafny.Map({d_2276_v0_: p0})
                d_2302_v15_: _dafny.Map
                d_2302_v15_ = _dafny.Map({_dafny.MultiSet([p0]): d_2301_v14_})
                d_2302_v15_ = d_2302_v15_
                d_2303_v16_: _dafny.Set
                d_2303_v16_ = _dafny.Set({not((d_2283_v5_)[default__.safeIndex(-719, len(d_2283_v5_))]), d_2276_v0_, d_2276_v0_})
                d_2304_v17_: D10
                d_2304_v17_ = D10_DC31(default__.fm0((self).f13, globalState), (d_2303_v16_) == (d_2303_v16_), ((self).f13) + (d_2298_i2_))
                pat_let_tv66_ = p0
                def iife175_(_pat_let57_0):
                    def iife176_(d_2305_dt__update__tmp_h4_):
                        def iife177_(_pat_let58_0):
                            def iife178_(d_2306_dt__update_hcf52_h0_):
                                return D10_DC31((d_2305_dt__update__tmp_h4_).cf50, (d_2305_dt__update__tmp_h4_).cf51, d_2306_dt__update_hcf52_h0_)
                            return iife178_(_pat_let58_0)
                        return iife177_(pat_let_tv66_)
                    return iife176_(_pat_let57_0)
                d_2304_v17_ = iife175_(d_2304_v17_)
                d_2307_v18_: C0
                nw355_ = C0()
                nw355_.ctor__()
                d_2307_v18_ = nw355_
            elif True:
                (globalState).f1 = default__.fm1(globalState)
                d_2308_v19_: _dafny.Map
                d_2308_v19_ = _dafny.Map({d_2298_i2_: d_2298_i2_})
                d_2309_v20_: T1
                nw356_ = C4()
                nw356_.ctor__()
                d_2309_v20_ = nw356_
                d_2310_v21_: _dafny.Map
                d_2310_v21_ = _dafny.Map({(self).f10: d_2309_v20_})
                d_2311_v22_: _dafny.Map
                d_2311_v22_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(((self).f14).set(default__.safeIndex(((d_2308_v19_)[(self).f13] if ((self).f13) in (d_2308_v19_) else len(default__.fm26(globalState))), len((self).f14)), d_2296_v10_)), d_2298_i2_])): len((d_2310_v21_) | (d_2310_v21_))})
                d_2312_v23_: _dafny.MultiSet
                d_2312_v23_ = _dafny.MultiSet([d_2296_v10_])
                d_2313_v24_: _dafny.MultiSet
                d_2313_v24_ = _dafny.MultiSet([(self).f13, (self).f13, len((_dafny.SeqWithoutIsStrInference([(d_2309_v20_).fm8((self).f13, (self).f10, (self).f10, globalState), d_2276_v0_, (self).f10, (self).f10, (self).f10])).set(default__.safeIndex(((d_2295_v9_)[d_2296_v10_] if (d_2296_v10_) in (d_2295_v9_) else ((d_2312_v23_)[d_2296_v10_] if (d_2296_v10_) in (d_2312_v23_) else 804)), len(_dafny.SeqWithoutIsStrInference([(d_2309_v20_).fm8((self).f13, (self).f10, (self).f10, globalState), d_2276_v0_, (self).f10, (self).f10, (self).f10]))), d_2276_v0_)), p0, d_2298_i2_])
                d_2308_v19_ = (d_2308_v19_).set((p0) - ((0) - ((d_2313_v24_).cardinality)), d_2298_i2_)
                d_2314_v25_: _dafny.Array
                nw357_ = _dafny.Array(int(0), 9)
                d_2314_v25_ = nw357_
                d_2314_v25_ = d_2314_v25_
                d_2315_v26_: _dafny.Seq
                d_2315_v26_ = _dafny.SeqWithoutIsStrInference([d_2309_v20_, d_2309_v20_, d_2309_v20_, d_2309_v20_, d_2309_v20_])
                d_2316_v27_: _dafny.Array
                nw358_ = _dafny.Array(None, 1)
                nw358_[int(0)] = (d_2315_v26_)[default__.safeIndex(384, len(d_2315_v26_))]
                d_2316_v27_ = nw358_
                d_2316_v27_ = d_2316_v27_
                d_2317_v28_: _dafny.Array
                nw359_ = _dafny.Array(None, 6)
                nw359_[int(0)] = d_2276_v0_
                nw359_[int(1)] = d_2276_v0_
                nw359_[int(2)] = d_2276_v0_
                nw359_[int(3)] = not((self).f10)
                nw359_[int(4)] = (self).f10
                nw359_[int(5)] = d_2276_v0_
                d_2317_v28_ = nw359_
                d_2318_v29_: _dafny.Seq
                d_2318_v29_ = _dafny.SeqWithoutIsStrInference([d_2317_v28_])
                d_2319_v30_: D0
                d_2319_v30_ = D0_DC0((d_2318_v29_)[default__.safeIndex(d_2298_i2_, len(d_2318_v29_))])
                d_2319_v30_ = d_2319_v30_
            d_2320_v31_: _dafny.Map
            d_2320_v31_ = _dafny.Map({d_2276_v0_: d_2296_v10_})
            d_2321_v32_: _dafny.Map
            d_2321_v32_ = _dafny.Map({d_2296_v10_: True})
            d_2322_v33_: _dafny.Map
            d_2322_v33_ = _dafny.Map({d_2320_v31_: ((d_2321_v32_)[d_2296_v10_] if (d_2296_v10_) in (d_2321_v32_) else d_2276_v0_)})
            d_2323_v34_: _dafny.Array
            nw360_ = _dafny.Array(None, 5)
            nw360_[int(0)] = not (d_2276_v0_) or (False)
            nw360_[int(1)] = not(d_2276_v0_)
            nw360_[int(2)] = d_2276_v0_
            nw360_[int(3)] = (self).f10
            nw360_[int(4)] = not((len((d_2322_v33_).set((d_2320_v31_).set((self).f10, d_2296_v10_), True))) < (len(d_2277_v1_)))
            d_2323_v34_ = nw360_
            index382_ = default__.safeIndex(618, (d_2323_v34_).length(0))
            (d_2323_v34_)[index382_] = d_2276_v0_
            index383_ = default__.safeIndex(618, (d_2323_v34_).length(0))
            (d_2323_v34_)[index383_] = d_2276_v0_
            d_2276_v0_ = d_2276_v0_
            d_2324_v35_: _dafny.Seq
            d_2325_v36_: _dafny.Seq
            out64_: _dafny.Seq
            out65_: _dafny.Seq
            out64_, out65_ = default__.m0((self).f13, d_2323_v34_, globalState)
            d_2324_v35_ = out64_
            d_2325_v36_ = out65_
        d_2326_v37_: _dafny.Map
        d_2326_v37_ = _dafny.Map({d_2276_v0_: d_2276_v0_})
        d_2327_v38_: _dafny.Set
        d_2327_v38_ = _dafny.Set({d_2326_v37_})
        d_2328_v39_: _dafny.Map
        d_2328_v39_ = _dafny.Map({d_2296_v10_: (self).f10})
        d_2329_v40_: _dafny.Map
        d_2329_v40_ = _dafny.Map({d_2328_v39_: (self).f10})
        d_2330_v41_: _dafny.Map
        d_2330_v41_ = _dafny.Map({d_2276_v0_: (self).f13})
        d_2331_v42_: _dafny.Map
        d_2331_v42_ = _dafny.Map({(self).fm9((self).f14, d_2327_v38_, globalState): (default__.fm18((self).f13, d_2330_v41_, True, globalState)).ispropersubset(_dafny.MultiSet([(self).f13, len((d_2329_v40_).set(d_2328_v39_, False)), (self).f13]))})
        d_2332_v43_: _dafny.Seq
        d_2332_v43_ = _dafny.SeqWithoutIsStrInference([(self).f14])
        d_2333_i3_: int
        d_2333_i3_ = 0
        with _dafny.label("13"):
            while ((d_2331_v42_)[p0] if (p0) in (d_2331_v42_) else (d_2296_v10_) in ((d_2332_v43_)[default__.safeIndex(p0, len(d_2332_v43_))])):
                with _dafny.c_label("13"):
                    if (d_2333_i3_) >= (100):
                        raise _dafny.Break("13")
                    d_2333_i3_ = (d_2333_i3_) + (1)
                    if default__.fm0((self).f13, globalState):
                        d_2326_v37_ = (d_2326_v37_) | ((d_2326_v37_) | (d_2326_v37_))
                        d_2334_v44_: _dafny.Array
                        nw361_ = _dafny.Array(None, 8)
                        d_2334_v44_ = nw361_
                        d_2334_v44_ = d_2334_v44_
                        d_2276_v0_ = (self).fm6((self).fm9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywsbusnf")), d_2327_v38_, globalState), len(_dafny.Map({(self).f10: len((_dafny.SeqWithoutIsStrInference([d_2296_v10_ for d_2335_i4_ in range(default__.abs(575))])).set(default__.safeIndex((self).f13, len(_dafny.SeqWithoutIsStrInference([d_2296_v10_ for d_2336_i4_ in range(default__.abs(575))]))), d_2296_v10_))})), (p0) - (default__.fm1(globalState)), (self).f10, globalState)
                        d_2337_v45_: T1
                        nw362_ = C4()
                        nw362_.ctor__()
                        d_2337_v45_ = nw362_
                        d_2338_v46_: _dafny.Array
                        nw363_ = _dafny.Array(False, 18)
                        d_2338_v46_ = nw363_
                        index384_ = default__.safeIndex(929, (d_2338_v46_).length(0))
                        (d_2338_v46_)[index384_] = ((self).f10 if d_2276_v0_ else True)
                        index385_ = default__.safeIndex(556, (d_2338_v46_).length(0))
                        (d_2338_v46_)[index385_] = d_2276_v0_
                        index386_ = default__.safeIndex(929, (d_2338_v46_).length(0))
                        index387_ = default__.safeIndex(556, (d_2338_v46_).length(0))
                        rhs364_ = d_2337_v45_
                        rhs365_ = (d_2337_v45_).fm9(default__.fm38(d_2276_v0_, p0, globalState), d_2327_v38_, globalState)
                        rhs366_ = d_2276_v0_
                        rhs367_ = (d_2283_v5_)[default__.safeIndex(p0, len(d_2283_v5_))]
                        lhs256_ = globalState
                        lhs257_ = d_2338_v46_
                        lhs258_ = default__.safeIndex(929, (d_2338_v46_).length(0))
                        lhs259_ = d_2338_v46_
                        lhs260_ = default__.safeIndex(556, (d_2338_v46_).length(0))
                        d_2337_v45_ = rhs364_
                        lhs256_.f7 = rhs365_
                        lhs257_[lhs258_] = rhs366_
                        lhs259_[lhs260_] = rhs367_
                        index388_ = default__.safeIndex(929, (d_2338_v46_).length(0))
                        (d_2338_v46_)[index388_] = (d_2283_v5_)[default__.safeIndex((self).f13, len(d_2283_v5_))]
                    elif True:
                        (globalState).f7 = ((default__.fm1(globalState)) * (207) if (self).f10 else (self).f13)
                        d_2339_v47_: _dafny.Seq
                        d_2339_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbiusl"))
                        d_2339_v47_ = _dafny.SeqWithoutIsStrInference([(d_2296_v10_ if (self).f10 else d_2296_v10_) for d_2340_i5_ in range(default__.abs(733))])
                        d_2276_v0_ = (self).f10
                        d_2341_v48_: C8
                        nw364_ = C8()
                        nw364_.ctor__(d_2296_v10_)
                        d_2341_v48_ = nw364_
                        d_2342_v49_: _dafny.Array
                        nw365_ = _dafny.Array(False, 6)
                        d_2342_v49_ = nw365_
                        index389_ = default__.safeIndex(71, (d_2342_v49_).length(0))
                        (d_2342_v49_)[index389_] = d_2276_v0_
                        index390_ = default__.safeIndex(71, (d_2342_v49_).length(0))
                        (d_2342_v49_)[index390_] = ((self).f10 if (self).f10 else (d_2283_v5_)[default__.safeIndex((self).f13, len(d_2283_v5_))])
                    d_2276_v0_ = (((self).f13) - (0)) >= ((self).f13)
                    (globalState).f1 = (0) - ((len((self).f14)) - ((self).f13))
                    d_2343_v50_: _dafny.Array
                    def lambda109_(d_2344_i6_):
                        return (self).f10

                    init58_ = lambda109_
                    nw366_ = _dafny.Array(None, 29)
                    for i0_58_ in range(nw366_.length(0)):
                        nw366_[i0_58_] = init58_(i0_58_)
                    d_2343_v50_ = nw366_
                    index391_ = default__.safeIndex(565, (d_2343_v50_).length(0))
                    (d_2343_v50_)[index391_] = (self).f10
                    index392_ = default__.safeIndex(565, (d_2343_v50_).length(0))
                    (d_2343_v50_)[index392_] = True
                    pass
            pass

    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14

class C12(T0, T1):
    def  __init__(self):
        self._f10: bool = False
        self.f12: bool = False
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C12"
    @property
    def f10(self):
        return self._f10
    def ctor__(self, f11, f12, f10):
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f10 = f10

    def fm6(self, p0, p1, p2, p3, globalState):
        return (self).f11

    def fm7(self, p0, globalState):
        return (-237) != ((-48) * ((_dafny.MultiSet([D2_DC6(_dafny.CodePoint('r'))])).cardinality))

    def fm8(self, p0, p1, p2, globalState):
        return not(((_dafny.Set({self.f12})) - (_dafny.Set({self.f12}))).ispropersubset((_dafny.Set({(self).f11})) | (_dafny.Set({(self).f11, (self).f11}))))

    def fm9(self, p0, p1, globalState):
        def iife179_():
            coll61_ = _dafny.Set()
            compr_61_: int
            for compr_61_ in _dafny.IntegerRange(930, 542):
                d_2345_v0_: int = compr_61_
                if ((930) <= (d_2345_v0_)) and ((d_2345_v0_) < (542)):
                    coll61_ = coll61_.union(_dafny.Set([(d_2345_v0_) + (-560)]))
            return _dafny.Set(coll61_)
        return (0) - ((((0) - (len(iife179_()
        ))) + (498)) + (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([405])), -418)))

    def fm10(self, globalState):
        return not((default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhyfgw"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhdnniyk"))))) >= ((len(_dafny.Set({self.f12}))) - (len(_dafny.SeqWithoutIsStrInference([210 for d_2346_i0_ in range(default__.abs(41))])))))

    def m3(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Map = _dafny.Map({})
        if not((self).f11):
            d_2347_v0_: str
            d_2347_v0_ = _dafny.CodePoint('r')
            d_2348_v1_: D2
            d_2348_v1_ = D2_DC6(d_2347_v0_)
            d_2349_v2_: _dafny.Seq
            d_2349_v2_ = _dafny.SeqWithoutIsStrInference([d_2347_v0_, (d_2348_v1_).cf8])
            d_2350_v3_: _dafny.Map
            d_2350_v3_ = _dafny.Map({self.f12: not(p1)})
            d_2351_v4_: int
            d_2351_v4_ = 926
            d_2352_v5_: _dafny.Set
            d_2352_v5_ = _dafny.Set({d_2350_v3_, default__.fm11((self).f11, d_2351_v4_, d_2351_v4_, globalState), d_2350_v3_})
            d_2353_v6_: _dafny.Set
            d_2353_v6_ = _dafny.Set({(d_2349_v2_)[default__.safeIndex((0) - ((self).fm9(_dafny.SeqWithoutIsStrInference([d_2347_v0_ for d_2354_i0_ in range(default__.abs(234))]), d_2352_v5_, globalState)), len(d_2349_v2_))]})
            d_2353_v6_ = (d_2353_v6_ if not(p1) else d_2353_v6_)
            if p1:
                d_2355_v7_: _dafny.Map
                d_2355_v7_ = _dafny.Map({d_2351_v4_: (self).f11})
                d_2355_v7_ = (d_2355_v7_).set(d_2351_v4_, True)
                d_2356_v8_: _dafny.MultiSet
                d_2356_v8_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2347_v0_ for d_2357_i1_ in range(default__.abs(934))]), d_2349_v2_])
                d_2358_v9_: bool
                d_2359_v10_: int
                out66_: bool
                out67_: int
                out66_, out67_ = (self).m6((d_2351_v4_) + (d_2351_v4_), self.f12, (d_2356_v8_) | (d_2356_v8_), globalState)
                d_2358_v9_ = out66_
                d_2359_v10_ = out67_
                r0 = self.f12
                d_2360_v11_: _dafny.MultiSet
                d_2360_v11_ = _dafny.MultiSet([True, not(p1), p1])
                d_2361_v12_: _dafny.Map
                d_2361_v12_ = _dafny.Map({(self).f11: d_2349_v2_})
                (self).f12 = not((((d_2360_v11_)[(self).f11] if ((self).f11) in (d_2360_v11_) else 239)) > (len(((d_2361_v12_)[d_2358_v9_] if (d_2358_v9_) in (d_2361_v12_) else d_2349_v2_))))
                d_2358_v9_ = d_2358_v9_
            elif True:
                d_2362_v13_: _dafny.Array
                nw367_ = _dafny.Array(False, 14)
                d_2362_v13_ = nw367_
                d_2363_v14_: D0
                d_2363_v14_ = D0_DC2((self).f11, d_2362_v13_)
                d_2364_v15_: _dafny.Map
                d_2364_v15_ = _dafny.Map({(self).f11: d_2350_v3_})
                d_2365_v16_: _dafny.Seq
                d_2365_v16_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({(self).f11: (d_2363_v14_).cf2}) if p1 else default__.fm11(self.f12, d_2351_v4_, d_2351_v4_, globalState)), d_2350_v3_, d_2350_v3_, d_2350_v3_, ((d_2364_v15_)[(self).f11] if ((self).f11) in (d_2364_v15_) else d_2350_v3_)])
                d_2366_v17_: _dafny.Seq
                d_2366_v17_ = _dafny.SeqWithoutIsStrInference([self.f12])
                rhs368_ = ((d_2365_v16_) + (d_2365_v16_)) + (_dafny.SeqWithoutIsStrInference([(D3_DC11(d_2351_v4_, d_2350_v3_)).cf18 for d_2367_i2_ in range(default__.abs(705))]))
                rhs369_ = d_2349_v2_
                rhs370_ = (d_2366_v17_)[default__.safeIndex(d_2351_v4_, len(d_2366_v17_))]
                lhs261_ = self
                d_2365_v16_ = rhs368_
                r1 = rhs369_
                lhs261_.f12 = rhs370_
                d_2368_v18_: _dafny.Array
                def lambda110_(d_2369_v4_):
                    def lambda111_(d_2370_i3_):
                        return (d_2370_i3_) * (d_2369_v4_)

                    return lambda111_

                init59_ = lambda110_(d_2351_v4_)
                nw368_ = _dafny.Array(None, 5)
                for i0_59_ in range(nw368_.length(0)):
                    nw368_[i0_59_] = init59_(i0_59_)
                d_2368_v18_ = nw368_
                d_2371_v19_: _dafny.Map
                d_2371_v19_ = _dafny.Map({_dafny.MultiSet([d_2368_v18_]): d_2351_v4_})
                d_2372_v20_: _dafny.MultiSet
                d_2372_v20_ = _dafny.MultiSet([((d_2371_v19_)[_dafny.MultiSet([d_2368_v18_])] if (_dafny.MultiSet([d_2368_v18_])) in (d_2371_v19_) else d_2351_v4_), d_2351_v4_, len(default__.fm12((self).f11, (self).f10, len(d_2349_v2_), d_2351_v4_, globalState))])
                d_2373_v21_: D4
                d_2373_v21_ = D4_DC15(d_2351_v4_, d_2372_v20_, d_2347_v0_, d_2351_v4_)
                d_2374_v22_: _dafny.Map
                d_2374_v22_ = _dafny.Map({d_2347_v0_: (d_2373_v21_ if False else d_2373_v21_)})
                pat_let_tv67_ = d_2347_v0_
                pat_let_tv68_ = d_2351_v4_
                def iife180_(_pat_let59_0):
                    def iife181_(d_2375_dt__update__tmp_h0_):
                        def iife182_(_pat_let60_0):
                            def iife183_(d_2376_dt__update_hcf25_h0_):
                                def iife184_(_pat_let61_0):
                                    def iife185_(d_2377_dt__update_hcf24_h0_):
                                        return D4_DC15((d_2375_dt__update__tmp_h0_).cf23, d_2377_dt__update_hcf24_h0_, d_2376_dt__update_hcf25_h0_, (d_2375_dt__update__tmp_h0_).cf26)
                                    return iife185_(_pat_let61_0)
                                return iife184_(_dafny.MultiSet([pat_let_tv68_]))
                            return iife183_(_pat_let60_0)
                        return iife182_(pat_let_tv67_)
                    return iife181_(_pat_let59_0)
                d_2374_v22_ = (d_2374_v22_).set(d_2347_v0_, iife180_(d_2373_v21_))
                index393_ = default__.safeIndex(5, (d_2362_v13_).length(0))
                (d_2362_v13_)[index393_] = ((d_2350_v3_)[(self).f10] if ((self).f10) in (d_2350_v3_) else not(True))
                d_2378_v23_: _dafny.Seq
                d_2378_v23_ = _dafny.SeqWithoutIsStrInference([8, d_2351_v4_])
                d_2379_v24_: _dafny.Map
                d_2379_v24_ = _dafny.Map({(self).f11: D4_DC14((0) - (d_2351_v4_), d_2351_v4_)})
                d_2380_v25_: _dafny.MultiSet
                d_2380_v25_ = _dafny.MultiSet([d_2379_v24_, d_2379_v24_])
                d_2381_v26_: _dafny.Map
                d_2381_v26_ = _dafny.Map({d_2362_v13_: p1})
                d_2382_v27_: _dafny.Set
                d_2382_v27_ = _dafny.Set({d_2351_v4_, 11, d_2351_v4_, (d_2372_v20_).cardinality})
                d_2383_v28_: _dafny.Set
                d_2383_v28_ = _dafny.Set({d_2351_v4_, len(d_2382_v27_)})
                d_2384_v29_: _dafny.Map
                d_2384_v29_ = _dafny.Map({len(d_2382_v27_): d_2381_v26_})
                index394_ = default__.safeIndex(5, (d_2362_v13_).length(0))
                rhs371_ = (len(d_2378_v23_)) + (d_2351_v4_)
                rhs372_ = p1
                rhs373_ = (d_2380_v25_).issubset(_dafny.MultiSet([d_2379_v24_]))
                rhs374_ = len((_dafny.Map({d_2362_v13_: (d_2366_v17_)[default__.safeIndex(d_2351_v4_, len(d_2366_v17_))]})) | ((d_2381_v26_) | (((d_2384_v29_)[len(d_2366_v17_)] if (len(d_2366_v17_)) in (d_2384_v29_) else _dafny.Map({d_2362_v13_: self.f12})))))
                lhs262_ = globalState
                lhs263_ = d_2362_v13_
                lhs264_ = default__.safeIndex(5, (d_2362_v13_).length(0))
                lhs265_ = globalState
                lhs262_.f1 = rhs371_
                r0 = rhs372_
                lhs263_[lhs264_] = rhs373_
                lhs265_.f0 = rhs374_
                d_2385_v30_: _dafny.Seq
                d_2386_v31_: _dafny.Seq
                out68_: _dafny.Seq
                out69_: _dafny.Seq
                out68_, out69_ = default__.m0(d_2351_v4_, d_2362_v13_, globalState)
                d_2385_v30_ = out68_
                d_2386_v31_ = out69_
                (globalState).f0 = (d_2351_v4_) - (d_2351_v4_)
            d_2387_v32_: _dafny.Seq
            d_2387_v32_ = _dafny.SeqWithoutIsStrInference([not(p1)])
            d_2388_v33_: _dafny.MultiSet
            d_2388_v33_ = _dafny.MultiSet([d_2349_v2_])
            d_2389_v35_: _dafny.Map
            d_2389_v35_ = _dafny.Map({d_2351_v4_: self.f12})
            d_2390_v36_: _dafny.Map
            d_2390_v36_ = _dafny.Map({len(d_2389_v35_): p1})
            def iife186_():
                coll62_ = _dafny.Map()
                compr_62_: int
                for compr_62_ in (d_2390_v36_).keys.Elements:
                    d_2391_v34_: int = compr_62_
                    if (d_2391_v34_) in (d_2390_v36_):
                        coll62_[default__.safeModuloInt(d_2391_v34_, d_2351_v4_)] = d_2350_v3_
                return _dafny.Map(coll62_)
            rhs375_ = ((d_2387_v32_) + ((_dafny.SeqWithoutIsStrInference([True])).set(default__.safeIndex(d_2351_v4_, len(_dafny.SeqWithoutIsStrInference([True]))), (self).fm8(len(iife186_()
            ), (self).f11, False, globalState)))) + (d_2387_v32_)
            rhs376_ = ((d_2349_v2_) + (d_2349_v2_)) + (d_2349_v2_)
            rhs377_ = ((_dafny.SeqWithoutIsStrInference([d_2347_v0_ for d_2392_i4_ in range(default__.abs(42))])) + (d_2349_v2_)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))
            rhs378_ = d_2388_v33_
            d_2387_v32_ = rhs375_
            r1 = rhs376_
            r0 = rhs377_
            d_2388_v33_ = rhs378_
            if not (False) or ((self).f10):
                d_2393_v37_: _dafny.Array
                nw369_ = _dafny.Array(int(0), 18)
                d_2393_v37_ = nw369_
                d_2394_v38_: _dafny.Seq
                d_2394_v38_ = _dafny.SeqWithoutIsStrInference([len(d_2387_v32_), d_2351_v4_])
                index395_ = default__.safeIndex(959, (d_2393_v37_).length(0))
                (d_2393_v37_)[index395_] = len(d_2394_v38_)
                index396_ = default__.safeIndex(959, (d_2393_v37_).length(0))
                (d_2393_v37_)[index396_] = (d_2351_v4_) * (d_2351_v4_)
                d_2395_v39_: _dafny.Array
                nw370_ = _dafny.Array(False, 9)
                d_2395_v39_ = nw370_
                index397_ = default__.safeIndex(409, (d_2395_v39_).length(0))
                (d_2395_v39_)[index397_] = self.f12
                d_2396_v40_: D4
                d_2396_v40_ = D4_DC14((d_2393_v37_)[default__.safeIndex(959, (d_2393_v37_).length(0))], d_2351_v4_)
                index398_ = default__.safeIndex(409, (d_2395_v39_).length(0))
                rhs379_ = (self).f10
                rhs380_ = d_2351_v4_
                rhs381_ = (d_2396_v40_).cf21
                lhs266_ = d_2395_v39_
                lhs267_ = default__.safeIndex(409, (d_2395_v39_).length(0))
                lhs268_ = globalState
                lhs269_ = globalState
                lhs266_[lhs267_] = rhs379_
                lhs268_.f0 = rhs380_
                lhs269_.f7 = rhs381_
                d_2397_v41_: _dafny.Map
                d_2397_v41_ = _dafny.Map({d_2351_v4_: len(d_2387_v32_)})
                (globalState).f0 = default__.safeModuloInt(579, ((d_2397_v41_)[(d_2393_v37_)[default__.safeIndex(959, (d_2393_v37_).length(0))]] if ((d_2393_v37_)[default__.safeIndex(959, (d_2393_v37_).length(0))]) in (d_2397_v41_) else (0) - (len(_dafny.Map({(self).f11: (self).f11})))))
                d_2397_v41_ = (d_2397_v41_).set(d_2351_v4_, ((d_2393_v37_)[default__.safeIndex(959, (d_2393_v37_).length(0))]) - (d_2351_v4_))
                (self).f12 = (d_2387_v32_)[default__.safeIndex(d_2351_v4_, len(d_2387_v32_))]
            elif True:
                d_2398_v42_: _dafny.Map
                d_2398_v42_ = _dafny.Map({(self).f10: d_2351_v4_})
                d_2399_v43_: C3
                nw371_ = C3()
                nw371_.ctor__(p1, ((d_2387_v32_)[default__.safeIndex(default__.fm1(globalState), len(d_2387_v32_))]) not in (d_2398_v42_))
                d_2399_v43_ = nw371_
                d_2400_v44_: _dafny.Array
                nw372_ = _dafny.Array(int(0), 24)
                d_2400_v44_ = nw372_
                index399_ = default__.safeIndex(133, (d_2400_v44_).length(0))
                (d_2400_v44_)[index399_] = len((d_2390_v36_) | (_dafny.Map({d_2351_v4_: True})))
                d_2401_v45_: _dafny.Seq
                d_2401_v45_ = _dafny.SeqWithoutIsStrInference([len(d_2387_v32_), d_2351_v4_, 963])
                index400_ = default__.safeIndex(133, (d_2400_v44_).length(0))
                (d_2400_v44_)[index400_] = (d_2401_v45_)[default__.safeIndex(len(d_2349_v2_), len(d_2401_v45_))]
                d_2402_v46_: _dafny.Map
                d_2402_v46_ = _dafny.Map({_dafny.Set({p1}): (self).f10})
                (globalState).f7 = ((0) - ((d_2400_v44_)[default__.safeIndex(133, (d_2400_v44_).length(0))])) * (len(((d_2387_v32_) + (d_2387_v32_)).set(default__.safeIndex(len((d_2402_v46_).set(_dafny.Set({False}), self.f12)), len((d_2387_v32_) + (d_2387_v32_))), (self).f11)))
                r0 = self.f12
                d_2403_v47_: C9
                nw373_ = C9()
                nw373_.ctor__((d_2399_v43_).f20, p1)
                d_2403_v47_ = nw373_
            r0 = default__.fm0(d_2351_v4_, globalState)
        elif True:
            d_2404_v48_: _dafny.Seq
            d_2404_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnhqf"))
            d_2405_v49_: int
            d_2405_v49_ = 633
            d_2406_v50_: str
            d_2406_v50_ = _dafny.CodePoint('u')
            d_2407_v51_: _dafny.Map
            d_2407_v51_ = _dafny.Map({False: (d_2404_v48_) <= ((d_2404_v48_).set(default__.safeIndex(d_2405_v49_, len(d_2404_v48_)), d_2406_v50_))})
            d_2407_v51_ = ((d_2407_v51_).set(self.f12, p1)) | (d_2407_v51_)
            r0 = (d_2406_v50_) in (d_2404_v48_)
            d_2408_v52_: _dafny.Seq
            d_2408_v52_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2409_v53_: D3
            d_2409_v53_ = D3_DC11(len(((d_2408_v52_) + (d_2408_v52_)).set(default__.safeIndex(d_2405_v49_, len((d_2408_v52_) + (d_2408_v52_))), True)), d_2407_v51_)
            source40_ = d_2409_v53_
            if source40_.is_DC10:
                d_2410___mcc_h0_ = source40_.cf14
                d_2411___mcc_h1_ = source40_.cf15
                d_2412___mcc_h2_ = source40_.cf16
                d_2413_cf16_ = d_2412___mcc_h2_
                d_2414_cf15_ = d_2411___mcc_h1_
                d_2415_cf14_ = d_2410___mcc_h0_
                d_2416_v54_: _dafny.Set
                d_2416_v54_ = _dafny.Set({(self).fm8(d_2405_v49_, self.f12, self.f12, globalState)})
                d_2417_v55_: _dafny.MultiSet
                d_2417_v55_ = _dafny.MultiSet([(0) - (d_2405_v49_)])
                d_2418_v57_: D13
                def iife187_():
                    coll63_ = _dafny.Map()
                    compr_63_: int
                    for compr_63_ in _dafny.IntegerRange(-537, 695):
                        d_2419_v56_: int = compr_63_
                        if ((-537) <= (d_2419_v56_)) and ((d_2419_v56_) < (695)):
                            coll63_[(d_2419_v56_) * (d_2405_v49_)] = D3_DC10(d_2415_cf14_, d_2414_cf15_, True)
                    return _dafny.Map(coll63_)
                d_2418_v57_ = D13_DC39(len(iife187_()
), d_2414_cf15_)
                d_2420_v58_: _dafny.MultiSet
                d_2420_v58_ = _dafny.MultiSet([self.f12])
                d_2421_v59_: _dafny.Set
                d_2421_v59_ = _dafny.Set({d_2407_v51_, d_2407_v51_})
                d_2422_v60_: _dafny.Map
                d_2422_v60_ = _dafny.Map({d_2405_v49_: (self).fm9(_dafny.SeqWithoutIsStrInference([d_2406_v50_ for d_2423_i5_ in range(default__.abs(-607))]), d_2421_v59_, globalState)})
                d_2424_v61_: _dafny.MultiSet
                d_2424_v61_ = _dafny.MultiSet([d_2422_v60_])
                d_2425_v62_: _dafny.Array
                nw374_ = _dafny.Array(None, 27)
                nw374_[int(0)] = ((D2_DC7(len(d_2416_v54_), d_2414_cf15_, d_2405_v49_)).cf10) or ((self).f10)
                nw374_[int(1)] = not((self).f10)
                nw374_[int(2)] = not(self.f12)
                nw374_[int(3)] = not((d_2417_v55_).ispropersubset(d_2417_v55_))
                nw374_[int(4)] = (self).f10
                nw374_[int(5)] = (self).f10
                nw374_[int(6)] = (d_2408_v52_)[default__.safeIndex(d_2405_v49_, len(d_2408_v52_))]
                nw374_[int(7)] = (d_2405_v49_) >= (d_2405_v49_)
                nw374_[int(8)] = self.f12
                nw374_[int(9)] = not(d_2414_cf15_)
                nw374_[int(10)] = d_2414_cf15_
                nw374_[int(11)] = (d_2414_cf15_) == (not((self).f10))
                nw374_[int(12)] = (d_2418_v57_).cf66
                nw374_[int(13)] = False
                nw374_[int(14)] = (d_2405_v49_) < ((d_2420_v58_).cardinality)
                nw374_[int(15)] = (self).f11
                nw374_[int(16)] = self.f12
                nw374_[int(17)] = d_2413_cf16_
                nw374_[int(18)] = self.f12
                nw374_[int(19)] = not(not(not((self).fm10(globalState))))
                nw374_[int(20)] = False
                nw374_[int(21)] = d_2413_cf16_
                nw374_[int(22)] = (d_2417_v55_).ispropersubset(d_2417_v55_)
                nw374_[int(23)] = (d_2415_cf14_) == (d_2415_cf14_)
                nw374_[int(24)] = p1
                nw374_[int(25)] = not(self.f12)
                nw374_[int(26)] = (d_2422_v60_) not in (d_2424_v61_)
                d_2425_v62_ = nw374_
                index401_ = default__.safeIndex(267, (d_2425_v62_).length(0))
                (d_2425_v62_)[index401_] = (self).f10
                d_2426_v63_: _dafny.Array
                nw375_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_2426_v63_ = nw375_
                d_2427_v64_: D17
                d_2427_v64_ = D17_DC50(d_2426_v63_)
                pat_let_tv69_ = d_2426_v63_
                d_2428_v65_: _dafny.Array
                nw376_ = _dafny.Array(None, 15)
                nw376_[int(0)] = d_2427_v64_
                nw376_[int(1)] = d_2427_v64_
                nw376_[int(2)] = d_2427_v64_
                nw376_[int(3)] = d_2427_v64_
                def iife188_(_pat_let62_0):
                    def iife189_(d_2429_dt__update__tmp_h1_):
                        def iife190_(_pat_let63_0):
                            def iife191_(d_2430_dt__update_hcf86_h0_):
                                return D17_DC50(d_2430_dt__update_hcf86_h0_)
                            return iife191_(_pat_let63_0)
                        return iife190_(pat_let_tv69_)
                    return iife189_(_pat_let62_0)
                nw376_[int(4)] = iife188_(d_2427_v64_)
                nw376_[int(5)] = D17_DC50(d_2426_v63_)
                nw376_[int(6)] = d_2427_v64_
                nw376_[int(7)] = D17_DC50(d_2426_v63_)
                nw376_[int(8)] = d_2427_v64_
                nw376_[int(9)] = D17_DC50(d_2426_v63_)
                nw376_[int(10)] = d_2427_v64_
                nw376_[int(11)] = d_2427_v64_
                nw376_[int(12)] = d_2427_v64_
                nw376_[int(13)] = d_2427_v64_
                nw376_[int(14)] = D17_DC50(d_2426_v63_)
                d_2428_v65_ = nw376_
                index402_ = default__.safeIndex(4, (d_2428_v65_).length(0))
                (d_2428_v65_)[index402_] = d_2427_v64_
                index403_ = default__.safeIndex(267, (d_2425_v62_).length(0))
                index404_ = default__.safeIndex(4, (d_2428_v65_).length(0))
                rhs382_ = not(p1)
                rhs383_ = D17_DC50(d_2426_v63_)
                lhs270_ = d_2425_v62_
                lhs271_ = default__.safeIndex(267, (d_2425_v62_).length(0))
                lhs272_ = d_2428_v65_
                lhs273_ = default__.safeIndex(4, (d_2428_v65_).length(0))
                lhs270_[lhs271_] = rhs382_
                lhs272_[lhs273_] = rhs383_
                (globalState).f7 = d_2405_v49_
                index405_ = default__.safeIndex(267, (d_2425_v62_).length(0))
                (d_2425_v62_)[index405_] = d_2414_cf15_
                index406_ = default__.safeIndex(267, (d_2425_v62_).length(0))
                (d_2425_v62_)[index406_] = not(self.f12)
            elif source40_.is_DC11:
                d_2431___mcc_h3_ = source40_.cf17
                d_2432___mcc_h4_ = source40_.cf18
                d_2433_cf18_ = d_2432___mcc_h4_
                d_2434_cf17_ = d_2431___mcc_h3_
                (self).f12 = (self).fm10(globalState)
                (globalState).f0 = (d_2434_cf17_) - (131)
                (self).f12 = (self).fm10(globalState)
                d_2435_v66_: _dafny.Array
                def lambda112_(d_2436_i6_):
                    return not((self).f11)

                init60_ = lambda112_
                nw377_ = _dafny.Array(None, 9)
                for i0_60_ in range(nw377_.length(0)):
                    nw377_[i0_60_] = init60_(i0_60_)
                d_2435_v66_ = nw377_
                d_2437_v67_: _dafny.Set
                d_2437_v67_ = _dafny.Set({d_2433_cf18_})
                d_2438_v68_: _dafny.Map
                d_2438_v68_ = _dafny.Map({d_2435_v66_: default__.safeDivisionInt((0) - ((self).fm9(d_2404_v48_, d_2437_v67_, globalState)), 98)})
                d_2439_v69_: _dafny.Map
                d_2439_v69_ = _dafny.Map({d_2406_v50_: _dafny.Map({d_2435_v66_: d_2405_v49_})})
                rhs384_ = -101
                rhs385_ = (((d_2438_v68_).set(d_2435_v66_, d_2405_v49_)) | (d_2438_v68_)) | (((d_2439_v69_)[d_2406_v50_] if (d_2406_v50_) in (d_2439_v69_) else d_2438_v68_))
                rhs386_ = (0) - ((0) - (d_2434_cf17_))
                rhs387_ = self.f12
                lhs274_ = globalState
                lhs275_ = self
                d_2434_cf17_ = rhs384_
                d_2438_v68_ = rhs385_
                lhs274_.f0 = rhs386_
                lhs275_.f12 = rhs387_
            elif source40_.is_DC9:
                d_2440___mcc_h5_ = source40_.cf13
                d_2441_cf13_ = d_2440___mcc_h5_
                d_2442_v70_: _dafny.Map
                d_2442_v70_ = _dafny.Map({d_2405_v49_: d_2404_v48_})
                d_2443_v71_: _dafny.Map
                d_2443_v71_ = _dafny.Map({d_2405_v49_: self.f12})
                d_2444_v72_: _dafny.MultiSet
                d_2444_v72_ = _dafny.MultiSet([d_2405_v49_])
                d_2445_v73_: _dafny.Array
                nw378_ = _dafny.Array(None, 23)
                nw378_[int(0)] = (d_2405_v49_) > (d_2405_v49_)
                nw378_[int(1)] = not(p1)
                nw378_[int(2)] = (((d_2442_v70_)[len(d_2443_v71_)] if (len(d_2443_v71_)) in (d_2442_v70_) else d_2404_v48_)) < (_dafny.SeqWithoutIsStrInference([d_2406_v50_ for d_2446_i7_ in range(default__.abs(934))]))
                nw378_[int(3)] = (self).f10
                nw378_[int(4)] = (self).f11
                nw378_[int(5)] = self.f12
                nw378_[int(6)] = (self).f10
                nw378_[int(7)] = p1
                nw378_[int(8)] = (self).f11
                nw378_[int(9)] = False
                nw378_[int(10)] = (self).f10
                nw378_[int(11)] = self.f12
                nw378_[int(12)] = p1
                nw378_[int(13)] = (self).f10
                nw378_[int(14)] = p1
                nw378_[int(15)] = not((self).fm7((self).f10, globalState))
                nw378_[int(16)] = not (not(True)) or (p1)
                nw378_[int(17)] = p1
                nw378_[int(18)] = not(not((d_2405_v49_) >= (d_2405_v49_)))
                nw378_[int(19)] = (self).f11
                nw378_[int(20)] = (_dafny.MultiSet([222])).issubset(d_2444_v72_)
                nw378_[int(21)] = (self).f10
                nw378_[int(22)] = self.f12
                d_2445_v73_ = nw378_
                index407_ = default__.safeIndex(458, (d_2445_v73_).length(0))
                (d_2445_v73_)[index407_] = not(((self).f11) or (p1))
                index408_ = default__.safeIndex(458, (d_2445_v73_).length(0))
                (d_2445_v73_)[index408_] = (d_2408_v52_) == ((default__.fm40(self.f12, (self).f11, globalState)) + (d_2408_v52_))
                d_2447_v74_: _dafny.Array
                def lambda113_(d_2448_i8_):
                    return _dafny.CodePoint('c')

                init61_ = lambda113_
                nw379_ = _dafny.Array(None, 6)
                for i0_61_ in range(nw379_.length(0)):
                    nw379_[i0_61_] = init61_(i0_61_)
                d_2447_v74_ = nw379_
                d_2449_v75_: _dafny.Map
                d_2449_v75_ = _dafny.Map({d_2447_v74_: default__.fm1(globalState)})
                d_2449_v75_ = d_2449_v75_
                d_2450_v76_: _dafny.Set
                d_2450_v76_ = _dafny.Set({len(d_2404_v48_), len(_dafny.Map({d_2408_v52_: d_2405_v49_}))})
                index409_ = default__.safeIndex(458, (d_2445_v73_).length(0))
                (d_2445_v73_)[index409_] = (d_2405_v49_) not in (d_2450_v76_)
                d_2451_v77_: _dafny.Map
                d_2451_v77_ = _dafny.Map({d_2405_v49_: d_2445_v73_})
                d_2452_v78_: C11
                nw380_ = C11()
                nw380_.ctor__(len(d_2451_v77_), (d_2404_v48_) + (d_2404_v48_), (d_2444_v72_).issubset(d_2444_v72_))
                d_2452_v78_ = nw380_
            elif True:
                d_2453___mcc_h6_ = source40_.cf19
                d_2454_cf19_ = d_2453___mcc_h6_
                d_2455_v79_: _dafny.Array
                nw381_ = _dafny.Array(False, 16)
                d_2455_v79_ = nw381_
                index410_ = default__.safeIndex(773, (d_2455_v79_).length(0))
                (d_2455_v79_)[index410_] = (not((self).f10)) == ((self).f10)
                d_2456_v80_: _dafny.MultiSet
                d_2456_v80_ = _dafny.MultiSet([d_2405_v49_])
                index411_ = default__.safeIndex(773, (d_2455_v79_).length(0))
                rhs388_ = ((d_2456_v80_ if (self).f11 else default__.fm18(d_2405_v49_, default__.fm51(globalState), (self).f10, globalState))).ispropersubset(d_2456_v80_)
                rhs389_ = self.f12
                lhs276_ = self
                lhs277_ = d_2455_v79_
                lhs278_ = default__.safeIndex(773, (d_2455_v79_).length(0))
                lhs276_.f12 = rhs388_
                lhs277_[lhs278_] = rhs389_
                d_2457_v81_: _dafny.Map
                d_2457_v81_ = _dafny.Map({631: d_2407_v51_})
                d_2458_v82_: D14
                d_2458_v82_ = D14_DC41(d_2457_v81_)
                d_2458_v82_ = d_2458_v82_
                d_2459_v83_: D20
                d_2459_v83_ = D20_DC60(len(d_2404_v48_))
                d_2460_v84_: _dafny.Set
                d_2460_v84_ = _dafny.Set({(d_2455_v79_)[default__.safeIndex(773, (d_2455_v79_).length(0))]})
                d_2461_v85_: _dafny.Seq
                d_2461_v85_ = _dafny.SeqWithoutIsStrInference([(d_2456_v80_).cardinality])
                d_2462_v86_: _dafny.Array
                nw382_ = _dafny.Array(None, 7)
                nw382_[int(0)] = (d_2405_v49_ if p1 else d_2405_v49_)
                nw382_[int(1)] = d_2405_v49_
                nw382_[int(2)] = d_2405_v49_
                nw382_[int(3)] = d_2405_v49_
                nw382_[int(4)] = (d_2459_v83_).cf102
                nw382_[int(5)] = (d_2405_v49_) + ((0) - (len(d_2460_v84_)))
                nw382_[int(6)] = (d_2405_v49_) - (len(d_2461_v85_))
                d_2462_v86_ = nw382_
                index412_ = default__.safeIndex(516, (d_2462_v86_).length(0))
                (d_2462_v86_)[index412_] = (d_2405_v49_) + (d_2405_v49_)
                d_2463_v87_: D11
                d_2463_v87_ = D11_DC32(_dafny.Map({d_2405_v49_: (self).f10}))
                index413_ = default__.safeIndex(516, (d_2462_v86_).length(0))
                rhs390_ = (d_2405_v49_) * ((0) - ((0) - (((d_2456_v80_) - (d_2456_v80_)).cardinality)))
                rhs391_ = d_2405_v49_
                rhs392_ = ((d_2456_v80_)[(255 if p1 else len((d_2463_v87_).cf53))] if ((255 if p1 else len((d_2463_v87_).cf53))) in (d_2456_v80_) else len(d_2460_v84_))
                lhs279_ = globalState
                lhs280_ = globalState
                lhs281_ = d_2462_v86_
                lhs282_ = default__.safeIndex(516, (d_2462_v86_).length(0))
                lhs279_.f0 = rhs390_
                lhs280_.f0 = rhs391_
                lhs281_[lhs282_] = rhs392_
                d_2464_v88_: _dafny.Map
                d_2464_v88_ = _dafny.Map({(d_2462_v86_)[default__.safeIndex(516, (d_2462_v86_).length(0))]: self.f12})
                d_2465_v89_: C9
                nw383_ = C9()
                nw383_.ctor__((self).f11, ((d_2464_v88_)[(d_2462_v86_)[default__.safeIndex(516, (d_2462_v86_).length(0))]] if ((d_2462_v86_)[default__.safeIndex(516, (d_2462_v86_).length(0))]) in (d_2464_v88_) else True))
                d_2465_v89_ = nw383_
            d_2466_v90_: _dafny.Array
            nw384_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_2466_v90_ = nw384_
            d_2467_v91_: _dafny.MultiSet
            d_2467_v91_ = _dafny.MultiSet([(self).f11, (self).f10])
            d_2468_v92_: _dafny.Seq
            d_2468_v92_ = _dafny.SeqWithoutIsStrInference([d_2405_v49_, (d_2467_v91_).cardinality])
            d_2469_v93_: C9
            nw385_ = C9()
            nw385_.ctor__(self.f12, not(True))
            d_2469_v93_ = nw385_
            d_2470_v94_: _dafny.Seq
            d_2470_v94_ = _dafny.SeqWithoutIsStrInference([d_2469_v93_])
            d_2471_v95_: D19
            d_2471_v95_ = D19_DC58(len(d_2470_v94_), (self).f10)
            d_2472_v96_: _dafny.Array
            nw386_ = _dafny.Array(None, 16)
            nw386_[int(0)] = d_2405_v49_
            nw386_[int(1)] = default__.fm1(globalState)
            nw386_[int(2)] = 629
            nw386_[int(3)] = len(d_2468_v92_)
            nw386_[int(4)] = d_2405_v49_
            nw386_[int(5)] = len(_dafny.SeqWithoutIsStrInference([d_2406_v50_ for d_2473_i9_ in range(default__.abs(-399))]))
            nw386_[int(6)] = len(d_2404_v48_)
            nw386_[int(7)] = (d_2471_v95_).cf99
            nw386_[int(8)] = 83
            nw386_[int(9)] = -750
            nw386_[int(10)] = d_2405_v49_
            nw386_[int(11)] = d_2405_v49_
            nw386_[int(12)] = 208
            nw386_[int(13)] = d_2405_v49_
            nw386_[int(14)] = default__.fm1(globalState)
            nw386_[int(15)] = d_2405_v49_
            d_2472_v96_ = nw386_
            index414_ = default__.safeIndex(285, (d_2466_v90_).length(0))
            (d_2466_v90_)[index414_] = d_2472_v96_
            d_2474_v97_: D3
            d_2474_v97_ = D3_DC12(D3_DC11((0) - (d_2405_v49_), d_2407_v51_))
            d_2475_v98_: _dafny.Set
            d_2475_v98_ = _dafny.Set({self.f12, self.f12})
            d_2476_v99_: _dafny.Set
            d_2476_v99_ = _dafny.Set({d_2475_v98_})
            d_2477_v100_: _dafny.Map
            d_2477_v100_ = _dafny.Map({D3_DC12(D3_DC9(d_2476_v99_)): True})
            d_2478_v101_: _dafny.Map
            d_2478_v101_ = _dafny.Map({(d_2469_v93_).f17: 698})
            d_2479_v102_: _dafny.MultiSet
            d_2479_v102_ = _dafny.MultiSet([d_2405_v49_])
            d_2480_v103_: _dafny.Array
            nw387_ = _dafny.Array(None, 19)
            nw387_[int(0)] = (d_2469_v93_).f17
            nw387_[int(1)] = self.f12
            nw387_[int(2)] = (d_2474_v97_) not in (d_2477_v100_)
            nw387_[int(3)] = (d_2404_v48_) < (d_2404_v48_)
            nw387_[int(4)] = (d_2405_v49_) == (d_2405_v49_)
            nw387_[int(5)] = (d_2405_v49_) >= (len(d_2407_v51_))
            nw387_[int(6)] = (self).f10
            nw387_[int(7)] = (self).f10
            nw387_[int(8)] = (d_2467_v91_).ispropersubset(_dafny.MultiSet([(d_2469_v93_).f17]))
            nw387_[int(9)] = (not((self).f10)) in (d_2478_v101_)
            nw387_[int(10)] = self.f12
            nw387_[int(11)] = (d_2479_v102_).ispropersubset(_dafny.MultiSet([d_2405_v49_, len(d_2404_v48_)]))
            nw387_[int(12)] = not(True)
            nw387_[int(13)] = p1
            nw387_[int(14)] = (d_2469_v93_).f17
            nw387_[int(15)] = (self).f11
            nw387_[int(16)] = (d_2469_v93_).f17
            nw387_[int(17)] = (d_2469_v93_).f17
            nw387_[int(18)] = (self).fm6(d_2405_v49_, (0) - (len((d_2404_v48_).set(default__.safeIndex(len(d_2408_v52_), len(d_2404_v48_)), d_2406_v50_))), d_2405_v49_, self.f12, globalState)
            d_2480_v103_ = nw387_
            index415_ = default__.safeIndex(377, (d_2480_v103_).length(0))
            (d_2480_v103_)[index415_] = (_dafny.SeqWithoutIsStrInference([d_2406_v50_ for d_2481_i10_ in range(default__.abs(-354))])) == (d_2404_v48_)
            d_2482_v104_: D15
            d_2482_v104_ = D15_DC46()
            index416_ = default__.safeIndex(285, (d_2466_v90_).length(0))
            index417_ = default__.safeIndex(377, (d_2480_v103_).length(0))
            rhs393_ = (_dafny.SeqWithoutIsStrInference([d_2406_v50_ for d_2483_i11_ in range(default__.abs(-732))])) < (d_2404_v48_)
            rhs394_ = d_2472_v96_
            rhs395_ = (d_2469_v93_).f17
            rhs396_ = D15_DC46()
            lhs283_ = d_2466_v90_
            lhs284_ = default__.safeIndex(285, (d_2466_v90_).length(0))
            lhs285_ = d_2480_v103_
            lhs286_ = default__.safeIndex(377, (d_2480_v103_).length(0))
            r0 = rhs393_
            lhs283_[lhs284_] = rhs394_
            lhs285_[lhs286_] = rhs395_
            d_2482_v104_ = rhs396_
            (globalState).f7 = (d_2405_v49_) + (d_2405_v49_)
        d_2484_v105_: _dafny.Seq
        d_2484_v105_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kaqhnb"))
        d_2485_v106_: int
        d_2485_v106_ = 557
        d_2486_v107_: T0
        nw388_ = C11()
        nw388_.ctor__(d_2485_v106_, d_2484_v105_, p1)
        d_2486_v107_ = nw388_
        d_2487_v108_: _dafny.Seq
        d_2487_v108_ = _dafny.SeqWithoutIsStrInference([d_2486_v107_, d_2486_v107_, d_2486_v107_])
        d_2488_v109_: _dafny.Map
        d_2488_v109_ = _dafny.Map({d_2487_v108_: self.f12})
        d_2489_v110_: _dafny.Seq
        d_2489_v110_ = _dafny.SeqWithoutIsStrInference([(d_2486_v107_).f10])
        d_2490_v111_: D0
        d_2490_v111_ = D0_DC1(d_2484_v105_)
        d_2491_v112_: _dafny.Map
        d_2491_v112_ = _dafny.Map({(self).f11: d_2484_v105_})
        d_2492_v113_: _dafny.Array
        nw389_ = _dafny.Array(None, 26)
        nw389_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prqda"))
        nw389_[int(1)] = d_2484_v105_
        nw389_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wx"))) + (d_2484_v105_)
        nw389_[int(3)] = d_2484_v105_
        nw389_[int(4)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjuvsa"))).set(default__.safeIndex(d_2485_v106_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjuvsa")))), _dafny.CodePoint('h')) if ((d_2488_v109_)[d_2487_v108_] if (d_2487_v108_) in (d_2488_v109_) else (self).f10) else d_2484_v105_)
        nw389_[int(5)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2493_i12_ in range(default__.abs(-557))])) + (d_2484_v105_)
        nw389_[int(6)] = (d_2484_v105_) + (d_2484_v105_)
        nw389_[int(7)] = d_2484_v105_
        nw389_[int(8)] = d_2484_v105_
        nw389_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnfj"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2494_i13_ in range(default__.abs(294))]))
        nw389_[int(10)] = (default__.fm13((d_2486_v107_).f10, (d_2489_v110_)[default__.safeIndex(d_2485_v106_, len(d_2489_v110_))], globalState)) + (d_2484_v105_)
        nw389_[int(11)] = d_2484_v105_
        nw389_[int(12)] = d_2484_v105_
        nw389_[int(13)] = d_2484_v105_
        nw389_[int(14)] = d_2484_v105_
        nw389_[int(15)] = d_2484_v105_
        nw389_[int(16)] = d_2484_v105_
        nw389_[int(17)] = d_2484_v105_
        nw389_[int(18)] = (d_2484_v105_ if not((self).f10) else d_2484_v105_)
        nw389_[int(19)] = d_2484_v105_
        nw389_[int(20)] = (d_2490_v111_).cf1
        nw389_[int(21)] = d_2484_v105_
        nw389_[int(22)] = d_2484_v105_
        nw389_[int(23)] = d_2484_v105_
        nw389_[int(24)] = d_2484_v105_
        nw389_[int(25)] = (((d_2491_v112_)[(d_2486_v107_).f10] if ((d_2486_v107_).f10) in (d_2491_v112_) else d_2484_v105_)) + (d_2484_v105_)
        d_2492_v113_ = nw389_
        index418_ = default__.safeIndex(176, (d_2492_v113_).length(0))
        (d_2492_v113_)[index418_] = d_2484_v105_
        index419_ = default__.safeIndex(176, (d_2492_v113_).length(0))
        (d_2492_v113_)[index419_] = d_2484_v105_
        d_2495_v114_: str
        d_2495_v114_ = _dafny.CodePoint('a')
        (self).f12 = (d_2495_v114_) in (d_2484_v105_)
        d_2496_v115_: D15
        d_2496_v115_ = D15_DC45(d_2485_v106_, p1, p1)
        (globalState).f0 = (d_2496_v115_).cf80
        (self).f12 = (d_2485_v106_) <= (default__.fm1(globalState))
        d_2497_v116_: _dafny.Map
        d_2497_v116_ = _dafny.Map({d_2485_v106_: 330})
        d_2498_v117_: D7
        d_2498_v117_ = D7_DC22((d_2497_v116_) | (d_2497_v116_))
        d_2498_v117_ = d_2498_v117_
        r0 = (d_2485_v106_) != ((0) - (default__.safeModuloInt(-380, d_2485_v106_)))
        r1 = d_2484_v105_
        r2 = d_2497_v116_
        return r0, r1, r2

    def m4(self, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: int = int(0)
        d_2499_v0_: int
        d_2499_v0_ = -364
        d_2500_v1_: _dafny.Array
        def lambda114_(d_2501_v0_):
            def lambda115_(d_2502_i1_):
                return (d_2502_i1_) * (d_2501_v0_)

            return lambda115_

        init62_ = lambda114_(d_2499_v0_)
        nw390_ = _dafny.Array(None, 10)
        for i0_62_ in range(nw390_.length(0)):
            nw390_[i0_62_] = init62_(i0_62_)
        d_2500_v1_ = nw390_
        d_2503_v2_: D10
        d_2503_v2_ = D10_DC31((self).f10, self.f12, d_2499_v0_)
        d_2504_v3_: _dafny.Seq
        d_2504_v3_ = _dafny.SeqWithoutIsStrInference([(d_2503_v2_).cf50])
        d_2505_v4_: D14
        d_2505_v4_ = D14_DC42(D1_DC4(d_2499_v0_, d_2500_v1_), d_2504_v3_, (self).f11)
        d_2506_i0_: int
        d_2506_i0_ = 0
        with _dafny.label("14"):
            while (d_2505_v4_).cf74:
                with _dafny.c_label("14"):
                    if (d_2506_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_2506_i0_ = (d_2506_i0_) + (1)
                    d_2507_v5_: _dafny.Array
                    nw391_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                    d_2507_v5_ = nw391_
                    d_2508_v6_: _dafny.Seq
                    d_2508_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbbgv"))
                    index420_ = default__.safeIndex(985, (d_2507_v5_).length(0))
                    (d_2507_v5_)[index420_] = (d_2508_v6_) + (d_2508_v6_)
                    index421_ = default__.safeIndex(985, (d_2507_v5_).length(0))
                    (d_2507_v5_)[index421_] = d_2508_v6_
                    d_2509_v7_: str
                    d_2509_v7_ = _dafny.CodePoint('u')
                    d_2509_v7_ = d_2509_v7_
                    d_2510_v8_: D6
                    d_2510_v8_ = D6_DC20(self.f12, False, (self).f11, d_2499_v0_)
                    d_2511_v9_: D4
                    d_2511_v9_ = D4_DC14(len(_dafny.Set({(self).f11, (self).f11, (d_2510_v8_).cf34})), d_2499_v0_)
                    source41_ = d_2511_v9_
                    if source41_.is_DC14:
                        d_2512___mcc_h0_ = source41_.cf21
                        d_2513___mcc_h1_ = source41_.cf22
                        d_2514_cf22_ = d_2513___mcc_h1_
                        d_2515_cf21_ = d_2512___mcc_h0_
                        d_2516_v10_: C8
                        nw392_ = C8()
                        nw392_.ctor__(d_2509_v7_)
                        d_2516_v10_ = nw392_
                        d_2517_v11_: _dafny.Array
                        nw393_ = _dafny.Array(False, 10)
                        d_2517_v11_ = nw393_
                        index422_ = default__.safeIndex(139, (d_2517_v11_).length(0))
                        (d_2517_v11_)[index422_] = self.f12
                        index423_ = default__.safeIndex(139, (d_2517_v11_).length(0))
                        (d_2517_v11_)[index423_] = (self).f10
                        d_2518_v12_: C3
                        nw394_ = C3()
                        nw394_.ctor__((d_2515_cf21_) == (d_2499_v0_), (len((d_2507_v5_)[default__.safeIndex(985, (d_2507_v5_).length(0))])) not in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len((d_2507_v5_)[default__.safeIndex(985, (d_2507_v5_).length(0))]) for d_2519_i3_ in range(default__.abs(-525))]))) for d_2520_i2_ in range(default__.abs(870))])))
                        d_2518_v12_ = nw394_
                        d_2521_v13_: _dafny.Map
                        d_2521_v13_ = _dafny.Map({d_2514_cf22_: (d_2517_v11_)[default__.safeIndex(139, (d_2517_v11_).length(0))]})
                        index424_ = default__.safeIndex(95, (d_2500_v1_).length(0))
                        (d_2500_v1_)[index424_] = 319
                        index425_ = default__.safeIndex(985, (d_2507_v5_).length(0))
                        index426_ = default__.safeIndex(95, (d_2500_v1_).length(0))
                        rhs397_ = (_dafny.Map({d_2515_cf21_: (self).f10})) | (_dafny.Map({d_2515_cf21_: (self).f11}))
                        rhs398_ = (d_2518_v12_).fm8(d_2514_cf22_, True, (d_2514_cf22_) == (304), globalState)
                        rhs399_ = (d_2507_v5_)[default__.safeIndex(985, (d_2507_v5_).length(0))]
                        rhs400_ = False
                        rhs401_ = d_2514_cf22_
                        lhs287_ = self
                        lhs288_ = d_2507_v5_
                        lhs289_ = default__.safeIndex(985, (d_2507_v5_).length(0))
                        lhs290_ = d_2500_v1_
                        lhs291_ = default__.safeIndex(95, (d_2500_v1_).length(0))
                        d_2521_v13_ = rhs397_
                        lhs287_.f12 = rhs398_
                        lhs288_[lhs289_] = rhs399_
                        r1 = rhs400_
                        lhs290_[lhs291_] = rhs401_
                    elif source41_.is_DC15:
                        d_2522___mcc_h2_ = source41_.cf23
                        d_2523___mcc_h3_ = source41_.cf24
                        d_2524___mcc_h4_ = source41_.cf25
                        d_2525___mcc_h5_ = source41_.cf26
                        d_2526_cf26_ = d_2525___mcc_h5_
                        d_2527_cf25_ = d_2524___mcc_h4_
                        d_2528_cf24_ = d_2523___mcc_h3_
                        d_2529_cf23_ = d_2522___mcc_h2_
                        (globalState).f0 = default__.fm1(globalState)
                        d_2530_v14_: _dafny.Array
                        nw395_ = _dafny.Array(False, 7)
                        d_2530_v14_ = nw395_
                        d_2531_v15_: _dafny.Array
                        nw396_ = _dafny.Array(None, 2)
                        nw396_[int(0)] = d_2530_v14_
                        nw396_[int(1)] = d_2530_v14_
                        d_2531_v15_ = nw396_
                        index427_ = default__.safeIndex(93, (d_2531_v15_).length(0))
                        (d_2531_v15_)[index427_] = d_2530_v14_
                        index428_ = default__.safeIndex(93, (d_2531_v15_).length(0))
                        (d_2531_v15_)[index428_] = d_2530_v14_
                        r1 = not(False)
                        rhs402_ = d_2500_v1_
                        rhs403_ = d_2529_cf23_
                        rhs404_ = (self).f11
                        rhs405_ = (self).fm7((self).f10, globalState)
                        lhs292_ = globalState
                        d_2500_v1_ = rhs402_
                        lhs292_.f7 = rhs403_
                        r1 = rhs404_
                        r1 = rhs405_
                    elif True:
                        d_2532___mcc_h6_ = source41_.cf20
                        d_2533_cf20_ = d_2532___mcc_h6_
                        d_2534_v16_: _dafny.Array
                        def lambda116_(d_2535_i4_):
                            return (self).f11

                        init63_ = lambda116_
                        nw397_ = _dafny.Array(None, 3)
                        for i0_63_ in range(nw397_.length(0)):
                            nw397_[i0_63_] = init63_(i0_63_)
                        d_2534_v16_ = nw397_
                        index429_ = default__.safeIndex(989, (d_2534_v16_).length(0))
                        (d_2534_v16_)[index429_] = False
                        d_2536_v17_: _dafny.Set
                        d_2536_v17_ = _dafny.Set({(d_2534_v16_ if (self).f10 else d_2534_v16_)})
                        index430_ = default__.safeIndex(989, (d_2534_v16_).length(0))
                        rhs406_ = (d_2509_v7_) not in (_dafny.SeqWithoutIsStrInference([d_2509_v7_ for d_2537_i5_ in range(default__.abs(400))]))
                        rhs407_ = d_2499_v0_
                        rhs408_ = default__.fm0(d_2499_v0_, globalState)
                        rhs409_ = (d_2536_v17_) | ((d_2536_v17_) | (d_2536_v17_))
                        rhs410_ = d_2499_v0_
                        lhs293_ = globalState
                        lhs294_ = d_2534_v16_
                        lhs295_ = default__.safeIndex(989, (d_2534_v16_).length(0))
                        lhs296_ = globalState
                        r1 = rhs406_
                        lhs293_.f7 = rhs407_
                        lhs294_[lhs295_] = rhs408_
                        d_2536_v17_ = rhs409_
                        lhs296_.f7 = rhs410_
                        index431_ = default__.safeIndex(131, (d_2500_v1_).length(0))
                        (d_2500_v1_)[index431_] = 768
                        d_2538_v18_: _dafny.Map
                        d_2538_v18_ = _dafny.Map({(d_2534_v16_)[default__.safeIndex(989, (d_2534_v16_).length(0))]: d_2499_v0_})
                        d_2539_v19_: _dafny.Map
                        d_2539_v19_ = _dafny.Map({d_2499_v0_: 653})
                        d_2540_v20_: _dafny.Seq
                        d_2540_v20_ = _dafny.SeqWithoutIsStrInference([d_2539_v19_])
                        d_2541_v21_: _dafny.MultiSet
                        d_2541_v21_ = _dafny.MultiSet([default__.fm0(11, globalState)])
                        index432_ = default__.safeIndex(131, (d_2500_v1_).length(0))
                        rhs411_ = (d_2538_v18_) == ((d_2538_v18_) | (d_2538_v18_))
                        rhs412_ = (d_2499_v0_) - (911)
                        rhs413_ = (self.f12) and (self.f12)
                        rhs414_ = (len((d_2540_v20_)[default__.safeIndex((d_2541_v21_).cardinality, len(d_2540_v20_))])) - (d_2499_v0_)
                        lhs297_ = globalState
                        lhs298_ = self
                        lhs299_ = d_2500_v1_
                        lhs300_ = default__.safeIndex(131, (d_2500_v1_).length(0))
                        r1 = rhs411_
                        lhs297_.f0 = rhs412_
                        lhs298_.f12 = rhs413_
                        lhs299_[lhs300_] = rhs414_
                        d_2509_v7_ = d_2509_v7_
                        d_2542_v22_: _dafny.MultiSet
                        d_2542_v22_ = _dafny.MultiSet([(d_2500_v1_)[default__.safeIndex(131, (d_2500_v1_).length(0))]])
                        d_2543_v23_: C11
                        nw398_ = C11()
                        nw398_.ctor__((d_2542_v22_).cardinality, d_2508_v6_, (self).f10)
                        d_2543_v23_ = nw398_
                        d_2544_v24_: _dafny.Map
                        d_2544_v24_ = _dafny.Map({d_2543_v23_: (self).f10})
                        d_2544_v24_ = d_2544_v24_
                    if (self).fm8(len((d_2507_v5_)[default__.safeIndex(985, (d_2507_v5_).length(0))]), (self).f10, (default__.fm40((self).f10, (self).f10, globalState)) <= (d_2504_v3_), globalState):
                        index433_ = default__.safeIndex(886, (d_2500_v1_).length(0))
                        (d_2500_v1_)[index433_] = (d_2499_v0_) * (d_2499_v0_)
                        index434_ = default__.safeIndex(59, (d_2500_v1_).length(0))
                        (d_2500_v1_)[index434_] = d_2499_v0_
                        index435_ = default__.safeIndex(886, (d_2500_v1_).length(0))
                        index436_ = default__.safeIndex(59, (d_2500_v1_).length(0))
                        rhs415_ = d_2499_v0_
                        rhs416_ = d_2499_v0_
                        lhs301_ = d_2500_v1_
                        lhs302_ = default__.safeIndex(886, (d_2500_v1_).length(0))
                        lhs303_ = d_2500_v1_
                        lhs304_ = default__.safeIndex(59, (d_2500_v1_).length(0))
                        lhs301_[lhs302_] = rhs415_
                        lhs303_[lhs304_] = rhs416_
                        d_2545_v25_: _dafny.Array
                        nw399_ = _dafny.Array(_dafny.Array(None, 0), 24)
                        d_2545_v25_ = nw399_
                        rhs417_ = default__.safeDivisionInt((d_2500_v1_)[default__.safeIndex(886, (d_2500_v1_).length(0))], (d_2500_v1_)[default__.safeIndex(886, (d_2500_v1_).length(0))])
                        rhs418_ = not((self).f10)
                        rhs419_ = (d_2545_v25_ if (self).f10 else d_2545_v25_)
                        lhs305_ = self
                        r2 = rhs417_
                        lhs305_.f12 = rhs418_
                        d_2545_v25_ = rhs419_
                        d_2546_v26_: _dafny.Map
                        d_2546_v26_ = _dafny.Map({-881: d_2499_v0_})
                        (self).f12 = ((d_2500_v1_)[default__.safeIndex(886, (d_2500_v1_).length(0))]) < ((0) - (default__.safeDivisionInt(d_2499_v0_, len(d_2546_v26_))))
                        d_2547_v27_: _dafny.MultiSet
                        d_2547_v27_ = _dafny.MultiSet([(self).f10])
                        d_2548_v28_: D15
                        d_2548_v28_ = D15_DC45((d_2500_v1_)[default__.safeIndex(886, (d_2500_v1_).length(0))], (self).f10, (self).f11)
                        d_2549_v29_: _dafny.Array
                        nw400_ = _dafny.Array(None, 14)
                        nw400_[int(0)] = d_2548_v28_
                        nw400_[int(1)] = d_2548_v28_
                        nw400_[int(2)] = d_2548_v28_
                        nw400_[int(3)] = d_2548_v28_
                        nw400_[int(4)] = d_2548_v28_
                        nw400_[int(5)] = d_2548_v28_
                        nw400_[int(6)] = D15_DC45(434, (self).f10, self.f12)
                        nw400_[int(7)] = d_2548_v28_
                        nw400_[int(8)] = d_2548_v28_
                        nw400_[int(9)] = default__.fm57(globalState)
                        nw400_[int(10)] = D15_DC45(942, not((self).f11), (self).f11)
                        nw400_[int(11)] = default__.fm57(globalState)
                        nw400_[int(12)] = d_2548_v28_
                        nw400_[int(13)] = d_2548_v28_
                        d_2549_v29_ = nw400_
                        index437_ = default__.safeIndex(105, (d_2549_v29_).length(0))
                        (d_2549_v29_)[index437_] = d_2548_v28_
                        d_2550_v30_: _dafny.MultiSet
                        d_2550_v30_ = _dafny.MultiSet([(d_2500_v1_)[default__.safeIndex(886, (d_2500_v1_).length(0))]])
                        index438_ = default__.safeIndex(105, (d_2549_v29_).length(0))
                        index439_ = default__.safeIndex(886, (d_2500_v1_).length(0))
                        rhs420_ = d_2547_v27_
                        rhs421_ = d_2548_v28_
                        rhs422_ = (d_2500_v1_)[default__.safeIndex(886, (d_2500_v1_).length(0))]
                        rhs423_ = d_2508_v6_
                        rhs424_ = default__.fm0((d_2550_v30_).cardinality, globalState)
                        lhs306_ = d_2549_v29_
                        lhs307_ = default__.safeIndex(105, (d_2549_v29_).length(0))
                        lhs308_ = d_2500_v1_
                        lhs309_ = default__.safeIndex(886, (d_2500_v1_).length(0))
                        d_2547_v27_ = rhs420_
                        lhs306_[lhs307_] = rhs421_
                        lhs308_[lhs309_] = rhs422_
                        d_2508_v6_ = rhs423_
                        r1 = rhs424_
                        (self).f12 = self.f12
                    elif True:
                        rhs425_ = d_2499_v0_
                        rhs426_ = d_2499_v0_
                        lhs310_ = globalState
                        lhs311_ = globalState
                        lhs310_.f7 = rhs425_
                        lhs311_.f7 = rhs426_
                        (globalState).f1 = (d_2499_v0_) - ((0) - (d_2499_v0_))
                        r1 = (self).f10
                        d_2551_v31_: _dafny.MultiSet
                        d_2551_v31_ = _dafny.MultiSet([d_2499_v0_, d_2499_v0_, d_2499_v0_])
                        d_2552_v32_: C6
                        nw401_ = C6()
                        nw401_.ctor__()
                        d_2552_v32_ = nw401_
                        d_2553_v33_: _dafny.Map
                        d_2553_v33_ = _dafny.Map({d_2552_v32_: (self).f11})
                        d_2554_v34_: _dafny.Map
                        d_2554_v34_ = _dafny.Map({d_2551_v31_: d_2553_v33_})
                        r2 = len(d_2554_v34_)
                        (self).f12 = self.f12
                    pass
            pass
        hi16_ = d_2499_v0_
        for d_2555_i6_ in range(default__.fm1(globalState), hi16_):
            index440_ = default__.safeIndex(169, (d_2500_v1_).length(0))
            (d_2500_v1_)[index440_] = d_2499_v0_
            index441_ = default__.safeIndex(169, (d_2500_v1_).length(0))
            (d_2500_v1_)[index441_] = default__.fm1(globalState)
            if (self).f11:
                (globalState).f7 = -961
                (globalState).f7 = d_2499_v0_
                (self).f12 = (self.f12) or ((self).f10)
                d_2556_v35_: _dafny.Array
                def lambda117_(d_2557_i7_):
                    return (_dafny.MultiSet([D5_DC16(_dafny.MultiSet([(self).f11]))])).isdisjoint(_dafny.MultiSet([D5_DC16(_dafny.MultiSet([(self).f11, (self).f10])), D5_DC16(_dafny.MultiSet([(self).f10]))]))

                init64_ = lambda117_
                nw402_ = _dafny.Array(None, 7)
                for i0_64_ in range(nw402_.length(0)):
                    nw402_[i0_64_] = init64_(i0_64_)
                d_2556_v35_ = nw402_
                index442_ = default__.safeIndex(888, (d_2556_v35_).length(0))
                (d_2556_v35_)[index442_] = (self).f11
                index443_ = default__.safeIndex(888, (d_2556_v35_).length(0))
                (d_2556_v35_)[index443_] = not ((d_2499_v0_) != (d_2555_i6_)) or ((self).f10)
                d_2558_v36_: _dafny.Set
                d_2558_v36_ = _dafny.Set({(d_2556_v35_)[default__.safeIndex(888, (d_2556_v35_).length(0))], (self).f11})
                (globalState).f7 = default__.safeModuloInt((d_2500_v1_)[default__.safeIndex(169, (d_2500_v1_).length(0))], len(d_2558_v36_))
            elif True:
                d_2559_v37_: str
                d_2559_v37_ = _dafny.CodePoint('i')
                d_2560_v38_: _dafny.Array
                nw403_ = _dafny.Array(None, 3)
                nw403_[int(0)] = d_2559_v37_
                nw403_[int(1)] = d_2559_v37_
                nw403_[int(2)] = d_2559_v37_
                d_2560_v38_ = nw403_
                d_2561_v39_: D12
                d_2561_v39_ = D12_DC37(-726, d_2560_v38_, d_2559_v37_)
                d_2562_v40_: D13
                d_2562_v40_ = D13_DC40(d_2561_v39_, (self).f10, d_2555_i6_, (d_2500_v1_)[default__.safeIndex(169, (d_2500_v1_).length(0))])
                (globalState).f0 = (d_2562_v40_).cf70
                d_2563_v41_: _dafny.Map
                d_2563_v41_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([self.f12]): (self).f11})
                d_2564_v42_: C7
                nw404_ = C7()
                nw404_.ctor__(((d_2563_v41_)[_dafny.SeqWithoutIsStrInference([(self).f11])] if (_dafny.SeqWithoutIsStrInference([(self).f11])) in (d_2563_v41_) else (self).f10), False)
                d_2564_v42_ = nw404_
                d_2565_v43_: C10
                nw405_ = C10()
                nw405_.ctor__(d_2499_v0_, (self).f10)
                d_2565_v43_ = nw405_
                d_2566_v44_: _dafny.Set
                d_2566_v44_ = _dafny.Set({((0) - (d_2499_v0_) if (d_2564_v42_).f19 else d_2499_v0_), (d_2565_v43_).f15, d_2555_i6_, d_2499_v0_, (d_2499_v0_) - ((0) - (d_2499_v0_))})
                d_2566_v44_ = ((default__.fm26(globalState) if (self).f10 else d_2566_v44_)) | (_dafny.Set({(d_2565_v43_).f15}))
                d_2567_v45_: _dafny.Array
                def lambda118_(d_2568_i8_):
                    return (self).f10

                init65_ = lambda118_
                nw406_ = _dafny.Array(None, 15)
                for i0_65_ in range(nw406_.length(0)):
                    nw406_[i0_65_] = init65_(i0_65_)
                d_2567_v45_ = nw406_
                d_2569_v46_: _dafny.Map
                d_2569_v46_ = _dafny.Map({752: d_2567_v45_})
                d_2570_v47_: _dafny.Array
                nw407_ = _dafny.Array(None, 18)
                nw407_[int(0)] = d_2567_v45_
                nw407_[int(1)] = d_2567_v45_
                nw407_[int(2)] = d_2567_v45_
                nw407_[int(3)] = d_2567_v45_
                nw407_[int(4)] = d_2567_v45_
                nw407_[int(5)] = d_2567_v45_
                nw407_[int(6)] = ((d_2569_v46_)[(d_2500_v1_)[default__.safeIndex(169, (d_2500_v1_).length(0))]] if ((d_2500_v1_)[default__.safeIndex(169, (d_2500_v1_).length(0))]) in (d_2569_v46_) else d_2567_v45_)
                nw407_[int(7)] = d_2567_v45_
                nw407_[int(8)] = d_2567_v45_
                nw407_[int(9)] = d_2567_v45_
                nw407_[int(10)] = d_2567_v45_
                nw407_[int(11)] = d_2567_v45_
                nw407_[int(12)] = d_2567_v45_
                nw407_[int(13)] = d_2567_v45_
                nw407_[int(14)] = d_2567_v45_
                nw407_[int(15)] = d_2567_v45_
                nw407_[int(16)] = d_2567_v45_
                nw407_[int(17)] = d_2567_v45_
                d_2570_v47_ = nw407_
                d_2571_v48_: D17
                d_2571_v48_ = D17_DC50(d_2570_v47_)
                d_2572_v49_: _dafny.Array
                nw408_ = _dafny.Array(None, 2)
                nw408_[int(0)] = d_2571_v48_
                nw408_[int(1)] = d_2571_v48_
                d_2572_v49_ = nw408_
                d_2572_v49_ = d_2572_v49_
            d_2573_v50_: _dafny.Seq
            d_2573_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idhb"))
            d_2574_v51_: _dafny.MultiSet
            d_2574_v51_ = _dafny.MultiSet([d_2555_i6_, (0) - (len(d_2573_v50_)), d_2555_i6_])
            d_2575_v52_: C6
            nw409_ = C6()
            nw409_.ctor__()
            d_2575_v52_ = nw409_
            d_2576_v53_: _dafny.Seq
            d_2576_v53_ = _dafny.SeqWithoutIsStrInference([(((d_2574_v51_)[d_2499_v0_] if (d_2499_v0_) in (d_2574_v51_) else (d_2500_v1_)[default__.safeIndex(169, (d_2500_v1_).length(0))])) - (len(_dafny.SeqWithoutIsStrInference([d_2575_v52_]))), (d_2500_v1_)[default__.safeIndex(169, (d_2500_v1_).length(0))], d_2555_i6_])
            (globalState).f7 = (d_2576_v53_)[default__.safeIndex(d_2555_i6_, len(d_2576_v53_))]
            (self).f12 = (d_2504_v3_)[default__.safeIndex(len((d_2573_v50_) + (d_2573_v50_)), len(d_2504_v3_))]
        d_2577_v54_: _dafny.Seq
        d_2577_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjbsmqq"))
        d_2578_v55_: str
        d_2578_v55_ = _dafny.CodePoint('i')
        d_2579_v56_: _dafny.Seq
        d_2579_v56_ = _dafny.SeqWithoutIsStrInference([(d_2577_v54_).set(default__.safeIndex(len(d_2577_v54_), len(d_2577_v54_)), d_2578_v55_)])
        d_2580_v57_: _dafny.MultiSet
        d_2580_v57_ = _dafny.MultiSet([d_2577_v54_, (d_2579_v56_)[default__.safeIndex(d_2499_v0_, len(d_2579_v56_))]])
        d_2581_v58_: bool
        d_2582_v59_: int
        out70_: bool
        out71_: int
        out70_, out71_ = (self).m6(d_2499_v0_, self.f12, d_2580_v57_, globalState)
        d_2581_v58_ = out70_
        d_2582_v59_ = out71_
        hi17_ = default__.safeModuloInt(d_2582_v59_, d_2582_v59_)
        for d_2583_i9_ in range(-382, hi17_):
            (self).f12 = (d_2583_i9_) in (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_2499_v0_ for d_2584_i10_ in range(default__.abs(730))]) if d_2581_v58_ else _dafny.SeqWithoutIsStrInference([d_2582_v59_]))))
            d_2585_v60_: _dafny.Map
            d_2585_v60_ = _dafny.Map({(self).f10: self.f12})
            d_2585_v60_ = (d_2585_v60_).set(d_2581_v58_, not(not(not((self).f11))))
            if default__.fm0(d_2582_v59_, globalState):
                (self).f12 = (self).f11
                d_2586_v61_: _dafny.Seq
                d_2586_v61_ = _dafny.SeqWithoutIsStrInference([(0) - (-361), (len(_dafny.Map({d_2499_v0_: (self).f11}))) * (d_2583_i9_)])
                (globalState).f7 = (d_2586_v61_)[default__.safeIndex(375, len(d_2586_v61_))]
                d_2587_v62_: _dafny.Array
                nw410_ = _dafny.Array(_dafny.Seq({}), 14)
                d_2587_v62_ = nw410_
                d_2587_v62_ = d_2587_v62_
                d_2588_v63_: _dafny.Array
                def lambda119_(d_2589_v61_, d_2590_v0_):
                    def lambda120_(d_2591_i11_):
                        return (len(d_2589_v61_)) != (d_2590_v0_)

                    return lambda120_

                init66_ = lambda119_(d_2586_v61_, d_2499_v0_)
                nw411_ = _dafny.Array(None, 2)
                for i0_66_ in range(nw411_.length(0)):
                    nw411_[i0_66_] = init66_(i0_66_)
                d_2588_v63_ = nw411_
                index444_ = default__.safeIndex(527, (d_2588_v63_).length(0))
                (d_2588_v63_)[index444_] = not((875) > (d_2499_v0_))
                d_2592_v64_: _dafny.Seq
                d_2592_v64_ = _dafny.SeqWithoutIsStrInference([d_2588_v63_])
                index445_ = default__.safeIndex(527, (d_2588_v63_).length(0))
                (d_2588_v63_)[index445_] = (d_2588_v63_) not in (d_2592_v64_)
                d_2581_v58_ = (d_2582_v59_) < (233)
            elif True:
                d_2593_v65_: _dafny.Seq
                d_2593_v65_ = _dafny.SeqWithoutIsStrInference([d_2504_v3_, (_dafny.SeqWithoutIsStrInference([self.f12])).set(default__.safeIndex(d_2499_v0_, len(_dafny.SeqWithoutIsStrInference([self.f12]))), (self).f11), d_2504_v3_])
                d_2581_v58_ = ((d_2504_v3_) + (d_2504_v3_)) not in (d_2593_v65_)
                d_2594_v66_: _dafny.Array
                nw412_ = _dafny.Array(False, 19)
                d_2594_v66_ = nw412_
                index446_ = default__.safeIndex(286, (d_2594_v66_).length(0))
                (d_2594_v66_)[index446_] = (self).f10
                index447_ = default__.safeIndex(286, (d_2594_v66_).length(0))
                (d_2594_v66_)[index447_] = d_2581_v58_
                d_2595_v67_: _dafny.Map
                d_2595_v67_ = _dafny.Map({d_2578_v55_: (d_2577_v54_).set(default__.safeIndex(d_2583_i9_, len(d_2577_v54_)), (d_2577_v54_)[default__.safeIndex(d_2583_i9_, len(d_2577_v54_))])})
                d_2596_v68_: _dafny.Set
                d_2596_v68_ = _dafny.Set({d_2581_v58_, d_2581_v58_, False, (self).f11})
                d_2597_v69_: _dafny.Map
                d_2597_v69_ = _dafny.Map({d_2596_v68_: _dafny.SeqWithoutIsStrInference([d_2578_v55_ for d_2598_i16_ in range(default__.abs(806))])})
                d_2599_v70_: _dafny.Array
                nw413_ = _dafny.Array(None, 24)
                nw413_[int(0)] = ((d_2595_v67_)[d_2578_v55_] if (d_2578_v55_) in (d_2595_v67_) else d_2577_v54_)
                nw413_[int(1)] = d_2577_v54_
                nw413_[int(2)] = d_2577_v54_
                nw413_[int(3)] = (d_2577_v54_) + (d_2577_v54_)
                nw413_[int(4)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2600_i12_ in range(default__.abs(354))])
                nw413_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcg")) if (self).f10 else default__.fm25(d_2499_v0_, (d_2594_v66_)[default__.safeIndex(286, (d_2594_v66_).length(0))], d_2583_i9_, globalState))
                nw413_[int(6)] = d_2577_v54_
                nw413_[int(7)] = d_2577_v54_
                nw413_[int(8)] = d_2577_v54_
                nw413_[int(9)] = d_2577_v54_
                nw413_[int(10)] = _dafny.SeqWithoutIsStrInference([d_2578_v55_ for d_2601_i13_ in range(default__.abs(949))])
                nw413_[int(11)] = _dafny.SeqWithoutIsStrInference([d_2578_v55_ for d_2602_i14_ in range(default__.abs(502))])
                nw413_[int(12)] = (d_2577_v54_) + (d_2577_v54_)
                nw413_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whgxodd"))
                nw413_[int(14)] = d_2577_v54_
                nw413_[int(15)] = ((d_2577_v54_) + (d_2577_v54_)).set(default__.safeIndex(d_2499_v0_, len((d_2577_v54_) + (d_2577_v54_))), d_2578_v55_)
                nw413_[int(16)] = d_2577_v54_
                nw413_[int(17)] = d_2577_v54_
                nw413_[int(18)] = (d_2577_v54_) + (d_2577_v54_)
                nw413_[int(19)] = d_2577_v54_
                nw413_[int(20)] = _dafny.SeqWithoutIsStrInference([d_2578_v55_ for d_2603_i15_ in range(default__.abs(855))])
                nw413_[int(21)] = ((d_2597_v69_)[d_2596_v68_] if (d_2596_v68_) in (d_2597_v69_) else d_2577_v54_)
                nw413_[int(22)] = (d_2577_v54_) + (d_2577_v54_)
                nw413_[int(23)] = d_2577_v54_
                d_2599_v70_ = nw413_
                d_2604_v71_: _dafny.Array
                nw414_ = _dafny.Array(_dafny.Map({}), 29)
                d_2604_v71_ = nw414_
                d_2605_v73_: _dafny.MultiSet
                d_2605_v73_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofs")))])
                d_2606_v74_: _dafny.Seq
                d_2606_v74_ = _dafny.SeqWithoutIsStrInference([d_2605_v73_])
                d_2607_v75_: _dafny.Seq
                d_2607_v75_ = _dafny.SeqWithoutIsStrInference([d_2606_v74_])
                index448_ = default__.safeIndex(386, (d_2604_v71_).length(0))
                def iife192_():
                    coll64_ = _dafny.Map()
                    compr_64_: _dafny.MultiSet
                    for compr_64_ in ((d_2607_v75_)[default__.safeIndex(default__.fm1(globalState), len(d_2607_v75_))]).Elements:
                        d_2608_v72_: _dafny.MultiSet = compr_64_
                        if (d_2608_v72_) in ((d_2607_v75_)[default__.safeIndex(default__.fm1(globalState), len(d_2607_v75_))]):
                            coll64_[d_2608_v72_] = _dafny.Map({d_2503_v2_: D5_DC18(D5_DC18(D5_DC17((d_2594_v66_)[default__.safeIndex(286, (d_2594_v66_).length(0))], d_2578_v55_)))})
                    return _dafny.Map(coll64_)
                (d_2604_v71_)[index448_] = iife192_()
                
                index449_ = default__.safeIndex(661, (d_2500_v1_).length(0))
                (d_2500_v1_)[index449_] = d_2499_v0_
                index450_ = default__.safeIndex(963, (d_2500_v1_).length(0))
                (d_2500_v1_)[index450_] = len(d_2596_v68_)
                d_2609_v76_: _dafny.Map
                d_2609_v76_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2581_v58_, (self).f11]): d_2583_i9_})
                d_2610_v77_: _dafny.Map
                d_2610_v77_ = _dafny.Map({not((self).f10): d_2583_i9_})
                d_2611_v78_: _dafny.MultiSet
                d_2611_v78_ = _dafny.MultiSet([d_2581_v58_])
                d_2612_v79_: D5
                d_2612_v79_ = D5_DC16((d_2611_v78_).set((d_2594_v66_)[default__.safeIndex(286, (d_2594_v66_).length(0))], default__.abs(d_2582_v59_)))
                d_2613_v80_: D5
                d_2613_v80_ = D5_DC18(d_2612_v79_)
                d_2614_v81_: D5
                d_2614_v81_ = D5_DC18(d_2613_v80_)
                d_2615_v82_: _dafny.Map
                d_2615_v82_ = _dafny.Map({D10_DC31(d_2581_v58_, d_2581_v58_, d_2582_v59_): d_2614_v81_})
                d_2616_v83_: _dafny.Map
                d_2616_v83_ = _dafny.Map({default__.fm18(d_2583_i9_, d_2610_v77_, (d_2594_v66_)[default__.safeIndex(286, (d_2594_v66_).length(0))], globalState): d_2615_v82_})
                d_2617_v84_: _dafny.Set
                d_2617_v84_ = _dafny.Set({default__.fm58(globalState)})
                d_2618_v85_: D19
                d_2618_v85_ = D19_DC58((d_2605_v73_).cardinality, (d_2594_v66_)[default__.safeIndex(286, (d_2594_v66_).length(0))])
                d_2619_v86_: _dafny.Seq
                d_2619_v86_ = _dafny.SeqWithoutIsStrInference([d_2617_v84_, _dafny.Set({d_2618_v85_, d_2618_v85_})])
                index451_ = default__.safeIndex(386, (d_2604_v71_).length(0))
                index452_ = default__.safeIndex(661, (d_2500_v1_).length(0))
                index453_ = default__.safeIndex(963, (d_2500_v1_).length(0))
                rhs427_ = (d_2504_v3_) not in (d_2609_v76_)
                rhs428_ = d_2599_v70_
                rhs429_ = d_2616_v83_
                rhs430_ = len(d_2619_v86_)
                rhs431_ = ((0) - (d_2583_i9_)) + ((d_2499_v0_) * (d_2582_v59_))
                lhs312_ = d_2604_v71_
                lhs313_ = default__.safeIndex(386, (d_2604_v71_).length(0))
                lhs314_ = d_2500_v1_
                lhs315_ = default__.safeIndex(661, (d_2500_v1_).length(0))
                lhs316_ = d_2500_v1_
                lhs317_ = default__.safeIndex(963, (d_2500_v1_).length(0))
                d_2581_v58_ = rhs427_
                d_2599_v70_ = rhs428_
                lhs312_[lhs313_] = rhs429_
                lhs314_[lhs315_] = rhs430_
                lhs316_[lhs317_] = rhs431_
                d_2610_v77_ = (d_2610_v77_).set(self.f12, d_2582_v59_)
                d_2620_v87_: _dafny.Seq
                d_2620_v87_ = _dafny.SeqWithoutIsStrInference([d_2594_v66_, d_2594_v66_, d_2594_v66_])
                d_2621_v88_: _dafny.Array
                nw415_ = _dafny.Array(None, 20)
                nw415_[int(0)] = d_2578_v55_
                nw415_[int(1)] = d_2578_v55_
                nw415_[int(2)] = d_2578_v55_
                nw415_[int(3)] = _dafny.CodePoint('r')
                nw415_[int(4)] = d_2578_v55_
                nw415_[int(5)] = d_2578_v55_
                nw415_[int(6)] = d_2578_v55_
                nw415_[int(7)] = d_2578_v55_
                nw415_[int(8)] = d_2578_v55_
                nw415_[int(9)] = default__.fm4(self.f12, (d_2594_v66_)[default__.safeIndex(286, (d_2594_v66_).length(0))], False, globalState)
                nw415_[int(10)] = d_2578_v55_
                nw415_[int(11)] = d_2578_v55_
                nw415_[int(12)] = (d_2578_v55_ if d_2581_v58_ else d_2578_v55_)
                nw415_[int(13)] = d_2578_v55_
                nw415_[int(14)] = d_2578_v55_
                nw415_[int(15)] = d_2578_v55_
                nw415_[int(16)] = d_2578_v55_
                nw415_[int(17)] = d_2578_v55_
                nw415_[int(18)] = d_2578_v55_
                nw415_[int(19)] = d_2578_v55_
                d_2621_v88_ = nw415_
                index454_ = default__.safeIndex(970, (d_2621_v88_).length(0))
                (d_2621_v88_)[index454_] = d_2578_v55_
                index455_ = default__.safeIndex(970, (d_2621_v88_).length(0))
                rhs432_ = d_2620_v87_
                rhs433_ = d_2581_v58_
                rhs434_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gim"))) + (default__.fm13((self).f11, (self).f10, globalState))) + (d_2577_v54_)
                rhs435_ = d_2578_v55_
                lhs318_ = d_2621_v88_
                lhs319_ = default__.safeIndex(970, (d_2621_v88_).length(0))
                d_2620_v87_ = rhs432_
                r1 = rhs433_
                d_2577_v54_ = rhs434_
                lhs318_[lhs319_] = rhs435_
            d_2622_v89_: C0
            nw416_ = C0()
            nw416_.ctor__()
            d_2622_v89_ = nw416_
        pat_let_tv70_ = d_2499_v0_
        pat_let_tv71_ = d_2499_v0_
        pat_let_tv72_ = d_2499_v0_
        def lambda121_(source42_):
            if source42_.is_DC31:
                d_2623___mcc_h7_ = source42_.cf50
                d_2624___mcc_h8_ = source42_.cf51
                d_2625___mcc_h9_ = source42_.cf52
                d_2626_cf52_ = d_2625___mcc_h9_
                d_2627_cf51_ = d_2624___mcc_h8_
                d_2628_cf50_ = d_2623___mcc_h7_
                return (0) - ((886) + (len(_dafny.Set({353, pat_let_tv70_, pat_let_tv71_, pat_let_tv72_}))))
            elif True:
                d_2629___mcc_h10_ = source42_.cf49
                d_2630_cf49_ = d_2629___mcc_h10_
                return 154

        def iife193_(_pat_let64_0):
            def iife194_(d_2631_dt__update__tmp_h0_):
                def iife195_(_pat_let65_0):
                    def iife196_(d_2632_dt__update_hcf50_h0_):
                        return D10_DC31(d_2632_dt__update_hcf50_h0_, (d_2631_dt__update__tmp_h0_).cf51, (d_2631_dt__update__tmp_h0_).cf52)
                    return iife196_(_pat_let65_0)
                return iife195_(self.f12)
            return iife194_(_pat_let64_0)
        (globalState).f0 = lambda121_(iife193_(d_2503_v2_))
        d_2633_v90_: T0
        nw417_ = C1()
        nw417_.ctor__(d_2581_v58_)
        d_2633_v90_ = nw417_
        d_2634_v91_: _dafny.MultiSet
        d_2634_v91_ = _dafny.MultiSet([(self).f10])
        rhs436_ = d_2633_v90_
        rhs437_ = (d_2634_v91_).issubset(d_2634_v91_)
        d_2633_v90_ = rhs436_
        d_2581_v58_ = rhs437_
        d_2635_v92_: _dafny.Set
        d_2635_v92_ = _dafny.Set({d_2577_v54_})
        r0 = ((d_2635_v92_) | (d_2635_v92_)) - ((d_2635_v92_) - (d_2635_v92_))
        d_2636_v93_: _dafny.MultiSet
        d_2636_v93_ = _dafny.MultiSet([len(d_2504_v3_)])
        d_2637_v94_: _dafny.Array
        nw418_ = _dafny.Array(None, 2)
        nw418_[int(0)] = not(self.f12)
        nw418_[int(1)] = d_2581_v58_
        d_2637_v94_ = nw418_
        d_2638_v95_: _dafny.Map
        d_2638_v95_ = _dafny.Map({len(d_2504_v3_): d_2637_v94_})
        r1 = not(default__.fm0(default__.safeDivisionInt(((d_2636_v93_)[len(d_2638_v95_)] if (len(d_2638_v95_)) in (d_2636_v93_) else len(d_2504_v3_)), d_2582_v59_), globalState))
        r2 = default__.fm1(globalState)
        return r0, r1, r2

    def m5(self, p0, globalState):
        if (self).f10:
            (globalState).f0 = p0
            d_2639_v0_: _dafny.Array
            nw419_ = _dafny.Array(int(0), 5)
            d_2639_v0_ = nw419_
            index456_ = default__.safeIndex(351, (d_2639_v0_).length(0))
            (d_2639_v0_)[index456_] = (p0) - (p0)
            d_2640_v1_: _dafny.Set
            d_2640_v1_ = _dafny.Set({self.f12, (self).f11})
            d_2641_v2_: D1
            d_2641_v2_ = D1_DC3(p0)
            d_2642_v3_: _dafny.Set
            d_2642_v3_ = _dafny.Set({p0, (d_2641_v2_).cf4})
            index457_ = default__.safeIndex(351, (d_2639_v0_).length(0))
            (d_2639_v0_)[index457_] = (len(d_2640_v1_)) * (len((d_2642_v3_).intersection(default__.fm31(p0, globalState))))
            d_2643_v4_: D6
            d_2643_v4_ = D6_DC20((self).f11, self.f12, (self).f11, default__.safeDivisionInt(p0, (d_2639_v0_)[default__.safeIndex(351, (d_2639_v0_).length(0))]))
            d_2644_v5_: str
            d_2644_v5_ = _dafny.CodePoint('a')
            d_2645_v6_: _dafny.MultiSet
            d_2645_v6_ = _dafny.MultiSet([p0])
            pat_let_tv73_ = p0
            index458_ = default__.safeIndex(351, (d_2639_v0_).length(0))
            def iife197_(_pat_let66_0):
                def iife198_(d_2646_dt__update__tmp_h0_):
                    def iife199_(_pat_let67_0):
                        def iife200_(d_2647_dt__update_hcf32_h0_):
                            def iife201_(_pat_let68_0):
                                def iife202_(d_2648_dt__update_hcf35_h0_):
                                    def iife203_(_pat_let69_0):
                                        def iife204_(d_2649_dt__update_hcf34_h0_):
                                            return D6_DC20(d_2647_dt__update_hcf32_h0_, (d_2646_dt__update__tmp_h0_).cf33, d_2649_dt__update_hcf34_h0_, d_2648_dt__update_hcf35_h0_)
                                        return iife204_(_pat_let69_0)
                                    return iife203_(True)
                                return iife202_(_pat_let68_0)
                            return iife201_(pat_let_tv73_)
                        return iife200_(_pat_let67_0)
                    return iife199_(self.f12)
                return iife198_(_pat_let66_0)
            rhs438_ = p0
            rhs439_ = iife197_(D6_DC20((self).f11, (self).f11, (D5_DC17(self.f12, d_2644_v5_)).cf28, (d_2645_v6_).cardinality))
            rhs440_ = self.f12
            rhs441_ = (((d_2639_v0_)[default__.safeIndex(351, (d_2639_v0_).length(0))] if self.f12 else (d_2639_v0_)[default__.safeIndex(351, (d_2639_v0_).length(0))])) - ((d_2639_v0_)[default__.safeIndex(351, (d_2639_v0_).length(0))])
            lhs320_ = globalState
            lhs321_ = self
            lhs322_ = d_2639_v0_
            lhs323_ = default__.safeIndex(351, (d_2639_v0_).length(0))
            lhs320_.f7 = rhs438_
            d_2643_v4_ = rhs439_
            lhs321_.f12 = rhs440_
            lhs322_[lhs323_] = rhs441_
            d_2650_v7_: _dafny.Array
            nw420_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_2650_v7_ = nw420_
            d_2651_v8_: _dafny.Array
            nw421_ = _dafny.Array(False, 3)
            d_2651_v8_ = nw421_
            index459_ = default__.safeIndex(93, (d_2650_v7_).length(0))
            (d_2650_v7_)[index459_] = d_2651_v8_
            index460_ = default__.safeIndex(93, (d_2650_v7_).length(0))
            (d_2650_v7_)[index460_] = d_2651_v8_
            d_2652_v9_: C10
            nw422_ = C10()
            nw422_.ctor__(((d_2639_v0_)[default__.safeIndex(351, (d_2639_v0_).length(0))]) * (p0), self.f12)
            d_2652_v9_ = nw422_
        elif True:
            d_2653_v10_: _dafny.Seq
            d_2653_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbgo"))
            d_2654_v11_: C3
            nw423_ = C3()
            nw423_.ctor__(False, self.f12)
            d_2654_v11_ = nw423_
            d_2655_v12_: _dafny.Map
            d_2655_v12_ = _dafny.Map({d_2654_v11_: p0})
            d_2656_v13_: _dafny.Array
            nw424_ = _dafny.Array(None, 18)
            nw424_[int(0)] = 965
            nw424_[int(1)] = p0
            nw424_[int(2)] = len(d_2653_v10_)
            nw424_[int(3)] = default__.fm1(globalState)
            nw424_[int(4)] = p0
            nw424_[int(5)] = len(d_2653_v10_)
            nw424_[int(6)] = (0) - (p0)
            nw424_[int(7)] = p0
            nw424_[int(8)] = p0
            nw424_[int(9)] = p0
            nw424_[int(10)] = 768
            nw424_[int(11)] = p0
            nw424_[int(12)] = p0
            nw424_[int(13)] = len(d_2653_v10_)
            nw424_[int(14)] = len(d_2655_v12_)
            nw424_[int(15)] = p0
            nw424_[int(16)] = -624
            nw424_[int(17)] = p0
            d_2656_v13_ = nw424_
            d_2657_v14_: _dafny.Seq
            d_2657_v14_ = _dafny.SeqWithoutIsStrInference([d_2656_v13_, d_2656_v13_])
            d_2658_v15_: _dafny.Map
            d_2658_v15_ = _dafny.Map({(d_2657_v14_)[default__.safeIndex(p0, len(d_2657_v14_))]: default__.fm0(p0, globalState)})
            d_2658_v15_ = (d_2658_v15_).set(d_2656_v13_, (self).f11)
            (globalState).f1 = (default__.safeDivisionInt(-29, p0) if (d_2654_v11_).f20 else p0)
            (self).f12 = not ((p0) > (p0)) or (not((d_2654_v11_).f20))
            d_2659_v16_: _dafny.Array
            def lambda122_(d_2660_i0_):
                return self.f12

            init67_ = lambda122_
            nw425_ = _dafny.Array(None, 10)
            for i0_67_ in range(nw425_.length(0)):
                nw425_[i0_67_] = init67_(i0_67_)
            d_2659_v16_ = nw425_
            d_2661_v17_: _dafny.MultiSet
            d_2661_v17_ = _dafny.MultiSet([d_2659_v16_])
            rhs442_ = d_2661_v17_
            rhs443_ = p0
            rhs444_ = (self).f10
            lhs324_ = globalState
            lhs325_ = self
            d_2661_v17_ = rhs442_
            lhs324_.f0 = rhs443_
            lhs325_.f12 = rhs444_
            d_2662_v18_: _dafny.Array
            def lambda123_(d_2663_p0_, d_2664_v10_):
                def lambda124_(d_2665_i1_):
                    return _dafny.MultiSet([(0) - (d_2663_p0_), len(_dafny.Map({(self).f10: len(d_2664_v10_)}))])

                return lambda124_

            init68_ = lambda123_(p0, d_2653_v10_)
            nw426_ = _dafny.Array(None, 21)
            for i0_68_ in range(nw426_.length(0)):
                nw426_[i0_68_] = init68_(i0_68_)
            d_2662_v18_ = nw426_
            d_2666_v19_: _dafny.Map
            d_2666_v19_ = _dafny.Map({(d_2662_v18_ if (d_2654_v11_).f20 else d_2662_v18_): d_2656_v13_})
            d_2666_v19_ = (d_2666_v19_).set(d_2662_v18_, d_2656_v13_)
        d_2667_v20_: _dafny.Seq
        d_2667_v20_ = _dafny.SeqWithoutIsStrInference([(self).f11, self.f12])
        d_2668_v21_: _dafny.Map
        d_2668_v21_ = _dafny.Map({d_2667_v20_: p0})
        hi18_ = len(d_2668_v21_)
        for d_2669_i2_ in range(p0, hi18_):
            d_2670_v22_: _dafny.Array
            def lambda125_(d_2671_p0_):
                def lambda126_(d_2672_i3_):
                    return (d_2672_i3_) + (d_2671_p0_)

                return lambda126_

            init69_ = lambda125_(p0)
            nw427_ = _dafny.Array(None, 20)
            for i0_69_ in range(nw427_.length(0)):
                nw427_[i0_69_] = init69_(i0_69_)
            d_2670_v22_ = nw427_
            d_2673_v23_: _dafny.MultiSet
            d_2673_v23_ = _dafny.MultiSet([d_2669_i2_])
            d_2674_v24_: _dafny.Map
            d_2674_v24_ = _dafny.Map({(self).f11: d_2673_v23_})
            index461_ = default__.safeIndex(994, (d_2670_v22_).length(0))
            (d_2670_v22_)[index461_] = len(d_2674_v24_)
            index462_ = default__.safeIndex(994, (d_2670_v22_).length(0))
            (d_2670_v22_)[index462_] = d_2669_i2_
            d_2675_v25_: _dafny.Array
            def lambda127_(d_2676_i4_):
                return (self).f10

            init70_ = lambda127_
            nw428_ = _dafny.Array(None, 10)
            for i0_70_ in range(nw428_.length(0)):
                nw428_[i0_70_] = init70_(i0_70_)
            d_2675_v25_ = nw428_
            rhs445_ = (default__.safeModuloInt(d_2669_i2_, p0)) - (p0)
            rhs446_ = d_2675_v25_
            lhs326_ = globalState
            lhs326_.f0 = rhs445_
            d_2675_v25_ = rhs446_
            d_2677_v26_: C0
            nw429_ = C0()
            nw429_.ctor__()
            d_2677_v26_ = nw429_
            d_2678_v27_: C1
            nw430_ = C1()
            nw430_.ctor__(not (self.f12) or ((self).f10))
            d_2678_v27_ = nw430_
        if self.f12:
            d_2679_v28_: _dafny.Seq
            d_2679_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgpnm"))
            d_2680_v29_: _dafny.Map
            d_2680_v29_ = _dafny.Map({D3_DC10(d_2679_v28_, (self).f11, (self).f10): p0})
            d_2681_v30_: D3
            d_2681_v30_ = D3_DC10(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlowkigp")), (self).f11, True)
            d_2680_v29_ = (d_2680_v29_).set(d_2681_v30_, len(d_2679_v28_))
            d_2679_v28_ = d_2679_v28_
            d_2682_v31_: _dafny.Array
            def lambda128_(d_2683_i5_):
                return (self).f11

            init71_ = lambda128_
            nw431_ = _dafny.Array(None, 29)
            for i0_71_ in range(nw431_.length(0)):
                nw431_[i0_71_] = init71_(i0_71_)
            d_2682_v31_ = nw431_
            index463_ = default__.safeIndex(575, (d_2682_v31_).length(0))
            (d_2682_v31_)[index463_] = self.f12
            index464_ = default__.safeIndex(575, (d_2682_v31_).length(0))
            (d_2682_v31_)[index464_] = self.f12
            d_2684_v32_: _dafny.Map
            d_2684_v32_ = _dafny.Map({self.f12: default__.safeModuloInt(p0, len(d_2679_v28_))})
            d_2684_v32_ = d_2684_v32_
            index465_ = default__.safeIndex(575, (d_2682_v31_).length(0))
            (d_2682_v31_)[index465_] = (p0) < (p0)
        elif True:
            d_2685_v33_: _dafny.Map
            d_2685_v33_ = _dafny.Map({-724: p0})
            d_2686_v34_: _dafny.Seq
            d_2686_v34_ = _dafny.SeqWithoutIsStrInference([955])
            d_2687_v35_: _dafny.Set
            d_2687_v35_ = _dafny.Set({(self).f11, (self).f11})
            d_2688_v36_: _dafny.MultiSet
            d_2688_v36_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([541 for d_2689_i7_ in range(default__.abs(419))])])
            d_2690_v37_: str
            d_2690_v37_ = _dafny.CodePoint('m')
            d_2691_v38_: _dafny.Seq
            d_2691_v38_ = _dafny.SeqWithoutIsStrInference([d_2690_v37_, d_2690_v37_])
            d_2692_v39_: _dafny.Array
            nw432_ = _dafny.Array(None, 19)
            nw432_[int(0)] = 648
            nw432_[int(1)] = len(_dafny.Map({False: p0}))
            nw432_[int(2)] = p0
            nw432_[int(3)] = p0
            nw432_[int(4)] = (0) - (p0)
            nw432_[int(5)] = p0
            nw432_[int(6)] = p0
            nw432_[int(7)] = p0
            nw432_[int(8)] = p0
            nw432_[int(9)] = p0
            nw432_[int(10)] = len(d_2685_v33_)
            nw432_[int(11)] = p0
            nw432_[int(12)] = 208
            nw432_[int(13)] = ((d_2685_v33_)[len(default__.fm45((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, p0]) for d_2693_i6_ in range(default__.abs(-165))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, p0]) for d_2694_i6_ in range(default__.abs(-165))]))), d_2686_v34_), p0, globalState))] if (len(default__.fm45((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, p0]) for d_2695_i6_ in range(default__.abs(-165))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, p0]) for d_2696_i6_ in range(default__.abs(-165))]))), d_2686_v34_), p0, globalState))) in (d_2685_v33_) else 541)
            nw432_[int(14)] = p0
            nw432_[int(15)] = p0
            nw432_[int(16)] = p0
            nw432_[int(17)] = default__.safeDivisionInt(len(d_2687_v35_), p0)
            nw432_[int(18)] = ((d_2688_v36_).cardinality) + (len(d_2691_v38_))
            d_2692_v39_ = nw432_
            index466_ = default__.safeIndex(2, (d_2692_v39_).length(0))
            (d_2692_v39_)[index466_] = 571
            index467_ = default__.safeIndex(2, (d_2692_v39_).length(0))
            (d_2692_v39_)[index467_] = p0
            (self).f12 = False
            (self).f12 = (self).f11
            d_2697_v40_: C3
            nw433_ = C3()
            nw433_.ctor__(self.f12, self.f12)
            d_2697_v40_ = nw433_
            d_2698_v41_: _dafny.Array
            nw434_ = _dafny.Array(False, 2)
            d_2698_v41_ = nw434_
            index468_ = default__.safeIndex(348, (d_2698_v41_).length(0))
            (d_2698_v41_)[index468_] = (d_2697_v40_).f20
            d_2699_v42_: D17
            d_2699_v42_ = D17_DC51(d_2690_v37_, (d_2697_v40_).f20)
            d_2700_v43_: _dafny.MultiSet
            d_2700_v43_ = _dafny.MultiSet([d_2699_v42_, d_2699_v42_, d_2699_v42_])
            index469_ = default__.safeIndex(348, (d_2698_v41_).length(0))
            (d_2698_v41_)[index469_] = ((d_2700_v43_) | (d_2700_v43_)).issubset(d_2700_v43_)
        d_2701_v44_: D15
        d_2701_v44_ = D15_DC46()
        source43_ = d_2701_v44_
        if source43_.is_DC45:
            d_2702___mcc_h0_ = source43_.cf80
            d_2703___mcc_h1_ = source43_.cf81
            d_2704___mcc_h2_ = source43_.cf82
            d_2705_cf82_ = d_2704___mcc_h2_
            d_2706_cf81_ = d_2703___mcc_h1_
            d_2707_cf80_ = d_2702___mcc_h0_
            d_2705_cf82_ = d_2706_cf81_
            d_2708_v45_: _dafny.Map
            d_2708_v45_ = _dafny.Map({d_2706_cf81_: d_2707_cf80_})
            d_2709_v46_: C7
            nw435_ = C7()
            nw435_.ctor__(self.f12, (False) not in (d_2708_v45_))
            d_2709_v46_ = nw435_
            (globalState).f1 = 746
            d_2706_cf81_ = self.f12
        elif source43_.is_DC46:
            d_2710_v47_: _dafny.Array
            nw436_ = _dafny.Array(False, 9)
            d_2710_v47_ = nw436_
            d_2711_v48_: _dafny.Seq
            d_2712_v49_: _dafny.Seq
            out72_: _dafny.Seq
            out73_: _dafny.Seq
            out72_, out73_ = default__.m0(p0, d_2710_v47_, globalState)
            d_2711_v48_ = out72_
            d_2712_v49_ = out73_
            if ((self).f10 if (self).f10 else (d_2667_v20_)[default__.safeIndex(390, len(d_2667_v20_))]):
                d_2713_v50_: C2
                nw437_ = C2()
                nw437_.ctor__((self).f11, p0)
                d_2713_v50_ = nw437_
                d_2714_v51_: _dafny.Map
                d_2714_v51_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])): d_2713_v50_})
                d_2715_v52_: _dafny.Map
                d_2715_v52_ = _dafny.Map({p0: (d_2713_v50_).f22})
                d_2714_v51_ = (d_2714_v51_).set(len(d_2715_v52_), d_2713_v50_)
                d_2716_v53_: C3
                nw438_ = C3()
                nw438_.ctor__(True, (self).fm10(globalState))
                d_2716_v53_ = nw438_
                d_2717_v54_: _dafny.Array
                nw439_ = _dafny.Array(None, 13)
                nw439_[int(0)] = d_2716_v53_
                nw439_[int(1)] = d_2716_v53_
                nw439_[int(2)] = d_2716_v53_
                nw439_[int(3)] = d_2716_v53_
                nw439_[int(4)] = d_2716_v53_
                nw439_[int(5)] = (d_2716_v53_ if self.f12 else d_2716_v53_)
                nw439_[int(6)] = d_2716_v53_
                nw439_[int(7)] = d_2716_v53_
                nw439_[int(8)] = d_2716_v53_
                nw439_[int(9)] = d_2716_v53_
                nw439_[int(10)] = d_2716_v53_
                nw439_[int(11)] = d_2716_v53_
                nw439_[int(12)] = d_2716_v53_
                d_2717_v54_ = nw439_
                index470_ = default__.safeIndex(31, (d_2717_v54_).length(0))
                (d_2717_v54_)[index470_] = d_2716_v53_
                d_2718_v55_: D21
                d_2718_v55_ = D21_DC61(d_2716_v53_)
                index471_ = default__.safeIndex(31, (d_2717_v54_).length(0))
                (d_2717_v54_)[index471_] = ((d_2716_v53_ if (d_2713_v50_).f21 else (d_2718_v55_).cf103) if (423) != ((0) - ((d_2713_v50_).f22)) else d_2716_v53_)
                index472_ = default__.safeIndex(295, (d_2710_v47_).length(0))
                (d_2710_v47_)[index472_] = (self).f10
                index473_ = default__.safeIndex(295, (d_2710_v47_).length(0))
                (d_2710_v47_)[index473_] = ((d_2667_v20_)[default__.safeIndex(p0, len(d_2667_v20_))] if (d_2713_v50_).f21 else (d_2713_v50_).f21)
                d_2719_v56_: C9
                nw440_ = C9()
                nw440_.ctor__((self).f10, (self).f11)
                d_2719_v56_ = nw440_
                d_2720_v57_: _dafny.Map
                d_2720_v57_ = _dafny.Map({d_2711_v48_: (d_2719_v56_).f17})
                d_2720_v57_ = (d_2720_v57_).set(d_2712_v49_, (d_2719_v56_).f17)
            elif True:
                (globalState).f1 = default__.safeModuloInt(p0, p0)
                d_2721_v58_: C2
                nw441_ = C2()
                nw441_.ctor__(not((d_2667_v20_)[default__.safeIndex(p0, len(d_2667_v20_))]), p0)
                d_2721_v58_ = nw441_
                d_2722_v59_: _dafny.Array
                nw442_ = _dafny.Array(int(0), 5)
                d_2722_v59_ = nw442_
                nw443_ = _dafny.Array(int(0), 8)
                d_2722_v59_ = nw443_
                (globalState).f0 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qe")))
                d_2723_v60_: C9
                nw444_ = C9()
                nw444_.ctor__((d_2667_v20_)[default__.safeIndex((0) - ((d_2721_v58_).f22), len(d_2667_v20_))], (self).f10)
                d_2723_v60_ = nw444_
            d_2724_v61_: str
            d_2724_v61_ = _dafny.CodePoint('l')
            d_2725_v62_: D2
            d_2725_v62_ = D2_DC6(d_2724_v61_)
            d_2726_v63_: _dafny.Array
            nw445_ = _dafny.Array(None, 17)
            nw445_[int(0)] = d_2724_v61_
            nw445_[int(1)] = d_2724_v61_
            nw445_[int(2)] = _dafny.CodePoint('p')
            nw445_[int(3)] = default__.fm4(self.f12, (self).f10, self.f12, globalState)
            nw445_[int(4)] = d_2724_v61_
            nw445_[int(5)] = (d_2712_v49_)[default__.safeIndex((0) - (p0), len(d_2712_v49_))]
            nw445_[int(6)] = d_2724_v61_
            nw445_[int(7)] = d_2724_v61_
            nw445_[int(8)] = d_2724_v61_
            nw445_[int(9)] = d_2724_v61_
            nw445_[int(10)] = d_2724_v61_
            nw445_[int(11)] = d_2724_v61_
            nw445_[int(12)] = _dafny.CodePoint('g')
            nw445_[int(13)] = (d_2725_v62_).cf8
            nw445_[int(14)] = d_2724_v61_
            nw445_[int(15)] = d_2724_v61_
            nw445_[int(16)] = default__.fm4((self).f10, (self).f11, (self).f11, globalState)
            d_2726_v63_ = nw445_
            index474_ = default__.safeIndex(226, (d_2726_v63_).length(0))
            (d_2726_v63_)[index474_] = d_2724_v61_
            index475_ = default__.safeIndex(226, (d_2726_v63_).length(0))
            (d_2726_v63_)[index475_] = d_2724_v61_
            index476_ = default__.safeIndex(427, (d_2710_v47_).length(0))
            (d_2710_v47_)[index476_] = (self).f10
            index477_ = default__.safeIndex(427, (d_2710_v47_).length(0))
            (d_2710_v47_)[index477_] = not(not(True))
        elif source43_.is_DC44:
            d_2727___mcc_h3_ = source43_.cf79
            d_2728_cf79_ = d_2727___mcc_h3_
            d_2729_v64_: _dafny.Array
            nw446_ = _dafny.Array(int(0), 28)
            d_2729_v64_ = nw446_
            d_2730_v65_: D21
            d_2730_v65_ = D21_DC63(p0, (self).f10, (self).f10, (self).f11)
            rhs447_ = ((d_2730_v65_).cf108) == ((self).f11)
            rhs448_ = d_2729_v64_
            lhs327_ = self
            lhs327_.f12 = rhs447_
            d_2729_v64_ = rhs448_
            d_2731_v66_: C5
            nw447_ = C5()
            nw447_.ctor__()
            d_2731_v66_ = nw447_
            d_2732_v67_: _dafny.Seq
            d_2732_v67_ = _dafny.SeqWithoutIsStrInference([(0) - (p0)])
            d_2733_v68_: _dafny.Map
            d_2733_v68_ = _dafny.Map({False: self.f12})
            d_2734_v69_: _dafny.Map
            d_2734_v69_ = _dafny.Map({d_2733_v68_: p0})
            d_2735_v70_: str
            d_2735_v70_ = _dafny.CodePoint('i')
            d_2736_v71_: _dafny.Seq
            d_2736_v71_ = _dafny.SeqWithoutIsStrInference([d_2735_v70_])
            d_2737_v72_: _dafny.MultiSet
            d_2737_v72_ = _dafny.MultiSet([((d_2734_v69_)[d_2733_v68_] if (d_2733_v68_) in (d_2734_v69_) else p0), p0, p0, len(d_2736_v71_), p0])
            d_2738_v73_: D4
            d_2738_v73_ = D4_DC15((default__.fm1(globalState) if not((self).f11) else len(d_2732_v67_)), (d_2737_v72_).set(p0, default__.abs(p0)), (d_2736_v71_)[default__.safeIndex(p0, len(d_2736_v71_))], (0) - (p0))
            source44_ = d_2738_v73_
            if source44_.is_DC14:
                d_2739___mcc_h5_ = source44_.cf21
                d_2740___mcc_h6_ = source44_.cf22
                d_2741_cf22_ = d_2740___mcc_h6_
                d_2742_cf21_ = d_2739___mcc_h5_
                d_2743_v74_: C2
                nw448_ = C2()
                nw448_.ctor__(self.f12, len(d_2733_v68_))
                d_2743_v74_ = nw448_
                d_2744_v75_: _dafny.Map
                d_2744_v75_ = _dafny.Map({d_2741_cf22_: not (self.f12) or (self.f12)})
                rhs449_ = d_2743_v74_
                rhs450_ = d_2736_v71_
                rhs451_ = d_2744_v75_
                rhs452_ = _dafny.SeqWithoutIsStrInference([default__.fm0(len(d_2736_v71_), globalState), (d_2743_v74_).f21, not ((d_2743_v74_).f21) or (not(self.f12))])
                rhs453_ = ((d_2732_v67_ if not((d_2743_v74_).f21) else d_2732_v67_)) + (_dafny.SeqWithoutIsStrInference([(d_2743_v74_).f22 for d_2745_i8_ in range(default__.abs(-929))]))
                d_2743_v74_ = rhs449_
                d_2736_v71_ = rhs450_
                d_2744_v75_ = rhs451_
                d_2667_v20_ = rhs452_
                d_2732_v67_ = rhs453_
                d_2746_v76_: _dafny.Set
                d_2746_v76_ = _dafny.Set({d_2733_v68_, d_2733_v68_})
                d_2733_v68_ = (d_2733_v68_).set((self).fm10(globalState), (_dafny.MultiSet([len(d_2667_v20_)])).issubset((_dafny.MultiSet(default__.fm3(d_2742_cf21_, 469, globalState))).set((0) - ((d_2743_v74_).fm9(d_2736_v71_, d_2746_v76_, globalState)), default__.abs((self).fm9(d_2736_v71_, d_2746_v76_, globalState)))))
                d_2747_v77_: _dafny.Map
                d_2747_v77_ = _dafny.Map({p0: d_2667_v20_})
                d_2747_v77_ = (d_2747_v77_).set(d_2741_cf22_, ((d_2747_v77_)[default__.fm1(globalState)] if (default__.fm1(globalState)) in (d_2747_v77_) else default__.fm40(not(self.f12), (self).f10, globalState)))
                d_2748_v78_: D4
                d_2748_v78_ = D4_DC14(d_2741_cf22_, d_2741_cf22_)
                (globalState).f7 = (d_2748_v78_).cf22
            elif source44_.is_DC15:
                d_2749___mcc_h7_ = source44_.cf23
                d_2750___mcc_h8_ = source44_.cf24
                d_2751___mcc_h9_ = source44_.cf25
                d_2752___mcc_h10_ = source44_.cf26
                d_2753_cf26_ = d_2752___mcc_h10_
                d_2754_cf25_ = d_2751___mcc_h9_
                d_2755_cf24_ = d_2750___mcc_h8_
                d_2756_cf23_ = d_2749___mcc_h7_
                d_2757_v79_: C5
                nw449_ = C5()
                nw449_.ctor__()
                d_2757_v79_ = nw449_
                (globalState).f1 = d_2756_cf23_
                (globalState).f7 = d_2753_cf26_
                d_2758_v80_: _dafny.Array
                nw450_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_2758_v80_ = nw450_
                index478_ = default__.safeIndex(154, (d_2758_v80_).length(0))
                (d_2758_v80_)[index478_] = d_2729_v64_
                index479_ = default__.safeIndex(154, (d_2758_v80_).length(0))
                (d_2758_v80_)[index479_] = d_2729_v64_
            elif True:
                d_2759___mcc_h11_ = source44_.cf20
                d_2760_cf20_ = d_2759___mcc_h11_
                (globalState).f1 = (0) - (p0)
                d_2729_v64_ = d_2729_v64_
                d_2761_v81_: _dafny.Array
                nw451_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_2761_v81_ = nw451_
                d_2762_v82_: D8
                d_2762_v82_ = D8_DC25(d_2761_v81_)
                d_2763_v83_: _dafny.Map
                d_2763_v83_ = _dafny.Map({(self).f10: (d_2762_v82_ if (self).f11 else d_2762_v82_)})
                d_2764_v84_: _dafny.MultiSet
                d_2764_v84_ = _dafny.MultiSet([False])
                d_2765_v85_: _dafny.MultiSet
                d_2765_v85_ = _dafny.MultiSet([d_2764_v84_])
                d_2763_v83_ = (d_2763_v83_).set((d_2765_v85_).issubset(d_2765_v85_), D8_DC25(d_2761_v81_))
                rhs454_ = (self).fm8(p0, (self).f11, self.f12, globalState)
                rhs455_ = False
                rhs456_ = p0
                lhs328_ = self
                lhs329_ = self
                lhs330_ = globalState
                lhs328_.f12 = rhs454_
                lhs329_.f12 = rhs455_
                lhs330_.f1 = rhs456_
            d_2766_v86_: _dafny.MultiSet
            d_2766_v86_ = _dafny.MultiSet([d_2735_v70_, d_2735_v70_, d_2735_v70_, _dafny.CodePoint('v')])
            d_2767_v87_: _dafny.Map
            d_2767_v87_ = _dafny.Map({d_2766_v86_: d_2733_v68_})
            d_2768_v88_: _dafny.Set
            d_2768_v88_ = _dafny.Set({d_2733_v68_, d_2733_v68_})
            d_2769_v90_: _dafny.Map
            d_2769_v90_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "able"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "able")))), default__.fm4((self).f11, self.f12, True, globalState)): d_2729_v64_})
            def iife205_():
                coll65_ = _dafny.Set()
                compr_65_: _dafny.Map
                for compr_65_ in (d_2768_v88_).Elements:
                    d_2770_v89_: _dafny.Map = compr_65_
                    if (d_2770_v89_) in (d_2768_v88_):
                        coll65_ = coll65_.union(_dafny.Set([d_2770_v89_]))
                return _dafny.Set(coll65_)
            rhs457_ = (self).fm9(d_2736_v71_, iife205_()
            , globalState)
            rhs458_ = p0
            rhs459_ = ((d_2769_v90_)[d_2736_v71_] if (d_2736_v71_) in (d_2769_v90_) else d_2729_v64_)
            rhs460_ = (self).f11
            rhs461_ = d_2767_v87_
            lhs331_ = globalState
            lhs332_ = globalState
            lhs333_ = self
            lhs331_.f0 = rhs457_
            lhs332_.f7 = rhs458_
            d_2729_v64_ = rhs459_
            lhs333_.f12 = rhs460_
            d_2767_v87_ = rhs461_
        elif True:
            d_2771___mcc_h4_ = source43_.cf83
            d_2772_cf83_ = d_2771___mcc_h4_
            (self).f12 = not((self).f10)
            (self).f12 = (self).f11
            if (656) != (p0):
                d_2773_v91_: str
                d_2773_v91_ = _dafny.CodePoint('n')
                d_2774_v92_: _dafny.Map
                d_2774_v92_ = _dafny.Map({self.f12: d_2773_v91_})
                d_2775_v93_: D6
                d_2775_v93_ = D6_DC21((self).f10, (self).f10, (self).f11, p0)
                d_2776_v94_: _dafny.Array
                nw452_ = _dafny.Array(None, 22)
                nw452_[int(0)] = d_2773_v91_
                nw452_[int(1)] = d_2773_v91_
                nw452_[int(2)] = d_2773_v91_
                nw452_[int(3)] = d_2773_v91_
                nw452_[int(4)] = d_2773_v91_
                nw452_[int(5)] = d_2773_v91_
                nw452_[int(6)] = d_2773_v91_
                nw452_[int(7)] = (d_2773_v91_ if (self).f11 else _dafny.CodePoint('i'))
                nw452_[int(8)] = default__.fm4((self).f11, self.f12, self.f12, globalState)
                nw452_[int(9)] = (d_2773_v91_ if (self).f10 else d_2773_v91_)
                nw452_[int(10)] = d_2773_v91_
                nw452_[int(11)] = _dafny.CodePoint('c')
                nw452_[int(12)] = (d_2773_v91_ if (self).f10 else ((d_2774_v92_)[(self).fm7(self.f12, globalState)] if ((self).fm7(self.f12, globalState)) in (d_2774_v92_) else d_2773_v91_))
                nw452_[int(13)] = _dafny.CodePoint('u')
                nw452_[int(14)] = d_2773_v91_
                nw452_[int(15)] = d_2773_v91_
                nw452_[int(16)] = _dafny.CodePoint('m')
                nw452_[int(17)] = d_2773_v91_
                nw452_[int(18)] = default__.fm4(not(self.f12), self.f12, (d_2775_v93_).cf37, globalState)
                nw452_[int(19)] = d_2773_v91_
                nw452_[int(20)] = d_2773_v91_
                nw452_[int(21)] = _dafny.CodePoint('v')
                d_2776_v94_ = nw452_
                index480_ = default__.safeIndex(154, (d_2776_v94_).length(0))
                (d_2776_v94_)[index480_] = d_2773_v91_
                d_2777_v95_: _dafny.Array
                nw453_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
                d_2777_v95_ = nw453_
                d_2778_v96_: _dafny.Seq
                d_2778_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjsrprqpl"))
                index481_ = default__.safeIndex(505, (d_2777_v95_).length(0))
                (d_2777_v95_)[index481_] = d_2778_v96_
                d_2779_v97_: D3
                d_2779_v97_ = D3_DC10(d_2778_v96_, (self).f11, (self).f10)
                index482_ = default__.safeIndex(154, (d_2776_v94_).length(0))
                index483_ = default__.safeIndex(505, (d_2777_v95_).length(0))
                rhs462_ = (self).f11
                rhs463_ = default__.safeModuloInt((0) - (len((d_2778_v96_) + (d_2778_v96_))), p0)
                rhs464_ = d_2773_v91_
                rhs465_ = (d_2778_v96_ if self.f12 else (d_2779_v97_).cf14)
                lhs334_ = self
                lhs335_ = globalState
                lhs336_ = d_2776_v94_
                lhs337_ = default__.safeIndex(154, (d_2776_v94_).length(0))
                lhs338_ = d_2777_v95_
                lhs339_ = default__.safeIndex(505, (d_2777_v95_).length(0))
                lhs334_.f12 = rhs462_
                lhs335_.f0 = rhs463_
                lhs336_[lhs337_] = rhs464_
                lhs338_[lhs339_] = rhs465_
                d_2780_v98_: _dafny.Set
                d_2780_v98_ = _dafny.Set({self.f12})
                d_2781_v99_: _dafny.Array
                nw454_ = _dafny.Array(int(0), 13)
                d_2781_v99_ = nw454_
                d_2782_v100_: _dafny.Seq
                d_2782_v100_ = _dafny.SeqWithoutIsStrInference([d_2781_v99_])
                d_2783_v101_: _dafny.Map
                d_2783_v101_ = _dafny.Map({(self).f11: (self).f11})
                d_2784_v102_: _dafny.Set
                d_2784_v102_ = _dafny.Set({d_2783_v101_})
                d_2785_v103_: _dafny.Array
                nw455_ = _dafny.Array(None, 13)
                nw455_[int(0)] = p0
                nw455_[int(1)] = p0
                nw455_[int(2)] = p0
                nw455_[int(3)] = len(d_2780_v98_)
                nw455_[int(4)] = p0
                nw455_[int(5)] = p0
                nw455_[int(6)] = p0
                nw455_[int(7)] = len(d_2782_v100_)
                nw455_[int(8)] = p0
                nw455_[int(9)] = p0
                nw455_[int(10)] = 473
                nw455_[int(11)] = (self).fm9(d_2778_v96_, d_2784_v102_, globalState)
                nw455_[int(12)] = p0
                d_2785_v103_ = nw455_
                d_2786_v104_: _dafny.Seq
                d_2786_v104_ = _dafny.SeqWithoutIsStrInference([d_2781_v99_, d_2785_v103_])
                d_2787_v105_: _dafny.Map
                d_2787_v105_ = _dafny.Map({d_2782_v100_: (self).f11})
                d_2787_v105_ = (d_2787_v105_).set(_dafny.SeqWithoutIsStrInference([d_2785_v103_]), (self).f11)
                d_2788_v106_: _dafny.Set
                d_2788_v106_ = _dafny.Set({((d_2777_v95_)[default__.safeIndex(505, (d_2777_v95_).length(0))]) + (d_2778_v96_), (d_2777_v95_)[default__.safeIndex(505, (d_2777_v95_).length(0))], (d_2777_v95_)[default__.safeIndex(505, (d_2777_v95_).length(0))]})
                d_2788_v106_ = d_2788_v106_
                (globalState).f1 = (p0) * (p0)
                d_2778_v96_ = d_2778_v96_
            elif True:
                (globalState).f0 = (0) - ((0) - (p0))
                rhs466_ = (self).f10
                rhs467_ = p0
                lhs340_ = self
                lhs341_ = globalState
                lhs340_.f12 = rhs466_
                lhs341_.f0 = rhs467_
                d_2789_v107_: _dafny.Array
                nw456_ = _dafny.Array(None, 29)
                nw456_[int(0)] = (self).f11
                nw456_[int(1)] = not(False)
                nw456_[int(2)] = (self).f11
                nw456_[int(3)] = (self).f11
                nw456_[int(4)] = (self).f10
                nw456_[int(5)] = (self).f10
                nw456_[int(6)] = self.f12
                nw456_[int(7)] = (self).f11
                nw456_[int(8)] = self.f12
                nw456_[int(9)] = (self).f11
                nw456_[int(10)] = False
                nw456_[int(11)] = (self).f11
                nw456_[int(12)] = self.f12
                nw456_[int(13)] = (self).f10
                nw456_[int(14)] = self.f12
                nw456_[int(15)] = (self).f11
                nw456_[int(16)] = self.f12
                nw456_[int(17)] = (self).f10
                nw456_[int(18)] = default__.fm0(p0, globalState)
                nw456_[int(19)] = self.f12
                nw456_[int(20)] = True
                nw456_[int(21)] = self.f12
                nw456_[int(22)] = (self).f10
                nw456_[int(23)] = (self).f10
                nw456_[int(24)] = (self).f11
                nw456_[int(25)] = (self).f10
                nw456_[int(26)] = (self).f11
                nw456_[int(27)] = (self).f10
                nw456_[int(28)] = self.f12
                d_2789_v107_ = nw456_
                d_2790_v108_: D0
                d_2790_v108_ = D0_DC0(d_2789_v107_)
                d_2791_v109_: D18
                d_2791_v109_ = D18_DC56((d_2667_v20_) + (_dafny.SeqWithoutIsStrInference([(self).f10])))
                d_2792_v110_: str
                d_2792_v110_ = _dafny.CodePoint('x')
                d_2793_v111_: _dafny.Seq
                d_2793_v111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wknhrdj"))
                rhs468_ = D0_DC0(d_2789_v107_)
                rhs469_ = d_2791_v109_
                rhs470_ = p0
                rhs471_ = d_2792_v110_
                rhs472_ = (d_2793_v111_)[default__.safeIndex(p0, len(d_2793_v111_))]
                lhs342_ = globalState
                d_2790_v108_ = rhs468_
                d_2791_v109_ = rhs469_
                lhs342_.f1 = rhs470_
                d_2792_v110_ = rhs471_
                d_2792_v110_ = rhs472_
                (self).f12 = self.f12
                d_2794_v112_: _dafny.Seq
                d_2795_v113_: _dafny.Seq
                out74_: _dafny.Seq
                out75_: _dafny.Seq
                out74_, out75_ = default__.m0((0) - ((p0) - (507)), d_2789_v107_, globalState)
                d_2794_v112_ = out74_
                d_2795_v113_ = out75_
            d_2796_v114_: str
            d_2796_v114_ = _dafny.CodePoint('o')
            d_2796_v114_ = d_2796_v114_
        d_2797_v115_: _dafny.Array
        nw457_ = _dafny.Array(D3.default()(), 25)
        d_2797_v115_ = nw457_
        rhs473_ = not (self.f12) or (self.f12)
        rhs474_ = d_2797_v115_
        lhs343_ = self
        lhs343_.f12 = rhs473_
        d_2797_v115_ = rhs474_
        d_2798_v116_: _dafny.Array
        def lambda129_(d_2799_i9_):
            return False

        init72_ = lambda129_
        nw458_ = _dafny.Array(None, 4)
        for i0_72_ in range(nw458_.length(0)):
            nw458_[i0_72_] = init72_(i0_72_)
        d_2798_v116_ = nw458_
        d_2800_v117_: _dafny.Map
        d_2800_v117_ = _dafny.Map({d_2798_v116_: self.f12})
        d_2800_v117_ = (d_2800_v117_).set((d_2798_v116_ if self.f12 else d_2798_v116_), self.f12)

    def m6(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: int = int(0)
        d_2801_i0_: int
        d_2801_i0_ = 0
        with _dafny.label("15"):
            while (self).f10:
                with _dafny.c_label("15"):
                    if (d_2801_i0_) >= (100):
                        raise _dafny.Break("15")
                    d_2801_i0_ = (d_2801_i0_) + (1)
                    (self).f12 = (self).f10
                    d_2802_v0_: _dafny.Map
                    d_2802_v0_ = _dafny.Map({(self).f10: p0})
                    r0 = (((d_2802_v0_)[p1] if (p1) in (d_2802_v0_) else p0)) != ((0) - (p0))
                    d_2803_v1_: _dafny.Array
                    nw459_ = _dafny.Array(int(0), 29)
                    d_2803_v1_ = nw459_
                    d_2804_v2_: _dafny.Seq
                    d_2804_v2_ = _dafny.SeqWithoutIsStrInference([(d_2803_v1_ if self.f12 else d_2803_v1_)])
                    d_2804_v2_ = d_2804_v2_
                    d_2805_v3_: _dafny.Array
                    nw460_ = _dafny.Array(False, 7)
                    d_2805_v3_ = nw460_
                    d_2806_v4_: _dafny.Seq
                    d_2806_v4_ = _dafny.SeqWithoutIsStrInference([False])
                    d_2807_v5_: _dafny.Array
                    nw461_ = _dafny.Array(None, 22)
                    nw461_[int(0)] = d_2805_v3_
                    nw461_[int(1)] = d_2805_v3_
                    nw461_[int(2)] = d_2805_v3_
                    nw461_[int(3)] = (d_2805_v3_ if (d_2806_v4_)[default__.safeIndex(p0, len(d_2806_v4_))] else d_2805_v3_)
                    nw461_[int(4)] = d_2805_v3_
                    nw461_[int(5)] = d_2805_v3_
                    nw461_[int(6)] = d_2805_v3_
                    nw461_[int(7)] = d_2805_v3_
                    nw461_[int(8)] = d_2805_v3_
                    nw461_[int(9)] = d_2805_v3_
                    nw461_[int(10)] = (d_2805_v3_ if p1 else d_2805_v3_)
                    nw461_[int(11)] = d_2805_v3_
                    nw461_[int(12)] = d_2805_v3_
                    nw461_[int(13)] = (d_2805_v3_ if not(not(True)) else d_2805_v3_)
                    nw461_[int(14)] = d_2805_v3_
                    nw461_[int(15)] = d_2805_v3_
                    nw461_[int(16)] = d_2805_v3_
                    nw461_[int(17)] = d_2805_v3_
                    nw461_[int(18)] = d_2805_v3_
                    nw461_[int(19)] = d_2805_v3_
                    nw461_[int(20)] = d_2805_v3_
                    nw461_[int(21)] = d_2805_v3_
                    d_2807_v5_ = nw461_
                    index484_ = default__.safeIndex(467, (d_2807_v5_).length(0))
                    (d_2807_v5_)[index484_] = (d_2805_v3_ if p1 else d_2805_v3_)
                    index485_ = default__.safeIndex(467, (d_2807_v5_).length(0))
                    (d_2807_v5_)[index485_] = d_2805_v3_
                    pass
            pass
        (globalState).f7 = p0
        d_2808_v6_: _dafny.Map
        d_2808_v6_ = _dafny.Map({(0) - (p0): (self).f11})
        d_2809_v7_: D11
        d_2809_v7_ = D11_DC32(d_2808_v6_)
        pat_let_tv74_ = p0
        def lambda130_(source45_):
            if source45_.is_DC33:
                d_2810___mcc_h0_ = source45_.cf54
                d_2811___mcc_h1_ = source45_.cf55
                d_2812_cf55_ = d_2811___mcc_h1_
                d_2813_cf54_ = d_2810___mcc_h0_
                return pat_let_tv74_
            elif source45_.is_DC34:
                d_2814___mcc_h2_ = source45_.cf56
                d_2815___mcc_h3_ = source45_.cf57
                d_2816___mcc_h4_ = source45_.cf58
                d_2817_cf58_ = d_2816___mcc_h4_
                d_2818_cf57_ = d_2815___mcc_h3_
                d_2819_cf56_ = d_2814___mcc_h2_
                return len(_dafny.SeqWithoutIsStrInference([self.f12]))
            elif True:
                d_2820___mcc_h5_ = source45_.cf53
                d_2821_cf53_ = d_2820___mcc_h5_
                return len((_dafny.SeqWithoutIsStrInference([(self).f11, self.f12, self.f12])) + (_dafny.SeqWithoutIsStrInference([(self).f11])))

        (globalState).f7 = lambda130_(d_2809_v7_)
        d_2822_v8_: _dafny.Seq
        d_2822_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "julndsliu"))
        d_2823_v9_: str
        d_2823_v9_ = _dafny.CodePoint('f')
        d_2822_v8_ = (((d_2822_v8_).set(default__.safeIndex(p0, len(d_2822_v8_)), d_2823_v9_)) + (d_2822_v8_)) + ((d_2822_v8_) + (d_2822_v8_))
        d_2824_v10_: _dafny.Array
        nw462_ = _dafny.Array(_dafny.Array(None, 0), 10)
        d_2824_v10_ = nw462_
        d_2825_v11_: D17
        d_2825_v11_ = D17_DC50(d_2824_v10_)
        d_2826_v12_: D17
        d_2826_v12_ = D17_DC52(d_2825_v11_)
        d_2827_v13_: D17
        d_2827_v13_ = D17_DC52(d_2826_v12_)
        d_2828_v14_: _dafny.Seq
        d_2828_v14_ = _dafny.SeqWithoutIsStrInference([d_2827_v13_])
        d_2829_v15_: _dafny.MultiSet
        d_2829_v15_ = _dafny.MultiSet([d_2828_v14_, d_2828_v14_])
        d_2830_v16_: _dafny.Seq
        d_2830_v16_ = _dafny.SeqWithoutIsStrInference([d_2828_v14_])
        d_2829_v15_ = ((d_2829_v15_) - (_dafny.MultiSet(d_2830_v16_))) | (d_2829_v15_)
        d_2831_v17_: _dafny.Array
        def lambda131_(d_2832_i1_):
            return (self).f10

        init73_ = lambda131_
        nw463_ = _dafny.Array(None, 25)
        for i0_73_ in range(nw463_.length(0)):
            nw463_[i0_73_] = init73_(i0_73_)
        d_2831_v17_ = nw463_
        index486_ = default__.safeIndex(923, (d_2831_v17_).length(0))
        (d_2831_v17_)[index486_] = (self).fm8(p0, self.f12, (self).f11, globalState)
        d_2833_v18_: D14
        d_2833_v18_ = D14_DC43(d_2823_v9_, p0, (0) - (p0), p0)
        index487_ = default__.safeIndex(923, (d_2831_v17_).length(0))
        rhs475_ = len(_dafny.SeqWithoutIsStrInference([(0) - ((d_2833_v18_).cf78)]))
        rhs476_ = not (self.f12) or (False)
        rhs477_ = len(default__.fm11(((self).f11 if True else self.f12), p0, p0, globalState))
        lhs344_ = globalState
        lhs345_ = d_2831_v17_
        lhs346_ = default__.safeIndex(923, (d_2831_v17_).length(0))
        lhs347_ = globalState
        lhs344_.f1 = rhs475_
        lhs345_[lhs346_] = rhs476_
        lhs347_.f7 = rhs477_
        r0 = not (self.f12) or ((self).f11)
        r1 = -942
        return r0, r1

    @property
    def f11(self):
        return self._f11

class C13:
    def  __init__(self):
        self._f8: bool = False
        self._f9: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C13"
    def ctor__(self, f8, f9):
        (self)._f8 = f8
        (self)._f9 = f9

    def fm5(self, p0, globalState):
        return ((len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(0) - (len((D4_DC13(_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f8, (self).f8]): False}))).cf20))]))]))) + (-400)) * ((0) - (len(_dafny.Set({(self).f9}))))

    def m1(self, p0, globalState):
        r0: bool = False
        d_2834_v0_: _dafny.Seq
        d_2834_v0_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
        d_2835_v1_: C11
        nw464_ = C11()
        nw464_.ctor__(len(d_2834_v0_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2836_i0_ in range(default__.abs(-595))]), (self).f9)
        d_2835_v1_ = nw464_
        d_2837_i1_: int
        d_2837_i1_ = 0
        with _dafny.label("16"):
            while (self).f8:
                with _dafny.c_label("16"):
                    if (d_2837_i1_) >= (100):
                        raise _dafny.Break("16")
                    d_2837_i1_ = (d_2837_i1_) + (1)
                    d_2838_v2_: C7
                    nw465_ = C7()
                    nw465_.ctor__(not ((self).f9) or ((self).f8), (self).f8)
                    d_2838_v2_ = nw465_
                    d_2839_v3_: C5
                    nw466_ = C5()
                    nw466_.ctor__()
                    d_2839_v3_ = nw466_
                    d_2840_v4_: _dafny.MultiSet
                    d_2840_v4_ = _dafny.MultiSet([d_2839_v3_])
                    d_2840_v4_ = _dafny.MultiSet([d_2839_v3_, d_2839_v3_])
                    d_2841_v5_: D7
                    d_2841_v5_ = D7_DC23()
                    source46_ = d_2841_v5_
                    if source46_.is_DC23:
                        (globalState).f1 = (d_2835_v1_).f13
                        d_2842_v6_: _dafny.Seq
                        d_2842_v6_ = _dafny.SeqWithoutIsStrInference([(d_2838_v2_).f19])
                        d_2843_v7_: _dafny.Map
                        d_2843_v7_ = _dafny.Map({(d_2838_v2_).f19: d_2842_v6_})
                        d_2844_v8_: _dafny.Array
                        def lambda132_(d_2845_v1_):
                            def lambda133_(d_2846_i2_):
                                return (d_2845_v1_).f14

                            return lambda133_

                        init74_ = lambda132_(d_2835_v1_)
                        nw467_ = _dafny.Array(None, 20)
                        for i0_74_ in range(nw467_.length(0)):
                            nw467_[i0_74_] = init74_(i0_74_)
                        d_2844_v8_ = nw467_
                        d_2847_v9_: _dafny.Set
                        d_2847_v9_ = _dafny.Set({(d_2835_v1_).f14})
                        d_2848_v10_: str
                        d_2848_v10_ = _dafny.CodePoint('k')
                        def iife206_():
                            coll66_ = _dafny.Set()
                            compr_66_: _dafny.Seq
                            for compr_66_ in ((_dafny.MultiSet([(d_2835_v1_).f14, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2849_i3_ in range(default__.abs(641))]), (d_2835_v1_).f14, (d_2835_v1_).f14, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwulthy"))])).set(((d_2835_v1_).f14).set(default__.safeIndex((d_2835_v1_).f13, len((d_2835_v1_).f14)), d_2848_v10_), default__.abs(293))).Elements:
                                d_2850_v11_: _dafny.Seq = compr_66_
                                if (d_2850_v11_) in ((_dafny.MultiSet([(d_2835_v1_).f14, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2849_i3_ in range(default__.abs(641))]), (d_2835_v1_).f14, (d_2835_v1_).f14, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwulthy"))])).set(((d_2835_v1_).f14).set(default__.safeIndex((d_2835_v1_).f13, len((d_2835_v1_).f14)), d_2848_v10_), default__.abs(293))):
                                    coll66_ = coll66_.union(_dafny.Set([d_2850_v11_]))
                            return _dafny.Set(coll66_)
                        rhs478_ = d_2843_v7_
                        rhs479_ = (d_2844_v8_ if (iife206_()
                        ).issubset(d_2847_v9_) else d_2844_v8_)
                        d_2843_v7_ = rhs478_
                        d_2844_v8_ = rhs479_
                        d_2851_v12_: _dafny.Array
                        nw468_ = _dafny.Array(int(0), 20)
                        d_2851_v12_ = nw468_
                        index488_ = default__.safeIndex(898, (d_2851_v12_).length(0))
                        (d_2851_v12_)[index488_] = default__.safeDivisionInt(p0, 116)
                        d_2852_v13_: _dafny.MultiSet
                        d_2852_v13_ = _dafny.MultiSet([False, (d_2838_v2_).f19])
                        d_2853_v14_: _dafny.Map
                        d_2853_v14_ = _dafny.Map({(d_2835_v1_).f13: 441})
                        index489_ = default__.safeIndex(898, (d_2851_v12_).length(0))
                        (d_2851_v12_)[index489_] = (((d_2852_v13_) - (_dafny.MultiSet((d_2842_v6_).set(default__.safeIndex(((d_2853_v14_)[(0) - (p0)] if ((0) - (p0)) in (d_2853_v14_) else (0) - (p0)), len(d_2842_v6_)), (d_2838_v2_).f19)))) - ((d_2852_v13_) - (d_2852_v13_))).cardinality
                        d_2854_v15_: _dafny.Map
                        d_2854_v15_ = _dafny.Map({(d_2838_v2_).f19: ((_dafny.MultiSet([d_2853_v14_])).set(d_2853_v14_, default__.abs((d_2835_v1_).f13))).cardinality})
                        d_2855_v16_: _dafny.Map
                        d_2855_v16_ = _dafny.Map({(d_2851_v12_)[default__.safeIndex(898, (d_2851_v12_).length(0))]: (self).f9})
                        d_2856_v17_: _dafny.Seq
                        d_2856_v17_ = _dafny.SeqWithoutIsStrInference([(d_2853_v14_).set(len(d_2855_v16_), p0)])
                        d_2854_v15_ = (d_2854_v15_).set((d_2838_v2_).f19, len((d_2856_v17_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hioaqwmh"))), len(d_2856_v17_))]))
                    elif source46_.is_DC24:
                        d_2857___mcc_h0_ = source46_.cf41
                        d_2858_cf41_ = d_2857___mcc_h0_
                        d_2859_v18_: _dafny.Seq
                        d_2859_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([859, (d_2835_v1_).f13]), d_2834_v0_])
                        d_2860_v19_: _dafny.Array
                        nw469_ = _dafny.Array(_dafny.CodePoint('D'), 12)
                        d_2860_v19_ = nw469_
                        d_2861_v20_: str
                        d_2861_v20_ = _dafny.CodePoint('t')
                        d_2862_v21_: _dafny.Seq
                        d_2862_v21_ = _dafny.SeqWithoutIsStrInference([(d_2859_v18_)[default__.safeIndex((D12_DC37((_dafny.MultiSet(d_2834_v0_)).cardinality, d_2860_v19_, d_2861_v20_)).cf61, len(d_2859_v18_))]])
                        d_2863_v22_: _dafny.Set
                        d_2863_v22_ = _dafny.Set({True})
                        d_2864_v23_: D10
                        d_2864_v23_ = D10_DC30(d_2863_v22_)
                        d_2859_v18_ = default__.fm59(True, ((d_2835_v1_).f14) + ((d_2835_v1_).f14), d_2864_v23_, (d_2835_v1_).f13, globalState)
                        d_2865_v24_: _dafny.Map
                        d_2865_v24_ = _dafny.Map({(d_2838_v2_).f19: (self).f8})
                        d_2866_v25_: _dafny.Map
                        d_2866_v25_ = _dafny.Map({_dafny.MultiSet(d_2834_v0_): _dafny.CodePoint('t')})
                        d_2867_v27_: _dafny.MultiSet
                        def iife207_():
                            coll67_ = _dafny.Map()
                            compr_67_: _dafny.Seq
                            for compr_67_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2834_v0_, d_2834_v0_]))).Elements:
                                d_2868_v26_: _dafny.Seq = compr_67_
                                if (d_2868_v26_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2834_v0_, d_2834_v0_]))):
                                    coll67_[d_2868_v26_] = -199
                            return _dafny.Map(coll67_)
                        d_2867_v27_ = _dafny.MultiSet([(d_2835_v1_).f13, len((iife207_()
                        ).set(_dafny.SeqWithoutIsStrInference([d_2858_cf41_ for d_2869_i4_ in range(default__.abs(572))]), len(_dafny.Map({d_2858_cf41_: (d_2835_v1_).f13}))))])
                        d_2870_v28_: D15
                        d_2870_v28_ = D15_DC46()
                        d_2871_v29_: bool
                        d_2872_v30_: _dafny.Seq
                        d_2873_v31_: _dafny.Map
                        out76_: bool
                        out77_: _dafny.Seq
                        out78_: _dafny.Map
                        out76_, out77_, out78_ = (d_2838_v2_).m3(default__.fm60(default__.fm61(not((self).f9), d_2865_v24_, (d_2835_v1_).f13, (self).f9, globalState), ((d_2866_v25_)[d_2867_v27_] if (d_2867_v27_) in (d_2866_v25_) else _dafny.CodePoint('e')), (self).f9, d_2870_v28_, globalState), (self).f8, globalState)
                        d_2871_v29_ = out76_
                        d_2872_v30_ = out77_
                        d_2873_v31_ = out78_
                        r0 = ((d_2858_cf41_) + (p0)) != (default__.safeModuloInt((d_2867_v27_).cardinality, 818))
                        d_2874_v32_: _dafny.Set
                        d_2874_v32_ = _dafny.Set({d_2861_v20_})
                        d_2875_v33_: _dafny.MultiSet
                        d_2875_v33_ = _dafny.MultiSet([(d_2867_v27_).ispropersubset(_dafny.MultiSet([(d_2835_v1_).f13, len(d_2874_v32_), (d_2835_v1_).f13])), ((d_2838_v2_).f19 if (d_2838_v2_).f19 else (d_2838_v2_).f19), (d_2838_v2_).f19])
                        rhs480_ = default__.safeDivisionInt(len((d_2835_v1_).f14), (305) - (d_2858_cf41_))
                        rhs481_ = (d_2835_v1_).f13
                        rhs482_ = d_2875_v33_
                        lhs348_ = globalState
                        lhs348_.f0 = rhs480_
                        d_2858_cf41_ = rhs481_
                        d_2875_v33_ = rhs482_
                    elif True:
                        d_2876___mcc_h1_ = source46_.cf40
                        d_2877_cf40_ = d_2876___mcc_h1_
                        d_2878_v34_: _dafny.Array
                        nw470_ = _dafny.Array(_dafny.Array(None, 0), 1)
                        d_2878_v34_ = nw470_
                        d_2879_v35_: T0
                        nw471_ = C12()
                        nw471_.ctor__((d_2838_v2_).f19, (d_2838_v2_).f19, (d_2838_v2_).f19)
                        d_2879_v35_ = nw471_
                        d_2880_v36_: _dafny.Map
                        d_2880_v36_ = _dafny.Map({((d_2835_v1_).f13) + ((d_2835_v1_).f13): _dafny.Map({(self).f8: default__.fm0(len(_dafny.Set({d_2879_v35_})), globalState)})})
                        rhs483_ = d_2878_v34_
                        rhs484_ = d_2880_v36_
                        d_2878_v34_ = rhs483_
                        d_2880_v36_ = rhs484_
                        r0 = (d_2879_v35_).f10
                        d_2881_v37_: _dafny.Map
                        d_2881_v37_ = _dafny.Map({p0: (d_2838_v2_).fm7(not((self).f8), globalState)})
                        d_2882_v38_: C2
                        nw472_ = C2()
                        nw472_.ctor__(((d_2881_v37_)[p0] if (p0) in (d_2881_v37_) else (d_2879_v35_).f10), 251)
                        d_2882_v38_ = nw472_
                        d_2883_v39_: _dafny.Map
                        d_2883_v39_ = _dafny.Map({(d_2879_v35_).f10: (self).f8})
                        d_2884_v40_: _dafny.MultiSet
                        d_2884_v40_ = _dafny.MultiSet([(self).f9, (d_2882_v38_).f21])
                        d_2885_v41_: _dafny.Seq
                        d_2885_v41_ = _dafny.SeqWithoutIsStrInference([d_2884_v40_])
                        d_2886_v42_: _dafny.Map
                        d_2886_v42_ = _dafny.Map({d_2835_v1_: d_2885_v41_})
                        d_2887_v43_: _dafny.Seq
                        d_2887_v43_ = _dafny.SeqWithoutIsStrInference([(self).f8])
                        d_2888_v44_: _dafny.Map
                        d_2888_v44_ = _dafny.Map({default__.safeModuloInt(len(d_2883_v39_), (d_2835_v1_).f13): ((d_2886_v42_)[d_2835_v1_] if (d_2835_v1_) in (d_2886_v42_) else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_2887_v43_)]))})
                        d_2889_v45_: _dafny.Array
                        nw473_ = _dafny.Array(_dafny.Array(None, 0), 28)
                        d_2889_v45_ = nw473_
                        d_2890_v46_: D8
                        d_2890_v46_ = D8_DC25(d_2889_v45_)
                        d_2891_v47_: _dafny.MultiSet
                        d_2891_v47_ = _dafny.MultiSet([d_2890_v46_, D8_DC25(d_2889_v45_)])
                        d_2888_v44_ = (d_2888_v44_).set(((d_2884_v40_)[(d_2838_v2_).f19] if ((d_2838_v2_).f19) in (d_2884_v40_) else (0) - ((d_2891_v47_).cardinality)), d_2885_v41_)
                    r0 = False
                    pass
            pass
        d_2892_v48_: _dafny.Seq
        d_2892_v48_ = _dafny.SeqWithoutIsStrInference([(self).f9, (self).f8, (self).f8, not((self).f9), (self).f9])
        d_2893_v49_: _dafny.Map
        d_2893_v49_ = _dafny.Map({(d_2835_v1_).fm8(p0, (self).f9, (self).f8, globalState): (d_2892_v48_)[default__.safeIndex((d_2834_v0_)[default__.safeIndex(p0, len(d_2834_v0_))], len(d_2892_v48_))]})
        d_2894_v50_: _dafny.Array
        nw474_ = _dafny.Array(None, 3)
        nw474_[int(0)] = False
        nw474_[int(1)] = (self).f9
        nw474_[int(2)] = ((d_2893_v49_)[(self).f9] if ((self).f9) in (d_2893_v49_) else (self).f9)
        d_2894_v50_ = nw474_
        index490_ = default__.safeIndex(235, (d_2894_v50_).length(0))
        (d_2894_v50_)[index490_] = (self).f9
        index491_ = default__.safeIndex(235, (d_2894_v50_).length(0))
        (d_2894_v50_)[index491_] = (self).f9
        d_2895_v51_: _dafny.Seq
        d_2895_v51_ = _dafny.SeqWithoutIsStrInference([(d_2835_v1_).f14, (d_2835_v1_).f14, (d_2835_v1_).f14, (d_2835_v1_).f14, (d_2835_v1_).f14])
        d_2895_v51_ = d_2895_v51_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_2894_v50_).length(0)):
            d_2896_i5_: int = guard_loop_5_
            if (True) and (((0) <= (d_2896_i5_)) and ((d_2896_i5_) < ((d_2894_v50_).length(0)))):
                (d_2894_v50_)[(d_2896_i5_)] = True
        d_2897_v52_: C3
        nw475_ = C3()
        nw475_.ctor__((self).f9, (self).f9)
        d_2897_v52_ = nw475_
        rhs485_ = 838
        rhs486_ = d_2897_v52_
        lhs349_ = globalState
        lhs349_.f1 = rhs485_
        d_2897_v52_ = rhs486_
        r0 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnr"))) <= ((d_2835_v1_).f14) if (self).f8 else (self).f8)
        return r0

    def m2(self, p0, p1, p2, globalState):
        (globalState).f0 = 20
        d_2898_v0_: _dafny.Array
        nw476_ = _dafny.Array(_dafny.Array(None, 0), 27)
        d_2898_v0_ = nw476_
        d_2899_v1_: str
        d_2899_v1_ = _dafny.CodePoint('x')
        d_2900_v2_: _dafny.Seq
        d_2900_v2_ = _dafny.SeqWithoutIsStrInference([d_2899_v1_, d_2899_v1_])
        d_2901_v3_: _dafny.Map
        d_2901_v3_ = _dafny.Map({d_2899_v1_: d_2899_v1_})
        d_2902_v4_: D3
        d_2902_v4_ = D3_DC10(d_2900_v2_, (self).f9, (self).f9)
        d_2903_v5_: _dafny.Array
        nw477_ = _dafny.Array(None, 29)
        nw477_[int(0)] = d_2900_v2_
        nw477_[int(1)] = _dafny.SeqWithoutIsStrInference([d_2899_v1_, ((d_2901_v3_)[d_2899_v1_] if (d_2899_v1_) in (d_2901_v3_) else d_2899_v1_)])
        nw477_[int(2)] = d_2900_v2_
        nw477_[int(3)] = d_2900_v2_
        nw477_[int(4)] = (d_2902_v4_).cf14
        nw477_[int(5)] = d_2900_v2_
        nw477_[int(6)] = d_2900_v2_
        nw477_[int(7)] = default__.fm25(p2, p0, p2, globalState)
        nw477_[int(8)] = d_2900_v2_
        nw477_[int(9)] = d_2900_v2_
        nw477_[int(10)] = d_2900_v2_
        nw477_[int(11)] = d_2900_v2_
        nw477_[int(12)] = d_2900_v2_
        nw477_[int(13)] = d_2900_v2_
        nw477_[int(14)] = d_2900_v2_
        nw477_[int(15)] = d_2900_v2_
        nw477_[int(16)] = d_2900_v2_
        nw477_[int(17)] = (d_2902_v4_).cf14
        nw477_[int(18)] = _dafny.SeqWithoutIsStrInference([d_2899_v1_, default__.fm4(p1, p1, (self).f8, globalState)])
        nw477_[int(19)] = _dafny.SeqWithoutIsStrInference([d_2899_v1_ for d_2904_i0_ in range(default__.abs(189))])
        nw477_[int(20)] = d_2900_v2_
        nw477_[int(21)] = d_2900_v2_
        nw477_[int(22)] = d_2900_v2_
        nw477_[int(23)] = d_2900_v2_
        nw477_[int(24)] = d_2900_v2_
        nw477_[int(25)] = d_2900_v2_
        nw477_[int(26)] = d_2900_v2_
        nw477_[int(27)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2905_i1_ in range(default__.abs(357))])
        nw477_[int(28)] = d_2900_v2_
        d_2903_v5_ = nw477_
        index492_ = default__.safeIndex(198, (d_2898_v0_).length(0))
        (d_2898_v0_)[index492_] = d_2903_v5_
        d_2906_v6_: _dafny.Map
        d_2906_v6_ = _dafny.Map({not(default__.fm0(488, globalState)): (0) - (p2)})
        d_2907_v7_: _dafny.Set
        d_2907_v7_ = _dafny.Set({10})
        index493_ = default__.safeIndex(198, (d_2898_v0_).length(0))
        rhs487_ = d_2903_v5_
        rhs488_ = d_2903_v5_
        rhs489_ = _dafny.Map({((self).f8 if p0 else (self).f9): len((_dafny.Set({(0) - (p2)})) | (d_2907_v7_))})
        lhs350_ = d_2898_v0_
        lhs351_ = default__.safeIndex(198, (d_2898_v0_).length(0))
        lhs350_[lhs351_] = rhs487_
        d_2903_v5_ = rhs488_
        d_2906_v6_ = rhs489_
        d_2908_v8_: _dafny.Seq
        d_2908_v8_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2909_v9_: D4
        d_2909_v9_ = D4_DC14(p2, p2)
        hi19_ = (d_2909_v9_).cf22
        for d_2910_i2_ in range(default__.safeDivisionInt(p2, len(d_2908_v8_)), hi19_):
            d_2911_v10_: _dafny.Map
            d_2911_v10_ = _dafny.Map({d_2910_i2_: p0})
            d_2912_v11_: _dafny.Map
            d_2912_v11_ = _dafny.Map({(self).f8: (self).f8})
            d_2913_v12_: _dafny.Set
            d_2913_v12_ = _dafny.Set({((d_2912_v11_)[p0] if (p0) in (d_2912_v11_) else False)})
            d_2914_v13_: _dafny.Map
            d_2914_v13_ = _dafny.Map({d_2911_v10_: default__.safeModuloInt(len(d_2913_v12_), p2)})
            d_2914_v13_ = (d_2914_v13_).set(d_2911_v10_, p2)
            d_2900_v2_ = default__.fm13(p1, ((0) - (p2)) != (d_2910_i2_), globalState)
            if p0:
                d_2915_v14_: _dafny.Array
                nw478_ = _dafny.Array(False, 7)
                d_2915_v14_ = nw478_
                index494_ = default__.safeIndex(209, (d_2915_v14_).length(0))
                (d_2915_v14_)[index494_] = (p0) == (p0)
                index495_ = default__.safeIndex(209, (d_2915_v14_).length(0))
                (d_2915_v14_)[index495_] = p0
                d_2916_v15_: _dafny.Seq
                d_2917_v16_: _dafny.Seq
                out79_: _dafny.Seq
                out80_: _dafny.Seq
                out79_, out80_ = default__.m0(782, d_2915_v14_, globalState)
                d_2916_v15_ = out79_
                d_2917_v16_ = out80_
                d_2918_v17_: C6
                nw479_ = C6()
                nw479_.ctor__()
                d_2918_v17_ = nw479_
                d_2919_v18_: C3
                nw480_ = C3()
                nw480_.ctor__((p2) >= (p2), p1)
                d_2919_v18_ = nw480_
                d_2920_v19_: C2
                nw481_ = C2()
                nw481_.ctor__(p0, (p2 if (self).f8 else p2))
                d_2920_v19_ = nw481_
            elif True:
                d_2921_v20_: _dafny.Seq
                d_2921_v20_ = _dafny.SeqWithoutIsStrInference([646, p2, 154])
                d_2922_v21_: D11
                d_2922_v21_ = D11_DC33((self).f9, p2)
                d_2923_v22_: _dafny.Array
                nw482_ = _dafny.Array(None, 14)
                nw482_[int(0)] = False
                nw482_[int(1)] = p0
                nw482_[int(2)] = (d_2922_v21_).cf54
                nw482_[int(3)] = (self).f9
                nw482_[int(4)] = (self).f9
                nw482_[int(5)] = (self).f9
                nw482_[int(6)] = (self).f9
                nw482_[int(7)] = (self).f9
                nw482_[int(8)] = p0
                nw482_[int(9)] = p1
                nw482_[int(10)] = (self).f9
                nw482_[int(11)] = p0
                nw482_[int(12)] = (self).f9
                nw482_[int(13)] = (self).f9
                d_2923_v22_ = nw482_
                d_2924_v23_: _dafny.Seq
                d_2925_v24_: _dafny.Seq
                out81_: _dafny.Seq
                out82_: _dafny.Seq
                out81_, out82_ = default__.m0(len(d_2921_v20_), d_2923_v22_, globalState)
                d_2924_v23_ = out81_
                d_2925_v24_ = out82_
                arr10_ = (d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]
                index496_ = default__.safeIndex(699, ((d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]).length(0))
                arr10_[index496_] = ((d_2924_v23_).set(default__.safeIndex(d_2910_i2_, len(d_2924_v23_)), d_2899_v1_)) + (d_2924_v23_)
                d_2926_v25_: _dafny.Seq
                d_2926_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxw"))])
                arr11_ = (d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]
                index497_ = default__.safeIndex(699, ((d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]).length(0))
                arr11_[index497_] = (d_2926_v25_)[default__.safeIndex(d_2910_i2_, len(d_2926_v25_))]
                index498_ = default__.safeIndex(291, (d_2923_v22_).length(0))
                def iife208_():
                    coll68_ = _dafny.Map()
                    compr_68_: str
                    for compr_68_ in (d_2924_v23_).Elements:
                        d_2927_v26_: str = compr_68_
                        if (d_2927_v26_) in (d_2924_v23_):
                            coll68_[d_2927_v26_] = p0
                    return _dafny.Map(coll68_)
                (d_2923_v22_)[index498_] = (len(((d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))])[default__.safeIndex(699, ((d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]).length(0))])) <= ((_dafny.MultiSet([p0, (d_2908_v8_)[default__.safeIndex(len(iife208_()
                ), len(d_2908_v8_))]])).cardinality)
                index499_ = default__.safeIndex(291, (d_2923_v22_).length(0))
                (d_2923_v22_)[index499_] = p0
                d_2928_v28_: D19
                def iife209_():
                    coll69_ = _dafny.Set()
                    compr_69_: int
                    for compr_69_ in _dafny.IntegerRange(219, -661):
                        d_2929_v27_: int = compr_69_
                        if ((219) <= (d_2929_v27_)) and ((d_2929_v27_) < (-661)):
                            coll69_ = coll69_.union(_dafny.Set([(d_2929_v27_) + (len(d_2925_v24_))]))
                    return _dafny.Set(coll69_)
                d_2928_v28_ = D19_DC57(iife209_()
)
                d_2930_v29_: _dafny.Seq
                d_2930_v29_ = _dafny.SeqWithoutIsStrInference([d_2907_v7_, d_2907_v7_, d_2907_v7_, (d_2928_v28_).cf98, d_2907_v7_])
                d_2907_v7_ = (d_2930_v29_)[default__.safeIndex(len(((d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))])[default__.safeIndex(699, ((d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]).length(0))]), len(d_2930_v29_))]
                d_2931_v30_: _dafny.Seq
                d_2931_v30_ = _dafny.SeqWithoutIsStrInference([d_2911_v10_, d_2911_v10_, _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aldu"))): (d_2923_v22_)[default__.safeIndex(291, (d_2923_v22_).length(0))]})])
                d_2932_v31_: _dafny.Array
                nw483_ = _dafny.Array(None, 24)
                nw483_[int(0)] = d_2899_v1_
                nw483_[int(1)] = d_2899_v1_
                nw483_[int(2)] = d_2899_v1_
                nw483_[int(3)] = d_2899_v1_
                nw483_[int(4)] = d_2899_v1_
                nw483_[int(5)] = d_2899_v1_
                nw483_[int(6)] = _dafny.CodePoint('b')
                nw483_[int(7)] = d_2899_v1_
                nw483_[int(8)] = d_2899_v1_
                nw483_[int(9)] = d_2899_v1_
                nw483_[int(10)] = d_2899_v1_
                nw483_[int(11)] = d_2899_v1_
                nw483_[int(12)] = default__.fm4((self).f8, p0, (d_2923_v22_)[default__.safeIndex(291, (d_2923_v22_).length(0))], globalState)
                nw483_[int(13)] = d_2899_v1_
                nw483_[int(14)] = d_2899_v1_
                nw483_[int(15)] = d_2899_v1_
                nw483_[int(16)] = d_2899_v1_
                nw483_[int(17)] = d_2899_v1_
                nw483_[int(18)] = d_2899_v1_
                nw483_[int(19)] = d_2899_v1_
                nw483_[int(20)] = d_2899_v1_
                nw483_[int(21)] = d_2899_v1_
                nw483_[int(22)] = d_2899_v1_
                nw483_[int(23)] = d_2899_v1_
                d_2932_v31_ = nw483_
                d_2933_v32_: D12
                d_2933_v32_ = D12_DC37(len(d_2931_v30_), d_2932_v31_, d_2899_v1_)
                d_2934_v33_: D13
                d_2934_v33_ = D13_DC40(d_2933_v32_, p0, d_2910_i2_, d_2910_i2_)
                (globalState).f7 = (default__.safeModuloInt(len(d_2907_v7_), (d_2934_v33_).cf69)) - (d_2910_i2_)
            d_2935_v34_: bool
            d_2935_v34_ = False
            d_2935_v34_ = (self).f8
        if ((p2) * (-824)) == ((0) - ((p2) - (p2))):
            d_2936_v35_: bool
            d_2936_v35_ = True
            d_2936_v35_ = p0
            d_2937_v36_: _dafny.Array
            nw484_ = _dafny.Array(_dafny.CodePoint('D'), 24)
            d_2937_v36_ = nw484_
            index500_ = default__.safeIndex(597, (d_2937_v36_).length(0))
            (d_2937_v36_)[index500_] = d_2899_v1_
            d_2938_v37_: _dafny.Map
            d_2938_v37_ = _dafny.Map({False: default__.fm25(p2, p1, p2, globalState)})
            arr12_ = (d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]
            index501_ = default__.safeIndex(715, ((d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]).length(0))
            arr12_[index501_] = ((d_2938_v37_)[p0] if (p0) in (d_2938_v37_) else d_2900_v2_)
            d_2939_v38_: _dafny.Array
            def lambda134_(d_2940_p2_):
                def lambda135_(d_2941_i3_):
                    return default__.safeDivisionInt(d_2941_i3_, d_2940_p2_)

                return lambda135_

            init75_ = lambda134_(p2)
            nw485_ = _dafny.Array(None, 17)
            for i0_75_ in range(nw485_.length(0)):
                nw485_[i0_75_] = init75_(i0_75_)
            d_2939_v38_ = nw485_
            index502_ = default__.safeIndex(921, (d_2939_v38_).length(0))
            (d_2939_v38_)[index502_] = (0) - ((p2) - (p2))
            d_2942_v39_: _dafny.Map
            d_2942_v39_ = _dafny.Map({p1: p0})
            d_2943_v40_: C9
            nw486_ = C9()
            nw486_.ctor__(p0, ((d_2942_v39_)[d_2936_v35_] if (d_2936_v35_) in (d_2942_v39_) else not(p1)))
            d_2943_v40_ = nw486_
            d_2944_v41_: _dafny.Map
            d_2944_v41_ = _dafny.Map({p2: d_2943_v40_})
            d_2945_v42_: _dafny.Array
            nw487_ = _dafny.Array(None, 13)
            nw487_[int(0)] = d_2943_v40_
            nw487_[int(1)] = d_2943_v40_
            nw487_[int(2)] = d_2943_v40_
            nw487_[int(3)] = d_2943_v40_
            nw487_[int(4)] = d_2943_v40_
            nw487_[int(5)] = d_2943_v40_
            nw487_[int(6)] = ((d_2944_v41_)[p2] if (p2) in (d_2944_v41_) else d_2943_v40_)
            nw487_[int(7)] = d_2943_v40_
            nw487_[int(8)] = d_2943_v40_
            nw487_[int(9)] = d_2943_v40_
            nw487_[int(10)] = d_2943_v40_
            nw487_[int(11)] = d_2943_v40_
            nw487_[int(12)] = d_2943_v40_
            d_2945_v42_ = nw487_
            index503_ = default__.safeIndex(597, (d_2937_v36_).length(0))
            arr13_ = (d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]
            index504_ = default__.safeIndex(715, ((d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]).length(0))
            index505_ = default__.safeIndex(921, (d_2939_v38_).length(0))
            rhs490_ = d_2899_v1_
            rhs491_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adr"))
            rhs492_ = (0) - (p2)
            rhs493_ = (D22_DC65(d_2945_v42_)).cf111
            lhs352_ = d_2937_v36_
            lhs353_ = default__.safeIndex(597, (d_2937_v36_).length(0))
            lhs354_ = (d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]
            lhs355_ = default__.safeIndex(715, ((d_2898_v0_)[default__.safeIndex(198, (d_2898_v0_).length(0))]).length(0))
            lhs356_ = d_2939_v38_
            lhs357_ = default__.safeIndex(921, (d_2939_v38_).length(0))
            lhs352_[lhs353_] = rhs490_
            lhs354_[lhs355_] = rhs491_
            lhs356_[lhs357_] = rhs492_
            d_2945_v42_ = rhs493_
            (globalState).f7 = len(default__.fm62(d_2907_v7_, globalState))
            source47_ = default__.fm63(globalState)
            if source47_.is_DC17:
                d_2946___mcc_h0_ = source47_.cf28
                d_2947___mcc_h1_ = source47_.cf29
                d_2948_cf29_ = d_2947___mcc_h1_
                d_2949_cf28_ = d_2946___mcc_h0_
                d_2907_v7_ = d_2907_v7_
                (globalState).f7 = 582
                d_2936_v35_ = d_2936_v35_
                d_2950_v43_: D2
                d_2950_v43_ = D2_DC6(d_2899_v1_)
                d_2950_v43_ = d_2950_v43_
            elif source47_.is_DC16:
                d_2951___mcc_h2_ = source47_.cf27
                d_2952_cf27_ = d_2951___mcc_h2_
                d_2953_v44_: _dafny.Array
                nw488_ = _dafny.Array(None, 8)
                nw488_[int(0)] = p1
                nw488_[int(1)] = (self).f8
                nw488_[int(2)] = not((d_2943_v40_).fm6(p2, len(d_2906_v6_), p2, p0, globalState))
                nw488_[int(3)] = p0
                nw488_[int(4)] = False
                nw488_[int(5)] = (self).f8
                nw488_[int(6)] = (d_2943_v40_).f17
                nw488_[int(7)] = (self).f9
                d_2953_v44_ = nw488_
                d_2954_v45_: D0
                d_2954_v45_ = D0_DC0(d_2953_v44_)
                d_2955_v46_: _dafny.Array
                nw489_ = _dafny.Array(None, 29)
                nw489_[int(0)] = d_2953_v44_
                nw489_[int(1)] = d_2953_v44_
                nw489_[int(2)] = d_2953_v44_
                nw489_[int(3)] = (d_2954_v45_).cf0
                nw489_[int(4)] = d_2953_v44_
                nw489_[int(5)] = d_2953_v44_
                nw489_[int(6)] = d_2953_v44_
                nw489_[int(7)] = d_2953_v44_
                nw489_[int(8)] = d_2953_v44_
                nw489_[int(9)] = d_2953_v44_
                nw489_[int(10)] = d_2953_v44_
                nw489_[int(11)] = d_2953_v44_
                nw489_[int(12)] = d_2953_v44_
                nw489_[int(13)] = (D0_DC2(d_2936_v35_, d_2953_v44_)).cf3
                nw489_[int(14)] = d_2953_v44_
                nw489_[int(15)] = d_2953_v44_
                nw489_[int(16)] = d_2953_v44_
                nw489_[int(17)] = d_2953_v44_
                nw489_[int(18)] = d_2953_v44_
                nw489_[int(19)] = d_2953_v44_
                nw489_[int(20)] = d_2953_v44_
                nw489_[int(21)] = d_2953_v44_
                nw489_[int(22)] = d_2953_v44_
                nw489_[int(23)] = d_2953_v44_
                nw489_[int(24)] = d_2953_v44_
                nw489_[int(25)] = d_2953_v44_
                nw489_[int(26)] = d_2953_v44_
                nw489_[int(27)] = d_2953_v44_
                nw489_[int(28)] = d_2953_v44_
                d_2955_v46_ = nw489_
                d_2956_v47_: D17
                d_2956_v47_ = D17_DC50(d_2955_v46_)
                pat_let_tv75_ = d_2955_v46_
                pat_let_tv76_ = d_2955_v46_
                pat_let_tv77_ = d_2955_v46_
                d_2957_v48_: _dafny.Array
                nw490_ = _dafny.Array(None, 19)
                nw490_[int(0)] = d_2956_v47_
                def iife210_(_pat_let70_0):
                    def iife211_(d_2958_dt__update__tmp_h0_):
                        def iife212_(_pat_let71_0):
                            def iife213_(d_2959_dt__update_hcf86_h0_):
                                return D17_DC50(d_2959_dt__update_hcf86_h0_)
                            return iife213_(_pat_let71_0)
                        return iife212_(pat_let_tv75_)
                    return iife211_(_pat_let70_0)
                nw490_[int(1)] = iife210_(d_2956_v47_)
                nw490_[int(2)] = d_2956_v47_
                nw490_[int(3)] = d_2956_v47_
                nw490_[int(4)] = (D17_DC50(d_2955_v46_) if p0 else D17_DC50(d_2955_v46_))
                nw490_[int(5)] = d_2956_v47_
                nw490_[int(6)] = d_2956_v47_
                nw490_[int(7)] = d_2956_v47_
                nw490_[int(8)] = d_2956_v47_
                nw490_[int(9)] = d_2956_v47_
                nw490_[int(10)] = d_2956_v47_
                nw490_[int(11)] = d_2956_v47_
                nw490_[int(12)] = D17_DC50(d_2955_v46_)
                nw490_[int(13)] = D17_DC50(d_2955_v46_)
                nw490_[int(14)] = d_2956_v47_
                nw490_[int(15)] = d_2956_v47_
                def iife214_(_pat_let72_0):
                    def iife215_(d_2960_dt__update__tmp_h1_):
                        def iife216_(_pat_let73_0):
                            def iife217_(d_2961_dt__update_hcf86_h1_):
                                return D17_DC50(d_2961_dt__update_hcf86_h1_)
                            return iife217_(_pat_let73_0)
                        return iife216_(pat_let_tv76_)
                    return iife215_(_pat_let72_0)
                nw490_[int(16)] = iife214_(d_2956_v47_)
                def iife218_(_pat_let74_0):
                    def iife219_(d_2962_dt__update__tmp_h2_):
                        def iife220_(_pat_let75_0):
                            def iife221_(d_2963_dt__update_hcf86_h2_):
                                return D17_DC50(d_2963_dt__update_hcf86_h2_)
                            return iife221_(_pat_let75_0)
                        return iife220_(pat_let_tv77_)
                    return iife219_(_pat_let74_0)
                nw490_[int(17)] = iife218_(d_2956_v47_)
                nw490_[int(18)] = D17_DC50(d_2955_v46_)
                d_2957_v48_ = nw490_
                index506_ = default__.safeIndex(65, (d_2957_v48_).length(0))
                (d_2957_v48_)[index506_] = d_2956_v47_
                index507_ = default__.safeIndex(65, (d_2957_v48_).length(0))
                rhs494_ = (d_2937_v36_)[default__.safeIndex(597, (d_2937_v36_).length(0))]
                rhs495_ = (d_2908_v8_) + (d_2908_v8_)
                rhs496_ = d_2956_v47_
                lhs358_ = d_2957_v48_
                lhs359_ = default__.safeIndex(65, (d_2957_v48_).length(0))
                d_2899_v1_ = rhs494_
                d_2908_v8_ = rhs495_
                lhs358_[lhs359_] = rhs496_
                d_2964_v49_: _dafny.MultiSet
                d_2964_v49_ = _dafny.MultiSet([p2, p2])
                index508_ = default__.safeIndex(921, (d_2939_v38_).length(0))
                rhs497_ = default__.fm18(p2, d_2906_v6_, (d_2900_v2_) != (_dafny.SeqWithoutIsStrInference([(d_2937_v36_)[default__.safeIndex(597, (d_2937_v36_).length(0))] for d_2965_i4_ in range(default__.abs(-547))])), globalState)
                rhs498_ = (d_2943_v40_).fm15(globalState)
                lhs360_ = d_2939_v38_
                lhs361_ = default__.safeIndex(921, (d_2939_v38_).length(0))
                d_2964_v49_ = rhs497_
                lhs360_[lhs361_] = rhs498_
                d_2966_v50_: D5
                d_2966_v50_ = D5_DC17(p1, d_2899_v1_)
                d_2936_v35_ = (not((d_2966_v50_).cf28) if not(p0) else (d_2943_v40_).f17)
                (globalState).f0 = (0) - ((0) - (p2))
            elif True:
                d_2967___mcc_h3_ = source47_.cf30
                d_2968_cf30_ = d_2967___mcc_h3_
                d_2969_v51_: D1
                d_2969_v51_ = D1_DC3(p2)
                d_2970_v52_: D1
                d_2970_v52_ = D1_DC5(d_2969_v51_)
                d_2971_v53_: _dafny.Set
                d_2971_v53_ = _dafny.Set({d_2970_v52_, d_2970_v52_, d_2970_v52_})
                d_2971_v53_ = (d_2971_v53_) - (d_2971_v53_)
                d_2972_v54_: _dafny.Array
                nw491_ = _dafny.Array(False, 3)
                d_2972_v54_ = nw491_
                index509_ = default__.safeIndex(822, (d_2972_v54_).length(0))
                (d_2972_v54_)[index509_] = ((0) - (p2)) == ((d_2939_v38_)[default__.safeIndex(921, (d_2939_v38_).length(0))])
                index510_ = default__.safeIndex(822, (d_2972_v54_).length(0))
                (d_2972_v54_)[index510_] = not((d_2943_v40_).f17)
                d_2973_v55_: C3
                nw492_ = C3()
                nw492_.ctor__(p0, False)
                d_2973_v55_ = nw492_
                d_2974_v56_: C9
                nw493_ = C9()
                nw493_.ctor__(((d_2939_v38_)[default__.safeIndex(921, (d_2939_v38_).length(0))]) > ((d_2939_v38_)[default__.safeIndex(921, (d_2939_v38_).length(0))]), (d_2972_v54_)[default__.safeIndex(822, (d_2972_v54_).length(0))])
                d_2974_v56_ = nw493_
            d_2975_v57_: _dafny.Seq
            d_2975_v57_ = _dafny.SeqWithoutIsStrInference([p2])
            d_2976_v58_: _dafny.Map
            d_2976_v58_ = _dafny.Map({len(d_2975_v57_): (d_2939_v38_)[default__.safeIndex(921, (d_2939_v38_).length(0))]})
            d_2976_v58_ = (d_2976_v58_).set((p2) + (p2), p2)
        elif True:
            d_2977_v59_: _dafny.Seq
            d_2977_v59_ = _dafny.SeqWithoutIsStrInference([-513, p2])
            d_2978_v60_: _dafny.MultiSet
            d_2978_v60_ = _dafny.MultiSet([(self).f9, (self).f9])
            d_2979_v61_: _dafny.Array
            nw494_ = _dafny.Array(None, 4)
            nw494_[int(0)] = (d_2977_v59_)[default__.safeIndex(p2, len(d_2977_v59_))]
            nw494_[int(1)] = ((d_2978_v60_)[p0] if (p0) in (d_2978_v60_) else p2)
            nw494_[int(2)] = p2
            nw494_[int(3)] = default__.safeModuloInt(942, (0) - (p2))
            d_2979_v61_ = nw494_
            d_2980_v62_: _dafny.Map
            d_2980_v62_ = _dafny.Map({p1: (self).f8})
            d_2981_v63_: _dafny.MultiSet
            d_2981_v63_ = _dafny.MultiSet([d_2899_v1_, d_2899_v1_])
            d_2982_v64_: _dafny.Array
            nw495_ = _dafny.Array(_dafny.CodePoint('D'), 18)
            d_2982_v64_ = nw495_
            d_2983_v65_: D12
            d_2983_v65_ = D12_DC37((0) - (p2), d_2982_v64_, d_2899_v1_)
            d_2984_v66_: _dafny.Array
            nw496_ = _dafny.Array(None, 28)
            nw496_[int(0)] = (self).f8
            nw496_[int(1)] = (self).f8
            nw496_[int(2)] = (self).f8
            nw496_[int(3)] = True
            nw496_[int(4)] = (self).f9
            nw496_[int(5)] = p1
            nw496_[int(6)] = False
            nw496_[int(7)] = p1
            nw496_[int(8)] = p0
            nw496_[int(9)] = p0
            nw496_[int(10)] = p1
            nw496_[int(11)] = (self).f9
            nw496_[int(12)] = (self).f9
            nw496_[int(13)] = p0
            nw496_[int(14)] = (self).f9
            nw496_[int(15)] = (self).f8
            nw496_[int(16)] = (self).f8
            nw496_[int(17)] = p1
            nw496_[int(18)] = False
            nw496_[int(19)] = (self).f8
            nw496_[int(20)] = p1
            nw496_[int(21)] = p0
            nw496_[int(22)] = p1
            nw496_[int(23)] = False
            nw496_[int(24)] = (self).f9
            nw496_[int(25)] = p0
            nw496_[int(26)] = p1
            nw496_[int(27)] = (self).f8
            d_2984_v66_ = nw496_
            d_2985_v67_: C10
            nw497_ = C10()
            nw497_.ctor__((d_2978_v60_).cardinality, p1)
            d_2985_v67_ = nw497_
            d_2986_v68_: _dafny.Map
            d_2986_v68_ = _dafny.Map({d_2984_v66_: d_2985_v67_})
            d_2987_v69_: _dafny.MultiSet
            d_2987_v69_ = _dafny.MultiSet([(0) - (len(d_2986_v68_))])
            d_2988_v70_: _dafny.Map
            d_2988_v70_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_2985_v67_).f15 for d_2989_i5_ in range(default__.abs(160))]): len(d_2980_v62_)})
            d_2990_v71_: _dafny.Set
            d_2990_v71_ = _dafny.Set({d_2980_v62_})
            d_2991_v72_: _dafny.Map
            d_2991_v72_ = _dafny.Map({(d_2985_v67_).fm9(d_2900_v2_, d_2990_v71_, globalState): p2})
            nw498_ = _dafny.Array(None, 29)
            nw498_[int(0)] = 600
            nw498_[int(1)] = p2
            nw498_[int(2)] = (p2) * ((0) - (p2))
            nw498_[int(3)] = p2
            nw498_[int(4)] = p2
            nw498_[int(5)] = p2
            nw498_[int(6)] = (0) - (((d_2978_v60_)[p0] if (p0) in (d_2978_v60_) else default__.fm1(globalState)))
            nw498_[int(7)] = p2
            nw498_[int(8)] = (0) - (default__.safeModuloInt((0) - (p2), p2))
            nw498_[int(9)] = p2
            nw498_[int(10)] = p2
            nw498_[int(11)] = default__.safeModuloInt(p2, (self).fm5(p2, globalState))
            nw498_[int(12)] = p2
            nw498_[int(13)] = (p2) * (p2)
            nw498_[int(14)] = p2
            nw498_[int(15)] = p2
            nw498_[int(16)] = p2
            nw498_[int(17)] = p2
            nw498_[int(18)] = len(_dafny.SeqWithoutIsStrInference([d_2980_v62_, _dafny.Map({(self).f8: (self).f8}), d_2980_v62_]))
            nw498_[int(19)] = ((d_2981_v63_)[d_2899_v1_] if (d_2899_v1_) in (d_2981_v63_) else (d_2977_v59_)[default__.safeIndex(len(d_2900_v2_), len(d_2977_v59_))])
            nw498_[int(20)] = (p2) + ((0) - (p2))
            nw498_[int(21)] = p2
            nw498_[int(22)] = (D13_DC40(d_2983_v65_, p1, ((d_2906_v6_)[(self).f8] if ((self).f8) in (d_2906_v6_) else p2), p2)).cf70
            nw498_[int(23)] = p2
            nw498_[int(24)] = (p2) + ((d_2987_v69_).cardinality)
            nw498_[int(25)] = (p2) * (p2)
            nw498_[int(26)] = (0) - (len(d_2988_v70_))
            nw498_[int(27)] = (p2) * (p2)
            nw498_[int(28)] = len(d_2991_v72_)
            d_2979_v61_ = nw498_
            index511_ = default__.safeIndex(484, (d_2984_v66_).length(0))
            (d_2984_v66_)[index511_] = p0
            index512_ = default__.safeIndex(484, (d_2984_v66_).length(0))
            (d_2984_v66_)[index512_] = (d_2908_v8_)[default__.safeIndex((0) - ((0) - (p2)), len(d_2908_v8_))]
            index513_ = default__.safeIndex(484, (d_2984_v66_).length(0))
            (d_2984_v66_)[index513_] = not((d_2985_v67_).f16)
            d_2992_v73_: _dafny.Map
            d_2992_v73_ = _dafny.Map({(d_2984_v66_)[default__.safeIndex(484, (d_2984_v66_).length(0))]: _dafny.CodePoint('s')})
            d_2993_v74_: D2
            d_2993_v74_ = D2_DC7((d_2985_v67_).f15, (d_2985_v67_).f16, len(d_2992_v73_))
            (globalState).f7 = (d_2993_v74_).cf11
            if not((d_2985_v67_).f16):
                index514_ = default__.safeIndex(360, (d_2979_v61_).length(0))
                (d_2979_v61_)[index514_] = (d_2985_v67_).f15
                d_2994_v75_: _dafny.Seq
                d_2994_v75_ = _dafny.SeqWithoutIsStrInference([d_2906_v6_])
                index515_ = default__.safeIndex(360, (d_2979_v61_).length(0))
                (d_2979_v61_)[index515_] = (0) - (len((((d_2994_v75_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_2985_v67_).f15 for d_2995_i6_ in range(default__.abs(120))])), len(d_2994_v75_))]).set((d_2984_v66_)[default__.safeIndex(484, (d_2984_v66_).length(0))], (0) - ((d_2985_v67_).f15)) if (self).f9 else (d_2906_v6_) | (default__.fm51(globalState)))))
                (globalState).f0 = 389
                index516_ = default__.safeIndex(360, (d_2979_v61_).length(0))
                (d_2979_v61_)[index516_] = ((p2) - (p2)) - (((d_2987_v69_).cardinality) - ((self).fm5((d_2985_v67_).f15, globalState)))
                d_2996_v76_: C6
                nw499_ = C6()
                nw499_.ctor__()
                d_2996_v76_ = nw499_
                d_2984_v66_ = d_2984_v66_
            elif True:
                d_2997_v77_: _dafny.Map
                d_2997_v77_ = _dafny.Map({d_2900_v2_: d_2899_v1_})
                d_2997_v77_ = (d_2997_v77_).set(d_2900_v2_, d_2899_v1_)
                d_2998_v78_: _dafny.Set
                d_2998_v78_ = _dafny.Set({d_2982_v64_})
                d_2999_v79_: _dafny.MultiSet
                d_2999_v79_ = _dafny.MultiSet([d_2998_v78_, d_2998_v78_, _dafny.Set({d_2982_v64_, d_2982_v64_}), d_2998_v78_, _dafny.Set({d_2982_v64_, d_2982_v64_})])
                d_3000_v80_: _dafny.Map
                d_3000_v80_ = _dafny.Map({d_2908_v8_: (d_2984_v66_)[default__.safeIndex(484, (d_2984_v66_).length(0))]})
                d_3001_v82_: _dafny.Array
                nw500_ = _dafny.Array(D1.default()(), 27)
                d_3001_v82_ = nw500_
                d_3002_v83_: _dafny.Seq
                d_3002_v83_ = _dafny.SeqWithoutIsStrInference([d_3001_v82_, d_3001_v82_, d_3001_v82_])
                d_3003_v84_: D8
                d_3003_v84_ = D8_DC27(p0, d_3002_v83_, d_2977_v59_, (d_2985_v67_).f15)
                d_3004_v85_: _dafny.Map
                d_3004_v85_ = _dafny.Map({(0) - (len(d_2980_v62_)): (d_3003_v84_).cf45})
                def iife222_():
                    coll70_ = _dafny.Set()
                    compr_70_: _dafny.Seq
                    for compr_70_ in (d_3000_v80_).keys.Elements:
                        d_3005_v81_: _dafny.Seq = compr_70_
                        if (d_3005_v81_) in (d_3000_v80_):
                            coll70_ = coll70_.union(_dafny.Set([d_3005_v81_]))
                    return _dafny.Set(coll70_)
                (globalState).f0 = ((d_2999_v79_)[d_2998_v78_] if (d_2998_v78_) in (d_2999_v79_) else (len(iife222_()
                )) + ((0) - (len(d_3004_v85_))))
                index517_ = default__.safeIndex(484, (d_2984_v66_).length(0))
                (d_2984_v66_)[index517_] = False
                index518_ = default__.safeIndex(484, (d_2984_v66_).length(0))
                (d_2984_v66_)[index518_] = (self).f9
                d_3006_v86_: D0
                d_3006_v86_ = D0_DC1(d_2900_v2_)
                index519_ = default__.safeIndex(614, (d_2979_v61_).length(0))
                (d_2979_v61_)[index519_] = (0) - ((d_2977_v59_)[default__.safeIndex(len((d_3006_v86_).cf1), len(d_2977_v59_))])
                index520_ = default__.safeIndex(614, (d_2979_v61_).length(0))
                (d_2979_v61_)[index520_] = default__.fm1(globalState)
        source48_ = default__.fm64(globalState)
        if source48_.is_DC31:
            d_3007___mcc_h4_ = source48_.cf50
            d_3008___mcc_h5_ = source48_.cf51
            d_3009___mcc_h6_ = source48_.cf52
            d_3010_cf52_ = d_3009___mcc_h6_
            d_3011_cf51_ = d_3008___mcc_h5_
            d_3012_cf50_ = d_3007___mcc_h4_
            d_3013_v87_: bool
            out83_: bool
            out83_ = (self).m1(p2, globalState)
            d_3013_v87_ = out83_
            d_3014_v88_: _dafny.Seq
            d_3014_v88_ = _dafny.SeqWithoutIsStrInference([d_3010_cf52_, d_3010_cf52_, p2])
            d_3015_v89_: _dafny.Map
            d_3015_v89_ = _dafny.Map({d_3012_cf50_: d_2906_v6_})
            d_3016_v90_: _dafny.Map
            d_3016_v90_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihsbrh")): d_2906_v6_})
            d_3017_v91_: _dafny.Array
            nw501_ = _dafny.Array(None, 17)
            nw501_[int(0)] = (d_2906_v6_) | (_dafny.Map({d_3013_v87_: (0) - (len(d_3014_v88_))}))
            nw501_[int(1)] = (((d_3015_v89_)[(self).f9] if ((self).f9) in (d_3015_v89_) else d_2906_v6_)) | (d_2906_v6_)
            nw501_[int(2)] = d_2906_v6_
            nw501_[int(3)] = (_dafny.Map({(self).f8: len(d_2900_v2_)})) | (d_2906_v6_)
            nw501_[int(4)] = d_2906_v6_
            nw501_[int(5)] = _dafny.Map({p0: 541})
            nw501_[int(6)] = d_2906_v6_
            nw501_[int(7)] = (d_2906_v6_).set(d_3013_v87_, p2)
            nw501_[int(8)] = (d_2906_v6_) | (d_2906_v6_)
            nw501_[int(9)] = d_2906_v6_
            nw501_[int(10)] = ((d_2906_v6_).set(not(d_3013_v87_), p2)) | (d_2906_v6_)
            nw501_[int(11)] = (((d_3016_v90_)[d_2900_v2_] if (d_2900_v2_) in (d_3016_v90_) else d_2906_v6_)) | (d_2906_v6_)
            nw501_[int(12)] = d_2906_v6_
            nw501_[int(13)] = _dafny.Map({not(d_3011_cf51_): d_3010_cf52_})
            nw501_[int(14)] = d_2906_v6_
            nw501_[int(15)] = _dafny.Map({not(False): p2})
            nw501_[int(16)] = d_2906_v6_
            d_3017_v91_ = nw501_
            d_3018_v92_: _dafny.Array
            def lambda136_(d_3019_i7_):
                return (self).f8

            init76_ = lambda136_
            nw502_ = _dafny.Array(None, 7)
            for i0_76_ in range(nw502_.length(0)):
                nw502_[i0_76_] = init76_(i0_76_)
            d_3018_v92_ = nw502_
            d_3020_v93_: T0
            nw503_ = C9()
            nw503_.ctor__(True, p1)
            d_3020_v93_ = nw503_
            d_3021_v94_: _dafny.MultiSet
            d_3021_v94_ = _dafny.MultiSet([d_3020_v93_, d_3020_v93_])
            index521_ = default__.safeIndex(601, (d_3018_v92_).length(0))
            (d_3018_v92_)[index521_] = default__.fm0((d_3021_v94_).cardinality, globalState)
            d_3022_v95_: _dafny.MultiSet
            d_3022_v95_ = _dafny.MultiSet([d_2907_v7_, d_2907_v7_])
            index522_ = default__.safeIndex(601, (d_3018_v92_).length(0))
            rhs499_ = d_3017_v91_
            rhs500_ = (_dafny.MultiSet([_dafny.Set({d_3010_cf52_})])).issubset((_dafny.MultiSet([_dafny.Set({d_3010_cf52_})]) if p1 else d_3022_v95_))
            lhs362_ = d_3018_v92_
            lhs363_ = default__.safeIndex(601, (d_3018_v92_).length(0))
            d_3017_v91_ = rhs499_
            lhs362_[lhs363_] = rhs500_
            (globalState).f0 = p2
            if (d_3010_cf52_) > (p2):
                d_3023_v96_: C5
                nw504_ = C5()
                nw504_.ctor__()
                d_3023_v96_ = nw504_
                d_3023_v96_ = (d_3023_v96_ if (d_2900_v2_) < (_dafny.SeqWithoutIsStrInference([d_2899_v1_ for d_3024_i8_ in range(default__.abs(431))])) else (d_3023_v96_ if not((self).f8) else d_3023_v96_))
                d_3025_v97_: C4
                nw505_ = C4()
                nw505_.ctor__()
                d_3025_v97_ = nw505_
                d_3026_v98_: C4
                d_3026_v98_ = d_3025_v97_
                d_3025_v97_ = (d_3026_v98_)
                d_3027_v99_: _dafny.Map
                d_3027_v99_ = _dafny.Map({(self).f8: p0})
                d_3028_v100_: _dafny.Set
                d_3028_v100_ = _dafny.Set({p0, not(not((d_3020_v93_).f10)), ((d_3027_v99_)[False] if (False) in (d_3027_v99_) else (self).f8), (d_3018_v92_)[default__.safeIndex(601, (d_3018_v92_).length(0))]})
                d_3029_v101_: C12
                nw506_ = C12()
                nw506_.ctor__((d_2908_v8_)[default__.safeIndex(p2, len(d_2908_v8_))], (d_3018_v92_)[default__.safeIndex(601, (d_3018_v92_).length(0))], (len(d_3028_v100_)) < (d_3010_cf52_))
                d_3029_v101_ = nw506_
                d_3030_v102_: _dafny.Array
                def lambda137_(d_3031_v1_):
                    def lambda138_(d_3032_i9_):
                        return d_3031_v1_

                    return lambda138_

                init77_ = lambda137_(d_2899_v1_)
                nw507_ = _dafny.Array(None, 2)
                for i0_77_ in range(nw507_.length(0)):
                    nw507_[i0_77_] = init77_(i0_77_)
                d_3030_v102_ = nw507_
                index523_ = default__.safeIndex(929, (d_3030_v102_).length(0))
                (d_3030_v102_)[index523_] = default__.fm4(d_3013_v87_, (d_3018_v92_)[default__.safeIndex(601, (d_3018_v92_).length(0))], d_3013_v87_, globalState)
                d_3033_v103_: _dafny.Array
                nw508_ = _dafny.Array(None, 2)
                nw508_[int(0)] = len(d_2906_v6_)
                nw508_[int(1)] = (d_3010_cf52_) + (len(_dafny.Map({d_3010_cf52_: d_3010_cf52_})))
                d_3033_v103_ = nw508_
                index524_ = default__.safeIndex(866, (d_3033_v103_).length(0))
                (d_3033_v103_)[index524_] = p2
                index525_ = default__.safeIndex(601, (d_3018_v92_).length(0))
                index526_ = default__.safeIndex(929, (d_3030_v102_).length(0))
                index527_ = default__.safeIndex(866, (d_3033_v103_).length(0))
                rhs501_ = (d_3029_v101_).f11
                rhs502_ = d_2899_v1_
                rhs503_ = default__.safeModuloInt((0) - ((p2) * (len(d_2900_v2_))), (d_3010_cf52_) - (len(d_3014_v88_)))
                rhs504_ = (d_2907_v7_).intersection(d_2907_v7_)
                lhs364_ = d_3018_v92_
                lhs365_ = default__.safeIndex(601, (d_3018_v92_).length(0))
                lhs366_ = d_3030_v102_
                lhs367_ = default__.safeIndex(929, (d_3030_v102_).length(0))
                lhs368_ = d_3033_v103_
                lhs369_ = default__.safeIndex(866, (d_3033_v103_).length(0))
                lhs364_[lhs365_] = rhs501_
                lhs366_[lhs367_] = rhs502_
                lhs368_[lhs369_] = rhs503_
                d_2907_v7_ = rhs504_
                d_3012_cf50_ = p0
            elif True:
                d_3034_v104_: _dafny.MultiSet
                d_3034_v104_ = _dafny.MultiSet([d_3013_v87_])
                d_3035_v105_: _dafny.Seq
                d_3035_v105_ = _dafny.SeqWithoutIsStrInference([d_3034_v104_])
                d_3036_v106_: _dafny.Array
                nw509_ = _dafny.Array(None, 17)
                nw509_[int(0)] = d_3034_v104_
                nw509_[int(1)] = (_dafny.MultiSet(d_2908_v8_)) - (d_3034_v104_)
                nw509_[int(2)] = d_3034_v104_
                nw509_[int(3)] = _dafny.MultiSet(d_2908_v8_)
                nw509_[int(4)] = d_3034_v104_
                nw509_[int(5)] = d_3034_v104_
                nw509_[int(6)] = default__.fm41(globalState)
                nw509_[int(7)] = d_3034_v104_
                nw509_[int(8)] = (d_3035_v105_)[default__.safeIndex(106, len(d_3035_v105_))]
                nw509_[int(9)] = d_3034_v104_
                nw509_[int(10)] = (D5_DC16(d_3034_v104_)).cf27
                nw509_[int(11)] = (d_3034_v104_).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_3012_cf50_])))
                nw509_[int(12)] = d_3034_v104_
                nw509_[int(13)] = _dafny.MultiSet([(d_3020_v93_).f10, (self).f8])
                nw509_[int(14)] = ((_dafny.MultiSet([not(p1)])).set(d_3012_cf50_, default__.abs(p2))) | (d_3034_v104_)
                nw509_[int(15)] = d_3034_v104_
                nw509_[int(16)] = d_3034_v104_
                d_3036_v106_ = nw509_
                index528_ = default__.safeIndex(125, (d_3036_v106_).length(0))
                (d_3036_v106_)[index528_] = _dafny.MultiSet([d_3011_cf51_])
                d_3037_v107_: D0
                d_3037_v107_ = D0_DC0(d_3018_v92_)
                index529_ = default__.safeIndex(125, (d_3036_v106_).length(0))
                index530_ = default__.safeIndex(601, (d_3018_v92_).length(0))
                rhs505_ = d_3034_v104_
                rhs506_ = (self).f8
                rhs507_ = (d_3020_v93_).fm7(not(d_3013_v87_), globalState)
                rhs508_ = d_3037_v107_
                lhs370_ = d_3036_v106_
                lhs371_ = default__.safeIndex(125, (d_3036_v106_).length(0))
                lhs372_ = d_3018_v92_
                lhs373_ = default__.safeIndex(601, (d_3018_v92_).length(0))
                lhs370_[lhs371_] = rhs505_
                d_3013_v87_ = rhs506_
                lhs372_[lhs373_] = rhs507_
                d_3037_v107_ = rhs508_
                d_3038_v108_: _dafny.Map
                d_3038_v108_ = _dafny.Map({(d_3010_cf52_) + (p2): d_2899_v1_})
                d_3038_v108_ = (d_3038_v108_).set(d_3010_cf52_, d_2899_v1_)
                d_3039_v109_: C5
                nw510_ = C5()
                nw510_.ctor__()
                d_3039_v109_ = nw510_
                (globalState).f0 = (d_3010_cf52_) * (len((d_2900_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))))
                (globalState).f1 = p2
        elif True:
            d_3040___mcc_h7_ = source48_.cf49
            d_3041_cf49_ = d_3040___mcc_h7_
            d_3042_v110_: _dafny.Array
            nw511_ = _dafny.Array(False, 12)
            d_3042_v110_ = nw511_
            nw512_ = _dafny.Array(False, 3)
            d_3042_v110_ = nw512_
            d_3043_v111_: bool
            d_3043_v111_ = True
            index531_ = default__.safeIndex(906, (d_3042_v110_).length(0))
            (d_3042_v110_)[index531_] = False
            d_3044_v112_: _dafny.MultiSet
            d_3044_v112_ = _dafny.MultiSet([(0) - (p2)])
            index532_ = default__.safeIndex(906, (d_3042_v110_).length(0))
            rhs509_ = not((self).f8)
            rhs510_ = (d_3044_v112_).isdisjoint(d_3044_v112_)
            rhs511_ = (d_2908_v8_)[default__.safeIndex(len((d_2900_v2_) + (d_2900_v2_)), len(d_2908_v8_))]
            rhs512_ = p2
            lhs374_ = d_3042_v110_
            lhs375_ = default__.safeIndex(906, (d_3042_v110_).length(0))
            lhs376_ = globalState
            d_3043_v111_ = rhs509_
            lhs374_[lhs375_] = rhs510_
            d_3043_v111_ = rhs511_
            lhs376_.f0 = rhs512_
            index533_ = default__.safeIndex(906, (d_3042_v110_).length(0))
            (d_3042_v110_)[index533_] = d_3043_v111_
            d_3045_v113_: _dafny.Set
            d_3045_v113_ = _dafny.Set({d_3042_v110_})
            d_3046_v115_: D10
            d_3046_v115_ = D10_DC30(d_3041_cf49_)
            d_3047_v116_: _dafny.Seq
            d_3047_v116_ = _dafny.SeqWithoutIsStrInference([p2])
            pat_let_tv78_ = p0
            index534_ = default__.safeIndex(906, (d_3042_v110_).length(0))
            index535_ = default__.safeIndex(906, (d_3042_v110_).length(0))
            def iife224_():
                coll71_ = _dafny.Map()
                compr_71_: int
                for compr_71_ in _dafny.IntegerRange(744, 853):
                    d_3048_v114_: int = compr_71_
                    if ((744) <= (d_3048_v114_)) and ((d_3048_v114_) < (853)):
                        coll71_[(d_3048_v114_) - (len(d_2908_v8_))] = len(d_2900_v2_)
                return _dafny.Map(coll71_)
            def iife223_(_pat_let76_0):
                def iife225_(d_3049_dt__update__tmp_h3_):
                    def iife226_(_pat_let77_0):
                        def iife227_(d_3050_dt__update_hcf54_h0_):
                            return D11_DC33(d_3050_dt__update_hcf54_h0_, (d_3049_dt__update__tmp_h3_).cf55)
                        return iife227_(_pat_let77_0)
                    return iife226_(pat_let_tv78_)
                return iife225_(_pat_let76_0)
            rhs513_ = p1
            rhs514_ = (_dafny.Set({p0, (iife223_(D11_DC33((self).f9, (0) - (len(iife224_()
))))).cf54})) | ((_dafny.Set({True, p1, (d_3042_v110_)[default__.safeIndex(906, (d_3042_v110_).length(0))]})).intersection((d_3046_v115_).cf49))
            rhs515_ = (d_3045_v113_) - (d_3045_v113_)
            rhs516_ = p2
            rhs517_ = (len(d_3047_v116_)) != (p2)
            lhs377_ = d_3042_v110_
            lhs378_ = default__.safeIndex(906, (d_3042_v110_).length(0))
            lhs379_ = globalState
            lhs380_ = d_3042_v110_
            lhs381_ = default__.safeIndex(906, (d_3042_v110_).length(0))
            lhs377_[lhs378_] = rhs513_
            d_3041_cf49_ = rhs514_
            d_3045_v113_ = rhs515_
            lhs379_.f0 = rhs516_
            lhs380_[lhs381_] = rhs517_
        d_3051_v117_: bool
        d_3051_v117_ = False
        d_3052_v118_: _dafny.Set
        d_3052_v118_ = _dafny.Set({(self).f8, (self).f8})
        d_3051_v117_ = (d_3052_v118_).issubset(d_3052_v118_)

    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
