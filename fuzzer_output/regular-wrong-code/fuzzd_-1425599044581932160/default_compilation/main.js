// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm6(p0, globalState) {
      if (false) {
        return _module.D0.create_DC0(true);
      } else {
        return _module.D0.create_DC0(true);
      }
    };
    static fm7(p0, p1, p2, globalState) {
      if (false) {
        return _module.D3.create_DC8(_dafny.Seq.of(_module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(615)), function (_0_i0) {
  return new _dafny.CodePoint('o'.codePointAt(0));
}))));
      } else {
        return _module.D3.create_DC8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(635)), function (_1_i1) {
  return _module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-926)), function (_2_i2) {
  return new _dafny.CodePoint('k'.codePointAt(0));
}));
}));
      }
    };
    static fm9(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qeexk"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(105)), function (_3_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tcrjsmo"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(995)), function (_4_i1) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      })));
    };
    static fm10(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("acxmpij"), _dafny.Seq.UnicodeFromString("xt"))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true, !(false), true, false)).length)), new BigNumber((_dafny.Seq.of(new BigNumber(500))).length)))).Merge(((false) ? (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(410))).cardinality())))) : (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(new BigNumber(115))))));
    };
    static fm11(p0, globalState) {
      let _source0 = _module.D1.create_DC3(_dafny.Seq.UnicodeFromString("kbtb"));
      if (_source0.is_DC4) {
        let _5___mcc_h0 = (_source0).cf5;
        let _6___mcc_h1 = (_source0).cf6;
        let _7_cf6 = _6___mcc_h1;
        let _8_cf5 = _5___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_8_cf5), _dafny.Seq.of(_8_cf5, _7_cf6, _8_cf5, _8_cf5));
      } else if (_source0.is_DC5) {
        let _9___mcc_h2 = (_source0).cf7;
        let _10___mcc_h3 = (_source0).cf8;
        let _11___mcc_h4 = (_source0).cf9;
        let _12___mcc_h5 = (_source0).cf10;
        let _13_cf10 = _12___mcc_h5;
        let _14_cf9 = _11___mcc_h4;
        let _15_cf8 = _10___mcc_h3;
        let _16_cf7 = _9___mcc_h2;
        if (!(_14_cf9)) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(145)), ((_17_cf7) => function (_18_i0) {
            return (_dafny.ZERO).minus(_17_cf7);
          })(_16_cf7));
        } else {
          return _dafny.Seq.of(_15_cf8);
        }
      } else {
        let _19___mcc_h6 = (_source0).cf4;
        let _20_cf4 = _19___mcc_h6;
        return _dafny.Seq.of(new BigNumber(767));
      }
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return (((false) ? (_dafny.Set.fromElements(new BigNumber(492))) : (_dafny.Set.fromElements(new BigNumber(135), new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(688), new BigNumber(805))).length)), new BigNumber(52))).Elements) {
          let _21_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(688), new BigNumber(805))).length)), new BigNumber(52)), _21_v0)) {
            _coll0.push([(_21_v0).multipliedBy(new BigNumber(82)),new BigNumber((_dafny.Seq.of(true)).length)]);
          }
        }
        return _coll0;
      }()).length))))).Intersect((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("lql")).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(210)), function (_22_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      })).length), new BigNumber(765), new BigNumber(79))).Intersect((_module.D10.create_DC28(_dafny.Set.fromElements(new BigNumber(701)))).dtor_cf44));
    };
    static fm13(p0, globalState) {
      return ((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(583), new BigNumber(764))) {
          let _23_v0 = _compr_1;
          if (((new BigNumber(583)).isLessThanOrEqualTo(_23_v0)) && ((_23_v0).isLessThan(new BigNumber(764)))) {
            _coll1.push([(_23_v0).multipliedBy(new BigNumber((_dafny.Set.fromElements(new BigNumber(37))).length)),new BigNumber(350)]);
          }
        }
        return _coll1;
      }()).Merge(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-272), new BigNumber(931))) {
          let _24_v1 = _compr_2;
          if (((new BigNumber(-272)).isLessThanOrEqualTo(_24_v1)) && ((_24_v1).isLessThan(new BigNumber(931)))) {
            _coll2.push([(_24_v1).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(506)), function (_25_i0) {
              return new _dafny.CodePoint('p'.codePointAt(0));
            })).length)),new BigNumber(569)]);
          }
        }
        return _coll2;
      }())).Merge(((true) ? (function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(900), new BigNumber(332))) {
          let _26_v2 = _compr_3;
          if (((new BigNumber(900)).isLessThanOrEqualTo(_26_v2)) && ((_26_v2).isLessThan(new BigNumber(332)))) {
            _coll3.push([_module.__default.safeDivisionInt(_26_v2, new BigNumber(574)),new BigNumber(492)]);
          }
        }
        return _coll3;
      }()) : (function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(410),false)).Keys.Elements) {
          let _27_v3 = _compr_4;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(410),false)).contains(_27_v3)) {
            _coll4.push([_module.__default.safeDivisionInt(_27_v3, new BigNumber(610)),new BigNumber(745)]);
          }
        }
        return _coll4;
      }())));
    };
    static fm14(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(_dafny.Set.fromElements(false))).Intersect(function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (_dafny.Seq.of(_dafny.Set.fromElements(!(!(true))), _dafny.Set.fromElements(true, !(false)))).Elements) {
          let _28_v0 = _compr_5;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(!(!(true))), _dafny.Set.fromElements(true, !(false))), _28_v0)) {
            _coll5.add(_28_v0);
          }
        }
        return _coll5;
      }());
    };
    static fm15(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(804)), function (_29_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })).length)).isEqualTo(new BigNumber(-960)),new BigNumber((_dafny.Seq.UnicodeFromString("xvbhie")).length));
    };
    static fm16(p0, globalState) {
      return ((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(724),new BigNumber((_dafny.Seq.of(new BigNumber(-52))).length))).Keys.Elements) {
          let _30_v0 = _compr_6;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(724),new BigNumber((_dafny.Seq.of(new BigNumber(-52))).length))).contains(_30_v0)) {
            _coll6.push([(_30_v0).multipliedBy(new BigNumber(628)),false]);
          }
        }
        return _coll6;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-925),false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(47),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(623),true)));
    };
    static fm17(p0, globalState) {
      return _dafny.Seq.of((_dafny.Set.fromElements(new BigNumber(33))).Union(_dafny.Set.fromElements(new BigNumber(-138))));
    };
    static fm18(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(!(!(true))), _dafny.Seq.of(false)), _dafny.Seq.Concat(_dafny.Seq.of(!(false), false, (_module.D1.create_DC5((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("kmamusrti")).length)), new BigNumber(983), false, function () {
  let _coll7 = new _dafny.Map();
  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(871), new BigNumber(44))) {
    let _31_v0 = _compr_7;
    if (((new BigNumber(871)).isLessThanOrEqualTo(_31_v0)) && ((_31_v0).isLessThan(new BigNumber(44)))) {
      _coll7.push([(_31_v0).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false))).cardinality())),_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)))]);
    }
  }
  return _coll7;
}())).dtor_cf9, (_module.D0.create_DC0(!(true))).dtor_cf0), _dafny.Seq.of(false, false, false)));
    };
    static fm19(globalState) {
      return (_dafny.ZERO).minus(((!_dafny.areEqual(_dafny.Seq.of(false), _dafny.Seq.of(true, false))) ? (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-775)), function (_32_i0) {
        return _dafny.MultiSet.fromElements(false);
      }), _dafny.Seq.of(_dafny.MultiSet.fromElements(true)))).length)) : (_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ltc")).length)), new BigNumber(236)))));
    };
    static fm20(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('d'.codePointAt(0));
    };
    static fm21(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D1.create_DC3(_dafny.Seq.UnicodeFromString("yohqnt"));
      if (_source1.is_DC4) {
        let _33___mcc_h0 = (_source1).cf5;
        let _34___mcc_h1 = (_source1).cf6;
        let _35_cf6 = _34___mcc_h1;
        let _36_cf5 = _33___mcc_h0;
        return false;
      } else if (_source1.is_DC5) {
        let _37___mcc_h2 = (_source1).cf7;
        let _38___mcc_h3 = (_source1).cf8;
        let _39___mcc_h4 = (_source1).cf9;
        let _40___mcc_h5 = (_source1).cf10;
        let _41_cf10 = _40___mcc_h5;
        let _42_cf9 = _39___mcc_h4;
        let _43_cf8 = _38___mcc_h3;
        let _44_cf7 = _37___mcc_h2;
        return (new BigNumber((_dafny.MultiSet.fromElements(_43_cf8, _43_cf8, _43_cf8, _44_cf7)).cardinality())).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(237), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_44_cf7,(_dafny.ZERO).minus(_43_cf8))).length))).cardinality()));
      } else {
        let _45___mcc_h6 = (_source1).cf4;
        let _46_cf4 = _45___mcc_h6;
        return (new BigNumber((_dafny.Seq.of(new BigNumber((_46_cf4).length), new BigNumber(390), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()), new BigNumber((_46_cf4).length), new BigNumber((_dafny.Seq.of(true)).length))).length)).isLessThanOrEqualTo(new BigNumber(-744));
      }
    };
    static fm22(globalState) {
      return _dafny.MultiSet.fromElements(!(true) || (true));
    };
    static fm23(p0, p1, globalState) {
      let _source2 = _module.D5.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(738)));
      if (_source2.is_DC14) {
        let _47___mcc_h0 = (_source2).cf25;
        let _48___mcc_h1 = (_source2).cf26;
        let _49___mcc_h2 = (_source2).cf27;
        let _50_cf27 = _49___mcc_h2;
        let _51_cf26 = _48___mcc_h1;
        let _52_cf25 = _47___mcc_h0;
        return _module.D3.create_DC9(_52_cf25);
      } else if (_source2.is_DC13) {
        let _53___mcc_h3 = (_source2).cf24;
        let _54_cf24 = _53___mcc_h3;
        return _module.D3.create_DC9((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length)))).cardinality())));
      } else {
        let _55___mcc_h4 = (_source2).cf28;
        let _56_cf28 = _55___mcc_h4;
        return _module.D3.create_DC9(new BigNumber(218));
      }
    };
    static m0(p0, globalState) {
      let _57_v0;
      _57_v0 = _dafny.Seq.UnicodeFromString("ir");
      let _58_v1;
      _58_v1 = new BigNumber(671);
      _57_v0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("drtgu"), _module.__default.fm9((_dafny.ZERO).minus(_58_v1), globalState)), _dafny.Seq.Concat(_57_v0, _57_v0));
      let _59_v2;
      _59_v2 = true;
      let _60_v3;
      _60_v3 = _dafny.MultiSet.fromElements(!(_59_v2));
      let _61_v4;
      _61_v4 = _dafny.Map.Empty.slice().updateUnsafe(_59_v2,_60_v3);
      let _62_v5;
      _62_v5 = _dafny.Seq.of(((((_61_v4).contains(_59_v2)) ? ((_61_v4).get(_59_v2)) : (_60_v3))).update(false, _module.__default.abs(_58_v1)));
      let _63_v6;
      _63_v6 = _dafny.MultiSet.fromElements(new BigNumber((_57_v0).length));
      let _64_v7;
      _64_v7 = _dafny.Map.Empty.slice().updateUnsafe(_63_v6,false);
      let _65_v8;
      _65_v8 = _dafny.Map.Empty.slice().updateUnsafe(_58_v1,new BigNumber((_64_v7).length));
      let _66_v9;
      _66_v9 = _dafny.Seq.of(_59_v2);
      let _67_v10;
      _67_v10 = _module.D2.create_DC7(_59_v2, false, _58_v1);
      let _68_v11;
      let _nw0 = Array((new BigNumber(20)).toNumber());
      _nw0[(_dafny.ZERO).toNumber()] = _module.__default.fm22(globalState);
      _nw0[(_dafny.ONE).toNumber()] = _60_v3;
      _nw0[(new BigNumber(2)).toNumber()] = (_module.__default.fm22(globalState)).update(_59_v2, _module.__default.abs(_58_v1));
      _nw0[(new BigNumber(3)).toNumber()] = (_62_v5)[_module.__default.safeIndex(_58_v1, new BigNumber((_62_v5).length))];
      _nw0[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(_59_v2, _59_v2, !(_module.__default.fm21(_58_v1, _59_v2, _59_v2, _65_v8, globalState)));
      _nw0[(new BigNumber(5)).toNumber()] = _60_v3;
      _nw0[(new BigNumber(6)).toNumber()] = _60_v3;
      _nw0[(new BigNumber(7)).toNumber()] = _60_v3;
      _nw0[(new BigNumber(8)).toNumber()] = _60_v3;
      _nw0[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(_59_v2, _59_v2));
      _nw0[(new BigNumber(10)).toNumber()] = _60_v3;
      _nw0[(new BigNumber(11)).toNumber()] = _60_v3;
      _nw0[(new BigNumber(12)).toNumber()] = (_60_v3).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.update(_66_v9, _module.__default.safeIndex(_58_v1, new BigNumber((_66_v9).length)), (_67_v10).dtor_cf12), _module.__default.safeIndex(_58_v1, new BigNumber((_dafny.Seq.update(_66_v9, _module.__default.safeIndex(_58_v1, new BigNumber((_66_v9).length)), (_67_v10).dtor_cf12)).length)), _59_v2)));
      _nw0[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(!(_59_v2)));
      _nw0[(new BigNumber(14)).toNumber()] = (_60_v3).Difference(_module.__default.fm22(globalState));
      _nw0[(new BigNumber(15)).toNumber()] = (_dafny.MultiSet.fromElements(_59_v2, _59_v2, _59_v2)).Union((_60_v3).update(_59_v2, _module.__default.abs(new BigNumber((_57_v0).length))));
      _nw0[(new BigNumber(16)).toNumber()] = (_60_v3).update(_59_v2, _module.__default.abs(_58_v1));
      _nw0[(new BigNumber(17)).toNumber()] = _60_v3;
      _nw0[(new BigNumber(18)).toNumber()] = _60_v3;
      _nw0[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.FromArray(_66_v9);
      _68_v11 = _nw0;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_68_v11).length))) {
        let _69_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_69_i0)) && ((_69_i0).isLessThan(new BigNumber((_68_v11).length))))) {
          (_68_v11)[(_69_i0)] = ((_59_v2) ? (_60_v3) : ((_60_v3).update(_59_v2, _module.__default.abs(_58_v1))));
        }
      }
      let _70_v12;
      let _nw1 = new _module.C0();
      _nw1.__ctor(_59_v2, _59_v2, _dafny.Seq.UnicodeFromString("jtxlh"), _58_v1, _59_v2);
      _70_v12 = _nw1;
      let _71_v13;
      _71_v13 = _dafny.Map.Empty.slice().updateUnsafe(_70_v12,_58_v1);
      (globalState).f11 = (((((_71_v13).contains(_70_v12)) ? ((_71_v13).get(_70_v12)) : (_70_v12.f14))).multipliedBy(_70_v12.f14)).plus(_70_v12.f14);
      _59_v2 = _59_v2;
      let _source3 = _module.D0.create_DC2(_70_v12.f14, (_dafny.MultiSet.fromElements(true, _59_v2, (_70_v12).f15, _59_v2)).Difference(_dafny.MultiSet.FromArray(_66_v9)));
      if (_source3.is_DC1) {
        let _72___mcc_h0 = (_source3).cf1;
        let _73_cf1 = _72___mcc_h0;
        let _74_v14;
        _74_v14 = _dafny.Map.Empty.slice().updateUnsafe(true,(_70_v12).f15);
        let _75_v15;
        let _nw2 = new _module.C0();
        _nw2.__ctor(_59_v2, ((((_74_v14).contains(_59_v2)) ? ((_74_v14).get(_59_v2)) : ((_70_v12).f15))) === ((_70_v12).f15), _70_v12.f16, _58_v1, (_70_v12.f14).isLessThanOrEqualTo(_70_v12.f14));
        _75_v15 = _nw2;
        _75_v15 = _75_v15;
        _58_v1 = (_dafny.ZERO).minus((((_dafny.MultiSet.fromElements(_58_v1)).IsSubsetOf(_63_v6)) ? ((_dafny.ZERO).minus(_73_cf1)) : (_70_v12.f14)));
        let _76_v16;
        _76_v16 = _dafny.Map.Empty.slice().updateUnsafe((_70_v12).f15,_70_v12);
        let _77_v17;
        let _nw3 = Array((_dafny.ONE).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = _75_v15;
        _77_v17 = _nw3;
        let _78_v18;
        _78_v18 = _dafny.MultiSet.fromElements(_77_v17, _77_v17, _77_v17, _77_v17);
        (globalState).f11 = ((_59_v2) ? (_module.__default.safeModuloInt(_70_v12.f14, new BigNumber((_76_v16).length))) : ((_dafny.ZERO).minus(new BigNumber(((_78_v18).update(_77_v17, _module.__default.abs(_70_v12.f14))).cardinality()))));
        _73_cf1 = (((_60_v3).contains((_75_v15).f17)) ? ((_60_v3).get((_75_v15).f17)) : (_73_cf1));
      } else if (_source3.is_DC2) {
        let _79___mcc_h1 = (_source3).cf2;
        let _80___mcc_h2 = (_source3).cf3;
        let _81_cf3 = _80___mcc_h2;
        let _82_cf2 = _79___mcc_h1;
        let _83_v19;
        let _nw4 = Array((new BigNumber(6)).toNumber()).fill(false);
        _83_v19 = _nw4;
        let _84_v20;
        _84_v20 = _dafny.Map.Empty.slice().updateUnsafe((_70_v12).f15,_83_v19);
        let _85_v21;
        _85_v21 = _dafny.Map.Empty.slice().updateUnsafe((_70_v12).fm0((_70_v12).f15, globalState),_59_v2);
        if (_module.__default.fm21(new BigNumber((_84_v20).length), (((_85_v21).contains((_70_v12).f15)) ? ((_85_v21).get((_70_v12).f15)) : ((_70_v12).f15)), (_70_v12).f15, _65_v8, globalState)) {
          (globalState).f11 = (_70_v12.f14).minus(_58_v1);
          let _86_v22;
          _86_v22 = _module.D3.create_DC9(_82_cf2);
          let _87_v23;
          let _nw5 = new _module.C1();
          _nw5.__ctor(_70_v12.f16, _66_v9, new BigNumber((_60_v3).cardinality()), (_70_v12).f15);
          _87_v23 = _nw5;
          let _88_v24;
          _88_v24 = _dafny.Map.Empty.slice().updateUnsafe((_70_v12).f15,_87_v23);
          _86_v22 = _module.__default.fm23(_dafny.Seq.of(_70_v12.f14, new BigNumber((_88_v24).length), _70_v12.f14), !(_59_v2), globalState);
          let _89_v25;
          let _init0 = ((_90_v12, _91_cf3, _92_cf2) => function (_93_i1) {
            return _dafny.Seq.of((((_91_cf3).contains((_90_v12).f15)) ? ((_91_cf3).get((_90_v12).f15)) : (_90_v12.f14)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(834),_92_cf2)).length), _90_v12.f14);
          })(_70_v12, _81_cf3, _82_cf2);
          let _nw6 = Array((new BigNumber(22)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw6.length); _i0_0++) {
            _nw6[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _89_v25 = _nw6;
          let _rhs0 = _70_v12.f14;
          let _rhs1 = !(_63_v6).contains((new BigNumber((_dafny.Seq.UnicodeFromString("yitpb")).length)).minus(_58_v1));
          let _rhs2 = !((_70_v12.f14).isLessThan((new BigNumber(949)).plus(new BigNumber(-463))));
          let _rhs3 = _57_v0;
          let _rhs4 = ((_59_v2) ? (_89_v25) : (_89_v25));
          let _lhs0 = globalState;
          let _lhs1 = globalState;
          _lhs0.f11 = _rhs0;
          _lhs1.f0 = _rhs1;
          _59_v2 = _rhs2;
          _57_v0 = _rhs3;
          _89_v25 = _rhs4;
          let _94_v26;
          _94_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(562),_59_v2);
          let _95_v27;
          _95_v27 = _module.D8.create_DC21(_94_v26);
          _94_v26 = ((_95_v27).dtor_cf36).Merge(_94_v26);
          let _96_v28;
          let _nw7 = Array((new BigNumber(28)).toNumber());
          _96_v28 = _nw7;
          _96_v28 = _96_v28;
        } else {
          (_70_v12).f16 = _dafny.Seq.Concat(_70_v12.f16, _70_v12.f16);
          let _index0 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_83_v19).length));
          (_83_v19)[_index0] = _module.__default.fm21(_70_v12.f14, false, (_70_v12).f15, _65_v8, globalState);
          let _index1 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_83_v19).length));
          (_83_v19)[_index1] = (_81_cf3).IsDisjointFrom(_60_v3);
          let _97_v29;
          let _init1 = ((_98_v9) => function (_99_i2) {
            return _98_v9;
          })(_66_v9);
          let _nw8 = Array((new BigNumber(24)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw8.length); _i0_1++) {
            _nw8[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _97_v29 = _nw8;
          let _100_v30;
          _100_v30 = _module.D5.create_DC14(_70_v12.f14, _97_v29, true);
          let _101_v31;
          let _nw9 = Array((new BigNumber(17)).toNumber());
          _nw9[(_dafny.ZERO).toNumber()] = (_70_v12).f15;
          _nw9[(_dafny.ONE).toNumber()] = false;
          _nw9[(new BigNumber(2)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(3)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(4)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(5)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(6)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(7)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(8)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(9)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(10)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(11)).toNumber()] = !((_100_v30).dtor_cf27);
          _nw9[(new BigNumber(12)).toNumber()] = !(!((_70_v12).f15));
          _nw9[(new BigNumber(13)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(14)).toNumber()] = (_70_v12).f15;
          _nw9[(new BigNumber(15)).toNumber()] = !(!(_59_v2));
          _nw9[(new BigNumber(16)).toNumber()] = (_70_v12).f15;
          _101_v31 = _nw9;
          let _102_v32;
          _102_v32 = _dafny.Map.Empty.slice().updateUnsafe(_101_v31,new BigNumber(420));
          _102_v32 = (_102_v32).update(_101_v31, new BigNumber((_63_v6).cardinality()));
          let _103_v33;
          _103_v33 = _dafny.Set.fromElements((_70_v12).f15, _59_v2);
          let _104_v34;
          _104_v34 = new _dafny.CodePoint('y'.codePointAt(0));
          let _105_v35;
          _105_v35 = _dafny.MultiSet.fromElements(_104_v34, _104_v34);
          let _pat_let_tv0 = _70_v12;
          let _pat_let_tv1 = _83_v19;
          let _pat_let_tv2 = _83_v19;
          let _pat_let_tv3 = _83_v19;
          let _pat_let_tv4 = _83_v19;
          let _pat_let_tv5 = globalState;
          let _pat_let_tv6 = globalState;
          let _pat_let_tv7 = _70_v12;
          let _106_v36;
          let _nw10 = new _module.C2();
          _nw10.__ctor(function (_pat_let0_0) {
            return function (_107_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_108_dt__update_hcf12_h0) {
                  return function (_pat_let2_0) {
                    return function (_109_dt__update_hcf13_h0) {
                      return _module.D2.create_DC7(_108_dt__update_hcf12_h0, _109_dt__update_hcf13_h0, (_107_dt__update__tmp_h0).dtor_cf14);
                    }(_pat_let2_0);
                  }((_pat_let_tv7).f15);
                }(_pat_let1_0);
              }(_module.__default.fm21(_pat_let_tv0.f14, (_pat_let_tv2)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_pat_let_tv1).length))], (_pat_let_tv4)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_pat_let_tv3).length))], _module.__default.fm13(new BigNumber((_dafny.Seq.UnicodeFromString("ffbfp")).length), _pat_let_tv5), _pat_let_tv6));
            }(_pat_let0_0);
          }(_module.D2.create_DC7(_module.__default.fm21(new BigNumber((_103_v33).length), (_70_v12).f15, (_70_v12).f15, _65_v8, globalState), (_70_v12).f15, _70_v12.f14)), (((_105_v35).contains(_104_v34)) ? ((_105_v35).get(_104_v34)) : (_82_cf2)), (_70_v12).f15);
          _106_v36 = _nw10;
          let _110_v37;
          let _nw11 = new _module.C0();
          _nw11.__ctor(!((_70_v12).f15), (_70_v12).f15, _57_v0, _82_cf2, (_70_v12).f15);
          _110_v37 = _nw11;
          let _111_v38;
          _111_v38 = _dafny.Map.Empty.slice().updateUnsafe((_66_v9)[_module.__default.safeIndex(_82_cf2, new BigNumber((_66_v9).length))],_110_v37);
          (globalState).f11 = new BigNumber((((_111_v38).Merge(_111_v38)).Merge(_111_v38)).length);
        }
        let _112_v39;
        let _nw12 = new _module.C1();
        _nw12.__ctor(_57_v0, _dafny.Seq.of(_59_v2), _58_v1, (_70_v12).f15);
        _112_v39 = _nw12;
        let _113_v40;
        _113_v40 = _dafny.Map.Empty.slice().updateUnsafe(_112_v39,(_70_v12).f15);
        let _114_v41;
        _114_v41 = _dafny.Map.Empty.slice().updateUnsafe((_70_v12).f15,new BigNumber((_113_v40).length));
        _114_v41 = (_114_v41).update((_70_v12.f14).isLessThanOrEqualTo(_70_v12.f14), _70_v12.f14);
        let _115_v42;
        _115_v42 = _dafny.Set.fromElements(_112_v39, _112_v39, _112_v39);
        let _116_v43;
        _116_v43 = _module.D1.create_DC4(_module.__default.safeDivisionInt(new BigNumber((_115_v42).length), new BigNumber((_70_v12.f16).length)), _70_v12.f14);
        _116_v43 = _116_v43;
        let _117_v44;
        let _nw13 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _117_v44 = _nw13;
        let _118_v45;
        _118_v45 = new _dafny.CodePoint('s'.codePointAt(0));
        let _index2 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_117_v44).length));
        (_117_v44)[_index2] = _118_v45;
        let _index3 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_117_v44).length));
        (_117_v44)[_index3] = _118_v45;
      } else {
        let _119___mcc_h3 = (_source3).cf0;
        let _120_cf0 = _119___mcc_h3;
        _66_v9 = _dafny.Seq.Concat(_dafny.Seq.of(false, (_70_v12).f15, (_70_v12).fm0((_70_v12).f15, globalState)), _66_v9);
        let _121_v46;
        let _nw14 = new _module.C0();
        _nw14.__ctor(_59_v2, !((_70_v12).f15), _57_v0, _70_v12.f14, (_70_v12).f15);
        _121_v46 = _nw14;
        let _122_v47;
        _122_v47 = _dafny.Map.Empty.slice().updateUnsafe(_63_v6,_121_v46);
        let _123_v48;
        _123_v48 = _dafny.Map.Empty.slice().updateUnsafe(_59_v2,_70_v12.f14);
        let _124_v49;
        _124_v49 = _dafny.Map.Empty.slice().updateUnsafe(_58_v1,_123_v48);
        let _125_v50;
        let _nw15 = Array((new BigNumber(12)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = _121_v46;
        _nw15[(_dafny.ONE).toNumber()] = _121_v46;
        _nw15[(new BigNumber(2)).toNumber()] = _121_v46;
        _nw15[(new BigNumber(3)).toNumber()] = _121_v46;
        _nw15[(new BigNumber(4)).toNumber()] = _121_v46;
        _nw15[(new BigNumber(5)).toNumber()] = _121_v46;
        _nw15[(new BigNumber(6)).toNumber()] = _121_v46;
        _nw15[(new BigNumber(7)).toNumber()] = _121_v46;
        _nw15[(new BigNumber(8)).toNumber()] = (((_70_v12).f15) ? (_121_v46) : (_121_v46));
        _nw15[(new BigNumber(9)).toNumber()] = _121_v46;
        _nw15[(new BigNumber(10)).toNumber()] = _121_v46;
        _nw15[(new BigNumber(11)).toNumber()] = (((_122_v47).contains(_dafny.MultiSet.fromElements(new BigNumber(((((_124_v49).contains(new BigNumber(512))) ? ((_124_v49).get(new BigNumber(512))) : (_123_v48))).length)))) ? ((_122_v47).get(_dafny.MultiSet.fromElements(new BigNumber(((((_124_v49).contains(new BigNumber(512))) ? ((_124_v49).get(new BigNumber(512))) : (_123_v48))).length)))) : (_121_v46));
        _125_v50 = _nw15;
        let _126_v51;
        _126_v51 = _module.D9.create_DC25(_125_v50);
        _125_v50 = (_126_v51).dtor_cf41;
        let _127_v52;
        _127_v52 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(globalState),(_70_v12.f14).isEqualTo(_70_v12.f14));
        let _128_v53;
        _128_v53 = _dafny.MultiSet.fromElements(_121_v46, _121_v46);
        let _129_v54;
        _129_v54 = _dafny.Seq.of(_58_v1);
        _120_cf0 = (((_127_v52).contains(((_120_cf0) ? (_58_v1) : ((((_128_v53).contains(_121_v46)) ? ((_128_v53).get(_121_v46)) : (new BigNumber((_129_v54).length))))))) ? ((_127_v52).get(((_120_cf0) ? (_58_v1) : ((((_128_v53).contains(_121_v46)) ? ((_128_v53).get(_121_v46)) : (new BigNumber((_129_v54).length))))))) : (_120_cf0));
        let _130_v55;
        let _init2 = ((_131_v46) => function (_132_i3) {
          return (_131_v46).f18;
        })(_121_v46);
        let _nw16 = Array((_dafny.ONE).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw16.length); _i0_2++) {
          _nw16[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _130_v55 = _nw16;
        _130_v55 = _130_v55;
      }
      let _133_v56;
      let _nw17 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _133_v56 = _nw17;
      let _index4 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_133_v56).length));
      (_133_v56)[_index4] = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(!((_70_v12).f15), _59_v2), _dafny.Seq.of(_59_v2, (_66_v9)[_module.__default.safeIndex(_58_v1, new BigNumber((_66_v9).length))], (_70_v12).f15, _59_v2)))).cardinality());
      let _index5 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_133_v56).length));
      (_133_v56)[_index5] = ((_dafny.ZERO).minus((_70_v12.f14).plus(new BigNumber((_66_v9).length)))).plus((_dafny.ZERO).minus(_58_v1));
      return;
    }
    static Main(__noArgsParameter) {
      let _134_v0;
      _134_v0 = false;
      let _135_v1;
      _135_v1 = _dafny.Seq.UnicodeFromString("ogvv");
      let _136_v2;
      _136_v2 = _dafny.Seq.of(_134_v0);
      let _137_v3;
      _137_v3 = new BigNumber(329);
      let _138_v4;
      _138_v4 = _dafny.Seq.of(_134_v0, _134_v0, _134_v0, (_136_v2)[_module.__default.safeIndex((_dafny.ZERO).minus(_137_v3), new BigNumber((_136_v2).length))], !(_134_v0));
      let _139_globalState;
      let _nw18 = new _module.GlobalState();
      _nw18.__ctor(false, true, new BigNumber(140), true, ((_134_v0) ? (_135_v1) : (_135_v1)), false, new BigNumber(188), true, new BigNumber(38), new BigNumber(937), new BigNumber(439), new BigNumber(8), false, _136_v2);
      _139_globalState = _nw18;
      if (_134_v0) {
        let _140_v5;
        _140_v5 = _dafny.Set.fromElements(_134_v0, _134_v0);
        let _141_v6;
        _141_v6 = _dafny.Seq.of(_140_v5);
        (_139_globalState).f0 = !_dafny.areEqual(_141_v6, _dafny.Seq.of(_140_v5, _140_v5, _140_v5, _140_v5));
        let _142_v7;
        _142_v7 = _dafny.Map.Empty.slice().updateUnsafe(_137_v3,!(_134_v0));
        let _143_v8;
        let _nw19 = Array((new BigNumber(20)).toNumber()).fill(false);
        _143_v8 = _nw19;
        _module.__default.m0(_dafny.Map.Empty.slice().updateUnsafe(_142_v7,_143_v8), _139_globalState);
        let _144_v9;
        _144_v9 = _dafny.Map.Empty.slice().updateUnsafe(_137_v3,_135_v1);
        _144_v9 = (_144_v9).update(_137_v3, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hkrlymcrd"), _135_v1));
        let _145_v10;
        let _nw20 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _145_v10 = _nw20;
        let _index6 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_145_v10).length));
        (_145_v10)[_index6] = _137_v3;
        let _146_v11;
        _146_v11 = _dafny.MultiSet.fromElements(_137_v3, new BigNumber(378));
        let _147_v12;
        _147_v12 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(((_146_v11).Intersect((_146_v11).update(new BigNumber((_dafny.Seq.of(_137_v3)).length), _module.__default.abs(_137_v3)))).cardinality()));
        let _148_v13;
        _148_v13 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_147_v12);
        let _index7 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_145_v10).length));
        let _rhs5 = true;
        let _rhs6 = _140_v5;
        let _rhs7 = _137_v3;
        let _rhs8 = (_dafny.Map.Empty.slice().updateUnsafe(false,_137_v3)).Merge(((((_148_v13).contains(!(_134_v0))) ? ((_148_v13).get(!(_134_v0))) : (_147_v12))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_134_v0,new BigNumber((_135_v1).length))));
        let _lhs2 = _145_v10;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_145_v10).length));
        _134_v0 = _rhs5;
        _140_v5 = _rhs6;
        _lhs2[_lhs3] = _rhs7;
        _147_v12 = _rhs8;
        let _149_v14;
        _149_v14 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_145_v10)[_module.__default.safeIndex(new BigNumber(759), new BigNumber((_145_v10).length))],_134_v0),_143_v8);
        _module.__default.m0(_149_v14, _139_globalState);
      } else {
        let _150_v15;
        _150_v15 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_137_v3),_134_v0);
        let _151_v16;
        let _init3 = ((_152_v0) => function (_153_i0) {
          return !(_152_v0);
        })(_134_v0);
        let _nw21 = Array((new BigNumber(7)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw21.length); _i0_3++) {
          _nw21[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _151_v16 = _nw21;
        let _154_v17;
        _154_v17 = _dafny.Map.Empty.slice().updateUnsafe(_150_v15,_151_v16);
        _module.__default.m0(_154_v17, _139_globalState);
        (_139_globalState).f11 = _137_v3;
        if ((((_150_v15).contains(_137_v3)) ? ((_150_v15).get(_137_v3)) : (_134_v0))) {
          let _155_v18;
          _155_v18 = _dafny.MultiSet.fromElements(_137_v3, _137_v3, (_dafny.ZERO).minus(_137_v3));
          _134_v0 = (((_150_v15).contains(new BigNumber((_155_v18).cardinality()))) ? ((_150_v15).get(new BigNumber((_155_v18).cardinality()))) : (_134_v0));
          _module.__default.m0((_154_v17).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_137_v3,_134_v0),_151_v16)), _139_globalState);
          let _156_v19;
          let _nw22 = new _module.C0();
          _nw22.__ctor(_134_v0, !((_134_v0) || (_134_v0)), _135_v1, new BigNumber(385), true);
          _156_v19 = _nw22;
          let _157_v21;
          _157_v21 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),_137_v3);
          (_139_globalState).f11 = new BigNumber((function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of (_157_v21).Keys.Elements) {
              let _158_v20 = _compr_8;
              if ((_157_v21).contains(_158_v20)) {
                _coll8.push([_158_v20,(_156_v19).f17]);
              }
            }
            return _coll8;
          }()).length);
          let _159_v22;
          _159_v22 = _module.D0.create_DC0(true);
          (_139_globalState).f0 = (_159_v22).dtor_cf0;
        } else {
          let _160_v23;
          _160_v23 = new _dafny.CodePoint('r'.codePointAt(0));
          let _161_v24;
          _161_v24 = _dafny.Seq.of(new BigNumber(382));
          let _162_v25;
          _162_v25 = _dafny.Map.Empty.slice().updateUnsafe(true,_160_v23);
          _160_v23 = (_135_v1)[_module.__default.safeIndex((_161_v24)[_module.__default.safeIndex(new BigNumber((_162_v25).length), new BigNumber((_161_v24).length))], new BigNumber((_135_v1).length))];
          _module.__default.m0((_154_v17).Merge(_154_v17), _139_globalState);
          let _163_v26;
          _163_v26 = _dafny.MultiSet.fromElements(_134_v0);
          let _rhs9 = !(_134_v0);
          let _rhs10 = (_137_v3).minus(((_134_v0) ? (new BigNumber((_136_v2).length)) : (_137_v3)));
          let _rhs11 = _137_v3;
          let _rhs12 = !(((_dafny.MultiSet.fromElements(_134_v0)).Difference(_163_v26)).IsSubsetOf(_163_v26));
          let _lhs4 = _139_globalState;
          let _lhs5 = _139_globalState;
          let _lhs6 = _139_globalState;
          _lhs4.f0 = _rhs9;
          _137_v3 = _rhs10;
          _lhs5.f11 = _rhs11;
          _lhs6.f3 = _rhs12;
          let _164_v27;
          _164_v27 = _dafny.Map.Empty.slice().updateUnsafe((_137_v3).plus((_dafny.ZERO).minus(_137_v3)),(_dafny.MultiSet.fromElements(_134_v0, _134_v0)).update(_134_v0, _module.__default.abs(_137_v3)));
          _164_v27 = (_module.D2.create_DC6(_164_v27)).dtor_cf11;
          _module.__default.m0(_dafny.Map.Empty.slice().updateUnsafe(_150_v15,_151_v16), _139_globalState);
        }
        (_139_globalState).f3 = true;
        (_139_globalState).f3 = !(_134_v0);
      }
      let _165_v28;
      let _init4 = ((_166_v0) => function (_167_i2) {
        return ((_166_v0) ? (_166_v0) : (_166_v0));
      })(_134_v0);
      let _nw23 = Array((new BigNumber(7)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw23.length); _i0_4++) {
        _nw23[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _165_v28 = _nw23;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_165_v28).length))) {
        let _168_i1 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_168_i1)) && ((_168_i1).isLessThan(new BigNumber((_165_v28).length))))) {
          (_165_v28)[(_168_i1)] = _134_v0;
        }
      }
      let _169_v29;
      _169_v29 = _dafny.Seq.of(_dafny.Seq.of(_134_v0), _dafny.Seq.update(_dafny.Seq.Concat(_138_v4, _138_v4), _module.__default.safeIndex((_dafny.ZERO).minus(_137_v3), new BigNumber((_dafny.Seq.Concat(_138_v4, _138_v4)).length)), _134_v0), _136_v2);
      _169_v29 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), ((_170_v4) => function (_171_i3) {
        return _170_v4;
      })(_138_v4)), _dafny.Seq.of(_138_v4, _dafny.Seq.update(_136_v2, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("tknb")).length), new BigNumber((_136_v2).length)), _134_v0), _138_v4)), _169_v29);
      _137_v3 = (_module.__default.safeModuloInt(_137_v3, _137_v3)).plus(_137_v3);
      if ((_137_v3).isLessThanOrEqualTo(_137_v3)) {
        let _172_v30;
        let _init5 = ((_173_v4) => function (_174_i4) {
          return _module.__default.safeModuloInt(_174_i4, new BigNumber((_173_v4).length));
        })(_138_v4);
        let _nw24 = Array((new BigNumber(9)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw24.length); _i0_5++) {
          _nw24[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _172_v30 = _nw24;
        let _175_v31;
        let _nw25 = new _module.C1();
        _nw25.__ctor(_dafny.Seq.UnicodeFromString("qsgjrgt"), _138_v4, _137_v3, _134_v0);
        _175_v31 = _nw25;
        let _176_v32;
        _176_v32 = _dafny.Map.Empty.slice().updateUnsafe(_175_v31,(_175_v31).f15);
        let _177_v33;
        let _nw26 = new _module.C0();
        _nw26.__ctor(true, _134_v0, _135_v1, _137_v3, (_138_v4)[_module.__default.safeIndex(new BigNumber((_176_v32).length), new BigNumber((_138_v4).length))]);
        _177_v33 = _nw26;
        _137_v3 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_172_v30,_177_v33)).length);
        let _178_v34;
        _178_v34 = _dafny.MultiSet.fromElements((_175_v31.f14).multipliedBy(_137_v3));
        _178_v34 = _dafny.MultiSet.fromElements(_177_v33.f14, (_175_v31.f14).multipliedBy((_175_v31).fm1(_178_v34, !((_175_v31).f15), _139_globalState)), _175_v31.f14, new BigNumber(297));
        let _179_v35;
        _179_v35 = _module.D1.create_DC4(_137_v3, (_dafny.ZERO).minus(_175_v31.f14));
        let _source4 = _179_v35;
        if (_source4.is_DC4) {
          let _180___mcc_h0 = (_source4).cf5;
          let _181___mcc_h1 = (_source4).cf6;
          let _182_cf6 = _181___mcc_h1;
          let _183_cf5 = _180___mcc_h0;
          let _184_v36;
          _184_v36 = _module.D0.create_DC0(_134_v0);
          let _185_v37;
          _185_v37 = _module.D2.create_DC7(!((_177_v33).f15), (_177_v33).f15, (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_184_v36, _184_v36, _184_v36)).cardinality())));
          let _186_v38;
          let _nw27 = new _module.C2();
          _nw27.__ctor(_185_v37, (_177_v33.f14).multipliedBy(_175_v31.f14), !(true));
          _186_v38 = _nw27;
          let _187_v39;
          let _nw28 = new _module.C2();
          _nw28.__ctor(_module.D2.create_DC7(_134_v0, _134_v0, _175_v31.f14), _module.__default.safeDivisionInt(_177_v33.f14, _182_cf6), (_186_v38).fm0((_136_v2)[_module.__default.safeIndex(_183_cf5, new BigNumber((_136_v2).length))], _139_globalState));
          _187_v39 = _nw28;
          let _188_v40;
          let _nw29 = Array((new BigNumber(24)).toNumber()).fill(_module.D2.Default());
          _188_v40 = _nw29;
          let _index8 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_188_v40).length));
          (_188_v40)[_index8] = (_186_v38).f19;
          let _index9 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_188_v40).length));
          (_188_v40)[_index9] = (_187_v39).f19;
          _138_v4 = _138_v4;
        } else if (_source4.is_DC5) {
          let _189___mcc_h2 = (_source4).cf7;
          let _190___mcc_h3 = (_source4).cf8;
          let _191___mcc_h4 = (_source4).cf9;
          let _192___mcc_h5 = (_source4).cf10;
          let _193_cf10 = _192___mcc_h5;
          let _194_cf9 = _191___mcc_h4;
          let _195_cf8 = _190___mcc_h3;
          let _196_cf7 = _189___mcc_h2;
          let _197_v41;
          _197_v41 = _dafny.Map.Empty.slice().updateUnsafe(_134_v0,(_177_v33).f15);
          _197_v41 = (_197_v41).update((_175_v31).f15, (_177_v33).f15);
          let _198_v42;
          let _nw30 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
          _198_v42 = _nw30;
          let _199_v43;
          _199_v43 = _module.D5.create_DC14(_177_v33.f14, _198_v42, _134_v0);
          let _200_v44;
          _200_v44 = _dafny.MultiSet.fromElements(_199_v43, _199_v43, _199_v43, _199_v43, _199_v43);
          _200_v44 = ((_200_v44).update(_199_v43, _module.__default.abs((_dafny.ZERO).minus(_195_cf8)))).Union(_200_v44);
          (_139_globalState).f3 = _134_v0;
          let _201_v45;
          _201_v45 = _module.D6.create_DC17(true, _177_v33);
          (_139_globalState).f0 = (_201_v45).dtor_cf30;
        } else {
          let _202___mcc_h6 = (_source4).cf4;
          let _203_cf4 = _202___mcc_h6;
          let _204_v46;
          let _nw31 = Array((new BigNumber(10)).toNumber());
          _204_v46 = _nw31;
          let _205_v47;
          let _nw32 = new _module.C1();
          _nw32.__ctor(_203_cf4, _138_v4, _177_v33.f14, _134_v0);
          _205_v47 = _nw32;
          let _index10 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_204_v46).length));
          (_204_v46)[_index10] = _205_v47;
          let _index11 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_165_v28).length));
          (_165_v28)[_index11] = (_177_v33).f15;
          let _index12 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_204_v46).length));
          let _index13 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_165_v28).length));
          let _rhs13 = _205_v47;
          let _rhs14 = (_175_v31).f15;
          let _rhs15 = _175_v31.f14;
          let _lhs7 = _204_v46;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_204_v46).length));
          let _lhs9 = _165_v28;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_165_v28).length));
          let _lhs11 = _139_globalState;
          _lhs7[_lhs8] = _rhs13;
          _lhs9[_lhs10] = _rhs14;
          _lhs11.f11 = _rhs15;
          let _206_v48;
          _206_v48 = _dafny.Seq.of(_177_v33);
          let _207_v49;
          _207_v49 = _dafny.Map.Empty.slice().updateUnsafe(_134_v0,new BigNumber((_206_v48).length));
          (_175_v31).f14 = _module.__default.safeModuloInt(new BigNumber((_207_v49).length), _177_v33.f14);
          let _index14 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_172_v30).length));
          (_172_v30)[_index14] = _177_v33.f14;
          let _208_v50;
          _208_v50 = new _dafny.CodePoint('c'.codePointAt(0));
          let _209_v51;
          _209_v51 = _dafny.Set.fromElements(_208_v50, _208_v50);
          let _210_v52;
          let _nw33 = Array((new BigNumber(2)).toNumber());
          _nw33[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_208_v50);
          _nw33[(_dafny.ONE).toNumber()] = _209_v51;
          _210_v52 = _nw33;
          let _index15 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_210_v52).length));
          (_210_v52)[_index15] = _209_v51;
          let _211_v53;
          let _nw34 = Array((new BigNumber(28)).toNumber());
          _211_v53 = _nw34;
          let _212_v54;
          let _nw35 = new _module.C0();
          _nw35.__ctor(_134_v0, (_175_v31).f15, _203_cf4, _175_v31.f14, _134_v0);
          _212_v54 = _nw35;
          let _index16 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_211_v53).length));
          (_211_v53)[_index16] = _212_v54;
          let _index17 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_172_v30).length));
          let _index18 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_210_v52).length));
          let _index19 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_211_v53).length));
          let _rhs16 = _137_v3;
          let _rhs17 = (_dafny.Set.fromElements(_208_v50)).Union(_209_v51);
          let _rhs18 = _212_v54;
          let _lhs12 = _172_v30;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_172_v30).length));
          let _lhs14 = _210_v52;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_210_v52).length));
          let _lhs16 = _211_v53;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_211_v53).length));
          _lhs12[_lhs13] = _rhs16;
          _lhs14[_lhs15] = _rhs17;
          _lhs16[_lhs17] = _rhs18;
          let _index20 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_165_v28).length));
          (_165_v28)[_index20] = _134_v0;
        }
        let _213_v55;
        let _init6 = function (_214_i5) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        };
        let _nw36 = Array((new BigNumber(3)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw36.length); _i0_6++) {
          _nw36[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _213_v55 = _nw36;
        let _215_v56;
        _215_v56 = new _dafny.CodePoint('y'.codePointAt(0));
        let _index21 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_213_v55).length));
        (_213_v55)[_index21] = (((_175_v31).f15) ? (_215_v56) : (new _dafny.CodePoint('t'.codePointAt(0))));
        let _216_v57;
        let _init7 = function (_217_i6) {
          return _dafny.Seq.UnicodeFromString("esgshe");
        };
        let _nw37 = Array((new BigNumber(15)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw37.length); _i0_7++) {
          _nw37[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _216_v57 = _nw37;
        let _218_v58;
        _218_v58 = _216_v57;
        let _219_v59;
        _219_v59 = _module.D6.create_DC18(_175_v31.f14, _215_v56);
        let _index22 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_172_v30).length));
        (_172_v30)[_index22] = (_219_v59).dtor_cf32;
        let _index23 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_213_v55).length));
        let _index24 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_172_v30).length));
        let _rhs19 = (_175_v31).f15;
        let _rhs20 = _215_v56;
        let _rhs21 = _216_v57;
        let _rhs22 = _177_v33.f14;
        let _rhs23 = (_dafny.ZERO).minus(new BigNumber(-341));
        let _lhs18 = _213_v55;
        let _lhs19 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_213_v55).length));
        let _lhs20 = _175_v31;
        let _lhs21 = _172_v30;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_172_v30).length));
        _134_v0 = _rhs19;
        _lhs18[_lhs19] = _rhs20;
        _218_v58 = _rhs21;
        _lhs20.f14 = _rhs22;
        _lhs21[_lhs22] = _rhs23;
        (_139_globalState).f3 = _134_v0;
      } else {
        (_139_globalState).f11 = _137_v3;
        let _220_v60;
        let _nw38 = new _module.C1();
        _nw38.__ctor(_dafny.Seq.UnicodeFromString("daadh"), _136_v2, _137_v3, _134_v0);
        _220_v60 = _nw38;
        let _221_v61;
        _221_v61 = _dafny.Map.Empty.slice().updateUnsafe(_220_v60,_134_v0);
        _221_v61 = (_221_v61).update(_220_v60, _134_v0);
        if ((_134_v0) || (!(_137_v3).isEqualTo(new BigNumber(612)))) {
          let _222_v62;
          _222_v62 = new _dafny.CodePoint('v'.codePointAt(0));
          let _223_v63;
          _223_v63 = _dafny.Set.fromElements((_dafny.ZERO).minus(_137_v3));
          let _224_v64;
          _224_v64 = _dafny.Seq.of(_module.__default.fm12(_137_v3, (_220_v60).fm0(_134_v0, _139_globalState), new BigNumber(699), _222_v62, _139_globalState), _223_v63);
          let _225_v65;
          let _nw39 = Array((new BigNumber(5)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_module.__default.fm17(_137_v3, _139_globalState), _224_v64);
          _nw39[(_dafny.ONE).toNumber()] = _224_v64;
          _nw39[(new BigNumber(2)).toNumber()] = _224_v64;
          _nw39[(new BigNumber(3)).toNumber()] = _224_v64;
          _nw39[(new BigNumber(4)).toNumber()] = _224_v64;
          _225_v65 = _nw39;
          let _index25 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_225_v65).length));
          (_225_v65)[_index25] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(188)), ((_226_v3) => function (_227_i7) {
            return _dafny.Set.fromElements(_226_v3);
          })(_137_v3));
          let _228_v66;
          _228_v66 = _dafny.Seq.of(_224_v64, _224_v64, _224_v64);
          let _229_v67;
          _229_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(595),_134_v0);
          let _230_v68;
          _230_v68 = _dafny.MultiSet.fromElements(new BigNumber(((_220_v60).fm8(_134_v0, _139_globalState)).length), _137_v3, (_dafny.ZERO).minus(_137_v3), _137_v3, _137_v3);
          let _index26 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_225_v65).length));
          (_225_v65)[_index26] = (_228_v66)[_module.__default.safeIndex((((((_229_v67).contains(_137_v3)) ? ((_229_v67).get(_137_v3)) : (!(_134_v0)))) ? (_137_v3) : ((_220_v60).fm1(_230_v68, _134_v0, _139_globalState))), new BigNumber((_228_v66).length))];
          let _231_v69;
          let _nw40 = new _module.C0();
          _nw40.__ctor(!(_134_v0), _134_v0, (_220_v60).f20, new BigNumber(224), _134_v0);
          _231_v69 = _nw40;
          let _232_v70;
          _232_v70 = _module.D2.create_DC7(false, false, new BigNumber(-533));
          let _233_v71;
          let _nw41 = new _module.C2();
          _nw41.__ctor(_232_v70, new BigNumber((_dafny.Set.fromElements(_137_v3)).length), (_231_v69).f15);
          _233_v71 = _nw41;
          let _234_v72;
          _234_v72 = _dafny.Map.Empty.slice().updateUnsafe(_231_v69,_233_v71);
          (_139_globalState).f11 = new BigNumber(((_234_v72).Merge(_dafny.Map.Empty.slice().updateUnsafe(_231_v69,_233_v71))).length);
          _229_v67 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("rhwepn")).length),!((_231_v69).f15))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_230_v68).cardinality()),_134_v0));
          (_139_globalState).f11 = (_dafny.ZERO).minus(_137_v3);
          let _index27 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_165_v28).length));
          (_165_v28)[_index27] = (_231_v69).f15;
          let _235_v73;
          _235_v73 = _dafny.Map.Empty.slice().updateUnsafe(_222_v62,_dafny.Seq.Create(_module.__default.abs(new BigNumber(50)), ((_236_v62) => function (_237_i8) {
            return _236_v62;
          })(_222_v62)));
          let _238_v74;
          _238_v74 = _dafny.Map.Empty.slice().updateUnsafe((_231_v69).f15,_134_v0);
          let _239_v75;
          _239_v75 = _dafny.Map.Empty.slice().updateUnsafe((((_220_v60).fm0(_134_v0, _139_globalState)) ? (_238_v74) : (_238_v74)),_134_v0);
          let _index28 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_165_v28).length));
          let _rhs24 = _134_v0;
          let _rhs25 = (_235_v73).update(_222_v62, _135_v1);
          let _rhs26 = ((_239_v75).Merge(_dafny.Map.Empty.slice().updateUnsafe(_238_v74,_134_v0))).Merge(function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of (_dafny.Seq.of(_238_v74, _238_v74, _238_v74)).Elements) {
              let _240_v76 = _compr_9;
              if (_dafny.Seq.contains(_dafny.Seq.of(_238_v74, _238_v74, _238_v74), _240_v76)) {
                _coll9.push([_240_v76,_134_v0]);
              }
            }
            return _coll9;
          }());
          let _rhs27 = _233_v71;
          let _rhs28 = _137_v3;
          let _lhs23 = _165_v28;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_165_v28).length));
          let _lhs25 = _139_globalState;
          _lhs23[_lhs24] = _rhs24;
          _235_v73 = _rhs25;
          _239_v75 = _rhs26;
          _233_v71 = _rhs27;
          _lhs25.f11 = _rhs28;
        } else {
          let _241_v77;
          _241_v77 = new _dafny.CodePoint('y'.codePointAt(0));
          let _242_v78;
          _242_v78 = _dafny.Set.fromElements(_241_v77);
          let _243_v79;
          let _nw42 = new _module.C1();
          _nw42.__ctor(_dafny.Seq.update(_135_v1, _module.__default.safeIndex(new BigNumber((_242_v78).length), new BigNumber((_135_v1).length)), _241_v77), (_220_v60).f21, _137_v3, (_220_v60).fm0(true, _139_globalState));
          _243_v79 = _nw42;
          _243_v79 = _243_v79;
          (_139_globalState).f11 = (_dafny.ZERO).minus(_243_v79.f14);
          let _244_v80;
          let _nw43 = new _module.C0();
          _nw43.__ctor(_134_v0, _134_v0, (_220_v60).f20, _137_v3, _134_v0);
          _244_v80 = _nw43;
          let _245_v81;
          _245_v81 = _dafny.Map.Empty.slice().updateUnsafe(_135_v1,_244_v80);
          let _246_v82;
          let _init8 = ((_247_v80) => function (_248_i9) {
            return _module.__default.safeModuloInt(_248_i9, _247_v80.f14);
          })(_244_v80);
          let _nw44 = Array((new BigNumber(21)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw44.length); _i0_8++) {
            _nw44[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _246_v82 = _nw44;
          let _249_v83;
          _249_v83 = _dafny.Set.fromElements(_246_v82, _246_v82, _246_v82);
          let _250_v84;
          _250_v84 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_137_v3, new BigNumber((_245_v81).length)),new BigNumber((_249_v83).length));
          _250_v84 = (_250_v84).update(_244_v80.f14, _244_v80.f14);
          let _251_v85;
          _251_v85 = _dafny.Map.Empty.slice().updateUnsafe(_137_v3,(_243_v79).f15);
          let _252_v86;
          _252_v86 = _dafny.Map.Empty.slice().updateUnsafe((_251_v85).update(_244_v80.f14, _134_v0),_165_v28);
          _module.__default.m0(_252_v86, _139_globalState);
          (_139_globalState).f11 = (_243_v79.f14).minus(new BigNumber(-435));
        }
        let _253_v87;
        _253_v87 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_137_v3));
        let _254_v88;
        let _nw45 = new _module.C1();
        _nw45.__ctor(_dafny.Seq.Concat((_220_v60).f20, (_220_v60).f20), _module.__default.fm18(_139_globalState), ((_134_v0) ? (new BigNumber((_136_v2).length)) : (new BigNumber((_253_v87).cardinality()))), !(_134_v0));
        _254_v88 = _nw45;
        let _255_v89;
        _255_v89 = _dafny.Seq.of(_254_v88);
        _254_v88 = (_255_v89)[_module.__default.safeIndex((_137_v3).minus(_254_v88.f14), new BigNumber((_255_v89).length))];
        let _256_v90;
        _256_v90 = _module.D2.create_DC7((_254_v88).f15, _134_v0, _137_v3);
        let _nw46 = new _module.C2();
        _nw46.__ctor(_256_v90, _254_v88.f14, (_254_v88).f15);
        _254_v88 = _nw46;
      }
      let _hi0 = new BigNumber(583);
      for (let _257_i10 = _module.__default.safeDivisionInt(_137_v3, _137_v3); _257_i10.isLessThan(_hi0); _257_i10 = _257_i10.plus(_dafny.ONE)) {
        let _258_v91;
        _258_v91 = _dafny.Map.Empty.slice().updateUnsafe(_257_i10,_257_i10);
        let _259_v92;
        _259_v92 = _module.D2.create_DC7(_134_v0, false, _257_i10);
        let _260_v93;
        _260_v93 = _dafny.Set.fromElements(_module.D2.create_DC7(false, _134_v0, (((_258_v91).contains(_137_v3)) ? ((_258_v91).get(_137_v3)) : (_137_v3))), _259_v92);
        let _261_v94;
        _261_v94 = _module.D6.create_DC16(_dafny.Set.fromElements(_260_v93));
        let _262_v95;
        _262_v95 = _dafny.Set.fromElements(_260_v93);
        let _pat_let_tv8 = _262_v95;
        _261_v94 = function (_pat_let3_0) {
          return function (_263_dt__update__tmp_h1) {
            return function (_pat_let4_0) {
              return function (_264_dt__update_hcf29_h0) {
                return _module.D6.create_DC16(_264_dt__update_hcf29_h0);
              }(_pat_let4_0);
            }(_pat_let_tv8);
          }(_pat_let3_0);
        }(_module.D6.create_DC16(_262_v95));
        let _265_v96;
        _265_v96 = new _dafny.CodePoint('q'.codePointAt(0));
        let _266_v97;
        let _nw47 = new _module.C1();
        _nw47.__ctor(_135_v1, _136_v2, _257_i10, _134_v0);
        _266_v97 = _nw47;
        let _267_v98;
        _267_v98 = _dafny.Map.Empty.slice().updateUnsafe(_266_v97,(_266_v97).f15);
        let _268_v99;
        _268_v99 = _dafny.Seq.of(new BigNumber((_267_v98).length));
        let _269_v100;
        let _nw48 = new _module.C0();
        _nw48.__ctor((_266_v97).f15, (_266_v97).f15, _135_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-637)), ((_270_v96) => function (_271_i11) {
          return _270_v96;
        })(_265_v96))).length), _134_v0);
        _269_v100 = _nw48;
        let _272_v101;
        _272_v101 = _dafny.Map.Empty.slice().updateUnsafe(_137_v3,_269_v100);
        let _273_v102;
        _273_v102 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm15(_265_v96, _257_i10, new BigNumber((_268_v99).length), _139_globalState),_module.__default.fm20(new BigNumber(647), new BigNumber((_272_v101).length), _134_v0, _139_globalState));
        let _274_v103;
        _274_v103 = _module.D6.create_DC17(_134_v0, _269_v100);
        let _275_v104;
        _275_v104 = _dafny.Seq.of((_274_v103).dtor_cf31);
        let _276_v106;
        _276_v106 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of _dafny.IntegerRange(new BigNumber(418), new BigNumber(104))) {
            let _277_v105 = _compr_10;
            if (((new BigNumber(418)).isLessThanOrEqualTo(_277_v105)) && ((_277_v105).isLessThan(new BigNumber(104)))) {
              _coll10.push([(_277_v105).minus(_266_v97.f14),_134_v0]);
            }
          }
          return _coll10;
        }(),(_266_v97).f15);
        let _278_v107;
        let _nw49 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _278_v107 = _nw49;
        let _279_v108;
        _279_v108 = _dafny.Seq.of(_278_v107, _278_v107);
        let _280_v109;
        _280_v109 = _dafny.Set.fromElements((_269_v100).f15);
        let _281_v110;
        _281_v110 = _dafny.Map.Empty.slice().updateUnsafe(_165_v28,new BigNumber((_280_v109).length));
        let _282_v111;
        _282_v111 = _dafny.Map.Empty.slice().updateUnsafe(_134_v0,_137_v3);
        let _283_v112;
        _283_v112 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_282_v111).length),(_266_v97).f15);
        let _284_v113;
        let _nw50 = Array((new BigNumber(27)).toNumber());
        _nw50[(_dafny.ZERO).toNumber()] = _257_i10;
        _nw50[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ninu")).length);
        _nw50[(new BigNumber(2)).toNumber()] = new BigNumber((_275_v104).length);
        _nw50[(new BigNumber(3)).toNumber()] = new BigNumber(-917);
        _nw50[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("q")).length);
        _nw50[(new BigNumber(5)).toNumber()] = _137_v3;
        _nw50[(new BigNumber(6)).toNumber()] = new BigNumber((_135_v1).length);
        _nw50[(new BigNumber(7)).toNumber()] = _137_v3;
        _nw50[(new BigNumber(8)).toNumber()] = new BigNumber((_276_v106).length);
        _nw50[(new BigNumber(9)).toNumber()] = new BigNumber((_279_v108).length);
        _nw50[(new BigNumber(10)).toNumber()] = _257_i10;
        _nw50[(new BigNumber(11)).toNumber()] = _257_i10;
        _nw50[(new BigNumber(12)).toNumber()] = new BigNumber((_136_v2).length);
        _nw50[(new BigNumber(13)).toNumber()] = _137_v3;
        _nw50[(new BigNumber(14)).toNumber()] = _257_i10;
        _nw50[(new BigNumber(15)).toNumber()] = _137_v3;
        _nw50[(new BigNumber(16)).toNumber()] = _137_v3;
        _nw50[(new BigNumber(17)).toNumber()] = new BigNumber((_135_v1).length);
        _nw50[(new BigNumber(18)).toNumber()] = (_268_v99)[_module.__default.safeIndex(new BigNumber((_136_v2).length), new BigNumber((_268_v99).length))];
        _nw50[(new BigNumber(19)).toNumber()] = new BigNumber(413);
        _nw50[(new BigNumber(20)).toNumber()] = _257_i10;
        _nw50[(new BigNumber(21)).toNumber()] = new BigNumber((_281_v110).length);
        _nw50[(new BigNumber(22)).toNumber()] = _269_v100.f14;
        _nw50[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_283_v112).length));
        _nw50[(new BigNumber(24)).toNumber()] = _269_v100.f14;
        _nw50[(new BigNumber(25)).toNumber()] = _257_i10;
        _nw50[(new BigNumber(26)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_134_v0,(_269_v100).f15)).length);
        _284_v113 = _nw50;
        let _285_v114;
        _285_v114 = _dafny.MultiSet.fromElements(_284_v113, _284_v113, _284_v113);
        let _286_v115;
        _286_v115 = _dafny.Seq.of(_module.__default.fm19(_139_globalState), new BigNumber((_273_v102).length), new BigNumber((_285_v114).cardinality()), _137_v3);
        _286_v115 = _268_v99;
        let _287_v116;
        _287_v116 = _dafny.Map.Empty.slice().updateUnsafe(_265_v96,_module.D1.create_DC3(_135_v1));
        let _288_v117;
        _288_v117 = _dafny.Map.Empty.slice().updateUnsafe(_284_v113,_287_v116);
        _137_v3 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_266_v97.f14).plus(_266_v97.f14),(((_288_v117).contains(_284_v113)) ? ((_288_v117).get(_284_v113)) : (_287_v116)))).length);
        (_266_v97).f14 = _module.__default.fm19(_139_globalState);
      }
      _137_v3 = _137_v3;
      let _289_v118;
      _289_v118 = _module.D3.create_DC10(_135_v1, _165_v28, _134_v0, _137_v3, new BigNumber(-289));
      let _290_v119;
      let _nw51 = Array((new BigNumber(21)).toNumber());
      _nw51[(_dafny.ZERO).toNumber()] = _135_v1;
      _nw51[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat((_289_v118).dtor_cf17, _135_v1);
      _nw51[(new BigNumber(2)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(3)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(4)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(5)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(6)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(7)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(8)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(9)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(10)).toNumber()] = _module.__default.fm9(_137_v3, _139_globalState);
      _nw51[(new BigNumber(11)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("qwdbsyk");
      _nw51[(new BigNumber(13)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_135_v1, _dafny.Seq.UnicodeFromString("iiayqcdiq"));
      _nw51[(new BigNumber(15)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(16)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("hbwlhhvq");
      _nw51[(new BigNumber(18)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(19)).toNumber()] = _135_v1;
      _nw51[(new BigNumber(20)).toNumber()] = _135_v1;
      _290_v119 = _nw51;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_290_v119).length))) {
        let _291_i12 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_291_i12)) && ((_291_i12).isLessThan(new BigNumber((_290_v119).length))))) {
          (_290_v119)[(_291_i12)] = (((false) === (_134_v0)) ? (_135_v1) : (_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(951)), function (_292_i13) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          }), _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(false, _134_v0)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(951)), function (_293_i13) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          })).length)), new _dafny.CodePoint('g'.codePointAt(0)))));
        }
      }
      let _294_v120;
      _294_v120 = _dafny.Map.Empty.slice().updateUnsafe(_137_v3,new BigNumber(622));
      let _295_i14;
      _295_i14 = _dafny.ZERO;
      L0: {
        while ((((_134_v0) || (_134_v0)) ? (_134_v0) : ((_134_v0) === (_module.__default.fm21(_137_v3, !(false), !(_134_v0), _294_v120, _139_globalState))))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_295_i14)) {
              break L0;
            }
            _295_i14 = (_295_i14).plus(_dafny.ONE);
            let _296_v121;
            _296_v121 = _dafny.Seq.of(_module.__default.safeModuloInt(new BigNumber((_135_v1).length), new BigNumber((_138_v4).length)));
            _296_v121 = _dafny.Seq.update(_dafny.Seq.update(_296_v121, _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm19(_139_globalState))), new BigNumber((_296_v121).length)), _module.__default.safeDivisionInt(new BigNumber(864), _137_v3)), _module.__default.safeIndex(_137_v3, new BigNumber((_dafny.Seq.update(_296_v121, _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm19(_139_globalState))), new BigNumber((_296_v121).length)), _module.__default.safeDivisionInt(new BigNumber(864), _137_v3))).length)), new BigNumber(366));
            _135_v1 = _dafny.Seq.UnicodeFromString("cuplutb");
            let _297_v122;
            _297_v122 = new _dafny.CodePoint('u'.codePointAt(0));
            if ((!((new BigNumber((_dafny.Seq.update(_135_v1, _module.__default.safeIndex(_137_v3, new BigNumber((_135_v1).length)), _297_v122)).length)).isLessThan(_137_v3))) || (!(_137_v3).isEqualTo(new BigNumber(911)))) {
              let _298_v123;
              let _nw52 = new _module.C0();
              _nw52.__ctor((_137_v3).isEqualTo(new BigNumber(79)), (_134_v0) || (_134_v0), _135_v1, _137_v3, ((_134_v0) ? (_134_v0) : (false)));
              _298_v123 = _nw52;
              let _299_v124;
              _299_v124 = _dafny.Map.Empty.slice().updateUnsafe(_137_v3,(_137_v3).isLessThan(_137_v3));
              _299_v124 = _299_v124;
              let _300_v125;
              let _nw53 = new _module.C1();
              _nw53.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-172)), ((_301_v122) => function (_302_i15) {
                return _301_v122;
              })(_297_v122)), _138_v4, new BigNumber((_135_v1).length), (_298_v123).f17);
              _300_v125 = _nw53;
              _300_v125 = _300_v125;
              let _303_v126;
              _303_v126 = _dafny.Map.Empty.slice().updateUnsafe(_299_v124,_165_v28);
              let _304_v127;
              _304_v127 = _dafny.Seq.of(_303_v126);
              _module.__default.m0(((_304_v127)[_module.__default.safeIndex(_137_v3, new BigNumber((_304_v127).length))]).Merge(_303_v126), _139_globalState);
              let _305_v128;
              _305_v128 = _dafny.MultiSet.fromElements(new BigNumber((_135_v1).length), _137_v3);
              (_139_globalState).f3 = (((_299_v124).contains(new BigNumber((_module.__default.fm13(new BigNumber((_305_v128).cardinality()), _139_globalState)).length))) ? ((_299_v124).get(new BigNumber((_module.__default.fm13(new BigNumber((_305_v128).cardinality()), _139_globalState)).length))) : ((_298_v123).f17));
            } else {
              let _306_v129;
              let _nw54 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
              _306_v129 = _nw54;
              let _index29 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_306_v129).length));
              (_306_v129)[_index29] = _137_v3;
              let _index30 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_306_v129).length));
              (_306_v129)[_index30] = _137_v3;
              _137_v3 = (_306_v129)[_module.__default.safeIndex(new BigNumber(293), new BigNumber((_306_v129).length))];
              let _307_v130;
              _307_v130 = _module.D3.create_DC8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(309)), function (_308_i16) {
  return _module.D1.create_DC3(_dafny.Seq.UnicodeFromString("psxioq"));
}));
              let _309_v131;
              _309_v131 = _dafny.MultiSet.fromElements(_134_v0);
              _307_v130 = (((_309_v131).IsSubsetOf(_309_v131)) ? (_307_v130) : (_307_v130));
              (_139_globalState).f0 = (!(_134_v0)) || (_134_v0);
              let _310_v132;
              _310_v132 = _dafny.Map.Empty.slice().updateUnsafe((_306_v129)[_module.__default.safeIndex(new BigNumber(293), new BigNumber((_306_v129).length))],_134_v0);
              let _311_v133;
              _311_v133 = _dafny.Map.Empty.slice().updateUnsafe(_310_v132,_165_v28);
              _module.__default.m0(_311_v133, _139_globalState);
            }
            let _312_v134;
            _312_v134 = _dafny.Map.Empty.slice().updateUnsafe(_134_v0,_137_v3);
            let _313_v135;
            _313_v135 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_135_v1).length),(_312_v134).Merge(_dafny.Map.Empty.slice().updateUnsafe(_134_v0,new BigNumber((_294_v120).length))));
            _313_v135 = (_313_v135).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(935),_312_v134));
          }
        }
      }
      let _314_v136;
      _314_v136 = _dafny.MultiSet.fromElements(_137_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(168)), function (_315_i17) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })).length), _137_v3);
      let _316_v137;
      _316_v137 = _dafny.Seq.of(_137_v3, _137_v3);
      let _317_v138;
      let _nw55 = new _module.C0();
      _nw55.__ctor((_137_v3).isLessThan(_137_v3), (_314_v136).IsProperSubsetOf(_314_v136), _135_v1, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_137_v3), _316_v137)).length)), !(_module.__default.fm21(_137_v3, _134_v0, _134_v0, _294_v120, _139_globalState)));
      _317_v138 = _nw55;
      let _318_v139;
      let _nw56 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
      _318_v139 = _nw56;
      _318_v139 = _318_v139;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_318_v139).length))) {
        let _319_i18 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_319_i18)) && ((_319_i18).isLessThan(new BigNumber((_318_v139).length))))) {
          (_318_v139)[(_319_i18)] = _module.__default.safeModuloInt(_319_i18, _137_v3);
        }
      }
      let _320_v140;
      _320_v140 = _dafny.MultiSet.fromElements(true);
      (_139_globalState).f0 = (_320_v140).IsDisjointFrom(_320_v140);
      if (!((_317_v138).f18)) {
        let _index31 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_165_v28).length));
        (_165_v28)[_index31] = (_317_v138).f17;
        let _index32 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_165_v28).length));
        (_165_v28)[_index32] = (_317_v138).f18;
        let _321_v141;
        let _init9 = ((_322_v3) => function (_323_i19) {
          return _dafny.Set.fromElements(_322_v3);
        })(_137_v3);
        let _nw57 = Array((new BigNumber(23)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw57.length); _i0_9++) {
          _nw57[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _321_v141 = _nw57;
        let _324_v142;
        _324_v142 = _dafny.Set.fromElements((((_314_v136).contains((((_320_v140).contains(_134_v0)) ? ((_320_v140).get(_134_v0)) : (_137_v3)))) ? ((_314_v136).get((((_320_v140).contains(_134_v0)) ? ((_320_v140).get(_134_v0)) : (_137_v3)))) : (_137_v3)));
        let _index33 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_321_v141).length));
        (_321_v141)[_index33] = _324_v142;
        let _index34 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_318_v139).length));
        (_318_v139)[_index34] = _137_v3;
        let _index35 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_321_v141).length));
        let _index36 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_318_v139).length));
        let _rhs29 = _324_v142;
        let _rhs30 = (_dafny.ZERO).minus(_137_v3);
        let _rhs31 = _137_v3;
        let _rhs32 = ((((_314_v136).contains(_137_v3)) ? ((_314_v136).get(_137_v3)) : ((_317_v138).fm2(_dafny.Seq.Create(_module.__default.abs(new BigNumber(241)), function (_325_i20) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }), (_317_v138).f18, (_165_v28)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_165_v28).length))], _137_v3, _139_globalState)))).multipliedBy(_137_v3);
        let _lhs26 = _321_v141;
        let _lhs27 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_321_v141).length));
        let _lhs28 = _139_globalState;
        let _lhs29 = _139_globalState;
        let _lhs30 = _318_v139;
        let _lhs31 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_318_v139).length));
        _lhs26[_lhs27] = _rhs29;
        _lhs28.f11 = _rhs30;
        _lhs29.f11 = _rhs31;
        _lhs30[_lhs31] = _rhs32;
        let _326_v143;
        _326_v143 = _module.D2.create_DC7((_317_v138).f18, _134_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_316_v137).length),(_165_v28)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_165_v28).length))])).length));
        let _327_v144;
        let _nw58 = new _module.C2();
        _nw58.__ctor(_326_v143, new BigNumber((_320_v140).cardinality()), true);
        _327_v144 = _nw58;
        let _328_v145;
        _328_v145 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm21((_318_v139)[_module.__default.safeIndex(new BigNumber(575), new BigNumber((_318_v139).length))], (_317_v138).f18, _134_v0, _294_v120, _139_globalState),_327_v144);
        let _329_v146;
        let _nw59 = Array((new BigNumber(24)).toNumber());
        _nw59[(_dafny.ZERO).toNumber()] = _327_v144;
        _nw59[(_dafny.ONE).toNumber()] = (((_317_v138).f18) ? (_327_v144) : (_327_v144));
        _nw59[(new BigNumber(2)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(3)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(4)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(5)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(6)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(7)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(8)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(9)).toNumber()] = (((_328_v145).contains(_134_v0)) ? ((_328_v145).get(_134_v0)) : (_327_v144));
        _nw59[(new BigNumber(10)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(11)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(12)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(13)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(14)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(15)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(16)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(17)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(18)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(19)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(20)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(21)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(22)).toNumber()] = _327_v144;
        _nw59[(new BigNumber(23)).toNumber()] = _327_v144;
        _329_v146 = _nw59;
        let _index37 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_329_v146).length));
        (_329_v146)[_index37] = _327_v144;
        let _index38 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_329_v146).length));
        (_329_v146)[_index38] = _327_v144;
        let _330_v147;
        let _nw60 = new _module.C0();
        _nw60.__ctor((_317_v138).f17, _134_v0, _135_v1, (new BigNumber((_dafny.Seq.UnicodeFromString("syla")).length)).minus(_137_v3), _134_v0);
        _330_v147 = _nw60;
        let _331_v148;
        _331_v148 = _dafny.Map.Empty.slice().updateUnsafe(_135_v1,(_165_v28)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_165_v28).length))]);
        let _332_v149;
        _332_v149 = _dafny.Map.Empty.slice().updateUnsafe(_331_v148,(_317_v138).f17);
        let _index39 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_165_v28).length));
        let _rhs33 = (_317_v138).f17;
        let _rhs34 = (_330_v147).f18;
        let _rhs35 = (_165_v28)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_165_v28).length))];
        let _rhs36 = (((_332_v149).contains(_331_v148)) ? ((_332_v149).get(_331_v148)) : ((_330_v147).f17));
        let _rhs37 = ((false) ? (_module.__default.fm21(_137_v3, !((_317_v138).f17), _134_v0, _294_v120, _139_globalState)) : (!((_317_v138).f17)));
        let _lhs32 = _165_v28;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_165_v28).length));
        let _lhs34 = _139_globalState;
        let _lhs35 = _139_globalState;
        _lhs32[_lhs33] = _rhs33;
        _134_v0 = _rhs34;
        _lhs34.f3 = _rhs35;
        _134_v0 = _rhs36;
        _lhs35.f0 = _rhs37;
      } else {
        (_139_globalState).f3 = (_317_v138).f18;
        if ((_317_v138).f18) {
          let _index40 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_165_v28).length));
          (_165_v28)[_index40] = (_317_v138).f17;
          let _index41 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_165_v28).length));
          (_165_v28)[_index41] = (_317_v138).f18;
          let _333_v150;
          _333_v150 = _dafny.Map.Empty.slice().updateUnsafe(_137_v3,(_317_v138).f17);
          let _334_v151;
          let _nw61 = Array((new BigNumber(11)).toNumber()).fill(false);
          _334_v151 = _nw61;
          let _335_v152;
          _335_v152 = _dafny.Map.Empty.slice().updateUnsafe(_333_v150,_334_v151);
          _module.__default.m0((_335_v152).Merge(_335_v152), _139_globalState);
          let _336_v153;
          _336_v153 = _module.D2.create_DC7((_317_v138).f17, (_317_v138).f18, _137_v3);
          let _337_v154;
          let _nw62 = new _module.C1();
          _nw62.__ctor(_135_v1, _138_v4, _137_v3, (_317_v138).f18);
          _337_v154 = _nw62;
          let _338_v155;
          _338_v155 = _dafny.Map.Empty.slice().updateUnsafe(_337_v154,!((_317_v138).f17));
          let _339_v156;
          let _nw63 = new _module.C2();
          _nw63.__ctor(_336_v153, new BigNumber((_338_v155).length), (_137_v3).isLessThanOrEqualTo(new BigNumber((_136_v2).length)));
          _339_v156 = _nw63;
          (_139_globalState).f0 = (((_317_v138).f17) ? (!(_314_v136).contains(_137_v3)) : ((_317_v138).f18));
          let _340_v157;
          _340_v157 = _dafny.Map.Empty.slice().updateUnsafe(_137_v3,_320_v140);
          let _341_v158;
          _341_v158 = _module.D2.create_DC6(_340_v157);
          let _342_v159;
          let _343_v160;
          let _out0;
          let _out1;
          let _outcollector0 = (_339_v156).m2(_341_v158, (_317_v138).f17, (_dafny.ZERO).minus((_dafny.ZERO).minus(_137_v3)), _137_v3, _139_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _342_v159 = _out0;
          _343_v160 = _out1;
        } else {
          let _344_v161;
          let _init10 = function (_345_i21) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          };
          let _nw64 = Array((new BigNumber(24)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw64.length); _i0_10++) {
            _nw64[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _344_v161 = _nw64;
          let _346_v162;
          _346_v162 = new _dafny.CodePoint('g'.codePointAt(0));
          let _index42 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_344_v161).length));
          (_344_v161)[_index42] = _346_v162;
          let _index43 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_344_v161).length));
          (_344_v161)[_index43] = _346_v162;
          let _347_v163;
          _347_v163 = _module.D1.create_DC3(_dafny.Seq.UnicodeFromString("tcir"));
          _347_v163 = _module.D1.create_DC3(_135_v1);
          (_139_globalState).f11 = _137_v3;
          (_139_globalState).f11 = (_317_v138).fm1(_314_v136, _134_v0, _139_globalState);
          _165_v28 = _165_v28;
        }
        _316_v137 = _dafny.Seq.Concat(_316_v137, _316_v137);
        (_139_globalState).f11 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_137_v3, _137_v3), _137_v3);
        let _348_v164;
        _348_v164 = _dafny.Set.fromElements((_317_v138).f17);
        let _rhs38 = _137_v3;
        let _rhs39 = (_137_v3).minus(_137_v3);
        let _rhs40 = _348_v164;
        _137_v3 = _rhs38;
        _137_v3 = _rhs39;
        _348_v164 = _rhs40;
      }
      let _index44 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_165_v28).length));
      (_165_v28)[_index44] = (_317_v138).fm0((_317_v138).f17, _139_globalState);
      let _349_v165;
      _349_v165 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(255),(_317_v138).f18);
      let _350_v166;
      _350_v166 = _module.D3.create_DC9(new BigNumber((_349_v165).length));
      let _pat_let_tv9 = _317_v138;
      let _pat_let_tv10 = _317_v138;
      let _pat_let_tv11 = _317_v138;
      let _pat_let_tv12 = _317_v138;
      let _index45 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_165_v28).length));
      (_165_v28)[_index45] = function (_source5) {
        if (_source5.is_DC9) {
          let _351___mcc_h7 = (_source5).cf16;
          let _352_cf16 = _351___mcc_h7;
          return !((_pat_let_tv9).f18);
        } else if (_source5.is_DC10) {
          let _353___mcc_h8 = (_source5).cf17;
          let _354___mcc_h9 = (_source5).cf18;
          let _355___mcc_h10 = (_source5).cf19;
          let _356___mcc_h11 = (_source5).cf20;
          let _357___mcc_h12 = (_source5).cf21;
          let _358_cf21 = _357___mcc_h12;
          let _359_cf20 = _356___mcc_h11;
          let _360_cf19 = _355___mcc_h10;
          let _361_cf18 = _354___mcc_h9;
          let _362_cf17 = _353___mcc_h8;
          return (_pat_let_tv10).f18;
        } else if (_source5.is_DC8) {
          let _363___mcc_h13 = (_source5).cf15;
          let _364_cf15 = _363___mcc_h13;
          return (_pat_let_tv11).f18;
        } else {
          let _365___mcc_h14 = (_source5).cf22;
          let _366_cf22 = _365___mcc_h14;
          return (_pat_let_tv12).f18;
        }
      }(_350_v166);
      let _367_v167;
      let _nw65 = Array((new BigNumber(26)).toNumber());
      _367_v167 = _nw65;
      _367_v167 = _367_v167;
      process.stdout.write(_dafny.toString(_134_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_135_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_136_v2, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_138_v4, _dafny.Seq.of(false, false, false, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_139_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_139_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_139_globalState).f4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_139_globalState.f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_139_globalState).f13, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v28)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v28)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v28)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v28)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v28)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v28)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v28)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_169_v29, _dafny.Seq.of(_dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false), _dafny.Seq.of(false, false, false, false, true), _dafny.Seq.of(false), _dafny.Seq.of(false, false, false, false, true, false, false, false, false, true), _dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_289_v118).dtor_cf17).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_289_v118).dtor_cf18)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_289_v118).dtor_cf18)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_289_v118).dtor_cf18)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_289_v118).dtor_cf18)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_289_v118).dtor_cf18)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_289_v118).dtor_cf18)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_289_v118).dtor_cf18)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_v118).dtor_cf19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_v118).dtor_cf20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_v118).dtor_cf21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(15)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(16)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(17)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(18)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(19)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_290_v119)[new BigNumber(20)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_294_v120).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(622)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_295_i14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_314_v136).equals(_dafny.MultiSet.fromElements(_dafny.ONE, _dafny.ONE, new BigNumber(168)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_316_v137, _dafny.Seq.of(_dafny.ONE, _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_317_v138).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_317_v138).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v139)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_320_v140).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_349_v165).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(255),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_350_v166).dtor_cf16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC2(cf2, cf3) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf5, cf6) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC5(cf7, cf8, cf9, cf10) {
      let $dt = new D1(1);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC3(cf4) {
      let $dt = new D1(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + this.cf4.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf12, cf13, cf14) {
      let $dt = new D2(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC6(cf11) {
      let $dt = new D2(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf12 === other.cf12 && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf16) {
      let $dt = new D3(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC10(cf17, cf18, cf19, cf20, cf21) {
      let $dt = new D3(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC8(cf15) {
      let $dt = new D3(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC11(cf22) {
      let $dt = new D3(3);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC11() { return this.$tag === 3; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + this.cf17.toVerbatimString(true) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17) && this.cf18 === other.cf18 && this.cf19 === other.cf19 && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12(cf23) {
      let $dt = new D4(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf25, cf26, cf27) {
      let $dt = new D5(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC13(cf24) {
      let $dt = new D5(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf28) {
      let $dt = new D5(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25) && this.cf26 === other.cf26 && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(_dafny.ZERO, [], false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf30, cf31) {
      let $dt = new D6(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC18(cf32, cf33) {
      let $dt = new D6(1);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC16(cf29) {
      let $dt = new D6(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC19(cf34) {
      let $dt = new D6(3);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get is_DC19() { return this.$tag === 3; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30 && this.cf31 === other.cf31;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(false, null);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf35) {
      let $dt = new D7(0);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf35 === other.cf35;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf37) {
      let $dt = new D8(0);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC23(cf38, cf39) {
      let $dt = new D8(1);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC21(cf36) {
      let $dt = new D8(2);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC24(cf40) {
      let $dt = new D8(3);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf38 === other.cf38 && this.cf39 === other.cf39;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC22(new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26(cf42, cf43) {
      let $dt = new D9(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC27() {
      let $dt = new D9(1);
      return $dt;
    }
    static create_DC25(cf41) {
      let $dt = new D9(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC27";
      } else if (this.$tag === 2) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42) && this.cf43 === other.cf43;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf41 === other.cf41;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC26(_dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC29(cf45) {
      let $dt = new D10(0);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC28(cf44) {
      let $dt = new D10(1);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf45 === other.cf45;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC29(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = false;
      this.f3 = false;
      this.f11 = _dafny.ZERO;
      this._f1 = false;
      this._f2 = _dafny.ZERO;
      this._f4 = _dafny.Seq.UnicodeFromString("");
      this._f5 = false;
      this._f6 = _dafny.ZERO;
      this._f7 = false;
      this._f8 = _dafny.ZERO;
      this._f9 = _dafny.ZERO;
      this._f10 = _dafny.ZERO;
      this._f12 = false;
      this._f13 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f14 = _dafny.ZERO;
      this._f16 = _dafny.Seq.UnicodeFromString("");
      this._f15 = false;
      this._f17 = false;
      this._f18 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    __ctor(f17, f18, f16, f14, f15) {
      let _this = this;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f16 = f16;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      return !(!((_module.D0.create_DC0((_this).f18)).dtor_cf0));
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return _this.f14;
    };
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.update((_module.D1.create_DC3(_this.f16)).dtor_cf4, _module.__default.safeIndex((_module.D1.create_DC4(new BigNumber((function () {
  let _coll11 = new _dafny.Map();
  for (const _compr_11 of _dafny.IntegerRange(new BigNumber(404), new BigNumber(615))) {
    let _368_v0 = _compr_11;
    if (((new BigNumber(404)).isLessThanOrEqualTo(_368_v0)) && ((_368_v0).isLessThan(new BigNumber(615)))) {
      _coll11.push([(_368_v0).plus(_this.f14),_this.f14]);
    }
  }
  return _coll11;
}()).length), _this.f14)).dtor_cf5, new BigNumber(((_module.D1.create_DC3(_this.f16)).dtor_cf4).length)), (_this.f16)[_module.__default.safeIndex(_this.f14, new BigNumber((_this.f16).length))])).length);
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements(_this.f14, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(523)), function (_369_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f18, (_this).f18, (_this).f15), _dafny.Seq.of((_this).f15))).length)));
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f14 = _dafny.ZERO;
      this._f15 = false;
      this._f20 = _dafny.Seq.UnicodeFromString("");
      this._f21 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    __ctor(f20, f21, f14, f15) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      return (_this).f15;
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("x"),new BigNumber((_dafny.Seq.of(new BigNumber(244), _this.f14)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,new BigNumber(((_this).f20).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f20,_this.f14)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,_this.f14)))).length));
    };
    fm8(p0, globalState) {
      let _this = this;
      let _source6 = _module.D0.create_DC2(_this.f14, _dafny.MultiSet.fromElements((_this).f15));
      if (_source6.is_DC1) {
        let _370___mcc_h0 = (_source6).cf1;
        let _371_cf1 = _370___mcc_h0;
        return _dafny.Seq.UnicodeFromString("uuycpn");
      } else if (_source6.is_DC2) {
        let _372___mcc_h1 = (_source6).cf2;
        let _373___mcc_h2 = (_source6).cf3;
        let _374_cf3 = _373___mcc_h2;
        let _375_cf2 = _372___mcc_h1;
        return (_this).f20;
      } else {
        let _376___mcc_h3 = (_source6).cf0;
        let _377_cf0 = _376___mcc_h3;
        return (_this).f20;
      }
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f14 = _dafny.ZERO;
      this._f15 = false;
      this._f19 = _module.D2.Default();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    set f14(value) {
      let _this = this;
      _this._f14 = value;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    __ctor(f19, f14, f15) {
      let _this = this;
      (_this)._f19 = f19;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      return (_this).f15;
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(292);
    };
    fm4(p0, globalState) {
      let _this = this;
      return new _dafny.CodePoint('m'.codePointAt(0));
    };
    fm5(p0, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements((_this).f15)).Union((_dafny.Set.fromElements((_this).f15)).Difference(_dafny.Set.fromElements((_this).f15)));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _module.D1.Default();
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _378_v0;
      let _nw66 = Array((new BigNumber(22)).toNumber()).fill(false);
      _378_v0 = _nw66;
      let _index46 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length));
      (_378_v0)[_index46] = (_this).f15;
      let _379_v1;
      let _init11 = function (_380_i0) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eafbti"), _dafny.Seq.UnicodeFromString("i"));
      };
      let _nw67 = Array((new BigNumber(5)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw67.length); _i0_11++) {
        _nw67[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _379_v1 = _nw67;
      let _381_v2;
      _381_v2 = _dafny.Seq.UnicodeFromString("kaueyjj");
      let _index47 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length));
      (_379_v1)[_index47] = _381_v2;
      let _382_v3;
      _382_v3 = new _dafny.CodePoint('q'.codePointAt(0));
      let _index48 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length));
      let _index49 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length));
      let _rhs41 = (_this).f15;
      let _rhs42 = _dafny.Seq.update(_dafny.Seq.Concat(_381_v2, _381_v2), _module.__default.safeIndex((_dafny.ZERO).minus(_this.f14), new BigNumber((_dafny.Seq.Concat(_381_v2, _381_v2)).length)), _382_v3);
      let _lhs36 = _378_v0;
      let _lhs37 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length));
      let _lhs38 = _379_v1;
      let _lhs39 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length));
      _lhs36[_lhs37] = _rhs41;
      _lhs38[_lhs39] = _rhs42;
      if ((_dafny.MultiSet.fromElements(_this.f14)).contains(_this.f14)) {
        (globalState).f11 = _this.f14;
        (globalState).f3 = ((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]) === ((_this).f15);
        let _383_v4;
        _383_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this.f14).minus(_this.f14));
        let _384_v5;
        _384_v5 = _dafny.MultiSet.fromElements((_this).f15);
        _383_v4 = (_383_v4).update((new BigNumber((_384_v5).cardinality())).isLessThan(new BigNumber((_381_v2).length)), _this.f14);
        let _385_v6;
        _385_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,!((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]) || ((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]));
        r2 = (((_385_v6).contains((_383_v4).equals(_383_v4))) ? ((_385_v6).get((_383_v4).equals(_383_v4))) : ((_this).fm0((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], globalState)));
        let _386_v7;
        _386_v7 = _module.D0.create_DC0(true);
        _386_v7 = _module.__default.fm6(false, globalState);
      } else {
        let _index50 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length));
        (_379_v1)[_index50] = _381_v2;
        let _387_v8;
        _387_v8 = _dafny.Map.Empty.slice().updateUnsafe(_382_v3,(_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))]);
        _387_v8 = (_387_v8).update(new _dafny.CodePoint('r'.codePointAt(0)), (_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))]);
        let _388_v9;
        _388_v9 = _dafny.Seq.of((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], (_this).f15);
        let _389_v10;
        let _nw68 = new _module.C1();
        _nw68.__ctor(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(823)), function (_390_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), _module.__default.safeIndex(_this.f14, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(823)), function (_391_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        })).length)), _382_v3), _dafny.Seq.update(_388_v9, _module.__default.safeIndex(_this.f14, new BigNumber((_388_v9).length)), (_this).f15), _this.f14, (_this).f15);
        _389_v10 = _nw68;
        let _392_v11;
        _392_v11 = _dafny.Seq.of(_389_v10, _389_v10);
        let _393_v12;
        _393_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(376),(_this.f14).multipliedBy(new BigNumber(((_module.__default.fm7((_this).f15, (_this).f15, new BigNumber((_392_v11).length), globalState)).dtor_cf15).length)));
        _393_v12 = (_393_v12).update(_this.f14, _389_v10.f14);
        (globalState).f0 = (_this).f15;
        (globalState).f11 = new BigNumber(727);
      }
      let _hi1 = (_this.f14).multipliedBy((_dafny.ZERO).minus(new BigNumber(-936)));
      for (let _394_i2 = _this.f14; _394_i2.isLessThan(_hi1); _394_i2 = _394_i2.plus(_dafny.ONE)) {
        let _395_v13;
        _395_v13 = _dafny.Set.fromElements(_382_v3);
        let _396_v15;
        _396_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(-98)),function () {
          let _coll12 = new _dafny.Set();
          for (const _compr_12 of ((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))]).Elements) {
            let _397_v14 = _compr_12;
            if (_dafny.Seq.contains((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))], _397_v14)) {
              _coll12.add(_397_v14);
            }
          }
          return _coll12;
        }());
        if ((_395_v13).equals((((_396_v15).contains(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_382_v3,_394_i2)).length)))) ? ((_396_v15).get(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_382_v3,_394_i2)).length)))) : (_395_v13)))) {
          let _398_v16;
          let _nw69 = new _module.C0();
          _nw69.__ctor(false, (_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], (_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))], new BigNumber((_381_v2).length), !((_this).f15));
          _398_v16 = _nw69;
          let _399_v17;
          let _init12 = function (_400_i3) {
            return _module.D3.create_DC9(_this.f14);
          };
          let _nw70 = Array((new BigNumber(26)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw70.length); _i0_12++) {
            _nw70[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _399_v17 = _nw70;
          let _401_v18;
          let _init13 = ((_402_v16) => function (_403_i4) {
            return _module.__default.safeModuloInt(_403_i4, new BigNumber((_dafny.MultiSet.fromElements((_402_v16).f17)).cardinality()));
          })(_398_v16);
          let _nw71 = Array((new BigNumber(5)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw71.length); _i0_13++) {
            _nw71[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _401_v18 = _nw71;
          let _index51 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_401_v18).length));
          (_401_v18)[_index51] = _this.f14;
          let _404_v19;
          _404_v19 = _dafny.Seq.of(_399_v17);
          let _405_v20;
          _405_v20 = _dafny.Seq.of((((_398_v16).f18) ? (_399_v17) : ((_404_v19)[_module.__default.safeIndex(_this.f14, new BigNumber((_404_v19).length))])));
          let _index52 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_401_v18).length));
          let _rhs43 = (_this).f15;
          let _rhs44 = _398_v16;
          let _rhs45 = (_405_v20)[_module.__default.safeIndex(_394_i2, new BigNumber((_405_v20).length))];
          let _rhs46 = new BigNumber(474);
          let _lhs40 = _401_v18;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_401_v18).length));
          r2 = _rhs43;
          _398_v16 = _rhs44;
          _399_v17 = _rhs45;
          _lhs40[_lhs41] = _rhs46;
          let _406_v21;
          _406_v21 = _dafny.Seq.of((_398_v16).f18);
          let _407_v22;
          _407_v22 = _dafny.MultiSet.fromElements(_406_v21);
          let _408_v23;
          _408_v23 = _dafny.Map.Empty.slice().updateUnsafe((_398_v16).f18,(_398_v16).f17);
          let _409_v24;
          _409_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(360)), ((_410_v3) => function (_411_i5) {
            return _410_v3;
          })(_382_v3))).length),_this.f14);
          let _index53 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length));
          let _rhs47 = (_398_v16).f17;
          let _rhs48 = (((((_408_v23).contains(true)) ? ((_408_v23).get(true)) : ((_398_v16).f17))) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f14)))) : ((new BigNumber((_module.__default.fm9(_this.f14, globalState)).length)).plus(_394_i2)));
          let _rhs49 = _407_v22;
          let _rhs50 = (_406_v21)[_module.__default.safeIndex((((_409_v24).contains(new BigNumber(769))) ? ((_409_v24).get(new BigNumber(769))) : (_this.f14)), new BigNumber((_406_v21).length))];
          let _lhs42 = globalState;
          let _lhs43 = globalState;
          let _lhs44 = _378_v0;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length));
          _lhs42.f0 = _rhs47;
          _lhs43.f11 = _rhs48;
          _407_v22 = _rhs49;
          _lhs44[_lhs45] = _rhs50;
          let _412_v25;
          _412_v25 = _379_v1;
          _379_v1 = (_412_v25);
          let _413_v26;
          _413_v26 = _dafny.Map.Empty.slice().updateUnsafe(!(!((_398_v16).f18)),_378_v0);
          _378_v0 = (((_413_v26).contains((_this).f15)) ? ((_413_v26).get((_this).f15)) : (_378_v0));
          let _414_v27;
          let _nw72 = new _module.C0();
          _nw72.__ctor((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], (_398_v16).f17, (_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))], _this.f14, (_398_v16).f18);
          _414_v27 = _nw72;
        } else {
          (globalState).f11 = _394_i2;
          let _415_v28;
          _415_v28 = _dafny.Seq.of(_this.f14);
          let _416_v29;
          _416_v29 = _dafny.Seq.of((_this).f15);
          let _417_v30;
          let _nw73 = new _module.C1();
          _nw73.__ctor(_381_v2, _dafny.Seq.of((_this).fm0((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], globalState), (_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]), (_394_i2).plus((_415_v28)[_module.__default.safeIndex(_394_i2, new BigNumber((_415_v28).length))]), (_416_v29)[_module.__default.safeIndex(_this.f14, new BigNumber((_416_v29).length))]);
          _417_v30 = _nw73;
          (globalState).f11 = (_dafny.ZERO).minus(_394_i2);
          (globalState).f3 = (_this).f15;
          let _418_v31;
          let _init14 = ((_419_v0, _420_v28) => function (_421_i6) {
            return (_dafny.Map.Empty.slice().updateUnsafe((_419_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_419_v0).length))],_420_v28)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,_420_v28));
          })(_378_v0, _415_v28);
          let _nw74 = Array((new BigNumber(11)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw74.length); _i0_14++) {
            _nw74[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _418_v31 = _nw74;
          let _index54 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_418_v31).length));
          (_418_v31)[_index54] = _module.__default.fm10((_this).f15, globalState);
          let _422_v32;
          _422_v32 = _dafny.MultiSet.fromElements(_394_i2);
          let _423_v33;
          _423_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_dafny.Seq.update(_dafny.Seq.update(_415_v28, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_415_v28)).cardinality()), new BigNumber((_415_v28).length)), _394_i2), _module.__default.safeIndex((_417_v30).fm1(_422_v32, (_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], globalState), new BigNumber((_dafny.Seq.update(_415_v28, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_415_v28)).cardinality()), new BigNumber((_415_v28).length)), _394_i2)).length)), new BigNumber((_381_v2).length)));
          let _index55 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_418_v31).length));
          (_418_v31)[_index55] = (_423_v33).update(false, _415_v28);
        }
        let _424_v34;
        _424_v34 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]);
        let _425_v35;
        _425_v35 = _dafny.Seq.of((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], (_this).f15);
        let _426_v36;
        let _nw75 = new _module.C0();
        _nw75.__ctor((_this).fm0((_this).fm0((_this).f15, globalState), globalState), !((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(101)), function (_427_i7) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }), _381_v2), new BigNumber(((_424_v34).update((_dafny.ZERO).minus(new BigNumber((_425_v35).length)), (((_424_v34).contains(new BigNumber(660))) ? ((_424_v34).get(new BigNumber(660))) : ((_this).fm0((_this).f15, globalState))))).length), (_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]);
        _426_v36 = _nw75;
        let _428_v37;
        _428_v37 = _dafny.Seq.of(new BigNumber(-45), _this.f14, new BigNumber(-309), _this.f14, _this.f14);
        let _429_v38;
        _429_v38 = _dafny.MultiSet.fromElements(_381_v2);
        let _430_v39;
        _430_v39 = _dafny.MultiSet.fromElements(_this.f14);
        let _431_v40;
        _431_v40 = _module.D3.create_DC10(_381_v2, _378_v0, (_430_v39).IsSubsetOf((_dafny.MultiSet.FromArray(_428_v37)).update((((_429_v38).contains(_dafny.Seq.UnicodeFromString("bxayt"))) ? ((_429_v38).get(_dafny.Seq.UnicodeFromString("bxayt"))) : (new BigNumber(((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))]).length))), _module.__default.abs(new BigNumber(((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))]).length)))), (_dafny.ZERO).minus(_394_i2), _394_i2);
        _431_v40 = _431_v40;
        _424_v34 = (_424_v34).update(_this.f14, (_394_i2).isLessThanOrEqualTo(_this.f14));
      }
      let _432_v41;
      _432_v41 = _dafny.Seq.of(_dafny.Seq.Concat((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))], _381_v2), ((true) ? ((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))]) : ((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))])), _dafny.Seq.UnicodeFromString("urdc"), _dafny.Seq.Concat((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))], (_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))]), (_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))]);
      let _433_v42;
      _433_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(828),_dafny.Seq.Create(_module.__default.abs(new BigNumber(471)), ((_434_v2) => function (_435_i8) {
        return _434_v2;
      })(_381_v2)));
      _432_v41 = (((_433_v42).contains(_this.f14)) ? ((_433_v42).get(_this.f14)) : (_432_v41));
      let _436_v43;
      let _init15 = ((_437_v0) => function (_438_i9) {
        return (_438_i9).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), ((_439_v0) => function (_440_i10) {
          return new BigNumber((_dafny.Seq.of((_this).f15, (_439_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_439_v0).length))])).length);
        })(_437_v0))).length));
      })(_378_v0);
      let _nw76 = Array((new BigNumber(19)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw76.length); _i0_15++) {
        _nw76[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _436_v43 = _nw76;
      let _index56 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length));
      (_436_v43)[_index56] = ((_dafny.ZERO).minus(_this.f14)).plus(_this.f14);
      let _441_v44;
      _441_v44 = _dafny.Seq.of(_this.f14);
      let _442_v45;
      _442_v45 = _dafny.Set.fromElements((_441_v44)[_module.__default.safeIndex(_this.f14, new BigNumber((_441_v44).length))]);
      let _index57 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_436_v43).length));
      (_436_v43)[_index57] = new BigNumber((_442_v45).length);
      let _index58 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length));
      let _index59 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_436_v43).length));
      let _rhs51 = new BigNumber(262);
      let _rhs52 = _this.f14;
      let _rhs53 = _this.f14;
      let _rhs54 = _436_v43;
      let _rhs55 = _dafny.Seq.Concat(_441_v44, _dafny.Seq.Create(_module.__default.abs(new BigNumber(281)), function (_443_i11) {
        return _this.f14;
      }));
      let _lhs46 = _436_v43;
      let _lhs47 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length));
      let _lhs48 = _this;
      let _lhs49 = _436_v43;
      let _lhs50 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_436_v43).length));
      _lhs46[_lhs47] = _rhs51;
      _lhs48.f14 = _rhs52;
      _lhs49[_lhs50] = _rhs53;
      _436_v43 = _rhs54;
      _441_v44 = _rhs55;
      let _444_v46;
      _444_v46 = _dafny.MultiSet.fromElements((_this).f15, (_this).f15, (_this).f15, (_this).f15);
      let _445_v47;
      _445_v47 = _module.D0.create_DC2(_this.f14, _444_v46);
      let _pat_let_tv13 = _436_v43;
      let _pat_let_tv14 = _436_v43;
      let _source7 = function (_source8) {
        if (_source8.is_DC1) {
          let _446___mcc_h8 = (_source8).cf1;
          let _447_cf1 = _446___mcc_h8;
          return _module.D3.create_DC9(_447_cf1);
        } else if (_source8.is_DC2) {
          let _448___mcc_h9 = (_source8).cf2;
          let _449___mcc_h10 = (_source8).cf3;
          let _450_cf3 = _449___mcc_h10;
          let _451_cf2 = _448___mcc_h9;
          return _module.D3.create_DC9(_451_cf2);
        } else {
          let _452___mcc_h11 = (_source8).cf0;
          let _453_cf0 = _452___mcc_h11;
          return _module.D3.create_DC9((_pat_let_tv14)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_pat_let_tv13).length))]);
        }
      }(_445_v47);
      if (_source7.is_DC9) {
        let _454___mcc_h0 = (_source7).cf16;
        let _455_cf16 = _454___mcc_h0;
        let _456_v48;
        _456_v48 = _dafny.Set.fromElements((_this).f15);
        let _457_v49;
        _457_v49 = _dafny.MultiSet.fromElements(_456_v48, _456_v48, _456_v48);
        let _458_v50;
        _458_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_457_v49);
        (_this).f14 = new BigNumber((_module.__default.fm11((((_458_v50).contains((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))])) ? ((_458_v50).get((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))])) : (_457_v49)), globalState)).length);
        let _459_v51;
        _459_v51 = _dafny.MultiSet.fromElements(new BigNumber(170));
        let _460_v52;
        _460_v52 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f15);
        let _461_v53;
        _461_v53 = _dafny.Map.Empty.slice().updateUnsafe(_459_v51,_460_v52);
        _461_v53 = (_461_v53).update((_459_v51).update(_this.f14, _module.__default.abs((_dafny.ZERO).minus((((_444_v46).contains((_this).f15)) ? ((_444_v46).get((_this).f15)) : (_this.f14))))), (_460_v52).Merge(_460_v52));
        let _462_v54;
        _462_v54 = _module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(746)), ((_463_v3) => function (_464_i12) {
  return _463_v3;
})(_382_v3)));
        let _index60 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length));
        let _rhs56 = (_455_cf16).plus(_455_cf16);
        let _rhs57 = ((_456_v48).Difference(_456_v48)).Difference((_this).fm5(_462_v54, globalState));
        let _lhs51 = _436_v43;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length));
        _lhs51[_lhs52] = _rhs56;
        _456_v48 = _rhs57;
        let _465_v55;
        _465_v55 = _dafny.Seq.of((_this).f15, true, (_this).f15, (_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], (_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]);
        let _466_v56;
        _466_v56 = _dafny.Map.Empty.slice().updateUnsafe((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))],_dafny.MultiSet.FromArray(_465_v55));
        let _467_v57;
        _467_v57 = _module.D2.create_DC6(_466_v56);
        let _468_v58;
        _468_v58 = _dafny.Map.Empty.slice().updateUnsafe(_467_v57,((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]) && ((_this).f15));
        let _469_v59;
        let _nw77 = new _module.C0();
        _nw77.__ctor((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], (_this).f15, (_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))], new BigNumber(670), (_this).f15);
        _469_v59 = _nw77;
        let _470_v60;
        _470_v60 = _dafny.Map.Empty.slice().updateUnsafe((((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]) ? (_469_v59) : (_469_v59)),_468_v58);
        _468_v58 = (((_470_v60).contains(_469_v59)) ? ((_470_v60).get(_469_v59)) : (_468_v58));
      } else if (_source7.is_DC10) {
        let _471___mcc_h1 = (_source7).cf17;
        let _472___mcc_h2 = (_source7).cf18;
        let _473___mcc_h3 = (_source7).cf19;
        let _474___mcc_h4 = (_source7).cf20;
        let _475___mcc_h5 = (_source7).cf21;
        let _476_cf21 = _475___mcc_h5;
        let _477_cf20 = _474___mcc_h4;
        let _478_cf19 = _473___mcc_h3;
        let _479_cf18 = _472___mcc_h2;
        let _480_cf17 = _471___mcc_h1;
        _381_v2 = (_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))];
        (globalState).f3 = (_this).f15;
        let _481_v61;
        _481_v61 = _dafny.Set.fromElements(_479_cf18);
        let _482_v62;
        _482_v62 = _dafny.Map.Empty.slice().updateUnsafe((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))],(_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]);
        let _483_v63;
        let _nw78 = new _module.C0();
        _nw78.__ctor((_481_v61).IsDisjointFrom(_481_v61), (_this).f15, _dafny.Seq.Concat(_381_v2, _480_cf17), _476_cf21, (((_482_v62).contains(_476_cf21)) ? ((_482_v62).get(_476_cf21)) : (!(true))));
        _483_v63 = _nw78;
        let _484_v64;
        _484_v64 = _dafny.Map.Empty.slice().updateUnsafe(_478_cf19,_this.f14);
        let _485_v65;
        _485_v65 = _module.D5.create_DC13(_484_v64);
        let _index61 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length));
        let _rhs58 = (_476_cf21).minus((((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]) ? ((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))]) : (_477_cf20)));
        let _rhs59 = true;
        let _rhs60 = (_483_v63).fm2(_dafny.Seq.Concat(_381_v2, _381_v2), (_483_v63).f17, !(_484_v64).equals((_485_v65).dtor_cf24), (_dafny.ZERO).minus(_this.f14), globalState);
        let _rhs61 = _483_v63;
        let _rhs62 = _477_cf20;
        let _lhs53 = _378_v0;
        let _lhs54 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length));
        let _lhs55 = _this;
        r3 = _rhs58;
        _lhs53[_lhs54] = _rhs59;
        r3 = _rhs60;
        _483_v63 = _rhs61;
        _lhs55.f14 = _rhs62;
        _476_cf21 = new BigNumber((_480_cf17).length);
      } else if (_source7.is_DC8) {
        let _486___mcc_h6 = (_source7).cf15;
        let _487_cf15 = _486___mcc_h6;
        let _488_v66;
        _488_v66 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber((_442_v45).length));
        _488_v66 = (_488_v66).update(_this.f14, new BigNumber(855));
        _442_v45 = (((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]) ? ((_442_v45).Union(_442_v45)) : (_module.__default.fm12(_this.f14, (_this).f15, (_dafny.ZERO).minus((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))]), _382_v3, globalState)));
        let _489_v67;
        _489_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(935),(_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]);
        let _490_v68;
        let _nw79 = new _module.C0();
        _nw79.__ctor((((_489_v67).contains(new BigNumber(239))) ? ((_489_v67).get(new BigNumber(239))) : (!((_this).f15))), (((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]) ? ((_this).f15) : ((_this).f15)), _381_v2, (_dafny.ZERO).minus((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))]), (_this).f15);
        _490_v68 = _nw79;
        (globalState).f11 = (_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))];
      } else {
        let _491___mcc_h7 = (_source7).cf22;
        let _492_cf22 = _491___mcc_h7;
        let _493_v69;
        _493_v69 = _dafny.Seq.of((_this).fm0((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], globalState));
        let _494_v70;
        _494_v70 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_493_v69);
        (globalState).f11 = ((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))]).plus(new BigNumber(((((_494_v70).contains(_this.f14)) ? ((_494_v70).get(_this.f14)) : (_dafny.Seq.of((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))])))).length));
        let _495_v71;
        _495_v71 = _module.D0.create_DC0((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]);
        _495_v71 = _495_v71;
        let _496_v72;
        _496_v72 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
        _496_v72 = ((_496_v72).update((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], !((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))]))).Merge(_496_v72);
        let _497_v73;
        let _nw80 = Array((new BigNumber(20)).toNumber());
        _nw80[(_dafny.ZERO).toNumber()] = (((_this).f15) ? (_382_v3) : ((_381_v2)[_module.__default.safeIndex(_this.f14, new BigNumber((_381_v2).length))]));
        _nw80[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
        _nw80[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
        _nw80[(new BigNumber(3)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(4)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(5)).toNumber()] = ((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))])[_module.__default.safeIndex(_this.f14, new BigNumber(((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))]).length))];
        _nw80[(new BigNumber(6)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(7)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(8)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(9)).toNumber()] = (_381_v2)[_module.__default.safeIndex((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))], new BigNumber((_381_v2).length))];
        _nw80[(new BigNumber(10)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(11)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(12)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(13)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
        _nw80[(new BigNumber(15)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(16)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(17)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(18)).toNumber()] = _382_v3;
        _nw80[(new BigNumber(19)).toNumber()] = _382_v3;
        _497_v73 = _nw80;
        let _index62 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_497_v73).length));
        (_497_v73)[_index62] = _382_v3;
        let _index63 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_497_v73).length));
        (_497_v73)[_index63] = _382_v3;
      }
      r0 = (_432_v41)[_module.__default.safeIndex(_this.f14, new BigNumber((_432_v41).length))];
      let _498_v74;
      _498_v74 = _dafny.Map.Empty.slice().updateUnsafe((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))],(_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))]);
      let _499_v75;
      _499_v75 = _dafny.Map.Empty.slice().updateUnsafe((_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))],(_this).f15);
      r1 = _module.D1.create_DC5((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))], new BigNumber((_dafny.Set.fromElements((_379_v1)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_379_v1).length))])).length), (_378_v0)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_378_v0).length))], (_498_v74).update(new BigNumber((_499_v75).length), _dafny.Seq.of(_382_v3, _382_v3, new _dafny.CodePoint('d'.codePointAt(0)))));
      r2 = !((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))]).isEqualTo((_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))]);
      r3 = (_436_v43)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_436_v43).length))];
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _500_v0;
      let _nw81 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
      _500_v0 = _nw81;
      let _501_v1;
      _501_v1 = _module.D5.create_DC14(p3, _500_v0, (_this).f15);
      let _502_v2;
      _502_v2 = _dafny.Seq.of((_this).f15);
      let _503_v3;
      let _nw82 = new _module.C0();
      _nw82.__ctor((_this).f15, p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(448)), function (_504_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), new BigNumber((_502_v2).length), p1);
      _503_v3 = _nw82;
      let _505_v4;
      _505_v4 = _dafny.MultiSet.fromElements(_503_v3, _503_v3);
      let _506_v5;
      _506_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,_505_v4);
      let _507_v6;
      _507_v6 = _dafny.Seq.of(_506_v5);
      let _508_v7;
      _508_v7 = _dafny.Seq.UnicodeFromString("ejencq");
      let _509_v8;
      let _nw83 = new _module.C1();
      _nw83.__ctor(_508_v7, _502_v2, new BigNumber(335), (_503_v3).f18);
      _509_v8 = _nw83;
      let _510_v9;
      _510_v9 = _dafny.Set.fromElements(_509_v8, _509_v8);
      let _511_v10;
      let _nw84 = Array((new BigNumber(16)).toNumber());
      _nw84[(_dafny.ZERO).toNumber()] = (_this).f15;
      _nw84[(_dafny.ONE).toNumber()] = true;
      _nw84[(new BigNumber(2)).toNumber()] = (((_this).f15) ? (p1) : (!((_501_v1).dtor_cf27)));
      _nw84[(new BigNumber(3)).toNumber()] = (_this).f15;
      _nw84[(new BigNumber(4)).toNumber()] = (_this).f15;
      _nw84[(new BigNumber(5)).toNumber()] = (_this).f15;
      _nw84[(new BigNumber(6)).toNumber()] = (new BigNumber(((_507_v6)[_module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_507_v6).length))]).length)).isEqualTo(p3);
      _nw84[(new BigNumber(7)).toNumber()] = (p2).isLessThanOrEqualTo((_dafny.ZERO).minus(p3));
      _nw84[(new BigNumber(8)).toNumber()] = (_503_v3).f17;
      _nw84[(new BigNumber(9)).toNumber()] = (_502_v2)[_module.__default.safeIndex(p2, new BigNumber((_502_v2).length))];
      _nw84[(new BigNumber(10)).toNumber()] = !((_503_v3).f18);
      _nw84[(new BigNumber(11)).toNumber()] = (_this).f15;
      _nw84[(new BigNumber(12)).toNumber()] = (_510_v9).IsSubsetOf(_dafny.Set.fromElements(_509_v8, _509_v8, _509_v8, _509_v8));
      _nw84[(new BigNumber(13)).toNumber()] = (_503_v3).f17;
      _nw84[(new BigNumber(14)).toNumber()] = p1;
      _nw84[(new BigNumber(15)).toNumber()] = p1;
      _511_v10 = _nw84;
      let _index64 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
      (_511_v10)[_index64] = (_509_v8).f15;
      let _512_v11;
      let _nw85 = Array((new BigNumber(17)).toNumber()).fill([]);
      _512_v11 = _nw85;
      let _513_v12;
      _513_v12 = _dafny.MultiSet.fromElements(_this.f14, p2);
      let _514_v13;
      _514_v13 = _dafny.Map.Empty.slice().updateUnsafe(_512_v11,new BigNumber((_513_v12).cardinality()));
      let _index65 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
      let _rhs63 = p3;
      let _rhs64 = (((_514_v13).contains(_512_v11)) ? ((_514_v13).get(_512_v11)) : (p2));
      let _rhs65 = false;
      let _lhs56 = _511_v10;
      let _lhs57 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
      r0 = _rhs63;
      r1 = _rhs64;
      _lhs56[_lhs57] = _rhs65;
      let _hi2 = new BigNumber((_508_v7).length);
      for (let _515_i1 = new BigNumber((_module.__default.fm13(p3, globalState)).length); _515_i1.isLessThan(_hi2); _515_i1 = _515_i1.plus(_dafny.ONE)) {
        let _516_v14;
        let _nw86 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
        _516_v14 = _nw86;
        let _517_v15;
        let _nw87 = Array((new BigNumber(8)).toNumber());
        _nw87[(_dafny.ZERO).toNumber()] = _516_v14;
        _nw87[(_dafny.ONE).toNumber()] = _516_v14;
        _nw87[(new BigNumber(2)).toNumber()] = _516_v14;
        _nw87[(new BigNumber(3)).toNumber()] = _516_v14;
        _nw87[(new BigNumber(4)).toNumber()] = _516_v14;
        _nw87[(new BigNumber(5)).toNumber()] = _516_v14;
        _nw87[(new BigNumber(6)).toNumber()] = _516_v14;
        _nw87[(new BigNumber(7)).toNumber()] = _516_v14;
        _517_v15 = _nw87;
        let _index66 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_517_v15).length));
        (_517_v15)[_index66] = _516_v14;
        let _index67 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_517_v15).length));
        (_517_v15)[_index67] = _516_v14;
        let _518_v16;
        _518_v16 = _dafny.Set.fromElements((_503_v3).f18, (_503_v3).f18);
        let _519_v17;
        let _nw88 = new _module.C0();
        _nw88.__ctor(p1, p1, _508_v7, new BigNumber((_513_v12).cardinality()), (_dafny.Set.fromElements((_503_v3).f18)).IsSubsetOf(_518_v16));
        _519_v17 = _nw88;
        _508_v7 = _508_v7;
        (globalState).f0 = p1;
      }
      (globalState).f11 = p2;
      let _520_v18;
      _520_v18 = _module.D3.create_DC11(_module.D3.create_DC10(_508_v7, _511_v10, (_503_v3).f18, new BigNumber(502), _509_v8.f14));
      let _521_v19;
      _521_v19 = _dafny.Set.fromElements(_520_v18);
      let _522_v20;
      _522_v20 = _dafny.Map.Empty.slice().updateUnsafe(_521_v19,(_dafny.ZERO).minus(p2));
      let _523_v21;
      _523_v21 = _dafny.Map.Empty.slice().updateUnsafe((p2).multipliedBy((_dafny.ZERO).minus(p2)),!(_522_v20).equals(_522_v20));
      let _524_v22;
      _524_v22 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-673));
      let _525_v23;
      _525_v23 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_511_v10)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length))],_509_v8.f14), _524_v22);
      _523_v21 = (_523_v21).update(new BigNumber((_525_v23).cardinality()), (_503_v3).f17);
      let _526_v24;
      _526_v24 = _module.D0.create_DC1(new BigNumber((_523_v21).length));
      let _source9 = _526_v24;
      if (_source9.is_DC1) {
        let _527___mcc_h0 = (_source9).cf1;
        let _528_cf1 = _527___mcc_h0;
        let _529_v25;
        _529_v25 = _dafny.Set.fromElements((_this).f15);
        let _530_v26;
        _530_v26 = _dafny.Map.Empty.slice().updateUnsafe(_529_v25,_509_v8.f14);
        let _rhs66 = (_this.f14).isLessThanOrEqualTo(_528_cf1);
        let _rhs67 = (_511_v10)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length))];
        let _rhs68 = (p3).plus(p2);
        let _rhs69 = (_530_v26).Merge(_530_v26);
        let _lhs58 = globalState;
        let _lhs59 = globalState;
        let _lhs60 = globalState;
        _lhs58.f0 = _rhs66;
        _lhs59.f0 = _rhs67;
        _lhs60.f11 = _rhs68;
        _530_v26 = _rhs69;
        let _index68 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
        (_511_v10)[_index68] = p1;
        _524_v22 = (_524_v22).update(!((_503_v3).f17), p3);
        let _531_v27;
        let _nw89 = new _module.C0();
        _nw89.__ctor((((_523_v21).contains(_528_cf1)) ? ((_523_v21).get(_528_cf1)) : ((_509_v8).f15)), false, _508_v7, new BigNumber((_508_v7).length), !(!((_503_v3).f17)));
        _531_v27 = _nw89;
      } else if (_source9.is_DC2) {
        let _532___mcc_h1 = (_source9).cf2;
        let _533___mcc_h2 = (_source9).cf3;
        let _534_cf3 = _533___mcc_h2;
        let _535_cf2 = _532___mcc_h1;
        r1 = new BigNumber(-780);
        if (!(_535_cf2).isEqualTo((new BigNumber(219)).minus(_509_v8.f14))) {
          (globalState).f11 = ((new BigNumber((_502_v2).length)).minus(_this.f14)).multipliedBy(p3);
          let _536_v28;
          let _nw90 = Array((new BigNumber(15)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = _508_v7;
          _nw90[(_dafny.ONE).toNumber()] = _508_v7;
          _nw90[(new BigNumber(2)).toNumber()] = _508_v7;
          _nw90[(new BigNumber(3)).toNumber()] = _508_v7;
          _nw90[(new BigNumber(4)).toNumber()] = _508_v7;
          _nw90[(new BigNumber(5)).toNumber()] = _508_v7;
          _nw90[(new BigNumber(6)).toNumber()] = _508_v7;
          _nw90[(new BigNumber(7)).toNumber()] = _508_v7;
          _nw90[(new BigNumber(8)).toNumber()] = _508_v7;
          _nw90[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_508_v7, _module.__default.fm9(p2, globalState));
          _nw90[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cwdgami"), _508_v7);
          _nw90[(new BigNumber(11)).toNumber()] = _508_v7;
          _nw90[(new BigNumber(12)).toNumber()] = _508_v7;
          _nw90[(new BigNumber(13)).toNumber()] = _508_v7;
          _nw90[(new BigNumber(14)).toNumber()] = ((false) ? (_dafny.Seq.UnicodeFromString("gamojfa")) : (_508_v7));
          _536_v28 = _nw90;
          let _index69 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_536_v28).length));
          (_536_v28)[_index69] = (((_503_v3).f17) ? (_508_v7) : (_508_v7));
          let _537_v30;
          let _init16 = ((_538_v3, _539_v8, _540_p3) => function (_541_i2) {
            return (_dafny.Set.fromElements(function () {
              let _coll13 = new _dafny.Set();
              for (const _compr_13 of (_dafny.MultiSet.fromElements(_module.D2.create_DC7((_538_v3).f17, (_539_v8).f15, _540_p3), (_this).f19)).Elements) {
                let _542_v29 = _compr_13;
                if ((_dafny.MultiSet.fromElements(_module.D2.create_DC7((_538_v3).f17, (_539_v8).f15, _540_p3), (_this).f19)).contains(_542_v29)) {
                  _coll13.add(_542_v29);
                }
              }
              return _coll13;
            }())).Difference(_dafny.Set.fromElements(_dafny.Set.fromElements((_this).f19)));
          })(_503_v3, _509_v8, p3);
          let _nw91 = Array((new BigNumber(17)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw91.length); _i0_16++) {
            _nw91[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _537_v30 = _nw91;
          let _543_v31;
          _543_v31 = _dafny.Set.fromElements(_dafny.Set.fromElements(_module.D2.create_DC7(p1, (_503_v3).f17, (_dafny.ZERO).minus(_509_v8.f14)), (_this).f19));
          let _544_v32;
          _544_v32 = _module.D6.create_DC16(_543_v31);
          let _index70 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_537_v30).length));
          (_537_v30)[_index70] = (_544_v32).dtor_cf29;
          let _545_v33;
          _545_v33 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_509_v8.f14),_511_v10);
          let _index71 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_536_v28).length));
          let _index72 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_537_v30).length));
          let _rhs70 = new BigNumber((_545_v33).length);
          let _rhs71 = _dafny.Seq.UnicodeFromString("kykbqtx");
          let _rhs72 = (_543_v31).Union(_543_v31);
          let _rhs73 = _513_v12;
          let _lhs61 = globalState;
          let _lhs62 = _536_v28;
          let _lhs63 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_536_v28).length));
          let _lhs64 = _537_v30;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_537_v30).length));
          _lhs61.f11 = _rhs70;
          _lhs62[_lhs63] = _rhs71;
          _lhs64[_lhs65] = _rhs72;
          _513_v12 = _rhs73;
          let _546_v34;
          let _nw92 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Set.Empty);
          _546_v34 = _nw92;
          let _547_v35;
          _547_v35 = _dafny.MultiSet.fromElements(_546_v34, _546_v34);
          let _548_v36;
          _548_v36 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),_509_v8.f14);
          let _549_v37;
          _549_v37 = _dafny.Map.Empty.slice().updateUnsafe(p2,_548_v36);
          (_509_v8).f14 = (((_547_v35).contains(_546_v34)) ? ((_547_v35).get(_546_v34)) : (_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of((_503_v3).f18, false)).length), new BigNumber((_549_v37).length))));
          let _550_v38;
          _550_v38 = _dafny.Set.fromElements((_dafny.ZERO).minus(_509_v8.f14));
          (globalState).f0 = (_550_v38).IsProperSubsetOf((((_503_v3).f18) ? (_550_v38) : (_550_v38)));
          let _551_v39;
          _551_v39 = _module.D1.create_DC3(_508_v7);
          let _index73 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_536_v28).length));
          (_536_v28)[_index73] = _dafny.Seq.Concat((_536_v28)[_module.__default.safeIndex(new BigNumber(437), new BigNumber((_536_v28).length))], (_551_v39).dtor_cf4);
        } else {
          let _552_v40;
          let _nw93 = new _module.C1();
          _nw93.__ctor(_dafny.Seq.Concat(_508_v7, _508_v7), _502_v2, _this.f14, false);
          _552_v40 = _nw93;
          _552_v40 = _552_v40;
          let _553_v41;
          let _nw94 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _553_v41 = _nw94;
          let _index74 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_553_v41).length));
          (_553_v41)[_index74] = (((_534_cf3).contains((_509_v8).f15)) ? ((_534_cf3).get((_509_v8).f15)) : (p3));
          let _index75 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_553_v41).length));
          (_553_v41)[_index75] = (_dafny.ZERO).minus(_this.f14);
          let _554_v42;
          _554_v42 = _dafny.Map.Empty.slice().updateUnsafe(_509_v8.f14,_module.D5.create_DC13((_524_v22).update((_509_v8).f15, _509_v8.f14)));
          let _index76 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
          let _index77 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_553_v41).length));
          let _index78 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_553_v41).length));
          let _index79 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
          let _rhs74 = !(new BigNumber(65)).isEqualTo(new BigNumber((_554_v42).length));
          let _rhs75 = _module.__default.safeModuloInt((_this.f14).plus(p2), _module.__default.safeModuloInt(p3, new BigNumber(109)));
          let _rhs76 = _509_v8.f14;
          let _rhs77 = !((_503_v3).f18) || ((_503_v3).f18);
          let _rhs78 = (_535_cf2).minus((((_503_v3).f18) ? ((_503_v3).fm2((_552_v40).f20, (_503_v3).f18, (_503_v3).f18, (_dafny.ZERO).minus((_dafny.ZERO).minus(_509_v8.f14)), globalState)) : ((_dafny.ZERO).minus(new BigNumber((_523_v21).length)))));
          let _lhs66 = _511_v10;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
          let _lhs68 = _553_v41;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_553_v41).length));
          let _lhs70 = _553_v41;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_553_v41).length));
          let _lhs72 = _511_v10;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
          let _lhs74 = globalState;
          _lhs66[_lhs67] = _rhs74;
          _lhs68[_lhs69] = _rhs75;
          _lhs70[_lhs71] = _rhs76;
          _lhs72[_lhs73] = _rhs77;
          _lhs74.f11 = _rhs78;
          let _555_v43;
          _555_v43 = new _dafny.CodePoint('m'.codePointAt(0));
          let _556_v44;
          _556_v44 = _dafny.Map.Empty.slice().updateUnsafe(_509_v8.f14,(_552_v40).f20);
          let _557_v45;
          let _nw95 = Array((new BigNumber(29)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = (_552_v40).f20;
          _nw95[(_dafny.ONE).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(2)).toNumber()] = _508_v7;
          _nw95[(new BigNumber(3)).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(4)).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(518)), ((_558_v43) => function (_559_i3) {
            return _558_v43;
          })(_555_v43));
          _nw95[(new BigNumber(6)).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(7)).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(8)).toNumber()] = _508_v7;
          _nw95[(new BigNumber(9)).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(10)).toNumber()] = _508_v7;
          _nw95[(new BigNumber(11)).toNumber()] = _508_v7;
          _nw95[(new BigNumber(12)).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(13)).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(14)).toNumber()] = (_552_v40).fm8((_511_v10)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length))], globalState);
          _nw95[(new BigNumber(15)).toNumber()] = _module.__default.fm9(_509_v8.f14, globalState);
          _nw95[(new BigNumber(16)).toNumber()] = _508_v7;
          _nw95[(new BigNumber(17)).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(18)).toNumber()] = _508_v7;
          _nw95[(new BigNumber(19)).toNumber()] = _508_v7;
          _nw95[(new BigNumber(20)).toNumber()] = _508_v7;
          _nw95[(new BigNumber(21)).toNumber()] = _508_v7;
          _nw95[(new BigNumber(22)).toNumber()] = _dafny.Seq.UnicodeFromString("gb");
          _nw95[(new BigNumber(23)).toNumber()] = _dafny.Seq.UnicodeFromString("h");
          _nw95[(new BigNumber(24)).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(25)).toNumber()] = _508_v7;
          _nw95[(new BigNumber(26)).toNumber()] = (_552_v40).f20;
          _nw95[(new BigNumber(27)).toNumber()] = _dafny.Seq.update(_508_v7, _module.__default.safeIndex(new BigNumber((_556_v44).length), new BigNumber((_508_v7).length)), _555_v43);
          _nw95[(new BigNumber(28)).toNumber()] = (_552_v40).f20;
          _557_v45 = _nw95;
          let _560_v46;
          _560_v46 = _dafny.Map.Empty.slice().updateUnsafe(_555_v43,_557_v45);
          _560_v46 = (_560_v46).update(_555_v43, _557_v45);
          (globalState).f3 = (_module.__default.safeDivisionInt(_509_v8.f14, (_dafny.ZERO).minus(new BigNumber((_524_v22).length)))).isLessThanOrEqualTo((_dafny.ZERO).minus(p3));
          let _561_v47;
          _561_v47 = _module.D3.create_DC9((_this.f14).plus((_553_v41)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_553_v41).length))]));
          _561_v47 = _561_v47;
        }
        let _init17 = ((_562_v10) => function (_563_i4) {
          return (_562_v10)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_562_v10).length))];
        })(_511_v10);
        let _nw96 = Array((new BigNumber(10)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw96.length); _i0_17++) {
          _nw96[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _511_v10 = _nw96;
        let _564_v48;
        let _nw97 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _564_v48 = _nw97;
        let _index80 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_512_v11).length));
        (_512_v11)[_index80] = _564_v48;
        let _565_v49;
        _565_v49 = new _dafny.CodePoint('r'.codePointAt(0));
        let _566_v50;
        _566_v50 = _564_v48;
        let _index81 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_512_v11).length));
        let _rhs79 = new BigNumber((_508_v7).length);
        let _rhs80 = ((!(new BigNumber(887)).isEqualTo(_535_cf2)) ? ((_566_v50)) : (_564_v48));
        let _rhs81 = (_503_v3).f18;
        let _rhs82 = (_this).fm4(_dafny.areEqual(_502_v2, _502_v2), globalState);
        let _rhs83 = (_509_v8.f14).isLessThan(p3);
        let _lhs75 = _512_v11;
        let _lhs76 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_512_v11).length));
        let _lhs77 = globalState;
        let _lhs78 = globalState;
        r0 = _rhs79;
        _lhs75[_lhs76] = _rhs80;
        _lhs77.f0 = _rhs81;
        _565_v49 = _rhs82;
        _lhs78.f0 = _rhs83;
      } else {
        let _567___mcc_h3 = (_source9).cf0;
        let _568_cf0 = _567___mcc_h3;
        let _index82 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
        (_511_v10)[_index82] = _568_cf0;
        let _569_v51;
        _569_v51 = _dafny.Set.fromElements((_503_v3).f18);
        let _index83 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
        let _index84 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
        let _rhs84 = (((_this).f15) ? (_569_v51) : (_569_v51));
        let _rhs85 = ((_this).f15) && ((_509_v8).f15);
        let _rhs86 = p1;
        let _lhs79 = _511_v10;
        let _lhs80 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
        let _lhs81 = _511_v10;
        let _lhs82 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
        _569_v51 = _rhs84;
        _lhs79[_lhs80] = _rhs85;
        _lhs81[_lhs82] = _rhs86;
        let _570_v52;
        let _nw98 = new _module.C0();
        _nw98.__ctor((_509_v8).f15, (_this).f15, _508_v7, _509_v8.f14, (_503_v3).f18);
        _570_v52 = _nw98;
        let _571_v53;
        _571_v53 = _module.D6.create_DC17((_509_v8).f15, _570_v52);
        _571_v53 = _571_v53;
        let _572_v54;
        _572_v54 = new _dafny.CodePoint('i'.codePointAt(0));
        let _573_v55;
        let _nw99 = new _module.C1();
        _nw99.__ctor(_dafny.Seq.update(_508_v7, _module.__default.safeIndex(_509_v8.f14, new BigNumber((_508_v7).length)), _572_v54), _502_v2, new BigNumber((_524_v22).length), (_503_v3).f18);
        _573_v55 = _nw99;
        let _574_v56;
        _574_v56 = _dafny.Set.fromElements(_570_v52);
        let _rhs87 = _500_v0;
        let _rhs88 = new BigNumber((_574_v56).length);
        let _rhs89 = new BigNumber(((_573_v55).f21).length);
        let _rhs90 = (((_511_v10)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length))]) ? ((_573_v55).f21) : ((_573_v55).f21));
        let _rhs91 = _573_v55;
        let _lhs83 = globalState;
        let _lhs84 = globalState;
        _500_v0 = _rhs87;
        _lhs83.f11 = _rhs88;
        _lhs84.f11 = _rhs89;
        _502_v2 = _rhs90;
        _573_v55 = _rhs91;
      }
      let _575_v57;
      _575_v57 = _dafny.Seq.of(_509_v8.f14);
      let _576_v58;
      _576_v58 = _dafny.Seq.of(_575_v57, _575_v57, _575_v57);
      let _577_v59;
      _577_v59 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).fm1(_513_v12, (_this).f15, globalState)));
      let _578_v60;
      _578_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf((_576_v58)[_module.__default.safeIndex(_509_v8.f14, new BigNumber((_576_v58).length))], _dafny.Seq.update(_dafny.Seq.update(_575_v57, _module.__default.safeIndex(p2, new BigNumber((_575_v57).length)), _509_v8.f14), _module.__default.safeIndex(_this.f14, new BigNumber((_dafny.Seq.update(_575_v57, _module.__default.safeIndex(p2, new BigNumber((_575_v57).length)), _509_v8.f14)).length)), p3)),(new BigNumber(-239)).isLessThan(new BigNumber((_577_v59).length)));
      if ((((_578_v60).contains(!(false))) ? ((_578_v60).get(!(false))) : ((_509_v8).f15))) {
        let _579_v61;
        _579_v61 = _dafny.Map.Empty.slice().updateUnsafe(_523_v21,_511_v10);
        _module.__default.m0(_579_v61, globalState);
        let _580_v62;
        let _init18 = ((_581_v8) => function (_582_i5) {
          return (_582_i5).multipliedBy(_581_v8.f14);
        })(_509_v8);
        let _nw100 = Array((new BigNumber(16)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw100.length); _i0_18++) {
          _nw100[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _580_v62 = _nw100;
        let _583_v63;
        _583_v63 = _580_v62;
        _583_v63 = (((_503_v3).f18) ? (_583_v63) : (_583_v63));
        (globalState).f0 = !((_503_v3).f18);
        let _rhs92 = _509_v8.f14;
        let _rhs93 = p2;
        let _lhs85 = _509_v8;
        let _lhs86 = globalState;
        _lhs85.f14 = _rhs92;
        _lhs86.f11 = _rhs93;
        let _584_v64;
        _584_v64 = _dafny.Set.fromElements((_503_v3).f18, p1, (_503_v3).f18);
        (globalState).f0 = (_584_v64).IsProperSubsetOf(_584_v64);
      } else {
        if (!((_503_v3).f17)) {
          let _585_v65;
          let _init19 = ((_586_v3, _587_v8, _588_v2) => function (_589_i6) {
            return (_589_i6).multipliedBy(new BigNumber((_dafny.Seq.of(_dafny.Seq.of((_586_v3).f17, (_586_v3).f17, (_587_v8).f15, (_586_v3).f18, (_587_v8).f15), _588_v2)).length));
          })(_503_v3, _509_v8, _502_v2);
          let _nw101 = Array((new BigNumber(4)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw101.length); _i0_19++) {
            _nw101[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _585_v65 = _nw101;
          let _index85 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_585_v65).length));
          (_585_v65)[_index85] = _509_v8.f14;
          let _590_v66;
          _590_v66 = _dafny.Seq.of(_511_v10);
          let _591_v67;
          let _nw102 = new _module.C1();
          _nw102.__ctor(_508_v7, _502_v2, (_dafny.ZERO).minus(p3), (_this).f15);
          _591_v67 = _nw102;
          let _592_v68;
          _592_v68 = _dafny.Map.Empty.slice().updateUnsafe(_591_v67,(_511_v10)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length))]);
          let _593_v69;
          _593_v69 = _dafny.Seq.of(_592_v68, _592_v68, _592_v68, _592_v68, _592_v68);
          let _594_v70;
          _594_v70 = _dafny.Set.fromElements((_503_v3).f18);
          let _index86 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_585_v65).length));
          let _rhs94 = (_509_v8.f14).minus(_509_v8.f14);
          let _rhs95 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_590_v66, _dafny.Seq.of(_511_v10, _511_v10, _511_v10)), _590_v66)).length);
          let _rhs96 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(((_593_v69)[_module.__default.safeIndex(new BigNumber((_594_v70).length), new BigNumber((_593_v69).length))]).length)), _module.__default.safeDivisionInt(_this.f14, _this.f14)))));
          let _rhs97 = false;
          let _lhs87 = _585_v65;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_585_v65).length));
          let _lhs89 = _509_v8;
          let _lhs90 = globalState;
          _lhs87[_lhs88] = _rhs94;
          r0 = _rhs95;
          _lhs89.f14 = _rhs96;
          _lhs90.f0 = _rhs97;
          let _595_v71;
          _595_v71 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(809)).plus(p2),_523_v21);
          let _596_v73;
          _596_v73 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_509_v8.f14),(_585_v65)[_module.__default.safeIndex(new BigNumber(983), new BigNumber((_585_v65).length))]);
          let _index87 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
          let _rhs98 = (((_595_v71).contains(new BigNumber((function () {
            let _coll15 = new _dafny.Set();
            for (const _compr_15 of (_577_v59).Elements) {
              let _599_v72 = _compr_15;
              if ((_577_v59).contains(_599_v72)) {
                _coll15.add(_module.__default.safeDivisionInt(_599_v72, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(253)), function (_600_i7) {
                  return new _dafny.CodePoint('p'.codePointAt(0));
                })).length))).length)));
              }
            }
            return _coll15;
          }()).length))) ? ((_595_v71).get(new BigNumber((function () {
            let _coll14 = new _dafny.Set();
            for (const _compr_14 of (_577_v59).Elements) {
              let _597_v72 = _compr_14;
              if ((_577_v59).contains(_597_v72)) {
                _coll14.add(_module.__default.safeDivisionInt(_597_v72, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(253)), function (_598_i7) {
                  return new _dafny.CodePoint('p'.codePointAt(0));
                })).length))).length)));
              }
            }
            return _coll14;
          }()).length))) : (_523_v21));
          let _rhs99 = (new BigNumber(19)).isLessThanOrEqualTo(new BigNumber(((_596_v73).Merge(_596_v73)).length));
          let _lhs91 = _511_v10;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length));
          _523_v21 = _rhs98;
          _lhs91[_lhs92] = _rhs99;
          let _rhs100 = _509_v8.f14;
          let _rhs101 = _577_v59;
          let _rhs102 = (_503_v3).f17;
          let _lhs93 = globalState;
          let _lhs94 = globalState;
          _lhs93.f11 = _rhs100;
          _577_v59 = _rhs101;
          _lhs94.f3 = _rhs102;
          r1 = _module.__default.safeDivisionInt(p2, _this.f14);
          let _601_v74;
          _601_v74 = new _dafny.CodePoint('d'.codePointAt(0));
          let _602_v75;
          _602_v75 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(681),_601_v74);
          r1 = new BigNumber((_module.__default.fm14((_module.__default.fm15(_601_v74, p3, new BigNumber((_602_v75).length), globalState)).update((_511_v10)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length))], p2), new _dafny.CodePoint('c'.codePointAt(0)), !(!_dafny.areEqual((_591_v67).f20, _dafny.Seq.update(_508_v7, _module.__default.safeIndex((_591_v67).fm1(_513_v12, true, globalState), new BigNumber((_508_v7).length)), _601_v74))), globalState)).length);
        } else {
          let _603_v76;
          _603_v76 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_508_v7, _508_v7),_575_v57);
          _603_v76 = _603_v76;
          let _604_v77;
          _604_v77 = _dafny.Seq.of(_502_v2);
          let _605_v78;
          let _nw103 = new _module.C1();
          _nw103.__ctor(_508_v7, (_604_v77)[_module.__default.safeIndex((_526_v24).dtor_cf1, new BigNumber((_604_v77).length))], new BigNumber(169), true);
          _605_v78 = _nw103;
          let _606_v79;
          _606_v79 = new _dafny.CodePoint('h'.codePointAt(0));
          let _607_v80;
          _607_v80 = _dafny.Map.Empty.slice().updateUnsafe((_503_v3).f18,_606_v79);
          let _608_v81;
          _608_v81 = _module.D6.create_DC18(_509_v8.f14, (((_607_v80).contains((_503_v3).f18)) ? ((_607_v80).get((_503_v3).f18)) : (_606_v79)));
          let _609_v82;
          _609_v82 = _module.D6.create_DC19(_608_v81);
          let _610_v83;
          _610_v83 = _module.D6.create_DC19(_609_v82);
          let _611_v84;
          _611_v84 = _dafny.Map.Empty.slice().updateUnsafe(_605_v78,_610_v83);
          _611_v84 = (_611_v84).update((((_509_v8).f15) ? (_605_v78) : (_605_v78)), _610_v83);
          let _612_v85;
          let _nw104 = new _module.C0();
          _nw104.__ctor((_503_v3).f17, (_503_v3).f18, (_605_v78).f20, _this.f14, (_503_v3).f17);
          _612_v85 = _nw104;
          let _613_v86;
          _613_v86 = _dafny.MultiSet.fromElements((_503_v3).f18, (_612_v85).fm0((_503_v3).f18, globalState), (_509_v8).f15, (_503_v3).f18, (_this).f15);
          let _614_v87;
          _614_v87 = _dafny.Map.Empty.slice().updateUnsafe(_509_v8.f14,_613_v86);
          let _615_v88;
          _615_v88 = _module.D2.create_DC6(_614_v87);
          _615_v88 = _615_v88;
          let _616_v89;
          _616_v89 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm16(new BigNumber((_577_v59).length), globalState),_511_v10);
          _module.__default.m0(_616_v89, globalState);
        }
        (globalState).f0 = (_502_v2)[_module.__default.safeIndex((_dafny.ZERO).minus((((_513_v12).contains(new BigNumber((_508_v7).length))) ? ((_513_v12).get(new BigNumber((_508_v7).length))) : (_509_v8.f14))), new BigNumber((_502_v2).length))];
        r1 = (_509_v8.f14).minus(_this.f14);
        (globalState).f0 = (_503_v3).fm0((_511_v10)[_module.__default.safeIndex(new BigNumber(920), new BigNumber((_511_v10).length))], globalState);
        let _617_v90;
        _617_v90 = _dafny.Seq.of(_502_v2, _502_v2, _502_v2, _dafny.Seq.of((_503_v3).f18), _dafny.Seq.of((_509_v8).f15, true));
        let _618_v91;
        let _nw105 = new _module.C1();
        _nw105.__ctor(_dafny.Seq.UnicodeFromString("aqvurq"), (_617_v90)[_module.__default.safeIndex(_509_v8.f14, new BigNumber((_617_v90).length))], _509_v8.f14, (_503_v3).f18);
        _618_v91 = _nw105;
      }
      r0 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus((_509_v8).fm1(_513_v12, (_this).f15, globalState))), p3);
      let _619_v92;
      _619_v92 = _dafny.Seq.of(_dafny.Set.fromElements(p3, p2));
      r1 = _module.__default.safeModuloInt(_509_v8.f14, (new BigNumber((_619_v92).length)).minus(new BigNumber(-105)));
      return [r0, r1];
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
