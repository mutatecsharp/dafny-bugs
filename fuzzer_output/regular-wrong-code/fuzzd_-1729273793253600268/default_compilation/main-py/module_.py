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
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(134, 865):
                d_0_v0_: int = compr_0_
                if ((134) <= (d_0_v0_)) and ((d_0_v0_) < (865)):
                    coll0_[default__.safeDivisionInt(d_0_v0_, 235)] = D1_DC5(_dafny.Set({False}))
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({653: False}))])).Elements:
                d_1_v1_: int = compr_1_
                if (d_1_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({653: False}))])):
                    coll1_[default__.safeDivisionInt(d_1_v1_, len(_dafny.Set({(0) - (len(_dafny.Set({True}))), (0) - (len(_dafny.Set({True}))), len(_dafny.Set({False, False})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2_i0_ in range(default__.abs(616))]))})))] = D1_DC5(_dafny.Set({True}))
            return _dafny.Map(coll1_)
        return _dafny.SeqWithoutIsStrInference([((0) - (len((D8_DC22(_dafny.Map({-702: True}))).cf41))) >= (len(_dafny.SeqWithoutIsStrInference([D8_DC25(True, _dafny.Map({False: 764}), _dafny.CodePoint('d'), iife0_()
, 350), D8_DC25(not(True), _dafny.Map({not(True): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kj")))}), _dafny.CodePoint('h'), iife1_()
, len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality, 123])))]))), (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_3_i1_ in range(default__.abs(-832))])])).issubset(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aphwvlvy"))])), not (True) or (True)])

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({True: not(False)})) | (_dafny.Map({False: True}))) | (_dafny.Map({True: True}))

    @staticmethod
    def fm2(p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "entkdvmve"))

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return D0_DC0(780)

    @staticmethod
    def fm4(p0, globalState):
        return default__.safeModuloInt(306, (0) - ((0) - ((0) - (default__.safeModuloInt(-211, len(_dafny.Map({True: False})))))))

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: str
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')])).Elements:
                d_4_v0_: str = compr_2_
                if (d_4_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')])):
                    coll2_[d_4_v0_] = (0) - ((0) - (-428))
            return _dafny.Map(coll2_)
        def iife3_():
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: str
                for compr_5_ in (_dafny.Map({_dafny.CodePoint('t'): _dafny.CodePoint('o')})).keys.Elements:
                    d_5_v1_: str = compr_5_
                    if (d_5_v1_) in (_dafny.Map({_dafny.CodePoint('t'): _dafny.CodePoint('o')})):
                        coll5_[d_5_v1_] = 748
                return _dafny.Map(coll5_)
            coll3_ = _dafny.Set()
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: str
                for compr_4_ in (_dafny.Map({_dafny.CodePoint('t'): _dafny.CodePoint('o')})).keys.Elements:
                    d_5_v1_: str = compr_4_
                    if (d_5_v1_) in (_dafny.Map({_dafny.CodePoint('t'): _dafny.CodePoint('o')})):
                        coll4_[d_5_v1_] = 748
                return _dafny.Map(coll4_)
            compr_3_: _dafny.Map
            for compr_3_ in (_dafny.Map({iife4_()
            : len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_6_i0_ in range(default__.abs(583))]))})).keys.Elements:
                d_7_v2_: _dafny.Map = compr_3_
                if (d_7_v2_) in (_dafny.Map({iife5_()
                : len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_6_i0_ in range(default__.abs(583))]))})):
                    coll3_ = coll3_.union(_dafny.Set([d_7_v2_]))
            return _dafny.Set(coll3_)
        return ((_dafny.Map({_dafny.CodePoint('m'): len(_dafny.SeqWithoutIsStrInference([(0) - (-129), -622]))})) | (_dafny.Map({_dafny.CodePoint('o'): 288}))) in ((_dafny.Set({_dafny.Map({_dafny.CodePoint('e'): len(_dafny.SeqWithoutIsStrInference([False]))}), iife2_()
        })) - (iife3_()
        ))

    @staticmethod
    def fm10(globalState):
        return D1_DC6(False, _dafny.CodePoint('f'), 469)

    @staticmethod
    def fm11(p0, globalState):
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: _dafny.Map
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): True}), _dafny.Map({True: False}), _dafny.Map({False: not(False)}), _dafny.Map({False: not(False)}), _dafny.Map({True: True})])).Elements:
                d_8_v0_: _dafny.Map = compr_6_
                if (d_8_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): True}), _dafny.Map({True: False}), _dafny.Map({False: not(False)}), _dafny.Map({False: not(False)}), _dafny.Map({True: True})])):
                    coll6_ = coll6_.union(_dafny.Set([d_8_v0_]))
            return _dafny.Set(coll6_)
        return (iife6_()
        ) - ((_dafny.Set({_dafny.Map({not(True): True}), _dafny.Map({False: True}), _dafny.Map({True: False}), (D3_DC10(_dafny.Map({True: True}))).cf20})) | (_dafny.Set({_dafny.Map({True: True}), _dafny.Map({False: not(True)})})))

    @staticmethod
    def fm14(globalState):
        return _dafny.Set({not(not(True))})

    @staticmethod
    def fm17(globalState):
        return ((_dafny.Set({True})) | (_dafny.Set({False, True}))).intersection((_dafny.Set({True, True})) - (_dafny.Set({True})))

    @staticmethod
    def fm18(p0, globalState):
        def iife7_():
            def iife9_():
                coll9_ = _dafny.Set()
                compr_9_: str
                for compr_9_ in (_dafny.Map({_dafny.CodePoint('w'): not(False)})).keys.Elements:
                    d_13_v0_: str = compr_9_
                    if (d_13_v0_) in (_dafny.Map({_dafny.CodePoint('w'): not(False)})):
                        coll9_ = coll9_.union(_dafny.Set([d_13_v0_]))
                return _dafny.Set(coll9_)
            coll7_ = _dafny.Set()
            def iife8_():
                coll8_ = _dafny.Set()
                compr_8_: str
                for compr_8_ in (_dafny.Map({_dafny.CodePoint('w'): not(False)})).keys.Elements:
                    d_11_v0_: str = compr_8_
                    if (d_11_v0_) in (_dafny.Map({_dafny.CodePoint('w'): not(False)})):
                        coll8_ = coll8_.union(_dafny.Set([d_11_v0_]))
                return _dafny.Set(coll8_)
            compr_7_: _dafny.Set
            for compr_7_ in (_dafny.SeqWithoutIsStrInference([iife8_()
            ])).Elements:
                d_12_v1_: _dafny.Set = compr_7_
                if (d_12_v1_) in (_dafny.SeqWithoutIsStrInference([iife9_()
                ])):
                    coll7_ = coll7_.union(_dafny.Set([d_12_v1_]))
            return _dafny.Set(coll7_)
        return _dafny.Map({(653) + (-319): (_dafny.SeqWithoutIsStrInference([493, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyeibuhk")))])) == (_dafny.SeqWithoutIsStrInference([(D24_DC71(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_9_i0_ in range(default__.abs(-137))])), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_10_i1_ in range(default__.abs(252))]), False)).cf133, -331, len(iife7_()
        ), len(_dafny.SeqWithoutIsStrInference([False]))]))})

    @staticmethod
    def fm19(p0, p1, p2, p3, globalState):
        return D3_DC10(_dafny.Map({False: False}))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        source0_ = D2_DC8(False, 551, len(_dafny.SeqWithoutIsStrInference([516 for d_14_i0_ in range(default__.abs(705))])))
        if source0_.is_DC8:
            d_15___mcc_h0_ = source0_.cf13
            d_16___mcc_h1_ = source0_.cf14
            d_17___mcc_h2_ = source0_.cf15
            d_18_cf15_ = d_17___mcc_h2_
            d_19_cf14_ = d_16___mcc_h1_
            d_20_cf13_ = d_15___mcc_h0_
            return _dafny.CodePoint('o')
        elif source0_.is_DC9:
            d_21___mcc_h3_ = source0_.cf16
            d_22___mcc_h4_ = source0_.cf17
            d_23___mcc_h5_ = source0_.cf18
            d_24___mcc_h6_ = source0_.cf19
            d_25_cf19_ = d_24___mcc_h6_
            d_26_cf18_ = d_23___mcc_h5_
            d_27_cf17_ = d_22___mcc_h4_
            d_28_cf16_ = d_21___mcc_h3_
            return _dafny.CodePoint('h')
        elif True:
            d_29___mcc_h7_ = source0_.cf12
            d_30_cf12_ = d_29___mcc_h7_
            if False:
                return _dafny.CodePoint('e')
            elif True:
                return _dafny.CodePoint('e')

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(736, 187):
                d_31_v0_: int = compr_10_
                if ((736) <= (d_31_v0_)) and ((d_31_v0_) < (187)):
                    coll10_[(d_31_v0_) + (708)] = True
            return _dafny.Map(coll10_)
        return _dafny.SeqWithoutIsStrInference([621, len((_dafny.SeqWithoutIsStrInference([D16_DC45(len(_dafny.SeqWithoutIsStrInference([True])), False, (0) - (-879), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thj"))), True), D16_DC45((_dafny.MultiSet([False, True, True, True])).cardinality, True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdngctndp"))), 579, True)]) if False else _dafny.SeqWithoutIsStrInference([D16_DC45(-429, not(True), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eiusfxqh")))])), len(_dafny.Map({False: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqyrbfsc"))))})), False), D16_DC45((0) - (len(_dafny.Map({True: len(iife10_()
)}))), True, len(_dafny.SeqWithoutIsStrInference([True, True])), -410, True)])))])

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        source1_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anfmose")): False})
        d_32___mcc_h0_ = source1_
        d_33_cf79_ = d_32___mcc_h0_
        return D6_DC16(-240, False)

    @staticmethod
    def fm23(p0, globalState):
        return (_dafny.Map({False: _dafny.CodePoint('w')})) | ((_dafny.Map({True: _dafny.CodePoint('r')})) | (_dafny.Map({True: _dafny.CodePoint('h')})))

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([730])) <= (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False])), (_dafny.MultiSet([140])).cardinality, 747, -170, 146]))

    @staticmethod
    def fm25(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_34_i0_ in range(default__.abs(459))])]) if False else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")) for d_35_i1_ in range(default__.abs(638))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbhlo"))]))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        return (_dafny.Set({-269})) | (_dafny.Set({774, (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_36_i0_ in range(default__.abs(416))]))), -437}))

    @staticmethod
    def fm27(p0, p1, globalState):
        return D1_DC6(False, _dafny.CodePoint('m'), (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_37_i0_ in range(default__.abs(872))]))) + (224))

    @staticmethod
    def fm28(globalState):
        return (D26_DC78(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False}))).cf149

    @staticmethod
    def fm29(p0, p1, p2, globalState):
        return D0_DC4(D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlky")), -639, (D27_DC82(len(_dafny.SeqWithoutIsStrInference([not(False)])), len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdlkrr")))})))).cf155))

    @staticmethod
    def fm30(globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: _dafny.Map
            for compr_11_ in (_dafny.MultiSet([_dafny.Map({False: False})])).Elements:
                d_38_v0_: _dafny.Map = compr_11_
                if (d_38_v0_) in (_dafny.MultiSet([_dafny.Map({False: False})])):
                    coll11_[d_38_v0_] = _dafny.Set({len(_dafny.Map({False: not(False)}))})
            return _dafny.Map(coll11_)
        def iife12_():
            coll12_ = _dafny.Set()
            compr_12_: int
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([357 for d_39_i0_ in range(default__.abs(881))])).Elements:
                d_40_v1_: int = compr_12_
                if (d_40_v1_) in (_dafny.SeqWithoutIsStrInference([357 for d_39_i0_ in range(default__.abs(881))])):
                    coll12_ = coll12_.union(_dafny.Set([default__.safeDivisionInt(d_40_v1_, (0) - ((0) - (len(_dafny.Map({D8_DC22(_dafny.Map({251: True})): True})))))]))
            return _dafny.Set(coll12_)
        return (iife11_()
        ) | ((_dafny.Map({_dafny.Map({True: False}): _dafny.Set({len(_dafny.Map({False: _dafny.CodePoint('s')}))})})) | (_dafny.Map({_dafny.Map({False: False}): iife12_()
        })))

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in (_dafny.SeqWithoutIsStrInference([693])).Elements:
                d_42_v0_: int = compr_13_
                if (d_42_v0_) in (_dafny.SeqWithoutIsStrInference([693])):
                    coll13_[(d_42_v0_) * (850)] = (0) - (-182)
            return _dafny.Map(coll13_)
        return D7_DC19(len((_dafny.Map({129: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_41_i0_ in range(default__.abs(246))])): len(_dafny.Map({True: 682}))}))})) | (iife13_()
)), False)

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return D2_DC9((0) - ((0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcoc")))) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_43_i0_ in range(default__.abs(771))])))))), (True) or (True), (_dafny.MultiSet([983])).cardinality, (0) - ((972) * (len(_dafny.SeqWithoutIsStrInference([False, True])))))

    @staticmethod
    def fm33(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([-985, len(_dafny.SeqWithoutIsStrInference([269, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_44_i0_ in range(default__.abs(641))]))]))])) | (_dafny.MultiSet([117, 286]))) - ((D10_DC30(_dafny.MultiSet([-804]))).cf57)

    @staticmethod
    def fm34(p0, p1, p2, p3, globalState):
        source2_ = D28_DC88(True, -846)
        if source2_.is_DC87:
            d_45___mcc_h0_ = source2_.cf165
            d_46___mcc_h1_ = source2_.cf166
            d_47___mcc_h2_ = source2_.cf167
            d_48_cf167_ = d_47___mcc_h2_
            d_49_cf166_ = d_46___mcc_h1_
            d_50_cf165_ = d_45___mcc_h0_
            return _dafny.Map({False: not(False)})
        elif source2_.is_DC88:
            d_51___mcc_h3_ = source2_.cf168
            d_52___mcc_h4_ = source2_.cf169
            d_53_cf169_ = d_52___mcc_h4_
            d_54_cf168_ = d_51___mcc_h3_
            return _dafny.Map({d_54_cf168_: not(d_54_cf168_)})
        elif True:
            d_55___mcc_h5_ = source2_.cf164
            d_56_cf164_ = d_55___mcc_h5_
            return (_dafny.Map({True: True})) | (_dafny.Map({True: False}))

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        source3_ = D15_DC42(_dafny.MultiSet([True]))
        if source3_.is_DC43:
            def iife14_():
                coll14_ = _dafny.Set()
                compr_14_: _dafny.Seq
                for compr_14_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('d')])).cardinality, 788])) for d_57_i1_ in range(default__.abs(-569))])), 871])): True})) for d_58_i0_ in range(default__.abs(841))]): 373})).keys.Elements:
                    d_59_v0_: _dafny.Seq = compr_14_
                    if (d_59_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('d')])).cardinality, 788])) for d_57_i1_ in range(default__.abs(-569))])), 871])): True})) for d_58_i0_ in range(default__.abs(841))]): 373})):
                        coll14_ = coll14_.union(_dafny.Set([d_59_v0_]))
                return _dafny.Set(coll14_)
            def iife15_():
                coll15_ = _dafny.Set()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(-309, 959):
                    d_60_v1_: int = compr_15_
                    if ((-309) <= (d_60_v1_)) and ((d_60_v1_) < (959)):
                        coll15_ = coll15_.union(_dafny.Set([(d_60_v1_) + (len(_dafny.Map({False: True})))]))
                return _dafny.Set(coll15_)
            return _dafny.MultiSet([D2_DC9(len(_dafny.SeqWithoutIsStrInference([len(iife14_()
), len(_dafny.Map({879: False}))])), True, len(iife15_()
), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxbwm"))))])
        elif True:
            d_61___mcc_h0_ = source3_.cf80
            d_62_cf80_ = d_61___mcc_h0_
            return _dafny.MultiSet([D2_DC9(254, True, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydbfhqa")))), len(_dafny.Map({True: True}))), D2_DC9(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jiass"))), True, len(_dafny.Set({not(True)})), -693)])

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('v') if False else _dafny.CodePoint('c')) for d_63_i0_ in range(default__.abs(-805))])

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([not(False), True, False])) + (_dafny.SeqWithoutIsStrInference([True, True]))) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm38(p0, globalState):
        if False:
            return D9_DC28()
        elif True:
            return D9_DC28()

    @staticmethod
    def fm39(p0, p1, p2, p3, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in _dafny.IntegerRange(266, -166):
                d_64_v0_: int = compr_16_
                if ((266) <= (d_64_v0_)) and ((d_64_v0_) < (-166)):
                    coll16_[default__.safeDivisionInt(d_64_v0_, (D28_DC88(not(False), 15)).cf169)] = D1_DC5(_dafny.Set({True}))
            return _dafny.Map(coll16_)
        return (iife16_()
        ) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))): D1_DC5(_dafny.Set({False, not(not(False))}))}) if True else _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False, True])])): D1_DC5(_dafny.Set({False, not(True), False, False, False}))})))

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        def iife17_():
            coll17_ = _dafny.Set()
            compr_17_: int
            for compr_17_ in _dafny.IntegerRange(-476, -249):
                d_65_v0_: int = compr_17_
                if ((-476) <= (d_65_v0_)) and ((d_65_v0_) < (-249)):
                    coll17_ = coll17_.union(_dafny.Set([(d_65_v0_) - ((0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkolr"))): 563}))))]))
            return _dafny.Set(coll17_)
        return (_dafny.Map({len(iife17_()
        ): 648})) | ((_dafny.Map({234: -527})) | (_dafny.Map({len(_dafny.Map({False: 244})): 461})))

    @staticmethod
    def fm41(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): False})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Set({True})): True})]))) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({193: True}), (D8_DC22(_dafny.Map({49: True}))).cf41]))

    @staticmethod
    def fm42(p0, globalState):
        return _dafny.CodePoint('j')

    @staticmethod
    def fm43(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdj"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbwindc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysl"))))

    @staticmethod
    def fm44(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([836, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppktkbrqj"))), 304]))])) | (_dafny.MultiSet([-356, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e'), _dafny.CodePoint('o')]))])).cardinality]))).intersection(_dafny.MultiSet([(0) - ((0) - ((0) - (len(_dafny.Map({True: False})))))]))

    @staticmethod
    def fm45(p0, p1, p2, globalState):
        source4_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wujr")): not(True)})
        d_66___mcc_h0_ = source4_
        d_67_cf79_ = d_66___mcc_h0_
        def iife18_():
            coll18_ = _dafny.Set()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(-369, 161):
                d_68_v0_: int = compr_18_
                if ((-369) <= (d_68_v0_)) and ((d_68_v0_) < (161)):
                    coll18_ = coll18_.union(_dafny.Set([(d_68_v0_) * (681)]))
            return _dafny.Set(coll18_)
        return _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(iife18_()
): _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnglsm")), (D24_DC71(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_69_i1_ in range(default__.abs(842))])), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omhtja")), False)).cf134])).cardinality])})) for d_70_i0_ in range(default__.abs(880))])

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        def iife19_():
            coll19_ = _dafny.Set()
            compr_19_: int
            for compr_19_ in _dafny.IntegerRange(869, -348):
                d_73_v0_: int = compr_19_
                if ((869) <= (d_73_v0_)) and ((d_73_v0_) < (-348)):
                    coll19_ = coll19_.union(_dafny.Set([default__.safeModuloInt(d_73_v0_, len(_dafny.SeqWithoutIsStrInference([False])))]))
            return _dafny.Set(coll19_)
        return _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False])])), len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_71_i0_ in range(default__.abs(620))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fcaxw")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_72_i1_ in range(default__.abs(48))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "noyjivij"))]))), ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvhcgs"))))) + (len(_dafny.SeqWithoutIsStrInference([len(iife19_()
        )])))})

    @staticmethod
    def fm47(globalState):
        return _dafny.Map({((D12_DC37(len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bevfc")): _dafny.CodePoint('v')})), -40, D3_DC10(_dafny.Map({False: False})), (0) - (-723), len(_dafny.SeqWithoutIsStrInference([False, True, False])))).cf73) >= (len(_dafny.SeqWithoutIsStrInference([(0) - (-676)]))): 126})

    @staticmethod
    def fm48(p0, p1, p2, globalState):
        source5_ = D22_DC62(504)
        if source5_.is_DC62:
            d_74___mcc_h0_ = source5_.cf119
            d_75_cf119_ = d_74___mcc_h0_
            if False:
                return D7_DC20(False, not(False))
            elif True:
                return D7_DC20(True, not(True))
        elif source5_.is_DC63:
            d_76___mcc_h1_ = source5_.cf120
            d_77_cf120_ = d_76___mcc_h1_
            return D7_DC20(d_77_cf120_, d_77_cf120_)
        elif source5_.is_DC64:
            d_78___mcc_h2_ = source5_.cf121
            d_79_cf121_ = d_78___mcc_h2_
            if d_79_cf121_:
                return D7_DC20(d_79_cf121_, d_79_cf121_)
            elif True:
                return D7_DC20(d_79_cf121_, not(d_79_cf121_))
        elif True:
            d_80___mcc_h3_ = source5_.cf118
            d_81_cf118_ = d_80___mcc_h3_
            return D7_DC20(False, False)

    @staticmethod
    def fm49(p0, globalState):
        return D0_DC2(not(not(not((False) or (True)))))

    @staticmethod
    def fm50(globalState):
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in _dafny.IntegerRange(535, 978):
                d_82_v0_: int = compr_20_
                if ((535) <= (d_82_v0_)) and ((d_82_v0_) < (978)):
                    coll20_[(d_82_v0_) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_83_i0_ in range(default__.abs(779))]))])))] = 467
            return _dafny.Map(coll20_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(False)])): 961})])) + ((_dafny.SeqWithoutIsStrInference([iife20_()
        ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(0) - ((_dafny.MultiSet([False])).cardinality): 129}), _dafny.Map({765: 200})])))

    @staticmethod
    def fm51(p0, p1, globalState):
        return (_dafny.Map({_dafny.Map({89: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True})), 204, 157, 840, 678]))}): 969})) | (_dafny.Map({_dafny.Map({len(_dafny.Map({_dafny.CodePoint('v'): 97})): (D22_DC62(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gh"))))).cf119}): (0) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpvehn"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xagwnx")))])).cardinality)}))

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        if (True if False else not(not(True))):
            return _dafny.Map({-602: D1_DC5(_dafny.Set({False, not(False)}))})
        elif True:
            def iife21_():
                coll21_ = _dafny.Map()
                compr_21_: int
                for compr_21_ in _dafny.IntegerRange(-925, 321):
                    d_84_v0_: int = compr_21_
                    if ((-925) <= (d_84_v0_)) and ((d_84_v0_) < (321)):
                        coll21_[(d_84_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twk"))))] = D1_DC5(_dafny.Set({False, False}))
                return _dafny.Map(coll21_)
            return iife21_()
            

    @staticmethod
    def fm53(p0, globalState):
        source6_ = D6_DC15(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_85_i1_ in range(default__.abs(965))])) for d_86_i0_ in range(default__.abs(-498))]))
        if source6_.is_DC16:
            d_87___mcc_h0_ = source6_.cf31
            d_88___mcc_h1_ = source6_.cf32
            d_89_cf32_ = d_88___mcc_h1_
            d_90_cf31_ = d_87___mcc_h0_
            return D2_DC8(d_89_cf32_, d_90_cf31_, d_90_cf31_)
        elif True:
            d_91___mcc_h2_ = source6_.cf30
            d_92_cf30_ = d_91___mcc_h2_
            return D2_DC8(False, 177, len(_dafny.SeqWithoutIsStrInference([True, False, False])))

    @staticmethod
    def fm54(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({_dafny.CodePoint('n'): -19})) | (_dafny.Map({_dafny.CodePoint('s'): 232}))) | ((_dafny.Map({_dafny.CodePoint('l'): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhlfqtxej"))) for d_93_i1_ in range(default__.abs(936))])) for d_94_i0_ in range(default__.abs(388))]))})) | (_dafny.Map({_dafny.CodePoint('k'): (_dafny.MultiSet([True])).cardinality})))

    @staticmethod
    def fm55(p0, p1, p2, p3, globalState):
        return D18_DC50(3, not(not(False)), _dafny.SeqWithoutIsStrInference([759, 915]), _dafny.CodePoint('q'), 157)

    @staticmethod
    def fm56(p0, globalState):
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wi"))]) if (True) or (True) else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_95_i0_ in range(default__.abs(-716))])])))

    @staticmethod
    def fm57(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D16_DC45((_dafny.MultiSet([not(False), False])).cardinality, True, 127, 708, False)).cf86, not(True)]))) | (_dafny.MultiSet([False])), (_dafny.MultiSet([False, False])) | (_dafny.MultiSet([True])), (_dafny.MultiSet([(D26_DC79(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_96_i0_ in range(default__.abs(996))])), 47, not(True))).cf152, True])) - (_dafny.MultiSet([(D28_DC88(True, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_97_i1_ in range(default__.abs(17))])))).cf168]))])

    @staticmethod
    def fm58(p0, p1, p2, p3, globalState):
        if not(True):
            return D22_DC62(819)
        elif True:
            return D22_DC62(len(_dafny.SeqWithoutIsStrInference([True, False, False])))

    @staticmethod
    def fm59(globalState):
        return D23_DC68(not((False) and (True)), ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([415])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqsqw"))), len(_dafny.Set({D12_DC38()}))])).cardinality) != (239))

    @staticmethod
    def fm60(p0, p1, p2, p3, globalState):
        def iife22_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in (_dafny.SeqWithoutIsStrInference([338])).Elements:
                d_98_v0_: int = compr_22_
                if (d_98_v0_) in (_dafny.SeqWithoutIsStrInference([338])):
                    coll22_[(d_98_v0_) + ((_dafny.MultiSet([213])).cardinality)] = True
            return _dafny.Map(coll22_)
        return (_dafny.Map({D23_DC68(False, not(False)): D18_DC50(len(iife22_()
), not(False), _dafny.SeqWithoutIsStrInference([-787]), _dafny.CodePoint('k'), 948)})) | ((D29_DC89(_dafny.Map({D23_DC68(False, False): D18_DC50((0) - (-614), False, _dafny.SeqWithoutIsStrInference([-690]), _dafny.CodePoint('u'), len(_dafny.Set({110, -397, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v')])), 697, len(_dafny.Map({-69: 554}))})))}))).cf170)

    @staticmethod
    def fm61(p0, p1, p2, globalState):
        return D21_DC57(_dafny.CodePoint('o'), (0) - (((_dafny.MultiSet([True, False])) | (_dafny.MultiSet([False]))).cardinality), len((_dafny.Map({True: True})) | (_dafny.Map({False: False}))), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_99_i0_ in range(default__.abs(424))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pae")))))

    @staticmethod
    def fm62(p0, p1, p2, globalState):
        return D24_DC71(default__.safeModuloInt((0) - (len(_dafny.SeqWithoutIsStrInference([(D24_DC72(True, 568, _dafny.CodePoint('d'), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_100_i0_ in range(default__.abs(257))]), False)).cf136, False]))), 878), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_101_i1_ in range(default__.abs(757))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmyvl"))), not(False))

    @staticmethod
    def fm63(p0, p1, globalState):
        return (_dafny.Map({_dafny.CodePoint('d'): False})) | ((_dafny.Map({_dafny.CodePoint('k'): True})) | (_dafny.Map({_dafny.CodePoint('m'): not(False)})))

    @staticmethod
    def fm64(p0, globalState):
        source7_ = D20_DC55(-723, True, 851, not(False))
        if source7_.is_DC55:
            d_102___mcc_h0_ = source7_.cf103
            d_103___mcc_h1_ = source7_.cf104
            d_104___mcc_h2_ = source7_.cf105
            d_105___mcc_h3_ = source7_.cf106
            d_106_cf106_ = d_105___mcc_h3_
            d_107_cf105_ = d_104___mcc_h2_
            d_108_cf104_ = d_103___mcc_h1_
            d_109_cf103_ = d_102___mcc_h0_
            def iife23_():
                coll23_ = _dafny.Map()
                compr_23_: _dafny.Map
                for compr_23_ in (_dafny.Map({_dafny.Map({d_106_cf106_: len(_dafny.SeqWithoutIsStrInference([d_107_cf105_]))}): d_106_cf106_})).keys.Elements:
                    d_110_v0_: _dafny.Map = compr_23_
                    if (d_110_v0_) in (_dafny.Map({_dafny.Map({d_106_cf106_: len(_dafny.SeqWithoutIsStrInference([d_107_cf105_]))}): d_106_cf106_})):
                        coll23_[d_110_v0_] = _dafny.CodePoint('v')
                return _dafny.Map(coll23_)
            return iife23_()
            
        elif True:
            d_111___mcc_h4_ = source7_.cf102
            d_112_cf102_ = d_111___mcc_h4_
            if not(not(True)):
                return _dafny.Map({_dafny.Map({True: 238}): _dafny.CodePoint('p')})
            elif True:
                return _dafny.Map({_dafny.Map({True: 824}): _dafny.CodePoint('t')})

    @staticmethod
    def m8(p0, p1, p2, p3, globalState):
        r0: C6 = None
        d_113_i0_: int
        d_113_i0_ = 0
        with _dafny.label("0"):
            while p0:
                with _dafny.c_label("0"):
                    if (d_113_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_113_i0_ = (d_113_i0_) + (1)
                    d_114_v0_: _dafny.Seq
                    d_114_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihpwnka"))
                    d_115_v1_: _dafny.Set
                    d_115_v1_ = _dafny.Set({not(p0)})
                    d_116_v2_: _dafny.Array
                    nw0_ = _dafny.Array(None, 24)
                    d_116_v2_ = nw0_
                    d_117_v3_: _dafny.Seq
                    d_117_v3_ = _dafny.SeqWithoutIsStrInference([d_116_v2_])
                    d_118_v4_: _dafny.Map
                    d_118_v4_ = _dafny.Map({(d_117_v3_)[default__.safeIndex(p2, len(d_117_v3_))]: p0})
                    d_119_v5_: _dafny.Map
                    d_119_v5_ = _dafny.Map({default__.safeDivisionInt(len(default__.fm0(p0, d_114_v0_, d_115_v1_, p2, globalState)), 459): d_118_v4_})
                    d_119_v5_ = (d_119_v5_).set(default__.safeDivisionInt(p2, p2), d_118_v4_)
                    (globalState).f9 = not(False)
                    d_120_v6_: _dafny.Map
                    d_120_v6_ = _dafny.Map({p0: p2})
                    d_121_v7_: str
                    d_121_v7_ = _dafny.CodePoint('k')
                    d_122_v8_: _dafny.Map
                    d_122_v8_ = _dafny.Map({d_120_v6_: d_121_v7_})
                    d_123_v9_: _dafny.Map
                    d_123_v9_ = _dafny.Map({p2: 834})
                    d_124_v10_: T0
                    nw1_ = C4()
                    nw1_.ctor__(p0, p2, d_114_v0_, d_123_v9_)
                    d_124_v10_ = nw1_
                    d_125_v11_: _dafny.Map
                    d_125_v11_ = _dafny.Map({(p3)[default__.safeIndex(p2, len(p3))]: d_124_v10_})
                    d_126_v12_: C6
                    nw2_ = C6()
                    nw2_.ctor__(d_122_v8_, len(d_125_v11_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvtsa")))
                    d_126_v12_ = nw2_
                    d_127_v13_: _dafny.Array
                    def lambda0_(d_128_p2_, d_129_p0_):
                        def lambda1_(d_130_i1_):
                            return _dafny.Map({d_128_p2_: d_129_p0_})

                        return lambda1_

                    init0_ = lambda0_(p2, p0)
                    nw3_ = _dafny.Array(None, 10)
                    for i0_0_ in range(nw3_.length(0)):
                        nw3_[i0_0_] = init0_(i0_0_)
                    d_127_v13_ = nw3_
                    d_131_v14_: _dafny.Map
                    d_131_v14_ = _dafny.Map({(0) - (p2): p0})
                    index0_ = default__.safeIndex(472, (d_127_v13_).length(0))
                    (d_127_v13_)[index0_] = d_131_v14_
                    index1_ = default__.safeIndex(472, (d_127_v13_).length(0))
                    (d_127_v13_)[index1_] = d_131_v14_
                    pass
            pass
        d_132_v15_: D7
        d_132_v15_ = D7_DC19(p2, p0)
        if (d_132_v15_).cf37:
            d_133_v16_: _dafny.Array
            nw4_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
            d_133_v16_ = nw4_
            index2_ = default__.safeIndex(683, (d_133_v16_).length(0))
            (d_133_v16_)[index2_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_134_i2_ in range(default__.abs(417))])
            d_135_v17_: _dafny.Seq
            d_135_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgesotcly"))
            d_136_v18_: D24
            d_136_v18_ = D24_DC71(p2, d_135_v17_, p0)
            index3_ = default__.safeIndex(683, (d_133_v16_).length(0))
            (d_133_v16_)[index3_] = ((d_136_v18_).cf134) + (d_135_v17_)
            d_137_v19_: D27
            d_137_v19_ = D27_DC82(p2, p2)
            source8_ = d_137_v19_
            if source8_.is_DC82:
                d_138___mcc_h0_ = source8_.cf155
                d_139___mcc_h1_ = source8_.cf156
                d_140_cf156_ = d_139___mcc_h1_
                d_141_cf155_ = d_138___mcc_h0_
                d_142_v20_: str
                d_142_v20_ = _dafny.CodePoint('o')
                d_142_v20_ = _dafny.CodePoint('x')
                d_143_v21_: _dafny.Array
                nw5_ = _dafny.Array(None, 8)
                nw5_[int(0)] = p0
                nw5_[int(1)] = p0
                nw5_[int(2)] = (p0 if True else True)
                nw5_[int(3)] = p0
                nw5_[int(4)] = p0
                nw5_[int(5)] = default__.fm24(p0, p0, False, d_140_cf156_, globalState)
                nw5_[int(6)] = p0
                nw5_[int(7)] = p0
                d_143_v21_ = nw5_
                d_144_v22_: _dafny.MultiSet
                d_144_v22_ = _dafny.MultiSet([p0])
                d_145_v23_: D26
                d_145_v23_ = D26_DC79((d_137_v19_).cf156, (d_144_v22_).cardinality, p0)
                index4_ = default__.safeIndex(381, (d_143_v21_).length(0))
                (d_143_v21_)[index4_] = (d_145_v23_).cf152
                index5_ = default__.safeIndex(381, (d_143_v21_).length(0))
                rhs0_ = default__.fm5((p2) - (d_140_cf156_), False, not (p0) or (p0), globalState)
                rhs1_ = (p2) - (d_140_cf156_)
                lhs0_ = d_143_v21_
                lhs1_ = default__.safeIndex(381, (d_143_v21_).length(0))
                lhs2_ = globalState
                lhs0_[lhs1_] = rhs0_
                lhs2_.f0 = rhs1_
                d_146_v24_: _dafny.Array
                nw6_ = _dafny.Array(None, 10)
                d_146_v24_ = nw6_
                nw7_ = _dafny.Array(None, 5)
                d_146_v24_ = nw7_
                d_147_v25_: C5
                nw8_ = C5()
                nw8_.ctor__(724, default__.safeDivisionInt(p2, d_140_cf156_), d_141_cf155_, (d_135_v17_) + (_dafny.SeqWithoutIsStrInference([d_142_v20_ for d_148_i3_ in range(default__.abs(-386))])))
                d_147_v25_ = nw8_
                d_147_v25_ = d_147_v25_
            elif source8_.is_DC83:
                d_149___mcc_h2_ = source8_.cf157
                d_150___mcc_h3_ = source8_.cf158
                d_151_cf158_ = d_150___mcc_h3_
                d_152_cf157_ = d_149___mcc_h2_
                d_153_v26_: _dafny.Array
                nw9_ = _dafny.Array(_dafny.MultiSet({}), 4)
                d_153_v26_ = nw9_
                d_154_v27_: C7
                nw10_ = C7()
                nw10_.ctor__()
                d_154_v27_ = nw10_
                d_155_v28_: _dafny.Map
                d_155_v28_ = _dafny.Map({d_152_cf157_: d_154_v27_})
                d_156_v29_: _dafny.MultiSet
                d_156_v29_ = _dafny.MultiSet([d_155_v28_])
                index6_ = default__.safeIndex(222, (d_153_v26_).length(0))
                (d_153_v26_)[index6_] = d_156_v29_
                index7_ = default__.safeIndex(222, (d_153_v26_).length(0))
                (d_153_v26_)[index7_] = (d_156_v29_).intersection(d_156_v29_)
                d_157_v30_: _dafny.Map
                d_157_v30_ = _dafny.Map({d_151_cf158_: -291})
                d_158_v31_: _dafny.Seq
                d_158_v31_ = _dafny.SeqWithoutIsStrInference([len(d_157_v30_)])
                d_159_v32_: _dafny.MultiSet
                d_159_v32_ = _dafny.MultiSet([d_152_cf157_])
                d_160_v33_: C2
                nw11_ = C2()
                nw11_.ctor__(d_158_v31_, _dafny.Map({(d_159_v32_).cardinality: d_152_cf157_}))
                d_160_v33_ = nw11_
                (globalState).f5 = d_152_cf157_
                d_161_v34_: _dafny.Array
                def lambda2_(d_162_p0_):
                    def lambda3_(d_163_i4_):
                        return _dafny.MultiSet([d_162_p0_])

                    return lambda3_

                init1_ = lambda2_(p0)
                nw12_ = _dafny.Array(None, 24)
                for i0_1_ in range(nw12_.length(0)):
                    nw12_[i0_1_] = init1_(i0_1_)
                d_161_v34_ = nw12_
                d_161_v34_ = d_161_v34_
            elif source8_.is_DC84:
                d_164___mcc_h4_ = source8_.cf159
                d_165___mcc_h5_ = source8_.cf160
                d_166___mcc_h6_ = source8_.cf161
                d_167___mcc_h7_ = source8_.cf162
                d_168_cf162_ = d_167___mcc_h7_
                d_169_cf161_ = d_166___mcc_h6_
                d_170_cf160_ = d_165___mcc_h5_
                d_171_cf159_ = d_164___mcc_h4_
                (globalState).f24 = ((_dafny.MultiSet([p0])).cardinality) - (len(p1))
                (globalState).f9 = default__.fm5(default__.safeModuloInt((0) - (p2), 506), p0, (p0 if d_169_cf161_ else True), globalState)
                d_172_v35_: _dafny.MultiSet
                d_172_v35_ = _dafny.MultiSet([-677])
                (globalState).f8 = (default__.safeDivisionInt(((d_172_v35_)[p2] if (p2) in (d_172_v35_) else p2), p2)) * (d_171_cf159_)
                (globalState).f9 = d_169_cf161_
            elif source8_.is_DC81:
                d_173___mcc_h8_ = source8_.cf154
                d_174_cf154_ = d_173___mcc_h8_
                (globalState).f25 = p0
                d_175_v36_: str
                d_175_v36_ = _dafny.CodePoint('l')
                d_175_v36_ = d_175_v36_
                index8_ = default__.safeIndex(683, (d_133_v16_).length(0))
                (d_133_v16_)[index8_] = (_dafny.SeqWithoutIsStrInference([d_175_v36_ for d_176_i5_ in range(default__.abs(277))])) + ((d_133_v16_)[default__.safeIndex(683, (d_133_v16_).length(0))])
                d_177_v37_: _dafny.Seq
                d_177_v37_ = _dafny.SeqWithoutIsStrInference([p0, (p2) <= (p2), default__.fm24(p0, p0, p0, (0) - (p2), globalState)])
                d_178_v38_: _dafny.MultiSet
                d_178_v38_ = _dafny.MultiSet([(d_135_v17_) + ((d_133_v16_)[default__.safeIndex(683, (d_133_v16_).length(0))]), d_135_v17_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxfhafvnj"))) + (d_135_v17_), (d_133_v16_)[default__.safeIndex(683, (d_133_v16_).length(0))]])
                d_179_v39_: _dafny.Set
                d_179_v39_ = _dafny.Set({p0})
                d_180_v40_: _dafny.Set
                d_180_v40_ = _dafny.Set({len(d_179_v39_)})
                rhs2_ = p0
                rhs3_ = d_177_v37_
                rhs4_ = d_175_v36_
                rhs5_ = default__.fm42(d_180_v40_, globalState)
                rhs6_ = (d_178_v38_).intersection(d_178_v38_)
                lhs3_ = globalState
                lhs3_.f25 = rhs2_
                d_177_v37_ = rhs3_
                d_175_v36_ = rhs4_
                d_175_v36_ = rhs5_
                d_178_v38_ = rhs6_
            elif True:
                d_181___mcc_h9_ = source8_.cf163
                d_182_cf163_ = d_181___mcc_h9_
                d_183_v41_: str
                d_183_v41_ = _dafny.CodePoint('e')
                d_184_v42_: _dafny.Map
                d_184_v42_ = _dafny.Map({470: p2})
                d_185_v43_: C4
                nw13_ = C4()
                nw13_.ctor__(p0, default__.fm4(d_183_v41_, globalState), (d_133_v16_)[default__.safeIndex(683, (d_133_v16_).length(0))], (d_184_v42_) | (d_184_v42_))
                d_185_v43_ = nw13_
                d_186_v44_: _dafny.MultiSet
                d_186_v44_ = _dafny.MultiSet([p2])
                d_187_v45_: _dafny.Array
                nw14_ = _dafny.Array(None, 23)
                nw14_[int(0)] = (d_185_v43_).fm8(d_186_v44_, p0, p2, p2, globalState)
                nw14_[int(1)] = p0
                nw14_[int(2)] = p0
                nw14_[int(3)] = p0
                nw14_[int(4)] = (d_185_v43_).f32
                nw14_[int(5)] = p0
                nw14_[int(6)] = p0
                nw14_[int(7)] = (d_185_v43_).f32
                nw14_[int(8)] = (d_185_v43_).f32
                nw14_[int(9)] = p0
                nw14_[int(10)] = p0
                nw14_[int(11)] = default__.fm24(p0, p0, (d_185_v43_).f32, p2, globalState)
                nw14_[int(12)] = not(not(not((d_185_v43_).f32)))
                nw14_[int(13)] = True
                nw14_[int(14)] = (d_185_v43_).f32
                nw14_[int(15)] = (d_185_v43_).f32
                nw14_[int(16)] = p0
                nw14_[int(17)] = p0
                nw14_[int(18)] = (d_185_v43_).f32
                nw14_[int(19)] = (d_185_v43_).f32
                nw14_[int(20)] = (d_185_v43_).f32
                nw14_[int(21)] = (d_185_v43_).f32
                nw14_[int(22)] = p0
                d_187_v45_ = nw14_
                d_188_v46_: _dafny.Map
                d_188_v46_ = _dafny.Map({d_183_v41_: d_187_v45_})
                d_189_v47_: _dafny.Seq
                d_189_v47_ = _dafny.SeqWithoutIsStrInference([d_187_v45_])
                d_190_v48_: _dafny.Seq
                d_190_v48_ = _dafny.SeqWithoutIsStrInference([((d_188_v46_)[d_183_v41_] if (d_183_v41_) in (d_188_v46_) else (d_189_v47_)[default__.safeIndex(p2, len(d_189_v47_))]), d_187_v45_])
                d_191_v49_: _dafny.Set
                d_191_v49_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([p2 for d_192_i6_ in range(default__.abs(798))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bngxvpx"))), p2})
                d_193_v50_: D21
                d_193_v50_ = D21_DC56(d_191_v49_)
                d_194_v51_: _dafny.Array
                nw15_ = _dafny.Array(None, 6)
                nw15_[int(0)] = p2
                nw15_[int(1)] = p2
                nw15_[int(2)] = len((d_189_v47_) + (d_190_v48_))
                nw15_[int(3)] = p2
                nw15_[int(4)] = (len((d_193_v50_).cf107)) + (len(d_135_v17_))
                nw15_[int(5)] = p2
                d_194_v51_ = nw15_
                d_194_v51_ = d_194_v51_
                d_195_v52_: _dafny.Seq
                d_195_v52_ = _dafny.SeqWithoutIsStrInference([(d_135_v17_) + ((d_133_v16_)[default__.safeIndex(683, (d_133_v16_).length(0))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gaoopuw")), d_135_v17_])
                index9_ = default__.safeIndex(683, (d_133_v16_).length(0))
                rhs7_ = d_135_v17_
                rhs8_ = d_194_v51_
                rhs9_ = _dafny.SeqWithoutIsStrInference([d_135_v17_, ((d_133_v16_)[default__.safeIndex(683, (d_133_v16_).length(0))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gch")))])
                lhs4_ = d_133_v16_
                lhs5_ = default__.safeIndex(683, (d_133_v16_).length(0))
                lhs4_[lhs5_] = rhs7_
                d_194_v51_ = rhs8_
                d_195_v52_ = rhs9_
                d_196_v53_: _dafny.Set
                d_196_v53_ = _dafny.Set({p0, (d_185_v43_).f32, p0, (d_185_v43_).fm8(d_186_v44_, (d_185_v43_).f32, p2, p2, globalState)})
                (globalState).f9 = (len((d_196_v53_) - (d_196_v53_))) >= (p2)
            (globalState).f17 = p2
            d_197_v54_: str
            d_197_v54_ = _dafny.CodePoint('u')
            index10_ = default__.safeIndex(683, (d_133_v16_).length(0))
            (d_133_v16_)[index10_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdenfwckf"))) + (d_135_v17_)).set(default__.safeIndex(p2, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdenfwckf"))) + (d_135_v17_))), d_197_v54_)
            d_198_v55_: _dafny.MultiSet
            d_198_v55_ = _dafny.MultiSet([p0])
            (globalState).f8 = ((d_198_v55_).cardinality) - (p2)
        elif True:
            d_199_v56_: _dafny.Array
            nw16_ = _dafny.Array(_dafny.Seq({}), 1)
            d_199_v56_ = nw16_
            d_200_v57_: _dafny.Seq
            d_200_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmukqaetm"))
            d_201_v58_: D0
            d_201_v58_ = D0_DC1(d_200_v57_, p2, p2)
            d_202_v59_: _dafny.Seq
            d_202_v59_ = _dafny.SeqWithoutIsStrInference([len((d_201_v58_).cf1)])
            index11_ = default__.safeIndex(84, (d_199_v56_).length(0))
            (d_199_v56_)[index11_] = d_202_v59_
            index12_ = default__.safeIndex(84, (d_199_v56_).length(0))
            (d_199_v56_)[index12_] = _dafny.SeqWithoutIsStrInference([(39) + (p2) for d_203_i7_ in range(default__.abs(197))])
            (globalState).f25 = not(p0)
            d_204_v60_: _dafny.Array
            nw17_ = _dafny.Array(_dafny.Set({}), 6)
            d_204_v60_ = nw17_
            d_205_v61_: _dafny.MultiSet
            d_205_v61_ = _dafny.MultiSet([p0])
            d_206_v62_: _dafny.Set
            d_206_v62_ = _dafny.Set({d_205_v61_})
            index13_ = default__.safeIndex(195, (d_204_v60_).length(0))
            (d_204_v60_)[index13_] = d_206_v62_
            d_207_v63_: _dafny.Array
            nw18_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_207_v63_ = nw18_
            d_208_v64_: _dafny.Array
            nw19_ = _dafny.Array(False, 22)
            d_208_v64_ = nw19_
            index14_ = default__.safeIndex(141, (d_207_v63_).length(0))
            (d_207_v63_)[index14_] = d_208_v64_
            index15_ = default__.safeIndex(486, (d_207_v63_).length(0))
            (d_207_v63_)[index15_] = d_208_v64_
            index16_ = default__.safeIndex(195, (d_204_v60_).length(0))
            index17_ = default__.safeIndex(141, (d_207_v63_).length(0))
            index18_ = default__.safeIndex(486, (d_207_v63_).length(0))
            rhs10_ = _dafny.Set({d_205_v61_})
            rhs11_ = p2
            rhs12_ = d_208_v64_
            rhs13_ = p0
            rhs14_ = d_208_v64_
            lhs6_ = d_204_v60_
            lhs7_ = default__.safeIndex(195, (d_204_v60_).length(0))
            lhs8_ = globalState
            lhs9_ = d_207_v63_
            lhs10_ = default__.safeIndex(141, (d_207_v63_).length(0))
            lhs11_ = globalState
            lhs12_ = d_207_v63_
            lhs13_ = default__.safeIndex(486, (d_207_v63_).length(0))
            lhs6_[lhs7_] = rhs10_
            lhs8_.f24 = rhs11_
            lhs9_[lhs10_] = rhs12_
            lhs11_.f9 = rhs13_
            lhs12_[lhs13_] = rhs14_
            d_209_v65_: str
            d_209_v65_ = _dafny.CodePoint('p')
            d_209_v65_ = d_209_v65_
            d_210_v67_: _dafny.Array
            def lambda4_(d_211_p0_):
                def lambda5_(d_212_i8_):
                    def iife24_():
                        coll24_ = _dafny.Map()
                        compr_24_: int
                        for compr_24_ in _dafny.IntegerRange(968, 27):
                            d_213_v66_: int = compr_24_
                            if ((968) <= (d_213_v66_)) and ((d_213_v66_) < (27)):
                                coll24_[(d_213_v66_) * (-584)] = d_211_p0_
                        return _dafny.Map(coll24_)
                    return iife24_()
                    

                return lambda5_

            init2_ = lambda4_(p0)
            nw20_ = _dafny.Array(None, 24)
            for i0_2_ in range(nw20_.length(0)):
                nw20_[i0_2_] = init2_(i0_2_)
            d_210_v67_ = nw20_
            d_214_v68_: _dafny.Set
            d_214_v68_ = _dafny.Set({p0})
            d_215_v69_: _dafny.Map
            d_215_v69_ = _dafny.Map({(0) - (len(d_214_v68_)): p0})
            index19_ = default__.safeIndex(809, (d_210_v67_).length(0))
            (d_210_v67_)[index19_] = d_215_v69_
            index20_ = default__.safeIndex(809, (d_210_v67_).length(0))
            (d_210_v67_)[index20_] = d_215_v69_
        d_216_i9_: int
        d_216_i9_ = 0
        with _dafny.label("1"):
            while default__.fm24(p0, p0, not (p0) or (p0), p2, globalState):
                with _dafny.c_label("1"):
                    if (d_216_i9_) >= (100):
                        raise _dafny.Break("1")
                    d_216_i9_ = (d_216_i9_) + (1)
                    d_217_v70_: _dafny.Seq
                    d_217_v70_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0])])
                    d_218_v71_: D23
                    d_218_v71_ = D23_DC66((p0) in ((d_217_v70_)[default__.safeIndex(p2, len(d_217_v70_))]))
                    d_219_v72_: str
                    d_219_v72_ = _dafny.CodePoint('j')
                    d_220_v73_: _dafny.Seq
                    d_220_v73_ = _dafny.SeqWithoutIsStrInference([p2])
                    rhs15_ = d_218_v71_
                    rhs16_ = default__.fm20(default__.fm24(p0, p0, p0, len(d_220_v73_), globalState), _dafny.CodePoint('t'), (0) - (default__.safeModuloInt(p2, p2)), d_219_v72_, globalState)
                    rhs17_ = p2
                    lhs14_ = globalState
                    d_218_v71_ = rhs15_
                    d_219_v72_ = rhs16_
                    lhs14_.f0 = rhs17_
                    d_221_v74_: _dafny.Map
                    d_221_v74_ = _dafny.Map({p0: p2})
                    d_222_v75_: _dafny.Map
                    d_222_v75_ = _dafny.Map({d_221_v74_: _dafny.CodePoint('r')})
                    d_223_v76_: _dafny.Seq
                    d_223_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsgbnqp"))
                    d_224_v77_: T0
                    nw21_ = C6()
                    nw21_.ctor__(d_222_v75_, p2, d_223_v76_)
                    d_224_v77_ = nw21_
                    source9_ = D28_DC87(d_224_v77_, d_220_v73_, _dafny.CodePoint('v'))
                    if source9_.is_DC87:
                        d_225___mcc_h10_ = source9_.cf165
                        d_226___mcc_h11_ = source9_.cf166
                        d_227___mcc_h12_ = source9_.cf167
                        d_228_cf167_ = d_227___mcc_h12_
                        d_229_cf166_ = d_226___mcc_h11_
                        d_230_cf165_ = d_225___mcc_h10_
                        (globalState).f25 = (p2) < (p2)
                        d_231_v78_: _dafny.Array
                        def lambda6_(d_232_p2_):
                            def lambda7_(d_233_i10_):
                                return (d_233_i10_) * (d_232_p2_)

                            return lambda7_

                        init3_ = lambda6_(p2)
                        nw22_ = _dafny.Array(None, 10)
                        for i0_3_ in range(nw22_.length(0)):
                            nw22_[i0_3_] = init3_(i0_3_)
                        d_231_v78_ = nw22_
                        index21_ = default__.safeIndex(502, (d_231_v78_).length(0))
                        (d_231_v78_)[index21_] = p2
                        index22_ = default__.safeIndex(502, (d_231_v78_).length(0))
                        (d_231_v78_)[index22_] = -813
                        (globalState).f25 = True
                        index23_ = default__.safeIndex(502, (d_231_v78_).length(0))
                        rhs18_ = default__.fm4(d_219_v72_, globalState)
                        rhs19_ = not(False)
                        rhs20_ = ((d_231_v78_)[default__.safeIndex(502, (d_231_v78_).length(0))] if False else (d_231_v78_)[default__.safeIndex(502, (d_231_v78_).length(0))])
                        rhs21_ = p0
                        rhs22_ = len(d_223_v76_)
                        lhs15_ = globalState
                        lhs16_ = globalState
                        lhs17_ = globalState
                        lhs18_ = globalState
                        lhs19_ = d_231_v78_
                        lhs20_ = default__.safeIndex(502, (d_231_v78_).length(0))
                        lhs15_.f17 = rhs18_
                        lhs16_.f9 = rhs19_
                        lhs17_.f0 = rhs20_
                        lhs18_.f25 = rhs21_
                        lhs19_[lhs20_] = rhs22_
                    elif source9_.is_DC88:
                        d_234___mcc_h13_ = source9_.cf168
                        d_235___mcc_h14_ = source9_.cf169
                        d_236_cf169_ = d_235___mcc_h14_
                        d_237_cf168_ = d_234___mcc_h13_
                        (globalState).f25 = (d_221_v74_) != ((d_221_v74_) | (d_221_v74_))
                        d_238_v79_: _dafny.Array
                        nw23_ = _dafny.Array(False, 11)
                        d_238_v79_ = nw23_
                        index24_ = default__.safeIndex(371, (d_238_v79_).length(0))
                        (d_238_v79_)[index24_] = p0
                        index25_ = default__.safeIndex(371, (d_238_v79_).length(0))
                        (d_238_v79_)[index25_] = (d_219_v72_) not in (d_223_v76_)
                        d_239_v80_: _dafny.Map
                        d_239_v80_ = _dafny.Map({p0: not(False)})
                        d_239_v80_ = (d_239_v80_).set(((d_238_v79_)[default__.safeIndex(371, (d_238_v79_).length(0))]) and (not(p0)), p0)
                        d_240_v81_: _dafny.Array
                        nw24_ = _dafny.Array(None, 4)
                        nw24_[int(0)] = d_236_cf169_
                        nw24_[int(1)] = p2
                        nw24_[int(2)] = d_236_cf169_
                        nw24_[int(3)] = d_236_cf169_
                        d_240_v81_ = nw24_
                        index26_ = default__.safeIndex(389, (d_240_v81_).length(0))
                        (d_240_v81_)[index26_] = len((d_223_v76_) + (d_223_v76_))
                        index27_ = default__.safeIndex(389, (d_240_v81_).length(0))
                        (d_240_v81_)[index27_] = (p2) - ((0) - (d_236_cf169_))
                    elif True:
                        d_241___mcc_h15_ = source9_.cf164
                        d_242_cf164_ = d_241___mcc_h15_
                        d_219_v72_ = d_219_v72_
                        (globalState).f0 = p2
                        d_243_v82_: _dafny.Array
                        nw25_ = _dafny.Array(_dafny.Seq({}), 2)
                        d_243_v82_ = nw25_
                        d_244_v83_: T1
                        nw26_ = C6()
                        nw26_.ctor__(d_222_v75_, p2, d_223_v76_)
                        d_244_v83_ = nw26_
                        d_245_v84_: _dafny.Seq
                        d_245_v84_ = _dafny.SeqWithoutIsStrInference([d_244_v83_])
                        index28_ = default__.safeIndex(673, (d_243_v82_).length(0))
                        (d_243_v82_)[index28_] = (d_245_v84_) + (_dafny.SeqWithoutIsStrInference([d_244_v83_]))
                        d_246_v85_: C0
                        nw27_ = C0()
                        nw27_.ctor__()
                        d_246_v85_ = nw27_
                        d_247_v86_: _dafny.Seq
                        d_247_v86_ = _dafny.SeqWithoutIsStrInference([d_246_v85_])
                        d_248_v87_: _dafny.Set
                        d_248_v87_ = _dafny.Set({(d_244_v83_).f26, len(_dafny.Set({len(d_247_v86_), p2, (d_244_v83_).f26, p2, (d_244_v83_).f26}))})
                        d_249_v88_: _dafny.Array
                        nw28_ = _dafny.Array(None, 6)
                        nw28_[int(0)] = (D6_DC16(len(d_220_v73_), p0)).cf32
                        nw28_[int(1)] = p0
                        nw28_[int(2)] = (p0) == (p0)
                        nw28_[int(3)] = (p2) == ((0) - ((d_244_v83_).f26))
                        nw28_[int(4)] = (d_248_v87_).ispropersubset(_dafny.Set({p2}))
                        nw28_[int(5)] = (p0) == (p0)
                        d_249_v88_ = nw28_
                        index29_ = default__.safeIndex(835, (d_249_v88_).length(0))
                        (d_249_v88_)[index29_] = p0
                        d_250_v89_: _dafny.Map
                        d_250_v89_ = _dafny.Map({p2: (d_244_v83_).f26})
                        index30_ = default__.safeIndex(673, (d_243_v82_).length(0))
                        index31_ = default__.safeIndex(835, (d_249_v88_).length(0))
                        rhs23_ = len(d_250_v89_)
                        rhs24_ = (d_245_v84_).set(default__.safeIndex((d_244_v83_).f26, len(d_245_v84_)), d_244_v83_)
                        rhs25_ = d_249_v88_
                        rhs26_ = p2
                        rhs27_ = not(p0)
                        lhs21_ = globalState
                        lhs22_ = d_243_v82_
                        lhs23_ = default__.safeIndex(673, (d_243_v82_).length(0))
                        lhs24_ = globalState
                        lhs25_ = globalState
                        lhs26_ = d_249_v88_
                        lhs27_ = default__.safeIndex(835, (d_249_v88_).length(0))
                        lhs21_.f0 = rhs23_
                        lhs22_[lhs23_] = rhs24_
                        lhs24_.f1 = rhs25_
                        lhs25_.f0 = rhs26_
                        lhs26_[lhs27_] = rhs27_
                        d_251_v90_: _dafny.Set
                        d_251_v90_ = _dafny.Set({p0})
                        d_251_v90_ = (d_251_v90_) | (_dafny.Set({p0}))
                    d_252_v91_: C0
                    nw29_ = C0()
                    nw29_.ctor__()
                    d_252_v91_ = nw29_
                    d_253_v92_: C3
                    nw30_ = C3()
                    nw30_.ctor__(p2, default__.fm4(_dafny.CodePoint('i'), globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))
                    d_253_v92_ = nw30_
                    pass
            pass
        hi0_ = p2
        for d_254_i11_ in range(p2, hi0_):
            d_255_v93_: _dafny.MultiSet
            d_255_v93_ = _dafny.MultiSet([p2])
            d_256_v94_: str
            d_256_v94_ = _dafny.CodePoint('s')
            d_257_v95_: _dafny.Seq
            d_257_v95_ = _dafny.SeqWithoutIsStrInference([d_254_i11_, default__.fm4(d_256_v94_, globalState)])
            (globalState).f9 = (p0) or ((d_255_v93_) != ((_dafny.MultiSet(d_257_v95_)).set(p2, default__.abs(d_254_i11_))))
            d_258_v96_: _dafny.Array
            def lambda8_(d_259_v95_):
                def lambda9_(d_260_i12_):
                    return (d_259_v95_) + (d_259_v95_)

                return lambda9_

            init4_ = lambda8_(d_257_v95_)
            nw31_ = _dafny.Array(None, 9)
            for i0_4_ in range(nw31_.length(0)):
                nw31_[i0_4_] = init4_(i0_4_)
            d_258_v96_ = nw31_
            index32_ = default__.safeIndex(396, (d_258_v96_).length(0))
            (d_258_v96_)[index32_] = d_257_v95_
            index33_ = default__.safeIndex(396, (d_258_v96_).length(0))
            (d_258_v96_)[index33_] = d_257_v95_
            d_261_v97_: _dafny.Seq
            d_261_v97_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klwb"))
            (globalState).f13 = (d_261_v97_) + (d_261_v97_)
            d_262_v98_: _dafny.Array
            nw32_ = _dafny.Array(False, 8)
            d_262_v98_ = nw32_
            d_263_v99_: _dafny.Map
            d_263_v99_ = _dafny.Map({d_262_v98_: p2})
            (globalState).f17 = ((d_263_v99_)[d_262_v98_] if (d_262_v98_) in (d_263_v99_) else p2)
        if p0:
            d_264_v100_: _dafny.Seq
            d_264_v100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgjojleha"))
            d_265_v101_: _dafny.Map
            d_265_v101_ = _dafny.Map({d_264_v100_: p0})
            d_266_v102_: _dafny.Seq
            d_266_v102_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            d_267_v103_: _dafny.Map
            d_267_v103_ = _dafny.Map({p2: p2})
            d_268_v104_: C2
            nw33_ = C2()
            nw33_.ctor__(d_266_v102_, (d_267_v103_).set(p2, p2))
            d_268_v104_ = nw33_
            d_269_v105_: _dafny.MultiSet
            d_269_v105_ = _dafny.MultiSet([d_268_v104_])
            d_270_v106_: str
            d_270_v106_ = _dafny.CodePoint('j')
            d_271_v107_: _dafny.MultiSet
            d_271_v107_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljyadfw"))])
            d_272_v108_: _dafny.Array
            nw34_ = _dafny.Array(None, 15)
            nw34_[int(0)] = not(p0)
            nw34_[int(1)] = ((d_265_v101_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qe"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qe"))) in (d_265_v101_) else p0)
            nw34_[int(2)] = not(p0)
            nw34_[int(3)] = default__.fm5(p2, not(p0), p0, globalState)
            nw34_[int(4)] = p0
            nw34_[int(5)] = (846) < (p2)
            nw34_[int(6)] = not (p0) or (p0)
            nw34_[int(7)] = p0
            nw34_[int(8)] = (_dafny.MultiSet([d_268_v104_])).isdisjoint(d_269_v105_)
            nw34_[int(9)] = p0
            nw34_[int(10)] = p0
            nw34_[int(11)] = p0
            nw34_[int(12)] = p0
            nw34_[int(13)] = not(p0)
            nw34_[int(14)] = (_dafny.MultiSet([d_264_v100_, d_264_v100_, (d_264_v100_).set(default__.safeIndex(p2, len(d_264_v100_)), d_270_v106_)])).issubset(d_271_v107_)
            d_272_v108_ = nw34_
            (globalState).f1 = d_272_v108_
            (globalState).f25 = p0
            d_273_v109_: C3
            nw35_ = C3()
            nw35_.ctor__(p2, default__.fm4(d_270_v106_, globalState), (d_264_v100_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yygjxck"))))
            d_273_v109_ = nw35_
            (globalState).f23 = (d_264_v100_) + (d_264_v100_)
            d_274_v110_: _dafny.MultiSet
            d_274_v110_ = _dafny.MultiSet([(d_264_v100_)[default__.safeIndex((d_273_v109_).f34, len(d_264_v100_))], d_270_v106_])
            d_274_v110_ = (d_274_v110_) | (d_274_v110_)
        elif True:
            d_275_v111_: _dafny.Seq
            d_275_v111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbsvhj"))
            d_276_v112_: _dafny.Map
            d_276_v112_ = _dafny.Map({p0: _dafny.Set({p0, p0})})
            d_277_v113_: _dafny.Set
            d_277_v113_ = _dafny.Set({p0})
            d_278_v114_: _dafny.Array
            nw36_ = _dafny.Array(False, 8)
            d_278_v114_ = nw36_
            d_279_v115_: _dafny.Seq
            d_279_v115_ = _dafny.SeqWithoutIsStrInference([d_278_v114_, d_278_v114_])
            d_280_v116_: _dafny.Map
            d_280_v116_ = _dafny.Map({p0: p0})
            d_281_v117_: str
            d_281_v117_ = _dafny.CodePoint('c')
            d_282_v119_: _dafny.Set
            d_282_v119_ = _dafny.Set({d_281_v117_, d_281_v117_})
            d_283_v120_: _dafny.Array
            nw37_ = _dafny.Array(None, 25)
            nw37_[int(0)] = p2
            nw37_[int(1)] = p2
            nw37_[int(2)] = len(_dafny.SeqWithoutIsStrInference([len(d_275_v111_), p2]))
            nw37_[int(3)] = len(((d_276_v112_).set(False, d_277_v113_)) | (_dafny.Map({p0: d_277_v113_})))
            nw37_[int(4)] = p2
            nw37_[int(5)] = (0) - (p2)
            nw37_[int(6)] = (p2) - (len(d_275_v111_))
            nw37_[int(7)] = p2
            nw37_[int(8)] = (p2) * (p2)
            nw37_[int(9)] = p2
            nw37_[int(10)] = -165
            nw37_[int(11)] = (p2) + ((0) - (len(d_277_v113_)))
            nw37_[int(12)] = (p2) * (p2)
            nw37_[int(13)] = default__.safeModuloInt(p2, len(d_275_v111_))
            nw37_[int(14)] = p2
            nw37_[int(15)] = 182
            nw37_[int(16)] = len(_dafny.Map({not(p0): (d_279_v115_)[default__.safeIndex(len(d_280_v116_), len(d_279_v115_))]}))
            nw37_[int(17)] = 77
            nw37_[int(18)] = default__.fm4(d_281_v117_, globalState)
            nw37_[int(19)] = (p2) - (p2)
            nw37_[int(20)] = len(_dafny.Map({_dafny.SeqWithoutIsStrInference([p0]): d_281_v117_}))
            nw37_[int(21)] = (default__.fm4(d_281_v117_, globalState)) + (p2)
            nw37_[int(22)] = 252
            def iife25_():
                coll25_ = _dafny.Map()
                compr_25_: _dafny.Set
                for compr_25_ in (_dafny.Map({d_282_v119_: True})).keys.Elements:
                    d_284_v118_: _dafny.Set = compr_25_
                    if (d_284_v118_) in (_dafny.Map({d_282_v119_: True})):
                        coll25_[d_284_v118_] = d_281_v117_
                return _dafny.Map(coll25_)
            nw37_[int(23)] = len(iife25_()
            )
            nw37_[int(24)] = p2
            d_283_v120_ = nw37_
            d_283_v120_ = d_283_v120_
            d_285_v121_: _dafny.Map
            d_285_v121_ = _dafny.Map({not(p0): p2})
            d_286_v122_: _dafny.MultiSet
            d_286_v122_ = _dafny.MultiSet([p2, (0) - (len(d_285_v121_))])
            d_287_v123_: C5
            nw38_ = C5()
            nw38_.ctor__(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwxwo"))).set(default__.safeIndex((0) - (p2), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwxwo")))), _dafny.CodePoint('k'))), ((d_286_v122_)[807] if (807) in (d_286_v122_) else default__.fm4(_dafny.CodePoint('m'), globalState)), 320, ((d_275_v111_).set(default__.safeIndex(p2, len(d_275_v111_)), d_281_v117_)) + (d_275_v111_))
            d_287_v123_ = nw38_
            index34_ = default__.safeIndex(360, (d_283_v120_).length(0))
            (d_283_v120_)[index34_] = d_287_v123_.f29
            index35_ = default__.safeIndex(360, (d_283_v120_).length(0))
            (d_283_v120_)[index35_] = d_287_v123_.f29
            d_288_v124_: D10
            d_288_v124_ = D10_DC32(p0, (d_283_v120_)[default__.safeIndex(360, (d_283_v120_).length(0))], p0)
            index36_ = default__.safeIndex(360, (d_283_v120_).length(0))
            rhs28_ = (len((d_275_v111_).set(default__.safeIndex(p2, len(d_275_v111_)), d_281_v117_)) if (d_288_v124_).cf64 else p2)
            rhs29_ = p0
            rhs30_ = p0
            lhs28_ = d_283_v120_
            lhs29_ = default__.safeIndex(360, (d_283_v120_).length(0))
            lhs30_ = globalState
            lhs31_ = globalState
            lhs28_[lhs29_] = rhs28_
            lhs30_.f9 = rhs29_
            lhs31_.f9 = rhs30_
            d_289_v125_: _dafny.Seq
            d_289_v125_ = _dafny.SeqWithoutIsStrInference([d_283_v120_, d_283_v120_, d_283_v120_, d_283_v120_, d_283_v120_])
            d_289_v125_ = ((d_289_v125_).set(default__.safeIndex(len(d_275_v111_), len(d_289_v125_)), d_283_v120_)) + (d_289_v125_)
        d_290_v126_: _dafny.Array
        nw39_ = _dafny.Array(False, 29)
        d_290_v126_ = nw39_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_290_v126_).length(0)):
            d_291_i13_: int = guard_loop_0_
            if (True) and (((0) <= (d_291_i13_)) and ((d_291_i13_) < ((d_290_v126_).length(0)))):
                (d_290_v126_)[(d_291_i13_)] = (default__.safeModuloInt(p2, (_dafny.MultiSet([p2])).cardinality)) <= (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_292_i14_ in range(default__.abs(-743))])))
        d_293_v127_: _dafny.Seq
        d_293_v127_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvwmon"))
        d_294_v128_: _dafny.Map
        d_294_v128_ = _dafny.Map({p2: p2})
        d_295_v129_: T2
        nw40_ = C4()
        nw40_.ctor__(p0, (0) - (p2), d_293_v127_, d_294_v128_)
        d_295_v129_ = nw40_
        d_296_v130_: _dafny.Map
        d_296_v130_ = _dafny.Map({d_295_v129_: p0})
        d_297_v131_: C6
        nw41_ = C6()
        nw41_.ctor__(default__.fm64(d_293_v127_, globalState), len(d_296_v130_), d_293_v127_)
        d_297_v131_ = nw41_
        r0 = (d_297_v131_ if p0 else d_297_v131_)
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_298_v0_: _dafny.Array
        nw42_ = _dafny.Array(False, 12)
        d_298_v0_ = nw42_
        d_299_v1_: _dafny.Array
        def lambda10_(d_300_i0_):
            return _dafny.Set({False, False})

        init5_ = lambda10_
        nw43_ = _dafny.Array(None, 22)
        for i0_5_ in range(nw43_.length(0)):
            nw43_[i0_5_] = init5_(i0_5_)
        d_299_v1_ = nw43_
        d_301_v2_: bool
        d_301_v2_ = True
        d_302_v3_: _dafny.Seq
        d_302_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uggid"))
        d_303_v4_: int
        d_303_v4_ = 655
        d_304_v5_: _dafny.Map
        d_304_v5_ = _dafny.Map({d_303_v4_: d_303_v4_})
        d_305_v6_: _dafny.Map
        d_305_v6_ = _dafny.Map({d_301_v2_: d_304_v5_})
        d_306_v7_: _dafny.Array
        def lambda11_(d_307_v4_):
            def lambda12_(d_308_i1_):
                return (d_308_i1_) - (d_307_v4_)

            return lambda12_

        init6_ = lambda11_(d_303_v4_)
        nw44_ = _dafny.Array(None, 1)
        for i0_6_ in range(nw44_.length(0)):
            nw44_[i0_6_] = init6_(i0_6_)
        d_306_v7_ = nw44_
        d_309_globalState_: GlobalState
        nw45_ = GlobalState()
        nw45_.ctor__(-603, d_298_v0_, 10, False, False, -559, _dafny.CodePoint('a'), d_299_v1_, -981, True, False, False, 800, (d_302_v3_ if d_301_v2_ else d_302_v3_), False, 148, True, 546, d_305_v6_, d_306_v7_, 911, 948, _dafny.CodePoint('g'), d_302_v3_, 30, False)
        d_309_globalState_ = nw45_
        d_310_v8_: _dafny.Array
        nw46_ = _dafny.Array(_dafny.Array(None, 0), 10)
        d_310_v8_ = nw46_
        d_310_v8_ = (d_310_v8_ if (d_302_v3_) <= (d_302_v3_) else d_310_v8_)
        d_311_v9_: _dafny.Seq
        d_311_v9_ = _dafny.SeqWithoutIsStrInference([(0) - (d_303_v4_), 731])
        d_312_v10_: _dafny.Set
        d_312_v10_ = _dafny.Set({d_301_v2_})
        (d_309_globalState_).f17 = (d_311_v9_)[default__.safeIndex(len((default__.fm0(d_301_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlj")), d_312_v10_, d_303_v4_, d_309_globalState_)).set(default__.safeIndex(d_303_v4_, len(default__.fm0(d_301_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlj")), d_312_v10_, d_303_v4_, d_309_globalState_))), d_301_v2_)), len(d_311_v9_))]
        d_313_v11_: D0
        d_313_v11_ = D0_DC2(d_301_v2_)
        source10_ = d_313_v11_
        if source10_.is_DC0:
            d_314___mcc_h0_ = source10_.cf0
            d_315_cf0_ = d_314___mcc_h0_
            d_311_v9_ = d_311_v9_
            d_316_v12_: str
            d_316_v12_ = _dafny.CodePoint('v')
            (d_309_globalState_).f24 = (d_303_v4_) - (len(((d_302_v3_).set(default__.safeIndex(d_303_v4_, len(d_302_v3_)), d_316_v12_) if d_301_v2_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itarlcvfg")))))
            d_317_v13_: _dafny.MultiSet
            d_317_v13_ = _dafny.MultiSet([(0) - (d_303_v4_), d_315_cf0_, d_315_cf0_])
            d_318_v14_: C5
            nw47_ = C5()
            nw47_.ctor__(d_315_cf0_, ((d_317_v13_)[-278] if (-278) in (d_317_v13_) else -444), default__.fm4(d_316_v12_, d_309_globalState_), d_302_v3_)
            d_318_v14_ = nw47_
            (d_309_globalState_).f24 = default__.safeDivisionInt(d_318_v14_.f29, d_318_v14_.f29)
        elif source10_.is_DC1:
            d_319___mcc_h1_ = source10_.cf1
            d_320___mcc_h2_ = source10_.cf2
            d_321___mcc_h3_ = source10_.cf3
            d_322_cf3_ = d_321___mcc_h3_
            d_323_cf2_ = d_320___mcc_h2_
            d_324_cf1_ = d_319___mcc_h1_
            d_301_v2_ = (d_301_v2_) == (d_301_v2_)
            index37_ = default__.safeIndex(895, (d_298_v0_).length(0))
            (d_298_v0_)[index37_] = False
            d_325_v15_: _dafny.Seq
            d_325_v15_ = _dafny.SeqWithoutIsStrInference([d_301_v2_, not(d_301_v2_), not(d_301_v2_), (d_323_cf2_) <= (d_322_cf3_)])
            index38_ = default__.safeIndex(895, (d_298_v0_).length(0))
            (d_298_v0_)[index38_] = not((d_325_v15_)[default__.safeIndex(d_303_v4_, len(d_325_v15_))])
            d_326_v16_: _dafny.Map
            d_326_v16_ = _dafny.Map({d_303_v4_: (d_324_cf1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sikrj")))})
            d_326_v16_ = (d_326_v16_).set((d_323_cf2_ if True else d_303_v4_), d_302_v3_)
            if (d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))]:
                d_327_v17_: C7
                nw48_ = C7()
                nw48_.ctor__()
                d_327_v17_ = nw48_
                d_328_v18_: _dafny.Seq
                d_328_v18_ = _dafny.SeqWithoutIsStrInference([d_304_v5_, d_304_v5_, d_304_v5_])
                d_329_v19_: C4
                nw49_ = C4()
                nw49_.ctor__((d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))], d_323_cf2_, d_302_v3_, (d_328_v18_)[default__.safeIndex(d_322_cf3_, len(d_328_v18_))])
                d_329_v19_ = nw49_
                d_330_v20_: _dafny.Map
                d_330_v20_ = _dafny.Map({(d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))]: 485})
                d_331_v21_: _dafny.MultiSet
                d_331_v21_ = _dafny.MultiSet([(d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))]])
                d_332_v22_: str
                d_332_v22_ = _dafny.CodePoint('h')
                d_333_v23_: _dafny.Set
                d_333_v23_ = _dafny.Set({d_332_v22_})
                (d_309_globalState_).f17 = (((d_330_v20_)[d_301_v2_] if (d_301_v2_) in (d_330_v20_) else ((d_331_v21_)[(d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))]] if ((d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))]) in (d_331_v21_) else len(d_333_v23_)))) + ((0) - (d_323_cf2_))
                (d_309_globalState_).f9 = d_301_v2_
                d_334_v24_: C5
                nw50_ = C5()
                nw50_.ctor__(d_323_cf2_, d_322_cf3_, d_323_cf2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvh")))
                d_334_v24_ = nw50_
                d_335_v25_: _dafny.Map
                d_335_v25_ = _dafny.Map({False: d_334_v24_})
                d_336_v26_: C5
                nw51_ = C5()
                nw51_.ctor__(d_303_v4_, len(d_335_v25_), d_334_v24_.f30, _dafny.SeqWithoutIsStrInference([d_332_v22_ for d_337_i2_ in range(default__.abs(301))]))
                d_336_v26_ = nw51_
                d_338_v27_: _dafny.Set
                d_338_v27_ = _dafny.Set({d_334_v24_})
                d_339_v28_: C4
                nw52_ = C4()
                nw52_.ctor__((d_338_v27_).issubset(d_338_v27_), d_323_cf2_, d_302_v3_, d_304_v5_)
                d_339_v28_ = nw52_
                d_340_v29_: _dafny.Map
                d_340_v29_ = _dafny.Map({default__.safeDivisionInt(d_336_v26_.f30, d_303_v4_): d_334_v24_})
                d_341_v30_: C3
                nw53_ = C3()
                nw53_.ctor__(d_303_v4_, d_336_v26_.f29, (d_302_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlblnafu"))))
                d_341_v30_ = nw53_
                d_342_v31_: _dafny.Array
                def lambda13_(d_343_v2_):
                    def lambda14_(d_344_i3_):
                        return d_343_v2_

                    return lambda14_

                init7_ = lambda13_(d_301_v2_)
                nw54_ = _dafny.Array(None, 14)
                for i0_7_ in range(nw54_.length(0)):
                    nw54_[i0_7_] = init7_(i0_7_)
                d_342_v31_ = nw54_
                d_345_v32_: D11
                d_345_v32_ = D11_DC34(d_342_v31_, (d_329_v19_).f32)
                d_346_v33_: _dafny.Seq
                d_346_v33_ = _dafny.SeqWithoutIsStrInference([d_345_v32_, D11_DC34(d_342_v31_, d_301_v2_), d_345_v32_, D11_DC34(d_342_v31_, (d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))]), d_345_v32_])
                d_347_v34_: _dafny.Set
                d_347_v34_ = _dafny.Set({d_346_v33_})
                rhs31_ = (d_312_v10_) == (_dafny.Set({(d_329_v19_).f32}))
                rhs32_ = (d_339_v28_ if (d_347_v34_).issubset(d_347_v34_) else d_339_v28_)
                rhs33_ = d_334_v24_.f29
                rhs34_ = d_340_v29_
                rhs35_ = d_341_v30_
                lhs32_ = d_309_globalState_
                lhs33_ = d_309_globalState_
                lhs32_.f9 = rhs31_
                d_339_v28_ = rhs32_
                lhs33_.f17 = rhs33_
                d_340_v29_ = rhs34_
                d_341_v30_ = rhs35_
            elif True:
                d_348_v35_: _dafny.Array
                nw55_ = _dafny.Array(D2.default()(), 19)
                d_348_v35_ = nw55_
                d_349_v36_: _dafny.Map
                d_349_v36_ = _dafny.Map({d_303_v4_: d_348_v35_})
                d_349_v36_ = (d_349_v36_).set(d_303_v4_, d_348_v35_)
                d_350_v37_: _dafny.Array
                nw56_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                d_350_v37_ = nw56_
                d_350_v37_ = d_350_v37_
                d_351_v38_: _dafny.Map
                d_351_v38_ = _dafny.Map({(d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))]: False})
                d_352_v39_: _dafny.Seq
                d_352_v39_ = _dafny.SeqWithoutIsStrInference([d_351_v38_])
                d_353_v40_: _dafny.Set
                d_353_v40_ = _dafny.Set({len((d_352_v39_)[default__.safeIndex(d_322_cf3_, len(d_352_v39_))])})
                (d_309_globalState_).f13 = default__.fm36(((d_311_v9_)[default__.safeIndex(len(d_353_v40_), len(d_311_v9_))]) < (d_323_cf2_), (d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))], default__.fm5(d_303_v4_, d_301_v2_, (d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))], d_309_globalState_), d_309_globalState_)
                d_354_v41_: _dafny.Map
                d_354_v41_ = _dafny.Map({(d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))]: d_303_v4_})
                d_355_v42_: str
                d_355_v42_ = _dafny.CodePoint('r')
                d_356_v44_: _dafny.Map
                def iife26_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in _dafny.IntegerRange(42, 708):
                        d_357_v43_: int = compr_26_
                        if ((42) <= (d_357_v43_)) and ((d_357_v43_) < (708)):
                            coll26_[default__.safeModuloInt(d_357_v43_, (_dafny.MultiSet([(D19_DC53(d_322_cf3_, d_301_v2_)).cf101])).cardinality)] = D1_DC5(d_312_v10_)
                    return _dafny.Map(coll26_)
                d_356_v44_ = iife26_()
                
                d_358_v45_: D8
                d_358_v45_ = D8_DC25(d_301_v2_, d_354_v41_, d_355_v42_, d_356_v44_, d_303_v4_)
                d_359_v47_: _dafny.Map
                d_359_v47_ = _dafny.Map({d_323_cf2_: d_355_v42_})
                d_360_v48_: _dafny.MultiSet
                d_360_v48_ = _dafny.MultiSet([(d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))]])
                d_361_v49_: D26
                def iife27_():
                    coll27_ = _dafny.Map()
                    compr_27_: int
                    for compr_27_ in _dafny.IntegerRange(986, 856):
                        d_362_v46_: int = compr_27_
                        if ((986) <= (d_362_v46_)) and ((d_362_v46_) < (856)):
                            coll27_[(d_362_v46_) - (d_322_cf3_)] = (d_298_v0_)[default__.safeIndex(895, (d_298_v0_).length(0))]
                    return _dafny.Map(coll27_)
                d_361_v49_ = D26_DC79(len((d_358_v45_).cf50), len((iife27_()
).set(default__.fm4(((d_359_v47_)[(d_360_v48_).cardinality] if ((d_360_v48_).cardinality) in (d_359_v47_) else (d_324_cf1_)[default__.safeIndex(d_303_v4_, len(d_324_cf1_))]), d_309_globalState_), d_301_v2_)), d_301_v2_)
                d_363_v50_: C6
                out0_: C6
                out0_ = default__.m8((d_361_v49_).cf152, d_325_v15_, len(d_324_cf1_), d_325_v15_, d_309_globalState_)
                d_363_v50_ = out0_
                d_364_v51_: T1
                nw57_ = C5()
                nw57_.ctor__(d_322_cf3_, 775, d_322_cf3_, d_324_cf1_)
                d_364_v51_ = nw57_
                d_365_v52_: _dafny.MultiSet
                d_365_v52_ = _dafny.MultiSet([d_303_v4_, len(_dafny.Set({d_364_v51_})), (d_363_v50_).fm7(720, d_301_v2_, d_309_globalState_)])
                d_366_v53_: _dafny.Map
                d_366_v53_ = _dafny.Map({d_365_v52_: default__.fm20(d_301_v2_, d_355_v42_, 363, d_355_v42_, d_309_globalState_)})
                (d_309_globalState_).f24 = len(d_366_v53_)
        elif source10_.is_DC2:
            d_367___mcc_h4_ = source10_.cf4
            d_368_cf4_ = d_367___mcc_h4_
            (d_309_globalState_).f25 = d_368_cf4_
            (d_309_globalState_).f8 = d_303_v4_
            rhs36_ = d_301_v2_
            rhs37_ = d_303_v4_
            rhs38_ = d_303_v4_
            rhs39_ = d_301_v2_
            lhs34_ = d_309_globalState_
            lhs35_ = d_309_globalState_
            lhs36_ = d_309_globalState_
            lhs37_ = d_309_globalState_
            lhs34_.f9 = rhs36_
            lhs35_.f17 = rhs37_
            lhs36_.f0 = rhs38_
            lhs37_.f9 = rhs39_
            d_301_v2_ = default__.fm5(default__.safeModuloInt(d_303_v4_, d_303_v4_), d_301_v2_, not (d_368_cf4_) or (d_368_cf4_), d_309_globalState_)
        elif source10_.is_DC3:
            d_369___mcc_h5_ = source10_.cf5
            d_370___mcc_h6_ = source10_.cf6
            d_371_cf6_ = d_370___mcc_h6_
            d_372_cf5_ = d_369___mcc_h5_
            d_373_v54_: _dafny.Map
            d_373_v54_ = _dafny.Map({d_303_v4_: d_301_v2_})
            d_374_v55_: D11
            d_374_v55_ = D11_DC34(d_298_v0_, d_301_v2_)
            d_375_v56_: _dafny.Map
            d_375_v56_ = _dafny.Map({(d_373_v54_) | (d_373_v54_): d_374_v55_})
            d_375_v56_ = (d_375_v56_).set(d_373_v54_, d_374_v55_)
            d_376_v57_: _dafny.Array
            nw58_ = _dafny.Array(None, 7)
            nw58_[int(0)] = d_306_v7_
            nw58_[int(1)] = d_306_v7_
            nw58_[int(2)] = d_306_v7_
            nw58_[int(3)] = d_306_v7_
            nw58_[int(4)] = d_306_v7_
            nw58_[int(5)] = d_306_v7_
            nw58_[int(6)] = d_306_v7_
            d_376_v57_ = nw58_
            index39_ = default__.safeIndex(503, (d_376_v57_).length(0))
            (d_376_v57_)[index39_] = d_306_v7_
            d_377_v58_: _dafny.Map
            d_377_v58_ = _dafny.Map({d_303_v4_: D1_DC5(d_312_v10_)})
            d_378_v59_: _dafny.Map
            d_378_v59_ = d_377_v58_
            index40_ = default__.safeIndex(503, (d_376_v57_).length(0))
            rhs40_ = d_306_v7_
            rhs41_ = (0) - ((d_372_cf5_ if d_301_v2_ else (0) - (default__.safeDivisionInt((0) - (len(_dafny.Map({d_378_v59_: -589}))), d_371_cf6_))))
            rhs42_ = (d_312_v10_) - (_dafny.Set({d_301_v2_}))
            rhs43_ = d_301_v2_
            lhs38_ = d_376_v57_
            lhs39_ = default__.safeIndex(503, (d_376_v57_).length(0))
            lhs40_ = d_309_globalState_
            lhs38_[lhs39_] = rhs40_
            d_303_v4_ = rhs41_
            d_312_v10_ = rhs42_
            lhs40_.f25 = rhs43_
            d_379_v60_: _dafny.Array
            nw59_ = _dafny.Array(None, 13)
            d_379_v60_ = nw59_
            d_379_v60_ = d_379_v60_
            d_380_v61_: _dafny.Seq
            d_380_v61_ = _dafny.SeqWithoutIsStrInference([d_301_v2_])
            d_381_v62_: C5
            nw60_ = C5()
            nw60_.ctor__(d_303_v4_, len((d_380_v61_) + (d_380_v61_)), ((_dafny.MultiSet([-523])).cardinality) * ((0) - (d_371_cf6_)), d_302_v3_)
            d_381_v62_ = nw60_
            d_382_v63_: _dafny.Map
            d_382_v63_ = _dafny.Map({d_301_v2_: d_371_cf6_})
            d_383_v64_: _dafny.Map
            d_383_v64_ = _dafny.Map({d_301_v2_: d_382_v63_})
            d_384_v65_: _dafny.Array
            nw61_ = _dafny.Array(_dafny.MultiSet({}), 26)
            d_384_v65_ = nw61_
            d_385_v66_: D11
            d_385_v66_ = D11_DC33(d_384_v65_)
            d_386_v67_: _dafny.MultiSet
            d_386_v67_ = _dafny.MultiSet([d_385_v66_])
            nw62_ = C5()
            nw62_.ctor__((d_371_cf6_) * (len(d_383_v64_)), ((d_386_v67_)[D11_DC33(d_384_v65_)] if (D11_DC33(d_384_v65_)) in (d_386_v67_) else d_371_cf6_), d_303_v4_, d_302_v3_)
            d_381_v62_ = nw62_
        elif True:
            d_387___mcc_h7_ = source10_.cf7
            d_388_cf7_ = d_387___mcc_h7_
            d_389_v68_: _dafny.Map
            d_389_v68_ = _dafny.Map({d_303_v4_: True})
            d_390_v69_: str
            d_390_v69_ = _dafny.CodePoint('q')
            d_391_v70_: _dafny.Map
            d_391_v70_ = _dafny.Map({d_301_v2_: default__.fm4(d_390_v69_, d_309_globalState_)})
            d_389_v68_ = (d_389_v68_).set((len(d_391_v70_)) - (-537), d_301_v2_)
            d_392_v71_: C1
            nw63_ = C1()
            nw63_.ctor__()
            d_392_v71_ = nw63_
            d_392_v71_ = d_392_v71_
            d_393_v72_: C2
            nw64_ = C2()
            nw64_.ctor__(d_311_v9_, d_304_v5_)
            d_393_v72_ = nw64_
            d_394_v73_: _dafny.Map
            d_394_v73_ = _dafny.Map({not(d_301_v2_): d_301_v2_})
            d_395_v74_: D0
            d_395_v74_ = D0_DC0(-285)
            d_396_v75_: _dafny.Map
            d_396_v75_ = _dafny.Map({len(d_394_v73_): d_395_v74_})
            d_397_v76_: D19
            d_397_v76_ = D19_DC52(d_396_v75_)
            source11_ = d_397_v76_
            if source11_.is_DC53:
                d_398___mcc_h8_ = source11_.cf100
                d_399___mcc_h9_ = source11_.cf101
                d_400_cf101_ = d_399___mcc_h9_
                d_401_cf100_ = d_398___mcc_h8_
                d_402_v77_: _dafny.Array
                nw65_ = _dafny.Array(_dafny.Seq({}), 21)
                d_402_v77_ = nw65_
                index41_ = default__.safeIndex(644, (d_402_v77_).length(0))
                (d_402_v77_)[index41_] = default__.fm45(d_401_cf100_, d_391_v70_, d_304_v5_, d_309_globalState_)
                index42_ = default__.safeIndex(644, (d_402_v77_).length(0))
                (d_402_v77_)[index42_] = (d_311_v9_) + ((d_393_v72_).f33)
                d_403_v78_: _dafny.Seq
                d_403_v78_ = _dafny.SeqWithoutIsStrInference([d_311_v9_, (d_393_v72_).f33])
                d_404_v79_: C2
                nw66_ = C2()
                nw66_.ctor__((d_403_v78_)[default__.safeIndex(d_303_v4_, len(d_403_v78_))], default__.fm40(d_400_cf101_, d_401_cf100_, d_389_v68_, False, d_309_globalState_))
                d_404_v79_ = nw66_
                d_405_v80_: D2
                d_405_v80_ = D2_DC8(d_301_v2_, d_303_v4_, d_401_cf100_)
                d_406_v81_: int
                out1_: int
                out1_ = (d_393_v72_).m4(d_303_v4_, d_405_v80_, d_309_globalState_)
                d_406_v81_ = out1_
                d_407_v82_: D18
                d_407_v82_ = D18_DC50(((d_393_v72_).f33)[default__.safeIndex(d_406_v81_, len((d_393_v72_).f33))], not(((d_394_v73_)[d_301_v2_] if (d_301_v2_) in (d_394_v73_) else False)), d_311_v9_, d_390_v69_, d_401_cf100_)
                (d_309_globalState_).f9 = (d_407_v82_).cf94
            elif True:
                d_408___mcc_h10_ = source11_.cf99
                d_409_cf99_ = d_408___mcc_h10_
                d_410_v83_: _dafny.Set
                d_411_v84_: _dafny.Map
                d_412_v85_: str
                out2_: _dafny.Set
                out3_: _dafny.Map
                out4_: str
                out2_, out3_, out4_ = (d_393_v72_).m5(d_298_v0_, (d_391_v70_).set(not(d_301_v2_), len(_dafny.SeqWithoutIsStrInference([d_301_v2_]))), d_309_globalState_)
                d_410_v83_ = out2_
                d_411_v84_ = out3_
                d_412_v85_ = out4_
                d_413_v86_: int
                d_414_v87_: _dafny.MultiSet
                d_415_v88_: _dafny.Array
                out5_: int
                out6_: _dafny.MultiSet
                out7_: _dafny.Array
                out5_, out6_, out7_ = (d_393_v72_).m2(d_301_v2_, d_313_v11_, d_309_globalState_)
                d_413_v86_ = out5_
                d_414_v87_ = out6_
                d_415_v88_ = out7_
                (d_309_globalState_).f24 = d_303_v4_
                (d_309_globalState_).f9 = d_301_v2_
        d_416_v89_: D7
        d_416_v89_ = D7_DC19(d_303_v4_, True)
        d_417_v90_: D7
        d_417_v90_ = D7_DC21(d_416_v89_)
        d_418_v91_: _dafny.Seq
        d_418_v91_ = _dafny.SeqWithoutIsStrInference([D7_DC20(d_301_v2_, not(d_301_v2_)), d_416_v89_])
        d_419_v92_: _dafny.Seq
        d_419_v92_ = _dafny.SeqWithoutIsStrInference([d_417_v90_, D7_DC21((d_418_v91_)[default__.safeIndex(d_303_v4_, len(d_418_v91_))])])
        d_420_v93_: _dafny.Seq
        d_420_v93_ = _dafny.SeqWithoutIsStrInference([d_419_v92_])
        d_421_v94_: D9
        d_421_v94_ = D9_DC29(D9_DC27(d_420_v93_))
        d_422_v95_: _dafny.Map
        d_422_v95_ = _dafny.Map({d_303_v4_: d_421_v94_})
        d_423_v96_: C5
        nw67_ = C5()
        nw67_.ctor__(d_303_v4_, len(d_422_v95_), d_303_v4_, d_302_v3_)
        d_423_v96_ = nw67_
        d_424_v97_: _dafny.Array
        nw68_ = _dafny.Array(_dafny.Map({}), 1)
        d_424_v97_ = nw68_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_424_v97_).length(0)):
            d_425_i4_: int = guard_loop_1_
            if (True) and (((0) <= (d_425_i4_)) and ((d_425_i4_) < ((d_424_v97_).length(0)))):
                (d_424_v97_)[(d_425_i4_)] = ((_dafny.Map({False: d_423_v96_.f29})) | (_dafny.Map({d_301_v2_: len(d_302_v3_)})) if d_301_v2_ else (_dafny.Map({d_301_v2_: d_423_v96_.f30}) if d_301_v2_ else _dafny.Map({d_301_v2_: (_dafny.MultiSet([d_423_v96_.f30])).cardinality})))
        d_426_v98_: int
        d_427_v99_: _dafny.MultiSet
        d_428_v100_: _dafny.Array
        out8_: int
        out9_: _dafny.MultiSet
        out10_: _dafny.Array
        out8_, out9_, out10_ = (d_423_v96_).m2(default__.fm24(d_301_v2_, d_301_v2_, d_301_v2_, d_423_v96_.f30, d_309_globalState_), d_313_v11_, d_309_globalState_)
        d_426_v98_ = out8_
        d_427_v99_ = out9_
        d_428_v100_ = out10_
        d_429_v101_: _dafny.Map
        d_429_v101_ = _dafny.Map({d_301_v2_: d_423_v96_.f29})
        d_430_v102_: _dafny.Seq
        d_430_v102_ = _dafny.SeqWithoutIsStrInference([d_429_v101_, d_429_v101_, d_429_v101_, d_429_v101_, d_429_v101_])
        d_431_v103_: _dafny.Map
        d_431_v103_ = _dafny.Map({d_301_v2_: len(d_430_v102_)})
        d_432_v104_: _dafny.Map
        d_432_v104_ = _dafny.Map({d_431_v103_: _dafny.CodePoint('e')})
        d_433_v105_: C6
        nw69_ = C6()
        nw69_.ctor__(d_432_v104_, (359) - (len(d_302_v3_)), d_302_v3_)
        d_433_v105_ = nw69_
        d_434_v106_: str
        d_434_v106_ = _dafny.CodePoint('l')
        hi1_ = d_426_v98_
        for d_435_i5_ in range(default__.fm4(d_434_v106_, d_309_globalState_), hi1_):
            index43_ = default__.safeIndex(805, (d_298_v0_).length(0))
            (d_298_v0_)[index43_] = d_301_v2_
            d_436_v107_: _dafny.Seq
            d_436_v107_ = _dafny.SeqWithoutIsStrInference([d_301_v2_, False])
            d_437_v108_: _dafny.MultiSet
            d_437_v108_ = _dafny.MultiSet([d_423_v96_.f30, len(d_436_v107_), d_423_v96_.f29, d_423_v96_.f29, d_423_v96_.f29])
            d_438_v109_: _dafny.Seq
            d_438_v109_ = _dafny.SeqWithoutIsStrInference([d_301_v2_, (d_433_v105_).fm8(d_437_v108_, d_301_v2_, d_423_v96_.f29, len(d_302_v3_), d_309_globalState_)])
            d_439_v110_: T1
            nw70_ = C6()
            nw70_.ctor__(d_432_v104_, len(d_302_v3_), d_302_v3_)
            d_439_v110_ = nw70_
            d_440_v111_: _dafny.Map
            d_440_v111_ = _dafny.Map({d_439_v110_: (0) - (d_423_v96_.f29)})
            d_441_v112_: _dafny.Seq
            d_441_v112_ = _dafny.SeqWithoutIsStrInference([d_439_v110_])
            index44_ = default__.safeIndex(805, (d_298_v0_).length(0))
            rhs44_ = (True) in (d_438_v109_)
            rhs45_ = d_304_v5_
            rhs46_ = ((d_433_v105_).fm7(d_303_v4_, d_301_v2_, d_309_globalState_)) > (len((d_440_v111_) | (_dafny.Map({(d_441_v112_)[default__.safeIndex(d_423_v96_.f30, len(d_441_v112_))]: (d_439_v110_).f26}))))
            lhs41_ = d_298_v0_
            lhs42_ = default__.safeIndex(805, (d_298_v0_).length(0))
            lhs43_ = d_309_globalState_
            lhs41_[lhs42_] = rhs44_
            d_304_v5_ = rhs45_
            lhs43_.f9 = rhs46_
            d_442_v113_: _dafny.Array
            nw71_ = _dafny.Array(_dafny.MultiSet({}), 2)
            d_442_v113_ = nw71_
            d_443_v114_: _dafny.MultiSet
            d_443_v114_ = _dafny.MultiSet([(d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))], d_301_v2_])
            index45_ = default__.safeIndex(694, (d_442_v113_).length(0))
            (d_442_v113_)[index45_] = (d_443_v114_).intersection(d_443_v114_)
            d_444_v115_: _dafny.Map
            d_444_v115_ = _dafny.Map({d_435_i5_: (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]})
            index46_ = default__.safeIndex(805, (d_298_v0_).length(0))
            index47_ = default__.safeIndex(694, (d_442_v113_).length(0))
            index48_ = default__.safeIndex(805, (d_298_v0_).length(0))
            rhs47_ = ((d_444_v115_)[((d_433_v105_).fm9((d_439_v110_).f26, len((d_439_v110_).f27), (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))], d_309_globalState_)) + (d_423_v96_.f30)] if (((d_433_v105_).fm9((d_439_v110_).f26, len((d_439_v110_).f27), (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))], d_309_globalState_)) + (d_423_v96_.f30)) in (d_444_v115_) else d_301_v2_)
            rhs48_ = (d_443_v114_) - ((_dafny.MultiSet([(d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]])) - (d_443_v114_))
            rhs49_ = d_301_v2_
            lhs44_ = d_298_v0_
            lhs45_ = default__.safeIndex(805, (d_298_v0_).length(0))
            lhs46_ = d_442_v113_
            lhs47_ = default__.safeIndex(694, (d_442_v113_).length(0))
            lhs48_ = d_298_v0_
            lhs49_ = default__.safeIndex(805, (d_298_v0_).length(0))
            lhs44_[lhs45_] = rhs47_
            lhs46_[lhs47_] = rhs48_
            lhs48_[lhs49_] = rhs49_
            d_445_v116_: _dafny.MultiSet
            d_445_v116_ = _dafny.MultiSet([d_306_v7_, d_306_v7_, d_306_v7_, d_306_v7_])
            if ((d_445_v116_) | (d_445_v116_)).issubset(d_445_v116_):
                index49_ = default__.safeIndex(820, (d_306_v7_).length(0))
                (d_306_v7_)[index49_] = d_426_v98_
                d_446_v117_: D21
                d_446_v117_ = D21_DC58(d_423_v96_.f30, (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))], d_310_v8_, (0) - (len(d_436_v107_)), (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))])
                index50_ = default__.safeIndex(820, (d_306_v7_).length(0))
                (d_306_v7_)[index50_] = (d_446_v117_).cf112
                d_447_v118_: _dafny.Array
                nw72_ = _dafny.Array(None, 1)
                d_447_v118_ = nw72_
                d_448_v119_: _dafny.Map
                d_448_v119_ = _dafny.Map({d_447_v118_: d_423_v96_.f29})
                d_449_v120_: T0
                nw73_ = C4()
                nw73_.ctor__(d_301_v2_, d_423_v96_.f29, d_302_v3_, d_304_v5_)
                d_449_v120_ = nw73_
                d_450_v121_: _dafny.Map
                d_450_v121_ = _dafny.Map({d_449_v120_: (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]})
                (d_309_globalState_).f25 = (d_448_v119_) != ((d_448_v119_) | (_dafny.Map({d_447_v118_: len(d_450_v121_)})))
                d_451_v122_: _dafny.Array
                nw74_ = _dafny.Array(int(0), 18)
                d_451_v122_ = nw74_
                d_452_v123_: C7
                nw75_ = C7()
                nw75_.ctor__()
                d_452_v123_ = nw75_
                d_453_v124_: D27
                d_453_v124_ = D27_DC81(d_452_v123_)
                d_454_v125_: _dafny.Array
                nw76_ = _dafny.Array(None, 24)
                nw76_[int(0)] = d_452_v123_
                nw76_[int(1)] = d_452_v123_
                nw76_[int(2)] = d_452_v123_
                nw76_[int(3)] = d_452_v123_
                nw76_[int(4)] = (d_452_v123_ if (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))] else d_452_v123_)
                nw76_[int(5)] = d_452_v123_
                nw76_[int(6)] = d_452_v123_
                nw76_[int(7)] = (d_453_v124_).cf154
                nw76_[int(8)] = d_452_v123_
                nw76_[int(9)] = d_452_v123_
                nw76_[int(10)] = d_452_v123_
                nw76_[int(11)] = d_452_v123_
                nw76_[int(12)] = d_452_v123_
                nw76_[int(13)] = d_452_v123_
                nw76_[int(14)] = d_452_v123_
                nw76_[int(15)] = d_452_v123_
                nw76_[int(16)] = d_452_v123_
                nw76_[int(17)] = d_452_v123_
                nw76_[int(18)] = d_452_v123_
                nw76_[int(19)] = d_452_v123_
                nw76_[int(20)] = d_452_v123_
                nw76_[int(21)] = d_452_v123_
                nw76_[int(22)] = (d_453_v124_).cf154
                nw76_[int(23)] = d_452_v123_
                d_454_v125_ = nw76_
                index51_ = default__.safeIndex(601, (d_454_v125_).length(0))
                (d_454_v125_)[index51_] = d_452_v123_
                d_455_v126_: _dafny.Seq
                d_455_v126_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len((d_439_v110_).f27): 258}), d_304_v5_, (d_304_v5_) | (default__.fm40(d_301_v2_, (d_423_v96_).fm7(-704, d_301_v2_, d_309_globalState_), d_444_v115_, False, d_309_globalState_)), (d_304_v5_) | (d_304_v5_)])
                index52_ = default__.safeIndex(601, (d_454_v125_).length(0))
                def iife28_():
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in (d_311_v9_).Elements:
                        d_456_v127_: int = compr_28_
                        if (d_456_v127_) in (d_311_v9_):
                            coll28_[(d_456_v127_) + ((d_439_v110_).f26)] = 998
                    return _dafny.Map(coll28_)
                rhs50_ = d_451_v122_
                rhs51_ = d_452_v123_
                rhs52_ = ((d_455_v126_) + (_dafny.SeqWithoutIsStrInference([iife28_()
                ]))) + (d_455_v126_)
                lhs50_ = d_454_v125_
                lhs51_ = default__.safeIndex(601, (d_454_v125_).length(0))
                d_451_v122_ = rhs50_
                lhs50_[lhs51_] = rhs51_
                d_455_v126_ = rhs52_
                d_457_v128_: C0
                nw77_ = C0()
                nw77_.ctor__()
                d_457_v128_ = nw77_
                d_458_v129_: _dafny.Seq
                d_458_v129_ = _dafny.SeqWithoutIsStrInference([d_457_v128_, d_457_v128_])
                d_459_v130_: D20
                d_459_v130_ = D20_DC55(d_423_v96_.f30, (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))], d_423_v96_.f29, False)
                d_460_v131_: _dafny.Map
                d_460_v131_ = _dafny.Map({(d_458_v129_).set(default__.safeIndex(d_423_v96_.f30, len(d_458_v129_)), d_457_v128_): _dafny.Map({d_459_v130_: d_423_v96_})})
                d_461_v132_: _dafny.Map
                d_461_v132_ = _dafny.Map({(d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]: d_301_v2_})
                rhs53_ = len(((d_460_v131_) | (d_460_v131_)) | (d_460_v131_))
                rhs54_ = default__.safeModuloInt((len(d_461_v132_) if (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))] else d_423_v96_.f30), d_435_i5_)
                rhs55_ = default__.safeDivisionInt(d_426_v98_, d_423_v96_.f30)
                lhs52_ = d_309_globalState_
                lhs53_ = d_309_globalState_
                lhs54_ = d_309_globalState_
                lhs52_.f0 = rhs53_
                lhs53_.f17 = rhs54_
                lhs54_.f8 = rhs55_
                index53_ = default__.safeIndex(805, (d_298_v0_).length(0))
                (d_298_v0_)[index53_] = ((d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))] if (d_438_v109_)[default__.safeIndex((d_311_v9_)[default__.safeIndex(642, len(d_311_v9_))], len(d_438_v109_))] else (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))])
            elif True:
                pat_let_tv0_ = d_301_v2_
                d_462_v133_: int
                d_463_v134_: _dafny.MultiSet
                d_464_v135_: _dafny.Array
                out11_: int
                out12_: _dafny.MultiSet
                out13_: _dafny.Array
                def iife29_(_pat_let0_0):
                    def iife30_(d_465_dt__update__tmp_h0_):
                        def iife31_(_pat_let1_0):
                            def iife32_(d_466_dt__update_hcf4_h0_):
                                return D0_DC2(d_466_dt__update_hcf4_h0_)
                            return iife32_(_pat_let1_0)
                        return iife31_(pat_let_tv0_)
                    return iife30_(_pat_let0_0)
                out11_, out12_, out13_ = (d_433_v105_).m2((True) in (d_312_v10_), iife29_(d_313_v11_), d_309_globalState_)
                d_462_v133_ = out11_
                d_463_v134_ = out12_
                d_464_v135_ = out13_
                d_467_v136_: C5
                nw78_ = C5()
                nw78_.ctor__(d_423_v96_.f30, default__.safeModuloInt(d_423_v96_.f30, (d_439_v110_).f26), (d_439_v110_).f26, (d_302_v3_ if d_301_v2_ else (d_439_v110_).f27))
                d_467_v136_ = nw78_
                d_468_v137_: _dafny.Seq
                d_468_v137_ = _dafny.SeqWithoutIsStrInference([(d_439_v110_).f27])
                index54_ = default__.safeIndex(717, (d_306_v7_).length(0))
                (d_306_v7_)[index54_] = (len(d_468_v137_)) * (-501)
                index55_ = default__.safeIndex(717, (d_306_v7_).length(0))
                (d_306_v7_)[index55_] = d_423_v96_.f29
                d_469_v138_: D22
                d_469_v138_ = D22_DC64((d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))])
                d_470_v139_: _dafny.Set
                d_470_v139_ = _dafny.Set({d_433_v105_})
                d_471_v140_: _dafny.Seq
                d_471_v140_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_423_v96_.f29: (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]}), default__.fm18(d_435_i5_, d_309_globalState_)])
                d_472_v142_: _dafny.Array
                nw79_ = _dafny.Array(None, 19)
                nw79_[int(0)] = d_444_v115_
                nw79_[int(1)] = d_444_v115_
                nw79_[int(2)] = _dafny.Map({len(d_302_v3_): (d_469_v138_).cf121})
                nw79_[int(3)] = d_444_v115_
                nw79_[int(4)] = (d_444_v115_).set(len(d_470_v139_), (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))])
                nw79_[int(5)] = d_444_v115_
                nw79_[int(6)] = (d_471_v140_)[default__.safeIndex(849, len(d_471_v140_))]
                nw79_[int(7)] = (d_444_v115_) | (d_444_v115_)
                nw79_[int(8)] = d_444_v115_
                nw79_[int(9)] = d_444_v115_
                nw79_[int(10)] = (d_444_v115_ if d_301_v2_ else d_444_v115_)
                nw79_[int(11)] = d_444_v115_
                nw79_[int(12)] = _dafny.Map({d_435_i5_: not((d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))])})
                nw79_[int(13)] = d_444_v115_
                def iife33_():
                    coll29_ = _dafny.Map()
                    compr_29_: int
                    for compr_29_ in _dafny.IntegerRange(149, 667):
                        d_473_v141_: int = compr_29_
                        if ((149) <= (d_473_v141_)) and ((d_473_v141_) < (667)):
                            coll29_[default__.safeModuloInt(d_473_v141_, d_423_v96_.f29)] = d_301_v2_
                    return _dafny.Map(coll29_)
                nw79_[int(14)] = iife33_()
                
                nw79_[int(15)] = (d_471_v140_)[default__.safeIndex(d_303_v4_, len(d_471_v140_))]
                nw79_[int(16)] = _dafny.Map({d_467_v136_.f30: (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]})
                nw79_[int(17)] = _dafny.Map({(d_439_v110_).fm6(_dafny.Map({d_423_v96_.f30: (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]}), d_423_v96_.f30, d_311_v9_, _dafny.Map({d_438_v109_: (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]}), d_309_globalState_): d_301_v2_})
                nw79_[int(18)] = d_444_v115_
                d_472_v142_ = nw79_
                index56_ = default__.safeIndex(533, (d_472_v142_).length(0))
                (d_472_v142_)[index56_] = (d_444_v115_ if d_301_v2_ else d_444_v115_)
                index57_ = default__.safeIndex(533, (d_472_v142_).length(0))
                index58_ = default__.safeIndex(805, (d_298_v0_).length(0))
                rhs56_ = default__.fm18(d_423_v96_.f29, d_309_globalState_)
                rhs57_ = ((d_423_v96_.f29) + ((d_439_v110_).f26)) > (d_426_v98_)
                rhs58_ = d_301_v2_
                rhs59_ = (default__.fm36(not((d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]), (d_433_v105_).fm8(d_437_v108_, ((d_444_v115_)[(d_437_v108_).cardinality] if ((d_437_v108_).cardinality) in (d_444_v115_) else not(d_301_v2_)), 447, (0) - (d_423_v96_.f30), d_309_globalState_), not(not((d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))])), d_309_globalState_)).set(default__.safeIndex(d_423_v96_.f30, len(default__.fm36(not((d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]), (d_433_v105_).fm8(d_437_v108_, ((d_444_v115_)[(d_437_v108_).cardinality] if ((d_437_v108_).cardinality) in (d_444_v115_) else not(d_301_v2_)), 447, (0) - (d_423_v96_.f30), d_309_globalState_), not(not((d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))])), d_309_globalState_))), d_434_v106_)
                lhs55_ = d_472_v142_
                lhs56_ = default__.safeIndex(533, (d_472_v142_).length(0))
                lhs57_ = d_309_globalState_
                lhs58_ = d_298_v0_
                lhs59_ = default__.safeIndex(805, (d_298_v0_).length(0))
                lhs60_ = d_309_globalState_
                lhs55_[lhs56_] = rhs56_
                lhs57_.f9 = rhs57_
                lhs58_[lhs59_] = rhs58_
                lhs60_.f23 = rhs59_
                d_474_v143_: _dafny.Array
                nw80_ = _dafny.Array(False, 5)
                d_474_v143_ = nw80_
                index59_ = default__.safeIndex(318, (d_310_v8_).length(0))
                (d_310_v8_)[index59_] = d_474_v143_
                index60_ = default__.safeIndex(318, (d_310_v8_).length(0))
                (d_310_v8_)[index60_] = d_474_v143_
            d_475_v144_: _dafny.Map
            d_475_v144_ = _dafny.Map({(d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]: (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]})
            d_476_v145_: _dafny.Seq
            d_476_v145_ = _dafny.SeqWithoutIsStrInference([(d_439_v110_).f27])
            d_477_v146_: _dafny.Seq
            d_477_v146_ = _dafny.SeqWithoutIsStrInference([d_302_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrj")), default__.fm36((d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))], default__.fm5((d_439_v110_).f26, ((d_475_v144_)[(d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]] if ((d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]) in (d_475_v144_) else (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))]), d_301_v2_, d_309_globalState_), not(True), d_309_globalState_), (d_476_v145_)[default__.safeIndex((d_439_v110_).f26, len(d_476_v145_))]])
            d_478_v147_: _dafny.Map
            d_478_v147_ = _dafny.Map({(0) - ((d_426_v98_ if (d_298_v0_)[default__.safeIndex(805, (d_298_v0_).length(0))] else (d_311_v9_)[default__.safeIndex(d_423_v96_.f30, len(d_311_v9_))])): d_477_v146_})
            d_478_v147_ = (d_478_v147_).set((0) - (len(d_304_v5_)), d_477_v146_)
        hi2_ = 782
        for d_479_i6_ in range((0) - (d_426_v98_), hi2_):
            if (d_302_v3_) < (d_302_v3_):
                (d_309_globalState_).f5 = (0) - (d_303_v4_)
                d_480_v148_: _dafny.Map
                d_480_v148_ = _dafny.Map({d_301_v2_: d_301_v2_})
                d_303_v4_ = (default__.safeDivisionInt(len(d_480_v148_), d_423_v96_.f29)) * (d_423_v96_.f29)
                d_481_v149_: _dafny.Seq
                d_481_v149_ = _dafny.SeqWithoutIsStrInference([d_301_v2_])
                (d_309_globalState_).f24 = len(((d_481_v149_) + (_dafny.SeqWithoutIsStrInference([d_301_v2_]))) + (d_481_v149_))
                d_482_v150_: _dafny.Array
                nw81_ = _dafny.Array(_dafny.CodePoint('D'), 27)
                d_482_v150_ = nw81_
                d_483_v151_: D20
                d_483_v151_ = D20_DC54(d_482_v150_)
                d_483_v151_ = d_483_v151_
                index61_ = default__.safeIndex(782, (d_298_v0_).length(0))
                (d_298_v0_)[index61_] = d_301_v2_
                d_484_v152_: D10
                d_484_v152_ = D10_DC31(d_301_v2_, d_311_v9_, d_301_v2_, 843)
                index62_ = default__.safeIndex(782, (d_298_v0_).length(0))
                rhs60_ = (d_311_v9_) <= ((d_484_v152_).cf59)
                rhs61_ = d_301_v2_
                rhs62_ = default__.safeModuloInt(d_423_v96_.f30, d_423_v96_.f30)
                rhs63_ = ((d_301_v2_) or (d_301_v2_)) == (d_301_v2_)
                lhs61_ = d_309_globalState_
                lhs62_ = d_309_globalState_
                lhs63_ = d_309_globalState_
                lhs64_ = d_298_v0_
                lhs65_ = default__.safeIndex(782, (d_298_v0_).length(0))
                lhs61_.f25 = rhs60_
                lhs62_.f25 = rhs61_
                lhs63_.f8 = rhs62_
                lhs64_[lhs65_] = rhs63_
            elif True:
                (d_309_globalState_).f9 = d_301_v2_
                d_485_v153_: _dafny.Map
                d_485_v153_ = _dafny.Map({len(d_429_v101_): d_301_v2_})
                d_486_v154_: _dafny.Map
                d_486_v154_ = _dafny.Map({default__.fm2(d_301_v2_, d_309_globalState_): default__.fm40(d_301_v2_, 339, d_485_v153_, d_301_v2_, d_309_globalState_)})
                d_487_v155_: C2
                nw82_ = C2()
                nw82_.ctor__(_dafny.SeqWithoutIsStrInference([898 for d_488_i7_ in range(default__.abs(621))]), (d_304_v5_ if True else _dafny.Map({d_423_v96_.f30: (0) - (d_426_v98_)})))
                d_487_v155_ = nw82_
                d_489_v156_: _dafny.Map
                d_489_v156_ = _dafny.Map({d_301_v2_: _dafny.Set({d_434_v106_})})
                d_490_v157_: _dafny.Set
                d_490_v157_ = _dafny.Set({d_434_v106_})
                def iife34_():
                    coll30_ = _dafny.Set()
                    compr_30_: str
                    for compr_30_ in (d_490_v157_).Elements:
                        d_491_v158_: str = compr_30_
                        if (d_491_v158_) in (d_490_v157_):
                            coll30_ = coll30_.union(_dafny.Set([d_491_v158_]))
                    return _dafny.Set(coll30_)
                rhs64_ = _dafny.Map({d_302_v3_: (default__.fm40(d_301_v2_, d_423_v96_.f30, d_485_v153_, d_301_v2_, d_309_globalState_)) | (d_304_v5_)})
                rhs65_ = (len(((d_489_v156_)[d_301_v2_] if (d_301_v2_) in (d_489_v156_) else iife34_()
                ))) - (d_479_i6_)
                rhs66_ = (0) - (d_303_v4_)
                rhs67_ = d_487_v155_
                lhs66_ = d_309_globalState_
                lhs67_ = d_309_globalState_
                d_486_v154_ = rhs64_
                lhs66_.f17 = rhs65_
                lhs67_.f17 = rhs66_
                d_487_v155_ = rhs67_
                d_492_v159_: C5
                nw83_ = C5()
                nw83_.ctor__((0) - (default__.safeDivisionInt(d_303_v4_, d_423_v96_.f29)), d_426_v98_, d_479_i6_, _dafny.SeqWithoutIsStrInference([d_434_v106_ for d_493_i8_ in range(default__.abs(511))]))
                d_492_v159_ = nw83_
                (d_309_globalState_).f25 = not(d_301_v2_)
                d_494_v160_: _dafny.Array
                nw84_ = _dafny.Array(None, 19)
                nw84_[int(0)] = d_302_v3_
                nw84_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwpepkb"))
                nw84_[int(2)] = d_302_v3_
                nw84_[int(3)] = _dafny.SeqWithoutIsStrInference([(D24_DC72(d_301_v2_, d_423_v96_.f29, d_434_v106_, (_dafny.SeqWithoutIsStrInference([d_434_v106_ for d_495_i10_ in range(default__.abs(551))])).set(default__.safeIndex(d_423_v96_.f29, len(_dafny.SeqWithoutIsStrInference([d_434_v106_ for d_496_i10_ in range(default__.abs(551))]))), d_434_v106_), d_301_v2_)).cf138 for d_497_i9_ in range(default__.abs(-754))])
                nw84_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krjjo"))
                nw84_[int(5)] = d_302_v3_
                nw84_[int(6)] = d_302_v3_
                nw84_[int(7)] = d_302_v3_
                nw84_[int(8)] = d_302_v3_
                nw84_[int(9)] = d_302_v3_
                nw84_[int(10)] = d_302_v3_
                nw84_[int(11)] = d_302_v3_
                nw84_[int(12)] = d_302_v3_
                nw84_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "of"))
                nw84_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbfnkt"))
                nw84_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbavnfdvc"))
                nw84_[int(16)] = d_302_v3_
                nw84_[int(17)] = d_302_v3_
                nw84_[int(18)] = default__.fm36(d_301_v2_, d_301_v2_, d_301_v2_, d_309_globalState_)
                d_494_v160_ = nw84_
                d_498_v161_: _dafny.Map
                d_498_v161_ = _dafny.Map({d_494_v160_: not(d_301_v2_)})
                d_498_v161_ = (d_498_v161_).set(d_494_v160_, d_301_v2_)
            d_499_v162_: C3
            nw85_ = C3()
            nw85_.ctor__(d_423_v96_.f30, d_303_v4_, d_302_v3_)
            d_499_v162_ = nw85_
            d_500_v163_: _dafny.Seq
            d_500_v163_ = _dafny.SeqWithoutIsStrInference([d_499_v162_, d_499_v162_, d_499_v162_, d_499_v162_])
            d_501_v164_: _dafny.MultiSet
            d_501_v164_ = _dafny.MultiSet([len(d_500_v163_)])
            d_502_v165_: _dafny.Map
            d_502_v165_ = _dafny.Map({d_301_v2_: d_434_v106_})
            d_503_v166_: _dafny.MultiSet
            d_503_v166_ = _dafny.MultiSet([d_426_v98_, (d_501_v164_).cardinality, (d_499_v162_).f34, len(d_502_v165_)])
            (d_309_globalState_).f9 = (default__.fm5(d_423_v96_.f29, d_301_v2_, (d_423_v96_).fm8(d_503_v166_, d_301_v2_, (d_499_v162_).f34, d_426_v98_, d_309_globalState_), d_309_globalState_)) or (not((259) < (d_423_v96_.f29)))
            d_504_v167_: _dafny.Array
            def lambda15_(d_505_v106_):
                def lambda16_(d_506_i11_):
                    return d_505_v106_

                return lambda16_

            init8_ = lambda15_(d_434_v106_)
            nw86_ = _dafny.Array(None, 26)
            for i0_8_ in range(nw86_.length(0)):
                nw86_[i0_8_] = init8_(i0_8_)
            d_504_v167_ = nw86_
            d_507_v168_: D20
            d_507_v168_ = D20_DC54(d_504_v167_)
            d_508_v169_: _dafny.Map
            d_508_v169_ = _dafny.Map({d_507_v168_: d_301_v2_})
            d_509_v170_: _dafny.Map
            d_509_v170_ = _dafny.Map({d_479_i6_: d_508_v169_})
            d_510_v171_: T1
            nw87_ = C4()
            nw87_.ctor__((D11_DC34(d_298_v0_, d_301_v2_)).cf67, d_479_i6_, d_302_v3_, d_304_v5_)
            d_510_v171_ = nw87_
            d_511_v172_: _dafny.Map
            d_511_v172_ = _dafny.Map({d_510_v171_: d_508_v169_})
            d_512_v173_: _dafny.Array
            nw88_ = _dafny.Array(None, 21)
            nw88_[int(0)] = d_508_v169_
            nw88_[int(1)] = (d_508_v169_).set(d_507_v168_, d_301_v2_)
            nw88_[int(2)] = d_508_v169_
            nw88_[int(3)] = d_508_v169_
            nw88_[int(4)] = ((d_509_v170_)[d_423_v96_.f29] if (d_423_v96_.f29) in (d_509_v170_) else d_508_v169_)
            nw88_[int(5)] = d_508_v169_
            nw88_[int(6)] = (d_508_v169_) | (d_508_v169_)
            nw88_[int(7)] = d_508_v169_
            nw88_[int(8)] = d_508_v169_
            nw88_[int(9)] = (_dafny.Map({d_507_v168_: d_301_v2_})) | (d_508_v169_)
            nw88_[int(10)] = d_508_v169_
            nw88_[int(11)] = d_508_v169_
            nw88_[int(12)] = (d_508_v169_ if d_301_v2_ else d_508_v169_)
            nw88_[int(13)] = (d_508_v169_) | (d_508_v169_)
            nw88_[int(14)] = (d_508_v169_).set(D20_DC54(d_504_v167_), d_301_v2_)
            nw88_[int(15)] = ((d_509_v170_)[(d_311_v9_)[default__.safeIndex(d_479_i6_, len(d_311_v9_))]] if ((d_311_v9_)[default__.safeIndex(d_479_i6_, len(d_311_v9_))]) in (d_509_v170_) else _dafny.Map({D20_DC54(d_504_v167_): True}))
            nw88_[int(16)] = ((d_511_v172_)[d_510_v171_] if (d_510_v171_) in (d_511_v172_) else _dafny.Map({d_507_v168_: d_301_v2_}))
            nw88_[int(17)] = d_508_v169_
            nw88_[int(18)] = d_508_v169_
            nw88_[int(19)] = d_508_v169_
            nw88_[int(20)] = d_508_v169_
            d_512_v173_ = nw88_
            index63_ = default__.safeIndex(384, (d_512_v173_).length(0))
            (d_512_v173_)[index63_] = (d_508_v169_) | (_dafny.Map({d_507_v168_: d_301_v2_}))
            index64_ = default__.safeIndex(384, (d_512_v173_).length(0))
            (d_512_v173_)[index64_] = (d_508_v169_).set(d_507_v168_, not(d_301_v2_))
            d_513_v174_: _dafny.MultiSet
            d_513_v174_ = _dafny.MultiSet([False])
            (d_423_v96_).f29 = (d_423_v96_.f30 if (d_433_v105_).fm8(d_503_v166_, not(d_301_v2_), (0) - ((d_513_v174_).cardinality), d_423_v96_.f29, d_309_globalState_) else (0) - ((len((d_510_v171_).f27)) * (len(_dafny.SeqWithoutIsStrInference([d_423_v96_.f30, d_423_v96_.f29])))))
        d_514_v175_: _dafny.Seq
        d_514_v175_ = _dafny.SeqWithoutIsStrInference([d_301_v2_, True])
        d_515_v176_: _dafny.Map
        d_515_v176_ = _dafny.Map({d_301_v2_: d_298_v0_})
        if default__.fm5((len(d_514_v175_)) + (len(d_515_v176_)), (d_303_v4_) <= (d_303_v4_), not (d_301_v2_) or (d_301_v2_), d_309_globalState_):
            d_516_v177_: D7
            d_516_v177_ = D7_DC17(d_298_v0_)
            d_517_v178_: C3
            nw89_ = C3()
            nw89_.ctor__(d_423_v96_.f30, d_423_v96_.f29, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvrlyal")))
            d_517_v178_ = nw89_
            d_518_v179_: _dafny.Map
            d_518_v179_ = _dafny.Map({(d_516_v177_).cf33: d_517_v178_})
            index65_ = default__.safeIndex(974, (d_306_v7_).length(0))
            (d_306_v7_)[index65_] = len(d_518_v179_)
            index66_ = default__.safeIndex(974, (d_306_v7_).length(0))
            (d_306_v7_)[index66_] = (d_517_v178_).f34
            d_519_v180_: D16
            d_519_v180_ = D16_DC45(d_423_v96_.f30, d_301_v2_, ((d_517_v178_).f34) - (162), (805) + ((d_306_v7_)[default__.safeIndex(974, (d_306_v7_).length(0))]), (not(d_301_v2_) if d_301_v2_ else d_301_v2_))
            d_520_v181_: _dafny.Set
            d_520_v181_ = _dafny.Set({(d_306_v7_)[default__.safeIndex(974, (d_306_v7_).length(0))]})
            d_519_v180_ = (D16_DC45(d_303_v4_, d_301_v2_, (d_517_v178_).f34, d_423_v96_.f30, d_301_v2_) if not((d_520_v181_).isdisjoint(d_520_v181_)) else d_519_v180_)
            d_521_v182_: C1
            nw90_ = C1()
            nw90_.ctor__()
            d_521_v182_ = nw90_
            d_522_v183_: C2
            nw91_ = C2()
            nw91_.ctor__(d_311_v9_, d_304_v5_)
            d_522_v183_ = nw91_
            (d_309_globalState_).f25 = False
        elif True:
            index67_ = default__.safeIndex(831, (d_306_v7_).length(0))
            (d_306_v7_)[index67_] = d_423_v96_.f30
            d_523_v184_: _dafny.Array
            nw92_ = _dafny.Array(None, 7)
            nw92_[int(0)] = default__.fm42(_dafny.Set({(0) - (d_423_v96_.f29)}), d_309_globalState_)
            nw92_[int(1)] = d_434_v106_
            nw92_[int(2)] = (d_434_v106_ if not(True) else d_434_v106_)
            nw92_[int(3)] = d_434_v106_
            nw92_[int(4)] = d_434_v106_
            nw92_[int(5)] = d_434_v106_
            nw92_[int(6)] = (d_302_v3_)[default__.safeIndex(len(d_311_v9_), len(d_302_v3_))]
            d_523_v184_ = nw92_
            index68_ = default__.safeIndex(831, (d_306_v7_).length(0))
            rhs68_ = d_423_v96_.f29
            rhs69_ = (d_423_v96_.f29) * ((0) - (d_423_v96_.f30))
            rhs70_ = d_523_v184_
            rhs71_ = d_311_v9_
            lhs68_ = d_423_v96_
            lhs69_ = d_306_v7_
            lhs70_ = default__.safeIndex(831, (d_306_v7_).length(0))
            lhs68_.f29 = rhs68_
            lhs69_[lhs70_] = rhs69_
            d_523_v184_ = rhs70_
            d_311_v9_ = rhs71_
            d_524_v185_: _dafny.Array
            nw93_ = _dafny.Array(D10.default()(), 27)
            d_524_v185_ = nw93_
            d_524_v185_ = d_524_v185_
            d_525_v186_: _dafny.Set
            d_525_v186_ = _dafny.Set({915})
            d_526_v187_: _dafny.Array
            nw94_ = _dafny.Array(None, 13)
            nw94_[int(0)] = d_423_v96_.f29
            nw94_[int(1)] = (d_306_v7_)[default__.safeIndex(831, (d_306_v7_).length(0))]
            nw94_[int(2)] = len(d_525_v186_)
            nw94_[int(3)] = (0) - ((d_306_v7_)[default__.safeIndex(831, (d_306_v7_).length(0))])
            nw94_[int(4)] = d_423_v96_.f30
            nw94_[int(5)] = (d_306_v7_)[default__.safeIndex(831, (d_306_v7_).length(0))]
            nw94_[int(6)] = d_426_v98_
            nw94_[int(7)] = d_423_v96_.f29
            nw94_[int(8)] = d_303_v4_
            nw94_[int(9)] = len(((d_302_v3_).set(default__.safeIndex(d_426_v98_, len(d_302_v3_)), d_434_v106_)).set(default__.safeIndex(d_303_v4_, len((d_302_v3_).set(default__.safeIndex(d_426_v98_, len(d_302_v3_)), d_434_v106_))), d_434_v106_))
            nw94_[int(10)] = ((d_431_v103_)[False] if (False) in (d_431_v103_) else -435)
            nw94_[int(11)] = d_423_v96_.f30
            nw94_[int(12)] = d_303_v4_
            d_526_v187_ = nw94_
            d_527_v188_: int
            d_528_v189_: bool
            out14_: int
            out15_: bool
            out14_, out15_ = (d_423_v96_).m1(d_526_v187_, (d_423_v96_.f29) == (len(d_525_v186_)), d_309_globalState_)
            d_527_v188_ = out14_
            d_528_v189_ = out15_
            index69_ = default__.safeIndex(779, (d_298_v0_).length(0))
            (d_298_v0_)[index69_] = True
            index70_ = default__.safeIndex(779, (d_298_v0_).length(0))
            (d_298_v0_)[index70_] = not ((d_301_v2_) or (d_301_v2_)) or (d_301_v2_)
            d_529_v190_: T0
            nw95_ = C6()
            nw95_.ctor__((d_433_v105_).f28, d_423_v96_.f29, (d_302_v3_).set(default__.safeIndex(d_426_v98_, len(d_302_v3_)), d_434_v106_))
            d_529_v190_ = nw95_
            d_530_v191_: _dafny.Seq
            d_530_v191_ = _dafny.SeqWithoutIsStrInference([d_529_v190_])
            d_531_v192_: _dafny.MultiSet
            d_531_v192_ = _dafny.MultiSet([len(d_530_v191_)])
            d_301_v2_ = (d_423_v96_).fm8((d_531_v192_) | (d_531_v192_), d_528_v189_, len(d_304_v5_), d_527_v188_, d_309_globalState_)
        d_532_v193_: D22
        d_532_v193_ = D22_DC62(116)
        d_533_v194_: _dafny.Set
        d_533_v194_ = _dafny.Set({d_423_v96_.f30})
        pat_let_tv1_ = d_423_v96_
        d_534_v195_: _dafny.Array
        nw96_ = _dafny.Array(None, 21)
        nw96_[int(0)] = d_532_v193_
        nw96_[int(1)] = d_532_v193_
        nw96_[int(2)] = D22_DC62(d_426_v98_)
        nw96_[int(3)] = d_532_v193_
        def iife35_(_pat_let2_0):
            def iife36_(d_535_dt__update__tmp_h1_):
                def iife37_(_pat_let3_0):
                    def iife38_(d_536_dt__update_hcf119_h0_):
                        return D22_DC62(d_536_dt__update_hcf119_h0_)
                    return iife38_(_pat_let3_0)
                return iife37_(pat_let_tv1_.f29)
            return iife36_(_pat_let2_0)
        nw96_[int(4)] = iife35_(d_532_v193_)
        nw96_[int(5)] = d_532_v193_
        nw96_[int(6)] = d_532_v193_
        nw96_[int(7)] = d_532_v193_
        nw96_[int(8)] = d_532_v193_
        nw96_[int(9)] = d_532_v193_
        nw96_[int(10)] = d_532_v193_
        nw96_[int(11)] = d_532_v193_
        nw96_[int(12)] = default__.fm58(d_302_v3_, d_311_v9_, _dafny.Map({len(d_533_v194_): d_423_v96_.f30}), d_423_v96_.f29, d_309_globalState_)
        nw96_[int(13)] = d_532_v193_
        nw96_[int(14)] = d_532_v193_
        nw96_[int(15)] = d_532_v193_
        nw96_[int(16)] = d_532_v193_
        nw96_[int(17)] = d_532_v193_
        nw96_[int(18)] = d_532_v193_
        nw96_[int(19)] = d_532_v193_
        nw96_[int(20)] = D22_DC62(d_423_v96_.f29)
        d_534_v195_ = nw96_
        d_537_v196_: D23
        d_537_v196_ = D23_DC68(d_301_v2_, d_301_v2_)
        d_538_v197_: _dafny.Set
        d_538_v197_ = _dafny.Set({d_537_v196_})
        d_539_v198_: _dafny.Map
        d_539_v198_ = _dafny.Map({d_534_v195_: d_538_v197_})
        (d_309_globalState_).f25 = (d_539_v198_) == (_dafny.Map({d_534_v195_: _dafny.Set({default__.fm59(d_309_globalState_), d_537_v196_})}))
        d_540_v199_: _dafny.Map
        d_540_v199_ = _dafny.Map({False: d_301_v2_})
        d_541_v200_: _dafny.Map
        d_541_v200_ = _dafny.Map({default__.fm0(d_301_v2_, d_302_v3_, default__.fm17(d_309_globalState_), d_423_v96_.f29, d_309_globalState_): ((d_540_v199_)[True] if (True) in (d_540_v199_) else d_301_v2_)})
        d_542_v201_: C1
        nw97_ = C1()
        nw97_.ctor__()
        d_542_v201_ = nw97_
        d_543_v202_: _dafny.MultiSet
        d_543_v202_ = _dafny.MultiSet([d_542_v201_])
        d_544_v203_: _dafny.Map
        d_544_v203_ = _dafny.Map({default__.safeDivisionInt(len((d_541_v200_).set(d_514_v175_, d_301_v2_)), (0) - (d_423_v96_.f29)): d_543_v202_})
        d_545_v205_: _dafny.Map
        d_545_v205_ = _dafny.Map({d_423_v96_.f30: d_434_v106_})
        d_546_v207_: _dafny.MultiSet
        def iife39_():
            coll31_ = _dafny.Map()
            compr_31_: int
            for compr_31_ in (d_533_v194_).Elements:
                d_547_v204_: int = compr_31_
                if (d_547_v204_) in (d_533_v194_):
                    coll31_[default__.safeDivisionInt(d_547_v204_, d_426_v98_)] = d_434_v106_
            return _dafny.Map(coll31_)
        def iife40_():
            coll32_ = _dafny.Map()
            compr_32_: int
            for compr_32_ in (d_533_v194_).Elements:
                d_548_v206_: int = compr_32_
                if (d_548_v206_) in (d_533_v194_):
                    coll32_[default__.safeDivisionInt(d_548_v206_, 296)] = d_434_v106_
            return _dafny.Map(coll32_)
        d_546_v207_ = _dafny.MultiSet([iife39_()
        , d_545_v205_, d_545_v205_, iife40_()
        ])
        d_544_v203_ = (d_544_v203_).set(((d_546_v207_) - (d_546_v207_)).cardinality, d_543_v202_)
        d_549_v208_: D8
        d_549_v208_ = D8_DC24(d_303_v4_, default__.safeDivisionInt((0) - (d_423_v96_.f30), d_426_v98_), d_423_v96_.f29)
        source12_ = d_549_v208_
        if source12_.is_DC23:
            d_550___mcc_h11_ = source12_.cf42
            d_551___mcc_h12_ = source12_.cf43
            d_552___mcc_h13_ = source12_.cf44
            d_553___mcc_h14_ = source12_.cf45
            d_554_cf45_ = d_553___mcc_h14_
            d_555_cf44_ = d_552___mcc_h13_
            d_556_cf43_ = d_551___mcc_h12_
            d_557_cf42_ = d_550___mcc_h11_
            d_558_v209_: _dafny.Array
            def lambda17_(d_559_i12_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ne"))

            init9_ = lambda17_
            nw98_ = _dafny.Array(None, 8)
            for i0_9_ in range(nw98_.length(0)):
                nw98_[i0_9_] = init9_(i0_9_)
            d_558_v209_ = nw98_
            rhs72_ = d_554_cf45_
            rhs73_ = d_558_v209_
            lhs71_ = d_309_globalState_
            lhs71_.f25 = rhs72_
            d_558_v209_ = rhs73_
            d_560_v210_: T0
            nw99_ = C6()
            nw99_.ctor__(((d_433_v105_).f28) | ((d_433_v105_).f28), d_426_v98_, default__.fm43(d_554_cf45_, d_303_v4_, d_309_globalState_))
            d_560_v210_ = nw99_
            d_560_v210_ = d_560_v210_
            d_561_v211_: _dafny.MultiSet
            d_561_v211_ = _dafny.MultiSet([d_423_v96_.f30])
            d_562_v212_: _dafny.MultiSet
            d_562_v212_ = _dafny.MultiSet([d_561_v211_, d_561_v211_, _dafny.MultiSet(d_311_v9_), (d_561_v211_ if d_301_v2_ else d_561_v211_), d_561_v211_])
            rhs74_ = 292
            rhs75_ = d_562_v212_
            lhs72_ = d_423_v96_
            lhs72_.f29 = rhs74_
            d_562_v212_ = rhs75_
            (d_309_globalState_).f24 = default__.safeDivisionInt(default__.safeModuloInt(d_423_v96_.f30, d_423_v96_.f30), d_423_v96_.f29)
        elif source12_.is_DC24:
            d_563___mcc_h15_ = source12_.cf46
            d_564___mcc_h16_ = source12_.cf47
            d_565___mcc_h17_ = source12_.cf48
            d_566_cf48_ = d_565___mcc_h17_
            d_567_cf47_ = d_564___mcc_h16_
            d_568_cf46_ = d_563___mcc_h15_
            (d_309_globalState_).f23 = d_302_v3_
            d_569_v213_: _dafny.Seq
            d_569_v213_ = _dafny.SeqWithoutIsStrInference([d_302_v3_])
            d_570_v214_: _dafny.Map
            d_570_v214_ = _dafny.Map({d_423_v96_.f30: d_301_v2_})
            d_571_v215_: _dafny.Seq
            d_571_v215_ = _dafny.SeqWithoutIsStrInference([default__.fm60(d_301_v2_, d_302_v3_, d_301_v2_, d_302_v3_, d_309_globalState_), default__.fm60(d_301_v2_, d_302_v3_, d_301_v2_, (d_569_v213_)[default__.safeIndex(len(d_570_v214_), len(d_569_v213_))], d_309_globalState_)])
            d_571_v215_ = (d_571_v215_) + (d_571_v215_)
            (d_309_globalState_).f24 = d_423_v96_.f30
            (d_309_globalState_).f1 = d_298_v0_
        elif source12_.is_DC25:
            d_572___mcc_h18_ = source12_.cf49
            d_573___mcc_h19_ = source12_.cf50
            d_574___mcc_h20_ = source12_.cf51
            d_575___mcc_h21_ = source12_.cf52
            d_576___mcc_h22_ = source12_.cf53
            d_577_cf53_ = d_576___mcc_h22_
            d_578_cf52_ = d_575___mcc_h21_
            d_579_cf51_ = d_574___mcc_h20_
            d_580_cf50_ = d_573___mcc_h19_
            d_581_cf49_ = d_572___mcc_h18_
            d_582_v216_: _dafny.Map
            d_582_v216_ = _dafny.Map({_dafny.MultiSet([d_577_cf53_]): d_514_v175_})
            d_583_v217_: _dafny.MultiSet
            d_583_v217_ = _dafny.MultiSet([d_423_v96_.f30])
            d_584_v218_: _dafny.Map
            d_584_v218_ = _dafny.Map({d_581_cf49_: d_583_v217_})
            d_585_v219_: D0
            d_585_v219_ = D0_DC3(d_577_cf53_, d_577_cf53_)
            d_586_v220_: D12
            d_586_v220_ = D12_DC37(d_423_v96_.f29, d_423_v96_.f29, default__.fm19(default__.fm32(d_581_cf49_, d_582_v216_, d_584_v218_, d_434_v106_, d_309_globalState_), d_585_v219_, d_579_cf51_, d_581_cf49_, d_309_globalState_), len(default__.fm0(d_581_cf49_, d_302_v3_, d_312_v10_, d_426_v98_, d_309_globalState_)), d_423_v96_.f30)
            pat_let_tv2_ = d_577_cf53_
            def iife41_(_pat_let4_0):
                def iife42_(d_587_dt__update__tmp_h2_):
                    def iife43_(_pat_let5_0):
                        def iife44_(d_588_dt__update_hcf74_h0_):
                            return D12_DC37((d_587_dt__update__tmp_h2_).cf70, (d_587_dt__update__tmp_h2_).cf71, (d_587_dt__update__tmp_h2_).cf72, (d_587_dt__update__tmp_h2_).cf73, d_588_dt__update_hcf74_h0_)
                        return iife44_(_pat_let5_0)
                    return iife43_(pat_let_tv2_)
                return iife42_(_pat_let4_0)
            source13_ = iife41_(d_586_v220_)
            if source13_.is_DC37:
                d_589___mcc_h25_ = source13_.cf70
                d_590___mcc_h26_ = source13_.cf71
                d_591___mcc_h27_ = source13_.cf72
                d_592___mcc_h28_ = source13_.cf73
                d_593___mcc_h29_ = source13_.cf74
                d_594_cf74_ = d_593___mcc_h29_
                d_595_cf73_ = d_592___mcc_h28_
                d_596_cf72_ = d_591___mcc_h27_
                d_597_cf71_ = d_590___mcc_h26_
                d_598_cf70_ = d_589___mcc_h25_
                d_599_v221_: _dafny.MultiSet
                d_599_v221_ = _dafny.MultiSet([d_581_cf49_, (d_433_v105_).fm8(d_583_v217_, False, d_597_cf71_, d_423_v96_.f29, d_309_globalState_), d_581_cf49_])
                d_600_v222_: _dafny.Seq
                d_600_v222_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_434_v106_ for d_601_i13_ in range(default__.abs(748))]), d_302_v3_])
                index71_ = default__.safeIndex(234, (d_306_v7_).length(0))
                (d_306_v7_)[index71_] = ((d_599_v221_) - (default__.fm37(d_581_cf49_, True, d_426_v98_, len(d_600_v222_), d_309_globalState_))).cardinality
                index72_ = default__.safeIndex(234, (d_306_v7_).length(0))
                (d_306_v7_)[index72_] = 377
                d_602_v223_: _dafny.Array
                nw100_ = _dafny.Array(None, 28)
                d_602_v223_ = nw100_
                d_603_v224_: _dafny.Map
                d_603_v224_ = _dafny.Map({d_434_v106_: d_602_v223_})
                d_604_v225_: D28
                d_604_v225_ = D28_DC86(((d_603_v224_)[d_434_v106_] if (d_434_v106_) in (d_603_v224_) else d_602_v223_))
                d_602_v223_ = (d_604_v225_).cf164
                d_580_cf50_ = (d_580_cf50_).set((d_423_v96_).fm8(d_583_v217_, d_581_cf49_, d_598_cf70_, d_423_v96_.f30, d_309_globalState_), d_423_v96_.f30)
                (d_309_globalState_).f5 = default__.safeModuloInt(d_423_v96_.f29, (d_306_v7_)[default__.safeIndex(234, (d_306_v7_).length(0))])
            elif source13_.is_DC38:
                nw101_ = _dafny.Array(int(0), 24)
                d_306_v7_ = nw101_
                (d_309_globalState_).f5 = d_577_cf53_
                (d_423_v96_).f30 = d_303_v4_
                d_605_v226_: C5
                nw102_ = C5()
                nw102_.ctor__(len((_dafny.SeqWithoutIsStrInference([d_579_cf51_ for d_606_i14_ in range(default__.abs(547))])) + (_dafny.SeqWithoutIsStrInference([d_434_v106_ for d_607_i15_ in range(default__.abs(284))]))), (d_423_v96_.f30) * (d_426_v98_), d_423_v96_.f30, d_302_v3_)
                d_605_v226_ = nw102_
            elif True:
                d_608___mcc_h30_ = source13_.cf69
                d_609_cf69_ = d_608___mcc_h30_
                d_610_v227_: int
                d_611_v228_: bool
                out16_: int
                out17_: bool
                out16_, out17_ = (d_542_v201_).m1(d_306_v7_, (d_301_v2_ if d_301_v2_ else d_581_cf49_), d_309_globalState_)
                d_610_v227_ = out16_
                d_611_v228_ = out17_
                d_612_v229_: _dafny.Map
                d_612_v229_ = _dafny.Map({d_610_v227_: d_581_cf49_})
                d_613_v230_: _dafny.Map
                d_613_v230_ = _dafny.Map({d_301_v2_: d_311_v9_})
                (d_423_v96_).f30 = (len(_dafny.Set({d_303_v4_, (d_433_v105_).fm6((d_612_v229_).set(len(((d_613_v230_)[d_581_cf49_] if (d_581_cf49_) in (d_613_v230_) else _dafny.SeqWithoutIsStrInference([d_423_v96_.f29]))), False), d_610_v227_, d_311_v9_, d_541_v200_, d_309_globalState_), d_423_v96_.f29})) if d_611_v228_ else 626)
                rhs76_ = True
                rhs77_ = d_301_v2_
                lhs73_ = d_309_globalState_
                lhs74_ = d_309_globalState_
                lhs73_.f9 = rhs76_
                lhs74_.f25 = rhs77_
                (d_309_globalState_).f25 = d_611_v228_
            d_614_v231_: _dafny.Array
            def lambda18_(d_615_cf51_):
                def lambda19_(d_616_i16_):
                    return d_615_cf51_

                return lambda19_

            init10_ = lambda18_(d_579_cf51_)
            nw103_ = _dafny.Array(None, 29)
            for i0_10_ in range(nw103_.length(0)):
                nw103_[i0_10_] = init10_(i0_10_)
            d_614_v231_ = nw103_
            index73_ = default__.safeIndex(132, (d_614_v231_).length(0))
            (d_614_v231_)[index73_] = d_434_v106_
            index74_ = default__.safeIndex(132, (d_614_v231_).length(0))
            (d_614_v231_)[index74_] = default__.fm20(d_301_v2_, (default__.fm61(d_581_cf49_, d_301_v2_, d_577_cf53_, d_309_globalState_)).cf108, ((d_583_v217_)[d_577_cf53_] if (d_577_cf53_) in (d_583_v217_) else d_423_v96_.f29), _dafny.CodePoint('p'), d_309_globalState_)
            d_617_v232_: _dafny.Array
            def lambda20_(d_618_cf49_, d_619_cf50_, d_620_cf51_, d_621_cf52_, d_622_v96_):
                def lambda21_(d_623_i17_):
                    return D8_DC25(d_618_cf49_, d_619_cf50_, d_620_cf51_, d_621_cf52_, d_622_v96_.f29)

                return lambda21_

            init11_ = lambda20_(d_581_cf49_, d_580_cf50_, d_579_cf51_, d_578_cf52_, d_423_v96_)
            nw104_ = _dafny.Array(None, 16)
            for i0_11_ in range(nw104_.length(0)):
                nw104_[i0_11_] = init11_(i0_11_)
            d_617_v232_ = nw104_
            rhs78_ = default__.fm36(not (d_301_v2_) or (d_581_cf49_), d_581_cf49_, d_581_cf49_, d_309_globalState_)
            rhs79_ = d_617_v232_
            rhs80_ = d_303_v4_
            lhs75_ = d_309_globalState_
            lhs76_ = d_309_globalState_
            lhs75_.f13 = rhs78_
            d_617_v232_ = rhs79_
            lhs76_.f0 = rhs80_
            d_624_v233_: _dafny.Map
            d_624_v233_ = _dafny.Map({((d_304_v5_)[len(d_514_v175_)] if (len(d_514_v175_)) in (d_304_v5_) else len(d_302_v3_)): d_581_cf49_})
            d_624_v233_ = (d_624_v233_).set(len(d_514_v175_), False)
        elif source12_.is_DC22:
            d_625___mcc_h23_ = source12_.cf41
            d_626_cf41_ = d_625___mcc_h23_
            d_627_v234_: C7
            nw105_ = C7()
            nw105_.ctor__()
            d_627_v234_ = nw105_
            d_628_v235_: _dafny.Map
            d_628_v235_ = _dafny.Map({425: d_627_v234_})
            d_629_v236_: _dafny.Seq
            d_629_v236_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_301_v2_}), d_312_v10_])
            d_630_v237_: D1
            d_630_v237_ = D1_DC5((d_629_v236_)[default__.safeIndex((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_423_v96_.f29 for d_631_i18_ in range(default__.abs(-203))]))).cardinality, len(d_629_v236_))])
            d_632_v238_: _dafny.Map
            d_632_v238_ = _dafny.Map({len(d_312_v10_): d_630_v237_})
            d_633_v239_: _dafny.Map
            d_633_v239_ = d_632_v238_
            d_634_v240_: D8
            d_634_v240_ = D8_DC25(not(d_301_v2_), d_431_v103_, d_434_v106_, d_633_v239_, len((d_629_v236_)[default__.safeIndex(d_423_v96_.f30, len(d_629_v236_))]))
            d_635_v241_: _dafny.Map
            d_635_v241_ = _dafny.Map({default__.fm42(_dafny.Set({726, d_303_v4_, d_423_v96_.f29, (0) - (len((d_634_v240_).cf50)), len(d_302_v3_)}), d_309_globalState_): d_303_v4_})
            rhs81_ = d_301_v2_
            rhs82_ = (len(d_628_v235_)) - (len(d_635_v241_))
            rhs83_ = (0) - (d_303_v4_)
            lhs77_ = d_309_globalState_
            lhs78_ = d_309_globalState_
            lhs79_ = d_309_globalState_
            lhs77_.f9 = rhs81_
            lhs78_.f0 = rhs82_
            lhs79_.f8 = rhs83_
            (d_309_globalState_).f5 = len(d_304_v5_)
            index75_ = default__.safeIndex(885, (d_310_v8_).length(0))
            (d_310_v8_)[index75_] = d_298_v0_
            index76_ = default__.safeIndex(885, (d_310_v8_).length(0))
            (d_310_v8_)[index76_] = d_298_v0_
            arr0_ = (d_310_v8_)[default__.safeIndex(885, (d_310_v8_).length(0))]
            index77_ = default__.safeIndex(614, ((d_310_v8_)[default__.safeIndex(885, (d_310_v8_).length(0))]).length(0))
            arr0_[index77_] = (default__.fm5(d_303_v4_, d_301_v2_, False, d_309_globalState_) if d_301_v2_ else d_301_v2_)
            arr1_ = (d_310_v8_)[default__.safeIndex(885, (d_310_v8_).length(0))]
            index78_ = default__.safeIndex(614, ((d_310_v8_)[default__.safeIndex(885, (d_310_v8_).length(0))]).length(0))
            arr1_[index78_] = (-433) >= (default__.safeModuloInt(d_303_v4_, (d_433_v105_).fm9(d_423_v96_.f30, len(d_302_v3_), d_301_v2_, d_309_globalState_)))
        elif True:
            d_636___mcc_h24_ = source12_.cf54
            d_637_cf54_ = d_636___mcc_h24_
            (d_309_globalState_).f0 = default__.safeModuloInt((0) - (d_303_v4_), len(d_514_v175_))
            d_638_v242_: C0
            nw106_ = C0()
            nw106_.ctor__()
            d_638_v242_ = nw106_
            d_639_v243_: _dafny.Map
            d_639_v243_ = _dafny.Map({d_434_v106_: d_638_v242_})
            d_639_v243_ = (d_639_v243_).set(d_434_v106_, d_638_v242_)
            d_640_v244_: D7
            d_640_v244_ = D7_DC18(d_303_v4_, d_306_v7_)
            source14_ = d_640_v244_
            if source14_.is_DC18:
                d_641___mcc_h31_ = source14_.cf34
                d_642___mcc_h32_ = source14_.cf35
                d_643_cf35_ = d_642___mcc_h32_
                d_644_cf34_ = d_641___mcc_h31_
                d_645_v245_: _dafny.Map
                d_645_v245_ = _dafny.Map({d_533_v194_: (d_423_v96_.f29) - (d_303_v4_)})
                d_645_v245_ = (d_645_v245_).set(d_533_v194_, d_423_v96_.f30)
                (d_309_globalState_).f25 = (len((d_514_v175_) + (_dafny.SeqWithoutIsStrInference([d_301_v2_])))) != (d_303_v4_)
                (d_309_globalState_).f25 = (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_301_v2_, True, d_301_v2_}))]))) >= ((0) - ((d_433_v105_).fm9(d_303_v4_, d_423_v96_.f30, d_301_v2_, d_309_globalState_)))
                (d_309_globalState_).f9 = d_301_v2_
            elif source14_.is_DC19:
                d_646___mcc_h33_ = source14_.cf36
                d_647___mcc_h34_ = source14_.cf37
                d_648_cf37_ = d_647___mcc_h34_
                d_649_cf36_ = d_646___mcc_h33_
                d_650_v246_: _dafny.Array
                nw107_ = _dafny.Array(_dafny.Map({}), 27)
                d_650_v246_ = nw107_
                index79_ = default__.safeIndex(891, (d_650_v246_).length(0))
                (d_650_v246_)[index79_] = d_304_v5_
                index80_ = default__.safeIndex(891, (d_650_v246_).length(0))
                (d_650_v246_)[index80_] = d_304_v5_
                (d_423_v96_).f30 = d_423_v96_.f30
                d_651_v247_: _dafny.MultiSet
                d_651_v247_ = _dafny.MultiSet([False])
                d_652_v248_: _dafny.Map
                d_652_v248_ = _dafny.Map({d_651_v247_: D3_DC10(d_540_v199_)})
                d_653_v249_: _dafny.Map
                d_653_v249_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgirbomb"))).set(default__.safeIndex(len(d_514_v175_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgirbomb")))), d_434_v106_): D12_DC36(d_652_v248_)})
                d_654_v250_: D12
                d_654_v250_ = D12_DC36(d_652_v248_)
                d_655_v251_: _dafny.Seq
                d_655_v251_ = _dafny.SeqWithoutIsStrInference([d_653_v249_, d_653_v249_, d_653_v249_, (d_653_v249_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsgljx")), d_654_v250_)])
                d_656_v252_: _dafny.Array
                nw108_ = _dafny.Array(None, 7)
                nw108_[int(0)] = d_653_v249_
                nw108_[int(1)] = d_653_v249_
                nw108_[int(2)] = d_653_v249_
                nw108_[int(3)] = (d_655_v251_)[default__.safeIndex(d_423_v96_.f29, len(d_655_v251_))]
                nw108_[int(4)] = d_653_v249_
                nw108_[int(5)] = d_653_v249_
                nw108_[int(6)] = d_653_v249_
                d_656_v252_ = nw108_
                index81_ = default__.safeIndex(278, (d_656_v252_).length(0))
                (d_656_v252_)[index81_] = _dafny.Map({d_302_v3_: d_654_v250_})
                index82_ = default__.safeIndex(278, (d_656_v252_).length(0))
                (d_656_v252_)[index82_] = _dafny.Map({d_302_v3_: d_654_v250_})
                d_657_v253_: D24
                d_657_v253_ = D24_DC71(d_303_v4_, default__.fm36(d_301_v2_, d_301_v2_, d_301_v2_, d_309_globalState_), d_301_v2_)
                d_658_v254_: int
                d_659_v255_: _dafny.MultiSet
                d_660_v256_: _dafny.Array
                out18_: int
                out19_: _dafny.MultiSet
                out20_: _dafny.Array
                out18_, out19_, out20_ = (d_423_v96_).m2(((d_657_v253_).cf133) >= ((0) - (len(_dafny.SeqWithoutIsStrInference([d_434_v106_ for d_661_i19_ in range(default__.abs(-636))])))), d_313_v11_, d_309_globalState_)
                d_658_v254_ = out18_
                d_659_v255_ = out19_
                d_660_v256_ = out20_
            elif source14_.is_DC20:
                d_662___mcc_h35_ = source14_.cf38
                d_663___mcc_h36_ = source14_.cf39
                d_664_cf39_ = d_663___mcc_h36_
                d_665_cf38_ = d_662___mcc_h35_
                (d_309_globalState_).f13 = (default__.fm2(d_664_cf39_, d_309_globalState_)) + (_dafny.SeqWithoutIsStrInference([d_434_v106_ for d_666_i20_ in range(default__.abs(827))]))
                d_306_v7_ = d_306_v7_
                d_429_v101_ = (d_429_v101_).set((611) <= (d_423_v96_.f30), d_423_v96_.f30)
                d_667_v257_: int
                d_668_v258_: _dafny.MultiSet
                d_669_v259_: _dafny.Array
                out21_: int
                out22_: _dafny.MultiSet
                out23_: _dafny.Array
                out21_, out22_, out23_ = (d_423_v96_).m2(d_301_v2_, d_313_v11_, d_309_globalState_)
                d_667_v257_ = out21_
                d_668_v258_ = out22_
                d_669_v259_ = out23_
            elif source14_.is_DC17:
                d_670___mcc_h37_ = source14_.cf33
                d_671_cf33_ = d_670___mcc_h37_
                d_434_v106_ = d_434_v106_
                d_672_v260_: int
                d_673_v261_: _dafny.MultiSet
                d_674_v262_: _dafny.Array
                out24_: int
                out25_: _dafny.MultiSet
                out26_: _dafny.Array
                out24_, out25_, out26_ = (d_423_v96_).m2(not(d_301_v2_), default__.fm49((0) - (d_423_v96_.f29), d_309_globalState_), d_309_globalState_)
                d_672_v260_ = out24_
                d_673_v261_ = out25_
                d_674_v262_ = out26_
                index83_ = default__.safeIndex(94, (d_534_v195_).length(0))
                (d_534_v195_)[index83_] = d_532_v193_
                index84_ = default__.safeIndex(94, (d_534_v195_).length(0))
                (d_534_v195_)[index84_] = d_532_v193_
                (d_309_globalState_).f9 = d_301_v2_
            elif True:
                d_675___mcc_h38_ = source14_.cf40
                d_676_cf40_ = d_675___mcc_h38_
                d_677_v263_: _dafny.Array
                nw109_ = _dafny.Array(None, 22)
                d_677_v263_ = nw109_
                d_678_v264_: _dafny.Seq
                d_678_v264_ = _dafny.SeqWithoutIsStrInference([d_677_v263_, d_677_v263_, d_677_v263_, (d_677_v263_ if d_301_v2_ else d_677_v263_)])
                d_678_v264_ = d_678_v264_
                d_679_v265_: D16
                d_679_v265_ = D16_DC45((d_433_v105_).fm9(d_423_v96_.f29, d_423_v96_.f30, d_301_v2_, d_309_globalState_), True, d_303_v4_, len(d_311_v9_), d_301_v2_)
                d_680_v266_: _dafny.MultiSet
                d_680_v266_ = _dafny.MultiSet([d_679_v265_, d_679_v265_, d_679_v265_, d_679_v265_])
                d_681_v267_: _dafny.Seq
                d_681_v267_ = _dafny.SeqWithoutIsStrInference([D16_DC45(d_423_v96_.f29, d_301_v2_, d_423_v96_.f30, d_423_v96_.f30, False), d_679_v265_])
                d_680_v266_ = _dafny.MultiSet((d_681_v267_ if d_301_v2_ else d_681_v267_))
                def iife45_():
                    coll33_ = _dafny.Map()
                    compr_33_: _dafny.Seq
                    for compr_33_ in (_dafny.SeqWithoutIsStrInference([d_514_v175_ for d_682_i21_ in range(default__.abs(609))])).Elements:
                        d_683_v268_: _dafny.Seq = compr_33_
                        if (d_683_v268_) in (_dafny.SeqWithoutIsStrInference([d_514_v175_ for d_682_i21_ in range(default__.abs(609))])):
                            coll33_[d_683_v268_] = True
                    return _dafny.Map(coll33_)
                (d_309_globalState_).f9 = ((0) - (default__.safeDivisionInt(875, (0) - (len(d_514_v175_))))) <= (((d_542_v201_).fm6(default__.fm18(d_423_v96_.f30, d_309_globalState_), len(default__.fm47(d_309_globalState_)), d_311_v9_, iife45_()
                , d_309_globalState_)) - (d_303_v4_))
                d_684_v269_: C6
                out27_: C6
                out27_ = default__.m8(d_301_v2_, d_514_v175_, 210, (d_514_v175_) + ((D2_DC7(_dafny.SeqWithoutIsStrInference([d_301_v2_]))).cf12), d_309_globalState_)
                d_684_v269_ = out27_
            d_685_v270_: T2
            nw110_ = C4()
            nw110_.ctor__(d_301_v2_, d_303_v4_, default__.fm2(False, d_309_globalState_), d_304_v5_)
            d_685_v270_ = nw110_
            d_686_v271_: _dafny.Seq
            d_686_v271_ = _dafny.SeqWithoutIsStrInference([d_685_v270_])
            source15_ = D17_DC48((d_686_v271_)[default__.safeIndex(d_423_v96_.f30, len(d_686_v271_))], len((_dafny.SeqWithoutIsStrInference([d_301_v2_])) + (d_514_v175_)), d_434_v106_)
            if source15_.is_DC48:
                d_687___mcc_h39_ = source15_.cf89
                d_688___mcc_h40_ = source15_.cf90
                d_689___mcc_h41_ = source15_.cf91
                d_690_cf91_ = d_689___mcc_h41_
                d_691_cf90_ = d_688___mcc_h40_
                d_692_cf89_ = d_687___mcc_h39_
                (d_309_globalState_).f0 = (d_423_v96_.f30) - ((0) - (d_426_v98_))
                d_693_v272_: _dafny.Map
                d_693_v272_ = _dafny.Map({d_303_v4_: False})
                index85_ = default__.safeIndex(521, (d_306_v7_).length(0))
                (d_306_v7_)[index85_] = (d_638_v242_).fm16(701, len(d_693_v272_), d_423_v96_.f30, d_301_v2_, d_309_globalState_)
                index86_ = default__.safeIndex(521, (d_306_v7_).length(0))
                (d_306_v7_)[index86_] = d_423_v96_.f29
                d_694_v273_: D24
                d_694_v273_ = D24_DC71((d_306_v7_)[default__.safeIndex(521, (d_306_v7_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqkcubwy")), d_301_v2_)
                (d_309_globalState_).f5 = ((0) - (((0) - (d_423_v96_.f29)) - (len(d_540_v199_))) if (d_694_v273_).cf135 else (d_303_v4_) + (d_423_v96_.f29))
                d_694_v273_ = default__.fm62(d_423_v96_.f29, d_690_cf91_, d_423_v96_.f29, d_309_globalState_)
            elif True:
                d_695___mcc_h42_ = source15_.cf88
                d_696_cf88_ = d_695___mcc_h42_
                d_697_v274_: _dafny.Array
                nw111_ = _dafny.Array(None, 22)
                d_697_v274_ = nw111_
                index87_ = default__.safeIndex(864, (d_697_v274_).length(0))
                (d_697_v274_)[index87_] = d_433_v105_
                d_698_v275_: _dafny.Map
                d_698_v275_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_301_v2_, d_301_v2_, d_301_v2_, True]): d_514_v175_})
                index88_ = default__.safeIndex(864, (d_697_v274_).length(0))
                rhs84_ = d_301_v2_
                rhs85_ = (d_433_v105_ if d_301_v2_ else d_433_v105_)
                rhs86_ = d_698_v275_
                rhs87_ = (d_540_v199_) | (d_540_v199_)
                lhs80_ = d_309_globalState_
                lhs81_ = d_697_v274_
                lhs82_ = default__.safeIndex(864, (d_697_v274_).length(0))
                lhs80_.f25 = rhs84_
                lhs81_[lhs82_] = rhs85_
                d_698_v275_ = rhs86_
                d_540_v199_ = rhs87_
                (d_309_globalState_).f8 = 622
                d_699_v276_: _dafny.Array
                nw112_ = _dafny.Array(_dafny.Set({}), 13)
                d_699_v276_ = nw112_
                index89_ = default__.safeIndex(849, (d_699_v276_).length(0))
                (d_699_v276_)[index89_] = d_696_cf88_
                index90_ = default__.safeIndex(849, (d_699_v276_).length(0))
                (d_699_v276_)[index90_] = d_696_cf88_
                d_700_v277_: D10
                d_700_v277_ = D10_DC31(d_301_v2_, d_311_v9_, d_301_v2_, len(d_302_v3_))
                pat_let_tv3_ = d_311_v9_
                pat_let_tv4_ = d_434_v106_
                def iife46_(_pat_let6_0):
                    def iife47_(d_701_dt__update__tmp_h3_):
                        def iife48_(_pat_let7_0):
                            def iife49_(d_702_dt__update_hcf59_h0_):
                                def iife50_(_pat_let8_0):
                                    def iife51_(d_704_dt__update_hcf61_h0_):
                                        return D10_DC31((d_701_dt__update__tmp_h3_).cf58, d_702_dt__update_hcf59_h0_, (d_701_dt__update__tmp_h3_).cf60, d_704_dt__update_hcf61_h0_)
                                    return iife51_(_pat_let8_0)
                                return iife50_(len(_dafny.SeqWithoutIsStrInference([pat_let_tv4_ for d_703_i22_ in range(default__.abs(910))])))
                            return iife49_(_pat_let7_0)
                        return iife48_(pat_let_tv3_)
                    return iife47_(_pat_let6_0)
                (d_309_globalState_).f5 = (0) - ((d_638_v242_).fm16((d_423_v96_.f29) - ((d_638_v242_).fm16(d_423_v96_.f30, len(_dafny.Set({d_533_v194_})), 379, d_301_v2_, d_309_globalState_)), d_423_v96_.f30, (iife46_(d_700_v277_)).cf61, not(d_301_v2_), d_309_globalState_))
        d_705_v278_: int
        d_706_v279_: _dafny.MultiSet
        d_707_v280_: _dafny.Array
        out28_: int
        out29_: _dafny.MultiSet
        out30_: _dafny.Array
        out28_, out29_, out30_ = (d_542_v201_).m2(d_301_v2_, (d_313_v11_ if d_301_v2_ else d_313_v11_), d_309_globalState_)
        d_705_v278_ = out28_
        d_706_v279_ = out29_
        d_707_v280_ = out30_
        d_708_v281_: _dafny.Map
        d_708_v281_ = _dafny.Map({d_431_v103_: default__.safeModuloInt(647, d_303_v4_)})
        d_708_v281_ = d_708_v281_
        hi3_ = d_423_v96_.f29
        for d_709_i23_ in range(default__.safeDivisionInt(d_705_v278_, d_423_v96_.f29), hi3_):
            index91_ = default__.safeIndex(422, (d_306_v7_).length(0))
            (d_306_v7_)[index91_] = d_423_v96_.f29
            index92_ = default__.safeIndex(422, (d_306_v7_).length(0))
            (d_306_v7_)[index92_] = d_423_v96_.f30
            d_710_v282_: C4
            nw113_ = C4()
            nw113_.ctor__(d_301_v2_, d_423_v96_.f29, d_302_v3_, _dafny.Map({d_423_v96_.f29: d_423_v96_.f30}))
            d_710_v282_ = nw113_
            d_711_v283_: _dafny.Seq
            d_711_v283_ = _dafny.SeqWithoutIsStrInference([d_710_v282_])
            d_712_v284_: _dafny.Map
            d_712_v284_ = _dafny.Map({d_434_v106_: (d_710_v282_).f32})
            d_713_v285_: _dafny.Map
            d_713_v285_ = _dafny.Map({d_711_v283_: d_712_v284_})
            d_713_v285_ = (d_713_v285_).set((d_711_v283_) + (d_711_v283_), default__.fm63((d_710_v282_).f32, d_311_v9_, d_309_globalState_))
            d_714_v286_: _dafny.Set
            d_714_v286_ = _dafny.Set({d_298_v0_, d_298_v0_})
            pat_let_tv5_ = d_714_v286_
            d_715_v287_: _dafny.Seq
            def iife52_(_pat_let9_0):
                def iife53_(d_716_dt__update__tmp_h4_):
                    def iife54_(_pat_let10_0):
                        def iife55_(d_717_dt__update_hcf92_h0_):
                            return D18_DC49(d_717_dt__update_hcf92_h0_)
                        return iife55_(_pat_let10_0)
                    return iife54_(pat_let_tv5_)
                return iife53_(_pat_let9_0)
            d_715_v287_ = _dafny.SeqWithoutIsStrInference([iife52_(D18_DC49(d_714_v286_))])
            d_718_v288_: _dafny.Map
            d_718_v288_ = _dafny.Map({_dafny.Map({d_298_v0_: (d_302_v3_)[default__.safeIndex(d_423_v96_.f30, len(d_302_v3_))]}): d_715_v287_})
            d_719_v289_: _dafny.Map
            d_719_v289_ = _dafny.Map({((d_304_v5_)[d_423_v96_.f30] if (d_423_v96_.f30) in (d_304_v5_) else d_709_i23_): _dafny.Map({d_298_v0_: _dafny.CodePoint('r')})})
            d_720_v290_: _dafny.Map
            d_720_v290_ = _dafny.Map({d_298_v0_: d_434_v106_})
            d_718_v288_ = (d_718_v288_).set(((d_719_v289_)[d_423_v96_.f30] if (d_423_v96_.f30) in (d_719_v289_) else d_720_v290_), d_715_v287_)
            index93_ = default__.safeIndex(422, (d_306_v7_).length(0))
            (d_306_v7_)[index93_] = ((d_303_v4_) + (len(d_514_v175_))) * (len(d_431_v103_))
        _dafny.print(_dafny.string_of((d_298_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_298_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[0]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[1]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[2]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[3]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[4]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[5]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[6]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[7]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[8]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[9]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[10]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[11]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[12]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[13]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[14]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[15]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[16]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[17]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[18]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[19]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[20]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_299_v1_)[21]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_301_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_302_v3_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_303_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v5_) == (_dafny.Map({655: 655}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_v6_) == (_dafny.Map({True: _dafny.Map({655: 655})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_306_v7_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_309_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_.f1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_.f1)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_309_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[0]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[1]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[2]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[3]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[4]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[5]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[6]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[7]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[8]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[9]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[10]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[11]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[12]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[13]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[14]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[15]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[16]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[17]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[18]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[19]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[20]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_309_globalState_).f7)[21]) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_309_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_309_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_309_globalState_.f13).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_309_globalState_.f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_.f18) == (_dafny.Map({True: _dafny.Map({655: 655})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_309_globalState_).f19)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_309_globalState_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_309_globalState_.f23).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_309_globalState_.f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_309_globalState_.f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v9_) == (_dafny.SeqWithoutIsStrInference([-655, 731]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v10_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v11_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_416_v89_).cf36))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_416_v89_).cf37))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v90_).cf40).cf36))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_417_v90_).cf40).cf37))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_418_v91_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_419_v92_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_420_v93_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_421_v94_).cf56).cf55)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_422_v95_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_423_v96_.f29))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_423_v96_.f30))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_424_v97_)[0]) == (_dafny.Map({False: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_426_v98_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_427_v99_) == (_dafny.MultiSet([D0_DC0(-854)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_429_v101_) == (_dafny.Map({False: 655}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_430_v102_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: 655}), _dafny.Map({False: 655}), _dafny.Map({False: 655}), _dafny.Map({False: 655}), _dafny.Map({False: 655})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_431_v103_) == (_dafny.Map({False: 5}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_432_v104_) == (_dafny.Map({_dafny.Map({False: 5}): _dafny.CodePoint('e')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_433_v105_).f28) == (_dafny.Map({_dafny.Map({False: 5}): _dafny.CodePoint('e')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_434_v106_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_514_v175_) == (_dafny.SeqWithoutIsStrInference([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_515_v176_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_532_v193_).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_533_v194_) == (_dafny.Set({1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[0]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[1]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[2]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[3]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[4]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[5]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[6]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[7]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[8]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[9]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[10]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[11]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[12]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[13]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[14]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[15]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[16]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[17]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[18]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[19]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_534_v195_)[20]).cf119))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_537_v196_).cf129))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_537_v196_).cf130))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_538_v197_) == (_dafny.Set({D23_DC68(True, True)}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_539_v198_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_540_v199_) == (_dafny.Map({False: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_541_v200_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, False, True]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_543_v202_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_544_v203_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_545_v205_) == (_dafny.Map({1: _dafny.CodePoint('l')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_546_v207_) == (_dafny.MultiSet([_dafny.Map({1: _dafny.CodePoint('l')}), _dafny.Map({1: _dafny.CodePoint('l')}), _dafny.Map({1: _dafny.CodePoint('l')}), _dafny.Map({0: _dafny.CodePoint('l')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_549_v208_).cf46))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_549_v208_).cf47))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_549_v208_).cf48))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_705_v278_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_706_v279_) == (_dafny.MultiSet([D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42), D0_DC0(42)]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_708_v281_) == (_dafny.Map({_dafny.Map({False: 5}): 647}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(int(0))
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
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D0_DC4)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({self.cf1.VerbatimString(True)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
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
        return lambda: D1_DC6(False, _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC6(D1, NamedTuple('DC6', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf16', Any), ('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
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
        return lambda: D3_DC11(_dafny.Map({}), False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC11(D3, NamedTuple('DC11', [('cf21', Any), ('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC13(D4, NamedTuple('DC13', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC14(D5, NamedTuple('DC14', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC16(D6, NamedTuple('DC16', [('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC18(int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC18(D7, NamedTuple('DC18', [('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(False, _dafny.Map({}), _dafny.MultiSet({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)

class D8_DC23(D8, NamedTuple('DC23', [('cf42', Any), ('cf43', Any), ('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf49', Any), ('cf50', Any), ('cf51', Any), ('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC26(D8, NamedTuple('DC26', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC28()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D9_DC29)

class D9_DC28(D9, NamedTuple('DC28', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC29(D9, NamedTuple('DC29', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC29({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC29) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC31(False, _dafny.Seq({}), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D10_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D10_DC32)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D10_DC30)

class D10_DC31(D10, NamedTuple('DC31', [('cf58', Any), ('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC31({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC31) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC32(D10, NamedTuple('DC32', [('cf62', Any), ('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC32({_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC32) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC30(D10, NamedTuple('DC30', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC34(_dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D11_DC34)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D11_DC35)

class D11_DC34(D11, NamedTuple('DC34', [('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC34({_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC34) and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC33(D11, NamedTuple('DC33', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC35(D11, NamedTuple('DC35', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC35({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC35) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC37(int(0), int(0), D3.default()(), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D12_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D12_DC38)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)

class D12_DC37(D12, NamedTuple('DC37', [('cf70', Any), ('cf71', Any), ('cf72', Any), ('cf73', Any), ('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37({_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37) and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73 and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC38(D12, NamedTuple('DC38', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC38'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC38)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC40(_dafny.Array(None, 0), _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D13_DC40)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D13_DC39)

class D13_DC40(D13, NamedTuple('DC40', [('cf76', Any), ('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC40({_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC40) and self.cf76 == __o.cf76 and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC39(D13, NamedTuple('DC39', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC39({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC39) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)

class D14_DC41(D14, NamedTuple('DC41', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC43()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)

class D15_DC43(D15, NamedTuple('DC43', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC42(D15, NamedTuple('DC42', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC45(int(0), False, int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D16_DC46)

class D16_DC45(D16, NamedTuple('DC45', [('cf82', Any), ('cf83', Any), ('cf84', Any), ('cf85', Any), ('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf82 == __o.cf82 and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC44(D16, NamedTuple('DC44', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC46(D16, NamedTuple('DC46', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC46({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC46) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC48(None, int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D17_DC48)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D17_DC47)

class D17_DC48(D17, NamedTuple('DC48', [('cf89', Any), ('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC48({_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC48) and self.cf89 == __o.cf89 and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC47(D17, NamedTuple('DC47', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC47({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC47) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC50(int(0), False, _dafny.Seq({}), _dafny.CodePoint('D'), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D18_DC50)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D18_DC49)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D18_DC51)

class D18_DC50(D18, NamedTuple('DC50', [('cf93', Any), ('cf94', Any), ('cf95', Any), ('cf96', Any), ('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)}, {_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf93 == __o.cf93 and self.cf94 == __o.cf94 and self.cf95 == __o.cf95 and self.cf96 == __o.cf96 and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC49(D18, NamedTuple('DC49', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC49({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC49) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC51(D18, NamedTuple('DC51', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC51({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC51) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC53(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D19_DC53)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D19_DC52)

class D19_DC53(D19, NamedTuple('DC53', [('cf100', Any), ('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC53({_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC53) and self.cf100 == __o.cf100 and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC52(D19, NamedTuple('DC52', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC52({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC52) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC55(int(0), False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D20_DC55)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D20_DC54)

class D20_DC55(D20, NamedTuple('DC55', [('cf103', Any), ('cf104', Any), ('cf105', Any), ('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC55({_dafny.string_of(self.cf103)}, {_dafny.string_of(self.cf104)}, {_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC55) and self.cf103 == __o.cf103 and self.cf104 == __o.cf104 and self.cf105 == __o.cf105 and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC54(D20, NamedTuple('DC54', [('cf102', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC54({_dafny.string_of(self.cf102)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC54) and self.cf102 == __o.cf102
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC57(_dafny.CodePoint('D'), int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D21_DC57)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D21_DC58)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D21_DC59)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D21_DC56)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D21_DC60)

class D21_DC57(D21, NamedTuple('DC57', [('cf108', Any), ('cf109', Any), ('cf110', Any), ('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC57({_dafny.string_of(self.cf108)}, {_dafny.string_of(self.cf109)}, {_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC57) and self.cf108 == __o.cf108 and self.cf109 == __o.cf109 and self.cf110 == __o.cf110 and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC58(D21, NamedTuple('DC58', [('cf112', Any), ('cf113', Any), ('cf114', Any), ('cf115', Any), ('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC58({_dafny.string_of(self.cf112)}, {_dafny.string_of(self.cf113)}, {_dafny.string_of(self.cf114)}, {_dafny.string_of(self.cf115)}, {_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC58) and self.cf112 == __o.cf112 and self.cf113 == __o.cf113 and self.cf114 == __o.cf114 and self.cf115 == __o.cf115 and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC59(D21, NamedTuple('DC59', [])):
    def __dafnystr__(self) -> str:
        return f'D21.DC59'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC59)
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC56(D21, NamedTuple('DC56', [('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC56({_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC56) and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC60(D21, NamedTuple('DC60', [('cf117', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC60({_dafny.string_of(self.cf117)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC60) and self.cf117 == __o.cf117
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC62(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D22_DC62)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D22_DC63)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D22_DC64)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D22_DC61)

class D22_DC62(D22, NamedTuple('DC62', [('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC62({_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC62) and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC63(D22, NamedTuple('DC63', [('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC63({_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC63) and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC64(D22, NamedTuple('DC64', [('cf121', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC64({_dafny.string_of(self.cf121)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC64) and self.cf121 == __o.cf121
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC61(D22, NamedTuple('DC61', [('cf118', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC61({_dafny.string_of(self.cf118)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC61) and self.cf118 == __o.cf118
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC66(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D23_DC66)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D23_DC67)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D23_DC68)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D23_DC65)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D23_DC69)

class D23_DC66(D23, NamedTuple('DC66', [('cf123', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC66({_dafny.string_of(self.cf123)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC66) and self.cf123 == __o.cf123
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC67(D23, NamedTuple('DC67', [('cf124', Any), ('cf125', Any), ('cf126', Any), ('cf127', Any), ('cf128', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC67({_dafny.string_of(self.cf124)}, {_dafny.string_of(self.cf125)}, {_dafny.string_of(self.cf126)}, {_dafny.string_of(self.cf127)}, {_dafny.string_of(self.cf128)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC67) and self.cf124 == __o.cf124 and self.cf125 == __o.cf125 and self.cf126 == __o.cf126 and self.cf127 == __o.cf127 and self.cf128 == __o.cf128
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC68(D23, NamedTuple('DC68', [('cf129', Any), ('cf130', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC68({_dafny.string_of(self.cf129)}, {_dafny.string_of(self.cf130)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC68) and self.cf129 == __o.cf129 and self.cf130 == __o.cf130
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC65(D23, NamedTuple('DC65', [('cf122', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC65({_dafny.string_of(self.cf122)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC65) and self.cf122 == __o.cf122
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC69(D23, NamedTuple('DC69', [('cf131', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC69({_dafny.string_of(self.cf131)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC69) and self.cf131 == __o.cf131
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC71(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D24_DC71)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D24_DC72)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D24_DC73)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D24_DC70)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D24_DC74)

class D24_DC71(D24, NamedTuple('DC71', [('cf133', Any), ('cf134', Any), ('cf135', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC71({_dafny.string_of(self.cf133)}, {self.cf134.VerbatimString(True)}, {_dafny.string_of(self.cf135)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC71) and self.cf133 == __o.cf133 and self.cf134 == __o.cf134 and self.cf135 == __o.cf135
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC72(D24, NamedTuple('DC72', [('cf136', Any), ('cf137', Any), ('cf138', Any), ('cf139', Any), ('cf140', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC72({_dafny.string_of(self.cf136)}, {_dafny.string_of(self.cf137)}, {_dafny.string_of(self.cf138)}, {self.cf139.VerbatimString(True)}, {_dafny.string_of(self.cf140)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC72) and self.cf136 == __o.cf136 and self.cf137 == __o.cf137 and self.cf138 == __o.cf138 and self.cf139 == __o.cf139 and self.cf140 == __o.cf140
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC73(D24, NamedTuple('DC73', [('cf141', Any), ('cf142', Any), ('cf143', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC73({_dafny.string_of(self.cf141)}, {_dafny.string_of(self.cf142)}, {_dafny.string_of(self.cf143)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC73) and self.cf141 == __o.cf141 and self.cf142 == __o.cf142 and self.cf143 == __o.cf143
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC70(D24, NamedTuple('DC70', [('cf132', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC70({_dafny.string_of(self.cf132)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC70) and self.cf132 == __o.cf132
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC74(D24, NamedTuple('DC74', [('cf144', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC74({_dafny.string_of(self.cf144)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC74) and self.cf144 == __o.cf144
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC76(False, _dafny.MultiSet({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D25_DC76)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D25_DC77)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D25_DC75)

class D25_DC76(D25, NamedTuple('DC76', [('cf146', Any), ('cf147', Any), ('cf148', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC76({_dafny.string_of(self.cf146)}, {_dafny.string_of(self.cf147)}, {_dafny.string_of(self.cf148)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC76) and self.cf146 == __o.cf146 and self.cf147 == __o.cf147 and self.cf148 == __o.cf148
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC77(D25, NamedTuple('DC77', [])):
    def __dafnystr__(self) -> str:
        return f'D25.DC77'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC77)
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC75(D25, NamedTuple('DC75', [('cf145', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC75({_dafny.string_of(self.cf145)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC75) and self.cf145 == __o.cf145
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC79(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D26_DC79)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D26_DC78)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D26_DC80)

class D26_DC79(D26, NamedTuple('DC79', [('cf150', Any), ('cf151', Any), ('cf152', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC79({_dafny.string_of(self.cf150)}, {_dafny.string_of(self.cf151)}, {_dafny.string_of(self.cf152)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC79) and self.cf150 == __o.cf150 and self.cf151 == __o.cf151 and self.cf152 == __o.cf152
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC78(D26, NamedTuple('DC78', [('cf149', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC78({_dafny.string_of(self.cf149)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC78) and self.cf149 == __o.cf149
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC80(D26, NamedTuple('DC80', [('cf153', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC80({_dafny.string_of(self.cf153)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC80) and self.cf153 == __o.cf153
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC82(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC82(self) -> bool:
        return isinstance(self, D27_DC82)
    @property
    def is_DC83(self) -> bool:
        return isinstance(self, D27_DC83)
    @property
    def is_DC84(self) -> bool:
        return isinstance(self, D27_DC84)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D27_DC81)
    @property
    def is_DC85(self) -> bool:
        return isinstance(self, D27_DC85)

class D27_DC82(D27, NamedTuple('DC82', [('cf155', Any), ('cf156', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC82({_dafny.string_of(self.cf155)}, {_dafny.string_of(self.cf156)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC82) and self.cf155 == __o.cf155 and self.cf156 == __o.cf156
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC83(D27, NamedTuple('DC83', [('cf157', Any), ('cf158', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC83({_dafny.string_of(self.cf157)}, {_dafny.string_of(self.cf158)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC83) and self.cf157 == __o.cf157 and self.cf158 == __o.cf158
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC84(D27, NamedTuple('DC84', [('cf159', Any), ('cf160', Any), ('cf161', Any), ('cf162', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC84({_dafny.string_of(self.cf159)}, {_dafny.string_of(self.cf160)}, {_dafny.string_of(self.cf161)}, {_dafny.string_of(self.cf162)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC84) and self.cf159 == __o.cf159 and self.cf160 == __o.cf160 and self.cf161 == __o.cf161 and self.cf162 == __o.cf162
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC81(D27, NamedTuple('DC81', [('cf154', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC81({_dafny.string_of(self.cf154)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC81) and self.cf154 == __o.cf154
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC85(D27, NamedTuple('DC85', [('cf163', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC85({_dafny.string_of(self.cf163)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC85) and self.cf163 == __o.cf163
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC87(None, _dafny.Seq({}), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC87(self) -> bool:
        return isinstance(self, D28_DC87)
    @property
    def is_DC88(self) -> bool:
        return isinstance(self, D28_DC88)
    @property
    def is_DC86(self) -> bool:
        return isinstance(self, D28_DC86)

class D28_DC87(D28, NamedTuple('DC87', [('cf165', Any), ('cf166', Any), ('cf167', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC87({_dafny.string_of(self.cf165)}, {_dafny.string_of(self.cf166)}, {_dafny.string_of(self.cf167)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC87) and self.cf165 == __o.cf165 and self.cf166 == __o.cf166 and self.cf167 == __o.cf167
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC88(D28, NamedTuple('DC88', [('cf168', Any), ('cf169', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC88({_dafny.string_of(self.cf168)}, {_dafny.string_of(self.cf169)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC88) and self.cf168 == __o.cf168 and self.cf169 == __o.cf169
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC86(D28, NamedTuple('DC86', [('cf164', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC86({_dafny.string_of(self.cf164)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC86) and self.cf164 == __o.cf164
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC90(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC90(self) -> bool:
        return isinstance(self, D29_DC90)
    @property
    def is_DC89(self) -> bool:
        return isinstance(self, D29_DC89)

class D29_DC90(D29, NamedTuple('DC90', [('cf171', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC90({_dafny.string_of(self.cf171)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC90) and self.cf171 == __o.cf171
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC89(D29, NamedTuple('DC89', [('cf170', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC89({_dafny.string_of(self.cf170)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC89) and self.cf170 == __o.cf170
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, p1, globalState):
        pass

    def m2(self, p0, p1, globalState):
        pass


class T1:
    pass

class T2:
    pass
    def m3(self, p0, p1, globalState):
        pass

    def m4(self, p0, p1, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f1: _dafny.Array = _dafny.Array(None, 0)
        self.f5: int = int(0)
        self.f8: int = int(0)
        self.f9: bool = False
        self.f13: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f17: int = int(0)
        self.f18: _dafny.Map = _dafny.Map({})
        self.f23: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f24: int = int(0)
        self.f25: bool = False
        self._f2: int = int(0)
        self._f3: bool = False
        self._f4: bool = False
        self._f6: str = _dafny.CodePoint('D')
        self._f7: _dafny.Array = _dafny.Array(None, 0)
        self._f10: bool = False
        self._f11: bool = False
        self._f12: int = int(0)
        self._f14: bool = False
        self._f15: int = int(0)
        self._f16: bool = False
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f20: int = int(0)
        self._f21: int = int(0)
        self._f22: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25):
        (self).f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self).f17 = f17
        (self).f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self)._f22 = f22
        (self).f23 = f23
        (self).f24 = f24
        (self).f25 = f25

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
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
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
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
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

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm16(self, p0, p1, p2, p3, globalState):
        source16_ = D3_DC10(_dafny.Map({False: False}))
        if source16_.is_DC11:
            d_721___mcc_h0_ = source16_.cf21
            d_722___mcc_h1_ = source16_.cf22
            d_723___mcc_h2_ = source16_.cf23
            d_724___mcc_h3_ = source16_.cf24
            d_725_cf24_ = d_724___mcc_h3_
            d_726_cf23_ = d_723___mcc_h2_
            d_727_cf22_ = d_722___mcc_h1_
            d_728_cf21_ = d_721___mcc_h0_
            return ((_dafny.MultiSet([d_726_cf23_, len(_dafny.SeqWithoutIsStrInference([d_726_cf23_, d_726_cf23_]))])) | (_dafny.MultiSet([d_726_cf23_, len(_dafny.SeqWithoutIsStrInference([d_726_cf23_, d_725_cf24_])), d_725_cf24_]))).cardinality
        elif source16_.is_DC12:
            d_729___mcc_h4_ = source16_.cf25
            d_730___mcc_h5_ = source16_.cf26
            d_731___mcc_h6_ = source16_.cf27
            d_732_cf27_ = d_731___mcc_h6_
            d_733_cf26_ = d_730___mcc_h5_
            d_734_cf25_ = d_729___mcc_h4_
            return d_733_cf26_
        elif True:
            d_735___mcc_h7_ = source16_.cf20
            d_736_cf20_ = d_735___mcc_h7_
            if True:
                return -807
            elif True:
                return 377


class C1(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm6(self, p0, p1, p2, p3, globalState):
        return (-551) * ((len(_dafny.SeqWithoutIsStrInference([919, 754, 508]))) + (-777))

    def fm7(self, p0, p1, globalState):
        return (623) + (564)

    def fm15(self, p0, p1, globalState):
        return not (not(True)) or ((True) == (False))

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        d_737_v0_: int
        d_737_v0_ = -868
        (globalState).f17 = (((0) - (d_737_v0_)) - (d_737_v0_)) - (d_737_v0_)
        d_738_i0_: int
        d_738_i0_ = 0
        with _dafny.label("2"):
            while (self).fm15((d_737_v0_ if p1 else d_737_v0_), (d_737_v0_) * (d_737_v0_), globalState):
                with _dafny.c_label("2"):
                    if (d_738_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_738_i0_ = (d_738_i0_) + (1)
                    d_739_v1_: str
                    d_739_v1_ = _dafny.CodePoint('d')
                    d_739_v1_ = d_739_v1_
                    d_740_v2_: C0
                    nw114_ = C0()
                    nw114_.ctor__()
                    d_740_v2_ = nw114_
                    d_741_v3_: _dafny.Set
                    d_741_v3_ = _dafny.Set({p1})
                    d_742_v4_: D1
                    d_742_v4_ = D1_DC5(d_741_v3_)
                    pat_let_tv6_ = d_741_v3_
                    d_743_v5_: _dafny.Map
                    def iife56_(_pat_let11_0):
                        def iife57_(d_744_dt__update__tmp_h0_):
                            def iife58_(_pat_let12_0):
                                def iife59_(d_745_dt__update_hcf8_h0_):
                                    return D1_DC5(d_745_dt__update_hcf8_h0_)
                                return iife59_(_pat_let12_0)
                            return iife58_(pat_let_tv6_)
                        return iife57_(_pat_let11_0)
                    d_743_v5_ = _dafny.Map({d_737_v0_: iife56_(d_742_v4_)})
                    d_746_v6_: _dafny.Map
                    d_746_v6_ = d_743_v5_
                    (globalState).f0 = len((d_746_v6_))
                    d_747_v7_: _dafny.Seq
                    d_747_v7_ = _dafny.SeqWithoutIsStrInference([(0) - (d_737_v0_)])
                    (globalState).f0 = (0) - ((d_747_v7_)[default__.safeIndex(default__.safeDivisionInt(d_737_v0_, d_737_v0_), len(d_747_v7_))])
                    pass
            pass
        r1 = False
        d_748_v8_: _dafny.MultiSet
        d_748_v8_ = _dafny.MultiSet([p1])
        d_749_v9_: _dafny.Map
        d_749_v9_ = _dafny.Map({(d_748_v8_).cardinality: p1})
        d_750_v10_: _dafny.Array
        def lambda22_(d_751_i1_):
            return _dafny.CodePoint('n')

        init12_ = lambda22_
        nw115_ = _dafny.Array(None, 19)
        for i0_12_ in range(nw115_.length(0)):
            nw115_[i0_12_] = init12_(i0_12_)
        d_750_v10_ = nw115_
        d_752_v11_: _dafny.Seq
        d_752_v11_ = _dafny.SeqWithoutIsStrInference([not(p1)])
        d_753_v12_: _dafny.Seq
        d_753_v12_ = _dafny.SeqWithoutIsStrInference([(self).fm15(len(_dafny.Map({d_750_v10_: p0})), (0) - ((0) - (d_737_v0_)), globalState), (d_752_v11_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([not(p1)])), len(d_752_v11_))], True])
        d_754_v13_: D2
        d_754_v13_ = D2_DC8(False, ((d_748_v8_)[not(p1)] if (not(p1)) in (d_748_v8_) else len(d_749_v9_)), (_dafny.MultiSet(d_752_v11_)).cardinality)
        d_754_v13_ = d_754_v13_
        if p1:
            d_755_v15_: _dafny.Seq
            d_755_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "boancx"))
            d_756_v16_: _dafny.Set
            d_756_v16_ = _dafny.Set({p1})
            d_757_v17_: _dafny.Seq
            d_757_v17_ = _dafny.SeqWithoutIsStrInference([default__.fm0(False, d_755_v15_, d_756_v16_, d_737_v0_, globalState)])
            d_758_v18_: _dafny.Map
            d_758_v18_ = _dafny.Map({d_752_v11_: d_737_v0_})
            d_759_v19_: _dafny.Map
            def iife60_():
                coll34_ = _dafny.Map()
                compr_34_: _dafny.Seq
                for compr_34_ in (d_757_v17_).Elements:
                    d_760_v14_: _dafny.Seq = compr_34_
                    if (d_760_v14_) in (d_757_v17_):
                        coll34_[d_760_v14_] = d_737_v0_
                return _dafny.Map(coll34_)
            d_759_v19_ = _dafny.Map({(iife60_()
            ) | (d_758_v18_): (p1) not in (default__.fm17(globalState))})
            d_759_v19_ = (d_759_v19_).set(_dafny.Map({d_753_v12_: d_737_v0_}), True)
            if p1:
                (globalState).f17 = (d_737_v0_ if p1 else d_737_v0_)
                (globalState).f9 = ((d_753_v12_) + (d_752_v11_)) < (d_753_v12_)
                d_761_v20_: _dafny.Seq
                d_761_v20_ = _dafny.SeqWithoutIsStrInference([d_755_v15_])
                d_762_v21_: _dafny.Map
                d_762_v21_ = _dafny.Map({p1: _dafny.MultiSet(d_761_v20_)})
                d_763_v22_: _dafny.MultiSet
                d_763_v22_ = _dafny.MultiSet([d_755_v15_])
                (globalState).f0 = (((d_762_v21_)[p1] if (p1) in (d_762_v21_) else d_763_v22_)).cardinality
                d_764_v24_: _dafny.Array
                nw116_ = _dafny.Array(None, 4)
                def iife61_():
                    coll35_ = _dafny.Map()
                    compr_35_: int
                    for compr_35_ in _dafny.IntegerRange(16, 413):
                        d_765_v23_: int = compr_35_
                        if ((16) <= (d_765_v23_)) and ((d_765_v23_) < (413)):
                            coll35_[(d_765_v23_) - (d_737_v0_)] = p1
                    return _dafny.Map(coll35_)
                nw116_[int(0)] = iife61_()
                
                nw116_[int(1)] = default__.fm18(d_737_v0_, globalState)
                nw116_[int(2)] = _dafny.Map({len(d_755_v15_): (d_753_v12_)[default__.safeIndex(d_737_v0_, len(d_753_v12_))]})
                nw116_[int(3)] = d_749_v9_
                d_764_v24_ = nw116_
                index94_ = default__.safeIndex(933, (d_764_v24_).length(0))
                (d_764_v24_)[index94_] = d_749_v9_
                d_766_v25_: str
                d_766_v25_ = _dafny.CodePoint('m')
                d_767_v26_: _dafny.Map
                d_767_v26_ = _dafny.Map({False: d_737_v0_})
                d_768_v27_: _dafny.Map
                d_768_v27_ = _dafny.Map({p0: (d_767_v26_) | (d_767_v26_)})
                d_769_v28_: _dafny.Seq
                d_769_v28_ = d_761_v20_
                d_770_v30_: _dafny.Map
                d_770_v30_ = _dafny.Map({d_737_v0_: d_737_v0_})
                d_771_v31_: _dafny.Seq
                d_771_v31_ = _dafny.SeqWithoutIsStrInference([len(d_770_v30_)])
                index95_ = default__.safeIndex(933, (d_764_v24_).length(0))
                def iife62_():
                    coll36_ = _dafny.Map()
                    compr_36_: int
                    for compr_36_ in (d_771_v31_).Elements:
                        d_773_v29_: int = compr_36_
                        if (d_773_v29_) in (d_771_v31_):
                            coll36_[default__.safeModuloInt(d_773_v29_, len(d_755_v15_))] = d_737_v0_
                    return _dafny.Map(coll36_)
                rhs88_ = _dafny.Map({d_737_v0_: not(p1)})
                rhs89_ = (((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "me")) for d_772_i2_ in range(default__.abs(571))])) + ((d_769_v28_))).set(default__.safeIndex(len(iife62_()
                ), len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "me")) for d_774_i2_ in range(default__.abs(571))])) + ((d_769_v28_)))), d_755_v15_)) + (_dafny.SeqWithoutIsStrInference([d_755_v15_ for d_775_i3_ in range(default__.abs(580))]))
                rhs90_ = _dafny.CodePoint('v')
                rhs91_ = d_768_v27_
                rhs92_ = not(p1)
                lhs83_ = d_764_v24_
                lhs84_ = default__.safeIndex(933, (d_764_v24_).length(0))
                lhs85_ = globalState
                lhs83_[lhs84_] = rhs88_
                d_761_v20_ = rhs89_
                d_766_v25_ = rhs90_
                d_768_v27_ = rhs91_
                lhs85_.f25 = rhs92_
                d_776_v32_: _dafny.Array
                nw117_ = _dafny.Array(_dafny.Seq({}), 26)
                d_776_v32_ = nw117_
                d_777_v33_: _dafny.Seq
                d_777_v33_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0])
                index96_ = default__.safeIndex(972, (d_776_v32_).length(0))
                (d_776_v32_)[index96_] = d_777_v33_
                index97_ = default__.safeIndex(972, (d_776_v32_).length(0))
                (d_776_v32_)[index97_] = d_777_v33_
            elif True:
                (globalState).f5 = (0) - ((0) - (d_737_v0_))
                d_778_v34_: _dafny.Map
                d_778_v34_ = _dafny.Map({d_737_v0_: d_737_v0_})
                d_778_v34_ = (d_778_v34_).set((0) - ((d_737_v0_) - (d_737_v0_)), d_737_v0_)
                r0 = (d_737_v0_ if not(not(not (not(p1)) or (not(p1)))) else d_737_v0_)
                (globalState).f9 = (p1) and (p1)
                d_779_v35_: _dafny.Array
                def lambda23_(d_780_i4_):
                    return True

                init13_ = lambda23_
                nw118_ = _dafny.Array(None, 27)
                for i0_13_ in range(nw118_.length(0)):
                    nw118_[i0_13_] = init13_(i0_13_)
                d_779_v35_ = nw118_
                index98_ = default__.safeIndex(230, (d_779_v35_).length(0))
                (d_779_v35_)[index98_] = p1
                index99_ = default__.safeIndex(230, (d_779_v35_).length(0))
                (d_779_v35_)[index99_] = p1
            (globalState).f5 = 712
            (globalState).f25 = True
            d_737_v0_ = ((d_737_v0_ if p1 else d_737_v0_)) * ((0) - (d_737_v0_))
        elif True:
            d_781_v36_: _dafny.Seq
            d_781_v36_ = _dafny.SeqWithoutIsStrInference([d_737_v0_])
            d_782_v37_: D6
            d_782_v37_ = D6_DC15(d_781_v36_)
            d_749_v9_ = (d_749_v9_).set((0) - ((175) - (len((d_782_v37_).cf30))), p1)
            (globalState).f9 = True
            d_737_v0_ = d_737_v0_
            if not(p1):
                d_783_v38_: _dafny.Array
                nw119_ = _dafny.Array(False, 12)
                d_783_v38_ = nw119_
                d_784_v39_: _dafny.Seq
                d_784_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "encpnwnn"))
                index100_ = default__.safeIndex(443, (d_783_v38_).length(0))
                (d_783_v38_)[index100_] = (d_784_v39_) != (d_784_v39_)
                index101_ = default__.safeIndex(443, (d_783_v38_).length(0))
                (d_783_v38_)[index101_] = (self).fm15(d_737_v0_, (0) - (d_737_v0_), globalState)
                index102_ = default__.safeIndex(443, (d_783_v38_).length(0))
                (d_783_v38_)[index102_] = ((d_737_v0_) + (d_737_v0_)) >= (len(d_784_v39_))
                d_785_v40_: C0
                nw120_ = C0()
                nw120_.ctor__()
                d_785_v40_ = nw120_
                d_786_v41_: C0
                nw121_ = C0()
                nw121_.ctor__()
                d_786_v41_ = nw121_
                d_787_v42_: str
                d_787_v42_ = _dafny.CodePoint('d')
                d_787_v42_ = d_787_v42_
            elif True:
                d_788_v43_: _dafny.Array
                def lambda24_(d_789_p1_):
                    def lambda25_(d_790_i5_):
                        return d_789_p1_

                    return lambda25_

                init14_ = lambda24_(p1)
                nw122_ = _dafny.Array(None, 22)
                for i0_14_ in range(nw122_.length(0)):
                    nw122_[i0_14_] = init14_(i0_14_)
                d_788_v43_ = nw122_
                d_791_v44_: _dafny.Set
                d_791_v44_ = _dafny.Set({d_788_v43_, d_788_v43_})
                d_791_v44_ = d_791_v44_
                index103_ = default__.safeIndex(878, (p0).length(0))
                (p0)[index103_] = d_737_v0_
                index104_ = default__.safeIndex(878, (p0).length(0))
                (p0)[index104_] = default__.safeModuloInt((d_737_v0_ if p1 else d_737_v0_), d_737_v0_)
                (globalState).f9 = p1
                d_792_v45_: _dafny.Array
                nw123_ = _dafny.Array(_dafny.Set({}), 20)
                d_792_v45_ = nw123_
                d_793_v46_: _dafny.Set
                d_793_v46_ = _dafny.Set({d_737_v0_})
                index105_ = default__.safeIndex(880, (d_792_v45_).length(0))
                (d_792_v45_)[index105_] = d_793_v46_
                index106_ = default__.safeIndex(880, (d_792_v45_).length(0))
                (d_792_v45_)[index106_] = d_793_v46_
                index107_ = default__.safeIndex(878, (p0).length(0))
                (p0)[index107_] = 881
            d_737_v0_ = (d_737_v0_) + ((0) - ((self).fm7(d_737_v0_, p1, globalState)))
        d_794_v47_: _dafny.Seq
        d_794_v47_ = _dafny.SeqWithoutIsStrInference([d_737_v0_, d_737_v0_])
        d_795_v48_: D6
        d_795_v48_ = D6_DC15(d_794_v47_)
        source17_ = d_795_v48_
        if source17_.is_DC16:
            d_796___mcc_h0_ = source17_.cf31
            d_797___mcc_h1_ = source17_.cf32
            d_798_cf32_ = d_797___mcc_h1_
            d_799_cf31_ = d_796___mcc_h0_
            d_800_v49_: str
            d_800_v49_ = _dafny.CodePoint('g')
            d_801_v50_: _dafny.Map
            d_801_v50_ = _dafny.Map({False: d_800_v49_})
            d_802_v51_: _dafny.Set
            d_802_v51_ = _dafny.Set({(len(d_801_v50_)) != ((0) - (d_737_v0_)), p1, True})
            d_803_v52_: _dafny.Map
            d_803_v52_ = _dafny.Map({(0) - (((0) - (d_799_cf31_)) * (d_737_v0_)): (d_802_v51_).intersection(d_802_v51_)})
            d_804_v53_: _dafny.Seq
            d_804_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjsfbwvme"))
            rhs93_ = ((d_803_v52_)[(d_737_v0_) + (len(d_794_v47_))] if ((d_737_v0_) + (len(d_794_v47_))) in (d_803_v52_) else d_802_v51_)
            rhs94_ = (0) - (d_799_cf31_)
            rhs95_ = default__.safeDivisionInt(d_737_v0_, d_799_cf31_)
            rhs96_ = d_804_v53_
            rhs97_ = (d_799_cf31_) * (d_799_cf31_)
            lhs86_ = globalState
            lhs87_ = globalState
            lhs88_ = globalState
            d_802_v51_ = rhs93_
            lhs86_.f0 = rhs94_
            d_737_v0_ = rhs95_
            lhs87_.f23 = rhs96_
            lhs88_.f0 = rhs97_
            r1 = p1
            d_805_v54_: _dafny.Array
            nw124_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_805_v54_ = nw124_
            d_806_v55_: _dafny.Array
            nw125_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
            d_806_v55_ = nw125_
            index108_ = default__.safeIndex(61, (d_805_v54_).length(0))
            (d_805_v54_)[index108_] = d_806_v55_
            index109_ = default__.safeIndex(61, (d_805_v54_).length(0))
            (d_805_v54_)[index109_] = d_806_v55_
            d_800_v49_ = d_800_v49_
        elif True:
            d_807___mcc_h2_ = source17_.cf30
            d_808_cf30_ = d_807___mcc_h2_
            d_752_v11_ = d_753_v12_
            d_809_v56_: _dafny.Map
            d_809_v56_ = _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_810_i6_ in range(default__.abs(179))]))})
            d_811_v57_: _dafny.Set
            d_811_v57_ = _dafny.Set({d_737_v0_, d_737_v0_, d_737_v0_, d_737_v0_, len(d_809_v56_)})
            def iife63_():
                coll37_ = _dafny.Set()
                compr_37_: int
                for compr_37_ in (d_811_v57_).Elements:
                    d_812_v58_: int = compr_37_
                    if (d_812_v58_) in (d_811_v57_):
                        coll37_ = coll37_.union(_dafny.Set([(d_812_v58_) - (270)]))
                return _dafny.Set(coll37_)
            if (iife63_()
            ).issubset(d_811_v57_):
                d_813_v59_: str
                d_813_v59_ = _dafny.CodePoint('f')
                d_814_v60_: _dafny.Map
                d_814_v60_ = _dafny.Map({d_813_v59_: -981})
                (globalState).f0 = (d_737_v0_ if not((d_752_v11_)[default__.safeIndex(len(d_814_v60_), len(d_752_v11_))]) else d_737_v0_)
                d_815_v61_: _dafny.Array
                def lambda26_(d_816_p1_, d_817_v59_, d_818_v0_):
                    def lambda27_(d_819_i7_):
                        return _dafny.Map({(D1_DC6(d_816_p1_, d_817_v59_, d_818_v0_)).cf9: _dafny.Map({-450: D1_DC5(_dafny.Set({d_816_p1_}))})})

                    return lambda27_

                init15_ = lambda26_(p1, d_813_v59_, d_737_v0_)
                nw126_ = _dafny.Array(None, 11)
                for i0_15_ in range(nw126_.length(0)):
                    nw126_[i0_15_] = init15_(i0_15_)
                d_815_v61_ = nw126_
                d_820_v62_: _dafny.Seq
                d_820_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                d_821_v63_: D1
                d_821_v63_ = D1_DC5(_dafny.Set({p1}))
                d_822_v64_: _dafny.Map
                d_822_v64_ = _dafny.Map({len(d_820_v62_): d_821_v63_})
                d_823_v65_: _dafny.Map
                d_823_v65_ = _dafny.Map({p1: d_822_v64_})
                index110_ = default__.safeIndex(875, (d_815_v61_).length(0))
                (d_815_v61_)[index110_] = (d_823_v65_) | (((d_823_v65_).set(not(p1), d_822_v64_)).set((d_754_v13_).cf13, d_822_v64_))
                index111_ = default__.safeIndex(875, (d_815_v61_).length(0))
                (d_815_v61_)[index111_] = (d_823_v65_) | ((_dafny.Map({p1: d_822_v64_})).set((self).fm15(-670, (0) - (d_737_v0_), globalState), d_822_v64_))
                d_824_v66_: _dafny.Map
                d_824_v66_ = _dafny.Map({p1: not(p1)})
                d_825_v67_: D3
                d_825_v67_ = D3_DC10((d_824_v66_).set(p1, True))
                d_826_v68_: D2
                d_826_v68_ = D2_DC9(len(d_749_v9_), not(p1), d_737_v0_, d_737_v0_)
                d_827_v69_: D0
                d_827_v69_ = D0_DC3(d_737_v0_, d_737_v0_)
                d_828_v70_: _dafny.Map
                d_828_v70_ = _dafny.Map({d_825_v67_: p1})
                d_825_v67_ = default__.fm19(d_826_v68_, d_827_v69_, default__.fm20(((d_828_v70_)[d_825_v67_] if (d_825_v67_) in (d_828_v70_) else p1), d_813_v59_, d_737_v0_, d_813_v59_, globalState), p1, globalState)
                d_813_v59_ = _dafny.CodePoint('p')
                d_813_v59_ = d_813_v59_
            elif True:
                d_829_v71_: _dafny.Array
                nw127_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_829_v71_ = nw127_
                d_830_v72_: _dafny.Seq
                d_830_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrdqs"))
                index112_ = default__.safeIndex(762, (d_829_v71_).length(0))
                (d_829_v71_)[index112_] = d_830_v72_
                d_831_v73_: str
                d_831_v73_ = _dafny.CodePoint('i')
                index113_ = default__.safeIndex(762, (d_829_v71_).length(0))
                (d_829_v71_)[index113_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adaeodbg"))).set(default__.safeIndex((len(d_830_v72_)) + (d_737_v0_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adaeodbg")))), d_831_v73_)
                d_832_v74_: _dafny.MultiSet
                d_832_v74_ = _dafny.MultiSet([d_737_v0_, d_737_v0_])
                d_833_v75_: _dafny.Map
                d_833_v75_ = _dafny.Map({d_753_v12_: False})
                d_834_v77_: _dafny.Seq
                def iife64_():
                    coll38_ = _dafny.Map()
                    compr_38_: int
                    for compr_38_ in _dafny.IntegerRange(2, 501):
                        d_835_v76_: int = compr_38_
                        if ((2) <= (d_835_v76_)) and ((d_835_v76_) < (501)):
                            coll38_[(d_835_v76_) - ((_dafny.MultiSet(d_794_v47_)).cardinality)] = len(_dafny.Map({p1: (d_752_v11_)[default__.safeIndex(((d_832_v74_)[((d_809_v56_)[p1] if (p1) in (d_809_v56_) else d_737_v0_)] if (((d_809_v56_)[p1] if (p1) in (d_809_v56_) else d_737_v0_)) in (d_832_v74_) else d_737_v0_), len(d_752_v11_))]}))
                    return _dafny.Map(coll38_)
                d_834_v77_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_737_v0_: d_737_v0_}), iife64_()
                ])
                d_836_v78_: D0
                d_836_v78_ = D0_DC1(d_830_v72_, len(d_830_v72_), len(d_834_v77_))
                (globalState).f9 = (_dafny.MultiSet([(self).fm6(d_749_v9_, d_737_v0_, d_794_v47_, d_833_v75_, globalState), d_737_v0_, len((d_836_v78_).cf1)])).issubset(d_832_v74_)
                d_837_v79_: _dafny.Set
                d_837_v79_ = _dafny.Set({True})
                d_838_v80_: _dafny.MultiSet
                d_838_v80_ = _dafny.MultiSet([(d_808_cf30_) + (_dafny.SeqWithoutIsStrInference([d_737_v0_, d_737_v0_, 310])), default__.fm21(len(d_837_v79_), d_737_v0_, p1, not(True), globalState), _dafny.SeqWithoutIsStrInference([d_737_v0_]), d_794_v47_])
                d_838_v80_ = (d_838_v80_) | (_dafny.MultiSet([d_808_cf30_, d_808_cf30_, d_794_v47_, d_794_v47_, d_808_cf30_]))
                (globalState).f0 = d_737_v0_
                (globalState).f9 = p1
            rhs98_ = p1
            rhs99_ = d_737_v0_
            rhs100_ = len(d_811_v57_)
            lhs89_ = globalState
            lhs90_ = globalState
            lhs91_ = globalState
            lhs89_.f25 = rhs98_
            lhs90_.f0 = rhs99_
            lhs91_.f17 = rhs100_
            (globalState).f17 = d_737_v0_
        r0 = d_737_v0_
        r1 = p1
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_839_v0_: int
        d_839_v0_ = 42
        d_840_v2_: _dafny.Map
        d_840_v2_ = _dafny.Map({d_839_v0_: p0})
        d_841_v3_: str
        d_841_v3_ = _dafny.CodePoint('s')
        def iife65_():
            coll39_ = _dafny.Map()
            compr_39_: int
            for compr_39_ in _dafny.IntegerRange(533, 824):
                d_844_v1_: int = compr_39_
                if ((533) <= (d_844_v1_)) and ((d_844_v1_) < (824)):
                    coll39_[(d_844_v1_) * (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([354])): not(False)})))] = _dafny.Map({d_839_v0_: p0})
            return _dafny.Map(coll39_)
        (globalState).f9 = not (((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_842_i0_ in range(default__.abs(577))])).set(default__.safeIndex(d_839_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_843_i0_ in range(default__.abs(577))]))), d_841_v3_)) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svrrlefr")))) or ((self).fm15(d_839_v0_, len((iife65_()
        ).set(d_839_v0_, d_840_v2_)), globalState))
        d_845_v4_: _dafny.Set
        d_845_v4_ = _dafny.Set({d_839_v0_})
        d_846_v5_: _dafny.Seq
        d_846_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
        d_847_v6_: _dafny.Map
        d_847_v6_ = _dafny.Map({d_845_v4_: d_846_v5_})
        d_848_v7_: _dafny.Seq
        d_848_v7_ = _dafny.SeqWithoutIsStrInference([d_846_v5_, d_846_v5_])
        d_847_v6_ = (d_847_v6_).set(_dafny.Set({d_839_v0_}), ((d_848_v7_)[default__.safeIndex(d_839_v0_, len(d_848_v7_))] if p0 else d_846_v5_))
        d_849_v8_: _dafny.Array
        nw128_ = _dafny.Array(int(0), 9)
        d_849_v8_ = nw128_
        index114_ = default__.safeIndex(733, (d_849_v8_).length(0))
        (d_849_v8_)[index114_] = d_839_v0_
        index115_ = default__.safeIndex(733, (d_849_v8_).length(0))
        (d_849_v8_)[index115_] = (d_839_v0_ if p0 else d_839_v0_)
        if not(not(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_850_i1_ in range(default__.abs(482))])) + (d_846_v5_)) < (d_846_v5_))):
            source18_ = (d_848_v7_) + (d_848_v7_)
            d_851___mcc_h0_ = source18_
            d_852_cf29_ = d_851___mcc_h0_
            (globalState).f5 = 793
            index116_ = default__.safeIndex(733, (d_849_v8_).length(0))
            (d_849_v8_)[index116_] = (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]
            d_853_v9_: _dafny.Array
            nw129_ = _dafny.Array(False, 7)
            d_853_v9_ = nw129_
            (globalState).f1 = d_853_v9_
            (globalState).f25 = (p0) == (p0)
            d_854_v10_: _dafny.Map
            d_854_v10_ = _dafny.Map({d_841_v3_: d_839_v0_})
            rhs101_ = (d_839_v0_) + ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])
            rhs102_ = (default__.safeModuloInt((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], -473)) * (d_839_v0_)
            rhs103_ = ((d_854_v10_)[d_841_v3_] if (d_841_v3_) in (d_854_v10_) else (0) - (d_839_v0_))
            rhs104_ = not(not (True) or (p0))
            lhs92_ = globalState
            lhs93_ = globalState
            lhs94_ = globalState
            lhs92_.f17 = rhs101_
            r0 = rhs102_
            lhs93_.f24 = rhs103_
            lhs94_.f9 = rhs104_
            d_855_v11_: C0
            nw130_ = C0()
            nw130_.ctor__()
            d_855_v11_ = nw130_
            (globalState).f9 = (d_839_v0_) != (len(d_846_v5_))
            d_856_v12_: _dafny.Seq
            d_856_v12_ = _dafny.SeqWithoutIsStrInference([d_839_v0_, d_839_v0_, (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]])
            if ((self).fm7(d_839_v0_, p0, globalState)) >= ((d_839_v0_ if p0 else (d_856_v12_)[default__.safeIndex((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], len(d_856_v12_))])):
                d_857_v13_: _dafny.Seq
                d_857_v13_ = _dafny.SeqWithoutIsStrInference([p0])
                d_858_v14_: D0
                d_858_v14_ = D0_DC3(default__.safeDivisionInt(d_839_v0_, (0) - ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])), (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])
                d_859_v15_: _dafny.Map
                d_859_v15_ = _dafny.Map({True: p0})
                d_860_v16_: D3
                d_860_v16_ = D3_DC10(d_859_v15_)
                d_861_v17_: _dafny.Set
                d_861_v17_ = _dafny.Set({True, p0})
                d_862_v18_: _dafny.MultiSet
                d_862_v18_ = _dafny.MultiSet([len(d_861_v17_)])
                d_863_v19_: _dafny.Map
                d_863_v19_ = _dafny.Map({p0: d_862_v18_})
                rhs105_ = default__.safeModuloInt((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], d_839_v0_)
                rhs106_ = default__.fm0(p0, (((d_848_v7_)[default__.safeIndex((((d_863_v19_)[True] if (True) in (d_863_v19_) else _dafny.MultiSet(d_856_v12_))).cardinality, len(d_848_v7_))]).set(default__.safeIndex(625, len((d_848_v7_)[default__.safeIndex((((d_863_v19_)[True] if (True) in (d_863_v19_) else _dafny.MultiSet(d_856_v12_))).cardinality, len(d_848_v7_))])), d_841_v3_)) + ((d_846_v5_).set(default__.safeIndex(d_839_v0_, len(d_846_v5_)), d_841_v3_)), (d_861_v17_) - (d_861_v17_), (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], globalState)
                rhs107_ = d_858_v14_
                rhs108_ = d_860_v16_
                rhs109_ = d_839_v0_
                lhs95_ = globalState
                lhs96_ = globalState
                lhs95_.f24 = rhs105_
                d_857_v13_ = rhs106_
                d_858_v14_ = rhs107_
                d_860_v16_ = rhs108_
                lhs96_.f0 = rhs109_
                d_864_v20_: D6
                d_864_v20_ = D6_DC16(d_839_v0_, p0)
                d_865_v21_: _dafny.Set
                d_865_v21_ = _dafny.Set({d_864_v20_, d_864_v20_, d_864_v20_, d_864_v20_})
                d_866_v22_: _dafny.MultiSet
                d_866_v22_ = _dafny.MultiSet([d_841_v3_])
                d_867_v23_: _dafny.Array
                nw131_ = _dafny.Array(None, 4)
                nw131_[int(0)] = _dafny.Set({d_864_v20_})
                nw131_[int(1)] = (_dafny.Set({d_864_v20_})).intersection(d_865_v21_)
                nw131_[int(2)] = _dafny.Set({d_864_v20_, default__.fm22((d_866_v22_).cardinality, d_839_v0_, d_839_v0_, (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], globalState)})
                nw131_[int(3)] = _dafny.Set({D6_DC16((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], p0)})
                d_867_v23_ = nw131_
                pat_let_tv7_ = p0
                index117_ = default__.safeIndex(676, (d_867_v23_).length(0))
                def iife66_(_pat_let13_0):
                    def iife67_(d_868_dt__update__tmp_h0_):
                        def iife68_(_pat_let14_0):
                            def iife69_(d_869_dt__update_hcf32_h0_):
                                return D6_DC16((d_868_dt__update__tmp_h0_).cf31, d_869_dt__update_hcf32_h0_)
                            return iife69_(_pat_let14_0)
                        return iife68_(pat_let_tv7_)
                    return iife67_(_pat_let13_0)
                (d_867_v23_)[index117_] = _dafny.Set({iife66_(d_864_v20_)})
                index118_ = default__.safeIndex(676, (d_867_v23_).length(0))
                rhs110_ = (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]
                rhs111_ = d_865_v21_
                rhs112_ = d_860_v16_
                lhs97_ = globalState
                lhs98_ = d_867_v23_
                lhs99_ = default__.safeIndex(676, (d_867_v23_).length(0))
                lhs97_.f8 = rhs110_
                lhs98_[lhs99_] = rhs111_
                d_860_v16_ = rhs112_
                d_870_v24_: D0
                d_870_v24_ = D0_DC1(d_846_v5_, len(d_846_v5_), (0) - (len(d_846_v5_)))
                d_871_v25_: D0
                d_871_v25_ = D0_DC4(d_870_v24_)
                d_872_v26_: D0
                d_872_v26_ = D0_DC4((d_871_v25_).cf7)
                d_873_v27_: _dafny.Map
                d_873_v27_ = _dafny.Map({(0) - ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]): d_872_v26_})
                def iife70_(_pat_let15_0):
                    def iife71_(d_874_dt__update__tmp_h1_):
                        def iife72_(_pat_let16_0):
                            def iife73_(d_875_dt__update_hcf7_h0_):
                                return D0_DC4(d_875_dt__update_hcf7_h0_)
                            return iife73_(_pat_let16_0)
                        return iife72_(D0_DC2(True))
                    return iife71_(_pat_let15_0)
                d_873_v27_ = (d_873_v27_).set((0) - ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]), iife70_(d_872_v26_))
                d_876_v28_: _dafny.Array
                nw132_ = _dafny.Array(False, 26)
                d_876_v28_ = nw132_
                index119_ = default__.safeIndex(808, (d_876_v28_).length(0))
                (d_876_v28_)[index119_] = p0
                index120_ = default__.safeIndex(808, (d_876_v28_).length(0))
                index121_ = default__.safeIndex(733, (d_849_v8_).length(0))
                rhs113_ = not(p0)
                rhs114_ = (0) - ((0) - (default__.safeModuloInt((d_839_v0_) - ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]), d_839_v0_)))
                rhs115_ = (d_855_v11_).fm16((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]) + ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]), (0) - (d_839_v0_), p0, globalState)
                lhs100_ = d_876_v28_
                lhs101_ = default__.safeIndex(808, (d_876_v28_).length(0))
                lhs102_ = d_849_v8_
                lhs103_ = default__.safeIndex(733, (d_849_v8_).length(0))
                lhs104_ = globalState
                lhs100_[lhs101_] = rhs113_
                lhs102_[lhs103_] = rhs114_
                lhs104_.f5 = rhs115_
                index122_ = default__.safeIndex(733, (d_849_v8_).length(0))
                (d_849_v8_)[index122_] = (self).fm7((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], p0, globalState)
            elif True:
                d_877_v29_: _dafny.Array
                nw133_ = _dafny.Array(None, 20)
                nw133_[int(0)] = default__.fm20(p0, _dafny.CodePoint('j'), -554, _dafny.CodePoint('b'), globalState)
                nw133_[int(1)] = d_841_v3_
                nw133_[int(2)] = _dafny.CodePoint('u')
                nw133_[int(3)] = d_841_v3_
                nw133_[int(4)] = d_841_v3_
                nw133_[int(5)] = d_841_v3_
                nw133_[int(6)] = d_841_v3_
                nw133_[int(7)] = d_841_v3_
                nw133_[int(8)] = (d_841_v3_ if p0 else d_841_v3_)
                nw133_[int(9)] = d_841_v3_
                nw133_[int(10)] = _dafny.CodePoint('m')
                nw133_[int(11)] = d_841_v3_
                nw133_[int(12)] = d_841_v3_
                nw133_[int(13)] = _dafny.CodePoint('y')
                nw133_[int(14)] = d_841_v3_
                nw133_[int(15)] = (d_841_v3_ if p0 else default__.fm20(p0, d_841_v3_, 423, d_841_v3_, globalState))
                nw133_[int(16)] = d_841_v3_
                nw133_[int(17)] = d_841_v3_
                nw133_[int(18)] = default__.fm20(p0, d_841_v3_, d_839_v0_, d_841_v3_, globalState)
                nw133_[int(19)] = d_841_v3_
                d_877_v29_ = nw133_
                index123_ = default__.safeIndex(857, (d_877_v29_).length(0))
                (d_877_v29_)[index123_] = (d_846_v5_)[default__.safeIndex(d_839_v0_, len(d_846_v5_))]
                index124_ = default__.safeIndex(857, (d_877_v29_).length(0))
                rhs116_ = d_839_v0_
                rhs117_ = d_841_v3_
                lhs105_ = globalState
                lhs106_ = d_877_v29_
                lhs107_ = default__.safeIndex(857, (d_877_v29_).length(0))
                lhs105_.f17 = rhs116_
                lhs106_[lhs107_] = rhs117_
                index125_ = default__.safeIndex(733, (d_849_v8_).length(0))
                (d_849_v8_)[index125_] = default__.safeDivisionInt(default__.safeDivisionInt((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], len(d_856_v12_)), d_839_v0_)
                (globalState).f9 = p0
                d_878_v30_: D3
                d_878_v30_ = D3_DC12(p0, (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], p0)
                d_878_v30_ = d_878_v30_
                d_879_v31_: _dafny.Set
                d_879_v31_ = _dafny.Set({True})
                d_880_v32_: D1
                d_880_v32_ = D1_DC5(d_879_v31_)
                d_881_v33_: _dafny.Seq
                d_881_v33_ = _dafny.SeqWithoutIsStrInference([d_880_v32_])
                pat_let_tv8_ = d_879_v31_
                def iife74_(_pat_let17_0):
                    def iife75_(d_882_dt__update__tmp_h2_):
                        def iife76_(_pat_let18_0):
                            def iife77_(d_883_dt__update_hcf8_h0_):
                                return D1_DC5(d_883_dt__update_hcf8_h0_)
                            return iife77_(_pat_let18_0)
                        return iife76_(pat_let_tv8_)
                    return iife75_(_pat_let17_0)
                d_881_v33_ = _dafny.SeqWithoutIsStrInference([iife74_(d_880_v32_)])
        elif True:
            (globalState).f9 = not(p0)
            (globalState).f25 = ((self).fm7((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], p0, globalState)) > ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])
            d_884_v34_: _dafny.Seq
            d_884_v34_ = _dafny.SeqWithoutIsStrInference([p0])
            if (self).fm15(d_839_v0_, len((d_884_v34_) + (d_884_v34_)), globalState):
                d_885_v35_: _dafny.Map
                d_885_v35_ = _dafny.Map({p0: (0) - ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])})
                d_885_v35_ = (d_885_v35_).set(p0, (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])
                (globalState).f17 = (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]
                (globalState).f24 = (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]
                d_886_v36_: _dafny.Array
                def lambda28_(d_887_p0_):
                    def lambda29_(d_888_i2_):
                        return d_887_p0_

                    return lambda29_

                init16_ = lambda28_(p0)
                nw134_ = _dafny.Array(None, 6)
                for i0_16_ in range(nw134_.length(0)):
                    nw134_[i0_16_] = init16_(i0_16_)
                d_886_v36_ = nw134_
                index126_ = default__.safeIndex(587, (d_886_v36_).length(0))
                (d_886_v36_)[index126_] = p0
                index127_ = default__.safeIndex(587, (d_886_v36_).length(0))
                (d_886_v36_)[index127_] = (_dafny.MultiSet([p0])) == (_dafny.MultiSet(d_884_v34_))
                d_886_v36_ = d_886_v36_
            elif True:
                d_889_v37_: _dafny.Set
                d_889_v37_ = _dafny.Set({p0})
                d_890_v38_: D2
                d_890_v38_ = D2_DC8(p0, len((default__.fm0(True, d_846_v5_, d_889_v37_, (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], globalState)).set(default__.safeIndex((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], len(default__.fm0(True, d_846_v5_, d_889_v37_, (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], globalState))), p0)), (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])
                d_891_v39_: _dafny.Array
                nw135_ = _dafny.Array(False, 26)
                d_891_v39_ = nw135_
                d_892_v40_: _dafny.Map
                d_892_v40_ = _dafny.Map({p0: d_891_v39_})
                d_893_v41_: _dafny.Map
                d_893_v41_ = _dafny.Map({d_890_v38_: ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]) > (len(d_892_v40_))})
                d_893_v41_ = (d_893_v41_).set((d_890_v38_ if p0 else d_890_v38_), p0)
                d_894_v42_: C0
                nw136_ = C0()
                nw136_.ctor__()
                d_894_v42_ = nw136_
                (globalState).f8 = (((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]) - ((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])) + (len(d_846_v5_))
                (globalState).f8 = default__.safeModuloInt(352, d_839_v0_)
                d_895_v43_: _dafny.Array
                def lambda30_(d_896_v4_):
                    def lambda31_(d_897_i3_):
                        return d_896_v4_

                    return lambda31_

                init17_ = lambda30_(d_845_v4_)
                nw137_ = _dafny.Array(None, 3)
                for i0_17_ in range(nw137_.length(0)):
                    nw137_[i0_17_] = init17_(i0_17_)
                d_895_v43_ = nw137_
                index128_ = default__.safeIndex(214, (d_895_v43_).length(0))
                def iife78_():
                    coll40_ = _dafny.Set()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(909, 79):
                        d_898_v44_: int = compr_40_
                        if ((909) <= (d_898_v44_)) and ((d_898_v44_) < (79)):
                            coll40_ = coll40_.union(_dafny.Set([default__.safeModuloInt(d_898_v44_, d_839_v0_)]))
                    return _dafny.Set(coll40_)
                (d_895_v43_)[index128_] = iife78_()
                
                d_899_v45_: _dafny.MultiSet
                d_899_v45_ = _dafny.MultiSet([477])
                index129_ = default__.safeIndex(733, (d_849_v8_).length(0))
                index130_ = default__.safeIndex(214, (d_895_v43_).length(0))
                rhs118_ = default__.safeModuloInt((((d_899_v45_)[d_839_v0_] if (d_839_v0_) in (d_899_v45_) else (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])) + (d_839_v0_), d_839_v0_)
                rhs119_ = (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]
                rhs120_ = p0
                rhs121_ = (self).fm15(164, d_839_v0_, globalState)
                rhs122_ = ((d_845_v4_) - (d_845_v4_)) - ((d_845_v4_) | (d_845_v4_))
                lhs108_ = globalState
                lhs109_ = d_849_v8_
                lhs110_ = default__.safeIndex(733, (d_849_v8_).length(0))
                lhs111_ = globalState
                lhs112_ = globalState
                lhs113_ = d_895_v43_
                lhs114_ = default__.safeIndex(214, (d_895_v43_).length(0))
                lhs108_.f8 = rhs118_
                lhs109_[lhs110_] = rhs119_
                lhs111_.f9 = rhs120_
                lhs112_.f25 = rhs121_
                lhs113_[lhs114_] = rhs122_
            d_900_v46_: _dafny.Set
            d_900_v46_ = _dafny.Set({p0, not(not(p0)), p0, p0, not(p0)})
            d_901_v47_: _dafny.Set
            d_901_v47_ = _dafny.Set({d_900_v46_, _dafny.Set({p0}), d_900_v46_})
            (globalState).f25 = not((d_901_v47_).ispropersubset((d_901_v47_) - (d_901_v47_)))
            d_902_v48_: _dafny.MultiSet
            d_902_v48_ = _dafny.MultiSet([p0])
            rhs123_ = ((d_902_v48_) - (_dafny.MultiSet(default__.fm0(p0, d_846_v5_, d_900_v46_, (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))], globalState)))) == (_dafny.MultiSet([(D1_DC6(p0, d_841_v3_, d_839_v0_)).cf9]))
            rhs124_ = p0
            lhs115_ = globalState
            lhs116_ = globalState
            lhs115_.f9 = rhs123_
            lhs116_.f9 = rhs124_
        d_903_v49_: _dafny.Array
        nw138_ = _dafny.Array(D1.default()(), 18)
        d_903_v49_ = nw138_
        d_904_v50_: _dafny.MultiSet
        d_904_v50_ = _dafny.MultiSet([d_903_v49_, d_903_v49_, d_903_v49_, (d_903_v49_ if p0 else d_903_v49_), d_903_v49_])
        d_905_v51_: _dafny.Seq
        d_905_v51_ = _dafny.SeqWithoutIsStrInference([(d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]])
        rhs125_ = not(not(((self).fm15(d_839_v0_, len(d_905_v51_), globalState)) not in (default__.fm23(d_846_v5_, globalState))))
        rhs126_ = d_904_v50_
        rhs127_ = (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]
        lhs117_ = globalState
        lhs118_ = globalState
        lhs117_.f9 = rhs125_
        d_904_v50_ = rhs126_
        lhs118_.f5 = rhs127_
        d_906_v52_: _dafny.Seq
        d_906_v52_ = _dafny.SeqWithoutIsStrInference([p0])
        d_907_v53_: _dafny.Map
        d_907_v53_ = _dafny.Map({(d_841_v3_) in ((d_848_v7_)[default__.safeIndex(len(d_906_v52_), len(d_848_v7_))]): (p0) and (p0)})
        d_907_v53_ = (d_907_v53_).set(p0, p0)
        d_908_v54_: _dafny.Map
        d_908_v54_ = _dafny.Map({p0: (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]})
        d_909_v55_: _dafny.Map
        d_909_v55_ = _dafny.Map({p0: d_841_v3_})
        d_910_v56_: _dafny.MultiSet
        d_910_v56_ = _dafny.MultiSet([d_839_v0_, 495])
        d_911_v57_: _dafny.MultiSet
        d_911_v57_ = _dafny.MultiSet([len(d_909_v55_), (d_910_v56_).cardinality, d_839_v0_])
        d_912_v58_: _dafny.MultiSet
        d_912_v58_ = _dafny.MultiSet([d_911_v57_, d_911_v57_, _dafny.MultiSet([d_839_v0_])])
        r0 = ((len(d_908_v54_) if p0 else (d_912_v58_).cardinality) if p0 else (d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])
        d_913_v59_: _dafny.MultiSet
        d_913_v59_ = _dafny.MultiSet([D0_DC0((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])])
        d_914_v60_: D0
        d_914_v60_ = D0_DC0((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))])
        r1 = ((d_913_v59_).set(d_914_v60_, default__.abs((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]))).set(d_914_v60_, default__.abs((d_849_v8_)[default__.safeIndex(733, (d_849_v8_).length(0))]))
        d_915_v61_: _dafny.Array
        nw139_ = _dafny.Array(_dafny.Array(None, 0), 19)
        d_915_v61_ = nw139_
        r2 = d_915_v61_
        return r0, r1, r2


class C2(T2, T0):
    def  __init__(self):
        self._f31: _dafny.Map = _dafny.Map({})
        self._f33: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f31(self):
        return self._f31
    def ctor__(self, f33, f31):
        (self)._f33 = f33
        (self)._f31 = f31

    def fm12(self, p0, p1, p2, p3, globalState):
        def iife79_():
            def iife81_():
                coll43_ = _dafny.Set()
                compr_43_: _dafny.Set
                for compr_43_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({-677, 50, len(_dafny.Set({15}))}), _dafny.Set({175, (_dafny.MultiSet([not(True)])).cardinality}), _dafny.Set({419, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plbidf")))})])).Elements:
                    d_918_v1_: _dafny.Set = compr_43_
                    if (d_918_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({-677, 50, len(_dafny.Set({15}))}), _dafny.Set({175, (_dafny.MultiSet([not(True)])).cardinality}), _dafny.Set({419, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plbidf")))})])):
                        coll43_ = coll43_.union(_dafny.Set([d_918_v1_]))
                return _dafny.Set(coll43_)
            coll41_ = _dafny.Map()
            def iife80_():
                coll42_ = _dafny.Set()
                compr_42_: _dafny.Set
                for compr_42_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({-677, 50, len(_dafny.Set({15}))}), _dafny.Set({175, (_dafny.MultiSet([not(True)])).cardinality}), _dafny.Set({419, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plbidf")))})])).Elements:
                    d_916_v1_: _dafny.Set = compr_42_
                    if (d_916_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({-677, 50, len(_dafny.Set({15}))}), _dafny.Set({175, (_dafny.MultiSet([not(True)])).cardinality}), _dafny.Set({419, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plbidf")))})])):
                        coll42_ = coll42_.union(_dafny.Set([d_916_v1_]))
                return _dafny.Set(coll42_)
            compr_41_: _dafny.Set
            for compr_41_ in ((_dafny.Set({_dafny.Set({len(_dafny.Map({-767: True}))}), _dafny.Set({len(_dafny.Set({(self).f31, _dafny.Map({-69: len(_dafny.SeqWithoutIsStrInference([True]))}), (self).f31})), 899, len(_dafny.SeqWithoutIsStrInference([False]))})})) | (iife80_()
            )).Elements:
                d_917_v0_: _dafny.Set = compr_41_
                if (d_917_v0_) in ((_dafny.Set({_dafny.Set({len(_dafny.Map({-767: True}))}), _dafny.Set({len(_dafny.Set({(self).f31, _dafny.Map({-69: len(_dafny.SeqWithoutIsStrInference([True]))}), (self).f31})), 899, len(_dafny.SeqWithoutIsStrInference([False]))})})) | (iife81_()
                )):
                    coll41_[d_917_v0_] = not((_dafny.MultiSet([False, True, not((D2_DC9(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_919_i0_ in range(default__.abs(481))])), True, 342, -357)).cf17), True, False])).ispropersubset(_dafny.MultiSet([True])))
            return _dafny.Map(coll41_)
        return iife79_()
        

    def fm6(self, p0, p1, p2, p3, globalState):
        return len((_dafny.SeqWithoutIsStrInference([True])) + ((_dafny.SeqWithoutIsStrInference([False, True])) + (_dafny.SeqWithoutIsStrInference([False]))))

    def fm7(self, p0, p1, globalState):
        return (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): not(True)}))) - (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), True]))).cardinality) + (769))

    def m3(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_920_v0_: int
        d_920_v0_ = 28
        (globalState).f8 = default__.safeDivisionInt(d_920_v0_, (_dafny.MultiSet((self).f33)).cardinality)
        if p0:
            (globalState).f25 = not((p0) or (p0))
            source19_ = D3_DC10(_dafny.Map({p0: p0}))
            if source19_.is_DC11:
                d_921___mcc_h0_ = source19_.cf21
                d_922___mcc_h1_ = source19_.cf22
                d_923___mcc_h2_ = source19_.cf23
                d_924___mcc_h3_ = source19_.cf24
                d_925_cf24_ = d_924___mcc_h3_
                d_926_cf23_ = d_923___mcc_h2_
                d_927_cf22_ = d_922___mcc_h1_
                d_928_cf21_ = d_921___mcc_h0_
                d_929_v1_: _dafny.Seq
                d_929_v1_ = _dafny.SeqWithoutIsStrInference([p0])
                d_930_v2_: _dafny.Seq
                d_930_v2_ = _dafny.SeqWithoutIsStrInference([d_929_v1_, d_929_v1_, d_929_v1_])
                d_930_v2_ = _dafny.SeqWithoutIsStrInference([(d_929_v1_ if p0 else d_929_v1_) for d_931_i0_ in range(default__.abs(-672))])
                (globalState).f0 = ((d_928_cf21_)[d_925_cf24_] if (d_925_cf24_) in (d_928_cf21_) else default__.safeDivisionInt(d_926_cf23_, d_926_cf23_))
                d_932_v3_: T0
                nw140_ = C1()
                nw140_.ctor__()
                d_932_v3_ = nw140_
                d_932_v3_ = d_932_v3_
                d_933_v4_: _dafny.Map
                d_933_v4_ = _dafny.Map({d_920_v0_: p0})
                d_934_v5_: str
                d_934_v5_ = _dafny.CodePoint('g')
                d_935_v6_: _dafny.Map
                d_935_v6_ = _dafny.Map({d_934_v5_: (self).f33})
                def iife82_():
                    coll44_ = _dafny.Map()
                    compr_44_: _dafny.Seq
                    for compr_44_ in (d_930_v2_).Elements:
                        d_936_v7_: _dafny.Seq = compr_44_
                        if (d_936_v7_) in (d_930_v2_):
                            coll44_[d_936_v7_] = d_927_cf22_
                    return _dafny.Map(coll44_)
                (globalState).f24 = (d_932_v3_).fm6((d_933_v4_).set(d_926_cf23_, d_927_cf22_), (0) - (d_925_cf24_), (((d_935_v6_)[d_934_v5_] if (d_934_v5_) in (d_935_v6_) else (self).f33)) + ((self).f33), iife82_()
                , globalState)
            elif source19_.is_DC12:
                d_937___mcc_h4_ = source19_.cf25
                d_938___mcc_h5_ = source19_.cf26
                d_939___mcc_h6_ = source19_.cf27
                d_940_cf27_ = d_939___mcc_h6_
                d_941_cf26_ = d_938___mcc_h5_
                d_942_cf25_ = d_937___mcc_h4_
                d_943_v8_: str
                d_943_v8_ = _dafny.CodePoint('a')
                d_944_v9_: _dafny.Map
                d_944_v9_ = _dafny.Map({d_943_v8_: _dafny.SeqWithoutIsStrInference([d_943_v8_ for d_945_i1_ in range(default__.abs(489))])})
                d_944_v9_ = d_944_v9_
                d_946_v10_: _dafny.Array
                def lambda32_(d_947_p0_):
                    def lambda33_(d_948_i2_):
                        return (d_948_i2_) + (len(_dafny.SeqWithoutIsStrInference([not(d_947_p0_), d_947_p0_, d_947_p0_])))

                    return lambda33_

                init18_ = lambda32_(p0)
                nw141_ = _dafny.Array(None, 9)
                for i0_18_ in range(nw141_.length(0)):
                    nw141_[i0_18_] = init18_(i0_18_)
                d_946_v10_ = nw141_
                d_949_v11_: _dafny.Seq
                d_949_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bb"))
                index131_ = default__.safeIndex(373, (d_946_v10_).length(0))
                (d_946_v10_)[index131_] = len(d_949_v11_)
                index132_ = default__.safeIndex(373, (d_946_v10_).length(0))
                (d_946_v10_)[index132_] = (d_941_cf26_) - (d_941_cf26_)
                rhs128_ = (d_920_v0_ if False else d_920_v0_)
                rhs129_ = (not (p0) or (d_940_cf27_)) == (d_942_cf25_)
                lhs119_ = globalState
                lhs119_.f5 = rhs128_
                d_942_cf25_ = rhs129_
                d_950_v12_: _dafny.Set
                d_950_v12_ = _dafny.Set({d_942_cf25_})
                d_951_v13_: _dafny.Seq
                d_951_v13_ = _dafny.SeqWithoutIsStrInference([d_950_v12_, d_950_v12_])
                d_952_v14_: _dafny.Set
                d_952_v14_ = _dafny.Set({d_950_v12_, d_950_v12_, d_950_v12_, d_950_v12_, d_950_v12_})
                d_953_v15_: _dafny.Array
                nw142_ = _dafny.Array(None, 26)
                nw142_[int(0)] = d_940_cf27_
                nw142_[int(1)] = d_940_cf27_
                nw142_[int(2)] = d_942_cf25_
                nw142_[int(3)] = not((p0) and (d_940_cf27_))
                nw142_[int(4)] = not (d_942_cf25_) or (p0)
                nw142_[int(5)] = True
                nw142_[int(6)] = d_942_cf25_
                nw142_[int(7)] = d_942_cf25_
                nw142_[int(8)] = default__.fm24(p0, not(d_942_cf25_), d_940_cf27_, 526, globalState)
                nw142_[int(9)] = d_940_cf27_
                nw142_[int(10)] = p0
                nw142_[int(11)] = p0
                nw142_[int(12)] = (369) != (len((self).f33))
                nw142_[int(13)] = d_940_cf27_
                nw142_[int(14)] = (d_951_v13_) != (d_951_v13_)
                nw142_[int(15)] = (len(d_952_v14_)) != ((d_946_v10_)[default__.safeIndex(373, (d_946_v10_).length(0))])
                nw142_[int(16)] = d_940_cf27_
                nw142_[int(17)] = d_940_cf27_
                nw142_[int(18)] = d_942_cf25_
                nw142_[int(19)] = p0
                nw142_[int(20)] = d_942_cf25_
                nw142_[int(21)] = (p0) and (p0)
                nw142_[int(22)] = True
                nw142_[int(23)] = not(d_940_cf27_)
                nw142_[int(24)] = p0
                nw142_[int(25)] = d_942_cf25_
                d_953_v15_ = nw142_
                index133_ = default__.safeIndex(985, (d_953_v15_).length(0))
                (d_953_v15_)[index133_] = d_942_cf25_
                d_954_v16_: _dafny.Set
                d_954_v16_ = _dafny.Set({(d_946_v10_)[default__.safeIndex(373, (d_946_v10_).length(0))]})
                index134_ = default__.safeIndex(985, (d_953_v15_).length(0))
                (d_953_v15_)[index134_] = (d_954_v16_).issubset((d_954_v16_) | (_dafny.Set({d_920_v0_, 365})))
            elif True:
                d_955___mcc_h7_ = source19_.cf20
                d_956_cf20_ = d_955___mcc_h7_
                d_957_v17_: str
                d_957_v17_ = _dafny.CodePoint('b')
                d_958_v18_: _dafny.Map
                d_958_v18_ = _dafny.Map({d_920_v0_: d_957_v17_})
                d_959_v21_: _dafny.Seq
                def iife83_():
                    coll45_ = _dafny.Map()
                    compr_45_: int
                    for compr_45_ in _dafny.IntegerRange(-690, 598):
                        d_960_v19_: int = compr_45_
                        if ((-690) <= (d_960_v19_)) and ((d_960_v19_) < (598)):
                            coll45_[(d_960_v19_) * (d_920_v0_)] = d_957_v17_
                    return _dafny.Map(coll45_)
                def iife84_():
                    coll46_ = _dafny.Map()
                    compr_46_: int
                    for compr_46_ in _dafny.IntegerRange(520, 671):
                        d_961_v20_: int = compr_46_
                        if ((520) <= (d_961_v20_)) and ((d_961_v20_) < (671)):
                            coll46_[(d_961_v20_) * (d_920_v0_)] = d_957_v17_
                    return _dafny.Map(coll46_)
                d_959_v21_ = _dafny.SeqWithoutIsStrInference([d_958_v18_, (d_958_v18_ if False else d_958_v18_), d_958_v18_, iife83_()
                , (iife84_()
                ).set(d_920_v0_, d_957_v17_)])
                d_958_v18_ = (d_959_v21_)[default__.safeIndex(d_920_v0_, len(d_959_v21_))]
                d_962_v22_: C0
                nw143_ = C0()
                nw143_.ctor__()
                d_962_v22_ = nw143_
                d_963_v23_: _dafny.Array
                nw144_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                d_963_v23_ = nw144_
                d_963_v23_ = (d_963_v23_ if p0 else d_963_v23_)
                d_964_v24_: _dafny.MultiSet
                d_964_v24_ = _dafny.MultiSet([d_920_v0_])
                (globalState).f25 = ((d_964_v24_) - (d_964_v24_)).isdisjoint(d_964_v24_)
            (globalState).f9 = p0
            (globalState).f5 = d_920_v0_
            d_965_v25_: _dafny.Array
            def lambda34_(d_966_p0_):
                def lambda35_(d_967_i3_):
                    return not(d_966_p0_)

                return lambda35_

            init19_ = lambda34_(p0)
            nw145_ = _dafny.Array(None, 17)
            for i0_19_ in range(nw145_.length(0)):
                nw145_[i0_19_] = init19_(i0_19_)
            d_965_v25_ = nw145_
            d_968_v26_: _dafny.Set
            d_968_v26_ = _dafny.Set({not(False)})
            index135_ = default__.safeIndex(443, (d_965_v25_).length(0))
            (d_965_v25_)[index135_] = (d_968_v26_).issubset(d_968_v26_)
            d_969_v28_: _dafny.Map
            d_969_v28_ = _dafny.Map({d_920_v0_: p0})
            index136_ = default__.safeIndex(443, (d_965_v25_).length(0))
            def iife85_():
                coll47_ = _dafny.Map()
                compr_47_: int
                for compr_47_ in (d_969_v28_).keys.Elements:
                    d_970_v27_: int = compr_47_
                    if (d_970_v27_) in (d_969_v28_):
                        coll47_[(d_970_v27_) + (d_920_v0_)] = 724
                return _dafny.Map(coll47_)
            (d_965_v25_)[index136_] = (len(((iife85_()
            ).set(d_920_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyynns"))))).set(d_920_v0_, 230))) > (-565)
        elif True:
            d_971_v29_: _dafny.Array
            def lambda36_(d_972_v0_):
                def lambda37_(d_973_i4_):
                    return (d_973_i4_) + (d_972_v0_)

                return lambda37_

            init20_ = lambda36_(d_920_v0_)
            nw146_ = _dafny.Array(None, 3)
            for i0_20_ in range(nw146_.length(0)):
                nw146_[i0_20_] = init20_(i0_20_)
            d_971_v29_ = nw146_
            d_974_v30_: _dafny.Seq
            d_974_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyfmfla"))
            index137_ = default__.safeIndex(466, (d_971_v29_).length(0))
            (d_971_v29_)[index137_] = (0) - (default__.safeModuloInt(len(d_974_v30_), d_920_v0_))
            d_975_v31_: _dafny.Seq
            d_975_v31_ = _dafny.SeqWithoutIsStrInference([(self).f33, (self).f33, (self).f33])
            index138_ = default__.safeIndex(466, (d_971_v29_).length(0))
            (d_971_v29_)[index138_] = default__.safeModuloInt(len((d_975_v31_)[default__.safeIndex(d_920_v0_, len(d_975_v31_))]), d_920_v0_)
            d_976_v32_: _dafny.Array
            def lambda38_(d_977_v0_, d_978_v29_):
                def lambda39_(d_979_i5_):
                    return (d_977_v0_) == ((d_978_v29_)[default__.safeIndex(466, (d_978_v29_).length(0))])

                return lambda39_

            init21_ = lambda38_(d_920_v0_, d_971_v29_)
            nw147_ = _dafny.Array(None, 9)
            for i0_21_ in range(nw147_.length(0)):
                nw147_[i0_21_] = init21_(i0_21_)
            d_976_v32_ = nw147_
            index139_ = default__.safeIndex(326, (d_976_v32_).length(0))
            (d_976_v32_)[index139_] = (614) >= ((d_971_v29_)[default__.safeIndex(466, (d_971_v29_).length(0))])
            index140_ = default__.safeIndex(326, (d_976_v32_).length(0))
            (d_976_v32_)[index140_] = not(p0)
            d_980_v33_: _dafny.Seq
            d_980_v33_ = default__.fm25((d_976_v32_)[default__.safeIndex(326, (d_976_v32_).length(0))], globalState)
            d_981_v34_: _dafny.Seq
            d_981_v34_ = _dafny.SeqWithoutIsStrInference([(d_976_v32_ if (d_976_v32_)[default__.safeIndex(326, (d_976_v32_).length(0))] else d_976_v32_), d_976_v32_, d_976_v32_])
            rhs130_ = (d_981_v34_)[default__.safeIndex(d_920_v0_, len(d_981_v34_))]
            rhs131_ = d_980_v33_
            lhs120_ = globalState
            lhs120_.f1 = rhs130_
            d_980_v33_ = rhs131_
            d_982_v35_: T0
            nw148_ = C1()
            nw148_.ctor__()
            d_982_v35_ = nw148_
            d_983_v36_: _dafny.Set
            d_983_v36_ = _dafny.Set({d_982_v35_, d_982_v35_})
            d_984_v37_: _dafny.Map
            d_984_v37_ = _dafny.Map({682: p0})
            (globalState).f0 = ((d_971_v29_)[default__.safeIndex(466, (d_971_v29_).length(0))]) * ((len(d_983_v36_) if p0 else len(d_984_v37_)))
            d_971_v29_ = d_971_v29_
        (globalState).f9 = p0
        (globalState).f25 = p0
        d_985_v39_: _dafny.MultiSet
        def iife86_():
            coll48_ = _dafny.Map()
            compr_48_: int
            for compr_48_ in _dafny.IntegerRange(28, 26):
                d_986_v38_: int = compr_48_
                if ((28) <= (d_986_v38_)) and ((d_986_v38_) < (26)):
                    coll48_[(d_986_v38_) * (d_920_v0_)] = p0
            return _dafny.Map(coll48_)
        d_985_v39_ = _dafny.MultiSet([len(iife86_()
        ), d_920_v0_, d_920_v0_])
        d_987_v40_: _dafny.Seq
        d_987_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukae"))
        d_988_v41_: _dafny.MultiSet
        d_988_v41_ = _dafny.MultiSet([p0])
        d_989_v42_: _dafny.Set
        d_989_v42_ = _dafny.Set({d_920_v0_, -439})
        d_990_v43_: _dafny.Map
        d_990_v43_ = _dafny.Map({d_920_v0_: p0})
        d_991_v44_: _dafny.Array
        nw149_ = _dafny.Array(None, 17)
        nw149_[int(0)] = d_920_v0_
        nw149_[int(1)] = d_920_v0_
        nw149_[int(2)] = ((d_985_v39_).cardinality) + (d_920_v0_)
        nw149_[int(3)] = default__.safeModuloInt(d_920_v0_, d_920_v0_)
        nw149_[int(4)] = default__.safeModuloInt(d_920_v0_, len(d_987_v40_))
        nw149_[int(5)] = (d_920_v0_) - (d_920_v0_)
        nw149_[int(6)] = d_920_v0_
        nw149_[int(7)] = d_920_v0_
        nw149_[int(8)] = (d_988_v41_).cardinality
        nw149_[int(9)] = len((self).f33)
        nw149_[int(10)] = d_920_v0_
        nw149_[int(11)] = d_920_v0_
        nw149_[int(12)] = 248
        nw149_[int(13)] = (d_985_v39_).cardinality
        nw149_[int(14)] = d_920_v0_
        nw149_[int(15)] = (((self).f33)[default__.safeIndex(len(d_989_v42_), len((self).f33))]) - (d_920_v0_)
        nw149_[int(16)] = default__.safeModuloInt(d_920_v0_, len(d_990_v43_))
        d_991_v44_ = nw149_
        d_991_v44_ = d_991_v44_
        d_992_v45_: _dafny.Map
        d_992_v45_ = _dafny.Map({d_991_v44_: (d_920_v0_) - (d_920_v0_)})
        d_992_v45_ = ((d_992_v45_) | (d_992_v45_)).set(d_991_v44_, (0) - (d_920_v0_))
        r0 = (d_987_v40_ if (d_985_v39_).issubset(d_985_v39_) else d_987_v40_)
        return r0

    def m4(self, p0, p1, globalState):
        r0: int = int(0)
        d_993_v0_: bool
        d_993_v0_ = True
        d_994_v1_: _dafny.MultiSet
        d_994_v1_ = _dafny.MultiSet([p0, p0, p0])
        source20_ = D2_DC9((p0 if d_993_v0_ else p0), ((d_994_v1_).set(p0, default__.abs(len(default__.fm21(p0, p0, d_993_v0_, d_993_v0_, globalState))))).isdisjoint(d_994_v1_), p0, (D0_DC0(p0)).cf0)
        if source20_.is_DC8:
            d_995___mcc_h0_ = source20_.cf13
            d_996___mcc_h1_ = source20_.cf14
            d_997___mcc_h2_ = source20_.cf15
            d_998_cf15_ = d_997___mcc_h2_
            d_999_cf14_ = d_996___mcc_h1_
            d_1000_cf13_ = d_995___mcc_h0_
            d_1001_v2_: _dafny.Array
            def lambda40_(d_1002_i0_):
                return False

            init22_ = lambda40_
            nw150_ = _dafny.Array(None, 4)
            for i0_22_ in range(nw150_.length(0)):
                nw150_[i0_22_] = init22_(i0_22_)
            d_1001_v2_ = nw150_
            d_1003_v3_: _dafny.MultiSet
            d_1003_v3_ = _dafny.MultiSet([d_1001_v2_, d_1001_v2_])
            (globalState).f9 = ((_dafny.MultiSet([d_1001_v2_, d_1001_v2_, d_1001_v2_]) if d_1000_cf13_ else d_1003_v3_)).issubset(d_1003_v3_)
            if d_1000_cf13_:
                d_1004_v4_: C1
                nw151_ = C1()
                nw151_.ctor__()
                d_1004_v4_ = nw151_
                (globalState).f0 = d_998_cf15_
                d_1005_v5_: _dafny.Map
                d_1005_v5_ = _dafny.Map({not(d_1000_cf13_): default__.safeDivisionInt(d_999_cf14_, len((self).f31))})
                d_1005_v5_ = (d_1005_v5_).set(d_1000_cf13_, 66)
                d_993_v0_ = True
                d_1006_v6_: C0
                nw152_ = C0()
                nw152_.ctor__()
                d_1006_v6_ = nw152_
            elif True:
                index141_ = default__.safeIndex(24, (d_1001_v2_).length(0))
                (d_1001_v2_)[index141_] = d_1000_cf13_
                index142_ = default__.safeIndex(24, (d_1001_v2_).length(0))
                (d_1001_v2_)[index142_] = not(not((((d_994_v1_)[-976] if (-976) in (d_994_v1_) else d_999_cf14_)) >= ((p0) * (d_999_cf14_))))
                (globalState).f17 = (0) - (p0)
                d_1007_v7_: _dafny.Seq
                d_1007_v7_ = _dafny.SeqWithoutIsStrInference([d_993_v0_])
                d_1008_v8_: _dafny.MultiSet
                d_1008_v8_ = _dafny.MultiSet([d_1007_v7_, (d_1007_v7_) + (d_1007_v7_), (d_1007_v7_).set(default__.safeIndex(d_998_cf15_, len(d_1007_v7_)), d_1000_cf13_)])
                d_1009_v9_: D3
                d_1009_v9_ = D3_DC11((self).f31, (d_1001_v2_)[default__.safeIndex(24, (d_1001_v2_).length(0))], d_998_cf15_, 478)
                (globalState).f0 = (0) - (((d_1008_v8_)[_dafny.SeqWithoutIsStrInference([d_993_v0_, d_993_v0_, (d_1009_v9_).cf22, d_993_v0_, (d_1001_v2_)[default__.safeIndex(24, (d_1001_v2_).length(0))]])] if (_dafny.SeqWithoutIsStrInference([d_993_v0_, d_993_v0_, (d_1009_v9_).cf22, d_993_v0_, (d_1001_v2_)[default__.safeIndex(24, (d_1001_v2_).length(0))]])) in (d_1008_v8_) else p0))
                d_1010_v10_: _dafny.Set
                d_1010_v10_ = _dafny.Set({d_999_cf14_, (0) - (p0)})
                d_1011_v11_: _dafny.Map
                d_1011_v11_ = _dafny.Map({p0: (d_1001_v2_)[default__.safeIndex(24, (d_1001_v2_).length(0))]})
                d_1012_v13_: str
                d_1012_v13_ = _dafny.CodePoint('b')
                d_1013_v14_: _dafny.Map
                def iife87_():
                    coll49_ = _dafny.Set()
                    compr_49_: int
                    for compr_49_ in (d_1011_v11_).keys.Elements:
                        d_1014_v12_: int = compr_49_
                        if (d_1014_v12_) in (d_1011_v11_):
                            coll49_ = coll49_.union(_dafny.Set([(d_1014_v12_) * (981)]))
                    return _dafny.Set(coll49_)
                d_1013_v14_ = _dafny.Map({(d_1010_v10_).isdisjoint(iife87_()
                ): d_1012_v13_})
                d_1015_v15_: _dafny.Seq
                d_1015_v15_ = _dafny.SeqWithoutIsStrInference([d_1010_v10_, d_1010_v10_])
                d_1016_v16_: _dafny.Seq
                d_1016_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxn"))
                rhs132_ = ((d_1015_v15_)[default__.safeIndex(d_999_cf14_, len(d_1015_v15_))]).issubset(d_1010_v10_)
                rhs133_ = (default__.fm23(d_1016_v16_, globalState)) | (d_1013_v14_)
                d_993_v0_ = rhs132_
                d_1013_v14_ = rhs133_
                d_1000_cf13_ = d_993_v0_
            d_1017_v17_: _dafny.Seq
            d_1017_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "revhj"))
            (globalState).f23 = d_1017_v17_
            d_1018_v18_: str
            d_1018_v18_ = _dafny.CodePoint('n')
            d_1019_v19_: _dafny.Map
            d_1019_v19_ = _dafny.Map({_dafny.Map({d_1018_v18_: d_1018_v18_}): d_1000_cf13_})
            d_1020_v20_: _dafny.Map
            d_1020_v20_ = _dafny.Map({d_1018_v18_: d_1018_v18_})
            d_1021_v21_: D2
            d_1021_v21_ = D2_DC7(_dafny.SeqWithoutIsStrInference([d_1000_cf13_, ((d_1019_v19_)[d_1020_v20_] if (d_1020_v20_) in (d_1019_v19_) else d_1000_cf13_), default__.fm24(d_1000_cf13_, d_1000_cf13_, d_1000_cf13_, p0, globalState)]))
            source21_ = d_1021_v21_
            if source21_.is_DC8:
                d_1022___mcc_h8_ = source21_.cf13
                d_1023___mcc_h9_ = source21_.cf14
                d_1024___mcc_h10_ = source21_.cf15
                d_1025_cf15_ = d_1024___mcc_h10_
                d_1026_cf14_ = d_1023___mcc_h9_
                d_1027_cf13_ = d_1022___mcc_h8_
                (globalState).f0 = p0
                (globalState).f17 = 17
                d_1028_v22_: D2
                d_1028_v22_ = D2_DC8((D1_DC6(d_1027_cf13_, d_1018_v18_, d_1025_cf15_)).cf9, p0, (_dafny.MultiSet((self).f33)).cardinality)
                d_1028_v22_ = d_1028_v22_
                d_1029_v23_: D0
                d_1029_v23_ = D0_DC1(d_1017_v17_, p0, 388)
                d_1030_v24_: _dafny.MultiSet
                d_1030_v24_ = _dafny.MultiSet([d_1029_v23_, d_1029_v23_])
                d_1031_v25_: _dafny.Seq
                d_1031_v25_ = _dafny.SeqWithoutIsStrInference([d_993_v0_])
                d_1032_v26_: _dafny.MultiSet
                d_1032_v26_ = _dafny.MultiSet([(d_1021_v21_).cf12, d_1031_v25_, d_1031_v25_])
                rhs134_ = d_1018_v18_
                rhs135_ = (((0) - (((d_1032_v26_)[d_1031_v25_] if (d_1031_v25_) in (d_1032_v26_) else d_1025_cf15_)) if d_993_v0_ else p0)) <= (259)
                rhs136_ = d_1030_v24_
                rhs137_ = d_1026_cf14_
                rhs138_ = (d_1025_cf15_) * ((d_1025_cf15_ if d_993_v0_ else 340))
                lhs121_ = globalState
                lhs122_ = globalState
                lhs123_ = globalState
                d_1018_v18_ = rhs134_
                lhs121_.f9 = rhs135_
                d_1030_v24_ = rhs136_
                lhs122_.f8 = rhs137_
                lhs123_.f8 = rhs138_
            elif source21_.is_DC9:
                d_1033___mcc_h11_ = source21_.cf16
                d_1034___mcc_h12_ = source21_.cf17
                d_1035___mcc_h13_ = source21_.cf18
                d_1036___mcc_h14_ = source21_.cf19
                d_1037_cf19_ = d_1036___mcc_h14_
                d_1038_cf18_ = d_1035___mcc_h13_
                d_1039_cf17_ = d_1034___mcc_h12_
                d_1040_cf16_ = d_1033___mcc_h11_
                d_1041_v27_: _dafny.Array
                nw153_ = _dafny.Array(_dafny.Map({}), 6)
                d_1041_v27_ = nw153_
                d_1041_v27_ = d_1041_v27_
                d_1042_v28_: _dafny.Map
                d_1042_v28_ = _dafny.Map({p0: d_1000_cf13_})
                d_1042_v28_ = (d_1042_v28_).set(-666, (d_1038_cf18_) > (d_1040_cf16_))
                d_1043_v29_: D7
                d_1043_v29_ = D7_DC17(d_1001_v2_)
                (globalState).f1 = (d_1043_v29_).cf33
                d_1044_v30_: _dafny.Array
                def lambda41_(d_1045_cf19_):
                    def lambda42_(d_1046_i1_):
                        return (d_1046_i1_) * (d_1045_cf19_)

                    return lambda42_

                init23_ = lambda41_(d_1037_cf19_)
                nw154_ = _dafny.Array(None, 13)
                for i0_23_ in range(nw154_.length(0)):
                    nw154_[i0_23_] = init23_(i0_23_)
                d_1044_v30_ = nw154_
                rhs139_ = d_1044_v30_
                rhs140_ = len(((self).f33) + (((self).f33 if d_1000_cf13_ else (self).f33)))
                lhs124_ = globalState
                d_1044_v30_ = rhs139_
                lhs124_.f8 = rhs140_
            elif True:
                d_1047___mcc_h15_ = source21_.cf12
                d_1048_cf12_ = d_1047___mcc_h15_
                d_1049_v31_: C0
                nw155_ = C0()
                nw155_.ctor__()
                d_1049_v31_ = nw155_
                (globalState).f17 = (d_1049_v31_).fm16(p0, 756, -892, (d_1017_v17_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trly"))), globalState)
                d_1050_v32_: _dafny.Set
                d_1050_v32_ = _dafny.Set({p0, 138, p0})
                d_1051_v33_: _dafny.Array
                nw156_ = _dafny.Array(None, 14)
                d_1051_v33_ = nw156_
                d_1052_v35_: _dafny.Set
                d_1052_v35_ = _dafny.Set({d_1018_v18_})
                def iife88_():
                    coll50_ = _dafny.Set()
                    compr_50_: int
                    for compr_50_ in ((self).f33).Elements:
                        d_1053_v34_: int = compr_50_
                        if (d_1053_v34_) in ((self).f33):
                            coll50_ = coll50_.union(_dafny.Set([(d_1053_v34_) * (-900)]))
                    return _dafny.Set(coll50_)
                rhs141_ = (default__.fm26(not(d_993_v0_), not(d_993_v0_), -40, globalState)) | (iife88_()
                )
                rhs142_ = d_1051_v33_
                rhs143_ = (d_1001_v2_ if not(not(d_993_v0_)) else d_1001_v2_)
                rhs144_ = (self).fm7(len(d_1052_v35_), (d_1000_cf13_) and (d_1000_cf13_), globalState)
                lhs125_ = globalState
                lhs126_ = globalState
                d_1050_v32_ = rhs141_
                d_1051_v33_ = rhs142_
                lhs125_.f1 = rhs143_
                lhs126_.f8 = rhs144_
                index143_ = default__.safeIndex(248, (d_1001_v2_).length(0))
                (d_1001_v2_)[index143_] = d_1000_cf13_
                d_1054_v36_: _dafny.Map
                d_1054_v36_ = _dafny.Map({d_1000_cf13_: d_1000_cf13_})
                d_1055_v37_: D3
                d_1055_v37_ = D3_DC10((d_1054_v36_).set(d_993_v0_, d_993_v0_))
                d_1056_v38_: C1
                nw157_ = C1()
                nw157_.ctor__()
                d_1056_v38_ = nw157_
                pat_let_tv9_ = d_1054_v36_
                index144_ = default__.safeIndex(248, (d_1001_v2_).length(0))
                def iife89_(_pat_let19_0):
                    def iife90_(d_1057_dt__update__tmp_h0_):
                        def iife91_(_pat_let20_0):
                            def iife92_(d_1058_dt__update_hcf20_h0_):
                                return D3_DC10(d_1058_dt__update_hcf20_h0_)
                            return iife92_(_pat_let20_0)
                        return iife91_(pat_let_tv9_)
                    return iife90_(_pat_let19_0)
                rhs145_ = d_993_v0_
                rhs146_ = iife89_(d_1055_v37_)
                rhs147_ = d_1056_v38_
                rhs148_ = p0
                rhs149_ = ((d_1017_v17_) + (_dafny.SeqWithoutIsStrInference([d_1018_v18_, _dafny.CodePoint('p'), d_1018_v18_]))) + ((d_1017_v17_) + (d_1017_v17_))
                lhs127_ = d_1001_v2_
                lhs128_ = default__.safeIndex(248, (d_1001_v2_).length(0))
                lhs129_ = globalState
                lhs130_ = globalState
                lhs127_[lhs128_] = rhs145_
                d_1055_v37_ = rhs146_
                d_1056_v38_ = rhs147_
                lhs129_.f5 = rhs148_
                lhs130_.f23 = rhs149_
        elif source20_.is_DC9:
            d_1059___mcc_h3_ = source20_.cf16
            d_1060___mcc_h4_ = source20_.cf17
            d_1061___mcc_h5_ = source20_.cf18
            d_1062___mcc_h6_ = source20_.cf19
            d_1063_cf19_ = d_1062___mcc_h6_
            d_1064_cf18_ = d_1061___mcc_h5_
            d_1065_cf17_ = d_1060___mcc_h4_
            d_1066_cf16_ = d_1059___mcc_h3_
            d_1067_v39_: C1
            nw158_ = C1()
            nw158_.ctor__()
            d_1067_v39_ = nw158_
            d_1068_v40_: D0
            d_1068_v40_ = D0_DC3(d_1066_cf16_, d_1063_cf19_)
            d_1069_v41_: _dafny.Map
            d_1069_v41_ = _dafny.Map({(d_1068_v40_).cf6: d_1065_cf17_})
            d_1070_v42_: _dafny.Map
            d_1070_v42_ = _dafny.Map({len((self).f33): d_1069_v41_})
            d_1071_v43_: D8
            d_1071_v43_ = D8_DC22(d_1069_v41_)
            d_1072_v44_: _dafny.Seq
            d_1072_v44_ = _dafny.SeqWithoutIsStrInference([d_993_v0_])
            d_1073_v46_: _dafny.Array
            nw159_ = _dafny.Array(None, 14)
            nw159_[int(0)] = default__.fm18(((d_994_v1_)[p0] if (p0) in (d_994_v1_) else d_1064_cf18_), globalState)
            nw159_[int(1)] = d_1069_v41_
            nw159_[int(2)] = d_1069_v41_
            nw159_[int(3)] = ((d_1070_v42_)[812] if (812) in (d_1070_v42_) else d_1069_v41_)
            nw159_[int(4)] = d_1069_v41_
            nw159_[int(5)] = (d_1071_v43_).cf41
            nw159_[int(6)] = (d_1069_v41_ if default__.fm24(d_1065_cf17_, d_1065_cf17_, d_993_v0_, (0) - ((0) - ((0) - (len(d_1072_v44_)))), globalState) else _dafny.Map({d_1063_cf19_: True}))
            nw159_[int(7)] = (d_1069_v41_) | ((d_1069_v41_).set(d_1066_cf16_, d_993_v0_))
            nw159_[int(8)] = d_1069_v41_
            nw159_[int(9)] = d_1069_v41_
            def iife93_():
                coll51_ = _dafny.Map()
                compr_51_: int
                for compr_51_ in _dafny.IntegerRange(381, 490):
                    d_1074_v45_: int = compr_51_
                    if ((381) <= (d_1074_v45_)) and ((d_1074_v45_) < (490)):
                        coll51_[default__.safeModuloInt(d_1074_v45_, d_1064_cf18_)] = d_1065_cf17_
                return _dafny.Map(coll51_)
            nw159_[int(10)] = (default__.fm18(len(_dafny.SeqWithoutIsStrInference([d_1065_cf17_])), globalState) if d_993_v0_ else iife93_()
            )
            nw159_[int(11)] = (d_1071_v43_).cf41
            nw159_[int(12)] = d_1069_v41_
            nw159_[int(13)] = d_1069_v41_
            d_1073_v46_ = nw159_
            d_1075_v47_: _dafny.Set
            d_1075_v47_ = _dafny.Set({p0})
            d_1076_v49_: _dafny.Map
            d_1076_v49_ = _dafny.Map({d_1066_cf16_: d_994_v1_})
            d_1077_v52_: _dafny.Seq
            d_1077_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpthx"))
            nw160_ = _dafny.Array(None, 17)
            nw160_[int(0)] = _dafny.Map({len(d_1075_v47_): d_1065_cf17_})
            nw160_[int(1)] = (_dafny.Map({len((self).f33): d_993_v0_})).set(d_1063_cf19_, d_1065_cf17_)
            def iife94_():
                coll52_ = _dafny.Map()
                compr_52_: int
                for compr_52_ in (d_1076_v49_).keys.Elements:
                    d_1078_v48_: int = compr_52_
                    if (d_1078_v48_) in (d_1076_v49_):
                        coll52_[default__.safeModuloInt(d_1078_v48_, d_1064_cf18_)] = d_1065_cf17_
                return _dafny.Map(coll52_)
            nw160_[int(2)] = (iife94_()
            ) | (d_1069_v41_)
            nw160_[int(3)] = (_dafny.Map({d_1063_cf19_: d_993_v0_})) | (d_1069_v41_)
            nw160_[int(4)] = d_1069_v41_
            nw160_[int(5)] = d_1069_v41_
            nw160_[int(6)] = d_1069_v41_
            nw160_[int(7)] = d_1069_v41_
            def iife95_():
                coll53_ = _dafny.Map()
                compr_53_: int
                for compr_53_ in _dafny.IntegerRange(22, 977):
                    d_1079_v50_: int = compr_53_
                    if ((22) <= (d_1079_v50_)) and ((d_1079_v50_) < (977)):
                        coll53_[(d_1079_v50_) * (d_1064_cf18_)] = not(d_1065_cf17_)
                return _dafny.Map(coll53_)
            nw160_[int(8)] = (iife95_()
            ) | (d_1069_v41_)
            def iife96_():
                coll54_ = _dafny.Map()
                compr_54_: int
                for compr_54_ in _dafny.IntegerRange(914, 552):
                    d_1080_v51_: int = compr_54_
                    if ((914) <= (d_1080_v51_)) and ((d_1080_v51_) < (552)):
                        coll54_[default__.safeModuloInt(d_1080_v51_, d_1064_cf18_)] = d_1065_cf17_
                return _dafny.Map(coll54_)
            nw160_[int(9)] = ((iife96_()
            ).set(d_1064_cf18_, d_993_v0_)) | (_dafny.Map({p0: d_1065_cf17_}))
            nw160_[int(10)] = d_1069_v41_
            nw160_[int(11)] = (d_1069_v41_) | (d_1069_v41_)
            nw160_[int(12)] = (d_1069_v41_) | (d_1069_v41_)
            nw160_[int(13)] = (d_1069_v41_) | (default__.fm18((0) - (len((self).f33)), globalState))
            nw160_[int(14)] = d_1069_v41_
            nw160_[int(15)] = d_1069_v41_
            nw160_[int(16)] = _dafny.Map({len(d_1077_v52_): not(d_1065_cf17_)})
            d_1073_v46_ = nw160_
            d_1081_v53_: D6
            d_1081_v53_ = D6_DC16(((self).f33)[default__.safeIndex(d_1064_cf18_, len((self).f33))], d_993_v0_)
            d_1065_cf17_ = (d_1081_v53_).cf32
            rhs150_ = not((d_1065_cf17_) or (d_1065_cf17_))
            rhs151_ = default__.safeDivisionInt(735, len(_dafny.Map({d_1066_cf16_: D2_DC7(d_1072_v44_)})))
            rhs152_ = d_1065_cf17_
            lhs131_ = globalState
            lhs132_ = globalState
            lhs133_ = globalState
            lhs131_.f25 = rhs150_
            lhs132_.f5 = rhs151_
            lhs133_.f25 = rhs152_
        elif True:
            d_1082___mcc_h7_ = source20_.cf12
            d_1083_cf12_ = d_1082___mcc_h7_
            d_1084_v54_: _dafny.Seq
            d_1084_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apacnlo"))
            d_1085_v55_: T0
            nw161_ = C1()
            nw161_.ctor__()
            d_1085_v55_ = nw161_
            d_1086_v56_: _dafny.MultiSet
            d_1086_v56_ = _dafny.MultiSet([d_1085_v55_, d_1085_v55_])
            d_1087_v57_: _dafny.Map
            d_1087_v57_ = _dafny.Map({(p0 if False else len(d_1084_v54_)): (p0 if d_993_v0_ else (d_1086_v56_).cardinality)})
            d_1088_v58_: _dafny.Map
            d_1088_v58_ = _dafny.Map({p0: d_993_v0_})
            d_1087_v57_ = (d_1087_v57_).set(p0, len((d_1088_v58_) | (d_1088_v58_)))
            if (d_993_v0_) and (d_993_v0_):
                (globalState).f17 = p0
                (globalState).f8 = (0) - (p0)
                d_1089_v59_: C1
                nw162_ = C1()
                nw162_.ctor__()
                d_1089_v59_ = nw162_
                d_1090_v60_: _dafny.MultiSet
                d_1090_v60_ = _dafny.MultiSet([d_993_v0_])
                (globalState).f8 = default__.safeModuloInt(p0, ((d_1090_v60_)[d_993_v0_] if (d_993_v0_) in (d_1090_v60_) else p0))
                (globalState).f24 = (949) * ((p0 if d_993_v0_ else p0))
            elif True:
                d_1091_v61_: _dafny.Set
                d_1091_v61_ = _dafny.Set({not(d_993_v0_)})
                d_1092_v62_: _dafny.Array
                nw163_ = _dafny.Array(False, 8)
                d_1092_v62_ = nw163_
                d_1093_v63_: _dafny.Set
                d_1093_v63_ = _dafny.Set({d_1092_v62_, d_1092_v62_, d_1092_v62_})
                d_1094_v64_: _dafny.Map
                d_1094_v64_ = _dafny.Map({d_993_v0_: p0})
                d_1095_v65_: _dafny.Array
                nw164_ = _dafny.Array(None, 27)
                nw164_[int(0)] = d_993_v0_
                nw164_[int(1)] = False
                nw164_[int(2)] = (default__.fm27(d_993_v0_, p0, globalState)).cf9
                nw164_[int(3)] = (True if d_993_v0_ else d_993_v0_)
                nw164_[int(4)] = d_993_v0_
                nw164_[int(5)] = (d_1091_v61_).ispropersubset(_dafny.Set({d_993_v0_}))
                nw164_[int(6)] = d_993_v0_
                nw164_[int(7)] = (p0) != (p0)
                nw164_[int(8)] = d_993_v0_
                nw164_[int(9)] = d_993_v0_
                nw164_[int(10)] = (p0) >= (p0)
                nw164_[int(11)] = d_993_v0_
                nw164_[int(12)] = d_993_v0_
                nw164_[int(13)] = d_993_v0_
                nw164_[int(14)] = d_993_v0_
                nw164_[int(15)] = (d_1093_v63_).isdisjoint(d_1093_v63_)
                nw164_[int(16)] = not(d_993_v0_)
                nw164_[int(17)] = d_993_v0_
                nw164_[int(18)] = d_993_v0_
                nw164_[int(19)] = d_993_v0_
                nw164_[int(20)] = d_993_v0_
                nw164_[int(21)] = d_993_v0_
                nw164_[int(22)] = d_993_v0_
                nw164_[int(23)] = d_993_v0_
                nw164_[int(24)] = d_993_v0_
                nw164_[int(25)] = (d_1083_cf12_)[default__.safeIndex(p0, len(d_1083_cf12_))]
                nw164_[int(26)] = (not(d_993_v0_)) not in (d_1094_v64_)
                d_1095_v65_ = nw164_
                (globalState).f1 = d_1092_v62_
                d_1096_v66_: _dafny.Array
                def lambda43_(d_1097_v1_):
                    def lambda44_(d_1098_i2_):
                        return (d_1097_v1_) | (d_1097_v1_)

                    return lambda44_

                init24_ = lambda43_(d_994_v1_)
                nw165_ = _dafny.Array(None, 3)
                for i0_24_ in range(nw165_.length(0)):
                    nw165_[i0_24_] = init24_(i0_24_)
                d_1096_v66_ = nw165_
                d_1099_v67_: _dafny.Array
                nw166_ = _dafny.Array(D2.default()(), 19)
                d_1099_v67_ = nw166_
                index145_ = default__.safeIndex(221, (d_1099_v67_).length(0))
                (d_1099_v67_)[index145_] = p1
                d_1100_v69_: _dafny.Map
                d_1100_v69_ = _dafny.Map({d_1083_cf12_: p0})
                d_1101_v70_: _dafny.Set
                d_1101_v70_ = _dafny.Set({p0})
                index146_ = default__.safeIndex(221, (d_1099_v67_).length(0))
                def iife97_():
                    coll55_ = _dafny.Map()
                    compr_55_: _dafny.Seq
                    for compr_55_ in (d_1100_v69_).keys.Elements:
                        d_1102_v68_: _dafny.Seq = compr_55_
                        if (d_1102_v68_) in (d_1100_v69_):
                            coll55_[d_1102_v68_] = d_993_v0_
                    return _dafny.Map(coll55_)
                rhs153_ = (self).fm6((d_1088_v58_) | (d_1088_v58_), p0, (self).f33, (iife97_()
                ) | (default__.fm28(globalState)), globalState)
                rhs154_ = d_1096_v66_
                rhs155_ = len(d_1101_v70_)
                rhs156_ = (p1 if d_993_v0_ else p1)
                lhs134_ = globalState
                lhs135_ = globalState
                lhs136_ = d_1099_v67_
                lhs137_ = default__.safeIndex(221, (d_1099_v67_).length(0))
                lhs134_.f8 = rhs153_
                d_1096_v66_ = rhs154_
                lhs135_.f17 = rhs155_
                lhs136_[lhs137_] = rhs156_
                d_1103_v71_: _dafny.Array
                nw167_ = _dafny.Array(int(0), 26)
                d_1103_v71_ = nw167_
                d_1104_v72_: _dafny.Array
                nw168_ = _dafny.Array(None, 8)
                nw168_[int(0)] = d_1103_v71_
                nw168_[int(1)] = d_1103_v71_
                nw168_[int(2)] = d_1103_v71_
                nw168_[int(3)] = d_1103_v71_
                nw168_[int(4)] = d_1103_v71_
                nw168_[int(5)] = d_1103_v71_
                nw168_[int(6)] = d_1103_v71_
                nw168_[int(7)] = d_1103_v71_
                d_1104_v72_ = nw168_
                index147_ = default__.safeIndex(504, (d_1104_v72_).length(0))
                (d_1104_v72_)[index147_] = d_1103_v71_
                index148_ = default__.safeIndex(504, (d_1104_v72_).length(0))
                (d_1104_v72_)[index148_] = d_1103_v71_
                rhs157_ = ((-653) > (p0)) == (d_993_v0_)
                rhs158_ = len(d_1083_cf12_)
                rhs159_ = default__.safeModuloInt(p0, p0)
                lhs138_ = globalState
                lhs139_ = globalState
                d_993_v0_ = rhs157_
                lhs138_.f24 = rhs158_
                lhs139_.f24 = rhs159_
                (globalState).f24 = (0) - ((0) - (p0))
            d_1105_v73_: _dafny.Array
            def lambda45_(d_1106_i3_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "golu"))

            init25_ = lambda45_
            nw169_ = _dafny.Array(None, 3)
            for i0_25_ in range(nw169_.length(0)):
                nw169_[i0_25_] = init25_(i0_25_)
            d_1105_v73_ = nw169_
            index149_ = default__.safeIndex(752, (d_1105_v73_).length(0))
            (d_1105_v73_)[index149_] = (d_1084_v54_) + (d_1084_v54_)
            d_1107_v74_: str
            d_1107_v74_ = _dafny.CodePoint('m')
            index150_ = default__.safeIndex(752, (d_1105_v73_).length(0))
            rhs160_ = d_1084_v54_
            rhs161_ = ((d_1088_v58_).set(p0, d_993_v0_)) | (d_1088_v58_)
            rhs162_ = (self).fm7(p0, (d_993_v0_ if d_993_v0_ else d_993_v0_), globalState)
            rhs163_ = (((d_1084_v54_).set(default__.safeIndex(p0, len(d_1084_v54_)), d_1107_v74_)) + (_dafny.SeqWithoutIsStrInference([d_1107_v74_ for d_1108_i4_ in range(default__.abs(603))]))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mogo")))
            lhs140_ = globalState
            lhs141_ = globalState
            lhs142_ = d_1105_v73_
            lhs143_ = default__.safeIndex(752, (d_1105_v73_).length(0))
            lhs140_.f13 = rhs160_
            d_1088_v58_ = rhs161_
            lhs141_.f5 = rhs162_
            lhs142_[lhs143_] = rhs163_
            if d_993_v0_:
                d_1109_v75_: C0
                nw170_ = C0()
                nw170_.ctor__()
                d_1109_v75_ = nw170_
                d_1110_v76_: D7
                d_1110_v76_ = D7_DC19(len(d_1083_cf12_), default__.fm24((d_1083_cf12_)[default__.safeIndex(len(d_1087_v57_), len(d_1083_cf12_))], d_993_v0_, d_993_v0_, p0, globalState))
                d_1111_v77_: _dafny.Map
                d_1111_v77_ = _dafny.Map({d_993_v0_: d_993_v0_})
                d_1112_v78_: _dafny.Set
                d_1112_v78_ = _dafny.Set({d_993_v0_, ((d_1111_v77_)[d_993_v0_] if (d_993_v0_) in (d_1111_v77_) else d_993_v0_)})
                d_1113_v79_: _dafny.Set
                d_1113_v79_ = _dafny.Set({not((_dafny.Set({d_993_v0_, not((d_1110_v76_).cf37)})).ispropersubset(d_1112_v78_))})
                d_1114_v80_: _dafny.MultiSet
                d_1114_v80_ = _dafny.MultiSet([False, d_993_v0_, d_993_v0_, d_993_v0_])
                d_1115_v81_: D2
                d_1115_v81_ = D2_DC9(p0, d_993_v0_, ((0) - ((0) - (p0))) * (482), (d_1114_v80_).cardinality)
                rhs164_ = d_1112_v78_
                rhs165_ = d_1115_v81_
                d_1112_v78_ = rhs164_
                d_1115_v81_ = rhs165_
                d_1116_v82_: _dafny.Array
                nw171_ = _dafny.Array(int(0), 8)
                d_1116_v82_ = nw171_
                d_1117_v83_: _dafny.MultiSet
                d_1117_v83_ = _dafny.MultiSet([d_1116_v82_, d_1116_v82_])
                rhs166_ = d_1117_v83_
                rhs167_ = d_993_v0_
                lhs144_ = globalState
                d_1117_v83_ = rhs166_
                lhs144_.f25 = rhs167_
                d_1118_v84_: C0
                nw172_ = C0()
                nw172_.ctor__()
                d_1118_v84_ = nw172_
                (globalState).f23 = (d_1105_v73_)[default__.safeIndex(752, (d_1105_v73_).length(0))]
            elif True:
                (globalState).f9 = d_993_v0_
                d_1119_v85_: _dafny.MultiSet
                d_1119_v85_ = _dafny.MultiSet([not(d_993_v0_)])
                d_1120_v86_: _dafny.Map
                d_1120_v86_ = _dafny.Map({d_993_v0_: d_1119_v85_})
                d_1121_v87_: _dafny.Map
                d_1121_v87_ = _dafny.Map({_dafny.CodePoint('x'): False})
                d_1120_v86_ = (d_1120_v86_).set(((d_1121_v87_)[d_1107_v74_] if (d_1107_v74_) in (d_1121_v87_) else default__.fm24(d_993_v0_, d_993_v0_, False, p0, globalState)), d_1119_v85_)
                (globalState).f9 = d_993_v0_
                d_1088_v58_ = d_1088_v58_
                d_1122_v88_: _dafny.Map
                d_1122_v88_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_993_v0_]): (d_1088_v58_).set(p0, d_993_v0_)})
                d_1122_v88_ = (d_1122_v88_).set(d_1083_cf12_, (d_1088_v58_) | (d_1088_v58_))
        d_1123_v89_: D0
        d_1123_v89_ = D0_DC2(d_993_v0_)
        d_1124_v90_: D0
        d_1124_v90_ = D0_DC4(d_1123_v89_)
        d_1125_v91_: _dafny.Seq
        d_1125_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uartbboah"))
        d_1126_v92_: str
        d_1126_v92_ = _dafny.CodePoint('g')
        def iife98_():
            coll56_ = _dafny.Map()
            compr_56_: int
            for compr_56_ in (_dafny.MultiSet((self).f33)).Elements:
                d_1127_v93_: int = compr_56_
                if (d_1127_v93_) in (_dafny.MultiSet((self).f33)):
                    coll56_[(d_1127_v93_) * (284)] = p0
            return _dafny.Map(coll56_)
        d_1124_v90_ = default__.fm29((d_1125_v91_).set(default__.safeIndex(-824, len(d_1125_v91_)), d_1126_v92_), _dafny.CodePoint('w'), iife98_()
        , globalState)
        d_1128_v94_: _dafny.Map
        d_1128_v94_ = _dafny.Map({d_1126_v92_: len((self).f31)})
        d_1129_v95_: _dafny.MultiSet
        d_1129_v95_ = _dafny.MultiSet([d_1128_v94_, _dafny.Map({d_1126_v92_: (0) - (p0)})])
        d_1130_v96_: _dafny.Map
        d_1130_v96_ = _dafny.Map({len(d_1125_v91_): d_1129_v95_})
        d_1131_v97_: _dafny.Map
        d_1131_v97_ = _dafny.Map({not(d_993_v0_): 418})
        d_1132_v98_: _dafny.Map
        d_1132_v98_ = _dafny.Map({(0) - ((p0 if not(True) else p0)): ((((d_1130_v96_)[p0] if (p0) in (d_1130_v96_) else d_1129_v95_)).set(d_1128_v94_, default__.abs(len(d_1131_v97_)))).issubset(d_1129_v95_)})
        d_1133_v99_: _dafny.Array
        nw173_ = _dafny.Array(False, 10)
        d_1133_v99_ = nw173_
        d_1134_v100_: D7
        d_1134_v100_ = D7_DC17(d_1133_v99_)
        d_1135_v101_: _dafny.Set
        d_1135_v101_ = _dafny.Set({d_1134_v100_})
        d_1132_v98_ = (d_1132_v98_).set(((self).f33)[default__.safeIndex(len(d_1135_v101_), len((self).f33))], not (d_993_v0_) or (d_993_v0_))
        d_1136_v102_: _dafny.Map
        d_1136_v102_ = _dafny.Map({d_993_v0_: not(d_993_v0_)})
        d_1137_v103_: _dafny.Set
        d_1137_v103_ = _dafny.Set({p0})
        d_1138_v104_: _dafny.Map
        d_1138_v104_ = _dafny.Map({d_1136_v102_: d_1137_v103_})
        d_1139_v106_: _dafny.Set
        d_1139_v106_ = _dafny.Set({d_1136_v102_, d_1136_v102_})
        d_1140_v108_: _dafny.Array
        nw174_ = _dafny.Array(None, 23)
        nw174_[int(0)] = (d_1138_v104_).set(d_1136_v102_, d_1137_v103_)
        nw174_[int(1)] = _dafny.Map({d_1136_v102_: d_1137_v103_})
        nw174_[int(2)] = d_1138_v104_
        nw174_[int(3)] = d_1138_v104_
        nw174_[int(4)] = d_1138_v104_
        nw174_[int(5)] = d_1138_v104_
        nw174_[int(6)] = d_1138_v104_
        nw174_[int(7)] = (d_1138_v104_).set(d_1136_v102_, d_1137_v103_)
        nw174_[int(8)] = (d_1138_v104_) | (d_1138_v104_)
        nw174_[int(9)] = d_1138_v104_
        nw174_[int(10)] = d_1138_v104_
        nw174_[int(11)] = d_1138_v104_
        nw174_[int(12)] = default__.fm30(globalState)
        nw174_[int(13)] = _dafny.Map({d_1136_v102_: d_1137_v103_})
        nw174_[int(14)] = d_1138_v104_
        nw174_[int(15)] = d_1138_v104_
        def iife99_():
            coll57_ = _dafny.Map()
            compr_57_: _dafny.Map
            for compr_57_ in (d_1139_v106_).Elements:
                d_1141_v105_: _dafny.Map = compr_57_
                if (d_1141_v105_) in (d_1139_v106_):
                    coll57_[d_1141_v105_] = d_1137_v103_
            return _dafny.Map(coll57_)
        nw174_[int(16)] = iife99_()
        
        nw174_[int(17)] = d_1138_v104_
        nw174_[int(18)] = d_1138_v104_
        def iife100_():
            coll58_ = _dafny.Map()
            compr_58_: _dafny.Map
            for compr_58_ in (d_1139_v106_).Elements:
                d_1142_v107_: _dafny.Map = compr_58_
                if (d_1142_v107_) in (d_1139_v106_):
                    coll58_[d_1142_v107_] = d_1137_v103_
            return _dafny.Map(coll58_)
        nw174_[int(19)] = (iife100_()
        ) | (d_1138_v104_)
        nw174_[int(20)] = d_1138_v104_
        nw174_[int(21)] = d_1138_v104_
        nw174_[int(22)] = d_1138_v104_
        d_1140_v108_ = nw174_
        d_1143_v111_: _dafny.Seq
        d_1143_v111_ = _dafny.SeqWithoutIsStrInference([d_1136_v102_, _dafny.Map({d_993_v0_: d_993_v0_})])
        index151_ = default__.safeIndex(769, (d_1140_v108_).length(0))
        def iife101_():
            def iife103_():
                coll61_ = _dafny.Map()
                compr_61_: _dafny.Map
                for compr_61_ in (d_1143_v111_).Elements:
                    d_1144_v110_: _dafny.Map = compr_61_
                    if (d_1144_v110_) in (d_1143_v111_):
                        coll61_[d_1144_v110_] = p0
                return _dafny.Map(coll61_)
            coll59_ = _dafny.Map()
            def iife102_():
                coll60_ = _dafny.Map()
                compr_60_: _dafny.Map
                for compr_60_ in (d_1143_v111_).Elements:
                    d_1144_v110_: _dafny.Map = compr_60_
                    if (d_1144_v110_) in (d_1143_v111_):
                        coll60_[d_1144_v110_] = p0
                return _dafny.Map(coll60_)
            compr_59_: _dafny.Map
            for compr_59_ in (iife102_()
            ).keys.Elements:
                d_1145_v109_: _dafny.Map = compr_59_
                if (d_1145_v109_) in (iife103_()
                ):
                    coll59_[d_1145_v109_] = _dafny.Set({p0})
            return _dafny.Map(coll59_)
        (d_1140_v108_)[index151_] = (d_1138_v104_ if d_993_v0_ else iife101_()
        )
        d_1146_v112_: _dafny.Seq
        d_1146_v112_ = _dafny.SeqWithoutIsStrInference([default__.fm31(False, False, d_993_v0_, d_993_v0_, globalState)])
        d_1147_v113_: D8
        d_1147_v113_ = D8_DC24(len(_dafny.SeqWithoutIsStrInference([D0_DC2(d_993_v0_) for d_1148_i5_ in range(default__.abs(356))])), p0, len(d_1146_v112_))
        pat_let_tv10_ = d_1138_v104_
        pat_let_tv11_ = d_1136_v102_
        pat_let_tv12_ = d_1136_v102_
        pat_let_tv13_ = d_1137_v103_
        pat_let_tv14_ = d_1143_v111_
        pat_let_tv15_ = d_1143_v111_
        pat_let_tv16_ = p0
        pat_let_tv17_ = d_1136_v102_
        pat_let_tv18_ = d_1137_v103_
        pat_let_tv19_ = d_1136_v102_
        pat_let_tv20_ = p0
        pat_let_tv21_ = d_1138_v104_
        index152_ = default__.safeIndex(769, (d_1140_v108_).length(0))
        def lambda46_(source22_):
            if source22_.is_DC23:
                d_1149___mcc_h16_ = source22_.cf42
                d_1150___mcc_h17_ = source22_.cf43
                d_1151___mcc_h18_ = source22_.cf44
                d_1152___mcc_h19_ = source22_.cf45
                d_1153_cf45_ = d_1152___mcc_h19_
                d_1154_cf44_ = d_1151___mcc_h18_
                d_1155_cf43_ = d_1150___mcc_h17_
                d_1156_cf42_ = d_1149___mcc_h16_
                return pat_let_tv10_
            elif source22_.is_DC24:
                d_1157___mcc_h20_ = source22_.cf46
                d_1158___mcc_h21_ = source22_.cf47
                d_1159___mcc_h22_ = source22_.cf48
                d_1160_cf48_ = d_1159___mcc_h22_
                d_1161_cf47_ = d_1158___mcc_h21_
                d_1162_cf46_ = d_1157___mcc_h20_
                def iife104_():
                    coll62_ = _dafny.Map()
                    compr_62_: _dafny.Map
                    for compr_62_ in (_dafny.Map({pat_let_tv11_: d_1161_cf47_})).keys.Elements:
                        d_1163_v114_: _dafny.Map = compr_62_
                        if (d_1163_v114_) in (_dafny.Map({pat_let_tv12_: d_1161_cf47_})):
                            coll62_[d_1163_v114_] = pat_let_tv13_
                    return _dafny.Map(coll62_)
                return iife104_()
                
            elif source22_.is_DC25:
                d_1164___mcc_h23_ = source22_.cf49
                d_1165___mcc_h24_ = source22_.cf50
                d_1166___mcc_h25_ = source22_.cf51
                d_1167___mcc_h26_ = source22_.cf52
                d_1168___mcc_h27_ = source22_.cf53
                d_1169_cf53_ = d_1168___mcc_h27_
                d_1170_cf52_ = d_1167___mcc_h26_
                d_1171_cf51_ = d_1166___mcc_h25_
                d_1172_cf50_ = d_1165___mcc_h24_
                d_1173_cf49_ = d_1164___mcc_h23_
                def iife105_():
                    coll63_ = _dafny.Map()
                    compr_63_: _dafny.Map
                    for compr_63_ in (pat_let_tv14_).Elements:
                        d_1174_v115_: _dafny.Map = compr_63_
                        if (d_1174_v115_) in (pat_let_tv15_):
                            coll63_[d_1174_v115_] = _dafny.Set({pat_let_tv16_})
                    return _dafny.Map(coll63_)
                return iife105_()
                
            elif source22_.is_DC22:
                d_1175___mcc_h28_ = source22_.cf41
                d_1176_cf41_ = d_1175___mcc_h28_
                return (_dafny.Map({pat_let_tv17_: pat_let_tv18_})).set(pat_let_tv19_, _dafny.Set({pat_let_tv20_}))
            elif True:
                d_1177___mcc_h29_ = source22_.cf54
                d_1178_cf54_ = d_1177___mcc_h29_
                return pat_let_tv21_

        rhs168_ = (0) - (p0)
        rhs169_ = lambda46_(d_1147_v113_)
        lhs145_ = globalState
        lhs146_ = d_1140_v108_
        lhs147_ = default__.safeIndex(769, (d_1140_v108_).length(0))
        lhs145_.f24 = rhs168_
        lhs146_[lhs147_] = rhs169_
        d_1179_v116_: _dafny.Array
        def lambda47_(d_1180_p0_):
            def lambda48_(d_1181_i7_):
                return default__.safeDivisionInt(d_1181_i7_, d_1180_p0_)

            return lambda48_

        init26_ = lambda47_(p0)
        nw175_ = _dafny.Array(None, 16)
        for i0_26_ in range(nw175_.length(0)):
            nw175_[i0_26_] = init26_(i0_26_)
        d_1179_v116_ = nw175_
        d_1182_v117_: _dafny.Map
        d_1182_v117_ = _dafny.Map({d_1179_v116_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kq"))})
        hi4_ = p0
        for d_1183_i6_ in range(len(d_1182_v117_), hi4_):
            (globalState).f25 = not((p0) < ((p0) - (942)))
            index153_ = default__.safeIndex(992, (d_1179_v116_).length(0))
            (d_1179_v116_)[index153_] = d_1183_i6_
            index154_ = default__.safeIndex(992, (d_1179_v116_).length(0))
            (d_1179_v116_)[index154_] = p0
            d_1184_v118_: _dafny.Seq
            d_1184_v118_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, d_993_v0_, d_993_v0_])])
            d_1185_v119_: _dafny.Seq
            d_1185_v119_ = _dafny.SeqWithoutIsStrInference([d_993_v0_])
            d_1186_v120_: _dafny.Seq
            d_1186_v120_ = _dafny.SeqWithoutIsStrInference([d_1184_v118_])
            d_1187_v121_: _dafny.Seq
            d_1187_v121_ = _dafny.SeqWithoutIsStrInference([d_1184_v118_, _dafny.SeqWithoutIsStrInference([d_1185_v119_]), (d_1186_v120_)[default__.safeIndex(d_1183_i6_, len(d_1186_v120_))], _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_993_v0_, default__.fm24(d_993_v0_, d_993_v0_, d_993_v0_, p0, globalState), d_993_v0_, d_993_v0_])])])
            index155_ = default__.safeIndex(992, (d_1179_v116_).length(0))
            rhs170_ = (((d_1187_v121_)[default__.safeIndex(d_1183_i6_, len(d_1187_v121_))]) + (_dafny.SeqWithoutIsStrInference([d_1185_v119_]))) <= ((d_1184_v118_) + (d_1184_v118_))
            rhs171_ = (p0) - (len(d_1185_v119_))
            rhs172_ = d_1125_v91_
            lhs148_ = globalState
            lhs149_ = d_1179_v116_
            lhs150_ = default__.safeIndex(992, (d_1179_v116_).length(0))
            lhs151_ = globalState
            lhs148_.f25 = rhs170_
            lhs149_[lhs150_] = rhs171_
            lhs151_.f13 = rhs172_
            d_1188_v122_: _dafny.Map
            d_1188_v122_ = _dafny.Map({d_994_v1_: d_1185_v119_})
            d_1189_v123_: _dafny.Map
            d_1189_v123_ = _dafny.Map({d_993_v0_: _dafny.MultiSet([(d_1179_v116_)[default__.safeIndex(992, (d_1179_v116_).length(0))], (d_1179_v116_)[default__.safeIndex(992, (d_1179_v116_).length(0))]])})
            source23_ = default__.fm32(d_993_v0_, (d_1188_v122_) | (d_1188_v122_), d_1189_v123_, d_1126_v92_, globalState)
            if source23_.is_DC8:
                d_1190___mcc_h30_ = source23_.cf13
                d_1191___mcc_h31_ = source23_.cf14
                d_1192___mcc_h32_ = source23_.cf15
                d_1193_cf15_ = d_1192___mcc_h32_
                d_1194_cf14_ = d_1191___mcc_h31_
                d_1195_cf13_ = d_1190___mcc_h30_
                d_1196_v124_: _dafny.Map
                d_1196_v124_ = _dafny.Map({d_1131_v97_: True})
                d_1197_v125_: _dafny.Map
                d_1197_v125_ = _dafny.Map({d_1196_v124_: (d_1179_v116_)[default__.safeIndex(992, (d_1179_v116_).length(0))]})
                d_1197_v125_ = (d_1197_v125_).set(d_1196_v124_, p0)
                def iife106_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in (d_994_v1_).Elements:
                        d_1198_v126_: int = compr_64_
                        if (d_1198_v126_) in (d_994_v1_):
                            coll64_[default__.safeDivisionInt(d_1198_v126_, d_1194_cf14_)] = d_1195_cf13_
                    return _dafny.Map(coll64_)
                (globalState).f5 = len((_dafny.SeqWithoutIsStrInference([d_1195_cf13_, d_993_v0_, True]) if (d_1193_cf15_) < ((self).fm6(iife106_()
                , d_1183_i6_, (self).f33, _dafny.Map({d_1185_v119_: d_1195_cf13_}), globalState)) else d_1185_v119_))
                (globalState).f25 = d_1195_cf13_
                index156_ = default__.safeIndex(225, (d_1133_v99_).length(0))
                (d_1133_v99_)[index156_] = d_993_v0_
                index157_ = default__.safeIndex(225, (d_1133_v99_).length(0))
                (d_1133_v99_)[index157_] = d_1195_cf13_
            elif source23_.is_DC9:
                d_1199___mcc_h33_ = source23_.cf16
                d_1200___mcc_h34_ = source23_.cf17
                d_1201___mcc_h35_ = source23_.cf18
                d_1202___mcc_h36_ = source23_.cf19
                d_1203_cf19_ = d_1202___mcc_h36_
                d_1204_cf18_ = d_1201___mcc_h35_
                d_1205_cf17_ = d_1200___mcc_h34_
                d_1206_cf16_ = d_1199___mcc_h33_
                d_1207_v127_: _dafny.Map
                d_1207_v127_ = _dafny.Map({d_1133_v99_: d_993_v0_})
                d_1208_v128_: bool
                d_1209_v129_: int
                d_1210_v130_: bool
                out31_: bool
                out32_: int
                out33_: bool
                out31_, out32_, out33_ = (self).m6(d_1205_cf17_, (_dafny.Map({d_1133_v99_: d_993_v0_})) | (d_1207_v127_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xk")), d_993_v0_, globalState)
                d_1208_v128_ = out31_
                d_1209_v129_ = out32_
                d_1210_v130_ = out33_
                (globalState).f5 = (d_1179_v116_)[default__.safeIndex(992, (d_1179_v116_).length(0))]
                d_1211_v131_: _dafny.Set
                d_1211_v131_ = _dafny.Set({False})
                (globalState).f9 = (default__.fm24(d_1210_v130_, d_1205_cf17_, d_1210_v130_, len(d_1211_v131_), globalState)) and ((not(d_993_v0_)) == (d_993_v0_))
                d_1212_v132_: _dafny.Set
                d_1212_v132_ = _dafny.Set({d_1133_v99_, d_1133_v99_, d_1133_v99_, d_1133_v99_})
                d_1213_v133_: _dafny.Map
                d_1213_v133_ = _dafny.Map({(d_1212_v132_) - (d_1212_v132_): d_1125_v91_})
                d_1213_v133_ = (d_1213_v133_).set(d_1212_v132_, _dafny.SeqWithoutIsStrInference([d_1126_v92_ for d_1214_i8_ in range(default__.abs(511))]))
            elif True:
                d_1215___mcc_h37_ = source23_.cf12
                d_1216_cf12_ = d_1215___mcc_h37_
                d_1217_v134_: _dafny.MultiSet
                d_1217_v134_ = _dafny.MultiSet([d_993_v0_, not(d_993_v0_)])
                (globalState).f0 = ((d_1217_v134_)[(d_1126_v92_) not in (_dafny.SeqWithoutIsStrInference([d_1126_v92_ for d_1218_i9_ in range(default__.abs(233))]))] if ((d_1126_v92_) not in (_dafny.SeqWithoutIsStrInference([d_1126_v92_ for d_1219_i9_ in range(default__.abs(233))]))) in (d_1217_v134_) else len(_dafny.Map({d_993_v0_: p0})))
                d_1220_v135_: _dafny.Array
                def lambda49_(d_1221_v103_):
                    def lambda50_(d_1222_i10_):
                        return d_1221_v103_

                    return lambda50_

                init27_ = lambda49_(d_1137_v103_)
                nw176_ = _dafny.Array(None, 16)
                for i0_27_ in range(nw176_.length(0)):
                    nw176_[i0_27_] = init27_(i0_27_)
                d_1220_v135_ = nw176_
                index158_ = default__.safeIndex(71, (d_1220_v135_).length(0))
                def iife107_():
                    coll65_ = _dafny.Set()
                    compr_65_: int
                    for compr_65_ in (d_1132_v98_).keys.Elements:
                        d_1223_v136_: int = compr_65_
                        if (d_1223_v136_) in (d_1132_v98_):
                            coll65_ = coll65_.union(_dafny.Set([(d_1223_v136_) - (495)]))
                    return _dafny.Set(coll65_)
                (d_1220_v135_)[index158_] = iife107_()
                
                d_1224_v138_: _dafny.Map
                def iife108_():
                    coll66_ = _dafny.Set()
                    compr_66_: int
                    for compr_66_ in _dafny.IntegerRange(18, 422):
                        d_1225_v137_: int = compr_66_
                        if ((18) <= (d_1225_v137_)) and ((d_1225_v137_) < (422)):
                            coll66_ = coll66_.union(_dafny.Set([(d_1225_v137_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlhoqiw"))))]))
                    return _dafny.Set(coll66_)
                d_1224_v138_ = _dafny.Map({d_993_v0_: iife108_()
                })
                d_1226_v139_: _dafny.Map
                d_1226_v139_ = _dafny.Map({d_993_v0_: d_1126_v92_})
                index159_ = default__.safeIndex(71, (d_1220_v135_).length(0))
                rhs173_ = not(d_993_v0_)
                rhs174_ = p0
                rhs175_ = not((p0) > (d_1183_i6_))
                rhs176_ = ((d_1224_v138_)[d_993_v0_] if (d_993_v0_) in (d_1224_v138_) else d_1137_v103_)
                rhs177_ = default__.fm24((_dafny.SeqWithoutIsStrInference([d_993_v0_])) <= (d_1185_v119_), ((d_1179_v116_)[default__.safeIndex(992, (d_1179_v116_).length(0))]) >= (len(d_1136_v102_)), d_993_v0_, ((d_1179_v116_)[default__.safeIndex(992, (d_1179_v116_).length(0))] if d_993_v0_ else len(d_1226_v139_)), globalState)
                lhs152_ = globalState
                lhs153_ = globalState
                lhs154_ = d_1220_v135_
                lhs155_ = default__.safeIndex(71, (d_1220_v135_).length(0))
                lhs156_ = globalState
                lhs152_.f25 = rhs173_
                lhs153_.f8 = rhs174_
                d_993_v0_ = rhs175_
                lhs154_[lhs155_] = rhs176_
                lhs156_.f25 = rhs177_
                d_1227_v140_: D1
                d_1227_v140_ = D1_DC6(d_993_v0_, d_1126_v92_, (d_1179_v116_)[default__.safeIndex(992, (d_1179_v116_).length(0))])
                (globalState).f9 = default__.fm24((d_993_v0_) or ((d_1227_v140_).cf9), (d_1183_i6_) == ((d_1179_v116_)[default__.safeIndex(992, (d_1179_v116_).length(0))]), False, ((d_1131_v97_)[d_993_v0_] if (d_993_v0_) in (d_1131_v97_) else 669), globalState)
                d_1228_v141_: _dafny.Seq
                d_1228_v141_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wswtmk"))])
                d_1228_v141_ = _dafny.SeqWithoutIsStrInference([(D0_DC1(d_1125_v91_, p0, (d_1179_v116_)[default__.safeIndex(992, (d_1179_v116_).length(0))])).cf1 for d_1229_i11_ in range(default__.abs(130))])
        source24_ = p1
        if source24_.is_DC8:
            d_1230___mcc_h38_ = source24_.cf13
            d_1231___mcc_h39_ = source24_.cf14
            d_1232___mcc_h40_ = source24_.cf15
            d_1233_cf15_ = d_1232___mcc_h40_
            d_1234_cf14_ = d_1231___mcc_h39_
            d_1235_cf13_ = d_1230___mcc_h38_
            d_1236_v142_: _dafny.Array
            def lambda51_(d_1237_v92_):
                def lambda52_(d_1238_i12_):
                    return d_1237_v92_

                return lambda52_

            init28_ = lambda51_(d_1126_v92_)
            nw177_ = _dafny.Array(None, 4)
            for i0_28_ in range(nw177_.length(0)):
                nw177_[i0_28_] = init28_(i0_28_)
            d_1236_v142_ = nw177_
            index160_ = default__.safeIndex(2, (d_1236_v142_).length(0))
            (d_1236_v142_)[index160_] = d_1126_v92_
            index161_ = default__.safeIndex(2, (d_1236_v142_).length(0))
            (d_1236_v142_)[index161_] = d_1126_v92_
            d_1133_v99_ = d_1133_v99_
            (globalState).f17 = ((0) - (d_1233_cf15_)) * (d_1234_cf14_)
            d_1239_v143_: _dafny.Set
            d_1239_v143_ = _dafny.Set({d_993_v0_})
            d_1240_v144_: D1
            d_1240_v144_ = D1_DC5(d_1239_v143_)
            source25_ = d_1240_v144_
            if source25_.is_DC6:
                d_1241___mcc_h46_ = source25_.cf9
                d_1242___mcc_h47_ = source25_.cf10
                d_1243___mcc_h48_ = source25_.cf11
                d_1244_cf11_ = d_1243___mcc_h48_
                d_1245_cf10_ = d_1242___mcc_h47_
                d_1246_cf9_ = d_1241___mcc_h46_
                (globalState).f25 = (((_dafny.MultiSet([d_1246_cf9_])).set(d_1235_cf13_, default__.abs(d_1244_cf11_))) != (_dafny.MultiSet([d_993_v0_])) if (d_1239_v143_).issubset(d_1239_v143_) else d_1246_cf9_)
                (globalState).f17 = default__.safeDivisionInt(d_1244_cf11_, 564)
                index162_ = default__.safeIndex(810, (d_1179_v116_).length(0))
                (d_1179_v116_)[index162_] = d_1244_cf11_
                d_1247_v145_: _dafny.Array
                def lambda53_(d_1248_v143_):
                    def lambda54_(d_1249_i13_):
                        return d_1248_v143_

                    return lambda54_

                init29_ = lambda53_(d_1239_v143_)
                nw178_ = _dafny.Array(None, 4)
                for i0_29_ in range(nw178_.length(0)):
                    nw178_[i0_29_] = init29_(i0_29_)
                d_1247_v145_ = nw178_
                index163_ = default__.safeIndex(474, (d_1247_v145_).length(0))
                (d_1247_v145_)[index163_] = _dafny.Set({d_1235_cf13_, default__.fm24(d_1235_cf13_, d_1246_cf9_, not(d_993_v0_), (0) - (d_1234_cf14_), globalState)})
                index164_ = default__.safeIndex(441, (d_1179_v116_).length(0))
                (d_1179_v116_)[index164_] = (d_994_v1_).cardinality
                d_1250_v146_: _dafny.Seq
                d_1250_v146_ = _dafny.SeqWithoutIsStrInference([d_1132_v98_])
                d_1251_v147_: _dafny.Map
                d_1251_v147_ = _dafny.Map({(0) - (len(d_1137_v103_)): d_1239_v143_})
                d_1252_v148_: _dafny.Seq
                d_1252_v148_ = _dafny.SeqWithoutIsStrInference([(d_1250_v146_)[default__.safeIndex(len(d_1251_v147_), len(d_1250_v146_))]])
                d_1253_v149_: _dafny.Map
                d_1253_v149_ = _dafny.Map({not(d_1235_cf13_): _dafny.Map({_dafny.Set({p0}): (0) - (d_1234_cf14_)})})
                d_1254_v150_: _dafny.Map
                d_1254_v150_ = _dafny.Map({d_1137_v103_: d_1233_cf15_})
                d_1255_v151_: _dafny.Seq
                d_1255_v151_ = _dafny.SeqWithoutIsStrInference([d_1246_cf9_, d_1235_cf13_])
                d_1256_v152_: _dafny.Map
                d_1256_v152_ = _dafny.Map({d_1255_v151_: d_993_v0_})
                index165_ = default__.safeIndex(810, (d_1179_v116_).length(0))
                index166_ = default__.safeIndex(474, (d_1247_v145_).length(0))
                index167_ = default__.safeIndex(441, (d_1179_v116_).length(0))
                rhs178_ = (d_1244_cf11_) * (len(d_1125_v91_))
                rhs179_ = d_1239_v143_
                rhs180_ = (self).fm6((d_1250_v146_)[default__.safeIndex(d_1234_cf14_, len(d_1250_v146_))], len(((d_1253_v149_)[d_993_v0_] if (d_993_v0_) in (d_1253_v149_) else d_1254_v150_)), ((self).f33) + ((self).f33), d_1256_v152_, globalState)
                rhs181_ = default__.safeDivisionInt(len(_dafny.Map({d_1133_v99_: d_1235_cf13_})), d_1233_cf15_)
                rhs182_ = (D7_DC19(d_1233_cf15_, d_1246_cf9_)).cf36
                lhs157_ = d_1179_v116_
                lhs158_ = default__.safeIndex(810, (d_1179_v116_).length(0))
                lhs159_ = d_1247_v145_
                lhs160_ = default__.safeIndex(474, (d_1247_v145_).length(0))
                lhs161_ = globalState
                lhs162_ = d_1179_v116_
                lhs163_ = default__.safeIndex(441, (d_1179_v116_).length(0))
                lhs157_[lhs158_] = rhs178_
                lhs159_[lhs160_] = rhs179_
                lhs161_.f8 = rhs180_
                lhs162_[lhs163_] = rhs181_
                d_1234_cf14_ = rhs182_
                d_1257_v153_: D8
                d_1257_v153_ = D8_DC22(d_1132_v98_)
                d_1258_v154_: _dafny.Array
                nw179_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_1258_v154_ = nw179_
                d_1259_v155_: _dafny.Array
                def lambda55_(d_1260_v91_, d_1261_v103_, d_1262_v102_):
                    def lambda56_(d_1263_i14_):
                        return _dafny.Map({D0_DC1(d_1260_v91_, len(d_1261_v103_), len(d_1260_v91_)): len(d_1262_v102_)})

                    return lambda56_

                init30_ = lambda55_(d_1125_v91_, d_1137_v103_, d_1136_v102_)
                nw180_ = _dafny.Array(None, 13)
                for i0_30_ in range(nw180_.length(0)):
                    nw180_[i0_30_] = init30_(i0_30_)
                d_1259_v155_ = nw180_
                index168_ = default__.safeIndex(368, (d_1258_v154_).length(0))
                (d_1258_v154_)[index168_] = d_1259_v155_
                index169_ = default__.safeIndex(368, (d_1258_v154_).length(0))
                rhs183_ = (d_1257_v153_ if d_1246_cf9_ else D8_DC22(d_1132_v98_))
                rhs184_ = d_1259_v155_
                rhs185_ = (0) - ((d_1234_cf14_) * (-400))
                lhs164_ = d_1258_v154_
                lhs165_ = default__.safeIndex(368, (d_1258_v154_).length(0))
                lhs166_ = globalState
                d_1257_v153_ = rhs183_
                lhs164_[lhs165_] = rhs184_
                lhs166_.f5 = rhs185_
            elif True:
                d_1264___mcc_h49_ = source25_.cf8
                d_1265_cf8_ = d_1264___mcc_h49_
                d_1132_v98_ = (_dafny.Map({d_1234_cf14_: (D0_DC2(d_1235_cf13_)).cf4})) | (d_1132_v98_)
                d_1266_v156_: C1
                nw181_ = C1()
                nw181_.ctor__()
                d_1266_v156_ = nw181_
                d_1126_v92_ = default__.fm20(d_993_v0_, d_1126_v92_, d_1234_cf14_, (d_1236_v142_)[default__.safeIndex(2, (d_1236_v142_).length(0))], globalState)
                (globalState).f13 = d_1125_v91_
        elif source24_.is_DC9:
            d_1267___mcc_h41_ = source24_.cf16
            d_1268___mcc_h42_ = source24_.cf17
            d_1269___mcc_h43_ = source24_.cf18
            d_1270___mcc_h44_ = source24_.cf19
            d_1271_cf19_ = d_1270___mcc_h44_
            d_1272_cf18_ = d_1269___mcc_h43_
            d_1273_cf17_ = d_1268___mcc_h42_
            d_1274_cf16_ = d_1267___mcc_h41_
            d_1275_v157_: C1
            nw182_ = C1()
            nw182_.ctor__()
            d_1275_v157_ = nw182_
            rhs186_ = d_1275_v157_
            rhs187_ = d_1272_cf18_
            lhs167_ = globalState
            d_1275_v157_ = rhs186_
            lhs167_.f24 = rhs187_
            d_1273_cf17_ = (d_993_v0_) or ((d_1272_cf18_) > (p0))
            (globalState).f5 = len(_dafny.SeqWithoutIsStrInference([d_1126_v92_ for d_1276_i15_ in range(default__.abs(656))]))
            index170_ = default__.safeIndex(618, (d_1179_v116_).length(0))
            (d_1179_v116_)[index170_] = -17
            index171_ = default__.safeIndex(618, (d_1179_v116_).length(0))
            (d_1179_v116_)[index171_] = d_1272_cf18_
        elif True:
            d_1277___mcc_h45_ = source24_.cf12
            d_1278_cf12_ = d_1277___mcc_h45_
            d_1279_v158_: _dafny.MultiSet
            d_1279_v158_ = _dafny.MultiSet([d_993_v0_, d_993_v0_])
            d_1280_v159_: _dafny.Map
            d_1280_v159_ = _dafny.Map({p0: (p0) * ((d_1279_v158_).cardinality)})
            d_1280_v159_ = (d_1280_v159_).set(p0, (447) * (803))
            if d_993_v0_:
                index172_ = default__.safeIndex(15, (d_1133_v99_).length(0))
                (d_1133_v99_)[index172_] = d_993_v0_
                d_1281_v160_: D7
                d_1281_v160_ = D7_DC18(p0, d_1179_v116_)
                d_1282_v161_: D7
                d_1282_v161_ = D7_DC21(d_1281_v160_)
                d_1283_v162_: _dafny.Seq
                d_1283_v162_ = _dafny.SeqWithoutIsStrInference([d_1282_v161_])
                d_1284_v163_: _dafny.Seq
                d_1284_v163_ = _dafny.SeqWithoutIsStrInference([d_1283_v162_, d_1283_v162_, d_1283_v162_, d_1283_v162_, (d_1283_v162_).set(default__.safeIndex(p0, len(d_1283_v162_)), d_1282_v161_)])
                d_1285_v164_: D9
                d_1285_v164_ = D9_DC27(d_1284_v163_)
                index173_ = default__.safeIndex(15, (d_1133_v99_).length(0))
                rhs188_ = default__.fm24((d_1137_v103_) != (_dafny.Set({p0})), d_993_v0_, d_993_v0_, p0, globalState)
                rhs189_ = d_993_v0_
                rhs190_ = d_993_v0_
                rhs191_ = len((d_1285_v164_).cf55)
                lhs168_ = globalState
                lhs169_ = d_1133_v99_
                lhs170_ = default__.safeIndex(15, (d_1133_v99_).length(0))
                lhs171_ = globalState
                lhs172_ = globalState
                lhs168_.f9 = rhs188_
                lhs169_[lhs170_] = rhs189_
                lhs171_.f25 = rhs190_
                lhs172_.f17 = rhs191_
                d_1286_v165_: _dafny.Array
                nw183_ = _dafny.Array(_dafny.CodePoint('D'), 14)
                d_1286_v165_ = nw183_
                index174_ = default__.safeIndex(884, (d_1286_v165_).length(0))
                (d_1286_v165_)[index174_] = d_1126_v92_
                index175_ = default__.safeIndex(884, (d_1286_v165_).length(0))
                rhs192_ = ((self).f33)[default__.safeIndex(p0, len((self).f33))]
                rhs193_ = default__.fm20((d_1133_v99_)[default__.safeIndex(15, (d_1133_v99_).length(0))], d_1126_v92_, p0, d_1126_v92_, globalState)
                lhs173_ = globalState
                lhs174_ = d_1286_v165_
                lhs175_ = default__.safeIndex(884, (d_1286_v165_).length(0))
                lhs173_.f5 = rhs192_
                lhs174_[lhs175_] = rhs193_
                d_1287_v166_: _dafny.Array
                def lambda57_(d_1288_p0_):
                    def lambda58_(d_1289_i16_):
                        return (D10_DC30(_dafny.MultiSet([d_1288_p0_]))).cf57

                    return lambda58_

                init31_ = lambda57_(p0)
                nw184_ = _dafny.Array(None, 23)
                for i0_31_ in range(nw184_.length(0)):
                    nw184_[i0_31_] = init31_(i0_31_)
                d_1287_v166_ = nw184_
                d_1290_v167_: _dafny.Map
                d_1290_v167_ = _dafny.Map({(self).fm7(p0, (d_1133_v99_)[default__.safeIndex(15, (d_1133_v99_).length(0))], globalState): d_1287_v166_})
                d_1290_v167_ = (d_1290_v167_).set(p0, (D11_DC33(d_1287_v166_)).cf65)
                (globalState).f13 = (d_1125_v91_) + (d_1125_v91_)
                d_1291_v168_: _dafny.Map
                d_1291_v168_ = _dafny.Map({p0: d_994_v1_})
                d_1292_v169_: _dafny.Set
                d_1292_v169_ = _dafny.Set({d_1125_v91_, d_1125_v91_})
                d_994_v1_ = (((d_1291_v168_)[len(d_1292_v169_)] if (len(d_1292_v169_)) in (d_1291_v168_) else default__.fm33(d_1126_v92_, ((d_1136_v102_)[(d_1133_v99_)[default__.safeIndex(15, (d_1133_v99_).length(0))]] if ((d_1133_v99_)[default__.safeIndex(15, (d_1133_v99_).length(0))]) in (d_1136_v102_) else not(d_993_v0_)), d_993_v0_, globalState))).intersection((d_994_v1_).intersection(d_994_v1_))
            elif True:
                (globalState).f25 = (d_993_v0_) and ((d_1278_cf12_) <= (d_1278_cf12_))
                d_1280_v159_ = (d_1280_v159_).set(len(_dafny.SeqWithoutIsStrInference([d_1126_v92_ for d_1293_i17_ in range(default__.abs(-525))])), p0)
                d_1294_v170_: _dafny.Seq
                d_1294_v170_ = _dafny.SeqWithoutIsStrInference([-396])
                d_1294_v170_ = (self).f33
                d_1295_v171_: _dafny.Set
                d_1295_v171_ = _dafny.Set({d_1294_v170_, (d_1294_v170_).set(default__.safeIndex(p0, len(d_1294_v170_)), len(_dafny.SeqWithoutIsStrInference([d_993_v0_])))})
                d_1296_v172_: _dafny.Map
                d_1296_v172_ = _dafny.Map({default__.fm21(len(d_1125_v91_), p0, d_993_v0_, d_993_v0_, globalState): d_1126_v92_})
                def iife109_():
                    coll67_ = _dafny.Set()
                    compr_67_: _dafny.Seq
                    for compr_67_ in (d_1296_v172_).keys.Elements:
                        d_1297_v173_: _dafny.Seq = compr_67_
                        if (d_1297_v173_) in (d_1296_v172_):
                            coll67_ = coll67_.union(_dafny.Set([d_1297_v173_]))
                    return _dafny.Set(coll67_)
                (globalState).f25 = (d_1295_v171_) == (iife109_()
                )
                index176_ = default__.safeIndex(304, (d_1133_v99_).length(0))
                (d_1133_v99_)[index176_] = True
                index177_ = default__.safeIndex(304, (d_1133_v99_).length(0))
                rhs194_ = d_993_v0_
                rhs195_ = d_993_v0_
                lhs176_ = d_1133_v99_
                lhs177_ = default__.safeIndex(304, (d_1133_v99_).length(0))
                lhs176_[lhs177_] = rhs194_
                d_993_v0_ = rhs195_
            (globalState).f5 = (0) - (p0)
            d_1298_v174_: D0
            d_1298_v174_ = D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmr"))))
            source26_ = d_1298_v174_
            if source26_.is_DC0:
                d_1299___mcc_h50_ = source26_.cf0
                d_1300_cf0_ = d_1299___mcc_h50_
                (globalState).f8 = default__.safeModuloInt((p0) + (d_1300_cf0_), (d_1279_v158_).cardinality)
                d_1301_v177_: _dafny.Seq
                def iife110_():
                    coll68_ = _dafny.Set()
                    compr_68_: int
                    for compr_68_ in _dafny.IntegerRange(-147, 573):
                        d_1302_v175_: int = compr_68_
                        if ((-147) <= (d_1302_v175_)) and ((d_1302_v175_) < (573)):
                            coll68_ = coll68_.union(_dafny.Set([(d_1302_v175_) - (-723)]))
                    return _dafny.Set(coll68_)
                def iife111_():
                    coll69_ = _dafny.Set()
                    compr_69_: int
                    for compr_69_ in (d_994_v1_).Elements:
                        d_1303_v176_: int = compr_69_
                        if (d_1303_v176_) in (d_994_v1_):
                            coll69_ = coll69_.union(_dafny.Set([(d_1303_v176_) * (39)]))
                    return _dafny.Set(coll69_)
                d_1301_v177_ = _dafny.SeqWithoutIsStrInference([d_1137_v103_, iife110_()
                , iife111_()
                , d_1137_v103_])
                (globalState).f25 = default__.fm24(d_993_v0_, d_993_v0_, d_993_v0_, len(d_1301_v177_), globalState)
                (globalState).f25 = d_993_v0_
                (globalState).f17 = ((d_994_v1_)[d_1300_cf0_] if (d_1300_cf0_) in (d_994_v1_) else default__.safeDivisionInt(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwviylb")))))
            elif source26_.is_DC1:
                d_1304___mcc_h51_ = source26_.cf1
                d_1305___mcc_h52_ = source26_.cf2
                d_1306___mcc_h53_ = source26_.cf3
                d_1307_cf3_ = d_1306___mcc_h53_
                d_1308_cf2_ = d_1305___mcc_h52_
                d_1309_cf1_ = d_1304___mcc_h51_
                d_1310_v178_: _dafny.Set
                d_1310_v178_ = _dafny.Set({(d_1278_cf12_)[default__.safeIndex(len(d_1125_v91_), len(d_1278_cf12_))]})
                rhs196_ = p0
                rhs197_ = ((d_993_v0_) == (d_993_v0_)) not in (d_1310_v178_)
                rhs198_ = d_993_v0_
                lhs178_ = globalState
                lhs179_ = globalState
                d_1308_cf2_ = rhs196_
                lhs178_.f25 = rhs197_
                lhs179_.f25 = rhs198_
                d_1280_v159_ = (self).f31
                d_1136_v102_ = (default__.fm34(d_993_v0_, d_993_v0_, d_993_v0_, ((d_1131_v97_)[d_993_v0_] if (d_993_v0_) in (d_1131_v97_) else p0), globalState)).set(d_993_v0_, d_993_v0_)
                index178_ = default__.safeIndex(277, (d_1179_v116_).length(0))
                (d_1179_v116_)[index178_] = d_1308_cf2_
                index179_ = default__.safeIndex(277, (d_1179_v116_).length(0))
                (d_1179_v116_)[index179_] = (0) - (d_1308_cf2_)
            elif source26_.is_DC2:
                d_1311___mcc_h54_ = source26_.cf4
                d_1312_cf4_ = d_1311___mcc_h54_
                (globalState).f5 = p0
                d_1131_v97_ = (d_1131_v97_).set((d_1126_v92_) in ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oc"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oc")))), _dafny.CodePoint('c'))), (self).fm7((d_1279_v158_).cardinality, d_1312_cf4_, globalState))
                d_1131_v97_ = (d_1131_v97_) | (d_1131_v97_)
                d_1313_v179_: _dafny.Seq
                d_1313_v179_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcmhqbkni"))])
                d_1314_v180_: _dafny.Seq
                d_1314_v180_ = d_1313_v179_
                d_1315_v181_: _dafny.Map
                d_1315_v181_ = _dafny.Map({len(d_1136_v102_): d_1314_v180_})
                d_1316_v182_: _dafny.MultiSet
                d_1316_v182_ = _dafny.MultiSet([d_1315_v181_])
                (globalState).f9 = default__.fm24(not((d_1312_cf4_ if d_993_v0_ else d_993_v0_)), (d_1316_v182_).isdisjoint(d_1316_v182_), d_993_v0_, p0, globalState)
            elif source26_.is_DC3:
                d_1317___mcc_h55_ = source26_.cf5
                d_1318___mcc_h56_ = source26_.cf6
                d_1319_cf6_ = d_1318___mcc_h56_
                d_1320_cf5_ = d_1317___mcc_h55_
                d_1179_v116_ = d_1179_v116_
                (globalState).f9 = (d_1320_cf5_) != (p0)
                d_1321_v183_: C0
                nw185_ = C0()
                nw185_.ctor__()
                d_1321_v183_ = nw185_
                d_1280_v159_ = (d_1280_v159_).set((D7_DC19(p0, d_993_v0_)).cf36, len(d_1125_v91_))
            elif True:
                d_1322___mcc_h57_ = source26_.cf7
                d_1323_cf7_ = d_1322___mcc_h57_
                d_1324_v184_: _dafny.Set
                d_1324_v184_ = _dafny.Set({d_1131_v97_})
                (globalState).f8 = len(d_1324_v184_)
                (globalState).f5 = p0
                (globalState).f25 = False
                (globalState).f23 = d_1125_v91_
        r0 = p0
        return r0

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        d_1325_v0_: _dafny.Set
        d_1325_v0_ = _dafny.Set({p1, p1})
        index180_ = default__.safeIndex(241, (p0).length(0))
        (p0)[index180_] = len(d_1325_v0_)
        d_1326_v1_: int
        d_1326_v1_ = -721
        d_1327_v2_: _dafny.Seq
        d_1327_v2_ = _dafny.SeqWithoutIsStrInference([False])
        d_1328_v3_: _dafny.MultiSet
        d_1328_v3_ = _dafny.MultiSet([d_1326_v1_])
        d_1329_v4_: _dafny.Seq
        d_1329_v4_ = _dafny.SeqWithoutIsStrInference([d_1328_v3_, d_1328_v3_])
        index181_ = default__.safeIndex(241, (p0).length(0))
        (p0)[index181_] = default__.safeDivisionInt((d_1326_v1_) + (len(_dafny.Map({d_1326_v1_: (d_1327_v2_)[default__.safeIndex(d_1326_v1_, len(d_1327_v2_))]}))), len((d_1329_v4_).set(default__.safeIndex(d_1326_v1_, len(d_1329_v4_)), d_1328_v3_)))
        (globalState).f25 = p1
        d_1330_v5_: _dafny.Array
        def lambda59_(d_1331_p1_):
            def lambda60_(d_1332_i0_):
                return not (d_1331_p1_) or (not(False))

            return lambda60_

        init32_ = lambda59_(p1)
        nw186_ = _dafny.Array(None, 1)
        for i0_32_ in range(nw186_.length(0)):
            nw186_[i0_32_] = init32_(i0_32_)
        d_1330_v5_ = nw186_
        index182_ = default__.safeIndex(494, (d_1330_v5_).length(0))
        (d_1330_v5_)[index182_] = p1
        index183_ = default__.safeIndex(494, (d_1330_v5_).length(0))
        (d_1330_v5_)[index183_] = (d_1325_v0_).issubset(d_1325_v0_)
        d_1333_v6_: D2
        d_1333_v6_ = D2_DC7(d_1327_v2_)
        if (d_1327_v2_)[default__.safeIndex(len((d_1333_v6_).cf12), len(d_1327_v2_))]:
            d_1334_v7_: _dafny.Array
            nw187_ = _dafny.Array(_dafny.MultiSet({}), 16)
            d_1334_v7_ = nw187_
            d_1335_v8_: str
            d_1335_v8_ = _dafny.CodePoint('m')
            index184_ = default__.safeIndex(807, (d_1334_v7_).length(0))
            (d_1334_v7_)[index184_] = default__.fm35(d_1335_v8_, p1, d_1326_v1_, globalState)
            d_1336_v9_: _dafny.Map
            d_1336_v9_ = _dafny.Map({(p0)[default__.safeIndex(241, (p0).length(0))]: True})
            d_1337_v10_: D2
            d_1337_v10_ = D2_DC9(d_1326_v1_, True, len((d_1336_v9_).set((0) - ((p0)[default__.safeIndex(241, (p0).length(0))]), p1)), -233)
            d_1338_v11_: _dafny.MultiSet
            d_1338_v11_ = _dafny.MultiSet([d_1337_v10_])
            index185_ = default__.safeIndex(807, (d_1334_v7_).length(0))
            (d_1334_v7_)[index185_] = d_1338_v11_
            d_1339_v12_: D0
            d_1339_v12_ = D0_DC2(p1)
            source27_ = d_1339_v12_
            if source27_.is_DC0:
                d_1340___mcc_h0_ = source27_.cf0
                d_1341_cf0_ = d_1340___mcc_h0_
                d_1342_v13_: _dafny.Array
                nw188_ = _dafny.Array(_dafny.MultiSet({}), 12)
                d_1342_v13_ = nw188_
                d_1343_v14_: _dafny.MultiSet
                d_1343_v14_ = _dafny.MultiSet([d_1330_v5_])
                index186_ = default__.safeIndex(343, (d_1342_v13_).length(0))
                (d_1342_v13_)[index186_] = d_1343_v14_
                index187_ = default__.safeIndex(343, (d_1342_v13_).length(0))
                (d_1342_v13_)[index187_] = d_1343_v14_
                d_1344_v15_: _dafny.Seq
                d_1344_v15_ = _dafny.SeqWithoutIsStrInference([len((self).f31)])
                d_1344_v15_ = _dafny.SeqWithoutIsStrInference([d_1326_v1_ for d_1345_i1_ in range(default__.abs(837))])
                d_1346_v16_: _dafny.Array
                nw189_ = _dafny.Array(_dafny.Map({}), 19)
                d_1346_v16_ = nw189_
                index188_ = default__.safeIndex(398, (d_1346_v16_).length(0))
                def iife112_():
                    coll70_ = _dafny.Map()
                    compr_70_: int
                    for compr_70_ in ((self).f33).Elements:
                        d_1347_v17_: int = compr_70_
                        if (d_1347_v17_) in ((self).f33):
                            coll70_[(d_1347_v17_) - (d_1326_v1_)] = (_dafny.MultiSet([p1])).cardinality
                    return _dafny.Map(coll70_)
                (d_1346_v16_)[index188_] = (_dafny.Map({len((self).f33): (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]})) | (_dafny.Map({len(iife112_()
                ): False}))
                d_1348_v18_: D8
                d_1348_v18_ = D8_DC22(d_1336_v9_)
                index189_ = default__.safeIndex(398, (d_1346_v16_).length(0))
                (d_1346_v16_)[index189_] = (d_1348_v18_).cf41
                d_1349_v19_: _dafny.Array
                def lambda61_(d_1350_v5_, d_1351_p1_):
                    def lambda62_(d_1352_i2_):
                        return _dafny.Map({(d_1350_v5_)[default__.safeIndex(494, (d_1350_v5_).length(0))]: d_1351_p1_})

                    return lambda62_

                init33_ = lambda61_(d_1330_v5_, p1)
                nw190_ = _dafny.Array(None, 21)
                for i0_33_ in range(nw190_.length(0)):
                    nw190_[i0_33_] = init33_(i0_33_)
                d_1349_v19_ = nw190_
                d_1353_v20_: _dafny.MultiSet
                d_1353_v20_ = _dafny.MultiSet([d_1349_v19_])
                index190_ = default__.safeIndex(241, (p0).length(0))
                (p0)[index190_] = ((d_1353_v20_)[d_1349_v19_] if (d_1349_v19_) in (d_1353_v20_) else ((p0)[default__.safeIndex(241, (p0).length(0))]) - ((0) - (d_1341_cf0_)))
            elif source27_.is_DC1:
                d_1354___mcc_h1_ = source27_.cf1
                d_1355___mcc_h2_ = source27_.cf2
                d_1356___mcc_h3_ = source27_.cf3
                d_1357_cf3_ = d_1356___mcc_h3_
                d_1358_cf2_ = d_1355___mcc_h2_
                d_1359_cf1_ = d_1354___mcc_h1_
                d_1360_v21_: _dafny.Set
                d_1360_v21_ = _dafny.Set({(d_1336_v9_) | (d_1336_v9_)})
                index191_ = default__.safeIndex(241, (p0).length(0))
                rhs199_ = len(d_1360_v21_)
                rhs200_ = (d_1337_v10_).cf17
                lhs180_ = p0
                lhs181_ = default__.safeIndex(241, (p0).length(0))
                lhs182_ = globalState
                lhs180_[lhs181_] = rhs199_
                lhs182_.f9 = rhs200_
                d_1361_v22_: _dafny.Array
                nw191_ = _dafny.Array(_dafny.MultiSet({}), 28)
                d_1361_v22_ = nw191_
                index192_ = default__.safeIndex(701, (d_1361_v22_).length(0))
                (d_1361_v22_)[index192_] = _dafny.MultiSet((self).f33)
                index193_ = default__.safeIndex(701, (d_1361_v22_).length(0))
                (d_1361_v22_)[index193_] = (_dafny.MultiSet([(p0)[default__.safeIndex(241, (p0).length(0))], (0) - (d_1326_v1_), d_1358_cf2_])).intersection(d_1328_v3_)
                d_1362_v23_: _dafny.Seq
                d_1362_v23_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({(0) - (((d_1361_v22_)[default__.safeIndex(701, (d_1361_v22_).length(0))]).cardinality), (p0)[default__.safeIndex(241, (p0).length(0))], (0) - (d_1326_v1_)}))])
                d_1363_v24_: _dafny.Map
                d_1363_v24_ = _dafny.Map({d_1357_cf3_: d_1362_v23_})
                d_1364_v25_: _dafny.MultiSet
                d_1364_v25_ = _dafny.MultiSet([(d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]])
                rhs201_ = d_1358_cf2_
                rhs202_ = (((d_1363_v24_)[(d_1364_v25_).cardinality] if ((d_1364_v25_).cardinality) in (d_1363_v24_) else d_1362_v23_)) + (_dafny.SeqWithoutIsStrInference([d_1326_v1_]))
                rhs203_ = ((len(_dafny.Map({d_1335_v8_: 647})) if not(p1) else (self).fm7((D10_DC31((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))], default__.fm21(d_1326_v1_, (p0)[default__.safeIndex(241, (p0).length(0))], (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))], p1, globalState), True, d_1326_v1_)).cf61, p1, globalState))) - (d_1326_v1_)
                rhs204_ = d_1357_cf3_
                lhs183_ = globalState
                lhs184_ = globalState
                lhs185_ = globalState
                lhs183_.f24 = rhs201_
                d_1362_v23_ = rhs202_
                lhs184_.f8 = rhs203_
                lhs185_.f24 = rhs204_
                d_1365_v26_: C0
                nw192_ = C0()
                nw192_.ctor__()
                d_1365_v26_ = nw192_
                d_1366_v27_: _dafny.MultiSet
                d_1366_v27_ = _dafny.MultiSet([d_1365_v26_, d_1365_v26_, d_1365_v26_, d_1365_v26_])
                d_1366_v27_ = d_1366_v27_
            elif source27_.is_DC2:
                d_1367___mcc_h4_ = source27_.cf4
                d_1368_cf4_ = d_1367___mcc_h4_
                d_1369_v28_: _dafny.Seq
                d_1369_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrvigief"))
                d_1370_v29_: D0
                d_1370_v29_ = D0_DC1(d_1369_v28_, len(d_1369_v28_), (p0)[default__.safeIndex(241, (p0).length(0))])
                (globalState).f5 = (self).fm7((d_1370_v29_).cf2, d_1368_cf4_, globalState)
                (globalState).f0 = default__.safeModuloInt((p0)[default__.safeIndex(241, (p0).length(0))], (p0)[default__.safeIndex(241, (p0).length(0))])
                d_1371_v30_: _dafny.Array
                nw193_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
                d_1371_v30_ = nw193_
                d_1372_v31_: _dafny.Map
                d_1372_v31_ = _dafny.Map({default__.fm24(False, (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))], d_1368_cf4_, d_1326_v1_, globalState): d_1371_v30_})
                d_1372_v31_ = (d_1372_v31_).set((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))], d_1371_v30_)
                (globalState).f23 = default__.fm36(d_1368_cf4_, not(d_1368_cf4_), (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))], globalState)
            elif source27_.is_DC3:
                d_1373___mcc_h5_ = source27_.cf5
                d_1374___mcc_h6_ = source27_.cf6
                d_1375_cf6_ = d_1374___mcc_h6_
                d_1376_cf5_ = d_1373___mcc_h5_
                rhs205_ = not ((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]) or ((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))])
                rhs206_ = (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]
                lhs186_ = globalState
                lhs186_.f9 = rhs205_
                r1 = rhs206_
                d_1377_v32_: _dafny.Array
                def lambda63_(d_1378_cf5_):
                    def lambda64_(d_1379_i3_):
                        return (d_1379_i3_) + (len(_dafny.Set({d_1378_cf5_, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubffx")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eliyqgiwy"))])), -437})))

                    return lambda64_

                init34_ = lambda63_(d_1376_cf5_)
                nw194_ = _dafny.Array(None, 28)
                for i0_34_ in range(nw194_.length(0)):
                    nw194_[i0_34_] = init34_(i0_34_)
                d_1377_v32_ = nw194_
                d_1380_v33_: D7
                d_1380_v33_ = D7_DC19(865, p1)
                rhs207_ = len(default__.fm36(p1, True, (d_1380_v33_).cf37, globalState))
                rhs208_ = (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]
                rhs209_ = d_1377_v32_
                lhs187_ = globalState
                lhs188_ = globalState
                lhs187_.f0 = rhs207_
                lhs188_.f9 = rhs208_
                d_1377_v32_ = rhs209_
                d_1381_v34_: _dafny.Map
                d_1381_v34_ = _dafny.Map({p1: (p0)[default__.safeIndex(241, (p0).length(0))]})
                d_1382_v35_: _dafny.Set
                d_1383_v36_: _dafny.Map
                d_1384_v37_: str
                out34_: _dafny.Set
                out35_: _dafny.Map
                out36_: str
                out34_, out35_, out36_ = (self).m5(d_1330_v5_, d_1381_v34_, globalState)
                d_1382_v35_ = out34_
                d_1383_v36_ = out35_
                d_1384_v37_ = out36_
                d_1385_v38_: _dafny.Map
                d_1385_v38_ = _dafny.Map({not((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]): _dafny.Map({d_1377_v32_: 708})})
                d_1386_v39_: _dafny.Map
                d_1386_v39_ = _dafny.Map({p0: d_1376_cf5_})
                d_1385_v38_ = (d_1385_v38_).set(p1, d_1386_v39_)
            elif True:
                d_1387___mcc_h7_ = source27_.cf7
                d_1388_cf7_ = d_1387___mcc_h7_
                d_1389_v40_: _dafny.Array
                def lambda65_(d_1390_i4_):
                    return (self).f33

                init35_ = lambda65_
                nw195_ = _dafny.Array(None, 20)
                for i0_35_ in range(nw195_.length(0)):
                    nw195_[i0_35_] = init35_(i0_35_)
                d_1389_v40_ = nw195_
                d_1391_v41_: D10
                d_1391_v41_ = D10_DC31((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))], (self).f33, p1, 869)
                d_1392_v42_: _dafny.Map
                d_1392_v42_ = _dafny.Map({d_1327_v2_: (d_1391_v41_).cf58})
                index194_ = default__.safeIndex(785, (d_1389_v40_).length(0))
                (d_1389_v40_)[index194_] = (((self).f33) + ((self).f33)).set(default__.safeIndex((self).fm6(d_1336_v9_, ((self).f33)[default__.safeIndex((p0)[default__.safeIndex(241, (p0).length(0))], len((self).f33))], (self).f33, d_1392_v42_, globalState), len(((self).f33) + ((self).f33))), d_1326_v1_)
                index195_ = default__.safeIndex(785, (d_1389_v40_).length(0))
                (d_1389_v40_)[index195_] = ((((self).f33) + ((self).f33)).set(default__.safeIndex(d_1326_v1_, len(((self).f33) + ((self).f33))), d_1326_v1_)) + ((self).f33)
                d_1393_v43_: C1
                nw196_ = C1()
                nw196_.ctor__()
                d_1393_v43_ = nw196_
                d_1394_v44_: _dafny.Set
                d_1394_v44_ = _dafny.Set({d_1326_v1_, d_1326_v1_, (p0)[default__.safeIndex(241, (p0).length(0))]})
                (globalState).f9 = (_dafny.Set({(self).fm7(d_1326_v1_, p1, globalState)})).issubset((d_1394_v44_) | (d_1394_v44_))
                d_1395_v45_: _dafny.MultiSet
                d_1395_v45_ = _dafny.MultiSet([((0) - ((p0)[default__.safeIndex(241, (p0).length(0))])) < (d_1326_v1_), p1])
                d_1395_v45_ = d_1395_v45_
            d_1396_v46_: C0
            nw197_ = C0()
            nw197_.ctor__()
            d_1396_v46_ = nw197_
            d_1397_v47_: _dafny.Map
            d_1397_v47_ = _dafny.Map({not((d_1327_v2_) < (d_1327_v2_)): (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]})
            d_1397_v47_ = (d_1397_v47_).set(((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))] if p1 else p1), p1)
            d_1398_v48_: _dafny.Map
            d_1398_v48_ = _dafny.Map({(d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]: d_1335_v8_})
            d_1399_v49_: _dafny.Map
            d_1399_v49_ = _dafny.Map({d_1398_v48_: d_1335_v8_})
            d_1399_v49_ = (d_1399_v49_).set(d_1398_v48_, d_1335_v8_)
        elif True:
            d_1400_v50_: C1
            nw198_ = C1()
            nw198_.ctor__()
            d_1400_v50_ = nw198_
            d_1401_v51_: C1
            nw199_ = C1()
            nw199_.ctor__()
            d_1401_v51_ = nw199_
            d_1402_v52_: str
            d_1402_v52_ = _dafny.CodePoint('o')
            d_1403_v53_: _dafny.Array
            nw200_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_1403_v53_ = nw200_
            d_1404_v54_: _dafny.Map
            d_1404_v54_ = _dafny.Map({d_1403_v53_: d_1326_v1_})
            rhs210_ = ((d_1404_v54_)[d_1403_v53_] if (d_1403_v53_) in (d_1404_v54_) else d_1326_v1_)
            rhs211_ = d_1402_v52_
            lhs189_ = globalState
            lhs189_.f0 = rhs210_
            d_1402_v52_ = rhs211_
            (globalState).f17 = default__.safeDivisionInt((d_1326_v1_ if p1 else d_1326_v1_), d_1326_v1_)
            d_1405_v55_: _dafny.MultiSet
            d_1405_v55_ = _dafny.MultiSet([(d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]])
            d_1406_v56_: D0
            d_1406_v56_ = D0_DC2(p1)
            d_1407_v57_: _dafny.Map
            d_1407_v57_ = _dafny.Map({(d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]: (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]})
            d_1408_v58_: _dafny.Map
            d_1408_v58_ = _dafny.Map({d_1405_v55_: D3_DC10(d_1407_v57_)})
            d_1409_v59_: _dafny.Map
            d_1409_v59_ = _dafny.Map({d_1406_v56_: d_1408_v58_})
            d_1410_v60_: D12
            d_1410_v60_ = D12_DC36(d_1408_v58_)
            if (d_1405_v55_) not in (((d_1409_v59_)[d_1406_v56_] if (d_1406_v56_) in (d_1409_v59_) else (d_1410_v60_).cf69)):
                (globalState).f8 = default__.safeModuloInt(((p0)[default__.safeIndex(241, (p0).length(0))]) - (d_1326_v1_), (p0)[default__.safeIndex(241, (p0).length(0))])
                (globalState).f9 = ((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]) and ((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))])
                d_1411_v61_: C1
                nw201_ = C1()
                nw201_.ctor__()
                d_1411_v61_ = nw201_
                d_1412_v62_: _dafny.Array
                nw202_ = _dafny.Array(_dafny.MultiSet({}), 22)
                d_1412_v62_ = nw202_
                index196_ = default__.safeIndex(514, (d_1412_v62_).length(0))
                (d_1412_v62_)[index196_] = _dafny.MultiSet([d_1330_v5_])
                d_1413_v63_: _dafny.MultiSet
                d_1413_v63_ = _dafny.MultiSet([d_1330_v5_, d_1330_v5_, d_1330_v5_, d_1330_v5_])
                index197_ = default__.safeIndex(514, (d_1412_v62_).length(0))
                (d_1412_v62_)[index197_] = (d_1413_v63_).set(d_1330_v5_, default__.abs(default__.safeModuloInt(d_1326_v1_, (0) - ((0) - ((p0)[default__.safeIndex(241, (p0).length(0))])))))
                d_1414_v64_: _dafny.MultiSet
                d_1414_v64_ = _dafny.MultiSet([p0])
                index198_ = default__.safeIndex(241, (p0).length(0))
                (p0)[index198_] = ((d_1414_v64_).set(p0, default__.abs((p0)[default__.safeIndex(241, (p0).length(0))]))).cardinality
            elif True:
                d_1327_v2_ = d_1327_v2_
                (globalState).f17 = 621
                index199_ = default__.safeIndex(241, (p0).length(0))
                (p0)[index199_] = (0) - (d_1326_v1_)
                (globalState).f0 = 753
                d_1415_v65_: D7
                d_1415_v65_ = D7_DC19(d_1326_v1_, (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))])
                d_1416_v66_: D7
                d_1416_v66_ = D7_DC21(d_1415_v65_)
                d_1417_v67_: _dafny.Seq
                d_1417_v67_ = _dafny.SeqWithoutIsStrInference([d_1416_v66_])
                d_1418_v68_: _dafny.Seq
                d_1418_v68_ = _dafny.SeqWithoutIsStrInference([d_1417_v67_, d_1417_v67_])
                d_1419_v69_: D9
                d_1419_v69_ = D9_DC27(d_1418_v68_)
                pat_let_tv22_ = d_1418_v68_
                d_1420_v70_: _dafny.Array
                nw203_ = _dafny.Array(None, 16)
                nw203_[int(0)] = d_1419_v69_
                def iife113_(_pat_let21_0):
                    def iife114_(d_1421_dt__update__tmp_h0_):
                        def iife115_(_pat_let22_0):
                            def iife116_(d_1422_dt__update_hcf55_h0_):
                                return D9_DC27(d_1422_dt__update_hcf55_h0_)
                            return iife116_(_pat_let22_0)
                        return iife115_(pat_let_tv22_)
                    return iife114_(_pat_let21_0)
                nw203_[int(1)] = iife113_(d_1419_v69_)
                nw203_[int(2)] = d_1419_v69_
                nw203_[int(3)] = d_1419_v69_
                nw203_[int(4)] = d_1419_v69_
                nw203_[int(5)] = D9_DC27(d_1418_v68_)
                nw203_[int(6)] = d_1419_v69_
                nw203_[int(7)] = d_1419_v69_
                nw203_[int(8)] = d_1419_v69_
                nw203_[int(9)] = d_1419_v69_
                nw203_[int(10)] = d_1419_v69_
                nw203_[int(11)] = D9_DC27(d_1418_v68_)
                nw203_[int(12)] = d_1419_v69_
                nw203_[int(13)] = D9_DC27(d_1418_v68_)
                nw203_[int(14)] = d_1419_v69_
                nw203_[int(15)] = d_1419_v69_
                d_1420_v70_ = nw203_
                index200_ = default__.safeIndex(645, (d_1420_v70_).length(0))
                (d_1420_v70_)[index200_] = D9_DC27((d_1418_v68_).set(default__.safeIndex((p0)[default__.safeIndex(241, (p0).length(0))], len(d_1418_v68_)), d_1417_v67_))
                index201_ = default__.safeIndex(645, (d_1420_v70_).length(0))
                (d_1420_v70_)[index201_] = D9_DC27(d_1418_v68_)
        d_1423_v71_: _dafny.Map
        d_1423_v71_ = _dafny.Map({d_1326_v1_: p1})
        d_1424_v72_: _dafny.Map
        d_1424_v72_ = _dafny.Map({d_1327_v2_: (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]})
        d_1326_v1_ = (self).fm6(d_1423_v71_, (d_1326_v1_) * (d_1326_v1_), ((self).f33 if (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))] else (self).f33), d_1424_v72_, globalState)
        d_1425_v73_: _dafny.Seq
        d_1425_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfostyx"))
        hi5_ = len(d_1425_v73_)
        for d_1426_i5_ in range((p0)[default__.safeIndex(241, (p0).length(0))], hi5_):
            d_1328_v3_ = _dafny.MultiSet(((self).f33) + (_dafny.SeqWithoutIsStrInference([(D1_DC6(p1, _dafny.CodePoint('h'), (p0)[default__.safeIndex(241, (p0).length(0))])).cf11 for d_1427_i6_ in range(default__.abs(-555))])))
            if p1:
                r0 = 895
                r1 = p1
                d_1428_v74_: C0
                nw204_ = C0()
                nw204_.ctor__()
                d_1428_v74_ = nw204_
                d_1429_v75_: str
                d_1429_v75_ = _dafny.CodePoint('s')
                index202_ = default__.safeIndex(494, (d_1330_v5_).length(0))
                rhs212_ = d_1428_v74_
                rhs213_ = (d_1330_v5_ if p1 else d_1330_v5_)
                rhs214_ = (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]
                rhs215_ = d_1429_v75_
                lhs190_ = globalState
                lhs191_ = d_1330_v5_
                lhs192_ = default__.safeIndex(494, (d_1330_v5_).length(0))
                d_1428_v74_ = rhs212_
                lhs190_.f1 = rhs213_
                lhs191_[lhs192_] = rhs214_
                d_1429_v75_ = rhs215_
                (globalState).f24 = (p0)[default__.safeIndex(241, (p0).length(0))]
                (globalState).f5 = d_1326_v1_
            elif True:
                index203_ = default__.safeIndex(494, (d_1330_v5_).length(0))
                (d_1330_v5_)[index203_] = (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]
                (globalState).f23 = d_1425_v73_
                d_1430_v76_: _dafny.Seq
                d_1430_v76_ = _dafny.SeqWithoutIsStrInference([d_1326_v1_, d_1426_i5_, default__.safeDivisionInt(len(default__.fm0((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))], d_1425_v73_, d_1325_v0_, -115, globalState)), (p0)[default__.safeIndex(241, (p0).length(0))]), (0) - ((p0)[default__.safeIndex(241, (p0).length(0))])])
                d_1430_v76_ = d_1430_v76_
                d_1431_v77_: _dafny.Seq
                d_1431_v77_ = _dafny.SeqWithoutIsStrInference([(((self).f31).set(len(d_1430_v76_), d_1426_i5_)) | ((self).f31)])
                d_1431_v77_ = _dafny.SeqWithoutIsStrInference([(self).f31, (self).f31])
                (globalState).f24 = (p0)[default__.safeIndex(241, (p0).length(0))]
            d_1432_v78_: _dafny.Array
            nw205_ = _dafny.Array(D7.default()(), 8)
            d_1432_v78_ = nw205_
            rhs216_ = ((d_1328_v3_) - (d_1328_v3_)) - ((d_1328_v3_).set((d_1328_v3_).cardinality, default__.abs(d_1326_v1_)))
            rhs217_ = (default__.fm17(globalState)) | (d_1325_v0_)
            rhs218_ = ((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]) and ((d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))])
            rhs219_ = (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))]
            rhs220_ = d_1432_v78_
            lhs193_ = globalState
            d_1328_v3_ = rhs216_
            d_1325_v0_ = rhs217_
            r1 = rhs218_
            lhs193_.f9 = rhs219_
            d_1432_v78_ = rhs220_
            d_1433_v79_: _dafny.Array
            nw206_ = _dafny.Array(_dafny.CodePoint('D'), 6)
            d_1433_v79_ = nw206_
            d_1434_v80_: str
            d_1434_v80_ = _dafny.CodePoint('f')
            index204_ = default__.safeIndex(680, (d_1433_v79_).length(0))
            (d_1433_v79_)[index204_] = d_1434_v80_
            index205_ = default__.safeIndex(680, (d_1433_v79_).length(0))
            rhs221_ = d_1434_v80_
            rhs222_ = len((self).f33)
            rhs223_ = d_1426_i5_
            rhs224_ = default__.safeModuloInt(d_1426_i5_, (d_1326_v1_ if (d_1330_v5_)[default__.safeIndex(494, (d_1330_v5_).length(0))] else d_1426_i5_))
            lhs194_ = d_1433_v79_
            lhs195_ = default__.safeIndex(680, (d_1433_v79_).length(0))
            lhs196_ = globalState
            lhs197_ = globalState
            lhs198_ = globalState
            lhs194_[lhs195_] = rhs221_
            lhs196_.f0 = rhs222_
            lhs197_.f24 = rhs223_
            lhs198_.f5 = rhs224_
        r0 = default__.safeDivisionInt(-473, (p0)[default__.safeIndex(241, (p0).length(0))])
        r1 = p1
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Array = _dafny.Array(None, 0)
        if not(not(not (True) or (p0))):
            d_1435_v0_: _dafny.Array
            nw207_ = _dafny.Array(_dafny.Seq({}), 20)
            d_1435_v0_ = nw207_
            d_1435_v0_ = d_1435_v0_
            (globalState).f25 = p0
            d_1436_v1_: int
            d_1436_v1_ = 525
            r0 = (d_1436_v1_) * (d_1436_v1_)
            d_1437_v2_: _dafny.Array
            nw208_ = _dafny.Array(_dafny.CodePoint('D'), 21)
            d_1437_v2_ = nw208_
            d_1438_v3_: str
            d_1438_v3_ = _dafny.CodePoint('b')
            index206_ = default__.safeIndex(605, (d_1437_v2_).length(0))
            (d_1437_v2_)[index206_] = d_1438_v3_
            d_1439_v4_: _dafny.Map
            d_1439_v4_ = _dafny.Map({d_1438_v3_: p0})
            d_1440_v5_: _dafny.MultiSet
            d_1440_v5_ = _dafny.MultiSet([(((d_1439_v4_)[d_1438_v3_] if (d_1438_v3_) in (d_1439_v4_) else p0)) and (p0), p0])
            d_1441_v6_: _dafny.Array
            def lambda66_(d_1442_i0_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obtc"))

            init36_ = lambda66_
            nw209_ = _dafny.Array(None, 29)
            for i0_36_ in range(nw209_.length(0)):
                nw209_[i0_36_] = init36_(i0_36_)
            d_1441_v6_ = nw209_
            d_1443_v7_: _dafny.Seq
            d_1443_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kn"))
            index207_ = default__.safeIndex(905, (d_1441_v6_).length(0))
            (d_1441_v6_)[index207_] = (d_1443_v7_) + (d_1443_v7_)
            d_1444_v8_: _dafny.Seq
            d_1444_v8_ = _dafny.SeqWithoutIsStrInference([p0])
            index208_ = default__.safeIndex(605, (d_1437_v2_).length(0))
            index209_ = default__.safeIndex(905, (d_1441_v6_).length(0))
            rhs225_ = True
            rhs226_ = d_1438_v3_
            rhs227_ = (_dafny.MultiSet((d_1444_v8_) + (_dafny.SeqWithoutIsStrInference([p0, p0])))).intersection((default__.fm37(p0, p0, d_1436_v1_, d_1436_v1_, globalState)) - (d_1440_v5_))
            rhs228_ = d_1443_v7_
            lhs199_ = globalState
            lhs200_ = d_1437_v2_
            lhs201_ = default__.safeIndex(605, (d_1437_v2_).length(0))
            lhs202_ = d_1441_v6_
            lhs203_ = default__.safeIndex(905, (d_1441_v6_).length(0))
            lhs199_.f9 = rhs225_
            lhs200_[lhs201_] = rhs226_
            d_1440_v5_ = rhs227_
            lhs202_[lhs203_] = rhs228_
            d_1445_v9_: C1
            nw210_ = C1()
            nw210_.ctor__()
            d_1445_v9_ = nw210_
        elif True:
            d_1446_v10_: int
            d_1446_v10_ = 896
            d_1447_v11_: _dafny.Seq
            d_1447_v11_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_1448_v12_: _dafny.MultiSet
            d_1448_v12_ = _dafny.MultiSet([len(d_1447_v11_), d_1446_v10_])
            d_1449_v13_: _dafny.Seq
            d_1449_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyukiyfki"))
            rhs229_ = d_1446_v10_
            rhs230_ = default__.fm24(p0, p0, ((d_1448_v12_).set(d_1446_v10_, default__.abs(len(d_1449_v13_)))).issubset(d_1448_v12_), d_1446_v10_, globalState)
            lhs204_ = globalState
            lhs205_ = globalState
            lhs204_.f8 = rhs229_
            lhs205_.f9 = rhs230_
            d_1450_v14_: _dafny.Map
            d_1450_v14_ = _dafny.Map({p0: d_1446_v10_})
            (globalState).f9 = (((self).f33).set(default__.safeIndex(-635, len((self).f33)), ((d_1450_v14_)[p0] if (p0) in (d_1450_v14_) else d_1446_v10_))) < (_dafny.SeqWithoutIsStrInference([d_1446_v10_]))
            (globalState).f9 = p0
            d_1451_v15_: _dafny.Array
            nw211_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_1451_v15_ = nw211_
            d_1452_v16_: D9
            d_1452_v16_ = D9_DC28()
            d_1453_v17_: _dafny.Array
            nw212_ = _dafny.Array(None, 29)
            nw212_[int(0)] = D9_DC28()
            nw212_[int(1)] = d_1452_v16_
            nw212_[int(2)] = d_1452_v16_
            nw212_[int(3)] = default__.fm38(not(p0), globalState)
            nw212_[int(4)] = d_1452_v16_
            nw212_[int(5)] = D9_DC28()
            nw212_[int(6)] = D9_DC28()
            nw212_[int(7)] = d_1452_v16_
            nw212_[int(8)] = d_1452_v16_
            nw212_[int(9)] = d_1452_v16_
            nw212_[int(10)] = D9_DC28()
            nw212_[int(11)] = d_1452_v16_
            nw212_[int(12)] = d_1452_v16_
            nw212_[int(13)] = d_1452_v16_
            nw212_[int(14)] = d_1452_v16_
            nw212_[int(15)] = d_1452_v16_
            nw212_[int(16)] = d_1452_v16_
            nw212_[int(17)] = d_1452_v16_
            nw212_[int(18)] = default__.fm38(p0, globalState)
            nw212_[int(19)] = d_1452_v16_
            nw212_[int(20)] = d_1452_v16_
            nw212_[int(21)] = d_1452_v16_
            nw212_[int(22)] = D9_DC28()
            nw212_[int(23)] = d_1452_v16_
            nw212_[int(24)] = d_1452_v16_
            nw212_[int(25)] = D9_DC28()
            nw212_[int(26)] = default__.fm38(p0, globalState)
            nw212_[int(27)] = D9_DC28()
            nw212_[int(28)] = d_1452_v16_
            d_1453_v17_ = nw212_
            index210_ = default__.safeIndex(147, (d_1451_v15_).length(0))
            (d_1451_v15_)[index210_] = d_1453_v17_
            index211_ = default__.safeIndex(147, (d_1451_v15_).length(0))
            nw213_ = _dafny.Array(D9.default()(), 28)
            (d_1451_v15_)[index211_] = nw213_
            d_1454_v18_: _dafny.Array
            nw214_ = _dafny.Array(_dafny.Seq({}), 29)
            d_1454_v18_ = nw214_
            d_1455_v19_: _dafny.Seq
            d_1455_v19_ = _dafny.SeqWithoutIsStrInference([d_1449_v13_])
            index212_ = default__.safeIndex(544, (d_1454_v18_).length(0))
            (d_1454_v18_)[index212_] = (d_1455_v19_) + (d_1455_v19_)
            index213_ = default__.safeIndex(544, (d_1454_v18_).length(0))
            (d_1454_v18_)[index213_] = ((d_1455_v19_) + (d_1455_v19_)) + (((d_1455_v19_).set(default__.safeIndex(d_1446_v10_, len(d_1455_v19_)), d_1449_v13_)).set(default__.safeIndex(d_1446_v10_, len((d_1455_v19_).set(default__.safeIndex(d_1446_v10_, len(d_1455_v19_)), d_1449_v13_))), d_1449_v13_))
        d_1456_v20_: int
        d_1456_v20_ = -221
        d_1457_v21_: _dafny.MultiSet
        d_1457_v21_ = _dafny.MultiSet([(0) - ((0) - (d_1456_v20_))])
        d_1458_v22_: _dafny.Map
        d_1458_v22_ = _dafny.Map({(p0 if p0 else True): d_1457_v21_})
        d_1459_v23_: _dafny.Seq
        d_1459_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwsksvh"))
        d_1460_v24_: _dafny.Set
        d_1460_v24_ = _dafny.Set({d_1459_v23_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmpjo")), (d_1459_v23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "luh"))) + (d_1459_v23_)})
        rhs231_ = _dafny.Map({p0: _dafny.MultiSet([d_1456_v20_, 307])})
        rhs232_ = p0
        rhs233_ = p0
        rhs234_ = d_1460_v24_
        rhs235_ = (d_1456_v20_) - ((d_1456_v20_) * (d_1456_v20_))
        lhs206_ = globalState
        lhs207_ = globalState
        lhs208_ = globalState
        d_1458_v22_ = rhs231_
        lhs206_.f25 = rhs232_
        lhs207_.f9 = rhs233_
        d_1460_v24_ = rhs234_
        lhs208_.f24 = rhs235_
        d_1461_v25_: _dafny.Array
        nw215_ = _dafny.Array(int(0), 7)
        d_1461_v25_ = nw215_
        d_1462_v26_: _dafny.MultiSet
        d_1462_v26_ = _dafny.MultiSet([p0, True])
        index214_ = default__.safeIndex(233, (d_1461_v25_).length(0))
        (d_1461_v25_)[index214_] = (d_1462_v26_).cardinality
        index215_ = default__.safeIndex(233, (d_1461_v25_).length(0))
        (d_1461_v25_)[index215_] = d_1456_v20_
        d_1463_v27_: _dafny.Map
        d_1463_v27_ = _dafny.Map({False: d_1459_v23_})
        d_1464_v28_: D0
        d_1464_v28_ = D0_DC1(d_1459_v23_, (d_1457_v21_).cardinality, (d_1461_v25_)[default__.safeIndex(233, (d_1461_v25_).length(0))])
        d_1465_v29_: _dafny.Array
        nw216_ = _dafny.Array(None, 13)
        nw216_[int(0)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yes"))) + (d_1459_v23_)
        nw216_[int(1)] = ((d_1463_v27_)[p0] if (p0) in (d_1463_v27_) else d_1459_v23_)
        nw216_[int(2)] = d_1459_v23_
        nw216_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1466_i1_ in range(default__.abs(385))])
        nw216_[int(4)] = d_1459_v23_
        nw216_[int(5)] = (d_1459_v23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vra")))
        nw216_[int(6)] = (d_1464_v28_).cf1
        nw216_[int(7)] = d_1459_v23_
        nw216_[int(8)] = d_1459_v23_
        nw216_[int(9)] = d_1459_v23_
        nw216_[int(10)] = (d_1459_v23_) + (d_1459_v23_)
        nw216_[int(11)] = d_1459_v23_
        nw216_[int(12)] = d_1459_v23_
        d_1465_v29_ = nw216_
        index216_ = default__.safeIndex(658, (d_1465_v29_).length(0))
        (d_1465_v29_)[index216_] = d_1459_v23_
        d_1467_v30_: str
        d_1467_v30_ = _dafny.CodePoint('b')
        d_1468_v31_: _dafny.Seq
        d_1468_v31_ = _dafny.SeqWithoutIsStrInference([(d_1459_v23_).set(default__.safeIndex(len(d_1459_v23_), len(d_1459_v23_)), default__.fm20(p0, d_1467_v30_, len((self).f33), d_1467_v30_, globalState)), (d_1459_v23_) + (d_1459_v23_)])
        index217_ = default__.safeIndex(658, (d_1465_v29_).length(0))
        (d_1465_v29_)[index217_] = (d_1468_v31_)[default__.safeIndex(d_1456_v20_, len(d_1468_v31_))]
        d_1469_v32_: _dafny.Set
        d_1469_v32_ = _dafny.Set({p0})
        d_1470_v33_: D1
        d_1470_v33_ = D1_DC5(d_1469_v32_)
        pat_let_tv23_ = p0
        pat_let_tv24_ = p0
        pat_let_tv25_ = p0
        d_1471_v34_: _dafny.Map
        def iife117_(_pat_let23_0):
            def iife118_(d_1472_dt__update__tmp_h0_):
                def iife119_(_pat_let24_0):
                    def iife120_(d_1473_dt__update_hcf8_h0_):
                        return D1_DC5(d_1473_dt__update_hcf8_h0_)
                    return iife120_(_pat_let24_0)
                return iife119_(_dafny.Set({not(pat_let_tv23_), pat_let_tv24_, pat_let_tv25_, not(not(True))}))
            return iife118_(_pat_let23_0)
        d_1471_v34_ = _dafny.Map({len(default__.fm36(not(p0), default__.fm24(default__.fm24(p0, False, p0, (d_1461_v25_)[default__.safeIndex(233, (d_1461_v25_).length(0))], globalState), p0, p0, d_1456_v20_, globalState), not(p0), globalState)): iife117_(d_1470_v33_)})
        d_1474_v35_: _dafny.Map
        d_1474_v35_ = d_1471_v34_
        d_1475_v36_: _dafny.Array
        nw217_ = _dafny.Array(_dafny.Array(None, 0), 9)
        d_1475_v36_ = nw217_
        index218_ = default__.safeIndex(658, (d_1465_v29_).length(0))
        rhs236_ = d_1474_v35_
        rhs237_ = d_1475_v36_
        rhs238_ = d_1456_v20_
        rhs239_ = (d_1468_v31_)[default__.safeIndex((d_1461_v25_)[default__.safeIndex(233, (d_1461_v25_).length(0))], len(d_1468_v31_))]
        lhs209_ = d_1465_v29_
        lhs210_ = default__.safeIndex(658, (d_1465_v29_).length(0))
        d_1474_v35_ = rhs236_
        r2 = rhs237_
        r0 = rhs238_
        lhs209_[lhs210_] = rhs239_
        d_1476_v37_: _dafny.Array
        def lambda67_(d_1477_p0_):
            def lambda68_(d_1478_i2_):
                return d_1477_p0_

            return lambda68_

        init37_ = lambda67_(p0)
        nw218_ = _dafny.Array(None, 3)
        for i0_37_ in range(nw218_.length(0)):
            nw218_[i0_37_] = init37_(i0_37_)
        d_1476_v37_ = nw218_
        (globalState).f1 = d_1476_v37_
        r0 = d_1456_v20_
        d_1479_v38_: D0
        d_1479_v38_ = D0_DC0(d_1456_v20_)
        d_1480_v39_: _dafny.MultiSet
        d_1480_v39_ = _dafny.MultiSet([d_1479_v38_])
        r1 = (d_1480_v39_).intersection(d_1480_v39_)
        r2 = d_1475_v36_
        return r0, r1, r2

    def m5(self, p0, p1, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: _dafny.Map = _dafny.Map({})
        r2: str = _dafny.CodePoint('D')
        d_1481_v0_: bool
        d_1481_v0_ = True
        (globalState).f25 = d_1481_v0_
        d_1482_v1_: int
        d_1482_v1_ = 132
        if (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1483_i0_ in range(default__.abs(723))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqcc"))))) != ((994 if d_1481_v0_ else d_1482_v1_)):
            d_1484_v2_: _dafny.MultiSet
            d_1484_v2_ = _dafny.MultiSet([d_1482_v1_, d_1482_v1_])
            d_1485_v3_: _dafny.Seq
            d_1485_v3_ = _dafny.SeqWithoutIsStrInference([d_1481_v0_, d_1481_v0_, True, d_1481_v0_])
            d_1486_v4_: _dafny.Map
            d_1486_v4_ = _dafny.Map({d_1484_v2_: d_1485_v3_})
            d_1487_v5_: str
            d_1487_v5_ = _dafny.CodePoint('c')
            d_1488_v6_: _dafny.Map
            d_1488_v6_ = _dafny.Map({d_1481_v0_: default__.fm33(d_1487_v5_, True, d_1481_v0_, globalState)})
            d_1489_v7_: _dafny.Map
            d_1489_v7_ = _dafny.Map({d_1482_v1_: d_1488_v6_})
            d_1490_v8_: _dafny.MultiSet
            d_1490_v8_ = _dafny.MultiSet([d_1487_v5_])
            (globalState).f0 = (default__.fm32(d_1481_v0_, d_1486_v4_, (((d_1489_v7_)[((d_1490_v8_)[d_1487_v5_] if (d_1487_v5_) in (d_1490_v8_) else d_1482_v1_)] if (((d_1490_v8_)[d_1487_v5_] if (d_1487_v5_) in (d_1490_v8_) else d_1482_v1_)) in (d_1489_v7_) else d_1488_v6_)).set(d_1481_v0_, _dafny.MultiSet([d_1482_v1_])), d_1487_v5_, globalState)).cf19
            index219_ = default__.safeIndex(223, (p0).length(0))
            (p0)[index219_] = (d_1481_v0_) == (d_1481_v0_)
            d_1491_v9_: _dafny.Seq
            d_1491_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erhvfc"))
            d_1492_v10_: _dafny.Array
            nw219_ = _dafny.Array(None, 9)
            nw219_[int(0)] = (d_1491_v9_) + (d_1491_v9_)
            nw219_[int(1)] = d_1491_v9_
            nw219_[int(2)] = d_1491_v9_
            nw219_[int(3)] = d_1491_v9_
            nw219_[int(4)] = (d_1491_v9_).set(default__.safeIndex(-755, len(d_1491_v9_)), d_1487_v5_)
            nw219_[int(5)] = d_1491_v9_
            nw219_[int(6)] = d_1491_v9_
            nw219_[int(7)] = d_1491_v9_
            nw219_[int(8)] = (d_1491_v9_) + (d_1491_v9_)
            d_1492_v10_ = nw219_
            index220_ = default__.safeIndex(223, (p0).length(0))
            rhs240_ = (0) - (d_1482_v1_)
            rhs241_ = False
            rhs242_ = d_1492_v10_
            lhs211_ = globalState
            lhs212_ = p0
            lhs213_ = default__.safeIndex(223, (p0).length(0))
            lhs211_.f24 = rhs240_
            lhs212_[lhs213_] = rhs241_
            d_1492_v10_ = rhs242_
            d_1493_v11_: D0
            d_1493_v11_ = D0_DC0(d_1482_v1_)
            (globalState).f25 = (d_1482_v1_) <= ((d_1493_v11_).cf0)
            source28_ = D0_DC2(not(d_1481_v0_))
            if source28_.is_DC0:
                d_1494___mcc_h0_ = source28_.cf0
                d_1495_cf0_ = d_1494___mcc_h0_
                d_1496_v12_: _dafny.Array
                nw220_ = _dafny.Array(False, 22)
                d_1496_v12_ = nw220_
                (globalState).f1 = d_1496_v12_
                d_1497_v13_: _dafny.Set
                d_1497_v13_ = _dafny.Set({d_1481_v0_, d_1481_v0_})
                d_1498_v14_: D1
                d_1498_v14_ = D1_DC5(d_1497_v13_)
                d_1499_v15_: _dafny.Map
                d_1499_v15_ = _dafny.Map({d_1495_cf0_: d_1498_v14_})
                d_1500_v16_: _dafny.Map
                d_1500_v16_ = d_1499_v15_
                d_1501_v17_: _dafny.Map
                d_1501_v17_ = _dafny.Map({d_1500_v16_: d_1481_v0_})
                rhs243_ = d_1501_v17_
                rhs244_ = 542
                lhs214_ = globalState
                d_1501_v17_ = rhs243_
                lhs214_.f17 = rhs244_
                rhs245_ = d_1495_cf0_
                rhs246_ = not(d_1481_v0_)
                rhs247_ = (0) - (d_1482_v1_)
                lhs215_ = globalState
                lhs216_ = globalState
                lhs217_ = globalState
                lhs215_.f8 = rhs245_
                lhs216_.f9 = rhs246_
                lhs217_.f24 = rhs247_
                d_1502_v18_: _dafny.Array
                nw221_ = _dafny.Array(None, 11)
                d_1502_v18_ = nw221_
                d_1503_v19_: _dafny.Map
                d_1503_v19_ = _dafny.Map({d_1502_v18_: d_1482_v1_})
                d_1503_v19_ = (d_1503_v19_) | (((d_1503_v19_).set(d_1502_v18_, d_1482_v1_)) | (d_1503_v19_))
            elif source28_.is_DC1:
                d_1504___mcc_h1_ = source28_.cf1
                d_1505___mcc_h2_ = source28_.cf2
                d_1506___mcc_h3_ = source28_.cf3
                d_1507_cf3_ = d_1506___mcc_h3_
                d_1508_cf2_ = d_1505___mcc_h2_
                d_1509_cf1_ = d_1504___mcc_h1_
                d_1487_v5_ = d_1487_v5_
                d_1510_v20_: _dafny.Set
                d_1510_v20_ = _dafny.Set({d_1507_cf3_})
                r0 = (d_1510_v20_) | ((d_1510_v20_).intersection(d_1510_v20_))
                index221_ = default__.safeIndex(223, (p0).length(0))
                (p0)[index221_] = (d_1508_cf2_) != (d_1482_v1_)
                d_1511_v21_: _dafny.Map
                d_1511_v21_ = _dafny.Map({(d_1510_v20_).issubset(d_1510_v20_): (p0)[default__.safeIndex(223, (p0).length(0))]})
                d_1511_v21_ = (d_1511_v21_).set((p0)[default__.safeIndex(223, (p0).length(0))], d_1481_v0_)
            elif source28_.is_DC2:
                d_1512___mcc_h4_ = source28_.cf4
                d_1513_cf4_ = d_1512___mcc_h4_
                d_1514_v22_: C1
                nw222_ = C1()
                nw222_.ctor__()
                d_1514_v22_ = nw222_
                (globalState).f0 = d_1482_v1_
                d_1515_v23_: C1
                nw223_ = C1()
                nw223_.ctor__()
                d_1515_v23_ = nw223_
                (globalState).f17 = (0) - (d_1482_v1_)
            elif source28_.is_DC3:
                d_1516___mcc_h5_ = source28_.cf5
                d_1517___mcc_h6_ = source28_.cf6
                d_1518_cf6_ = d_1517___mcc_h6_
                d_1519_cf5_ = d_1516___mcc_h5_
                (globalState).f24 = default__.safeModuloInt(d_1482_v1_, d_1519_cf5_)
                d_1520_v24_: _dafny.Map
                d_1520_v24_ = _dafny.Map({(d_1482_v1_ if (d_1485_v3_)[default__.safeIndex(d_1482_v1_, len(d_1485_v3_))] else d_1482_v1_): (((self).f33).set(default__.safeIndex(len(p1), len((self).f33)), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('j')])))) == ((self).f33)})
                def iife121_():
                    coll71_ = _dafny.Map()
                    compr_71_: int
                    for compr_71_ in _dafny.IntegerRange(53, 864):
                        d_1521_v25_: int = compr_71_
                        if ((53) <= (d_1521_v25_)) and ((d_1521_v25_) < (864)):
                            coll71_[(d_1521_v25_) + (775)] = (p0)[default__.safeIndex(223, (p0).length(0))]
                    return _dafny.Map(coll71_)
                d_1520_v24_ = ((d_1520_v24_).set(d_1482_v1_, (p0)[default__.safeIndex(223, (p0).length(0))])) | (iife121_()
                )
                (globalState).f5 = d_1482_v1_
                d_1522_v26_: _dafny.Array
                nw224_ = _dafny.Array(False, 27)
                d_1522_v26_ = nw224_
                d_1523_v27_: _dafny.Map
                d_1523_v27_ = _dafny.Map({d_1522_v26_: (d_1518_cf6_) != (d_1518_cf6_)})
                d_1524_v28_: D13
                d_1524_v28_ = D13_DC39(d_1523_v27_)
                d_1523_v27_ = (d_1524_v28_).cf75
            elif True:
                d_1525___mcc_h7_ = source28_.cf7
                d_1526_cf7_ = d_1525___mcc_h7_
                index222_ = default__.safeIndex(223, (p0).length(0))
                (p0)[index222_] = (p0)[default__.safeIndex(223, (p0).length(0))]
                index223_ = default__.safeIndex(223, (p0).length(0))
                (p0)[index223_] = d_1481_v0_
                d_1527_v29_: _dafny.Set
                d_1527_v29_ = _dafny.Set({d_1481_v0_})
                index224_ = default__.safeIndex(223, (p0).length(0))
                rhs248_ = d_1481_v0_
                rhs249_ = d_1482_v1_
                rhs250_ = (default__.fm0(False, d_1491_v9_, d_1527_v29_, d_1482_v1_, globalState)) != ((d_1485_v3_).set(default__.safeIndex(len((self).f33), len(d_1485_v3_)), d_1481_v0_))
                lhs218_ = globalState
                lhs219_ = globalState
                lhs220_ = p0
                lhs221_ = default__.safeIndex(223, (p0).length(0))
                lhs218_.f25 = rhs248_
                lhs219_.f17 = rhs249_
                lhs220_[lhs221_] = rhs250_
                d_1528_v30_: _dafny.Array
                nw225_ = _dafny.Array(_dafny.Map({}), 26)
                d_1528_v30_ = nw225_
                d_1529_v31_: _dafny.Map
                d_1529_v31_ = _dafny.Map({(0) - (d_1482_v1_): not((p0)[default__.safeIndex(223, (p0).length(0))])})
                index225_ = default__.safeIndex(556, (d_1528_v30_).length(0))
                (d_1528_v30_)[index225_] = d_1529_v31_
                index226_ = default__.safeIndex(556, (d_1528_v30_).length(0))
                (d_1528_v30_)[index226_] = (d_1529_v31_) | (_dafny.Map({len(d_1491_v9_): d_1481_v0_}))
            index227_ = default__.safeIndex(223, (p0).length(0))
            (p0)[index227_] = d_1481_v0_
        elif True:
            d_1530_v32_: _dafny.Seq
            d_1530_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "updu"))
            d_1531_v33_: _dafny.Set
            d_1531_v33_ = _dafny.Set({d_1530_v32_})
            d_1531_v33_ = d_1531_v33_
            d_1532_v34_: _dafny.Seq
            d_1532_v34_ = _dafny.SeqWithoutIsStrInference([d_1481_v0_])
            d_1532_v34_ = ((_dafny.SeqWithoutIsStrInference([d_1481_v0_, d_1481_v0_])).set(default__.safeIndex(d_1482_v1_, len(_dafny.SeqWithoutIsStrInference([d_1481_v0_, d_1481_v0_]))), d_1481_v0_)) + ((d_1532_v34_) + (d_1532_v34_))
            (globalState).f5 = d_1482_v1_
            d_1533_v35_: _dafny.Array
            nw226_ = _dafny.Array(_dafny.Seq({}), 10)
            d_1533_v35_ = nw226_
            d_1534_v36_: _dafny.Seq
            d_1534_v36_ = _dafny.SeqWithoutIsStrInference([p0])
            index228_ = default__.safeIndex(937, (d_1533_v35_).length(0))
            (d_1533_v35_)[index228_] = d_1534_v36_
            index229_ = default__.safeIndex(937, (d_1533_v35_).length(0))
            (d_1533_v35_)[index229_] = d_1534_v36_
            d_1535_v37_: _dafny.Seq
            d_1535_v37_ = _dafny.SeqWithoutIsStrInference([d_1482_v1_])
            d_1535_v37_ = _dafny.SeqWithoutIsStrInference([(len(_dafny.Map({d_1481_v0_: not(d_1481_v0_)}))) * (d_1482_v1_), d_1482_v1_])
        d_1536_v38_: _dafny.Set
        d_1536_v38_ = _dafny.Set({((self).f31) != ((self).f31)})
        d_1536_v38_ = _dafny.Set({d_1481_v0_})
        if not (d_1481_v0_) or ((d_1481_v0_) and (d_1481_v0_)):
            d_1537_v39_: _dafny.MultiSet
            d_1537_v39_ = _dafny.MultiSet([d_1481_v0_])
            (globalState).f25 = default__.fm24((d_1537_v39_).isdisjoint((d_1537_v39_).set(d_1481_v0_, default__.abs(d_1482_v1_))), d_1481_v0_, d_1481_v0_, d_1482_v1_, globalState)
            d_1538_v40_: _dafny.Map
            d_1538_v40_ = _dafny.Map({p0: d_1481_v0_})
            d_1539_v41_: _dafny.Seq
            d_1539_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvgw"))
            d_1540_v42_: bool
            d_1541_v43_: int
            d_1542_v44_: bool
            out37_: bool
            out38_: int
            out39_: bool
            out37_, out38_, out39_ = (self).m6(not(d_1481_v0_), d_1538_v40_, d_1539_v41_, not(True), globalState)
            d_1540_v42_ = out37_
            d_1541_v43_ = out38_
            d_1542_v44_ = out39_
            d_1543_v45_: _dafny.Map
            d_1543_v45_ = _dafny.Map({476: d_1481_v0_})
            d_1544_v46_: _dafny.Seq
            d_1544_v46_ = _dafny.SeqWithoutIsStrInference([d_1540_v42_, d_1481_v0_])
            d_1545_v47_: D2
            d_1545_v47_ = D2_DC8(True, d_1482_v1_, d_1482_v1_)
            d_1546_v48_: _dafny.Array
            nw227_ = _dafny.Array(None, 9)
            nw227_[int(0)] = (d_1541_v43_) * ((0) - ((0) - ((self).fm6((D8_DC22(d_1543_v45_)).cf41, d_1482_v1_, (self).f33, _dafny.Map({d_1544_v46_: (d_1544_v46_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1547_i1_ in range(default__.abs(-610))])), len(d_1544_v46_))]}), globalState))))
            nw227_[int(1)] = d_1541_v43_
            nw227_[int(2)] = d_1541_v43_
            nw227_[int(3)] = default__.safeDivisionInt(d_1482_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayp"))))
            nw227_[int(4)] = (0) - (default__.safeDivisionInt(d_1541_v43_, d_1482_v1_))
            nw227_[int(5)] = d_1541_v43_
            nw227_[int(6)] = d_1541_v43_
            nw227_[int(7)] = ((d_1545_v47_).cf14 if d_1540_v42_ else d_1482_v1_)
            nw227_[int(8)] = d_1541_v43_
            d_1546_v48_ = nw227_
            index230_ = default__.safeIndex(518, (d_1546_v48_).length(0))
            (d_1546_v48_)[index230_] = d_1541_v43_
            index231_ = default__.safeIndex(518, (d_1546_v48_).length(0))
            (d_1546_v48_)[index231_] = d_1541_v43_
            d_1548_v49_: str
            d_1548_v49_ = _dafny.CodePoint('h')
            d_1549_v50_: D1
            d_1549_v50_ = D1_DC5(d_1536_v38_)
            d_1550_v51_: _dafny.Map
            d_1550_v51_ = _dafny.Map({d_1482_v1_: d_1549_v50_})
            d_1551_v52_: D8
            d_1551_v52_ = D8_DC25(False, p1, d_1548_v49_, d_1550_v51_, d_1541_v43_)
            d_1552_v53_: _dafny.Map
            d_1552_v53_ = _dafny.Map({d_1539_v41_: d_1551_v52_})
            d_1553_v55_: _dafny.Map
            d_1553_v55_ = _dafny.Map({d_1539_v41_: d_1542_v44_})
            def iife122_():
                coll72_ = _dafny.Map()
                compr_72_: _dafny.Seq
                for compr_72_ in (d_1553_v55_).keys.Elements:
                    d_1554_v54_: _dafny.Seq = compr_72_
                    if (d_1554_v54_) in (d_1553_v55_):
                        coll72_[d_1554_v54_] = d_1551_v52_
                return _dafny.Map(coll72_)
            d_1552_v53_ = (iife122_()
            ) | ((d_1552_v53_) | (_dafny.Map({d_1539_v41_: d_1551_v52_})))
            (globalState).f25 = (d_1539_v41_) <= (d_1539_v41_)
        elif True:
            d_1555_v56_: str
            d_1555_v56_ = _dafny.CodePoint('j')
            r2 = d_1555_v56_
            d_1556_v57_: D10
            d_1556_v57_ = D10_DC32(d_1481_v0_, d_1482_v1_, d_1481_v0_)
            d_1557_v58_: _dafny.MultiSet
            d_1557_v58_ = _dafny.MultiSet([d_1481_v0_, d_1481_v0_, (d_1556_v57_).cf62, d_1481_v0_])
            d_1558_v59_: _dafny.Map
            d_1558_v59_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1482_v1_ for d_1559_i2_ in range(default__.abs(415))]): 602})
            if default__.fm24(d_1481_v0_, (((self).f33)[default__.safeIndex((d_1557_v58_).cardinality, len((self).f33))]) <= ((self).fm7(d_1482_v1_, d_1481_v0_, globalState)), d_1481_v0_, (len(d_1558_v59_) if default__.fm24(d_1481_v0_, d_1481_v0_, d_1481_v0_, d_1482_v1_, globalState) else (0) - (d_1482_v1_)), globalState):
                d_1560_v60_: _dafny.Array
                nw228_ = _dafny.Array(_dafny.MultiSet({}), 10)
                d_1560_v60_ = nw228_
                index232_ = default__.safeIndex(644, (d_1560_v60_).length(0))
                (d_1560_v60_)[index232_] = _dafny.MultiSet([d_1481_v0_, d_1481_v0_])
                index233_ = default__.safeIndex(644, (d_1560_v60_).length(0))
                (d_1560_v60_)[index233_] = d_1557_v58_
                (globalState).f1 = p0
                (globalState).f25 = d_1481_v0_
                d_1561_v61_: _dafny.Map
                d_1561_v61_ = _dafny.Map({len(_dafny.Map({d_1482_v1_: d_1482_v1_})): d_1481_v0_})
                d_1562_v62_: _dafny.Map
                d_1562_v62_ = _dafny.Map({d_1555_v56_: False})
                d_1563_v64_: _dafny.MultiSet
                d_1563_v64_ = _dafny.MultiSet([d_1555_v56_])
                d_1564_v65_: _dafny.Seq
                d_1564_v65_ = _dafny.SeqWithoutIsStrInference([d_1562_v62_])
                d_1565_v67_: _dafny.Array
                nw229_ = _dafny.Array(None, 24)
                nw229_[int(0)] = d_1562_v62_
                nw229_[int(1)] = _dafny.Map({d_1555_v56_: d_1481_v0_})
                nw229_[int(2)] = d_1562_v62_
                nw229_[int(3)] = d_1562_v62_
                def iife123_():
                    coll73_ = _dafny.Map()
                    compr_73_: str
                    for compr_73_ in (d_1563_v64_).Elements:
                        d_1566_v63_: str = compr_73_
                        if (d_1566_v63_) in (d_1563_v64_):
                            coll73_[d_1566_v63_] = d_1481_v0_
                    return _dafny.Map(coll73_)
                nw229_[int(4)] = iife123_()
                
                nw229_[int(5)] = d_1562_v62_
                nw229_[int(6)] = d_1562_v62_
                nw229_[int(7)] = d_1562_v62_
                nw229_[int(8)] = d_1562_v62_
                nw229_[int(9)] = d_1562_v62_
                nw229_[int(10)] = d_1562_v62_
                nw229_[int(11)] = d_1562_v62_
                nw229_[int(12)] = d_1562_v62_
                nw229_[int(13)] = (d_1562_v62_).set(d_1555_v56_, d_1481_v0_)
                nw229_[int(14)] = (d_1564_v65_)[default__.safeIndex(d_1482_v1_, len(d_1564_v65_))]
                nw229_[int(15)] = d_1562_v62_
                nw229_[int(16)] = d_1562_v62_
                nw229_[int(17)] = d_1562_v62_
                nw229_[int(18)] = _dafny.Map({d_1555_v56_: d_1481_v0_})
                nw229_[int(19)] = d_1562_v62_
                nw229_[int(20)] = _dafny.Map({d_1555_v56_: d_1481_v0_})
                nw229_[int(21)] = d_1562_v62_
                def iife124_():
                    coll74_ = _dafny.Map()
                    compr_74_: str
                    for compr_74_ in (d_1562_v62_).keys.Elements:
                        d_1567_v66_: str = compr_74_
                        if (d_1567_v66_) in (d_1562_v62_):
                            coll74_[d_1567_v66_] = not(d_1481_v0_)
                    return _dafny.Map(coll74_)
                nw229_[int(22)] = iife124_()
                
                nw229_[int(23)] = d_1562_v62_
                d_1565_v67_ = nw229_
                d_1568_v68_: _dafny.Map
                d_1568_v68_ = _dafny.Map({d_1561_v61_: d_1565_v67_})
                d_1568_v68_ = (d_1568_v68_).set(d_1561_v61_, d_1565_v67_)
                (globalState).f5 = (((self).f31)[(self).fm7(742, d_1481_v0_, globalState)] if ((self).fm7(742, d_1481_v0_, globalState)) in ((self).f31) else d_1482_v1_)
            elif True:
                d_1569_v69_: _dafny.Array
                nw230_ = _dafny.Array(_dafny.Map({}), 14)
                d_1569_v69_ = nw230_
                d_1570_v70_: _dafny.Map
                d_1570_v70_ = _dafny.Map({d_1481_v0_: d_1481_v0_})
                d_1571_v71_: D3
                d_1571_v71_ = D3_DC10(d_1570_v70_)
                index234_ = default__.safeIndex(655, (d_1569_v69_).length(0))
                (d_1569_v69_)[index234_] = _dafny.Map({d_1571_v71_: d_1482_v1_})
                d_1572_v72_: _dafny.Map
                d_1572_v72_ = _dafny.Map({d_1571_v71_: (0) - ((_dafny.MultiSet([d_1482_v1_])).cardinality)})
                d_1573_v73_: _dafny.Map
                d_1573_v73_ = _dafny.Map({d_1482_v1_: d_1481_v0_})
                d_1574_v74_: D2
                d_1574_v74_ = D2_DC9(d_1482_v1_, default__.fm24(not(not(d_1481_v0_)), d_1481_v0_, ((d_1573_v73_)[d_1482_v1_] if (d_1482_v1_) in (d_1573_v73_) else d_1481_v0_), d_1482_v1_, globalState), d_1482_v1_, d_1482_v1_)
                d_1575_v75_: D0
                d_1575_v75_ = D0_DC3(d_1482_v1_, len(_dafny.SeqWithoutIsStrInference([d_1482_v1_ for d_1576_i3_ in range(default__.abs(-734))])))
                index235_ = default__.safeIndex(655, (d_1569_v69_).length(0))
                (d_1569_v69_)[index235_] = ((d_1572_v72_).set(d_1571_v71_, d_1482_v1_)).set(default__.fm19(d_1574_v74_, d_1575_v75_, d_1555_v56_, d_1481_v0_, globalState), d_1482_v1_)
                d_1577_v76_: C1
                nw231_ = C1()
                nw231_.ctor__()
                d_1577_v76_ = nw231_
                d_1578_v77_: _dafny.Array
                nw232_ = _dafny.Array(int(0), 4)
                d_1578_v77_ = nw232_
                rhs251_ = d_1482_v1_
                rhs252_ = d_1578_v77_
                lhs222_ = globalState
                lhs222_.f8 = rhs251_
                d_1578_v77_ = rhs252_
                d_1579_v78_: _dafny.Array
                nw233_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_1579_v78_ = nw233_
                d_1580_v79_: _dafny.Map
                d_1580_v79_ = default__.fm39(d_1555_v56_, True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbydpr"))), d_1481_v0_, globalState)
                d_1581_v80_: _dafny.Map
                d_1581_v80_ = _dafny.Map({d_1579_v78_: d_1580_v79_})
                d_1581_v80_ = (d_1581_v80_).set(d_1579_v78_, d_1580_v79_)
                d_1582_v81_: _dafny.Seq
                d_1582_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpalrcg"))
                (globalState).f9 = ((default__.fm20(d_1481_v0_, _dafny.CodePoint('j'), len(d_1582_v81_), d_1555_v56_, globalState) if False else d_1555_v56_)) not in (d_1582_v81_)
            d_1481_v0_ = (d_1482_v1_) <= ((self).fm7(d_1482_v1_, d_1481_v0_, globalState))
            (globalState).f1 = p0
            d_1583_v82_: _dafny.Array
            nw234_ = _dafny.Array(int(0), 27)
            d_1583_v82_ = nw234_
            d_1584_v83_: _dafny.Seq
            d_1584_v83_ = _dafny.SeqWithoutIsStrInference([d_1583_v82_])
            d_1584_v83_ = d_1584_v83_
        if (d_1482_v1_) >= (default__.safeModuloInt(d_1482_v1_, (0) - ((0) - (d_1482_v1_)))):
            d_1585_v84_: _dafny.Seq
            d_1585_v84_ = _dafny.SeqWithoutIsStrInference([d_1482_v1_])
            d_1585_v84_ = (self).f33
            (globalState).f5 = (d_1482_v1_) - (d_1482_v1_)
            def lambda69_(d_1586_v0_):
                def lambda70_(d_1587_i4_):
                    return d_1586_v0_

                return lambda70_

            init38_ = lambda69_(d_1481_v0_)
            nw235_ = _dafny.Array(None, 11)
            for i0_38_ in range(nw235_.length(0)):
                nw235_[i0_38_] = init38_(i0_38_)
            (globalState).f1 = nw235_
            index236_ = default__.safeIndex(197, (p0).length(0))
            (p0)[index236_] = d_1481_v0_
            index237_ = default__.safeIndex(197, (p0).length(0))
            (p0)[index237_] = True
            index238_ = default__.safeIndex(197, (p0).length(0))
            (p0)[index238_] = not(d_1481_v0_)
        elif True:
            (globalState).f0 = (0) - (d_1482_v1_)
            if d_1481_v0_:
                index239_ = default__.safeIndex(459, (p0).length(0))
                (p0)[index239_] = True
                index240_ = default__.safeIndex(459, (p0).length(0))
                (p0)[index240_] = d_1481_v0_
                d_1588_v85_: _dafny.Map
                d_1588_v85_ = _dafny.Map({len((self).f33): (p0)[default__.safeIndex(459, (p0).length(0))]})
                d_1589_v86_: _dafny.Seq
                d_1589_v86_ = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(459, (p0).length(0))], (p0)[default__.safeIndex(459, (p0).length(0))]])
                d_1590_v87_: _dafny.Map
                d_1590_v87_ = _dafny.Map({d_1589_v86_: (p0)[default__.safeIndex(459, (p0).length(0))]})
                d_1591_v88_: _dafny.Map
                d_1591_v88_ = _dafny.Map({(p0)[default__.safeIndex(459, (p0).length(0))]: d_1481_v0_})
                d_1592_v89_: _dafny.Set
                d_1592_v89_ = _dafny.Set({len(d_1591_v88_), 212})
                d_1593_v90_: _dafny.Array
                nw236_ = _dafny.Array(False, 2)
                d_1593_v90_ = nw236_
                d_1594_v91_: _dafny.Map
                d_1594_v91_ = _dafny.Map({d_1593_v90_: d_1481_v0_})
                d_1595_v92_: _dafny.Seq
                d_1595_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ixqyqnbi"))
                d_1596_v93_: _dafny.MultiSet
                d_1596_v93_ = _dafny.MultiSet([d_1482_v1_, d_1482_v1_, d_1482_v1_])
                d_1597_v94_: str
                d_1597_v94_ = _dafny.CodePoint('t')
                d_1598_v95_: bool
                d_1599_v96_: int
                d_1600_v97_: bool
                out40_: bool
                out41_: int
                out42_: bool
                out40_, out41_, out42_ = (self).m6(((self).fm6(d_1588_v85_, d_1482_v1_, (self).f33, d_1590_v87_, globalState)) > (len(d_1592_v89_)), d_1594_v91_, ((d_1595_v92_) + (d_1595_v92_)).set(default__.safeIndex((d_1596_v93_).cardinality, len((d_1595_v92_) + (d_1595_v92_))), d_1597_v94_), (((d_1588_v85_)[d_1482_v1_] if (d_1482_v1_) in (d_1588_v85_) else (p0)[default__.safeIndex(459, (p0).length(0))])) or (d_1481_v0_), globalState)
                d_1598_v95_ = out40_
                d_1599_v96_ = out41_
                d_1600_v97_ = out42_
                d_1600_v97_ = (d_1536_v38_).ispropersubset((default__.fm17(globalState)) - (d_1536_v38_))
                d_1601_v98_: D1
                d_1601_v98_ = D1_DC5(d_1536_v38_)
                d_1602_v99_: _dafny.Map
                d_1602_v99_ = _dafny.Map({d_1482_v1_: d_1601_v98_})
                d_1603_v100_: _dafny.Map
                d_1603_v100_ = d_1602_v99_
                d_1604_v101_: D8
                d_1604_v101_ = D8_DC25(False, p1, d_1597_v94_, d_1603_v100_, d_1482_v1_)
                d_1605_v102_: _dafny.Seq
                d_1605_v102_ = _dafny.SeqWithoutIsStrInference([D8_DC25((p0)[default__.safeIndex(459, (p0).length(0))], p1, _dafny.CodePoint('f'), d_1603_v100_, d_1599_v96_), d_1604_v101_])
                d_1606_v103_: _dafny.Array
                nw237_ = _dafny.Array(None, 5)
                nw237_[int(0)] = _dafny.Map({(self).fm7(len(_dafny.Map({d_1482_v1_: d_1605_v102_})), (p0)[default__.safeIndex(459, (p0).length(0))], globalState): d_1599_v96_})
                nw237_[int(1)] = (self).f31
                nw237_[int(2)] = (self).f31
                nw237_[int(3)] = (self).f31
                nw237_[int(4)] = (self).f31
                d_1606_v103_ = nw237_
                d_1607_v104_: _dafny.Map
                d_1607_v104_ = _dafny.Map({d_1482_v1_: (self).f31})
                index241_ = default__.safeIndex(280, (d_1606_v103_).length(0))
                (d_1606_v103_)[index241_] = ((d_1607_v104_)[d_1599_v96_] if (d_1599_v96_) in (d_1607_v104_) else default__.fm40(d_1600_v97_, d_1482_v1_, d_1588_v85_, d_1600_v97_, globalState))
                index242_ = default__.safeIndex(280, (d_1606_v103_).length(0))
                def iife125_():
                    coll75_ = _dafny.Map()
                    compr_75_: int
                    for compr_75_ in (d_1596_v93_).Elements:
                        d_1608_v105_: int = compr_75_
                        if (d_1608_v105_) in (d_1596_v93_):
                            coll75_[(d_1608_v105_) + (d_1599_v96_)] = d_1598_v95_
                    return _dafny.Map(coll75_)
                (d_1606_v103_)[index242_] = (default__.fm40(d_1598_v95_, 141, iife125_()
                , (p0)[default__.safeIndex(459, (p0).length(0))], globalState)) | (_dafny.Map({len(d_1589_v86_): d_1599_v96_}))
                (globalState).f5 = 535
            elif True:
                (globalState).f17 = (0) - (d_1482_v1_)
                d_1609_v106_: _dafny.Seq
                d_1609_v106_ = _dafny.SeqWithoutIsStrInference([not (d_1481_v0_) or (True), d_1481_v0_, d_1481_v0_])
                d_1609_v106_ = d_1609_v106_
                d_1610_v107_: _dafny.MultiSet
                d_1610_v107_ = _dafny.MultiSet([not((d_1609_v106_)[default__.safeIndex(len(d_1609_v106_), len(d_1609_v106_))])])
                d_1611_v108_: _dafny.Array
                nw238_ = _dafny.Array(None, 21)
                nw238_[int(0)] = d_1482_v1_
                nw238_[int(1)] = d_1482_v1_
                nw238_[int(2)] = (0) - (d_1482_v1_)
                nw238_[int(3)] = d_1482_v1_
                nw238_[int(4)] = d_1482_v1_
                nw238_[int(5)] = d_1482_v1_
                nw238_[int(6)] = d_1482_v1_
                nw238_[int(7)] = d_1482_v1_
                nw238_[int(8)] = d_1482_v1_
                nw238_[int(9)] = d_1482_v1_
                nw238_[int(10)] = d_1482_v1_
                nw238_[int(11)] = d_1482_v1_
                nw238_[int(12)] = 923
                nw238_[int(13)] = d_1482_v1_
                nw238_[int(14)] = len((self).f33)
                nw238_[int(15)] = len(d_1609_v106_)
                nw238_[int(16)] = 342
                nw238_[int(17)] = d_1482_v1_
                nw238_[int(18)] = ((d_1610_v107_).set(False, default__.abs(d_1482_v1_))).cardinality
                nw238_[int(19)] = d_1482_v1_
                nw238_[int(20)] = d_1482_v1_
                d_1611_v108_ = nw238_
                d_1612_v109_: _dafny.Map
                d_1612_v109_ = _dafny.Map({len(d_1536_v38_): d_1611_v108_})
                d_1612_v109_ = (d_1612_v109_).set(d_1482_v1_, d_1611_v108_)
                d_1613_v110_: _dafny.Map
                d_1613_v110_ = _dafny.Map({len((self).f31): False})
                d_1614_v111_: _dafny.Map
                d_1614_v111_ = _dafny.Map({d_1609_v106_: ((d_1613_v110_)[-540] if (-540) in (d_1613_v110_) else d_1481_v0_)})
                index243_ = default__.safeIndex(682, (d_1611_v108_).length(0))
                (d_1611_v108_)[index243_] = default__.safeDivisionInt((self).fm6(default__.fm18(d_1482_v1_, globalState), d_1482_v1_, (self).f33, d_1614_v111_, globalState), (0) - (d_1482_v1_))
                index244_ = default__.safeIndex(682, (d_1611_v108_).length(0))
                (d_1611_v108_)[index244_] = len(d_1536_v38_)
                d_1615_v112_: C1
                nw239_ = C1()
                nw239_.ctor__()
                d_1615_v112_ = nw239_
            d_1616_v113_: _dafny.Array
            def lambda71_(d_1617_v0_):
                def lambda72_(d_1618_i5_):
                    return default__.safeModuloInt(d_1618_i5_, (_dafny.MultiSet([d_1617_v0_, d_1617_v0_])).cardinality)

                return lambda72_

            init39_ = lambda71_(d_1481_v0_)
            nw240_ = _dafny.Array(None, 5)
            for i0_39_ in range(nw240_.length(0)):
                nw240_[i0_39_] = init39_(i0_39_)
            d_1616_v113_ = nw240_
            rhs253_ = _dafny.CodePoint('t')
            rhs254_ = d_1481_v0_
            rhs255_ = d_1616_v113_
            r2 = rhs253_
            d_1481_v0_ = rhs254_
            d_1616_v113_ = rhs255_
            d_1619_v114_: _dafny.Seq
            d_1619_v114_ = _dafny.SeqWithoutIsStrInference([d_1481_v0_])
            d_1620_v115_: _dafny.Map
            d_1620_v115_ = _dafny.Map({d_1482_v1_: d_1619_v114_})
            d_1621_v116_: _dafny.MultiSet
            d_1621_v116_ = _dafny.MultiSet([d_1481_v0_])
            d_1620_v115_ = (d_1620_v115_).set((d_1621_v116_).cardinality, d_1619_v114_)
            index245_ = default__.safeIndex(851, (p0).length(0))
            (p0)[index245_] = d_1481_v0_
            index246_ = default__.safeIndex(851, (p0).length(0))
            (p0)[index246_] = not(d_1481_v0_)
        d_1622_v117_: _dafny.Array
        nw241_ = _dafny.Array(int(0), 13)
        d_1622_v117_ = nw241_
        index247_ = default__.safeIndex(353, (d_1622_v117_).length(0))
        (d_1622_v117_)[index247_] = d_1482_v1_
        index248_ = default__.safeIndex(714, (d_1622_v117_).length(0))
        (d_1622_v117_)[index248_] = d_1482_v1_
        d_1623_v118_: D0
        d_1623_v118_ = D0_DC3(d_1482_v1_, d_1482_v1_)
        d_1624_v121_: D0
        d_1624_v121_ = D0_DC2(d_1481_v0_)
        d_1625_v122_: _dafny.Seq
        d_1625_v122_ = _dafny.SeqWithoutIsStrInference([D0_DC2(d_1481_v0_), D0_DC2(d_1481_v0_), d_1624_v121_])
        pat_let_tv26_ = d_1481_v0_
        pat_let_tv27_ = d_1481_v0_
        pat_let_tv28_ = d_1481_v0_
        pat_let_tv29_ = d_1481_v0_
        pat_let_tv30_ = d_1481_v0_
        pat_let_tv31_ = d_1482_v1_
        pat_let_tv32_ = d_1482_v1_
        pat_let_tv33_ = d_1482_v1_
        pat_let_tv34_ = d_1481_v0_
        index249_ = default__.safeIndex(353, (d_1622_v117_).length(0))
        index250_ = default__.safeIndex(714, (d_1622_v117_).length(0))
        def lambda73_(source29_):
            if source29_.is_DC0:
                d_1626___mcc_h8_ = source29_.cf0
                d_1627_cf0_ = d_1626___mcc_h8_
                return (0) - (len((self).f33))
            elif source29_.is_DC1:
                d_1628___mcc_h9_ = source29_.cf1
                d_1629___mcc_h10_ = source29_.cf2
                d_1630___mcc_h11_ = source29_.cf3
                d_1631_cf3_ = d_1630___mcc_h11_
                d_1632_cf2_ = d_1629___mcc_h10_
                d_1633_cf1_ = d_1628___mcc_h9_
                def iife126_():
                    coll76_ = _dafny.Map()
                    compr_76_: _dafny.Set
                    for compr_76_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({605})])).Elements:
                        d_1634_v119_: _dafny.Set = compr_76_
                        if (d_1634_v119_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({605})])):
                            coll76_[d_1634_v119_] = not(not(pat_let_tv26_))
                    return _dafny.Map(coll76_)
                def iife127_():
                    coll77_ = _dafny.Map()
                    compr_77_: _dafny.Set
                    for compr_77_ in (_dafny.Map({_dafny.Set({135}): D10_DC32(pat_let_tv27_, d_1632_cf2_, pat_let_tv28_)})).keys.Elements:
                        d_1635_v120_: _dafny.Set = compr_77_
                        if (d_1635_v120_) in (_dafny.Map({_dafny.Set({135}): D10_DC32(pat_let_tv29_, d_1632_cf2_, pat_let_tv30_)})):
                            coll77_[d_1635_v120_] = True
                    return _dafny.Map(coll77_)
                return len((iife126_()
                ) | (iife127_()
                ))
            elif source29_.is_DC2:
                d_1636___mcc_h12_ = source29_.cf4
                d_1637_cf4_ = d_1636___mcc_h12_
                return (pat_let_tv31_) + ((0) - (pat_let_tv32_))
            elif source29_.is_DC3:
                d_1638___mcc_h13_ = source29_.cf5
                d_1639___mcc_h14_ = source29_.cf6
                d_1640_cf6_ = d_1639___mcc_h14_
                d_1641_cf5_ = d_1638___mcc_h13_
                return (0) - (d_1640_cf6_)
            elif True:
                d_1642___mcc_h15_ = source29_.cf7
                d_1643_cf7_ = d_1642___mcc_h15_
                return pat_let_tv33_

        def iife128_(_pat_let25_0):
            def iife129_(d_1644_dt__update__tmp_h0_):
                def iife130_(_pat_let26_0):
                    def iife131_(d_1645_dt__update_hcf4_h0_):
                        return D0_DC2(d_1645_dt__update_hcf4_h0_)
                    return iife131_(_pat_let26_0)
                return iife130_(pat_let_tv34_)
            return iife129_(_pat_let25_0)
        rhs256_ = d_1482_v1_
        rhs257_ = lambda73_(d_1623_v118_)
        rhs258_ = d_1482_v1_
        rhs259_ = (0) - (d_1482_v1_)
        rhs260_ = (iife128_(d_1624_v121_)) not in (d_1625_v122_)
        lhs223_ = d_1622_v117_
        lhs224_ = default__.safeIndex(353, (d_1622_v117_).length(0))
        lhs225_ = globalState
        lhs226_ = d_1622_v117_
        lhs227_ = default__.safeIndex(714, (d_1622_v117_).length(0))
        lhs228_ = globalState
        d_1482_v1_ = rhs256_
        lhs223_[lhs224_] = rhs257_
        lhs225_.f8 = rhs258_
        lhs226_[lhs227_] = rhs259_
        lhs228_.f9 = rhs260_
        d_1646_v123_: _dafny.Seq
        d_1646_v123_ = _dafny.SeqWithoutIsStrInference([True, d_1481_v0_])
        d_1647_v124_: _dafny.Set
        d_1647_v124_ = _dafny.Set({(d_1622_v117_)[default__.safeIndex(353, (d_1622_v117_).length(0))], len(d_1646_v123_), d_1482_v1_})
        r0 = ((d_1647_v124_).intersection(d_1647_v124_)) - ((d_1647_v124_ if d_1481_v0_ else d_1647_v124_))
        d_1648_v125_: _dafny.Map
        d_1648_v125_ = _dafny.Map({(_dafny.MultiSet([(d_1622_v117_)[default__.safeIndex(353, (d_1622_v117_).length(0))]])).set((d_1622_v117_)[default__.safeIndex(353, (d_1622_v117_).length(0))], default__.abs(-533)): len(default__.fm21(d_1482_v1_, d_1482_v1_, d_1481_v0_, default__.fm24(d_1481_v0_, d_1481_v0_, d_1481_v0_, d_1482_v1_, globalState), globalState))})
        r1 = d_1648_v125_
        d_1649_v126_: str
        d_1649_v126_ = _dafny.CodePoint('s')
        r2 = d_1649_v126_
        return r0, r1, r2

    def m6(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        d_1650_v0_: _dafny.Map
        d_1650_v0_ = _dafny.Map({not(p3): p3})
        d_1650_v0_ = (d_1650_v0_).set((p0 if p0 else p0), p0)
        d_1651_v1_: _dafny.Array
        def lambda74_(d_1652_i0_):
            return (d_1652_i0_) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nq")))))

        init40_ = lambda74_
        nw242_ = _dafny.Array(None, 9)
        for i0_40_ in range(nw242_.length(0)):
            nw242_[i0_40_] = init40_(i0_40_)
        d_1651_v1_ = nw242_
        d_1653_v2_: int
        d_1653_v2_ = 867
        index251_ = default__.safeIndex(737, (d_1651_v1_).length(0))
        (d_1651_v1_)[index251_] = d_1653_v2_
        index252_ = default__.safeIndex(737, (d_1651_v1_).length(0))
        (d_1651_v1_)[index252_] = d_1653_v2_
        d_1654_v3_: _dafny.Map
        d_1654_v3_ = _dafny.Map({(d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))]: default__.fm24(False, p3, p0, len(p2), globalState)})
        d_1655_v4_: _dafny.Seq
        d_1655_v4_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1656_v5_: _dafny.Map
        d_1656_v5_ = _dafny.Map({default__.fm0(p0, p2, default__.fm17(globalState), d_1653_v2_, globalState): p0})
        d_1653_v2_ = default__.safeDivisionInt(d_1653_v2_, (self).fm6(d_1654_v3_, len(d_1655_v4_), (self).f33, d_1656_v5_, globalState))
        hi6_ = -524
        for d_1657_i1_ in range(d_1653_v2_, hi6_):
            d_1658_v6_: C0
            nw243_ = C0()
            nw243_.ctor__()
            d_1658_v6_ = nw243_
            r0 = (d_1657_i1_) < (96)
            d_1659_v7_: _dafny.Array
            def lambda75_(d_1660_p0_):
                def lambda76_(d_1661_i2_):
                    return d_1660_p0_

                return lambda76_

            init41_ = lambda75_(p0)
            nw244_ = _dafny.Array(None, 21)
            for i0_41_ in range(nw244_.length(0)):
                nw244_[i0_41_] = init41_(i0_41_)
            d_1659_v7_ = nw244_
            (globalState).f1 = d_1659_v7_
            d_1655_v4_ = d_1655_v4_
        d_1651_v1_ = d_1651_v1_
        if p0:
            d_1662_v8_: _dafny.Array
            nw245_ = _dafny.Array(_dafny.Array(None, 0), 3)
            d_1662_v8_ = nw245_
            d_1663_v9_: _dafny.Array
            nw246_ = _dafny.Array(_dafny.MultiSet({}), 25)
            d_1663_v9_ = nw246_
            index253_ = default__.safeIndex(911, (d_1662_v8_).length(0))
            (d_1662_v8_)[index253_] = d_1663_v9_
            d_1664_v10_: _dafny.Array
            def lambda77_(d_1665_p0_, d_1666_v1_):
                def lambda78_(d_1667_i3_):
                    return (_dafny.Map({d_1665_p0_: (d_1666_v1_)[default__.safeIndex(737, (d_1666_v1_).length(0))]})) | (_dafny.Map({d_1665_p0_: (d_1666_v1_)[default__.safeIndex(737, (d_1666_v1_).length(0))]}))

                return lambda78_

            init42_ = lambda77_(p0, d_1651_v1_)
            nw247_ = _dafny.Array(None, 26)
            for i0_42_ in range(nw247_.length(0)):
                nw247_[i0_42_] = init42_(i0_42_)
            d_1664_v10_ = nw247_
            d_1668_v11_: _dafny.Map
            d_1668_v11_ = _dafny.Map({p0: d_1664_v10_})
            index254_ = default__.safeIndex(911, (d_1662_v8_).length(0))
            rhs261_ = (d_1663_v9_ if p3 else d_1663_v9_)
            rhs262_ = ((d_1668_v11_)[p3] if (p3) in (d_1668_v11_) else d_1664_v10_)
            rhs263_ = p0
            rhs264_ = (d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))]
            lhs229_ = d_1662_v8_
            lhs230_ = default__.safeIndex(911, (d_1662_v8_).length(0))
            lhs231_ = globalState
            lhs232_ = globalState
            lhs229_[lhs230_] = rhs261_
            d_1664_v10_ = rhs262_
            lhs231_.f25 = rhs263_
            lhs232_.f5 = rhs264_
            if not(p0):
                d_1669_v12_: D2
                d_1669_v12_ = D2_DC7(d_1655_v4_)
                d_1670_v13_: _dafny.Array
                def lambda79_(d_1671_p3_):
                    def lambda80_(d_1672_i4_):
                        return d_1671_p3_

                    return lambda80_

                init43_ = lambda79_(p3)
                nw248_ = _dafny.Array(None, 6)
                for i0_43_ in range(nw248_.length(0)):
                    nw248_[i0_43_] = init43_(i0_43_)
                d_1670_v13_ = nw248_
                rhs265_ = p0
                rhs266_ = d_1670_v13_
                rhs267_ = D2_DC7(d_1655_v4_)
                lhs233_ = globalState
                lhs234_ = globalState
                lhs233_.f9 = rhs265_
                lhs234_.f1 = rhs266_
                d_1669_v12_ = rhs267_
                d_1673_v14_: _dafny.Array
                nw249_ = _dafny.Array(None, 1)
                d_1673_v14_ = nw249_
                d_1674_v15_: C0
                nw250_ = C0()
                nw250_.ctor__()
                d_1674_v15_ = nw250_
                index255_ = default__.safeIndex(442, (d_1673_v14_).length(0))
                (d_1673_v14_)[index255_] = d_1674_v15_
                index256_ = default__.safeIndex(625, (d_1670_v13_).length(0))
                (d_1670_v13_)[index256_] = default__.fm24(p3, p3, p0, d_1653_v2_, globalState)
                d_1675_v16_: _dafny.Array
                def lambda81_(d_1676_p2_):
                    def lambda82_(d_1677_i5_):
                        return (d_1676_p2_) + (d_1676_p2_)

                    return lambda82_

                init44_ = lambda81_(p2)
                nw251_ = _dafny.Array(None, 16)
                for i0_44_ in range(nw251_.length(0)):
                    nw251_[i0_44_] = init44_(i0_44_)
                d_1675_v16_ = nw251_
                index257_ = default__.safeIndex(510, (d_1675_v16_).length(0))
                (d_1675_v16_)[index257_] = (p2) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1678_i6_ in range(default__.abs(723))]))
                d_1679_v17_: _dafny.MultiSet
                d_1679_v17_ = _dafny.MultiSet([355, d_1653_v2_])
                d_1680_v18_: _dafny.Set
                d_1680_v18_ = _dafny.Set({d_1653_v2_})
                d_1681_v19_: _dafny.Map
                d_1681_v19_ = _dafny.Map({not(((d_1654_v3_)[len(p2)] if (len(p2)) in (d_1654_v3_) else p3)): len(d_1680_v18_)})
                index258_ = default__.safeIndex(442, (d_1673_v14_).length(0))
                index259_ = default__.safeIndex(625, (d_1670_v13_).length(0))
                index260_ = default__.safeIndex(510, (d_1675_v16_).length(0))
                rhs268_ = d_1674_v15_
                rhs269_ = not(default__.fm24(True, (d_1679_v17_).issubset(d_1679_v17_), p3, len(d_1681_v19_), globalState))
                rhs270_ = default__.fm24(p0, ((d_1650_v0_)[p0] if (p0) in (d_1650_v0_) else p0), p3, (d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))], globalState)
                rhs271_ = (((self).f33)[default__.safeIndex((d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))], len((self).f33))]) != (len(d_1655_v4_))
                rhs272_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmftdfysg"))).set(default__.safeIndex(default__.safeDivisionInt((d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))], 76), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmftdfysg")))), _dafny.CodePoint('i'))
                lhs235_ = d_1673_v14_
                lhs236_ = default__.safeIndex(442, (d_1673_v14_).length(0))
                lhs237_ = globalState
                lhs238_ = d_1670_v13_
                lhs239_ = default__.safeIndex(625, (d_1670_v13_).length(0))
                lhs240_ = globalState
                lhs241_ = d_1675_v16_
                lhs242_ = default__.safeIndex(510, (d_1675_v16_).length(0))
                lhs235_[lhs236_] = rhs268_
                lhs237_.f25 = rhs269_
                lhs238_[lhs239_] = rhs270_
                lhs240_.f9 = rhs271_
                lhs241_[lhs242_] = rhs272_
                d_1682_v20_: C1
                nw252_ = C1()
                nw252_.ctor__()
                d_1682_v20_ = nw252_
                d_1651_v1_ = d_1651_v1_
                def iife132_():
                    coll78_ = _dafny.Set()
                    compr_78_: int
                    for compr_78_ in _dafny.IntegerRange(-481, 367):
                        d_1683_v21_: int = compr_78_
                        if ((-481) <= (d_1683_v21_)) and ((d_1683_v21_) < (367)):
                            coll78_ = coll78_.union(_dafny.Set([default__.safeModuloInt(d_1683_v21_, 834)]))
                    return _dafny.Set(coll78_)
                d_1680_v18_ = iife132_()
                
            elif True:
                d_1684_v22_: C0
                nw253_ = C0()
                nw253_.ctor__()
                d_1684_v22_ = nw253_
                (globalState).f9 = (p0 if not (True) or (p0) else p0)
                d_1685_v23_: str
                d_1685_v23_ = _dafny.CodePoint('e')
                d_1685_v23_ = _dafny.CodePoint('n')
                r2 = p0
                d_1686_v24_: _dafny.Array
                nw254_ = _dafny.Array(_dafny.Seq({}), 3)
                d_1686_v24_ = nw254_
                index261_ = default__.safeIndex(217, (d_1686_v24_).length(0))
                (d_1686_v24_)[index261_] = (self).f33
                index262_ = default__.safeIndex(217, (d_1686_v24_).length(0))
                (d_1686_v24_)[index262_] = (self).f33
            d_1687_v25_: _dafny.Array
            nw255_ = _dafny.Array(False, 26)
            d_1687_v25_ = nw255_
            index263_ = default__.safeIndex(313, (d_1687_v25_).length(0))
            (d_1687_v25_)[index263_] = True
            index264_ = default__.safeIndex(313, (d_1687_v25_).length(0))
            (d_1687_v25_)[index264_] = not (p3) or ((d_1655_v4_)[default__.safeIndex(d_1653_v2_, len(d_1655_v4_))])
            (globalState).f0 = (d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))]
            (globalState).f9 = (d_1687_v25_)[default__.safeIndex(313, (d_1687_v25_).length(0))]
        elif True:
            d_1688_v26_: _dafny.Seq
            d_1688_v26_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            d_1689_v27_: _dafny.Map
            d_1689_v27_ = _dafny.Map({_dafny.Set({p3}): len(d_1650_v0_)})
            d_1690_v28_: _dafny.Array
            nw256_ = _dafny.Array(None, 28)
            nw256_[int(0)] = p0
            nw256_[int(1)] = p0
            nw256_[int(2)] = p3
            nw256_[int(3)] = not(False)
            nw256_[int(4)] = not(((self).f33) < ((self).f33))
            nw256_[int(5)] = p3
            nw256_[int(6)] = p0
            nw256_[int(7)] = p3
            nw256_[int(8)] = not (p0) or (p0)
            nw256_[int(9)] = p0
            nw256_[int(10)] = ((_dafny.MultiSet([d_1653_v2_])).cardinality) < (611)
            nw256_[int(11)] = not((not(p0)) and (p0))
            nw256_[int(12)] = p0
            nw256_[int(13)] = p3
            nw256_[int(14)] = p0
            nw256_[int(15)] = p3
            nw256_[int(16)] = p0
            nw256_[int(17)] = p3
            nw256_[int(18)] = (not(p0) if False else p0)
            nw256_[int(19)] = p0
            nw256_[int(20)] = (d_1653_v2_) != (379)
            nw256_[int(21)] = p3
            nw256_[int(22)] = p3
            nw256_[int(23)] = p0
            nw256_[int(24)] = p3
            nw256_[int(25)] = p3
            nw256_[int(26)] = p0
            nw256_[int(27)] = (p2) == ((d_1688_v26_)[default__.safeIndex(len(d_1689_v27_), len(d_1688_v26_))])
            d_1690_v28_ = nw256_
            index265_ = default__.safeIndex(641, (d_1690_v28_).length(0))
            (d_1690_v28_)[index265_] = True
            d_1691_v29_: _dafny.Set
            d_1691_v29_ = _dafny.Set({p3, False, p0, p0})
            d_1692_v30_: D1
            d_1692_v30_ = D1_DC5((d_1691_v29_).intersection(d_1691_v29_))
            d_1693_v31_: _dafny.Map
            d_1693_v31_ = _dafny.Map({True: p2})
            d_1694_v33_: _dafny.Map
            def iife133_():
                coll79_ = _dafny.Map()
                compr_79_: int
                for compr_79_ in _dafny.IntegerRange(372, 695):
                    d_1695_v32_: int = compr_79_
                    if ((372) <= (d_1695_v32_)) and ((d_1695_v32_) < (695)):
                        coll79_[(d_1695_v32_) * (((self).f33)[default__.safeIndex(d_1653_v2_, len((self).f33))])] = False
                return _dafny.Map(coll79_)
            d_1694_v33_ = _dafny.Map({p0: iife133_()
            })
            d_1696_v34_: _dafny.Map
            d_1696_v34_ = _dafny.Map({(d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))]: d_1691_v29_})
            pat_let_tv35_ = d_1696_v34_
            pat_let_tv36_ = d_1651_v1_
            pat_let_tv37_ = d_1651_v1_
            pat_let_tv38_ = d_1651_v1_
            pat_let_tv39_ = d_1651_v1_
            pat_let_tv40_ = d_1696_v34_
            pat_let_tv41_ = d_1691_v29_
            pat_let_tv42_ = d_1691_v29_
            index266_ = default__.safeIndex(641, (d_1690_v28_).length(0))
            def iife134_(_pat_let27_0):
                def iife135_(d_1697_dt__update__tmp_h0_):
                    def iife136_(_pat_let28_0):
                        def iife137_(d_1698_dt__update_hcf8_h0_):
                            return D1_DC5(d_1698_dt__update_hcf8_h0_)
                        return iife137_(_pat_let28_0)
                    return iife136_((((pat_let_tv35_)[(pat_let_tv37_)[default__.safeIndex(737, (pat_let_tv36_).length(0))]] if ((pat_let_tv39_)[default__.safeIndex(737, (pat_let_tv38_).length(0))]) in (pat_let_tv40_) else pat_let_tv41_)) - (pat_let_tv42_))
                return iife135_(_pat_let27_0)
            rhs273_ = (self).fm6((d_1654_v3_) | (d_1654_v3_), len(((d_1693_v31_).set(p0, p2)) | (d_1693_v31_)), (self).f33, d_1656_v5_, globalState)
            rhs274_ = (0) - (((0) - (d_1653_v2_)) * (((0) - ((d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))])) - ((d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))])))
            rhs275_ = p3
            rhs276_ = ((len(((d_1694_v33_)[p3] if (p3) in (d_1694_v33_) else ((d_1694_v33_)[p3] if (p3) in (d_1694_v33_) else _dafny.Map({d_1653_v2_: p3}))))) <= (d_1653_v2_) if p0 else p3)
            rhs277_ = iife134_((d_1692_v30_ if p3 else D1_DC5(d_1691_v29_)))
            lhs243_ = globalState
            lhs244_ = globalState
            lhs245_ = d_1690_v28_
            lhs246_ = default__.safeIndex(641, (d_1690_v28_).length(0))
            lhs247_ = globalState
            lhs243_.f17 = rhs273_
            lhs244_.f5 = rhs274_
            lhs245_[lhs246_] = rhs275_
            lhs247_.f25 = rhs276_
            d_1692_v30_ = rhs277_
            (globalState).f0 = (0) - ((d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))])
            d_1699_v35_: _dafny.MultiSet
            d_1699_v35_ = _dafny.MultiSet([p0, (d_1690_v28_)[default__.safeIndex(641, (d_1690_v28_).length(0))]])
            r1 = (default__.safeModuloInt(len(_dafny.Set({(d_1690_v28_)[default__.safeIndex(641, (d_1690_v28_).length(0))]})), len(p2))) * ((d_1699_v35_).cardinality)
            d_1700_v36_: C1
            nw257_ = C1()
            nw257_.ctor__()
            d_1700_v36_ = nw257_
            (globalState).f0 = (self).fm7((d_1700_v36_).fm7((d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))], p0, globalState), p3, globalState)
        r0 = not (p0) or (((0) - ((d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))])) >= ((d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))]))
        d_1701_v37_: D7
        d_1701_v37_ = D7_DC18(d_1653_v2_, d_1651_v1_)
        d_1702_v38_: str
        d_1702_v38_ = _dafny.CodePoint('p')
        r1 = len((_dafny.SeqWithoutIsStrInference([(d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))] for d_1703_i7_ in range(default__.abs(912))])).set(default__.safeIndex(((0) - (d_1653_v2_)) - ((d_1701_v37_).cf34), len(_dafny.SeqWithoutIsStrInference([(d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))] for d_1704_i7_ in range(default__.abs(912))]))), (0) - (len(_dafny.Map({d_1702_v38_: (d_1651_v1_)[default__.safeIndex(737, (d_1651_v1_).length(0))]})))))
        r2 = p3
        return r0, r1, r2

    @property
    def f33(self):
        return self._f33

class C3(T1, T0):
    def  __init__(self):
        self._f26: int = int(0)
        self._f27: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f34: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f34, f26, f27):
        (self)._f34 = f34
        (self)._f26 = f26
        (self)._f27 = f27

    def fm8(self, p0, p1, p2, p3, globalState):
        return (((self).f27) + ((self).f27)) not in ((_dafny.Set({(self).f27})) - (_dafny.Set({(self).f27})))

    def fm6(self, p0, p1, p2, p3, globalState):
        def iife138_():
            def iife141_():
                coll83_ = _dafny.Map()
                compr_83_: _dafny.Seq
                for compr_83_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('y')])])).Elements:
                    d_1705_v1_: _dafny.Seq = compr_83_
                    if (d_1705_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('y')])])):
                        coll83_[d_1705_v1_] = D8_DC22(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aq"))): False}))
                return _dafny.Map(coll83_)
            def iife142_():
                coll84_ = _dafny.Map()
                compr_84_: int
                for compr_84_ in (_dafny.SeqWithoutIsStrInference([(self).f34])).Elements:
                    d_1706_v2_: int = compr_84_
                    if (d_1706_v2_) in (_dafny.SeqWithoutIsStrInference([(self).f34])):
                        coll84_[(d_1706_v2_) + (-351)] = True
                return _dafny.Map(coll84_)
            coll80_ = _dafny.Map()
            def iife139_():
                coll81_ = _dafny.Map()
                compr_81_: _dafny.Seq
                for compr_81_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('y')])])).Elements:
                    d_1705_v1_: _dafny.Seq = compr_81_
                    if (d_1705_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('y')])])):
                        coll81_[d_1705_v1_] = D8_DC22(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aq"))): False}))
                return _dafny.Map(coll81_)
            def iife140_():
                coll82_ = _dafny.Map()
                compr_82_: int
                for compr_82_ in (_dafny.SeqWithoutIsStrInference([(self).f34])).Elements:
                    d_1706_v2_: int = compr_82_
                    if (d_1706_v2_) in (_dafny.SeqWithoutIsStrInference([(self).f34])):
                        coll82_[(d_1706_v2_) + (-351)] = True
                return _dafny.Map(coll82_)
            compr_80_: _dafny.Seq
            for compr_80_ in ((iife139_()
            ) | (_dafny.Map({(self).f27: D8_DC22(iife140_()
)}))).keys.Elements:
                d_1707_v0_: _dafny.Seq = compr_80_
                if (d_1707_v0_) in ((iife141_()
                ) | (_dafny.Map({(self).f27: D8_DC22(iife142_()
)}))):
                    coll80_[d_1707_v0_] = True
            return _dafny.Map(coll80_)
        return len(iife138_()
        )

    def fm7(self, p0, p1, globalState):
        return default__.safeModuloInt(13, (self).f26)

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        d_1708_v1_: _dafny.Seq
        def iife143_():
            coll85_ = _dafny.Set()
            compr_85_: int
            for compr_85_ in _dafny.IntegerRange(768, 144):
                d_1709_v0_: int = compr_85_
                if ((768) <= (d_1709_v0_)) and ((d_1709_v0_) < (144)):
                    coll85_ = coll85_.union(_dafny.Set([default__.safeDivisionInt(d_1709_v0_, (self).f26)]))
            return _dafny.Set(coll85_)
        d_1708_v1_ = _dafny.SeqWithoutIsStrInference([(-310) * (len(iife143_()
        ))])
        d_1708_v1_ = ((_dafny.SeqWithoutIsStrInference([(self).f26 for d_1710_i0_ in range(default__.abs(443))]) if p1 else d_1708_v1_) if (p1 if p1 else p1) else (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_1708_v1_, d_1708_v1_}))])) + (d_1708_v1_))
        d_1711_v2_: _dafny.MultiSet
        d_1711_v2_ = _dafny.MultiSet([p1])
        r1 = (p1) and ((d_1711_v2_).ispropersubset(d_1711_v2_))
        d_1712_v3_: _dafny.Map
        d_1712_v3_ = _dafny.Map({p1: (self).f34})
        hi7_ = (self).f26
        for d_1713_i1_ in range((d_1708_v1_)[default__.safeIndex(len((d_1712_v3_).set(p1, len((self).f27))), len(d_1708_v1_))], hi7_):
            d_1714_v4_: D15
            d_1714_v4_ = D15_DC42(d_1711_v2_)
            if ((d_1711_v2_) - (d_1711_v2_)) != ((_dafny.MultiSet([True, p1])).intersection(((d_1714_v4_).cf80).set(not(p1), default__.abs(d_1713_i1_)))):
                d_1715_v5_: _dafny.MultiSet
                d_1715_v5_ = _dafny.MultiSet([(self).f34])
                d_1716_v6_: _dafny.Set
                d_1716_v6_ = _dafny.Set({(self).fm8(d_1715_v5_, p1, (self).f26, (self).f26, globalState)})
                d_1717_v7_: _dafny.Seq
                d_1718_v8_: int
                d_1719_v9_: int
                d_1720_v10_: bool
                out43_: _dafny.Seq
                out44_: int
                out45_: int
                out46_: bool
                out43_, out44_, out45_, out46_ = (self).m7(d_1716_v6_, globalState)
                d_1717_v7_ = out43_
                d_1718_v8_ = out44_
                d_1719_v9_ = out45_
                d_1720_v10_ = out46_
                d_1721_v11_: str
                d_1721_v11_ = _dafny.CodePoint('b')
                d_1722_v12_: _dafny.Map
                d_1722_v12_ = _dafny.Map({d_1718_v8_: (d_1721_v11_) not in (d_1717_v7_)})
                d_1722_v12_ = (d_1722_v12_).set(d_1713_i1_, not(((_dafny.MultiSet([d_1720_v10_, True])).cardinality) < ((self).f26)))
                d_1723_v13_: _dafny.Seq
                d_1723_v13_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
                d_1724_v14_: _dafny.Map
                d_1724_v14_ = _dafny.Map({d_1720_v10_: (p1 if not((d_1723_v13_)[default__.safeIndex(438, len(d_1723_v13_))]) else ((d_1722_v12_)[d_1718_v8_] if (d_1718_v8_) in (d_1722_v12_) else p1))})
                d_1725_v15_: _dafny.Map
                d_1725_v15_ = _dafny.Map({not(d_1720_v10_): d_1721_v11_})
                (globalState).f9 = not(((d_1724_v14_)[(self).fm8(d_1715_v5_, p1, (0) - ((self).f34), len(d_1725_v15_), globalState)] if ((self).fm8(d_1715_v5_, p1, (0) - ((self).f34), len(d_1725_v15_), globalState)) in (d_1724_v14_) else p1))
                d_1726_v16_: C0
                nw258_ = C0()
                nw258_.ctor__()
                d_1726_v16_ = nw258_
                d_1727_v17_: C1
                nw259_ = C1()
                nw259_.ctor__()
                d_1727_v17_ = nw259_
            elif True:
                d_1728_v18_: _dafny.Array
                def lambda83_(d_1729_v4_):
                    def lambda84_(d_1730_i2_):
                        return d_1729_v4_

                    return lambda84_

                init45_ = lambda83_(d_1714_v4_)
                nw260_ = _dafny.Array(None, 21)
                for i0_45_ in range(nw260_.length(0)):
                    nw260_[i0_45_] = init45_(i0_45_)
                d_1728_v18_ = nw260_
                rhs278_ = d_1728_v18_
                rhs279_ = p1
                lhs248_ = globalState
                d_1728_v18_ = rhs278_
                lhs248_.f25 = rhs279_
                d_1731_v19_: _dafny.Set
                d_1731_v19_ = _dafny.Set({((d_1708_v1_)[default__.safeIndex(198, len(d_1708_v1_))]) + ((self).f34), len(((self).f27) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1732_i3_ in range(default__.abs(391))])))})
                d_1733_v20_: _dafny.Array
                def lambda85_(d_1734_i4_):
                    return (self).f27

                init46_ = lambda85_
                nw261_ = _dafny.Array(None, 13)
                for i0_46_ in range(nw261_.length(0)):
                    nw261_[i0_46_] = init46_(i0_46_)
                d_1733_v20_ = nw261_
                index267_ = default__.safeIndex(941, (d_1733_v20_).length(0))
                (d_1733_v20_)[index267_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1735_i5_ in range(default__.abs(77))])) + ((self).f27)
                d_1736_v21_: _dafny.Map
                d_1736_v21_ = _dafny.Map({p1: False})
                d_1737_v22_: D2
                d_1737_v22_ = D2_DC8(p1, (self).f34, d_1713_i1_)
                index268_ = default__.safeIndex(941, (d_1733_v20_).length(0))
                rhs280_ = d_1731_v19_
                rhs281_ = ((self).f27) + ((self).f27)
                rhs282_ = len(((d_1736_v21_) | (d_1736_v21_)) | (_dafny.Map({False: (d_1737_v22_).cf13})))
                rhs283_ = not(p1)
                rhs284_ = d_1733_v20_
                lhs249_ = d_1733_v20_
                lhs250_ = default__.safeIndex(941, (d_1733_v20_).length(0))
                lhs251_ = globalState
                lhs252_ = globalState
                d_1731_v19_ = rhs280_
                lhs249_[lhs250_] = rhs281_
                lhs251_.f8 = rhs282_
                lhs252_.f25 = rhs283_
                d_1733_v20_ = rhs284_
                d_1738_v23_: str
                d_1738_v23_ = _dafny.CodePoint('n')
                d_1738_v23_ = d_1738_v23_
                d_1739_v24_: _dafny.MultiSet
                d_1739_v24_ = _dafny.MultiSet([(0) - ((self).f26), len((self).f27)])
                def iife144_():
                    coll86_ = _dafny.Map()
                    compr_86_: int
                    for compr_86_ in _dafny.IntegerRange(1, -737):
                        d_1740_v25_: int = compr_86_
                        if ((1) <= (d_1740_v25_)) and ((d_1740_v25_) < (-737)):
                            coll86_[default__.safeModuloInt(d_1740_v25_, d_1713_i1_)] = p1
                    return _dafny.Map(coll86_)
                (globalState).f9 = (self).fm8(d_1739_v24_, False, len(_dafny.SeqWithoutIsStrInference([(self).f34, d_1713_i1_, d_1713_i1_, (self).f34, (0) - (len((self).f27))])), len(iife144_()
                ), globalState)
                d_1741_v26_: _dafny.Map
                d_1741_v26_ = _dafny.Map({(self).f34: ((d_1711_v2_)[p1] if (p1) in (d_1711_v2_) else d_1713_i1_)})
                d_1742_v27_: C2
                nw262_ = C2()
                nw262_.ctor__(_dafny.SeqWithoutIsStrInference([(self).f26, len(_dafny.SeqWithoutIsStrInference([d_1713_i1_ for d_1743_i6_ in range(default__.abs(489))])), (self).f26, (self).f26]), d_1741_v26_)
                d_1742_v27_ = nw262_
            d_1744_v28_: _dafny.Array
            nw263_ = _dafny.Array(_dafny.Map({}), 6)
            d_1744_v28_ = nw263_
            d_1745_v29_: T0
            nw264_ = C1()
            nw264_.ctor__()
            d_1745_v29_ = nw264_
            d_1746_v30_: _dafny.Map
            d_1746_v30_ = _dafny.Map({d_1745_v29_: (self).f34})
            index269_ = default__.safeIndex(628, (d_1744_v28_).length(0))
            (d_1744_v28_)[index269_] = d_1746_v30_
            index270_ = default__.safeIndex(53, (p0).length(0))
            (p0)[index270_] = d_1713_i1_
            d_1747_v31_: str
            d_1747_v31_ = _dafny.CodePoint('f')
            index271_ = default__.safeIndex(628, (d_1744_v28_).length(0))
            index272_ = default__.safeIndex(53, (p0).length(0))
            rhs285_ = ((self).f26) <= (len(_dafny.Map({p1: d_1747_v31_})))
            rhs286_ = (self).f26
            rhs287_ = _dafny.Map({d_1745_v29_: (self).f26})
            rhs288_ = (default__.safeModuloInt(len((self).f27), (self).f26)) + (814)
            lhs253_ = globalState
            lhs254_ = d_1744_v28_
            lhs255_ = default__.safeIndex(628, (d_1744_v28_).length(0))
            lhs256_ = p0
            lhs257_ = default__.safeIndex(53, (p0).length(0))
            lhs253_.f9 = rhs285_
            r0 = rhs286_
            lhs254_[lhs255_] = rhs287_
            lhs256_[lhs257_] = rhs288_
            (globalState).f25 = p1
            (globalState).f25 = p1
        d_1748_v32_: _dafny.MultiSet
        d_1748_v32_ = _dafny.MultiSet([(self).f26, (self).f34])
        d_1749_v33_: _dafny.Map
        d_1749_v33_ = _dafny.Map({(self).f26: p1})
        d_1750_v34_: _dafny.Seq
        d_1750_v34_ = _dafny.SeqWithoutIsStrInference([(self).fm8((d_1748_v32_).set((self).f34, default__.abs(len(d_1749_v33_))), True, (self).f34, (self).f34, globalState), p1, False])
        d_1751_v35_: _dafny.Seq
        d_1751_v35_ = _dafny.SeqWithoutIsStrInference([(d_1750_v34_)[default__.safeIndex((self).f34, len(d_1750_v34_))]])
        (globalState).f9 = (d_1750_v34_)[default__.safeIndex((self).f26, len(d_1750_v34_))]
        if True:
            d_1752_v37_: _dafny.Set
            d_1752_v37_ = _dafny.Set({d_1711_v2_})
            d_1753_v38_: D12
            def iife145_():
                coll87_ = _dafny.Map()
                compr_87_: _dafny.MultiSet
                for compr_87_ in (d_1752_v37_).Elements:
                    d_1754_v36_: _dafny.MultiSet = compr_87_
                    if (d_1754_v36_) in (d_1752_v37_):
                        coll87_[d_1754_v36_] = D3_DC10(_dafny.Map({(D3_DC11(_dafny.Map({-850: len((self).f27)}), False, len((self).f27), (self).f26)).cf22: p1}))
                return _dafny.Map(coll87_)
            d_1753_v38_ = D12_DC36(iife145_()
)
            d_1755_v39_: _dafny.Array
            def lambda86_(d_1756_p1_, d_1757_v2_):
                def lambda87_(d_1758_i7_):
                    return (_dafny.Set({len((D0_DC1((self).f27, ((d_1757_v2_)[d_1756_p1_] if (d_1756_p1_) in (d_1757_v2_) else (self).f26), 86)).cf1)})) - (_dafny.Set({(self).f34}))

                return lambda87_

            init47_ = lambda86_(p1, d_1711_v2_)
            nw265_ = _dafny.Array(None, 26)
            for i0_47_ in range(nw265_.length(0)):
                nw265_[i0_47_] = init47_(i0_47_)
            d_1755_v39_ = nw265_
            d_1759_v40_: _dafny.Set
            d_1759_v40_ = _dafny.Set({-412})
            index273_ = default__.safeIndex(437, (d_1755_v39_).length(0))
            (d_1755_v39_)[index273_] = d_1759_v40_
            d_1760_v41_: _dafny.Array
            def lambda88_(d_1761_i8_):
                return (_dafny.CodePoint('q')) in ((self).f27)

            init48_ = lambda88_
            nw266_ = _dafny.Array(None, 17)
            for i0_48_ in range(nw266_.length(0)):
                nw266_[i0_48_] = init48_(i0_48_)
            d_1760_v41_ = nw266_
            d_1762_v42_: _dafny.Seq
            d_1762_v42_ = _dafny.SeqWithoutIsStrInference([d_1749_v33_, d_1749_v33_, _dafny.Map({len(d_1749_v33_): p1})])
            d_1763_v44_: _dafny.Seq
            def iife146_():
                coll88_ = _dafny.Map()
                compr_88_: int
                for compr_88_ in _dafny.IntegerRange(280, 789):
                    d_1764_v43_: int = compr_88_
                    if ((280) <= (d_1764_v43_)) and ((d_1764_v43_) < (789)):
                        coll88_[(d_1764_v43_) - (len((self).f27))] = p1
                return _dafny.Map(coll88_)
            d_1763_v44_ = _dafny.SeqWithoutIsStrInference([(d_1762_v42_)[default__.safeIndex((self).f26, len(d_1762_v42_))], d_1749_v33_, d_1749_v33_, iife146_()
            ])
            d_1765_v45_: _dafny.Seq
            d_1765_v45_ = _dafny.SeqWithoutIsStrInference([d_1708_v1_])
            d_1766_v47_: _dafny.Seq
            d_1766_v47_ = _dafny.SeqWithoutIsStrInference([d_1711_v2_, d_1711_v2_, d_1711_v2_])
            d_1767_v48_: _dafny.Seq
            d_1767_v48_ = _dafny.SeqWithoutIsStrInference([d_1759_v40_, d_1759_v40_])
            index274_ = default__.safeIndex(437, (d_1755_v39_).length(0))
            def iife147_():
                coll89_ = _dafny.Map()
                compr_89_: _dafny.MultiSet
                for compr_89_ in (d_1766_v47_).Elements:
                    d_1768_v46_: _dafny.MultiSet = compr_89_
                    if (d_1768_v46_) in (d_1766_v47_):
                        coll89_[d_1768_v46_] = D3_DC10(_dafny.Map({not(p1): not((D10_DC31(p1, d_1708_v1_, p1, (self).f34)).cf58)}))
                return _dafny.Map(coll89_)
            rhs289_ = d_1760_v41_
            rhs290_ = (d_1763_v44_) != ((default__.fm41(p1, (d_1765_v45_)[default__.safeIndex((self).f34, len(d_1765_v45_))], globalState)) + (d_1762_v42_))
            rhs291_ = D12_DC36(iife147_()
)
            rhs292_ = ((_dafny.Set({(self).f26, (self).f34})) | ((d_1767_v48_)[default__.safeIndex((self).f34, len(d_1767_v48_))])) - (d_1759_v40_)
            rhs293_ = (len(d_1759_v40_) if p1 else (self).f26)
            lhs258_ = globalState
            lhs259_ = globalState
            lhs260_ = d_1755_v39_
            lhs261_ = default__.safeIndex(437, (d_1755_v39_).length(0))
            lhs262_ = globalState
            lhs258_.f1 = rhs289_
            lhs259_.f25 = rhs290_
            d_1753_v38_ = rhs291_
            lhs260_[lhs261_] = rhs292_
            lhs262_.f0 = rhs293_
            d_1711_v2_ = d_1711_v2_
            index275_ = default__.safeIndex(599, (d_1760_v41_).length(0))
            (d_1760_v41_)[index275_] = ((self).fm8(d_1748_v32_, p1, (self).f26, (self).f34, globalState)) == (p1)
            index276_ = default__.safeIndex(599, (d_1760_v41_).length(0))
            (d_1760_v41_)[index276_] = p1
            (globalState).f25 = p1
            d_1769_v49_: str
            d_1769_v49_ = _dafny.CodePoint('k')
            d_1769_v49_ = default__.fm42(d_1759_v40_, globalState)
        elif True:
            d_1770_v50_: _dafny.Array
            nw267_ = _dafny.Array(int(0), 13)
            d_1770_v50_ = nw267_
            d_1770_v50_ = p0
            if p1:
                (globalState).f9 = not (p1) or (True)
                (globalState).f25 = p1
                (globalState).f0 = ((self).f26) * (len((d_1751_v35_) + (d_1751_v35_)))
                d_1771_v51_: _dafny.Array
                nw268_ = _dafny.Array(_dafny.Set({}), 10)
                d_1771_v51_ = nw268_
                d_1772_v52_: _dafny.Set
                d_1772_v52_ = _dafny.Set({d_1770_v50_})
                index277_ = default__.safeIndex(119, (d_1771_v51_).length(0))
                (d_1771_v51_)[index277_] = (d_1772_v52_) | (_dafny.Set({d_1770_v50_}))
                index278_ = default__.safeIndex(119, (d_1771_v51_).length(0))
                (d_1771_v51_)[index278_] = (d_1772_v52_).intersection(d_1772_v52_)
                d_1749_v33_ = (d_1749_v33_).set(default__.safeDivisionInt((self).f26, (self).f34), not(p1))
            elif True:
                r1 = p1
                d_1773_v53_: _dafny.Map
                d_1773_v53_ = _dafny.Map({d_1751_v35_: (0) - ((self).f34)})
                d_1773_v53_ = d_1773_v53_
                d_1774_v54_: _dafny.Array
                def lambda89_(d_1775_p1_):
                    def lambda90_(d_1776_i9_):
                        return d_1775_p1_

                    return lambda90_

                init49_ = lambda89_(p1)
                nw269_ = _dafny.Array(None, 1)
                for i0_49_ in range(nw269_.length(0)):
                    nw269_[i0_49_] = init49_(i0_49_)
                d_1774_v54_ = nw269_
                d_1777_v55_: _dafny.Array
                nw270_ = _dafny.Array(None, 13)
                nw270_[int(0)] = d_1774_v54_
                nw270_[int(1)] = d_1774_v54_
                nw270_[int(2)] = d_1774_v54_
                nw270_[int(3)] = d_1774_v54_
                nw270_[int(4)] = d_1774_v54_
                nw270_[int(5)] = (d_1774_v54_ if False else d_1774_v54_)
                nw270_[int(6)] = d_1774_v54_
                nw270_[int(7)] = d_1774_v54_
                nw270_[int(8)] = d_1774_v54_
                nw270_[int(9)] = d_1774_v54_
                nw270_[int(10)] = d_1774_v54_
                nw270_[int(11)] = d_1774_v54_
                nw270_[int(12)] = d_1774_v54_
                d_1777_v55_ = nw270_
                d_1777_v55_ = d_1777_v55_
                d_1778_v56_: _dafny.Array
                def lambda91_(d_1779_i10_):
                    return _dafny.MultiSet([(self).f26, (self).f34, (self).f34])

                init50_ = lambda91_
                nw271_ = _dafny.Array(None, 24)
                for i0_50_ in range(nw271_.length(0)):
                    nw271_[i0_50_] = init50_(i0_50_)
                d_1778_v56_ = nw271_
                d_1780_v57_: D11
                d_1780_v57_ = D11_DC33(d_1778_v56_)
                d_1781_v58_: _dafny.Seq
                d_1781_v58_ = _dafny.SeqWithoutIsStrInference([d_1780_v57_, d_1780_v57_])
                d_1782_v59_: D11
                d_1782_v59_ = D11_DC35((d_1781_v58_)[default__.safeIndex((self).f34, len(d_1781_v58_))])
                d_1783_v60_: D11
                d_1783_v60_ = D11_DC35(d_1782_v59_)
                pat_let_tv43_ = d_1778_v56_
                def iife148_(_pat_let29_0):
                    def iife149_(d_1784_dt__update__tmp_h0_):
                        def iife150_(_pat_let30_0):
                            def iife151_(d_1785_dt__update_hcf68_h0_):
                                return D11_DC35(d_1785_dt__update_hcf68_h0_)
                            return iife151_(_pat_let30_0)
                        return iife150_(D11_DC33(pat_let_tv43_))
                    return iife149_(_pat_let29_0)
                d_1783_v60_ = iife148_(D11_DC35(d_1782_v59_))
                index279_ = default__.safeIndex(384, (d_1774_v54_).length(0))
                (d_1774_v54_)[index279_] = p1
                d_1786_v61_: D7
                d_1786_v61_ = D7_DC19((self).f34, p1)
                index280_ = default__.safeIndex(384, (d_1774_v54_).length(0))
                (d_1774_v54_)[index280_] = ((len(_dafny.SeqWithoutIsStrInference([(d_1786_v61_).cf37]))) * ((self).f26)) < (((self).f26) * ((self).f34))
            d_1787_v62_: D7
            d_1787_v62_ = D7_DC20(p1, p1)
            d_1788_v63_: D7
            d_1788_v63_ = D7_DC21(d_1787_v62_)
            d_1789_v64_: _dafny.Seq
            d_1789_v64_ = _dafny.SeqWithoutIsStrInference([d_1788_v63_, d_1788_v63_])
            d_1790_v65_: _dafny.Set
            d_1790_v65_ = _dafny.Set({d_1789_v64_, _dafny.SeqWithoutIsStrInference([d_1788_v63_, D7_DC21(D7_DC20(p1, p1)), d_1788_v63_]), d_1789_v64_})
            d_1791_v66_: _dafny.Array
            def lambda92_(d_1792_v1_):
                def lambda93_(d_1793_i11_):
                    return d_1792_v1_

                return lambda93_

            init51_ = lambda92_(d_1708_v1_)
            nw272_ = _dafny.Array(None, 23)
            for i0_51_ in range(nw272_.length(0)):
                nw272_[i0_51_] = init51_(i0_51_)
            d_1791_v66_ = nw272_
            d_1794_v67_: _dafny.Map
            d_1794_v67_ = _dafny.Map({d_1790_v65_: d_1791_v66_})
            d_1794_v67_ = (d_1794_v67_).set(((D16_DC44(d_1790_v65_)).cf81) - (d_1790_v65_), d_1791_v66_)
            r1 = ((self).f34) > ((self).f26)
            d_1708_v1_ = d_1708_v1_
        (globalState).f9 = not(False)
        r0 = (0) - ((self).f26)
        d_1795_v68_: _dafny.Set
        d_1795_v68_ = _dafny.Set({p1})
        r1 = not (p1) or ((len(d_1795_v68_)) >= ((self).f34))
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_1796_v0_: _dafny.Seq
        d_1796_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1797_v1_: _dafny.Seq
        d_1797_v1_ = _dafny.SeqWithoutIsStrInference([d_1796_v0_, d_1796_v0_])
        if ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0]) for d_1798_i0_ in range(default__.abs(657))])) + (d_1797_v1_)) == (((_dafny.SeqWithoutIsStrInference([d_1796_v0_])).set(default__.safeIndex((self).f34, len(_dafny.SeqWithoutIsStrInference([d_1796_v0_]))), d_1796_v0_)) + (d_1797_v1_)):
            (globalState).f24 = ((self).f26) + ((self).f34)
            d_1799_v2_: D3
            d_1799_v2_ = D3_DC11(_dafny.Map({(self).f26: (self).f34}), p0, (self).f26, (self).f26)
            d_1800_v3_: _dafny.Set
            d_1800_v3_ = _dafny.Set({(d_1799_v2_).cf22})
            d_1801_v4_: _dafny.MultiSet
            d_1801_v4_ = _dafny.MultiSet([(self).f34, len(d_1800_v3_), (self).f26])
            d_1802_v5_: _dafny.Seq
            d_1802_v5_ = _dafny.SeqWithoutIsStrInference([(0) - (len((self).f27))])
            d_1803_v6_: _dafny.Array
            nw273_ = _dafny.Array(None, 13)
            nw273_[int(0)] = (self).f26
            nw273_[int(1)] = (self).f34
            nw273_[int(2)] = 136
            nw273_[int(3)] = (self).f34
            nw273_[int(4)] = (self).f34
            nw273_[int(5)] = len((self).f27)
            nw273_[int(6)] = (self).f34
            nw273_[int(7)] = len(d_1802_v5_)
            nw273_[int(8)] = 459
            nw273_[int(9)] = len((self).f27)
            nw273_[int(10)] = (self).f26
            nw273_[int(11)] = (self).f34
            nw273_[int(12)] = (self).f26
            d_1803_v6_ = nw273_
            d_1804_v7_: D7
            d_1804_v7_ = D7_DC18((self).f34, d_1803_v6_)
            d_1805_v8_: _dafny.Seq
            d_1805_v8_ = _dafny.SeqWithoutIsStrInference([(d_1804_v7_).cf34, (self).f26])
            (globalState).f9 = (self).fm8(d_1801_v4_, p0, (self).f34, (len(d_1805_v8_)) - ((0) - ((self).f34)), globalState)
            d_1806_v9_: _dafny.Map
            d_1806_v9_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0, p0])): (self).f26})
            d_1807_v10_: C2
            nw274_ = C2()
            nw274_.ctor__(d_1802_v5_, d_1806_v9_)
            d_1807_v10_ = nw274_
            d_1808_v11_: _dafny.Array
            nw275_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
            d_1808_v11_ = nw275_
            index281_ = default__.safeIndex(379, (d_1808_v11_).length(0))
            (d_1808_v11_)[index281_] = ((self).f27) + ((self).f27)
            d_1809_v12_: str
            d_1809_v12_ = _dafny.CodePoint('a')
            index282_ = default__.safeIndex(379, (d_1808_v11_).length(0))
            (d_1808_v11_)[index282_] = ((self).f27) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxbou"))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f26 for d_1810_i1_ in range(default__.abs(323))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxbou")))), d_1809_v12_))
            d_1806_v9_ = (d_1806_v9_).set(-946, (0) - (default__.safeModuloInt(457, (self).f26)))
        elif True:
            d_1811_v13_: _dafny.Set
            d_1811_v13_ = _dafny.Set({p0})
            d_1812_v14_: _dafny.Seq
            d_1813_v15_: int
            d_1814_v16_: int
            d_1815_v17_: bool
            out47_: _dafny.Seq
            out48_: int
            out49_: int
            out50_: bool
            out47_, out48_, out49_, out50_ = (self).m7((d_1811_v13_) | (d_1811_v13_), globalState)
            d_1812_v14_ = out47_
            d_1813_v15_ = out48_
            d_1814_v16_ = out49_
            d_1815_v17_ = out50_
            d_1816_v18_: _dafny.Array
            nw276_ = _dafny.Array(False, 9)
            d_1816_v18_ = nw276_
            index283_ = default__.safeIndex(387, (d_1816_v18_).length(0))
            (d_1816_v18_)[index283_] = (952) > ((self).f26)
            d_1817_v19_: _dafny.Array
            nw277_ = _dafny.Array(int(0), 17)
            d_1817_v19_ = nw277_
            d_1818_v20_: _dafny.Map
            d_1818_v20_ = _dafny.Map({(self).f34: _dafny.CodePoint('p')})
            index284_ = default__.safeIndex(982, (d_1817_v19_).length(0))
            (d_1817_v19_)[index284_] = default__.safeModuloInt(len(d_1818_v20_), d_1814_v16_)
            d_1819_v21_: D6
            d_1819_v21_ = D6_DC16((self).f34, d_1815_v17_)
            index285_ = default__.safeIndex(387, (d_1816_v18_).length(0))
            index286_ = default__.safeIndex(982, (d_1817_v19_).length(0))
            rhs294_ = ((d_1819_v21_).cf32) == (d_1815_v17_)
            rhs295_ = default__.safeDivisionInt(default__.safeModuloInt(d_1814_v16_, d_1813_v15_), default__.safeModuloInt(930, (self).f34))
            rhs296_ = d_1814_v16_
            rhs297_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
            lhs263_ = d_1816_v18_
            lhs264_ = default__.safeIndex(387, (d_1816_v18_).length(0))
            lhs265_ = d_1817_v19_
            lhs266_ = default__.safeIndex(982, (d_1817_v19_).length(0))
            lhs267_ = globalState
            lhs268_ = globalState
            lhs263_[lhs264_] = rhs294_
            lhs265_[lhs266_] = rhs295_
            lhs267_.f17 = rhs296_
            lhs268_.f23 = rhs297_
            (globalState).f25 = (d_1816_v18_)[default__.safeIndex(387, (d_1816_v18_).length(0))]
            if (d_1816_v18_)[default__.safeIndex(387, (d_1816_v18_).length(0))]:
                d_1820_v22_: _dafny.Map
                d_1820_v22_ = _dafny.Map({d_1813_v15_: d_1815_v17_})
                d_1821_v23_: D2
                d_1821_v23_ = D2_DC8(d_1815_v17_, d_1813_v15_, (0) - ((self).f26))
                d_1820_v22_ = (d_1820_v22_).set((d_1821_v23_).cf14, (d_1816_v18_)[default__.safeIndex(387, (d_1816_v18_).length(0))])
                d_1822_v24_: _dafny.MultiSet
                d_1822_v24_ = _dafny.MultiSet([(d_1816_v18_)[default__.safeIndex(387, (d_1816_v18_).length(0))]])
                d_1820_v22_ = (d_1820_v22_).set((self).fm7((d_1822_v24_).cardinality, d_1815_v17_, globalState), not((d_1816_v18_)[default__.safeIndex(387, (d_1816_v18_).length(0))]))
                d_1823_v25_: _dafny.Seq
                d_1823_v25_ = _dafny.SeqWithoutIsStrInference([d_1816_v18_, d_1816_v18_])
                d_1816_v18_ = (d_1823_v25_)[default__.safeIndex(-832, len(d_1823_v25_))]
                index287_ = default__.safeIndex(387, (d_1816_v18_).length(0))
                (d_1816_v18_)[index287_] = p0
                d_1824_v26_: _dafny.Array
                nw278_ = _dafny.Array(_dafny.Seq({}), 10)
                d_1824_v26_ = nw278_
                index288_ = default__.safeIndex(287, (d_1824_v26_).length(0))
                (d_1824_v26_)[index288_] = (_dafny.SeqWithoutIsStrInference([d_1815_v17_])) + (d_1796_v0_)
                index289_ = default__.safeIndex(287, (d_1824_v26_).length(0))
                (d_1824_v26_)[index289_] = d_1796_v0_
            elif True:
                d_1825_v27_: _dafny.Map
                d_1825_v27_ = _dafny.Map({d_1816_v18_: (self).f26})
                rhs298_ = d_1825_v27_
                rhs299_ = default__.safeModuloInt(144, (d_1817_v19_)[default__.safeIndex(982, (d_1817_v19_).length(0))])
                lhs269_ = globalState
                d_1825_v27_ = rhs298_
                lhs269_.f24 = rhs299_
                index290_ = default__.safeIndex(982, (d_1817_v19_).length(0))
                (d_1817_v19_)[index290_] = (self).f26
                index291_ = default__.safeIndex(982, (d_1817_v19_).length(0))
                (d_1817_v19_)[index291_] = (0) - ((self).f34)
                d_1826_v28_: str
                d_1826_v28_ = _dafny.CodePoint('f')
                d_1826_v28_ = d_1826_v28_
                (globalState).f9 = p0
            (globalState).f25 = (d_1816_v18_)[default__.safeIndex(387, (d_1816_v18_).length(0))]
        d_1827_v29_: _dafny.Map
        d_1827_v29_ = _dafny.Map({p0: True})
        d_1828_v30_: _dafny.Map
        d_1828_v30_ = _dafny.Map({p0: (self).f34})
        d_1829_v31_: str
        d_1829_v31_ = _dafny.CodePoint('n')
        d_1830_v32_: _dafny.Set
        d_1830_v32_ = _dafny.Set({p0})
        d_1831_v33_: D1
        d_1831_v33_ = D1_DC5(d_1830_v32_)
        d_1832_v34_: _dafny.Map
        d_1832_v34_ = _dafny.Map({len((self).f27): d_1831_v33_})
        d_1833_v35_: _dafny.Map
        d_1833_v35_ = d_1832_v34_
        d_1834_v36_: _dafny.Set
        d_1834_v36_ = _dafny.Set({(self).f26})
        d_1835_v37_: _dafny.Set
        d_1835_v37_ = _dafny.Set({p0, p0, p0, p0, (D8_DC25(p0, d_1828_v30_, d_1829_v31_, d_1833_v35_, len(d_1834_v36_))).cf49})
        d_1827_v29_ = (d_1827_v29_).set((_dafny.Set({p0})).issubset(d_1835_v37_), p0)
        d_1836_v38_: _dafny.MultiSet
        d_1836_v38_ = _dafny.MultiSet([p0, p0, p0, p0, p0])
        d_1837_v39_: D10
        d_1837_v39_ = D10_DC32(p0, (self).f26, p0)
        d_1828_v30_ = (d_1828_v30_).set(not((p0) in (d_1836_v38_)), (d_1837_v39_).cf63)
        d_1838_v40_: _dafny.MultiSet
        d_1838_v40_ = _dafny.MultiSet([(self).f34, (self).f26, (self).f34])
        d_1839_v41_: _dafny.Map
        d_1839_v41_ = _dafny.Map({(self).fm8(d_1838_v40_, p0, (self).f26, (self).f26, globalState): ((self).f27 if p0 else (self).f27)})
        d_1839_v41_ = (d_1839_v41_).set(not(p0), (self).f27)
        if True:
            d_1840_v42_: _dafny.Seq
            d_1840_v42_ = _dafny.SeqWithoutIsStrInference([d_1836_v38_])
            (globalState).f25 = ((d_1840_v42_) + (d_1840_v42_)) < (d_1840_v42_)
            d_1841_v44_: _dafny.Array
            def lambda94_(d_1842_i2_):
                def iife152_():
                    coll90_ = _dafny.Map()
                    compr_90_: int
                    for compr_90_ in _dafny.IntegerRange(932, -221):
                        d_1843_v43_: int = compr_90_
                        if ((932) <= (d_1843_v43_)) and ((d_1843_v43_) < (-221)):
                            coll90_[(d_1843_v43_) * (len((self).f27))] = (self).f26
                    return _dafny.Map(coll90_)
                return iife152_()
                

            init52_ = lambda94_
            nw279_ = _dafny.Array(None, 10)
            for i0_52_ in range(nw279_.length(0)):
                nw279_[i0_52_] = init52_(i0_52_)
            d_1841_v44_ = nw279_
            index292_ = default__.safeIndex(159, (d_1841_v44_).length(0))
            def iife153_():
                coll91_ = _dafny.Map()
                compr_91_: int
                for compr_91_ in _dafny.IntegerRange(615, 625):
                    d_1844_v45_: int = compr_91_
                    if ((615) <= (d_1844_v45_)) and ((d_1844_v45_) < (625)):
                        coll91_[(d_1844_v45_) + ((0) - ((self).f26))] = (d_1836_v38_).cardinality
                return _dafny.Map(coll91_)
            (d_1841_v44_)[index292_] = iife153_()
            
            d_1845_v46_: _dafny.Map
            d_1845_v46_ = _dafny.Map({(self).f26: (self).f34})
            index293_ = default__.safeIndex(159, (d_1841_v44_).length(0))
            (d_1841_v44_)[index293_] = (d_1845_v46_ if p0 else _dafny.Map({300: len(((self).f27).set(default__.safeIndex((self).f26, len((self).f27)), d_1829_v31_))}))
            d_1846_v47_: _dafny.Array
            nw280_ = _dafny.Array(None, 6)
            nw280_[int(0)] = p0
            nw280_[int(1)] = p0
            nw280_[int(2)] = p0
            nw280_[int(3)] = p0
            nw280_[int(4)] = p0
            nw280_[int(5)] = p0
            d_1846_v47_ = nw280_
            d_1847_v48_: _dafny.MultiSet
            d_1847_v48_ = _dafny.MultiSet([d_1846_v47_])
            (globalState).f9 = ((d_1847_v48_).intersection(d_1847_v48_)).isdisjoint(d_1847_v48_)
            d_1848_v49_: _dafny.Array
            nw281_ = _dafny.Array(None, 11)
            d_1848_v49_ = nw281_
            d_1849_v50_: C1
            nw282_ = C1()
            nw282_.ctor__()
            d_1849_v50_ = nw282_
            index294_ = default__.safeIndex(862, (d_1848_v49_).length(0))
            (d_1848_v49_)[index294_] = d_1849_v50_
            index295_ = default__.safeIndex(862, (d_1848_v49_).length(0))
            rhs300_ = d_1849_v50_
            rhs301_ = default__.safeDivisionInt(default__.safeModuloInt((self).f34, (self).f34), 641)
            rhs302_ = (self).f26
            rhs303_ = (self).f34
            lhs270_ = d_1848_v49_
            lhs271_ = default__.safeIndex(862, (d_1848_v49_).length(0))
            lhs272_ = globalState
            lhs273_ = globalState
            lhs274_ = globalState
            lhs270_[lhs271_] = rhs300_
            lhs272_.f5 = rhs301_
            lhs273_.f8 = rhs302_
            lhs274_.f0 = rhs303_
            d_1850_v51_: _dafny.Map
            d_1850_v51_ = _dafny.Map({(self).f26: d_1834_v36_})
            rhs304_ = default__.safeModuloInt(249, (self).f34)
            rhs305_ = default__.safeModuloInt((0) - (len((((d_1850_v51_)[-458] if (-458) in (d_1850_v51_) else d_1834_v36_)).intersection(d_1834_v36_))), (self).f26)
            rhs306_ = p0
            rhs307_ = ((self).fm8((d_1838_v40_).set(len(_dafny.Map({p0: p0})), default__.abs((d_1836_v38_).cardinality)), p0, (self).f34, (self).f26, globalState) if ((self).f34) != ((self).f34) else p0)
            lhs275_ = globalState
            lhs276_ = globalState
            lhs277_ = globalState
            lhs278_ = globalState
            lhs275_.f5 = rhs304_
            lhs276_.f8 = rhs305_
            lhs277_.f9 = rhs306_
            lhs278_.f25 = rhs307_
        elif True:
            (globalState).f8 = (0) - ((self).f26)
            (globalState).f9 = ((0) - ((self).f26)) == ((self).f34)
            d_1851_v52_: _dafny.Array
            def lambda95_(d_1852_p0_):
                def lambda96_(d_1853_i3_):
                    return d_1852_p0_

                return lambda96_

            init53_ = lambda95_(p0)
            nw283_ = _dafny.Array(None, 15)
            for i0_53_ in range(nw283_.length(0)):
                nw283_[i0_53_] = init53_(i0_53_)
            d_1851_v52_ = nw283_
            index296_ = default__.safeIndex(803, (d_1851_v52_).length(0))
            (d_1851_v52_)[index296_] = p0
            d_1854_v53_: _dafny.Array
            nw284_ = _dafny.Array(int(0), 15)
            d_1854_v53_ = nw284_
            index297_ = default__.safeIndex(342, (d_1854_v53_).length(0))
            (d_1854_v53_)[index297_] = (self).f34
            d_1855_v54_: _dafny.Seq
            d_1855_v54_ = _dafny.SeqWithoutIsStrInference([(self).f26, 733])
            index298_ = default__.safeIndex(803, (d_1851_v52_).length(0))
            index299_ = default__.safeIndex(342, (d_1854_v53_).length(0))
            rhs308_ = (self).fm8(_dafny.MultiSet([(d_1855_v54_)[default__.safeIndex(len(d_1835_v37_), len(d_1855_v54_))]]), (93) not in (_dafny.SeqWithoutIsStrInference([(self).f26 for d_1856_i4_ in range(default__.abs(-155))])), (0) - ((self).f26), (self).f26, globalState)
            rhs309_ = (self).f26
            rhs310_ = default__.safeDivisionInt((self).f34, len((self).f27))
            rhs311_ = default__.safeModuloInt((self).f34, ((self).f34 if not(True) else (self).f34))
            lhs279_ = d_1851_v52_
            lhs280_ = default__.safeIndex(803, (d_1851_v52_).length(0))
            lhs281_ = globalState
            lhs282_ = d_1854_v53_
            lhs283_ = default__.safeIndex(342, (d_1854_v53_).length(0))
            lhs284_ = globalState
            lhs279_[lhs280_] = rhs308_
            lhs281_.f17 = rhs309_
            lhs282_[lhs283_] = rhs310_
            lhs284_.f17 = rhs311_
            rhs312_ = default__.safeModuloInt(((self).f34 if False else (d_1854_v53_)[default__.safeIndex(342, (d_1854_v53_).length(0))]), (self).f34)
            rhs313_ = len((self).f27)
            lhs285_ = globalState
            lhs285_.f17 = rhs312_
            r0 = rhs313_
            (globalState).f9 = not (True) or (((self).f26) < ((d_1854_v53_)[default__.safeIndex(342, (d_1854_v53_).length(0))]))
        (globalState).f13 = ((self).f27) + ((self).f27)
        r0 = (self).f34
        d_1857_v55_: _dafny.Map
        d_1857_v55_ = _dafny.Map({(self).f26: p0})
        d_1858_v56_: _dafny.Map
        d_1858_v56_ = _dafny.Map({len(d_1857_v55_): p0})
        d_1859_v57_: D10
        d_1859_v57_ = D10_DC30(d_1838_v40_)
        d_1860_v58_: _dafny.Seq
        d_1860_v58_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        d_1861_v59_: _dafny.Map
        d_1861_v59_ = _dafny.Map({d_1796_v0_: p0})
        d_1862_v60_: D0
        d_1862_v60_ = D0_DC0((self).fm6(d_1857_v55_, ((d_1859_v57_).cf57).cardinality, d_1860_v58_, d_1861_v59_, globalState))
        d_1863_v61_: _dafny.MultiSet
        d_1863_v61_ = _dafny.MultiSet([d_1862_v60_, d_1862_v60_, d_1862_v60_])
        r1 = ((d_1863_v61_).set(d_1862_v60_, default__.abs((self).f26))) | (_dafny.MultiSet([d_1862_v60_]))
        d_1864_v62_: _dafny.Array
        def lambda97_(d_1865_p0_):
            def lambda98_(d_1866_i5_):
                return default__.safeModuloInt(d_1866_i5_, len(_dafny.Map({_dafny.MultiSet([True, d_1865_p0_]): 2})))

            return lambda98_

        init54_ = lambda97_(p0)
        nw285_ = _dafny.Array(None, 1)
        for i0_54_ in range(nw285_.length(0)):
            nw285_[i0_54_] = init54_(i0_54_)
        d_1864_v62_ = nw285_
        d_1867_v63_: _dafny.Map
        d_1867_v63_ = _dafny.Map({(d_1838_v40_).cardinality: d_1864_v62_})
        d_1868_v64_: _dafny.Map
        d_1868_v64_ = _dafny.Map({(self).fm6(d_1857_v55_, (self).f26, _dafny.SeqWithoutIsStrInference([len((self).f27) for d_1869_i6_ in range(default__.abs(802))]), _dafny.Map({d_1796_v0_: p0}), globalState): len((self).f27)})
        d_1870_v65_: _dafny.Map
        d_1870_v65_ = _dafny.Map({d_1829_v31_: d_1864_v62_})
        d_1871_v66_: _dafny.Array
        nw286_ = _dafny.Array(None, 21)
        nw286_[int(0)] = d_1864_v62_
        nw286_[int(1)] = d_1864_v62_
        nw286_[int(2)] = d_1864_v62_
        nw286_[int(3)] = d_1864_v62_
        nw286_[int(4)] = d_1864_v62_
        nw286_[int(5)] = d_1864_v62_
        nw286_[int(6)] = d_1864_v62_
        nw286_[int(7)] = d_1864_v62_
        nw286_[int(8)] = ((d_1867_v63_)[len(d_1868_v64_)] if (len(d_1868_v64_)) in (d_1867_v63_) else d_1864_v62_)
        nw286_[int(9)] = d_1864_v62_
        nw286_[int(10)] = ((d_1870_v65_)[d_1829_v31_] if (d_1829_v31_) in (d_1870_v65_) else d_1864_v62_)
        nw286_[int(11)] = d_1864_v62_
        nw286_[int(12)] = d_1864_v62_
        nw286_[int(13)] = d_1864_v62_
        nw286_[int(14)] = d_1864_v62_
        nw286_[int(15)] = d_1864_v62_
        nw286_[int(16)] = d_1864_v62_
        nw286_[int(17)] = d_1864_v62_
        nw286_[int(18)] = d_1864_v62_
        nw286_[int(19)] = d_1864_v62_
        nw286_[int(20)] = d_1864_v62_
        d_1871_v66_ = nw286_
        r2 = d_1871_v66_
        return r0, r1, r2

    def m7(self, p0, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: int = int(0)
        r2: int = int(0)
        r3: bool = False
        d_1872_v0_: _dafny.Array
        nw287_ = _dafny.Array(False, 12)
        d_1872_v0_ = nw287_
        d_1873_v1_: str
        d_1873_v1_ = _dafny.CodePoint('u')
        d_1874_v2_: _dafny.Map
        d_1874_v2_ = _dafny.Map({d_1872_v0_: d_1873_v1_})
        hi8_ = (_dafny.MultiSet([(self).f26, -980, (0) - ((self).f26), 31, (self).f34])).cardinality
        for d_1875_i0_ in range(len(d_1874_v2_), hi8_):
            d_1876_v3_: bool
            d_1876_v3_ = True
            d_1877_v4_: _dafny.Map
            d_1877_v4_ = _dafny.Map({d_1876_v3_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xehn")))})
            d_1878_v5_: D1
            d_1878_v5_ = D1_DC5(p0)
            d_1879_v6_: _dafny.Map
            d_1879_v6_ = _dafny.Map({(self).f26: d_1878_v5_})
            d_1880_v7_: D8
            d_1880_v7_ = D8_DC25((39) <= (d_1875_i0_), d_1877_v4_, d_1873_v1_, d_1879_v6_, (self).f34)
            d_1880_v7_ = d_1880_v7_
            if d_1876_v3_:
                d_1881_v8_: _dafny.MultiSet
                d_1881_v8_ = _dafny.MultiSet([323])
                d_1882_v9_: _dafny.Map
                d_1882_v9_ = _dafny.Map({len(((self).f27).set(default__.safeIndex((self).f34, len((self).f27)), d_1873_v1_)): (self).fm7((self).f26, True, globalState)})
                d_1883_v10_: D0
                d_1883_v10_ = D0_DC1((self).f27, (self).f34, len((self).f27))
                d_1884_v11_: _dafny.MultiSet
                d_1884_v11_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1873_v1_ for d_1885_i1_ in range(default__.abs(188))]), default__.fm43((self).fm8(d_1881_v8_, d_1876_v3_, (self).f26, 411, globalState), len(d_1882_v9_), globalState), ((self).f27) + ((d_1883_v10_).cf1), (d_1883_v10_).cf1])
                d_1884_v11_ = d_1884_v11_
                d_1886_v12_: _dafny.Seq
                d_1886_v12_ = _dafny.SeqWithoutIsStrInference([True, d_1876_v3_, d_1876_v3_])
                d_1887_v13_: D2
                d_1887_v13_ = D2_DC7(d_1886_v12_)
                pat_let_tv44_ = d_1876_v3_
                pat_let_tv45_ = d_1876_v3_
                d_1888_v14_: _dafny.Map
                def iife154_(_pat_let31_0):
                    def iife155_(d_1889_dt__update__tmp_h0_):
                        def iife156_(_pat_let32_0):
                            def iife157_(d_1890_dt__update_hcf12_h0_):
                                return D2_DC7(d_1890_dt__update_hcf12_h0_)
                            return iife157_(_pat_let32_0)
                        return iife156_(_dafny.SeqWithoutIsStrInference([pat_let_tv44_, pat_let_tv45_]))
                    return iife155_(_pat_let31_0)
                d_1888_v14_ = _dafny.Map({iife154_(d_1887_v13_): (d_1883_v10_).cf1})
                d_1888_v14_ = (d_1888_v14_).set(d_1887_v13_, _dafny.SeqWithoutIsStrInference([d_1873_v1_ for d_1891_i2_ in range(default__.abs(-973))]))
                d_1892_v15_: _dafny.Array
                nw288_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                d_1892_v15_ = nw288_
                d_1892_v15_ = d_1892_v15_
                d_1893_v16_: _dafny.Set
                d_1893_v16_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktjv"))})
                rhs314_ = d_1875_i0_
                rhs315_ = d_1893_v16_
                lhs286_ = globalState
                lhs286_.f8 = rhs314_
                d_1893_v16_ = rhs315_
                d_1894_v17_: _dafny.Seq
                d_1894_v17_ = _dafny.SeqWithoutIsStrInference([d_1886_v12_, d_1886_v12_])
                (globalState).f25 = (d_1894_v17_) < ((d_1894_v17_) + (d_1894_v17_))
            elif True:
                (globalState).f13 = (self).f27
                d_1895_v18_: C0
                nw289_ = C0()
                nw289_.ctor__()
                d_1895_v18_ = nw289_
                (globalState).f25 = d_1876_v3_
                d_1896_v19_: _dafny.MultiSet
                d_1896_v19_ = _dafny.MultiSet([d_1875_i0_, len((self).f27)])
                (globalState).f24 = ((d_1896_v19_ if d_1876_v3_ else default__.fm44(d_1876_v3_, d_1876_v3_, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1876_v3_, d_1876_v3_])), (self).fm7((self).f26, d_1876_v3_, globalState), (d_1895_v18_).fm16((self).f26, d_1875_i0_, d_1875_i0_, d_1876_v3_, globalState), (0) - (d_1875_i0_), d_1875_i0_]), globalState))).cardinality
                d_1897_v20_: _dafny.Array
                nw290_ = _dafny.Array(_dafny.Map({}), 25)
                d_1897_v20_ = nw290_
                index300_ = default__.safeIndex(517, (d_1897_v20_).length(0))
                (d_1897_v20_)[index300_] = (d_1874_v2_) | (_dafny.Map({d_1872_v0_: d_1873_v1_}))
                index301_ = default__.safeIndex(517, (d_1897_v20_).length(0))
                (d_1897_v20_)[index301_] = d_1874_v2_
            (globalState).f8 = (default__.safeModuloInt((0) - (-541), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhkdo"))))) - (d_1875_i0_)
            d_1898_v21_: C1
            nw291_ = C1()
            nw291_.ctor__()
            d_1898_v21_ = nw291_
            d_1899_v22_: _dafny.Array
            nw292_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
            d_1899_v22_ = nw292_
            d_1900_v23_: _dafny.Seq
            d_1900_v23_ = _dafny.SeqWithoutIsStrInference([d_1898_v21_, d_1898_v21_])
            rhs316_ = (d_1900_v23_)[default__.safeIndex(len((_dafny.Set({d_1872_v0_})) - (_dafny.Set({d_1872_v0_}))), len(d_1900_v23_))]
            rhs317_ = d_1899_v22_
            d_1898_v21_ = rhs316_
            d_1899_v22_ = rhs317_
        d_1901_v24_: bool
        d_1901_v24_ = True
        if d_1901_v24_:
            d_1902_v25_: _dafny.Map
            d_1902_v25_ = _dafny.Map({d_1901_v24_: (self).f26})
            d_1902_v25_ = d_1902_v25_
            d_1903_v26_: _dafny.Array
            nw293_ = _dafny.Array(int(0), 5)
            d_1903_v26_ = nw293_
            d_1904_v27_: _dafny.Array
            nw294_ = _dafny.Array(None, 12)
            nw294_[int(0)] = d_1903_v26_
            nw294_[int(1)] = d_1903_v26_
            nw294_[int(2)] = d_1903_v26_
            nw294_[int(3)] = d_1903_v26_
            nw294_[int(4)] = d_1903_v26_
            nw294_[int(5)] = d_1903_v26_
            nw294_[int(6)] = d_1903_v26_
            nw294_[int(7)] = (d_1903_v26_ if d_1901_v24_ else d_1903_v26_)
            nw294_[int(8)] = d_1903_v26_
            nw294_[int(9)] = d_1903_v26_
            nw294_[int(10)] = d_1903_v26_
            nw294_[int(11)] = d_1903_v26_
            d_1904_v27_ = nw294_
            nw295_ = _dafny.Array(None, 3)
            nw295_[int(0)] = d_1903_v26_
            nw295_[int(1)] = d_1903_v26_
            nw295_[int(2)] = d_1903_v26_
            d_1904_v27_ = nw295_
            d_1905_v28_: _dafny.Seq
            d_1905_v28_ = _dafny.SeqWithoutIsStrInference([d_1901_v24_])
            d_1906_v29_: D2
            d_1906_v29_ = D2_DC7((d_1905_v28_).set(default__.safeIndex((self).f34, len(d_1905_v28_)), d_1901_v24_))
            d_1907_v30_: _dafny.Seq
            d_1907_v30_ = _dafny.SeqWithoutIsStrInference([d_1906_v29_, d_1906_v29_, d_1906_v29_, d_1906_v29_])
            (globalState).f8 = len(d_1907_v30_)
            d_1908_v31_: D3
            d_1908_v31_ = D3_DC12(d_1901_v24_, (self).f26, d_1901_v24_)
            d_1909_v33_: _dafny.MultiSet
            d_1909_v33_ = _dafny.MultiSet([(self).f26, (self).f26])
            d_1910_v35_: _dafny.Map
            d_1910_v35_ = _dafny.Map({(self).f26: d_1901_v24_})
            d_1911_v36_: _dafny.Map
            d_1911_v36_ = _dafny.Map({d_1905_v28_: d_1901_v24_})
            d_1912_v37_: _dafny.Seq
            d_1912_v37_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).fm6(d_1910_v35_, (self).f34, _dafny.SeqWithoutIsStrInference([(self).f34, 683]), d_1911_v36_, globalState)), ((d_1902_v25_)[d_1901_v24_] if (d_1901_v24_) in (d_1902_v25_) else (self).f34), (self).f34, (self).f26])
            def iife158_():
                coll92_ = _dafny.Map()
                compr_92_: int
                for compr_92_ in (d_1909_v33_).Elements:
                    d_1913_v32_: int = compr_92_
                    if (d_1913_v32_) in (d_1909_v33_):
                        coll92_[(d_1913_v32_) * ((self).f34)] = d_1901_v24_
                return _dafny.Map(coll92_)
            def iife159_():
                coll93_ = _dafny.Map()
                compr_93_: int
                for compr_93_ in _dafny.IntegerRange(975, 905):
                    d_1914_v34_: int = compr_93_
                    if ((975) <= (d_1914_v34_)) and ((d_1914_v34_) < (905)):
                        coll93_[(d_1914_v34_) - ((d_1909_v33_).cardinality)] = (self).f26
                return _dafny.Map(coll93_)
            (globalState).f17 = (((self).f34) - ((d_1908_v31_).cf26)) + ((self).fm6(iife158_()
            , len(iife159_()
            ), d_1912_v37_, d_1911_v36_, globalState))
            d_1915_v38_: _dafny.Seq
            d_1915_v38_ = _dafny.SeqWithoutIsStrInference([d_1912_v37_, d_1912_v37_])
            d_1912_v37_ = ((d_1915_v38_)[default__.safeIndex((self).f34, len(d_1915_v38_))]) + (d_1912_v37_)
        elif True:
            d_1916_v39_: _dafny.MultiSet
            d_1916_v39_ = _dafny.MultiSet([(self).f34, (self).f34])
            d_1917_v40_: _dafny.Map
            d_1917_v40_ = _dafny.Map({d_1916_v39_: d_1901_v24_})
            d_1918_v41_: _dafny.Seq
            d_1918_v41_ = _dafny.SeqWithoutIsStrInference([(self).f34])
            d_1917_v40_ = (d_1917_v40_).set(default__.fm44(d_1901_v24_, False, d_1918_v41_, globalState), d_1901_v24_)
            d_1919_v42_: _dafny.Seq
            d_1919_v42_ = _dafny.SeqWithoutIsStrInference([not(d_1901_v24_), True, d_1901_v24_])
            d_1920_v44_: _dafny.Map
            d_1920_v44_ = _dafny.Map({-612: (self).f34})
            d_1921_v45_: T2
            nw296_ = C2()
            def iife160_():
                coll94_ = _dafny.Set()
                compr_94_: int
                for compr_94_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1873_v1_ for d_1922_i4_ in range(default__.abs(500))])): d_1901_v24_})).keys.Elements:
                    d_1923_v43_: int = compr_94_
                    if (d_1923_v43_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1873_v1_ for d_1922_i4_ in range(default__.abs(500))])): d_1901_v24_})):
                        coll94_ = coll94_.union(_dafny.Set([(d_1923_v43_) - ((0) - ((_dafny.MultiSet([not(False)])).cardinality))]))
                return _dafny.Set(coll94_)
            nw296_.ctor__(_dafny.SeqWithoutIsStrInference([len(iife160_()
) for d_1924_i3_ in range(default__.abs(707))]), d_1920_v44_)
            d_1921_v45_ = nw296_
            d_1925_v46_: _dafny.Seq
            d_1925_v46_ = _dafny.SeqWithoutIsStrInference([d_1921_v45_])
            d_1926_v47_: _dafny.Map
            d_1926_v47_ = _dafny.Map({len(d_1925_v46_): (self).f27})
            d_1927_v48_: _dafny.Map
            d_1927_v48_ = _dafny.Map({(d_1919_v42_)[default__.safeIndex(-454, len(d_1919_v42_))]: (d_1916_v39_).intersection(_dafny.MultiSet([(self).f34, (self).f26, len(d_1926_v47_)]))})
            d_1927_v48_ = d_1927_v48_
            d_1928_v49_: D6
            d_1928_v49_ = D6_DC16((self).f26, d_1901_v24_)
            if (d_1928_v49_).cf32:
                d_1929_v50_: _dafny.Array
                def lambda99_(d_1930_v39_):
                    def lambda100_(d_1931_i5_):
                        return default__.safeDivisionInt(d_1931_i5_, (d_1930_v39_).cardinality)

                    return lambda100_

                init55_ = lambda99_(d_1916_v39_)
                nw297_ = _dafny.Array(None, 27)
                for i0_55_ in range(nw297_.length(0)):
                    nw297_[i0_55_] = init55_(i0_55_)
                d_1929_v50_ = nw297_
                index302_ = default__.safeIndex(831, (d_1929_v50_).length(0))
                (d_1929_v50_)[index302_] = (0) - ((-786) + ((self).f26))
                d_1932_v51_: _dafny.Map
                d_1932_v51_ = _dafny.Map({d_1901_v24_: d_1873_v1_})
                index303_ = default__.safeIndex(831, (d_1929_v50_).length(0))
                (d_1929_v50_)[index303_] = len(d_1932_v51_)
                d_1933_v52_: D11
                d_1933_v52_ = D11_DC34(d_1872_v0_, d_1901_v24_)
                (globalState).f9 = (not(d_1901_v24_) if d_1901_v24_ else (d_1933_v52_).cf67)
                r1 = (self).fm7((self).f26, (p0).isdisjoint(_dafny.Set({d_1901_v24_})), globalState)
                d_1919_v42_ = d_1919_v42_
                d_1934_v54_: _dafny.Map
                d_1934_v54_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1873_v1_ for d_1935_i6_ in range(default__.abs(487))]): (self).f34})
                d_1936_v56_: _dafny.Map
                def iife161_():
                    coll95_ = _dafny.Set()
                    compr_95_: _dafny.Seq
                    for compr_95_ in (d_1934_v54_).keys.Elements:
                        d_1937_v55_: _dafny.Seq = compr_95_
                        if (d_1937_v55_) in (d_1934_v54_):
                            coll95_ = coll95_.union(_dafny.Set([d_1937_v55_]))
                    return _dafny.Set(coll95_)
                d_1936_v56_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([True, (d_1919_v42_)[default__.safeIndex((self).f34, len(d_1919_v42_))]])): iife161_()
                })
                d_1938_v57_: _dafny.Set
                d_1938_v57_ = _dafny.Set({(self).f27})
                d_1939_v58_: _dafny.Map
                def iife162_():
                    coll96_ = _dafny.Map()
                    compr_96_: _dafny.Seq
                    for compr_96_ in ((D17_DC47(((d_1936_v56_)[(d_1929_v50_)[default__.safeIndex(831, (d_1929_v50_).length(0))]] if ((d_1929_v50_)[default__.safeIndex(831, (d_1929_v50_).length(0))]) in (d_1936_v56_) else d_1938_v57_))).cf88).Elements:
                        d_1940_v53_: _dafny.Seq = compr_96_
                        if (d_1940_v53_) in ((D17_DC47(((d_1936_v56_)[(d_1929_v50_)[default__.safeIndex(831, (d_1929_v50_).length(0))]] if ((d_1929_v50_)[default__.safeIndex(831, (d_1929_v50_).length(0))]) in (d_1936_v56_) else d_1938_v57_))).cf88):
                            coll96_[d_1940_v53_] = (self).f27
                    return _dafny.Map(coll96_)
                d_1939_v58_ = _dafny.Map({(d_1929_v50_)[default__.safeIndex(831, (d_1929_v50_).length(0))]: iife162_()
                })
                d_1941_v59_: _dafny.Map
                d_1941_v59_ = _dafny.Map({(self).f27: ((self).f27).set(default__.safeIndex((self).f34, len((self).f27)), ((self).f27)[default__.safeIndex((d_1929_v50_)[default__.safeIndex(831, (d_1929_v50_).length(0))], len((self).f27))])})
                d_1939_v58_ = (d_1939_v58_).set(743, d_1941_v59_)
            elif True:
                d_1920_v44_ = (d_1920_v44_).set((self).f26, (self).f26)
                d_1942_v60_: _dafny.Array
                nw298_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
                d_1942_v60_ = nw298_
                index304_ = default__.safeIndex(768, (d_1942_v60_).length(0))
                (d_1942_v60_)[index304_] = ((self).f27) + ((self).f27)
                index305_ = default__.safeIndex(768, (d_1942_v60_).length(0))
                (d_1942_v60_)[index305_] = ((self).f27) + (((self).f27) + ((self).f27))
                r1 = (self).f34
                d_1943_v61_: C1
                nw299_ = C1()
                nw299_.ctor__()
                d_1943_v61_ = nw299_
                d_1944_v62_: _dafny.Map
                d_1944_v62_ = _dafny.Map({(self).f34: not(d_1901_v24_)})
                d_1944_v62_ = (d_1944_v62_).set((self).f34, d_1901_v24_)
            d_1873_v1_ = d_1873_v1_
            d_1945_v63_: _dafny.Map
            d_1945_v63_ = _dafny.Map({d_1901_v24_: d_1901_v24_})
            d_1945_v63_ = (d_1945_v63_).set(d_1901_v24_, d_1901_v24_)
        d_1946_v64_: _dafny.Set
        d_1946_v64_ = _dafny.Set({d_1872_v0_})
        d_1947_v65_: D18
        d_1947_v65_ = D18_DC49(d_1946_v64_)
        d_1946_v64_ = (d_1946_v64_ if True else (d_1947_v65_).cf92)
        d_1948_i7_: int
        d_1948_i7_ = 0
        with _dafny.label("3"):
            while not(d_1901_v24_):
                with _dafny.c_label("3"):
                    if (d_1948_i7_) >= (100):
                        raise _dafny.Break("3")
                    d_1948_i7_ = (d_1948_i7_) + (1)
                    if d_1901_v24_:
                        d_1949_v66_: D0
                        d_1949_v66_ = D0_DC4(D0_DC0(len((self).f27)))
                        d_1950_v67_: D0
                        d_1950_v67_ = D0_DC0((self).f34)
                        d_1951_v68_: D0
                        d_1951_v68_ = D0_DC4(d_1950_v67_)
                        d_1952_v69_: _dafny.Seq
                        d_1952_v69_ = _dafny.SeqWithoutIsStrInference([d_1951_v68_, d_1950_v67_])
                        d_1953_v71_: _dafny.Seq
                        d_1953_v71_ = _dafny.SeqWithoutIsStrInference([(self).f34])
                        d_1954_v72_: _dafny.Seq
                        d_1954_v72_ = _dafny.SeqWithoutIsStrInference([d_1901_v24_])
                        d_1955_v73_: _dafny.Map
                        d_1955_v73_ = _dafny.Map({(self).f26: d_1954_v72_})
                        d_1956_v74_: _dafny.Map
                        d_1956_v74_ = _dafny.Map({((d_1955_v73_)[(self).f34] if ((self).f34) in (d_1955_v73_) else d_1954_v72_): d_1901_v24_})
                        def iife163_():
                            coll97_ = _dafny.Map()
                            compr_97_: int
                            for compr_97_ in _dafny.IntegerRange(719, 936):
                                d_1957_v70_: int = compr_97_
                                if ((719) <= (d_1957_v70_)) and ((d_1957_v70_) < (936)):
                                    coll97_[(d_1957_v70_) - ((self).f26)] = d_1901_v24_
                            return _dafny.Map(coll97_)
                        d_1949_v66_ = D0_DC4((d_1952_v69_)[default__.safeIndex((self).fm6(iife163_()
, (self).f34, d_1953_v71_, d_1956_v74_, globalState), len(d_1952_v69_))])
                        rhs318_ = len((d_1953_v71_) + (d_1953_v71_))
                        rhs319_ = default__.safeDivisionInt(((self).f26) + ((self).f34), (self).f34)
                        lhs287_ = globalState
                        lhs288_ = globalState
                        lhs287_.f17 = rhs318_
                        lhs288_.f8 = rhs319_
                        d_1958_v75_: _dafny.Map
                        d_1958_v75_ = _dafny.Map({d_1901_v24_: _dafny.SeqWithoutIsStrInference([d_1873_v1_ for d_1959_i8_ in range(default__.abs(-901))])})
                        d_1960_v76_: _dafny.MultiSet
                        d_1960_v76_ = _dafny.MultiSet([len(d_1958_v75_), (self).f34])
                        d_1961_v77_: _dafny.Map
                        d_1961_v77_ = _dafny.Map({(self).f26: ((d_1960_v76_)[(self).f34] if ((self).f34) in (d_1960_v76_) else (self).fm7((self).f26, d_1901_v24_, globalState))})
                        d_1962_v78_: _dafny.Map
                        d_1962_v78_ = _dafny.Map({(0) - ((self).f34): (d_1961_v77_) | (d_1961_v77_)})
                        d_1963_v79_: _dafny.Map
                        d_1963_v79_ = _dafny.Map({(self).f34: d_1958_v75_})
                        d_1964_v80_: _dafny.Map
                        d_1964_v80_ = _dafny.Map({(self).f34: (self).f27})
                        d_1965_v81_: _dafny.Map
                        d_1965_v81_ = _dafny.Map({211: d_1901_v24_})
                        rhs320_ = (d_1962_v78_) | ((d_1962_v78_) | (d_1962_v78_))
                        rhs321_ = ((d_1963_v79_)[len(((d_1964_v80_)[(self).fm6(d_1965_v81_, -216, d_1953_v71_, d_1956_v74_, globalState)] if ((self).fm6(d_1965_v81_, -216, d_1953_v71_, d_1956_v74_, globalState)) in (d_1964_v80_) else (self).f27))] if (len(((d_1964_v80_)[(self).fm6(d_1965_v81_, -216, d_1953_v71_, d_1956_v74_, globalState)] if ((self).fm6(d_1965_v81_, -216, d_1953_v71_, d_1956_v74_, globalState)) in (d_1964_v80_) else (self).f27))) in (d_1963_v79_) else (d_1958_v75_) | (_dafny.Map({True: (self).f27})))
                        d_1962_v78_ = rhs320_
                        d_1958_v75_ = rhs321_
                        d_1966_v82_: _dafny.Array
                        nw300_ = _dafny.Array(_dafny.Map({}), 27)
                        d_1966_v82_ = nw300_
                        d_1967_v83_: _dafny.Map
                        d_1967_v83_ = _dafny.Map({d_1901_v24_: len(d_1954_v72_)})
                        d_1968_v84_: _dafny.Set
                        d_1968_v84_ = _dafny.Set({len(_dafny.Map({425: d_1967_v83_})), (self).f34, len((self).f27)})
                        d_1969_v85_: _dafny.Map
                        d_1969_v85_ = _dafny.Map({d_1901_v24_: d_1967_v83_})
                        d_1970_v86_: D0
                        d_1970_v86_ = D0_DC0((self).f34)
                        d_1971_v87_: _dafny.Map
                        d_1971_v87_ = _dafny.Map({(self).f34: d_1970_v86_})
                        d_1972_v88_: _dafny.Map
                        d_1972_v88_ = _dafny.Map({(0) - ((self).f26): d_1971_v87_})
                        d_1973_v89_: D19
                        d_1973_v89_ = D19_DC52(((d_1972_v88_)[-646] if (-646) in (d_1972_v88_) else d_1971_v87_))
                        d_1974_v90_: _dafny.Array
                        nw301_ = _dafny.Array(None, 27)
                        nw301_[int(0)] = default__.safeDivisionInt((self).f26, len((self).f27))
                        nw301_[int(1)] = default__.safeModuloInt((self).f34, (self).f26)
                        nw301_[int(2)] = (self).f34
                        nw301_[int(3)] = (self).f26
                        nw301_[int(4)] = (self).f26
                        nw301_[int(5)] = len(d_1968_v84_)
                        nw301_[int(6)] = (self).f26
                        nw301_[int(7)] = len((_dafny.Map({d_1901_v24_: d_1967_v83_})) | (d_1969_v85_))
                        nw301_[int(8)] = 544
                        nw301_[int(9)] = (self).f34
                        nw301_[int(10)] = (self).f26
                        nw301_[int(11)] = (self).f26
                        nw301_[int(12)] = default__.safeDivisionInt((self).f34, (self).f34)
                        nw301_[int(13)] = (self).f26
                        nw301_[int(14)] = (0) - ((self).f26)
                        nw301_[int(15)] = 551
                        nw301_[int(16)] = (self).f34
                        nw301_[int(17)] = len((p0 if d_1901_v24_ else p0))
                        nw301_[int(18)] = ((d_1960_v76_)[(0) - ((self).fm6(d_1965_v81_, (self).f26, d_1953_v71_, _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1901_v24_]): d_1901_v24_}), globalState))] if ((0) - ((self).fm6(d_1965_v81_, (self).f26, d_1953_v71_, _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1901_v24_]): d_1901_v24_}), globalState))) in (d_1960_v76_) else len(d_1946_v64_))
                        nw301_[int(19)] = -612
                        nw301_[int(20)] = (self).f34
                        nw301_[int(21)] = (self).f34
                        nw301_[int(22)] = len((d_1973_v89_).cf99)
                        nw301_[int(23)] = (self).f26
                        nw301_[int(24)] = ((self).f34) * ((self).f34)
                        nw301_[int(25)] = ((self).f34) * (317)
                        nw301_[int(26)] = ((self).f34) + ((self).f34)
                        d_1974_v90_ = nw301_
                        index306_ = default__.safeIndex(777, (d_1974_v90_).length(0))
                        (d_1974_v90_)[index306_] = default__.safeDivisionInt(281, (self).f26)
                        index307_ = default__.safeIndex(777, (d_1974_v90_).length(0))
                        rhs322_ = (d_1968_v84_).ispropersubset((_dafny.Set({(self).f26, 335, len(d_1954_v72_), len(d_1965_v81_), (self).f34})) | (_dafny.Set({(self).f26, (self).f34})))
                        rhs323_ = d_1966_v82_
                        rhs324_ = (0) - (default__.safeDivisionInt(len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) + ((self).f27)).set(default__.safeIndex((self).f26, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) + ((self).f27))), d_1873_v1_)), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khcfajjp"))))))
                        rhs325_ = _dafny.Set({(self).f26})
                        lhs289_ = globalState
                        lhs290_ = d_1974_v90_
                        lhs291_ = default__.safeIndex(777, (d_1974_v90_).length(0))
                        lhs289_.f25 = rhs322_
                        d_1966_v82_ = rhs323_
                        lhs290_[lhs291_] = rhs324_
                        d_1968_v84_ = rhs325_
                        d_1975_v91_: C1
                        nw302_ = C1()
                        nw302_.ctor__()
                        d_1975_v91_ = nw302_
                    elif True:
                        index308_ = default__.safeIndex(723, (d_1872_v0_).length(0))
                        (d_1872_v0_)[index308_] = d_1901_v24_
                        d_1976_v92_: _dafny.Seq
                        d_1976_v92_ = _dafny.SeqWithoutIsStrInference([d_1901_v24_])
                        d_1977_v93_: _dafny.Seq
                        d_1977_v93_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f34, (self).f26, (self).f26, (0) - (len(d_1976_v92_))])
                        d_1978_v94_: D6
                        d_1978_v94_ = D6_DC15(d_1977_v93_)
                        d_1979_v95_: _dafny.Array
                        nw303_ = _dafny.Array(None, 3)
                        nw303_[int(0)] = d_1978_v94_
                        nw303_[int(1)] = d_1978_v94_
                        nw303_[int(2)] = d_1978_v94_
                        d_1979_v95_ = nw303_
                        d_1980_v96_: _dafny.Map
                        d_1980_v96_ = _dafny.Map({d_1979_v95_: d_1901_v24_})
                        index309_ = default__.safeIndex(723, (d_1872_v0_).length(0))
                        (d_1872_v0_)[index309_] = ((d_1980_v96_)[d_1979_v95_] if (d_1979_v95_) in (d_1980_v96_) else (False) == (d_1901_v24_))
                        d_1981_v97_: _dafny.Map
                        d_1981_v97_ = _dafny.Map({(self).f27: not(not(True))})
                        d_1982_v98_: _dafny.Map
                        d_1982_v98_ = d_1981_v97_
                        d_1982_v98_ = d_1982_v98_
                        (globalState).f9 = (d_1901_v24_ if (d_1872_v0_)[default__.safeIndex(723, (d_1872_v0_).length(0))] else (d_1872_v0_)[default__.safeIndex(723, (d_1872_v0_).length(0))])
                        d_1983_v99_: _dafny.Map
                        d_1983_v99_ = _dafny.Map({len(_dafny.Set({d_1901_v24_})): len((self).f27)})
                        d_1984_v100_: T2
                        nw304_ = C2()
                        nw304_.ctor__(d_1977_v93_, d_1983_v99_)
                        d_1984_v100_ = nw304_
                        d_1985_v101_: D17
                        d_1985_v101_ = D17_DC48(d_1984_v100_, (self).f34, _dafny.CodePoint('a'))
                        d_1985_v101_ = d_1985_v101_
                        d_1986_v102_: _dafny.Array
                        nw305_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                        d_1986_v102_ = nw305_
                        index310_ = default__.safeIndex(505, (d_1986_v102_).length(0))
                        (d_1986_v102_)[index310_] = d_1873_v1_
                        index311_ = default__.safeIndex(505, (d_1986_v102_).length(0))
                        (d_1986_v102_)[index311_] = d_1873_v1_
                    d_1987_v103_: _dafny.Seq
                    d_1987_v103_ = _dafny.SeqWithoutIsStrInference([d_1901_v24_])
                    d_1987_v103_ = d_1987_v103_
                    (globalState).f23 = (self).f27
                    d_1988_v104_: C0
                    nw306_ = C0()
                    nw306_.ctor__()
                    d_1988_v104_ = nw306_
                    d_1989_v105_: D17
                    d_1989_v105_ = D17_DC47(_dafny.Set({(self).f27, (self).f27, (self).f27}))
                    d_1990_v106_: _dafny.Map
                    d_1990_v106_ = _dafny.Map({d_1988_v104_: d_1989_v105_})
                    d_1990_v106_ = (d_1990_v106_).set(d_1988_v104_, d_1989_v105_)
                    pass
            pass
        d_1991_v107_: _dafny.MultiSet
        d_1991_v107_ = _dafny.MultiSet([not(True), False, not(d_1901_v24_)])
        d_1992_i9_: int
        d_1992_i9_ = 0
        with _dafny.label("4"):
            while not(((D15_DC42(_dafny.MultiSet([d_1901_v24_]))).cf80).isdisjoint(d_1991_v107_)):
                with _dafny.c_label("4"):
                    if (d_1992_i9_) >= (100):
                        raise _dafny.Break("4")
                    d_1992_i9_ = (d_1992_i9_) + (1)
                    d_1993_v108_: _dafny.Map
                    d_1993_v108_ = _dafny.Map({d_1901_v24_: (self).f26})
                    d_1994_v109_: _dafny.Map
                    d_1994_v109_ = _dafny.Map({(self).f26: 736})
                    d_1995_v110_: _dafny.Seq
                    d_1995_v110_ = _dafny.SeqWithoutIsStrInference([not(d_1901_v24_), (self).fm8(_dafny.MultiSet(default__.fm45(530, d_1993_v108_, d_1994_v109_, globalState)), d_1901_v24_, 30, (self).f26, globalState)])
                    d_1996_v111_: _dafny.Map
                    d_1996_v111_ = _dafny.Map({d_1873_v1_: (d_1995_v110_) + (d_1995_v110_)})
                    d_1996_v111_ = (d_1996_v111_).set(_dafny.CodePoint('u'), d_1995_v110_)
                    d_1997_v112_: _dafny.Array
                    nw307_ = _dafny.Array(D0.default()(), 13)
                    d_1997_v112_ = nw307_
                    d_1998_v113_: D1
                    d_1998_v113_ = D1_DC6(not(d_1901_v24_), default__.fm42(default__.fm46(d_1901_v24_, len((self).f27), d_1901_v24_, globalState), globalState), (21) - ((self).f26))
                    rhs326_ = not(d_1901_v24_)
                    rhs327_ = d_1997_v112_
                    rhs328_ = not (d_1901_v24_) or ((p0).issubset(p0))
                    rhs329_ = d_1901_v24_
                    rhs330_ = D1_DC6((d_1901_v24_ if d_1901_v24_ else not(d_1901_v24_)), d_1873_v1_, (self).f26)
                    r3 = rhs326_
                    d_1997_v112_ = rhs327_
                    r3 = rhs328_
                    d_1901_v24_ = rhs329_
                    d_1998_v113_ = rhs330_
                    d_1993_v108_ = (d_1993_v108_) | ((d_1993_v108_ if d_1901_v24_ else d_1993_v108_))
                    if d_1901_v24_:
                        d_1999_v114_: _dafny.Map
                        d_1999_v114_ = _dafny.Map({(self).fm7(33, d_1901_v24_, globalState): d_1901_v24_})
                        d_1999_v114_ = (d_1999_v114_).set(956, True)
                        d_2000_v115_: C1
                        nw308_ = C1()
                        nw308_.ctor__()
                        d_2000_v115_ = nw308_
                        d_2001_v116_: _dafny.MultiSet
                        d_2001_v116_ = _dafny.MultiSet([(self).f34, (self).f34])
                        d_2002_v117_: D16
                        d_2002_v117_ = D16_DC45(len(d_1999_v114_), True, (d_2001_v116_).cardinality, (self).f34, d_1901_v24_)
                        d_2003_v118_: _dafny.Map
                        d_2003_v118_ = _dafny.Map({d_1901_v24_: True})
                        d_2004_v119_: _dafny.Map
                        d_2004_v119_ = _dafny.Map({d_1901_v24_: d_2003_v118_})
                        d_2005_v120_: _dafny.Seq
                        d_2005_v120_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xltohr"))), len(d_2004_v119_)])
                        d_2006_v121_: _dafny.MultiSet
                        d_2006_v121_ = _dafny.MultiSet([d_2005_v120_, _dafny.SeqWithoutIsStrInference([(self).f26]), d_2005_v120_, d_2005_v120_, d_2005_v120_])
                        d_2007_v122_: _dafny.Array
                        nw309_ = _dafny.Array(None, 19)
                        nw309_[int(0)] = (self).f34
                        nw309_[int(1)] = (self).f34
                        nw309_[int(2)] = (d_2002_v117_).cf84
                        nw309_[int(3)] = (self).f34
                        nw309_[int(4)] = ((d_2006_v121_)[d_2005_v120_] if (d_2005_v120_) in (d_2006_v121_) else (self).f26)
                        nw309_[int(5)] = -801
                        nw309_[int(6)] = (self).f34
                        nw309_[int(7)] = (0) - (len((self).f27))
                        nw309_[int(8)] = (self).f26
                        nw309_[int(9)] = (self).f34
                        nw309_[int(10)] = default__.safeModuloInt((self).f26, (self).f34)
                        nw309_[int(11)] = (d_2000_v115_).fm7(((d_1993_v108_)[d_1901_v24_] if (d_1901_v24_) in (d_1993_v108_) else (self).f26), d_1901_v24_, globalState)
                        nw309_[int(12)] = (self).f34
                        nw309_[int(13)] = ((self).f26 if d_1901_v24_ else (self).f34)
                        nw309_[int(14)] = (self).f34
                        nw309_[int(15)] = len(d_2005_v120_)
                        nw309_[int(16)] = (self).f34
                        nw309_[int(17)] = (self).f34
                        nw309_[int(18)] = ((_dafny.MultiSet([d_2005_v120_])) - ((d_2006_v121_).set(_dafny.SeqWithoutIsStrInference([(self).f34 for d_2008_i10_ in range(default__.abs(53))]), default__.abs((self).f34)))).cardinality
                        d_2007_v122_ = nw309_
                        index312_ = default__.safeIndex(574, (d_2007_v122_).length(0))
                        (d_2007_v122_)[index312_] = len((self).f27)
                        index313_ = default__.safeIndex(574, (d_2007_v122_).length(0))
                        (d_2007_v122_)[index313_] = (self).f34
                        (globalState).f25 = (769) >= ((self).f34)
                        (globalState).f9 = (D11_DC34(d_1872_v0_, d_1901_v24_)).cf67
                    elif True:
                        index314_ = default__.safeIndex(644, (d_1872_v0_).length(0))
                        (d_1872_v0_)[index314_] = d_1901_v24_
                        index315_ = default__.safeIndex(644, (d_1872_v0_).length(0))
                        (d_1872_v0_)[index315_] = d_1901_v24_
                        d_1873_v1_ = (d_1998_v113_).cf10
                        d_2009_v123_: _dafny.Array
                        nw310_ = _dafny.Array(int(0), 8)
                        d_2009_v123_ = nw310_
                        index316_ = default__.safeIndex(587, (d_2009_v123_).length(0))
                        (d_2009_v123_)[index316_] = (0) - ((self).f34)
                        index317_ = default__.safeIndex(587, (d_2009_v123_).length(0))
                        (d_2009_v123_)[index317_] = 736
                        index318_ = default__.safeIndex(587, (d_2009_v123_).length(0))
                        index319_ = default__.safeIndex(587, (d_2009_v123_).length(0))
                        rhs331_ = (d_2009_v123_ if (d_1872_v0_)[default__.safeIndex(644, (d_1872_v0_).length(0))] else d_2009_v123_)
                        rhs332_ = ((self).f26) + (default__.safeDivisionInt((self).f34, (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f27, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltlyxnx")), ((self).f27).set(default__.safeIndex(96, len((self).f27)), d_1873_v1_), default__.fm43((d_1872_v0_)[default__.safeIndex(644, (d_1872_v0_).length(0))], (self).f34, globalState), (self).f27])))))
                        rhs333_ = len(((self).f27) + ((self).f27))
                        rhs334_ = (d_2009_v123_)[default__.safeIndex(587, (d_2009_v123_).length(0))]
                        lhs292_ = globalState
                        lhs293_ = d_2009_v123_
                        lhs294_ = default__.safeIndex(587, (d_2009_v123_).length(0))
                        lhs295_ = d_2009_v123_
                        lhs296_ = default__.safeIndex(587, (d_2009_v123_).length(0))
                        d_2009_v123_ = rhs331_
                        lhs292_.f5 = rhs332_
                        lhs293_[lhs294_] = rhs333_
                        lhs295_[lhs296_] = rhs334_
                        d_1873_v1_ = d_1873_v1_
                    pass
            pass
        (globalState).f0 = (self).f34
        r0 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ioocdnpt"))
        r1 = (self).f26
        d_2010_v124_: D2
        d_2010_v124_ = D2_DC8(d_1901_v24_, (self).f34, (self).f34)
        pat_let_tv46_ = d_1901_v24_
        pat_let_tv47_ = d_1901_v24_
        pat_let_tv48_ = d_1901_v24_
        pat_let_tv49_ = d_1901_v24_
        def lambda101_(source30_):
            if source30_.is_DC8:
                d_2011___mcc_h0_ = source30_.cf13
                d_2012___mcc_h1_ = source30_.cf14
                d_2013___mcc_h2_ = source30_.cf15
                d_2014_cf15_ = d_2013___mcc_h2_
                d_2015_cf14_ = d_2012___mcc_h1_
                d_2016_cf13_ = d_2011___mcc_h0_
                def iife164_():
                    coll98_ = _dafny.Map()
                    compr_98_: int
                    for compr_98_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2014_cf15_]))).Elements:
                        d_2017_v125_: int = compr_98_
                        if (d_2017_v125_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2014_cf15_]))):
                            def iife165_():
                                coll99_ = _dafny.Map()
                                compr_99_: int
                                for compr_99_ in _dafny.IntegerRange(-533, 417):
                                    d_2018_v126_: int = compr_99_
                                    if ((-533) <= (d_2018_v126_)) and ((d_2018_v126_) < (417)):
                                        coll99_[(d_2018_v126_) * ((self).f26)] = -841
                                return _dafny.Map(coll99_)
                            coll98_[(d_2017_v125_) - ((self).f26)] = iife165_()

                    return _dafny.Map(coll98_)
                return (iife164_()
                ) | (_dafny.Map({d_2014_cf15_: _dafny.Map({426: -407})}))
            elif source30_.is_DC9:
                d_2019___mcc_h3_ = source30_.cf16
                d_2020___mcc_h4_ = source30_.cf17
                d_2021___mcc_h5_ = source30_.cf18
                d_2022___mcc_h6_ = source30_.cf19
                d_2023_cf19_ = d_2022___mcc_h6_
                d_2024_cf18_ = d_2021___mcc_h5_
                d_2025_cf17_ = d_2020___mcc_h4_
                d_2026_cf16_ = d_2019___mcc_h3_
                return _dafny.Map({d_2026_cf16_: _dafny.Map({len(_dafny.SeqWithoutIsStrInference([pat_let_tv46_])): len((self).f27)})})
            elif True:
                d_2027___mcc_h7_ = source30_.cf12
                d_2028_cf12_ = d_2027___mcc_h7_
                def iife166_():
                    coll100_ = _dafny.Map()
                    compr_100_: int
                    for compr_100_ in _dafny.IntegerRange(-400, -56):
                        d_2029_v127_: int = compr_100_
                        if ((-400) <= (d_2029_v127_)) and ((d_2029_v127_) < (-56)):
                            coll100_[(d_2029_v127_) - ((self).f34)] = _dafny.Map({(self).f34: len((D10_DC31(pat_let_tv47_, _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len((self).f27): pat_let_tv48_})) for d_2030_i11_ in range(default__.abs(-128))]), pat_let_tv49_, (self).f34)).cf59)})
                    return _dafny.Map(coll100_)
                def iife167_():
                    coll101_ = _dafny.Map()
                    compr_101_: int
                    for compr_101_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f26]))).Elements:
                        d_2031_v128_: int = compr_101_
                        if (d_2031_v128_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f26]))):
                            coll101_[default__.safeModuloInt(d_2031_v128_, len((self).f27))] = 407
                    return _dafny.Map(coll101_)
                return (iife166_()
                ) | (_dafny.Map({(self).f34: iife167_()
                }))

        r2 = len(lambda101_(d_2010_v124_))
        r3 = d_1901_v24_
        return r0, r1, r2, r3

    @property
    def f34(self):
        return self._f34

class C4(T1, T0, T2):
    def  __init__(self):
        self._f26: int = int(0)
        self._f27: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f31: _dafny.Map = _dafny.Map({})
        self._f32: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    @property
    def f31(self):
        return self._f31
    def ctor__(self, f32, f26, f27, f31):
        (self)._f32 = f32
        (self)._f26 = f26
        (self)._f27 = f27
        (self)._f31 = f31

    def fm8(self, p0, p1, p2, p3, globalState):
        return (self).f32

    def fm6(self, p0, p1, p2, p3, globalState):
        return ((self).f26) * (((self).f26) * ((self).f26))

    def fm7(self, p0, p1, globalState):
        if True:
            return (self).f26
        elif True:
            return (self).f26

    def fm12(self, p0, p1, p2, p3, globalState):
        def iife168_():
            coll102_ = _dafny.Map()
            compr_102_: _dafny.Set
            for compr_102_ in (_dafny.Map({_dafny.Set({(0) - (len(_dafny.Map({True: _dafny.CodePoint('l')}))), (self).f26, (self).f26}): D0_DC3(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2032_i0_ in range(default__.abs(650))])), (self).f26)})).keys.Elements:
                d_2033_v0_: _dafny.Set = compr_102_
                if (d_2033_v0_) in (_dafny.Map({_dafny.Set({(0) - (len(_dafny.Map({True: _dafny.CodePoint('l')}))), (self).f26, (self).f26}): D0_DC3(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2032_i0_ in range(default__.abs(650))])), (self).f26)})):
                    coll102_[d_2033_v0_] = True
            return _dafny.Map(coll102_)
        return (_dafny.Map({_dafny.Set({(self).f26}): (self).f32})) | (iife168_()
        )

    def fm13(self, p0, p1, p2, globalState):
        return (self).f27

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        if not((_dafny.CodePoint('f')) not in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2034_i0_ in range(default__.abs(674))]))):
            (globalState).f25 = False
            d_2035_v0_: _dafny.Seq
            d_2035_v0_ = _dafny.SeqWithoutIsStrInference([False])
            (globalState).f25 = ((d_2035_v0_)[default__.safeIndex((self).f26, len(d_2035_v0_))]) and ((d_2035_v0_)[default__.safeIndex((self).f26, len(d_2035_v0_))])
            index320_ = default__.safeIndex(640, (p0).length(0))
            (p0)[index320_] = 111
            index321_ = default__.safeIndex(640, (p0).length(0))
            (p0)[index321_] = (self).fm7(((self).f26) * ((self).f26), (self).f32, globalState)
            d_2036_v1_: str
            d_2036_v1_ = _dafny.CodePoint('l')
            d_2037_v2_: _dafny.MultiSet
            d_2037_v2_ = _dafny.MultiSet([d_2036_v1_])
            if (d_2037_v2_).isdisjoint((_dafny.MultiSet((self).f27)).intersection(_dafny.MultiSet([_dafny.CodePoint('u')]))):
                d_2038_v3_: _dafny.MultiSet
                d_2038_v3_ = _dafny.MultiSet([(p0)[default__.safeIndex(640, (p0).length(0))]])
                d_2039_v4_: _dafny.Map
                d_2039_v4_ = _dafny.Map({False: (self).f32})
                d_2040_v5_: _dafny.Set
                d_2040_v5_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([((d_2038_v3_)[(p0)[default__.safeIndex(640, (p0).length(0))]] if ((p0)[default__.safeIndex(640, (p0).length(0))]) in (d_2038_v3_) else (self).f26), len(d_2039_v4_), 811])})
                d_2041_v6_: _dafny.MultiSet
                d_2041_v6_ = _dafny.MultiSet([(d_2040_v5_).ispropersubset(d_2040_v5_)])
                d_2042_v7_: _dafny.Seq
                d_2042_v7_ = _dafny.SeqWithoutIsStrInference([(d_2041_v6_).cardinality])
                d_2043_v9_: _dafny.Map
                def iife169_():
                    coll103_ = _dafny.Map()
                    compr_103_: int
                    for compr_103_ in _dafny.IntegerRange(-412, 574):
                        d_2044_v8_: int = compr_103_
                        if ((-412) <= (d_2044_v8_)) and ((d_2044_v8_) < (574)):
                            coll103_[(d_2044_v8_) - ((p0)[default__.safeIndex(640, (p0).length(0))])] = (p0)[default__.safeIndex(640, (p0).length(0))]
                    return _dafny.Map(coll103_)
                d_2043_v9_ = _dafny.Map({(self).f32: D3_DC11(iife169_()
, (self).f32, (self).f26, (self).f26)})
                d_2045_v10_: _dafny.MultiSet
                d_2045_v10_ = _dafny.MultiSet([d_2043_v9_])
                rhs335_ = (default__.safeDivisionInt(85, (p0)[default__.safeIndex(640, (p0).length(0))])) >= ((d_2042_v7_)[default__.safeIndex((p0)[default__.safeIndex(640, (p0).length(0))], len(d_2042_v7_))])
                rhs336_ = d_2036_v1_
                rhs337_ = (len((self).fm13((self).f32, (self).f27, len((self).f27), globalState))) != ((p0)[default__.safeIndex(640, (p0).length(0))])
                rhs338_ = (d_2041_v6_).set((self).fm8(d_2038_v3_, (d_2035_v0_)[default__.safeIndex((d_2045_v10_).cardinality, len(d_2035_v0_))], (self).f26, 201, globalState), default__.abs((self).f26))
                rhs339_ = (d_2035_v0_ if (False) == ((self).f32) else (_dafny.SeqWithoutIsStrInference([(self).fm8(d_2038_v3_, p1, 154, len(((d_2035_v0_).set(default__.safeIndex((p0)[default__.safeIndex(640, (p0).length(0))], len(d_2035_v0_)), (self).f32)).set(default__.safeIndex((self).f26, len((d_2035_v0_).set(default__.safeIndex((p0)[default__.safeIndex(640, (p0).length(0))], len(d_2035_v0_)), (self).f32))), p1)), globalState), p1])).set(default__.safeIndex((p0)[default__.safeIndex(640, (p0).length(0))], len(_dafny.SeqWithoutIsStrInference([(self).fm8(d_2038_v3_, p1, 154, len(((d_2035_v0_).set(default__.safeIndex((p0)[default__.safeIndex(640, (p0).length(0))], len(d_2035_v0_)), (self).f32)).set(default__.safeIndex((self).f26, len((d_2035_v0_).set(default__.safeIndex((p0)[default__.safeIndex(640, (p0).length(0))], len(d_2035_v0_)), (self).f32))), p1)), globalState), p1]))), not(p1)))
                lhs297_ = globalState
                lhs298_ = globalState
                lhs297_.f25 = rhs335_
                d_2036_v1_ = rhs336_
                lhs298_.f25 = rhs337_
                d_2041_v6_ = rhs338_
                d_2035_v0_ = rhs339_
                d_2046_v11_: _dafny.Set
                d_2046_v11_ = _dafny.Set({(self).f32, (self).f32})
                index322_ = default__.safeIndex(640, (p0).length(0))
                rhs340_ = (d_2035_v0_) + ((_dafny.SeqWithoutIsStrInference([p1, p1, p1])) + (d_2035_v0_))
                rhs341_ = (self).f26
                rhs342_ = (0) - ((0) - (len((d_2046_v11_) - ((d_2046_v11_) | (default__.fm14(globalState))))))
                rhs343_ = default__.safeModuloInt((self).f26, (self).f26)
                rhs344_ = (self).f26
                lhs299_ = p0
                lhs300_ = default__.safeIndex(640, (p0).length(0))
                lhs301_ = globalState
                lhs302_ = globalState
                lhs303_ = globalState
                d_2035_v0_ = rhs340_
                lhs299_[lhs300_] = rhs341_
                lhs301_.f5 = rhs342_
                lhs302_.f0 = rhs343_
                lhs303_.f17 = rhs344_
                r1 = p1
                def lambda102_(d_2047_i1_):
                    return (self).f32

                init56_ = lambda102_
                nw311_ = _dafny.Array(None, 7)
                for i0_56_ in range(nw311_.length(0)):
                    nw311_[i0_56_] = init56_(i0_56_)
                (globalState).f1 = nw311_
                (globalState).f13 = (self).f27
            elif True:
                (globalState).f24 = (p0)[default__.safeIndex(640, (p0).length(0))]
                d_2048_v12_: _dafny.Seq
                d_2048_v12_ = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(640, (p0).length(0))], (p0)[default__.safeIndex(640, (p0).length(0))], (self).f26])
                d_2049_v13_: _dafny.Seq
                d_2049_v13_ = _dafny.SeqWithoutIsStrInference([d_2048_v12_, d_2048_v12_, _dafny.SeqWithoutIsStrInference([(self).f26 for d_2050_i2_ in range(default__.abs(79))])])
                d_2051_v14_: _dafny.Map
                d_2051_v14_ = _dafny.Map({(self).f32: (self).f26})
                d_2052_v15_: _dafny.MultiSet
                d_2052_v15_ = _dafny.MultiSet([(p0)[default__.safeIndex(640, (p0).length(0))], (p0)[default__.safeIndex(640, (p0).length(0))]])
                d_2053_v16_: _dafny.MultiSet
                d_2053_v16_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isc"))])
                d_2054_v17_: C2
                nw312_ = C2()
                nw312_.ctor__((((d_2049_v13_)[default__.safeIndex((p0)[default__.safeIndex(640, (p0).length(0))], len(d_2049_v13_))]) + (_dafny.SeqWithoutIsStrInference([(self).f26, len((self).f27)]))).set(default__.safeIndex((self).f26, len(((d_2049_v13_)[default__.safeIndex((p0)[default__.safeIndex(640, (p0).length(0))], len(d_2049_v13_))]) + (_dafny.SeqWithoutIsStrInference([(self).f26, len((self).f27)])))), 1), ((self).f31).set(((d_2051_v14_)[(self).f32] if ((self).f32) in (d_2051_v14_) else (d_2052_v15_).cardinality), (d_2053_v16_).cardinality))
                d_2054_v17_ = nw312_
                (globalState).f25 = not(p1)
                d_2055_v18_: _dafny.Map
                d_2055_v18_ = _dafny.Map({p1: not((self).f32)})
                d_2056_v19_: _dafny.Map
                d_2056_v19_ = _dafny.Map({(d_2054_v17_).fm7((p0)[default__.safeIndex(640, (p0).length(0))], False, globalState): False})
                d_2055_v18_ = (_dafny.Map({p1: (self).f32})).set(((d_2056_v19_)[(self).f26] if ((self).f26) in (d_2056_v19_) else (self).f32), p1)
                d_2057_v20_: _dafny.Array
                nw313_ = _dafny.Array(None, 2)
                nw313_[int(0)] = (self).f32
                nw313_[int(1)] = p1
                d_2057_v20_ = nw313_
                rhs345_ = p1
                rhs346_ = d_2057_v20_
                lhs304_ = globalState
                lhs305_ = globalState
                lhs304_.f9 = rhs345_
                lhs305_.f1 = rhs346_
            if p1:
                d_2035_v0_ = d_2035_v0_
                d_2058_v21_: _dafny.Array
                nw314_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_2058_v21_ = nw314_
                d_2059_v22_: _dafny.Array
                def lambda103_(d_2060_p0_):
                    def lambda104_(d_2061_i3_):
                        return default__.safeModuloInt(d_2061_i3_, (d_2060_p0_)[default__.safeIndex(640, (d_2060_p0_).length(0))])

                    return lambda104_

                init57_ = lambda103_(p0)
                nw315_ = _dafny.Array(None, 11)
                for i0_57_ in range(nw315_.length(0)):
                    nw315_[i0_57_] = init57_(i0_57_)
                d_2059_v22_ = nw315_
                index323_ = default__.safeIndex(27, (d_2058_v21_).length(0))
                (d_2058_v21_)[index323_] = d_2059_v22_
                index324_ = default__.safeIndex(27, (d_2058_v21_).length(0))
                (d_2058_v21_)[index324_] = d_2059_v22_
                d_2062_v23_: _dafny.Map
                d_2062_v23_ = _dafny.Map({d_2036_v1_: (self).f32})
                d_2062_v23_ = _dafny.Map({d_2036_v1_: (257) not in (_dafny.Map({(p0)[default__.safeIndex(640, (p0).length(0))]: (p0)[default__.safeIndex(640, (p0).length(0))]}))})
                (globalState).f9 = (d_2035_v0_)[default__.safeIndex((845) - ((((self).f31)[(p0)[default__.safeIndex(640, (p0).length(0))]] if ((p0)[default__.safeIndex(640, (p0).length(0))]) in ((self).f31) else (0) - ((self).f26))), len(d_2035_v0_))]
                d_2063_v24_: _dafny.Array
                nw316_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
                d_2063_v24_ = nw316_
                rhs347_ = d_2063_v24_
                rhs348_ = p1
                lhs306_ = globalState
                d_2063_v24_ = rhs347_
                lhs306_.f9 = rhs348_
            elif True:
                d_2064_v25_: _dafny.MultiSet
                d_2064_v25_ = _dafny.MultiSet([p1])
                (globalState).f0 = (default__.safeModuloInt((self).f26, (0) - ((self).f26)) if not (p1) or (p1) else ((self).f26) + ((d_2064_v25_).cardinality))
                index325_ = default__.safeIndex(640, (p0).length(0))
                rhs349_ = (self).f27
                rhs350_ = (p0)[default__.safeIndex(640, (p0).length(0))]
                rhs351_ = (self).f26
                rhs352_ = (self).f27
                lhs307_ = globalState
                lhs308_ = globalState
                lhs309_ = p0
                lhs310_ = default__.safeIndex(640, (p0).length(0))
                lhs311_ = globalState
                lhs307_.f23 = rhs349_
                lhs308_.f24 = rhs350_
                lhs309_[lhs310_] = rhs351_
                lhs311_.f13 = rhs352_
                (globalState).f8 = (((self).f26) * ((0) - ((self).f26))) + ((self).f26)
                (globalState).f23 = ((self).f27) + (((self).f27) + ((self).f27))
                d_2065_v26_: _dafny.Array
                nw317_ = _dafny.Array(D7.default()(), 11)
                d_2065_v26_ = nw317_
                d_2066_v27_: D7
                d_2066_v27_ = D7_DC20(not((self).f32), not(p1))
                index326_ = default__.safeIndex(618, (d_2065_v26_).length(0))
                (d_2065_v26_)[index326_] = d_2066_v27_
                index327_ = default__.safeIndex(618, (d_2065_v26_).length(0))
                (d_2065_v26_)[index327_] = d_2066_v27_
        elif True:
            d_2067_v28_: _dafny.Seq
            d_2067_v28_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2068_v29_: _dafny.Seq
            d_2068_v29_ = _dafny.SeqWithoutIsStrInference([d_2067_v28_, d_2067_v28_])
            (globalState).f17 = len((d_2068_v29_) + (d_2068_v29_))
            if p1:
                d_2069_v30_: D7
                d_2069_v30_ = D7_DC20(p1, p1)
                d_2070_v31_: _dafny.MultiSet
                d_2070_v31_ = _dafny.MultiSet([d_2069_v30_, d_2069_v30_, D7_DC20(not(not(False)), p1)])
                d_2071_v32_: _dafny.Map
                d_2071_v32_ = _dafny.Map({((self).f31) | ((self).f31): _dafny.MultiSet([d_2069_v30_, d_2069_v30_, d_2069_v30_, d_2069_v30_])})
                d_2070_v31_ = ((d_2071_v32_)[(self).f31] if ((self).f31) in (d_2071_v32_) else _dafny.MultiSet([d_2069_v30_, d_2069_v30_, d_2069_v30_, d_2069_v30_]))
                d_2072_v33_: str
                d_2072_v33_ = _dafny.CodePoint('n')
                d_2073_v34_: _dafny.Map
                d_2073_v34_ = _dafny.Map({p1: p1})
                nw318_ = _dafny.Array(None, 9)
                nw318_[int(0)] = p1
                nw318_[int(1)] = p1
                nw318_[int(2)] = p1
                nw318_[int(3)] = (self).fm8(default__.fm33(d_2072_v33_, (self).f32, (self).f32, globalState), (self).f32, (self).f26, (self).f26, globalState)
                nw318_[int(4)] = not((self).f32)
                nw318_[int(5)] = (self).f32
                nw318_[int(6)] = p1
                nw318_[int(7)] = ((d_2073_v34_)[p1] if (p1) in (d_2073_v34_) else p1)
                nw318_[int(8)] = not(p1)
                (globalState).f1 = nw318_
                d_2074_v35_: _dafny.Map
                d_2074_v35_ = _dafny.Map({(self).f26: (self).f32})
                d_2075_v36_: _dafny.Map
                d_2075_v36_ = _dafny.Map({(self).f26: d_2072_v33_})
                d_2076_v37_: _dafny.Map
                d_2076_v37_ = _dafny.Map({len(d_2075_v36_): d_2072_v33_})
                d_2077_v38_: _dafny.Map
                d_2077_v38_ = _dafny.Map({(d_2073_v34_).set((self).f32, p1): (_dafny.Map({d_2067_v28_: p1})).set(d_2067_v28_, p1)})
                d_2078_v39_: _dafny.Map
                d_2078_v39_ = _dafny.Map({d_2067_v28_: p1})
                d_2079_v40_: _dafny.Set
                d_2079_v40_ = _dafny.Set({(self).fm6((d_2074_v35_).set((self).f26, not((self).f32)), (self).f26, _dafny.SeqWithoutIsStrInference([len(d_2076_v37_), (self).f26, (self).f26]), ((d_2077_v38_)[d_2073_v34_] if (d_2073_v34_) in (d_2077_v38_) else d_2078_v39_), globalState)})
                (globalState).f0 = len(d_2079_v40_)
                d_2080_v41_: _dafny.Array
                def lambda105_(d_2081_i4_):
                    return (self).f32

                init58_ = lambda105_
                nw319_ = _dafny.Array(None, 14)
                for i0_58_ in range(nw319_.length(0)):
                    nw319_[i0_58_] = init58_(i0_58_)
                d_2080_v41_ = nw319_
                d_2082_v42_: _dafny.Set
                d_2082_v42_ = _dafny.Set({d_2080_v41_, d_2080_v41_, d_2080_v41_})
                (globalState).f9 = (d_2082_v42_).issubset(d_2082_v42_)
                (globalState).f9 = p1
            elif True:
                d_2083_v44_: _dafny.Set
                d_2083_v44_ = _dafny.Set({(self).f26})
                d_2084_v45_: _dafny.Map
                def iife170_():
                    coll104_ = _dafny.Set()
                    compr_104_: int
                    for compr_104_ in _dafny.IntegerRange(-369, 333):
                        d_2085_v43_: int = compr_104_
                        if ((-369) <= (d_2085_v43_)) and ((d_2085_v43_) < (333)):
                            coll104_ = coll104_.union(_dafny.Set([(d_2085_v43_) + ((self).f26)]))
                    return _dafny.Set(coll104_)
                d_2084_v45_ = _dafny.Map({(iife170_()
                ) | (d_2083_v44_): 249})
                d_2086_v46_: _dafny.MultiSet
                d_2086_v46_ = _dafny.MultiSet([(self).f26, (0) - ((self).f26)])
                d_2087_v49_: _dafny.Seq
                d_2087_v49_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])
                d_2088_v50_: _dafny.MultiSet
                d_2088_v50_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f26, len(d_2067_v28_)]), _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (self).f26, (self).f26]), d_2087_v49_])
                def iife171_():
                    coll105_ = _dafny.Set()
                    compr_105_: int
                    for compr_105_ in (d_2086_v46_).Elements:
                        d_2089_v47_: int = compr_105_
                        if (d_2089_v47_) in (d_2086_v46_):
                            def iife172_():
                                coll106_ = _dafny.Map()
                                compr_106_: int
                                for compr_106_ in (_dafny.SeqWithoutIsStrInference([976 for d_2090_i5_ in range(default__.abs(-300))])).Elements:
                                    d_2091_v48_: int = compr_106_
                                    if (d_2091_v48_) in (_dafny.SeqWithoutIsStrInference([976 for d_2090_i5_ in range(default__.abs(-300))])):
                                        coll106_[(d_2091_v48_) - (len(_dafny.SeqWithoutIsStrInference([933, 419])))] = (_dafny.MultiSet([len(_dafny.Map({True: -378})), -317])).cardinality
                                return _dafny.Map(coll106_)
                            coll105_ = coll105_.union(_dafny.Set([(d_2089_v47_) * (len(iife172_()
))]))
                    return _dafny.Set(coll105_)
                d_2084_v45_ = (d_2084_v45_).set((iife171_()
                ).intersection(d_2083_v44_), ((d_2088_v50_)[d_2087_v49_] if (d_2087_v49_) in (d_2088_v50_) else (self).f26))
                (globalState).f9 = p1
                (globalState).f17 = (0) - ((0) - (((self).f26) * (default__.safeModuloInt((self).f26, (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f32])))))))
                d_2092_v51_: _dafny.Map
                d_2092_v51_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbghxyl")): ((self).f26) <= (65)})
                d_2093_v53_: _dafny.Set
                d_2093_v53_ = _dafny.Set({(self).f27})
                d_2094_v54_: _dafny.Map
                def iife173_():
                    coll107_ = _dafny.Map()
                    compr_107_: _dafny.Seq
                    for compr_107_ in (d_2093_v53_).Elements:
                        d_2095_v52_: _dafny.Seq = compr_107_
                        if (d_2095_v52_) in (d_2093_v53_):
                            coll107_[d_2095_v52_] = p1
                    return _dafny.Map(coll107_)
                d_2094_v54_ = iife173_()
                
                d_2096_v56_: _dafny.Seq
                d_2096_v56_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])
                def iife174_():
                    coll108_ = _dafny.Map()
                    compr_108_: _dafny.Seq
                    for compr_108_ in (d_2096_v56_).Elements:
                        d_2097_v55_: _dafny.Seq = compr_108_
                        if (d_2097_v55_) in (d_2096_v56_):
                            coll108_[d_2097_v55_] = p1
                    return _dafny.Map(coll108_)
                d_2092_v51_ = ((d_2092_v51_) | ((d_2094_v54_))) | (iife174_()
                )
                d_2098_v57_: _dafny.Set
                d_2098_v57_ = _dafny.Set({((self).f26) <= ((self).f26), False, (self).f32, (self).f32, not(((self).f26) == ((self).f26))})
                d_2098_v57_ = (d_2098_v57_ if not(p1) else (d_2098_v57_).intersection(d_2098_v57_))
            (globalState).f5 = (self).f26
            d_2099_v58_: _dafny.MultiSet
            d_2099_v58_ = _dafny.MultiSet([p1, p1])
            d_2100_v59_: _dafny.Array
            nw320_ = _dafny.Array(None, 25)
            nw320_[int(0)] = (self).f32
            nw320_[int(1)] = p1
            nw320_[int(2)] = p1
            nw320_[int(3)] = not(not(p1))
            nw320_[int(4)] = p1
            nw320_[int(5)] = p1
            nw320_[int(6)] = (self).f32
            nw320_[int(7)] = p1
            nw320_[int(8)] = (self).f32
            nw320_[int(9)] = (self).f32
            nw320_[int(10)] = p1
            nw320_[int(11)] = p1
            nw320_[int(12)] = ((self).f27) < ((self).f27)
            nw320_[int(13)] = not(False)
            nw320_[int(14)] = (self).f32
            nw320_[int(15)] = (self).f32
            nw320_[int(16)] = True
            nw320_[int(17)] = (self).f32
            nw320_[int(18)] = (self).f32
            nw320_[int(19)] = ((self).f26) == ((self).f26)
            nw320_[int(20)] = (self).f32
            nw320_[int(21)] = p1
            nw320_[int(22)] = (d_2099_v58_).issubset(d_2099_v58_)
            nw320_[int(23)] = not(p1)
            nw320_[int(24)] = ((self).f32) == ((self).f32)
            d_2100_v59_ = nw320_
            index328_ = default__.safeIndex(477, (d_2100_v59_).length(0))
            (d_2100_v59_)[index328_] = p1
            index329_ = default__.safeIndex(477, (d_2100_v59_).length(0))
            (d_2100_v59_)[index329_] = p1
            d_2101_v60_: _dafny.Seq
            d_2101_v60_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_2102_v61_: _dafny.MultiSet
            d_2102_v61_ = _dafny.MultiSet([(0) - ((self).f26), (self).f26, len(d_2101_v60_)])
            d_2103_v62_: _dafny.Map
            d_2103_v62_ = _dafny.Map({d_2102_v61_: d_2067_v28_})
            d_2103_v62_ = d_2103_v62_
        d_2104_v63_: _dafny.Seq
        d_2104_v63_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        d_2105_v64_: C2
        nw321_ = C2()
        nw321_.ctor__(d_2104_v63_, (self).f31)
        d_2105_v64_ = nw321_
        d_2106_v65_: _dafny.Seq
        d_2106_v65_ = _dafny.SeqWithoutIsStrInference([False])
        d_2107_v66_: _dafny.Map
        d_2107_v66_ = _dafny.Map({d_2106_v65_: p1})
        hi9_ = default__.safeDivisionInt((self).f26, (self).f26)
        for d_2108_i6_ in range((self).fm6(_dafny.Map({(self).f26: p1}), 855, (d_2105_v64_).f33, d_2107_v66_, globalState), hi9_):
            (globalState).f25 = p1
            d_2109_v67_: _dafny.MultiSet
            d_2109_v67_ = _dafny.MultiSet([(self).f32])
            d_2110_v68_: D15
            d_2110_v68_ = D15_DC42(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1])))
            d_2111_v69_: _dafny.Map
            d_2111_v69_ = _dafny.Map({(0) - ((0) - (d_2108_i6_)): d_2109_v67_})
            d_2112_v70_: _dafny.Map
            d_2112_v70_ = _dafny.Map({(self).f32: d_2109_v67_})
            d_2113_v71_: _dafny.Seq
            d_2113_v71_ = _dafny.SeqWithoutIsStrInference([d_2109_v67_])
            d_2114_v72_: _dafny.Array
            nw322_ = _dafny.Array(None, 20)
            nw322_[int(0)] = d_2109_v67_
            nw322_[int(1)] = (d_2109_v67_).intersection(d_2109_v67_)
            nw322_[int(2)] = (d_2110_v68_).cf80
            nw322_[int(3)] = d_2109_v67_
            nw322_[int(4)] = _dafny.MultiSet((d_2106_v65_) + (d_2106_v65_))
            nw322_[int(5)] = (d_2109_v67_) - (_dafny.MultiSet([p1]))
            nw322_[int(6)] = (d_2109_v67_).intersection(((d_2111_v69_)[(self).f26] if ((self).f26) in (d_2111_v69_) else d_2109_v67_))
            nw322_[int(7)] = (_dafny.MultiSet([p1, p1, (self).f32, (self).f32])) - (d_2109_v67_)
            nw322_[int(8)] = ((d_2112_v70_)[p1] if (p1) in (d_2112_v70_) else d_2109_v67_)
            nw322_[int(9)] = d_2109_v67_
            nw322_[int(10)] = d_2109_v67_
            nw322_[int(11)] = (d_2113_v71_)[default__.safeIndex((self).f26, len(d_2113_v71_))]
            nw322_[int(12)] = d_2109_v67_
            nw322_[int(13)] = d_2109_v67_
            nw322_[int(14)] = (d_2109_v67_) | (d_2109_v67_)
            nw322_[int(15)] = (d_2109_v67_) - (d_2109_v67_)
            nw322_[int(16)] = _dafny.MultiSet(d_2106_v65_)
            nw322_[int(17)] = d_2109_v67_
            nw322_[int(18)] = d_2109_v67_
            nw322_[int(19)] = _dafny.MultiSet([p1, p1, (self).f32])
            d_2114_v72_ = nw322_
            d_2114_v72_ = d_2114_v72_
            d_2115_v73_: _dafny.Set
            d_2115_v73_ = _dafny.Set({True})
            d_2116_v74_: _dafny.Seq
            d_2116_v74_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, (self).f27])
            d_2117_v75_: _dafny.Seq
            out51_: _dafny.Seq
            out51_ = (d_2105_v64_).m3(default__.fm24((self).f32, not((self).f32), p1, len(d_2115_v73_), globalState), d_2116_v74_, globalState)
            d_2117_v75_ = out51_
            d_2115_v73_ = (d_2115_v73_).intersection(d_2115_v73_)
        d_2118_v76_: _dafny.Map
        d_2118_v76_ = _dafny.Map({p1: (self).f26})
        d_2119_v77_: _dafny.Map
        d_2119_v77_ = _dafny.Map({default__.safeModuloInt(((d_2105_v64_).f33)[default__.safeIndex((0) - ((self).f26), len((d_2105_v64_).f33))], 14): d_2118_v76_})
        d_2119_v77_ = d_2119_v77_
        d_2120_v78_: C0
        nw323_ = C0()
        nw323_.ctor__()
        d_2120_v78_ = nw323_
        hi10_ = (self).f26
        for d_2121_i7_ in range((self).f26, hi10_):
            (globalState).f0 = default__.safeDivisionInt(default__.safeModuloInt((self).f26, d_2121_i7_), (self).f26)
            d_2122_v79_: D0
            d_2122_v79_ = D0_DC2(not((self).f32))
            d_2123_v80_: int
            d_2124_v81_: _dafny.MultiSet
            d_2125_v82_: _dafny.Array
            out52_: int
            out53_: _dafny.MultiSet
            out54_: _dafny.Array
            out52_, out53_, out54_ = (d_2105_v64_).m2(p1, d_2122_v79_, globalState)
            d_2123_v80_ = out52_
            d_2124_v81_ = out53_
            d_2125_v82_ = out54_
            rhs353_ = p1
            rhs354_ = (self).f26
            lhs312_ = globalState
            lhs313_ = globalState
            lhs312_.f25 = rhs353_
            lhs313_.f5 = rhs354_
            d_2126_v83_: _dafny.Array
            nw324_ = _dafny.Array(False, 14)
            d_2126_v83_ = nw324_
            (globalState).f1 = (d_2126_v83_ if not((self).f32) else d_2126_v83_)
        r0 = (self).f26
        r1 = False
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Array = _dafny.Array(None, 0)
        if (self).f32:
            d_2127_v0_: _dafny.Seq
            d_2127_v0_ = _dafny.SeqWithoutIsStrInference([(self).f26])
            (globalState).f9 = ((self).f26) not in ((d_2127_v0_ if p0 else _dafny.SeqWithoutIsStrInference([(self).f26])))
            d_2128_v1_: _dafny.Seq
            d_2128_v1_ = _dafny.SeqWithoutIsStrInference([not((self).f32), True])
            (globalState).f25 = (d_2128_v1_) <= (d_2128_v1_)
            d_2129_v2_: _dafny.Array
            def lambda106_(d_2130_i0_):
                return D3_DC11((self).f31, (self).f32, (self).f26, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2131_i1_ in range(default__.abs(252))])))

            init59_ = lambda106_
            nw325_ = _dafny.Array(None, 6)
            for i0_59_ in range(nw325_.length(0)):
                nw325_[i0_59_] = init59_(i0_59_)
            d_2129_v2_ = nw325_
            d_2132_v3_: _dafny.Map
            d_2132_v3_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndhq")): d_2129_v2_})
            d_2133_v4_: str
            d_2133_v4_ = _dafny.CodePoint('p')
            d_2132_v3_ = (d_2132_v3_).set((((self).f27) + ((self).f27)).set(default__.safeIndex((self).f26, len(((self).f27) + ((self).f27))), d_2133_v4_), d_2129_v2_)
            (globalState).f5 = ((self).f26) * ((self).f26)
            d_2134_v5_: _dafny.MultiSet
            d_2134_v5_ = _dafny.MultiSet([(self).f32, p0])
            d_2135_v6_: T1
            nw326_ = C3()
            nw326_.ctor__(532, (d_2134_v5_).cardinality, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ui")))
            d_2135_v6_ = nw326_
            d_2136_v7_: _dafny.Set
            d_2136_v7_ = _dafny.Set({p0})
            d_2137_v8_: _dafny.Map
            d_2137_v8_ = _dafny.Map({len(_dafny.Map({d_2135_v6_: (self).f32})): d_2136_v7_})
            d_2138_v9_: _dafny.Map
            d_2138_v9_ = _dafny.Map({p0: len(((d_2137_v8_)[-576] if (-576) in (d_2137_v8_) else d_2136_v7_))})
            d_2139_v10_: _dafny.Map
            d_2139_v10_ = _dafny.Map({d_2138_v9_: (self).f27})
            d_2139_v10_ = (d_2139_v10_).set(d_2138_v9_, (self).f27)
        elif True:
            d_2140_v11_: C1
            nw327_ = C1()
            nw327_.ctor__()
            d_2140_v11_ = nw327_
            d_2141_v12_: _dafny.Array
            nw328_ = _dafny.Array(False, 28)
            d_2141_v12_ = nw328_
            (globalState).f1 = d_2141_v12_
            d_2142_v13_: C2
            nw329_ = C2()
            nw329_.ctor__(_dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (self).f26]), ((self).f31) | ((self).f31))
            d_2142_v13_ = nw329_
            (globalState).f9 = (self).f32
            d_2143_v14_: _dafny.Seq
            d_2143_v14_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2144_v15_: _dafny.Map
            d_2144_v15_ = _dafny.Map({(d_2143_v14_) + (d_2143_v14_): p0})
            d_2144_v15_ = (d_2144_v15_).set(d_2143_v14_, (self).f32)
        d_2145_v16_: _dafny.Array
        nw330_ = _dafny.Array(None, 5)
        nw330_[int(0)] = (self).f26
        nw330_[int(1)] = (self).f26
        nw330_[int(2)] = (self).f26
        nw330_[int(3)] = (self).f26
        nw330_[int(4)] = 855
        d_2145_v16_ = nw330_
        d_2146_v17_: _dafny.Array
        nw331_ = _dafny.Array(None, 3)
        nw331_[int(0)] = d_2145_v16_
        nw331_[int(1)] = d_2145_v16_
        nw331_[int(2)] = d_2145_v16_
        d_2146_v17_ = nw331_
        index330_ = default__.safeIndex(584, (d_2146_v17_).length(0))
        (d_2146_v17_)[index330_] = d_2145_v16_
        index331_ = default__.safeIndex(584, (d_2146_v17_).length(0))
        rhs355_ = d_2145_v16_
        rhs356_ = (self).fm7((self).f26, (self).f32, globalState)
        rhs357_ = (self).f26
        lhs314_ = d_2146_v17_
        lhs315_ = default__.safeIndex(584, (d_2146_v17_).length(0))
        lhs316_ = globalState
        lhs314_[lhs315_] = rhs355_
        lhs316_.f24 = rhs356_
        r0 = rhs357_
        d_2147_v18_: _dafny.Array
        nw332_ = _dafny.Array(False, 26)
        d_2147_v18_ = nw332_
        index332_ = default__.safeIndex(454, (d_2147_v18_).length(0))
        (d_2147_v18_)[index332_] = p0
        d_2148_v19_: _dafny.Seq
        d_2148_v19_ = _dafny.SeqWithoutIsStrInference([(self).f31, (self).f31])
        d_2149_v20_: str
        d_2149_v20_ = _dafny.CodePoint('q')
        d_2150_v21_: _dafny.Map
        d_2150_v21_ = _dafny.Map({(self).f32: (self).f26})
        index333_ = default__.safeIndex(454, (d_2147_v18_).length(0))
        rhs358_ = d_2147_v18_
        rhs359_ = ((self).f26) + (((self).f26) - (len(((self).f27).set(default__.safeIndex(len(d_2148_v19_), len((self).f27)), d_2149_v20_))))
        rhs360_ = p0
        rhs361_ = (self).f26
        rhs362_ = ((self).f26) >= (len((default__.fm47(globalState)) | (d_2150_v21_)))
        lhs317_ = globalState
        lhs318_ = globalState
        lhs319_ = globalState
        lhs320_ = globalState
        lhs321_ = d_2147_v18_
        lhs322_ = default__.safeIndex(454, (d_2147_v18_).length(0))
        lhs317_.f1 = rhs358_
        lhs318_.f0 = rhs359_
        lhs319_.f25 = rhs360_
        lhs320_.f0 = rhs361_
        lhs321_[lhs322_] = rhs362_
        (globalState).f9 = not(p0)
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_2147_v18_).length(0)):
            d_2151_i2_: int = guard_loop_2_
            if (True) and (((0) <= (d_2151_i2_)) and ((d_2151_i2_) < ((d_2147_v18_).length(0)))):
                (d_2147_v18_)[(d_2151_i2_)] = (p0) or (not(p0))
        d_2152_v22_: _dafny.Map
        d_2152_v22_ = _dafny.Map({(d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))]: (self).f32})
        d_2153_v23_: D3
        d_2153_v23_ = D3_DC10(d_2152_v22_)
        d_2154_v24_: D12
        d_2154_v24_ = D12_DC37((self).f26, ((self).f26) + ((self).f26), d_2153_v23_, (self).f26, (self).f26)
        source31_ = d_2154_v24_
        if source31_.is_DC37:
            d_2155___mcc_h0_ = source31_.cf70
            d_2156___mcc_h1_ = source31_.cf71
            d_2157___mcc_h2_ = source31_.cf72
            d_2158___mcc_h3_ = source31_.cf73
            d_2159___mcc_h4_ = source31_.cf74
            d_2160_cf74_ = d_2159___mcc_h4_
            d_2161_cf73_ = d_2158___mcc_h3_
            d_2162_cf72_ = d_2157___mcc_h2_
            d_2163_cf71_ = d_2156___mcc_h1_
            d_2164_cf70_ = d_2155___mcc_h0_
            (globalState).f9 = p0
            d_2165_v25_: _dafny.Map
            d_2165_v25_ = _dafny.Map({d_2163_cf71_: default__.safeModuloInt(d_2163_cf71_, d_2163_cf71_)})
            d_2166_v26_: _dafny.MultiSet
            d_2166_v26_ = _dafny.MultiSet([584, d_2164_cf70_, d_2161_cf73_, len((self).f27)])
            d_2167_v27_: _dafny.MultiSet
            d_2167_v27_ = _dafny.MultiSet([(self).f32])
            d_2168_v28_: _dafny.Seq
            d_2168_v28_ = _dafny.SeqWithoutIsStrInference([d_2161_cf73_])
            index334_ = default__.safeIndex(454, (d_2147_v18_).length(0))
            rhs363_ = (self).fm8(d_2166_v26_, not((self).f32), ((self).f26 if (self).f32 else (d_2167_v27_).cardinality), d_2160_cf74_, globalState)
            rhs364_ = (d_2168_v28_) != (_dafny.SeqWithoutIsStrInference([(self).f26, -150]))
            rhs365_ = d_2164_cf70_
            rhs366_ = (self).f31
            lhs323_ = d_2147_v18_
            lhs324_ = default__.safeIndex(454, (d_2147_v18_).length(0))
            lhs325_ = globalState
            lhs326_ = globalState
            lhs323_[lhs324_] = rhs363_
            lhs325_.f9 = rhs364_
            lhs326_.f8 = rhs365_
            d_2165_v25_ = rhs366_
            index335_ = default__.safeIndex(290, (d_2145_v16_).length(0))
            (d_2145_v16_)[index335_] = d_2163_cf71_
            d_2169_v29_: _dafny.MultiSet
            d_2169_v29_ = _dafny.MultiSet([d_2147_v18_, d_2147_v18_])
            index336_ = default__.safeIndex(290, (d_2145_v16_).length(0))
            rhs367_ = (866) - ((0) - (d_2161_cf73_))
            rhs368_ = (not((d_2167_v27_) != (default__.fm37((d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))], p0, (default__.fm31((d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))], (d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))], True, p0, globalState)).cf36, (d_2166_v26_).cardinality, globalState))) if (self).f32 else ((self).f27) < ((D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aewnvt")), d_2161_cf73_, len((self).f31))).cf1))
            rhs369_ = ((d_2150_v21_).set(p0, d_2163_cf71_)).set((d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))], (self).f26)
            rhs370_ = d_2149_v20_
            rhs371_ = default__.safeModuloInt((((d_2169_v29_).set(d_2147_v18_, default__.abs(d_2161_cf73_))) - (d_2169_v29_)).cardinality, (d_2163_cf71_) + (len((self).f27)))
            lhs327_ = d_2145_v16_
            lhs328_ = default__.safeIndex(290, (d_2145_v16_).length(0))
            lhs329_ = globalState
            lhs330_ = globalState
            lhs327_[lhs328_] = rhs367_
            lhs329_.f25 = rhs368_
            d_2150_v21_ = rhs369_
            d_2149_v20_ = rhs370_
            lhs330_.f17 = rhs371_
            d_2170_v30_: _dafny.Map
            d_2170_v30_ = _dafny.Map({(self).f26: (d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))]})
            d_2171_v31_: C2
            nw333_ = C2()
            nw333_.ctor__((d_2168_v28_).set(default__.safeIndex(d_2160_cf74_, len(d_2168_v28_)), d_2160_cf74_), default__.fm40((d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))], d_2164_cf70_, d_2170_v30_, (self).f32, globalState))
            d_2171_v31_ = nw333_
        elif source31_.is_DC38:
            arr2_ = (d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]
            index337_ = default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))
            arr2_[index337_] = (0) - ((self).f26)
            d_2172_v32_: _dafny.Set
            d_2172_v32_ = _dafny.Set({(self).f32})
            d_2173_v33_: D1
            d_2173_v33_ = D1_DC5(d_2172_v32_)
            d_2174_v34_: _dafny.Map
            d_2174_v34_ = _dafny.Map({len((self).f27): d_2173_v33_})
            d_2175_v35_: _dafny.Map
            d_2175_v35_ = d_2174_v34_
            d_2176_v36_: D8
            d_2176_v36_ = D8_DC25((d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))], d_2150_v21_, d_2149_v20_, d_2175_v35_, (self).f26)
            d_2177_v37_: _dafny.Map
            d_2177_v37_ = _dafny.Map({(d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))]: (d_2176_v36_).cf51})
            d_2178_v38_: _dafny.Seq
            d_2178_v38_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2179_v39_: _dafny.Seq
            d_2179_v39_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxrlhps"))), len(d_2178_v38_), (self).f26, (self).f26])
            d_2180_v40_: _dafny.Map
            d_2180_v40_ = _dafny.Map({(d_2178_v38_).set(default__.safeIndex(len(d_2178_v38_), len(d_2178_v38_)), True): (self).f32})
            index338_ = default__.safeIndex(454, (d_2147_v18_).length(0))
            arr3_ = (d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]
            index339_ = default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))
            rhs372_ = True
            rhs373_ = (self).fm6(default__.fm18((self).f26, globalState), (0) - ((self).f26), d_2179_v39_, (d_2180_v40_) | ((d_2180_v40_).set(d_2178_v38_, False)), globalState)
            rhs374_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exfichunx"))) < ((self).f27)
            rhs375_ = default__.fm23((self).f27, globalState)
            lhs331_ = d_2147_v18_
            lhs332_ = default__.safeIndex(454, (d_2147_v18_).length(0))
            lhs333_ = (d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]
            lhs334_ = default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))
            lhs335_ = globalState
            lhs331_[lhs332_] = rhs372_
            lhs333_[lhs334_] = rhs373_
            lhs335_.f25 = rhs374_
            d_2177_v37_ = rhs375_
            d_2181_v41_: _dafny.Array
            nw334_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_2181_v41_ = nw334_
            d_2182_v42_: _dafny.Array
            nw335_ = _dafny.Array(_dafny.Set({}), 12)
            d_2182_v42_ = nw335_
            index340_ = default__.safeIndex(713, (d_2181_v41_).length(0))
            (d_2181_v41_)[index340_] = d_2182_v42_
            index341_ = default__.safeIndex(713, (d_2181_v41_).length(0))
            nw336_ = _dafny.Array(_dafny.Set({}), 6)
            (d_2181_v41_)[index341_] = nw336_
            if False:
                arr4_ = (d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]
                index342_ = default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))
                arr4_[index342_] = (self).f26
                d_2183_v43_: _dafny.Array
                nw337_ = _dafny.Array(_dafny.Map({}), 12)
                d_2183_v43_ = nw337_
                index343_ = default__.safeIndex(319, (d_2183_v43_).length(0))
                (d_2183_v43_)[index343_] = d_2174_v34_
                d_2184_v44_: _dafny.Array
                nw338_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_2184_v44_ = nw338_
                d_2185_v45_: D13
                d_2185_v45_ = D13_DC40(d_2184_v44_, d_2147_v18_, (d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))])
                d_2186_v46_: _dafny.Seq
                d_2186_v46_ = _dafny.SeqWithoutIsStrInference([d_2147_v18_, d_2147_v18_, d_2147_v18_])
                d_2187_v47_: _dafny.Map
                d_2187_v47_ = _dafny.Map({(d_2186_v46_)[default__.safeIndex(((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))], len(d_2186_v46_))]: (self).f32})
                d_2188_v48_: _dafny.Array
                nw339_ = _dafny.Array(None, 4)
                nw339_[int(0)] = (d_2187_v47_ if (d_2185_v45_).cf78 else d_2187_v47_)
                nw339_[int(1)] = d_2187_v47_
                nw339_[int(2)] = (d_2187_v47_) | (d_2187_v47_)
                nw339_[int(3)] = d_2187_v47_
                d_2188_v48_ = nw339_
                index344_ = default__.safeIndex(634, (d_2188_v48_).length(0))
                (d_2188_v48_)[index344_] = d_2187_v47_
                d_2189_v49_: _dafny.Array
                nw340_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_2189_v49_ = nw340_
                index345_ = default__.safeIndex(424, (d_2189_v49_).length(0))
                (d_2189_v49_)[index345_] = d_2145_v16_
                d_2190_v50_: D18
                d_2190_v50_ = D18_DC50((self).f26, (d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))], d_2179_v39_, d_2149_v20_, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))])
                d_2191_v51_: D18
                d_2191_v51_ = D18_DC51(d_2190_v50_)
                d_2192_v52_: D18
                d_2192_v52_ = D18_DC51(d_2190_v50_)
                d_2193_v53_: D18
                d_2193_v53_ = D18_DC51(d_2190_v50_)
                d_2194_v54_: _dafny.Array
                nw341_ = _dafny.Array(None, 11)
                nw341_[int(0)] = d_2193_v53_
                nw341_[int(1)] = d_2193_v53_
                nw341_[int(2)] = d_2193_v53_
                nw341_[int(3)] = d_2193_v53_
                nw341_[int(4)] = d_2193_v53_
                nw341_[int(5)] = D18_DC51(d_2192_v52_)
                nw341_[int(6)] = d_2193_v53_
                nw341_[int(7)] = d_2193_v53_
                nw341_[int(8)] = d_2193_v53_
                nw341_[int(9)] = d_2193_v53_
                nw341_[int(10)] = D18_DC51(d_2190_v50_)
                d_2194_v54_ = nw341_
                d_2195_v55_: _dafny.Map
                d_2195_v55_ = _dafny.Map({d_2194_v54_: d_2145_v16_})
                index346_ = default__.safeIndex(319, (d_2183_v43_).length(0))
                index347_ = default__.safeIndex(634, (d_2188_v48_).length(0))
                index348_ = default__.safeIndex(424, (d_2189_v49_).length(0))
                rhs376_ = ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))]
                rhs377_ = d_2175_v35_
                rhs378_ = d_2187_v47_
                rhs379_ = ((d_2195_v55_)[d_2194_v54_] if (d_2194_v54_) in (d_2195_v55_) else d_2145_v16_)
                rhs380_ = (self).f26
                lhs336_ = globalState
                lhs337_ = d_2183_v43_
                lhs338_ = default__.safeIndex(319, (d_2183_v43_).length(0))
                lhs339_ = d_2188_v48_
                lhs340_ = default__.safeIndex(634, (d_2188_v48_).length(0))
                lhs341_ = d_2189_v49_
                lhs342_ = default__.safeIndex(424, (d_2189_v49_).length(0))
                lhs343_ = globalState
                lhs336_.f8 = rhs376_
                lhs337_[lhs338_] = rhs377_
                lhs339_[lhs340_] = rhs378_
                lhs341_[lhs342_] = rhs379_
                lhs343_.f8 = rhs380_
                d_2196_v56_: C2
                nw342_ = C2()
                nw342_.ctor__(((_dafny.SeqWithoutIsStrInference([((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))] for d_2197_i3_ in range(default__.abs(289))])) + (d_2179_v39_)).set(default__.safeIndex(65, len((_dafny.SeqWithoutIsStrInference([((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))] for d_2198_i3_ in range(default__.abs(289))])) + (d_2179_v39_))), (self).f26), (self).f31)
                d_2196_v56_ = nw342_
                nw343_ = C2()
                nw343_.ctor__(default__.fm45(((d_2196_v56_).f33)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))]])), len((d_2196_v56_).f33))], _dafny.Map({(d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))]: 428}), (self).f31, globalState), _dafny.Map({((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))]: (self).f26}))
                d_2196_v56_ = nw343_
                d_2199_v57_: _dafny.Map
                d_2199_v57_ = _dafny.Map({((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))]: d_2150_v21_})
                d_2200_v58_: D8
                d_2200_v58_ = D8_DC24(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukglkivlo"))), len(d_2199_v57_), ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))])
                d_2201_v59_: _dafny.Map
                d_2201_v59_ = _dafny.Map({d_2200_v58_: d_2179_v39_})
                d_2202_v60_: _dafny.Seq
                d_2202_v60_ = _dafny.SeqWithoutIsStrInference([d_2179_v39_])
                d_2203_v61_: C2
                nw344_ = C2()
                nw344_.ctor__(((d_2201_v59_)[d_2200_v58_] if (d_2200_v58_) in (d_2201_v59_) else (d_2202_v60_)[default__.safeIndex(len(d_2178_v38_), len(d_2202_v60_))]), (self).f31)
                d_2203_v61_ = nw344_
                d_2204_v62_: _dafny.Map
                d_2204_v62_ = _dafny.Map({(d_2172_v32_).intersection(d_2172_v32_): p0})
                d_2205_v63_: _dafny.Map
                d_2205_v63_ = _dafny.Map({(self).f26: (self).f32})
                def iife175_():
                    coll109_ = _dafny.Map()
                    compr_109_: int
                    for compr_109_ in (d_2179_v39_).Elements:
                        d_2206_v64_: int = compr_109_
                        if (d_2206_v64_) in (d_2179_v39_):
                            coll109_[(d_2206_v64_) - (((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))])] = (d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))]
                    return _dafny.Map(coll109_)
                d_2204_v62_ = (d_2204_v62_).set(d_2172_v32_, not(not(((d_2196_v56_).fm6(d_2205_v63_, (d_2196_v56_).fm6(iife175_()
                , ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))], d_2179_v39_, default__.fm28(globalState), globalState), _dafny.SeqWithoutIsStrInference([(0) - ((self).f26) for d_2207_i4_ in range(default__.abs(580))]), d_2180_v40_, globalState)) < ((self).f26))))
            elif True:
                d_2208_v65_: _dafny.Set
                d_2208_v65_ = _dafny.Set({999, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))]})
                d_2208_v65_ = d_2208_v65_
                arr5_ = (d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]
                index349_ = default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))
                arr5_[index349_] = ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))]
                index350_ = default__.safeIndex(454, (d_2147_v18_).length(0))
                (d_2147_v18_)[index350_] = (self).f32
                arr6_ = (d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]
                index351_ = default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))
                arr6_[index351_] = ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))]
                arr7_ = (d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]
                index352_ = default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))
                arr7_[index352_] = (self).fm7(((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))])[default__.safeIndex(438, ((d_2146_v17_)[default__.safeIndex(584, (d_2146_v17_).length(0))]).length(0))], p0, globalState)
            d_2209_v66_: _dafny.MultiSet
            d_2209_v66_ = _dafny.MultiSet([(self).f26])
            (globalState).f9 = (p0 if (d_2209_v66_).issubset(d_2209_v66_) else (self).f32)
        elif True:
            d_2210___mcc_h5_ = source31_.cf69
            d_2211_cf69_ = d_2210___mcc_h5_
            (globalState).f0 = (0) - ((0) - (default__.safeDivisionInt((self).f26, (self).f26)))
            d_2212_v67_: _dafny.Map
            d_2212_v67_ = _dafny.Map({(self).f26: (self).f32})
            d_2213_v68_: _dafny.Set
            d_2213_v68_ = _dafny.Set({((d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))] if ((d_2212_v67_)[(self).f26] if ((self).f26) in (d_2212_v67_) else False) else (self).f32)})
            d_2214_v69_: _dafny.Seq
            d_2214_v69_ = _dafny.SeqWithoutIsStrInference([d_2213_v68_, (d_2213_v68_) | (d_2213_v68_), (d_2213_v68_) | (d_2213_v68_), d_2213_v68_, d_2213_v68_])
            d_2215_v70_: _dafny.Map
            d_2215_v70_ = _dafny.Map({d_2213_v68_: (self).f26})
            d_2213_v68_ = (d_2214_v69_)[default__.safeIndex((0) - (len(d_2215_v70_)), len(d_2214_v69_))]
            index353_ = default__.safeIndex(454, (d_2147_v18_).length(0))
            (d_2147_v18_)[index353_] = (d_2147_v18_)[default__.safeIndex(454, (d_2147_v18_).length(0))]
            d_2216_v71_: _dafny.MultiSet
            d_2216_v71_ = _dafny.MultiSet([(self).f26])
            d_2217_v72_: _dafny.Seq
            d_2217_v72_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (self).f26])
            rhs381_ = d_2147_v18_
            rhs382_ = _dafny.MultiSet(d_2217_v72_)
            d_2147_v18_ = rhs381_
            d_2216_v71_ = rhs382_
        r0 = (self).f26
        d_2218_v73_: D0
        d_2218_v73_ = D0_DC0((self).f26)
        d_2219_v74_: _dafny.MultiSet
        d_2219_v74_ = _dafny.MultiSet([d_2218_v73_, d_2218_v73_, d_2218_v73_])
        r1 = (d_2219_v74_).intersection((_dafny.MultiSet([D0_DC0((self).f26)])).intersection(d_2219_v74_))
        r2 = d_2146_v17_
        return r0, r1, r2

    def m3(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2220_v0_: _dafny.Array
        nw345_ = _dafny.Array(False, 28)
        d_2220_v0_ = nw345_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_2220_v0_).length(0)):
            d_2221_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_2221_i0_)) and ((d_2221_i0_) < ((d_2220_v0_).length(0)))):
                (d_2220_v0_)[(d_2221_i0_)] = (self).f32
        (globalState).f17 = (self).f26
        d_2222_v1_: _dafny.Set
        d_2222_v1_ = _dafny.Set({p0, (self).f32})
        if ((d_2222_v1_).isdisjoint(d_2222_v1_)) == (p0):
            (globalState).f17 = (self).f26
            d_2223_v2_: C3
            nw346_ = C3()
            nw346_.ctor__(len(_dafny.Map({(self).f32: (self).f32})), (self).f26, (self).fm13(p0, (self).f27, (self).f26, globalState))
            d_2223_v2_ = nw346_
            d_2224_v3_: _dafny.Seq
            d_2224_v3_ = _dafny.SeqWithoutIsStrInference([((self).f32 if p0 else not((self).f32))])
            if (d_2224_v3_)[default__.safeIndex((self).f26, len(d_2224_v3_))]:
                d_2225_v4_: _dafny.Map
                d_2225_v4_ = _dafny.Map({p0: p0})
                d_2226_v5_: D3
                d_2226_v5_ = D3_DC10((d_2225_v4_).set(default__.fm24((self).f32, (self).f32, p0, (d_2223_v2_).f34, globalState), (self).f32))
                d_2227_v6_: _dafny.MultiSet
                d_2227_v6_ = _dafny.MultiSet([(d_2223_v2_).f34])
                d_2228_v7_: D12
                d_2228_v7_ = D12_DC37((self).f26, (self).f26, d_2226_v5_, ((d_2227_v6_)[(d_2223_v2_).f34] if ((d_2223_v2_).f34) in (d_2227_v6_) else 727), 412)
                d_2229_v8_: _dafny.Map
                d_2229_v8_ = _dafny.Map({(d_2223_v2_).f34: d_2228_v7_})
                d_2230_v9_: _dafny.MultiSet
                d_2230_v9_ = _dafny.MultiSet([d_2224_v3_])
                d_2231_v10_: _dafny.Map
                d_2231_v10_ = _dafny.Map({(0) - ((self).f26): p0})
                d_2229_v8_ = (d_2229_v8_).set(default__.safeDivisionInt(((d_2230_v9_)[d_2224_v3_] if (d_2224_v3_) in (d_2230_v9_) else (self).f26), (self).f26), D12_DC37((d_2223_v2_).f34, (d_2223_v2_).f34, d_2226_v5_, (d_2223_v2_).f34, len(d_2231_v10_)))
                d_2231_v10_ = (d_2231_v10_).set(((d_2223_v2_).f34) * ((d_2223_v2_).f34), p0)
                d_2232_v11_: _dafny.Array
                def lambda107_(d_2233_v3_):
                    def lambda108_(d_2234_i1_):
                        return default__.safeModuloInt(d_2234_i1_, len(d_2233_v3_))

                    return lambda108_

                init60_ = lambda107_(d_2224_v3_)
                nw347_ = _dafny.Array(None, 21)
                for i0_60_ in range(nw347_.length(0)):
                    nw347_[i0_60_] = init60_(i0_60_)
                d_2232_v11_ = nw347_
                d_2235_v12_: _dafny.Map
                d_2235_v12_ = _dafny.Map({d_2232_v11_: (self).f32})
                d_2235_v12_ = (d_2235_v12_).set(d_2232_v11_, True)
                (globalState).f1 = d_2220_v0_
                d_2236_v13_: _dafny.Array
                nw348_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_2236_v13_ = nw348_
                index354_ = default__.safeIndex(960, (d_2236_v13_).length(0))
                (d_2236_v13_)[index354_] = d_2220_v0_
                d_2237_v14_: _dafny.Map
                d_2237_v14_ = _dafny.Map({(self).f32: (d_2223_v2_).f34})
                d_2238_v15_: _dafny.Seq
                d_2238_v15_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_2223_v2_).f34: (self).f32})), len((d_2237_v14_) | ((_dafny.Map({p0: (d_2223_v2_).f34})).set((self).f32, len(d_2225_v4_)))), (d_2223_v2_).f34])
                d_2239_v16_: _dafny.Map
                d_2239_v16_ = _dafny.Map({(d_2223_v2_).f34: d_2220_v0_})
                d_2240_v17_: _dafny.MultiSet
                d_2240_v17_ = _dafny.MultiSet([False, p0])
                d_2241_v18_: D15
                d_2241_v18_ = D15_DC42(_dafny.MultiSet([p0, (self).f32]))
                index355_ = default__.safeIndex(960, (d_2236_v13_).length(0))
                rhs383_ = ((d_2239_v16_)[default__.safeDivisionInt((d_2223_v2_).f34, (d_2223_v2_).f34)] if (default__.safeDivisionInt((d_2223_v2_).f34, (d_2223_v2_).f34)) in (d_2239_v16_) else d_2220_v0_)
                rhs384_ = (d_2238_v15_) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f26, (d_2223_v2_).f34, (d_2223_v2_).f34, (d_2223_v2_).f34])), (d_2223_v2_).f34, (0) - ((d_2223_v2_).f34), (d_2223_v2_).f34, (d_2223_v2_).f34]))
                rhs385_ = (_dafny.Map({d_2224_v3_: 786})) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f32, (self).f32]): len(d_2222_v1_)}))
                rhs386_ = not (p0) or (not(((d_2241_v18_).cf80).issubset(d_2240_v17_)))
                lhs344_ = d_2236_v13_
                lhs345_ = default__.safeIndex(960, (d_2236_v13_).length(0))
                lhs346_ = globalState
                lhs347_ = globalState
                lhs344_[lhs345_] = rhs383_
                d_2238_v15_ = rhs384_
                lhs346_.f25 = rhs385_
                lhs347_.f25 = rhs386_
            elif True:
                d_2242_v19_: _dafny.Seq
                d_2242_v19_ = _dafny.SeqWithoutIsStrInference([(d_2223_v2_).f34])
                d_2243_v20_: D10
                d_2243_v20_ = D10_DC31(p0, (d_2242_v19_).set(default__.safeIndex((d_2223_v2_).f34, len(d_2242_v19_)), (d_2223_v2_).f34), (self).f32, (self).f26)
                d_2244_v21_: _dafny.Map
                d_2244_v21_ = _dafny.Map({_dafny.CodePoint('t'): (d_2223_v2_).f34})
                d_2245_v22_: _dafny.Array
                nw349_ = _dafny.Array(None, 8)
                nw349_[int(0)] = (d_2243_v20_).cf61
                nw349_[int(1)] = (d_2223_v2_).f34
                nw349_[int(2)] = len(d_2244_v21_)
                nw349_[int(3)] = (d_2223_v2_).f34
                nw349_[int(4)] = (d_2223_v2_).f34
                nw349_[int(5)] = (d_2242_v19_)[default__.safeIndex((d_2223_v2_).f34, len(d_2242_v19_))]
                nw349_[int(6)] = (d_2223_v2_).f34
                nw349_[int(7)] = (d_2223_v2_).f34
                d_2245_v22_ = nw349_
                nw350_ = _dafny.Array(int(0), 23)
                d_2245_v22_ = nw350_
                (globalState).f24 = (d_2223_v2_).f34
                index356_ = default__.safeIndex(152, (d_2220_v0_).length(0))
                (d_2220_v0_)[index356_] = p0
                index357_ = default__.safeIndex(152, (d_2220_v0_).length(0))
                (d_2220_v0_)[index357_] = not((d_2224_v3_) <= (d_2224_v3_))
                d_2246_v23_: C1
                nw351_ = C1()
                nw351_.ctor__()
                d_2246_v23_ = nw351_
                d_2247_v24_: _dafny.Map
                d_2247_v24_ = _dafny.Map({p0: ((d_2220_v0_)[default__.safeIndex(152, (d_2220_v0_).length(0))] if True else (self).f32)})
                d_2248_v25_: _dafny.Set
                d_2248_v25_ = _dafny.Set({len((self).f27), (d_2242_v19_)[default__.safeIndex((self).f26, len(d_2242_v19_))], (self).f26, (d_2242_v19_)[default__.safeIndex(344, len(d_2242_v19_))]})
                (globalState).f25 = ((d_2247_v24_)[True] if (True) in (d_2247_v24_) else (d_2248_v25_).ispropersubset(d_2248_v25_))
            (globalState).f24 = (0) - ((self).fm7(((self).f26) + ((self).f26), (self).f32, globalState))
            d_2249_v26_: _dafny.MultiSet
            d_2249_v26_ = _dafny.MultiSet([(d_2223_v2_).f34, (d_2223_v2_).f34])
            (globalState).f8 = default__.safeDivisionInt((0) - ((self).f26), len(_dafny.Map({(self).f32: (d_2223_v2_).fm8(d_2249_v26_, (self).f32, (self).f26, (d_2223_v2_).f34, globalState)})))
        elif True:
            d_2250_v27_: D12
            d_2250_v27_ = D12_DC38()
            d_2250_v27_ = d_2250_v27_
            d_2251_v28_: _dafny.Array
            def lambda109_(d_2252_p0_):
                def lambda110_(d_2253_i2_):
                    return (d_2253_i2_) * (len(_dafny.Map({len((self).f27): d_2252_p0_})))

                return lambda110_

            init61_ = lambda109_(p0)
            nw352_ = _dafny.Array(None, 4)
            for i0_61_ in range(nw352_.length(0)):
                nw352_[i0_61_] = init61_(i0_61_)
            d_2251_v28_ = nw352_
            index358_ = default__.safeIndex(401, (d_2251_v28_).length(0))
            (d_2251_v28_)[index358_] = (self).f26
            index359_ = default__.safeIndex(401, (d_2251_v28_).length(0))
            (d_2251_v28_)[index359_] = (self).f26
            if p0:
                index360_ = default__.safeIndex(150, (d_2220_v0_).length(0))
                (d_2220_v0_)[index360_] = (self).f32
                index361_ = default__.safeIndex(150, (d_2220_v0_).length(0))
                (d_2220_v0_)[index361_] = (self).f32
                d_2254_v29_: _dafny.Array
                nw353_ = _dafny.Array(_dafny.CodePoint('D'), 6)
                d_2254_v29_ = nw353_
                d_2255_v30_: _dafny.Array
                nw354_ = _dafny.Array(None, 5)
                nw354_[int(0)] = d_2254_v29_
                nw354_[int(1)] = d_2254_v29_
                nw354_[int(2)] = d_2254_v29_
                nw354_[int(3)] = d_2254_v29_
                nw354_[int(4)] = d_2254_v29_
                d_2255_v30_ = nw354_
                index362_ = default__.safeIndex(650, (d_2255_v30_).length(0))
                (d_2255_v30_)[index362_] = d_2254_v29_
                d_2256_v31_: str
                d_2256_v31_ = _dafny.CodePoint('w')
                d_2257_v32_: D20
                d_2257_v32_ = D20_DC54(d_2254_v29_)
                index363_ = default__.safeIndex(650, (d_2255_v30_).length(0))
                rhs387_ = ((d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))]) - ((self).fm7((0) - ((self).f26), p0, globalState))
                rhs388_ = (d_2220_v0_)[default__.safeIndex(150, (d_2220_v0_).length(0))]
                rhs389_ = ((d_2257_v32_).cf102 if (d_2220_v0_)[default__.safeIndex(150, (d_2220_v0_).length(0))] else d_2254_v29_)
                rhs390_ = ((d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))]) - ((d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))])
                rhs391_ = _dafny.CodePoint('y')
                lhs348_ = globalState
                lhs349_ = globalState
                lhs350_ = d_2255_v30_
                lhs351_ = default__.safeIndex(650, (d_2255_v30_).length(0))
                lhs352_ = globalState
                lhs348_.f17 = rhs387_
                lhs349_.f9 = rhs388_
                lhs350_[lhs351_] = rhs389_
                lhs352_.f17 = rhs390_
                d_2256_v31_ = rhs391_
                d_2258_v33_: _dafny.Map
                d_2258_v33_ = _dafny.Map({(d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))]: p0})
                d_2259_v34_: D0
                d_2259_v34_ = D0_DC0((self).f26)
                d_2260_v35_: _dafny.Map
                d_2260_v35_ = _dafny.Map({len(d_2258_v33_): d_2259_v34_})
                d_2261_v36_: D19
                d_2261_v36_ = D19_DC52(d_2260_v35_)
                d_2262_v37_: _dafny.Map
                d_2262_v37_ = _dafny.Map({(default__.fm48((self).f26, (self).f32, d_2261_v36_, globalState)).cf39: False})
                (globalState).f25 = (((d_2262_v37_)[(self).f32] if ((self).f32) in (d_2262_v37_) else not(True))) == (not ((d_2220_v0_)[default__.safeIndex(150, (d_2220_v0_).length(0))]) or (p0))
                d_2263_v38_: _dafny.MultiSet
                d_2263_v38_ = _dafny.MultiSet([(d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], (d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))]])
                d_2264_v39_: _dafny.Seq
                d_2264_v39_ = _dafny.SeqWithoutIsStrInference([(self).f32])
                d_2265_v40_: _dafny.Seq
                d_2265_v40_ = _dafny.SeqWithoutIsStrInference([(D1_DC6((self).f32, _dafny.CodePoint('m'), (_dafny.MultiSet([(d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))]])).cardinality)).cf11])
                d_2266_v41_: _dafny.Seq
                d_2266_v41_ = _dafny.SeqWithoutIsStrInference([(self).fm8(d_2263_v38_, (d_2220_v0_)[default__.safeIndex(150, (d_2220_v0_).length(0))], 905, len(d_2264_v39_), globalState), p0, default__.fm24((d_2264_v39_)[default__.safeIndex(len(d_2265_v40_), len(d_2264_v39_))], (d_2220_v0_)[default__.safeIndex(150, (d_2220_v0_).length(0))], (d_2220_v0_)[default__.safeIndex(150, (d_2220_v0_).length(0))], (self).f26, globalState), ((d_2258_v33_)[len(default__.fm21((d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], 111, (self).f32, (self).f32, globalState))] if (len(default__.fm21((d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], 111, (self).f32, (self).f32, globalState))) in (d_2258_v33_) else False)])
                d_2267_v42_: _dafny.MultiSet
                d_2267_v42_ = _dafny.MultiSet([(d_2220_v0_)[default__.safeIndex(150, (d_2220_v0_).length(0))]])
                (globalState).f17 = ((_dafny.MultiSet(d_2266_v41_)) - (d_2267_v42_)).cardinality
                (globalState).f13 = (self).f27
            elif True:
                d_2268_v43_: _dafny.Map
                d_2268_v43_ = _dafny.Map({p0: (self).f32})
                d_2269_v44_: _dafny.Seq
                d_2269_v44_ = _dafny.SeqWithoutIsStrInference([d_2268_v43_, d_2268_v43_, d_2268_v43_])
                d_2270_v45_: str
                d_2270_v45_ = _dafny.CodePoint('i')
                d_2271_v46_: D0
                d_2271_v46_ = D0_DC1(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2272_i5_ in range(default__.abs(862))]), len((self).f27), (self).f26)
                d_2273_v47_: _dafny.Array
                nw355_ = _dafny.Array(None, 27)
                nw355_[int(0)] = (self).f27
                nw355_[int(1)] = (self).f27
                nw355_[int(2)] = (self).f27
                nw355_[int(3)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqbwkhq"))) + ((self).f27)
                nw355_[int(4)] = (self).f27
                nw355_[int(5)] = (default__.fm43(p0, len((d_2269_v44_).set(default__.safeIndex((d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], len(d_2269_v44_)), d_2268_v43_)), globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxptv")))
                nw355_[int(6)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2274_i3_ in range(default__.abs(871))])
                nw355_[int(7)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnoknyvk"))) + ((self).f27)).set(default__.safeIndex((d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnoknyvk"))) + ((self).f27))), d_2270_v45_)
                nw355_[int(8)] = (default__.fm36(p0, p0, p0, globalState)) + (_dafny.SeqWithoutIsStrInference([d_2270_v45_ for d_2275_i4_ in range(default__.abs(588))]))
                nw355_[int(9)] = (self).f27
                nw355_[int(10)] = ((self).f27) + ((d_2271_v46_).cf1)
                nw355_[int(11)] = (self).f27
                nw355_[int(12)] = (self).f27
                nw355_[int(13)] = (self).f27
                nw355_[int(14)] = (self).f27
                nw355_[int(15)] = ((self).f27) + ((self).f27)
                nw355_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hncieu"))
                nw355_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
                nw355_[int(18)] = ((self).f27) + ((self).f27)
                nw355_[int(19)] = (self).f27
                nw355_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhaagcqq"))
                nw355_[int(21)] = (self).f27
                nw355_[int(22)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "paglhji"))
                nw355_[int(23)] = (self).f27
                nw355_[int(24)] = (self).f27
                nw355_[int(25)] = ((self).f27) + ((self).f27)
                nw355_[int(26)] = ((self).f27 if p0 else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxpqa"))).set(default__.safeIndex((self).f26, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxpqa")))), d_2270_v45_))
                d_2273_v47_ = nw355_
                index364_ = default__.safeIndex(118, (d_2273_v47_).length(0))
                (d_2273_v47_)[index364_] = (self).f27
                index365_ = default__.safeIndex(118, (d_2273_v47_).length(0))
                (d_2273_v47_)[index365_] = (self).f27
                d_2276_v48_: C1
                nw356_ = C1()
                nw356_.ctor__()
                d_2276_v48_ = nw356_
                (globalState).f24 = (self).f26
                d_2277_v49_: _dafny.Seq
                d_2277_v49_ = _dafny.SeqWithoutIsStrInference([(d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))]])
                (globalState).f9 = not(default__.fm24(default__.fm24((self).f32, (self).f32, not((self).f32), (d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], globalState), (self).f32, ((0) - ((self).f26)) >= ((d_2277_v49_)[default__.safeIndex((d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], len(d_2277_v49_))]), (d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], globalState))
                d_2278_v50_: _dafny.Map
                d_2278_v50_ = _dafny.Map({(self).f26: (self).f26})
                def iife176_():
                    coll110_ = _dafny.Map()
                    compr_110_: int
                    for compr_110_ in (_dafny.SeqWithoutIsStrInference([465, (d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], len(d_2277_v49_)])).Elements:
                        d_2279_v51_: int = compr_110_
                        if (d_2279_v51_) in (_dafny.SeqWithoutIsStrInference([465, (d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], len(d_2277_v49_)])):
                            coll110_[(d_2279_v51_) + ((self).f26)] = (d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))]
                    return _dafny.Map(coll110_)
                rhs392_ = (iife176_()
                ).set((self).f26, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmf"))))
                rhs393_ = default__.safeDivisionInt((self).f26, len(((self).f27) + ((d_2273_v47_)[default__.safeIndex(118, (d_2273_v47_).length(0))])))
                rhs394_ = True
                rhs395_ = default__.safeDivisionInt(default__.safeDivisionInt((d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_2270_v45_ for d_2280_i6_ in range(default__.abs(535))]))), len(d_2222_v1_))
                rhs396_ = p0
                lhs353_ = globalState
                lhs354_ = globalState
                lhs355_ = globalState
                lhs356_ = globalState
                d_2278_v50_ = rhs392_
                lhs353_.f24 = rhs393_
                lhs354_.f25 = rhs394_
                lhs355_.f5 = rhs395_
                lhs356_.f9 = rhs396_
            (globalState).f9 = True
            index366_ = default__.safeIndex(401, (d_2251_v28_).length(0))
            (d_2251_v28_)[index366_] = (d_2251_v28_)[default__.safeIndex(401, (d_2251_v28_).length(0))]
        d_2281_v52_: _dafny.Map
        d_2281_v52_ = _dafny.Map({(self).f26: (self).f32})
        d_2282_v53_: _dafny.Map
        d_2282_v53_ = _dafny.Map({(self).f26: ((d_2281_v52_)[len((self).f27)] if (len((self).f27)) in (d_2281_v52_) else (self).f32)})
        d_2283_v54_: _dafny.Map
        d_2283_v54_ = _dafny.Map({p0: (self).fm7((self).f26, p0, globalState)})
        d_2284_v55_: _dafny.Seq
        d_2284_v55_ = _dafny.SeqWithoutIsStrInference([not(p0)])
        d_2285_v56_: _dafny.Map
        d_2285_v56_ = _dafny.Map({d_2284_v55_: p0})
        if (self).fm8(_dafny.MultiSet([756, len((self).f27)]), (not((self).f32)) == ((self).f32), (0) - ((self).fm6(d_2281_v52_, len((self).f27), default__.fm45(-163, d_2283_v54_, _dafny.Map({(self).f26: (self).f26}), globalState), d_2285_v56_, globalState)), (self).f26, globalState):
            (globalState).f5 = (0) - ((self).f26)
            d_2286_v57_: _dafny.Map
            d_2286_v57_ = _dafny.Map({(self).f27: (self).f26})
            d_2287_v58_: _dafny.Array
            nw357_ = _dafny.Array(None, 2)
            nw357_[int(0)] = d_2286_v57_
            nw357_[int(1)] = d_2286_v57_
            d_2287_v58_ = nw357_
            index367_ = default__.safeIndex(807, (d_2287_v58_).length(0))
            (d_2287_v58_)[index367_] = d_2286_v57_
            index368_ = default__.safeIndex(807, (d_2287_v58_).length(0))
            (d_2287_v58_)[index368_] = d_2286_v57_
            d_2283_v54_ = (d_2283_v54_).set(p0, (self).f26)
            d_2288_v59_: _dafny.Seq
            d_2288_v59_ = _dafny.SeqWithoutIsStrInference([(d_2284_v55_).set(default__.safeIndex(len((self).f27), len(d_2284_v55_)), p0), d_2284_v55_])
            (globalState).f9 = (d_2288_v59_) == (d_2288_v59_)
            if p0:
                index369_ = default__.safeIndex(517, (d_2220_v0_).length(0))
                (d_2220_v0_)[index369_] = (self).f32
                index370_ = default__.safeIndex(517, (d_2220_v0_).length(0))
                rhs397_ = not((self).f32)
                rhs398_ = False
                lhs357_ = globalState
                lhs358_ = d_2220_v0_
                lhs359_ = default__.safeIndex(517, (d_2220_v0_).length(0))
                lhs357_.f25 = rhs397_
                lhs358_[lhs359_] = rhs398_
                d_2289_v60_: _dafny.Array
                def lambda111_(d_2290_i7_):
                    return _dafny.CodePoint('u')

                init62_ = lambda111_
                nw358_ = _dafny.Array(None, 25)
                for i0_62_ in range(nw358_.length(0)):
                    nw358_[i0_62_] = init62_(i0_62_)
                d_2289_v60_ = nw358_
                d_2291_v61_: _dafny.Map
                d_2291_v61_ = _dafny.Map({d_2289_v60_: (self).f32})
                d_2292_v62_: _dafny.Seq
                d_2292_v62_ = _dafny.SeqWithoutIsStrInference([d_2291_v61_])
                d_2293_v63_: _dafny.Map
                d_2293_v63_ = _dafny.Map({(d_2292_v62_)[default__.safeIndex((self).f26, len(d_2292_v62_))]: (0) - (((self).f26 if (d_2220_v0_)[default__.safeIndex(517, (d_2220_v0_).length(0))] else len((self).f27)))})
                d_2293_v63_ = (d_2293_v63_).set(d_2291_v61_, (self).f26)
                d_2294_v64_: _dafny.Array
                nw359_ = _dafny.Array(int(0), 26)
                d_2294_v64_ = nw359_
                index371_ = default__.safeIndex(0, (d_2294_v64_).length(0))
                (d_2294_v64_)[index371_] = (self).f26
                index372_ = default__.safeIndex(0, (d_2294_v64_).length(0))
                (d_2294_v64_)[index372_] = (self).f26
                d_2295_v65_: _dafny.Array
                nw360_ = _dafny.Array(None, 2)
                nw360_[int(0)] = (d_2220_v0_)[default__.safeIndex(517, (d_2220_v0_).length(0))]
                nw360_[int(1)] = (self).f32
                d_2295_v65_ = nw360_
                d_2296_v66_: D7
                d_2296_v66_ = D7_DC17(d_2295_v65_)
                pat_let_tv50_ = d_2295_v65_
                def iife177_(_pat_let33_0):
                    def iife178_(d_2297_dt__update__tmp_h0_):
                        def iife179_(_pat_let34_0):
                            def iife180_(d_2298_dt__update_hcf33_h0_):
                                return D7_DC17(d_2298_dt__update_hcf33_h0_)
                            return iife180_(_pat_let34_0)
                        return iife179_(pat_let_tv50_)
                    return iife178_(_pat_let33_0)
                d_2296_v66_ = iife177_(d_2296_v66_)
                d_2299_v67_: _dafny.Map
                d_2299_v67_ = _dafny.Map({False: False})
                d_2299_v67_ = (d_2299_v67_).set((d_2220_v0_)[default__.safeIndex(517, (d_2220_v0_).length(0))], (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyixnofqf"))) <= ((self).f27))
            elif True:
                d_2300_v68_: _dafny.Set
                d_2300_v68_ = _dafny.Set({(self).f26})
                (globalState).f24 = len(_dafny.Map({((self).f26) >= (len(d_2300_v68_)): (self).f32}))
                d_2301_v69_: _dafny.Map
                d_2301_v69_ = _dafny.Map({p0: (self).f31})
                (globalState).f25 = (p0) in (d_2301_v69_)
                d_2302_v70_: C3
                nw361_ = C3()
                nw361_.ctor__((self).f26, (self).f26, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2303_i8_ in range(default__.abs(-647))])) + ((self).f27))
                d_2302_v70_ = nw361_
                d_2304_v71_: str
                d_2304_v71_ = _dafny.CodePoint('r')
                d_2305_v72_: _dafny.Seq
                d_2305_v72_ = _dafny.SeqWithoutIsStrInference([(self).f31])
                d_2306_v73_: _dafny.Seq
                d_2306_v73_ = _dafny.SeqWithoutIsStrInference([(self).f26, len(d_2222_v1_), 674, (self).f26])
                d_2307_v74_: _dafny.MultiSet
                d_2307_v74_ = _dafny.MultiSet([p0])
                d_2304_v71_ = default__.fm42(_dafny.Set({(d_2302_v70_).f34, len(d_2305_v72_), len(d_2306_v73_), (d_2307_v74_).cardinality}), globalState)
                (globalState).f5 = (self).f26
        elif True:
            d_2308_v75_: _dafny.MultiSet
            d_2308_v75_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2309_i9_ in range(default__.abs(808))])])
            d_2310_v76_: _dafny.MultiSet
            d_2310_v76_ = _dafny.MultiSet([(self).f26, len(_dafny.SeqWithoutIsStrInference([(d_2308_v75_).cardinality])), (self).f26])
            d_2311_v77_: D10
            d_2311_v77_ = D10_DC30(d_2310_v76_)
            source32_ = d_2311_v77_
            if source32_.is_DC31:
                d_2312___mcc_h0_ = source32_.cf58
                d_2313___mcc_h1_ = source32_.cf59
                d_2314___mcc_h2_ = source32_.cf60
                d_2315___mcc_h3_ = source32_.cf61
                d_2316_cf61_ = d_2315___mcc_h3_
                d_2317_cf60_ = d_2314___mcc_h2_
                d_2318_cf59_ = d_2313___mcc_h1_
                d_2319_cf58_ = d_2312___mcc_h0_
                (globalState).f25 = (((self).f26) * ((self).f26)) > ((self).f26)
                d_2316_cf61_ = (d_2316_cf61_ if (self).f32 else (0) - ((self).f26))
                (globalState).f25 = (((self).f26) <= (533) if (d_2284_v55_)[default__.safeIndex(len(_dafny.Map({len(d_2318_cf59_): p0})), len(d_2284_v55_))] else p0)
                d_2320_v78_: _dafny.Array
                nw362_ = _dafny.Array(int(0), 28)
                d_2320_v78_ = nw362_
                index373_ = default__.safeIndex(663, (d_2320_v78_).length(0))
                (d_2320_v78_)[index373_] = (self).f26
                index374_ = default__.safeIndex(663, (d_2320_v78_).length(0))
                (d_2320_v78_)[index374_] = (self).f26
            elif source32_.is_DC32:
                d_2321___mcc_h4_ = source32_.cf62
                d_2322___mcc_h5_ = source32_.cf63
                d_2323___mcc_h6_ = source32_.cf64
                d_2324_cf64_ = d_2323___mcc_h6_
                d_2325_cf63_ = d_2322___mcc_h5_
                d_2326_cf62_ = d_2321___mcc_h4_
                d_2327_v79_: _dafny.Seq
                d_2327_v79_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])
                d_2328_v80_: _dafny.Map
                d_2328_v80_ = _dafny.Map({(d_2327_v79_) + (_dafny.SeqWithoutIsStrInference([(self).f26 for d_2329_i10_ in range(default__.abs(240))])): (self).f27})
                d_2330_v81_: _dafny.Map
                d_2330_v81_ = _dafny.Map({(self).f32: d_2328_v80_})
                d_2328_v80_ = ((d_2328_v80_) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_2325_cf63_]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smx"))}))) | (((d_2330_v81_)[d_2326_cf62_] if (d_2326_cf62_) in (d_2330_v81_) else d_2328_v80_))
                (globalState).f25 = (self).f32
                d_2331_v82_: _dafny.MultiSet
                d_2331_v82_ = _dafny.MultiSet([_dafny.CodePoint('w')])
                d_2281_v52_ = _dafny.Map({(d_2331_v82_).cardinality: (self).f32})
                (globalState).f9 = (d_2326_cf62_ if (self).f32 else ((self).f32 if d_2324_cf64_ else default__.fm24(d_2324_cf64_, True, d_2324_cf64_, (self).f26, globalState)))
            elif True:
                d_2332___mcc_h7_ = source32_.cf57
                d_2333_cf57_ = d_2332___mcc_h7_
                d_2334_v83_: _dafny.Map
                d_2334_v83_ = _dafny.Map({True: d_2220_v0_})
                d_2335_v84_: str
                d_2335_v84_ = _dafny.CodePoint('s')
                d_2336_v85_: _dafny.Map
                d_2336_v85_ = _dafny.Map({d_2335_v84_: d_2220_v0_})
                d_2337_v86_: _dafny.Array
                nw363_ = _dafny.Array(None, 15)
                nw363_[int(0)] = d_2220_v0_
                nw363_[int(1)] = d_2220_v0_
                nw363_[int(2)] = ((d_2334_v83_)[not(p0)] if (not(p0)) in (d_2334_v83_) else d_2220_v0_)
                nw363_[int(3)] = d_2220_v0_
                nw363_[int(4)] = d_2220_v0_
                nw363_[int(5)] = d_2220_v0_
                nw363_[int(6)] = d_2220_v0_
                nw363_[int(7)] = d_2220_v0_
                nw363_[int(8)] = d_2220_v0_
                nw363_[int(9)] = d_2220_v0_
                nw363_[int(10)] = d_2220_v0_
                nw363_[int(11)] = d_2220_v0_
                nw363_[int(12)] = d_2220_v0_
                nw363_[int(13)] = (d_2220_v0_ if not((self).f32) else d_2220_v0_)
                nw363_[int(14)] = ((d_2336_v85_)[d_2335_v84_] if (d_2335_v84_) in (d_2336_v85_) else d_2220_v0_)
                d_2337_v86_ = nw363_
                index375_ = default__.safeIndex(123, (d_2337_v86_).length(0))
                (d_2337_v86_)[index375_] = d_2220_v0_
                index376_ = default__.safeIndex(123, (d_2337_v86_).length(0))
                nw364_ = _dafny.Array(False, 9)
                (d_2337_v86_)[index376_] = nw364_
                d_2338_v87_: _dafny.Seq
                d_2338_v87_ = _dafny.SeqWithoutIsStrInference([(self).fm7((self).f26, p0, globalState)])
                d_2339_v88_: _dafny.Map
                d_2339_v88_ = _dafny.Map({d_2335_v84_: (self).f32})
                d_2340_v89_: _dafny.Array
                nw365_ = _dafny.Array(None, 15)
                nw365_[int(0)] = (self).f26
                nw365_[int(1)] = len((d_2281_v52_) | (d_2282_v53_))
                nw365_[int(2)] = (self).f26
                nw365_[int(3)] = len(_dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (D8_DC24((self).f26, (self).f26, (self).f26)).cf46, (self).f26, 568]))
                nw365_[int(4)] = -710
                nw365_[int(5)] = (self).f26
                nw365_[int(6)] = default__.safeDivisionInt((self).f26, (self).f26)
                nw365_[int(7)] = (self).f26
                nw365_[int(8)] = (d_2338_v87_)[default__.safeIndex((self).f26, len(d_2338_v87_))]
                nw365_[int(9)] = (self).f26
                nw365_[int(10)] = (self).f26
                nw365_[int(11)] = ((self).f26) - ((self).f26)
                nw365_[int(12)] = ((self).f26) * (len(d_2284_v55_))
                nw365_[int(13)] = (self).f26
                nw365_[int(14)] = len(d_2339_v88_)
                d_2340_v89_ = nw365_
                index377_ = default__.safeIndex(941, (d_2340_v89_).length(0))
                (d_2340_v89_)[index377_] = (self).f26
                d_2341_v91_: D9
                d_2341_v91_ = D9_DC28()
                d_2342_v92_: _dafny.Map
                def iife181_():
                    coll111_ = _dafny.Map()
                    compr_111_: _dafny.Seq
                    for compr_111_ in (_dafny.SeqWithoutIsStrInference([d_2284_v55_ for d_2343_i11_ in range(default__.abs(656))])).Elements:
                        d_2344_v90_: _dafny.Seq = compr_111_
                        if (d_2344_v90_) in (_dafny.SeqWithoutIsStrInference([d_2284_v55_ for d_2343_i11_ in range(default__.abs(656))])):
                            coll111_[d_2344_v90_] = (D0_DC2(p0)).cf4
                    return _dafny.Map(coll111_)
                d_2342_v92_ = _dafny.Map({(self).fm6(d_2281_v52_, (self).f26, d_2338_v87_, iife181_()
                , globalState): d_2341_v91_})
                index378_ = default__.safeIndex(941, (d_2340_v89_).length(0))
                index379_ = default__.safeIndex(123, (d_2337_v86_).length(0))
                rhs399_ = (self).f26
                rhs400_ = (0) - (((self).f26) + ((0) - ((self).f26)))
                rhs401_ = default__.safeDivisionInt((self).f26, len((d_2342_v92_).set(937, d_2341_v91_)))
                rhs402_ = d_2220_v0_
                rhs403_ = ((((d_2310_v76_)[(self).fm7((self).f26, p0, globalState)] if ((self).fm7((self).f26, p0, globalState)) in (d_2310_v76_) else (self).f26)) + (len(d_2284_v55_)) if (((self).f27).set(default__.safeIndex((self).f26, len((self).f27)), d_2335_v84_)) == (((self).f27).set(default__.safeIndex(519, len((self).f27)), d_2335_v84_)) else (self).f26)
                lhs360_ = globalState
                lhs361_ = d_2340_v89_
                lhs362_ = default__.safeIndex(941, (d_2340_v89_).length(0))
                lhs363_ = globalState
                lhs364_ = d_2337_v86_
                lhs365_ = default__.safeIndex(123, (d_2337_v86_).length(0))
                lhs366_ = globalState
                lhs360_.f5 = rhs399_
                lhs361_[lhs362_] = rhs400_
                lhs363_.f24 = rhs401_
                lhs364_[lhs365_] = rhs402_
                lhs366_.f17 = rhs403_
                def lambda112_(d_2345_v89_):
                    def lambda113_(d_2346_i12_):
                        return (d_2346_i12_) + ((d_2345_v89_)[default__.safeIndex(941, (d_2345_v89_).length(0))])

                    return lambda113_

                init63_ = lambda112_(d_2340_v89_)
                nw366_ = _dafny.Array(None, 22)
                for i0_63_ in range(nw366_.length(0)):
                    nw366_[i0_63_] = init63_(i0_63_)
                d_2340_v89_ = nw366_
                d_2347_v93_: C1
                nw367_ = C1()
                nw367_.ctor__()
                d_2347_v93_ = nw367_
            d_2348_v94_: _dafny.Array
            def lambda114_(d_2349_i13_):
                return (d_2349_i13_) - ((self).f26)

            init64_ = lambda114_
            nw368_ = _dafny.Array(None, 1)
            for i0_64_ in range(nw368_.length(0)):
                nw368_[i0_64_] = init64_(i0_64_)
            d_2348_v94_ = nw368_
            index380_ = default__.safeIndex(184, (d_2348_v94_).length(0))
            (d_2348_v94_)[index380_] = (self).f26
            index381_ = default__.safeIndex(184, (d_2348_v94_).length(0))
            (d_2348_v94_)[index381_] = (0) - (((self).f26) - (((self).f26) * (len(d_2284_v55_))))
            d_2350_v95_: _dafny.Seq
            d_2350_v95_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_2348_v94_)[default__.safeIndex(184, (d_2348_v94_).length(0))])])
            d_2351_v96_: _dafny.MultiSet
            d_2351_v96_ = _dafny.MultiSet([(self).f32, (self).f32, p0])
            d_2352_v97_: _dafny.Set
            d_2352_v97_ = _dafny.Set({(self).f26, (d_2348_v94_)[default__.safeIndex(184, (d_2348_v94_).length(0))], (self).f26, (self).f26, (d_2310_v76_).cardinality})
            d_2350_v95_ = default__.fm21((self).f26, len(((self).f27) + ((self).f27)), not(((self).f26) != (((d_2310_v76_)[(self).f26] if ((self).f26) in (d_2310_v76_) else (d_2351_v96_).cardinality))), (d_2352_v97_).ispropersubset(d_2352_v97_), globalState)
            d_2353_v98_: _dafny.Seq
            d_2353_v98_ = _dafny.SeqWithoutIsStrInference([d_2284_v55_])
            d_2354_v99_: C3
            nw369_ = C3()
            nw369_.ctor__((0) - ((self).f26), (d_2350_v95_)[default__.safeIndex(len(((d_2353_v98_)[default__.safeIndex((self).f26, len(d_2353_v98_))]).set(default__.safeIndex((self).f26, len((d_2353_v98_)[default__.safeIndex((self).f26, len(d_2353_v98_))])), p0)), len(d_2350_v95_))], default__.fm36((self).f32, p0, p0, globalState))
            d_2354_v99_ = nw369_
            index382_ = default__.safeIndex(892, (d_2220_v0_).length(0))
            (d_2220_v0_)[index382_] = not(p0)
            index383_ = default__.safeIndex(892, (d_2220_v0_).length(0))
            (d_2220_v0_)[index383_] = ((d_2354_v99_).f34) < ((self).f26)
        hi11_ = (self).f26
        for d_2355_i14_ in range((self).f26, hi11_):
            d_2356_v100_: _dafny.Map
            d_2356_v100_ = _dafny.Map({(self).f32: (d_2284_v55_)[default__.safeIndex(-684, len(d_2284_v55_))]})
            d_2356_v100_ = (d_2356_v100_).set(p0, not((self).f32))
            d_2357_v101_: _dafny.Seq
            d_2357_v101_ = _dafny.SeqWithoutIsStrInference([d_2355_i14_])
            d_2357_v101_ = (d_2357_v101_) + (d_2357_v101_)
            d_2358_v102_: _dafny.Set
            d_2358_v102_ = _dafny.Set({len((d_2357_v101_) + (d_2357_v101_)), (self).f26})
            d_2358_v102_ = ((d_2358_v102_).intersection(d_2358_v102_)) | (_dafny.Set({(self).f26}))
            d_2359_v103_: _dafny.Array
            nw370_ = _dafny.Array(int(0), 25)
            d_2359_v103_ = nw370_
            d_2360_v104_: D7
            d_2360_v104_ = D7_DC18(len(d_2358_v102_), d_2359_v103_)
            d_2361_v105_: _dafny.Array
            nw371_ = _dafny.Array(None, 29)
            nw371_[int(0)] = d_2359_v103_
            nw371_[int(1)] = d_2359_v103_
            nw371_[int(2)] = d_2359_v103_
            nw371_[int(3)] = d_2359_v103_
            nw371_[int(4)] = d_2359_v103_
            nw371_[int(5)] = d_2359_v103_
            nw371_[int(6)] = d_2359_v103_
            nw371_[int(7)] = d_2359_v103_
            nw371_[int(8)] = d_2359_v103_
            nw371_[int(9)] = d_2359_v103_
            nw371_[int(10)] = (d_2360_v104_).cf35
            nw371_[int(11)] = d_2359_v103_
            nw371_[int(12)] = (d_2359_v103_ if (self).f32 else d_2359_v103_)
            nw371_[int(13)] = d_2359_v103_
            nw371_[int(14)] = d_2359_v103_
            nw371_[int(15)] = (d_2359_v103_ if True else d_2359_v103_)
            nw371_[int(16)] = d_2359_v103_
            nw371_[int(17)] = d_2359_v103_
            nw371_[int(18)] = d_2359_v103_
            nw371_[int(19)] = d_2359_v103_
            nw371_[int(20)] = d_2359_v103_
            nw371_[int(21)] = d_2359_v103_
            nw371_[int(22)] = d_2359_v103_
            nw371_[int(23)] = d_2359_v103_
            nw371_[int(24)] = (d_2359_v103_ if p0 else d_2359_v103_)
            nw371_[int(25)] = d_2359_v103_
            nw371_[int(26)] = d_2359_v103_
            nw371_[int(27)] = d_2359_v103_
            nw371_[int(28)] = d_2359_v103_
            d_2361_v105_ = nw371_
            index384_ = default__.safeIndex(269, (d_2361_v105_).length(0))
            (d_2361_v105_)[index384_] = d_2359_v103_
            index385_ = default__.safeIndex(269, (d_2361_v105_).length(0))
            def lambda115_(d_2362_i15_):
                return (d_2362_i15_) * ((0) - ((self).f26))

            init65_ = lambda115_
            nw372_ = _dafny.Array(None, 2)
            for i0_65_ in range(nw372_.length(0)):
                nw372_[i0_65_] = init65_(i0_65_)
            (d_2361_v105_)[index385_] = nw372_
        if True:
            d_2363_v106_: _dafny.Map
            d_2363_v106_ = _dafny.Map({p0: (self).f32})
            d_2364_v107_: D3
            d_2364_v107_ = D3_DC10(d_2363_v106_)
            d_2365_v108_: D12
            d_2365_v108_ = D12_DC37((self).f26, 903, d_2364_v107_, (self).f26, (self).f26)
            d_2366_v109_: str
            d_2366_v109_ = _dafny.CodePoint('a')
            (globalState).f13 = ((((self).f27).set(default__.safeIndex((self).f26, len((self).f27)), _dafny.CodePoint('q'))) + ((self).f27)).set(default__.safeIndex((d_2365_v108_).cf74, len((((self).f27).set(default__.safeIndex((self).f26, len((self).f27)), _dafny.CodePoint('q'))) + ((self).f27))), d_2366_v109_)
            index386_ = default__.safeIndex(270, (d_2220_v0_).length(0))
            (d_2220_v0_)[index386_] = False
            d_2367_v110_: _dafny.Seq
            d_2367_v110_ = _dafny.SeqWithoutIsStrInference([(self).f26])
            d_2368_v111_: D6
            d_2368_v111_ = D6_DC15(d_2367_v110_)
            d_2369_v112_: _dafny.MultiSet
            d_2369_v112_ = _dafny.MultiSet([True])
            d_2370_v113_: _dafny.Map
            d_2370_v113_ = _dafny.Map({(d_2369_v112_) - (d_2369_v112_): (self).f32})
            d_2371_v114_: _dafny.Set
            d_2371_v114_ = _dafny.Set({len((self).fm13((self).f32, (self).f27, (self).fm6(d_2282_v53_, (self).f26, d_2367_v110_, d_2285_v56_, globalState), globalState)), 718, (self).f26})
            d_2372_v115_: D21
            d_2372_v115_ = D21_DC56(d_2371_v114_)
            pat_let_tv51_ = d_2282_v53_
            index387_ = default__.safeIndex(270, (d_2220_v0_).length(0))
            def iife182_(_pat_let35_0):
                def iife183_(d_2373_dt__update__tmp_h1_):
                    def iife184_(_pat_let36_0):
                        def iife185_(d_2374_dt__update_hcf30_h0_):
                            return D6_DC15(d_2374_dt__update_hcf30_h0_)
                        return iife185_(_pat_let36_0)
                    return iife184_(_dafny.SeqWithoutIsStrInference([len(pat_let_tv51_), 210, (self).f26]))
                return iife183_(_pat_let35_0)
            rhs404_ = True
            rhs405_ = iife182_(d_2368_v111_)
            rhs406_ = not(((d_2372_v115_).cf107).issubset(d_2371_v114_))
            rhs407_ = (self).f26
            rhs408_ = ((_dafny.Map({d_2369_v112_: (self).f32}) if (self).f32 else d_2370_v113_)) | (d_2370_v113_)
            lhs367_ = d_2220_v0_
            lhs368_ = default__.safeIndex(270, (d_2220_v0_).length(0))
            lhs369_ = globalState
            lhs370_ = globalState
            lhs367_[lhs368_] = rhs404_
            d_2368_v111_ = rhs405_
            lhs369_.f25 = rhs406_
            lhs370_.f0 = rhs407_
            d_2370_v113_ = rhs408_
            d_2375_v116_: _dafny.MultiSet
            d_2375_v116_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2366_v109_ for d_2376_i16_ in range(default__.abs(102))])])
            d_2377_v117_: _dafny.Map
            d_2377_v117_ = _dafny.Map({(self).f26: (self).f27})
            d_2378_v118_: _dafny.Array
            nw373_ = _dafny.Array(None, 5)
            nw373_[int(0)] = (self).f27
            nw373_[int(1)] = (self).f27
            nw373_[int(2)] = (self).f27
            nw373_[int(3)] = (D0_DC1((self).f27, len((self).f31), ((d_2375_v116_).set(_dafny.SeqWithoutIsStrInference([d_2366_v109_ for d_2379_i17_ in range(default__.abs(587))]), default__.abs((self).f26))).cardinality)).cf1
            nw373_[int(4)] = ((d_2377_v117_)[len((self).f27)] if (len((self).f27)) in (d_2377_v117_) else (self).f27)
            d_2378_v118_ = nw373_
            d_2380_v119_: _dafny.Map
            d_2380_v119_ = _dafny.Map({(self).f26: d_2378_v118_})
            d_2380_v119_ = (d_2380_v119_).set((self).f26, d_2378_v118_)
            d_2366_v109_ = _dafny.CodePoint('r')
            d_2363_v106_ = (d_2363_v106_).set(p0, default__.fm24((self).f32, p0, (d_2220_v0_)[default__.safeIndex(270, (d_2220_v0_).length(0))], len((self).f27), globalState))
        elif True:
            d_2381_v120_: _dafny.Map
            d_2381_v120_ = _dafny.Map({p0: default__.fm24(p0, (self).f32, (self).f32, (self).f26, globalState)})
            d_2381_v120_ = d_2381_v120_
            if ((self).f26) != ((-465) * ((self).f26)):
                d_2382_v121_: D0
                d_2382_v121_ = D0_DC1((self).f27, (self).f26, ((d_2283_v54_)[p0] if (p0) in (d_2283_v54_) else (self).f26))
                d_2383_v122_: D0
                d_2383_v122_ = D0_DC4(d_2382_v121_)
                d_2384_v123_: _dafny.Map
                d_2384_v123_ = _dafny.Map({d_2383_v122_: (self).f26})
                pat_let_tv52_ = d_2382_v121_
                def iife186_(_pat_let37_0):
                    def iife187_(d_2385_dt__update__tmp_h2_):
                        def iife188_(_pat_let38_0):
                            def iife189_(d_2386_dt__update_hcf7_h0_):
                                return D0_DC4(d_2386_dt__update_hcf7_h0_)
                            return iife189_(_pat_let38_0)
                        return iife188_(pat_let_tv52_)
                    return iife187_(_pat_let37_0)
                d_2384_v123_ = (d_2384_v123_).set(iife186_(d_2383_v122_), (self).f26)
                d_2387_v124_: str
                d_2387_v124_ = _dafny.CodePoint('p')
                d_2388_v125_: _dafny.MultiSet
                d_2388_v125_ = _dafny.MultiSet([d_2387_v124_, d_2387_v124_, d_2387_v124_, d_2387_v124_, d_2387_v124_])
                (globalState).f9 = (d_2387_v124_) in ((d_2388_v125_) - (d_2388_v125_))
                d_2389_v126_: _dafny.Array
                nw374_ = _dafny.Array(None, 20)
                d_2389_v126_ = nw374_
                d_2390_v127_: C0
                nw375_ = C0()
                nw375_.ctor__()
                d_2390_v127_ = nw375_
                index388_ = default__.safeIndex(745, (d_2389_v126_).length(0))
                (d_2389_v126_)[index388_] = d_2390_v127_
                index389_ = default__.safeIndex(745, (d_2389_v126_).length(0))
                (d_2389_v126_)[index389_] = d_2390_v127_
                d_2391_v128_: _dafny.Seq
                d_2391_v128_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_2392_v129_: _dafny.MultiSet
                d_2392_v129_ = _dafny.MultiSet([True, (d_2284_v55_)[default__.safeIndex(-429, len(d_2284_v55_))], p0, p0])
                d_2393_v130_: _dafny.Set
                d_2393_v130_ = _dafny.Set({d_2392_v129_})
                d_2394_v131_: _dafny.MultiSet
                d_2394_v131_ = _dafny.MultiSet([d_2393_v130_, d_2393_v130_])
                rhs409_ = (len(d_2391_v128_)) + (((self).f26) - ((self).f26))
                rhs410_ = (d_2390_v127_).fm16(((d_2394_v131_)[_dafny.Set({d_2392_v129_})] if (_dafny.Set({d_2392_v129_})) in (d_2394_v131_) else (self).f26), (self).f26, (self).f26, (self).f32, globalState)
                rhs411_ = p0
                lhs371_ = globalState
                lhs372_ = globalState
                lhs373_ = globalState
                lhs371_.f8 = rhs409_
                lhs372_.f17 = rhs410_
                lhs373_.f25 = rhs411_
                d_2395_v132_: _dafny.Array
                nw376_ = _dafny.Array(_dafny.CodePoint('D'), 5)
                d_2395_v132_ = nw376_
                d_2396_v133_: _dafny.Map
                d_2396_v133_ = _dafny.Map({(self).f26: d_2395_v132_})
                d_2396_v133_ = (d_2396_v133_).set(default__.safeModuloInt((self).f26, (self).f26), d_2395_v132_)
            elif True:
                d_2397_v134_: _dafny.Array
                nw377_ = _dafny.Array(int(0), 3)
                d_2397_v134_ = nw377_
                d_2398_v135_: _dafny.Seq
                d_2398_v135_ = _dafny.SeqWithoutIsStrInference([d_2397_v134_, d_2397_v134_])
                d_2398_v135_ = (d_2398_v135_) + ((_dafny.SeqWithoutIsStrInference([d_2397_v134_, d_2397_v134_])) + (d_2398_v135_))
                d_2399_v136_: _dafny.Seq
                d_2399_v136_ = _dafny.SeqWithoutIsStrInference([d_2220_v0_])
                d_2400_v137_: _dafny.MultiSet
                d_2400_v137_ = _dafny.MultiSet([(self).f26, -351])
                d_2401_v138_: _dafny.Map
                d_2401_v138_ = _dafny.Map({(self).f32: d_2400_v137_})
                rhs412_ = p0
                rhs413_ = ((d_2399_v136_ if (self).f32 else d_2399_v136_)) + (_dafny.SeqWithoutIsStrInference([d_2220_v0_, d_2220_v0_]))
                rhs414_ = (d_2401_v138_ if p0 else d_2401_v138_)
                lhs374_ = globalState
                lhs374_.f25 = rhs412_
                d_2399_v136_ = rhs413_
                d_2401_v138_ = rhs414_
                d_2402_v139_: D0
                d_2402_v139_ = D0_DC2((self).f32)
                d_2403_v140_: D3
                d_2403_v140_ = D3_DC10(d_2381_v120_)
                d_2404_v141_: _dafny.Map
                d_2404_v141_ = _dafny.Map({p0: (self).f27})
                d_2405_v142_: str
                d_2405_v142_ = _dafny.CodePoint('e')
                d_2406_v143_: _dafny.Map
                d_2406_v143_ = _dafny.Map({d_2405_v142_: d_2381_v120_})
                d_2407_v144_: _dafny.Map
                d_2407_v144_ = _dafny.Map({p0: d_2381_v120_})
                d_2408_v145_: _dafny.Array
                nw378_ = _dafny.Array(None, 27)
                nw378_[int(0)] = ((d_2381_v120_).set(p0, (d_2402_v139_).cf4)) | (_dafny.Map({False: (self).f32}))
                nw378_[int(1)] = (d_2403_v140_).cf20
                nw378_[int(2)] = (d_2381_v120_) | (d_2381_v120_)
                nw378_[int(3)] = d_2381_v120_
                nw378_[int(4)] = _dafny.Map({p0: p0})
                nw378_[int(5)] = _dafny.Map({not((self).f32): p0})
                nw378_[int(6)] = _dafny.Map({(self).f32: p0})
                nw378_[int(7)] = d_2381_v120_
                nw378_[int(8)] = d_2381_v120_
                nw378_[int(9)] = (d_2381_v120_) | (d_2381_v120_)
                nw378_[int(10)] = d_2381_v120_
                nw378_[int(11)] = d_2381_v120_
                nw378_[int(12)] = (d_2381_v120_) | (d_2381_v120_)
                nw378_[int(13)] = (default__.fm34(not(True), p0, (self).f32, (self).f26, globalState) if (self).f32 else d_2381_v120_)
                nw378_[int(14)] = d_2381_v120_
                nw378_[int(15)] = (_dafny.Map({(d_2284_v55_)[default__.safeIndex(len(((d_2404_v141_)[(self).f32] if ((self).f32) in (d_2404_v141_) else (self).f27)), len(d_2284_v55_))]: (self).f32})) | (d_2381_v120_)
                nw378_[int(16)] = (d_2381_v120_) | (d_2381_v120_)
                nw378_[int(17)] = _dafny.Map({p0: (self).f32})
                nw378_[int(18)] = ((d_2406_v143_)[d_2405_v142_] if (d_2405_v142_) in (d_2406_v143_) else _dafny.Map({p0: p0}))
                nw378_[int(19)] = d_2381_v120_
                nw378_[int(20)] = d_2381_v120_
                nw378_[int(21)] = ((d_2407_v144_)[not((self).f32)] if (not((self).f32)) in (d_2407_v144_) else d_2381_v120_)
                nw378_[int(22)] = (d_2381_v120_) | (d_2381_v120_)
                nw378_[int(23)] = (d_2381_v120_) | (d_2381_v120_)
                nw378_[int(24)] = d_2381_v120_
                nw378_[int(25)] = (d_2381_v120_).set(p0, (self).f32)
                nw378_[int(26)] = _dafny.Map({not(p0): (self).f32})
                d_2408_v145_ = nw378_
                index390_ = default__.safeIndex(375, (d_2408_v145_).length(0))
                (d_2408_v145_)[index390_] = _dafny.Map({default__.fm24(p0, (self).f32, p0, (self).f26, globalState): False})
                index391_ = default__.safeIndex(375, (d_2408_v145_).length(0))
                (d_2408_v145_)[index391_] = d_2381_v120_
                index392_ = default__.safeIndex(826, (d_2397_v134_).length(0))
                (d_2397_v134_)[index392_] = (0) - (len((self).f31))
                index393_ = default__.safeIndex(826, (d_2397_v134_).length(0))
                (d_2397_v134_)[index393_] = (self).f26
                d_2409_v146_: C0
                nw379_ = C0()
                nw379_.ctor__()
                d_2409_v146_ = nw379_
                d_2410_v147_: _dafny.Seq
                d_2410_v147_ = _dafny.SeqWithoutIsStrInference([d_2409_v146_, d_2409_v146_, d_2409_v146_])
                d_2411_v148_: _dafny.MultiSet
                d_2411_v148_ = _dafny.MultiSet([d_2410_v147_])
                d_2412_v149_: _dafny.MultiSet
                d_2412_v149_ = _dafny.MultiSet([(self).f27, (self).f27])
                d_2413_v150_: _dafny.Map
                d_2413_v150_ = _dafny.Map({d_2411_v148_: (d_2412_v149_) | (d_2412_v149_)})
                d_2413_v150_ = (d_2413_v150_).set((d_2411_v148_).intersection(d_2411_v148_), d_2412_v149_)
            d_2414_v151_: _dafny.Map
            d_2414_v151_ = _dafny.Map({(self).f26: len((self).f27)})
            d_2414_v151_ = (d_2414_v151_).set(len(d_2284_v55_), (self).f26)
            d_2283_v54_ = (d_2283_v54_).set((self).f32, (self).f26)
            d_2415_v152_: _dafny.Array
            nw380_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_2415_v152_ = nw380_
            d_2416_v153_: D21
            d_2416_v153_ = D21_DC58((self).f26, p0, d_2415_v152_, (self).f26, p0)
            d_2417_v154_: D21
            d_2417_v154_ = D21_DC60(d_2416_v153_)
            d_2418_v155_: str
            d_2418_v155_ = _dafny.CodePoint('i')
            d_2419_v156_: D21
            d_2419_v156_ = D21_DC57(d_2418_v155_, (self).f26, (self).f26, (self).f26)
            d_2420_v157_: _dafny.Map
            d_2420_v157_ = _dafny.Map({d_2417_v154_: (d_2419_v156_).cf110})
            d_2421_v158_: D19
            d_2421_v158_ = D19_DC53((self).f26, (self).f32)
            rhs415_ = (self).f26
            rhs416_ = (default__.fm24(p0, (d_2421_v158_).cf101, (self).f32, (self).f26, globalState) if (len(d_2420_v157_)) != (579) else (self).f32)
            lhs375_ = globalState
            lhs376_ = globalState
            lhs375_.f0 = rhs415_
            lhs376_.f25 = rhs416_
        r0 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "typo"))) + (((self).f27) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xc"))))
        return r0

    def m4(self, p0, p1, globalState):
        r0: int = int(0)
        d_2422_v0_: _dafny.Array
        nw381_ = _dafny.Array(_dafny.CodePoint('D'), 16)
        d_2422_v0_ = nw381_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_2422_v0_).length(0)):
            d_2423_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_2423_i0_)) and ((d_2423_i0_) < ((d_2422_v0_).length(0)))):
                (d_2422_v0_)[(d_2423_i0_)] = _dafny.CodePoint('d')
        (globalState).f8 = (self).f26
        (globalState).f17 = p0
        d_2424_v1_: _dafny.Array
        def lambda116_(d_2425_p0_):
            def lambda117_(d_2426_i1_):
                return (d_2425_p0_) != ((_dafny.MultiSet([(self).f26])).cardinality)

            return lambda117_

        init66_ = lambda116_(p0)
        nw382_ = _dafny.Array(None, 8)
        for i0_66_ in range(nw382_.length(0)):
            nw382_[i0_66_] = init66_(i0_66_)
        d_2424_v1_ = nw382_
        index394_ = default__.safeIndex(697, (d_2424_v1_).length(0))
        (d_2424_v1_)[index394_] = ((self).f32) == ((self).f32)
        d_2427_v2_: _dafny.MultiSet
        d_2427_v2_ = _dafny.MultiSet([(self).f32, (self).f32, (self).f32])
        d_2428_v3_: D15
        d_2428_v3_ = D15_DC42(d_2427_v2_)
        index395_ = default__.safeIndex(697, (d_2424_v1_).length(0))
        (d_2424_v1_)[index395_] = ((d_2428_v3_).cf80).issubset((d_2427_v2_ if (self).f32 else d_2427_v2_))
        if not((self).f32):
            (globalState).f24 = (0) - (default__.safeModuloInt((0) - (p0), p0))
            r0 = (self).f26
            d_2429_v4_: _dafny.Map
            d_2429_v4_ = _dafny.Map({(self).f26: (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]})
            d_2430_v5_: D3
            d_2430_v5_ = D3_DC11((self).f31, True, (self).f26, p0)
            d_2431_v6_: _dafny.Set
            d_2431_v6_ = _dafny.Set({(self).f32, (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]})
            d_2432_v7_: _dafny.Array
            nw383_ = _dafny.Array(int(0), 14)
            d_2432_v7_ = nw383_
            d_2433_v8_: _dafny.Set
            d_2433_v8_ = _dafny.Set({d_2432_v7_, d_2432_v7_})
            pat_let_tv53_ = p0
            pat_let_tv54_ = d_2431_v6_
            pat_let_tv55_ = d_2433_v8_
            d_2434_v9_: _dafny.Map
            def iife190_(_pat_let39_0):
                def iife191_(d_2435_dt__update__tmp_h0_):
                    def iife192_(_pat_let40_0):
                        def iife193_(d_2436_dt__update_hcf23_h0_):
                            def iife194_(_pat_let41_0):
                                def iife195_(d_2437_dt__update_hcf21_h0_):
                                    return D3_DC11(d_2437_dt__update_hcf21_h0_, (d_2435_dt__update__tmp_h0_).cf22, d_2436_dt__update_hcf23_h0_, (d_2435_dt__update__tmp_h0_).cf24)
                                return iife195_(_pat_let41_0)
                            return iife194_(_dafny.Map({len(pat_let_tv54_): len(pat_let_tv55_)}))
                        return iife193_(_pat_let40_0)
                    return iife192_(pat_let_tv53_)
                return iife191_(_pat_let39_0)
            d_2434_v9_ = _dafny.Map({(self).f26: (iife190_(d_2430_v5_)).cf21})
            d_2438_v10_: _dafny.Array
            nw384_ = _dafny.Array(None, 15)
            nw384_[int(0)] = (self).f31
            nw384_[int(1)] = (self).f31
            nw384_[int(2)] = (self).f31
            nw384_[int(3)] = _dafny.Map({p0: (self).f26})
            nw384_[int(4)] = (self).f31
            nw384_[int(5)] = (self).f31
            nw384_[int(6)] = (self).f31
            nw384_[int(7)] = (self).f31
            nw384_[int(8)] = (self).f31
            nw384_[int(9)] = (self).f31
            nw384_[int(10)] = default__.fm40((d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))], (self).f26, d_2429_v4_, (self).f32, globalState)
            nw384_[int(11)] = ((self).f31).set(len((_dafny.SeqWithoutIsStrInference([(self).f32, (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))], not((self).f32), (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]])).set(default__.safeIndex(len(_dafny.Map({(d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]: len(_dafny.SeqWithoutIsStrInference([p0 for d_2439_i2_ in range(default__.abs(376))]))})), len(_dafny.SeqWithoutIsStrInference([(self).f32, (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))], not((self).f32), (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]]))), (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))])), p0)
            nw384_[int(12)] = (self).f31
            nw384_[int(13)] = ((d_2434_v9_)[(self).f26] if ((self).f26) in (d_2434_v9_) else (self).f31)
            nw384_[int(14)] = (self).f31
            d_2438_v10_ = nw384_
            index396_ = default__.safeIndex(890, (d_2438_v10_).length(0))
            (d_2438_v10_)[index396_] = (self).f31
            index397_ = default__.safeIndex(697, (d_2424_v1_).length(0))
            index398_ = default__.safeIndex(890, (d_2438_v10_).length(0))
            rhs417_ = (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]
            rhs418_ = (self).f26
            rhs419_ = p0
            rhs420_ = (self).f31
            lhs377_ = d_2424_v1_
            lhs378_ = default__.safeIndex(697, (d_2424_v1_).length(0))
            lhs379_ = globalState
            lhs380_ = globalState
            lhs381_ = d_2438_v10_
            lhs382_ = default__.safeIndex(890, (d_2438_v10_).length(0))
            lhs377_[lhs378_] = rhs417_
            lhs379_.f24 = rhs418_
            lhs380_.f8 = rhs419_
            lhs381_[lhs382_] = rhs420_
            d_2440_v11_: C0
            nw385_ = C0()
            nw385_.ctor__()
            d_2440_v11_ = nw385_
            d_2441_v12_: _dafny.Seq
            d_2441_v12_ = _dafny.SeqWithoutIsStrInference([(self).f32])
            (globalState).f24 = ((0) - (default__.safeModuloInt(len((d_2441_v12_).set(default__.safeIndex(p0, len(d_2441_v12_)), (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))])), (self).f26))) - (p0)
        elif True:
            d_2442_v13_: C1
            nw386_ = C1()
            nw386_.ctor__()
            d_2442_v13_ = nw386_
            d_2443_v14_: int
            d_2444_v15_: _dafny.MultiSet
            d_2445_v16_: _dafny.Array
            out55_: int
            out56_: _dafny.MultiSet
            out57_: _dafny.Array
            out55_, out56_, out57_ = (d_2442_v13_).m2((d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))], default__.fm49(len(_dafny.SeqWithoutIsStrInference([(self).f26])), globalState), globalState)
            d_2443_v14_ = out55_
            d_2444_v15_ = out56_
            d_2445_v16_ = out57_
            d_2446_v17_: _dafny.Map
            d_2446_v17_ = _dafny.Map({(self).f32: (self).f32})
            d_2447_v18_: _dafny.Set
            d_2447_v18_ = _dafny.Set({d_2443_v14_, d_2443_v14_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjs")))})
            d_2448_v19_: _dafny.Seq
            d_2448_v19_ = _dafny.SeqWithoutIsStrInference([d_2443_v14_, d_2443_v14_])
            d_2446_v17_ = (d_2446_v17_).set((d_2443_v14_) != ((self).f26), (len(d_2447_v18_)) != (len(d_2448_v19_)))
            d_2449_v20_: D0
            d_2449_v20_ = D0_DC2((self).f32)
            d_2450_v21_: int
            d_2451_v22_: _dafny.MultiSet
            d_2452_v23_: _dafny.Array
            out58_: int
            out59_: _dafny.MultiSet
            out60_: _dafny.Array
            out58_, out59_, out60_ = (d_2442_v13_).m2((d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))], d_2449_v20_, globalState)
            d_2450_v21_ = out58_
            d_2451_v22_ = out59_
            d_2452_v23_ = out60_
            d_2453_v24_: _dafny.Array
            def lambda118_(d_2454_v21_):
                def lambda119_(d_2455_i3_):
                    return default__.safeModuloInt(d_2455_i3_, d_2454_v21_)

                return lambda119_

            init67_ = lambda118_(d_2450_v21_)
            nw387_ = _dafny.Array(None, 29)
            for i0_67_ in range(nw387_.length(0)):
                nw387_[i0_67_] = init67_(i0_67_)
            d_2453_v24_ = nw387_
            index399_ = default__.safeIndex(278, (d_2452_v23_).length(0))
            (d_2452_v23_)[index399_] = d_2453_v24_
            index400_ = default__.safeIndex(278, (d_2452_v23_).length(0))
            (d_2452_v23_)[index400_] = d_2453_v24_
        if (D0_DC2((d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))])).cf4:
            d_2456_v25_: str
            d_2456_v25_ = _dafny.CodePoint('b')
            d_2457_v26_: _dafny.Map
            d_2457_v26_ = _dafny.Map({d_2456_v25_: (self).f26})
            d_2458_v27_: _dafny.Seq
            d_2458_v27_ = _dafny.SeqWithoutIsStrInference([(d_2457_v26_) | (d_2457_v26_)])
            d_2458_v27_ = ((d_2458_v27_) + (d_2458_v27_)).set(default__.safeIndex((self).f26, len((d_2458_v27_) + (d_2458_v27_))), d_2457_v26_)
            d_2459_v28_: T1
            nw388_ = C3()
            nw388_.ctor__((self).f26, p0, (self).f27)
            d_2459_v28_ = nw388_
            d_2460_v29_: _dafny.Map
            d_2460_v29_ = _dafny.Map({d_2459_v28_: (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]})
            d_2461_v30_: _dafny.Map
            d_2461_v30_ = _dafny.Map({p0: (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]})
            d_2462_v31_: _dafny.Seq
            d_2462_v31_ = _dafny.SeqWithoutIsStrInference([(self).f26])
            d_2463_v32_: _dafny.Seq
            d_2463_v32_ = _dafny.SeqWithoutIsStrInference([(self).f32])
            d_2464_v33_: _dafny.Map
            d_2464_v33_ = _dafny.Map({d_2463_v32_: False})
            d_2465_v34_: _dafny.Array
            nw389_ = _dafny.Array(None, 12)
            nw389_[int(0)] = p0
            nw389_[int(1)] = (self).f26
            nw389_[int(2)] = (self).f26
            nw389_[int(3)] = len((self).f27)
            nw389_[int(4)] = (self).f26
            nw389_[int(5)] = (self).fm7(p0, (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))], globalState)
            nw389_[int(6)] = (p0) - (len(d_2460_v29_))
            nw389_[int(7)] = (self).f26
            nw389_[int(8)] = default__.safeDivisionInt(len(default__.fm50(globalState)), len((d_2459_v28_).f27))
            nw389_[int(9)] = default__.safeDivisionInt(p0, (self).f26)
            nw389_[int(10)] = (d_2459_v28_).fm6(d_2461_v30_, (self).f26, d_2462_v31_, d_2464_v33_, globalState)
            nw389_[int(11)] = (self).f26
            d_2465_v34_ = nw389_
            d_2465_v34_ = d_2465_v34_
            d_2466_v35_: _dafny.Map
            d_2466_v35_ = _dafny.Map({d_2428_v3_: d_2459_v28_})
            if (len(d_2466_v35_)) not in (default__.fm40((d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))], p0, d_2461_v30_, True, globalState)):
                d_2467_v36_: _dafny.Map
                d_2467_v36_ = _dafny.Map({(self).f31: default__.safeDivisionInt((self).f26, (d_2459_v28_).f26)})
                d_2467_v36_ = ((default__.fm51(p0, False, globalState)) | (d_2467_v36_)) | (d_2467_v36_)
                d_2468_v37_: _dafny.Map
                d_2468_v37_ = _dafny.Map({d_2424_v1_: d_2422_v0_})
                d_2469_v38_: _dafny.Map
                d_2469_v38_ = _dafny.Map({(d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]: d_2468_v37_})
                d_2470_v39_: _dafny.Array
                nw390_ = _dafny.Array(None, 15)
                nw390_[int(0)] = _dafny.Map({d_2424_v1_: d_2422_v0_})
                nw390_[int(1)] = _dafny.Map({d_2424_v1_: d_2422_v0_})
                nw390_[int(2)] = d_2468_v37_
                nw390_[int(3)] = (d_2468_v37_) | (d_2468_v37_)
                nw390_[int(4)] = (((d_2469_v38_)[(self).f32] if ((self).f32) in (d_2469_v38_) else d_2468_v37_)) | (d_2468_v37_)
                nw390_[int(5)] = d_2468_v37_
                nw390_[int(6)] = (((d_2469_v38_)[(d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]] if ((d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]) in (d_2469_v38_) else d_2468_v37_)) | (d_2468_v37_)
                nw390_[int(7)] = d_2468_v37_
                nw390_[int(8)] = _dafny.Map({d_2424_v1_: d_2422_v0_})
                nw390_[int(9)] = d_2468_v37_
                nw390_[int(10)] = d_2468_v37_
                nw390_[int(11)] = d_2468_v37_
                nw390_[int(12)] = d_2468_v37_
                nw390_[int(13)] = d_2468_v37_
                nw390_[int(14)] = (d_2468_v37_) | (d_2468_v37_)
                d_2470_v39_ = nw390_
                index401_ = default__.safeIndex(94, (d_2470_v39_).length(0))
                (d_2470_v39_)[index401_] = d_2468_v37_
                index402_ = default__.safeIndex(94, (d_2470_v39_).length(0))
                (d_2470_v39_)[index402_] = (d_2468_v37_) | (d_2468_v37_)
                d_2471_v40_: _dafny.Map
                d_2471_v40_ = _dafny.Map({((d_2459_v28_).f27) < ((self).f27): (self).f31})
                d_2471_v40_ = (d_2471_v40_).set(True, ((self).f31) | ((self).f31))
                d_2472_v41_: _dafny.Map
                d_2472_v41_ = _dafny.Map({True: (self).f26})
                d_2473_v42_: C2
                nw391_ = C2()
                nw391_.ctor__(default__.fm45((self).f26, d_2472_v41_, (self).f31, globalState), (self).f31)
                d_2473_v42_ = nw391_
                index403_ = default__.safeIndex(697, (d_2424_v1_).length(0))
                (d_2424_v1_)[index403_] = not (default__.fm24((self).f32, (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))], (self).f32, len(_dafny.Set({d_2473_v42_})), globalState)) or (not(False))
                d_2465_v34_ = d_2465_v34_
            elif True:
                d_2474_v43_: _dafny.Map
                d_2474_v43_ = _dafny.Map({(self).f32: (d_2459_v28_).f26})
                d_2475_v44_: _dafny.Map
                d_2475_v44_ = _dafny.Map({d_2474_v43_: (self).f32})
                d_2476_v45_: _dafny.Map
                d_2476_v45_ = _dafny.Map({((d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))] if ((d_2475_v44_)[_dafny.Map({(self).f32: (d_2427_v2_).cardinality})] if (_dafny.Map({(self).f32: (d_2427_v2_).cardinality})) in (d_2475_v44_) else (self).f32) else (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]): not((self).f32)})
                d_2476_v45_ = (d_2476_v45_).set(not ((self).f32) or ((d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]), not((self).f32))
                (globalState).f25 = (self).f32
                (globalState).f9 = (self).f32
                index404_ = default__.safeIndex(697, (d_2424_v1_).length(0))
                (d_2424_v1_)[index404_] = not((self).f32)
                (globalState).f23 = (d_2459_v28_).f27
            rhs421_ = (p1).cf13
            rhs422_ = 96
            rhs423_ = p0
            rhs424_ = (d_2459_v28_).f26
            lhs383_ = globalState
            lhs384_ = globalState
            lhs385_ = globalState
            lhs386_ = globalState
            lhs383_.f9 = rhs421_
            lhs384_.f17 = rhs422_
            lhs385_.f8 = rhs423_
            lhs386_.f0 = rhs424_
            d_2477_v46_: _dafny.Array
            nw392_ = _dafny.Array(_dafny.Map({}), 23)
            d_2477_v46_ = nw392_
            d_2478_v48_: _dafny.Set
            d_2478_v48_ = _dafny.Set({d_2456_v25_, d_2456_v25_, _dafny.CodePoint('a')})
            index405_ = default__.safeIndex(98, (d_2477_v46_).length(0))
            def iife196_():
                coll112_ = _dafny.Map()
                compr_112_: str
                for compr_112_ in (d_2478_v48_).Elements:
                    d_2479_v47_: str = compr_112_
                    if (d_2479_v47_) in (d_2478_v48_):
                        coll112_[d_2479_v47_] = len((d_2459_v28_).f27)
                return _dafny.Map(coll112_)
            (d_2477_v46_)[index405_] = (iife196_()
            ) | (d_2457_v26_)
            index406_ = default__.safeIndex(98, (d_2477_v46_).length(0))
            rhs425_ = ((d_2457_v26_) | (d_2457_v26_)) | (d_2457_v26_)
            rhs426_ = (0) - (default__.safeDivisionInt(p0, (d_2459_v28_).f26))
            rhs427_ = (0) - (((d_2459_v28_).f26 if not(not ((self).f32) or (True)) else (d_2459_v28_).f26))
            lhs387_ = d_2477_v46_
            lhs388_ = default__.safeIndex(98, (d_2477_v46_).length(0))
            lhs389_ = globalState
            lhs387_[lhs388_] = rhs425_
            r0 = rhs426_
            lhs389_.f17 = rhs427_
        elif True:
            d_2480_v49_: _dafny.Array
            nw393_ = _dafny.Array(int(0), 3)
            d_2480_v49_ = nw393_
            index407_ = default__.safeIndex(673, (d_2480_v49_).length(0))
            (d_2480_v49_)[index407_] = (0) - (((self).f26) * (len(default__.fm14(globalState))))
            index408_ = default__.safeIndex(673, (d_2480_v49_).length(0))
            (d_2480_v49_)[index408_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2481_i4_ in range(default__.abs(-478))]))
            (globalState).f25 = (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]
            index409_ = default__.safeIndex(697, (d_2424_v1_).length(0))
            (d_2424_v1_)[index409_] = (self).f32
            d_2482_v50_: _dafny.Array
            nw394_ = _dafny.Array(None, 11)
            nw394_[int(0)] = (self).f27
            nw394_[int(1)] = (self).f27
            nw394_[int(2)] = (self).f27
            nw394_[int(3)] = (self).f27
            nw394_[int(4)] = (self).f27
            nw394_[int(5)] = (self).f27
            nw394_[int(6)] = ((self).f27) + ((self).f27)
            nw394_[int(7)] = (self).f27
            nw394_[int(8)] = (self).f27
            nw394_[int(9)] = (self).f27
            nw394_[int(10)] = ((self).f27) + ((self).f27)
            d_2482_v50_ = nw394_
            index410_ = default__.safeIndex(540, (d_2482_v50_).length(0))
            (d_2482_v50_)[index410_] = (self).f27
            index411_ = default__.safeIndex(540, (d_2482_v50_).length(0))
            (d_2482_v50_)[index411_] = (self).f27
            d_2483_v51_: _dafny.Map
            d_2483_v51_ = _dafny.Map({False: False})
            d_2483_v51_ = (d_2483_v51_).set(not((d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]), ((d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))] if (self).f32 else (self).f32))
        d_2484_v52_: _dafny.Seq
        d_2484_v52_ = _dafny.SeqWithoutIsStrInference([(self).f26, p0, p0])
        d_2485_v53_: _dafny.Map
        d_2485_v53_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([not((self).f32), not((self).f32), (d_2424_v1_)[default__.safeIndex(697, (d_2424_v1_).length(0))]]): (self).f32})
        r0 = (self).fm6(_dafny.Map({p0: not((self).f32)}), (self).f26, d_2484_v52_, d_2485_v53_, globalState)
        return r0

    @property
    def f32(self):
        return self._f32

class C5(T1, T0):
    def  __init__(self):
        self._f26: int = int(0)
        self._f27: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f29: int = int(0)
        self.f30: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f29, f30, f26, f27):
        (self).f29 = f29
        (self).f30 = f30
        (self)._f26 = f26
        (self)._f27 = f27

    def fm8(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Set({False})) - (_dafny.Set({False, False}))).ispropersubset((D1_DC5(_dafny.Set({False, True}))).cf8)

    def fm6(self, p0, p1, p2, p3, globalState):
        def iife197_():
            coll113_ = _dafny.Map()
            compr_113_: str
            for compr_113_ in (_dafny.MultiSet([_dafny.CodePoint('t'), _dafny.CodePoint('r')])).Elements:
                d_2486_v0_: str = compr_113_
                if (d_2486_v0_) in (_dafny.MultiSet([_dafny.CodePoint('t'), _dafny.CodePoint('r')])):
                    coll113_[d_2486_v0_] = True
            return _dafny.Map(coll113_)
        return len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([self.f29, (_dafny.MultiSet([not(True)])).cardinality]), _dafny.SeqWithoutIsStrInference([628, len(_dafny.Map({_dafny.Set({_dafny.CodePoint('e')}): self.f30})), (0) - (self.f30)]), _dafny.SeqWithoutIsStrInference([len(iife197_()
        ), self.f30, len((D2_DC7(_dafny.SeqWithoutIsStrInference([False, False]))).cf12)])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f26 for d_2487_i1_ in range(default__.abs(-719))]) for d_2488_i0_ in range(default__.abs(-840))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(((self).f27).set(default__.safeIndex(self.f29, len((self).f27)), _dafny.CodePoint('s'))) for d_2489_i2_ in range(default__.abs(486))]))]), _dafny.SeqWithoutIsStrInference([(self).f26, self.f30])]))))

    def fm7(self, p0, p1, globalState):
        return (self.f29) + ((len(_dafny.Map({True: not(True)}))) + (self.f30))

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        d_2490_i0_: int
        d_2490_i0_ = 0
        with _dafny.label("5"):
            while not(p1):
                with _dafny.c_label("5"):
                    if (d_2490_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_2490_i0_ = (d_2490_i0_) + (1)
                    (globalState).f24 = self.f30
                    d_2491_v0_: _dafny.Array
                    def lambda120_(d_2492_i1_):
                        return _dafny.CodePoint('o')

                    init68_ = lambda120_
                    nw395_ = _dafny.Array(None, 12)
                    for i0_68_ in range(nw395_.length(0)):
                        nw395_[i0_68_] = init68_(i0_68_)
                    d_2491_v0_ = nw395_
                    d_2493_v1_: str
                    d_2493_v1_ = _dafny.CodePoint('h')
                    index412_ = default__.safeIndex(337, (d_2491_v0_).length(0))
                    (d_2491_v0_)[index412_] = d_2493_v1_
                    index413_ = default__.safeIndex(337, (d_2491_v0_).length(0))
                    (d_2491_v0_)[index413_] = d_2493_v1_
                    d_2494_v2_: _dafny.Seq
                    d_2494_v2_ = _dafny.SeqWithoutIsStrInference([not(p1), p1, p1])
                    d_2495_v3_: _dafny.Array
                    nw396_ = _dafny.Array(None, 11)
                    nw396_[int(0)] = p1
                    nw396_[int(1)] = p1
                    nw396_[int(2)] = (p1) or (p1)
                    nw396_[int(3)] = (d_2494_v2_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2493_v1_ for d_2496_i2_ in range(default__.abs(839))])), len(d_2494_v2_))]
                    nw396_[int(4)] = p1
                    nw396_[int(5)] = p1
                    nw396_[int(6)] = True
                    nw396_[int(7)] = (D2_DC9((0) - (self.f29), p1, (self).f26, self.f30)).cf17
                    nw396_[int(8)] = p1
                    nw396_[int(9)] = p1
                    nw396_[int(10)] = p1
                    d_2495_v3_ = nw396_
                    index414_ = default__.safeIndex(220, (d_2495_v3_).length(0))
                    (d_2495_v3_)[index414_] = p1
                    index415_ = default__.safeIndex(220, (d_2495_v3_).length(0))
                    (d_2495_v3_)[index415_] = p1
                    d_2497_v4_: _dafny.Array
                    nw397_ = _dafny.Array(_dafny.Array(None, 0), 27)
                    d_2497_v4_ = nw397_
                    index416_ = default__.safeIndex(527, (d_2497_v4_).length(0))
                    (d_2497_v4_)[index416_] = d_2495_v3_
                    index417_ = default__.safeIndex(527, (d_2497_v4_).length(0))
                    (d_2497_v4_)[index417_] = d_2495_v3_
                    pass
            pass
        d_2498_v5_: _dafny.Array
        nw398_ = _dafny.Array(False, 18)
        d_2498_v5_ = nw398_
        index418_ = default__.safeIndex(220, (d_2498_v5_).length(0))
        (d_2498_v5_)[index418_] = p1
        d_2499_v6_: _dafny.Seq
        d_2499_v6_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2500_v7_: _dafny.Set
        d_2500_v7_ = _dafny.Set({len(d_2499_v6_), (self).f26})
        d_2501_v8_: str
        d_2501_v8_ = _dafny.CodePoint('q')
        index419_ = default__.safeIndex(220, (d_2498_v5_).length(0))
        rhs428_ = p1
        rhs429_ = (((self).f27 if not (p1) or (not(p1)) else ((self).f27) + ((self).f27))).set(default__.safeIndex((832) + ((0) - ((0) - (len(d_2500_v7_)))), len(((self).f27 if not (p1) or (not(p1)) else ((self).f27) + ((self).f27)))), d_2501_v8_)
        rhs430_ = 566
        rhs431_ = p1
        lhs390_ = globalState
        lhs391_ = globalState
        lhs392_ = d_2498_v5_
        lhs393_ = default__.safeIndex(220, (d_2498_v5_).length(0))
        r1 = rhs428_
        lhs390_.f23 = rhs429_
        lhs391_.f17 = rhs430_
        lhs392_[lhs393_] = rhs431_
        index420_ = default__.safeIndex(616, (p0).length(0))
        (p0)[index420_] = (self).f26
        index421_ = default__.safeIndex(616, (p0).length(0))
        (p0)[index421_] = (self).f26
        d_2502_v9_: _dafny.Map
        d_2502_v9_ = _dafny.Map({self.f29: _dafny.Map({(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]: (self).f26})})
        d_2503_v10_: _dafny.Map
        d_2503_v10_ = _dafny.Map({(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]: (self).f26})
        d_2504_v11_: _dafny.Seq
        d_2504_v11_ = _dafny.SeqWithoutIsStrInference([self.f30, len(((d_2502_v9_)[544] if (544) in (d_2502_v9_) else d_2503_v10_))])
        index422_ = default__.safeIndex(616, (p0).length(0))
        (p0)[index422_] = len((d_2504_v11_ if (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))] else _dafny.SeqWithoutIsStrInference([(0) - ((default__.fm10(globalState)).cf11), self.f29, len(d_2504_v11_)])))
        if p1:
            d_2505_v12_: _dafny.MultiSet
            d_2505_v12_ = _dafny.MultiSet([(self).f26, (self).f26])
            d_2506_v13_: _dafny.Set
            d_2506_v13_ = _dafny.Set({(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], p1, False, p1})
            d_2507_v14_: _dafny.Map
            d_2507_v14_ = _dafny.Map({(self).fm8((d_2505_v12_).set(len(d_2504_v11_), default__.abs((self).f26)), p1, -248, (p0)[default__.safeIndex(616, (p0).length(0))], globalState): not((self).fm8(_dafny.MultiSet(d_2504_v11_), not((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]), len(d_2506_v13_), (p0)[default__.safeIndex(616, (p0).length(0))], globalState))})
            d_2508_v15_: _dafny.Set
            d_2508_v15_ = _dafny.Set({d_2507_v14_, (D3_DC10(d_2507_v14_)).cf20})
            index423_ = default__.safeIndex(220, (d_2498_v5_).length(0))
            (d_2498_v5_)[index423_] = (default__.fm11((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], globalState)) != (d_2508_v15_)
            d_2509_v16_: _dafny.Map
            d_2509_v16_ = _dafny.Map({self.f29: False})
            d_2509_v16_ = _dafny.Map({(self).fm7((self).f26, not(True), globalState): (d_2499_v6_)[default__.safeIndex(self.f30, len(d_2499_v6_))]})
            d_2510_v17_: C3
            nw399_ = C3()
            nw399_.ctor__((self).f26, (self).f26, _dafny.SeqWithoutIsStrInference([d_2501_v8_ for d_2511_i3_ in range(default__.abs(118))]))
            d_2510_v17_ = nw399_
            (globalState).f9 = (_dafny.SeqWithoutIsStrInference([not((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))])])) <= (_dafny.SeqWithoutIsStrInference([(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]]))
            d_2512_v18_: _dafny.Seq
            d_2512_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]: p1})])
            d_2513_v20_: _dafny.Map
            def iife198_():
                coll114_ = _dafny.Set()
                compr_114_: _dafny.Map
                for compr_114_ in (d_2512_v18_).Elements:
                    d_2514_v19_: _dafny.Map = compr_114_
                    if (d_2514_v19_) in (d_2512_v18_):
                        coll114_ = coll114_.union(_dafny.Set([d_2514_v19_]))
                return _dafny.Set(coll114_)
            d_2513_v20_ = _dafny.Map({(self).f27: (d_2508_v15_) - (iife198_()
            )})
            d_2515_v21_: _dafny.Map
            d_2515_v21_ = _dafny.Map({d_2506_v13_: p1})
            d_2513_v20_ = (d_2513_v20_).set(default__.fm43((d_2499_v6_)[default__.safeIndex(len(d_2515_v21_), len(d_2499_v6_))], (0) - ((self).f26), globalState), _dafny.Set({d_2507_v14_}))
        elif True:
            d_2516_v22_: _dafny.Array
            nw400_ = _dafny.Array(_dafny.Map({}), 7)
            d_2516_v22_ = nw400_
            d_2517_v23_: _dafny.Map
            d_2517_v23_ = _dafny.Map({d_2498_v5_: d_2503_v10_})
            index424_ = default__.safeIndex(763, (d_2516_v22_).length(0))
            (d_2516_v22_)[index424_] = d_2517_v23_
            d_2518_v24_: _dafny.Map
            d_2518_v24_ = _dafny.Map({((0) - ((p0)[default__.safeIndex(616, (p0).length(0))])) - (len((D18_DC50(self.f30, (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], d_2504_v11_, d_2501_v8_, (self).f26)).cf95)): ((self).f27).set(default__.safeIndex((self).f26, len((self).f27)), d_2501_v8_)})
            index425_ = default__.safeIndex(763, (d_2516_v22_).length(0))
            rhs432_ = ((_dafny.Map({d_2498_v5_: d_2503_v10_})) | ((d_2517_v23_).set(d_2498_v5_, d_2503_v10_))) | (_dafny.Map({d_2498_v5_: d_2503_v10_}))
            rhs433_ = ((d_2518_v24_)[(p0)[default__.safeIndex(616, (p0).length(0))]] if ((p0)[default__.safeIndex(616, (p0).length(0))]) in (d_2518_v24_) else ((self).f27) + (((self).f27).set(default__.safeIndex((self).f26, len((self).f27)), d_2501_v8_)))
            lhs394_ = d_2516_v22_
            lhs395_ = default__.safeIndex(763, (d_2516_v22_).length(0))
            lhs396_ = globalState
            lhs394_[lhs395_] = rhs432_
            lhs396_.f23 = rhs433_
            d_2519_v25_: _dafny.Array
            nw401_ = _dafny.Array(_dafny.CodePoint('D'), 16)
            d_2519_v25_ = nw401_
            d_2520_v26_: _dafny.Map
            d_2520_v26_ = _dafny.Map({(p0)[default__.safeIndex(616, (p0).length(0))]: d_2519_v25_})
            d_2521_v27_: D20
            d_2521_v27_ = D20_DC54(((d_2520_v26_)[self.f29] if (self.f29) in (d_2520_v26_) else d_2519_v25_))
            d_2522_v28_: _dafny.Map
            d_2522_v28_ = _dafny.Map({p1: d_2519_v25_})
            pat_let_tv56_ = d_2522_v28_
            pat_let_tv57_ = d_2498_v5_
            pat_let_tv58_ = d_2498_v5_
            pat_let_tv59_ = d_2498_v5_
            pat_let_tv60_ = d_2498_v5_
            pat_let_tv61_ = d_2522_v28_
            pat_let_tv62_ = d_2519_v25_
            def iife199_(_pat_let42_0):
                def iife200_(d_2523_dt__update__tmp_h0_):
                    def iife201_(_pat_let43_0):
                        def iife202_(d_2524_dt__update_hcf102_h0_):
                            return D20_DC54(d_2524_dt__update_hcf102_h0_)
                        return iife202_(_pat_let43_0)
                    return iife201_(((pat_let_tv56_)[(pat_let_tv58_)[default__.safeIndex(220, (pat_let_tv57_).length(0))]] if ((pat_let_tv60_)[default__.safeIndex(220, (pat_let_tv59_).length(0))]) in (pat_let_tv61_) else pat_let_tv62_))
                return iife200_(_pat_let42_0)
            source33_ = iife199_(d_2521_v27_)
            if source33_.is_DC55:
                d_2525___mcc_h0_ = source33_.cf103
                d_2526___mcc_h1_ = source33_.cf104
                d_2527___mcc_h2_ = source33_.cf105
                d_2528___mcc_h3_ = source33_.cf106
                d_2529_cf106_ = d_2528___mcc_h3_
                d_2530_cf105_ = d_2527___mcc_h2_
                d_2531_cf104_ = d_2526___mcc_h1_
                d_2532_cf103_ = d_2525___mcc_h0_
                (globalState).f9 = (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]
                d_2533_v29_: _dafny.Array
                nw402_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_2533_v29_ = nw402_
                d_2533_v29_ = d_2533_v29_
                d_2534_v30_: _dafny.Set
                d_2534_v30_ = _dafny.Set({(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], d_2531_cf104_, d_2529_cf106_})
                d_2535_v31_: D1
                d_2535_v31_ = D1_DC5(d_2534_v30_)
                d_2536_v32_: _dafny.Map
                d_2536_v32_ = _dafny.Map({(p0)[default__.safeIndex(616, (p0).length(0))]: d_2535_v31_})
                d_2537_v33_: _dafny.Map
                d_2537_v33_ = d_2536_v32_
                d_2538_v34_: D8
                d_2538_v34_ = D8_DC25(d_2529_cf106_, d_2503_v10_, _dafny.CodePoint('c'), d_2537_v33_, d_2530_cf105_)
                (globalState).f25 = ((0) - (self.f30)) == ((d_2538_v34_).cf53)
                d_2539_v35_: _dafny.Map
                d_2539_v35_ = _dafny.Map({d_2530_cf105_: (d_2499_v6_)[default__.safeIndex((self).f26, len(d_2499_v6_))]})
                index426_ = default__.safeIndex(616, (p0).length(0))
                rhs434_ = (len(d_2539_v35_)) * (d_2532_cf103_)
                rhs435_ = d_2532_cf103_
                lhs397_ = globalState
                lhs398_ = p0
                lhs399_ = default__.safeIndex(616, (p0).length(0))
                lhs397_.f8 = rhs434_
                lhs398_[lhs399_] = rhs435_
            elif True:
                d_2540___mcc_h4_ = source33_.cf102
                d_2541_cf102_ = d_2540___mcc_h4_
                (globalState).f13 = (self).f27
                (globalState).f0 = default__.safeDivisionInt(self.f29, self.f30)
                d_2542_v36_: _dafny.Set
                d_2542_v36_ = _dafny.Set({(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]})
                d_2543_v38_: _dafny.Map
                d_2543_v38_ = _dafny.Map({self.f30: (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]})
                d_2544_v39_: _dafny.Map
                d_2544_v39_ = _dafny.Map({not((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]): d_2543_v38_})
                d_2545_v41_: C4
                nw403_ = C4()
                def iife203_():
                    def iife204_():
                        coll116_ = _dafny.Set()
                        compr_116_: int
                        for compr_116_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (self.f30): self.f30})), (D10_DC32((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], (_dafny.MultiSet([len(_dafny.Map({p1: p1}))])).cardinality, False)).cf63])).Elements:
                            d_2547_v40_: int = compr_116_
                            if (d_2547_v40_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (self.f30): self.f30})), (D10_DC32((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], (_dafny.MultiSet([len(_dafny.Map({p1: p1}))])).cardinality, False)).cf63])):
                                coll116_ = coll116_.union(_dafny.Set([default__.safeModuloInt(d_2547_v40_, -362)]))
                        return _dafny.Set(coll116_)
                    coll115_ = _dafny.Map()
                    compr_115_: int
                    for compr_115_ in (((d_2544_v39_)[(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]] if ((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]) in (d_2544_v39_) else _dafny.Map({896: (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]}))).keys.Elements:
                        d_2546_v37_: int = compr_115_
                        if (d_2546_v37_) in (((d_2544_v39_)[(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]] if ((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]) in (d_2544_v39_) else _dafny.Map({896: (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]}))):
                            coll115_[default__.safeModuloInt(d_2546_v37_, len(iife204_()
                            ))] = (p0)[default__.safeIndex(616, (p0).length(0))]
                    return _dafny.Map(coll115_)
                nw403_.ctor__((d_2499_v6_)[default__.safeIndex(len((d_2518_v24_).set(len(default__.fm0((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], (self).f27, d_2542_v36_, self.f29, globalState)), (self).f27)), len(d_2499_v6_))], (p0)[default__.safeIndex(616, (p0).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enn")), iife203_()
                )
                d_2545_v41_ = nw403_
                index427_ = default__.safeIndex(616, (p0).length(0))
                (p0)[index427_] = (self.f29) * (((_dafny.MultiSet([True])).cardinality) - (len(((d_2503_v10_).set((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], self.f29)).set((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rol")))))))
            index428_ = default__.safeIndex(616, (p0).length(0))
            (p0)[index428_] = len(_dafny.SeqWithoutIsStrInference([self.f30 for d_2548_i4_ in range(default__.abs(983))]))
            d_2549_v42_: _dafny.Map
            d_2549_v42_ = _dafny.Map({(self).f27: p0})
            d_2550_v43_: _dafny.MultiSet
            d_2550_v43_ = _dafny.MultiSet([p0])
            d_2551_v44_: D8
            d_2551_v44_ = D8_DC23(p1, (d_2549_v42_) | (d_2549_v42_), d_2550_v43_, not(((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))] if False else p1)))
            pat_let_tv63_ = p0
            def iife205_(_pat_let44_0):
                def iife206_(d_2552_dt__update__tmp_h1_):
                    def iife207_(_pat_let45_0):
                        def iife208_(d_2553_dt__update_hcf44_h0_):
                            return D8_DC23((d_2552_dt__update__tmp_h1_).cf42, (d_2552_dt__update__tmp_h1_).cf43, d_2553_dt__update_hcf44_h0_, (d_2552_dt__update__tmp_h1_).cf45)
                        return iife208_(_pat_let45_0)
                    return iife207_(_dafny.MultiSet([pat_let_tv63_]))
                return iife206_(_pat_let44_0)
            d_2551_v44_ = iife205_(d_2551_v44_)
            d_2554_v45_: D16
            d_2554_v45_ = D16_DC45(self.f30, (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], (p0)[default__.safeIndex(616, (p0).length(0))], len((self).f27), False)
            d_2555_v46_: D16
            d_2555_v46_ = D16_DC46(d_2554_v45_)
            source34_ = d_2555_v46_
            if source34_.is_DC45:
                d_2556___mcc_h5_ = source34_.cf82
                d_2557___mcc_h6_ = source34_.cf83
                d_2558___mcc_h7_ = source34_.cf84
                d_2559___mcc_h8_ = source34_.cf85
                d_2560___mcc_h9_ = source34_.cf86
                d_2561_cf86_ = d_2560___mcc_h9_
                d_2562_cf85_ = d_2559___mcc_h8_
                d_2563_cf84_ = d_2558___mcc_h7_
                d_2564_cf83_ = d_2557___mcc_h6_
                d_2565_cf82_ = d_2556___mcc_h5_
                index429_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                (d_2498_v5_)[index429_] = d_2564_cf83_
                d_2566_v47_: _dafny.Set
                d_2566_v47_ = _dafny.Set({d_2499_v6_})
                d_2567_v48_: _dafny.Seq
                d_2567_v48_ = _dafny.SeqWithoutIsStrInference([d_2499_v6_, d_2499_v6_])
                def iife209_():
                    coll117_ = _dafny.Set()
                    compr_117_: _dafny.Seq
                    for compr_117_ in (d_2567_v48_).Elements:
                        d_2568_v49_: _dafny.Seq = compr_117_
                        if (d_2568_v49_) in (d_2567_v48_):
                            coll117_ = coll117_.union(_dafny.Set([d_2568_v49_]))
                    return _dafny.Set(coll117_)
                (globalState).f25 = ((d_2566_v47_) | (iife209_()
                )).issubset(d_2566_v47_)
                (globalState).f24 = len((d_2500_v7_).intersection((d_2500_v7_) | (d_2500_v7_)))
                (globalState).f8 = default__.safeDivisionInt(d_2563_cf84_, (0) - ((0) - (self.f30)))
            elif source34_.is_DC44:
                d_2569___mcc_h10_ = source34_.cf81
                d_2570_cf81_ = d_2569___mcc_h10_
                d_2571_v50_: D10
                d_2571_v50_ = D10_DC31(p1, _dafny.SeqWithoutIsStrInference([len((self).f27)]), (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], self.f29)
                d_2572_v51_: _dafny.MultiSet
                d_2572_v51_ = _dafny.MultiSet([len((self).f27)])
                d_2573_v52_: _dafny.Map
                d_2573_v52_ = _dafny.Map({(d_2499_v6_).set(default__.safeIndex((p0)[default__.safeIndex(616, (p0).length(0))], len(d_2499_v6_)), p1): (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]})
                index430_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                index431_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                rhs436_ = p1
                rhs437_ = (self.f30) != ((0) - (self.f30))
                rhs438_ = (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]
                rhs439_ = len((d_2499_v6_) + (((_dafny.SeqWithoutIsStrInference([p1, (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]])).set(default__.safeIndex((d_2504_v11_)[default__.safeIndex(self.f30, len(d_2504_v11_))], len(_dafny.SeqWithoutIsStrInference([p1, (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]]))), (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]) if p1 else d_2499_v6_)))
                rhs440_ = ((d_2503_v10_)[(d_2571_v50_).cf60] if ((d_2571_v50_).cf60) in (d_2503_v10_) else (0) - (((d_2572_v51_)[self.f29] if (self.f29) in (d_2572_v51_) else (self).fm6(_dafny.Map({(self).f26: (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]}), 182, d_2504_v11_, d_2573_v52_, globalState))))
                lhs400_ = d_2498_v5_
                lhs401_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                lhs402_ = d_2498_v5_
                lhs403_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                lhs404_ = globalState
                lhs405_ = globalState
                lhs406_ = globalState
                lhs400_[lhs401_] = rhs436_
                lhs402_[lhs403_] = rhs437_
                lhs404_.f9 = rhs438_
                lhs405_.f8 = rhs439_
                lhs406_.f0 = rhs440_
                d_2574_v53_: _dafny.Array
                nw404_ = _dafny.Array(None, 15)
                nw404_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbh"))
                nw404_[int(1)] = ((self).f27) + (_dafny.SeqWithoutIsStrInference([d_2501_v8_ for d_2575_i5_ in range(default__.abs(138))]))
                nw404_[int(2)] = (self).f27
                nw404_[int(3)] = (self).f27
                nw404_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhed"))
                nw404_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wklq"))) + ((self).f27)
                nw404_[int(6)] = (self).f27
                nw404_[int(7)] = ((self).f27) + ((self).f27)
                nw404_[int(8)] = (self).f27
                nw404_[int(9)] = (self).f27
                nw404_[int(10)] = ((self).f27) + ((self).f27)
                nw404_[int(11)] = (self).f27
                nw404_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                nw404_[int(13)] = ((self).f27) + (((self).f27).set(default__.safeIndex((p0)[default__.safeIndex(616, (p0).length(0))], len((self).f27)), _dafny.CodePoint('n')))
                nw404_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjsu"))
                d_2574_v53_ = nw404_
                index432_ = default__.safeIndex(172, (d_2574_v53_).length(0))
                (d_2574_v53_)[index432_] = (self).f27
                d_2576_v54_: _dafny.Array
                nw405_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_2576_v54_ = nw405_
                index433_ = default__.safeIndex(138, (d_2576_v54_).length(0))
                (d_2576_v54_)[index433_] = d_2498_v5_
                d_2577_v55_: _dafny.Map
                d_2577_v55_ = _dafny.Map({(self).fm8(_dafny.MultiSet([(p0)[default__.safeIndex(616, (p0).length(0))], 796]), (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], len(d_2500_v7_), (p0)[default__.safeIndex(616, (p0).length(0))], globalState): d_2498_v5_})
                d_2578_v56_: _dafny.Set
                d_2578_v56_ = _dafny.Set({d_2501_v8_})
                index434_ = default__.safeIndex(172, (d_2574_v53_).length(0))
                index435_ = default__.safeIndex(138, (d_2576_v54_).length(0))
                index436_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                rhs441_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqkfolo"))
                rhs442_ = ((d_2577_v55_)[(d_2501_v8_) not in (_dafny.SeqWithoutIsStrInference([d_2501_v8_, d_2501_v8_]))] if ((d_2501_v8_) not in (_dafny.SeqWithoutIsStrInference([d_2501_v8_, d_2501_v8_]))) in (d_2577_v55_) else d_2498_v5_)
                rhs443_ = p1
                rhs444_ = (p1 if (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))] else (d_2578_v56_).issubset(d_2578_v56_))
                rhs445_ = (d_2576_v54_ if True else d_2576_v54_)
                lhs407_ = d_2574_v53_
                lhs408_ = default__.safeIndex(172, (d_2574_v53_).length(0))
                lhs409_ = d_2576_v54_
                lhs410_ = default__.safeIndex(138, (d_2576_v54_).length(0))
                lhs411_ = d_2498_v5_
                lhs412_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                lhs407_[lhs408_] = rhs441_
                lhs409_[lhs410_] = rhs442_
                r1 = rhs443_
                lhs411_[lhs412_] = rhs444_
                d_2576_v54_ = rhs445_
                d_2579_v57_: _dafny.Set
                d_2579_v57_ = _dafny.Set({p1})
                d_2580_v58_: D16
                d_2580_v58_ = D16_DC45(len(d_2504_v11_), p1, self.f30, ((p0)[default__.safeIndex(616, (p0).length(0))]) - ((0) - ((p0)[default__.safeIndex(616, (p0).length(0))])), (default__.fm0(p1, (self).f27, d_2579_v57_, self.f29, globalState)) == (d_2499_v6_))
                d_2581_v59_: _dafny.Array
                nw406_ = _dafny.Array(_dafny.Seq({}), 6)
                d_2581_v59_ = nw406_
                index437_ = default__.safeIndex(707, (d_2581_v59_).length(0))
                (d_2581_v59_)[index437_] = d_2504_v11_
                d_2582_v61_: T2
                nw407_ = C2()
                def iife210_():
                    coll118_ = _dafny.Map()
                    compr_118_: int
                    for compr_118_ in (d_2504_v11_).Elements:
                        d_2583_v60_: int = compr_118_
                        if (d_2583_v60_) in (d_2504_v11_):
                            coll118_[(d_2583_v60_) + ((self).f26)] = (self).f26
                    return _dafny.Map(coll118_)
                nw407_.ctor__(d_2504_v11_, iife210_()
                )
                d_2582_v61_ = nw407_
                d_2584_v62_: _dafny.Seq
                d_2584_v62_ = _dafny.SeqWithoutIsStrInference([d_2582_v61_, d_2582_v61_])
                d_2585_v63_: _dafny.Map
                d_2585_v63_ = _dafny.Map({(d_2584_v62_)[default__.safeIndex(self.f29, len(d_2584_v62_))]: False})
                d_2586_v64_: D17
                d_2586_v64_ = D17_DC48(d_2582_v61_, (self).f26, d_2501_v8_)
                index438_ = default__.safeIndex(707, (d_2581_v59_).length(0))
                rhs446_ = d_2580_v58_
                rhs447_ = d_2504_v11_
                rhs448_ = not ((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]) or (((d_2585_v63_)[(d_2586_v64_).cf89] if ((d_2586_v64_).cf89) in (d_2585_v63_) else p1))
                rhs449_ = (d_2576_v54_)[default__.safeIndex(138, (d_2576_v54_).length(0))]
                lhs413_ = d_2581_v59_
                lhs414_ = default__.safeIndex(707, (d_2581_v59_).length(0))
                lhs415_ = globalState
                d_2580_v58_ = rhs446_
                lhs413_[lhs414_] = rhs447_
                lhs415_.f9 = rhs448_
                d_2498_v5_ = rhs449_
                (globalState).f5 = (self).f26
            elif True:
                d_2587___mcc_h11_ = source34_.cf87
                d_2588_cf87_ = d_2587___mcc_h11_
                index439_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                (d_2498_v5_)[index439_] = (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]
                d_2589_v65_: D7
                d_2589_v65_ = D7_DC19(-77, (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))])
                d_2590_v66_: D7
                d_2590_v66_ = D7_DC21(d_2589_v65_)
                d_2591_v67_: _dafny.Seq
                d_2591_v67_ = _dafny.SeqWithoutIsStrInference([d_2590_v66_])
                d_2592_v68_: _dafny.Set
                d_2592_v68_ = _dafny.Set({d_2591_v67_})
                d_2593_v69_: D16
                d_2593_v69_ = D16_DC44(d_2592_v68_)
                d_2594_v70_: _dafny.Map
                d_2594_v70_ = _dafny.Map({d_2504_v11_: (0) - (self.f29)})
                d_2595_v71_: _dafny.Map
                d_2595_v71_ = _dafny.Map({d_2593_v69_: (d_2594_v70_) | (d_2594_v70_)})
                d_2595_v71_ = (d_2595_v71_).set(d_2593_v69_, d_2594_v70_)
                d_2501_v8_ = d_2501_v8_
                d_2501_v8_ = d_2501_v8_
        d_2596_v72_: _dafny.Map
        d_2596_v72_ = _dafny.Map({(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]: (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]})
        if ((p1) in (d_2596_v72_)) and (not(p1)):
            d_2597_v73_: D18
            d_2597_v73_ = D18_DC50(self.f29, p1, d_2504_v11_, d_2501_v8_, self.f29)
            d_2598_v74_: _dafny.Array
            nw408_ = _dafny.Array(None, 24)
            nw408_[int(0)] = default__.fm43((d_2597_v73_).cf94, self.f29, globalState)
            nw408_[int(1)] = _dafny.SeqWithoutIsStrInference([d_2501_v8_ for d_2599_i6_ in range(default__.abs(520))])
            nw408_[int(2)] = _dafny.SeqWithoutIsStrInference([d_2501_v8_ for d_2600_i7_ in range(default__.abs(614))])
            nw408_[int(3)] = (self).f27
            nw408_[int(4)] = (self).f27
            nw408_[int(5)] = default__.fm43(p1, (p0)[default__.safeIndex(616, (p0).length(0))], globalState)
            nw408_[int(6)] = (self).f27
            nw408_[int(7)] = (self).f27
            nw408_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
            nw408_[int(9)] = (self).f27
            nw408_[int(10)] = (self).f27
            nw408_[int(11)] = (self).f27
            nw408_[int(12)] = (self).f27
            nw408_[int(13)] = (self).f27
            nw408_[int(14)] = (self).f27
            nw408_[int(15)] = (self).f27
            nw408_[int(16)] = (self).f27
            nw408_[int(17)] = (self).f27
            nw408_[int(18)] = _dafny.SeqWithoutIsStrInference([d_2501_v8_ for d_2601_i8_ in range(default__.abs(711))])
            nw408_[int(19)] = (self).f27
            nw408_[int(20)] = (self).f27
            nw408_[int(21)] = ((self).f27).set(default__.safeIndex((p0)[default__.safeIndex(616, (p0).length(0))], len((self).f27)), d_2501_v8_)
            nw408_[int(22)] = (self).f27
            nw408_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlbn"))
            d_2598_v74_ = nw408_
            d_2602_v75_: D13
            d_2602_v75_ = D13_DC40(d_2598_v74_, d_2498_v5_, (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))])
            d_2603_v76_: D6
            d_2603_v76_ = D6_DC16((self).f26, p1)
            index440_ = default__.safeIndex(616, (p0).length(0))
            rhs450_ = (self.f30) * ((p0)[default__.safeIndex(616, (p0).length(0))])
            rhs451_ = (_dafny.SeqWithoutIsStrInference([p1, ((d_2596_v72_)[p1] if (p1) in (d_2596_v72_) else True)])) == (d_2499_v6_)
            rhs452_ = (0) - (default__.safeModuloInt((p0)[default__.safeIndex(616, (p0).length(0))], (self).fm6((_dafny.Map({(self).f26: not(p1)})).set(len(d_2500_v7_), (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]), self.f29, d_2504_v11_, default__.fm28(globalState), globalState)))
            rhs453_ = (d_2602_v75_).cf77
            rhs454_ = (d_2603_v76_).cf31
            lhs416_ = p0
            lhs417_ = default__.safeIndex(616, (p0).length(0))
            lhs418_ = globalState
            lhs416_[lhs417_] = rhs450_
            r1 = rhs451_
            r0 = rhs452_
            d_2498_v5_ = rhs453_
            lhs418_.f24 = rhs454_
            (globalState).f25 = p1
            (globalState).f8 = (p0)[default__.safeIndex(616, (p0).length(0))]
            d_2604_v77_: C0
            nw409_ = C0()
            nw409_.ctor__()
            d_2604_v77_ = nw409_
            d_2605_v78_: _dafny.MultiSet
            d_2605_v78_ = _dafny.MultiSet([d_2604_v77_])
            d_2606_v79_: _dafny.MultiSet
            d_2606_v79_ = _dafny.MultiSet([self.f30, len((self).f27)])
            index441_ = default__.safeIndex(220, (d_2498_v5_).length(0))
            rhs455_ = not((d_2605_v78_).ispropersubset(d_2605_v78_))
            rhs456_ = (self).fm8(((d_2606_v79_).set(self.f29, default__.abs((p0)[default__.safeIndex(616, (p0).length(0))]))) | (d_2606_v79_), (len(d_2499_v6_)) != ((self).f26), ((d_2606_v79_)[(self).f26] if ((self).f26) in (d_2606_v79_) else len((self).f27)), self.f29, globalState)
            lhs419_ = d_2498_v5_
            lhs420_ = default__.safeIndex(220, (d_2498_v5_).length(0))
            lhs421_ = globalState
            lhs419_[lhs420_] = rhs455_
            lhs421_.f25 = rhs456_
            (self).f29 = ((self).f26) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdy"))))
        elif True:
            index442_ = default__.safeIndex(616, (p0).length(0))
            (p0)[index442_] = (0) - (-389)
            (globalState).f13 = (self).f27
            d_2607_v80_: _dafny.MultiSet
            d_2607_v80_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([self.f29 for d_2608_i9_ in range(default__.abs(135))])), (p0)[default__.safeIndex(616, (p0).length(0))], (p0)[default__.safeIndex(616, (p0).length(0))]])
            if (d_2499_v6_)[default__.safeIndex(((p0)[default__.safeIndex(616, (p0).length(0))] if (self).fm8(d_2607_v80_, not(p1), self.f29, self.f30, globalState) else self.f29), len(d_2499_v6_))]:
                d_2609_v81_: D10
                d_2609_v81_ = D10_DC31((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], d_2504_v11_, (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], (self).f26)
                (globalState).f5 = (d_2609_v81_).cf61
                d_2610_v82_: _dafny.Array
                nw410_ = _dafny.Array(None, 11)
                nw410_[int(0)] = d_2499_v6_
                nw410_[int(1)] = d_2499_v6_
                nw410_[int(2)] = d_2499_v6_
                nw410_[int(3)] = d_2499_v6_
                nw410_[int(4)] = d_2499_v6_
                nw410_[int(5)] = d_2499_v6_
                nw410_[int(6)] = d_2499_v6_
                nw410_[int(7)] = d_2499_v6_
                nw410_[int(8)] = d_2499_v6_
                nw410_[int(9)] = (_dafny.SeqWithoutIsStrInference([not(p1)])).set(default__.safeIndex(self.f29, len(_dafny.SeqWithoutIsStrInference([not(p1)]))), (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))])
                nw410_[int(10)] = d_2499_v6_
                d_2610_v82_ = nw410_
                d_2611_v83_: _dafny.Map
                d_2611_v83_ = _dafny.Map({d_2596_v72_: p1})
                d_2612_v84_: _dafny.Map
                d_2612_v84_ = _dafny.Map({d_2610_v82_: d_2611_v83_})
                def iife211_():
                    coll119_ = _dafny.Map()
                    compr_119_: _dafny.Map
                    for compr_119_ in (d_2611_v83_).keys.Elements:
                        d_2613_v85_: _dafny.Map = compr_119_
                        if (d_2613_v85_) in (d_2611_v83_):
                            coll119_[d_2613_v85_] = p1
                    return _dafny.Map(coll119_)
                d_2612_v84_ = (d_2612_v84_).set(d_2610_v82_, iife211_()
                )
                d_2614_v86_: _dafny.MultiSet
                d_2614_v86_ = _dafny.MultiSet([(self).f27])
                (globalState).f25 = ((d_2614_v86_) - (d_2614_v86_)).isdisjoint(d_2614_v86_)
                index443_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                (d_2498_v5_)[index443_] = not((((p0)[default__.safeIndex(616, (p0).length(0))]) + (self.f30)) <= (self.f30))
                d_2615_v87_: D7
                d_2615_v87_ = D7_DC18((p0)[default__.safeIndex(616, (p0).length(0))], p0)
                d_2616_v88_: D7
                d_2616_v88_ = D7_DC21(d_2615_v87_)
                pat_let_tv64_ = d_2615_v87_
                d_2617_v89_: _dafny.Seq
                def iife212_(_pat_let46_0):
                    def iife213_(d_2618_dt__update__tmp_h2_):
                        def iife214_(_pat_let47_0):
                            def iife215_(d_2619_dt__update_hcf40_h0_):
                                return D7_DC21(d_2619_dt__update_hcf40_h0_)
                            return iife215_(_pat_let47_0)
                        return iife214_(pat_let_tv64_)
                    return iife213_(_pat_let46_0)
                d_2617_v89_ = _dafny.SeqWithoutIsStrInference([iife212_(d_2616_v88_)])
                d_2620_v90_: _dafny.Seq
                d_2620_v90_ = _dafny.SeqWithoutIsStrInference([d_2617_v89_])
                d_2621_v91_: D9
                d_2621_v91_ = D9_DC27(d_2620_v90_)
                d_2622_v92_: _dafny.Map
                d_2622_v92_ = _dafny.Map({p0: d_2621_v91_})
                index444_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                rhs457_ = (d_2501_v8_) in ((self).f27)
                rhs458_ = (0) - ((0) - (len(((d_2622_v92_) | (d_2622_v92_)) | (d_2622_v92_))))
                lhs422_ = d_2498_v5_
                lhs423_ = default__.safeIndex(220, (d_2498_v5_).length(0))
                lhs424_ = globalState
                lhs422_[lhs423_] = rhs457_
                lhs424_.f0 = rhs458_
            elif True:
                d_2623_v94_: _dafny.Map
                d_2623_v94_ = _dafny.Map({self.f29: (p0)[default__.safeIndex(616, (p0).length(0))]})
                def iife216_():
                    coll120_ = _dafny.Set()
                    compr_120_: int
                    for compr_120_ in _dafny.IntegerRange(399, 311):
                        d_2624_v93_: int = compr_120_
                        if ((399) <= (d_2624_v93_)) and ((d_2624_v93_) < (311)):
                            coll120_ = coll120_.union(_dafny.Set([default__.safeModuloInt(d_2624_v93_, self.f30)]))
                    return _dafny.Set(coll120_)
                (globalState).f8 = ((len(iife216_()
                ) if p1 else ((d_2623_v94_)[632] if (632) in (d_2623_v94_) else (self).f26))) - (459)
                d_2625_v95_: _dafny.Set
                d_2625_v95_ = _dafny.Set({True})
                d_2626_v96_: D1
                d_2626_v96_ = D1_DC5(d_2625_v95_)
                pat_let_tv65_ = d_2625_v95_
                d_2627_v97_: _dafny.Map
                def iife217_(_pat_let48_0):
                    def iife218_(d_2628_dt__update__tmp_h3_):
                        def iife219_(_pat_let49_0):
                            def iife220_(d_2629_dt__update_hcf8_h0_):
                                return D1_DC5(d_2629_dt__update_hcf8_h0_)
                            return iife220_(_pat_let49_0)
                        return iife219_(pat_let_tv65_)
                    return iife218_(_pat_let48_0)
                d_2627_v97_ = _dafny.Map({(_dafny.MultiSet([(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]])).cardinality: iife217_(d_2626_v96_)})
                d_2630_v98_: _dafny.Map
                d_2630_v98_ = d_2627_v97_
                d_2631_v99_: _dafny.Seq
                d_2631_v99_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_2632_v100_: _dafny.Seq
                d_2632_v100_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))]])])
                d_2633_v101_: _dafny.Map
                d_2633_v101_ = _dafny.Map({d_2498_v5_: len(d_2632_v100_)})
                pat_let_tv66_ = d_2625_v95_
                d_2634_v102_: _dafny.Array
                nw411_ = _dafny.Array(None, 25)
                nw411_[int(0)] = d_2630_v98_
                nw411_[int(1)] = d_2630_v98_
                nw411_[int(2)] = d_2630_v98_
                nw411_[int(3)] = d_2630_v98_
                nw411_[int(4)] = d_2630_v98_
                nw411_[int(5)] = d_2630_v98_
                nw411_[int(6)] = _dafny.Map({(self).f26: d_2626_v96_})
                nw411_[int(7)] = d_2630_v98_
                nw411_[int(8)] = d_2627_v97_
                nw411_[int(9)] = d_2627_v97_
                nw411_[int(10)] = d_2630_v98_
                nw411_[int(11)] = d_2630_v98_
                nw411_[int(12)] = (d_2630_v98_ if (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))] else d_2627_v97_)
                nw411_[int(13)] = _dafny.Map({51: d_2626_v96_})
                nw411_[int(14)] = default__.fm52(p1, p1, len((self).f27), (self).f26, globalState)
                nw411_[int(15)] = d_2627_v97_
                nw411_[int(16)] = d_2627_v97_
                nw411_[int(17)] = d_2630_v98_
                nw411_[int(18)] = d_2630_v98_
                nw411_[int(19)] = d_2627_v97_
                def iife221_(_pat_let50_0):
                    def iife222_(d_2635_dt__update__tmp_h7_):
                        def iife223_(_pat_let51_0):
                            def iife224_(d_2636_dt__update_hcf8_h1_):
                                return D1_DC5(d_2636_dt__update_hcf8_h1_)
                            return iife224_(_pat_let51_0)
                        return iife223_(pat_let_tv66_)
                    return iife222_(_pat_let50_0)
                nw411_[int(20)] = (_dafny.Map({(p0)[default__.safeIndex(616, (p0).length(0))]: iife221_(d_2626_v96_)})).set(len(d_2631_v99_), d_2626_v96_)
                nw411_[int(21)] = d_2630_v98_
                nw411_[int(22)] = d_2630_v98_
                nw411_[int(23)] = _dafny.Map({len(d_2633_v101_): d_2626_v96_})
                nw411_[int(24)] = d_2630_v98_
                d_2634_v102_ = nw411_
                index445_ = default__.safeIndex(339, (d_2634_v102_).length(0))
                (d_2634_v102_)[index445_] = d_2630_v98_
                d_2637_v103_: T0
                nw412_ = C1()
                nw412_.ctor__()
                d_2637_v103_ = nw412_
                d_2638_v104_: D10
                d_2638_v104_ = D10_DC32((d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))], len((self).f27), (d_2498_v5_)[default__.safeIndex(220, (d_2498_v5_).length(0))])
                index446_ = default__.safeIndex(616, (p0).length(0))
                index447_ = default__.safeIndex(339, (d_2634_v102_).length(0))
                rhs459_ = (d_2638_v104_).cf64
                rhs460_ = ((self).f26) + ((self.f30 if p1 else len(_dafny.SeqWithoutIsStrInference([False, p1]))))
                rhs461_ = d_2627_v97_
                rhs462_ = d_2637_v103_
                lhs425_ = globalState
                lhs426_ = p0
                lhs427_ = default__.safeIndex(616, (p0).length(0))
                lhs428_ = d_2634_v102_
                lhs429_ = default__.safeIndex(339, (d_2634_v102_).length(0))
                lhs425_.f9 = rhs459_
                lhs426_[lhs427_] = rhs460_
                lhs428_[lhs429_] = rhs461_
                d_2637_v103_ = rhs462_
                d_2501_v8_ = (d_2501_v8_ if not((_dafny.MultiSet(d_2504_v11_)).issubset(d_2607_v80_)) else d_2501_v8_)
                (globalState).f24 = ((self).f26) * ((self).f26)
                index448_ = default__.safeIndex(616, (p0).length(0))
                (p0)[index448_] = default__.safeModuloInt((self).f26, (self).f26)
            d_2639_v105_: _dafny.Array
            def lambda121_(d_2640_v7_, d_2641_p1_):
                def lambda122_(d_2642_i10_):
                    return (D8_DC22(_dafny.Map({len(d_2640_v7_): d_2641_p1_}))).cf41

                return lambda122_

            init69_ = lambda121_(d_2500_v7_, p1)
            nw413_ = _dafny.Array(None, 25)
            for i0_69_ in range(nw413_.length(0)):
                nw413_[i0_69_] = init69_(i0_69_)
            d_2639_v105_ = nw413_
            d_2639_v105_ = d_2639_v105_
            d_2643_v106_: C1
            nw414_ = C1()
            nw414_.ctor__()
            d_2643_v106_ = nw414_
        r0 = ((self).f26 if p1 else (p0)[default__.safeIndex(616, (p0).length(0))])
        r1 = not(p1)
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_2644_v0_: _dafny.Array
        nw415_ = _dafny.Array(False, 21)
        d_2644_v0_ = nw415_
        d_2645_v1_: D7
        d_2645_v1_ = D7_DC17(d_2644_v0_)
        d_2646_v2_: _dafny.MultiSet
        d_2646_v2_ = _dafny.MultiSet([d_2644_v0_, (d_2645_v1_).cf33, d_2644_v0_])
        d_2647_v3_: _dafny.Array
        def lambda123_(d_2648_p0_):
            def lambda124_(d_2649_i0_):
                return (_dafny.Map({(self).f27: 616}) if not(d_2648_p0_) else _dafny.Map({(self).f27: 233}))

            return lambda124_

        init70_ = lambda123_(p0)
        nw416_ = _dafny.Array(None, 28)
        for i0_70_ in range(nw416_.length(0)):
            nw416_[i0_70_] = init70_(i0_70_)
        d_2647_v3_ = nw416_
        d_2650_v4_: _dafny.Map
        d_2650_v4_ = _dafny.Map({(self).f27: self.f30})
        index449_ = default__.safeIndex(533, (d_2647_v3_).length(0))
        (d_2647_v3_)[index449_] = d_2650_v4_
        index450_ = default__.safeIndex(533, (d_2647_v3_).length(0))
        rhs463_ = _dafny.MultiSet([d_2644_v0_])
        rhs464_ = (d_2650_v4_) | (d_2650_v4_)
        lhs430_ = d_2647_v3_
        lhs431_ = default__.safeIndex(533, (d_2647_v3_).length(0))
        d_2646_v2_ = rhs463_
        lhs430_[lhs431_] = rhs464_
        if p0:
            d_2651_v5_: str
            d_2651_v5_ = _dafny.CodePoint('h')
            d_2652_v6_: _dafny.Map
            d_2652_v6_ = _dafny.Map({self.f30: True})
            d_2653_v7_: _dafny.Seq
            d_2653_v7_ = _dafny.SeqWithoutIsStrInference([self.f29, (self).f26])
            d_2654_v8_: _dafny.Seq
            d_2654_v8_ = _dafny.SeqWithoutIsStrInference([False, p0])
            d_2655_v9_: _dafny.Map
            d_2655_v9_ = _dafny.Map({d_2654_v8_: p0})
            d_2656_v10_: _dafny.Array
            nw417_ = _dafny.Array(None, 12)
            nw417_[int(0)] = _dafny.CodePoint('f')
            nw417_[int(1)] = d_2651_v5_
            nw417_[int(2)] = ((self).f27)[default__.safeIndex((self).fm6(d_2652_v6_, self.f29, d_2653_v7_, d_2655_v9_, globalState), len((self).f27))]
            nw417_[int(3)] = d_2651_v5_
            nw417_[int(4)] = d_2651_v5_
            nw417_[int(5)] = (_dafny.CodePoint('b') if p0 else _dafny.CodePoint('y'))
            nw417_[int(6)] = (d_2651_v5_ if p0 else d_2651_v5_)
            nw417_[int(7)] = d_2651_v5_
            nw417_[int(8)] = d_2651_v5_
            nw417_[int(9)] = _dafny.CodePoint('h')
            nw417_[int(10)] = d_2651_v5_
            nw417_[int(11)] = d_2651_v5_
            d_2656_v10_ = nw417_
            d_2656_v10_ = d_2656_v10_
            d_2653_v7_ = d_2653_v7_
            d_2657_v11_: C3
            nw418_ = C3()
            nw418_.ctor__((self).f26, self.f29, (self).f27)
            d_2657_v11_ = nw418_
            d_2658_v12_: _dafny.Map
            d_2658_v12_ = _dafny.Map({d_2657_v11_: not((d_2654_v8_)[default__.safeIndex(self.f29, len(d_2654_v8_))])})
            d_2658_v12_ = (d_2658_v12_).set(d_2657_v11_, p0)
            (globalState).f9 = p0
            d_2659_v13_: _dafny.Array
            nw419_ = _dafny.Array(_dafny.Map({}), 27)
            d_2659_v13_ = nw419_
            d_2660_v14_: _dafny.Map
            d_2660_v14_ = _dafny.Map({p0: (d_2659_v13_ if p0 else d_2659_v13_)})
            d_2660_v14_ = (d_2660_v14_).set(True, d_2659_v13_)
        elif True:
            if not(p0):
                index451_ = default__.safeIndex(475, (d_2644_v0_).length(0))
                (d_2644_v0_)[index451_] = p0
                index452_ = default__.safeIndex(475, (d_2644_v0_).length(0))
                (d_2644_v0_)[index452_] = not((self.f30) >= (-410))
                (globalState).f23 = (self).f27
                (globalState).f17 = self.f30
                d_2661_v15_: _dafny.Array
                nw420_ = _dafny.Array(None, 19)
                nw420_[int(0)] = (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]
                nw420_[int(1)] = (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]
                nw420_[int(2)] = (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]
                nw420_[int(3)] = (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]
                nw420_[int(4)] = p0
                nw420_[int(5)] = (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]
                nw420_[int(6)] = p0
                nw420_[int(7)] = p0
                nw420_[int(8)] = p0
                nw420_[int(9)] = p0
                nw420_[int(10)] = p0
                nw420_[int(11)] = p0
                nw420_[int(12)] = (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]
                nw420_[int(13)] = False
                nw420_[int(14)] = (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]
                nw420_[int(15)] = (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]
                nw420_[int(16)] = True
                nw420_[int(17)] = False
                nw420_[int(18)] = p0
                d_2661_v15_ = nw420_
                d_2662_v16_: D11
                d_2662_v16_ = D11_DC35(D11_DC34(d_2661_v15_, p0))
                d_2663_v17_: D11
                d_2663_v17_ = D11_DC35(d_2662_v16_)
                pat_let_tv67_ = d_2662_v16_
                def iife225_(_pat_let52_0):
                    def iife226_(d_2664_dt__update__tmp_h0_):
                        def iife227_(_pat_let53_0):
                            def iife228_(d_2665_dt__update_hcf68_h0_):
                                return D11_DC35(d_2665_dt__update_hcf68_h0_)
                            return iife228_(_pat_let53_0)
                        return iife227_(D11_DC35(pat_let_tv67_))
                    return iife226_(_pat_let52_0)
                d_2663_v17_ = iife225_(d_2663_v17_)
                d_2666_v18_: _dafny.Seq
                d_2666_v18_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_2667_v19_: _dafny.Map
                d_2667_v19_ = _dafny.Map({len(d_2666_v18_): (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]})
                d_2668_v20_: _dafny.Map
                d_2668_v20_ = _dafny.Map({((d_2667_v19_)[self.f30] if (self.f30) in (d_2667_v19_) else p0): d_2661_v15_})
                rhs465_ = d_2668_v20_
                rhs466_ = (0) - ((self).f26)
                rhs467_ = (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]
                rhs468_ = (d_2644_v0_)[default__.safeIndex(475, (d_2644_v0_).length(0))]
                rhs469_ = (0) - (len((self).f27))
                lhs432_ = globalState
                lhs433_ = globalState
                lhs434_ = globalState
                lhs435_ = globalState
                d_2668_v20_ = rhs465_
                lhs432_.f8 = rhs466_
                lhs433_.f9 = rhs467_
                lhs434_.f25 = rhs468_
                lhs435_.f24 = rhs469_
            elif True:
                d_2669_v21_: _dafny.Map
                d_2669_v21_ = _dafny.Map({len(_dafny.Map({self.f29: p0})): self.f30})
                d_2670_v22_: _dafny.Map
                d_2670_v22_ = _dafny.Map({d_2669_v21_: p0})
                d_2671_v23_: _dafny.Map
                d_2671_v23_ = _dafny.Map({p0: (((d_2670_v22_)[d_2669_v21_] if (d_2669_v21_) in (d_2670_v22_) else p0) if p0 else p0)})
                d_2672_v24_: _dafny.MultiSet
                d_2672_v24_ = _dafny.MultiSet([self.f29])
                d_2671_v23_ = (d_2671_v23_).set((self).fm8(d_2672_v24_, p0, len((self).f27), (self).f26, globalState), p0)
                d_2673_v25_: _dafny.Array
                nw421_ = _dafny.Array(_dafny.Map({}), 20)
                d_2673_v25_ = nw421_
                d_2674_v27_: _dafny.Map
                d_2674_v27_ = _dafny.Map({(self).f26: p0})
                def iife229_():
                    coll121_ = _dafny.Map()
                    compr_121_: int
                    for compr_121_ in (d_2674_v27_).keys.Elements:
                        d_2675_v26_: int = compr_121_
                        if (d_2675_v26_) in (d_2674_v27_):
                            coll121_[(d_2675_v26_) + ((self).f26)] = p0
                    return _dafny.Map(coll121_)
                rhs470_ = (self).f26
                rhs471_ = p0
                rhs472_ = d_2673_v25_
                rhs473_ = self.f30
                rhs474_ = ((self).f26) >= (len(iife229_()
                ))
                lhs436_ = globalState
                lhs437_ = globalState
                lhs438_ = self
                lhs439_ = globalState
                lhs436_.f8 = rhs470_
                lhs437_.f25 = rhs471_
                d_2673_v25_ = rhs472_
                lhs438_.f29 = rhs473_
                lhs439_.f9 = rhs474_
                d_2676_v28_: str
                d_2676_v28_ = _dafny.CodePoint('f')
                d_2677_v29_: _dafny.Map
                d_2677_v29_ = _dafny.Map({(self).f27: d_2676_v28_})
                d_2678_v30_: _dafny.Set
                d_2678_v30_ = _dafny.Set({self.f29})
                d_2679_v31_: _dafny.Array
                nw422_ = _dafny.Array(None, 25)
                nw422_[int(0)] = d_2676_v28_
                nw422_[int(1)] = d_2676_v28_
                nw422_[int(2)] = d_2676_v28_
                nw422_[int(3)] = d_2676_v28_
                nw422_[int(4)] = d_2676_v28_
                nw422_[int(5)] = d_2676_v28_
                nw422_[int(6)] = d_2676_v28_
                nw422_[int(7)] = d_2676_v28_
                nw422_[int(8)] = d_2676_v28_
                nw422_[int(9)] = d_2676_v28_
                nw422_[int(10)] = d_2676_v28_
                nw422_[int(11)] = d_2676_v28_
                nw422_[int(12)] = _dafny.CodePoint('m')
                nw422_[int(13)] = d_2676_v28_
                nw422_[int(14)] = d_2676_v28_
                nw422_[int(15)] = d_2676_v28_
                nw422_[int(16)] = d_2676_v28_
                nw422_[int(17)] = d_2676_v28_
                nw422_[int(18)] = d_2676_v28_
                nw422_[int(19)] = d_2676_v28_
                nw422_[int(20)] = d_2676_v28_
                nw422_[int(21)] = ((d_2677_v29_)[(self).f27] if ((self).f27) in (d_2677_v29_) else d_2676_v28_)
                nw422_[int(22)] = default__.fm42(d_2678_v30_, globalState)
                nw422_[int(23)] = d_2676_v28_
                nw422_[int(24)] = d_2676_v28_
                d_2679_v31_ = nw422_
                d_2680_v32_: D20
                d_2680_v32_ = D20_DC54(d_2679_v31_)
                d_2680_v32_ = d_2680_v32_
                d_2681_v33_: C1
                nw423_ = C1()
                nw423_.ctor__()
                d_2681_v33_ = nw423_
                d_2681_v33_ = d_2681_v33_
                d_2682_v34_: _dafny.Map
                d_2682_v34_ = _dafny.Map({d_2644_v0_: (self).f27})
                d_2682_v34_ = (_dafny.Map({d_2644_v0_: (self).f27})) | (d_2682_v34_)
            if not(p0):
                (globalState).f25 = p0
                d_2683_v35_: _dafny.Array
                def lambda125_(d_2684_i1_):
                    return default__.safeModuloInt(d_2684_i1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "in"))))

                init71_ = lambda125_
                nw424_ = _dafny.Array(None, 18)
                for i0_71_ in range(nw424_.length(0)):
                    nw424_[i0_71_] = init71_(i0_71_)
                d_2683_v35_ = nw424_
                index453_ = default__.safeIndex(996, (d_2683_v35_).length(0))
                (d_2683_v35_)[index453_] = (129) * (self.f30)
                d_2685_v36_: _dafny.Set
                d_2685_v36_ = _dafny.Set({p0, p0})
                d_2686_v37_: _dafny.Seq
                d_2686_v37_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                index454_ = default__.safeIndex(996, (d_2683_v35_).length(0))
                rhs475_ = p0
                rhs476_ = (0) - (default__.safeDivisionInt(len(d_2685_v36_), -509))
                rhs477_ = ((d_2686_v37_)[default__.safeIndex((self).f26, len(d_2686_v37_))]) and ((self.f29) >= (self.f29))
                rhs478_ = self.f30
                lhs440_ = globalState
                lhs441_ = d_2683_v35_
                lhs442_ = default__.safeIndex(996, (d_2683_v35_).length(0))
                lhs443_ = globalState
                lhs444_ = globalState
                lhs440_.f25 = rhs475_
                lhs441_[lhs442_] = rhs476_
                lhs443_.f9 = rhs477_
                lhs444_.f0 = rhs478_
                d_2687_v38_: _dafny.Map
                d_2687_v38_ = _dafny.Map({len((self).f27): p0})
                d_2688_v40_: _dafny.Seq
                def iife230_():
                    coll122_ = _dafny.Map()
                    compr_122_: _dafny.Seq
                    for compr_122_ in (_dafny.SeqWithoutIsStrInference([d_2686_v37_, d_2686_v37_, d_2686_v37_])).Elements:
                        d_2689_v39_: _dafny.Seq = compr_122_
                        if (d_2689_v39_) in (_dafny.SeqWithoutIsStrInference([d_2686_v37_, d_2686_v37_, d_2686_v37_])):
                            coll122_[d_2689_v39_] = p0
                    return _dafny.Map(coll122_)
                d_2688_v40_ = _dafny.SeqWithoutIsStrInference([(self).fm6(d_2687_v38_, (0) - ((self).f26), _dafny.SeqWithoutIsStrInference([self.f29]), iife230_()
                , globalState), (self).f26, self.f30])
                d_2688_v40_ = d_2688_v40_
                d_2690_v41_: D7
                d_2690_v41_ = D7_DC19((self).fm7(len((self).f27), p0, globalState), default__.fm24(p0, p0, p0, len(d_2688_v40_), globalState))
                d_2691_v42_: D7
                d_2691_v42_ = D7_DC21(d_2690_v41_)
                d_2692_v43_: D7
                d_2692_v43_ = D7_DC21(d_2690_v41_)
                d_2693_v44_: _dafny.Seq
                d_2693_v44_ = _dafny.SeqWithoutIsStrInference([d_2692_v43_, d_2692_v43_, d_2692_v43_])
                d_2694_v45_: _dafny.Seq
                d_2694_v45_ = _dafny.SeqWithoutIsStrInference([d_2693_v44_])
                d_2695_v46_: _dafny.Set
                d_2695_v46_ = _dafny.Set({d_2693_v44_, (d_2694_v45_)[default__.safeIndex((self).f26, len(d_2694_v45_))]})
                d_2696_v47_: D16
                d_2696_v47_ = D16_DC44((d_2695_v46_) | (d_2695_v46_))
                d_2696_v47_ = d_2696_v47_
                d_2697_v48_: _dafny.Map
                d_2697_v48_ = _dafny.Map({self.f29: self.f29})
                d_2698_v49_: D11
                d_2698_v49_ = D11_DC34(d_2644_v0_, p0)
                d_2697_v48_ = (d_2697_v48_).set((self).fm7(self.f30, (d_2698_v49_).cf67, globalState), default__.safeModuloInt(self.f29, self.f30))
            elif True:
                d_2699_v50_: _dafny.Array
                def lambda126_(d_2700_i2_):
                    return D15_DC43()

                init72_ = lambda126_
                nw425_ = _dafny.Array(None, 25)
                for i0_72_ in range(nw425_.length(0)):
                    nw425_[i0_72_] = init72_(i0_72_)
                d_2699_v50_ = nw425_
                d_2701_v51_: D15
                d_2701_v51_ = D15_DC43()
                index455_ = default__.safeIndex(287, (d_2699_v50_).length(0))
                (d_2699_v50_)[index455_] = d_2701_v51_
                index456_ = default__.safeIndex(287, (d_2699_v50_).length(0))
                (d_2699_v50_)[index456_] = d_2701_v51_
                index457_ = default__.safeIndex(376, (d_2644_v0_).length(0))
                (d_2644_v0_)[index457_] = p0
                d_2702_v52_: T0
                nw426_ = C1()
                nw426_.ctor__()
                d_2702_v52_ = nw426_
                d_2703_v53_: _dafny.Seq
                d_2703_v53_ = _dafny.SeqWithoutIsStrInference([d_2702_v52_])
                index458_ = default__.safeIndex(376, (d_2644_v0_).length(0))
                (d_2644_v0_)[index458_] = (d_2703_v53_) != ((_dafny.SeqWithoutIsStrInference([d_2702_v52_, d_2702_v52_]) if p0 else d_2703_v53_))
                d_2704_v54_: _dafny.MultiSet
                d_2704_v54_ = _dafny.MultiSet([True])
                d_2705_v55_: _dafny.Map
                d_2705_v55_ = _dafny.Map({self.f29: 601})
                d_2706_v56_: C4
                nw427_ = C4()
                nw427_.ctor__(p0, (d_2704_v54_).cardinality, _dafny.SeqWithoutIsStrInference([((self).f27)[default__.safeIndex(self.f30, len((self).f27))] for d_2707_i3_ in range(default__.abs(759))]), d_2705_v55_)
                d_2706_v56_ = nw427_
                d_2708_v57_: C1
                nw428_ = C1()
                nw428_.ctor__()
                d_2708_v57_ = nw428_
                d_2709_v59_: _dafny.Seq
                def iife231_():
                    coll123_ = _dafny.Map()
                    compr_123_: int
                    for compr_123_ in _dafny.IntegerRange(626, -472):
                        d_2710_v58_: int = compr_123_
                        if ((626) <= (d_2710_v58_)) and ((d_2710_v58_) < (-472)):
                            coll123_[(d_2710_v58_) * (self.f29)] = D6_DC16((self).f26, (d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))])
                    return _dafny.Map(coll123_)
                d_2709_v59_ = _dafny.SeqWithoutIsStrInference([iife231_()
                ])
                d_2711_v60_: D6
                d_2711_v60_ = D6_DC16(self.f29, (d_2706_v56_).f32)
                d_2712_v61_: _dafny.Map
                d_2712_v61_ = _dafny.Map({self.f30: d_2711_v60_})
                d_2713_v62_: _dafny.MultiSet
                d_2713_v62_ = _dafny.MultiSet([(self).f26, self.f30])
                d_2714_v63_: _dafny.Map
                d_2714_v63_ = _dafny.Map({p0: False})
                d_2715_v64_: _dafny.Array
                nw429_ = _dafny.Array(None, 24)
                nw429_[int(0)] = (d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))]
                nw429_[int(1)] = (d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))]
                nw429_[int(2)] = (d_2706_v56_).f32
                nw429_[int(3)] = p0
                nw429_[int(4)] = (self.f29) == ((self).f26)
                nw429_[int(5)] = (d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))]
                nw429_[int(6)] = (d_2706_v56_).f32
                nw429_[int(7)] = (d_2706_v56_).f32
                nw429_[int(8)] = ((d_2709_v59_)[default__.safeIndex(self.f30, len(d_2709_v59_))]) != (d_2712_v61_)
                nw429_[int(9)] = (d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))]
                nw429_[int(10)] = (self.f30) >= (self.f30)
                nw429_[int(11)] = p0
                nw429_[int(12)] = not ((d_2706_v56_).fm8(d_2713_v62_, p0, self.f29, (self).f26, globalState)) or ((d_2706_v56_).f32)
                nw429_[int(13)] = not (False) or (not(p0))
                nw429_[int(14)] = (d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))]
                nw429_[int(15)] = not(not((d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))]))
                nw429_[int(16)] = (d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))]
                nw429_[int(17)] = (d_2706_v56_).f32
                nw429_[int(18)] = p0
                nw429_[int(19)] = (d_2706_v56_).f32
                nw429_[int(20)] = (D10_DC31((d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))], _dafny.SeqWithoutIsStrInference([len((self).f27), self.f29]), p0, (self).f26)).cf60
                nw429_[int(21)] = (d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))]
                nw429_[int(22)] = (d_2644_v0_)[default__.safeIndex(376, (d_2644_v0_).length(0))]
                nw429_[int(23)] = (d_2714_v63_) != (d_2714_v63_)
                d_2715_v64_ = nw429_
                (globalState).f1 = d_2715_v64_
            index459_ = default__.safeIndex(938, (d_2644_v0_).length(0))
            (d_2644_v0_)[index459_] = p0
            index460_ = default__.safeIndex(938, (d_2644_v0_).length(0))
            (d_2644_v0_)[index460_] = False
            (globalState).f8 = (0) - (-274)
            d_2716_v65_: _dafny.Array
            def lambda127_(d_2717_i4_):
                return (d_2717_i4_) * (self.f29)

            init73_ = lambda127_
            nw430_ = _dafny.Array(None, 23)
            for i0_73_ in range(nw430_.length(0)):
                nw430_[i0_73_] = init73_(i0_73_)
            d_2716_v65_ = nw430_
            index461_ = default__.safeIndex(148, (d_2716_v65_).length(0))
            (d_2716_v65_)[index461_] = self.f29
            index462_ = default__.safeIndex(148, (d_2716_v65_).length(0))
            (d_2716_v65_)[index462_] = (((self).f26) - (-609) if not(True) else self.f30)
        d_2718_v66_: _dafny.Map
        d_2718_v66_ = _dafny.Map({p0: not(False)})
        d_2719_v67_: _dafny.Set
        d_2719_v67_ = _dafny.Set({(0) - (self.f30)})
        d_2720_v68_: _dafny.Map
        d_2720_v68_ = _dafny.Map({((d_2718_v66_)[p0] if (p0) in (d_2718_v66_) else p0): d_2719_v67_})
        d_2720_v68_ = (d_2720_v68_).set(p0, d_2719_v67_)
        source35_ = p1
        if source35_.is_DC0:
            d_2721___mcc_h0_ = source35_.cf0
            d_2722_cf0_ = d_2721___mcc_h0_
            (globalState).f8 = -295
            d_2723_v69_: _dafny.Array
            nw431_ = _dafny.Array(int(0), 25)
            d_2723_v69_ = nw431_
            index463_ = default__.safeIndex(960, (d_2723_v69_).length(0))
            (d_2723_v69_)[index463_] = self.f30
            d_2724_v70_: _dafny.Seq
            d_2724_v70_ = _dafny.SeqWithoutIsStrInference([(self).f26])
            d_2725_v71_: D9
            d_2725_v71_ = D9_DC28()
            index464_ = default__.safeIndex(617, (d_2723_v69_).length(0))
            (d_2723_v69_)[index464_] = 582
            index465_ = default__.safeIndex(994, (d_2644_v0_).length(0))
            (d_2644_v0_)[index465_] = default__.fm24(p0, p0, default__.fm24(p0, p0, p0, (self).f26, globalState), d_2722_cf0_, globalState)
            d_2726_v72_: _dafny.Seq
            d_2726_v72_ = _dafny.SeqWithoutIsStrInference([not(p0)])
            d_2727_v74_: _dafny.MultiSet
            def iife232_():
                coll124_ = _dafny.Map()
                compr_124_: int
                for compr_124_ in _dafny.IntegerRange(821, 47):
                    d_2728_v73_: int = compr_124_
                    if ((821) <= (d_2728_v73_)) and ((d_2728_v73_) < (47)):
                        coll124_[default__.safeDivisionInt(d_2728_v73_, self.f29)] = (_dafny.MultiSet([self.f30])).cardinality
                return _dafny.Map(coll124_)
            d_2727_v74_ = _dafny.MultiSet([iife232_()
            ])
            d_2729_v75_: _dafny.Map
            d_2729_v75_ = _dafny.Map({self.f30: len(d_2724_v70_)})
            index466_ = default__.safeIndex(960, (d_2723_v69_).length(0))
            index467_ = default__.safeIndex(617, (d_2723_v69_).length(0))
            index468_ = default__.safeIndex(994, (d_2644_v0_).length(0))
            rhs479_ = len(((d_2726_v72_) + (d_2726_v72_)) + (d_2726_v72_))
            rhs480_ = (((d_2724_v70_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2726_v72_, (d_2726_v72_).set(default__.safeIndex(self.f29, len(d_2726_v72_)), p0)])), len(d_2724_v70_)), ((d_2727_v74_)[d_2729_v75_] if (d_2729_v75_) in (d_2727_v74_) else self.f29))).set(default__.safeIndex(d_2722_cf0_, len((d_2724_v70_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2726_v72_, (d_2726_v72_).set(default__.safeIndex(self.f29, len(d_2726_v72_)), p0)])), len(d_2724_v70_)), ((d_2727_v74_)[d_2729_v75_] if (d_2729_v75_) in (d_2727_v74_) else self.f29)))), len((self).f27))).set(default__.safeIndex((self).f26, len(((d_2724_v70_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2726_v72_, (d_2726_v72_).set(default__.safeIndex(self.f29, len(d_2726_v72_)), p0)])), len(d_2724_v70_)), ((d_2727_v74_)[d_2729_v75_] if (d_2729_v75_) in (d_2727_v74_) else self.f29))).set(default__.safeIndex(d_2722_cf0_, len((d_2724_v70_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2726_v72_, (d_2726_v72_).set(default__.safeIndex(self.f29, len(d_2726_v72_)), p0)])), len(d_2724_v70_)), ((d_2727_v74_)[d_2729_v75_] if (d_2729_v75_) in (d_2727_v74_) else self.f29)))), len((self).f27)))), (self).f26)
            rhs481_ = default__.fm38(p0, globalState)
            rhs482_ = ((self).f26) - (730)
            rhs483_ = p0
            lhs445_ = d_2723_v69_
            lhs446_ = default__.safeIndex(960, (d_2723_v69_).length(0))
            lhs447_ = d_2723_v69_
            lhs448_ = default__.safeIndex(617, (d_2723_v69_).length(0))
            lhs449_ = d_2644_v0_
            lhs450_ = default__.safeIndex(994, (d_2644_v0_).length(0))
            lhs445_[lhs446_] = rhs479_
            d_2724_v70_ = rhs480_
            d_2725_v71_ = rhs481_
            lhs447_[lhs448_] = rhs482_
            lhs449_[lhs450_] = rhs483_
            d_2730_v76_: D7
            d_2730_v76_ = D7_DC20(not(p0), p0)
            d_2731_v77_: _dafny.Map
            d_2731_v77_ = _dafny.Map({(self).f26: d_2730_v76_})
            d_2732_v78_: _dafny.Seq
            d_2732_v78_ = _dafny.SeqWithoutIsStrInference([d_2731_v77_, d_2731_v77_])
            (globalState).f9 = (d_2731_v77_) not in (d_2732_v78_)
            d_2733_v79_: C1
            nw432_ = C1()
            nw432_.ctor__()
            d_2733_v79_ = nw432_
        elif source35_.is_DC1:
            d_2734___mcc_h1_ = source35_.cf1
            d_2735___mcc_h2_ = source35_.cf2
            d_2736___mcc_h3_ = source35_.cf3
            d_2737_cf3_ = d_2736___mcc_h3_
            d_2738_cf2_ = d_2735___mcc_h2_
            d_2739_cf1_ = d_2734___mcc_h1_
            r0 = d_2738_cf2_
            d_2740_v80_: _dafny.Map
            d_2740_v80_ = _dafny.Map({d_2738_cf2_: p0})
            d_2741_v81_: _dafny.Seq
            d_2741_v81_ = _dafny.SeqWithoutIsStrInference([(self).f27, d_2739_cf1_])
            d_2742_v82_: _dafny.Seq
            d_2742_v82_ = _dafny.SeqWithoutIsStrInference([d_2740_v80_, default__.fm18(len(d_2741_v81_), globalState), d_2740_v80_])
            d_2743_v83_: _dafny.Seq
            d_2743_v83_ = _dafny.SeqWithoutIsStrInference([d_2742_v82_, d_2742_v82_])
            d_2744_v84_: _dafny.Seq
            d_2744_v84_ = _dafny.SeqWithoutIsStrInference([self.f29])
            d_2745_v85_: _dafny.Seq
            d_2745_v85_ = _dafny.SeqWithoutIsStrInference([len(d_2744_v84_)])
            d_2746_v86_: _dafny.Seq
            d_2746_v86_ = _dafny.SeqWithoutIsStrInference([(self).f26, len(d_2745_v85_)])
            (globalState).f24 = len((d_2742_v82_ if p0 else (((d_2743_v83_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2737_cf3_ for d_2747_i5_ in range(default__.abs(181))])), len(d_2743_v83_))]) + (_dafny.SeqWithoutIsStrInference([d_2740_v80_ for d_2748_i6_ in range(default__.abs(383))]))).set(default__.safeIndex(self.f29, len(((d_2743_v83_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2737_cf3_ for d_2749_i5_ in range(default__.abs(181))])), len(d_2743_v83_))]) + (_dafny.SeqWithoutIsStrInference([d_2740_v80_ for d_2750_i6_ in range(default__.abs(383))])))), default__.fm18((self).fm7(len(d_2746_v86_), p0, globalState), globalState))))
            d_2751_v87_: _dafny.Seq
            d_2751_v87_ = _dafny.SeqWithoutIsStrInference([d_2644_v0_])
            pat_let_tv68_ = d_2751_v87_
            pat_let_tv69_ = d_2738_cf2_
            pat_let_tv70_ = d_2751_v87_
            pat_let_tv71_ = d_2644_v0_
            d_2752_v88_: _dafny.Seq
            def iife234_(_pat_let55_0):
                def iife235_(d_2753_dt__update__tmp_h1_):
                    def iife236_(_pat_let56_0):
                        def iife237_(d_2754_dt__update_hcf33_h0_):
                            return D7_DC17(d_2754_dt__update_hcf33_h0_)
                        return iife237_(_pat_let56_0)
                    return iife236_((pat_let_tv68_)[default__.safeIndex(pat_let_tv69_, len(pat_let_tv70_))])
                return iife235_(_pat_let55_0)
            def iife233_(_pat_let54_0):
                def iife238_(d_2755_dt__update__tmp_h2_):
                    def iife239_(_pat_let57_0):
                        def iife240_(d_2756_dt__update_hcf33_h1_):
                            return D7_DC17(d_2756_dt__update_hcf33_h1_)
                        return iife240_(_pat_let57_0)
                    return iife239_(pat_let_tv71_)
                return iife238_(_pat_let54_0)
            d_2752_v88_ = _dafny.SeqWithoutIsStrInference([iife233_(iife234_(d_2645_v1_)), d_2645_v1_])
            (globalState).f24 = default__.safeDivisionInt((len((d_2752_v88_).set(default__.safeIndex(d_2738_cf2_, len(d_2752_v88_)), d_2645_v1_))) - (self.f29), d_2738_cf2_)
            d_2757_v89_: _dafny.Seq
            d_2757_v89_ = _dafny.SeqWithoutIsStrInference([p0])
            index469_ = default__.safeIndex(2, (d_2644_v0_).length(0))
            (d_2644_v0_)[index469_] = (d_2757_v89_)[default__.safeIndex(self.f29, len(d_2757_v89_))]
            index470_ = default__.safeIndex(2, (d_2644_v0_).length(0))
            (d_2644_v0_)[index470_] = p0
        elif source35_.is_DC2:
            d_2758___mcc_h4_ = source35_.cf4
            d_2759_cf4_ = d_2758___mcc_h4_
            (globalState).f13 = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('n') if d_2759_cf4_ else _dafny.CodePoint('x')) for d_2760_i7_ in range(default__.abs(971))])
            (globalState).f24 = self.f29
            (globalState).f9 = (((self).f27) + ((self).f27)) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2761_i8_ in range(default__.abs(823))]))
            d_2762_v90_: _dafny.MultiSet
            d_2762_v90_ = _dafny.MultiSet([self.f30, self.f30, (self).f26])
            d_2763_v91_: str
            d_2763_v91_ = _dafny.CodePoint('j')
            d_2764_v92_: _dafny.Array
            nw433_ = _dafny.Array(None, 21)
            nw433_[int(0)] = ((self).f27) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxbnntp")))
            nw433_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2765_i9_ in range(default__.abs(-932))])
            nw433_[int(2)] = ((self).f27 if (self).fm8(d_2762_v90_, not(d_2759_cf4_), self.f29, (d_2762_v90_).cardinality, globalState) else (self).f27)
            nw433_[int(3)] = ((self).f27 if d_2759_cf4_ else (self).f27)
            nw433_[int(4)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2766_i10_ in range(default__.abs(444))])
            nw433_[int(5)] = default__.fm36(d_2759_cf4_, d_2759_cf4_, p0, globalState)
            nw433_[int(6)] = ((self).f27) + ((self).f27)
            nw433_[int(7)] = (self).f27
            nw433_[int(8)] = (self).f27
            nw433_[int(9)] = ((self).f27) + ((self).f27)
            nw433_[int(10)] = (self).f27
            nw433_[int(11)] = (self).f27
            nw433_[int(12)] = (self).f27
            nw433_[int(13)] = (self).f27
            nw433_[int(14)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2767_i11_ in range(default__.abs(-465))])
            nw433_[int(15)] = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2768_i12_ in range(default__.abs(846))])) + ((self).f27)).set(default__.safeIndex(-90, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2769_i12_ in range(default__.abs(846))])) + ((self).f27))), d_2763_v91_)
            nw433_[int(16)] = (self).f27
            nw433_[int(17)] = (self).f27
            nw433_[int(18)] = (self).f27
            nw433_[int(19)] = _dafny.SeqWithoutIsStrInference([d_2763_v91_ for d_2770_i13_ in range(default__.abs(-359))])
            nw433_[int(20)] = ((self).f27) + (_dafny.SeqWithoutIsStrInference([((self).f27)[default__.safeIndex(-859, len((self).f27))] for d_2771_i14_ in range(default__.abs(-391))]))
            d_2764_v92_ = nw433_
            index471_ = default__.safeIndex(123, (d_2764_v92_).length(0))
            (d_2764_v92_)[index471_] = (self).f27
            d_2772_v93_: _dafny.Seq
            d_2772_v93_ = _dafny.SeqWithoutIsStrInference([687])
            d_2773_v94_: T2
            nw434_ = C2()
            nw434_.ctor__(d_2772_v93_, _dafny.Map({self.f30: (self).f26}))
            d_2773_v94_ = nw434_
            d_2774_v95_: C1
            nw435_ = C1()
            nw435_.ctor__()
            d_2774_v95_ = nw435_
            d_2775_v96_: _dafny.Map
            d_2775_v96_ = _dafny.Map({d_2774_v95_: -577})
            d_2776_v97_: D17
            d_2776_v97_ = D17_DC48(d_2773_v94_, len(d_2775_v96_), default__.fm42((D21_DC56(d_2719_v67_)).cf107, globalState))
            index472_ = default__.safeIndex(123, (d_2764_v92_).length(0))
            rhs484_ = p0
            rhs485_ = p0
            rhs486_ = p0
            rhs487_ = self.f29
            rhs488_ = (default__.fm36(p0, d_2759_cf4_, p0, globalState)).set(default__.safeIndex((679) - ((self).f26), len(default__.fm36(p0, d_2759_cf4_, p0, globalState))), default__.fm42(_dafny.Set({(d_2762_v90_).cardinality, self.f29, (d_2776_v97_).cf90}), globalState))
            lhs451_ = globalState
            lhs452_ = globalState
            lhs453_ = d_2764_v92_
            lhs454_ = default__.safeIndex(123, (d_2764_v92_).length(0))
            d_2759_cf4_ = rhs484_
            lhs451_.f25 = rhs485_
            d_2759_cf4_ = rhs486_
            lhs452_.f24 = rhs487_
            lhs453_[lhs454_] = rhs488_
        elif source35_.is_DC3:
            d_2777___mcc_h5_ = source35_.cf5
            d_2778___mcc_h6_ = source35_.cf6
            d_2779_cf6_ = d_2778___mcc_h6_
            d_2780_cf5_ = d_2777___mcc_h5_
            d_2781_v98_: _dafny.Seq
            d_2781_v98_ = _dafny.SeqWithoutIsStrInference([p0])
            if ((self).f26) < (len(d_2781_v98_)):
                d_2782_v99_: _dafny.Array
                def lambda128_(d_2783_i15_):
                    return default__.safeModuloInt(d_2783_i15_, len((self).f27))

                init74_ = lambda128_
                nw436_ = _dafny.Array(None, 4)
                for i0_74_ in range(nw436_.length(0)):
                    nw436_[i0_74_] = init74_(i0_74_)
                d_2782_v99_ = nw436_
                d_2784_v100_: _dafny.Map
                d_2784_v100_ = _dafny.Map({(d_2782_v99_ if not(p0) else d_2782_v99_): d_2644_v0_})
                d_2784_v100_ = (d_2784_v100_).set(d_2782_v99_, d_2644_v0_)
                (globalState).f13 = (self).f27
                d_2785_v101_: _dafny.Set
                d_2785_v101_ = _dafny.Set({not(p0), p0})
                d_2786_v102_: _dafny.Map
                d_2786_v102_ = _dafny.Map({(0) - ((0) - (len(d_2785_v101_))): self.f30})
                d_2787_v103_: T0
                nw437_ = C4()
                nw437_.ctor__(False, 760, (self).f27, d_2786_v102_)
                d_2787_v103_ = nw437_
                d_2788_v104_: _dafny.Map
                d_2788_v104_ = _dafny.Map({384: d_2787_v103_})
                d_2789_v105_: _dafny.Seq
                d_2789_v105_ = _dafny.SeqWithoutIsStrInference([((d_2788_v104_)[d_2779_cf6_] if (d_2779_cf6_) in (d_2788_v104_) else d_2787_v103_), d_2787_v103_])
                d_2790_v106_: _dafny.Array
                nw438_ = _dafny.Array(None, 5)
                nw438_[int(0)] = d_2787_v103_
                nw438_[int(1)] = (d_2789_v105_)[default__.safeIndex(d_2779_cf6_, len(d_2789_v105_))]
                nw438_[int(2)] = (d_2787_v103_ if True else d_2787_v103_)
                nw438_[int(3)] = d_2787_v103_
                nw438_[int(4)] = d_2787_v103_
                d_2790_v106_ = nw438_
                index473_ = default__.safeIndex(560, (d_2790_v106_).length(0))
                (d_2790_v106_)[index473_] = d_2787_v103_
                index474_ = default__.safeIndex(560, (d_2790_v106_).length(0))
                (d_2790_v106_)[index474_] = d_2787_v103_
                (globalState).f25 = (d_2779_cf6_) != (((0) - (self.f30)) + ((self).f26))
                (globalState).f25 = False
            elif True:
                d_2791_v107_: _dafny.MultiSet
                d_2791_v107_ = _dafny.MultiSet([d_2780_cf5_])
                (globalState).f25 = (self).fm8(d_2791_v107_, p0, d_2780_cf5_, self.f29, globalState)
                d_2792_v108_: _dafny.Seq
                d_2792_v108_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrl")), (self).f27, (self).f27, (self).f27, (self).f27])
                (globalState).f9 = not(((self).f27) in (d_2792_v108_))
                (globalState).f25 = (p0) or (p0)
                d_2793_v109_: _dafny.Map
                d_2793_v109_ = _dafny.Map({(self).f27: not((d_2779_cf6_) != (self.f30))})
                d_2793_v109_ = (d_2793_v109_).set((self).f27, p0)
                d_2794_v110_: _dafny.Seq
                d_2794_v110_ = _dafny.SeqWithoutIsStrInference([self.f30, d_2779_cf6_])
                d_2795_v111_: _dafny.Map
                d_2795_v111_ = _dafny.Map({992: len((d_2718_v66_).set(p0, (d_2781_v98_)[default__.safeIndex(d_2780_cf5_, len(d_2781_v98_))]))})
                d_2796_v112_: C2
                nw439_ = C2()
                nw439_.ctor__(d_2794_v110_, d_2795_v111_)
                d_2796_v112_ = nw439_
                d_2797_v113_: D0
                d_2797_v113_ = D0_DC1((self).f27, len(_dafny.Map({d_2796_v112_: True})), d_2780_cf5_)
                d_2798_v114_: C3
                nw440_ = C3()
                nw440_.ctor__((0) - (len((d_2718_v66_) | (default__.fm34(p0, not(p0), p0, self.f29, globalState)))), (self).f26, (d_2797_v113_).cf1)
                d_2798_v114_ = nw440_
            d_2799_v115_: D2
            d_2799_v115_ = D2_DC8(p0, self.f30, (self).f26)
            d_2799_v115_ = default__.fm53(p0, globalState)
            (globalState).f9 = p0
            (globalState).f9 = p0
        elif True:
            d_2800___mcc_h7_ = source35_.cf7
            d_2801_cf7_ = d_2800___mcc_h7_
            (globalState).f25 = (p0) or (p0)
            if p0:
                d_2802_v116_: _dafny.Array
                nw441_ = _dafny.Array(None, 26)
                d_2802_v116_ = nw441_
                d_2802_v116_ = d_2802_v116_
                (globalState).f23 = (((self).f27) + (default__.fm43(p0, (self).f26, globalState))) + ((self).f27)
                d_2650_v4_ = (d_2650_v4_).set((self).f27, (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2803_i16_ in range(default__.abs(-982))]))) + ((self).f26))
                d_2804_v117_: C0
                nw442_ = C0()
                nw442_.ctor__()
                d_2804_v117_ = nw442_
                d_2805_v118_: _dafny.Map
                d_2805_v118_ = _dafny.Map({p0: self.f29})
                d_2806_v119_: str
                d_2806_v119_ = _dafny.CodePoint('b')
                d_2807_v120_: _dafny.Set
                d_2807_v120_ = _dafny.Set({p0})
                d_2808_v121_: D1
                d_2808_v121_ = D1_DC5(d_2807_v120_)
                d_2809_v122_: _dafny.Map
                d_2809_v122_ = _dafny.Map({(0) - (len((self).f27)): d_2808_v121_})
                d_2810_v123_: _dafny.Map
                d_2810_v123_ = d_2809_v122_
                d_2811_v124_: D8
                d_2811_v124_ = D8_DC25(p0, d_2805_v118_, d_2806_v119_, d_2810_v123_, 147)
                d_2812_v125_: _dafny.Array
                nw443_ = _dafny.Array(None, 4)
                nw443_[int(0)] = (self).f26
                nw443_[int(1)] = len((self).f27)
                nw443_[int(2)] = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orpypysup"))).set(default__.safeIndex((d_2811_v124_).cf53, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orpypysup")))), d_2806_v119_)).set(default__.safeIndex(981, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orpypysup"))).set(default__.safeIndex((d_2811_v124_).cf53, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orpypysup")))), d_2806_v119_))), d_2806_v119_))
                nw443_[int(3)] = len((self).f27)
                d_2812_v125_ = nw443_
                d_2813_v126_: _dafny.Map
                d_2813_v126_ = _dafny.Map({(self).f26: True})
                index475_ = default__.safeIndex(28, (d_2812_v125_).length(0))
                (d_2812_v125_)[index475_] = default__.safeModuloInt((self).fm6(d_2813_v126_, self.f30, _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, self.f29, self.f30]), _dafny.Map({_dafny.SeqWithoutIsStrInference([p0, p0]): p0}), globalState), (0) - ((self).f26))
                index476_ = default__.safeIndex(28, (d_2812_v125_).length(0))
                rhs489_ = (self.f30) < ((self).f26)
                rhs490_ = p0
                rhs491_ = d_2804_v117_
                rhs492_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvoen")))
                lhs455_ = globalState
                lhs456_ = globalState
                lhs457_ = d_2812_v125_
                lhs458_ = default__.safeIndex(28, (d_2812_v125_).length(0))
                lhs455_.f9 = rhs489_
                lhs456_.f9 = rhs490_
                d_2804_v117_ = rhs491_
                lhs457_[lhs458_] = rhs492_
                index477_ = default__.safeIndex(227, (d_2644_v0_).length(0))
                (d_2644_v0_)[index477_] = p0
                index478_ = default__.safeIndex(227, (d_2644_v0_).length(0))
                (d_2644_v0_)[index478_] = not(p0)
            elif True:
                (globalState).f5 = default__.safeModuloInt((self.f30) - (self.f30), -759)
                d_2814_v127_: _dafny.Map
                d_2814_v127_ = _dafny.Map({self.f30: self.f29})
                d_2815_v128_: T2
                nw444_ = C4()
                nw444_.ctor__(p0, (0) - (self.f30), (self).f27, d_2814_v127_)
                d_2815_v128_ = nw444_
                d_2816_v129_: str
                d_2816_v129_ = _dafny.CodePoint('l')
                d_2817_v130_: D17
                d_2817_v130_ = D17_DC48(d_2815_v128_, self.f29, d_2816_v129_)
                d_2817_v130_ = d_2817_v130_
                d_2818_v131_: _dafny.Array
                nw445_ = _dafny.Array(None, 2)
                nw445_[int(0)] = self.f30
                nw445_[int(1)] = self.f30
                d_2818_v131_ = nw445_
                index479_ = default__.safeIndex(669, (d_2818_v131_).length(0))
                (d_2818_v131_)[index479_] = (self).f26
                d_2819_v132_: _dafny.Map
                d_2819_v132_ = _dafny.Map({p0: len((self).f27)})
                d_2820_v133_: _dafny.Seq
                d_2820_v133_ = _dafny.SeqWithoutIsStrInference([True, p0])
                d_2821_v134_: _dafny.MultiSet
                d_2821_v134_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([self.f29])), 131])
                d_2822_v135_: _dafny.MultiSet
                d_2822_v135_ = _dafny.MultiSet([p0])
                d_2823_v136_: _dafny.Map
                d_2823_v136_ = _dafny.Map({d_2822_v135_: d_2818_v131_})
                d_2824_v137_: _dafny.Set
                d_2824_v137_ = _dafny.Set({d_2644_v0_})
                index480_ = default__.safeIndex(669, (d_2818_v131_).length(0))
                rhs493_ = self.f30
                rhs494_ = ((d_2819_v132_)[(d_2820_v133_)[default__.safeIndex((d_2821_v134_).cardinality, len(d_2820_v133_))]] if ((d_2820_v133_)[default__.safeIndex((d_2821_v134_).cardinality, len(d_2820_v133_))]) in (d_2819_v132_) else 525)
                rhs495_ = default__.fm24(p0, p0, p0, ((0) - (len(d_2718_v66_))) * (self.f29), globalState)
                rhs496_ = not(not((d_2822_v135_) not in (d_2823_v136_)))
                rhs497_ = ((d_2824_v137_) - (d_2824_v137_)).ispropersubset(d_2824_v137_)
                lhs459_ = d_2818_v131_
                lhs460_ = default__.safeIndex(669, (d_2818_v131_).length(0))
                lhs461_ = globalState
                lhs462_ = globalState
                lhs463_ = globalState
                lhs464_ = globalState
                lhs459_[lhs460_] = rhs493_
                lhs461_.f0 = rhs494_
                lhs462_.f25 = rhs495_
                lhs463_.f9 = rhs496_
                lhs464_.f9 = rhs497_
                d_2825_v138_: D16
                d_2825_v138_ = D16_DC45(len(d_2820_v133_), p0, self.f29, self.f30, p0)
                (self).f29 = (self.f30) * ((d_2825_v138_).cf85)
                d_2826_v139_: _dafny.Array
                nw446_ = _dafny.Array(None, 5)
                nw446_[int(0)] = d_2818_v131_
                nw446_[int(1)] = d_2818_v131_
                nw446_[int(2)] = d_2818_v131_
                nw446_[int(3)] = d_2818_v131_
                nw446_[int(4)] = d_2818_v131_
                d_2826_v139_ = nw446_
                index481_ = default__.safeIndex(949, (d_2826_v139_).length(0))
                (d_2826_v139_)[index481_] = d_2818_v131_
                index482_ = default__.safeIndex(949, (d_2826_v139_).length(0))
                (d_2826_v139_)[index482_] = d_2818_v131_
            d_2827_v140_: _dafny.Array
            nw447_ = _dafny.Array(int(0), 15)
            d_2827_v140_ = nw447_
            d_2828_v141_: _dafny.Map
            d_2828_v141_ = _dafny.Map({d_2827_v140_: self.f30})
            d_2829_v142_: _dafny.Map
            d_2829_v142_ = _dafny.Map({len((self).f27): len(default__.fm34(p0, p0, True, (self).f26, globalState))})
            d_2830_v143_: C4
            nw448_ = C4()
            nw448_.ctor__(p0, len(d_2828_v141_), (self).f27, d_2829_v142_)
            d_2830_v143_ = nw448_
            d_2831_v144_: _dafny.Seq
            d_2831_v144_ = _dafny.SeqWithoutIsStrInference([75])
            d_2832_v145_: D7
            d_2832_v145_ = D7_DC19(len((d_2831_v144_ if (d_2830_v143_).f32 else d_2831_v144_)), p0)
            d_2833_v146_: _dafny.MultiSet
            d_2833_v146_ = _dafny.MultiSet([(d_2830_v143_).f32, False, (d_2830_v143_).f32])
            rhs498_ = p0
            rhs499_ = default__.safeModuloInt(876, ((self).f26) - (-636))
            rhs500_ = D7_DC19((d_2833_v146_).cardinality, p0)
            lhs465_ = globalState
            lhs466_ = globalState
            lhs465_.f25 = rhs498_
            lhs466_.f17 = rhs499_
            d_2832_v145_ = rhs500_
        if not(p0):
            (globalState).f0 = self.f30
            d_2834_v147_: str
            d_2834_v147_ = _dafny.CodePoint('j')
            d_2834_v147_ = default__.fm42((d_2719_v67_) | (d_2719_v67_), globalState)
            (globalState).f5 = self.f30
            d_2718_v66_ = (d_2718_v66_).set(p0, p0)
            d_2835_v148_: _dafny.Array
            nw449_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_2835_v148_ = nw449_
            r2 = d_2835_v148_
        elif True:
            (globalState).f9 = p0
            d_2836_v149_: _dafny.Array
            nw450_ = _dafny.Array(None, 19)
            nw450_[int(0)] = (self).f27
            nw450_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qksflf"))
            nw450_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujenvwrw"))
            nw450_[int(3)] = (self).f27
            nw450_[int(4)] = (self).f27
            nw450_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsytxhdyk"))
            nw450_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wv"))
            nw450_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vioievb"))
            nw450_[int(8)] = (self).f27
            nw450_[int(9)] = (self).f27
            nw450_[int(10)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_2837_i17_ in range(default__.abs(107))])
            nw450_[int(11)] = (self).f27
            nw450_[int(12)] = (self).f27
            nw450_[int(13)] = (self).f27
            nw450_[int(14)] = (self).f27
            nw450_[int(15)] = (self).f27
            nw450_[int(16)] = (self).f27
            nw450_[int(17)] = (self).f27
            nw450_[int(18)] = (self).f27
            d_2836_v149_ = nw450_
            d_2838_v150_: _dafny.Map
            d_2838_v150_ = _dafny.Map({default__.safeDivisionInt(self.f30, self.f29): d_2836_v149_})
            d_2838_v150_ = (d_2838_v150_).set(self.f30, d_2836_v149_)
            d_2839_v151_: D0
            d_2839_v151_ = D0_DC3(self.f30, self.f29)
            d_2840_v152_: _dafny.Seq
            d_2840_v152_ = _dafny.SeqWithoutIsStrInference([d_2839_v151_, d_2839_v151_, d_2839_v151_])
            d_2841_v153_: D0
            d_2841_v153_ = D0_DC4((d_2840_v152_)[default__.safeIndex((0) - (self.f29), len(d_2840_v152_))])
            source36_ = d_2841_v153_
            if source36_.is_DC0:
                d_2842___mcc_h8_ = source36_.cf0
                d_2843_cf0_ = d_2842___mcc_h8_
                d_2844_v154_: _dafny.Seq
                d_2844_v154_ = _dafny.SeqWithoutIsStrInference([self.f29, self.f30, self.f29])
                d_2845_v155_: str
                d_2845_v155_ = _dafny.CodePoint('j')
                d_2846_v156_: _dafny.Map
                d_2846_v156_ = _dafny.Map({d_2845_v155_: p0})
                d_2847_v157_: _dafny.Array
                nw451_ = _dafny.Array(None, 8)
                nw451_[int(0)] = len(d_2846_v156_)
                nw451_[int(1)] = (d_2843_cf0_) + (self.f29)
                nw451_[int(2)] = d_2843_cf0_
                nw451_[int(3)] = self.f30
                nw451_[int(4)] = len(_dafny.Set({_dafny.CodePoint('c'), d_2845_v155_, _dafny.CodePoint('q'), _dafny.CodePoint('n')}))
                nw451_[int(5)] = default__.safeDivisionInt((self).f26, (self).f26)
                nw451_[int(6)] = d_2843_cf0_
                nw451_[int(7)] = d_2843_cf0_
                d_2847_v157_ = nw451_
                index483_ = default__.safeIndex(849, (d_2847_v157_).length(0))
                (d_2847_v157_)[index483_] = d_2843_cf0_
                d_2848_v158_: _dafny.Array
                nw452_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_2848_v158_ = nw452_
                d_2849_v159_: _dafny.Seq
                d_2849_v159_ = _dafny.SeqWithoutIsStrInference([d_2848_v158_])
                d_2850_v160_: _dafny.Map
                d_2850_v160_ = _dafny.Map({(d_2849_v159_)[default__.safeIndex(d_2843_cf0_, len(d_2849_v159_))]: d_2844_v154_})
                index484_ = default__.safeIndex(849, (d_2847_v157_).length(0))
                rhs501_ = ((d_2850_v160_)[d_2848_v158_] if (d_2848_v158_) in (d_2850_v160_) else d_2844_v154_)
                rhs502_ = d_2836_v149_
                rhs503_ = (p0) == ((self.f30) == (d_2843_cf0_))
                rhs504_ = (0) - (self.f30)
                rhs505_ = self.f29
                lhs467_ = globalState
                lhs468_ = d_2847_v157_
                lhs469_ = default__.safeIndex(849, (d_2847_v157_).length(0))
                d_2844_v154_ = rhs501_
                d_2836_v149_ = rhs502_
                lhs467_.f9 = rhs503_
                r0 = rhs504_
                lhs468_[lhs469_] = rhs505_
                (globalState).f9 = p0
                d_2844_v154_ = d_2844_v154_
                (globalState).f8 = (self).fm7(-705, p0, globalState)
            elif source36_.is_DC1:
                d_2851___mcc_h9_ = source36_.cf1
                d_2852___mcc_h10_ = source36_.cf2
                d_2853___mcc_h11_ = source36_.cf3
                d_2854_cf3_ = d_2853___mcc_h11_
                d_2855_cf2_ = d_2852___mcc_h10_
                d_2856_cf1_ = d_2851___mcc_h9_
                d_2857_v161_: _dafny.Seq
                d_2857_v161_ = _dafny.SeqWithoutIsStrInference([d_2719_v67_, d_2719_v67_, d_2719_v67_, (d_2719_v67_) | (d_2719_v67_), d_2719_v67_])
                d_2858_v162_: _dafny.MultiSet
                d_2858_v162_ = _dafny.MultiSet([len(d_2719_v67_)])
                def iife241_():
                    coll125_ = _dafny.Set()
                    compr_125_: int
                    for compr_125_ in (d_2858_v162_).Elements:
                        d_2859_v163_: int = compr_125_
                        if (d_2859_v163_) in (d_2858_v162_):
                            coll125_ = coll125_.union(_dafny.Set([(d_2859_v163_) * ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2860_i18_ in range(default__.abs(-237))]))])).cardinality)]))
                    return _dafny.Set(coll125_)
                d_2857_v161_ = (_dafny.SeqWithoutIsStrInference([d_2719_v67_, iife241_()
                ]) if p0 else _dafny.SeqWithoutIsStrInference([d_2719_v67_ for d_2861_i19_ in range(default__.abs(-204))]))
                nw453_ = _dafny.Array(False, 4)
                d_2644_v0_ = nw453_
                d_2862_v164_: _dafny.Map
                d_2862_v164_ = _dafny.Map({d_2644_v0_: d_2644_v0_})
                d_2862_v164_ = (d_2862_v164_).set(d_2644_v0_, d_2644_v0_)
                d_2863_v165_: C3
                nw454_ = C3()
                nw454_.ctor__(len((self).f27), self.f30, (self).f27)
                d_2863_v165_ = nw454_
            elif source36_.is_DC2:
                d_2864___mcc_h12_ = source36_.cf4
                d_2865_cf4_ = d_2864___mcc_h12_
                d_2866_v166_: _dafny.Map
                d_2866_v166_ = _dafny.Map({(self).f26: p0})
                d_2867_v167_: _dafny.Seq
                d_2867_v167_ = _dafny.SeqWithoutIsStrInference([600])
                d_2868_v168_: _dafny.Seq
                d_2868_v168_ = _dafny.SeqWithoutIsStrInference([True, p0])
                d_2869_v169_: _dafny.Map
                d_2869_v169_ = _dafny.Map({d_2868_v168_: p0})
                d_2870_v170_: _dafny.Seq
                d_2870_v170_ = _dafny.SeqWithoutIsStrInference([(self).fm6((d_2866_v166_).set(717, p0), self.f30, d_2867_v167_, d_2869_v169_, globalState), 481, self.f30])
                d_2867_v167_ = (d_2867_v167_) + (_dafny.SeqWithoutIsStrInference([(self).f26]))
                d_2871_v171_: _dafny.Array
                nw455_ = _dafny.Array(D19.default()(), 14)
                d_2871_v171_ = nw455_
                d_2871_v171_ = (d_2871_v171_ if p0 else d_2871_v171_)
                d_2872_v172_: _dafny.Array
                nw456_ = _dafny.Array(_dafny.Seq({}), 3)
                d_2872_v172_ = nw456_
                d_2873_v173_: _dafny.Seq
                d_2873_v173_ = _dafny.SeqWithoutIsStrInference([d_2872_v172_])
                d_2874_v174_: D22
                d_2874_v174_ = D22_DC61((d_2873_v173_)[default__.safeIndex(self.f29, len(d_2873_v173_))])
                rhs506_ = self.f30
                rhs507_ = (d_2874_v174_).cf118
                rhs508_ = d_2865_cf4_
                lhs470_ = globalState
                r0 = rhs506_
                d_2872_v172_ = rhs507_
                lhs470_.f25 = rhs508_
                d_2875_v175_: _dafny.Array
                def lambda129_(d_2876_v168_, d_2877_p0_):
                    def lambda130_(d_2878_i20_):
                        return _dafny.Map({(d_2876_v168_)[default__.safeIndex(393, len(d_2876_v168_))]: d_2877_p0_})

                    return lambda130_

                init75_ = lambda129_(d_2868_v168_, p0)
                nw457_ = _dafny.Array(None, 27)
                for i0_75_ in range(nw457_.length(0)):
                    nw457_[i0_75_] = init75_(i0_75_)
                d_2875_v175_ = nw457_
                d_2879_v176_: _dafny.MultiSet
                d_2879_v176_ = _dafny.MultiSet([self.f30, self.f30])
                index485_ = default__.safeIndex(823, (d_2875_v175_).length(0))
                (d_2875_v175_)[index485_] = _dafny.Map({(self).fm8(d_2879_v176_, d_2865_cf4_, self.f30, self.f30, globalState): d_2865_cf4_})
                index486_ = default__.safeIndex(442, (d_2644_v0_).length(0))
                (d_2644_v0_)[index486_] = default__.fm24(d_2865_cf4_, p0, (self).fm8(d_2879_v176_, d_2865_cf4_, -445, 652, globalState), (0) - ((self).f26), globalState)
                d_2880_v177_: _dafny.Seq
                d_2880_v177_ = _dafny.SeqWithoutIsStrInference([(d_2718_v66_) | (d_2718_v66_)])
                index487_ = default__.safeIndex(823, (d_2875_v175_).length(0))
                index488_ = default__.safeIndex(442, (d_2644_v0_).length(0))
                rhs509_ = default__.safeModuloInt((self.f29) - (self.f30), (self.f29) - ((self).f26))
                rhs510_ = (d_2880_v177_)[default__.safeIndex(-25, len(d_2880_v177_))]
                rhs511_ = p0
                rhs512_ = p0
                lhs471_ = globalState
                lhs472_ = d_2875_v175_
                lhs473_ = default__.safeIndex(823, (d_2875_v175_).length(0))
                lhs474_ = d_2644_v0_
                lhs475_ = default__.safeIndex(442, (d_2644_v0_).length(0))
                lhs471_.f24 = rhs509_
                lhs472_[lhs473_] = rhs510_
                lhs474_[lhs475_] = rhs511_
                d_2865_cf4_ = rhs512_
            elif source36_.is_DC3:
                d_2881___mcc_h13_ = source36_.cf5
                d_2882___mcc_h14_ = source36_.cf6
                d_2883_cf6_ = d_2882___mcc_h14_
                d_2884_cf5_ = d_2881___mcc_h13_
                d_2885_v178_: str
                d_2885_v178_ = _dafny.CodePoint('p')
                d_2885_v178_ = _dafny.CodePoint('g')
                (globalState).f9 = False
                d_2886_v179_: D21
                d_2886_v179_ = D21_DC56(d_2719_v67_)
                d_2887_v180_: _dafny.Set
                d_2887_v180_ = _dafny.Set({d_2886_v179_})
                d_2888_v181_: _dafny.Array
                nw458_ = _dafny.Array(int(0), 6)
                d_2888_v181_ = nw458_
                d_2889_v182_: D7
                d_2889_v182_ = D7_DC18((self).f26, d_2888_v181_)
                d_2890_v184_: _dafny.Seq
                def iife242_():
                    coll126_ = _dafny.Set()
                    compr_126_: D21
                    for compr_126_ in (_dafny.Map({d_2886_v179_: p0})).keys.Elements:
                        d_2891_v183_: D21 = compr_126_
                        if (d_2891_v183_) in (_dafny.Map({d_2886_v179_: p0})):
                            coll126_ = coll126_.union(_dafny.Set([d_2891_v183_]))
                    return _dafny.Set(coll126_)
                d_2890_v184_ = _dafny.SeqWithoutIsStrInference([d_2887_v180_, d_2887_v180_, (_dafny.Set({d_2886_v179_, d_2886_v179_, d_2886_v179_})).intersection(iife242_()
                ), d_2887_v180_])
                rhs513_ = (d_2889_v182_).cf34
                rhs514_ = (d_2890_v184_)[default__.safeIndex(-221, len(d_2890_v184_))]
                rhs515_ = d_2888_v181_
                rhs516_ = (self.f29) - (120)
                lhs476_ = globalState
                lhs477_ = globalState
                lhs476_.f5 = rhs513_
                d_2887_v180_ = rhs514_
                d_2888_v181_ = rhs515_
                lhs477_.f24 = rhs516_
                d_2892_v185_: C1
                nw459_ = C1()
                nw459_.ctor__()
                d_2892_v185_ = nw459_
            elif True:
                d_2893___mcc_h15_ = source36_.cf7
                d_2894_cf7_ = d_2893___mcc_h15_
                d_2895_v186_: _dafny.Map
                d_2895_v186_ = _dafny.Map({(0) - ((self).f26): p0})
                d_2896_v187_: _dafny.Seq
                d_2896_v187_ = _dafny.SeqWithoutIsStrInference([self.f30, self.f29, (self).f26])
                d_2897_v188_: _dafny.Seq
                d_2897_v188_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2898_v189_: _dafny.Map
                d_2898_v189_ = _dafny.Map({d_2897_v188_: p0})
                (globalState).f0 = (self).fm7(default__.safeDivisionInt((self).f26, self.f29), ((self).fm6(d_2895_v186_, self.f30, (d_2896_v187_).set(default__.safeIndex(len((self).f27), len(d_2896_v187_)), len(_dafny.Map({p0: (self).f26}))), d_2898_v189_, globalState)) == (self.f29), globalState)
                d_2899_v190_: _dafny.Set
                d_2899_v190_ = _dafny.Set({p0})
                d_2899_v190_ = d_2899_v190_
                d_2900_v191_: C1
                nw460_ = C1()
                nw460_.ctor__()
                d_2900_v191_ = nw460_
                d_2901_v192_: _dafny.Map
                d_2901_v192_ = _dafny.Map({p0: (self).f26})
                d_2902_v193_: str
                d_2902_v193_ = _dafny.CodePoint('m')
                d_2903_v194_: _dafny.Map
                d_2903_v194_ = _dafny.Map({len((self).f27): D1_DC5(d_2899_v190_)})
                d_2904_v195_: D8
                d_2904_v195_ = D8_DC25(p0, d_2901_v192_, d_2902_v193_, d_2903_v194_, (self).f26)
                d_2905_v196_: _dafny.Map
                d_2905_v196_ = _dafny.Map({p0: d_2904_v195_})
                d_2906_v197_: _dafny.Map
                d_2906_v197_ = _dafny.Map({p0: d_2905_v196_})
                d_2906_v197_ = (d_2906_v197_).set(p0, (d_2905_v196_) | (d_2905_v196_))
            d_2907_v198_: _dafny.Set
            d_2907_v198_ = _dafny.Set({p0, p0})
            d_2908_v199_: _dafny.Set
            d_2908_v199_ = _dafny.Set({d_2907_v198_, d_2907_v198_, d_2907_v198_})
            d_2909_v200_: _dafny.Map
            d_2909_v200_ = _dafny.Map({_dafny.Set({p0, p0}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))})
            d_2910_v202_: _dafny.Array
            nw461_ = _dafny.Array(None, 15)
            nw461_[int(0)] = d_2908_v199_
            nw461_[int(1)] = d_2908_v199_
            nw461_[int(2)] = _dafny.Set({d_2907_v198_, d_2907_v198_})
            nw461_[int(3)] = d_2908_v199_
            nw461_[int(4)] = d_2908_v199_
            nw461_[int(5)] = d_2908_v199_
            nw461_[int(6)] = d_2908_v199_
            nw461_[int(7)] = d_2908_v199_
            nw461_[int(8)] = d_2908_v199_
            nw461_[int(9)] = _dafny.Set({d_2907_v198_})
            nw461_[int(10)] = d_2908_v199_
            def iife243_():
                coll127_ = _dafny.Set()
                compr_127_: _dafny.Set
                for compr_127_ in (d_2909_v200_).keys.Elements:
                    d_2911_v201_: _dafny.Set = compr_127_
                    if (d_2911_v201_) in (d_2909_v200_):
                        coll127_ = coll127_.union(_dafny.Set([d_2911_v201_]))
                return _dafny.Set(coll127_)
            nw461_[int(11)] = iife243_()
            
            nw461_[int(12)] = d_2908_v199_
            nw461_[int(13)] = d_2908_v199_
            nw461_[int(14)] = _dafny.Set({d_2907_v198_})
            d_2910_v202_ = nw461_
            index489_ = default__.safeIndex(433, (d_2910_v202_).length(0))
            (d_2910_v202_)[index489_] = d_2908_v199_
            d_2912_v203_: _dafny.MultiSet
            d_2912_v203_ = _dafny.MultiSet([d_2907_v198_])
            index490_ = default__.safeIndex(433, (d_2910_v202_).length(0))
            def iife244_():
                coll128_ = _dafny.Set()
                compr_128_: _dafny.Set
                for compr_128_ in ((d_2912_v203_).set(d_2907_v198_, default__.abs(143))).Elements:
                    d_2913_v204_: _dafny.Set = compr_128_
                    if (d_2913_v204_) in ((d_2912_v203_).set(d_2907_v198_, default__.abs(143))):
                        coll128_ = coll128_.union(_dafny.Set([d_2913_v204_]))
                return _dafny.Set(coll128_)
            (d_2910_v202_)[index490_] = iife244_()
            
            d_2914_v205_: _dafny.MultiSet
            d_2914_v205_ = _dafny.MultiSet([(self).f26, self.f29])
            d_2914_v205_ = (d_2914_v205_).intersection(((_dafny.MultiSet([(self).f26])).set((self).f26, default__.abs(((d_2914_v205_)[self.f29] if (self.f29) in (d_2914_v205_) else self.f30)))) - (d_2914_v205_))
        d_2915_v206_: _dafny.Array
        nw462_ = _dafny.Array(int(0), 2)
        d_2915_v206_ = nw462_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_2915_v206_).length(0)):
            d_2916_i21_: int = guard_loop_5_
            if (True) and (((0) <= (d_2916_i21_)) and ((d_2916_i21_) < ((d_2915_v206_).length(0)))):
                (d_2915_v206_)[(d_2916_i21_)] = (d_2916_i21_) - (len(_dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])))
        r0 = ((self).f26) - (self.f29)
        d_2917_v207_: D0
        d_2917_v207_ = D0_DC0(-854)
        d_2918_v208_: _dafny.MultiSet
        d_2918_v208_ = _dafny.MultiSet([d_2917_v207_])
        r1 = (d_2918_v208_).intersection(d_2918_v208_)
        nw463_ = _dafny.Array(_dafny.Array(None, 0), 22)
        r2 = nw463_
        return r0, r1, r2


class C6(T0, T1):
    def  __init__(self):
        self._f26: int = int(0)
        self._f27: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f28: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f28, f26, f27):
        (self)._f28 = f28
        (self)._f26 = f26
        (self)._f27 = f27

    def fm6(self, p0, p1, p2, p3, globalState):
        def lambda131_(source37_):
            if source37_.is_DC0:
                d_2919___mcc_h0_ = source37_.cf0
                d_2920_cf0_ = d_2919___mcc_h0_
                return (_dafny.Map({D0_DC0(d_2920_cf0_): True})) | (_dafny.Map({D0_DC0((self).f26): False}))
            elif source37_.is_DC1:
                d_2921___mcc_h1_ = source37_.cf1
                d_2922___mcc_h2_ = source37_.cf2
                d_2923___mcc_h3_ = source37_.cf3
                d_2924_cf3_ = d_2923___mcc_h3_
                d_2925_cf2_ = d_2922___mcc_h2_
                d_2926_cf1_ = d_2921___mcc_h1_
                def iife245_():
                    coll129_ = _dafny.Map()
                    compr_129_: D0
                    for compr_129_ in (_dafny.SeqWithoutIsStrInference([D0_DC0(d_2925_cf2_), D0_DC0((self).f26)])).Elements:
                        d_2927_v0_: D0 = compr_129_
                        if (d_2927_v0_) in (_dafny.SeqWithoutIsStrInference([D0_DC0(d_2925_cf2_), D0_DC0((self).f26)])):
                            coll129_[d_2927_v0_] = True
                    return _dafny.Map(coll129_)
                return (iife245_()
                ) | (_dafny.Map({D0_DC0((self).f26): False}))
            elif source37_.is_DC2:
                d_2928___mcc_h4_ = source37_.cf4
                d_2929_cf4_ = d_2928___mcc_h4_
                return _dafny.Map({D0_DC0(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_2929_cf4_: (self).f26})) for d_2930_i0_ in range(default__.abs(527))]))): d_2929_cf4_})
            elif source37_.is_DC3:
                d_2931___mcc_h5_ = source37_.cf5
                d_2932___mcc_h6_ = source37_.cf6
                d_2933_cf6_ = d_2932___mcc_h6_
                d_2934_cf5_ = d_2931___mcc_h5_
                return _dafny.Map({D0_DC0(d_2934_cf5_): True})
            elif True:
                d_2935___mcc_h7_ = source37_.cf7
                d_2936_cf7_ = d_2935___mcc_h7_
                return _dafny.Map({D0_DC0((self).f26): not(True)})

        return len(lambda131_(D0_DC2(False)))

    def fm7(self, p0, p1, globalState):
        return len((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")), (self).f27, (D0_DC1((self).f27, 61, len((self).f27))).cf1])): (self).f27})) | (_dafny.Map({(self).f26: (self).f27})))

    def fm8(self, p0, p1, p2, p3, globalState):
        return not(True)

    def fm9(self, p0, p1, p2, globalState):
        return (D0_DC3((self).f26, (self).f26)).cf6

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        d_2937_v0_: D0
        d_2937_v0_ = D0_DC0((self).f26)
        source38_ = d_2937_v0_
        if source38_.is_DC0:
            d_2938___mcc_h0_ = source38_.cf0
            d_2939_cf0_ = d_2938___mcc_h0_
            d_2940_v1_: str
            d_2940_v1_ = _dafny.CodePoint('g')
            d_2940_v1_ = d_2940_v1_
            (globalState).f24 = (len(((self).f27 if p1 else (self).f27))) - ((self).f26)
            d_2941_v2_: _dafny.Map
            d_2941_v2_ = _dafny.Map({(self).f26: (self).f27})
            d_2942_v3_: _dafny.Array
            nw464_ = _dafny.Array(None, 8)
            nw464_[int(0)] = ((d_2941_v2_)[(self).f26] if ((self).f26) in (d_2941_v2_) else ((d_2941_v2_)[d_2939_cf0_] if (d_2939_cf0_) in (d_2941_v2_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2943_i0_ in range(default__.abs(102))])))
            nw464_[int(1)] = (self).f27
            nw464_[int(2)] = ((self).f27) + ((self).f27)
            nw464_[int(3)] = (self).f27
            nw464_[int(4)] = (self).f27
            nw464_[int(5)] = (self).f27
            nw464_[int(6)] = (self).f27
            nw464_[int(7)] = (self).f27
            d_2942_v3_ = nw464_
            index491_ = default__.safeIndex(789, (d_2942_v3_).length(0))
            (d_2942_v3_)[index491_] = (D0_DC1((self).f27, 226, 97)).cf1
            index492_ = default__.safeIndex(789, (d_2942_v3_).length(0))
            (d_2942_v3_)[index492_] = ((self).f27) + ((self).f27)
            d_2942_v3_ = d_2942_v3_
        elif source38_.is_DC1:
            d_2944___mcc_h1_ = source38_.cf1
            d_2945___mcc_h2_ = source38_.cf2
            d_2946___mcc_h3_ = source38_.cf3
            d_2947_cf3_ = d_2946___mcc_h3_
            d_2948_cf2_ = d_2945___mcc_h2_
            d_2949_cf1_ = d_2944___mcc_h1_
            d_2950_v4_: _dafny.Map
            d_2950_v4_ = _dafny.Map({not(p1): p1})
            d_2951_v5_: _dafny.Map
            d_2951_v5_ = _dafny.Map({((d_2950_v4_)[p1] if (p1) in (d_2950_v4_) else p1): True})
            d_2952_v6_: _dafny.Map
            d_2952_v6_ = _dafny.Map({p1: d_2948_cf2_})
            d_2951_v5_ = (d_2951_v5_).set(p1, (d_2952_v6_) == (d_2952_v6_))
            (globalState).f0 = (self).f26
            (globalState).f25 = True
            d_2953_v7_: _dafny.Array
            def lambda132_(d_2954_i1_):
                return _dafny.SeqWithoutIsStrInference([True])

            init76_ = lambda132_
            nw465_ = _dafny.Array(None, 29)
            for i0_76_ in range(nw465_.length(0)):
                nw465_[i0_76_] = init76_(i0_76_)
            d_2953_v7_ = nw465_
            d_2955_v8_: _dafny.Seq
            d_2955_v8_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2956_v9_: _dafny.Seq
            d_2956_v9_ = _dafny.SeqWithoutIsStrInference([d_2955_v8_])
            index493_ = default__.safeIndex(619, (d_2953_v7_).length(0))
            (d_2953_v7_)[index493_] = ((d_2956_v9_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(d_2949_cf1_), d_2947_cf3_])), len(d_2956_v9_))]) + (d_2955_v8_)
            index494_ = default__.safeIndex(619, (d_2953_v7_).length(0))
            (d_2953_v7_)[index494_] = _dafny.SeqWithoutIsStrInference([((self).f26) <= (len(_dafny.Map({p1: d_2948_cf2_})))])
        elif source38_.is_DC2:
            d_2957___mcc_h4_ = source38_.cf4
            d_2958_cf4_ = d_2957___mcc_h4_
            d_2959_v10_: _dafny.MultiSet
            d_2959_v10_ = _dafny.MultiSet([(self).f26])
            d_2960_v11_: _dafny.Seq
            d_2960_v11_ = _dafny.SeqWithoutIsStrInference([d_2958_cf4_, (self).fm8(d_2959_v10_, p1, (self).f26, (self).f26, globalState)])
            d_2961_v12_: T1
            nw466_ = C3()
            nw466_.ctor__(len(d_2960_v11_), (self).f26, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2962_i2_ in range(default__.abs(726))]))
            d_2961_v12_ = nw466_
            d_2963_v13_: _dafny.Map
            d_2963_v13_ = _dafny.Map({p1: d_2961_v12_})
            d_2964_v14_: str
            d_2964_v14_ = _dafny.CodePoint('x')
            d_2965_v15_: _dafny.Array
            def lambda133_(d_2966_i3_):
                return (d_2966_i3_) * ((self).f26)

            init77_ = lambda133_
            nw467_ = _dafny.Array(None, 18)
            for i0_77_ in range(nw467_.length(0)):
                nw467_[i0_77_] = init77_(i0_77_)
            d_2965_v15_ = nw467_
            d_2967_v16_: _dafny.Array
            def lambda134_(d_2968_i4_):
                return True

            init78_ = lambda134_
            nw468_ = _dafny.Array(None, 25)
            for i0_78_ in range(nw468_.length(0)):
                nw468_[i0_78_] = init78_(i0_78_)
            d_2967_v16_ = nw468_
            d_2969_v17_: D7
            d_2969_v17_ = D7_DC17(d_2967_v16_)
            d_2970_v18_: _dafny.Seq
            d_2970_v18_ = _dafny.SeqWithoutIsStrInference([d_2969_v17_])
            d_2971_v19_: D7
            d_2971_v19_ = D7_DC21((d_2970_v18_)[default__.safeIndex((self).f26, len(d_2970_v18_))])
            d_2972_v20_: _dafny.Seq
            d_2972_v20_ = _dafny.SeqWithoutIsStrInference([d_2971_v19_])
            d_2973_v21_: _dafny.Set
            d_2973_v21_ = _dafny.Set({d_2972_v20_, _dafny.SeqWithoutIsStrInference([d_2971_v19_])})
            d_2974_v22_: D16
            d_2974_v22_ = D16_DC44(d_2973_v21_)
            d_2975_v23_: _dafny.MultiSet
            d_2975_v23_ = _dafny.MultiSet([d_2974_v22_])
            d_2976_v24_: _dafny.Map
            d_2976_v24_ = _dafny.Map({717: (d_2961_v12_).f26})
            d_2977_v25_: T0
            nw469_ = C4()
            nw469_.ctor__((d_2958_cf4_) and (d_2958_cf4_), (0) - (((d_2975_v23_)[d_2974_v22_] if (d_2974_v22_) in (d_2975_v23_) else (d_2959_v10_).cardinality)), (self).f27, (d_2976_v24_) | (d_2976_v24_))
            d_2977_v25_ = nw469_
            rhs517_ = d_2963_v13_
            rhs518_ = d_2964_v14_
            rhs519_ = p1
            rhs520_ = d_2965_v15_
            rhs521_ = d_2977_v25_
            lhs478_ = globalState
            d_2963_v13_ = rhs517_
            d_2964_v14_ = rhs518_
            lhs478_.f9 = rhs519_
            d_2965_v15_ = rhs520_
            d_2977_v25_ = rhs521_
            d_2978_v26_: _dafny.Map
            d_2978_v26_ = _dafny.Map({d_2964_v14_: d_2960_v11_})
            r1 = (((d_2978_v26_)[((self).f27)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sy"))), len((self).f27))]] if (((self).f27)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sy"))), len((self).f27))]) in (d_2978_v26_) else d_2960_v11_)) == (d_2960_v11_)
            d_2979_v27_: _dafny.Set
            d_2979_v27_ = _dafny.Set({d_2958_cf4_, (self).fm8(d_2959_v10_, p1, 296, (self).f26, globalState), d_2958_cf4_})
            d_2979_v27_ = _dafny.Set({d_2958_cf4_})
            d_2980_v28_: _dafny.Seq
            d_2980_v28_ = _dafny.SeqWithoutIsStrInference([len(d_2960_v11_), (self).f26])
            (globalState).f17 = default__.safeModuloInt((d_2961_v12_).f26, default__.safeModuloInt((self).f26, (d_2980_v28_)[default__.safeIndex((d_2961_v12_).f26, len(d_2980_v28_))]))
        elif source38_.is_DC3:
            d_2981___mcc_h5_ = source38_.cf5
            d_2982___mcc_h6_ = source38_.cf6
            d_2983_cf6_ = d_2982___mcc_h6_
            d_2984_cf5_ = d_2981___mcc_h5_
            (globalState).f25 = p1
            d_2985_v29_: str
            d_2985_v29_ = _dafny.CodePoint('r')
            d_2985_v29_ = d_2985_v29_
            d_2986_v30_: _dafny.Seq
            d_2986_v30_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2987_v31_: D9
            d_2987_v31_ = D9_DC28()
            d_2988_v32_: D9
            d_2988_v32_ = D9_DC29(d_2987_v31_)
            d_2989_v33_: _dafny.Map
            d_2989_v33_ = _dafny.Map({d_2988_v32_: d_2986_v30_})
            d_2986_v30_ = (d_2986_v30_) + (((d_2989_v33_)[d_2988_v32_] if (d_2988_v32_) in (d_2989_v33_) else d_2986_v30_))
            d_2990_v34_: C1
            nw470_ = C1()
            nw470_.ctor__()
            d_2990_v34_ = nw470_
            d_2991_v37_: _dafny.Map
            def iife246_():
                def iife248_():
                    coll132_ = _dafny.Set()
                    compr_132_: int
                    for compr_132_ in _dafny.IntegerRange(823, 793):
                        d_2994_v36_: int = compr_132_
                        if ((823) <= (d_2994_v36_)) and ((d_2994_v36_) < (793)):
                            coll132_ = coll132_.union(_dafny.Set([(d_2994_v36_) * ((self).f26)]))
                    return _dafny.Set(coll132_)
                coll130_ = _dafny.Map()
                def iife247_():
                    coll131_ = _dafny.Set()
                    compr_131_: int
                    for compr_131_ in _dafny.IntegerRange(823, 793):
                        d_2992_v36_: int = compr_131_
                        if ((823) <= (d_2992_v36_)) and ((d_2992_v36_) < (793)):
                            coll131_ = coll131_.union(_dafny.Set([(d_2992_v36_) * ((self).f26)]))
                    return _dafny.Set(coll131_)
                compr_130_: int
                for compr_130_ in (iife247_()
                ).Elements:
                    d_2993_v35_: int = compr_130_
                    if (d_2993_v35_) in (iife248_()
                    ):
                        coll130_[(d_2993_v35_) * (d_2983_cf6_)] = D21_DC59()
                return _dafny.Map(coll130_)
            d_2991_v37_ = _dafny.Map({d_2990_v34_: iife246_()
            })
            d_2995_v38_: _dafny.Map
            d_2995_v38_ = _dafny.Map({len(d_2991_v37_): (self).f26})
            d_2996_v39_: C4
            nw471_ = C4()
            nw471_.ctor__(p1, (self).fm9(d_2983_cf6_, len((self).f27), p1, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wduolc")), d_2995_v38_)
            d_2996_v39_ = nw471_
            d_2997_v40_: _dafny.Map
            d_2997_v40_ = _dafny.Map({d_2996_v39_: p1})
            d_2998_v41_: _dafny.Map
            d_2998_v41_ = _dafny.Map({p1: len(d_2997_v40_)})
            d_2999_v42_: D19
            d_2999_v42_ = D19_DC53(len((self).f27), (d_2996_v39_).f32)
            d_3000_v43_: _dafny.Map
            d_3000_v43_ = _dafny.Map({(d_2999_v42_).cf100: (d_2996_v39_).f32})
            d_3001_v44_: _dafny.Seq
            d_3001_v44_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f26)])
            d_3002_v45_: _dafny.Map
            d_3002_v45_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_2996_v39_).f32, (d_2996_v39_).f32]): (d_2996_v39_).f32})
            (globalState).f9 = default__.fm24(p1, p1, not((((d_2998_v41_)[(d_2996_v39_).f32] if ((d_2996_v39_).f32) in (d_2998_v41_) else 745)) < ((0) - ((self).fm6(d_3000_v43_, (0) - (d_2984_cf5_), d_3001_v44_, d_3002_v45_, globalState)))), 986, globalState)
        elif True:
            d_3003___mcc_h7_ = source38_.cf7
            d_3004_cf7_ = d_3003___mcc_h7_
            if not(p1):
                d_3005_v46_: str
                d_3005_v46_ = _dafny.CodePoint('r')
                d_3006_v47_: _dafny.Set
                d_3006_v47_ = _dafny.Set({False, p1})
                r1 = (not(((self).f27) == (((self).f27).set(default__.safeIndex((self).f26, len((self).f27)), d_3005_v46_)))) and ((d_3006_v47_).isdisjoint(d_3006_v47_))
                d_3007_v48_: _dafny.MultiSet
                d_3007_v48_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emrp"))), (self).f26]))])
                d_3008_v49_: _dafny.Seq
                d_3008_v49_ = _dafny.SeqWithoutIsStrInference([p1, (self).fm8(d_3007_v48_, p1, (self).f26, 170, globalState), p1])
                d_3009_v50_: D2
                d_3009_v50_ = D2_DC7(d_3008_v49_)
                d_3008_v49_ = ((d_3009_v50_).cf12).set(default__.safeIndex((self).f26, len((d_3009_v50_).cf12)), p1)
                d_3010_v51_: _dafny.Seq
                d_3010_v51_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_3011_v53_: T2
                nw472_ = C2()
                def iife249_():
                    coll133_ = _dafny.Map()
                    compr_133_: int
                    for compr_133_ in _dafny.IntegerRange(198, 278):
                        d_3012_v52_: int = compr_133_
                        if ((198) <= (d_3012_v52_)) and ((d_3012_v52_) < (278)):
                            coll133_[(d_3012_v52_) + (len(_dafny.Map({p1: p1})))] = (self).f26
                    return _dafny.Map(coll133_)
                nw472_.ctor__(d_3010_v51_, iife249_()
                )
                d_3011_v53_ = nw472_
                d_3013_v54_: _dafny.Seq
                d_3013_v54_ = _dafny.SeqWithoutIsStrInference([421, 662, (D17_DC48(d_3011_v53_, (self).f26, d_3005_v46_)).cf90, (self).f26])
                d_3014_v56_: C2
                nw473_ = C2()
                def iife250_():
                    coll134_ = _dafny.Map()
                    compr_134_: int
                    for compr_134_ in (d_3010_v51_).Elements:
                        d_3015_v55_: int = compr_134_
                        if (d_3015_v55_) in (d_3010_v51_):
                            coll134_[(d_3015_v55_) - ((self).f26)] = (self).f26
                    return _dafny.Map(coll134_)
                nw473_.ctor__(d_3013_v54_, iife250_()
                )
                d_3014_v56_ = nw473_
                index495_ = default__.safeIndex(292, (p0).length(0))
                (p0)[index495_] = (self).f26
                index496_ = default__.safeIndex(292, (p0).length(0))
                (p0)[index496_] = (self).f26
                d_3016_v57_: _dafny.Array
                nw474_ = _dafny.Array(int(0), 10)
                d_3016_v57_ = nw474_
                d_3017_v58_: _dafny.Seq
                d_3017_v58_ = _dafny.SeqWithoutIsStrInference([d_3006_v47_, d_3006_v47_, _dafny.Set({p1, p1})])
                rhs522_ = d_3016_v57_
                rhs523_ = default__.fm0(p1, _dafny.SeqWithoutIsStrInference([d_3005_v46_ for d_3018_i5_ in range(default__.abs(-641))]), (d_3017_v58_)[default__.safeIndex((p0)[default__.safeIndex(292, (p0).length(0))], len(d_3017_v58_))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wd"))), globalState)
                d_3016_v57_ = rhs522_
                d_3008_v49_ = rhs523_
            elif True:
                d_3019_v59_: _dafny.Map
                d_3019_v59_ = _dafny.Map({p1: D8_DC24((self).f26, (self).f26, (self).f26)})
                d_3019_v59_ = (d_3019_v59_).set(p1, D8_DC24((self).f26, (self).f26, (self).f26))
                d_3020_v60_: _dafny.Map
                d_3020_v60_ = _dafny.Map({(self).f26: (self).f26})
                d_3021_v61_: _dafny.Array
                nw475_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                d_3021_v61_ = nw475_
                d_3022_v62_: _dafny.Array
                nw476_ = _dafny.Array(None, 25)
                nw476_[int(0)] = p1
                nw476_[int(1)] = False
                nw476_[int(2)] = p1
                nw476_[int(3)] = p1
                nw476_[int(4)] = p1
                nw476_[int(5)] = p1
                nw476_[int(6)] = p1
                nw476_[int(7)] = p1
                nw476_[int(8)] = p1
                nw476_[int(9)] = p1
                nw476_[int(10)] = p1
                nw476_[int(11)] = True
                nw476_[int(12)] = True
                nw476_[int(13)] = p1
                nw476_[int(14)] = p1
                nw476_[int(15)] = p1
                nw476_[int(16)] = p1
                nw476_[int(17)] = False
                nw476_[int(18)] = True
                nw476_[int(19)] = p1
                nw476_[int(20)] = p1
                nw476_[int(21)] = p1
                nw476_[int(22)] = p1
                nw476_[int(23)] = p1
                nw476_[int(24)] = p1
                d_3022_v62_ = nw476_
                d_3023_v63_: D13
                d_3023_v63_ = D13_DC40(d_3021_v61_, d_3022_v62_, p1)
                d_3024_v64_: _dafny.MultiSet
                d_3024_v64_ = _dafny.MultiSet([d_3023_v63_])
                d_3025_v65_: _dafny.Seq
                d_3025_v65_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_3026_v66_: _dafny.Map
                d_3026_v66_ = _dafny.Map({len((self).f27): p1})
                d_3027_v67_: _dafny.Map
                d_3027_v67_ = _dafny.Map({len(d_3026_v66_): ((d_3026_v66_)[-222] if (-222) in (d_3026_v66_) else p1)})
                d_3028_v68_: _dafny.Seq
                d_3028_v68_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_3029_v69_: _dafny.Map
                d_3029_v69_ = _dafny.Map({d_3025_v65_: p1})
                d_3030_v70_: C5
                nw477_ = C5()
                nw477_.ctor__(((d_3024_v64_)[d_3023_v63_] if (d_3023_v63_) in (d_3024_v64_) else 787), (_dafny.MultiSet([len(d_3025_v65_)])).cardinality, (self).fm6(d_3026_v66_, (self).f26, d_3028_v68_, d_3029_v69_, globalState), (self).f27)
                d_3030_v70_ = nw477_
                d_3031_v71_: _dafny.Map
                d_3031_v71_ = _dafny.Map({(0) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mliwwul"))), ((d_3020_v60_)[(self).f26] if ((self).f26) in (d_3020_v60_) else (self).f26))): (d_3030_v70_ if p1 else d_3030_v70_)})
                d_3031_v71_ = d_3031_v71_
                d_3032_v72_: _dafny.Map
                d_3032_v72_ = _dafny.Map({p1: d_3030_v70_.f30})
                (globalState).f25 = (len(d_3032_v72_)) != ((0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ug"))) + ((self).f27))))
                d_3033_v73_: _dafny.Map
                d_3033_v73_ = _dafny.Map({(self).f26: (self).f27})
                d_3034_v74_: _dafny.Set
                d_3034_v74_ = _dafny.Set({p1, p1, True})
                d_3035_v75_: _dafny.Map
                d_3035_v75_ = _dafny.Map({352: d_3027_v67_})
                d_3036_v76_: _dafny.MultiSet
                d_3036_v76_ = _dafny.MultiSet([len(((d_3035_v75_)[d_3030_v70_.f29] if (d_3030_v70_.f29) in (d_3035_v75_) else d_3027_v67_))])
                d_3037_v77_: _dafny.Seq
                d_3037_v77_ = _dafny.SeqWithoutIsStrInference([d_3036_v76_])
                rhs524_ = d_3030_v70_.f30
                rhs525_ = len((_dafny.SeqWithoutIsStrInference([p1])) + (default__.fm0(True, ((d_3033_v73_)[(self).f26] if ((self).f26) in (d_3033_v73_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohsnbkvuj"))), d_3034_v74_, d_3030_v70_.f30, globalState)))
                rhs526_ = not(not((d_3030_v70_).fm8(((d_3037_v77_)[default__.safeIndex((0) - ((D10_DC32(p1, -184, p1)).cf63), len(d_3037_v77_))]).set(237, default__.abs(len(_dafny.SeqWithoutIsStrInference([len(d_3020_v60_) for d_3038_i6_ in range(default__.abs(344))])))), p1, d_3030_v70_.f29, d_3030_v70_.f29, globalState)))
                rhs527_ = p1
                lhs479_ = d_3030_v70_
                lhs480_ = globalState
                lhs481_ = globalState
                lhs482_ = globalState
                lhs479_.f29 = rhs524_
                lhs480_.f0 = rhs525_
                lhs481_.f9 = rhs526_
                lhs482_.f9 = rhs527_
                d_3039_v78_: str
                d_3039_v78_ = _dafny.CodePoint('y')
                d_3039_v78_ = d_3039_v78_
            d_3040_v79_: _dafny.Array
            def lambda135_(d_3041_p1_):
                def lambda136_(d_3042_i7_):
                    return d_3041_p1_

                return lambda136_

            init79_ = lambda135_(p1)
            nw478_ = _dafny.Array(None, 4)
            for i0_79_ in range(nw478_.length(0)):
                nw478_[i0_79_] = init79_(i0_79_)
            d_3040_v79_ = nw478_
            index497_ = default__.safeIndex(746, (d_3040_v79_).length(0))
            (d_3040_v79_)[index497_] = p1
            d_3043_v80_: _dafny.MultiSet
            d_3043_v80_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_3044_i8_ in range(default__.abs(245))])])
            index498_ = default__.safeIndex(746, (d_3040_v79_).length(0))
            (d_3040_v79_)[index498_] = ((_dafny.MultiSet([(self).f27, (self).f27])) - (d_3043_v80_)) == (d_3043_v80_)
            d_3045_v81_: _dafny.Seq
            d_3045_v81_ = _dafny.SeqWithoutIsStrInference([p1, p1, (d_3040_v79_)[default__.safeIndex(746, (d_3040_v79_).length(0))]])
            d_3046_v82_: _dafny.Map
            d_3046_v82_ = _dafny.Map({len(d_3045_v81_): (self).f27})
            d_3047_v83_: _dafny.Seq
            d_3047_v83_ = _dafny.SeqWithoutIsStrInference([(self).f26, default__.safeDivisionInt((self).f26, (self).f26), len(d_3046_v82_), ((self).f26) + ((self).f26)])
            d_3047_v83_ = d_3047_v83_
            d_3048_v84_: _dafny.Seq
            d_3048_v84_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dm"))])
            d_3049_v85_: _dafny.Seq
            d_3049_v85_ = _dafny.SeqWithoutIsStrInference([(d_3048_v84_)[default__.safeIndex((self).f26, len(d_3048_v84_))]])
            d_3050_v86_: _dafny.Seq
            d_3050_v86_ = (d_3049_v85_).set(default__.safeIndex((0) - ((self).f26), len(d_3049_v85_)), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_3051_i9_ in range(default__.abs(-221))]))
            source39_ = d_3050_v86_
            d_3052___mcc_h8_ = source39_
            d_3053_cf29_ = d_3052___mcc_h8_
            d_3054_v87_: _dafny.Set
            d_3054_v87_ = _dafny.Set({(self).f26, (self).f26, (self).f26, (0) - ((self).f26)})
            (globalState).f9 = ((d_3054_v87_).intersection(d_3054_v87_)).issubset((_dafny.Set({(self).f26})).intersection(d_3054_v87_))
            d_3055_v88_: C0
            nw479_ = C0()
            nw479_.ctor__()
            d_3055_v88_ = nw479_
            d_3056_v89_: _dafny.MultiSet
            d_3056_v89_ = _dafny.MultiSet([d_3055_v88_, d_3055_v88_, d_3055_v88_, d_3055_v88_])
            (globalState).f25 = (d_3056_v89_).issubset(d_3056_v89_)
            d_3057_v90_: _dafny.Set
            d_3057_v90_ = _dafny.Set({d_3040_v79_})
            d_3058_v91_: D18
            d_3058_v91_ = D18_DC49(d_3057_v90_)
            d_3059_v92_: _dafny.MultiSet
            d_3059_v92_ = _dafny.MultiSet([D18_DC49(d_3057_v90_), D18_DC49(_dafny.Set({d_3040_v79_})), d_3058_v91_, d_3058_v91_])
            (globalState).f9 = (d_3059_v92_).ispropersubset((d_3059_v92_).set(d_3058_v91_, default__.abs(len(d_3045_v81_))))
            (globalState).f9 = (d_3040_v79_)[default__.safeIndex(746, (d_3040_v79_).length(0))]
        if p1:
            d_3060_v93_: _dafny.Set
            d_3060_v93_ = _dafny.Set({(self).f26})
            if (_dafny.Set({(self).f26, (self).f26})).isdisjoint(d_3060_v93_):
                d_3061_v94_: str
                d_3061_v94_ = _dafny.CodePoint('f')
                d_3062_v95_: _dafny.Map
                d_3062_v95_ = _dafny.Map({d_3061_v94_: (self).f26})
                d_3063_v96_: _dafny.Array
                nw480_ = _dafny.Array(_dafny.CodePoint('D'), 13)
                d_3063_v96_ = nw480_
                d_3064_v97_: _dafny.Map
                d_3064_v97_ = _dafny.Map({d_3063_v96_: p1})
                d_3065_v99_: _dafny.Map
                d_3065_v99_ = _dafny.Map({d_3061_v94_: p1})
                d_3066_v100_: _dafny.Array
                nw481_ = _dafny.Array(None, 11)
                nw481_[int(0)] = d_3062_v95_
                nw481_[int(1)] = (d_3062_v95_) | (d_3062_v95_)
                nw481_[int(2)] = (d_3062_v95_).set(d_3061_v94_, (0) - ((self).f26))
                nw481_[int(3)] = d_3062_v95_
                nw481_[int(4)] = d_3062_v95_
                nw481_[int(5)] = d_3062_v95_
                nw481_[int(6)] = (d_3062_v95_) | (default__.fm54(p1, p1, d_3061_v94_, p1, globalState))
                nw481_[int(7)] = d_3062_v95_
                nw481_[int(8)] = d_3062_v95_
                nw481_[int(9)] = default__.fm54(((d_3064_v97_)[d_3063_v96_] if (d_3063_v96_) in (d_3064_v97_) else p1), default__.fm24(p1, p1, p1, (self).f26, globalState), d_3061_v94_, p1, globalState)
                def iife251_():
                    coll135_ = _dafny.Map()
                    compr_135_: str
                    for compr_135_ in (d_3065_v99_).keys.Elements:
                        d_3067_v98_: str = compr_135_
                        if (d_3067_v98_) in (d_3065_v99_):
                            coll135_[d_3067_v98_] = (self).f26
                    return _dafny.Map(coll135_)
                nw481_[int(10)] = iife251_()
                
                d_3066_v100_ = nw481_
                index499_ = default__.safeIndex(905, (d_3066_v100_).length(0))
                def iife252_():
                    coll136_ = _dafny.Map()
                    compr_136_: str
                    for compr_136_ in (default__.fm36(False, p1, p1, globalState)).Elements:
                        d_3068_v101_: str = compr_136_
                        if (d_3068_v101_) in (default__.fm36(False, p1, p1, globalState)):
                            coll136_[d_3068_v101_] = len(_dafny.Set({p1, True}))
                    return _dafny.Map(coll136_)
                (d_3066_v100_)[index499_] = iife252_()
                
                d_3069_v102_: _dafny.Map
                d_3069_v102_ = _dafny.Map({-81: d_2937_v0_})
                d_3070_v103_: D19
                d_3070_v103_ = D19_DC52(d_3069_v102_)
                d_3071_v104_: _dafny.Array
                nw482_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_3071_v104_ = nw482_
                d_3072_v105_: _dafny.Array
                nw483_ = _dafny.Array(False, 12)
                d_3072_v105_ = nw483_
                index500_ = default__.safeIndex(564, (d_3071_v104_).length(0))
                (d_3071_v104_)[index500_] = d_3072_v105_
                d_3073_v106_: _dafny.Map
                d_3073_v106_ = _dafny.Map({(self).f26: (self).f26})
                d_3074_v107_: C4
                nw484_ = C4()
                nw484_.ctor__((p1 if p1 else p1), 232, (self).f27, d_3073_v106_)
                d_3074_v107_ = nw484_
                index501_ = default__.safeIndex(905, (d_3066_v100_).length(0))
                index502_ = default__.safeIndex(564, (d_3071_v104_).length(0))
                rhs528_ = d_3062_v95_
                rhs529_ = D19_DC52(d_3069_v102_)
                rhs530_ = d_3072_v105_
                rhs531_ = d_3074_v107_
                lhs483_ = d_3066_v100_
                lhs484_ = default__.safeIndex(905, (d_3066_v100_).length(0))
                lhs485_ = d_3071_v104_
                lhs486_ = default__.safeIndex(564, (d_3071_v104_).length(0))
                lhs483_[lhs484_] = rhs528_
                d_3070_v103_ = rhs529_
                lhs485_[lhs486_] = rhs530_
                d_3074_v107_ = rhs531_
                (globalState).f8 = (self).f26
                d_3075_v108_: D2
                d_3075_v108_ = D2_DC8((d_3074_v107_).f32, (0) - ((self).f26), (0) - ((self).f26))
                d_3076_v109_: int
                out61_: int
                out61_ = (d_3074_v107_).m4((self).f26, d_3075_v108_, globalState)
                d_3076_v109_ = out61_
                d_3061_v94_ = d_3061_v94_
                d_3077_v110_: C1
                nw485_ = C1()
                nw485_.ctor__()
                d_3077_v110_ = nw485_
            elif True:
                d_3078_v111_: _dafny.Array
                nw486_ = _dafny.Array(_dafny.CodePoint('D'), 25)
                d_3078_v111_ = nw486_
                d_3079_v112_: str
                d_3079_v112_ = _dafny.CodePoint('v')
                index503_ = default__.safeIndex(684, (d_3078_v111_).length(0))
                (d_3078_v111_)[index503_] = d_3079_v112_
                index504_ = default__.safeIndex(684, (d_3078_v111_).length(0))
                (d_3078_v111_)[index504_] = _dafny.CodePoint('u')
                index505_ = default__.safeIndex(154, (p0).length(0))
                (p0)[index505_] = len((self).f27)
                d_3080_v113_: _dafny.Set
                d_3080_v113_ = _dafny.Set({p1})
                d_3081_v114_: _dafny.Seq
                d_3081_v114_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_3082_v115_: _dafny.Seq
                d_3082_v115_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p1, p1}), d_3080_v113_, d_3080_v113_])
                d_3083_v116_: C5
                nw487_ = C5()
                nw487_.ctor__(len((d_3081_v114_) + (d_3081_v114_)), len(default__.fm43(p1, (self).f26, globalState)), len(d_3082_v115_), ((self).f27) + ((self).f27))
                d_3083_v116_ = nw487_
                d_3084_v117_: _dafny.Map
                d_3084_v117_ = _dafny.Map({(self).f26: p1})
                index506_ = default__.safeIndex(154, (p0).length(0))
                index507_ = default__.safeIndex(684, (d_3078_v111_).length(0))
                rhs532_ = len(((self).f27) + ((self).f27))
                rhs533_ = ((d_3084_v117_)[(self).f26] if ((self).f26) in (d_3084_v117_) else p1)
                rhs534_ = d_3080_v113_
                rhs535_ = d_3083_v116_
                rhs536_ = (d_3078_v111_)[default__.safeIndex(684, (d_3078_v111_).length(0))]
                lhs487_ = p0
                lhs488_ = default__.safeIndex(154, (p0).length(0))
                lhs489_ = globalState
                lhs490_ = d_3078_v111_
                lhs491_ = default__.safeIndex(684, (d_3078_v111_).length(0))
                lhs487_[lhs488_] = rhs532_
                lhs489_.f25 = rhs533_
                d_3080_v113_ = rhs534_
                d_3083_v116_ = rhs535_
                lhs490_[lhs491_] = rhs536_
                d_3085_v118_: _dafny.Set
                d_3085_v118_ = _dafny.Set({d_3079_v112_, (d_3078_v111_)[default__.safeIndex(684, (d_3078_v111_).length(0))]})
                (globalState).f25 = not (((self).f26) <= (777)) or (((d_3078_v111_)[default__.safeIndex(684, (d_3078_v111_).length(0))]) in (d_3085_v118_))
                r1 = (((self).f27) + ((self).f27)) < ((self).f27)
                d_3086_v119_: _dafny.Array
                nw488_ = _dafny.Array(False, 21)
                d_3086_v119_ = nw488_
                d_3087_v120_: _dafny.Seq
                d_3087_v120_ = _dafny.SeqWithoutIsStrInference([d_3086_v119_, d_3086_v119_, d_3086_v119_])
                d_3088_v121_: _dafny.Map
                d_3088_v121_ = _dafny.Map({(self).f27: p1})
                d_3089_v122_: _dafny.Map
                d_3089_v122_ = _dafny.Map({((d_3087_v120_)[default__.safeIndex((p0)[default__.safeIndex(154, (p0).length(0))], len(d_3087_v120_))] if True else d_3086_v119_): (d_3088_v121_) | (d_3088_v121_)})
                d_3089_v122_ = (d_3089_v122_).set(d_3086_v119_, (d_3088_v121_) | (d_3088_v121_))
            (globalState).f25 = not ((p1 if p1 else p1)) or (p1)
            (globalState).f24 = 233
            (globalState).f25 = (self).fm8(_dafny.MultiSet([(self).f26, (self).f26]), p1, default__.safeDivisionInt((self).f26, (self).f26), (self).f26, globalState)
            if p1:
                d_3090_v123_: _dafny.Seq
                d_3090_v123_ = _dafny.SeqWithoutIsStrInference([p1])
                d_3091_v124_: _dafny.Map
                d_3091_v124_ = _dafny.Map({(self).f26: p1})
                d_3092_v125_: _dafny.Set
                d_3092_v125_ = _dafny.Set({p1})
                d_3093_v126_: _dafny.Seq
                d_3093_v126_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_3094_v128_: _dafny.Seq
                d_3094_v128_ = _dafny.SeqWithoutIsStrInference([d_3090_v123_])
                def iife253_():
                    coll137_ = _dafny.Map()
                    compr_137_: _dafny.Seq
                    for compr_137_ in (d_3094_v128_).Elements:
                        d_3096_v127_: _dafny.Seq = compr_137_
                        if (d_3096_v127_) in (d_3094_v128_):
                            coll137_[d_3096_v127_] = not(p1)
                    return _dafny.Map(coll137_)
                rhs537_ = (d_3090_v123_ if p1 else default__.fm0(p1, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_3095_i10_ in range(default__.abs(-927))]), default__.fm14(globalState), (self).f26, globalState))
                rhs538_ = (self).fm6((d_3091_v124_ if p1 else _dafny.Map({len(d_3092_v125_): p1})), 871, d_3093_v126_, iife253_()
                , globalState)
                rhs539_ = (d_3090_v123_) + (d_3090_v123_)
                lhs492_ = globalState
                d_3090_v123_ = rhs537_
                lhs492_.f8 = rhs538_
                d_3090_v123_ = rhs539_
                d_3097_v129_: _dafny.Array
                def lambda137_(d_3098_i11_):
                    return _dafny.CodePoint('p')

                init80_ = lambda137_
                nw489_ = _dafny.Array(None, 10)
                for i0_80_ in range(nw489_.length(0)):
                    nw489_[i0_80_] = init80_(i0_80_)
                d_3097_v129_ = nw489_
                d_3099_v130_: str
                d_3099_v130_ = _dafny.CodePoint('o')
                index508_ = default__.safeIndex(930, (d_3097_v129_).length(0))
                (d_3097_v129_)[index508_] = d_3099_v130_
                index509_ = default__.safeIndex(930, (d_3097_v129_).length(0))
                (d_3097_v129_)[index509_] = d_3099_v130_
                (globalState).f25 = (_dafny.SeqWithoutIsStrInference([d_3099_v130_ for d_3100_i12_ in range(default__.abs(77))])) <= ((self).f27)
                index510_ = default__.safeIndex(458, (p0).length(0))
                (p0)[index510_] = (self).f26
                index511_ = default__.safeIndex(458, (p0).length(0))
                rhs540_ = False
                rhs541_ = (0) - (((0) - ((self).f26)) - ((self).f26))
                lhs493_ = globalState
                lhs494_ = p0
                lhs495_ = default__.safeIndex(458, (p0).length(0))
                lhs493_.f25 = rhs540_
                lhs494_[lhs495_] = rhs541_
                d_3101_v131_: _dafny.Map
                d_3101_v131_ = _dafny.Map({p1: p1})
                d_3102_v132_: _dafny.Map
                d_3102_v132_ = _dafny.Map({(d_3093_v126_)[default__.safeIndex((self).f26, len(d_3093_v126_))]: d_3101_v131_})
                d_3103_v133_: _dafny.MultiSet
                d_3103_v133_ = _dafny.MultiSet([p1, p1, p1])
                rhs542_ = p1
                rhs543_ = d_3102_v132_
                rhs544_ = ((p0)[default__.safeIndex(458, (p0).length(0))]) + (((_dafny.MultiSet((d_3094_v128_)[default__.safeIndex((0) - ((self).f26), len(d_3094_v128_))])) - (d_3103_v133_)).cardinality)
                lhs496_ = globalState
                r1 = rhs542_
                d_3102_v132_ = rhs543_
                lhs496_.f8 = rhs544_
            elif True:
                d_3104_v134_: _dafny.Map
                d_3104_v134_ = _dafny.Map({(self).f26: p1})
                d_3105_v135_: _dafny.Seq
                d_3105_v135_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_3106_v136_: _dafny.Seq
                d_3106_v136_ = _dafny.SeqWithoutIsStrInference([True])
                d_3107_v137_: _dafny.Map
                d_3107_v137_ = _dafny.Map({d_3106_v136_: p1})
                (globalState).f17 = (self).fm6(d_3104_v134_, -474, d_3105_v135_, d_3107_v137_, globalState)
                d_3108_v138_: C0
                nw490_ = C0()
                nw490_.ctor__()
                d_3108_v138_ = nw490_
                d_3109_v139_: str
                d_3109_v139_ = _dafny.CodePoint('e')
                d_3110_v140_: _dafny.Array
                nw491_ = _dafny.Array(None, 4)
                nw491_[int(0)] = (self).f27
                nw491_[int(1)] = (self).f27
                nw491_[int(2)] = (self).f27
                nw491_[int(3)] = (((self).f27) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gehqfs")))).set(default__.safeIndex((self).f26, len(((self).f27) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gehqfs"))))), d_3109_v139_)
                d_3110_v140_ = nw491_
                index512_ = default__.safeIndex(317, (d_3110_v140_).length(0))
                (d_3110_v140_)[index512_] = (self).f27
                index513_ = default__.safeIndex(317, (d_3110_v140_).length(0))
                (d_3110_v140_)[index513_] = ((self).f27) + ((self).f27)
                d_3111_v141_: _dafny.Array
                nw492_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_3111_v141_ = nw492_
                index514_ = default__.safeIndex(604, (d_3111_v141_).length(0))
                (d_3111_v141_)[index514_] = p0
                index515_ = default__.safeIndex(604, (d_3111_v141_).length(0))
                (d_3111_v141_)[index515_] = p0
                d_3112_v142_: _dafny.Array
                def lambda138_(d_3113_v135_):
                    def lambda139_(d_3114_i13_):
                        return _dafny.MultiSet([(self).f26, (self).f26, len(d_3113_v135_), 855, 708])

                    return lambda139_

                init81_ = lambda138_(d_3105_v135_)
                nw493_ = _dafny.Array(None, 10)
                for i0_81_ in range(nw493_.length(0)):
                    nw493_[i0_81_] = init81_(i0_81_)
                d_3112_v142_ = nw493_
                d_3112_v142_ = d_3112_v142_
        elif True:
            if p1:
                r1 = p1
                d_3115_v143_: _dafny.Set
                d_3115_v143_ = _dafny.Set({p1, p1})
                d_3116_v144_: _dafny.Set
                d_3116_v144_ = _dafny.Set({d_3115_v143_})
                d_3116_v144_ = d_3116_v144_
                d_3117_v145_: _dafny.Set
                d_3117_v145_ = _dafny.Set({(self).f26, -782})
                d_3118_v146_: _dafny.Seq
                d_3118_v146_ = _dafny.SeqWithoutIsStrInference([p1, p1, not((d_3117_v145_) == (d_3117_v145_)), p1])
                d_3118_v146_ = d_3118_v146_
                d_3119_v147_: _dafny.Array
                nw494_ = _dafny.Array(False, 23)
                d_3119_v147_ = nw494_
                d_3120_v148_: _dafny.Map
                d_3120_v148_ = _dafny.Map({(self).f26: d_3119_v147_})
                d_3121_v149_: _dafny.Set
                d_3121_v149_ = _dafny.Set({d_3119_v147_, d_3119_v147_, ((d_3120_v148_)[(self).f26] if ((self).f26) in (d_3120_v148_) else d_3119_v147_), d_3119_v147_})
                d_3121_v149_ = ((_dafny.Set({d_3119_v147_, d_3119_v147_})).intersection(_dafny.Set({d_3119_v147_})) if False else (d_3121_v149_ if p1 else d_3121_v149_))
                d_3122_v150_: str
                d_3122_v150_ = _dafny.CodePoint('t')
                d_3122_v150_ = d_3122_v150_
            elif True:
                d_3123_v151_: _dafny.Map
                d_3123_v151_ = _dafny.Map({(self).f26: (self).f26})
                d_3124_v152_: C2
                nw495_ = C2()
                nw495_.ctor__((default__.fm55(p1, 328, True, (self).f26, globalState)).cf95, d_3123_v151_)
                d_3124_v152_ = nw495_
                d_3125_v153_: _dafny.Array
                nw496_ = _dafny.Array(_dafny.MultiSet({}), 5)
                d_3125_v153_ = nw496_
                d_3126_v154_: _dafny.Set
                d_3126_v154_ = _dafny.Set({True, p1})
                index516_ = default__.safeIndex(608, (d_3125_v153_).length(0))
                (d_3125_v153_)[index516_] = (_dafny.MultiSet([(self).f27, (self).f27, (self).f27])).intersection(default__.fm56(len(d_3126_v154_), globalState))
                d_3127_v155_: _dafny.MultiSet
                d_3127_v155_ = _dafny.MultiSet([(self).f27])
                index517_ = default__.safeIndex(608, (d_3125_v153_).length(0))
                (d_3125_v153_)[index517_] = d_3127_v155_
                (globalState).f24 = 556
                (globalState).f8 = (self).f26
                d_3128_v156_: _dafny.Seq
                d_3128_v156_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_3129_v157_: _dafny.Seq
                out62_: _dafny.Seq
                out62_ = (d_3124_v152_).m3(not(p1), d_3128_v156_, globalState)
                d_3129_v157_ = out62_
            if p1:
                d_3130_v158_: _dafny.Set
                d_3130_v158_ = _dafny.Set({p1})
                d_3131_v159_: _dafny.MultiSet
                d_3131_v159_ = _dafny.MultiSet([((self).f26) >= (len(d_3130_v158_)), ((self).f27) <= ((self).f27)])
                d_3131_v159_ = d_3131_v159_
                d_3132_v160_: _dafny.Array
                nw497_ = _dafny.Array(int(0), 20)
                d_3132_v160_ = nw497_
                d_3133_v161_: _dafny.Seq
                d_3133_v161_ = _dafny.SeqWithoutIsStrInference([p1])
                d_3134_v162_: _dafny.Seq
                d_3134_v162_ = _dafny.SeqWithoutIsStrInference([d_3133_v161_])
                d_3135_v163_: _dafny.Map
                d_3135_v163_ = _dafny.Map({(self).f26: 944})
                d_3136_v164_: _dafny.Seq
                d_3136_v164_ = _dafny.SeqWithoutIsStrInference([-462])
                d_3137_v165_: str
                d_3137_v165_ = _dafny.CodePoint('s')
                d_3138_v166_: D18
                d_3138_v166_ = D18_DC50((self).f26, p1, d_3136_v164_, d_3137_v165_, (self).f26)
                nw498_ = _dafny.Array(None, 14)
                nw498_[int(0)] = (self).f26
                nw498_[int(1)] = ((self).f26) + (len(_dafny.Map({p1: p1})))
                nw498_[int(2)] = (self).f26
                nw498_[int(3)] = (0) - (default__.safeDivisionInt(len((d_3134_v162_)[default__.safeIndex((self).f26, len(d_3134_v162_))]), (self).f26))
                nw498_[int(4)] = (0) - (len(d_3135_v163_))
                nw498_[int(5)] = 719
                nw498_[int(6)] = (self).f26
                nw498_[int(7)] = ((self).f26) * ((self).f26)
                nw498_[int(8)] = (self).f26
                nw498_[int(9)] = len((d_3138_v166_).cf95)
                nw498_[int(10)] = (self).f26
                nw498_[int(11)] = (self).f26
                nw498_[int(12)] = (-141) + ((self).f26)
                nw498_[int(13)] = (self).fm9((0) - (len((self).f27)), (self).f26, p1, globalState)
                d_3132_v160_ = nw498_
                d_3139_v167_: _dafny.Array
                nw499_ = _dafny.Array(_dafny.Map({}), 3)
                d_3139_v167_ = nw499_
                d_3139_v167_ = d_3139_v167_
                d_3140_v168_: _dafny.Seq
                d_3140_v168_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_3141_v169_: _dafny.Seq
                d_3141_v169_ = (d_3140_v168_) + ((_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_3137_v165_ for d_3142_i14_ in range(default__.abs(946))])), len(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27]))), (self).f27))
                d_3141_v169_ = (_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, _dafny.SeqWithoutIsStrInference([d_3137_v165_ for d_3143_i15_ in range(default__.abs(538))])])) + (_dafny.SeqWithoutIsStrInference([(self).f27]))
                d_3144_v170_: _dafny.Seq
                d_3144_v170_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(p1), False])])
                d_3145_v171_: _dafny.Seq
                d_3145_v171_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([default__.fm37(p1, p1, (self).f26, (self).f26, globalState)])])
                d_3146_v172_: _dafny.Array
                nw500_ = _dafny.Array(None, 28)
                nw500_[int(0)] = d_3144_v170_
                nw500_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_3133_v161_), d_3131_v159_])
                nw500_[int(2)] = (d_3144_v170_) + (d_3144_v170_)
                nw500_[int(3)] = (d_3144_v170_) + (d_3144_v170_)
                nw500_[int(4)] = d_3144_v170_
                nw500_[int(5)] = d_3144_v170_
                nw500_[int(6)] = d_3144_v170_
                nw500_[int(7)] = (d_3144_v170_) + (d_3144_v170_)
                nw500_[int(8)] = d_3144_v170_
                nw500_[int(9)] = ((d_3144_v170_).set(default__.safeIndex((self).f26, len(d_3144_v170_)), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1])))) + (default__.fm57(p1, p1, (self).f26, len(_dafny.SeqWithoutIsStrInference([p1])), globalState))
                nw500_[int(10)] = d_3144_v170_
                nw500_[int(11)] = d_3144_v170_
                nw500_[int(12)] = d_3144_v170_
                nw500_[int(13)] = _dafny.SeqWithoutIsStrInference([d_3131_v159_])
                nw500_[int(14)] = d_3144_v170_
                nw500_[int(15)] = (d_3145_v171_)[default__.safeIndex((self).f26, len(d_3145_v171_))]
                nw500_[int(16)] = (d_3144_v170_) + (default__.fm57(p1, p1, (self).f26, (self).f26, globalState))
                nw500_[int(17)] = (d_3144_v170_) + (_dafny.SeqWithoutIsStrInference([d_3131_v159_]))
                nw500_[int(18)] = d_3144_v170_
                nw500_[int(19)] = d_3144_v170_
                nw500_[int(20)] = d_3144_v170_
                nw500_[int(21)] = d_3144_v170_
                nw500_[int(22)] = _dafny.SeqWithoutIsStrInference([d_3131_v159_ for d_3147_i16_ in range(default__.abs(183))])
                nw500_[int(23)] = d_3144_v170_
                nw500_[int(24)] = (d_3144_v170_).set(default__.safeIndex((self).f26, len(d_3144_v170_)), d_3131_v159_)
                nw500_[int(25)] = (_dafny.SeqWithoutIsStrInference([d_3131_v159_, d_3131_v159_])) + (d_3144_v170_)
                nw500_[int(26)] = (default__.fm57(p1, p1, len((self).f27), (d_3131_v159_).cardinality, globalState)) + (d_3144_v170_)
                nw500_[int(27)] = d_3144_v170_
                d_3146_v172_ = nw500_
                index518_ = default__.safeIndex(985, (d_3146_v172_).length(0))
                (d_3146_v172_)[index518_] = (d_3144_v170_) + (d_3144_v170_)
                index519_ = default__.safeIndex(985, (d_3146_v172_).length(0))
                (d_3146_v172_)[index519_] = d_3144_v170_
            elif True:
                d_3148_v173_: C1
                nw501_ = C1()
                nw501_.ctor__()
                d_3148_v173_ = nw501_
                d_3149_v174_: _dafny.Map
                d_3149_v174_ = _dafny.Map({p1: d_3148_v173_})
                d_3148_v173_ = ((d_3149_v174_)[p1] if (p1) in (d_3149_v174_) else d_3148_v173_)
                d_3150_v175_: _dafny.Seq
                d_3150_v175_ = _dafny.SeqWithoutIsStrInference([(self).f26, len((self).f27)])
                d_3151_v177_: T2
                nw502_ = C2()
                def iife254_():
                    coll138_ = _dafny.Map()
                    compr_138_: int
                    for compr_138_ in _dafny.IntegerRange(87, 478):
                        d_3152_v176_: int = compr_138_
                        if ((87) <= (d_3152_v176_)) and ((d_3152_v176_) < (478)):
                            coll138_[(d_3152_v176_) - (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_3153_i17_ in range(default__.abs(-430))])).set(default__.safeIndex(651, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_3154_i17_ in range(default__.abs(-430))]))), _dafny.CodePoint('t'))))] = len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([p1]))}))
                    return _dafny.Map(coll138_)
                nw502_.ctor__(d_3150_v175_, (iife254_()
                ).set((self).f26, (self).f26))
                d_3151_v177_ = nw502_
                d_3155_v178_: _dafny.Map
                d_3155_v178_ = _dafny.Map({p1: (self).f26})
                nw503_ = C2()
                nw503_.ctor__(default__.fm45((0) - ((self).f26), d_3155_v178_, (d_3151_v177_).f31, globalState), (d_3151_v177_).f31)
                d_3151_v177_ = nw503_
                (globalState).f24 = (self).f26
                r1 = ((self).f26) >= ((self).f26)
                (globalState).f17 = (self).f26
            d_3156_v179_: _dafny.Set
            d_3156_v179_ = _dafny.Set({p1, not(p1)})
            if not((len(d_3156_v179_)) != ((self).fm7((self).f26, p1, globalState))):
                (globalState).f9 = True
                d_3157_v180_: _dafny.Array
                nw504_ = _dafny.Array(_dafny.Map({}), 24)
                d_3157_v180_ = nw504_
                d_3158_v181_: _dafny.Map
                d_3158_v181_ = _dafny.Map({(self).f26: p1})
                index520_ = default__.safeIndex(439, (d_3157_v180_).length(0))
                (d_3157_v180_)[index520_] = d_3158_v181_
                d_3159_v182_: D8
                d_3159_v182_ = D8_DC22(d_3158_v181_)
                index521_ = default__.safeIndex(439, (d_3157_v180_).length(0))
                (d_3157_v180_)[index521_] = (((d_3159_v182_).cf41) | (d_3158_v181_)) | (d_3158_v181_)
                d_3160_v183_: _dafny.Seq
                d_3160_v183_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
                d_3161_v184_: _dafny.Map
                d_3161_v184_ = _dafny.Map({d_3160_v183_: p1})
                d_3162_v185_: str
                d_3162_v185_ = _dafny.CodePoint('a')
                d_3163_v186_: C5
                nw505_ = C5()
                nw505_.ctor__((self).f26, (self).f26, (self).fm6((d_3157_v180_)[default__.safeIndex(439, (d_3157_v180_).length(0))], (self).f26, _dafny.SeqWithoutIsStrInference([(self).f26]), d_3161_v184_, globalState), ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afigneams"))) + ((self).f27)).set(default__.safeIndex((self).f26, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "afigneams"))) + ((self).f27))), d_3162_v185_))
                d_3163_v186_ = nw505_
                d_3160_v183_ = ((d_3160_v183_) + (d_3160_v183_)).set(default__.safeIndex((len((self).f27)) - (d_3163_v186_.f29), len((d_3160_v183_) + (d_3160_v183_))), not(p1))
                (globalState).f24 = -907
            elif True:
                (globalState).f0 = ((0) - (default__.safeDivisionInt((self).f26, len(d_3156_v179_)))) + ((self).f26)
                (globalState).f5 = (self).f26
                d_3164_v187_: _dafny.MultiSet
                d_3164_v187_ = _dafny.MultiSet([not(not(p1)), p1])
                d_3165_v188_: _dafny.Map
                d_3165_v188_ = _dafny.Map({(self).f26: default__.fm24(True, p1, p1, (self).f26, globalState)})
                d_3166_v189_: _dafny.Array
                nw506_ = _dafny.Array(None, 7)
                nw506_[int(0)] = d_3164_v187_
                nw506_[int(1)] = _dafny.MultiSet([p1])
                nw506_[int(2)] = (d_3164_v187_).intersection(d_3164_v187_)
                nw506_[int(3)] = _dafny.MultiSet([((d_3165_v188_)[(self).f26] if ((self).f26) in (d_3165_v188_) else p1)])
                nw506_[int(4)] = (d_3164_v187_).set(True, default__.abs((self).f26))
                nw506_[int(5)] = d_3164_v187_
                nw506_[int(6)] = d_3164_v187_
                d_3166_v189_ = nw506_
                nw507_ = _dafny.Array(None, 4)
                nw507_[int(0)] = d_3164_v187_
                nw507_[int(1)] = d_3164_v187_
                nw507_[int(2)] = _dafny.MultiSet([p1])
                nw507_[int(3)] = d_3164_v187_
                d_3166_v189_ = nw507_
                (globalState).f9 = p1
                d_3167_v190_: _dafny.Seq
                d_3167_v190_ = _dafny.SeqWithoutIsStrInference([p1])
                d_3168_v191_: _dafny.MultiSet
                d_3168_v191_ = _dafny.MultiSet([(self).f26, (self).f26, (self).f26])
                d_3169_v192_: _dafny.Seq
                d_3169_v192_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                (globalState).f0 = (0) - ((self).fm6((d_3165_v188_) | (_dafny.Map({len(d_3167_v190_): (self).fm8(d_3168_v191_, p1, -178, (self).f26, globalState)})), (self).f26, d_3169_v192_, _dafny.Map({d_3167_v190_: p1}), globalState))
            d_3170_v193_: str
            d_3170_v193_ = _dafny.CodePoint('a')
            d_3171_v194_: _dafny.Set
            d_3171_v194_ = _dafny.Set({d_3170_v193_})
            d_3172_v195_: _dafny.Set
            d_3172_v195_ = _dafny.Set({(0) - ((self).f26), len(d_3171_v194_), (self).f26})
            d_3173_v196_: _dafny.Map
            d_3173_v196_ = _dafny.Map({p0: p1})
            d_3174_v197_: _dafny.Map
            d_3174_v197_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_3175_i18_ in range(default__.abs(-834))])) <= ((self).f27): _dafny.Map({default__.fm42(d_3172_v195_, globalState): ((d_3173_v196_)[p0] if (p0) in (d_3173_v196_) else p1)})})
            d_3176_v200_: _dafny.MultiSet
            d_3176_v200_ = _dafny.MultiSet([d_3170_v193_])
            def iife255_():
                def iife257_():
                    coll141_ = _dafny.Map()
                    compr_141_: str
                    for compr_141_ in (d_3176_v200_).Elements:
                        d_3177_v199_: str = compr_141_
                        if (d_3177_v199_) in (d_3176_v200_):
                            coll141_[d_3177_v199_] = (self).f26
                    return _dafny.Map(coll141_)
                coll139_ = _dafny.Map()
                def iife256_():
                    coll140_ = _dafny.Map()
                    compr_140_: str
                    for compr_140_ in (d_3176_v200_).Elements:
                        d_3177_v199_: str = compr_140_
                        if (d_3177_v199_) in (d_3176_v200_):
                            coll140_[d_3177_v199_] = (self).f26
                    return _dafny.Map(coll140_)
                compr_139_: str
                for compr_139_ in (iife256_()
                ).keys.Elements:
                    d_3178_v198_: str = compr_139_
                    if (d_3178_v198_) in (iife257_()
                    ):
                        coll139_[d_3178_v198_] = p1
                return _dafny.Map(coll139_)
            d_3174_v197_ = (d_3174_v197_).set(p1, (iife255_()
            ).set(d_3170_v193_, p1))
            (globalState).f9 = not(p1)
        d_3179_v202_: _dafny.Seq
        d_3179_v202_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        d_3180_v203_: _dafny.Map
        d_3180_v203_ = _dafny.Map({p1: d_3179_v202_})
        d_3181_v204_: _dafny.Map
        d_3181_v204_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p1, p1]): p1})
        def iife258_():
            coll142_ = _dafny.Map()
            compr_142_: int
            for compr_142_ in (d_3179_v202_).Elements:
                d_3182_v201_: int = compr_142_
                if (d_3182_v201_) in (d_3179_v202_):
                    coll142_[default__.safeModuloInt(d_3182_v201_, (self).f26)] = p1
            return _dafny.Map(coll142_)
        r0 = default__.safeModuloInt((self).fm6(iife258_()
        , len(d_3179_v202_), ((d_3180_v203_)[not(p1)] if (not(p1)) in (d_3180_v203_) else d_3179_v202_), d_3181_v204_, globalState), (self).f26)
        d_3183_v205_: _dafny.Array
        nw508_ = _dafny.Array(_dafny.Array(None, 0), 14)
        d_3183_v205_ = nw508_
        d_3184_v206_: _dafny.Array
        nw509_ = _dafny.Array(_dafny.Map({}), 24)
        d_3184_v206_ = nw509_
        index522_ = default__.safeIndex(925, (d_3183_v205_).length(0))
        (d_3183_v205_)[index522_] = d_3184_v206_
        index523_ = default__.safeIndex(925, (d_3183_v205_).length(0))
        (d_3183_v205_)[index523_] = d_3184_v206_
        d_3185_i19_: int
        d_3185_i19_ = 0
        with _dafny.label("6"):
            while p1:
                with _dafny.c_label("6"):
                    if (d_3185_i19_) >= (100):
                        raise _dafny.Break("6")
                    d_3185_i19_ = (d_3185_i19_) + (1)
                    d_3186_v207_: _dafny.MultiSet
                    d_3186_v207_ = _dafny.MultiSet([(self).f27, (self).f27, (self).f27, (self).f27])
                    d_3187_v208_: _dafny.Map
                    d_3187_v208_ = _dafny.Map({d_3186_v207_: (self).f27})
                    d_3187_v208_ = (d_3187_v208_).set(d_3186_v207_, (self).f27)
                    d_3188_v209_: _dafny.Array
                    nw510_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                    d_3188_v209_ = nw510_
                    d_3189_v210_: str
                    d_3189_v210_ = _dafny.CodePoint('e')
                    index524_ = default__.safeIndex(621, (d_3188_v209_).length(0))
                    (d_3188_v209_)[index524_] = default__.fm20(p1, d_3189_v210_, (self).f26, _dafny.CodePoint('l'), globalState)
                    index525_ = default__.safeIndex(621, (d_3188_v209_).length(0))
                    (d_3188_v209_)[index525_] = d_3189_v210_
                    (globalState).f17 = ((self).f26) - ((self).f26)
                    if (367) == ((self).f26):
                        d_3190_v211_: _dafny.Array
                        def lambda140_(d_3191_p1_):
                            def lambda141_(d_3192_i20_):
                                return not (d_3191_p1_) or (d_3191_p1_)

                            return lambda141_

                        init82_ = lambda140_(p1)
                        nw511_ = _dafny.Array(None, 14)
                        for i0_82_ in range(nw511_.length(0)):
                            nw511_[i0_82_] = init82_(i0_82_)
                        d_3190_v211_ = nw511_
                        index526_ = default__.safeIndex(940, (d_3190_v211_).length(0))
                        (d_3190_v211_)[index526_] = p1
                        d_3193_v212_: D0
                        d_3193_v212_ = D0_DC2(p1)
                        index527_ = default__.safeIndex(940, (d_3190_v211_).length(0))
                        (d_3190_v211_)[index527_] = (d_3193_v212_).cf4
                        d_3189_v210_ = d_3189_v210_
                        d_3194_v213_: _dafny.MultiSet
                        d_3194_v213_ = _dafny.MultiSet([(self).f26])
                        d_3195_v214_: _dafny.Map
                        d_3195_v214_ = _dafny.Map({(d_3194_v213_).cardinality: (d_3190_v211_)[default__.safeIndex(940, (d_3190_v211_).length(0))]})
                        d_3196_v215_: D7
                        d_3196_v215_ = D7_DC19(len(d_3195_v214_), p1)
                        d_3197_v216_: D7
                        d_3197_v216_ = D7_DC21(d_3196_v215_)
                        d_3198_v217_: _dafny.Seq
                        d_3198_v217_ = _dafny.SeqWithoutIsStrInference([d_3197_v216_])
                        d_3199_v218_: _dafny.Set
                        d_3199_v218_ = _dafny.Set({d_3198_v217_, d_3198_v217_, d_3198_v217_})
                        d_3200_v219_: D16
                        d_3200_v219_ = D16_DC44((d_3199_v218_) - (d_3199_v218_))
                        d_3200_v219_ = d_3200_v219_
                        (globalState).f25 = ((self).f26) <= (len((self).f27))
                        d_3201_v220_: _dafny.Array
                        def lambda142_(d_3202_p1_):
                            def lambda143_(d_3203_i21_):
                                return ((self).f27 if d_3202_p1_ else (self).f27)

                            return lambda143_

                        init83_ = lambda142_(p1)
                        nw512_ = _dafny.Array(None, 9)
                        for i0_83_ in range(nw512_.length(0)):
                            nw512_[i0_83_] = init83_(i0_83_)
                        d_3201_v220_ = nw512_
                        index528_ = default__.safeIndex(113, (d_3201_v220_).length(0))
                        (d_3201_v220_)[index528_] = ((self).f27) + ((self).f27)
                        index529_ = default__.safeIndex(113, (d_3201_v220_).length(0))
                        (d_3201_v220_)[index529_] = default__.fm43((d_3190_v211_)[default__.safeIndex(940, (d_3190_v211_).length(0))], (self).f26, globalState)
                    elif True:
                        (globalState).f0 = (self).f26
                        (globalState).f9 = p1
                        d_3204_v221_: _dafny.Array
                        nw513_ = _dafny.Array(D18.default()(), 12)
                        d_3204_v221_ = nw513_
                        d_3205_v222_: _dafny.Seq
                        d_3205_v222_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                        d_3206_v223_: _dafny.Map
                        d_3206_v223_ = _dafny.Map({d_3204_v221_: len(d_3205_v222_)})
                        d_3207_v224_: _dafny.Map
                        d_3207_v224_ = _dafny.Map({True: d_3206_v223_})
                        d_3207_v224_ = (d_3207_v224_).set(p1, _dafny.Map({d_3204_v221_: -701}))
                        (globalState).f8 = (0) - (len(((self).f27) + ((self).f27)))
                        (globalState).f25 = ((self).f26) == ((self).f26)
                    pass
            pass
        d_3208_v225_: _dafny.Set
        d_3208_v225_ = _dafny.Set({(self).f26})
        index530_ = default__.safeIndex(379, (p0).length(0))
        (p0)[index530_] = ((0) - ((self).f26)) * (len(d_3208_v225_))
        d_3209_v226_: _dafny.MultiSet
        d_3209_v226_ = _dafny.MultiSet([p1, (p1 if not(p1) else p1), True])
        index531_ = default__.safeIndex(379, (p0).length(0))
        rhs545_ = default__.safeDivisionInt(default__.safeDivisionInt(471, (self).f26), ((self).f26) - ((self).f26))
        rhs546_ = ((d_3209_v226_)[p1] if (p1) in (d_3209_v226_) else (self).f26)
        lhs497_ = p0
        lhs498_ = default__.safeIndex(379, (p0).length(0))
        lhs499_ = globalState
        lhs497_[lhs498_] = rhs545_
        lhs499_.f0 = rhs546_
        r0 = (D20_DC55((self).f26, p1, (self).f26, p1)).cf105
        r1 = (p1 if ((p0)[default__.safeIndex(379, (p0).length(0))]) != ((self).f26) else p1)
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_3210_v0_: _dafny.MultiSet
        d_3210_v0_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f26 for d_3211_i0_ in range(default__.abs(-957))])])
        d_3210_v0_ = d_3210_v0_
        d_3212_v1_: _dafny.Map
        d_3212_v1_ = _dafny.Map({default__.fm24(p0, p0, not(p0), (self).f26, globalState): (self).f26})
        d_3212_v1_ = (d_3212_v1_).set(p0, (self).f26)
        d_3213_i1_: int
        d_3213_i1_ = 0
        with _dafny.label("7"):
            while p0:
                with _dafny.c_label("7"):
                    if (d_3213_i1_) >= (100):
                        raise _dafny.Break("7")
                    d_3213_i1_ = (d_3213_i1_) + (1)
                    (globalState).f5 = (self).f26
                    d_3214_v2_: _dafny.Array
                    nw514_ = _dafny.Array(_dafny.Seq({}), 12)
                    d_3214_v2_ = nw514_
                    d_3215_v3_: _dafny.Array
                    nw515_ = _dafny.Array(_dafny.MultiSet({}), 4)
                    d_3215_v3_ = nw515_
                    d_3216_v4_: D11
                    d_3216_v4_ = D11_DC33(d_3215_v3_)
                    d_3217_v5_: _dafny.Array
                    def lambda144_(d_3218_i2_):
                        return default__.safeModuloInt(d_3218_i2_, (self).f26)

                    init84_ = lambda144_
                    nw516_ = _dafny.Array(None, 29)
                    for i0_84_ in range(nw516_.length(0)):
                        nw516_[i0_84_] = init84_(i0_84_)
                    d_3217_v5_ = nw516_
                    index532_ = default__.safeIndex(328, (d_3217_v5_).length(0))
                    (d_3217_v5_)[index532_] = default__.safeDivisionInt((0) - ((self).f26), (self).f26)
                    d_3219_v6_: _dafny.Array
                    nw517_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
                    d_3219_v6_ = nw517_
                    d_3220_v7_: _dafny.Array
                    def lambda145_(d_3221_p0_):
                        def lambda146_(d_3222_i3_):
                            return d_3221_p0_

                        return lambda146_

                    init85_ = lambda145_(p0)
                    nw518_ = _dafny.Array(None, 7)
                    for i0_85_ in range(nw518_.length(0)):
                        nw518_[i0_85_] = init85_(i0_85_)
                    d_3220_v7_ = nw518_
                    d_3223_v8_: D13
                    d_3223_v8_ = D13_DC40(d_3219_v6_, d_3220_v7_, p0)
                    pat_let_tv72_ = d_3220_v7_
                    pat_let_tv73_ = d_3219_v6_
                    d_3224_v9_: _dafny.Seq
                    def iife259_(_pat_let58_0):
                        def iife260_(d_3225_dt__update__tmp_h0_):
                            def iife261_(_pat_let59_0):
                                def iife262_(d_3226_dt__update_hcf77_h0_):
                                    def iife263_(_pat_let60_0):
                                        def iife264_(d_3227_dt__update_hcf76_h0_):
                                            return D13_DC40(d_3227_dt__update_hcf76_h0_, d_3226_dt__update_hcf77_h0_, (d_3225_dt__update__tmp_h0_).cf78)
                                        return iife264_(_pat_let60_0)
                                    return iife263_(pat_let_tv73_)
                                return iife262_(_pat_let59_0)
                            return iife261_(pat_let_tv72_)
                        return iife260_(_pat_let58_0)
                    d_3224_v9_ = _dafny.SeqWithoutIsStrInference([not(True), not(p0), (iife259_(d_3223_v8_)).cf78])
                    index533_ = default__.safeIndex(328, (d_3217_v5_).length(0))
                    rhs547_ = (d_3224_v9_)[default__.safeIndex((self).f26, len(d_3224_v9_))]
                    rhs548_ = d_3214_v2_
                    rhs549_ = d_3216_v4_
                    rhs550_ = (self).f26
                    rhs551_ = (d_3212_v1_ if p0 else _dafny.Map({p0: (self).f26}))
                    lhs500_ = globalState
                    lhs501_ = d_3217_v5_
                    lhs502_ = default__.safeIndex(328, (d_3217_v5_).length(0))
                    lhs500_.f9 = rhs547_
                    d_3214_v2_ = rhs548_
                    d_3216_v4_ = rhs549_
                    lhs501_[lhs502_] = rhs550_
                    d_3212_v1_ = rhs551_
                    (globalState).f24 = len((self).f27)
                    d_3228_v10_: _dafny.Map
                    d_3228_v10_ = _dafny.Map({(self).f26: True})
                    (globalState).f5 = default__.safeDivisionInt((d_3217_v5_)[default__.safeIndex(328, (d_3217_v5_).length(0))], len((_dafny.Map({463: p0})) | (d_3228_v10_)))
                    pass
            pass
        d_3229_i4_: int
        d_3229_i4_ = 0
        with _dafny.label("8"):
            while p0:
                with _dafny.c_label("8"):
                    if (d_3229_i4_) >= (100):
                        raise _dafny.Break("8")
                    d_3229_i4_ = (d_3229_i4_) + (1)
                    (globalState).f9 = p0
                    if ((self).f26) > (((self).f26) - ((self).f26)):
                        d_3230_v11_: _dafny.Seq
                        d_3230_v11_ = _dafny.SeqWithoutIsStrInference([(self).f26, (0) - ((self).f26)])
                        d_3231_v12_: _dafny.Map
                        d_3231_v12_ = _dafny.Map({(p0 if p0 else p0): (p0 if not(False) else (self).fm8(_dafny.MultiSet(d_3230_v11_), p0, (self).fm9((self).f26, (self).f26, p0, globalState), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_3232_i5_ in range(default__.abs(-711))])), globalState))})
                        d_3231_v12_ = ((d_3231_v12_) | (_dafny.Map({(p1).cf4: p0}))) | ((d_3231_v12_) | (_dafny.Map({p0: not(p0)})))
                        (globalState).f17 = (self).f26
                        d_3233_v13_: _dafny.Array
                        def lambda147_(d_3234_i6_):
                            return (d_3234_i6_) + ((self).f26)

                        init86_ = lambda147_
                        nw519_ = _dafny.Array(None, 14)
                        for i0_86_ in range(nw519_.length(0)):
                            nw519_[i0_86_] = init86_(i0_86_)
                        d_3233_v13_ = nw519_
                        index534_ = default__.safeIndex(407, (d_3233_v13_).length(0))
                        (d_3233_v13_)[index534_] = (self).f26
                        d_3235_v14_: str
                        d_3235_v14_ = _dafny.CodePoint('k')
                        d_3236_v15_: _dafny.Map
                        d_3236_v15_ = _dafny.Map({d_3235_v14_: (self).f27})
                        d_3237_v16_: _dafny.MultiSet
                        d_3237_v16_ = _dafny.MultiSet([(self).f27, (self).f27, ((self).f27) + (((d_3236_v15_)[d_3235_v14_] if (d_3235_v14_) in (d_3236_v15_) else (self).f27)), (self).f27])
                        index535_ = default__.safeIndex(407, (d_3233_v13_).length(0))
                        (d_3233_v13_)[index535_] = ((d_3237_v16_)[_dafny.SeqWithoutIsStrInference([((self).f27)[default__.safeIndex((self).f26, len((self).f27))] for d_3238_i7_ in range(default__.abs(-378))])] if (_dafny.SeqWithoutIsStrInference([((self).f27)[default__.safeIndex((self).f26, len((self).f27))] for d_3239_i7_ in range(default__.abs(-378))])) in (d_3237_v16_) else (self).f26)
                        (globalState).f13 = ((self).f27) + (((self).f27) + ((self).f27))
                        d_3240_v17_: _dafny.Array
                        nw520_ = _dafny.Array(D11.default()(), 1)
                        d_3240_v17_ = nw520_
                        d_3241_v18_: _dafny.Array
                        nw521_ = _dafny.Array(False, 6)
                        d_3241_v18_ = nw521_
                        d_3242_v19_: D11
                        d_3242_v19_ = D11_DC34(d_3241_v18_, p0)
                        index536_ = default__.safeIndex(83, (d_3240_v17_).length(0))
                        (d_3240_v17_)[index536_] = d_3242_v19_
                        index537_ = default__.safeIndex(776, (d_3241_v18_).length(0))
                        (d_3241_v18_)[index537_] = True
                        d_3243_v20_: _dafny.MultiSet
                        d_3243_v20_ = _dafny.MultiSet([True, (p0) or (p0), p0])
                        d_3244_v22_: _dafny.MultiSet
                        d_3244_v22_ = _dafny.MultiSet([len((self).f27)])
                        d_3245_v23_: _dafny.Map
                        d_3245_v23_ = _dafny.Map({(self).f26: (d_3244_v22_).cardinality})
                        d_3246_v24_: _dafny.Seq
                        d_3246_v24_ = _dafny.SeqWithoutIsStrInference([(d_3245_v23_).set(len((self).f27), (d_3233_v13_)[default__.safeIndex(407, (d_3233_v13_).length(0))])])
                        index538_ = default__.safeIndex(83, (d_3240_v17_).length(0))
                        index539_ = default__.safeIndex(776, (d_3241_v18_).length(0))
                        def iife265_():
                            coll143_ = _dafny.Map()
                            compr_143_: int
                            for compr_143_ in _dafny.IntegerRange(252, 324):
                                d_3247_v21_: int = compr_143_
                                if ((252) <= (d_3247_v21_)) and ((d_3247_v21_) < (324)):
                                    coll143_[default__.safeDivisionInt(d_3247_v21_, (D3_DC12(p0, (d_3233_v13_)[default__.safeIndex(407, (d_3233_v13_).length(0))], p0)).cf26)] = (d_3233_v13_)[default__.safeIndex(407, (d_3233_v13_).length(0))]
                            return _dafny.Map(coll143_)
                        rhs552_ = d_3242_v19_
                        rhs553_ = ((_dafny.SeqWithoutIsStrInference([iife265_()
 for d_3248_i8_ in range(default__.abs(593))])) + (d_3246_v24_)) <= (d_3246_v24_)
                        rhs554_ = not(p0)
                        rhs555_ = False
                        rhs556_ = _dafny.MultiSet([p0])
                        lhs503_ = d_3240_v17_
                        lhs504_ = default__.safeIndex(83, (d_3240_v17_).length(0))
                        lhs505_ = d_3241_v18_
                        lhs506_ = default__.safeIndex(776, (d_3241_v18_).length(0))
                        lhs507_ = globalState
                        lhs508_ = globalState
                        lhs503_[lhs504_] = rhs552_
                        lhs505_[lhs506_] = rhs553_
                        lhs507_.f25 = rhs554_
                        lhs508_.f9 = rhs555_
                        d_3243_v20_ = rhs556_
                    elif True:
                        d_3249_v25_: _dafny.Map
                        d_3249_v25_ = _dafny.Map({p0: p0})
                        d_3249_v25_ = (d_3249_v25_).set(False, p0)
                        (globalState).f5 = (self).f26
                        d_3250_v26_: _dafny.Map
                        d_3250_v26_ = _dafny.Map({(self).f26: D12_DC38()})
                        d_3251_v27_: D12
                        d_3251_v27_ = D12_DC38()
                        d_3250_v26_ = (d_3250_v26_).set((self).f26, d_3251_v27_)
                        d_3252_v28_: _dafny.Map
                        d_3252_v28_ = _dafny.Map({p0: (self).f27})
                        d_3252_v28_ = (d_3252_v28_).set(False, ((d_3252_v28_)[p0] if (p0) in (d_3252_v28_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_3253_i9_ in range(default__.abs(413))])))
                        d_3254_v29_: _dafny.Array
                        nw522_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
                        d_3254_v29_ = nw522_
                        index540_ = default__.safeIndex(729, (d_3254_v29_).length(0))
                        (d_3254_v29_)[index540_] = (self).f27
                        d_3255_v30_: _dafny.Map
                        d_3255_v30_ = _dafny.Map({(self).f27: p0})
                        d_3256_v31_: _dafny.Seq
                        d_3256_v31_ = _dafny.SeqWithoutIsStrInference([len(d_3212_v1_)])
                        d_3257_v32_: _dafny.Map
                        d_3257_v32_ = _dafny.Map({(self).f26: 852})
                        d_3258_v33_: C2
                        nw523_ = C2()
                        nw523_.ctor__(d_3256_v31_, d_3257_v32_)
                        d_3258_v33_ = nw523_
                        d_3259_v34_: D2
                        d_3259_v34_ = D2_DC9((self).f26, p0, (self).f26, (self).f26)
                        d_3260_v35_: _dafny.Map
                        d_3260_v35_ = _dafny.Map({d_3258_v33_: d_3259_v34_})
                        index541_ = default__.safeIndex(729, (d_3254_v29_).length(0))
                        rhs557_ = ((d_3255_v30_)[(self).f27] if ((self).f27) in (d_3255_v30_) else not(p0))
                        rhs558_ = not (p0) or (p0)
                        rhs559_ = ((self).f27) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fn"))).set(default__.safeIndex((self).f26, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fn")))), _dafny.CodePoint('d')))
                        rhs560_ = len((_dafny.Map({d_3258_v33_: d_3259_v34_})) | (d_3260_v35_))
                        lhs509_ = globalState
                        lhs510_ = globalState
                        lhs511_ = d_3254_v29_
                        lhs512_ = default__.safeIndex(729, (d_3254_v29_).length(0))
                        lhs513_ = globalState
                        lhs509_.f25 = rhs557_
                        lhs510_.f25 = rhs558_
                        lhs511_[lhs512_] = rhs559_
                        lhs513_.f8 = rhs560_
                    if p0:
                        d_3261_v36_: _dafny.Array
                        nw524_ = _dafny.Array(False, 24)
                        d_3261_v36_ = nw524_
                        d_3262_v37_: str
                        d_3262_v37_ = _dafny.CodePoint('r')
                        index542_ = default__.safeIndex(160, (d_3261_v36_).length(0))
                        (d_3261_v36_)[index542_] = (d_3262_v37_) in ((self).f27)
                        index543_ = default__.safeIndex(160, (d_3261_v36_).length(0))
                        (d_3261_v36_)[index543_] = p0
                        d_3263_v38_: _dafny.Map
                        d_3263_v38_ = _dafny.Map({(self).f26: p0})
                        d_3264_v39_: _dafny.MultiSet
                        d_3264_v39_ = _dafny.MultiSet([p0, p0])
                        d_3265_v40_: _dafny.Seq
                        d_3265_v40_ = _dafny.SeqWithoutIsStrInference([((d_3264_v39_)[(d_3261_v36_)[default__.safeIndex(160, (d_3261_v36_).length(0))]] if ((d_3261_v36_)[default__.safeIndex(160, (d_3261_v36_).length(0))]) in (d_3264_v39_) else len(d_3212_v1_))])
                        d_3266_v41_: _dafny.Map
                        d_3266_v41_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0]): p0})
                        d_3267_v42_: _dafny.Seq
                        d_3267_v42_ = _dafny.SeqWithoutIsStrInference([not((d_3261_v36_)[default__.safeIndex(160, (d_3261_v36_).length(0))])])
                        (globalState).f9 = ((self).f26) <= ((self).fm6(d_3263_v38_, (self).f26, d_3265_v40_, (d_3266_v41_).set(d_3267_v42_, p0), globalState))
                        d_3268_v43_: C3
                        nw525_ = C3()
                        nw525_.ctor__((self).f26, 337, (self).f27)
                        d_3268_v43_ = nw525_
                        d_3269_v44_: D23
                        d_3269_v44_ = D23_DC65(d_3268_v43_)
                        d_3270_v45_: _dafny.Seq
                        d_3270_v45_ = _dafny.SeqWithoutIsStrInference([d_3268_v43_])
                        d_3271_v46_: _dafny.Array
                        nw526_ = _dafny.Array(None, 27)
                        nw526_[int(0)] = d_3268_v43_
                        nw526_[int(1)] = d_3268_v43_
                        nw526_[int(2)] = d_3268_v43_
                        nw526_[int(3)] = d_3268_v43_
                        nw526_[int(4)] = d_3268_v43_
                        nw526_[int(5)] = (d_3269_v44_).cf122
                        nw526_[int(6)] = d_3268_v43_
                        nw526_[int(7)] = d_3268_v43_
                        nw526_[int(8)] = d_3268_v43_
                        nw526_[int(9)] = d_3268_v43_
                        nw526_[int(10)] = (d_3268_v43_ if False else d_3268_v43_)
                        nw526_[int(11)] = d_3268_v43_
                        nw526_[int(12)] = d_3268_v43_
                        nw526_[int(13)] = d_3268_v43_
                        nw526_[int(14)] = d_3268_v43_
                        nw526_[int(15)] = d_3268_v43_
                        nw526_[int(16)] = d_3268_v43_
                        nw526_[int(17)] = d_3268_v43_
                        nw526_[int(18)] = d_3268_v43_
                        nw526_[int(19)] = d_3268_v43_
                        nw526_[int(20)] = d_3268_v43_
                        nw526_[int(21)] = (d_3270_v45_)[default__.safeIndex((0) - ((self).f26), len(d_3270_v45_))]
                        nw526_[int(22)] = d_3268_v43_
                        nw526_[int(23)] = d_3268_v43_
                        nw526_[int(24)] = (d_3270_v45_)[default__.safeIndex((self).f26, len(d_3270_v45_))]
                        nw526_[int(25)] = d_3268_v43_
                        nw526_[int(26)] = (d_3270_v45_)[default__.safeIndex((self).f26, len(d_3270_v45_))]
                        d_3271_v46_ = nw526_
                        index544_ = default__.safeIndex(535, (d_3271_v46_).length(0))
                        (d_3271_v46_)[index544_] = d_3268_v43_
                        index545_ = default__.safeIndex(535, (d_3271_v46_).length(0))
                        (d_3271_v46_)[index545_] = d_3268_v43_
                        d_3272_v47_: _dafny.Set
                        d_3272_v47_ = _dafny.Set({(d_3264_v39_).cardinality})
                        d_3273_v48_: _dafny.MultiSet
                        d_3273_v48_ = _dafny.MultiSet([d_3272_v47_, d_3272_v47_, d_3272_v47_])
                        d_3274_v49_: C5
                        nw527_ = C5()
                        nw527_.ctor__(((self).f26) + ((d_3268_v43_).f34), (0) - ((_dafny.MultiSet([(self).f26, (d_3268_v43_).f34, (self).f26, ((d_3273_v48_).set(d_3272_v47_, default__.abs((d_3268_v43_).f34))).cardinality])).cardinality), (self).f26, (self).f27)
                        d_3274_v49_ = nw527_
                        d_3275_v50_: _dafny.Array
                        def lambda148_(d_3276_i10_):
                            return default__.safeModuloInt(d_3276_i10_, len((self).f27))

                        init87_ = lambda148_
                        nw528_ = _dafny.Array(None, 23)
                        for i0_87_ in range(nw528_.length(0)):
                            nw528_[i0_87_] = init87_(i0_87_)
                        d_3275_v50_ = nw528_
                        index546_ = default__.safeIndex(500, (d_3275_v50_).length(0))
                        (d_3275_v50_)[index546_] = len((_dafny.SeqWithoutIsStrInference([(d_3268_v43_).f34])) + (d_3265_v40_))
                        index547_ = default__.safeIndex(500, (d_3275_v50_).length(0))
                        (d_3275_v50_)[index547_] = (d_3268_v43_).f34
                    elif True:
                        d_3277_v51_: _dafny.Seq
                        d_3277_v51_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f26), (self).f26, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktwuwhbfj")))])
                        (globalState).f8 = (d_3277_v51_)[default__.safeIndex((self).f26, len(d_3277_v51_))]
                        d_3278_v52_: _dafny.Array
                        nw529_ = _dafny.Array(_dafny.MultiSet({}), 8)
                        d_3278_v52_ = nw529_
                        d_3279_v53_: _dafny.MultiSet
                        d_3279_v53_ = _dafny.MultiSet([p0, p0])
                        index548_ = default__.safeIndex(804, (d_3278_v52_).length(0))
                        (d_3278_v52_)[index548_] = d_3279_v53_
                        index549_ = default__.safeIndex(804, (d_3278_v52_).length(0))
                        (d_3278_v52_)[index549_] = default__.fm37(p0, p0, ((self).f26) - ((self).f26), len(default__.fm36(p0, p0, p0, globalState)), globalState)
                        d_3280_v54_: _dafny.Array
                        def lambda149_(d_3281_p0_):
                            def lambda150_(d_3282_i11_):
                                return (D1_DC6(d_3281_p0_, _dafny.CodePoint('g'), (self).f26)).cf9

                            return lambda150_

                        init88_ = lambda149_(p0)
                        nw530_ = _dafny.Array(None, 3)
                        for i0_88_ in range(nw530_.length(0)):
                            nw530_[i0_88_] = init88_(i0_88_)
                        d_3280_v54_ = nw530_
                        index550_ = default__.safeIndex(723, (d_3280_v54_).length(0))
                        (d_3280_v54_)[index550_] = p0
                        d_3283_v55_: _dafny.Seq
                        d_3283_v55_ = _dafny.SeqWithoutIsStrInference([False])
                        d_3284_v56_: _dafny.Map
                        d_3284_v56_ = _dafny.Map({(self).f26: _dafny.SeqWithoutIsStrInference([p0])})
                        index551_ = default__.safeIndex(804, (d_3278_v52_).length(0))
                        index552_ = default__.safeIndex(723, (d_3280_v54_).length(0))
                        rhs561_ = (((self).f27) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "duqjf")))) + ((self).f27)
                        rhs562_ = _dafny.MultiSet(d_3283_v55_)
                        rhs563_ = True
                        rhs564_ = (d_3283_v55_) + (((d_3284_v56_)[(self).f26] if ((self).f26) in (d_3284_v56_) else d_3283_v55_))
                        rhs565_ = d_3280_v54_
                        lhs514_ = globalState
                        lhs515_ = d_3278_v52_
                        lhs516_ = default__.safeIndex(804, (d_3278_v52_).length(0))
                        lhs517_ = d_3280_v54_
                        lhs518_ = default__.safeIndex(723, (d_3280_v54_).length(0))
                        lhs519_ = globalState
                        lhs514_.f23 = rhs561_
                        lhs515_[lhs516_] = rhs562_
                        lhs517_[lhs518_] = rhs563_
                        d_3283_v55_ = rhs564_
                        lhs519_.f1 = rhs565_
                        d_3285_v57_: D2
                        d_3285_v57_ = D2_DC8(p0, (self).f26, 202)
                        d_3286_v58_: _dafny.Array
                        nw531_ = _dafny.Array(None, 1)
                        d_3286_v58_ = nw531_
                        d_3287_v59_: _dafny.Map
                        def iife266_(_pat_let61_0):
                            def iife267_(d_3288_dt__update__tmp_h1_):
                                def iife268_(_pat_let62_0):
                                    def iife269_(d_3289_dt__update_hcf15_h0_):
                                        return D2_DC8((d_3288_dt__update__tmp_h1_).cf13, (d_3288_dt__update__tmp_h1_).cf14, d_3289_dt__update_hcf15_h0_)
                                    return iife269_(_pat_let62_0)
                                return iife268_(len((self).f27))
                            return iife267_(_pat_let61_0)
                        d_3287_v59_ = _dafny.Map({iife266_(d_3285_v57_): d_3286_v58_})
                        rhs566_ = d_3287_v59_
                        rhs567_ = p0
                        lhs520_ = globalState
                        d_3287_v59_ = rhs566_
                        lhs520_.f25 = rhs567_
                        (globalState).f25 = (d_3280_v54_)[default__.safeIndex(723, (d_3280_v54_).length(0))]
                    (globalState).f9 = ((self).f26) > ((self).f26)
                    pass
            pass
        d_3290_v60_: C0
        nw532_ = C0()
        nw532_.ctor__()
        d_3290_v60_ = nw532_
        d_3291_v61_: _dafny.Map
        d_3291_v61_ = _dafny.Map({d_3290_v60_: default__.fm43(p0, (self).f26, globalState)})
        d_3292_v62_: D24
        d_3292_v62_ = D24_DC70(d_3290_v60_)
        d_3291_v61_ = (d_3291_v61_).set((d_3292_v62_).cf132, (self).f27)
        d_3293_v63_: _dafny.Array
        nw533_ = _dafny.Array(_dafny.Array(None, 0), 13)
        d_3293_v63_ = nw533_
        d_3294_v65_: _dafny.Array
        def lambda151_(d_3295_i12_):
            def iife270_():
                coll144_ = _dafny.Map()
                compr_144_: int
                for compr_144_ in (_dafny.MultiSet([(self).f26, (self).f26])).Elements:
                    d_3296_v64_: int = compr_144_
                    if (d_3296_v64_) in (_dafny.MultiSet([(self).f26, (self).f26])):
                        coll144_[(d_3296_v64_) - ((self).f26)] = _dafny.CodePoint('n')
                return _dafny.Map(coll144_)
            return (d_3295_i12_) + (len(iife270_()
            ))

        init89_ = lambda151_
        nw534_ = _dafny.Array(None, 1)
        for i0_89_ in range(nw534_.length(0)):
            nw534_[i0_89_] = init89_(i0_89_)
        d_3294_v65_ = nw534_
        d_3297_v66_: _dafny.Seq
        d_3297_v66_ = _dafny.SeqWithoutIsStrInference([d_3294_v65_, d_3294_v65_, d_3294_v65_])
        d_3298_v67_: _dafny.MultiSet
        d_3298_v67_ = _dafny.MultiSet([p0])
        index553_ = default__.safeIndex(388, (d_3293_v63_).length(0))
        (d_3293_v63_)[index553_] = (D7_DC18(446, (d_3297_v66_)[default__.safeIndex(((d_3298_v67_).set(p0, default__.abs((self).f26))).cardinality, len(d_3297_v66_))])).cf35
        index554_ = default__.safeIndex(388, (d_3293_v63_).length(0))
        (d_3293_v63_)[index554_] = d_3294_v65_
        r0 = default__.safeModuloInt((0) - ((self).f26), ((self).f26) + (400))
        d_3299_v68_: D0
        d_3299_v68_ = D0_DC0((self).f26)
        d_3300_v69_: _dafny.MultiSet
        def iife271_(_pat_let63_0):
            def iife272_(d_3301_dt__update__tmp_h2_):
                def iife273_(_pat_let64_0):
                    def iife274_(d_3302_dt__update_hcf0_h0_):
                        return D0_DC0(d_3302_dt__update_hcf0_h0_)
                    return iife274_(_pat_let64_0)
                return iife273_((self).f26)
            return iife272_(_pat_let63_0)
        d_3300_v69_ = _dafny.MultiSet([d_3299_v68_, iife271_(d_3299_v68_)])
        r1 = ((d_3300_v69_) | ((d_3300_v69_).set(D0_DC0((self).f26), default__.abs(5)))).intersection((d_3300_v69_).intersection(d_3300_v69_))
        d_3303_v70_: _dafny.Array
        def lambda152_(d_3304_p0_):
            def lambda153_(d_3305_i13_):
                return d_3304_p0_

            return lambda153_

        init90_ = lambda152_(p0)
        nw535_ = _dafny.Array(None, 23)
        for i0_90_ in range(nw535_.length(0)):
            nw535_[i0_90_] = init90_(i0_90_)
        d_3303_v70_ = nw535_
        d_3306_v71_: D7
        d_3306_v71_ = D7_DC17(d_3303_v70_)
        pat_let_tv74_ = d_3303_v70_
        nw536_ = _dafny.Array(None, 18)
        nw536_[int(0)] = d_3294_v65_
        nw536_[int(1)] = (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))]
        nw536_[int(2)] = (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))]
        nw536_[int(3)] = (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))]
        nw536_[int(4)] = d_3294_v65_
        nw536_[int(5)] = d_3294_v65_
        nw536_[int(6)] = (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))]
        def iife275_(_pat_let65_0):
            def iife276_(d_3307_dt__update__tmp_h3_):
                def iife277_(_pat_let66_0):
                    def iife278_(d_3308_dt__update_hcf33_h0_):
                        return D7_DC17(d_3308_dt__update_hcf33_h0_)
                    return iife278_(_pat_let66_0)
                return iife277_(pat_let_tv74_)
            return iife276_(_pat_let65_0)
        nw536_[int(7)] = (D23_DC67(p0, iife275_(d_3306_v71_), (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))], (self).f26, len(_dafny.Map({(self).f26: (self).f26})))).cf126
        nw536_[int(8)] = (d_3297_v66_)[default__.safeIndex((self).f26, len(d_3297_v66_))]
        nw536_[int(9)] = d_3294_v65_
        nw536_[int(10)] = (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))]
        nw536_[int(11)] = d_3294_v65_
        nw536_[int(12)] = (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))]
        nw536_[int(13)] = d_3294_v65_
        nw536_[int(14)] = (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))]
        nw536_[int(15)] = (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))]
        nw536_[int(16)] = (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))]
        nw536_[int(17)] = (d_3293_v63_)[default__.safeIndex(388, (d_3293_v63_).length(0))]
        r2 = nw536_
        return r0, r1, r2

    @property
    def f28(self):
        return self._f28

class C7:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def m0(self, p0, p1, p2, p3, globalState):
        d_3309_v0_: _dafny.Array
        def lambda154_(d_3310_p2_):
            def lambda155_(d_3311_i0_):
                return d_3310_p2_

            return lambda155_

        init91_ = lambda154_(p2)
        nw537_ = _dafny.Array(None, 23)
        for i0_91_ in range(nw537_.length(0)):
            nw537_[i0_91_] = init91_(i0_91_)
        d_3309_v0_ = nw537_
        index555_ = default__.safeIndex(971, (d_3309_v0_).length(0))
        (d_3309_v0_)[index555_] = (p0 if p0 else p0)
        d_3312_v1_: _dafny.Seq
        d_3312_v1_ = _dafny.SeqWithoutIsStrInference([p2])
        index556_ = default__.safeIndex(971, (d_3309_v0_).length(0))
        (d_3309_v0_)[index556_] = (_dafny.SeqWithoutIsStrInference([True, p0, p2, not(p0), p0])) < (d_3312_v1_)
        d_3313_v2_: D0
        d_3313_v2_ = D0_DC3(p1, p1)
        source40_ = d_3313_v2_
        if source40_.is_DC0:
            d_3314___mcc_h0_ = source40_.cf0
            d_3315_cf0_ = d_3314___mcc_h0_
            (globalState).f8 = p1
            d_3316_v3_: D0
            d_3316_v3_ = D0_DC2((d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))])
            d_3317_v4_: _dafny.Map
            d_3317_v4_ = _dafny.Map({(d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]: p0})
            d_3318_v5_: _dafny.Set
            d_3318_v5_ = _dafny.Set({_dafny.Map({p0: (d_3316_v3_).cf4}), _dafny.Map({(d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]: p0}), d_3317_v4_, default__.fm1(default__.fm2(p2, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uywfafmwc")), p0, d_3315_cf0_, globalState), d_3317_v4_})
            d_3319_v6_: _dafny.Map
            d_3319_v6_ = _dafny.Map({D0_DC0(-26): d_3318_v5_})
            d_3319_v6_ = (d_3319_v6_).set(default__.fm3(True, d_3315_cf0_, not(p2), globalState), d_3318_v5_)
            (globalState).f9 = p2
            d_3320_v7_: D0
            d_3320_v7_ = D0_DC0(p1)
            pat_let_tv75_ = p1
            d_3321_v8_: _dafny.Map
            def iife279_(_pat_let67_0):
                def iife280_(d_3322_dt__update__tmp_h0_):
                    def iife281_(_pat_let68_0):
                        def iife282_(d_3323_dt__update_hcf0_h0_):
                            return D0_DC0(d_3323_dt__update_hcf0_h0_)
                        return iife282_(_pat_let68_0)
                    return iife281_(pat_let_tv75_)
                return iife280_(_pat_let67_0)
            d_3321_v8_ = _dafny.Map({iife279_(d_3320_v7_): True})
            d_3321_v8_ = d_3321_v8_
        elif source40_.is_DC1:
            d_3324___mcc_h1_ = source40_.cf1
            d_3325___mcc_h2_ = source40_.cf2
            d_3326___mcc_h3_ = source40_.cf3
            d_3327_cf3_ = d_3326___mcc_h3_
            d_3328_cf2_ = d_3325___mcc_h2_
            d_3329_cf1_ = d_3324___mcc_h1_
            index557_ = default__.safeIndex(971, (d_3309_v0_).length(0))
            (d_3309_v0_)[index557_] = (((d_3313_v2_).cf5) == (p1)) and (p2)
            d_3330_v9_: str
            d_3330_v9_ = _dafny.CodePoint('w')
            (globalState).f5 = ((default__.fm4(d_3330_v9_, globalState)) + (p1)) + (len(p3))
            d_3331_v10_: _dafny.MultiSet
            d_3331_v10_ = _dafny.MultiSet([d_3328_cf2_])
            d_3332_v11_: _dafny.Map
            d_3332_v11_ = _dafny.Map({d_3327_cf3_: d_3331_v10_})
            d_3332_v11_ = (d_3332_v11_).set(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unan"))) + (p3)), d_3331_v10_)
            index558_ = default__.safeIndex(971, (d_3309_v0_).length(0))
            (d_3309_v0_)[index558_] = p0
        elif source40_.is_DC2:
            d_3333___mcc_h4_ = source40_.cf4
            d_3334_cf4_ = d_3333___mcc_h4_
            d_3335_v12_: _dafny.Seq
            d_3335_v12_ = _dafny.SeqWithoutIsStrInference([p3, default__.fm2(p0, globalState)])
            d_3336_v13_: _dafny.Map
            d_3336_v13_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_3337_i1_ in range(default__.abs(457))])})
            d_3338_v14_: str
            d_3338_v14_ = _dafny.CodePoint('w')
            d_3335_v12_ = _dafny.SeqWithoutIsStrInference([p3, ((d_3336_v13_)[d_3334_cf4_] if (d_3334_cf4_) in (d_3336_v13_) else (p3).set(default__.safeIndex(p1, len(p3)), d_3338_v14_)), (p3) + (p3)])
            d_3339_v15_: _dafny.Array
            nw538_ = _dafny.Array(int(0), 26)
            d_3339_v15_ = nw538_
            index559_ = default__.safeIndex(896, (d_3339_v15_).length(0))
            (d_3339_v15_)[index559_] = p1
            index560_ = default__.safeIndex(896, (d_3339_v15_).length(0))
            (d_3339_v15_)[index560_] = p1
            d_3340_v16_: _dafny.Array
            nw539_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_3340_v16_ = nw539_
            index561_ = default__.safeIndex(19, (d_3340_v16_).length(0))
            (d_3340_v16_)[index561_] = d_3309_v0_
            d_3341_v17_: _dafny.Map
            d_3341_v17_ = _dafny.Map({p1: (d_3339_v15_)[default__.safeIndex(896, (d_3339_v15_).length(0))]})
            index562_ = default__.safeIndex(19, (d_3340_v16_).length(0))
            rhs568_ = d_3309_v0_
            rhs569_ = p1
            rhs570_ = not ((((d_3341_v17_)[(d_3339_v15_)[default__.safeIndex(896, (d_3339_v15_).length(0))]] if ((d_3339_v15_)[default__.safeIndex(896, (d_3339_v15_).length(0))]) in (d_3341_v17_) else (d_3339_v15_)[default__.safeIndex(896, (d_3339_v15_).length(0))])) == (len(_dafny.Map({(d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]: d_3338_v14_})))) or (p0)
            lhs521_ = d_3340_v16_
            lhs522_ = default__.safeIndex(19, (d_3340_v16_).length(0))
            lhs523_ = globalState
            lhs524_ = globalState
            lhs521_[lhs522_] = rhs568_
            lhs523_.f0 = rhs569_
            lhs524_.f25 = rhs570_
            d_3342_v18_: _dafny.Set
            d_3342_v18_ = _dafny.Set({310, (p1) + ((d_3339_v15_)[default__.safeIndex(896, (d_3339_v15_).length(0))])})
            index563_ = default__.safeIndex(19, (d_3340_v16_).length(0))
            rhs571_ = len(d_3342_v18_)
            rhs572_ = d_3309_v0_
            lhs525_ = globalState
            lhs526_ = d_3340_v16_
            lhs527_ = default__.safeIndex(19, (d_3340_v16_).length(0))
            lhs525_.f24 = rhs571_
            lhs526_[lhs527_] = rhs572_
        elif source40_.is_DC3:
            d_3343___mcc_h5_ = source40_.cf5
            d_3344___mcc_h6_ = source40_.cf6
            d_3345_cf6_ = d_3344___mcc_h6_
            d_3346_cf5_ = d_3343___mcc_h5_
            d_3347_v19_: str
            d_3347_v19_ = _dafny.CodePoint('s')
            d_3348_v20_: _dafny.MultiSet
            d_3348_v20_ = _dafny.MultiSet([d_3345_cf6_])
            d_3349_v21_: _dafny.Seq
            d_3349_v21_ = _dafny.SeqWithoutIsStrInference([len((p3).set(default__.safeIndex((0) - (d_3346_cf5_), len(p3)), d_3347_v19_)), 258, (d_3348_v20_).cardinality, d_3346_cf5_, (0) - (p1)])
            d_3350_v22_: _dafny.Map
            d_3350_v22_ = _dafny.Map({len(_dafny.Map({p2: False})): p2})
            d_3351_v23_: _dafny.Seq
            d_3351_v23_ = _dafny.SeqWithoutIsStrInference([d_3350_v22_])
            d_3352_v24_: _dafny.MultiSet
            d_3352_v24_ = _dafny.MultiSet([d_3350_v22_])
            d_3353_v25_: _dafny.MultiSet
            d_3353_v25_ = _dafny.MultiSet([(d_3349_v21_)[default__.safeIndex(d_3345_cf6_, len(d_3349_v21_))], ((_dafny.MultiSet(d_3351_v23_)) | (d_3352_v24_)).cardinality, p1])
            d_3354_v26_: _dafny.Map
            d_3354_v26_ = _dafny.Map({d_3346_cf5_: p1})
            d_3348_v20_ = (_dafny.MultiSet([len(d_3354_v26_)])) | (d_3353_v25_)
            d_3355_v27_: _dafny.MultiSet
            d_3355_v27_ = _dafny.MultiSet([not(default__.fm5(p1, p0, p2, globalState))])
            index564_ = default__.safeIndex(971, (d_3309_v0_).length(0))
            rhs573_ = (d_3355_v27_).ispropersubset(d_3355_v27_)
            rhs574_ = d_3345_cf6_
            rhs575_ = p2
            lhs528_ = globalState
            lhs529_ = globalState
            lhs530_ = d_3309_v0_
            lhs531_ = default__.safeIndex(971, (d_3309_v0_).length(0))
            lhs528_.f25 = rhs573_
            lhs529_.f5 = rhs574_
            lhs530_[lhs531_] = rhs575_
            d_3356_v28_: _dafny.Map
            d_3356_v28_ = _dafny.Map({p2: p1})
            d_3357_v29_: C3
            nw540_ = C3()
            nw540_.ctor__(d_3345_cf6_, ((d_3356_v28_)[p2] if (p2) in (d_3356_v28_) else d_3346_cf5_), (p3) + (p3))
            d_3357_v29_ = nw540_
            d_3358_v30_: _dafny.Map
            d_3358_v30_ = _dafny.Map({d_3357_v29_: _dafny.Set({p2, (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]})})
            d_3359_v31_: D2
            d_3359_v31_ = D2_DC9(65, (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], len(d_3356_v28_), 275)
            d_3360_v32_: _dafny.Set
            d_3360_v32_ = _dafny.Set({len(d_3312_v1_), (d_3359_v31_).cf19})
            d_3361_v33_: _dafny.Map
            d_3361_v33_ = _dafny.Map({d_3346_cf5_: (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({553: d_3347_v19_})) for d_3362_i2_ in range(default__.abs(689))])).set(default__.safeIndex(((d_3356_v28_)[False] if (False) in (d_3356_v28_) else len(d_3358_v30_)), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({553: d_3347_v19_})) for d_3363_i2_ in range(default__.abs(689))]))), len(d_3360_v32_))})
            d_3364_v34_: _dafny.Map
            d_3364_v34_ = _dafny.Map({(d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]: p3})
            d_3365_v35_: _dafny.Array
            nw541_ = _dafny.Array(None, 7)
            nw541_[int(0)] = p3
            nw541_[int(1)] = p3
            nw541_[int(2)] = p3
            nw541_[int(3)] = p3
            nw541_[int(4)] = p3
            nw541_[int(5)] = (p3).set(default__.safeIndex(d_3346_cf5_, len(p3)), (p3)[default__.safeIndex(len(d_3361_v33_), len(p3))])
            nw541_[int(6)] = (((d_3364_v34_)[p0] if (p0) in (d_3364_v34_) else p3)) + (p3)
            d_3365_v35_ = nw541_
            index565_ = default__.safeIndex(219, (d_3365_v35_).length(0))
            (d_3365_v35_)[index565_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcv"))
            d_3366_v36_: D21
            d_3366_v36_ = D21_DC57(d_3347_v19_, (0) - (p1), (d_3357_v29_).f34, (d_3357_v29_).f34)
            index566_ = default__.safeIndex(219, (d_3365_v35_).length(0))
            index567_ = default__.safeIndex(971, (d_3309_v0_).length(0))
            rhs576_ = p3
            rhs577_ = ((0) - (p1)) == (d_3346_cf5_)
            rhs578_ = d_3345_cf6_
            rhs579_ = (0) - (default__.fm4((d_3366_v36_).cf108, globalState))
            lhs532_ = d_3365_v35_
            lhs533_ = default__.safeIndex(219, (d_3365_v35_).length(0))
            lhs534_ = d_3309_v0_
            lhs535_ = default__.safeIndex(971, (d_3309_v0_).length(0))
            lhs536_ = globalState
            lhs537_ = globalState
            lhs532_[lhs533_] = rhs576_
            lhs534_[lhs535_] = rhs577_
            lhs536_.f17 = rhs578_
            lhs537_.f5 = rhs579_
        elif True:
            d_3367___mcc_h7_ = source40_.cf7
            d_3368_cf7_ = d_3367___mcc_h7_
            d_3369_v37_: _dafny.MultiSet
            d_3369_v37_ = _dafny.MultiSet([p2, True])
            (globalState).f17 = default__.safeDivisionInt(p1, default__.safeDivisionInt(p1, (0) - (((d_3369_v37_)[p0] if (p0) in (d_3369_v37_) else -898))))
            d_3370_v38_: _dafny.Set
            d_3370_v38_ = _dafny.Set({p1})
            d_3371_v40_: _dafny.Seq
            d_3371_v40_ = _dafny.SeqWithoutIsStrInference([d_3370_v38_])
            d_3372_v42_: _dafny.Set
            d_3372_v42_ = _dafny.Set({_dafny.Set({p1})})
            d_3373_v43_: _dafny.MultiSet
            d_3373_v43_ = _dafny.MultiSet([_dafny.CodePoint('i')])
            d_3374_v44_: T1
            nw542_ = C5()
            nw542_.ctor__(p1, p1, (d_3373_v43_).cardinality, p3)
            d_3374_v44_ = nw542_
            d_3375_v45_: _dafny.Map
            d_3375_v45_ = _dafny.Map({p0: len(_dafny.SeqWithoutIsStrInference([p0]))})
            d_3376_v46_: _dafny.Map
            d_3376_v46_ = _dafny.Map({d_3374_v44_: d_3375_v45_})
            d_3377_v47_: str
            d_3377_v47_ = _dafny.CodePoint('f')
            d_3378_v48_: C6
            nw543_ = C6()
            nw543_.ctor__(_dafny.Map({((d_3376_v46_)[d_3374_v44_] if (d_3374_v44_) in (d_3376_v46_) else d_3375_v45_): d_3377_v47_}), len(_dafny.SeqWithoutIsStrInference([(d_3374_v44_).f26 for d_3379_i3_ in range(default__.abs(175))])), (d_3374_v44_).f27)
            d_3378_v48_ = nw543_
            d_3380_v49_: _dafny.Map
            d_3380_v49_ = _dafny.Map({p1: d_3378_v48_})
            def iife283_():
                coll145_ = _dafny.Set()
                compr_145_: int
                for compr_145_ in _dafny.IntegerRange(811, 384):
                    d_3381_v39_: int = compr_145_
                    if ((811) <= (d_3381_v39_)) and ((d_3381_v39_) < (384)):
                        coll145_ = coll145_.union(_dafny.Set([(d_3381_v39_) - (198)]))
                return _dafny.Set(coll145_)
            def iife284_():
                coll146_ = _dafny.Set()
                compr_146_: _dafny.Set
                for compr_146_ in (d_3371_v40_).Elements:
                    d_3382_v41_: _dafny.Set = compr_146_
                    if (d_3382_v41_) in (d_3371_v40_):
                        coll146_ = coll146_.union(_dafny.Set([d_3382_v41_]))
                return _dafny.Set(coll146_)
            if default__.fm5(len((d_3370_v38_) - (iife283_()
            )), (d_3372_v42_).ispropersubset(iife284_()
            ), (default__.fm5(len((d_3380_v49_).set(-360, d_3378_v48_)), p2, p2, globalState) if (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))] else False), globalState):
                d_3383_v50_: _dafny.MultiSet
                d_3383_v50_ = _dafny.MultiSet([p1, p1, 511])
                d_3384_v51_: _dafny.Seq
                d_3384_v51_ = _dafny.SeqWithoutIsStrInference([(d_3374_v44_).f26])
                d_3385_v52_: _dafny.Seq
                d_3385_v52_ = _dafny.SeqWithoutIsStrInference([d_3383_v50_, _dafny.MultiSet(d_3384_v51_)])
                d_3386_v53_: _dafny.Map
                d_3386_v53_ = _dafny.Map({d_3385_v52_: p0})
                d_3387_v54_: _dafny.Set
                d_3387_v54_ = _dafny.Set({p0})
                d_3386_v53_ = (d_3386_v53_).set(d_3385_v52_, (d_3387_v54_).ispropersubset(d_3387_v54_))
                d_3388_v55_: _dafny.Array
                nw544_ = _dafny.Array(None, 13)
                nw544_[int(0)] = len((d_3374_v44_).f27)
                nw544_[int(1)] = (d_3374_v44_).f26
                nw544_[int(2)] = len((d_3374_v44_).f27)
                nw544_[int(3)] = (0) - ((d_3374_v44_).f26)
                nw544_[int(4)] = (d_3374_v44_).f26
                nw544_[int(5)] = (d_3374_v44_).f26
                nw544_[int(6)] = (d_3378_v48_).fm9((d_3374_v44_).f26, (d_3374_v44_).f26, p0, globalState)
                nw544_[int(7)] = p1
                nw544_[int(8)] = p1
                nw544_[int(9)] = p1
                nw544_[int(10)] = p1
                nw544_[int(11)] = p1
                nw544_[int(12)] = (d_3374_v44_).f26
                d_3388_v55_ = nw544_
                d_3389_v56_: D7
                d_3389_v56_ = D7_DC18((d_3369_v37_).cardinality, d_3388_v55_)
                (globalState).f8 = ((d_3389_v56_).cf34) - (default__.safeModuloInt(p1, p1))
                d_3390_v57_: int
                d_3391_v58_: _dafny.MultiSet
                d_3392_v59_: _dafny.Array
                out63_: int
                out64_: _dafny.MultiSet
                out65_: _dafny.Array
                out63_, out64_, out65_ = (d_3378_v48_).m2(((d_3374_v44_).f26) != (p1), D0_DC2(p0), globalState)
                d_3390_v57_ = out63_
                d_3391_v58_ = out64_
                d_3392_v59_ = out65_
                (globalState).f23 = default__.fm36(p2, (d_3312_v1_)[default__.safeIndex((d_3374_v44_).f26, len(d_3312_v1_))], (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], globalState)
                index568_ = default__.safeIndex(996, (d_3388_v55_).length(0))
                (d_3388_v55_)[index568_] = ((0) - ((d_3374_v44_).f26) if (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))] else (d_3374_v44_).f26)
                index569_ = default__.safeIndex(996, (d_3388_v55_).length(0))
                (d_3388_v55_)[index569_] = 290
            elif True:
                (globalState).f25 = ((d_3374_v44_).f26) > ((0) - ((d_3378_v48_).fm7(len(p3), p0, globalState)))
                index570_ = default__.safeIndex(971, (d_3309_v0_).length(0))
                (d_3309_v0_)[index570_] = p0
                d_3393_v60_: _dafny.Map
                d_3393_v60_ = _dafny.Map({p1: d_3312_v1_})
                d_3394_v61_: D24
                d_3394_v61_ = D24_DC73(False, p2, d_3393_v60_)
                d_3394_v61_ = (d_3394_v61_ if not(p2) else D24_DC73((d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], d_3393_v60_))
                d_3395_v62_: _dafny.Map
                d_3395_v62_ = _dafny.Map({(d_3374_v44_).f26: p1})
                d_3396_v63_: _dafny.Seq
                d_3396_v63_ = _dafny.SeqWithoutIsStrInference([d_3312_v1_, d_3312_v1_])
                d_3395_v62_ = (d_3395_v62_).set(p1, len((d_3312_v1_) + ((d_3396_v63_)[default__.safeIndex(p1, len(d_3396_v63_))])))
                (globalState).f9 = p0
            if p2:
                d_3397_v64_: _dafny.Array
                nw545_ = _dafny.Array(int(0), 11)
                d_3397_v64_ = nw545_
                index571_ = default__.safeIndex(691, (d_3397_v64_).length(0))
                (d_3397_v64_)[index571_] = ((d_3374_v44_).f26) * (len(_dafny.Set({True, p2})))
                d_3398_v65_: _dafny.Array
                nw546_ = _dafny.Array(_dafny.CodePoint('D'), 8)
                d_3398_v65_ = nw546_
                index572_ = default__.safeIndex(695, (d_3398_v65_).length(0))
                (d_3398_v65_)[index572_] = d_3377_v47_
                d_3399_v66_: _dafny.MultiSet
                d_3399_v66_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_3377_v47_ for d_3400_i4_ in range(default__.abs(218))])])
                index573_ = default__.safeIndex(691, (d_3397_v64_).length(0))
                index574_ = default__.safeIndex(695, (d_3398_v65_).length(0))
                rhs580_ = p1
                rhs581_ = default__.safeModuloInt(((d_3399_v66_).cardinality) * ((d_3374_v44_).f26), (d_3374_v44_).f26)
                rhs582_ = d_3377_v47_
                lhs538_ = globalState
                lhs539_ = d_3397_v64_
                lhs540_ = default__.safeIndex(691, (d_3397_v64_).length(0))
                lhs541_ = d_3398_v65_
                lhs542_ = default__.safeIndex(695, (d_3398_v65_).length(0))
                lhs538_.f24 = rhs580_
                lhs539_[lhs540_] = rhs581_
                lhs541_[lhs542_] = rhs582_
                d_3401_v67_: _dafny.Seq
                d_3401_v67_ = _dafny.SeqWithoutIsStrInference([p1])
                d_3402_v68_: _dafny.Map
                d_3402_v68_ = _dafny.Map({(d_3312_v1_)[default__.safeIndex((d_3374_v44_).f26, len(d_3312_v1_))]: p0})
                d_3403_v69_: _dafny.Map
                d_3403_v69_ = _dafny.Map({len(d_3402_v68_): p0})
                d_3404_v70_: D25
                d_3404_v70_ = D25_DC75(d_3374_v44_)
                d_3405_v72_: _dafny.Set
                d_3405_v72_ = _dafny.Set({(d_3312_v1_).set(default__.safeIndex(len(d_3402_v68_), len(d_3312_v1_)), (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))])})
                index575_ = default__.safeIndex(691, (d_3397_v64_).length(0))
                def iife285_():
                    coll147_ = _dafny.Map()
                    compr_147_: _dafny.Seq
                    for compr_147_ in (d_3405_v72_).Elements:
                        d_3406_v71_: _dafny.Seq = compr_147_
                        if (d_3406_v71_) in (d_3405_v72_):
                            coll147_[d_3406_v71_] = (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]
                    return _dafny.Map(coll147_)
                rhs583_ = not ((p2 if p2 else p0)) or (p0)
                rhs584_ = (d_3401_v67_)[default__.safeIndex((d_3378_v48_).fm6(d_3403_v69_, (d_3374_v44_).f26, _dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([d_3374_v44_, d_3374_v44_, (d_3404_v70_).cf145])).set(default__.safeIndex((d_3397_v64_)[default__.safeIndex(691, (d_3397_v64_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_3374_v44_, d_3374_v44_, (d_3404_v70_).cf145]))), d_3374_v44_)), (d_3369_v37_).cardinality]), iife285_()
                , globalState), len(d_3401_v67_))]
                rhs585_ = d_3377_v47_
                lhs543_ = globalState
                lhs544_ = d_3397_v64_
                lhs545_ = default__.safeIndex(691, (d_3397_v64_).length(0))
                lhs543_.f25 = rhs583_
                lhs544_[lhs545_] = rhs584_
                d_3377_v47_ = rhs585_
                d_3407_v73_: _dafny.Array
                nw547_ = _dafny.Array(_dafny.Set({}), 3)
                d_3407_v73_ = nw547_
                index576_ = default__.safeIndex(691, (d_3397_v64_).length(0))
                rhs586_ = d_3407_v73_
                rhs587_ = d_3377_v47_
                rhs588_ = d_3309_v0_
                rhs589_ = (d_3397_v64_)[default__.safeIndex(691, (d_3397_v64_).length(0))]
                lhs546_ = globalState
                lhs547_ = d_3397_v64_
                lhs548_ = default__.safeIndex(691, (d_3397_v64_).length(0))
                d_3407_v73_ = rhs586_
                d_3377_v47_ = rhs587_
                lhs546_.f1 = rhs588_
                lhs547_[lhs548_] = rhs589_
                d_3408_v74_: D11
                d_3408_v74_ = D11_DC34(d_3309_v0_, p2)
                d_3409_v75_: _dafny.Array
                nw548_ = _dafny.Array(None, 21)
                nw548_[int(0)] = d_3309_v0_
                nw548_[int(1)] = (d_3408_v74_).cf66
                nw548_[int(2)] = d_3309_v0_
                nw548_[int(3)] = d_3309_v0_
                nw548_[int(4)] = d_3309_v0_
                nw548_[int(5)] = d_3309_v0_
                nw548_[int(6)] = d_3309_v0_
                nw548_[int(7)] = d_3309_v0_
                nw548_[int(8)] = d_3309_v0_
                nw548_[int(9)] = d_3309_v0_
                nw548_[int(10)] = d_3309_v0_
                nw548_[int(11)] = d_3309_v0_
                nw548_[int(12)] = d_3309_v0_
                nw548_[int(13)] = d_3309_v0_
                nw548_[int(14)] = d_3309_v0_
                nw548_[int(15)] = d_3309_v0_
                nw548_[int(16)] = d_3309_v0_
                nw548_[int(17)] = d_3309_v0_
                nw548_[int(18)] = d_3309_v0_
                nw548_[int(19)] = d_3309_v0_
                nw548_[int(20)] = d_3309_v0_
                d_3409_v75_ = nw548_
                d_3410_v76_: _dafny.Map
                d_3410_v76_ = _dafny.Map({(d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]: d_3409_v75_})
                d_3410_v76_ = (d_3410_v76_).set((d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], d_3409_v75_)
                (globalState).f25 = (d_3312_v1_)[default__.safeIndex(len(d_3370_v38_), len(d_3312_v1_))]
            elif True:
                d_3411_v77_: _dafny.Map
                d_3411_v77_ = _dafny.Map({(d_3374_v44_).f27: (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]})
                d_3412_v78_: _dafny.Map
                d_3412_v78_ = d_3411_v77_
                d_3412_v78_ = d_3412_v78_
                index577_ = default__.safeIndex(971, (d_3309_v0_).length(0))
                (d_3309_v0_)[index577_] = ((d_3374_v44_).f26) < (len(d_3370_v38_))
                d_3413_v79_: C5
                nw549_ = C5()
                nw549_.ctor__(((d_3374_v44_).f26) - ((d_3374_v44_).f26), (0) - ((0) - (p1)), p1, (d_3374_v44_).f27)
                d_3413_v79_ = nw549_
                d_3414_v80_: _dafny.Set
                d_3414_v80_ = _dafny.Set({p0, not(p2), (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], p0})
                d_3414_v80_ = d_3414_v80_
                d_3415_v81_: _dafny.Map
                d_3415_v81_ = _dafny.Map({693: (d_3374_v44_).f26})
                d_3416_v82_: _dafny.Map
                d_3416_v82_ = _dafny.Map({d_3377_v47_: d_3415_v81_})
                d_3417_v83_: _dafny.MultiSet
                d_3417_v83_ = _dafny.MultiSet([(d_3374_v44_).f26])
                d_3418_v84_: _dafny.Map
                d_3418_v84_ = _dafny.Map({d_3377_v47_: (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]})
                d_3419_v86_: _dafny.Map
                def iife286_():
                    coll148_ = _dafny.Set()
                    compr_148_: int
                    for compr_148_ in _dafny.IntegerRange(-597, 250):
                        d_3420_v85_: int = compr_148_
                        if ((-597) <= (d_3420_v85_)) and ((d_3420_v85_) < (250)):
                            coll148_ = coll148_.union(_dafny.Set([default__.safeModuloInt(d_3420_v85_, d_3413_v79_.f29)]))
                    return _dafny.Set(coll148_)
                d_3419_v86_ = _dafny.Map({len(iife286_()
                ): (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]})
                d_3421_v87_: _dafny.Seq
                d_3421_v87_ = _dafny.SeqWithoutIsStrInference([349])
                d_3422_v88_: _dafny.Map
                d_3422_v88_ = _dafny.Map({d_3312_v1_: p2})
                d_3423_v89_: T0
                nw550_ = C1()
                nw550_.ctor__()
                d_3423_v89_ = nw550_
                d_3424_v90_: _dafny.Seq
                d_3424_v90_ = _dafny.SeqWithoutIsStrInference([d_3423_v89_])
                d_3425_v91_: _dafny.Map
                d_3425_v91_ = _dafny.Map({p2: d_3424_v90_})
                d_3426_v92_: _dafny.Array
                nw551_ = _dafny.Array(None, 29)
                nw551_[int(0)] = (0) - (d_3413_v79_.f29)
                nw551_[int(1)] = default__.fm4(_dafny.CodePoint('x'), globalState)
                nw551_[int(2)] = len(d_3416_v82_)
                nw551_[int(3)] = (0) - (p1)
                nw551_[int(4)] = d_3413_v79_.f29
                nw551_[int(5)] = ((d_3373_v43_)[d_3377_v47_] if (d_3377_v47_) in (d_3373_v43_) else d_3413_v79_.f30)
                nw551_[int(6)] = (d_3374_v44_).f26
                nw551_[int(7)] = (d_3413_v79_.f29) + ((d_3374_v44_).f26)
                nw551_[int(8)] = default__.safeModuloInt(d_3413_v79_.f30, (d_3417_v83_).cardinality)
                nw551_[int(9)] = (d_3374_v44_).f26
                nw551_[int(10)] = len(d_3418_v84_)
                nw551_[int(11)] = p1
                nw551_[int(12)] = (d_3378_v48_).fm6(d_3419_v86_, len(_dafny.SeqWithoutIsStrInference([(d_3374_v44_).f26])), d_3421_v87_, ((D26_DC78(d_3422_v88_)).cf149).set(d_3312_v1_, p0), globalState)
                nw551_[int(13)] = default__.safeDivisionInt((d_3374_v44_).f26, d_3413_v79_.f29)
                nw551_[int(14)] = d_3413_v79_.f30
                nw551_[int(15)] = (d_3374_v44_).f26
                nw551_[int(16)] = d_3413_v79_.f30
                nw551_[int(17)] = (0) - (((d_3374_v44_).f26) * (p1))
                nw551_[int(18)] = len(_dafny.Map({d_3309_v0_: p0}))
                nw551_[int(19)] = d_3413_v79_.f30
                nw551_[int(20)] = (d_3374_v44_).f26
                nw551_[int(21)] = len(d_3370_v38_)
                nw551_[int(22)] = (d_3374_v44_).f26
                nw551_[int(23)] = d_3413_v79_.f30
                nw551_[int(24)] = 179
                nw551_[int(25)] = (d_3413_v79_.f29) + ((d_3374_v44_).f26)
                nw551_[int(26)] = (len(d_3370_v38_) if p2 else p1)
                nw551_[int(27)] = d_3413_v79_.f30
                nw551_[int(28)] = len(((d_3425_v91_).set(p2, d_3424_v90_)) | (d_3425_v91_))
                d_3426_v92_ = nw551_
                d_3426_v92_ = d_3426_v92_
            d_3377_v47_ = ((d_3374_v44_).f27)[default__.safeIndex((d_3374_v44_).f26, len((d_3374_v44_).f27))]
        if p0:
            (globalState).f5 = p1
            d_3309_v0_ = d_3309_v0_
            d_3427_v93_: str
            d_3427_v93_ = _dafny.CodePoint('m')
            (globalState).f8 = default__.fm4(d_3427_v93_, globalState)
            d_3428_v94_: _dafny.Array
            nw552_ = _dafny.Array(int(0), 6)
            d_3428_v94_ = nw552_
            index578_ = default__.safeIndex(704, (d_3428_v94_).length(0))
            (d_3428_v94_)[index578_] = p1
            d_3429_v95_: _dafny.MultiSet
            d_3429_v95_ = _dafny.MultiSet([932])
            d_3430_v96_: C0
            nw553_ = C0()
            nw553_.ctor__()
            d_3430_v96_ = nw553_
            d_3431_v97_: D24
            d_3431_v97_ = D24_DC70(d_3430_v96_)
            d_3432_v98_: _dafny.Seq
            d_3432_v98_ = _dafny.SeqWithoutIsStrInference([d_3431_v97_])
            d_3433_v99_: _dafny.MultiSet
            d_3433_v99_ = _dafny.MultiSet([d_3432_v98_])
            d_3434_v100_: _dafny.Map
            d_3434_v100_ = _dafny.Map({p1: 91})
            index579_ = default__.safeIndex(704, (d_3428_v94_).length(0))
            rhs590_ = (d_3429_v95_).isdisjoint((_dafny.MultiSet([p1])) | (d_3429_v95_))
            rhs591_ = ((d_3433_v99_)[d_3432_v98_] if (d_3432_v98_) in (d_3433_v99_) else (0) - ((len(d_3434_v100_)) + (p1)))
            rhs592_ = default__.safeModuloInt(p1, p1)
            lhs549_ = globalState
            lhs550_ = d_3428_v94_
            lhs551_ = default__.safeIndex(704, (d_3428_v94_).length(0))
            lhs552_ = globalState
            lhs549_.f25 = rhs590_
            lhs550_[lhs551_] = rhs591_
            lhs552_.f0 = rhs592_
            if (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]:
                d_3435_v101_: _dafny.Map
                d_3435_v101_ = _dafny.Map({D6_DC16(p1, False): (d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))]})
                d_3436_v102_: D6
                d_3436_v102_ = D6_DC16(p1, p0)
                d_3435_v101_ = (d_3435_v101_).set(d_3436_v102_, (d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))])
                (globalState).f25 = p2
                nw554_ = _dafny.Array(False, 24)
                d_3309_v0_ = nw554_
                d_3437_v103_: C4
                nw555_ = C4()
                nw555_.ctor__(not(((_dafny.MultiSet([len(p3), p1, (d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))]])).cardinality) > ((d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))])), (d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))], p3, _dafny.Map({(d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))]: len(p3)}))
                d_3437_v103_ = nw555_
                d_3438_v104_: _dafny.Map
                d_3438_v104_ = _dafny.Map({(d_3437_v103_).fm8(d_3429_v95_, (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], len(d_3312_v1_), (d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))], globalState): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxi"))})
                d_3439_v105_: C6
                nw556_ = C6()
                nw556_.ctor__(_dafny.Map({_dafny.Map({(d_3437_v103_).f32: (d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))]}): d_3427_v93_}), p1, (p3) + (((d_3438_v104_)[p0] if (p0) in (d_3438_v104_) else p3)))
                d_3439_v105_ = nw556_
            elif True:
                index580_ = default__.safeIndex(971, (d_3309_v0_).length(0))
                (d_3309_v0_)[index580_] = p2
                (globalState).f5 = -69
                d_3440_v106_: _dafny.Array
                nw557_ = _dafny.Array(None, 10)
                nw557_[int(0)] = d_3427_v93_
                nw557_[int(1)] = d_3427_v93_
                nw557_[int(2)] = d_3427_v93_
                nw557_[int(3)] = d_3427_v93_
                nw557_[int(4)] = d_3427_v93_
                nw557_[int(5)] = d_3427_v93_
                nw557_[int(6)] = d_3427_v93_
                nw557_[int(7)] = _dafny.CodePoint('s')
                nw557_[int(8)] = d_3427_v93_
                nw557_[int(9)] = d_3427_v93_
                d_3440_v106_ = nw557_
                d_3441_v107_: D20
                d_3441_v107_ = D20_DC54(d_3440_v106_)
                d_3442_v108_: _dafny.Array
                nw558_ = _dafny.Array(None, 15)
                nw558_[int(0)] = d_3441_v107_
                nw558_[int(1)] = d_3441_v107_
                nw558_[int(2)] = d_3441_v107_
                nw558_[int(3)] = D20_DC54(d_3440_v106_)
                nw558_[int(4)] = d_3441_v107_
                nw558_[int(5)] = d_3441_v107_
                nw558_[int(6)] = d_3441_v107_
                nw558_[int(7)] = d_3441_v107_
                nw558_[int(8)] = d_3441_v107_
                nw558_[int(9)] = d_3441_v107_
                nw558_[int(10)] = d_3441_v107_
                nw558_[int(11)] = d_3441_v107_
                nw558_[int(12)] = d_3441_v107_
                nw558_[int(13)] = d_3441_v107_
                nw558_[int(14)] = d_3441_v107_
                d_3442_v108_ = nw558_
                index581_ = default__.safeIndex(828, (d_3442_v108_).length(0))
                (d_3442_v108_)[index581_] = d_3441_v107_
                index582_ = default__.safeIndex(828, (d_3442_v108_).length(0))
                (d_3442_v108_)[index582_] = d_3441_v107_
                d_3443_v109_: _dafny.MultiSet
                d_3443_v109_ = _dafny.MultiSet([(p1) < (p1), ((d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))]) < (p1), default__.fm5(p1, (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], p2, globalState)])
                d_3444_v110_: _dafny.Set
                d_3444_v110_ = _dafny.Set({p2, p0, p2, p0})
                d_3445_v111_: _dafny.Map
                d_3445_v111_ = _dafny.Map({default__.fm0(p2, p3, d_3444_v110_, default__.fm4(d_3427_v93_, globalState), globalState): (d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))]})
                d_3443_v109_ = (d_3443_v109_) - ((default__.fm37(p0, default__.fm24(p0, p2, (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], len(d_3445_v111_), globalState), (d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))], (d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))], globalState)) | (_dafny.MultiSet([default__.fm5((d_3428_v94_)[default__.safeIndex(704, (d_3428_v94_).length(0))], (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], not(p2), globalState), False])))
                d_3444_v110_ = d_3444_v110_
        elif True:
            d_3446_v112_: _dafny.Array
            nw559_ = _dafny.Array(int(0), 9)
            d_3446_v112_ = nw559_
            d_3447_v113_: _dafny.MultiSet
            d_3447_v113_ = _dafny.MultiSet([p1])
            index583_ = default__.safeIndex(981, (d_3446_v112_).length(0))
            (d_3446_v112_)[index583_] = (d_3447_v113_).cardinality
            index584_ = default__.safeIndex(981, (d_3446_v112_).length(0))
            (d_3446_v112_)[index584_] = (0) - (p1)
            (globalState).f23 = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_3448_i5_ in range(default__.abs(643))])
            d_3449_v114_: str
            d_3449_v114_ = _dafny.CodePoint('u')
            d_3450_v115_: _dafny.Map
            d_3450_v115_ = _dafny.Map({p2: p1})
            nw560_ = _dafny.Array(None, 12)
            nw560_[int(0)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efiixvi")))
            nw560_[int(1)] = default__.safeModuloInt(default__.fm4(d_3449_v114_, globalState), p1)
            nw560_[int(2)] = (d_3446_v112_)[default__.safeIndex(981, (d_3446_v112_).length(0))]
            nw560_[int(3)] = (d_3446_v112_)[default__.safeIndex(981, (d_3446_v112_).length(0))]
            nw560_[int(4)] = (d_3446_v112_)[default__.safeIndex(981, (d_3446_v112_).length(0))]
            nw560_[int(5)] = p1
            nw560_[int(6)] = (d_3446_v112_)[default__.safeIndex(981, (d_3446_v112_).length(0))]
            nw560_[int(7)] = (len((p3).set(default__.safeIndex((d_3446_v112_)[default__.safeIndex(981, (d_3446_v112_).length(0))], len(p3)), d_3449_v114_))) - ((d_3446_v112_)[default__.safeIndex(981, (d_3446_v112_).length(0))])
            nw560_[int(8)] = -432
            nw560_[int(9)] = default__.safeDivisionInt(((d_3450_v115_)[p0] if (p0) in (d_3450_v115_) else p1), (d_3446_v112_)[default__.safeIndex(981, (d_3446_v112_).length(0))])
            nw560_[int(10)] = p1
            nw560_[int(11)] = p1
            d_3446_v112_ = nw560_
            (globalState).f0 = (d_3446_v112_)[default__.safeIndex(981, (d_3446_v112_).length(0))]
            d_3451_v116_: _dafny.Seq
            d_3451_v116_ = _dafny.SeqWithoutIsStrInference([(D2_DC7(d_3312_v1_)).cf12, _dafny.SeqWithoutIsStrInference([p2, p0]), d_3312_v1_, _dafny.SeqWithoutIsStrInference([p0, False, (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))], p2]), _dafny.SeqWithoutIsStrInference([p0])])
            d_3312_v1_ = (d_3451_v116_)[default__.safeIndex(-318, len(d_3451_v116_))]
        if (len(default__.fm47(globalState))) < (p1):
            d_3452_v117_: _dafny.Map
            d_3452_v117_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_3453_i6_ in range(default__.abs(635))]): (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))]})
            d_3454_v118_: _dafny.Map
            d_3454_v118_ = d_3452_v117_
            rhs593_ = d_3454_v118_
            rhs594_ = (p0) or (not (p2) or (p0))
            lhs553_ = globalState
            d_3454_v118_ = rhs593_
            lhs553_.f25 = rhs594_
            d_3455_v119_: str
            d_3455_v119_ = _dafny.CodePoint('l')
            d_3455_v119_ = (d_3455_v119_ if (d_3309_v0_)[default__.safeIndex(971, (d_3309_v0_).length(0))] else d_3455_v119_)
            d_3456_v120_: _dafny.Map
            d_3456_v120_ = _dafny.Map({len(d_3312_v1_): d_3455_v119_})
            d_3457_v121_: C5
            nw561_ = C5()
            nw561_.ctor__(p1, (0) - (p1), len(d_3456_v120_), (p3 if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gns"))))
            d_3457_v121_ = nw561_
            d_3458_v122_: _dafny.Array
            def lambda156_(d_3459_p1_):
                def lambda157_(d_3460_i7_):
                    return default__.safeModuloInt(d_3460_i7_, d_3459_p1_)

                return lambda157_

            init92_ = lambda156_(p1)
            nw562_ = _dafny.Array(None, 27)
            for i0_92_ in range(nw562_.length(0)):
                nw562_[i0_92_] = init92_(i0_92_)
            d_3458_v122_ = nw562_
            index585_ = default__.safeIndex(0, (d_3458_v122_).length(0))
            (d_3458_v122_)[index585_] = (_dafny.MultiSet([d_3457_v121_.f29, d_3457_v121_.f30, d_3457_v121_.f30])).cardinality
            d_3461_v123_: _dafny.Seq
            d_3461_v123_ = _dafny.SeqWithoutIsStrInference([p3, p3])
            d_3462_v124_: _dafny.Map
            d_3462_v124_ = _dafny.Map({len((d_3312_v1_).set(default__.safeIndex(d_3457_v121_.f29, len(d_3312_v1_)), not(True))): d_3457_v121_.f30})
            d_3463_v125_: _dafny.MultiSet
            d_3463_v125_ = _dafny.MultiSet([d_3457_v121_.f30, d_3457_v121_.f30, p1])
            d_3464_v126_: _dafny.Seq
            d_3464_v126_ = _dafny.SeqWithoutIsStrInference([d_3457_v121_.f29, (0) - ((0) - (d_3457_v121_.f29)), -432])
            d_3465_v127_: D21
            d_3465_v127_ = D21_DC57(d_3455_v119_, 1, d_3457_v121_.f29, 873)
            d_3466_v128_: _dafny.Map
            d_3466_v128_ = _dafny.Map({(d_3465_v127_).cf111: d_3463_v125_})
            index586_ = default__.safeIndex(0, (d_3458_v122_).length(0))
            index587_ = default__.safeIndex(971, (d_3309_v0_).length(0))
            rhs595_ = ((0) - (-434)) + (p1)
            rhs596_ = ((d_3457_v121_.f30 if False else len((d_3461_v123_)[default__.safeIndex(d_3457_v121_.f29, len(d_3461_v123_))]))) - ((d_3457_v121_.f30) + (len(d_3462_v124_)))
            rhs597_ = ((d_3463_v125_).set(len(p3), default__.abs((d_3464_v126_)[default__.safeIndex(d_3457_v121_.f30, len(d_3464_v126_))]))).isdisjoint(((d_3466_v128_)[len(_dafny.SeqWithoutIsStrInference([d_3457_v121_.f30, len(p3)]))] if (len(_dafny.SeqWithoutIsStrInference([d_3457_v121_.f30, len(p3)]))) in (d_3466_v128_) else d_3463_v125_))
            rhs598_ = d_3312_v1_
            lhs554_ = d_3458_v122_
            lhs555_ = default__.safeIndex(0, (d_3458_v122_).length(0))
            lhs556_ = globalState
            lhs557_ = d_3309_v0_
            lhs558_ = default__.safeIndex(971, (d_3309_v0_).length(0))
            lhs554_[lhs555_] = rhs595_
            lhs556_.f0 = rhs596_
            lhs557_[lhs558_] = rhs597_
            d_3312_v1_ = rhs598_
            (globalState).f25 = not(((p3) + (p3)) < ((p3) + (p3)))
        elif True:
            (globalState).f5 = p1
            d_3467_v129_: _dafny.Map
            d_3467_v129_ = _dafny.Map({p1: 915})
            (globalState).f9 = (d_3312_v1_)[default__.safeIndex(len(d_3467_v129_), len(d_3312_v1_))]
            d_3468_v130_: str
            d_3468_v130_ = _dafny.CodePoint('s')
            (globalState).f9 = (p1) != ((default__.fm4(d_3468_v130_, globalState)) * (821))
            (globalState).f8 = (default__.safeModuloInt(len(p3), -650)) * (p1)
            index588_ = default__.safeIndex(971, (d_3309_v0_).length(0))
            rhs599_ = _dafny.SeqWithoutIsStrInference([d_3468_v130_ for d_3469_i8_ in range(default__.abs(659))])
            rhs600_ = d_3468_v130_
            rhs601_ = (p1) > ((p1) - (p1))
            rhs602_ = p1
            lhs559_ = globalState
            lhs560_ = d_3309_v0_
            lhs561_ = default__.safeIndex(971, (d_3309_v0_).length(0))
            lhs562_ = globalState
            lhs559_.f13 = rhs599_
            d_3468_v130_ = rhs600_
            lhs560_[lhs561_] = rhs601_
            lhs562_.f8 = rhs602_
        hi12_ = p1
        for d_3470_i9_ in range(p1, hi12_):
            d_3471_v131_: C0
            nw563_ = C0()
            nw563_.ctor__()
            d_3471_v131_ = nw563_
            d_3472_v132_: _dafny.Array
            nw564_ = _dafny.Array(int(0), 9)
            d_3472_v132_ = nw564_
            d_3473_v133_: str
            d_3473_v133_ = _dafny.CodePoint('a')
            index589_ = default__.safeIndex(951, (d_3472_v132_).length(0))
            (d_3472_v132_)[index589_] = default__.safeDivisionInt(default__.fm4(d_3473_v133_, globalState), p1)
            index590_ = default__.safeIndex(951, (d_3472_v132_).length(0))
            (d_3472_v132_)[index590_] = p1
            d_3474_v134_: _dafny.MultiSet
            d_3474_v134_ = _dafny.MultiSet([d_3472_v132_])
            (globalState).f5 = (0) - (((d_3474_v134_)[d_3472_v132_] if (d_3472_v132_) in (d_3474_v134_) else (d_3472_v132_)[default__.safeIndex(951, (d_3472_v132_).length(0))]))
            d_3475_v135_: _dafny.Map
            d_3475_v135_ = _dafny.Map({default__.safeDivisionInt(p1, (0) - ((d_3472_v132_)[default__.safeIndex(951, (d_3472_v132_).length(0))])): p1})
            d_3475_v135_ = (d_3475_v135_).set((-976) + ((d_3472_v132_)[default__.safeIndex(951, (d_3472_v132_).length(0))]), default__.fm4(d_3473_v133_, globalState))
        (globalState).f24 = ((0) - (len(d_3312_v1_))) * (p1)

